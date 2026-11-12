<a id="catalog-pg-user-mapping"></a>

## `pg_user_mapping`


 The catalog `pg_user_mapping` stores the mappings from local user to remote. Access to this catalog is restricted from normal users, use the view [`pg_user_mappings`](../system-views/pg_user_mappings.md#view-pg-user-mappings) instead.


**Table: `pg_user_mapping` Columns**

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
<td><p><code>umuser</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>OID of the local role being mapped, or zero if the user mapping is public</p></td>
</tr>
<tr>
<td><p><code>umserver</code> <code>oid</code> (references <a href="pg_foreign_server.md#catalog-pg-foreign-server"><code>pg_foreign_server</code></a>.<code>oid</code>)</p>
<p>The OID of the foreign server that contains this mapping</p></td>
</tr>
<tr>
<td><p><code>umoptions</code> <code>text[]</code></p>
<p>User mapping specific options, as “keyword=value” strings</p></td>
</tr>
</tbody>
</table>
