## Release 17 { #release-17 }


**Release date:.**


2024-09-26


### Overview { #release-17-highlights }


 PostgreSQL 17 contains many new features and enhancements, including:


-  New memory management system for `VACUUM`, which reduces memory consumption and can improve overall vacuuming performance.
-  New SQL/JSON capabilities, including constructors, identity functions, and the [`JSON_TABLE()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-table) function, which converts JSON data into a table representation.
-  Various query performance improvements, including for sequential reads using streaming I/O, write throughput under high concurrency, and searches over multiple values in a [btree](../../internals/built-in-index-access-methods/b-tree-indexes.md#btree) index.
-  Logical replication enhancements, including:

-  Failover control
-  [pg_createsubscriber](../../reference/postgresql-server-applications/pg_createsubscriber.md#app-pgcreatesubscriber), a utility that creates logical replicas from physical standbys
-  [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) now preserves logical replication slots on publishers and full subscription state on subscribers. This will allow upgrades to future major versions to continue logical replication without requiring copy to resynchronize.

-  New client-side connection option, [`sslnegotiation=direct`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-sslnegotiation), that performs a direct TLS handshake to avoid a round-trip negotiation.
-  [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) now supports incremental backup.
-  [`COPY`](../../reference/sql-commands/copy.md#sql-copy) adds a new option, `ON_ERROR ignore`, that allows a copy operation to continue in the event of an error.


 The above items and other new features of PostgreSQL 17 are explained in more detail in the sections below.


### Migration to Version 17 { #release-17-migration }


 A dump/restore using [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall) or use of [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) or logical replication is required for those wishing to migrate data from any previous release. See [Upgrading a PostgreSQL Cluster](../../server-administration/server-setup-and-operation/upgrading-a-postgresql-cluster.md#upgrading) for general information on migrating to new major releases.


 Version 17 contains a number of changes that may affect compatibility with previous releases. Observe the following incompatibilities:


-  Change functions to use a safe [search_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-search-path) during maintenance operations (Jeff Davis) [&sect;](https://postgr.es/c/2af07e2f7) [&sect;](https://postgr.es/c/b4da732fd64)

   This prevents maintenance operations (`ANALYZE`, `CLUSTER`, `CREATE INDEX`, `CREATE MATERIALIZED VIEW`, `REFRESH MATERIALIZED VIEW`, `REINDEX`, or `VACUUM`) from performing unsafe access. Functions used by expression indexes and materialized views that need to reference non-default schemas must specify a search path during function creation.
-  Restrict `ago` to only appear at the end in `interval` values (Joseph Koshakow) [&sect;](https://postgr.es/c/165d581f1) [&sect;](https://postgr.es/c/617f9b7d4)

   Also, prevent empty interval units from appearing multiple times.
-  Remove server variable old_snapshot_threshold (Thomas Munro) [&sect;](https://postgr.es/c/f691f5b80)

   This variable allowed vacuum to remove rows that potentially could be still visible to running transactions, causing "snapshot too old" errors later if accessed. This feature might be re-added to PostgreSQL later if an improved implementation is found.
-  Change [`SET SESSION AUTHORIZATION`](../../reference/sql-commands/set-session-authorization.md#sql-set-session-authorization) handling of the initial session user's superuser status (Joseph Koshakow) [&sect;](https://postgr.es/c/a0363ab7a)

   The new behavior is based on the session user's superuser status at the time the `SET SESSION AUTHORIZATION` command is issued, rather than their superuser status at connection time.
-  Remove feature which simulated per-database users (Nathan Bossart) [&sect;](https://postgr.es/c/884eee5bf)

   The feature, `db_user_namespace`, was rarely used.
-  Remove adminpack contrib extension (Daniel Gustafsson) [&sect;](https://postgr.es/c/cc09e6549)

   This was used by now end-of-life pgAdmin III.
-  Remove [wal_sync_method](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-sync-method) value `fsync_writethrough` on `Windows` (Thomas Munro) [&sect;](https://postgr.es/c/d0c28601e)

   This value was the same as `fsync` on `Windows`.
-  Change file boundary handling of two WAL file name functions (Kyotaro Horiguchi, Andres Freund, Bruce Momjian) [&sect;](https://postgr.es/c/344afc776)

   The functions [`pg_walfile_name()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-backup-table) and `pg_walfile_name_offset()` used to report the previous LSN segment number when the LSN was on a file segment boundary; it now returns the current LSN segment.
-  Remove server variable `trace_recovery_messages` since it is no longer needed (Bharath Rupireddy) [&sect;](https://postgr.es/c/c7a3e6b46)
-  Remove [information schema](../../client-interfaces/the-information-schema/index.md#information-schema) column `element_types`.`domain_default` (Peter Eisentraut) [&sect;](https://postgr.es/c/78806a950)
-  Change [pgrowlocks](../additional-supplied-modules-and-extensions/pgrowlocks-show-a-tables-row-locking-information.md#pgrowlocks) lock mode output labels (Bruce Momjian) [&sect;](https://postgr.es/c/15d5d7405)
-  Remove `buffers_backend` and `buffers_backend_fsync` from [`pg_stat_bgwriter`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-bgwriter-view) (Bharath Rupireddy) [&sect;](https://postgr.es/c/74604a37f)

   These fields are considered redundant to similar columns in [`pg_stat_io`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-io-view).
-  Rename I/O block read/write timing statistics columns of [pg_stat_statements](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements) (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/13d00729d)

   This renames `blk_read_time` to `shared_blk_read_time`, and `blk_write_time` to `shared_blk_write_time`.
-  Change [`pg_attribute`.`attstattarget`](../../internals/system-catalogs/pg_attribute.md#catalog-pg-attribute) and `pg_statistic_ext`.`stxstattarget` to represent the default statistics target as `NULL` (Peter Eisentraut) [&sect;](https://postgr.es/c/4f622503d) [&sect;](https://postgr.es/c/012460ee9)
-  Rename [`pg_collation`.`colliculocale`](../../internals/system-catalogs/pg_collation.md#catalog-pg-collation) to `colllocale` and [`pg_database`.`daticulocale`](../../internals/system-catalogs/pg_database.md#catalog-pg-database) to `datlocale` (Jeff Davis) [&sect;](https://postgr.es/c/f696c0cd5)
-  Rename [`pg_stat_progress_vacuum`](../../server-administration/monitoring-database-activity/progress-reporting.md#vacuum-progress-reporting) column `max_dead_tuples` to `max_dead_tuple_bytes`, rename `num_dead_tuples` to `num_dead_item_ids`, and add `dead_tuple_bytes` (Masahiko Sawada) [&sect;](https://postgr.es/c/667e65aac) [&sect;](https://postgr.es/c/f1affb670)
-  Rename SLRU columns in system view [`pg_stat_slru`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-slru-view) (Alvaro Herrera) [&sect;](https://postgr.es/c/bcdfa5f2e)

   The column names accepted by [`pg_stat_reset_slru()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-funcs-table) are also changed.


### Changes { #release-17-changes }


 Below you will find a detailed account of the changes between PostgreSQL 17 and the previous major release.


#### Server { #release-17-server }


##### Optimizer { #release-17-optimizer }


-  Allow the optimizer to improve CTE plans by considering the statistics and sort order of columns referenced in earlier row output clauses (Jian Guo, Richard Guo, Tom Lane) [&sect;](https://postgr.es/c/f7816aec2) [&sect;](https://postgr.es/c/a65724dfa)
-  Improve optimization of `IS NOT NULL` and `IS NULL` query restrictions (David Rowley, Richard Guo, Andy Fan) [&sect;](https://postgr.es/c/b262ad440) [&sect;](https://postgr.es/c/3af704098)

   Remove `IS NOT NULL` restrictions from queries on `NOT NULL` columns and eliminate scans on `NOT NULL` columns if `IS NULL` is specified.
-  Allow partition pruning on boolean columns on `IS [NOT] UNKNOWN` conditionals (David Rowley) [&sect;](https://postgr.es/c/07c36c133)
-  Improve optimization of range values when using containment operators <@ and @> (Kim Johan Andersson, Jian He) [&sect;](https://postgr.es/c/075df6b20)
-  Allow correlated `IN` subqueries to be transformed into joins (Andy Fan, Tom Lane) [&sect;](https://postgr.es/c/9f1337639)
-  Improve optimization of the `LIMIT` clause on partitioned tables, inheritance parents, and `UNION ALL` queries (Andy Fan, David Rowley) [&sect;](https://postgr.es/c/a8a968a82)
-  Allow query nodes to be run in parallel in more cases (Tom Lane) [&sect;](https://postgr.es/c/e08d74ca1)
-  Allow `GROUP BY` columns to be internally ordered to match `ORDER BY` (Andrei Lepikhov, Teodor Sigaev) [&sect;](https://postgr.es/c/0452b461b)

   This can be disabled using server variable [enable_group_by_reordering](../../server-administration/server-configuration/query-planning.md#guc-enable-groupby-reordering).
-  Allow `UNION` (without `ALL`) to use MergeAppend (David Rowley) [&sect;](https://postgr.es/c/66c0185a3)
-  Fix MergeAppend plans to more accurately compute the number of rows that need to be sorted (Alexander Kuzmenkov) [&sect;](https://postgr.es/c/9d1a5354f)
-  Allow [GiST](../../internals/built-in-index-access-methods/gist-indexes.md#gist) and [SP-GiST](../../internals/built-in-index-access-methods/sp-gist-indexes.md#spgist) indexes to be part of incremental sorts (Miroslav Bendik) [&sect;](https://postgr.es/c/625d5b3ca)

   This is particularly useful for `ORDER BY` clauses where the first column has a GiST and SP-GiST index, and other columns do not.
-  Add columns to [`pg_stats`](../../internals/system-views/pg_stats.md#view-pg-stats) to report range-type histogram information (Egor Rogov, Soumyadeep Chakraborty) [&sect;](https://postgr.es/c/bc3c8db8a)


##### Indexes { #release-17-indexes }


-  Allow [btree](../../internals/built-in-index-access-methods/b-tree-indexes.md#btree) indexes to more efficiently find a set of values, such as those supplied by `IN` clauses using constants (Peter Geoghegan, Matthias van de Meent) [&sect;](https://postgr.es/c/5bf748b86)
-  Allow [BRIN](../../internals/built-in-index-access-methods/brin-indexes.md#brin) indexes to be created using parallel workers (Tomas Vondra, Matthias van de Meent) [&sect;](https://postgr.es/c/b43757171)


##### General Performance { #release-17-performance }


-  Allow vacuum to more efficiently remove and freeze tuples (Melanie Plageman, Heikki Linnakangas) [&sect;](https://postgr.es/c/6dbb49026)

   WAL traffic caused by vacuum is also more compact.
-  Allow vacuum to more efficiently store tuple references (Masahiko Sawada, John Naylor) [&sect;](https://postgr.es/c/ee1b30f12) [&sect;](https://postgr.es/c/30e144287) [&sect;](https://postgr.es/c/667e65aac) [&sect;](https://postgr.es/c/6dbb49026)

   Additionally, vacuum is no longer silently limited to one gigabyte of memory when [maintenance_work_mem](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-work-mem) or [autovacuum_work_mem](../../server-administration/server-configuration/resource-consumption.md#guc-autovacuum-work-mem) are higher.
-  Optimize vacuuming of relations with no indexes (Melanie Plageman) [&sect;](https://postgr.es/c/c120550ed)
-  Increase default [vacuum_buffer_usage_limit](../../server-administration/server-configuration/resource-consumption.md#guc-vacuum-buffer-usage-limit) to 2MB (Thomas Munro) [&sect;](https://postgr.es/c/98f320eb2)
-  Improve performance when checking roles with many memberships (Nathan Bossart) [&sect;](https://postgr.es/c/d365ae705)
-  Improve performance of heavily-contended WAL writes (Bharath Rupireddy) [&sect;](https://postgr.es/c/71e4cc6b8)
-  Improve performance when transferring large blocks of data to a client (Melih Mutlu) [&sect;](https://postgr.es/c/c4ab7da60)
-  Allow the grouping of file system reads with the new system variable [io_combine_limit](../../server-administration/server-configuration/resource-consumption.md#guc-io-combine-limit) (Thomas Munro, Andres Freund, Melanie Plageman, Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/210622c60) [&sect;](https://postgr.es/c/b7b0f3f27) [&sect;](https://postgr.es/c/041b96802)


##### Monitoring { #release-17-monitoring }


-  Create system view [`pg_stat_checkpointer`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-checkpointer-view) (Bharath Rupireddy, Anton A. Melnikov, Alexander Korotkov) [&sect;](https://postgr.es/c/96f052613) [&sect;](https://postgr.es/c/12915a58e) [&sect;](https://postgr.es/c/e820db5b5)

   Relevant columns have been removed from [`pg_stat_bgwriter`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-bgwriter-view) and added to this new system view.
-  Improve control over resetting statistics (Atsushi Torikoshi, Bharath Rupireddy) [&sect;](https://postgr.es/c/23c8c0c8f) [&sect;](https://postgr.es/c/2e8a0edc2) [&sect;](https://postgr.es/c/e5cca6288)

   Allow [`pg_stat_reset_shared()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-funcs-table) (with no arguments) and pg_stat_reset_shared(`NULL`) to reset all shared statistics. Allow pg_stat_reset_shared('slru') and [`pg_stat_reset_slru()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-funcs-table) (with no arguments) to reset SLRU statistics, which was already possible with pg_stat_reset_slru(NULL).
-  Add log messages related to WAL recovery from backups (Andres Freund) [&sect;](https://postgr.es/c/1d35f705e)
-  Add [log_connections](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-connections) log line for `trust` connections (Jacob Champion) [&sect;](https://postgr.es/c/e48b19c5d)
-  Add log message to report walsender acquisition and release of replication slots (Bharath Rupireddy) [&sect;](https://postgr.es/c/7c3fb505b)

   This is enabled by the server variable [log_replication_commands](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-replication-commands).
-  Add system view [`pg_wait_events`](../../internals/system-views/pg_wait_events.md#view-pg-wait-events) that reports wait event types (Bertrand Drouvot) [&sect;](https://postgr.es/c/1e68e43d3)

   This is useful for adding descriptions to wait events reported in [`pg_stat_activity`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-activity-view).
-  Add [wait events](../../internals/system-views/pg_wait_events.md#view-pg-wait-events) for checkpoint delays (Thomas Munro) [&sect;](https://postgr.es/c/0013ba290)
-  Allow vacuum to report the progress of index processing (Sami Imseih) [&sect;](https://postgr.es/c/46ebdfe16)

   This appears in system view [`pg_stat_progress_vacuum`](../../server-administration/monitoring-database-activity/progress-reporting.md#pg-stat-progress-vacuum-view) columns `indexes_total` and `indexes_processed`.


##### Privileges { #release-17-privileges }


-  Allow granting the right to perform maintenance operations (Nathan Bossart) [&sect;](https://postgr.es/c/ecb0fd337)

   The permission can be granted on a per-table basis using the [`MAINTAIN`](../../the-sql-language/data-definition/privileges.md#ddl-priv-maintain) privilege and on a per-role basis via the [`pg_maintain`](../../server-administration/database-roles/predefined-roles.md#predefined-roles) predefined role. Permitted operations are `VACUUM`, `ANALYZE`, `REINDEX`, `REFRESH MATERIALIZED VIEW`, `CLUSTER`, and `LOCK TABLE`.
-  Allow roles with [`pg_monitor`](../../server-administration/database-roles/predefined-roles.md#predefined-roles) membership to execute [`pg_current_logfile()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-session-table) (Pavlo Golub, Nathan Bossart) [&sect;](https://postgr.es/c/8d8afd48d)


##### Server Configuration { #release-17-server-config }


-  Add system variable [allow_alter_system](../../server-administration/server-configuration/version-and-platform-compatibility.md#guc-allow-alter-system) to disallow [`ALTER SYSTEM`](../../reference/sql-commands/alter-system.md#sql-altersystem) (Jelte Fennema-Nio, Gabriele Bartolini) [&sect;](https://postgr.es/c/d3ae2a24f)
-  Allow [`ALTER SYSTEM`](../../reference/sql-commands/alter-system.md#sql-altersystem) to set unrecognized custom server variables (Tom Lane) [&sect;](https://postgr.es/c/2d870b4ae)

   This is also possible with [`GRANT ON PARAMETER`](../../reference/sql-commands/grant.md#sql-grant).
-  Add server variable [transaction_timeout](../../server-administration/server-configuration/client-connection-defaults.md#guc-transaction-timeout) to restrict the duration of transactions (Andrey Borodin, Japin Li, Junwang Zhao, Alexander Korotkov) [&sect;](https://postgr.es/c/51efe38cb) [&sect;](https://postgr.es/c/bf82f4379) [&sect;](https://postgr.es/c/28e858c0f)
-  Add a builtin platform-independent collation provider (Jeff Davis) [&sect;](https://postgr.es/c/2d819a08a) [&sect;](https://postgr.es/c/846311051) [&sect;](https://postgr.es/c/f69319f2f) [&sect;](https://postgr.es/c/9acae56ce)

   This supports `C` and `C.UTF-8` collations.
-  Add server variable [huge_pages_status](../../server-administration/server-configuration/preset-options.md#guc-huge-pages-status) to report the use of huge pages by Postgres (Justin Pryzby) [&sect;](https://postgr.es/c/a14354cac)

   This is useful when [huge_pages](../../server-administration/server-configuration/resource-consumption.md#guc-huge-pages) is set to `try`.
-  Add server variable to disable event triggers (Daniel Gustafsson) [&sect;](https://postgr.es/c/7750fefdb)

   The setting, [event_triggers](../../server-administration/server-configuration/client-connection-defaults.md#guc-event-triggers), allows for the temporary disabling of event triggers for debugging.
-  Allow the [SLRU](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-slru-view) cache sizes to be configured (Andrey Borodin, Dilip Kumar, Alvaro Herrera) [&sect;](https://postgr.es/c/53c2a97a9)

   The new server variables are [commit_timestamp_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-commit-timestamp-buffers), [multixact_member_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-multixact-member-buffers), [multixact_offset_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-multixact-offset-buffers), [notify_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-notify-buffers), [serializable_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-serializable-buffers), [subtransaction_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-subtransaction-buffers), and [transaction_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-transaction-buffers). [commit_timestamp_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-commit-timestamp-buffers), [transaction_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-transaction-buffers), and [subtransaction_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-subtransaction-buffers) scale up automatically with [shared_buffers](../../server-administration/server-configuration/resource-consumption.md#guc-shared-buffers).


##### Streaming Replication and Recovery { #release-17-replication }


-  Add support for incremental file system backup (Robert Haas, Jakub Wartak, Tomas Vondra) [&sect;](https://postgr.es/c/dc2123400) [&sect;](https://postgr.es/c/f8ce4ed78)

   Incremental backups can be created using [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup)'s new `--incremental` option. The new application [pg_combinebackup](../../reference/postgresql-client-applications/pg_combinebackup.md#app-pgcombinebackup) allows manipulation of base and incremental file system backups.
-  Allow the creation of WAL summarization files (Robert Haas, Nathan Bossart, Hubert Depesz Lubaczewski) [&sect;](https://postgr.es/c/174c48050) [&sect;](https://postgr.es/c/d97ef756a) [&sect;](https://postgr.es/c/f896057e4) [&sect;](https://postgr.es/c/d9ef650fc)

   These files record the block numbers that have changed within an [LSN](../../the-sql-language/data-types/pg_lsn-type.md#datatype-pg-lsn) range and are useful for incremental file system backups. This is controlled by the server variables [summarize_wal](../../server-administration/server-configuration/write-ahead-log.md#guc-summarize-wal) and [wal_summary_keep_time](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-summary-keep-time), and introspected with [`pg_available_wal_summaries()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-wal-summary), `pg_wal_summary_contents()`, and `pg_get_wal_summarizer_state()`.
-  Add the system identifier to file system [backup manifest](../../internals/backup-manifest-format/index.md#backup-manifest-format) files (Amul Sul) [&sect;](https://postgr.es/c/2041bc427)

   This helps detect invalid WAL usage.
-  Allow connection string value `dbname` to be written when [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) writes connection information to `postgresql.auto.conf` (Vignesh C, Hayato Kuroda) [&sect;](https://postgr.es/c/a145f424d)
-  Add column [`pg_replication_slots`.`invalidation_reason`](../../internals/system-views/pg_replication_slots.md#view-pg-replication-slots) to report the reason for invalid slots (Shveta Malik, Bharath Rupireddy) [&sect;](https://postgr.es/c/007693f2a) [&sect;](https://postgr.es/c/6ae701b43)
-  Add column [`pg_replication_slots`.`inactive_since`](../../internals/system-views/pg_replication_slots.md#view-pg-replication-slots) to report slot inactivity duration (Bharath Rupireddy) [&sect;](https://postgr.es/c/a11f330b5) [&sect;](https://postgr.es/c/6d49c8d4b) [&sect;](https://postgr.es/c/6f132ed69)
-  Add function [`pg_sync_replication_slots()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-replication-table) to synchronize logical replication slots (Hou Zhijie, Shveta Malik, Ajin Cherian, Peter Eisentraut) [&sect;](https://postgr.es/c/ddd5f4f54) [&sect;](https://postgr.es/c/7a424ece4)
-  Add the `failover` property to the [replication protocol](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication) (Hou Zhijie, Shveta Malik) [&sect;](https://postgr.es/c/732924043)


##### [Logical Replication] { #release-17-logical }


-  Add application [pg_createsubscriber](../../reference/postgresql-server-applications/pg_createsubscriber.md#app-pgcreatesubscriber) to create a logical replica from a physical standby server (Euler Taveira) [&sect;](https://postgr.es/c/d44032d01)
-  Have [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) migrate valid logical slots and subscriptions (Hayato Kuroda, Hou Zhijie, Vignesh C, Julien Rouhaud, Shlok Kyal) [&sect;](https://postgr.es/c/29d0a77fa) [&sect;](https://postgr.es/c/9a17be1e2)

   This allows logical replication to continue quickly after the upgrade. This only works for old PostgreSQL clusters that are version 17 or later.
-  Enable the failover of [logical slots](../../server-administration/logical-replication/subscription.md#logical-replication-subscription-slot) (Hou Zhijie, Shveta Malik, Ajin Cherian) [&sect;](https://postgr.es/c/c393308b6)

   This is controlled by an optional fifth argument to [`pg_create_logical_replication_slot()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-replication-table).
-  Add server variable [sync_replication_slots](../../server-administration/server-configuration/replication.md#guc-sync-replication-slots) to enable failover logical slot synchronization (Shveta Malik, Hou Zhijie, Peter Smith) [&sect;](https://postgr.es/c/93db6cbda) [&sect;](https://postgr.es/c/60c07820d)
-  Add logical replication failover control to [`CREATE/ALTER SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription) (Shveta Malik, Hou Zhijie, Ajin Cherian) [&sect;](https://postgr.es/c/776621a5e) [&sect;](https://postgr.es/c/22f7e61a6)
-  Allow the application of logical replication changes to use [hash](../../internals/built-in-index-access-methods/hash-indexes.md#hash-index) indexes on the subscriber (Hayato Kuroda) [&sect;](https://postgr.es/c/edca34243)

   Previously only [btree](../../internals/built-in-index-access-methods/b-tree-indexes.md#btree) indexes could be used for this purpose.
-  Improve [logical decoding](../../server-programming/logical-decoding/index.md#logicaldecoding) performance in cases where there are many subtransactions (Masahiko Sawada) [&sect;](https://postgr.es/c/5bec1d6bc)
-  Restart apply workers if subscription owner's superuser privileges are revoked (Vignesh C) [&sect;](https://postgr.es/c/79243de13)

   This forces reauthentication.
-  Add `flush` option to [`pg_logical_emit_message()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-replication-table) (Michael Paquier) [&sect;](https://postgr.es/c/173b56f1e)

   This makes the message durable.
-  Allow specification of physical standbys that must be synchronized before they are visible to subscribers (Hou Zhijie, Shveta Malik) [&sect;](https://postgr.es/c/bf279ddd1) [&sect;](https://postgr.es/c/0f934b073)

   The new server variable is [synchronized_standby_slots](../../server-administration/server-configuration/replication.md#guc-synchronized-standby-slots).
-  Add worker type column to [`pg_stat_subscription`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-subscription) (Peter Smith) [&sect;](https://postgr.es/c/13aeaf079)


#### Utility Commands { #release-17-utility }


-  Add new [`COPY`](../../reference/sql-commands/copy.md#sql-copy) option `ON_ERROR ignore` to discard error rows (Damir Belyalov, Atsushi Torikoshi, Alex Shulgin, Jian He, Yugo Nagata) [&sect;](https://postgr.es/c/9e2d87011) [&sect;](https://postgr.es/c/b725b7eec) [&sect;](https://postgr.es/c/40bbc8cf0) [&sect;](https://postgr.es/c/a6d0fa5ef)

   The default behavior is `ON_ERROR stop`.
-  Add new `COPY` option `LOG_VERBOSITY` which reports `COPY FROM` ignored error rows (Bharath Rupireddy) [&sect;](https://postgr.es/c/f5a227895)
-  Allow `COPY FROM` to report the number of skipped rows during processing (Atsushi Torikoshi) [&sect;](https://postgr.es/c/729439607)

   This appears in system view column [`pg_stat_progress_copy`.`tuples_skipped`](../../server-administration/monitoring-database-activity/progress-reporting.md#copy-progress-reporting).
-  In `COPY FROM`, allow easy specification that all columns should be forced null or not null (Zhang Mingli) [&sect;](https://postgr.es/c/f6d4c9cf1)
-  Allow partitioned tables to have identity columns (Ashutosh Bapat) [&sect;](https://postgr.es/c/699586315)
-  Allow [exclusion constraints](../../the-sql-language/data-definition/constraints.md#ddl-constraints-exclusion) on partitioned tables (Paul A. Jungwirth) [&sect;](https://postgr.es/c/8c852ba9a)

   As long as exclusion constraints compare partition key columns for equality, other columns can use exclusion constraint-specific comparisons.
-  Add clearer [`ALTER TABLE`](../../reference/sql-commands/alter-table.md#sql-altertable) method to set a column to the default statistics target (Peter Eisentraut) [&sect;](https://postgr.es/c/4f622503d)

   The new syntax is `ALTER TABLE ... SET STATISTICS DEFAULT`; using `SET STATISTICS -1` is still supported.
-  Allow `ALTER TABLE` to change a column's generation expression (Amul Sul) [&sect;](https://postgr.es/c/5d06e99a3)

   The syntax is `ALTER TABLE ... ALTER COLUMN ... SET EXPRESSION`.
-  Allow specification of [table access methods](../../internals/table-access-method-interface-definition.md#tableam) on partitioned tables (Justin Pryzby, Soumyadeep Chakraborty, Michael Paquier) [&sect;](https://postgr.es/c/374c7a229) [&sect;](https://postgr.es/c/e2395cdbe)
-  Add `DEFAULT` setting for `ALTER TABLE .. SET ACCESS METHOD` (Michael Paquier) [&sect;](https://postgr.es/c/d61a6cad6)
-  Add support for [event triggers](../../reference/sql-commands/create-event-trigger.md#sql-createeventtrigger) that fire at connection time (Konstantin Knizhnik, Mikhail Gribkov) [&sect;](https://postgr.es/c/e83d1b0c4)
-  Add event trigger support for [`REINDEX`](../../reference/sql-commands/reindex.md#sql-reindex) (Garrett Thornburg, Jian He) [&sect;](https://postgr.es/c/f21848de2)
-  Allow parenthesized syntax for [`CLUSTER`](../../reference/sql-commands/cluster.md#sql-cluster) options if a table name is not specified (Nathan Bossart) [&sect;](https://postgr.es/c/cdaedfc96)


##### [`EXPLAIN`] { #release-17-explain }


-  Allow `EXPLAIN` to report optimizer memory usage (Ashutosh Bapat) [&sect;](https://postgr.es/c/5de890e36)

   The option is called `MEMORY`.
-  Add `EXPLAIN` option `SERIALIZE` to report the cost of converting data for network transmission (Stepan Rutz, Matthias van de Meent) [&sect;](https://postgr.es/c/06286709e)
-  Add local I/O block read/write timing statistics to `EXPLAIN`'s `BUFFERS` output (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/295c36c0c)
-  Improve `EXPLAIN`'s display of SubPlan nodes and output parameters (Tom Lane, Dean Rasheed) [&sect;](https://postgr.es/c/fd0398fcb)
-  Add JIT `deform_counter` details to `EXPLAIN` (Dmitry Dolgov) [&sect;](https://postgr.es/c/5a3423ad8)


#### Data Types { #release-17-datatypes }


-  Allow the `interval` data type to support `+/-infinity` values (Joseph Koshakow, Jian He, Ashutosh Bapat) [&sect;](https://postgr.es/c/519fc1bd9)
-  Allow the use of an [`ENUM`](../../the-sql-language/data-types/enumerated-types.md#datatype-enum) added via [`ALTER TYPE`](../../reference/sql-commands/alter-type.md#sql-altertype) if the type was created in the same transaction (Tom Lane) [&sect;](https://postgr.es/c/af1d39584)

   This was previously disallowed.


#### [MERGE] { #release-17-merge }


-  Allow `MERGE` to modify updatable views (Dean Rasheed) [&sect;](https://postgr.es/c/5f2e179bd)
-  Add `WHEN NOT MATCHED BY SOURCE` to `MERGE` (Dean Rasheed) [&sect;](https://postgr.es/c/0294df2f1)

   `WHEN NOT MATCHED` on target rows was already supported.
-  Allow `MERGE` to use the `RETURNING` clause (Dean Rasheed) [&sect;](https://postgr.es/c/c649fa24a)

   The new `RETURNING` function `merge_action()` reports on the DML that generated the row.


#### Functions { #release-17-functions }


-  Add function [`JSON_TABLE()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-table) to convert `JSON` data to a table representation (Nikita Glukhov, Teodor Sigaev, Oleg Bartunov, Alexander Korotkov, Andrew Dunstan, Amit Langote, Jian He) [&sect;](https://postgr.es/c/de3600452) [&sect;](https://postgr.es/c/bb766cde6)

   This function can be used in the `FROM` clause of `SELECT` queries as a tuple source.
-  Add SQL/JSON constructor functions [`JSON()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-json-creation-table), `JSON_SCALAR()`, and `JSON_SERIALIZE()` (Nikita Glukhov, Teodor Sigaev, Oleg Bartunov, Alexander Korotkov, Andrew Dunstan, Amit Langote) [&sect;](https://postgr.es/c/03734a7fe)
-  Add SQL/JSON query functions [`JSON_EXISTS()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-querying), `JSON_QUERY()`, and `JSON_VALUE()` (Nikita Glukhov, Teodor Sigaev, Oleg Bartunov, Alexander Korotkov, Andrew Dunstan, Amit Langote, Peter Eisentraut, Jian He) [&sect;](https://postgr.es/c/aaaf9449e) [&sect;](https://postgr.es/c/1edb3b491) [&sect;](https://postgr.es/c/6185c9737) [&sect;](https://postgr.es/c/c0fc07518) [&sect;](https://postgr.es/c/ef744ebb7)
-  Add [jsonpath](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators) methods to convert `JSON` values to other `JSON` data types (Jeevan Chalke) [&sect;](https://postgr.es/c/66ea94e8e)

   The jsonpath methods are `.bigint()`, `.boolean()`, `.date()`, `.decimal([precision [, scale]])`, `.integer()`, `.number()`, `.string()`, `.time()`, `.time_tz()`, `.timestamp()`, and `.timestamp_tz()`.
-  Add [`to_timestamp()`](../../the-sql-language/functions-and-operators/data-type-formatting-functions.md#functions-formatting-table) time zone format specifiers (Tom Lane) [&sect;](https://postgr.es/c/8ba6fdf90)

   `TZ` accepts time zone abbreviations or numeric offsets, while `OF` accepts only numeric offsets.
-  Allow the session [time zone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) to be specified by `AT LOCAL` (Vik Fearing) [&sect;](https://postgr.es/c/97957fdba)

   This is useful when converting adding and removing time zones from time stamps values, rather than specifying the literal session time zone.
-  Add functions [`uuid_extract_timestamp()`](../../the-sql-language/functions-and-operators/uuid-functions.md#functions-uuid) and `uuid_extract_version()` to return UUID information (Andrey Borodin) [&sect;](https://postgr.es/c/794f10f6b)
-  Add functions to generate random numbers in a specified range (Dean Rasheed) [&sect;](https://postgr.es/c/e6341323a)

   The functions are [`random(min, max)`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-random-table) and they take values of type `integer`, `bigint`, and `numeric`.
-  Add functions to convert integers to binary and octal strings (Eric Radman, Nathan Bossart) [&sect;](https://postgr.es/c/260a1f18d)

   The functions are [`to_bin()`](../../the-sql-language/functions-and-operators/string-functions-and-operators.md#functions-string-other) and `to_oct()`.
-  Add Unicode informational functions (Jeff Davis) [&sect;](https://postgr.es/c/a02b37fc0)

   Function [`unicode_version()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-version) returns the Unicode version, `icu_unicode_version()` returns the ICU version, and `unicode_assigned()` returns if the characters are assigned Unicode codepoints.
-  Add function [`xmltext()`](../../the-sql-language/functions-and-operators/xml-functions.md#functions-producing-xml-xmltext) to convert text to a single `XML` text node (Jim Jones) [&sect;](https://postgr.es/c/526fe0d79)
-  Add function [`to_regtypemod()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-catalog-table) to return the type modifier of a type specification (David Wheeler, Erik Wienhold) [&sect;](https://postgr.es/c/1218ca995)
-  Add [`pg_basetype()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-catalog-table) function to return a domain's base type (Steve Chavez) [&sect;](https://postgr.es/c/b154d8a6d)
-  Add function [`pg_column_toast_chunk_id()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-dbsize) to return a value's [TOAST](../../internals/database-physical-storage/toast.md#storage-toast) identifier (Yugo Nagata) [&sect;](https://postgr.es/c/d1162cfda)

   This returns `NULL` if the value is not stored in TOAST.


#### [PL/pgSQL] { #release-17-plpgsql }


-  Allow plpgsql [`%TYPE`](../../server-programming/pl-pgsql-sql-procedural-language/declarations.md#plpgsql-declaration-type) and `%ROWTYPE` specifications to represent arrays of non-array types (Quan Zongliang, Pavel Stehule) [&sect;](https://postgr.es/c/5e8674dc8)
-  Allow plpgsql `%TYPE` specification to reference composite column (Tom Lane) [&sect;](https://postgr.es/c/43b46aae1)


#### [libpq] { #release-17-libpq }


-  Add libpq function to change role passwords (Joe Conway) [&sect;](https://postgr.es/c/a7be2a6c2)

   The new function, [`PQchangePassword()`](../../client-interfaces/libpq-c-library/miscellaneous-functions.md#libpq-PQchangePassword), hashes the new password before sending it to the server.
-  Add libpq functions to close portals and prepared statements (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/28b572656)

   The functions are [`PQclosePrepared()`](../../client-interfaces/libpq-c-library/command-execution-functions.md#libpq-PQclosePrepared), [`PQclosePortal()`](../../client-interfaces/libpq-c-library/command-execution-functions.md#libpq-PQclosePortal), [`PQsendClosePrepared()`](../../client-interfaces/libpq-c-library/asynchronous-command-processing.md#libpq-PQsendClosePrepared), and [`PQsendClosePortal()`](../../client-interfaces/libpq-c-library/asynchronous-command-processing.md#libpq-PQsendClosePortal).
-  Add libpq API which allows for blocking and non-blocking [cancel requests](../../client-interfaces/libpq-c-library/canceling-queries-in-progress.md#libpq-cancel), with encryption if already in use (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/61461a300)

   Previously only blocking, unencrypted cancel requests were supported.
-  Add libpq function [`PQsocketPoll()`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-PQsocketPoll) to allow polling of network sockets (Tristan Partin, Tom Lane) [&sect;](https://postgr.es/c/f5e4dedfa) [&sect;](https://postgr.es/c/105024a47)
-  Add libpq function [`PQsendPipelineSync()`](../../client-interfaces/libpq-c-library/pipeline-mode.md#libpq-PQsendPipelineSync) to send a pipeline synchronization point (Anton Kirilov) [&sect;](https://postgr.es/c/4794c2d31)

   This is similar to [`PQpipelineSync()`](../../client-interfaces/libpq-c-library/pipeline-mode.md#libpq-PQpipelineSync) but it does not flush to the server unless the size threshold of the output buffer is reached.
-  Add libpq function [`PQsetChunkedRowsMode()`](../../client-interfaces/libpq-c-library/retrieving-query-results-in-chunks.md#libpq-PQsetChunkedRowsMode) to allow retrieval of results in chunks (Daniel V&eacute;rit&eacute;) [&sect;](https://postgr.es/c/4643a2b26)
-  Allow TLS connections without requiring a network round-trip negotiation (Greg Stark, Heikki Linnakangas, Peter Eisentraut, Michael Paquier, Daniel Gustafsson) [&sect;](https://postgr.es/c/d39a49c1e) [&sect;](https://postgr.es/c/91044ae4b) [&sect;](https://postgr.es/c/44e27f0a6) [&sect;](https://postgr.es/c/d80f2ce29) [&sect;](https://postgr.es/c/03a0e0d4b) [&sect;](https://postgr.es/c/17a834a04) [&sect;](https://postgr.es/c/407e0b023) [&sect;](https://postgr.es/c/fb5718f35)

   This is enabled with the client-side option [`sslnegotiation=direct`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-sslnegotiation), requires ALPN, and only works on PostgreSQL 17 and later servers.


#### [app-psql] { #release-17-psql }


-  Improve psql display of default and empty privileges (Erik Wienhold, Laurenz Albe) [&sect;](https://postgr.es/c/d1379ebf4)

   Command `\dp` now displays `(none)` for empty privileges; default still displays as empty.
-  Have backslash commands honor `\pset null` (Erik Wienhold, Laurenz Albe) [&sect;](https://postgr.es/c/d1379ebf4)

   Previously `\pset null` was ignored.
-  Allow psql's `\watch` to stop after a minimum number of rows returned (Greg Sabino Mullane) [&sect;](https://postgr.es/c/f347ec76e)

   The parameter is `min_rows`.
-  Allow psql connection attempts to be canceled with control-C (Tristan Partin) [&sect;](https://postgr.es/c/cafe10565)
-  Allow psql to honor `FETCH_COUNT` for non-`SELECT` queries (Daniel V&eacute;rit&eacute;) [&sect;](https://postgr.es/c/90f517821)
-  Improve psql tab completion (Dagfinn Ilmari Manns&aring;ker, Gilles Darold, Christoph Heiss, Steve Chavez, Vignesh C, Pavel Borisov, Jian He) [&sect;](https://postgr.es/c/c951e9042) [&sect;](https://postgr.es/c/d16eb83ab) [&sect;](https://postgr.es/c/cd3424748) [&sect;](https://postgr.es/c/816f10564) [&sect;](https://postgr.es/c/927332b95) [&sect;](https://postgr.es/c/f1bb9284f) [&sect;](https://postgr.es/c/304b6b1a6) [&sect;](https://postgr.es/c/2800fbb2b)


#### Server Applications { #release-17-server-apps }


-  Add application [pg_walsummary](../../reference/postgresql-server-applications/pg_walsummary.md#app-pgwalsummary) to dump WAL summary files (Robert Haas) [&sect;](https://postgr.es/c/ee1bfd168)
-  Allow [pg_dump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump)'s large objects to be restorable in batches (Tom Lane) [&sect;](https://postgr.es/c/a45c78e32)

   This allows the restoration of many large objects to avoid transaction limits and to be restored in parallel.
-  Add pg_dump option `--exclude-extension` (Ayush Vatsa) [&sect;](https://postgr.es/c/522ed12f7)
-  Allow [pg_dump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump), [pg_dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall), and [pg_restore](../../reference/postgresql-client-applications/pg_restore.md#app-pgrestore) to specify include/exclude objects in a file (Pavel Stehule, Daniel Gustafsson) [&sect;](https://postgr.es/c/a5cf808be)

   The option is called `--filter`.
-  Add the `--sync-method` parameter to several client applications (Justin Pryzby, Nathan Bossart) [&sect;](https://postgr.es/c/8c16ad3b4)

   The applications are [initdb](../../reference/postgresql-server-applications/initdb.md#app-initdb), [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup), [pg_checksums](../../reference/postgresql-server-applications/pg_checksums.md#app-pgchecksums), [pg_dump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump), [pg_rewind](../../reference/postgresql-server-applications/pg_rewind.md#app-pgrewind), and [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade).
-  Add [pg_restore](../../reference/postgresql-client-applications/pg_restore.md#app-pgrestore) option `--transaction-size` to allow object restores in transaction batches (Tom Lane) [&sect;](https://postgr.es/c/959b38d77)

   This allows the performance benefits of transaction batches without the problems of excessively large transaction blocks.
-  Change [pgbench](../../reference/postgresql-client-applications/pgbench.md#pgbench) debug mode option from `-d` to `--debug` (Greg Sabino Mullane) [&sect;](https://postgr.es/c/3ff01b2b6)

   Option `-d` is now used for the database name, and the new `--dbname` option can be used as well.
-  Add pgbench option `--exit-on-abort` to exit after any client aborts (Yugo Nagata) [&sect;](https://postgr.es/c/3c662643c)
-  Add pgbench command `\syncpipeline` to allow sending of sync messages (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/94edfe250)
-  Allow [pg_archivecleanup](../../reference/postgresql-server-applications/pg_archivecleanup.md#pgarchivecleanup) to remove backup history files (Atsushi Torikoshi) [&sect;](https://postgr.es/c/3f8c98d0b)

   The option is `--clean-backup-history`.
-  Add some long options to pg_archivecleanup (Atsushi Torikoshi) [&sect;](https://postgr.es/c/dd7c60f19)

   The long options are `--debug`, `--dry-run`, and `--strip-extension`.
-  Allow [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) and [pg_receivewal](../../reference/postgresql-client-applications/pg_receivewal.md#app-pgreceivewal) to use dbname in their connection specification (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/cca97ce6a)

   This is useful for connection poolers that are sensitive to the database name.
-  Add [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) option `--copy-file-range` (Thomas Munro) [&sect;](https://postgr.es/c/d93627bcb)

   This is supported on `Linux` and `FreeBSD`.
-  Allow [reindexdb](../../reference/postgresql-client-applications/reindexdb.md#app-reindexdb) `--index` to process indexes from different tables in parallel (Maxim Orlov, Svetlana Derevyanko, Alexander Korotkov) [&sect;](https://postgr.es/c/47f99a407)
-  Allow [reindexdb](../../reference/postgresql-client-applications/reindexdb.md#app-reindexdb), [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb), and [clusterdb](../../reference/postgresql-client-applications/clusterdb.md#app-clusterdb) to process objects in all databases matching a pattern (Nathan Bossart) [&sect;](https://postgr.es/c/24c928ad9) [&sect;](https://postgr.es/c/648928c79) [&sect;](https://postgr.es/c/1b49d56d3)

   The new option `--all` controls this behavior.


#### Source Code { #release-17-source-code }


-  Remove support for OpenSSL 1.0.1 (Michael Paquier) [&sect;](https://postgr.es/c/8e278b657)
-  Allow tests to pass in OpenSSL FIPS mode (Peter Eisentraut) [&sect;](https://postgr.es/c/284cbaea7) [&sect;](https://postgr.es/c/3c44e7d8d)
-  Use CPU AVX-512 instructions for bit counting (Paul Amonson, Nathan Bossart, Ants Aasma) [&sect;](https://postgr.es/c/792752af4) [&sect;](https://postgr.es/c/41c51f0c6)
-  Require LLVM version 10 or later (Thomas Munro) [&sect;](https://postgr.es/c/820b5af73)
-  Use native CRC instructions on 64-bit LoongArch CPUs (Xudong Yang) [&sect;](https://postgr.es/c/4d14ccd6a)
-  Remove `AIX` support (Heikki Linnakangas) [&sect;](https://postgr.es/c/0b16bb877)
-  Remove the Microsoft Visual Studio-specific PostgreSQL build option (Michael Paquier) [&sect;](https://postgr.es/c/1301c80b2)

   Meson is now the only available method for Visual Studio builds.
-  Remove configure option `--disable-thread-safety` (Thomas Munro, Heikki Linnakangas) [&sect;](https://postgr.es/c/68a4b58ec) [&sect;](https://postgr.es/c/ce0b0fa3e)

   We now assume all supported platforms have sufficient thread support.
-  Remove configure option `--with-CC` (Heikki Linnakangas) [&sect;](https://postgr.es/c/1c1eec0f2)

   Setting the `CC` environment variable is now the only supported method for specifying the compiler.
-  User-defined data type receive functions will no longer receive their data null-terminated (David Rowley) [&sect;](https://postgr.es/c/f0efa5aec)
-  Add incremental `JSON` parser for use with huge `JSON` documents (Andrew Dunstan) [&sect;](https://postgr.es/c/3311ea86e)
-  Convert top-level `README` file to Markdown (Nathan Bossart) [&sect;](https://postgr.es/c/363eb0599)
-  Remove no longer needed top-level `INSTALL` file (Tom Lane) [&sect;](https://postgr.es/c/e2b73f4a4)
-  Remove make's `distprep` option (Peter Eisentraut) [&sect;](https://postgr.es/c/721856ff2)
-  Add make support for Android shared libraries (Peter Eisentraut) [&sect;](https://postgr.es/c/79b03dbb3)
-  Add backend support for injection points (Michael Paquier) [&sect;](https://postgr.es/c/d86d20f0b) [&sect;](https://postgr.es/c/37b369dc6) [&sect;](https://postgr.es/c/f587338de) [&sect;](https://postgr.es/c/bb93640a6)

   This is used for server debugging and they must be enabled at server compile time.
-  Add dynamic shared memory registry (Nathan Bossart) [&sect;](https://postgr.es/c/8b2bcf3f2)

   This allows shared libraries which are not initialized at startup to coordinate dynamic shared memory access.
-  Fix `emit_log_hook` to use the same time value as other log records for the same query (Kambam Vinay, Michael Paquier) [&sect;](https://postgr.es/c/2a217c371)
-  Improve documentation for using `jsonpath` for predicate checks (David Wheeler) [&sect;](https://postgr.es/c/7014c9a4b)


#### Additional Modules { #release-17-modules }


-  Allow joins with non-join qualifications to be pushed down to foreign servers and custom scans (Richard Guo, Etsuro Fujita) [&sect;](https://postgr.es/c/9e9931d2b)

   Foreign data wrappers and custom scans will need to be modified to handle these cases.
-  Allow pushdown of `EXISTS` and `IN` subqueries to [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) foreign servers (Alexander Pyhalov) [&sect;](https://postgr.es/c/824dbea3e)
-  Increase the default foreign data wrapper tuple cost (David Rowley, Umair Shahid) [&sect;](https://postgr.es/c/cac169d68) [&sect;](https://postgr.es/c/f7f694b21)

   This value is used by the optimizer.
-  Allow [dblink](../additional-supplied-modules-and-extensions/dblink-connect-to-other-postgresql-databases.md#dblink) database operations to be interrupted (Noah Misch) [&sect;](https://postgr.es/c/d3c5f37dd)
-  Allow the creation of hash indexes on [ltree](../additional-supplied-modules-and-extensions/ltree-hierarchical-tree-like-data-type.md#ltree) columns (Tommy Pavlicek) [&sect;](https://postgr.es/c/485f0aa85)

   This also enables hash join and hash aggregation on ltree columns.
-  Allow [unaccent](../additional-supplied-modules-and-extensions/unaccent-a-text-search-dictionary-which-removes-diacritics.md#unaccent) character translation rules to contain whitespace and quotes (Michael Paquier) [&sect;](https://postgr.es/c/59f47fb98)

   The syntax for the `unaccent.rules` file has changed.
-  Allow [amcheck](../additional-supplied-modules-and-extensions/amcheck-tools-to-verify-table-and-index-consistency.md#amcheck) to check for unique constraint violations using new option `--checkunique` (Anastasia Lubennikova, Pavel Borisov, Maxim Orlov) [&sect;](https://postgr.es/c/5ae208720)
-  Allow [citext](../additional-supplied-modules-and-extensions/citext-a-case-insensitive-character-string-type.md#citext) tests to pass in OpenSSL FIPS mode (Peter Eisentraut) [&sect;](https://postgr.es/c/3c551ebed)
-  Allow [pgcrypto](../additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto) tests to pass in OpenSSL FIPS mode (Peter Eisentraut) [&sect;](https://postgr.es/c/795592865)
-  Remove some unused [SPI](../../server-programming/server-programming-interface/index.md#spi) macros (Bharath Rupireddy) [&sect;](https://postgr.es/c/75680c3d8)
-  Allow [`ALTER OPERATOR`](../../reference/sql-commands/alter-operator.md#sql-alteroperator) to set more optimization attributes (Tommy Pavlicek) [&sect;](https://postgr.es/c/2b5154bea)

   This is useful for extensions.
-  Allow extensions to define [custom wait events](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-wait-events) (Masahiro Ikeda) [&sect;](https://postgr.es/c/c9af05465) [&sect;](https://postgr.es/c/c8e318b1b) [&sect;](https://postgr.es/c/d61f2538a) [&sect;](https://postgr.es/c/c789f0f6c)

   Custom wait events have been added to [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) and [dblink](../additional-supplied-modules-and-extensions/dblink-connect-to-other-postgresql-databases.md#dblink).
-  Add [pg_buffercache](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache) function `pg_buffercache_evict()` to allow shared buffer eviction (Palak Chaturvedi, Thomas Munro) [&sect;](https://postgr.es/c/13453eedd)

   This is useful for testing.


##### [pg_stat_statements] { #release-17-pgstatstatements }


-  Replace [`CALL`](../../reference/sql-commands/call.md#sql-call) parameters in pg_stat_statements with placeholders (Sami Imseih) [&sect;](https://postgr.es/c/11c34b342)
-  Replace savepoint names stored in `pg_stat_statements` with placeholders (Greg Sabino Mullane) [&sect;](https://postgr.es/c/31de7e60d)

   This greatly reduces the number of entries needed to record [`SAVEPOINT`](../../reference/sql-commands/savepoint.md#sql-savepoint), [`RELEASE SAVEPOINT`](../../reference/sql-commands/release-savepoint.md#sql-release-savepoint), and [`ROLLBACK TO SAVEPOINT`](../../reference/sql-commands/rollback-to-savepoint.md#sql-rollback-to) commands.
-  Replace the two-phase commit GIDs stored in `pg_stat_statements` with placeholders (Michael Paquier) [&sect;](https://postgr.es/c/638d42a3c)

   This greatly reduces the number of entries needed to record [`PREPARE TRANSACTION`](../../reference/sql-commands/prepare-transaction.md#sql-prepare-transaction), [`COMMIT PREPARED`](../../reference/sql-commands/commit-prepared.md#sql-commit-prepared), and [`ROLLBACK PREPARED`](../../reference/sql-commands/rollback-prepared.md#sql-rollback-prepared).
-  Track [`DEALLOCATE`](../../reference/sql-commands/deallocate.md#sql-deallocate) in `pg_stat_statements` (Dagfinn Ilmari Manns&aring;ker, Michael Paquier) [&sect;](https://postgr.es/c/bb45156f3)

   `DEALLOCATE` names are stored in `pg_stat_statements` as placeholders.
-  Add local I/O block read/write timing statistics columns of `pg_stat_statements` (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/295c36c0c) [&sect;](https://postgr.es/c/5147ab1dd)

   The new columns are `local_blk_read_time` and `local_blk_write_time`.
-  Add JIT deform_counter details to `pg_stat_statements` (Dmitry Dolgov) [&sect;](https://postgr.es/c/5a3423ad8)
-  Add optional fourth argument (`minmax_only`) to `pg_stat_statements_reset()` to allow for the resetting of only min/max statistics (Andrei Zubkov) [&sect;](https://postgr.es/c/dc9f8a798)

   This argument defaults to `false`.
-  Add `pg_stat_statements` columns `stats_since` and `minmax_stats_since` to track entry creation time and last min/max reset time (Andrei Zubkov) [&sect;](https://postgr.es/c/dc9f8a798)


### Acknowledgments { #release-17-acknowledgements }


 The following individuals (in alphabetical order) have contributed to this release as patch authors, committers, reviewers, testers, or reporters of issues.


- Abhijit Menon-Sen
- Adnan Dautovic
- Aidar Imamov
- Ajin Cherian
- Akash Shankaran
- Akshat Jaimini
- Alaa Attya
- Aleksander Alekseev
- Aleksej Orlov
- Alena Rybakina
- Alex Hsieh
- Alex Malek
- Alex Shulgin
- Alex Work
- Alexander Korotkov
- Alexander Kozhemyakin
- Alexander Kuzmenkov
- Alexander Lakhin
- Alexander Pyhalov
- Alexey Palazhchenko
- Alfons Kemper
- Álvaro Herrera
- Amadeo Gallardo
- Amit Kapila
- Amit Langote
- Amul Sul
- Anastasia Lubennikova
- Anatoly Zaretsky
- Andreas Karlsson
- Andreas Ulbrich
- Andrei Lepikhov
- Andrei Zubkov
- Andres Freund
- Andrew Alsup
- Andrew Atkinson
- Andrew Bille
- Andrew Dunstan
- Andrew Kane
- Andrey Borodin
- Andrey Rachitskiy
- Andrey Sokolov
- Andy Fan
- Anthonin Bonnefoy
- Anthony Hsu
- Anton Kirilov
- Anton Melnikov
- Anton Voloshin
- Antonin Houska
- Ants Aasma
- Antti Lampinen
- Aramaki Zyake
- Artem Anisimov
- Artur Zakirov
- Ashutosh Bapat
- Ashutosh Sharma
- Atsushi Torikoshi
- Attila Gulyás
- Ayush Tiwari
- Ayush Vatsa
- Bartosz Chrol
- Benoît Ryder
- Bernd Helmle
- Bertrand Drouvot
- Bharath Rupireddy
- Bo Andreson
- Boshomi Phenix
- Bowen Shi
- Boyu Yang
- Bruce Momjian
- Cameron Vogt
- Cary Huang
- Cédric Villemain
- Changhong Fei
- Chantal Keller
- Chapman Flack
- Chengxi Sun
- Chris Travers
- Christian Maurer
- Christian Stork
- Christoph Berg
- Christoph Heiss
- Christophe Courtois
- Christopher Kline
- Claudio Freire
- Colin Caine
- Corey Huinker
- Curt Kolovson
- Dag Lem
- Dagfinn Ilmari Mannsåker
- Damir Belyalov
- Daniel Fredouille
- Daniel Gustafsson
- Daniel Shelepanov
- Daniel Vérité
- Daniel Westermann
- Darren Rush
- Dave Cramer
- Dave Page
- David Christensen
- David Cook
- David G. Johnston
- David Geier
- David Hillman
- David Perez
- David Rowley
- David Steele
- David Wheeler
- David Zhang
- Dean Rasheed
- Denis Erokhin
- Denis Laxalde
- Devrim Gündüz
- Dilip Kumar
- Dimitrios Apostolou
- Dmitry Dolgov
- Dmitry Koval
- Dmitry Vasiliev
- Dominique Devienne
- Dong Wook Lee
- Donghang Lin
- Dongming Liu
- Drew Callahan
- Drew Kimball
- Dzmitry Jachnik
- Egor Chindyaskin
- Egor Rogov
- Ekaterina Kiryanova
- Elena Indrupskaya
- Elizabeth Christensen
- Emre Hasegeli
- Eric Cyr
- Eric Mutta
- Eric Radman
- Eric Ridge
- Erik Rijkers
- Erik Wienhold
- Erki Eessaar
- Ethan Mertz
- Etsuro Fujita
- Eugen Konkov
- Euler Taveira
- Evan Macbeth
- Evgeny Morozov
- Fabien Coelho
- Fabrízio de Royes Mello
- Farias de Oliveira
- Feliphe Pozzer
- Fire Emerald
- Flavien Guedez
- Floris Van Nee
- Francesco Degrassi
- Frank Streitzig
- Gabriele Bartolini
- Garrett Thornburg
- Gavin Flower
- Gavin Panella
- Gilles Darold
- Gilles Parc
- Grant Gryczan
- Greg Nancarrow
- Greg Sabino Mullane
- Greg Stark
- Gurjeet Singh
- Haiying Tang
- Hajime Matsunaga
- Hal Takahara
- Hanefi Onaldi
- Hannu Krosing
- Hans Buschmann
- Hao Wu
- Hao Zhang
- Hayato Kuroda
- Heikki Linnakangas
- Hemanth Sandrana
- Himanshu Upadhyaya
- Hironobu Suzuki
- Holger Reise
- Hongxu Ma
- Hongyu Song
- Horst Reiterer
- Hubert Lubaczewski
- Hywel Carver
- Ian Barwick
- Ian Ilyasov
- Ilya Nenashev
- Isaac Morland
- Israel Barth Rubio
- Ivan Kartyshov
- Ivan Kolombet
- Ivan Lazarev
- Ivan Panchenko
- Ivan Trofimov
- Jacob Champion
- Jacob Speidel
- Jacques Combrink
- Jaime Casanova
- Jakub Wartak
- James Coleman
- James Pang
- Jani Rahkola
- Japin Li
- Jeevan Chalke
- Jeff Davis
- Jeff Janes
- Jelte Fennema-Nio
- Jeremy Schneider
- Jian Guo
- Jian He
- Jim Jones
- Jim Keener
- Jim Nasby
- Jingtang Zhang
- Jingxian Li
- Jingzhou Fu
- Joe Conway
- Joel Jacobson
- John Ekins
- John Hsu
- John Morris
- John Naylor
- John Russell
- Jonathan Katz
- Jordi Gutiérrez
- Joseph Koshakow
- Josh Kupershmidt
- Joshua D. Drake
- Joshua Uyehara
- Jubilee Young
- Julien Rouhaud
- Junwang Zhao
- Justin Pryzby
- Kaido Vaikla
- Kambam Vinay
- Karen Talarico
- Karina Litskevich
- Karl O. Pinc
- Kashif Zeeshan
- Kim Johan Andersson
- Kirill Reshke
- Kirk Parker
- Kirk Wolak
- Kisoon Kwon
- Koen De Groote
- Kohei KaiGai
- Kong Man
- Konstantin Knizhnik
- Kouhei Sutou
- Krishnakumar R
- Kuntal Ghosh
- Kurt Roeckx
- Kyotaro Horiguchi
- Lang Liu
- Lars Kanis
- Laurenz Albe
- Lauri Laanmets
- Legs Mansion
- Lukas Fittl
- Magnus Hagander
- Mahendrakar Srinivasarao
- Maiquel Grassi
- Manos Emmanouilidis
- Marcel Hofstetter
- Marcos Pegoraro
- Marian Krucina
- Marina Polyakova
- Mark Dilger
- Mark Guertin
- Mark Sloan
- Markus Winand
- Marlene Reiterer
- Martín Marqués
- Martin Nash
- Martin Schlossarek
- Masahiko Sawada
- Masahiro Ikeda
- Masaki Kuwamura
- Masao Fujii
- Mason Sharp
- Matheus Alcantara
- Mats Kindahl
- Matthias Kuhn
- Matthias van de Meent
- Maxim Boguk
- Maxim Orlov
- Maxim Yablokov
- Maxime Boyer
- Melanie Plageman
- Melih Mutlu
- Merlin Moncure
- Micah Gate
- Michael Banck
- Michael Bondarenko
- Michael Paquier
- Michael Wang
- Michael Zhilin
- Michail Nikolaev
- Michal Bartak
- Michal Kleczek
- Mikhail Gribkov
- Mingli Zhang
- Miroslav Bendik
- Mitsuru Hinata
- Moaaz Assali
- Muralikrishna Bandaru
- Nathan Bossart
- Nazir Bilal Yavuz
- Neil Tiffin
- Ngigi Waithaka
- Nikhil Benesch
- Nikhil Raj
- Nikita Glukhov
- Nikita Kalinin
- Nikita Malakhov
- Nikolay Samokhvalov
- Nikolay Shaplov
- Nisha Moond
- Nishant Sharma
- Nitin Jadhav
- Noah Misch
- Noriyoshi Shinoda
- Ole Peder Brandtzæg
- Oleg Bartunov
- Oleg Sibiryakov
- Oleg Tselebrovskiy
- Olleg Samoylov
- Onder Kalaci
- Ondrej Navratil
- Pablo Kharo
- Palak Chaturvedi
- Pantelis Theodosiou
- Paul Amonson
- Paul Jungwirth
- Pavel Borisov
- Pavel Kulakov
- Pavel Luzanov
- Pavel Stehule
- Pavlo Golub
- Pedro Gallegos
- Pete Storer
- Peter Eisentraut
- Peter Geoghegan
- Peter Smith
- Philip Warner
- Philipp Salvisberg
- Pierre Ducroquet
- Pierre Fortin
- Przemyslaw Sztoch
- Quynh Tran
- Raghuveer Devulapalli
- Ranier Vilela
- Reid Thompson
- Rian McGuire
- Richard Guo
- Richard Vesely
- Ridvan Korkmaz
- Robert Haas
- Robert Scott
- Robert Treat
- Roberto Mello
- Robins Tharakan
- Roman Lozko
- Ronan Dunklau
- Rui Zhao
- Ryo Matsumura
- Ryoga Yoshida
- Sameer Kumar
- Sami Imseih
- Samuel Dussault
- Sanjay Minni
- Satoru Koizumi
- Sebastian Skalacki
- Sergei Glukhov
- Sergei Kornilov
- Sergey Prokhorenko
- Sergey Sargsyan
- Sergey Shinderuk
- Shaozhong Shi
- Shaun Thomas
- Shay Rojansky
- Shihao Zhong
- Shinya Kato
- Shlok Kyal
- Shruthi Gowda
- Shubham Khanna
- Shulin Zhou
- Shveta Malik
- Simon Riggs
- Soumyadeep Chakraborty
- Sravan Velagandula
- Stan Hu
- Stepan Neretin
- Stepan Rutz
- Stéphane Schildknecht
- Stephane Tachoires
- Stephen Frost
- Steve Atkins
- Steve Chavez
- Suraj Khamkar
- Suraj Kharage
- Svante Richter
- Svetlana Derevyanko
- Sylvain Frandaz
- Takayuki Tsunakawa
- Tatsuo Ishii
- Tatsuro Yamada
- Tender Wang
- Teodor Sigaev
- Thom Brown
- Thomas Munro
- Tim Carey-Smith
- Tim Needham
- Tim Palmer
- Tobias Bussmann
- Tom Lane
- Tomas Vondra
- Tommy Pavlicek
- Tomonari Katsumata
- Tristan Partin
- Tristen Raab
- Tung Nguyen
- Umair Shahid
- Uwe Binder
- Valerie Woolard
- Vallimaharajan G
- Vasya Boytsov
- Victor Wagner
- Victor Yegorov
- Victoria Shepard
- Vidushi Gupta
- Vignesh C
- Vik Fearing
- Viktor Leis
- Vinayak Pokale
- Vitaly Burovoy
- Vojtech Benes
- Wei Sun
- Wei Wang
- Wenjiang Zhang
- Will Mortensen
- Willi Mann
- Wolfgang Walther
- Xiang Liu
- Xiaoran Wang
- Xing Guo
- Xudong Yang
- Yahor Yuzefovich
- Yajun Hu
- Yaroslav Saburov
- Yong Li
- Yongtao Huang
- Yugo Nagata
- Yuhang Qiu
- Yuki Seino
- Yura Sokolov
- Yurii Rashkovskii
- Yuuki Fujii
- Yuya Watari
- Yves Colin
- Zhihong Yu
- Zhijie Hou
- Zongliang Quan
- Zubeyr Eryilmaz
- Zuming Jiang
