<a id="sql-createdatabase"></a>

# CREATE DATABASE

create a new database

## Synopsis


```

CREATE DATABASE NAME
    [ WITH ] [ OWNER [=] USER_NAME ]
           [ TEMPLATE [=] TEMPLATE ]
           [ ENCODING [=] ENCODING ]
           [ STRATEGY [=] STRATEGY ]
           [ LOCALE [=] LOCALE ]
           [ LC_COLLATE [=] LC_COLLATE ]
           [ LC_CTYPE [=] LC_CTYPE ]
           [ BUILTIN_LOCALE [=] BUILTIN_LOCALE ]
           [ ICU_LOCALE [=] ICU_LOCALE ]
           [ ICU_RULES [=] ICU_RULES ]
           [ LOCALE_PROVIDER [=] LOCALE_PROVIDER ]
           [ COLLATION_VERSION = COLLATION_VERSION ]
           [ TABLESPACE [=] TABLESPACE_NAME ]
           [ ALLOW_CONNECTIONS [=] ALLOWCONN ]
           [ CONNECTION LIMIT [=] CONNLIMIT ]
           [ IS_TEMPLATE [=] ISTEMPLATE ]
           [ OID [=] OID ]
```


## Description


 `CREATE DATABASE` creates a new PostgreSQL database.


 To create a database, you must be a superuser or have the special `CREATEDB` privilege. See [sql-createrole](create-role.md#sql-createrole).


 By default, the new database will be created by cloning the standard system database `template1`. A different template can be specified by writing <code>TEMPLATE
   </code><em>name</em>. In particular, by writing `TEMPLATE template0`, you can create a pristine database (one where no user-defined objects exist and where the system objects have not been altered) containing only the standard objects predefined by your version of PostgreSQL. This is useful if you wish to avoid copying any installation-local objects that might have been added to `template1`.


## Parameters


<a id="create-database-name"></a>

*name*
:   The name of a database to create.
<a id="create-database-user-name"></a>

*user_name*
:   The role name of the user who will own the new database, or `DEFAULT` to use the default (namely, the user executing the command). To create a database owned by another role, you must be able to `SET ROLE` to that role.
<a id="create-database-template"></a>

*template*
:   The name of the template from which to create the new database, or `DEFAULT` to use the default template (`template1`).
<a id="create-database-encoding"></a>

*encoding*
:   Character set encoding to use in the new database. Specify a string constant (e.g., `'SQL_ASCII'`), or an integer encoding number, or `DEFAULT` to use the default encoding (namely, the encoding of the template database). The character sets supported by the PostgreSQL server are described in [Supported Character Sets](../../server-administration/localization/character-set-support.md#multibyte-charset-supported). See below for additional restrictions.
<a id="create-database-strategy"></a>

*strategy*
:   Strategy to be used in creating the new database. If the `WAL_LOG` strategy is used, the database will be copied block by block and each block will be separately written to the write-ahead log. This is the most efficient strategy in cases where the template database is small, and therefore it is the default. The older `FILE_COPY` strategy is also available. This strategy writes a small record to the write-ahead log for each tablespace used by the target database. Each such record represents copying an entire directory to a new location at the filesystem level. While this does reduce the write-ahead log volume substantially, especially if the template database is large, it also forces the system to perform a checkpoint both before and after the creation of the new database. In some situations, this may have a noticeable negative impact on overall system performance. The `FILE_COPY` strategy is affected by the [file_copy_method](../../server-administration/server-configuration/resource-consumption.md#guc-file-copy-method) setting.
<a id="create-database-locale"></a>

*locale*
:   Sets the default collation order and character classification in the new database. Collation affects the sort order applied to strings, e.g., in queries with `ORDER BY`, as well as the order used in indexes on text columns. Character classification affects the categorization of characters, e.g., lower, upper, and digit. Also sets the `LC_CTYPE` aspect of the operating system environment. The default is the same setting as the template database. See [libc Collations](../../server-administration/localization/collation-support.md#collation-managing-create-libc) and [ICU Collations](../../server-administration/localization/collation-support.md#collation-managing-create-icu) for details.


     Can be overridden by setting [lc_collate](#create-database-lc-collate), [lc_ctype](#create-database-lc-ctype), [builtin_locale](#create-database-builtin-locale), or [icu_locale](#create-database-icu-locale) individually.


     If [locale_provider](#create-database-locale-provider) is `builtin`, then *locale* or *builtin_locale* must be specified and set to either `C`, `C.UTF-8`, or `PG_UNICODE_FAST`.


    !!! tip

        The other locale settings [lc_messages](../../server-administration/server-configuration/client-connection-defaults.md#guc-lc-messages), [lc_monetary](../../server-administration/server-configuration/client-connection-defaults.md#guc-lc-monetary), [lc_numeric](../../server-administration/server-configuration/client-connection-defaults.md#guc-lc-numeric), and [lc_time](../../server-administration/server-configuration/client-connection-defaults.md#guc-lc-time) are not fixed per database and are not set by this command. If you want to make them the default for a specific database, you can use `ALTER DATABASE ... SET`.
<a id="create-database-lc-collate"></a>

*lc_collate*
:   If [locale_provider](#create-database-locale-provider) is `libc`, sets the default collation order to use in the new database, overriding the setting [locale](#create-database-locale). Otherwise, this setting is ignored.


     The default is the setting of [locale](#create-database-locale) if specified, otherwise the same setting as the template database. See below for additional restrictions.
<a id="create-database-lc-ctype"></a>

*lc_ctype*
:   Sets `LC_CTYPE` in the database server's operating system environment.


     If [locale_provider](#create-database-locale-provider) is `libc`, sets the default character classification to use in the new database, overriding the setting [locale](#create-database-locale).


     The default is the setting of [locale](#create-database-locale) if specified, otherwise the same setting as the template database. See below for additional restrictions.
<a id="create-database-builtin-locale"></a>

*builtin_locale*
:   Specifies the builtin provider locale for the database default collation order and character classification, overriding the setting [locale](#create-database-locale). The [locale provider](#create-database-locale-provider) must be `builtin`. The default is the setting of [locale](#create-database-locale) if specified; otherwise the same setting as the template database.


     The locales available for the `builtin` provider are `C`, `C.UTF-8` and `PG_UNICODE_FAST`.
<a id="create-database-icu-locale"></a>

*icu_locale*
:   Specifies the ICU locale (see [ICU Collations](../../server-administration/localization/collation-support.md#collation-managing-create-icu)) for the database default collation order and character classification, overriding the setting [locale](#create-database-locale). The [locale provider](#create-database-locale-provider) must be ICU. The default is the setting of [locale](#create-database-locale) if specified; otherwise the same setting as the template database.
<a id="create-database-icu-rules"></a>

*icu_rules*
:   Specifies additional collation rules to customize the behavior of the default collation of this database. This is supported for ICU only. See [ICU Tailoring Rules](../../server-administration/localization/collation-support.md#icu-tailoring-rules) for details.
<a id="create-database-locale-provider"></a>

*locale_provider*
:   Specifies the provider to use for the default collation in this database. Possible values are `builtin`, `icu` (if the server was built with ICU support) or `libc`. By default, the provider is the same as that of the [template](#create-database-template). See [Locale Providers](../../server-administration/localization/locale-support.md#locale-providers) for details.
<a id="create-database-collation-version"></a>

*collation_version*
:   Specifies the collation version string to store with the database. Normally, this should be omitted, which will cause the version to be computed from the actual version of the database collation as provided by the operating system. This option is intended to be used by `pg_upgrade` for copying the version from an existing installation.


     See also [sql-alterdatabase](alter-database.md#sql-alterdatabase) for how to handle database collation version mismatches.
<a id="create-database-tablespace-name"></a>

*tablespace_name*
:   The name of the tablespace that will be associated with the new database, or `DEFAULT` to use the template database's tablespace. This tablespace will be the default tablespace used for objects created in this database. See [sql-createtablespace](create-tablespace.md#sql-createtablespace) for more information.
<a id="create-database-allowconn"></a>

*allowconn*
:   If false then no one can connect to this database. The default is true, allowing connections (except as restricted by other mechanisms, such as `GRANT`/`REVOKE CONNECT`).
<a id="create-database-connlimit"></a>

*connlimit*
:   How many concurrent connections can be made to this database. -1 (the default) means no limit.
<a id="create-database-istemplate"></a>

*istemplate*
:   If true, then this database can be cloned by any user with `CREATEDB` privileges; if false (the default), then only superusers or the owner of the database can clone it.
<a id="create-database-oid"></a>

*oid*
:   The object identifier to be used for the new database. If this parameter is not specified, PostgreSQL will choose a suitable OID automatically. This parameter is primarily intended for internal use by pg_upgrade, and only pg_upgrade can specify a value less than 16384.


 Optional parameters can be written in any order, not only the order illustrated above.


## Notes


 `CREATE DATABASE` cannot be executed inside a transaction block.


 Errors along the line of “could not initialize database directory” are most likely related to insufficient permissions on the data directory, a full disk, or other file system problems.


 Use [`DROP DATABASE`](drop-database.md#sql-dropdatabase) to remove a database.


 The program [app-createdb](../postgresql-client-applications/createdb.md#app-createdb) is a wrapper program around this command, provided for convenience.


 Database-level configuration parameters (set via [`ALTER DATABASE`](alter-database.md#sql-alterdatabase)) and database-level permissions (set via [`GRANT`](grant.md#sql-grant)) are not copied from the template database.


 Although it is possible to copy a database other than `template1` by specifying its name as the template, this is not (yet) intended as a general-purpose “`COPY DATABASE`” facility. The principal limitation is that no other sessions can be connected to the template database while it is being copied. `CREATE DATABASE` will fail if any other connection exists when it starts; otherwise, new connections to the template database are locked out until `CREATE DATABASE` completes. See [Template Databases](../../server-administration/managing-databases/template-databases.md#manage-ag-templatedbs) for more information.


 The character set encoding specified for the new database must be compatible with the chosen locale settings (`LC_COLLATE` and `LC_CTYPE`). If the locale is `C` (or equivalently `POSIX`), then all encodings are allowed, but for other locale settings there is only one encoding that will work properly. (On Windows, however, UTF-8 encoding can be used with any locale.) `CREATE DATABASE` will allow superusers to specify `SQL_ASCII` encoding regardless of the locale settings, but this choice is deprecated and may result in misbehavior of character-string functions if data that is not encoding-compatible with the locale is stored in the database.


 The encoding and locale settings must match those of the template database, except when `template0` is used as template. This is because other databases might contain data that does not match the specified encoding, or might contain indexes whose sort ordering is affected by `LC_COLLATE` and `LC_CTYPE`. Copying such data would result in a database that is corrupt according to the new settings. `template0`, however, is known to not contain any data or indexes that would be affected.


 There is currently no option to use a database locale with nondeterministic comparisons (see [`CREATE COLLATION`](create-collation.md#sql-createcollation) for an explanation). If this is needed, then per-column collations would need to be used.


 The `CONNECTION LIMIT` option is only enforced approximately; if two new sessions start at about the same time when just one connection “slot” remains for the database, it is possible that both will fail. Also, the limit is not enforced against superusers or background worker processes.


## Examples


 To create a new database:

```sql

CREATE DATABASE lusiadas;
```


 To create a database `sales` owned by user `salesapp` with a default tablespace of `salesspace`:

```sql

CREATE DATABASE sales OWNER salesapp TABLESPACE salesspace;
```


 To create a database `music` with a different locale:

```sql

CREATE DATABASE music
    LOCALE 'sv_SE.utf8'
    TEMPLATE template0;
```
 In this example, the `TEMPLATE template0` clause is required if the specified locale is different from the one in `template1`. (If it is not, then specifying the locale explicitly is redundant.)


 To create a database `music2` with a different locale and a different character set encoding:

```sql

CREATE DATABASE music2
    LOCALE 'sv_SE.iso885915'
    ENCODING LATIN9
    TEMPLATE template0;
```
 The specified locale and encoding settings must match, or an error will be reported.


 Note that locale names are specific to the operating system, so that the above commands might not work in the same way everywhere.


## Compatibility


 There is no `CREATE DATABASE` statement in the SQL standard. Databases are equivalent to catalogs, whose creation is implementation-defined.


## See Also
  [sql-alterdatabase](alter-database.md#sql-alterdatabase), [sql-dropdatabase](drop-database.md#sql-dropdatabase)
