# CHECKPOINT { #sql-checkpoint }

force a write-ahead log checkpoint

## Synopsis


```

CHECKPOINT [ ( option [, ...] ) ]

where OPTION can be one of:

    FLUSH_UNLOGGED [ BOOLEAN ]
    MODE { FAST | SPREAD }
```


## Description


 A checkpoint is a point in the write-ahead log sequence at which all data files have been updated to reflect the information in the log. All data files will be flushed to disk. Refer to [WAL Configuration](../../server-administration/reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration) for more details about what happens during a checkpoint.


 By default, the `CHECKPOINT` command forces a fast checkpoint when the command is issued, without waiting for a regular checkpoint scheduled by the system (controlled by the settings in [Checkpoints](../../server-administration/server-configuration/write-ahead-log.md#runtime-config-wal-checkpoints)). To request the checkpoint be spread over a longer interval, set the `MODE` option to `SPREAD`. `CHECKPOINT` is not intended for use during normal operation.


 The server may consolidate concurrently requested checkpoints. Such consolidated requests will contain a combined set of options. For example, if one session requests a fast checkpoint and another requests a spread checkpoint, the server may combine those requests and perform one fast checkpoint.


 If executed during recovery, the `CHECKPOINT` command will force a restartpoint (see [WAL Configuration](../../server-administration/reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration)) rather than writing a new checkpoint.


 Only superusers or users with the privileges of the [pg_checkpoint](../../server-administration/database-roles/predefined-roles.md#predefined-role-pg-checkpoint) role can call `CHECKPOINT`.


## Parameters


`FLUSH_UNLOGGED`
:   Normally, `CHECKPOINT` does not flush dirty buffers of unlogged relations. This option, which is disabled by default, enables flushing unlogged relations to disk.

`MODE`
:   When set to `FAST`, which is the default, the requested checkpoint will be completed as fast as possible, which may result in a significantly higher rate of I/O during the checkpoint.


     `MODE` can also be set to `SPREAD` to request the checkpoint be spread over a longer interval (controlled via the settings in [Checkpoints](../../server-administration/server-configuration/write-ahead-log.md#runtime-config-wal-checkpoints)), like a regular checkpoint scheduled by the system. This can reduce the rate of I/O during the checkpoint.

*boolean*
:   Specifies whether the selected option should be turned on or off. You can write `TRUE`, `ON`, or `1` to enable the option, and `FALSE`, `OFF`, or `0` to disable it. The *boolean* value can also be omitted, in which case `TRUE` is assumed.


## Compatibility


 The `CHECKPOINT` command is a PostgreSQL language extension.
