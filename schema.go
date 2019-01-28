package schema

// Omit indicates that a payload value is/should be omitted.
type Omit struct{}

// SkipReadOnly can be used within an un-validated payload containing the
// internal representation to explicitly skip checks on the ReadOnly and
// CreateOnly Schema properties.
type SkipReadOnly struct {
	Value interface{}
}

// Schema lets you define a schema for performing payload parsing, validation
// and serialization. It is JSON marshaled to JSON Schema Draft v7.
type Schema struct {
	Type
	Default     interface{}       `json:"default,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	ReadOnly    bool              `json:"readOnly,omitempty"`
	WriteOnly   bool              `json:"writeOnly,omitempty"`
	Dependency  map[string]Schema `json:"dependency,omitempty"`
	CreateOnly  bool              `json:"-"`
}

// Parser returns the ParserFunc for the schema.
func (s *Schema) Parser() ParserFunc {
	if s.Type == nil {
		return func(in interface{}) (interface{}, error) {
			return in, nil
		}
	}
	return s.Type.Parser()
}

// Validate returns the ValidatorFunc for the schema.
func (s *Schema) Validate() ValidatorFunc {
	var typeValidator ValidatorFunc
	if s.Type != nil {
		typeValidator = s.Type.Validator()
	}
	if s.ReadOnly {
		return func(in, original interface{}) (interface{}, error) {
			if skip, ok := in.(SkipReadOnly); ok {
				in = skip.Value
			} else if in != original {
				return nil, ErrCreateOnly
			}
			if typeValidator == nil {
				return in, nil
			}
			return typeValidator(in, original)
		}
	} else if s.CreateOnly {
		return func(in, original interface{}) (interface{}, error) {
			if skip, ok := in.(SkipReadOnly); ok {
				in = skip.Value
			} else if in != original {
				return nil, ErrCreateOnly
			}
			if typeValidator == nil {
				return in, nil
			}
			return typeValidator(in, original)
		}
	}

	if typeValidator == nil {
		return func(in, original interface{}) (interface{}, error) {
			return in, nil
		}
	}
	return typeValidator
}

// Serializer returns the SerializerFunc for the schema.
func (s *Schema) Serializer() SerializerFunc {
	if s.WriteOnly {
		return func(in interface{}) (interface{}, error) {
			return Omit{}, nil
		}
	}
	if s.Type == nil {
		return func(in interface{}) (interface{}, error) {
			return in, nil
		}
	}
	return s.Type.Serializer()
}
