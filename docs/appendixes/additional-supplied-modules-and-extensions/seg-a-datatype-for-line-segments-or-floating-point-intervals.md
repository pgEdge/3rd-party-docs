<a id="seg"></a>

## seg — a datatype for line segments or floating point intervals


 This module implements a data type `seg` for representing line segments, or floating point intervals. `seg` can represent uncertainty in the interval endpoints, making it especially useful for representing laboratory measurements.


 This module is considered “trusted”, that is, it can be installed by non-superusers who have `CREATE` privilege on the current database.
 <a id="seg-rationale"></a>

### Rationale


 The geometry of measurements is usually more complex than that of a point in a numeric continuum. A measurement is usually a segment of that continuum with somewhat fuzzy limits. The measurements come out as intervals because of uncertainty and randomness, as well as because the value being measured may naturally be an interval indicating some condition, such as the temperature range of stability of a protein.


 Using just common sense, it appears more convenient to store such data as intervals, rather than pairs of numbers. In practice, it even turns out more efficient in most applications.


 Further along the line of common sense, the fuzziness of the limits suggests that the use of traditional numeric data types leads to a certain loss of information. Consider this: your instrument reads 6.50, and you input this reading into the database. What do you get when you fetch it? Watch:

```

test=> SELECT 6.50::float8 AS "pH";
 pH
---
6.5
(1 row)
```
 In the world of measurements, 6.50 is not the same as 6.5. It may sometimes be critically different. The experimenters usually write down (and publish) the digits they trust. 6.50 is actually a fuzzy interval contained within a bigger and even fuzzier interval, 6.5, with their center points being (probably) the only common feature they share. We definitely do not want such different data items to appear the same.


 Conclusion? It is nice to have a special data type that can record the limits of an interval with arbitrarily variable precision. Variable in the sense that each data element records its own precision.


 Check this out:

```

test=> SELECT '6.25 .. 6.50'::seg AS "pH";
          pH
------------
6.25 .. 6.50
(1 row)
```

  <a id="seg-syntax"></a>

### Syntax


 The external representation of an interval is formed using one or two floating-point numbers joined by the range operator (`..` or `...`). Alternatively, it can be specified as a center point plus or minus a deviation. Optional certainty indicators (`<`, `>` or `~`) can be stored as well. (Certainty indicators are ignored by all the built-in operators, however.) [`seg` External Representations](#seg-repr-table) gives an overview of allowed representations; [Examples of Valid `seg` Input](#seg-input-examples) shows some examples.


 In [`seg` External Representations](#seg-repr-table), *x*, *y*, and *delta* denote floating-point numbers. *x* and *y*, but not *delta*, can be preceded by a certainty indicator.
 <a id="seg-repr-table"></a>

**Table: `seg` External Representations**

| <em>x</em> | Single value (zero-length interval) |
| <em>x</em><code> .. </code><em>y</em> | Interval from *x* to *y* |
| <em>x</em><code> (+-) </code><em>delta</em> | Interval from *x* - *delta* to *x* + *delta* |
| <em>x</em><code> ..</code> | Open interval with lower bound *x* |
| <code>.. </code><em>x</em> | Open interval with upper bound *x* |
 <a id="seg-input-examples"></a>

**Table: Examples of Valid `seg` Input**

|  |  |
| --- | --- |
| `5.0` | Creates a zero-length segment (a point, if you will) |
| `~5.0` | Creates a zero-length segment and records `~` in the data. `~` is ignored by `seg` operations, but is preserved as a comment. |
| `<5.0` | Creates a point at 5.0. `<` is ignored but is preserved as a comment. |
| `>5.0` | Creates a point at 5.0. `>` is ignored but is preserved as a comment. |
| `5(+-)0.3` | Creates an interval `4.7 .. 5.3`. Note that the `(+-)` notation isn't preserved. |
| `50 .. ` | Everything that is greater than or equal to 50 |
| `.. 0` | Everything that is less than or equal to 0 |
| `1.5e-2 .. 2E-2 ` | Creates an interval `0.015 .. 0.02` |
| `1 ... 2` | The same as `1...2`, or `1 .. 2`, or `1..2` (spaces around the range operator are ignored) |


 Because the `...` operator is widely used in data sources, it is allowed as an alternative spelling of the `..` operator. Unfortunately, this creates a parsing ambiguity: it is not clear whether the upper bound in `0...23` is meant to be `23` or `0.23`. This is resolved by requiring at least one digit before the decimal point in all numbers in `seg` input.


 As a sanity check, `seg` rejects intervals with the lower bound greater than the upper, for example `5 .. 2`.
  <a id="seg-precision"></a>

### Precision


 `seg` values are stored internally as pairs of 32-bit floating point numbers. This means that numbers with more than 7 significant digits will be truncated.


 Numbers with 7 or fewer significant digits retain their original precision. That is, if your query returns 0.00, you will be sure that the trailing zeroes are not the artifacts of formatting: they reflect the precision of the original data. The number of leading zeroes does not affect precision: the value 0.0067 is considered to have just 2 significant digits.
  <a id="seg-usage"></a>

### Usage


 The `seg` module includes a GiST index operator class for `seg` values. The operators supported by the GiST operator class are shown in [Seg GiST Operators](#seg-gist-operators).
 <a id="seg-gist-operators"></a>

**Table: Seg GiST Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>seg</code> <code>&lt;&lt;</code> <code>seg</code> <code>boolean</code></td>
<td>Is the first <code>seg</code> entirely to the left of the second? [a, b] << [c, d] is true if b < c.</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>&gt;&gt;</code> <code>seg</code> <code>boolean</code></td>
<td>Is the first <code>seg</code> entirely to the right of the second? [a, b] >> [c, d] is true if a > d.</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>&amp;&lt;</code> <code>seg</code> <code>boolean</code></td>
<td>Does the first <code>seg</code> not extend to the right of the second? [a, b] &< [c, d] is true if b <= d.</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>&amp;&gt;</code> <code>seg</code> <code>boolean</code></td>
<td>Does the first <code>seg</code> not extend to the left of the second? [a, b] &> [c, d] is true if a >= c.</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>=</code> <code>seg</code> <code>boolean</code></td>
<td>Are the two <code>seg</code>s equal?</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>&amp;&amp;</code> <code>seg</code> <code>boolean</code></td>
<td>Do the two <code>seg</code>s overlap?</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>@&gt;</code> <code>seg</code> <code>boolean</code></td>
<td>Does the first <code>seg</code> contain the second?</td>
<td></td>
</tr>
<tr>
<td><code>seg</code> <code>&lt;@</code> <code>seg</code> <code>boolean</code></td>
<td>Is the first <code>seg</code> contained in the second?</td>
<td></td>
</tr>
</tbody>
</table>


 In addition to the above operators, the usual comparison operators shown in [Comparison Operators](../../the-sql-language/functions-and-operators/comparison-functions-and-operators.md#functions-comparison-op-table) are available for type `seg`. These operators first compare (a) to (c), and if these are equal, compare (b) to (d). That results in reasonably good sorting in most cases, which is useful if you want to use ORDER BY with this type.
  <a id="seg-notes"></a>

### Notes


 For examples of usage, see the regression test `sql/seg.sql`.


 The mechanism that converts `(+-)` to regular ranges isn't completely accurate in determining the number of significant digits for the boundaries. For example, it adds an extra digit to the lower boundary if the resulting interval includes a power of ten:

```

postgres=> SELECT '10(+-)1'::seg AS seg;
      seg
---------
9.0 .. 11             -- should be: 9 .. 11
```


 The performance of an R-tree index can largely depend on the initial order of input values. It may be very helpful to sort the input table on the `seg` column; see the script `sort-segments.pl` for an example.
  <a id="seg-credits"></a>

### Credits


 Original author: Gene Selkov, Jr. [selkovjr@mcs.anl.gov](mailto:selkovjr@mcs.anl.gov), Mathematics and Computer Science Division, Argonne National Laboratory.


 My thanks are primarily to Prof. Joe Hellerstein ([https://dsf.berkeley.edu/jmh/](https://dsf.berkeley.edu/jmh/)) for elucidating the gist of the GiST ([http://gist.cs.berkeley.edu/](http://gist.cs.berkeley.edu/)). I am also grateful to all Postgres developers, present and past, for enabling myself to create my own world and live undisturbed in it. And I would like to acknowledge my gratitude to Argonne Lab and to the U.S. Department of Energy for the years of faithful support of my database research.
