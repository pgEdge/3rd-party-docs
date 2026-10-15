<a id="catalog-pg-ts-template"></a>

## `pg_ts_template`


 The `pg_ts_template` catalog contains entries defining text search templates. A template is the implementation skeleton for a class of text search dictionaries. Since a template must be implemented by C-language-level functions, creation of new templates is restricted to database superusers.


 PostgreSQL's text search features are described at length in [Full Text Search](../../the-sql-language/full-text-search/index.md#textsearch).


**Table: `pg_ts_template` Columns**

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
<td><p><code>tmplname</code> <code>name</code></p>
<p>Text search template name</p></td>
</tr>
<tr>
<td><p><code>tmplnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this template</p></td>
</tr>
<tr>
<td><p><code>tmplinit</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the template's initialization function (zero if none)</p></td>
</tr>
<tr>
<td><p><code>tmpllexize</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the template's lexize function</p></td>
</tr>
</tbody>
</table>
