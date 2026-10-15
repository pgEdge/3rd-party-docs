<a id="catalog-pg-event-trigger"></a>

## `pg_event_trigger`


 The catalog `pg_event_trigger` stores event triggers. See [Event Triggers](../../server-programming/event-triggers/index.md#event-triggers) for more information.


**Table: `pg_event_trigger` Columns**

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
<td><p><code>evtname</code> <code>name</code></p>
<p>Trigger name (must be unique)</p></td>
</tr>
<tr>
<td><p><code>evtevent</code> <code>name</code></p>
<p>Identifies the event for which this trigger fires</p></td>
</tr>
<tr>
<td><p><code>evtowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the event trigger</p></td>
</tr>
<tr>
<td><p><code>evtfoid</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>The function to be called</p></td>
</tr>
<tr>
<td><p><code>evtenabled</code> <code>char</code></p>
<p>Controls in which <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-session-replication-role">session_replication_role</a> modes the event trigger fires. <code>O</code> = trigger fires in “origin” and “local” modes, <code>D</code> = trigger is disabled, <code>R</code> = trigger fires in “replica” mode, <code>A</code> = trigger fires always.</p></td>
</tr>
<tr>
<td><p><code>evttags</code> <code>text[]</code></p>
<p>Command tags for which this trigger will fire. If NULL, the firing of this trigger is not restricted on the basis of the command tag.</p></td>
</tr>
</tbody>
</table>
