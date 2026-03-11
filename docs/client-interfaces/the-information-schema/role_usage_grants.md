<a id="infoschema-role-usage-grants"></a>

## `role_usage_grants`


 The view `role_usage_grants` identifies `USAGE` privileges granted on various kinds of objects where the grantor or grantee is a currently enabled role. Further information can be found under `usage_privileges`. The only effective difference between this view and `usage_privileges` is that this view omits objects that have been made accessible to the current user by way of a grant to `PUBLIC`.


**Table: `role_usage_grants` Columns**

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
<p>The name of the role that granted the privilege</p></td>
</tr>
<tr>
<td><p><code>grantee</code> <code>sql_identifier</code></p>
<p>The name of the role that the privilege was granted to</p></td>
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
