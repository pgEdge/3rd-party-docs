<a id="view-pg-publication-tables"></a>

## `pg_publication_tables`


 The view `pg_publication_tables` provides information about the mapping between publications and information of tables they contain. Unlike the underlying catalog [`pg_publication_rel`](../system-catalogs/pg_publication_rel.md#catalog-pg-publication-rel), this view expands publications defined as [`FOR ALL TABLES`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-for-all-tables) and [`FOR TABLES IN SCHEMA`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-for-tables-in-schema), so for such publications there will be a row for each eligible table.


**Table: `pg_publication_tables` Columns**

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
<p>Name of schema containing table</p></td>
</tr>
<tr>
<td><p><code>tablename</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of table</p></td>
</tr>
<tr>
<td><p><code>attnames</code> <code>name[]</code> (references <a href="../system-catalogs/pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attname</code>)</p>
<p>Names of table columns included in the publication. This contains all the columns of the table when the user didn't specify the column list for the table.</p></td>
</tr>
<tr>
<td><p><code>rowfilter</code> <code>text</code></p>
<p>Expression for the table's publication qualifying condition</p></td>
</tr>
</tbody>
</table>
