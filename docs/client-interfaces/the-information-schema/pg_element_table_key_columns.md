<a id="infoschema-pg-element-table-key-columns"></a>

## `pg_element_table_key_columns`


 The view `pg_element_table_key_columns` identifies which columns are part of the keys of the element tables of property graphs defined in the current database. Only those property graphs are shown that the current user has access to (by way of being the owner or having some privilege).


 The key of an element table uniquely identifies the rows in it. It is either specified using the `KEY` clause in `CREATE PROPERTY GRAPH` or defaults to the primary key.


**Table: `pg_element_table_key_columns` Columns**

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
<td><p><code>column_name</code> <code>sql_identifier</code></p>
<p>Name of the column that is part of the key</p></td>
</tr>
<tr>
<td><p><code>ordinal_position</code> <code>cardinal_number</code></p>
<p>Ordinal position of the column within the key (count starts at 1)</p></td>
</tr>
</tbody>
</table>
