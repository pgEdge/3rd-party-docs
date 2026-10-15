<a id="view-pg-policies"></a>

## `pg_policies`


 The view `pg_policies` provides access to useful information about each row-level security policy in the database.


**Table: `pg_policies` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>schemaname</code> <code>name</code> (references <a href="../system-catalogs/pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>nspname</code>)</p>
<p>Name of schema containing table policy is on</p></td>
</tr>
<tr>
<td><p><code>tablename</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of table policy is on</p></td>
</tr>
<tr>
<td><p><code>policyname</code> <code>name</code> (references <a href="../system-catalogs/pg_policy.md#catalog-pg-policy"><code>pg_policy</code></a>.<code>polname</code>)</p>
<p>Name of policy</p></td>
</tr>
<tr>
<td><p><code>permissive</code> <code>text</code></p>
<p>Is the policy permissive or restrictive?</p></td>
</tr>
<tr>
<td><p><code>roles</code> <code>name[]</code></p>
<p>The roles to which this policy applies</p></td>
</tr>
<tr>
<td><p><code>cmd</code> <code>text</code></p>
<p>The command type to which the policy is applied</p></td>
</tr>
<tr>
<td><p><code>qual</code> <code>text</code></p>
<p>The expression added to the security barrier qualifications for queries that this policy applies to</p></td>
</tr>
<tr>
<td><p><code>with_check</code> <code>text</code></p>
<p>The expression added to the WITH CHECK qualifications for queries that attempt to add rows to this table</p></td>
</tr>
</tbody>
</table>
