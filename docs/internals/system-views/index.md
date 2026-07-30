<a id="views"></a>

# System Views

 In addition to the system catalogs, PostgreSQL provides a number of built-in views. Some system views provide convenient access to some commonly used queries on the system catalogs. Other views provide access to internal server state.

 The information schema ([The Information Schema](../../client-interfaces/the-information-schema/index.md#information-schema)) provides an alternative set of views which overlap the functionality of the system views. Since the information schema is SQL-standard whereas the views described here are PostgreSQL-specific, it's usually better to use the information schema if it provides all the information you need.

 [System Views](overview.md#view-table) lists the system views described here. More detailed documentation of each view follows below. There are some additional views that provide access to accumulated statistics; they are described in [Collected Statistics Views](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-views-table).

- [Overview](overview.md#views-overview)
- [`pg_aios`](pg_aios.md#view-pg-aios)
- [`pg_available_extensions`](pg_available_extensions.md#view-pg-available-extensions)
- [`pg_available_extension_versions`](pg_available_extension_versions.md#view-pg-available-extension-versions)
- [`pg_backend_memory_contexts`](pg_backend_memory_contexts.md#view-pg-backend-memory-contexts)
- [`pg_config`](pg_config.md#view-pg-config)
- [`pg_cursors`](pg_cursors.md#view-pg-cursors)
- [`pg_dsm_registry_allocations`](pg_dsm_registry_allocations.md#view-pg-dsm-registry-allocations)
- [`pg_file_settings`](pg_file_settings.md#view-pg-file-settings)
- [`pg_group`](pg_group.md#view-pg-group)
- [`pg_hba_file_rules`](pg_hba_file_rules.md#view-pg-hba-file-rules)
- [`pg_ident_file_mappings`](pg_ident_file_mappings.md#view-pg-ident-file-mappings)
- [`pg_indexes`](pg_indexes.md#view-pg-indexes)
- [`pg_locks`](pg_locks.md#view-pg-locks)
- [`pg_matviews`](pg_matviews.md#view-pg-matviews)
- [`pg_policies`](pg_policies.md#view-pg-policies)
- [`pg_prepared_statements`](pg_prepared_statements.md#view-pg-prepared-statements)
- [`pg_prepared_xacts`](pg_prepared_xacts.md#view-pg-prepared-xacts)
- [`pg_publication_sequences`](pg_publication_sequences.md#view-pg-publication-sequences)
- [`pg_publication_tables`](pg_publication_tables.md#view-pg-publication-tables)
- [`pg_replication_origin_status`](pg_replication_origin_status.md#view-pg-replication-origin-status)
- [`pg_replication_slots`](pg_replication_slots.md#view-pg-replication-slots)
- [`pg_roles`](pg_roles.md#view-pg-roles)
- [`pg_rules`](pg_rules.md#view-pg-rules)
- [`pg_seclabels`](pg_seclabels.md#view-pg-seclabels)
- [`pg_sequences`](pg_sequences.md#view-pg-sequences)
- [`pg_settings`](pg_settings.md#view-pg-settings)
- [`pg_shadow`](pg_shadow.md#view-pg-shadow)
- [`pg_shmem_allocations`](pg_shmem_allocations.md#view-pg-shmem-allocations)
- [`pg_shmem_allocations_numa`](pg_shmem_allocations_numa.md#view-pg-shmem-allocations-numa)
- [`pg_stats`](pg_stats.md#view-pg-stats)
- [`pg_stats_ext`](pg_stats_ext.md#view-pg-stats-ext)
- [`pg_stats_ext_exprs`](pg_stats_ext_exprs.md#view-pg-stats-ext-exprs)
- [`pg_tables`](pg_tables.md#view-pg-tables)
- [`pg_timezone_abbrevs`](pg_timezone_abbrevs.md#view-pg-timezone-abbrevs)
- [`pg_timezone_names`](pg_timezone_names.md#view-pg-timezone-names)
- [`pg_user`](pg_user.md#view-pg-user)
- [`pg_user_mappings`](pg_user_mappings.md#view-pg-user-mappings)
- [`pg_views`](pg_views.md#view-pg-views)
- [`pg_wait_events`](pg_wait_events.md#view-pg-wait-events)
