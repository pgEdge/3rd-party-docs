<a id="infoschema-pg-labels"></a>

## `pg_labels`


 The view `pg_labels` contains all the labels defined in property graphs defined in the current database. Only those property graphs are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `pg_labels` Columns**

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
<td><p><code>label_name</code> <code>sql_identifier</code></p>
<p>Name of the label</p></td>
</tr>
</tbody>
</table>
