<a id="view-pg-user"></a>

## `pg_user`


 The view `pg_user` provides access to information about database users. This is simply a publicly readable view of [`pg_shadow`](pg_shadow.md#view-pg-shadow) that blanks out the password field.


**Table: `pg_user` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>usename</code> <code>name</code></p>
<p>User name</p></td>
</tr>
<tr>
<td><p><code>usesysid</code> <code>oid</code></p>
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
<p>Not the password (always reads as <code>********</code>)</p></td>
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
