<a id="infoschema-pg-element-tables"></a>

## `pg_element_tables`


 The view `pg_element_tables` contains information about the element tables of property graphs defined in the current database. Only those property graphs are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `pg_element_tables` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>property_graph_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the property graph (always the current database)</p></td>
</tr>
<tr>
<td><p><code>property_graph_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the property graph</p></td>
</tr>
<tr>
<td><p><code>property_graph_name</code> <code>sql_identifier</code></p>
<p>Name of the property graph</p></td>
</tr>
<tr>
<td><p><code>element_table_alias</code> <code>sql_identifier</code></p>
<p>Element table alias (unique identifier of an element table within a property graph)</p></td>
</tr>
<tr>
<td><p><code>element_table_kind</code> <code>character_data</code></p>
<p>The kind of the element table: <code>EDGE</code> or <code>VERTEX</code></p></td>
</tr>
<tr>
<td><p><code>table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the referenced table (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the referenced table</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the table being referenced by the element table definition</p></td>
</tr>
<tr>
<td><p><code>element_table_definition</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
</tbody>
</table>
