<a id="functions-array"></a>

## Array Functions and Operators


 [Array Operators](#array-operators-table) shows the specialized operators available for array types. In addition to those, the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) are available for arrays. The comparison operators compare the array contents element-by-element, using the default B-tree comparison function for the element data type, and sort based on the first difference. In multidimensional arrays the elements are visited in row-major order (last subscript varies most rapidly). If the contents of two arrays are equal but the dimensionality is different, the first difference in the dimensionality information determines the sort order.
 <a id="array-operators-table"></a>

**Table: Array Operators**

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
<td><code>anyarray</code> <code>@&gt;</code> <code>anyarray</code> <code>boolean</code></td>
<td>Does the first array contain the second, that is, does each element appearing in the second array equal some element of the first array? (Duplicates are not treated specially, thus <code>ARRAY[1]</code> and <code>ARRAY[1,1]</code> are each considered to contain the other.)</td>
<td><code>ARRAY[1,4,3] @&gt; ARRAY[3,1,3]</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyarray</code> <code>&lt;@</code> <code>anyarray</code> <code>boolean</code></td>
<td>Is the first array contained by the second?</td>
<td><code>ARRAY[2,2,7] &lt;@ ARRAY[1,7,4,2,6]</code> <code>t</code></td>
</tr>
<tr>
<td><code>anyarray</code> <code>&amp;&amp;</code> <code>anyarray</code> <code>boolean</code></td>
<td>Do the arrays overlap, that is, have any elements in common?</td>
<td><code>ARRAY[1,4,3] &amp;&amp; ARRAY[2,1]</code> <code>t</code></td>
</tr>
<tr>
<td><code>anycompatiblearray</code> <code>||</code> <code>anycompatiblearray</code> <code>anycompatiblearray</code></td>
<td>Concatenates the two arrays. Concatenating a null or empty array is a no-op; otherwise the arrays must have the same number of dimensions (as illustrated by the first example) or differ in number of dimensions by one (as illustrated by the second). If the arrays are not of identical element types, they will be coerced to a common type (see <a href="../type-conversion/union-case-and-related-constructs.md#typeconv-union-case"><code>UNION</code>, <code>CASE</code>, and Related Constructs</a>).</td>
<td><code>ARRAY[1,2,3] || ARRAY[4,5,6,7]</code> <code>{1,2,3,4,5,6,7}</code><br><code>ARRAY[1,2,3] || ARRAY[[4,5,6],[7,8,9.9]]</code> <code>{{1,2,3},{4,5,6},{7,8,9.9}}</code></td>
</tr>
<tr>
<td><code>anycompatible</code> <code>||</code> <code>anycompatiblearray</code> <code>anycompatiblearray</code></td>
<td>Concatenates an element onto the front of an array (which must be empty or one-dimensional).</td>
<td><code>3 || ARRAY[4,5,6]</code> <code>{3,4,5,6}</code></td>
</tr>
<tr>
<td><code>anycompatiblearray</code> <code>||</code> <code>anycompatible</code> <code>anycompatiblearray</code></td>
<td>Concatenates an element onto the end of an array (which must be empty or one-dimensional).</td>
<td><code>ARRAY[4,5,6] || 7</code> <code>{4,5,6,7}</code></td>
</tr>
</tbody>
</table>


 See [Arrays](../data-types/arrays.md#arrays) for more details about array operator behavior. See [Index Types](../indexes/index-types.md#indexes-types) for more details about which operators support indexed operations.


 [Array Functions](#array-functions-table) shows the functions available for use with array types. See [Arrays](../data-types/arrays.md#arrays) for more information and examples of the use of these functions.
 <a id="array-functions-table"></a>

**Table: Array Functions**

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
<td><code>array_append</code> ( <code>anycompatiblearray</code>, <code>anycompatible</code> ) <code>anycompatiblearray</code></td>
<td>Appends an element to the end of an array (same as the <code>anycompatiblearray</code> <code>||</code> <code>anycompatible</code> operator).</td>
<td><code>array_append(ARRAY[1,2], 3)</code> <code>{1,2,3}</code></td>
</tr>
<tr>
<td><code>array_cat</code> ( <code>anycompatiblearray</code>, <code>anycompatiblearray</code> ) <code>anycompatiblearray</code></td>
<td>Concatenates two arrays (same as the <code>anycompatiblearray</code> <code>||</code> <code>anycompatiblearray</code> operator).</td>
<td><code>array_cat(ARRAY[1,2,3], ARRAY[4,5])</code> <code>{1,2,3,4,5}</code></td>
</tr>
<tr>
<td><code>array_dims</code> ( <code>anyarray</code> ) <code>text</code></td>
<td>Returns a text representation of the array's dimensions.</td>
<td><code>array_dims(ARRAY[[1,2,3], [4,5,6]])</code> <code>[1:2][1:3]</code></td>
</tr>
<tr>
<td><code>array_fill</code> ( <code>anyelement</code>, <code>integer[]</code> [, <code>integer[]</code> ] ) <code>anyarray</code></td>
<td>Returns an array filled with copies of the given value, having dimensions of the lengths specified by the second argument. The optional third argument supplies lower-bound values for each dimension (which default to all <code>1</code>).</td>
<td><code>array_fill(11, ARRAY[2,3])</code> <code>{{11,11,11},{11,11,11}}</code><br><code>array_fill(7, ARRAY[3], ARRAY[2])</code> <code>[2:4]={7,7,7}</code></td>
</tr>
<tr>
<td><code>array_length</code> ( <code>anyarray</code>, <code>integer</code> ) <code>integer</code></td>
<td>Returns the length of the requested array dimension. (Produces NULL instead of 0 for empty or missing array dimensions.)</td>
<td><code>array_length(array[1,2,3], 1)</code> <code>3</code><br><code>array_length(array[]::int[], 1)</code> <code>NULL</code><br><code>array_length(array['text'], 2)</code> <code>NULL</code></td>
</tr>
<tr>
<td><code>array_lower</code> ( <code>anyarray</code>, <code>integer</code> ) <code>integer</code></td>
<td>Returns the lower bound of the requested array dimension.</td>
<td><code>array_lower('[0:2]={1,2,3}'::integer[], 1)</code> <code>0</code></td>
</tr>
<tr>
<td><code>array_ndims</code> ( <code>anyarray</code> ) <code>integer</code></td>
<td>Returns the number of dimensions of the array.</td>
<td><code>array_ndims(ARRAY[[1,2,3], [4,5,6]])</code> <code>2</code></td>
</tr>
<tr>
<td><code>array_position</code> ( <code>anycompatiblearray</code>, <code>anycompatible</code> [, <code>integer</code> ] ) <code>integer</code></td>
<td>Returns the subscript of the first occurrence of the second argument in the array, or <code>NULL</code> if it's not present. If the third argument is given, the search begins at that subscript. The array must be one-dimensional. Comparisons are done using <code>IS NOT DISTINCT FROM</code> semantics, so it is possible to search for <code>NULL</code>.</td>
<td><code>array_position(ARRAY['sun', 'mon', 'tue', 'wed', 'thu', 'fri', 'sat'], 'mon')</code> <code>2</code></td>
</tr>
<tr>
<td><code>array_positions</code> ( <code>anycompatiblearray</code>, <code>anycompatible</code> ) <code>integer[]</code></td>
<td>Returns an array of the subscripts of all occurrences of the second argument in the array given as first argument. The array must be one-dimensional. Comparisons are done using <code>IS NOT DISTINCT FROM</code> semantics, so it is possible to search for <code>NULL</code>. <code>NULL</code> is returned only if the array is <code>NULL</code>; if the value is not found in the array, an empty array is returned.</td>
<td><code>array_positions(ARRAY['A','A','B','A'], 'A')</code> <code>{1,2,4}</code></td>
</tr>
<tr>
<td><code>array_prepend</code> ( <code>anycompatible</code>, <code>anycompatiblearray</code> ) <code>anycompatiblearray</code></td>
<td>Prepends an element to the beginning of an array (same as the <code>anycompatible</code> <code>||</code> <code>anycompatiblearray</code> operator).</td>
<td><code>array_prepend(1, ARRAY[2,3])</code> <code>{1,2,3}</code></td>
</tr>
<tr>
<td><code>array_remove</code> ( <code>anycompatiblearray</code>, <code>anycompatible</code> ) <code>anycompatiblearray</code></td>
<td>Removes all elements equal to the given value from the array. The array must be one-dimensional. Comparisons are done using <code>IS NOT DISTINCT FROM</code> semantics, so it is possible to remove <code>NULL</code>s.</td>
<td><code>array_remove(ARRAY[1,2,3,2], 2)</code> <code>{1,3}</code></td>
</tr>
<tr>
<td><code>array_replace</code> ( <code>anycompatiblearray</code>, <code>anycompatible</code>, <code>anycompatible</code> ) <code>anycompatiblearray</code></td>
<td>Replaces each array element equal to the second argument with the third argument.</td>
<td><code>array_replace(ARRAY[1,2,5,4], 5, 3)</code> <code>{1,2,3,4}</code></td>
</tr>
<tr>
<td><code>array_sample</code> ( <code>array</code> <code>anyarray</code>, <code>n</code> <code>integer</code> ) <code>anyarray</code></td>
<td>Returns an array of <code>n</code> items randomly selected from <code>array</code>. <code>n</code> may not exceed the length of <code>array</code>'s first dimension. If <code>array</code> is multi-dimensional, an “item” is a slice having a given first subscript.</td>
<td><code>array_sample(ARRAY[1,2,3,4,5,6], 3)</code> <code>{2,6,1}</code><br><code>array_sample(ARRAY[[1,2],[3,4],[5,6]], 2)</code> <code>{{5,6},{1,2}}</code></td>
</tr>
<tr>
<td><code>array_shuffle</code> ( <code>anyarray</code> ) <code>anyarray</code></td>
<td>Randomly shuffles the first dimension of the array.</td>
<td><code>array_shuffle(ARRAY[[1,2],[3,4],[5,6]])</code> <code>{{5,6},{1,2},{3,4}}</code></td>
</tr>
<tr>
<td><a id="function-array-to-string"></a>
 `array_to_string` ( `array` `anyarray`, `delimiter` `text` [, `null_string` `text` ] ) `text`</td>
<td>Converts each array element to its text representation, and concatenates those separated by the <code>delimiter</code> string. If <code>null_string</code> is given and is not <code>NULL</code>, then <code>NULL</code> array entries are represented by that string; otherwise, they are omitted. See also <a href="string-functions-and-operators.md#function-string-to-array"><code>string_to_array</code></a>.</td>
<td><code>array_to_string(ARRAY[1, 2, 3, NULL, 5], ',', '<em>')</code> <code>1,2,3,</em>,5</code></td>
</tr>
<tr>
<td><code>array_upper</code> ( <code>anyarray</code>, <code>integer</code> ) <code>integer</code></td>
<td>Returns the upper bound of the requested array dimension.</td>
<td><code>array_upper(ARRAY[1,8,3,7], 1)</code> <code>4</code></td>
</tr>
<tr>
<td><code>cardinality</code> ( <code>anyarray</code> ) <code>integer</code></td>
<td>Returns the total number of elements in the array, or 0 if the array is empty.</td>
<td><code>cardinality(ARRAY[[1,2],[3,4]])</code> <code>4</code></td>
</tr>
<tr>
<td><code>trim_array</code> ( <code>array</code> <code>anyarray</code>, <code>n</code> <code>integer</code> ) <code>anyarray</code></td>
<td>Trims an array by removing the last <code>n</code> elements. If the array is multidimensional, only the first dimension is trimmed.</td>
<td><code>trim_array(ARRAY[1,2,3,4,5,6], 2)</code> <code>{1,2,3,4}</code></td>
</tr>
<tr>
<td><code>unnest</code> ( <code>anyarray</code> ) <code>setof anyelement</code></td>
<td>Expands an array into a set of rows. The array's elements are read out in storage order.</td>
<td><p><code>unnest(ARRAY[1,2])</code></p>
<pre><code>
 1
 2</code></pre><br><p><code>unnest(ARRAY[['foo','bar'],['baz','quux']])</code></p>
<pre><code>
 foo
 bar
 baz
 quux</code></pre></td>
</tr>
<tr>
<td><code>unnest</code> ( <code>anyarray</code>, <code>anyarray</code> [, ... ] ) <code>setof anyelement, anyelement [, ... ]</code></td>
<td>Expands multiple arrays (possibly of different data types) into a set of rows. If the arrays are not all the same length then the shorter ones are padded with <code>NULL</code>s. This form is only allowed in a query's FROM clause; see <a href="../queries/table-expressions.md#queries-tablefunctions">Table Functions</a>.</td>
<td><p><code>select * from unnest(ARRAY[1,2], ARRAY['foo','bar','baz']) as x(a,b)</code></p>
<pre><code>
 a |  b
---+-----
 1 | foo
 2 | bar
   | baz</code></pre></td>
</tr>
</tbody>
</table>


 See also [Aggregate Functions](aggregate-functions.md#functions-aggregate) about the aggregate function `array_agg` for use with arrays.
