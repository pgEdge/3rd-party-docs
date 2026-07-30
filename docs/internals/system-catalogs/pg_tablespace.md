<a id="catalog-pg-tablespace"></a>

## `pg_tablespace`


 The catalog `pg_tablespace` stores information about the available tablespaces. Tables can be placed in particular tablespaces to aid administration of disk layout.


 Unlike most system catalogs, `pg_tablespace` is shared across all databases of a cluster: there is only one copy of `pg_tablespace` per cluster, not one per database.


**Table: `pg_tablespace` Columns**

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
<td><p><code>spcname</code> <code>name</code></p>
<p>Tablespace name</p></td>
</tr>
<tr>
<td><p><code>spcowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the tablespace, usually the user who created it</p></td>
</tr>
<tr>
<td><p><code>spcacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
<tr>
<td><p><code>spcoptions</code> <code>text[]</code></p>
<p>Tablespace-level options, as “keyword=value” strings</p></td>
</tr>
</tbody>
</table>
