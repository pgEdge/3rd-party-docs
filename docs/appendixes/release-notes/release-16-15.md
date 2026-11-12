<a id="release-16-15"></a>

## Release 16.15


**Release date:.**


2026-08-13


 This release contains a variety of fixes from 16.14. For information about new features in major release 16, see [Release 16](release-16.md#release-16).
 <a id="release-16-15-migration"></a>

### Migration to Version 16.15


 A dump/restore is not required for those running 16.X.


 However, the first three security entries below describe configuration adjustments and data cleanups that you may need to make after updating.


 Also, if you use `contrib/btree_gist` or `contrib/ltree`, you may need to reindex indexes made with those extensions; see the relevant entries below.


 Also, if you are upgrading from a version earlier than 16.10, see [Release 16.10](release-16-10.md#release-16-10).
  <a id="release-16-15-changes"></a>

### Changes


-  Restrict logical decoding output plugins to the set specified by a new server parameter `output_plugin_libraries` (Jacob Champion) [&sect;](https://postgr.es/c/4e1252ee5)

   Previously, a replication user could select any loadable library for logical decoding, allowing exploits of various sorts. To allow locking this down without breaking setups that worked before, introduce a whitelist of allowed output plugins.

   By default, only the output plugins shipped as part of PostgreSQL (`pgoutput` and `test_decoding`) are included in `output_plugin_libraries`. Installations that rely on other output plugins must add them after updating the server, for example

```

output_plugin_libraries = 'pgoutput, test_decoding, my_trusted_decoder'
```


   Additionally, `pg_upgrade --check` will fail if the `output_plugin_libraries` parameter on the new cluster does not permit the plugins of logical replication slots on the old cluster, when migrating from versions 17 and later. Make necessary additions to the new cluster's setting before performing pg_upgrade.

   The PostgreSQL Project thanks Vladimir Tokarev and Yu Kunpeng for reporting this problem. (CVE-2026-6471)
-  Fix `contrib/pgcrypto`'s PGP encryption to detect unsupported ciphers (Daniel Gustafsson) [&sect;](https://postgr.es/c/6ca6023ca) [&sect;](https://postgr.es/c/ba2a63faf)

   Previously, if OpenSSL rejected the requested cipher (for example, because it is running in FIPS mode, or the legacy provider hasn't been loaded), pgcrypto failed to notice the failure and simply XOR'd the non-encrypted block with the plaintext, rendering the “encryption” trivially breakable. This will typically occur with deprecated or non-FIPS cipher algorithms (cipher-algo=blowfish/bf, twofish, cast5, or 3des).

   By default, pgcrypto will now fail to decrypt any messages that were affected in this way. To allow retrieval of such data, a new option `ignore-cipher-failure` has been added to `pgp_pub_decrypt()` and `pgp_sym_decrypt()`. Setting `ignore-cipher-failure=1` will restore their previous behavior, allowing the faulty encryption wrapper to be stripped off:

```

pgp_sym_decrypt(encrypted_column, ANY KEY, 'ignore-cipher-failure=1')
```
   Once the affected messages are identified and stripped of their wrappers, they can then be re-encrypted with a modern algorithm. It is important however that the behavior of OpenSSL be the same as it was when the faulty messages were created: if the set of unsupported algorithms is not the same, this approach will not work. See the documentation for `ignore-cipher-failure`.

   The PostgreSQL Project thanks Shishir Sharma for reporting this problem. (CVE-2026-14663)
-  Fix psql to skip in-line data following a scripted `COPY ... FROM STDIN` command, even if the `COPY` fails before sending `PGRES_COPY_IN` (Tom Lane) [&sect;](https://postgr.es/c/2bdfd5cdc) [&sect;](https://postgr.es/c/db96e87f9)

   Previously, if a `COPY` command failed at startup (for instance, because the target table doesn't exist) psql would not realize that and would proceed to read the following in-line data as SQL commands. In the best case that's wrong and in the worst case it's a SQL-injection hazard. Teach psql to recognize syntactically-valid `COPY ... FROM STDIN` commands and to skip data on its own authority if the server doesn't respond with `PGRES_COPY_IN`.

   While this fix is unlikely to affect any production SQL scripts, test scripts might intentionally exercise failing `COPY ... FROM STDIN` commands. Those will need to gain a `\.` data terminator line after each such command.

   The PostgreSQL Project thanks Alexander Lakhin for reporting this problem. (CVE-2026-6464)
-  Cross-check the output row type of a portal running `EXECUTE` or `FETCH` (Robert Haas) [&sect;](https://postgr.es/c/7d150b5c9)

   `EXECUTE` and `FETCH` use two portals: an outer one for the statement itself, and an inner one running the query being executed on its behalf. It was previously possible to make the declared row types of the two portals diverge, leading to server memory disclosure and arbitrary code execution.

   The PostgreSQL Project thanks Ben Morris (in collaboration with Claude and Anthropic Research) and Peter Geoghegan for reporting this problem. (CVE-2026-16239)
-  Fix buffer overrun with long time zone abbreviation in `to_char()` (Tom Lane) [&sect;](https://postgr.es/c/5dbeb69bc)

   This can easily crash the server, and exploits leading to arbitrary code execution have been reported.

   The PostgreSQL Project thanks Hcamael, Amjad Shahzad, Tan Zhen of AntAISecurityLab, Tomer Fichman, Zheng Yu, Amy Burnett (OpenAI Codex Security), Rick de Jager, Heewon Song, Sylvie Mayer, Aleksander Alekseev, and Hillai Ben Sasson for reporting this problem. (CVE-2026-14669)
-  Fix buffer overrun in regexp match/split functions (Masahiko Sawada) [&sect;](https://postgr.es/c/3179253c2)

   If passed invalidly-encoded data, these functions could write past the end of their conversion buffer.

   The PostgreSQL Project thanks Francesco Verardi for reporting this problem. (CVE-2026-14664)
-  Harden the `ascii()` function against invalid input (Michael Paquier) [&sect;](https://postgr.es/c/ac450853d)

   By supplying invalidly-encoded input, this function could be coaxed to read and return a few bytes of data that it shouldn't. In assert-enabled builds, its assertions could be triggered too.

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-18024)
-  Make `scalarineqsel()` check that a constant it expects to be of type `tid` actually is (Tom Lane) [&sect;](https://postgr.es/c/591192e2e)

   This expectation will hold for all the built-in operators that use this estimator, but a maliciously-constructed operator could violate it, leading to a crash or server memory disclosure.

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-14668)
-  Harden `tsvector` and `tsquery` code against overly long values (both individual lexemes and total vector/query length) (Tom Lane) [&sect;](https://postgr.es/c/7c1a8805a) [&sect;](https://postgr.es/c/84b6a7a06)

   The documented limits were not enforced in all code paths.

   The PostgreSQL Project thanks Yuhang Wu, Zhenpeng Lin, Zheng Yu, and Hcamael for reporting these problems. (CVE-2026-14662)
-  Fix various places that mistakenly assumed they would not have to deal with more than `FUNC_MAX_ARGS` function arguments (Tom Lane) [&sect;](https://postgr.es/c/8c67c4cf5) [&sect;](https://postgr.es/c/32e0d25d4)

   Notably, the server's actual limit on the number of arguments to an aggregate function is `FUNC_MAX_ARGS - 1`, but the parser failed to enforce that, creating hazards downstream.

   The PostgreSQL Project thanks Zheng Yu, ylwangtju, and Masahiko Sawada for reporting these problems. (CVE-2026-14679)
-  Reject calls from SQL to functions that take or return type `internal` (Tom Lane) [&sect;](https://postgr.es/c/4c0c1b2cd) [&sect;](https://postgr.es/c/7f308dd7c)

   The existing defenses against doing this have been shown to be insufficient, so add more explicit checks.

   The PostgreSQL Project thanks Amy Burnett (OpenAI Codex Security) for reporting this problem. (CVE-2026-14680)
-  Preserve the ownership of extended statistics objects when they are rebuilt by `ALTER TABLE` (Masahiko Sawada) [&sect;](https://postgr.es/c/7c3367ab5)

   Previously, the role running `ALTER TABLE` gained ownership of such objects, but that seems inappropriate.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-6469)
-  When deparsing an `EXTRACT()` function call, quote the field name if needed (Nathan Bossart) [&sect;](https://postgr.es/c/5e5ea74e3)

   The parser accepts any string literal as a field name in `EXTRACT()`, deferring validation to execution. If the call is stored and deparsed (for example during pg_dump), the string body was regurgitated verbatim, allowing SQL injection.

   The PostgreSQL Project thanks Ben Morris (in collaboration with Claude and Anthropic Research) for reporting this problem. (CVE-2026-15741)
-  Check for `USAGE` privilege on data types in places that formerly failed to check that (Nathan Bossart) [&sect;](https://postgr.es/c/24855357c) [&sect;](https://postgr.es/c/7a761fe40) [&sect;](https://postgr.es/c/7ce056b17)

   `CREATE TYPE AS RANGE` did not check, nor did `ALTER TABLE OF`, nor did commands that create stored expressions. These omissions allowed roles without `USAGE` privilege to nonetheless create objects depending on the type, possibly blocking the type's owner from changing the type later.

   The PostgreSQL Project thanks Jingzhou Fu for reporting this problem. (CVE-2026-6470)
-  Invalidate role-dependent cached plans after role changes (Ilya Staroverov, Shinya Kato, Nathan Bossart) [&sect;](https://postgr.es/c/f5f2da795)

   Role membership, role attribute, and database ownership changes may impact the expected behavior of row-level security policies, but previously we'd continue to use cached plans that were made according to the old state of affairs.

   The PostgreSQL Project thanks Ilya Staroverov and Shinya Kato for reporting this problem. (CVE-2026-14666)
-  Make mock SCRAM authentication secrets more plausible (Nathan Bossart) [&sect;](https://postgr.es/c/fadbe882d)

   If a SCRAM login is attempted against a role that doesn't exist or doesn't have a SCRAM secret, we generate a mock secret and carry out the authentication handshake anyway, to avoid revealing these facts to an attacker. But the mock secret was made with a fixed iteration count, which in itself can be an observable response discrepancy. Use the configuration setting `scram_iterations` instead, to make the mock secret look more like the installation's real secrets.

   The PostgreSQL Project thanks Radim Marek for reporting this problem. (CVE-2026-14672)
-  Fix out-of-bounds writes in ecpg applications caused by invalid `bytea` data received from the server (Michael Paquier) [&sect;](https://postgr.es/c/39f855e0e)

   ecpg assumed without checking that any `bytea` value must begin with `\x`. A broken or malicious server might send a string shorter than 2 bytes, resulting in memory clobber in the application.

   The PostgreSQL Project thanks ylwangtju for reporting this problem. (CVE-2026-16241)
-  Do not do backquote expansion on the argument of psql's `\unrestrict` command (Nathan Bossart) [&sect;](https://postgr.es/c/33d0c63fb)

   This oversight in the fix for CVE-2025-8714 allows a malicious server to inject shell commands into plain-text dump output that will be run at restore time on the machine running psql, the exact scenario that CVE-2025-8714 intended to prevent.

   The PostgreSQL Project thanks Lucas Velgus, Filip Janus, and Daniel Bakker for reporting this problem. (CVE-2026-18408)
-  Remove pg_dump's assumption that `pg_proc`.`protrftypes` cannot have more than `FUNC_MAX_ARGS` entries (Tom Lane) [&sect;](https://postgr.es/c/3299c2ba0)

   Since there could be entries for both input and output arguments, it's feasible for this array's length to exceed `FUNC_MAX_ARGS` (which constrains only input arguments). Even if that were not so, pg_dump cannot assume that the server was built with the same value of `FUNC_MAX_ARGS` that it has. An overrun would lead to a memory clobber inside pg_dump.

   The PostgreSQL Project thanks Masahiko Sawada for reporting this problem. (CVE-2026-19385)
-  Harden PL/Perl against “tied” Perl arrays and hashes (Tom Lane) [&sect;](https://postgr.es/c/477a6bdb0)

   A tied object that doesn't behave like a regular one could lead to memory overwrite, or to constructing a corrupt result array (which would likely cause problems later).

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-14670)
-  Fix integer overflows in memory-allocation calculations in PL/Perl and PL/Tcl (Heikki Linnakangas) [&sect;](https://postgr.es/c/8bbdc0540)

   This is the same type of problem as CVE-2026-6473, just in a different part of the code, and is fixed in the same way.

   The PostgreSQL Project thanks the Tulya Project (Team Dhiutsa, Bitecope Technologies Private Ltd) for reporting this problem. (CVE-2026-14677)
-  Ensure that `contrib/amcheck` functions restrict `search_path` before executing index expressions (Noah Misch) [&sect;](https://postgr.es/c/2cc147318)

   Because amcheck will run such index expressions as the owner of their tables, a caller could potentially hijack `search_path`-dependent functions to run arbitrary code as the table owner. By default this is not a vulnerability because only superusers are allowed to call amcheck functions; but if that privilege was granted out, it created a larger hazard than the documentation suggests.

   The PostgreSQL Project thanks Yuelin Wang and Jacob Brazeal for reporting this problem. (CVE-2026-14673)
-  Fix integer overflows in `contrib/fuzzystrmatch`'s `levenshtein()` and `levenshtein_less_equal()` functions (Nathan Bossart) [&sect;](https://postgr.es/c/61481aa0e)

   Passing large cost values to these functions could cause integer overflows, thereby producing nonsensical results, and even causing out-of-bounds writes in some cases.

   The PostgreSQL Project thanks Ben Morris (in collaboration with Claude and Anthropic Research) for reporting this problem. (CVE-2026-15742)
-  Fix datatype error in `contrib/pg_trgm`'s GiST picksplit function (Heikki Linnakangas) [&sect;](https://postgr.es/c/24c88cd39)

   This mistake resulted in reading past the end of the buffer, typically causing bad split decisions; but a crash could ensue if you're very unlucky.

   The PostgreSQL Project thanks Mehmet D. Ince for reporting this problem. (CVE-2026-14678)
-  Remove the plan cache in `contrib/refint` (Ayush Tiwari) [&sect;](https://postgr.es/c/a572dd213)

   This caching behavior has several serious bugs, notably that `check_foreign_key()` embeds the new key values in its cascade-UPDATE queries, so a cached plan reuses the originally-needed values rather than the key values that should be used. The simplest solution is to remove it.

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-14671)
-  Fix self-deadlock when replaying WAL generated by an older minor version (Andrey Borodin) [&sect;](https://postgr.es/c/42a3194e5)

   This error was introduced in the previous set of minor releases. It caused standby servers that were following a primary of an older minor release version to get stuck in some scenarios.
-  Fix mis-handling of asynchronous reads when rescanning an asynchronous Append plan node (Alexander Korotkov, Gleb Kashkin, Etsuro Fujita) [&sect;](https://postgr.es/c/894da35b3)

   When an upper plan node rescans an Append before having read the entire Append output, we need to discard any in-flight requests sent to external servers (by `postgres_fdw` for example). This was not done correctly in cases where a subplan has parameter changes or is discarded by partition pruning in the next scan. The outcome could be incorrect query results, an infinite loop, or an assertion failure.
-  Fix error in partition pruning for RANGE-partitioned tables (David Rowley) [&sect;](https://postgr.es/c/9a0cd8e73)

   In some cases the DEFAULT partition would be skipped when it should not be, which could lead to rows missing from query results.
-  Fix planner's nullability and strictness checks for <em>value</em><code> IN
      (</code><em>array</em><code>)</code> expressions (Ayush Tiwari) [&sect;](https://postgr.es/c/6fcee189b)

   These checks should only succeed if the array operand is known to be non-empty, but that consideration was missed, allowing optimizations to be applied that should not be. This could result in wrong query answers if the array actually was empty.
-  Fix incorrect join removal logic (Matheus Alcantara, Richard Guo) [&sect;](https://postgr.es/c/3e1fe25e6) [&sect;](https://postgr.es/c/d610d8e8b)

   In edge cases, it was possible for a constant output value coming from within the nullable side of an outer join to not be replaced by NULL when it should be.
-  Add missed checks for hashability of equality comparisons on container datatypes (arrays, composites, ranges) (Andrei Lepikhov, Tom Lane) [&sect;](https://postgr.es/c/cf2bfe073)

   The planner must verify hashability of the container's component type(s) before deciding it can use a hash-based plan type. This step was missed in some places, leading to “could not identify a hash function” failures at execution.
-  Fix mis-optimization of `COUNT` window functions that have an `EXCLUDE` clause or lack `ORDER BY` (Chengpeng Yan, David Rowley) [&sect;](https://postgr.es/c/a85732162)

   These window functions were treated as monotonic when they should not be, allowing wrong answers to be computed.
-  Fix `ALTER COLUMN ... DROP EXPRESSION` to work when there are multiple levels of partitions (Alberto Piai) [&sect;](https://postgr.es/c/34785a0d5)
-  Disallow renaming a rule to `_RETURN` (Tom Lane) [&sect;](https://postgr.es/c/a69503fb1)

   That name is reserved for a view's `ON SELECT` rule, but `ALTER RULE` allowed renaming other rules to `_RETURN`, causing trouble later.
-  Fix missing lock release for role membership grants in `DROP OWNED BY` (Jeff Davis) [&sect;](https://postgr.es/c/e61d44fde)

   This oversight resulted in a warning message, followed by retaining a lock on the membership grant until the end of the transaction.
-  Fix failure of `EXPLAIN` when deparsing `SQL/JSON` aggregates (Richard Guo) [&sect;](https://postgr.es/c/485527190)

   Some plan structures resulted in “invalid JsonConstructorExpr underlying node type” errors.
-  Fix use of `REINDEX CONCURRENTLY` with a deferred uniqueness constraint (Nitin Motiani) [&sect;](https://postgr.es/c/ac222bea5)

   The transient index copy created during `REINDEX CONCURRENTLY` was incorrectly marked as enforcing immediate uniqueness, causing spurious reports of constraint violation.
-  Fix matching of localized month/day names in `to_date()` (Heikki Linnakangas) [&sect;](https://postgr.es/c/b594efe52)

   The matching logic misbehaved in cases where case-folding changes the byte length of the string.
-  Fix incorrect NFC recomposition for Hangul U+11A7 (TBASE) (Diego Frias, Michael Paquier) [&sect;](https://postgr.es/c/82116023e)

   This character was treated as a valid T syllable, which it is not, and hence silently swallowed during normalization.
-  Avoid possible truncation of output lexemes in case-insensitive `synonym` dictionaries (Jeff Davis) [&sect;](https://postgr.es/c/b32df590c)

   If folding to lower case increased the byte length of a lexeme, it was incorrectly truncated to its original byte length when emitted.
-  Fix typo in `hash_record_extended()` (Man Zeng) [&sect;](https://postgr.es/c/8b5580a05)

   The code failed to initialize the second isnull argument passed to FunctionCallInvoke(). This is harmless for existing in-core extended hash support functions, which will not examine that value. However, extension-provided hash functions could be affected if they inspect `PG_ARGISNULL(1)`.
-  Fix `pg_get_publication_tables()` to not fail if a publishable table is dropped concurrently (Bharath Rupireddy) [&sect;](https://postgr.es/c/6c760f6b6)
-  Prevent `satisfies_hash_partition()` from crashing with `VARIADIC NULL` (Robert Haas) [&sect;](https://postgr.es/c/6de480156)
-  Report invalid-weight errors more cleanly and consistently in `tsvector_filter()` and allied functions (Ewan Young) [&sect;](https://postgr.es/c/ed5ca5828)

   In particular, report weight characters that are not printable ASCII in octal form (<code>\</code><em>nnn</em>), as `charout()` would render them. This avoids possibly producing an invalidly-encoded error message.
-  Fix mishandling of namespace nodes in `xpath()` (Michael Paquier) [&sect;](https://postgr.es/c/174076601) [&sect;](https://postgr.es/c/fac26fd41)

   This fix avoids an unexpected “could not copy node” error.
-  Treat an undefined jsonpath variable as an error even when no variables are supplied (Andrey Rachitskiy) [&sect;](https://postgr.es/c/1daeef6e0)

   The `jsonb` `@?` and `@@` operators cannot supply any variables to be used in their jsonpath expressions. This code path erroneously treated an unknown jsonpath variable as a JSON null, rather than raising an error as expected. Aside from not being the expected behavior, this mistake could result in unbounded memory consumption.
-  Fix NULL-pointer crash when `IS JSON` or similar constructs have an argument that is of string category but lacks a cast to type `text` (Ayush Tiwari) [&sect;](https://postgr.es/c/60abb3c73)

   There are no such data types in core PostgreSQL, but the problem is reachable with some extension types.
-  Avoid machine-dependent behavior when dividing the smallest possible `money` value by -1 (Andrey Rachitskiy) [&sect;](https://postgr.es/c/86f42357c)
-  Fix crash after out-of-memory failure partway through creation of a cache entry for a text search dictionary (Tom Lane) [&sect;](https://postgr.es/c/83336e3ed)
-  Fix memory-safety bugs in processing of incorrect ispell/hunspell dictionary files (Andrey Rachitskiy) [&sect;](https://postgr.es/c/444038bb7)
-  Prevent access to other sessions' temporary tables (Jim Jones, Daniil Davydov, Alexander Korotkov) [&sect;](https://postgr.es/c/cc3fe7e2a)

   Some code paths failed to prevent this, leading to silently wrong (inconsistent) results.
-  Honor query cancel and vacuum delay during GIN index posting-tree cleanup (Paul Kim, Alexander Korotkov) [&sect;](https://postgr.es/c/006ac761e)

   The posting tree for a common value can be large, so that this missed check could allow vacuum to run for a long time before noticing an interrupt.
-  Fix possible mis-decoding of index tuples during GiST and SP-GiST index-only scans (Peter Geoghegan) [&sect;](https://postgr.es/c/e58192546)

   This error could lead to emitting corrupted data from an index-only scan plan. The only affected core opclass is GiST's range_ops, and it could only fail if the range column were not the first index column.
-  Ensure that the new last block of a bulk-extended table is added to its free space map promptly (Jingtang Zhang) [&sect;](https://postgr.es/c/0fd5595aa)

   An off-by-one error caused the last block of a multi-block table extension to not be marked as free in the map. This would eventually get corrected by vacuum, but meanwhile the space wouldn't be used.
-  When creating directories, tolerate concurrent creation of the same directory (Andrew Dunstan, Tom Lane) [&sect;](https://postgr.es/c/9951e3d38)
-  Prevent creation of dangling object dependencies by acquiring a shared lock on any object being depended on (Bertrand Drouvot) [&sect;](https://postgr.es/c/d9bc0d96c) [&sect;](https://postgr.es/c/d44cd4674)

   The shared lock will conflict with any attempt to drop the depended-on object, eliminating the race condition that formerly existed. For example, if one session drops a schema (that appears empty to it) concurrently with some other session creating a function in that schema, previously both transactions could commit, leaving an invalid function definition behind. Now, one transaction or the other will fail.
-  Fix race condition in conflict detection for `SERIALIZABLE` isolation mode (Peter Geoghegan) [&sect;](https://postgr.es/c/e321faaae)

   A conflict could be missed when examining an initially-empty btree index, allowing failure of serializability due to improperly allowing conflicting transactions to commit.
-  Fix race condition in ProcSignalBarrier code (Masahiko Sawada) [&sect;](https://postgr.es/c/cdb9b2830)

   This error could result in processes getting stuck, typically after reporting “still waiting for backend with PID *nnnn* to accept ProcSignalBarrier”.
-  Fix race conditions when a set of processes that belong to the same lock group exit at the same time (Vlad Lesin) [&sect;](https://postgr.es/c/65d04df60) [&sect;](https://postgr.es/c/4e8edf53f)

   These errors could lead to PANIC aborts, with messages such as “latch already owned”. The issue does not normally arise in regular parallel query, since the leader won't exit before seeing its workers finish; but some extensions reach the problem.
-  Fix race with timeline selection in logical decoding during standby promotion (Bertrand Drouvot) [&sect;](https://postgr.es/c/8cd687c44) [&sect;](https://postgr.es/c/2a04624da) [&sect;](https://postgr.es/c/d9b49e5b4)

   Logical decoding being performed on the standby could fail with a “requested WAL segment has already been removed” error. A repeat attempt would succeed, so there was no permanent problem but there was an availability hazard.
-  Avoid exposing a WAL receiver's full connection string during timeline jumps (Chao Li) [&sect;](https://postgr.es/c/e037a4199)

   The `pg_stat_wal_receiver` view should show a sanitized version of the connection string, without sensitive data. But it transiently showed the full string when we re-use an existing WAL receiver.
-  Use run-time checks, not just Asserts, to verify the correct number of columns in tuples received during logical replication (Varik Matevosyan) [&sect;](https://postgr.es/c/59759e1a5)

   A malicious or buggy publisher could send inconsistent numbers of columns. While we could not find a scenario in which this would have serious ill effects, extra caution seems warranted.
-  Clean up quoting of string parameters within constructed replication commands (Tom Lane) [&sect;](https://postgr.es/c/26d6c19d0)

   Various places that generate replication commands were not being adequately careful about quoting replication slot names and other parameters that need to be inserted into those commands. This could result in unexpected syntax errors in those commands. In principle, a crafted replication slot name could result in SQL injection; but such a scenario seems very unlikely to occur in practice, since replication operations can only be invoked by highly-privileged users and there is no reason for them to use a slot name coming from an untrustworthy source.
-  Fix logical decoding of empty prepared transactions (Masahiko Sawada) [&sect;](https://postgr.es/c/744618346)

   A prepared transaction that did not cause any decodable updates could result in sending `COMMIT/ROLLBACK PREPARED` to the output plugin with no preceding `PREPARE`. For the built-in subscriber this breaks replication, and other plugins will probably not like it either.
-  Fix corruption of unlogged sequences after standby promotion (Fujii Masao) [&sect;](https://postgr.es/c/913d3b610)

   Previously, if an unlogged sequence was created on the primary and replicated to a standby, accessing the sequence after promoting the standby could fail with “bad magic number in sequence” or related errors.
-  Fix cascading standby reconnect failure after archive fallback (Marco Nenciarini) [&sect;](https://postgr.es/c/5dbeb69bc)

   A cascading standby could fail to reconnect to its upstream standby with “requested starting point ... is ahead of the WAL flush position” after falling back to archive recovery.
-  Avoid race condition while dropping ephemeral replication slots (Zhijie Hou) [&sect;](https://postgr.es/c/51e79222c)

   The slot-releasing code performed some additional updates to the replication slot's shared-memory entry after releasing the slot. This is unsafe since another session could immediately re-use the dropped slot's shared-memory entry. Skip those updates in the case of an ephemeral slot.
-  Fix stale progress reports during logical replication table synchronization (Shinya Kato) [&sect;](https://postgr.es/c/d140237da)

   Previously, the `pg_stat_progress_copy` view in the subscriber would continue to show the initial `COPY` operation as active even after the data copy had finished. The stale entry remained visible until synchronization caught up with the publisher.
-  Clear base backup progress on backup failure (Chao Li) [&sect;](https://postgr.es/c/7cbd80340) [&sect;](https://postgr.es/c/fffa4d870)

   Previously the `pg_stat_progress_basebackup` view would continue to show a stale progress entry after a failure, until the replication client disconnected. pg_basebackup normally disconnects immediately, but other clients might not.
-  Fix possible PANIC due to concurrent drop of pgstats entries when `track_functions` is enabled (Sami Imseih, Michael Paquier) [&sect;](https://postgr.es/c/bf4616b59) [&sect;](https://postgr.es/c/7c457ea15) [&sect;](https://postgr.es/c/e34b1ff5d)
-  Clean up broken local pgstats entry after failing to obtain space for the corresponding shared hashtable entry (Niall Newman) [&sect;](https://postgr.es/c/636bcd6a5)

   Failure to do this led to a null-pointer dereference the next time the local entry was used.
-  In PL/Perl, avoid NULL pointer dereference crash when working with an invalid `PostgreSQL::InServer::ARRAY` object (Xing Guo) [&sect;](https://postgr.es/c/3854f4afc)
-  In PL/Python, properly check for errors when working with sequence and mapping objects (Richard Guo) [&sect;](https://postgr.es/c/e92f23bde)

   Previously, a broken object or an unhandled exception could result in a NULL pointer dereference crash.
-  In libpq, always drain all pending bytes from the SSL or GSS decryption buffer during `pqReadData()` (Jacob Champion) [&sect;](https://postgr.es/c/d3f329656) [&sect;](https://postgr.es/c/477bc27a5) [&sect;](https://postgr.es/c/6243ea6c3) [&sect;](https://postgr.es/c/56988064b)

   This avoids edge cases where libpq or its calling application waits for more data to arrive on the socket, but actually all the data has already arrived.
-  Allow libpq to accept ParameterDescription messages exceeding 30000 bytes (Ning Sun) [&sect;](https://postgr.es/c/8d516d2b9)

   Previously, this message type was not among those that libpq's validity heuristics believed could be long. The limit resulted in failure for prepared queries having more than 7498 parameters, which is unlikely but supported.
-  Reject multiple descriptor header items in ecpg's `GET/SET DESCRIPTOR` statements (Masashi Kamura) [&sect;](https://postgr.es/c/4d4b73cfd)

   Previously the grammar allowed this syntax, but broken C code was generated. Adjust the grammar and the documentation to allow only one header item.
-  Make line widths match in psql's expanded aligned output format (Pavel Stehule) [&sect;](https://postgr.es/c/1edbe6695)

   When the table's data rows are narrower than the record header lines, widen the data rows to match the headers, avoiding unsightly output.
-  Fix psql's privilege check for showing database size in `\l+` (Christoph Berg) [&sect;](https://postgr.es/c/844bc718a)

   The underlying server function permits users who have `pg_read_all_stats` privileges to see the sizes of all databases, even if they lack `CONNECT` privilege. But psql was unaware of that provision and would not call the function unless the user has `CONNECT` privilege.
-  Fix psql's tab completion for `\df` to consider procedures too (Erik Wienhold) [&sect;](https://postgr.es/c/f9afdb6d1)
-  Fix thread-safety bug in pgbench (Fujii Masao) [&sect;](https://postgr.es/c/6432a4cd6)

   When pgbench runs with multiple threads and the `--verbose-errors` option, different threads could attempt to use the same buffer to construct error messages, leading to corrupted log output.
-  Use the source cluster's group-read file permissions for pg_recvlogical output files (Fujii Masao) [&sect;](https://postgr.es/c/9590dcfca)

   pg_recvlogical was documented to behave this way, but it never actually enabled group-read.
-  In `contrib/amcheck`, handle short-header varlena datums correctly (Andrey Borodin) [&sect;](https://postgr.es/c/c89a1ca01)

   This error could result in doing excess work while verifying a btree index, but seems not to have had any worse consequences.
-  In `contrib/btree_gist`, fix `NaN` handling in the `float4` and `float8` opclasses (Bill Kim, Tom Lane) [&sect;](https://postgr.es/c/d569ccd40)

   Comparisons, as well as the GiST penalty and distance functions, did not account for `NaN` and would give the wrong answer when handed one. It is recommended to reindex `btree_gist` indexes on float columns after installing this update, if there is any possibility that there are `NaN` entries in those columns.
-  In `contrib/btree_gist`, fix searches using a not-equal operator (Ayush Tiwari) [&sect;](https://postgr.es/c/0663382c9)

   For variable-length data types, the code for scanning non-leaf index pages applied the wrong comparison function, leading to wrong results and potentially crashes.
-  Fix unguarded recursion and loops in `contrib/hstore_plperl`, `contrib/jsonb_plperl`, and `contrib/jsonb_plpython` (Aleksander Alekseev) [&sect;](https://postgr.es/c/b0f3465b6) [&sect;](https://postgr.es/c/60a1d712a)

   Prevent stack overflow when dealing with deeply nested `jsonb` values, and allow interruption of the infinite loop caused when attempting to dereference circular chains of Perl object references.
-  Fix missed release of statistics catcache entry in `contrib/intarray` (Man Zeng) [&sect;](https://postgr.es/c/273be484a)

   This oversight led to warnings like “resource was not closed: cache pg_statistic”.
-  In `contrib/ltree`, fix integer overflow in comparisons (Ayush Tiwari) [&sect;](https://postgr.es/c/aca944e33)

   `ltree` values containing more than about 14,653 labels resulted in wrong comparison answers due to overflow. If a btree index contains such values, it is probably corrupt and should be reindexed after installing this update.
-  Fix array overrun in `contrib/pg_surgery`'s `heap_force_kill` and `heap_force_freeze` functions (Michael Paquier) [&sect;](https://postgr.es/c/daf8bc7d4)

   Attempting to change a TID whose offset number equals MaxHeapTuplesPerPage wrote one byte past the end of the allocated array, potentially crashing the server.
-  In `contrib/pg_surgery`, avoid infinite loop with TID arrays having more than 64K elements (Andrey Rachitskiy) [&sect;](https://postgr.es/c/eb28b6f24)
-  Avoid NULL-pointer dereference in `contrib/refint`'s `check_foreign_key()` (Ayush Tiwari) [&sect;](https://postgr.es/c/eb93f10e2)

   In the on-update-cascade case, a null value of a referenced column led to a crash. This is an oversight in the fix for CVE-2026-6637, but the code that was there before that wasn't really right either.
-  Fix `contrib/seg` to print segments with `~` certainty indicators correctly (Ewan Young) [&sect;](https://postgr.es/c/504ca0513)

   Due to a typo, `seg_out()` did not print a `~` certainty indicator attached to a segment's upper boundary. Worse, if the lower boundary had `~` while the upper boundary had no indicator, the upper boundary was not printed at all, incorrectly converting the value into an open interval.
-  Fix crash with namespace nodes in `contrib/xml2`'s `xpath_nodeset()` function (Andrey Chernyy, Michael Paquier) [&sect;](https://postgr.es/c/18955d412)
-  Support building PostgreSQL with Visual Studio 2026 (Andrew Dunstan) [&sect;](https://postgr.es/c/bc5a873a8)
-  Support building PostgreSQL with OpenSSL 4 (Daniel Gustafsson) [&sect;](https://postgr.es/c/fb8befa7b)
-  Update time zone data files to tzdata release 2026c (Tom Lane) [&sect;](https://postgr.es/c/796c275a0)

   Alberta (America/Edmonton) will be on year-round UTC-06 (effectively, permanent DST) beginning in November 2026. This release assumes that their TZ abbreviation will be `CST` from that time forward. That seems likely to change, but it's unclear what new abbreviation will be used.

   Morocco (Africa/Casablanca) will move to permanent UTC+00, without daylight saving time transitions, on 2026-09-20.
