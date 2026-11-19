<a id="catalog-pg-index"></a>

## `pg_index`


 The catalog `pg_index` contains part of the information about indexes. The rest is mostly in [`pg_class`](pg_class.md#catalog-pg-class).


**Table: `pg_index` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>indexrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry for this index</p></td>
</tr>
<tr>
<td><p><code>indrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry for the table this index is for</p></td>
</tr>
<tr>
<td><p><code>indnatts</code> <code>int2</code></p>
<p>The total number of columns in the index (duplicates <code>pg_class.relnatts</code>); this number includes both key and included attributes</p></td>
</tr>
<tr>
<td><p><code>indnkeyatts</code> <code>int2</code></p>
<p>The number of <em>key columns</em> in the index, not counting any <em>included columns</em>, which are merely stored and do not participate in the index semantics</p></td>
</tr>
<tr>
<td><p><code>indisunique</code> <code>bool</code></p>
<p>If true, this is a unique index</p></td>
</tr>
<tr>
<td><p><code>indnullsnotdistinct</code> <code>bool</code></p>
<p>This value is only used for unique indexes. If false, this unique index will consider null values distinct (so the index can contain multiple null values in a column, the default PostgreSQL behavior). If it is true, it will consider null values to be equal (so the index can only contain one null value in a column).</p></td>
</tr>
<tr>
<td><p><code>indisprimary</code> <code>bool</code></p>
<p>If true, this index represents the primary key of the table (<code>indisunique</code> should always be true when this is true)</p></td>
</tr>
<tr>
<td><p><code>indisexclusion</code> <code>bool</code></p>
<p>If true, this index supports an exclusion constraint</p></td>
</tr>
<tr>
<td><p><code>indimmediate</code> <code>bool</code></p>
<p>If true, the uniqueness check is enforced immediately on insertion (irrelevant if <code>indisunique</code> is not true)</p></td>
</tr>
<tr>
<td><p><code>indisclustered</code> <code>bool</code></p>
<p>If true, the table was last clustered on this index</p></td>
</tr>
<tr>
<td><p><code>indisvalid</code> <code>bool</code></p>
<p>If true, the index is currently valid for queries. False means the index is possibly incomplete: it must still be modified by <a href="../../reference/sql-commands/insert.md#sql-insert"><code>INSERT</code></a>/<a href="../../reference/sql-commands/update.md#sql-update"><code>UPDATE</code></a> operations, but it cannot safely be used for queries. If it is unique, the uniqueness property is not guaranteed true either.</p></td>
</tr>
<tr>
<td><p><code>indcheckxmin</code> <code>bool</code></p>
<p>If true, queries must not use the index until the <code>xmin</code> of this <code>pg_index</code> row is below their <code>TransactionXmin</code> event horizon, because the table may contain broken <a href="../database-physical-storage/heap-only-tuples-hot.md#storage-hot">HOT chains</a> with incompatible rows that they can see</p></td>
</tr>
<tr>
<td><p><code>indisready</code> <code>bool</code></p>
<p>If true, the index is currently ready for inserts. False means the index must be ignored by <a href="../../reference/sql-commands/insert.md#sql-insert"><code>INSERT</code></a>/<a href="../../reference/sql-commands/update.md#sql-update"><code>UPDATE</code></a> operations.</p></td>
</tr>
<tr>
<td><p><code>indislive</code> <code>bool</code></p>
<p>If false, the index is in process of being dropped, and should be ignored for all purposes (including HOT-safety decisions)</p></td>
</tr>
<tr>
<td><p><code>indisreplident</code> <code>bool</code></p>
<p>If true this index has been chosen as “replica identity” using <a href="../../reference/sql-commands/alter-table.md#sql-altertable-replica-identity"><code>ALTER TABLE ... REPLICA IDENTITY USING INDEX ...</code></a></p></td>
</tr>
<tr>
<td><p><code>indkey</code> <code>int2vector</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>This is an array of <code>indnatts</code> values that indicate which table columns this index indexes. For example, a value of <code>1 3</code> would mean that the first and the third table columns make up the index entries. Key columns come before non-key (included) columns. A zero in this array indicates that the corresponding index attribute is an expression over the table columns, rather than a simple column reference.</p></td>
</tr>
<tr>
<td><p><code>indcollation</code> <code>oidvector</code> (references <a href="pg_collation.md#catalog-pg-collation"><code>pg_collation</code></a>.<code>oid</code>)</p>
<p>For each column in the index key (<code>indnkeyatts</code> values), this contains the OID of the collation to use for the index, or zero if the column is not of a collatable data type.</p></td>
</tr>
<tr>
<td><p><code>indclass</code> <code>oidvector</code> (references <a href="pg_opclass.md#catalog-pg-opclass"><code>pg_opclass</code></a>.<code>oid</code>)</p>
<p>For each column in the index key (<code>indnkeyatts</code> values), this contains the OID of the operator class to use. See <a href="pg_opclass.md#catalog-pg-opclass"><code>pg_opclass</code></a> for details.</p></td>
</tr>
<tr>
<td><p><code>indoption</code> <code>int2vector</code></p>
<p>This is an array of <code>indnkeyatts</code> values that store per-column flag bits. The meaning of the bits is defined by the index's access method.</p></td>
</tr>
<tr>
<td><p><code>indexprs</code> <code>pg_node_tree</code></p>
<p>Expression trees (in <code>nodeToString()</code> representation) for index attributes that are not simple column references. This is a list with one element for each zero entry in <code>indkey</code>. Null if all index attributes are simple references.</p></td>
</tr>
<tr>
<td><p><code>indpred</code> <code>pg_node_tree</code></p>
<p>Expression tree (in <code>nodeToString()</code> representation) for partial index predicate. Null if not a partial index.</p></td>
</tr>
</tbody>
</table>
