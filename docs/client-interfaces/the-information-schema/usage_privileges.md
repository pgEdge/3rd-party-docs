<a id="infoschema-usage-privileges"></a>

## `usage_privileges`


 The view `usage_privileges` identifies `USAGE` privileges granted on various kinds of objects to a currently enabled role or by a currently enabled role. In PostgreSQL, this currently applies to collations, domains, foreign-data wrappers, foreign servers, and sequences. There is one row for each combination of object, grantor, and grantee.


 Since collations do not have real privileges in PostgreSQL, this view shows implicit non-grantable `USAGE` privileges granted by the owner to `PUBLIC` for all collations. The other object types, however, show real privileges.


 In PostgreSQL, sequences also support `SELECT` and `UPDATE` privileges in addition to the `USAGE` privilege. These are nonstandard and therefore not visible in the information schema.


**Table: `usage_privileges` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>grantor</code> <code>sql_identifier</code></p>
<p>Name of the role that granted the privilege</p></td>
</tr>
<tr>
<td><p><code>grantee</code> <code>sql_identifier</code></p>
<p>Name of the role that the privilege was granted to</p></td>
</tr>
<tr>
<td><p><code>object_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the object (always the current database)</p></td>
</tr>
<tr>
<td><p><code>object_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the object, if applicable, else an empty string</p></td>
</tr>
<tr>
<td><p><code>object_name</code> <code>sql_identifier</code></p>
<p>Name of the object</p></td>
</tr>
<tr>
<td><p><code>object_type</code> <code>character_data</code></p>
<p><code>COLLATION</code> or <code>DOMAIN</code> or <code>FOREIGN DATA WRAPPER</code> or <code>FOREIGN SERVER</code> or <code>SEQUENCE</code></p></td>
</tr>
<tr>
<td><p><code>privilege_type</code> <code>character_data</code></p>
<p>Always <code>USAGE</code></p></td>
</tr>
<tr>
<td><p><code>is_grantable</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the privilege is grantable, <code>NO</code> if not</p></td>
</tr>
</tbody>
</table>
