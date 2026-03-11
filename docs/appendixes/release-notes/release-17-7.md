## Release 17.7 { #release-17-7 }


**Release date:.**


2025-11-13


 This release contains a variety of fixes from 17.6. For information about new features in major release 17, see [Release 17](release-17.md#release-17).


### Migration to Version 17.7 { #release-17-7-migration }


 A dump/restore is not required for those running 17.X.


 However, if you are upgrading from a version earlier than 17.6, see [Release 17.6](release-17-6.md#release-17-6).


### Changes { #release-17-7-changes }


-  Check for `CREATE` privileges on the schema in `CREATE STATISTICS` (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/e2fb3dfa8)

   This omission allowed table owners to create statistics in any schema, potentially leading to unexpected naming conflicts.

   The PostgreSQL Project thanks Jelte Fennema-Nio for reporting this problem. (CVE-2025-12817)
-  Avoid integer overflow in allocation-size calculations within libpq (Jacob Champion) [&sect;](https://postgr.es/c/f5999f018)

   Several places in libpq were not sufficiently careful about computing the required size of a memory allocation. Sufficiently large inputs could cause integer overflow, resulting in an undersized buffer, which would then lead to writing past the end of the buffer.

   The PostgreSQL Project thanks Aleksey Solovev of Positive Technologies for reporting this problem. (CVE-2025-12818)
-  Prevent “unrecognized node type” errors when a SQL/JSON function such as `JSON_VALUE` has a `DEFAULT` clause containing a `COLLATE` expression (Jian He) [&sect;](https://postgr.es/c/09f86a42f) [&sect;](https://postgr.es/c/1e6dfdaa0)
-  Correctly treat JSON constructor expressions, such as `JSON_OBJECT()`, as non-strict (Tender Wang, Richard Guo) [&sect;](https://postgr.es/c/d719e2ecb)

   In some cases these expressions can yield a non-null result despite having one or more null inputs, making them non-strict. The planner incorrectly classified them as strict and could perform incorrect query transformations as a result.
-  Further fix processing of character classes within `SIMILAR TO` regular expressions (Laurenz Albe) [&sect;](https://postgr.es/c/e09adb5b9)

   The previous fix for translating `SIMILAR TO` pattern matching expressions to POSIX-style regular expressions broke a corner case that formerly worked: if there is an escape character right after the opening bracket and then a closing bracket right after the escape sequence (for example `[\w]`), the closing bracket was no longer seen as terminating the character class.
-  Fix parsing of aggregate functions whose arguments contain a sub-select with a `FROM` reference to a CTE outside the aggregate function (Tom Lane) [&sect;](https://postgr.es/c/e830896c1)

   Such a CTE reference must act like a outer-level column reference when determining the aggregate's semantic level; but it was not being accounted for, leading to obscure planner or executor errors.
-  Fix “no relation entry for relid” errors in corner cases while estimating SubPlan costs (Richard Guo) [&sect;](https://postgr.es/c/f34202f51)
-  Avoid unlikely use-after-free in planner's expansion of partitioned tables (Bernd Reiß) [&sect;](https://postgr.es/c/ed394c4bd)

   There was a hazard only when the last live partition was concurrently dropped.
-  Remove faulty assertion in btree index cleanup (Peter Geoghegan) [&sect;](https://postgr.es/c/ae15cebc2)
-  Fix possible infinite loop in GIN index scans with multiple scan conditions (Tom Lane) [&sect;](https://postgr.es/c/456c6a05d)

   GIN can handle scan conditions that can reject non-matching entries but are not useful for searching for relevant entries, for example a `tsquery` clause like `!term`. But such a condition must not be first in the array of scan conditions. The code failed to ensure that in all cases, with the result that a query having a mix of such conditions with normal conditions might work or not depending on the order in which the conditions were given in the query.
-  Ensure that GIN index scans can be canceled (Tom Lane) [&sect;](https://postgr.es/c/d17abaea8)

   Some code paths were capable of running for a long time without checking for interrupts.
-  Ensure that BRIN autosummarization provides a snapshot for index expressions that need one (Álvaro Herrera) [&sect;](https://postgr.es/c/f4b68b033) [&sect;](https://postgr.es/c/3b5007347)

   Previously, autosummarization would fail for such indexes, and then leave placeholder index tuples behind, causing the index to bloat over time.
-  Fix integer-overflow hazard in BRIN index scans when the table contains close to 2<sup>32</sup> pages (Sunil S) [&sect;](https://postgr.es/c/c4f5a59ab)

   This oversight could result in an infinite loop or scanning of unneeded table pages.
-  Fix incorrect zero-extension of stored values in JIT-generated tuple deforming code (David Rowley) [&sect;](https://postgr.es/c/10945148e)

   When not using JIT, the equivalent code does sign-extension not zero-extension, leading to a different Datum representation of small integer data types. This inconsistency was masked in most cases, but it is known to lead to “could not find memoization table entry” errors when using Memoize plan nodes, and there might be other symptoms.
-  Fix incorrect logic for caching result-relation information for triggers (David Rowley, Amit Langote) [&sect;](https://postgr.es/c/0d3074615)

   In cases where partitions' column sets aren't physically identical to their parent partitioned tables' column sets, this oversight could lead to crashes.
-  Add missing EvalPlanQual rechecks for TID Scan and TID Range Scan plan nodes (Sophie Alpert, David Rowley) [&sect;](https://postgr.es/c/0fb06e893) [&sect;](https://postgr.es/c/3d939a9b1)

   This omission led to possibly not rechecking a condition on `ctid` during concurrent-update situations, causing the update's behavior to vary depending on which plan type had been selected.
-  Fix EvalPlanQual handling of foreign or custom joins that do not have an alternative local-join plan prepared for EPQ (Masahiko Sawada, Etsuro Fujita) [&sect;](https://postgr.es/c/2bb84ea7e)

   In such cases the foreign or custom access method should be invoked normally, but that did not happen, typically leading to a crash.
-  Avoid duplicating hash partition constraints during `DETACH CONCURRENTLY` (Haiyang Li) [&sect;](https://postgr.es/c/ea06f97ee)

   `ALTER TABLE DETACH PARTITION CONCURRENTLY` was written to add a copy of the partitioning constraint to the now-detached partition. This was misguided, partially because non-concurrent `DETACH` doesn't do that, but mostly because in the case of hash partitioning the constraint expression contains references to the parent table's OID. That causes problems during dump/restore, or if the parent table is dropped after `DETACH`. In v19 and later, we'll no longer create any such copied constraints at all. In released branches, to minimize the risk of unforeseen consequences, only skip adding a copied constraint if it is for hash partitioning.
-  Disallow generated columns in partition keys (Jian He, Ashutosh Bapat) [&sect;](https://postgr.es/c/0b44f2443)

   This was already not allowed, but the check missed some cases, such as where the column reference is implicit in a whole-row reference.
-  Disallow generated columns in `COPY ... FROM ... WHERE` clauses (Peter Eisentraut, Jian He) [&sect;](https://postgr.es/c/07f787e57)

   Previously, incorrect behavior or an obscure error message resulted from attempting to reference such a column, since generated columns have not yet been computed at the point where `WHERE` filtering is done.
-  Avoid potential use-after-free in parallel vacuum (Kevin Oommen Anish) [&sect;](https://postgr.es/c/3549ffb6a)

   This bug seems to have no consequences in standard builds, but it's theoretically a hazard.
-  Fix visibility checking for statistics objects in `pg_temp` (Noah Misch) [&sect;](https://postgr.es/c/6778fbca6)

   A statistics object located in a temporary schema cannot be named without schema qualification, but `pg_statistics_obj_is_visible()` missed that memo and could return “true” regardless. In turn, functions such as `pg_describe_object()` could fail to schema-qualify the object's name as expected.
-  Fix `pg_event_trigger_dropped_objects()`'s reporting of temporary status (Antoine Violin, Tom Lane) [&sect;](https://postgr.es/c/c0c8ee23c) [&sect;](https://postgr.es/c/a220e40d1)

   If a dropped column default, trigger, or RLS policy belongs to a temporary table, report it with `is_temporary` true.
-  Fix memory leakage in hashed subplans (Haiyang Li) [&sect;](https://postgr.es/c/862980f92)

   Any memory consumed by the hash functions used for hashing tuples constituted a query-lifespan memory leak. One way that could happen is if the values being hashed require de-toasting.
-  Avoid leaking `SMgrRelation` objects in the startup process (Jingtang Zhang) [&sect;](https://postgr.es/c/e2dd7b2ac)

   In a long-running standby server, the hashtable holding these objects could bloat substantially, because there was no mechanism for freeing no-longer-interesting entries.
-  Fix minor memory leak during WAL replay of database creation (Nathan Bossart) [&sect;](https://postgr.es/c/f9993ac64)
-  Fix corruption of the shared statistics table after out-of-memory failures (Mikhail Kot) [&sect;](https://postgr.es/c/3e6dfcfb0)

   Previously, an out-of-memory failure partway through creating a new hash table entry left a broken entry behind, potentially causing errors in other sessions later.
-  Fix concurrent update issue in `MERGE` (Yugo Nagata) [&sect;](https://postgr.es/c/6195afbe5)

   When executing a `MERGE UPDATE` action, if there is more than one concurrent update of the target row, the lock-and-retry code would sometimes incorrectly identify the latest version of the target tuple, leading to incorrect results.
-  Add missing replica identity checks in `MERGE` and `INSERT ... ON CONFLICT DO UPDATE` (Zhijie Hou) [&sect;](https://postgr.es/c/76f45be93) [&sect;](https://postgr.es/c/0b934d399) [&sect;](https://postgr.es/c/57dfb64ec)

   If `MERGE` may require update or delete actions, and the target table publishes updates or deletes, insist that it have a `REPLICA IDENTITY` defined. Failing to require this can silently break replication. Likewise, `INSERT` with an `UPDATE` option must require `REPLICA IDENTITY` if the target table publishes either inserts or updates.
-  Avoid deadlock during `DROP SUBSCRIPTION` when publisher is on the same server as subscriber (Dilip Kumar) [&sect;](https://postgr.es/c/288a817bc)
-  Fix incorrect reporting of replication lag in `pg_stat_replication` view (Fujii Masao) [&sect;](https://postgr.es/c/62d5ee75b)

   If any standby server's replay LSN stopped advancing, the `write_lag` and `flush_lag` columns would eventually stop updating.
-  Avoid duplicative log messages about invalid `primary_slot_name` settings (Fujii Masao) [&sect;](https://postgr.es/c/1db2870bb)
-  Avoid failures when `synchronized_standby_slots` references nonexistent replication slots (Shlok Kyal) [&sect;](https://postgr.es/c/0024f5a10)
-  Remove the unfinished slot state file after failing to write a replication slot's state to disk (Michael Paquier) [&sect;](https://postgr.es/c/42348839d)

   Previously, a failure such as out-of-disk-space resulted in leaving a temporary `state.tmp` file behind. That's problematic because it would block all subsequent attempts to write the state, requiring manual intervention to clean up.
-  Fix mishandling of lock timeout signals in parallel apply workers for logical replication (Hayato Kuroda) [&sect;](https://postgr.es/c/2f6e1a490)

   The same signal number was being used for both worker shutdown and lock timeout, leading to confusion.
-  Avoid unwanted WAL receiver shutdown when switching from streaming to archive WAL source (Xuneng Zhou) [&sect;](https://postgr.es/c/e7340b484)

   During a timeline change, a standby server's WAL receiver should remain alive, waiting for a new WAL streaming start point. Instead it was repeatedly shutting down and immediately getting restarted, which could confuse status monitoring code.
-  Avoid failures in logical replication due to chance collisions of file numbers between regular and temporary tables (Vignesh C) [&sect;](https://postgr.es/c/dcdc95cb4)

   This low-probability problem manifested as transient errors like “unexpected duplicate for tablespace *X*, relfilenode *Y*”. `contrib/autoprewarm` was also affected. A side-effect of the fix is that the SQL function `pg_filenode_relation()` will now ignore temporary tables.
-  Fix use-after-free issue in the relation synchronization cache maintained by the pgoutput logical decoding plugin (Vignesh C, Masahiko Sawada) [&sect;](https://postgr.es/c/a61592253)

   An error during logical decoding could result in crashes in subsequent logical decoding attempts in the same session. The case is only reachable when pgoutput is invoked via SQL functions.
-  Avoid unnecessary invalidation of logical replication slots (Bertrand Drouvot) [&sect;](https://postgr.es/c/f3fb6bc9f)
-  Avoid assertion failure when trying to release a replication slot in single-user mode (Hayato Kuroda) [&sect;](https://postgr.es/c/07a302387)
-  Fix incorrect printing of messages about failures in checking whether the user has Windows administrator privilege (Bryan Green) [&sect;](https://postgr.es/c/4c53519e1)

   This code would have crashed or at least printed garbage. No such cases have been reported though, indicating that failure of these system calls is extremely rare.
-  Avoid startup failure on macOS and BSD platforms when there is a collision with a pre-existing semaphore set (Tom Lane) [&sect;](https://postgr.es/c/ab92f0e7f)

   If the pre-existing set has fewer semaphores than we asked for, these platforms return `EINVAL` not `EEXIST` as our code expected, resulting in failure to start the database.
-  Avoid crash when attempting to test PostgreSQL with certain libsanitizer options (Emmanuel Sibi, Jacob Champion) [&sect;](https://postgr.es/c/a9515f294)
-  Fix false memory-context-checking warnings in debug builds on 64-bit Windows (David Rowley) [&sect;](https://postgr.es/c/bd6f986c9)
-  Correctly handle `GROUP BY DISTINCT` in PL/pgSQL assignment statements (Tom Lane) [&sect;](https://postgr.es/c/3fc9aa5b0)

   The parser failed to record the `DISTINCT` option in this context, so that the command would act as if it were plain `GROUP BY`.
-  Avoid leaking memory when handling a SQL error within PL/Python (Tom Lane) [&sect;](https://postgr.es/c/fbc41a145)

   This fixes a session-lifespan memory leak introduced in our previous minor releases.
-  Fix libpq's trace output of characters with the high bit set (Ran Benita) [&sect;](https://postgr.es/c/0fedb3a27)

   On platforms where `char` is considered signed, the output included unsightly `\xffffff` decoration.
-  Fix libpq's handling of socket-related errors on Windows within its GSSAPI logic (Ning Wu, Tom Lane) [&sect;](https://postgr.es/c/1c4671f7b)

   The code for encrypting/decrypting transmitted data using GSSAPI did not correctly recognize error conditions on the connection socket, since Windows reports those differently than other platforms. This led to failure to make such connections on Windows.
-  Fix dumping of non-inherited not-null constraints on inherited table columns (Dilip Kumar) [&sect;](https://postgr.es/c/c945b06d5)

   pg_dump failed to preserve such constraints when dumping from a pre-v18 server.
-  In pg_dump, dump security labels on subscriptions and event triggers (Jian He, Fujii Masao) [&sect;](https://postgr.es/c/968141898)

   Labels on these types of objects were previously missed.
-  Fix pg_dump's sorting of default ACLs and foreign key constraints (Kirill Reshke, Álvaro Herrera) [&sect;](https://postgr.es/c/e8d22095e) [&sect;](https://postgr.es/c/49a09c6c5) [&sect;](https://postgr.es/c/7419c99a2)

   Ensure consistent ordering of these database object types, as was already done for other object types.
-  In pg_dump, label comments for separately-dumped domain constraints with the proper dependency (Noah Misch) [&sect;](https://postgr.es/c/e127764b6)

   This error could lead to parallel pg_restore attempting to create the comment before the constraint itself has been restored.
-  In pg_restore, skip comments and security labels for publications and subscriptions that are not being restored (Jian He, Fujii Masao) [&sect;](https://postgr.es/c/f7f9c5d65) [&sect;](https://postgr.es/c/dc8aa2f58)

   Do not emit `COMMENT` or `SECURITY LABEL` commands for these objects when `--no-publications` or `--no-subscriptions` is specified.
-  Fix assorted errors in the data compression logic in pg_dump and pg_restore (Daniel Gustafsson, Tom Lane) [&sect;](https://postgr.es/c/92268b35d) [&sect;](https://postgr.es/c/bf18e9bd7) [&sect;](https://postgr.es/c/2efca1633)

   Error checking was missing or incorrect in several places, and there were also portability issues that would manifest on big-endian hardware. These problems had been missed because this code is only used to read compressed TOC files within directory-format dumps. pg_dump never produces such a dump; the case can be reached only by manually compressing the TOC file after the fact, which is a supported thing to do but very uncommon.
-  Fix pgbench to error out cleanly if a `COPY` operation is started (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/de6de069d)

   pgbench doesn't intend to support this case, but previously it went into an infinite loop.
-  Fix pgbench's reporting of multiple errors (Yugo Nagata) [&sect;](https://postgr.es/c/a912118c6)

   In cases where two successive `PQgetResult` calls both fail, pgbench might report the wrong error message.
-  In pgbench, fix faulty assertion about errors in pipeline mode (Yugo Nagata) [&sect;](https://postgr.es/c/f39d9164b)
-  Fix per-file memory leakage in pg_combinebackup (Tom Lane) [&sect;](https://postgr.es/c/4eb6992af)
-  Ensure that `contrib/pg_buffercache` functions can be canceled (Satyanarayana Narlapuram, Yuhang Qiu) [&sect;](https://postgr.es/c/b6090ed96)

   Some code paths were capable of running for a long time without checking for interrupts.
-  Fix `contrib/pg_prewarm`'s privilege checks for indexes (Ayush Vatsa, Nathan Bossart) [&sect;](https://postgr.es/c/a0551bc57) [&sect;](https://postgr.es/c/d4e8c37cc)

   `pg_prewarm()` requires `SELECT` privilege on relations to be prewarmed. However, since indexes have no SQL privileges of their own, this resulted in non-superusers being unable to prewarm indexes. Instead, check for `SELECT` privilege on the index's table.
-  Make `contrib/pgstattuple` more robust about empty or invalid index pages (Nitin Motiani) [&sect;](https://postgr.es/c/036decbba)

   Count all-zero pages as free space, and ignore pages that are invalid according to a check of the page's special-space size. The code for btree indexes already counted all-zero pages as free, but the hash and gist code would error out, which has been found to be much less user-friendly. Similarly, make all three cases agree on ignoring corrupted pages rather than throwing errors.
-  Harden our read and write barrier macros to satisfy Clang (Thomas Munro) [&sect;](https://postgr.es/c/03d9140cb)

   We supposed that `__atomic_thread_fence()` is a sufficient barrier to prevent the C compiler from re-ordering memory accesses around it, but it appears that that's not true for Clang, allowing it to generate incorrect code for at least RISC-V, MIPS, and LoongArch machines. Add explicit compiler barriers to fix that.
-  Fix building with LLVM version 21 and later (Holger Hoffstätte) [&sect;](https://postgr.es/c/755f01ad7)
-  When building with meson, apply the same special optimization flags for `numeric.c` and `checksum.c` as the makefile build does (Nathan Bossart, Jeff Davis) [&sect;](https://postgr.es/c/15f9eeef6) [&sect;](https://postgr.es/c/e25453a36)

   Use `-ftree-vectorize` for both files, as well as `-funroll-loops` for `checksum.c`, to match what the makefiles have long done.
-  Fix PGXS build infrastructure to support building NLS `po` files for extensions (Ryo Matsumura) [&sect;](https://postgr.es/c/a8933194e)
