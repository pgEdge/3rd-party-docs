<a id="catalog-pg-statistic-ext"></a>

## `pg_statistic_ext`


 The catalog `pg_statistic_ext` holds definitions of extended planner statistics. Each row in this catalog corresponds to a *statistics object* created with [`CREATE STATISTICS`](../../reference/sql-commands/create-statistics.md#sql-createstatistics).


**Table: `pg_statistic_ext` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>oid</code> <code>oid</code></p>
<p>Row identifier</p></td>
</tr>
<tr>
<td><p><code>stxrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Table containing the columns described by this object</p></td>
</tr>
<tr>
<td><p><code>stxname</code> <code>name</code></p>
<p>Name of the statistics object</p></td>
</tr>
<tr>
<td><p><code>stxnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this statistics object</p></td>
</tr>
<tr>
<td><p><code>stxowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the statistics object</p></td>
</tr>
<tr>
<td><p><code>stxkeys</code> <code>int2vector</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>An array of attribute numbers, indicating which table columns are covered by this statistics object; for example a value of <code>1 3</code> would mean that the first and the third table columns are covered</p></td>
</tr>
<tr>
<td><p><code>stxstattarget</code> <code>int2</code></p>
<p><code>stxstattarget</code> controls the level of detail of statistics accumulated for this statistics object by <a href="../../reference/sql-commands/analyze.md#sql-analyze"><code>ANALYZE</code></a>. A zero value indicates that no statistics should be collected. A null value says to use the maximum of the statistics targets of the referenced columns, if set, or the system default statistics target. Positive values of <code>stxstattarget</code> determine the target number of “most common values” to collect.</p></td>
</tr>
<tr>
<td><p><code>stxkind</code> <code>char[]</code></p>
<p>An array containing codes for the enabled statistics kinds; valid values are: <code>d</code> for n-distinct statistics, <code>f</code> for functional dependency statistics, <code>m</code> for most common values (MCV) list statistics, and <code>e</code> for expression statistics</p></td>
</tr>
<tr>
<td><p><code>stxexprs</code> <code>pg_node_tree</code></p>
<p>Expression trees (in <code>nodeToString()</code> representation) for statistics object attributes that are not simple column references. This is a list with one element per expression. Null if all statistics object attributes are simple references.</p></td>
</tr>
</tbody>
</table>


 The `pg_statistic_ext` entry is filled in completely during [`CREATE STATISTICS`](../../reference/sql-commands/create-statistics.md#sql-createstatistics), but the actual statistical values are not computed then. Subsequent [`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze) commands compute the desired values and populate an entry in the [`pg_statistic_ext_data`](pg_statistic_ext_data.md#catalog-pg-statistic-ext-data) catalog.
