<a id="view-pg-shadow"></a>

## `pg_shadow`


 The view `pg_shadow` exists for backwards compatibility: it emulates a catalog that existed in PostgreSQL before version 8.1. It shows properties of all roles that are marked as `rolcanlogin` in [`pg_authid`](../system-catalogs/pg_authid.md#catalog-pg-authid).


 The name stems from the fact that this table should not be readable by the public since it contains passwords. [`pg_user`](pg_user.md#view-pg-user) is a publicly readable view on `pg_shadow` that blanks out the password field.


**Table: `pg_shadow` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>usename</code> <code>name</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>rolname</code>)</p>
<p>User name</p></td>
</tr>
<tr>
<td><p><code>usesysid</code> <code>oid</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>ID of this user</p></td>
</tr>
<tr>
<td><p><code>usecreatedb</code> <code>bool</code></p>
<p>User can create databases</p></td>
</tr>
<tr>
<td><p><code>usesuper</code> <code>bool</code></p>
<p>User is a superuser</p></td>
</tr>
<tr>
<td><p><code>userepl</code> <code>bool</code></p>
<p>User can initiate streaming replication and put the system in and out of backup mode.</p></td>
</tr>
<tr>
<td><p><code>usebypassrls</code> <code>bool</code></p>
<p>User bypasses every row-level security policy, see <a href="../../the-sql-language/data-definition/row-security-policies.md#ddl-rowsecurity">Row Security Policies</a> for more information.</p></td>
</tr>
<tr>
<td><p><code>passwd</code> <code>text</code></p>
<p>Encrypted password; null if none. See <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a> for details of how encrypted passwords are stored.</p></td>
</tr>
<tr>
<td><p><code>valuntil</code> <code>timestamptz</code></p>
<p>Password expiry time (only used for password authentication)</p></td>
</tr>
<tr>
<td><p><code>useconfig</code> <code>text[]</code></p>
<p>Session defaults for run-time configuration variables</p></td>
</tr>
</tbody>
</table>
