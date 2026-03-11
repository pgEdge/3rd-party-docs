# CLUSTER { #sql-cluster }

cluster a table according to an index

## Synopsis


```

CLUSTER [ ( OPTION [, ...] ) ] [ TABLE_NAME [ USING INDEX_NAME ] ]

where OPTION can be one of:

    VERBOSE [ BOOLEAN ]
```


## Description


 The `CLUSTER` command is equivalent to [sql-repack](repack.md#sql-repack) with an `USING INDEX` clause. See there for more details.


## Parameters


*table_name*
:   The name (possibly schema-qualified) of a table.

*index_name*
:   The name of an index.

`VERBOSE`
:   Prints a progress report as each table is clustered at `INFO` level.

*boolean*
:   Specifies whether the selected option should be turned on or off. You can write `TRUE`, `ON`, or `1` to enable the option, and `FALSE`, `OFF`, or `0` to disable it. The *boolean* value can also be omitted, in which case `TRUE` is assumed.


## Notes


 To cluster a table, one must have the `MAINTAIN` privilege on the table.


 While `CLUSTER` is running, the [search_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-search-path) is temporarily changed to `pg_catalog, pg_temp`.


 Because `CLUSTER` remembers which indexes are clustered, one can cluster the tables one wants clustered manually the first time, then set up a periodic maintenance script that executes `CLUSTER` without any parameters, so that the desired tables are periodically reclustered.


 Each backend running `CLUSTER` will report its progress in the `pg_stat_progress_cluster` view. See [CLUSTER Progress Reporting](../../server-administration/monitoring-database-activity/progress-reporting.md#cluster-progress-reporting) for details.


 Clustering a partitioned table clusters each of its partitions using the partition of the specified partitioned index. When clustering a partitioned table, the index may not be omitted. `CLUSTER` on a partitioned table cannot be executed inside a transaction block.


## Examples


 Cluster the table `employees` on the basis of its index `employees_ind`:

```

CLUSTER employees USING employees_ind;
```


 Cluster the `employees` table using the same index that was used before:

```

CLUSTER employees;
```


 Cluster all tables in the database that have previously been clustered:

```

CLUSTER;
```


## Compatibility


 There is no `CLUSTER` statement in the SQL standard.


 The following syntax was used before PostgreSQL 17 and is still supported:

```

CLUSTER [ VERBOSE ] [ TABLE_NAME [ USING INDEX_NAME ] ]
```


 The following syntax was used before PostgreSQL 8.3 and is still supported:

```

CLUSTER INDEX_NAME ON TABLE_NAME
```


## See Also
  [sql-repack](repack.md#sql-repack), [app-clusterdb](../postgresql-client-applications/clusterdb.md#app-clusterdb), [CLUSTER Progress Reporting](../../server-administration/monitoring-database-activity/progress-reporting.md#cluster-progress-reporting)
