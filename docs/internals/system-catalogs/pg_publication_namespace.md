<a id="catalog-pg-publication-namespace"></a>

## `pg_publication_namespace`


 The catalog `pg_publication_namespace` contains the mapping between schemas and publications in the database. This is a many-to-many mapping.


**Table: `pg_publication_namespace` Columns**

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
<td><p><code>pnpubid</code> <code>oid</code> (references <a href="pg_publication.md#catalog-pg-publication"><code>pg_publication</code></a>.<code>oid</code>)</p>
<p>Reference to publication</p></td>
</tr>
<tr>
<td><p><code>pnnspid</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>Reference to schema</p></td>
</tr>
</tbody>
</table>
