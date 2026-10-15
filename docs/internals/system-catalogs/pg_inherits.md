<a id="catalog-pg-inherits"></a>

## `pg_inherits`


 The catalog `pg_inherits` records information about table and index inheritance hierarchies. There is one entry for each direct parent-child table or index relationship in the database. (Indirect inheritance can be determined by following chains of entries.)


**Table: `pg_inherits` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>inhrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the child table or index</p></td>
</tr>
<tr>
<td><p><code>inhparent</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the parent table or index</p></td>
</tr>
<tr>
<td><p><code>inhseqno</code> <code>int4</code></p>
<p>If there is more than one direct parent for a child table (multiple inheritance), this number tells the order in which the inherited columns are to be arranged. The count starts at 1.</p>
<p>Indexes cannot have multiple inheritance, since they can only inherit when using declarative partitioning.</p></td>
</tr>
<tr>
<td><p><code>inhdetachpending</code> <code>bool</code></p>
<p><code>true</code> for a partition that is in the process of being detached; <code>false</code> otherwise.</p></td>
</tr>
</tbody>
</table>
