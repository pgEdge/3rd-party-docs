<a id="catalog-pg-init-privs"></a>

## `pg_init_privs`


 The catalog `pg_init_privs` records information about the initial privileges of objects in the system. There is one entry for each object in the database which has a non-default (non-NULL) initial set of privileges.


 Objects can have initial privileges either by having those privileges set when the system is initialized (by initdb) or when the object is created during a [`CREATE EXTENSION`](../../reference/sql-commands/create-extension.md#sql-createextension) and the extension script sets initial privileges using the [`GRANT`](../../reference/sql-commands/grant.md#sql-grant) system. Note that the system will automatically handle recording of the privileges during the extension script and that extension authors need only use the `GRANT` and `REVOKE` statements in their script to have the privileges recorded. The `privtype` column indicates if the initial privilege was set by initdb or during a `CREATE EXTENSION` command.


 Objects which have initial privileges set by initdb will have entries where `privtype` is `'i'`, while objects which have initial privileges set by `CREATE EXTENSION` will have entries where `privtype` is `'e'`.


**Table: `pg_init_privs` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>objoid</code> <code>oid</code> (references any OID column)</p>
<p>The OID of the specific object</p></td>
</tr>
<tr>
<td><p><code>classoid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the system catalog the object is in</p></td>
</tr>
<tr>
<td><p><code>objsubid</code> <code>int4</code></p>
<p>For a table column, this is the column number (the <code>objoid</code> and <code>classoid</code> refer to the table itself). For all other object types, this column is zero.</p></td>
</tr>
<tr>
<td><p><code>privtype</code> <code>char</code></p>
<p>A code defining the type of initial privilege of this object; see text</p></td>
</tr>
<tr>
<td><p><code>initprivs</code> <code>aclitem[]</code></p>
<p>The initial access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>
