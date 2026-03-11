<a id="functions-window"></a>

## Window Functions


 *Window functions* provide the ability to perform calculations across sets of rows that are related to the current query row. See [Window Functions](../../tutorial/advanced-features/window-functions.md#tutorial-window) for an introduction to this feature, and [Window Function Calls](../sql-syntax/value-expressions.md#syntax-window-functions) for syntax details.


 The built-in window functions are listed in [General-Purpose Window Functions](#functions-window-table). Note that these functions *must* be invoked using window function syntax, i.e., an `OVER` clause is required.


 In addition to these functions, any built-in or user-defined ordinary aggregate (i.e., not ordered-set or hypothetical-set aggregates) can be used as a window function; see [Aggregate Functions](aggregate-functions.md#functions-aggregate) for a list of the built-in aggregates. Aggregate functions act as window functions only when an `OVER` clause follows the call; otherwise they act as plain aggregates and return a single row for the entire set.
 <a id="functions-window-table"></a>

**Table: General-Purpose Window Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>row_number</code> () <code>bigint</code></td>
<td>Returns the number of the current row within its partition, counting from 1.</td>
<td></td>
</tr>
<tr>
<td><code>rank</code> () <code>bigint</code></td>
<td>Returns the rank of the current row, with gaps; that is, the <code>row_number</code> of the first row in its peer group.</td>
<td></td>
</tr>
<tr>
<td><code>dense_rank</code> () <code>bigint</code></td>
<td>Returns the rank of the current row, without gaps; this function effectively counts peer groups.</td>
<td></td>
</tr>
<tr>
<td><code>percent_rank</code> () <code>double precision</code></td>
<td>Returns the relative rank of the current row, that is (<code>rank</code> - 1) / (total partition rows - 1). The value thus ranges from 0 to 1 inclusive.</td>
<td></td>
</tr>
<tr>
<td><code>cume_dist</code> () <code>double precision</code></td>
<td>Returns the cumulative distribution, that is (number of partition rows preceding or peers with current row) / (total partition rows). The value thus ranges from 1/<code>N</code> to 1.</td>
<td></td>
</tr>
<tr>
<td><code>ntile</code> ( <code>num_buckets</code> <code>integer</code> ) <code>integer</code></td>
<td>Returns an integer ranging from 1 to the argument value, dividing the partition as equally as possible.</td>
<td></td>
</tr>
<tr>
<td><code>lag</code> ( <code>value</code> <code>anycompatible</code> [, <code>offset</code> <code>integer</code> [, <code>default</code> <code>anycompatible</code> ]] ) <code>anycompatible</code></td>
<td>Returns <code>value</code> evaluated at the row that is <code>offset</code> rows before the current row within the partition; if there is no such row, instead returns <code>default</code> (which must be of a type compatible with <code>value</code>). Both <code>offset</code> and <code>default</code> are evaluated with respect to the current row. If omitted, <code>offset</code> defaults to 1 and <code>default</code> to <code>NULL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>lead</code> ( <code>value</code> <code>anycompatible</code> [, <code>offset</code> <code>integer</code> [, <code>default</code> <code>anycompatible</code> ]] ) <code>anycompatible</code></td>
<td>Returns <code>value</code> evaluated at the row that is <code>offset</code> rows after the current row within the partition; if there is no such row, instead returns <code>default</code> (which must be of a type compatible with <code>value</code>). Both <code>offset</code> and <code>default</code> are evaluated with respect to the current row. If omitted, <code>offset</code> defaults to 1 and <code>default</code> to <code>NULL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>first_value</code> ( <code>value</code> <code>anyelement</code> ) <code>anyelement</code></td>
<td>Returns <code>value</code> evaluated at the row that is the first row of the window frame.</td>
<td></td>
</tr>
<tr>
<td><code>last_value</code> ( <code>value</code> <code>anyelement</code> ) <code>anyelement</code></td>
<td>Returns <code>value</code> evaluated at the row that is the last row of the window frame.</td>
<td></td>
</tr>
<tr>
<td><code>nth_value</code> ( <code>value</code> <code>anyelement</code>, <code>n</code> <code>integer</code> ) <code>anyelement</code></td>
<td>Returns <code>value</code> evaluated at the row that is the <code>n</code>'th row of the window frame (counting from 1); returns <code>NULL</code> if there is no such row.</td>
<td></td>
</tr>
</tbody>
</table>


 All of the functions listed in [General-Purpose Window Functions](#functions-window-table) depend on the sort ordering specified by the `ORDER BY` clause of the associated window definition. Rows that are not distinct when considering only the `ORDER BY` columns are said to be *peers*. The four ranking functions (including `cume_dist`) are defined so that they give the same answer for all rows of a peer group.


 Note that `first_value`, `last_value`, and `nth_value` consider only the rows within the “window frame”, which by default contains the rows from the start of the partition through the last peer of the current row. This is likely to give unhelpful results for `last_value` and sometimes also `nth_value`. You can redefine the frame by adding a suitable frame specification (`RANGE`, `ROWS` or `GROUPS`) to the `OVER` clause. See [Window Function Calls](../sql-syntax/value-expressions.md#syntax-window-functions) for more information about frame specifications.


 When an aggregate function is used as a window function, it aggregates over the rows within the current row's window frame. An aggregate used with `ORDER BY` and the default window frame definition produces a “running sum” type of behavior, which may or may not be what's wanted. To obtain aggregation over the whole partition, omit `ORDER BY` or use `ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING`. Other frame specifications can be used to obtain other effects.


!!! note

    The SQL standard defines a `RESPECT NULLS` or `IGNORE NULLS` option for `lead`, `lag`, `first_value`, `last_value`, and `nth_value`. This is not implemented in PostgreSQL: the behavior is always the same as the standard's default, namely `RESPECT NULLS`. Likewise, the standard's `FROM FIRST` or `FROM LAST` option for `nth_value` is not implemented: only the default `FROM FIRST` behavior is supported. (You can achieve the result of `FROM LAST` by reversing the `ORDER BY` ordering.)
