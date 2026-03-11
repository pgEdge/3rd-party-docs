## JSON Functions and Operators { #functions-json }


 This section describes:

-  functions and operators for processing and creating JSON data
-  the SQL/JSON path language
-  the SQL/JSON query functions


 To provide native support for JSON data types within the SQL environment, PostgreSQL implements the *SQL/JSON data model*. This model comprises sequences of items. Each item can hold SQL scalar values, with an additional SQL/JSON null value, and composite data structures that use JSON arrays and objects. The model is a formalization of the implied data model in the JSON specification [RFC 7159](https://datatracker.ietf.org/doc/html/rfc7159).


 SQL/JSON allows you to handle JSON data alongside regular SQL data, with transaction support, including:

-  Uploading JSON data into the database and storing it in regular SQL columns as character or binary strings.
-  Generating JSON objects and arrays from relational data.
-  Querying JSON data using SQL/JSON query functions and SQL/JSON path language expressions.


 To learn more about the SQL/JSON standard, see [SQL Technical Report](../../bibliography.md#sqltr-19075-6). For details on JSON types supported in PostgreSQL, see [JSON Types](../data-types/json-types.md#datatype-json).


### Processing and Creating JSON Data { #functions-json-processing }


 [`json` and `jsonb` Operators](#functions-json-op-table) shows the operators that are available for use with JSON data types (see [JSON Types](../data-types/json-types.md#datatype-json)). In addition, the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) are available for `jsonb`, though not for `json`. The comparison operators follow the ordering rules for B-tree operations outlined in [`jsonb` Indexing](../data-types/json-types.md#json-indexing). See also [Aggregate Functions](aggregate-functions.md#functions-aggregate) for the aggregate function `json_agg` which aggregates record values as JSON, the aggregate function `json_object_agg` which aggregates pairs of values into a JSON object, and their `jsonb` equivalents, `jsonb_agg` and `jsonb_object_agg`.


<a id="functions-json-op-table"></a>
**Table: `json` and `jsonb` Operators**

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
<td><code>json</code> <code>-&gt;</code> <code>integer</code> <code>json</code></td>
<td><code>jsonb</code> <code>-&gt;</code> <code>integer</code> <code>jsonb</code></td>
<td>Extracts <code>n</code>'th element of JSON array (array elements are indexed from zero, but negative integers count from the end).<br><code>'[{"a":"foo"},{"b":"bar"},{"c":"baz"}]'::json -&gt; 2</code> <code>{"c":"baz"}</code><br><code>'[{"a":"foo"},{"b":"bar"},{"c":"baz"}]'::json -&gt; -3</code> <code>{"a":"foo"}</code></td>
</tr>
<tr>
<td><code>json</code> <code>-&gt;</code> <code>text</code> <code>json</code></td>
<td><code>jsonb</code> <code>-&gt;</code> <code>text</code> <code>jsonb</code></td>
<td>Extracts JSON object field with the given key.<br><code>'{"a": {"b":"foo"}}'::json -&gt; 'a'</code> <code>{"b":"foo"}</code></td>
</tr>
<tr>
<td><code>json</code> <code>-&gt;&gt;</code> <code>integer</code> <code>text</code></td>
<td><code>jsonb</code> <code>-&gt;&gt;</code> <code>integer</code> <code>text</code></td>
<td>Extracts <code>n</code>'th element of JSON array, as <code>text</code>.<br><code>'[1,2,3]'::json -&gt;&gt; 2</code> <code>3</code></td>
</tr>
<tr>
<td><code>json</code> <code>-&gt;&gt;</code> <code>text</code> <code>text</code></td>
<td><code>jsonb</code> <code>-&gt;&gt;</code> <code>text</code> <code>text</code></td>
<td>Extracts JSON object field with the given key, as <code>text</code>.<br><code>'{"a":1,"b":2}'::json -&gt;&gt; 'b'</code> <code>2</code></td>
</tr>
<tr>
<td><code>json</code> <code>#&gt;</code> <code>text[]</code> <code>json</code></td>
<td><code>jsonb</code> <code>#&gt;</code> <code>text[]</code> <code>jsonb</code></td>
<td>Extracts JSON sub-object at the specified path, where path elements can be either field keys or array indexes.<br><code>'{"a": {"b": ["foo","bar"]}}'::json #&gt; '{a,b,1}'</code> <code>"bar"</code></td>
</tr>
<tr>
<td><code>json</code> <code>#&gt;&gt;</code> <code>text[]</code> <code>text</code></td>
<td><code>jsonb</code> <code>#&gt;&gt;</code> <code>text[]</code> <code>text</code></td>
<td>Extracts JSON sub-object at the specified path as <code>text</code>.<br><code>'{"a": {"b": ["foo","bar"]}}'::json #&gt;&gt; '{a,b,1}'</code> <code>bar</code></td>
</tr>
</tbody>
</table>


!!! note

    The field/element/path extraction operators return NULL, rather than failing, if the JSON input does not have the right structure to match the request; for example if no such key or array element exists.


 Some further operators exist only for `jsonb`, as shown in [Additional `jsonb` Operators](#functions-jsonb-op-table). [`jsonb` Indexing](../data-types/json-types.md#json-indexing) describes how these operators can be used to effectively search indexed `jsonb` data.


<a id="functions-jsonb-op-table"></a>
**Table: Additional `jsonb` Operators**

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
<td><code>jsonb</code> <code>@&gt;</code> <code>jsonb</code> <code>boolean</code></td>
<td>Does the first JSON value contain the second? (See <a href="../data-types/json-types.md#json-containment"><code>jsonb</code> Containment and Existence</a> for details about containment.)</td>
<td><code>'{"a":1, "b":2}'::jsonb @&gt; '{"b":2}'::jsonb</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>&lt;@</code> <code>jsonb</code> <code>boolean</code></td>
<td>Is the first JSON value contained in the second?</td>
<td><code>'{"b":2}'::jsonb &lt;@ '{"a":1, "b":2}'::jsonb</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>?</code> <code>text</code> <code>boolean</code></td>
<td>Does the text string exist as a top-level key or array element within the JSON value?</td>
<td><code>'{"a":1, "b":2}'::jsonb ? 'b'</code> <code>t</code><br><code>'["a", "b", "c"]'::jsonb ? 'b'</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>?|</code> <code>text[]</code> <code>boolean</code></td>
<td>Do any of the strings in the text array exist as top-level keys or array elements?</td>
<td><code>'{"a":1, "b":2, "c":3}'::jsonb ?| array['b', 'd']</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>?&amp;</code> <code>text[]</code> <code>boolean</code></td>
<td>Do all of the strings in the text array exist as top-level keys or array elements?</td>
<td><code>'["a", "b", "c"]'::jsonb ?&amp; array['a', 'b']</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>||</code> <code>jsonb</code> <code>jsonb</code></td>
<td>Concatenates two <code>jsonb</code> values. Concatenating two arrays generates an array containing all the elements of each input. Concatenating two objects generates an object containing the union of their keys, taking the second object's value when there are duplicate keys. All other cases are treated by converting a non-array input into a single-element array, and then proceeding as for two arrays. Does not operate recursively: only the top-level array or object structure is merged.</td>
<td><code>'["a", "b"]'::jsonb || '["a", "d"]'::jsonb</code> <code>["a", "b", "a", "d"]</code><br><code>'{"a": "b"}'::jsonb || '{"c": "d"}'::jsonb</code> <code>{"a": "b", "c": "d"}</code><br><code>'[1, 2]'::jsonb || '3'::jsonb</code> <code>[1, 2, 3]</code><br><code>'{"a": "b"}'::jsonb || '42'::jsonb</code> <code>[{"a": "b"}, 42]</code><br>To append an array to another array as a single entry, wrap it in an additional layer of array, for example:<br><code>'[1, 2]'::jsonb || jsonb_build_array('[3, 4]'::jsonb)</code> <code>[1, 2, [3, 4]]</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>-</code> <code>text</code> <code>jsonb</code></td>
<td>Deletes a key (and its value) from a JSON object, or matching string value(s) from a JSON array.</td>
<td><code>'{"a": "b", "c": "d"}'::jsonb - 'a'</code> <code>{"c": "d"}</code><br><code>'["a", "b", "c", "b"]'::jsonb - 'b'</code> <code>["a", "c"]</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>-</code> <code>text[]</code> <code>jsonb</code></td>
<td>Deletes all matching keys or array elements from the left operand.</td>
<td><code>'{"a": "b", "c": "d"}'::jsonb - '{a,c}'::text[]</code> <code>{}</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>-</code> <code>integer</code> <code>jsonb</code></td>
<td>Deletes the array element with specified index (negative integers count from the end). Throws an error if JSON value is not an array.</td>
<td><code>'["a", "b"]'::jsonb - 1 </code> <code>["a"]</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>#-</code> <code>text[]</code> <code>jsonb</code></td>
<td>Deletes the field or array element at the specified path, where path elements can be either field keys or array indexes.</td>
<td><code>'["a", {"b":1}]'::jsonb #- '{1,b}'</code> <code>["a", {}]</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>@?</code> <code>jsonpath</code> <code>boolean</code></td>
<td>Does JSON path return any item for the specified JSON value? (This is useful only with SQL-standard JSON path expressions, not <a href="#functions-sqljson-check-expressions">predicate check expressions</a>, since those always return a value.)</td>
<td><code>'{"a":[1,2,3,4,5]}'::jsonb @? '$.a[*] ? (@ &gt; 2)'</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb</code> <code>@@</code> <code>jsonpath</code> <code>boolean</code></td>
<td>Returns the result of a JSON path predicate check for the specified JSON value. (This is useful only with <a href="#functions-sqljson-check-expressions">predicate check expressions</a>, not SQL-standard JSON path expressions, since it will return <code>NULL</code> if the path result is not a single boolean value.)</td>
<td><code>'{"a":[1,2,3,4,5]}'::jsonb @@ '$.a[*] &gt; 2'</code> <code>t</code></td>
</tr>
</tbody>
</table>


!!! note

    The `jsonpath` operators `@?` and `@@` suppress the following errors: missing object field or array element, unexpected JSON item type, datetime and numeric errors. The `jsonpath`-related functions described below can also be told to suppress these types of errors. This behavior might be helpful when searching JSON document collections of varying structure.


 [JSON Creation Functions](#functions-json-creation-table) shows the functions that are available for constructing `json` and `jsonb` values. Some functions in this table have a `RETURNING` clause, which specifies the data type returned. It must be one of `json`, `jsonb`, `bytea`, a character string type (`text`, `char`, or `varchar`), or a type that can be cast to `json`. By default, the `json` type is returned.


<a id="functions-json-creation-table"></a>
**Table: JSON Creation Functions**

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
<td><code>to_json</code> ( <code>anyelement</code> ) <code>json</code></td>
<td><code>to_jsonb</code> ( <code>anyelement</code> ) <code>jsonb</code></td>
<td>Converts any SQL value to <code>json</code> or <code>jsonb</code>. Arrays and composites are converted recursively to arrays and objects (multidimensional arrays become arrays of arrays in JSON). Otherwise, if there is a cast from the SQL data type to <code>json</code>, the cast function will be used to perform the conversion; (For example, the <a href="../../appendixes/additional-supplied-modules-and-extensions/hstore-hstore-key-value-datatype.md#hstore">hstore</a> extension has a cast from <code>hstore</code> to <code>json</code>, so that <code>hstore</code> values converted via the JSON creation functions will be represented as JSON objects, not as primitive string values.) otherwise, a scalar JSON value is produced. For any scalar other than a number, a Boolean, or a null value, the text representation will be used, with escaping as necessary to make it a valid JSON string value.<br><code>to_json('Fred said "Hi."'::text)</code> <code>"Fred said \"Hi.\""</code><br><code>to_jsonb(row(42, 'Fred said "Hi."'::text))</code> <code>{"f1": 42, "f2": "Fred said \"Hi.\""}</code></td>
</tr>
<tr>
<td><code>array_to_json</code> ( <code>anyarray</code> [, <code>boolean</code> ] ) <code>json</code></td>
<td>Converts an SQL array to a JSON array. The behavior is the same as <code>to_json</code> except that line feeds will be added between top-level array elements if the optional boolean parameter is true.</td>
<td><code>array_to_json('{{1,5},{99,100}}'::int[])</code> <code>[[1,5],[99,100]]</code></td>
</tr>
<tr>
<td><code>json_array</code> ( [ { <em>value_expression</em> [ <code>FORMAT JSON</code> ] } [, ...] ] [ { <code>NULL</code> | <code>ABSENT</code> } <code>ON NULL</code> ] [ <code>RETURNING</code> <em>data_type</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] ])</td>
<td><code>json_array</code> ( [ <em>query_expression</em> ] [ <code>RETURNING</code> <em>data_type</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] ])</td>
<td>Constructs a JSON array from either a series of <em>value_expression</em> parameters or from the results of <em>query_expression</em>, which must be a SELECT query returning a single column. If <code>ABSENT ON NULL</code> is specified, NULL values are ignored. This is always the case if a <em>query_expression</em> is used.<br><code>json_array(1,true,json '{"a":null}')</code> <code>[1, true, {"a":null}]</code><br><code>json_array(SELECT * FROM (VALUES(1),(2)) t)</code> <code>[1, 2]</code></td>
</tr>
<tr>
<td><code>row_to_json</code> ( <code>record</code> [, <code>boolean</code> ] ) <code>json</code></td>
<td>Converts an SQL composite value to a JSON object. The behavior is the same as <code>to_json</code> except that line feeds will be added between top-level elements if the optional boolean parameter is true.</td>
<td><code>row_to_json(row(1,'foo'))</code> <code>{"f1":1,"f2":"foo"}</code></td>
</tr>
<tr>
<td><code>json_build_array</code> ( <code>VARIADIC</code> <code>"any"</code> ) <code>json</code></td>
<td><code>jsonb_build_array</code> ( <code>VARIADIC</code> <code>"any"</code> ) <code>jsonb</code></td>
<td>Builds a possibly-heterogeneously-typed JSON array out of a variadic argument list. Each argument is converted as per <code>to_json</code> or <code>to_jsonb</code>.<br><code>json_build_array(1, 2, 'foo', 4, 5)</code> <code>[1, 2, "foo", 4, 5]</code></td>
</tr>
<tr>
<td><code>json_build_object</code> ( <code>VARIADIC</code> <code>"any"</code> ) <code>json</code></td>
<td><code>jsonb_build_object</code> ( <code>VARIADIC</code> <code>"any"</code> ) <code>jsonb</code></td>
<td>Builds a JSON object out of a variadic argument list. By convention, the argument list consists of alternating keys and values. Key arguments are coerced to text; value arguments are converted as per <code>to_json</code> or <code>to_jsonb</code>.<br><code>json_build_object('foo', 1, 2, row(3,'bar'))</code> <code>{"foo" : 1, "2" : {"f1":3,"f2":"bar"}}</code></td>
</tr>
<tr>
<td><code>json_object</code> ( [ { <em>key_expression</em> { <code>VALUE</code> | ':' } <em>value_expression</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] }[, ...] ] [ { <code>NULL</code> | <code>ABSENT</code> } <code>ON NULL</code> ] [ { <code>WITH</code> | <code>WITHOUT</code> } <code>UNIQUE</code> [ <code>KEYS</code> ] ] [ <code>RETURNING</code> <em>data_type</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] ])</td>
<td>Constructs a JSON object of all the key/value pairs given, or an empty object if none are given. <em>key_expression</em> is a scalar expression defining the JSON key, which is converted to the <code>text</code> type. It cannot be <code>NULL</code> nor can it belong to a type that has a cast to the <code>json</code> type. If <code>WITH UNIQUE KEYS</code> is specified, there must not be any duplicate <em>key_expression</em>. Any pair for which the <em>value_expression</em> evaluates to <code>NULL</code> is omitted from the output if <code>ABSENT ON NULL</code> is specified; if <code>NULL ON NULL</code> is specified or the clause omitted, the key is included with value <code>NULL</code>.</td>
<td><code>json_object('code' VALUE 'P123', 'title': 'Jaws')</code> <code>{"code" : "P123", "title" : "Jaws"}</code></td>
</tr>
<tr>
<td><code>json_object</code> ( <code>text[]</code> ) <code>json</code></td>
<td><code>jsonb_object</code> ( <code>text[]</code> ) <code>jsonb</code></td>
<td>Builds a JSON object out of a text array. The array must have either exactly one dimension with an even number of members, in which case they are taken as alternating key/value pairs, or two dimensions such that each inner array has exactly two elements, which are taken as a key/value pair. All values are converted to JSON strings.<br><code>json_object('{a, 1, b, "def", c, 3.5}')</code> <code>{"a" : "1", "b" : "def", "c" : "3.5"}</code><br><code>json_object('{{a, 1}, {b, "def"}, {c, 3.5}}')</code> <code>{"a" : "1", "b" : "def", "c" : "3.5"}</code></td>
</tr>
<tr>
<td><code>json_object</code> ( <code>keys</code> <code>text[]</code>, <code>values</code> <code>text[]</code> ) <code>json</code></td>
<td><code>jsonb_object</code> ( <code>keys</code> <code>text[]</code>, <code>values</code> <code>text[]</code> ) <code>jsonb</code></td>
<td>This form of <code>json_object</code> takes keys and values pairwise from separate text arrays. Otherwise it is identical to the one-argument form.<br><code>json_object('{a,b}', '{1,2}')</code> <code>{"a": "1", "b": "2"}</code></td>
</tr>
<tr>
<td><code>json</code> ( <em>expression</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ]] [ { <code>WITH</code> | <code>WITHOUT</code> } <code>UNIQUE</code> [ <code>KEYS</code> ]] ) <code>json</code></td>
<td>Converts a given expression specified as <code>text</code> or <code>bytea</code> string (in UTF8 encoding) into a JSON value. If <em>expression</em> is NULL, an SQL null value is returned. If <code>WITH UNIQUE</code> is specified, the <em>expression</em> must not contain any duplicate object keys.</td>
<td><code>json('{"a":123, "b":[true,"foo"], "a":"bar"}')</code> <code>{"a":123, "b":[true,"foo"], "a":"bar"}</code></td>
</tr>
<tr>
<td><code>json_scalar</code> ( <em>expression</em> )</td>
<td>Converts a given SQL scalar value into a JSON scalar value. If the input is NULL, an SQL null is returned. If the input is number or a boolean value, a corresponding JSON number or boolean value is returned. For any other value, a JSON string is returned.</td>
<td><code>json_scalar(123.45)</code> <code>123.45</code><br><code>json_scalar(CURRENT_TIMESTAMP)</code> <code>"2022-05-10T10:51:04.62128-04:00"</code></td>
</tr>
<tr>
<td><code>json_serialize</code> ( <em>expression</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] [ <code>RETURNING</code> <em>data_type</em> [ <code>FORMAT JSON</code> [ <code>ENCODING UTF8</code> ] ] ] )</td>
<td>Converts an SQL/JSON expression into a character or binary string. The <em>expression</em> can be of any JSON type, any character string type, or <code>bytea</code> in UTF8 encoding. The returned type used in <code> RETURNING</code> can be any character string type or <code>bytea</code>. The default is <code>text</code>.</td>
<td><code>json_serialize('{ "a" : 1 } ' RETURNING bytea)</code> <code>\x7b20226122203a2031207d20</code></td>
</tr>
</tbody>
</table>


 [SQL/JSON Testing Functions](#functions-sqljson-misc) details SQL/JSON facilities for testing JSON.


<a id="functions-sqljson-misc"></a>
**Table: SQL/JSON Testing Functions**

<table>
<thead>
<tr>
<th>Function signature</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>expression</em> <code>IS</code> [ <code>NOT</code> ] <code>JSON</code> [ { <code>VALUE</code> | <code>SCALAR</code> | <code>ARRAY</code> | <code>OBJECT</code> } ] [ { <code>WITH</code> | <code>WITHOUT</code> } <code>UNIQUE</code> [ <code>KEYS</code> ] ]</td>
<td>This predicate tests whether <em>expression</em> can be parsed as JSON, possibly of a specified type. If <code>SCALAR</code> or <code>ARRAY</code> or <code>OBJECT</code> is specified, the test is whether or not the JSON is of that particular type. If <code>WITH UNIQUE KEYS</code> is specified, then any object in the <em>expression</em> is also tested to see if it has duplicate keys.</td>
<td><pre><code class="language-sql">
SELECT js,
  js IS JSON "json?",
  js IS JSON SCALAR "scalar?",
  js IS JSON OBJECT "object?",
  js IS JSON ARRAY "array?"
FROM (VALUES
      ('123'), ('"abc"'), ('{"a": "b"}'), ('[1,2]'),('abc')) foo(js);
     js     | json? | scalar? | object? | array?
------------+-------+---------+---------+--------
 123        | t     | t       | f       | f
 "abc"      | t     | t       | f       | f
 {"a": "b"} | t     | f       | t       | f
 [1,2]      | t     | f       | f       | t
 abc        | f     | f       | f       | f</code></pre><br><pre><code class="language-sql">
SELECT js,
  js IS JSON OBJECT "object?",
  js IS JSON ARRAY "array?",
  js IS JSON ARRAY WITH UNIQUE KEYS "array w. UK?",
  js IS JSON ARRAY WITHOUT UNIQUE KEYS "array w/o UK?"
FROM (VALUES ('[{"a":"1"},
 {"b":"2","b":"3"}]')) foo(js);
-[ RECORD 1 ]-+--------------------
js            | [{"a":"1"},        +
              |  {"b":"2","b":"3"}]
object?       | f
array?        | t
array w. UK?  | f
array w/o UK? | t</code></pre></td>
</tr>
</tbody>
</table>


 [JSON Processing Functions](#functions-json-processing-table) shows the functions that are available for processing `json` and `jsonb` values.


<a id="functions-json-processing-table"></a>
**Table: JSON Processing Functions**

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
<td><code>json_array_elements</code> ( <code>json</code> ) <code>setof json</code></td>
<td><code>jsonb_array_elements</code> ( <code>jsonb</code> ) <code>setof jsonb</code></td>
<td>Expands the top-level JSON array into a set of JSON values.<br><p><code>SELECT * FROM json_array_elements('[1,true, [2,false]]')</code></p>
<pre><code>
   value
-----------
 1
 true
 [2,false]</code></pre></td>
</tr>
<tr>
<td><code>json_array_elements_text</code> ( <code>json</code> ) <code>setof text</code></td>
<td><code>jsonb_array_elements_text</code> ( <code>jsonb</code> ) <code>setof text</code></td>
<td>Expands the top-level JSON array into a set of <code>text</code> values.<br><p><code>SELECT * FROM json_array_elements_text('["foo", "bar"]')</code></p>
<pre><code>
   value
-----------
 foo
 bar</code></pre></td>
</tr>
<tr>
<td><code>json_array_length</code> ( <code>json</code> ) <code>integer</code></td>
<td><code>jsonb_array_length</code> ( <code>jsonb</code> ) <code>integer</code></td>
<td>Returns the number of elements in the top-level JSON array.<br><code>json_array_length('[1,2,3,{"f1":1,"f2":[5,6]},4]')</code> <code>5</code><br><code>jsonb_array_length('[]')</code> <code>0</code></td>
</tr>
<tr>
<td><code>json_each</code> ( <code>json</code> ) <code>setof record</code> ( <code>key</code> <code>text</code>, <code>value</code> <code>json</code> )</td>
<td><code>jsonb_each</code> ( <code>jsonb</code> ) <code>setof record</code> ( <code>key</code> <code>text</code>, <code>value</code> <code>jsonb</code> )</td>
<td>Expands the top-level JSON object into a set of key/value pairs.<br><p><code>SELECT * FROM json_each('{"a":"foo", "b":"bar"}')</code></p>
<pre><code>
 key | value
-----+-------
 a   | "foo"
 b   | "bar"</code></pre></td>
</tr>
<tr>
<td><code>json_each_text</code> ( <code>json</code> ) <code>setof record</code> ( <code>key</code> <code>text</code>, <code>value</code> <code>text</code> )</td>
<td><code>jsonb_each_text</code> ( <code>jsonb</code> ) <code>setof record</code> ( <code>key</code> <code>text</code>, <code>value</code> <code>text</code> )</td>
<td>Expands the top-level JSON object into a set of key/value pairs. The returned <code>value</code>s will be of type <code>text</code>.<br><p><code>SELECT * FROM json_each_text('{"a":"foo", "b":"bar"}')</code></p>
<pre><code>
 key | value
-----+-------
 a   | foo
 b   | bar</code></pre></td>
</tr>
<tr>
<td><code>json_extract_path</code> ( <code>from_json</code> <code>json</code>, <code>VARIADIC</code> <code>path_elems</code> <code>text[]</code> ) <code>json</code></td>
<td><code>jsonb_extract_path</code> ( <code>from_json</code> <code>jsonb</code>, <code>VARIADIC</code> <code>path_elems</code> <code>text[]</code> ) <code>jsonb</code></td>
<td>Extracts JSON sub-object at the specified path. (This is functionally equivalent to the <code>#&gt;</code> operator, but writing the path out as a variadic list can be more convenient in some cases.)<br><code>json_extract_path('{"f2":{"f3":1},"f4":{"f5":99,"f6":"foo"}}', 'f4', 'f6')</code> <code>"foo"</code></td>
</tr>
<tr>
<td><code>json_extract_path_text</code> ( <code>from_json</code> <code>json</code>, <code>VARIADIC</code> <code>path_elems</code> <code>text[]</code> ) <code>text</code></td>
<td><code>jsonb_extract_path_text</code> ( <code>from_json</code> <code>jsonb</code>, <code>VARIADIC</code> <code>path_elems</code> <code>text[]</code> ) <code>text</code></td>
<td>Extracts JSON sub-object at the specified path as <code>text</code>. (This is functionally equivalent to the <code>#&gt;&gt;</code> operator.)<br><code>json_extract_path_text('{"f2":{"f3":1},"f4":{"f5":99,"f6":"foo"}}', 'f4', 'f6')</code> <code>foo</code></td>
</tr>
<tr>
<td><code>json_object_keys</code> ( <code>json</code> ) <code>setof text</code></td>
<td><code>jsonb_object_keys</code> ( <code>jsonb</code> ) <code>setof text</code></td>
<td>Returns the set of keys in the top-level JSON object.<br><p><code>SELECT * FROM json_object_keys('{"f1":"abc","f2":{"f3":"a", "f4":"b"}}')</code></p>
<pre><code>
 json_object_keys
------------------
 f1
 f2</code></pre></td>
</tr>
<tr>
<td><code>json_populate_record</code> ( <code>base</code> <code>anyelement</code>, <code>from_json</code> <code>json</code> ) <code>anyelement</code></td>
<td><code>jsonb_populate_record</code> ( <code>base</code> <code>anyelement</code>, <code>from_json</code> <code>jsonb</code> ) <code>anyelement</code></td>
<td>Expands the top-level JSON object to a row having the composite type of the <code>base</code> argument. The JSON object is scanned for fields whose names match column names of the output row type, and their values are inserted into those columns of the output. (Fields that do not correspond to any output column name are ignored.) In typical use, the value of <code>base</code> is just <code>NULL</code>, which means that any output columns that do not match any object field will be filled with nulls. However, if <code>base</code> isn't <code>NULL</code> then the values it contains will be used for unmatched columns.<br><p>To convert a JSON value to the SQL type of an output column, the following rules are applied in sequence:</p>
<p>-  A JSON null value is converted to an SQL null in all cases. <br>
-  If the output column is of type <code>json</code> or <code>jsonb</code>, the JSON value is just reproduced exactly. <br>
-  If the output column is a composite (row) type, and the JSON value is a JSON object, the fields of the object are converted to columns of the output row type by recursive application of these rules. <br>
-  Likewise, if the output column is an array type and the JSON value is a JSON array, the elements of the JSON array are converted to elements of the output array by recursive application of these rules. <br>
-  Otherwise, if the JSON value is a string, the contents of the string are fed to the input conversion function for the column's data type. <br>
-  Otherwise, the ordinary text representation of the JSON value is fed to the input conversion function for the column's data type.</p><br>While the example below uses a constant JSON value, typical use would be to reference a <code>json</code> or <code>jsonb</code> column laterally from another table in the query's <code>FROM</code> clause. Writing <code>json_populate_record</code> in the <code>FROM</code> clause is good practice, since all of the extracted columns are available for use without duplicate function calls.<br><code>CREATE TYPE subrowtype AS (d int, e text);</code> <code>CREATE TYPE myrowtype AS (a int, b text[], c subrowtype);</code><br><p><code>SELECT * FROM json_populate_record(NULL::myrowtype, '{"a": 1, "b": ["2", "a b"], "c": {"d": 4, "e": "a b c"}, "x": "foo"}')</code></p>
<pre><code>
 a |   b       |      c
---+-----------+-------------
 1 | {2,"a b"} | (4,"a b c")</code></pre></td>
</tr>
<tr>
<td><code>jsonb_populate_record_valid</code> ( <code>base</code> <code>anyelement</code>, <code>from_json</code> <code>json</code> ) <code>boolean</code></td>
<td>Function for testing <code>jsonb_populate_record</code>. Returns <code>true</code> if the input <code>jsonb_populate_record</code> would finish without an error for the given input JSON object; that is, it's valid input, <code>false</code> otherwise.</td>
<td><code>CREATE TYPE jsb_char2 AS (a char(2));</code><br><p><code>SELECT jsonb_populate_record_valid(NULL::jsb_char2, '{"a": "aaa"}');</code></p>
<pre><code>
 jsonb_populate_record_valid
-----------------------------
 f
(1 row)</code></pre>
<pre><code>
ERROR:  value too long for type character(2)</code></pre>
<pre><code>
 jsonb_populate_record_valid
-----------------------------
 t
(1 row)</code></pre>
<pre><code>
 a
----
 aa
(1 row)</code></pre></td>
</tr>
<tr>
<td><code>json_populate_recordset</code> ( <code>base</code> <code>anyelement</code>, <code>from_json</code> <code>json</code> ) <code>setof anyelement</code></td>
<td><code>jsonb_populate_recordset</code> ( <code>base</code> <code>anyelement</code>, <code>from_json</code> <code>jsonb</code> ) <code>setof anyelement</code></td>
<td>Expands the top-level JSON array of objects to a set of rows having the composite type of the <code>base</code> argument. Each element of the JSON array is processed as described above for <code>json[b]_populate_record</code>.<br><code>CREATE TYPE twoints AS (a int, b int);</code><br><p><code>SELECT * FROM json_populate_recordset(NULL::twoints, '[{"a":1,"b":2}, {"a":3,"b":4}]')</code></p>
<pre><code>
 a | b
---+---
 1 | 2
 3 | 4</code></pre></td>
</tr>
<tr>
<td><code>json_to_record</code> ( <code>json</code> ) <code>record</code></td>
<td><code>jsonb_to_record</code> ( <code>jsonb</code> ) <code>record</code></td>
<td>Expands the top-level JSON object to a row having the composite type defined by an <code>AS</code> clause. (As with all functions returning <code>record</code>, the calling query must explicitly define the structure of the record with an <code>AS</code> clause.) The output record is filled from fields of the JSON object, in the same way as described above for <code>json[b]_populate_record</code>. Since there is no input record value, unmatched columns are always filled with nulls.<br><code>CREATE TYPE myrowtype AS (a int, b text);</code><br><p><code>SELECT * FROM json_to_record('{"a":1,"b":[1,2,3],"c":[1,2,3],"e":"bar","r": {"a": 123, "b": "a b c"}}') AS x(a int, b text, c int[], d text, r myrowtype)</code></p>
<pre><code>
 a |    b    |    c    | d |       r
---+---------+---------+---+---------------
 1 | [1,2,3] | {1,2,3} |   | (123,"a b c")</code></pre></td>
</tr>
<tr>
<td><code>json_to_recordset</code> ( <code>json</code> ) <code>setof record</code></td>
<td><code>jsonb_to_recordset</code> ( <code>jsonb</code> ) <code>setof record</code></td>
<td>Expands the top-level JSON array of objects to a set of rows having the composite type defined by an <code>AS</code> clause. (As with all functions returning <code>record</code>, the calling query must explicitly define the structure of the record with an <code>AS</code> clause.) Each element of the JSON array is processed as described above for <code>json[b]_populate_record</code>.<br><p><code>SELECT * FROM json_to_recordset('[{"a":1,"b":"foo"}, {"a":"2","c":"bar"}]') AS x(a int, b text)</code></p>
<pre><code>
 a |  b
---+-----
 1 | foo
 2 |</code></pre></td>
</tr>
<tr>
<td><code>jsonb_set</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>text[]</code>, <code>new_value</code> <code>jsonb</code> [, <code>create_if_missing</code> <code>boolean</code> ] ) <code>jsonb</code></td>
<td>Returns <code>target</code> with the item designated by <code>path</code> replaced by <code>new_value</code>, or with <code>new_value</code> added if <code>create_if_missing</code> is true (which is the default) and the item designated by <code>path</code> does not exist. All earlier steps in the path must exist, or the <code>target</code> is returned unchanged. As with the path oriented operators, negative integers that appear in the <code>path</code> count from the end of JSON arrays. If the last path step is an array index that is out of range, and <code>create_if_missing</code> is true, the new value is added at the beginning of the array if the index is negative, or at the end of the array if it is positive.</td>
<td><code>jsonb_set('[{"f1":1,"f2":null},2,null,3]', '{0,f1}', '[2,3,4]', false)</code> <code>[{"f1": [2, 3, 4], "f2": null}, 2, null, 3]</code><br><code>jsonb_set('[{"f1":1,"f2":null},2]', '{0,f3}', '[2,3,4]')</code> <code>[{"f1": 1, "f2": null, "f3": [2, 3, 4]}, 2]</code></td>
</tr>
<tr>
<td><code>jsonb_set_lax</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>text[]</code>, <code>new_value</code> <code>jsonb</code> [, <code>create_if_missing</code> <code>boolean</code> [, <code>null_value_treatment</code> <code>text</code> ]] ) <code>jsonb</code></td>
<td>If <code>new_value</code> is not <code>NULL</code>, behaves identically to <code>jsonb_set</code>. Otherwise behaves according to the value of <code>null_value_treatment</code> which must be one of <code>'raise_exception'</code>, <code>'use_json_null'</code>, <code>'delete_key'</code>, or <code>'return_target'</code>. The default is <code>'use_json_null'</code>.</td>
<td><code>jsonb_set_lax('[{"f1":1,"f2":null},2,null,3]', '{0,f1}', null)</code> <code>[{"f1": null, "f2": null}, 2, null, 3]</code><br><code>jsonb_set_lax('[{"f1":99,"f2":null},2]', '{0,f3}', null, true, 'return_target')</code> <code>[{"f1": 99, "f2": null}, 2]</code></td>
</tr>
<tr>
<td><code>jsonb_insert</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>text[]</code>, <code>new_value</code> <code>jsonb</code> [, <code>insert_after</code> <code>boolean</code> ] ) <code>jsonb</code></td>
<td>Returns <code>target</code> with <code>new_value</code> inserted. If the item designated by the <code>path</code> is an array element, <code>new_value</code> will be inserted before that item if <code>insert_after</code> is false (which is the default), or after it if <code>insert_after</code> is true. If the item designated by the <code>path</code> is an object field, <code>new_value</code> will be inserted only if the object does not already contain that key. All earlier steps in the path must exist, or the <code>target</code> is returned unchanged. As with the path oriented operators, negative integers that appear in the <code>path</code> count from the end of JSON arrays. If the last path step is an array index that is out of range, the new value is added at the beginning of the array if the index is negative, or at the end of the array if it is positive.</td>
<td><code>jsonb_insert('{"a": [0,1,2]}', '{a, 1}', '"new_value"')</code> <code>{"a": [0, "new_value", 1, 2]}</code><br><code>jsonb_insert('{"a": [0,1,2]}', '{a, 1}', '"new_value"', true)</code> <code>{"a": [0, 1, "new_value", 2]}</code></td>
</tr>
<tr>
<td><code>json_strip_nulls</code> ( <code>target</code> <code>json</code> [,<code>strip_in_arrays</code> <code>boolean</code> ] ) <code>json</code></td>
<td><code>jsonb_strip_nulls</code> ( <code>target</code> <code>jsonb</code> [,<code>strip_in_arrays</code> <code>boolean</code> ] ) <code>jsonb</code></td>
<td>Deletes all object fields that have null values from the given JSON value, recursively. If <code>strip_in_arrays</code> is true (the default is false), null array elements are also stripped. Otherwise they are not stripped. Bare null values are never stripped.<br><code>json_strip_nulls('[{"f1":1, "f2":null}, 2, null, 3]')</code> <code>[{"f1":1},2,null,3]</code><br><code>jsonb_strip_nulls('[1,2,null,3,4]', true)</code> <code>[1,2,3,4]</code></td>
</tr>
<tr>
<td><code>jsonb_path_exists</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>boolean</code></td>
<td>Checks whether the JSON path returns any item for the specified JSON value. (This is useful only with SQL-standard JSON path expressions, not <a href="#functions-sqljson-check-expressions">predicate check expressions</a>, since those always return a value.) If the <code>vars</code> argument is specified, it must be a JSON object, and its fields provide named values to be substituted into the <code>jsonpath</code> expression. If the <code>silent</code> argument is specified and is <code>true</code>, the function suppresses the same errors as the <code>@?</code> and <code>@@</code> operators do.</td>
<td><code>jsonb_path_exists('{"a":[1,2,3,4,5]}', '$.a[*] ? (@ &gt;= $min &amp;&amp; @ &lt;= $max)', '{"min":2, "max":4}')</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb_path_match</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>boolean</code></td>
<td>Returns the SQL boolean result of a JSON path predicate check for the specified JSON value. (This is useful only with <a href="#functions-sqljson-check-expressions">predicate check expressions</a>, not SQL-standard JSON path expressions, since it will either fail or return <code>NULL</code> if the path result is not a single boolean value.) The optional <code>vars</code> and <code>silent</code> arguments act the same as for <code>jsonb_path_exists</code>.</td>
<td><code>jsonb_path_match('{"a":[1,2,3,4,5]}', 'exists($.a[*] ? (@ &gt;= $min &amp;&amp; @ &lt;= $max))', '{"min":2, "max":4}')</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb_path_query</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>setof jsonb</code></td>
<td>Returns all JSON items returned by the JSON path for the specified JSON value. For SQL-standard JSON path expressions it returns the JSON values selected from <code>target</code>. For <a href="#functions-sqljson-check-expressions">predicate check expressions</a> it returns the result of the predicate check: <code>true</code>, <code>false</code>, or <code>null</code>. The optional <code>vars</code> and <code>silent</code> arguments act the same as for <code>jsonb_path_exists</code>.</td>
<td><p><code>SELECT <em> FROM jsonb_path_query('{"a":[1,2,3,4,5]}', '$.a[</em>] ? (@ &gt;= $min &amp;&amp; @ &lt;= $max)', '{"min":2, "max":4}')</code></p>
<pre><code>
 jsonb_path_query
------------------
 2
 3
 4</code></pre></td>
</tr>
<tr>
<td><code>jsonb_path_query_array</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>jsonb</code></td>
<td>Returns all JSON items returned by the JSON path for the specified JSON value, as a JSON array. The parameters are the same as for <code>jsonb_path_query</code>.</td>
<td><code>jsonb_path_query_array('{"a":[1,2,3,4,5]}', '$.a[*] ? (@ &gt;= $min &amp;&amp; @ &lt;= $max)', '{"min":2, "max":4}')</code> <code>[2, 3, 4]</code></td>
</tr>
<tr>
<td><code>jsonb_path_query_first</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>jsonb</code></td>
<td>Returns the first JSON item returned by the JSON path for the specified JSON value, or <code>NULL</code> if there are no results. The parameters are the same as for <code>jsonb_path_query</code>.</td>
<td><code>jsonb_path_query_first('{"a":[1,2,3,4,5]}', '$.a[*] ? (@ &gt;= $min &amp;&amp; @ &lt;= $max)', '{"min":2, "max":4}')</code> <code>2</code></td>
</tr>
<tr>
<td><code>jsonb_path_exists_tz</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>boolean</code></td>
<td><code>jsonb_path_match_tz</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>boolean</code></td>
<td><code>jsonb_path_query_tz</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>setof jsonb</code><br><code>jsonb_path_query_array_tz</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>jsonb</code><br><code>jsonb_path_query_first_tz</code> ( <code>target</code> <code>jsonb</code>, <code>path</code> <code>jsonpath</code> [, <code>vars</code> <code>jsonb</code> [, <code>silent</code> <code>boolean</code> ]] ) <code>jsonb</code><br>These functions act like their counterparts described above without the <code>_tz</code> suffix, except that these functions support comparisons of date/time values that require timezone-aware conversions. The example below requires interpretation of the date-only value <code>2015-08-02</code> as a timestamp with time zone, so the result depends on the current <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone">TimeZone</a> setting. Due to this dependency, these functions are marked as stable, which means these functions cannot be used in indexes. Their counterparts are immutable, and so can be used in indexes; but they will throw errors if asked to make such comparisons.<br><code>jsonb_path_exists_tz('["2015-08-01 12:00:00-05"]', '$[*] ? (@.datetime() &lt; "2015-08-02".datetime())')</code> <code>t</code></td>
</tr>
<tr>
<td><code>jsonb_pretty</code> ( <code>jsonb</code> ) <code>text</code></td>
<td>Converts the given JSON value to pretty-printed, indented text.</td>
<td><p><code>jsonb_pretty('[{"f1":1,"f2":null}, 2]')</code></p>
<pre><code>
[
    {
        "f1": 1,
        "f2": null
    },
    2
]</code></pre></td>
</tr>
<tr>
<td><code>json_typeof</code> ( <code>json</code> ) <code>text</code></td>
<td><code>jsonb_typeof</code> ( <code>jsonb</code> ) <code>text</code></td>
<td>Returns the type of the top-level JSON value as a text string. Possible types are <code>object</code>, <code>array</code>, <code>string</code>, <code>number</code>, <code>boolean</code>, and <code>null</code>. (The <code>null</code> result should not be confused with an SQL NULL; see the examples.)<br><code>json_typeof('-123.4')</code> <code>number</code><br><code>json_typeof('null'::json)</code> <code>null</code><br><code>json_typeof(NULL::json) IS NULL</code> <code>t</code></td>
</tr>
</tbody>
</table>


### The SQL/JSON Path Language { #functions-sqljson-path }


 SQL/JSON path expressions specify item(s) to be retrieved from a JSON value, similarly to XPath expressions used for access to XML content. In PostgreSQL, path expressions are implemented as the `jsonpath` data type and can use any elements described in [jsonpath Type](../data-types/json-types.md#datatype-jsonpath).


 JSON query functions and operators pass the provided path expression to the *path engine* for evaluation. If the expression matches the queried JSON data, the corresponding JSON item, or set of items, is returned. If there is no match, the result will be `NULL`, `false`, or an error, depending on the function. Path expressions are written in the SQL/JSON path language and can include arithmetic expressions and functions.


 A path expression consists of a sequence of elements allowed by the `jsonpath` data type. The path expression is normally evaluated from left to right, but you can use parentheses to change the order of operations. If the evaluation is successful, a sequence of JSON items is produced, and the evaluation result is returned to the JSON query function that completes the specified computation.


 To refer to the JSON value being queried (the *context item*), use the `$` variable in the path expression. The first element of a path must always be `$`. It can be followed by one or more [accessor operators](../data-types/json-types.md#type-jsonpath-accessors), which go down the JSON structure level by level to retrieve sub-items of the context item. Each accessor operator acts on the result(s) of the previous evaluation step, producing zero, one, or more output items from each input item.


 For example, suppose you have some JSON data from a GPS tracker that you would like to parse, such as:

```sql

SELECT '{
  "track": {
    "segments": [
      {
        "location":   [ 47.763, 13.4034 ],
        "start time": "2018-10-14 10:05:14",
        "HR": 73
      },
      {
        "location":   [ 47.706, 13.2635 ],
        "start time": "2018-10-14 10:39:21",
        "HR": 135
      }
    ]
  }
}' AS json \gset
```
 (The above example can be copied-and-pasted into psql to set things up for the following examples. Then psql will expand `:'json'` into a suitably-quoted string constant containing the JSON value.)


 To retrieve the available track segments, you need to use the <code>.</code><em>key</em> accessor operator to descend through surrounding JSON objects, for example:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments');
                                                                         jsonb_path_query
-----------------------------------------------------------​-----------------------------------------------------------​---------------------------------------------
 [{"HR": 73, "location": [47.763, 13.4034], "start time": "2018-10-14 10:05:14"}, {"HR": 135, "location": [47.706, 13.2635], "start time": "2018-10-14 10:39:21"}]
```


 To retrieve the contents of an array, you typically use the `[*]` operator. The following example will return the location coordinates for all the available track segments:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[*].location');
 jsonb_path_query
-------------------
 [47.763, 13.4034]
 [47.706, 13.2635]
```
 Here we started with the whole JSON input value (`$`), then the `.track` accessor selected the JSON object associated with the `"track"` object key, then the `.segments` accessor selected the JSON array associated with the `"segments"` key within that object, then the `[*]` accessor selected each element of that array (producing a series of items), then the `.location` accessor selected the JSON array associated with the `"location"` key within each of those objects. In this example, each of those objects had a `"location"` key; but if any of them did not, the `.location` accessor would have simply produced no output for that input item.


 To return the coordinates of the first segment only, you can specify the corresponding subscript in the `[]` accessor operator. Recall that JSON array indexes are 0-relative:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[0].location');
 jsonb_path_query
-------------------
 [47.763, 13.4034]
```


 The result of each path evaluation step can be processed by one or more of the `jsonpath` operators and methods listed in [SQL/JSON Path Operators and Methods](#functions-sqljson-path-operators). Each method name must be preceded by a dot. For example, you can get the size of an array:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments.size()');
 jsonb_path_query
------------------
 2
```
 More examples of using `jsonpath` operators and methods within path expressions appear below in [SQL/JSON Path Operators and Methods](#functions-sqljson-path-operators).


 A path can also contain *filter expressions* that work similarly to the `WHERE` clause in SQL. A filter expression begins with a question mark and provides a condition in parentheses:

```

? (CONDITION)
```


 Filter expressions must be written just after the path evaluation step to which they should apply. The result of that step is filtered to include only those items that satisfy the provided condition. SQL/JSON defines three-valued logic, so the condition can produce `true`, `false`, or `unknown`. The `unknown` value plays the same role as SQL `NULL` and can be tested for with the `IS UNKNOWN` predicate. Further path evaluation steps use only those items for which the filter expression returned `true`.


 The functions and operators that can be used in filter expressions are listed in [`jsonpath` Filter Expression Elements](#functions-sqljson-filter-ex-table). Within a filter expression, the `@` variable denotes the value being considered (i.e., one result of the preceding path step). You can write accessor operators after `@` to retrieve component items.


 For example, suppose you would like to retrieve all heart rate values higher than 130. You can achieve this as follows:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[*].HR ? (@ > 130)');
 jsonb_path_query
------------------
 135
```


 To get the start times of segments with such values, you have to filter out irrelevant segments before selecting the start times, so the filter expression is applied to the previous step, and the path used in the condition is different:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[*] ? (@.HR > 130)."start time"');
   jsonb_path_query
-----------------------
 "2018-10-14 10:39:21"
```


 You can use several filter expressions in sequence, if required. The following example selects start times of all segments that contain locations with relevant coordinates and high heart rate values:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[*] ? (@.location[1] < 13.4) ? (@.HR > 130)."start time"');
   jsonb_path_query
-----------------------
 "2018-10-14 10:39:21"
```


 Using filter expressions at different nesting levels is also allowed. The following example first filters all segments by location, and then returns high heart rate values for these segments, if available:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[*] ? (@.location[1] < 13.4).HR ? (@ > 130)');
 jsonb_path_query
------------------
 135
```


 You can also nest filter expressions within each other. This example returns the size of the track if it contains any segments with high heart rate values, or an empty sequence otherwise:

```

=> SELECT jsonb_path_query(:'json', '$.track ? (exists(@.segments[*] ? (@.HR > 130))).segments.size()');
 jsonb_path_query
------------------
 2
```


#### Deviations from the SQL Standard { #functions-sqljson-deviations }


 PostgreSQL's implementation of the SQL/JSON path language has the following deviations from the SQL/JSON standard.


##### Boolean Predicate Check Expressions { #functions-sqljson-check-expressions }


 As an extension to the SQL standard, a PostgreSQL path expression can be a Boolean predicate, whereas the SQL standard allows predicates only within filters. While SQL-standard path expressions return the relevant element(s) of the queried JSON value, predicate check expressions return the single three-valued `jsonb` result of the predicate: `true`, `false`, or `null`. For example, we could write this SQL-standard filter expression:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments ?(@[*].HR > 130)');
                                jsonb_path_query
-----------------------------------------------------------​----------------------
 {"HR": 135, "location": [47.706, 13.2635], "start time": "2018-10-14 10:39:21"}
```
 The similar predicate check expression simply returns `true`, indicating that a match exists:

```

=> SELECT jsonb_path_query(:'json', '$.track.segments[*].HR > 130');
 jsonb_path_query
------------------
 true
```


!!! note

    Predicate check expressions are required in the `@@` operator (and the `jsonb_path_match` function), and should not be used with the `@?` operator (or the `jsonb_path_exists` function).


##### Regular Expression Interpretation { #functions-sqljson-regular-expression-deviation }


 There are minor differences in the interpretation of regular expression patterns used in `like_regex` filters, as described in [SQL/JSON Regular Expressions](#jsonpath-regular-expressions).


#### Strict and Lax Modes { #functions-sqljson-strict-and-lax-modes }


 When you query JSON data, the path expression may not match the actual JSON data structure. An attempt to access a non-existent member of an object or element of an array is defined as a structural error. SQL/JSON path expressions have two modes of handling structural errors:


-  lax (default) — the path engine implicitly adapts the queried data to the specified path. Any structural errors that cannot be fixed as described below are suppressed, producing no match.
-  strict — if a structural error occurs, an error is raised.


 Lax mode facilitates matching of a JSON document and path expression when the JSON data does not conform to the expected schema. If an operand does not match the requirements of a particular operation, it can be automatically wrapped as an SQL/JSON array, or unwrapped by converting its elements into an SQL/JSON sequence before performing the operation. Also, comparison operators automatically unwrap their operands in lax mode, so you can compare SQL/JSON arrays out-of-the-box. An array of size 1 is considered equal to its sole element. Automatic unwrapping is not performed when:

-  The path expression contains `type()` or `size()` methods that return the type and the number of elements in the array, respectively.
-  The queried JSON data contain nested arrays. In this case, only the outermost array is unwrapped, while all the inner arrays remain unchanged. Thus, implicit unwrapping can only go one level down within each path evaluation step.


 For example, when querying the GPS data listed above, you can abstract from the fact that it stores an array of segments when using lax mode:

```

=> SELECT jsonb_path_query(:'json', 'lax $.track.segments.location');
 jsonb_path_query
-------------------
 [47.763, 13.4034]
 [47.706, 13.2635]
```


 In strict mode, the specified path must exactly match the structure of the queried JSON document, so using this path expression will cause an error:

```

=> SELECT jsonb_path_query(:'json', 'strict $.track.segments.location');
ERROR:  jsonpath member accessor can only be applied to an object
```
 To get the same result as in lax mode, you have to explicitly unwrap the `segments` array:

```

=> SELECT jsonb_path_query(:'json', 'strict $.track.segments[*].location');
 jsonb_path_query
-------------------
 [47.763, 13.4034]
 [47.706, 13.2635]
```


 The unwrapping behavior of lax mode can lead to surprising results. For instance, the following query using the `.**` accessor selects every `HR` value twice:

```

=> SELECT jsonb_path_query(:'json', 'lax $.**.HR');
 jsonb_path_query
------------------
 73
 135
 73
 135
```
 This happens because the `.**` accessor selects both the `segments` array and each of its elements, while the `.HR` accessor automatically unwraps arrays when using lax mode. To avoid surprising results, we recommend using the `.**` accessor only in strict mode. The following query selects each `HR` value just once:

```

=> SELECT jsonb_path_query(:'json', 'strict $.**.HR');
 jsonb_path_query
------------------
 73
 135
```


 The unwrapping of arrays can also lead to unexpected results. Consider this example, which selects all the `location` arrays:

```

=> SELECT jsonb_path_query(:'json', 'lax $.track.segments[*].location');
 jsonb_path_query
-------------------
 [47.763, 13.4034]
 [47.706, 13.2635]
(2 rows)
```
 As expected it returns the full arrays. But applying a filter expression causes the arrays to be unwrapped to evaluate each item, returning only the items that match the expression:

```

=> SELECT jsonb_path_query(:'json', 'lax $.track.segments[*].location ?(@[*] > 15)');
 jsonb_path_query
------------------
 47.763
 47.706
(2 rows)
```
 This despite the fact that the full arrays are selected by the path expression. Use strict mode to restore selecting the arrays:

```

=> SELECT jsonb_path_query(:'json', 'strict $.track.segments[*].location ?(@[*] > 15)');
 jsonb_path_query
-------------------
 [47.763, 13.4034]
 [47.706, 13.2635]
(2 rows)
```


#### SQL/JSON Path Operators and Methods { #functions-sqljson-path-operators }


 [`jsonpath` Operators and Methods](#functions-sqljson-op-table) shows the operators and methods available in `jsonpath`. Note that while the unary operators and methods can be applied to multiple values resulting from a preceding path step, the binary operators (addition etc.) can only be applied to single values. In lax mode, methods applied to an array will be executed for each value in the array. The exceptions are `.type()` and `.size()`, which apply to the array itself.


<a id="functions-sqljson-op-table"></a>
**Table: `jsonpath` Operators and Methods**

<table>
<thead>
<tr>
<th>Operator/Method</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>number</em> <code>+</code> <em>number</em> <em>number</em></td>
<td>Addition</td>
<td><code>jsonb_path_query('[2]', '$[0] + 3')</code> <code>5</code></td>
</tr>
<tr>
<td><code>+</code> <em>number</em> <em>number</em></td>
<td>Unary plus (no operation); unlike addition, this can iterate over multiple values</td>
<td><code>jsonb_path_query_array('{"x": [2,3,4]}', '+ $.x')</code> <code>[2, 3, 4]</code></td>
</tr>
<tr>
<td><em>number</em> <code>-</code> <em>number</em> <em>number</em></td>
<td>Subtraction</td>
<td><code>jsonb_path_query('[2]', '7 - $[0]')</code> <code>5</code></td>
</tr>
<tr>
<td><code>-</code> <em>number</em> <em>number</em></td>
<td>Negation; unlike subtraction, this can iterate over multiple values</td>
<td><code>jsonb_path_query_array('{"x": [2,3,4]}', '- $.x')</code> <code>[-2, -3, -4]</code></td>
</tr>
<tr>
<td><em>number</em> <code><em></code> </em>number* <em>number</em></td>
<td>Multiplication</td>
<td><code>jsonb_path_query('[4]', '2 * $[0]')</code> <code>8</code></td>
</tr>
<tr>
<td><em>number</em> <code>/</code> <em>number</em> <em>number</em></td>
<td>Division</td>
<td><code>jsonb_path_query('[8.5]', '$[0] / 2')</code> <code>4.2500000000000000</code></td>
</tr>
<tr>
<td><em>number</em> <code>%</code> <em>number</em> <em>number</em></td>
<td>Modulo (remainder)</td>
<td><code>jsonb_path_query('[32]', '$[0] % 10')</code> <code>2</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>type()</code> <em>string</em></td>
<td>Type of the JSON item (see <code>json_typeof</code>)</td>
<td><code>jsonb_path_query_array('[1, "2", {}]', '$[*].type()')</code> <code>["number", "string", "object"]</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>size()</code> <em>number</em></td>
<td>Size of the JSON item (number of array elements, or 1 if not an array)</td>
<td><code>jsonb_path_query('{"m": [11, 15]}', '$.m.size()')</code> <code>2</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>boolean()</code> <em>boolean</em></td>
<td>Boolean value converted from a JSON boolean, number, or string</td>
<td><code>jsonb_path_query_array('[1, "yes", false]', '$[*].boolean()')</code> <code>[true, true, false]</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>string()</code> <em>string</em></td>
<td>String value converted from a JSON boolean, number, string, or datetime</td>
<td><code>jsonb_path_query_array('[1.23, "xyz", false]', '$[*].string()')</code> <code>["1.23", "xyz", "false"]</code><br><code>jsonb_path_query('"2023-08-15 12:34:56"', '$.timestamp().string()')</code> <code>"2023-08-15T12:34:56"</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>double()</code> <em>number</em></td>
<td>Approximate floating-point number converted from a JSON number or string</td>
<td><code>jsonb_path_query('{"len": "1.9"}', '$.len.double() * 2')</code> <code>3.8</code></td>
</tr>
<tr>
<td><em>number</em> <code>.</code> <code>ceiling()</code> <em>number</em></td>
<td>Nearest integer greater than or equal to the given number</td>
<td><code>jsonb_path_query('{"h": 1.3}', '$.h.ceiling()')</code> <code>2</code></td>
</tr>
<tr>
<td><em>number</em> <code>.</code> <code>floor()</code> <em>number</em></td>
<td>Nearest integer less than or equal to the given number</td>
<td><code>jsonb_path_query('{"h": 1.7}', '$.h.floor()')</code> <code>1</code></td>
</tr>
<tr>
<td><em>number</em> <code>.</code> <code>abs()</code> <em>number</em></td>
<td>Absolute value of the given number</td>
<td><code>jsonb_path_query('{"z": -0.3}', '$.z.abs()')</code> <code>0.3</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>bigint()</code> <em>bigint</em></td>
<td>Big integer value converted from a JSON number or string</td>
<td><code>jsonb_path_query('{"len": "9876543219"}', '$.len.bigint()')</code> <code>9876543219</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>decimal( [ </code><em>precision</em><code> [ , </code><em>scale</em><code> ] ] )</code> <em>decimal</em></td>
<td>Rounded decimal value converted from a JSON number or string (<code>precision</code> and <code>scale</code> must be integer values)</td>
<td><code>jsonb_path_query('1234.5678', '$.decimal(6, 2)')</code> <code>1234.57</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>integer()</code> <em>integer</em></td>
<td>Integer value converted from a JSON number or string</td>
<td><code>jsonb_path_query('{"len": "12345"}', '$.len.integer()')</code> <code>12345</code></td>
</tr>
<tr>
<td><em>value</em> <code>.</code> <code>number()</code> <em>numeric</em></td>
<td>Numeric value converted from a JSON number or string</td>
<td><code>jsonb_path_query('{"len": "123.45"}', '$.len.number()')</code> <code>123.45</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>datetime()</code> <em>datetime_type</em> (see note)</td>
<td>Date/time value converted from a string</td>
<td><code>jsonb_path_query('["2015-8-1", "2015-08-12"]', '$[*] ? (@.datetime() &lt; "2015-08-2".datetime())')</code> <code>"2015-8-1"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>datetime(</code><em>template</em><code>)</code> <em>datetime_type</em> (see note)</td>
<td>Date/time value converted from a string using the specified <code>to_timestamp</code> template</td>
<td><code>jsonb_path_query_array('["12:30", "18:40"]', '$[*].datetime("HH24:MI")')</code> <code>["12:30:00", "18:40:00"]</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>date()</code> <em>date</em></td>
<td>Date value converted from a string</td>
<td><code>jsonb_path_query('"2023-08-15"', '$.date()')</code> <code>"2023-08-15"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>time()</code> <em>time without time zone</em></td>
<td>Time without time zone value converted from a string</td>
<td><code>jsonb_path_query('"12:34:56"', '$.time()')</code> <code>"12:34:56"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>time(</code><em>precision</em><code>)</code> <em>time without time zone</em></td>
<td>Time without time zone value converted from a string, with fractional seconds adjusted to the given precision</td>
<td><code>jsonb_path_query('"12:34:56.789"', '$.time(2)')</code> <code>"12:34:56.79"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>time_tz()</code> <em>time with time zone</em></td>
<td>Time with time zone value converted from a string</td>
<td><code>jsonb_path_query('"12:34:56 +05:30"', '$.time_tz()')</code> <code>"12:34:56+05:30"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>time_tz(</code><em>precision</em><code>)</code> <em>time with time zone</em></td>
<td>Time with time zone value converted from a string, with fractional seconds adjusted to the given precision</td>
<td><code>jsonb_path_query('"12:34:56.789 +05:30"', '$.time_tz(2)')</code> <code>"12:34:56.79+05:30"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>timestamp()</code> <em>timestamp without time zone</em></td>
<td>Timestamp without time zone value converted from a string</td>
<td><code>jsonb_path_query('"2023-08-15 12:34:56"', '$.timestamp()')</code> <code>"2023-08-15T12:34:56"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>timestamp(</code><em>precision</em><code>)</code> <em>timestamp without time zone</em></td>
<td>Timestamp without time zone value converted from a string, with fractional seconds adjusted to the given precision</td>
<td><code>jsonb_path_query('"2023-08-15 12:34:56.789"', '$.timestamp(2)')</code> <code>"2023-08-15T12:34:56.79"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>timestamp_tz()</code> <em>timestamp with time zone</em></td>
<td>Timestamp with time zone value converted from a string</td>
<td><code>jsonb_path_query('"2023-08-15 12:34:56 +05:30"', '$.timestamp_tz()')</code> <code>"2023-08-15T12:34:56+05:30"</code></td>
</tr>
<tr>
<td><em>string</em> <code>.</code> <code>timestamp_tz(</code><em>precision</em><code>)</code> <em>timestamp with time zone</em></td>
<td>Timestamp with time zone value converted from a string, with fractional seconds adjusted to the given precision</td>
<td><code>jsonb_path_query('"2023-08-15 12:34:56.789 +05:30"', '$.timestamp_tz(2)')</code> <code>"2023-08-15T12:34:56.79+05:30"</code></td>
</tr>
<tr>
<td><em>object</em> <code>.</code> <code>keyvalue()</code> <em>array</em></td>
<td>The object's key-value pairs, represented as an array of objects containing three fields: <code>"key"</code>, <code>"value"</code>, and <code>"id"</code>; <code>"id"</code> is a unique identifier of the object the key-value pair belongs to</td>
<td><code>jsonb_path_query_array('{"x": "20", "y": 32}', '$.keyvalue()')</code> <code>[{"id": 0, "key": "x", "value": "20"}, {"id": 0, "key": "y", "value": 32}]</code></td>
</tr>
</tbody>
</table>


!!! note

    The result type of the `datetime()` and <code>datetime(</code><em>template</em><code>)</code> methods can be `date`, `timetz`, `time`, `timestamptz`, or `timestamp`. Both methods determine their result type dynamically.


     The `datetime()` method sequentially tries to match its input string to the ISO formats for `date`, `timetz`, `time`, `timestamptz`, and `timestamp`. It stops on the first matching format and emits the corresponding data type.


     The <code>datetime(</code><em>template</em><code>)</code> method determines the result type according to the fields used in the provided template string.


     The `datetime()` and <code>datetime(</code><em>template</em><code>)</code> methods use the same parsing rules as the `to_timestamp` SQL function does (see [Data Type Formatting Functions](data-type-formatting-functions.md#functions-formatting)), with three exceptions. First, these methods don't allow unmatched template patterns. Second, only the following separators are allowed in the template string: minus sign, period, solidus (slash), comma, apostrophe, semicolon, colon and space. Third, separators in the template string must exactly match the input string.


     If different date/time types need to be compared, an implicit cast is applied. A `date` value can be cast to `timestamp` or `timestamptz`, `timestamp` can be cast to `timestamptz`, and `time` to `timetz`. However, all but the first of these conversions depend on the current [TimeZone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) setting, and thus can only be performed within timezone-aware `jsonpath` functions. Similarly, other date/time-related methods that convert strings to date/time types also do this casting, which may involve the current [TimeZone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) setting. Therefore, these conversions can also only be performed within timezone-aware `jsonpath` functions.


 [`jsonpath` Filter Expression Elements](#functions-sqljson-filter-ex-table) shows the available filter expression elements.


<a id="functions-sqljson-filter-ex-table"></a>
**Table: `jsonpath` Filter Expression Elements**

<table>
<thead>
<tr>
<th>Predicate/Value</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>value</em> <code>==</code> <em>value</em> <code>boolean</code></td>
<td>Equality comparison (this, and the other comparison operators, work on all JSON scalar values)</td>
<td><code>jsonb_path_query_array('[1, "a", 1, 3]', '$[*] ? (@ == 1)')</code> <code>[1, 1]</code><br><code>jsonb_path_query_array('[1, "a", 1, 3]', '$[*] ? (@ == "a")')</code> <code>["a"]</code></td>
</tr>
<tr>
<td><em>value</em> <code>!=</code> <em>value</em> <code>boolean</code></td>
<td><em>value</em> <code>&lt;&gt;</code> <em>value</em> <code>boolean</code></td>
<td>Non-equality comparison<br><code>jsonb_path_query_array('[1, 2, 1, 3]', '$[*] ? (@ != 1)')</code> <code>[2, 3]</code><br><code>jsonb_path_query_array('["a", "b", "c"]', '$[*] ? (@ &lt;&gt; "b")')</code> <code>["a", "c"]</code></td>
</tr>
<tr>
<td><em>value</em> <code>&lt;</code> <em>value</em> <code>boolean</code></td>
<td>Less-than comparison</td>
<td><code>jsonb_path_query_array('[1, 2, 3]', '$[*] ? (@ &lt; 2)')</code> <code>[1]</code></td>
</tr>
<tr>
<td><em>value</em> <code>&lt;=</code> <em>value</em> <code>boolean</code></td>
<td>Less-than-or-equal-to comparison</td>
<td><code>jsonb_path_query_array('["a", "b", "c"]', '$[*] ? (@ &lt;= "b")')</code> <code>["a", "b"]</code></td>
</tr>
<tr>
<td><em>value</em> <code>&gt;</code> <em>value</em> <code>boolean</code></td>
<td>Greater-than comparison</td>
<td><code>jsonb_path_query_array('[1, 2, 3]', '$[*] ? (@ &gt; 2)')</code> <code>[3]</code></td>
</tr>
<tr>
<td><em>value</em> <code>&gt;=</code> <em>value</em> <code>boolean</code></td>
<td>Greater-than-or-equal-to comparison</td>
<td><code>jsonb_path_query_array('[1, 2, 3]', '$[*] ? (@ &gt;= 2)')</code> <code>[2, 3]</code></td>
</tr>
<tr>
<td><code>true</code> <code>boolean</code></td>
<td>JSON constant <code>true</code></td>
<td><code>jsonb_path_query('[{"name": "John", "parent": false}, {"name": "Chris", "parent": true}]', '$[*] ? (@.parent == true)')</code> <code>{"name": "Chris", "parent": true}</code></td>
</tr>
<tr>
<td><code>false</code> <code>boolean</code></td>
<td>JSON constant <code>false</code></td>
<td><code>jsonb_path_query('[{"name": "John", "parent": false}, {"name": "Chris", "parent": true}]', '$[*] ? (@.parent == false)')</code> <code>{"name": "John", "parent": false}</code></td>
</tr>
<tr>
<td><code>null</code> <em>value</em></td>
<td>JSON constant <code>null</code> (note that, unlike in SQL, comparison to <code>null</code> works normally)</td>
<td><code>jsonb_path_query('[{"name": "Mary", "job": null}, {"name": "Michael", "job": "driver"}]', '$[*] ? (@.job == null) .name')</code> <code>"Mary"</code></td>
</tr>
<tr>
<td><em>boolean</em> <code>&amp;&amp;</code> <em>boolean</em> <code>boolean</code></td>
<td>Boolean AND</td>
<td><code>jsonb_path_query('[1, 3, 7]', '$[*] ? (@ &gt; 1 &amp;&amp; @ &lt; 5)')</code> <code>3</code></td>
</tr>
<tr>
<td><em>boolean</em> <code>||</code> <em>boolean</em> <code>boolean</code></td>
<td>Boolean OR</td>
<td><code>jsonb_path_query('[1, 3, 7]', '$[*] ? (@ &lt; 1 || @ &gt; 5)')</code> <code>7</code></td>
</tr>
<tr>
<td><code>!</code> <em>boolean</em> <code>boolean</code></td>
<td>Boolean NOT</td>
<td><code>jsonb_path_query('[1, 3, 7]', '$[*] ? (!(@ &lt; 5))')</code> <code>7</code></td>
</tr>
<tr>
<td><em>boolean</em> <code>is unknown</code> <code>boolean</code></td>
<td>Tests whether a Boolean condition is <code>unknown</code>.</td>
<td><code>jsonb_path_query('[-1, 2, 7, "foo"]', '$[*] ? ((@ &gt; 0) is unknown)')</code> <code>"foo"</code></td>
</tr>
<tr>
<td><em>string</em> <code>like_regex</code> <em>string</em> [ <code>flag</code> <em>string</em> ] <code>boolean</code></td>
<td>Tests whether the first operand matches the regular expression given by the second operand, optionally with modifications described by a string of <code>flag</code> characters (see <a href="#jsonpath-regular-expressions">SQL/JSON Regular Expressions</a>).</td>
<td><code>jsonb_path_query_array('["abc", "abd", "aBdC", "abdacb", "babc"]', '$[<em>] ? (@ like_regex "^ab.</em>c")')</code> <code>["abc", "abdacb"]</code><br><code>jsonb_path_query_array('["abc", "abd", "aBdC", "abdacb", "babc"]', '$[<em>] ? (@ like_regex "^ab.</em>c" flag "i")')</code> <code>["abc", "aBdC", "abdacb"]</code></td>
</tr>
<tr>
<td><em>string</em> <code>starts with</code> <em>string</em> <code>boolean</code></td>
<td>Tests whether the second operand is an initial substring of the first operand.</td>
<td><code>jsonb_path_query('["John Smith", "Mary Stone", "Bob Johnson"]', '$[*] ? (@ starts with "John")')</code> <code>"John Smith"</code></td>
</tr>
<tr>
<td><code>exists</code> <code>(</code> <em>path_expression</em> <code>)</code> <code>boolean</code></td>
<td>Tests whether a path expression matches at least one SQL/JSON item. Returns <code>unknown</code> if the path expression would result in an error; the second example uses this to avoid a no-such-key error in strict mode.</td>
<td><code>jsonb_path_query('{"x": [1, 2], "y": [2, 4]}', 'strict $.<em> ? (exists (@ ? (@[</em>] &gt; 2)))')</code> <code>[2, 4]</code><br><code>jsonb_path_query_array('{"value": 41}', 'strict $ ? (exists (@.name)) .name')</code> <code>[]</code></td>
</tr>
</tbody>
</table>


#### SQL/JSON Regular Expressions { #jsonpath-regular-expressions }


 SQL/JSON path expressions allow matching text to a regular expression with the `like_regex` filter. For example, the following SQL/JSON path query would case-insensitively match all strings in an array that start with an English vowel:

```

$[*] ? (@ like_regex "^[aeiou]" flag "i")
```


 The optional `flag` string may include one or more of the characters `i` for case-insensitive match, `m` to allow `^` and `$` to match at newlines, `s` to allow `.` to match a newline, and `q` to quote the whole pattern (reducing the behavior to a simple substring match).


 The SQL/JSON standard borrows its definition for regular expressions from the `LIKE_REGEX` operator, which in turn uses the XQuery standard. PostgreSQL does not currently support the `LIKE_REGEX` operator. Therefore, the `like_regex` filter is implemented using the POSIX regular expression engine described in [POSIX Regular Expressions](pattern-matching.md#functions-posix-regexp). This leads to various minor discrepancies from standard SQL/JSON behavior, which are cataloged in [Differences from SQL Standard and XQuery](pattern-matching.md#posix-vs-xquery). Note, however, that the flag-letter incompatibilities described there do not apply to SQL/JSON, as it translates the XQuery flag letters to match what the POSIX engine expects.


 Keep in mind that the pattern argument of `like_regex` is a JSON path string literal, written according to the rules given in [jsonpath Type](../data-types/json-types.md#datatype-jsonpath). This means in particular that any backslashes you want to use in the regular expression must be doubled. For example, to match string values of the root document that contain only digits:

```

$.* ? (@ like_regex "^\\d+$")
```


### SQL/JSON Query Functions { #sqljson-query-functions }


 SQL/JSON functions `JSON_EXISTS()`, `JSON_QUERY()`, and `JSON_VALUE()` described in [SQL/JSON Query Functions](#functions-sqljson-querying) can be used to query JSON documents. Each of these functions apply a *path_expression* (an SQL/JSON path query) to a *context_item* (the document). See [The SQL/JSON Path Language](#functions-sqljson-path) for more details on what the *path_expression* can contain. The *path_expression* can also reference variables, whose values are specified with their respective names in the `PASSING` clause that is supported by each function. *context_item* can be a `jsonb` value or a character string that can be successfully cast to `jsonb`.


<a id="functions-sqljson-querying"></a>
**Table: SQL/JSON Query Functions**

<table>
<thead>
<tr>
<th>Function signature</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>
JSON_EXISTS (
CONTEXT_ITEM, PATH_EXPRESSION
[ PASSING { VALUE AS VARNAME } [, ...]]
[{ TRUE | FALSE | UNKNOWN | ERROR } ON ERROR ]) boolean</code></pre></td>
<td>Examples:</td>
<td><code>JSON_EXISTS(jsonb '{"key1": [1,2,3]}', 'strict $.key1[*] ? (@ &gt; $x)' PASSING 2 AS x)</code> <code>t</code><br><code>JSON_EXISTS(jsonb '{"a": [1,2,3]}', 'lax $.a[5]' ERROR ON ERROR)</code> <code>f</code><br><p><code>JSON_EXISTS(jsonb '{"a": [1,2,3]}', 'strict $.a[5]' ERROR ON ERROR)</code></p>
<pre><code>
ERROR:  jsonpath array subscript is out of bounds</code></pre></td>
</tr>
<tr>
<td><pre><code>
JSON_QUERY (
CONTEXT_ITEM, PATH_EXPRESSION
[ PASSING { VALUE AS VARNAME } [, ...]]
[ RETURNING DATA_TYPE [ FORMAT JSON [ ENCODING UTF8 ] ] ]
[ { WITHOUT | WITH { CONDITIONAL | [UNCONDITIONAL] } } [ ARRAY ] WRAPPER ]
[ { KEEP | OMIT } QUOTES [ ON SCALAR STRING ] ]
[ { ERROR | NULL | EMPTY { [ ARRAY ] | OBJECT } | DEFAULT EXPRESSION } ON EMPTY ]
[ { ERROR | NULL | EMPTY { [ ARRAY ] | OBJECT } | DEFAULT EXPRESSION } ON ERROR ]) jsonb</code></pre></td>
<td>Examples:</td>
<td><code>JSON_QUERY(jsonb '[1,[2,3],null]', 'lax $[*][$off]' PASSING 1 AS off WITH CONDITIONAL WRAPPER)</code> <code>3</code><br><code>JSON_QUERY(jsonb '{"a": "[1, 2]"}', 'lax $.a' OMIT QUOTES)</code> <code>[1, 2]</code><br><p><code>JSON_QUERY(jsonb '{"a": "[1, 2]"}', 'lax $.a' RETURNING int[] OMIT QUOTES ERROR ON ERROR)</code></p>
<pre><code>
ERROR:  malformed array literal: "[1, 2]"
DETAIL:  Missing "]" after array dimensions.</code></pre></td>
</tr>
<tr>
<td><pre><code>
JSON_VALUE (
CONTEXT_ITEM, PATH_EXPRESSION
[ PASSING { VALUE AS VARNAME } [, ...]]
[ RETURNING DATA_TYPE ]
[ { ERROR | NULL | DEFAULT EXPRESSION } ON EMPTY ]
[ { ERROR | NULL | DEFAULT EXPRESSION } ON ERROR ]) text</code></pre></td>
<td>Examples:</td>
<td><code>JSON_VALUE(jsonb '"123.45"', '$' RETURNING float)</code> <code>123.45</code><br><code>JSON_VALUE(jsonb '"03:04 2015-02-01"', '$.datetime("HH24:MI YYYY-MM-DD")' RETURNING date)</code> <code>2015-02-01</code><br><code>JSON_VALUE(jsonb '[1,2]', 'strict $[$off]' PASSING 1 AS off)</code> <code>2</code><br><code>JSON_VALUE(jsonb '[1,2]', 'strict $[*]' DEFAULT 9 ON ERROR)</code> <code>9</code></td>
</tr>
</tbody>
</table>


!!! note

    The *context_item* expression is converted to `jsonb` by an implicit cast if the expression is not already of type `jsonb`. Note, however, that any parsing errors that occur during that conversion are thrown unconditionally, that is, are not handled according to the (specified or implicit) `ON ERROR` clause.


!!! note

    `JSON_VALUE()` returns an SQL NULL if *path_expression* returns a JSON `null`, whereas `JSON_QUERY()` returns the JSON `null` as is.


### JSON_TABLE { #functions-sqljson-table }


 `JSON_TABLE` is an SQL/JSON function which queries JSON data and presents the results as a relational view, which can be accessed as a regular SQL table. You can use `JSON_TABLE` inside the `FROM` clause of a `SELECT`, `UPDATE`, or `DELETE` and as data source in a `MERGE` statement.


 Taking JSON data as input, `JSON_TABLE` uses a JSON path expression to extract a part of the provided data to use as a *row pattern* for the constructed view. Each SQL/JSON value given by the row pattern serves as source for a separate row in the constructed view.


 To split the row pattern into columns, `JSON_TABLE` provides the `COLUMNS` clause that defines the schema of the created view. For each column, a separate JSON path expression can be specified to be evaluated against the row pattern to get an SQL/JSON value that will become the value for the specified column in a given output row.


 JSON data stored at a nested level of the row pattern can be extracted using the `NESTED PATH` clause. Each `NESTED PATH` clause can be used to generate one or more columns using the data from a nested level of the row pattern. Those columns can be specified using a `COLUMNS` clause that looks similar to the top-level COLUMNS clause. Rows constructed from NESTED COLUMNS are called *child rows* and are joined against the row constructed from the columns specified in the parent `COLUMNS` clause to get the row in the final view. Child columns themselves may contain a `NESTED PATH` specification thus allowing to extract data located at arbitrary nesting levels. Columns produced by multiple `NESTED PATH`s at the same level are considered to be *siblings* of each other and their rows after joining with the parent row are combined using UNION.


 The rows produced by `JSON_TABLE` are laterally joined to the row that generated them, so you do not have to explicitly join the constructed view with the original table holding JSON data.


 The syntax is:


```

JSON_TABLE (
    CONTEXT_ITEM, PATH_EXPRESSION [ AS JSON_PATH_NAME ] [ PASSING { VALUE AS VARNAME } [, ...] ]
    COLUMNS ( JSON_TABLE_COLUMN [, ...] )
    [ { ERROR | EMPTY [ARRAY]} ON ERROR ]
)


where JSON_TABLE_COLUMN is:

  NAME FOR ORDINALITY
  | NAME TYPE
        [ FORMAT JSON [ENCODING UTF8]]
        [ PATH PATH_EXPRESSION ]
        [ { WITHOUT | WITH { CONDITIONAL | [UNCONDITIONAL] } } [ ARRAY ] WRAPPER ]
        [ { KEEP | OMIT } QUOTES [ ON SCALAR STRING ] ]
        [ { ERROR | NULL | EMPTY { [ARRAY] | OBJECT } | DEFAULT EXPRESSION } ON EMPTY ]
        [ { ERROR | NULL | EMPTY { [ARRAY] | OBJECT } | DEFAULT EXPRESSION } ON ERROR ]
  | NAME TYPE EXISTS [ PATH PATH_EXPRESSION ]
        [ { ERROR | TRUE | FALSE | UNKNOWN } ON ERROR ]
  | NESTED [ PATH ] PATH_EXPRESSION [ AS JSON_PATH_NAME ] COLUMNS ( JSON_TABLE_COLUMN [, ...] )
```


 Each syntax element is described below in more detail.


<em>context_item</em><code>, </code><em>path_expression</em><code> [ `AS` *json_path_name* ] [ `PASSING` { *value* `AS` *varname* } [, ...]]</code>
:   The *context_item* specifies the input document to query, the *path_expression* is an SQL/JSON path expression defining the query, and *json_path_name* is an optional name for the *path_expression*. The optional `PASSING` clause provides data values for the variables mentioned in the *path_expression*. The result of the input data evaluation using the aforementioned elements is called the *row pattern*, which is used as the source for row values in the constructed view.

`COLUMNS` ( *json_table_column* [, ...] )
:   The `COLUMNS` clause defining the schema of the constructed view. In this clause, you can specify each column to be filled with an SQL/JSON value obtained by applying a JSON path expression against the row pattern. *json_table_column* has the following variants:


    *name* `FOR ORDINALITY`
    :   Adds an ordinality column that provides sequential row numbering starting from 1. Each `NESTED PATH` (see below) gets its own counter for any nested ordinality columns.

    <em>name</em><code> </code><em>type</em><code>
              [`FORMAT JSON` [ENCODING `UTF8`]]
              [ `PATH` *path_expression* ]</code>
    :   Inserts an SQL/JSON value obtained by applying *path_expression* against the row pattern into the view's output row after coercing it to specified *type*.


         Specifying `FORMAT JSON` makes it explicit that you expect the value to be a valid `json` object. It only makes sense to specify `FORMAT JSON` if *type* is one of `bpchar`, `bytea`, `character varying`, `name`, `json`, `jsonb`, `text`, or a domain over these types.


         Optionally, you can specify `WRAPPER` and `QUOTES` clauses to format the output. Note that specifying `OMIT QUOTES` overrides `FORMAT JSON` if also specified, because unquoted literals do not constitute valid `json` values.


         Optionally, you can use `ON EMPTY` and `ON ERROR` clauses to specify whether to throw the error or return the specified value when the result of JSON path evaluation is empty and when an error occurs during JSON path evaluation or when coercing the SQL/JSON value to the specified type, respectively. The default for both is to return a `NULL` value.


        !!! note

            This clause is internally turned into and has the same semantics as `JSON_VALUE` or `JSON_QUERY`. The latter if the specified type is not a scalar type or if either of `FORMAT JSON`, `WRAPPER`, or `QUOTES` clause is present.

    *name* *type* `EXISTS` [ `PATH` *path_expression* ]
    :   Inserts a boolean value obtained by applying *path_expression* against the row pattern into the view's output row after coercing it to specified *type*.


         The value corresponds to whether applying the `PATH` expression to the row pattern yields any values.


         The specified *type* should have a cast from the `boolean` type.


         Optionally, you can use `ON ERROR` to specify whether to throw the error or return the specified value when an error occurs during JSON path evaluation or when coercing SQL/JSON value to the specified type. The default is to return a boolean value `FALSE`.


        !!! note

            This clause is internally turned into and has the same semantics as `JSON_EXISTS`.

    `NESTED [ PATH ]` *path_expression* [ `AS` *json_path_name* ] `COLUMNS` ( *json_table_column* [, ...] )
    :   Extracts SQL/JSON values from nested levels of the row pattern, generates one or more columns as defined by the `COLUMNS` subclause, and inserts the extracted SQL/JSON values into those columns. The *json_table_column* expression in the `COLUMNS` subclause uses the same syntax as in the parent `COLUMNS` clause.


         The `NESTED PATH` syntax is recursive, so you can go down multiple nested levels by specifying several `NESTED PATH` subclauses within each other. It allows to unnest the hierarchy of JSON objects and arrays in a single function invocation rather than chaining several `JSON_TABLE` expressions in an SQL statement.


    !!! note

        In each variant of *json_table_column* described above, if the `PATH` clause is omitted, path expression <code>$.</code><em>name</em> is used, where *name* is the provided column name.

`AS` *json_path_name*
:   The optional *json_path_name* serves as an identifier of the provided *path_expression*. The name must be unique and distinct from the column names.

{ `ERROR` | `EMPTY` } `ON ERROR`
:   The optional `ON ERROR` can be used to specify how to handle errors when evaluating the top-level *path_expression*. Use `ERROR` if you want the errors to be thrown and `EMPTY` to return an empty table, that is, a table containing 0 rows. Note that this clause does not affect the errors that occur when evaluating columns, for which the behavior depends on whether the `ON ERROR` clause is specified against a given column.


Examples


 In the examples that follow, the following table containing JSON data will be used:

```sql

CREATE TABLE my_films ( js jsonb );

INSERT INTO my_films VALUES (
'{ "favorites" : [
   { "kind" : "comedy", "films" : [
     { "title" : "Bananas",
       "director" : "Woody Allen"},
     { "title" : "The Dinner Game",
       "director" : "Francis Veber" } ] },
   { "kind" : "horror", "films" : [
     { "title" : "Psycho",
       "director" : "Alfred Hitchcock" } ] },
   { "kind" : "thriller", "films" : [
     { "title" : "Vertigo",
       "director" : "Alfred Hitchcock" } ] },
   { "kind" : "drama", "films" : [
     { "title" : "Yojimbo",
       "director" : "Akira Kurosawa" } ] }
  ] }');
```


 The following query shows how to use `JSON_TABLE` to turn the JSON objects in the `my_films` table to a view containing columns for the keys `kind`, `title`, and `director` contained in the original JSON along with an ordinality column:

```sql

SELECT jt.* FROM
 my_films,
 JSON_TABLE (js, '$.favorites[*]' COLUMNS (
   id FOR ORDINALITY,
   kind text PATH '$.kind',
   title text PATH '$.films[*].title' WITH WRAPPER,
   director text PATH '$.films[*].director' WITH WRAPPER)) AS jt;
```


```

 id |   kind   |             title              |             director
----+----------+--------------------------------+----------------------------------
  1 | comedy   | ["Bananas", "The Dinner Game"] | ["Woody Allen", "Francis Veber"]
  2 | horror   | ["Psycho"]                     | ["Alfred Hitchcock"]
  3 | thriller | ["Vertigo"]                    | ["Alfred Hitchcock"]
  4 | drama    | ["Yojimbo"]                    | ["Akira Kurosawa"]
(4 rows)
```


 The following is a modified version of the above query to show the usage of `PASSING` arguments in the filter specified in the top-level JSON path expression and the various options for the individual columns:

```sql

SELECT jt.* FROM
 my_films,
 JSON_TABLE (js, '$.favorites[*] ? (@.films[*].director == $filter)'
   PASSING 'Alfred Hitchcock' AS filter
     COLUMNS (
     id FOR ORDINALITY,
     kind text PATH '$.kind',
     title text FORMAT JSON PATH '$.films[*].title' OMIT QUOTES,
     director text PATH '$.films[*].director' KEEP QUOTES)) AS jt;
```


```

 id |   kind   |  title  |      director
----+----------+---------+--------------------
  1 | horror   | Psycho  | "Alfred Hitchcock"
  2 | thriller | Vertigo | "Alfred Hitchcock"
(2 rows)
```


 The following is a modified version of the above query to show the usage of `NESTED PATH` for populating title and director columns, illustrating how they are joined to the parent columns id and kind:

```sql

SELECT jt.* FROM
 my_films,
 JSON_TABLE ( js, '$.favorites[*] ? (@.films[*].director == $filter)'
   PASSING 'Alfred Hitchcock' AS filter
   COLUMNS (
    id FOR ORDINALITY,
    kind text PATH '$.kind',
    NESTED PATH '$.films[*]' COLUMNS (
      title text FORMAT JSON PATH '$.title' OMIT QUOTES,
      director text PATH '$.director' KEEP QUOTES))) AS jt;
```


```

 id |   kind   |  title  |      director
----+----------+---------+--------------------
  1 | horror   | Psycho  | "Alfred Hitchcock"
  2 | thriller | Vertigo | "Alfred Hitchcock"
(2 rows)
```


 The following is the same query but without the filter in the root path:

```sql

SELECT jt.* FROM
 my_films,
 JSON_TABLE ( js, '$.favorites[*]'
   COLUMNS (
    id FOR ORDINALITY,
    kind text PATH '$.kind',
    NESTED PATH '$.films[*]' COLUMNS (
      title text FORMAT JSON PATH '$.title' OMIT QUOTES,
      director text PATH '$.director' KEEP QUOTES))) AS jt;
```


```

 id |   kind   |      title      |      director
----+----------+-----------------+--------------------
  1 | comedy   | Bananas         | "Woody Allen"
  1 | comedy   | The Dinner Game | "Francis Veber"
  2 | horror   | Psycho          | "Alfred Hitchcock"
  3 | thriller | Vertigo         | "Alfred Hitchcock"
  4 | drama    | Yojimbo         | "Akira Kurosawa"
(5 rows)
```


 The following shows another query using a different `JSON` object as input. It shows the UNION "sibling join" between `NESTED` paths `$.movies[*]` and `$.books[*]` and also the usage of `FOR ORDINALITY` column at `NESTED` levels (columns `movie_id`, `book_id`, and `author_id`):

```sql

SELECT * FROM JSON_TABLE (
'{"favorites":
    [{"movies":
      [{"name": "One", "director": "John Doe"},
       {"name": "Two", "director": "Don Joe"}],
     "books":
      [{"name": "Mystery", "authors": [{"name": "Brown Dan"}]},
       {"name": "Wonder", "authors": [{"name": "Jun Murakami"}, {"name":"Craig Doe"}]}]
}]}'::json, '$.favorites[*]'
COLUMNS (
  user_id FOR ORDINALITY,
  NESTED '$.movies[*]'
    COLUMNS (
    movie_id FOR ORDINALITY,
    mname text PATH '$.name',
    director text),
  NESTED '$.books[*]'
    COLUMNS (
      book_id FOR ORDINALITY,
      bname text PATH '$.name',
      NESTED '$.authors[*]'
        COLUMNS (
          author_id FOR ORDINALITY,
          author_name text PATH '$.name'))));
```


```

 user_id | movie_id | mname | director | book_id |  bname  | author_id | author_name
---------+----------+-------+----------+---------+---------+-----------+--------------
       1 |        1 | One   | John Doe |         |         |           |
       1 |        2 | Two   | Don Joe  |         |         |           |
       1 |          |       |          |       1 | Mystery |         1 | Brown Dan
       1 |          |       |          |       2 | Wonder  |         1 | Jun Murakami
       1 |          |       |          |       2 | Wonder  |         2 | Craig Doe
(5 rows)
```
