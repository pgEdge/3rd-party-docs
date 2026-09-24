<a id="catalog-pg-trigger"></a>

## `pg_trigger`


 The catalog `pg_trigger` stores triggers on tables and views. See [sql-createtrigger](../../reference/sql-commands/create-trigger.md#sql-createtrigger) for more information.


**Table: `pg_trigger` Columns**

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
<td><p><code>tgrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table this trigger is on</p></td>
</tr>
<tr>
<td><p><code>tgparentid</code> <code>oid</code> (references <a href="#catalog-pg-trigger"><code>pg_trigger</code></a>.<code>oid</code>)</p>
<p>Parent trigger that this trigger is cloned from (this happens when partitions are created or attached to a partitioned table); zero if not a clone</p></td>
</tr>
<tr>
<td><p><code>tgname</code> <code>name</code></p>
<p>Trigger name (must be unique among triggers of same table)</p></td>
</tr>
<tr>
<td><p><code>tgfoid</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>The function to be called</p></td>
</tr>
<tr>
<td><p><code>tgtype</code> <code>int2</code></p>
<p>Bit mask identifying trigger firing conditions</p></td>
</tr>
<tr>
<td><p><code>tgenabled</code> <code>char</code></p>
<p>Controls in which <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-session-replication-role">session_replication_role</a> modes the trigger fires. <code>O</code> = trigger fires in “origin” and “local” modes, <code>D</code> = trigger is disabled, <code>R</code> = trigger fires in “replica” mode, <code>A</code> = trigger fires always.</p></td>
</tr>
<tr>
<td><p><code>tgisinternal</code> <code>bool</code></p>
<p>True if trigger is internally generated (usually, to enforce the constraint identified by <code>tgconstraint</code>)</p></td>
</tr>
<tr>
<td><p><code>tgconstrrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table referenced by a referential integrity constraint (zero if trigger is not for a referential integrity constraint)</p></td>
</tr>
<tr>
<td><p><code>tgconstrindid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The index supporting a unique, primary key, referential integrity, or exclusion constraint (zero if trigger is not for one of these types of constraint)</p></td>
</tr>
<tr>
<td><p><code>tgconstraint</code> <code>oid</code> (references <a href="pg_constraint.md#catalog-pg-constraint"><code>pg_constraint</code></a>.<code>oid</code>)</p>
<p>The <a href="pg_constraint.md#catalog-pg-constraint"><code>pg_constraint</code></a> entry associated with the trigger (zero if trigger is not for a constraint)</p></td>
</tr>
<tr>
<td><p><code>tgdeferrable</code> <code>bool</code></p>
<p>True if constraint trigger is deferrable</p></td>
</tr>
<tr>
<td><p><code>tginitdeferred</code> <code>bool</code></p>
<p>True if constraint trigger is initially deferred</p></td>
</tr>
<tr>
<td><p><code>tgnargs</code> <code>int2</code></p>
<p>Number of argument strings passed to trigger function</p></td>
</tr>
<tr>
<td><p><code>tgattr</code> <code>int2vector</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>Column numbers, if trigger is column-specific; otherwise an empty array</p></td>
</tr>
<tr>
<td><p><code>tgargs</code> <code>bytea</code></p>
<p>Argument strings to pass to trigger, each NULL-terminated</p></td>
</tr>
<tr>
<td><p><code>tgqual</code> <code>pg_node_tree</code></p>
<p>Expression tree (in <code>nodeToString()</code> representation) for the trigger's <code>WHEN</code> condition, or null if none</p></td>
</tr>
<tr>
<td><p><code>tgoldtable</code> <code>name</code></p>
<p><code>REFERENCING</code> clause name for <code>OLD TABLE</code>, or null if none</p></td>
</tr>
<tr>
<td><p><code>tgnewtable</code> <code>name</code></p>
<p><code>REFERENCING</code> clause name for <code>NEW TABLE</code>, or null if none</p></td>
</tr>
</tbody>
</table>


 Currently, column-specific triggering is supported only for `UPDATE` events, and so `tgattr` is relevant only for that event type. `tgtype` might contain bits for other event types as well, but those are presumed to be table-wide regardless of what is in `tgattr`.


!!! note

    When `tgconstraint` is nonzero, `tgconstrrelid`, `tgconstrindid`, `tgdeferrable`, and `tginitdeferred` are largely redundant with the referenced [`pg_constraint`](pg_constraint.md#catalog-pg-constraint) entry. However, it is possible for a non-deferrable trigger to be associated with a deferrable constraint: foreign key constraints can have some deferrable and some non-deferrable triggers.


!!! note

    `pg_class.relhastriggers` must be true if a relation has any triggers in this catalog.
