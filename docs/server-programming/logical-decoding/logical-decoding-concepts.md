<a id="logicaldecoding-explanation"></a>

## Logical Decoding Concepts
  <a id="logicaldecoding-explanation-log-dec"></a>

### Logical Decoding


 Logical decoding is the process of extracting all persistent changes to a database's tables into a coherent, easy to understand format which can be interpreted without detailed knowledge of the database's internal state.


 In PostgreSQL, logical decoding is implemented by decoding the contents of the [write-ahead log](../../server-administration/reliability-and-the-write-ahead-log/index.md#wal), which describe changes on a storage level, into an application-specific form such as a stream of tuples or SQL statements.


 Logical decoding becomes available in two conditions:


-  When [wal_level](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-level) is set to `logical`.
-  When [wal_level](../../server-administration/server-configuration/write-ahead-log.md#guc-wal-level) is set to `replica` and at least one valid logical replication slot exists on the system.


 If either condition is met, the operational WAL level becomes equivalent to `logical`, which can be monitored through the [effective_wal_level](../../server-administration/server-configuration/preset-options.md#guc-effective-wal-level) parameter.


 When `wal_level` is set to `replica`, logical decoding is automatically activated upon creation of the first logical replication slot. This activation process involves several steps and requires synchronization among processes, ensuring system-wide consistency. Conversely, if `wal_level` is set to `replica` and the last logical replication slot is dropped or invalidated, logical decoding is automatically disabled. Note that the deactivation of logical decoding might take some time as it is performed asynchronously by the checkpointer process.


!!! caution

    When `wal_level` is set to `replica`, dropping or invalidating the last logical slot disables logical decoding on the primary, resulting in slots on standbys being invalidated.
  <a id="logicaldecoding-replication-slots"></a>

### Replication Slots


 In the context of logical replication, a slot represents a stream of changes that can be replayed to a client in the order they were made on the origin server. Each slot streams a sequence of changes from a single database.


!!! note

    PostgreSQL can also use streaming replication slots to maintain a standby server (see [Streaming Replication](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication)), but typically those use physical replication, not logical.


 A replication slot has an identifier that is unique across all databases in a PostgreSQL cluster. Slots persist independently of the connection using them and are crash-safe.


 A logical slot will emit each change just once in normal operation. The current position of each slot is persisted only at checkpoint, so in the case of a crash the slot might return to an earlier LSN, which will then cause recent changes to be sent again when the server restarts. Logical decoding clients are responsible for avoiding ill effects from handling the same message more than once. Clients may wish to record the last LSN they saw when decoding and skip over any repeated data or (when using the replication protocol) request that decoding start from that LSN rather than letting the server determine the start point. The Replication Progress Tracking feature is designed for this purpose, refer to [replication origins](../replication-progress-tracking.md#replication-origins).


 Multiple independent slots may exist for a single database. Each slot has its own state, allowing different consumers to receive changes from different points in the database change stream. For most applications, a separate slot will be required for each consumer.


 A logical replication slot knows nothing about the state of the receiver(s). It's even possible to have multiple different receivers using the same slot at different times; they'll just get the changes following on from when the last receiver stopped consuming them. Only one receiver may consume changes from a slot at any given time.


 A logical replication slot can also be created on a hot standby. To prevent `VACUUM` from removing required rows from the system catalogs, `hot_standby_feedback` should be set on the standby. In spite of that, if any required rows get removed, the slot gets invalidated. It's highly recommended to use a physical slot between the primary and the standby. Otherwise, `hot_standby_feedback` will work but only while the connection is alive (for example a node restart would break it). Then, the primary may delete system catalog rows that could be needed by the logical decoding on the standby (as it does not know about the `catalog_xmin` on the standby). Existing logical slots on standby also get invalidated if `effective_wal_level` on the primary is reduced to less than `logical`. This is done as soon as the standby detects such a change in the WAL stream. It means that, for walsenders that are lagging (if any), some WAL records up to the `wal_level` parameter change on the primary won't be decoded.


 Creation of a logical slot requires information about all the currently running transactions. On the primary, this information is available directly, but on a standby, this information has to be obtained from primary. Thus, slot creation may need to wait for some activity to happen on the primary. If the primary is idle, creating a logical slot on standby may take noticeable time. This can be sped up by calling the `pg_log_standby_snapshot` function on the primary.


!!! caution

    Replication slots persist across crashes and know nothing about the state of their consumer(s). They will prevent removal of required resources even when there is no connection using them. This consumes storage because neither required WAL nor required rows from the system catalogs can be removed by `VACUUM` as long as they are required by a replication slot. In extreme cases this could cause the database to shut down to prevent transaction ID wraparound (see [Preventing Transaction ID Wraparound Failures](../../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#vacuum-for-wraparound)). So if a slot is no longer required it should be dropped.
  <a id="logicaldecoding-replication-slots-synchronization"></a>

### Replication Slot Synchronization


 The logical replication slots on the primary can be synchronized to the hot standby by using the `failover` parameter of [`pg_create_logical_replication_slot`](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-create-logical-replication-slot), or by using the [`failover`](../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-failover) option of `CREATE SUBSCRIPTION` during slot creation. Additionally, enabling [`sync_replication_slots`](../../server-administration/server-configuration/replication.md#guc-sync-replication-slots) on the standby is required. By enabling [`sync_replication_slots`](../../server-administration/server-configuration/replication.md#guc-sync-replication-slots) on the standby, the failover slots can be synchronized periodically in the slotsync worker. For the synchronization to work, it is mandatory to have a physical replication slot between the primary and the standby (i.e., [`primary_slot_name`](../../server-administration/server-configuration/replication.md#guc-primary-slot-name) should be configured on the standby), and [`hot_standby_feedback`](../../server-administration/server-configuration/replication.md#guc-hot-standby-feedback) must be enabled on the standby. It is also necessary to specify a valid `dbname` in the [`primary_conninfo`](../../server-administration/server-configuration/replication.md#guc-primary-conninfo). It's highly recommended that the said physical replication slot is named in [`synchronized_standby_slots`](../../server-administration/server-configuration/replication.md#guc-synchronized-standby-slots) list on the primary, to prevent the subscriber from consuming changes faster than the hot standby. Even when correctly configured, some latency is expected when sending changes to logical subscribers due to the waiting on slots named in [`synchronized_standby_slots`](../../server-administration/server-configuration/replication.md#guc-synchronized-standby-slots). When `synchronized_standby_slots` is utilized, the primary server will not completely shut down until the corresponding standbys, associated with the physical replication slots specified in `synchronized_standby_slots`, have confirmed receiving the WAL up to the latest flushed position on the primary server.


!!! note

    While enabling [`sync_replication_slots`](../../server-administration/server-configuration/replication.md#guc-sync-replication-slots) allows for automatic periodic synchronization of failover slots, they can also be manually synchronized using the [`pg_sync_replication_slots`](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-sync-replication-slots) function on the standby. However, unlike automatic synchronization, it does not perform incremental updates. It retries cyclically until all the failover slots that existed on primary at the start of the function call are synchronized. Any slots created after the function begins will not be synchronized. In contrast, automatic synchronization via `sync_replication_slots` provides continuous slot updates, enabling seamless failover and supporting high availability. Therefore, it is the recommended method for synchronizing slots.


 When slot synchronization is configured as recommended, and the initial synchronization is performed either automatically or manually via `pg_sync_replication_slots`, the standby can persist the synchronized slot only if the following condition is met: The logical replication slot on the primary must retain WALs and system catalog rows that are still available on the standby. This ensures data integrity and allows logical replication to continue smoothly after promotion. If the required WALs or catalog rows have already been purged from the standby, the slot will not be persisted to avoid data loss. In such cases, the following log message may appear:

```

LOG:  could not synchronize replication slot "failover_slot"
DETAIL:  Synchronization could lead to data loss, because the remote slot needs WAL at LSN 0/03003F28 and catalog xmin 754, but the standby has LSN 0/03003F28 and catalog xmin 756.
```
 If the logical replication slot is actively used by a consumer, no manual intervention is needed; the slot will advance automatically, and synchronization will resume in the next cycle. However, if no consumer is configured, it is advisable to manually advance the slot on the primary using [`pg_logical_slot_get_changes`](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-slot-get-changes) or [`pg_logical_slot_get_binary_changes`](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-slot-get-binary-changes), allowing synchronization to proceed.


 The ability to resume logical replication after failover depends upon the [pg_replication_slots](../../internals/system-views/pg_replication_slots.md#view-pg-replication-slots).`synced` value for the synchronized slots on the standby at the time of failover. Only persistent slots that have attained synced state as true on the standby before failover can be used for logical replication after failover. Temporary synced slots cannot be used for logical decoding, therefore logical replication for those slots cannot be resumed. For example, if the synchronized slot could not become persistent on the standby due to a disabled subscription, then the subscription cannot be resumed after failover even when it is enabled.


 To resume logical replication after failover from the synced logical slots, the subscription's 'conninfo' must be altered to point to the new primary server. This is done using [`ALTER SUBSCRIPTION ... CONNECTION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-connection). It is recommended that subscriptions are first disabled before promoting the standby and are re-enabled after altering the connection string.


!!! caution

    There is a chance that the old primary is up again during the promotion and if subscriptions are not disabled, the logical subscribers may continue to receive data from the old primary server even after promotion until the connection string is altered. This might result in data inconsistency issues, preventing the logical subscribers from being able to continue replication from the new primary server.
  <a id="logicaldecoding-explanation-output-plugins"></a>

### Output Plugins


 Output plugins transform the data from the write-ahead log's internal representation into the format the consumer of a replication slot desires.
  <a id="logicaldecoding-explanation-exported-snapshots"></a>

### Exported Snapshots


 When a new replication slot is created using the streaming replication interface (see [CREATE_REPLICATION_SLOT](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication-create-replication-slot)), a snapshot is exported (see [Snapshot Synchronization Functions](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-snapshot-synchronization)), which will show exactly the state of the database after which all changes will be included in the change stream. This can be used to create a new replica by using [`SET TRANSACTION SNAPSHOT`](../../reference/sql-commands/set-transaction.md#sql-set-transaction) to read the state of the database at the moment the slot was created. This transaction can then be used to dump the database's state at that point in time, which afterwards can be updated using the slot's contents without losing any changes.


 Applications that do not require snapshot export may suppress it with the `SNAPSHOT 'nothing'` option.
