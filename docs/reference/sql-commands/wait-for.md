<a id="sql-wait-for"></a>

# WAIT FOR

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


 On timeout, an error is emitted unless `NO_THROW` is specified in the WITH clause. For standby modes (`standby_replay`, `standby_write`, `standby_flush`), an error is also emitted if the server is promoted before the `lsn` is reached. If `NO_THROW` is specified, the command returns a status string instead of throwing errors.


 The possible return values are `success`, `timeout`, and `not in recovery`.


## Parameters


*lsn*
:   Specifies the target LSN to wait for.

<code>WITH ( </code><em>option</em><code> [, ...] )</code>
:   This clause specifies optional parameters for the wait operation. The following parameters are supported:

    `MODE` '*mode*'
    :   Specifies the type of LSN processing to wait for. If not specified, the default is `standby_replay`. The valid modes are:


        -  `standby_replay`: Wait for the LSN to be replayed (applied to the database) on a standby server. After successful completion, `pg_last_wal_replay_lsn()` will return a value greater than or equal to the target LSN. This mode can only be used during recovery.
        -  `standby_write`: Wait for the WAL containing the LSN to be received from the primary and written to disk on a standby server, but not yet flushed. This is faster than `standby_flush` but provides weaker durability guarantees since the data may still be in operating system buffers. After successful completion, the `written_lsn` column in [`pg_stat_wal_receiver`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-wal-receiver-view) will show a value greater than or equal to the target LSN. This mode can only be used during recovery.
        -  `standby_flush`: Wait for the WAL containing the LSN to be received from the primary and flushed to disk on a standby server. This provides a durability guarantee without waiting for the WAL to be applied. After successful completion, `pg_last_wal_receive_lsn()` will return a value greater than or equal to the target LSN. This value is also available as the `flushed_lsn` column in [`pg_stat_wal_receiver`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-wal-receiver-view). This mode can only be used during recovery.
        -  `primary_flush`: Wait for the WAL containing the LSN to be flushed to disk on a primary server. After successful completion, `pg_current_wal_flush_lsn()` will return a value greater than or equal to the target LSN. This mode can only be used on a primary server (not during recovery).

    `TIMEOUT` '*timeout*'
    :   When specified and `timeout` is greater than zero, the command waits until `lsn` is reached or the specified `timeout` has elapsed.


         The `timeout` might be given as integer number of milliseconds. Also it might be given as string literal with integer number of milliseconds or a number with unit (see [Parameter Names and Values](../../server-administration/server-configuration/setting-parameters.md#config-setting-names-values)).

    `NO_THROW`
    :   Specify to not throw an error in the case of timeout or running on the primary. In this case the result status can be get from the return value.


## Outputs


`success`
:   This return value denotes that we have successfully reached the target `lsn`.

`timeout`
:   This return value denotes that the timeout happened before reaching the target `lsn`.

`not in recovery`
:   This return value denotes that the database server is not in a recovery state. This might mean either the database server was not in recovery at the moment of receiving the command (i.e., executed on a primary), or it was promoted before reaching the target `lsn`. In the promotion case, this status indicates a timeline change occurred, and the application should re-evaluate whether the target LSN is still relevant.


## Notes


 `WAIT FOR` must be executed as a top-level command. It cannot be executed from a function, procedure, or `DO` block. It also requires that no active or registered snapshot be held, and therefore cannot be used in contexts where such a snapshot must remain active, including transactions running at isolation levels higher than `READ COMMITTED`.


 `WAIT FOR` waits until the specified `lsn` is reached according to the specified `mode`. The `standby_replay` mode waits for the LSN to be replayed (applied to the database), which is useful to achieve read-your-writes consistency while using an async replica for reads and the primary for writes. The `standby_flush` mode waits for the WAL to be flushed to durable storage on the replica, providing a durability guarantee without waiting for replay. The `standby_write` mode waits for the WAL to be written to the operating system, which is faster than flush but provides weaker durability guarantees. The `primary_flush` mode waits for WAL to be flushed on a primary server. In all cases, the LSN of the last modification should be stored on the client application side or the connection pooler side.


 The standby modes (`standby_replay`, `standby_write`, `standby_flush`) can only be used during recovery, and `primary_flush` can only be used on a primary server. Using the wrong mode for the current server state will result in an error. If a standby is promoted while waiting with a standby mode, the command will return `not in recovery` (or throw an error if `NO_THROW` is not specified). Promotion creates a new timeline, and the LSN being waited for may refer to WAL from the old timeline.


 On a standby server, `WAIT FOR` sessions may be interrupted by recovery conflicts. Some recovery conflicts are unavoidable: for example, replaying a tablespace drop resolves conflicts by terminating all backends, regardless of what they are doing. Applications using `WAIT FOR` on a standby should be prepared to handle such interruptions, for example by retrying the command or falling back to an alternative mechanism.


## Examples


 You can use `WAIT FOR` command to wait for the `pg_lsn` value. For example, an application could update the `movie` table and get the lsn after changes just made. This example uses `pg_current_wal_insert_lsn` on primary server to get the lsn given that `synchronous_commit` could be set to `off`.

```

postgres=# UPDATE movie SET genre = 'Dramatic' WHERE genre = 'Drama';
UPDATE 100
postgres=# SELECT pg_current_wal_insert_lsn();
 pg_current_wal_insert_lsn
---------------------------
 0/306EE20
(1 row)
```
 Then an application could run `WAIT FOR` with the `lsn` obtained from primary. After that the changes made on primary should be guaranteed to be visible on replica.

```

postgres=# WAIT FOR LSN '0/306EE20';
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

postgres=# WAIT FOR LSN '0/306EE20' WITH (MODE 'standby_flush');
 status
---------
 success
(1 row)
```


 Wait for write with timeout:

```

postgres=# WAIT FOR LSN '0/306EE20' WITH (MODE 'standby_write', TIMEOUT '100ms', NO_THROW);
 status
---------
 success
(1 row)
```


 Wait for flush on primary:

```

postgres=# WAIT FOR LSN '0/306EE20' WITH (MODE 'primary_flush');
 status
---------
 success
(1 row)
```


 If the target LSN is not reached before the timeout, an error is thrown:

```

postgres=# WAIT FOR LSN '0/306EE20' WITH (TIMEOUT '0.1s');
ERROR:  timed out while waiting for target LSN 0/306EE20 to be replayed; current replay LSN 0/306EA60
```


 The same example uses `WAIT FOR` with `NO_THROW` option:

```

postgres=# WAIT FOR LSN '0/306EE20' WITH (TIMEOUT '100ms', NO_THROW);
 status
---------
 timeout
(1 row)
```
