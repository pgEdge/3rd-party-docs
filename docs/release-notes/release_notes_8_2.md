# Version 8.2

Release date: 2024-01-11

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.1.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, and 16

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.0

# New features

[Issue #2483](https://github.com/pgadmin-org/pgadmin4/issues/2483) -  Administer pgAdmin Users and Preferences Using the Command Line Interface (CLI).<br>
[Issue #5908](https://github.com/pgadmin-org/pgadmin4/issues/5908) -  Allow users to convert View/Edit table into a Query tool to enable editing the SQL generated.<br>
[Issue #6085](https://github.com/pgadmin-org/pgadmin4/issues/6085) -  Added copy server support, allowing the duplication of existing servers with the option to make certain modifications.<br>
[Issue #7016](https://github.com/pgadmin-org/pgadmin4/issues/7016) -  Added keep-alive support for SSH sessions when connecting to a PostgreSQL server via an SSH tunnel.<br>

# Housekeeping

[Issue #6926](https://github.com/pgadmin-org/pgadmin4/issues/6926) -  Ensure that eventlet's subprocess should be used following the resolution of an issue with Python 3.12 by eventlet.<br>

# Bug fixes

[Issue #6193](https://github.com/pgadmin-org/pgadmin4/issues/6193) -  Fixed an issue where query tool title did not change after "Save As" until any new change is made.<br>
[Issue #6781](https://github.com/pgadmin-org/pgadmin4/issues/6781) -  Fixed an issue where export servers was not adding extension if not specified.<br>
[Issue #6815](https://github.com/pgadmin-org/pgadmin4/issues/6815) -  Fixed an issue where pgAdmin imports servers to the wrong accounts for the external authentication.<br>
[Issue #7002](https://github.com/pgadmin-org/pgadmin4/issues/7002) -  Fixed an issue where an error occurred in the SQL tab when using an extended index(pgroonga).<br>
[Issue #7041](https://github.com/pgadmin-org/pgadmin4/issues/7041) -  Fixed an issue where changes done to a node using edit dialog are not reflecting on the properties tab if the properties tab is active.<br>
[Issue #7059](https://github.com/pgadmin-org/pgadmin4/issues/7059) -  Fixed an issue where DB Restrictions were not visible on the server dialog.<br>
[Issue #7061](https://github.com/pgadmin-org/pgadmin4/issues/7061) -  Ensure that the 'Dbo' schema is displayed as a regular schema rather than a system catalog schema.<br>
[Issue #7062](https://github.com/pgadmin-org/pgadmin4/issues/7062) -  Introduce LDAP configuration parameter LDAP_IGNORE_MALFORMED_SCHEMA to ignore fetching schema from the LDAP server.<br>
[Issue #7064](https://github.com/pgadmin-org/pgadmin4/issues/7064) -  Fixed an error-'amname' when generating ERD for database containing partition tables.<br>
[Issue #7066](https://github.com/pgadmin-org/pgadmin4/issues/7066) -  Fixed an issue where object explorer last tree state was not saving.<br>
[Issue #7070](https://github.com/pgadmin-org/pgadmin4/issues/7070) -  Fixed an issue where pgAgent job schedule dialog is not opening for edit.<br>
[Issue #7078](https://github.com/pgadmin-org/pgadmin4/issues/7078) -  Fixed an issue where user is not able to cancel or terminate active queries from dashboard.<br>
[Issue #7082](https://github.com/pgadmin-org/pgadmin4/issues/7082) -  Fixed browser autocomplete related issues on pgAdmin authentication related pages.<br>
[Issue #7091](https://github.com/pgadmin-org/pgadmin4/issues/7091) -  Fixed an issue where auto commit/rollback setting not persisting across query tool connection change.<br>
[Issue #7104](https://github.com/pgadmin-org/pgadmin4/issues/7104) -  Fixed an issue where Schema Diff not generating difference for missing columns.<br>
