# Schema

This repo holds a proto-type for new schema package for [rest-layer](https://github.com/rs/rest-layer).

More info on [rs/rest-layer#77](https://github.com/rs/rest-layer/issues/77)

Goals:

- Merge Schema and Field types.
- Allow [any validator type](https://github.com/rs/rest-layer/issues/77) at the root-level.
- Make deep schemas work as well as shallow (single level) ones.
- Allow [nested serialization](https://github.com/rs/rest-layer/issues/184).
- Allow simpler interface for [buypassing read-only fields](https://github.com/rs/rest-layer/issues/225).
- Avoid [data race for parallel tests](https://github.com/rs/rest-layer/issues/194) when a validator definition is reused.
- Allow a simpler interface for generating schema documentation - as JSON Schema. E.g. do so automatically on struct-inheritance.

Specifically fo references:

- Allow reference checks to operate with the [correct context](https://github.com/rs/rest-layer/issues/192).
- Allow [embedding] to rely on an interface.

The proto-type will only focus on the schema parts, not the other rest-layer packages. It will not replicate/include the `schmea/query` package.

## Example schema

TODO: All the definitions and types to construct this example is not yet in place.

```go
package main

import "github.com/searis/schema"

var userSchema = schema.Schema{
    Title: "User",
    Description: "Holds information about a user",
    Type: schema.Object{
        Required: []string{"id", "name"},
        Properties: map[string]schema.Schema{
            "id": schema.Schema{
                Type: schema.XIDField{},
                ReadOnly: true,
            },
            "name": schema.Schema{
                Type: schema.String{
                    MinLen: 1,
                    MaxLen: 255,
                },
            },
            "email": schema.Schema{
                Type: schema.Email(), // A sepcialized initalizer for String.
            },
        },
    },
}

func main() {
    // Compile the interfaces you need.
    parser := userSchema.Parser()
    validator := userSchema.Validator()
    serializer := userSchema.Serializer()

    ...
}
```

## Test format

Most unit tests are written on a [GWT format](https://www.agilealliance.org/glossary/gwt) using Go sub-tests.
