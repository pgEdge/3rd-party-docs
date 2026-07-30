<a id="functions-math"></a>

## Mathematical Functions and Operators


 Mathematical operators are provided for many PostgreSQL types. For types without standard mathematical conventions (e.g., date/time types) we describe the actual behavior in subsequent sections.


 [Mathematical Operators](#functions-math-op-table) shows the mathematical operators that are available for the standard numeric types. Unless otherwise noted, operators shown as accepting *numeric_type* are available for all the types `smallint`, `integer`, `bigint`, `numeric`, `real`, and `double precision`. Operators shown as accepting *integral_type* are available for the types `smallint`, `integer`, and `bigint`. Except where noted, each form of an operator returns the same data type as its argument(s). Calls involving multiple argument data types, such as `integer` `+` `numeric`, are resolved by using the type appearing later in these lists.
 <a id="functions-math-op-table"></a>

**Table: Mathematical Operators**

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
<td><em>numeric_type</em> <code>+</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Addition</td>
<td><code>2 + 3</code> <code>5</code></td>
</tr>
<tr>
<td><code>+</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Unary plus (no operation)</td>
<td><code>+ 3.5</code> <code>3.5</code></td>
</tr>
<tr>
<td><em>numeric_type</em> <code>-</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Subtraction</td>
<td><code>2 - 3</code> <code>-1</code></td>
</tr>
<tr>
<td><code>-</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Negation</td>
<td><code>- (-4)</code> <code>4</code></td>
</tr>
<tr>
<td><em>numeric_type</em> <code><em></code> </em>numeric_type* <em>numeric_type</em></td>
<td>Multiplication</td>
<td><code>2 * 3</code> <code>6</code></td>
</tr>
<tr>
<td><em>numeric_type</em> <code>/</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Division (for integral types, division truncates the result towards zero)</td>
<td><code>5.0 / 2</code> <code>2.5000000000000000</code><br><code>5 / 2</code> <code>2</code><br><code>(-5) / 2</code> <code>-2</code></td>
</tr>
<tr>
<td><em>numeric_type</em> <code>%</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Modulo (remainder); available for <code>smallint</code>, <code>integer</code>, <code>bigint</code>, and <code>numeric</code></td>
<td><code>5 % 4</code> <code>1</code></td>
</tr>
<tr>
<td><code>numeric</code> <code>^</code> <code>numeric</code> <code>numeric</code></td>
<td><code>double precision</code> <code>^</code> <code>double precision</code> <code>double precision</code></td>
<td>Exponentiation<br><code>2 ^ 3</code> <code>8</code><br>Unlike typical mathematical practice, multiple uses of <code>^</code> will associate left to right by default:<br><code>2 ^ 3 ^ 3</code> <code>512</code><br><code>2 ^ (3 ^ 3)</code> <code>134217728</code></td>
</tr>
<tr>
<td><code>|/</code> <code>double precision</code> <code>double precision</code></td>
<td>Square root</td>
<td><code>|/ 25.0</code> <code>5</code></td>
</tr>
<tr>
<td><code>||/</code> <code>double precision</code> <code>double precision</code></td>
<td>Cube root</td>
<td><code>||/ 64.0</code> <code>4</code></td>
</tr>
<tr>
<td><code>@</code> <em>numeric_type</em> <em>numeric_type</em></td>
<td>Absolute value</td>
<td><code>@ -5.0</code> <code>5.0</code></td>
</tr>
<tr>
<td><em>integral_type</em> <code>&amp;</code> <em>integral_type</em> <em>integral_type</em></td>
<td>Bitwise AND</td>
<td><code>91 &amp; 15</code> <code>11</code></td>
</tr>
<tr>
<td><em>integral_type</em> <code>|</code> <em>integral_type</em> <em>integral_type</em></td>
<td>Bitwise OR</td>
<td><code>32 | 3</code> <code>35</code></td>
</tr>
<tr>
<td><em>integral_type</em> <code>#</code> <em>integral_type</em> <em>integral_type</em></td>
<td>Bitwise exclusive OR</td>
<td><code>17 # 5</code> <code>20</code></td>
</tr>
<tr>
<td><code>~</code> <em>integral_type</em> <em>integral_type</em></td>
<td>Bitwise NOT</td>
<td><code>~1</code> <code>-2</code></td>
</tr>
<tr>
<td><em>integral_type</em> <code>&lt;&lt;</code> <code>integer</code> <em>integral_type</em></td>
<td>Bitwise shift left</td>
<td><code>1 &lt;&lt; 4</code> <code>16</code></td>
</tr>
<tr>
<td><em>integral_type</em> <code>&gt;&gt;</code> <code>integer</code> <em>integral_type</em></td>
<td>Bitwise shift right</td>
<td><code>8 &gt;&gt; 2</code> <code>2</code></td>
</tr>
</tbody>
</table>


 [Mathematical Functions](#functions-math-func-table) shows the available mathematical functions. Many of these functions are provided in multiple forms with different argument types. Except where noted, any given form of a function returns the same data type as its argument(s); cross-type cases are resolved in the same way as explained above for operators. The functions working with `double precision` data are mostly implemented on top of the host system's C library; accuracy and behavior in boundary cases can therefore vary depending on the host system.
 <a id="functions-math-func-table"></a>

**Table: Mathematical Functions**

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
<td><code>abs</code> ( <em>numeric_type</em> ) <em>numeric_type</em></td>
<td>Absolute value</td>
<td><code>abs(-17.4)</code> <code>17.4</code></td>
</tr>
<tr>
<td><code>cbrt</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Cube root</td>
<td><code>cbrt(64.0)</code> <code>4</code></td>
</tr>
<tr>
<td><code>ceil</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>ceil</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Nearest integer greater than or equal to argument<br><code>ceil(42.2)</code> <code>43</code><br><code>ceil(-42.8)</code> <code>-42</code></td>
</tr>
<tr>
<td><code>ceiling</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>ceiling</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Nearest integer greater than or equal to argument (same as <code>ceil</code>)<br><code>ceiling(95.3)</code> <code>96</code></td>
</tr>
<tr>
<td><code>degrees</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Converts radians to degrees</td>
<td><code>degrees(0.5)</code> <code>28.64788975654116</code></td>
</tr>
<tr>
<td><code>div</code> ( <code>y</code> <code>numeric</code>, <code>x</code> <code>numeric</code> ) <code>numeric</code></td>
<td>Integer quotient of <code>y</code>/<code>x</code> (truncates towards zero)</td>
<td><code>div(9, 4)</code> <code>2</code></td>
</tr>
<tr>
<td><code>erf</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Error function</td>
<td><code>erf(1.0)</code> <code>0.8427007929497149</code></td>
</tr>
<tr>
<td><code>erfc</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Complementary error function (<code>1 - erf(x)</code>, without loss of precision for large inputs)</td>
<td><code>erfc(1.0)</code> <code>0.15729920705028513</code></td>
</tr>
<tr>
<td><code>exp</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>exp</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Exponential (<code>e</code> raised to the given power)<br><code>exp(1.0)</code> <code>2.7182818284590452</code></td>
</tr>
<tr>
<td><a id="function-factorial"></a>
 `factorial` ( `bigint` ) `numeric`</td>
<td>Factorial</td>
<td><code>factorial(5)</code> <code>120</code></td>
</tr>
<tr>
<td><code>floor</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>floor</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Nearest integer less than or equal to argument<br><code>floor(42.8)</code> <code>42</code><br><code>floor(-42.8)</code> <code>-43</code></td>
</tr>
<tr>
<td><code>gamma</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Gamma function</td>
<td><code>gamma(0.5)</code> <code>1.772453850905516</code><br><code>gamma(6)</code> <code>120</code></td>
</tr>
<tr>
<td><code>gcd</code> ( <em>numeric_type</em>, <em>numeric_type</em> ) <em>numeric_type</em></td>
<td>Greatest common divisor (the largest positive number that divides both inputs with no remainder); returns <code>0</code> if both inputs are zero; available for <code>integer</code>, <code>bigint</code>, and <code>numeric</code></td>
<td><code>gcd(1071, 462)</code> <code>21</code></td>
</tr>
<tr>
<td><code>lcm</code> ( <em>numeric_type</em>, <em>numeric_type</em> ) <em>numeric_type</em></td>
<td>Least common multiple (the smallest strictly positive number that is an integral multiple of both inputs); returns <code>0</code> if either input is zero; available for <code>integer</code>, <code>bigint</code>, and <code>numeric</code></td>
<td><code>lcm(1071, 462)</code> <code>23562</code></td>
</tr>
<tr>
<td><code>lgamma</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Natural logarithm of the absolute value of the gamma function</td>
<td><code>lgamma(1000)</code> <code>5905.220423209181</code></td>
</tr>
<tr>
<td><code>ln</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>ln</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Natural logarithm<br><code>ln(2.0)</code> <code>0.6931471805599453</code></td>
</tr>
<tr>
<td><code>log</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>log</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Base 10 logarithm<br><code>log(100)</code> <code>2</code></td>
</tr>
<tr>
<td><code>log10</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>log10</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Base 10 logarithm (same as <code>log</code>)<br><code>log10(1000)</code> <code>3</code></td>
</tr>
<tr>
<td><code>log</code> ( <code>b</code> <code>numeric</code>, <code>x</code> <code>numeric</code> ) <code>numeric</code></td>
<td>Logarithm of <code>x</code> to base <code>b</code></td>
<td><code>log(2.0, 64.0)</code> <code>6.0000000000000000</code></td>
</tr>
<tr>
<td><code>min_scale</code> ( <code>numeric</code> ) <code>integer</code></td>
<td>Minimum scale (number of fractional decimal digits) needed to represent the supplied value precisely</td>
<td><code>min_scale(8.4100)</code> <code>2</code></td>
</tr>
<tr>
<td><code>mod</code> ( <code>y</code> <em>numeric_type</em>, <code>x</code> <em>numeric_type</em> ) <em>numeric_type</em></td>
<td>Remainder of <code>y</code>/<code>x</code>; available for <code>smallint</code>, <code>integer</code>, <code>bigint</code>, and <code>numeric</code></td>
<td><code>mod(9, 4)</code> <code>1</code></td>
</tr>
<tr>
<td><code>pi</code> ( ) <code>double precision</code></td>
<td>Approximate value of π</td>
<td><code>pi()</code> <code>3.141592653589793</code></td>
</tr>
<tr>
<td><code>power</code> ( <code>a</code> <code>numeric</code>, <code>b</code> <code>numeric</code> ) <code>numeric</code></td>
<td><code>power</code> ( <code>a</code> <code>double precision</code>, <code>b</code> <code>double precision</code> ) <code>double precision</code></td>
<td><code>a</code> raised to the power of <code>b</code><br><code>power(9, 3)</code> <code>729</code></td>
</tr>
<tr>
<td><code>radians</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Converts degrees to radians</td>
<td><code>radians(45.0)</code> <code>0.7853981633974483</code></td>
</tr>
<tr>
<td><code>round</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>round</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Rounds to nearest integer. For <code>numeric</code>, ties are broken by rounding away from zero. For <code>double precision</code>, the tie-breaking behavior is platform dependent, but “round to nearest even” is the most common rule.<br><code>round(42.4)</code> <code>42</code></td>
</tr>
<tr>
<td><code>round</code> ( <code>v</code> <code>numeric</code>, <code>s</code> <code>integer</code> ) <code>numeric</code></td>
<td>Rounds <code>v</code> to <code>s</code> decimal places. Ties are broken by rounding away from zero.</td>
<td><code>round(42.4382, 2)</code> <code>42.44</code><br><code>round(1234.56, -1)</code> <code>1230</code></td>
</tr>
<tr>
<td><code>scale</code> ( <code>numeric</code> ) <code>integer</code></td>
<td>Scale of the argument (the number of decimal digits in the fractional part)</td>
<td><code>scale(8.4100)</code> <code>4</code></td>
</tr>
<tr>
<td><code>sign</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>sign</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Sign of the argument (-1, 0, or +1)<br><code>sign(-8.4)</code> <code>-1</code></td>
</tr>
<tr>
<td><code>sqrt</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>sqrt</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Square root<br><code>sqrt(2)</code> <code>1.4142135623730951</code></td>
</tr>
<tr>
<td><code>trim_scale</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td>Reduces the value's scale (number of fractional decimal digits) by removing trailing zeroes</td>
<td><code>trim_scale(8.4100)</code> <code>8.41</code></td>
</tr>
<tr>
<td><code>trunc</code> ( <code>numeric</code> ) <code>numeric</code></td>
<td><code>trunc</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Truncates to integer (towards zero)<br><code>trunc(42.8)</code> <code>42</code><br><code>trunc(-42.8)</code> <code>-42</code></td>
</tr>
<tr>
<td><code>trunc</code> ( <code>v</code> <code>numeric</code>, <code>s</code> <code>integer</code> ) <code>numeric</code></td>
<td>Truncates <code>v</code> to <code>s</code> decimal places</td>
<td><code>trunc(42.4382, 2)</code> <code>42.43</code></td>
</tr>
<tr>
<td><code>width_bucket</code> ( <code>operand</code> <code>numeric</code>, <code>low</code> <code>numeric</code>, <code>high</code> <code>numeric</code>, <code>count</code> <code>integer</code> ) <code>integer</code></td>
<td><code>width_bucket</code> ( <code>operand</code> <code>double precision</code>, <code>low</code> <code>double precision</code>, <code>high</code> <code>double precision</code>, <code>count</code> <code>integer</code> ) <code>integer</code></td>
<td>Returns the number of the bucket in which <code>operand</code> falls in a histogram having <code>count</code> equal-width buckets spanning the range <code>low</code> to <code>high</code>. The buckets have inclusive lower bounds and exclusive upper bounds. Returns <code>0</code> for an input less than <code>low</code>, or `<code>count</code>+1<code> for an input greater than or equal to </code>high<code>. If </code>low<code> &gt; </code>high<code>, the behavior is mirror-reversed, with bucket </code>1<code> now being the one just below </code>low`, and the inclusive bounds now being on the upper side.<br><code>width_bucket(5.35, 0.024, 10.06, 5)</code> <code>3</code><br><code>width_bucket(9, 10, 0, 10)</code> <code>2</code></td>
</tr>
<tr>
<td><code>width_bucket</code> ( <code>operand</code> <code>anycompatible</code>, <code>thresholds</code> <code>anycompatiblearray</code> ) <code>integer</code></td>
<td>Returns the number of the bucket in which <code>operand</code> falls given an array listing the inclusive lower bounds of the buckets. Returns <code>0</code> for an input less than the first lower bound. <code>operand</code> and the array elements can be of any type having standard comparison operators. The <code>thresholds</code> array <em>must be sorted</em>, smallest first, or unexpected results will be obtained.</td>
<td><code>width_bucket(now(), array['yesterday', 'today', 'tomorrow']::timestamptz[])</code> <code>2</code></td>
</tr>
</tbody>
</table>


 [Random Functions](#functions-math-random-table) shows functions for generating random numbers.
 <a id="functions-math-random-table"></a>

**Table: Random Functions**

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
<td><code>random</code> ( ) <code>double precision</code></td>
<td>Returns a random value in the range 0.0 <= x < 1.0</td>
<td><code>random()</code> <code>0.897124072839091</code></td>
</tr>
<tr>
<td><code>random</code> ( <code>min</code> <code>integer</code>, <code>max</code> <code>integer</code> ) <code>integer</code></td>
<td><code>random</code> ( <code>min</code> <code>bigint</code>, <code>max</code> <code>bigint</code> ) <code>bigint</code></td>
<td><code>random</code> ( <code>min</code> <code>numeric</code>, <code>max</code> <code>numeric</code> ) <code>numeric</code><br>Returns a random value in the range <code>min</code> <= x <= <code>max</code>. For type <code>numeric</code>, the result will have the same number of fractional decimal digits as <code>min</code> or <code>max</code>, whichever has more.<br><code>random(1, 10)</code> <code>7</code><br><code>random(-0.499, 0.499)</code> <code>0.347</code></td>
</tr>
<tr>
<td><code>random_normal</code> ( [ <code>mean</code> <code>double precision</code> [, <code>stddev</code> <code>double precision</code> ]] ) <code>double precision</code></td>
<td>Returns a random value from the normal distribution with the given parameters; <code>mean</code> defaults to 0.0 and <code>stddev</code> defaults to 1.0</td>
<td><code>random_normal(0.0, 1.0)</code> <code>0.051285419</code></td>
</tr>
<tr>
<td><a id="function-setseed"></a>
 `setseed` ( `double precision` ) `void`</td>
<td>Sets the seed for subsequent <code>random()</code> and <code>random_normal()</code> calls; argument must be between -1.0 and 1.0, inclusive</td>
<td><code>setseed(0.12345)</code></td>
</tr>
</tbody>
</table>


 The `random()` and `random_normal()` functions listed in [Random Functions](#functions-math-random-table) and [Date/Time Functions](date-time-functions-and-operators.md#functions-datetime-table) use a deterministic pseudo-random number generator. It is fast but not suitable for cryptographic applications; see the [pgcrypto](../../appendixes/additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto) module for a more secure alternative. If `setseed()` is called, the series of results of subsequent calls to these functions in the current session can be repeated by re-issuing `setseed()` with the same argument. Without any prior `setseed()` call in the same session, the first call to any of these functions obtains a seed from a platform-dependent source of random bits.


 [Trigonometric Functions](#functions-math-trig-table) shows the available trigonometric functions. Each of these functions comes in two variants, one that measures angles in radians and one that measures angles in degrees.
 <a id="functions-math-trig-table"></a>

**Table: Trigonometric Functions**

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
<td><code>acos</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse cosine, result in radians</td>
<td><code>acos(1)</code> <code>0</code></td>
</tr>
<tr>
<td><code>acosd</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse cosine, result in degrees</td>
<td><code>acosd(0.5)</code> <code>60</code></td>
</tr>
<tr>
<td><code>asin</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse sine, result in radians</td>
<td><code>asin(1)</code> <code>1.5707963267948966</code></td>
</tr>
<tr>
<td><code>asind</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse sine, result in degrees</td>
<td><code>asind(0.5)</code> <code>30</code></td>
</tr>
<tr>
<td><code>atan</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse tangent, result in radians</td>
<td><code>atan(1)</code> <code>0.7853981633974483</code></td>
</tr>
<tr>
<td><code>atand</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse tangent, result in degrees</td>
<td><code>atand(1)</code> <code>45</code></td>
</tr>
<tr>
<td><code>atan2</code> ( <code>y</code> <code>double precision</code>, <code>x</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse tangent of <code>y</code>/<code>x</code>, result in radians</td>
<td><code>atan2(1, 0)</code> <code>1.5707963267948966</code></td>
</tr>
<tr>
<td><code>atan2d</code> ( <code>y</code> <code>double precision</code>, <code>x</code> <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse tangent of <code>y</code>/<code>x</code>, result in degrees</td>
<td><code>atan2d(1, 0)</code> <code>90</code></td>
</tr>
<tr>
<td><code>cos</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Cosine, argument in radians</td>
<td><code>cos(0)</code> <code>1</code></td>
</tr>
<tr>
<td><code>cosd</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Cosine, argument in degrees</td>
<td><code>cosd(60)</code> <code>0.5</code></td>
</tr>
<tr>
<td><code>cot</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Cotangent, argument in radians</td>
<td><code>cot(0.5)</code> <code>1.830487721712452</code></td>
</tr>
<tr>
<td><code>cotd</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Cotangent, argument in degrees</td>
<td><code>cotd(45)</code> <code>1</code></td>
</tr>
<tr>
<td><code>sin</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Sine, argument in radians</td>
<td><code>sin(1)</code> <code>0.8414709848078965</code></td>
</tr>
<tr>
<td><code>sind</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Sine, argument in degrees</td>
<td><code>sind(30)</code> <code>0.5</code></td>
</tr>
<tr>
<td><code>tan</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Tangent, argument in radians</td>
<td><code>tan(1)</code> <code>1.5574077246549023</code></td>
</tr>
<tr>
<td><code>tand</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Tangent, argument in degrees</td>
<td><code>tand(45)</code> <code>1</code></td>
</tr>
</tbody>
</table>


!!! note

    Another way to work with angles measured in degrees is to use the unit transformation functions ``radians()`` and ``degrees()`` shown earlier. However, using the degree-based trigonometric functions is preferred, as that way avoids round-off error for special cases such as `sind(30)`.


 [Hyperbolic Functions](#functions-math-hyp-table) shows the available hyperbolic functions.
 <a id="functions-math-hyp-table"></a>

**Table: Hyperbolic Functions**

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
<td><code>sinh</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Hyperbolic sine</td>
<td><code>sinh(1)</code> <code>1.1752011936438014</code></td>
</tr>
<tr>
<td><code>cosh</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Hyperbolic cosine</td>
<td><code>cosh(0)</code> <code>1</code></td>
</tr>
<tr>
<td><code>tanh</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Hyperbolic tangent</td>
<td><code>tanh(1)</code> <code>0.7615941559557649</code></td>
</tr>
<tr>
<td><code>asinh</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse hyperbolic sine</td>
<td><code>asinh(1)</code> <code>0.881373587019543</code></td>
</tr>
<tr>
<td><code>acosh</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse hyperbolic cosine</td>
<td><code>acosh(1)</code> <code>0</code></td>
</tr>
<tr>
<td><code>atanh</code> ( <code>double precision</code> ) <code>double precision</code></td>
<td>Inverse hyperbolic tangent</td>
<td><code>atanh(0.5)</code> <code>0.5493061443340548</code></td>
</tr>
</tbody>
</table>
