<a id="catalog-pg-partitioned-table"></a>

## `pg_partitioned_table`


 The catalog `pg_partitioned_table` stores information about how tables are partitioned.


**Table: `pg_partitioned_table` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>partrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry for this partitioned table</p></td>
</tr>
<tr>
<td><p><code>partstrat</code> <code>char</code></p>
<p>Partitioning strategy; <code>h</code> = hash partitioned table, <code>l</code> = list partitioned table, <code>r</code> = range partitioned table</p></td>
</tr>
<tr>
<td><p><code>partnatts</code> <code>int2</code></p>
<p>The number of columns in the partition key</p></td>
</tr>
<tr>
<td><p><code>partdefid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry for the default partition of this partitioned table, or zero if this partitioned table does not have a default partition</p></td>
</tr>
<tr>
<td><p><code>partattrs</code> <code>int2vector</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>This is an array of <code>partnatts</code> values that indicate which table columns are part of the partition key. For example, a value of <code>1 3</code> would mean that the first and the third table columns make up the partition key. A zero in this array indicates that the corresponding partition key column is an expression, rather than a simple column reference.</p></td>
</tr>
<tr>
<td><p><code>partclass</code> <code>oidvector</code> (references <a href="pg_opclass.md#catalog-pg-opclass"><code>pg_opclass</code></a>.<code>oid</code>)</p>
<p>For each column in the partition key, this contains the OID of the operator class to use. See <a href="pg_opclass.md#catalog-pg-opclass"><code>pg_opclass</code></a> for details.</p></td>
</tr>
<tr>
<td><p><code>partcollation</code> <code>oidvector</code> (references <a href="pg_collation.md#catalog-pg-collation"><code>pg_collation</code></a>.<code>oid</code>)</p>
<p>For each column in the partition key, this contains the OID of the collation to use for partitioning, or zero if the column is not of a collatable data type.</p></td>
</tr>
<tr>
<td><p><code>partexprs</code> <code>pg_node_tree</code></p>
<p>Expression trees (in <code>nodeToString()</code> representation) for partition key columns that are not simple column references. This is a list with one element for each zero entry in <code>partattrs</code>. Null if all partition key columns are simple references.</p></td>
</tr>
</tbody>
</table>
