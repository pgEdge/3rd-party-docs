## Release 16.3 { #release-16-3 }


**Release date:.**


2024-05-09


 This release contains a variety of fixes from 16.2. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.3 { #release-16-3-migration }


 A dump/restore is not required for those running 16.X.


 However, a security vulnerability was found in the system views `pg_stats_ext` and `pg_stats_ext_exprs`, potentially allowing authenticated database users to see data they shouldn't. If this is of concern in your installation, follow the steps in the first changelog entry below to rectify it.


 Also, if you are upgrading from a version earlier than 16.2, see [Release 16.2](release-16-2.md#release-16-2).


### Changes { #release-16-3-changes }


-  Restrict visibility of `pg_stats_ext` and `pg_stats_ext_exprs` entries to the table owner (Nathan Bossart) [&sect;](https://postgr.es/c/2485a85e9)

   These views failed to hide statistics for expressions that involve columns the accessing user does not have permission to read. View columns such as `most_common_vals` might expose security-relevant data. The potential interactions here are not fully clear, so in the interest of erring on the side of safety, make rows in these views visible only to the owner of the associated table.

   The PostgreSQL Project thanks Lukas Fittl for reporting this problem. (CVE-2024-4317)

   By itself, this fix will only fix the behavior in newly initdb'd database clusters. If you wish to apply this change in an existing cluster, you will need to do the following:

1.  Find the SQL script `fix-CVE-2024-4317.sql` in the *share* directory of the PostgreSQL installation (typically located someplace like `/usr/share/postgresql/`). Be sure to use the script appropriate to your PostgreSQL major version. If you do not see this file, either your version is not vulnerable (only v14–v16 are affected) or your minor version is too old to have the fix.
2.  In *each* database of the cluster, run the `fix-CVE-2024-4317.sql` script as superuser. In psql this would look like

```

\i /usr/share/postgresql/fix-CVE-2024-4317.sql
```
      (adjust the file path as appropriate). Any error probably indicates that you've used the wrong script version. It will not hurt to run the script more than once.
3.  Do not forget to include the `template0` and `template1` databases, or the vulnerability will still exist in databases you create later. To fix `template0`, you'll need to temporarily make it accept connections. Do that with

```sql

ALTER DATABASE template0 WITH ALLOW_CONNECTIONS true;
```
      and then after fixing `template0`, undo it with

```sql

ALTER DATABASE template0 WITH ALLOW_CONNECTIONS false;
```

-  Fix `INSERT` from multiple `VALUES` rows into a target column that is a domain over an array or composite type (Tom Lane) [&sect;](https://postgr.es/c/52898c63e)

   Such cases would either fail with surprising complaints about mismatched datatypes, or insert unexpected coercions that could lead to odd results.
-  Require `SELECT` privilege on the target table for `MERGE` with a `DO NOTHING` clause (&Aacute;lvaro Herrera) [&sect;](https://postgr.es/c/a3f5d2056)

   `SELECT` privilege would be required in all practical cases anyway, but require it even if the query reads no columns of the target table. This avoids an edge case in which `MERGE` would require no privileges whatever, which seems undesirable even when it's a do-nothing command.
-  Fix handling of self-modified tuples in `MERGE` (Dean Rasheed) [&sect;](https://postgr.es/c/dd73d10ad)

   Throw an error if a target row joins to more than one source row, as required by the SQL standard. (The previous coding could silently ignore this condition if a concurrent update was involved.) Also, throw a non-misleading error if a target row is already updated by a later command in the current transaction, thanks to a `BEFORE` trigger or a volatile function used in the query.
-  Fix incorrect pruning of NULL partition when a table is partitioned on a boolean column and the query has a boolean `IS NOT` clause (David Rowley) [&sect;](https://postgr.es/c/fb95cc72b)

   A NULL value satisfies a clause such as <em>boolcol</em><code> IS NOT
      FALSE</code>, so pruning away a partition containing NULLs yielded incorrect answers.
-  Make `ALTER FOREIGN TABLE SET SCHEMA` move any owned sequences into the new schema (Tom Lane) [&sect;](https://postgr.es/c/7445f0928)

   Moving a regular table to a new schema causes any sequences owned by the table to be moved to that schema too (along with indexes and constraints). This was overlooked for foreign tables, however.
-  Make `ALTER TABLE ... ADD COLUMN` create identity/serial sequences with the same persistence as their owning tables (Peter Eisentraut) [&sect;](https://postgr.es/c/86d2b434c)

   `CREATE UNLOGGED TABLE` will make any owned sequences be unlogged too. `ALTER TABLE` missed that consideration, so that an added identity column would have a logged sequence, which seems pointless.
-  Improve `ALTER TABLE ... ALTER COLUMN TYPE`'s error message when there is a dependent function or publication (Tom Lane) [&sect;](https://postgr.es/c/11d40a44b) [&sect;](https://postgr.es/c/b19255ca6)
-  In `CREATE DATABASE`, recognize strategy keywords case-insensitively for consistency with other options (Tomas Vondra) [&sect;](https://postgr.es/c/9e6faeb32)
-  Fix `EXPLAIN`'s counting of heap pages accessed by a bitmap heap scan (Melanie Plageman) [&sect;](https://postgr.es/c/1f4eb7342)

   Previously, heap pages that contain no visible tuples were not counted; but it seems more consistent to count all pages returned by the bitmap index scan.
-  Fix `EXPLAIN`'s output for subplans in `MERGE` (Dean Rasheed) [&sect;](https://postgr.es/c/34c854b93)

   `EXPLAIN` would sometimes fail to properly display subplan Params referencing variables in other parts of the plan tree.
-  Avoid deadlock during removal of orphaned temporary tables (Mikhail Zhilin) [&sect;](https://postgr.es/c/cbfbb14bd)

   If the session that creates a temporary table crashes without removing the table, autovacuum will eventually try to remove the orphaned table. However, an incoming session that's been assigned the same temporary namespace will do that too. If a temporary table has a dependency (such as an owned sequence) then a deadlock could result between these two cleanup attempts.
-  Fix updating of visibility map state in `VACUUM` with the `DISABLE_PAGE_SKIPPING` option (Heikki Linnakangas) [&sect;](https://postgr.es/c/407cb6c65)

   Due to an oversight, this mode caused all heap pages to be dirtied, resulting in excess I/O. Also, visibility map bits that were incorrectly set would not get cleared.
-  Avoid race condition while examining per-relation frozen-XID values (Noah Misch) [&sect;](https://postgr.es/c/92685c389)

   `VACUUM`'s computation of per-database frozen-XID values from per-relation values could get confused by a concurrent update of those values by another `VACUUM`.
-  Fix buffer usage reporting for parallel vacuuming (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/f199436c1)

   Buffer accesses performed by parallel workers were not getting counted in the statistics reported in `VERBOSE` mode.
-  Ensure that join conditions generated from equivalence classes are applied at the correct plan level (Tom Lane) [&sect;](https://postgr.es/c/91800af13)

   In versions before PostgreSQL 16, it was possible for generated conditions to be evaluated below outer joins when they should be evaluated above (after) the outer join, leading to incorrect query results. All versions have a similar hazard when considering joins to `UNION ALL` trees that have constant outputs for the join column in some `SELECT ` arms.
-  Fix “could not find pathkey item to sort” errors occurring while planning aggregate functions with `ORDER BY` or `DISTINCT` options (David Rowley) [&sect;](https://postgr.es/c/4e1ff2aad)

   This is similar to a fix applied in 16.1, but it solves the problem for parallel plans.
-  Prevent potentially-incorrect optimization of some window functions (David Rowley) [&sect;](https://postgr.es/c/9d36b883b)

   Disable “run condition” optimization of `ntile()` and `count()` with non-constant arguments. This avoids possible misbehavior with sub-selects, typically leading to errors like “WindowFunc not found in subplan target lists”.
-  Avoid unnecessary use of moving-aggregate mode with a non-moving window frame (Vallimaharajan G) [&sect;](https://postgr.es/c/a94f51a7b)

   When a plain aggregate is used as a window function, and the window frame start is specified as `UNBOUNDED PRECEDING`, the frame's head cannot move so we do not need to use the special (and more expensive) moving-aggregate mode. This optimization was intended all along, but due to a coding error it never triggered.
-  Avoid use of already-freed data while planning partition-wise joins under GEQO (Tom Lane) [&sect;](https://postgr.es/c/ef0333e67)

   This would typically end in a crash or unexpected error message.
-  Avoid freeing still-in-use data in Memoize (Tender Wang, Andrei Lepikhov) [&sect;](https://postgr.es/c/348233cb1)

   In production builds this error frequently didn't cause any problems, as the freed data would most likely not get overwritten before it was used.
-  Fix incorrectly-reported statistics kind codes in “requested statistics kind *X* is not yet built” error messages (David Rowley) [&sect;](https://postgr.es/c/ac7e6a01c)
-  Use a hash table instead of linear search for “catcache list” objects (Tom Lane) [&sect;](https://postgr.es/c/14e991db8)

   This change solves performance problems that were reported for certain operations in installations with many thousands of roles.
-  Be more careful with `RECORD`-returning functions in `FROM` (Tom Lane) [&sect;](https://postgr.es/c/cc1eb6a3c) [&sect;](https://postgr.es/c/1b3029be5)

   The output columns of such a function call must be defined by an `AS` clause that specifies the column names and data types. If the actual function output value doesn't match that, an error is supposed to be thrown at runtime. However, some code paths would examine the actual value prematurely, and potentially issue strange errors or suffer assertion failures if it doesn't match expectations.
-  Fix confusion about the return rowtype of SQL-language procedures (Tom Lane) [&sect;](https://postgr.es/c/40d1bdeb7)

   A procedure implemented in SQL language that returns a single composite-type column would cause an assertion failure or core dump.
-  Add protective stack depth checks to some recursive functions (Egor Chindyaskin) [&sect;](https://postgr.es/c/760767182)
-  Fix mis-rounding and overflow hazards in `date_bin()` (Moaaz Assali) [&sect;](https://postgr.es/c/17db5436e)

   In the case where the source timestamp is before the origin timestamp and their difference is already an exact multiple of the stride, the code incorrectly subtracted the stride anyway. Also, detect some integer-overflow cases that would have produced incorrect results.
-  Detect integer overflow when adding or subtracting an `interval` to/from a `timestamp` (Joseph Koshakow) [&sect;](https://postgr.es/c/3752e3d21)

   Some cases that should cause an out-of-range error produced an incorrect result instead.
-  Avoid race condition in `pg_get_expr()` (Tom Lane) [&sect;](https://postgr.es/c/4eb261165)

   If the relation referenced by the argument is dropped concurrently, the function's intention is to return NULL, but sometimes it failed instead.
-  Fix detection of old transaction IDs in XID status functions (Karina Litskevich) [&sect;](https://postgr.es/c/e3e05adde)

   Transaction IDs more than 2<sup>31</sup> transactions in the past could be misidentified as recent, leading to misbehavior of `pg_xact_status()` or `txid_status()`.
-  Ensure that a table's freespace map won't return a page that's past the end of the table (Ronan Dunklau) [&sect;](https://postgr.es/c/4e62ba21a)

   Because the freespace map isn't WAL-logged, this was possible in edge cases involving an OS crash, a replica promote, or a PITR restore. The result would be a “could not read block” error.
-  Fix file descriptor leakage when an error is thrown while waiting in `WaitEventSetWait` (Etsuro Fujita) [&sect;](https://postgr.es/c/e79ceafe9)
-  Avoid corrupting exception stack if an FDW implements async append but doesn't configure any wait conditions for the Append plan node to wait for (Alexander Pyhalov) [&sect;](https://postgr.es/c/f6f61a4bd)
-  Throw an error if an index is accessed while it is being reindexed (Tom Lane) [&sect;](https://postgr.es/c/8c785d354)

   Previously this was just an assertion check, but promote it into a regular runtime error. This will provide a more on-point error message when reindexing a user-defined index expression that attempts to access its own table.
-  Ensure that index-only scans on `name` columns return a fully-padded value (David Rowley) [&sect;](https://postgr.es/c/68d358545)

   The value physically stored in the index is truncated, and previously a pointer to that value was returned to callers. This provoked complaints when testing under valgrind. In theory it could result in crashes, though none have been reported.
-  Fix race condition that could lead to reporting an incorrect conflict cause when invalidating a replication slot (Bertrand Drouvot) [&sect;](https://postgr.es/c/59cea09f0)
-  Fix race condition in deciding whether a table sync operation is needed in logical replication (Vignesh C) [&sect;](https://postgr.es/c/a9155efc7)

   An invalidation event arriving while a subscriber identifies which tables need to be synced would be forgotten about, so that any tables newly in need of syncing might not get processed in a timely fashion.
-  Fix crash with DSM allocations larger than 4GB (Heikki Linnakangas) [&sect;](https://postgr.es/c/f2f09b825)
-  Disconnect if a new server session's client socket cannot be put into non-blocking mode (Heikki Linnakangas) [&sect;](https://postgr.es/c/539e328b1)

   It was once theoretically possible for us to operate with a socket that's in blocking mode; but that hasn't worked fully in a long time, so fail at connection start rather than misbehave later.
-  Fix inadequate error reporting with OpenSSL 3.0.0 and later (Heikki Linnakangas, Tom Lane) [&sect;](https://postgr.es/c/6a2c80e95)

   System-reported errors passed through by OpenSSL were reported with a numeric error code rather than anything readable.
-  Fix thread-safety of error reporting for `getaddrinfo()` on Windows (Thomas Munro) [&sect;](https://postgr.es/c/0460e4ecc)

   A multi-threaded libpq client program could get an incorrect or corrupted error message after a network lookup failure.
-  Avoid concurrent calls to `bindtextdomain()` in libpq and ecpglib (Tom Lane) [&sect;](https://postgr.es/c/52afe5632) [&sect;](https://postgr.es/c/9440d23a0)

   Although GNU gettext's implementation seems to be fine with concurrent calls, the version available on Windows is not.
-  Fix crash in ecpg's preprocessor if the program tries to redefine a macro that was defined on the preprocessor command line (Tom Lane) [&sect;](https://postgr.es/c/392e6e9e6) [&sect;](https://postgr.es/c/0018f0af5) [&sect;](https://postgr.es/c/dd3fddc85)
-  In ecpg, avoid issuing false “unsupported feature will be passed to server” warnings (Tom Lane) [&sect;](https://postgr.es/c/118558e6d)
-  Ensure that the string result of ecpg's `intoasc()` function is correctly zero-terminated (Oleg Tselebrovskiy) [&sect;](https://postgr.es/c/88e03d055)
-  In initdb's `-c` option, match parameter names case-insensitively (Tom Lane) [&sect;](https://postgr.es/c/b78f4d22b)

   The server treats parameter names case-insensitively, so this code should too. This avoids putting redundant entries into the generated `postgresql.conf` file.
-  In psql, avoid leaking a query result after the query is cancelled (Tom Lane) [&sect;](https://postgr.es/c/a85e3ba1c)

   This happened only when cancelling a non-last query in a query string made with `\;` separators.
-  Fix pg_dumpall so that role comments, if present, will be dumped regardless of the setting of `--no-role-passwords` (Daniel Gustafsson, &Aacute;lvaro Herrera) [&sect;](https://postgr.es/c/5863bacb8)
-  Skip files named `.DS_Store` in pg_basebackup, pg_checksums, and pg_rewind (Daniel Gustafsson) [&sect;](https://postgr.es/c/103235888)

   This avoids problems on macOS, where the Finder may create such files.
-  Fix PL/pgSQL's parsing of single-line comments (`--`-style comments) following expressions (Erik Wienhold, Tom Lane) [&sect;](https://postgr.es/c/48f216dc6)

   This mistake caused parse errors if such a comment followed a `WHEN` expression in a PL/pgSQL `CASE` statement.
-  In `contrib/amcheck`, don't report false match failures due to short- versus long-header values (Andrey Borodin, Michael Zhilin) [&sect;](https://postgr.es/c/3676b846b) [&sect;](https://postgr.es/c/a6ddb8ad0)

   A variable-length datum in a heap tuple or index tuple could have either a short or a long header, depending on compression parameters that applied when it was made. Treat these cases as equivalent rather than complaining if there's a difference.
-  Fix bugs in BRIN output functions (Tomas Vondra) [&sect;](https://postgr.es/c/8cea358b1) [&sect;](https://postgr.es/c/ccd8f0fa1)

   These output functions are only used for displaying index entries in `contrib/pageinspect`, so the errors are of limited practical concern.
-  In `contrib/postgres_fdw`, avoid emitting requests to sort by a constant (David Rowley) [&sect;](https://postgr.es/c/6a9e2cb2b)

   This could occur in cases involving `UNION ALL` with constant-emitting subqueries. Sorting by a constant is useless of course, but it also risks being misinterpreted by the remote server, leading to “ORDER BY position *N* is not in select list” errors.
-  Make `contrib/postgres_fdw` set the remote session's time zone to `GMT` not `UTC` (Tom Lane) [&sect;](https://postgr.es/c/75929b6cf)

   This should have the same results for practical purposes. However, `GMT` is recognized by hard-wired code in the server, while `UTC` is looked up in the timezone database. So the old code could fail in the unlikely event that the remote server's timezone database is missing entries.
-  In `contrib/xml2`, avoid use of library functions that have been deprecated in recent versions of libxml2 (Dmitry Koval) [&sect;](https://postgr.es/c/7c93f31de)
-  Fix incompatibility with LLVM 18 (Thomas Munro, Dmitry Dolgov) [&sect;](https://postgr.es/c/bf1cfe77e)
-  Allow `make check` to work with the musl C library (Thomas Munro, Bruce Momjian, Tom Lane) [&sect;](https://postgr.es/c/7651fd387)
