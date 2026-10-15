<a id="view-pg-user-mappings"></a>

## `pg_user_mappings`


 The view `pg_user_mappings` provides access to information about user mappings. This is essentially a publicly readable view of [`pg_user_mapping`](../system-catalogs/pg_user_mapping.md#catalog-pg-user-mapping) that leaves out the options field if the user has no rights to use it.


**Table: `pg_user_mappings` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>umid</code> <code>oid</code> (references <a href="../system-catalogs/pg_user_mapping.md#catalog-pg-user-mapping"><code>pg_user_mapping</code></a>.<code>oid</code>)</p>
<p>OID of the user mapping</p></td>
</tr>
<tr>
<td><p><code>srvid</code> <code>oid</code> (references <a href="../system-catalogs/pg_foreign_server.md#catalog-pg-foreign-server"><code>pg_foreign_server</code></a>.<code>oid</code>)</p>
<p>The OID of the foreign server that contains this mapping</p></td>
</tr>
<tr>
<td><p><code>srvname</code> <code>name</code> (references <a href="../system-catalogs/pg_foreign_server.md#catalog-pg-foreign-server"><code>pg_foreign_server</code></a>.<code>srvname</code>)</p>
<p>Name of the foreign server</p></td>
</tr>
<tr>
<td><p><code>umuser</code> <code>oid</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>OID of the local role being mapped, or zero if the user mapping is public</p></td>
</tr>
<tr>
<td><p><code>usename</code> <code>name</code></p>
<p>Name of the local user to be mapped</p></td>
</tr>
<tr>
<td><p><code>umoptions</code> <code>text[]</code></p>
<p>User mapping specific options, as “keyword=value” strings</p></td>
</tr>
</tbody>
</table>


 To protect password information stored as a user mapping option, the `umoptions` column will read as null unless one of the following applies:

-  current user is the user being mapped, and owns the server or holds `USAGE` privilege on it
-  current user is the server owner and mapping is for `PUBLIC`
-  current user is a superuser
