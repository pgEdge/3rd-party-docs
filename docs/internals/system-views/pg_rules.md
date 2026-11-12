<a id="view-pg-rules"></a>

## `pg_rules`


 The view `pg_rules` provides access to useful information about query rewrite rules.


**Table: `pg_rules` Columns**

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
<p>Name of table the rule is for</p></td>
</tr>
<tr>
<td><p><code>rulename</code> <code>name</code> (references <a href="../system-catalogs/pg_rewrite.md#catalog-pg-rewrite"><code>pg_rewrite</code></a>.<code>rulename</code>)</p>
<p>Name of rule</p></td>
</tr>
<tr>
<td><p><code>definition</code> <code>text</code></p>
<p>Rule definition (a reconstructed creation command)</p></td>
</tr>
</tbody>
</table>


 The `pg_rules` view excludes the `ON SELECT` rules of views and materialized views; those can be seen in [`pg_views`](pg_views.md#view-pg-views) and [`pg_matviews`](pg_matviews.md#view-pg-matviews).
