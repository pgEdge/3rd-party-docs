<a id="view-pg-stats-ext"></a>

## `pg_stats_ext`


 The view `pg_stats_ext` provides access to information about each extended statistics object in the database, combining information stored in the [`pg_statistic_ext`](../system-catalogs/pg_statistic_ext.md#catalog-pg-statistic-ext) and [`pg_statistic_ext_data`](../system-catalogs/pg_statistic_ext_data.md#catalog-pg-statistic-ext-data) catalogs. This view allows access only to rows of [`pg_statistic_ext`](../system-catalogs/pg_statistic_ext.md#catalog-pg-statistic-ext) and [`pg_statistic_ext_data`](../system-catalogs/pg_statistic_ext_data.md#catalog-pg-statistic-ext-data) that correspond to tables the user owns, and therefore it is safe to allow public read access to this view.


 `pg_stats_ext` is also designed to present the information in a more readable format than the underlying catalogs — at the cost that its schema must be extended whenever new types of extended statistics are added to [`pg_statistic_ext`](../system-catalogs/pg_statistic_ext.md#catalog-pg-statistic-ext).


**Table: `pg_stats_ext` Columns**

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
<td><p><code>statistics_schemaname</code> <code>name</code> (references <a href="../system-catalogs/pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>nspname</code>)</p>
<p>Name of schema containing extended statistics object</p></td>
</tr>
<tr>
<td><p><code>statistics_name</code> <code>name</code> (references <a href="../system-catalogs/pg_statistic_ext.md#catalog-pg-statistic-ext"><code>pg_statistic_ext</code></a>.<code>stxname</code>)</p>
<p>Name of extended statistics object</p></td>
</tr>
<tr>
<td><p><code>statistics_owner</code> <code>name</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>rolname</code>)</p>
<p>Owner of the extended statistics object</p></td>
</tr>
<tr>
<td><p><code>attnames</code> <code>name[]</code> (references <a href="../system-catalogs/pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attname</code>)</p>
<p>Names of the columns included in the extended statistics object</p></td>
</tr>
<tr>
<td><p><code>exprs</code> <code>text[]</code></p>
<p>Expressions included in the extended statistics object</p></td>
</tr>
<tr>
<td><p><code>kinds</code> <code>char[]</code></p>
<p>Types of extended statistics object enabled for this record</p></td>
</tr>
<tr>
<td><p><code>inherited</code> <code>bool</code> (references <a href="../system-catalogs/pg_statistic_ext_data.md#catalog-pg-statistic-ext-data"><code>pg_statistic_ext_data</code></a>.<code>stxdinherit</code>)</p>
<p>If true, the stats include values from child tables, not just the values in the specified relation</p></td>
</tr>
<tr>
<td><p><code>n_distinct</code> <code>pg_ndistinct</code></p>
<p>N-distinct counts for combinations of column values. If greater than zero, the estimated number of distinct values in the combination. If less than zero, the negative of the number of distinct values divided by the number of rows. (The negated form is used when <code>ANALYZE</code> believes that the number of distinct values is likely to increase as the table grows; the positive form is used when the column seems to have a fixed number of possible values.) For example, -1 indicates a unique combination of columns in which the number of distinct combinations is the same as the number of rows.</p></td>
</tr>
<tr>
<td><p><code>dependencies</code> <code>pg_dependencies</code></p>
<p>Functional dependency statistics</p></td>
</tr>
<tr>
<td><p><code>most_common_vals</code> <code>text[]</code></p>
<p>A list of the most common combinations of values in the columns. (Null if no combinations seem to be more common than any others.)</p></td>
</tr>
<tr>
<td><p><code>most_common_val_nulls</code> <code>bool[]</code></p>
<p>A list of NULL flags for the most common combinations of values. (Null when <code>most_common_vals</code> is.)</p></td>
</tr>
<tr>
<td><p><code>most_common_freqs</code> <code>float8[]</code></p>
<p>A list of the frequencies of the most common combinations, i.e., number of occurrences of each divided by total number of rows. (Null when <code>most_common_vals</code> is.)</p></td>
</tr>
<tr>
<td><p><code>most_common_base_freqs</code> <code>float8[]</code></p>
<p>A list of the base frequencies of the most common combinations, i.e., product of per-value frequencies. (Null when <code>most_common_vals</code> is.)</p></td>
</tr>
</tbody>
</table>


 The maximum number of entries in the array fields can be controlled on a column-by-column basis using the [`ALTER TABLE SET STATISTICS`](../../reference/sql-commands/alter-table.md#sql-altertable) command, or globally by setting the [default_statistics_target](../../server-administration/server-configuration/query-planning.md#guc-default-statistics-target) run-time parameter.
