<a id="catalog-pg-propgraph-element-label"></a>

## `pg_propgraph_element_label`


 The catalog `pg_propgraph_element_label` stores information about which labels apply to which elements.


**Table: `pg_propgraph_element_label` Columns**

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
<td><p><code>pgellabelid</code> <code>oid</code> (references <a href="pg_propgraph_label.md#catalog-pg-propgraph-label"><code>pg_propgraph_label</code></a>.<code>oid</code>)</p>
<p>Reference to the label</p></td>
</tr>
<tr>
<td><p><code>pgelelid</code> <code>oid</code> (references <a href="pg_propgraph_element.md#catalog-pg-propgraph-element"><code>pg_propgraph_element</code></a>.<code>oid</code>)</p>
<p>Reference to the element</p></td>
</tr>
</tbody>
</table>
