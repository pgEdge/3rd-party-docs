<a id="infoschema-user-mapping-options"></a>

## `user_mapping_options`


 The view `user_mapping_options` contains all the options defined for user mappings in the current database. Only those user mappings are shown where the current user has access to the corresponding foreign server (by way of being the owner or having some privilege).


**Table: `user_mapping_options` Columns**

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
<tr>
<td><p><code>option_name</code> <code>sql_identifier</code></p>
<p>Name of an option</p></td>
</tr>
<tr>
<td><p><code>option_value</code> <code>character_data</code></p>
<p>Value of the option. This column will show as null unless the current user is the user being mapped, or the mapping is for <code>PUBLIC</code> and the current user is the server owner, or the current user is a superuser. The intent is to protect password information stored as user mapping option.</p></td>
</tr>
</tbody>
</table>
