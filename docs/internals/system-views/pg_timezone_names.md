<a id="view-pg-timezone-names"></a>

## `pg_timezone_names`


 The view `pg_timezone_names` provides a list of time zone names that are recognized by `SET TIMEZONE`, along with their associated abbreviations, UTC offsets, and daylight-savings status. (Technically, PostgreSQL does not use UTC because leap seconds are not handled.) Unlike the abbreviations shown in [`pg_timezone_abbrevs`](pg_timezone_abbrevs.md#view-pg-timezone-abbrevs), many of these names imply a set of daylight-savings transition date rules. Therefore, the associated information changes across local DST boundaries. The displayed information is computed based on the current value of `CURRENT_TIMESTAMP`.


**Table: `pg_timezone_names` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>name</code> <code>text</code></p>
<p>Time zone name</p></td>
</tr>
<tr>
<td><p><code>abbrev</code> <code>text</code></p>
<p>Time zone abbreviation</p></td>
</tr>
<tr>
<td><p><code>utc_offset</code> <code>interval</code></p>
<p>Offset from UTC (positive means east of Greenwich)</p></td>
</tr>
<tr>
<td><p><code>is_dst</code> <code>bool</code></p>
<p>True if currently observing daylight savings</p></td>
</tr>
</tbody>
</table>
