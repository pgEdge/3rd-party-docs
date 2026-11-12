<a id="catalog-pg-statistic-ext-data"></a>

## `pg_statistic_ext_data`


 The catalog `pg_statistic_ext_data` holds data for extended planner statistics defined in [`pg_statistic_ext`](pg_statistic_ext.md#catalog-pg-statistic-ext). Each row in this catalog corresponds to a *statistics object* created with [`CREATE STATISTICS`](../../reference/sql-commands/create-statistics.md#sql-createstatistics).


 Normally there is one entry, with `stxdinherit` = `false`, for each statistics object that has been analyzed. If the table has inheritance children or partitions, a second entry with `stxdinherit` = `true` is also created. This row represents the statistics object over the inheritance tree, i.e., statistics for the data you'd see with <code>SELECT * FROM </code><em>table</em><code>*</code>, whereas the `stxdinherit` = `false` row represents the results of <code>SELECT * FROM ONLY </code><em>table</em>.


 Like [`pg_statistic`](pg_statistic.md#catalog-pg-statistic), `pg_statistic_ext_data` should not be readable by the public, since the contents might be considered sensitive. (Example: most common combinations of values in columns might be quite interesting.) [`pg_stats_ext`](../system-views/pg_stats_ext.md#view-pg-stats-ext) is a publicly readable view on `pg_statistic_ext_data` (after joining with [`pg_statistic_ext`](pg_statistic_ext.md#catalog-pg-statistic-ext)) that only exposes information about tables the current user owns.


**Table: `pg_statistic_ext_data` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>stxoid</code> <code>oid</code> (references <a href="pg_statistic_ext.md#catalog-pg-statistic-ext"><code>pg_statistic_ext</code></a>.<code>oid</code>)</p>
<p>Extended statistics object containing the definition for this data</p></td>
</tr>
<tr>
<td><p><code>stxdinherit</code> <code>bool</code></p>
<p>If true, the stats include values from child tables, not just the values in the specified relation</p></td>
</tr>
<tr>
<td><p><code>stxdndistinct</code> <code>pg_ndistinct</code></p>
<p>N-distinct counts, serialized as <code>pg_ndistinct</code> type</p></td>
</tr>
<tr>
<td><p><code>stxddependencies</code> <code>pg_dependencies</code></p>
<p>Functional dependency statistics, serialized as <code>pg_dependencies</code> type</p></td>
</tr>
<tr>
<td><p><code>stxdmcv</code> <code>pg_mcv_list</code></p>
<p>MCV (most-common values) list statistics, serialized as <code>pg_mcv_list</code> type</p></td>
</tr>
<tr>
<td><p><code>stxdexpr</code> <code>pg_statistic[]</code></p>
<p>Per-expression statistics, serialized as an array of <code>pg_statistic</code> type</p></td>
</tr>
</tbody>
</table>
