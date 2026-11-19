<a id="app-vacuumdb"></a>

# vacuumdb

garbage-collect and analyze a PostgreSQL database

## Synopsis


```
vacuumdb [CONNECTION-OPTION...] [OPTION...]
     [
      {-t | --table}
     TABLE
      [( COLUMN [,...] )]
    ]
   ... [
     {DBNAME | -a | --all}
   ]
```


```
vacuumdb [CONNECTION-OPTION...] [OPTION...]
     [
      {-n | --schema}
     SCHEMA
    ]
   ... [
     {DBNAME | -a | --all}
   ]
```


```
vacuumdb [CONNECTION-OPTION...] [OPTION...]
     [
      {-N | --exclude-schema}
     SCHEMA
    ]
   ... [
     {DBNAME | -a | --all}
   ]
```


## Description


 vacuumdb is a utility for cleaning a PostgreSQL database. vacuumdb will also generate internal statistics used by the PostgreSQL query optimizer.


 vacuumdb is a wrapper around the SQL command [`VACUUM`](../sql-commands/vacuum.md#sql-vacuum). There is no effective difference between vacuuming and analyzing databases via this utility and via other methods for accessing the server.


## Options


 vacuumdb accepts the following command-line arguments:

`-a`, `--all`
:   Vacuum all databases.

<code>--buffer-usage-limit </code><em>size</em>
:   Specifies the *Buffer Access Strategy* ring buffer size for a given invocation of vacuumdb. This size is used to calculate the number of shared buffers which will be reused as part of this strategy. See [sql-vacuum](../sql-commands/vacuum.md#sql-vacuum).

<code>[-d] </code><em>dbname</em>, <code>[--dbname=]</code><em>dbname</em>
:   Specifies the name of the database to be cleaned or analyzed, when `-a`/`--all` is not used. If this is not specified, the database name is read from the environment variable `PGDATABASE`. If that is not set, the user name specified for the connection is used. The *dbname* can be a [connection string](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connstring). If so, connection string parameters will override any conflicting command line options.

`--disable-page-skipping`
:   Disable skipping pages based on the contents of the visibility map.

`--dry-run`
:   Print, but do not execute, the vacuum and analyze commands that would have been sent to the server.

`-e`, `--echo`
:   Echo the commands that vacuumdb generates and sends to the server.

`-f`, `--full`
:   Perform “full” vacuuming.

`-F`, `--freeze`
:   Aggressively “freeze” tuples.

`--force-index-cleanup`
:   Always remove index entries pointing to dead tuples.

<code>-j </code><em>njobs</em>, <code>--jobs=</code><em>njobs</em>
:   Execute the vacuum or analyze commands in parallel by running *njobs* commands simultaneously. This option may reduce the processing time but it also increases the load on the database server.


     vacuumdb will open *njobs* connections to the database, so make sure your [max_connections](../../server-administration/server-configuration/connections-and-authentication.md#guc-max-connections) setting is high enough to accommodate all connections.


     Note that using this mode together with the `-f` (`FULL`) option might cause deadlock failures if certain system catalogs are processed in parallel.

<code>--min-mxid-age </code><em>mxid_age</em>
:   Only execute the vacuum or analyze commands on tables with a multixact ID age of at least *mxid_age*. This setting is useful for prioritizing tables to process to prevent multixact ID wraparound (see [Multixacts and Wraparound](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-multixact-wraparound)).


     For the purposes of this option, the multixact ID age of a relation is the greatest of the ages of the main relation and its associated TOAST table, if one exists. Since the commands issued by vacuumdb will also process the TOAST table for the relation if necessary, it does not need to be considered separately.

<code>--min-xid-age </code><em>xid_age</em>
:   Only execute the vacuum or analyze commands on tables with a transaction ID age of at least *xid_age*. This setting is useful for prioritizing tables to process to prevent transaction ID wraparound (see [Preventing Transaction ID Wraparound Failures](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound)).


     For the purposes of this option, the transaction ID age of a relation is the greatest of the ages of the main relation and its associated TOAST table, if one exists. Since the commands issued by vacuumdb will also process the TOAST table for the relation if necessary, it does not need to be considered separately.

`--missing-stats-only`
:   Only analyze relations that are missing statistics for a column, index expression, or extended statistics object. When used with `--analyze-in-stages`, this option prevents vacuumdb from temporarily replacing existing statistics with ones generated with lower statistics targets, thus avoiding transiently worse query optimizer choices.


     This option can only be used in conjunction with `--analyze-only` or `--analyze-in-stages`.


     Note that `--missing-stats-only` requires `SELECT` privileges on [`pg_statistic`](../../internals/system-catalogs/pg_statistic.md#catalog-pg-statistic) and [`pg_statistic_ext_data`](../../internals/system-catalogs/pg_statistic_ext_data.md#catalog-pg-statistic-ext-data), which are restricted to superusers by default.

<code>-n </code><em>schema</em>, <code>--schema=</code><em>schema</em>
:   Clean or analyze all tables in *schema* only. Multiple schemas can be vacuumed by writing multiple `-n` switches.

<code>-N </code><em>schema</em>, <code>--exclude-schema=</code><em>schema</em>
:   Do not clean or analyze any tables in *schema*. Multiple schemas can be excluded by writing multiple `-N` switches.

`--no-index-cleanup`
:   Do not remove index entries pointing to dead tuples.

`--no-process-main`
:   Skip the main relation.

`--no-process-toast`
:   Skip the TOAST table associated with the table to vacuum, if any.

`--no-truncate`
:   Do not truncate empty pages at the end of the table.

<code>-P </code><em>parallel_workers</em>, <code>--parallel=</code><em>parallel_workers</em>
:   Specify the number of parallel workers for *parallel vacuum*. This allows the vacuum to leverage multiple CPUs to process indexes. See [sql-vacuum](../sql-commands/vacuum.md#sql-vacuum).

`-q`, `--quiet`
:   Do not display progress messages.

`--skip-locked`
:   Skip relations that cannot be immediately locked for processing.

<code>-t </code><em>table</em><code> [ (</code><em>column</em><code> [,...]) ]</code>, <code>--table=</code><em>table</em><code> [ (</code><em>column</em><code> [,...]) ]</code>
:   Clean or analyze *table* only. Column names can be specified only in conjunction with the `--analyze` or `--analyze-only` options. Multiple tables can be vacuumed by writing multiple `-t` switches.


     If no tables are specified with the `--table` option, vacuumdb will clean all regular tables and materialized views in the connected database. If `--analyze-only` or `--analyze-in-stages` is also specified, it will analyze all regular tables, partitioned tables, and materialized views (but not foreign tables).


    !!! tip

        If you specify columns, you probably have to escape the parentheses from the shell. (See examples below.)

`-v`, `--verbose`
:   Print detailed information during processing.

`-V`, `--version`
:   Print the vacuumdb version and exit.

`-z`, `--analyze`
:   Also calculate statistics for use by the optimizer.

`-Z`, `--analyze-only`
:   Only calculate statistics for use by the optimizer (no vacuum).

`--analyze-in-stages`
:   Only calculate statistics for use by the optimizer (no vacuum), like `--analyze-only`. Run three stages of analyze; the first stage uses the lowest possible statistics target (see [default_statistics_target](../../server-administration/server-configuration/query-planning.md#guc-default-statistics-target)) to produce usable statistics faster, and subsequent stages build the full statistics.


     This option is only useful to analyze a database that currently has no statistics or has wholly incorrect ones, such as if it is newly populated from a restored dump or by `pg_upgrade`. Be aware that running with this option in a database with existing statistics may cause the query optimizer choices to become transiently worse due to the low statistics targets of the early stages.

`-?`, `--help`
:   Show help about vacuumdb command line arguments, and exit.


 vacuumdb also accepts the following command-line arguments for connection parameters:

<code>-h </code><em>host</em>, <code>--host=</code><em>host</em>
:   Specifies the host name of the machine on which the server is running. If the value begins with a slash, it is used as the directory for the Unix domain socket.

<code>-p </code><em>port</em>, <code>--port=</code><em>port</em>
:   Specifies the TCP port or local Unix domain socket file extension on which the server is listening for connections.

<code>-U </code><em>username</em>, <code>--username=</code><em>username</em>
:   User name to connect as.

`-w`, `--no-password`
:   Never issue a password prompt. If the server requires password authentication and a password is not available by other means such as a `.pgpass` file, the connection attempt will fail. This option can be useful in batch jobs and scripts where no user is present to enter a password.

`-W`, `--password`
:   Force vacuumdb to prompt for a password before connecting to a database.


     This option is never essential, since vacuumdb will automatically prompt for a password if the server demands password authentication. However, vacuumdb will waste a connection attempt finding out that the server wants a password. In some cases it is worth typing `-W` to avoid the extra connection attempt.

<code>--maintenance-db=</code><em>dbname</em>
:   When the `-a`/`--all` is used, connect to this database to gather the list of databases to vacuum. If not specified, the `postgres` database will be used, or if that does not exist, `template1` will be used. This can be a [connection string](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connstring). If so, connection string parameters will override any conflicting command line options. Also, connection string parameters other than the database name itself will be re-used when connecting to other databases.


## Environment


`PGDATABASE`, `PGHOST`, `PGPORT`, `PGUSER`
:   Default connection parameters

`PG_COLOR`
:   Specifies whether to use color in diagnostic messages. Possible values are `always`, `auto` and `never`.


 This utility, like most other PostgreSQL utilities, also uses the environment variables supported by libpq (see [Environment Variables](../../client-interfaces/libpq-c-library/environment-variables.md#libpq-envars)).


## Diagnostics


 In case of difficulty, see [sql-vacuum](../sql-commands/vacuum.md#sql-vacuum) and [app-psql](psql.md#app-psql) for discussions of potential problems and error messages. The database server must be running at the targeted host. Also, any default connection settings and environment variables used by the libpq front-end library will apply.


## Examples


 To clean the database `test`:

```

$ vacuumdb test
```


 To clean and analyze for the optimizer a database named `bigdb`:

```

$ vacuumdb --analyze bigdb
```


 To clean a single table `foo` in a database named `xyzzy`, and analyze a single column `bar` of the table for the optimizer:

```

$ vacuumdb --analyze --verbose --table='foo(bar)' xyzzy
```


 To clean all tables in the `foo` and `bar` schemas in a database named `xyzzy`:

```

$ vacuumdb --schema='foo' --schema='bar' xyzzy
```


## See Also
  [sql-vacuum](../sql-commands/vacuum.md#sql-vacuum)
