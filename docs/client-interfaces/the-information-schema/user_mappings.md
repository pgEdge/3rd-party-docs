<a id="infoschema-user-mappings"></a>

## `user_mappings`


 The view `user_mappings` contains all user mappings defined in the current database. Only those user mappings are shown where the current user has access to the corresponding foreign server (by way of being the owner or having some privilege).


**Table: `user_mappings` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>authorization_identifier</code> <code>sql_identifier</code></p>
<p>Name of the user being mapped, or <code>PUBLIC</code> if the mapping is public</p></td>
</tr>
<tr>
<td><p><code>foreign_server_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the foreign server used by this mapping is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_server_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign server used by this mapping</p></td>
</tr>
</tbody>
</table>
