<a id="infoschema-pg-element-table-properties"></a>

## `pg_element_table_properties`


 The view `pg_element_table_properties` shows the definitions of the properties for the element tables of property graphs defined in the current database. Only those property graphs are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `pg_element_table_properties` Columns**

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
<td><p><code>property_name</code> <code>sql_identifier</code></p>
<p>Name of the property</p></td>
</tr>
<tr>
<td><p><code>property_expression</code> <code>character_data</code></p>
<p>Expression of the property definition for this element table</p></td>
</tr>
</tbody>
</table>
