## Release 17.8 { #release-17-8 }


**Release date:.**


2026-02-12


 This release contains a variety of fixes from 17.7. For information about new features in major release 17, see [Release 17](release-17.md#release-17).


### Migration to Version 17.8 { #release-17-8-migration }


 A dump/restore is not required for those running 17.X.


 However, if you are upgrading from a version earlier than 17.6, see [Release 17.6](release-17-6.md#release-17-6).


### Changes { #release-17-8-changes }


-  Guard against unexpected dimensions of `oidvector`/`int2vector` (Tom Lane) [&sect;](https://postgr.es/c/3d160401b)

   These data types are expected to be 1-dimensional arrays containing no nulls, but there are cast pathways that permit violating those expectations. Add checks to some functions that were depending on those expectations without verifying them, and could misbehave in consequence.

   The PostgreSQL Project thanks Altan Birler for reporting this problem. (CVE-2026-2003)
-  Harden selectivity estimators against being attached to operators that accept unexpected data types (Tom Lane) [&sect;](https://postgr.es/c/bbf5bcf58) [&sect;](https://postgr.es/c/dd3ad2a4d) [&sect;](https://postgr.es/c/dbb09fd8e)

   `contrib/intarray` contained a selectivity estimation function that could be abused for arbitrary code execution, because it did not check that its input was of the expected data type. Third-party extensions should check for similar hazards and add defenses using the technique intarray now uses. Since such extension fixes will take time, we now require superuser privilege to attach a non-built-in selectivity estimator to an operator.

   The PostgreSQL Project thanks Daniel Firer, as part of zeroday.cloud, for reporting this problem. (CVE-2026-2004)
-  Fix buffer overrun in `contrib/pgcrypto`'s PGP decryption functions (Michael Paquier) [&sect;](https://postgr.es/c/7a7d9693c)

   Decrypting a crafted message with an overlength session key caused a buffer overrun, with consequences as bad as arbitrary code execution.

   The PostgreSQL Project thanks Team Xint Code, as part of zeroday.cloud, for reporting this problem. (CVE-2026-2005)
-  Fix inadequate validation of multibyte character lengths (Thomas Munro, Noah Misch) [&sect;](https://postgr.es/c/838248b1b) [&sect;](https://postgr.es/c/7a522039f) [&sect;](https://postgr.es/c/319e8a644) [&sect;](https://postgr.es/c/10ebc4bd6) [&sect;](https://postgr.es/c/dc072a09a) [&sect;](https://postgr.es/c/955433ebd)

   Assorted bugs allowed an attacker able to issue crafted SQL to overrun string buffers, with consequences as bad as arbitrary code execution. After these fixes, applications may observe “invalid byte sequence for encoding” errors when string functions process invalid text that has been stored in the database.

   The PostgreSQL Project thanks Paul Gerste and Moritz Sanft, as part of zeroday.cloud, for reporting this problem. (CVE-2026-2006)
-  Don't allow CTE references in sub-selects to determine semantic levels of aggregate functions (Tom Lane) [&sect;](https://postgr.es/c/075a763e2)

   This change undoes a change made two minor releases ago, instead throwing an error if a sub-select references a CTE that's below the semantic level that standard SQL rules would assign to the aggregate based on contained column references and aggregates. The attempted fix turned out to cause problems of its own, and it's unclear what to do instead. Since sub-selects within aggregates are disallowed altogether by the SQL standard, treating such cases as errors seems sufficient.
-  Fix trigger transition table capture for `MERGE` in CTE queries (Dean Rasheed) [&sect;](https://postgr.es/c/c5fc17dda)

   When executing a data-modifying CTE query containing both a `MERGE` and another DML operation on a table with statement-level `AFTER` triggers, the transition tables passed to the triggers would not include the rows affected by the `MERGE`, only those affected by the other operation(s).
-  Fix failure when all children of a partitioned target table of an update or delete have been pruned (Amit Langote) [&sect;](https://postgr.es/c/933f67fb6)

   In such cases, the executor could report “could not find junk ctid column” errors, even though nothing needs to be done.
-  Avoid possible planner failure when a query contains duplicate window function calls (Meng Zhang, David Rowley) [&sect;](https://postgr.es/c/cae812741)

   Confusion over de-duplication of such calls could result in errors like “WindowFunc with winref 2 assigned to WindowAgg with winref 1”.
-  Allow indexscans on partial hash indexes even when the index's predicate implies the truth of the WHERE clause (Tom Lane) [&sect;](https://postgr.es/c/e79b27662)

   Normally we drop a WHERE clause that is implied by the predicate, since it's pointless to test it; it must hold for every index entry. However that can prevent creation of an indexscan plan if the index is one that requires a WHERE clause on the leading index key, as hash indexes do. Don't drop implied clauses when considering such an index.
-  Do not emit WAL for unlogged BRIN indexes (Kirill Reshke) [&sect;](https://postgr.es/c/4b6d096a0)

   One seldom-taken code path incorrectly emitted a WAL record relating to a BRIN index even if the index was marked unlogged. Crash recovery would then fail to replay that record, complaining that the file already exists.
-  Prevent truncation of CLOG that is still needed by unread `NOTIFY` messages (Joel Jacobson, Heikki Linnakangas) [&sect;](https://postgr.es/c/d02c03ddc) [&sect;](https://postgr.es/c/c2682810a) [&sect;](https://postgr.es/c/d80d5f099)

   This fix prevents “could not access status of transaction” errors when a backend is slow to absorb `NOTIFY` messages.
-  Escalate errors occurring during `NOTIFY` message processing to FATAL, i.e. close the connection (Heikki Linnakangas) [&sect;](https://postgr.es/c/b821c9292)

   Formerly, if a backend got an error while absorbing a `NOTIFY` message, it would advance past that message, report the error to the client, and move on. That behavior was fraught with problems though. One big concern is that the client has no good way to know that a notification was lost, and certainly no way to know what was in it. Depending on the application logic, missing a notification could cause the application to get stuck waiting, for example. Also, any remaining messages would not get processed until someone sent a new `NOTIFY`.

   Also, if the connection is idle at the time of receiving a `NOTIFY` signal, any ERROR would be escalated to FATAL anyway, due to unrelated concerns. Therefore, we've chosen to make that happen in all cases, for consistency and to provide a clear signal to the application that it might have missed some notifications.
-  Fix erroneous counting of updates in `EXPLAIN ANALYZE MERGE` with a concurrent update (Dean Rasheed) [&sect;](https://postgr.es/c/d6c415c4b)

   This situation led to an incorrect count of “skipped” tuples in `EXPLAIN`'s output, or to an assertion failure in an assert-enabled build.
-  Fix bug in following update chain when locking a tuple (Jasper Smit) [&sect;](https://postgr.es/c/bb87d7fef)

   This code path neglected to check the xmin of the first new tuple in the update chain, making it possible to lock an unrelated tuple if the original updater aborted and the space was immediately reclaimed by `VACUUM` and then re-used. That could cause unexpected transaction delays or deadlocks. Errors associated with having identified the wrong tuple have also been observed.
-  Fix issues around in-place catalog updates (Noah Misch) [&sect;](https://postgr.es/c/0f69bedde) [&sect;](https://postgr.es/c/d3e5d8950) [&sect;](https://postgr.es/c/bcb784e7d)

   Send a nontransactional invalidation message for an in-place update, since such an update will survive transaction rollback. Also ensure that the update is WAL-logged before other sessions can see it. These fixes primarily prevent scenarios in which relations' frozen-XID attributes become inconsistent, possibly allowing premature CLOG truncation and subsequent “could not access status of transaction” errors.
-  Fix incorrect handling of incremental backups of large tables (Robert Haas, Oleg Tkachenko) [&sect;](https://postgr.es/c/ad569b54a)

   If a table exceeding 1GB (or in general, the installation's segment size) is truncated by `VACUUM` between the base backup and the incremental backup, pg_combinebackup could fail with an error about “truncation block length in excess of segment size”. This prevented restoring the incremental backup.
-  Fix potential backend process crash at process exit due to trying to release a lock in an already-unmapped shared memory segment (Rahila Syed) [&sect;](https://postgr.es/c/4071fe900)
-  Guard against incorrect truncation of the multixact log after a crash (Heikki Linnakangas) [&sect;](https://postgr.es/c/d3ad4cef6)
-  Fix possibly mis-encoded result of `pg_stat_get_backend_activity()` (Chao Li) [&sect;](https://postgr.es/c/52b27f585)

   The shared-memory buffer holding a session's activity string can end with an incomplete multibyte character. Readers are supposed to truncate off any such incomplete character, but this function failed to do so.
-  Guard against recursive memory context logging (Fujii Masao) [&sect;](https://postgr.es/c/699293d27)

   A constant flow of signals requesting memory context logging could cause recursive execution of the logging code, which in theory could lead to stack overflow.
-  Fix memory context usage when reinitializing a parallel execution context (Jakub Wartak, Jeevan Chalke) [&sect;](https://postgr.es/c/1d0fc2499)

   This error could result in a crash due to a subsidiary data structure having a shorter lifespan than the parallel context. The problem is not known to be reachable using only core PostgreSQL, but we have reports of trouble in extensions.
-  Set next multixid's offset when creating a new multixid, to remove the wait loop that was needed in corner cases (Andrey Borodin) [&sect;](https://postgr.es/c/8ba61bc06) [&sect;](https://postgr.es/c/cad40cec2)

   The previous logic could get stuck waiting for an update that would never occur.
-  Avoid rewriting data-modifying CTEs more than once (Bernice Southey, Dean Rasheed) [&sect;](https://postgr.es/c/c09096503)

   Formerly, when updating an auto-updatable view or a relation with rules, if the original query had any data-modifying CTEs, the rewriter would rewrite those CTEs multiple times due to recursion. This was inefficient and could produce false errors if a CTE included an update of an always-generated column.
-  Allow retrying initialization of a DSM registry entry (Nathan Bossart) [&sect;](https://postgr.es/c/2fc5c5062)

   If we fail partway through initialization of a dynamic shared memory entry, allow the next attempt to use that entry to retry initialization. Previously the entry was left in a permanently-failed state.
-  Fail recovery if WAL does not exist back to the redo point indicated by the checkpoint record (Nitin Jadhav) [&sect;](https://postgr.es/c/f5927da4f)

   Add an explicit check for this before starting recovery, so that no harm is done and a useful error message is provided. Previously, recovery might crash or corrupt the database in this situation.
-  Avoid scribbling on the source query tree during `ALTER PUBLICATION` (Sunil S) [&sect;](https://postgr.es/c/bb08ac7ac)

   This error had the visible effect that an event trigger fired for the query would see only the first `publish` option, even if several had been specified. If such a query were set up as a prepared statement, re-executions would misbehave too.
-  Pass connection options specified in `CREATE SUBSCRIPTION ... CONNECTION` to the publisher's walsender (Fujii Masao) [&sect;](https://postgr.es/c/7a990e801)

   Before this fix, the `options` connection option (if any) was ignored, thus for example preventing setting custom server parameter values in the walsender session. It was intended for that to work, and it did work before refactoring in PostgreSQL version 15 broke it, so restore the previous behavior.
-  Prevent invalidation of newly created or newly synced replication slots (Zhijie Hou) [&sect;](https://postgr.es/c/3243c0177) [&sect;](https://postgr.es/c/9649f1adf) [&sect;](https://postgr.es/c/3510ebeb0)

   A race condition with a concurrent checkpoint could allow WAL to be removed that is needed by the replication slot, causing the slot to immediately get marked invalid.
-  Fix race condition in computing a replication slot's required xmin (Zhijie Hou) [&sect;](https://postgr.es/c/123b851ab)

   This could lead to the error “cannot build an initial slot snapshot as oldest safe xid follows snapshot's xmin”.
-  During initial synchronization of a logical replication subscription, commit the addition of a `pg_replication_origin` entry before starting to copy data (Zhijie Hou) [&sect;](https://postgr.es/c/e063ccc72)

   Previously, if the copy step failed, the new `pg_replication_origin` entry would be lost due to transaction rollback. This led to inconsistent state in shared memory.
-  Don't advance logical replication progress after a parallel worker apply failure (Zhijie Hou) [&sect;](https://postgr.es/c/0ed8f1afb)

   The previous behavior allowed transactions to be lost by a subscriber.
-  Fix logical replication slotsync worker processes to handle LOCK_TIMEOUT signals correctly (Zhijie Hou) [&sect;](https://postgr.es/c/f2818868a)

   Previously, timeout signals were effectively ignored.
-  Fix possible failure with “unexpected data beyond EOF” during restart of a streaming replica server (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/c3770181c)
-  Fix error reporting for SQL/JSON path type mismatches (Jian He) [&sect;](https://postgr.es/c/b5511fed5)

   The code could produce a “cache lookup failed for type 0” error instead of the intended complaint about the path expression not being of the right type.
-  Fix erroneous tracking of column position when parsing partition range bounds (myzhen) [&sect;](https://postgr.es/c/84b787ae6)

   This could, for example, lead to the wrong column name being cited in error messages about casting partition bound values to the column's data type.
-  Fix assorted minor errors in error messages (Man Zeng, Tianchen Zhang) [&sect;](https://postgr.es/c/67ad4387b) [&sect;](https://postgr.es/c/263af458e) [&sect;](https://postgr.es/c/5995135f1) [&sect;](https://postgr.es/c/05ef2371a) [&sect;](https://postgr.es/c/5449fd261)

   For example, an error report about mismatched timeline number in a backup manifest showed the starting timeline number where it meant to show the ending timeline number.
-  Fix failure to perform function inlining when doing JIT compilation with LLVM version 17 or later (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/d0bb0e5b3)
-  Adjust our JIT code to work with LLVM 21 (Holger Hoffstätte) [&sect;](https://postgr.es/c/60215eae7)

   The previous coding failed to compile on aarch64 machines.
-  Add new server parameter [file_extend_method](../../server-administration/server-configuration/resource-consumption.md#guc-file-extend-method) to control use of `posix_fallocate()` (Thomas Munro) [&sect;](https://postgr.es/c/4dac22aa1)

   PostgreSQL version 16 and later will use `posix_fallocate()`, if the platform provides it, to extend relation files. However, this has been reported to interact poorly with some file systems: BTRFS compression is disabled by the use of `posix_fallocate()`, and XFS could produce spurious `ENOSPC` errors in older Linux kernel versions. To provide a workaround, introduce this new server parameter. Setting `file_extend_method` to `write_zeros` will cause the server to return to the old method of extending files by writing blocks of zeroes.
-  Honor `open()`'s `O_CLOEXEC` flag on Windows (Bryan Green, Thomas Munro) [&sect;](https://postgr.es/c/f24af0e04) [&sect;](https://postgr.es/c/045185913) [&sect;](https://postgr.es/c/b3c8119e2)

   Make this flag work like it does on POSIX platforms, so that we don't leak file handles into child processes such as `COPY TO/FROM PROGRAM`. While that leakage hasn't caused many problems, it seems undesirable.
-  Fix failure to parse long options on the server command line in Solaris executables built with meson (Tom Lane) [&sect;](https://postgr.es/c/59c2f7efa)
-  Support process title changes on GNU/Hurd (Michael Banck) [&sect;](https://postgr.es/c/d66a922f9)
-  Avoid pg_dump assertion failure in binary-upgrade mode (Vignesh C) [&sect;](https://postgr.es/c/1cdc07ad5)

   Failure to handle subscription-relation objects in the object sorting code triggered an assertion, though there were no serious ill effects in production builds.
-  Fix incorrect error handling in pgbench with multiple `\syncpipeline` commands in pipeline mode (Yugo Nagata) [&sect;](https://postgr.es/c/5bc251b28)

   If multiple `\syncpipeline` commands are encountered after a query error, pgbench would report “failed to exit pipeline mode”, or get an assertion failure in an assert-enabled build.
-  Make pg_resetwal print the updated value when changing OldestXID (Heikki Linnakangas) [&sect;](https://postgr.es/c/f2e0ca0af)

   It already did that for every other variable it can change.
-  Make pg_resetwal allow setting next multixact xid to 0 or next multixact offset to UINT32_MAX (Maxim Orlov) [&sect;](https://postgr.es/c/cb2ef0e92)

   These are valid values, so rejecting them was incorrect. In the worst case, if a pg_upgrade is attempted when exactly at the point of multixact wraparound, the upgrade would fail.
-  In `contrib/amcheck`, use the correct snapshot for btree index parent checks (Mihail Nikalayeu) [&sect;](https://postgr.es/c/ce2f575b7) [&sect;](https://postgr.es/c/e1a327dc4)

   The previous coding caused spurious errors when examining indexes created with `CREATE INDEX CONCURRENTLY`.
-  Fix `contrib/amcheck` to handle “half-dead” btree index pages correctly (Heikki Linnakangas) [&sect;](https://postgr.es/c/e8ae59445)

   `amcheck` expected such a page to have a parent downlink, but it does not, leading to a false error report about “mismatch between parent key and child high key”.
-  Fix `contrib/amcheck` to handle incomplete btree root page splits correctly (Heikki Linnakangas) [&sect;](https://postgr.es/c/5a2d1df00)

   `amcheck` could report a false error about “block is not true root”.
-  Fix edge-case integer overflow in `contrib/intarray`'s selectivity estimator for `@@` (Chao Li) [&sect;](https://postgr.es/c/a5f2dc421)

   This could cause poor selectivity estimates to be produced for cases involving the maximum integer value.
-  Fix multibyte-encoding issue in `contrib/ltree` (Jeff Davis) [&sect;](https://postgr.es/c/b8cfe9dc2)

   The previous coding could pass an incomplete multibyte character to `lower()`, probably resulting in incorrect behavior.
-  Update time zone data files to tzdata release 2025c (Tom Lane) [&sect;](https://postgr.es/c/f87c0b84e)

   The only change is in historical data for pre-1976 timestamps in Baja California.
