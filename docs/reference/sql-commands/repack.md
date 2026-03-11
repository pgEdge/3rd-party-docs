# REPACK { #sql-repack }

rewrite a table to reclaim disk space

## Synopsis


```

REPACK [ ( OPTION [, ...] ) ] [ TABLE_AND_COLUMNS [ USING INDEX [ INDEX_NAME ] ] ]
REPACK [ ( OPTION [, ...] ) ] USING INDEX

where OPTION can be one of:

    VERBOSE [ BOOLEAN ]
    ANALYZE [ BOOLEAN ]

and TABLE_AND_COLUMNS is:

    TABLE_NAME [ ( COLUMN_NAME [, ...] ) ]
```


## Description


 `REPACK` reclaims storage occupied by dead tuples. Unlike `VACUUM`, it does so by rewriting the entire contents of the table specified by *table_name* into a new disk file with no extra space (except for the space guaranteed by the `fillfactor` storage parameter), allowing unused space to be returned to the operating system.


 Without a *table_name*, `REPACK` processes every table and materialized view in the current database that the current user has the `MAINTAIN` privilege on. This form of `REPACK` cannot be executed inside a transaction block.


 If a `USING INDEX` clause is specified, the rows are physically reordered based on information from an index. Please see the notes on clustering below.


 When a table is being repacked, an `ACCESS EXCLUSIVE` lock is acquired on it. This prevents any other database operations (both reads and writes) from operating on the table until the `REPACK` is finished.


### Notes on Clustering { #sql-repack-notes-on-clustering }


 If the `USING INDEX` clause is specified, the rows in the table are stored in the order that the index specifies; *clustering*, because rows are physically clustered afterwards. If an index name is specified in the command, the order implied by that index is used, and that index is configured as the index to cluster on. (This also applies to an index given to the `CLUSTER` command.) If no index name is specified, then the index that has been configured as the index to cluster on is used; an error is thrown if none has. An index can be set manually using `ALTER TABLE ... CLUSTER ON`, and reset with `ALTER TABLE ... SET WITHOUT CLUSTER`.


 If no table name is specified in `REPACK USING INDEX`, all tables which have a clustering index defined and which the calling user has privileges for are processed.


 Clustering is a one-time operation: when the table is subsequently updated, the changes are not clustered. That is, no attempt is made to store new or updated rows according to their index order. (If one wishes, one can periodically recluster by issuing the command again. Also, setting the table's `fillfactor` storage parameter to less than 100% can aid in preserving cluster ordering during updates, since updated rows are kept on the same page if enough space is available there.)


 In cases where you are accessing single rows randomly within a table, the actual order of the data in the table is unimportant. However, if you tend to access some data more than others, and there is an index that groups them together, you will benefit from using clustering. If you are requesting a range of indexed values from a table, or a single indexed value that has multiple rows that match, clustering will help because once the index identifies the table page for the first row that matches, all other rows that match are probably already on the same table page, and so you save disk accesses and speed up the query.


 `REPACK` can re-sort the table using either an index scan on the specified index (if the index is a b-tree), or a sequential scan followed by sorting. It will attempt to choose the method that will be faster, based on planner cost parameters and available statistical information.


 Because the planner records statistics about the ordering of tables, it is advisable to run [`ANALYZE`](analyze.md#sql-analyze) on the newly repacked table. Otherwise, the planner might make poor choices of query plans.


### Notes on Resources { #sql-repack-notes-on-resources }


 When an index scan or a sequential scan without sort is used, a temporary copy of the table is created that contains the table data in the index order. Temporary copies of each index on the table are created as well. Therefore, you need free space on disk at least equal to the sum of the table size and the index sizes.


 When a sequential scan and sort is used, a temporary sort file is also created, so that the peak temporary space requirement is as much as double the table size, plus the index sizes. This method is often faster than the index scan method, but if the disk space requirement is intolerable, you can disable this choice by temporarily setting [enable_sort](../../server-administration/server-configuration/query-planning.md#guc-enable-sort) to `off`.


 It is advisable to set [maintenance_work_mem](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-work-mem) to a reasonably large value (but not more than the amount of RAM you can dedicate to the `REPACK` operation) before repacking.


## Parameters


*table_name*
:   The name (possibly schema-qualified) of a table.

*column_name*
:   The name of a specific column to analyze. Defaults to all columns. If a column list is specific, `ANALYZE` must also be specified.

*index_name*
:   The name of an index.

`VERBOSE`
:   Prints a progress report as each table is repacked at `INFO` level.

`ANALYZE`, `ANALYSE`
:   Applies [sql-analyze](analyze.md#sql-analyze) on the table after repacking. This is currently only supported when a single (non-partitioned) table is specified.

*boolean*
:   Specifies whether the selected option should be turned on or off. You can write `TRUE`, `ON`, or `1` to enable the option, and `FALSE`, `OFF`, or `0` to disable it. The *boolean* value can also be omitted, in which case `TRUE` is assumed.


## Notes


 To repack a table, one must have the `MAINTAIN` privilege on the table.


 While `REPACK` is running, the [search_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-search-path) is temporarily changed to `pg_catalog, pg_temp`.


 Each backend running `REPACK` will report its progress in the `pg_stat_progress_repack` view. See [REPACK Progress Reporting](../../server-administration/monitoring-database-activity/progress-reporting.md#repack-progress-reporting) for details.


 Repacking a partitioned table repacks each of its partitions. If an index is specified, each partition is repacked using the partition of that index. `REPACK` on a partitioned table cannot be executed inside a transaction block.


## Examples


 Repack the table `employees`:

```

REPACK employees;
```


 Repack the table `employees` on the basis of its index `employees_ind` (Since index is used here, this is effectively clustering):

```

REPACK employees USING INDEX employees_ind;
```


 Repack the table `cases` on physical ordering, running an `ANALYZE` on the given columns once repacking is done, showing informational messages:

```

REPACK (ANALYZE, VERBOSE) cases (district, case_nr);
```


 Repack all tables in the database on which you have the `MAINTAIN` privilege:

```

REPACK;
```


 Repack all tables for which a clustering index has previously been configured on which you have the `MAINTAIN` privilege, showing informational messages:

```

REPACK (VERBOSE) USING INDEX;
```


## Compatibility


 There is no `REPACK` statement in the SQL standard.


## See Also
  [REPACK Progress Reporting](../../server-administration/monitoring-database-activity/progress-reporting.md#repack-progress-reporting)
