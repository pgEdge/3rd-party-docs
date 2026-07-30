<a id="catalog-pg-ts-config-map"></a>

## `pg_ts_config_map`


 The `pg_ts_config_map` catalog contains entries showing which text search dictionaries should be consulted, and in what order, for each output token type of each text search configuration's parser.


 PostgreSQL's text search features are described at length in [Full Text Search](../../the-sql-language/full-text-search/index.md#textsearch).


**Table: `pg_ts_config_map` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>mapcfg</code> <code>oid</code> (references <a href="pg_ts_config.md#catalog-pg-ts-config"><code>pg_ts_config</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_ts_config.md#catalog-pg-ts-config"><code>pg_ts_config</code></a> entry owning this map entry</p></td>
</tr>
<tr>
<td><p><code>maptokentype</code> <code>int4</code></p>
<p>A token type emitted by the configuration's parser</p></td>
</tr>
<tr>
<td><p><code>mapseqno</code> <code>int4</code></p>
<p>Order in which to consult this entry (lower <code>mapseqno</code>s first)</p></td>
</tr>
<tr>
<td><p><code>mapdict</code> <code>oid</code> (references <a href="pg_ts_dict.md#catalog-pg-ts-dict"><code>pg_ts_dict</code></a>.<code>oid</code>)</p>
<p>The OID of the text search dictionary to consult</p></td>
</tr>
</tbody>
</table>
