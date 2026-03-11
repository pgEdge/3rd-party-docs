<a id="app-initdb"></a>

# initdb

create a new PostgreSQL database cluster

## Synopsis


```
initdb [OPTION...] {
     --pgdata
     -D
     |  DIRECTORY}
```
 <a id="r1-app-initdb-1"></a>

## Description


 `initdb` creates a new PostgreSQL *database cluster*.


 Creating a database cluster consists of creating the *directories* in which the cluster data will live, generating the shared catalog tables (tables that belong to the whole cluster rather than to any particular database), and creating the `postgres`, `template1`, and `template0` databases. The `postgres` database is a default database meant for use by users, utilities and third party applications. `template1` and `template0` are meant as source databases to be copied by later `CREATE DATABASE` commands. `template0` should never be modified, but you can add objects to `template1`, which by default will be copied into databases created later. See [Template Databases](../../server-administration/managing-databases/template-databases.md#manage-ag-templatedbs) for more details.


 Although `initdb` will attempt to create the specified data directory, it might not have permission if the parent directory of the desired data directory is root-owned. To initialize in such a setup, create an empty data directory as root, then use `chown` to assign ownership of that directory to the database user account, then `su` to become the database user to run `initdb`.


 `initdb` must be run as the user that will own the server process, because the server needs to have access to the files and directories that `initdb` creates. Since the server cannot be run as root, you must not run `initdb` as root either. (It will in fact refuse to do so.)


 For security reasons the new cluster created by `initdb` will only be accessible by the cluster owner by default. The `--allow-group-access` option allows any user in the same group as the cluster owner to read files in the cluster. This is useful for performing backups as a non-privileged user.


 `initdb` initializes the database cluster's default locale and character set encoding. These can also be set separately for each database when it is created. `initdb` determines those settings for the template databases, which will serve as the default for all other databases.


 By default, `initdb` uses the locale provider `libc` (see [Locale Providers](../../server-administration/localization/locale-support.md#locale-providers)). The `libc` locale provider takes the locale settings from the environment, and determines the encoding from the locale settings.


 To choose a different locale for the cluster, use the option `--locale`. There are also individual options `--lc-*` and `--icu-locale` (see below) to set values for the individual locale categories. Note that inconsistent settings for different locale categories can give nonsensical results, so this should be used with care.


 Alternatively, `initdb` can use the ICU library to provide locale services by specifying `--locale-provider=icu`. The server must be built with ICU support. To choose the specific ICU locale ID to apply, use the option `--icu-locale`. Note that for implementation reasons and to support legacy code, `initdb` will still select and initialize libc locale settings when the ICU locale provider is used.


 When `initdb` runs, it will print out the locale settings it has chosen. If you have complex requirements or specified multiple options, it is advisable to check that the result matches what was intended.


 More details about locale settings can be found in [Locale Support](../../server-administration/localization/locale-support.md#locale).


 To alter the default encoding, use the `--encoding`. More details can be found in [Character Set Support](../../server-administration/localization/character-set-support.md#multibyte).


## Options


<a id="app-initdb-option-auth"></a>

<code>-A </code><em>authmethod</em>, <code>--auth=</code><em>authmethod</em>
:   This option specifies the default authentication method for local users used in `pg_hba.conf` (`host` and `local` lines). See [The `pg_hba.conf` File](../../server-administration/client-authentication/the-pg_hba-conf-file.md#auth-pg-hba-conf) for an overview of valid values.


     `initdb` will prepopulate `pg_hba.conf` entries using the specified authentication method for non-replication as well as replication connections.


     Do not use `trust` unless you trust all local users on your system. `trust` is the default for ease of installation.
<a id="app-initdb-option-auth-host"></a>

<code>--auth-host=</code><em>authmethod</em>
:   This option specifies the authentication method for local users via TCP/IP connections used in `pg_hba.conf` (`host` lines).
<a id="app-initdb-option-auth-local"></a>

<code>--auth-local=</code><em>authmethod</em>
:   This option specifies the authentication method for local users via Unix-domain socket connections used in `pg_hba.conf` (`local` lines).
<a id="app-initdb-option-pgdata"></a>

<code>-D </code><em>directory</em>, <code>--pgdata=</code><em>directory</em>
:   This option specifies the directory where the database cluster should be stored. This is the only information required by `initdb`, but you can avoid writing it by setting the `PGDATA` environment variable, which can be convenient since the database server (`postgres`) can find the data directory later by the same variable.
<a id="app-initdb-option-encoding"></a>

<code>-E </code><em>encoding</em>, <code>--encoding=</code><em>encoding</em>
:   Selects the encoding of the template databases. This will also be the default encoding of any database you create later, unless you override it then. The character sets supported by the PostgreSQL server are described in [Supported Character Sets](../../server-administration/localization/character-set-support.md#multibyte-charset-supported).


     By default, the template database encoding is derived from the locale. If [--no-locale](#app-initdb-option-no-locale) is specified (or equivalently, if the locale is `C` or `POSIX`), then the default is `UTF8` for the ICU provider and `SQL_ASCII` for the `libc` provider.
<a id="app-initdb-allow-group-access"></a>

`-g`, `--allow-group-access`
:   Allows users in the same group as the cluster owner to read all cluster files created by `initdb`. This option is ignored on Windows as it does not support POSIX-style group permissions.
<a id="app-initdb-icu-locale"></a>

<code>--icu-locale=</code><em>locale</em>
:   Specifies the ICU locale when the ICU provider is used. Locale support is described in [Locale Support](../../server-administration/localization/locale-support.md#locale).
<a id="app-initdb-icu-rules"></a>

<code>--icu-rules=</code><em>rules</em>
:   Specifies additional collation rules to customize the behavior of the default collation. This is supported for ICU only.
<a id="app-initdb-data-checksums"></a>

`-k`, `--data-checksums`
:   Use checksums on data pages to help detect corruption by the I/O system that would otherwise be silent. This is enabled by default; use [--no-data-checksums](#app-initdb-no-data-checksums) to disable checksums.


     Enabling checksums might incur a small performance penalty. If set, checksums are calculated for all objects, in all databases. All checksum failures will be reported in the [`pg_stat_database`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-database-view) view. See [Data Checksums](../../server-administration/reliability-and-the-write-ahead-log/data-checksums.md#checksums) for details.
<a id="app-initdb-option-locale"></a>

<code>--locale=</code><em>locale</em>
:   Sets the default locale for the database cluster. If this option is not specified, the locale is inherited from the environment that `initdb` runs in. Locale support is described in [Locale Support](../../server-administration/localization/locale-support.md#locale).


     If `--locale-provider` is `builtin`, `--locale` or `--builtin-locale` must be specified and set to `C`, `C.UTF-8` or `PG_UNICODE_FAST`.
<a id="app-initdb-option-lc-collate"></a>

<code>--lc-collate=</code><em>locale</em>, <code>--lc-ctype=</code><em>locale</em>, <code>--lc-messages=</code><em>locale</em>, <code>--lc-monetary=</code><em>locale</em>, <code>--lc-numeric=</code><em>locale</em>, <code>--lc-time=</code><em>locale</em>
:   Like `--locale`, but only sets the locale in the specified category.
<a id="app-initdb-option-no-locale"></a>

`--no-locale`
:   Equivalent to `--locale=C`.
<a id="app-initdb-builtin-locale"></a>

<code>--builtin-locale=</code><em>locale</em>
:   Specifies the locale name when the builtin provider is used. Locale support is described in [Locale Support](../../server-administration/localization/locale-support.md#locale).
<a id="app-initdb-option-locale-provider"></a>

`--locale-provider={`builtin`|`libc`|`icu`}`
:   This option sets the locale provider for databases created in the new cluster. It can be overridden in the `CREATE DATABASE` command when new databases are subsequently created. The default is `libc` (see [Locale Providers](../../server-administration/localization/locale-support.md#locale-providers)).
<a id="app-initdb-no-data-checksums"></a>

`--no-data-checksums`
:   Do not enable data checksums.
<a id="app-initdb-option-pwfile"></a>

<code>--pwfile=</code><em>filename</em>
:   Makes `initdb` read the bootstrap superuser's password from a file. The first line of the file is taken as the password.
<a id="app-initdb-option-text-search-config"></a>

<code>-T </code><em>config</em>, <code>--text-search-config=</code><em>config</em>
:   Sets the default text search configuration. See [default_text_search_config](../../server-administration/server-configuration/client-connection-defaults.md#guc-default-text-search-config) for further information.
<a id="app-initdb-option-username"></a>

<code>-U </code><em>username</em>, <code>--username=</code><em>username</em>
:   Sets the user name of the *bootstrap superuser*. This defaults to the name of the operating-system user running `initdb`.
<a id="app-initdb-option-pwprompt"></a>

`-W`, `--pwprompt`
:   Makes `initdb` prompt for a password to give the bootstrap superuser. If you don't plan on using password authentication, this is not important. Otherwise you won't be able to use password authentication until you have a password set up.
<a id="app-initdb-option-waldir"></a>

<code>-X </code><em>directory</em>, <code>--waldir=</code><em>directory</em>
:   This option specifies the directory where the write-ahead log should be stored.
<a id="app-initdb-option-wal-segsize"></a>

<code>--wal-segsize=</code><em>size</em>
:   Set the *WAL segment size*, in megabytes. This is the size of each individual file in the WAL log. The default size is 16 megabytes. The value must be a power of 2 between 1 and 1024 (megabytes). This option can only be set during initialization, and cannot be changed later.


     It may be useful to adjust this size to control the granularity of WAL log shipping or archiving. Also, in databases with a high volume of WAL, the sheer number of WAL files per directory can become a performance and management problem. Increasing the WAL file size will reduce the number of WAL files.


 Other, less commonly used, options are also available:

<a id="app-initdb-option-set"></a>

<code>-c </code><em>name</em><code>=</code><em>value</em>, <code>--set </code><em>name</em><code>=</code><em>value</em>
:   Forcibly set the server parameter *name* to *value* during `initdb`, and also install that setting in the generated `postgresql.conf` file, so that it will apply during future server runs. This option can be given more than once to set several parameters. It is primarily useful when the environment is such that the server will not start at all using the default parameters.
<a id="app-initdb-option-debug"></a>

`-d`, `--debug`
:   Print debugging output from the bootstrap backend and a few other messages of lesser interest for the general public. The bootstrap backend is the program `initdb` uses to create the catalog tables. This option generates a tremendous amount of extremely boring output.
<a id="app-initdb-option-discard-caches"></a>

`--discard-caches`
:   Run the bootstrap backend with the `debug_discard_caches=1` option. This takes a very long time and is only of use for deep debugging.
<a id="app-initdb-option-l"></a>

<code>-L </code><em>directory</em>
:   Specifies where `initdb` should find its input files to initialize the database cluster. This is normally not necessary. You will be told if you need to specify their location explicitly.
<a id="app-initdb-option-no-clean"></a>

`-n`, `--no-clean`
:   By default, when `initdb` determines that an error prevented it from completely creating the database cluster, it removes any files it might have created before discovering that it cannot finish the job. This option inhibits tidying-up and is thus useful for debugging.
<a id="app-initdb-option-no-sync"></a>

`-N`, `--no-sync`
:   By default, `initdb` will wait for all files to be written safely to disk. This option causes `initdb` to return without waiting, which is faster, but means that a subsequent operating system crash can leave the data directory corrupt. Generally, this option is useful for testing, but should not be used when creating a production installation.
<a id="app-initdb-option-no-sync-data-files"></a>

`--no-sync-data-files`
:   By default, `initdb` safely writes all database files to disk. This option instructs `initdb` to skip synchronizing all files in the individual database directories, the database directories themselves, and the tablespace directories, i.e., everything in the `base` subdirectory and any other tablespace directories. Other files, such as those in `pg_wal` and `pg_xact`, will still be synchronized unless the `--no-sync` option is also specified.


     Note that if `--no-sync-data-files` is used in conjunction with `--sync-method=syncfs`, some or all of the aforementioned files and directories will be synchronized because `syncfs` processes entire file systems.


     This option is primarily intended for internal use by tools that separately ensure the skipped files are synchronized to disk.
<a id="app-initdb-option-no-instructions"></a>

`--no-instructions`
:   By default, `initdb` will write instructions for how to start the cluster at the end of its output. This option causes those instructions to be left out. This is primarily intended for use by tools that wrap `initdb` in platform-specific behavior, where those instructions are likely to be incorrect.
<a id="app-initdb-option-show"></a>

`-s`, `--show`
:   Show internal settings and exit, without doing anything else. This can be used to debug the initdb installation.
<a id="app-initdb-option-sync-method"></a>

<code>--sync-method=</code><em>method</em>
:   When set to `fsync`, which is the default, `initdb` will recursively open and synchronize all files in the data directory. The search for files will follow symbolic links for the WAL directory and each configured tablespace.


     On Linux, `syncfs` may be used instead to ask the operating system to synchronize the whole file systems that contain the data directory, the WAL files, and each tablespace. See [recovery_init_sync_method](../../server-administration/server-configuration/error-handling.md#guc-recovery-init-sync-method) for information about the caveats to be aware of when using `syncfs`.


     This option has no effect when `--no-sync` is used.
<a id="app-initdb-option-sync-only"></a>

`-S`, `--sync-only`
:   Safely write all database files to disk and exit. This does not perform any of the normal initdb operations. Generally, this option is useful for ensuring reliable recovery after changing [fsync](../../server-administration/server-configuration/write-ahead-log.md#guc-fsync) from `off` to `on`.


 Other options:

<a id="app-initdb-option-version"></a>

`-V`, `--version`
:   Print the initdb version and exit.
<a id="app-initdb-option-help"></a>

`-?`, `--help`
:   Show help about initdb command line arguments, and exit.


## Environment


<a id="app-initdb-environment-pgdata"></a>

`PGDATA`
:   Specifies the directory where the database cluster is to be stored; can be overridden using the `-D` option.
<a id="app-initdb-environment-pg-color"></a>

`PG_COLOR`
:   Specifies whether to use color in diagnostic messages. Possible values are `always`, `auto` and `never`.
<a id="app-initdb-environment-tz"></a>

`TZ`
:   Specifies the default time zone of the created database cluster. The value should be a full time zone name (see [Time Zones](../../the-sql-language/data-types/date-time-types.md#datatype-timezones)).


## Notes


 `initdb` can also be invoked via `pg_ctl initdb`.


## See Also
  [app-pg-ctl](pg_ctl.md#app-pg-ctl), [app-postgres](postgres.md#app-postgres), [The `pg_hba.conf` File](../../server-administration/client-authentication/the-pg_hba-conf-file.md#auth-pg-hba-conf)
