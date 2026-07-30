<a id="catalog-pg-auth-members"></a>

## `pg_auth_members`


 The catalog `pg_auth_members` shows the membership relations between roles. Any non-circular set of relationships is allowed.


 Because user identities are cluster-wide, `pg_auth_members` is shared across all databases of a cluster: there is only one copy of `pg_auth_members` per cluster, not one per database.


**Table: `pg_auth_members` Columns**

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
<td><p><code>roleid</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>ID of a role that has a member</p></td>
</tr>
<tr>
<td><p><code>member</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>ID of a role that is a member of <code>roleid</code></p></td>
</tr>
<tr>
<td><p><code>grantor</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>ID of the role that granted this membership</p></td>
</tr>
<tr>
<td><p><code>admin_option</code> <code>bool</code></p>
<p>True if <code>member</code> can grant membership in <code>roleid</code> to others</p></td>
</tr>
<tr>
<td><p><code>inherit_option</code> <code>bool</code></p>
<p>True if the member automatically inherits the privileges of the granted role</p></td>
</tr>
<tr>
<td><p><code>set_option</code> <code>bool</code></p>
<p>True if the member can <a href="../../reference/sql-commands/set-role.md#sql-set-role"><code>SET ROLE</code></a> to the granted role</p></td>
</tr>
</tbody>
</table>
