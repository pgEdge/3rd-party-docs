<a id="catalog-pg-ts-parser"></a>

## `pg_ts_parser`


 The `pg_ts_parser` catalog contains entries defining text search parsers. A parser is responsible for splitting input text into lexemes and assigning a token type to each lexeme. Since a parser must be implemented by C-language-level functions, creation of new parsers is restricted to database superusers.


 PostgreSQL's text search features are described at length in [Full Text Search](../../the-sql-language/full-text-search/index.md#textsearch).


**Table: `pg_ts_parser` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>oid</code> <code>oid</code></p>
<p>Row identifier</p></td>
</tr>
<tr>
<td><p><code>prsname</code> <code>name</code></p>
<p>Text search parser name</p></td>
</tr>
<tr>
<td><p><code>prsnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this parser</p></td>
</tr>
<tr>
<td><p><code>prsstart</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the parser's startup function</p></td>
</tr>
<tr>
<td><p><code>prstoken</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the parser's next-token function</p></td>
</tr>
<tr>
<td><p><code>prsend</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the parser's shutdown function</p></td>
</tr>
<tr>
<td><p><code>prsheadline</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the parser's headline function (zero if none)</p></td>
</tr>
<tr>
<td><p><code>prslextype</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the parser's lextype function</p></td>
</tr>
</tbody>
</table>
