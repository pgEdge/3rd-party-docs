<a id="catalog-pg-largeobject-metadata"></a>

## `pg_largeobject_metadata`


 The catalog `pg_largeobject_metadata` holds metadata associated with large objects. The actual large object data is stored in [`pg_largeobject`](pg_largeobject.md#catalog-pg-largeobject).


**Table: `pg_largeobject_metadata` Columns**

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
<td><p><code>lomowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the large object</p></td>
</tr>
<tr>
<td><p><code>lomacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>
