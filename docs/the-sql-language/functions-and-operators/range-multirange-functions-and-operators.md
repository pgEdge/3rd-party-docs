<a id="functions-range"></a>

## Range/Multirange Functions and Operators


 See [Range Types](../data-types/range-types.md#rangetypes) for an overview of range types.


 [Range Operators](#range-operators-table) shows the specialized operators available for range types. [Multirange Operators](#multirange-operators-table) shows the specialized operators available for multirange types. In addition to those, the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) are available for range and multirange types. The comparison operators order first by the range lower bounds, and only if those are equal do they compare the upper bounds. The multirange operators compare each range until one is unequal. This does not usually result in a useful overall ordering, but the operators are provided to allow unique indexes to be constructed on ranges.
 <a id="range-operators-table"></a>

**Table: Range Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>anyrange</code> <code>@&gt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the first range contain the second?</td>
<td><code>int4range(2,4) @&gt; int4range(2,3)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>@&gt;</code> <code>anyelement</code> <code>boolean</code></td>
<td>Does the range contain the element?</td>
<td><code>'[2011-01-01,2011-03-01)'::tsrange @&gt; '2011-01-10'::timestamp</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&lt;@</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the first range contained by the second?</td>
<td><code>int4range(2,4) &lt;@ int4range(1,7)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyelement</code> <code>&lt;@</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the element contained in the range?</td>
<td><code>42 &lt;@ int4range(1,7)</code> <code>f</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&amp;&amp;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Do the ranges overlap, that is, have any elements in common?</td>
<td><code>int8range(3,7) &amp;&amp; int8range(4,12)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&lt;&lt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the first range strictly left of the second?</td>
<td><code>int8range(1,10) &lt;&lt; int8range(100,110)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&gt;&gt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the first range strictly right of the second?</td>
<td><code>int8range(50,60) &gt;&gt; int8range(20,30)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&amp;&lt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the first range not extend to the right of the second?</td>
<td><code>int8range(1,20) &amp;&lt; int8range(18,20)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&amp;&gt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the first range not extend to the left of the second?</td>
<td><code>int8range(7,20) &amp;&gt; int8range(5,10)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>-|-</code> <code>anyrange</code> <code>boolean</code></td>
<td>Are the ranges adjacent?</td>
<td><code>numrange(1.1,2.2) -|- numrange(2.2,3.3)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>+</code> <code>anyrange</code> <code>anyrange</code></td>
<td>Computes the union of the ranges. The ranges must overlap or be adjacent, so that the union is a single range (but see <code>range_merge()</code>).</td>
<td><code>numrange(5,15) + numrange(10,20)</code> <code>[5,20)</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>*</code> <code>anyrange</code> <code>anyrange</code></td>
<td>Computes the intersection of the ranges.</td>
<td><code>int8range(5,15) * int8range(10,20)</code> <code>[10,15)</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>-</code> <code>anyrange</code> <code>anyrange</code></td>
<td>Computes the difference of the ranges. The second range must not be contained in the first in such a way that the difference would not be a single range.</td>
<td><code>int8range(5,15) - int8range(10,20)</code> <code>[5,10)</code></td>
</tr>
</tbody>
</table>
 <a id="multirange-operators-table"></a>

**Table: Multirange Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>anymultirange</code> <code>@&gt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the first multirange contain the second?</td>
<td><code>'{[2,4)}'::int4multirange @&gt; '{[2,3)}'::int4multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>@&gt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the multirange contain the range?</td>
<td><code>'{[2,4)}'::int4multirange @&gt; int4range(2,3)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>@&gt;</code> <code>anyelement</code> <code>boolean</code></td>
<td>Does the multirange contain the element?</td>
<td><code>'{[2011-01-01,2011-03-01)}'::tsmultirange @&gt; '2011-01-10'::timestamp</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>@&gt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the range contain the multirange?</td>
<td><code>'[2,4)'::int4range @&gt; '{[2,3)}'::int4multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&lt;@</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the first multirange contained by the second?</td>
<td><code>'{[2,4)}'::int4multirange &lt;@ '{[1,7)}'::int4multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&lt;@</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the multirange contained by the range?</td>
<td><code>'{[2,4)}'::int4multirange &lt;@ int4range(1,7)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&lt;@</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the range contained by the multirange?</td>
<td><code>int4range(2,4) &lt;@ '{[1,7)}'::int4multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyelement</code> <code>&lt;@</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the element contained by the multirange?</td>
<td><code>4 &lt;@ '{[1,7)}'::int4multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&amp;&amp;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Do the multiranges overlap, that is, have any elements in common?</td>
<td><code>'{[3,7)}'::int8multirange &amp;&amp; '{[4,12)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&amp;&amp;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the multirange overlap the range?</td>
<td><code>'{[3,7)}'::int8multirange &amp;&amp; int8range(4,12)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&amp;&amp;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the range overlap the multirange?</td>
<td><code>int8range(3,7) &amp;&amp; '{[4,12)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&lt;&lt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the first multirange strictly left of the second?</td>
<td><code>'{[1,10)}'::int8multirange &lt;&lt; '{[100,110)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&lt;&lt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the multirange strictly left of the range?</td>
<td><code>'{[1,10)}'::int8multirange &lt;&lt; int8range(100,110)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&lt;&lt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the range strictly left of the multirange?</td>
<td><code>int8range(1,10) &lt;&lt; '{[100,110)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&gt;&gt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the first multirange strictly right of the second?</td>
<td><code>'{[50,60)}'::int8multirange &gt;&gt; '{[20,30)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&gt;&gt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the multirange strictly right of the range?</td>
<td><code>'{[50,60)}'::int8multirange &gt;&gt; int8range(20,30)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&gt;&gt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the range strictly right of the multirange?</td>
<td><code>int8range(50,60) &gt;&gt; '{[20,30)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&amp;&lt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the first multirange not extend to the right of the second?</td>
<td><code>'{[1,20)}'::int8multirange &amp;&lt; '{[18,20)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&amp;&lt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the multirange not extend to the right of the range?</td>
<td><code>'{[1,20)}'::int8multirange &amp;&lt; int8range(18,20)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&amp;&lt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the range not extend to the right of the multirange?</td>
<td><code>int8range(1,20) &amp;&lt; '{[18,20)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&amp;&gt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the first multirange not extend to the left of the second?</td>
<td><code>'{[7,20)}'::int8multirange &amp;&gt; '{[5,10)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>&amp;&gt;</code> <code>anyrange</code> <code>boolean</code></td>
<td>Does the multirange not extend to the left of the range?</td>
<td><code>'{[7,20)}'::int8multirange &amp;&gt; int8range(5,10)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>&amp;&gt;</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Does the range not extend to the left of the multirange?</td>
<td><code>int8range(7,20) &amp;&gt; '{[5,10)}'::int8multirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>-|-</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Are the multiranges adjacent?</td>
<td><code>'{[1.1,2.2)}'::nummultirange -|- '{[2.2,3.3)}'::nummultirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>-|-</code> <code>anyrange</code> <code>boolean</code></td>
<td>Is the multirange adjacent to the range?</td>
<td><code>'{[1.1,2.2)}'::nummultirange -|- numrange(2.2,3.3)</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyrange</code> <code>-|-</code> <code>anymultirange</code> <code>boolean</code></td>
<td>Is the range adjacent to the multirange?</td>
<td><code>numrange(1.1,2.2) -|- '{[2.2,3.3)}'::nummultirange</code> <code>t</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>+</code> <code>anymultirange</code> <code>anymultirange</code></td>
<td>Computes the union of the multiranges. The multiranges need not overlap or be adjacent.</td>
<td><code>'{[5,10)}'::nummultirange + '{[15,20)}'::nummultirange</code> <code>{[5,10), [15,20)}</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>*</code> <code>anymultirange</code> <code>anymultirange</code></td>
<td>Computes the intersection of the multiranges.</td>
<td><code>'{[5,15)}'::int8multirange * '{[10,20)}'::int8multirange</code> <code>{[10,15)}</code></td>
</tr>
<tr>
<td><code>anymultirange</code> <code>-</code> <code>anymultirange</code> <code>anymultirange</code></td>
<td>Computes the difference of the multiranges.</td>
<td><code>'{[5,20)}'::int8multirange - '{[10,15)}'::int8multirange</code> <code>{[5,10), [15,20)}</code></td>
</tr>
</tbody>
</table>


 The left-of/right-of/adjacent operators always return false when an empty range or multirange is involved; that is, an empty range is not considered to be either before or after any other range.


 Elsewhere empty ranges and multiranges are treated as the additive identity: anything unioned with an empty value is itself. Anything minus an empty value is itself. An empty multirange has exactly the same points as an empty range. Every range contains the empty range. Every multirange contains as many empty ranges as you like.


 The range union and difference operators will fail if the resulting range would need to contain two disjoint sub-ranges, as such a range cannot be represented. There are separate operators for union and difference that take multirange parameters and return a multirange, and they do not fail even if their arguments are disjoint. So if you need a union or difference operation for ranges that may be disjoint, you can avoid errors by first casting your ranges to multiranges.


 [Range Functions](#range-functions-table) shows the functions available for use with range types. [Multirange Functions](#multirange-functions-table) shows the functions available for use with multirange types.
 <a id="range-functions-table"></a>

**Table: Range Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>lower</code> ( <code>anyrange</code> ) <code>anyelement</code></td>
<td>Extracts the lower bound of the range (<code>NULL</code> if the range is empty or has no lower bound).</td>
<td><code>lower(numrange(1.1,2.2))</code> <code>1.1</code></td>
</tr>
<tr>
<td><code>upper</code> ( <code>anyrange</code> ) <code>anyelement</code></td>
<td>Extracts the upper bound of the range (<code>NULL</code> if the range is empty or has no upper bound).</td>
<td><code>upper(numrange(1.1,2.2))</code> <code>2.2</code></td>
</tr>
<tr>
<td><code>isempty</code> ( <code>anyrange</code> ) <code>boolean</code></td>
<td>Is the range empty?</td>
<td><code>isempty(numrange(1.1,2.2))</code> <code>f</code></td>
</tr>
<tr>
<td><code>lower_inc</code> ( <code>anyrange</code> ) <code>boolean</code></td>
<td>Is the range's lower bound inclusive?</td>
<td><code>lower_inc(numrange(1.1,2.2))</code> <code>t</code></td>
</tr>
<tr>
<td><code>upper_inc</code> ( <code>anyrange</code> ) <code>boolean</code></td>
<td>Is the range's upper bound inclusive?</td>
<td><code>upper_inc(numrange(1.1,2.2))</code> <code>f</code></td>
</tr>
<tr>
<td><code>lower_inf</code> ( <code>anyrange</code> ) <code>boolean</code></td>
<td>Does the range have no lower bound? (A lower bound of <code>-Infinity</code> returns false.)</td>
<td><code>lower_inf('(,)'::daterange)</code> <code>t</code></td>
</tr>
<tr>
<td><code>upper_inf</code> ( <code>anyrange</code> ) <code>boolean</code></td>
<td>Does the range have no upper bound? (An upper bound of <code>Infinity</code> returns false.)</td>
<td><code>upper_inf('(,)'::daterange)</code> <code>t</code></td>
</tr>
<tr>
<td><code>range_merge</code> ( <code>anyrange</code>, <code>anyrange</code> ) <code>anyrange</code></td>
<td>Computes the smallest range that includes both of the given ranges.</td>
<td><code>range_merge('[1,2)'::int4range, '[3,4)'::int4range)</code> <code>[1,4)</code></td>
</tr>
<tr>
<td><code>range_minus_multi</code> ( <code>anyrange</code>, <code>anyrange</code> ) <code>setof anyrange</code></td>
<td>Returns the non-empty range(s) remaining after subtracting the second range from the first. One row is returned for each range, so if the second range splits the first into two parts, there will be two results. If the subtraction yields an empty range, no rows are returned.</td>
<td><p><code>range_minus_multi('[0,10)'::int4range, '[3,4)'::int4range)</code></p>
<pre><code>
 [0,3)
 [4,10)</code></pre></td>
</tr>
</tbody>
</table>
 <a id="multirange-functions-table"></a>

**Table: Multirange Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>lower</code> ( <code>anymultirange</code> ) <code>anyelement</code></td>
<td>Extracts the lower bound of the multirange (<code>NULL</code> if the multirange is empty or has no lower bound).</td>
<td><code>lower('{[1.1,2.2)}'::nummultirange)</code> <code>1.1</code></td>
</tr>
<tr>
<td><code>upper</code> ( <code>anymultirange</code> ) <code>anyelement</code></td>
<td>Extracts the upper bound of the multirange (<code>NULL</code> if the multirange is empty or has no upper bound).</td>
<td><code>upper('{[1.1,2.2)}'::nummultirange)</code> <code>2.2</code></td>
</tr>
<tr>
<td><code>isempty</code> ( <code>anymultirange</code> ) <code>boolean</code></td>
<td>Is the multirange empty?</td>
<td><code>isempty('{[1.1,2.2)}'::nummultirange)</code> <code>f</code></td>
</tr>
<tr>
<td><code>lower_inc</code> ( <code>anymultirange</code> ) <code>boolean</code></td>
<td>Is the multirange's lower bound inclusive?</td>
<td><code>lower_inc('{[1.1,2.2)}'::nummultirange)</code> <code>t</code></td>
</tr>
<tr>
<td><code>upper_inc</code> ( <code>anymultirange</code> ) <code>boolean</code></td>
<td>Is the multirange's upper bound inclusive?</td>
<td><code>upper_inc('{[1.1,2.2)}'::nummultirange)</code> <code>f</code></td>
</tr>
<tr>
<td><code>lower_inf</code> ( <code>anymultirange</code> ) <code>boolean</code></td>
<td>Does the multirange have no lower bound? (A lower bound of <code>-Infinity</code> returns false.)</td>
<td><code>lower_inf('{(,)}'::datemultirange)</code> <code>t</code></td>
</tr>
<tr>
<td><code>upper_inf</code> ( <code>anymultirange</code> ) <code>boolean</code></td>
<td>Does the multirange have no upper bound? (An upper bound of <code>Infinity</code> returns false.)</td>
<td><code>upper_inf('{(,)}'::datemultirange)</code> <code>t</code></td>
</tr>
<tr>
<td><code>range_merge</code> ( <code>anymultirange</code> ) <code>anyrange</code></td>
<td>Computes the smallest range that includes the entire multirange.</td>
<td><code>range_merge('{[1,2), [3,4)}'::int4multirange)</code> <code>[1,4)</code></td>
</tr>
<tr>
<td><code>multirange</code> ( <code>anyrange</code> ) <code>anymultirange</code></td>
<td>Returns a multirange containing just the given range.</td>
<td><code>multirange('[1,2)'::int4range)</code> <code>{[1,2)}</code></td>
</tr>
<tr>
<td><code>unnest</code> ( <code>anymultirange</code> ) <code>setof anyrange</code></td>
<td>Expands a multirange into a set of ranges in ascending order.</td>
<td><p><code>unnest('{[1,2), [3,4)}'::int4multirange)</code></p>
<pre><code>
 [1,2)
 [3,4)</code></pre></td>
</tr>
<tr>
<td><code>multirange_minus_multi</code> ( <code>anymultirange</code>, <code>anymultirange</code> ) <code>setof anymultirange</code></td>
<td>Returns the non-empty multirange(s) remaining after subtracting the second multirange from the first. If the subtraction yields an empty multirange, no rows are returned. Two rows are never returned, because a single multirange can always accommodate any result.</td>
<td><code>multirange_minus_multi('{[0,10)}'::int4multirange, '{[3,4)}'::int4multirange)</code> <code>{[0,3), [4,10)}</code></td>
</tr>
</tbody>
</table>


 The `lower_inc`, `upper_inc`, `lower_inf`, and `upper_inf` functions all return false for an empty range or multirange.
