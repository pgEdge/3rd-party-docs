<a id="infoschema-role-routine-grants"></a>

## `role_routine_grants`


 The view `role_routine_grants` identifies all privileges granted on functions where the grantor or grantee is a currently enabled role. Further information can be found under `routine_privileges`. The only effective difference between this view and `routine_privileges` is that this view omits functions that have been made accessible to the current user by way of a grant to `PUBLIC`.


**Table: `role_routine_grants` Columns**

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
<td><p><code>specific_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the function (always the current database)</p></td>
</tr>
<tr>
<td><p><code>specific_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the function</p></td>
</tr>
<tr>
<td><p><code>specific_name</code> <code>sql_identifier</code></p>
<p>The “specific name” of the function. See <a href="routines.md#infoschema-routines"><code>routines</code></a> for more information.</p></td>
</tr>
<tr>
<td><p><code>routine_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the function (always the current database)</p></td>
</tr>
<tr>
<td><p><code>routine_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the function</p></td>
</tr>
<tr>
<td><p><code>routine_name</code> <code>sql_identifier</code></p>
<p>Name of the function (might be duplicated in case of overloading)</p></td>
</tr>
<tr>
<td><p><code>privilege_type</code> <code>character_data</code></p>
<p>Always <code>EXECUTE</code> (the only privilege type for functions)</p></td>
</tr>
<tr>
<td><p><code>is_grantable</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the privilege is grantable, <code>NO</code> if not</p></td>
</tr>
</tbody>
</table>
