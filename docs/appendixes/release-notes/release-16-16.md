<a id="release-16-16"></a>

## Release 16.16


**Release date:.**


2026-11-12


 This release contains a variety of fixes from 16.15. For information about new features in major release 16, see [Release 16](release-16.md#release-16).
 <a id="release-16-16-migration"></a>

### Migration to Version 16.16


 A dump/restore is not required for those running 16.X.


 However, several different bugs have been discovered that could result in index corruption. Examine the changelog entries below that mention `REINDEX`, and reindex any indexes you have that might be affected.


 Also, if you are upgrading from a version earlier than 16.15, see [Release 16.15](release-16-15.md#release-16-15).
  <a id="release-16-16-changes"></a>

### Changes


-  Add defenses against problems caused by altering composite types concurrently with operations on data of those types (Heikki Linnakangas, Andres Freund, Nathan Bossart, Nikhil Sontakke) [&sect;](https://postgr.es/c/c67a4a24b) [&sect;](https://postgr.es/c/9528f2f31) [&sect;](https://postgr.es/c/1d1650e52) [&sect;](https://postgr.es/c/148fc8249) [&sect;](https://postgr.es/c/e5c581cd4) [&sect;](https://postgr.es/c/b219eed24) [&sect;](https://postgr.es/c/7fe0b6c8e) [&sect;](https://postgr.es/c/62761302b) [&sect;](https://postgr.es/c/4a090c2a7) [&sect;](https://postgr.es/c/74681445e) [&sect;](https://postgr.es/c/5024810ad) [&sect;](https://postgr.es/c/cf3fcd765) [&sect;](https://postgr.es/c/1c0c88109) [&sect;](https://postgr.es/c/3429a10d4) [&sect;](https://postgr.es/c/ecdfe0307) [&sect;](https://postgr.es/c/2ce625266) [&sect;](https://postgr.es/c/726ef27b8) [&sect;](https://postgr.es/c/dd0278aa3) [&sect;](https://postgr.es/c/0dcf804f2) [&sect;](https://postgr.es/c/cfe611199) [&sect;](https://postgr.es/c/eb64484eb) [&sect;](https://postgr.es/c/7f8be5d19) [&sect;](https://postgr.es/c/813d92afc) [&sect;](https://postgr.es/c/0c8ffa1ad) [&sect;](https://postgr.es/c/788915adf) [&sect;](https://postgr.es/c/0453c5539) [&sect;](https://postgr.es/c/5968dfa67) [&sect;](https://postgr.es/c/8d0db2258)

   Operations such as `ALTER TABLE ... ALTER COLUMN TYPE` are hazardous because tuples of the older row layout may still exist in the system. While we have long had some checks to defend against that, they were never really complete, and they entirely failed to detect cases where composite types were changed concurrently with queries using them. This voluminous set of patches attempts to improve that.

   Introduce versioning logic to permit detection of composite datums that don't match the tuple descriptor we wish to interpret them with. For full protection, this change requires rebuilding C-language extensions, and there are some extensions that will require source-code changes. If necessary, the new `composite_type_mismatch` server parameter can be used to disable these checks, but doing so carries security risks.

   Add locking to prevent cases such as `ALTER TYPE` on a composite type used in a `CREATE TABLE` that hasn't committed yet.

   Tighten the rules about whether `ALTER COLUMN TYPE` is allowed when dependencies exist on the composite type. For example, it's no longer allowed when a function's default expression or new-style SQL function body depends on the type.

   The PostgreSQL Project thanks Andrew Gierth, Alexander Lakhin, Francesco Verardi, Hcamael, Tomer Fichman, OpenAI Security Research Team, Peter Geoghegan, Mehmet Ince, Bryan Green, Sumedh Pathak, Sergey Shinderuk, Yuxiao Guo, Nikhil Sontakke, Daniel Bakker, the Anthropic OSS program, Claude and Ada Logics for reporting this problem. (CVE-2026-6460)
-  Cope with concurrent rowtype changes in `ExecEvalFieldStore[De]Form()` (Tomas Vondra, Heikki Linnakangas) [&sect;](https://postgr.es/c/71c8a6ac4)

   The PostgreSQL Project thanks OpenAI Codex Security, Bronson Yen, Mehmet Ince, Claude and Ada Logics for reporting this problem. (CVE-2026-116306)
-  Enforce view row visibility in `INSERT ... ON CONFLICT` (Peter Geoghegan, Zsolt Parragi) [&sect;](https://postgr.es/c/629f54970)

   When applied to a view, `INSERT ... ON CONFLICT DO UPDATE` could update rows that should be hidden by the view's `WHERE` clause. Although it was documented, this behavior is surely unexpected, and it gives rise to security problems: a user can update rows he should have no privilege to reach. Updating is incorrect, but inserting is impossible too (it would fail due to the unique arbiter index), so throw an error instead when the case arises. `INSERT ... ON CONFLICT DO SELECT` could show such rows as well, but no longer does.

   The PostgreSQL Project thanks Qing Xu, Carl Andersen, Zsolt Parragi, the Anthropic OSS program, Shihao Zhong, and Josef Šimánek for reporting this problem. (CVE-2026-16237)
-  Align `REASSIGN OWNED`'s caller-privilege checks with `ALTER ... OWNER TO`'s checks (Zsolt Parragi) [&sect;](https://postgr.es/c/dbc8a36d7)

   Since v16, giving an object to another role requires the ability to `SET ROLE` to that role. However, a role granted `INHERIT TRUE, SET FALSE` could bypass that via `REASSIGN OWNED`. A role granted `INHERIT FALSE, SET TRUE` saw the inverse problem: `REASSIGN OWNED` failed while `ALTER OWNER` succeeded. Make `REASSIGN` follow the `ALTER OWNER` rules.

   `REASSIGN OWNED` also bypassed the `CREATEDB` privilege check in `ALTER DATABASE OWNER`.

   In v15 and v14, `REASSIGN OWNED`'s privilege check is stricter than the membership checks in `ALTER OWNER`. Leave that alone, but still fix the missing `CREATEDB` check.

   The PostgreSQL Project thanks Zsolt Parragi for reporting this problem. (CVE-2026-120146)
-  Disallow closing a portal that is still being set up (Michael Paquier) [&sect;](https://postgr.es/c/38783780f)

   This change forestalls various use-after-free issues.

   The PostgreSQL Project thanks Tomas Vondra for reporting this problem. (CVE-2026-116332)
-  Forcibly close portals associated with transition tables at the end of the trigger function (Jeff Davis) [&sect;](https://postgr.es/c/5b2bd1dc1)

   This change avoids use-after-free hazards.

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120143)
-  Fix race condition leading to use-after-free in shared tupledesc management (Robert Haas) [&sect;](https://postgr.es/c/ba00c3576)

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-116304)
-  Fix use-after-free in `pg_show_all_settings()` (Masahiko Sawada) [&sect;](https://postgr.es/c/15b78dc39)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120152)
-  Fix use-after-free in expanded-record domain constraint checking (Tom Lane) [&sect;](https://postgr.es/c/4614c78a8)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120150)
-  Prevent memory corruption after a large allocation in a “generation” memory context is freed (Heikki Linnakangas) [&sect;](https://postgr.es/c/735b0dc1a)

   The PostgreSQL Project thanks Jayden Sarveshkumar for reporting this problem. (CVE-2026-116328)
-  Fix buffer overrun in `tsvector_delete()` (Anthropic OSS program) [&sect;](https://postgr.es/c/834255ca8)

   With crafted input, `tsvector_delete()` could be induced to write more bytes than it had allocated for its result.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116315)
-  Fix buffer overrun in `ts_rank_cd()` (Anthropic OSS program) [&sect;](https://postgr.es/c/2d7573b00)

   With crafted input, `ts_rank_cd()` could be induced to write more bytes than it had allocated for its result.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116317)
-  Prevent construction of invalid `name` values via element assignment (Tom Lane) [&sect;](https://postgr.es/c/c673c0a61)

   `array_set_element()` could be abused to make the last byte of a `name`-type value nonzero. Subsequent uses of the value would then read past the end, resulting in crashes or disclosure of server memory. While at it, disallow construction of invalidly-encoded strings through this mechanism.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116321)
-  Treat application of a leaky function to a `CaseTestExpr` or internal `Param` node as leaky (Tom Lane) [&sect;](https://postgr.es/c/fad845598)

   A `CaseTestExpr` or internal `Param` represents the value of an expression somewhere else in the query. If that expression contains a `Var` then there is potential for information leakage. Therefore, reject such cases when deciding whether an expression is leak-proof.

   The PostgreSQL Project thanks Only see myself for reporting this problem. (CVE-2026-116327)
-  Don't generate an index-based plan for an aggregate whose input and output types don't match (Tom Lane) [&sect;](https://postgr.es/c/63b204d9f)

   Previously, we'd blindly do this if the aggregate has a `SORTOP` option attached. The index-based plan will return the aggregate's input column, leading to type violations and potentially to arbitrary code execution if its declared output type is different. In the master branch, disallow such aggregate declarations. In released branches, add a plan-time check to not generate such plans if the types don't match.

   The PostgreSQL Project thanks Zheng Yu, Claude and Ada Logics for reporting this problem. (CVE-2026-116331)
-  Avoid inserting the wrong tuple into a trigger transition table in a `MERGE` that does an `INSERT` followed by an `UPDATE` (Heikki Linnakangas, Jacek Bylina, Jacob Robertson, Nigel Gott) [&sect;](https://postgr.es/c/e48446c3e)

   The PostgreSQL Project thanks Jacek Bylina, Timothy Ball, Claude and Ada Logics for reporting this problem. (CVE-2026-116312)
-  Check that commutator and negator operator pairs return the same data type (Tom Lane) [&sect;](https://postgr.es/c/85420b358)

   Previously, substitution of a commutator or negator operator could lead to a type violation. This patch also cleans up some cases where one operator had a commutator/negator link to another, but the second operator didn't link back to the first.

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-116335)
-  Prevent functions in index predicates and expressions from breaking `REINDEX CONCURRENTLY` by manipulating the indexes (Michael Paquier) [&sect;](https://postgr.es/c/2362bc619)

   The PostgreSQL Project thanks Florent Saudel for reporting this problem. (CVE-2026-116336)
-  Prevent corrupting a newly-installed extension by renaming its target schema during `CREATE EXTENSION` or by installing it into a temporary schema (Noah Misch) [&sect;](https://postgr.es/c/1dc59af89) [&sect;](https://postgr.es/c/60341198b)

   With some ingenuity, an attacker could substitute objects of his creation for objects in the extension.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-6467)
-  Reject more unsafe names in extension script substitutions (Nathan Bossart) [&sect;](https://postgr.es/c/b0e85b99b)

   We already rejected schema and owner names containing quotes, backslashes, and dollar signs, but names containing newlines, carriage returns, comment delimiters, or the substitution tokens themselves can still achieve SQL injection. Reject all those cases too.

   The PostgreSQL Project thanks the Anthropic OSS program, Sven Klemm, and Aaron Chen for reporting this problem. (CVE-2026-120153)
-  Check `EXECUTE` privilege before calling cast functions within `to_json()` and other conversion-to-`JSON` code (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/ab7594cd1)

   Typically such functions would have public `EXECUTE` privilege anyway, but this puts the behavior on par with normal execution of casts.

   The PostgreSQL Project thanks Filip Janus and the Anthropic OSS program for reporting this problem. (CVE-2026-116303)
-  Check `EXECUTE` privilege before calling cast functions within foreign key enforcement checks (Nathan Bossart) [&sect;](https://postgr.es/c/3409adfbe)

   Typically such functions would have public `EXECUTE` privilege anyway, but this puts the behavior on par with normal execution of casts.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-116303)
-  Prevent `DROP OWNED BY` from changing row-level security policies on tables that the current user does not own (Zsolt Parragi, Anthropic OSS program) [&sect;](https://postgr.es/c/d90159929)

   Previously, an unprivileged role could remove itself from restrictive policies on tables it does not own. Instead issue a warning, as is already done for privileges that the current user cannot revoke.

   The PostgreSQL Project thanks Zsolt Parragi and the Anthropic OSS program for reporting this problem. (CVE-2026-116311)
-  Disallow security-restricted operations from creating types in temporary schemas (Anthropic OSS program, Nathan Bossart) [&sect;](https://postgr.es/c/0ed462c1f)

   We already disallowed creating temporary tables in this context, but a type also presents a hazard of shadowing permanent objects and thus allowing privilege escalation.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116316)
-  Avoid out-of-bounds accesses in FDW batch inserts (Nathan Bossart) [&sect;](https://postgr.es/c/b7c0062c1)

   The PostgreSQL Project thanks Sumedh Pathak and Sven Klemm for reporting this problem. (CVE-2026-120149)
-  Refuse to create extended statistics on expressions of transient types (Masahiko Sawada) [&sect;](https://postgr.es/c/30f67bdac)

   Disallow `CREATE STATISTICS` on expressions whose types would not be allowed as table column types, most notably anonymous record types. This prevents other sessions from misinterpreting the stored values.

   The PostgreSQL Project thanks Mehmet Ince for reporting this problem. (CVE-2026-120148)
-  Reject system catalogs as logical replication targets (Masahiko Sawada) [&sect;](https://postgr.es/c/71bb0b587)

   The PostgreSQL Project thanks Qian Zhang for reporting this problem. (CVE-2026-120145)
-  Check `ON CONFLICT` target row against RLS policies before applying user quals (Peter Geoghegan, Zsolt Parragi) [&sect;](https://postgr.es/c/9ee0723c0)

   Previously, leaky qual expressions could see rows that RLS policies should forbid them from seeing.

   The PostgreSQL Project thanks Peter Geoghegan, Zsolt Parragi, Only see myself, the Anthropic OSS program, Joao Detomini, Josef Šimánek, and Qing Xu for reporting this problem. (CVE-2026-120154)
-  Validate all connection options in replication subscriptions and `contrib/dblink` (Jacob Champion) [&sect;](https://postgr.es/c/f7b12b9f7)

   Previously, these checks were applied inconsistently.

   The PostgreSQL Project thanks Andrey Krasichkov, Bashir Altawil, and Zsolt Parragi for reporting this problem. (CVE-2026-14661)
-  Disallow declaring SQL functions as returning type `"any"` (Peter Geoghegan) [&sect;](https://postgr.es/c/4ac399b6e)

   This allowed type violations.

   The PostgreSQL Project thanks Peter Geoghegan, Claude and Ada Logics for reporting this problem. (CVE-2026-18447)
-  Harden more memory-size calculations against integer overflow (Álvaro Herrera) [&sect;](https://postgr.es/c/aa2d0aefe) [&sect;](https://postgr.es/c/c152f62ac)

   The PostgreSQL Project thanks Hillai Ben Sasson, Jayden Sarveshkumar, and Mehmet Ince for reporting this problem. (CVE-2026-120139)
-  Fix type confusion in set-returning functions that return `RECORD`, when materialized in a query's target list (Melanie Plageman, Robert Haas) [&sect;](https://postgr.es/c/19cff5432)

   The same tuple descriptor was used across multiple calls, which is unsafe since the actual record type could vary.

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120141)
-  Cope with concurrent rowtype changes in PL/pgSQL record field assignment (Tom Lane) [&sect;](https://postgr.es/c/cac94a040)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120151)
-  Support “tied” Perl objects in PL/Perl and associated contrib modules (Andrey Rachitskiy) [&sect;](https://postgr.es/c/26a89a81a)

   Up to now, if you passed a tied object to a PL/Perl operation, it would most likely get treated as `undef`, since we failed to perform Perl get magic on values we're passed. Tied hashes worked partially, in that you could read out their keys, but the values always read as `undef`.

   Originally this was intended as a master-only change, since it looks like a feature addition. But afterwards we discovered that assorted security-grade misbehaviors could be triggered by tied objects even without this patch. Rather than trying to secure two sets of behaviors, it seems best to back-patch this change first. (CVE-2026-116326)
-  Fix PL/Perl to cope with errors thrown from tied objects (Tom Lane, Andrew Dunstan) [&sect;](https://postgr.es/c/ab70d013e) [&sect;](https://postgr.es/c/8d6ca4358)

   The PostgreSQL Project thanks Ngoc Tran, the Anthropic OSS program, and Dwij Mehta for reporting this problem. (CVE-2026-116326)
-  Fix type confusion in PL/Perl's `get_perl_array_ref()` (Andrew Dunstan) [&sect;](https://postgr.es/c/25baebeb2)

   The PostgreSQL Project thanks Zheng Yu for reporting this problem. (CVE-2026-116309)
-  Fix use-after-free in scans of Perl arrays and hashes (Tom Lane, Anthropic OSS program) [&sect;](https://postgr.es/c/adbbabe18)

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116320)
-  Fix use-after-free in `plperl_return_next_internal()` (Anthropic OSS program) [&sect;](https://postgr.es/c/a78974d34)

   The PostgreSQL Project thanks the Anthropic OSS program and Ngoc Tran for reporting this problem. (CVE-2026-116319)
-  In PL/Perl, fix use-after-free triggered by releasing a prepared query that is being executed (Masahiko Sawada) [&sect;](https://postgr.es/c/b4ebd7240)

   The PostgreSQL Project thanks Zheng Yu and the Anthropic OSS program for reporting this problem. (CVE-2026-116307)
-  Fix use-after-free hazards in re-entrant use of PL/Tcl (Tom Lane) [&sect;](https://postgr.es/c/05114072c)

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116323)
-  Reject re-entrant use of `contrib/dblink` (Masahiko Sawada) [&sect;](https://postgr.es/c/f0aa15186)

   The PostgreSQL Project thanks Jiyong Ahn for reporting this problem. (CVE-2026-120147)
-  In `contrib/dict_int`, return NULL for zero-length tokens (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/dec3a342d)

   This behavior is consistent with other text-search dictionary types, and it avoids disclosure of server memory as was possible previously.

   The PostgreSQL Project thanks Nathan Bossart and the Anthropic OSS program for reporting this problem. (CVE-2026-116308)
-  Make the `\[un]restrict` commands functional in pg_dump's tar-format output (Noah Misch) [&sect;](https://postgr.es/c/40a81579e)

   This provides the intended level of protection to users who make use of the undocumented `restore.sql` tar member.

   The PostgreSQL Project thanks Pavel Kopylov for reporting this problem. (CVE-2026-116330)
-  Harden pg_restore against corrupt archive files (Dilip Kumar, Yonghwa Lee, Nathan Bossart) [&sect;](https://postgr.es/c/cc13d5e0c)

   This patch fixes out-of-bounds accesses and file path traversal hazards. There remain other issues that don't seem to amount to security bugs.

   The PostgreSQL Project thanks Dilip Kumar, Yonghwa Lee, OpenSec, and Ben Dixon for reporting this problem. (CVE-2026-116314)
-  Fix use of externally-controlled format string in pg_basebackup (Peter Eisentraut) [&sect;](https://postgr.es/c/885dfee08)

   This problem could only be reached if files with unusual names exist in the server's data directory.

   The PostgreSQL Project thanks Peter Eisentraut for reporting this problem. (CVE-2026-16240)
-  Disable btree index de-duplication support for the `bpchar` data type (Shihao Zhong, Peter Geoghegan) [&sect;](https://postgr.es/c/80c32ab93)

   Index de-duplication is unsafe for `bpchar` because its notion of equality is weaker than bitwise equality. In particular, because `bpchar` equality ignores trailing spaces, values of different lengths can compare equal. De-duplication then allows index-only scans to return subtly wrong results.

   This patch disables de-duplication for new index entries, but does nothing about possible corruption of existing entries. Therefore, users who have btree indexes on `bpchar` columns should `REINDEX` them after installing this update.

   Note that this problem only applies to columns declared as `bpchar`, that is, lacking any length constraint. The more common usage of <code>CHAR(</code><em>n</em><code>)</code> is not at risk, even though it shares the same underlying code, because the length constraint forces otherwise-equal strings to have the same length and thus to be bitwise equal.
-  Ensure that `VACUUM` cleans up all dead tuples in a GIN index (Peter Geoghegan) [&sect;](https://postgr.es/c/fb7c2a514) [&sect;](https://postgr.es/c/dd1634ddb)

   When a single `VACUUM` operation requires multiple index scans (because it is removing more tuples than it can hold in memory at once), it was possible to miss removing some index entries that should be removed. An additional bug involving a race with leaf page splits had the same effect. These bugs would eventually lead to a corrupt index with entries pointing at unrelated table rows. `REINDEX` is recommended if you suspect this has happened.
-  Fix GIN posting tree page deletion after an incomplete page split (Peter Geoghegan) [&sect;](https://postgr.es/c/4dd227f51)

   The deletion step could incorrectly disconnect sibling pages, causing subsequent index searches to miss rows. `REINDEX` is recommended if you suspect this has happened.
-  Fix failure to update mutable `pg_index` fields for TOAST-table indexes (Bharath Rupireddy) [&sect;](https://postgr.es/c/17c5a0659)

   `REINDEX CONCURRENTLY` can change these fields, but the relation cache was not expecting that, allowing index updates to be missed. The end result could be errors about “missing chunk number” due to missing index entries. Fortunately, these only represent index corruption; the TOAST table itself is fine, so a `REINDEX` will suffice to recover.
-  Perform join removal by editing the query's parse tree, rather than trying to fix up the planner's derived data structures (Tom Lane, Richard Guo, Shihao Zhong, Ayush Tiwari) [&sect;](https://postgr.es/c/986870baa) [&sect;](https://postgr.es/c/38fe5a936) [&sect;](https://postgr.es/c/d15ab66f5) [&sect;](https://postgr.es/c/fc8ed6362) [&sect;](https://postgr.es/c/54ac06f31) [&sect;](https://postgr.es/c/8ea97ee34) [&sect;](https://postgr.es/c/9df23331f) [&sect;](https://postgr.es/c/533915857)

   The original implementation of join removal in the planner has had numerous bugs over the years, including some that were never fixed. It was inherently fragile because changes in the planner's derived data structures often required adjustments in join removal, yet there was no easy way to keep that code in sync. Replace it with code that modifies the far-simpler parse tree representation and then loops around to re-generate all of the derived data. This is a significant change that normally we'd never consider back-patching, but the prospect of continuing bugs from the old implementation has driven us to take the risk of doing so.
-  Fix assorted planner bugs triggered by lateral references (Richard Guo) [&sect;](https://postgr.es/c/e2a7f784c) [&sect;](https://postgr.es/c/aad7571d0) [&sect;](https://postgr.es/c/7fb765945) [&sect;](https://postgr.es/c/b03eb54b1) [&sect;](https://postgr.es/c/333d566c4) [&sect;](https://postgr.es/c/e89049cba) [&sect;](https://postgr.es/c/cdc21ba1b) [&sect;](https://postgr.es/c/8e9288030) [&sect;](https://postgr.es/c/a8b9a7fe4) [&sect;](https://postgr.es/c/d87257b2c)

   These mistakes could lead to incorrect query plans causing wrong answers, or to internal planner errors such as “failed to build any N-way joins”.
-  Ensure proper placement of qual clauses containing PlaceHolderVars by assigning a relation ID to antijoins made from SubLinks (Richard Guo) [&sect;](https://postgr.es/c/ca4273866)

   This avoids edge cases where such a qual could be pushed into the right-hand side of the antijoin, giving it the opposite of the desired semantics.
-  Fix placement of PlaceHolderVars as nestloop parameters (Richard Guo) [&sect;](https://postgr.es/c/071fb9af2)

   In rare cases, planning could fail with “failed to assign all NestLoopParams to plan nodes”, “plan should not reference subplan's variable”, or “variable not found in subplan target list” errors.
-  Avoid planner failure for sort keys that are PlaceHolderVars (Richard Guo) [&sect;](https://postgr.es/c/d9452ecf5) [&sect;](https://postgr.es/c/617555d74)

   Depending on the plan shape selected, such cases could produce wrong results, or could fail with “variable not found in subplan target lists”.
-  Fix incorrect dropping of outer join clauses after re-associating outer joins (Richard Guo, Nitin Motiani) [&sect;](https://postgr.es/c/78bc3b96c)
-  Avoid “unsupported join alias expression” failures when a whole-row Var references a join and appears below an outer join (Richard Guo) [&sect;](https://postgr.es/c/f73794bc1)
-  Handle SubLinks correctly when expanding a whole-row join alias reference (Richard Guo) [&sect;](https://postgr.es/c/6bd71489d)

   This oversight led to errors like “cannot handle unplanned sub-select” or “unrecognized node type”.
-  Don't assume that `DISTINCT ON` ensures uniqueness of a subquery's output rows when the subquery's targetlist contains any set-returning functions (Richard Guo) [&sect;](https://postgr.es/c/8fc5b57ae)

   This doesn't follow, because the set-returning functions may be evaluated after the de-duplication step.
-  Fix incorrect constant-folding of a `JSON` constructor with a `RETURNING` clause inside a simple-form `CASE` with a constant test expression (Richard Guo) [&sect;](https://postgr.es/c/27ef9097b)
-  Fix incorrect behavior of `MIN()`/`MAX()` index scans when query specifies `FETCH ... WITH TIES` (Chaitanya Choudhary) [&sect;](https://postgr.es/c/a91d664d5)

   The `WITH TIES` option was accidentally copied into the sub-SELECT generated to perform the index scan, possibly triggering “more than one row returned by a subquery used as an expression” errors.
-  Use the correct collation when unique-ifying a semijoin's righthand side (Andrey Rachitskiy) [&sect;](https://postgr.es/c/a9c9c4f76)

   This is an issue only if the join's comparison clause uses a non-deterministic collation, since all deterministic collations agree on string equality. But if it does, it was possible to produce incorrect results with too many rows.
-  Acquire a predicate lock on conflicting rows during `INSERT ... ON CONFLICT` (Zsolt Parragi, Andrey Borodin) [&sect;](https://postgr.es/c/80c9a58b4) [&sect;](https://postgr.es/c/b76127a88)

   Failure to do this resulted in incorrect behavior in serializable mode, allowing concurrent updates that should have failed with a serialization error.
-  Fix incorrect multi-column `RANGE` partition pruning (David Rowley) [&sect;](https://postgr.es/c/a0396c8e3)

   When performing partition pruning with a `RANGE` partitioned table where pruning quals are only present for a leading prefix of the partition key, it was possible for partition pruning to prune away some partitions that shouldn't be pruned (leading to missing some rows that should be found) and/or include some partitions that were not needed.
-  Fix incorrect pruning of `DEFAULT` partitions in range-partitioned tables that use `MINVALUE` or `MAXVALUE` bounds (Ewan Young) [&sect;](https://postgr.es/c/1d8a93336)

   In some cases the code would incorrectly conclude that the `DEFAULT` partition needn't be scanned, causing failure to find some rows that should match the query.
-  Fix problems with queries applying multiple sort keys to a single column (David Rowley, Ayush Tiwari) [&sect;](https://postgr.es/c/0500fb3ab) [&sect;](https://postgr.es/c/1d57f6ef7)

   Some queries could crash the backend, while in other cases the additional sort specifications were ignored. Ordinarily such additional sort specifications are redundant anyway, since the first one would fully determine the result ordering. However, this is not necessarily so with non-deterministic collations, or with queries using different `ORDER BY` operators.
-  Avoid possible re-use of stale hash tables when a hashed SubPlan is referenced in multiple places (Rahul Yadav) [&sect;](https://postgr.es/c/06e46b651) [&sect;](https://postgr.es/c/1e9c734c2)
-  Ensure that Memoize plan nodes use the correct input collation (Shihao Zhong) [&sect;](https://postgr.es/c/083d60e34)

   This matters only when using non-deterministic collations.
-  Fix wrong results from Memoize plans when a Param is a subexpression of a cache key (Richard Guo) [&sect;](https://postgr.es/c/b5ad785be) [&sect;](https://postgr.es/c/f1d27a922)
-  Fix loss of parallel worker statistics for Memoize plan nodes underneath a Gather (Tatsuya Kawata) [&sect;](https://postgr.es/c/0896fb027)

   Only the statistics since the last rescan of the Gather were reported by `EXPLAIN ANALYZE`.
-  Don't create a shell type if a function definition specifies an unknown return type name with array or type modifier syntax (Heikki Linnakangas) [&sect;](https://postgr.es/c/e5e841cc5)

   Such cases are surely errors, since you should create the type before doing anything advanced with it.
-  Fix snapshot export and import to work properly on a snapshot taken during recovery (Peter Geoghegan, Bertrand Drouvot) [&sect;](https://postgr.es/c/d03ff3b75)

   The data layout is different in such a snapshot, but the export/import logic didn't know that, so that the imported snapshot behaved incorrectly. This led not only to wrong answers within the importing transaction, but possibly to persistent data corruption due to it setting hint bits incorrectly.

   This bug is known to be triggered by parallel pg_dump, when run on a standby server. In that context it could corrupt both the dump output and the standby server's data.
-  Fix atomicity problem in transaction snapshot import (Chee Wooson) [&sect;](https://postgr.es/c/7a392b0f6)

   There was a window in which onlookers might not see the source transaction as still running, possibly allowing wrong decisions by `VACUUM` for example.
-  Fix GIN index searches for plain scalar strings in strict `jsonpath` queries (Rui Zhao) [&sect;](https://postgr.es/c/0cf2c7e38)

   A `jsonb_ops` index scan for a `jsonpath` such as `strict $ == "scalar"` failed to find table entries that were simply `"scalar"`.
-  Fix errors in extraction of GIN index scan conditions from jsonpath queries (Tom Lane) [&sect;](https://postgr.es/c/c138ab8a9)

   An `AND` operator underneath a `NOT` operator could be misimplemented. Also a method applied to a literal constant, for example `"42".double()`, was not handled correctly. These mistakes did not cause visible errors, but the GIN index search would miss some or all of the rows that should match.
-  Fix edge-case error in SP-GiST quad-tree `<@` searches (John Naylor) [&sect;](https://postgr.es/c/9cba87fe9)

   Due to roundoff error, the search might sometimes not descend into a quadrant that it needs to search, leading to incomplete query results.
-  Fix edge-case failures in SP-GiST quad-tree indexes (Ayush Tiwari) [&sect;](https://postgr.es/c/235e15ec7)

   Roundoff errors could allow searches and insertions to fail with “getQuadrant: impossible case” errors.
-  Accept empty `tsquery` values in binary input (Ayush Tiwari) [&sect;](https://postgr.es/c/774c47de1)

   An empty `tsquery` value is allowed elsewhere, but `tsqueryrecv()` failed.
-  Disallow zero-length lexemes and out-of-range phrase distance values in binary input of `tsquery` values (Tom Lane) [&sect;](https://postgr.es/c/cdffe3375)

   These cases were already rejected during text input, but `tsqueryrecv()` omitted making equivalent checks. That allowed accepting values that would fail subsequent dump/reload.
-  Treat operators-nested-too-deeply errors as soft errors in `tsquery` input and `ltxtquery` input (Tom Lane) [&sect;](https://postgr.es/c/d4fae2993)
-  Consistently enforce maximum lengths in `tsvector` and `tsquery` (Tom Lane) [&sect;](https://postgr.es/c/1fc1b82b2)

   Some places thought the maximum lexeme length was `MAXSTRLEN` (2047), while others thought it was `MAXSTRLEN-1`. Fortunately, the looser bound is correct, so make them all use `MAXSTRLEN`. Likewise, some places were stricter than others about the total data length limit, `MAXSTRPOS`. These fixes avoid situations where a validly-constructed value would fail dump/reload.
-  Fix `right()` with a length argument of `INT_MIN` (-2147483648) (Ewan Young) [&sect;](https://postgr.es/c/48af47dcd)

   Due to integer overflow, the function returned the whole string when it should return an empty string.
-  Fix `pg_blocking_pids()` to account correctly for relation extension locks (Bertrand Drouvot) [&sect;](https://postgr.es/c/af42a7a82)
-  Fix 32-vs-64-bit confusion in `pg_indexam_progress_phasename()` (Markos Fountoulakis) [&sect;](https://postgr.es/c/6f126b328)

   This function accidentally worked on common hardware, but could misbehave on 32-bit or big-endian platforms.
-  Detect integer overflow in `to_char()` with `V` format (Andrey Rachitskiy) [&sect;](https://postgr.es/c/55af98a45)

   If the data value is an integer, and multiplying it by the power of ten implied by the `V` format causes integer overflow, `to_char()` silently emitted incorrect digits. Throw an integer overflow error instead.
-  Avoid integer overflow while calculating BRIN minmax_multi distances for `tid` values (Tomas Vondra) [&sect;](https://postgr.es/c/7e0d03a91)

   Very large `tid` values could result in negative distances, which would end in making the index inefficient, though not actually corrupt.
-  Avoid overflow while calculating BRIN minmax_multi distances for numeric values, and handle NaNs properly (Kirill Reshke, Tomas Vondra) [&sect;](https://postgr.es/c/444b85032)

   These oversights could result in errors being thrown (preventing index summarization), or in inefficient indexes.
-  Avoid overflow when computing differences for numeric ranges (Kirill Reshke) [&sect;](https://postgr.es/c/5c78685fc)

   In this context we want to return plus or minus infinity rather than throwing an overflow error.
-  Support out-of-order `FLAG` lines in Hunspell affix files (Ewan Young) [&sect;](https://postgr.es/c/5827ddfc2)

   Hunspell affix files usually put the `FLAG` line first, and our implementation silently depended on that. If it came after some lines whose meaning depends on the `FLAG` setting, the result would be crashes or failures to match compound words. Fix the logic so that the `FLAG` line can be anywhere.
-  Fix assorted errors in rule deparsing (Henson Choi) [&sect;](https://postgr.es/c/f0595c3d8) [&sect;](https://postgr.es/c/b27d566f0) [&sect;](https://postgr.es/c/715bf0305) [&sect;](https://postgr.es/c/a5e96ebe7)

   Always print `AS` before a target-relation alias in `UPDATE`, `DELETE`, or `MERGE`, to avoid ambiguity if that alias is `set`. Also ensure that <em>xml</em><code> IS DOCUMENT</code> constructs and subquery left-hand arguments are adequately parenthesized. Also fix deparsing of `AT TIME ZONE` and similar constructs in a view's `FROM` clause; these need to be wrapped in a no-op `CAST`. These errors could prevent dump restoration due to a syntax error in a view, or cause reloaded views to be parsed as something different from before.
-  Fix parser failure when a multi-row `INSERT ... VALUES` has a sub-SELECT in a target column subscript (Amit Langote) [&sect;](https://postgr.es/c/2dd680013)

   This led to “unexpected non-SELECT command in SubLink”.
-  Fix missed preprocessing of `WHERE` clauses in `COPY FROM` (Nathan Bossart) [&sect;](https://postgr.es/c/1955fd0b4)

   This oversight led to “cache lookup failed for function 0” errors with negated operators, and probably in other cases as well.
-  Fix race conditions with `DROP TABLESPACE` (Ayush Tiwari) [&sect;](https://postgr.es/c/749dd8d1f)

   A `DROP` executed concurrently with commands that attempt to add new objects to the tablespace could end with inconsistent catalog contents.
-  Cleanly reject operations on unlogged tables during recovery (Jim Jones) [&sect;](https://postgr.es/c/044887b60) [&sect;](https://postgr.es/c/391fded1e)

   An unlogged table will have no data on a standby, so reading it must fail. Most operations give an appropriate error for this, but `COPY ... TO` failed with an internal “could not open file” error, as did `currtid2()`.
-  Fix inconsistent permissions checking for role membership changes (ChangAo Chen) [&sect;](https://postgr.es/c/dbda203fe)

   This fixes cases where what should be a permissions failure was reported as an internal “no possible grantors” error.
-  Revert previous changes to free space map persistence (Alexander Korotkov) [&sect;](https://postgr.es/c/a8004e161)

   Commits c06d1a4ba et al. (included in the May 2026 releases) turned out to cause more trouble than they fixed.
-  Avoid backend hang during temporary table cleanup at exit, if the default transaction mode is `SERIALIZABLE READ ONLY DEFERRABLE` (Andrey Rachitskiy) [&sect;](https://postgr.es/c/a90cbbc86)

   Cleanup would block until concurrent serializable transactions complete, which could be indefinitely if some of them are prepared transactions.
-  Avoid crash when describing a `FETCH` statement whose cursor is gone (Dirkjan Bussink) [&sect;](https://postgr.es/c/caa19f80d)
-  If a vacuum delay sleep is requested while holding locks, don't sleep (Kevin Rocker, Andrey Borodin, Mostafa) [&sect;](https://postgr.es/c/226e86cda)

   Also, move several existing delay calls to points where it's actually safe to sleep.
-  Avoid freeing the plan for a foreign-key enforcement query if the plan is still in use (Lucas Jeffrey, Trakshan Mishra, Amit Langote) [&sect;](https://postgr.es/c/e37a8075f)

   It was possible to reach this behavior in contrived cases, such as a self-referencing ON DELETE CASCADE foreign key on tables with triggers that perform DML.
-  Fix WAL-logging of multi-level page deletion in btree indexes (Peter Geoghegan) [&sect;](https://postgr.es/c/4bd8a9613)

   The WAL log could contain stale values for a to-be-deleted page's sibling links, resulting in complaints on standby servers if `wal_consistency_checking` is enabled. Barring a crash, the damage would shortly be repaired by completion of the deletion action, so that this error was usually only cosmetic.
-  Ensure that state files for pending prepared transactions are written out atomically during a checkpoint (Shihao Zhong) [&sect;](https://postgr.es/c/11c87ab37)

   Previously it was possible for a crash to leave an invalid state file behind, causing subsequent database startup to fail.
-  Do not start a checkpoint in the middle of aborting a prepared transaction (Shihao Zhong) [&sect;](https://postgr.es/c/bb4eabd6a)

   A crash would allow the transaction to be committed afterwards, leading to inconsistent state.
-  Do not start a checkpoint in the middle of creating a multixact (Shihao Zhong) [&sect;](https://postgr.es/c/0d0bde0b0)

   This avoids possible data loss after a subsequent crash.
-  Fix possible loss of tuples updated by later-aborted subtransactions (Shihao Zhong, Michael Paquier) [&sect;](https://postgr.es/c/07d1b4418)

   A table rewrite (for example, `CLUSTER` or `VACUUM FULL`) could discard such tuples as invalid.
-  Do not lose tuple's `tid` when storing it into a tuple table slot (Virender Singla, Greg Burd) [&sect;](https://postgr.es/c/c5c36ae6f)

   Some code paths neglected to copy the `tid` field, which could lead to data corruption or wrong query answers in unusual cases.
-  Avoid race condition in marking a partitioned index valid (Andrey Borodin, Kirill Reshke) [&sect;](https://postgr.es/c/3bd5a4755)

   This could, for example, cause parallel pg_restore to fail due to trying to recreate foreign key constraints too soon.
-  Fix checks for an aborted updating transaction in multixids (Chee Wooson) [&sect;](https://postgr.es/c/7cafc4abb)

   Known consequences of this coding error include unexpected “new multixact has more than one updating member” errors, as well as incorrect failures to update a row.
-  A hash index scan that starts during a bucket split could miss some index entries (Peter Geoghegan, Amit Kapila) [&sect;](https://postgr.es/c/99ae176e4)
-  Fix incorrect position tracking in TID Range Scans (Ewan Young, David Rowley) [&sect;](https://postgr.es/c/6ceac2e12)

   Sequences involving fetching backwards and then forwards again could fail to scan all the way to the end of the range.
-  Avoid reading a potentially-torn copy of the `pg_control` file for base backups (David Steele) [&sect;](https://postgr.es/c/5dc3cfbb9) [&sect;](https://postgr.es/c/38ecdfdc1) [&sect;](https://postgr.es/c/0cfd56254)

   A concurrent write to the file could result in including inconsistent data in the backup.
-  Prevent making incorrect copies of non-relation files in base backups (David Steele) [&sect;](https://postgr.es/c/61aec3f3b)

   Files that routinely get replaced by atomic renames, such as `postgresql.auto.conf`, could be truncated or zero-padded in a backup made concurrently with such a replacement.
-  Ensure that base backups report checksum failures affecting single pages (Zsolt Parragi) [&sect;](https://postgr.es/c/45c88d621)

   Due to a logic error, checksum failures were counted in `pg_stat_database` only when more than one page in the same file was corrupt.
-  Verify the syntax of the `output_plugin_libraries` parameter when a value is assigned (Jacob Champion, Hayato Kuroda) [&sect;](https://postgr.es/c/079638030)

   Previously, a syntactically-incorrect value could be put into `postgresql.conf`, causing logical replication to fail.
-  Avoid incorrect update of backend status flags during replication slot invalidation (Vlad Lesin) [&sect;](https://postgr.es/c/02c32116d)

   An invalidation done by an auxiliary process such as the checkpointer could clobber the status field of an unrelated backend. The consequences of this are difficult to predict, but one possibility is that hot standby feedback could be effectively disabled.
-  In logical replication, handle subtransaction aborts better (Bingshuai Li, Hayato Kuroda, Zhijie Hou) [&sect;](https://postgr.es/c/ff5ba5f19)

   If a subtransaction that made a catalog update aborts, logical decoding mistakenly still believed that the update was effective, potentially causing incorrect decoding output.
-  Avoid race condition when building an initial decoding snapshot (Antonin Houska, Rui Zhao) [&sect;](https://postgr.es/c/a33ffaf7b)

   There was a small risk of data corruption from using a snapshot that includes transactions that have committed but not yet set their CLOG status.
-  Avoid null-pointer crash with missing replication option values in `pgoutput` logical replication sender (Zhijie Hou) [&sect;](https://postgr.es/c/e8191e13f)
-  Fix small memory leak in pgoutput with streamed prepared transactions (Ayush Tiwari) [&sect;](https://postgr.es/c/feb975917)
-  Fix lost wakeups in physical WAL senders (Alex Speller) [&sect;](https://postgr.es/c/2ac618d71)

   This fix corrects some cases where a WAL sender process could sleep even though there is data waiting to be sent.
-  Fix duplicate creation of temporary replication slot when a WAL receiver follows a timeline switch (ChangAo Chen) [&sect;](https://postgr.es/c/58fb12f4a)

   This mistake led to a “could not create replication slot” error followed by a restart of the WAL receiver, which would then recover, so there was no permanent damage.
-  Fix race condition between logical replication subscriber and concurrent index DDL (Mikhail Nikalayeu, Vignesh C) [&sect;](https://postgr.es/c/2389f898e)

   Changing the state of a replication target relation's primary key or replica identity index could cause the replication apply logic to miss finding a matching row (but only once per such event).
-  Fix incorrect LSNs in WAL decoding error messages (Rahul Yadav) [&sect;](https://postgr.es/c/bcb9d9051)

   These error messages reported the LSN of a WAL record that wasn't the one they were actually complaining about.
-  Don't allow `REPLICA IDENTITY FULL` lookups to select invalid indexes (Mikhail Nikalayeu) [&sect;](https://postgr.es/c/b3272b49a)

   Previously, logical replication could try to use an index left behind by a failed `CREATE INDEX CONCURRENTLY`, leading to wrong results since the index is likely incomplete.
-  Fix potential use-after-free cases in WAL sender (Ayush Tiwari) [&sect;](https://postgr.es/c/078110c00)
-  Avoid use-after-free in `jsonb_set_lax()` (John Naylor) [&sect;](https://postgr.es/c/4a3177299)
-  Fix double-free when `XMLSERIALIZE(... INDENT)` output exceeds 1GB (Jim Jones) [&sect;](https://postgr.es/c/582db3ee0)
-  Fix null-pointer dereferences that could occur when a schema is concurrently dropped (Baji Shaik) [&sect;](https://postgr.es/c/a2732c986)
-  Avoid null-pointer crash in logical subscription refresh when a relevant relation is concurrently dropped (Satyanarayana Narlapuram, Bharath Rupireddy) [&sect;](https://postgr.es/c/ec8aaf845)
-  Fix possible null-pointer dereference when building extended statistics on expressions (Shihao Zhong, Michael Paquier) [&sect;](https://postgr.es/c/0cea8f85e)

   This error was reachable only with unexpected (but legal) behavior of a custom `typanalyze` function.
-  Fix possibly-stale values of `pg_locks`.`waitstart` (Shihao Zhong) [&sect;](https://postgr.es/c/6f2f51084)
-  Harden pgstats code against partially-completed local pgstats entries (Michael Paquier) [&sect;](https://postgr.es/c/203b96953)
-  Fix cleanup of reused pgstats entries (Michael Paquier) [&sect;](https://postgr.es/c/ae97ef28d)

   This error led to leakage of pgstats entries, which could result in errors at backend exit or database shutdown.
-  Fix memory leaks caused by not releasing all data associated with a WAL reader (Tomás Senart) [&sect;](https://postgr.es/c/4c271dc21)

   In many cases this oversight did little harm because the containing memory context would soon be reset or deleted. However there are cases where it could cause long-lived memory bloat, notably in the checkpointer process.
-  Avoid memory bloat while computing `ndistinct` extended statistics across many columns (Ilia Evdokimov) [&sect;](https://postgr.es/c/050ad7edb)
-  Fix buffer pin leak after a failed hash index bucket split cleanup (Peter Geoghegan) [&sect;](https://postgr.es/c/c180f965d)

   This coding error resulted in nothing worse than warnings about resource leakage.
-  Ensure all relations opened for `AFTER` triggers are closed again (Hayato Kuroda, David Rowley) [&sect;](https://postgr.es/c/059aae612)

   This oversight triggered “resource was not closed” warnings, but nothing more serious than that.
-  Fix PL/pgSQL's handling of `anycompatiblemultirange` arguments and results (Jeevan Chalke) [&sect;](https://postgr.es/c/658e15092)

   This type had not been added to PL/pgSQL's knowledge of polymorphic types, resulting in errors when trying to create or use a function declared with it.
-  Harden PL/Perl and associated contrib modules against “tied” Perl hashes and arrays (Tom Lane) [&sect;](https://postgr.es/c/794473ec2) [&sect;](https://postgr.es/c/d54d3f5d9)

   Rely on `av_count()` to decide the length of a Perl array, rather than looking for a null pointer; and contrarily be more wary about unexpected null value pointers from array and hash fetches. These cases are only known to arise in connection with “tied” objects.
-  Avoid memory leakage in libpq with repeated GSS connections (Matt Carter, Tom Lane) [&sect;](https://postgr.es/c/72216a245)
-  Ensure that psql counts all the `COPY FROM STDIN` commands appearing in a multi-query command string (Zsolt Parragi) [&sect;](https://postgr.es/c/52f1eaae6)

   After the fix for CVE-2026-6464, psql would reject multiple `COPY FROM STDIN` commands appearing in a single string, because it was not expecting more than one `COPY` operation to occur.
-  Make psql cope better with out-of-memory conditions during backslash command processing (Fujii Masao) [&sect;](https://postgr.es/c/3d41f014f) [&sect;](https://postgr.es/c/1c62efd14)
-  Fix some small memory leaks in psql (Fujii Masao) [&sect;](https://postgr.es/c/4d2c7beb6)
-  Make pg_dump dump column statistics targets for indexes backing constraints (Zsolt Parragi) [&sect;](https://postgr.es/c/cb157339c)

   Formerly pg_dump would emit `ALTER INDEX ... SET STATISTICS` commands when needed for plain indexes, but not for indexes created via `EXCLUDE` constraints on expressions.
-  In pg_dump's tar-format output, don't put invalid `\restrict`/`\unrestrict` commands into the `restore.sql` tar member (Pavel Kopylov) [&sect;](https://postgr.es/c/c5fcb19d0)
-  Fix `contrib/file_fdw` to not crash in `ANALYZE` when the table has column default clauses (Fujii Masao) [&sect;](https://postgr.es/c/e9c45974a)
-  Fix possible index corruption in `contrib/pg_trgm` GiST indexes (Nate Clark) [&sect;](https://postgr.es/c/8ef710eb3)

   Combination of index entries could lose the flag identifying their format, so that subsequent index searches could miss matching rows.

   This fix does not repair existing indexes. Users should `REINDEX` existing `gist_trgm_ops` indexes after updating.
-  Fix error check in timer setup on Windows (Sehrope Sarkuni) [&sect;](https://postgr.es/c/908589bdb)

   Failure to create a thread to manage timeouts was not detected, leading to silent failure of subsequent timeout requests.
-  Use a non-locking initial test in TAS_SPIN on AArch64 (Nathan Bossart) [&sect;](https://postgr.es/c/d45816c24)

   This code change has been shown to improve performance on many-core AArch64 machines while not hurting performance on smaller ones; moreover, it avoids a performance regression with newer Linux kernels.
-  Update time zone data files to tzdata release 2026f (Tom Lane) [&sect;](https://postgr.es/c/f1762b51c) [&sect;](https://postgr.es/c/3868ed7b3) [&sect;](https://postgr.es/c/cfd4b54cf)

   Canada's Northwest Territories are on year-round UTC-06 (effectively, permanent DST) beginning in November 2026. This affects only America/Inuvik as the rest of the territory is covered by America/Edmonton, for which the equivalent change was released in 2026c. Similarly, Manitoba (America/Winnipeg) moves to year-round UTC-05. For now assume their TZ abbreviations will be CST and EST respectively.

   Cuba did not fall back on 2026-11-01; they will stay on DST this winter (but apparently plan to revert to normal next year).

   The backward-compatibility zone names EST5EDT, CST6CDT, MST7MDT, and PST8PDT now have simpler behavior for old time stamps. They use the indicated UTC offsets and TZ abbreviations indefinitely far back, rather than allowing aspects of the associated city-named zone's historical behavior to leak through. This conforms better to POSIX.

   Historical corrections for Brazil, Colombia, Iran, and Ireland.

   On Windows, initdb will now recognize the new time zone names that Microsoft has devised for the updated Canadian zones.
