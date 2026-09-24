<a id="functions-textsearch"></a>

## Text Search Functions and Operators


 [Text Search Operators](#textsearch-operators-table), [Text Search Functions](#textsearch-functions-table) and [Text Search Debugging Functions](#textsearch-functions-debug-table) summarize the functions and operators that are provided for full text searching. See [Full Text Search](../full-text-search/index.md#textsearch) for a detailed explanation of PostgreSQL's text search facility.
 <a id="textsearch-operators-table"></a>

**Table: Text Search Operators**

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
<td><code>tsvector</code> <code>@@</code> <code>tsquery</code> <code>boolean</code></td>
<td><code>tsquery</code> <code>@@</code> <code>tsvector</code> <code>boolean</code></td>
<td>Does <code>tsvector</code> match <code>tsquery</code>? (The arguments can be given in either order.)<br><code>to_tsvector('fat cats ate rats') @@ to_tsquery('cat &amp; rat')</code> <code>t</code></td>
</tr>
<tr>
<td><code>text</code> <code>@@</code> <code>tsquery</code> <code>boolean</code></td>
<td>Does text string, after implicit invocation of <code>to_tsvector()</code>, match <code>tsquery</code>?</td>
<td><code>'fat cats ate rats' @@ to_tsquery('cat &amp; rat')</code> <code>t</code></td>
</tr>
<tr>
<td><code>tsvector</code> <code>||</code> <code>tsvector</code> <code>tsvector</code></td>
<td>Concatenates two <code>tsvector</code>s. If both inputs contain lexeme positions, the second input's positions are adjusted accordingly.</td>
<td><code>'a:1 b:2'::tsvector || 'c:1 d:2 b:3'::tsvector</code> <code>'a':1 'b':2,5 'c':3 'd':4</code></td>
</tr>
<tr>
<td><code>tsquery</code> <code>&amp;&amp;</code> <code>tsquery</code> <code>tsquery</code></td>
<td>ANDs two <code>tsquery</code>s together, producing a query that matches documents that match both input queries.</td>
<td><code>'fat | rat'::tsquery &amp;&amp; 'cat'::tsquery</code> <code>( 'fat' | 'rat' ) &amp; 'cat'</code></td>
</tr>
<tr>
<td><code>tsquery</code> <code>||</code> <code>tsquery</code> <code>tsquery</code></td>
<td>ORs two <code>tsquery</code>s together, producing a query that matches documents that match either input query.</td>
<td><code>'fat | rat'::tsquery || 'cat'::tsquery</code> <code>'fat' | 'rat' | 'cat'</code></td>
</tr>
<tr>
<td><code>!!</code> <code>tsquery</code> <code>tsquery</code></td>
<td>Negates a <code>tsquery</code>, producing a query that matches documents that do not match the input query.</td>
<td><code>!! 'cat'::tsquery</code> <code>!'cat'</code></td>
</tr>
<tr>
<td><code>tsquery</code> <code>&lt;-&gt;</code> <code>tsquery</code> <code>tsquery</code></td>
<td>Constructs a phrase query, which matches if the two input queries match at successive lexemes.</td>
<td><code>to_tsquery('fat') &lt;-&gt; to_tsquery('rat')</code> <code>'fat' &lt;-&gt; 'rat'</code></td>
</tr>
<tr>
<td><code>tsquery</code> <code>@&gt;</code> <code>tsquery</code> <code>boolean</code></td>
<td>Does first <code>tsquery</code> contain the second? (This considers only whether all the lexemes appearing in one query appear in the other, ignoring the combining operators.)</td>
<td><code>'cat'::tsquery @&gt; 'cat &amp; rat'::tsquery</code> <code>f</code></td>
</tr>
<tr>
<td><code>tsquery</code> <code>&lt;@</code> <code>tsquery</code> <code>boolean</code></td>
<td>Is first <code>tsquery</code> contained in the second? (This considers only whether all the lexemes appearing in one query appear in the other, ignoring the combining operators.)</td>
<td><code>'cat'::tsquery &lt;@ 'cat &amp; rat'::tsquery</code> <code>t</code><br><code>'cat'::tsquery &lt;@ '!cat &amp; rat'::tsquery</code> <code>t</code></td>
</tr>
</tbody>
</table>


 In addition to these specialized operators, the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) are available for types `tsvector` and `tsquery`. These are not very useful for text searching but allow, for example, unique indexes to be built on columns of these types.
 <a id="textsearch-functions-table"></a>

**Table: Text Search Functions**

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
<td><code>array_to_tsvector</code> ( <code>text[]</code> ) <code>tsvector</code></td>
<td>Converts an array of text strings to a <code>tsvector</code>. The given strings are used as lexemes as-is, without further processing. Array elements must not be empty strings or <code>NULL</code>.</td>
<td><code>array_to_tsvector('{fat,cat,rat}'::text[])</code> <code>'cat' 'fat' 'rat'</code></td>
</tr>
<tr>
<td><code>get_current_ts_config</code> ( ) <code>regconfig</code></td>
<td>Returns the OID of the current default text search configuration (as set by <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-default-text-search-config">default_text_search_config</a>).</td>
<td><code>get_current_ts_config()</code> <code>english</code></td>
</tr>
<tr>
<td><code>length</code> ( <code>tsvector</code> ) <code>integer</code></td>
<td>Returns the number of lexemes in the <code>tsvector</code>.</td>
<td><code>length('fat:2,4 cat:3 rat:5A'::tsvector)</code> <code>3</code></td>
</tr>
<tr>
<td><code>numnode</code> ( <code>tsquery</code> ) <code>integer</code></td>
<td>Returns the number of lexemes plus operators in the <code>tsquery</code>.</td>
<td><code>numnode('(fat &amp; rat) | cat'::tsquery)</code> <code>5</code></td>
</tr>
<tr>
<td><code>plainto_tsquery</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>query</code> <code>text</code> ) <code>tsquery</code></td>
<td>Converts text to a <code>tsquery</code>, normalizing words according to the specified or default configuration. Any punctuation in the string is ignored (it does not determine query operators). The resulting query matches documents containing all non-stopwords in the text.</td>
<td><code>plainto_tsquery('english', 'The Fat Rats')</code> <code>'fat' &amp; 'rat'</code></td>
</tr>
<tr>
<td><code>phraseto_tsquery</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>query</code> <code>text</code> ) <code>tsquery</code></td>
<td>Converts text to a <code>tsquery</code>, normalizing words according to the specified or default configuration. Any punctuation in the string is ignored (it does not determine query operators). The resulting query matches phrases containing all non-stopwords in the text.</td>
<td><code>phraseto_tsquery('english', 'The Fat Rats')</code> <code>'fat' &lt;-&gt; 'rat'</code><br><code>phraseto_tsquery('english', 'The Cat and Rats')</code> <code>'cat' &lt;2&gt; 'rat'</code></td>
</tr>
<tr>
<td><code>websearch_to_tsquery</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>query</code> <code>text</code> ) <code>tsquery</code></td>
<td>Converts text to a <code>tsquery</code>, normalizing words according to the specified or default configuration. Quoted word sequences are converted to phrase tests. The word “or” is understood as producing an OR operator, and a dash produces a NOT operator; other punctuation is ignored. This approximates the behavior of some common web search tools.</td>
<td><code>websearch_to_tsquery('english', '"fat rat" or cat dog')</code> <code>'fat' &lt;-&gt; 'rat' | 'cat' &amp; 'dog'</code></td>
</tr>
<tr>
<td><code>querytree</code> ( <code>tsquery</code> ) <code>text</code></td>
<td>Produces a representation of the indexable portion of a <code>tsquery</code>. A result that is empty or just <code>T</code> indicates a non-indexable query.</td>
<td><code>querytree('foo &amp; ! bar'::tsquery)</code> <code>'foo'</code></td>
</tr>
<tr>
<td><code>setweight</code> ( <code>vector</code> <code>tsvector</code>, <code>weight</code> <code>"char"</code> ) <code>tsvector</code></td>
<td>Assigns the specified <code>weight</code> to each element of the <code>vector</code>.</td>
<td><code>setweight('fat:2,4 cat:3 rat:5B'::tsvector, 'A')</code> <code>'cat':3A 'fat':2A,4A 'rat':5A</code></td>
</tr>
<tr>
<td><code>setweight</code> ( <code>vector</code> <code>tsvector</code>, <code>weight</code> <code>"char"</code>, <code>lexemes</code> <code>text[]</code> ) <code>tsvector</code></td>
<td>Assigns the specified <code>weight</code> to elements of the <code>vector</code> that are listed in <code>lexemes</code>. The strings in <code>lexemes</code> are taken as lexemes as-is, without further processing. Strings that do not match any lexeme in <code>vector</code> are ignored.</td>
<td><code>setweight('fat:2,4 cat:3 rat:5,6B'::tsvector, 'A', '{cat,rat}')</code> <code>'cat':3A 'fat':2,4 'rat':5A,6A</code></td>
</tr>
<tr>
<td><code>strip</code> ( <code>tsvector</code> ) <code>tsvector</code></td>
<td>Removes positions and weights from the <code>tsvector</code>.</td>
<td><code>strip('fat:2,4 cat:3 rat:5A'::tsvector)</code> <code>'cat' 'fat' 'rat'</code></td>
</tr>
<tr>
<td><code>to_tsquery</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>query</code> <code>text</code> ) <code>tsquery</code></td>
<td>Converts text to a <code>tsquery</code>, normalizing words according to the specified or default configuration. The words must be combined by valid <code>tsquery</code> operators.</td>
<td><code>to_tsquery('english', 'The &amp; Fat &amp; Rats')</code> <code>'fat' &amp; 'rat'</code></td>
</tr>
<tr>
<td><code>to_tsvector</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>text</code> ) <code>tsvector</code></td>
<td>Converts text to a <code>tsvector</code>, normalizing words according to the specified or default configuration. Position information is included in the result.</td>
<td><code>to_tsvector('english', 'The Fat Rats')</code> <code>'fat':2 'rat':3</code></td>
</tr>
<tr>
<td><code>to_tsvector</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>json</code> ) <code>tsvector</code></td>
<td><code>to_tsvector</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>jsonb</code> ) <code>tsvector</code></td>
<td>Converts each string value in the JSON document to a <code>tsvector</code>, normalizing words according to the specified or default configuration. The results are then concatenated in document order to produce the output. Position information is generated as though one stopword exists between each pair of string values. (Beware that “document order” of the fields of a JSON object is implementation-dependent when the input is <code>jsonb</code>; observe the difference in the examples.)<br><code>to_tsvector('english', '{"aa": "The Fat Rats", "b": "dog"}'::json)</code> <code>'dog':5 'fat':2 'rat':3</code><br><code>to_tsvector('english', '{"aa": "The Fat Rats", "b": "dog"}'::jsonb)</code> <code>'dog':1 'fat':4 'rat':5</code></td>
</tr>
<tr>
<td><code>json_to_tsvector</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>json</code>, <code>filter</code> <code>jsonb</code> ) <code>tsvector</code></td>
<td><code>jsonb_to_tsvector</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>jsonb</code>, <code>filter</code> <code>jsonb</code> ) <code>tsvector</code></td>
<td>Selects each item in the JSON document that is requested by the <code>filter</code> and converts each one to a <code>tsvector</code>, normalizing words according to the specified or default configuration. The results are then concatenated in document order to produce the output. Position information is generated as though one stopword exists between each pair of selected items. (Beware that “document order” of the fields of a JSON object is implementation-dependent when the input is <code>jsonb</code>.) The <code>filter</code> must be a <code>jsonb</code> array containing zero or more of these keywords: <code>"string"</code> (to include all string values), <code>"numeric"</code> (to include all numeric values), <code>"boolean"</code> (to include all boolean values), <code>"key"</code> (to include all keys), or <code>"all"</code> (to include all the above). As a special case, the <code>filter</code> can also be a simple JSON value that is one of these keywords.<br><code>json_to_tsvector('english', '{"a": "The Fat Rats", "b": 123}'::json, '["string", "numeric"]')</code> <code>'123':5 'fat':2 'rat':3</code><br><code>json_to_tsvector('english', '{"cat": "The Fat Rats", "dog": 123}'::json, '"all"')</code> <code>'123':9 'cat':1 'dog':7 'fat':4 'rat':5</code></td>
</tr>
<tr>
<td><code>ts_delete</code> ( <code>vector</code> <code>tsvector</code>, <code>lexeme</code> <code>text</code> ) <code>tsvector</code></td>
<td>Removes any occurrence of the given <code>lexeme</code> from the <code>vector</code>. The <code>lexeme</code> string is treated as a lexeme as-is, without further processing.</td>
<td><code>ts_delete('fat:2,4 cat:3 rat:5A'::tsvector, 'fat')</code> <code>'cat':3 'rat':5A</code></td>
</tr>
<tr>
<td><code>ts_delete</code> ( <code>vector</code> <code>tsvector</code>, <code>lexemes</code> <code>text[]</code> ) <code>tsvector</code></td>
<td>Removes any occurrences of the lexemes in <code>lexemes</code> from the <code>vector</code>. The strings in <code>lexemes</code> are taken as lexemes as-is, without further processing. Strings that do not match any lexeme in <code>vector</code> are ignored.</td>
<td><code>ts_delete('fat:2,4 cat:3 rat:5A'::tsvector, ARRAY['fat','rat'])</code> <code>'cat':3</code></td>
</tr>
<tr>
<td><code>ts_filter</code> ( <code>vector</code> <code>tsvector</code>, <code>weights</code> <code>"char"[]</code> ) <code>tsvector</code></td>
<td>Selects only elements with the given <code>weights</code> from the <code>vector</code>.</td>
<td><code>ts_filter('fat:2,4 cat:3b,7c rat:5A'::tsvector, '{a,b}')</code> <code>'cat':3B 'rat':5A</code></td>
</tr>
<tr>
<td><code>ts_headline</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>text</code>, <code>query</code> <code>tsquery</code> [, <code>options</code> <code>text</code> ] ) <code>text</code></td>
<td>Displays, in an abbreviated form, the match(es) for the <code>query</code> in the <code>document</code>, which must be raw text not a <code>tsvector</code>. Words in the document are normalized according to the specified or default configuration before matching to the query. Use of this function is discussed in <a href="../full-text-search/controlling-text-search.md#textsearch-headline">Highlighting Results</a>, which also describes the available <code>options</code>.</td>
<td><code>ts_headline('The fat cat ate the rat.', 'cat')</code> <code>The fat &lt;b&gt;cat&lt;/b&gt; ate the rat.</code></td>
</tr>
<tr>
<td><code>ts_headline</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>json</code>, <code>query</code> <code>tsquery</code> [, <code>options</code> <code>text</code> ] ) <code>text</code></td>
<td><code>ts_headline</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>jsonb</code>, <code>query</code> <code>tsquery</code> [, <code>options</code> <code>text</code> ] ) <code>text</code></td>
<td>Displays, in an abbreviated form, match(es) for the <code>query</code> that occur in string values within the JSON <code>document</code>. See <a href="../full-text-search/controlling-text-search.md#textsearch-headline">Highlighting Results</a> for more details.<br><code>ts_headline('{"cat":"raining cats and dogs"}'::jsonb, 'cat')</code> <code>{"cat": "raining &lt;b&gt;cats&lt;/b&gt; and dogs"}</code></td>
</tr>
<tr>
<td><code>ts_rank</code> ( [ <code>weights</code> <code>real[]</code>, ] <code>vector</code> <code>tsvector</code>, <code>query</code> <code>tsquery</code> [, <code>normalization</code> <code>integer</code> ] ) <code>real</code></td>
<td>Computes a score showing how well the <code>vector</code> matches the <code>query</code>. See <a href="../full-text-search/controlling-text-search.md#textsearch-ranking">Ranking Search Results</a> for details.</td>
<td><code>ts_rank(to_tsvector('raining cats and dogs'), 'cat')</code> <code>0.06079271</code></td>
</tr>
<tr>
<td><code>ts_rank_cd</code> ( [ <code>weights</code> <code>real[]</code>, ] <code>vector</code> <code>tsvector</code>, <code>query</code> <code>tsquery</code> [, <code>normalization</code> <code>integer</code> ] ) <code>real</code></td>
<td>Computes a score showing how well the <code>vector</code> matches the <code>query</code>, using a cover density algorithm. See <a href="../full-text-search/controlling-text-search.md#textsearch-ranking">Ranking Search Results</a> for details.</td>
<td><code>ts_rank_cd(to_tsvector('raining cats and dogs'), 'cat')</code> <code>0.1</code></td>
</tr>
<tr>
<td><code>ts_rewrite</code> ( <code>query</code> <code>tsquery</code>, <code>target</code> <code>tsquery</code>, <code>substitute</code> <code>tsquery</code> ) <code>tsquery</code></td>
<td>Replaces occurrences of <code>target</code> with <code>substitute</code> within the <code>query</code>. See <a href="../full-text-search/additional-features.md#textsearch-query-rewriting">Query Rewriting</a> for details.</td>
<td><code>ts_rewrite('a &amp; b'::tsquery, 'a'::tsquery, 'foo|bar'::tsquery)</code> <code>'b' &amp; ( 'foo' | 'bar' )</code></td>
</tr>
<tr>
<td><code>ts_rewrite</code> ( <code>query</code> <code>tsquery</code>, <code>select</code> <code>text</code> ) <code>tsquery</code></td>
<td>Replaces portions of the <code>query</code> according to target(s) and substitute(s) obtained by executing a <code>SELECT</code> command. See <a href="../full-text-search/additional-features.md#textsearch-query-rewriting">Query Rewriting</a> for details.</td>
<td><code>SELECT ts_rewrite('a &amp; b'::tsquery, 'SELECT t,s FROM aliases')</code> <code>'b' &amp; ( 'foo' | 'bar' )</code></td>
</tr>
<tr>
<td><code>tsquery_phrase</code> ( <code>query1</code> <code>tsquery</code>, <code>query2</code> <code>tsquery</code> ) <code>tsquery</code></td>
<td>Constructs a phrase query that searches for matches of <code>query1</code> and <code>query2</code> at successive lexemes (same as <code>&lt;-&gt;</code> operator).</td>
<td><code>tsquery_phrase(to_tsquery('fat'), to_tsquery('cat'))</code> <code>'fat' &lt;-&gt; 'cat'</code></td>
</tr>
<tr>
<td><code>tsquery_phrase</code> ( <code>query1</code> <code>tsquery</code>, <code>query2</code> <code>tsquery</code>, <code>distance</code> <code>integer</code> ) <code>tsquery</code></td>
<td>Constructs a phrase query that searches for matches of <code>query1</code> and <code>query2</code> that occur exactly <code>distance</code> lexemes apart.</td>
<td><code>tsquery_phrase(to_tsquery('fat'), to_tsquery('cat'), 10)</code> <code>'fat' &lt;10&gt; 'cat'</code></td>
</tr>
<tr>
<td><code>tsvector_to_array</code> ( <code>tsvector</code> ) <code>text[]</code></td>
<td>Converts a <code>tsvector</code> to an array of lexemes.</td>
<td><code>tsvector_to_array('fat:2,4 cat:3 rat:5A'::tsvector)</code> <code>{cat,fat,rat}</code></td>
</tr>
<tr>
<td><code>unnest</code> ( <code>tsvector</code> ) <code>setof record</code> ( <code>lexeme</code> <code>text</code>, <code>positions</code> <code>smallint[]</code>, <code>weights</code> <code>text</code> )</td>
<td>Expands a <code>tsvector</code> into a set of rows, one per lexeme.</td>
<td><p><code>SELECT * FROM unnest('cat:3 fat:2,4 rat:5A'::tsvector)</code></p>
<pre><code>
 lexeme | positions | weights
--------+-----------+---------
 cat    | {3}       | {D}
 fat    | {2,4}     | {D,D}
 rat    | {5}       | {A}</code></pre></td>
</tr>
</tbody>
</table>


!!! note

    All the text search functions that accept an optional `regconfig` argument will use the configuration specified by [default_text_search_config](../../server-administration/server-configuration/client-connection-defaults.md#guc-default-text-search-config) when that argument is omitted.


 The functions in [Text Search Debugging Functions](#textsearch-functions-debug-table) are listed separately because they are not usually used in everyday text searching operations. They are primarily helpful for development and debugging of new text search configurations.
 <a id="textsearch-functions-debug-table"></a>

**Table: Text Search Debugging Functions**

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
<td><code>ts_debug</code> ( [ <code>config</code> <code>regconfig</code>, ] <code>document</code> <code>text</code> ) <code>setof record</code> ( <code>alias</code> <code>text</code>, <code>description</code> <code>text</code>, <code>token</code> <code>text</code>, <code>dictionaries</code> <code>regdictionary[]</code>, <code>dictionary</code> <code>regdictionary</code>, <code>lexemes</code> <code>text[]</code> )</td>
<td>Extracts and normalizes tokens from the <code>document</code> according to the specified or default text search configuration, and returns information about how each token was processed. See <a href="../full-text-search/testing-and-debugging-text-search.md#textsearch-configuration-testing">Configuration Testing</a> for details.</td>
<td><code>ts_debug('english', 'The Brightest supernovaes')</code> <code>(asciiword,"Word, all ASCII",The,{english_stem},english_stem,{}) ...</code></td>
</tr>
<tr>
<td><code>ts_lexize</code> ( <code>dict</code> <code>regdictionary</code>, <code>token</code> <code>text</code> ) <code>text[]</code></td>
<td>Returns an array of replacement lexemes if the input token is known to the dictionary, or an empty array if the token is known to the dictionary but it is a stop word, or NULL if it is not a known word. See <a href="../full-text-search/testing-and-debugging-text-search.md#textsearch-dictionary-testing">Dictionary Testing</a> for details.</td>
<td><code>ts_lexize('english_stem', 'stars')</code> <code>{star}</code></td>
</tr>
<tr>
<td><code>ts_parse</code> ( <code>parser_name</code> <code>text</code>, <code>document</code> <code>text</code> ) <code>setof record</code> ( <code>tokid</code> <code>integer</code>, <code>token</code> <code>text</code> )</td>
<td>Extracts tokens from the <code>document</code> using the named parser. See <a href="../full-text-search/testing-and-debugging-text-search.md#textsearch-parser-testing">Parser Testing</a> for details.</td>
<td><code>ts_parse('default', 'foo - bar')</code> <code>(1,foo) ...</code></td>
</tr>
<tr>
<td><code>ts_parse</code> ( <code>parser_oid</code> <code>oid</code>, <code>document</code> <code>text</code> ) <code>setof record</code> ( <code>tokid</code> <code>integer</code>, <code>token</code> <code>text</code> )</td>
<td>Extracts tokens from the <code>document</code> using a parser specified by OID. See <a href="../full-text-search/testing-and-debugging-text-search.md#textsearch-parser-testing">Parser Testing</a> for details.</td>
<td><code>ts_parse(3722, 'foo - bar')</code> <code>(1,foo) ...</code></td>
</tr>
<tr>
<td><code>ts_token_type</code> ( <code>parser_name</code> <code>text</code> ) <code>setof record</code> ( <code>tokid</code> <code>integer</code>, <code>alias</code> <code>text</code>, <code>description</code> <code>text</code> )</td>
<td>Returns a table that describes each type of token the named parser can recognize. See <a href="../full-text-search/testing-and-debugging-text-search.md#textsearch-parser-testing">Parser Testing</a> for details.</td>
<td><code>ts_token_type('default')</code> <code>(1,asciiword,"Word, all ASCII") ...</code></td>
</tr>
<tr>
<td><code>ts_token_type</code> ( <code>parser_oid</code> <code>oid</code> ) <code>setof record</code> ( <code>tokid</code> <code>integer</code>, <code>alias</code> <code>text</code>, <code>description</code> <code>text</code> )</td>
<td>Returns a table that describes each type of token a parser specified by OID can recognize. See <a href="../full-text-search/testing-and-debugging-text-search.md#textsearch-parser-testing">Parser Testing</a> for details.</td>
<td><code>ts_token_type(3722)</code> <code>(1,asciiword,"Word, all ASCII") ...</code></td>
</tr>
<tr>
<td><code>ts_stat</code> ( <code>sqlquery</code> <code>text</code> [, <code>weights</code> <code>text</code> ] ) <code>setof record</code> ( <code>word</code> <code>text</code>, <code>ndoc</code> <code>integer</code>, <code>nentry</code> <code>integer</code> )</td>
<td>Executes the <code>sqlquery</code>, which must return a single <code>tsvector</code> column, and returns statistics about each distinct lexeme contained in the data. See <a href="../full-text-search/additional-features.md#textsearch-statistics">Gathering Document Statistics</a> for details.</td>
<td><code>ts_stat('SELECT vector FROM apod')</code> <code>(foo,10,15) ...</code></td>
</tr>
</tbody>
</table>
