# Version 8.10

Release date: 2024-07-29

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.9.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, 16 and 17

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.3

# New features

[Issue #3981](https://github.com/pgadmin-org/pgadmin4/issues/3981) -  Add support for Postgres Server Logs for Text, CSV and JSON format in plain and tabular formats. Upgrade React to version 18.<br>
[Issue #6572](https://github.com/pgadmin-org/pgadmin4/issues/6572) -  Add a keyboard shortcut to close active tab panel.<br>
[Issue #7530](https://github.com/pgadmin-org/pgadmin4/issues/7530) -  Add support for highlighting selection matches in the query editor.<br>

# Housekeeping

[Issue #7494](https://github.com/pgadmin-org/pgadmin4/issues/7494) -  Replace pgAdmin NW.js container with Electron container.<br>
[Issue #7501](https://github.com/pgadmin-org/pgadmin4/issues/7501) -  Updated to the latest version of the Notistack library.<br>
[Issue #7537](https://github.com/pgadmin-org/pgadmin4/issues/7537) -  Ensure that pgAdmin 4 is compatible with PostgreSQL v17.<br>
[Issue #7607](https://github.com/pgadmin-org/pgadmin4/issues/7607) -  Automatically apply virtualization in the DataGridView of SchemaView if the schema contains only one collection.<br>
[Issue #7623](https://github.com/pgadmin-org/pgadmin4/issues/7623) -  Add the git commit hash details to the About dialog.<br>

# Bug fixes

[Issue #3199](https://github.com/pgadmin-org/pgadmin4/issues/3199) -  Fixed an issue where paste operation in query tool data grid should skip bytea columns and put the value as NULL instead.<br>
[Issue #4165](https://github.com/pgadmin-org/pgadmin4/issues/4165) -  Fixed an issue where the taskbar icon appeared as a red square for the query tool and schema diff when opened in a new window.<br>
[Issue #5345](https://github.com/pgadmin-org/pgadmin4/issues/5345) -  Fix issue with missing new added records in download file.<br>
[Issue #5610](https://github.com/pgadmin-org/pgadmin4/issues/5610) -  Fixed an issue where the File Open dialog did not show files without a dot extension.<br>
[Issue #6548](https://github.com/pgadmin-org/pgadmin4/issues/6548) -  Ensure pgAdmin never makes network requests to Google etc.<br>
[Issue #6571](https://github.com/pgadmin-org/pgadmin4/issues/6571) -  Fixed an issue where pop-up notifications from Object Explorer wouldn't get dismissed automatically if the Query Tool was opened.<br>
[Issue #7035](https://github.com/pgadmin-org/pgadmin4/issues/7035) -  Fixed the permission denied issue for functions of the pgstattuple extension when accessing statistics with a non-admin user.<br>
[Issue #7219](https://github.com/pgadmin-org/pgadmin4/issues/7219) -  Ensure processes related notifiers disappears.<br>
[Issue #7297](https://github.com/pgadmin-org/pgadmin4/issues/7297) -  Updated entrypoint.sh to utilize the email-validator package for email validation.<br>
[Issue #7511](https://github.com/pgadmin-org/pgadmin4/issues/7511) -  Fixed an issue where users could not insert characters at the desired location, as it was added to the end of the line.<br>
[Issue #7554](https://github.com/pgadmin-org/pgadmin4/issues/7554) -  Fixed an issue where sorting the database activity table on the dashboard by any column caused the details to expand in the wrong position.<br>
[Issue #7618](https://github.com/pgadmin-org/pgadmin4/issues/7618) -  Fix an issue where the preferences JSON file has no effect when an external database is used.<br>
[Issue #7626](https://github.com/pgadmin-org/pgadmin4/issues/7626) -  Fixed an issue where theme preview under theme options was broken in pgAdmin server mode.<br>
[Issue #7627](https://github.com/pgadmin-org/pgadmin4/issues/7627) -  Fixed an issue where users could not autofill their saved passwords in the connect server dialog in the browser.<br>
[Issue #7638](https://github.com/pgadmin-org/pgadmin4/issues/7638) -  Fixed an issue where Generate Script button should be disabled if no objects are selected in the schema diff result.<br>
[Issue #7639](https://github.com/pgadmin-org/pgadmin4/issues/7639) -  Fixed an issue where ERD Open/Save shorcuts were not working on Windows/Linux.<br>
[Issue #7660](https://github.com/pgadmin-org/pgadmin4/issues/7660) -  Add a precautionary check for the query tool connection cursor to fix the error 'NoneType' object has no attribute '_query'.<br>
[Issue #7662](https://github.com/pgadmin-org/pgadmin4/issues/7662) -  Fixed an issue where boolean values in node details of graphical explain plan were not interpreted correctly.<br>
[Issue #7663](https://github.com/pgadmin-org/pgadmin4/issues/7663) -  Fixed an issue where Reassign/Drop Owned dialog not opening for Role.<br>
[Issue #7679](https://github.com/pgadmin-org/pgadmin4/issues/7679) -  Ensure pgadmin does not try to connect to the server if saved password is not available.<br>
[Issue #7681](https://github.com/pgadmin-org/pgadmin4/issues/7681) -  Ensure that pgAdmin works when opened in an iframe.<br>
