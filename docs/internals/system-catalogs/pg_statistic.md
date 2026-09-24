<a id="catalog-pg-statistic"></a>

## `pg_statistic`


 The catalog `pg_statistic` stores statistical data about the contents of the database. Entries are created by [`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze) and subsequently used by the query planner. Note that all the statistical data is inherently approximate, even assuming that it is up-to-date.


 Normally there is one entry, with `stainherit` = `false`, for each table column that has been analyzed. If the table has inheritance children or partitions, a second entry with `stainherit` = `true` is also created. This row represents the column's statistics over the inheritance tree, i.e., statistics for the data you'd see with <code>SELECT </code><em>column</em><code> FROM </code><em>table</em><code>*</code>, whereas the `stainherit` = `false` row represents the results of <code>SELECT </code><em>column</em><code> FROM ONLY </code><em>table</em>.


 `pg_statistic` also stores statistical data about the values of index expressions. These are described as if they were actual data columns; in particular, `starelid` references the index. No entry is made for an ordinary non-expression index column, however, since it would be redundant with the entry for the underlying table column. Currently, entries for index expressions always have `stainherit` = `false`.


 Since different kinds of statistics might be appropriate for different kinds of data, `pg_statistic` is designed not to assume very much about what sort of statistics it stores. Only extremely general statistics (such as nullness) are given dedicated columns in `pg_statistic`. Everything else is stored in “slots”, which are groups of associated columns whose content is identified by a code number in one of the slot's columns. For more information see `src/include/catalog/pg_statistic.h`.


 `pg_statistic` should not be readable by the public, since even statistical information about a table's contents might be considered sensitive. (Example: minimum and maximum values of a salary column might be quite interesting.) [`pg_stats`](../system-views/pg_stats.md#view-pg-stats) is a publicly readable view on `pg_statistic` that only exposes information about those tables that are readable by the current user.


**Table: `pg_statistic` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>starelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table or index that the described column belongs to</p></td>
</tr>
<tr>
<td><p><code>staattnum</code> <code>int2</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>The number of the described column</p></td>
</tr>
<tr>
<td><p><code>stainherit</code> <code>bool</code></p>
<p>If true, the stats include values from child tables, not just the values in the specified relation</p></td>
</tr>
<tr>
<td><p><code>stanullfrac</code> <code>float4</code></p>
<p>The fraction of the column's entries that are null</p></td>
</tr>
<tr>
<td><p><code>stawidth</code> <code>int4</code></p>
<p>The average stored width, in bytes, of nonnull entries</p></td>
</tr>
<tr>
<td><p><code>stadistinct</code> <code>float4</code></p>
<p>The number of distinct nonnull data values in the column. A value greater than zero is the actual number of distinct values. A value less than zero is the negative of a multiplier for the number of rows in the table; for example, a column in which about 80% of the values are nonnull and each nonnull value appears about twice on average could be represented by <code>stadistinct</code> = -0.4. A zero value means the number of distinct values is unknown.</p></td>
</tr>
<tr>
<td><p><code>stakind</code><em>N</em> <code>int2</code></p>
<p>A code number indicating the kind of statistics stored in the <em>N</em>th “slot” of the <code>pg_statistic</code> row.</p></td>
</tr>
<tr>
<td><p><code>staop</code><em>N</em> <code>oid</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>An operator used to derive the statistics stored in the <em>N</em>th “slot”. For example, a histogram slot would show the <code>&lt;</code> operator that defines the sort order of the data. Zero if the statistics kind does not require an operator.</p></td>
</tr>
<tr>
<td><p><code>stacoll</code><em>N</em> <code>oid</code> (references <a href="pg_collation.md#catalog-pg-collation"><code>pg_collation</code></a>.<code>oid</code>)</p>
<p>The collation used to derive the statistics stored in the <em>N</em>th “slot”. For example, a histogram slot for a collatable column would show the collation that defines the sort order of the data. Zero for noncollatable data.</p></td>
</tr>
<tr>
<td><p><code>stanumbers</code><em>N</em> <code>float4[]</code></p>
<p>Numerical statistics of the appropriate kind for the <em>N</em>th “slot”, or null if the slot kind does not involve numerical values</p></td>
</tr>
<tr>
<td><p><code>stavalues</code><em>N</em> <code>anyarray</code></p>
<p>Column data values of the appropriate kind for the <em>N</em>th “slot”, or null if the slot kind does not store any data values. Each array's element values are actually of the specific column's data type, or a related type such as an array's element type, so there is no way to define these columns' type more specifically than <code>anyarray</code>.</p></td>
</tr>
</tbody>
</table>
