<a id="catalog-pg-sequence"></a>

## `pg_sequence`


 The catalog `pg_sequence` contains information about sequences. Some of the information about sequences, such as the name and the schema, is in [`pg_class`](pg_class.md#catalog-pg-class)


**Table: `pg_sequence` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>seqrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry for this sequence</p></td>
</tr>
<tr>
<td><p><code>seqtypid</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type of the sequence</p></td>
</tr>
<tr>
<td><p><code>seqstart</code> <code>int8</code></p>
<p>Start value of the sequence</p></td>
</tr>
<tr>
<td><p><code>seqincrement</code> <code>int8</code></p>
<p>Increment value of the sequence</p></td>
</tr>
<tr>
<td><p><code>seqmax</code> <code>int8</code></p>
<p>Maximum value of the sequence</p></td>
</tr>
<tr>
<td><p><code>seqmin</code> <code>int8</code></p>
<p>Minimum value of the sequence</p></td>
</tr>
<tr>
<td><p><code>seqcache</code> <code>int8</code></p>
<p>Cache size of the sequence</p></td>
</tr>
<tr>
<td><p><code>seqcycle</code> <code>bool</code></p>
<p>Whether the sequence cycles</p></td>
</tr>
</tbody>
</table>
