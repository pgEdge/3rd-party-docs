<a id="functions-aggregate"></a>

## Aggregate Functions


 *Aggregate functions* compute a single result from a set of input values. The built-in general-purpose aggregate functions are listed in [General-Purpose Aggregate Functions](#functions-aggregate-table) while statistical aggregates are in [Aggregate Functions for Statistics](#functions-aggregate-statistics-table). The built-in within-group ordered-set aggregate functions are listed in [Ordered-Set Aggregate Functions](#functions-orderedset-table) while the built-in within-group hypothetical-set ones are in [Hypothetical-Set Aggregate Functions](#functions-hypothetical-table). Grouping operations, which are closely related to aggregate functions, are listed in [Grouping Operations](#functions-grouping-table). The special syntax considerations for aggregate functions are explained in [Aggregate Expressions](../sql-syntax/value-expressions.md#syntax-aggregates). Consult [Aggregate Functions](../../tutorial/the-sql-language/aggregate-functions.md#tutorial-agg) for additional introductory information.


 Aggregate functions that support *Partial Mode* are eligible to participate in various optimizations, such as parallel aggregation.


 While all aggregates below accept an optional `ORDER BY` clause (as outlined in [Aggregate Expressions](../sql-syntax/value-expressions.md#syntax-aggregates)), the clause has only been added to aggregates whose output is affected by ordering.
 <a id="functions-aggregate-table"></a>

**Table: General-Purpose Aggregate Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>any_value</code> ( <code>anyelement</code> ) <em>same as input type</em></td>
<td>Returns an arbitrary value from the non-null input values.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>array_agg</code> ( <code>anynonarray</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>anyarray</code></td>
<td>Collects all the input values, including nulls, into an array.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>array_agg</code> ( <code>anyarray</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>anyarray</code></td>
<td>Concatenates all the input arrays into an array of one higher dimension. (The inputs must all have the same dimensionality, and cannot be empty or null.)</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>avg</code> ( <code>smallint</code> ) <code>numeric</code></td>
<td><code>avg</code> ( <code>integer</code> ) <code>numeric</code></td>
<td><code>avg</code> ( <code>bigint</code> ) <code>numeric</code><br><code>avg</code> ( <code>numeric</code> ) <code>numeric</code><br><code>avg</code> ( <code>real</code> ) <code>double precision</code><br><code>avg</code> ( <code>double precision</code> ) <code>double precision</code><br><code>avg</code> ( <code>interval</code> ) <code>interval</code><br>Computes the average (arithmetic mean) of all the non-null input values.</td>
<td></td>
</tr>
<tr>
<td><code>bit_and</code> ( <code>smallint</code> ) <code>smallint</code></td>
<td><code>bit_and</code> ( <code>integer</code> ) <code>integer</code></td>
<td><code>bit_and</code> ( <code>bigint</code> ) <code>bigint</code><br><code>bit_and</code> ( <code>bit</code> ) <code>bit</code><br>Computes the bitwise AND of all non-null input values.</td>
<td></td>
</tr>
<tr>
<td><code>bit_or</code> ( <code>smallint</code> ) <code>smallint</code></td>
<td><code>bit_or</code> ( <code>integer</code> ) <code>integer</code></td>
<td><code>bit_or</code> ( <code>bigint</code> ) <code>bigint</code><br><code>bit_or</code> ( <code>bit</code> ) <code>bit</code><br>Computes the bitwise OR of all non-null input values.</td>
<td></td>
</tr>
<tr>
<td><code>bit_xor</code> ( <code>smallint</code> ) <code>smallint</code></td>
<td><code>bit_xor</code> ( <code>integer</code> ) <code>integer</code></td>
<td><code>bit_xor</code> ( <code>bigint</code> ) <code>bigint</code><br><code>bit_xor</code> ( <code>bit</code> ) <code>bit</code><br>Computes the bitwise exclusive OR of all non-null input values. Can be useful as a checksum for an unordered set of values.</td>
<td></td>
</tr>
<tr>
<td><code>bool_and</code> ( <code>boolean</code> ) <code>boolean</code></td>
<td>Returns true if all non-null input values are true, otherwise false.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>bool_or</code> ( <code>boolean</code> ) <code>boolean</code></td>
<td>Returns true if any non-null input value is true, otherwise false.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>count</code> ( <code>*</code> ) <code>bigint</code></td>
<td>Computes the number of input rows.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>count</code> ( <code>"any"</code> ) <code>bigint</code></td>
<td>Computes the number of input rows in which the input value is not null.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>every</code> ( <code>boolean</code> ) <code>boolean</code></td>
<td>This is the SQL standard's equivalent to <code>bool_and</code>.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>json_agg</code> ( <code>anyelement</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>json</code></td>
<td><code>jsonb_agg</code> ( <code>anyelement</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>jsonb</code></td>
<td>Collects all the input values, including nulls, into a JSON array. Values are converted to JSON as per <code>to_json</code> or <code>to_jsonb</code>.</td>
<td></td>
</tr>
<tr>
<td><code>json_agg_strict</code> ( <code>anyelement</code> ) <code>json</code></td>
<td><code>jsonb_agg_strict</code> ( <code>anyelement</code> ) <code>jsonb</code></td>
<td>Collects all the input values, skipping nulls, into a JSON array. Values are converted to JSON as per <code>to_json</code> or <code>to_jsonb</code>.</td>
<td></td>
</tr>
<tr>
<td><code>json_arrayagg</code> ( [ <em>value_expression</em> ] [ <code>ORDER BY</code> <em>sort_expression</em> ] [ { <code>NULL</code> | <code>ABSENT</code> } <code>ON NULL</code> ] [ <code>RETURNING</code> <em>data_type</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] ])</td>
<td>Behaves in the same way as <code>json_array</code> but as an aggregate function so it only takes one <em>value_expression</em> parameter. If <code>ABSENT ON NULL</code> is specified, any NULL values are omitted. If <code>ORDER BY</code> is specified, the elements will appear in the array in that order rather than in the input order.</td>
<td><code>SELECT json_arrayagg(v) FROM (VALUES(2),(1)) t(v)</code> <code>[2, 1]</code></td>
<td></td>
</tr>
<tr>
<td><code>json_objectagg</code> ( [ { <em>key_expression</em> { <code>VALUE</code> | ':' } <em>value_expression</em> } ] [ { <code>NULL</code> | <code>ABSENT</code> } <code>ON NULL</code> ] [ { <code>WITH</code> | <code>WITHOUT</code> } <code>UNIQUE</code> [ <code>KEYS</code> ] ] [ <code>RETURNING</code> <em>data_type</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] ])</td>
<td>Behaves like <code>json_object</code>, but as an aggregate function, so it only takes one <em>key_expression</em> and one <em>value_expression</em> parameter.</td>
<td><code>SELECT json_objectagg(k:v) FROM (VALUES ('a'::text,current_date),('b',current_date + 1)) AS t(k,v)</code> <code>{ "a" : "2022-05-10", "b" : "2022-05-11" }</code></td>
<td></td>
</tr>
<tr>
<td><code>json_object_agg</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>json</code></td>
<td><code>jsonb_object_agg</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>jsonb</code></td>
<td>Collects all the key/value pairs into a JSON object. Key arguments are coerced to text; value arguments are converted as per <code>to_json</code> or <code>to_jsonb</code>. Values can be null, but keys cannot.</td>
<td></td>
</tr>
<tr>
<td><code>json_object_agg_strict</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> ) <code>json</code></td>
<td><code>jsonb_object_agg_strict</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> ) <code>jsonb</code></td>
<td>Collects all the key/value pairs into a JSON object. Key arguments are coerced to text; value arguments are converted as per <code>to_json</code> or <code>to_jsonb</code>. The <code>key</code> cannot be null. If the <code>value</code> is null then the entry is skipped,</td>
<td></td>
</tr>
<tr>
<td><code>json_object_agg_unique</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> ) <code>json</code></td>
<td><code>jsonb_object_agg_unique</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> ) <code>jsonb</code></td>
<td>Collects all the key/value pairs into a JSON object. Key arguments are coerced to text; value arguments are converted as per <code>to_json</code> or <code>to_jsonb</code>. Values can be null, but keys cannot. If there is a duplicate key an error is thrown.</td>
<td></td>
</tr>
<tr>
<td><code>json_object_agg_unique_strict</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> ) <code>json</code></td>
<td><code>jsonb_object_agg_unique_strict</code> ( <code>key</code> <code>"any"</code>, <code>value</code> <code>"any"</code> ) <code>jsonb</code></td>
<td>Collects all the key/value pairs into a JSON object. Key arguments are coerced to text; value arguments are converted as per <code>to_json</code> or <code>to_jsonb</code>. The <code>key</code> cannot be null. If the <code>value</code> is null then the entry is skipped. If there is a duplicate key an error is thrown.</td>
<td></td>
</tr>
<tr>
<td><code>max</code> ( <em>see text</em> ) <em>same as input type</em></td>
<td>Computes the maximum of the non-null input values. Available for any numeric, string, date/time, or enum type, as well as <code>bytea</code>, <code>inet</code>, <code>interval</code>, <code>money</code>, <code>oid</code>, <code>oid8</code>, <code>pg_lsn</code>, <code>tid</code>, <code>xid8</code>, and also arrays and composite types containing sortable data types.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>min</code> ( <em>see text</em> ) <em>same as input type</em></td>
<td>Computes the minimum of the non-null input values. Available for any numeric, string, date/time, or enum type, as well as <code>bytea</code>, <code>inet</code>, <code>interval</code>, <code>money</code>, <code>oid</code>, <code>oid8</code>, <code>pg_lsn</code>, <code>tid</code>, <code>xid8</code>, and also arrays and composite types containing sortable data types.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>range_agg</code> ( <code>value</code> <code>anyrange</code> ) <code>anymultirange</code></td>
<td><code>range_agg</code> ( <code>value</code> <code>anymultirange</code> ) <code>anymultirange</code></td>
<td>Computes the union of the non-null input values.</td>
<td></td>
</tr>
<tr>
<td><code>range_intersect_agg</code> ( <code>value</code> <code>anyrange</code> ) <code>anyrange</code></td>
<td><code>range_intersect_agg</code> ( <code>value</code> <code>anymultirange</code> ) <code>anymultirange</code></td>
<td>Computes the intersection of the non-null input values.</td>
<td></td>
</tr>
<tr>
<td><code>string_agg</code> ( <code>value</code> <code>text</code>, <code>delimiter</code> <code>text</code> ) <code>text</code></td>
<td><code>string_agg</code> ( <code>value</code> <code>bytea</code>, <code>delimiter</code> <code>bytea</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>bytea</code></td>
<td>Concatenates the non-null input values into a string. Each value after the first is preceded by the corresponding <code>delimiter</code> (if it's not null).</td>
<td></td>
</tr>
<tr>
<td><code>sum</code> ( <code>smallint</code> ) <code>bigint</code></td>
<td><code>sum</code> ( <code>integer</code> ) <code>bigint</code></td>
<td><code>sum</code> ( <code>bigint</code> ) <code>numeric</code><br><code>sum</code> ( <code>numeric</code> ) <code>numeric</code><br><code>sum</code> ( <code>real</code> ) <code>real</code><br><code>sum</code> ( <code>double precision</code> ) <code>double precision</code><br><code>sum</code> ( <code>interval</code> ) <code>interval</code><br><code>sum</code> ( <code>money</code> ) <code>money</code><br>Computes the sum of the non-null input values.</td>
<td></td>
</tr>
<tr>
<td><code>xmlagg</code> ( <code>xml</code> <code>ORDER BY</code> <code>input_sort_columns</code> ) <code>xml</code></td>
<td>Concatenates the non-null XML input values (see <a href="xml-functions.md#functions-xml-xmlagg"><code>xmlagg</code></a>).</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>


 It should be noted that except for `count`, these functions return a null value when no rows are selected. In particular, `sum` of no rows returns null, not zero as one might expect, and `array_agg` returns null rather than an empty array when there are no input rows. The `coalesce` function can be used to substitute zero or an empty array for null when necessary.


 The aggregate functions `array_agg`, `json_agg`, `jsonb_agg`, `json_agg_strict`, `jsonb_agg_strict`, `json_object_agg`, `jsonb_object_agg`, `json_object_agg_strict`, `jsonb_object_agg_strict`, `json_object_agg_unique`, `jsonb_object_agg_unique`, `json_object_agg_unique_strict`, `jsonb_object_agg_unique_strict`, `string_agg`, and `xmlagg`, as well as similar user-defined aggregate functions, produce meaningfully different result values depending on the order of the input values. This ordering is unspecified by default, but can be controlled by writing an `ORDER BY` clause within the aggregate call, as shown in [Aggregate Expressions](../sql-syntax/value-expressions.md#syntax-aggregates). Alternatively, supplying the input values from a sorted subquery will usually work. For example:

```

SELECT xmlagg(x) FROM (SELECT x FROM test ORDER BY y DESC) AS tab;
```
 Beware that this approach can fail if the outer query level contains additional processing, such as a join, because that might cause the subquery's output to be reordered before the aggregate is computed.


!!! note

    The boolean aggregates `bool_and` and `bool_or` correspond to the standard SQL aggregates `every` and `any` or `some`. PostgreSQL supports `every`, but not `any` or `some`, because there is an ambiguity built into the standard syntax:

    ```sql

    SELECT b1 = ANY((SELECT b2 FROM t2 ...)) FROM t1 ...;
    ```
     Here `ANY` can be considered either as introducing a subquery, or as being an aggregate function, if the subquery returns one row with a Boolean value. Thus the standard name cannot be given to these aggregates.


!!! note

    Users accustomed to working with other SQL database management systems might be disappointed by the performance of the `count` aggregate when it is applied to the entire table. A query like:

    ```sql

    SELECT count(*) FROM sometable;
    ```
     will require effort proportional to the size of the table: PostgreSQL will need to scan either the entire table or the entirety of an index that includes all rows in the table.


 [Aggregate Functions for Statistics](#functions-aggregate-statistics-table) shows aggregate functions typically used in statistical analysis. (These are separated out merely to avoid cluttering the listing of more-commonly-used aggregates.) Functions shown as accepting *numeric_type* are available for all the types `smallint`, `integer`, `bigint`, `numeric`, `real`, and `double precision`. Where the description mentions `N`, it means the number of input rows for which all the input expressions are non-null. In all cases, null is returned if the computation is meaningless, for example when `N` is zero.
   <a id="functions-aggregate-statistics-table"></a>

**Table: Aggregate Functions for Statistics**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>corr</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the correlation coefficient.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>covar_pop</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the population covariance.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>covar_samp</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the sample covariance.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_avgx</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the average of the independent variable, <code>sum(</code>X<code>)/</code>N``.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_avgy</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the average of the dependent variable, <code>sum(</code>Y<code>)/</code>N``.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_count</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>bigint</code></td>
<td>Computes the number of rows in which both inputs are non-null.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_intercept</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the y-intercept of the least-squares-fit linear equation determined by the (<code>X</code>, <code>Y</code>) pairs.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_r2</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the square of the correlation coefficient.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_slope</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the slope of the least-squares-fit linear equation determined by the (<code>X</code>, <code>Y</code>) pairs.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_sxx</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the “sum of squares” of the independent variable, <code>sum(</code>X<code>^2) - sum(</code>X<code>)^2/</code>N``.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_sxy</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the “sum of products” of independent times dependent variables, <code>sum(</code>X<code><em></code>Y<code>) - sum(</code>X<code>) </em> sum(</code>Y<code>)/</code>N``.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>regr_syy</code> ( <code>Y</code> <code>double precision</code>, <code>X</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Computes the “sum of squares” of the dependent variable, <code>sum(</code>Y<code>^2) - sum(</code>Y<code>)^2/</code>N``.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>stddev</code> ( <em>numeric_type</em> )  <code>double precision</code> for <code>real</code> or <code>double precision</code>, otherwise <code>numeric</code></td>
<td>This is a historical alias for <code>stddev_samp</code>.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>stddev_pop</code> ( <em>numeric_type</em> )  <code>double precision</code> for <code>real</code> or <code>double precision</code>, otherwise <code>numeric</code></td>
<td>Computes the population standard deviation of the input values.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>stddev_samp</code> ( <em>numeric_type</em> )  <code>double precision</code> for <code>real</code> or <code>double precision</code>, otherwise <code>numeric</code></td>
<td>Computes the sample standard deviation of the input values.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>variance</code> ( <em>numeric_type</em> )  <code>double precision</code> for <code>real</code> or <code>double precision</code>, otherwise <code>numeric</code></td>
<td>This is a historical alias for <code>var_samp</code>.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>var_pop</code> ( <em>numeric_type</em> )  <code>double precision</code> for <code>real</code> or <code>double precision</code>, otherwise <code>numeric</code></td>
<td>Computes the population variance of the input values (square of the population standard deviation).</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>var_samp</code> ( <em>numeric_type</em> )  <code>double precision</code> for <code>real</code> or <code>double precision</code>, otherwise <code>numeric</code></td>
<td>Computes the sample variance of the input values (square of the sample standard deviation).</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>


 [Ordered-Set Aggregate Functions](#functions-orderedset-table) shows some aggregate functions that use the *ordered-set aggregate* syntax. These functions are sometimes referred to as “inverse distribution” functions. Their aggregated input is introduced by `ORDER BY`, and they may also take a *direct argument* that is not aggregated, but is computed only once. All these functions ignore null values in their aggregated input. For those that take a `fraction` parameter, the fraction value must be between 0 and 1; an error is thrown if not. However, a null `fraction` value simply produces a null result.
   <a id="functions-orderedset-table"></a>

**Table: Ordered-Set Aggregate Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>mode</code> () <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>anyelement</code> ) <code>anyelement</code></td>
<td>Computes the <em>mode</em>, the most frequent value of the aggregated argument (arbitrarily choosing the first one if there are multiple equally-frequent values). The aggregated argument must be of a sortable type.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>percentile_cont</code> ( <code>fraction</code> <code>double precision</code> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>double precision</code> ) <code>double precision</code></td>
<td><code>percentile_cont</code> ( <code>fraction</code> <code>double precision</code> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>interval</code> ) <code>interval</code></td>
<td>Computes the <em>continuous percentile</em>, a value corresponding to the specified <code>fraction</code> within the ordered set of aggregated argument values. This will interpolate between adjacent input items if needed.</td>
<td></td>
</tr>
<tr>
<td><code>percentile_cont</code> ( <code>fractions</code> <code>double precision[]</code> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>double precision</code> ) <code>double precision[]</code></td>
<td><code>percentile_cont</code> ( <code>fractions</code> <code>double precision[]</code> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>interval</code> ) <code>interval[]</code></td>
<td>Computes multiple continuous percentiles. The result is an array of the same dimensions as the <code>fractions</code> parameter, with each non-null element replaced by the (possibly interpolated) value corresponding to that percentile.</td>
<td></td>
</tr>
<tr>
<td><code>percentile_disc</code> ( <code>fraction</code> <code>double precision</code> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>anyelement</code> ) <code>anyelement</code></td>
<td>Computes the <em>discrete percentile</em>, the first value within the ordered set of aggregated argument values whose position in the ordering equals or exceeds the specified <code>fraction</code>. The aggregated argument must be of a sortable type.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>percentile_disc</code> ( <code>fractions</code> <code>double precision[]</code> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <code>anyelement</code> ) <code>anyarray</code></td>
<td>Computes multiple discrete percentiles. The result is an array of the same dimensions as the <code>fractions</code> parameter, with each non-null element replaced by the input value corresponding to that percentile. The aggregated argument must be of a sortable type.</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>


 Each of the “hypothetical-set” aggregates listed in [Hypothetical-Set Aggregate Functions](#functions-hypothetical-table) is associated with a window function of the same name defined in [Window Functions](window-functions.md#functions-window). In each case, the aggregate's result is the value that the associated window function would have returned for the “hypothetical” row constructed from *args*, if such a row had been added to the sorted group of rows represented by the *sorted_args*. For each of these functions, the list of direct arguments given in *args* must match the number and types of the aggregated arguments given in *sorted_args*. Unlike most built-in aggregates, these aggregates are not strict, that is they do not drop input rows containing nulls. Null values sort according to the rule specified in the `ORDER BY` clause.
 <a id="functions-hypothetical-table"></a>

**Table: Hypothetical-Set Aggregate Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>rank</code> ( <em>args</em> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <em>sorted_args</em> ) <code>bigint</code></td>
<td>Computes the rank of the hypothetical row, with gaps; that is, the row number of the first row in its peer group.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>dense_rank</code> ( <em>args</em> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <em>sorted_args</em> ) <code>bigint</code></td>
<td>Computes the rank of the hypothetical row, without gaps; this function effectively counts peer groups.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>percent_rank</code> ( <em>args</em> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <em>sorted_args</em> ) <code>double precision</code></td>
<td>Computes the relative rank of the hypothetical row, that is (<code>rank</code> - 1) / (total rows - 1). The value thus ranges from 0 to 1 inclusive.</td>
<td></td>
<td></td>
</tr>
<tr>
<td><code>cume_dist</code> ( <em>args</em> ) <code>WITHIN GROUP</code> ( <code>ORDER BY</code> <em>sorted_args</em> ) <code>double precision</code></td>
<td>Computes the cumulative distribution, that is (number of rows preceding or peers with hypothetical row) / (total rows). The value thus ranges from 1/<code>N</code> to 1.</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
 <a id="functions-grouping-table"></a>

**Table: Grouping Operations**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>GROUPING</code> ( <em>group_by_expression(s)</em> ) <code>integer</code></td>
<td>Returns a bit mask indicating which <code>GROUP BY</code> expressions are not included in the current grouping set. Bits are assigned with the rightmost argument corresponding to the least-significant bit; each bit is 0 if the corresponding expression is included in the grouping criteria of the grouping set generating the current result row, and 1 if it is not included.</td>
<td></td>
</tr>
</tbody>
</table>


 The grouping operations shown in [Grouping Operations](#functions-grouping-table) are used in conjunction with grouping sets (see [`GROUPING SETS`, `CUBE`, and `ROLLUP`](../queries/table-expressions.md#queries-grouping-sets)) to distinguish result rows. The arguments to the `GROUPING` function are not actually evaluated, but they must exactly match expressions given in the `GROUP BY` clause of the associated query level. For example:

```

=> SELECT * FROM items_sold;
 make  | model | sales
-------+-------+-------
 Foo   | GT    |  10
 Foo   | Tour  |  20
 Bar   | City  |  15
 Bar   | Sport |  5
(4 rows)

=> SELECT make, model, GROUPING(make,model), sum(sales) FROM items_sold GROUP BY ROLLUP(make,model);
 make  | model | grouping | sum
-------+-------+----------+-----
 Foo   | GT    |        0 | 10
 Foo   | Tour  |        0 | 20
 Bar   | City  |        0 | 15
 Bar   | Sport |        0 | 5
 Foo   |       |        1 | 30
 Bar   |       |        1 | 20
       |       |        3 | 50
(7 rows)
```
 Here, the `grouping` value `0` in the first four rows shows that those have been grouped normally, over both the grouping columns. The value `1` indicates that `model` was not grouped by in the next-to-last two rows, and the value `3` indicates that neither `make` nor `model` was grouped by in the last row (which therefore is an aggregate over all the input rows).
