<a id="intarray"></a>

## intarray — manipulate arrays of integers


 The `intarray` module provides a number of useful functions and operators for manipulating null-free arrays of integers. There is also support for indexed searches using some of the operators.


 All of these operations will throw an error if a supplied array contains any NULL elements.


 Many of these operations are only sensible for one-dimensional arrays. Although they will accept input arrays of more dimensions, the data is treated as though it were a linear array in storage order.


 This module is considered “trusted”, that is, it can be installed by non-superusers who have `CREATE` privilege on the current database.
 <a id="intarray-funcs-ops"></a>

### `intarray` Functions and Operators


 The functions provided by the `intarray` module are shown in [`intarray` Functions](#intarray-func-table), the operators in [`intarray` Operators](#intarray-op-table).
 <a id="intarray-func-table"></a>

**Table: `intarray` Functions**

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
<td><code>icount</code> ( <code>integer[]</code> ) <code>integer</code></td>
<td>Returns the number of elements in the array.</td>
<td><code>icount('{1,2,3}'::integer[])</code> <code>3</code></td>
</tr>
<tr>
<td><code>sort</code> ( <code>integer[]</code>, <code>dir</code> <code>text</code> ) <code>integer[]</code></td>
<td>Sorts the array in either ascending or descending order. <code>dir</code> must be <code>asc</code> or <code>desc</code>.</td>
<td><code>sort('{1,3,2}'::integer[], 'desc')</code> <code>{3,2,1}</code></td>
</tr>
<tr>
<td><code>sort</code> ( <code>integer[]</code> ) <code>integer[]</code></td>
<td><code>sort_asc</code> ( <code>integer[]</code> ) <code>integer[]</code></td>
<td>Sorts in ascending order.<br><code>sort(array[11,77,44])</code> <code>{11,44,77}</code></td>
</tr>
<tr>
<td><code>sort_desc</code> ( <code>integer[]</code> ) <code>integer[]</code></td>
<td>Sorts in descending order.</td>
<td><code>sort_desc(array[11,77,44])</code> <code>{77,44,11}</code></td>
</tr>
<tr>
<td><code>uniq</code> ( <code>integer[]</code> ) <code>integer[]</code></td>
<td>Removes adjacent duplicates. Often used with <code>sort</code> to remove all duplicates.</td>
<td><code>uniq('{1,2,2,3,1,1}'::integer[])</code> <code>{1,2,3,1}</code><br><code>uniq(sort('{1,2,3,2,1}'::integer[]))</code> <code>{1,2,3}</code></td>
</tr>
<tr>
<td><code>idx</code> ( <code>integer[]</code>, <code>item</code> <code>integer</code> ) <code>integer</code></td>
<td>Returns index of the first array element matching <code>item</code>, or 0 if no match.</td>
<td><code>idx(array[11,22,33,22,11], 22)</code> <code>2</code></td>
</tr>
<tr>
<td><code>subarray</code> ( <code>integer[]</code>, <code>start</code> <code>integer</code>, <code>len</code> <code>integer</code> ) <code>integer[]</code></td>
<td>Extracts the portion of the array starting at position <code>start</code>, with <code>len</code> elements.</td>
<td><code>subarray('{1,2,3,2,1}'::integer[], 2, 3)</code> <code>{2,3,2}</code></td>
</tr>
<tr>
<td><code>subarray</code> ( <code>integer[]</code>, <code>start</code> <code>integer</code> ) <code>integer[]</code></td>
<td>Extracts the portion of the array starting at position <code>start</code>.</td>
<td><code>subarray('{1,2,3,2,1}'::integer[], 2)</code> <code>{2,3,2,1}</code></td>
</tr>
<tr>
<td><code>intset</code> ( <code>integer</code> ) <code>integer[]</code></td>
<td>Makes a single-element array.</td>
<td><code>intset(42)</code> <code>{42}</code></td>
</tr>
</tbody>
</table>
 <a id="intarray-op-table"></a>

**Table: `intarray` Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>integer[]</code> <code>&amp;&amp;</code> <code>integer[]</code> <code>boolean</code></td>
<td>Do arrays overlap (have at least one element in common)?</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>@&gt;</code> <code>integer[]</code> <code>boolean</code></td>
<td>Does left array contain right array?</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>&lt;@</code> <code>integer[]</code> <code>boolean</code></td>
<td>Is left array contained in right array?</td>
<td></td>
</tr>
<tr>
<td><code>#</code> <code>integer[]</code> <code>integer</code></td>
<td>Returns the number of elements in the array.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>#</code> <code>integer</code> <code>integer</code></td>
<td>Returns index of the first array element matching the right argument, or 0 if no match. (Same as <code>idx</code> function.)</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>+</code> <code>integer</code> <code>integer[]</code></td>
<td>Adds element to end of array.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>+</code> <code>integer[]</code> <code>integer[]</code></td>
<td>Concatenates the arrays.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>-</code> <code>integer</code> <code>integer[]</code></td>
<td>Removes entries matching the right argument from the array.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>-</code> <code>integer[]</code> <code>integer[]</code></td>
<td>Removes elements of the right array from the left array.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>|</code> <code>integer</code> <code>integer[]</code></td>
<td>Computes the union of the arguments.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>|</code> <code>integer[]</code> <code>integer[]</code></td>
<td>Computes the union of the arguments.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>&amp;</code> <code>integer[]</code> <code>integer[]</code></td>
<td>Computes the intersection of the arguments.</td>
<td></td>
</tr>
<tr>
<td><code>integer[]</code> <code>@@</code> <code>query_int</code> <code>boolean</code></td>
<td>Does array satisfy query? (see below)</td>
<td></td>
</tr>
<tr>
<td><code>query_int</code> <code>~~</code> <code>integer[]</code> <code>boolean</code></td>
<td>Does array satisfy query? (commutator of <code>@@</code>)</td>
<td></td>
</tr>
</tbody>
</table>


 The operators `&&`, `@>` and `<@` are equivalent to PostgreSQL's built-in operators of the same names, except that they work only on integer arrays that do not contain nulls, while the built-in operators work for any array type. This restriction makes them faster than the built-in operators in many cases.


 The `@@` and `~~` operators test whether an array satisfies a *query*, which is expressed as a value of a specialized data type `query_int`. A *query* consists of integer values that are checked against the elements of the array, possibly combined using the operators `&` (AND), `|` (OR), and `!` (NOT). Parentheses can be used as needed. For example, the query `1&(2|3)` matches arrays that contain 1 and also contain either 2 or 3.
  <a id="intarray-index"></a>

### Index Support


 `intarray` provides index support for the `&&`, `@>`, and `@@` operators, as well as regular array equality.


 Two parameterized GiST index operator classes are provided: `gist__int_ops` (used by default) is suitable for small- to medium-size data sets, while `gist__intbig_ops` uses a larger signature and is more suitable for indexing large data sets (i.e., columns containing a large number of distinct array values). The implementation uses an RD-tree data structure with built-in lossy compression.


 `gist__int_ops` approximates an integer set as an array of integer ranges. Its optional integer parameter `numranges` determines the maximum number of ranges in one index key. The default value of `numranges` is 100. Valid values are between 1 and 253. Using larger arrays as GiST index keys leads to a more precise search (scanning a smaller fraction of the index and fewer heap pages), at the cost of a larger index.


 `gist__intbig_ops` approximates an integer set as a bitmap signature. Its optional integer parameter `siglen` determines the signature length in bytes. The default signature length is 16 bytes. Valid values of signature length are between 1 and 2024 bytes. Longer signatures lead to a more precise search (scanning a smaller fraction of the index and fewer heap pages), at the cost of a larger index.


 There is also a non-default GIN operator class `gin__int_ops`, which supports these operators as well as `<@`.


 The choice between GiST and GIN indexing depends on the relative performance characteristics of GiST and GIN, which are discussed elsewhere.
  <a id="intarray-example"></a>

### Example


```

-- a message can be in one or more sections
CREATE TABLE message (mid INT PRIMARY KEY, sections INT[], ...);

-- create specialized index with signature length of 32 bytes
CREATE INDEX message_rdtree_idx ON message USING GIST (sections gist__intbig_ops (siglen = 32));

-- select messages in section 1 OR 2 - OVERLAP operator
SELECT message.mid FROM message WHERE message.sections && '{1,2}';

-- select messages in sections 1 AND 2 - CONTAINS operator
SELECT message.mid FROM message WHERE message.sections @> '{1,2}';

-- the same, using QUERY operator
SELECT message.mid FROM message WHERE message.sections @@ '1&2'::query_int;
```
  <a id="intarray-benchmark"></a>

### Benchmark


 The source directory `contrib/intarray/bench` contains a benchmark test suite, which can be run against an installed PostgreSQL server. (It also requires `DBD::Pg` to be installed.) To run:


```

cd .../contrib/intarray/bench
createdb TEST
psql -c "CREATE EXTENSION intarray" TEST
./create_test.pl | psql TEST
./bench.pl
```


 The `bench.pl` script has numerous options, which are displayed when it is run without any arguments.
  <a id="intarray-Authors"></a>

### Authors


 All work was done by Teodor Sigaev ([teodor@sigaev.ru](mailto:teodor@sigaev.ru)) and Oleg Bartunov ([oleg@sai.msu.su](mailto:oleg@sai.msu.su)). See [http://www.sai.msu.su/~megera/postgres/gist/](http://www.sai.msu.su/~megera/postgres/gist/) for additional information. Andrey Oktyabrski did a great work on adding new functions and operations.
