<a id="catalog-pg-policy"></a>

## `pg_policy`


 The catalog `pg_policy` stores row-level security policies for tables. A policy includes the kind of command that it applies to (possibly all commands), the roles that it applies to, the expression to be added as a security-barrier qualification to queries that include the table, and the expression to be added as a `WITH CHECK` option for queries that attempt to add new records to the table.


**Table: `pg_policy` Columns**

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
<td><p><code>polname</code> <code>name</code></p>
<p>The name of the policy</p></td>
</tr>
<tr>
<td><p><code>polrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table to which the policy applies</p></td>
</tr>
<tr>
<td><p><code>polcmd</code> <code>char</code></p>
<p>The command type to which the policy is applied: <code>r</code> for <a href="../../reference/sql-commands/select.md#sql-select">sql-select</a>, <code>a</code> for <a href="../../reference/sql-commands/insert.md#sql-insert">sql-insert</a>, <code>w</code> for <a href="../../reference/sql-commands/update.md#sql-update">sql-update</a>, <code>d</code> for <a href="../../reference/sql-commands/delete.md#sql-delete">sql-delete</a>, or <code>*</code> for all</p></td>
</tr>
<tr>
<td><p><code>polpermissive</code> <code>bool</code></p>
<p>Is the policy permissive or restrictive?</p></td>
</tr>
<tr>
<td><p><code>polroles</code> <code>oid[]</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>The roles to which the policy is applied; zero means <code>PUBLIC</code> (and normally appears alone in the array)</p></td>
</tr>
<tr>
<td><p><code>polqual</code> <code>pg_node_tree</code></p>
<p>The expression tree to be added to the security barrier qualifications for queries that use the table</p></td>
</tr>
<tr>
<td><p><code>polwithcheck</code> <code>pg_node_tree</code></p>
<p>The expression tree to be added to the WITH CHECK qualifications for queries that attempt to add rows to the table</p></td>
</tr>
</tbody>
</table>


!!! note

    Policies stored in `pg_policy` are applied only when [`pg_class`](pg_class.md#catalog-pg-class).`relrowsecurity` is set for their table.
