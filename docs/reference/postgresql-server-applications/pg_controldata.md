<a id="app-pgcontroldata"></a>

# pg_controldata

display control information of a PostgreSQL database cluster

## Synopsis


```
pg_controldata [OPTION] [
     -D
     --pgdata
     | DATADIR]
```
 <a id="r1-app-pgcontroldata-1"></a>

## Description


 `pg_controldata` prints information initialized during `initdb`, such as the catalog version. It also shows information about write-ahead logging and checkpoint processing. This information is cluster-wide, and not specific to any one database.


 This utility can only be run by the user who initialized the cluster because it requires read access to the data directory. You can specify the data directory on the command line, or use the environment variable `PGDATA`.


## Options


<code>-D </code><em>datadir</em>, <code>--pgdata=</code><em>datadir</em>
:   Specifies the directory where the database cluster is stored.

`-V`, `--version`
:   Print the pg_controldata version and exit.

`-?`, `--help`
:   Show help about pg_controldata command line arguments, and exit.


## Environment


`PGDATA`
:   Default data directory location

`PG_COLOR`
:   Specifies whether to use color in diagnostic messages. Possible values are `always`, `auto` and `never`.
