<a id="release-17-11"></a>

## Release 17.11


**Release date:.**


2026-08-13


 This release contains a variety of fixes from 17.10. For information about new features in major release 17, see [Release 17](release-17.md#release-17).
 <a id="release-17-11-migration"></a>

### Migration to Version 17.11


 A dump/restore is not required for those running 17.X.


 However, the first three security entries below describe configuration adjustments and data cleanups that you may need to make after updating.


 Also, if you use `contrib/btree_gist` or `contrib/ltree`, you may need to reindex indexes made with those extensions; see the relevant entries below.


 Also, if you are upgrading from a version earlier than 17.6, see [Release 17.6](release-17-6.md#release-17-6).
  <a id="release-17-11-changes"></a>

### Changes


-  Restrict logical decoding output plugins to the set specified by a new server parameter `output_plugin_libraries` (Jacob Champion) [&sect;](https://postgr.es/c/01992176e) [&sect;](https://postgr.es/c/4fcccea97)

   Previously, a replication user could select any loadable library for logical decoding, allowing exploits of various sorts. To allow locking this down without breaking setups that worked before, introduce a whitelist of allowed output plugins.

   By default, only the output plugins shipped as part of PostgreSQL (`pgoutput` and `test_decoding`) are included in `output_plugin_libraries`. Installations that rely on other output plugins must add them after updating the server, for example

```

output_plugin_libraries = 'pgoutput, test_decoding, my_trusted_decoder'
```


   Additionally, `pg_upgrade --check` will fail if the `output_plugin_libraries` parameter on the new cluster does not permit the plugins of logical replication slots on the old cluster, when migrating from versions 17 and later. Make necessary additions to the new cluster's setting before performing pg_upgrade.

   The PostgreSQL Project thanks Vladimir Tokarev and Yu Kunpeng for reporting this problem. (CVE-2026-6471)
-  Fix `contrib/pgcrypto`'s PGP encryption to detect unsupported ciphers (Daniel Gustafsson) [&sect;](https://postgr.es/c/b05cfc693) [&sect;](https://postgr.es/c/7eed42aa8)

   Previously, if OpenSSL rejected the requested cipher (for example, because it is running in FIPS mode, or the legacy provider hasn't been loaded), pgcrypto failed to notice the failure and simply XOR'd the non-encrypted block with the plaintext, rendering the “encryption” trivially breakable. This will typically occur with deprecated or non-FIPS cipher algorithms (cipher-algo=blowfish/bf, twofish, cast5, or 3des).

   By default, pgcrypto will now fail to decrypt any messages that were affected in this way. To allow retrieval of such data, a new option `ignore-cipher-failure` has been added to `pgp_pub_decrypt()` and `pgp_sym_decrypt()`. Setting `ignore-cipher-failure=1` will restore their previous behavior, allowing the faulty encryption wrapper to be stripped off:

```

pgp_sym_decrypt(encrypted_column, ANY KEY, 'ignore-cipher-failure=1')
```
   Once the affected messages are identified and stripped of their wrappers, they can then be re-encrypted with a modern algorithm. It is important however that the behavior of OpenSSL be the same as it was when the faulty messages were created: if the set of unsupported algorithms is not the same, this approach will not work. See the documentation for `ignore-cipher-failure`.

   The PostgreSQL Project thanks Shishir Sharma for reporting this problem. (CVE-2026-14663)
-  Fix psql to skip in-line data following a scripted `COPY ... FROM STDIN` command, even if the `COPY` fails before sending `PGRES_COPY_IN` (Tom Lane) [&sect;](https://postgr.es/c/46fa1f6f3) [&sect;](https://postgr.es/c/dca6627de)

   Previously, if a `COPY` command failed at startup (for instance, because the target table doesn't exist) psql would not realize that and would proceed to read the following in-line data as SQL commands. In the best case that's wrong and in the worst case it's a SQL-injection hazard. Teach psql to recognize syntactically-valid `COPY ... FROM STDIN` commands and to skip data on its own authority if the server doesn't respond with `PGRES_COPY_IN`.

   While this fix is unlikely to affect any production SQL scripts, test scripts might intentionally exercise failing `COPY ... FROM STDIN` commands. Those will need to gain a `\.` data terminator line after each such command.

   The PostgreSQL Project thanks Alexander Lakhin for reporting this problem. (CVE-2026-6464)
-  Cross-check the output row type of a portal running `EXECUTE` or `FETCH` (Robert Haas) [&sect;](https://postgr.es/c/60887d348)

   `EXECUTE` and `FETCH` use two portals: an outer one for the statement itself, and an inner one running the query being executed on its behalf. It was previously possible to make the declared row types of the two portals diverge, leading to server memory disclosure and arbitrary code execution.

   The PostgreSQL Project thanks Ben Morris (in collaboration with Claude and Anthropic Research) and Peter Geoghegan for reporting this problem. (CVE-2026-16239)
-  Fix buffer overrun with long time zone abbreviation in `to_char()` (Tom Lane) [&sect;](https://postgr.es/c/12a620686)

   This can easily crash the server, and exploits leading to arbitrary code execution have been reported.

   The PostgreSQL Project thanks Hcamael, Amjad Shahzad, Tan Zhen of AntAISecurityLab, Tomer Fichman, Zheng Yu, Amy Burnett (OpenAI Codex Security), Rick de Jager, Heewon Song, Sylvie Mayer, Aleksander Alekseev, and Hillai Ben Sasson for reporting this problem. (CVE-2026-14669)
-  Fix buffer overrun in regexp match/split functions (Masahiko Sawada) [&sect;](https://postgr.es/c/e91dcfcca)

   If passed invalidly-encoded data, these functions could write past the end of their conversion buffer.

   The PostgreSQL Project thanks Francesco Verardi for reporting this problem. (CVE-2026-14664)
-  Harden the `ascii()` function against invalid input (Michael Paquier) [&sect;](https://postgr.es/c/849da8210)

   By supplying invalidly-encoded input, this function could be coaxed to read and return a few bytes of data that it shouldn't. In assert-enabled builds, its assertions could be triggered too.

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-18024)
-  Make `scalarineqsel()` check that a constant it expects to be of type `tid` actually is (Tom Lane) [&sect;](https://postgr.es/c/0ebf896f4)

   This expectation will hold for all the built-in operators that use this estimator, but a maliciously-constructed operator could violate it, leading to a crash or server memory disclosure.

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-14668)
-  Harden `tsvector` and `tsquery` code against overly long values (both individual lexemes and total vector/query length) (Tom Lane) [&sect;](https://postgr.es/c/fe0b5bd6d) [&sect;](https://postgr.es/c/8e87bd473)

   The documented limits were not enforced in all code paths.

   The PostgreSQL Project thanks Yuhang Wu, Zhenpeng Lin, Zheng Yu, and Hcamael for reporting these problems. (CVE-2026-14662)
-  Fix various places that mistakenly assumed they would not have to deal with more than `FUNC_MAX_ARGS` function arguments (Tom Lane) [&sect;](https://postgr.es/c/797e3cc4c) [&sect;](https://postgr.es/c/8a40f4b09)

   Notably, the server's actual limit on the number of arguments to an aggregate function is `FUNC_MAX_ARGS - 1`, but the parser failed to enforce that, creating hazards downstream.

   The PostgreSQL Project thanks Zheng Yu, ylwangtju, and Masahiko Sawada for reporting these problems. (CVE-2026-14679)
-  Reject calls from SQL to functions that take or return type `internal` (Tom Lane) [&sect;](https://postgr.es/c/eb9e55297) [&sect;](https://postgr.es/c/83d0a083f)

   The existing defenses against doing this have been shown to be insufficient, so add more explicit checks.

   The PostgreSQL Project thanks Amy Burnett (OpenAI Codex Security) for reporting this problem. (CVE-2026-14680)
-  Preserve the ownership of extended statistics objects when they are rebuilt by `ALTER TABLE` (Masahiko Sawada) [&sect;](https://postgr.es/c/75a03c569)

   Previously, the role running `ALTER TABLE` gained ownership of such objects, but that seems inappropriate.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-6469)
-  When deparsing an `EXTRACT()` function call, quote the field name if needed (Nathan Bossart) [&sect;](https://postgr.es/c/5981fe370)

   The parser accepts any string literal as a field name in `EXTRACT()`, deferring validation to execution. If the call is stored and deparsed (for example during pg_dump), the string body was regurgitated verbatim, allowing SQL injection.

   The PostgreSQL Project thanks Ben Morris (in collaboration with Claude and Anthropic Research) for reporting this problem. (CVE-2026-15741)
-  Check for `USAGE` privilege on data types in places that formerly failed to check that (Nathan Bossart) [&sect;](https://postgr.es/c/fd6d3e7e9) [&sect;](https://postgr.es/c/671359605) [&sect;](https://postgr.es/c/d1c8aa0b0)

   `CREATE TYPE AS RANGE` did not check, nor did `ALTER TABLE OF`, nor did commands that create stored expressions. These omissions allowed roles without `USAGE` privilege to nonetheless create objects depending on the type, possibly blocking the type's owner from changing the type later.

   The PostgreSQL Project thanks Jingzhou Fu for reporting this problem. (CVE-2026-6470)
-  Invalidate role-dependent cached plans after role changes (Ilya Staroverov, Shinya Kato, Nathan Bossart) [&sect;](https://postgr.es/c/1974acf23)

   Role membership, role attribute, and database ownership changes may impact the expected behavior of row-level security policies, but previously we'd continue to use cached plans that were made according to the old state of affairs.

   The PostgreSQL Project thanks Ilya Staroverov and Shinya Kato for reporting this problem. (CVE-2026-14666)
-  Reject GSSEncRequest after direct SSL connection (Michael Paquier) [&sect;](https://postgr.es/c/067a64d40)

   After establishing a TLS-encrypted connection, the server would still accept a request for GSSAPI encryption. If that succeeded, the connection would proceed using TLS encryption, but it would look like a GSS connection to the `pg_hba` rules. Thus, a `pg_hba` policy intending to disallow TLS would not be enforced correctly.

   The PostgreSQL Project thanks p4p3r for reporting this problem. (CVE-2026-14681)
-  Make mock SCRAM authentication secrets more plausible (Nathan Bossart) [&sect;](https://postgr.es/c/dec60e8ad)

   If a SCRAM login is attempted against a role that doesn't exist or doesn't have a SCRAM secret, we generate a mock secret and carry out the authentication handshake anyway, to avoid revealing these facts to an attacker. But the mock secret was made with a fixed iteration count, which in itself can be an observable response discrepancy. Use the configuration setting `scram_iterations` instead, to make the mock secret look more like the installation's real secrets.

   The PostgreSQL Project thanks Radim Marek for reporting this problem. (CVE-2026-14672)
-  Fix out-of-bounds writes in ecpg applications caused by invalid `bytea` data received from the server (Michael Paquier) [&sect;](https://postgr.es/c/c3c830272)

   ecpg assumed without checking that any `bytea` value must begin with `\x`. A broken or malicious server might send a string shorter than 2 bytes, resulting in memory clobber in the application.

   The PostgreSQL Project thanks ylwangtju for reporting this problem. (CVE-2026-16241)
-  Do not do backquote expansion on the argument of psql's `\unrestrict` command (Nathan Bossart) [&sect;](https://postgr.es/c/0bfac9e1f)

   This oversight in the fix for CVE-2025-8714 allows a malicious server to inject shell commands into plain-text dump output that will be run at restore time on the machine running psql, the exact scenario that CVE-2025-8714 intended to prevent.

   The PostgreSQL Project thanks Lucas Velgus, Filip Janus, and Daniel Bakker for reporting this problem. (CVE-2026-18408)
-  Remove pg_dump's assumption that `pg_proc`.`protrftypes` cannot have more than `FUNC_MAX_ARGS` entries (Tom Lane) [&sect;](https://postgr.es/c/c2b16f5d4)

   Since there could be entries for both input and output arguments, it's feasible for this array's length to exceed `FUNC_MAX_ARGS` (which constrains only input arguments). Even if that were not so, pg_dump cannot assume that the server was built with the same value of `FUNC_MAX_ARGS` that it has. An overrun would lead to a memory clobber inside pg_dump.

   The PostgreSQL Project thanks Masahiko Sawada for reporting this problem. (CVE-2026-19385)
-  Harden PL/Perl against “tied” Perl arrays and hashes (Tom Lane) [&sect;](https://postgr.es/c/2d78c34f8)

   A tied object that doesn't behave like a regular one could lead to memory overwrite, or to constructing a corrupt result array (which would likely cause problems later).

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-14670)
-  Fix integer overflows in memory-allocation calculations in PL/Perl and PL/Tcl (Heikki Linnakangas) [&sect;](https://postgr.es/c/b56cc7deb)

   This is the same type of problem as CVE-2026-6473, just in a different part of the code, and is fixed in the same way.

   The PostgreSQL Project thanks the Tulya Project (Team Dhiutsa, Bitecope Technologies Private Ltd) for reporting this problem. (CVE-2026-14677)
-  Ensure that `contrib/amcheck` functions restrict `search_path` before executing index expressions (Noah Misch) [&sect;](https://postgr.es/c/e41c72aff)

   Because amcheck will run such index expressions as the owner of their tables, a caller could potentially hijack `search_path`-dependent functions to run arbitrary code as the table owner. By default this is not a vulnerability because only superusers are allowed to call amcheck functions; but if that privilege was granted out, it created a larger hazard than the documentation suggests.

   The PostgreSQL Project thanks Yuelin Wang and Jacob Brazeal for reporting this problem. (CVE-2026-14673)
-  Fix integer overflows in `contrib/fuzzystrmatch`'s `levenshtein()` and `levenshtein_less_equal()` functions (Nathan Bossart) [&sect;](https://postgr.es/c/c4d51b627)

   Passing large cost values to these functions could cause integer overflows, thereby producing nonsensical results, and even causing out-of-bounds writes in some cases.

   The PostgreSQL Project thanks Ben Morris (in collaboration with Claude and Anthropic Research) for reporting this problem. (CVE-2026-15742)
-  Fix datatype error in `contrib/pg_trgm`'s GiST picksplit function (Heikki Linnakangas) [&sect;](https://postgr.es/c/1af08af69)

   This mistake resulted in reading past the end of the buffer, typically causing bad split decisions; but a crash could ensue if you're very unlucky.

   The PostgreSQL Project thanks Mehmet D. Ince for reporting this problem. (CVE-2026-14678)
-  Remove the plan cache in `contrib/refint` (Ayush Tiwari) [&sect;](https://postgr.es/c/66a5146dc)

   This caching behavior has several serious bugs, notably that `check_foreign_key()` embeds the new key values in its cascade-UPDATE queries, so a cached plan reuses the originally-needed values rather than the key values that should be used. The simplest solution is to remove it.

   The PostgreSQL Project thanks Hcamael for reporting this problem. (CVE-2026-14671)
-  Fix mis-handling of asynchronous reads when rescanning an asynchronous Append plan node (Alexander Korotkov, Gleb Kashkin, Etsuro Fujita) [&sect;](https://postgr.es/c/3704870b2)

   When an upper plan node rescans an Append before having read the entire Append output, we need to discard any in-flight requests sent to external servers (by `postgres_fdw` for example). This was not done correctly in cases where a subplan has parameter changes or is discarded by partition pruning in the next scan. The outcome could be incorrect query results, an infinite loop, or an assertion failure.
-  Fix error in partition pruning for RANGE-partitioned tables (David Rowley) [&sect;](https://postgr.es/c/31f2acde5)

   In some cases the DEFAULT partition would be skipped when it should not be, which could lead to rows missing from query results.
-  Fix planner's nullability and strictness checks for <em>value</em><code> IN
      (</code><em>array</em><code>)</code> expressions (Ayush Tiwari) [&sect;](https://postgr.es/c/26d6b7dc9)

   These checks should only succeed if the array operand is known to be non-empty, but that consideration was missed, allowing optimizations to be applied that should not be. This could result in wrong query answers if the array actually was empty.
-  Fix incorrect join removal logic (Matheus Alcantara, Richard Guo) [&sect;](https://postgr.es/c/cdcec567d) [&sect;](https://postgr.es/c/b308eb366)

   In edge cases, it was possible for a constant output value coming from within the nullable side of an outer join to not be replaced by NULL when it should be.
-  Add missed checks for hashability of equality comparisons on container datatypes (arrays, composites, ranges) (Andrei Lepikhov, Tom Lane) [&sect;](https://postgr.es/c/19152e3c2)

   The planner must verify hashability of the container's component type(s) before deciding it can use a hash-based plan type. This step was missed in some places, leading to “could not identify a hash function” failures at execution.
-  Fix mis-optimization of `COUNT` window functions that have an `EXCLUDE` clause or lack `ORDER BY` (Chengpeng Yan, David Rowley) [&sect;](https://postgr.es/c/be63b285e)

   These window functions were treated as monotonic when they should not be, allowing wrong answers to be computed.
-  Fix `ALTER COLUMN ... DROP EXPRESSION` to work when there are multiple levels of partitions (Alberto Piai) [&sect;](https://postgr.es/c/18006c1bd)
-  Fix attaching partitions of indexes that are exclusion constraints (Japin Li) [&sect;](https://postgr.es/c/1d6c654c8)

   Notably, this oversight broke dump/restore of partitioned exclusion constraints.
-  Disallow renaming a rule to `_RETURN` (Tom Lane) [&sect;](https://postgr.es/c/e99fb3262)

   That name is reserved for a view's `ON SELECT` rule, but `ALTER RULE` allowed renaming other rules to `_RETURN`, causing trouble later.
-  Fix missing lock release for role membership grants in `DROP OWNED BY` (Jeff Davis) [&sect;](https://postgr.es/c/25e54cec7)

   This oversight resulted in a warning message, followed by retaining a lock on the membership grant until the end of the transaction.
-  Fix failure of `EXPLAIN` when deparsing `SQL/JSON` aggregates (Richard Guo) [&sect;](https://postgr.es/c/dcda1f07d)

   Some plan structures resulted in “invalid JsonConstructorExpr underlying node type” errors.
-  Fix use of `REINDEX CONCURRENTLY` with a deferred uniqueness constraint (Nitin Motiani) [&sect;](https://postgr.es/c/28269fed6)

   The transient index copy created during `REINDEX CONCURRENTLY` was incorrectly marked as enforcing immediate uniqueness, causing spurious reports of constraint violation.
-  Fix matching of localized month/day names in `to_date()` (Heikki Linnakangas) [&sect;](https://postgr.es/c/8acfaa12a)

   The matching logic misbehaved in cases where case-folding changes the byte length of the string.
-  Fix incorrect NFC recomposition for Hangul U+11A7 (TBASE) (Diego Frias, Michael Paquier) [&sect;](https://postgr.es/c/0c9cbbfb5)

   This character was treated as a valid T syllable, which it is not, and hence silently swallowed during normalization.
-  Avoid possible truncation of output lexemes in case-insensitive `synonym` dictionaries (Jeff Davis) [&sect;](https://postgr.es/c/3805641cb)

   If folding to lower case increased the byte length of a lexeme, it was incorrectly truncated to its original byte length when emitted.
-  Defend against truncated UTF-8 characters in case-conversion logic (Jeff Davis) [&sect;](https://postgr.es/c/5e78ebca5)
-  Fix typo in `hash_record_extended()` (Man Zeng) [&sect;](https://postgr.es/c/203e238bb)

   The code failed to initialize the second isnull argument passed to FunctionCallInvoke(). This is harmless for existing in-core extended hash support functions, which will not examine that value. However, extension-provided hash functions could be affected if they inspect `PG_ARGISNULL(1)`.
-  Fix `pg_get_publication_tables()` to not fail if a publishable table is dropped concurrently (Bharath Rupireddy) [&sect;](https://postgr.es/c/dcbc96685)
-  Prevent `satisfies_hash_partition()` from crashing with `VARIADIC NULL` (Robert Haas) [&sect;](https://postgr.es/c/52af6fef4)
-  Report invalid-weight errors more cleanly and consistently in `tsvector_filter()` and allied functions (Ewan Young) [&sect;](https://postgr.es/c/0626fbfeb)

   In particular, report weight characters that are not printable ASCII in octal form (<code>\</code><em>nnn</em>), as `charout()` would render them. This avoids possibly producing an invalidly-encoded error message.
-  Fix mishandling of namespace nodes in `xpath()` (Michael Paquier) [&sect;](https://postgr.es/c/6c08cbb7a) [&sect;](https://postgr.es/c/940916549)

   This fix avoids an unexpected “could not copy node” error.
-  Fix `jsonpath`'s `.decimal` method to not throw a hard error for incorrect precision or scale (Ewan Young) [&sect;](https://postgr.es/c/ab35b8d25) [&sect;](https://postgr.es/c/c768637d6)

   Silent mode should suppress these errors, but failed to.
-  Fix NULL-pointer crash when `IS JSON` or similar constructs have an argument that is of string category but lacks a cast to type `text` (Ayush Tiwari) [&sect;](https://postgr.es/c/d0acd2535)

   There are no such data types in core PostgreSQL, but the problem is reachable with some extension types.
-  Ensure that `SQL/JSON` `ON EMPTY / ON ERROR DEFAULT` values are coerced to the correct typmod (Ewan Young) [&sect;](https://postgr.es/c/71cd10cd2)

   For example, the declared precision and scale of a `numeric` target column were not applied to the default value.
-  Avoid machine-dependent behavior when dividing the smallest possible `money` value by -1 (Andrey Rachitskiy) [&sect;](https://postgr.es/c/1416f304d)
-  Fix crash after out-of-memory failure partway through creation of a cache entry for a text search dictionary (Tom Lane) [&sect;](https://postgr.es/c/634a8dcb8)
-  Fix memory-safety bugs in processing of incorrect ispell/hunspell dictionary files (Andrey Rachitskiy) [&sect;](https://postgr.es/c/5fdea3aa3)
-  Prevent access to other sessions' temporary tables (Jim Jones, Daniil Davydov, Alexander Korotkov) [&sect;](https://postgr.es/c/4dfae59a1) [&sect;](https://postgr.es/c/4e49f68b7)

   Some code paths failed to prevent this, leading to silently wrong (inconsistent) results.
-  Fix the order in which autovacuum processes databases (Rustam Khamidullin) [&sect;](https://postgr.es/c/288d4e83f)

   It was unintentionally processing databases from lowest to highest score, when it should be doing the reverse.
-  Fix memory leak in parallel vacuum worker processes (Baji Shaik) [&sect;](https://postgr.es/c/8ad414831)

   Progress reports from a parallel worker leaked about 1kB per report, with the waste accumulating for the life of the worker process.
-  Honor query cancel and vacuum delay during GIN index posting-tree cleanup (Paul Kim, Alexander Korotkov) [&sect;](https://postgr.es/c/ba5e46329)

   The posting tree for a common value can be large, so that this missed check could allow vacuum to run for a long time before noticing an interrupt.
-  Fix possible mis-decoding of index tuples during GiST and SP-GiST index-only scans (Peter Geoghegan) [&sect;](https://postgr.es/c/355faed5a)

   This error could lead to emitting corrupted data from an index-only scan plan. The only affected core opclass is GiST's range_ops, and it could only fail if the range column were not the first index column.
-  Ensure that the new last block of a bulk-extended table is added to its free space map promptly (Jingtang Zhang) [&sect;](https://postgr.es/c/768ae083e)

   An off-by-one error caused the last block of a multi-block table extension to not be marked as free in the map. This would eventually get corrected by vacuum, but meanwhile the space wouldn't be used.
-  Avoid possible double-free or infinite error recovery loop in resource cleanup during transaction abort (Tom Lane) [&sect;](https://postgr.es/c/011eedcdc)
-  When creating directories, tolerate concurrent creation of the same directory (Andrew Dunstan, Tom Lane) [&sect;](https://postgr.es/c/f0a831ef3)
-  Prevent creation of dangling object dependencies by acquiring a shared lock on any object being depended on (Bertrand Drouvot) [&sect;](https://postgr.es/c/3a9909eda) [&sect;](https://postgr.es/c/c1588f92a)

   The shared lock will conflict with any attempt to drop the depended-on object, eliminating the race condition that formerly existed. For example, if one session drops a schema (that appears empty to it) concurrently with some other session creating a function in that schema, previously both transactions could commit, leaving an invalid function definition behind. Now, one transaction or the other will fail.
-  Fix race condition in conflict detection for `SERIALIZABLE` isolation mode (Peter Geoghegan) [&sect;](https://postgr.es/c/8434c9385)

   A conflict could be missed when examining an initially-empty btree index, allowing failure of serializability due to improperly allowing conflicting transactions to commit.
-  Fix race condition in ProcSignalBarrier code (Masahiko Sawada) [&sect;](https://postgr.es/c/a651b8a89)

   This error could result in processes getting stuck, typically after reporting “still waiting for backend with PID *nnnn* to accept ProcSignalBarrier”.
-  Fix race conditions when a set of processes that belong to the same lock group exit at the same time (Vlad Lesin) [&sect;](https://postgr.es/c/d489c4439) [&sect;](https://postgr.es/c/e14b4ea4a)

   These errors could lead to PANIC aborts, with messages such as “latch already owned”. The issue does not normally arise in regular parallel query, since the leader won't exit before seeing its workers finish; but some extensions reach the problem.
-  Fix WAL logging of operations that clear bits in tables' visibility maps (Melanie Plageman, Andres Freund) [&sect;](https://postgr.es/c/7edec8b57) [&sect;](https://postgr.es/c/c0d9864f5) [&sect;](https://postgr.es/c/067213430)

   Such VM changes were missed by the WAL summarizer, potentially leading to incorrect incremental backups. We also failed to log full-page images of such VM pages when needed, potentially allowing torn page writes to go uncorrected. This could lead to misbehavior later, such as wrong results from index-only scans.
-  Prevent WAL summarizer process from getting stuck at a timeline switch (Robert Haas) [&sect;](https://postgr.es/c/499d9eabb) [&sect;](https://postgr.es/c/d28cdf46e)
-  Fix race with timeline selection in logical decoding during standby promotion (Bertrand Drouvot) [&sect;](https://postgr.es/c/16b89ff04) [&sect;](https://postgr.es/c/ab5334d8b)

   Logical decoding being performed on the standby could fail with a “requested WAL segment has already been removed” error. A repeat attempt would succeed, so there was no permanent problem but there was an availability hazard.
-  Avoid exposing a WAL receiver's full connection string during timeline jumps (Chao Li) [&sect;](https://postgr.es/c/c89499a79)

   The `pg_stat_wal_receiver` view should show a sanitized version of the connection string, without sensitive data. But it transiently showed the full string when we re-use an existing WAL receiver.
-  Use run-time checks, not just Asserts, to verify the correct number of columns in tuples received during logical replication (Varik Matevosyan) [&sect;](https://postgr.es/c/15f4e3d0c)

   A malicious or buggy publisher could send inconsistent numbers of columns. While we could not find a scenario in which this would have serious ill effects, extra caution seems warranted.
-  Clean up quoting of string parameters within constructed replication commands (Tom Lane) [&sect;](https://postgr.es/c/14810cc0d)

   Various places that generate replication commands were not being adequately careful about quoting replication slot names and other parameters that need to be inserted into those commands. This could result in unexpected syntax errors in those commands. In principle, a crafted replication slot name could result in SQL injection; but such a scenario seems very unlikely to occur in practice, since replication operations can only be invoked by highly-privileged users and there is no reason for them to use a slot name coming from an untrustworthy source.
-  Fix logical decoding of empty prepared transactions (Masahiko Sawada) [&sect;](https://postgr.es/c/8c4519c71)

   A prepared transaction that did not cause any decodable updates could result in sending `COMMIT/ROLLBACK PREPARED` to the output plugin with no preceding `PREPARE`. For the built-in subscriber this breaks replication, and other plugins will probably not like it either.
-  Fix corruption of unlogged sequences after standby promotion (Fujii Masao) [&sect;](https://postgr.es/c/a1ed6a9a0)

   Previously, if an unlogged sequence was created on the primary and replicated to a standby, accessing the sequence after promoting the standby could fail with “bad magic number in sequence” or related errors.
-  Fix cascading standby reconnect failure after archive fallback (Marco Nenciarini) [&sect;](https://postgr.es/c/2cf28d1b9)

   A cascading standby could fail to reconnect to its upstream standby with “requested starting point ... is ahead of the WAL flush position” after falling back to archive recovery.
-  Do not try to clear `pg_database`.`dathasloginevt` locally on a standby server (Ayush Tiwari) [&sect;](https://postgr.es/c/4a375527a)

   Event trigger cleanup tried to perform that action on standby servers as well as the primary. That can't work on a standby, and there's no need anyway since replay of the primary's database change will soon fix it.
-  Avoid race condition while dropping obsolete replication slots (Xuneng Zhou) [&sect;](https://postgr.es/c/ea834d747)

   An incorrect unlock and log message could occur if another session immediately re-used the dropped slot's shared-memory entry.
-  Avoid race condition while dropping ephemeral replication slots (Zhijie Hou) [&sect;](https://postgr.es/c/080d61f07)

   The slot-releasing code performed some additional updates to the replication slot's shared-memory entry after releasing the slot. This is unsafe since another session could immediately re-use the dropped slot's shared-memory entry. Skip those updates in the case of an ephemeral slot.
-  Fix stale progress reports during logical replication table synchronization (Shinya Kato) [&sect;](https://postgr.es/c/f2acab534)

   Previously, the `pg_stat_progress_copy` view in the subscriber would continue to show the initial `COPY` operation as active even after the data copy had finished. The stale entry remained visible until synchronization caught up with the publisher.
-  Clear base backup progress on backup failure (Chao Li) [&sect;](https://postgr.es/c/e2ea6bfed) [&sect;](https://postgr.es/c/a240e8cd2)

   Previously the `pg_stat_progress_basebackup` view would continue to show a stale progress entry after a failure, until the replication client disconnected. pg_basebackup normally disconnects immediately, but other clients might not.
-  Fix possible PANIC due to concurrent drop of pgstats entries when `track_functions` is enabled (Sami Imseih, Michael Paquier) [&sect;](https://postgr.es/c/2e0c61aed) [&sect;](https://postgr.es/c/39e649d44) [&sect;](https://postgr.es/c/afb076b29)
-  Clean up broken local pgstats entry after failing to obtain space for the corresponding shared hashtable entry (Niall Newman) [&sect;](https://postgr.es/c/fc9283b21)

   Failure to do this led to a null-pointer dereference the next time the local entry was used.
-  In PL/Perl, avoid NULL pointer dereference crash when working with an invalid `PostgreSQL::InServer::ARRAY` object (Xing Guo) [&sect;](https://postgr.es/c/d424d06ed)
-  In PL/Python, properly check for errors when working with sequence and mapping objects (Richard Guo) [&sect;](https://postgr.es/c/3dc59c173)

   Previously, a broken object or an unhandled exception could result in a NULL pointer dereference crash.
-  In libpq, always drain all pending bytes from the SSL or GSS decryption buffer during `pqReadData()` (Jacob Champion) [&sect;](https://postgr.es/c/c162810a5) [&sect;](https://postgr.es/c/177a2a2e3) [&sect;](https://postgr.es/c/0958c3ea4) [&sect;](https://postgr.es/c/bab0a2db7)

   This avoids edge cases where libpq or its calling application waits for more data to arrive on the socket, but actually all the data has already arrived.
-  Allow libpq to accept ParameterDescription messages exceeding 30000 bytes (Ning Sun) [&sect;](https://postgr.es/c/e4183c667)

   Previously, this message type was not among those that libpq's validity heuristics believed could be long. The limit resulted in failure for prepared queries having more than 7498 parameters, which is unlikely but supported.
-  Reject multiple descriptor header items in ecpg's `GET/SET DESCRIPTOR` statements (Masashi Kamura) [&sect;](https://postgr.es/c/fe8c0a762)

   Previously the grammar allowed this syntax, but broken C code was generated. Adjust the grammar and the documentation to allow only one header item.
-  Make line widths match in psql's expanded aligned output format (Pavel Stehule) [&sect;](https://postgr.es/c/efd885d05)

   When the table's data rows are narrower than the record header lines, widen the data rows to match the headers, avoiding unsightly output.
-  Fix psql's privilege check for showing database size in `\l+` (Christoph Berg) [&sect;](https://postgr.es/c/41949c6f3)

   The underlying server function permits users who have `pg_read_all_stats` privileges to see the sizes of all databases, even if they lack `CONNECT` privilege. But psql was unaware of that provision and would not call the function unless the user has `CONNECT` privilege.
-  Fix psql's tab completion for `\df` to consider procedures too (Erik Wienhold) [&sect;](https://postgr.es/c/355a4fcf9)
-  Fix thread-safety bug in pgbench (Fujii Masao) [&sect;](https://postgr.es/c/52b3e7001)

   When pgbench runs with multiple threads and the `--verbose-errors` option, different threads could attempt to use the same buffer to construct error messages, leading to corrupted log output.
-  In pg_combinebackup, prevent infinite loop if the source file is shorter than expected (Peter Eisentraut) [&sect;](https://postgr.es/c/090ce6934)
-  Fix cleanup of publisher-side objects after errors in pg_createsubscriber (Nisha Moond) [&sect;](https://postgr.es/c/c03784a21)

   When pg_createsubscriber fails after creating logical replication objects, it should remove the publication and replication slot that it created on the publisher. Some error cases failed to do so.
-  Use the source cluster's group-read file permissions for pg_recvlogical output files (Fujii Masao) [&sect;](https://postgr.es/c/ddd12d1a5)

   pg_recvlogical was documented to behave this way, but it never actually enabled group-read.
-  In `contrib/amcheck`, handle short-header varlena datums correctly (Andrey Borodin) [&sect;](https://postgr.es/c/8abb8a155)

   This error could result in doing excess work while verifying a btree index, but seems not to have had any worse consequences.
-  In `contrib/btree_gist`, fix `NaN` handling in the `float4` and `float8` opclasses (Bill Kim, Tom Lane) [&sect;](https://postgr.es/c/d215d2cc2)

   Comparisons, as well as the GiST penalty and distance functions, did not account for `NaN` and would give the wrong answer when handed one. It is recommended to reindex `btree_gist` indexes on float columns after installing this update, if there is any possibility that there are `NaN` entries in those columns.
-  In `contrib/btree_gist`, fix searches using a not-equal operator (Ayush Tiwari) [&sect;](https://postgr.es/c/86992769e)

   For variable-length data types, the code for scanning non-leaf index pages applied the wrong comparison function, leading to wrong results and potentially crashes.
-  Fix unguarded recursion and loops in `contrib/hstore_plperl`, `contrib/jsonb_plperl`, and `contrib/jsonb_plpython` (Aleksander Alekseev) [&sect;](https://postgr.es/c/3df0b7755) [&sect;](https://postgr.es/c/4efef9d18)

   Prevent stack overflow when dealing with deeply nested `jsonb` values, and allow interruption of the infinite loop caused when attempting to dereference circular chains of Perl object references.
-  Fix missed release of statistics catcache entry in `contrib/intarray` (Man Zeng) [&sect;](https://postgr.es/c/94b57ab54)

   This oversight led to warnings like “resource was not closed: cache pg_statistic”.
-  In `contrib/ltree`, fix integer overflow in comparisons (Ayush Tiwari) [&sect;](https://postgr.es/c/c391c00d9)

   `ltree` values containing more than about 14,653 labels resulted in wrong comparison answers due to overflow. If a btree index contains such values, it is probably corrupt and should be reindexed after installing this update.
-  In `contrib/pgcrypto`, avoid double-free crash after encountering an error while using an OSSLCipher object (Yuelin Wang) [&sect;](https://postgr.es/c/2aa6be6e6)
-  Fix array overrun in `contrib/pg_surgery`'s `heap_force_kill` and `heap_force_freeze` functions (Michael Paquier) [&sect;](https://postgr.es/c/0bcf19c9e)

   Attempting to change a TID whose offset number equals MaxHeapTuplesPerPage wrote one byte past the end of the allocated array, potentially crashing the server.
-  In `contrib/pg_surgery`, avoid infinite loop with TID arrays having more than 64K elements (Andrey Rachitskiy) [&sect;](https://postgr.es/c/1b6c99d96)
-  Avoid NULL-pointer dereference in `contrib/refint`'s `check_foreign_key()` (Ayush Tiwari) [&sect;](https://postgr.es/c/6b4de201e)

   In the on-update-cascade case, a null value of a referenced column led to a crash. This is an oversight in the fix for CVE-2026-6637, but the code that was there before that wasn't really right either.
-  Fix `contrib/seg` to print segments with `~` certainty indicators correctly (Ewan Young) [&sect;](https://postgr.es/c/bcbbd070d)

   Due to a typo, `seg_out()` did not print a `~` certainty indicator attached to a segment's upper boundary. Worse, if the lower boundary had `~` while the upper boundary had no indicator, the upper boundary was not printed at all, incorrectly converting the value into an open interval.
-  Fix crash with namespace nodes in `contrib/xml2`'s `xpath_nodeset()` function (Andrey Chernyy, Michael Paquier) [&sect;](https://postgr.es/c/4a49ab289)
-  Support building PostgreSQL with OpenSSL 4 (Daniel Gustafsson) [&sect;](https://postgr.es/c/8bcabfcce)
-  Update time zone data files to tzdata release 2026c (Tom Lane) [&sect;](https://postgr.es/c/669438aed)

   Alberta (America/Edmonton) will be on year-round UTC-06 (effectively, permanent DST) beginning in November 2026. This release assumes that their TZ abbreviation will be `CST` from that time forward. That seems likely to change, but it's unclear what new abbreviation will be used.

   Morocco (Africa/Casablanca) will move to permanent UTC+00, without daylight saving time transitions, on 2026-09-20.
