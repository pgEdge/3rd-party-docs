<a id="view-pg-timezone-abbrevs"></a>

## `pg_timezone_abbrevs`


 The view `pg_timezone_abbrevs` provides a list of time zone abbreviations that are currently recognized by the datetime input routines. The contents of this view change when the [TimeZone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) or [timezone_abbreviations](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone-abbreviations) run-time parameters are modified.


**Table: `pg_timezone_abbrevs` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
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
<p>True if this is a daylight-savings abbreviation</p></td>
</tr>
</tbody>
</table>


 While most timezone abbreviations represent fixed offsets from UTC, there are some that have historically varied in value (see [Date/Time Configuration Files](../../appendixes/date-time-support/date-time-configuration-files.md#datetime-config-files) for more information). In such cases this view presents their current meaning.
