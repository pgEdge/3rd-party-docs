<a id="functions-comparison"></a>

## Comparison Functions and Operators


 The usual comparison operators are available, as shown in [Comparison Operators](#functions-comparison-op-table).
 <a id="functions-comparison-op-table"></a>

**Table: Comparison Operators**

| Operator | Description |
| --- | --- |
| *datatype* `<` *datatype* `boolean` | Less than |
| *datatype* `>` *datatype* `boolean` | Greater than |
| *datatype* `<=` *datatype* `boolean` | Less than or equal to |
| *datatype* `>=` *datatype* `boolean` | Greater than or equal to |
| *datatype* `=` *datatype* `boolean` | Equal |
| *datatype* `<>` *datatype* `boolean` | Not equal |
| *datatype* `!=` *datatype* `boolean` | Not equal |


!!! note

    `<>` is the standard SQL notation for “not equal”. `!=` is an alias, which is converted to `<>` at a very early stage of parsing. Hence, it is not possible to implement `!=` and `<>` operators that do different things.


 These comparison operators are available for all built-in data types that have a natural ordering, including numeric, string, and date/time types. In addition, arrays, composite types, and ranges can be compared if their component data types are comparable.


 It is usually possible to compare values of related data types as well; for example `integer` `>` `bigint` will work. Some cases of this sort are implemented directly by “cross-type” comparison operators, but if no such operator is available, the parser will coerce the less-general type to the more-general type and apply the latter's comparison operator.


 As shown above, all comparison operators are binary operators that return values of type `boolean`. Thus, expressions like `1 < 2 < 3` are not valid (because there is no `<` operator to compare a Boolean value with `3`). Use the `BETWEEN` predicates shown below to perform range tests.


 There are also some comparison predicates, as shown in [Comparison Predicates](#functions-comparison-pred-table). These behave much like operators, but have special syntax mandated by the SQL standard.
 <a id="functions-comparison-pred-table"></a>

**Table: Comparison Predicates**

<table>
<thead>
<tr>
<th>Predicate</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>datatype</em> <code>BETWEEN</code> <em>datatype</em> <code>AND</code> <em>datatype</em> <code>boolean</code></td>
<td>Between (inclusive of the range endpoints).</td>
<td><code>2 BETWEEN 1 AND 3</code> <code>t</code><br><code>2 BETWEEN 3 AND 1</code> <code>f</code></td>
</tr>
<tr>
<td><em>datatype</em> <code>NOT BETWEEN</code> <em>datatype</em> <code>AND</code> <em>datatype</em> <code>boolean</code></td>
<td>Not between (the negation of <code>BETWEEN</code>).</td>
<td><code>2 NOT BETWEEN 1 AND 3</code> <code>f</code></td>
</tr>
<tr>
<td><em>datatype</em> <code>BETWEEN SYMMETRIC</code> <em>datatype</em> <code>AND</code> <em>datatype</em> <code>boolean</code></td>
<td>Between, after sorting the two endpoint values.</td>
<td><code>2 BETWEEN SYMMETRIC 3 AND 1</code> <code>t</code></td>
</tr>
<tr>
<td><em>datatype</em> <code>NOT BETWEEN SYMMETRIC</code> <em>datatype</em> <code>AND</code> <em>datatype</em> <code>boolean</code></td>
<td>Not between, after sorting the two endpoint values.</td>
<td><code>2 NOT BETWEEN SYMMETRIC 3 AND 1</code> <code>f</code></td>
</tr>
<tr>
<td><em>datatype</em> <code>IS DISTINCT FROM</code> <em>datatype</em> <code>boolean</code></td>
<td>Not equal, treating null as a comparable value.</td>
<td><code>1 IS DISTINCT FROM NULL</code> <code>t</code> (rather than <code>NULL</code>)<br><code>NULL IS DISTINCT FROM NULL</code> <code>f</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><em>datatype</em> <code>IS NOT DISTINCT FROM</code> <em>datatype</em> <code>boolean</code></td>
<td>Equal, treating null as a comparable value.</td>
<td><code>1 IS NOT DISTINCT FROM NULL</code> <code>f</code> (rather than <code>NULL</code>)<br><code>NULL IS NOT DISTINCT FROM NULL</code> <code>t</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><em>datatype</em> <code>IS NULL</code> <code>boolean</code></td>
<td>Test whether value is null.</td>
<td><code>1.5 IS NULL</code> <code>f</code></td>
</tr>
<tr>
<td><em>datatype</em> <code>IS NOT NULL</code> <code>boolean</code></td>
<td>Test whether value is not null.</td>
<td><code>'null' IS NOT NULL</code> <code>t</code></td>
</tr>
<tr>
<td><em>datatype</em> <code>ISNULL</code> <code>boolean</code></td>
<td>Test whether value is null (nonstandard syntax).</td>
<td></td>
</tr>
<tr>
<td><em>datatype</em> <code>NOTNULL</code> <code>boolean</code></td>
<td>Test whether value is not null (nonstandard syntax).</td>
<td></td>
</tr>
<tr>
<td><code>boolean</code> <code>IS TRUE</code> <code>boolean</code></td>
<td>Test whether boolean expression yields true.</td>
<td><code>true IS TRUE</code> <code>t</code><br><code>NULL::boolean IS TRUE</code> <code>f</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><code>boolean</code> <code>IS NOT TRUE</code> <code>boolean</code></td>
<td>Test whether boolean expression yields false or unknown.</td>
<td><code>true IS NOT TRUE</code> <code>f</code><br><code>NULL::boolean IS NOT TRUE</code> <code>t</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><code>boolean</code> <code>IS FALSE</code> <code>boolean</code></td>
<td>Test whether boolean expression yields false.</td>
<td><code>true IS FALSE</code> <code>f</code><br><code>NULL::boolean IS FALSE</code> <code>f</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><code>boolean</code> <code>IS NOT FALSE</code> <code>boolean</code></td>
<td>Test whether boolean expression yields true or unknown.</td>
<td><code>true IS NOT FALSE</code> <code>t</code><br><code>NULL::boolean IS NOT FALSE</code> <code>t</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><code>boolean</code> <code>IS UNKNOWN</code> <code>boolean</code></td>
<td>Test whether boolean expression yields unknown.</td>
<td><code>true IS UNKNOWN</code> <code>f</code><br><code>NULL::boolean IS UNKNOWN</code> <code>t</code> (rather than <code>NULL</code>)</td>
</tr>
<tr>
<td><code>boolean</code> <code>IS NOT UNKNOWN</code> <code>boolean</code></td>
<td>Test whether boolean expression yields true or false.</td>
<td><code>true IS NOT UNKNOWN</code> <code>t</code><br><code>NULL::boolean IS NOT UNKNOWN</code> <code>f</code> (rather than <code>NULL</code>)</td>
</tr>
</tbody>
</table>


   The `BETWEEN` predicate simplifies range tests:

```

A BETWEEN X AND Y
```
 is equivalent to

```

A >= X AND A <= Y
```
 Notice that `BETWEEN` treats the endpoint values as included in the range. `BETWEEN SYMMETRIC` is like `BETWEEN` except there is no requirement that the argument to the left of `AND` be less than or equal to the argument on the right. If it is not, those two arguments are automatically swapped, so that a nonempty range is always implied.


 The various variants of `BETWEEN` are implemented in terms of the ordinary comparison operators, and therefore will work for any data type(s) that can be compared.


!!! note

    The use of `AND` in the `BETWEEN` syntax creates an ambiguity with the use of `AND` as a logical operator. To resolve this, only a limited set of expression types are allowed as the second argument of a `BETWEEN` clause. If you need to write a more complex sub-expression in `BETWEEN`, write parentheses around the sub-expression.


   Ordinary comparison operators yield null (signifying “unknown”), not true or false, when either input is null. For example, `7 = NULL` yields null, as does `7 <> NULL`. When this behavior is not suitable, use the `IS [ NOT ] DISTINCT FROM` predicates:

```

A IS DISTINCT FROM B
A IS NOT DISTINCT FROM B
```
 For non-null inputs, `IS DISTINCT FROM` is the same as the `<>` operator. However, if both inputs are null it returns false, and if only one input is null it returns true. Similarly, `IS NOT DISTINCT FROM` is identical to `=` for non-null inputs, but it returns true when both inputs are null, and false when only one input is null. Thus, these predicates effectively act as though null were a normal data value, rather than “unknown”.


     To check whether a value is or is not null, use the predicates:

```

EXPRESSION IS NULL
EXPRESSION IS NOT NULL
```
 or the equivalent, but nonstandard, predicates:

```

EXPRESSION ISNULL
EXPRESSION NOTNULL
```


 Do *not* write <em>expression</em><code> = NULL</code> because `NULL` is not “equal to” `NULL`. (The null value represents an unknown value, and it is not known whether two unknown values are equal.)


!!! tip

    Some applications might expect that <em>expression</em><code> = NULL</code> returns true if *expression* evaluates to the null value. It is highly recommended that these applications be modified to comply with the SQL standard. However, if that cannot be done the [transform_null_equals](../../server-administration/server-configuration/version-and-platform-compatibility.md#guc-transform-null-equals) configuration variable is available. If it is enabled, PostgreSQL will convert `x = NULL` clauses to `x IS NULL`.


 If the *expression* is row-valued, then `IS NULL` is true when the row expression itself is null or when all the row's fields are null, while `IS NOT NULL` is true when the row expression itself is non-null and all the row's fields are non-null. Because of this behavior, `IS NULL` and `IS NOT NULL` do not always return inverse results for row-valued expressions; in particular, a row-valued expression that contains both null and non-null fields will return false for both tests. For example:

```sql

SELECT ROW(1,2.5,'this is a test') = ROW(1, 3, 'not the same');

SELECT ROW(table.*) IS NULL FROM table;  -- detect all-null rows

SELECT ROW(table.*) IS NOT NULL FROM table;  -- detect all-non-null rows

SELECT NOT(ROW(table.*) IS NOT NULL) FROM TABLE; -- detect at least one null in rows
```
 In some cases, it may be preferable to write *row* `IS DISTINCT FROM NULL` or *row* `IS NOT DISTINCT FROM NULL`, which will simply check whether the overall row value is null without any additional tests on the row fields.


       Boolean values can also be tested using the predicates

```

BOOLEAN_EXPRESSION IS TRUE
BOOLEAN_EXPRESSION IS NOT TRUE
BOOLEAN_EXPRESSION IS FALSE
BOOLEAN_EXPRESSION IS NOT FALSE
BOOLEAN_EXPRESSION IS UNKNOWN
BOOLEAN_EXPRESSION IS NOT UNKNOWN
```
 These will always return true or false, never a null value, even when the operand is null. A null input is treated as the logical value “unknown”. Notice that `IS UNKNOWN` and `IS NOT UNKNOWN` are effectively the same as `IS NULL` and `IS NOT NULL`, respectively, except that the input expression must be of Boolean type.


 Some comparison-related functions are also available, as shown in [Comparison Functions](#functions-comparison-func-table).
 <a id="functions-comparison-func-table"></a>

**Table: Comparison Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>error_on_null</code> ( <code>anyelement</code> ) <code>anyelement</code></td>
<td>Checks if the input is the null value, generating an error if so; otherwise, returns the input.</td>
<td><code>error_on_null(42)</code> <code>42</code><br><code>error_on_null(row(null,null))</code> <code>(,)</code></td>
</tr>
<tr>
<td><code>num_nonnulls</code> ( <code>VARIADIC</code> <code>"any"</code> ) <code>integer</code></td>
<td>Returns the number of non-null arguments.</td>
<td><code>num_nonnulls(1, NULL, 2)</code> <code>2</code></td>
</tr>
<tr>
<td><code>num_nulls</code> ( <code>VARIADIC</code> <code>"any"</code> ) <code>integer</code></td>
<td>Returns the number of null arguments.</td>
<td><code>num_nulls(1, NULL, 2)</code> <code>1</code></td>
</tr>
</tbody>
</table>
