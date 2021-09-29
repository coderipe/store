package dart

import (
	"bytes"
	pgs "github.com/lyft/protoc-gen-star"
	"os/exec"
	"strings"
)

type dartFmt struct{}

// DartFmt returns a PostProcessor that runs gofmt on any files ending in ".go"
func DartFmt() pgs.PostProcessor { return dartFmt{} }

func (p dartFmt) Match(a pgs.Artifact) bool {
	var n string

	switch a := a.(type) {
	case pgs.GeneratorFile:
		n = a.Name
	case pgs.GeneratorTemplateFile:
		n = a.Name
	case pgs.CustomFile:
		n = a.Name
	case pgs.CustomTemplateFile:
		n = a.Name
	default:
		return false
	}

	return strings.HasSuffix(n, ".dart")
}

func (p dartFmt) Process(in []byte) ([]byte, error) {
	cmd := exec.Command("dart", "format", "-o", "show")
	cmd.Stdin = bytes.NewReader(in)
	return cmd.Output()
}

var _ pgs.PostProcessor = dartFmt{}