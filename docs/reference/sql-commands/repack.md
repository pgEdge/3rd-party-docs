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


 If a `USING INDEX` clause is specified, the rows are physically reordered based on information from an index. Please see [Notes on Clustering](#sql-repack-notes-on-clustering) below.


 When a table is being repacked, an `ACCESS EXCLUSIVE` lock is acquired on it, unless the `CONCURRENTLY` option is given. See [Notes on Concurrent Operation](#sql-repack-notes-on-concurrently) for a discussion on the effects of this option.


## Parameters


*table_name*
:   The name (possibly schema-qualified) of a table.

*column_name*
:   The name of a specific column to analyze. Defaults to all columns. If a column list is specific, `ANALYZE` must also be specified.

*index_name*
:   The name of an index.

`CONCURRENTLY`
:   Allow other transactions to use the table while it is being repacked. See [Notes on Concurrent Operation](#sql-repack-notes-on-concurrently) for more details.

`VERBOSE`
:   Prints a progress report as each table is repacked at `INFO` level.

`ANALYZE`, `ANALYSE`
:   Applies [sql-analyze](analyze.md#sql-analyze) on the table after repacking. This is currently only supported when a single (non-partitioned) table is specified. This option cannot be used inside a transaction block, or from a function, procedure, or `DO` block.

*boolean*
:   Specifies whether the selected option should be turned on or off. You can write `TRUE`, `ON`, or `1` to enable the option, and `FALSE`, `OFF`, or `0` to disable it. The *boolean* value can also be omitted, in which case `TRUE` is assumed.
 <a id="sql-repack-notes"></a>

## Notes


 To repack a table, one must have the `MAINTAIN` privilege on the table.


 `REPACK` is primarily meant to remove bloat and, with `USING INDEX`, to cluster the table, in order to improve performance. Although it also advances the table's `relfrozenxid` and `relminmxid`, it is not a good way to prevent transaction ID or multixact ID wraparound: it rewrites the whole table and all of its indexes, so it takes much longer than `VACUUM`, and it can fail after most of the work is done, especially with `CONCURRENTLY` on a busy table. `VACUUM` is recommended for that instead (see [Preventing Transaction ID Wraparound Failures](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound)). When a table gets close to wraparound, `VACUUM` skips index vacuuming on its own (see [vacuum_failsafe_age](../../server-administration/server-configuration/vacuuming.md#guc-vacuum-failsafe-age)), and its [`INDEX_CLEANUP OFF`](vacuum.md#sql-vacuum) option can do the same earlier.


 `REPACK` refuses to process a table on which an invalid index exists. Such indexes must be dropped or reindexed by the user ahead of time.


 While `REPACK` is running, the [search_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-search-path) is temporarily changed to `pg_catalog, pg_temp`.


 Each backend running `REPACK` will report its progress in the `pg_stat_progress_repack` view. See [REPACK Progress Reporting](../../server-administration/monitoring-database-activity/progress-reporting.md#repack-progress-reporting) for details.


 Repacking a partitioned table repacks each of its partitions. If an index is specified, each partition is repacked using the partition of that index. `REPACK` on a partitioned table cannot be executed inside a transaction block.
 <a id="sql-repack-notes-on-clustering"></a>

### Notes on Clustering


 If the `USING INDEX` clause is specified, the rows in the table are physically rearranged according to the ordering implied by the specified index; this is known as *clustering*. This can have performance implications: in cases where you are accessing single rows randomly within a table, the actual order of the data in the table is unimportant. However, if you tend to access some data more than others, and there is an index that groups them together, you will benefit from using clustering. If you are requesting a range of indexed values from a table, or a single indexed value that has multiple matching rows, clustering will help because once the index identifies the table page for the first row that matches, all other rows that match are probably already on the same table page, and so you save disk accesses and speed up the query.


 For B-Tree indexes, the ordering used by clustering is the index's linear sort order. Other clusterable index access methods may use a different ordering strategy, one which may not necessarily correspond to any SQL sort order. If an index name is specified in the command, that index is used and is recorded as the table's clustering index. (This also applies to an index given to the `CLUSTER` command.) If no index name is specified, then the index that has been configured as the clustering one is used; if none has been configured, an error is thrown. An index can be set manually using `ALTER TABLE ... CLUSTER ON`, and reset with `ALTER TABLE ... SET WITHOUT CLUSTER`.


 Clustering is a one-time operation: when the table is subsequently updated, the changes are not clustered. That is, no attempt is made to store new or updated rows according to the clustering order. (If one wishes, one can periodically recluster by issuing the command again. Also, setting the table's `fillfactor` storage parameter to less than 100% can aid in preserving cluster ordering during updates, since updated rows are kept on the same page if enough space is available there.)


 When clustering on a B-Tree index, `REPACK` can rewrite the table using either an index scan on the specified index, or a sequential scan followed by sorting. It will attempt to choose the method that will be faster, based on planner cost parameters and available statistical information. When clustering on an index of an access method other than B-Tree, `REPACK` always uses an index scan.


 Because the planner records statistics about the ordering of tables, it is advisable to specify the `ANALYZE` option, or to run [`ANALYZE`](analyze.md#sql-analyze) on the newly repacked table. Otherwise, the planner might make poor choices of query plans.


 If no table name is specified in `REPACK USING INDEX`, all tables which have a clustering index defined and which the calling user has privileges for are processed.
  <a id="sql-repack-notes-on-resources"></a>

### Notes on Resources


 `REPACK` creates a temporary copy of the table and all of its indexes, so you need free disk space at least equal to the sum of the table size and the index sizes.


 If the `USING INDEX` clause is given and the planner chooses a sequential scan and sort to implement it, a temporary sort file of up to the size of the table is needed in addition to the above. This method is usually faster than an index scan, but if the disk space requirement is intolerable, you can disable this choice by temporarily setting [enable_sort](../../server-administration/server-configuration/query-planning.md#guc-enable-sort) to `off`; or omitting `USING INDEX` if the table doesn't need to be reordered.


 For the purposes of transaction ID wraparound (see [Preventing Transaction ID Wraparound Failures](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound)), `REPACK` is considered a single long-running transaction, which prevents `VACUUM` from cleaning up dead rows from other tables. It is advisable to monitor `pg_stat_activity.backend_xid` for the process running `REPACK` when repacking very large tables, to avoid causing excessive bloat in other tables.


 It is advisable to set [maintenance_work_mem](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-work-mem) to a reasonably large value (but not more than the amount of RAM you can dedicate to the `REPACK` operation) before repacking.
  <a id="sql-repack-notes-on-concurrently"></a>

### Notes on Concurrent Operation


!!! warning

    `REPACK` with the `CONCURRENTLY` option is not MVCC-safe. After the repacking transaction commits, the table will appear empty to concurrent transactions, if they are using a snapshot taken before repacking started. See [Caveats](../../the-sql-language/concurrency-control/caveats.md#mvcc-caveats) for more details.


 `REPACK` copies the contents of the table (ignoring dead tuples) into a new file, sorted by the specified index, and also creates a new file for each index. It then swaps the old and new files for the table and all the indexes, and deletes the old files. Without the `CONCURRENTLY` option, an `ACCESS EXCLUSIVE` lock is acquired at the beginning and remains held throughout the operation to make sure that the old files do not change during the processing; otherwise, any concurrent changes would get lost due to the swap.


 By contrast, when the `CONCURRENTLY` option is specified, the bulk of the operation is run with `SHARE UPDATE EXCLUSIVE` lock, and the `ACCESS EXCLUSIVE` lock is only acquired during the final phase to swap the table and index files. The data changes that took place during the creation of the new table and index files are captured using logical decoding (see [Logical Decoding](../../server-programming/logical-decoding/index.md#logicaldecoding)) and applied before the `ACCESS EXCLUSIVE` lock is requested. Once that lock is obtained, a final pass over any remaining captured concurrent data changes is done and the files are swapped. Thus the lock is typically held for a short time, depending on the amount of changes accumulated while the lock was being waited for.


 Processing of these concurrent changes requires a fixed small amount of memory for each tuple concurrently updated or deleted, not limited by [maintenance_work_mem](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-work-mem); if more than about 104 million rows are concurrently updated or deleted during the execution of `REPACK`, the command fails. In addition, the complete tuple data for all concurrent changes (except for TOAST columns that are not modified by `UPDATE`) are stored on disk separately.


 `REPACK (CONCURRENTLY)` might fail to complete if DDL commands are executed on the table by other transactions during the repacking.


 `REPACK (CONCURRENTLY) USING INDEX` does not try to order the rows inserted into the table after the repacking started.


 The `CONCURRENTLY` option cannot be used in the following cases:

-  The relation is not a table (e.g., it is a materialized view).
-  The table is `UNLOGGED`.
-  The table is partitioned.
-  The table lacks a primary key and index-based replica identity.
-  The table is a system catalog or a TOAST table.
-  The table's access method is not `heap`.
-  The table is declared as a catalog table using the [`user_catalog_table`](create-table.md#reloption-user-catalog-table) storage parameter.
-  `REPACK` is executed inside a transaction block.
-  The [`max_repack_replication_slots`](../../server-administration/server-configuration/replication.md#guc-max-repack-replication-slots) configuration parameter does not allow for the creation of an additional replication slot.


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
