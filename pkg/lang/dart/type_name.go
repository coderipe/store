package dart

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func (c context) Type(f pgs.Field) TypeName {
	ft := f.Type()

	var t TypeName
	switch {
	case ft.IsMap():
		key := scalarType(ft.Key().ProtoType())
		return TypeName(fmt.Sprintf("Map<%s,%s>", key, c.elType(ft)))
	case ft.IsRepeated():
		return TypeName(fmt.Sprintf("Iterable<%s>", c.elType(ft)))
	case ft.IsEmbed():
		return c.importableTypeName(f, ft.Embed())
	case ft.IsEnum():
		t = c.importableTypeName(f, ft.Enum())
	default:
		t = scalarType(ft.ProtoType())
	}

	if f.HasPresence() && f.HasOptionalKeyword() {
		return t.Nullable()
	}

	return t
}

func (c context) importableTypeName(f pgs.Field, e pgs.Entity) TypeName {
	return TypeName(c.Name(e))
}

func (c context) elType(ft pgs.FieldType) TypeName {
	el := ft.Element()
	switch {
	case el.IsEnum():
		return c.importableTypeName(ft.Field(), el.Enum())
	case el.IsEmbed():
		return c.importableTypeName(ft.Field(), el.Embed())
	default:
		return scalarType(el.ProtoType())
	}
}

func scalarType(t pgs.ProtoType) TypeName {
	switch t {
	case pgs.DoubleT:
		return "double"
	case pgs.FloatT:
		return "double"
	case pgs.Int64T, pgs.SFixed64, pgs.SInt64:
		return "Int64"
	case pgs.UInt64T, pgs.Fixed64T:
		return "Int64"
	case pgs.Int32T, pgs.SFixed32, pgs.SInt32:
		return "int"
	case pgs.UInt32T, pgs.Fixed32T:
		return "int"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "String"
	case pgs.BytesT:
		return "List<int>"
	default:
		panic("unreachable: invalid scalar type")
	}
}

// A TypeName describes the name of a type (type on a field, or method signature)
type TypeName string

// String satisfies the strings.Stringer interface.
func (n TypeName) String() string { return string(n) }

// Element returns the TypeName of the element of n. For types other than
// lists and maps, this just returns n.
func (n TypeName) Element() TypeName {
	if strings.HasPrefix(string(n), "Map") {
		s := strings.IndexRune(string(n), ',')
		e := strings.IndexRune(string(n), '>')
		if s < 0 || e < 0 {
			return ""
		}

		return TypeName(strings.TrimSpace(string(n[s+1 : e])))
	}

	if strings.HasPrefix(string(n), "List") || strings.HasPrefix(string(n), "Iterable") {
		s := strings.IndexRune(string(n), '<')
		e := strings.IndexRune(string(n), '>')
		if s < 0 || e < 0 {
			return ""
		}

		return TypeName(strings.TrimSpace(string(n[s+1 : e])))
	}

	return ""
}

// Key returns the TypeName of the key of n. For lists, the return TypeName is
// always "int", and for non slice/map types an empty TypeName is returned.
func (n TypeName) Key() TypeName {
	if strings.HasPrefix(string(n), "Map") {
		s := strings.IndexRune(string(n), '<')
		e := strings.IndexRune(string(n), ',')
		if s < 0 || e < 0 {
			return ""
		}

		return TypeName(strings.TrimSpace(string(n[s+1 : e])))
	}

	if strings.HasPrefix(string(n), "List") || strings.HasPrefix(string(n), "Iterable") {
		return "int"
	}

	return ""
}

// IsNullable reports whether TypeName n is a nullable type
func (n TypeName) IsNullable() bool {
	ns := string(n)
	return strings.HasSuffix(ns, "?")
}

// Nullable converts TypeName n to it's nullable type. If n is already a nullable,
// slice, or map, it is returned unmodified.
func (n TypeName) Nullable() TypeName {
	if n.IsNullable() {
		return n
	}
	return TypeName(string(n) + "?")
}

// NonNullable converts TypeName n to it's non nullable type. If n is already a non nullable
// type it is returned unmodified.
func (n TypeName) NonNullable() TypeName {
	return TypeName(strings.TrimSuffix(string(n), "?"))
}
