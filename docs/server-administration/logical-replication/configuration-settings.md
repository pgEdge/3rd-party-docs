<a id="logical-replication-config"></a>

## Configuration Settings


 Logical replication requires several configuration options to be set. These options are relevant only on one side of the replication.
 <a id="logical-replication-config-publisher"></a>

### Publishers


 [`wal_level`](../server-configuration/write-ahead-log.md#guc-wal-level) must be set to `replica` or `logical`.


 [`max_replication_slots`](../server-configuration/replication.md#guc-max-replication-slots) must be set to at least the number of subscriptions expected to connect, plus some reserve for table synchronization.


 Logical replication slots are also affected by [`idle_replication_slot_timeout`](../server-configuration/replication.md#guc-idle-replication-slot-timeout).


 [`max_wal_senders`](../server-configuration/replication.md#guc-max-wal-senders) should be set to at least the same as `max_replication_slots`, plus the number of physical replicas that are connected at the same time.


 Logical replication walsender is also affected by [`wal_sender_timeout`](../server-configuration/replication.md#guc-wal-sender-timeout).
  <a id="logical-replication-config-subscriber"></a>

### Subscribers


 [`max_active_replication_origins`](../server-configuration/replication.md#guc-max-active-replication-origins) must be set to at least the number of subscriptions that will be added to the subscriber, plus some reserve for table synchronization.


 [`max_replication_slots`](../server-configuration/replication.md#guc-max-replication-slots) must be set to at least 1 when [`retain_dead_tuples`](../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-retain-dead-tuples) is enabled for any subscription.


 [`max_logical_replication_workers`](../server-configuration/replication.md#guc-max-logical-replication-workers) must be set to at least the number of subscriptions (for leader apply workers), plus some reserve for the parallel apply workers, and table/sequence synchronization workers.


 [`max_worker_processes`](../server-configuration/resource-consumption.md#guc-max-worker-processes) may need to be adjusted to accommodate for replication workers, at least ([`max_logical_replication_workers`](../server-configuration/replication.md#guc-max-logical-replication-workers) + `1`). Note, some extensions and parallel queries also take worker slots from `max_worker_processes`.


 [`max_sync_workers_per_subscription`](../server-configuration/replication.md#guc-max-sync-workers-per-subscription) controls how many tables can be synchronized in parallel during subscription initialization or when new tables are added. One additional worker is also needed for sequence synchronization.


 [`max_parallel_apply_workers_per_subscription`](../server-configuration/replication.md#guc-max-parallel-apply-workers-per-subscription) controls the amount of parallelism for streaming of in-progress transactions with subscription parameter `streaming = parallel`.


 Logical replication workers are also affected by [`wal_receiver_timeout`](../server-configuration/replication.md#guc-wal-receiver-timeout), [`wal_receiver_status_interval`](../server-configuration/replication.md#guc-wal-receiver-status-interval) and [`wal_retrieve_retry_interval`](../server-configuration/replication.md#guc-wal-retrieve-retry-interval).
