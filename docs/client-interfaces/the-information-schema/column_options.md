<a id="infoschema-column-options"></a>

## `column_options`


 The view `column_options` contains all the options defined for foreign table columns in the current database. Only those foreign table columns are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `column_options` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the foreign table (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the foreign table</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign table</p></td>
</tr>
<tr>
<td><p><code>column_name</code> <code>sql_identifier</code></p>
<p>Name of the column</p></td>
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
