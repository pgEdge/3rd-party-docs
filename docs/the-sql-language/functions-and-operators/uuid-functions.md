## UUID Functions { #functions-uuid }


 PostgreSQL includes one function to generate a UUID:

```

gen_random_uuid () uuid
```
 This function returns a version 4 (random) UUID. This is the most commonly used type of UUID and is appropriate for most applications.


 The [uuid-ossp](../../appendixes/additional-supplied-modules-and-extensions/uuid-ossp-a-uuid-generator.md#uuid-ossp) module provides additional functions that implement other standard algorithms for generating UUIDs.


 PostgreSQL also provides the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) for UUIDs.
