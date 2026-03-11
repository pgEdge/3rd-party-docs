<a id="recovery-config"></a>

## `recovery.conf` file merged into `postgresql.conf`


 PostgreSQL 11 and below used a configuration file named `recovery.conf`  to manage replicas and standbys. Support for this file was removed in PostgreSQL 12. See [the release notes for PostgreSQL 12](../release-notes/prior-releases.md#release-prior) for details on this change.


 On PostgreSQL 12 and above, [archive recovery, streaming replication, and PITR](../../server-administration/backup-and-restore/continuous-archiving-and-point-in-time-recovery-pitr.md#continuous-archiving) are configured using [normal server configuration parameters](../../server-administration/server-configuration/replication.md#runtime-config-replication-standby). These are set in `postgresql.conf` or via [ALTER SYSTEM](../../reference/sql-commands/alter-system.md#sql-altersystem) like any other parameter.


 The server will not start if a `recovery.conf` exists.


 PostgreSQL 15 and below had a setting `promote_trigger_file`, or `trigger_file` before 12. Use `pg_ctl promote` or call `pg_promote()` to promote a standby instead.


 The `standby_mode`  setting has been removed. A `standby.signal` file in the data directory is used instead. See [Standby Server Operation](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#standby-server-operation) for details.
