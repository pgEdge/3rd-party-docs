<a id="infoschema-pg-edge-table-components"></a>

## `pg_edge_table_components`


 The view `pg_edge_table_components` identifies which columns are part of the source or destination vertex keys, as well as their corresponding columns in the vertex tables being linked to, in the edge tables of property graphs defined in the current database. Only those property graphs are shown that the current user has access to (by way of being the owner or having some privilege).


 The source and destination vertex links of edge tables are specified in `CREATE PROPERTY GRAPH` and default to foreign keys in certain cases.


**Table: `pg_edge_table_components` Columns**

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
<td><p><code>edge_table_alias</code> <code>sql_identifier</code></p>
<p>The element table alias of the edge table being described</p></td>
</tr>
<tr>
<td><p><code>vertex_table_alias</code> <code>sql_identifier</code></p>
<p>The element table alias of the source or destination vertex table being linked to</p></td>
</tr>
<tr>
<td><p><code>edge_end</code> <code>character_data</code></p>
<p>Either <code>SOURCE</code> or <code>DESTINATION</code>; specifies which edge link is being described.</p></td>
</tr>
<tr>
<td><p><code>edge_table_column_name</code> <code>sql_identifier</code></p>
<p>Name of the column that is part of the source or destination vertex key in this edge table</p></td>
</tr>
<tr>
<td><p><code>vertex_table_column_name</code> <code>sql_identifier</code></p>
<p>Name of the column that is part of the key in the source or destination vertex table being linked to</p></td>
</tr>
<tr>
<td><p><code>ordinal_position</code> <code>cardinal_number</code></p>
<p>Ordinal position of the columns within the key (count starts at 1)</p></td>
</tr>
</tbody>
</table>
