<a id="logical-replication-config"></a>

## Configuration Settings


 Logical replication requires several configuration options to be set. Most options are relevant only on one side of the replication. However, `max_replication_slots` is used on both the publisher and the subscriber, but it has a different meaning for each.
 <a id="logical-replication-config-publisher"></a>

### Publishers


 [`wal_level`](../server-configuration/write-ahead-log.md#guc-wal-level) must be set to `logical`.


 [`max_replication_slots`](../server-configuration/replication.md#guc-max-replication-slots) must be set to at least the number of subscriptions expected to connect, plus some reserve for table synchronization.


 [`max_wal_senders`](../server-configuration/replication.md#guc-max-wal-senders) should be set to at least the same as `max_replication_slots`, plus the number of physical replicas that are connected at the same time.


 Logical replication walsender is also affected by [`wal_sender_timeout`](../server-configuration/replication.md#guc-wal-sender-timeout).
  <a id="logical-replication-config-subscriber"></a>

### Subscribers


 [`max_replication_slots`](../server-configuration/replication.md#guc-max-replication-slots-subscriber) must be set to at least the number of subscriptions that will be added to the subscriber, plus some reserve for table synchronization.


 [`max_logical_replication_workers`](../server-configuration/replication.md#guc-max-logical-replication-workers) must be set to at least the number of subscriptions (for leader apply workers), plus some reserve for the table synchronization workers and parallel apply workers.


 [`max_worker_processes`](../server-configuration/resource-consumption.md#guc-max-worker-processes) may need to be adjusted to accommodate for replication workers, at least ([`max_logical_replication_workers`](../server-configuration/replication.md#guc-max-logical-replication-workers) + `1`). Note, some extensions and parallel queries also take worker slots from `max_worker_processes`.


 [`max_sync_workers_per_subscription`](../server-configuration/replication.md#guc-max-sync-workers-per-subscription) controls the amount of parallelism of the initial data copy during the subscription initialization or when new tables are added.


 [`max_parallel_apply_workers_per_subscription`](../server-configuration/replication.md#guc-max-parallel-apply-workers-per-subscription) controls the amount of parallelism for streaming of in-progress transactions with subscription parameter `streaming = parallel`.


 Logical replication workers are also affected by [`wal_receiver_timeout`](../server-configuration/replication.md#guc-wal-receiver-timeout), [`wal_receiver_status_interval`](../server-configuration/replication.md#guc-wal-receiver-status-interval) and [`wal_retrieve_retry_interval`](../server-configuration/replication.md#guc-wal-retrieve-retry-interval).
