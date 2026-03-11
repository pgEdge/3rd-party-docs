<a id="monitoring-stats"></a>

## The Cumulative Statistics System


 PostgreSQL's *cumulative statistics system* supports collection and reporting of information about server activity. Presently, accesses to tables and indexes in both disk-block and individual-row terms are counted. The total number of rows in each table, and information about vacuum and analyze actions for each table are also counted. If enabled, calls to user-defined functions and the total time spent in each one are counted as well.


 PostgreSQL also supports reporting dynamic information about exactly what is going on in the system right now, such as the exact command currently being executed by other server processes, and which other connections exist in the system. This facility is independent of the cumulative statistics system.
 <a id="monitoring-stats-setup"></a>

### Statistics Collection Configuration


 Since collection of statistics adds some overhead to query execution, the system can be configured to collect or not collect information. This is controlled by configuration parameters that are normally set in `postgresql.conf`. (See [Server Configuration](../server-configuration/index.md#runtime-config) for details about setting configuration parameters.)


 The parameter [track_activities](../server-configuration/run-time-statistics.md#guc-track-activities) enables monitoring of the current command being executed by any server process.


 The parameter [track_cost_delay_timing](../server-configuration/run-time-statistics.md#guc-track-cost-delay-timing) enables monitoring of cost-based vacuum delay.


 The parameter [track_counts](../server-configuration/run-time-statistics.md#guc-track-counts) controls whether cumulative statistics are collected about table and index accesses.


 The parameter [track_functions](../server-configuration/run-time-statistics.md#guc-track-functions) enables tracking of usage of user-defined functions.


 The parameter [track_io_timing](../server-configuration/run-time-statistics.md#guc-track-io-timing) enables monitoring of block read, write, extend, and fsync times.


 The parameter [track_wal_io_timing](../server-configuration/run-time-statistics.md#guc-track-wal-io-timing) enables monitoring of WAL read, write and fsync times.


 Normally these parameters are set in `postgresql.conf` so that they apply to all server processes, but it is possible to turn them on or off in individual sessions using the [sql-set](../../reference/sql-commands/set.md#sql-set) command. (To prevent ordinary users from hiding their activity from the administrator, only superusers are allowed to change these parameters with `SET`.)


 Cumulative statistics are collected in shared memory. Every PostgreSQL process collects statistics locally, then updates the shared data at appropriate intervals. When a server, including a physical replica, shuts down cleanly, a permanent copy of the statistics data is stored in the `pg_stat` subdirectory, so that statistics can be retained across server restarts. In contrast, when starting from an unclean shutdown (e.g., after an immediate shutdown, a server crash, starting from a base backup, and point-in-time recovery), all statistics counters are reset.
  <a id="monitoring-stats-views"></a>

### Viewing Statistics


 Several predefined views, listed in [Dynamic Statistics Views](#monitoring-stats-dynamic-views-table), are available to show the current state of the system. There are also several other views, listed in [Collected Statistics Views](#monitoring-stats-views-table), available to show the accumulated statistics. Alternatively, one can build custom views using the underlying cumulative statistics functions, as discussed in [Statistics Functions](#monitoring-stats-functions).


 When using the cumulative statistics views and functions to monitor collected data, it is important to realize that the information does not update instantaneously. Each individual server process flushes out accumulated statistics to shared memory just before going idle, but not more frequently than once per `PGSTAT_MIN_INTERVAL` milliseconds (1 second unless altered while building the server); so a query or transaction still in progress does not affect the displayed totals and the displayed information lags behind actual activity. However, current-query information collected by `track_activities` is always up-to-date.


 Another important point is that when a server process is asked to display any of the accumulated statistics, accessed values are cached until the end of its current transaction in the default configuration. So the statistics will show static information as long as you continue the current transaction. Similarly, information about the current queries of all sessions is collected when any such information is first requested within a transaction, and the same information will be displayed throughout the transaction. This is a feature, not a bug, because it allows you to perform several queries on the statistics and correlate the results without worrying that the numbers are changing underneath you. When analyzing statistics interactively, or with expensive queries, the time delta between accesses to individual statistics can lead to significant skew in the cached statistics. To minimize skew, `stats_fetch_consistency` can be set to `snapshot`, at the price of increased memory usage for caching not-needed statistics data. Conversely, if it's known that statistics are only accessed once, caching accessed statistics is unnecessary and can be avoided by setting `stats_fetch_consistency` to `none`. You can invoke `pg_stat_clear_snapshot()` to discard the current transaction's statistics snapshot or cached values (if any). The next use of statistical information will (when in snapshot mode) cause a new snapshot to be built or (when in cache mode) accessed statistics to be cached.


 A transaction can also see its own statistics (not yet flushed out to the shared memory statistics) in the views `pg_stat_xact_all_tables`, `pg_stat_xact_sys_tables`, `pg_stat_xact_user_tables`, and `pg_stat_xact_user_functions`. These numbers do not act as stated above; instead they update continuously throughout the transaction.


 Some of the information in the dynamic statistics views shown in [Dynamic Statistics Views](#monitoring-stats-dynamic-views-table) is security restricted. Ordinary users can only see all the information about their own sessions (sessions belonging to a role that they are a member of). In rows about other sessions, many columns will be null. Note, however, that the existence of a session and its general properties such as its sessions user and database are visible to all users. Superusers and roles with privileges of built-in role [`pg_read_all_stats`](../database-roles/predefined-roles.md#predefined-role-pg-monitor) can see all the information about all sessions.
 <a id="monitoring-stats-dynamic-views-table"></a>

**Table: Dynamic Statistics Views**

| View Name | Description |
| --- | --- |
| `pg_stat_activity` | One row per server process, showing information related to the current activity of that process, such as state and current query. See [`pg_stat_activity`](#monitoring-pg-stat-activity-view) for details. |
| `pg_stat_replication` | One row per WAL sender process, showing statistics about replication to that sender's connected standby server. See [`pg_stat_replication`](#monitoring-pg-stat-replication-view) for details. |
| `pg_stat_wal_receiver` | Only one row, showing statistics about the WAL receiver from that receiver's connected server. See [`pg_stat_wal_receiver`](#monitoring-pg-stat-wal-receiver-view) for details. |
| `pg_stat_recovery_prefetch` | Only one row, showing statistics about blocks prefetched during recovery. See [`pg_stat_recovery_prefetch`](#monitoring-pg-stat-recovery-prefetch) for details. |
| `pg_stat_subscription` | At least one row per subscription, showing information about the subscription workers. See [`pg_stat_subscription`](#monitoring-pg-stat-subscription) for details. |
| `pg_stat_ssl` | One row per connection (regular and replication), showing information about SSL used on this connection. See [`pg_stat_ssl`](#monitoring-pg-stat-ssl-view) for details. |
| `pg_stat_gssapi` | One row per connection (regular and replication), showing information about GSSAPI authentication and encryption used on this connection. See [`pg_stat_gssapi`](#monitoring-pg-stat-gssapi-view) for details. |
| `pg_stat_progress_analyze` | One row for each backend (including autovacuum worker processes) running `ANALYZE`, showing current progress. See [ANALYZE Progress Reporting](progress-reporting.md#analyze-progress-reporting). |
| `pg_stat_progress_create_index` | One row for each backend running `CREATE INDEX` or `REINDEX`, showing current progress. See [CREATE INDEX Progress Reporting](progress-reporting.md#create-index-progress-reporting). |
| `pg_stat_progress_vacuum` | One row for each backend (including autovacuum worker processes) running `VACUUM`, showing current progress. See [VACUUM Progress Reporting](progress-reporting.md#vacuum-progress-reporting). |
| `pg_stat_progress_cluster` | One row for each backend running `CLUSTER` or `VACUUM FULL`, showing current progress. See [CLUSTER Progress Reporting](progress-reporting.md#cluster-progress-reporting). |
| `pg_stat_progress_basebackup` | One row for each WAL sender process streaming a base backup, showing current progress. See [Base Backup Progress Reporting](progress-reporting.md#basebackup-progress-reporting). |
| `pg_stat_progress_copy` | One row for each backend running `COPY`, showing current progress. See [COPY Progress Reporting](progress-reporting.md#copy-progress-reporting). |
 <a id="monitoring-stats-views-table"></a>

**Table: Collected Statistics Views**

| View Name | Description |
| --- | --- |
| `pg_stat_archiver` | One row only, showing statistics about the WAL archiver process's activity. See [`pg_stat_archiver`](#monitoring-pg-stat-archiver-view) for details. |
| `pg_stat_bgwriter` | One row only, showing statistics about the background writer process's activity. See [`pg_stat_bgwriter`](#monitoring-pg-stat-bgwriter-view) for details. |
| `pg_stat_checkpointer` | One row only, showing statistics about the checkpointer process's activity. See [`pg_stat_checkpointer`](#monitoring-pg-stat-checkpointer-view) for details. |
| `pg_stat_database` | One row per database, showing database-wide statistics. See [`pg_stat_database`](#monitoring-pg-stat-database-view) for details. |
| `pg_stat_database_conflicts` | One row per database, showing database-wide statistics about query cancels due to conflict with recovery on standby servers. See [`pg_stat_database_conflicts`](#monitoring-pg-stat-database-conflicts-view) for details. |
| `pg_stat_io` | One row for each combination of backend type, context, and target object containing cluster-wide I/O statistics. See [`pg_stat_io`](#monitoring-pg-stat-io-view) for details. |
| `pg_stat_replication_slots` | One row per replication slot, showing statistics about the replication slot's usage. See [`pg_stat_replication_slots`](#monitoring-pg-stat-replication-slots-view) for details. |
| `pg_stat_slru` | One row per SLRU, showing statistics of operations. See [`pg_stat_slru`](#monitoring-pg-stat-slru-view) for details. |
| `pg_stat_subscription_stats` | One row per subscription, showing statistics about errors and conflicts. See [`pg_stat_subscription_stats`](#monitoring-pg-stat-subscription-stats) for details. |
| `pg_stat_wal` | One row only, showing statistics about WAL activity. See [`pg_stat_wal`](#monitoring-pg-stat-wal-view) for details. |
| `pg_stat_all_tables` | One row for each table in the current database, showing statistics about accesses to that specific table. See [`pg_stat_all_tables`](#monitoring-pg-stat-all-tables-view) for details. |
| `pg_stat_sys_tables` | Same as `pg_stat_all_tables`, except that only system tables are shown. |
| `pg_stat_user_tables` | Same as `pg_stat_all_tables`, except that only user tables are shown. |
| `pg_stat_xact_all_tables` | Similar to `pg_stat_all_tables`, but counts actions taken so far within the current transaction (which are *not* yet included in `pg_stat_all_tables` and related views). The columns for numbers of live and dead rows and vacuum and analyze actions are not present in this view. |
| `pg_stat_xact_sys_tables` | Same as `pg_stat_xact_all_tables`, except that only system tables are shown. |
| `pg_stat_xact_user_tables` | Same as `pg_stat_xact_all_tables`, except that only user tables are shown. |
| `pg_stat_all_indexes` | One row for each index in the current database, showing statistics about accesses to that specific index. See [`pg_stat_all_indexes`](#monitoring-pg-stat-all-indexes-view) for details. |
| `pg_stat_sys_indexes` | Same as `pg_stat_all_indexes`, except that only indexes on system tables are shown. |
| `pg_stat_user_indexes` | Same as `pg_stat_all_indexes`, except that only indexes on user tables are shown. |
| `pg_stat_user_functions` | One row for each tracked function, showing statistics about executions of that function. See [`pg_stat_user_functions`](#monitoring-pg-stat-user-functions-view) for details. |
| `pg_stat_xact_user_functions` | Similar to `pg_stat_user_functions`, but counts only calls during the current transaction (which are *not* yet included in `pg_stat_user_functions`). |
| `pg_statio_all_tables` | One row for each table in the current database, showing statistics about I/O on that specific table. See [`pg_statio_all_tables`](#monitoring-pg-statio-all-tables-view) for details. |
| `pg_statio_sys_tables` | Same as `pg_statio_all_tables`, except that only system tables are shown. |
| `pg_statio_user_tables` | Same as `pg_statio_all_tables`, except that only user tables are shown. |
| `pg_statio_all_indexes` | One row for each index in the current database, showing statistics about I/O on that specific index. See [`pg_statio_all_indexes`](#monitoring-pg-statio-all-indexes-view) for details. |
| `pg_statio_sys_indexes` | Same as `pg_statio_all_indexes`, except that only indexes on system tables are shown. |
| `pg_statio_user_indexes` | Same as `pg_statio_all_indexes`, except that only indexes on user tables are shown. |
| `pg_statio_all_sequences` | One row for each sequence in the current database, showing statistics about I/O on that specific sequence. See [`pg_statio_all_sequences`](#monitoring-pg-statio-all-sequences-view) for details. |
| `pg_statio_sys_sequences` | Same as `pg_statio_all_sequences`, except that only system sequences are shown. (Presently, no system sequences are defined, so this view is always empty.) |
| `pg_statio_user_sequences` | Same as `pg_statio_all_sequences`, except that only user sequences are shown. |


 The per-index statistics are particularly useful to determine which indexes are being used and how effective they are.


 The `pg_stat_io` and `pg_statio_` set of views are useful for determining the effectiveness of the buffer cache. They can be used to calculate a cache hit ratio. Note that while PostgreSQL's I/O statistics capture most instances in which the kernel was invoked in order to perform I/O, they do not differentiate between data which had to be fetched from disk and that which already resided in the kernel page cache. Users are advised to use the PostgreSQL statistics views in combination with operating system utilities for a more complete picture of their database's I/O performance.
  <a id="monitoring-pg-stat-activity-view"></a>

### `pg_stat_activity`


 The `pg_stat_activity` view will have one row per server process, showing information related to the current activity of that process.
 <a id="pg-stat-activity-view"></a>

**Table: `pg_stat_activity` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of the database this backend is connected to</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of the database this backend is connected to</p></td>
</tr>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of this backend</p></td>
</tr>
<tr>
<td><p><code>leader_pid</code> <code>integer</code></p>
<p>Process ID of the parallel group leader if this process is a parallel query worker, or process ID of the leader apply worker if this process is a parallel apply worker. <code>NULL</code> indicates that this process is a parallel group leader or leader apply worker, or does not participate in any parallel operation.</p></td>
</tr>
<tr>
<td><p><code>usesysid</code> <code>oid</code></p>
<p>OID of the user logged into this backend</p></td>
</tr>
<tr>
<td><p><code>usename</code> <code>name</code></p>
<p>Name of the user logged into this backend</p></td>
</tr>
<tr>
<td><p><code>application_name</code> <code>text</code></p>
<p>Name of the application that is connected to this backend</p></td>
</tr>
<tr>
<td><p><code>client_addr</code> <code>inet</code></p>
<p>IP address of the client connected to this backend. If this field is null, it indicates either that the client is connected via a Unix socket on the server machine or that this is an internal process such as autovacuum.</p></td>
</tr>
<tr>
<td><p><code>client_hostname</code> <code>text</code></p>
<p>Host name of the connected client, as reported by a reverse DNS lookup of <code>client_addr</code>. This field will only be non-null for IP connections, and only when <a href="../server-configuration/error-reporting-and-logging.md#guc-log-hostname">log_hostname</a> is enabled.</p></td>
</tr>
<tr>
<td><p><code>client_port</code> <code>integer</code></p>
<p>TCP port number that the client is using for communication with this backend, or <code>-1</code> if a Unix socket is used. If this field is null, it indicates that this is an internal server process.</p></td>
</tr>
<tr>
<td><p><code>backend_start</code> <code>timestamp with time zone</code></p>
<p>Time when this process was started. For client backends, this is the time the client connected to the server.</p></td>
</tr>
<tr>
<td><p><code>xact_start</code> <code>timestamp with time zone</code></p>
<p>Time when this process' current transaction was started, or null if no transaction is active. If the current query is the first of its transaction, this column is equal to the <code>query_start</code> column.</p></td>
</tr>
<tr>
<td><p><code>query_start</code> <code>timestamp with time zone</code></p>
<p>Time when the currently active query was started, or if <code>state</code> is not <code>active</code>, when the last query was started</p></td>
</tr>
<tr>
<td><p><code>state_change</code> <code>timestamp with time zone</code></p>
<p>Time when the <code>state</code> was last changed</p></td>
</tr>
<tr>
<td><p><code>wait_event_type</code> <code>text</code></p>
<p>The type of event for which the backend is waiting, if any; otherwise NULL. See <a href="#wait-event-table">Wait Event Types</a>.</p></td>
</tr>
<tr>
<td><p><code>wait_event</code> <code>text</code></p>
<p>Wait event name if backend is currently waiting, otherwise NULL. See <a href="#wait-event-activity-table">Wait Events of Type <code>Activity</code></a> through <a href="#wait-event-timeout-table">Wait Events of Type <code>Timeout</code></a>.</p></td>
</tr>
<tr>
<td><p><code>state</code> <code>text</code></p>
<p>Current overall state of this backend. Possible values are:</p>
<p>-  <code>starting</code>: The backend is in initial startup. Client authentication is performed during this phase. <br>
-  <code>active</code>: The backend is executing a query. <br>
-  <code>idle</code>: The backend is waiting for a new client command. <br>
-  <code>idle in transaction</code>: The backend is in a transaction, but is not currently executing a query. <br>
-  <code>idle in transaction (aborted)</code>: This state is similar to <code>idle in transaction</code>, except one of the statements in the transaction caused an error. <br>
-  <code>fastpath function call</code>: The backend is executing a fast-path function. <br>
-  <code>disabled</code>: This state is reported if <a href="../server-configuration/run-time-statistics.md#guc-track-activities">track_activities</a> is disabled in this backend.</p></td>
</tr>
<tr>
<td><p><code>backend_xid</code> <code>xid</code></p>
<p>Top-level transaction identifier of this backend, if any; see <a href="../../internals/transaction-processing/transactions-and-identifiers.md#transaction-id">Transactions and Identifiers</a>.</p></td>
</tr>
<tr>
<td><p><code>backend_xmin</code> <code>xid</code></p>
<p>The current backend's <code>xmin</code> horizon.</p></td>
</tr>
<tr>
<td><p><code>query_id</code> <code>bigint</code></p>
<p>Identifier of this backend's most recent query. If <code>state</code> is <code>active</code> this field shows the identifier of the currently executing query. In all other states, it shows the identifier of last query that was executed. Query identifiers are not computed by default so this field will be null unless <a href="../server-configuration/run-time-statistics.md#guc-compute-query-id">compute_query_id</a> parameter is enabled or a third-party module that computes query identifiers is configured.</p></td>
</tr>
<tr>
<td><p><code>query</code> <code>text</code></p>
<p>Text of this backend's most recent query. If <code>state</code> is <code>active</code> this field shows the currently executing query. In all other states, it shows the last query that was executed. By default the query text is truncated at 1024 bytes; this value can be changed via the parameter <a href="../server-configuration/run-time-statistics.md#guc-track-activity-query-size">track_activity_query_size</a>.</p></td>
</tr>
<tr>
<td><p><code>backend_type</code> <code>text</code></p>
<p>Type of current backend. Possible types are <code>autovacuum launcher</code>, <code>autovacuum worker</code>, <code>logical replication launcher</code>, <code>logical replication worker</code>, <code>parallel worker</code>, <code>background writer</code>, <code>client backend</code>, <code>checkpointer</code>, <code>archiver</code>, <code>standalone backend</code>, <code>startup</code>, <code>walreceiver</code>, <code>walsender</code>, <code>walwriter</code> and <code>walsummarizer</code>. In addition, background workers registered by extensions may have additional types.</p></td>
</tr>
</tbody>
</table>


!!! note

    The `wait_event` and `state` columns are independent. If a backend is in the `active` state, it may or may not be `waiting` on some event. If the state is `active` and `wait_event` is non-null, it means that a query is being executed, but is being blocked somewhere in the system. To keep the reporting overhead low, the system does not attempt to synchronize different aspects of activity data for a backend. As a result, ephemeral discrepancies may exist between the view's columns.
 <a id="wait-event-table"></a>

**Table: Wait Event Types**

| Wait Event Type | Description |
| --- | --- |
| `Activity` | The server process is idle. This event type indicates a process waiting for activity in its main processing loop. `wait_event` will identify the specific wait point; see [Wait Events of Type `Activity`](#wait-event-activity-table). |
| `BufferPin` | The server process is waiting for exclusive access to a data buffer. Buffer pin waits can be protracted if another process holds an open cursor that last read data from the buffer in question. See [Wait Events of Type `Bufferpin`](#wait-event-bufferpin-table). |
| `Client` | The server process is waiting for activity on a socket connected to a user application. Thus, the server expects something to happen that is independent of its internal processes. `wait_event` will identify the specific wait point; see [Wait Events of Type `Client`](#wait-event-client-table). |
| `Extension` | The server process is waiting for some condition defined by an extension module. See [Wait Events of Type `Extension`](#wait-event-extension-table). |
| `InjectionPoint` | The server process is waiting for an injection point to reach an outcome defined in a test. See [Injection Points](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-injection-points) for more details. This type has no predefined wait points. |
| `IO` | The server process is waiting for an I/O operation to complete. `wait_event` will identify the specific wait point; see [Wait Events of Type `Io`](#wait-event-io-table). |
| `IPC` | The server process is waiting for some interaction with another server process. `wait_event` will identify the specific wait point; see [Wait Events of Type `Ipc`](#wait-event-ipc-table). |
| `Lock` | The server process is waiting for a heavyweight lock. Heavyweight locks, also known as lock manager locks or simply locks, primarily protect SQL-visible objects such as tables. However, they are also used to ensure mutual exclusion for certain internal operations such as relation extension. `wait_event` will identify the type of lock awaited; see [Wait Events of Type `Lock`](#wait-event-lock-table). |
| `LWLock` | The server process is waiting for a lightweight lock. Most such locks protect a particular data structure in shared memory. `wait_event` will contain a name identifying the purpose of the lightweight lock. (Some locks have specific names; others are part of a group of locks each with a similar purpose.) See [Wait Events of Type `Lwlock`](#wait-event-lwlock-table). |
| `Timeout` | The server process is waiting for a timeout to expire. `wait_event` will identify the specific wait point; see [Wait Events of Type `Timeout`](#wait-event-timeout-table). |
 <a id="wait-event-activity-table"></a>

**Table: Wait Events of Type `Activity`**

| `Activity` Wait Event | Description |
| --- | --- |
| `ArchiverMain` | Waiting in main loop of archiver process. |
| `AutovacuumMain` | Waiting in main loop of autovacuum launcher process. |
| `BgwriterHibernate` | Waiting in background writer process, hibernating. |
| `BgwriterMain` | Waiting in main loop of background writer process. |
| `CheckpointerMain` | Waiting in main loop of checkpointer process. |
| `CheckpointerShutdown` | Waiting for checkpointer process to be terminated. |
| `IoWorkerMain` | Waiting in main loop of IO Worker process. |
| `LogicalApplyMain` | Waiting in main loop of logical replication apply process. |
| `LogicalLauncherMain` | Waiting in main loop of logical replication launcher process. |
| `LogicalParallelApplyMain` | Waiting in main loop of logical replication parallel apply process. |
| `RecoveryWalStream` | Waiting in main loop of startup process for WAL to arrive, during streaming recovery. |
| `ReplicationSlotsyncMain` | Waiting in main loop of slot sync worker. |
| `ReplicationSlotsyncShutdown` | Waiting for slot sync worker to shut down. |
| `SysloggerMain` | Waiting in main loop of syslogger process. |
| `WalReceiverMain` | Waiting in main loop of WAL receiver process. |
| `WalSenderMain` | Waiting in main loop of WAL sender process. |
| `WalSummarizerWal` | Waiting in WAL summarizer for more WAL to be generated. |
| `WalWriterMain` | Waiting in main loop of WAL writer process. |
 <a id="wait-event-bufferpin-table"></a>

**Table: Wait Events of Type `Bufferpin`**

| `BufferPin` Wait Event | Description |
| --- | --- |
| `BufferPin` | Waiting to acquire an exclusive pin on a buffer. |
 <a id="wait-event-client-table"></a>

**Table: Wait Events of Type `Client`**

| `Client` Wait Event | Description |
| --- | --- |
| `ClientRead` | Waiting to read data from the client. |
| `ClientWrite` | Waiting to write data to the client. |
| `GssOpenServer` | Waiting to read data from the client while establishing a GSSAPI session. |
| `LibpqwalreceiverConnect` | Waiting in WAL receiver to establish connection to remote server. |
| `LibpqwalreceiverReceive` | Waiting in WAL receiver to receive data from remote server. |
| `SslOpenServer` | Waiting for SSL while attempting connection. |
| `WaitForStandbyConfirmation` | Waiting for WAL to be received and flushed by the physical standby. |
| `WalSenderWaitForWal` | Waiting for WAL to be flushed in WAL sender process. |
| `WalSenderWriteData` | Waiting for any activity when processing replies from WAL receiver in WAL sender process. |
 <a id="wait-event-extension-table"></a>

**Table: Wait Events of Type `Extension`**

| `Extension` Wait Event | Description |
| --- | --- |
| `Extension` | Waiting in an extension. |
 <a id="wait-event-io-table"></a>

**Table: Wait Events of Type `Io`**

| `IO` Wait Event | Description |
| --- | --- |
| `AioIoCompletion` | Waiting for another process to complete IO. |
| `AioIoUringExecution` | Waiting for IO execution via io_uring. |
| `AioIoUringSubmit` | Waiting for IO submission via io_uring. |
| `BasebackupRead` | Waiting for base backup to read from a file. |
| `BasebackupSync` | Waiting for data written by a base backup to reach durable storage. |
| `BasebackupWrite` | Waiting for base backup to write to a file. |
| `BuffileRead` | Waiting for a read from a buffered file. |
| `BuffileTruncate` | Waiting for a buffered file to be truncated. |
| `BuffileWrite` | Waiting for a write to a buffered file. |
| `ControlFileRead` | Waiting for a read from the `pg_control` file. |
| `ControlFileSync` | Waiting for the `pg_control` file to reach durable storage. |
| `ControlFileSyncUpdate` | Waiting for an update to the `pg_control` file to reach durable storage. |
| `ControlFileWrite` | Waiting for a write to the `pg_control` file. |
| `ControlFileWriteUpdate` | Waiting for a write to update the `pg_control` file. |
| `CopyFileCopy` | Waiting for a file copy operation. |
| `CopyFileRead` | Waiting for a read during a file copy operation. |
| `CopyFileWrite` | Waiting for a write during a file copy operation. |
| `DataFileExtend` | Waiting for a relation data file to be extended. |
| `DataFileFlush` | Waiting for a relation data file to reach durable storage. |
| `DataFileImmediateSync` | Waiting for an immediate synchronization of a relation data file to durable storage. |
| `DataFilePrefetch` | Waiting for an asynchronous prefetch from a relation data file. |
| `DataFileRead` | Waiting for a read from a relation data file. |
| `DataFileSync` | Waiting for changes to a relation data file to reach durable storage. |
| `DataFileTruncate` | Waiting for a relation data file to be truncated. |
| `DataFileWrite` | Waiting for a write to a relation data file. |
| `DsmAllocate` | Waiting for a dynamic shared memory segment to be allocated. |
| `DsmFillZeroWrite` | Waiting to fill a dynamic shared memory backing file with zeroes. |
| `LockFileAddtodatadirRead` | Waiting for a read while adding a line to the data directory lock file. |
| `LockFileAddtodatadirSync` | Waiting for data to reach durable storage while adding a line to the data directory lock file. |
| `LockFileAddtodatadirWrite` | Waiting for a write while adding a line to the data directory lock file. |
| `LockFileCreateRead` | Waiting to read while creating the data directory lock file. |
| `LockFileCreateSync` | Waiting for data to reach durable storage while creating the data directory lock file. |
| `LockFileCreateWrite` | Waiting for a write while creating the data directory lock file. |
| `LockFileRecheckdatadirRead` | Waiting for a read during recheck of the data directory lock file. |
| `LogicalRewriteCheckpointSync` | Waiting for logical rewrite mappings to reach durable storage during a checkpoint. |
| `LogicalRewriteMappingSync` | Waiting for mapping data to reach durable storage during a logical rewrite. |
| `LogicalRewriteMappingWrite` | Waiting for a write of mapping data during a logical rewrite. |
| `LogicalRewriteSync` | Waiting for logical rewrite mappings to reach durable storage. |
| `LogicalRewriteTruncate` | Waiting for truncate of mapping data during a logical rewrite. |
| `LogicalRewriteWrite` | Waiting for a write of logical rewrite mappings. |
| `RelationMapRead` | Waiting for a read of the relation map file. |
| `RelationMapReplace` | Waiting for durable replacement of a relation map file. |
| `RelationMapWrite` | Waiting for a write to the relation map file. |
| `ReorderBufferRead` | Waiting for a read during reorder buffer management. |
| `ReorderBufferWrite` | Waiting for a write during reorder buffer management. |
| `ReorderLogicalMappingRead` | Waiting for a read of a logical mapping during reorder buffer management. |
| `ReplicationSlotRead` | Waiting for a read from a replication slot control file. |
| `ReplicationSlotRestoreSync` | Waiting for a replication slot control file to reach durable storage while restoring it to memory. |
| `ReplicationSlotSync` | Waiting for a replication slot control file to reach durable storage. |
| `ReplicationSlotWrite` | Waiting for a write to a replication slot control file. |
| `SlruFlushSync` | Waiting for SLRU data to reach durable storage during a checkpoint or database shutdown. |
| `SlruRead` | Waiting for a read of an SLRU page. |
| `SlruSync` | Waiting for SLRU data to reach durable storage following a page write. |
| `SlruWrite` | Waiting for a write of an SLRU page. |
| `SnapbuildRead` | Waiting for a read of a serialized historical catalog snapshot. |
| `SnapbuildSync` | Waiting for a serialized historical catalog snapshot to reach durable storage. |
| `SnapbuildWrite` | Waiting for a write of a serialized historical catalog snapshot. |
| `TimelineHistoryFileSync` | Waiting for a timeline history file received via streaming replication to reach durable storage. |
| `TimelineHistoryFileWrite` | Waiting for a write of a timeline history file received via streaming replication. |
| `TimelineHistoryRead` | Waiting for a read of a timeline history file. |
| `TimelineHistorySync` | Waiting for a newly created timeline history file to reach durable storage. |
| `TimelineHistoryWrite` | Waiting for a write of a newly created timeline history file. |
| `TwophaseFileRead` | Waiting for a read of a two phase state file. |
| `TwophaseFileSync` | Waiting for a two phase state file to reach durable storage. |
| `TwophaseFileWrite` | Waiting for a write of a two phase state file. |
| `VersionFileSync` | Waiting for the version file to reach durable storage while creating a database. |
| `VersionFileWrite` | Waiting for the version file to be written while creating a database. |
| `WalsenderTimelineHistoryRead` | Waiting for a read from a timeline history file during a walsender timeline command. |
| `WalBootstrapSync` | Waiting for WAL to reach durable storage during bootstrapping. |
| `WalBootstrapWrite` | Waiting for a write of a WAL page during bootstrapping. |
| `WalCopyRead` | Waiting for a read when creating a new WAL segment by copying an existing one. |
| `WalCopySync` | Waiting for a new WAL segment created by copying an existing one to reach durable storage. |
| `WalCopyWrite` | Waiting for a write when creating a new WAL segment by copying an existing one. |
| `WalInitSync` | Waiting for a newly initialized WAL file to reach durable storage. |
| `WalInitWrite` | Waiting for a write while initializing a new WAL file. |
| `WalRead` | Waiting for a read from a WAL file. |
| `WalSummaryRead` | Waiting for a read from a WAL summary file. |
| `WalSummaryWrite` | Waiting for a write to a WAL summary file. |
| `WalSync` | Waiting for a WAL file to reach durable storage. |
| `WalSyncMethodAssign` | Waiting for data to reach durable storage while assigning a new WAL sync method. |
| `WalWrite` | Waiting for a write to a WAL file. |
 <a id="wait-event-ipc-table"></a>

**Table: Wait Events of Type `Ipc`**

| `IPC` Wait Event | Description |
| --- | --- |
| `AppendReady` | Waiting for subplan nodes of an `Append` plan node to be ready. |
| `ArchiveCleanupCommand` | Waiting for [archive_cleanup_command](../server-configuration/write-ahead-log.md#guc-archive-cleanup-command) to complete. |
| `ArchiveCommand` | Waiting for [archive_command](../server-configuration/write-ahead-log.md#guc-archive-command) to complete. |
| `BackendTermination` | Waiting for the termination of another backend. |
| `BackupWaitWalArchive` | Waiting for WAL files required for a backup to be successfully archived. |
| `BgworkerShutdown` | Waiting for background worker to shut down. |
| `BgworkerStartup` | Waiting for background worker to start up. |
| `BtreePage` | Waiting for the page number needed to continue a parallel B-tree scan to become available. |
| `BufferIo` | Waiting for buffer I/O to complete. |
| `CheckpointDelayComplete` | Waiting for a backend that blocks a checkpoint from completing. |
| `CheckpointDelayStart` | Waiting for a backend that blocks a checkpoint from starting. |
| `CheckpointDone` | Waiting for a checkpoint to complete. |
| `CheckpointStart` | Waiting for a checkpoint to start. |
| `ExecuteGather` | Waiting for activity from a child process while executing a `Gather` plan node. |
| `HashBatchAllocate` | Waiting for an elected Parallel Hash participant to allocate a hash table. |
| `HashBatchElect` | Waiting to elect a Parallel Hash participant to allocate a hash table. |
| `HashBatchLoad` | Waiting for other Parallel Hash participants to finish loading a hash table. |
| `HashBuildAllocate` | Waiting for an elected Parallel Hash participant to allocate the initial hash table. |
| `HashBuildElect` | Waiting to elect a Parallel Hash participant to allocate the initial hash table. |
| `HashBuildHashInner` | Waiting for other Parallel Hash participants to finish hashing the inner relation. |
| `HashBuildHashOuter` | Waiting for other Parallel Hash participants to finish partitioning the outer relation. |
| `HashGrowBatchesDecide` | Waiting to elect a Parallel Hash participant to decide on future batch growth. |
| `HashGrowBatchesElect` | Waiting to elect a Parallel Hash participant to allocate more batches. |
| `HashGrowBatchesFinish` | Waiting for an elected Parallel Hash participant to decide on future batch growth. |
| `HashGrowBatchesReallocate` | Waiting for an elected Parallel Hash participant to allocate more batches. |
| `HashGrowBatchesRepartition` | Waiting for other Parallel Hash participants to finish repartitioning. |
| `HashGrowBucketsElect` | Waiting to elect a Parallel Hash participant to allocate more buckets. |
| `HashGrowBucketsReallocate` | Waiting for an elected Parallel Hash participant to finish allocating more buckets. |
| `HashGrowBucketsReinsert` | Waiting for other Parallel Hash participants to finish inserting tuples into new buckets. |
| `LogicalApplySendData` | Waiting for a logical replication leader apply process to send data to a parallel apply process. |
| `LogicalParallelApplyStateChange` | Waiting for a logical replication parallel apply process to change state. |
| `LogicalSyncData` | Waiting for a logical replication remote server to send data for initial table synchronization. |
| `LogicalSyncStateChange` | Waiting for a logical replication remote server to change state. |
| `MessageQueueInternal` | Waiting for another process to be attached to a shared message queue. |
| `MessageQueuePutMessage` | Waiting to write a protocol message to a shared message queue. |
| `MessageQueueReceive` | Waiting to receive bytes from a shared message queue. |
| `MessageQueueSend` | Waiting to send bytes to a shared message queue. |
| `MultixactCreation` | Waiting for a multixact creation to complete. |
| `ParallelBitmapScan` | Waiting for parallel bitmap scan to become initialized. |
| `ParallelCreateIndexScan` | Waiting for parallel `CREATE INDEX` workers to finish heap scan. |
| `ParallelFinish` | Waiting for parallel workers to finish computing. |
| `ProcarrayGroupUpdate` | Waiting for the group leader to clear the transaction ID at transaction end. |
| `ProcSignalBarrier` | Waiting for a barrier event to be processed by all backends. |
| `Promote` | Waiting for standby promotion. |
| `RecoveryConflictSnapshot` | Waiting for recovery conflict resolution for a vacuum cleanup. |
| `RecoveryConflictTablespace` | Waiting for recovery conflict resolution for dropping a tablespace. |
| `RecoveryEndCommand` | Waiting for [recovery_end_command](../server-configuration/write-ahead-log.md#guc-recovery-end-command) to complete. |
| `RecoveryPause` | Waiting for recovery to be resumed. |
| `ReplicationOriginDrop` | Waiting for a replication origin to become inactive so it can be dropped. |
| `ReplicationSlotDrop` | Waiting for a replication slot to become inactive so it can be dropped. |
| `RestoreCommand` | Waiting for [restore_command](../server-configuration/write-ahead-log.md#guc-restore-command) to complete. |
| `SafeSnapshot` | Waiting to obtain a valid snapshot for a `READ ONLY DEFERRABLE` transaction. |
| `SyncRep` | Waiting for confirmation from a remote server during synchronous replication. |
| `WalReceiverExit` | Waiting for the WAL receiver to exit. |
| `WalReceiverWaitStart` | Waiting for startup process to send initial data for streaming replication. |
| `WalSummaryReady` | Waiting for a new WAL summary to be generated. |
| `XactGroupUpdate` | Waiting for the group leader to update transaction status at transaction end. |
 <a id="wait-event-lock-table"></a>

**Table: Wait Events of Type `Lock`**

| `Lock` Wait Event | Description |
| --- | --- |
| `advisory` | Waiting to acquire an advisory user lock. |
| `applytransaction` | Waiting to acquire a lock on a remote transaction being applied by a logical replication subscriber. |
| `extend` | Waiting to extend a relation. |
| `frozenid` | Waiting to update `pg_database`.`datfrozenxid` and `pg_database`.`datminmxid`. |
| `object` | Waiting to acquire a lock on a non-relation database object. |
| `page` | Waiting to acquire a lock on a page of a relation. |
| `relation` | Waiting to acquire a lock on a relation. |
| `spectoken` | Waiting to acquire a speculative insertion lock. |
| `transactionid` | Waiting for a transaction to finish. |
| `tuple` | Waiting to acquire a lock on a tuple. |
| `userlock` | Waiting to acquire a user lock. |
| `virtualxid` | Waiting to acquire a virtual transaction ID lock; see [Transactions and Identifiers](../../internals/transaction-processing/transactions-and-identifiers.md#transaction-id). |
 <a id="wait-event-lwlock-table"></a>

**Table: Wait Events of Type `Lwlock`**

| `LWLock` Wait Event | Description |
| --- | --- |
| `AddinShmemInit` | Waiting to manage an extension's space allocation in shared memory. |
| `AioUringCompletion` | Waiting for another process to complete IO via io_uring. |
| `AioWorkerSubmissionQueue` | Waiting to access AIO worker submission queue. |
| `AutoFile` | Waiting to update the `postgresql.auto.conf` file. |
| `Autovacuum` | Waiting to read or update the current state of autovacuum workers. |
| `AutovacuumSchedule` | Waiting to ensure that a table selected for autovacuum still needs vacuuming. |
| `BackgroundWorker` | Waiting to read or update background worker state. |
| `BtreeVacuum` | Waiting to read or update vacuum-related information for a B-tree index. |
| `BufferContent` | Waiting to access a data page in memory. |
| `BufferMapping` | Waiting to associate a data block with a buffer in the buffer pool. |
| `CheckpointerComm` | Waiting to manage fsync requests. |
| `CommitTs` | Waiting to read or update the last value set for a transaction commit timestamp. |
| `CommitTsBuffer` | Waiting for I/O on a commit timestamp SLRU buffer. |
| `CommitTsSLRU` | Waiting to access the commit timestamp SLRU cache. |
| `ControlFile` | Waiting to read or update the `pg_control` file or create a new WAL file. |
| `DSMRegistry` | Waiting to read or update the dynamic shared memory registry. |
| `DSMRegistryDSA` | Waiting to access dynamic shared memory registry's dynamic shared memory allocator. |
| `DSMRegistryHash` | Waiting to access dynamic shared memory registry's shared hash table. |
| `DynamicSharedMemoryControl` | Waiting to read or update dynamic shared memory allocation information. |
| `InjectionPoint` | Waiting to read or update information related to injection points. |
| `LockFastPath` | Waiting to read or update a process' fast-path lock information. |
| `LockManager` | Waiting to read or update information about “heavyweight” locks. |
| `LogicalRepLauncherDSA` | Waiting to access logical replication launcher's dynamic shared memory allocator. |
| `LogicalRepLauncherHash` | Waiting to access logical replication launcher's shared hash table. |
| `LogicalRepWorker` | Waiting to read or update the state of logical replication workers. |
| `MultiXactGen` | Waiting to read or update shared multixact state. |
| `MultiXactMemberBuffer` | Waiting for I/O on a multixact member SLRU buffer. |
| `MultiXactMemberSLRU` | Waiting to access the multixact member SLRU cache. |
| `MultiXactOffsetBuffer` | Waiting for I/O on a multixact offset SLRU buffer. |
| `MultiXactOffsetSLRU` | Waiting to access the multixact offset SLRU cache. |
| `MultiXactTruncation` | Waiting to read or truncate multixact information. |
| `NotifyBuffer` | Waiting for I/O on a `NOTIFY` message SLRU buffer. |
| `NotifyQueue` | Waiting to read or update `NOTIFY` messages. |
| `NotifyQueueTail` | Waiting to update limit on `NOTIFY` message storage. |
| `NotifySLRU` | Waiting to access the `NOTIFY` message SLRU cache. |
| `OidGen` | Waiting to allocate a new OID. |
| `ParallelAppend` | Waiting to choose the next subplan during Parallel Append plan execution. |
| `ParallelBtreeScan` | Waiting to synchronize workers during Parallel B-tree scan plan execution. |
| `ParallelHashJoin` | Waiting to synchronize workers during Parallel Hash Join plan execution. |
| `ParallelQueryDSA` | Waiting for parallel query dynamic shared memory allocation. |
| `ParallelVacuumDSA` | Waiting for parallel vacuum dynamic shared memory allocation. |
| `PerSessionDSA` | Waiting for parallel query dynamic shared memory allocation. |
| `PerSessionRecordType` | Waiting to access a parallel query's information about composite types. |
| `PerSessionRecordTypmod` | Waiting to access a parallel query's information about type modifiers that identify anonymous record types. |
| `PerXactPredicateList` | Waiting to access the list of predicate locks held by the current serializable transaction during a parallel query. |
| `PgStatsData` | Waiting for shared memory stats data access. |
| `PgStatsDSA` | Waiting for stats dynamic shared memory allocator access. |
| `PgStatsHash` | Waiting for stats shared memory hash table access. |
| `PredicateLockManager` | Waiting to access predicate lock information used by serializable transactions. |
| `ProcArray` | Waiting to access the shared per-process data structures (typically, to get a snapshot or report a session's transaction ID). |
| `RelationMapping` | Waiting to read or update a `pg_filenode.map` file (used to track the filenode assignments of certain system catalogs). |
| `RelCacheInit` | Waiting to read or update a `pg_internal.init` relation cache initialization file. |
| `ReplicationOrigin` | Waiting to create, drop or use a replication origin. |
| `ReplicationOriginState` | Waiting to read or update the progress of one replication origin. |
| `ReplicationSlotAllocation` | Waiting to allocate or free a replication slot. |
| `ReplicationSlotControl` | Waiting to read or update replication slot state. |
| `ReplicationSlotIO` | Waiting for I/O on a replication slot. |
| `SerialBuffer` | Waiting for I/O on a serializable transaction conflict SLRU buffer. |
| `SerialControl` | Waiting to read or update shared `pg_serial` state. |
| `SerializableFinishedList` | Waiting to access the list of finished serializable transactions. |
| `SerializablePredicateList` | Waiting to access the list of predicate locks held by serializable transactions. |
| `SerializableXactHash` | Waiting to read or update information about serializable transactions. |
| `SerialSLRU` | Waiting to access the serializable transaction conflict SLRU cache. |
| `SharedTidBitmap` | Waiting to access a shared TID bitmap during a parallel bitmap index scan. |
| `SharedTupleStore` | Waiting to access a shared tuple store during parallel query. |
| `ShmemIndex` | Waiting to find or allocate space in shared memory. |
| `SInvalRead` | Waiting to retrieve messages from the shared catalog invalidation queue. |
| `SInvalWrite` | Waiting to add a message to the shared catalog invalidation queue. |
| `SubtransBuffer` | Waiting for I/O on a sub-transaction SLRU buffer. |
| `SubtransSLRU` | Waiting to access the sub-transaction SLRU cache. |
| `SyncRep` | Waiting to read or update information about the state of synchronous replication. |
| `SyncScan` | Waiting to select the starting location of a synchronized table scan. |
| `TablespaceCreate` | Waiting to create or drop a tablespace. |
| `TwoPhaseState` | Waiting to read or update the state of prepared transactions. |
| `WaitEventCustom` | Waiting to read or update custom wait events information. |
| `WALBufMapping` | Waiting to replace a page in WAL buffers. |
| `WALInsert` | Waiting to insert WAL data into a memory buffer. |
| `WALSummarizer` | Waiting to read or update WAL summarization state. |
| `WALWrite` | Waiting for WAL buffers to be written to disk. |
| `WrapLimitsVacuum` | Waiting to update limits on transaction id and multixact consumption. |
| `XactBuffer` | Waiting for I/O on a transaction status SLRU buffer. |
| `XactSLRU` | Waiting to access the transaction status SLRU cache. |
| `XactTruncation` | Waiting to execute `pg_xact_status` or update the oldest transaction ID available to it. |
| `XidGen` | Waiting to allocate a new transaction ID. |
 <a id="wait-event-timeout-table"></a>

**Table: Wait Events of Type `Timeout`**

| `Timeout` Wait Event | Description |
| --- | --- |
| `BaseBackupThrottle` | Waiting during base backup when throttling activity. |
| `CheckpointWriteDelay` | Waiting between writes while performing a checkpoint. |
| `PgSleep` | Waiting due to a call to `pg_sleep` or a sibling function. |
| `RecoveryApplyDelay` | Waiting to apply WAL during recovery because of a delay setting. |
| `RecoveryRetrieveRetryInterval` | Waiting during recovery when WAL data is not available from any source (`pg_wal`, archive or stream). |
| `RegisterSyncRequest` | Waiting while sending synchronization requests to the checkpointer, because the request queue is full. |
| `SpinDelay` | Waiting while acquiring a contended spinlock. |
| `VacuumDelay` | Waiting in a cost-based vacuum delay point. |
| `VacuumTruncate` | Waiting to acquire an exclusive lock to truncate off any empty pages at the end of a table vacuumed. |
| `WalSummarizerError` | Waiting after a WAL summarizer error. |


 Here are examples of how wait events can be viewed:

```sql

SELECT pid, wait_event_type, wait_event FROM pg_stat_activity WHERE wait_event is NOT NULL;
 pid  | wait_event_type | wait_event
------+-----------------+------------
 2540 | Lock            | relation
 6644 | LWLock          | ProcArray
(2 rows)
```


```sql

SELECT a.pid, a.wait_event, w.description
  FROM pg_stat_activity a JOIN
       pg_wait_events w ON (a.wait_event_type = w.type AND
                            a.wait_event = w.name)
  WHERE a.wait_event is NOT NULL and a.state = 'active';
-[ RECORD 1 ]------------------------------------------------------​------------
pid         | 686674
wait_event  | WALInitSync
description | Waiting for a newly initialized WAL file to reach durable storage
```


!!! note

    Extensions can add `Extension`, `InjectionPoint`, and `LWLock` events to the lists shown in [Wait Events of Type `Extension`](#wait-event-extension-table) and [Wait Events of Type `Lwlock`](#wait-event-lwlock-table). In some cases, the name of an `LWLock` assigned by an extension will not be available in all server processes. It might be reported as just “`extension`” rather than the extension-assigned name.
  <a id="monitoring-pg-stat-replication-view"></a>

### `pg_stat_replication`


 The `pg_stat_replication` view will contain one row per WAL sender process, showing statistics about replication to that sender's connected standby server. Only directly connected standbys are listed; no information is available about downstream standby servers.
 <a id="pg-stat-replication-view"></a>

**Table: `pg_stat_replication` View**

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
<p>Process ID of a WAL sender process</p></td>
</tr>
<tr>
<td><p><code>usesysid</code> <code>oid</code></p>
<p>OID of the user logged into this WAL sender process</p></td>
</tr>
<tr>
<td><p><code>usename</code> <code>name</code></p>
<p>Name of the user logged into this WAL sender process</p></td>
</tr>
<tr>
<td><p><code>application_name</code> <code>text</code></p>
<p>Name of the application that is connected to this WAL sender</p></td>
</tr>
<tr>
<td><p><code>client_addr</code> <code>inet</code></p>
<p>IP address of the client connected to this WAL sender. If this field is null, it indicates that the client is connected via a Unix socket on the server machine.</p></td>
</tr>
<tr>
<td><p><code>client_hostname</code> <code>text</code></p>
<p>Host name of the connected client, as reported by a reverse DNS lookup of <code>client_addr</code>. This field will only be non-null for IP connections, and only when <a href="../server-configuration/error-reporting-and-logging.md#guc-log-hostname">log_hostname</a> is enabled.</p></td>
</tr>
<tr>
<td><p><code>client_port</code> <code>integer</code></p>
<p>TCP port number that the client is using for communication with this WAL sender, or <code>-1</code> if a Unix socket is used</p></td>
</tr>
<tr>
<td><p><code>backend_start</code> <code>timestamp with time zone</code></p>
<p>Time when this process was started, i.e., when the client connected to this WAL sender</p></td>
</tr>
<tr>
<td><p><code>backend_xmin</code> <code>xid</code></p>
<p>This standby's <code>xmin</code> horizon reported by <a href="../server-configuration/replication.md#guc-hot-standby-feedback">hot_standby_feedback</a>.</p></td>
</tr>
<tr>
<td><p><code>state</code> <code>text</code></p>
<p>Current WAL sender state. Possible values are:</p>
<p>-  <code>startup</code>: This WAL sender is starting up. <br>
-  <code>catchup</code>: This WAL sender's connected standby is catching up with the primary. <br>
-  <code>streaming</code>: This WAL sender is streaming changes after its connected standby server has caught up with the primary. <br>
-  <code>backup</code>: This WAL sender is sending a backup. <br>
-  <code>stopping</code>: This WAL sender is stopping.</p></td>
</tr>
<tr>
<td><p><code>sent_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location sent on this connection</p></td>
</tr>
<tr>
<td><p><code>write_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location written to disk by this standby server</p></td>
</tr>
<tr>
<td><p><code>flush_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location flushed to disk by this standby server</p></td>
</tr>
<tr>
<td><p><code>replay_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location replayed into the database on this standby server</p></td>
</tr>
<tr>
<td><p><code>write_lag</code> <code>interval</code></p>
<p>Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written it (but not yet flushed it or applied it). This can be used to gauge the delay that <code>synchronous_commit</code> level <code>remote_write</code> incurred while committing if this server was configured as a synchronous standby.</p></td>
</tr>
<tr>
<td><p><code>flush_lag</code> <code>interval</code></p>
<p>Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written and flushed it (but not yet applied it). This can be used to gauge the delay that <code>synchronous_commit</code> level <code>on</code> incurred while committing if this server was configured as a synchronous standby.</p></td>
</tr>
<tr>
<td><p><code>replay_lag</code> <code>interval</code></p>
<p>Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written, flushed and applied it. This can be used to gauge the delay that <code>synchronous_commit</code> level <code>remote_apply</code> incurred while committing if this server was configured as a synchronous standby.</p></td>
</tr>
<tr>
<td><p><code>sync_priority</code> <code>integer</code></p>
<p>Priority of this standby server for being chosen as the synchronous standby in a priority-based synchronous replication. This has no effect in a quorum-based synchronous replication.</p></td>
</tr>
<tr>
<td><p><code>sync_state</code> <code>text</code></p>
<p>Synchronous state of this standby server. Possible values are:</p>
<p>-  <code>async</code>: This standby server is asynchronous. <br>
-  <code>potential</code>: This standby server is now asynchronous, but can potentially become synchronous if one of current synchronous ones fails. <br>
-  <code>sync</code>: This standby server is synchronous. <br>
-  <code>quorum</code>: This standby server is considered as a candidate for quorum standbys.</p></td>
</tr>
<tr>
<td><p><code>reply_time</code> <code>timestamp with time zone</code></p>
<p>Send time of last reply message received from standby server</p></td>
</tr>
</tbody>
</table>


 The lag times reported in the `pg_stat_replication` view are measurements of the time taken for recent WAL to be written, flushed and replayed and for the sender to know about it. These times represent the commit delay that was (or would have been) introduced by each synchronous commit level, if the remote server was configured as a synchronous standby. For an asynchronous standby, the `replay_lag` column approximates the delay before recent transactions became visible to queries. If the standby server has entirely caught up with the sending server and there is no more WAL activity, the most recently measured lag times will continue to be displayed for a short time and then show NULL.


 Lag times work automatically for physical replication. Logical decoding plugins may optionally emit tracking messages; if they do not, the tracking mechanism will simply display NULL lag.


!!! note

    The reported lag times are not predictions of how long it will take for the standby to catch up with the sending server assuming the current rate of replay. Such a system would show similar times while new WAL is being generated, but would differ when the sender becomes idle. In particular, when the standby has caught up completely, `pg_stat_replication` shows the time taken to write, flush and replay the most recent reported WAL location rather than zero as some users might expect. This is consistent with the goal of measuring synchronous commit and transaction visibility delays for recent write transactions. To reduce confusion for users expecting a different model of lag, the lag columns revert to NULL after a short time on a fully replayed idle system. Monitoring systems should choose whether to represent this as missing data, zero or continue to display the last known value.
  <a id="monitoring-pg-stat-replication-slots-view"></a>

### `pg_stat_replication_slots`


 The `pg_stat_replication_slots` view will contain one row per logical replication slot, showing statistics about its usage.
 <a id="pg-stat-replication-slots-view"></a>

**Table: `pg_stat_replication_slots` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>slot_name</code> <code>text</code></p>
<p>A unique, cluster-wide identifier for the replication slot</p></td>
</tr>
<tr>
<td><p><code>spill_txns</code> <code>bigint</code></p>
<p>Number of transactions spilled to disk once the memory used by logical decoding to decode changes from WAL has exceeded <code>logical_decoding_work_mem</code>. The counter gets incremented for both top-level transactions and subtransactions.</p></td>
</tr>
<tr>
<td><p><code>spill_count</code> <code>bigint</code></p>
<p>Number of times transactions were spilled to disk while decoding changes from WAL for this slot. This counter is incremented each time a transaction is spilled, and the same transaction may be spilled multiple times.</p></td>
</tr>
<tr>
<td><p><code>spill_bytes</code> <code>bigint</code></p>
<p>Amount of decoded transaction data spilled to disk while performing decoding of changes from WAL for this slot. This and other spill counters can be used to gauge the I/O which occurred during logical decoding and allow tuning <code>logical_decoding_work_mem</code>.</p></td>
</tr>
<tr>
<td><p><code>stream_txns</code> <code>bigint</code></p>
<p>Number of in-progress transactions streamed to the decoding output plugin after the memory used by logical decoding to decode changes from WAL for this slot has exceeded <code>logical_decoding_work_mem</code>. Streaming only works with top-level transactions (subtransactions can't be streamed independently), so the counter is not incremented for subtransactions.</p></td>
</tr>
<tr>
<td><p><code>stream_count</code><code>bigint</code></p>
<p>Number of times in-progress transactions were streamed to the decoding output plugin while decoding changes from WAL for this slot. This counter is incremented each time a transaction is streamed, and the same transaction may be streamed multiple times.</p></td>
</tr>
<tr>
<td><p><code>stream_bytes</code><code>bigint</code></p>
<p>Amount of transaction data decoded for streaming in-progress transactions to the decoding output plugin while decoding changes from WAL for this slot. This and other streaming counters for this slot can be used to tune <code>logical_decoding_work_mem</code>.</p></td>
</tr>
<tr>
<td><p><code>total_txns</code> <code>bigint</code></p>
<p>Number of decoded transactions sent to the decoding output plugin for this slot. This counts top-level transactions only, and is not incremented for subtransactions. Note that this includes the transactions that are streamed and/or spilled.</p></td>
</tr>
<tr>
<td><p><code>total_bytes</code><code>bigint</code></p>
<p>Amount of transaction data decoded for sending transactions to the decoding output plugin while decoding changes from WAL for this slot. Note that this includes data that is streamed and/or spilled.</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-wal-receiver-view"></a>

### `pg_stat_wal_receiver`


 The `pg_stat_wal_receiver` view will contain only one row, showing statistics about the WAL receiver from that receiver's connected server.
 <a id="pg-stat-wal-receiver-view"></a>

**Table: `pg_stat_wal_receiver` View**

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
<p>Process ID of the WAL receiver process</p></td>
</tr>
<tr>
<td><p><code>status</code> <code>text</code></p>
<p>Activity status of the WAL receiver process</p></td>
</tr>
<tr>
<td><p><code>receive_start_lsn</code> <code>pg_lsn</code></p>
<p>First write-ahead log location used when WAL receiver is started</p></td>
</tr>
<tr>
<td><p><code>receive_start_tli</code> <code>integer</code></p>
<p>First timeline number used when WAL receiver is started</p></td>
</tr>
<tr>
<td><p><code>written_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location already received and written to disk, but not flushed. This should not be used for data integrity checks.</p></td>
</tr>
<tr>
<td><p><code>flushed_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location already received and flushed to disk, the initial value of this field being the first log location used when WAL receiver is started</p></td>
</tr>
<tr>
<td><p><code>received_tli</code> <code>integer</code></p>
<p>Timeline number of last write-ahead log location received and flushed to disk, the initial value of this field being the timeline number of the first log location used when WAL receiver is started</p></td>
</tr>
<tr>
<td><p><code>last_msg_send_time</code> <code>timestamp with time zone</code></p>
<p>Send time of last message received from origin WAL sender</p></td>
</tr>
<tr>
<td><p><code>last_msg_receipt_time</code> <code>timestamp with time zone</code></p>
<p>Receipt time of last message received from origin WAL sender</p></td>
</tr>
<tr>
<td><p><code>latest_end_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location reported to origin WAL sender</p></td>
</tr>
<tr>
<td><p><code>latest_end_time</code> <code>timestamp with time zone</code></p>
<p>Time of last write-ahead log location reported to origin WAL sender</p></td>
</tr>
<tr>
<td><p><code>slot_name</code> <code>text</code></p>
<p>Replication slot name used by this WAL receiver</p></td>
</tr>
<tr>
<td><p><code>sender_host</code> <code>text</code></p>
<p>Host of the PostgreSQL instance this WAL receiver is connected to. This can be a host name, an IP address, or a directory path if the connection is via Unix socket. (The path case can be distinguished because it will always be an absolute path, beginning with <code>/</code>.)</p></td>
</tr>
<tr>
<td><p><code>sender_port</code> <code>integer</code></p>
<p>Port number of the PostgreSQL instance this WAL receiver is connected to.</p></td>
</tr>
<tr>
<td><p><code>conninfo</code> <code>text</code></p>
<p>Connection string used by this WAL receiver, with security-sensitive fields obfuscated.</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-recovery-prefetch"></a>

### `pg_stat_recovery_prefetch`


 The `pg_stat_recovery_prefetch` view will contain only one row. The columns `wal_distance`, `block_distance` and `io_depth` show current values, and the other columns show cumulative counters that can be reset with the `pg_stat_reset_shared` function.
 <a id="pg-stat-recovery-prefetch-view"></a>

**Table: `pg_stat_recovery_prefetch` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
<tr>
<td><p><code>prefetch</code> <code>bigint</code></p>
<p>Number of blocks prefetched because they were not in the buffer pool</p></td>
</tr>
<tr>
<td><p><code>hit</code> <code>bigint</code></p>
<p>Number of blocks not prefetched because they were already in the buffer pool</p></td>
</tr>
<tr>
<td><p><code>skip_init</code> <code>bigint</code></p>
<p>Number of blocks not prefetched because they would be zero-initialized</p></td>
</tr>
<tr>
<td><p><code>skip_new</code> <code>bigint</code></p>
<p>Number of blocks not prefetched because they didn't exist yet</p></td>
</tr>
<tr>
<td><p><code>skip_fpw</code> <code>bigint</code></p>
<p>Number of blocks not prefetched because a full page image was included in the WAL</p></td>
</tr>
<tr>
<td><p><code>skip_rep</code> <code>bigint</code></p>
<p>Number of blocks not prefetched because they were already recently prefetched</p></td>
</tr>
<tr>
<td><p><code>wal_distance</code> <code>int</code></p>
<p>How many bytes ahead the prefetcher is looking</p></td>
</tr>
<tr>
<td><p><code>block_distance</code> <code>int</code></p>
<p>How many blocks ahead the prefetcher is looking</p></td>
</tr>
<tr>
<td><p><code>io_depth</code> <code>int</code></p>
<p>How many prefetches have been initiated but are not yet known to have completed</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-subscription"></a>

### `pg_stat_subscription`
   <a id="pg-stat-subscription"></a>

**Table: `pg_stat_subscription` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>subid</code> <code>oid</code></p>
<p>OID of the subscription</p></td>
</tr>
<tr>
<td><p><code>subname</code> <code>name</code></p>
<p>Name of the subscription</p></td>
</tr>
<tr>
<td><p><code>worker_type</code> <code>text</code></p>
<p>Type of the subscription worker process. Possible types are <code>apply</code>, <code>parallel apply</code>, and <code>table synchronization</code>.</p></td>
</tr>
<tr>
<td><p><code>pid</code> <code>integer</code></p>
<p>Process ID of the subscription worker process</p></td>
</tr>
<tr>
<td><p><code>leader_pid</code> <code>integer</code></p>
<p>Process ID of the leader apply worker if this process is a parallel apply worker; NULL if this process is a leader apply worker or a table synchronization worker</p></td>
</tr>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the relation that the worker is synchronizing; NULL for the leader apply worker and parallel apply workers</p></td>
</tr>
<tr>
<td><p><code>received_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location received, the initial value of this field being 0; NULL for parallel apply workers</p></td>
</tr>
<tr>
<td><p><code>last_msg_send_time</code> <code>timestamp with time zone</code></p>
<p>Send time of last message received from origin WAL sender; NULL for parallel apply workers</p></td>
</tr>
<tr>
<td><p><code>last_msg_receipt_time</code> <code>timestamp with time zone</code></p>
<p>Receipt time of last message received from origin WAL sender; NULL for parallel apply workers</p></td>
</tr>
<tr>
<td><p><code>latest_end_lsn</code> <code>pg_lsn</code></p>
<p>Last write-ahead log location reported to origin WAL sender; NULL for parallel apply workers</p></td>
</tr>
<tr>
<td><p><code>latest_end_time</code> <code>timestamp with time zone</code></p>
<p>Time of last write-ahead log location reported to origin WAL sender; NULL for parallel apply workers</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-subscription-stats"></a>

### `pg_stat_subscription_stats`


 The `pg_stat_subscription_stats` view will contain one row per subscription.
 <a id="pg-stat-subscription-stats"></a>

**Table: `pg_stat_subscription_stats` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>subid</code> <code>oid</code></p>
<p>OID of the subscription</p></td>
</tr>
<tr>
<td><p><code>subname</code> <code>name</code></p>
<p>Name of the subscription</p></td>
</tr>
<tr>
<td><p><code>apply_error_count</code> <code>bigint</code></p>
<p>Number of times an error occurred while applying changes. Note that any conflict resulting in an apply error will be counted in both <code>apply_error_count</code> and the corresponding conflict count (e.g., <code>confl_*</code>).</p></td>
</tr>
<tr>
<td><p><code>sync_error_count</code> <code>bigint</code></p>
<p>Number of times an error occurred during the initial table synchronization</p></td>
</tr>
<tr>
<td><p><code>confl_insert_exists</code> <code>bigint</code></p>
<p>Number of times a row insertion violated a <code>NOT DEFERRABLE</code> unique constraint during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-insert-exists">insert_exists</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>confl_update_origin_differs</code> <code>bigint</code></p>
<p>Number of times an update was applied to a row that had been previously modified by another source during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-update-origin-differs">update_origin_differs</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>confl_update_exists</code> <code>bigint</code></p>
<p>Number of times that an updated row value violated a <code>NOT DEFERRABLE</code> unique constraint during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-update-exists">update_exists</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>confl_update_missing</code> <code>bigint</code></p>
<p>Number of times the tuple to be updated was not found during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-update-missing">update_missing</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>confl_delete_origin_differs</code> <code>bigint</code></p>
<p>Number of times a delete operation was applied to row that had been previously modified by another source during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-delete-origin-differs">delete_origin_differs</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>confl_delete_missing</code> <code>bigint</code></p>
<p>Number of times the tuple to be deleted was not found during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-delete-missing">delete_missing</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>confl_multiple_unique_conflicts</code> <code>bigint</code></p>
<p>Number of times a row insertion or an updated row values violated multiple <code>NOT DEFERRABLE</code> unique constraints during the application of changes. See <a href="../logical-replication/conflicts.md#conflict-multiple-unique-conflicts">multiple_unique_conflicts</a> for details about this conflict.</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-ssl-view"></a>

### `pg_stat_ssl`


 The `pg_stat_ssl` view will contain one row per backend or WAL sender process, showing statistics about SSL usage on this connection. It can be joined to `pg_stat_activity` or `pg_stat_replication` on the `pid` column to get more details about the connection.
 <a id="pg-stat-ssl-view"></a>

**Table: `pg_stat_ssl` View**

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
<p>Process ID of a backend or WAL sender process</p></td>
</tr>
<tr>
<td><p><code>ssl</code> <code>boolean</code></p>
<p>True if SSL is used on this connection</p></td>
</tr>
<tr>
<td><p><code>version</code> <code>text</code></p>
<p>Version of SSL in use, or NULL if SSL is not in use on this connection</p></td>
</tr>
<tr>
<td><p><code>cipher</code> <code>text</code></p>
<p>Name of SSL cipher in use, or NULL if SSL is not in use on this connection</p></td>
</tr>
<tr>
<td><p><code>bits</code> <code>integer</code></p>
<p>Number of bits in the encryption algorithm used, or NULL if SSL is not used on this connection</p></td>
</tr>
<tr>
<td><p><code>client_dn</code> <code>text</code></p>
<p>Distinguished Name (DN) field from the client certificate used, or NULL if no client certificate was supplied or if SSL is not in use on this connection. This field is truncated if the DN field is longer than <code>NAMEDATALEN</code> (64 characters in a standard build).</p></td>
</tr>
<tr>
<td><p><code>client_serial</code> <code>numeric</code></p>
<p>Serial number of the client certificate, or NULL if no client certificate was supplied or if SSL is not in use on this connection. The combination of certificate serial number and certificate issuer uniquely identifies a certificate (unless the issuer erroneously reuses serial numbers).</p></td>
</tr>
<tr>
<td><p><code>issuer_dn</code> <code>text</code></p>
<p>DN of the issuer of the client certificate, or NULL if no client certificate was supplied or if SSL is not in use on this connection. This field is truncated like <code>client_dn</code>.</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-gssapi-view"></a>

### `pg_stat_gssapi`


 The `pg_stat_gssapi` view will contain one row per backend, showing information about GSSAPI usage on this connection. It can be joined to `pg_stat_activity` or `pg_stat_replication` on the `pid` column to get more details about the connection.
 <a id="pg-stat-gssapi-view"></a>

**Table: `pg_stat_gssapi` View**

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
<p>Process ID of a backend</p></td>
</tr>
<tr>
<td><p><code>gss_authenticated</code> <code>boolean</code></p>
<p>True if GSSAPI authentication was used for this connection</p></td>
</tr>
<tr>
<td><p><code>principal</code> <code>text</code></p>
<p>Principal used to authenticate this connection, or NULL if GSSAPI was not used to authenticate this connection. This field is truncated if the principal is longer than <code>NAMEDATALEN</code> (64 characters in a standard build).</p></td>
</tr>
<tr>
<td><p><code>encrypted</code> <code>boolean</code></p>
<p>True if GSSAPI encryption is in use on this connection</p></td>
</tr>
<tr>
<td><p><code>credentials_delegated</code> <code>boolean</code></p>
<p>True if GSSAPI credentials were delegated on this connection.</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-archiver-view"></a>

### `pg_stat_archiver`


 The `pg_stat_archiver` view will always have a single row, containing data about the archiver process of the cluster.
 <a id="pg-stat-archiver-view"></a>

**Table: `pg_stat_archiver` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>archived_count</code> <code>bigint</code></p>
<p>Number of WAL files that have been successfully archived</p></td>
</tr>
<tr>
<td><p><code>last_archived_wal</code> <code>text</code></p>
<p>Name of the WAL file most recently successfully archived</p></td>
</tr>
<tr>
<td><p><code>last_archived_time</code> <code>timestamp with time zone</code></p>
<p>Time of the most recent successful archive operation</p></td>
</tr>
<tr>
<td><p><code>failed_count</code> <code>bigint</code></p>
<p>Number of failed attempts for archiving WAL files</p></td>
</tr>
<tr>
<td><p><code>last_failed_wal</code> <code>text</code></p>
<p>Name of the WAL file of the most recent failed archival operation</p></td>
</tr>
<tr>
<td><p><code>last_failed_time</code> <code>timestamp with time zone</code></p>
<p>Time of the most recent failed archival operation</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>


 Normally, WAL files are archived in order, oldest to newest, but that is not guaranteed, and does not hold under special circumstances like when promoting a standby or after crash recovery. Therefore it is not safe to assume that all files older than `last_archived_wal` have also been successfully archived.
  <a id="monitoring-pg-stat-io-view"></a>

### `pg_stat_io`


 The `pg_stat_io` view will contain one row for each combination of backend type, target I/O object, and I/O context, showing cluster-wide I/O statistics. Combinations which do not make sense are omitted.


 Currently, I/O on relations (e.g. tables, indexes) and WAL activity are tracked. However, relation I/O which bypasses shared buffers (e.g. when moving a table from one tablespace to another) is currently not tracked.
 <a id="pg-stat-io-view"></a>

**Table: `pg_stat_io` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>backend_type</code> <code>text</code></p>
<p>Type of backend (e.g. background worker, autovacuum worker). See <a href="#monitoring-pg-stat-activity-view"><code>pg_stat_activity</code></a> for more information on <code>backend_type</code>s. Some <code>backend_type</code>s do not accumulate I/O operation statistics and will not be included in the view.</p></td>
</tr>
<tr>
<td><p><code>object</code> <code>text</code></p>
<p>Target object of an I/O operation. Possible values are:</p>
<p>-  <code>relation</code>: Permanent relations. <br>
-  <code>temp relation</code>: Temporary relations. <br>
-  <code>wal</code>: Write Ahead Logs.</p></td>
</tr>
<tr>
<td><p><code>context</code> <code>text</code></p>
<p>The context of an I/O operation. Possible values are:</p>
<p>-  <code>normal</code>: The default or standard <code>context</code> for a type of I/O operation. For example, by default, relation data is read into and written out from shared buffers. Thus, reads and writes of relation data to and from shared buffers are tracked in <code>context</code> <code>normal</code>. <br>
-  <code>init</code>: I/O operations performed while creating the WAL segments are tracked in <code>context</code> <code>init</code>. <br>
-  <code>vacuum</code>: I/O operations performed outside of shared buffers while vacuuming and analyzing permanent relations. Temporary table vacuums use the same local buffer pool as other temporary table I/O operations and are tracked in <code>context</code> <code>normal</code>. <br>
-  <code>bulkread</code>: Certain large read I/O operations done outside of shared buffers, for example, a sequential scan of a large table. <br>
-  <code>bulkwrite</code>: Certain large write I/O operations done outside of shared buffers, such as <code>COPY</code>.</p></td>
</tr>
<tr>
<td><p><code>reads</code> <code>bigint</code></p>
<p>Number of read operations.</p></td>
</tr>
<tr>
<td><p><code>read_bytes</code> <code>numeric</code></p>
<p>The total size of read operations in bytes.</p></td>
</tr>
<tr>
<td><p><code>read_time</code> <code>double precision</code></p>
<p>Time spent waiting for read operations in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled and <code>object</code> is not <code>wal</code>, or if <a href="../server-configuration/run-time-statistics.md#guc-track-wal-io-timing">track_wal_io_timing</a> is enabled and <code>object</code> is <code>wal</code>, otherwise zero)</p></td>
</tr>
<tr>
<td><p><code>writes</code> <code>bigint</code></p>
<p>Number of write operations.</p></td>
</tr>
<tr>
<td><p><code>write_bytes</code> <code>numeric</code></p>
<p>The total size of write operations in bytes.</p></td>
</tr>
<tr>
<td><p><code>write_time</code> <code>double precision</code></p>
<p>Time spent waiting for write operations in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled and <code>object</code> is not <code>wal</code>, or if <a href="../server-configuration/run-time-statistics.md#guc-track-wal-io-timing">track_wal_io_timing</a> is enabled and <code>object</code> is <code>wal</code>, otherwise zero)</p></td>
</tr>
<tr>
<td><p><code>writebacks</code> <code>bigint</code></p>
<p>Number of units of size <code>BLCKSZ</code> (typically 8kB) which the process requested the kernel write out to permanent storage.</p></td>
</tr>
<tr>
<td><p><code>writeback_time</code> <code>double precision</code></p>
<p>Time spent waiting for writeback operations in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled, otherwise zero). This includes the time spent queueing write-out requests and, potentially, the time spent to write out the dirty data.</p></td>
</tr>
<tr>
<td><p><code>extends</code> <code>bigint</code></p>
<p>Number of relation extend operations.</p></td>
</tr>
<tr>
<td><p><code>extend_bytes</code> <code>numeric</code></p>
<p>The total size of relation extend operations in bytes.</p></td>
</tr>
<tr>
<td><p><code>extend_time</code> <code>double precision</code></p>
<p>Time spent waiting for extend operations in milliseconds. (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled and <code>object</code> is not <code>wal</code>, or if <a href="../server-configuration/run-time-statistics.md#guc-track-wal-io-timing">track_wal_io_timing</a> is enabled and <code>object</code> is <code>wal</code>, otherwise zero)</p></td>
</tr>
<tr>
<td><p><code>hits</code> <code>bigint</code></p>
<p>The number of times a desired block was found in a shared buffer.</p></td>
</tr>
<tr>
<td><p><code>evictions</code> <code>bigint</code></p>
<p>Number of times a block has been written out from a shared or local buffer in order to make it available for another use.</p>
<p>In <code>context</code> <code>normal</code>, this counts the number of times a block was evicted from a buffer and replaced with another block. In <code>context</code>s <code>bulkwrite</code>, <code>bulkread</code>, and <code>vacuum</code>, this counts the number of times a block was evicted from shared buffers in order to add the shared buffer to a separate, size-limited ring buffer for use in a bulk I/O operation.</p></td>
</tr>
<tr>
<td><p><code>reuses</code> <code>bigint</code></p>
<p>The number of times an existing buffer in a size-limited ring buffer outside of shared buffers was reused as part of an I/O operation in the <code>bulkread</code>, <code>bulkwrite</code>, or <code>vacuum</code> <code>context</code>s.</p></td>
</tr>
<tr>
<td><p><code>fsyncs</code> <code>bigint</code></p>
<p>Number of <code>fsync</code> calls. These are only tracked in <code>context</code> <code>normal</code>.</p></td>
</tr>
<tr>
<td><p><code>fsync_time</code> <code>double precision</code></p>
<p>Time spent waiting for fsync operations in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled and <code>object</code> is not <code>wal</code>, or if <a href="../server-configuration/run-time-statistics.md#guc-track-wal-io-timing">track_wal_io_timing</a> is enabled and <code>object</code> is <code>wal</code>, otherwise zero)</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset.</p></td>
</tr>
</tbody>
</table>


 Some backend types never perform I/O operations on some I/O objects and/or in some I/O contexts. These rows are omitted from the view. For example, the checkpointer does not checkpoint temporary tables, so there will be no rows for `backend_type` `checkpointer` and `object` `temp relation`.


 In addition, some I/O operations will never be performed either by certain backend types or on certain I/O objects and/or in certain I/O contexts. These cells will be NULL. For example, temporary tables are not `fsync`ed, so `fsyncs` will be NULL for `object` `temp relation`. Also, the background writer does not perform reads, so `reads` will be NULL in rows for `backend_type` `background writer`.


 For the `object` `wal`, `fsyncs` and `fsync_time` track the fsync activity of WAL files done in `issue_xlog_fsync`. `writes` and `write_time` track the write activity of WAL files done in `XLogWrite`. See [WAL Configuration](../reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration) for more information.


 `pg_stat_io` can be used to inform database tuning. For example:

-  A high `evictions` count can indicate that shared buffers should be increased.
-  Client backends rely on the checkpointer to ensure data is persisted to permanent storage. Large numbers of `fsyncs` by `client backend`s could indicate a misconfiguration of shared buffers or of the checkpointer. More information on configuring the checkpointer can be found in [WAL Configuration](../reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration).
-  Normally, client backends should be able to rely on auxiliary processes like the checkpointer and the background writer to write out dirty data as much as possible. Large numbers of writes by client backends could indicate a misconfiguration of shared buffers or of the checkpointer. More information on configuring the checkpointer can be found in [WAL Configuration](../reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration).


!!! note

    Columns tracking I/O wait time will only be non-zero when [track_io_timing](../server-configuration/run-time-statistics.md#guc-track-io-timing) is enabled. The user should be careful when referencing these columns in combination with their corresponding I/O operations in case `track_io_timing` was not enabled for the entire time since the last stats reset.
  <a id="monitoring-pg-stat-bgwriter-view"></a>

### `pg_stat_bgwriter`


 The `pg_stat_bgwriter` view will always have a single row, containing data about the background writer of the cluster.
 <a id="pg-stat-bgwriter-view"></a>

**Table: `pg_stat_bgwriter` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>buffers_clean</code> <code>bigint</code></p>
<p>Number of buffers written by the background writer</p></td>
</tr>
<tr>
<td><p><code>maxwritten_clean</code> <code>bigint</code></p>
<p>Number of times the background writer stopped a cleaning scan because it had written too many buffers</p></td>
</tr>
<tr>
<td><p><code>buffers_alloc</code> <code>bigint</code></p>
<p>Number of buffers allocated</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-checkpointer-view"></a>

### `pg_stat_checkpointer`


 The `pg_stat_checkpointer` view will always have a single row, containing data about the checkpointer process of the cluster.
 <a id="pg-stat-checkpointer-view"></a>

**Table: `pg_stat_checkpointer` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>num_timed</code> <code>bigint</code></p>
<p>Number of scheduled checkpoints due to timeout</p></td>
</tr>
<tr>
<td><p><code>num_requested</code> <code>bigint</code></p>
<p>Number of requested checkpoints</p></td>
</tr>
<tr>
<td><p><code>num_done</code> <code>bigint</code></p>
<p>Number of checkpoints that have been performed</p></td>
</tr>
<tr>
<td><p><code>restartpoints_timed</code> <code>bigint</code></p>
<p>Number of scheduled restartpoints due to timeout or after a failed attempt to perform it</p></td>
</tr>
<tr>
<td><p><code>restartpoints_req</code> <code>bigint</code></p>
<p>Number of requested restartpoints</p></td>
</tr>
<tr>
<td><p><code>restartpoints_done</code> <code>bigint</code></p>
<p>Number of restartpoints that have been performed</p></td>
</tr>
<tr>
<td><p><code>write_time</code> <code>double precision</code></p>
<p>Total amount of time that has been spent in the portion of processing checkpoints and restartpoints where files are written to disk, in milliseconds</p></td>
</tr>
<tr>
<td><p><code>sync_time</code> <code>double precision</code></p>
<p>Total amount of time that has been spent in the portion of processing checkpoints and restartpoints where files are synchronized to disk, in milliseconds</p></td>
</tr>
<tr>
<td><p><code>buffers_written</code> <code>bigint</code></p>
<p>Number of shared buffers written during checkpoints and restartpoints</p></td>
</tr>
<tr>
<td><p><code>slru_written</code> <code>bigint</code></p>
<p>Number of SLRU buffers written during checkpoints and restartpoints</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>


 Checkpoints may be skipped if the server has been idle since the last one. `num_timed` and `num_requested` count both completed and skipped checkpoints, while `num_done` tracks only the completed ones. Similarly, restartpoints may be skipped if the last replayed checkpoint record is already the last restartpoint. `restartpoints_timed` and `restartpoints_req` count both completed and skipped restartpoints, while `restartpoints_done` tracks only the completed ones.
  <a id="monitoring-pg-stat-wal-view"></a>

### `pg_stat_wal`


 The `pg_stat_wal` view will always have a single row, containing data about WAL activity of the cluster.
 <a id="pg-stat-wal-view"></a>

**Table: `pg_stat_wal` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>wal_records</code> <code>bigint</code></p>
<p>Total number of WAL records generated</p></td>
</tr>
<tr>
<td><p><code>wal_fpi</code> <code>bigint</code></p>
<p>Total number of WAL full page images generated</p></td>
</tr>
<tr>
<td><p><code>wal_bytes</code> <code>numeric</code></p>
<p>Total amount of WAL generated in bytes</p></td>
</tr>
<tr>
<td><p><code>wal_buffers_full</code> <code>bigint</code></p>
<p>Number of times WAL data was written to disk because WAL buffers became full</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-database-view"></a>

### `pg_stat_database`


 The `pg_stat_database` view will contain one row for each database in the cluster, plus one for shared objects, showing database-wide statistics.
 <a id="pg-stat-database-view"></a>

**Table: `pg_stat_database` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of this database, or 0 for objects belonging to a shared relation</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of this database, or <code>NULL</code> for shared objects.</p></td>
</tr>
<tr>
<td><p><code>numbackends</code> <code>integer</code></p>
<p>Number of backends currently connected to this database, or <code>NULL</code> for shared objects. This is the only column in this view that returns a value reflecting current state; all other columns return the accumulated values since the last reset.</p></td>
</tr>
<tr>
<td><p><code>xact_commit</code> <code>bigint</code></p>
<p>Number of transactions in this database that have been committed</p></td>
</tr>
<tr>
<td><p><code>xact_rollback</code> <code>bigint</code></p>
<p>Number of transactions in this database that have been rolled back</p></td>
</tr>
<tr>
<td><p><code>blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read in this database</p></td>
</tr>
<tr>
<td><p><code>blks_hit</code> <code>bigint</code></p>
<p>Number of times disk blocks were found already in the buffer cache, so that a read was not necessary (this only includes hits in the PostgreSQL buffer cache, not the operating system's file system cache)</p></td>
</tr>
<tr>
<td><p><code>tup_returned</code> <code>bigint</code></p>
<p>Number of live rows fetched by sequential scans and index entries returned by index scans in this database</p></td>
</tr>
<tr>
<td><p><code>tup_fetched</code> <code>bigint</code></p>
<p>Number of live rows fetched by index scans in this database</p></td>
</tr>
<tr>
<td><p><code>tup_inserted</code> <code>bigint</code></p>
<p>Number of rows inserted by queries in this database</p></td>
</tr>
<tr>
<td><p><code>tup_updated</code> <code>bigint</code></p>
<p>Number of rows updated by queries in this database</p></td>
</tr>
<tr>
<td><p><code>tup_deleted</code> <code>bigint</code></p>
<p>Number of rows deleted by queries in this database</p></td>
</tr>
<tr>
<td><p><code>conflicts</code> <code>bigint</code></p>
<p>Number of queries canceled due to conflicts with recovery in this database. (Conflicts occur only on standby servers; see <a href="#monitoring-pg-stat-database-conflicts-view"><code>pg_stat_database_conflicts</code></a> for details.)</p></td>
</tr>
<tr>
<td><p><code>temp_files</code> <code>bigint</code></p>
<p>Number of temporary files created by queries in this database. All temporary files are counted, regardless of why the temporary file was created (e.g., sorting or hashing), and regardless of the <a href="../server-configuration/error-reporting-and-logging.md#guc-log-temp-files">log_temp_files</a> setting.</p></td>
</tr>
<tr>
<td><p><code>temp_bytes</code> <code>bigint</code></p>
<p>Total amount of data written to temporary files by queries in this database. All temporary files are counted, regardless of why the temporary file was created, and regardless of the <a href="../server-configuration/error-reporting-and-logging.md#guc-log-temp-files">log_temp_files</a> setting.</p></td>
</tr>
<tr>
<td><p><code>deadlocks</code> <code>bigint</code></p>
<p>Number of deadlocks detected in this database</p></td>
</tr>
<tr>
<td><p><code>checksum_failures</code> <code>bigint</code></p>
<p>Number of data page checksum failures detected in this database (or on a shared object), or NULL if data checksums are disabled.</p></td>
</tr>
<tr>
<td><p><code>checksum_last_failure</code> <code>timestamp with time zone</code></p>
<p>Time at which the last data page checksum failure was detected in this database (or on a shared object), or NULL if data checksums are disabled.</p></td>
</tr>
<tr>
<td><p><code>blk_read_time</code> <code>double precision</code></p>
<p>Time spent reading data file blocks by backends in this database, in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled, otherwise zero)</p></td>
</tr>
<tr>
<td><p><code>blk_write_time</code> <code>double precision</code></p>
<p>Time spent writing data file blocks by backends in this database, in milliseconds (if <a href="../server-configuration/run-time-statistics.md#guc-track-io-timing">track_io_timing</a> is enabled, otherwise zero)</p></td>
</tr>
<tr>
<td><p><code>session_time</code> <code>double precision</code></p>
<p>Time spent by database sessions in this database, in milliseconds (note that statistics are only updated when the state of a session changes, so if sessions have been idle for a long time, this idle time won't be included)</p></td>
</tr>
<tr>
<td><p><code>active_time</code> <code>double precision</code></p>
<p>Time spent executing SQL statements in this database, in milliseconds (this corresponds to the states <code>active</code> and <code>fastpath function call</code> in <a href="#monitoring-pg-stat-activity-view"><code>pg_stat_activity</code></a>)</p></td>
</tr>
<tr>
<td><p><code>idle_in_transaction_time</code> <code>double precision</code></p>
<p>Time spent idling while in a transaction in this database, in milliseconds (this corresponds to the states <code>idle in transaction</code> and <code>idle in transaction (aborted)</code> in <a href="#monitoring-pg-stat-activity-view"><code>pg_stat_activity</code></a>)</p></td>
</tr>
<tr>
<td><p><code>sessions</code> <code>bigint</code></p>
<p>Total number of sessions established to this database</p></td>
</tr>
<tr>
<td><p><code>sessions_abandoned</code> <code>bigint</code></p>
<p>Number of database sessions to this database that were terminated because connection to the client was lost</p></td>
</tr>
<tr>
<td><p><code>sessions_fatal</code> <code>bigint</code></p>
<p>Number of database sessions to this database that were terminated by fatal errors</p></td>
</tr>
<tr>
<td><p><code>sessions_killed</code> <code>bigint</code></p>
<p>Number of database sessions to this database that were terminated by operator intervention</p></td>
</tr>
<tr>
<td><p><code>parallel_workers_to_launch</code> <code>bigint</code></p>
<p>Number of parallel workers planned to be launched by queries on this database</p></td>
</tr>
<tr>
<td><p><code>parallel_workers_launched</code> <code>bigint</code></p>
<p>Number of parallel workers launched by queries on this database</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-database-conflicts-view"></a>

### `pg_stat_database_conflicts`


 The `pg_stat_database_conflicts` view will contain one row per database, showing database-wide statistics about query cancels occurring due to conflicts with recovery on standby servers. This view will only contain information on standby servers, since conflicts do not occur on primary servers.
 <a id="pg-stat-database-conflicts-view"></a>

**Table: `pg_stat_database_conflicts` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>datid</code> <code>oid</code></p>
<p>OID of a database</p></td>
</tr>
<tr>
<td><p><code>datname</code> <code>name</code></p>
<p>Name of this database</p></td>
</tr>
<tr>
<td><p><code>confl_tablespace</code> <code>bigint</code></p>
<p>Number of queries in this database that have been canceled due to dropped tablespaces</p></td>
</tr>
<tr>
<td><p><code>confl_lock</code> <code>bigint</code></p>
<p>Number of queries in this database that have been canceled due to lock timeouts</p></td>
</tr>
<tr>
<td><p><code>confl_snapshot</code> <code>bigint</code></p>
<p>Number of queries in this database that have been canceled due to old snapshots</p></td>
</tr>
<tr>
<td><p><code>confl_bufferpin</code> <code>bigint</code></p>
<p>Number of queries in this database that have been canceled due to pinned buffers</p></td>
</tr>
<tr>
<td><p><code>confl_deadlock</code> <code>bigint</code></p>
<p>Number of queries in this database that have been canceled due to deadlocks</p></td>
</tr>
<tr>
<td><p><code>confl_active_logicalslot</code> <code>bigint</code></p>
<p>Number of uses of logical slots in this database that have been canceled due to old snapshots or too low a <a href="../server-configuration/write-ahead-log.md#guc-wal-level">wal_level</a> on the primary</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-all-tables-view"></a>

### `pg_stat_all_tables`


 The `pg_stat_all_tables` view will contain one row for each table in the current database (including TOAST tables), showing statistics about accesses to that specific table. The `pg_stat_user_tables` and `pg_stat_sys_tables` views contain the same information, but filtered to only show user and system tables respectively.
 <a id="pg-stat-all-tables-view"></a>

**Table: `pg_stat_all_tables` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of a table</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code></p>
<p>Name of the schema that this table is in</p></td>
</tr>
<tr>
<td><p><code>relname</code> <code>name</code></p>
<p>Name of this table</p></td>
</tr>
<tr>
<td><p><code>seq_scan</code> <code>bigint</code></p>
<p>Number of sequential scans initiated on this table</p></td>
</tr>
<tr>
<td><p><code>last_seq_scan</code> <code>timestamp with time zone</code></p>
<p>The time of the last sequential scan on this table, based on the most recent transaction stop time</p></td>
</tr>
<tr>
<td><p><code>seq_tup_read</code> <code>bigint</code></p>
<p>Number of live rows fetched by sequential scans</p></td>
</tr>
<tr>
<td><p><code>idx_scan</code> <code>bigint</code></p>
<p>Number of index scans initiated on this table</p></td>
</tr>
<tr>
<td><p><code>last_idx_scan</code> <code>timestamp with time zone</code></p>
<p>The time of the last index scan on this table, based on the most recent transaction stop time</p></td>
</tr>
<tr>
<td><p><code>idx_tup_fetch</code> <code>bigint</code></p>
<p>Number of live rows fetched by index scans</p></td>
</tr>
<tr>
<td><p><code>n_tup_ins</code> <code>bigint</code></p>
<p>Total number of rows inserted</p></td>
</tr>
<tr>
<td><p><code>n_tup_upd</code> <code>bigint</code></p>
<p>Total number of rows updated. (This includes row updates counted in <code>n_tup_hot_upd</code> and <code>n_tup_newpage_upd</code>, and remaining non-HOT updates.)</p></td>
</tr>
<tr>
<td><p><code>n_tup_del</code> <code>bigint</code></p>
<p>Total number of rows deleted</p></td>
</tr>
<tr>
<td><p><code>n_tup_hot_upd</code> <code>bigint</code></p>
<p>Number of rows <a href="../../internals/database-physical-storage/heap-only-tuples-hot.md#storage-hot">HOT updated</a>. These are updates where no successor versions are required in indexes.</p></td>
</tr>
<tr>
<td><p><code>n_tup_newpage_upd</code> <code>bigint</code></p>
<p>Number of rows updated where the successor version goes onto a <em>new</em> heap page, leaving behind an original version with a <a href="../../internals/database-physical-storage/database-page-layout.md#storage-tuple-layout"><code>t_ctid</code> field</a> that points to a different heap page. These are always non-HOT updates.</p></td>
</tr>
<tr>
<td><p><code>n_live_tup</code> <code>bigint</code></p>
<p>Estimated number of live rows</p></td>
</tr>
<tr>
<td><p><code>n_dead_tup</code> <code>bigint</code></p>
<p>Estimated number of dead rows</p></td>
</tr>
<tr>
<td><p><code>n_mod_since_analyze</code> <code>bigint</code></p>
<p>Estimated number of rows modified since this table was last analyzed</p></td>
</tr>
<tr>
<td><p><code>n_ins_since_vacuum</code> <code>bigint</code></p>
<p>Estimated number of rows inserted since this table was last vacuumed (not counting <code>VACUUM FULL</code>)</p></td>
</tr>
<tr>
<td><p><code>last_vacuum</code> <code>timestamp with time zone</code></p>
<p>Last time at which this table was manually vacuumed (not counting <code>VACUUM FULL</code>)</p></td>
</tr>
<tr>
<td><p><code>last_autovacuum</code> <code>timestamp with time zone</code></p>
<p>Last time at which this table was vacuumed by the autovacuum daemon</p></td>
</tr>
<tr>
<td><p><code>last_analyze</code> <code>timestamp with time zone</code></p>
<p>Last time at which this table was manually analyzed</p></td>
</tr>
<tr>
<td><p><code>last_autoanalyze</code> <code>timestamp with time zone</code></p>
<p>Last time at which this table was analyzed by the autovacuum daemon</p></td>
</tr>
<tr>
<td><p><code>vacuum_count</code> <code>bigint</code></p>
<p>Number of times this table has been manually vacuumed (not counting <code>VACUUM FULL</code>)</p></td>
</tr>
<tr>
<td><p><code>autovacuum_count</code> <code>bigint</code></p>
<p>Number of times this table has been vacuumed by the autovacuum daemon</p></td>
</tr>
<tr>
<td><p><code>analyze_count</code> <code>bigint</code></p>
<p>Number of times this table has been manually analyzed</p></td>
</tr>
<tr>
<td><p><code>autoanalyze_count</code> <code>bigint</code></p>
<p>Number of times this table has been analyzed by the autovacuum daemon</p></td>
</tr>
<tr>
<td><p><code>total_vacuum_time</code> <code>double precision</code></p>
<p>Total time this table has been manually vacuumed, in milliseconds (not counting <code>VACUUM FULL</code>). (This includes the time spent sleeping due to cost-based delays.)</p></td>
</tr>
<tr>
<td><p><code>total_autovacuum_time</code> <code>double precision</code></p>
<p>Total time this table has been vacuumed by the autovacuum daemon, in milliseconds. (This includes the time spent sleeping due to cost-based delays.)</p></td>
</tr>
<tr>
<td><p><code>total_analyze_time</code> <code>double precision</code></p>
<p>Total time this table has been manually analyzed, in milliseconds. (This includes the time spent sleeping due to cost-based delays.)</p></td>
</tr>
<tr>
<td><p><code>total_autoanalyze_time</code> <code>double precision</code></p>
<p>Total time this table has been analyzed by the autovacuum daemon, in milliseconds. (This includes the time spent sleeping due to cost-based delays.)</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-all-indexes-view"></a>

### `pg_stat_all_indexes`


 The `pg_stat_all_indexes` view will contain one row for each index in the current database, showing statistics about accesses to that specific index. The `pg_stat_user_indexes` and `pg_stat_sys_indexes` views contain the same information, but filtered to only show user and system indexes respectively.
 <a id="pg-stat-all-indexes-view"></a>

**Table: `pg_stat_all_indexes` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table for this index</p></td>
</tr>
<tr>
<td><p><code>indexrelid</code> <code>oid</code></p>
<p>OID of this index</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code></p>
<p>Name of the schema this index is in</p></td>
</tr>
<tr>
<td><p><code>relname</code> <code>name</code></p>
<p>Name of the table for this index</p></td>
</tr>
<tr>
<td><p><code>indexrelname</code> <code>name</code></p>
<p>Name of this index</p></td>
</tr>
<tr>
<td><p><code>idx_scan</code> <code>bigint</code></p>
<p>Number of index scans initiated on this index</p></td>
</tr>
<tr>
<td><p><code>last_idx_scan</code> <code>timestamp with time zone</code></p>
<p>The time of the last scan on this index, based on the most recent transaction stop time</p></td>
</tr>
<tr>
<td><p><code>idx_tup_read</code> <code>bigint</code></p>
<p>Number of index entries returned by scans on this index</p></td>
</tr>
<tr>
<td><p><code>idx_tup_fetch</code> <code>bigint</code></p>
<p>Number of live table rows fetched by simple index scans using this index</p></td>
</tr>
</tbody>
</table>


 Indexes can be used by simple index scans, “bitmap” index scans, and the optimizer. In a bitmap scan the output of several indexes can be combined via AND or OR rules, so it is difficult to associate individual heap row fetches with specific indexes when a bitmap scan is used. Therefore, a bitmap scan increments the `pg_stat_all_indexes`.`idx_tup_read` count(s) for the index(es) it uses, and it increments the `pg_stat_all_tables`.`idx_tup_fetch` count for the table, but it does not affect `pg_stat_all_indexes`.`idx_tup_fetch`. The optimizer also accesses indexes to check for supplied constants whose values are outside the recorded range of the optimizer statistics because the optimizer statistics might be stale.


!!! note

    The `idx_tup_read` and `idx_tup_fetch` counts can be different even without any use of bitmap scans, because `idx_tup_read` counts index entries retrieved from the index while `idx_tup_fetch` counts live rows fetched from the table. The latter will be less if any dead or not-yet-committed rows are fetched using the index, or if any heap fetches are avoided by means of an index-only scan.


!!! note

    Index scans may sometimes perform multiple index searches per execution. Each index search increments `pg_stat_all_indexes`.`idx_scan`, so it's possible for the count of index scans to significantly exceed the total number of index scan executor node executions.


     This can happen with queries that use certain SQL constructs to search for rows matching any value out of a list or array of multiple scalar values (see [Row and Array Comparisons](../../the-sql-language/functions-and-operators/row-and-array-comparisons.md#functions-comparisons)). It can also happen to queries with a <em>column_name</em><code> =
         </code><em>value1</em><code> OR
         </code><em>column_name</em><code> =
         </code><em>value2</em><code> ...</code> construct, though only when the optimizer transforms the construct into an equivalent multi-valued array representation. Similarly, when B-tree index scans use the skip scan optimization, an index search is performed each time the scan is repositioned to the next index leaf page that might have matching tuples (see [Multicolumn Indexes](../../the-sql-language/indexes/multicolumn-indexes.md#indexes-multicolumn)).


!!! tip

    `EXPLAIN ANALYZE` outputs the total number of index searches performed by each index scan node. See [`EXPLAIN ANALYZE`](../../the-sql-language/performance-tips/using-explain.md#using-explain-analyze) for an example demonstrating how this works.
  <a id="monitoring-pg-statio-all-tables-view"></a>

### `pg_statio_all_tables`


 The `pg_statio_all_tables` view will contain one row for each table in the current database (including TOAST tables), showing statistics about I/O on that specific table. The `pg_statio_user_tables` and `pg_statio_sys_tables` views contain the same information, but filtered to only show user and system tables respectively.
 <a id="pg-statio-all-tables-view"></a>

**Table: `pg_statio_all_tables` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of a table</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code></p>
<p>Name of the schema that this table is in</p></td>
</tr>
<tr>
<td><p><code>relname</code> <code>name</code></p>
<p>Name of this table</p></td>
</tr>
<tr>
<td><p><code>heap_blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read from this table</p></td>
</tr>
<tr>
<td><p><code>heap_blks_hit</code> <code>bigint</code></p>
<p>Number of buffer hits in this table</p></td>
</tr>
<tr>
<td><p><code>idx_blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read from all indexes on this table</p></td>
</tr>
<tr>
<td><p><code>idx_blks_hit</code> <code>bigint</code></p>
<p>Number of buffer hits in all indexes on this table</p></td>
</tr>
<tr>
<td><p><code>toast_blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read from this table's TOAST table (if any)</p></td>
</tr>
<tr>
<td><p><code>toast_blks_hit</code> <code>bigint</code></p>
<p>Number of buffer hits in this table's TOAST table (if any)</p></td>
</tr>
<tr>
<td><p><code>tidx_blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read from this table's TOAST table indexes (if any)</p></td>
</tr>
<tr>
<td><p><code>tidx_blks_hit</code> <code>bigint</code></p>
<p>Number of buffer hits in this table's TOAST table indexes (if any)</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-statio-all-indexes-view"></a>

### `pg_statio_all_indexes`


 The `pg_statio_all_indexes` view will contain one row for each index in the current database, showing statistics about I/O on that specific index. The `pg_statio_user_indexes` and `pg_statio_sys_indexes` views contain the same information, but filtered to only show user and system indexes respectively.
 <a id="pg-statio-all-indexes-view"></a>

**Table: `pg_statio_all_indexes` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of the table for this index</p></td>
</tr>
<tr>
<td><p><code>indexrelid</code> <code>oid</code></p>
<p>OID of this index</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code></p>
<p>Name of the schema this index is in</p></td>
</tr>
<tr>
<td><p><code>relname</code> <code>name</code></p>
<p>Name of the table for this index</p></td>
</tr>
<tr>
<td><p><code>indexrelname</code> <code>name</code></p>
<p>Name of this index</p></td>
</tr>
<tr>
<td><p><code>idx_blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read from this index</p></td>
</tr>
<tr>
<td><p><code>idx_blks_hit</code> <code>bigint</code></p>
<p>Number of buffer hits in this index</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-statio-all-sequences-view"></a>

### `pg_statio_all_sequences`


 The `pg_statio_all_sequences` view will contain one row for each sequence in the current database, showing statistics about I/O on that specific sequence.
 <a id="pg-statio-all-sequences-view"></a>

**Table: `pg_statio_all_sequences` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>relid</code> <code>oid</code></p>
<p>OID of a sequence</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code></p>
<p>Name of the schema this sequence is in</p></td>
</tr>
<tr>
<td><p><code>relname</code> <code>name</code></p>
<p>Name of this sequence</p></td>
</tr>
<tr>
<td><p><code>blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read from this sequence</p></td>
</tr>
<tr>
<td><p><code>blks_hit</code> <code>bigint</code></p>
<p>Number of buffer hits in this sequence</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-user-functions-view"></a>

### `pg_stat_user_functions`


 The `pg_stat_user_functions` view will contain one row for each tracked function, showing statistics about executions of that function. The [track_functions](../server-configuration/run-time-statistics.md#guc-track-functions) parameter controls exactly which functions are tracked.
 <a id="pg-stat-user-functions-view"></a>

**Table: `pg_stat_user_functions` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>funcid</code> <code>oid</code></p>
<p>OID of a function</p></td>
</tr>
<tr>
<td><p><code>schemaname</code> <code>name</code></p>
<p>Name of the schema this function is in</p></td>
</tr>
<tr>
<td><p><code>funcname</code> <code>name</code></p>
<p>Name of this function</p></td>
</tr>
<tr>
<td><p><code>calls</code> <code>bigint</code></p>
<p>Number of times this function has been called</p></td>
</tr>
<tr>
<td><p><code>total_time</code> <code>double precision</code></p>
<p>Total time spent in this function and all other functions called by it, in milliseconds</p></td>
</tr>
<tr>
<td><p><code>self_time</code> <code>double precision</code></p>
<p>Total time spent in this function itself, not including other functions called by it, in milliseconds</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-pg-stat-slru-view"></a>

### `pg_stat_slru`


 PostgreSQL accesses certain on-disk information via `SLRU` (*simple least-recently-used*) caches. The `pg_stat_slru` view will contain one row for each tracked SLRU cache, showing statistics about access to cached pages.


 For each `SLRU` cache that's part of the core server, there is a configuration parameter that controls its size, with the suffix `_buffers` appended.
 <a id="pg-stat-slru-view"></a>

**Table: `pg_stat_slru` View**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>name</code> <code>text</code></p>
<p>Name of the SLRU</p></td>
</tr>
<tr>
<td><p><code>blks_zeroed</code> <code>bigint</code></p>
<p>Number of blocks zeroed during initializations</p></td>
</tr>
<tr>
<td><p><code>blks_hit</code> <code>bigint</code></p>
<p>Number of times disk blocks were found already in the SLRU, so that a read was not necessary (this only includes hits in the SLRU, not the operating system's file system cache)</p></td>
</tr>
<tr>
<td><p><code>blks_read</code> <code>bigint</code></p>
<p>Number of disk blocks read for this SLRU</p></td>
</tr>
<tr>
<td><p><code>blks_written</code> <code>bigint</code></p>
<p>Number of disk blocks written for this SLRU</p></td>
</tr>
<tr>
<td><p><code>blks_exists</code> <code>bigint</code></p>
<p>Number of blocks checked for existence for this SLRU</p></td>
</tr>
<tr>
<td><p><code>flushes</code> <code>bigint</code></p>
<p>Number of flushes of dirty data for this SLRU</p></td>
</tr>
<tr>
<td><p><code>truncates</code> <code>bigint</code></p>
<p>Number of truncates for this SLRU</p></td>
</tr>
<tr>
<td><p><code>stats_reset</code> <code>timestamp with time zone</code></p>
<p>Time at which these statistics were last reset</p></td>
</tr>
</tbody>
</table>
  <a id="monitoring-stats-functions"></a>

### Statistics Functions


 Other ways of looking at the statistics can be set up by writing queries that use the same underlying statistics access functions used by the standard views shown above. For details such as the functions' names, consult the definitions of the standard views. (For example, in psql you could issue `\d+ pg_stat_activity`.) The access functions for per-database statistics take a database OID as an argument to identify which database to report on. The per-table and per-index functions take a table or index OID. The functions for per-function statistics take a function OID. Note that only tables, indexes, and functions in the current database can be seen with these functions.


 Additional functions related to the cumulative statistics system are listed in [Additional Statistics Functions](#monitoring-stats-funcs-table).
 <a id="monitoring-stats-funcs-table"></a>

**Table: Additional Statistics Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_backend_pid</code> () <code>integer</code></td>
<td>Returns the process ID of the server process attached to the current session.</td>
<td></td>
</tr>
<tr id="pg-stat-get-backend-io">
<td><code>pg_stat_get_backend_io</code> ( <code>integer</code> ) <code>setof record</code></td>
<td>Returns I/O statistics about the backend with the specified process ID. The output fields are exactly the same as the ones in the <code>pg_stat_io</code> view.</td>
<td>The function does not return I/O statistics for the checkpointer, the background writer, the startup process and the autovacuum launcher as they are already visible in the <code>pg_stat_io</code> view and there is only one of each.</td>
</tr>
<tr>
<td><code>pg_stat_get_activity</code> ( <code>integer</code> ) <code>setof record</code></td>
<td>Returns a record of information about the backend with the specified process ID, or one record for each active backend in the system if <code>NULL</code> is specified. The fields returned are a subset of those in the <code>pg_stat_activity</code> view.</td>
<td></td>
</tr>
<tr id="pg-stat-get-backend-wal">
<td><code>pg_stat_get_backend_wal</code> ( <code>integer</code> ) <code>record</code></td>
<td>Returns WAL statistics about the backend with the specified process ID. The output fields are exactly the same as the ones in the <code>pg_stat_wal</code> view.</td>
<td>The function does not return WAL statistics for the checkpointer, the background writer, the startup process and the autovacuum launcher.</td>
</tr>
<tr>
<td><code>pg_stat_get_snapshot_timestamp</code> () <code>timestamp with time zone</code></td>
<td>Returns the timestamp of the current statistics snapshot, or NULL if no statistics snapshot has been taken. A snapshot is taken the first time cumulative statistics are accessed in a transaction if <code>stats_fetch_consistency</code> is set to <code>snapshot</code></td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_xact_blocks_fetched</code> ( <code>oid</code> ) <code>bigint</code></td>
<td>Returns the number of block read requests for table or index, in the current transaction. This number minus <code>pg_stat_get_xact_blocks_hit</code> gives the number of kernel <code>read()</code> calls; the number of actual physical reads is usually lower due to kernel-level buffering.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_xact_blocks_hit</code> ( <code>oid</code> ) <code>bigint</code></td>
<td>Returns the number of block read requests for table or index, in the current transaction, found in cache (not triggering kernel <code>read()</code> calls).</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_clear_snapshot</code> () <code>void</code></td>
<td>Discards the current statistics snapshot or cached information.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_reset</code> () <code>void</code></td>
<td>Resets all statistics counters for the current database to zero.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_shared</code> ( [ <code>target</code> <code>text</code> <code>DEFAULT</code> <code>NULL</code> ] ) <code>void</code></td>
<td><p>Resets some cluster-wide statistics counters to zero, depending on the argument. <code>target</code> can be:</p>
<p>-  <code>archiver</code>: Reset all the counters shown in the <code>pg_stat_archiver</code> view. <br>
-  <code>bgwriter</code>: Reset all the counters shown in the <code>pg_stat_bgwriter</code> view. <br>
-  <code>checkpointer</code>: Reset all the counters shown in the <code>pg_stat_checkpointer</code> view. <br>
-  <code>io</code>: Reset all the counters shown in the <code>pg_stat_io</code> view. <br>
-  <code>recovery_prefetch</code>: Reset all the counters shown in the <code>pg_stat_recovery_prefetch</code> view. <br>
-  <code>slru</code>: Reset all the counters shown in the <code>pg_stat_slru</code> view. <br>
-  <code>wal</code>: Reset all the counters shown in the <code>pg_stat_wal</code> view. <br>
-  <code>NULL</code> or not specified: All the counters from the views listed above are reset.</p></td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_single_table_counters</code> ( <code>oid</code> ) <code>void</code></td>
<td>Resets statistics for a single table or index in the current database or shared across all databases in the cluster to zero.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_backend_stats</code> ( <code>integer</code> ) <code>void</code></td>
<td>Resets statistics for a single backend with the specified process ID to zero.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_single_function_counters</code> ( <code>oid</code> ) <code>void</code></td>
<td>Resets statistics for a single function in the current database to zero.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_slru</code> ( [ <code>target</code> <code>text</code> <code>DEFAULT</code> <code>NULL</code> ] ) <code>void</code></td>
<td>Resets statistics to zero for a single SLRU cache, or for all SLRUs in the cluster. If <code>target</code> is <code>NULL</code> or is not specified, all the counters shown in the <code>pg_stat_slru</code> view for all SLRU caches are reset. The argument can be one of <code>commit_timestamp</code>, <code>multixact_member</code>, <code>multixact_offset</code>, <code>notify</code>, <code>serializable</code>, <code>subtransaction</code>, or <code>transaction</code> to reset the counters for only that entry. If the argument is <code>other</code> (or indeed, any unrecognized name), then the counters for all other SLRU caches, such as extension-defined caches, are reset.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_replication_slot</code> ( <code>text</code> ) <code>void</code></td>
<td>Resets statistics of the replication slot defined by the argument. If the argument is <code>NULL</code>, resets statistics for all the replication slots.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_stat_reset_subscription_stats</code> ( <code>oid</code> ) <code>void</code></td>
<td>Resets statistics for a single subscription shown in the <code>pg_stat_subscription_stats</code> view to zero. If the argument is <code>NULL</code>, reset statistics for all subscriptions.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
</tbody>
</table>


!!! warning

    Using `pg_stat_reset()` also resets counters that autovacuum uses to determine when to trigger a vacuum or an analyze. Resetting these counters can cause autovacuum to not perform necessary work, which can cause problems such as table bloat or out-dated table statistics. A database-wide `ANALYZE` is recommended after the statistics have been reset.


 `pg_stat_get_activity`, the underlying function of the `pg_stat_activity` view, returns a set of records containing all the available information about each backend process. Sometimes it may be more convenient to obtain just a subset of this information. In such cases, another set of per-backend statistics access functions can be used; these are shown in [Per-Backend Statistics Functions](#monitoring-stats-backend-funcs-table). These access functions use the session's backend ID number, which is a small integer (>= 0) that is distinct from the backend ID of any concurrent session, although a session's ID can be recycled as soon as it exits. The backend ID is used, among other things, to identify the session's temporary schema if it has one. The function `pg_stat_get_backend_idset` provides a convenient way to list all the active backends' ID numbers for invoking these functions. For example, to show the PIDs and current queries of all backends:

```sql

SELECT pg_stat_get_backend_pid(backendid) AS pid,
       pg_stat_get_backend_activity(backendid) AS query
FROM pg_stat_get_backend_idset() AS backendid;
```

 <a id="monitoring-stats-backend-funcs-table"></a>

**Table: Per-Backend Statistics Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_stat_get_backend_activity</code> ( <code>integer</code> ) <code>text</code></td>
<td>Returns the text of this backend's most recent query.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_activity_start</code> ( <code>integer</code> ) <code>timestamp with time zone</code></td>
<td>Returns the time when the backend's most recent query was started.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_client_addr</code> ( <code>integer</code> ) <code>inet</code></td>
<td>Returns the IP address of the client connected to this backend.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_client_port</code> ( <code>integer</code> ) <code>integer</code></td>
<td>Returns the TCP port number that the client is using for communication.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_dbid</code> ( <code>integer</code> ) <code>oid</code></td>
<td>Returns the OID of the database this backend is connected to.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_idset</code> () <code>setof integer</code></td>
<td>Returns the set of currently active backend ID numbers.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_pid</code> ( <code>integer</code> ) <code>integer</code></td>
<td>Returns the process ID of this backend.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_start</code> ( <code>integer</code> ) <code>timestamp with time zone</code></td>
<td>Returns the time when this process was started.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_subxact</code> ( <code>integer</code> ) <code>record</code></td>
<td>Returns a record of information about the subtransactions of the backend with the specified ID. The fields returned are <code>subxact_count</code>, which is the number of subtransactions in the backend's subtransaction cache, and <code>subxact_overflow</code>, which indicates whether the backend's subtransaction cache is overflowed or not.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_userid</code> ( <code>integer</code> ) <code>oid</code></td>
<td>Returns the OID of the user logged into this backend.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_wait_event</code> ( <code>integer</code> ) <code>text</code></td>
<td>Returns the wait event name if this backend is currently waiting, otherwise NULL. See <a href="#wait-event-activity-table">Wait Events of Type <code>Activity</code></a> through <a href="#wait-event-timeout-table">Wait Events of Type <code>Timeout</code></a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_wait_event_type</code> ( <code>integer</code> ) <code>text</code></td>
<td>Returns the wait event type name if this backend is currently waiting, otherwise NULL. See <a href="#wait-event-table">Wait Event Types</a> for details.</td>
<td></td>
</tr>
<tr>
<td><code>pg_stat_get_backend_xact_start</code> ( <code>integer</code> ) <code>timestamp with time zone</code></td>
<td>Returns the time when the backend's current transaction was started.</td>
<td></td>
</tr>
</tbody>
</table>
