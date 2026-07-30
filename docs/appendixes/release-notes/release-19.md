<a id="release-19"></a>

## Release 19


**Release date:.**


2026-??-??, AS OF 2026-06-17
  <a id="release-19-highlights"></a>

### Overview


 PostgreSQL 19 contains many new features and enhancements, including:


- *fill in later*


 The above items and other new features of PostgreSQL 19 are explained in more detail in the sections below.
  <a id="release-19-migration"></a>

### Migration to Version 19


 A dump/restore using [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall) or use of [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) or logical replication is required for those wishing to migrate data from any previous release. See [Upgrading a PostgreSQL Cluster](../../server-administration/server-setup-and-operation/upgrading-a-postgresql-cluster.md#upgrading) for general information on migrating to new major releases.


 Version 19 contains a number of changes that may affect compatibility with previous releases. Observe the following incompatibilities:


-  Add server variable [`password_expiration_warning_threshold`](../../server-administration/server-configuration/connections-and-authentication.md#guc-password-expiration-warning-threshold) to warn about password expiration (Gilles Darold, Nathan Bossart) [&sect;](https://postgr.es/c/1d92e0c2c)

   The default warning period is seven days.
-  Issue a warning after successful MD5 password authentication (Nathan Bossart) [&sect;](https://postgr.es/c/bc60ee860)

   The warning can be disabled via server variable [`md5_password_warnings`](../../server-administration/server-configuration/connections-and-authentication.md#guc-md5-password-warnings). MD5 passwords were marked as deprecated in PostgreSQL 18.
-  Remove RADIUS support (Thomas Munro) [&sect;](https://postgr.es/c/a1643d40b)

   PostgreSQL only supported RADIUS over UDP, which is unfixably insecure.
-  Force [`standard_conforming_strings`](../../server-administration/server-configuration/version-and-platform-compatibility.md#guc-standard-conforming-strings) to always be `on` in the database server (Tom Lane) [&sect;](https://postgr.es/c/457620845)

   Dumps created using pre-PostgreSQL 19 versions of [pg_dump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump) or [pg_dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall), and using `standard_conforming_strings = off`, will not properly load into PostgreSQL 19 and later servers. Users should create dumps using PostgreSQL 19 or later versions of these applications, or use `standard_conforming_strings = on`.

   Client applications still support operations with servers having `standard_conforming_strings = off`, for compatibility with old servers. The server variable `escape_string_warning` has been removed as unnecessary.
-  Disallow carriage returns and line feeds in database, role, and tablespace names (Mahendra Singh Thalor) [&sect;](https://postgr.es/c/b380a56a3)

   [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) will also disallow upgrading of clusters that use such names. This was changed to avoid security problems.
-  Change the default index opclasses for [`inet`](../../the-sql-language/data-types/network-address-types.md#datatype-net-types) and [`cidr`](../../the-sql-language/data-types/network-address-types.md#datatype-net-types) data types from [btree_gist](../additional-supplied-modules-and-extensions/btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist) to [GiST](../../internals/built-in-index-access-methods/gist-indexes.md#gist) (Tom Lane) [&sect;](https://postgr.es/c/b3b0b4571) [&sect;](https://postgr.es/c/b352d3d80)

   The [btree_gist](../additional-supplied-modules-and-extensions/btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist) `inet`/`cidr` opclasses are broken because they can exclude rows that should be returned. [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) will disallow upgrading of clusters with [btree_gist](../additional-supplied-modules-and-extensions/btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist) `inet`/`cidr` indexes.
-  Stop reordering non-schema objects created by [`CREATE SCHEMA`](../../reference/sql-commands/create-schema.md#sql-createschema) (Tom Lane, Jian He) [&sect;](https://postgr.es/c/a9c350d9e) [&sect;](https://postgr.es/c/404db8f9e)

   The goal of the reordering was to avoid dependencies, but it was imperfect. PostgreSQL now uses the specified object ordering, except for foreign keys which are created last.
-  Disallow system columns from being used in [`COPY FROM ... WHERE`](../../reference/sql-commands/copy.md#sql-copy) (Tom Lane) [&sect;](https://postgr.es/c/21c69dc73)

   The values of such columns were not well-defined.
-  Change a [`json_array()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-json) call which returns no rows to return an empty JSON array (Richard Guo) [&sect;](https://postgr.es/c/8d829f5a0)

   This previously returned `NULL`.
-  Cause transactions to pass their `READ ONLY` and `DEFERRABLE` status to [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) sessions (Etsuro Fujita) [&sect;](https://postgr.es/c/de28140de)

   This means `READ ONLY` transactions can no longer modify rows processed by [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) sessions.
-  Change default of [`max_locks_per_transaction`](../../server-administration/server-configuration/lock-management.md#guc-max-locks-per-transaction) from 64 to 128 (Heikki Linnakangas) [&sect;](https://postgr.es/c/79534f906)

   Lock size allocation has changed, so effectively settings must now be doubled to match their capacity in previous releases.
-  Change [JIT](../../server-administration/just-in-time-compilation-jit/index.md#jit) to be disabled by default (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/7f8c88c2b)

   Previously JIT was enabled by default, and activated based on optimizer costs, but this costing has been determined to be unreliable. This change requires sites that are doing many large analytical queries to manually enable JIT.
-  Rename column `sync_error_count` to `sync_table_error_count` in system view [`pg_stat_subscription_stats`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-subscription-stats) (Vignesh C) [&sect;](https://postgr.es/c/3edaf29fa)

   This is necessary since sequence errors are now tracked separately.
-  Rename wait event type `BUFFERPIN` to `BUFFER` (Andres Freund) [&sect;](https://postgr.es/c/6c5c393b7)
-  Change index access method handlers to use a static [`IndexAmRoutines`](../../internals/index-access-method-interface-definition/index.md#indexam) structure, rather than dynamically allocated ones (Matthias van de Meent) [&sect;](https://postgr.es/c/bc6374cd7)
-  Remove optimizer hook get_relation_info_hook and add better-placed hook build_simple_rel_hook (Robert Haas) [&sect;](https://postgr.es/c/91f33a2ae)
-  Remove `MULE_INTERNAL` encoding (Thomas Munro) [&sect;](https://postgr.es/c/77645d44e)

   This encoding was complex and rarely used. Databases using it will need to be dumped and restored with a different encoding.
  <a id="release-19-changes"></a>

### Changes


 Below you will find a detailed account of the changes between PostgreSQL 19 and the previous major release.
 <a id="release-19-server"></a>

#### Server
  <a id="release-19-optimizer"></a>

##### Optimizer


-  Allow `NOT IN` clauses to be converted to more efficient `ANTI JOIN`s when NULLs are not present (Richard Guo) [&sect;](https://postgr.es/c/383eb21eb)
-  Allow more `LEFT JOIN`s to be converted to `ANTI JOIN`s (Tender Wang, Richard Guo) [&sect;](https://postgr.es/c/cf74558fe)
-  Allow use of Memoize for `ANTI JOIN`s with unique inner sides (Richard Guo) [&sect;](https://postgr.es/c/0da29e4cb)
-  Allow some aggregate processing to be performed before joins (Richard Guo, Antonin Houska) [&sect;](https://postgr.es/c/8e1185910) [&sect;](https://postgr.es/c/bd94845e8) [&sect;](https://postgr.es/c/3a08a2a8b)

   This can reduce the number of rows needed to be processed.
-  Improve hash join's handling of tuples with `NULL` join keys (Tom Lane) [&sect;](https://postgr.es/c/1811f1af9)
-  Improve the planning of semijoins (Richard Guo) [&sect;](https://postgr.es/c/24225ad9a)
-  Allow Append and MergeAppend to consider explicit incremental sorts (Richard Guo) [&sect;](https://postgr.es/c/55a780e94)
-  Convert [`IS [NOT] DISTINCT FROM NULL`](../../the-sql-language/functions-and-operators/comparison-functions-and-operators.md#functions-comparison-pred-table) to `IS [NOT] NULL` during constant folding (Richard Guo) [&sect;](https://postgr.es/c/f41ab5157)

   The latter form is more easily optimized.
-  Simplify [`IS [NOT] DISTINCT FROM`](../../the-sql-language/functions-and-operators/comparison-functions-and-operators.md#functions-comparison-pred-table) to equality/inequality operators when inputs are proven non-nullable (Richard Guo) [&sect;](https://postgr.es/c/0a3796125)
-  Perform earlier constant folding of *var* `IS [NOT] NULL` in the optimizer (Richard Guo) [&sect;](https://postgr.es/c/e2debb643)

   This allows for later optimizations.
-  Simplify [`COALESCE()`](../../the-sql-language/functions-and-operators/conditional-expressions.md#functions-coalesce-nvl-ifnull) and `ROW(...) IS [NOT] NULL` to avoid evaluating unnecessary arguments (Richard Guo) [&sect;](https://postgr.es/c/10c4fe074) [&sect;](https://postgr.es/c/cb7b7ec7a)
-  Simplify `IS [NOT] TRUE/FALSE/UNKNOWN` to plain boolean expressions when the input is proven non-nullable (Richard Guo) [&sect;](https://postgr.es/c/0aaf0de7f)
-  Speed up join selectivity computations for large optimizer statistics targets (Ilia Evdokimov, David Geier) [&sect;](https://postgr.es/c/057012b20)
-  Enable proper optimizer statistics for functions returning boolean values (Tom Lane) [&sect;](https://postgr.es/c/1eccb9315)
-  Allow [extended statistics](../../reference/sql-commands/create-statistics.md#sql-createstatistics) on virtual generated columns (Yugo Nagata) [&sect;](https://postgr.es/c/f7f4052a4)
-  Allow function [`pg_restore_extended_stats()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-statsmod) to restore optimizer extended statistics (Corey Huinker, Michael Paquier, Chao Li) [&sect;](https://postgr.es/c/0e80f3f88) [&sect;](https://postgr.es/c/302879bd6) [&sect;](https://postgr.es/c/efbebb4e8) [&sect;](https://postgr.es/c/ba97bf9cb)
-  Add function [`pg_clear_extended_stats()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-statsmod) to remove extended statistics (Corey Huinker, Michael Paquier) [&sect;](https://postgr.es/c/d756fa101)
-  Adjust the optimizer to consider startup costs of partial paths (Robert Haas, Tomas Vondra) [&sect;](https://postgr.es/c/8300d3ad4)
-  Allow negative values of [`pg_aggregate`](../../internals/system-catalogs/pg_aggregate.md#catalog-pg-aggregate).[`aggtransspace`](../../internals/system-catalogs/pg_aggregate.md#catalog-pg-aggregate) to indicate unbounded memory usage (Richard Guo) [&sect;](https://postgr.es/c/185e30426)

   This information is used by the optimizer in planning memory usage.
  <a id="release-19-performance"></a>

##### General Performance


-  Improve performance of foreign key constraint checks (Junwang Zhao, Amit Langote, Chao Li) [&sect;](https://postgr.es/c/2da86c1ef) [&sect;](https://postgr.es/c/e484b0eea) [&sect;](https://postgr.es/c/b7b27eb41) [&sect;](https://postgr.es/c/5c54c3ed1)
-  Improve [asynchronous I/O](../../server-administration/server-configuration/resource-consumption.md#guc-io-method) read-ahead scheduling for large requests (Andres Freund) [&sect;](https://postgr.es/c/a9ee66881) [&sect;](https://postgr.es/c/8ca147d58) [&sect;](https://postgr.es/c/f63ca3379)
-  Allow [`io_method`](../../server-administration/server-configuration/resource-consumption.md#guc-io-method) method `worker` to automatically control needed background workers (Thomas Munro) [&sect;](https://postgr.es/c/d1c01b79d)

   The new server variables are [`io_min_workers`](../../server-administration/server-configuration/resource-consumption.md#guc-io-min-workers), [`io_max_workers`](../../server-administration/server-configuration/resource-consumption.md#guc-io-max-workers), [`io_worker_idle_timeout`](../../server-administration/server-configuration/resource-consumption.md#guc-io-worker-idle-timeout), and [`io_worker_launch_interval`](../../server-administration/server-configuration/resource-consumption.md#guc-io-worker-launch-interval).
-  Allow query table scans to mark pages as all-visible in the [visibility map](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-visibility-map) (Melanie Plageman) [&sect;](https://postgr.es/c/b46e1e54d)

   Previously only [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) and [`COPY ... FREEZE`](../../reference/sql-commands/copy.md#sql-copy) could do this.
-  Allow [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) to use [`parallel autovacuum workers`](../../reference/sql-commands/create-table.md#reloption-autovacuum-parallel-workers) (Daniil Davydov) [&sect;](https://postgr.es/c/1ff3180ca) [&sect;](https://postgr.es/c/2a3d2f9f6)

   The maximum number of workers is controlled by server variable [`autovacuum_max_parallel_workers`](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-max-parallel-workers) and per-table storage parameter `autovacuum_parallel_workers`.
-  Allow [`TID`](../../the-sql-language/data-types/object-identifier-types.md#datatype-oid) Range Scans to be parallelized (Cary Huang, David Rowley) [&sect;](https://postgr.es/c/0ca3b1697)
-  Improve [`COPY FROM`](../../reference/sql-commands/copy.md#sql-copy) performance for text and CSV input using SIMD CPU instructions (Nazir Bilal Yavuz, Shinya Kato) [&sect;](https://postgr.es/c/e0a3a3fd5)
-  Improve [`NOTIFY`](../../reference/sql-commands/notify.md#sql-notify) to only wake up backends that are listening to specified notifications (Joel Jacobson) [&sect;](https://postgr.es/c/282b1cde9)

   Previously most backends were woken by [`NOTIFY`](../../reference/sql-commands/notify.md#sql-notify).
-  Change the default [TOAST](../../internals/database-physical-storage/toast.md#storage-toast) compression method from `pglz` to the more efficient `lz4` (Euler Taveira) [&sect;](https://postgr.es/c/34dfca293)

   This is done by changing the default for server variable [`default_toast_compression`](../../server-administration/server-configuration/client-connection-defaults.md#guc-default-toast-compression).
-  Improve performance of internal row deformation (David Rowley) [&sect;](https://postgr.es/c/c456e3911)
-  Improve performance of repeated UTF-8 case-folding operations (Andreas Karlsson) [&sect;](https://postgr.es/c/c4ff35f10)
-  Improve performance of hash index bulk-deletion and [`GIN`](../../internals/built-in-index-access-methods/gin-indexes.md#gin) index vacuuming using streaming reads (Xuneng Zhou) [&sect;](https://postgr.es/c/bfa3c4f10) [&sect;](https://postgr.es/c/6c228755a)
-  Improve sort performance using radix sort (John Naylor) [&sect;](https://postgr.es/c/ef3c3cf6d)
-  Improve timing performance measurements (Lukas Fittl, Andres Freund, David Geier) [&sect;](https://postgr.es/c/294520c44) [&sect;](https://postgr.es/c/16fca4825)

   This benefits [`EXPLAIN (ANALYZE, TIMING)`](../../reference/sql-commands/explain.md#sql-explain) and [pg_test_timing](../../reference/postgresql-server-applications/pg_test_timing.md#pgtesttiming), and is controlled via server variable [`timing_clock_source`](../../server-administration/server-configuration/resource-consumption.md#guc-timing-clock-source).
-  Optimize [plpgsql](../../server-programming/pl-pgsql-sql-procedural-language/index.md#plpgsql) syntax `SELECT simple-expression INTO` *var* (Tom Lane) [&sect;](https://postgr.es/c/ce8d5fe0e)
  <a id="release-19-system-views"></a>

##### System Views


-  Add system view [`pg_stat_lock`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-lock-view) and function [`pg_stat_get_lock()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-lock-view) to report per-lock-type statistics (Bertrand Drouvot) [&sect;](https://postgr.es/c/4019f725f)
-  Add system view [`pg_stat_recovery`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-recovery-view) to report recovery status (Xuneng Zhou, Shinya Kato) [&sect;](https://postgr.es/c/01d485b14) [&sect;](https://postgr.es/c/2d4ead6f4)
-  Add system view [`pg_stat_autovacuum_scores`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-autovacuum-scores-view) to report per-table [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) details (Sami Imseih) [&sect;](https://postgr.es/c/87f61f0c8)
-  Add system view [`pg_dsm_registry_allocations`](../../internals/system-views/pg_dsm_registry_allocations.md#view-pg-dsm-registry-allocations) to report [dynamic shared memory](../../server-administration/server-configuration/resource-consumption.md#guc-dynamic-shared-memory-type) details (Florents Tselai, Nathan Bossart) [&sect;](https://postgr.es/c/167ed8082) [&sect;](https://postgr.es/c/f894acb24)
-  Add vacuum initiation details to system view [`pg_stat_progress_vacuum`](../../server-administration/monitoring-database-activity/progress-reporting.md#pg-stat-progress-vacuum-view) (Shinya Kato) [&sect;](https://postgr.es/c/0d7895206)

   The new `started_by` column reports the initiator of the vacuum, and `mode` indicates its aggressiveness.
-  Add analyze initiation details to system view [`pg_stat_progress_analyze`](../../server-administration/monitoring-database-activity/progress-reporting.md#pg-stat-progress-analyze-view) (Shinya Kato) [&sect;](https://postgr.es/c/ab40db385)

   The new `started_by` column reports the initiator of the analyze.
-  Add `mem_exceeded_count` column to system view [`pg_stat_replication_slots`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-replication-slots-view) (Bertrand Drouvot) [&sect;](https://postgr.es/c/d3b6183dd)

   This reports the number of times that [`logical_decoding_work_mem`](../../server-administration/server-configuration/resource-consumption.md#guc-logical-decoding-work-mem) was exceeded.
-  Add slot synchronization skip information to [`pg_stat_replication_slots`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-replication-slots-view) and [`pg_replication_slots`](../../internals/system-views/pg_replication_slots.md#view-pg-replication-slots) (Shlok Kyal) [&sect;](https://postgr.es/c/76b78721c) [&sect;](https://postgr.es/c/e68b6adad) [&sect;](https://postgr.es/c/5db6a344a)

   The new columns are `slotsync_skip_count`, `slotsync_last_skip`, and `slotsync_skip_reason`.
-  Add `update_deleted` column to system view [`pg_stat_subscription_stats`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-subscription-stats) (Zhijie Hou) [&sect;](https://postgr.es/c/fd5a1a0c3)

   This reports the number of rows where updates were ignored due to concurrent deletes. This requires the subscriber have `retain_dead_tuples` enabled.
-  Add `sync_seq_error_count` column to system view [`pg_stat_subscription_stats`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-subscription-stats) to report sequence synchronization errors (Vignesh C) [&sect;](https://postgr.es/c/f6a4c498d) [&sect;](https://postgr.es/c/3edaf29fa)
-  Add `stats_reset` column to system views [`pg_stat_all_tables`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-all-tables-view), [`pg_stat_all_indexes`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-all-indexes-view), and [`pg_statio_all_sequences`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-statio-all-sequences-view) (Bertrand Drouvot, Sami Imseih, Shihao Zhong) [&sect;](https://postgr.es/c/a5b543258)

   It also appears in the `sys` and `user` view variants.
-  Add `stats_reset` column to system views [`pg_stat_user_functions`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-user-functions-view) and [`pg_stat_database_conflicts`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-database-conflicts-view) (Bertrand Drouvot, Shihao Zhong) [&sect;](https://postgr.es/c/b71bae41a) [&sect;](https://postgr.es/c/8fe315f18)
-  Add `location` column to system views [`pg_available_extensions`](../../internals/system-views/pg_available_extensions.md#view-pg-available-extensions) and [`pg_available_extension_versions`](../../internals/system-views/pg_available_extension_versions.md#view-pg-available-extension-versions) to report the file system directory of extensions (Matheus Alcantara) [&sect;](https://postgr.es/c/f3c9e341c)
-  Add `backup_type` column to system view [`pg_stat_progress_basebackup`](../../server-administration/monitoring-database-activity/progress-reporting.md#pg-stat-progress-basebackup-view) to report the type of backup (Shinya Kato) [&sect;](https://postgr.es/c/deb674454)

   Possible values are `full` or `incremental`.
-  Add `connecting` value to system view column [`pg_stat_wal_receiver`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-wal-receiver-view).[`status`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-wal-receiver-view) (Xuneng Zhou) [&sect;](https://postgr.es/c/a36164e74)
-  Add reporting of the bytes written to WAL for full page images (Shinya Kato) [&sect;](https://postgr.es/c/f9a09aa29)

   This is accessible via system view [`pg_stat_wal`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-wal-view) and function [`pg_stat_get_backend_wal()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-get-backend-wal).
-  Add columns to system views [`pg_stats`](../../internals/system-views/pg_stats.md#view-pg-stats), [`pg_stats_ext`](../../internals/system-views/pg_stats_ext.md#view-pg-stats-ext), and [`pg_stats_ext_exprs`](../../internals/system-views/pg_stats_ext_exprs.md#view-pg-stats-ext-exprs) (Corey Huinker) [&sect;](https://postgr.es/c/3b88e50d6)

   Adds table `OID` and attribute number columns to [`pg_stats`](../../internals/system-views/pg_stats.md#view-pg-stats), and table `OID` and statistics object `OID` columns to the other two.
-  Add information about range type [extended statistics](../../reference/sql-commands/create-statistics.md#sql-createstatistics) to system view [`pg_stats_ext_exprs`](../../internals/system-views/pg_stats_ext_exprs.md#view-pg-stats-ext-exprs) (Corey Huinker, Michael Paquier) [&sect;](https://postgr.es/c/307447e6d)
  <a id="release-19-monitoring"></a>

##### Monitoring


-  Allow [`log_min_messages`](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-min-messages) log levels to be specified by process type (Euler Taveira) [&sect;](https://postgr.es/c/38e0190ce)

   The new format is *type*:*level*. A value without a colon controls all process types, allowing backward compatibility.
-  Add server variable [`log_autoanalyze_min_duration`](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-autoanalyze-min-duration) to log long-running analyze operations by [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) operations (Shinya Kato) [&sect;](https://postgr.es/c/dd3ae3783)

   Server variable [`log_autovacuum_min_duration`](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-autovacuum-min-duration) now only controls logging of vacuum operations by autovacuum.
-  Enable server variable [`log_lock_waits`](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-lock-waits) by default (Laurenz Albe) [&sect;](https://postgr.es/c/2aac62be8)
-  Add server variable [`debug_print_raw_parse`](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-debug-print-raw-parse) to log raw parse trees (Chao Li) [&sect;](https://postgr.es/c/06473f5a3)

   This is also enabled when the server is started with debug level three and higher.
-  Make messages coming from remote servers appear in the server logs in the same format as local server messages (Vignesh C) [&sect;](https://postgr.es/c/112faf137)

   These include replication, [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw), and [dblink](../additional-supplied-modules-and-extensions/dblink-connect-to-other-postgresql-databases.md#dblink) servers.
-  Add reporting of WAL [full-page write](../../server-administration/server-configuration/write-ahead-log.md#guc-full-page-writes) bytes to [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) and [`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze) logging (Shinya Kato) [&sect;](https://postgr.es/c/ad25744f4)
-  Add IO wait events for [`COPY FROM/TO`](../../reference/sql-commands/copy.md#sql-copy) on a pipe, file, or program (Nikolay Samokhvalov) [&sect;](https://postgr.es/c/e05a24c2d)
-  Add wait events for WAL write and flush [LSN](../../server-administration/reliability-and-the-write-ahead-log/wal-internals.md#wal-internals)s (Xuneng Zhou) [&sect;](https://postgr.es/c/7a39f43d8)
-  Have [`pg_get_sequence_data()`](../../the-sql-language/functions-and-operators/sequence-manipulation-functions.md#func-pg-get-sequence-data) return the sequence page [LSN](../../server-administration/reliability-and-the-write-ahead-log/wal-internals.md#wal-internals) (Vignesh C) [&sect;](https://postgr.es/c/b93172ca5)
-  Add function [`pg_get_multixact_stats()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-snapshot) to report [multixact](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-multixact-wraparound) activity (Naga Appani) [&sect;](https://postgr.es/c/97b101776)
-  Issue warnings when the [wraparound](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound) of xid and multi-xids is less than 100 million (Nathan Bossart) [&sect;](https://postgr.es/c/48f11bfa0)

   The previous warning was 40 million. Warnings are issued to clients and in the server log.
  <a id="release-19-server-config"></a>

##### Server Configuration


-  Allow [online enabling](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-checksum) and disabling of [data checksums](../../server-administration/reliability-and-the-write-ahead-log/data-checksums.md#checksums) (Daniel Gustafsson, Magnus Hagander, Tomas Vondra) [&sect;](https://postgr.es/c/f19c0ecca) [&sect;](https://postgr.es/c/b364828f8)

   Previously the checksum status could only be changed while the cluster was offline using [pg_checksums](../../reference/postgresql-server-applications/pg_checksums.md#app-pgchecksums).
-  Add scoring system to control the order that tables are processed by [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) (Nathan Bossart) [&sect;](https://postgr.es/c/d7965d65f)

   The new server variables are [`autovacuum_freeze_score_weight`](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-freeze-score-weight), [`autovacuum_multixact_freeze_score_weight`](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-multixact-freeze-score-weight), [`autovacuum_vacuum_score_weight`](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-vacuum-score-weight), [`autovacuum_vacuum_insert_score_weight`](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-vacuum-insert-score-weight), and [`autovacuum_analyze_score_weight`](../../server-administration/server-configuration/vacuuming.md#guc-autovacuum-analyze-score-weight).
-  Add server-side support for [SNI](../../server-administration/server-setup-and-operation/secure-tcp-ip-connections-with-ssl.md#ssl-sni) (Server Name Indication) (Daniel Gustafsson, Jacob Champion) [&sect;](https://postgr.es/c/4f433025f)

   New configuration file [`PGDATA/pg_hosts.conf`](../../server-administration/server-configuration/file-locations.md#guc-hosts-file) specifies hostname/key pairs.
-  Add a new OAUTH flow hook [`PQAUTHDATA_OAUTH_BEARER_TOKEN_V2`](../../client-interfaces/libpq-c-library/oauth-support.md#libpq-oauth-authdata-oauth-bearer-token-v2) (Jacob Champion) [&sect;](https://postgr.es/c/e982331b5) [&sect;](https://postgr.es/c/0af4d402c)

   This is an improved version of [`PQAUTHDATA_OAUTH_BEARER_TOKEN`](../../client-interfaces/libpq-c-library/oauth-support.md#libpq-oauth-authdata-oauth-bearer-token) by adding the issuer identifier and error message specification.
-  Allow roles [`pg_read_all_data`](../../server-administration/database-roles/predefined-roles.md#predefined-role-pg-read-all-data) and [`pg_write_all_data`](../../server-administration/database-roles/predefined-roles.md#predefined-role-pg-read-all-data) to read/write large objects (Nitin Motiani, Nathan Bossart) [&sect;](https://postgr.es/c/d98197602)

   These roles are designed to allow non-super users to run [pg_dump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump).
-  Allow background workers to be configured to terminate before database-level operations (Aya Iwata) [&sect;](https://postgr.es/c/f1e251be8)

   This allows database-level operations to complete more quickly since blocking background workers can now be terminated.
-  Allow [server variables](../../server-administration/server-configuration/setting-parameters.md#config-setting) that represent lists to be emptied by setting the value to `NULL` (Tom Lane) [&sect;](https://postgr.es/c/ff4597acd)
-  Update [GB18030 encoding](../../server-administration/localization/character-set-support.md#multibyte-charset-supported) from version 2000 to 2022 (Chao Li, Zheng Tao) [&sect;](https://postgr.es/c/5334620ee)

   See the commit message for compatibility details.
  <a id="release-19-replication"></a>

##### [Streaming Replication and Recovery]


-  Add [`WAIT FOR`](../../reference/sql-commands/wait-for.md#sql-wait-for) command to allow standbys to wait for [LSN](../../server-administration/reliability-and-the-write-ahead-log/wal-internals.md#wal-internals) values to be written, flushed, or replayed (Kartyshov Ivan, Alexander Korotkov, Xuneng Zhou) [&sect;](https://postgr.es/c/447aae13b) [&sect;](https://postgr.es/c/49a181b5d)
-  Improve function [`pg_sync_replication_slots()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-sync-replication-slots) to wait for replication synchronization completion (Ajin Cherian, Zhijie Hou) [&sect;](https://postgr.es/c/0d2d4a0ec)

   Previously, certain synchronization failures would not be reported.
-  Add server variable [`wal_sender_shutdown_timeout`](../../server-administration/server-configuration/replication.md#guc-wal-sender-shutdown-timeout) to limit replica synchronization waits during shutdown (Andrey Silitskiy, Hayato Kuroda) [&sect;](https://postgr.es/c/a8f45dee9)

   By default, senders still wait forever for synchronization.
-  Allow [`wal_receiver_timeout`](../../server-administration/server-configuration/replication.md#guc-wal-receiver-timeout) to be set per-subscription and user (Fujii Masao) [&sect;](https://postgr.es/c/8a6af3ad0) [&sect;](https://postgr.es/c/fb80f388f)

   This allows subscribers to use different [`wal_receiver_timeout`](../../server-administration/server-configuration/replication.md#guc-wal-receiver-timeout) values.
-  Add optional pid parameter to [`pg_replication_origin_session_setup()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-replication-origin-session-setup) to allow parallelization of SQL-level replication solutions (Doruk Yilmaz, Hayato Kuroda) [&sect;](https://postgr.es/c/5b148706c)
  <a id="release-19-logical"></a>

##### [Logical Replication]


-  Allow sequence values stored in subscribers to match the publisher (Vignesh C) [&sect;](https://postgr.es/c/f0b3573c3) [&sect;](https://postgr.es/c/5509055d6) [&sect;](https://postgr.es/c/55cefadde)

   This is enabled during [`CREATE SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription), [`ALTER SUBSCRIPTION ... REFRESH PUBLICATION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription), and [`ALTER SUBSCRIPTION ... REFRESH SEQUENCES`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription). The latter only updates values, not sequence existence. Function [`pg_get_sequence_data()`](../../the-sql-language/functions-and-operators/sequence-manipulation-functions.md#func-pg-get-sequence-data) allows inspection of sequence synchronization.
-  Allow [`CREATE`](../../reference/sql-commands/create-publication.md#sql-createpublication)/[`ALTER PUBLICATION`](../../reference/sql-commands/alter-publication.md#sql-alterpublication) to publish all sequences (Vignesh C, Tomas Vondra) [&sect;](https://postgr.es/c/96b378497)

   This is enabled with the `ALL SEQUENCES` clause.
-  Allow [`ALTER SUBSCRIPTION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription) on publications to synchronize the existence of sequences on subscribers to match the publisher (Vignesh C) [&sect;](https://postgr.es/c/f0b3573c3)

   This is enabled with the `REFRESH SEQUENCES` clause.
-  Allow [`CREATE`](../../reference/sql-commands/create-publication.md#sql-createpublication)/[`ALTER PUBLICATION`](../../reference/sql-commands/alter-publication.md#sql-alterpublication) to exclude some tables (Vignesh C, Shlok Kyal) [&sect;](https://postgr.es/c/493f8c643) [&sect;](https://postgr.es/c/6b0550c45) [&sect;](https://postgr.es/c/fd366065e) [&sect;](https://postgr.es/c/5984ea868)

   This is controlled with the `EXCEPT` clause, and is useful when specifying `ALL TABLES`.
-  Add [`CREATE`](../../reference/sql-commands/create-publication.md#sql-createpublication)/[`ALTER PUBLICATION`](../../reference/sql-commands/alter-publication.md#sql-alterpublication) setting [`retain_dead_tuples`](../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-retain-dead-tuples) to retain information needed for conflict resolution (Zhijie Hou) [&sect;](https://postgr.es/c/228c37086) [&sect;](https://postgr.es/c/0d48d393d) [&sect;](https://postgr.es/c/a850be2fe)

   Also add setting [`max_retention_duration`](../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-max-retention-duration) to limit `retain_dead_tuples` retention.
-  Allow [`CREATE SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription) to use [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) foreign data wrapper connection parameters (Jeff Davis) [&sect;](https://postgr.es/c/8185bb534)

   The connection parameters are referenced via [`CREATE SUBSCRIPTION ... SERVER`](../../reference/sql-commands/create-subscription.md#sql-createsubscription).
-  When server variable [`wal_level`](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-level) is `replica`, allow automatic enablement of logical replication when needed (Masahiko Sawada) [&sect;](https://postgr.es/c/67c20979c)

   New server variable [`effective_wal_level`](../../server-administration/server-configuration/preset-options.md#guc-effective-wal-level) reports the effective WAL level.
   <a id="release-19-query"></a>

#### Query Commands


-  Add support for SQL Property Graph Queries (SQL/PGQ) (Peter Eisentraut, Ashutosh Bapat) [&sect;](https://postgr.es/c/2f094e7ac) [&sect;](https://postgr.es/c/c5b3253b8) [&sect;](https://postgr.es/c/a0dd0702e)

   Internally these are processed like views so are written as standard relational queries.
-  Add `FOR PORTION OF` clause to [`UPDATE`](../../reference/sql-commands/update.md#sql-update) and [`DELETE`](../../reference/sql-commands/delete.md#sql-delete) (Paul A. Jungwirth) [&sect;](https://postgr.es/c/8e72d914c) [&sect;](https://postgr.es/c/b6ccd30d8)

   This allows operations on temporal ranges.
-  Add `GROUP BY ALL` syntax to [`SELECT`](../../reference/sql-commands/select.md#sql-select) to automatically group all non-aggregate and non-window-function target list parameters (David Christensen) [&sect;](https://postgr.es/c/ef38a4d97)
-  Allow `GROUP BY` to process target list subqueries that have expressions referencing non-subquery columns (Tom Lane) [&sect;](https://postgr.es/c/415100aa6)

   Also fix a bug in how [`GROUPING()`](../../the-sql-language/functions-and-operators/aggregate-functions.md#functions-grouping-table) handles target list subquery aliases.
-  Allow [window functions](../../the-sql-language/functions-and-operators/window-functions.md#functions-window) to ignore NULLs with the `IGNORE NULLS`/`RESPECT NULLS` clause (Oliver Ford, Tatsuo Ishii) [&sect;](https://postgr.es/c/25a30bbd4)

   Supported window functions are `lead()`, `lag()`, `first_value()`, `last_value()`, and `nth_value()`.
-  Add support for [`INSERT ... ON CONFLICT DO SELECT ... RETURNING`](../../reference/sql-commands/insert.md#sql-insert) (Andreas Karlsson, Marko Tiikkaja, Viktor Holmberg) [&sect;](https://postgr.es/c/88327092f)

   This allows conflicting rows to be returned, and optionally locked with `FOR UPDATE`/`SHARE`.
  <a id="release-19-utility"></a>

#### Utility Commands


-  Add [`REPACK`](../../reference/sql-commands/repack.md#sql-repack) command which replaces [`VACUUM FULL`](../../reference/sql-commands/vacuum.md#sql-vacuum) and [`CLUSTER`](../../reference/sql-commands/cluster.md#sql-cluster) (Antonin Houska) [&sect;](https://postgr.es/c/ac58465e0)

   The two former commands did similar things, but with confusing names, so unify them as [`REPACK`](../../reference/sql-commands/repack.md#sql-repack). The old commands have been retained for compatibility.
-  Allow [`REPACK`](../../reference/sql-commands/repack.md#sql-repack) to rebuild tables without access-exclusive locking (Antonin Houska, Mihail Nikalayeu, Álvaro Herrera) [&sect;](https://postgr.es/c/28d534e2a) [&sect;](https://postgr.es/c/8fb95a8ab) [&sect;](https://postgr.es/c/e76d8c749)

   This is enabled via the `CONCURRENTLY` option. Server variable [`max_repack_replication_slots`](../../server-administration/server-configuration/replication.md#guc-max-repack-replication-slots) was also added.
-  Allow partitions to be merged and split using [`ALTER TABLE ... MERGE/SPLIT PARTITIONS`](../../reference/sql-commands/alter-table.md#sql-altertable) (Dmitry Koval, Alexander Korotkov, Tender Wang, Richard Guo, Dagfinn Ilmari Mannsåker, Fujii Masao, Jian He) [&sect;](https://postgr.es/c/f2e4cc427) [&sect;](https://postgr.es/c/4b3d17362)
-  Allow [`GRANT`](../../reference/sql-commands/grant.md#sql-grant)/[`REVOKE`](../../reference/sql-commands/revoke.md#sql-revoke) to specify the effective role performing the privileges adjustment (Nathan Bossart, Tom Lane) [&sect;](https://postgr.es/c/dd1398f13)

   The `GRANTED BY` clause controls this.
-  Allow [`CREATE SCHEMA`](../../reference/sql-commands/create-schema.md#sql-createschema) to create more types of objects in newly-created schemas (Kirill Reshke, Jian He, Tom Lane) [&sect;](https://postgr.es/c/d51697484)
-  Allow [`CHECKPOINT`](../../reference/sql-commands/checkpoint.md#sql-checkpoint) to accept a list of options (Christoph Berg) [&sect;](https://postgr.es/c/a4f126516) [&sect;](https://postgr.es/c/2f698d7f4) [&sect;](https://postgr.es/c/8d33fbacb)

   Supported options are `MODE` and `FLUSH_UNLOGGED`.
-  Add `CONNECTION` clause to [`CREATE FOREIGN DATA WRAPPER`](../../reference/sql-commands/create-foreign-data-wrapper.md#sql-createforeigndatawrapper) to specify a function to be called for subscription connection parameters (Jeff Davis, Noriyoshi Shinoda) [&sect;](https://postgr.es/c/8185bb534) [&sect;](https://postgr.es/c/90630ec42)
-  Add memory usage and parallelism reporting to [`VACUUM (VERBOSE)`](../../reference/sql-commands/vacuum.md#sql-vacuum) and [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) logs (Tatsuya Kawata, Daniil Davydov) [&sect;](https://postgr.es/c/736f754ee) [&sect;](https://postgr.es/c/adcdbe938)
 <a id="release-19-constraints"></a>

##### [Constraints]


-  Allow [`ALTER TABLE ALTER CONSTRAINT ... [NOT] ENFORCED`](../../reference/sql-commands/alter-table.md#sql-altertable) for `CHECK` constraints (Jian He) [&sect;](https://postgr.es/c/342051d73)

   Previously enforcement changes were only supported for foreign key constraints.
-  Allow [`ALTER TABLE ... COLUMN SET EXPRESSION`](../../reference/sql-commands/alter-table.md#sql-altertable) to succeed on virtual columns with `CHECK` constraints (Jian He) [&sect;](https://postgr.es/c/f80bedd52)

   This was previously prohibited.
  <a id="release-19-copy"></a>

##### [sql-copy]


-  Allow multiple headers lines to be skipped by [`COPY FROM`](../../reference/sql-commands/copy.md#sql-copy) (Shinya Kato, Fujii Masao) [&sect;](https://postgr.es/c/bc2f348e8)

   Previously only a single header line could be skipped.
-  Allow [`COPY FROM`](../../reference/sql-commands/copy.md#sql-copy) to set invalid input values to `NULL` (Jian He, Kirill Reshke) [&sect;](https://postgr.es/c/2a525cc97)

   This is done using the [`COPY`](../../reference/sql-commands/copy.md#sql-copy) option `ON_ERROR SET_NULL`.
-  Allow [`COPY TO`](../../reference/sql-commands/copy.md#sql-copy) to output JSON format (Joe Conway, Jian He, Andrew Dunstan) [&sect;](https://postgr.es/c/7dadd38cd) [&sect;](https://postgr.es/c/4c0390ac5)

   JSON output can also be a single JSON array using the [`COPY`](../../reference/sql-commands/copy.md#sql-copy) option `FORCE_ARRAY`.
-  Allow [`COPY TO`](../../reference/sql-commands/copy.md#sql-copy) to process partitioned tables (Jian He, Ajin Cherian) [&sect;](https://postgr.es/c/4bea91f21) [&sect;](https://postgr.es/c/266543a62)

   Previously [`COPY (SELECT ...)`](../../reference/sql-commands/copy.md#sql-copy) had to be used to output partitioned tables. This also improves logical replication table synchronization.
  <a id="release-19-explain"></a>

##### [sql-explain]


-  Add [`EXPLAIN ANALYZE`](../../reference/sql-commands/explain.md#sql-explain) option `IO` to report asynchronous IO activity (Tomas Vondra) [&sect;](https://postgr.es/c/681daed93) [&sect;](https://postgr.es/c/3b1117d6e) [&sect;](https://postgr.es/c/e157fe6f7)
-  Add reporting of WAL [full-page write](../../server-administration/server-configuration/write-ahead-log.md#guc-full-page-writes) bytes to [`EXPLAIN (ANALYZE, WAL)`](../../reference/sql-commands/explain.md#sql-explain) output (Shinya Kato) [&sect;](https://postgr.es/c/5ab0b6a24)
-  Add Memoize cache and lookup estimates to [`EXPLAIN`](../../reference/sql-commands/explain.md#sql-explain) output (Ilia Evdokimov, Lukas Fittl) [&sect;](https://postgr.es/c/4bc62b868)

   This can show why Memoize was chosen.
   <a id="release-19-datatypes"></a>

#### Data Types


-  Add the 64-bit unsigned data type [`oid8`](../../the-sql-language/data-types/object-identifier-types.md#datatype-oid) (Michael Paquier) [&sect;](https://postgr.es/c/b139bd3b6)
-  Add more [`jsonpath`](../../the-sql-language/data-types/json-types.md#datatype-jsonpath) string methods (Florents Tselai, David E. Wheeler) [&sect;](https://postgr.es/c/bd4f879a9)

   They are [`ltrim()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), [`rtrim()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), [`btrim()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), [`lower()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), [`upper()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), [`initcap()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), [`replace()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators), and [`split_part()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-path-operators). These are immutable like their non-JSON string variants.
-  Allow casts between [`bytea`](../../the-sql-language/data-types/binary-data-types.md#datatype-binary) and [`uuid`](../../the-sql-language/data-types/uuid-type.md#datatype-uuid) data types (Dagfinn Ilmari Mannsåker, Aleksander Alekseev) [&sect;](https://postgr.es/c/ba21f5bf8)
-  Add ability to cast between database names and [`oid8`](../../the-sql-language/data-types/object-identifier-types.md#datatype-oid)s using [`regdatabase`](../../the-sql-language/data-types/object-identifier-types.md#datatype-oid) (Ian Lawrence Barwick) [&sect;](https://postgr.es/c/bd09f024a)
-  Add functions [`tid_block()`](../../the-sql-language/functions-and-operators/tid-functions.md#functions-tid) and [`tid_offset()`](../../the-sql-language/functions-and-operators/tid-functions.md#functions-tid) to extract block numbers and offsets from [`tid`](../../the-sql-language/data-types/object-identifier-types.md#datatype-oid) values (Ayush Tiwari) [&sect;](https://postgr.es/c/df6949ccf)
  <a id="release-19-functions"></a>

#### Functions


-  Add [`date`](../../the-sql-language/data-types/date-time-types.md#datatype-datetime), [`timestamp`](../../the-sql-language/data-types/date-time-types.md#datatype-datetime), and [`timestamptz`](../../the-sql-language/data-types/date-time-types.md#datatype-datetime) versions of [`random(min, max)`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-random-table) (Damien Clochard, Dean Rasheed) [&sect;](https://postgr.es/c/faf071b55) [&sect;](https://postgr.es/c/9c24111c4)
-  Allow [`encode()`](../../the-sql-language/functions-and-operators/binary-string-functions-and-operators.md#functions-binarystring-conversions) and [`decode()`](../../the-sql-language/functions-and-operators/binary-string-functions-and-operators.md#functions-binarystring-conversions) to process data in `base64url` and `base32hex` formats (Andrey Borodin, Aleksander Alekseev, Florents Tselai) [&sect;](https://postgr.es/c/497c1170c) [&sect;](https://postgr.es/c/e752a2ccc) [&sect;](https://postgr.es/c/e1d917182)

   This format retains ordering, unlike `base32`.
-  Add functions to return a set of ranges resulting from range subtraction (Paul A. Jungwirth) [&sect;](https://postgr.es/c/5eed8ce50)

   The functions are [`range_minus_multi()`](../../the-sql-language/functions-and-operators/range-multirange-functions-and-operators.md#range-functions-table) and [`multirange_minus_multi()`](../../the-sql-language/functions-and-operators/range-multirange-functions-and-operators.md#multirange-functions-table). This is useful to represent range subtraction results with gaps.
-  Add function [`error_on_null()`](../../the-sql-language/functions-and-operators/comparison-functions-and-operators.md#functions-comparison-func-table) to return the supplied parameter, or error on `NULL` input (Joel Jacobson) [&sect;](https://postgr.es/c/2b75c38b7)
-  Allow [`IS JSON`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-misc) to work on domains defined over supported base types (Jian He) [&sect;](https://postgr.es/c/3b4c2b9db)

   The supported base types are [`TEXT`](../../the-sql-language/data-types/character-types.md#datatype-character), [`JSON`](../../the-sql-language/data-types/json-types.md#datatype-json), [`JSONB`](../../the-sql-language/data-types/json-types.md#datatype-json), and [`BYTEA`](../../the-sql-language/data-types/binary-data-types.md#datatype-binary).
-  Add [full text stemmers](../../the-sql-language/full-text-search/dictionaries.md#textsearch-snowball-dictionary) for Polish and Esperanto (Tom Lane) [&sect;](https://postgr.es/c/7dc95cc3b)

   The Dutch stemmer has also been updated. The old Dutch stemmer is available via [`dutch_porter`](../../the-sql-language/full-text-search/dictionaries.md#textsearch-snowball-dictionary).
-  Add function [`pg_get_role_ddl()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-get-object-ddl-table) to output role creation commands (Mario Gonzalez, Bryan Green, Andrew Dunstan, Euler Taveira) [&sect;](https://postgr.es/c/76e514ebb)
-  Add function [`pg_get_tablespace_ddl()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-get-object-ddl-table) to output tablespace creation commands (Nishant Sharma, Manni Wood, Andrew Dunstan, Euler Taveira) [&sect;](https://postgr.es/c/b99fd9fd7)
-  Add function [`pg_get_database_ddl()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-get-object-ddl-table) to output database creation commands (Akshay Joshi, Andrew Dunstan, Euler Taveira) [&sect;](https://postgr.es/c/a4f774cf1)
-  Allow event triggers to be written using [PL/Python](../../server-programming/pl-python-python-procedural-language/index.md#plpython) (Euler Taveira, Dimitri Fontaine) [&sect;](https://postgr.es/c/53eff471c)
  <a id="release-19-libpq"></a>

#### [Libpq]


-  Allow libpq connections to specify a service file via [`servicefile`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-servicefile) (Torsten Förtsch, Ryo Kanbayashi) [&sect;](https://postgr.es/c/092f3c63e)
-  Add special libpq protocol version 3.9999 for version testing (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/d8d7c5dc8)
-  Add libpq function [`PQgetThreadLock()`](../../client-interfaces/libpq-c-library/behavior-in-threaded-programs.md#libpq-threading) to retrieve the current locking callback (Jacob Champion) [&sect;](https://postgr.es/c/b8d768583)
-  Add libpq connection parameter [`oauth_ca_file`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-oauth-ca-file) to specify the OAUTH certificate authority file (Jonathan Gonzalez V., Jacob Champion) [&sect;](https://postgr.es/c/993368113)

   This can also be set via the [`PGOAUTHCAFILE`](../../client-interfaces/libpq-c-library/environment-variables.md#libpq-envars) environment variable. The default is to use curl's built-in certificates.
-  Allow custom OAUTH validators to register custom [`pg_hba.conf`](../../server-administration/client-authentication/the-pg_hba-conf-file.md#auth-pg-hba-conf) authentication options (Jacob Champion) [&sect;](https://postgr.es/c/b977bd308)
-  Allow OAUTH validators to supply failure details (Jacob Champion) [&sect;](https://postgr.es/c/d438a3659)

   This is done by setting the [`ValidatorModuleResult`](../../server-programming/oauth-validator-modules/oauth-validator-callbacks.md#oauth-validator-callback-validate) structure member error_detail.
-  Allow libpq environment variable [`PGOAUTHDEBUG`](../../client-interfaces/libpq-c-library/environment-variables.md#libpq-envars) to specify particular debug options (Zsolt Parragi, Jacob Champion) [&sect;](https://postgr.es/c/6d00fb904)

   The `UNSAFE` option still generates all debugging output.
  <a id="release-19-psql"></a>

#### [app-psql]


-  Allow the search path to appear in the [psql](../../reference/postgresql-client-applications/psql.md#app-psql) prompt via [`%S`](../../reference/postgresql-client-applications/psql.md#app-psql-prompting-S) (Florents Tselai) [&sect;](https://postgr.es/c/b3ce55f41)

   This works when [psql](../../reference/postgresql-client-applications/psql.md#app-psql) is connected to PostgreSQL 18 or later.
-  Allow the hot standby status to appear in the [psql](../../reference/postgresql-client-applications/psql.md#app-psql) prompt via [`%i`](../../reference/postgresql-client-applications/psql.md#app-psql-prompting-i) (Jim Jones) [&sect;](https://postgr.es/c/dddbbc253)
-  Modify [psql](../../reference/postgresql-client-applications/psql.md#app-psql) backslash commands to show comments for publications, subscriptions, and extended statistics (Fujii Masao, Jim Jones) [&sect;](https://postgr.es/c/aecc55866)

   The modified commands are [`\dRp+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-drp), [`\dRs+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-drs), and [`\dX+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-dx-uc).
-  Allow control over how booleans are displayed in [psql](../../reference/postgresql-client-applications/psql.md#app-psql) (David G. Johnston) [&sect;](https://postgr.es/c/645cb44c5)

   The \pset variables are [`display_true`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-pset-display-true) and [`display_false`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-pset-display-false).
-  Add [psql](../../reference/postgresql-client-applications/psql.md#app-psql) variable [`SERVICEFILE`](../../reference/postgresql-client-applications/psql.md#app-psql-variables-servicefile) to reference the service file location (Ryo Kanbayashi) [&sect;](https://postgr.es/c/6b1c4d326)
-  Allow [psql](../../reference/postgresql-client-applications/psql.md#app-psql) to more accurately determine if the [pager](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-pset-pager) is needed (Erik Wienhold) [&sect;](https://postgr.es/c/27da1a796)
-  Add or improve [psql](../../reference/postgresql-client-applications/psql.md#app-psql) tab completion (Yamaguchi Atsuo, Yugo Nagata, Haruna Miwa, Xuneng Zhou, Dagfinn Ilmari Mannsåker, Fujii Masao, Álvaro Herrera, Jian He, Tatsuya Kawata, Ian Lawrence Barwick, Vasuki M) [&sect;](https://postgr.es/c/5fa7837d9) [&sect;](https://postgr.es/c/c6a7d3bab) [&sect;](https://postgr.es/c/81966c545) [&sect;](https://postgr.es/c/a1f7f91be) [&sect;](https://postgr.es/c/6d2ff1de4) [&sect;](https://postgr.es/c/02fd47dbf) [&sect;](https://postgr.es/c/14ee8e640) [&sect;](https://postgr.es/c/ff0bcb248) [&sect;](https://postgr.es/c/86c539c5a) [&sect;](https://postgr.es/c/a604affad) [&sect;](https://postgr.es/c/a4c10de92) [&sect;](https://postgr.es/c/28c4b8a05) [&sect;](https://postgr.es/c/0bf7d4ca9) [&sect;](https://postgr.es/c/344b572e3)
  <a id="release-19-server-apps"></a>

#### Server Applications


-  Change [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb)'s `--analyze-only` and `--analyze-in-stages` options to analyze partitioned tables when no targets are specified (Laurenz Albe, Mircea Cadariu, Chao Li) [&sect;](https://postgr.es/c/6429e5b77) [&sect;](https://postgr.es/c/95b6ec52e)

   Previously it skipped partitioned tables. This now matches the behavior of [`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze).
-  Allow [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) to report its commands without running them using option `--dry-run` (Corey Huinker) [&sect;](https://postgr.es/c/d107176d2)
-  Allow [pg_verifybackup](../../reference/postgresql-client-applications/pg_verifybackup.md#app-pgverifybackup) to read WAL files stored in tar archives (Amul Sul) [&sect;](https://postgr.es/c/b3cf461b3)

   Add option `--wal-path` as an alias for the existing and deprecated `--wal-directory` option.
-  Allow [pg_waldump](../../reference/postgresql-server-applications/pg_waldump.md#pgwaldump) to read WAL files stored in tar archives (Amul Sul) [&sect;](https://postgr.es/c/b15c15139)
-  Improve performance of [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) copying large object metadata (Nathan Bossart) [&sect;](https://postgr.es/c/3bcfcd815) [&sect;](https://postgr.es/c/158408fef) [&sect;](https://postgr.es/c/161a3e8b6) [&sect;](https://postgr.es/c/b33f75361)

   Various methods are used, depending on the PostgreSQL version of the old cluster.
-  Allow [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) to process non-default tablespaces stored in the `PGDATA` directory (Nathan Bossart) [&sect;](https://postgr.es/c/412036c22)

   Previously such tablespaces generated an error.
-  Add [pgbench](../../reference/postgresql-client-applications/pgbench.md#pgbench) option `--continue-on-error` to continue after SQL errors (Rintaro Ikeda, Yugo Nagata, Fujii Masao) [&sect;](https://postgr.es/c/0ab208fa5)
-  Improve the usability of [pg_test_timing](../../reference/postgresql-server-applications/pg_test_timing.md#pgtesttiming) (Hannu Krosing, Tom Lane) [&sect;](https://postgr.es/c/0b096e379) [&sect;](https://postgr.es/c/9dcc76414)

   Report nanoseconds instead of microseconds. In addition to histogram output, output a second table that reports exact timings, with an optional cutoff set by `--cutoff`.
 <a id="release-19-pgdump"></a>

##### [pg_dump]/[pg_dumpall]/[pg_restore]


-  Allow [pg_dump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump) to include restorable extended statistics (Corey Huinker) [&sect;](https://postgr.es/c/c32fb29e9)
  <a id="release-19-pg-createsubscriber"></a>

##### [pg_createsubscriber]


-  Allow [pg_createsubscriber](../../reference/postgresql-server-applications/pg_createsubscriber.md#app-pgcreatesubscriber) to ignore specified publications that already exist (Shubham Khanna) [&sect;](https://postgr.es/c/85ddcc2f4)

   Previously this generated an error.
-  Change the way [pg_createsubscriber](../../reference/postgresql-server-applications/pg_createsubscriber.md#app-pgcreatesubscriber) stores recovery parameters (Alyona Vinter) [&sect;](https://postgr.es/c/639352d90)

   Changes are stored in optionally-included `pg_createsubscriber.conf` rather than directly in `postgresql.auto.conf`.
-  Add [pg_createsubscriber](../../reference/postgresql-server-applications/pg_createsubscriber.md#app-pgcreatesubscriber) option `-l`/`--logdir` to redirect output to files (Gyan Sreejith, Hayato Kuroda) [&sect;](https://postgr.es/c/6b5b7eae3)
   <a id="release-19-source-code"></a>

#### Source Code


-  Restore support for [`AIX`](../../server-administration/installation-from-source-code/platform-specific-notes.md#installation-notes-aix) (Aditya Kamath, Srirama Kucherlapati, Peter Eisentraut) [&sect;](https://postgr.es/c/ecae09725) [&sect;](https://postgr.es/c/4a1b05caa)

   This uses gcc and only supports 64-bit builds.
-  Change [`Solaris`](../../server-administration/installation-from-source-code/platform-specific-notes.md#installation-notes-solaris) to use unnamed POSIX semaphores (Tom Lane) [&sect;](https://postgr.es/c/0123ce131)

   Previously it used System V semaphores.
-  Require [Visual Studio 2019](../../server-administration/installation-from-source-code/platform-specific-notes.md#installation-notes-visual-studio) or later (Peter Eisentraut) [&sect;](https://postgr.es/c/8fd9bb1d9)
-  Allow [MSVC](../../server-administration/installation-from-source-code/platform-specific-notes.md#installation-notes-visual-studio) to create [PL/Python](../../server-programming/pl-python-python-procedural-language/index.md#plpython) using the Python Limited API (Bryan Green) [&sect;](https://postgr.es/c/2bc60f862)
-  Allow building on AArch64 using [MSVC](../../server-administration/installation-from-source-code/platform-specific-notes.md#installation-notes-visual-studio) (Niyas Sait, Greg Burd, Dave Cramer) [&sect;](https://postgr.es/c/a516b3f00)
-  Allow execution stack backtraces on [Windows](../../server-administration/installation-from-source-code/platform-specific-notes.md#installation-notes-visual-studio) using `DbgHelp` (Bryan Green) [&sect;](https://postgr.es/c/65707ed9a)
-  Change the supported C language version to C11 (Peter Eisentraut) [&sect;](https://postgr.es/c/f5e0186f8) [&sect;](https://postgr.es/c/4fbe01514)

   Previously C99 was used.
-  Use standard C23 and C++ attributes if available (Peter Eisentraut) [&sect;](https://postgr.es/c/76f4b92ba)
-  Use AVX2 CPU instructions for calculating [page checksums](../../server-administration/reliability-and-the-write-ahead-log/data-checksums.md#checksums) (Matthew Sterrett, Andrew Kim) [&sect;](https://postgr.es/c/5e13b0f24)
-  Use ARM Crypto Extension to Compute CRC32C (John Naylor) [&sect;](https://postgr.es/c/fbc57f2bc)
-  Change [`hex_encode()`](../../the-sql-language/functions-and-operators/binary-string-functions-and-operators.md#functions-binarystring) and [`hex_decode()`](../../the-sql-language/functions-and-operators/binary-string-functions-and-operators.md#functions-binarystring) to use SIMD CPU instructions (Nathan Bossart, Chiranmoy Bhattacharya) [&sect;](https://postgr.es/c/ec8719ccb)
-  Require [Meson](../../server-administration/installation-from-source-code/building-and-installation-with-meson.md#install-meson) version 0.57.2 or later (Peter Eisentraut) [&sect;](https://postgr.es/c/f039c2244)
-  Add Meson option to build both shared and static libraries, or only shared (Peter Eisentraut) [&sect;](https://postgr.es/c/78727dcba)
-  Update [Unicode](../../server-administration/localization/collation-support.md#collation-managing-standard) data to version 17.0.0 (Peter Eisentraut) [&sect;](https://postgr.es/c/57ee39795)
-  Add hooks `planner_setup_hook`, `planner_shutdown_hook`, `joinrel_setup_hook`, and `join_path_setup_hook` (Robert Haas) [&sect;](https://postgr.es/c/94f3ad396) [&sect;](https://postgr.es/c/4020b370f)
-  Allow extensions to replace [set-returning functions](../../server-programming/extending-sql/query-language-sql-functions.md#xfunc-sql-functions-returning-set) in the `FROM` clause with SQL queries (Paul A. Jungwirth) [&sect;](https://postgr.es/c/b140c8d7a)
-  Make [multixid members](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-snapshot) 64-bit (Maxim Orlov) [&sect;](https://postgr.es/c/bd8d9c9bd)
-  Change function prototypes to use `uint*` instead of `bit*` typedefs (Nathan Bossart) [&sect;](https://postgr.es/c/bab2f27ea)
-  Allow [logical decoding plugins](../../server-programming/logical-decoding/logical-decoding-output-plugins.md#logicaldecoding-pgoutput) to specify if they do not access shared catalogs (Antonin Houska) [&sect;](https://postgr.es/c/0d3dba38c)
-  Add simplified and improved shared memory registration function [`ShmemRequestStruct()`](../../server-programming/extending-sql/c-language-functions.md#xfunc-shared-addin-at-startup) (Heikki Linnakangas, Ashutosh Bapat) [&sect;](https://postgr.es/c/283e823f9)

   Functions [`ShmemInitStruct()`](../../server-programming/extending-sql/c-language-functions.md#xfunc-shared-addin-at-startup) and [`ShmemInitHash()`](../../server-programming/extending-sql/c-language-functions.md#xfunc-shared-addin-at-startup) remain for backward compatibility.
-  Add server variable [`debug_exec_backend`](../../server-administration/server-configuration/preset-options.md#guc-debug-exec-backend) to report how parameters are passed to new backends (Daniel Gustafsson) [&sect;](https://postgr.es/c/b3fe098d3)
-  Add documentation section about [temporal tables](../../the-sql-language/data-definition/temporal-tables.md#ddl-temporal-tables) (Paul A. Jungwirth) [&sect;](https://postgr.es/c/e4d8a2af0)
-  Document the environment variables that control the [regression tests](../../server-administration/regression-tests/running-the-tests.md#regress-run-path-substitution) (Michael Paquier) [&sect;](https://postgr.es/c/02976b0a1)
-  Update documented [systemd example](../../server-administration/server-setup-and-operation/starting-the-database-server.md#server-start) to include a restart setting (Andrew Jackson) [&sect;](https://postgr.es/c/b30656ce0)
  <a id="release-19-modules"></a>

#### Additional Modules


-  Add [pg_plan_advice](../additional-supplied-modules-and-extensions/pg_plan_advice-help-the-planner-get-the-right-plan.md#pgplanadvice) module to stabilize and control planner decisions (Robert Haas) [&sect;](https://postgr.es/c/5883ff30b) [&sect;](https://postgr.es/c/6455e55b0)
-  Add extension [pg_stash_advice](../additional-supplied-modules-and-extensions/pg_stash_advice-store-and-automatically-apply-plan-advice.md#pgstashadvice) to allow per-query-id advice to be specified (Robert Haas, Lukas Fittl) [&sect;](https://postgr.es/c/e8ec19aa3) [&sect;](https://postgr.es/c/c10edb102)
-  Show sizes of [`FETCH`](../../reference/sql-commands/fetch.md#sql-fetch) queries as constants in [pg_stat_statements](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements) (Sami Imseih) [&sect;](https://postgr.es/c/bee23ea4d)

   Fetches of different sizes will now be grouped together in [pg_stat_statements](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements) output.
-  Add [generic and custom plan](../../reference/sql-commands/prepare.md#sql-prepare) counts to [pg_stat_statements](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements) (Sami Imseih) [&sect;](https://postgr.es/c/3357471cf)
-  Refactor [pg_buffercache](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache) reporting of shared memory mapping (Bertrand Drouvot) [&sect;](https://postgr.es/c/4b203d499)

   New function `pg_buffercache_os_pages()` and system view `pg_buffercache_os_pages` allow reporting of shared memory mapping; the function optionally includes NUMA details. Function `pg_buffercache_numa_pages()` remains for backward compatibility.
-  Add functions to [pg_buffercache](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache) to mark buffers as dirty (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/9ccc049df)

   The functions are `pg_buffercache_mark_dirty()`, `pg_buffercache_mark_dirty_relation()`, and `pg_buffercache_mark_dirty_all()`.
-  Allow pushdown of array comparisons in [prepared statements](../../reference/sql-commands/prepare.md#sql-prepare) to [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) foreign servers (Alexander Pyhalov) [&sect;](https://postgr.es/c/62c3b4cd9)
-  Allow the retrieval of statistics from [foreign data wrapper servers](../../reference/sql-commands/create-foreign-data-wrapper.md#sql-createforeigndatawrapper) (Corey Huinker, Etsuro Fujita) [&sect;](https://postgr.es/c/28972b6fc)

   This is enabled for [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) by using the option `restore_stats`. The default is for [`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze) to retrieve rows from the remote server to locally generate statistics.
-  Allow [file_fdw](../additional-supplied-modules-and-extensions/file_fdw-access-data-files-in-the-servers-file-system.md#file-fdw) to read files or program output that uses multi-line headers (Shinya Kato) [&sect;](https://postgr.es/c/26cb14aea)
-  Add server variable `auto_explain.log_io` to add IO reporting to [auto_explain](../additional-supplied-modules-and-extensions/auto_explain-log-execution-plans-of-slow-queries.md#auto-explain) (Tomas Vondra) [&sect;](https://postgr.es/c/61c36a34a)
-  Allow [auto_explain](../additional-supplied-modules-and-extensions/auto_explain-log-execution-plans-of-slow-queries.md#auto-explain) to add extension-specific [`EXPLAIN`](../../reference/sql-commands/explain.md#sql-explain) options via server variable `auto_explain.log_extension_options` (Robert Haas) [&sect;](https://postgr.es/c/e972dff6c)
-  Change [btree_gin](../additional-supplied-modules-and-extensions/btree_gin-gin-operator-classes-with-b-tree-behavior.md#btree-gin) to support all btree-supported cross-type comparisons (Tom Lane) [&sect;](https://postgr.es/c/e2b64fcef) [&sect;](https://postgr.es/c/fc896821c)
-  Improve performance of [bloom](../additional-supplied-modules-and-extensions/bloom-bloom-filter-index-access-method.md#bloom) indexes by using streaming reads (Xuneng Zhou) [&sect;](https://postgr.es/c/4c910f3bb) [&sect;](https://postgr.es/c/d841ca2d1)
-  Improve performance of [pgstattuple](../additional-supplied-modules-and-extensions/pgstattuple-obtain-tuple-level-statistics.md#pgstattuple) by using streaming reads (Xuneng Zhou) [&sect;](https://postgr.es/c/213f0079b) [&sect;](https://postgr.es/c/ae58189a4)
-  Allow [fuzzystrmatch](../additional-supplied-modules-and-extensions/fuzzystrmatch-determine-string-similarities-and-distance.md#fuzzystrmatch)'s `dmetaphone()` to use single-byte encodings beyond ASCII (Peter Eisentraut) [&sect;](https://postgr.es/c/e39ece034)
-  Modify [oid2name](../additional-supplied-programs/client-applications.md#oid2name) `--extended` to report the relation file path (David Bidoc) [&sect;](https://postgr.es/c/3c5ec35de)
   <a id="release-19-acknowledgements"></a>

### Acknowledgments


 The following individuals (in alphabetical order) have contributed to this release as patch authors, committers, reviewers, testers, or reporters of issues.


- *fill in later*
