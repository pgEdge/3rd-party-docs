<a id="non-durability"></a>

## Non-Durable Settings


 Durability is a database feature that guarantees the recording of committed transactions even if the server crashes or loses power. However, durability adds significant database overhead, so if your site does not require such a guarantee, PostgreSQL can be configured to run much faster. The following are configuration changes you can make to improve performance in such cases. Except as noted below, durability is still guaranteed in case of a crash of the database software; only an abrupt operating system crash creates a risk of data loss or corruption when these settings are used.

-  Place the database cluster's data directory in a memory-backed file system (i.e., RAM disk). This eliminates all database disk I/O, but limits data storage to the amount of available memory (and perhaps swap).
-  Turn off [fsync](../../server-administration/server-configuration/write-ahead-log.md#guc-fsync); there is no need to flush data to disk.
-  Turn off [synchronous_commit](../../server-administration/server-configuration/write-ahead-log.md#guc-synchronous-commit); there might be no need to force WAL writes to disk on every commit. This setting does risk transaction loss (though not data corruption) in case of a crash of the *database*.
-  Turn off [full_page_writes](../../server-administration/server-configuration/write-ahead-log.md#guc-full-page-writes); there is no need to guard against partial page writes.
-  Increase [max_wal_size](../../server-administration/server-configuration/write-ahead-log.md#guc-max-wal-size) and [checkpoint_timeout](../../server-administration/server-configuration/write-ahead-log.md#guc-checkpoint-timeout); this reduces the frequency of checkpoints, but increases the storage requirements of `/pg_wal`.
-  Create [unlogged tables](../../reference/sql-commands/create-table.md#sql-createtable-unlogged) to avoid WAL writes, though it makes the tables non-crash-safe.
