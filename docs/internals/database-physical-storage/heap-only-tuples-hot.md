<a id="storage-hot"></a>

## Heap-Only Tuples (HOT)


 To allow for high concurrency, PostgreSQL uses [multiversion concurrency control](../../the-sql-language/concurrency-control/introduction.md#mvcc-intro) (MVCC) to store rows. However, MVCC has some downsides for update queries. Specifically, updates require new versions of rows to be added to tables. This can also require new index entries for each updated row, and removal of old versions of rows and their index entries can be expensive.


 To help reduce the overhead of updates, PostgreSQL has an optimization called heap-only tuples (HOT). This optimization is possible when:

-  The update does not modify any columns referenced by the table's indexes, not including summarizing indexes. The only summarizing index method in the core PostgreSQL distribution is [BRIN](../built-in-index-access-methods/brin-indexes.md#brin).
-  There is sufficient free space on the page containing the old row for the updated row.
 In such cases, heap-only tuples provide two optimizations:

-  New index entries are not needed to represent updated rows, however, summary indexes may still need to be updated.
-  When a row is updated multiple times, row versions other than the oldest and the newest can be completely removed during normal operation, including `SELECT`s, instead of requiring periodic vacuum operations. (Indexes always refer to the [page item identifier](database-page-layout.md#storage-page-layout) of the original row version. The tuple data associated with that row version is removed, and its item identifier is converted to a redirect that points to the oldest version that may still be visible to some concurrent transaction. Intermediate row versions that are no longer visible to anyone are completely removed, and the associated page item identifiers are made available for reuse.)


 You can increase the likelihood of sufficient page space for HOT updates by decreasing a table's [`fillfactor`](../../reference/sql-commands/create-table.md#reloption-fillfactor). If you don't, HOT updates will still happen because new rows will naturally migrate to new pages and existing pages with sufficient free space for new row versions. The system view [pg_stat_all_tables](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-all-tables-view) allows monitoring of the occurrence of HOT and non-HOT updates.
