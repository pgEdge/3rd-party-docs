<a id="infoschema-schemata"></a>

## `schemata`


 The view `schemata` contains all schemas in the current database that the current user has access to (by way of being the owner or having some privilege).


**Table: `schemata` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>catalog_name</code> <code>sql_identifier</code></p>
<p>Name of the database that the schema is contained in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>schema_name</code> <code>sql_identifier</code></p>
<p>Name of the schema</p></td>
</tr>
<tr>
<td><p><code>schema_owner</code> <code>sql_identifier</code></p>
<p>Name of the owner of the schema</p></td>
</tr>
<tr>
<td><p><code>default_character_set_catalog</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>default_character_set_schema</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>default_character_set_name</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>sql_path</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
</tbody>
</table>
