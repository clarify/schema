package schema

import "errors"

// Commone schema validation errors.
var (
	ErrReadOnly   = errors.New("read-only")
	ErrCreateOnly = errors.New("create-only")
)

// O is a type-alias for payloads of type JSON Objects.
type O = map[string]interface{}

// A is a type-alias for payloads of type JSON Arrays.
type A = []interface{}

// Type is the interface for the type-specific parts of a schema. It should be
// JSON marshaled to JSON Schema Draft v7.
type Type interface {
	Type() string
	Parser() ParserFunc
	Validator() ValidatorFunc
	Serializer() SerializerFunc
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

// ParserFunc is a function that converts external representation types to
// internall representation types.
type ParserFunc func(in interface{}) (out interface{}, err error)

// ValidatorFunc is a function that validates internal representation types. If
// original is not Omit, the function should validate that the described change
// are allowed.
type ValidatorFunc func(in, original interface{}) (out interface{}, err error)

// SerializerFunc is a function that converts an internal representation type to
// an external one.
type SerializerFunc func(in interface{}) (out interface{}, err error)
