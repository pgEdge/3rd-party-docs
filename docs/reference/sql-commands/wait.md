<a id="sql-wait"></a>

# WAIT

wait for WAL to reach a target LSN

## Synopsis


```

WAIT FOR LSN 'LSN'
    [ WITH ( OPTION [, ...] ) ]

where OPTION can be:

    MODE 'MODE'
    TIMEOUT 'TIMEOUT'
    NO_THROW

and MODE can be:

    standby_replay | standby_write | standby_flush | primary_flush
```


## Description


 Waits until the specified `lsn` is reached according to the specified `mode`, which determines whether to wait for WAL to be written, flushed, or replayed. If no `timeout` is specified or it is set to zero, this command waits indefinitely for the `lsn`.


 On timeout, an error is emitted unless `NO_THROW` is specified in the WITH clause. For standby modes (`standby_replay`, `standby_write`, `standby_flush`), an error is also emitted if the server is promoted before the `lsn` is reached. If `NO_THROW` is specified, the command reports these outcomes as a status string instead of raising an error.


 The possible return values are `success`, `timeout`, and `not in recovery`.


## Parameters


*lsn*
:   Specifies the target LSN to wait for.

<code>WITH ( </code><em>option</em><code> [, ...] )</code>
:   This clause specifies optional parameters for the wait operation. The following parameters are supported:

    `MODE` '*mode*'
    :   Specifies the type of LSN processing to wait for. If not specified, the default is `standby_replay`. The valid modes are:


        -  `standby_replay`: Wait for the LSN to be replayed (applied to the database) on a standby server. After successful completion, `pg_last_wal_replay_lsn()` will return a value greater than or equal to the target LSN. This mode can only be used during recovery.
        -  `standby_write`: Wait for the WAL containing the LSN to be written to disk on a standby server, but not yet necessarily flushed. This is faster than `standby_flush` but provides weaker durability guarantees since the data may still be in operating system buffers. This is satisfied by WAL already present on the standby from a base backup, archive restore, or prior streaming, as well as WAL newly received from the primary. This mode can only be used during recovery.
        -  `standby_flush`: Wait for the WAL containing the LSN to be flushed to disk on a standby server. This provides a durability guarantee without waiting for the WAL to be applied. This is satisfied by WAL already present on the standby from a base backup, archive restore, or prior streaming, as well as WAL newly received from the primary. This mode can only be used during recovery.
        -  `primary_flush`: Wait for the WAL containing the LSN to be flushed to disk on a primary server. After successful completion, `pg_current_wal_flush_lsn()` will return a value greater than or equal to the target LSN. This mode can only be used on a primary server (not during recovery).

    `TIMEOUT` '*timeout*'
    :   When specified and `timeout` is greater than zero, the command waits until `lsn` is reached or the specified `timeout` has elapsed. A value of zero (the default) means the command waits indefinitely.


         The `timeout` is an amount of time in milliseconds. It may also be specified as a string containing the numerical value followed by a time unit (see [Parameter Names and Values](../../server-administration/server-configuration/setting-parameters.md#config-setting-names-values)). The maximum value is `2147483647 ms`.


         Fractional values are rounded to the nearest millisecond. Note that a `timeout` of half a millisecond or less therefore rounds down to zero, which means waiting indefinitely.

    `NO_THROW`
    :   Specify to not throw an error in the case of timeout or running on the primary. In this case the result status can be obtained from the return value.


         Use this option when `timeout` or `not in recovery` is an expected result that the application intends to handle, for example by retrying the wait, reporting replication delay, or choosing another server for a subsequent operation. The command then returns the result as a status, which the application must check before assuming that the target LSN was reached. Omit the option when the application must not proceed unless the target LSN is reached, so that an unsuccessful wait stops normal execution with an error.


         Returning a status also leaves an explicit transaction usable; without this option, the corresponding error requires rolling back the transaction, or rolling back to a savepoint, before further commands can be issued.


         This option changes only how `timeout` and `not in recovery` are reported. Other errors are still raised. That covers invalid input, such as a malformed LSN or an unrecognized option value, and every condition that is checked before the wait begins, such as requesting `primary_flush` during recovery or holding a lock while waiting for a standby LSN. The option also does not limit the duration of the wait; specify `TIMEOUT` for that purpose.


## Outputs


`success`
:   This return value denotes that we have successfully reached the target `lsn`.

`timeout`
:   This return value denotes that the timeout happened before reaching the target `lsn`.

`not in recovery`
:   This return value denotes that the database server is not in a recovery state. This might mean either the database server was not in recovery at the moment of receiving the command (i.e., executed on a primary), or it was promoted before reaching the target `lsn`. In the promotion case, this status indicates a timeline change occurred, and the application should re-evaluate whether the target LSN is still relevant.


## Notes


 `WAIT` must be executed as a top-level command. It cannot be executed from a function, procedure, or `DO` block. It also cannot be executed while the current transaction holds a snapshot. `WAIT` itself acquires none, so it can run before the first snapshot-taking statement of a `REPEATABLE READ` or `SERIALIZABLE` transaction, but not after it, and not while a cursor or an exported snapshot holds one at any isolation level. A snapshot held here could delay replay, which `standby_replay` waits for and which other standby modes can end up waiting for too. That is also why `WAIT` is a command rather than a function or a procedure, which execute with one held.


 While recovery is in progress, a wait in `standby_replay` (the default), `standby_write`, or `standby_flush` mode is rejected when the session already holds a lock and the target `lsn` has not been reached yet. Such a lock can make the startup process wait for this session, either directly or through another session, while this session waits for the startup process to advance recovery. That cycle involves no lock wait on this side, so deadlock detection does not see it and nothing breaks it. A wait whose target has already been reached returns immediately and is therefore always allowed.


 Issue `WAIT FOR` outside a transaction block, or as the first statement of one, before running anything that takes locks. That is also the natural order for the read-your-writes pattern shown in the examples below: wait for the target `lsn` first, then run the queries that have to see it. Note that a lock taken by an earlier statement is still held at `READ COMMITTED`, even though its snapshot is gone, so a wait placed after such a statement is rejected even when the isolation level permits it.


 The restriction covers `standby_write` and `standby_flush` as well, even though streaming replication can advance those positions without the startup process. Both positions are at least the replay position, so without an active walreceiver replay can be their only source of progress. If a held lock blocks replay, the session waits for replay to advance while replay waits for the session to release the lock. Under streaming replication the positions advance independently only while WAL keeps arriving. If reception stops before the target is reached, a blocked startup process cannot restart the walreceiver. It also cannot replay newer checkpoint records needed to advance restartpoints and recycle WAL, so `pg_wal` can fill up and reception can stop before the target is reached. The restriction therefore also applies when streaming is active at the start of the wait.


 `WAIT` waits until the specified `lsn` is reached according to the specified `mode`. The `standby_replay` mode waits for the LSN to be replayed (applied to the database), which is useful to achieve read-your-writes consistency while using an async replica for reads and the primary for writes, provided that the target LSN is at or after the end of the relevant write transaction's `COMMIT` record on the primary. The `standby_flush` mode waits for the WAL to be flushed to durable storage on the replica, or to have already been replayed from WAL present on the standby. The `standby_write` mode waits for the WAL to be written to the operating system, or to have already been replayed, which is faster than flush for newly received WAL but provides weaker durability guarantees. The `primary_flush` mode waits for WAL to be flushed on a primary server. In all cases, the LSN of the last modification should be stored on the client application side or the connection pooler side.


 The standby modes (`standby_replay`, `standby_write`, `standby_flush`) can only be used during recovery, and `primary_flush` can only be used on a primary server. Using the wrong mode for the current server state will result in an error. If a standby is promoted while waiting with a standby mode, the command will return `not in recovery` (or throw an error if `NO_THROW` is not specified). Promotion creates a new timeline, and the LSN being waited for may refer to WAL from the old timeline.


 `WAIT` compares only the numeric LSN; it has no notion of which timeline a WAL record belongs to. This matters when a standby continues recovery across an upstream timeline switch — for example, a cascading standby whose upstream gets promoted. In that case `WAIT` will return `success` as soon as the position used by the selected wait mode reaches or passes the numeric LSN, regardless of which timeline that LSN belongs to. Applications that need to confirm the target refers to the expected timeline must validate the timeline themselves.


 On a standby server, `WAIT` sessions may be interrupted by recovery conflicts. Some recovery conflicts are unavoidable: for example, replaying a tablespace drop resolves conflicts by terminating all backends, regardless of what they are doing. Applications using `WAIT` on a standby should be prepared to handle such interruptions, for example by retrying the command or falling back to an alternative mechanism.


## Examples


 You can use the `WAIT` command to wait for the `pg_lsn` value. For example, an application could update the `movie` table and get an lsn that is at or after the end of the relevant write transaction's `COMMIT` record. In the default autocommit mode shown here, the `UPDATE` commits before the subsequent `SELECT`. This example uses `pg_current_wal_insert_lsn` on primary server to get the lsn given that `synchronous_commit` could be set to `off`.

```

postgres=# UPDATE movie SET genre = 'Dramatic' WHERE genre = 'Drama';
UPDATE 100
postgres=# SELECT pg_current_wal_insert_lsn();
 pg_current_wal_insert_lsn
---------------------------
 0/0306EE20
(1 row)
```
 Then an application could run `WAIT` with the `lsn` obtained from the primary after the commit. After that, the changes made on the primary should be guaranteed to be visible on the replica.

```

postgres=# WAIT FOR LSN '0/0306EE20';
 status
---------
 success
(1 row)
postgres=# SELECT * FROM movie WHERE genre = 'Drama';
 genre
-------
(0 rows)
```


 Wait for flush (data durable on replica):

```

postgres=# WAIT FOR LSN '0/0306EE20' WITH (MODE 'standby_flush');
 status
---------
 success
(1 row)
```


 Wait for write with timeout:

```

postgres=# WAIT FOR LSN '0/0306EE20' WITH (MODE 'standby_write', TIMEOUT '100ms', NO_THROW);
 status
---------
 success
(1 row)
```


 Wait for flush on primary:

```

postgres=# WAIT FOR LSN '0/0306EE20' WITH (MODE 'primary_flush');
 status
---------
 success
(1 row)
```


 If the target LSN is not reached before the timeout, an error is thrown:

```

postgres=# WAIT FOR LSN '0/0306EE20' WITH (TIMEOUT '0.1s');
ERROR:  timed out while waiting for target LSN 0/0306EE20 to be replayed; current standby_replay LSN 0/0306EA60
```


 The same example uses `WAIT` with the `NO_THROW` option:

```

postgres=# WAIT FOR LSN '0/0306EE20' WITH (TIMEOUT '100ms', NO_THROW);
 status
---------
 timeout
(1 row)
```
