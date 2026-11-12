<a id="release-17-10"></a>

## Release 17.10


**Release date:.**


2026-05-14


 This release contains a variety of fixes from 17.9. For information about new features in major release 17, see [Release 17](release-17.md#release-17).
 <a id="release-17-10-migration"></a>

### Migration to Version 17.10


 A dump/restore is not required for those running 17.X.


 However, if you are upgrading from a version earlier than 17.6, see [Release 17.6](release-17-6.md#release-17-6).
  <a id="release-17-10-changes"></a>

### Changes


-  Prevent unbounded recursion while processing startup packets (Michael Paquier) [&sect;](https://postgr.es/c/32a4ce55c) [&sect;](https://postgr.es/c/6dffaeb8e)

   A malicious client could crash the connected backend by alternating rejected SSL and GSS encryption requests indefinitely.

   The PostgreSQL Project thanks Calif.io (in collaboration with Claude and Anthropic Research) for reporting this problem. (CVE-2026-6479)
-  Fix assorted integer overflows in memory-allocation calculations (Tom Lane, Nathan Bossart, Heikki Linnakangas) [&sect;](https://postgr.es/c/fe2720c45) [&sect;](https://postgr.es/c/01b5ef7df) [&sect;](https://postgr.es/c/e3a2bea41) [&sect;](https://postgr.es/c/26dd3cac2) [&sect;](https://postgr.es/c/87357a606) [&sect;](https://postgr.es/c/3c41f5534) [&sect;](https://postgr.es/c/ebcfa7867) [&sect;](https://postgr.es/c/00e243e67) [&sect;](https://postgr.es/c/e5babf754) [&sect;](https://postgr.es/c/8e909812d)

   Various places were incautious about the possibility of integer overflow in calculations of how much memory to allocate. Overflow would lead to allocating a too-small buffer which the caller would then write past the end of. This would at least trigger server crashes, and probably could be exploited for arbitrary code execution. In many but by no means all cases, the hazard exists only in 32-bit builds.

   The PostgreSQL Project thanks Xint Code, Bruce Dang, Sven Klemm, and Pavel Kohout for reporting these problems. (CVE-2026-6473)
-  Properly quote subscription names in pg_createsubscriber (Nathan Bossart) [&sect;](https://postgr.es/c/d7de7fa84)

   The given subscription name was inserted into SQL commands without quoting, so that SQL injection could be achieved in the (perhaps unlikely) case that the subscription name comes from an untrusted source.

   The PostgreSQL Project thanks Yu Kunpeng for reporting this problem. (CVE-2026-6476)
-  Properly quote object names in logical replication origin checks (Pavel Kohout) [&sect;](https://postgr.es/c/f0f59b658)

   `ALTER SUBSCRIPTION ... REFRESH PUBLICATION` interpolated schema and relation names into SQL commands without quoting them, allowing execution of arbitrary SQL on the publisher.

   The PostgreSQL Project thanks Pavel Kohout for reporting this problem. (CVE-2026-6638)
-  Reject over-length options in `ts_headline()` (Michael Paquier) [&sect;](https://postgr.es/c/3ed3dbbf4)

   The `StartSel`, `StopSel` and `FragmentDelimiter` strings must not exceed 32Kb in length, but this was not checked for. An over-length value would typically crash the server.

   The PostgreSQL Project thanks Xint Code for reporting this problem. (CVE-2026-6473)
-  Guard against malicious time zone names in `timeofday()` and `pg_strftime()` (Tom Lane) [&sect;](https://postgr.es/c/4197c880c) [&sect;](https://postgr.es/c/a386d14fe)

   A crafted time zone setting could pass `%` sequences to `snprintf()`, potentially causing crashes or disclosure of server memory. Another path to similar results was to overflow the limited-size output buffer used by `pg_strftime()`.

   The PostgreSQL Project thanks Xint Code for reporting this problem. (CVE-2026-6474)
-  When creating a multirange type, ensure the user has `CREATE` privilege on the schema specified for the multirange type (Jelte Fennema-Nio) [&sect;](https://postgr.es/c/c27ba08cd)

   The multirange type can be put into a different schema than its parent range type, but we neglected to apply the required privilege check when doing so.

   The PostgreSQL Project thanks Jelte Fennema-Nio for reporting this problem. (CVE-2026-6472)
-  Use timing-safe string comparisons in authentication code (Michael Paquier) [&sect;](https://postgr.es/c/c4e7435b3) [&sect;](https://postgr.es/c/8e34acfda)

   Use `timingsafe_bcmp()` instead of `memcmp()` or `strcmp()` when checking passwords, hashes, etc. It is not known whether the data dependency of those functions is usefully exploitable in any of these places, but in the interests of safety, replace them.

   The PostgreSQL Project thanks Joe Conway for reporting this problem. (CVE-2026-6478)
-  Mark `PQfn()` as unsafe, and avoid using it within libpq (Nathan Bossart) [&sect;](https://postgr.es/c/d88c7be15)

   For a non-integral result type, `PQfn()` is not passed the size of the output buffer, so it cannot check that the data returned by the server will fit. A malicious server could therefore overwrite client memory. This is unfixable without an API change, so mark the function as deprecated. Internally to libpq, use a variant version that can apply the missing check.

   The PostgreSQL Project thanks Yu Kunpeng and Martin Heistermann for reporting this problem. (CVE-2026-6477)
-  Prevent path traversal in pg_basebackup and pg_rewind (Michael Paquier) [&sect;](https://postgr.es/c/8f881e188)

   These applications failed to validate output file paths read from their input, so that a malicious source could overwrite any file writable by these applications. Constrain where data can be written by rejecting paths that are absolute or contain parent-directory references.

   The PostgreSQL Project thanks XlabAI Team of Tencent Xuanwu Lab and Valery Gubanov for reporting this problem. (CVE-2026-6475)
-  Guard against field overflow within `contrib/intarray`'s `query_int` type and `contrib/ltree`'s `ltxtquery` type (Tom Lane) [&sect;](https://postgr.es/c/c4d04cc48) [&sect;](https://postgr.es/c/2b429d887)

   Parsing of these query structures did not check for overflow of 16-bit fields, so that construction of an invalid query tree was possible. This can crash the server when executing the query.

   The PostgreSQL Project thanks Xint Code for reporting this problem. (CVE-2026-6473)
-  Guard against overly long values of `contrib/ltree`'s `lquery` type (Michael Paquier) [&sect;](https://postgr.es/c/8c3426110)

   Values with more than 64K items caused internal overflows, potentially resulting in stack smashes or wrong answers.

   The PostgreSQL Project thanks Vergissmeinnicht, A1ex, and Jihe Wang for reporting this problem. (CVE-2026-6473)
-  Prevent SQL injection and buffer overruns in `contrib/spi` (Nathan Bossart) [&sect;](https://postgr.es/c/2dc64ef28)

   `check_foreign_key()` was insufficiently careful about quoting key values, and also used fixed-length buffers for constructing queries. While this module is only meant as example code, it still shouldn't contain such dangerous errors.

   The PostgreSQL Project thanks Nikolay Samokhvalov for reporting this problem. (CVE-2026-6637)
-  Check for nondeterministic collations before assuming that an equality condition on a collatable type implies uniqueness (Richard Guo) [&sect;](https://postgr.es/c/d0e73bb18) [&sect;](https://postgr.es/c/13226050e)

   Numerous planner optimizations assume that, for example, at most one table row can satisfy `WHERE x = 'abc'` if there is a unique index on `x`. However this conclusion is unsafe in general if the index and the `WHERE` clause have different collations attached. It is safe when both collations are deterministic, because that property essentially requires that equality of two strings means bitwise equality. But nondeterministic collations don't act that way, so that optimizing on the assumption of unique matches can give wrong query answers if either the `WHERE` clause or the index has a nondeterministic collation.
-  Fix incomplete removal of relation references in `RestrictInfo` structs during join removal (Tom Lane) [&sect;](https://postgr.es/c/766d40286) [&sect;](https://postgr.es/c/53cb4ec1d)

   This oversight has been shown to result in planner failures such as unexpected “FULL JOIN is only supported with merge-joinable or hash-joinable join conditions” errors. It may also have caused failure to consider valid plans in other cases.
-  Fix incorrect handling of `NEW` generated columns in rule actions and rule qualifications (Richard Guo, Dean Rasheed) [&sect;](https://postgr.es/c/9d6208939)

   Previously, such column references would produce NULL in `INSERT` cases, or be equivalent to the `OLD` value in `UPDATE` cases.
-  Fix spurious “generated columns are not supported in COPY FROM WHERE conditions” errors (Tom Lane) [&sect;](https://postgr.es/c/681a91d29)

   Use of a system column in a `COPY FROM WHERE` condition could sometimes incorrectly report this error.
-  Correctly report a serialization failure when `MERGE` encounters a concurrently-updated tuple in repeatable-read or serializable mode (Tender Wang) [&sect;](https://postgr.es/c/2dcac93c0)

   Previously, such cases behaved the same as in lower isolation levels.
-  Fix `CREATE TABLE ... LIKE ... INCLUDING STATISTICS` for cases where the source table has dropped column(s) (Julien Tachoires) [&sect;](https://postgr.es/c/a0104b447)

   In such cases, extended statistics objects could be copied incorrectly, or the command could give an incorrect error.
-  Allow `ALTER INDEX ... ATTACH PARTITION` to mark the parent index valid if appropriate (Sami Imseih) [&sect;](https://postgr.es/c/becf6d269)

   There are edge cases in which a partitioned index might remain marked as invalid even when all its leaf indexes are valid. This change provides a mechanism whereby a user can correct such a situation without resorting to manual catalog updates.
-  Fix `ALTER FOREIGN DATA WRAPPER` to not drop the wrapper object's dependency on its handler function (Jeff Davis) [&sect;](https://postgr.es/c/876fa84a2)
-  Disallow making a composite type be a member of itself via a multirange (Heikki Linnakangas) [&sect;](https://postgr.es/c/54343f6f9)

   We already forbade such cases when the intermediate type is a domain, array, composite type, or range; but multiranges were overlooked.
-  Fix datum-image comparisons to be insensitive to sign-extension variations (David Rowley) [&sect;](https://postgr.es/c/d29808e35)

   This fixes some situations that previously led to “could not find memoization table entry” errors or wrong query results.
-  Fix incorrect logic for hashed `IN`/`NOT IN` with non-strict equality operator (Chengpeng Yan) [&sect;](https://postgr.es/c/3fda3e12f)

   The previous coding could crash or give wrong answers. All built-in data types have strict equality operators, so that this issue could only arise with an extension data type.
-  Truncate overly-long locale-specific numeric symbols in `to_char()` (Tom Lane) [&sect;](https://postgr.es/c/c97a28618)

   If a locale specified a currency symbol, thousands separator, or decimal or sign symbol more than 8 bytes long, a buffer overrun was possible. No such locales exist in the real world, and it's impractical for an unprivileged attacker to install a malicious locale definition underneath a Postgres server; but for safety's sake check for overlength symbols and truncate if needed.
-  Prevent buffer overruns when parsing an affix file for an `Ispell` dictionary (Tom Lane) [&sect;](https://postgr.es/c/ea5f0d176) [&sect;](https://postgr.es/c/a5426dbf8)

   A corrupt or malicious affix file could crash the server. This is not considered a security issue because text search configuration files are presumed trustworthy, but it still seems worth fixing.
-  Guard against integer overflow in calculations of frame start and end positions for window aggregates (Richard Guo) [&sect;](https://postgr.es/c/f8736f8bc)

   Very large user-specified offsets (close to INT64_MAX) could result in errors or incorrect query results.
-  Fix `array_agg_array_combine()` to combine the arrays' null bitmaps correctly (Dmytro Astapov) [&sect;](https://postgr.es/c/d6c9432cb)

   This mistake resulted in sometimes-incorrect output from parallelized `array_agg(anyarray)` calculations.
-  Retry `sync_file_range()` if it returns error code `EINTR` (DaeMyung Kang) [&sect;](https://postgr.es/c/5499be332)
-  Fix incorrect behavior of `pg_stat_reset_single_table_counters()` on a shared catalog (Chao Li) [&sect;](https://postgr.es/c/a4fefb3e0)

   Such cases had a side-effect of resetting the current database's `stat_reset_timestamp`, which was unintended.
-  Update activity statistics when a parallel apply worker is idle (Zhijie Hou) [&sect;](https://postgr.es/c/88d7fdcc9)

   Previously, statistics from a recently-completed transaction might go unreported for long intervals, particularly if the workload is light.
-  Fix “no relation entry for relid 0” failure while estimating array lengths in set operations (Tender Wang) [&sect;](https://postgr.es/c/93ed18720)
-  Fix buffer overread when `pglz_decompress()` receives corrupt input (Andrew Dunstan) [&sect;](https://postgr.es/c/c05c3baf1)

   It was possible to read a few bytes past the end of the input, which in very unlucky cases might cause a crash.
-  Fix incremental JSON parser's handling of numeric tokens that cross input buffer boundaries (Andrew Dunstan) [&sect;](https://postgr.es/c/2e373785e)

   It was possible to accept an incorrectly-formatted number, leading to failures later.
-  Prevent bloating relation visibility maps during restore of an incremental backup (Robert Haas) [&sect;](https://postgr.es/c/076bc57fa)

   Restore could append many blocks of zeroes to a visibility map, due to incorrect computation of the expected file length. This does not result in data corruption, but it could waste a substantial amount of disk space.
-  Use C collation, not the database's default collation, in catalog cache lookups on text columns (Jeff Davis) [&sect;](https://postgr.es/c/9dda30dd3)

   This avoids failures in edge cases such as physical replication startup, where there is no identified database so that a default collation cannot be determined.
-  Prevent stuck slotsync worker processes from blocking promotion of a standby server (Nisha Moond, Ajin Cherian) [&sect;](https://postgr.es/c/15910b1c3) [&sect;](https://postgr.es/c/586f4266f) [&sect;](https://postgr.es/c/4bed04d39)

   A worker process that was vainly waiting for a response from the primary would delay promotion for an unreasonable amount of time.
-  Fix excessive log output from idle slotsync worker processes (Zhijie Hou) [&sect;](https://postgr.es/c/91741b7cb)
-  Ensure that tuplestore data structures are internally consistent even after an error (Tom Lane) [&sect;](https://postgr.es/c/1f5b6a5e5)

   The code was previously careless about this, which is fine most of the time but is problematic for the tuplestore backing a `WITH HOLD` cursor. In v15 and before this leads to easily-reproducible crashes; later branches are not known to be vulnerable, but it seems best to preserve consistency in all.
-  Fix premature NULL lag reporting in `pg_stat_replication` (Shinya Kato) [&sect;](https://postgr.es/c/fdce5de55)

   The lag columns frequently read as NULL even while replication activity was happening.
-  Avoid rare flush failure when working with non-WAL-logged GiST indexes (Tomas Vondra) [&sect;](https://postgr.es/c/6ef36bb35)

   A non-logged GiST index could nonetheless sometimes produce “xlog flush request *n/nnnn* is not satisfied” errors, due to incorrect selection of a “fake LSN” to represent an insertion point.
-  Fix underestimate of required size of DSA page maps for odd-size segments (Paul Bunn) [&sect;](https://postgr.es/c/2543b9ea9)

   This miscalculation led to out-of-bounds accesses and hence server crashes.
-  Fix indexing of oldest-multixact arrays in shared memory (Yura Sokolov) [&sect;](https://postgr.es/c/dcd9c06a4) [&sect;](https://postgr.es/c/969576dab)

   This mistake could cause a prepared-but-not-yet-committed transaction's row locks to appear invisible to other sessions, or other visibility issues for the results of such a transaction. With a very small max_connections setting, memory stomps were also possible.
-  Fix possible server crash when processing extended statistics on expressions of extension data types (Michael Paquier) [&sect;](https://postgr.es/c/530b6b02f)

   NULL pointer dereferences were possible if the data type's typanalyze function does not compute any useful statistics. No in-core typanalyze function behaves that way, but extensions could.
-  Fix minor memory leaks in ICU-based string processing (Jeff Davis) [&sect;](https://postgr.es/c/4761f2eee)
-  If the startup process fails, properly shut down other child processes before exiting the postmaster (Ayush Tiwari) [&sect;](https://postgr.es/c/e381843cf)

   The handling of this situation relied on a long-obsolete assumption that no other postmaster children exist while the startup process is running, so that immediate postmaster exit is acceptable. Orphaned children would eventually notice the postmaster's death and exit on their own, but a cleaner shutdown procedure is desirable.
-  Fix race condition between WAL replay of checkpoints and multixact ID creations (Heikki Linnakangas) [&sect;](https://postgr.es/c/1ca385032)

   A standby server following WAL from a primary of an older minor version could get into a crash-and-restart loop complaining about “could not access status of transaction”.
-  Prevent indefinite wait in shutdown of a walsender process (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/bbbc0888b) [&sect;](https://postgr.es/c/8ee536c89)

   At shutdown of a cluster that is publishing logical replication data, the walsender waits for all pending WAL to be written out. But it did not correctly request that to happen, so that in some cases this could become an indefinite wait.
-  Ensure that changes to tables' free space maps are persisted during recovery (Alexey Makhmutov) [&sect;](https://postgr.es/c/1cf010f21)

   Previously, while WAL replay did update the free space map while replaying operations that should change it, the map page buffer did not get marked dirty if checksums are enabled, so that the changes might never get written out. On a standby server, over time this would result in a map wildly at variance with the table's actual contents. While the map is only used as a hint, this condition could cause significant performance degradation for some period of time after the standby server is promoted to be active, until most of the map has been repaired by updates.
-  Fix crashes in some ecpg functions when called without any established connection (Shruthi Gowda) [&sect;](https://postgr.es/c/5d67549d9)
-  Fix assorted bugs in backup decompression and tar-parsing code (Andrew Dunstan, Tom Lane, Chao Li) [&sect;](https://postgr.es/c/f1298a4c2) [&sect;](https://postgr.es/c/8b198b093) [&sect;](https://postgr.es/c/2640c5ba7)

   The decompression and tar-file reading code used in pg_basebackup and pg_verifybackup mishandled tar-file padding data, could corrupt LZ4-compressed data in edge cases, failed to check for some unusual error conditions, failed to exit after compression/decompression errors (leading to cascading error reports), and leaked memory.
-  In pg_dumpall, don't skip role `GRANT`s with dangling grantor OIDs (Tom Lane) [&sect;](https://postgr.es/c/1cd783d20)

   Instead, handle such cases by emitting `GRANT` without any `GRANTED BY` clause, as we did before v16. This avoids losing the grant in foreseeable cases, since pre-v16 servers didn't prevent dropping the grantor role. Continue to emit a warning about the missing grantor, but only if the source server is v16 or later.
-  In pg_upgrade, take care to use the correct protocol version when connecting to older source servers (Jacob Champion) [&sect;](https://postgr.es/c/ad7fc3f1f)

   This could be problematic when attempting to upgrade from a pre-2018 server.
-  In `contrib/basic_archive`, allow the archive directory to be missing at startup (Nathan Bossart) [&sect;](https://postgr.es/c/f510577de)

   Previously, the setting of `basic_archive.archive_directory` was rejected if it didn't point to an existing directory. This is undesirable because archiving will be stuck indefinitely, even if the directory appears later.
-  Fix `contrib/ltree` to cope when case-folding changes a string's byte length (Jeff Davis) [&sect;](https://postgr.es/c/d1bd9a7dc)

   Previously, `lquery` patterns specifying case-insensitive matching might fail to match labels they should match.
-  In `contrib/pg_stat_statements`, don't leak memory if an error occurs while parsing the `pgss_query_texts.stat` file (Heikki Linnakangas) [&sect;](https://postgr.es/c/351e59f34)
-  In `contrib/postgres_fdw`, avoid crash due to premature cleanup of a failed connection (Etsuro Fujita) [&sect;](https://postgr.es/c/af8f9248f)

   If a remote connection fails abort cleanup, we can't use it any longer. But delay closing the connection object until end of transaction, because there might still be references to it within data structures such as open cursors.
-  Update time zone data files to tzdata release 2026b (Tom Lane) [&sect;](https://postgr.es/c/4c0eab6f0)

   British Columbia (America/Vancouver) will be on year-round UTC-07 (effectively, permanent DST) beginning in November 2026. This release assumes that their TZ abbreviation will be `MST` from that time forward. That seems likely to change, but it's unclear what new abbreviation will be used. Also a historical correction for Moldova: they have followed EU DST transition times since 2022.
