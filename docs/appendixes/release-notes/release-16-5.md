## Release 16.5 { #release-16-5 }


**Release date:.**


2024-11-14


 This release contains a variety of fixes from 16.4. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.5 { #release-16-5-migration }


 A dump/restore is not required for those running 16.X.


 However, if you have ever detached a partition from a partitioned table that has a foreign-key reference to another partitioned table, and not dropped the former partition, then you may have catalog and/or data corruption to repair, as detailed in the fifth changelog entry below.


 Also, if you are upgrading from a version earlier than 16.3, see [Release 16.3](release-16-3.md#release-16-3).


### Changes { #release-16-5-changes }


-  Ensure cached plans are marked as dependent on the calling role when RLS applies to a non-top-level table reference (Nathan Bossart) [&sect;](https://postgr.es/c/562289460)

   If a CTE, subquery, sublink, security invoker view, or coercion projection in a query references a table with row-level security policies, we neglected to mark the resulting plan as potentially dependent on which role is executing it. This could lead to later query executions in the same session using the wrong plan, and then returning or hiding rows that should have been hidden or returned instead.

   The PostgreSQL Project thanks Wolfgang Walther for reporting this problem. (CVE-2024-10976)
-  Make libpq discard error messages received during SSL or GSS protocol negotiation (Jacob Champion) [&sect;](https://postgr.es/c/67d28bd02)

   An error message received before encryption negotiation is completed might have been injected by a man-in-the-middle, rather than being real server output. Reporting it opens the door to various security hazards; for example, the message might spoof a query result that a careless user could mistake for correct output. The best answer seems to be to discard such data and rely only on libpq's own report of the connection failure.

   The PostgreSQL Project thanks Jacob Champion for reporting this problem. (CVE-2024-10977)
-  Fix unintended interactions between `SET SESSION AUTHORIZATION` and `SET ROLE` (Tom Lane) [&sect;](https://postgr.es/c/ae340d031) [&sect;](https://postgr.es/c/95f5a5237)

   The SQL standard mandates that `SET SESSION AUTHORIZATION` have a side-effect of doing `SET ROLE NONE`. Our implementation of that was flawed, creating more interaction between the two settings than intended. Notably, rolling back a transaction that had done `SET SESSION AUTHORIZATION` would revert `ROLE` to `NONE` even if that had not been the previous state, so that the effective user ID might now be different from what it had been before the transaction. Transiently setting `session_authorization` in a function `SET` clause had a similar effect. A related bug was that if a parallel worker inspected `current_setting('role')`, it saw `none` even when it should see something else.

   The PostgreSQL Project thanks Tom Lane for reporting this problem. (CVE-2024-10978)
-  Prevent trusted PL/Perl code from changing environment variables (Andrew Dunstan, Noah Misch) [&sect;](https://postgr.es/c/8fe3e697a) [&sect;](https://postgr.es/c/88269df4d) [&sect;](https://postgr.es/c/168579e23) [&sect;](https://postgr.es/c/c335264c9) [&sect;](https://postgr.es/c/64df88700)

   The ability to manipulate process environment variables such as `PATH` gives an attacker opportunities to execute arbitrary code. Therefore, “trusted” PLs must not offer the ability to do that. To fix `plperl`, replace `%ENV` with a tied hash that rejects any modification attempt with a warning. Untrusted `plperlu` retains the ability to change the environment.

   The PostgreSQL Project thanks Coby Abrams for reporting this problem. (CVE-2024-10979)
-  Fix updates of catalog state for foreign-key constraints when attaching or detaching table partitions (Jehan-Guillaume de Rorthais, Tender Wang, Álvaro Herrera) [&sect;](https://postgr.es/c/2aaf2a28b) [&sect;](https://postgr.es/c/f7d510a38)

   If the referenced table is partitioned, then different catalog entries are needed for a referencing table that is stand-alone versus one that is a partition. `ATTACH/DETACH PARTITION` commands failed to perform this conversion correctly. In particular, after `DETACH` the now stand-alone table would be missing foreign-key enforcement triggers, which could result in the table later containing rows that fail the foreign-key constraint. A subsequent re-`ATTACH` could fail with surprising errors, too.

   The way to fix this is to do `ALTER TABLE DROP CONSTRAINT` on the now stand-alone table for each faulty constraint, and then re-add the constraint. If re-adding the constraint fails, then some erroneous data has crept in. You will need to manually re-establish consistency between the referencing and referenced tables, then re-add the constraint.

   This query can be used to identify broken constraints and construct the commands needed to recreate them:

```sql

SELECT conrelid::pg_catalog.regclass AS "constrained table",
       conname AS constraint,
       confrelid::pg_catalog.regclass AS "references",
       pg_catalog.format('ALTER TABLE %s DROP CONSTRAINT %I;',
                         conrelid::pg_catalog.regclass, conname) AS "drop",
       pg_catalog.format('ALTER TABLE %s ADD CONSTRAINT %I %s;',
                         conrelid::pg_catalog.regclass, conname,
                         pg_catalog.pg_get_constraintdef(oid)) AS "add"
FROM pg_catalog.pg_constraint c
WHERE contype = 'f' AND conparentid = 0 AND
   (SELECT count(*) FROM pg_catalog.pg_constraint c2
    WHERE c2.conparentid = c.oid) <>
   ((SELECT count(*) FROM pg_catalog.pg_inherits i
    WHERE (i.inhparent = c.conrelid OR i.inhparent = c.confrelid) AND
      EXISTS (SELECT 1 FROM pg_catalog.pg_partitioned_table
              WHERE partrelid = i.inhparent)) +
    CASE WHEN pg_catalog.pg_partition_root(conrelid) = confrelid THEN
              (SELECT count(*) FROM pg_catalog.pg_partition_tree(confrelid)
                WHERE level = 1)
         ELSE 0 END);
```
   Since it is possible that one or more of the `ADD CONSTRAINT` steps will fail, you should save the query's output in a file and then attempt to perform each step.
-  Avoid possible crashes and “could not open relation” errors in queries on a partitioned table occurring concurrently with a `DETACH CONCURRENTLY` and immediate drop of a partition (Álvaro Herrera, Kuntal Gosh) [&sect;](https://postgr.es/c/a6ff329e7) [&sect;](https://postgr.es/c/1b9dd6b05)
-  Disallow `ALTER TABLE ATTACH PARTITION` if the table to be attached has a foreign key referencing the partitioned table (Álvaro Herrera) [&sect;](https://postgr.es/c/ada34d714) [&sect;](https://postgr.es/c/57c8b8726)

   This arrangement is not supported, and other ways of creating it already fail.
-  Don't use partitionwise joins or grouping if the query's collation for the key column doesn't match the partition key's collation (Jian He, Webbo Han) [&sect;](https://postgr.es/c/f734b6b4d) [&sect;](https://postgr.es/c/dd2f8ebee)

   Such plans could produce incorrect results.
-  Fix possible “could not find pathkey item to sort” error when the output of a `UNION ALL` member query needs to be sorted, and the sort column is an expression (Andrei Lepikhov, Tom Lane) [&sect;](https://postgr.es/c/64635c8af)
-  Fix performance regressions involving flattening of subqueries underneath outer joins that are later reduced to plain joins (Tom Lane) [&sect;](https://postgr.es/c/80d9c07a4)

   v16 failed to optimize some queries as well as prior versions had, because of overoptimistic simplification of query-pullup logic.
-  Allow cancellation of the second stage of index build for large hash indexes (Pavel Borisov) [&sect;](https://postgr.es/c/d23109f4b)
-  Fix assertion failure or confusing error message for <code>COPY
      (</code><em>query</em><code>) TO ...</code>, when the *query* is rewritten by a `DO INSTEAD NOTIFY` rule (Tender Wang, Tom Lane) [&sect;](https://postgr.es/c/6c3b2d204)
-  Fix server crash when a `json_objectagg()` call contains a volatile function (Amit Langote) [&sect;](https://postgr.es/c/fa4f11854)
-  Fix checking of key uniqueness in JSON object constructors (Junwang Zhao, Tomas Vondra) [&sect;](https://postgr.es/c/8e65d9ff9)

   When building an object larger than a kilobyte, it was possible to accept invalid input that includes duplicate object keys, or to falsely report that duplicate keys are present.
-  Fix detection of skewed data during parallel hash join (Thomas Munro) [&sect;](https://postgr.es/c/53edc9485)

   After repartitioning the inner side of a hash join because one partition has accumulated too many tuples, we check to see if all the partition's tuples went into the same child partition, which suggests that they all have the same hash value and further repartitioning cannot improve matters. This check malfunctioned in some cases, allowing repeated futile repartitioning which would eventually end in a resource-exhaustion error.
-  Disallow locale names containing non-ASCII characters (Thomas Munro) [&sect;](https://postgr.es/c/ce17de580)

   This is only an issue on Windows, as such locale names are not used elsewhere. They are problematic because it's quite unclear what encoding such names are represented in (since the locale itself defines the encoding to use). In recent PostgreSQL releases, an abort in the Windows runtime library could occur because of confusion about that.

   Anyone who encounters the new error message should either create a new duplicated locale with an ASCII-only name using Windows Locale Builder, or consider using BCP 47-compliant locale names like `tr-TR`.
-  Fix race condition in committing a serializable transaction (Heikki Linnakangas) [&sect;](https://postgr.es/c/22665f210)

   Mis-processing of a recently committed transaction could lead to an assertion failure or a “could not access status of transaction” error.
-  Fix race condition in `COMMIT PREPARED` that resulted in orphaned 2PC files (wuchengwen) [&sect;](https://postgr.es/c/7de9b64a5)

   A concurrent `PREPARE TRANSACTION` could cause `COMMIT PREPARED` to not remove the on-disk two-phase state file for the completed transaction. There was no immediate ill effect, but a subsequent crash-and-recovery could fail with “could not access status of transaction”, requiring manual removal of the orphaned file to restore service.
-  Avoid invalid memory accesses after skipping an invalid toast index during `VACUUM FULL` (Tender Wang) [&sect;](https://postgr.es/c/afbd3dc7d)

   A list tracking yet-to-be-rebuilt indexes was not properly updated in this code path, risking assertion failures or crashes later on.
-  Fix ways in which an “in place” catalog update could be lost (Noah Misch) [&sect;](https://postgr.es/c/63f019805) [&sect;](https://postgr.es/c/51ff46de2) [&sect;](https://postgr.es/c/4c922821e) [&sect;](https://postgr.es/c/2d63c964f) [&sect;](https://postgr.es/c/370bc7740) [&sect;](https://postgr.es/c/6c837c237) [&sect;](https://postgr.es/c/f8f9110b4)

   Normal row updates write a new version of the row to preserve rollback-ability of the transaction. However, certain system catalog updates are intentionally non-transactional and are done with an in-place update of the row. These patches fix race conditions that could cause the effects of an in-place update to be lost. As an example, it was possible to forget having set `pg_class`.`relhasindex` to true, preventing updates of the new index and thus causing index corruption.
-  Reset catalog caches at end of recovery (Noah Misch) [&sect;](https://postgr.es/c/d36b4d8ec)

   This prevents scenarios wherein an in-place catalog update could be lost due to using stale data from a catalog cache.
-  Avoid using parallel query while holding off interrupts (Francesco Degrassi, Noah Misch, Tom Lane) [&sect;](https://postgr.es/c/6f6521de9) [&sect;](https://postgr.es/c/06424e9a2)

   This situation cannot arise normally, but it can be reached with test scenarios such as using a SQL-language function as B-tree support (which would be far too slow for production usage). If it did occur it would result in an indefinite wait.
-  Report the active query ID for statistics purposes at the start of processing of Bind and Execute protocol messages (Sami Imseih) [&sect;](https://postgr.es/c/21aad4bea)

   This allows more of the work done in extended query protocol to be attributed to the correct query.
-  Guard against stack overflow in libxml2 with too-deeply-nested XML input (Tom Lane, with hat tip to Nick Wellnhofer) [&sect;](https://postgr.es/c/4c9bf947a)

   Use `xmlXPathCtxtCompile()` rather than `xmlXPathCompile()`, because the latter fails to protect itself against recursion-to-stack-overflow in libxml2 releases before 2.13.4.
-  Fix some whitespace issues in the result of `XMLSERIALIZE(... INDENT)` (Jim Jones) [&sect;](https://postgr.es/c/06c285018)

   Fix failure to indent nodes separated by whitespace, and ensure that a trailing newline is not added.
-  Do not ignore a concurrent `REINDEX CONCURRENTLY` that is working on an index with predicates or expressions (Michail Nikolaev) [&sect;](https://postgr.es/c/edb0f6e41)

   Normally, `REINDEX CONCURRENTLY` does not need to wait for other `REINDEX CONCURRENTLY` operations on other tables. However, this optimization is not applied if the other `REINDEX CONCURRENTLY` is processing an index with predicates or expressions, on the chance that such expressions contain user-defined code that accesses other tables. Careless coding created a race condition such that that rule was not applied uniformly, possibly allowing inconsistent behavior.
-  Fix mis-deparsing of `ORDER BY` lists when there is a name conflict (Tom Lane) [&sect;](https://postgr.es/c/9fe6319dc)

   If an `ORDER BY` item in `SELECT` is a bare identifier, the parser first seeks it as an output column name of the `SELECT`, for SQL92 compatibility. However, ruleutils.c expects the SQL99 interpretation where such a name is an input column name. So it was possible to produce an incorrect display of a view in the (rather ill-advised) case where some other column is renamed in the `SELECT` output list to match an input column used in `ORDER BY`. Fix by table-qualifying such names in the dumped view text.
-  Fix “failed to find plan for subquery/CTE” errors in `EXPLAIN` (Richard Guo, Tom Lane) [&sect;](https://postgr.es/c/9db6650a5) [&sect;](https://postgr.es/c/03f679475)

   This case arose while trying to print references to fields of a RECORD-type output of a subquery when the subquery has been optimized out of the plan altogether (which is possible at least in the case that it has a constant-false `WHERE` condition). Nothing remains in the plan to identify the original field names, so fall back to printing <code>f</code><em>N</em> for the *N*'th record column. (That's actually the right thing anyway, if the record output arose from a `ROW()` constructor.)
-  Disallow a `USING` clause when altering the type of a generated column (Peter Eisentraut) [&sect;](https://postgr.es/c/5867ee005)

   A generated column already has an expression specifying the column contents, so including `USING` doesn't make sense.
-  Ignore not-yet-defined Portals in the `pg_cursors` view (Tom Lane) [&sect;](https://postgr.es/c/5de77b609)

   It is possible for user-defined code that inspects this view to be called while a new cursor is being set up, and if that happens a null pointer dereference would ensue. Avoid the problem by defining the view to exclude incompletely-set-up cursors.
-  Fix incorrect output of the `pg_stat_io` view on 32-bit machines (Bertrand Drouvot) [&sect;](https://postgr.es/c/dd20f950d)

   The `stats_reset` timestamp column contained garbage on such hardware.
-  Prevent mis-encoding of “trailing junk after numeric literal” error messages (Karina Litskevich) [&sect;](https://postgr.es/c/4fd4d7653)

   We do not allow identifiers to appear immediately following numeric literals (there must be some whitespace between). If a multibyte character immediately followed a numeric literal, the syntax error message about it included only the first byte of that character, causing bad-encoding problems both in the report to the client and in the postmaster log file.
-  Avoid “unexpected table_index_fetch_tuple call during logical decoding” error while decoding a transaction involving insertion of a column default value (Takeshi Ideriha, Hou Zhijie) [&sect;](https://postgr.es/c/0f0e253db) [&sect;](https://postgr.es/c/0c40d9019)
-  Reduce memory consumption of logical decoding (Masahiko Sawada) [&sect;](https://postgr.es/c/05e982cdc)

   Use a smaller default block size to store tuple data received during logical replication. This reduces memory wastage, which has been reported to be severe while processing long-running transactions, even leading to out-of-memory failures.
-  In a logical replication apply worker, ensure that origin progress is not advanced during an error or apply worker shutdown (Hayato Kuroda, Shveta Malik) [&sect;](https://postgr.es/c/b39c5272c)

   This avoids possible loss of a transaction, since once the origin progress point is advanced the source server won't send that data again.
-  Re-disable sending of stateless (TLSv1.2) session tickets (Daniel Gustafsson) [&sect;](https://postgr.es/c/9333174af)

   A previous change to prevent sending of stateful (TLSv1.3) session tickets accidentally re-enabled sending of stateless ones. Thus, while we intended to prevent clients from thinking that TLS session resumption is supported, some still did.
-  Avoid “wrong tuple length” failure when dropping a database with many ACL (permission) entries (Ayush Tiwari) [&sect;](https://postgr.es/c/545794515) [&sect;](https://postgr.es/c/f6991cafa)
-  Allow adjusting the `session_authorization` and `role` settings in parallel workers (Tom Lane) [&sect;](https://postgr.es/c/f3ab5d3a2)

   Our code intends to allow modifiable server settings to be set by function `SET` clauses, but not otherwise within a parallel worker. `SET` clauses failed for these two settings, though.
-  Fix behavior of stable functions called from a `CALL` statement's argument list, when the `CALL` is within a PL/pgSQL `EXCEPTION` block (Tom Lane) [&sect;](https://postgr.es/c/25d639eea)

   As with a similar fix in our previous quarterly releases, this case allowed such functions to be passed the wrong snapshot, causing them to see stale values of rows modified since the start of the outer transaction.
-  Fix “cache lookup failed for function” errors in edge cases in PL/pgSQL's `CALL` (Tom Lane) [&sect;](https://postgr.es/c/a073835c1)
-  Fix thread safety of our fallback (non-OpenSSL) MD5 implementation on big-endian hardware (Heikki Linnakangas) [&sect;](https://postgr.es/c/0583863e9)

   Thread safety is not currently a concern in the server, but it is for libpq.
-  Parse libpq's `keepalives` connection option in the same way as other integer-valued options (Yuto Sasaki) [&sect;](https://postgr.es/c/65f431aff)

   The coding used here rejected trailing whitespace in the option value, unlike other cases. This turns out to be problematic in ecpg's usage, for example.
-  Avoid use of `pnstrdup()` in ecpglib (Jacob Champion) [&sect;](https://postgr.es/c/ee2997c67)

   That function will call `exit()` on out-of-memory, which is undesirable in a library. The calling code already handles allocation failures properly.
-  In ecpglib, fix out-of-bounds read when parsing incorrect datetime input (Bruce Momjian, Pavel Nekrasov) [&sect;](https://postgr.es/c/a1e613b81)

   It was possible to try to read the location just before the start of a constant array. Real-world consequences seem minimal, though.
-  Fix memory leak in psql during repeated use of `\bind` (Michael Paquier) [&sect;](https://postgr.es/c/c2fb2f9e2)
-  Avoid hanging if an interval less than 1ms is specified in psql's `\watch` command (Andrey Borodin, Michael Paquier) [&sect;](https://postgr.es/c/6331972c7)

   Instead, treat this the same as an interval of zero (no wait between executions).
-  Fix pg_dump's handling of identity sequences that have persistence different from their owning table's persistence (Tom Lane) [&sect;](https://postgr.es/c/b8b175a4c)

   Since v15, it's been possible to set an identity sequence to be LOGGED when its owning table is UNLOGGED or vice versa. However, pg_dump's method for recreating that situation failed in binary-upgrade mode, causing pg_upgrade to fail when such sequences are present. Fix by introducing a new option for `ADD/ALTER COLUMN GENERATED AS IDENTITY` to allow the sequence's persistence to be set correctly at creation. Note that this means a dump from a database containing such a sequence will only load into a server of this minor version or newer.
-  Include the source timeline history in pg_rewind's debug output (Heikki Linnakangas) [&sect;](https://postgr.es/c/e8240dbd8)

   This was the intention to begin with, but a coding error caused the source history to always print as empty.
-  Avoid trying to reindex temporary tables and indexes in vacuumdb and in parallel reindexdb (VaibhaveS, Michael Paquier, Fujii Masao, Nathan Bossart) [&sect;](https://postgr.es/c/1ea4d9c00) [&sect;](https://postgr.es/c/653ce5b8b) [&sect;](https://postgr.es/c/eba8cc1af)

   Reindexing other sessions' temporary tables cannot work, but the check to skip them was missing in some code paths, leading to unwanted failures.
-  Allow inspection of sequence relations in relevant functions of `contrib/pageinspect` and `contrib/pgstattuple` (Nathan Bossart, Ayush Vatsa) [&sect;](https://postgr.es/c/2bd4c06bb) [&sect;](https://postgr.es/c/0938a4ecd)

   This had been allowed in the past, but it got broken during the introduction of non-default access methods for tables.
-  Fix incorrect LLVM-generated code on ARM64 platforms (Thomas Munro, Anthonin Bonnefoy) [&sect;](https://postgr.es/c/ee67b73f5)

   When using JIT compilation on ARM platforms, the generated code could not support relocation distances exceeding 32 bits, allowing unlucky placement of generated code to cause server crashes on large-memory systems.
-  Fix a few places that assumed that process start time (represented as a `time_t`) will fit into a `long` value (Max Johnson, Nathan Bossart) [&sect;](https://postgr.es/c/8aaf88b63)

   On platforms where `long` is 32 bits (notably Windows), this coding would fail after Y2038. Most of the failures appear only cosmetic, but notably `pg_ctl start` would hang.
-  Fix building with Strawberry Perl on Windows (Andrew Dunstan) [&sect;](https://postgr.es/c/0a0db4631)
-  Update time zone data files to tzdata release 2024b (Tom Lane) [&sect;](https://postgr.es/c/a0c8d600b) [&sect;](https://postgr.es/c/2abc88958)

   This tzdata release changes the old System-V-compatibility zone names to duplicate the corresponding geographic zones; for example `PST8PDT` is now an alias for `America/Los_Angeles`. The main visible consequence is that for timestamps before the introduction of standardized time zones, the zone is considered to represent local mean solar time for the named location. For example, in `PST8PDT`, `timestamptz` input such as `1801-01-01 00:00` would previously have been rendered as `1801-01-01 00:00:00-08`, but now it is rendered as `1801-01-01 00:00:00-07:52:58`.

   Also, historical corrections for Mexico, Mongolia, and Portugal. Notably, `Asia/Choibalsan` is now an alias for `Asia/Ulaanbaatar` rather than being a separate zone, mainly because the differences between those zones were found to be based on untrustworthy data.
