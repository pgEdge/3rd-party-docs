<a id="datatype-boolean"></a>

## Boolean Type


 PostgreSQL provides the standard SQL type `boolean`; see [Boolean Data Type](#datatype-boolean-table). The `boolean` type can have several states: “true”, “false”, and a third state, “unknown”, which is represented by the SQL null value.
 <a id="datatype-boolean-table"></a>

**Table: Boolean Data Type**

| Name | Storage Size | Description |
| --- | --- | --- |
| `boolean` | 1 byte | state of true or false |


 Boolean constants can be represented in SQL queries by the SQL key words `TRUE`, `FALSE`, and `NULL`.


 The datatype input function for type `boolean` accepts these string representations for the “true” state:

- `true`
- `yes`
- `on`
- `1`
 and these representations for the “false” state:

- `false`
- `no`
- `off`
- `0`
 Unique prefixes of these strings are also accepted, for example `t` or `n`. Leading or trailing whitespace is ignored, and case does not matter.


 The datatype output function for type `boolean` always emits either `t` or `f`, as shown in [Using the `boolean` Type](#datatype-boolean-example).
 <a id="datatype-boolean-example"></a>

**Example: Using the `boolean` Type**


```sql

CREATE TABLE test1 (a boolean, b text);
INSERT INTO test1 VALUES (TRUE, 'sic est');
INSERT INTO test1 VALUES (FALSE, 'non est');
SELECT * FROM test1;
 a |    b
---+---------
 t | sic est
 f | non est

SELECT * FROM test1 WHERE a;
 a |    b
---+---------
 t | sic est
```


 The key words `TRUE` and `FALSE` are the preferred (SQL-compliant) method for writing Boolean constants in SQL queries. But you can also use the string representations by following the generic string-literal constant syntax described in [Constants of Other Types](../sql-syntax/lexical-structure.md#sql-syntax-constants-generic), for example `'yes'::boolean`.


 Note that the parser automatically understands that `TRUE` and `FALSE` are of type `boolean`, but this is not so for `NULL` because that can have any type. So in some contexts you might have to cast `NULL` to `boolean` explicitly, for example `NULL::boolean`. Conversely, the cast can be omitted from a string-literal Boolean value in contexts where the parser can deduce that the literal must be of type `boolean`.
