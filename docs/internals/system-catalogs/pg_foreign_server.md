<a id="catalog-pg-foreign-server"></a>

## `pg_foreign_server`


 The catalog `pg_foreign_server` stores foreign server definitions. A foreign server describes a source of external data, such as a remote server. Foreign servers are accessed via foreign-data wrappers.


**Table: `pg_foreign_server` Columns**

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
<td><p><code>srvname</code> <code>name</code></p>
<p>Name of the foreign server</p></td>
</tr>
<tr>
<td><p><code>srvowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the foreign server</p></td>
</tr>
<tr>
<td><p><code>srvfdw</code> <code>oid</code> (references <a href="pg_foreign_data_wrapper.md#catalog-pg-foreign-data-wrapper"><code>pg_foreign_data_wrapper</code></a>.<code>oid</code>)</p>
<p>OID of the foreign-data wrapper of this foreign server</p></td>
</tr>
<tr>
<td><p><code>srvtype</code> <code>text</code></p>
<p>Type of the server (optional)</p></td>
</tr>
<tr>
<td><p><code>srvversion</code> <code>text</code></p>
<p>Version of the server (optional)</p></td>
</tr>
<tr>
<td><p><code>srvacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
<tr>
<td><p><code>srvoptions</code> <code>text[]</code></p>
<p>Foreign server specific options, as “keyword=value” strings</p></td>
</tr>
</tbody>
</table>
