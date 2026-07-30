<a id="catalog-pg-extension"></a>

## `pg_extension`


 The catalog `pg_extension` stores information about the installed extensions. See [Packaging Related Objects into an Extension](../../server-programming/extending-sql/packaging-related-objects-into-an-extension.md#extend-extensions) for details about extensions.


**Table: `pg_extension` Columns**

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
<td><p><code>extname</code> <code>name</code></p>
<p>Name of the extension</p></td>
</tr>
<tr>
<td><p><code>extowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the extension</p></td>
</tr>
<tr>
<td><p><code>extnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>Schema containing the extension's exported objects</p></td>
</tr>
<tr>
<td><p><code>extrelocatable</code> <code>bool</code></p>
<p>True if extension can be relocated to another schema</p></td>
</tr>
<tr>
<td><p><code>extversion</code> <code>text</code></p>
<p>Version name for the extension</p></td>
</tr>
<tr>
<td><p><code>extconfig</code> <code>oid[]</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Array of <code>regclass</code> OIDs for the extension's configuration table(s), or <code>NULL</code> if none</p></td>
</tr>
<tr>
<td><p><code>extcondition</code> <code>text[]</code></p>
<p>Array of <code>WHERE</code>-clause filter conditions for the extension's configuration table(s), or <code>NULL</code> if none</p></td>
</tr>
</tbody>
</table>


 Note that unlike most catalogs with a “namespace” column, `extnamespace` is not meant to imply that the extension belongs to that schema. Extension names are never schema-qualified. Rather, `extnamespace` indicates the schema that contains most or all of the extension's objects. If `extrelocatable` is true, then this schema must in fact contain all schema-qualifiable objects belonging to the extension.
