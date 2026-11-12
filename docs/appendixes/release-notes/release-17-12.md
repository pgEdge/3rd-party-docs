<a id="release-17-12"></a>

## Release 17.12


**Release date:.**


2026-11-12


 This release contains a variety of fixes from 17.11. For information about new features in major release 17, see [Release 17](release-17.md#release-17).
 <a id="release-17-12-migration"></a>

### Migration to Version 17.12


 A dump/restore is not required for those running 17.X.


 However, several different bugs have been discovered that could result in index corruption. Examine the changelog entries below that mention `REINDEX`, and reindex any indexes you have that might be affected.


 Also, existing incremental backup files should not be trusted, thanks to the discovery of several ways that they could become corrupt. It's recommended that users of incremental backups take a new full backup after updating and use it as the basis for subsequent incremental backups.


 Also, if you are upgrading from a version earlier than 17.11, see [Release 17.11](release-17-11.md#release-17-11).
  <a id="release-17-12-changes"></a>

### Changes


-  Add defenses against problems caused by altering composite types concurrently with operations on data of those types (Heikki Linnakangas, Andres Freund, Nathan Bossart, Nikhil Sontakke) [&sect;](https://postgr.es/c/bfcb6625e) [&sect;](https://postgr.es/c/3e8bb75c8) [&sect;](https://postgr.es/c/11658b4df) [&sect;](https://postgr.es/c/fbee8e90f) [&sect;](https://postgr.es/c/7d089b28c) [&sect;](https://postgr.es/c/1e3b37da9) [&sect;](https://postgr.es/c/7e1aeaad3) [&sect;](https://postgr.es/c/b65b69f2c) [&sect;](https://postgr.es/c/23a1d8825) [&sect;](https://postgr.es/c/f234ae72b) [&sect;](https://postgr.es/c/b019147c9) [&sect;](https://postgr.es/c/d4f8c3e15) [&sect;](https://postgr.es/c/800e67351) [&sect;](https://postgr.es/c/83ec4accd) [&sect;](https://postgr.es/c/5da4d4b1b) [&sect;](https://postgr.es/c/d5cf40e02) [&sect;](https://postgr.es/c/98e9001cf) [&sect;](https://postgr.es/c/05283881a) [&sect;](https://postgr.es/c/d3d599625) [&sect;](https://postgr.es/c/f397e0ae3) [&sect;](https://postgr.es/c/3b2c73d0e) [&sect;](https://postgr.es/c/db59ee006) [&sect;](https://postgr.es/c/deebcc772) [&sect;](https://postgr.es/c/bcbcbd5a2) [&sect;](https://postgr.es/c/9f6302533) [&sect;](https://postgr.es/c/d0b52be9b) [&sect;](https://postgr.es/c/5f8d39154) [&sect;](https://postgr.es/c/3643ce425)

   Operations such as `ALTER TABLE ... ALTER COLUMN TYPE` are hazardous because tuples of the older row layout may still exist in the system. While we have long had some checks to defend against that, they were never really complete, and they entirely failed to detect cases where composite types were changed concurrently with queries using them. This voluminous set of patches attempts to improve that.

   Introduce versioning logic to permit detection of composite datums that don't match the tuple descriptor we wish to interpret them with. For full protection, this change requires rebuilding C-language extensions, and there are some extensions that will require source-code changes. If necessary, the new `composite_type_mismatch` server parameter can be used to disable these checks, but doing so carries security risks.

   Add locking to prevent cases such as `ALTER TYPE` on a composite type used in a `CREATE TABLE` that hasn't committed yet.

   Tighten the rules about whether `ALTER COLUMN TYPE` is allowed when dependencies exist on the composite type. For example, it's no longer allowed when a function's default expression or new-style SQL function body depends on the type.

   The PostgreSQL Project thanks Andrew Gierth, Alexander Lakhin, Francesco Verardi, Hcamael, Tomer Fichman, OpenAI Security Research Team, Peter Geoghegan, Mehmet Ince, Bryan Green, Sumedh Pathak, Sergey Shinderuk, Yuxiao Guo, Nikhil Sontakke, Daniel Bakker, the Anthropic OSS program, Claude and Ada Logics for reporting this problem. (CVE-2026-6460)
-  Cope with concurrent rowtype changes in `ExecEvalFieldStore[De]Form()` (Tomas Vondra, Heikki Linnakangas) [&sect;](https://postgr.es/c/02ca67429)

   The PostgreSQL Project thanks OpenAI Codex Security, Bronson Yen, Mehmet Ince, Claude and Ada Logics for reporting this problem. (CVE-2026-116306)
-  Enforce view row visibility in `INSERT ... ON CONFLICT` (Peter Geoghegan, Zsolt Parragi) [&sect;](https://postgr.es/c/a14aacebe)

   When applied to a view, `INSERT ... ON CONFLICT DO UPDATE` could update rows that should be hidden by the view's `WHERE` clause. Although it was documented, this behavior is surely unexpected, and it gives rise to security problems: a user can update rows he should have no privilege to reach. Updating is incorrect, but inserting is impossible too (it would fail due to the unique arbiter index), so throw an error instead when the case arises. `INSERT ... ON CONFLICT DO SELECT` could show such rows as well, but no longer does.

   The PostgreSQL Project thanks Qing Xu, Carl Andersen, Zsolt Parragi, the Anthropic OSS program, Shihao Zhong, and Josef Šimánek for reporting this problem. (CVE-2026-16237)
-  Align `REASSIGN OWNED`'s caller-privilege checks with `ALTER ... OWNER TO`'s checks (Zsolt Parragi) [&sect;](https://postgr.es/c/0b8d1768f)

   Since v16, giving an object to another role requires the ability to `SET ROLE` to that role. However, a role granted `INHERIT TRUE, SET FALSE` could bypass that via `REASSIGN OWNED`. A role granted `INHERIT FALSE, SET TRUE` saw the inverse problem: `REASSIGN OWNED` failed while `ALTER OWNER` succeeded. Make `REASSIGN` follow the `ALTER OWNER` rules.

   `REASSIGN OWNED` also bypassed the `CREATEDB` privilege check in `ALTER DATABASE OWNER`.

   In v15 and v14, `REASSIGN OWNED`'s privilege check is stricter than the membership checks in `ALTER OWNER`. Leave that alone, but still fix the missing `CREATEDB` check.

   The PostgreSQL Project thanks Zsolt Parragi for reporting this problem. (CVE-2026-120146)
-  Fix use-after-free in `COPY FROM ON_ERROR` reporting (Masahiko Sawada) [&sect;](https://postgr.es/c/3094cdb7d)

   An error report raised in `COPY FROM` could disclose server memory.

   The PostgreSQL Project thanks Jiyong Ahn, Claude and Ada Logics for reporting this problem. (CVE-2026-116310)
-  Disallow closing a portal that is still being set up (Michael Paquier) [&sect;](https://postgr.es/c/36f42216e)

   This change forestalls various use-after-free issues.

   The PostgreSQL Project thanks Tomas Vondra for reporting this problem. (CVE-2026-116332)
-  Forcibly close portals associated with transition tables at the end of the trigger function (Jeff Davis) [&sect;](https://postgr.es/c/695d72c57)

   This change avoids use-after-free hazards.

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120143)
-  Fix race condition leading to use-after-free in shared tupledesc management (Robert Haas) [&sect;](https://postgr.es/c/a4d40a6ae)

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-116304)
-  Fix use-after-free in `pg_show_all_settings()` (Masahiko Sawada) [&sect;](https://postgr.es/c/46c7529fa)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120152)
-  Fix use-after-free in expanded-record domain constraint checking (Tom Lane) [&sect;](https://postgr.es/c/f4d200615)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120150)
-  Prevent memory corruption after a large allocation in a “generation” memory context is freed (Heikki Linnakangas) [&sect;](https://postgr.es/c/f34ccffdd)

   The PostgreSQL Project thanks Jayden Sarveshkumar for reporting this problem. (CVE-2026-116328)
-  Fix buffer overrun in `tsvector_delete()` (Anthropic OSS program) [&sect;](https://postgr.es/c/cacc71935)

   With crafted input, `tsvector_delete()` could be induced to write more bytes than it had allocated for its result.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116315)
-  Fix buffer overrun in `ts_rank_cd()` (Anthropic OSS program) [&sect;](https://postgr.es/c/7f8424ece)

   With crafted input, `ts_rank_cd()` could be induced to write more bytes than it had allocated for its result.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116317)
-  Prevent construction of invalid `name` values via element assignment (Tom Lane) [&sect;](https://postgr.es/c/5788eaee3)

   `array_set_element()` could be abused to make the last byte of a `name`-type value nonzero. Subsequent uses of the value would then read past the end, resulting in crashes or disclosure of server memory. While at it, disallow construction of invalidly-encoded strings through this mechanism.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116321)
-  Treat application of a leaky function to a `CaseTestExpr` or internal `Param` node as leaky (Tom Lane) [&sect;](https://postgr.es/c/606c23540)

   A `CaseTestExpr` or internal `Param` represents the value of an expression somewhere else in the query. If that expression contains a `Var` then there is potential for information leakage. Therefore, reject such cases when deciding whether an expression is leak-proof.

   The PostgreSQL Project thanks Only see myself for reporting this problem. (CVE-2026-116327)
-  Don't generate an index-based plan for an aggregate whose input and output types don't match (Tom Lane) [&sect;](https://postgr.es/c/dd2cd3803)

   Previously, we'd blindly do this if the aggregate has a `SORTOP` option attached. The index-based plan will return the aggregate's input column, leading to type violations and potentially to arbitrary code execution if its declared output type is different. In the master branch, disallow such aggregate declarations. In released branches, add a plan-time check to not generate such plans if the types don't match.

   The PostgreSQL Project thanks Zheng Yu, Claude and Ada Logics for reporting this problem. (CVE-2026-116331)
-  Avoid inserting the wrong tuple into a trigger transition table in a `MERGE` that does an `INSERT` followed by an `UPDATE` (Heikki Linnakangas, Jacek Bylina, Jacob Robertson, Nigel Gott) [&sect;](https://postgr.es/c/a5a82b4d2)

   The PostgreSQL Project thanks Jacek Bylina, Timothy Ball, Claude and Ada Logics for reporting this problem. (CVE-2026-116312)
-  Check that commutator and negator operator pairs return the same data type (Tom Lane) [&sect;](https://postgr.es/c/22f838a6d)

   Previously, substitution of a commutator or negator operator could lead to a type violation. This patch also cleans up some cases where one operator had a commutator/negator link to another, but the second operator didn't link back to the first.

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-116335)
-  Prevent functions in index predicates and expressions from breaking `REINDEX CONCURRENTLY` by manipulating the indexes (Michael Paquier) [&sect;](https://postgr.es/c/23e7e2bde)

   The PostgreSQL Project thanks Florent Saudel for reporting this problem. (CVE-2026-116336)
-  Prevent corrupting a newly-installed extension by renaming its target schema during `CREATE EXTENSION` or by installing it into a temporary schema (Noah Misch) [&sect;](https://postgr.es/c/79f9793fc) [&sect;](https://postgr.es/c/5c7515337)

   With some ingenuity, an attacker could substitute objects of his creation for objects in the extension.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-6467)
-  Reject more unsafe names in extension script substitutions (Nathan Bossart) [&sect;](https://postgr.es/c/438b435f0)

   We already rejected schema and owner names containing quotes, backslashes, and dollar signs, but names containing newlines, carriage returns, comment delimiters, or the substitution tokens themselves can still achieve SQL injection. Reject all those cases too.

   The PostgreSQL Project thanks the Anthropic OSS program, Sven Klemm, and Aaron Chen for reporting this problem. (CVE-2026-120153)
-  Check `EXECUTE` privilege before calling cast functions within `to_json()` and other conversion-to-`JSON` code (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/9520dbea7)

   Typically such functions would have public `EXECUTE` privilege anyway, but this puts the behavior on par with normal execution of casts.

   The PostgreSQL Project thanks Filip Janus and the Anthropic OSS program for reporting this problem. (CVE-2026-116303)
-  Check `EXECUTE` privilege before calling cast functions within foreign key enforcement checks (Nathan Bossart) [&sect;](https://postgr.es/c/445e10c50)

   Typically such functions would have public `EXECUTE` privilege anyway, but this puts the behavior on par with normal execution of casts.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-116303)
-  Prevent `DROP OWNED BY` from changing row-level security policies on tables that the current user does not own (Zsolt Parragi, Anthropic OSS program) [&sect;](https://postgr.es/c/21c732aad)

   Previously, an unprivileged role could remove itself from restrictive policies on tables it does not own. Instead issue a warning, as is already done for privileges that the current user cannot revoke.

   The PostgreSQL Project thanks Zsolt Parragi and the Anthropic OSS program for reporting this problem. (CVE-2026-116311)
-  Disallow security-restricted operations from creating types in temporary schemas (Anthropic OSS program, Nathan Bossart) [&sect;](https://postgr.es/c/3f082dcc7)

   We already disallowed creating temporary tables in this context, but a type also presents a hazard of shadowing permanent objects and thus allowing privilege escalation.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116316)
-  Avoid out-of-bounds accesses in FDW batch inserts (Nathan Bossart) [&sect;](https://postgr.es/c/e0bd77667)

   The PostgreSQL Project thanks Sumedh Pathak and Sven Klemm for reporting this problem. (CVE-2026-120149)
-  Avoid out-of-bounds reads in parallel btree scans with array keys (Nathan Bossart) [&sect;](https://postgr.es/c/4cd18df09)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-116333)
-  Refuse to create extended statistics on expressions of transient types (Masahiko Sawada) [&sect;](https://postgr.es/c/7f51d62e1)

   Disallow `CREATE STATISTICS` on expressions whose types would not be allowed as table column types, most notably anonymous record types. This prevents other sessions from misinterpreting the stored values.

   The PostgreSQL Project thanks Mehmet Ince for reporting this problem. (CVE-2026-120148)
-  Reject system catalogs as logical replication targets (Masahiko Sawada) [&sect;](https://postgr.es/c/9532d4c94)

   The PostgreSQL Project thanks Qian Zhang for reporting this problem. (CVE-2026-120145)
-  Check `ON CONFLICT` target row against RLS policies before applying user quals (Peter Geoghegan, Zsolt Parragi) [&sect;](https://postgr.es/c/9c677791d)

   Previously, leaky qual expressions could see rows that RLS policies should forbid them from seeing.

   The PostgreSQL Project thanks Peter Geoghegan, Zsolt Parragi, Only see myself, the Anthropic OSS program, Joao Detomini, Josef Šimánek, and Qing Xu for reporting this problem. (CVE-2026-120154)
-  Validate all connection options in replication subscriptions and `contrib/dblink` (Jacob Champion) [&sect;](https://postgr.es/c/ea9b427fa)

   Previously, these checks were applied inconsistently.

   The PostgreSQL Project thanks Andrey Krasichkov, Bashir Altawil, and Zsolt Parragi for reporting this problem. (CVE-2026-14661)
-  Handle dropped columns safely while comparing row types (Peter Geoghegan) [&sect;](https://postgr.es/c/354d5c71b)

   Previously, two row types that had dropped columns at the same position could be considered equal even if those columns were of different physical widths. That leads to problems when trying to deconstruct tuples of those types.

   The PostgreSQL Project thanks Peter Geoghegan and Mehmet Ince for reporting this problem. (CVE-2026-120142)
-  Disallow declaring SQL functions as returning type `"any"` (Peter Geoghegan) [&sect;](https://postgr.es/c/5e99fab61)

   This allowed type violations.

   The PostgreSQL Project thanks Peter Geoghegan, Claude and Ada Logics for reporting this problem. (CVE-2026-18447)
-  Harden more memory-size calculations against integer overflow (Álvaro Herrera) [&sect;](https://postgr.es/c/bef80ad61) [&sect;](https://postgr.es/c/a17bc8b3c)

   The PostgreSQL Project thanks Hillai Ben Sasson, Jayden Sarveshkumar, and Mehmet Ince for reporting this problem. (CVE-2026-120139)
-  Fix type confusion in set-returning functions that return `RECORD`, when materialized in a query's target list (Melanie Plageman, Robert Haas) [&sect;](https://postgr.es/c/55c00bffe)

   The same tuple descriptor was used across multiple calls, which is unsafe since the actual record type could vary.

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120141)
-  Cope with concurrent rowtype changes in PL/pgSQL record field assignment (Tom Lane) [&sect;](https://postgr.es/c/3a3eb7616)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120151)
-  Support “tied” Perl objects in PL/Perl and associated contrib modules (Andrey Rachitskiy) [&sect;](https://postgr.es/c/217eb08b4)

   Up to now, if you passed a tied object to a PL/Perl operation, it would most likely get treated as `undef`, since we failed to perform Perl get magic on values we're passed. Tied hashes worked partially, in that you could read out their keys, but the values always read as `undef`.

   Originally this was intended as a master-only change, since it looks like a feature addition. But afterwards we discovered that assorted security-grade misbehaviors could be triggered by tied objects even without this patch. Rather than trying to secure two sets of behaviors, it seems best to back-patch this change first. (CVE-2026-116326)
-  Fix PL/Perl to cope with errors thrown from tied objects (Tom Lane, Andrew Dunstan) [&sect;](https://postgr.es/c/62a19e158) [&sect;](https://postgr.es/c/6e5524cda)

   The PostgreSQL Project thanks Ngoc Tran, the Anthropic OSS program, and Dwij Mehta for reporting this problem. (CVE-2026-116326)
-  Fix type confusion in PL/Perl's `get_perl_array_ref()` (Andrew Dunstan) [&sect;](https://postgr.es/c/dcc32e7d9)

   The PostgreSQL Project thanks Zheng Yu for reporting this problem. (CVE-2026-116309)
-  Fix use-after-free in scans of Perl arrays and hashes (Tom Lane, Anthropic OSS program) [&sect;](https://postgr.es/c/d1883981a)

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116320)
-  Fix use-after-free in `plperl_return_next_internal()` (Anthropic OSS program) [&sect;](https://postgr.es/c/eab800134)

   The PostgreSQL Project thanks the Anthropic OSS program and Ngoc Tran for reporting this problem. (CVE-2026-116319)
-  In PL/Perl, fix use-after-free triggered by releasing a prepared query that is being executed (Masahiko Sawada) [&sect;](https://postgr.es/c/2009539be)

   The PostgreSQL Project thanks Zheng Yu and the Anthropic OSS program for reporting this problem. (CVE-2026-116307)
-  Fix use-after-free hazards in re-entrant use of PL/Tcl (Tom Lane) [&sect;](https://postgr.es/c/9d15ca063)

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116323)
-  Reject re-entrant use of `contrib/dblink` (Masahiko Sawada) [&sect;](https://postgr.es/c/33f4232cb)

   The PostgreSQL Project thanks Jiyong Ahn for reporting this problem. (CVE-2026-120147)
-  In `contrib/dict_int`, return NULL for zero-length tokens (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/b24a45486)

   This behavior is consistent with other text-search dictionary types, and it avoids disclosure of server memory as was possible previously.

   The PostgreSQL Project thanks Nathan Bossart and the Anthropic OSS program for reporting this problem. (CVE-2026-116308)
-  Ensure that reindexdb correctly quotes index names it is dealing with (Anthropic OSS program) [&sect;](https://postgr.es/c/2680b0781)

   Crafted index names could yield a SQL-injection attack against the user running reindexdb.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116325)
-  Make the `\[un]restrict` commands functional in pg_dump's tar-format output (Noah Misch) [&sect;](https://postgr.es/c/4d3c13724)

   This provides the intended level of protection to users who make use of the undocumented `restore.sql` tar member.

   The PostgreSQL Project thanks Pavel Kopylov for reporting this problem. (CVE-2026-116330)
-  Harden pg_restore against corrupt archive files (Dilip Kumar, Yonghwa Lee, Nathan Bossart) [&sect;](https://postgr.es/c/a6dc99b28)

   This patch fixes out-of-bounds accesses and file path traversal hazards. There remain other issues that don't seem to amount to security bugs.

   The PostgreSQL Project thanks Dilip Kumar, Yonghwa Lee, OpenSec, and Ben Dixon for reporting this problem. (CVE-2026-116314)
-  Fix use of externally-controlled format string in pg_basebackup (Peter Eisentraut) [&sect;](https://postgr.es/c/04986ad59)

   This problem could only be reached if files with unusual names exist in the server's data directory.

   The PostgreSQL Project thanks Peter Eisentraut for reporting this problem. (CVE-2026-16240)
-  Disable btree index de-duplication support for the `bpchar` data type (Shihao Zhong, Peter Geoghegan) [&sect;](https://postgr.es/c/74f26a193)

   Index de-duplication is unsafe for `bpchar` because its notion of equality is weaker than bitwise equality. In particular, because `bpchar` equality ignores trailing spaces, values of different lengths can compare equal. De-duplication then allows index-only scans to return subtly wrong results.

   This patch disables de-duplication for new index entries, but does nothing about possible corruption of existing entries. Therefore, users who have btree indexes on `bpchar` columns should `REINDEX` them after installing this update.

   Note that this problem only applies to columns declared as `bpchar`, that is, lacking any length constraint. The more common usage of <code>CHAR(</code><em>n</em><code>)</code> is not at risk, even though it shares the same underlying code, because the length constraint forces otherwise-equal strings to have the same length and thus to be bitwise equal.
-  Ensure that `VACUUM` cleans up all dead tuples in a GIN index (Peter Geoghegan) [&sect;](https://postgr.es/c/d778f00c1) [&sect;](https://postgr.es/c/af7cbf4e2)

   When a single `VACUUM` operation requires multiple index scans (because it is removing more tuples than it can hold in memory at once), it was possible to miss removing some index entries that should be removed. An additional bug involving a race with leaf page splits had the same effect. These bugs would eventually lead to a corrupt index with entries pointing at unrelated table rows. `REINDEX` is recommended if you suspect this has happened.
-  Fix GIN posting tree page deletion after an incomplete page split (Peter Geoghegan) [&sect;](https://postgr.es/c/aa14281fd)

   The deletion step could incorrectly disconnect sibling pages, causing subsequent index searches to miss rows. `REINDEX` is recommended if you suspect this has happened.
-  Fix failure to update mutable `pg_index` fields for TOAST-table indexes (Bharath Rupireddy) [&sect;](https://postgr.es/c/b3852eec6)

   `REINDEX CONCURRENTLY` can change these fields, but the relation cache was not expecting that, allowing index updates to be missed. The end result could be errors about “missing chunk number” due to missing index entries. Fortunately, these only represent index corruption; the TOAST table itself is fine, so a `REINDEX` will suffice to recover.
-  Perform join removal by editing the query's parse tree, rather than trying to fix up the planner's derived data structures (Tom Lane, Richard Guo, Shihao Zhong, Ayush Tiwari) [&sect;](https://postgr.es/c/13466d1f7) [&sect;](https://postgr.es/c/1000f7b14) [&sect;](https://postgr.es/c/12f10c25c) [&sect;](https://postgr.es/c/197a41aa6) [&sect;](https://postgr.es/c/ccb7d9dba) [&sect;](https://postgr.es/c/51f9e30f3) [&sect;](https://postgr.es/c/bd003dca0) [&sect;](https://postgr.es/c/00d1b4598)

   The original implementation of join removal in the planner has had numerous bugs over the years, including some that were never fixed. It was inherently fragile because changes in the planner's derived data structures often required adjustments in join removal, yet there was no easy way to keep that code in sync. Replace it with code that modifies the far-simpler parse tree representation and then loops around to re-generate all of the derived data. This is a significant change that normally we'd never consider back-patching, but the prospect of continuing bugs from the old implementation has driven us to take the risk of doing so.
-  Fix assorted planner bugs triggered by lateral references (Richard Guo) [&sect;](https://postgr.es/c/70cf32506) [&sect;](https://postgr.es/c/b83b7ddf7) [&sect;](https://postgr.es/c/c37a2bb54) [&sect;](https://postgr.es/c/d9595e7e6) [&sect;](https://postgr.es/c/9e8f4ce28) [&sect;](https://postgr.es/c/cf24f8e61) [&sect;](https://postgr.es/c/55b9f54dc) [&sect;](https://postgr.es/c/65e944d92) [&sect;](https://postgr.es/c/3e914438a) [&sect;](https://postgr.es/c/94035972a)

   These mistakes could lead to incorrect query plans causing wrong answers, or to internal planner errors such as “failed to build any N-way joins”.
-  Ensure proper placement of qual clauses containing PlaceHolderVars by assigning a relation ID to antijoins made from SubLinks (Richard Guo) [&sect;](https://postgr.es/c/854f81f07)

   This avoids edge cases where such a qual could be pushed into the right-hand side of the antijoin, giving it the opposite of the desired semantics.
-  Fix placement of PlaceHolderVars as nestloop parameters (Richard Guo) [&sect;](https://postgr.es/c/f2e9e3aa1)

   In rare cases, planning could fail with “failed to assign all NestLoopParams to plan nodes”, “plan should not reference subplan's variable”, or “variable not found in subplan target list” errors.
-  Avoid planner failure for sort keys that are PlaceHolderVars (Richard Guo) [&sect;](https://postgr.es/c/05d073e2c) [&sect;](https://postgr.es/c/2b66c9c8a)

   Depending on the plan shape selected, such cases could produce wrong results, or could fail with “variable not found in subplan target lists”.
-  Fix incorrect dropping of outer join clauses after re-associating outer joins (Richard Guo, Nitin Motiani) [&sect;](https://postgr.es/c/ba7127f24)
-  Avoid “unsupported join alias expression” failures when a whole-row Var references a join and appears below an outer join (Richard Guo) [&sect;](https://postgr.es/c/0c4235013)
-  Handle SubLinks correctly when expanding a whole-row join alias reference (Richard Guo) [&sect;](https://postgr.es/c/6c0698339)

   This oversight led to errors like “cannot handle unplanned sub-select” or “unrecognized node type”.
-  Don't assume that `DISTINCT ON` ensures uniqueness of a subquery's output rows when the subquery's targetlist contains any set-returning functions (Richard Guo) [&sect;](https://postgr.es/c/a4eb938b3)

   This doesn't follow, because the set-returning functions may be evaluated after the de-duplication step.
-  Fix incorrect constant-folding of a `JSON` constructor with a `RETURNING` clause inside a simple-form `CASE` with a constant test expression (Richard Guo) [&sect;](https://postgr.es/c/23088673d)
-  Fix incorrect behavior of `MIN()`/`MAX()` index scans when query specifies `FETCH ... WITH TIES` (Chaitanya Choudhary) [&sect;](https://postgr.es/c/70d0ac3c9)

   The `WITH TIES` option was accidentally copied into the sub-SELECT generated to perform the index scan, possibly triggering “more than one row returned by a subquery used as an expression” errors.
-  Use the correct collation when unique-ifying a semijoin's righthand side (Andrey Rachitskiy) [&sect;](https://postgr.es/c/a10d14afc)

   This is an issue only if the join's comparison clause uses a non-deterministic collation, since all deterministic collations agree on string equality. But if it does, it was possible to produce incorrect results with too many rows.
-  Acquire a predicate lock on conflicting rows during `INSERT ... ON CONFLICT` (Zsolt Parragi, Andrey Borodin) [&sect;](https://postgr.es/c/5b4ebfbfa) [&sect;](https://postgr.es/c/0d83f9e0e)

   Failure to do this resulted in incorrect behavior in serializable mode, allowing concurrent updates that should have failed with a serialization error.
-  Fix incorrect multi-column `RANGE` partition pruning (David Rowley) [&sect;](https://postgr.es/c/2d8494e24)

   When performing partition pruning with a `RANGE` partitioned table where pruning quals are only present for a leading prefix of the partition key, it was possible for partition pruning to prune away some partitions that shouldn't be pruned (leading to missing some rows that should be found) and/or include some partitions that were not needed.
-  Fix incorrect pruning of `DEFAULT` partitions in range-partitioned tables that use `MINVALUE` or `MAXVALUE` bounds (Ewan Young) [&sect;](https://postgr.es/c/d21d4f1ca)

   In some cases the code would incorrectly conclude that the `DEFAULT` partition needn't be scanned, causing failure to find some rows that should match the query.
-  Fix problems with queries applying multiple sort keys to a single column (David Rowley, Ayush Tiwari) [&sect;](https://postgr.es/c/9171fca99) [&sect;](https://postgr.es/c/2a3ac1554)

   Some queries could crash the backend, while in other cases the additional sort specifications were ignored. Ordinarily such additional sort specifications are redundant anyway, since the first one would fully determine the result ordering. However, this is not necessarily so with non-deterministic collations, or with queries using different `ORDER BY` operators.
-  Avoid possible re-use of stale hash tables when a hashed SubPlan is referenced in multiple places (Rahul Yadav) [&sect;](https://postgr.es/c/fcb44c5e5) [&sect;](https://postgr.es/c/25b55788d)
-  Ensure that Memoize plan nodes use the correct input collation (Shihao Zhong) [&sect;](https://postgr.es/c/f79bf8d8f)

   This matters only when using non-deterministic collations.
-  Fix wrong results from Memoize plans when a Param is a subexpression of a cache key (Richard Guo) [&sect;](https://postgr.es/c/d6f60a031) [&sect;](https://postgr.es/c/db7e885d0)
-  Fix loss of parallel worker statistics for Memoize plan nodes underneath a Gather (Tatsuya Kawata) [&sect;](https://postgr.es/c/04609efb3)

   Only the statistics since the last rescan of the Gather were reported by `EXPLAIN ANALYZE`.
-  Don't create a shell type if a function definition specifies an unknown return type name with array or type modifier syntax (Heikki Linnakangas) [&sect;](https://postgr.es/c/b62206043)

   Such cases are surely errors, since you should create the type before doing anything advanced with it.
-  Fix snapshot export and import to work properly on a snapshot taken during recovery (Peter Geoghegan, Bertrand Drouvot) [&sect;](https://postgr.es/c/2ae67ada7)

   The data layout is different in such a snapshot, but the export/import logic didn't know that, so that the imported snapshot behaved incorrectly. This led not only to wrong answers within the importing transaction, but possibly to persistent data corruption due to it setting hint bits incorrectly.

   This bug is known to be triggered by parallel pg_dump, when run on a standby server. In that context it could corrupt both the dump output and the standby server's data.
-  Fix atomicity problem in transaction snapshot import (Chee Wooson) [&sect;](https://postgr.es/c/24bf23fe6)

   There was a window in which onlookers might not see the source transaction as still running, possibly allowing wrong decisions by `VACUUM` for example.
-  Fix double parse transformation of `PASSING` arguments in `JSON_TABLE` (Ayush Tiwari, Amit Langote) [&sect;](https://postgr.es/c/8793978e3)

   This mistake caused unexpected failures such as “unexpected non-SELECT command in SubLink”.
-  Fix failure to assign collations to `PASSING` expressions in `JSON_TABLE` (Venu Kummari) [&sect;](https://postgr.es/c/03550d3cf)

   This oversight could result in wrong answers or “could not determine which collation to use for string comparison” errors.
-  Fix usage of toasted text values in the `PASSING` arguments of `JSON_TABLE` and related functions (Chaitanya Choudhary, Shihao Zhong) [&sect;](https://postgr.es/c/8bccf0159)

   The code failed to apply detoasting, resulting in incorrect output.
-  Fix NULL result persistence in `JSON_VALUE() RETURNING` `json`/`jsonb` (Srinath Reddy Sadipiralla) [&sect;](https://postgr.es/c/31af62f38)

   Once `JSON_VALUE(... RETURNING ...)` with a `json` or `jsonb` result value returned NULL, every later evaluation of the same expression in the same statement also returned NULL, whether it should or not.
-  Fix `ON EMPTY/ERROR` result persistence in `JSON_VALUE` and related constructs (Jian He, Andrey Rachitskiy) [&sect;](https://postgr.es/c/775e03191)

   A NULL input could result in returning the value of the previous row's `ON EMPTY / ON ERROR` expression.
-  Fix handling of failures in `DEFAULT` expressions in `SQL/JSON` operations (Jian He) [&sect;](https://postgr.es/c/6dfe31d48)

   An error arising while executing a `DEFAULT` expression could be ignored, instead producing a NULL result.
-  Fix handling of domain coercion failures in `SQL/JSON` operations (Jian He) [&sect;](https://postgr.es/c/46798ab1a)

   `SQL/JSON` query functions could throw an error, or report the wrong one, in cases where an `ON ERROR` or `ON EMPTY` clause specified some non-error behavior.
-  Fix GIN index searches for plain scalar strings in strict `jsonpath` queries (Rui Zhao) [&sect;](https://postgr.es/c/19656b77d)

   A `jsonb_ops` index scan for a `jsonpath` such as `strict $ == "scalar"` failed to find table entries that were simply `"scalar"`.
-  Fix errors in extraction of GIN index scan conditions from jsonpath queries (Tom Lane) [&sect;](https://postgr.es/c/65f1b71c3)

   An `AND` operator underneath a `NOT` operator could be misimplemented. Also a method applied to a literal constant, for example `"42".double()`, was not handled correctly. These mistakes did not cause visible errors, but the GIN index search would miss some or all of the rows that should match.
-  Fix edge-case error in SP-GiST quad-tree `<@` searches (John Naylor) [&sect;](https://postgr.es/c/9ad421a9b)

   Due to roundoff error, the search might sometimes not descend into a quadrant that it needs to search, leading to incomplete query results.
-  Fix edge-case failures in SP-GiST quad-tree indexes (Ayush Tiwari) [&sect;](https://postgr.es/c/3d8c50317)

   Roundoff errors could allow searches and insertions to fail with “getQuadrant: impossible case” errors.
-  Accept empty `tsquery` values in binary input (Ayush Tiwari) [&sect;](https://postgr.es/c/e05b7e2a4)

   An empty `tsquery` value is allowed elsewhere, but `tsqueryrecv()` failed.
-  Disallow zero-length lexemes and out-of-range phrase distance values in binary input of `tsquery` values (Tom Lane) [&sect;](https://postgr.es/c/91989795d)

   These cases were already rejected during text input, but `tsqueryrecv()` omitted making equivalent checks. That allowed accepting values that would fail subsequent dump/reload.
-  Treat operators-nested-too-deeply errors as soft errors in `tsquery` input and `ltxtquery` input (Tom Lane) [&sect;](https://postgr.es/c/ffbe37115)
-  Consistently enforce maximum lengths in `tsvector` and `tsquery` (Tom Lane) [&sect;](https://postgr.es/c/b42d0c310)

   Some places thought the maximum lexeme length was `MAXSTRLEN` (2047), while others thought it was `MAXSTRLEN-1`. Fortunately, the looser bound is correct, so make them all use `MAXSTRLEN`. Likewise, some places were stricter than others about the total data length limit, `MAXSTRPOS`. These fixes avoid situations where a validly-constructed value would fail dump/reload.
-  Fix `right()` with a length argument of `INT_MIN` (-2147483648) (Ewan Young) [&sect;](https://postgr.es/c/3103bdafd)

   Due to integer overflow, the function returned the whole string when it should return an empty string.
-  Fix `pg_blocking_pids()` to account correctly for relation extension locks (Bertrand Drouvot) [&sect;](https://postgr.es/c/c0df269c5)
-  Fix 32-vs-64-bit confusion in `pg_indexam_progress_phasename()` (Markos Fountoulakis) [&sect;](https://postgr.es/c/8af1a12c9)

   This function accidentally worked on common hardware, but could misbehave on 32-bit or big-endian platforms.
-  Detect integer overflow in `to_char()` with `V` format (Andrey Rachitskiy) [&sect;](https://postgr.es/c/04c25af21)

   If the data value is an integer, and multiplying it by the power of ten implied by the `V` format causes integer overflow, `to_char()` silently emitted incorrect digits. Throw an integer overflow error instead.
-  Avoid integer overflow while calculating BRIN minmax_multi distances for `tid` values (Tomas Vondra) [&sect;](https://postgr.es/c/c83ddfcfe)

   Very large `tid` values could result in negative distances, which would end in making the index inefficient, though not actually corrupt.
-  Avoid overflow while calculating BRIN minmax_multi distances for numeric values, and handle NaNs properly (Kirill Reshke, Tomas Vondra) [&sect;](https://postgr.es/c/9ab2a624e)

   These oversights could result in errors being thrown (preventing index summarization), or in inefficient indexes.
-  Avoid overflow when computing differences for numeric ranges (Kirill Reshke) [&sect;](https://postgr.es/c/dfd98ce58)

   In this context we want to return plus or minus infinity rather than throwing an overflow error.
-  Support out-of-order `FLAG` lines in Hunspell affix files (Ewan Young) [&sect;](https://postgr.es/c/e96b27296)

   Hunspell affix files usually put the `FLAG` line first, and our implementation silently depended on that. If it came after some lines whose meaning depends on the `FLAG` setting, the result would be crashes or failures to match compound words. Fix the logic so that the `FLAG` line can be anywhere.
-  Fix assorted errors in rule deparsing (Henson Choi) [&sect;](https://postgr.es/c/91246c726) [&sect;](https://postgr.es/c/80da124a4) [&sect;](https://postgr.es/c/467bb7ea7) [&sect;](https://postgr.es/c/7933cf45a)

   Always print `AS` before a target-relation alias in `UPDATE`, `DELETE`, or `MERGE`, to avoid ambiguity if that alias is `set`. Also ensure that <em>xml</em><code> IS DOCUMENT</code> constructs and subquery left-hand arguments are adequately parenthesized. Also fix deparsing of `AT TIME ZONE` and similar constructs in a view's `FROM` clause; these need to be wrapped in a no-op `CAST`. These errors could prevent dump restoration due to a syntax error in a view, or cause reloaded views to be parsed as something different from before.
-  Fix parser failure when a multi-row `INSERT ... VALUES` has a sub-SELECT in a target column subscript (Amit Langote) [&sect;](https://postgr.es/c/f11673768)

   This led to “unexpected non-SELECT command in SubLink”.
-  Fix missed preprocessing of `WHERE` clauses in `COPY FROM` (Nathan Bossart) [&sect;](https://postgr.es/c/6b428e47b)

   This oversight led to “cache lookup failed for function 0” errors with negated operators, and probably in other cases as well.
-  Fix concurrent delete handling in `MERGE ... WHEN NOT MATCHED BY SOURCE` (Tender Wang) [&sect;](https://postgr.es/c/745152e45)

   When executing a `WHEN NOT MATCHED BY SOURCE` action in `MERGE`, if the tuple from the target table was concurrently deleted, the retry code would incorrectly execute a `WHEN NOT MATCHED BY TARGET` action for the row instead, giving incorrect results.
-  Fix race conditions with `DROP TABLESPACE` (Ayush Tiwari) [&sect;](https://postgr.es/c/960b6b532)

   A `DROP` executed concurrently with commands that attempt to add new objects to the tablespace could end with inconsistent catalog contents.
-  Cleanly reject operations on unlogged tables during recovery (Jim Jones) [&sect;](https://postgr.es/c/17f66a681) [&sect;](https://postgr.es/c/3e53104dc)

   An unlogged table will have no data on a standby, so reading it must fail. Most operations give an appropriate error for this, but `COPY ... TO` failed with an internal “could not open file” error, as did `currtid2()`.
-  Fix inconsistent permissions checking for role membership changes (ChangAo Chen) [&sect;](https://postgr.es/c/be0076b71)

   This fixes cases where what should be a permissions failure was reported as an internal “no possible grantors” error.
-  Revert previous changes to free space map persistence (Alexander Korotkov) [&sect;](https://postgr.es/c/bee33974e)

   Commits c06d1a4ba et al. (included in the May 2026 releases) turned out to cause more trouble than they fixed.
-  Avoid backend hang during temporary table cleanup at exit, if the default transaction mode is `SERIALIZABLE READ ONLY DEFERRABLE` (Andrey Rachitskiy) [&sect;](https://postgr.es/c/4639b6cfe)

   Cleanup would block until concurrent serializable transactions complete, which could be indefinitely if some of them are prepared transactions.
-  Avoid crash when describing a `FETCH` statement whose cursor is gone (Dirkjan Bussink) [&sect;](https://postgr.es/c/2affc1964)
-  Fix memory accounting for datum sorts on pass-by-reference data types (Mario Karuza) [&sect;](https://postgr.es/c/2414028ab)

   This error could lead to failure to enforce `work_mem` limits correctly, or to misleading output from `EXPLAIN ANALYZE`.
-  If a vacuum delay sleep is requested while holding locks, don't sleep (Kevin Rocker, Andrey Borodin, Mostafa) [&sect;](https://postgr.es/c/71d1ace33)

   Also, move several existing delay calls to points where it's actually safe to sleep.
-  Fix parallel vacuum I/O timing reports (Bharath Rupireddy) [&sect;](https://postgr.es/c/35c508af5)

   `VACUUM VERBOSE` and autovacuum logs included parallel workers' blocks-read and blocks-dirtied counts, but not the time they spent doing that I/O.
-  Avoid freeing the plan for a foreign-key enforcement query if the plan is still in use (Lucas Jeffrey, Trakshan Mishra, Amit Langote) [&sect;](https://postgr.es/c/f3ea7cb46)

   It was possible to reach this behavior in contrived cases, such as a self-referencing ON DELETE CASCADE foreign key on tables with triggers that perform DML.
-  Fix WAL-logging of multi-level page deletion in btree indexes (Peter Geoghegan) [&sect;](https://postgr.es/c/eea2ff293)

   The WAL log could contain stale values for a to-be-deleted page's sibling links, resulting in complaints on standby servers if `wal_consistency_checking` is enabled. Barring a crash, the damage would shortly be repaired by completion of the deletion action, so that this error was usually only cosmetic.
-  Avoid overly-strong WAL consistency check on buffers that are only pinned not locked (Greg Burd, Thomas Munro) [&sect;](https://postgr.es/c/fd1f61e43)

   A buffer that is not locked could be concurrently changed by other processes, so that the image captured for consistency checking is not self-consistent, leading to “incorrect resource manager data checksum in record” failures on standbys.
-  Ensure that state files for pending prepared transactions are written out atomically during a checkpoint (Shihao Zhong) [&sect;](https://postgr.es/c/15f57b683)

   Previously it was possible for a crash to leave an invalid state file behind, causing subsequent database startup to fail.
-  Do not start a checkpoint in the middle of aborting a prepared transaction (Shihao Zhong) [&sect;](https://postgr.es/c/56e831ea2)

   A crash would allow the transaction to be committed afterwards, leading to inconsistent state.
-  Do not start a checkpoint in the middle of creating a multixact (Shihao Zhong) [&sect;](https://postgr.es/c/4022edb93)

   This avoids possible data loss after a subsequent crash.
-  Fix possible loss of tuples updated by later-aborted subtransactions (Shihao Zhong, Michael Paquier) [&sect;](https://postgr.es/c/08f0cec8f)

   A table rewrite (for example, `CLUSTER` or `VACUUM FULL`) could discard such tuples as invalid.
-  Do not lose tuple's `tid` when storing it into a tuple table slot (Virender Singla, Greg Burd) [&sect;](https://postgr.es/c/7e594c30f)

   Some code paths neglected to copy the `tid` field, which could lead to data corruption or wrong query answers in unusual cases.
-  Avoid race condition in marking a partitioned index valid (Andrey Borodin, Kirill Reshke) [&sect;](https://postgr.es/c/bd8070126)

   This could, for example, cause parallel pg_restore to fail due to trying to recreate foreign key constraints too soon.
-  Fix checks for an aborted updating transaction in multixids (Chee Wooson) [&sect;](https://postgr.es/c/ff9eacc56)

   Known consequences of this coding error include unexpected “new multixact has more than one updating member” errors, as well as incorrect failures to update a row.
-  A hash index scan that starts during a bucket split could miss some index entries (Peter Geoghegan, Amit Kapila) [&sect;](https://postgr.es/c/ab2233743)
-  Fix incorrect position tracking in TID Range Scans (Ewan Young, David Rowley) [&sect;](https://postgr.es/c/25ebf0097)

   Sequences involving fetching backwards and then forwards again could fail to scan all the way to the end of the range.
-  Avoid unnecessary buffer locking when recovery replays freeze-only and VM-only WAL records (Melanie Plageman) [&sect;](https://postgr.es/c/5897bd1d8)

   This prevents unnecessary replay delays and cancellations of standby queries.
-  Avoid reading a potentially-torn copy of the `pg_control` file for base backups (David Steele) [&sect;](https://postgr.es/c/4aeefb26d) [&sect;](https://postgr.es/c/59bc83d41)

   A concurrent write to the file could result in including inconsistent data in the backup.
-  Prevent making incorrect copies of non-relation files in base backups (David Steele) [&sect;](https://postgr.es/c/1fbe31b9d)

   Files that routinely get replaced by atomic renames, such as `postgresql.auto.conf`, could be truncated or zero-padded in a backup made concurrently with such a replacement.
-  Ensure that base backups report checksum failures affecting single pages (Zsolt Parragi) [&sect;](https://postgr.es/c/7f92983dc)

   Due to a logic error, checksum failures were counted in `pg_stat_database` only when more than one page in the same file was corrupt.
-  Fix stale visibility maps in incremental backups after heap truncation (Ke, Fujii Masao, Shihao Zhong) [&sect;](https://postgr.es/c/0733c7219)

   An incremental backup taken after `VACUUM` truncates a table can restore stale visibility map bits when both data checksums and wal_log_hints are disabled, leading to data corruption.

   This fix does not correct existing WAL summaries, therefore incremental backups taken after updating can still be affected. It's recommended that users of incremental backups take a new full backup after updating and use it as the basis for subsequent incremental backups.
-  Reject incremental backups based on a backup from a divergent history (Sai Pranav, Fujii Masao) [&sect;](https://postgr.es/c/ecb43083f)

   After a failover, an incremental backup could accept a prior backup whose WAL range extends beyond the point where the new primary left the old timeline. Combining these backups could retain pages from the abandoned history, producing a corrupt backup.
-  Ensure that incremental backup merges WAL summaries in LSN order (Ke, Fujii Masao) [&sect;](https://postgr.es/c/9580e29da)

   This error could result in stale data being included in a combined backup.
-  Fix erroneous edge-case decisions about when to include WAL summaries in incremental backups (Ayush Tiwari) [&sect;](https://postgr.es/c/ec756e591)

   This error could result in stale data being included in a combined backup.
-  Fix incorrect size estimates for incremental backups (Daria Lepikhova) [&sect;](https://postgr.es/c/4c674b952)

   This problem affected progress reporting in pg_basebackup and other places.
-  Verify the syntax of the `output_plugin_libraries` parameter when a value is assigned (Jacob Champion, Hayato Kuroda) [&sect;](https://postgr.es/c/7c78c8cff)

   Previously, a syntactically-incorrect value could be put into `postgresql.conf`, causing logical replication to fail.
-  Avoid incorrect update of backend status flags during replication slot invalidation (Vlad Lesin) [&sect;](https://postgr.es/c/295f75b4b)

   An invalidation done by an auxiliary process such as the checkpointer could clobber the status field of an unrelated backend. The consequences of this are difficult to predict, but one possibility is that hot standby feedback could be effectively disabled.
-  In logical replication, handle subtransaction aborts better (Bingshuai Li, Hayato Kuroda, Zhijie Hou) [&sect;](https://postgr.es/c/624bd2b1f)

   If a subtransaction that made a catalog update aborts, logical decoding mistakenly still believed that the update was effective, potentially causing incorrect decoding output.
-  Avoid race condition when building an initial decoding snapshot (Antonin Houska, Rui Zhao) [&sect;](https://postgr.es/c/b93df6eac)

   There was a small risk of data corruption from using a snapshot that includes transactions that have committed but not yet set their CLOG status.
-  Avoid null-pointer crash with missing replication option values in `pgoutput` logical replication sender (Zhijie Hou) [&sect;](https://postgr.es/c/c3ef3f228)
-  Fix small memory leak in pgoutput with streamed prepared transactions (Ayush Tiwari) [&sect;](https://postgr.es/c/3d6a4a582)
-  Fix lost wakeups in physical WAL senders (Alex Speller) [&sect;](https://postgr.es/c/6d34d5500)

   This fix corrects some cases where a WAL sender process could sleep even though there is data waiting to be sent.
-  Fix duplicate creation of temporary replication slot when a WAL receiver follows a timeline switch (ChangAo Chen) [&sect;](https://postgr.es/c/4c8fa02c5)

   This mistake led to a “could not create replication slot” error followed by a restart of the WAL receiver, which would then recover, so there was no permanent damage.
-  Fix race condition between logical replication subscriber and concurrent index DDL (Mikhail Nikalayeu, Vignesh C) [&sect;](https://postgr.es/c/0405dcf7c)

   Changing the state of a replication target relation's primary key or replica identity index could cause the replication apply logic to miss finding a matching row (but only once per such event).
-  Fix incorrect LSNs in WAL decoding error messages (Rahul Yadav) [&sect;](https://postgr.es/c/5f472ab56)

   These error messages reported the LSN of a WAL record that wasn't the one they were actually complaining about.
-  Don't allow `REPLICA IDENTITY FULL` lookups to select invalid indexes (Mikhail Nikalayeu) [&sect;](https://postgr.es/c/92c018214)

   Previously, logical replication could try to use an index left behind by a failed `CREATE INDEX CONCURRENTLY`, leading to wrong results since the index is likely incomplete.
-  Fix potential use-after-free cases in WAL sender (Ayush Tiwari) [&sect;](https://postgr.es/c/39173a9bd)
-  Avoid use-after-free in `jsonb_set_lax()` (John Naylor) [&sect;](https://postgr.es/c/723075b03)
-  Fix double-free when `XMLSERIALIZE(... INDENT)` output exceeds 1GB (Jim Jones) [&sect;](https://postgr.es/c/7fa0a5f21)
-  Fix null-pointer dereferences that could occur when a schema is concurrently dropped (Baji Shaik) [&sect;](https://postgr.es/c/f1ba2b826)
-  Avoid null-pointer crash in logical subscription refresh when a relevant relation is concurrently dropped (Satyanarayana Narlapuram, Bharath Rupireddy) [&sect;](https://postgr.es/c/ebf7971da)
-  Fix possible null-pointer dereference when building extended statistics on expressions (Shihao Zhong, Michael Paquier) [&sect;](https://postgr.es/c/8b517a30c)

   This error was reachable only with unexpected (but legal) behavior of a custom `typanalyze` function.
-  Fix possibly-stale values of `pg_locks`.`waitstart` (Shihao Zhong) [&sect;](https://postgr.es/c/0165f38d1)
-  Harden pgstats code against partially-completed local pgstats entries (Michael Paquier) [&sect;](https://postgr.es/c/6e0011596)
-  Fix cleanup of reused pgstats entries (Michael Paquier) [&sect;](https://postgr.es/c/1613efa1a)

   This error led to leakage of pgstats entries, which could result in errors at backend exit or database shutdown.
-  Fix memory leaks caused by not releasing all data associated with a WAL reader (Tomás Senart) [&sect;](https://postgr.es/c/aa6f14da4)

   In many cases this oversight did little harm because the containing memory context would soon be reset or deleted. However there are cases where it could cause long-lived memory bloat, notably in the checkpointer process.
-  Avoid memory bloat while computing `ndistinct` extended statistics across many columns (Ilia Evdokimov) [&sect;](https://postgr.es/c/a8c98c884)
-  Fix buffer pin leak after a failed hash index bucket split cleanup (Peter Geoghegan) [&sect;](https://postgr.es/c/b57161883)

   This coding error resulted in nothing worse than warnings about resource leakage.
-  Ensure all relations opened for `AFTER` triggers are closed again (Hayato Kuroda, David Rowley) [&sect;](https://postgr.es/c/70d5b0668)

   This oversight triggered “resource was not closed” warnings, but nothing more serious than that.
-  Fix PL/pgSQL's handling of `anycompatiblemultirange` arguments and results (Jeevan Chalke) [&sect;](https://postgr.es/c/fa10d961d)

   This type had not been added to PL/pgSQL's knowledge of polymorphic types, resulting in errors when trying to create or use a function declared with it.
-  Harden PL/Perl and associated contrib modules against “tied” Perl hashes and arrays (Tom Lane) [&sect;](https://postgr.es/c/3c4d77480) [&sect;](https://postgr.es/c/2860d322a)

   Rely on `av_count()` to decide the length of a Perl array, rather than looking for a null pointer; and contrarily be more wary about unexpected null value pointers from array and hash fetches. These cases are only known to arise in connection with “tied” objects.
-  Avoid memory leakage in libpq with repeated GSS connections (Matt Carter, Tom Lane) [&sect;](https://postgr.es/c/df9e6d7c7)
-  Avoid incorrect behavior when `getopt_long()` rearranges a program's command line arguments (Sehrope Sarkuni) [&sect;](https://postgr.es/c/a11bce64a)

   For example, `vacuumdb postgres --jobs` took `postgres` as the number of jobs instead of complaining that `--jobs` is missing its argument.
-  Ensure that psql counts all the `COPY FROM STDIN` commands appearing in a multi-query command string (Zsolt Parragi) [&sect;](https://postgr.es/c/3828ab704)

   After the fix for CVE-2026-6464, psql would reject multiple `COPY FROM STDIN` commands appearing in a single string, because it was not expecting more than one `COPY` operation to occur.
-  Make psql cope better with out-of-memory conditions during backslash command processing (Fujii Masao) [&sect;](https://postgr.es/c/ab001a3a6) [&sect;](https://postgr.es/c/d2e19268d)
-  Fix some small memory leaks in psql (Fujii Masao) [&sect;](https://postgr.es/c/4c5485aed)
-  In pg_combinebackup, reject out-of-range block numbers in incremental backup files (Peter Eisentraut) [&sect;](https://postgr.es/c/0bffea8f6)

   This omission meant that pg_combinebackup could crash on corrupted input.
-  Make pg_dump dump column statistics targets for indexes backing constraints (Zsolt Parragi) [&sect;](https://postgr.es/c/5e158a2e4)

   Formerly pg_dump would emit `ALTER INDEX ... SET STATISTICS` commands when needed for plain indexes, but not for indexes created via `EXCLUDE` constraints on expressions.
-  In pg_dump's tar-format output, don't put invalid `\restrict`/`\unrestrict` commands into the `restore.sql` tar member (Pavel Kopylov) [&sect;](https://postgr.es/c/45fc3ce22)
-  Make pg_walsummary's `--quiet` switch suppress limit block output (Koshi Shibagaki) [&sect;](https://postgr.es/c/d843a801b)
-  Make reindexdb error out cleanly if an index-level `REINDEX CONCURRENTLY` is requested with multiple jobs (Kirill Reshke) [&sect;](https://postgr.es/c/869c53ff8)

   Previously, this combination of options would lead to a server-side error, complaining that `REINDEX CONCURRENTLY` cannot run within a transaction block.
-  Fix `contrib/file_fdw` to not crash in `ANALYZE` when the table has column default clauses (Fujii Masao) [&sect;](https://postgr.es/c/b2b1910d9)
-  Fix possible index corruption in `contrib/pg_trgm` GiST indexes (Nate Clark) [&sect;](https://postgr.es/c/145a7175e)

   Combination of index entries could lose the flag identifying their format, so that subsequent index searches could miss matching rows.

   This fix does not repair existing indexes. Users should `REINDEX` existing `gist_trgm_ops` indexes after updating.
-  Fix error check in timer setup on Windows (Sehrope Sarkuni) [&sect;](https://postgr.es/c/2bc73b3e8)

   Failure to create a thread to manage timeouts was not detected, leading to silent failure of subsequent timeout requests.
-  Use a non-locking initial test in TAS_SPIN on AArch64 (Nathan Bossart) [&sect;](https://postgr.es/c/e10dc5b85)

   This code change has been shown to improve performance on many-core AArch64 machines while not hurting performance on smaller ones; moreover, it avoids a performance regression with newer Linux kernels.
-  Update time zone data files to tzdata release 2026f (Tom Lane) [&sect;](https://postgr.es/c/1ad631142) [&sect;](https://postgr.es/c/8e1414de3) [&sect;](https://postgr.es/c/5f0f69019)

   Canada's Northwest Territories are on year-round UTC-06 (effectively, permanent DST) beginning in November 2026. This affects only America/Inuvik as the rest of the territory is covered by America/Edmonton, for which the equivalent change was released in 2026c. Similarly, Manitoba (America/Winnipeg) moves to year-round UTC-05. For now assume their TZ abbreviations will be CST and EST respectively.

   Cuba did not fall back on 2026-11-01; they will stay on DST this winter (but apparently plan to revert to normal next year).

   The backward-compatibility zone names EST5EDT, CST6CDT, MST7MDT, and PST8PDT now have simpler behavior for old time stamps. They use the indicated UTC offsets and TZ abbreviations indefinitely far back, rather than allowing aspects of the associated city-named zone's historical behavior to leak through. This conforms better to POSIX.

   Historical corrections for Brazil, Colombia, Iran, and Ireland.

   On Windows, initdb will now recognize the new time zone names that Microsoft has devised for the updated Canadian zones.
