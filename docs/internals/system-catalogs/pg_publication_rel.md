<a id="catalog-pg-publication-rel"></a>

## `pg_publication_rel`


 The catalog `pg_publication_rel` contains the mapping between relations and publications in the database. This is a many-to-many mapping. See also [`pg_publication_tables`](../system-views/pg_publication_tables.md#view-pg-publication-tables) for a more user-friendly view of this information.


**Table: `pg_publication_rel` Columns**

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
<td><p><code>prpubid</code> <code>oid</code> (references <a href="pg_publication.md#catalog-pg-publication"><code>pg_publication</code></a>.<code>oid</code>)</p>
<p>Reference to publication</p></td>
</tr>
<tr>
<td><p><code>prrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Reference to relation</p></td>
</tr>
<tr>
<td><p><code>prqual</code> <code>pg_node_tree</code></p>
<p>Expression tree (in <code>nodeToString()</code> representation) for the relation's publication qualifying condition. Null if there is no publication qualifying condition.</p></td>
</tr>
<tr>
<td><p><code>prattrs</code> <code>int2vector</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>This is an array of values that indicates which table columns are part of the publication. For example, a value of <code>1 3</code> would mean that the first and the third table columns are published. A null value indicates that all columns are published.</p></td>
</tr>
</tbody>
</table>
