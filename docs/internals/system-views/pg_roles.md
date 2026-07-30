<a id="view-pg-roles"></a>

## `pg_roles`


 The view `pg_roles` provides access to information about database roles. This is simply a publicly readable view of [`pg_authid`](../system-catalogs/pg_authid.md#catalog-pg-authid) that blanks out the password field.


**Table: `pg_roles` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>rolname</code> <code>name</code></p>
<p>Role name</p></td>
</tr>
<tr>
<td><p><code>rolsuper</code> <code>bool</code></p>
<p>Role has superuser privileges</p></td>
</tr>
<tr>
<td><p><code>rolinherit</code> <code>bool</code></p>
<p>Role automatically inherits privileges of roles it is a member of</p></td>
</tr>
<tr>
<td><p><code>rolcreaterole</code> <code>bool</code></p>
<p>Role can create more roles</p></td>
</tr>
<tr>
<td><p><code>rolcreatedb</code> <code>bool</code></p>
<p>Role can create databases</p></td>
</tr>
<tr>
<td><p><code>rolcanlogin</code> <code>bool</code></p>
<p>Role can log in. That is, this role can be given as the initial session authorization identifier</p></td>
</tr>
<tr>
<td><p><code>rolreplication</code> <code>bool</code></p>
<p>Role is a replication role. A replication role can initiate replication connections and create and drop replication slots.</p></td>
</tr>
<tr>
<td><p><code>rolconnlimit</code> <code>int4</code></p>
<p>For roles that can log in, this sets maximum number of concurrent connections this role can make. -1 means no limit.</p></td>
</tr>
<tr>
<td><p><code>rolpassword</code> <code>text</code></p>
<p>Not the password (always reads as <code>********</code>)</p></td>
</tr>
<tr>
<td><p><code>rolvaliduntil</code> <code>timestamptz</code></p>
<p>Password expiry time (only used for password authentication); null if no expiration</p></td>
</tr>
<tr>
<td><p><code>rolbypassrls</code> <code>bool</code></p>
<p>Role bypasses every row-level security policy, see <a href="../../the-sql-language/data-definition/row-security-policies.md#ddl-rowsecurity">Row Security Policies</a> for more information.</p></td>
</tr>
<tr>
<td><p><code>rolconfig</code> <code>text[]</code></p>
<p>Role-specific defaults for run-time configuration variables</p></td>
</tr>
<tr>
<td><p><code>oid</code> <code>oid</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>ID of role</p></td>
</tr>
</tbody>
</table>
