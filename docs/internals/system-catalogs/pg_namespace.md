<a id="catalog-pg-namespace"></a>

## `pg_namespace`


 The catalog `pg_namespace` stores namespaces. A namespace is the structure underlying SQL schemas: each namespace can have a separate collection of relations, types, etc. without name conflicts.


**Table: `pg_namespace` Columns**

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
<td><p><code>nspname</code> <code>name</code></p>
<p>Name of the namespace</p></td>
</tr>
<tr>
<td><p><code>nspowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the namespace</p></td>
</tr>
<tr>
<td><p><code>nspacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>
