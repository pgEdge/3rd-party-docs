<a id="catalog-pg-conversion"></a>

## `pg_conversion`


 The catalog `pg_conversion` describes encoding conversion functions. See [sql-createconversion](../../reference/sql-commands/create-conversion.md#sql-createconversion) for more information.


**Table: `pg_conversion` Columns**

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
<td><p><code>conname</code> <code>name</code></p>
<p>Conversion name (unique within a namespace)</p></td>
</tr>
<tr>
<td><p><code>connamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this conversion</p></td>
</tr>
<tr>
<td><p><code>conowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the conversion</p></td>
</tr>
<tr>
<td><p><code>conforencoding</code> <code>int4</code></p>
<p>Source encoding ID (<a href="../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#pg-encoding-to-char"><code>pg_encoding_to_char()</code></a> can translate this number to the encoding name)</p></td>
</tr>
<tr>
<td><p><code>contoencoding</code> <code>int4</code></p>
<p>Destination encoding ID (<a href="../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#pg-encoding-to-char"><code>pg_encoding_to_char()</code></a> can translate this number to the encoding name)</p></td>
</tr>
<tr>
<td><p><code>conproc</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Conversion function</p></td>
</tr>
<tr>
<td><p><code>condefault</code> <code>bool</code></p>
<p>True if this is the default conversion</p></td>
</tr>
</tbody>
</table>
