<a id="functions-admin"></a>

## System Administration Functions


 The functions described in this section are used to control and monitor a PostgreSQL installation.
 <a id="functions-admin-set"></a>

### Configuration Settings Functions


 [Configuration Settings Functions](#functions-admin-set-table) shows the functions available to query and alter run-time configuration parameters.
 <a id="functions-admin-set-table"></a>

**Table: Configuration Settings Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>current_setting</code> ( <code>setting_name</code> <code>text</code> [, <code>missing_ok</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Returns the current value of the setting <code>setting_name</code>. If there is no such setting, <code>current_setting</code> throws an error unless <code>missing_ok</code> is supplied and is <code>true</code> (in which case NULL is returned). This function corresponds to the SQL command <a href="../../reference/sql-commands/show.md#sql-show">sql-show</a>.</td>
<td><code>current_setting('datestyle')</code> <code>ISO, MDY</code></td>
</tr>
<tr>
<td><code>set_config</code> ( <code>setting_name</code> <code>text</code>, <code>new_value</code> <code>text</code>, <code>is_local</code> <code>boolean</code> ) <code>text</code></td>
<td>Sets the parameter <code>setting_name</code> to <code>new_value</code>, and returns that value. If <code>is_local</code> is <code>true</code>, the new value will only apply during the current transaction. If you want the new value to apply for the rest of the current session, use <code>false</code> instead. This function corresponds to the SQL command <a href="../../reference/sql-commands/set.md#sql-set">sql-set</a>.</td>
<td><code>set_config('log_statement_stats', 'off', false)</code> <code>off</code></td>
</tr>
</tbody>
</table>
  <a id="functions-admin-signal"></a>

### Server Signaling Functions


 The functions shown in [Server Signaling Functions](#functions-admin-signal-table) send control signals to other server processes. Use of these functions is restricted to superusers by default but access may be granted to others using `GRANT`, with noted exceptions.


 Each of these functions returns `true` if the signal was successfully sent and `false` if sending the signal failed.
 <a id="functions-admin-signal-table"></a>

**Table: Server Signaling Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_cancel_backend</code> ( <code>pid</code> <code>integer</code> ) <code>boolean</code></td>
<td>Cancels the current query of the session whose backend process has the specified process ID. This is also allowed if the calling role is a member of the role whose backend is being canceled or the calling role has privileges of <code>pg_signal_backend</code>, however only superusers can cancel superuser backends.</td>
<td></td>
</tr>
<tr>
<td><code>pg_log_backend_memory_contexts</code> ( <code>pid</code> <code>integer</code> ) <code>boolean</code></td>
<td>Requests to log the memory contexts of the backend with the specified process ID. This function can send the request to backends and auxiliary processes except logger. These memory contexts will be logged at <code>LOG</code> message level. They will appear in the server log based on the log configuration set (see <a href="../../server-administration/server-configuration/error-reporting-and-logging.md#runtime-config-logging">Error Reporting and Logging</a> for more information), but will not be sent to the client regardless of <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-client-min-messages">client_min_messages</a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_reload_conf</code> () <code>boolean</code></td>
<td>Causes all processes of the PostgreSQL server to reload their configuration files. (This is initiated by sending a <code>SIGHUP</code> signal to the postmaster process, which in turn sends <code>SIGHUP</code> to each of its children.) You can use the <a href="../../internals/system-views/pg_file_settings.md#view-pg-file-settings"><code>pg_file_settings</code></a>, <a href="../../internals/system-views/pg_hba_file_rules.md#view-pg-hba-file-rules"><code>pg_hba_file_rules</code></a> and <a href="../../internals/system-views/pg_ident_file_mappings.md#view-pg-ident-file-mappings"><code>pg_ident_file_mappings</code></a> views to check the configuration files for possible errors, before reloading.</td>
<td></td>
</tr>
<tr>
<td><code>pg_rotate_logfile</code> () <code>boolean</code></td>
<td>Signals the log-file manager to switch to a new output file immediately. This works only when the built-in log collector is running, since otherwise there is no log-file manager subprocess.</td>
<td></td>
</tr>
<tr>
<td><code>pg_terminate_backend</code> ( <code>pid</code> <code>integer</code>, <code>timeout</code> <code>bigint</code> <code>DEFAULT</code> <code>0</code> ) <code>boolean</code></td>
<td>Terminates the session whose backend process has the specified process ID. This is also allowed if the calling role is a member of the role whose backend is being terminated or the calling role has privileges of <code>pg_signal_backend</code>, however only superusers can terminate superuser backends.</td>
<td>If <code>timeout</code> is not specified or zero, this function returns <code>true</code> whether the process actually terminates or not, indicating only that the sending of the signal was successful. If the <code>timeout</code> is specified (in milliseconds) and greater than zero, the function waits until the process is actually terminated or until the given time has passed. If the process is terminated, the function returns <code>true</code>. On timeout, a warning is emitted and <code>false</code> is returned.</td>
</tr>
</tbody>
</table>


 `pg_cancel_backend` and `pg_terminate_backend` send signals (`SIGINT` or `SIGTERM` respectively) to backend processes identified by process ID. The process ID of an active backend can be found from the `pid` column of the `pg_stat_activity` view, or by listing the `postgres` processes on the server (using ps on Unix or the Task Manager on Windows). The role of an active backend can be found from the `usename` column of the `pg_stat_activity` view.


 `pg_log_backend_memory_contexts` can be used to log the memory contexts of a backend process. For example:

```

postgres=# SELECT pg_log_backend_memory_contexts(pg_backend_pid());
 pg_log_backend_memory_contexts
--------------------------------
 t
(1 row)
```
 One message for each memory context will be logged. For example:

```

LOG:  logging memory contexts of PID 10377
STATEMENT:  SELECT pg_log_backend_memory_contexts(pg_backend_pid());
LOG:  level: 0; TopMemoryContext: 80800 total in 6 blocks; 14432 free (5 chunks); 66368 used
LOG:  level: 1; pgstat TabStatusArray lookup hash table: 8192 total in 1 blocks; 1408 free (0 chunks); 6784 used
LOG:  level: 1; TopTransactionContext: 8192 total in 1 blocks; 7720 free (1 chunks); 472 used
LOG:  level: 1; RowDescriptionContext: 8192 total in 1 blocks; 6880 free (0 chunks); 1312 used
LOG:  level: 1; MessageContext: 16384 total in 2 blocks; 5152 free (0 chunks); 11232 used
LOG:  level: 1; Operator class cache: 8192 total in 1 blocks; 512 free (0 chunks); 7680 used
LOG:  level: 1; smgr relation table: 16384 total in 2 blocks; 4544 free (3 chunks); 11840 used
LOG:  level: 1; TransactionAbortContext: 32768 total in 1 blocks; 32504 free (0 chunks); 264 used
...
LOG:  level: 1; ErrorContext: 8192 total in 1 blocks; 7928 free (3 chunks); 264 used
LOG:  Grand total: 1651920 bytes in 201 blocks; 622360 free (88 chunks); 1029560 used
```
 If there are more than 100 child contexts under the same parent, the first 100 child contexts are logged, along with a summary of the remaining contexts. Note that frequent calls to this function could incur significant overhead, because it may generate a large number of log messages.
  <a id="functions-admin-backup"></a>

### Backup Control Functions


 The functions shown in [Backup Control Functions](#functions-admin-backup-table) assist in making on-line backups. These functions cannot be executed during recovery (except `pg_backup_start`, `pg_backup_stop`, and `pg_wal_lsn_diff`).


 For details about proper usage of these functions, see [Continuous Archiving and Point-in-Time Recovery (PITR)](../../server-administration/backup-and-restore/continuous-archiving-and-point-in-time-recovery-pitr.md#continuous-archiving).
 <a id="functions-admin-backup-table"></a>

**Table: Backup Control Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_create_restore_point</code> ( <code>name</code> <code>text</code> ) <code>pg_lsn</code></td>
<td>Creates a named marker record in the write-ahead log that can later be used as a recovery target, and returns the corresponding write-ahead log location. The given name can then be used with <a href="../../server-administration/server-configuration/write-ahead-log.md#guc-recovery-target-name">recovery_target_name</a> to specify the point up to which recovery will proceed. Avoid creating multiple restore points with the same name, since recovery will stop at the first one whose name matches the recovery target.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_current_wal_flush_lsn</code> () <code>pg_lsn</code></td>
<td>Returns the current write-ahead log flush location (see notes below).</td>
<td></td>
</tr>
<tr>
<td><code>pg_current_wal_insert_lsn</code> () <code>pg_lsn</code></td>
<td>Returns the current write-ahead log insert location (see notes below).</td>
<td></td>
</tr>
<tr>
<td><code>pg_current_wal_lsn</code> () <code>pg_lsn</code></td>
<td>Returns the current write-ahead log write location (see notes below).</td>
<td></td>
</tr>
<tr>
<td><code>pg_backup_start</code> ( <code>label</code> <code>text</code> [, <code>fast</code> <code>boolean</code> ] ) <code>pg_lsn</code></td>
<td>Prepares the server to begin an on-line backup. The only required parameter is an arbitrary user-defined label for the backup. (Typically this would be the name under which the backup dump file will be stored.) If the optional second parameter is given as <code>true</code>, it specifies executing <code>pg_backup_start</code> as quickly as possible. This forces an immediate checkpoint which will cause a spike in I/O operations, slowing any concurrently executing queries.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_backup_stop</code> ( [<code>wait_for_archive</code> <code>boolean</code> ] ) <code>record</code> ( <code>lsn</code> <code>pg_lsn</code>, <code>labelfile</code> <code>text</code>, <code>spcmapfile</code> <code>text</code> )</td>
<td>Finishes performing an on-line backup. The desired contents of the backup label file and the tablespace map file are returned as part of the result of the function and must be written to files in the backup area. These files must not be written to the live data directory (doing so will cause PostgreSQL to fail to restart in the event of a crash).</td>
<td>There is an optional parameter of type <code>boolean</code>. If false, the function will return immediately after the backup is completed, without waiting for WAL to be archived. This behavior is only useful with backup software that independently monitors WAL archiving. Otherwise, WAL required to make the backup consistent might be missing and make the backup useless. By default or when this parameter is true, <code>pg_backup_stop</code> will wait for WAL to be archived when archiving is enabled. (On a standby, this means that it will wait only when <code>archive_mode</code> = <code>always</code>. If write activity on the primary is low, it may be useful to run <code>pg_switch_wal</code> on the primary in order to trigger an immediate segment switch.)<br>When executed on a primary, this function also creates a backup history file in the write-ahead log archive area. The history file includes the label given to <code>pg_backup_start</code>, the starting and ending write-ahead log locations for the backup, and the starting and ending times of the backup. After recording the ending location, the current write-ahead log insertion point is automatically advanced to the next write-ahead log file, so that the ending write-ahead log file can be archived immediately to complete the backup.<br>The result of the function is a single record. The <code>lsn</code> column holds the backup's ending write-ahead log location (which again can be ignored). The second column returns the contents of the backup label file, and the third column returns the contents of the tablespace map file. These must be stored as part of the backup and are required as part of the restore process.<br>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_switch_wal</code> () <code>pg_lsn</code></td>
<td>Forces the server to switch to a new write-ahead log file, which allows the current file to be archived (assuming you are using continuous archiving). The result is the ending write-ahead log location plus 1 within the just-completed write-ahead log file. If there has been no write-ahead log activity since the last write-ahead log switch, <code>pg_switch_wal</code> does nothing and returns the start location of the write-ahead log file currently in use.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_walfile_name</code> ( <code>lsn</code> <code>pg_lsn</code> ) <code>text</code></td>
<td>Converts a write-ahead log location to the name of the WAL file holding that location.</td>
<td></td>
</tr>
<tr>
<td><code>pg_walfile_name_offset</code> ( <code>lsn</code> <code>pg_lsn</code> ) <code>record</code> ( <code>file_name</code> <code>text</code>, <code>file_offset</code> <code>integer</code> )</td>
<td>Converts a write-ahead log location to a WAL file name and byte offset within that file.</td>
<td></td>
</tr>
<tr>
<td><code>pg_split_walfile_name</code> ( <code>file_name</code> <code>text</code> ) <code>record</code> ( <code>segment_number</code> <code>numeric</code>, <code>timeline_id</code> <code>bigint</code> )</td>
<td>Extracts the sequence number and timeline ID from a WAL file name.</td>
<td></td>
</tr>
<tr>
<td><code>pg_wal_lsn_diff</code> ( <code>lsn1</code> <code>pg_lsn</code>, <code>lsn2</code> <code>pg_lsn</code> ) <code>numeric</code></td>
<td>Calculates the difference in bytes (<code>lsn1</code> - <code>lsn2</code>) between two write-ahead log locations. This can be used with <code>pg_stat_replication</code> or some of the functions shown in <a href="#functions-admin-backup-table">Backup Control Functions</a> to get the replication lag.</td>
<td></td>
</tr>
</tbody>
</table>


 `pg_current_wal_lsn` displays the current write-ahead log write location in the same format used by the above functions. Similarly, `pg_current_wal_insert_lsn` displays the current write-ahead log insertion location and `pg_current_wal_flush_lsn` displays the current write-ahead log flush location. The insertion location is the “logical” end of the write-ahead log at any instant, while the write location is the end of what has actually been written out from the server's internal buffers, and the flush location is the last location known to be written to durable storage. The write location is the end of what can be examined from outside the server, and is usually what you want if you are interested in archiving partially-complete write-ahead log files. The insertion and flush locations are made available primarily for server debugging purposes. These are all read-only operations and do not require superuser permissions.


 You can use `pg_walfile_name_offset` to extract the corresponding write-ahead log file name and byte offset from a `pg_lsn` value. For example:

```

postgres=# SELECT * FROM pg_walfile_name_offset((pg_backup_stop()).lsn);
        file_name         | file_offset
--------------------------+-------------
 00000001000000000000000D |     4039624
(1 row)
```
 Similarly, `pg_walfile_name` extracts just the write-ahead log file name.


 `pg_split_walfile_name` is useful to compute a LSN from a file offset and WAL file name, for example:

```

postgres=# \set file_name '000000010000000100C000AB'
postgres=# \set offset 256
postgres=# SELECT '0/0'::pg_lsn + pd.segment_number * ps.setting::int + :offset AS lsn
  FROM pg_split_walfile_name(:'file_name') pd,
       pg_show_all_settings() ps
  WHERE ps.name = 'wal_segment_size';
      lsn
---------------
 C001/AB000100
(1 row)
```

  <a id="functions-recovery-control"></a>

### Recovery Control Functions


 The functions shown in [Recovery Information Functions](#functions-recovery-info-table) provide information about the current status of a standby server. These functions may be executed both during recovery and in normal running.
 <a id="functions-recovery-info-table"></a>

**Table: Recovery Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_is_in_recovery</code> () <code>boolean</code></td>
<td>Returns true if recovery is still in progress.</td>
<td></td>
</tr>
<tr>
<td><code>pg_last_wal_receive_lsn</code> () <code>pg_lsn</code></td>
<td>Returns the last write-ahead log location that has been received and synced to disk by streaming replication. While streaming replication is in progress this will increase monotonically. If recovery has completed then this will remain static at the location of the last WAL record received and synced to disk during recovery. If streaming replication is disabled, or if it has not yet started, the function returns <code>NULL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_last_wal_replay_lsn</code> () <code>pg_lsn</code></td>
<td>Returns the last write-ahead log location that has been replayed during recovery. If recovery is still in progress this will increase monotonically. If recovery has completed then this will remain static at the location of the last WAL record applied during recovery. When the server has been started normally without recovery, the function returns <code>NULL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_last_xact_replay_timestamp</code> () <code>timestamp with time zone</code></td>
<td>Returns the time stamp of the last transaction replayed during recovery. This is the time at which the commit or abort WAL record for that transaction was generated on the primary. If no transactions have been replayed during recovery, the function returns <code>NULL</code>. Otherwise, if recovery is still in progress this will increase monotonically. If recovery has completed then this will remain static at the time of the last transaction applied during recovery. When the server has been started normally without recovery, the function returns <code>NULL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_wal_resource_managers</code> () <code>setof record</code> ( <code>rm_id</code> <code>integer</code>, <code>rm_name</code> <code>text</code>, <code>rm_builtin</code> <code>boolean</code> )</td>
<td>Returns the currently-loaded WAL resource managers in the system. The column <code>rm_builtin</code> indicates whether it's a built-in resource manager, or a custom resource manager loaded by an extension.</td>
<td></td>
</tr>
</tbody>
</table>


 The functions shown in [Recovery Control Functions](#functions-recovery-control-table) control the progress of recovery. These functions may be executed only during recovery.
 <a id="functions-recovery-control-table"></a>

**Table: Recovery Control Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_is_wal_replay_paused</code> () <code>boolean</code></td>
<td>Returns true if recovery pause is requested.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_wal_replay_pause_state</code> () <code>text</code></td>
<td>Returns recovery pause state. The return values are <code> not paused</code> if pause is not requested, <code> pause requested</code> if pause is requested but recovery is not yet paused, and <code>paused</code> if the recovery is actually paused.</td>
<td></td>
</tr>
<tr>
<td><code>pg_promote</code> ( <code>wait</code> <code>boolean</code> <code>DEFAULT</code> <code>true</code>, <code>wait_seconds</code> <code>integer</code> <code>DEFAULT</code> <code>60</code> ) <code>boolean</code></td>
<td>Promotes a standby server to primary status. With <code>wait</code> set to <code>true</code> (the default), the function waits until promotion is completed or <code>wait_seconds</code> seconds have passed, and returns <code>true</code> if promotion is successful and <code>false</code> otherwise. If <code>wait</code> is set to <code>false</code>, the function returns <code>true</code> immediately after sending a <code>SIGUSR1</code> signal to the postmaster to trigger promotion.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_wal_replay_pause</code> () <code>void</code></td>
<td>Request to pause recovery. A request doesn't mean that recovery stops right away. If you want a guarantee that recovery is actually paused, you need to check for the recovery pause state returned by <code>pg_get_wal_replay_pause_state()</code>. Note that <code>pg_is_wal_replay_paused()</code> returns whether a request is made. While recovery is paused, no further database changes are applied. If hot standby is active, all new queries will see the same consistent snapshot of the database, and no further query conflicts will be generated until recovery is resumed.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_wal_replay_resume</code> () <code>void</code></td>
<td>Restarts recovery if it was paused.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
</tbody>
</table>


 `pg_wal_replay_pause` and `pg_wal_replay_resume` cannot be executed while a promotion is ongoing. If a promotion is triggered while recovery is paused, the paused state ends and promotion continues.


 If streaming replication is disabled, the paused state may continue indefinitely without a problem. If streaming replication is in progress then WAL records will continue to be received, which will eventually fill available disk space, depending upon the duration of the pause, the rate of WAL generation and available disk space.
  <a id="functions-snapshot-synchronization"></a>

### Snapshot Synchronization Functions


 PostgreSQL allows database sessions to synchronize their snapshots. A *snapshot* determines which data is visible to the transaction that is using the snapshot. Synchronized snapshots are necessary when two or more sessions need to see identical content in the database. If two sessions just start their transactions independently, there is always a possibility that some third transaction commits between the executions of the two `START TRANSACTION` commands, so that one session sees the effects of that transaction and the other does not.


 To solve this problem, PostgreSQL allows a transaction to *export* the snapshot it is using. As long as the exporting transaction remains open, other transactions can *import* its snapshot, and thereby be guaranteed that they see exactly the same view of the database that the first transaction sees. But note that any database changes made by any one of these transactions remain invisible to the other transactions, as is usual for changes made by uncommitted transactions. So the transactions are synchronized with respect to pre-existing data, but act normally for changes they make themselves.


 Snapshots are exported with the `pg_export_snapshot` function, shown in [Snapshot Synchronization Functions](#functions-snapshot-synchronization-table), and imported with the [sql-set-transaction](../../reference/sql-commands/set-transaction.md#sql-set-transaction) command.
 <a id="functions-snapshot-synchronization-table"></a>

**Table: Snapshot Synchronization Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_export_snapshot</code> () <code>text</code></td>
<td>Saves the transaction's current snapshot and returns a <code>text</code> string identifying the snapshot. This string must be passed (outside the database) to clients that want to import the snapshot. The snapshot is available for import only until the end of the transaction that exported it.</td>
<td>A transaction can export more than one snapshot, if needed. Note that doing so is only useful in <code>READ COMMITTED</code> transactions, since in <code>REPEATABLE READ</code> and higher isolation levels, transactions use the same snapshot throughout their lifetime. Once a transaction has exported any snapshots, it cannot be prepared with <a href="../../reference/sql-commands/prepare-transaction.md#sql-prepare-transaction">sql-prepare-transaction</a>.</td>
</tr>
<tr>
<td><code>pg_log_standby_snapshot</code> () <code>pg_lsn</code></td>
<td>Take a snapshot of running transactions and write it to WAL, without having to wait for bgwriter or checkpointer to log one. This is useful for logical decoding on standby, as logical slot creation has to wait until such a record is replayed on the standby.</td>
<td></td>
</tr>
</tbody>
</table>
  <a id="functions-replication"></a>

### Replication Management Functions


 The functions shown in [Replication Management Functions](#functions-replication-table) are for controlling and interacting with replication features. See [Streaming Replication](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication), [Replication Slots](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication-slots), and [Replication Progress Tracking](../../server-programming/replication-progress-tracking.md#replication-origins) for information about the underlying features. Use of functions for replication origin is only allowed to the superuser by default, but may be allowed to other users by using the `GRANT` command. Use of functions for replication slots is restricted to superusers and users having `REPLICATION` privilege.


 Many of these functions have equivalent commands in the replication protocol; see [Streaming Replication Protocol](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication).


 The functions described in [Backup Control Functions](#functions-admin-backup), [Recovery Control Functions](#functions-recovery-control), and [Snapshot Synchronization Functions](#functions-snapshot-synchronization) are also relevant for replication.
 <a id="functions-replication-table"></a>

**Table: Replication Management Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_create_physical_replication_slot</code> ( <code>slot_name</code> <code>name</code> [, <code>immediately_reserve</code> <code>boolean</code>, <code>temporary</code> <code>boolean</code> ] ) <code>record</code> ( <code>slot_name</code> <code>name</code>, <code>lsn</code> <code>pg_lsn</code> )</td>
<td>Creates a new physical replication slot named <code>slot_name</code>. The optional second parameter, when <code>true</code>, specifies that the LSN for this replication slot be reserved immediately; otherwise the LSN is reserved on first connection from a streaming replication client. Streaming changes from a physical slot is only possible with the streaming-replication protocol — see <a href="../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication">Streaming Replication Protocol</a>. The optional third parameter, <code>temporary</code>, when set to true, specifies that the slot should not be permanently stored to disk and is only meant for use by the current session. Temporary slots are also released upon any error. This function corresponds to the replication protocol command <code>CREATE_REPLICATION_SLOT ... PHYSICAL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_drop_replication_slot</code> ( <code>slot_name</code> <code>name</code> ) <code>void</code></td>
<td>Drops the physical or logical replication slot named <code>slot_name</code>. Same as replication protocol command <code>DROP_REPLICATION_SLOT</code>.</td>
<td></td>
</tr>
<tr id="pg-create-logical-replication-slot">
<td><code>pg_create_logical_replication_slot</code> ( <code>slot_name</code> <code>name</code>, <code>plugin</code> <code>name</code> [, <code>temporary</code> <code>boolean</code>, <code>twophase</code> <code>boolean</code>, <code>failover</code> <code>boolean</code> ] ) <code>record</code> ( <code>slot_name</code> <code>name</code>, <code>lsn</code> <code>pg_lsn</code> )</td>
<td>Creates a new logical (decoding) replication slot named <code>slot_name</code> using the output plugin <code>plugin</code>. The optional third parameter, <code>temporary</code>, when set to true, specifies that the slot should not be permanently stored to disk and is only meant for use by the current session. Temporary slots are also released upon any error. The optional fourth parameter, <code>twophase</code>, when set to true, specifies that the decoding of prepared transactions is enabled for this slot. The optional fifth parameter, <code>failover</code>, when set to true, specifies that this slot is enabled to be synced to the standbys so that logical replication can be resumed after failover. A call to this function has the same effect as the replication protocol command <code>CREATE_REPLICATION_SLOT ... LOGICAL</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_copy_physical_replication_slot</code> ( <code>src_slot_name</code> <code>name</code>, <code>dst_slot_name</code> <code>name</code> [, <code>temporary</code> <code>boolean</code> ] ) <code>record</code> ( <code>slot_name</code> <code>name</code>, <code>lsn</code> <code>pg_lsn</code> )</td>
<td>Copies an existing physical replication slot named <code>src_slot_name</code> to a physical replication slot named <code>dst_slot_name</code>. The copied physical slot starts to reserve WAL from the same LSN as the source slot. <code>temporary</code> is optional. If <code>temporary</code> is omitted, the same value as the source slot is used. Copy of an invalidated slot is not allowed.</td>
<td></td>
</tr>
<tr>
<td><code>pg_copy_logical_replication_slot</code> ( <code>src_slot_name</code> <code>name</code>, <code>dst_slot_name</code> <code>name</code> [, <code>temporary</code> <code>boolean</code> [, <code>plugin</code> <code>name</code> ]] ) <code>record</code> ( <code>slot_name</code> <code>name</code>, <code>lsn</code> <code>pg_lsn</code> )</td>
<td>Copies an existing logical replication slot named <code>src_slot_name</code> to a logical replication slot named <code>dst_slot_name</code>, optionally changing the output plugin and persistence. The copied logical slot starts from the same LSN as the source logical slot. Both <code>temporary</code> and <code>plugin</code> are optional; if they are omitted, the values of the source slot are used. The <code>failover</code> option of the source logical slot is not copied and is set to <code>false</code> by default. This is to avoid the risk of being unable to continue logical replication after failover to standby where the slot is being synchronized. Copy of an invalidated slot is not allowed.</td>
<td></td>
</tr>
<tr id="pg-logical-slot-get-changes">
<td><code>pg_logical_slot_get_changes</code> ( <code>slot_name</code> <code>name</code>, <code>upto_lsn</code> <code>pg_lsn</code>, <code>upto_nchanges</code> <code>integer</code>, <code>VARIADIC</code> <code>options</code> <code>text[]</code> ) <code>setof record</code> ( <code>lsn</code> <code>pg_lsn</code>, <code>xid</code> <code>xid</code>, <code>data</code> <code>text</code> )</td>
<td>Returns changes in the slot <code>slot_name</code>, starting from the point from which changes have been consumed last. If <code>upto_lsn</code> and <code>upto_nchanges</code> are NULL, logical decoding will continue until end of WAL. If <code>upto_lsn</code> is non-NULL, decoding will include only those transactions which commit prior to the specified LSN. If <code>upto_nchanges</code> is non-NULL, decoding will stop when the number of rows produced by decoding exceeds the specified value. Note, however, that the actual number of rows returned may be larger, since this limit is only checked after adding the rows produced when decoding each new transaction commit. If the specified slot is a logical failover slot then the function will not return until all physical slots specified in <a href="../../server-administration/server-configuration/replication.md#guc-synchronized-standby-slots"><code>synchronized_standby_slots</code></a> have confirmed WAL receipt.</td>
<td></td>
</tr>
<tr id="pg-logical-slot-peek-changes">
<td><code>pg_logical_slot_peek_changes</code> ( <code>slot_name</code> <code>name</code>, <code>upto_lsn</code> <code>pg_lsn</code>, <code>upto_nchanges</code> <code>integer</code>, <code>VARIADIC</code> <code>options</code> <code>text[]</code> ) <code>setof record</code> ( <code>lsn</code> <code>pg_lsn</code>, <code>xid</code> <code>xid</code>, <code>data</code> <code>text</code> )</td>
<td>Behaves just like the <code>pg_logical_slot_get_changes()</code> function, except that changes are not consumed; that is, they will be returned again on future calls.</td>
<td></td>
</tr>
<tr id="pg-logical-slot-get-binary-changes">
<td><code>pg_logical_slot_get_binary_changes</code> ( <code>slot_name</code> <code>name</code>, <code>upto_lsn</code> <code>pg_lsn</code>, <code>upto_nchanges</code> <code>integer</code>, <code>VARIADIC</code> <code>options</code> <code>text[]</code> ) <code>setof record</code> ( <code>lsn</code> <code>pg_lsn</code>, <code>xid</code> <code>xid</code>, <code>data</code> <code>bytea</code> )</td>
<td>Behaves just like the <code>pg_logical_slot_get_changes()</code> function, except that changes are returned as <code>bytea</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_logical_slot_peek_binary_changes</code> ( <code>slot_name</code> <code>name</code>, <code>upto_lsn</code> <code>pg_lsn</code>, <code>upto_nchanges</code> <code>integer</code>, <code>VARIADIC</code> <code>options</code> <code>text[]</code> ) <code>setof record</code> ( <code>lsn</code> <code>pg_lsn</code>, <code>xid</code> <code>xid</code>, <code>data</code> <code>bytea</code> )</td>
<td>Behaves just like the <code>pg_logical_slot_peek_changes()</code> function, except that changes are returned as <code>bytea</code>.</td>
<td></td>
</tr>
<tr id="pg-replication-slot-advance">
<td><code>pg_replication_slot_advance</code> ( <code>slot_name</code> <code>name</code>, <code>upto_lsn</code> <code>pg_lsn</code> ) <code>record</code> ( <code>slot_name</code> <code>name</code>, <code>end_lsn</code> <code>pg_lsn</code> )</td>
<td>Advances the current confirmed position of a replication slot named <code>slot_name</code>. The slot will not be moved backwards, and it will not be moved beyond the current insert location. Returns the name of the slot and the actual position that it was advanced to. The updated slot position information is written out at the next checkpoint if any advancing is done. So in the event of a crash, the slot may return to an earlier position. If the specified slot is a logical failover slot then the function will not return until all physical slots specified in <a href="../../server-administration/server-configuration/replication.md#guc-synchronized-standby-slots"><code>synchronized_standby_slots</code></a> have confirmed WAL receipt.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-create">
<td><code>pg_replication_origin_create</code> ( <code>node_name</code> <code>text</code> ) <code>oid</code></td>
<td>Creates a replication origin with the given external name, and returns the internal ID assigned to it.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-drop">
<td><code>pg_replication_origin_drop</code> ( <code>node_name</code> <code>text</code> ) <code>void</code></td>
<td>Deletes a previously-created replication origin, including any associated replay progress.</td>
<td></td>
</tr>
<tr>
<td><code>pg_replication_origin_oid</code> ( <code>node_name</code> <code>text</code> ) <code>oid</code></td>
<td>Looks up a replication origin by name and returns the internal ID. If no such replication origin is found, <code>NULL</code> is returned.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-session-setup">
<td><code>pg_replication_origin_session_setup</code> ( <code>node_name</code> <code>text</code> ) <code>void</code></td>
<td>Marks the current session as replaying from the given origin, allowing replay progress to be tracked. Can only be used if no origin is currently selected. Use <code>pg_replication_origin_session_reset</code> to undo.</td>
<td></td>
</tr>
<tr>
<td><code>pg_replication_origin_session_reset</code> () <code>void</code></td>
<td>Cancels the effects of <code>pg_replication_origin_session_setup()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_replication_origin_session_is_setup</code> () <code>boolean</code></td>
<td>Returns true if a replication origin has been selected in the current session.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-session-progress">
<td><code>pg_replication_origin_session_progress</code> ( <code>flush</code> <code>boolean</code> ) <code>pg_lsn</code></td>
<td>Returns the replay location for the replication origin selected in the current session. The parameter <code>flush</code> determines whether the corresponding local transaction will be guaranteed to have been flushed to disk or not.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-xact-setup">
<td><code>pg_replication_origin_xact_setup</code> ( <code>origin_lsn</code> <code>pg_lsn</code>, <code>origin_timestamp</code> <code>timestamp with time zone</code> ) <code>void</code></td>
<td>Marks the current transaction as replaying a transaction that has committed at the given LSN and timestamp. Can only be called when a replication origin has been selected using <code>pg_replication_origin_session_setup</code>.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-xact-reset">
<td><code>pg_replication_origin_xact_reset</code> () <code>void</code></td>
<td>Cancels the effects of <code>pg_replication_origin_xact_setup()</code>.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-advance">
<td><code>pg_replication_origin_advance</code> ( <code>node_name</code> <code>text</code>, <code>lsn</code> <code>pg_lsn</code> ) <code>void</code></td>
<td>Sets replication progress for the given node to the given location. This is primarily useful for setting up the initial location, or setting a new location after configuration changes and similar. Be aware that careless use of this function can lead to inconsistently replicated data.</td>
<td></td>
</tr>
<tr id="pg-replication-origin-progress">
<td><code>pg_replication_origin_progress</code> ( <code>node_name</code> <code>text</code>, <code>flush</code> <code>boolean</code> ) <code>pg_lsn</code></td>
<td>Returns the replay location for the given replication origin. The parameter <code>flush</code> determines whether the corresponding local transaction will be guaranteed to have been flushed to disk or not.</td>
<td></td>
</tr>
<tr id="pg-logical-emit-message">
<td><code>pg_logical_emit_message</code> ( <code>transactional</code> <code>boolean</code>, <code>prefix</code> <code>text</code>, <code>content</code> <code>text</code> [, <code>flush</code> <code>boolean</code> <code>DEFAULT</code> <code>false</code>] ) <code>pg_lsn</code></td>
<td><code>pg_logical_emit_message</code> ( <code>transactional</code> <code>boolean</code>, <code>prefix</code> <code>text</code>, <code>content</code> <code>bytea</code> [, <code>flush</code> <code>boolean</code> <code>DEFAULT</code> <code>false</code>] ) <code>pg_lsn</code></td>
<td>Emits a logical decoding message. This can be used to pass generic messages to logical decoding plugins through WAL. The <code>transactional</code> parameter specifies if the message should be part of the current transaction, or if it should be written immediately and decoded as soon as the logical decoder reads the record. The <code>prefix</code> parameter is a textual prefix that can be used by logical decoding plugins to easily recognize messages that are interesting for them. The <code>content</code> parameter is the content of the message, given either in text or binary form. The <code>flush</code> parameter (default set to <code>false</code>) controls if the message is immediately flushed to WAL or not. <code>flush</code> has no effect with <code>transactional</code>, as the message's WAL record is flushed along with its transaction.</td>
</tr>
<tr id="pg-sync-replication-slots">
<td><code>pg_sync_replication_slots</code> () <code>void</code></td>
<td>Synchronize the logical failover replication slots from the primary server to the standby server. This function can only be executed on the standby server. Temporary synced slots, if any, cannot be used for logical decoding and must be dropped after promotion. See <a href="../../server-programming/logical-decoding/logical-decoding-concepts.md#logicaldecoding-replication-slots-synchronization">Replication Slot Synchronization</a> for details. Note that this function is primarily intended for testing and debugging purposes and should be used with caution. Additionally, this function cannot be executed if <a href="../../server-administration/server-configuration/replication.md#guc-sync-replication-slots"><code> sync_replication_slots</code></a> is enabled and the slotsync worker is already running to perform the synchronization of slots.</td>
<td></td>
</tr>
</tbody>
</table>
  <a id="functions-admin-dbobject"></a>

### Database Object Management Functions


 The functions shown in [Database Object Size Functions](#functions-admin-dbsize) calculate the disk space usage of database objects, or assist in presentation or understanding of usage results. `bigint` results are measured in bytes. If an OID that does not represent an existing object is passed to one of these functions, `NULL` is returned.
 <a id="functions-admin-dbsize"></a>

**Table: Database Object Size Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_column_size</code> ( <code>"any"</code> ) <code>integer</code></td>
<td>Shows the number of bytes used to store any individual data value. If applied directly to a table column value, this reflects any compression that was done.</td>
<td></td>
</tr>
<tr>
<td><code>pg_column_compression</code> ( <code>"any"</code> ) <code>text</code></td>
<td>Shows the compression algorithm that was used to compress an individual variable-length value. Returns <code>NULL</code> if the value is not compressed.</td>
<td></td>
</tr>
<tr>
<td><code>pg_column_toast_chunk_id</code> ( <code>"any"</code> ) <code>oid</code></td>
<td>Shows the <code>chunk_id</code> of an on-disk TOASTed value. Returns <code>NULL</code> if the value is un-TOASTed or not on-disk. See <a href="../../internals/database-physical-storage/toast.md#storage-toast">TOAST</a> for more information about TOAST.</td>
<td></td>
</tr>
<tr>
<td><code>pg_database_size</code> ( <code>name</code> ) <code>bigint</code></td>
<td><code>pg_database_size</code> ( <code>oid</code> ) <code>bigint</code></td>
<td>Computes the total disk space used by the database with the specified name or OID. To use this function, you must have <code>CONNECT</code> privilege on the specified database (which is granted by default) or have privileges of the <code>pg_read_all_stats</code> role.</td>
</tr>
<tr>
<td><code>pg_indexes_size</code> ( <code>regclass</code> ) <code>bigint</code></td>
<td>Computes the total disk space used by indexes attached to the specified table.</td>
<td></td>
</tr>
<tr>
<td><code>pg_relation_size</code> ( <code>relation</code> <code>regclass</code> [, <code>fork</code> <code>text</code> ] ) <code>bigint</code></td>
<td><p>Computes the disk space used by one “fork” of the specified relation. (Note that for most purposes it is more convenient to use the higher-level functions <code>pg_total_relation_size</code> or <code>pg_table_size</code>, which sum the sizes of all forks.) With one argument, this returns the size of the main data fork of the relation. The second argument can be provided to specify which fork to examine:</p>
<p>-  <code>main</code> returns the size of the main data fork of the relation. <br>
-  <code>fsm</code> returns the size of the Free Space Map (see <a href="../../internals/database-physical-storage/free-space-map.md#storage-fsm">Free Space Map</a>) associated with the relation. <br>
-  <code>vm</code> returns the size of the Visibility Map (see <a href="../../internals/database-physical-storage/visibility-map.md#storage-vm">Visibility Map</a>) associated with the relation. <br>
-  <code>init</code> returns the size of the initialization fork, if any, associated with the relation.</p></td>
<td></td>
</tr>
<tr>
<td><code>pg_size_bytes</code> ( <code>text</code> ) <code>bigint</code></td>
<td>Converts a size in human-readable format (as returned by <code>pg_size_pretty</code>) into bytes. Valid units are <code>bytes</code>, <code>B</code>, <code>kB</code>, <code>MB</code>, <code>GB</code>, <code>TB</code>, and <code>PB</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_size_pretty</code> ( <code>bigint</code> ) <code>text</code></td>
<td><code>pg_size_pretty</code> ( <code>numeric</code> ) <code>text</code></td>
<td>Converts a size in bytes into a more easily human-readable format with size units (bytes, kB, MB, GB, TB, or PB as appropriate). Note that the units are powers of 2 rather than powers of 10, so 1kB is 1024 bytes, 1MB is 1024<sup>2</sup> = 1048576 bytes, and so on.</td>
</tr>
<tr>
<td><code>pg_table_size</code> ( <code>regclass</code> ) <code>bigint</code></td>
<td>Computes the disk space used by the specified table, excluding indexes (but including its TOAST table if any, free space map, and visibility map).</td>
<td></td>
</tr>
<tr>
<td><code>pg_tablespace_size</code> ( <code>name</code> ) <code>bigint</code></td>
<td><code>pg_tablespace_size</code> ( <code>oid</code> ) <code>bigint</code></td>
<td>Computes the total disk space used in the tablespace with the specified name or OID. To use this function, you must have <code>CREATE</code> privilege on the specified tablespace or have privileges of the <code>pg_read_all_stats</code> role, unless it is the default tablespace for the current database.</td>
</tr>
<tr>
<td><code>pg_total_relation_size</code> ( <code>regclass</code> ) <code>bigint</code></td>
<td>Computes the total disk space used by the specified table, including all indexes and TOAST data. The result is equivalent to <code>pg_table_size</code> <code>+</code> <code>pg_indexes_size</code>.</td>
<td></td>
</tr>
</tbody>
</table>


 The functions above that operate on tables or indexes accept a `regclass` argument, which is simply the OID of the table or index in the `pg_class` system catalog. You do not have to look up the OID by hand, however, since the `regclass` data type's input converter will do the work for you. See [Object Identifier Types](../data-types/object-identifier-types.md#datatype-oid) for details.


 The functions shown in [Database Object Location Functions](#functions-admin-dblocation) assist in identifying the specific disk files associated with database objects.
 <a id="functions-admin-dblocation"></a>

**Table: Database Object Location Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_relation_filenode</code> ( <code>relation</code> <code>regclass</code> ) <code>oid</code></td>
<td>Returns the “filenode” number currently assigned to the specified relation. The filenode is the base component of the file name(s) used for the relation (see <a href="../../internals/database-physical-storage/database-file-layout.md#storage-file-layout">Database File Layout</a> for more information). For most relations the result is the same as <code>pg_class</code>.<code>relfilenode</code>, but for certain system catalogs <code>relfilenode</code> is zero and this function must be used to get the correct value. The function returns NULL if passed a relation that does not have storage, such as a view.</td>
<td></td>
</tr>
<tr>
<td><code>pg_relation_filepath</code> ( <code>relation</code> <code>regclass</code> ) <code>text</code></td>
<td>Returns the entire file path name (relative to the database cluster's data directory, <code>PGDATA</code>) of the relation.</td>
<td></td>
</tr>
<tr>
<td><code>pg_filenode_relation</code> ( <code>tablespace</code> <code>oid</code>, <code>filenode</code> <code>oid</code> ) <code>regclass</code></td>
<td>Returns a relation's OID given the tablespace OID and filenode it is stored under. This is essentially the inverse mapping of <code>pg_relation_filepath</code>. For a relation in the database's default tablespace, the tablespace can be specified as zero. Returns <code>NULL</code> if no relation in the current database is associated with the given values, or if dealing with a temporary relation.</td>
<td></td>
</tr>
</tbody>
</table>


 [Collation Management Functions](#functions-admin-collation) lists functions used to manage collations.
 <a id="functions-admin-collation"></a>

**Table: Collation Management Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_collation_actual_version</code> ( <code>oid</code> ) <code>text</code></td>
<td>Returns the actual version of the collation object as it is currently installed in the operating system. If this is different from the value in <code>pg_collation</code>.<code>collversion</code>, then objects depending on the collation might need to be rebuilt. See also <a href="../../reference/sql-commands/alter-collation.md#sql-altercollation">sql-altercollation</a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_database_collation_actual_version</code> ( <code>oid</code> ) <code>text</code></td>
<td>Returns the actual version of the database's collation as it is currently installed in the operating system. If this is different from the value in <code>pg_database</code>.<code>datcollversion</code>, then objects depending on the collation might need to be rebuilt. See also <a href="../../reference/sql-commands/alter-database.md#sql-alterdatabase">sql-alterdatabase</a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_import_system_collations</code> ( <code>schema</code> <code>regnamespace</code> ) <code>integer</code></td>
<td>Adds collations to the system catalog <code>pg_collation</code> based on all the locales it finds in the operating system. This is what <code>initdb</code> uses; see <a href="../../server-administration/localization/collation-support.md#collation-managing">Managing Collations</a> for more details. If additional locales are installed into the operating system later on, this function can be run again to add collations for the new locales. Locales that match existing entries in <code>pg_collation</code> will be skipped. (But collation objects based on locales that are no longer present in the operating system are not removed by this function.) The <code>schema</code> parameter would typically be <code>pg_catalog</code>, but that is not a requirement; the collations could be installed into some other schema as well. The function returns the number of new collation objects it created. Use of this function is restricted to superusers.</td>
<td></td>
</tr>
</tbody>
</table>


 [Partitioning Information Functions](#functions-info-partition) lists functions that provide information about the structure of partitioned tables.
 <a id="functions-info-partition"></a>

**Table: Partitioning Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_partition_tree</code> ( <code>regclass</code> ) <code>setof record</code> ( <code>relid</code> <code>regclass</code>, <code>parentrelid</code> <code>regclass</code>, <code>isleaf</code> <code>boolean</code>, <code>level</code> <code>integer</code> )</td>
<td>Lists the tables or indexes in the partition tree of the given partitioned table or partitioned index, with one row for each partition. Information provided includes the OID of the partition, the OID of its immediate parent, a boolean value telling if the partition is a leaf, and an integer telling its level in the hierarchy. The level value is 0 for the input table or index, 1 for its immediate child partitions, 2 for their partitions, and so on. Returns no rows if the relation does not exist or is not a partition or partitioned table.</td>
<td></td>
</tr>
<tr>
<td><code>pg_partition_ancestors</code> ( <code>regclass</code> ) <code>setof regclass</code></td>
<td>Lists the ancestor relations of the given partition, including the relation itself. Returns no rows if the relation does not exist or is not a partition or partitioned table.</td>
<td></td>
</tr>
<tr>
<td><code>pg_partition_root</code> ( <code>regclass</code> ) <code>regclass</code></td>
<td>Returns the top-most parent of the partition tree to which the given relation belongs. Returns <code>NULL</code> if the relation does not exist or is not a partition or partitioned table.</td>
<td></td>
</tr>
</tbody>
</table>


 For example, to check the total size of the data contained in a partitioned table `measurement`, one could use the following query:

```sql

SELECT pg_size_pretty(sum(pg_relation_size(relid))) AS total_size
  FROM pg_partition_tree('measurement');
```

  <a id="functions-admin-index"></a>

### Index Maintenance Functions


 [Index Maintenance Functions](#functions-admin-index-table) shows the functions available for index maintenance tasks. (Note that these maintenance tasks are normally done automatically by autovacuum; use of these functions is only required in special cases.) These functions cannot be executed during recovery. Use of these functions is restricted to superusers and the owner of the given index.
 <a id="functions-admin-index-table"></a>

**Table: Index Maintenance Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>brin_summarize_new_values</code> ( <code>index</code> <code>regclass</code> ) <code>integer</code></td>
<td>Scans the specified BRIN index to find page ranges in the base table that are not currently summarized by the index; for any such range it creates a new summary index tuple by scanning those table pages. Returns the number of new page range summaries that were inserted into the index.</td>
<td></td>
</tr>
<tr>
<td><code>brin_summarize_range</code> ( <code>index</code> <code>regclass</code>, <code>blockNumber</code> <code>bigint</code> ) <code>integer</code></td>
<td>Summarizes the page range covering the given block, if not already summarized. This is like <code>brin_summarize_new_values</code> except that it only processes the page range that covers the given table block number.</td>
<td></td>
</tr>
<tr>
<td><code>brin_desummarize_range</code> ( <code>index</code> <code>regclass</code>, <code>blockNumber</code> <code>bigint</code> ) <code>void</code></td>
<td>Removes the BRIN index tuple that summarizes the page range covering the given table block, if there is one.</td>
<td></td>
</tr>
<tr>
<td><code>gin_clean_pending_list</code> ( <code>index</code> <code>regclass</code> ) <code>bigint</code></td>
<td>Cleans up the “pending” list of the specified GIN index by moving entries in it, in bulk, to the main GIN data structure. Returns the number of pages removed from the pending list. If the argument is a GIN index built with the <code>fastupdate</code> option disabled, no cleanup happens and the result is zero, because the index doesn't have a pending list. See <a href="../../internals/built-in-index-access-methods/gin-indexes.md#gin-fast-update">GIN Fast Update Technique</a> and <a href="../../internals/built-in-index-access-methods/gin-indexes.md#gin-tips">GIN Tips and Tricks</a> for details about the pending list and <code>fastupdate</code> option.</td>
<td></td>
</tr>
</tbody>
</table>
  <a id="functions-admin-genfile"></a>

### Generic File Access Functions


 The functions shown in [Generic File Access Functions](#functions-admin-genfile-table) provide native access to files on the machine hosting the server. Only files within the database cluster directory and the `log_directory` can be accessed, unless the user is a superuser or is granted the role `pg_read_server_files`. Use a relative path for files in the cluster directory, and a path matching the `log_directory` configuration setting for log files.


 Note that granting users the EXECUTE privilege on `pg_read_file()`, or related functions, allows them the ability to read any file on the server that the database server process can read; these functions bypass all in-database privilege checks. This means that, for example, a user with such access is able to read the contents of the `pg_authid` table where authentication information is stored, as well as read any table data in the database. Therefore, granting access to these functions should be carefully considered.


 When granting privilege on these functions, note that the table entries showing optional parameters are mostly implemented as several physical functions with different parameter lists. Privilege must be granted separately on each such function, if it is to be used. psql's `\df` command can be useful to check what the actual function signatures are.


 Some of these functions take an optional `missing_ok` parameter, which specifies the behavior when the file or directory does not exist. If `true`, the function returns `NULL` or an empty result set, as appropriate. If `false`, an error is raised. (Failure conditions other than “file not found” are reported as errors in any case.) The default is `false`.
 <a id="functions-admin-genfile-table"></a>

**Table: Generic File Access Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_ls_dir</code> ( <code>dirname</code> <code>text</code> [, <code>missing_ok</code> <code>boolean</code>, <code>include_dot_dirs</code> <code>boolean</code> ] ) <code>setof text</code></td>
<td>Returns the names of all files (and directories and other special files) in the specified directory. The <code>include_dot_dirs</code> parameter indicates whether “.” and “..” are to be included in the result set; the default is to exclude them. Including them can be useful when <code>missing_ok</code> is <code>true</code>, to distinguish an empty directory from a non-existent directory.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_logdir</code> () <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the server's log directory. Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and roles with privileges of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_waldir</code> () <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the server's write-ahead log (WAL) directory. Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and roles with privileges of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_logicalmapdir</code> () <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the server's <code>pg_logical/mappings</code> directory. Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and members of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_logicalsnapdir</code> () <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the server's <code>pg_logical/snapshots</code> directory. Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and members of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_replslotdir</code> ( <code>slot_name</code> <code>text</code> ) <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the server's <code>pg_replslot/slot_name</code> directory, where <code>slot_name</code> is the name of the replication slot provided as input of the function. Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and members of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_archive_statusdir</code> () <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the server's WAL archive status directory (<code>pg_wal/archive_status</code>). Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and members of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_ls_tmpdir</code> ( [ <code>tablespace</code> <code>oid</code> ] ) <code>setof record</code> ( <code>name</code> <code>text</code>, <code>size</code> <code>bigint</code>, <code>modification</code> <code>timestamp with time zone</code> )</td>
<td>Returns the name, size, and last modification time (mtime) of each ordinary file in the temporary file directory for the specified <code>tablespace</code>. If <code>tablespace</code> is not provided, the <code>pg_default</code> tablespace is examined. Filenames beginning with a dot, directories, and other special files are excluded.</td>
<td>This function is restricted to superusers and members of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_read_file</code> ( <code>filename</code> <code>text</code> [, <code>offset</code> <code>bigint</code>, <code>length</code> <code>bigint</code> ] [, <code>missing_ok</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Returns all or part of a text file, starting at the given byte <code>offset</code>, returning at most <code>length</code> bytes (less if the end of file is reached first). If <code>offset</code> is negative, it is relative to the end of the file. If <code>offset</code> and <code>length</code> are omitted, the entire file is returned. The bytes read from the file are interpreted as a string in the database's encoding; an error is thrown if they are not valid in that encoding.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_read_binary_file</code> ( <code>filename</code> <code>text</code> [, <code>offset</code> <code>bigint</code>, <code>length</code> <code>bigint</code> ] [, <code>missing_ok</code> <code>boolean</code> ] ) <code>bytea</code></td>
<td>Returns all or part of a file. This function is identical to <code>pg_read_file</code> except that it can read arbitrary binary data, returning the result as <code>bytea</code> not <code>text</code>; accordingly, no encoding checks are performed.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.<br><p>In combination with the <code>convert_from</code> function, this function can be used to read a text file in a specified encoding and convert to the database's encoding:</p>
<pre><code class="language-sql">
SELECT convert_from(pg_read_binary_file('file_in_utf8.txt'), 'UTF8');</code></pre></td>
</tr>
<tr>
<td><code>pg_stat_file</code> ( <code>filename</code> <code>text</code> [, <code>missing_ok</code> <code>boolean</code> ] ) <code>record</code> ( <code>size</code> <code>bigint</code>, <code>access</code> <code>timestamp with time zone</code>, <code>modification</code> <code>timestamp with time zone</code>, <code>change</code> <code>timestamp with time zone</code>, <code>creation</code> <code>timestamp with time zone</code>, <code>isdir</code> <code>boolean</code> )</td>
<td>Returns a record containing the file's size, last access time stamp, last modification time stamp, last file status change time stamp (Unix platforms only), file creation time stamp (Windows only), and a flag indicating if it is a directory.</td>
<td>This function is restricted to superusers by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
</tbody>
</table>
  <a id="functions-advisory-locks"></a>

### Advisory Lock Functions


 The functions shown in [Advisory Lock Functions](#functions-advisory-locks-table) manage advisory locks. For details about proper use of these functions, see [Advisory Locks](../concurrency-control/explicit-locking.md#advisory-locks).


 All these functions are intended to be used to lock application-defined resources, which can be identified either by a single 64-bit key value or two 32-bit key values (note that these two key spaces do not overlap). If another session already holds a conflicting lock on the same resource identifier, the functions will either wait until the resource becomes available, or return a `false` result, as appropriate for the function. Locks can be either shared or exclusive: a shared lock does not conflict with other shared locks on the same resource, only with exclusive locks. Locks can be taken at session level (so that they are held until released or the session ends) or at transaction level (so that they are held until the current transaction ends; there is no provision for manual release). Multiple session-level lock requests stack, so that if the same resource identifier is locked three times there must then be three unlock requests to release the resource in advance of session end.
 <a id="functions-advisory-locks-table"></a>

**Table: Advisory Lock Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_advisory_lock</code> ( <code>key</code> <code>bigint</code> ) <code>void</code></td>
<td><code>pg_advisory_lock</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>void</code></td>
<td>Obtains an exclusive session-level advisory lock, waiting if necessary.</td>
</tr>
<tr>
<td><code>pg_advisory_lock_shared</code> ( <code>key</code> <code>bigint</code> ) <code>void</code></td>
<td><code>pg_advisory_lock_shared</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>void</code></td>
<td>Obtains a shared session-level advisory lock, waiting if necessary.</td>
</tr>
<tr>
<td><code>pg_advisory_unlock</code> ( <code>key</code> <code>bigint</code> ) <code>boolean</code></td>
<td><code>pg_advisory_unlock</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>boolean</code></td>
<td>Releases a previously-acquired exclusive session-level advisory lock. Returns <code>true</code> if the lock is successfully released. If the lock was not held, <code>false</code> is returned, and in addition, an SQL warning will be reported by the server.</td>
</tr>
<tr>
<td><code>pg_advisory_unlock_all</code> () <code>void</code></td>
<td>Releases all session-level advisory locks held by the current session. (This function is implicitly invoked at session end, even if the client disconnects ungracefully.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_advisory_unlock_shared</code> ( <code>key</code> <code>bigint</code> ) <code>boolean</code></td>
<td><code>pg_advisory_unlock_shared</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>boolean</code></td>
<td>Releases a previously-acquired shared session-level advisory lock. Returns <code>true</code> if the lock is successfully released. If the lock was not held, <code>false</code> is returned, and in addition, an SQL warning will be reported by the server.</td>
</tr>
<tr>
<td><code>pg_advisory_xact_lock</code> ( <code>key</code> <code>bigint</code> ) <code>void</code></td>
<td><code>pg_advisory_xact_lock</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>void</code></td>
<td>Obtains an exclusive transaction-level advisory lock, waiting if necessary.</td>
</tr>
<tr>
<td><code>pg_advisory_xact_lock_shared</code> ( <code>key</code> <code>bigint</code> ) <code>void</code></td>
<td><code>pg_advisory_xact_lock_shared</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>void</code></td>
<td>Obtains a shared transaction-level advisory lock, waiting if necessary.</td>
</tr>
<tr>
<td><code>pg_try_advisory_lock</code> ( <code>key</code> <code>bigint</code> ) <code>boolean</code></td>
<td><code>pg_try_advisory_lock</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>boolean</code></td>
<td>Obtains an exclusive session-level advisory lock if available. This will either obtain the lock immediately and return <code>true</code>, or return <code>false</code> without waiting if the lock cannot be acquired immediately.</td>
</tr>
<tr>
<td><code>pg_try_advisory_lock_shared</code> ( <code>key</code> <code>bigint</code> ) <code>boolean</code></td>
<td><code>pg_try_advisory_lock_shared</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>boolean</code></td>
<td>Obtains a shared session-level advisory lock if available. This will either obtain the lock immediately and return <code>true</code>, or return <code>false</code> without waiting if the lock cannot be acquired immediately.</td>
</tr>
<tr>
<td><code>pg_try_advisory_xact_lock</code> ( <code>key</code> <code>bigint</code> ) <code>boolean</code></td>
<td><code>pg_try_advisory_xact_lock</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>boolean</code></td>
<td>Obtains an exclusive transaction-level advisory lock if available. This will either obtain the lock immediately and return <code>true</code>, or return <code>false</code> without waiting if the lock cannot be acquired immediately.</td>
</tr>
<tr>
<td><code>pg_try_advisory_xact_lock_shared</code> ( <code>key</code> <code>bigint</code> ) <code>boolean</code></td>
<td><code>pg_try_advisory_xact_lock_shared</code> ( <code>key1</code> <code>integer</code>, <code>key2</code> <code>integer</code> ) <code>boolean</code></td>
<td>Obtains a shared transaction-level advisory lock if available. This will either obtain the lock immediately and return <code>true</code>, or return <code>false</code> without waiting if the lock cannot be acquired immediately.</td>
</tr>
</tbody>
</table>
