# schmea
Proto-type for new schema package for [rest-layer](https://github.com/rs/rest-layer).

More info on [rs/rest-layer#77](https://github.com/rs/rest-layer/issues/77)

Goals:
- Merge Schema and Field types.
- Allow [any validator type](https://github.com/rs/rest-layer/issues/77) at the root-level.
- Make deep schemas work as well as shallow (single level) ones.
- Allow [nested serialization](https://github.com/rs/rest-layer/issues/184).
- Allow simpler interface for [buypassing read-only fields](https://github.com/rs/rest-layer/issues/225).
- Avoid [data race for parallel tests](https://github.com/rs/rest-layer/issues/194) when a validator definintion is reused.
- Allow a simpler interface for generating schema documentation - as JSON Schema. E.g. do so automatically on struct-inheritance.

Specifically fo references:
- Allow reference checks to operate with the [correct context](https://github.com/rs/rest-layer/issues/192).
- Allow [embedding] to rely on an interface.

The proto-type will only focus on the schema parts, not the other rest-layer packages. It will not replicate/include the `schmea/query` package.
