<a id="view-pg-hba-file-rules"></a>

## `pg_hba_file_rules`


 The view `pg_hba_file_rules` provides a summary of the contents of the client authentication configuration file, [`pg_hba.conf`](../../server-administration/client-authentication/the-pg_hba-conf-file.md#auth-pg-hba-conf). A row appears in this view for each non-empty, non-comment line in the file, with annotations indicating whether the rule could be applied successfully.


 This view can be helpful for checking whether planned changes in the authentication configuration file will work, or for diagnosing a previous failure. Note that this view reports on the *current* contents of the file, not on what was last loaded by the server.


 By default, the `pg_hba_file_rules` view can be read only by superusers.


**Table: `pg_hba_file_rules` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>rule_number</code> <code>int4</code></p>
<p>Number of this rule, if valid, otherwise <code>NULL</code>. This indicates the order in which each rule is considered until a match is found during authentication.</p></td>
</tr>
<tr>
<td><p><code>file_name</code> <code>text</code></p>
<p>Name of the file containing this rule</p></td>
</tr>
<tr>
<td><p><code>line_number</code> <code>int4</code></p>
<p>Line number of this rule in <code>file_name</code></p></td>
</tr>
<tr>
<td><p><code>type</code> <code>text</code></p>
<p>Type of connection</p></td>
</tr>
<tr>
<td><p><code>database</code> <code>text[]</code></p>
<p>List of database name(s) to which this rule applies</p></td>
</tr>
<tr>
<td><p><code>user_name</code> <code>text[]</code></p>
<p>List of user and group name(s) to which this rule applies</p></td>
</tr>
<tr>
<td><p><code>address</code> <code>text</code></p>
<p>Host name or IP address, or one of <code>all</code>, <code>samehost</code>, or <code>samenet</code>, or null for local connections</p></td>
</tr>
<tr>
<td><p><code>netmask</code> <code>text</code></p>
<p>IP address mask, or null if not applicable</p></td>
</tr>
<tr>
<td><p><code>auth_method</code> <code>text</code></p>
<p>Authentication method</p></td>
</tr>
<tr>
<td><p><code>options</code> <code>text[]</code></p>
<p>Options specified for authentication method, if any</p></td>
</tr>
<tr>
<td><p><code>error</code> <code>text</code></p>
<p>If not null, an error message indicating why this line could not be processed</p></td>
</tr>
</tbody>
</table>


 Usually, a row reflecting an incorrect entry will have values for only the `line_number` and `error` fields.


 See [Client Authentication](../../server-administration/client-authentication/index.md#client-authentication) for more information about client authentication configuration.
