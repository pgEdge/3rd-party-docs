<a id="release-18"></a>

## Release 18


**Release date:.**


2025-09-25
  <a id="release-18-highlights"></a>

### Overview


 PostgreSQL 18 contains many new features and enhancements, including:


-  An asynchronous I/O (AIO) subsystem that can improve performance of sequential scans, bitmap heap scans, vacuums, and other operations.
-  [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) now retains optimizer statistics.
-  Support for "skip scan" lookups that allow using [multicolumn B-tree indexes](../../the-sql-language/indexes/multicolumn-indexes.md#indexes-multicolumn) in more cases.
-  [`uuidv7()`](../../the-sql-language/functions-and-operators/uuid-functions.md#func_uuid_gen_table) function for generating timestamp-ordered [UUIDs](../../the-sql-language/data-types/uuid-type.md#datatype-uuid).
-  Virtual [generated columns](../../reference/sql-commands/create-table.md#sql-createtable-parms-generated-stored) that compute their values during read operations. This is now the default for generated columns.
-  [OAuth authentication](../../server-administration/client-authentication/oauth-authorization-authentication.md#auth-oauth) support.
-  `OLD` and `NEW` support for [`RETURNING`](../../the-sql-language/data-manipulation/returning-data-from-modified-rows.md#dml-returning) clauses in [sql-insert](../../reference/sql-commands/insert.md#sql-insert), [sql-update](../../reference/sql-commands/update.md#sql-update), [sql-delete](../../reference/sql-commands/delete.md#sql-delete), and [sql-merge](../../reference/sql-commands/merge.md#sql-merge) commands.
-  Temporal constraints, or constraints over ranges, for [PRIMARY KEY](../../reference/sql-commands/create-table.md#sql-createtable-parms-primary-key), [UNIQUE](../../reference/sql-commands/create-table.md#sql-createtable-parms-unique), and [FOREIGN KEY](../../reference/sql-commands/create-table.md#sql-createtable-parms-references) constraints.


 The above items and other new features of PostgreSQL 18 are explained in more detail in the sections below.
  <a id="release-18-migration"></a>

### Migration to Version 18


 A dump/restore using [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall) or use of [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) or logical replication is required for those wishing to migrate data from any previous release. See [Upgrading a PostgreSQL Cluster](../../server-administration/server-setup-and-operation/upgrading-a-postgresql-cluster.md#upgrading) for general information on migrating to new major releases.


 Version 18 contains a number of changes that may affect compatibility with previous releases. Observe the following incompatibilities:


-  Change [app-initdb](../../reference/postgresql-server-applications/initdb.md#app-initdb) default to enable data checksums (Greg Sabino Mullane) [&sect;](https://postgr.es/c/04bec894a04)

   Checksums can be disabled with the new initdb option `--no-data-checksums`. [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) requires matching cluster checksum settings, so this new option can be useful to upgrade non-checksum old clusters.
-  Change time zone abbreviation handling (Tom Lane) [&sect;](https://postgr.es/c/d7674c9fa)

   The system will now favor the current session's time zone abbreviations before checking the server variable [timezone_abbreviations](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone-abbreviations). Previously `timezone_abbreviations` was checked first.
-  Deprecate [MD5 password](../../server-administration/client-authentication/password-authentication.md#auth-password) authentication (Nathan Bossart) [&sect;](https://postgr.es/c/db6a4a985)

   Support for MD5 passwords will be removed in a future major version release. [sql-createrole](../../reference/sql-commands/create-role.md#sql-createrole) and [sql-alterrole](../../reference/sql-commands/alter-role.md#sql-alterrole) now emit deprecation warnings when setting MD5 passwords. These warnings can be disabled by setting the [md5_password_warnings](../../server-administration/server-configuration/connections-and-authentication.md#guc-md5-password-warnings) parameter to `off`.
-  Change [sql-vacuum](../../reference/sql-commands/vacuum.md#sql-vacuum) and [sql-analyze](../../reference/sql-commands/analyze.md#sql-analyze) to process the inheritance children of a parent (Michael Harris) [&sect;](https://postgr.es/c/62ddf7ee9)

   The previous behavior can be performed by using the new `ONLY` option.
-  Prevent [`COPY FROM`](../../reference/sql-commands/copy.md#sql-copy) from treating `\.` as an end-of-file marker when reading CSV files (Daniel Vérité, Tom Lane) [&sect;](https://postgr.es/c/770233748) [&sect;](https://postgr.es/c/da8a4c166)

   [app-psql](../../reference/postgresql-client-applications/psql.md#app-psql) will still treat `\.` as an end-of-file marker when reading CSV files from `STDIN`. Older psql clients connecting to PostgreSQL 18 servers might experience [`\copy`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-commands-copy) problems. This release also enforces that `\.` must appear alone on a line.
-  Disallow unlogged partitioned tables (Michael Paquier) [&sect;](https://postgr.es/c/e2bab2d79)

   Previously [`ALTER TABLE SET [UN]LOGGED`](../../reference/sql-commands/alter-table.md#sql-altertable) did nothing, and the creation of an unlogged partitioned table did not cause its children to be unlogged.
-  Execute `AFTER` [triggers](../../server-programming/triggers/index.md#triggers) as the role that was active when trigger events were queued (Laurenz Albe) [&sect;](https://postgr.es/c/01463e1cc)

   Previously such triggers were run as the role that was active at trigger execution time (e.g., at [sql-commit](../../reference/sql-commands/commit.md#sql-commit)). This is significant for cases where the role is changed between queue time and transaction commit.
-  Remove non-functional support for rule privileges in [sql-grant](../../reference/sql-commands/grant.md#sql-grant)/[sql-revoke](../../reference/sql-commands/revoke.md#sql-revoke) (Fujii Masao) [&sect;](https://postgr.es/c/fefa76f70)

   These have been non-functional since PostgreSQL 8.2.
-  Remove column [`pg_backend_memory_contexts`](../../internals/system-views/pg_backend_memory_contexts.md#view-pg-backend-memory-contexts).`parent` (Melih Mutlu) [&sect;](https://postgr.es/c/f0d112759)

   This is no longer needed since `pg_backend_memory_contexts`.`path` was added.
-  Change `pg_backend_memory_contexts`.`level` and [`pg_log_backend_memory_contexts()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-signal-table) to be one-based (Melih Mutlu, Atsushi Torikoshi, David Rowley, Fujii Masao) [&sect;](https://postgr.es/c/32d3ed816) [&sect;](https://postgr.es/c/d9e03864b) [&sect;](https://postgr.es/c/706cbed35)

   These were previously zero-based.
-  Change [full text search](../../the-sql-language/full-text-search/index.md#textsearch) to use the default collation provider of the cluster to read configuration files and dictionaries, rather than always using libc (Peter Eisentraut) [&sect;](https://postgr.es/c/fb1a18810f0)

   Clusters that default to non-libc collation providers (e.g., ICU, builtin) that behave differently than libc for characters processed by LC_CTYPE could observe changes in behavior of some full-text search functions, as well as the [pg_trgm](../additional-supplied-modules-and-extensions/pg_trgm-support-for-similarity-of-text-using-trigram-matching.md#pgtrgm) extension. When upgrading such clusters using [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade), it is recommended to reindex all indexes related to full-text search and pg_trgm after the upgrade.
  <a id="release-18-changes"></a>

### Changes


 Below you will find a detailed account of the changes between PostgreSQL 18 and the previous major release.
 <a id="release-18-server"></a>

#### Server
  <a id="release-18-optimizer"></a>

##### Optimizer


-  Automatically remove some unnecessary table self-joins (Andrey Lepikhov, Alexander Kuzmenkov, Alexander Korotkov, Alena Rybakina) [&sect;](https://postgr.es/c/fc069a3a6)

   This optimization can be disabled using server variable [enable_self_join_elimination](../../server-administration/server-configuration/query-planning.md#guc-enable-self-join-elimination).
-  Convert some [`IN (VALUES ...)`](../../the-sql-language/functions-and-operators/row-and-array-comparisons.md#functions-comparisons-in-scalar) to `x = ANY ...` for better optimizer statistics (Alena Rybakina, Andrei Lepikhov) [&sect;](https://postgr.es/c/c0962a113)
-  Allow transforming [`OR`](../../the-sql-language/functions-and-operators/logical-operators.md#functions-logical)-clauses to arrays for faster index processing (Alexander Korotkov, Andrey Lepikhov) [&sect;](https://postgr.es/c/ae4569161)
-  Speed up the processing of [`INTERSECT`](../../reference/sql-commands/select.md#sql-intersect), [`EXCEPT`](../../reference/sql-commands/select.md#sql-except), [window aggregates](../../tutorial/advanced-features/window-functions.md#tutorial-window), and [view column aliases](../../reference/sql-commands/create-view.md#sql-createview) (Tom Lane, David Rowley) [&sect;](https://postgr.es/c/52c707483) [&sect;](https://postgr.es/c/276279295) [&sect;](https://postgr.es/c/8d96f57d5) [&sect;](https://postgr.es/c/908a96861)
-  Allow the keys of [`SELECT DISTINCT`](../../reference/sql-commands/select.md#sql-distinct) to be internally reordered to avoid sorting (Richard Guo) [&sect;](https://postgr.es/c/a8ccf4e93)

   This optimization can be disabled using [enable_distinct_reordering](../../server-administration/server-configuration/query-planning.md#guc-enable-distinct-reordering).
-  Ignore [`GROUP BY`](../../reference/sql-commands/select.md#sql-groupby) columns that are functionally dependent on other columns (Zhang Mingli, Jian He, David Rowley) [&sect;](https://postgr.es/c/bd10ec529)

   If a `GROUP BY` clause includes all columns of a unique index, as well as other columns of the same table, those other columns are redundant and can be dropped from the grouping. This was already true for non-deferred primary keys.
-  Allow some [`HAVING`](../../reference/sql-commands/select.md#sql-having) clauses on [`GROUPING SETS`](../../the-sql-language/queries/table-expressions.md#queries-grouping-sets) to be pushed to [`WHERE`](../../reference/sql-commands/select.md#sql-where) clauses (Richard Guo) [&sect;](https://postgr.es/c/67a54b9e8) [&sect;](https://postgr.es/c/247dea89f) [&sect;](https://postgr.es/c/f5050f795) [&sect;](https://postgr.es/c/cc5d98525)

   This allows earlier row filtering. This release also fixes some `GROUPING SETS` queries that used to return incorrect results.
-  Improve row estimates for [`generate_series()`](../../the-sql-language/functions-and-operators/set-returning-functions.md#functions-srf-series) using [`numeric`](../../the-sql-language/data-types/numeric-types.md#datatype-numeric) and [`timestamp`](../../the-sql-language/data-types/date-time-types.md#datatype-datetime) values (David Rowley, Song Jinzhou) [&sect;](https://postgr.es/c/036bdcec9) [&sect;](https://postgr.es/c/97173536e)
-  Allow the optimizer to use `Right Semi Join` plans (Richard Guo) [&sect;](https://postgr.es/c/aa86129e1)

   Semi-joins are used when needing to find if there is at least one match.
-  Allow merge joins to use [incremental sorts](../../server-administration/server-configuration/query-planning.md#guc-enable-incremental-sort) (Richard Guo) [&sect;](https://postgr.es/c/828e94c9d)
-  Improve the efficiency of planning queries accessing many partitions (Ashutosh Bapat, Yuya Watari, David Rowley) [&sect;](https://postgr.es/c/88f55bc97) [&sect;](https://postgr.es/c/d69d45a5a)
-  Allow [partitionwise joins](../../server-administration/server-configuration/query-planning.md#guc-enable-partitionwise-join) in more cases, and reduce its memory usage (Richard Guo, Tom Lane, Ashutosh Bapat) [&sect;](https://postgr.es/c/9b282a935) [&sect;](https://postgr.es/c/513f4472a)
-  Improve cost estimates of partition queries (Nikita Malakhov, Andrei Lepikhov) [&sect;](https://postgr.es/c/fae535da0)
-  Improve [SQL-language function](../../server-programming/extending-sql/query-language-sql-functions.md#xfunc-sql) plan caching (Alexander Pyhalov, Tom Lane) [&sect;](https://postgr.es/c/0dca5d68d) [&sect;](https://postgr.es/c/09b07c295)
-  Improve handling of disabled optimizer features (Robert Haas) [&sect;](https://postgr.es/c/e22253467)
  <a id="release-18-indexes"></a>

##### Indexes


-  Allow skip scans of [btree](../../server-programming/extending-sql/query-language-sql-functions.md#xfunc-sql) indexes (Peter Geoghegan) [&sect;](https://postgr.es/c/92fe23d93) [&sect;](https://postgr.es/c/8a510275d)

   This allows multi-column btree indexes to be used in more cases such as when there are no restrictions on the first or early indexed columns (or there are non-equality ones), and there are useful restrictions on later indexed columns.
-  Allow non-btree unique indexes to be used as partition keys and in materialized views (Mark Dilger) [&sect;](https://postgr.es/c/f278e1fe3) [&sect;](https://postgr.es/c/9d6db8bec)

   The index type must still support equality.
-  Allow [`GIN`](../../internals/built-in-index-access-methods/gin-indexes.md#gin) indexes to be created in parallel (Tomas Vondra, Matthias van de Meent) [&sect;](https://postgr.es/c/8492feb98)
-  Allow values to be sorted to speed range-type [GiST](../../internals/built-in-index-access-methods/gist-indexes.md#gist) and [btree](../../internals/built-in-index-access-methods/b-tree-indexes.md#btree) index builds (Bernd Helmle) [&sect;](https://postgr.es/c/e9e7b6604)
  <a id="release-18-performance"></a>

##### General Performance


-  Add an asynchronous I/O subsystem (Andres Freund, Thomas Munro, Nazir Bilal Yavuz, Melanie Plageman) [&sect;](https://postgr.es/c/02844012b) [&sect;](https://postgr.es/c/da7226993) [&sect;](https://postgr.es/c/55b454d0e) [&sect;](https://postgr.es/c/247ce06b8) [&sect;](https://postgr.es/c/10f664684) [&sect;](https://postgr.es/c/06fb5612c) [&sect;](https://postgr.es/c/c325a7633) [&sect;](https://postgr.es/c/50cb7505b) [&sect;](https://postgr.es/c/047cba7fa) [&sect;](https://postgr.es/c/12ce89fd0) [&sect;](https://postgr.es/c/2a5e709e7)

   This feature allows backends to queue multiple read requests, which allows for more efficient sequential scans, bitmap heap scans, vacuums, etc. This is enabled by server variable [io_method](../../server-administration/server-configuration/resource-consumption.md#guc-io-method), with server variables [io_combine_limit](../../server-administration/server-configuration/resource-consumption.md#guc-io-combine-limit) and [io_max_combine_limit](../../server-administration/server-configuration/resource-consumption.md#guc-io-max-combine-limit) added to control it. This also enables [effective_io_concurrency](../../server-administration/server-configuration/resource-consumption.md#guc-effective-io-concurrency) and [maintenance_io_concurrency](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-io-concurrency) values greater than zero for systems without `fadvise()` support. The new system view [`pg_aios`](../../internals/system-views/pg_aios.md#view-pg-aios) shows the file handles being used for asynchronous I/O.
-  Improve the locking performance of queries that access many relations (Tomas Vondra) [&sect;](https://postgr.es/c/c4d5cb71d)
-  Improve the performance and reduce memory usage of hash joins and [`GROUP BY`](../../reference/sql-commands/select.md#sql-groupby) (David Rowley, Jeff Davis) [&sect;](https://postgr.es/c/adf97c156) [&sect;](https://postgr.es/c/0f5738202) [&sect;](https://postgr.es/c/4d143509c) [&sect;](https://postgr.es/c/a0942f441) [&sect;](https://postgr.es/c/626df47ad)

   This also improves hash set operations used by [`EXCEPT`](../../reference/sql-commands/select.md#sql-except), and hash lookups of subplan values.
-  Allow normal vacuums to freeze some pages, even though they are all-visible (Melanie Plageman) [&sect;](https://postgr.es/c/052026c9b) [&sect;](https://postgr.es/c/06eae9e62)

   This reduces the overhead of later full-relation freezing. The aggressiveness of this can be controlled by server variable and per-table setting [vacuum_max_eager_freeze_failure_rate](../../server-administration/server-configuration/vacuuming.md#guc-vacuum-max-eager-freeze-failure-rate). Previously vacuum never processed all-visible pages until freezing was required.
-  Add server variable [vacuum_truncate](../../server-administration/server-configuration/vacuuming.md#guc-vacuum-truncate) to control file truncation during [sql-vacuum](../../reference/sql-commands/vacuum.md#sql-vacuum) (Nathan Bossart, Gurjeet Singh) [&sect;](https://postgr.es/c/0164a0f9e)

   A storage-level parameter with the same name and behavior already existed.
-  Increase server variables [effective_io_concurrency](../../server-administration/server-configuration/resource-consumption.md#guc-effective-io-concurrency)'s and [maintenance_io_concurrency](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-io-concurrency)'s default values to 16 (Melanie Plageman) [&sect;](https://postgr.es/c/ff79b5b2a) [&sect;](https://postgr.es/c/cc6be07eb)

   This more accurately reflects modern hardware.
  <a id="release-18-monitoring"></a>

##### Monitoring


-  Increase the logging granularity of server variable [log_connections](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-connections) (Melanie Plageman) [&sect;](https://postgr.es/c/9219093ca)

   This server variable was previously only boolean, which is still supported.
-  Add `log_connections` option to report the duration of connection stages (Melanie Plageman) [&sect;](https://postgr.es/c/18cd15e70)
-  Add [log_line_prefix](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-line-prefix) escape `%L` to output the client IP address (Greg Sabino Mullane) [&sect;](https://postgr.es/c/3516ea768)
-  Add server variable [log_lock_failures](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-lock-failures) to log lock acquisition failures (Yuki Seino, Fujii Masao) [&sect;](https://postgr.es/c/6d376c3b0) [&sect;](https://postgr.es/c/73bdcfab3)

   Specifically it reports [`SELECT ... NOWAIT`](../../reference/sql-commands/select.md#sql-for-update-share) lock failures.
-  Modify [`pg_stat_all_tables`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-all-tables-view) and its variants to report the time spent in [sql-vacuum](../../reference/sql-commands/vacuum.md#sql-vacuum), [sql-analyze](../../reference/sql-commands/analyze.md#sql-analyze), and their [automatic](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) variants (Sami Imseih) [&sect;](https://postgr.es/c/30a6ed0ce)

   The new columns are `total_vacuum_time`, `total_autovacuum_time`, `total_analyze_time`, and `total_autoanalyze_time`.
-  Add delay time reporting to [sql-vacuum](../../reference/sql-commands/vacuum.md#sql-vacuum) and [sql-analyze](../../reference/sql-commands/analyze.md#sql-analyze) (Bertrand Drouvot, Nathan Bossart) [&sect;](https://postgr.es/c/bb8dff999) [&sect;](https://postgr.es/c/7720082ae)

   This information appears in the server log, the system views [`pg_stat_progress_vacuum`](../../server-administration/monitoring-database-activity/progress-reporting.md#vacuum-progress-reporting) and [`pg_stat_progress_analyze`](../../server-administration/monitoring-database-activity/progress-reporting.md#pg-stat-progress-analyze-view), and the output of [sql-vacuum](../../reference/sql-commands/vacuum.md#sql-vacuum) and [sql-analyze](../../reference/sql-commands/analyze.md#sql-analyze) when in `VERBOSE` mode; tracking must be enabled with the server variable [track_cost_delay_timing](../../server-administration/server-configuration/run-time-statistics.md#guc-track-cost-delay-timing).
-  Add WAL, CPU, and average read statistics output to `ANALYZE VERBOSE` (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/4c1b4cdb8) [&sect;](https://postgr.es/c/bb7775234)
-  Add full WAL buffer count to `VACUUM`/`ANALYZE (VERBOSE)` and autovacuum log output (Bertrand Drouvot) [&sect;](https://postgr.es/c/6a8a7ce47)
-  Add per-backend I/O statistics reporting (Bertrand Drouvot) [&sect;](https://postgr.es/c/9aea73fc6) [&sect;](https://postgr.es/c/3f1db99bf)

   The statistics are accessed via [`pg_stat_get_backend_io()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-get-backend-io). Per-backend I/O statistics can be cleared via [`pg_stat_reset_backend_stats()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-funcs-table).
-  Add [`pg_stat_io`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-io-view) columns to report I/O activity in bytes (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/f92c854cf)

   The new columns are `read_bytes`, `write_bytes`, and `extend_bytes`. The `op_bytes` column, which always equaled [`BLCKSZ`](../../server-administration/server-configuration/preset-options.md#guc-block-size), has been removed.
-  Add WAL I/O activity rows to `pg_stat_io` (Nazir Bilal Yavuz, Bertrand Drouvot, Michael Paquier) [&sect;](https://postgr.es/c/a051e71e2) [&sect;](https://postgr.es/c/4538bd3f1) [&sect;](https://postgr.es/c/7f7f324eb)

   This includes WAL receiver activity and a wait event for such writes.
-  Change server variable [track_wal_io_timing](../../server-administration/server-configuration/run-time-statistics.md#guc-track-wal-io-timing) to control tracking WAL timing in `pg_stat_io` instead of [`pg_stat_wal`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-wal-view) (Bertrand Drouvot) [&sect;](https://postgr.es/c/6c349d83b)
-  Remove read/sync columns from `pg_stat_wal` (Bertrand Drouvot) [&sect;](https://postgr.es/c/2421e9a51) [&sect;](https://postgr.es/c/6c349d83b)

   This removes columns `wal_write`, `wal_sync`, `wal_write_time`, and `wal_sync_time`.
-  Add function [`pg_stat_get_backend_wal()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-get-backend-wal) to return per-backend WAL statistics (Bertrand Drouvot) [&sect;](https://postgr.es/c/76def4cdd)

   Per-backend WAL statistics can be cleared via [`pg_stat_reset_backend_stats()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-funcs-table).
-  Add function [`pg_ls_summariesdir()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-genfile-table) to specifically list the contents of [`PGDATA`](../../internals/database-physical-storage/database-file-layout.md#storage-file-layout)/[`pg_wal/summaries`](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-summary-keep-time) (Yushi Ogiwara) [&sect;](https://postgr.es/c/4e1fad378)
-  Add column [`pg_stat_checkpointer`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-checkpointer-view).`num_done` to report the number of completed checkpoints (Anton A. Melnikov) [&sect;](https://postgr.es/c/559efce1d)

   Columns `num_timed` and `num_requested` count both completed and skipped checkpoints.
-  Add column `pg_stat_checkpointer`.`slru_written` to report SLRU buffers written (Nitin Jadhav) [&sect;](https://postgr.es/c/17cc5f666)

   Also, modify the checkpoint server log message to report separate shared buffer and SLRU buffer values.
-  Add columns to [`pg_stat_database`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-database-view) to report parallel worker activity (Benoit Lobréau) [&sect;](https://postgr.es/c/e7a9496de)

   The new columns are `parallel_workers_to_launch` and `parallel_workers_launched`.
-  Have [query id](../../server-administration/server-configuration/run-time-statistics.md#guc-compute-query-id) computation of constant lists consider only the first and last constants (Dmitry Dolgov, Sami Imseih) [&sect;](https://postgr.es/c/62d712ecf) [&sect;](https://postgr.es/c/9fbd53dea) [&sect;](https://postgr.es/c/c2da1a5d6)

   Jumbling is used by [pg_stat_statements](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements).
-  Adjust query id computations to group together queries using the same relation name (Michael Paquier, Sami Imseih) [&sect;](https://postgr.es/c/787514b30)

   This is true even if the tables in different schemas have different column names.
-  Add column [`pg_backend_memory_contexts`](../../internals/system-views/pg_backend_memory_contexts.md#view-pg-backend-memory-contexts).`type` to report the type of memory context (David Rowley) [&sect;](https://postgr.es/c/12227a1d5)
-  Add column `pg_backend_memory_contexts`.`path` to show memory context parents (Melih Mutlu) [&sect;](https://postgr.es/c/32d3ed816)
  <a id="release-18-privileges"></a>

##### Privileges


-  Add function [`pg_get_acl()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-object-table) to retrieve database access control details (Joel Jacobson) [&sect;](https://postgr.es/c/4564f1ceb) [&sect;](https://postgr.es/c/d898665bf)
-  Add function [`has_largeobject_privilege()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-access-table) to check large object privileges (Yugo Nagata) [&sect;](https://postgr.es/c/4eada203a)
-  Allow [sql-alterdefaultprivileges](../../reference/sql-commands/alter-default-privileges.md#sql-alterdefaultprivileges) to define large object default privileges (Takatsuka Haruka, Yugo Nagata, Laurenz Albe) [&sect;](https://postgr.es/c/0d6c47766)
-  Add predefined role [`pg_signal_autovacuum_worker`](../../server-administration/database-roles/predefined-roles.md#predefined-roles) (Kirill Reshke) [&sect;](https://postgr.es/c/ccd38024b)

   This allows sending signals to autovacuum workers.
  <a id="release-18-server-config"></a>

##### Server Configuration


-  Add support for the [OAuth authentication method](../../server-administration/client-authentication/oauth-authorization-authentication.md#auth-oauth) (Jacob Champion, Daniel Gustafsson, Thomas Munro) [&sect;](https://postgr.es/c/b3f0be788)

   This adds an `oauth` authentication method to [`pg_hba.conf`](../../server-administration/client-authentication/the-pg_hba-conf-file.md#auth-pg-hba-conf), libpq OAuth options, a server variable [oauth_validator_libraries](../../server-administration/server-configuration/connections-and-authentication.md#guc-oauth-validator-libraries) to load token validation libraries, and a configure flag [`--with-libcurl`](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-option-with-libcurl) to add the required compile-time libraries.
-  Add server variable [ssl_tls13_ciphers](../../server-administration/server-configuration/connections-and-authentication.md#guc-ssl-tls13-ciphers) to allow specification of multiple colon-separated TLSv1.3 cipher suites (Erica Zhang, Daniel Gustafsson) [&sect;](https://postgr.es/c/45188c2ea)
-  Change server variable [ssl_groups](../../server-administration/server-configuration/connections-and-authentication.md#guc-ssl-groups)'s default to include elliptic curve X25519 (Daniel Gustafsson, Jacob Champion) [&sect;](https://postgr.es/c/daa02c6bd)
-  Rename server variable `ssl_ecdh_curve` to [ssl_groups](../../server-administration/server-configuration/connections-and-authentication.md#guc-ssl-groups) and allow multiple colon-separated ECDH curves to be specified (Erica Zhang, Daniel Gustafsson) [&sect;](https://postgr.es/c/3d1ef3a15)

   The previous name still works.
-  Make [cancel request keys](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-signal-table) 256 bits (Heikki Linnakangas, Jelte Fennema-Nio) [&sect;](https://postgr.es/c/a460251f0) [&sect;](https://postgr.es/c/9d9b9d46f)

   This is only possible when the server and client support wire protocol version 3.2, introduced in this release.
-  Add server variable [autovacuum_worker_slots](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-worker-slots) to specify the maximum number of background workers (Nathan Bossart) [&sect;](https://postgr.es/c/c758119e5)

   With this variable set, [autovacuum_max_workers](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-max-workers) can be adjusted at runtime up to this maximum without a server restart.
-  Allow specification of the fixed number of dead tuples that will trigger an [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) (Nathan Bossart, Frédéric Yhuel) [&sect;](https://postgr.es/c/306dc520b)

   The server variable is [autovacuum_vacuum_max_threshold](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-vacuum-max-threshold). Percentages are still used for triggering.
-  Change server variable [max_files_per_process](../../server-administration/server-configuration/resource-consumption.md#guc-max-files-per-process) to limit only files opened by a backend (Andres Freund) [&sect;](https://postgr.es/c/adb5f85fa)

   Previously files opened by the postmaster were also counted toward this limit.
-  Add server variable [num_os_semaphores](../../server-administration/server-configuration/preset-options.md#guc-num-os-semaphores) to report the required number of semaphores (Nathan Bossart) [&sect;](https://postgr.es/c/0dcaea569)

   This is useful for operating system configuration.
-  Add server variable [extension_control_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-extension-control-path) to specify the location of extension control files (Peter Eisentraut, Matheus Alcantara) [&sect;](https://postgr.es/c/4f7f7b037) [&sect;](https://postgr.es/c/81eaaa2c4)
  <a id="release-18-replication"></a>

##### Streaming Replication and Recovery


-  Allow inactive replication slots to be automatically invalidated using server variable [idle_replication_slot_timeout](../../server-administration/server-configuration/replication.md#guc-idle-replication-slot-timeout) (Nisha Moond, Bharath Rupireddy) [&sect;](https://postgr.es/c/ac0e33136)
-  Add server variable [max_active_replication_origins](../../server-administration/server-configuration/replication.md#guc-max-active-replication-origins) to control the maximum active replication origins (Euler Taveira) [&sect;](https://postgr.es/c/04ff636cb)

   This was previously controlled by [max_replication_slots](../../server-administration/server-configuration/replication.md#guc-max-replication-slots), but this new setting allows a higher origin count in cases where fewer slots are required.
  <a id="release-18-logical"></a>

##### [Logical Replication]


-  Allow the values of [generated columns](../../reference/sql-commands/create-table.md#sql-createtable-parms-generated-stored) to be logically replicated (Shubham Khanna, Vignesh C, Zhijie Hou, Shlok Kyal, Peter Smith) [&sect;](https://postgr.es/c/745217a05) [&sect;](https://postgr.es/c/7054186c4) [&sect;](https://postgr.es/c/87ce27de6) [&sect;](https://postgr.es/c/6252b1eaf)

   If the publication specifies a column list, all specified columns, generated and non-generated, are published. Without a specified column list, publication option `publish_generated_columns` controls whether generated columns are published. Previously generated columns were not replicated and the subscriber had to compute the values if possible; this is particularly useful for non-PostgreSQL subscribers which lack such a capability.
-  Change the default [sql-createsubscription](../../reference/sql-commands/create-subscription.md#sql-createsubscription) streaming option from `off` to `parallel` (Vignesh C) [&sect;](https://postgr.es/c/1bf1140be)
-  Allow [sql-altersubscription](../../reference/sql-commands/alter-subscription.md#sql-altersubscription) to change the replication slot's two-phase commit behavior (Hayato Kuroda, Ajin Cherian, Amit Kapila, Zhijie Hou) [&sect;](https://postgr.es/c/1462aad2e) [&sect;](https://postgr.es/c/4868c96bc)
-  Log [conflicts](../../server-administration/high-availability-load-balancing-and-replication/hot-standby.md#hot-standby-conflict) while applying logical replication changes (Zhijie Hou, Nisha Moond) [&sect;](https://postgr.es/c/9758174e2) [&sect;](https://postgr.es/c/edcb71258) [&sect;](https://postgr.es/c/640178c92) [&sect;](https://postgr.es/c/6c2b5edec) [&sect;](https://postgr.es/c/73eba5004)

   Also report in new columns of [`pg_stat_subscription_stats`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-subscription-stats).
   <a id="release-18-utility"></a>

#### Utility Commands


-  Allow [generated columns](../../reference/sql-commands/create-table.md#sql-createtable-parms-generated-stored) to be virtual, and make them the default (Peter Eisentraut, Jian He, Richard Guo, Dean Rasheed) [&sect;](https://postgr.es/c/83ea6c540) [&sect;](https://postgr.es/c/cdc168ad4) [&sect;](https://postgr.es/c/1e4351af3)

   Virtual generated columns generate their values when the columns are read, not written. The write behavior can still be specified via the `STORED` option.
-  Add `OLD`/`NEW` support to [`RETURNING`](../../the-sql-language/data-manipulation/returning-data-from-modified-rows.md#dml-returning) in DML queries (Dean Rasheed) [&sect;](https://postgr.es/c/80feb727c)

   Previously `RETURNING` only returned new values for [sql-insert](../../reference/sql-commands/insert.md#sql-insert) and [sql-update](../../reference/sql-commands/update.md#sql-update), and old values for [sql-delete](../../reference/sql-commands/delete.md#sql-delete); [sql-merge](../../reference/sql-commands/merge.md#sql-merge) would return the appropriate value for the internal query executed. This new syntax allows the `RETURNING` list of `INSERT`/`UPDATE`/`DELETE`/`MERGE` to explicitly return old and new values by using the special aliases `old` and `new`. These aliases can be renamed to avoid identifier conflicts.
-  Allow foreign tables to be created like existing local tables (Zhang Mingli) [&sect;](https://postgr.es/c/302cf1575)

   The syntax is [`CREATE FOREIGN TABLE ... LIKE`](../../reference/sql-commands/create-foreign-table.md#sql-createforeigntable).
-  Allow [`LIKE`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-like) with [nondeterministic collations](../../server-administration/localization/collation-support.md#collation-nondeterministic) (Peter Eisentraut) [&sect;](https://postgr.es/c/85b7efa1c)
-  Allow text position search functions with nondeterministic collations (Peter Eisentraut) [&sect;](https://postgr.es/c/329304c90)

   These used to generate an error.
-  Add builtin collation provider [`PG_UNICODE_FAST`](../../server-administration/localization/locale-support.md#locale-providers) (Jeff Davis) [&sect;](https://postgr.es/c/d3d098316)

   This locale supports case mapping, but sorts in code point order, not natural language order.
-  Allow [sql-vacuum](../../reference/sql-commands/vacuum.md#sql-vacuum) and [sql-analyze](../../reference/sql-commands/analyze.md#sql-analyze) to process partitioned tables without processing their children (Michael Harris) [&sect;](https://postgr.es/c/62ddf7ee9)

   This is enabled with the new `ONLY` option. This is useful since autovacuum does not process partitioned tables, just its children.
-  Add functions to modify per-relation and per-column optimizer statistics (Corey Huinker) [&sect;](https://postgr.es/c/e839c8ecc) [&sect;](https://postgr.es/c/d32d14639) [&sect;](https://postgr.es/c/650ab8aaf)

   The functions are [`pg_restore_relation_stats()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-statsmod), `pg_restore_attribute_stats()`, `pg_clear_relation_stats()`, and `pg_clear_attribute_stats()`.
-  Add server variable [file_copy_method](../../server-administration/server-configuration/resource-consumption.md#guc-file-copy-method) to control the file copying method (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/f78ca6f3e)

   This controls whether [`CREATE DATABASE ... STRATEGY=FILE_COPY`](../../reference/sql-commands/create-database.md#sql-createdatabase) and [`ALTER DATABASE ... SET TABLESPACE`](../../reference/sql-commands/alter-database.md#sql-alterdatabase) uses file copy or clone.
 <a id="release-18-constraints"></a>

##### [Constraints]


-  Allow the specification of non-overlapping [`PRIMARY KEY`](../../reference/sql-commands/create-table.md#sql-createtable-parms-primary-key), [`UNIQUE`](../../reference/sql-commands/create-table.md#sql-createtable-parms-unique), and [foreign key](../../reference/sql-commands/create-table.md#sql-createtable-parms-references) constraints (Paul A. Jungwirth) [&sect;](https://postgr.es/c/fc0438b4e) [&sect;](https://postgr.es/c/89f908a6d)

   This is specified by `WITHOUT OVERLAPS` for `PRIMARY KEY` and `UNIQUE`, and by `PERIOD` for foreign keys, all applied to the last specified column.
-  Allow [`CHECK`](../../reference/sql-commands/create-table.md#sql-createtable-parms-check) and [foreign key](../../reference/sql-commands/create-table.md#sql-createtable-parms-references) constraints to be specified as `NOT ENFORCED` (Amul Sul) [&sect;](https://postgr.es/c/ca87c415e) [&sect;](https://postgr.es/c/eec0040c4)

   This also adds column [`pg_constraint`](../../internals/system-catalogs/pg_constraint.md#catalog-pg-constraint).`conenforced`.
-  Require [primary/foreign key](../../reference/sql-commands/create-table.md#sql-createtable-parms-references) relationships to use either deterministic collations or the the same nondeterministic collations (Peter Eisentraut) [&sect;](https://postgr.es/c/9321d2fdf)

   The restore of a [app-pgdump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump), also used by [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade), will fail if these requirements are not met; schema changes must be made for these upgrade methods to succeed.
-  Store column [`NOT NULL`](../../reference/sql-commands/create-table.md#sql-createtable-parms-not-null) specifications in [`pg_constraint`](../../internals/system-catalogs/pg_constraint.md#catalog-pg-constraint) (Álvaro Herrera, Bernd Helmle) [&sect;](https://postgr.es/c/14e87ffa5) [&sect;](https://postgr.es/c/81ce602d4)

   This allows names to be specified for `NOT NULL` constraint. This also adds `NOT NULL` constraints to foreign tables and `NOT NULL` inheritance control to local tables.
-  Allow [sql-altertable](../../reference/sql-commands/alter-table.md#sql-altertable) to set the `NOT VALID` attribute of `NOT NULL` constraints (Rushabh Lathia, Jian He) [&sect;](https://postgr.es/c/a379061a2)
-  Allow modification of the inheritability of `NOT NULL` constraints (Suraj Kharage, Álvaro Herrera) [&sect;](https://postgr.es/c/f4e53e10b) [&sect;](https://postgr.es/c/4a02af8b1)

   The syntax is [`ALTER TABLE ... ALTER CONSTRAINT ... [NO] INHERIT`](../../reference/sql-commands/alter-table.md#sql-altertable).
-  Allow `NOT VALID` foreign key constraints on partitioned tables (Amul Sul) [&sect;](https://postgr.es/c/b663b9436)
-  Allow [dropping](../../reference/sql-commands/alter-table.md#sql-altertable-desc-drop-constraint) of constraints `ONLY` on partitioned tables (Álvaro Herrera) [&sect;](https://postgr.es/c/4dea33ce7)

   This was previously erroneously prohibited.
  <a id="release-18-copy"></a>

##### [sql-copy]


-  Add `REJECT_LIMIT` to control the number of invalid rows `COPY FROM` can ignore (Atsushi Torikoshi) [&sect;](https://postgr.es/c/4ac2a9bec)

   This is available when `ON_ERROR = 'ignore'`.
-  Allow `COPY TO` to copy rows from populated materialized views (Jian He) [&sect;](https://postgr.es/c/534874fac)
-  Add `COPY` `LOG_VERBOSITY` level `silent` to suppress log output of ignored rows (Atsushi Torikoshi) [&sect;](https://postgr.es/c/e7834a1a2)

   This new level suppresses output for discarded input rows when `on_error = 'ignore'`.
-  Disallow `COPY FREEZE` on foreign tables (Nathan Bossart) [&sect;](https://postgr.es/c/401a6956f)

   Previously, the `COPY` worked but the `FREEZE` was ignored, so disallow this command.
  <a id="release-18-explain"></a>

##### [sql-explain]


-  Automatically include `BUFFERS` output in `EXPLAIN ANALYZE` (Guillaume Lelarge, David Rowley) [&sect;](https://postgr.es/c/c2a4078eb)
-  Add full WAL buffer count to `EXPLAIN (WAL)` output (Bertrand Drouvot) [&sect;](https://postgr.es/c/320545bfc)
-  In `EXPLAIN ANALYZE`, report the number of index lookups used per index scan node (Peter Geoghegan) [&sect;](https://postgr.es/c/0fbceae84)
-  Modify `EXPLAIN` to output fractional row counts (Ibrar Ahmed, Ilia Evdokimov, Robert Haas) [&sect;](https://postgr.es/c/ddb17e387) [&sect;](https://postgr.es/c/95dbd827f)
-  Add memory and disk usage details to `Material`, `Window Aggregate`, and common table expression nodes to `EXPLAIN` output (David Rowley, Tatsuo Ishii) [&sect;](https://postgr.es/c/1eff8279d) [&sect;](https://postgr.es/c/53abb1e0e) [&sect;](https://postgr.es/c/95d6e9af0) [&sect;](https://postgr.es/c/40708acd6)
-  Add details about window function arguments to `EXPLAIN` output (Tom Lane) [&sect;](https://postgr.es/c/8b1b34254)
-  Add `Parallel Bitmap Heap Scan` worker cache statistics to `EXPLAIN ANALYZE` (David Geier, Heikki Linnakangas, Donghang Lin, Alena Rybakina, David Rowley) [&sect;](https://postgr.es/c/5a1e6df3b)
-  Indicate disabled nodes in `EXPLAIN ANALYZE` output (Robert Haas, David Rowley, Laurenz Albe) [&sect;](https://postgr.es/c/c01743aa4) [&sect;](https://postgr.es/c/161320b4b) [&sect;](https://postgr.es/c/84b8fccbe)
   <a id="release-18-datatypes"></a>

#### Data Types


-  Improve [Unicode](../../server-administration/localization/collation-support.md#collation-managing-standard) full case mapping and conversion (Jeff Davis) [&sect;](https://postgr.es/c/4e7f62bc3) [&sect;](https://postgr.es/c/286a365b9)

   This adds the ability to do conditional and title case mapping, and case map single characters to multiple characters.
-  Allow [`jsonb`](../../the-sql-language/data-types/json-types.md#datatype-json) `null` values to be cast to scalar types as `NULL` (Tom Lane) [&sect;](https://postgr.es/c/a5579a90a)

   Previously such casts generated an error.
-  Add optional parameter to [`json{b}_strip_nulls`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-json-processing-table) to allow removal of null array elements (Florents Tselai) [&sect;](https://postgr.es/c/4603903d2)
-  Add function [`array_sort()`](../../the-sql-language/functions-and-operators/array-functions-and-operators.md#array-functions-table) which sorts an array's first dimension (Junwang Zhao, Jian He) [&sect;](https://postgr.es/c/6c12ae09f)
-  Add function [`array_reverse()`](../../the-sql-language/functions-and-operators/array-functions-and-operators.md#array-functions-table) which reverses an array's first dimension (Aleksander Alekseev) [&sect;](https://postgr.es/c/49d6c7d8d)
-  Add function [`reverse()`](../../the-sql-language/functions-and-operators/string-functions-and-operators.md#functions-string-other) to reverse bytea bytes (Aleksander Alekseev) [&sect;](https://postgr.es/c/0697b2390)
-  Allow casting between integer types and [`bytea`](../../the-sql-language/data-types/binary-data-types.md#datatype-binary) (Aleksander Alekseev) [&sect;](https://postgr.es/c/6da469bad)

   The integer values are stored as `bytea` two's complement values.
-  Update Unicode data to [Unicode](../../server-administration/localization/collation-support.md#collation-managing-standard) 16.0.0 (Peter Eisentraut) [&sect;](https://postgr.es/c/82a46cca9)
-  Add full text search [stemming](../../the-sql-language/full-text-search/dictionaries.md#textsearch-snowball-dictionary) for Estonian (Tom Lane) [&sect;](https://postgr.es/c/b464e51ab)
-  Improve the [`XML`](../../the-sql-language/data-types/xml-type.md#datatype-xml) error codes to more closely match the SQL standard (Tom Lane) [&sect;](https://postgr.es/c/cd838e200)

   These errors are reported via [`SQLSTATE`](../postgresql-error-codes.md#errcodes-appendix).
  <a id="release-18-functions"></a>

#### Functions


-  Add function [`casefold()`](../../the-sql-language/functions-and-operators/string-functions-and-operators.md#functions-string-other) to allow for more sophisticated case-insensitive matching (Jeff Davis) [&sect;](https://postgr.es/c/bfc599206)

   This allows more accurate comparisons, i.e., a character can have multiple upper or lower case equivalents, or upper or lower case conversion changes the number of characters.
-  Allow [`MIN()`](../../the-sql-language/functions-and-operators/aggregate-functions.md#functions-aggregate-table)/[`MAX()`](../../the-sql-language/functions-and-operators/aggregate-functions.md#functions-aggregate-table) aggregates on arrays and composite types (Aleksander Alekseev, Marat Buharov) [&sect;](https://postgr.es/c/a0f1fce80) [&sect;](https://postgr.es/c/2d24fd942)
-  Add a `WEEK` option to [`EXTRACT()`](../../the-sql-language/functions-and-operators/date-time-functions-and-operators.md#functions-datetime-extract) (Tom Lane) [&sect;](https://postgr.es/c/6be39d77a)
-  Improve the output `EXTRACT(QUARTER ...)` for negative values (Tom Lane) [&sect;](https://postgr.es/c/6be39d77a)
-  Add roman numeral support to [`to_number()`](../../the-sql-language/functions-and-operators/data-type-formatting-functions.md#functions-formatting-table) (Hunaid Sohail) [&sect;](https://postgr.es/c/172e6b3ad)

   This is accessed via the `RN` pattern.
-  Add [`UUID`](../../the-sql-language/data-types/uuid-type.md#datatype-uuid) version 7 generation function [`uuidv7()`](../../the-sql-language/functions-and-operators/uuid-functions.md#func_uuid_gen_table) (Andrey Borodin) [&sect;](https://postgr.es/c/78c5e141e)

   This `UUID` value is temporally sortable. Function alias [`uuidv4()`](../../the-sql-language/functions-and-operators/uuid-functions.md#func_uuid_gen_table) has been added to explicitly generate version 4 UUIDs.
-  Add functions [`crc32()`](../../the-sql-language/functions-and-operators/binary-string-functions-and-operators.md#functions-binarystring-other) and [`crc32c()`](../../the-sql-language/functions-and-operators/binary-string-functions-and-operators.md#functions-binarystring-other) to compute CRC values (Aleksander Alekseev) [&sect;](https://postgr.es/c/760162fed)
-  Add math functions [`gamma()`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-func-table) and [`lgamma()`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-func-table) (Dean Rasheed) [&sect;](https://postgr.es/c/a3b6dfd41)
-  Allow `=>` syntax for named cursor arguments in [PL/pgSQL](../../server-programming/pl-pgsql-sql-procedural-language/index.md#plpgsql) (Pavel Stehule) [&sect;](https://postgr.es/c/246dedc5d)

   We previously only accepted `:=`.
-  Allow [`regexp_match[es]()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_like()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_replace()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_count()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_instr()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_substr()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_split_to_table()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp)/[`regexp_split_to_array()`](../../the-sql-language/functions-and-operators/pattern-matching.md#functions-posix-regexp) to use named arguments (Jian He) [&sect;](https://postgr.es/c/580f8727c)
  <a id="release-18-libpq"></a>

#### [Libpq]


-  Add function [`PQfullProtocolVersion()`](../../client-interfaces/libpq-c-library/connection-status-functions.md#libpq-PQfullProtocolVersion) to report the full, including minor, protocol version number (Jacob Champion, Jelte Fennema-Nio) [&sect;](https://postgr.es/c/cdb6b0fdb)
-  Add libpq connection [parameters](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-ssl-max-protocol-version) and [environment variables](../../client-interfaces/libpq-c-library/environment-variables.md#libpq-envars) to specify the minimum and maximum acceptable protocol version for connections (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/285613c60) [&sect;](https://postgr.es/c/507034910)
-  Report [search_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-search-path) changes to the client (Alexander Kukushkin, Jelte Fennema-Nio, Tomas Vondra) [&sect;](https://postgr.es/c/28a1121fd) [&sect;](https://postgr.es/c/0d06a7eac)
-  Add [`PQtrace()`](../../client-interfaces/libpq-c-library/control-functions.md#libpq-PQtrace) output for all message types, including authentication (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/ea92f3a0a) [&sect;](https://postgr.es/c/a5c6b8f22) [&sect;](https://postgr.es/c/b8b3f861f) [&sect;](https://postgr.es/c/e87c14b19) [&sect;](https://postgr.es/c/7adec2d5f)
-  Add libpq connection parameter [`sslkeylogfile`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-sslkeylogfile) which dumps out SSL key material (Abhishek Chanda, Daniel Gustafsson) [&sect;](https://postgr.es/c/2da74d8d6)

   This is useful for debugging.
-  Modify some libpq function signatures to use `int64_t` (Thomas Munro) [&sect;](https://postgr.es/c/3c86223c9)

   These previously used `pg_int64`, which is now deprecated.
  <a id="release-18-psql"></a>

#### [app-psql]


-  Allow psql to parse, bind, and close named prepared statements (Anthonin Bonnefoy, Michael Paquier) [&sect;](https://postgr.es/c/d55322b0d) [&sect;](https://postgr.es/c/fc39b286a)

   This is accomplished with new commands [`\parse`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-parse), [`\bind_named`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-bind-named), and [`\close_prepared`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-close-prepared).
-  Add psql backslash commands to allowing issuance of pipeline queries (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/41625ab8e) [&sect;](https://postgr.es/c/17caf6644) [&sect;](https://postgr.es/c/2cce0fe44)

   The new commands are [`\startpipeline`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-pipeline), `\syncpipeline`, `\sendpipeline`, `\endpipeline`, `\flushrequest`, `\flush`, and `\getresults`.
-  Allow adding pipeline status to the psql prompt and add related state variables (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/3ce357584)

   The new prompt character is `%P` and the new psql variables are [`PIPELINE_SYNC_COUNT`](../../reference/postgresql-client-applications/psql.md#app-psql-variables-pipeline-sync-count), [`PIPELINE_COMMAND_COUNT`](../../reference/postgresql-client-applications/psql.md#app-psql-variables-pipeline-command-count), and [`PIPELINE_RESULT_COUNT`](../../reference/postgresql-client-applications/psql.md#app-psql-variables-pipeline-result-count).
-  Allow adding the connection service name to the psql prompt or access it via psql variable (Michael Banck) [&sect;](https://postgr.es/c/477728b5d)
-  Add psql option to use expanded mode on all list commands (Dean Rasheed) [&sect;](https://postgr.es/c/00f4c2959)

   Adding backslash suffix `x` enables this.
-  Change psql's [\conninfo](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-conninfo) to use tabular format and include more information (Álvaro Herrera, Maiquel Grassi, Hunaid Sohail) [&sect;](https://postgr.es/c/bba2fbc62)
-  Add function's leakproof indicator to psql's [`\df+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-df-lc), `\do+`, `\dAo+`, and `\dC+` outputs (Yugo Nagata) [&sect;](https://postgr.es/c/2355e5111)
-  Add access method details for partitioned relations in [`\dP+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-dp-uc) (Justin Pryzby) [&sect;](https://postgr.es/c/978f38c77)
-  Add `default_version` to the psql [`\dx`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-dx-lc) extension output (Magnus Hagander) [&sect;](https://postgr.es/c/d696406a9)
-  Add psql variable [WATCH_INTERVAL](../../reference/postgresql-client-applications/psql.md#app-psql-variables-watch-interval) to set the default [`\watch`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-watch) wait time (Daniel Gustafsson) [&sect;](https://postgr.es/c/1a759c832)
  <a id="release-18-server-apps"></a>

#### Server Applications


-  Change [app-initdb](../../reference/postgresql-server-applications/initdb.md#app-initdb) to default to enabling checksums (Greg Sabino Mullane) [&sect;](https://postgr.es/c/983a588e0) [&sect;](https://postgr.es/c/04bec894a)

   The new initdb option `--no-data-checksums` disables checksums.
-  Add initdb option `--no-sync-data-files` to avoid syncing heap/index files (Nathan Bossart) [&sect;](https://postgr.es/c/cf131fa94)

   initdb option `--no-sync` is still available to avoid syncing any files.
-  Add [app-vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) option `--missing-stats-only` to compute only missing optimizer statistics (Corey Huinker, Nathan Bossart) [&sect;](https://postgr.es/c/edba754f0) [&sect;](https://postgr.es/c/987910502)

   This option can only be run by superusers and can only be used with options `--analyze-only` and `--analyze-in-stages`.
-  Add [app-pgcombinebackup](../../reference/postgresql-client-applications/pg_combinebackup.md#app-pgcombinebackup) option `-k`/`--link` to enable hard linking (Israel Barth Rubio, Robert Haas) [&sect;](https://postgr.es/c/99aeb8470)

   Only some files can be hard linked. This should not be used if the backups will be used independently.
-  Allow [app-pgverifybackup](../../reference/postgresql-client-applications/pg_verifybackup.md#app-pgverifybackup) to verify tar-format backups (Amul Sul) [&sect;](https://postgr.es/c/8dfd31290)
-  If [app-pgrewind](../../reference/postgresql-server-applications/pg_rewind.md#app-pgrewind)'s `--source-server` specifies a database name, use it in `--write-recovery-conf` output (Masahiko Sawada) [&sect;](https://postgr.es/c/4ecdd4110)
-  Add [app-pgresetwal](../../reference/postgresql-server-applications/pg_resetwal.md#app-pgresetwal) option `--char-signedness` to change the default `char` signedness (Masahiko Sawada) [&sect;](https://postgr.es/c/30666d185)
 <a id="release-18-pgdump"></a>

##### [pg_dump]/[pg_dumpall]/[pg_restore]


-  Add [app-pgdump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump) option `--statistics` (Jeff Davis) [&sect;](https://postgr.es/c/bde2fb797) [&sect;](https://postgr.es/c/a3e8dc143)
-  Add pg_dump and [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall) option `--sequence-data` to dump sequence data that would normally be excluded (Nathan Bossart) [&sect;](https://postgr.es/c/9c49f0e8c) [&sect;](https://postgr.es/c/acea3fc49)
-  Add [app-pgdump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump), [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall), and [app-pgrestore](../../reference/postgresql-client-applications/pg_restore.md#app-pgrestore) options `--statistics-only`, `--no-statistics`, `--no-data`, and `--no-schema` (Corey Huinker, Jeff Davis) [&sect;](https://postgr.es/c/1fd1bd871)
-  Add option `--no-policies` to disable row level security policy processing in [app-pgdump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump), [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall), [app-pgrestore](../../reference/postgresql-client-applications/pg_restore.md#app-pgrestore) (Nikolay Samokhvalov) [&sect;](https://postgr.es/c/cd3c45125)

   This is useful for migrating to systems with different policies.
  <a id="release-18-pgupgrade"></a>

##### [pgupgrade]


-  Allow pg_upgrade to preserve optimizer statistics (Corey Huinker, Jeff Davis, Nathan Bossart) [&sect;](https://postgr.es/c/1fd1bd871) [&sect;](https://postgr.es/c/c9d502eb6) [&sect;](https://postgr.es/c/d5f1b6a75) [&sect;](https://postgr.es/c/1fd1bd871)

   Extended statistics are not preserved. Also add pg_upgrade option `--no-statistics` to disable statistics preservation.
-  Allow pg_upgrade to process database checks in parallel (Nathan Bossart) [&sect;](https://postgr.es/c/40e2e5e92) [&sect;](https://postgr.es/c/6d3d2e8e5) [&sect;](https://postgr.es/c/7baa36de5) [&sect;](https://postgr.es/c/46cad8b31) [&sect;](https://postgr.es/c/6ab8f27bc) [&sect;](https://postgr.es/c/bbf83cab9) [&sect;](https://postgr.es/c/9db3018cf) [&sect;](https://postgr.es/c/c34eabfbb) [&sect;](https://postgr.es/c/cf2f82a37) [&sect;](https://postgr.es/c/f93f5f7b9) [&sect;](https://postgr.es/c/c880cf258)

   This is controlled by the existing `--jobs` option.
-  Add pg_upgrade option `--swap` to swap directories rather than copy, clone, or link files (Nathan Bossart) [&sect;](https://postgr.es/c/626d7236b)

   This mode is potentially the fastest.
-  Add pg_upgrade option `--set-char-signedness` to set the default `char` signedness of new cluster (Masahiko Sawada) [&sect;](https://postgr.es/c/a8238f87f) [&sect;](https://postgr.es/c/1aab68059)

   This is to handle cases where a pre-PostgreSQL 18 cluster's default CPU signedness does not match the new cluster.
  <a id="release-18-logicalrep-app"></a>

##### Logical Replication Applications


-  Add [app-pgcreatesubscriber](../../reference/postgresql-server-applications/pg_createsubscriber.md#app-pgcreatesubscriber) option `--all` to create logical replicas for all databases (Shubham Khanna) [&sect;](https://postgr.es/c/fb2ea12f4)
-  Add pg_createsubscriber option `--clean` to remove publications (Shubham Khanna) [&sect;](https://postgr.es/c/e5aeed4b8) [&sect;](https://postgr.es/c/60dda7bbc)
-  Add pg_createsubscriber option `--enable-two-phase` to enable prepared transactions (Shubham Khanna) [&sect;](https://postgr.es/c/e117cfb2f)
-  Add [app-pgrecvlogical](../../reference/postgresql-client-applications/pg_recvlogical.md#app-pgrecvlogical) option `--enable-failover` to specify failover slots (Hayato Kuroda) [&sect;](https://postgr.es/c/cf2655a90)

   Also add option `--enable-two-phase` as a synonym for `--two-phase`, and deprecate the latter.
-  Allow pg_recvlogical `--drop-slot` to work without `--dbname` (Hayato Kuroda) [&sect;](https://postgr.es/c/c68100aa4)
   <a id="release-18-source-code"></a>

#### Source Code


-  Separate the loading and running of [injection points](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-injection-points) (Michael Paquier, Heikki Linnakangas) [&sect;](https://postgr.es/c/4b211003e) [&sect;](https://postgr.es/c/a0a5869a8)

   Injection points can now be created, but not run, via [`INJECTION_POINT_LOAD()`](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-injection-points), and such injection points can be run via [`INJECTION_POINT_CACHED()`](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-injection-points).
-  Support runtime arguments in injection points (Michael Paquier) [&sect;](https://postgr.es/c/371f2db8b)
-  Allow inline injection point test code with [`IS_INJECTION_POINT_ATTACHED()`](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-injection-points) (Heikki Linnakangas) [&sect;](https://postgr.es/c/20e0e7da9)
-  Improve the performance of processing long [`JSON`](../../the-sql-language/data-types/json-types.md#datatype-json) strings using SIMD (Single Instruction Multiple Data) (David Rowley) [&sect;](https://postgr.es/c/ca6fde922)
-  Speed up CRC32C calculations using x86 AVX-512 instructions (Raghuveer Devulapalli, Paul Amonson) [&sect;](https://postgr.es/c/3c6e8c123)
-  Add ARM Neon and SVE CPU intrinsics for popcount (integer bit counting) (Chiranmoy Bhattacharya, Devanga Susmitha, Rama Malladi) [&sect;](https://postgr.es/c/6be53c276) [&sect;](https://postgr.es/c/519338ace)
-  Improve the speed of numeric multiplication and division (Joel Jacobson, Dean Rasheed) [&sect;](https://postgr.es/c/ca481d3c9) [&sect;](https://postgr.es/c/c4e44224c) [&sect;](https://postgr.es/c/8dc28d7eb) [&sect;](https://postgr.es/c/9428c001f)
-  Add configure option [`--with-libnuma`](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-option-with-libnuma) to enable NUMA awareness (Jakub Wartak, Bertrand Drouvot) [&sect;](https://postgr.es/c/65c298f61) [&sect;](https://postgr.es/c/8cc139bec) [&sect;](https://postgr.es/c/ba2a3c230)

   The function [`pg_numa_available()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-session-table) reports on NUMA awareness, and system views [`pg_shmem_allocations_numa`](../../internals/system-views/pg_shmem_allocations_numa.md#view-pg-shmem-allocations-numa) and [`pg_buffercache_numa`](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache-pg-buffercache-numa) which report on shared memory distribution across NUMA nodes.
-  Add [TOAST](../../internals/database-physical-storage/toast.md#storage-toast) table to [`pg_index`](../../internals/system-catalogs/pg_index.md#catalog-pg-index) to allow for very large expression indexes (Nathan Bossart) [&sect;](https://postgr.es/c/b52c4fc3c)
-  Remove column [`pg_attribute`](../../internals/system-catalogs/pg_attribute.md#catalog-pg-attribute).`attcacheoff` (David Rowley) [&sect;](https://postgr.es/c/02a8d0c45)
-  Add column [`pg_class`](../../internals/system-catalogs/pg_class.md#catalog-pg-class).`relallfrozen` (Melanie Plageman) [&sect;](https://postgr.es/c/99f8f3fbb)
-  Add [`amgettreeheight`](../../internals/index-access-method-interface-definition/index.md#indexam), `amconsistentequality`, and `amconsistentordering` to the index access method API (Mark Dilger) [&sect;](https://postgr.es/c/56fead44d) [&sect;](https://postgr.es/c/af4002b38)
-  Add GiST support function [`stratnum()`](../../internals/built-in-index-access-methods/gist-indexes.md#gist-extensibility) (Paul A. Jungwirth) [&sect;](https://postgr.es/c/7406ab623)
-  Record the default CPU signedness of `char` in [app-pgcontroldata](../../reference/postgresql-server-applications/pg_controldata.md#app-pgcontroldata) (Masahiko Sawada) [&sect;](https://postgr.es/c/44fe30fda)
-  Add support for Python "Limited API" in [PL/Python](../../server-programming/pl-python-python-procedural-language/index.md#plpython) (Peter Eisentraut) [&sect;](https://postgr.es/c/72a3d0462) [&sect;](https://postgr.es/c/0793ab810)

   This helps prevent problems caused by Python 3.x version mismatches.
-  Change the minimum supported Python version to 3.6.8 (Jacob Champion) [&sect;](https://postgr.es/c/45363fca6)
-  Remove support for OpenSSL versions older than 1.1.1 (Daniel Gustafsson) [&sect;](https://postgr.es/c/a70e01d43) [&sect;](https://postgr.es/c/6c66b7443)
-  If LLVM is enabled, require version 14 or later (Thomas Munro) [&sect;](https://postgr.es/c/972c2cd28)
-  Add macro [`PG_MODULE_MAGIC_EXT`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info) to allow extensions to report their name and version (Andrei Lepikhov) [&sect;](https://postgr.es/c/9324c8c58)

   This information can be access via the new function [`pg_get_loaded_modules()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-session-table).
-  Document that [`SPI_connect()`](../../server-programming/server-programming-interface/interface-functions.md#spi-spi-connect)/[`SPI_connect_ext()`](../../server-programming/server-programming-interface/interface-functions.md#spi-spi-connect) always returns success (`SPI_OK_CONNECT`) (Stepan Neretin) [&sect;](https://postgr.es/c/218527d01)

   Errors are always reported via `ereport()`.
-  Add [documentation section](../../server-programming/extending-sql/c-language-functions.md#xfunc-api-abi-stability-guidance) about API and ABI compatibility (David Wheeler, Peter Eisentraut) [&sect;](https://postgr.es/c/e54a42ac9)
-  Remove the experimental designation of Meson builds on `Windows` (Aleksander Alekseev) [&sect;](https://postgr.es/c/5afaba629)
-  Remove configure options `--disable-spinlocks` and `--disable-atomics` (Thomas Munro) [&sect;](https://postgr.es/c/e25626677) [&sect;](https://postgr.es/c/813852613)

   Thirty-two-bit atomic operations are now required.
-  Remove support for the HPPA/PA-RISC architecture (Tom Lane) [&sect;](https://postgr.es/c/edadeb071)
  <a id="release-18-modules"></a>

#### Additional Modules


-  Add extension [pg_logicalinspect](../additional-supplied-modules-and-extensions/pg_logicalinspect-logical-decoding-components-inspection.md#pglogicalinspect) to inspect logical snapshots (Bertrand Drouvot) [&sect;](https://postgr.es/c/7cdfeee32)
-  Add extension [pg_overexplain](../additional-supplied-modules-and-extensions/pg_overexplain-allow-explain-to-dump-even-more-details.md#pgoverexplain) which adds debug details to [`EXPLAIN`](../../reference/sql-commands/explain.md#sql-explain) output (Robert Haas) [&sect;](https://postgr.es/c/8d5ceb113)
-  Add output columns to [`postgres_fdw_get_connections()`](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-functions) (Hayato Kuroda, Sagar Dilip Shedge) [&sect;](https://postgr.es/c/c297a47c5) [&sect;](https://postgr.es/c/857df3cef) [&sect;](https://postgr.es/c/4f08ab554) [&sect;](https://postgr.es/c/fe186bda7)

   New output column `used_in_xact` indicates if the foreign data wrapper is being used by a current transaction, `closed` indicates if it is closed, `user_name` indicates the user name, and `remote_backend_pid` indicates the remote backend process identifier.
-  Allow [SCRAM](../../server-administration/client-authentication/password-authentication.md#auth-password) authentication from the client to be passed to [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) servers (Matheus Alcantara, Peter Eisentraut) [&sect;](https://postgr.es/c/761c79508)

   This avoids storing postgres_fdw authentication information in the database, and is enabled with the postgres_fdw [`use_scram_passthrough`](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-option-use-scram-passthrough) connection option. libpq uses new connection parameters [scram_client_key](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-scram-client-key) and [scram_server_key](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-scram-server-key).
-  Allow SCRAM authentication from the client to be passed to [dblink](../additional-supplied-modules-and-extensions/dblink-connect-to-other-postgresql-databases.md#dblink) servers (Matheus Alcantara) [&sect;](https://postgr.es/c/3642df265)
-  Add `on_error` and `log_verbosity` options to [file_fdw](../additional-supplied-modules-and-extensions/file_fdw-access-data-files-in-the-servers-file-system.md#file-fdw) (Atsushi Torikoshi) [&sect;](https://postgr.es/c/a1c4c8a9e)

   These control how file_fdw handles and reports invalid file rows.
-  Add `reject_limit` to control the number of invalid rows file_fdw can ignore (Atsushi Torikoshi) [&sect;](https://postgr.es/c/6c8f67032)

   This is active when `ON_ERROR = 'ignore'`.
-  Add configurable variable `min_password_length` to [passwordcheck](../additional-supplied-modules-and-extensions/passwordcheck-verify-password-strength.md#passwordcheck) (Emanuele Musella, Maurizio Boriani) [&sect;](https://postgr.es/c/f7e1b3828)

   This controls the minimum password length.
-  Have [pgbench](../../reference/postgresql-client-applications/pgbench.md#pgbench) report the number of failed, retried, or skipped transactions in per-script reports (Yugo Nagata) [&sect;](https://postgr.es/c/cae0f3c40)
-  Add [isn](../additional-supplied-modules-and-extensions/isn-data-types-for-international-standard-numbers-isbn-ean-upc-etc.md#isn) server variable `weak` to control invalid check digit acceptance (Viktor Holmberg) [&sect;](https://postgr.es/c/448904423)

   This was previously only controlled by function [`isn_weak()`](../additional-supplied-modules-and-extensions/isn-data-types-for-international-standard-numbers-isbn-ean-upc-etc.md#isn-functions).
-  Allow values to be sorted to speed [btree_gist](../additional-supplied-modules-and-extensions/btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist) index builds (Bernd Helmle, Andrey Borodin) [&sect;](https://postgr.es/c/e4309f73f)
-  Add [amcheck](../additional-supplied-modules-and-extensions/amcheck-tools-to-verify-table-and-index-consistency.md#amcheck) check function [`gin_index_check()`](../additional-supplied-modules-and-extensions/amcheck-tools-to-verify-table-and-index-consistency.md#amcheck-functions) to verify `GIN` indexes (Grigory Kryachko, Heikki Linnakangas, Andrey Borodin) [&sect;](https://postgr.es/c/14ffaece0)
-  Add functions [`pg_buffercache_evict_relation()`](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache-pg-buffercache-evict-relation) and [`pg_buffercache_evict_all()`](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache-pg-buffercache-evict-all) to evict unpinned shared buffers (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/dcf7e1697)

   The existing function [`pg_buffercache_evict()`](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache-pg-buffercache-evict) now returns the buffer flush status.
-  Allow extensions to install custom [sql-explain](../../reference/sql-commands/explain.md#sql-explain) options (Robert Haas, Sami Imseih) [&sect;](https://postgr.es/c/c65bc2e1d) [&sect;](https://postgr.es/c/4fd02bf7c) [&sect;](https://postgr.es/c/50ba65e73)
-  Allow extensions to use the server's cumulative statistics API (Michael Paquier) [&sect;](https://postgr.es/c/7949d9594) [&sect;](https://postgr.es/c/2eff9e678)
 <a id="release-18-pgstatstatements"></a>

##### [pgstatstatements]


-  Allow the queries of [sql-createtableas](../../reference/sql-commands/create-table-as.md#sql-createtableas) and [sql-declare](../../reference/sql-commands/declare.md#sql-declare) to be tracked by pg_stat_statements (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/6b652e6ce)

   They are also now assigned query ids.
-  Allow the parameterization of [sql-set](../../reference/sql-commands/set.md#sql-set) values in pg_stat_statements (Greg Sabino Mullane, Michael Paquier) [&sect;](https://postgr.es/c/dc6851596)

   This reduces the bloat caused by `SET` statements with differing constants.
-  Add [`pg_stat_statements`](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements-pg-stat-statements) columns to report parallel activity (Guillaume Lelarge) [&sect;](https://postgr.es/c/cf54a2c00)

   The new columns are `parallel_workers_to_launch` and `parallel_workers_launched`.
-  Add `pg_stat_statements`.`wal_buffers_full` to report full WAL buffers (Bertrand Drouvot) [&sect;](https://postgr.es/c/ce5bcc4a9)
  <a id="release-18-pgcrypto"></a>

##### [pgcrypto]


-  Add pgcrypto algorithms [`sha256crypt`](../additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto-crypt-algorithms) and [`sha512crypt`](../additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto-crypt-algorithms) (Bernd Helmle) [&sect;](https://postgr.es/c/749a9e20c)
-  Add [CFB](../additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto-raw-enc-funcs) mode to pgcrypto encryption and decryption (Umar Hayat) [&sect;](https://postgr.es/c/9ad1b3d01)
-  Add function [`fips_mode()`](../additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto-openssl-support-funcs) to report the server's FIPS mode (Daniel Gustafsson) [&sect;](https://postgr.es/c/924d89a35)
-  Add pgcrypto server variable [`builtin_crypto_enabled`](../additional-supplied-modules-and-extensions/pgcrypto-cryptographic-functions.md#pgcrypto-configuration-parameters-builtin_crypto_enabled) to allow disabling builtin non-FIPS mode cryptographic functions (Daniel Gustafsson, Joe Conway) [&sect;](https://postgr.es/c/035f99cbe)

   This is useful for guaranteeing FIPS mode behavior.
    <a id="release-18-acknowledgements"></a>

### Acknowledgments


 The following individuals (in alphabetical order) have contributed to this release as patch authors, committers, reviewers, testers, or reporters of issues.


- Abhishek Chanda
- Adam Guo
- Adam Rauch
- Aidar Imamov
- Ajin Cherian
- Alastair Turner
- Alec Cozens
- Aleksander Alekseev
- Alena Rybakina
- Alex Friedman
- Alex Richman
- Alexander Alehin
- Alexander Borisov
- Alexander Korotkov
- Alexander Kozhemyakin
- Alexander Kukushkin
- Alexander Kuzmenkov
- Alexander Kuznetsov
- Alexander Lakhin
- Alexander Pyhalov
- Alexandra Wang
- Alexey Dvoichenkov
- Alexey Makhmutov
- Alexey Shishkin
- Ali Akbar
- Álvaro Herrera
- Álvaro Mongil
- Amit Kapila
- Amit Langote
- Amul Sul
- Andreas Karlsson
- Andreas Scherbaum
- Andreas Ulbrich
- Andrei Lepikhov
- Andres Freund
- Andrew
- Andrew Bille
- Andrew Dunstan
- Andrew Jackson
- Andrew Kane
- Andrew Watkins
- Andrey Borodin
- Andrey Chudnovsky
- Andrey Rachitskiy
- Andrey Rudometov
- Andy Alsup
- Andy Fan
- Anthonin Bonnefoy
- Anthony Hsu
- Anthony Leung
- Anton Melnikov
- Anton Voloshin
- Antonin Houska
- Antti Lampinen
- Arseniy Mukhin
- Artur Zakirov
- Arun Thirupathi
- Ashutosh Bapat
- Asphator
- Atsushi Torikoshi
- Avi Weinberg
- Aya Iwata
- Ayush Tiwari
- Ayush Vatsa
- Bastien Roucariès
- Ben Peachey Higdon
- Benoit Lobréau
- Bernd Helmle
- Bernd Reiß
- Bernhard Wiedemann
- Bertrand Drouvot
- Bertrand Mamasam
- Bharath Rupireddy
- Bogdan Grigorenko
- Boyu Yang
- Braulio Fdo Gonzalez
- Bruce Momjian
- Bykov Ivan
- Cameron Vogt
- Cary Huang
- Cédric Villemain
- Cees van Zeeland
- ChangAo Chen
- Chao Li
- Chapman Flack
- Charles Samborski
- Chengwen Wu
- Chengxi Sun
- Chiranmoy Bhattacharya
- Chris Gooch
- Christian Charukiewicz
- Christoph Berg
- Christophe Courtois
- Christopher Inokuchi
- Clemens Ruck
- Corey Huinker
- Craig Milhiser
- Crisp Lee
- Dagfinn Ilmari Mannsåker
- Daniel Elishakov
- Daniel Gustafsson
- Daniel Vérité
- Daniel Westermann
- Daniele Varrazzo
- Daniil Davydov
- Daria Shanina
- Dave Cramer
- Dave Page
- David Benjamin
- David Christensen
- David Fiedler
- David G. Johnston
- David Geier
- David Rowley
- David Steele
- David Wheeler
- David Zhang
- Davinder Singh
- Dean Rasheed
- Devanga Susmitha
- Devrim Gündüz
- Dian Fay
- Dilip Kumar
- Dimitrios Apostolou
- Dipesh Dhameliya
- Dmitrii Bondar
- Dmitry Dolgov
- Dmitry Koval
- Dmitry Kovalenko
- Dmitry Yurichev
- Dominique Devienne
- Donghang Lin
- Dorjpalam Batbaatar
- Drew Callahan
- Duncan Sands
- Dwayne Towell
- Dzmitry Jachnik
- Egor Chindyaskin
- Egor Rogov
- Emanuel Ionescu
- Emanuele Musella
- Emre Hasegeli
- Eric Cyr
- Erica Zhang
- Erik Nordström
- Erik Rijkers
- Erik Wienhold
- Erki Eessaar
- Ethan Mertz
- Etienne LAFARGE
- Etsuro Fujita
- Euler Taveira
- Evan Si
- Evgeniy Gorbanev
- Fabio R. Sluzala
- Fabrízio de Royes Mello
- Feike Steenbergen
- Feliphe Pozzer
- Felix
- Fire Emerald
- Florents Tselai
- Francesco Degrassi
- Frank Streitzig
- Frédéric Yhuel
- Fredrik Widlert
- Gabriele Bartolini
- Gavin Panella
- Geoff Winkless
- George MacKerron
- Gilles Darold
- Grant Gryczan
- Greg Burd
- Greg Sabino Mullane
- Greg Stark
- Grigory Kryachko
- Guillaume Lelarge
- Gunnar Morling
- Gunnar Wagner
- Gurjeet Singh
- Haifang Wang
- Hajime Matsunaga
- Hamid Akhtar
- Hannu Krosing
- Hari Krishna Sunder
- Haruka Takatsuka
- Hayato Kuroda
- Heikki Linnakangas
- Hironobu Suzuki
- Holger Jakobs
- Hubert Lubaczewski
- Hugo Dubois
- Hugo Zhang
- Hunaid Sohail
- Hywel Carver
- Ian Barwick
- Ibrar Ahmed
- Igor Gnatyuk
- Igor Korot
- Ilia Evdokimov
- Ilya Gladyshev
- Ilyasov Ian
- Imran Zaheer
- Isaac Morland
- Israel Barth Rubio
- Ivan Kush
- Jacob Brazeal
- Jacob Champion
- Jaime Casanova
- Jakob Egger
- Jakub Wartak
- James Coleman
- James Hunter
- Jan Behrens
- Japin Li
- Jason Smith
- Jayesh Dehankar
- Jeevan Chalke
- Jeff Davis
- Jehan-Guillaume de Rorthais
- Jelte Fennema-Nio
- Jian He
- Jianghua Yang
- Jiao Shuntian
- Jim Jones
- Jim Nasby
- Jingtang Zhang
- Jingzhou Fu
- Joe Conway
- Joel Jacobson
- John Hutchins
- John Naylor
- Jonathan Katz
- Jorge Solórzano
- José Villanova
- Josef Šimánek
- Joseph Koshakow
- Julien Rouhaud
- Junwang Zhao
- Justin Pryzby
- Kaido Vaikla
- Kaimeh
- Karina Litskevich
- Karthik S
- Kartyshov Ivan
- Kashif Zeeshan
- Keisuke Kuroda
- Kevin Hale Boyes
- Kevin K Biju
- Kirill Reshke
- Kirill Zdornyy
- Koen De Groote
- Koichi Suzuki
- Koki Nakamura
- Konstantin Knizhnik
- Kouhei Sutou
- Kuntal Ghosh
- Kyotaro Horiguchi
- Lakshmi Narayana Velayudam
- Lars Kanis
- Laurence Parry
- Laurenz Albe
- Lele Gaifax
- Li Yong
- Lilian Ontowhee
- Lingbin Meng
- Luboslav Špilák
- Luca Vallisa
- Lukas Fittl
- Maciek Sakrejda
- Magnus Hagander
- Mahendra Singh Thalor
- Mahendrakar Srinivasarao
- Maiquel Grassi
- Maksim Korotkov
- Maksim Melnikov
- Man Zeng
- Marat Buharov
- Marc Balmer
- Marco Nenciarini
- Marcos Pegoraro
- Marina Polyakova
- Mark Callaghan
- Mark Dilger
- Marlene Brandstaetter
- Marlene Reiterer
- Martin Rakhmanov
- Masahiko Sawada
- Masahiro Ikeda
- Masao Fujii
- Mason Mackaman
- Mat Arye
- Matheus Alcantara
- Mats Kindahl
- Matthew Gabeler-Lee
- Matthew Kim
- Matthew Sterrett
- Matthew Woodcraft
- Matthias van de Meent
- Matthieu Denais
- Maurizio Boriani
- Max Johnson
- Max Madden
- Maxim Boguk
- Maxim Orlov
- Maximilian Chrzan
- Melanie Plageman
- Melih Mutlu
- Mert Alev
- Michael Banck
- Michael Bondarenko
- Michael Christofides
- Michael Guissine
- Michael Harris
- Michaël Paquier
- Michail Nikolaev
- Michal Kleczek
- Michel Pelletier
- Mikaël Gourlaouen
- Mikhail Gribkov
- Mikhail Kot
- Milosz Chmura
- Muralikrishna Bandaru
- Murat Efendioglu
- Mutaamba Maasha
- Naeem Akhter
- Nat Makarevitch
- Nathan Bossart
- Navneet Kumar
- Nazir Bilal Yavuz
- Neil Conway
- Niccolò Fei
- Nick Davies
- Nicolas Maus
- Niek Brasa
- Nikhil Raj
- Nikita
- Nikita Kalinin
- Nikita Malakhov
- Nikolay Samokhvalov
- Nikolay Shaplov
- Nisha Moond
- Nitin Jadhav
- Nitin Motiani
- Noah Misch
- Noboru Saito
- Noriyoshi Shinoda
- Ole Peder Brandtzæg
- Oleg Sibiryakov
- Oleg Tselebrovskiy
- Olleg Samoylov
- Onder Kalaci
- Ondrej Navratil
- Patrick Stählin
- Paul Amonson
- Paul Jungwirth
- Paul Ramsey
- Pavel Borisov
- Pavel Luzanov
- Pavel Nekrasov
- Pavel Stehule
- Peter Eisentraut
- Peter Geoghegan
- Peter Mittere
- Peter Smith
- Phil Eaton
- Philipp Salvisberg
- Philippe Beaudoin
- Pierre Giraud
- Pixian Shi
- Polina Bungina
- Przemyslaw Sztoch
- Quynh Tran
- Rafia Sabih
- Raghuveer Devulapalli
- Rahila Syed
- Rama Malladi
- Ran Benita
- Ranier Vilela
- Renan Alves Fonseca
- Richard Guo
- Richard Neill
- Rintaro Ikeda
- Robert Haas
- Robert Treat
- Robins Tharakan
- Roman Zharkov
- Ronald Cruz
- Ronan Dunklau
- Rui Zhao
- Rushabh Lathia
- Rustam Allakov
- Ryo Kanbayashi
- Ryohei Takahashi
- RyotaK
- Sagar Dilip Shedge
- Salvatore Dipietro
- Sam Gabrielsson
- Sam James
- Sameer Kumar
- Sami Imseih
- Samuel Thibault
- Satyanarayana Narlapuram
- Sebastian Skalacki
- Senglee Choi
- Sergei Kornilov
- Sergey Belyashov
- Sergey Dudoladov
- Sergey Prokhorenko
- Sergey Sargsyan
- Sergey Soloviev
- Sergey Tatarintsev
- Shaik Mohammad Mujeeb
- Shawn McCoy
- Shenhao Wang
- Shihao Zhong
- Shinya Kato
- Shlok Kyal
- Shubham Khanna
- Shveta Malik
- Simon Riggs
- Smolkin Grigory
- Sofia Kopikova
- Song Hongyu
- Song Jinzhou
- Soumyadeep Chakraborty
- Sravan Kumar
- Srinath Reddy
- Stan Hu
- Stepan Neretin
- Stephen Fewer
- Stephen Frost
- Steve Chavez
- Steven Niu
- Suraj Kharage
- Sven Klemm
- Takamichi Osumi
- Takeshi Ideriha
- Tatsuo Ishii
- Ted Yu
- Tels
- Tender Wang
- Teodor Sigaev
- Thom Brown
- Thomas Baehler
- Thomas Krennwallner
- Thomas Munro
- Tim Wood
- Timur Magomedov
- Tobias Wendorff
- Todd Cook
- Tofig Aliev
- Tom Lane
- Tomas Vondra
- Tomasz Rybak
- Tomasz Szypowski
- Torsten Foertsch
- Toshi Harada
- Tristan Partin
- Triveni N
- Umar Hayat
- Vallimaharajan G
- Vasya Boytsov
- Victor Yegorov
- Vignesh C
- Viktor Holmberg
- Vinícius Abrahão
- Vinod Sridharan
- Virender Singla
- Vitaly Davydov
- Vladlen Popolitov
- Vladyslav Nebozhyn
- Walid Ibrahim
- Webbo Han
- Wenhui Qiu
- Will Mortensen
- Will Storey
- Wolfgang Walther
- Xin Zhang
- Xing Guo
- Xuneng Zhou
- Yan Chengpen
- Yang Lei
- Yaroslav Saburov
- Yaroslav Syrytsia
- Yasir Hussain
- Yasuo Honda
- Yogesh Sharma
- Yonghao Lee
- Yoran Heling
- Yu Liang
- Yugo Nagata
- Yuhang Qiu
- Yuki Seino
- Yura Sokolov
- Yurii Rashkovskii
- Yushi Ogiwara
- Yusuke Sugie
- Yuta Katsuragi
- Yuto Sasaki
- Yuuki Fujii
- Yuya Watari
- Zane Duffield
- Zeyuan Hu
- Zhang Mingli
- Zhihong Yu
- Zhijie Hou
- Zsolt Parragi
