<a id="view-pg-publication-sequences"></a>

## `pg_publication_sequences`


 The view `pg_publication_sequences` provides information about the mapping between publications and sequences.


**Table: `pg_publication_sequences` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pubname</code> <code>name</code> (references <a href="../system-catalogs/pg_publication.md#catalog-pg-publication"><code>pg_publication</code></a>.<code>pubname</code>)</p>
<p>Name of publication</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code> (references <a href="../system-catalogs/pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>nspname</code>)</p>
<p>Name of schema containing sequence</p></td>
</tr>
<tr>
<td><p><code>sequencename</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of sequence</p></td>
</tr>
</tbody>
</table>
