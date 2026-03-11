<a id="catalog-pg-default-acl"></a>

## `pg_default_acl`


 The catalog `pg_default_acl` stores initial privileges to be assigned to newly created objects.


**Table: `pg_default_acl` Columns**

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
<td><p><code>defaclrole</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>The OID of the role associated with this entry</p></td>
</tr>
<tr>
<td><p><code>defaclnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace associated with this entry, or zero if none</p></td>
</tr>
<tr>
<td><p><code>defaclobjtype</code> <code>char</code></p>
<p>Type of object this entry is for: <code>r</code> = relation (table, view), <code>S</code> = sequence, <code>f</code> = function, <code>T</code> = type, <code>n</code> = schema</p></td>
</tr>
<tr>
<td><p><code>defaclacl</code> <code>aclitem[]</code></p>
<p>Access privileges that this type of object should have on creation</p></td>
</tr>
</tbody>
</table>


 A `pg_default_acl` entry shows the initial privileges to be assigned to an object belonging to the indicated user. There are currently two types of entry: “global” entries with `defaclnamespace` = zero, and “per-schema” entries that reference a particular schema. If a global entry is present then it *overrides* the normal hard-wired default privileges for the object type. A per-schema entry, if present, represents privileges to be *added to* the global or hard-wired default privileges.


 Note that when an ACL entry in another catalog is null, it is taken to represent the hard-wired default privileges for its object, *not* whatever might be in `pg_default_acl` at the moment. `pg_default_acl` is only consulted during object creation.
