package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// TODO: Natural language map.
// TODO: Kind serialize/deserialize use Method/Function.

const (
	// Method names for generated code
	getMethod                 = "Get"
	setMethod                 = "Set"
	hasMethod                 = "Has"
	clearMethod               = "Clear"
	iteratorClearMethod       = "clear"
	isMethod                  = "Is"
	appendMethod              = "Append"
	prependMethod             = "Prepend"
	removeMethod              = "Remove"
	lenMethod                 = "Len"
	swapMethod                = "Swap"
	lessMethod                = "Less"
	kindIndexMethod           = "kindIndex"
	serializeMethod           = "Serialize"
	deserializeMethod         = "Deserialize"
	nameMethod                = "Name"
	serializeIteratorMethod   = "serialize"
	deserializeIteratorMethod = "deserialize"
	// Member names for generated code
	unknownMemberName = "unknown"
)

// join appends a bunch of Go Code together, each on their own line.
func join(s []jen.Code) *jen.Statement {
	r := jen.Empty()
	for i, stmt := range s {
		if i > 0 {
			r.Line()
		}
		r.Add(stmt)
	}
	return r
}

// Identifier determines how a name will appear in documentation and Go code.
type Identifier struct {
	// LowerName is the typical name used in documentation.
	LowerName string
	// CamelName is the typical name used in identifiers in code.
	CamelName string
}

// Kind is data that describes a concrete Go type, how to serialize and
// deserialize such types, compare the types, and other meta-information to use
// during Go code generation.
type Kind struct {
	Name                  Identifier
	ConcreteKind          string
	Nilable               bool
	HasNaturalLanguageMap bool
	SerializeFnName       string
	DeserializeFnName     string
	LessFnName            string
}

// PropertyGenerator is a common base struct used in both Functional and
// NonFunctional ActivityStreams properties. It provides common naming patterns,
// logic, and common Go code to be generated.
//
// It also properly handles the concept of generating Go code for property
// iterators, which are needed for NonFunctional properties.
type PropertyGenerator struct {
	Package    string
	Name       Identifier
	Kinds      []Kind
	asIterator bool
}

// packageName returns the name of the package for the property to be generated.
func (p *PropertyGenerator) packageName() string {
	return p.Package
}

// structName returns the name of the type, which may or may not be a struct,
// to generate.
func (p *PropertyGenerator) structName() string {
	if p.asIterator {
		return p.Name.CamelName
	}
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
}

// propertyName returns the name of this property, as defined in
// specifications. It is not suitable for use in generated code function
// identifiers.
func (p *PropertyGenerator) propertyName() string {
	return p.Name.LowerName
}

// deserializeFnName returns the identifier of the function that deserializes
// raw JSON into the generated Go type.
func (p *PropertyGenerator) deserializeFnName() string {
	if p.asIterator {
		return fmt.Sprintf("%s%s", deserializeIteratorMethod, p.Name.CamelName)
	}
	return fmt.Sprintf("%s%sProperty", deserializeMethod, p.Name.CamelName)
}

// getFnName returns the identifier of the function that fetches concrete types
// of the property.
func (p *PropertyGenerator) getFnName(i int) string {
	if len(p.Kinds) == 1 {
		return getMethod
	}
	return fmt.Sprintf("%s%s", getMethod, p.kindCamelName(i))
}

// serializeFnName returns the identifier of the function that serializes the
// generated Go type into raw JSON.
func (p *PropertyGenerator) serializeFnName() string {
	if p.asIterator {
		return serializeIteratorMethod
	}
	return serializeMethod
}

// kindCamelName returns an identifier-friendly name for the kind at the
// specified index.
//
// It will panic if 'i' is out of range.
func (p *PropertyGenerator) kindCamelName(i int) string {
	return p.Kinds[i].Name.CamelName
}

// memberName returns the identifier to use for the kind at the specified index.
//
// It will panic if 'i' is out of range.
func (p *PropertyGenerator) memberName(i int) string {
	return fmt.Sprintf("%sMember", p.Kinds[i].Name.LowerName)
}

// hasMemberName returns the identifier to use for struct members that determine
// whether non-nilable types have been set. Panics if called for a Kind that is
// nilable.
func (p *PropertyGenerator) hasMemberName(i int) string {
	if len(p.Kinds) == 1 && p.Kinds[0].Nilable {
		panic("PropertyGenerator.hasMemberName called for nilable single value")
	}
	return fmt.Sprintf("has%sMember", p.Kinds[i].Name.CamelName)
}

// clearMethodName returns the identifier to use for methods that clear all
// values from the property.
func (p *PropertyGenerator) clearMethodName() string {
	if p.asIterator {
		return iteratorClearMethod
	}
	return clearMethod
}

// commonMethods returns methods common to every property.
func (p *PropertyGenerator) commonMethods() []*Method {
	return []*Method{
		NewCommentedValueMethod(
			p.packageName(),
			nameMethod,
			p.structName(),
			/*params=*/ nil,
			[]jen.Code{jen.String()},
			[]jen.Code{
				jen.Return(
					jen.Lit(p.propertyName()),
				),
			},
			jen.Commentf("%s returns the name of this property: %q.", nameMethod, p.propertyName()),
		),
	}
}

// isMethodName returns the identifier to use for methods that determine if a
// property holds a specific Kind of value.
func (p *PropertyGenerator) isMethodName(i int) string {
	return fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i))
}