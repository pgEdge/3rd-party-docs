<a id="xact-locking"></a>

## Transactions and Locking


 The transaction IDs of currently executing transactions are shown in [`pg_locks`](../system-views/pg_locks.md#view-pg-locks) in columns `virtualxid` and `transactionid`. Read-only transactions will have `virtualxid`s but NULL `transactionid`s, while both columns will be set in read-write transactions.


 Some lock types wait on `virtualxid`, while other types wait on `transactionid`. Row-level read and write locks are recorded directly in the locked rows and can be inspected using the [pgrowlocks](../../appendixes/additional-supplied-modules-and-extensions/pgrowlocks-show-a-tables-row-locking-information.md#pgrowlocks) extension. Row-level read locks might also require the assignment of multixact IDs (`mxid`; see [Multixacts and Wraparound](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-multixact-wraparound)).
