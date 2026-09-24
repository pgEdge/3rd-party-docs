<a id="catalog-pg-db-role-setting"></a>

## `pg_db_role_setting`


 The catalog `pg_db_role_setting` records the default values that have been set for run-time configuration variables, for each role and database combination.


 Unlike most system catalogs, `pg_db_role_setting` is shared across all databases of a cluster: there is only one copy of `pg_db_role_setting` per cluster, not one per database.


**Table: `pg_db_role_setting` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>setdatabase</code> <code>oid</code> (references <a href="pg_database.md#catalog-pg-database"><code>pg_database</code></a>.<code>oid</code>)</p>
<p>The OID of the database the setting is applicable to, or zero if not database-specific</p></td>
</tr>
<tr>
<td><p><code>setrole</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>The OID of the role the setting is applicable to, or zero if not role-specific</p></td>
</tr>
<tr>
<td><p><code>setconfig</code> <code>text[]</code></p>
<p>Defaults for run-time configuration variables</p></td>
</tr>
</tbody>
</table>
