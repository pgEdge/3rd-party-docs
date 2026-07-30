<a id="catalog-pg-ts-config"></a>

## `pg_ts_config`


 The `pg_ts_config` catalog contains entries representing text search configurations. A configuration specifies a particular text search parser and a list of dictionaries to use for each of the parser's output token types. The parser is shown in the `pg_ts_config` entry, but the token-to-dictionary mapping is defined by subsidiary entries in [`pg_ts_config_map`](pg_ts_config_map.md#catalog-pg-ts-config-map).


 PostgreSQL's text search features are described at length in [Full Text Search](../../the-sql-language/full-text-search/index.md#textsearch).


**Table: `pg_ts_config` Columns**

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
<td><p><code>cfgname</code> <code>name</code></p>
<p>Text search configuration name</p></td>
</tr>
<tr>
<td><p><code>cfgnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this configuration</p></td>
</tr>
<tr>
<td><p><code>cfgowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the configuration</p></td>
</tr>
<tr>
<td><p><code>cfgparser</code> <code>oid</code> (references <a href="pg_ts_parser.md#catalog-pg-ts-parser"><code>pg_ts_parser</code></a>.<code>oid</code>)</p>
<p>The OID of the text search parser for this configuration</p></td>
</tr>
</tbody>
</table>
