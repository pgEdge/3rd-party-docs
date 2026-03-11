## String Functions and Operators { #functions-string }


 This section describes functions and operators for examining and manipulating string values. Strings in this context include values of the types `character`, `character varying`, and `text`. Except where noted, these functions and operators are declared to accept and return type `text`. They will interchangeably accept `character varying` arguments. Values of type `character` will be converted to `text` before the function or operator is applied, resulting in stripping any trailing spaces in the `character` value.


 SQL defines some string functions that use key words, rather than commas, to separate arguments. Details are in [SQL String Functions and Operators](#functions-string-sql). PostgreSQL also provides versions of these functions that use the regular function invocation syntax (see [Other String Functions and Operators](#functions-string-other)).


!!! note

    The string concatenation operator (`||`) will accept non-string input, so long as at least one input is of string type, as shown in [SQL String Functions and Operators](#functions-string-sql). For other cases, inserting an explicit coercion to `text` can be used to have non-string input accepted.


<a id="functions-string-sql"></a>
**Table: SQL String Functions and Operators**

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
<td><code>text</code> <code>||</code> <code>text</code> <code>text</code></td>
<td>Concatenates the two strings.</td>
<td><code>'Post' || 'greSQL'</code> <code>PostgreSQL</code></td>
</tr>
<tr>
<td><code>text</code> <code>||</code> <code>anynonarray</code> <code>text</code></td>
<td><code>anynonarray</code> <code>||</code> <code>text</code> <code>text</code></td>
<td>Converts the non-string input to text, then concatenates the two strings. (The non-string input cannot be of an array type, because that would create ambiguity with the array <code>||</code> operators. If you want to concatenate an array's text equivalent, cast it to <code>text</code> explicitly.)<br><code>'Value: ' || 42</code> <code>Value: 42</code></td>
</tr>
<tr>
<td><code>btrim</code> ( <code>string</code> <code>text</code> [, <code>characters</code> <code>text</code> ] ) <code>text</code></td>
<td>Removes the longest string containing only characters in <code>characters</code> (a space by default) from the start and end of <code>string</code>.</td>
<td><code>btrim('xyxtrimyyx', 'xyz')</code> <code>trim</code></td>
</tr>
<tr>
<td><code>text</code> <code>IS</code> [<code>NOT</code>] [<code>form</code>] <code>NORMALIZED</code> <code>boolean</code></td>
<td>Checks whether the string is in the specified Unicode normalization form. The optional <code>form</code> key word specifies the form: <code>NFC</code> (the default), <code>NFD</code>, <code>NFKC</code>, or <code>NFKD</code>. This expression can only be used when the server encoding is <code>UTF8</code>. Note that checking for normalization using this expression is often faster than normalizing possibly already normalized strings.</td>
<td><code>U&amp;'\0061\0308bc' IS NFD NORMALIZED</code> <code>t</code></td>
</tr>
<tr>
<td><code>bit_length</code> ( <code>text</code> ) <code>integer</code></td>
<td>Returns number of bits in the string (8 times the <code>octet_length</code>).</td>
<td><code>bit_length('jose')</code> <code>32</code></td>
</tr>
<tr>
<td><code>char_length</code> ( <code>text</code> ) <code>integer</code></td>
<td><code>character_length</code> ( <code>text</code> ) <code>integer</code></td>
<td>Returns number of characters in the string.<br><code>char_length('jos&amp;eacute;')</code> <code>4</code></td>
</tr>
<tr>
<td><code>lower</code> ( <code>text</code> ) <code>text</code></td>
<td>Converts the string to all lower case, according to the rules of the database's locale.</td>
<td><code>lower('TOM')</code> <code>tom</code></td>
</tr>
<tr>
<td><code>lpad</code> ( <code>string</code> <code>text</code>, <code>length</code> <code>integer</code> [, <code>fill</code> <code>text</code> ] ) <code>text</code></td>
<td>Extends the <code>string</code> to length <code>length</code> by prepending the characters <code>fill</code> (a space by default). If the <code>string</code> is already longer than <code>length</code> then it is truncated (on the right).</td>
<td><code>lpad('hi', 5, 'xy')</code> <code>xyxhi</code></td>
</tr>
<tr>
<td><code>ltrim</code> ( <code>string</code> <code>text</code> [, <code>characters</code> <code>text</code> ] ) <code>text</code></td>
<td>Removes the longest string containing only characters in <code>characters</code> (a space by default) from the start of <code>string</code>.</td>
<td><code>ltrim('zzzytest', 'xyz')</code> <code>test</code></td>
</tr>
<tr>
<td><code>normalize</code> ( <code>text</code> [, <code>form</code> ] ) <code>text</code></td>
<td>Converts the string to the specified Unicode normalization form. The optional <code>form</code> key word specifies the form: <code>NFC</code> (the default), <code>NFD</code>, <code>NFKC</code>, or <code>NFKD</code>. This function can only be used when the server encoding is <code>UTF8</code>.</td>
<td><code>normalize(U&amp;'\0061\0308bc', NFC)</code> <code>U&amp;'\00E4bc'</code></td>
</tr>
<tr>
<td><code>octet_length</code> ( <code>text</code> ) <code>integer</code></td>
<td>Returns number of bytes in the string.</td>
<td><code>octet_length('jos&amp;eacute;')</code> <code>5</code> (if server encoding is UTF8)</td>
</tr>
<tr>
<td><code>octet_length</code> ( <code>character</code> ) <code>integer</code></td>
<td>Returns number of bytes in the string. Since this version of the function accepts type <code>character</code> directly, it will not strip trailing spaces.</td>
<td><code>octet_length('abc '::character(4))</code> <code>4</code></td>
</tr>
<tr>
<td><code>overlay</code> ( <code>string</code> <code>text</code> <code>PLACING</code> <code>newsubstring</code> <code>text</code> <code>FROM</code> <code>start</code> <code>integer</code> [ <code>FOR</code> <code>count</code> <code>integer</code> ] ) <code>text</code></td>
<td>Replaces the substring of <code>string</code> that starts at the <code>start</code>'th character and extends for <code>count</code> characters with <code>newsubstring</code>. If <code>count</code> is omitted, it defaults to the length of <code>newsubstring</code>.</td>
<td><code>overlay('Txxxxas' PLACING 'hom' FROM 2 FOR 4)</code> <code>Thomas</code></td>
</tr>
<tr>
<td><code>position</code> ( <code>substring</code> <code>text</code> <code>IN</code> <code>string</code> <code>text</code> ) <code>integer</code></td>
<td>Returns first starting index of the specified <code>substring</code> within <code>string</code>, or zero if it's not present.</td>
<td><code>position('om' IN 'Thomas')</code> <code>3</code></td>
</tr>
<tr>
<td><code>rpad</code> ( <code>string</code> <code>text</code>, <code>length</code> <code>integer</code> [, <code>fill</code> <code>text</code> ] ) <code>text</code></td>
<td>Extends the <code>string</code> to length <code>length</code> by appending the characters <code>fill</code> (a space by default). If the <code>string</code> is already longer than <code>length</code> then it is truncated.</td>
<td><code>rpad('hi', 5, 'xy')</code> <code>hixyx</code></td>
</tr>
<tr>
<td><code>rtrim</code> ( <code>string</code> <code>text</code> [, <code>characters</code> <code>text</code> ] ) <code>text</code></td>
<td>Removes the longest string containing only characters in <code>characters</code> (a space by default) from the end of <code>string</code>.</td>
<td><code>rtrim('testxxzx', 'xyz')</code> <code>test</code></td>
</tr>
<tr>
<td><code>substring</code> ( <code>string</code> <code>text</code> [ <code>FROM</code> <code>start</code> <code>integer</code> ] [ <code>FOR</code> <code>count</code> <code>integer</code> ] ) <code>text</code></td>
<td>Extracts the substring of <code>string</code> starting at the <code>start</code>'th character if that is specified, and stopping after <code>count</code> characters if that is specified. Provide at least one of <code>start</code> and <code>count</code>.</td>
<td><code>substring('Thomas' FROM 2 FOR 3)</code> <code>hom</code><br><code>substring('Thomas' FROM 3)</code> <code>omas</code><br><code>substring('Thomas' FOR 2)</code> <code>Th</code></td>
</tr>
<tr>
<td><code>substring</code> ( <code>string</code> <code>text</code> <code>FROM</code> <code>pattern</code> <code>text</code> ) <code>text</code></td>
<td>Extracts the first substring matching POSIX regular expression; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>substring('Thomas' FROM '...$')</code> <code>mas</code></td>
</tr>
<tr>
<td><code>substring</code> ( <code>string</code> <code>text</code> <code>SIMILAR</code> <code>pattern</code> <code>text</code> <code>ESCAPE</code> <code>escape</code> <code>text</code> ) <code>text</code></td>
<td><code>substring</code> ( <code>string</code> <code>text</code> <code>FROM</code> <code>pattern</code> <code>text</code> <code>FOR</code> <code>escape</code> <code>text</code> ) <code>text</code></td>
<td>Extracts the first substring matching SQL regular expression; see <a href="pattern-matching.md#functions-similarto-regexp"><code>SIMILAR TO</code> Regular Expressions</a>. The first form has been specified since SQL:2003; the second form was only in SQL:1999 and should be considered obsolete.<br><code>substring('Thomas' SIMILAR '%#"o_a#"_' ESCAPE '#')</code> <code>oma</code></td>
</tr>
<tr>
<td><code>trim</code> ( [ <code>LEADING</code> | <code>TRAILING</code> | <code>BOTH</code> ] [ <code>characters</code> <code>text</code> ] <code>FROM</code> <code>string</code> <code>text</code> ) <code>text</code></td>
<td>Removes the longest string containing only characters in <code>characters</code> (a space by default) from the start, end, or both ends (<code>BOTH</code> is the default) of <code>string</code>.</td>
<td><code>trim(both 'xyz' from 'yxTomxx')</code> <code>Tom</code></td>
</tr>
<tr>
<td><code>trim</code> ( [ <code>LEADING</code> | <code>TRAILING</code> | <code>BOTH</code> ] [ <code>FROM</code> ] <code>string</code> <code>text</code> [, <code>characters</code> <code>text</code> ] ) <code>text</code></td>
<td>This is a non-standard syntax for <code>trim()</code>.</td>
<td><code>trim(both from 'yxTomxx', 'xyz')</code> <code>Tom</code></td>
</tr>
<tr>
<td><code>unicode_assigned</code> ( <code>text</code> ) <code>boolean</code></td>
<td>Returns <code>true</code> if all characters in the string are assigned Unicode codepoints; <code>false</code> otherwise. This function can only be used when the server encoding is <code>UTF8</code>.</td>
<td></td>
</tr>
<tr>
<td><code>upper</code> ( <code>text</code> ) <code>text</code></td>
<td>Converts the string to all upper case, according to the rules of the database's locale.</td>
<td><code>upper('tom')</code> <code>TOM</code></td>
</tr>
</tbody>
</table>


 Additional string manipulation functions and operators are available and are listed in [Other String Functions and Operators](#functions-string-other). (Some of these are used internally to implement the SQL-standard string functions listed in [SQL String Functions and Operators](#functions-string-sql).) There are also pattern-matching operators, which are described in [Pattern Matching](pattern-matching.md#functions-matching), and operators for full-text search, which are described in [Full Text Search](../full-text-search/index.md#textsearch).


<a id="functions-string-other"></a>
**Table: Other String Functions and Operators**

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
<td><code>text</code> <code>^@</code> <code>text</code> <code>boolean</code></td>
<td>Returns true if the first string starts with the second string (equivalent to the <code>starts_with()</code> function).</td>
<td><code>'alphabet' ^@ 'alph'</code> <code>t</code></td>
</tr>
<tr>
<td><code>ascii</code> ( <code>text</code> ) <code>integer</code></td>
<td>Returns the numeric code of the first character of the argument. In UTF8 encoding, returns the Unicode code point of the character. In other multibyte encodings, the argument must be an ASCII character.</td>
<td><code>ascii('x')</code> <code>120</code></td>
</tr>
<tr>
<td><code>chr</code> ( <code>integer</code> ) <code>text</code></td>
<td>Returns the character with the given code. In UTF8 encoding the argument is treated as a Unicode code point. In other multibyte encodings the argument must designate an ASCII character. <code>chr(0)</code> is disallowed because text data types cannot store that character.</td>
<td><code>chr(65)</code> <code>A</code></td>
</tr>
<tr>
<td><code>concat</code> ( <code>val1</code> <code>"any"</code> [, <code>val2</code> <code>"any"</code> [, ...] ] ) <code>text</code></td>
<td>Concatenates the text representations of all the arguments. NULL arguments are ignored.</td>
<td><code>concat('abcde', 2, NULL, 22)</code> <code>abcde222</code></td>
</tr>
<tr>
<td><code>concat_ws</code> ( <code>sep</code> <code>text</code>, <code>val1</code> <code>"any"</code> [, <code>val2</code> <code>"any"</code> [, ...] ] ) <code>text</code></td>
<td>Concatenates all but the first argument, with separators. The first argument is used as the separator string, and should not be NULL. Other NULL arguments are ignored.</td>
<td><code>concat_ws(',', 'abcde', 2, NULL, 22)</code> <code>abcde,2,22</code></td>
</tr>
<tr>
<td><code>format</code> ( <code>formatstr</code> <code>text</code> [, <code>formatarg</code> <code>"any"</code> [, ...] ] ) <code>text</code></td>
<td>Formats arguments according to a format string; see <a href="#functions-string-format"><code>format</code></a>. This function is similar to the C function <code>sprintf</code>.</td>
<td><code>format('Hello %s, %1$s', 'World')</code> <code>Hello World, World</code></td>
</tr>
<tr>
<td><code>initcap</code> ( <code>text</code> ) <code>text</code></td>
<td>Converts the first letter of each word to upper case (or title case if the letter is a digraph and locale is <code>ICU</code> or <code>builtin</code> <code>PG_UNICODE_FAST</code>) and the rest to lower case. When using the <code>libc</code> or <code>builtin</code> locale provider, words are sequences of alphanumeric characters separated by non-alphanumeric characters; when using the ICU locale provider, words are separated according to <a href="https://unicode-org.github.io/icu-docs/apidoc/dev/icu4c/ustring_8h.html#a47602e2c2012d77ee91908b9bbfdc063">u_strToTitle ICU function</a>.</td>
<td>This function is primarily used for convenient display, and the specific result should not be relied upon because of the differences between locale providers and between different ICU versions. If specific word boundary rules are desired, it is recommended to write a custom function.<br><code>initcap('hi THOMAS')</code> <code>Hi Thomas</code></td>
</tr>
<tr>
<td><code>casefold</code> ( <code>text</code> ) <code>text</code></td>
<td>Performs case folding of the input string according to the collation. Case folding is similar to case conversion, but the purpose of case folding is to facilitate case-insensitive matching of strings, whereas the purpose of case conversion is to convert to a particular cased form. This function can only be used when the server encoding is <code>UTF8</code>.</td>
<td>Ordinarily, case folding simply converts to lowercase, but there may be exceptions depending on the collation. For instance, some characters have more than two lowercase variants, or fold to uppercase.<br>Case folding may change the length of the string. For instance, in the <code>PG_UNICODE_FAST</code> collation, <code>ß</code> (U+00DF) folds to <code>ss</code>.<br><code>casefold</code> can be used for Unicode Default Caseless Matching. It does not always preserve the normalized form of the input string (see <a href="#function-normalize">function-normalize</a>).<br>The <code>libc</code> provider doesn't support case folding, so <code>casefold</code> is identical to <a href="#function-lower">function-lower</a>.</td>
</tr>
<tr>
<td><code>left</code> ( <code>string</code> <code>text</code>, <code>n</code> <code>integer</code> ) <code>text</code></td>
<td>Returns first <code>n</code> characters in the string, or when <code>n</code> is negative, returns all but last |<code>n</code>| characters.</td>
<td><code>left('abcde', 2)</code> <code>ab</code></td>
</tr>
<tr>
<td><code>length</code> ( <code>text</code> ) <code>integer</code></td>
<td>Returns the number of characters in the string.</td>
<td><code>length('jose')</code> <code>4</code></td>
</tr>
<tr>
<td><code>md5</code> ( <code>text</code> ) <code>text</code></td>
<td>Computes the MD5 <a href="binary-string-functions-and-operators.md#functions-hash-note">hash</a> of the argument, with the result written in hexadecimal.</td>
<td><code>md5('abc')</code> <code>900150983cd24fb0​d6963f7d28e17f72</code></td>
</tr>
<tr>
<td><code>parse_ident</code> ( <code>qualified_identifier</code> <code>text</code> [, <code>strict_mode</code> <code>boolean</code> <code>DEFAULT</code> <code>true</code> ] ) <code>text[]</code></td>
<td>Splits <code>qualified_identifier</code> into an array of identifiers, removing any quoting of individual identifiers. By default, extra characters after the last identifier are considered an error; but if the second parameter is <code>false</code>, then such extra characters are ignored. (This behavior is useful for parsing names for objects like functions.) Note that this function does not truncate over-length identifiers. If you want truncation you can cast the result to <code>name[]</code>.</td>
<td><code>parse_ident('"SomeSchema".someTable')</code> <code>{SomeSchema,sometable}</code></td>
</tr>
<tr>
<td><code>pg_client_encoding</code> ( ) <code>name</code></td>
<td>Returns current client encoding name.</td>
<td><code>pg_client_encoding()</code> <code>UTF8</code></td>
</tr>
<tr>
<td><code>quote_ident</code> ( <code>text</code> ) <code>text</code></td>
<td>Returns the given string suitably quoted to be used as an identifier in an SQL statement string. Quotes are added only if necessary (i.e., if the string contains non-identifier characters or would be case-folded). Embedded quotes are properly doubled. See also <a href="../../server-programming/pl-pgsql-sql-procedural-language/basic-statements.md#plpgsql-quote-literal-example">Quoting Values in Dynamic Queries</a>.</td>
<td><code>quote_ident('Foo bar')</code> <code>"Foo bar"</code></td>
</tr>
<tr>
<td><code>quote_literal</code> ( <code>text</code> ) <code>text</code></td>
<td>Returns the given string suitably quoted to be used as a string literal in an SQL statement string. Embedded single-quotes and backslashes are properly doubled. Note that <code>quote_literal</code> returns null on null input; if the argument might be null, <code>quote_nullable</code> is often more suitable. See also <a href="../../server-programming/pl-pgsql-sql-procedural-language/basic-statements.md#plpgsql-quote-literal-example">Quoting Values in Dynamic Queries</a>.</td>
<td><code>quote_literal(E'O\'Reilly')</code> <code>'O''Reilly'</code></td>
</tr>
<tr>
<td><code>quote_literal</code> ( <code>anyelement</code> ) <code>text</code></td>
<td>Converts the given value to text and then quotes it as a literal. Embedded single-quotes and backslashes are properly doubled.</td>
<td><code>quote_literal(42.5)</code> <code>'42.5'</code></td>
</tr>
<tr>
<td><code>quote_nullable</code> ( <code>text</code> ) <code>text</code></td>
<td>Returns the given string suitably quoted to be used as a string literal in an SQL statement string; or, if the argument is null, returns <code>NULL</code>. Embedded single-quotes and backslashes are properly doubled. See also <a href="../../server-programming/pl-pgsql-sql-procedural-language/basic-statements.md#plpgsql-quote-literal-example">Quoting Values in Dynamic Queries</a>.</td>
<td><code>quote_nullable(NULL)</code> <code>NULL</code></td>
</tr>
<tr>
<td><code>quote_nullable</code> ( <code>anyelement</code> ) <code>text</code></td>
<td>Converts the given value to text and then quotes it as a literal; or, if the argument is null, returns <code>NULL</code>. Embedded single-quotes and backslashes are properly doubled.</td>
<td><code>quote_nullable(42.5)</code> <code>'42.5'</code></td>
</tr>
<tr>
<td><code>regexp_count</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>start</code> <code>integer</code> [, <code>flags</code> <code>text</code> ] ] ) <code>integer</code></td>
<td>Returns the number of times the POSIX regular expression <code>pattern</code> matches in the <code>string</code>; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_count('123456789012', '\d\d\d', 2)</code> <code>3</code></td>
</tr>
<tr>
<td><code>regexp_instr</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>start</code> <code>integer</code> [, <code>N</code> <code>integer</code> [, <code>endoption</code> <code>integer</code> [, <code>flags</code> <code>text</code> [, <code>subexpr</code> <code>integer</code> ] ] ] ] ] ) <code>integer</code></td>
<td>Returns the position within <code>string</code> where the <code>N</code>'th match of the POSIX regular expression <code>pattern</code> occurs, or zero if there is no such match; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_instr('ABCDEF', 'c(.)(..)', 1, 1, 0, 'i')</code> <code>3</code><br><code>regexp_instr('ABCDEF', 'c(.)(..)', 1, 1, 0, 'i', 2)</code> <code>5</code></td>
</tr>
<tr>
<td><code>regexp_like</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>flags</code> <code>text</code> ] ) <code>boolean</code></td>
<td>Checks whether a match of the POSIX regular expression <code>pattern</code> occurs within <code>string</code>; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_like('Hello World', 'world$', 'i')</code> <code>t</code></td>
</tr>
<tr>
<td><code>regexp_match</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>flags</code> <code>text</code> ] ) <code>text[]</code></td>
<td>Returns substrings within the first match of the POSIX regular expression <code>pattern</code> to the <code>string</code>; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_match('foobarbequebaz', '(bar)(beque)')</code> <code>{bar,beque}</code></td>
</tr>
<tr>
<td><code>regexp_matches</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>flags</code> <code>text</code> ] ) <code>setof text[]</code></td>
<td>Returns substrings within the first match of the POSIX regular expression <code>pattern</code> to the <code>string</code>, or substrings within all such matches if the <code>g</code> flag is used; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><p><code>regexp_matches('foobarbequebaz', 'ba.', 'g')</code></p>
<pre><code>
 {bar}
 {baz}</code></pre></td>
</tr>
<tr>
<td><code>regexp_replace</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code>, <code>replacement</code> <code>text</code> [, <code>flags</code> <code>text</code> ] ) <code>text</code></td>
<td>Replaces the substring that is the first match to the POSIX regular expression <code>pattern</code>, or all such matches if the <code>g</code> flag is used; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_replace('Thomas', '.[mN]a.', 'M')</code> <code>ThM</code></td>
</tr>
<tr>
<td><code>regexp_replace</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code>, <code>replacement</code> <code>text</code>, <code>start</code> <code>integer</code> [, <code>N</code> <code>integer</code> [, <code>flags</code> <code>text</code> ] ] ) <code>text</code></td>
<td>Replaces the substring that is the <code>N</code>'th match to the POSIX regular expression <code>pattern</code>, or all such matches if <code>N</code> is zero, with the search beginning at the <code>start</code>'th character of <code>string</code>. If <code>N</code> is omitted, it defaults to 1. See <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_replace('Thomas', '.', 'X', 3, 2)</code> <code>ThoXas</code><br><code>regexp_replace(string=&gt;'hello world', pattern=&gt;'l', replacement=&gt;'XX', start=&gt;1, "N"=&gt;2)</code> <code>helXXo world</code></td>
</tr>
<tr>
<td><code>regexp_split_to_array</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>flags</code> <code>text</code> ] ) <code>text[]</code></td>
<td>Splits <code>string</code> using a POSIX regular expression as the delimiter, producing an array of results; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_split_to_array('hello world', '\s+')</code> <code>{hello,world}</code></td>
</tr>
<tr>
<td><code>regexp_split_to_table</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>flags</code> <code>text</code> ] ) <code>setof text</code></td>
<td>Splits <code>string</code> using a POSIX regular expression as the delimiter, producing a set of results; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><p><code>regexp_split_to_table('hello world', '\s+')</code></p>
<pre><code>
 hello
 world</code></pre></td>
</tr>
<tr>
<td><code>regexp_substr</code> ( <code>string</code> <code>text</code>, <code>pattern</code> <code>text</code> [, <code>start</code> <code>integer</code> [, <code>N</code> <code>integer</code> [, <code>flags</code> <code>text</code> [, <code>subexpr</code> <code>integer</code> ] ] ] ] ) <code>text</code></td>
<td>Returns the substring within <code>string</code> that matches the <code>N</code>'th occurrence of the POSIX regular expression <code>pattern</code>, or <code>NULL</code> if there is no such match; see <a href="pattern-matching.md#functions-posix-regexp">POSIX Regular Expressions</a>.</td>
<td><code>regexp_substr('ABCDEF', 'c(.)(..)', 1, 1, 'i')</code> <code>CDEF</code><br><code>regexp_substr('ABCDEF', 'c(.)(..)', 1, 1, 'i', 2)</code> <code>EF</code></td>
</tr>
<tr>
<td><code>repeat</code> ( <code>string</code> <code>text</code>, <code>number</code> <code>integer</code> ) <code>text</code></td>
<td>Repeats <code>string</code> the specified <code>number</code> of times.</td>
<td><code>repeat('Pg', 4)</code> <code>PgPgPgPg</code></td>
</tr>
<tr>
<td><code>replace</code> ( <code>string</code> <code>text</code>, <code>from</code> <code>text</code>, <code>to</code> <code>text</code> ) <code>text</code></td>
<td>Replaces all occurrences in <code>string</code> of substring <code>from</code> with substring <code>to</code>.</td>
<td><code>replace('abcdefabcdef', 'cd', 'XX')</code> <code>abXXefabXXef</code></td>
</tr>
<tr>
<td><code>reverse</code> ( <code>text</code> ) <code>text</code></td>
<td>Reverses the order of the characters in the string.</td>
<td><code>reverse('abcde')</code> <code>edcba</code></td>
</tr>
<tr>
<td><code>right</code> ( <code>string</code> <code>text</code>, <code>n</code> <code>integer</code> ) <code>text</code></td>
<td>Returns last <code>n</code> characters in the string, or when <code>n</code> is negative, returns all but first |<code>n</code>| characters.</td>
<td><code>right('abcde', 2)</code> <code>de</code></td>
</tr>
<tr>
<td><code>split_part</code> ( <code>string</code> <code>text</code>, <code>delimiter</code> <code>text</code>, <code>n</code> <code>integer</code> ) <code>text</code></td>
<td>Splits <code>string</code> at occurrences of <code>delimiter</code> and returns the <code>n</code>'th field (counting from one), or when <code>n</code> is negative, returns the |<code>n</code>|'th-from-last field.</td>
<td><code>split_part('abc~@~def~@~ghi', '~@~', 2)</code> <code>def</code><br><code>split_part('abc,def,ghi,jkl', ',', -2)</code> <code>ghi</code></td>
</tr>
<tr>
<td><code>starts_with</code> ( <code>string</code> <code>text</code>, <code>prefix</code> <code>text</code> ) <code>boolean</code></td>
<td>Returns true if <code>string</code> starts with <code>prefix</code>.</td>
<td><code>starts_with('alphabet', 'alph')</code> <code>t</code></td>
</tr>
<tr>
<td><code>string_to_array</code> ( <code>string</code> <code>text</code>, <code>delimiter</code> <code>text</code> [, <code>null_string</code> <code>text</code> ] ) <code>text[]</code></td>
<td>Splits the <code>string</code> at occurrences of <code>delimiter</code> and forms the resulting fields into a <code>text</code> array. If <code>delimiter</code> is <code>NULL</code>, each character in the <code>string</code> will become a separate element in the array. If <code>delimiter</code> is an empty string, then the <code>string</code> is treated as a single field. If <code>null_string</code> is supplied and is not <code>NULL</code>, fields matching that string are replaced by <code>NULL</code>. See also <a href="array-functions-and-operators.md#function-array-to-string"><code>array_to_string</code></a>.</td>
<td><code>string_to_array('xx~~yy~~zz', '~~', 'yy')</code> <code>{xx,NULL,zz}</code></td>
</tr>
<tr>
<td><code>string_to_table</code> ( <code>string</code> <code>text</code>, <code>delimiter</code> <code>text</code> [, <code>null_string</code> <code>text</code> ] ) <code>setof text</code></td>
<td>Splits the <code>string</code> at occurrences of <code>delimiter</code> and returns the resulting fields as a set of <code>text</code> rows. If <code>delimiter</code> is <code>NULL</code>, each character in the <code>string</code> will become a separate row of the result. If <code>delimiter</code> is an empty string, then the <code>string</code> is treated as a single field. If <code>null_string</code> is supplied and is not <code>NULL</code>, fields matching that string are replaced by <code>NULL</code>.</td>
<td><p><code>string_to_table('xx~^~yy~^~zz', '~^~', 'yy')</code></p>
<pre><code>
 xx
 NULL
 zz</code></pre></td>
</tr>
<tr>
<td><code>strpos</code> ( <code>string</code> <code>text</code>, <code>substring</code> <code>text</code> ) <code>integer</code></td>
<td>Returns first starting index of the specified <code>substring</code> within <code>string</code>, or zero if it's not present. (Same as <code>position(</code>substring<code> in </code>string<code>)</code>, but note the reversed argument order.)</td>
<td><code>strpos('high', 'ig')</code> <code>2</code></td>
</tr>
<tr>
<td><code>substr</code> ( <code>string</code> <code>text</code>, <code>start</code> <code>integer</code> [, <code>count</code> <code>integer</code> ] ) <code>text</code></td>
<td>Extracts the substring of <code>string</code> starting at the <code>start</code>'th character, and extending for <code>count</code> characters if that is specified. (Same as <code>substring(</code>string<code> from </code>start<code> for </code>count<code>)</code>.)</td>
<td><code>substr('alphabet', 3)</code> <code>phabet</code><br><code>substr('alphabet', 3, 2)</code> <code>ph</code></td>
</tr>
<tr>
<td><code>to_ascii</code> ( <code>string</code> <code>text</code> ) <code>text</code></td>
<td><code>to_ascii</code> ( <code>string</code> <code>text</code>, <code>encoding</code> <code>name</code> ) <code>text</code></td>
<td><code>to_ascii</code> ( <code>string</code> <code>text</code>, <code>encoding</code> <code>integer</code> ) <code>text</code><br>Converts <code>string</code> to ASCII from another encoding, which may be identified by name or number. If <code>encoding</code> is omitted the database encoding is assumed (which in practice is the only useful case). The conversion consists primarily of dropping accents. Conversion is only supported from <code>LATIN1</code>, <code>LATIN2</code>, <code>LATIN9</code>, and <code>WIN1250</code> encodings. (See the <a href="../../appendixes/additional-supplied-modules-and-extensions/unaccent-a-text-search-dictionary-which-removes-diacritics.md#unaccent">unaccent</a> module for another, more flexible solution.)<br><code>to_ascii('Kar&amp;eacute;l')</code> <code>Karel</code></td>
</tr>
<tr>
<td><code>to_bin</code> ( <code>integer</code> ) <code>text</code></td>
<td><code>to_bin</code> ( <code>bigint</code> ) <code>text</code></td>
<td>Converts the number to its equivalent two's complement binary representation.<br><code>to_bin(2147483647)</code> <code>1111111111111111111111111111111</code><br><code>to_bin(-1234)</code> <code>11111111111111111111101100101110</code></td>
</tr>
<tr>
<td><code>to_hex</code> ( <code>integer</code> ) <code>text</code></td>
<td><code>to_hex</code> ( <code>bigint</code> ) <code>text</code></td>
<td>Converts the number to its equivalent two's complement hexadecimal representation.<br><code>to_hex(2147483647)</code> <code>7fffffff</code><br><code>to_hex(-1234)</code> <code>fffffb2e</code></td>
</tr>
<tr>
<td><code>to_oct</code> ( <code>integer</code> ) <code>text</code></td>
<td><code>to_oct</code> ( <code>bigint</code> ) <code>text</code></td>
<td>Converts the number to its equivalent two's complement octal representation.<br><code>to_oct(2147483647)</code> <code>17777777777</code><br><code>to_oct(-1234)</code> <code>37777775456</code></td>
</tr>
<tr>
<td><code>translate</code> ( <code>string</code> <code>text</code>, <code>from</code> <code>text</code>, <code>to</code> <code>text</code> ) <code>text</code></td>
<td>Replaces each character in <code>string</code> that matches a character in the <code>from</code> set with the corresponding character in the <code>to</code> set. If <code>from</code> is longer than <code>to</code>, occurrences of the extra characters in <code>from</code> are deleted.</td>
<td><code>translate('12345', '143', 'ax')</code> <code>a2x5</code></td>
</tr>
<tr>
<td><code>unistr</code> ( <code>text</code> ) <code>text</code></td>
<td>Evaluate escaped Unicode characters in the argument. Unicode characters can be specified as <code>\</code><em>XXXX</em> (4 hexadecimal digits), <code>\+</code><em>XXXXXX</em> (6 hexadecimal digits), <code>\u</code><em>XXXX</em> (4 hexadecimal digits), or <code>\U</code><em>XXXXXXXX</em> (8 hexadecimal digits). To specify a backslash, write two backslashes. All other characters are taken literally.</td>
<td>If the server encoding is not UTF-8, the Unicode code point identified by one of these escape sequences is converted to the actual server encoding; an error is reported if that's not possible.<br>This function provides a (non-standard) alternative to string constants with Unicode escapes (see <a href="../sql-syntax/lexical-structure.md#sql-syntax-strings-uescape">String Constants with Unicode Escapes</a>).<br><code>unistr('d\0061t\+000061')</code> <code>data</code><br><code>unistr('d\u0061t\U00000061')</code> <code>data</code></td>
</tr>
</tbody>
</table>


 The `concat`, `concat_ws` and `format` functions are variadic, so it is possible to pass the values to be concatenated or formatted as an array marked with the `VARIADIC` keyword (see [SQL Functions with Variable Numbers of Arguments](../../server-programming/extending-sql/query-language-sql-functions.md#xfunc-sql-variadic-functions)). The array's elements are treated as if they were separate ordinary arguments to the function. If the variadic array argument is NULL, `concat` and `concat_ws` return NULL, but `format` treats a NULL as a zero-element array.


 See also the aggregate function `string_agg` in [Aggregate Functions](aggregate-functions.md#functions-aggregate), and the functions for converting between strings and the `bytea` type in [Text/Binary String Conversion Functions](binary-string-functions-and-operators.md#functions-binarystring-conversions).


### `format` { #functions-string-format }


 The function `format` produces output formatted according to a format string, in a style similar to the C function `sprintf`.


```

format(formatstr text [, formatarg "any" [, ...] ])
```
 `formatstr` is a format string that specifies how the result should be formatted. Text in the format string is copied directly to the result, except where *format specifiers* are used. Format specifiers act as placeholders in the string, defining how subsequent function arguments should be formatted and inserted into the result. Each `formatarg` argument is converted to text according to the usual output rules for its data type, and then formatted and inserted into the result string according to the format specifier(s).


 Format specifiers are introduced by a `%` character and have the form

```

%[position][flags][width]type
```
 where the component fields are:

`position` (optional)
:   A string of the form ``n`$` where `n` is the index of the argument to print. Index 1 means the first argument after `formatstr`. If the `position` is omitted, the default is to use the next argument in sequence.

`flags` (optional)
:   Additional options controlling how the format specifier's output is formatted. Currently the only supported flag is a minus sign (`-`) which will cause the format specifier's output to be left-justified. This has no effect unless the `width` field is also specified.

`width` (optional)
:   Specifies the *minimum* number of characters to use to display the format specifier's output. The output is padded on the left or right (depending on the `-` flag) with spaces as needed to fill the width. A too-small width does not cause truncation of the output, but is simply ignored. The width may be specified using any of the following: a positive integer; an asterisk (`*`) to use the next function argument as the width; or a string of the form `*`n`$` to use the `n`th function argument as the width.


     If the width comes from a function argument, that argument is consumed before the argument that is used for the format specifier's value. If the width argument is negative, the result is left aligned (as if the `-` flag had been specified) within a field of length `abs`(`width`).

`type` (required)
:   The type of format conversion to use to produce the format specifier's output. The following types are supported:

    -  `s` formats the argument value as a simple string. A null value is treated as an empty string.
    -  `I` treats the argument value as an SQL identifier, double-quoting it if necessary. It is an error for the value to be null (equivalent to `quote_ident`).
    -  `L` quotes the argument value as an SQL literal. A null value is displayed as the string `NULL`, without quotes (equivalent to `quote_nullable`).


 In addition to the format specifiers described above, the special sequence `%%` may be used to output a literal `%` character.


 Here are some examples of the basic format conversions:

```

SELECT format('Hello %s', 'World');
Result: Hello World

SELECT format('Testing %s, %s, %s, %%', 'one', 'two', 'three');
Result: Testing one, two, three, %

SELECT format('INSERT INTO %I VALUES(%L)', 'Foo bar', E'O\'Reilly');
Result: INSERT INTO "Foo bar" VALUES('O''Reilly')

SELECT format('INSERT INTO %I VALUES(%L)', 'locations', 'C:\Program Files');
Result: INSERT INTO locations VALUES('C:\Program Files')
```


 Here are examples using `width` fields and the `-` flag:

```

SELECT format('|%10s|', 'foo');
Result: |       foo|

SELECT format('|%-10s|', 'foo');
Result: |foo       |

SELECT format('|%*s|', 10, 'foo');
Result: |       foo|

SELECT format('|%*s|', -10, 'foo');
Result: |foo       |

SELECT format('|%-*s|', 10, 'foo');
Result: |foo       |

SELECT format('|%-*s|', -10, 'foo');
Result: |foo       |
```


 These examples show use of `position` fields:

```

SELECT format('Testing %3$s, %2$s, %1$s', 'one', 'two', 'three');
Result: Testing three, two, one

SELECT format('|%*2$s|', 'foo', 10, 'bar');
Result: |       bar|

SELECT format('|%1$*2$s|', 'foo', 10, 'bar');
Result: |       foo|
```


 Unlike the standard C function `sprintf`, PostgreSQL's `format` function allows format specifiers with and without `position` fields to be mixed in the same format string. A format specifier without a `position` field always uses the next argument after the last argument consumed. In addition, the `format` function does not require all function arguments to be used in the format string. For example:

```

SELECT format('Testing %3$s, %2$s, %s', 'one', 'two', 'three');
Result: Testing three, two, three
```


 The `%I` and `%L` format specifiers are particularly useful for safely constructing dynamic SQL statements. See [Quoting Values in Dynamic Queries](../../server-programming/pl-pgsql-sql-procedural-language/basic-statements.md#plpgsql-quote-literal-example).
