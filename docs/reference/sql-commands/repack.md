<a id="sql-repack"></a>

# REPACK

rewrite a table to reclaim disk space

## Synopsis


```

REPACK [ ( OPTION [, ...] ) ] [ TABLE_AND_COLUMNS [ USING INDEX [ INDEX_NAME ] ] ]
REPACK [ ( OPTION [, ...] ) ] USING INDEX

where OPTION can be one of:

    VERBOSE [ BOOLEAN ]
    ANALYZE [ BOOLEAN ]
    CONCURRENTLY [ BOOLEAN ]

and TABLE_AND_COLUMNS is:

    TABLE_NAME [ ( COLUMN_NAME [, ...] ) ]
```


## Description


 `REPACK` reclaims storage occupied by dead tuples. Unlike `VACUUM`, it does so by rewriting the entire contents of the table specified by *table_name* into a new disk file with no extra space (except for the space guaranteed by the `fillfactor` storage parameter), allowing unused space to be returned to the operating system.


 Without a *table_name*, `REPACK` processes every table and materialized view in the current database that the current user has the `MAINTAIN` privilege on. This form of `REPACK` cannot be executed inside a transaction block. Also, this form is not allowed if the `CONCURRENTLY` option is used.


 If a `USING INDEX` clause is specified, the rows are physically reordered based on information from an index. Please see the notes on clustering below.


 When a table is being repacked, an `ACCESS EXCLUSIVE` lock is acquired on it. This prevents any other database operations (both reads and writes) from operating on the table until the `REPACK` is finished. If you want to keep the table accessible during the repacking, consider using the `CONCURRENTLY` option.
 <a id="sql-repack-notes-on-clustering"></a>

### Notes on Clustering


 If the `USING INDEX` clause is specified, the rows in the table are stored in the order that the index specifies; *clustering*, because rows are physically clustered afterwards. If an index name is specified in the command, the order implied by that index is used, and that index is configured as the index to cluster on. (This also applies to an index given to the `CLUSTER` command.) If no index name is specified, then the index that has been configured as the index to cluster on is used; an error is thrown if none has. An index can be set manually using `ALTER TABLE ... CLUSTER ON`, and reset with `ALTER TABLE ... SET WITHOUT CLUSTER`.


 If no table name is specified in `REPACK USING INDEX`, all tables which have a clustering index defined and which the calling user has privileges for are processed.


 Clustering is a one-time operation: when the table is subsequently updated, the changes are not clustered. That is, no attempt is made to store new or updated rows according to their index order. (If one wishes, one can periodically recluster by issuing the command again. Also, setting the table's `fillfactor` storage parameter to less than 100% can aid in preserving cluster ordering during updates, since updated rows are kept on the same page if enough space is available there.)


 In cases where you are accessing single rows randomly within a table, the actual order of the data in the table is unimportant. However, if you tend to access some data more than others, and there is an index that groups them together, you will benefit from using clustering. If you are requesting a range of indexed values from a table, or a single indexed value that has multiple rows that match, clustering will help because once the index identifies the table page for the first row that matches, all other rows that match are probably already on the same table page, and so you save disk accesses and speed up the query.


 `REPACK` can re-sort the table using either an index scan on the specified index (if the index is a b-tree), or a sequential scan followed by sorting. It will attempt to choose the method that will be faster, based on planner cost parameters and available statistical information.


 Because the planner records statistics about the ordering of tables, it is advisable to run [`ANALYZE`](analyze.md#sql-analyze) on the newly repacked table. Otherwise, the planner might make poor choices of query plans.
  <a id="sql-repack-notes-on-resources"></a>

### Notes on Resources


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

`CONCURRENTLY`
:   Allow other transactions to use the table while it is being repacked.


     Internally, `REPACK` copies the contents of the table (ignoring dead tuples) into a new file, sorted by the specified index, and also creates a new file for each index. Then it swaps the old and new files for the table and all the indexes, and deletes the old files. The `ACCESS EXCLUSIVE` lock is needed to make sure that the old files do not change during the processing because the changes would get lost due to the swap.


     With the `CONCURRENTLY` option, the `ACCESS EXCLUSIVE` lock is only acquired to swap the table and index files. The data changes that took place during the creation of the new table and index files are captured using logical decoding ([Logical Decoding](../../server-programming/logical-decoding/index.md#logicaldecoding)) and applied before the `ACCESS EXCLUSIVE` lock is requested. Thus the lock is typically held only for the time needed to swap the files, which should be pretty short. However, the time might still be noticeable if too many data changes have been done to the table while `REPACK` was waiting for the lock: those changes must be processed just before the files are swapped, while the `ACCESS EXCLUSIVE` lock is being held.


     Note that `REPACK` with the `CONCURRENTLY` option does not try to order the rows inserted into the table after the repacking started. Also note `REPACK` might fail to complete due to DDL commands executed on the table by other transactions during the repacking.


    !!! note

        In addition to the temporary space requirements explained in [Notes on Resources](#sql-repack-notes-on-resources), the `CONCURRENTLY` option can add to the usage of temporary space a bit more. The reason is that other transactions can perform DML operations which cannot be applied to the new file until `REPACK` has copied all the existing tuples from the old file. Thus the tuples inserted into the old file during the copying are also stored separately in a temporary file, until they can be processed.


     The `CONCURRENTLY` option cannot be used in the following cases:

    -  The table is `UNLOGGED`.
    -  The table is partitioned.
    -  The table lacks a primary key and index-based replica identity.
    -  The table is a system catalog or a TOAST table.
    -  `REPACK` is executed inside a transaction block.
    -  The [`max_repack_replication_slots`](../../server-administration/server-configuration/replication.md#guc-max-repack-replication-slots) configuration parameter does not allow for the creation of an additional replication slot.


    !!! warning

        `REPACK` with the `CONCURRENTLY` option is not MVCC-safe, see [Caveats](../../the-sql-language/concurrency-control/caveats.md#mvcc-caveats) for details.

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


 Repack the table `employees` on the basis of its index `employees_ind` (since an index is specified, this is effectively clustering):

```

REPACK employees USING INDEX employees_ind;
```


 Repack the `employees` table following the same index as was used before, in concurrent mode:

```

REPACK (CONCURRENTLY) employees USING INDEX;
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
