package schema

// Compile-time checks.
var (
	_ Parser     = ParserFunc(nil)
	_ Serializer = SerializerFunc(nil)
	_ Validator  = ValidatorFunc(nil)
	_ Lesser     = LesserFunc(nil)
)

// O is a type-alias for payloads of type JSON Objects.
type O = map[string]interface{}

// A is a type-alias for payloads of type JSON Arrays.
type A = []interface{}

// DocType describes a type that can produce JSON-Schema documentation.
type DocType interface {
	// Doc should return a JSON-encodable type that includes fields from s.
	Doc(s Schema) Doc
}

// ElementType is the interface for any Type that may allow a JSON Array
// payload.
type ElementType interface {
	ElementSchema(i int) *Schema
}

// PropertyType is the interface for any Type that may allow a JSON Object
// payload.
type PropertyType interface {
	PropertySchema(name string) *Schema
}

// Type is the interface for the type-specific parts of a schema. It should be
// JSON marshaled to JSON Schema Draft v7.
type Type interface {
	Parser() Parser
	Validator() Validator
	Serializer() Serializer
}

// Parser allows converting an external representation types to the internall
// representation types. A parser should only handles type conversion, and never
// perform any semantic validation.
type Parser interface {
	Parse(in interface{}) (out interface{}, err error)
}

// ParserFunc allows a function to implement Parser. ParserFunc(nil) is also
// valid, and implements a no-op Parser.
type ParserFunc func(in interface{}) (out interface{}, err error)

func (f ParserFunc) Parse(in interface{}) (out interface{}, err error) {
	if f == nil {
		return in, nil
	}
	return f(in)
}

// Serializer allows converting an internal representation type to an external
// one. May return an Omit{} value if the content should not be exported. nil
// on the other hand is considered a value to be exported.
type Serializer interface {
	Serialize(in interface{}) (out interface{}, err error)
}

// SerializerFunc allows a function to implement Serializer. SerializerFunc(nil)
// is also valid, and implements a no-op Serializer.
type SerializerFunc func(in interface{}) (out interface{}, err error)

func (f SerializerFunc) Serialize(in interface{}) (out interface{}, err error) {
	if f == nil {
		return in, nil
	}
	return f(in)
}

// Validator validates an internal representation. This involves validating that
// the input is of the correct type. original can be ignored in most-cases, but
// must be passed-along when validating sub-schemas.
type Validator interface {
	Validate(in, original interface{}) (out interface{}, err error)
}

// ValidatorFunc allows a function to implement Validator. ValidatorFunc(nil)
// is also valid, and implements a no-op Validator.
type ValidatorFunc func(in, original interface{}) (out interface{}, err error)

func (f ValidatorFunc) Validate(in, original interface{}) (out interface{}, err error) {
	if f == nil {
		return in, nil
	}
	return f(in, original)
}

// LesserType is a type that allows le,lt,gt and ge comparisons of internally
// represented (parsed) values.
type LesserType interface {
	Lesser() Lesser
}

// Lesser allows comparing two different internally represented (parsed) values
// of the same type.
type Lesser interface {
	// Less returns true if both a and b are of the correct type, and a is less
	// than b.
	Less(a, b interface{}) bool
}

// LesserFunc allows a function to implement Lesser. Note that LesserFunc(nil)
// is NOT valid and wil panic if called.
type LesserFunc func(a, b interface{}) bool

func (f LesserFunc) Less(a, b interface{}) bool {
	return f(a, b)
}
