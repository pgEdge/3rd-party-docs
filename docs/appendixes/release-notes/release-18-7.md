<a id="release-18-7"></a>

## Release 18.7


**Release date:.**


2026-11-12


 This release contains a variety of fixes from 18.6. For information about new features in major release 18, see [Release 18](release-18.md#release-18).
 <a id="release-18-7-migration"></a>

### Migration to Version 18.7


 A dump/restore is not required for those running 18.X.


 However, several different bugs have been discovered that could result in index corruption. Examine the changelog entries below that mention `REINDEX`, and reindex any indexes you have that might be affected.


 Also, existing incremental backup files should not be trusted, thanks to the discovery of several ways that they could become corrupt. It's recommended that users of incremental backups take a new full backup after updating and use it as the basis for subsequent incremental backups.


 Also, if you are upgrading from a version earlier than 18.6, see [Release 18.6](release-18-6.md#release-18-6).
  <a id="release-18-7-changes"></a>

### Changes


-  Add defenses against problems caused by altering composite types concurrently with operations on data of those types (Heikki Linnakangas, Andres Freund, Nathan Bossart, Nikhil Sontakke) [&sect;](https://postgr.es/c/afaff70f0) [&sect;](https://postgr.es/c/692f4b3c4) [&sect;](https://postgr.es/c/9ef1cd650) [&sect;](https://postgr.es/c/9ddd897b3) [&sect;](https://postgr.es/c/1bb9942e2) [&sect;](https://postgr.es/c/bcc692493) [&sect;](https://postgr.es/c/bb66a4e82) [&sect;](https://postgr.es/c/05fc87f95) [&sect;](https://postgr.es/c/d48d28c4b) [&sect;](https://postgr.es/c/67d89bb22) [&sect;](https://postgr.es/c/e72f4e91c) [&sect;](https://postgr.es/c/5db4823d0) [&sect;](https://postgr.es/c/26ff7cbd3) [&sect;](https://postgr.es/c/fdbb7ffbb) [&sect;](https://postgr.es/c/ff6d9deca) [&sect;](https://postgr.es/c/7bc2d46d1) [&sect;](https://postgr.es/c/dcb2ff56d) [&sect;](https://postgr.es/c/7c32010ca) [&sect;](https://postgr.es/c/da2cb6090) [&sect;](https://postgr.es/c/8b1a2a8ac) [&sect;](https://postgr.es/c/391884d5c) [&sect;](https://postgr.es/c/edcd260d9) [&sect;](https://postgr.es/c/d835dc88d) [&sect;](https://postgr.es/c/cfdd09589) [&sect;](https://postgr.es/c/5f0add5cf) [&sect;](https://postgr.es/c/30d81a0ba) [&sect;](https://postgr.es/c/edcc027ee) [&sect;](https://postgr.es/c/412297476)

   Operations such as `ALTER TABLE ... ALTER COLUMN TYPE` are hazardous because tuples of the older row layout may still exist in the system. While we have long had some checks to defend against that, they were never really complete, and they entirely failed to detect cases where composite types were changed concurrently with queries using them. This voluminous set of patches attempts to improve that.

   Introduce versioning logic to permit detection of composite datums that don't match the tuple descriptor we wish to interpret them with. For full protection, this change requires rebuilding C-language extensions, and there are some extensions that will require source-code changes. If necessary, the new `composite_type_mismatch` server parameter can be used to disable these checks, but doing so carries security risks.

   Add locking to prevent cases such as `ALTER TYPE` on a composite type used in a `CREATE TABLE` that hasn't committed yet.

   Tighten the rules about whether `ALTER COLUMN TYPE` is allowed when dependencies exist on the composite type. For example, it's no longer allowed when a function's default expression or new-style SQL function body depends on the type.

   The PostgreSQL Project thanks Andrew Gierth, Alexander Lakhin, Francesco Verardi, Hcamael, Tomer Fichman, OpenAI Security Research Team, Peter Geoghegan, Mehmet Ince, Bryan Green, Sumedh Pathak, Sergey Shinderuk, Yuxiao Guo, Nikhil Sontakke, Daniel Bakker, the Anthropic OSS program, Claude and Ada Logics for reporting this problem. (CVE-2026-6460)
-  Cope with concurrent rowtype changes in `ExecEvalFieldStore[De]Form()` (Tomas Vondra, Heikki Linnakangas) [&sect;](https://postgr.es/c/4ccea70a1)

   The PostgreSQL Project thanks OpenAI Codex Security, Bronson Yen, Mehmet Ince, Claude and Ada Logics for reporting this problem. (CVE-2026-116306)
-  Guard against malicious alteration of a table while a statistics import function is running on it (Masahiko Sawada) [&sect;](https://postgr.es/c/780900dcf)

   Previously it was possible to store statistics that no longer match the table, leading to crashes or other problems later.

   The PostgreSQL Project thanks Mehmet Ince for reporting this problem. (CVE-2026-116329)
-  Enforce view row visibility in `INSERT ... ON CONFLICT` (Peter Geoghegan, Zsolt Parragi) [&sect;](https://postgr.es/c/a244ef271)

   When applied to a view, `INSERT ... ON CONFLICT DO UPDATE` could update rows that should be hidden by the view's `WHERE` clause. Although it was documented, this behavior is surely unexpected, and it gives rise to security problems: a user can update rows he should have no privilege to reach. Updating is incorrect, but inserting is impossible too (it would fail due to the unique arbiter index), so throw an error instead when the case arises. `INSERT ... ON CONFLICT DO SELECT` could show such rows as well, but no longer does.

   The PostgreSQL Project thanks Qing Xu, Carl Andersen, Zsolt Parragi, the Anthropic OSS program, Shihao Zhong, and Josef Šimánek for reporting this problem. (CVE-2026-16237)
-  Align `REASSIGN OWNED`'s caller-privilege checks with `ALTER ... OWNER TO`'s checks (Zsolt Parragi) [&sect;](https://postgr.es/c/ea228688e)

   Since v16, giving an object to another role requires the ability to `SET ROLE` to that role. However, a role granted `INHERIT TRUE, SET FALSE` could bypass that via `REASSIGN OWNED`. A role granted `INHERIT FALSE, SET TRUE` saw the inverse problem: `REASSIGN OWNED` failed while `ALTER OWNER` succeeded. Make `REASSIGN` follow the `ALTER OWNER` rules.

   `REASSIGN OWNED` also bypassed the `CREATEDB` privilege check in `ALTER DATABASE OWNER`.

   In v15 and v14, `REASSIGN OWNED`'s privilege check is stricter than the membership checks in `ALTER OWNER`. Leave that alone, but still fix the missing `CREATEDB` check.

   The PostgreSQL Project thanks Zsolt Parragi for reporting this problem. (CVE-2026-120146)
-  Fix use-after-free in `COPY FROM ON_ERROR` reporting (Masahiko Sawada) [&sect;](https://postgr.es/c/8d9f203b8)

   An error report raised in `COPY FROM` could disclose server memory.

   The PostgreSQL Project thanks Jiyong Ahn, Claude and Ada Logics for reporting this problem. (CVE-2026-116310)
-  Disallow closing a portal that is still being set up (Michael Paquier) [&sect;](https://postgr.es/c/8a2982de9)

   This change forestalls various use-after-free issues.

   The PostgreSQL Project thanks Tomas Vondra for reporting this problem. (CVE-2026-116332)
-  Forcibly close portals associated with transition tables at the end of the trigger function (Jeff Davis) [&sect;](https://postgr.es/c/983bcb768)

   This change avoids use-after-free hazards.

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120143)
-  Fix race condition leading to use-after-free in shared tupledesc management (Robert Haas) [&sect;](https://postgr.es/c/861b5a869)

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-116304)
-  Fix use-after-free in `pg_show_all_settings()` (Masahiko Sawada) [&sect;](https://postgr.es/c/921b93f25)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120152)
-  Fix use-after-free in expanded-record domain constraint checking (Tom Lane) [&sect;](https://postgr.es/c/2d91c1c78)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120150)
-  Prevent memory corruption after a large allocation in a “generation” memory context is freed (Heikki Linnakangas) [&sect;](https://postgr.es/c/53a34a4de)

   The PostgreSQL Project thanks Jayden Sarveshkumar for reporting this problem. (CVE-2026-116328)
-  Fix buffer overrun in `tsvector_delete()` (Anthropic OSS program) [&sect;](https://postgr.es/c/9c2cadacd)

   With crafted input, `tsvector_delete()` could be induced to write more bytes than it had allocated for its result.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116315)
-  Fix buffer overrun in `ts_rank_cd()` (Anthropic OSS program) [&sect;](https://postgr.es/c/3e97c9a1f)

   With crafted input, `ts_rank_cd()` could be induced to write more bytes than it had allocated for its result.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116317)
-  Prevent construction of invalid `name` values via element assignment (Tom Lane) [&sect;](https://postgr.es/c/b3ef92518)

   `array_set_element()` could be abused to make the last byte of a `name`-type value nonzero. Subsequent uses of the value would then read past the end, resulting in crashes or disclosure of server memory. While at it, disallow construction of invalidly-encoded strings through this mechanism.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116321)
-  Treat application of a leaky function to a `CaseTestExpr` or internal `Param` node as leaky (Tom Lane) [&sect;](https://postgr.es/c/9a96032e1)

   A `CaseTestExpr` or internal `Param` represents the value of an expression somewhere else in the query. If that expression contains a `Var` then there is potential for information leakage. Therefore, reject such cases when deciding whether an expression is leak-proof.

   The PostgreSQL Project thanks Only see myself for reporting this problem. (CVE-2026-116327)
-  Don't generate an index-based plan for an aggregate whose input and output types don't match (Tom Lane) [&sect;](https://postgr.es/c/2c03c7c48)

   Previously, we'd blindly do this if the aggregate has a `SORTOP` option attached. The index-based plan will return the aggregate's input column, leading to type violations and potentially to arbitrary code execution if its declared output type is different. In the master branch, disallow such aggregate declarations. In released branches, add a plan-time check to not generate such plans if the types don't match.

   The PostgreSQL Project thanks Zheng Yu, Claude and Ada Logics for reporting this problem. (CVE-2026-116331)
-  Delay processing of cancel request packets until it's safe to handle them (Michael Paquier) [&sect;](https://postgr.es/c/e1cc4a893)

   With luck and good timing, an attacker could induce the process receiving a cancel request packet to exit while holding a spinlock, thereby causing a database crash and restart.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116322)
-  Avoid inserting the wrong tuple into a trigger transition table in a `MERGE` that does an `INSERT` followed by an `UPDATE` (Heikki Linnakangas, Jacek Bylina, Jacob Robertson, Nigel Gott) [&sect;](https://postgr.es/c/f64f71188)

   The PostgreSQL Project thanks Jacek Bylina, Timothy Ball, Claude and Ada Logics for reporting this problem. (CVE-2026-116312)
-  Check that commutator and negator operator pairs return the same data type (Tom Lane) [&sect;](https://postgr.es/c/165bbc265)

   Previously, substitution of a commutator or negator operator could lead to a type violation. This patch also cleans up some cases where one operator had a commutator/negator link to another, but the second operator didn't link back to the first.

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-116335)
-  Prevent functions in index predicates and expressions from breaking `REINDEX CONCURRENTLY` by manipulating the indexes (Michael Paquier) [&sect;](https://postgr.es/c/061cc0ee6)

   The PostgreSQL Project thanks Florent Saudel for reporting this problem. (CVE-2026-116336)
-  Prevent corrupting a newly-installed extension by renaming its target schema during `CREATE EXTENSION` or by installing it into a temporary schema (Noah Misch) [&sect;](https://postgr.es/c/68ee5a816) [&sect;](https://postgr.es/c/b552d6efb)

   With some ingenuity, an attacker could substitute objects of his creation for objects in the extension.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-6467)
-  Reject more unsafe names in extension script substitutions (Nathan Bossart) [&sect;](https://postgr.es/c/d3c9f8c61)

   We already rejected schema and owner names containing quotes, backslashes, and dollar signs, but names containing newlines, carriage returns, comment delimiters, or the substitution tokens themselves can still achieve SQL injection. Reject all those cases too.

   The PostgreSQL Project thanks the Anthropic OSS program, Sven Klemm, and Aaron Chen for reporting this problem. (CVE-2026-120153)
-  Check `EXECUTE` privilege before calling cast functions within `to_json()` and other conversion-to-`JSON` code (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/3531f4b34)

   Typically such functions would have public `EXECUTE` privilege anyway, but this puts the behavior on par with normal execution of casts.

   The PostgreSQL Project thanks Filip Janus and the Anthropic OSS program for reporting this problem. (CVE-2026-116303)
-  Check `EXECUTE` privilege before calling cast functions within foreign key enforcement checks (Nathan Bossart) [&sect;](https://postgr.es/c/d9b15fda5)

   Typically such functions would have public `EXECUTE` privilege anyway, but this puts the behavior on par with normal execution of casts.

   The PostgreSQL Project thanks Noah Misch for reporting this problem. (CVE-2026-116303)
-  Prevent `DROP OWNED BY` from changing row-level security policies on tables that the current user does not own (Zsolt Parragi, Anthropic OSS program) [&sect;](https://postgr.es/c/7f30f607a)

   Previously, an unprivileged role could remove itself from restrictive policies on tables it does not own. Instead issue a warning, as is already done for privileges that the current user cannot revoke.

   The PostgreSQL Project thanks Zsolt Parragi and the Anthropic OSS program for reporting this problem. (CVE-2026-116311)
-  Disallow security-restricted operations from creating types in temporary schemas (Anthropic OSS program, Nathan Bossart) [&sect;](https://postgr.es/c/448db1451)

   We already disallowed creating temporary tables in this context, but a type also presents a hazard of shadowing permanent objects and thus allowing privilege escalation.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116316)
-  Avoid out-of-bounds accesses in FDW batch inserts (Nathan Bossart) [&sect;](https://postgr.es/c/1e0a55a0f)

   The PostgreSQL Project thanks Sumedh Pathak and Sven Klemm for reporting this problem. (CVE-2026-120149)
-  Avoid out-of-bounds reads in parallel btree scans with array keys (Nathan Bossart) [&sect;](https://postgr.es/c/1dfde5b37)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-116333)
-  Replace the type cache's `in_progress_list` mechanism with something simpler and safer (Robert Haas) [&sect;](https://postgr.es/c/2950cc789)

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120144)
-  Refuse to create extended statistics on expressions of transient types (Masahiko Sawada) [&sect;](https://postgr.es/c/74a999de1)

   Disallow `CREATE STATISTICS` on expressions whose types would not be allowed as table column types, most notably anonymous record types. This prevents other sessions from misinterpreting the stored values.

   The PostgreSQL Project thanks Mehmet Ince for reporting this problem. (CVE-2026-120148)
-  Reject system catalogs as logical replication targets (Masahiko Sawada) [&sect;](https://postgr.es/c/cc2d1a4c9)

   The PostgreSQL Project thanks Qian Zhang for reporting this problem. (CVE-2026-120145)
-  Check `ON CONFLICT` target row against RLS policies before applying user quals (Peter Geoghegan, Zsolt Parragi) [&sect;](https://postgr.es/c/a748ed909)

   Previously, leaky qual expressions could see rows that RLS policies should forbid them from seeing.

   The PostgreSQL Project thanks Peter Geoghegan, Zsolt Parragi, Only see myself, the Anthropic OSS program, Joao Detomini, Josef Šimánek, and Qing Xu for reporting this problem. (CVE-2026-120154)
-  Validate all connection options in replication subscriptions and `contrib/dblink` (Jacob Champion) [&sect;](https://postgr.es/c/780809ba6)

   Previously, these checks were applied inconsistently.

   The PostgreSQL Project thanks Andrey Krasichkov, Bashir Altawil, and Zsolt Parragi for reporting this problem. (CVE-2026-14661)
-  Handle dropped columns safely while comparing row types (Peter Geoghegan) [&sect;](https://postgr.es/c/35d851811)

   Previously, two row types that had dropped columns at the same position could be considered equal even if those columns were of different physical widths. That leads to problems when trying to deconstruct tuples of those types.

   The PostgreSQL Project thanks Peter Geoghegan and Mehmet Ince for reporting this problem. (CVE-2026-120142)
-  Disallow declaring SQL functions as returning type `"any"` (Peter Geoghegan) [&sect;](https://postgr.es/c/ff961cbbf)

   This allowed type violations.

   The PostgreSQL Project thanks Peter Geoghegan, Claude and Ada Logics for reporting this problem. (CVE-2026-18447)
-  Harden more memory-size calculations against integer overflow (Álvaro Herrera) [&sect;](https://postgr.es/c/317b7bcab) [&sect;](https://postgr.es/c/f1b226dee)

   The PostgreSQL Project thanks Hillai Ben Sasson, Jayden Sarveshkumar, and Mehmet Ince for reporting this problem. (CVE-2026-120139)
-  Fix type confusion in set-returning functions that return `RECORD`, when materialized in a query's target list (Melanie Plageman, Robert Haas) [&sect;](https://postgr.es/c/4fb349500)

   The same tuple descriptor was used across multiple calls, which is unsafe since the actual record type could vary.

   The PostgreSQL Project thanks Peter Geoghegan for reporting this problem. (CVE-2026-120141)
-  Cope with concurrent rowtype changes in PL/pgSQL record field assignment (Tom Lane) [&sect;](https://postgr.es/c/6b1ca1333)

   The PostgreSQL Project thanks Claude and Ada Logics for reporting this problem. (CVE-2026-120151)
-  Support “tied” Perl objects in PL/Perl and associated contrib modules (Andrey Rachitskiy) [&sect;](https://postgr.es/c/bbe476c74)

   Up to now, if you passed a tied object to a PL/Perl operation, it would most likely get treated as `undef`, since we failed to perform Perl get magic on values we're passed. Tied hashes worked partially, in that you could read out their keys, but the values always read as `undef`.

   Originally this was intended as a master-only change, since it looks like a feature addition. But afterwards we discovered that assorted security-grade misbehaviors could be triggered by tied objects even without this patch. Rather than trying to secure two sets of behaviors, it seems best to back-patch this change first. (CVE-2026-116326)
-  Fix PL/Perl to cope with errors thrown from tied objects (Tom Lane, Andrew Dunstan) [&sect;](https://postgr.es/c/8d308caee) [&sect;](https://postgr.es/c/3ce3d94c0)

   The PostgreSQL Project thanks Ngoc Tran, the Anthropic OSS program, and Dwij Mehta for reporting this problem. (CVE-2026-116326)
-  Fix type confusion in PL/Perl's `get_perl_array_ref()` (Andrew Dunstan) [&sect;](https://postgr.es/c/030b28a8d)

   The PostgreSQL Project thanks Zheng Yu for reporting this problem. (CVE-2026-116309)
-  Fix use-after-free in scans of Perl arrays and hashes (Tom Lane, Anthropic OSS program) [&sect;](https://postgr.es/c/5215e9d6a)

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116320)
-  Fix use-after-free in `plperl_return_next_internal()` (Anthropic OSS program) [&sect;](https://postgr.es/c/343663966)

   The PostgreSQL Project thanks the Anthropic OSS program and Ngoc Tran for reporting this problem. (CVE-2026-116319)
-  In PL/Perl, fix use-after-free triggered by releasing a prepared query that is being executed (Masahiko Sawada) [&sect;](https://postgr.es/c/250405afb)

   The PostgreSQL Project thanks Zheng Yu and the Anthropic OSS program for reporting this problem. (CVE-2026-116307)
-  Fix use-after-free hazards in re-entrant use of PL/Tcl (Tom Lane) [&sect;](https://postgr.es/c/f01727316)

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116323)
-  Reject re-entrant use of `contrib/dblink` (Masahiko Sawada) [&sect;](https://postgr.es/c/b8fbeeb76)

   The PostgreSQL Project thanks Jiyong Ahn for reporting this problem. (CVE-2026-120147)
-  In `contrib/dict_int`, return NULL for zero-length tokens (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/8373a4215)

   This behavior is consistent with other text-search dictionary types, and it avoids disclosure of server memory as was possible previously.

   The PostgreSQL Project thanks Nathan Bossart and the Anthropic OSS program for reporting this problem. (CVE-2026-116308)
-  Ensure that reindexdb correctly quotes index names it is dealing with (Anthropic OSS program) [&sect;](https://postgr.es/c/aac5e6f50)

   Crafted index names could yield a SQL-injection attack against the user running reindexdb.

   The PostgreSQL Project thanks the Anthropic OSS program for reporting this problem. (CVE-2026-116325)
-  Use fully-qualified operator names in vacuumdb (Nathan Bossart, Anthropic OSS program) [&sect;](https://postgr.es/c/e10d38d47)

   This prevents attacks on the user running vacuumdb.

   The PostgreSQL Project thanks Noah Misch and the Anthropic OSS program for reporting this problem. (CVE-2026-116313)
-  Ensure that pg_dump quotes the names of `NOT NULL` constraints (Anthropic OSS program) [&sect;](https://postgr.es/c/9e4cde792)

   Its failure to do so in some cases risked SQL injection, or even just restoring the constraint with the wrong name.

   The PostgreSQL Project thanks the Anthropic OSS program, Sarath Kumar, and Nikolay Samokhvalov for reporting this problem. (CVE-2026-116318)
-  Make the `\[un]restrict` commands functional in pg_dump's tar-format output (Noah Misch) [&sect;](https://postgr.es/c/17e57d06c)

   This provides the intended level of protection to users who make use of the undocumented `restore.sql` tar member.

   The PostgreSQL Project thanks Pavel Kopylov for reporting this problem. (CVE-2026-116330)
-  Harden pg_restore against corrupt archive files (Dilip Kumar, Yonghwa Lee, Nathan Bossart) [&sect;](https://postgr.es/c/673ef2341)

   This patch fixes out-of-bounds accesses and file path traversal hazards. There remain other issues that don't seem to amount to security bugs.

   The PostgreSQL Project thanks Dilip Kumar, Yonghwa Lee, OpenSec, and Ben Dixon for reporting this problem. (CVE-2026-116314)
-  Fix use of externally-controlled format string in pg_basebackup (Peter Eisentraut) [&sect;](https://postgr.es/c/af250ef26)

   This problem could only be reached if files with unusual names exist in the server's data directory.

   The PostgreSQL Project thanks Peter Eisentraut for reporting this problem. (CVE-2026-16240)
-  Disable btree index de-duplication support for the `bpchar` data type (Shihao Zhong, Peter Geoghegan) [&sect;](https://postgr.es/c/882fd416d)

   Index de-duplication is unsafe for `bpchar` because its notion of equality is weaker than bitwise equality. In particular, because `bpchar` equality ignores trailing spaces, values of different lengths can compare equal. De-duplication then allows index-only scans to return subtly wrong results.

   This patch disables de-duplication for new index entries, but does nothing about possible corruption of existing entries. Therefore, users who have btree indexes on `bpchar` columns should `REINDEX` them after installing this update.

   Note that this problem only applies to columns declared as `bpchar`, that is, lacking any length constraint. The more common usage of <code>CHAR(</code><em>n</em><code>)</code> is not at risk, even though it shares the same underlying code, because the length constraint forces otherwise-equal strings to have the same length and thus to be bitwise equal.
-  Ensure that `VACUUM` cleans up all dead tuples in a GIN index (Peter Geoghegan) [&sect;](https://postgr.es/c/b0a36462b) [&sect;](https://postgr.es/c/9900966c7)

   When a single `VACUUM` operation requires multiple index scans (because it is removing more tuples than it can hold in memory at once), it was possible to miss removing some index entries that should be removed. An additional bug involving a race with leaf page splits had the same effect. These bugs would eventually lead to a corrupt index with entries pointing at unrelated table rows. `REINDEX` is recommended if you suspect this has happened.
-  Fix GIN posting tree page deletion after an incomplete page split (Peter Geoghegan) [&sect;](https://postgr.es/c/9448c3e7b)

   The deletion step could incorrectly disconnect sibling pages, causing subsequent index searches to miss rows. `REINDEX` is recommended if you suspect this has happened.
-  Fix failure to update mutable `pg_index` fields for TOAST-table indexes (Bharath Rupireddy) [&sect;](https://postgr.es/c/e4807a0f3)

   `REINDEX CONCURRENTLY` can change these fields, but the relation cache was not expecting that, allowing index updates to be missed. The end result could be errors about “missing chunk number” due to missing index entries. Fortunately, these only represent index corruption; the TOAST table itself is fine, so a `REINDEX` will suffice to recover.
-  Perform join removal by editing the query's parse tree, rather than trying to fix up the planner's derived data structures (Tom Lane, Richard Guo, Shihao Zhong, Ayush Tiwari) [&sect;](https://postgr.es/c/9f25197bf) [&sect;](https://postgr.es/c/37c2bc950) [&sect;](https://postgr.es/c/476b9eb3c) [&sect;](https://postgr.es/c/18aca7257) [&sect;](https://postgr.es/c/e509be9ce) [&sect;](https://postgr.es/c/37b6e2657) [&sect;](https://postgr.es/c/83bd1c1ba) [&sect;](https://postgr.es/c/051db7737) [&sect;](https://postgr.es/c/3bfc7c849)

   The original implementation of join removal in the planner has had numerous bugs over the years, including some that were never fixed. It was inherently fragile because changes in the planner's derived data structures often required adjustments in join removal, yet there was no easy way to keep that code in sync. Replace it with code that modifies the far-simpler parse tree representation and then loops around to re-generate all of the derived data. This is a significant change that normally we'd never consider back-patching, but the prospect of continuing bugs from the old implementation has driven us to take the risk of doing so.
-  Fix assorted planner bugs triggered by lateral references (Richard Guo) [&sect;](https://postgr.es/c/cbec57291) [&sect;](https://postgr.es/c/30baaacdd) [&sect;](https://postgr.es/c/b7866c934) [&sect;](https://postgr.es/c/d8a311e8d) [&sect;](https://postgr.es/c/96c869df7) [&sect;](https://postgr.es/c/7432ecba1) [&sect;](https://postgr.es/c/707b3738d) [&sect;](https://postgr.es/c/01a9e573f) [&sect;](https://postgr.es/c/d7938ff30) [&sect;](https://postgr.es/c/77c4a9037) [&sect;](https://postgr.es/c/17b45abef)

   These mistakes could lead to incorrect query plans causing wrong answers, or to internal planner errors such as “failed to build any N-way joins”.
-  Ensure proper placement of qual clauses containing PlaceHolderVars by assigning a relation ID to antijoins made from SubLinks (Richard Guo) [&sect;](https://postgr.es/c/64fb20a52)

   This avoids edge cases where such a qual could be pushed into the right-hand side of the antijoin, giving it the opposite of the desired semantics.
-  Fix placement of PlaceHolderVars as nestloop parameters (Richard Guo) [&sect;](https://postgr.es/c/e84fb3f1c) [&sect;](https://postgr.es/c/e0f5f79d2)

   In rare cases, planning could fail with “failed to assign all NestLoopParams to plan nodes”, “plan should not reference subplan's variable”, or “variable not found in subplan target list” errors.
-  Avoid planner failure for sort keys that are PlaceHolderVars (Richard Guo) [&sect;](https://postgr.es/c/5fb5bef3e) [&sect;](https://postgr.es/c/f6642a97f)

   Depending on the plan shape selected, such cases could produce wrong results, or could fail with “variable not found in subplan target lists”.
-  Fix incorrect dropping of outer join clauses after re-associating outer joins (Richard Guo, Nitin Motiani) [&sect;](https://postgr.es/c/27a53c857)
-  Avoid “unsupported join alias expression” failures when a whole-row Var references a join and appears below an outer join (Richard Guo) [&sect;](https://postgr.es/c/53b058d3b)
-  Handle SubLinks correctly when expanding a whole-row join alias reference (Richard Guo) [&sect;](https://postgr.es/c/7a776fc9d)

   This oversight led to errors like “cannot handle unplanned sub-select” or “unrecognized node type”.
-  Do not perform self-join elimination with inheritance trees (Tender Wang) [&sect;](https://postgr.es/c/ebf82fae6)

   In traditional table inheritance, a unique index on the parent table isn't sufficient to prove uniqueness across its children too, so that the proof required for SJE isn't possible. Moreover, the code was willing to merge `parent_table` with `parent_table ONLY`, which doesn't even represent the same set of rows.
-  In self-join elimination, check that clauses used to prove uniqueness actually match (Tender Wang) [&sect;](https://postgr.es/c/10b2d9986)

   Previously, this code assumed that equality clauses with the same operands were equivalent; but in fact they could use different equality operators, or more likely different collations.
-  Don't assume that `DISTINCT ON` ensures uniqueness of a subquery's output rows when the subquery's targetlist contains any set-returning functions (Richard Guo) [&sect;](https://postgr.es/c/4c66f172a)

   This doesn't follow, because the set-returning functions may be evaluated after the de-duplication step.
-  Fix incorrect constant-folding of a `JSON` constructor with a `RETURNING` clause inside a simple-form `CASE` with a constant test expression (Richard Guo) [&sect;](https://postgr.es/c/29aef0f68)
-  Fix test for pushing `WHERE` clauses down past a grouping step that has a different equivalence rule (Richard Guo) [&sect;](https://postgr.es/c/dd7c135bb)

   The previous release cycle's fix for this situation misbehaved for some clauses involving `CASE` expressions.
-  Fix incorrect behavior of `MIN()`/`MAX()` index scans when query specifies `FETCH ... WITH TIES` (Chaitanya Choudhary) [&sect;](https://postgr.es/c/720d3df25)

   The `WITH TIES` option was accidentally copied into the sub-SELECT generated to perform the index scan, possibly triggering “more than one row returned by a subquery used as an expression” errors.
-  Use the correct collation when unique-ifying a semijoin's righthand side (Andrey Rachitskiy) [&sect;](https://postgr.es/c/1ac292cb1)

   This is an issue only if the join's comparison clause uses a non-deterministic collation, since all deterministic collations agree on string equality. But if it does, it was possible to produce incorrect results with too many rows.
-  Fix handling of disabled plan node types below single-child Append nodes (Tender Wang) [&sect;](https://postgr.es/c/718f28167)

   The Append node failed to copy its child's disabled_nodes count, allowing it to possibly be chosen over other alternatives that contain no disabled nodes.
-  Acquire a predicate lock on conflicting rows during `INSERT ... ON CONFLICT` (Zsolt Parragi, Andrey Borodin) [&sect;](https://postgr.es/c/6567f3f1f) [&sect;](https://postgr.es/c/558c00f96)

   Failure to do this resulted in incorrect behavior in serializable mode, allowing concurrent updates that should have failed with a serialization error.
-  Fix integer overflow during parallel GIN index build (Ewan Young) [&sect;](https://postgr.es/c/937a0e68c)

   Index keys were stored in an interim data structure whose size field was only 16 bits, so that keys larger than 65535 bytes were truncated, resulting in crashes or corrupted index entries.
-  Fix incorrect multi-column `RANGE` partition pruning (David Rowley) [&sect;](https://postgr.es/c/7c503d683)

   When performing partition pruning with a `RANGE` partitioned table where pruning quals are only present for a leading prefix of the partition key, it was possible for partition pruning to prune away some partitions that shouldn't be pruned (leading to missing some rows that should be found) and/or include some partitions that were not needed.
-  Fix incorrect pruning of `DEFAULT` partitions in range-partitioned tables that use `MINVALUE` or `MAXVALUE` bounds (Ewan Young) [&sect;](https://postgr.es/c/63d9e26c2)

   In some cases the code would incorrectly conclude that the `DEFAULT` partition needn't be scanned, causing failure to find some rows that should match the query.
-  Fix mishandling of run-time partition pruning during EPQ rechecks (David Rowley, Zsolt Parragi) [&sect;](https://postgr.es/c/c7328770b) [&sect;](https://postgr.es/c/57193b3a0)

   This error could result in silently skipping rows that should be updated, as well as possible use-after-free accesses.
-  Don't use SubPlan-containing expressions in plan-startup-time partition pruning (Richard Guo) [&sect;](https://postgr.es/c/70ffc30e1)

   This avoids “SubPlan found with no parent plan” errors in edge cases.
-  Fix problems with queries applying multiple sort keys to a single column (David Rowley, Ayush Tiwari) [&sect;](https://postgr.es/c/bfbb3d842) [&sect;](https://postgr.es/c/744ca40cb)

   Some queries could crash the backend, while in other cases the additional sort specifications were ignored. Ordinarily such additional sort specifications are redundant anyway, since the first one would fully determine the result ordering. However, this is not necessarily so with non-deterministic collations, or with queries using different `ORDER BY` operators.
-  Avoid possible re-use of stale hash tables when a hashed SubPlan is referenced in multiple places (Rahul Yadav) [&sect;](https://postgr.es/c/5b730e81f) [&sect;](https://postgr.es/c/9d0081c01)
-  Ensure that Memoize plan nodes use the correct input collation (Shihao Zhong) [&sect;](https://postgr.es/c/9dd027535)

   This matters only when using non-deterministic collations.
-  Fix wrong results from Memoize plans when a Param is a subexpression of a cache key (Richard Guo) [&sect;](https://postgr.es/c/dbe6cf558) [&sect;](https://postgr.es/c/2bcf2fa49)
-  Fix loss of parallel worker statistics for Memoize plan nodes underneath a Gather (Tatsuya Kawata) [&sect;](https://postgr.es/c/0aba37abf)

   Only the statistics since the last rescan of the Gather were reported by `EXPLAIN ANALYZE`.
-  Don't create a shell type if a function definition specifies an unknown return type name with array or type modifier syntax (Heikki Linnakangas) [&sect;](https://postgr.es/c/b90eeb29e)

   Such cases are surely errors, since you should create the type before doing anything advanced with it.
-  Fix snapshot export and import to work properly on a snapshot taken during recovery (Peter Geoghegan, Bertrand Drouvot) [&sect;](https://postgr.es/c/74c9832ca)

   The data layout is different in such a snapshot, but the export/import logic didn't know that, so that the imported snapshot behaved incorrectly. This led not only to wrong answers within the importing transaction, but possibly to persistent data corruption due to it setting hint bits incorrectly.

   This bug is known to be triggered by parallel pg_dump, when run on a standby server. In that context it could corrupt both the dump output and the standby server's data.
-  Fix atomicity problem in transaction snapshot import (Chee Wooson) [&sect;](https://postgr.es/c/3977cde27)

   There was a window in which onlookers might not see the source transaction as still running, possibly allowing wrong decisions by `VACUUM` for example.
-  Fix import of statistics for columns that are domains over range types, multirange types, or `tsvector` (Michael Paquier) [&sect;](https://postgr.es/c/753057e34)

   The restore logic failed to look through the domain while checking applicability of statistics types that are for these data types.
-  Fix double parse transformation of `PASSING` arguments in `JSON_TABLE` (Ayush Tiwari, Amit Langote) [&sect;](https://postgr.es/c/6cb715e2d)

   This mistake caused unexpected failures such as “unexpected non-SELECT command in SubLink”.
-  Fix failure to assign collations to `PASSING` expressions in `JSON_TABLE` (Venu Kummari) [&sect;](https://postgr.es/c/c7656ac9e)

   This oversight could result in wrong answers or “could not determine which collation to use for string comparison” errors.
-  Fix usage of toasted text values in the `PASSING` arguments of `JSON_TABLE` and related functions (Chaitanya Choudhary, Shihao Zhong) [&sect;](https://postgr.es/c/b8736db6e)

   The code failed to apply detoasting, resulting in incorrect output.
-  Fix NULL result persistence in `JSON_VALUE() RETURNING` `json`/`jsonb` (Srinath Reddy Sadipiralla) [&sect;](https://postgr.es/c/7150d0fd4)

   Once `JSON_VALUE(... RETURNING ...)` with a `json` or `jsonb` result value returned NULL, every later evaluation of the same expression in the same statement also returned NULL, whether it should or not.
-  Fix `ON EMPTY/ERROR` result persistence in `JSON_VALUE` and related constructs (Jian He, Andrey Rachitskiy) [&sect;](https://postgr.es/c/52033b887)

   A NULL input could result in returning the value of the previous row's `ON EMPTY / ON ERROR` expression.
-  Fix handling of failures in `DEFAULT` expressions in `SQL/JSON` operations (Jian He) [&sect;](https://postgr.es/c/5209a736a)

   An error arising while executing a `DEFAULT` expression could be ignored, instead producing a NULL result.
-  Fix handling of domain coercion failures in `SQL/JSON` operations (Jian He) [&sect;](https://postgr.es/c/b9a415360)

   `SQL/JSON` query functions could throw an error, or report the wrong one, in cases where an `ON ERROR` or `ON EMPTY` clause specified some non-error behavior.
-  Fix GIN index searches for plain scalar strings in strict `jsonpath` queries (Rui Zhao) [&sect;](https://postgr.es/c/65f3b7096)

   A `jsonb_ops` index scan for a `jsonpath` such as `strict $ == "scalar"` failed to find table entries that were simply `"scalar"`.
-  Fix errors in extraction of GIN index scan conditions from jsonpath queries (Tom Lane) [&sect;](https://postgr.es/c/154e2f101)

   An `AND` operator underneath a `NOT` operator could be misimplemented. Also a method applied to a literal constant, for example `"42".double()`, was not handled correctly. These mistakes did not cause visible errors, but the GIN index search would miss some or all of the rows that should match.
-  Fix edge-case error in SP-GiST quad-tree `<@` searches (John Naylor) [&sect;](https://postgr.es/c/03ba8e591)

   Due to roundoff error, the search might sometimes not descend into a quadrant that it needs to search, leading to incomplete query results.
-  Fix edge-case failures in SP-GiST quad-tree indexes (Ayush Tiwari) [&sect;](https://postgr.es/c/fc834ccab)

   Roundoff errors could allow searches and insertions to fail with “getQuadrant: impossible case” errors.
-  Accept empty `tsquery` values in binary input (Ayush Tiwari) [&sect;](https://postgr.es/c/6c4172b34)

   An empty `tsquery` value is allowed elsewhere, but `tsqueryrecv()` failed.
-  Disallow zero-length lexemes and out-of-range phrase distance values in binary input of `tsquery` values (Tom Lane) [&sect;](https://postgr.es/c/69915dac4)

   These cases were already rejected during text input, but `tsqueryrecv()` omitted making equivalent checks. That allowed accepting values that would fail subsequent dump/reload.
-  Treat operators-nested-too-deeply errors as soft errors in `tsquery` input and `ltxtquery` input (Tom Lane) [&sect;](https://postgr.es/c/90cc09f57)
-  Consistently enforce maximum lengths in `tsvector` and `tsquery` (Tom Lane) [&sect;](https://postgr.es/c/a68beda42)

   Some places thought the maximum lexeme length was `MAXSTRLEN` (2047), while others thought it was `MAXSTRLEN-1`. Fortunately, the looser bound is correct, so make them all use `MAXSTRLEN`. Likewise, some places were stricter than others about the total data length limit, `MAXSTRPOS`. These fixes avoid situations where a validly-constructed value would fail dump/reload.
-  Fix `right()` with a length argument of `INT_MIN` (-2147483648) (Ewan Young) [&sect;](https://postgr.es/c/bfb08d094)

   Due to integer overflow, the function returned the whole string when it should return an empty string.
-  Fix `pg_blocking_pids()` to account correctly for relation extension locks (Bertrand Drouvot) [&sect;](https://postgr.es/c/dc60d6aa1)
-  Fix 32-vs-64-bit confusion in `pg_indexam_progress_phasename()` (Markos Fountoulakis) [&sect;](https://postgr.es/c/58998c33d)

   This function accidentally worked on common hardware, but could misbehave on 32-bit or big-endian platforms.
-  Detect integer overflow in `to_char()` with `V` format (Andrey Rachitskiy) [&sect;](https://postgr.es/c/ec434db1f)

   If the data value is an integer, and multiplying it by the power of ten implied by the `V` format causes integer overflow, `to_char()` silently emitted incorrect digits. Throw an integer overflow error instead.
-  Avoid integer overflow while calculating BRIN minmax_multi distances for `tid` values (Tomas Vondra) [&sect;](https://postgr.es/c/1be7010d1)

   Very large `tid` values could result in negative distances, which would end in making the index inefficient, though not actually corrupt.
-  Avoid overflow while calculating BRIN minmax_multi distances for numeric values, and handle NaNs properly (Kirill Reshke, Tomas Vondra) [&sect;](https://postgr.es/c/5c8629bb5)

   These oversights could result in errors being thrown (preventing index summarization), or in inefficient indexes.
-  Avoid overflow when computing differences for numeric ranges (Kirill Reshke) [&sect;](https://postgr.es/c/8a043da94)

   In this context we want to return plus or minus infinity rather than throwing an overflow error.
-  Support out-of-order `FLAG` lines in Hunspell affix files (Ewan Young) [&sect;](https://postgr.es/c/1f02a1e48)

   Hunspell affix files usually put the `FLAG` line first, and our implementation silently depended on that. If it came after some lines whose meaning depends on the `FLAG` setting, the result would be crashes or failures to match compound words. Fix the logic so that the `FLAG` line can be anywhere.
-  Fix `OLD`/`NEW` aliases that `EXPLAIN` prints in `RETURNING` SubPlans (Henson Choi) [&sect;](https://postgr.es/c/4ee8af2e8) [&sect;](https://postgr.es/c/b049b4e7b)

   In plans involving multiple aliases, the wrong one could be printed, resulting in confusing output.
-  Fix assorted errors in rule deparsing (Henson Choi) [&sect;](https://postgr.es/c/e79df9dd9) [&sect;](https://postgr.es/c/eb1454ecc) [&sect;](https://postgr.es/c/b08c4df3b) [&sect;](https://postgr.es/c/17a0604c2)

   Always print `AS` before a target-relation alias in `UPDATE`, `DELETE`, or `MERGE`, to avoid ambiguity if that alias is `set`. Also ensure that <em>xml</em><code> IS DOCUMENT</code> constructs and subquery left-hand arguments are adequately parenthesized. Also fix deparsing of `AT TIME ZONE` and similar constructs in a view's `FROM` clause; these need to be wrapped in a no-op `CAST`. These errors could prevent dump restoration due to a syntax error in a view, or cause reloaded views to be parsed as something different from before.
-  Fix parser failure when a multi-row `INSERT ... VALUES` has a sub-SELECT in a target column subscript (Amit Langote) [&sect;](https://postgr.es/c/fa0806647)

   This led to “unexpected non-SELECT command in SubLink”.
-  Fix missed preprocessing of `WHERE` clauses in `COPY FROM` (Nathan Bossart) [&sect;](https://postgr.es/c/95d960447)

   This oversight led to “cache lookup failed for function 0” errors with negated operators, and probably in other cases as well.
-  Disallow `NOT NULL NO INHERIT` constraints on partitioned tables (David Rowley, Nitin Motiani) [&sect;](https://postgr.es/c/35f60e5ce)

   This was already rejected by plain `CREATE TABLE`, but cases such as `CREATE TABLE LIKE` missed checking copied constraints, potentially allowing a partition to contain null values when it should not.

   If you suspect this might have happened in your database, issue this command to find broken constraints:

```sql

SELECT conrelid::regclass, conname
FROM pg_constraint
WHERE contype = 'n' AND connoinherit
      AND conrelid IN (SELECT oid FROM pg_class WHERE relkind = 'p');
```
   Unwanted constraints can be removed with `ALTER TABLE ... DROP CONSTRAINT ...`, while ones that are wanted should be repaired with `ALTER TABLE ... ALTER CONSTRAINT ... INHERIT`. The latter command will fail if any null values are present in the table's partitions.
-  Fix test for an existing NOT NULL constraint in `ALTER TABLE ADD PRIMARY KEY` (Fujii Masao, Manuel Reyes Bravo) [&sect;](https://postgr.es/c/5659e549a)

   Previously, dropping a NOT NULL column, adding a new column with the same name, and making it a primary key in a single `ALTER TABLE` command could fail with “primary key column ... is not marked NOT NULL”.
-  Fix concurrent delete handling in `MERGE ... WHEN NOT MATCHED BY SOURCE` (Tender Wang) [&sect;](https://postgr.es/c/ce9bd103e)

   When executing a `WHEN NOT MATCHED BY SOURCE` action in `MERGE`, if the tuple from the target table was concurrently deleted, the retry code would incorrectly execute a `WHEN NOT MATCHED BY TARGET` action for the row instead, giving incorrect results.
-  Fix race conditions with `DROP TABLESPACE` (Ayush Tiwari) [&sect;](https://postgr.es/c/d77627daa)

   A `DROP` executed concurrently with commands that attempt to add new objects to the tablespace could end with inconsistent catalog contents.
-  Cleanly reject operations on unlogged tables during recovery (Jim Jones) [&sect;](https://postgr.es/c/794c29dc4) [&sect;](https://postgr.es/c/a5b925cfc)

   An unlogged table will have no data on a standby, so reading it must fail. Most operations give an appropriate error for this, but `COPY ... TO` failed with an internal “could not open file” error, as did `currtid2()`.
-  Fix inconsistent permissions checking for role membership changes (ChangAo Chen) [&sect;](https://postgr.es/c/0468246f5)

   This fixes cases where what should be a permissions failure was reported as an internal “no possible grantors” error.
-  Revert previous changes to free space map persistence (Alexander Korotkov) [&sect;](https://postgr.es/c/5d63dbd97)

   Commits c06d1a4ba et al. (included in the May 2026 releases) turned out to cause more trouble than they fixed.
-  Avoid backend hang during temporary table cleanup at exit, if the default transaction mode is `SERIALIZABLE READ ONLY DEFERRABLE` (Andrey Rachitskiy) [&sect;](https://postgr.es/c/d2943b94b)

   Cleanup would block until concurrent serializable transactions complete, which could be indefinitely if some of them are prepared transactions.
-  Avoid crash when describing a `FETCH` statement whose cursor is gone (Dirkjan Bussink) [&sect;](https://postgr.es/c/95f6bfb51)
-  Fix memory accounting for datum sorts on pass-by-reference data types (Mario Karuza) [&sect;](https://postgr.es/c/8d5b228ae)

   This error could lead to failure to enforce `work_mem` limits correctly, or to misleading output from `EXPLAIN ANALYZE`.
-  If a vacuum delay sleep is requested while holding locks, don't sleep (Kevin Rocker, Andrey Borodin, Mostafa) [&sect;](https://postgr.es/c/5f671c431)

   Also, move several existing delay calls to points where it's actually safe to sleep.
-  Fix parallel vacuum I/O timing reports (Bharath Rupireddy) [&sect;](https://postgr.es/c/6bfb33088)

   `VACUUM VERBOSE` and autovacuum logs included parallel workers' blocks-read and blocks-dirtied counts, but not the time they spent doing that I/O.
-  Avoid freeing the plan for a foreign-key enforcement query if the plan is still in use (Lucas Jeffrey, Trakshan Mishra, Amit Langote) [&sect;](https://postgr.es/c/9e9d24c33)

   It was possible to reach this behavior in contrived cases, such as a self-referencing ON DELETE CASCADE foreign key on tables with triggers that perform DML.
-  Fix postmaster failing to exit when the startup subprocess crashes during crash restart (Zexin Li) [&sect;](https://postgr.es/c/6a1759a02)

   The postmaster should give up and quit in this case, but it failed to do so, instead becoming stuck.
-  Be more wary about scanner failures when parsing a GUC file (Michael Paquier) [&sect;](https://postgr.es/c/0391fee2d)
-  Recheck `oauth_validator_libraries` during OAuth authentication (Zsolt Parragi) [&sect;](https://postgr.es/c/8e54a4145) [&sect;](https://postgr.es/c/dcf112bd4)

   This addition ensures that changes to `oauth_validator_libraries` are enforced immediately. Previously, if such a change made an existing `oauth` line in `pg_hba.conf` invalid, the new setting was not enforced until that situation was fixed.
-  Fix postmaster crash when `oauth_validator_libraries` is set to a string containing only whitespace (Grigorev Jurij, Zsolt Parragi) [&sect;](https://postgr.es/c/99c3f5e20) [&sect;](https://postgr.es/c/459ef520d) [&sect;](https://postgr.es/c/af28fc13e) [&sect;](https://postgr.es/c/e83298599)
-  Fix WAL-logging of multi-level page deletion in btree indexes (Peter Geoghegan) [&sect;](https://postgr.es/c/e88259b67)

   The WAL log could contain stale values for a to-be-deleted page's sibling links, resulting in complaints on standby servers if `wal_consistency_checking` is enabled. Barring a crash, the damage would shortly be repaired by completion of the deletion action, so that this error was usually only cosmetic.
-  Avoid overly-strong WAL consistency check on buffers that are only pinned not locked (Greg Burd, Thomas Munro) [&sect;](https://postgr.es/c/7ac2f74d1)

   A buffer that is not locked could be concurrently changed by other processes, so that the image captured for consistency checking is not self-consistent, leading to “incorrect resource manager data checksum in record” failures on standbys.
-  Ensure that state files for pending prepared transactions are written out atomically during a checkpoint (Shihao Zhong) [&sect;](https://postgr.es/c/606f3db55)

   Previously it was possible for a crash to leave an invalid state file behind, causing subsequent database startup to fail.
-  Do not start a checkpoint in the middle of aborting a prepared transaction (Shihao Zhong) [&sect;](https://postgr.es/c/3a2fc7662)

   A crash would allow the transaction to be committed afterwards, leading to inconsistent state.
-  Do not start a checkpoint in the middle of creating a multixact (Shihao Zhong) [&sect;](https://postgr.es/c/31538e165)

   This avoids possible data loss after a subsequent crash.
-  Fix possible loss of tuples updated by later-aborted subtransactions (Shihao Zhong, Michael Paquier) [&sect;](https://postgr.es/c/5df84b355)

   A table rewrite (for example, `CLUSTER` or `VACUUM FULL`) could discard such tuples as invalid.
-  Do not lose tuple's `tid` when storing it into a tuple table slot (Virender Singla, Greg Burd) [&sect;](https://postgr.es/c/badcb62d6)

   Some code paths neglected to copy the `tid` field, which could lead to data corruption or wrong query answers in unusual cases.
-  Avoid race condition in marking a partitioned index valid (Andrey Borodin, Kirill Reshke) [&sect;](https://postgr.es/c/1735ad557)

   This could, for example, cause parallel pg_restore to fail due to trying to recreate foreign key constraints too soon.
-  Fix checks for an aborted updating transaction in multixids (Chee Wooson) [&sect;](https://postgr.es/c/94d53e9ec)

   Known consequences of this coding error include unexpected “new multixact has more than one updating member” errors, as well as incorrect failures to update a row.
-  A hash index scan that starts during a bucket split could miss some index entries (Peter Geoghegan, Amit Kapila) [&sect;](https://postgr.es/c/a85b5523e)
-  Fix incorrect position tracking in TID Range Scans (Ewan Young, David Rowley) [&sect;](https://postgr.es/c/39a0db101)

   Sequences involving fetching backwards and then forwards again could fail to scan all the way to the end of the range.
-  Avoid unnecessary buffer locking when recovery replays freeze-only and VM-only WAL records (Melanie Plageman) [&sect;](https://postgr.es/c/9733eb9e8)

   This prevents unnecessary replay delays and cancellations of standby queries.
-  Avoid reading a potentially-torn copy of the `pg_control` file for base backups (David Steele) [&sect;](https://postgr.es/c/8cddd5fc8) [&sect;](https://postgr.es/c/4f857f399)

   A concurrent write to the file could result in including inconsistent data in the backup.
-  Prevent making incorrect copies of non-relation files in base backups (David Steele) [&sect;](https://postgr.es/c/b09b26ac8)

   Files that routinely get replaced by atomic renames, such as `postgresql.auto.conf`, could be truncated or zero-padded in a backup made concurrently with such a replacement.
-  Ensure that base backups report checksum failures affecting single pages (Zsolt Parragi) [&sect;](https://postgr.es/c/39faf4de2)

   Due to a logic error, checksum failures were counted in `pg_stat_database` only when more than one page in the same file was corrupt.
-  Fix stale visibility maps in incremental backups after heap truncation (Ke, Fujii Masao, Shihao Zhong) [&sect;](https://postgr.es/c/41b4b8e8b)

   An incremental backup taken after `VACUUM` truncates a table can restore stale visibility map bits when both data checksums and wal_log_hints are disabled, leading to data corruption.

   This fix does not correct existing WAL summaries, therefore incremental backups taken after updating can still be affected. It's recommended that users of incremental backups take a new full backup after updating and use it as the basis for subsequent incremental backups.
-  Reject incremental backups based on a backup from a divergent history (Sai Pranav, Fujii Masao) [&sect;](https://postgr.es/c/2afbad0ce)

   After a failover, an incremental backup could accept a prior backup whose WAL range extends beyond the point where the new primary left the old timeline. Combining these backups could retain pages from the abandoned history, producing a corrupt backup.
-  Ensure that incremental backup merges WAL summaries in LSN order (Ke, Fujii Masao) [&sect;](https://postgr.es/c/9e20db542)

   This error could result in stale data being included in a combined backup.
-  Fix erroneous edge-case decisions about when to include WAL summaries in incremental backups (Ayush Tiwari) [&sect;](https://postgr.es/c/5787f8785)

   This error could result in stale data being included in a combined backup.
-  Fix incorrect size estimates for incremental backups (Daria Lepikhova) [&sect;](https://postgr.es/c/f47c63d7a)

   This problem affected progress reporting in pg_basebackup and other places.
-  Avoid resetting a replication subscription's `two_phase` state to pending if it is already enabled (Zhijie Hou) [&sect;](https://postgr.es/c/fbaebf450)

   Executing `ALTER SUBSCRIPTION ... SET (two_phase = true)` twice resulted in confused state, potentially allowing data corruption on the subscriber.
-  Verify the syntax of the `output_plugin_libraries` parameter when a value is assigned (Jacob Champion, Hayato Kuroda) [&sect;](https://postgr.es/c/e2b3549a4)

   Previously, a syntactically-incorrect value could be put into `postgresql.conf`, causing logical replication to fail.
-  Reject `REPLICA IDENTITY USING INDEX` if a key column has an unvalidated `NOT NULL` constraint (Ante Kresic) [&sect;](https://postgr.es/c/864ef0fb1)

   A replica identity key must be not-null, but this code path would accept cases where that was not yet known to be true.
-  Fix logical replication with a lossy identity index (Paul A. Jungwirth, Zhijie Hou) [&sect;](https://postgr.es/c/23fffd382)

   Typically identity indexes are btrees, which deliver exact results. However, GiST indexes on multiranges can be used as identity indexes, and they deliver results that require rechecking. Logical replication failed to perform the recheck, so that updates and deletes could target the wrong row on the subscriber.
-  Avoid incorrect update of backend status flags during replication slot invalidation (Vlad Lesin) [&sect;](https://postgr.es/c/c5666f7e1)

   An invalidation done by an auxiliary process such as the checkpointer could clobber the status field of an unrelated backend. The consequences of this are difficult to predict, but one possibility is that hot standby feedback could be effectively disabled.
-  In logical replication, handle subtransaction aborts better (Bingshuai Li, Hayato Kuroda, Zhijie Hou) [&sect;](https://postgr.es/c/5eae2800a)

   If a subtransaction that made a catalog update aborts, logical decoding mistakenly still believed that the update was effective, potentially causing incorrect decoding output.
-  Avoid race condition when building an initial decoding snapshot (Antonin Houska, Rui Zhao) [&sect;](https://postgr.es/c/6210f00ca)

   There was a small risk of data corruption from using a snapshot that includes transactions that have committed but not yet set their CLOG status.
-  In logical decoding, avoid processing a stream abort event for a subtransaction that was never actually streamed (Andrey Rachitskiy) [&sect;](https://postgr.es/c/caa463e3f)

   This error could lead to null pointer dereferences or other misbehavior.
-  Avoid null-pointer crash with missing replication option values in `pgoutput` logical replication sender (Zhijie Hou) [&sect;](https://postgr.es/c/6437d9774)
-  Fix small memory leak in pgoutput with streamed prepared transactions (Ayush Tiwari) [&sect;](https://postgr.es/c/a477847eb)
-  Fix lost wakeups in physical WAL senders (Alex Speller) [&sect;](https://postgr.es/c/1ff9f9380)

   This fix corrects some cases where a WAL sender process could sleep even though there is data waiting to be sent.
-  Fix duplicate creation of temporary replication slot when a WAL receiver follows a timeline switch (ChangAo Chen) [&sect;](https://postgr.es/c/2bf6144b4)

   This mistake led to a “could not create replication slot” error followed by a restart of the WAL receiver, which would then recover, so there was no permanent damage.
-  Fix race condition between logical replication subscriber and concurrent index DDL (Mikhail Nikalayeu, Vignesh C) [&sect;](https://postgr.es/c/af48a6702)

   Changing the state of a replication target relation's primary key or replica identity index could cause the replication apply logic to miss finding a matching row (but only once per such event).
-  Fix incorrect LSNs in WAL decoding error messages (Rahul Yadav) [&sect;](https://postgr.es/c/092da6511)

   These error messages reported the LSN of a WAL record that wasn't the one they were actually complaining about.
-  Don't allow `REPLICA IDENTITY FULL` lookups to select invalid indexes (Mikhail Nikalayeu) [&sect;](https://postgr.es/c/38eb0cfa3)

   Previously, logical replication could try to use an index left behind by a failed `CREATE INDEX CONCURRENTLY`, leading to wrong results since the index is likely incomplete.
-  Fix potential use-after-free cases in WAL sender (Ayush Tiwari) [&sect;](https://postgr.es/c/d21a3e1fe)
-  Avoid use-after-free in `jsonb_set_lax()` (John Naylor) [&sect;](https://postgr.es/c/52af12e9d)
-  Fix double-free when `XMLSERIALIZE(... INDENT)` output exceeds 1GB (Jim Jones) [&sect;](https://postgr.es/c/f10500777)
-  Fix null-pointer dereferences that could occur when a schema is concurrently dropped (Baji Shaik) [&sect;](https://postgr.es/c/d3a03158c)
-  Avoid null-pointer crash in logical subscription refresh when a relevant relation is concurrently dropped (Satyanarayana Narlapuram, Bharath Rupireddy) [&sect;](https://postgr.es/c/ec074c41e)
-  Fix possible null-pointer dereference when building extended statistics on expressions (Shihao Zhong, Michael Paquier) [&sect;](https://postgr.es/c/4bcd0d31c)

   This error was reachable only with unexpected (but legal) behavior of a custom `typanalyze` function.
-  Fix possibly-stale values of `pg_locks`.`waitstart` (Shihao Zhong) [&sect;](https://postgr.es/c/ccffd3439)
-  Harden pgstats code against partially-completed local pgstats entries (Michael Paquier) [&sect;](https://postgr.es/c/e99e00b27)
-  Fix cleanup of reused pgstats entries (Michael Paquier) [&sect;](https://postgr.es/c/2c056b37f)

   This error led to leakage of pgstats entries, which could result in errors at backend exit or database shutdown.
-  Fix memory leaks caused by not releasing all data associated with a WAL reader (Tomás Senart) [&sect;](https://postgr.es/c/7418d9e96)

   In many cases this oversight did little harm because the containing memory context would soon be reset or deleted. However there are cases where it could cause long-lived memory bloat, notably in the checkpointer process.
-  Avoid memory bloat while computing `ndistinct` extended statistics across many columns (Ilia Evdokimov) [&sect;](https://postgr.es/c/942266c7f)
-  Fix buffer pin leak after a failed hash index bucket split cleanup (Peter Geoghegan) [&sect;](https://postgr.es/c/f4284d9aa)

   This coding error resulted in nothing worse than warnings about resource leakage.
-  Ensure all relations opened for `AFTER` triggers are closed again (Hayato Kuroda, David Rowley) [&sect;](https://postgr.es/c/715d83971)

   This oversight triggered “resource was not closed” warnings, but nothing more serious than that.
-  Fix `io_uring` resource leakage across crash restarts (Lucas Draescher) [&sect;](https://postgr.es/c/5bfe88cf6)

   Eventually, a restart would fail with “too many open files” or a similar error, at which point the postmaster would exit and the resources would be released. This is only an issue when using `io_method=io_uring`, which is not the default.
-  Fix PL/pgSQL's handling of `anycompatiblemultirange` arguments and results (Jeevan Chalke) [&sect;](https://postgr.es/c/a34d61616)

   This type had not been added to PL/pgSQL's knowledge of polymorphic types, resulting in errors when trying to create or use a function declared with it.
-  Harden PL/Perl and associated contrib modules against “tied” Perl hashes and arrays (Tom Lane) [&sect;](https://postgr.es/c/af659efcb) [&sect;](https://postgr.es/c/b6193ff03)

   Rely on `av_count()` to decide the length of a Perl array, rather than looking for a null pointer; and contrarily be more wary about unexpected null value pointers from array and hash fetches. These cases are only known to arise in connection with “tied” objects.
-  Avoid memory leakage in libpq with repeated GSS connections (Matt Carter, Tom Lane) [&sect;](https://postgr.es/c/94d625f5d)
-  Avoid incorrect behavior when `getopt_long()` rearranges a program's command line arguments (Sehrope Sarkuni) [&sect;](https://postgr.es/c/d33449024)

   For example, `vacuumdb postgres --jobs` took `postgres` as the number of jobs instead of complaining that `--jobs` is missing its argument.
-  Fix psql's display of `REPLICA IDENTITY NOTHING` table properties (Shinya Kato) [&sect;](https://postgr.es/c/baa7b142a)
-  Ensure that psql counts all the `COPY FROM STDIN` commands appearing in a multi-query command string (Zsolt Parragi) [&sect;](https://postgr.es/c/bcb776671)

   After the fix for CVE-2026-6464, psql would reject multiple `COPY FROM STDIN` commands appearing in a single string, because it was not expecting more than one `COPY` operation to occur.
-  Prevent psql from getting confused if the argument of `\getresults` is invalid (Fujii Masao) [&sect;](https://postgr.es/c/fd9bb365d)
-  Make psql cope better with out-of-memory conditions during backslash command processing (Fujii Masao) [&sect;](https://postgr.es/c/3f01b0c4a) [&sect;](https://postgr.es/c/789ae09a5)
-  Fix some small memory leaks in psql (Fujii Masao) [&sect;](https://postgr.es/c/8753b6904)
-  In pg_combinebackup, reject out-of-range block numbers in incremental backup files (Peter Eisentraut) [&sect;](https://postgr.es/c/15e940a10)

   This omission meant that pg_combinebackup could crash on corrupted input.
-  Make pg_dump dump column statistics targets for indexes backing constraints (Zsolt Parragi) [&sect;](https://postgr.es/c/37bbf5bba)

   Formerly pg_dump would emit `ALTER INDEX ... SET STATISTICS` commands when needed for plain indexes, but not for indexes created via `EXCLUDE` constraints on expressions.
-  In pg_dump's tar-format output, don't put invalid `\restrict`/`\unrestrict` commands into the `restore.sql` tar member (Pavel Kopylov) [&sect;](https://postgr.es/c/2d748cfe3)
-  Make pg_walsummary's `--quiet` switch suppress limit block output (Koshi Shibagaki) [&sect;](https://postgr.es/c/1370a7832)
-  Make reindexdb error out cleanly if an index-level `REINDEX CONCURRENTLY` is requested with multiple jobs (Kirill Reshke) [&sect;](https://postgr.es/c/acccaf0d8)

   Previously, this combination of options would lead to a server-side error, complaining that `REINDEX CONCURRENTLY` cannot run within a transaction block.
-  Fix `contrib/file_fdw` to not crash in `ANALYZE` when the table has column default clauses (Fujii Masao) [&sect;](https://postgr.es/c/813a74bc5)
-  Fix possible index corruption in `contrib/pg_trgm` GiST indexes (Nate Clark) [&sect;](https://postgr.es/c/60ccd730d)

   Combination of index entries could lose the flag identifying their format, so that subsequent index searches could miss matching rows.

   This fix does not repair existing indexes. Users should `REINDEX` existing `gist_trgm_ops` indexes after updating.
-  Fix error check in timer setup on Windows (Sehrope Sarkuni) [&sect;](https://postgr.es/c/4d62ef92f)

   Failure to create a thread to manage timeouts was not detected, leading to silent failure of subsequent timeout requests.
-  Update time zone data files to tzdata release 2026f (Tom Lane) [&sect;](https://postgr.es/c/c45ba888d) [&sect;](https://postgr.es/c/f9d7dee5f) [&sect;](https://postgr.es/c/fd97624fb)

   Canada's Northwest Territories are on year-round UTC-06 (effectively, permanent DST) beginning in November 2026. This affects only America/Inuvik as the rest of the territory is covered by America/Edmonton, for which the equivalent change was released in 2026c. Similarly, Manitoba (America/Winnipeg) moves to year-round UTC-05. For now assume their TZ abbreviations will be CST and EST respectively.

   Cuba did not fall back on 2026-11-01; they will stay on DST this winter (but apparently plan to revert to normal next year).

   The backward-compatibility zone names EST5EDT, CST6CDT, MST7MDT, and PST8PDT now have simpler behavior for old time stamps. They use the indicated UTC offsets and TZ abbreviations indefinitely far back, rather than allowing aspects of the associated city-named zone's historical behavior to leak through. This conforms better to POSIX.

   Historical corrections for Brazil, Colombia, Iran, and Ireland.

   On Windows, initdb will now recognize the new time zone names that Microsoft has devised for the updated Canadian zones.
