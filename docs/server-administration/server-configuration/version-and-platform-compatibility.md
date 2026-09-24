<a id="runtime-config-compatible"></a>

## Version and Platform Compatibility
  <a id="runtime-config-compatible-version"></a>

### Previous PostgreSQL Versions


<a id="guc-array-nulls"></a>

`array_nulls` (`boolean`)
:   This controls whether the array input parser recognizes unquoted `NULL` as specifying a null array element. By default, this is `on`, allowing array values containing null values to be entered. However, PostgreSQL versions before 8.2 did not support null values in arrays, and therefore would treat `NULL` as specifying a normal array element with the string value “NULL”. For backward compatibility with applications that require the old behavior, this variable can be turned `off`.


     Note that it is possible to create array values containing null values even when this variable is `off`.
<a id="guc-backslash-quote"></a>

`backslash_quote` (`enum`)
:   This parameter controls whether a quote mark can be represented by `\'` in the escape string syntax (`E'...'`). The preferred, SQL-standard way to represent a quote mark is by doubling it (`''`) but PostgreSQL has historically also accepted `\'`. However, use of `\'` creates security risks because in some client character set encodings, there are multibyte characters in which the last byte is numerically equivalent to ASCII `\`. If client-side code does escaping incorrectly then an SQL-injection attack is possible. This risk can be prevented by making the server reject queries in which a quote mark appears to be escaped by a backslash. The allowed values of `backslash_quote` are `on` (allow `\'` always), `off` (reject always), and `safe_encoding` (allow only if client encoding does not allow ASCII `\` within a multibyte character). `safe_encoding` is the default setting.


     Note that in an ordinary string literal, `\` just means `\` anyway. This parameter only affects the handling of escape string syntax.
<a id="guc-lo-compat-privileges"></a>

`lo_compat_privileges` (`boolean`)
:   In PostgreSQL releases prior to 9.0, large objects did not have access privileges and were, therefore, always readable and writable by all users. Setting this variable to `on` disables the new privilege checks, for compatibility with prior releases. The default is `off`. Only superusers and users with the appropriate `SET` privilege can change this setting.


     Setting this variable does not disable all security checks related to large objects — only those for which the default behavior has changed in PostgreSQL 9.0.
<a id="guc-quote-all-identifiers"></a>

`quote_all_identifiers` (`boolean`)
:   When the database generates SQL, force all identifiers to be quoted, even if they are not (currently) keywords. This will affect the output of `EXPLAIN` as well as the results of functions like `pg_get_viewdef`. See also the `--quote-all-identifiers` option of [app-pgdump](../../reference/postgresql-client-applications/pg_dump.md#app-pgdump) and [app-pg-dumpall](../../reference/postgresql-client-applications/pg_dumpall.md#app-pg-dumpall).
<a id="guc-standard-conforming-strings"></a>

`standard_conforming_strings` (`boolean`)
:   Beginning in PostgreSQL 19, this parameter is always `on`. String literals are always parsed as specified in the SQL standard (that is, backslashes are ordinary characters within a string literal). This parameter continues to exist because applications may consult it; but it cannot be set to `off`. Escape string syntax ([String Constants with C-Style Escapes](../../the-sql-language/sql-syntax/lexical-structure.md#sql-syntax-strings-escape)) should be used if an application desires backslashes to be treated as escape characters.
<a id="guc-synchronize-seqscans"></a>

`synchronize_seqscans` (`boolean`)
:   This allows sequential scans of large tables to synchronize with each other, so that concurrent scans read the same block at about the same time and hence share the I/O workload. When this is enabled, a scan might start in the middle of the table and then “wrap around” the end to cover all rows, so as to synchronize with the activity of scans already in progress. This can result in unpredictable changes in the row ordering returned by queries that have no `ORDER BY` clause. Setting this parameter to `off` ensures the pre-8.3 behavior in which a sequential scan always starts from the beginning of the table. The default is `on`.
  <a id="runtime-config-compatible-clients"></a>

### Platform and Client Compatibility


<a id="guc-transform-null-equals"></a>

`transform_null_equals` (`boolean`)
:   When on, expressions of the form <em>expr</em><code> =
            NULL</code> (or <code>NULL =
            </code><em>expr</em>) are treated as <em>expr</em><code> IS NULL</code>, that is, they return true if *expr* evaluates to the null value, and false otherwise. The correct SQL-spec-compliant behavior of <em>expr</em><code> = NULL</code> is to always return null (unknown). Therefore this parameter defaults to `off`.


     However, filtered forms in Microsoft Access generate queries that appear to use <em>expr</em><code> = NULL</code> to test for null values, so if you use that interface to access the database you might want to turn this option on. Since expressions of the form <em>expr</em><code> = NULL</code> always return the null value (using the SQL standard interpretation), they are not very useful and do not appear often in normal applications so this option does little harm in practice. But new users are frequently confused about the semantics of expressions involving null values, so this option is off by default.


     Note that this option only affects the exact form `= NULL`, not other comparison operators or other expressions that are computationally equivalent to some expression involving the equals operator (such as `IN`). Thus, this option is not a general fix for bad programming.


     Refer to [Comparison Functions and Operators](../../the-sql-language/functions-and-operators/comparison-functions-and-operators.md#functions-comparison) for related information.
<a id="guc-allow-alter-system"></a>

`allow_alter_system` (`boolean`)
:   When `allow_alter_system` is set to `off`, an error is returned if the `ALTER SYSTEM` command is executed. This parameter can only be set in the `postgresql.conf` file or on the server command line. The default value is `on`.


     Note that this setting must not be regarded as a security feature. It only disables the `ALTER SYSTEM` command. It does not prevent a superuser from changing the configuration using other SQL commands. A superuser has many ways of executing shell commands at the operating system level, and can therefore modify `postgresql.auto.conf` regardless of the value of this setting.


     Turning this setting off is intended for environments where the configuration of PostgreSQL is managed by some external tool. In such environments, a well-intentioned superuser might *mistakenly* use `ALTER SYSTEM` to change the configuration instead of using the external tool. This might result in unintended behavior, such as the external tool overwriting the change at some later point in time when it updates the configuration. Setting this parameter to `off` can help avoid such mistakes.


     This parameter only controls the use of `ALTER SYSTEM`. The settings stored in `postgresql.auto.conf` take effect even if `allow_alter_system` is set to `off`.
