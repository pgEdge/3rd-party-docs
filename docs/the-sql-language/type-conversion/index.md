<a id="typeconv"></a>

# Type Conversion

 SQL statements can, intentionally or not, require the mixing of different data types in the same expression. PostgreSQL has extensive facilities for evaluating mixed-type expressions.

 In many cases a user does not need to understand the details of the type conversion mechanism. However, implicit conversions done by PostgreSQL can affect the results of a query. When necessary, these results can be tailored by using *explicit* type conversion.

 This chapter introduces the PostgreSQL type conversion mechanisms and conventions. Refer to the relevant sections in [Data Types](../data-types/index.md#datatype) and [Functions and Operators](../functions-and-operators/index.md#functions) for more information on specific data types and allowed functions and operators.

- [Overview](overview.md#typeconv-overview)
- [Operators](operators.md#typeconv-oper)
- [Functions](functions.md#typeconv-func)
- [Value Storage](value-storage.md#typeconv-query)
- [`UNION`, `CASE`, and Related Constructs](union-case-and-related-constructs.md#typeconv-union-case)
- [`SELECT` Output Columns](select-output-columns.md#typeconv-select)
