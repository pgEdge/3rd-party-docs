<a id="infoschema-foreign-server-options"></a>

## `foreign_server_options`


 The view `foreign_server_options` contains all the options defined for foreign servers in the current database. Only those foreign servers are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `foreign_server_options` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>foreign_server_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the foreign server is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_server_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign server</p></td>
</tr>
<tr>
<td><p><code>option_name</code> <code>sql_identifier</code></p>
<p>Name of an option</p></td>
</tr>
<tr>
<td><p><code>option_value</code> <code>character_data</code></p>
<p>Value of the option</p></td>
</tr>
</tbody>
</table>
