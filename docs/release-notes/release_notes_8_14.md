# Version 8.14

Release date: 2024-12-12

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.13.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, 16 and 17

**EDB Advanced Server**: 12, 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.0

# New features

[Issue #3317](https://github.com/pgadmin-org/pgadmin4/issues/3317) -  Allow setting NULL ordering for columns in view/edit data filter dialog.<br>
[Issue #3751](https://github.com/pgadmin-org/pgadmin4/issues/3751) -  Show auto-complete column names in filtered rows dialog of table and filter options of view/edit data tool.<br>
[Issue #5786](https://github.com/pgadmin-org/pgadmin4/issues/5786) -  Allow the use of a pgpass file in the pgAdmin container via Docker secrets.<br>
[Issue #6592](https://github.com/pgadmin-org/pgadmin4/issues/6592) -  Fixed multiple issues and improved ERD auto-layout.<br>
[Issue #6794](https://github.com/pgadmin-org/pgadmin4/issues/6794) -  Add support for passing connection details as placeholders in the passexec command.<br>
[Issue #7330](https://github.com/pgadmin-org/pgadmin4/issues/7330) -  Add support to deploy pgAdmin in a container with readOnlyRootFilesystem to true.<br>
[Issue #8095](https://github.com/pgadmin-org/pgadmin4/issues/8095) -  Added support for a builtin locale provider in the Database dialog.<br>
[Issue #8192](https://github.com/pgadmin-org/pgadmin4/issues/8192) -  Added support for adding tags on a server node.<br>

# Housekeeping

# Bug fixes

[Issue #5099](https://github.com/pgadmin-org/pgadmin4/issues/5099) -  Fixed an issue where Ctrl/Cmd + A was not selecting all data in query tool data grid.<br>
[Issue #7384](https://github.com/pgadmin-org/pgadmin4/issues/7384) -  Fixed an issue where querying a foreign table gives the error 'ForeignTableCommand' object has no attribute 'auto_commit'.<br>
[Issue #7486](https://github.com/pgadmin-org/pgadmin4/issues/7486) -  Fixed an issue where indent with space was not aligning to next tab position.<br>
[Issue #7809](https://github.com/pgadmin-org/pgadmin4/issues/7809) -  Fixed an issue where pgAdmin crashes on accessing the dashboard state tab when database is super busy.<br>
[Issue #7865](https://github.com/pgadmin-org/pgadmin4/issues/7865) -  Fixed an issue related to the query tool update connection after the server disconnected from the object explorer.<br>
[Issue #7892](https://github.com/pgadmin-org/pgadmin4/issues/7892) -  Fixed an issue where a column name change in the edit dialog of the table node does not show the updated column name.<br>
[Issue #8010](https://github.com/pgadmin-org/pgadmin4/issues/8010) -  Fixed an issue where query tool should show results and messages only from the last executed query.<br>
[Issue #8028](https://github.com/pgadmin-org/pgadmin4/issues/8028) -  Fixed an issue where query tool throws syntax error if a newly added row is untouched and saved.<br>
[Issue #8065](https://github.com/pgadmin-org/pgadmin4/issues/8065) -  Ensure the crypt key is retrieved correctly on backend server restart.<br>
[Issue #8098](https://github.com/pgadmin-org/pgadmin4/issues/8098) -  Fixed an issue in schema diff where an error message popup was showing some garbage without any info.<br>
[Issue #8127](https://github.com/pgadmin-org/pgadmin4/issues/8127) -  Fixed an issue where query tool should not prompt for unsaved changes when there are no changes.<br>
[Issue #8134](https://github.com/pgadmin-org/pgadmin4/issues/8134) -  Add a user preference to enable/disable alternating row background colors in the data output of query tool.<br>
[Issue #8157](https://github.com/pgadmin-org/pgadmin4/issues/8157) -  Fixed an issue where doing Ctrl/Cmd+C on a selected text in a cell editor of data output in the query tool copied the complete text.<br>
[Issue #8158](https://github.com/pgadmin-org/pgadmin4/issues/8158) -  Fixed an issue where auto-width of wide columns in data output is incorrectly calculated.<br>
[Issue #8215](https://github.com/pgadmin-org/pgadmin4/issues/8215) -  Ensure correct custom SSL certificate passed in connection string.<br>
