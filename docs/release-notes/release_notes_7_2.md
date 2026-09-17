# Version 7.2

Release date: 2023-06-01

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v7.1.

# Supported Database Servers

**PostgreSQL**: 11, 12, 13, 14 and 15

**EDB Advanced Server**: 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.2

# New features

[Issue #3831](https://github.com/pgadmin-org/pgadmin4/issues/3831) -  Add Option to only show active connections on Dashboard.<br>
[Issue #4769](https://github.com/pgadmin-org/pgadmin4/issues/4769) -  Allow pgAdmin to retrive master password from external script/program.<br>
[Issue #5048](https://github.com/pgadmin-org/pgadmin4/issues/5048) -  Add an option to hide/show empty object collection nodes.<br>
[Issue #5123](https://github.com/pgadmin-org/pgadmin4/issues/5123) -  Added support to use standard OS secret store to save server/ssh tunnel passwords instead of master password in pgAdmin desktop mode.<br>
[Issue #5868](https://github.com/pgadmin-org/pgadmin4/issues/5868) -  Implement new PostgreSQL 15 features in publication dialog and SQL.<br>

# Housekeeping

# Bug fixes

[Issue #5789](https://github.com/pgadmin-org/pgadmin4/issues/5789) -  Fixed an issue where Foreign Key columns are shown in the wrong order in SQL and Properties.<br>
[Issue #5817](https://github.com/pgadmin-org/pgadmin4/issues/5817) -  Ensure that a new row should be added on top in the User Management dialog.<br>
[Issue #5926](https://github.com/pgadmin-org/pgadmin4/issues/5926) -  Fixed an issue where REVOKE ALL DDL in table SQL was added only for one role.<br>
[Issue #6003](https://github.com/pgadmin-org/pgadmin4/issues/6003) -  Indicate the user if all the server's children nodes are hidden from the preferences setting.<br>
[Issue #6026](https://github.com/pgadmin-org/pgadmin4/issues/6026) -  Tools menu should be toggled for "pause replay of wal" and "resume replay of wal".<br>
[Issue #6043](https://github.com/pgadmin-org/pgadmin4/issues/6043) -  Make the 'Connect to server' dialog a modal dialog.<br>
[Issue #6080](https://github.com/pgadmin-org/pgadmin4/issues/6080) -  pgAdmin icon not showing on taskbar on Windows 10.<br>
[Issue #6127](https://github.com/pgadmin-org/pgadmin4/issues/6127) -  Fixed an issue where properties were not visible for FTS Parsers, FTS Templates, MViews, and Rules in Catalog objects.<br>
[Issue #6147](https://github.com/pgadmin-org/pgadmin4/issues/6147) -  Heartbeat is getting logged, though no server is connected in pgAdmin.<br>
[Issue #6204](https://github.com/pgadmin-org/pgadmin4/issues/6204) -  Ensure that name can't be empty in edit mode for Primary Key and Index.<br>
[Issue #6221](https://github.com/pgadmin-org/pgadmin4/issues/6221) -  Fix circular reference error for the multirange data types in the query tool.<br>
[Issue #6247](https://github.com/pgadmin-org/pgadmin4/issues/6247) -  Fixed an issue where PSQL tool prompts for password if using password exec command.<br>
[Issue #6253](https://github.com/pgadmin-org/pgadmin4/issues/6253) -  Fix an issue in the register server when setting the role, an arbitrary SQL query can be fired.<br>
[Issue #6267](https://github.com/pgadmin-org/pgadmin4/issues/6267) -  Ensure the user is able to log in if the specified OAUTH2_USERNAME_CLAIM is present in the OAuth2 profile.<br>
[Issue #6278](https://github.com/pgadmin-org/pgadmin4/issues/6278) -  Use dependent instead of dependant in the message.<br>
[Issue #6279](https://github.com/pgadmin-org/pgadmin4/issues/6279) -  Fix incorrect number of foreign tables displayed. Show column comments in RE-SQL.<br>
[Issue #6280](https://github.com/pgadmin-org/pgadmin4/issues/6280) -  View SQL tab not quoting column comments.<br>
[Issue #6281](https://github.com/pgadmin-org/pgadmin4/issues/6281) -  VarChar Field Sizes are missing from Query's Grid header.<br>
[Issue #6331](https://github.com/pgadmin-org/pgadmin4/issues/6331) -  Separate multiple Blocking PIDs with delimiter on Dashboard.<br>
