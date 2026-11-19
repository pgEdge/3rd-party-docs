<a id="catalog-pg-aggregate"></a>

## `pg_aggregate`


 The catalog `pg_aggregate` stores information about aggregate functions. An aggregate function is a function that operates on a set of values (typically one column from each row that matches a query condition) and returns a single value computed from all these values. Typical aggregate functions are `sum`, `count`, and `max`. Each entry in `pg_aggregate` is an extension of an entry in [`pg_proc`](pg_proc.md#catalog-pg-proc). The `pg_proc` entry carries the aggregate's name, input and output data types, and other information that is similar to ordinary functions.


**Table: `pg_aggregate` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>aggfnoid</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p><code>pg_proc</code> OID of the aggregate function</p></td>
</tr>
<tr>
<td><p><code>aggkind</code> <code>char</code></p>
<p>Aggregate kind: <code>n</code> for “normal” aggregates, <code>o</code> for “ordered-set” aggregates, or <code>h</code> for “hypothetical-set” aggregates</p></td>
</tr>
<tr>
<td><p><code>aggnumdirectargs</code> <code>int2</code></p>
<p>Number of direct (non-aggregated) arguments of an ordered-set or hypothetical-set aggregate, counting a variadic array as one argument. If equal to <code>pronargs</code>, the aggregate must be variadic and the variadic array describes the aggregated arguments as well as the final direct arguments. Always zero for normal aggregates.</p></td>
</tr>
<tr>
<td><p><code>aggtransfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Transition function</p></td>
</tr>
<tr>
<td><p><code>aggfinalfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Final function (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggcombinefn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Combine function (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggserialfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Serialization function (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggdeserialfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Deserialization function (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggmtransfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Forward transition function for moving-aggregate mode (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggminvtransfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Inverse transition function for moving-aggregate mode (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggmfinalfn</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Final function for moving-aggregate mode (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggfinalextra</code> <code>bool</code></p>
<p>True to pass extra dummy arguments to <code>aggfinalfn</code></p></td>
</tr>
<tr>
<td><p><code>aggmfinalextra</code> <code>bool</code></p>
<p>True to pass extra dummy arguments to <code>aggmfinalfn</code></p></td>
</tr>
<tr>
<td><p><code>aggfinalmodify</code> <code>char</code></p>
<p>Whether <code>aggfinalfn</code> modifies the transition state value: <code>r</code> if it is read-only, <code>s</code> if the <code>aggtransfn</code> cannot be applied after the <code>aggfinalfn</code>, or <code>w</code> if it writes on the value</p></td>
</tr>
<tr>
<td><p><code>aggmfinalmodify</code> <code>char</code></p>
<p>Like <code>aggfinalmodify</code>, but for the <code>aggmfinalfn</code></p></td>
</tr>
<tr>
<td><p><code>aggsortop</code> <code>oid</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>Associated sort operator (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggtranstype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type of the aggregate function's internal transition (state) data</p></td>
</tr>
<tr>
<td><p><code>aggtransspace</code> <code>int4</code></p>
<p>Approximate average size (in bytes) of the transition state data. A positive value provides an estimate; zero means to use a default estimate. A negative value indicates the state data can grow unboundedly in size, such as when the aggregate accumulates input rows (e.g., array_agg, string_agg).</p></td>
</tr>
<tr>
<td><p><code>aggmtranstype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type of the aggregate function's internal transition (state) data for moving-aggregate mode (zero if none)</p></td>
</tr>
<tr>
<td><p><code>aggmtransspace</code> <code>int4</code></p>
<p>Approximate average size (in bytes) of the transition state data for moving-aggregate mode, or zero to use a default estimate</p></td>
</tr>
<tr>
<td><p><code>agginitval</code> <code>text</code></p>
<p>The initial value of the transition state. This is a text field containing the initial value in its external string representation. If this field is null, the transition state value starts out null.</p></td>
</tr>
<tr>
<td><p><code>aggminitval</code> <code>text</code></p>
<p>The initial value of the transition state for moving-aggregate mode. This is a text field containing the initial value in its external string representation. If this field is null, the transition state value starts out null.</p></td>
</tr>
</tbody>
</table>


 New aggregate functions are registered with the [`CREATE AGGREGATE`](../../reference/sql-commands/create-aggregate.md#sql-createaggregate) command. See [User-Defined Aggregates](../../server-programming/extending-sql/user-defined-aggregates.md#xaggr) for more information about writing aggregate functions and the meaning of the transition functions, etc.
