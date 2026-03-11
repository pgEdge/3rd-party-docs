<a id="infoschema-foreign-tables"></a>

## `foreign_tables`


 The view `foreign_tables` contains all foreign tables defined in the current database. Only those foreign tables are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `foreign_tables` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>foreign_table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the foreign table is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the foreign table</p></td>
</tr>
<tr>
<td><p><code>foreign_table_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign table</p></td>
</tr>
<tr>
<td><p><code>foreign_server_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the foreign server is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_server_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign server</p></td>
</tr>
</tbody>
</table>
