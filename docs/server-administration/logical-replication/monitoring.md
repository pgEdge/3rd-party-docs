<a id="logical-replication-monitoring"></a>

## Monitoring


 Because streaming logical replication is based on a similar architecture as [streaming physical replication](../high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication), the monitoring on a publication node is similar to monitoring of a physical replication primary (see [Monitoring](../high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication-monitoring)).


 The monitoring information about subscription is visible in [`pg_stat_subscription`](../monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-subscription). This view contains one row for every subscription worker. A subscription can have zero or more active subscription workers depending on its state.


 Normally, there is a single apply process running for an enabled subscription. A disabled subscription or a crashed subscription will have zero rows in this view. If the initial data synchronization of any table is in progress, there will be additional workers for the tables being synchronized. Moreover, if the [`streaming`](../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-streaming) transaction is applied in parallel, there may be additional parallel apply workers.
