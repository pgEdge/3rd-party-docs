<a id="infoschema-foreign-table-options"></a>

## `foreign_table_options`


 The view `foreign_table_options` contains all the options defined for foreign tables in the current database. Only those foreign tables are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `foreign_table_options` Columns**

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
<p>Name of the database that contains the foreign table (always the current database)</p></td>
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
<td><p><code>option_name</code> <code>sql_identifier</code></p>
<p>Name of an option</p></td>
</tr>
<tr>
<td><p><code>option_value</code> <code>character_data</code></p>
<p>Value of the option</p></td>
</tr>
</tbody>
</table>
