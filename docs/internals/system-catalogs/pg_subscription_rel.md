<a id="catalog-pg-subscription-rel"></a>

## `pg_subscription_rel`


 The catalog `pg_subscription_rel` stores the state of each replicated table and sequence for each subscription. This is a many-to-many mapping.


 This catalog contains tables and sequences known to the subscription after running: [`CREATE SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription), [`ALTER SUBSCRIPTION ... REFRESH PUBLICATION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-refresh-publication), or [`ALTER SUBSCRIPTION ... REFRESH SEQUENCES`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-refresh-sequences).


**Table: `pg_subscription_rel` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>srsubid</code> <code>oid</code> (references <a href="pg_subscription.md#catalog-pg-subscription"><code>pg_subscription</code></a>.<code>oid</code>)</p>
<p>Reference to subscription</p></td>
</tr>
<tr>
<td><p><code>srrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>Reference to table or sequence</p></td>
</tr>
<tr>
<td><p><code>srsubstate</code> <code>char</code></p>
<p>State code for the table or sequence.</p>
<p>State codes for tables: <code>i</code> = initialize, <code>d</code> = data is being copied, <code>f</code> = finished table copy, <code>s</code> = synchronized, <code>r</code> = ready (normal replication)</p>
<p>State codes for sequences: <code>i</code> = initialize, <code>r</code> = ready</p></td>
</tr>
<tr>
<td><p><code>srsublsn</code> <code>pg_lsn</code></p>
<p>Remote LSN of the state change used for synchronization coordination when in <code>s</code> or <code>r</code> states, otherwise null</p></td>
</tr>
</tbody>
</table>
