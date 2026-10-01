# Version 9.1

Release date: 2025-02-28

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.0.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.2

# New features

[Issue #1235](https://github.com/pgadmin-org/pgadmin4/issues/1235) -  Added an option to open the file in a new tab within the query tool.<br>
[Issue #5128](https://github.com/pgadmin-org/pgadmin4/issues/5128) -  Add support for one to one relationship in the ERD tool.<br>
[Issue #8436](https://github.com/pgadmin-org/pgadmin4/issues/8436) -  Enable the ability to close the 'Welcome' tab in the Query Tool/PSQL Workspace.<br>

# Housekeeping

[Issue #8068](https://github.com/pgadmin-org/pgadmin4/issues/8068) -  Added support for Python 3.13<br>

# Bug fixes

[Issue #8181](https://github.com/pgadmin-org/pgadmin4/issues/8181) -  Fixed an issue where pgAdmin does not support pg_vector column length/precision.<br>
[Issue #8296](https://github.com/pgadmin-org/pgadmin4/issues/8296) -  Fixed an issue where pasting text containing JSON data into the query tool grid would result in incorrect parsing.<br>
[Issue #8341](https://github.com/pgadmin-org/pgadmin4/issues/8341) -  Fixed an issue where the query tool was not treating IDENTITY columns as columns with default values when inserting new rows.<br>
[Issue #8389](https://github.com/pgadmin-org/pgadmin4/issues/8389) -  Fixed an issue where the ERD tool fails to open a saved file containing parent-child relationship within the same table.<br>
[Issue #8410](https://github.com/pgadmin-org/pgadmin4/issues/8410) -  Fixed Docker image entrypoint.sh email validation.<br>
[Issue #8418](https://github.com/pgadmin-org/pgadmin4/issues/8418) -  Fixed an issue where the User Management and Change Password dialogs were hidden when selecting a menu while a workspace other than 'Default' was active.<br>
[Issue #8430](https://github.com/pgadmin-org/pgadmin4/issues/8430) -  Fixed an issue where the column order displayed was incorrect for exclusion constraints with multiple columns.<br>
[Issue #8435](https://github.com/pgadmin-org/pgadmin4/issues/8435) -  Ensure the saved passwords are decrypted with the correct encryption key for external authentication in server mode.<br>
[Issue #8439](https://github.com/pgadmin-org/pgadmin4/issues/8439) -  Fixed an issue where drop-down menus were hidden behind the dock.<br>
[Issue #8460](https://github.com/pgadmin-org/pgadmin4/issues/8460) -  Fixed an issue where deleting rows in the query tool would delete all rows in the table when 'Select All Remaining Rows' was used.<br>
