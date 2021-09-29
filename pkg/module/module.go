package module

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/srikrsna/protoc-gen-store/pkg/lang/dart"

	pgs "github.com/lyft/protoc-gen-star"
	storepb "github.com/srikrsna/protoc-gen-store/store"
)

type Module struct {
	*pgs.ModuleBase

	dartCtx dart.Context
}

func New() *Module {
	return &Module{ModuleBase: &pgs.ModuleBase{}}
}

func (m *Module) Name() string {
	return "store"
}

func (m *Module) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.dartCtx = dart.InitContext(c.Parameters())
}

func (m *Module) Execute(files map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {

	buf := &bytes.Buffer{}
	for _, file := range files {
		if len(file.Services()) == 0 {
			continue
		}

		importMap := map[string]pgs.Entity{}
		buf.Reset()
		for _, service := range file.Services() {
			fmt.Fprintf(buf, "class %sBloc {\n", m.dartCtx.Name(service))
			fmt.Fprintf(buf, "final %s api;\n", m.dartCtx.ClientName(service))
			fmt.Fprintf(buf, "final Store store;\n")
			fmt.Fprintf(buf, "%sBloc(this.store, this.api);\n\n", m.dartCtx.Name(service))
			for _, method := range service.Methods() {
				if method.ServerStreaming() || method.ClientStreaming() {
					continue
				}

				var (
					req = method.Input()
					res = method.Output()
				)

				for _, f := range req.Imports() {
					importMap[string(f.Name())] = f
				}
				for _, f := range method.Imports() {
					importMap[string(f.Name())] = f
				}

				var sigOpts storepb.MethodSignature
				if _, err := method.Extension(storepb.E_MethodSignature, &sigOpts); err != nil {
					m.Fail(err)
				}

				requiredFields := strings.Split(sigOpts.Required, ",")
				requiredFieldMap := map[string]bool{}
				for _, rf := range requiredFields {
					requiredFieldMap[rf] = true
				}

				// Method Signture
				fmt.Fprintf(buf, "Future<Result<%s>> %s({\n", m.dartCtx.Name(res), m.dartCtx.Name(method))
				for _, field := range req.Fields() {
					if requiredFieldMap[field.Name().String()] {
						fmt.Fprintf(buf, "required %s %s,\n", m.dartCtx.Type(field).NonNullable(), m.dartCtx.Name(field))
					} else {
						fmt.Fprintf(buf, "%s %s,\n", m.dartCtx.Type(field).Nullable(), m.dartCtx.Name(field))
					}
				}
				fmt.Fprintln(buf, "}) async {")

				// Body

				// Request
				fmt.Fprintf(buf, "final req = %s(\n", m.dartCtx.Name(req))
				for _, field := range req.Fields() {
					fmt.Fprintf(buf, "%[1]s: %[1]s,\n", m.dartCtx.Name(field))
				}
				fmt.Fprintln(buf, ");")

				// Refs
				fmt.Fprintf(buf, "final baseRef = '%s(${req.hashCode})';\n", m.dartCtx.Name(method))
				fmt.Fprintln(buf, "final loadingRef = '$baseRef:loading';")
				fmt.Fprintln(buf, "final errorRef = '$baseRef:error';")

				// CustomRefCode
				fmt.Fprintf(buf, "\n\n// @@protoc_insertion_point(on_start:%s)\n\n", method.FullyQualifiedName())

				// Set Loading
				fmt.Fprintln(buf, "store.publishRecord(loadingRef, true);")
				fmt.Fprintln(buf, "store.publishRecord<GrpcError?>(errorRef, null);")
				fmt.Fprintln(buf, "store.notify();")

				// Call
				fmt.Fprintln(buf, "try {")
				fmt.Fprintf(buf, "final res = await api.%s(req);\n", m.dartCtx.Name(method))
				fmt.Fprintf(buf, "store.publishRecord<%s>(baseRef, res);\n", m.dartCtx.Name(res))
				fmt.Fprintf(buf, "\n\n// @@protoc_insertion_point(on_success:%s)\n\n", method.FullyQualifiedName())
				fmt.Fprintln(buf, "return Result(res, null);")
				fmt.Fprintln(buf, "} on GrpcError catch (err) {")
				fmt.Fprintln(buf, "store.publishRecord<GrpcError>(errorRef, err);")
				fmt.Fprintf(buf, "\n\n// @@protoc_insertion_point(on_error:%s)\n\n", method.FullyQualifiedName())
				fmt.Fprintln(buf, "return Result(null, err);")
				fmt.Fprintln(buf, "} finally {")
				fmt.Fprintf(buf, "\n\n// @@protoc_insertion_point(on_end:%s)\n\n", method.FullyQualifiedName())
				fmt.Fprintln(buf, "store.publishRecord(loadingRef, false);")
				fmt.Fprintln(buf, "store.notify();")
				fmt.Fprintln(buf, "}")

				fmt.Fprintln(buf, "}")
			}

			fmt.Fprintln(buf, "}")
		}

		impBuf := &bytes.Buffer{}
		impBuf.WriteString(`
			import 'package:grpc/grpc.dart';
			import 'package:sane/sane.dart';	
			import 'package:fixnum/fixnum.dart';		
		`)

		importPrefix := strings.Repeat("../", len(strings.Split(string(m.dartCtx.OutputPath(file)), "/"))-1)

		for _, f := range importMap {
			fmt.Fprintf(impBuf, "import '%s%s';\n", importPrefix, m.dartCtx.OutputPath(f))
		}

		fmt.Fprintf(impBuf, "import '%s%s';\n", importPrefix, m.dartCtx.OutputPath(file).SetExt("grpc.dart"))
		fmt.Fprintf(impBuf, "\n\n// @@protoc_insertion_point(imports:%s)\n\n", file.FullyQualifiedName())

		m.AddGeneratorFile(string(file.InputPath().SetExt(".pbstore.dart")), impBuf.String()+buf.String())
	}

	return m.Artifacts()
}
