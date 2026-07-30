<a id="catalog-pg-foreign-table"></a>

## `pg_foreign_table`


 The catalog `pg_foreign_table` contains auxiliary information about foreign tables. A foreign table is primarily represented by a [`pg_class`](pg_class.md#catalog-pg-class) entry, just like a regular table. Its `pg_foreign_table` entry contains the information that is pertinent only to foreign tables and not any other kind of relation.


**Table: `pg_foreign_table` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>ftrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry for this foreign table</p></td>
</tr>
<tr>
<td><p><code>ftserver</code> <code>oid</code> (references <a href="pg_foreign_server.md#catalog-pg-foreign-server"><code>pg_foreign_server</code></a>.<code>oid</code>)</p>
<p>OID of the foreign server for this foreign table</p></td>
</tr>
<tr>
<td><p><code>ftoptions</code> <code>text[]</code></p>
<p>Foreign table options, as “keyword=value” strings</p></td>
</tr>
</tbody>
</table>
