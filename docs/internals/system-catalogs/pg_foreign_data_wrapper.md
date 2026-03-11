<a id="catalog-pg-foreign-data-wrapper"></a>

## `pg_foreign_data_wrapper`


 The catalog `pg_foreign_data_wrapper` stores foreign-data wrapper definitions. A foreign-data wrapper is the mechanism by which external data, residing on foreign servers, is accessed.


**Table: `pg_foreign_data_wrapper` Columns**

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
<td><p><code>fdwname</code> <code>name</code></p>
<p>Name of the foreign-data wrapper</p></td>
</tr>
<tr>
<td><p><code>fdwowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the foreign-data wrapper</p></td>
</tr>
<tr>
<td><p><code>fdwhandler</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>References a handler function that is responsible for supplying execution routines for the foreign-data wrapper. Zero if no handler is provided</p></td>
</tr>
<tr>
<td><p><code>fdwvalidator</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>References a validator function that is responsible for checking the validity of the options given to the foreign-data wrapper, as well as options for foreign servers and user mappings using the foreign-data wrapper. Zero if no validator is provided</p></td>
</tr>
<tr>
<td><p><code>fdwacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
<tr>
<td><p><code>fdwoptions</code> <code>text[]</code></p>
<p>Foreign-data wrapper specific options, as “keyword=value” strings</p></td>
</tr>
</tbody>
</table>
