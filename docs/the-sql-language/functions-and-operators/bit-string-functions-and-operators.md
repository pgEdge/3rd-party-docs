<a id="functions-bitstring"></a>

## Bit String Functions and Operators


 This section describes functions and operators for examining and manipulating bit strings, that is values of the types `bit` and `bit varying`. (While only type `bit` is mentioned in these tables, values of type `bit varying` can be used interchangeably.) Bit strings support the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table), as well as the operators shown in [Bit String Operators](#functions-bit-string-op-table).
 <a id="functions-bit-string-op-table"></a>

**Table: Bit String Operators**

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
<td><code>bit</code> <code>||</code> <code>bit</code> <code>bit</code></td>
<td>Concatenation</td>
<td><code>B'10001' || B'011'</code> <code>10001011</code></td>
</tr>
<tr>
<td><code>bit</code> <code>&amp;</code> <code>bit</code> <code>bit</code></td>
<td>Bitwise AND (inputs must be of equal length)</td>
<td><code>B'10001' &amp; B'01101'</code> <code>00001</code></td>
</tr>
<tr>
<td><code>bit</code> <code>|</code> <code>bit</code> <code>bit</code></td>
<td>Bitwise OR (inputs must be of equal length)</td>
<td><code>B'10001' | B'01101'</code> <code>11101</code></td>
</tr>
<tr>
<td><code>bit</code> <code>#</code> <code>bit</code> <code>bit</code></td>
<td>Bitwise exclusive OR (inputs must be of equal length)</td>
<td><code>B'10001' # B'01101'</code> <code>11100</code></td>
</tr>
<tr>
<td><code>~</code> <code>bit</code> <code>bit</code></td>
<td>Bitwise NOT</td>
<td><code>~ B'10001'</code> <code>01110</code></td>
</tr>
<tr>
<td><code>bit</code> <code>&lt;&lt;</code> <code>integer</code> <code>bit</code></td>
<td>Bitwise shift left (string length is preserved)</td>
<td><code>B'10001' &lt;&lt; 3</code> <code>01000</code></td>
</tr>
<tr>
<td><code>bit</code> <code>&gt;&gt;</code> <code>integer</code> <code>bit</code></td>
<td>Bitwise shift right (string length is preserved)</td>
<td><code>B'10001' &gt;&gt; 2</code> <code>00100</code></td>
</tr>
</tbody>
</table>


 Some of the functions available for binary strings are also available for bit strings, as shown in [Bit String Functions](#functions-bit-string-table).
 <a id="functions-bit-string-table"></a>

**Table: Bit String Functions**

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
<td><code>bit_count</code> ( <code>bit</code> ) <code>bigint</code></td>
<td>Returns the number of bits set in the bit string (also known as “popcount”).</td>
<td><code>bit_count(B'10111')</code> <code>4</code></td>
</tr>
<tr>
<td><code>bit_length</code> ( <code>bit</code> ) <code>integer</code></td>
<td>Returns number of bits in the bit string.</td>
<td><code>bit_length(B'10111')</code> <code>5</code></td>
</tr>
<tr>
<td><code>length</code> ( <code>bit</code> ) <code>integer</code></td>
<td>Returns number of bits in the bit string.</td>
<td><code>length(B'10111')</code> <code>5</code></td>
</tr>
<tr>
<td><code>octet_length</code> ( <code>bit</code> ) <code>integer</code></td>
<td>Returns number of bytes in the bit string.</td>
<td><code>octet_length(B'1011111011')</code> <code>2</code></td>
</tr>
<tr>
<td><code>overlay</code> ( <code>bits</code> <code>bit</code> <code>PLACING</code> <code>newsubstring</code> <code>bit</code> <code>FROM</code> <code>start</code> <code>integer</code> [ <code>FOR</code> <code>count</code> <code>integer</code> ] ) <code>bit</code></td>
<td>Replaces the substring of <code>bits</code> that starts at the <code>start</code>'th bit and extends for <code>count</code> bits with <code>newsubstring</code>. If <code>count</code> is omitted, it defaults to the length of <code>newsubstring</code>.</td>
<td><code>overlay(B'01010101010101010' PLACING B'11111' FROM 2 FOR 3)</code> <code>0111110101010101010</code></td>
</tr>
<tr>
<td><code>position</code> ( <code>substring</code> <code>bit</code> <code>IN</code> <code>bits</code> <code>bit</code> ) <code>integer</code></td>
<td>Returns first starting index of the specified <code>substring</code> within <code>bits</code>, or zero if it's not present.</td>
<td><code>position(B'010' IN B'000001101011')</code> <code>8</code></td>
</tr>
<tr>
<td><code>substring</code> ( <code>bits</code> <code>bit</code> [ <code>FROM</code> <code>start</code> <code>integer</code> ] [ <code>FOR</code> <code>count</code> <code>integer</code> ] ) <code>bit</code></td>
<td>Extracts the substring of <code>bits</code> starting at the <code>start</code>'th bit if that is specified, and stopping after <code>count</code> bits if that is specified. Provide at least one of <code>start</code> and <code>count</code>.</td>
<td><code>substring(B'110010111111' FROM 3 FOR 2)</code> <code>00</code></td>
</tr>
<tr>
<td><code>get_bit</code> ( <code>bits</code> <code>bit</code>, <code>n</code> <code>integer</code> ) <code>integer</code></td>
<td>Extracts <code>n</code>'th bit from bit string; the first (leftmost) bit is bit 0.</td>
<td><code>get_bit(B'101010101010101010', 6)</code> <code>1</code></td>
</tr>
<tr>
<td><code>set_bit</code> ( <code>bits</code> <code>bit</code>, <code>n</code> <code>integer</code>, <code>newvalue</code> <code>integer</code> ) <code>bit</code></td>
<td>Sets <code>n</code>'th bit in bit string to <code>newvalue</code>; the first (leftmost) bit is bit 0.</td>
<td><code>set_bit(B'101010101010101010', 6, 0)</code> <code>101010001010101010</code></td>
</tr>
</tbody>
</table>


 In addition, it is possible to cast integral values to and from type `bit`. Casting an integer to `bit(n)` copies the rightmost `n` bits. Casting an integer to a bit string width wider than the integer itself will sign-extend on the left. Some examples:

```

44::bit(10)                    0000101100
44::bit(3)                     100
cast(-44 AS bit(12))           111111010100
'1110'::bit(4)::integer        14
```
 Note that casting to just “bit” means casting to `bit(1)`, and so will deliver only the least significant bit of the integer.
