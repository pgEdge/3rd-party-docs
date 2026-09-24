<a id="catalog-pg-rewrite"></a>

## `pg_rewrite`


 The catalog `pg_rewrite` stores rewrite rules for tables and views.


**Table: `pg_rewrite` Columns**

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
<td><p><code>rulename</code> <code>name</code></p>
<p>Rule name</p></td>
</tr>
<tr>
<td><p><code>ev_class</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table this rule is for</p></td>
</tr>
<tr>
<td><p><code>ev_type</code> <code>char</code></p>
<p>Event type that the rule is for: 1 = <a href="../../reference/sql-commands/select.md#sql-select">sql-select</a>, 2 = <a href="../../reference/sql-commands/update.md#sql-update">sql-update</a>, 3 = <a href="../../reference/sql-commands/insert.md#sql-insert">sql-insert</a>, 4 = <a href="../../reference/sql-commands/delete.md#sql-delete">sql-delete</a></p></td>
</tr>
<tr>
<td><p><code>ev_enabled</code> <code>char</code></p>
<p>Controls in which <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-session-replication-role">session_replication_role</a> modes the rule fires. <code>O</code> = rule fires in “origin” and “local” modes, <code>D</code> = rule is disabled, <code>R</code> = rule fires in “replica” mode, <code>A</code> = rule fires always.</p></td>
</tr>
<tr>
<td><p><code>is_instead</code> <code>bool</code></p>
<p>True if the rule is an <code>INSTEAD</code> rule</p></td>
</tr>
<tr>
<td><p><code>ev_qual</code> <code>pg_node_tree</code></p>
<p>Expression tree (in the form of a <code>nodeToString()</code> representation) for the rule's qualifying condition</p></td>
</tr>
<tr>
<td><p><code>ev_action</code> <code>pg_node_tree</code></p>
<p>Query tree (in the form of a <code>nodeToString()</code> representation) for the rule's action</p></td>
</tr>
</tbody>
</table>


!!! note

    `pg_class.relhasrules` must be true if a table has any rules in this catalog.
