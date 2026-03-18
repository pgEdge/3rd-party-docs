<a id="catalog-pg-propgraph-property"></a>

## `pg_propgraph_property`


 The catalog `pg_propgraph_property` stores information about the properties in a property graph. This only stores information that applies to a property throughout the graph, independent of what label or element it is on. Additional information, including the actual expressions that define the properties are in the catalog [`pg_propgraph_label_property`](pg_propgraph_label_property.md#catalog-pg-propgraph-label-property).


**Table: `pg_propgraph_property` Columns**

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
<td><p><code>pgppgid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Reference to the property graph that this property belongs to</p></td>
</tr>
<tr>
<td><p><code>pgpname</code> <code>name</code></p>
<p>The name of the property. This is unique among the properties in a graph.</p></td>
</tr>
<tr>
<td><p><code>pgptypid</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>The data type of this property. (This is required to be fixed for a given property in a property graph, even if the property is defined multiple times in different elements and labels.)</p></td>
</tr>
<tr>
<td><p><code>pgptypmod</code> <code>int4</code></p>
<p><code>typmod</code> to be applied to the data type of this property. (This is required to be fixed for a given property in a property graph, even if the property is defined multiple times in different elements and labels.)</p></td>
</tr>
<tr>
<td><p><code>pgpcollation</code> <code>oid</code> (references <a href="pg_collation.md#catalog-pg-collation"><code>pg_collation</code></a>.<code>oid</code>)</p>
<p>The defined collation of this property, or zero if the property is not of a collatable data type. (This is required to be fixed for a given property in a property graph, even if the property is defined multiple times in different elements and labels.)</p></td>
</tr>
</tbody>
</table>
