## Progress Reporting { #progress-reporting }


 PostgreSQL has the ability to report the progress of certain commands during command execution. Currently, the only commands which support progress reporting are `ANALYZE`, `COPY`, `CREATE INDEX`, `REPACK` (and its obsolete spelling `CLUSTER`), `VACUUM`, and [BASE_BACKUP](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication-base-backup) (i.e., replication command that [app-pgbasebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) issues to take a base backup). This may be expanded in the future.


### ANALYZE Progress Reporting { #analyze-progress-reporting }


 Whenever `ANALYZE` is running, the `pg_stat_progress_analyze` view will contain a row for each backend that is currently running that command. The tables below describe the information that will be reported and provide information about how to interpret it.


<a id="pg-stat-progress-analyze-view"></a>
**Table: `pg_stat_progress_analyze` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of backend.</p></td>
</tr>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table being analyzed.</p></td>
</tr>
<tr>
<td><p><code>phase</code> <code>text</code></p>
<p>Current processing phase. See <a href="#analyze-phases">ANALYZE Phases</a>.</p></td>
</tr>
<tr>
<td><p><code>sample_blks_total</code> <code>bigint</code></p>
<p>Total number of heap blocks that will be sampled.</p></td>
</tr>
<tr>
<td><p><code>sample_blks_scanned</code> <code>bigint</code></p>
<p>Number of heap blocks scanned.</p></td>
</tr>
<tr>
<td><p><code>ext_stats_total</code> <code>bigint</code></p>
<p>Number of extended statistics.</p></td>
</tr>
<tr>
<td><p><code>ext_stats_computed</code> <code>bigint</code></p>
<p>Number of extended statistics computed. This counter only advances when the phase is <code>computing extended statistics</code>.</p></td>
</tr>
<tr>
<td><p><code>child_tables_total</code> <code>bigint</code></p>
<p>Number of child tables.</p></td>
</tr>
<tr>
<td><p><code>child_tables_done</code> <code>bigint</code></p>
<p>Number of child tables scanned. This counter only advances when the phase is <code>acquiring inherited sample rows</code>.</p></td>
</tr>
<tr>
<td><p><code>current_child_table_relid</code> <code>oid</code></p>
<p>OID of the child table currently being scanned. This field is only valid when the phase is <code>acquiring inherited sample rows</code>.</p></td>
</tr>
<tr>
<td><p><code>delay_time</code> <code>double precision</code></p>
<p>Total time spent sleeping due to cost-based delay (see <a href="../server-configuration/vacuuming.md#runtime-config-resource-vacuum-cost">Cost-based Vacuum Delay</a>), in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-cost-delay-timing">track_cost_delay_timing</a> is enabled, otherwise zero).</p></td>
</tr>
<tr>
<td><p><code>started_by</code> <code>text</code></p>
<p>Shows what caused the current <code>ANALYZE</code> operation to be started. Possible values are:</p>
<p>-  <code>manual</code>: The analyze was started by an explicit <code>ANALYZE</code>, or by <code>VACUUM</code> with the <code>ANALYZE</code> option. <br>
-  <code>autovacuum</code>: The analyze was started by an autovacuum worker.</p></td>
</tr>
</tbody>
</table>


<a id="analyze-phases"></a>
**Table: ANALYZE Phases**

| Phase | Description |
| --- | --- |
| `initializing` | The command is preparing to begin scanning the heap. This phase is expected to be very brief. |
| `acquiring sample rows` | The command is currently scanning the table given by `relid` to obtain sample rows. |
| `acquiring inherited sample rows` | The command is currently scanning child tables to obtain sample rows. Columns `child_tables_total`, `child_tables_done`, and `current_child_table_relid` contain the progress information for this phase. |
| `computing statistics` | The command is computing statistics from the sample rows obtained during the table scan. |
| `computing extended statistics` | The command is computing extended statistics from the sample rows obtained during the table scan. |
| `finalizing analyze` | The command is updating `pg_class`. When this phase is completed, `ANALYZE` will end. |


!!! note

    Note that when `ANALYZE` is run on a partitioned table without the `ONLY` keyword, all of its partitions are also recursively analyzed. In that case, `ANALYZE` progress is reported first for the parent table, whereby its inheritance statistics are collected, followed by that for each partition.


### CLUSTER Progress Reporting { #cluster-progress-reporting }


 Whenever `CLUSTER` or `VACUUM FULL` is running, the `pg_stat_progress_cluster` view will contain a row for each backend that is currently running either command. The tables below describe the information that will be reported and provide information about how to interpret it.


<a id="pg-stat-progress-cluster-view"></a>
**Table: `pg_stat_progress_cluster` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of backend.</p></td>
</tr>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table being clustered.</p></td>
</tr>
<tr>
<td><p><code>command</code> <code>text</code></p>
<p>The command that is running. Either <code>CLUSTER</code> or <code>VACUUM FULL</code>.</p></td>
</tr>
<tr>
<td><p><code>phase</code> <code>text</code></p>
<p>Current processing phase. See <a href="#cluster-phases">CLUSTER and VACUUM FULL Phases</a>.</p></td>
</tr>
<tr>
<td><p><code>cluster_index_relid</code> <code>oid</code></p>
<p>If the table is being scanned using an index, this is the OID of the index being used; otherwise, it is zero.</p></td>
</tr>
<tr>
<td><p><code>heap_tuples_scanned</code> <code>bigint</code></p>
<p>Number of heap tuples scanned. This counter only advances when the phase is <code>seq scanning heap</code>, <code>index scanning heap</code> or <code>writing new heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_tuples_written</code> <code>bigint</code></p>
<p>Number of heap tuples written. This counter only advances when the phase is <code>seq scanning heap</code>, <code>index scanning heap</code> or <code>writing new heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_total</code> <code>bigint</code></p>
<p>Total number of heap blocks in the table. This number is reported as of the beginning of <code>seq scanning heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_scanned</code> <code>bigint</code></p>
<p>Number of heap blocks scanned. This counter only advances when the phase is <code>seq scanning heap</code>.</p></td>
</tr>
<tr>
<td><p><code>index_rebuild_count</code> <code>bigint</code></p>
<p>Number of indexes rebuilt. This counter only advances when the phase is <code>rebuilding index</code>.</p></td>
</tr>
</tbody>
</table>


<a id="cluster-phases"></a>
**Table: CLUSTER and VACUUM FULL Phases**

| Phase | Description |
| --- | --- |
| `initializing` | The command is preparing to begin scanning the heap. This phase is expected to be very brief. |
| `seq scanning heap` | The command is currently scanning the table using a sequential scan. |
| `index scanning heap` | `CLUSTER` is currently scanning the table using an index scan. |
| `sorting tuples` | `CLUSTER` is currently sorting tuples. |
| `writing new heap` | `CLUSTER` is currently writing the new heap. |
| `swapping relation files` | The command is currently swapping newly-built files into place. |
| `rebuilding index` | The command is currently rebuilding an index. |
| `performing final cleanup` | The command is performing final cleanup. When this phase is completed, `CLUSTER` or `VACUUM FULL` will end. |


### COPY Progress Reporting { #copy-progress-reporting }


 Whenever `COPY` is running, the `pg_stat_progress_copy` view will contain one row for each backend that is currently running a `COPY` command. The table below describes the information that will be reported and provides information about how to interpret it.


<a id="pg-stat-progress-copy-view"></a>
**Table: `pg_stat_progress_copy` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of backend.</p></td>
</tr>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table on which the <code>COPY</code> command is executed. It is set to <code>0</code> if copying from a <code>SELECT</code> query.</p></td>
</tr>
<tr>
<td><p><code>command</code> <code>text</code></p>
<p>The command that is running: <code>COPY FROM</code>, or <code>COPY TO</code>.</p></td>
</tr>
<tr>
<td><p><code>type</code> <code>text</code></p>
<p>The I/O type that the data is read from or written to: <code>FILE</code>, <code>PROGRAM</code>, <code>PIPE</code> (for <code>COPY FROM STDIN</code> and <code>COPY TO STDOUT</code>), or <code>CALLBACK</code> (used for example during the initial table synchronization in logical replication).</p></td>
</tr>
<tr>
<td><p><code>bytes_processed</code> <code>bigint</code></p>
<p>Number of bytes already processed by <code>COPY</code> command.</p></td>
</tr>
<tr>
<td><p><code>bytes_total</code> <code>bigint</code></p>
<p>Size of source file for <code>COPY FROM</code> command in bytes. It is set to <code>0</code> if not available.</p></td>
</tr>
<tr>
<td><p><code>tuples_processed</code> <code>bigint</code></p>
<p>Number of tuples already processed by <code>COPY</code> command.</p></td>
</tr>
<tr>
<td><p><code>tuples_excluded</code> <code>bigint</code></p>
<p>Number of tuples not processed because they were excluded by the <code>WHERE</code> clause of the <code>COPY</code> command.</p></td>
</tr>
<tr>
<td><p><code>tuples_skipped</code> <code>bigint</code></p>
<p>Number of tuples skipped because they contain malformed data. This counter only advances when <code>ignore</code> is specified to the <code>ON_ERROR</code> option.</p></td>
</tr>
</tbody>
</table>


### CREATE INDEX Progress Reporting { #create-index-progress-reporting }


 Whenever `CREATE INDEX` or `REINDEX` is running, the `pg_stat_progress_create_index` view will contain one row for each backend that is currently creating indexes. The tables below describe the information that will be reported and provide information about how to interpret it.


<a id="pg-stat-progress-create-index-view"></a>
**Table: `pg_stat_progress_create_index` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of the backend creating indexes.</p></td>
</tr>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table on which the index is being created.</p></td>
</tr>
<tr>
<td><p><code>index_relid</code> <code>oid</code></p>
<p>OID of the index being created or reindexed. During a non-concurrent <code>CREATE INDEX</code>, this is 0.</p></td>
</tr>
<tr>
<td><p><code>command</code> <code>text</code></p>
<p>Specific command type: <code>CREATE INDEX</code>, <code>CREATE INDEX CONCURRENTLY</code>, <code>REINDEX</code>, or <code>REINDEX CONCURRENTLY</code>.</p></td>
</tr>
<tr>
<td><p><code>phase</code> <code>text</code></p>
<p>Current processing phase of index creation. See <a href="#create-index-phases">CREATE INDEX Phases</a>.</p></td>
</tr>
<tr>
<td><p><code>lockers_total</code> <code>bigint</code></p>
<p>Total number of lockers to wait for, when applicable.</p></td>
</tr>
<tr>
<td><p><code>lockers_done</code> <code>bigint</code></p>
<p>Number of lockers already waited for.</p></td>
</tr>
<tr>
<td><p><code>current_locker_pid</code> <code>bigint</code></p>
<p>Process ID of the locker currently being waited for.</p></td>
</tr>
<tr>
<td><p><code>blocks_total</code> <code>bigint</code></p>
<p>Total number of blocks to be processed in the current phase.</p></td>
</tr>
<tr>
<td><p><code>blocks_done</code> <code>bigint</code></p>
<p>Number of blocks already processed in the current phase.</p></td>
</tr>
<tr>
<td><p><code>tuples_total</code> <code>bigint</code></p>
<p>Total number of tuples to be processed in the current phase.</p></td>
</tr>
<tr>
<td><p><code>tuples_done</code> <code>bigint</code></p>
<p>Number of tuples already processed in the current phase.</p></td>
</tr>
<tr>
<td><p><code>partitions_total</code> <code>bigint</code></p>
<p>Total number of partitions on which the index is to be created or attached, including both direct and indirect partitions. <code>0</code> during a <code>REINDEX</code>, or when the index is not partitioned.</p></td>
</tr>
<tr>
<td><p><code>partitions_done</code> <code>bigint</code></p>
<p>Number of partitions on which the index has already been created or attached, including both direct and indirect partitions. <code>0</code> during a <code>REINDEX</code>, or when the index is not partitioned.</p></td>
</tr>
</tbody>
</table>


<a id="create-index-phases"></a>
**Table: CREATE INDEX Phases**

| Phase | Description |
| --- | --- |
| `initializing` | `CREATE INDEX` or `REINDEX` is preparing to create the index. This phase is expected to be very brief. |
| `waiting for writers before build` | `CREATE INDEX CONCURRENTLY` or `REINDEX CONCURRENTLY` is waiting for transactions with write locks that can potentially see the table to finish. This phase is skipped when not in concurrent mode. Columns `lockers_total`, `lockers_done` and `current_locker_pid` contain the progress information for this phase. |
| `building index` | The index is being built by the access method-specific code. In this phase, access methods that support progress reporting fill in their own progress data, and the subphase is indicated in this column. Typically, `blocks_total` and `blocks_done` will contain progress data, as well as potentially `tuples_total` and `tuples_done`. |
| `waiting for writers before validation` | `CREATE INDEX CONCURRENTLY` or `REINDEX CONCURRENTLY` is waiting for transactions with write locks that can potentially write into the table to finish. This phase is skipped when not in concurrent mode. Columns `lockers_total`, `lockers_done` and `current_locker_pid` contain the progress information for this phase. |
| `index validation: scanning index` | `CREATE INDEX CONCURRENTLY` is scanning the index searching for tuples that need to be validated. This phase is skipped when not in concurrent mode. Columns `blocks_total` (set to the total size of the index) and `blocks_done` contain the progress information for this phase. |
| `index validation: sorting tuples` | `CREATE INDEX CONCURRENTLY` is sorting the output of the index scanning phase. |
| `index validation: scanning table` | `CREATE INDEX CONCURRENTLY` is scanning the table to validate the index tuples collected in the previous two phases. This phase is skipped when not in concurrent mode. Columns `blocks_total` (set to the total size of the table) and `blocks_done` contain the progress information for this phase. |
| `waiting for old snapshots` | `CREATE INDEX CONCURRENTLY` or `REINDEX CONCURRENTLY` is waiting for transactions that can potentially see the table to release their snapshots. This phase is skipped when not in concurrent mode. Columns `lockers_total`, `lockers_done` and `current_locker_pid` contain the progress information for this phase. |
| `waiting for readers before marking dead` | `REINDEX CONCURRENTLY` is waiting for transactions with read locks on the table to finish, before marking the old index dead. This phase is skipped when not in concurrent mode. Columns `lockers_total`, `lockers_done` and `current_locker_pid` contain the progress information for this phase. |
| `waiting for readers before dropping` | `REINDEX CONCURRENTLY` is waiting for transactions with read locks on the table to finish, before dropping the old index. This phase is skipped when not in concurrent mode. Columns `lockers_total`, `lockers_done` and `current_locker_pid` contain the progress information for this phase. |


### REPACK Progress Reporting { #repack-progress-reporting }


 Whenever `REPACK` is running, the `pg_stat_progress_repack` view will contain a row for each backend that is currently running the command. The tables below describe the information that will be reported and provide information about how to interpret it.


<a id="pg-stat-progress-repack-view"></a>
**Table: `pg_stat_progress_repack` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of backend.</p></td>
</tr>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table being repacked.</p></td>
</tr>
<tr>
<td><p><code>phase</code> <code>text</code></p>
<p>Current processing phase. See <a href="#repack-phases">REPACK Phases</a>.</p></td>
</tr>
<tr>
<td><p><code>repack_index_relid</code> <code>oid</code></p>
<p>If the table is being scanned using an index, this is the OID of the index being used; otherwise, it is zero.</p></td>
</tr>
<tr>
<td><p><code>heap_tuples_scanned</code> <code>bigint</code></p>
<p>Number of heap tuples scanned. This counter only advances when the phase is <code>seq scanning heap</code>, <code>index scanning heap</code> or <code>writing new heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_tuples_written</code> <code>bigint</code></p>
<p>Number of heap tuples written. This counter only advances when the phase is <code>seq scanning heap</code>, <code>index scanning heap</code> or <code>writing new heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_total</code> <code>bigint</code></p>
<p>Total number of heap blocks in the table. This number is reported as of the beginning of <code>seq scanning heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_scanned</code> <code>bigint</code></p>
<p>Number of heap blocks scanned. This counter only advances when the phase is <code>seq scanning heap</code>.</p></td>
</tr>
<tr>
<td><p><code>index_rebuild_count</code> <code>bigint</code></p>
<p>Number of indexes rebuilt. This counter only advances when the phase is <code>rebuilding index</code>.</p></td>
</tr>
</tbody>
</table>


<a id="repack-phases"></a>
**Table: REPACK Phases**

| Phase | Description |
| --- | --- |
| `initializing` | The command is preparing to begin scanning the heap. This phase is expected to be very brief. |
| `seq scanning heap` | The command is currently scanning the table using a sequential scan. |
| `index scanning heap` | `REPACK` is currently scanning the table using an index scan. |
| `sorting tuples` | `REPACK` is currently sorting tuples. |
| `writing new heap` | `REPACK` is currently writing the new heap. |
| `swapping relation files` | The command is currently swapping newly-built files into place. |
| `rebuilding index` | The command is currently rebuilding an index. |
| `performing final cleanup` | The command is performing final cleanup. When this phase is completed, `REPACK` will end. |


### VACUUM Progress Reporting { #vacuum-progress-reporting }


 Whenever `VACUUM` is running, the `pg_stat_progress_vacuum` view will contain one row for each backend (including autovacuum worker processes) that is currently vacuuming. The tables below describe the information that will be reported and provide information about how to interpret it. Progress for `VACUUM FULL` commands is reported via `pg_stat_progress_cluster` because both `VACUUM FULL` and `CLUSTER` rewrite the table, while regular `VACUUM` only modifies it in place. See [CLUSTER Progress Reporting](#cluster-progress-reporting).


<a id="pg-stat-progress-vacuum-view"></a>
**Table: `pg_stat_progress_vacuum` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of backend.</p></td>
</tr>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database to which this backend is connected.</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table being vacuumed.</p></td>
</tr>
<tr>
<td><p><code>phase</code> <code>text</code></p>
<p>Current processing phase of vacuum. See <a href="#vacuum-phases">VACUUM Phases</a>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_total</code> <code>bigint</code></p>
<p>Total number of heap blocks in the table. This number is reported as of the beginning of the scan; blocks added later will not be (and need not be) visited by this <code>VACUUM</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_scanned</code> <code>bigint</code></p>
<p>Number of heap blocks scanned. Because the <a href="../../internals/database-physical-storage/visibility-map.md#storage-vm">visibility map</a> is used to optimize scans, some blocks will be skipped without inspection; skipped blocks are included in this total, so that this number will eventually become equal to <code>heap_blks_total</code> when the vacuum is complete. This counter only advances when the phase is <code>scanning heap</code>.</p></td>
</tr>
<tr>
<td><p><code>heap_blks_vacuumed</code> <code>bigint</code></p>
<p>Number of heap blocks vacuumed. Unless the table has no indexes, this counter only advances when the phase is <code>vacuuming heap</code>. Blocks that contain no dead tuples are skipped, so the counter may sometimes skip forward in large increments.</p></td>
</tr>
<tr>
<td><p><code>index_vacuum_count</code> <code>bigint</code></p>
<p>Number of completed index vacuum cycles.</p></td>
</tr>
<tr>
<td><p><code>max_dead_tuple_bytes</code> <code>bigint</code></p>
<p>Amount of dead tuple data that we can store before needing to perform an index vacuum cycle, based on <a href="../server-configuration/resource-consumption.md#guc-maintenance-work-mem">maintenance_work_mem</a>.</p></td>
</tr>
<tr>
<td><p><code>dead_tuple_bytes</code> <code>bigint</code></p>
<p>Amount of dead tuple data collected since the last index vacuum cycle.</p></td>
</tr>
<tr>
<td><p><code>num_dead_item_ids</code> <code>bigint</code></p>
<p>Number of dead item identifiers collected since the last index vacuum cycle.</p></td>
</tr>
<tr>
<td><p><code>indexes_total</code> <code>bigint</code></p>
<p>Total number of indexes that will be vacuumed or cleaned up. This number is reported at the beginning of the <code>vacuuming indexes</code> phase or the <code>cleaning up indexes</code> phase.</p></td>
</tr>
<tr>
<td><p><code>indexes_processed</code> <code>bigint</code></p>
<p>Number of indexes processed. This counter only advances when the phase is <code>vacuuming indexes</code> or <code>cleaning up indexes</code>.</p></td>
</tr>
<tr>
<td><p><code>delay_time</code> <code>double precision</code></p>
<p>Total time spent sleeping due to cost-based delay (see <a href="../server-configuration/vacuuming.md#runtime-config-resource-vacuum-cost">Cost-based Vacuum Delay</a>), in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-cost-delay-timing">track_cost_delay_timing</a> is enabled, otherwise zero). This includes the time that any associated parallel workers have slept. However, parallel workers report their sleep time no more frequently than once per second, so the reported value may be slightly stale.</p></td>
</tr>
<tr>
<td><p><code>mode</code> <code>text</code></p>
<p>The mode in which the current <code>VACUUM</code> operation is running. See <a href="../routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound">Preventing Transaction ID Wraparound Failures</a> for details of each mode. Possible values are:</p>
<p>-  <code>normal</code>: The operation is performing a standard vacuum. It is neither required to run in aggressive mode nor operating in failsafe mode. <br>
-  <code>aggressive</code>: The operation is running an aggressive vacuum, which must scan every page that is not marked all-frozen. The parameters <a href="../server-configuration/vacuuming.md#guc-vacuum-freeze-table-age">vacuum_freeze_table_age</a> and <a href="../server-configuration/vacuuming.md#guc-vacuum-multixact-freeze-table-age">vacuum_multixact_freeze_table_age</a> determine when a table requires aggressive vacuuming. <br>
-  <code>failsafe</code>: The vacuum has entered failsafe mode, in which it performs only the minimum work necessary to avoid transaction ID or multixact ID wraparound failure. The parameters <a href="../server-configuration/vacuuming.md#guc-vacuum-failsafe-age">vacuum_failsafe_age</a> and <a href="../server-configuration/vacuuming.md#guc-vacuum-multixact-failsafe-age">vacuum_multixact_failsafe_age</a> determine when the vacuum enters failsafe mode. The vacuum may start in this mode or switch to it while running; the value of the <code>mode</code> column may transition from another mode to <code>failsafe</code> during the operation.</p></td>
</tr>
<tr>
<td><p><code>started_by</code> <code>text</code></p>
<p>Shows what caused the current <code>VACUUM</code> operation to be started. Possible values are:</p>
<p>-  <code>manual</code>: The vacuum was started by an explicit <code>VACUUM</code> command. <br>
-  <code>autovacuum</code>: The vacuum was started by an autovacuum worker. Vacuums run by autovacuum workers may be interrupted due to lock conflicts. <br>
-  <code>autovacuum_wraparound</code>: The vacuum was started by an autovacuum worker to prevent transaction ID or multixact ID wraparound. Vacuums run for wraparound protection are not interrupted due to lock conflicts.</p></td>
</tr>
</tbody>
</table>


<a id="vacuum-phases"></a>
**Table: VACUUM Phases**

| Phase | Description |
| --- | --- |
| `initializing` | `VACUUM` is preparing to begin scanning the heap. This phase is expected to be very brief. |
| `scanning heap` | `VACUUM` is currently scanning the heap. It will prune and defragment each page if required, and possibly perform freezing activity. The `heap_blks_scanned` column can be used to monitor the progress of the scan. |
| `vacuuming indexes` | `VACUUM` is currently vacuuming the indexes. If a table has any indexes, this will happen at least once per vacuum, after the heap has been completely scanned. It may happen multiple times per vacuum if [maintenance_work_mem](../server-configuration/resource-consumption.md#guc-maintenance-work-mem) (or, in the case of autovacuum, [autovacuum_work_mem](../server-configuration/resource-consumption.md#guc-autovacuum-work-mem) if set) is insufficient to store the number of dead tuples found. |
| `vacuuming heap` | `VACUUM` is currently vacuuming the heap. Vacuuming the heap is distinct from scanning the heap, and occurs after each instance of vacuuming indexes. If `heap_blks_scanned` is less than `heap_blks_total`, the system will return to scanning the heap after this phase is completed; otherwise, it will begin cleaning up indexes after this phase is completed. |
| `cleaning up indexes` | `VACUUM` is currently cleaning up indexes. This occurs after the heap has been completely scanned and all vacuuming of the indexes and the heap has been completed. |
| `truncating heap` | `VACUUM` is currently truncating the heap so as to return empty pages at the end of the relation to the operating system. This occurs after cleaning up indexes. |
| `performing final cleanup` | `VACUUM` is performing final cleanup. During this phase, `VACUUM` will vacuum the free space map, update statistics in `pg_class`, and report statistics to the cumulative statistics system. When this phase is completed, `VACUUM` will end. |


### Base Backup Progress Reporting { #basebackup-progress-reporting }


 Whenever an application like pg_basebackup is taking a base backup, the `pg_stat_progress_basebackup` view will contain a row for each WAL sender process that is currently running the `BASE_BACKUP` replication command and streaming the backup. The tables below describe the information that will be reported and provide information about how to interpret it.


<a id="pg-stat-progress-basebackup-view"></a>
**Table: `pg_stat_progress_basebackup` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of a WAL sender process.</p></td>
</tr>
<tr>
<td><p><code>phase</code> <code>text</code></p>
<p>Current processing phase. See <a href="#basebackup-phases">Base Backup Phases</a>.</p></td>
</tr>
<tr>
<td><p><code>backup_total</code> <code>bigint</code></p>
<p>Total amount of data that will be streamed. This is estimated and reported as of the beginning of <code>streaming database files</code> phase. Note that this is only an approximation since the database may change during <code>streaming database files</code> phase and WAL log may be included in the backup later. This is always the same value as <code>backup_streamed</code> once the amount of data streamed exceeds the estimated total size. If the estimation is disabled in pg_basebackup (i.e., <code>--no-estimate-size</code> option is specified), this is <code>NULL</code>.</p></td>
</tr>
<tr>
<td><p><code>backup_streamed</code> <code>bigint</code></p>
<p>Amount of data streamed. This counter only advances when the phase is <code>streaming database files</code> or <code>transferring wal files</code>.</p></td>
</tr>
<tr>
<td><p><code>tablespaces_total</code> <code>bigint</code></p>
<p>Total number of tablespaces that will be streamed.</p></td>
</tr>
<tr>
<td><p><code>tablespaces_streamed</code> <code>bigint</code></p>
<p>Number of tablespaces streamed. This counter only advances when the phase is <code>streaming database files</code>.</p></td>
</tr>
<tr>
<td><p><code>backup_type</code> <code>text</code></p>
<p>Backup type. Either <code>full</code> or <code>incremental</code>.</p></td>
</tr>
</tbody>
</table>


<a id="basebackup-phases"></a>
**Table: Base Backup Phases**

| Phase | Description |
| --- | --- |
| `initializing` | The WAL sender process is preparing to begin the backup. This phase is expected to be very brief. |
| `waiting for checkpoint to finish` | The WAL sender process is currently performing `pg_backup_start` to prepare to take a base backup, and waiting for the start-of-backup checkpoint to finish. |
| `estimating backup size` | The WAL sender process is currently estimating the total amount of database files that will be streamed as a base backup. |
| `streaming database files` | The WAL sender process is currently streaming database files as a base backup. |
| `waiting for wal archiving to finish` | The WAL sender process is currently performing `pg_backup_stop` to finish the backup, and waiting for all the WAL files required for the base backup to be successfully archived. If either `--wal-method=none` or `--wal-method=stream` is specified in pg_basebackup, the backup will end when this phase is completed. |
| `transferring wal files` | The WAL sender process is currently transferring all WAL logs generated during the backup. This phase occurs after `waiting for wal archiving to finish` phase if `--wal-method=fetch` is specified in pg_basebackup. The backup will end when this phase is completed. |
