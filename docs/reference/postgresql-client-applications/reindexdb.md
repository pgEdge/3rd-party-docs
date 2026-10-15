<a id="app-reindexdb"></a>

# reindexdb

reindex a PostgreSQL database

## Synopsis


```
reindexdb [CONNECTION-OPTION...] [OPTION...]
     [
      {-S | --schema}
     SCHEMA
    ]
   ...
     [
      {-t | --table}
     TABLE
    ]
   ...
     [
      {-i | --index}
     INDEX
    ]
   ...
     [
      {-s | --system}
    ]
    [
     {DBNAME | -a | --all}
   ]
```


## Description


 reindexdb is a utility for rebuilding indexes in a PostgreSQL database.


 reindexdb is a wrapper around the SQL command [`REINDEX`](../sql-commands/reindex.md#sql-reindex). There is no effective difference between reindexing databases via this utility and via other methods for accessing the server.


## Options


 reindexdb accepts the following command-line arguments:

`-a`, `--all`
:   Reindex all databases.

`--concurrently`
:   Use the `CONCURRENTLY` option. See [sql-reindex](../sql-commands/reindex.md#sql-reindex), where all the caveats of this option are explained in detail.

<code>[-d] </code><em>dbname</em>, <code>[--dbname=]</code><em>dbname</em>
:   Specifies the name of the database to be reindexed, when `-a`/`--all` is not used. If this is not specified, the database name is read from the environment variable `PGDATABASE`. If that is not set, the user name specified for the connection is used. The *dbname* can be a [connection string](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connstring). If so, connection string parameters will override any conflicting command line options.

`-e`, `--echo`
:   Echo the commands that reindexdb generates and sends to the server.

<code>-i </code><em>index</em>, <code>--index=</code><em>index</em>
:   Recreate *index* only. Multiple indexes can be recreated by writing multiple `-i` switches.

<code>-j </code><em>njobs</em>, <code>--jobs=</code><em>njobs</em>
:   Execute the reindex commands in parallel by running *njobs* commands simultaneously. This option may reduce the processing time but it also increases the load on the database server.


     reindexdb will open *njobs* connections to the database, so make sure your [max_connections](../../server-administration/server-configuration/connections-and-authentication.md#guc-max-connections) setting is high enough to accommodate all connections.


     Note that this option is incompatible with the `--system` option.

`-q`, `--quiet`
:   Do not display progress messages.

`-s`, `--system`
:   Reindex database's system catalogs only.

<code>-S </code><em>schema</em>, <code>--schema=</code><em>schema</em>
:   Reindex *schema* only. Multiple schemas can be reindexed by writing multiple `-S` switches.

<code>-t </code><em>table</em>, <code>--table=</code><em>table</em>
:   Reindex *table* only. Multiple tables can be reindexed by writing multiple `-t` switches.

<code>--tablespace=</code><em>tablespace</em>
:   Specifies the tablespace where indexes are rebuilt. (This name is processed as a double-quoted identifier.)

`-v`, `--verbose`
:   Print detailed information during processing.

`-V`, `--version`
:   Print the reindexdb version and exit.

`-?`, `--help`
:   Show help about reindexdb command line arguments, and exit.


 reindexdb also accepts the following command-line arguments for connection parameters:

<code>-h </code><em>host</em>, <code>--host=</code><em>host</em>
:   Specifies the host name of the machine on which the server is running. If the value begins with a slash, it is used as the directory for the Unix domain socket.

<code>-p </code><em>port</em>, <code>--port=</code><em>port</em>
:   Specifies the TCP port or local Unix domain socket file extension on which the server is listening for connections.

<code>-U </code><em>username</em>, <code>--username=</code><em>username</em>
:   User name to connect as.

`-w`, `--no-password`
:   Never issue a password prompt. If the server requires password authentication and a password is not available by other means such as a `.pgpass` file, the connection attempt will fail. This option can be useful in batch jobs and scripts where no user is present to enter a password.

`-W`, `--password`
:   Force reindexdb to prompt for a password before connecting to a database.


     This option is never essential, since reindexdb will automatically prompt for a password if the server demands password authentication. However, reindexdb will waste a connection attempt finding out that the server wants a password. In some cases it is worth typing `-W` to avoid the extra connection attempt.

<code>--maintenance-db=</code><em>dbname</em>
:   When the `-a`/`--all` is used, connect to this database to gather the list of databases to reindex. If not specified, the `postgres` database will be used, or if that does not exist, `template1` will be used. This can be a [connection string](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connstring). If so, connection string parameters will override any conflicting command line options. Also, connection string parameters other than the database name itself will be re-used when connecting to other databases.


## Environment


`PGDATABASE`, `PGHOST`, `PGPORT`, `PGUSER`
:   Default connection parameters

`PG_COLOR`
:   Specifies whether to use color in diagnostic messages. Possible values are `always`, `auto` and `never`.


 This utility, like most other PostgreSQL utilities, also uses the environment variables supported by libpq (see [Environment Variables](../../client-interfaces/libpq-c-library/environment-variables.md#libpq-envars)).


## Diagnostics


 In case of difficulty, see [sql-reindex](../sql-commands/reindex.md#sql-reindex) and [app-psql](psql.md#app-psql) for discussions of potential problems and error messages. The database server must be running at the targeted host. Also, any default connection settings and environment variables used by the libpq front-end library will apply.


## Examples


 To reindex the database `test`:

```

$ reindexdb test
```


 To reindex the table `foo` and the index `bar` in a database named `abcd`:

```

$ reindexdb --table=foo --index=bar abcd
```


## See Also
  [sql-reindex](../sql-commands/reindex.md#sql-reindex)
