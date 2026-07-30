<a id="view-pg-indexes"></a>

## `pg_indexes`


 The view `pg_indexes` provides access to useful information about each index in the database.


**Table: `pg_indexes` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>schemaname</code> <code>name</code> (references <a href="../system-catalogs/pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>nspname</code>)</p>
<p>Name of schema containing table and index</p></td>
</tr>
<tr>
<td><p><code>tablename</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of table the index is for</p></td>
</tr>
<tr>
<td><p><code>indexname</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of index</p></td>
</tr>
<tr>
<td><p><code>tablespace</code> <code>name</code> (references <a href="../system-catalogs/pg_tablespace.md#catalog-pg-tablespace"><code>pg_tablespace</code></a>.<code>spcname</code>)</p>
<p>Name of tablespace containing index (null if default for database)</p></td>
</tr>
<tr>
<td><p><code>indexdef</code> <code>text</code></p>
<p>Index definition (a reconstructed <a href="../../reference/sql-commands/create-index.md#sql-createindex">sql-createindex</a> command)</p></td>
</tr>
</tbody>
</table>
