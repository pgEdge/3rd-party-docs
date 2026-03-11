## Built-in Operator Classes { #gin-builtin-opclasses }


 The core PostgreSQL distribution includes the GIN operator classes shown in [Built-in GIN Operator Classes](#gin-builtin-opclasses-table). (Some of the optional modules described in [Additional Supplied Modules and Extensions](../../appendixes/additional-supplied-modules-and-extensions/index.md#contrib) provide additional GIN operator classes.)


<a id="gin-builtin-opclasses-table"></a>
**Table: Built-in GIN Operator Classes**

<table>
<thead>
<tr>
<th>Name</th>
<th>Indexable Operators</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan="3"><code>array_ops</code></td>
<td><code>&amp;&amp; (anyarray,anyarray)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyarray,anyarray)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anyarray,anyarray)</code></td>
</tr>
<tr>
<td><code>= (anyarray,anyarray)</code></td>
</tr>
<tr>
<td rowspan="5"><code>jsonb_ops</code></td>
<td><code>@&gt; (jsonb,jsonb)</code></td>
</tr>
<tr>
<td><code>@? (jsonb,jsonpath)</code></td>
</tr>
<tr>
<td><code>@@ (jsonb,jsonpath)</code></td>
</tr>
<tr>
<td><code>? (jsonb,text)</code></td>
</tr>
<tr>
<td><code>?| (jsonb,text[])</code></td>
</tr>
<tr>
<td><code>?&amp; (jsonb,text[])</code></td>
</tr>
<tr>
<td rowspan="2"><code>jsonb_path_ops</code></td>
<td><code>@&gt; (jsonb,jsonb)</code></td>
</tr>
<tr>
<td><code>@? (jsonb,jsonpath)</code></td>
</tr>
<tr>
<td><code>@@ (jsonb,jsonpath)</code></td>
</tr>
<tr>
<td rowspan="1"><code>tsvector_ops</code></td>
<td><code>@@ (tsvector,tsquery)</code></td>
</tr>
<tr>
<td><code>@@@ (tsvector,tsquery)</code></td>
</tr>
</tbody>
</table>


 Of the two operator classes for type `jsonb`, `jsonb_ops` is the default. `jsonb_path_ops` supports fewer operators but offers better performance for those operators. See [`jsonb` Indexing](../../the-sql-language/data-types/json-types.md#json-indexing) for details.
