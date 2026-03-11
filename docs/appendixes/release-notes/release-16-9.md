## Release 16.9 { #release-16-9 }


**Release date:.**


2025-05-08


 This release contains a variety of fixes from 16.8. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.9 { #release-16-9-migration }


 A dump/restore is not required for those running 16.X.


 However, if you have any self-referential foreign key constraints on partitioned tables, it may be necessary to recreate those constraints to ensure that they are being enforced correctly. See the second changelog entry below.


 Also, if you have any BRIN bloom indexes, it may be advisable to reindex them after updating. See the third changelog entry below.


 Also, if you are upgrading from a version earlier than 16.5, see [Release 16.5](release-16-5.md#release-16-5).


### Changes { #release-16-9-changes }


-  Avoid one-byte buffer overread when examining invalidly-encoded strings that are claimed to be in GB18030 encoding (Noah Misch, Andres Freund) [&sect;](https://postgr.es/c/d1264948f) [&sect;](https://postgr.es/c/f3bb0b2c4)

   While unlikely, a SIGSEGV crash could occur if an incomplete multibyte character appeared at the end of memory. This was possible both in the server and in libpq-using applications. (CVE-2025-4207)
-  Handle self-referential foreign keys on partitioned tables correctly (Álvaro Herrera) [&sect;](https://postgr.es/c/1817d62ec)

   Creating or attaching partitions failed to make the required catalog entries for a foreign-key constraint, if the table referenced by the constraint was the same partitioned table. This resulted in failure to enforce the constraint fully.

   To fix this, you should drop and recreate any self-referential foreign keys on partitioned tables, if partitions have been created or attached since the constraint was created. Bear in mind that violating rows might already be present, in which case recreating the constraint will fail, and you'll need to fix up those rows before trying again.
-  Avoid data loss when merging compressed BRIN summaries in `brin_bloom_union()` (Tomas Vondra) [&sect;](https://postgr.es/c/ebcc799a7)

   The code failed to account for decompression results not being identical to the input objects, which would result in failure to add some of the data to the merged summary, leading to missed rows in index searches.

   This mistake was present back to v14 where BRIN bloom indexes were introduced, but this code path was only rarely reached then. It's substantially more likely to be hit in v17 because parallel index builds now use the code.
-  Fix unexpected “attribute has wrong type” errors in `UPDATE`, `DELETE`, and `MERGE` queries that use whole-row table references to views or functions in `FROM` (Tom Lane) [&sect;](https://postgr.es/c/fec43428c)
-  Fix `MERGE` into a partitioned table with `DO NOTHING` actions (Tender Wang) [&sect;](https://postgr.es/c/8d4cd3b4a)

   Some cases failed with “unknown action in MERGE WHEN clause” errors.
-  Prevent failure in `INSERT` commands when the table has a `GENERATED` column of a domain data type and the domain's constraints disallow null values (Jian He) [&sect;](https://postgr.es/c/f04e0faa3)

   Constraint failure was reported even if the generation expression produced a perfectly okay result.
-  Correctly process references to outer CTE names that appear within a `WITH` clause attached to an `INSERT`/`UPDATE`/`DELETE`/`MERGE` command that's inside `WITH` (Tom Lane) [&sect;](https://postgr.es/c/1980ec2bc)

   The parser failed to detect disallowed recursion cases, nor did it account for such references when sorting CTEs into a usable order.
-  Don't try to parallelize `array_agg()` when the argument is of an anonymous record type (Richard Guo, Tom Lane) [&sect;](https://postgr.es/c/a7aa9f21f)

   The protocol for communicating with parallel workers doesn't support identifying the concrete record type that a worker is returning.
-  Fix <code>ARRAY(</code><em>subquery</em><code>)</code> and <code>ARRAY[</code><em>expression, ...</em><code>]</code> constructs to produce sane results when the input is of type `int2vector` or `oidvector` (Tom Lane) [&sect;](https://postgr.es/c/0405982c7)

   This patch restores the behavior that existed before PostgreSQL 9.5: the result is of type `int2vector[]` or `oidvector[]`.
-  Fix possible erroneous reports of invalid affixes while parsing Ispell dictionaries (Jacob Brazeal) [&sect;](https://postgr.es/c/61513da08)
-  Fix `ALTER TABLE ADD COLUMN` to correctly handle the case of a domain type that has a default (Jian He, Tom Lane, Tender Wang) [&sect;](https://postgr.es/c/edc3bccd0) [&sect;](https://postgr.es/c/053222a97)

   If a domain type has a default, adding a column of that type (without any explicit `DEFAULT` clause) failed to install the domain's default value in existing rows, instead leaving the new column null.
-  Repair misbehavior when there are duplicate column names in a foreign key constraint's `ON DELETE SET DEFAULT` or `SET NULL` action (Tom Lane) [&sect;](https://postgr.es/c/fb3a77fd9)
-  Improve the error message for disallowed attempts to alter the properties of a foreign key constraint (Álvaro Herrera) [&sect;](https://postgr.es/c/9ea5fe5b6)
-  Avoid error when resetting the `relhassubclass` flag of a temporary table that's marked `ON COMMIT DELETE ROWS` (Noah Misch) [&sect;](https://postgr.es/c/5905e9935)
-  Add missing deparsing of the `INDENT` option of `XMLSERIALIZE()` (Jim Jones) [&sect;](https://postgr.es/c/0af3ae468) [&sect;](https://postgr.es/c/514d47dfb)

   Previously, views or rules using `XMLSERIALIZE(... INDENT)` were dumped without the `INDENT` clause, causing incorrect results after restore.
-  Avoid premature evaluation of the arguments of an aggregate function that has both `FILTER` and `ORDER BY` (or `DISTINCT`) options (David Rowley) [&sect;](https://postgr.es/c/887a23237)

   If there is `ORDER BY` or `DISTINCT`, we consider pre-sorting the aggregate input values rather than doing the sort within the Agg plan node. But this is problematic if the aggregate inputs include expressions that could fail (for example, a division where some of the input divisors could be zero) and there is a `FILTER` clause that's meant to prevent such failures. Pre-sorting would push the expression evaluations to before the `FILTER` test, allowing the failures to happen anyway. Avoid this by not pre-sorting if there's a `FILTER` and the input expressions are anything more complex than a simple Var or Const.
-  Fix planner's failure to identify more than one hashable ScalarArrayOpExpr subexpression within a top-level expression (David Geier) [&sect;](https://postgr.es/c/5a0840b76)

   This resulted in unnecessarily-inefficient execution of any additional subexpressions that could have been processed with a hash table (that is, `IN`, `NOT IN`, or `= ANY` clauses with all-constant right-hand sides).
-  Disable “skip fetch” optimization in bitmap heap scan (Matthias van de Meent) [&sect;](https://postgr.es/c/980727b84)

   It turns out that this optimization can result in returning dead tuples when a concurrent vacuum marks a page all-visible.
-  Fix performance issues in GIN index search startup when there are many search keys (Tom Lane, Vinod Sridharan) [&sect;](https://postgr.es/c/d52221cf0) [&sect;](https://postgr.es/c/4b65b085a)

   An indexable clause with many keys (for example, `jsonbcol ?| array[...]` with tens of thousands of array elements) took O(N<sup>2</sup>) time to start up, and was uncancelable for that interval too.
-  Detect missing support procedures in a BRIN index operator class, and report an error instead of crashing (Álvaro Herrera) [&sect;](https://postgr.es/c/e0d8f49a3)
-  Respond to interrupts (such as query cancel) while waiting for asynchronous subplans of an Append plan node (Heikki Linnakangas) [&sect;](https://postgr.es/c/004dbbd72)

   Previously, nothing would happen until one of the subplans becomes ready.
-  Report the I/O statistics of active WAL senders more frequently (Bertrand Drouvot) [&sect;](https://postgr.es/c/e2a82cd23)

   Previously, the `pg_stat_io` view failed to accumulate I/O performed by a WAL sender until that process exited. Now such I/O will be reported after at most one second's delay.
-  Fix race condition in handling of `synchronous_standby_names` immediately after startup (Melnikov Maksim, Michael Paquier) [&sect;](https://postgr.es/c/c922ae2c4)

   For a short period after system startup, backends might fail to wait for synchronous commit even though `synchronous_standby_names` is enabled.
-  Avoid infinite loop if `scram_iterations` is set to `INT_MAX` (Kevin K Biju) [&sect;](https://postgr.es/c/de1484736)
-  Avoid possible crashes due to double transformation of `json_array()`'s subquery (Tom Lane) [&sect;](https://postgr.es/c/ca54f9b70)
-  Fix `pg_strtof()` to not crash with null endptr (Alexander Lakhin, Tom Lane) [&sect;](https://postgr.es/c/5c64ece8a)
-  Fix crash after out-of-memory in certain GUC assignments (Daniel Gustafsson) [&sect;](https://postgr.es/c/8d48e84c5)
-  Avoid crash when a Snowball stemmer encounters an out-of-memory condition (Maksim Korotkov) [&sect;](https://postgr.es/c/c0c364fa1)
-  Disallow copying of invalidated replication slots (Shlok Kyal) [&sect;](https://postgr.es/c/87e8599e0)

   This prevents trouble when the invalid slot points to WAL that's already been removed.
-  Disallow restoring logical replication slots on standby servers that are not in hot-standby mode (Masahiko Sawada) [&sect;](https://postgr.es/c/cc628f661)

   This prevents a scenario where the slot could remain valid after promotion even if `wal_level` is too low.
-  Prevent over-advancement of catalog xmin in “fast forward” mode of logical decoding (Zhijie Hou) [&sect;](https://postgr.es/c/21a7caeeb)

   This mistake could allow deleted catalog entries to be vacuumed away even though they were still potentially needed by the WAL-reading process.
-  Avoid data loss when DDL operations that don't take a strong lock affect tables that are being logically replicated (Shlok Kyal, Hayato Kuroda) [&sect;](https://postgr.es/c/9a2f8b4f0) [&sect;](https://postgr.es/c/9987c9466)

   The catalog changes caused by the DDL command were not reflected into WAL-decoding processes, allowing them to decode subsequent changes using stale catalog data, probably resulting in data corruption.
-  Prevent incorrect reset of replication origin when an apply worker encounters an error but the error is caught and does not result in worker exit (Hayato Kuroda) [&sect;](https://postgr.es/c/0de091a4b)

   This mistake could allow duplicate data to be applied.
-  Avoid duplicate snapshot creation in logical replication index lookups (Heikki Linnakangas) [&sect;](https://postgr.es/c/8171d2dae) [&sect;](https://postgr.es/c/324e0b656)
-  Improve detection of mixed-origin subscriptions (Hou Zhijie, Shlok Kyal) [&sect;](https://postgr.es/c/1c2a2354c)

   Subscription creation gives a warning if a subscribed-to table is also being followed through other publications, since that could cause duplicate data to be received. This change improves that logic to also detect cases where a partition parent or child table is the one being followed through another publication.
-  Fix wrong checkpoint details in error message about incorrect recovery timeline choice (David Steele) [&sect;](https://postgr.es/c/b4969a296)

   If the requested recovery timeline is not reachable, the reported checkpoint and timeline should be the values read from the backup_label, if there is one. This message previously reported values from the control file, which is correct when recovering from the control file without a backup_label, but not when there is a backup_label.
-  Remove incorrect assertion in `pgstat_report_stat()` (Michael Paquier) [&sect;](https://postgr.es/c/e9ab8677b)
-  Fix overly-strict assertion in `gistFindCorrectParent()` (Heikki Linnakangas) [&sect;](https://postgr.es/c/419321398)
-  Fix rare assertion failure in standby servers when the primary is restarted (Heikki Linnakangas) [&sect;](https://postgr.es/c/2f33de3cd)
-  In PL/pgSQL, avoid “unexpected plan node type” error when a scrollable cursor is defined on a simple <code>SELECT </code><em>expression</em> query (Andrei Lepikhov) [&sect;](https://postgr.es/c/a28c1fb61)
-  Don't try to drop individual index partitions in pg_dump's `--clean` mode (Jian He) [&sect;](https://postgr.es/c/a25f21d99)

   The server rejects such `DROP` commands. That has no real consequences, since the partitions will go away anyway in the subsequent `DROP`s of either their parent tables or their partitioned index. However, the error reported for the attempted drop causes problems when restoring in `--single-transaction` mode.
-  In pg_dumpall, avoid emitting invalid role `GRANT` commands if `pg_auth_members` contains invalid role OIDs (Tom Lane) [&sect;](https://postgr.es/c/d850a6600)

   Instead, print a warning and skip the entry. This copes better with catalog corruption that has been seen to occur in back branches as a result of race conditions between `GRANT` and `DROP ROLE`.
-  In pg_amcheck and pg_upgrade, use the correct function to free allocations made by libpq (Michael Paquier, Ranier Vilela) [&sect;](https://postgr.es/c/9ca2145b0) [&sect;](https://postgr.es/c/57467ec7b) [&sect;](https://postgr.es/c/816149dc6)

   These oversights could result in crashes in certain Windows build configurations, such as a debug build of libpq used by a non-debug build of the calling application.
-  Allow `contrib/dblink` queries to be interrupted by query cancel (Noah Misch) [&sect;](https://postgr.es/c/82a8f0f46)

   This change back-patches a v17-era fix. It prevents possible hangs in `CREATE DATABASE` and `DROP DATABASE` due to failure to detect deadlocks.
-  Avoid crashing with corrupt input data in `contrib/pageinspect`'s `heap_page_items()` (Dmitry Kovalenko) [&sect;](https://postgr.es/c/2d33cf7b8)
-  Prevent assertion failure in `contrib/pg_freespacemap`'s `pg_freespacemap()` (Tender Wang) [&sect;](https://postgr.es/c/41ed749e4)

   Applying `pg_freespacemap()` to a relation lacking storage (such as a view) caused an assertion failure, although there was no ill effect in non-assert builds. Add an error check to reject that case.
-  Fix build failure on macOS 15.4 (Tom Lane, Peter Eisentraut) [&sect;](https://postgr.es/c/a39eb9c77)

   This macOS update broke our configuration probe for `strchrnul()`.
-  Update time zone data files to tzdata release 2025b for DST law changes in Chile, plus historical corrections for Iran (Tom Lane) [&sect;](https://postgr.es/c/e076120d9)

   There is a new time zone America/Coyhaique for Chile's Aysén Region, to account for it changing to UTC-03 year-round and thus diverging from America/Santiago.
