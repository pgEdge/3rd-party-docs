# Version 7.8

Release date: 2023-10-19

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v7.7.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, and 16

**EDB Advanced Server**: 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.0

# New features

[Issue #640](https://github.com/pgadmin-org/pgadmin4/issues/640) -    Add support for foreign table's new functionality for PG 11 and above.<br>
[Issue #6229](https://github.com/pgadmin-org/pgadmin4/issues/6229) -  Allow setting custom username for shared servers, with default as username of server being shared.<br>
[Issue #6373](https://github.com/pgadmin-org/pgadmin4/issues/6373) -  Add 'GENERATED ALWAYS AS..' option while creating column constraints for Foreign Table.<br>
[Issue #6797](https://github.com/pgadmin-org/pgadmin4/issues/6797) -  GUI representation of the system's activity using the 'system_stats' extension.<br>
[Issue #6802](https://github.com/pgadmin-org/pgadmin4/issues/6802) -  Added 'load_balance_hosts' connection string parameter for PG 16 and above.<br>

# Housekeeping

[Issue #6782](https://github.com/pgadmin-org/pgadmin4/issues/6782) -  Use PG16 as the default PostgreSQL version.<br>

# Bug fixes

[Issue #4995](https://github.com/pgadmin-org/pgadmin4/issues/4995) -  Fixed an issue in ERD tool where the downloaded images have a few links cut.<br>
[Issue #5749](https://github.com/pgadmin-org/pgadmin4/issues/5749) -  Fixed an issue where user was not able to assign new/old columns as primary key once column with primary key is deleted.<br>
[Issue #6285](https://github.com/pgadmin-org/pgadmin4/issues/6285) -  Add support for setting prepare threshold in server connection.<br>
[Issue #6482](https://github.com/pgadmin-org/pgadmin4/issues/6482) -  Fixed an issue where the wrong message "Current database has been moved or renamed" is displayed when debugging any function.<br>
[Issue #6538](https://github.com/pgadmin-org/pgadmin4/issues/6538) -  Fixed an issue where Processes tab displays wrong server name in some scenario.<br>
[Issue #6579](https://github.com/pgadmin-org/pgadmin4/issues/6579) -  Fix an issue where global/native keyboard shortcuts are not working when any cell of data output grid has focus.<br>
[Issue #6666](https://github.com/pgadmin-org/pgadmin4/issues/6666) -  Fixed query history slowness issue by storing query only for those having certain threshold max length.<br>
[Issue #6674](https://github.com/pgadmin-org/pgadmin4/issues/6674) -  Fix an issue where foreign table column name becomes "none" if the user changes any column data type.<br>
[Issue #6718](https://github.com/pgadmin-org/pgadmin4/issues/6718) -  Pin the cryptography version to fix PyO3 modules initialisation error.<br>
[Issue #6790](https://github.com/pgadmin-org/pgadmin4/issues/6790) -  Ensure that the backup works properly for PG 16 on the latest docker image.<br>
[Issue #6799](https://github.com/pgadmin-org/pgadmin4/issues/6799) -  Fixed an issue where the user is unable to select objects on the backup dialog due to tree flickering.<br>
[Issue #6836](https://github.com/pgadmin-org/pgadmin4/issues/6836) -  Fixed an issue where non-super PostgreSQL users are not able to terminate their own connections from dashboard.<br>
[Issue #6851](https://github.com/pgadmin-org/pgadmin4/issues/6851) -  Fix an issue where scale in columns is not allowed to have value as 0 or below.<br>
[Issue #6858](https://github.com/pgadmin-org/pgadmin4/issues/6858) -  Fix an issue in graphical explain plan where query tool crashes when the plan has parallel workers details and sort node is clicked for details.<br>
[Issue #6865](https://github.com/pgadmin-org/pgadmin4/issues/6865) -  Fix an issue where user login is not working if username/email contains single quote in server mode.<br>
