<a id="catalog-pg-propgraph-label-property"></a>

## `pg_propgraph_label_property`


 The catalog `pg_propgraph_label_property` stores information about the properties in a property graph that are specific to a label. In particular, this stores the expression that defines the property.


**Table: `pg_propgraph_label_property` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>oid</code> <code>oid</code></p>
<p>Row identifier</p></td>
</tr>
<tr>
<td><p><code>plppropid</code> <code>oid</code> (references <a href="pg_propgraph_property.md#catalog-pg-propgraph-property"><code>pg_propgraph_property</code></a>.<code>oid</code>)</p>
<p>Reference to the property</p></td>
</tr>
<tr>
<td><p><code>plpellabelid</code> <code>oid</code> (references <a href="pg_propgraph_element_label.md#catalog-pg-propgraph-element-label"><code>pg_propgraph_element_label</code></a>.<code>oid</code>)</p>
<p>Reference to the label (indirectly via <code>pg_propgraph_element_label</code>, which then links to <code>pg_propgraph_label</code>)</p></td>
</tr>
<tr>
<td><p><code>plpexpr</code> <code>pg_node_tree</code></p>
<p>Expression tree (in <code>nodeToString()</code> representation) for the property's definition. The expression references the table reached via <code>pg_propgraph_element_label</code> and <code>pg_propgraph_element</code>.</p></td>
</tr>
</tbody>
</table>
