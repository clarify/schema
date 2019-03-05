package schema

// Omit indicates that a payload value should not be included after
// serialization or validation, except for at the top-level of a (nested)
// schema.
type Omit struct{}

// SkipReadOnly can be used within an un-validated payload containing the
// internal representation to explicitly skip checks on the ReadOnly and
// CreateOnly Schema properties.
type SkipReadOnly struct {
	Value interface{}
}

// Schema lets you define a schema for performing payload parsing, validation
// and serialization.
type Schema struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Type        Type   `json:"-"`

	CreateOnly bool `json:"-"`
	ReadOnly   bool `json:"readOnly,omitempty"`
	WriteOnly  bool `json:"writeOnly,omitempty"`

	Default interface{} `json:"default,omitempty"`
}

// Doc returns a JSON-encodable type describing the schema. The current form
// is based on JSON Schema Draft 7.
func (s Schema) Doc() Doc {
	if td, ok := s.Type.(DocType); ok {
		return td.Doc(s)
	}
	return s
}

// Parser returns the ParserFunc for the schema.
func (s Schema) Parser() Parser {
	if s.Type == nil {
		return ParserFunc(nil)
	}
	return s.Type.Parser()
}

// Validator returns the ValidatorFunc for the schema.
func (s Schema) Validator() Validator {
	var val Validator
	if s.Type != nil {
		val = s.Type.Validator()
	} else {
		val = ValidatorFunc(nil)
	}

	if s.ReadOnly {
		val = ValidatorFunc(func(in, original interface{}) (interface{}, error) {
			if skip, ok := in.(SkipReadOnly); ok {
				in = skip.Value
			} else if in != original {
				return nil, ErrCreateOnly
			}
			return val.Validate(in, original)
		})
	} else if s.CreateOnly {
		val = ValidatorFunc(func(in, original interface{}) (interface{}, error) {
			if skip, ok := in.(SkipReadOnly); ok {
				in = skip.Value
			} else if in != original {
				return nil, ErrCreateOnly
			}
			return val.Validate(in, original)
		})
	}

	return val
}

// Serializer returns the SerializerFunc for the schema.
func (s Schema) Serializer() Serializer {
	if s.WriteOnly {
		return SerializerFunc(func(in interface{}) (interface{}, error) {
			return Omit{}, nil
		})
	}
	if s.Type == nil {
		return SerializerFunc(nil)
	}
	return s.Type.Serializer()
}
