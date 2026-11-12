<a id="infoschema-foreign-servers"></a>

## `foreign_servers`


 The view `foreign_servers` contains all foreign servers defined in the current database. Only those foreign servers are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `foreign_servers` Columns**

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
<td><p><code>foreign_data_wrapper_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the foreign-data wrapper used by the foreign server (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_data_wrapper_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign-data wrapper used by the foreign server</p></td>
</tr>
<tr>
<td><p><code>foreign_server_type</code> <code>character_data</code></p>
<p>Foreign server type information, if specified upon creation</p></td>
</tr>
<tr>
<td><p><code>foreign_server_version</code> <code>character_data</code></p>
<p>Foreign server version information, if specified upon creation</p></td>
</tr>
<tr>
<td><p><code>authorization_identifier</code> <code>sql_identifier</code></p>
<p>Name of the owner of the foreign server</p></td>
</tr>
</tbody>
</table>
