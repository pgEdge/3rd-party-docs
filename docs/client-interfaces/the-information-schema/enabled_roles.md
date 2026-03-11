<a id="infoschema-enabled-roles"></a>

## `enabled_roles`


 The view `enabled_roles` identifies the currently “enabled roles”. The enabled roles are recursively defined as the current user together with all roles that have been granted to the enabled roles with automatic inheritance. In other words, these are all roles that the current user has direct or indirect, automatically inheriting membership in.


 For permission checking, the set of “applicable roles” is applied, which can be broader than the set of enabled roles. So generally, it is better to use the view `applicable_roles` instead of this one; See [`applicable_roles`](applicable_roles.md#infoschema-applicable-roles) for details on `applicable_roles` view.


**Table: `enabled_roles` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>role_name</code> <code>sql_identifier</code></p>
<p>Name of a role</p></td>
</tr>
</tbody>
</table>
