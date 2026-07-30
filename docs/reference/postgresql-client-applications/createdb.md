<a id="app-createdb"></a>

# createdb

create a new PostgreSQL database

## Synopsis


```
createdb [CONNECTION-OPTION...] [OPTION...] [DBNAME
    [DESCRIPTION]]
```
 <a id="r1-app-createdb-1"></a>

## Description


 createdb creates a new PostgreSQL database.


 Normally, the database user who executes this command becomes the owner of the new database. However, a different owner can be specified via the `-O` option, if the executing user has appropriate privileges.


 createdb is a wrapper around the SQL command [`CREATE DATABASE`](../sql-commands/create-database.md#sql-createdatabase). There is no effective difference between creating databases via this utility and via other methods for accessing the server.


## Options


 createdb accepts the following command-line arguments:

*dbname*
:   Specifies the name of the database to be created. The name must be unique among all PostgreSQL databases in this cluster. The default is to create a database with the same name as the current system user.

*description*
:   Specifies a comment to be associated with the newly created database.

<code>-D </code><em>tablespace</em>, <code>--tablespace=</code><em>tablespace</em>
:   Specifies the default tablespace for the database. (This name is processed as a double-quoted identifier.)

`-e`, `--echo`
:   Echo the commands that createdb generates and sends to the server.

<code>-E </code><em>encoding</em>, <code>--encoding=</code><em>encoding</em>
:   Specifies the character encoding scheme to be used in this database. The character sets supported by the PostgreSQL server are described in [Supported Character Sets](../../server-administration/localization/character-set-support.md#multibyte-charset-supported).

<code>-l </code><em>locale</em>, <code>--locale=</code><em>locale</em>
:   Specifies the locale to be used in this database. This is equivalent to specifying `--lc-collate`, `--lc-ctype`, and `--icu-locale` to the same value. Some locales are only valid for ICU and must be set with `--icu-locale`.

<code>--lc-collate=</code><em>locale</em>
:   Specifies the LC_COLLATE setting to be used in this database (ignored unless the locale provider is `libc`).

<code>--lc-ctype=</code><em>locale</em>
:   Specifies the LC_CTYPE setting to be used in this database.

<code>--builtin-locale=</code><em>locale</em>
:   Specifies the locale name when the builtin provider is used. Locale support is described in [Locale Support](../../server-administration/localization/locale-support.md#locale).

<code>--icu-locale=</code><em>locale</em>
:   Specifies the ICU locale ID to be used in this database, if the ICU locale provider is selected.

<code>--icu-rules=</code><em>rules</em>
:   Specifies additional collation rules to customize the behavior of the default collation of this database. This is supported for ICU only.

`--locale-provider={`builtin`|`libc`|`icu`}`
:   Specifies the locale provider for the database's default collation.

<code>-O </code><em>owner</em>, <code>--owner=</code><em>owner</em>
:   Specifies the database user who will own the new database. (This name is processed as a double-quoted identifier.)

<code>-S </code><em>strategy</em>, <code>--strategy=</code><em>strategy</em>
:   Specifies the database creation strategy. See [CREATE DATABASE STRATEGY](../sql-commands/create-database.md#create-database-strategy) for more details.

<code>-T </code><em>template</em>, <code>--template=</code><em>template</em>
:   Specifies the template database from which to build this database. (This name is processed as a double-quoted identifier.)

`-V`, `--version`
:   Print the createdb version and exit.

`-?`, `--help`
:   Show help about createdb command line arguments, and exit.


 The options `-D`, `-l`, `-E`, `-O`, and `-T` correspond to options of the underlying SQL command [`CREATE DATABASE`](../sql-commands/create-database.md#sql-createdatabase); see there for more information about them.


 createdb also accepts the following command-line arguments for connection parameters:

<code>-h </code><em>host</em>, <code>--host=</code><em>host</em>
:   Specifies the host name of the machine on which the server is running. If the value begins with a slash, it is used as the directory for the Unix domain socket.

<code>-p </code><em>port</em>, <code>--port=</code><em>port</em>
:   Specifies the TCP port or the local Unix domain socket file extension on which the server is listening for connections.

<code>-U </code><em>username</em>, <code>--username=</code><em>username</em>
:   User name to connect as.

`-w`, `--no-password`
:   Never issue a password prompt. If the server requires password authentication and a password is not available by other means such as a `.pgpass` file, the connection attempt will fail. This option can be useful in batch jobs and scripts where no user is present to enter a password.

`-W`, `--password`
:   Force createdb to prompt for a password before connecting to a database.


     This option is never essential, since createdb will automatically prompt for a password if the server demands password authentication. However, createdb will waste a connection attempt finding out that the server wants a password. In some cases it is worth typing `-W` to avoid the extra connection attempt.

<code>--maintenance-db=</code><em>dbname</em>
:   Specifies the name of the database to connect to when creating the new database. If not specified, the `postgres` database will be used; if that does not exist (or if it is the name of the new database being created), `template1` will be used. This can be a [connection string](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connstring). If so, connection string parameters will override any conflicting command line options.


## Environment


`PGDATABASE`
:   If set, the name of the database to create, unless overridden on the command line.

`PGHOST`, `PGPORT`, `PGUSER`
:   Default connection parameters. `PGUSER` also determines the name of the database to create, if it is not specified on the command line or by `PGDATABASE`.

`PG_COLOR`
:   Specifies whether to use color in diagnostic messages. Possible values are `always`, `auto` and `never`.


 This utility, like most other PostgreSQL utilities, also uses the environment variables supported by libpq (see [Environment Variables](../../client-interfaces/libpq-c-library/environment-variables.md#libpq-envars)).


## Diagnostics


 In case of difficulty, see [sql-createdatabase](../sql-commands/create-database.md#sql-createdatabase) and [app-psql](psql.md#app-psql) for discussions of potential problems and error messages. The database server must be running at the targeted host. Also, any default connection settings and environment variables used by the libpq front-end library will apply.


## Examples


 To create the database `demo` using the default database server:

```

$ createdb demo
```


 To create the database `demo` using the server on host `eden`, port 5000, using the `template0` template database, here is the command-line command and the underlying SQL command:

```

$ createdb -p 5000 -h eden -T template0 -e demo
CREATE DATABASE demo TEMPLATE template0;
```


## See Also
  [app-dropdb](dropdb.md#app-dropdb), [sql-createdatabase](../sql-commands/create-database.md#sql-createdatabase)
