<a id="catalog-pg-authid"></a>

## `pg_authid`


 The catalog `pg_authid` contains information about database authorization identifiers (roles). A role subsumes the concepts of “users” and “groups”. A user is essentially just a role with the `rolcanlogin` flag set. Any role (with or without `rolcanlogin`) can have other roles as members; see [`pg_auth_members`](pg_auth_members.md#catalog-pg-auth-members).


 Since this catalog contains passwords, it must not be publicly readable. [`pg_roles`](../system-views/pg_roles.md#view-pg-roles) is a publicly readable view on `pg_authid` that blanks out the password field.


 [Database Roles](../../server-administration/database-roles/index.md#user-manag) contains detailed information about user and privilege management.


 Because user identities are cluster-wide, `pg_authid` is shared across all databases of a cluster: there is only one copy of `pg_authid` per cluster, not one per database.


**Table: `pg_authid` Columns**

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
<p>Role can log in. That is, this role can be given as the initial session authorization identifier.</p></td>
</tr>
<tr>
<td><p><code>rolreplication</code> <code>bool</code></p>
<p>Role is a replication role. A replication role can initiate replication connections and create and drop replication slots.</p></td>
</tr>
<tr>
<td><p><code>rolbypassrls</code> <code>bool</code></p>
<p>Role bypasses every row-level security policy, see <a href="../../the-sql-language/data-definition/row-security-policies.md#ddl-rowsecurity">Row Security Policies</a> for more information.</p></td>
</tr>
<tr>
<td><p><code>rolconnlimit</code> <code>int4</code></p>
<p>For roles that can log in, this sets maximum number of concurrent connections this role can make. -1 means no limit.</p></td>
</tr>
<tr>
<td><p><code>rolpassword</code> <code>text</code></p>
<p>Encrypted password; null if none. The format depends on the form of encryption used.</p></td>
</tr>
<tr>
<td><p><code>rolvaliduntil</code> <code>timestamptz</code></p>
<p>Password expiry time (only used for password authentication); null if no expiration</p></td>
</tr>
</tbody>
</table>


 For an MD5 encrypted password, `rolpassword` column will begin with the string `md5` followed by a 32-character hexadecimal MD5 hash. The MD5 hash will be of the user's password concatenated to their user name. For example, if user `joe` has password `xyzzy`, PostgreSQL will store the md5 hash of `xyzzyjoe`.


!!! warning

    Support for MD5-encrypted passwords is deprecated and will be removed in a future release of PostgreSQL. Refer to [Password Authentication](../../server-administration/client-authentication/password-authentication.md#auth-password) for details about migrating to another password type.


 If the password is encrypted with SCRAM-SHA-256, it has the format:

```

SCRAM-SHA-256$:$:
```
 where *salt*, *StoredKey* and *ServerKey* are in Base64 encoded format. This format is the same as that specified by [RFC 5803](https://datatracker.ietf.org/doc/html/rfc5803).
