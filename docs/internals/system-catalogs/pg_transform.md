<a id="catalog-pg-transform"></a>

## `pg_transform`


 The catalog `pg_transform` stores information about transforms, which are a mechanism to adapt data types to procedural languages. See [sql-createtransform](../../reference/sql-commands/create-transform.md#sql-createtransform) for more information.


**Table: `pg_transform` Columns**

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
<td><p><code>trftype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>OID of the data type this transform is for</p></td>
</tr>
<tr>
<td><p><code>trflang</code> <code>oid</code> (references <a href="pg_language.md#catalog-pg-language"><code>pg_language</code></a>.<code>oid</code>)</p>
<p>OID of the language this transform is for</p></td>
</tr>
<tr>
<td><p><code>trffromsql</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>The OID of the function to use when converting the data type for input to the procedural language (e.g., function parameters). Zero is stored if the default behavior should be used.</p></td>
</tr>
<tr>
<td><p><code>trftosql</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>The OID of the function to use when converting output from the procedural language (e.g., return values) to the data type. Zero is stored if the default behavior should be used.</p></td>
</tr>
</tbody>
</table>
