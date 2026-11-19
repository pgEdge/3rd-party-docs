<a id="catalog-pg-ts-dict"></a>

## `pg_ts_dict`


 The `pg_ts_dict` catalog contains entries defining text search dictionaries. A dictionary depends on a text search template, which specifies all the implementation functions needed; the dictionary itself provides values for the user-settable parameters supported by the template. This division of labor allows dictionaries to be created by unprivileged users. The parameters are specified by a text string `dictinitoption`, whose format and meaning vary depending on the template.


 PostgreSQL's text search features are described at length in [Full Text Search](../../the-sql-language/full-text-search/index.md#textsearch).


**Table: `pg_ts_dict` Columns**

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
<td><p><code>dictname</code> <code>name</code></p>
<p>Text search dictionary name</p></td>
</tr>
<tr>
<td><p><code>dictnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this dictionary</p></td>
</tr>
<tr>
<td><p><code>dictowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the dictionary</p></td>
</tr>
<tr>
<td><p><code>dicttemplate</code> <code>oid</code> (references <a href="pg_ts_template.md#catalog-pg-ts-template"><code>pg_ts_template</code></a>.<code>oid</code>)</p>
<p>The OID of the text search template for this dictionary</p></td>
</tr>
<tr>
<td><p><code>dictinitoption</code> <code>text</code></p>
<p>Initialization option string for the template</p></td>
</tr>
</tbody>
</table>
