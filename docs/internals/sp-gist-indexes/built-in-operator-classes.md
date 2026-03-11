## Built-in Operator Classes { #spgist-builtin-opclasses }


 The core PostgreSQL distribution includes the SP-GiST operator classes shown in [Built-in SP-GiST Operator Classes](#spgist-builtin-opclasses-table).


<a id="spgist-builtin-opclasses-table"></a>
**Table: Built-in SP-GiST Operator Classes**

<table>
<thead>
<tr>
<th>Name</th>
<th>Indexable Operators</th>
<th>Ordering Operators</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan="11"><code>box_ops</code></td>
<td><code>&lt;&lt; (box,box)</code></td>
<td rowspan="11"><code>&lt;-&gt; (box,point)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>&lt;@ (box,box)</code></td>
</tr>
<tr>
<td><code>@&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>~= (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (box,box)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&lt;| (box,box)</code></td>
</tr>
<tr>
<td><code>|&amp;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>|&gt;&gt; (box,box)</code></td>
</tr>
<tr>
<td rowspan="10"><code>inet_ops</code></td>
<td><code>&lt;&lt; (inet,inet)</code></td>
<td rowspan="10"></td>
</tr>
<tr>
<td><code>&lt;&lt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;&gt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt;&gt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (inet,inet)</code></td>
</tr>
<tr>
<td rowspan="5"><code>kd_point_ops</code></td>
<td><code>|&gt;&gt; (point,point)</code></td>
<td rowspan="5"><code>&lt;-&gt; (point,point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (point,point)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (point,point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (point,point)</code></td>
</tr>
<tr>
<td><code>~= (point,point)</code></td>
</tr>
<tr>
<td><code>&lt;@ (point,box)</code></td>
</tr>
<tr>
<td rowspan="11"><code>poly_ops</code></td>
<td><code>&lt;&lt; (polygon,polygon)</code></td>
<td rowspan="11"><code>&lt;-&gt; (polygon,point)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>&lt;@ (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>@&gt; (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>~= (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>&amp;&lt;| (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>|&gt;&gt; (polygon,polygon)</code></td>
</tr>
<tr>
<td><code>|&amp;&gt; (polygon,polygon)</code></td>
</tr>
<tr>
<td rowspan="5"><code>quad_point_ops</code></td>
<td><code>|&gt;&gt; (point,point)</code></td>
<td rowspan="5"><code>&lt;-&gt; (point,point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (point,point)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (point,point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (point,point)</code></td>
</tr>
<tr>
<td><code>~= (point,point)</code></td>
</tr>
<tr>
<td><code>&lt;@ (point,box)</code></td>
</tr>
<tr>
<td rowspan="9"><code>range_ops</code></td>
<td><code>= (anyrange,anyrange)</code></td>
<td rowspan="9"></td>
</tr>
<tr>
<td><code>&amp;&amp; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange,anyelement)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>-|- (anyrange,anyrange)</code></td>
</tr>
<tr>
<td rowspan="9"><code>text_ops</code></td>
<td><code>= (text,text)</code></td>
<td rowspan="9"></td>
</tr>
<tr>
<td><code>&lt; (text,text)</code></td>
</tr>
<tr>
<td><code>&lt;= (text,text)</code></td>
</tr>
<tr>
<td><code>&gt; (text,text)</code></td>
</tr>
<tr>
<td><code>&gt;= (text,text)</code></td>
</tr>
<tr>
<td><code>~&lt;~ (text,text)</code></td>
</tr>
<tr>
<td><code>~&lt;=~ (text,text)</code></td>
</tr>
<tr>
<td><code>~&gt;=~ (text,text)</code></td>
</tr>
<tr>
<td><code>~&gt;~ (text,text)</code></td>
</tr>
<tr>
<td><code>^@ (text,text)</code></td>
</tr>
</tbody>
</table>


 Of the two operator classes for type `point`, `quad_point_ops` is the default. `kd_point_ops` supports the same operators but uses a different index data structure that may offer better performance in some applications.


 The `quad_point_ops`, `kd_point_ops` and `poly_ops` operator classes support the `<->` ordering operator, which enables the k-nearest neighbor (`k-NN`) search over indexed point or polygon data sets.
