<a id="view-pg-replication-slots"></a>

## `pg_replication_slots`


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
<p>True if this slot is currently being streamed</p></td>
</tr>
<tr>
<td><p><code>active_pid</code> <code>int4</code></p>
<p>The process ID of the session streaming data for this slot. <code>NULL</code> if inactive.</p></td>
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
<td><p><code>two_phase_at</code> <code>pg_lsn</code></p>
<p>The address (<code>LSN</code>) from which the decoding of prepared transactions is enabled. <code>NULL</code> for logical slots where <code>two_phase</code> is false and for physical slots.</p></td>
</tr>
<tr>
<td><p><code>inactive_since</code> <code>timestamptz</code></p>
<p>The time when the slot became inactive. <code>NULL</code> if the slot is currently being streamed. If the slot becomes invalid, this value will never be updated. For standby slots that are being synced from a primary server (whose <code>synced</code> field is <code>true</code>), the <code>inactive_since</code> indicates the time when slot synchronization (see <a href="../../server-programming/logical-decoding/logical-decoding-concepts.md#logicaldecoding-replication-slots-synchronization">Replication Slot Synchronization</a>) was most recently stopped. <code>NULL</code> if the slot has always been synchronized. This helps standby slots track when synchronization was interrupted.</p></td>
</tr>
<tr>
<td><p><code>conflicting</code> <code>bool</code></p>
<p>True if this logical slot conflicted with recovery (and so is now invalidated). When this column is true, check <code>invalidation_reason</code> column for the conflict reason. Always <code>NULL</code> for physical slots.</p></td>
</tr>
<tr>
<td><p><code>invalidation_reason</code> <code>text</code></p>
<p>The reason for the slot's invalidation. It is set for both logical and physical slots. <code>NULL</code> if the slot is not invalidated. Possible values are:</p>
<p>-  <code>wal_removed</code> means that the required WAL has been removed. <br>
-  <code>rows_removed</code> means that the required rows have been removed. It is set only for logical slots. <br>
-  <code>wal_level_insufficient</code> means that the primary doesn't have an <a href="../../server-administration/server-configuration/preset-options.md#guc-effective-wal-level">effective_wal_level</a> sufficient to perform logical decoding. It is set only for logical slots. <br>
-  <code>idle_timeout</code> means that the slot has remained inactive longer than the configured <a href="../../server-administration/server-configuration/replication.md#guc-idle-replication-slot-timeout">idle_replication_slot_timeout</a> duration.</p></td>
</tr>
<tr>
<td><p><code>failover</code> <code>bool</code></p>
<p>True if this is a logical slot enabled to be synced to the standbys so that logical replication can be resumed from the new primary after failover. Always false for physical slots.</p></td>
</tr>
<tr>
<td><p><code>synced</code> <code>bool</code></p>
<p>True if this is a logical slot that was synced from a primary server. On a hot standby, the slots with the synced column marked as true can neither be used for logical decoding nor dropped manually. The value of this column has no meaning on the primary server; the column value on the primary is default false for all slots but may (if leftover from a promoted standby) also be true.</p></td>
</tr>
<tr>
<td><p><code>slotsync_skip_reason</code><code>text</code></p>
<p>The reason for the last slot synchronization skip. Slot synchronization occurs only on standby servers and thus this column has no meaning on the primary server. It is relevant mainly for logical slots on standby servers whose <code>synced</code> field is <code>true</code>. It is <code>NULL</code> if slot synchronization is successful. Possible values are:</p>
<p>-  <code>wal_or_rows_removed</code> means that the required WALs or catalog rows have already been removed or are at the risk of removal from the standby. <br>
-  <code>wal_not_flushed</code> means that the standby had not flushed the WAL corresponding to the position reserved on the failover slot. <br>
-  <code>no_consistent_snapshot</code> means that the standby could not build a consistent snapshot to decode WALs from <code>restart_lsn</code>. <br>
-  <code>slot_invalidated</code> means that the synced slot is invalidated.</p></td>
</tr>
</tbody>
</table>
