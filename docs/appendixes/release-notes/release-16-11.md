## Release 16.11 { #release-16-11 }


**Release date:.**


2025-11-13


 This release contains a variety of fixes from 16.10. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.11 { #release-16-11-migration }


 A dump/restore is not required for those running 16.X.


 However, if you are upgrading from a version earlier than 16.10, see [Release 16.10](release-16-10.md#release-16-10).


### Changes { #release-16-11-changes }


-  Check for `CREATE` privileges on the schema in `CREATE STATISTICS` (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/d20abb587)

   This omission allowed table owners to create statistics in any schema, potentially leading to unexpected naming conflicts.

   The PostgreSQL Project thanks Jelte Fennema-Nio for reporting this problem. (CVE-2025-12817)
-  Avoid integer overflow in allocation-size calculations within libpq (Jacob Champion) [&sect;](https://postgr.es/c/585fd9b3c)

   Several places in libpq were not sufficiently careful about computing the required size of a memory allocation. Sufficiently large inputs could cause integer overflow, resulting in an undersized buffer, which would then lead to writing past the end of the buffer.

   The PostgreSQL Project thanks Aleksey Solovev of Positive Technologies for reporting this problem. (CVE-2025-12818)
-  Correctly treat JSON constructor expressions, such as `JSON_OBJECT()`, as non-strict (Tender Wang, Richard Guo) [&sect;](https://postgr.es/c/62397bb18)

   In some cases these expressions can yield a non-null result despite having one or more null inputs, making them non-strict. The planner incorrectly classified them as strict and could perform incorrect query transformations as a result.
-  Further fix processing of character classes within `SIMILAR TO` regular expressions (Laurenz Albe) [&sect;](https://postgr.es/c/281ad4ed1)

   The previous fix for translating `SIMILAR TO` pattern matching expressions to POSIX-style regular expressions broke a corner case that formerly worked: if there is an escape character right after the opening bracket and then a closing bracket right after the escape sequence (for example `[\w]`), the closing bracket was no longer seen as terminating the character class.
-  Fix parsing of aggregate functions whose arguments contain a sub-select with a `FROM` reference to a CTE outside the aggregate function (Tom Lane) [&sect;](https://postgr.es/c/7df74e635)

   Such a CTE reference must act like a outer-level column reference when determining the aggregate's semantic level; but it was not being accounted for, leading to obscure planner or executor errors.
-  Fix “no relation entry for relid” errors in corner cases while estimating SubPlan costs (Richard Guo) [&sect;](https://postgr.es/c/79ade5873)
-  Avoid unlikely use-after-free in planner's expansion of partitioned tables (Bernd Reiß) [&sect;](https://postgr.es/c/f0fe1da50)

   There was a hazard only when the last live partition was concurrently dropped.
-  Remove faulty assertion in btree index cleanup (Peter Geoghegan) [&sect;](https://postgr.es/c/c160fd469)
-  Fix possible infinite loop in GIN index scans with multiple scan conditions (Tom Lane) [&sect;](https://postgr.es/c/d532069c3)

   GIN can handle scan conditions that can reject non-matching entries but are not useful for searching for relevant entries, for example a `tsquery` clause like `!term`. But such a condition must not be first in the array of scan conditions. The code failed to ensure that in all cases, with the result that a query having a mix of such conditions with normal conditions might work or not depending on the order in which the conditions were given in the query.
-  Ensure that GIN index scans can be canceled (Tom Lane) [&sect;](https://postgr.es/c/25eadfd0f)

   Some code paths were capable of running for a long time without checking for interrupts.
-  Ensure that BRIN autosummarization provides a snapshot for index expressions that need one (Álvaro Herrera) [&sect;](https://postgr.es/c/6ef33c805) [&sect;](https://postgr.es/c/20442cf50)

   Previously, autosummarization would fail for such indexes, and then leave placeholder index tuples behind, causing the index to bloat over time.
-  Fix integer-overflow hazard in BRIN index scans when the table contains close to 2<sup>32</sup> pages (Sunil S) [&sect;](https://postgr.es/c/ef915bf93)

   This oversight could result in an infinite loop or scanning of unneeded table pages.
-  Fix incorrect zero-extension of stored values in JIT-generated tuple deforming code (David Rowley) [&sect;](https://postgr.es/c/3398b0d02)

   When not using JIT, the equivalent code does sign-extension not zero-extension, leading to a different Datum representation of small integer data types. This inconsistency was masked in most cases, but it is known to lead to “could not find memoization table entry” errors when using Memoize plan nodes, and there might be other symptoms.
-  Fix incorrect logic for caching result-relation information for triggers (David Rowley, Amit Langote) [&sect;](https://postgr.es/c/a546964db)

   In cases where partitions' column sets aren't physically identical to their parent partitioned tables' column sets, this oversight could lead to crashes.
-  Add missing EvalPlanQual rechecks for TID Scan and TID Range Scan plan nodes (Sophie Alpert, David Rowley) [&sect;](https://postgr.es/c/ba0203880) [&sect;](https://postgr.es/c/d6539f88b)

   This omission led to possibly not rechecking a condition on `ctid` during concurrent-update situations, causing the update's behavior to vary depending on which plan type had been selected.
-  Fix EvalPlanQual handling of foreign or custom joins that do not have an alternative local-join plan prepared for EPQ (Masahiko Sawada, Etsuro Fujita) [&sect;](https://postgr.es/c/5a9af4868)

   In such cases the foreign or custom access method should be invoked normally, but that did not happen, typically leading to a crash.
-  Avoid duplicating hash partition constraints during `DETACH CONCURRENTLY` (Haiyang Li) [&sect;](https://postgr.es/c/b835759ec)

   `ALTER TABLE DETACH PARTITION CONCURRENTLY` was written to add a copy of the partitioning constraint to the now-detached partition. This was misguided, partially because non-concurrent `DETACH` doesn't do that, but mostly because in the case of hash partitioning the constraint expression contains references to the parent table's OID. That causes problems during dump/restore, or if the parent table is dropped after `DETACH`. In v19 and later, we'll no longer create any such copied constraints at all. In released branches, to minimize the risk of unforeseen consequences, only skip adding a copied constraint if it is for hash partitioning.
-  Disallow generated columns in partition keys (Jian He, Ashutosh Bapat) [&sect;](https://postgr.es/c/7180d56c5)

   This was already not allowed, but the check missed some cases, such as where the column reference is implicit in a whole-row reference.
-  Disallow generated columns in `COPY ... FROM ... WHERE` clauses (Peter Eisentraut, Jian He) [&sect;](https://postgr.es/c/26958f4d9)

   Previously, incorrect behavior or an obscure error message resulted from attempting to reference such a column, since generated columns have not yet been computed at the point where `WHERE` filtering is done.
-  Fix visibility checking for statistics objects in `pg_temp` (Noah Misch) [&sect;](https://postgr.es/c/ab16418ee)

   A statistics object located in a temporary schema cannot be named without schema qualification, but `pg_statistics_obj_is_visible()` missed that memo and could return “true” regardless. In turn, functions such as `pg_describe_object()` could fail to schema-qualify the object's name as expected.
-  Fix `pg_event_trigger_dropped_objects()`'s reporting of temporary status (Antoine Violin, Tom Lane) [&sect;](https://postgr.es/c/8856f1acc) [&sect;](https://postgr.es/c/f6c8e7824)

   If a dropped column default, trigger, or RLS policy belongs to a temporary table, report it with `is_temporary` true.
-  Fix memory leakage in hashed subplans (Haiyang Li) [&sect;](https://postgr.es/c/e1da9c072)

   Any memory consumed by the hash functions used for hashing tuples constituted a query-lifespan memory leak. One way that could happen is if the values being hashed require de-toasting.
-  Fix minor memory leak during WAL replay of database creation (Nathan Bossart) [&sect;](https://postgr.es/c/c72b5c536)
-  Fix corruption of the shared statistics table after out-of-memory failures (Mikhail Kot) [&sect;](https://postgr.es/c/12f57681c)

   Previously, an out-of-memory failure partway through creating a new hash table entry left a broken entry behind, potentially causing errors in other sessions later.
-  Fix concurrent update issue in `MERGE` (Yugo Nagata) [&sect;](https://postgr.es/c/21a61b87f)

   When executing a `MERGE UPDATE` action, if there is more than one concurrent update of the target row, the lock-and-retry code would sometimes incorrectly identify the latest version of the target tuple, leading to incorrect results.
-  Add missing replica identity checks in `MERGE` and `INSERT ... ON CONFLICT DO UPDATE` (Zhijie Hou) [&sect;](https://postgr.es/c/421d7a157) [&sect;](https://postgr.es/c/0c4d5a45d) [&sect;](https://postgr.es/c/d37694b97)

   If `MERGE` may require update or delete actions, and the target table publishes updates or deletes, insist that it have a `REPLICA IDENTITY` defined. Failing to require this can silently break replication. Likewise, `INSERT` with an `UPDATE` option must require `REPLICA IDENTITY` if the target table publishes either inserts or updates.
-  Avoid deadlock during `DROP SUBSCRIPTION` when publisher is on the same server as subscriber (Dilip Kumar) [&sect;](https://postgr.es/c/7ece76129)
-  Fix incorrect reporting of replication lag in `pg_stat_replication` view (Fujii Masao) [&sect;](https://postgr.es/c/2e55cf4ef)

   If any standby server's replay LSN stopped advancing, the `write_lag` and `flush_lag` columns would eventually stop updating.
-  Avoid duplicative log messages about invalid `primary_slot_name` settings (Fujii Masao) [&sect;](https://postgr.es/c/4fd916eab)
-  Remove the unfinished slot state file after failing to write a replication slot's state to disk (Michael Paquier) [&sect;](https://postgr.es/c/bfdd1a12d)

   Previously, a failure such as out-of-disk-space resulted in leaving a temporary `state.tmp` file behind. That's problematic because it would block all subsequent attempts to write the state, requiring manual intervention to clean up.
-  Fix mishandling of lock timeout signals in parallel apply workers for logical replication (Hayato Kuroda) [&sect;](https://postgr.es/c/a54c7a113)

   The same signal number was being used for both worker shutdown and lock timeout, leading to confusion.
-  Avoid unwanted WAL receiver shutdown when switching from streaming to archive WAL source (Xuneng Zhou) [&sect;](https://postgr.es/c/9b6109607)

   During a timeline change, a standby server's WAL receiver should remain alive, waiting for a new WAL streaming start point. Instead it was repeatedly shutting down and immediately getting restarted, which could confuse status monitoring code.
-  Avoid failures in logical replication due to chance collisions of file numbers between regular and temporary tables (Vignesh C) [&sect;](https://postgr.es/c/ab874faaa)

   This low-probability problem manifested as transient errors like “unexpected duplicate for tablespace *X*, relfilenode *Y*”. `contrib/autoprewarm` was also affected. A side-effect of the fix is that the SQL function `pg_filenode_relation()` will now ignore temporary tables.
-  Fix use-after-free issue in the relation synchronization cache maintained by the pgoutput logical decoding plugin (Vignesh C, Masahiko Sawada) [&sect;](https://postgr.es/c/b07682bce)

   An error during logical decoding could result in crashes in subsequent logical decoding attempts in the same session. The case is only reachable when pgoutput is invoked via SQL functions.
-  Avoid unnecessary invalidation of logical replication slots (Bertrand Drouvot) [&sect;](https://postgr.es/c/e3714dc05)
-  Avoid assertion failure when trying to release a replication slot in single-user mode (Hayato Kuroda) [&sect;](https://postgr.es/c/fea1cc3f7)
-  Fix incorrect printing of messages about failures in checking whether the user has Windows administrator privilege (Bryan Green) [&sect;](https://postgr.es/c/9883e3cd1)

   This code would have crashed or at least printed garbage. No such cases have been reported though, indicating that failure of these system calls is extremely rare.
-  Avoid startup failure on macOS and BSD platforms when there is a collision with a pre-existing semaphore set (Tom Lane) [&sect;](https://postgr.es/c/e67d5f7ba)

   If the pre-existing set has fewer semaphores than we asked for, these platforms return `EINVAL` not `EEXIST` as our code expected, resulting in failure to start the database.
-  Avoid crash when attempting to test PostgreSQL with certain libsanitizer options (Emmanuel Sibi, Jacob Champion) [&sect;](https://postgr.es/c/c775bf048)
-  Fix false memory-context-checking warnings in debug builds on 64-bit Windows (David Rowley) [&sect;](https://postgr.es/c/cdc04a6c3)
-  Correctly handle `GROUP BY DISTINCT` in PL/pgSQL assignment statements (Tom Lane) [&sect;](https://postgr.es/c/b7f6798c0)

   The parser failed to record the `DISTINCT` option in this context, so that the command would act as if it were plain `GROUP BY`.
-  Avoid leaking memory when handling a SQL error within PL/Python (Tom Lane) [&sect;](https://postgr.es/c/cbfd4d0f8)

   This fixes a session-lifespan memory leak introduced in our previous minor releases.
-  Fix libpq's trace output of characters with the high bit set (Ran Benita) [&sect;](https://postgr.es/c/701a0bd56)

   On platforms where `char` is considered signed, the output included unsightly `\xffffff` decoration.
-  Fix libpq's handling of socket-related errors on Windows within its GSSAPI logic (Ning Wu, Tom Lane) [&sect;](https://postgr.es/c/46c4478db)

   The code for encrypting/decrypting transmitted data using GSSAPI did not correctly recognize error conditions on the connection socket, since Windows reports those differently than other platforms. This led to failure to make such connections on Windows.
-  In pg_dump, dump security labels on subscriptions and event triggers (Jian He, Fujii Masao) [&sect;](https://postgr.es/c/20b23784f)

   Labels on these types of objects were previously missed.
-  Fix pg_dump's sorting of default ACLs and foreign key constraints (Kirill Reshke, Álvaro Herrera) [&sect;](https://postgr.es/c/e68fa9a83) [&sect;](https://postgr.es/c/412d29fd2) [&sect;](https://postgr.es/c/06c1ee6b7)

   Ensure consistent ordering of these database object types, as was already done for other object types.
-  In pg_dump, label comments for separately-dumped domain constraints with the proper dependency (Noah Misch) [&sect;](https://postgr.es/c/3cf328eca)

   This error could lead to parallel pg_restore attempting to create the comment before the constraint itself has been restored.
-  In pg_restore, skip comments and security labels for publications and subscriptions that are not being restored (Jian He, Fujii Masao) [&sect;](https://postgr.es/c/97527a5e6) [&sect;](https://postgr.es/c/0870397cc)

   Do not emit `COMMENT` or `SECURITY LABEL` commands for these objects when `--no-publications` or `--no-subscriptions` is specified.
-  Fix assorted errors in the data compression logic in pg_dump and pg_restore (Daniel Gustafsson, Tom Lane) [&sect;](https://postgr.es/c/ec017a305) [&sect;](https://postgr.es/c/1518b7d76) [&sect;](https://postgr.es/c/c865f5b9f)

   Error checking was missing or incorrect in several places, and there were also portability issues that would manifest on big-endian hardware. These problems had been missed because this code is only used to read compressed TOC files within directory-format dumps. pg_dump never produces such a dump; the case can be reached only by manually compressing the TOC file after the fact, which is a supported thing to do but very uncommon.
-  Fix pgbench to error out cleanly if a `COPY` operation is started (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/640590bb4)

   pgbench doesn't intend to support this case, but previously it went into an infinite loop.
-  Fix pgbench's reporting of multiple errors (Yugo Nagata) [&sect;](https://postgr.es/c/36c4d30c8)

   In cases where two successive `PQgetResult` calls both fail, pgbench might report the wrong error message.
-  In pgbench, fix faulty assertion about errors in pipeline mode (Yugo Nagata) [&sect;](https://postgr.es/c/8b2e290bd)
-  Ensure that `contrib/pg_buffercache` functions can be canceled (Satyanarayana Narlapuram, Yuhang Qiu) [&sect;](https://postgr.es/c/815fcfb20)

   Some code paths were capable of running for a long time without checking for interrupts.
-  Fix `contrib/pg_prewarm`'s privilege checks for indexes (Ayush Vatsa, Nathan Bossart) [&sect;](https://postgr.es/c/fae0ce5e3) [&sect;](https://postgr.es/c/c26a8eaf6)

   `pg_prewarm()` requires `SELECT` privilege on relations to be prewarmed. However, since indexes have no SQL privileges of their own, this resulted in non-superusers being unable to prewarm indexes. Instead, check for `SELECT` privilege on the index's table.
-  Make `contrib/pgstattuple` more robust about empty or invalid index pages (Nitin Motiani) [&sect;](https://postgr.es/c/c0f9fe877)

   Count all-zero pages as free space, and ignore pages that are invalid according to a check of the page's special-space size. The code for btree indexes already counted all-zero pages as free, but the hash and gist code would error out, which has been found to be much less user-friendly. Similarly, make all three cases agree on ignoring corrupted pages rather than throwing errors.
-  Harden our read and write barrier macros to satisfy Clang (Thomas Munro) [&sect;](https://postgr.es/c/2f76ffe5e)

   We supposed that `__atomic_thread_fence()` is a sufficient barrier to prevent the C compiler from re-ordering memory accesses around it, but it appears that that's not true for Clang, allowing it to generate incorrect code for at least RISC-V, MIPS, and LoongArch machines. Add explicit compiler barriers to fix that.
-  Fix building with LLVM version 21 and later (Holger Hoffstätte) [&sect;](https://postgr.es/c/2670881af)
-  When building with meson, apply the same special optimization flags for `numeric.c` and `checksum.c` as the makefile build does (Nathan Bossart, Jeff Davis) [&sect;](https://postgr.es/c/509c77929) [&sect;](https://postgr.es/c/2de24ca6c)

   Use `-ftree-vectorize` for both files, as well as `-funroll-loops` for `checksum.c`, to match what the makefiles have long done.
-  Fix PGXS build infrastructure to support building NLS `po` files for extensions (Ryo Matsumura) [&sect;](https://postgr.es/c/a506b0c0a)
