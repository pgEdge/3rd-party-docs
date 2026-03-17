# Version 3.5

Release date: 2018-11-01

This release contains a number of features and fixes reported since the release of pgAdmin4 3.4

# Features

[Issue #1253](https://redmine.postgresql.org/issues/1253) - Save the treeview state periodically, and restore it automatically when reconnecting.<br>
[Issue #3562](https://redmine.postgresql.org/issues/3562) - Migrate from Bootstrap 3 to Bootstrap 4.<br>

# Bug fixes

[Issue #3232](https://redmine.postgresql.org/issues/3232) - Ensure that Utilities(Backup/Restore/Maintenence/Import-Export) should not be started if binary path is wrong and also added 'Stop Process' button to cancel the process.<br>
[Issue #3638](https://redmine.postgresql.org/issues/3638) - Fix syntax error when creating new pgAgent schedules with a start date/time and exception.<br>
[Issue #3674](https://redmine.postgresql.org/issues/3674) - Cleanup session files periodically.<br>
[Issue #3660](https://redmine.postgresql.org/issues/3660) - Rename the 'SQL Editor' section of the Preferences to 'Query Tool' as it applies to the whole tool, not just the editor.<br>
[Issue #3676](https://redmine.postgresql.org/issues/3676) - Fix CREATE Script functionality for EDB-Wrapped functions.<br>
[Issue #3700](https://redmine.postgresql.org/issues/3700) - Fix connection garbage collector.<br>
[Issue #3703](https://redmine.postgresql.org/issues/3703) - Purge connections from the cache on logout.<br>
[Issue #3722](https://redmine.postgresql.org/issues/3722) - Ensure that utility existence check should work for schema and other child objects while taking Backup/Restore.<br>
[Issue #3730](https://redmine.postgresql.org/issues/3730) - Fixed fatal error while launching the pgAdmin4 3.5. Update the version of the Flask to 0.12.4 for release.<br>
