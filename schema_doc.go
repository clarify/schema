package schema

import "encoding/json"

// Valid type names in JSON-schema documentation.
const (
	TypeNull    = "null"
	TypeObject  = "object"
	TypeString  = "string"
	TypeNumber  = "number"
	TypeInteger = "integer"
	TypeArray   = "array"
)

// Pre-defined formats for JSON-schema documentation.
const (
	FormatDateTime = "date-time"
	FormatDate     = "date"
	FormatTime     = "time"

	FormatEmail    = "email"
	FormatIDNEmail = "idn-email"

	FormatHostname    = "hostname"
	FormatIDNHostname = "idn-hostname"

	FormatURI          = "uri"
	FormatURIReference = "uri-reference"
	FormatURITemplate  = "uri-template"
	FormatIRI          = "iri"
	FormatIRIReference = "iri-reference"

	FormatJSONPointer         = "json-pointer"
	FormatRelativeJSONPointer = "relative-json-pointer"

	FormatRegex = "regex"
)

// Doc represents a JSON-Encodable type of format JSON Schema Draft7.
type Doc interface{}

// StringDoc implements the JSON Schema fields for JSON string types.
type StringDoc struct {
	Schema
	Type string `json:"type"`

	Format  string `json:"format,omitempty"`
	Pattern string `json:"pattern,omitempty"`

	MinLength string `json:"minLength,omitempty"`
	MaxLength string `json:"maxLength,omitempty"`
}

// NumberDoc implements the JSON Schema fields for JSON numeric types.
type NumberDoc struct {
	Schema
	Type string `json:"type"`

	MultipleOf       json.Number `json:"multipleOf,omitempty"`
	Minimum          json.Number `json:"minimum,omitempty"`
	ExclusiveMinimum json.Number `json:"exclusiveMinimum,omitempty"`
	Maximum          json.Number `json:"maximum,omitempty"`
	ExclusiveMaximum json.Number `json:"exclusiveMaximum,omitempty"`
}

// ObjectDoc implements the JSON Schema fields for JSON object types.
type ObjectDoc struct {
	Schema
	Type string `json:"type"`

	Properties           map[string]Doc `json:"properties,omitempty"`
	PatternProperties    map[string]Doc `json:"patternProperties,omitempty"`
	AdditionalProperties map[string]Doc `json:"additionalProperties,omitempty"`
	Dependency           map[string]Doc `json:"dependency,omitempty"`

	Required      []string   `json:"required"`
	PropertyNames *StringDoc `json:"propertyNames,omitempty"`
	MinProperties string     `json:"minProperties,omitempty"`
	MaxProperties string     `json:"maxProperties,omitempty"`
}

// ArrayDoc implements the JSON Schema fields for JSON array types.
type ArrayDoc struct {
	Schema
	Type string `json:"type"`

	Items                map[string]Doc `json:"properties,omitempty"`
	PatternProperties    map[string]Doc `json:"patternProperties,omitempty"`
	AdditionalProperties map[string]Doc `json:"additionalProperties,omitempty"`
	Dependency           map[string]Doc `json:"dependency,omitempty"`

	Required      []string   `json:"required"`
	PropertyNames *StringDoc `json:"propertyNames,omitempty"`
	MinProperties string     `json:"minProperties,omitempty"`
	MaxProperties string     `json:"maxProperties,omitempty"`
}
