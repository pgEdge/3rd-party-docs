# Version 4.17

Release date: 2020-01-09

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.16.

# New features

[Issue #4764](https://redmine.postgresql.org/issues/4764) -  Allow screen-reader to read relationship attributes in nested elements.<br>
[Issue #5060](https://redmine.postgresql.org/issues/5060) -  Ensure all binaries are securely signed and linked with the hardened runtime in the macOS bundle<br>

# Housekeeping

[Issue #4988](https://redmine.postgresql.org/issues/4988) -  Refactored SQL of Table's and it's child nodes.<br>
[Issue #5023](https://redmine.postgresql.org/issues/5023) -  Refactored SQL of Views and Materialized Views.<br>
[Issue #5024](https://redmine.postgresql.org/issues/5024) -  Refactored SQL of Functions and Procedures.<br>
[Issue #5038](https://redmine.postgresql.org/issues/5038) -  Added support for on-demand loading of items in Select2.<br>
[Issue #5048](https://redmine.postgresql.org/issues/5048) -  Added code coverage tool for pgAdmin.<br>

# Bug fixes

[Issue #4198](https://redmine.postgresql.org/issues/4198) -  Fix syntax highlighting in code mirror for backslash and escape constant.<br>
[Issue #4506](https://redmine.postgresql.org/issues/4506) -  Fix an issue where clicking on an empty textbox like fill factor or comments, considers it as change and enabled the save button.<br>
[Issue #4633](https://redmine.postgresql.org/issues/4633) -  Added support to view multilevel partitioned tables.<br>
[Issue #4842](https://redmine.postgresql.org/issues/4842) -  Ensure that constraints, indexes, rules, triggers, and compound triggers should be created on partitions.<br>
[Issue #4943](https://redmine.postgresql.org/issues/4943) -  Added more information to the 'Database connected/disconnected' message.<br>
[Issue #4950](https://redmine.postgresql.org/issues/4950) -  Ensure that the user should be able to select/modify tablespace for the partitioned table on v12 and above.<br>
[Issue #4999](https://redmine.postgresql.org/issues/4999) -  Rename some internal environment variables that could conflict with Kubernetes.<br>
[Issue #5004](https://redmine.postgresql.org/issues/5004) -  Fix vulnerability issues reported by 'yarn audit'. Replace the deprecated uglifyjs-webpack-plugin with a terser-webpack-plugin.<br>
[Issue #5008](https://redmine.postgresql.org/issues/5008) -  Ensure that the error message should not be displayed if Tablespace is not selected while creating the index.<br>
[Issue #5009](https://redmine.postgresql.org/issues/5009) -  Fix an issue where operator, access method and operator class is not visible for exclusion constraints.<br>
[Issue #5013](https://redmine.postgresql.org/issues/5013) -  Add a note to the documentation about the use of non-privileged ports on filesystems that don't support extended attributes when running the container.<br>
[Issue #5047](https://redmine.postgresql.org/issues/5047) -  Added tab navigation for tabs under explain panel in query tool.<br>
[Issue #5068](https://redmine.postgresql.org/issues/5068) -  Fix an issue where the table is not created with autovacuum_enabled and toast.autovacuum_enabled for PG/EPAS 12.<br>
