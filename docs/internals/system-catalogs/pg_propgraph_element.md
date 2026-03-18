<a id="catalog-pg-propgraph-element"></a>

## `pg_propgraph_element`


 The catalog `pg_propgraph_element` stores information about the vertices and edges of a property graph, collectively called the elements of the property graph.


**Table: `pg_propgraph_element` Columns**

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
<td><p><code>pgepgid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Reference to the property graph that this element belongs to</p></td>
</tr>
<tr>
<td><p><code>pgerelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Reference to the table that contains the data for this property graph element</p></td>
</tr>
<tr>
<td><p><code>pgealias</code> <code>name</code></p>
<p>The alias of the element. This is a unique identifier for the element within the graph. It is set when the property graph is defined and defaults to the name of the underlying element table.</p></td>
</tr>
<tr>
<td><p><code>pgekind</code> <code>char</code></p>
<p><code>v</code> for a vertex, <code>e</code> for an edge</p></td>
</tr>
<tr>
<td><p><code>pgesrcvertexid</code> <code>oid</code> (references <a href="#catalog-pg-propgraph-element"><code>pg_propgraph_element</code></a>.<code>oid</code>)</p>
<p>For an edge, a link to the source vertex. (Zero for a vertex.)</p></td>
</tr>
<tr>
<td><p><code>pgedestvertexid</code> <code>oid</code> (references <a href="#catalog-pg-propgraph-element"><code>pg_propgraph_element</code></a>.<code>oid</code>)</p>
<p>For an edge, a link to the destination vertex. (Zero for a vertex.)</p></td>
</tr>
<tr>
<td><p><code>pgekey</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>An array of column numbers in the table referenced by <code>pgerelid</code> that defines the key to use for this element table. (This defaults to the primary key when the property graph is created.)</p></td>
</tr>
<tr>
<td><p><code>pgesrckey</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>For an edge, an array of column numbers in the table referenced by <code>pgerelid</code> that defines the source key to use for this element table. (Null for a vertex.) The combination of <code>pgesrckey</code> and <code>pgesrcref</code> creates the link between the edge and the source vertex.</p></td>
</tr>
<tr>
<td><p><code>pgesrcref</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>For an edge, an array of column numbers in the table reached via <code>pgesrcvertexid</code>. (Null for a vertex.) The combination of <code>pgesrckey</code> and <code>pgesrcref</code> creates the link between the edge and the source vertex.</p></td>
</tr>
<tr>
<td><p><code>pgesrceqop</code> <code>oid[]</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>For an edge, an array of equality operators for <code>pgesrcref</code> = <code>pgesrckey</code> comparison. (Null for a vertex.)</p></td>
</tr>
<tr>
<td><p><code>pgedestkey</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>For an edge, an array of column numbers in the table referenced by <code>pgerelid</code> that defines the destination key to use for this element table. (Null for a vertex.) The combination of <code>pgedestkey</code> and <code>pgedestref</code> creates the link between the edge and the destination vertex.</p></td>
</tr>
<tr>
<td><p><code>pgedestref</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>For an edge, an array of column numbers in the table reached via <code>pgedestvertexid</code>. (Null for a vertex.) The combination of <code>pgedestkey</code> and <code>pgedestref</code> creates the link between the edge and the destination vertex.</p></td>
</tr>
<tr>
<td><p><code>pgedesteqop</code> <code>oid[]</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>For an edge, an array of equality operators for <code>pgedestref</code> = <code>pgedestkey</code> comparison. (Null for a vertex.)</p></td>
</tr>
</tbody>
</table>
