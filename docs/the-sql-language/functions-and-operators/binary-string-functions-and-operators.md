<a id="functions-binarystring"></a>

## Binary String Functions and Operators


 This section describes functions and operators for examining and manipulating binary strings, that is values of type `bytea`. Many of these are equivalent, in purpose and syntax, to the text-string functions described in the previous section.


 SQL defines some string functions that use key words, rather than commas, to separate arguments. Details are in [SQL Binary String Functions and Operators](#functions-binarystring-sql). PostgreSQL also provides versions of these functions that use the regular function invocation syntax (see [Other Binary String Functions](#functions-binarystring-other)).
 <a id="functions-binarystring-sql"></a>

**Table: SQL Binary String Functions and Operators**

<table>
<thead>
<tr>
<th>Function/Operator</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>bytea</code> <code>||</code> <code>bytea</code> <code>bytea</code></td>
<td>Concatenates the two binary strings.</td>
<td><code>'\x123456'::bytea || '\x789a00bcde'::bytea</code> <code>\x123456789a00bcde</code></td>
</tr>
<tr>
<td><code>bit_length</code> ( <code>bytea</code> ) <code>integer</code></td>
<td>Returns number of bits in the binary string (8 times the <code>octet_length</code>).</td>
<td><code>bit_length('\x123456'::bytea)</code> <code>24</code></td>
</tr>
<tr>
<td><code>btrim</code> ( <code>bytes</code> <code>bytea</code>, <code>bytesremoved</code> <code>bytea</code> ) <code>bytea</code></td>
<td>Removes the longest string containing only bytes appearing in <code>bytesremoved</code> from the start and end of <code>bytes</code>.</td>
<td><code>btrim('\x1234567890'::bytea, '\x9012'::bytea)</code> <code>\x345678</code></td>
</tr>
<tr>
<td><code>ltrim</code> ( <code>bytes</code> <code>bytea</code>, <code>bytesremoved</code> <code>bytea</code> ) <code>bytea</code></td>
<td>Removes the longest string containing only bytes appearing in <code>bytesremoved</code> from the start of <code>bytes</code>.</td>
<td><code>ltrim('\x1234567890'::bytea, '\x9012'::bytea)</code> <code>\x34567890</code></td>
</tr>
<tr>
<td><code>octet_length</code> ( <code>bytea</code> ) <code>integer</code></td>
<td>Returns number of bytes in the binary string.</td>
<td><code>octet_length('\x123456'::bytea)</code> <code>3</code></td>
</tr>
<tr>
<td><code>overlay</code> ( <code>bytes</code> <code>bytea</code> <code>PLACING</code> <code>newsubstring</code> <code>bytea</code> <code>FROM</code> <code>start</code> <code>integer</code> [ <code>FOR</code> <code>count</code> <code>integer</code> ] ) <code>bytea</code></td>
<td>Replaces the substring of <code>bytes</code> that starts at the <code>start</code>'th byte and extends for <code>count</code> bytes with <code>newsubstring</code>. If <code>count</code> is omitted, it defaults to the length of <code>newsubstring</code>.</td>
<td><code>overlay('\x1234567890'::bytea placing '\002\003'::bytea from 2 for 3)</code> <code>\x12020390</code></td>
</tr>
<tr>
<td><code>position</code> ( <code>substring</code> <code>bytea</code> <code>IN</code> <code>bytes</code> <code>bytea</code> ) <code>integer</code></td>
<td>Returns first starting index of the specified <code>substring</code> within <code>bytes</code>, or zero if it's not present.</td>
<td><code>position('\x5678'::bytea in '\x1234567890'::bytea)</code> <code>3</code></td>
</tr>
<tr>
<td><code>rtrim</code> ( <code>bytes</code> <code>bytea</code>, <code>bytesremoved</code> <code>bytea</code> ) <code>bytea</code></td>
<td>Removes the longest string containing only bytes appearing in <code>bytesremoved</code> from the end of <code>bytes</code>.</td>
<td><code>rtrim('\x1234567890'::bytea, '\x9012'::bytea)</code> <code>\x12345678</code></td>
</tr>
<tr>
<td><code>substring</code> ( <code>bytes</code> <code>bytea</code> [ <code>FROM</code> <code>start</code> <code>integer</code> ] [ <code>FOR</code> <code>count</code> <code>integer</code> ] ) <code>bytea</code></td>
<td>Extracts the substring of <code>bytes</code> starting at the <code>start</code>'th byte if that is specified, and stopping after <code>count</code> bytes if that is specified. Provide at least one of <code>start</code> and <code>count</code>.</td>
<td><code>substring('\x1234567890'::bytea from 3 for 2)</code> <code>\x5678</code></td>
</tr>
<tr>
<td><code>trim</code> ( [ <code>LEADING</code> | <code>TRAILING</code> | <code>BOTH</code> ] <code>bytesremoved</code> <code>bytea</code> <code>FROM</code> <code>bytes</code> <code>bytea</code> ) <code>bytea</code></td>
<td>Removes the longest string containing only bytes appearing in <code>bytesremoved</code> from the start, end, or both ends (<code>BOTH</code> is the default) of <code>bytes</code>.</td>
<td><code>trim('\x9012'::bytea from '\x1234567890'::bytea)</code> <code>\x345678</code></td>
</tr>
<tr>
<td><code>trim</code> ( [ <code>LEADING</code> | <code>TRAILING</code> | <code>BOTH</code> ] [ <code>FROM</code> ] <code>bytes</code> <code>bytea</code>, <code>bytesremoved</code> <code>bytea</code> ) <code>bytea</code></td>
<td>This is a non-standard syntax for <code>trim()</code>.</td>
<td><code>trim(both from '\x1234567890'::bytea, '\x9012'::bytea)</code> <code>\x345678</code></td>
</tr>
</tbody>
</table>


 Additional binary string manipulation functions are available and are listed in [Other Binary String Functions](#functions-binarystring-other). Some of them are used internally to implement the SQL-standard string functions listed in [SQL Binary String Functions and Operators](#functions-binarystring-sql).
 <a id="functions-binarystring-other"></a>

**Table: Other Binary String Functions**

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
<td><code>bit_count</code> ( <code>bytes</code> <code>bytea</code> ) <code>bigint</code></td>
<td>Returns the number of bits set in the binary string (also known as “popcount”).</td>
<td><code>bit_count('\x1234567890'::bytea)</code> <code>15</code></td>
</tr>
<tr>
<td><code>get_bit</code> ( <code>bytes</code> <code>bytea</code>, <code>n</code> <code>bigint</code> ) <code>integer</code></td>
<td>Extracts <a href="#functions-zerobased-note">n'th</a> bit from binary string.</td>
<td><code>get_bit('\x1234567890'::bytea, 30)</code> <code>1</code></td>
</tr>
<tr>
<td><code>get_byte</code> ( <code>bytes</code> <code>bytea</code>, <code>n</code> <code>integer</code> ) <code>integer</code></td>
<td>Extracts <a href="#functions-zerobased-note">n'th</a> byte from binary string.</td>
<td><code>get_byte('\x1234567890'::bytea, 4)</code> <code>144</code></td>
</tr>
<tr>
<td><code>length</code> ( <code>bytea</code> ) <code>integer</code></td>
<td>Returns the number of bytes in the binary string.</td>
<td><code>length('\x1234567890'::bytea)</code> <code>5</code></td>
</tr>
<tr>
<td><code>length</code> ( <code>bytes</code> <code>bytea</code>, <code>encoding</code> <code>name</code> ) <code>integer</code></td>
<td>Returns the number of characters in the binary string, assuming that it is text in the given <code>encoding</code>.</td>
<td><code>length('jose'::bytea, 'UTF8')</code> <code>4</code></td>
</tr>
<tr>
<td><code>md5</code> ( <code>bytea</code> ) <code>text</code></td>
<td>Computes the MD5 <a href="#functions-hash-note">hash</a> of the binary string, with the result written in hexadecimal.</td>
<td><code>md5('Th\000omas'::bytea)</code> <code>8ab2d3c9689aaf18​b4958c334c82d8b1</code></td>
</tr>
<tr>
<td><code>set_bit</code> ( <code>bytes</code> <code>bytea</code>, <code>n</code> <code>bigint</code>, <code>newvalue</code> <code>integer</code> ) <code>bytea</code></td>
<td>Sets <a href="#functions-zerobased-note">n'th</a> bit in binary string to <code>newvalue</code>.</td>
<td><code>set_bit('\x1234567890'::bytea, 30, 0)</code> <code>\x1234563890</code></td>
</tr>
<tr>
<td><code>set_byte</code> ( <code>bytes</code> <code>bytea</code>, <code>n</code> <code>integer</code>, <code>newvalue</code> <code>integer</code> ) <code>bytea</code></td>
<td>Sets <a href="#functions-zerobased-note">n'th</a> byte in binary string to <code>newvalue</code>.</td>
<td><code>set_byte('\x1234567890'::bytea, 4, 64)</code> <code>\x1234567840</code></td>
</tr>
<tr>
<td><code>sha224</code> ( <code>bytea</code> ) <code>bytea</code></td>
<td>Computes the SHA-224 <a href="#functions-hash-note">hash</a> of the binary string.</td>
<td><code>sha224('abc'::bytea)</code> <code>\x23097d223405d8228642a477bda2​55b32aadbce4bda0b3f7e36c9da7</code></td>
</tr>
<tr>
<td><code>sha256</code> ( <code>bytea</code> ) <code>bytea</code></td>
<td>Computes the SHA-256 <a href="#functions-hash-note">hash</a> of the binary string.</td>
<td><code>sha256('abc'::bytea)</code> <code>\xba7816bf8f01cfea414140de5dae2223​b00361a396177a9cb410ff61f20015ad</code></td>
</tr>
<tr>
<td><code>sha384</code> ( <code>bytea</code> ) <code>bytea</code></td>
<td>Computes the SHA-384 <a href="#functions-hash-note">hash</a> of the binary string.</td>
<td><code>sha384('abc'::bytea)</code> <code>\xcb00753f45a35e8bb5a03d699ac65007​272c32ab0eded1631a8b605a43ff5bed​8086072ba1e7cc2358baeca134c825a7</code></td>
</tr>
<tr>
<td><code>sha512</code> ( <code>bytea</code> ) <code>bytea</code></td>
<td>Computes the SHA-512 <a href="#functions-hash-note">hash</a> of the binary string.</td>
<td><code>sha512('abc'::bytea)</code> <code>\xddaf35a193617abacc417349ae204131​12e6fa4e89a97ea20a9eeee64b55d39a​2192992a274fc1a836ba3c23a3feebbd​454d4423643ce80e2a9ac94fa54ca49f</code></td>
</tr>
<tr>
<td><code>substr</code> ( <code>bytes</code> <code>bytea</code>, <code>start</code> <code>integer</code> [, <code>count</code> <code>integer</code> ] ) <code>bytea</code></td>
<td>Extracts the substring of <code>bytes</code> starting at the <code>start</code>'th byte, and extending for <code>count</code> bytes if that is specified. (Same as <code>substring(</code>bytes<code> from </code>start<code> for </code>count<code>)</code>.)</td>
<td><code>substr('\x1234567890'::bytea, 3, 2)</code> <code>\x5678</code></td>
</tr>
</tbody>
</table>
 <a id="functions-zerobased-note"></a>

 Functions `get_byte` and `set_byte` number the first byte of a binary string as byte 0. Functions `get_bit` and `set_bit` number bits from the right within each byte; for example bit 0 is the least significant bit of the first byte, and bit 15 is the most significant bit of the second byte.
 <a id="functions-hash-note"></a>

 For historical reasons, the function `md5` returns a hex-encoded value of type `text` whereas the SHA-2 functions return type `bytea`. Use the functions [`encode`](#function-encode) and [`decode`](#function-decode) to convert between the two. For example write `encode(sha256('abc'), 'hex')` to get a hex-encoded text representation, or `decode(md5('abc'), 'hex')` to get a `bytea` value.


   Functions for converting strings between different character sets (encodings), and for representing arbitrary binary data in textual form, are shown in [Text/Binary String Conversion Functions](#functions-binarystring-conversions). For these functions, an argument or result of type `text` is expressed in the database's default encoding, while arguments or results of type `bytea` are in an encoding named by another argument.
 <a id="functions-binarystring-conversions"></a>

**Table: Text/Binary String Conversion Functions**

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
<td><code>convert</code> ( <code>bytes</code> <code>bytea</code>, <code>src_encoding</code> <code>name</code>, <code>dest_encoding</code> <code>name</code> ) <code>bytea</code></td>
<td>Converts a binary string representing text in encoding <code>src_encoding</code> to a binary string in encoding <code>dest_encoding</code> (see <a href="../../server-administration/localization/character-set-support.md#multibyte-conversions-supported">Available Character Set Conversions</a> for available conversions).</td>
<td><code>convert('text_in_utf8', 'UTF8', 'LATIN1')</code> <code>\x746578745f696e5f75746638</code></td>
</tr>
<tr>
<td><code>convert_from</code> ( <code>bytes</code> <code>bytea</code>, <code>src_encoding</code> <code>name</code> ) <code>text</code></td>
<td>Converts a binary string representing text in encoding <code>src_encoding</code> to <code>text</code> in the database encoding (see <a href="../../server-administration/localization/character-set-support.md#multibyte-conversions-supported">Available Character Set Conversions</a> for available conversions).</td>
<td><code>convert_from('text_in_utf8', 'UTF8')</code> <code>text_in_utf8</code></td>
</tr>
<tr>
<td><code>convert_to</code> ( <code>string</code> <code>text</code>, <code>dest_encoding</code> <code>name</code> ) <code>bytea</code></td>
<td>Converts a <code>text</code> string (in the database encoding) to a binary string encoded in encoding <code>dest_encoding</code> (see <a href="../../server-administration/localization/character-set-support.md#multibyte-conversions-supported">Available Character Set Conversions</a> for available conversions).</td>
<td><code>convert_to('some_text', 'UTF8')</code> <code>\x736f6d655f74657874</code></td>
</tr>
<tr>
<td><a id="function-encode"></a>
 `encode` ( `bytes` `bytea`, `format` `text` ) `text`</td>
<td>Encodes binary data into a textual representation; supported <code>format</code> values are: <a href="#encode-format-base64"><code>base64</code></a>, <a href="#encode-format-escape"><code>escape</code></a>, <a href="#encode-format-hex"><code>hex</code></a>.</td>
<td><code>encode('123\000\001', 'base64')</code> <code>MTIzAAE=</code></td>
</tr>
<tr>
<td><a id="function-decode"></a>
 `decode` ( `string` `text`, `format` `text` ) `bytea`</td>
<td>Decodes binary data from a textual representation; supported <code>format</code> values are the same as for <code>encode</code>.</td>
<td><code>decode('MTIzAAE=', 'base64')</code> <code>\x3132330001</code></td>
</tr>
</tbody>
</table>


 The `encode` and `decode` functions support the following textual formats:

<a id="encode-format-base64"></a>

base64
:   The `base64` format is that of [RFC 2045 Section 6.8](https://datatracker.ietf.org/doc/html/rfc2045#section-6.8). As per the RFC, encoded lines are broken at 76 characters. However instead of the MIME CRLF end-of-line marker, only a newline is used for end-of-line. The `decode` function ignores carriage-return, newline, space, and tab characters. Otherwise, an error is raised when `decode` is supplied invalid base64 data — including when trailing padding is incorrect.
<a id="encode-format-escape"></a>

escape
:   The `escape` format converts zero bytes and bytes with the high bit set into octal escape sequences (`\`*nnn*), and it doubles backslashes. Other byte values are represented literally. The `decode` function will raise an error if a backslash is not followed by either a second backslash or three octal digits; it accepts other byte values unchanged.
<a id="encode-format-hex"></a>

hex
:   The `hex` format represents each 4 bits of data as one hexadecimal digit, `0` through `f`, writing the higher-order digit of each byte first. The `encode` function outputs the `a`-`f` hex digits in lower case. Because the smallest unit of data is 8 bits, there are always an even number of characters returned by `encode`. The `decode` function accepts the `a`-`f` characters in either upper or lower case. An error is raised when `decode` is given invalid hex data — including when given an odd number of characters.


 See also the aggregate function `string_agg` in [Aggregate Functions](aggregate-functions.md#functions-aggregate) and the large object functions in [Server-Side Functions](../../client-interfaces/large-objects/server-side-functions.md#lo-funcs).
