<a id="catalog-pg-attrdef"></a>

## `pg_attrdef`


 The catalog `pg_attrdef` stores column default expressions and generation expressions. The main information about columns is stored in [`pg_attribute`](pg_attribute.md#catalog-pg-attribute). Only columns for which a default expression or generation expression has been explicitly set will have an entry here.


**Table: `pg_attrdef` Columns**

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
<td><p><code>adrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table this column belongs to</p></td>
</tr>
<tr>
<td><p><code>adnum</code> <code>int2</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>The number of the column</p></td>
</tr>
<tr>
<td><p><code>adbin</code> <code>pg_node_tree</code></p>
<p>The column default or generation expression, in <code>nodeToString()</code> representation. Use <code>pg_get_expr(adbin, adrelid)</code> to convert it to an SQL expression.</p></td>
</tr>
</tbody>
</table>
