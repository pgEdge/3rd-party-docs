## Release 16.1 { #release-16-1 }


**Release date:.**


2023-11-09


 This release contains a variety of fixes from 16.0. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.1 { #release-16-1-migration }


 A dump/restore is not required for those running 16.X.


 However, several mistakes have been discovered that could lead to certain types of indexes yielding wrong search results or being unnecessarily inefficient. It is advisable to `REINDEX` potentially-affected indexes after installing this update. See the fourth through seventh changelog entries below.


### Changes { #release-16-1-changes }


-  Fix handling of unknown-type arguments in `DISTINCT` `"any"` aggregate functions (Tom Lane) [&sect;](https://postgr.es/c/d3d1e2509)

   This error led to a `text`-type value being interpreted as an `unknown`-type value (that is, a zero-terminated string) at runtime. This could result in disclosure of server memory following the `text` value.

   The PostgreSQL Project thanks Jingzhou Fu for reporting this problem. (CVE-2023-5868)
-  Detect integer overflow while computing new array dimensions (Tom Lane) [&sect;](https://postgr.es/c/e24daa94b)

   When assigning new elements to array subscripts that are outside the current array bounds, an undetected integer overflow could occur in edge cases. Memory stomps that are potentially exploitable for arbitrary code execution are possible, and so is disclosure of server memory.

   The PostgreSQL Project thanks Pedro Gallegos for reporting this problem. (CVE-2023-5869)
-  Prevent the `pg_signal_backend` role from signalling background workers and autovacuum processes (Noah Misch, Jelte Fennema-Nio) [&sect;](https://postgr.es/c/785412731) [&sect;](https://postgr.es/c/2c3c5ec49)

   The documentation says that `pg_signal_backend` cannot issue signals to superuser-owned processes. It was able to signal these background processes, though, because they advertise a role OID of zero. Treat that as indicating superuser ownership. The security implications of cancelling one of these process types are fairly small so far as the core code goes (we'll just start another one), but extensions might add background workers that are more vulnerable.

   Also ensure that the `is_superuser` parameter is set correctly in such processes. No specific security consequences are known for that oversight, but it might be significant for some extensions.

   The PostgreSQL Project thanks Hemanth Sandrana and Mahendrakar Srinivasarao for reporting this problem. (CVE-2023-5870)
-  Fix misbehavior during recursive page split in GiST index build (Heikki Linnakangas) [&sect;](https://postgr.es/c/d7f521325)

   Fix a case where the location of a page downlink was incorrectly tracked, and introduce some logic to allow recovering from such situations rather than silently doing the wrong thing. This error could result in incorrect answers from subsequent index searches. It may be advisable to reindex all GiST indexes after installing this update.
-  Prevent de-duplication of btree index entries for `interval` columns (Noah Misch) [&sect;](https://postgr.es/c/bf1c21c4f)

   There are `interval` values that are distinguishable but compare equal, for example `24:00:00` and `1 day`. This breaks assumptions made by btree de-duplication, so `interval` columns need to be excluded from de-duplication. This oversight can cause incorrect results from index-only scans. Moreover, after updating amcheck will report an error for almost all such indexes. Users should reindex any btree indexes on `interval` columns.
-  Process `date` values more sanely in BRIN `datetime_minmax_multi_ops` indexes (Tomas Vondra) [&sect;](https://postgr.es/c/e7965226d)

   The distance calculation for dates was backward, causing poor decisions about which entries to merge. The index still produces correct results, but is much less efficient than it should be. Reindexing BRIN `minmax_multi` indexes on `date` columns is advisable.
-  Process large `timestamp` and `timestamptz` values more sanely in BRIN `datetime_minmax_multi_ops` indexes (Tomas Vondra) [&sect;](https://postgr.es/c/497fc9208) [&sect;](https://postgr.es/c/0635fe02b)

   Infinities were mistakenly treated as having distance zero rather than a large distance from other values, causing poor decisions about which entries to merge. Also, finite-but-very-large values (near the endpoints of the representable timestamp range) could result in internal overflows, again causing poor decisions. The index still produces correct results, but is much less efficient than it should be. Reindexing BRIN `minmax_multi` indexes on `timestamp` and `timestamptz` columns is advisable if the column contains, or has contained, infinities or large finite values.
-  Avoid calculation overflows in BRIN `interval_minmax_multi_ops` indexes with extreme interval values (Tomas Vondra) [&sect;](https://postgr.es/c/924e0e2ee)

   This bug might have caused unexpected failures while trying to insert large interval values into such an index.
-  Fix partition step generation and runtime partition pruning for hash-partitioned tables with multiple partition keys (David Rowley) [&sect;](https://postgr.es/c/595db9e9c) [&sect;](https://postgr.es/c/6352f1627)

   Some cases involving an `IS NULL` condition on one of the partition keys could result in a crash.
-  Fix inconsistent rechecking of concurrently-updated rows during `MERGE` (Dean Rasheed) [&sect;](https://postgr.es/c/6d2de076c)

   In `READ COMMITTED` mode, an update that finds that its target row was just updated by a concurrent transaction will recheck the query's `WHERE` conditions on the updated row. `MERGE` failed to ensure that the proper rows of other joined tables were used during this recheck, possibly resulting in incorrect decisions about whether the newly-updated row should be updated again by `MERGE`.
-  Correctly identify the target table in an inherited `UPDATE`/`DELETE`/`MERGE` even when the parent table is excluded by constraints (Amit Langote, Tom Lane) [&sect;](https://postgr.es/c/b1444a09d) [&sect;](https://postgr.es/c/178ee1d85) [&sect;](https://postgr.es/c/2bf99b48d)

   If the initially-named table is excluded by constraints, but not all its inheritance descendants are, the first non-excluded descendant was identified as the primary target table. This would lead to firing statement-level triggers associated with that table, rather than the initially-named table as should happen. In v16, the same oversight could also lead to “invalid perminfoindex 0 in RTE with relid NNNN” errors.
-  Fix edge case in btree mark/restore processing of ScalarArrayOpExpr clauses (Peter Geoghegan) [&sect;](https://postgr.es/c/3fa81b62e)

   When restoring an indexscan to a previously marked position, the code could miss required setup steps if the scan had advanced exactly to the end of the matches for a ScalarArrayOpExpr (that is, an `indexcol = ANY(ARRAY[])`) clause. This could result in missing some rows that should have been fetched.
-  Fix intra-query memory leak in Memoize execution (Orlov Aleksej, David Rowley) [&sect;](https://postgr.es/c/31b2b2d72)
-  Fix intra-query memory leak when a set-returning function repeatedly returns zero rows (Tom Lane) [&sect;](https://postgr.es/c/07494a0df)
-  Don't crash if `cursor_to_xmlschema()` is applied to a non-data-returning Portal (Boyu Yang) [&sect;](https://postgr.es/c/ec693a3f3)
-  Fix improper sharing of origin filter condition across successive `pg_logical_slot_get_changes()` calls (Hou Zhijie) [&sect;](https://postgr.es/c/8d05be931)

   The origin condition set by one call of this function would be re-used by later calls that did not specify the origin argument. This was not intended.
-  Throw the intended error if `pgrowlocks()` is applied to a partitioned table (David Rowley) [&sect;](https://postgr.es/c/a98f01c93)

   Previously, a not-on-point complaint “only heap AM is supported” would be raised.
-  Handle invalid indexes more cleanly in assorted SQL functions (Noah Misch) [&sect;](https://postgr.es/c/1a368dd3e)

   Report an error if `pgstatindex()`, `pgstatginindex()`, `pgstathashindex()`, or `pgstattuple()` is applied to an invalid index. If `brin_desummarize_range()`, `brin_summarize_new_values()`, `brin_summarize_range()`, or `gin_clean_pending_list()` is applied to an invalid index, do nothing except to report a debug-level message. Formerly these functions attempted to process the index, and might fail in strange ways depending on what the failed `CREATE INDEX` had left behind.
-  Avoid premature memory allocation failure with long inputs to `to_tsvector()` (Tom Lane) [&sect;](https://postgr.es/c/8465efc1a)
-  Fix over-allocation of the constructed `tsvector` in `tsvectorrecv()` (Denis Erokhin) [&sect;](https://postgr.es/c/5c34a7374)

   If the incoming vector includes position data, the binary receive function left wasted space (roughly equal to the size of the position data) in the finished `tsvector`. In extreme cases this could lead to “maximum total lexeme length exceeded” failures for vectors that were under the length limit when emitted. In any case it could lead to wasted space on-disk.
-  Improve checks for corrupt PGLZ compressed data (Flavien Guedez) [&sect;](https://postgr.es/c/cfa4eba02)
-  Fix `ALTER SUBSCRIPTION` so that a commanded change in the `run_as_owner` option is actually applied (Hou Zhijie) [&sect;](https://postgr.es/c/a81e5516f)
-  Fix bulk table insertion into partitioned tables (Andres Freund) [&sect;](https://postgr.es/c/0002feb82)

   Improper sharing of insertion state across partitions could result in failures during `COPY FROM`, typically manifesting as “could not read block NNNN in file XXXX: read only 0 of 8192 bytes” errors.
-  In `COPY FROM`, avoid evaluating column default values that will not be needed by the command (Laurenz Albe) [&sect;](https://postgr.es/c/910eb61b2)

   This avoids a possible error if the default value isn't actually valid for the column, or if the default's expression would fail in the current execution context. Such edge cases sometimes arise while restoring dumps, for example. Previous releases did not fail in this situation, so prevent v16 from doing so.
-  In `COPY FROM`, fail cleanly when an unsupported encoding conversion is needed (Tom Lane) [&sect;](https://postgr.es/c/ea0e7cd6b)

   Recent refactoring accidentally removed the intended error check for this, such that it ended in “cache lookup failed for function 0” instead of a useful error message.
-  Avoid crash in `EXPLAIN` if a parameter marked to be displayed by `EXPLAIN` has a NULL boot-time value (Xing Guo, Aleksander Alekseev, Tom Lane) [&sect;](https://postgr.es/c/82063edd4)

   No built-in parameter fits this description, but an extension could define such a parameter.
-  Ensure we have a snapshot while dropping `ON COMMIT DROP` temp tables (Tom Lane) [&sect;](https://postgr.es/c/57e6e861d)

   This prevents possible misbehavior if any catalog entries for the temp tables have fields wide enough to require toasting (such as a very complex `CHECK` condition).
-  Avoid improper response to shutdown signals in child processes just forked by `system()` (Nathan Bossart) [&sect;](https://postgr.es/c/ee06199fc)

   This fix avoids a race condition in which a child process that has been forked off by `system()`, but hasn't yet exec'd the intended child program, might receive and act on a signal intended for the parent server process. That would lead to duplicate cleanup actions being performed, which will not end well.
-  Cope with torn reads of `pg_control` in frontend programs (Thomas Munro) [&sect;](https://postgr.es/c/5725e4ebe)

   On some file systems, reading `pg_control` may not be an atomic action when the server concurrently writes that file. This is detectable via a bad CRC. Retry a few times to see if the file becomes valid before we report error.
-  Avoid torn reads of `pg_control` in relevant SQL functions (Thomas Munro) [&sect;](https://postgr.es/c/2371432cd)

   Acquire the appropriate lock before reading `pg_control`, to ensure we get a consistent view of that file.
-  Fix “could not find pathkey item to sort” errors occurring while planning aggregate functions with `ORDER BY` or `DISTINCT` options (David Rowley) [&sect;](https://postgr.es/c/9154ededf)
-  Avoid integer overflow when computing size of backend activity string array (Jakub Wartak) [&sect;](https://postgr.es/c/75f31a3f2)

   On 64-bit machines we will allow values of `track_activity_query_size` large enough to cause 32-bit overflow when multiplied by the allowed number of connections. The code actually allocating the per-backend local array was careless about this though, and allocated the array incorrectly.
-  Fix briefly showing inconsistent progress statistics for `ANALYZE` on inherited tables (Heikki Linnakangas) [&sect;](https://postgr.es/c/992d2ca81)

   The block-level counters should be reset to zero at the same time we update the current-relation field.
-  Fix the background writer to report any WAL writes it makes to the statistics counters (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/4a97a43a7)
-  Fix confusion about forced-flush behavior in `pgstat_report_wal()` (Ryoga Yoshida, Michael Paquier) [&sect;](https://postgr.es/c/280f70221)

   This could result in some statistics about WAL I/O being forgotten in a shutdown.
-  Fix statistics tracking of temporary-table extensions (Karina Litskevich, Andres Freund) [&sect;](https://postgr.es/c/c4758649b)

   These were counted as normal-table writes when they should be counted as temp-table writes.
-  When `track_io_timing` is enabled, include the time taken by relation extension operations as write time (Nazir Bilal Yavuz) [&sect;](https://postgr.es/c/2308f18c0)
-  Track the dependencies of cached `CALL` statements, and re-plan them when needed (Tom Lane) [&sect;](https://postgr.es/c/055f786ea)

   DDL commands, such as replacement of a function that has been inlined into a `CALL` argument, can create the need to re-plan a `CALL` that has been cached by PL/pgSQL. That was not happening, leading to misbehavior or strange errors such as “cache lookup failed”.
-  Avoid a possible pfree-a-NULL-pointer crash after an error in OpenSSL connection setup (Sergey Shinderuk) [&sect;](https://postgr.es/c/f720875a4)
-  Track nesting depth correctly when inspecting `RECORD`-type Vars from outer query levels (Richard Guo) [&sect;](https://postgr.es/c/53630f12d)

   This oversight could lead to assertion failures, core dumps, or “bogus varno” errors.
-  Track hash function and negator function dependencies of ScalarArrayOpExpr plan nodes (David Rowley) [&sect;](https://postgr.es/c/1a6900e58)

   In most cases this oversight was harmless, since these functions would be unlikely to disappear while the node's original operator remains present.
-  Fix error-handling bug in `RECORD` type cache management (Thomas Munro) [&sect;](https://postgr.es/c/f899c7f1e)

   An out-of-memory error occurring at just the wrong point could leave behind inconsistent state that would lead to an infinite loop.
-  Treat out-of-memory failures as fatal while reading WAL (Michael Paquier) [&sect;](https://postgr.es/c/a06efbc3a)

   Previously this would be treated as a bogus-data condition, leading to the conclusion that we'd reached the end of WAL, which is incorrect and could lead to inconsistent WAL replay.
-  Fix possible recovery failure due to trying to allocate memory based on a bogus WAL record length field (Thomas Munro, Michael Paquier) [&sect;](https://postgr.es/c/ce497f648) [&sect;](https://postgr.es/c/10d0591ea)
-  Fix “could not duplicate handle” error occurring on Windows when `min_dynamic_shared_memory` is set above zero (Thomas Munro) [&sect;](https://postgr.es/c/174ccda5e)
-  Fix order of operations in `GenericXLogFinish` (Jeff Davis) [&sect;](https://postgr.es/c/b8963e8a2)

   This code violated the conditions required for crash safety by writing WAL before marking changed buffers dirty. No core code uses this function, but extensions do (`contrib/bloom` does, for example).
-  Remove incorrect assertion in PL/Python exception handling (Alexander Lakhin) [&sect;](https://postgr.es/c/f171430f0)
-  Fix pg_dump to dump the new `run_as_owner` option of subscriptions (Philip Warner) [&sect;](https://postgr.es/c/67738dbf9)

   Due to this oversight, subscriptions would always be restored with `run_as_owner` set to `false`, which is not equivalent to their behavior in pre-v16 releases.
-  Fix pg_restore so that selective restores will include both table-level and column-level ACLs for selected tables (Euler Taveira, Tom Lane) [&sect;](https://postgr.es/c/aaaf8fbb6)

   Formerly, only the table-level ACL would get restored if both types were present.
-  Add logic to pg_upgrade to check for use of `abstime`, `reltime`, and `tinterval` data types (&Aacute;lvaro Herrera) [&sect;](https://postgr.es/c/fb9ddd0fa)

   These obsolete data types were removed in PostgreSQL version 12, so check to make sure they aren't present in an older database before claiming it can be upgraded.
-  Avoid false “too many client connections” errors in pgbench on Windows (Noah Misch) [&sect;](https://postgr.es/c/06ff06484)
-  Fix vacuumdb's handling of multiple `-N` switches (Nathan Bossart, Kuwamura Masaki) [&sect;](https://postgr.es/c/2143d96dc)

   Multiple `-N` switches should exclude tables in multiple schemas, but in fact excluded nothing due to faulty construction of a generated query.
-  Fix vacuumdb to honor its `--buffer-usage-limit` option in analyze-only mode (Ryoga Yoshida, David Rowley) [&sect;](https://postgr.es/c/f7dbdab05)
-  In `contrib/amcheck`, do not report interrupted page deletion as corruption (Noah Misch) [&sect;](https://postgr.es/c/3c6a05b80)

   This fix prevents false-positive reports of “the first child of leftmost target page is not leftmost of its level”, “block NNNN is not leftmost” or “left link/right link pair in index XXXX not in agreement”. They appeared if amcheck ran after an unfinished btree index page deletion and before `VACUUM` had cleaned things up.
-  Fix failure of `contrib/btree_gin` indexes on `interval` columns, when an indexscan using the `<` or `<=` operator is performed (Dean Rasheed) [&sect;](https://postgr.es/c/ab73a37e9)

   Such an indexscan failed to return all the entries it should.
-  Add support for LLVM 16 and 17 (Thomas Munro, Dmitry Dolgov) [&sect;](https://postgr.es/c/774185056) [&sect;](https://postgr.es/c/74d19ec09) [&sect;](https://postgr.es/c/60596f148)
-  Suppress assorted build-time warnings on recent macOS (Tom Lane) [&sect;](https://postgr.es/c/75c562653) [&sect;](https://postgr.es/c/e73d6a0df)

   Xcode 15 (released with macOS Sonoma) changed the linker's behavior in a way that causes many duplicate-library warnings while building PostgreSQL. These were harmless, but they're annoying so avoid citing the same libraries twice. Also remove use of the `-multiply_defined suppress` linker switch, which apparently has been a no-op for a long time, and is now actively complained of.
-  When building `contrib/unaccent`'s rules file, fall back to using `python` if `--with-python` was not given and make variable `PYTHON` was not set (Japin Li) [&sect;](https://postgr.es/c/641db601b)
-  Remove `PHOT` (Phoenix Islands Time) from the default timezone abbreviations list (Tom Lane) [&sect;](https://postgr.es/c/d1537afe3)

   Presence of this abbreviation in the default list can cause failures on recent Debian and Ubuntu releases, as they no longer install the underlying tzdb entry by default. Since this is a made-up abbreviation for a zone with a total human population of about two dozen, it seems unlikely that anyone will miss it. If someone does, they can put it back via a custom abbreviations file.
