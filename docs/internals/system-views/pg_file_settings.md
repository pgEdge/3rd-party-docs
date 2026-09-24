<a id="view-pg-file-settings"></a>

## `pg_file_settings`


 The view `pg_file_settings` provides a summary of the contents of the server's configuration file(s). A row appears in this view for each “name = value” entry appearing in the files, with annotations indicating whether the value could be applied successfully. Additional row(s) may appear for problems not linked to a “name = value” entry, such as syntax errors in the files.


 This view is helpful for checking whether planned changes in the configuration files will work, or for diagnosing a previous failure. Note that this view reports on the *current* contents of the files, not on what was last applied by the server. (The [`pg_settings`](pg_settings.md#view-pg-settings) view is usually sufficient to determine that.)


 By default, the `pg_file_settings` view can be read only by superusers.


**Table: `pg_file_settings` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>sourcefile</code> <code>text</code></p>
<p>Full path name of the configuration file</p></td>
</tr>
<tr>
<td><p><code>sourceline</code> <code>int4</code></p>
<p>Line number within the configuration file where the entry appears</p></td>
</tr>
<tr>
<td><p><code>seqno</code> <code>int4</code></p>
<p>Order in which the entries are processed (1..<em>n</em>)</p></td>
</tr>
<tr>
<td><p><code>name</code> <code>text</code></p>
<p>Configuration parameter name</p></td>
</tr>
<tr>
<td><p><code>setting</code> <code>text</code></p>
<p>Value to be assigned to the parameter</p></td>
</tr>
<tr>
<td><p><code>applied</code> <code>bool</code></p>
<p>True if the value can be applied successfully</p></td>
</tr>
<tr>
<td><p><code>error</code> <code>text</code></p>
<p>If not null, an error message indicating why this entry could not be applied</p></td>
</tr>
</tbody>
</table>


 If the configuration file contains syntax errors or invalid parameter names, the server will not attempt to apply any settings from it, and therefore all the `applied` fields will read as false. In such a case there will be one or more rows with non-null `error` fields indicating the problem(s). Otherwise, individual settings will be applied if possible. If an individual setting cannot be applied (e.g., invalid value, or the setting cannot be changed after server start) it will have an appropriate message in the `error` field. Another way that an entry might have `applied` = false is that it is overridden by a later entry for the same parameter name; this case is not considered an error so nothing appears in the `error` field.


 See [Setting Parameters](../../server-administration/server-configuration/setting-parameters.md#config-setting) for more information about the various ways to change run-time parameters.
