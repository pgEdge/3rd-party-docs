<a id="release-19"></a>

## Release 19


**Release date:.**


2026-??-??, AS OF 2026-04-13
  <a id="release-19-highlights"></a>

### Overview


 PostgreSQL 19 contains many new features and enhancements, including:


- *fill in later*


 The above items and other new features of PostgreSQL 19 are explained in more detail in the sections below.
  <a id="release-19-migration"></a>

### Migration to Version 19


 A dump/restore using [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall) or use of [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) or logical replication is required for those wishing to migrate data from any previous release. See [Upgrading a PostgreSQL Cluster](../../server-administration/server-setup-and-operation/upgrading-a-postgresql-cluster.md#upgrading) for general information on migrating to new major releases.


 Version 19 contains a number of changes that may affect compatibility with previous releases. Observe the following incompatibilities:


-  Add server variable password_expiration_warning_threshold to warn about password expiration (Gilles Darold, Nathan Bossart) [&sect;](https://postgr.es/c/1d92e0c2c)

   The default warning period is seven days.
-  Issue a warning after successful MD5 password authentication (Nathan Bossart) [&sect;](https://postgr.es/c/bc60ee860)

   The warning can be disabled via server variable md5_password_warnings. MD5 passwords were marked marked as deprecated in Postgres 18.
-  Remove RADIUS support (Thomas Munro) [&sect;](https://postgr.es/c/a1643d40b)

   Postgres only supported RADIUS over UDP, which is unfixably insecure.
-  Force standard_conforming_strings to always be "on" in the database server (Tom Lane) [&sect;](https://postgr.es/c/457620845)

   Server variable escape_string_warning has been removed as unnecessary. Client applications still support "standard_conforming_strings = off" for compatibility with old servers.
-  Prevent carriage returns and line feeds in database, role, and tablespace names (Mahendra Singh Thalor) [&sect;](https://postgr.es/c/b380a56a3)

   This was changed to avoid security problems. pg_upgrade will also disallow upgrading from clusters that use such names.
-  Change the default index opclasses for inet and cidr data types from btree_gist to GiST (Tom Lane) [&sect;](https://postgr.es/c/b3b0b4571) [&sect;](https://postgr.es/c/b352d3d80)

   The btree_gist inet/cidr opclasses are broken because they can exclude rows that should be returned. Pg_upgrade will fail to upgrade if btree_gist inet/cidr indexes exist in the old server.
-  Stop reordering non-schema objects created by CREATE SCHEMA (Tom Lane, Jian He) [&sect;](https://postgr.es/c/a9c350d9e) [&sect;](https://postgr.es/c/404db8f9e)

   The goal of the reordering was to avoid dependencies, but it was imperfect. Postgres now uses the specified object ordering, except for foreign keys which are created last.
-  Disallow system columns from being used in COPY FROM ... WHERE (Tom Lane) [&sect;](https://postgr.es/c/21c69dc73)

   The values of such columns were not well-defined.
-  Cause transactions to pass their READ ONLY and DEFERRABLE status to postgres_fdw sessions (Etsuro Fujita) [&sect;](https://postgr.es/c/de28140de)

   This means READ ONLY transactions can no longer modify rows processed by postgres_fdw sessions.
-  Change default of max_locks_per_transactions from 64 to 128 (Heikki Linnakangas) [&sect;](https://postgr.es/c/79534f906)

   Lock size allocation has changed, so effectively settings must now be doubled to match their capacity in previous releases.
-  Change JIT to be disabled by default (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/7f8c88c2b)

   Previously JIT was enabled by default, and activated based on optimizer costs. Unfortunately, this costing has been determined to be unreliable, so require sites that are doing many large analytical queries to manually enable JIT.
-  Rename wait event type BUFFERPIN to BUFFER (Andres Freund) [&sect;](https://postgr.es/c/6c5c393b7)
-  Change index access method handlers to use a static IndexAmRoutines structure, rather than dynamically allocated ones (Matthias van de Meent) [&sect;](https://postgr.es/c/bc6374cd7)

   This is a backwardly incompatible.
-  Remove optimizer hook get_relation_info_hook and add better-placed hook build_simple_rel_hook (Robert Haas) [&sect;](https://postgr.es/c/91f33a2ae)
-  Remove MULE_INTERNAL encoding (Thomas Munro) [&sect;](https://postgr.es/c/77645d44e)

   This encoding was complex and rarely used. Databases using it will need to be dumped and restored with a different encoding.
  <a id="release-19-changes"></a>

### Changes


 Below you will find a detailed account of the changes between PostgreSQL 19 and the previous major release.
 <a id="release-19-server"></a>

#### Server
  <a id="release-19-optimizer"></a>

##### Optimizer


-  Allow NOT INs to be converted to more efficient ANTI JOINs when NULLs are not present (Richard Guo) [&sect;](https://postgr.es/c/383eb21eb)
-  Allow more LEFT JOINs to be converted to ANTI JOINs (Tender Wang, Richard Guo) [&sect;](https://postgr.es/c/cf74558fe)
-  Allow use of Memoize for ANTI JOINS with unique inner sides (Richard Guo) [&sect;](https://postgr.es/c/0da29e4cb)
-  Improve the planning of semijoins (Richard Guo) [&sect;](https://postgr.es/c/24225ad9a)
-  Improve hash join's handling of tuples with NULL join keys (Tom Lane) [&sect;](https://postgr.es/c/1811f1af9)
-  Convert IS [NOT] DISTINCT FROM NULL to IS [NOT] NULL during constant folding (Richard Guo) [&sect;](https://postgr.es/c/f41ab5157)

   The latter form is more easily optimized.
-  Perform earlier constant folding of "Var IS [NOT] NULL" in the optimizer (Richard Guo) [&sect;](https://postgr.es/c/e2debb643)

   This allows for later optimizations.
-  Allow Append and MergeAppend to consider explicit incremental sorts (Richard Guo) [&sect;](https://postgr.es/c/55a780e94)
-  Allow some aggregate processing to be performed before joins (Richard Guo, Antonin Houska) [&sect;](https://postgr.es/c/8e1185910) [&sect;](https://postgr.es/c/bd94845e8) [&sect;](https://postgr.es/c/3a08a2a8b)

   This can reduce the number of rows needed to be processed.
-  Allow negative values of pg_aggregate.aggtransspace to indicate unbounded memory usage (Richard Guo) [&sect;](https://postgr.es/c/185e30426)

   This information is used by the optimizer in planning memory usage.
-  Simplify IS [NOT] TRUE/FALSE/UNKNOWN to plain boolean expressions when the input is proven non-nullable (Richard Guo) [&sect;](https://postgr.es/c/0aaf0de7f)
-  Simplify COALESCE and ROW(...) IS [NOT] NULL to avoid evaluating unnecessary arguments (Richard Guo) [&sect;](https://postgr.es/c/10c4fe074) [&sect;](https://postgr.es/c/cb7b7ec7a)
-  Simplify IS [NOT] DISTINCT FROM to equality/inequality operators when inputs are proven non-nullable (Richard Guo) [&sect;](https://postgr.es/c/0a3796125)
-  Speed up join selectivity computations for large optimizer statistics targets (Ilia Evdokimov, David Geier) [&sect;](https://postgr.es/c/057012b20)
-  Enable proper optimizer statistics for functions returning boolean values (Tom Lane) [&sect;](https://postgr.es/c/1eccb9315)
-  Allow extended statistics on virtual generated columns (Yugo Nagata) [&sect;](https://postgr.es/c/f7f4052a4)
-  Allow function pg_restore_extended_stats() to restore optimizer extended statistics (Corey Huinker, Michael Paquier, Chao Li) [&sect;](https://postgr.es/c/0e80f3f88) [&sect;](https://postgr.es/c/302879bd6) [&sect;](https://postgr.es/c/efbebb4e8) [&sect;](https://postgr.es/c/ba97bf9cb)
-  Add function pg_clear_extended_stats() to remove extended statistics (Corey Huinker, Michael Paquier) [&sect;](https://postgr.es/c/d756fa101)
-  Adjust the optimizer to consider startup costs of partial paths (Robert Haas, Tomas Vondra) [&sect;](https://postgr.es/c/8300d3ad4)
  <a id="release-19-performance"></a>

##### General Performance


-  Improve performance of foreign key constraint checks (Junwang Zhao, Amit Langote, Chao Li) [&sect;](https://postgr.es/c/2da86c1ef) [&sect;](https://postgr.es/c/e484b0eea) [&sect;](https://postgr.es/c/b7b27eb41) [&sect;](https://postgr.es/c/5c54c3ed1)
-  Improve asynchronous I/O read-ahead scheduling for large requests (Andres Freund) [&sect;](https://postgr.es/c/a9ee66881) [&sect;](https://postgr.es/c/8ca147d58) [&sect;](https://postgr.es/c/f63ca3379)
-  Allow io_method method "worker" to automatically control needed background workers (Thomas Munro) [&sect;](https://postgr.es/c/d1c01b79d)

   New server variables are io_min_workers, io_max_workers, io_worker_idle_timeout, and io_worker_launch_interval.
-  Allow query table scans to mark pages as all-visible in the visibility map (Melanie Plageman) [&sect;](https://postgr.es/c/b46e1e54d)

   Previously only VACUUM and COPY FREEZE could do this.
-  Allow autovacuum to use parallel vacuum workers (Daniil Davydov) [&sect;](https://postgr.es/c/1ff3180ca) [&sect;](https://postgr.es/c/2a3d2f9f6)

   This is enabled via server variable autovacuum_max_parallel_workers and per-table storage parameter autovacuum_parallel_workers.
-  Allow TID Range Scans to be parallelized (Cary Huang, David Rowley) [&sect;](https://postgr.es/c/0ca3b1697)
-  Improve COPY FROM performance for text and CSV output using SIMD CPU instructions (Nazir Bilal Yavuz, Shinya Kato) [&sect;](https://postgr.es/c/e0a3a3fd5)
-  Improve NOTIFY to only wake up backends that are listening to specified notifications (Joel Jacobson) [&sect;](https://postgr.es/c/282b1cde9)

   Previously most backends were woken by NOTIFY.
-  Allow the addition of columns based on domains containing constraints to usually avoid a table rewrite (Jian He) [&sect;](https://postgr.es/c/a0b6ef29a)

   Previously this always required a table rewrite.
-  Change the default TOAST compression method from pglz to the more efficient lz4 (Euler Taveira) [&sect;](https://postgr.es/c/34dfca293)

   This is done by changing the default for server variable default_toast_compression.
-  Improve performance of internal row deformation (David Rowley) [&sect;](https://postgr.es/c/c456e3911)
-  Improve performance of hash index bulk-deletion by using streaming reads (Xuneng Zhou) [&sect;](https://postgr.es/c/bfa3c4f10)
-  Improve sort performance using radix sorts (John Naylor) [&sect;](https://postgr.es/c/ef3c3cf6d)
-  Improve timing performance measurements (Lukas Fittl, Andres Freund, David Geier) [&sect;](https://postgr.es/c/294520c44) [&sect;](https://postgr.es/c/16fca4825)

   This benefits EXPLAIN (ANALYZE, TIMING) and pg_test_timing, and is controlled via server variable timing_clock_source.
-  Optimize plpgsql syntax SELECT simple-expression INTO var (Tom Lane) [&sect;](https://postgr.es/c/ce8d5fe0e)
-  Improve performance of numeric operations on platforms without 128-bit integer support (Dean Rasheed) [&sect;](https://postgr.es/c/d699687b3)
  <a id="release-19-system-views"></a>

##### System Views


-  Add system view pg_stat_lock and function pg_stat_get_lock() to report per-lock type statistics (Bertrand Drouvot) [&sect;](https://postgr.es/c/4019f725f)
-  Add system view pg_stat_recovery to report recovery status (Xuneng Zhou, Shinya Kato) [&sect;](https://postgr.es/c/01d485b14) [&sect;](https://postgr.es/c/2d4ead6f4)
-  Add mem_exceeded_count column to system view pg_stat_replication_slots (Bertrand Drouvot) [&sect;](https://postgr.es/c/d3b6183dd)

   This reports the number of times that logical_decoding_work_mem was exceeded.
-  Add stats_reset column to system views pg_stat_all_tables, pg_stat_all_indexes, and pg_statio_all_sequences (Bertrand Drouvot, Sami Imseih, Shihao Zhong) [&sect;](https://postgr.es/c/a5b543258)

   It also appears in the "sys" and "user" view variants.
-  Add stats_reset column to system views pg_stat_user_functions and pg_stat_database_conflicts (Bertrand Drouvot, Shihao Zhong) [&sect;](https://postgr.es/c/b71bae41a) [&sect;](https://postgr.es/c/8fe315f18)
-  Add system view pg_stat_autovacuum_scores to report per-table autovacuum details (Sami Imseih) [&sect;](https://postgr.es/c/87f61f0c8)
-  Add vacuum initiation details to system view pg_stat_progress_vacuum (Shinya Kato) [&sect;](https://postgr.es/c/0d7895206)

   The new "started_by" column reports the initiator of the vacuum, and "mode" indicates its aggressiveness.
-  Add analyze initiation details to system view pg_stat_progress_analyze (Shinya Kato) [&sect;](https://postgr.es/c/ab40db385)

   The new "started_by" column reports the initiator of the analyze.
-  Add a column to system view pg_stat_progress_basebackup to report the type of backup (Shinya Kato) [&sect;](https://postgr.es/c/deb674454)

   Possible values are "full" or "incremental".
-  Add reporting of the bytes written to WAL for full page images (Shinya Kato) [&sect;](https://postgr.es/c/f9a09aa29)

   This is accessible via system view pg_stat_wal and function pg_stat_get_backend_wal().
-  Add "connecting" status to system view column pg_stat_wal_receiver.status (Xuneng Zhou) [&sect;](https://postgr.es/c/a36164e74)
-  Add columns to system views pg_stats, pg_stats_ext, and pg_stats_ext_exprs (Corey Huinker) [&sect;](https://postgr.es/c/3b88e50d6)

   Adds table OID and attribute number columns to pg_stats, and table OID and statistics object OID columns to the other two.
-  Add information about range type extended statistics to system view pg_stats_ext_exprs (Corey Huinker, Michael Paquier) [&sect;](https://postgr.es/c/307447e6d)
-  Add system view pg_dsm_registry_allocations to report dynamic shared memory details (Florents Tselai, Nathan Bossart) [&sect;](https://postgr.es/c/167ed8082) [&sect;](https://postgr.es/c/f894acb24)
-  Add column "location" to system views pg_available_extensions and pg_available_extension_versions to report the file system directory of extensions (Matheus Alcantara) [&sect;](https://postgr.es/c/f3c9e341c)
  <a id="release-19-monitoring"></a>

##### Monitoring


-  Allow log_min_messages log levels to be specified by process type (Euler Taveira) [&sect;](https://postgr.es/c/38e0190ce)

   The new format is "type:level". A value without a colon controls unspecified process types, enabling backward compatibility.
-  Add server variable log_autoanalyze_min_duration to log long-running autoanalyze operations (Shinya Kato) [&sect;](https://postgr.es/c/dd3ae3783)

   Server variable log_autovacuum_min_duration now only controls logging of automatic vacuum operations.
-  Enable server variable log_lock_waits by default (Laurenz Albe) [&sect;](https://postgr.es/c/2aac62be8)
-  Add server variable debug_print_raw_parse to log the raw parse tree (Chao Li) [&sect;](https://postgr.es/c/06473f5a3)

   This is also enabled when the server is started with debug level 3 and higher.
-  Make messages coming from remote servers appear in the server logs in the same format as local server messages (Vignesh C) [&sect;](https://postgr.es/c/112faf137)

   These include replication, postgres_fdw, and dblink servers.
-  Add WAL full page write bytes reporting to VACUUM and ANALYZE logging (Shinya Kato) [&sect;](https://postgr.es/c/ad25744f4)
-  Add IO wait events for COPY FROM/TO on a pipe/file/program (Nikolay Samokhvalov) [&sect;](https://postgr.es/c/e05a24c2d)
-  Add wait events for WAL write and flush LSNs (Xuneng Zhou) [&sect;](https://postgr.es/c/7a39f43d8)
-  Have pg_get_sequence_data function return the sequence page LSN (Vignesh C) [&sect;](https://postgr.es/c/b93172ca5)
-  Add function pg_get_multixact_stats() to report multixact activity (Naga Appani) [&sect;](https://postgr.es/c/97b101776)
-  Issue warnings when the wraparound of xid and multi-xids is less then 100 million (Nathan Bossart) [&sect;](https://postgr.es/c/48f11bfa0)

   The previous warning was 40 million. Warnings are issued to clients and the server log.
  <a id="release-19-server-config"></a>

##### Server Configuration


-  Allow online enabling and disabling of data checksums (Daniel Gustafsson, Magnus Hagander, Tomas Vondra) [&sect;](https://postgr.es/c/f19c0ecca) [&sect;](https://postgr.es/c/b364828f8)

   Previously the checksum status could only be set at initialization and changed only while the cluster was offline using pg_checksums.
-  Add scoring system to control the order that tables are autovacuumed (Nathan Bossart) [&sect;](https://postgr.es/c/d7965d65f)

   The new server variables are autovacuum_freeze_score_weight, autovacuum_multixact_freeze_score_weight, autovacuum_vacuum_score_weight, autovacuum_vacuum_insert_score_weight, and autovacuum_analyze_score_weight.
-  Add server-side report for SNI (Server Name Indication) (Daniel Gustafsson, Jacob Champion) [&sect;](https://postgr.es/c/4f433025f)

   New configuration file PGDATA/pg_hosts.conf specifies hostname/key pairs.
-  Add a new OAUTH flow hook PQAUTHDATA_OAUTH_BEARER_TOKEN_V2 (Jacob Champion) [&sect;](https://postgr.es/c/e982331b5) [&sect;](https://postgr.es/c/0af4d402c)

   This is an improved version of PQAUTHDATA_OAUTH_BEARER_TOKEN by adding the issuer identifier and error message specification.
-  Allow background workers to be configured to terminate before database-level operations (Aya Iwata) [&sect;](https://postgr.es/c/f1e251be8)
-  Allow server variables that represent lists to be emptied by setting the value to NULL (Tom Lane) [&sect;](https://postgr.es/c/ff4597acd)
-  Update GB18030 encoding from version 2000 to 2022 (Chao Li, Zheng Tao) [&sect;](https://postgr.es/c/5334620ee)

   See the commit message for compatibility details.
  <a id="release-19-replication"></a>

##### Streaming Replication and Recovery


-  Allow standbys to wait for LSN values to be replayed via WAIT FOR (Kartyshov Ivan, Alexander Korotkov, Xuneng Zhou) [&sect;](https://postgr.es/c/447aae13b) [&sect;](https://postgr.es/c/49a181b5d)
-  Improve function pg_sync_replication_slots() to wait for the synchronization completion (Ajin Cherian, Zhijie Hou) [&sect;](https://postgr.es/c/0d2d4a0ec)

   Previously, certain synchronization failures would not be reported.
-  Add server variable wal_sender_shutdown_timeout to limit replica synchronization waits during shutdown (Andrey Silitskiy, Hayato Kuroda) [&sect;](https://postgr.es/c/a8f45dee9)

   By default, senders still wait forever for synchronization.
-  Allow wal_receiver_timeout to be set per subscription and user (Fujii Masao) [&sect;](https://postgr.es/c/8a6af3ad0) [&sect;](https://postgr.es/c/fb80f388f)

   This allows subscriptions to use different wal_receiver_timeout values.
-  Add optional pid parameter to pg_replication_origin_session_setup() to allow parallelization of SQL-level replication solutions (Doruk Yilmaz, Hayato Kuroda) [&sect;](https://postgr.es/c/5b148706c)
  <a id="release-19-logical"></a>

##### [Logical Replication]


-  Allow sequence values stored in subscribers to match the publisher (Vignesh C) [&sect;](https://postgr.es/c/f0b3573c3) [&sect;](https://postgr.es/c/5509055d6) [&sect;](https://postgr.es/c/55cefadde)

   This is enabled during CREATE SUBSCRIPTION, ALTER SUBSCRIPTION ... REFRESH PUBLICATION, and ALTER SUBSCRIPTION ... REFRESH SEQUENCES. The latter only updates values, not sequence existence. Function pg_get_sequence_data() allows inspection of sequence synchronization.
-  Allow publications to publish all sequences via the ALL SEQUENCES clause (Vignesh C, Tomas Vondra) [&sect;](https://postgr.es/c/96b378497)
-  Enhance ALTER SUBSCRIPTION on publications to synchronize the existence of sequences on subscribers to match the publisher (Vignesh C) [&sect;](https://postgr.es/c/f0b3573c3)
-  Allow CREATE/ALTER PUBLICATION to exclude some tables using the EXCEPT clause (Vignesh C, Shlok Kyal) [&sect;](https://postgr.es/c/493f8c643) [&sect;](https://postgr.es/c/6b0550c45) [&sect;](https://postgr.es/c/fd366065e) [&sect;](https://postgr.es/c/5984ea868)

   This is useful when specifying ALL TABLES.
-  Allow CREATE SUBSCRIPTION to use postgres_fdw foreign data wrapper connection parameters (Jeff Davis) [&sect;](https://postgr.es/c/8185bb534)

   The connection parameters are referenced via CREATE SUBSCRIPTION ... SERVER.
-  When server variable wal_level is "replica", allow the automatic enablement of logical replication when needed (Masahiko Sawada) [&sect;](https://postgr.es/c/67c20979c)

   New server variable effective_wal_level reports the effective WAL level.
-  Add logical subscriber setting retain_conflict_info to retain information needed for conflict resolution (Zhijie Hou) [&sect;](https://postgr.es/c/228c37086)
-  Report cases where an update is applied to a row that was already deleted on a subscriber (Zhijie Hou) [&sect;](https://postgr.es/c/fd5a1a0c3)

   This requires the subscriber have retain_dead_tuples enabled.
-  Re-enable retain_dead_tuples when the necessary transaction retention falls below max_retention_duration (Zhijie Hou) [&sect;](https://postgr.es/c/0d48d393d)
-  Add subscription option max_retention_duration to limit retain_dead_tuples retention (Zhijie Hou) [&sect;](https://postgr.es/c/a850be2fe)

   When the limit is reached, dead tuple retention until manually re-enabled or a new subscription is created.
-  Add column pg_stat_subscription_stats.sync_seq_error_count to report sequence synchronization errors (Vignesh C) [&sect;](https://postgr.es/c/f6a4c498d) [&sect;](https://postgr.es/c/3edaf29fa)
-  Rename column sync_error_count to sync_table_error_count in system view pg_stat_subscription_stats (Vignesh C) [&sect;](https://postgr.es/c/3edaf29fa)

   This is necessary since sequences errors are now also tracked.
-  Add slot synchronization skip information to pg_stat_replication_slots (Shlok Kyal) [&sect;](https://postgr.es/c/76b78721c) [&sect;](https://postgr.es/c/e68b6adad) [&sect;](https://postgr.es/c/5db6a344a)

   The new columns are slotsync_skip_count, slotsync_last_skip, and slotsync_skip_reason.
   <a id="release-19-query"></a>

#### Query Commands


-  Add support for SQL Property Graph Queries (SQL/PGQ) (Peter Eisentraut, Ashutosh Bapat) [&sect;](https://postgr.es/c/2f094e7ac) [&sect;](https://postgr.es/c/c5b3253b8) [&sect;](https://postgr.es/c/a0dd0702e)

   Internally these are processed like views so are written as standard relational queries.
-  Add UPDATE/DELETE FOR PORTION OF (Paul A. Jungwirth) [&sect;](https://postgr.es/c/8e72d914c) [&sect;](https://postgr.es/c/b6ccd30d8)

   This allows operations on a temporal range.
-  Add GROUP BY ALL syntax to automatically group all non-aggregate and non-window function target list parameters (David Christensen) [&sect;](https://postgr.es/c/ef38a4d97)
-  Allow GROUP BY to process target list subqueries that have expressions referencing non-subquery columns (Tom Lane) [&sect;](https://postgr.es/c/415100aa6)

   Also fix a bug in how GROUPING() handles target list subquery aliases.
-  Allow window functions to ignore NULLs with IGNORE NULLS/RESPECT NULLS option (Oliver Ford, Tatsuo Ishii) [&sect;](https://postgr.es/c/25a30bbd4)

   Supported window functions are lead, lag, first_value, last_value and nth_value.
-  Add support for INSERT ... ON CONFLICT DO SELECT ... RETURNING (Andreas Karlsson, Marko Tiikkaja, Viktor Holmberg) [&sect;](https://postgr.es/c/88327092f)

   This allows conflicting rows to be returned, and optionally locked with FOR UPDATE/SHARE.
  <a id="release-19-utility"></a>

#### Utility Commands


-  Create a REPACK command that replaces VACUUM FULL and CLUSTER (Antonin Houska) [&sect;](https://postgr.es/c/ac58465e0)

   The two former commands did similar things, but with confusing names, so unify them as REPACK.
-  Allow REPACK to rebuild tables without access-exclusive locking (Antonin Houska, Mihail Nikalayeu, Álvaro Herrera) [&sect;](https://postgr.es/c/28d534e2a) [&sect;](https://postgr.es/c/8fb95a8ab) [&sect;](https://postgr.es/c/e76d8c749)

   This is enabled via the CONCURRENTLY option. Server variables max_repack_replication_slots was also added.
-  Allow partitions to be merged and split using ALTER TABLE ... MERGE/SPLIT PARTITIONS (Dmitry Koval, Alexander Korotkov, Tender Wang, Richard Guo, Dagfinn Ilmari Mannsåker, Fujii Masao, Jian He) [&sect;](https://postgr.es/c/f2e4cc427) [&sect;](https://postgr.es/c/4b3d17362)
-  Allow GRANT/REVOKE to specify the effective role performing the privileges adjustment (Nathan Bossart, Tom Lane) [&sect;](https://postgr.es/c/dd1398f13)

   The GRANTED BY clause controls this.
-  Allow CREATE SCHEMA to create more types of non-schema objects (Kirill Reshke, Jian He, Tom Lane) [&sect;](https://postgr.es/c/d51697484)
-  Allow CHECKPOINT to accept a list of options (Christoph Berg) [&sect;](https://postgr.es/c/a4f126516) [&sect;](https://postgr.es/c/2f698d7f4) [&sect;](https://postgr.es/c/8d33fbacb)

   Supported options are MODE and FLUSH_UNLOGGED.
-  Add CONNECTION clause to CREATE FOREIGN DATA WRAPPER to specify a function to be called for subscription connection parameters (Jeff Davis, Noriyoshi Shinoda) [&sect;](https://postgr.es/c/8185bb534) [&sect;](https://postgr.es/c/90630ec42)
-  Add memory usage and parallelism reporting to VACUUM (VERBOSE) and autovacuum logs (Tatsuya Kawata, Daniil Davydov) [&sect;](https://postgr.es/c/736f754ee) [&sect;](https://postgr.es/c/adcdbe938)
 <a id="release-19-constraints"></a>

##### [Constraints]


-  Allow ALTER TABLE ALTER CONSTRAINT ... [NOT] ENFORCED for CHECK constraints (Jian He) [&sect;](https://postgr.es/c/342051d73)

   Previously enforcement changes were only supported for foreign key constraints.
-  Allow ALTER COLUMN SET EXPRESSION to succeed on virtual columns with CHECK constraints (Jian He) [&sect;](https://postgr.es/c/f80bedd52)

   This was previously prohibited.
-  Reduce lock level of ALTER DOMAIN ... VALIDATE CONSTRAINT to match ALTER TABLE ... VALIDATE CONSTRAINT (Jian He) [&sect;](https://postgr.es/c/16a0039dc)
  <a id="release-19-copy"></a>

##### [sql-copy]


-  Allow multiple headers lines be skipped by COPY FROM (Shinya Kato, Fujii Masao) [&sect;](https://postgr.es/c/bc2f348e8)

   Previously only a single header line could be skipped.
-  Allow COPY FROM to set invalid input values to NULL (Jian He, Kirill Reshke) [&sect;](https://postgr.es/c/2a525cc97)

   This is done using the COPY option ON_ERROR SET_NULL.
-  Allow COPY TO to output JSON format (Joe Conway, Jian He, Andrew Dunstan) [&sect;](https://postgr.es/c/7dadd38cd)
-  Allow COPY TO in JSON format to output its results as a single JSON array (Joe Conway, Jian He) [&sect;](https://postgr.es/c/4c0390ac5)

   The COPY option is FORCE_ARRAY.
-  Allow COPY TO to output partitioned tables (Jian He, Ajin Cherian) [&sect;](https://postgr.es/c/4bea91f21) [&sect;](https://postgr.es/c/266543a62)

   Previously COPY (SELECT ...) had to be used to output partitioned tables. This also improves logical replication table synchronization.
  <a id="release-19-explain"></a>

##### [sql-explain]


-  Add EXPLAIN ANALYZE option IO to report asynchronous IO activity (Tomas Vondra) [&sect;](https://postgr.es/c/681daed93) [&sect;](https://postgr.es/c/3b1117d6e) [&sect;](https://postgr.es/c/e157fe6f7)
-  Add WAL full page write bytes reporting to EXPLAIN (ANALYZE, WAL) (Shinya Kato) [&sect;](https://postgr.es/c/5ab0b6a24)
-  Add Memoize cache and lookup estimates to EXPLAIN output (Ilia Evdokimov, Lukas Fittl) [&sect;](https://postgr.es/c/4bc62b868)

   This will help illustrate why Memoize was chosen.
   <a id="release-19-datatypes"></a>

#### Data Types


-  Add the 64-bit unsigned data type oid8 (Michael Paquier) [&sect;](https://postgr.es/c/b139bd3b6)
-  Add more jsonpath string methods (Florents Tselai, David E. Wheeler) [&sect;](https://postgr.es/c/bd4f879a9)

   They are l/r/btrim(), lower(), upper(), initcap(), replace(), and split_part(). These are immutable like their non-JSON string variants.
-  Allow casts between bytea and uuid data types (Dagfinn Ilmari Mannsåker, Aleksander Alekseev) [&sect;](https://postgr.es/c/ba21f5bf8)
-  Add ability to cast between database names and oids using regdatabase (Ian Lawrence Barwick) [&sect;](https://postgr.es/c/bd09f024a)
-  Add functions tid_block() and tid_offset() to extract block numbers and offsets from tid values (Ayush Tiwari) [&sect;](https://postgr.es/c/df6949ccf)
  <a id="release-19-functions"></a>

#### Functions


-  Add date, timestamp, and timestamptz versions of random(min, max) (Damien Clochard, Dean Rasheed) [&sect;](https://postgr.es/c/faf071b55) [&sect;](https://postgr.es/c/9c24111c4)
-  Allow encode() and decode() to process data in base64url and base32hex formats (Andrey Borodin, Aleksander Alekseev, Florents Tselai) [&sect;](https://postgr.es/c/497c1170c) [&sect;](https://postgr.es/c/e752a2ccc) [&sect;](https://postgr.es/c/e1d917182)

   This format retains ordering, unlike base32.
-  Add functions to return a set of ranges resulting from range subtraction (Paul A. Jungwirth) [&sect;](https://postgr.es/c/5eed8ce50)

   The functions are range_minus_multi() and multirange_minus_multi(). This is useful to represent range subtractions results with gaps.
-  Add function error_on_null() to return the supplied parameter, or error on NULL input (Joel Jacobson) [&sect;](https://postgr.es/c/2b75c38b7)
-  Allow IS JSON to work on domains defined over supported base types (Jian He) [&sect;](https://postgr.es/c/3b4c2b9db)

   The supported base domains are TEXT, JSON, JSONB, and BYTEA.
-  Add full text stemmers for Polish and Esperanto (Tom Lane) [&sect;](https://postgr.es/c/7dc95cc3b)

   The Dutch stemmer has also been updated. The old Dutch stemmer is available via "dutch_porter".
-  Modify pg_read_all_data() and pg_write_all_data() to read/write large objects (Nitin Motiani, Nathan Bossart) [&sect;](https://postgr.es/c/d98197602)

   These functions are designed to allow non-super users to run pg_dump.
-  Add function pg_get_role_ddl() to output role creation commands (Mario Gonzalez, Bryan Green, Andrew Dunstan, Euler Taveira) [&sect;](https://postgr.es/c/76e514ebb)
-  Add function pg_get_tablespace_ddl() to output tablespace creation commands (Nishant Sharma, Manni Wood, Andrew Dunstan, Euler Taveira) [&sect;](https://postgr.es/c/b99fd9fd7)
-  Add function pg_get_database_ddl() to output database creation commands (Akshay Joshi, Andrew Dunstan, Euler Taveira) [&sect;](https://postgr.es/c/a4f774cf1)
-  Allow event triggers to be written using PL/Python (Euler Taveira, Dimitri Fontaine) [&sect;](https://postgr.es/c/53eff471c)
  <a id="release-19-libpq"></a>

#### [Libpq]


-  Allow libpq connections to specify a service file via "servicefile" (Torsten Förtsch, Ryo Kanbayashi) [&sect;](https://postgr.es/c/092f3c63e)
-  Add special libpq protocol version 3.9999 for version testing (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/d8d7c5dc8)
-  Add libpq function PQgetThreadLock() to retrieve the current locking callback (Jacob Champion) [&sect;](https://postgr.es/c/b8d768583)
-  Add libpq connection setting oauth_ca_file to specify the OAUTH certificate authority file (Jonathan Gonzalez V., Jacob Champion) [&sect;](https://postgr.es/c/993368113)

   This can also be set via the PGOAUTHCAFILE environment variable. The default is to use curl's built-in certificates.
-  Allow custom OAUTH validators to register custom pg_hba.conf authentication options (Jacob Champion) [&sect;](https://postgr.es/c/b977bd308)
-  Allow OAUTH validators to supply failure details (Jacob Champion) [&sect;](https://postgr.es/c/d438a3659)

   This is done by setting the ValidatorModuleResult structure member error_detail.
-  Allow libpq environment variable PGOAUTHDEBUG to specify specific debug options (Zsolt Parragi, Jacob Champion) [&sect;](https://postgr.es/c/6d00fb904)

   The UNSAFE option still generates all debugging output.
  <a id="release-19-psql"></a>

#### [app-psql]


-  Allow the search path to appear in the psql prompt via "%S" (Florents Tselai) [&sect;](https://postgr.es/c/b3ce55f41)

   This works when psql is connected to Postgres 18 or later.
-  Allow the hot standby status to appear in the psql prompt via "%i" (Jim Jones) [&sect;](https://postgr.es/c/dddbbc253)
-  Modify psql backslash commands to show comments for publications, subscriptions, and extended statistics (Fujii Masao, Jim Jones) [&sect;](https://postgr.es/c/aecc55866)

   The modified commands are \dRp+, \dRs+, and \dX+.
-  Allow control over how booleans are displayed in psql (David G. Johnston) [&sect;](https://postgr.es/c/645cb44c5)

   The \pset variables are display_true and display_false.
-  Add psql variable SERVICEFILE to reference the service file location (Ryo Kanbayashi) [&sect;](https://postgr.es/c/6b1c4d326)
-  Allow psql to more accurately determine if the pager is needed (Erik Wienhold) [&sect;](https://postgr.es/c/27da1a796)
-  Add or improve psql tab completion (Yamaguchi Atsuo, Yugo Nagata, Haruna Miwa, Xuneng Zhou, Yugo Nagata, Dagfinn Ilmari Mannsåker, Fujii Masao, Álvaro Herrera, Jian He, Fujii Masao, Tatsuya Kawata, Ian Lawrence Barwick, Vasuki M) [&sect;](https://postgr.es/c/5fa7837d9) [&sect;](https://postgr.es/c/c6a7d3bab) [&sect;](https://postgr.es/c/81966c545) [&sect;](https://postgr.es/c/a1f7f91be) [&sect;](https://postgr.es/c/6d2ff1de4) [&sect;](https://postgr.es/c/02fd47dbf) [&sect;](https://postgr.es/c/14ee8e640) [&sect;](https://postgr.es/c/ff0bcb248) [&sect;](https://postgr.es/c/86c539c5a) [&sect;](https://postgr.es/c/a604affad) [&sect;](https://postgr.es/c/a4c10de92) [&sect;](https://postgr.es/c/28c4b8a05) [&sect;](https://postgr.es/c/0bf7d4ca9) [&sect;](https://postgr.es/c/344b572e3)
  <a id="release-19-server-apps"></a>

#### Server Applications


-  Change vacuumdb's --analyze-only option to analyze partitioned tables when no targets are specified (Laurenz Albe, Mircea Cadariu) [&sect;](https://postgr.es/c/6429e5b77)

   Previously it skipped partitioned tables. This now matches the behavior of ANALYZE.
-  Allow vacuumdb to report its commands without running them using option --dry-run (Corey Huinker) [&sect;](https://postgr.es/c/d107176d2)
-  Allow pg_verifybackup to read WAL files stored in tar archives (Amul Sul) [&sect;](https://postgr.es/c/b3cf461b3)

   Add option --wal-path as an alias for the existing and deprecated --wal-directory option.
-  Allow pg_waldump to read WAL files stored in tar archives (Amul Sul) [&sect;](https://postgr.es/c/b15c15139)
-  Add pgbench option --continue-on-error to continue after SQL errors (Rintaro Ikeda, Yugo Nagata, Fujii Masao) [&sect;](https://postgr.es/c/0ab208fa5)
-  Improve the usability of pg_test_timing (Hannu Krosing, Tom Lane) [&sect;](https://postgr.es/c/0b096e379) [&sect;](https://postgr.es/c/9dcc76414)

   Report nanoseconds instead of microseconds. In addition to histogram output, output a second table that reports exact timings, with an optional cutoff set by --cutoff.
 <a id="release-19-pgdump"></a>

##### [pg_dump]/[pg_dumpall]/[pg_restore]


-  Allow pg_dumpall to product output in non-text formats (Mahendra Singh Thalor, Andrew Dunstan) [&sect;](https://postgr.es/c/763aaa06f) [&sect;](https://postgr.es/c/3c19983cc)

   The new output formats are custom, directory, or tar.
-  Allow pg_dump to include restorable extended statistics (Corey Huinker) [&sect;](https://postgr.es/c/c32fb29e9)
  <a id="release-19-pgupgrade"></a>

##### [pgupgrade]


-  Have pg_upgrade copy large object metadata files rather than use COPY (Nathan Bossart) [&sect;](https://postgr.es/c/3bcfcd815) [&sect;](https://postgr.es/c/158408fef)

   This is possible when upgrading from Postgres 16 and later.
-  Allow pg_upgrade to use COPY for large object metadata (Nathan Bossart) [&sect;](https://postgr.es/c/161a3e8b6)

   This is used when upgrading from Postgres major versions 12-15.
-  Improve pg_upgrade performance when restoring large object metadata for origin servers version 11 and earlier (Nathan Bossart) [&sect;](https://postgr.es/c/b33f75361)
-  Allow pg_upgrade to process non-default tablespaces stored in the PGDATA directory (Nathan Bossart) [&sect;](https://postgr.es/c/412036c22)

   Previously such tablespaces generated an error.
  <a id="release-19-logicalrep-app"></a>

##### Logical Replication Applications


-  Allow pg_createsubscriber to ignore specified publications that already exist (Shubham Khanna) [&sect;](https://postgr.es/c/85ddcc2f4)

   Previously this generated an error.
-  Change the way pg_createsubscriber stores recovery parameters (Alyona Vinter) [&sect;](https://postgr.es/c/639352d90)

   Changes are stored in optionally-included pg_createsubscriber.conf rather than directly in postgresql.auto.conf.
-  Add pg_createsubscriber option -l/--logdir to redirect output to files (Gyan Sreejith, Hayato Kuroda) [&sect;](https://postgr.es/c/6b5b7eae3)
   <a id="release-19-source-code"></a>

#### Source Code


-  Restore support for AIX (Aditya Kamath, Srirama Kucherlapati, Peter Eisentraut) [&sect;](https://postgr.es/c/ecae09725) [&sect;](https://postgr.es/c/4a1b05caa)

   This uses gcc and only supports 64-bit builds.
-  Change Solaris to use unnamed POSIX semaphores (Tom Lane) [&sect;](https://postgr.es/c/0123ce131)

   Previously it used System V semaphores.
-  Require Visual Studio 2019 or later (Peter Eisentraut) [&sect;](https://postgr.es/c/8fd9bb1d9)
-  Allow MSVC to create PL/Python using the Python Limited API (Bryan Green) [&sect;](https://postgr.es/c/2bc60f862)
-  Allow building on AArch64 using MSVC (Niyas Sait, Greg Burd, Dave Cramer) [&sect;](https://postgr.es/c/a516b3f00)
-  Allow execution stack backtraces on Windows using DbgHelp (Bryan Green) [&sect;](https://postgr.es/c/65707ed9a)
-  Change the supported C language version to C11 (Peter Eisentraut) [&sect;](https://postgr.es/c/f5e0186f8) [&sect;](https://postgr.es/c/4fbe01514)

   Previously C99 was used.
-  Use standard C23 and C++ attributes if available (Peter Eisentraut) [&sect;](https://postgr.es/c/76f4b92ba)
-  Allow C++ compiler mode to be used with ICU (John Naylor) [&sect;](https://postgr.es/c/ed26c4e25)
-  Optionally use AVX2 CPU instructions for calculating page checksums (Matthew Sterrett, Andrew Kim) [&sect;](https://postgr.es/c/5e13b0f24)
-  Optionally use ARM Crypto Extension to Compute CRC32C (John Naylor) [&sect;](https://postgr.es/c/fbc57f2bc)
-  Change hex_encode() and hex_decode() to use SIMD CPU instructions (Nathan Bossart, Chiranmoy Bhattacharya) [&sect;](https://postgr.es/c/ec8719ccb)
-  Require Meson version 0.57.2 or later (Peter Eisentraut) [&sect;](https://postgr.es/c/f039c2244)
-  Add Meson option to build both shared and static libraries, or only shared (Peter Eisentraut) [&sect;](https://postgr.es/c/78727dcba)
-  Update Unicode data to version 17.0.0 (Peter Eisentraut) [&sect;](https://postgr.es/c/57ee39795)
-  Add hooks planner_setup_hook and planner_shutdown_hook (Robert Haas) [&sect;](https://postgr.es/c/94f3ad396)
-  Allow extensions to replace set-returning functions in the FROM clause with SQL queries (Paul A. Jungwirth) [&sect;](https://postgr.es/c/b140c8d7a)
-  Make multixid members 64-bit (Maxim Orlov) [&sect;](https://postgr.es/c/bd8d9c9bd)
-  Add fake LSN support to hash index AM (Peter Geoghegan) [&sect;](https://postgr.es/c/e5836f7b7)
-  Change FDW function prototypes to use uint* instead of bit* typedefs (Nathan Bossart) [&sect;](https://postgr.es/c/bab2f27ea)
-  Allow logical decoding plugins to specify if they do not access shared catalogs (Antonin Houska) [&sect;](https://postgr.es/c/0d3dba38c)
-  Add simplified and improved shared memory registration function ShmemRequestStruct (Heikki Linnakangas, Ashutosh Bapat) [&sect;](https://postgr.es/c/283e823f9)

   Functions ShmemInitStruct() and ShmemInitHash() remain for backward compatibility.
-  Add server variable debug_exec_backend to report how parameters are passed to new backends (Daniel Gustafsson) [&sect;](https://postgr.es/c/b3fe098d3)
-  Document the environment variables that control the regression tests (Michael Paquier) [&sect;](https://postgr.es/c/02976b0a1)
-  Add documentation section about temporal tables (Paul A. Jungwirth) [&sect;](https://postgr.es/c/e4d8a2af0)
-  Update documented systemd example to include a restart setting (Andrew Jackson) [&sect;](https://postgr.es/c/b30656ce0)
  <a id="release-19-modules"></a>

#### Additional Modules


-  Add pg_plan_advice module to stabilize and control planner decisions (Robert Haas) [&sect;](https://postgr.es/c/5883ff30b) [&sect;](https://postgr.es/c/6455e55b0)
-  Add extension pg_stash_advice to allow per-query-id advice to be specified (Robert Haas, Lukas Fittl) [&sect;](https://postgr.es/c/e8ec19aa3) [&sect;](https://postgr.es/c/c10edb102)
-  Refactor pg_buffercache reporting of shared memory mapping (Bertrand Drouvot) [&sect;](https://postgr.es/c/4b203d499)

   New function pg_buffercache_os_pages() and system view pg_buffercache_os_pages allow reporting of shared memory mapping; the function optionally includes NUMA details. Function pg_buffercache_numa_pages() remains for backward compatibility.
-  Add functions to pg_buffercache to mark buffers as dirty (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/9ccc049df)

   The functions are pg_buffercache_mark_dirty(), pg_buffercache_mark_dirty_relation(), and pg_buffercache_mark_dirty_all().
-  Allow pushdown of array comparisons in prepared statements to postgres_fdw foreign servers (Alexander Pyhalov) [&sect;](https://postgr.es/c/62c3b4cd9)
-  Allow the retrieval of statistics from foreign data wrapper servers (Corey Huinker, Etsuro Fujita) [&sect;](https://postgr.es/c/28972b6fc)

   This is enabled for postgres_fdw by using the option restore_stats. The default is for ANALYZE to retrieve rows from the remote server to locally generate statistics.
-  Allow file_fdw to read files or program output that uses multi-line headers (Shinya Kato) [&sect;](https://postgr.es/c/26cb14aea)
-  Add server variable auto_explain.log_io to add IO reporting to auto_explain (Tomas Vondra) [&sect;](https://postgr.es/c/61c36a34a)
-  Allow auto_explain to add extension-specific EXPLAIN options via server variable auto_explain.log_extension_options (Robert Haas) [&sect;](https://postgr.es/c/e972dff6c)
-  Allow btree_gin to match partial qualifications (Tom Lane) [&sect;](https://postgr.es/c/e2b64fcef) [&sect;](https://postgr.es/c/fc896821c)
-  Improve performance of bloom indexes by using streaming reads (Xuneng Zhou) [&sect;](https://postgr.es/c/4c910f3bb) [&sect;](https://postgr.es/c/d841ca2d1)
-  Improve performance of pgstattuple by using streaming reads (Xuneng Zhou) [&sect;](https://postgr.es/c/213f0079b)
-  Allow fuzzystrmatch's dmetaphone to use single-byte encodings beyond ASCII (Peter Eisentraut) [&sect;](https://postgr.es/c/e39ece034)
-  Modify oid2name --extended to report the relation file path (David Bidoc) [&sect;](https://postgr.es/c/3c5ec35de)
 <a id="release-19-pgstatstatements"></a>

##### [pgstatstatements]


-  Show sizes of FETCH queries as constants in pg_stat_statements (Sami Imseih) [&sect;](https://postgr.es/c/bee23ea4d)

   Fetches of different sizes will now be grouped together in pg_stat_statements output.
-  Add generic and custom plans counts to pg_stat_statements (Sami Imseih) [&sect;](https://postgr.es/c/3357471cf)
    <a id="release-19-acknowledgements"></a>

### Acknowledgments


 The following individuals (in alphabetical order) have contributed to this release as patch authors, committers, reviewers, testers, or reporters of issues.


- *fill in later*
