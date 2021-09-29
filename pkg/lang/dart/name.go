package dart

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	pgs "github.com/lyft/protoc-gen-star"
)

func (c context) Name(node pgs.Node) pgs.Name {
	// Message or Enum
	type ChildEntity interface {
		Name() pgs.Name
		Parent() pgs.ParentEntity
	}

	switch en := node.(type) {
	case pgs.File, pgs.Package: // the package name for this file
		return ""
	case ChildEntity: // Message or Enum types, which may be nested
		if p, ok := en.Parent().(pgs.Message); ok {
			return pgs.Name(joinChild(c.Name(p), en.Name()))
		}
		return en.Name().UpperCamelCase()
	case pgs.Field: // field names cannot conflict with other generated methods
		return replaceProtected(en.Name().LowerCamelCase())
	case pgs.OneOf: // oneof field names cannot conflict with other generated methods
		return replaceProtected(en.Name().UpperCamelCase())
	case pgs.EnumValue: // EnumValue are prefixed with the enum name
		if _, ok := en.Enum().Parent().(pgs.File); ok {
			return pgs.Name(joinNames(c.Name(en.Enum()), en.Name()))
		}
		return pgs.Name(joinNames(c.Name(en.Enum().Parent()), en.Name()))
	case pgs.Service: // always return the server name
		return c.ServerName(en)
	case pgs.Method:
		return en.Name().LowerCamelCase()
	case pgs.Entity: // any other entity should be just upper-camel-cased
		return en.Name().UpperCamelCase()
	default:
		panic("unreachable")
	}
}

func (c context) OneofOption(field pgs.Field) pgs.Name {
	n := pgs.Name(joinNames(c.Name(field.Message()), c.Name(field)))

	for _, msg := range field.Message().Messages() {
		if c.Name(msg) == n {
			return n + "_"
		}
	}

	for _, en := range field.Message().Enums() {
		if c.Name(en) == n {
			return n + "_"
		}
	}

	return n
}

func (c context) ServerName(s pgs.Service) pgs.Name {
	n := s.Name().UpperCamelCase()
	return pgs.Name(fmt.Sprintf("%sServer", n))
}

func (c context) ClientName(s pgs.Service) pgs.Name {
	n := s.Name().UpperCamelCase()
	return pgs.Name(fmt.Sprintf("%sClient", n))
}

func (c context) ServerStream(m pgs.Method) pgs.Name {
	s := m.Service().Name().UpperCamelCase()
	n := m.Name().UpperCamelCase()
	return joinNames(s, n) + "Server"
}

var protectedNames = map[pgs.Name]pgs.Name{
	"Reset":               "Reset_",
	"String":              "String_",
	"ProtoMessage":        "ProtoMessage_",
	"Marshal":             "Marshal_",
	"Unmarshal":           "Unmarshal_",
	"ExtensionRangeArray": "ExtensionRangeArray_",
	"ExtensionMap":        "ExtensionMap_",
	"Descriptor":          "Descriptor_",
}

func replaceProtected(n pgs.Name) pgs.Name {
	if use, protected := protectedNames[n]; protected {
		return use
	}
	return n
}

func joinChild(a, b pgs.Name) pgs.Name {
	if r, _ := utf8.DecodeRuneInString(b.String()); unicode.IsLetter(r) && unicode.IsLower(r) {
		return pgs.Name(fmt.Sprintf("%s%s", a, b.UpperCamelCase()))
	}
	return joinNames(a, b.UpperCamelCase())
}

func joinNames(a, b pgs.Name) pgs.Name {
	return pgs.Name(fmt.Sprintf("%s_%s", a, b))
}
