## Release 16.12 { #release-16-12 }


**Release date:.**


2026-02-12


 This release contains a variety of fixes from 16.11. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.12 { #release-16-12-migration }


 A dump/restore is not required for those running 16.X.


 However, if you are upgrading from a version earlier than 16.10, see [Release 16.10](release-16-10.md#release-16-10).


### Changes { #release-16-12-changes }


-  Guard against unexpected dimensions of `oidvector`/`int2vector` (Tom Lane) [&sect;](https://postgr.es/c/595956fc7)

   These data types are expected to be 1-dimensional arrays containing no nulls, but there are cast pathways that permit violating those expectations. Add checks to some functions that were depending on those expectations without verifying them, and could misbehave in consequence.

   The PostgreSQL Project thanks Altan Birler for reporting this problem. (CVE-2026-2003)
-  Harden selectivity estimators against being attached to operators that accept unexpected data types (Tom Lane) [&sect;](https://postgr.es/c/91d7c0bfd) [&sect;](https://postgr.es/c/c0887b39d) [&sect;](https://postgr.es/c/d484bc260)

   `contrib/intarray` contained a selectivity estimation function that could be abused for arbitrary code execution, because it did not check that its input was of the expected data type. Third-party extensions should check for similar hazards and add defenses using the technique intarray now uses. Since such extension fixes will take time, we now require superuser privilege to attach a non-built-in selectivity estimator to an operator.

   The PostgreSQL Project thanks Daniel Firer, as part of zeroday.cloud, for reporting this problem. (CVE-2026-2004)
-  Fix buffer overrun in `contrib/pgcrypto`'s PGP decryption functions (Michael Paquier) [&sect;](https://postgr.es/c/527b730f4)

   Decrypting a crafted message with an overlength session key caused a buffer overrun, with consequences as bad as arbitrary code execution.

   The PostgreSQL Project thanks Team Xint Code, as part of zeroday.cloud, for reporting this problem. (CVE-2026-2005)
-  Fix inadequate validation of multibyte character lengths (Thomas Munro, Noah Misch) [&sect;](https://postgr.es/c/70ff9ede5) [&sect;](https://postgr.es/c/b0e3f5cf9) [&sect;](https://postgr.es/c/d837fb029) [&sect;](https://postgr.es/c/4c08960d9) [&sect;](https://postgr.es/c/0c33d5608) [&sect;](https://postgr.es/c/763671b74)

   Assorted bugs allowed an attacker able to issue crafted SQL to overrun string buffers, with consequences as bad as arbitrary code execution. After these fixes, applications may observe “invalid byte sequence for encoding” errors when string functions process invalid text that has been stored in the database.

   The PostgreSQL Project thanks Paul Gerste and Moritz Sanft, as part of zeroday.cloud, for reporting this problem. (CVE-2026-2006)
-  Don't allow CTE references in sub-selects to determine semantic levels of aggregate functions (Tom Lane) [&sect;](https://postgr.es/c/1c8c3206f)

   This change undoes a change made two minor releases ago, instead throwing an error if a sub-select references a CTE that's below the semantic level that standard SQL rules would assign to the aggregate based on contained column references and aggregates. The attempted fix turned out to cause problems of its own, and it's unclear what to do instead. Since sub-selects within aggregates are disallowed altogether by the SQL standard, treating such cases as errors seems sufficient.
-  Fix trigger transition table capture for `MERGE` in CTE queries (Dean Rasheed) [&sect;](https://postgr.es/c/e7391bbf1)

   When executing a data-modifying CTE query containing both a `MERGE` and another DML operation on a table with statement-level `AFTER` triggers, the transition tables passed to the triggers would not include the rows affected by the `MERGE`, only those affected by the other operation(s).
-  Fix failure when all children of a partitioned target table of an update or delete have been pruned (Amit Langote) [&sect;](https://postgr.es/c/fab386f74)

   In such cases, the executor could report “could not find junk ctid column” errors, even though nothing needs to be done.
-  Avoid possible planner failure when a query contains duplicate window function calls (Meng Zhang, David Rowley) [&sect;](https://postgr.es/c/4297a3519)

   Confusion over de-duplication of such calls could result in errors like “WindowFunc with winref 2 assigned to WindowAgg with winref 1”.
-  Allow indexscans on partial hash indexes even when the index's predicate implies the truth of the WHERE clause (Tom Lane) [&sect;](https://postgr.es/c/b497766a8)

   Normally we drop a WHERE clause that is implied by the predicate, since it's pointless to test it; it must hold for every index entry. However that can prevent creation of an indexscan plan if the index is one that requires a WHERE clause on the leading index key, as hash indexes do. Don't drop implied clauses when considering such an index.
-  Do not emit WAL for unlogged BRIN indexes (Kirill Reshke) [&sect;](https://postgr.es/c/a5277700e)

   One seldom-taken code path incorrectly emitted a WAL record relating to a BRIN index even if the index was marked unlogged. Crash recovery would then fail to replay that record, complaining that the file already exists.
-  Prevent truncation of CLOG that is still needed by unread `NOTIFY` messages (Joel Jacobson, Heikki Linnakangas) [&sect;](https://postgr.es/c/053e1868b) [&sect;](https://postgr.es/c/44e8c60be) [&sect;](https://postgr.es/c/0e8eaa218)

   This fix prevents “could not access status of transaction” errors when a backend is slow to absorb `NOTIFY` messages.
-  Escalate errors occurring during `NOTIFY` message processing to FATAL, i.e. close the connection (Heikki Linnakangas) [&sect;](https://postgr.es/c/c1a5bde00)

   Formerly, if a backend got an error while absorbing a `NOTIFY` message, it would advance past that message, report the error to the client, and move on. That behavior was fraught with problems though. One big concern is that the client has no good way to know that a notification was lost, and certainly no way to know what was in it. Depending on the application logic, missing a notification could cause the application to get stuck waiting, for example. Also, any remaining messages would not get processed until someone sent a new `NOTIFY`.

   Also, if the connection is idle at the time of receiving a `NOTIFY` signal, any ERROR would be escalated to FATAL anyway, due to unrelated concerns. Therefore, we've chosen to make that happen in all cases, for consistency and to provide a clear signal to the application that it might have missed some notifications.
-  Fix bug in following update chain when locking a tuple (Jasper Smit) [&sect;](https://postgr.es/c/7efef18ff)

   This code path neglected to check the xmin of the first new tuple in the update chain, making it possible to lock an unrelated tuple if the original updater aborted and the space was immediately reclaimed by `VACUUM` and then re-used. That could cause unexpected transaction delays or deadlocks. Errors associated with having identified the wrong tuple have also been observed.
-  Fix issues around in-place catalog updates (Noah Misch) [&sect;](https://postgr.es/c/1d7b02711) [&sect;](https://postgr.es/c/720e9304f) [&sect;](https://postgr.es/c/27e4fad98)

   Send a nontransactional invalidation message for an in-place update, since such an update will survive transaction rollback. Also ensure that the update is WAL-logged before other sessions can see it. These fixes primarily prevent scenarios in which relations' frozen-XID attributes become inconsistent, possibly allowing premature CLOG truncation and subsequent “could not access status of transaction” errors.
-  Fix potential backend process crash at process exit due to trying to release a lock in an already-unmapped shared memory segment (Rahila Syed) [&sect;](https://postgr.es/c/980b7c736)
-  Guard against incorrect truncation of the multixact log after a crash (Heikki Linnakangas) [&sect;](https://postgr.es/c/c7946e6f3)
-  Fix possibly mis-encoded result of `pg_stat_get_backend_activity()` (Chao Li) [&sect;](https://postgr.es/c/c48829ed8)

   The shared-memory buffer holding a session's activity string can end with an incomplete multibyte character. Readers are supposed to truncate off any such incomplete character, but this function failed to do so.
-  Guard against recursive memory context logging (Fujii Masao) [&sect;](https://postgr.es/c/3853f6168)

   A constant flow of signals requesting memory context logging could cause recursive execution of the logging code, which in theory could lead to stack overflow.
-  Fix memory context usage when reinitializing a parallel execution context (Jakub Wartak, Jeevan Chalke) [&sect;](https://postgr.es/c/12c2f843c)

   This error could result in a crash due to a subsidiary data structure having a shorter lifespan than the parallel context. The problem is not known to be reachable using only core PostgreSQL, but we have reports of trouble in extensions.
-  Set next multixid's offset when creating a new multixid, to remove the wait loop that was needed in corner cases (Andrey Borodin) [&sect;](https://postgr.es/c/635166913) [&sect;](https://postgr.es/c/4d689a176)

   The previous logic could get stuck waiting for an update that would never occur.
-  Avoid rewriting data-modifying CTEs more than once (Bernice Southey, Dean Rasheed) [&sect;](https://postgr.es/c/4d288e33b)

   Formerly, when updating an auto-updatable view or a relation with rules, if the original query had any data-modifying CTEs, the rewriter would rewrite those CTEs multiple times due to recursion. This was inefficient and could produce false errors if a CTE included an update of an always-generated column.
-  Fail recovery if WAL does not exist back to the redo point indicated by the checkpoint record (Nitin Jadhav) [&sect;](https://postgr.es/c/1aa57e9ed)

   Add an explicit check for this before starting recovery, so that no harm is done and a useful error message is provided. Previously, recovery might crash or corrupt the database in this situation.
-  Avoid scribbling on the source query tree during `ALTER PUBLICATION` (Sunil S) [&sect;](https://postgr.es/c/0687a6eb0)

   This error had the visible effect that an event trigger fired for the query would see only the first `publish` option, even if several had been specified. If such a query were set up as a prepared statement, re-executions would misbehave too.
-  Pass connection options specified in `CREATE SUBSCRIPTION ... CONNECTION` to the publisher's walsender (Fujii Masao) [&sect;](https://postgr.es/c/75f3428f2)

   Before this fix, the `options` connection option (if any) was ignored, thus for example preventing setting custom server parameter values in the walsender session. It was intended for that to work, and it did work before refactoring in PostgreSQL version 15 broke it, so restore the previous behavior.
-  Prevent invalidation of newly created or newly synced replication slots (Zhijie Hou) [&sect;](https://postgr.es/c/24cce33c3)

   A race condition with a concurrent checkpoint could allow WAL to be removed that is needed by the replication slot, causing the slot to immediately get marked invalid.
-  Fix race condition in computing a replication slot's required xmin (Zhijie Hou) [&sect;](https://postgr.es/c/821466722)

   This could lead to the error “cannot build an initial slot snapshot as oldest safe xid follows snapshot's xmin”.
-  During initial synchronization of a logical replication subscription, commit the addition of a `pg_replication_origin` entry before starting to copy data (Zhijie Hou) [&sect;](https://postgr.es/c/e22e9ab0c)

   Previously, if the copy step failed, the new `pg_replication_origin` entry would be lost due to transaction rollback. This led to inconsistent state in shared memory.
-  Don't advance logical replication progress after a parallel worker apply failure (Zhijie Hou) [&sect;](https://postgr.es/c/63a65adf4)

   The previous behavior allowed transactions to be lost by a subscriber.
-  Fix possible failure with “unexpected data beyond EOF” during restart of a streaming replica server (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/a2eeb04f3)
-  Fix erroneous tracking of column position when parsing partition range bounds (myzhen) [&sect;](https://postgr.es/c/821c4d27b)

   This could, for example, lead to the wrong column name being cited in error messages about casting partition bound values to the column's data type.
-  Fix assorted minor errors in error messages (Man Zeng, Tianchen Zhang) [&sect;](https://postgr.es/c/977a17a3e) [&sect;](https://postgr.es/c/e8fd6c9fd) [&sect;](https://postgr.es/c/a7bdbbada)

   For example, an error report about mismatched timeline number in a backup manifest showed the starting timeline number where it meant to show the ending timeline number.
-  Fix failure to perform function inlining when doing JIT compilation with LLVM version 17 or later (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/7600dc79c)
-  Adjust our JIT code to work with LLVM 21 (Holger Hoffstätte) [&sect;](https://postgr.es/c/600acd34b)

   The previous coding failed to compile on aarch64 machines.
-  Add new server parameter [file_extend_method](../../server-administration/server-configuration/resource-consumption.md#guc-file-extend-method) to control use of `posix_fallocate()` (Thomas Munro) [&sect;](https://postgr.es/c/e37b59802)

   PostgreSQL version 16 and later will use `posix_fallocate()`, if the platform provides it, to extend relation files. However, this has been reported to interact poorly with some file systems: BTRFS compression is disabled by the use of `posix_fallocate()`, and XFS could produce spurious `ENOSPC` errors in older Linux kernel versions. To provide a workaround, introduce this new server parameter. Setting `file_extend_method` to `write_zeros` will cause the server to return to the old method of extending files by writing blocks of zeroes.
-  Honor `open()`'s `O_CLOEXEC` flag on Windows (Bryan Green, Thomas Munro) [&sect;](https://postgr.es/c/d62a258cd) [&sect;](https://postgr.es/c/0666ccc16) [&sect;](https://postgr.es/c/80e8ec772)

   Make this flag work like it does on POSIX platforms, so that we don't leak file handles into child processes such as `COPY TO/FROM PROGRAM`. While that leakage hasn't caused many problems, it seems undesirable.
-  Fix failure to parse long options on the server command line in Solaris executables built with meson (Tom Lane) [&sect;](https://postgr.es/c/7369656fa)
-  Support process title changes on GNU/Hurd (Michael Banck) [&sect;](https://postgr.es/c/a1407dade)
-  Make pg_resetwal print the updated value when changing OldestXID (Heikki Linnakangas) [&sect;](https://postgr.es/c/890cc81b6)

   It already did that for every other variable it can change.
-  Make pg_resetwal allow setting next multixact xid to 0 or next multixact offset to UINT32_MAX (Maxim Orlov) [&sect;](https://postgr.es/c/e039b09f8)

   These are valid values, so rejecting them was incorrect. In the worst case, if a pg_upgrade is attempted when exactly at the point of multixact wraparound, the upgrade would fail.
-  In `contrib/amcheck`, use the correct snapshot for btree index parent checks (Mihail Nikalayeu) [&sect;](https://postgr.es/c/098a1fab8)

   The previous coding caused spurious errors when examining indexes created with `CREATE INDEX CONCURRENTLY`.
-  Fix `contrib/amcheck` to handle “half-dead” btree index pages correctly (Heikki Linnakangas) [&sect;](https://postgr.es/c/182901626)

   `amcheck` expected such a page to have a parent downlink, but it does not, leading to a false error report about “mismatch between parent key and child high key”.
-  Fix `contrib/amcheck` to handle incomplete btree root page splits correctly (Heikki Linnakangas) [&sect;](https://postgr.es/c/f2a6df9fd)

   `amcheck` could report a false error about “block is not true root”.
-  Fix edge-case integer overflow in `contrib/intarray`'s selectivity estimator for `@@` (Chao Li) [&sect;](https://postgr.es/c/54f82c4aa)

   This could cause poor selectivity estimates to be produced for cases involving the maximum integer value.
-  Fix multibyte-encoding issue in `contrib/ltree` (Jeff Davis) [&sect;](https://postgr.es/c/b80227c0a)

   The previous coding could pass an incomplete multibyte character to `lower()`, probably resulting in incorrect behavior.
-  Update time zone data files to tzdata release 2025c (Tom Lane) [&sect;](https://postgr.es/c/d852d105e)

   The only change is in historical data for pre-1976 timestamps in Baja California.
