<a id="hstore"></a>

## hstore — hstore key/value datatype


 This module implements the `hstore` data type for storing sets of key/value pairs within a single PostgreSQL value. This can be useful in various scenarios, such as rows with many attributes that are rarely examined, or semi-structured data. Keys and values are simply text strings.


 This module is considered “trusted”, that is, it can be installed by non-superusers who have `CREATE` privilege on the current database.
 <a id="hstore-external-rep"></a>

### `hstore` External Representation


 The text representation of an `hstore`, used for input and output, includes zero or more *key* `=>` *value* pairs separated by commas. Some examples:

```

k => v
foo => bar, baz => whatever
"1-a" => "anything at all"
```
 The order of the pairs is not significant (and may not be reproduced on output). Whitespace between pairs or around the `=>` sign is ignored. Double-quote keys and values that include whitespace, commas, `=`s or `>`s. To include a double quote or a backslash in a key or value, escape it with a backslash.


 Each key in an `hstore` is unique. If you declare an `hstore` with duplicate keys, only one will be stored in the `hstore` and there is no guarantee as to which will be kept:

```sql

SELECT 'a=>1,a=>2'::hstore;
  hstore
----------
 "a"=>"1"
```


 A value (but not a key) can be an SQL `NULL`. For example:

```

key => NULL
```
 The `NULL` keyword is case-insensitive. Double-quote the `NULL` to treat it as the ordinary string “NULL”.


!!! note

    Keep in mind that the `hstore` text format, when used for input, applies *before* any required quoting or escaping. If you are passing an `hstore` literal via a parameter, then no additional processing is needed. But if you're passing it as a quoted literal constant, then any single-quote characters need to be escaped correctly. See [String Constants](../../the-sql-language/sql-syntax/lexical-structure.md#sql-syntax-strings) for more on the handling of string constants.


 On output, double quotes always surround keys and values, even when it's not strictly necessary.
  <a id="hstore-ops-funcs"></a>

### `hstore` Operators and Functions


 The operators provided by the `hstore` module are shown in [`hstore` Operators](#hstore-op-table), the functions in [`hstore` Functions](#hstore-func-table).
 <a id="hstore-op-table"></a>

**Table: `hstore` Operators**

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
<td><code>hstore</code> <code>-&gt;</code> <code>text</code> <code>text</code></td>
<td>Returns value associated with given key, or <code>NULL</code> if not present.</td>
<td><code>'a=&gt;x, b=&gt;y'::hstore -&gt; 'a'</code> <code>x</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>-&gt;</code> <code>text[]</code> <code>text[]</code></td>
<td>Returns values associated with given keys, or <code>NULL</code> if not present.</td>
<td><code>'a=&gt;x, b=&gt;y, c=&gt;z'::hstore -&gt; ARRAY['c','a']</code> <code>{"z","x"}</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>||</code> <code>hstore</code> <code>hstore</code></td>
<td>Concatenates two <code>hstore</code>s.</td>
<td><code>'a=&gt;b, c=&gt;d'::hstore || 'c=&gt;x, d=&gt;q'::hstore</code> <code>"a"=&gt;"b", "c"=&gt;"x", "d"=&gt;"q"</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>?</code> <code>text</code> <code>boolean</code></td>
<td>Does <code>hstore</code> contain key?</td>
<td><code>'a=&gt;1'::hstore ? 'a'</code> <code>t</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>?&amp;</code> <code>text[]</code> <code>boolean</code></td>
<td>Does <code>hstore</code> contain all the specified keys?</td>
<td><code>'a=&gt;1,b=&gt;2'::hstore ?&amp; ARRAY['a','b']</code> <code>t</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>?|</code> <code>text[]</code> <code>boolean</code></td>
<td>Does <code>hstore</code> contain any of the specified keys?</td>
<td><code>'a=&gt;1,b=&gt;2'::hstore ?| ARRAY['b','c']</code> <code>t</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>@&gt;</code> <code>hstore</code> <code>boolean</code></td>
<td>Does left operand contain right?</td>
<td><code>'a=&gt;b, b=&gt;1, c=&gt;NULL'::hstore @&gt; 'b=&gt;1'</code> <code>t</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>&lt;@</code> <code>hstore</code> <code>boolean</code></td>
<td>Is left operand contained in right?</td>
<td><code>'a=&gt;c'::hstore &lt;@ 'a=&gt;b, b=&gt;1, c=&gt;NULL'</code> <code>f</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>-</code> <code>text</code> <code>hstore</code></td>
<td>Deletes key from left operand.</td>
<td><code>'a=&gt;1, b=&gt;2, c=&gt;3'::hstore - 'b'::text</code> <code>"a"=&gt;"1", "c"=&gt;"3"</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>-</code> <code>text[]</code> <code>hstore</code></td>
<td>Deletes keys from left operand.</td>
<td><code>'a=&gt;1, b=&gt;2, c=&gt;3'::hstore - ARRAY['a','b']</code> <code>"c"=&gt;"3"</code></td>
</tr>
<tr>
<td><code>hstore</code> <code>-</code> <code>hstore</code> <code>hstore</code></td>
<td>Deletes pairs from left operand that match pairs in the right operand.</td>
<td><code>'a=&gt;1, b=&gt;2, c=&gt;3'::hstore - 'a=&gt;4, b=&gt;2'::hstore</code> <code>"a"=&gt;"1", "c"=&gt;"3"</code></td>
</tr>
<tr>
<td><code>anyelement</code> <code>#=</code> <code>hstore</code> <code>anyelement</code></td>
<td>Replaces fields in the left operand (which must be a composite type) with matching values from <code>hstore</code>.</td>
<td><code>ROW(1,3) #= 'f1=&gt;11'::hstore</code> <code>(11,3)</code></td>
</tr>
<tr>
<td><code>%%</code> <code>hstore</code> <code>text[]</code></td>
<td>Converts <code>hstore</code> to an array of alternating keys and values.</td>
<td><code>%% 'a=&gt;foo, b=&gt;bar'::hstore</code> <code>{a,foo,b,bar}</code></td>
</tr>
<tr>
<td><code>%#</code> <code>hstore</code> <code>text[]</code></td>
<td>Converts <code>hstore</code> to a two-dimensional key/value array.</td>
<td><code>%# 'a=&gt;foo, b=&gt;bar'::hstore</code> <code>{{a,foo},{b,bar}}</code></td>
</tr>
</tbody>
</table>
 <a id="hstore-func-table"></a>

**Table: `hstore` Functions**

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
<td><code>hstore</code> ( <code>record</code> ) <code>hstore</code></td>
<td>Constructs an <code>hstore</code> from a record or row.</td>
<td><code>hstore(ROW(1,2))</code> <code>"f1"=&gt;"1", "f2"=&gt;"2"</code></td>
</tr>
<tr>
<td><code>hstore</code> ( <code>text[]</code> ) <code>hstore</code></td>
<td>Constructs an <code>hstore</code> from an array, which may be either a key/value array, or a two-dimensional array.</td>
<td><code>hstore(ARRAY['a','1','b','2'])</code> <code>"a"=&gt;"1", "b"=&gt;"2"</code><br><code>hstore(ARRAY[['c','3'],['d','4']])</code> <code>"c"=&gt;"3", "d"=&gt;"4"</code></td>
</tr>
<tr>
<td><code>hstore</code> ( <code>text[]</code>, <code>text[]</code> ) <code>hstore</code></td>
<td>Constructs an <code>hstore</code> from separate key and value arrays.</td>
<td><code>hstore(ARRAY['a','b'], ARRAY['1','2'])</code> <code>"a"=&gt;"1", "b"=&gt;"2"</code></td>
</tr>
<tr>
<td><code>hstore</code> ( <code>text</code>, <code>text</code> ) <code>hstore</code></td>
<td>Makes a single-item <code>hstore</code>.</td>
<td><code>hstore('a', 'b')</code> <code>"a"=&gt;"b"</code></td>
</tr>
<tr>
<td><code>akeys</code> ( <code>hstore</code> ) <code>text[]</code></td>
<td>Extracts an <code>hstore</code>'s keys as an array.</td>
<td><code>akeys('a=&gt;1,b=&gt;2')</code> <code>{a,b}</code></td>
</tr>
<tr>
<td><code>skeys</code> ( <code>hstore</code> ) <code>setof text</code></td>
<td>Extracts an <code>hstore</code>'s keys as a set.</td>
<td><p><code>skeys('a=&gt;1,b=&gt;2')</code></p>
<pre><code>
a
b</code></pre></td>
</tr>
<tr>
<td><code>avals</code> ( <code>hstore</code> ) <code>text[]</code></td>
<td>Extracts an <code>hstore</code>'s values as an array.</td>
<td><code>avals('a=&gt;1,b=&gt;2')</code> <code>{1,2}</code></td>
</tr>
<tr>
<td><code>svals</code> ( <code>hstore</code> ) <code>setof text</code></td>
<td>Extracts an <code>hstore</code>'s values as a set.</td>
<td><p><code>svals('a=&gt;1,b=&gt;2')</code></p>
<pre><code>
1
2</code></pre></td>
</tr>
<tr>
<td><code>hstore_to_array</code> ( <code>hstore</code> ) <code>text[]</code></td>
<td>Extracts an <code>hstore</code>'s keys and values as an array of alternating keys and values.</td>
<td><code>hstore_to_array('a=&gt;1,b=&gt;2')</code> <code>{a,1,b,2}</code></td>
</tr>
<tr>
<td><code>hstore_to_matrix</code> ( <code>hstore</code> ) <code>text[]</code></td>
<td>Extracts an <code>hstore</code>'s keys and values as a two-dimensional array.</td>
<td><code>hstore_to_matrix('a=&gt;1,b=&gt;2')</code> <code>{{a,1},{b,2}}</code></td>
</tr>
<tr>
<td><code>hstore_to_json</code> ( <code>hstore</code> ) <code>json</code></td>
<td>Converts an <code>hstore</code> to a <code>json</code> value, converting all non-null values to JSON strings.</td>
<td>This function is used implicitly when an <code>hstore</code> value is cast to <code>json</code>.<br><code>hstore_to_json('"a key"=&gt;1, b=&gt;t, c=&gt;null, d=&gt;12345, e=&gt;012345, f=&gt;1.234, g=&gt;2.345e+4')</code> <code>{"a key": "1", "b": "t", "c": null, "d": "12345", "e": "012345", "f": "1.234", "g": "2.345e+4"}</code></td>
</tr>
<tr>
<td><code>hstore_to_jsonb</code> ( <code>hstore</code> ) <code>jsonb</code></td>
<td>Converts an <code>hstore</code> to a <code>jsonb</code> value, converting all non-null values to JSON strings.</td>
<td>This function is used implicitly when an <code>hstore</code> value is cast to <code>jsonb</code>.<br><code>hstore_to_jsonb('"a key"=&gt;1, b=&gt;t, c=&gt;null, d=&gt;12345, e=&gt;012345, f=&gt;1.234, g=&gt;2.345e+4')</code> <code>{"a key": "1", "b": "t", "c": null, "d": "12345", "e": "012345", "f": "1.234", "g": "2.345e+4"}</code></td>
</tr>
<tr>
<td><code>hstore_to_json_loose</code> ( <code>hstore</code> ) <code>json</code></td>
<td>Converts an <code>hstore</code> to a <code>json</code> value, but attempts to distinguish numerical and Boolean values so they are unquoted in the JSON.</td>
<td><code>hstore_to_json_loose('"a key"=&gt;1, b=&gt;t, c=&gt;null, d=&gt;12345, e=&gt;012345, f=&gt;1.234, g=&gt;2.345e+4')</code> <code>{"a key": 1, "b": true, "c": null, "d": 12345, "e": "012345", "f": 1.234, "g": 2.345e+4}</code></td>
</tr>
<tr>
<td><code>hstore_to_jsonb_loose</code> ( <code>hstore</code> ) <code>jsonb</code></td>
<td>Converts an <code>hstore</code> to a <code>jsonb</code> value, but attempts to distinguish numerical and Boolean values so they are unquoted in the JSON.</td>
<td><code>hstore_to_jsonb_loose('"a key"=&gt;1, b=&gt;t, c=&gt;null, d=&gt;12345, e=&gt;012345, f=&gt;1.234, g=&gt;2.345e+4')</code> <code>{"a key": 1, "b": true, "c": null, "d": 12345, "e": "012345", "f": 1.234, "g": 2.345e+4}</code></td>
</tr>
<tr>
<td><code>slice</code> ( <code>hstore</code>, <code>text[]</code> ) <code>hstore</code></td>
<td>Extracts a subset of an <code>hstore</code> containing only the specified keys.</td>
<td><code>slice('a=&gt;1,b=&gt;2,c=&gt;3'::hstore, ARRAY['b','c','x'])</code> <code>"b"=&gt;"2", "c"=&gt;"3"</code></td>
</tr>
<tr>
<td><code>each</code> ( <code>hstore</code> ) <code>setof record</code> ( <code>key</code> <code>text</code>, <code>value</code> <code>text</code> )</td>
<td>Extracts an <code>hstore</code>'s keys and values as a set of records.</td>
<td><p><code>SELECT * FROM each('a=&gt;1,b=&gt;2')</code></p>
<pre><code>
 key | value
-----+-------
 a   | 1
 b   | 2</code></pre></td>
</tr>
<tr>
<td><code>exist</code> ( <code>hstore</code>, <code>text</code> ) <code>boolean</code></td>
<td>Does <code>hstore</code> contain key?</td>
<td><code>exist('a=&gt;1', 'a')</code> <code>t</code></td>
</tr>
<tr>
<td><code>defined</code> ( <code>hstore</code>, <code>text</code> ) <code>boolean</code></td>
<td>Does <code>hstore</code> contain a non-<code>NULL</code> value for key?</td>
<td><code>defined('a=&gt;NULL', 'a')</code> <code>f</code></td>
</tr>
<tr>
<td><code>delete</code> ( <code>hstore</code>, <code>text</code> ) <code>hstore</code></td>
<td>Deletes pair with matching key.</td>
<td><code>delete('a=&gt;1,b=&gt;2', 'b')</code> <code>"a"=&gt;"1"</code></td>
</tr>
<tr>
<td><code>delete</code> ( <code>hstore</code>, <code>text[]</code> ) <code>hstore</code></td>
<td>Deletes pairs with matching keys.</td>
<td><code>delete('a=&gt;1,b=&gt;2,c=&gt;3', ARRAY['a','b'])</code> <code>"c"=&gt;"3"</code></td>
</tr>
<tr>
<td><code>delete</code> ( <code>hstore</code>, <code>hstore</code> ) <code>hstore</code></td>
<td>Deletes pairs matching those in the second argument.</td>
<td><code>delete('a=&gt;1,b=&gt;2', 'a=&gt;4,b=&gt;2'::hstore)</code> <code>"a"=&gt;"1"</code></td>
</tr>
<tr>
<td><code>populate_record</code> ( <code>anyelement</code>, <code>hstore</code> ) <code>anyelement</code></td>
<td>Replaces fields in the left operand (which must be a composite type) with matching values from <code>hstore</code>.</td>
<td><code>populate_record(ROW(1,2), 'f1=&gt;42'::hstore)</code> <code>(42,2)</code></td>
</tr>
</tbody>
</table>


 In addition to these operators and functions, values of the `hstore` type can be subscripted, allowing them to act like associative arrays. Only a single subscript of type `text` can be specified; it is interpreted as a key and the corresponding value is fetched or stored. For example,

```sql

CREATE TABLE mytable (h hstore);
INSERT INTO mytable VALUES ('a=>b, c=>d');
SELECT h['a'] FROM mytable;
 h
---
 b
(1 row)

UPDATE mytable SET h['c'] = 'new';
SELECT h FROM mytable;
          h
----------------------
 "a"=>"b", "c"=>"new"
(1 row)
```
 A subscripted fetch returns `NULL` if the subscript is `NULL` or that key does not exist in the `hstore`. (Thus, a subscripted fetch is not greatly different from the `->` operator.) A subscripted update fails if the subscript is `NULL`; otherwise, it replaces the value for that key, adding an entry to the `hstore` if the key does not already exist.
  <a id="hstore-indexes"></a>

### Indexes


 `hstore` has GiST and GIN index support for the `@>`, `?`, `?&` and `?|` operators. For example:


```sql

CREATE INDEX hidx ON testhstore USING GIST (h);

CREATE INDEX hidx ON testhstore USING GIN (h);
```


 `gist_hstore_ops` GiST opclass approximates a set of key/value pairs as a bitmap signature. Its optional integer parameter `siglen` determines the signature length in bytes. The default length is 16 bytes. Valid values of signature length are between 1 and 2024 bytes. Longer signatures lead to a more precise search (scanning a smaller fraction of the index and fewer heap pages), at the cost of a larger index.


 Example of creating such an index with a signature length of 32 bytes:

```sql

CREATE INDEX hidx ON testhstore USING GIST (h gist_hstore_ops(siglen=32));
```


 `hstore` also supports `btree` or `hash` indexes for the `=` operator. This allows `hstore` columns to be declared `UNIQUE`, or to be used in `GROUP BY`, `ORDER BY` or `DISTINCT` expressions. The sort ordering for `hstore` values is not particularly useful, but these indexes may be useful for equivalence lookups. Create indexes for `=` comparisons as follows:


```sql

CREATE INDEX hidx ON testhstore USING BTREE (h);

CREATE INDEX hidx ON testhstore USING HASH (h);
```
  <a id="hstore-examples"></a>

### Examples


 Add a key, or update an existing key with a new value:

```sql

UPDATE tab SET h['c'] = '3';
```
 Another way to do the same thing is:

```sql

UPDATE tab SET h = h || hstore('c', '3');
```
 If multiple keys are to be added or changed in one operation, the concatenation approach is more efficient than subscripting:

```sql

UPDATE tab SET h = h || hstore(ARRAY['q', 'w'], ARRAY['11', '12']);
```


 Delete a key:

```sql

UPDATE tab SET h = delete(h, 'k1');
```


 Convert a `record` to an `hstore`:

```sql

CREATE TABLE test (col1 integer, col2 text, col3 text);
INSERT INTO test VALUES (123, 'foo', 'bar');

SELECT hstore(t) FROM test AS t;
                   hstore
---------------------------------------------
 "col1"=>"123", "col2"=>"foo", "col3"=>"bar"
(1 row)
```


 Convert an `hstore` to a predefined `record` type:

```sql

CREATE TABLE test (col1 integer, col2 text, col3 text);

SELECT * FROM populate_record(null::test,
                              '"col1"=>"456", "col2"=>"zzz"');
 col1 | col2 | col3
------+------+------
  456 | zzz  |
(1 row)
```


 Modify an existing record using the values from an `hstore`:

```sql

CREATE TABLE test (col1 integer, col2 text, col3 text);
INSERT INTO test VALUES (123, 'foo', 'bar');

SELECT (r).* FROM (SELECT t #= '"col3"=>"baz"' AS r FROM test t) s;
 col1 | col2 | col3
------+------+------
  123 | foo  | baz
(1 row)
```

  <a id="hstore-statistics"></a>

### Statistics


 The `hstore` type, because of its intrinsic liberality, could contain a lot of different keys. Checking for valid keys is the task of the application. The following examples demonstrate several techniques for checking keys and obtaining statistics.


 Simple example:

```sql

SELECT * FROM each('aaa=>bq, b=>NULL, ""=>1');
```


 Using a table:

```sql

CREATE TABLE stat AS SELECT (each(h)).key, (each(h)).value FROM testhstore;
```


 Online statistics:

```sql

SELECT key, count(*) FROM
  (SELECT (each(h)).key FROM testhstore) AS stat
  GROUP BY key
  ORDER BY count DESC, key;
    key    | count
-----------+-------
 line      |   883
 query     |   207
 pos       |   203
 node      |   202
 space     |   197
 status    |   195
 public    |   194
 title     |   190
 org       |   189
...................
```

  <a id="hstore-compatibility"></a>

### Compatibility


 As of PostgreSQL 9.0, `hstore` uses a different internal representation than previous versions. This presents no obstacle for dump/restore upgrades since the text representation (used in the dump) is unchanged.


 In the event of a binary upgrade, upward compatibility is maintained by having the new code recognize old-format data. This will entail a slight performance penalty when processing data that has not yet been modified by the new code. It is possible to force an upgrade of all values in a table column by doing an `UPDATE` statement as follows:

```sql

UPDATE tablename SET hstorecol = hstorecol || '';
```


 Another way to do it is:

```sql

ALTER TABLE tablename ALTER hstorecol TYPE hstore USING hstorecol || '';
```
 The `ALTER TABLE` method requires an `ACCESS EXCLUSIVE` lock on the table, but does not result in bloating the table with old row versions.
  <a id="hstore-transforms"></a>

### Transforms


 Additional extensions are available that implement transforms for the `hstore` type for the languages PL/Perl and PL/Python. The extensions for PL/Perl are called `hstore_plperl` and `hstore_plperlu`, for trusted and untrusted PL/Perl. If you install these transforms and specify them when creating a function, `hstore` values are mapped to Perl hashes. The extension for PL/Python is called `hstore_plpython3u`. If you use it, `hstore` values are mapped to Python dictionaries.
  <a id="hstore-authors"></a>

### Authors


 Oleg Bartunov [oleg@sai.msu.su](mailto:oleg@sai.msu.su), Moscow, Moscow University, Russia


 Teodor Sigaev [teodor@sigaev.ru](mailto:teodor@sigaev.ru), Moscow, Delta-Soft Ltd., Russia


 Additional enhancements by Andrew Gierth [andrew@tao11.riddles.org.uk](mailto:andrew@tao11.riddles.org.uk), United Kingdom
