<a id="view-pg-ident-file-mappings"></a>

## `pg_ident_file_mappings`


 The view `pg_ident_file_mappings` provides a summary of the contents of the client user name mapping configuration file, [`pg_ident.conf`](../../server-administration/client-authentication/user-name-maps.md#auth-username-maps). A row appears in this view for each non-empty, non-comment line in the file, with annotations indicating whether the map could be applied successfully.


 This view can be helpful for checking whether planned changes in the authentication configuration file will work, or for diagnosing a previous failure. Note that this view reports on the *current* contents of the file, not on what was last loaded by the server.


 By default, the `pg_ident_file_mappings` view can be read only by superusers.


**Table: `pg_ident_file_mappings` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>map_number</code> <code>int4</code></p>
<p>Number of this map, in priority order, if valid, otherwise <code>NULL</code></p></td>
</tr>
<tr>
<td><p><code>file_name</code> <code>text</code></p>
<p>Name of the file containing this map</p></td>
</tr>
<tr>
<td><p><code>line_number</code> <code>int4</code></p>
<p>Line number of this map in <code>file_name</code></p></td>
</tr>
<tr>
<td><p><code>map_name</code> <code>text</code></p>
<p>Name of the map</p></td>
</tr>
<tr>
<td><p><code>sys_name</code> <code>text</code></p>
<p>Detected user name of the client</p></td>
</tr>
<tr>
<td><p><code>pg_username</code> <code>text</code></p>
<p>Requested PostgreSQL user name</p></td>
</tr>
<tr>
<td><p><code>error</code> <code>text</code></p>
<p>If not <code>NULL</code>, an error message indicating why this line could not be processed</p></td>
</tr>
</tbody>
</table>


 Usually, a row reflecting an incorrect entry will have values for only the `line_number` and `error` fields.


 See [Client Authentication](../../server-administration/client-authentication/index.md#client-authentication) for more information about client authentication configuration.
