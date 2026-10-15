<a id="view-pg-views"></a>

## `pg_views`


 The view `pg_views` provides access to useful information about each view in the database.


**Table: `pg_views` Columns**

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
<p>Name of schema containing view</p></td>
</tr>
<tr>
<td><p><code>viewname</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of view</p></td>
</tr>
<tr>
<td><p><code>viewowner</code> <code>name</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>rolname</code>)</p>
<p>Name of view's owner</p></td>
</tr>
<tr>
<td><p><code>definition</code> <code>text</code></p>
<p>View definition (a reconstructed <a href="../../reference/sql-commands/select.md#sql-select">sql-select</a> query)</p></td>
</tr>
</tbody>
</table>
