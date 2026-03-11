## Release 16.4 { #release-16-4 }


**Release date:.**


2024-08-08


 This release contains a variety of fixes from 16.3. For information about new features in major release 16, see [Release 16](release-16.md#release-16).


### Migration to Version 16.4 { #release-16-4-migration }


 A dump/restore is not required for those running 16.X.


 However, if you are upgrading from a version earlier than 16.3, see [Release 16.3](release-16-3.md#release-16-3).


### Changes { #release-16-4-changes }


-  Prevent unauthorized code execution during pg_dump (Masahiko Sawada) [&sect;](https://postgr.es/c/6aba85a4b)

   An attacker able to create and drop non-temporary objects could inject SQL code that would be executed by a concurrent pg_dump session with the privileges of the role running pg_dump (which is often a superuser). The attack involves replacing a sequence or similar object with a view or foreign table that will execute malicious code. To prevent this, introduce a new server parameter `restrict_nonsystem_relation_kind` that can disable expansion of non-builtin views as well as access to foreign tables, and teach pg_dump to set it when available. Note that the attack is prevented only if both pg_dump and the server it is dumping from are new enough to have this fix.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2024-7348)
-  Avoid incorrect results from Merge Right Anti Join plans (Richard Guo) [&sect;](https://postgr.es/c/507f2347e)

   If the inner relation is known to have unique join keys, the merge could misbehave when there are duplicated join keys in the outer relation.
-  Prevent infinite loop in `VACUUM` (Melanie Plageman) [&sect;](https://postgr.es/c/06bf404cd)

   After a disconnected standby server with an old running transaction reconnected to the primary, it was possible for `VACUUM` on the primary to get confused about which tuples are removable, resulting in an infinite loop.
-  Fix failure after attaching a table as a partition, if the table had previously had inheritance children (&Aacute;lvaro Herrera) [&sect;](https://postgr.es/c/084814d88)
-  Fix `ALTER TABLE DETACH PARTITION` for cases involving inconsistent index-based constraints (&Aacute;lvaro Herrera, Tender Wang) [&sect;](https://postgr.es/c/00a40e33c) [&sect;](https://postgr.es/c/34eb37f79)

   When a partitioned table has an index that is not associated with a constraint, but a partition has an equivalent index that is, then detaching the partition would misbehave, leaving the ex-partition's constraint with an incorrect `coninhcount` value. This would cause trouble during any further manipulations of that constraint.
-  Fix partition pruning setup during `ALTER TABLE DETACH PARTITION CONCURRENTLY` (&Aacute;lvaro Herrera) [&sect;](https://postgr.es/c/96105ebfe) [&sect;](https://postgr.es/c/bf78abebf)

   The executor assumed that no partition could be detached between planning and execution of a query on a partitioned table. This is no longer true since the introduction of `DETACH PARTITION`'s `CONCURRENTLY` option, making it possible for query execution to fail transiently when that is used.
-  Correctly update a partitioned table's `pg_class`.`reltuples` field to zero after its last child partition is dropped (Noah Misch) [&sect;](https://postgr.es/c/e81deeefc)

   The first `ANALYZE` on such a partitioned table must update `relhassubclass` as well, and that caused the `reltuples` update to be lost.
-  Fix handling of polymorphic output arguments for procedures (Tom Lane) [&sect;](https://postgr.es/c/8e0e99972) [&sect;](https://postgr.es/c/bb331af4a)

   The SQL `CALL` statement did not resolve the correct data types for such arguments, leading to errors such as “cannot display a value of type anyelement”, or even outright crashes. (But `CALL` in PL/pgSQL worked correctly.)
-  Fix behavior of stable functions called from a `CALL` statement's argument list (Tom Lane) [&sect;](https://postgr.es/c/0d18b8eb4)

   If the `CALL` is within an atomic context (e.g. there's an outer transaction block), such functions were passed the wrong snapshot, causing them to see stale values of rows modified since the start of the outer transaction.
-  Fix input of ISO-8601 “extended” time format for types `time` and `timetz` (Tom Lane) [&sect;](https://postgr.es/c/019ea7675)

   Re-allow cases such as `T12:34:56`.
-  Detect integer overflow in `money` calculations (Joseph Koshakow) [&sect;](https://postgr.es/c/34e9dce69)

   None of the arithmetic functions for the `money` type checked for overflow before, so they would silently give wrong answers for overflowing cases.
-  Fix over-aggressive clamping of the scale argument in `round(numeric)` and `trunc(numeric)` (Dean Rasheed) [&sect;](https://postgr.es/c/f7aec8c1d)

   These functions clamped their scale argument to +/-2000, but there are valid use-cases for it to be larger; the functions returned incorrect results in such cases. Instead clamp to the actual allowed range of type `numeric`.
-  Fix result for `pg_size_pretty()` when applied to the smallest possible `bigint` value (Joseph Koshakow) [&sect;](https://postgr.es/c/6f6b0f193)
-  Prevent `pg_sequence_last_value()` from failing on unlogged sequences on standby servers and on temporary sequences of other sessions (Nathan Bossart) [&sect;](https://postgr.es/c/c1664c8ee)

   Make it return NULL in these cases instead of throwing an error.
-  Fix parsing of ignored operators in `websearch_to_tsquery()` (Tom Lane) [&sect;](https://postgr.es/c/086ecd12b)

   Per the manual, punctuation in the input of `websearch_to_tsquery()` is ignored except for the special cases of dashes and quotes. However, parentheses and a few other characters appearing immediately before an `or` could cause `or` to be treated as a data word, rather than as an `OR` operator as expected.
-  Detect another integer overflow case while computing new array dimensions (Joseph Koshakow) [&sect;](https://postgr.es/c/a57d16865)

   Reject applying array dimensions `[-2147483648:2147483647]` to an empty array. This is closely related to CVE-2023-5869, but appears harmless since the array still ends up empty.
-  Fix unportable usage of `strnxfrm()` (Jeff Davis) [&sect;](https://postgr.es/c/403cbd210)

   Some code paths for non-deterministic collations could fail with errors like “pg_strnxfrm() returned unexpected result”.
-  Detect another case of a new catalog cache entry becoming stale while detoasting its fields (Noah Misch) [&sect;](https://postgr.es/c/e4afd7153)

   An in-place update occurring while we expand out-of-line fields in a catalog tuple could be missed, leading to a catalog cache entry that lacks the in-place change but is not known to be stale. This is only possible in the `pg_database` catalog, so the effects are narrow, but misbehavior is possible.
-  Correctly check updatability of view columns targeted by `INSERT` ... `DEFAULT` (Tom Lane) [&sect;](https://postgr.es/c/fd958bbbd)

   If such a column is non-updatable, we should give an error reporting that. But the check was missed and then later code would report an unhelpful error such as “attribute number *N* not found in view targetlist”.
-  Avoid reporting an unhelpful internal error for incorrect recursive queries (Tom Lane) [&sect;](https://postgr.es/c/8fc487614)

   Rearrange the order of error checks so that we throw an on-point error when a `WITH RECURSIVE` query does not have a self-reference within the second arm of the `UNION`, but does have one self-reference in some other place such as `ORDER BY`.
-  Lock owned sequences during `ALTER TABLE SET LOGGED|UNLOGGED` (Noah Misch) [&sect;](https://postgr.es/c/112d05570)

   These commands change the persistence of a table's owned sequences along with the table, but they failed to acquire lock on the sequences while doing so. This could result in losing the effects of concurrent `nextval()` calls.
-  Don't throw an error if a queued `AFTER` trigger no longer exists (Tom Lane) [&sect;](https://postgr.es/c/4f1966676)

   It's possible for a transaction to execute an operation that queues a deferred `AFTER` trigger for later execution, and then to drop the trigger before that happens. Formerly this led to weird errors such as “could not find trigger *NNNN*”. It seems better to silently do nothing if the trigger no longer exists at the time when it would have been executed.
-  Fix failure to remove `pg_init_privs` entries for column-level privileges when their table is dropped (Tom Lane) [&sect;](https://postgr.es/c/9cf4beb9e)

   If an extension grants some column-level privileges on a table it creates, relevant catalog entries would remain behind after the extension is dropped. This was harmless until/unless the table's OID was re-used for another relation, when it could interfere with what pg_dump dumps for that relation.
-  Fix selection of an arbiter index for `ON CONFLICT` when the desired index has expressions or predicates (Tom Lane) [&sect;](https://postgr.es/c/b188e1bf7)

   If a query using `ON CONFLICT` accesses the target table through an updatable view, it could fail with “there is no unique or exclusion constraint matching the ON CONFLICT specification”, even though a matching index does exist.
-  Refuse to modify a temporary table of another session with `ALTER TABLE` (Tom Lane) [&sect;](https://postgr.es/c/8397f161e)

   Permissions checks normally would prevent this case from arising, but it is possible to reach it by altering a parent table whose child is another session's temporary table. Throw an error if we discover that such a child table belongs to another session.
-  Fix handling of extended statistics on expressions in `CREATE TABLE LIKE STATISTICS` (Tom Lane) [&sect;](https://postgr.es/c/2aa90c02d)

   The `CREATE` command failed to adjust column references in statistics expressions to the possibly-different column numbering of the new table. This resulted in invalid statistics objects that would cause problems later. A typical scenario where renumbering columns is needed is when the source table contains some dropped columns.
-  Fix failure to recalculate sub-queries generated from `MIN()` or `MAX()` aggregates (Tom Lane) [&sect;](https://postgr.es/c/ce0d16544)

   In some cases the aggregate result computed at one row of the outer query could be re-used for later rows when it should not be. This has only been seen to happen when the outer query uses `DISTINCT` that is implemented with hash aggregation, but other cases may exist.
-  Re-forbid underscore in positional parameters (Erik Wienhold) [&sect;](https://postgr.es/c/315661eca)

   As of v16 we allow integer literals to contain underscores. This change caused input such as `$1_234` to be taken as a single token, but it did not work correctly. It seems better to revert to the original definition in which a parameter symbol is only `$` followed by digits.
-  Avoid crashing when a JIT-inlined backend function throws an error (Tom Lane) [&sect;](https://postgr.es/c/07d66d3cc)

   The error state can include pointers into the dynamically loaded module holding the JIT-compiled code (for error location strings). In some code paths the module could get unloaded before the error report is processed, leading to SIGSEGV when the location strings are accessed.
-  Cope with behavioral changes in libxml2 version 2.13.x (Erik Wienhold, Tom Lane) [&sect;](https://postgr.es/c/f85c91a18)

   Notably, we now suppress “chunk is not well balanced” errors from libxml2, unless that is the only reported error. This is to make error reports consistent between 2.13.x and earlier libxml2 versions. In earlier versions, that message was almost always redundant or outright incorrect, so 2.13.x substantially reduced the number of cases in which it's reported.
-  Fix handling of subtransactions of prepared transactions when starting a hot standby server (Heikki Linnakangas) [&sect;](https://postgr.es/c/b5b418b68)

   When starting a standby's replay at a shutdown checkpoint WAL record, transactions that had been prepared but not yet committed on the primary are correctly understood as being still in progress. But subtransactions of a prepared transaction (created by savepoints or PL/pgSQL exception blocks) were not accounted for and would be treated as aborted. That led to inconsistency if the prepared transaction was later committed.
-  Prevent incorrect initialization of logical replication slots (Masahiko Sawada) [&sect;](https://postgr.es/c/2f3304ce1)

   In some cases a replication slot's start point within the WAL stream could be set to a point within a transaction, leading to assertion failures or incorrect decoding results.
-  Avoid “can only drop stats once” error during replication slot creation and drop (Floris Van Nee) [&sect;](https://postgr.es/c/f2c922ff2)
-  Fix resource leakage in logical replication WAL sender (Hou Zhijie) [&sect;](https://postgr.es/c/b8f953d8d)

   The walsender process leaked memory when publishing changes to a partitioned table whose partitions have row types physically different from the partitioned table's.
-  Avoid memory leakage after servicing a notify or sinval interrupt (Tom Lane) [&sect;](https://postgr.es/c/54a7b21b3)

   The processing functions for these events could switch the current memory context to TopMemoryContext, resulting in session-lifespan leakage of any data allocated before the incorrect setting gets replaced. There were observable leaks associated with (at least) encoding conversion of incoming queries and parameters attached to Bind messages.
-  Prevent leakage of reference counts for the shared memory block used for statistics (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/6f61d0e7e)

   A new backend process attaching to the statistics shared memory incremented its reference count, but failed to decrement the count when exiting. After 2<sup>32</sup> sessions had been created, the reference count would overflow to zero, causing failures in all subsequent backend process starts.
-  Prevent deadlocks and assertion failures during truncation of the multixact SLRU log (Heikki Linnakangas) [&sect;](https://postgr.es/c/e7cbe5a85)

   A process trying to delete SLRU segments could deadlock with the checkpointer process.
-  Avoid possibly missing end-of-input events on Windows sockets (Thomas Munro) [&sect;](https://postgr.es/c/a622095bc)

   Windows reports an FD_CLOSE event only once after the remote end of the connection disconnects. With unlucky timing, we could miss that report and wait indefinitely, or at least until a timeout elapsed, expecting more input.
-  Fix buffer overread in JSON parse error reports for incomplete byte sequences (Jacob Champion) [&sect;](https://postgr.es/c/5396a2987)

   It was possible to walk off the end of the input buffer by a few bytes when the last bytes comprise an incomplete multi-byte character. While usually harmless, in principle this could cause a crash.
-  Disable creation of stateful TLS session tickets by OpenSSL (Daniel Gustafsson) [&sect;](https://postgr.es/c/cc606afce) [&sect;](https://postgr.es/c/83b4a6358) [&sect;](https://postgr.es/c/441eba34d)

   This avoids possible failures with clients that think receipt of a session ticket means that TLS session resumption is supported.
-  When replanning a PL/pgSQL “simple expression”, check it's still simple (Tom Lane) [&sect;](https://postgr.es/c/82a931d3d)

   Certain fairly-artificial cases, such as dropping a referenced function and recreating it as an aggregate, could lead to surprising failures such as “unexpected plan node type”.
-  Fix PL/pgSQL's handling of integer ranges containing underscores (Erik Wienhold) [&sect;](https://postgr.es/c/b4e909082)

   As of v16 we allow integer literals to contain underscores, but PL/pgSQL failed to handle examples such as `FOR i IN 1_001..1_003`.
-  Fix recursive `RECORD`-returning PL/Python functions (Tom Lane) [&sect;](https://postgr.es/c/52ea653aa)

   If we recurse to a new call of the same function that passes a different column definition list (`AS` clause), it would fail because the inner call would overwrite the outer call's idea of what rowtype to return.
-  Don't corrupt PL/Python's `TD` dictionary during a recursive trigger call (Tom Lane) [&sect;](https://postgr.es/c/be18a12b6)

   If a PL/Python-language trigger caused another one to be invoked, the `TD` dictionary created for the inner one would overwrite the outer one's `TD` dictionary.
-  Fix PL/Tcl's reporting of invalid list syntax in the result of a function returning tuple (Erik Wienhold, Tom Lane) [&sect;](https://postgr.es/c/c236ecc82)

   Such a case could result in a crash, or in emission of misleading context information that actually refers to the previous Tcl error.
-  Avoid non-thread-safe usage of `strerror()` in libpq (Peter Eisentraut) [&sect;](https://postgr.es/c/c53016860)

   Certain error messages returned by OpenSSL could become garbled in multi-threaded applications.
-  Avoid memory leak within pg_dump during a binary upgrade (Daniel Gustafsson) [&sect;](https://postgr.es/c/0ae05c18e)
-  Ensure that `pg_restore` `-l` reports dependent TOC entries correctly (Tom Lane) [&sect;](https://postgr.es/c/5dce8ce0a)

   If `-l` was specified together with selective-restore options such as `-n` or `-N`, dependent TOC entries such as comments would be omitted from the listing, even when an actual restore would have selected them.
-  Allow `contrib/pg_stat_statements` to distinguish among utility statements appearing within SQL-language functions (Anthonin Bonnefoy) [&sect;](https://postgr.es/c/9cd365f28)

   The SQL-language function executor failed to pass along the query ID that is computed for a utility (non `SELECT`/`INSERT`/`UPDATE`/`DELETE`/`MERGE`) statement.
-  Avoid “cursor can only scan forward” error in `contrib/postgres_fdw` (Etsuro Fujita) [&sect;](https://postgr.es/c/d97f2ee50)

   This error could occur if the remote server is v15 or later and a foreign table is mapped to a non-trivial remote view.
-  In `contrib/postgres_fdw`, do not send `FETCH FIRST WITH TIES` clauses to the remote server (Japin Li) [&sect;](https://postgr.es/c/8405d5a37)

   The remote server might not implement this clause, or might interpret it differently than we would locally, so don't risk attempting remote execution.
-  Avoid clashing with system-provided `<regex.h>` headers (Thomas Munro) [&sect;](https://postgr.es/c/31423bc44)

   This fixes a compilation failure on macOS version 15 and up.
-  Fix otherwise-harmless assertion failure in Memoize cost estimation (David Rowley) [&sect;](https://postgr.es/c/6143c9c03)
-  Fix otherwise-harmless assertion failures in `REINDEX CONCURRENTLY` applied to an SP-GiST index (Tom Lane) [&sect;](https://postgr.es/c/06f81fed3)
