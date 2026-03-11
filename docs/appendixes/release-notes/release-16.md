## Release 16 { #release-16 }


**Release date:.**


2023-09-14


### Overview { #release-16-highlights }


 PostgreSQL 16 contains many new features and enhancements, including:


-  Allow parallelization of `FULL` and internal right `OUTER` hash joins
-  Allow logical replication from standby servers
-  Allow logical replication subscribers to apply large transactions in parallel
-  Allow monitoring of I/O statistics using the new `pg_stat_io` view
-  Add SQL/JSON constructors and identity functions
-  Improve performance of vacuum freezing
-  Add support for regular expression matching of user and database names in `pg_hba.conf`, and user names in `pg_ident.conf`


 The above items and other new features of PostgreSQL 16 are explained in more detail in the sections below.


### Migration to Version 16 { #release-16-migration }


 A dump/restore using [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall) or use of [pgupgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) or logical replication is required for those wishing to migrate data from any previous release. See [Upgrading a PostgreSQL Cluster](../../server-administration/server-setup-and-operation/upgrading-a-postgresql-cluster.md#upgrading) for general information on migrating to new major releases.


 Version 16 contains a number of changes that may affect compatibility with previous releases. Observe the following incompatibilities:


-  Change assignment rules for [PL/pgSQL](../../server-programming/pl-pgsql-sql-procedural-language/cursors.md#plpgsql-open-bound-cursor) bound cursor variables (Tom Lane) [&sect;](https://postgr.es/c/d747dc85a)

   Previously, the string value of such variables was set to match the variable name during cursor assignment; now it will be assigned during [`OPEN`](../../server-programming/pl-pgsql-sql-procedural-language/cursors.md#plpgsql-cursor-opening), and will not match the variable name. To restore the previous behavior, assign the desired portal name to the cursor variable before `OPEN`.
-  Disallow [`NULLS NOT DISTINCT`](../../reference/sql-commands/create-index.md#sql-createindex) indexes for primary keys (Daniel Gustafsson) [&sect;](https://postgr.es/c/d95952325)
-  Change [`REINDEX DATABASE`](../../reference/sql-commands/reindex.md#sql-reindex) and [reindexdb](../../reference/postgresql-client-applications/reindexdb.md#app-reindexdb) to not process indexes on system catalogs (Simon Riggs) [&sect;](https://postgr.es/c/2cbc3c17a) [&sect;](https://postgr.es/c/0a5f06b84)

   Processing such indexes is still possible using `REINDEX SYSTEM` and [`reindexdb --system`](../../reference/postgresql-client-applications/reindexdb.md#app-reindexdb).
-  Tighten [`GENERATED`](../../the-sql-language/data-definition/generated-columns.md#ddl-generated-columns) expression restrictions on inherited and partitioned tables (Amit Langote, Tom Lane) [&sect;](https://postgr.es/c/8bf6ec3ba)

   Columns of parent/partitioned and child/partition tables must all have the same generation status, though now the actual generation expressions can be different.
-  Remove [pg_walinspect](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect) functions `pg_get_wal_records_info_till_end_of_wal()` and `pg_get_wal_stats_till_end_of_wal()` (Bharath Rupireddy) [&sect;](https://postgr.es/c/5c1b66280)
-  Rename server variable `force_parallel_mode` to [`debug_parallel_query`](../../server-administration/server-configuration/developer-options.md#guc-debug-parallel-query) (David Rowley) [&sect;](https://postgr.es/c/5352ca22e) [&sect;](https://postgr.es/c/0981846b9)
-  Remove the ability to [create views](../../reference/sql-commands/create-view.md#sql-createview) manually with `ON SELECT` rules (Tom Lane) [&sect;](https://postgr.es/c/b23cd185f)
-  Remove the server variable `vacuum_defer_cleanup_age` (Andres Freund) [&sect;](https://postgr.es/c/1118cd37e)

   This has been unnecessary since [`hot_standby_feedback`](../../server-administration/server-configuration/replication.md#guc-hot-standby-feedback) and [replication slots](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication-slots) were added.
-  Remove server variable `promote_trigger_file` (Simon Riggs) [&sect;](https://postgr.es/c/cd4329d93)

   This was used to promote a standby to primary, but is now more easily accomplished with [`pg_ctl promote`](../../reference/postgresql-server-applications/pg_ctl.md#app-pg-ctl) or [`pg_promote()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-recovery-control-table).
-  Remove read-only server variables `lc_collate` and `lc_ctype` (Peter Eisentraut) [&sect;](https://postgr.es/c/b0f6c4371)

   Collations and locales can vary between databases so having them as read-only server variables was unhelpful.
-  Role inheritance now controls the default inheritance status of member roles added during [`GRANT`](../../reference/sql-commands/grant.md#sql-grant) (Robert Haas) [&sect;](https://postgr.es/c/e3ce2de09)

   The role's default inheritance behavior can be overridden with the new `GRANT ... WITH INHERIT` clause. This allows inheritance of some roles and not others because the members' inheritance status is set at `GRANT` time. Previously the inheritance status of member roles was controlled only by the role's inheritance status, and changes to a role's inheritance status affected all previous and future member roles.
-  Restrict the privileges of [`CREATEROLE`](../../reference/sql-commands/create-role.md#sql-createrole) and its ability to modify other roles (Robert Haas) [&sect;](https://postgr.es/c/cf5eb37c5) [&sect;](https://postgr.es/c/f1358ca52)

   Previously roles with `CREATEROLE` privileges could change many aspects of any non-superuser role. Such changes, including adding members, now require the role requesting the change to have `ADMIN OPTION` permission. For example, they can now change the `CREATEDB`, `REPLICATION`, and `BYPASSRLS` properties only if they also have those permissions.
-  Remove symbolic links for the postmaster binary (Peter Eisentraut) [&sect;](https://postgr.es/c/37e267335)


### Changes { #release-16-changes }


 Below you will find a detailed account of the changes between PostgreSQL 16 and the previous major release.


#### Server { #release-16-server }


##### Optimizer { #release-16-optimizer }


-  Allow incremental sorts in more cases, including `DISTINCT` (David Rowley) [&sect;](https://postgr.es/c/b59242209) [&sect;](https://postgr.es/c/3c6fc5820)
-  Add the ability for aggregates having `ORDER BY` or `DISTINCT` to use pre-sorted data (David Rowley) [&sect;](https://postgr.es/c/1349d2790) [&sect;](https://postgr.es/c/3226f4728) [&sect;](https://postgr.es/c/da5800d5f)

   The new server variable [`enable_presorted_aggregate`](../../server-administration/server-configuration/query-planning.md#guc-enable-presorted-aggregate) can be used to disable this.
-  Allow memoize atop a `UNION ALL` (Richard Guo) [&sect;](https://postgr.es/c/9bfd2822b)
-  Allow anti-joins to be performed with the non-nullable input as the inner relation (Richard Guo) [&sect;](https://postgr.es/c/16dc2703c)
-  Allow parallelization of [`FULL`](../../the-sql-language/queries/table-expressions.md#queries-join) and internal right `OUTER` hash joins (Melanie Plageman, Thomas Munro) [&sect;](https://postgr.es/c/11c2d6fdf)
-  Improve the accuracy of [`GIN`](../../internals/gin-indexes/index.md#gin) index access optimizer costs (Ronan Dunklau) [&sect;](https://postgr.es/c/cd9479af2)


##### General Performance { #release-16-performance }


-  Allow more efficient addition of heap and index pages (Andres Freund) [&sect;](https://postgr.es/c/00d1e02be) [&sect;](https://postgr.es/c/26158b852)
-  During non-freeze operations, perform page [freezing](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound) where appropriate (Peter Geoghegan) [&sect;](https://postgr.es/c/d977ffd92) [&sect;](https://postgr.es/c/9e5405993) [&sect;](https://postgr.es/c/1de58df4f)

   This makes full-table freeze vacuums less necessary.
-  Allow window functions to use the faster [`ROWS`](../../the-sql-language/sql-syntax/value-expressions.md#syntax-window-functions) mode internally when `RANGE` mode is active but unnecessary (David Rowley) [&sect;](https://postgr.es/c/ed1a88dda)
-  Allow optimization of always-increasing window functions [`ntile()`](../../the-sql-language/functions-and-operators/window-functions.md#functions-window-table), `cume_dist()` and `percent_rank()` (David Rowley) [&sect;](https://postgr.es/c/456fa635a)
-  Allow aggregate functions [`string_agg()`](../../the-sql-language/functions-and-operators/aggregate-functions.md#functions-aggregate-table) and `array_agg()` to be parallelized (David Rowley) [&sect;](https://postgr.es/c/16fd03e95)
-  Improve performance by caching [`RANGE`](../../the-sql-language/data-definition/table-partitioning.md#ddl-partitioning-overview) and `LIST` partition lookups (Amit Langote, Hou Zhijie, David Rowley) [&sect;](https://postgr.es/c/3592e0ff9)
-  Allow control of the shared buffer usage by vacuum and analyze (Melanie Plageman) [&sect;](https://postgr.es/c/1cbbee033) [&sect;](https://postgr.es/c/ae78cae3b) [&sect;](https://postgr.es/c/b72f564d8)

   The [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum)/[`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze) option is `BUFFER_USAGE_LIMIT`, and the [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) option is `--buffer-usage-limit`. The default value is set by server variable [`vacuum_buffer_usage_limit`](../../server-administration/server-configuration/resource-consumption.md#guc-vacuum-buffer-usage-limit), which also controls autovacuum.
-  Support [`wal_sync_method=fdatasync`](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-sync-method) on `Windows` (Thomas Munro) [&sect;](https://postgr.es/c/9430fb407)
-  Allow [HOT](../../internals/database-physical-storage/heap-only-tuples-hot.md#storage-hot) updates if only `BRIN`-indexed columns are updated (Matthias van de Meent, Josef Simanek, Tomas Vondra) [&sect;](https://postgr.es/c/19d8e2308)
-  Improve the speed of updating the [process title](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-update-process-title) (David Rowley) [&sect;](https://postgr.es/c/2cb82e2ac)
-  Allow `xid`/`subxid` searches and ASCII string detection to use vector operations (Nathan Bossart, John Naylor) [&sect;](https://postgr.es/c/37a6e5df3) [&sect;](https://postgr.es/c/121d2d3d7) [&sect;](https://postgr.es/c/b6ef16756) [&sect;](https://postgr.es/c/e813e0e16)

   ASCII detection is particularly useful for [`COPY FROM`](../../reference/sql-commands/copy.md#sql-copy). Vector operations are also used for some C array searches.
-  Reduce overhead of memory allocations (Andres Freund, David Rowley) [&sect;](https://postgr.es/c/c6e0fe1f2)


##### Monitoring { #release-16-monitoring }


-  Add system view [`pg_stat_io`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-io-view) view to track I/O statistics (Melanie Plageman) [&sect;](https://postgr.es/c/a9c70b46d) [&sect;](https://postgr.es/c/8aaa04b32) [&sect;](https://postgr.es/c/ac8d53dae) [&sect;](https://postgr.es/c/0ecb87e1f) [&sect;](https://postgr.es/c/093e5c57d)
-  Record statistics on the last sequential and index scans on tables (Dave Page) [&sect;](https://postgr.es/c/c03747183)

   This information appears in [`pg_stat_*_tables`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-all-tables-view) and [`pg_stat_*_indexes`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-all-indexes-view).
-  Record statistics on the occurrence of updated rows moving to new pages (Corey Huinker) [&sect;](https://postgr.es/c/ae4fdde13)

   The `pg_stat_*_tables` column is [`n_tup_newpage_upd`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-all-tables-view).
-  Add speculative lock information to the [`pg_locks`](../../internals/system-views/pg_locks.md#view-pg-locks) system view (Masahiko Sawada, Noriyoshi Shinoda) [&sect;](https://postgr.es/c/f74573969)

   The transaction id is displayed in the `transactionid` column and the speculative insertion token is displayed in the `objid` column.
-  Add the display of prepared statement result types to the [`pg_prepared_statements`](../../internals/system-views/pg_prepared_statements.md#view-pg-prepared-statements) view (Dagfinn Ilmari Mannsåker) [&sect;](https://postgr.es/c/84ad713cf) [&sect;](https://postgr.es/c/6ffff0fd2)
-  Create subscription statistics entries at subscription creation time so [`stats_reset`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-database-view) is accurate (Andres Freund) [&sect;](https://postgr.es/c/e0b014295)

   Previously entries were created only when the first statistics were reported.
-  Correct the I/O accounting for temp relation writes shown in [`pg_stat_database`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#pg-stat-database-view) (Melanie Plageman) [&sect;](https://postgr.es/c/704261ecc)
-  Add function [`pg_stat_get_backend_subxact()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-backend-funcs-table) to report on a session's subtransaction cache (Dilip Kumar) [&sect;](https://postgr.es/c/10ea0f924)
-  Have [`pg_stat_get_backend_idset()`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-backend-funcs-table), `pg_stat_get_backend_activity()`, and related functions use the unchanging backend id (Nathan Bossart) [&sect;](https://postgr.es/c/d7e39d72c)

   Previously the index values might change during the lifetime of the session.
-  Report stand-alone backends with a special backend type (Melanie Plageman) [&sect;](https://postgr.es/c/0c679464a)
-  Add wait event [`SpinDelay`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#wait-event-timeout-table) to report spinlock sleep delays (Andres Freund) [&sect;](https://postgr.es/c/92daeca45)
-  Create new wait event [`DSMAllocate`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#wait-event-io-table) to indicate waiting for dynamic shared memory allocation (Thomas Munro) [&sect;](https://postgr.es/c/7bae3bbf6)

   Previously this type of wait was reported as `DSMFillZeroWrite`, which was also used by `mmap()` allocations.
-  Add the database name to the [process title](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-update-process-title) of logical WAL senders (Tatsuhiro Nakamori) [&sect;](https://postgr.es/c/af205152e)

   Physical WAL senders do not display a database name.
-  Add checkpoint and `REDO LSN` information to [`log_checkpoints`](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-checkpoints) messages (Bharath Rupireddy, Kyotaro Horiguchi) [&sect;](https://postgr.es/c/62c46eee2)
-  Provide additional details during client certificate failures (Jacob Champion) [&sect;](https://postgr.es/c/3a0e38504)


##### Privileges { #release-16-privileges }


-  Add predefined role [`pg_create_subscription`](../../server-administration/database-roles/predefined-roles.md#predefined-roles) with permission to create subscriptions (Robert Haas) [&sect;](https://postgr.es/c/c3afe8cf5)
-  Allow subscriptions to not require passwords (Robert Haas) [&sect;](https://postgr.es/c/c3afe8cf5) [&sect;](https://postgr.es/c/c1cc4e688) [&sect;](https://postgr.es/c/19e65dff3)

   This is accomplished with the option [`password_required=false`](../../reference/sql-commands/create-subscription.md#sql-createsubscription).
-  Simplify permissions for [`LOCK TABLE`](../../reference/sql-commands/lock.md#sql-lock) (Jeff Davis) [&sect;](https://postgr.es/c/c44f6334c)

   Previously a user's ability to perform `LOCK TABLE` at various lock levels was limited to the lock levels required by the commands they had permission to execute on the table. For example, someone with [`UPDATE`](../../reference/sql-commands/update.md#sql-update) permission could perform all lock levels except `ACCESS SHARE`, even though it was a lesser lock level. Now users can issue lesser lock levels if they already have permission for greater lock levels.
-  Allow [`ALTER GROUP group_name ADD USER user_name`](../../reference/sql-commands/alter-group.md#sql-altergroup) to be performed with `ADMIN OPTION` (Robert Haas) [&sect;](https://postgr.es/c/ce6b672e4)

   Previously `CREATEROLE` permission was required.
-  Allow [`GRANT`](../../reference/sql-commands/grant.md#sql-grant) to use `WITH ADMIN TRUE`/`FALSE` syntax (Robert Haas) [&sect;](https://postgr.es/c/e3ce2de09)

   Previously only the `WITH ADMIN OPTION` syntax was supported.
-  Allow roles that create other roles to automatically inherit the new role's rights or the ability to [`SET ROLE`](../../reference/sql-commands/set-role.md#sql-set-role) to the new role (Robert Haas, Shi Yu) [&sect;](https://postgr.es/c/e5b8a4c09) [&sect;](https://postgr.es/c/e00bc6c92)

   This is controlled by server variable [`createrole_self_grant`](../../server-administration/server-configuration/client-connection-defaults.md#guc-createrole-self-grant).
-  Prevent users from changing the default privileges of non-inherited roles (Robert Haas) [&sect;](https://postgr.es/c/48a257d44)

   This is now only allowed for inherited roles.
-  When granting role membership, require the granted-by role to be a role that has appropriate permissions (Robert Haas) [&sect;](https://postgr.es/c/ce6b672e4)

   This is a requirement even when a non-bootstrap superuser is granting role membership.
-  Allow non-superusers to grant permissions using a granted-by user that is not the current user (Robert Haas) [&sect;](https://postgr.es/c/ce6b672e4)

   The current user still must have sufficient permissions given by the specified granted-by user.
-  Add [`GRANT`](../../reference/sql-commands/grant.md#sql-grant) to control permission to use [`SET ROLE`](../../reference/sql-commands/set-role.md#sql-set-role) (Robert Haas) [&sect;](https://postgr.es/c/3d14e171e)

   This is controlled by a new `GRANT ... SET` option.
-  Add dependency tracking to roles which have granted privileges (Robert Haas) [&sect;](https://postgr.es/c/ce6b672e4)

   For example, removing `ADMIN OPTION` will fail if there are privileges using that option; `CASCADE` must be used to revoke dependent permissions.
-  Add dependency tracking of grantors for [`GRANT`](../../reference/sql-commands/grant.md#sql-grant) records (Robert Haas) [&sect;](https://postgr.es/c/6566133c5)

   This guarantees that [`pg_auth_members`](../../internals/system-catalogs/pg_auth_members.md#catalog-pg-auth-members).`grantor` values are always valid.
-  Allow multiple role membership records (Robert Haas) [&sect;](https://postgr.es/c/ce6b672e4) [&sect;](https://postgr.es/c/0101f770a)

   Previously a new membership grant would remove a previous matching membership grant, even if other aspects of the grant did not match.
-  Prevent removal of superuser privileges for the bootstrap user (Robert Haas) [&sect;](https://postgr.es/c/e530be2c5)

   Restoring such users could lead to errors.
-  Allow [`makeaclitem()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-aclitem-fn-table) to accept multiple privilege names (Robins Tharakan) [&sect;](https://postgr.es/c/b762bbde3)

   Previously only a single privilege name, like [`SELECT`](../../reference/sql-commands/select.md#sql-select), was accepted.


##### Server Configuration { #release-16-server-config }


-  Add support for Kerberos credential delegation (Stephen Frost) [&sect;](https://postgr.es/c/6633cfb21) [&sect;](https://postgr.es/c/9c0a0e2ed) [&sect;](https://postgr.es/c/f4001a553) [&sect;](https://postgr.es/c/a2eb99a01)

   This is enabled with server variable [`gss_accept_delegation`](../../server-administration/server-configuration/connections-and-authentication.md#guc-gss-accept-delegation) and libpq connection parameter [`gssdelegation`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-gssdelegation).
-  Allow the SCRAM iteration count to be set with server variable [`scram_iterations`](../../server-administration/server-configuration/connections-and-authentication.md#guc-scram-iterations) (Daniel Gustafsson) [&sect;](https://postgr.es/c/b57774300)
-  Improve performance of server variable management (Tom Lane) [&sect;](https://postgr.es/c/3057465ac) [&sect;](https://postgr.es/c/f13b2088f)
-  Tighten restrictions on which server variables can be reset (Masahiko Sawada) [&sect;](https://postgr.es/c/385366426)

   Previously, while certain variables, like [`transaction_isolation`](../../server-administration/server-configuration/client-connection-defaults.md#guc-default-transaction-isolation), were not affected by [`RESET ALL`](../../reference/sql-commands/reset.md#sql-reset), they could be individually reset in inappropriate situations.
-  Move various [`postgresql.conf`](../../server-administration/server-configuration/setting-parameters.md#config-setting-configuration-file) items into new categories (Shinya Kato) [&sect;](https://postgr.es/c/0b039e3a8)

   This also affects the categories displayed in the [`pg_settings`](../../internals/system-views/pg_settings.md#view-pg-settings) view.
-  Prevent configuration file recursion beyond 10 levels (Julien Rouhaud) [&sect;](https://postgr.es/c/d13b68411)
-  Allow [autovacuum](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum) to more frequently honor changes to delay settings (Melanie Plageman) [&sect;](https://postgr.es/c/7d71d3dd0) [&sect;](https://postgr.es/c/a9781ae11)

   Rather than honor changes only at the start of each relation, honor them at the start of each block.
-  Remove restrictions that archive files be durably renamed (Nathan Bossart) [&sect;](https://postgr.es/c/756e221db) [&sect;](https://postgr.es/c/3cabe45a8)

   The [`archive_command`](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-command) command is now more likely to be called with already-archived files after a crash.
-  Prevent [`archive_library`](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-library) and [`archive_command`](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-command) from being set at the same time (Nathan Bossart) [&sect;](https://postgr.es/c/d627ce3b7)

   Previously `archive_library` would override `archive_command`.
-  Allow the postmaster to terminate children with an abort signal (Tom Lane) [&sect;](https://postgr.es/c/51b5834cd)

   This allows collection of a core dump for a stuck child process. This is controlled by [`send_abort_for_crash`](../../server-administration/server-configuration/developer-options.md#guc-send-abort-for-crash) and [`send_abort_for_kill`](../../server-administration/server-configuration/developer-options.md#guc-send-abort-for-kill). The postmaster's `-T` switch is now the same as setting `send_abort_for_crash`.
-  Remove the non-functional postmaster `-n` option (Tom Lane) [&sect;](https://postgr.es/c/51b5834cd)
-  Allow the server to reserve backend slots for roles with [`pg_use_reserved_connections`](../../server-administration/database-roles/predefined-roles.md#predefined-roles) membership (Nathan Bossart) [&sect;](https://postgr.es/c/6e2775e4d)

   The number of reserved slots is set by server variable [`reserved_connections`](../../server-administration/server-configuration/connections-and-authentication.md#guc-reserved-connections).
-  Allow [huge pages](../../server-administration/server-configuration/resource-consumption.md#guc-huge-pages) to work on newer versions of `Windows 10` (Thomas Munro) [&sect;](https://postgr.es/c/fdd8937c0)

   This adds the special handling required to enable huge pages on newer versions of `Windows 10`.
-  Add [`debug_io_direct`](../../server-administration/server-configuration/developer-options.md#guc-debug-io-direct) setting for developer usage (Thomas Munro, Andres Freund, Bharath Rupireddy) [&sect;](https://postgr.es/c/d4e71df6d) [&sect;](https://postgr.es/c/319bae9a8)

   While primarily for developers, [`wal_sync_method=open_sync`](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-sync-method)/`open_datasync` has been modified to not use direct I/O with `wal_level=minimal`; this is now enabled with `debug_io_direct=wal`.
-  Add function [`pg_split_walfile_name()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-backup-table) to report the segment and timeline values of WAL file names (Bharath Rupireddy) [&sect;](https://postgr.es/c/cca186348) [&sect;](https://postgr.es/c/13e0d7a60)


##### [pg_hba.conf] { #release-16-pg-hba }


-  Add support for regular expression matching on database and role entries in `pg_hba.conf` (Bertrand Drouvot) [&sect;](https://postgr.es/c/8fea86830)

   Regular expression patterns are prefixed with a slash. Database and role names that begin with slashes need to be double-quoted if referenced in `pg_hba.conf`.
-  Improve user-column handling of [`pg_ident.conf`](../../server-administration/server-configuration/file-locations.md#runtime-config-file-locations) to match `pg_hba.conf` (Jelte Fennema) [&sect;](https://postgr.es/c/efb6f4a4f)

   Specifically, add support for `all`, role membership with `+`, and regular expressions with a leading slash. Any user name that matches these patterns must be double-quoted.
-  Allow include files in `pg_hba.conf` and `pg_ident.conf` (Julien Rouhaud) [&sect;](https://postgr.es/c/a54b658ce)

   These are controlled by `include`, `include_if_exists`, and `include_dir`. System views [`pg_hba_file_rules`](../../internals/system-views/pg_hba_file_rules.md#view-pg-hba-file-rules) and [`pg_ident_file_mappings`](../../internals/system-views/pg_ident_file_mappings.md#view-pg-ident-file-mappings) now display the file name.
-  Allow `pg_hba.conf` tokens to be of unlimited length (Tom Lane) [&sect;](https://postgr.es/c/de3f0e3fe)
-  Add rule and map numbers to the system view [`pg_hba_file_rules`](../../internals/system-views/pg_hba_file_rules.md#view-pg-hba-file-rules) (Julien Rouhaud) [&sect;](https://postgr.es/c/c591300a8)


##### [Localization] { #release-16-localization }


-  Determine the default encoding from the locale when using ICU (Jeff Davis) [&sect;](https://postgr.es/c/c45dc7ffb)

   Previously the default was always `UTF-8`.
-  Have [`CREATE DATABASE`](../../reference/sql-commands/create-database.md#sql-createdatabase) and [`CREATE COLLATION`](../../reference/sql-commands/create-collation.md#sql-createcollation)'s `LOCALE` options, and [initdb](../../reference/postgresql-server-applications/initdb.md#app-initdb) and [createdb](../../reference/postgresql-client-applications/createdb.md#app-createdb) `--locale` options, control non-libc collation providers (Jeff Davis)

   Previously they only controlled libc providers.
-  Add predefined collations `unicode` and `ucs_basic` (Peter Eisentraut) [&sect;](https://postgr.es/c/0d21d4b9b)

   This only works if ICU support is enabled.
-  Allow custom ICU collation rules to be created (Peter Eisentraut) [&sect;](https://postgr.es/c/30a53b792)

   This is done using [`CREATE COLLATION`](../../reference/sql-commands/create-collation.md#sql-createcollation)'s new `RULES` clause, as well as new options for [`CREATE DATABASE`](../../reference/sql-commands/create-database.md#sql-createdatabase), [createdb](../../reference/postgresql-client-applications/createdb.md#app-createdb), and [initdb](../../reference/postgresql-server-applications/initdb.md#app-initdb).
-  Allow `Windows` to import system locales automatically (Juan José Santamaría Flecha) [&sect;](https://postgr.es/c/bf03cfd16)

   Previously, only ICU locales could be imported on `Windows`.


#### [Logical Replication] { #release-16-logical }


-  Allow [logical decoding](../../server-programming/logical-decoding/index.md#logicaldecoding) on standbys (Bertrand Drouvot, Andres Freund, Amit Khandekar) [&sect;](https://postgr.es/c/0fdab27ad) [&sect;](https://postgr.es/c/be87200ef) [&sect;](https://postgr.es/c/26669757b)

   Snapshot WAL records are required for logical slot creation but cannot be created on standbys. To avoid delays, the new function [`pg_log_standby_snapshot()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-snapshot-synchronization-table) allows creation of such records.
-  Add server variable to control how logical decoding publishers transfer changes and how subscribers apply them (Shi Yu) [&sect;](https://postgr.es/c/5de94a041) [&sect;](https://postgr.es/c/1e8b61735) [&sect;](https://postgr.es/c/9f2213a7c)

   The variable is [`debug_logical_replication_streaming`](../../server-administration/server-configuration/developer-options.md#guc-debug-logical-replication-streaming).
-  Allow logical replication initial table synchronization to copy rows in binary format (Melih Mutlu) [&sect;](https://postgr.es/c/ecb696527)

   This is only possible for subscriptions marked as binary.
-  Allow parallel application of logical replication (Hou Zhijie, Wang Wei, Amit Kapila) [&sect;](https://postgr.es/c/216a78482) [&sect;](https://postgr.es/c/cd06ccd78) [&sect;](https://postgr.es/c/fce003cfd)

   The [`CREATE SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription) `STREAMING` option now supports `parallel` to enable application of large transactions by parallel workers. The number of parallel workers is controlled by the new server variable [`max_parallel_apply_workers_per_subscription`](../../server-administration/server-configuration/replication.md#guc-max-parallel-apply-workers-per-subscription). Wait events [`LogicalParallelApplyMain`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#wait-event-activity-table), `LogicalParallelApplyStateChange`, and `LogicalApplySendData` were also added. Column `leader_pid` was added to system view [`pg_stat_subscription`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-subscription) to track parallel activity.
-  Improve performance for [logical replication apply](../../server-administration/logical-replication/architecture.md#logical-replication-architecture) without a primary key (Onder Kalaci, Amit Kapila) [&sect;](https://postgr.es/c/89e46da5e)

   Specifically, `REPLICA IDENTITY FULL` can now use btree indexes rather than sequentially scanning the table to find matches.
-  Allow logical replication subscribers to process only changes that have no origin (Vignesh C, Amit Kapila) [&sect;](https://postgr.es/c/366283961) [&sect;](https://postgr.es/c/875693019)

   This can be used to avoid replication loops. This is controlled by the new `CREATE SUBSCRIPTION ... ORIGIN` option.
-  Perform logical replication [`SELECT`](../../reference/sql-commands/select.md#sql-select) and DML actions as the table owner (Robert Haas) [&sect;](https://postgr.es/c/1e10d49b6) [&sect;](https://postgr.es/c/482675987)

   This improves security and now requires subscription owners to be either superusers or to have [`SET ROLE`](../../reference/sql-commands/set-role.md#sql-set-role) permission on all roles owning tables in the replication set. The previous behavior of performing all operations as the subscription owner can be enabled with the subscription [`run_as_owner`](../../reference/sql-commands/create-subscription.md#sql-createsubscription) option.
-  Have [`wal_retrieve_retry_interval`](../../server-administration/server-configuration/replication.md#guc-wal-retrieve-retry-interval) operate on a per-subscription basis (Nathan Bossart) [&sect;](https://postgr.es/c/5a3a95385)

   Previously the retry time was applied globally. This also adds wait events [>`LogicalRepLauncherDSA`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#wait-event-lwlock-table) and `LogicalRepLauncherHash`.


#### Utility Commands { #release-16-utility }


-  Add [`EXPLAIN`](../../reference/sql-commands/explain.md#sql-explain) option `GENERIC_PLAN` to display the generic plan for a parameterized query (Laurenz Albe) [&sect;](https://postgr.es/c/3c05284d8)
-  Allow a [`COPY FROM`](../../reference/sql-commands/copy.md#sql-copy) value to map to a column's `DEFAULT` (Israel Barth Rubio) [&sect;](https://postgr.es/c/9f8377f7a)
-  Allow [`COPY`](../../reference/sql-commands/copy.md#sql-copy) into foreign tables to add rows in batches (Andrey Lepikhov, Etsuro Fujita) [&sect;](https://postgr.es/c/97da48246)

   This is controlled by the [postgres_fdw](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw) option [`batch_size`](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-options-cost-estimation).
-  Allow the `STORAGE` type to be specified by [`CREATE TABLE`](../../reference/sql-commands/create-table.md#sql-createtable) (Teodor Sigaev, Aleksander Alekseev) [&sect;](https://postgr.es/c/784cedda0) [&sect;](https://postgr.es/c/b9424d014)

   Previously only [`ALTER TABLE`](../../reference/sql-commands/alter-table.md#sql-altertable) could control this.
-  Allow [truncate triggers](../../reference/sql-commands/create-trigger.md#sql-createtrigger) on foreign tables (Yugo Nagata) [&sect;](https://postgr.es/c/3b00a944a)
-  Allow [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) and [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) to only process [`TOAST`](../../internals/database-physical-storage/toast.md#storage-toast) tables (Nathan Bossart) [&sect;](https://postgr.es/c/4211fbd84)

   This is accomplished by having [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) turn off `PROCESS_MAIN` or by [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) using the `--no-process-main` option.
-  Add [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) options to skip or update all [frozen](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound) statistics (Tom Lane, Nathan Bossart) [&sect;](https://postgr.es/c/a46a7011b)

   The options are `SKIP_DATABASE_STATS` and `ONLY_DATABASE_STATS`.
-  Change [`REINDEX DATABASE`](../../reference/sql-commands/reindex.md#sql-reindex) and [`REINDEX SYSTEM`](../../reference/sql-commands/reindex.md#sql-reindex) to no longer require an argument (Simon Riggs) [&sect;](https://postgr.es/c/2cbc3c17a) [&sect;](https://postgr.es/c/0a5f06b84)

   Previously the database name had to be specified.
-  Allow [`CREATE STATISTICS`](../../reference/sql-commands/create-statistics.md#sql-createstatistics) to generate a statistics name if none is specified (Simon Riggs) [&sect;](https://postgr.es/c/624aa2a13)


#### Data Types { #release-16-datatypes }


-  Allow non-decimal [integer literals](../../the-sql-language/sql-syntax/lexical-structure.md#sql-syntax-bit-strings) (Peter Eisentraut) [&sect;](https://postgr.es/c/6fcda9aba)

   For example, `0x42F`, `0o273`, and `0b100101`.
-  Allow [`NUMERIC`](../../the-sql-language/data-types/numeric-types.md#datatype-numeric) to process hexadecimal, octal, and binary integers of any size (Dean Rasheed) [&sect;](https://postgr.es/c/6dfacbf72)

   Previously only unquoted eight-byte integers were supported with these non-decimal bases.
-  Allow underscores in integer and numeric [constants](../../the-sql-language/sql-syntax/lexical-structure.md#sql-syntax-bit-strings) (Peter Eisentraut, Dean Rasheed) [&sect;](https://postgr.es/c/faff8f8e4)

   This can improve readability for long strings of digits.
-  Accept the spelling `+infinity` in datetime input (Vik Fearing) [&sect;](https://postgr.es/c/2ceea5adb)
-  Prevent the specification of `epoch` and `infinity` together with other fields in datetime strings (Joseph Koshakow) [&sect;](https://postgr.es/c/bcc704b52)
-  Remove undocumented support for date input in the form <code>Y</code><em>year</em><code>M</code><em>month</em><code>D</code><em>day</em> (Joseph Koshakow) [&sect;](https://postgr.es/c/5b3c59535)
-  Add functions [`pg_input_is_valid()`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-validity-table) and `pg_input_error_info()` to check for type conversion errors (Tom Lane) [&sect;](https://postgr.es/c/1939d2628) [&sect;](https://postgr.es/c/b8da37b3a)


#### General Queries { #release-16-general }


-  Allow subqueries in the `FROM` clause to omit aliases (Dean Rasheed) [&sect;](https://postgr.es/c/bcedd8f5f)
-  Add support for enhanced numeric literals in SQL/JSON paths (Peter Eisentraut) [&sect;](https://postgr.es/c/102a5c164)

   For example, allow hexadecimal, octal, and binary integers and underscores between digits.


#### Functions { #release-16-functions }


-  Add SQL/JSON constructors (Nikita Glukhov, Teodor Sigaev, Oleg Bartunov, Alexander Korotkov, Amit Langote) [&sect;](https://postgr.es/c/7081ac46a)

   The new functions [`JSON_ARRAY()`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-json-creation-table), [`JSON_ARRAYAGG()`](../../the-sql-language/functions-and-operators/aggregate-functions.md#functions-aggregate-table), `JSON_OBJECT()`, and `JSON_OBJECTAGG()` are part of the SQL standard.
-  Add SQL/JSON object checks (Nikita Glukhov, Teodor Sigaev, Oleg Bartunov, Alexander Korotkov, Amit Langote, Andrew Dunstan) [&sect;](https://postgr.es/c/6ee30209a)

   The [`IS JSON`](../../the-sql-language/functions-and-operators/json-functions-and-operators.md#functions-sqljson-misc) checks include checks for values, arrays, objects, scalars, and unique keys.
-  Allow JSON string parsing to use vector operations (John Naylor) [&sect;](https://postgr.es/c/0a8de93a4)
-  Improve the handling of full text highlighting function [`ts_headline()`](../../the-sql-language/functions-and-operators/text-search-functions-and-operators.md#textsearch-functions-table) for `OR` and `NOT` expressions (Tom Lane) [&sect;](https://postgr.es/c/5a617d75d)
-  Add functions to add, subtract, and generate `timestamptz` values in a specified time zone (Przemyslaw Sztoch, Gurjeet Singh) [&sect;](https://postgr.es/c/75bd846b6)

   The functions are [`date_add()`](../../the-sql-language/functions-and-operators/date-time-functions-and-operators.md#functions-datetime-table), `date_subtract()`, and [`generate_series()`](../../the-sql-language/functions-and-operators/set-returning-functions.md#functions-srf-series).
-  Change [`date_trunc(unit, timestamptz, time_zone)`](../../the-sql-language/functions-and-operators/date-time-functions-and-operators.md#functions-datetime-table) to be an immutable function (Przemyslaw Sztoch) [&sect;](https://postgr.es/c/533e02e92)

   This allows the creation of expression indexes using this function.
-  Add server variable [`SYSTEM_USER`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-session-table) (Bertrand Drouvot) [&sect;](https://postgr.es/c/0823d061b)

   This reports the authentication method and its authenticated user.
-  Add functions [`array_sample()`](../../the-sql-language/functions-and-operators/array-functions-and-operators.md#array-functions-table) and `array_shuffle()` (Martin Kalcher) [&sect;](https://postgr.es/c/888f2ea0a)
-  Add aggregate function [`ANY_VALUE()`](../../the-sql-language/functions-and-operators/aggregate-functions.md#functions-aggregate-table) which returns any value from a set (Vik Fearing) [&sect;](https://postgr.es/c/2ddab010c)
-  Add function [`random_normal()`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-random-table) to supply normally-distributed random numbers (Paul Ramsey) [&sect;](https://postgr.es/c/38d81760c)
-  Add error function [`erf()`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-func-table) and its complement `erfc()` (Dean Rasheed) [&sect;](https://postgr.es/c/d5d574146)
-  Improve the accuracy of numeric [`power()`](../../the-sql-language/functions-and-operators/mathematical-functions-and-operators.md#functions-math-func-table) for integer exponents (Dean Rasheed) [&sect;](https://postgr.es/c/40c7fcbbe)
-  Add [`XMLSERIALIZE()`](../../the-sql-language/data-types/xml-type.md#datatype-xml-creating) option `INDENT` to pretty-print its output (Jim Jones) [&sect;](https://postgr.es/c/483bdb2af)
-  Change [`pg_collation_actual_version()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-collation) to return a reasonable value for the default collation (Jeff Davis) [&sect;](https://postgr.es/c/10932ed5e)

   Previously it returned `NULL`.
-  Allow [`pg_read_file()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-genfile-table) and `pg_read_binary_file()` to ignore missing files (Kyotaro Horiguchi) [&sect;](https://postgr.es/c/283129e32)
-  Add byte specification (`B`) to [`pg_size_bytes()`](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-dbsize) (Peter Eisentraut) [&sect;](https://postgr.es/c/ce1215d9b)
-  Allow [`to_reg`](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-catalog-table)* functions to accept numeric OIDs as input (Tom Lane) [&sect;](https://postgr.es/c/3ea7329c9)


#### [PL/pgSQL] { #release-16-plpgsql }


-  Add the ability to get the current function's OID in PL/pgSQL (Pavel Stehule) [&sect;](https://postgr.es/c/d3d53f955)

   This is accomplished with [`GET DIAGNOSTICS variable = PG_ROUTINE_OID`](../../server-programming/pl-pgsql-sql-procedural-language/basic-statements.md#plpgsql-statements-diagnostics).


#### [libpq] { #release-16-libpq }


-  Add libpq connection option [`require_auth`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-require-auth) to specify a list of acceptable authentication methods (Jacob Champion) [&sect;](https://postgr.es/c/3a465cc67)

   This can also be used to disallow certain authentication methods.
-  Allow multiple libpq-specified hosts to be randomly selected (Jelte Fennema) [&sect;](https://postgr.es/c/7f5b19817) [&sect;](https://postgr.es/c/0a16512d4)

   This is enabled with [`load_balance_hosts=random`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-load-balance-hosts) and can be used for load balancing.
-  Add libpq option [`sslcertmode`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-sslcertmode) to control transmission of the client certificate (Jacob Champion) [&sect;](https://postgr.es/c/36f40ce2d)

   The option values are `disable`, `allow`, and `require`.
-  Allow libpq to use the system certificate pool for certificate verification (Jacob Champion, Thomas Habets) [&sect;](https://postgr.es/c/8eda73146)

   This is enabled with [`sslrootcert=system`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-sslrootcert), which also enables [`sslmode=verify-full`](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-sslmode).


#### Client Applications { #release-16-client-apps }


-  Allow [`ECPG`](../../client-interfaces/ecpg-embedded-sql-in-c/index.md#ecpg) variable declarations to use typedef names that match unreserved SQL keywords (Tom Lane) [&sect;](https://postgr.es/c/83f1c7b74)

   This change does prevent keywords which match C typedef names from being processed as keywords in later `EXEC SQL` blocks.


##### [app-psql] { #release-16-psql }


-  Allow psql to control the maximum width of header lines in expanded format (Platon Pronko) [&sect;](https://postgr.es/c/a45388d6e)

   This is controlled by [`xheader_width`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-pset-xheader-width).
-  Add psql command [`\drg`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-drg) to show role membership details (Pavel Luzanov) [&sect;](https://postgr.es/c/d913928c9) [&sect;](https://postgr.es/c/d65ddaca9)

   The `Member of` output column has been removed from `\du` and `\dg` because this new command displays this information in more detail.
-  Allow psql's access privilege commands to show system objects (Nathan Bossart) [&sect;](https://postgr.es/c/d913928c9) [&sect;](https://postgr.es/c/d65ddaca9)

   The options are [`\dpS`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-dp-lc) and [`\zS`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-z).
-  Add `FOREIGN` designation to psql [`\d+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-d) for foreign table children and partitions (Ian Lawrence Barwick) [&sect;](https://postgr.es/c/bd95816f7)
-  Prevent [`\df+`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-df-uc) from showing function source code (Isaac Morland) [&sect;](https://postgr.es/c/3dfae91f7)

   Function bodies are more easily viewed with [`\sf`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-sf).
-  Allow psql to submit queries using the extended query protocol (Peter Eisentraut) [&sect;](https://postgr.es/c/5b66de343)

   Passing arguments to such queries is done using the new psql [`\bind`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-bind) command.
-  Allow psql [`\watch`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-watch) to limit the number of executions (Andrey Borodin) [&sect;](https://postgr.es/c/00beecfe8)

   The `\watch` options can now be named when specified.
-  Detect invalid values for psql [`\watch`](../../reference/postgresql-client-applications/psql.md#app-psql-meta-command-watch), and allow zero to specify no delay (Andrey Borodin) [&sect;](https://postgr.es/c/6f9ee74d4)
-  Allow psql scripts to obtain the exit status of shell commands and queries (Corey Huinker, Tom Lane) [&sect;](https://postgr.es/c/b0d8f2d98) [&sect;](https://postgr.es/c/31ae2aa9d)

   The new psql control variables are [`SHELL_ERROR`](../../reference/postgresql-client-applications/psql.md#app-psql-variables-shell-error) and [`SHELL_EXIT_CODE`](../../reference/postgresql-client-applications/psql.md#app-psql-variables-shell-exit-code).
-  Various psql tab completion improvements (Vignesh C, Aleksander Alekseev, Dagfinn Ilmari Mannsåker, Shi Yu, Michael Paquier, Ken Kato, Peter Smith) [&sect;](https://postgr.es/c/f6c750d31) [&sect;](https://postgr.es/c/4cbe57974) [&sect;](https://postgr.es/c/6afcab6ac) [&sect;](https://postgr.es/c/9aa58d48f) [&sect;](https://postgr.es/c/3cf2f7af7) [&sect;](https://postgr.es/c/2ea5de296) [&sect;](https://postgr.es/c/07f7237c2) [&sect;](https://postgr.es/c/9d0cf5749) [&sect;](https://postgr.es/c/a3bc631ea) [&sect;](https://postgr.es/c/2ff5ca86e) [&sect;](https://postgr.es/c/9e1e9d656) [&sect;](https://postgr.es/c/96c498d2f)


##### [pg_dump] { #release-16-pgdump }


-  Add pg_dump control of dumping child tables and partitions (Gilles Darold) [&sect;](https://postgr.es/c/a563c24c9)

   The new options are `--table-and-children`, `--exclude-table-and-children`, and `--exclude-table-data-and-children`.
-  Add LZ4 and Zstandard compression to pg_dump (Georgios Kokolatos, Justin Pryzby)
-  Allow pg_dump and [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) to use `long` mode for compression (Justin Pryzby) [&sect;](https://postgr.es/c/0da243fed) [&sect;](https://postgr.es/c/0070b66fe) [&sect;](https://postgr.es/c/84adc8e20) [&sect;](https://postgr.es/c/2820adf77)
-  Improve pg_dump to accept a more consistent compression syntax (Georgios Kokolatos) [&sect;](https://postgr.es/c/5e73a6048)

   Options like `--compress=gzip:5`.


#### Server Applications { #release-16-server-apps }


-  Add [initdb](../../reference/postgresql-server-applications/initdb.md#app-initdb) option to set server variables for the duration of initdb and all future server starts (Tom Lane) [&sect;](https://postgr.es/c/3e51b278d)

   The option is `-c name=value`.
-  Add options to [createuser](../../reference/postgresql-client-applications/createuser.md#app-createuser) to control more user options (Shinya Kato) [&sect;](https://postgr.es/c/08951a7c9) [&sect;](https://postgr.es/c/2dcd1578c)

   Specifically, the new options control the valid-until date, bypassing of row-level security, and role membership.
-  Deprecate [createuser](../../reference/postgresql-client-applications/createuser.md#app-createuser) option `--role` (Nathan Bossart) [&sect;](https://postgr.es/c/2dcd1578c) [&sect;](https://postgr.es/c/381d19b3e)

   This option could be easily confused with new createuser role membership options, so option `--member-of` has been added with the same functionality. The `--role` option can still be used.
-  Allow control of [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) schema processing (Gilles Darold) [&sect;](https://postgr.es/c/7781f4e3e)

   These are controlled by options `--schema` and `--exclude-schema`.
-  Use new [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) options to improve the performance of [vacuumdb](../../reference/postgresql-client-applications/vacuumdb.md#app-vacuumdb) (Tom Lane, Nathan Bossart) [&sect;](https://postgr.es/c/a46a7011b)
-  Have [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) set the new cluster's locale and encoding (Jeff Davis) [&sect;](https://postgr.es/c/9637badd9)

   This removes the requirement that the new cluster be created with the same locale and encoding settings.
-  Add [pg_upgrade](../../reference/postgresql-server-applications/pg_upgrade.md#pgupgrade) option to specify the default transfer mode (Peter Eisentraut) [&sect;](https://postgr.es/c/746915c68)

   The option is `--copy`.
-  Improve [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) to accept numeric compression options (Georgios Kokolatos, Michael Paquier) [&sect;](https://postgr.es/c/d18655cc0)

   Options like `--compress=server-5` are now supported.
-  Fix [pg_basebackup](../../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup) to handle tablespaces stored in the `PGDATA` directory (Robert Haas) [&sect;](https://postgr.es/c/363e8f911)
-  Add [pg_waldump](../../reference/postgresql-server-applications/pg_waldump.md#pgwaldump) option `--save-fullpage` to dump full page images (David Christensen) [&sect;](https://postgr.es/c/d497093cb)
-  Allow [pg_waldump](../../reference/postgresql-server-applications/pg_waldump.md#pgwaldump) options `-t`/`--timeline` to accept hexadecimal values (Peter Eisentraut) [&sect;](https://postgr.es/c/4c8044c04)
-  Add support for progress reporting to [pg_verifybackup](../../reference/postgresql-client-applications/pg_verifybackup.md#app-pgverifybackup) (Masahiko Sawada) [&sect;](https://postgr.es/c/d07c2948b)
-  Allow [pg_rewind](../../reference/postgresql-server-applications/pg_rewind.md#app-pgrewind) to properly track timeline changes (Heikki Linnakangas) [&sect;](https://postgr.es/c/009eeee74) [&sect;](https://postgr.es/c/0a0500207)

   Previously if pg_rewind was run after a timeline switch but before a checkpoint was issued, it might incorrectly determine that a rewind was unnecessary.
-  Have [pg_receivewal](../../reference/postgresql-client-applications/pg_receivewal.md#app-pgreceivewal) and [pg_recvlogical](../../reference/postgresql-client-applications/pg_recvlogical.md#app-pgrecvlogical) cleanly exit on `SIGTERM` (Christoph Berg) [&sect;](https://postgr.es/c/8b60db774)

   This signal is often used by systemd.


#### Source Code { #release-16-source-code }


-  Build ICU support by default (Jeff Davis) [&sect;](https://postgr.es/c/fcb21b3ac)

   This removes [build flag](../../server-administration/installation-from-source-code/index.md#installation) `--with-icu` and adds flag `--without-icu`.
-  Add support for SSE2 (Streaming SIMD Extensions 2) vector operations on x86-64 architectures (John Naylor) [&sect;](https://postgr.es/c/56f2c7b58)
-  Add support for Advanced SIMD (Single Instruction Multiple Data) (NEON) instructions on ARM architectures (Nathan Bossart) [&sect;](https://postgr.es/c/82739d4a8)
-  Have `Windows` binaries built with MSVC use `RandomizedBaseAddress` (ASLR) (Michael Paquier) [&sect;](https://postgr.es/c/36389a060)

   This was already enabled on MinGW builds.
-  Prevent extension libraries from exporting their symbols by default (Andres Freund, Tom Lane) [&sect;](https://postgr.es/c/089480c07) [&sect;](https://postgr.es/c/8cf64d35e)

   Functions that need to be called from the core backend or other extensions must now be explicitly marked `PGDLLEXPORT`.
-  Require `Windows 10` or newer versions (Michael Paquier, Juan José Santamaría Flecha) [&sect;](https://postgr.es/c/495ed0ef2)

   Previously `Windows Vista` and `Windows XP` were supported.
-  Require Perl version 5.14 or later (John Naylor) [&sect;](https://postgr.es/c/4c1532763)
-  Require Bison version 2.3 or later (John Naylor) [&sect;](https://postgr.es/c/b086a47a2)
-  Require Flex version 2.5.35 or later (John Naylor) [&sect;](https://postgr.es/c/8b878bffa)
-  Require MIT Kerberos for GSSAPI support (Stephen Frost) [&sect;](https://postgr.es/c/f7431bca8)
-  Remove support for Visual Studio 2013 (Michael Paquier) [&sect;](https://postgr.es/c/6203583b7)
-  Remove support for `HP-UX` (Thomas Munro) [&sect;](https://postgr.es/c/9db300ce6)
-  Remove support for HP/Intel Itanium (Thomas Munro) [&sect;](https://postgr.es/c/0ad5b48e5)
-  Remove support for M68K, M88K, M32R, and SuperH CPU architectures (Thomas Munro) [&sect;](https://postgr.es/c/718aa43a4) [&sect;](https://postgr.es/c/14168d3c6)
-  Remove [libpq](../../client-interfaces/libpq-c-library/index.md#libpq) support for SCM credential authentication (Michael Paquier) [&sect;](https://postgr.es/c/98ae2c84a)

   Backend support for this authentication method was removed in PostgresSQL 9.1.
-  Add [meson](../../server-administration/installation-from-source-code/building-and-installation-with-meson.md#install-meson) build system (Andres Freund, Nazir Bilal Yavuz, Peter Eisentraut) [&sect;](https://postgr.es/c/e6927270c)

   This eventually will replace the Autoconf and `Windows`-based MSVC build systems.
-  Allow control of the location of the openssl binary used by the build system (Peter Eisentraut) [&sect;](https://postgr.es/c/c8e4030d1)

   Make finding openssl program a configure or meson option
-  Add build option to allow testing of small table segment sizes (Andres Freund) [&sect;](https://postgr.es/c/d3b111e32)

   The build options are [`--with-segsize-blocks`](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-option-with-segsize) and `-Dsegsize_blocks`.
-  Add [pgindent](../../internals/postgresql-coding-conventions/index.md#source) options (Andrew Dunstan) [&sect;](https://postgr.es/c/b90f0b574) [&sect;](https://postgr.es/c/62e1e28bf) [&sect;](https://postgr.es/c/124937163) [&sect;](https://postgr.es/c/a1c4cd6f2) [&sect;](https://postgr.es/c/068a243b7) [&sect;](https://postgr.es/c/dab07e8c6) [&sect;](https://postgr.es/c/b16259b3c)

   The new options are `--show-diff`, `--silent-diff`, `--commit`, and `--help`, and allow multiple `--exclude` options. Also require the typedef file to be explicitly specified. Options `--code-base` and `--build` were also removed.
-  Add [pg_bsd_indent](../../internals/postgresql-coding-conventions/index.md#source) source code to the main tree (Tom Lane) [&sect;](https://postgr.es/c/4e831f4ce)
-  Improve make_ctags and make_etags (Yugo Nagata) [&sect;](https://postgr.es/c/d1e2a380c)
-  Adjust [`pg_attribute`](../../internals/system-catalogs/pg_attribute.md#catalog-pg-attribute) columns for efficiency (Peter Eisentraut) [&sect;](https://postgr.es/c/90189eefc)


#### Additional Modules { #release-16-modules }


-  Improve use of extension-based indexes on boolean columns (Zongliang Quan, Tom Lane) [&sect;](https://postgr.es/c/ff720a597)
-  Add support for Daitch-Mokotoff Soundex to [fuzzystrmatch](../additional-supplied-modules-and-extensions/fuzzystrmatch-determine-string-similarities-and-distance.md#fuzzystrmatch) (Dag Lem) [&sect;](https://postgr.es/c/a290378a3)
-  Allow [auto_explain](../additional-supplied-modules-and-extensions/auto_explain-log-execution-plans-of-slow-queries.md#auto-explain) to log values passed to parameterized statements (Dagfinn Ilmari Mannsåker) [&sect;](https://postgr.es/c/d4bfe4128)

   This affects queries using server-side [`PREPARE`](../../reference/sql-commands/prepare.md#sql-prepare)/[`EXECUTE`](../../reference/sql-commands/execute.md#sql-execute) and client-side parse/bind. Logging is controlled by [`auto_explain.log_parameter_max_length`](../additional-supplied-modules-and-extensions/auto_explain-log-execution-plans-of-slow-queries.md#auto-explain-configuration-parameters-log-parameter-max-length); by default query parameters will be logged with no length restriction.
-  Have [auto_explain](../additional-supplied-modules-and-extensions/auto_explain-log-execution-plans-of-slow-queries.md#auto-explain)'s `log_verbose` mode honor the value of [`compute_query_id`](../../server-administration/server-configuration/run-time-statistics.md#guc-compute-query-id) (Atsushi Torikoshi) [&sect;](https://postgr.es/c/9d2d9728b)

   Previously even if `compute_query_id` was enabled, [`log_verbose`](../additional-supplied-modules-and-extensions/auto_explain-log-execution-plans-of-slow-queries.md#auto-explain-configuration-parameters-log-verbose) was not showing the query identifier.
-  Change the maximum length of [ltree](../additional-supplied-modules-and-extensions/ltree-hierarchical-tree-like-data-type.md#ltree) labels from 256 to 1000 and allow hyphens (Garen Torikian) [&sect;](https://postgr.es/c/b1665bf01)
-  Have [`pg_stat_statements`](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements) normalize constants used in utility commands (Michael Paquier) [&sect;](https://postgr.es/c/daa8365a9)

   Previously constants appeared instead of placeholders, e.g., `$1`.
-  Add [pg_walinspect](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect) function [`pg_get_wal_block_info()`](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect-funcs-pg-get-wal-block-info) to report WAL block information (Michael Paquier, Melanie Plageman, Bharath Rupireddy) [&sect;](https://postgr.es/c/c31cf1c03) [&sect;](https://postgr.es/c/9ecb134a9) [&sect;](https://postgr.es/c/122376f02) [&sect;](https://postgr.es/c/df4f3ab51)
-  Change how [pg_walinspect](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect) functions [`pg_get_wal_records_info()`](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect-funcs-pg-get-wal-records-info) and [`pg_get_wal_stats()`](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect-funcs-pg-get-wal-stats) interpret ending LSNs (Bharath Rupireddy) [&sect;](https://postgr.es/c/5c1b66280)

   Previously ending LSNs which represent nonexistent WAL locations would generate an error, while they will now be interpreted as the end of the WAL.
-  Add detailed descriptions of WAL records in [pg_walinspect](../additional-supplied-modules-and-extensions/pg_walinspect-low-level-wal-inspection.md#pgwalinspect) and [pg_waldump](../../reference/postgresql-server-applications/pg_waldump.md#pgwaldump) (Melanie Plageman, Peter Geoghegan) [&sect;](https://postgr.es/c/7d8219a44) [&sect;](https://postgr.es/c/1c453cfd8) [&sect;](https://postgr.es/c/96149a180) [&sect;](https://postgr.es/c/50547a3fa)
-  Add [pageinspect](../additional-supplied-modules-and-extensions/pageinspect-low-level-inspection-of-database-pages.md#pageinspect) function [`bt_multi_page_stats()`](../additional-supplied-modules-and-extensions/pageinspect-low-level-inspection-of-database-pages.md#pageinspect-b-tree-funcs) to report statistics on multiple pages (Hamid Akhtar) [&sect;](https://postgr.es/c/1fd3dd204)

   This is similar to `bt_page_stats()` except it can report on a range of pages.
-  Add empty range output column to [pageinspect](../additional-supplied-modules-and-extensions/pageinspect-low-level-inspection-of-database-pages.md#pageinspect) function [`brin_page_items()`](../additional-supplied-modules-and-extensions/pageinspect-low-level-inspection-of-database-pages.md#pageinspect-brin-funcs) (Tomas Vondra) [&sect;](https://postgr.es/c/1fd3dd204)
-  Redesign archive modules to be more flexible (Nathan Bossart) [&sect;](https://postgr.es/c/35739b87d)

   Initialization changes will require modules written for older versions of Postgres to be updated.
-  Correct inaccurate [pg_stat_statements](../additional-supplied-modules-and-extensions/pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements) row tracking extended query protocol statements (Sami Imseih) [&sect;](https://postgr.es/c/1d477a907)
-  Add [pg_buffercache](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache) function `pg_buffercache_usage_counts()` to report usage totals (Nathan Bossart) [&sect;](https://postgr.es/c/f3fa31327)
-  Add [pg_buffercache](../additional-supplied-modules-and-extensions/pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache) function `pg_buffercache_summary()` to report summarized buffer statistics (Melih Mutlu) [&sect;](https://postgr.es/c/2589434ae)
-  Allow the schemas of required extensions to be referenced in extension scripts using the new syntax `@extschema:referenced_extension_name@` (Regina Obe) [&sect;](https://postgr.es/c/72a5b1fc8)
-  Allow required extensions to be marked as non-relocatable using [`no_relocate`](../../server-programming/extending-sql/packaging-related-objects-into-an-extension.md#extend-extensions-files-no-relocate) (Regina Obe) [&sect;](https://postgr.es/c/72a5b1fc8)

   This allows `@extschema:referenced_extension_name@` to be treated as a constant for the lifetime of the extension.


##### [postgres_fdw] { #release-16-pgfdw }


-  Allow postgres_fdw to do aborts in parallel (Etsuro Fujita) [&sect;](https://postgr.es/c/983ec2300)

   This is enabled with postgres_fdw option [`parallel_abort`](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-options-transaction-management).
-  Make [`ANALYZE`](../../reference/sql-commands/analyze.md#sql-analyze) on foreign postgres_fdw tables more efficient (Tomas Vondra) [&sect;](https://postgr.es/c/8ad51b5f4)

   The postgres_fdw option [`analyze_sampling`](../additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-options-cost-estimation) controls the sampling method.
-  Restrict shipment of [`reg`](../../the-sql-language/data-types/object-identifier-types.md#datatype-oid)* type constants in postgres_fdw to those referencing built-in objects or extensions marked as shippable (Tom Lane) [&sect;](https://postgr.es/c/31e5b5029)
-  Have postgres_fdw and [dblink](../additional-supplied-modules-and-extensions/dblink-connect-to-other-postgresql-databases.md#dblink) handle interrupts during connection establishment (Andres Freund) [&sect;](https://postgr.es/c/e4602483e)


### Acknowledgments { #release-16-acknowledgements }


 The following individuals (in alphabetical order) have contributed to this release as patch authors, committers, reviewers, testers, or reporters of issues.


- Abhijit Menon-Sen
- Adam Mackler
- Adrian Klaver
- Ahsan Hadi
- Ajin Cherian
- Ajit Awekar
- Alan Hodgson
- Aleksander Alekseev
- Alex Denman
- Alex Kozhemyakin
- Alexander Korolev
- Alexander Korotkov
- Alexander Lakhin
- Alexander Pyhalov
- Alexey Borzov
- Alexey Ermakov
- Alexey Makhmutov
- Álvaro Herrera
- Amit Kapila
- Amit Khandekar
- Amit Langote
- Amul Sul
- Anastasia Lubennikova
- Anban Company
- Andreas Dijkman
- Andreas Karlsson
- Andreas Scherbaum
- Andrei Zubkov
- Andres Freund
- Andrew Alsup
- Andrew Bille
- Andrew Dunstan
- Andrew Gierth
- Andrew Kesper
- Andrey Borodin
- Andrey Lepikhov
- Andrey Sokolov
- Ankit Kumar Pandey
- Ante Kresic
- Anton Melnikov
- Anton Sidyakin
- Anton Voloshin
- Antonin Houska
- Arne Roland
- Artem Anisimov
- Arthur Zakirov
- Ashutosh Bapat
- Ashutosh Sharma
- Asim Praveen
- Atsushi Torikoshi
- Ayaki Tachikake
- Balazs Szilfai
- Benoit Lobréau
- Bernd Helmle
- Bertrand Drouvot
- Bharath Rupireddy
- Bilva Sanaba
- Bob Krier
- Boris Zentner
- Brad Nicholson
- Brar Piening
- Bruce Momjian
- Bruno da Silva
- Carl Sopchak
- Cary Huang
- Changhong Fei
- Chris Travers
- Christoph Berg
- Christophe Pettus
- Corey Huinker
- Craig Ringer
- Curt Kolovson
- Dag Lem
- Dagfinn Ilmari Mannsåker
- Daniel Gustafsson
- Daniel Vérité
- Daniel Watzinger
- Daniel Westermann
- Daniele Varrazzo
- Daniil Anisimov
- Danny Shemesh
- Dave Page
- David Christensen
- David G. Johnston
- David Geier
- David Gilman
- David Kimura
- David Rowley
- David Steele
- David Turon
- David Zhang
- Davinder Singh
- Dean Rasheed
- Denis Laxalde
- Dilip Kumar
- Dimos Stamatakis
- Dmitriy Kuzmin
- Dmitry Astapov
- Dmitry Dolgov
- Dmitry Koval
- Dong Wook Lee
- Dongming Liu
- Drew DeVault
- Duncan Sands
- Ed Maste
- Egor Chindyaskin
- Ekaterina Kiryanova
- Elena Indrupskaya
- Emmanuel Quincerot
- Eric Mutta
- Erik Rijkers
- Erki Eessaar
- Erwin Brandstetter
- Etsuro Fujita
- Eugeny Zhuzhnev
- Euler Taveira
- Evan Jones
- Evgeny Morozov
- Fabrízio de Royes Mello
- Farias de Oliveira
- Florin Irion
- Franz-Josef Färber
- Garen Torikian
- Georgios Kokolatos
- Gilles Darold
- Greg Stark
- Guillaume Lelarge
- Gunnar Bluth
- Gunnar Morling
- Gurjeet Singh
- Haiyang Wang
- Haiying Tang
- Hamid Akhtar
- Hans Buschmann
- Hao Wu
- Hayato Kuroda
- Heath Lord
- Heikki Linnakangas
- Himanshu Upadhyaya
- Hisahiro Kauchi
- Hongyu Song
- Hubert Lubaczewski
- Hung Nguyen
- Ian Barwick
- Ibrar Ahmed
- Ilya Gladyshev
- Ilya Nenashev
- Isaac Morland
- Israel Barth Rubio
- Jacob Champion
- Jacob Speidel
- Jaime Casanova
- Jakub Wartak
- James Coleman
- James Inform
- James Vanns
- Jan Wieck
- Japin Li
- Jeevan Ladhe
- Jeff Davis
- Jeff Janes
- Jehan-Guillaume de Rorthais
- Jelte Fennema
- Jian He
- Jim Jones
- Jinbao Chen
- Joe Conway
- Joel Jacobson
- John Naylor
- Jonathan Katz
- Josef Simanek
- Joseph Koshakow
- Juan José Santamaría Flecha
- Julien Rouhaud
- Julien Roze
- Junwang Zhao
- Justin Pryzby
- Justin Zhang
- Karina Litskevich
- Karl O. Pinc
- Keisuke Kuroda
- Ken Kato
- Kevin McKibbin
- Kieran McCusker
- Kirk Wolak
- Konstantin Knizhnik
- Koshi Shibagaki
- Kotaro Kawamoto
- Kui Liu
- Kyotaro Horiguchi
- Lakshmi Narayanan Sreethar
- Laurence Parry
- Laurenz Albe
- Luca Ferrari
- Lukas Fittl
- Maciek Sakrejda
- Magnus Hagander
- Maja Zaloznik
- Marcel Hofstetter
- Marina Polyakova
- Mark Dilger
- Marko Tiikkaja
- Markus Winand
- Martijn van Oosterhout
- Martin Jurca
- Martin Kalcher
- Mary Xu
- Masahiko Sawada
- Masahiro Ikeda
- Masao Fujii
- Mason Sharp
- Matheus Alcantara
- Mats Kindahl
- Matthias van de Meent
- Matthijs van der Vleuten
- Maxim Orlov
- Maxim Yablokov
- Mehmet Emin Karakas
- Melanie Plageman
- Melih Mutlu
- Micah Gates
- Michael Banck
- Michael Paquier
- Michail Nikolaev
- Michel Pelletier
- Mike Oh
- Mikhail Gribkov
- Mingli Zhang
- Miroslav Bendik
- Mitsuru Hinata
- Myo Wai Thant
- Naeem Akhter
- Naoki Okano
- Nathan Bossart
- Nazir Bilal Yavuz
- Neha Sharma
- Nick Babadzhanian
- Nicola Contu
- Nikhil Shetty
- Nikita Glukhov
- Nikolay Samokhvalov
- Nikolay Shaplov
- Nishant Sharma
- Nitin Jadhav
- Noah Misch
- Noboru Saito
- Noriyoshi Shinoda
- Nuko Yokohama
- Oleg Bartunov
- Oleg Tselebrovskiy
- Olly Betts
- Onder Kalaci
- Onur Tirtir
- Pablo Federico
- Palle Girgensohn
- Paul Guo
- Paul Jungwirth
- Paul Ramsey
- Pavel Borisov
- Pavel Kulakov
- Pavel Luzanov
- Pavel Stehule
- Peifeng Qiu
- Peter Eisentraut
- Peter Geoghegan
- Peter Smith
- Phil Florent
- Philippe Godfrin
- Platon Pronko
- Przemyslaw Sztoch
- Rachel Heaton
- Ranier Vilela
- Regina Obe
- Reid Thompson
- Reiner Peterke
- Richard Guo
- Riivo Kolka
- Rishu Bagga
- Robert Haas
- Robert Sjöblom
- Robert Treat
- Roberto Mello
- Robins Tharakan
- Roman Zharkov
- Ronan Dunklau
- Rushabh Lathia
- Ryo Matsumura
- Samay Sharma
- Sami Imseih
- Sandeep Thakkar
- Sandro Santilli
- Sebastien Flaesch
- Sébastien Lardière
- Sehrope Sarkuni
- Sergey Belyashov
- Sergey Pankov
- Sergey Shinderuk
- Shi Yu
- Shinya Kato
- Sho Kato
- Shruthi Gowda
- Shveta Mallik
- Simon Riggs
- Sindy Senorita
- Sirisha Chamarthi
- Sravan Kumar
- Stéphane Tachoires
- Stephen Frost
- Steve Chavez
- Stone Tickle
- Sven Klemm
- Takamichi Osumi
- Takeshi Ideriha
- Tatsuhiro Nakamori
- Tatsuo Ishii
- Teja Mupparti
- Tender Wang
- Teodor Sigaev
- Thiago Nunes
- Thom Brown
- Thomas Habets
- Thomas Mc Kay
- Thomas Munro
- Tim Carey-Smith
- Tim Field
- Timo Stolz
- Tom Lane
- Tomas Vondra
- Tor Erik Linnerud
- Torsten Förtsch
- Tristan Partin
- Troy Frericks
- Tushar Ahuja
- Valerie Woolard
- Vibhor Kumar
- Victor Spirin
- Victoria Shepard
- Vignesh C
- Vik Fearing
- Vitaly Burovoy
- Vitaly Davydov
- Wang Wei
- Wenjing Zeng
- Whale Song
- Will Mortensen
- Wolfgang Walther
- Xin Wen
- Xing Guo
- Xingwang Xu
- XueJing Zhao
- Yanliang Lei
- Youmiu Mo
- Yugo Nagata
- Yura Sokolov
- Yuta Katsuragi
- Zhen Mingyang
- Zheng Li
- Zhihong Yu
- Zhijie Hou
- Zongliang Quan
- Zuming Jiang
