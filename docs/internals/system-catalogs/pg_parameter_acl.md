<a id="catalog-pg-parameter-acl"></a>

## `pg_parameter_acl`


 The catalog `pg_parameter_acl` records configuration parameters for which privileges have been granted to one or more roles. No entry is made for parameters that have default privileges.


 Unlike most system catalogs, `pg_parameter_acl` is shared across all databases of a cluster: there is only one copy of `pg_parameter_acl` per cluster, not one per database.


**Table: `pg_parameter_acl` Columns**

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
<td><p><code>parname</code> <code>text</code></p>
<p>The name of a configuration parameter for which privileges are granted</p></td>
</tr>
<tr>
<td><p><code>paracl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>
