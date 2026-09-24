<a id="cube"></a>

## cube — a multi-dimensional cube data type


 This module implements a data type `cube` for representing multidimensional cubes.


 This module is considered “trusted”, that is, it can be installed by non-superusers who have `CREATE` privilege on the current database.
 <a id="cube-syntax"></a>

### Syntax


 [Cube External Representations](#cube-repr-table) shows the valid external representations for the `cube` type. *x*, *y*, etc. denote floating-point numbers.
 <a id="cube-repr-table"></a>

**Table: Cube External Representations**

| External Syntax | Meaning |
| --- | --- |
| <em>x</em> | A one-dimensional point (or, zero-length one-dimensional interval) |
| <code>(</code><em>x</em><code>)</code> | Same as above |
| <em>x1</em><code>,</code><em>x2</em><code>,...,</code><em>xn</em> | A point in n-dimensional space, represented internally as a zero-volume cube |
| <code>(</code><em>x1</em><code>,</code><em>x2</em><code>,...,</code><em>xn</em><code>)</code> | Same as above |
| <code>(</code><em>x</em><code>),(</code><em>y</em><code>)</code> | A one-dimensional interval starting at *x* and ending at *y* or vice versa; the order does not matter |
| <code>[(</code><em>x</em><code>),(</code><em>y</em><code>)]</code> | Same as above |
| <code>(</code><em>x1</em><code>,...,</code><em>xn</em><code>),(</code><em>y1</em><code>,...,</code><em>yn</em><code>)</code> | An n-dimensional cube represented by a pair of its diagonally opposite corners |
| <code>[(</code><em>x1</em><code>,...,</code><em>xn</em><code>),(</code><em>y1</em><code>,...,</code><em>yn</em><code>)]</code> | Same as above |


 It does not matter which order the opposite corners of a cube are entered in. The `cube` functions automatically swap values if needed to create a uniform “lower left — upper right” internal representation. When the corners coincide, `cube` stores only one corner along with an “is point” flag to avoid wasting space.


 White space is ignored on input, so <code>[(</code><em>x</em><code>),(</code><em>y</em><code>)]</code> is the same as <code>[ ( </code><em>x</em><code> ), ( </code><em>y</em><code> ) ]</code>.
  <a id="cube-precision"></a>

### Precision


 Values are stored internally as 64-bit floating point numbers. This means that numbers with more than about 16 significant digits will be truncated.
  <a id="cube-usage"></a>

### Usage


 [Cube Operators](#cube-operators-table) shows the specialized operators provided for type `cube`.
 <a id="cube-operators-table"></a>

**Table: Cube Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>cube</code> <code>&amp;&amp;</code> <code>cube</code> <code>boolean</code></td>
<td>Do the cubes overlap?</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>@&gt;</code> <code>cube</code> <code>boolean</code></td>
<td>Does the first cube contain the second?</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>&lt;@</code> <code>cube</code> <code>boolean</code></td>
<td>Is the first cube contained in the second?</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>-&gt;</code> <code>integer</code> <code>float8</code></td>
<td>Extracts the <code>n</code>-th coordinate of the cube (counting from 1).</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>~&gt;</code> <code>integer</code> <code>float8</code></td>
<td>Extracts the <code>n</code>-th coordinate of the cube, counting in the following way: <code>n</code> = 2 <em> <code>k</code> - 1 means lower bound of <code>k</code>-th dimension, <code>n</code> = 2 </em> <code>k</code> means upper bound of <code>k</code>-th dimension. Negative <code>n</code> denotes the inverse value of the corresponding positive coordinate. This operator is designed for KNN-GiST support.</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>&lt;-&gt;</code> <code>cube</code> <code>float8</code></td>
<td>Computes the Euclidean distance between the two cubes.</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>&lt;#&gt;</code> <code>cube</code> <code>float8</code></td>
<td>Computes the taxicab (L-1 metric) distance between the two cubes.</td>
<td></td>
</tr>
<tr>
<td><code>cube</code> <code>&lt;=&gt;</code> <code>cube</code> <code>float8</code></td>
<td>Computes the Chebyshev (L-inf metric) distance between the two cubes.</td>
<td></td>
</tr>
</tbody>
</table>


 In addition to the above operators, the usual comparison operators shown in [Comparison Operators](../../the-sql-language/functions-and-operators/comparison-functions-and-operators.md#functions-comparison-op-table) are available for type `cube`. These operators first compare the first coordinates, and if those are equal, compare the second coordinates, etc. They exist mainly to support the b-tree index operator class for `cube`, which can be useful for example if you would like a UNIQUE constraint on a `cube` column. Otherwise, this ordering is not of much practical use.


 The `cube` module also provides a GiST index operator class for `cube` values. A `cube` GiST index can be used to search for values using the `=`, `&&`, `@>`, and `<@` operators in `WHERE` clauses.


 In addition, a `cube` GiST index can be used to find nearest neighbors using the metric operators `<->`, `<#>`, and `<=>` in `ORDER BY` clauses. For example, the nearest neighbor of the 3-D point (0.5, 0.5, 0.5) could be found efficiently with:

```sql

SELECT c FROM test ORDER BY c <-> cube(ARRAY[0.5, 0.5, 0.5]) LIMIT 1;
```


 The `~>` operator can also be used in this way to efficiently retrieve the first few values sorted by a selected coordinate. For example, to get the first few cubes ordered by the first coordinate (lower left corner) ascending one could use the following query:

```sql

SELECT c FROM test ORDER BY c ~> 1 LIMIT 5;
```
 And to get 2-D cubes ordered by the first coordinate of the upper right corner descending:

```sql

SELECT c FROM test ORDER BY c ~> 3 DESC LIMIT 5;
```


 [Cube Functions](#cube-functions-table) shows the available functions.
 <a id="cube-functions-table"></a>

**Table: Cube Functions**

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
<td><code>cube</code> ( <code>float8</code> ) <code>cube</code></td>
<td>Makes a one dimensional cube with both coordinates the same.</td>
<td><code>cube(1)</code> <code>(1)</code></td>
</tr>
<tr>
<td><code>cube</code> ( <code>float8</code>, <code>float8</code> ) <code>cube</code></td>
<td>Makes a one dimensional cube.</td>
<td><code>cube(1, 2)</code> <code>(1),(2)</code></td>
</tr>
<tr>
<td><code>cube</code> ( <code>float8[]</code> ) <code>cube</code></td>
<td>Makes a zero-volume cube using the coordinates defined by the array.</td>
<td><code>cube(ARRAY[1,2,3])</code> <code>(1, 2, 3)</code></td>
</tr>
<tr>
<td><code>cube</code> ( <code>float8[]</code>, <code>float8[]</code> ) <code>cube</code></td>
<td>Makes a cube with upper right and lower left coordinates as defined by the two arrays, which must be of the same length.</td>
<td><code>cube(ARRAY[1,2], ARRAY[3,4])</code> <code>(1, 2),(3, 4)</code></td>
</tr>
<tr>
<td><code>cube</code> ( <code>cube</code>, <code>float8</code> ) <code>cube</code></td>
<td>Makes a new cube by adding a dimension on to an existing cube, with the same values for both endpoints of the new coordinate. This is useful for building cubes piece by piece from calculated values.</td>
<td><code>cube('(1,2),(3,4)'::cube, 5)</code> <code>(1, 2, 5),(3, 4, 5)</code></td>
</tr>
<tr>
<td><code>cube</code> ( <code>cube</code>, <code>float8</code>, <code>float8</code> ) <code>cube</code></td>
<td>Makes a new cube by adding a dimension on to an existing cube. This is useful for building cubes piece by piece from calculated values.</td>
<td><code>cube('(1,2),(3,4)'::cube, 5, 6)</code> <code>(1, 2, 5),(3, 4, 6)</code></td>
</tr>
<tr>
<td><code>cube_dim</code> ( <code>cube</code> ) <code>integer</code></td>
<td>Returns the number of dimensions of the cube.</td>
<td><code>cube_dim('(1,2),(3,4)')</code> <code>2</code></td>
</tr>
<tr>
<td><code>cube_ll_coord</code> ( <code>cube</code>, <code>integer</code> ) <code>float8</code></td>
<td>Returns the <code>n</code>-th coordinate value for the lower left corner of the cube.</td>
<td><code>cube_ll_coord('(1,2),(3,4)', 2)</code> <code>2</code></td>
</tr>
<tr>
<td><code>cube_ur_coord</code> ( <code>cube</code>, <code>integer</code> ) <code>float8</code></td>
<td>Returns the <code>n</code>-th coordinate value for the upper right corner of the cube.</td>
<td><code>cube_ur_coord('(1,2),(3,4)', 2)</code> <code>4</code></td>
</tr>
<tr>
<td><code>cube_is_point</code> ( <code>cube</code> ) <code>boolean</code></td>
<td>Returns true if the cube is a point, that is, the two defining corners are the same.</td>
<td><code>cube_is_point(cube(1,1))</code> <code>t</code></td>
</tr>
<tr>
<td><code>cube_distance</code> ( <code>cube</code>, <code>cube</code> ) <code>float8</code></td>
<td>Returns the distance between two cubes. If both cubes are points, this is the normal distance function.</td>
<td><code>cube_distance('(1,2)', '(3,4)')</code> <code>2.8284271247461903</code></td>
</tr>
<tr>
<td><code>cube_subset</code> ( <code>cube</code>, <code>integer[]</code> ) <code>cube</code></td>
<td>Makes a new cube from an existing cube, using a list of dimension indexes from an array. Can be used to extract the endpoints of a single dimension, or to drop dimensions, or to reorder them as desired.</td>
<td><code>cube_subset(cube('(1,3,5),(6,7,8)'), ARRAY[2])</code> <code>(3),(7)</code><br><code>cube_subset(cube('(1,3,5),(6,7,8)'), ARRAY[3,2,1,1])</code> <code>(5, 3, 1, 1),(8, 7, 6, 6)</code></td>
</tr>
<tr>
<td><code>cube_union</code> ( <code>cube</code>, <code>cube</code> ) <code>cube</code></td>
<td>Produces the union of two cubes.</td>
<td><code>cube_union('(1,2)', '(3,4)')</code> <code>(1, 2),(3, 4)</code></td>
</tr>
<tr>
<td><code>cube_inter</code> ( <code>cube</code>, <code>cube</code> ) <code>cube</code></td>
<td>Produces the intersection of two cubes.</td>
<td><code>cube_inter('(1,2)', '(3,4)')</code> <code>(3, 4),(1, 2)</code></td>
</tr>
<tr>
<td><code>cube_enlarge</code> ( <code>c</code> <code>cube</code>, <code>r</code> <code>double</code>, <code>n</code> <code>integer</code> ) <code>cube</code></td>
<td>Increases the size of the cube by the specified radius <code>r</code> in at least <code>n</code> dimensions. If the radius is negative the cube is shrunk instead. All defined dimensions are changed by the radius <code>r</code>. Lower-left coordinates are decreased by <code>r</code> and upper-right coordinates are increased by <code>r</code>. If a lower-left coordinate is increased to more than the corresponding upper-right coordinate (this can only happen when <code>r</code> < 0) than both coordinates are set to their average. If <code>n</code> is greater than the number of defined dimensions and the cube is being enlarged (<code>r</code> > 0), then extra dimensions are added to make <code>n</code> altogether; 0 is used as the initial value for the extra coordinates. This function is useful for creating bounding boxes around a point for searching for nearby points.</td>
<td><code>cube_enlarge('(1,2),(3,4)', 0.5, 3)</code> <code>(0.5, 1.5, -0.5),(3.5, 4.5, 0.5)</code></td>
</tr>
</tbody>
</table>
  <a id="cube-defaults"></a>

### Defaults


 This union:


```sql

SELECT cube_union('(0,5,2),(2,3,1)', '0');
cube_union
-------------------
(0, 0, 0),(2, 5, 2)
(1 row)
```


 does not contradict common sense, neither does the intersection:


```sql

SELECT cube_inter('(0,-1),(1,1)', '(-2),(2)');
cube_inter
-------------
(0, 0),(1, 0)
(1 row)
```


 In all binary operations on differently-dimensioned cubes, the lower-dimensional one is assumed to be a Cartesian projection, i. e., having zeroes in place of coordinates omitted in the string representation. The above examples are equivalent to:


```

cube_union('(0,5,2),(2,3,1)','(0,0,0),(0,0,0)');
cube_inter('(0,-1),(1,1)','(-2,0),(2,0)');
```


 The following containment predicate uses the point syntax, while in fact the second argument is internally represented by a box. This syntax makes it unnecessary to define a separate point type and functions for (box,point) predicates.


```sql

SELECT cube_contains('(0,0),(1,1)', '0.5,0.5');
cube_contains
--------------
t
(1 row)
```
  <a id="cube-notes"></a>

### Notes


 For examples of usage, see the regression test `sql/cube.sql`.


 To make it harder for people to break things, there is a limit of 100 on the number of dimensions of cubes. This is set in `cubedata.h` if you need something bigger.
  <a id="cube-credits"></a>

### Credits


 Original author: Gene Selkov, Jr. [selkovjr@mcs.anl.gov](mailto:selkovjr@mcs.anl.gov), Mathematics and Computer Science Division, Argonne National Laboratory.


 My thanks are primarily to Prof. Joe Hellerstein ([https://dsf.berkeley.edu/jmh/](https://dsf.berkeley.edu/jmh/)) for elucidating the gist of the GiST ([http://gist.cs.berkeley.edu/](http://gist.cs.berkeley.edu/)), and to his former student Andy Dong for his example written for Illustra. I am also grateful to all Postgres developers, present and past, for enabling myself to create my own world and live undisturbed in it. And I would like to acknowledge my gratitude to Argonne Lab and to the U.S. Department of Energy for the years of faithful support of my database research.


 Minor updates to this package were made by Bruno Wolff III [bruno@wolff.to](mailto:bruno@wolff.to) in August/September of 2002. These include changing the precision from single precision to double precision and adding some new functions.


 Additional updates were made by Joshua Reich [josh@root.net](mailto:josh@root.net) in July 2006. These include `cube(float8[], float8[])` and cleaning up the code to use the V1 call protocol instead of the deprecated V0 protocol.
