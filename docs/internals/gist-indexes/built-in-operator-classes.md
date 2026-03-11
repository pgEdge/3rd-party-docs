## Built-in Operator Classes { #gist-builtin-opclasses }


 The core PostgreSQL distribution includes the GiST operator classes shown in [Built-in GiST Operator Classes](#gist-builtin-opclasses-table). (Some of the optional modules described in [Additional Supplied Modules and Extensions](../../appendixes/additional-supplied-modules-and-extensions/index.md#contrib) provide additional GiST operator classes.)


<a id="gist-builtin-opclasses-table"></a>
**Table: Built-in GiST Operator Classes**

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
<td><code>&lt;&lt; (box, box)</code></td>
<td rowspan="11"><code>&lt;-&gt; (box, point)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (box, box)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (box, box)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (box, box)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (box, box)</code></td>
</tr>
<tr>
<td><code>~= (box, box)</code></td>
</tr>
<tr>
<td><code>@&gt; (box, box)</code></td>
</tr>
<tr>
<td><code>&lt;@ (box, box)</code></td>
</tr>
<tr>
<td><code>&amp;&lt;| (box, box)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (box, box)</code></td>
</tr>
<tr>
<td><code>|&gt;&gt; (box, box)</code></td>
</tr>
<tr>
<td><code>|&amp;&gt; (box, box)</code></td>
</tr>
<tr>
<td rowspan="11"><code>circle_ops</code></td>
<td><code>&lt;&lt; (circle, circle)</code></td>
<td rowspan="11"><code>&lt;-&gt; (circle, point)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (circle, circle)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (circle, circle)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (circle, circle)</code></td>
</tr>
<tr>
<td><code>&lt;@ (circle, circle)</code></td>
</tr>
<tr>
<td><code>@&gt; (circle, circle)</code></td>
</tr>
<tr>
<td><code>~= (circle, circle)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (circle, circle)</code></td>
</tr>
<tr>
<td><code>|&gt;&gt; (circle, circle)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (circle, circle)</code></td>
</tr>
<tr>
<td><code>&amp;&lt;| (circle, circle)</code></td>
</tr>
<tr>
<td><code>|&amp;&gt; (circle, circle)</code></td>
</tr>
<tr>
<td rowspan="10"><code>inet_ops</code></td>
<td><code>&lt;&lt; (inet, inet)</code></td>
<td rowspan="10"></td>
</tr>
<tr>
<td><code>&lt;&lt;= (inet, inet)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (inet, inet)</code></td>
</tr>
<tr>
<td><code>&gt;&gt;= (inet, inet)</code></td>
</tr>
<tr>
<td><code>= (inet, inet)</code></td>
</tr>
<tr>
<td><code>&lt;&gt; (inet, inet)</code></td>
</tr>
<tr>
<td><code>&lt; (inet, inet)</code></td>
</tr>
<tr>
<td><code>&lt;= (inet, inet)</code></td>
</tr>
<tr>
<td><code>&gt; (inet, inet)</code></td>
</tr>
<tr>
<td><code>&gt;= (inet, inet)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (inet, inet)</code></td>
</tr>
<tr>
<td rowspan="17"><code>multirange_ops</code></td>
<td><code>= (anymultirange, anymultirange)</code></td>
<td rowspan="17"></td>
</tr>
<tr>
<td><code>&amp;&amp; (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>@&gt; (anymultirange, anyelement)</code></td>
</tr>
<tr>
<td><code>@&gt; (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>@&gt; (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td><code>-|- (anymultirange, anymultirange)</code></td>
</tr>
<tr>
<td><code>-|- (anymultirange, anyrange)</code></td>
</tr>
<tr>
<td rowspan="7"><code>point_ops</code></td>
<td><code>|&gt;&gt; (point, point)</code></td>
<td rowspan="7"><code>&lt;-&gt; (point, point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (point, point)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (point, point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (point, point)</code></td>
</tr>
<tr>
<td><code>~= (point, point)</code></td>
</tr>
<tr>
<td><code>&lt;@ (point, box)</code></td>
</tr>
<tr>
<td><code>&lt;@ (point, polygon)</code></td>
</tr>
<tr>
<td><code>&lt;@ (point, circle)</code></td>
</tr>
<tr>
<td rowspan="11"><code>poly_ops</code></td>
<td><code>&lt;&lt; (polygon, polygon)</code></td>
<td rowspan="11"><code>&lt;-&gt; (polygon, point)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>&lt;@ (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>@&gt; (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>~= (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>&amp;&lt;| (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>|&amp;&gt; (polygon, polygon)</code></td>
</tr>
<tr>
<td><code>|&gt;&gt; (polygon, polygon)</code></td>
</tr>
<tr>
<td rowspan="17"><code>range_ops</code></td>
<td><code>= (anyrange, anyrange)</code></td>
<td rowspan="17"></td>
</tr>
<tr>
<td><code>&amp;&amp; (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange, anyelement)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td><code>-|- (anyrange, anyrange)</code></td>
</tr>
<tr>
<td><code>-|- (anyrange, anymultirange)</code></td>
</tr>
<tr>
<td rowspan="1"><code>tsquery_ops</code></td>
<td><code>&lt;@ (tsquery, tsquery)</code></td>
<td rowspan="1"></td>
</tr>
<tr>
<td><code>@&gt; (tsquery, tsquery)</code></td>
</tr>
<tr>
<td><code>tsvector_ops</code></td>
<td><code>@@ (tsvector, tsquery)</code></td>
<td></td>
</tr>
</tbody>
</table>


 For historical reasons, the `inet_ops` operator class is not the default class for types `inet` and `cidr`. To use it, mention the class name in `CREATE INDEX`, for example

```sql

CREATE INDEX ON my_table USING GIST (my_inet_column inet_ops);
```
