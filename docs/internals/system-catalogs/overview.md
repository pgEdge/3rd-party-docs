<a id="catalogs-overview"></a>

## Overview


 [System Catalogs](#catalog-table) lists the system catalogs. More detailed documentation of each catalog follows below.


 Most system catalogs are copied from the template database during database creation and are thereafter database-specific. A few catalogs are physically shared across all databases in a cluster; these are noted in the descriptions of the individual catalogs.
 <a id="catalog-table"></a>

**Table: System Catalogs**

| Catalog Name | Purpose |
| --- | --- |
| [`pg_aggregate`](pg_aggregate.md#catalog-pg-aggregate) | aggregate functions |
| [`pg_am`](pg_am.md#catalog-pg-am) | relation access methods |
| [`pg_amop`](pg_amop.md#catalog-pg-amop) | access method operators |
| [`pg_amproc`](pg_amproc.md#catalog-pg-amproc) | access method support functions |
| [`pg_attrdef`](pg_attrdef.md#catalog-pg-attrdef) | column default values |
| [`pg_attribute`](pg_attribute.md#catalog-pg-attribute) | table columns (“attributes”) |
| [`pg_authid`](pg_authid.md#catalog-pg-authid) | authorization identifiers (roles) |
| [`pg_auth_members`](pg_auth_members.md#catalog-pg-auth-members) | authorization identifier membership relationships |
| [`pg_cast`](pg_cast.md#catalog-pg-cast) | casts (data type conversions) |
| [`pg_class`](pg_class.md#catalog-pg-class) | tables, indexes, sequences, views (“relations”) |
| [`pg_collation`](pg_collation.md#catalog-pg-collation) | collations (locale information) |
| [`pg_constraint`](pg_constraint.md#catalog-pg-constraint) | check constraints, unique constraints, primary key constraints, foreign key constraints |
| [`pg_conversion`](pg_conversion.md#catalog-pg-conversion) | encoding conversion information |
| [`pg_database`](pg_database.md#catalog-pg-database) | databases within this database cluster |
| [`pg_db_role_setting`](pg_db_role_setting.md#catalog-pg-db-role-setting) | per-role and per-database settings |
| [`pg_default_acl`](pg_default_acl.md#catalog-pg-default-acl) | default privileges for object types |
| [`pg_depend`](pg_depend.md#catalog-pg-depend) | dependencies between database objects |
| [`pg_description`](pg_description.md#catalog-pg-description) | descriptions or comments on database objects |
| [`pg_enum`](pg_enum.md#catalog-pg-enum) | enum label and value definitions |
| [`pg_event_trigger`](pg_event_trigger.md#catalog-pg-event-trigger) | event triggers |
| [`pg_extension`](pg_extension.md#catalog-pg-extension) | installed extensions |
| [`pg_foreign_data_wrapper`](pg_foreign_data_wrapper.md#catalog-pg-foreign-data-wrapper) | foreign-data wrapper definitions |
| [`pg_foreign_server`](pg_foreign_server.md#catalog-pg-foreign-server) | foreign server definitions |
| [`pg_foreign_table`](pg_foreign_table.md#catalog-pg-foreign-table) | additional foreign table information |
| [`pg_index`](pg_index.md#catalog-pg-index) | additional index information |
| [`pg_inherits`](pg_inherits.md#catalog-pg-inherits) | table inheritance hierarchy |
| [`pg_init_privs`](pg_init_privs.md#catalog-pg-init-privs) | object initial privileges |
| [`pg_language`](pg_language.md#catalog-pg-language) | languages for writing functions |
| [`pg_largeobject`](pg_largeobject.md#catalog-pg-largeobject) | data pages for large objects |
| [`pg_largeobject_metadata`](pg_largeobject_metadata.md#catalog-pg-largeobject-metadata) | metadata for large objects |
| [`pg_namespace`](pg_namespace.md#catalog-pg-namespace) | schemas |
| [`pg_opclass`](pg_opclass.md#catalog-pg-opclass) | access method operator classes |
| [`pg_operator`](pg_operator.md#catalog-pg-operator) | operators |
| [`pg_opfamily`](pg_opfamily.md#catalog-pg-opfamily) | access method operator families |
| [`pg_parameter_acl`](pg_parameter_acl.md#catalog-pg-parameter-acl) | configuration parameters for which privileges have been granted |
| [`pg_partitioned_table`](pg_partitioned_table.md#catalog-pg-partitioned-table) | information about partition key of tables |
| [`pg_policy`](pg_policy.md#catalog-pg-policy) | row-security policies |
| [`pg_proc`](pg_proc.md#catalog-pg-proc) | functions and procedures |
| [`pg_publication`](pg_publication.md#catalog-pg-publication) | publications for logical replication |
| [`pg_publication_namespace`](pg_publication_namespace.md#catalog-pg-publication-namespace) | schema to publication mapping |
| [`pg_publication_rel`](pg_publication_rel.md#catalog-pg-publication-rel) | relation to publication mapping |
| [`pg_range`](pg_range.md#catalog-pg-range) | information about range types |
| [`pg_replication_origin`](pg_replication_origin.md#catalog-pg-replication-origin) | registered replication origins |
| [`pg_rewrite`](pg_rewrite.md#catalog-pg-rewrite) | query rewrite rules |
| [`pg_seclabel`](pg_seclabel.md#catalog-pg-seclabel) | security labels on database objects |
| [`pg_sequence`](pg_sequence.md#catalog-pg-sequence) | information about sequences |
| [`pg_shdepend`](pg_shdepend.md#catalog-pg-shdepend) | dependencies on shared objects |
| [`pg_shdescription`](pg_shdescription.md#catalog-pg-shdescription) | comments on shared objects |
| [`pg_shseclabel`](pg_shseclabel.md#catalog-pg-shseclabel) | security labels on shared database objects |
| [`pg_statistic`](pg_statistic.md#catalog-pg-statistic) | planner statistics |
| [`pg_statistic_ext`](pg_statistic_ext.md#catalog-pg-statistic-ext) | extended planner statistics (definition) |
| [`pg_statistic_ext_data`](pg_statistic_ext_data.md#catalog-pg-statistic-ext-data) | extended planner statistics (built statistics) |
| [`pg_subscription`](pg_subscription.md#catalog-pg-subscription) | logical replication subscriptions |
| [`pg_subscription_rel`](pg_subscription_rel.md#catalog-pg-subscription-rel) | relation state for subscriptions |
| [`pg_tablespace`](pg_tablespace.md#catalog-pg-tablespace) | tablespaces within this database cluster |
| [`pg_transform`](pg_transform.md#catalog-pg-transform) | transforms (data type to procedural language conversions) |
| [`pg_trigger`](pg_trigger.md#catalog-pg-trigger) | triggers |
| [`pg_ts_config`](pg_ts_config.md#catalog-pg-ts-config) | text search configurations |
| [`pg_ts_config_map`](pg_ts_config_map.md#catalog-pg-ts-config-map) | text search configurations' token mappings |
| [`pg_ts_dict`](pg_ts_dict.md#catalog-pg-ts-dict) | text search dictionaries |
| [`pg_ts_parser`](pg_ts_parser.md#catalog-pg-ts-parser) | text search parsers |
| [`pg_ts_template`](pg_ts_template.md#catalog-pg-ts-template) | text search templates |
| [`pg_type`](pg_type.md#catalog-pg-type) | data types |
| [`pg_user_mapping`](pg_user_mapping.md#catalog-pg-user-mapping) | mappings of users to foreign servers |
