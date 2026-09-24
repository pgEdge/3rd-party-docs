<a id="views-overview"></a>

## Overview


 [System Views](#view-table) lists the system views. More detailed documentation of each catalog follows below. Except where noted, all the views described here are read-only.
 <a id="view-table"></a>

**Table: System Views**

| View Name | Purpose |
| --- | --- |
| [`pg_aios`](pg_aios.md#view-pg-aios) | in-use asynchronous IO handles |
| [`pg_available_extensions`](pg_available_extensions.md#view-pg-available-extensions) | available extensions |
| [`pg_available_extension_versions`](pg_available_extension_versions.md#view-pg-available-extension-versions) | available versions of extensions |
| [`pg_backend_memory_contexts`](pg_backend_memory_contexts.md#view-pg-backend-memory-contexts) | backend memory contexts |
| [`pg_config`](pg_config.md#view-pg-config) | compile-time configuration parameters |
| [`pg_cursors`](pg_cursors.md#view-pg-cursors) | open cursors |
| [`pg_dsm_registry_allocations`](pg_dsm_registry_allocations.md#view-pg-dsm-registry-allocations) | shared memory allocations tracked in the DSM registry |
| [`pg_file_settings`](pg_file_settings.md#view-pg-file-settings) | summary of configuration file contents |
| [`pg_group`](pg_group.md#view-pg-group) | groups of database users |
| [`pg_hba_file_rules`](pg_hba_file_rules.md#view-pg-hba-file-rules) | summary of client authentication configuration file contents |
| [`pg_ident_file_mappings`](pg_ident_file_mappings.md#view-pg-ident-file-mappings) | summary of client user name mapping configuration file contents |
| [`pg_indexes`](pg_indexes.md#view-pg-indexes) | indexes |
| [`pg_locks`](pg_locks.md#view-pg-locks) | locks currently held or awaited |
| [`pg_matviews`](pg_matviews.md#view-pg-matviews) | materialized views |
| [`pg_policies`](pg_policies.md#view-pg-policies) | policies |
| [`pg_prepared_statements`](pg_prepared_statements.md#view-pg-prepared-statements) | prepared statements |
| [`pg_prepared_xacts`](pg_prepared_xacts.md#view-pg-prepared-xacts) | prepared transactions |
| [`pg_publication_sequences`](pg_publication_sequences.md#view-pg-publication-sequences) | publications and information of their associated sequences |
| [`pg_publication_tables`](pg_publication_tables.md#view-pg-publication-tables) | publications and information of their associated tables |
| [`pg_replication_origin_status`](pg_replication_origin_status.md#view-pg-replication-origin-status) | information about replication origins, including replication progress |
| [`pg_replication_slots`](pg_replication_slots.md#view-pg-replication-slots) | replication slot information |
| [`pg_roles`](pg_roles.md#view-pg-roles) | database roles |
| [`pg_rules`](pg_rules.md#view-pg-rules) | rules |
| [`pg_seclabels`](pg_seclabels.md#view-pg-seclabels) | security labels |
| [`pg_sequences`](pg_sequences.md#view-pg-sequences) | sequences |
| [`pg_settings`](pg_settings.md#view-pg-settings) | parameter settings |
| [`pg_shadow`](pg_shadow.md#view-pg-shadow) | database users |
| [`pg_shmem_allocations`](pg_shmem_allocations.md#view-pg-shmem-allocations) | shared memory allocations |
| [`pg_shmem_allocations_numa`](pg_shmem_allocations_numa.md#view-pg-shmem-allocations-numa) | NUMA node mappings for shared memory allocations |
| [`pg_stats`](pg_stats.md#view-pg-stats) | planner statistics |
| [`pg_stats_ext`](pg_stats_ext.md#view-pg-stats-ext) | extended planner statistics |
| [`pg_stats_ext_exprs`](pg_stats_ext_exprs.md#view-pg-stats-ext-exprs) | extended planner statistics for expressions |
| [`pg_tables`](pg_tables.md#view-pg-tables) | tables |
| [`pg_timezone_abbrevs`](pg_timezone_abbrevs.md#view-pg-timezone-abbrevs) | time zone abbreviations |
| [`pg_timezone_names`](pg_timezone_names.md#view-pg-timezone-names) | time zone names |
| [`pg_user`](pg_user.md#view-pg-user) | database users |
| [`pg_user_mappings`](pg_user_mappings.md#view-pg-user-mappings) | user mappings |
| [`pg_views`](pg_views.md#view-pg-views) | views |
| [`pg_wait_events`](pg_wait_events.md#view-pg-wait-events) | wait events |
