<a id="view-pg-tables"></a>

## `pg_tables`


 The view `pg_tables` provides access to useful information about each table in the database.


**Table: `pg_tables` Columns**

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
<p>Name of schema containing table</p></td>
</tr>
<tr>
<td><p><code>tablename</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of table</p></td>
</tr>
<tr>
<td><p><code>tableowner</code> <code>name</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>rolname</code>)</p>
<p>Name of table's owner</p></td>
</tr>
<tr>
<td><p><code>tablespace</code> <code>name</code> (references <a href="../system-catalogs/pg_tablespace.md#catalog-pg-tablespace"><code>pg_tablespace</code></a>.<code>spcname</code>)</p>
<p>Name of tablespace containing table (null if default for database)</p></td>
</tr>
<tr>
<td><p><code>hasindexes</code> <code>bool</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relhasindex</code>)</p>
<p>True if table has (or recently had) any indexes</p></td>
</tr>
<tr>
<td><p><code>hasrules</code> <code>bool</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relhasrules</code>)</p>
<p>True if table has (or once had) rules</p></td>
</tr>
<tr>
<td><p><code>hastriggers</code> <code>bool</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relhastriggers</code>)</p>
<p>True if table has (or once had) triggers</p></td>
</tr>
<tr>
<td><p><code>rowsecurity</code> <code>bool</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relrowsecurity</code>)</p>
<p>True if row security is enabled on the table</p></td>
</tr>
</tbody>
</table>
