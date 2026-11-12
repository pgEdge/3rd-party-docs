<a id="view-pg-matviews"></a>

## `pg_matviews`


 The view `pg_matviews` provides access to useful information about each materialized view in the database.


**Table: `pg_matviews` Columns**

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
<p>Name of schema containing materialized view</p></td>
</tr>
<tr>
<td><p><code>matviewname</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of materialized view</p></td>
</tr>
<tr>
<td><p><code>matviewowner</code> <code>name</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>rolname</code>)</p>
<p>Name of materialized view's owner</p></td>
</tr>
<tr>
<td><p><code>tablespace</code> <code>name</code> (references <a href="../system-catalogs/pg_tablespace.md#catalog-pg-tablespace"><code>pg_tablespace</code></a>.<code>spcname</code>)</p>
<p>Name of tablespace containing materialized view (null if default for database)</p></td>
</tr>
<tr>
<td><p><code>hasindexes</code> <code>bool</code></p>
<p>True if materialized view has (or recently had) any indexes</p></td>
</tr>
<tr>
<td><p><code>ispopulated</code> <code>bool</code></p>
<p>True if materialized view is currently populated</p></td>
</tr>
<tr>
<td><p><code>definition</code> <code>text</code></p>
<p>Materialized view definition (a reconstructed <a href="../../reference/sql-commands/select.md#sql-select">sql-select</a> query)</p></td>
</tr>
</tbody>
</table>
