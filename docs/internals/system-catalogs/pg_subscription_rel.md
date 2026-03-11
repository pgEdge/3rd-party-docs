<a id="catalog-pg-subscription-rel"></a>

## `pg_subscription_rel`


 The catalog `pg_subscription_rel` contains the state for each replicated relation in each subscription. This is a many-to-many mapping.


 This catalog only contains tables known to the subscription after running either [`CREATE SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription) or [`ALTER SUBSCRIPTION ... REFRESH PUBLICATION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription).


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
<p>Reference to relation</p></td>
</tr>
<tr>
<td><p><code>srsubstate</code> <code>char</code></p>
<p>State code: <code>i</code> = initialize, <code>d</code> = data is being copied, <code>f</code> = finished table copy, <code>s</code> = synchronized, <code>r</code> = ready (normal replication)</p></td>
</tr>
<tr>
<td><p><code>srsublsn</code> <code>pg_lsn</code></p>
<p>Remote LSN of the state change used for synchronization coordination when in <code>s</code> or <code>r</code> states, otherwise null</p></td>
</tr>
</tbody>
</table>
