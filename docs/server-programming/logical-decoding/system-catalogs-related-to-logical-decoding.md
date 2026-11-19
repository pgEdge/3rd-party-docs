<a id="logicaldecoding-catalogs"></a>

## System Catalogs Related to Logical Decoding


 The [`pg_replication_slots`](../../internals/system-views/pg_replication_slots.md#view-pg-replication-slots) view and the [`pg_stat_replication`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-replication-view) view provide information about the current state of replication slots and streaming replication connections respectively. These views apply to both physical and logical replication. The [`pg_stat_replication_slots`](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-pg-stat-replication-slots-view) view provides statistics information about the logical replication slots.
