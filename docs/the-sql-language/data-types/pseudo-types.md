<a id="datatype-pseudo"></a>

## Pseudo-Types


 The PostgreSQL type system contains a number of special-purpose entries that are collectively called *pseudo-types*. A pseudo-type cannot be used as a column data type, but it can be used to declare a function's argument or result type. Each of the available pseudo-types is useful in situations where a function's behavior does not correspond to simply taking or returning a value of a specific SQL data type. [Pseudo-Types](#datatype-pseudotypes-table) lists the existing pseudo-types.
 <a id="datatype-pseudotypes-table"></a>

**Table: Pseudo-Types**

| Name | Description |
| --- | --- |
| `any` | Indicates that a function accepts any input data type. |
| `anyelement` | Indicates that a function accepts any data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic)). |
| `anyarray` | Indicates that a function accepts any array data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic)). |
| `anynonarray` | Indicates that a function accepts any non-array data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic)). |
| `anyenum` | Indicates that a function accepts any enum data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic) and [Enumerated Types](enumerated-types.md#datatype-enum)). |
| `anyrange` | Indicates that a function accepts any range data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic) and [Range Types](range-types.md#rangetypes)). |
| `anymultirange` | Indicates that a function accepts any multirange data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic) and [Range Types](range-types.md#rangetypes)). |
| `anycompatible` | Indicates that a function accepts any data type, with automatic promotion of multiple arguments to a common data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic)). |
| `anycompatiblearray` | Indicates that a function accepts any array data type, with automatic promotion of multiple arguments to a common data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic)). |
| `anycompatiblenonarray` | Indicates that a function accepts any non-array data type, with automatic promotion of multiple arguments to a common data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic)). |
| `anycompatiblerange` | Indicates that a function accepts any range data type, with automatic promotion of multiple arguments to a common data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic) and [Range Types](range-types.md#rangetypes)). |
| `anycompatiblemultirange` | Indicates that a function accepts any multirange data type, with automatic promotion of multiple arguments to a common data type (see [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic) and [Range Types](range-types.md#rangetypes)). |
| `cstring` | Indicates that a function accepts or returns a null-terminated C string. |
| `internal` | Indicates that a function accepts or returns a server-internal data type. |
| `language_handler` | A procedural language call handler is declared to return `language_handler`. |
| `fdw_handler` | A foreign-data wrapper handler is declared to return `fdw_handler`. |
| `table_am_handler` | A table access method handler is declared to return `table_am_handler`. |
| `index_am_handler` | An index access method handler is declared to return `index_am_handler`. |
| `tsm_handler` | A tablesample method handler is declared to return `tsm_handler`. |
| `record` | Identifies a function taking or returning an unspecified row type. |
| `trigger` | A trigger function is declared to return `trigger.` |
| `event_trigger` | An event trigger function is declared to return `event_trigger.` |
| `pg_ddl_command` | Identifies a representation of DDL commands that is available to event triggers. |
| `void` | Indicates that a function returns no value. |
| `unknown` | Identifies a not-yet-resolved type, e.g., of an undecorated string literal. |


 Functions coded in C (whether built-in or dynamically loaded) can be declared to accept or return any of these pseudo-types. It is up to the function author to ensure that the function will behave safely when a pseudo-type is used as an argument type.


 Functions coded in procedural languages can use pseudo-types only as allowed by their implementation languages. At present most procedural languages forbid use of a pseudo-type as an argument type, and allow only `void` and `record` as a result type (plus `trigger` or `event_trigger` when the function is used as a trigger or event trigger). Some also support polymorphic functions using the polymorphic pseudo-types, which are shown above and discussed in detail in [Polymorphic Types](../../server-programming/extending-sql/the-postgresql-type-system.md#extend-types-polymorphic).


 The `internal` pseudo-type is used to declare functions that are meant only to be called internally by the database system, and not by direct invocation in an SQL query. If a function has at least one `internal`-type argument then it cannot be called from SQL. To preserve the type safety of this restriction it is important to follow this coding rule: do not create any function that is declared to return `internal` unless it has at least one `internal` argument.
