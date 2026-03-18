<a id="catalog-pg-propgraph-label"></a>

## `pg_propgraph_label`


 The catalog `pg_propgraph_label` stores information about the labels in a property graph.


**Table: `pg_propgraph_label` Columns**

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
<td><p><code>pglpgid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Reference to the property graph that this label belongs to</p></td>
</tr>
<tr>
<td><p><code>pgllabel</code> <code>name</code></p>
<p>The name of the label. This is unique among the labels in a graph.</p></td>
</tr>
</tbody>
</table>
