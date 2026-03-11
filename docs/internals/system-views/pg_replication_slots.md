## `pg_replication_slots` { #view-pg-replication-slots }


 The `pg_replication_slots` view provides a listing of all replication slots that currently exist on the database cluster, along with their current state.


 For more on replication slots, see [Replication Slots](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication-slots) and [Logical Decoding](../../server-programming/logical-decoding/index.md#logicaldecoding).


**Table: `pg_replication_slots` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>slot_name</code> <code>name</code></p>
<p>A unique, cluster-wide identifier for the replication slot</p></td>
</tr>
<tr>
<td><p><code>plugin</code> <code>name</code></p>
<p>The base name of the shared object containing the output plugin this logical slot is using, or null for physical slots.</p></td>
</tr>
<tr>
<td><p><code>slot_type</code> <code>text</code></p>
<p>The slot type: <code>physical</code> or <code>logical</code></p></td>
</tr>
<tr>
<td><p><code>datoid</code> <code>oid</code> (references <a href="../system-catalogs/pg_database.md#catalog-pg-database"><code>pg_database</code></a>.<code>oid</code>)</p>
<p>The OID of the database this slot is associated with, or null. Only logical slots have an associated database.</p></td>
</tr>
<tr>
<td><p><code>database</code> <code>name</code> (references <a href="../system-catalogs/pg_database.md#catalog-pg-database"><code>pg_database</code></a>.<code>datname</code>)</p>
<p>The name of the database this slot is associated with, or null. Only logical slots have an associated database.</p></td>
</tr>
<tr>
<td><p><code>temporary</code> <code>bool</code></p>
<p>True if this is a temporary replication slot. Temporary slots are not saved to disk and are automatically dropped on error or when the session has finished.</p></td>
</tr>
<tr>
<td><p><code>active</code> <code>bool</code></p>
<p>True if this slot is currently actively being used</p></td>
</tr>
<tr>
<td><p><code>active_pid</code> <code>int4</code></p>
<p>The process ID of the session using this slot if the slot is currently actively being used. <code>NULL</code> if inactive.</p></td>
</tr>
<tr>
<td><p><code>xmin</code> <code>xid</code></p>
<p>The oldest transaction that this slot needs the database to retain. <code>VACUUM</code> cannot remove tuples deleted by any later transaction.</p></td>
</tr>
<tr>
<td><p><code>catalog_xmin</code> <code>xid</code></p>
<p>The oldest transaction affecting the system catalogs that this slot needs the database to retain. <code>VACUUM</code> cannot remove catalog tuples deleted by any later transaction.</p></td>
</tr>
<tr>
<td><p><code>restart_lsn</code> <code>pg_lsn</code></p>
<p>The address (<code>LSN</code>) of oldest WAL which still might be required by the consumer of this slot and thus won't be automatically removed during checkpoints unless this LSN gets behind more than <a href="../../server-administration/server-configuration/replication.md#guc-max-slot-wal-keep-size">max_slot_wal_keep_size</a> from the current LSN. <code>NULL</code> if the <code>LSN</code> of this slot has never been reserved.</p></td>
</tr>
<tr>
<td><p><code>confirmed_flush_lsn</code> <code>pg_lsn</code></p>
<p>The address (<code>LSN</code>) up to which the logical slot's consumer has confirmed receiving data. Data corresponding to the transactions committed before this <code>LSN</code> is not available anymore. <code>NULL</code> for physical slots.</p></td>
</tr>
<tr>
<td><p><code>wal_status</code> <code>text</code></p>
<p>Availability of WAL files claimed by this slot. Possible values are:</p>
<p>- <code>reserved</code> means that the claimed files are within <code>max_wal_size</code>.<br>
- <code>extended</code> means that <code>max_wal_size</code> is exceeded but the files are still retained, either by the replication slot or by <code>wal_keep_size</code>. <br>
-  <code>unreserved</code> means that the slot no longer retains the required WAL files and some of them are to be removed at the next checkpoint. This typically occurs when <a href="../../server-administration/server-configuration/replication.md#guc-max-slot-wal-keep-size">max_slot_wal_keep_size</a> is set to a non-negative value. This state can return to <code>reserved</code> or <code>extended</code>. <br>
-  <code>lost</code> means that this slot is no longer usable.</p></td>
</tr>
<tr>
<td><p><code>safe_wal_size</code> <code>int8</code></p>
<p>The number of bytes that can be written to WAL such that this slot is not in danger of getting in state "lost". It is NULL for lost slots, as well as if <code>max_slot_wal_keep_size</code> is <code>-1</code>.</p></td>
</tr>
<tr>
<td><p><code>two_phase</code> <code>bool</code></p>
<p>True if the slot is enabled for decoding prepared transactions. Always false for physical slots.</p></td>
</tr>
<tr>
<td><p><code>conflicting</code> <code>bool</code></p>
<p>True if this logical slot conflicted with recovery (and so is now invalidated). Always NULL for physical slots.</p></td>
</tr>
</tbody>
</table>
