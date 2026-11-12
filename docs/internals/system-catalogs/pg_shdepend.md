<a id="catalog-pg-shdepend"></a>

## `pg_shdepend`


 The catalog `pg_shdepend` records the dependency relationships between database objects and shared objects, such as roles. This information allows PostgreSQL to ensure that those objects are unreferenced before attempting to delete them.


 See also [`pg_depend`](pg_depend.md#catalog-pg-depend), which performs a similar function for dependencies involving objects within a single database.


 Unlike most system catalogs, `pg_shdepend` is shared across all databases of a cluster: there is only one copy of `pg_shdepend` per cluster, not one per database.


**Table: `pg_shdepend` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>dbid</code> <code>oid</code> (references <a href="pg_database.md#catalog-pg-database"><code>pg_database</code></a>.<code>oid</code>)</p>
<p>The OID of the database the dependent object is in, or zero for a shared object</p></td>
</tr>
<tr>
<td><p><code>classid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the system catalog the dependent object is in</p></td>
</tr>
<tr>
<td><p><code>objid</code> <code>oid</code> (references any OID column)</p>
<p>The OID of the specific dependent object</p></td>
</tr>
<tr>
<td><p><code>objsubid</code> <code>int4</code></p>
<p>For a table column, this is the column number (the <code>objid</code> and <code>classid</code> refer to the table itself). For all other object types, this column is zero.</p></td>
</tr>
<tr>
<td><p><code>refclassid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the system catalog the referenced object is in (must be a shared catalog)</p></td>
</tr>
<tr>
<td><p><code>refobjid</code> <code>oid</code> (references any OID column)</p>
<p>The OID of the specific referenced object</p></td>
</tr>
<tr>
<td><p><code>deptype</code> <code>char</code></p>
<p>A code defining the specific semantics of this dependency relationship; see text</p></td>
</tr>
</tbody>
</table>


 In all cases, a `pg_shdepend` entry indicates that the referenced object cannot be dropped without also dropping the dependent object. However, there are several subflavors identified by `deptype`:

`SHARED_DEPENDENCY_OWNER` (`o`)
:   The referenced object (which must be a role) is the owner of the dependent object.

`SHARED_DEPENDENCY_ACL` (`a`)
:   The referenced object (which must be a role) is mentioned in the ACL of the dependent object. (A `SHARED_DEPENDENCY_ACL` entry is not made for the owner of the object, since the owner will have a `SHARED_DEPENDENCY_OWNER` entry anyway.)

`SHARED_DEPENDENCY_INITACL` (`i`)
:   The referenced object (which must be a role) is mentioned in a [`pg_init_privs`](pg_init_privs.md#catalog-pg-init-privs) entry for the dependent object.

`SHARED_DEPENDENCY_POLICY` (`r`)
:   The referenced object (which must be a role) is mentioned as the target of a dependent policy object.

`SHARED_DEPENDENCY_TABLESPACE` (`t`)
:   The referenced object (which must be a tablespace) is mentioned as the tablespace for a relation that doesn't have storage.
 Other dependency flavors might be needed in future. Note in particular that the current definition only supports roles and tablespaces as referenced objects.


 As in the `pg_depend` catalog, most objects created during initdb are considered “pinned”. No entries are made in `pg_shdepend` that would have a pinned object as either referenced or dependent object.
