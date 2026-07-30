<a id="infoschema-administrable-role-authorizations"></a>

## `administrable_role_​authorizations`


 The view `administrable_role_authorizations` identifies all roles that the current user has the admin option for.


**Table: `administrable_role_authorizations` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>grantee</code> <code>sql_identifier</code></p>
<p>Name of the role to which this role membership was granted (can be the current user, or a different role in case of nested role memberships)</p></td>
</tr>
<tr>
<td><p><code>role_name</code> <code>sql_identifier</code></p>
<p>Name of a role</p></td>
</tr>
<tr>
<td><p><code>is_grantable</code> <code>yes_or_no</code></p>
<p>Always <code>YES</code></p></td>
</tr>
</tbody>
</table>
