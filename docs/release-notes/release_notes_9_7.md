# Version 9.7

Release date: 2025-08-21

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.6.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.2

# New features

[Issue #5766](https://github.com/pgadmin-org/pgadmin4/issues/5766) -  Add support for automatic updates in the pgAdmin 4 Desktop application on macOS.<br>
[Issue #6456](https://github.com/pgadmin-org/pgadmin4/issues/6456) -  Added GENERIC_PLAN, MEMORY, SERIALIZE option to EXPLAIN/EXPLAIN ANALYZE command.<br>
[Issue #8712](https://github.com/pgadmin-org/pgadmin4/issues/8712) -  Add support for changing cursor blink rate within the editors.<br>
[Issue #8917](https://github.com/pgadmin-org/pgadmin4/issues/8917) -  Add support for server tag-based filtering in the Object Explorer.<br>
[Issue #8931](https://github.com/pgadmin-org/pgadmin4/issues/8931) -  Added support for builtin locale provider while creating Collation.<br>

# Housekeeping

[Issue #6384](https://github.com/pgadmin-org/pgadmin4/issues/6384) -  Replace keyword PROCEDURE with FUNCTION while creating trigger and event trigger.<br>
[Issue #8861](https://github.com/pgadmin-org/pgadmin4/issues/8861) -  Introduced an ‘Editor’ preferences category and migrated all editor related settings into it.<br>

# Bug fixes

[Issue #7057](https://github.com/pgadmin-org/pgadmin4/issues/7057) -  Fixed an issue where custom column widths in the result grid of Query Tool or View/Edit Data were reset after re-executing a query.<br>
[Issue #7617](https://github.com/pgadmin-org/pgadmin4/issues/7617) -  Fixed the issue where updating the name of a table column does not reflect in the corresponding primary key constraint.<br>
[Issue #8149](https://github.com/pgadmin-org/pgadmin4/issues/8149) -  Fixed an issue where pgAdmin failed to update the server connection status when the server was disconnected in the background and a refresh was performed on that server.<br>
[Issue #8650](https://github.com/pgadmin-org/pgadmin4/issues/8650) -  Make Dashboard tables to be vertically resizable.<br>
[Issue #8756](https://github.com/pgadmin-org/pgadmin4/issues/8756) -  Fixed an issue in Firefox where the query window would shift to the left after opening the history tab or selecting a column header in the results grid.<br>
[Issue #8864](https://github.com/pgadmin-org/pgadmin4/issues/8864) -  Fixed an issue where CPU usage was very high on Windows when opening the psql tool.<br>
[Issue #8867](https://github.com/pgadmin-org/pgadmin4/issues/8867) -  Ensure DB restriction type is preserved while import and export server.<br>
[Issue #8969](https://github.com/pgadmin-org/pgadmin4/issues/8969) -  Fixed incorrect behaviour of the option deduplicate items after creating the index.<br>
[Issue #8971](https://github.com/pgadmin-org/pgadmin4/issues/8971) -  Added PKEY index in the index statistics summary.<br>
[Issue #9073](https://github.com/pgadmin-org/pgadmin4/issues/9073) -  Fixed an issue where adding breakpoints caused errors, and stepping out of a nested function removed breakpoints from the parent function.<br>
[Issue #9007](https://github.com/pgadmin-org/pgadmin4/issues/9007) -  Ensure the scratch pad in the Query Tool is not restored after it is closed.<br>
[Issue #9008](https://github.com/pgadmin-org/pgadmin4/issues/9008) -  Update the documentation for parameters that require file paths.<br>
[Issue #9047](https://github.com/pgadmin-org/pgadmin4/issues/9047) -  Fixed an issue where downloading images on the ERD tool was not working in desktop mode.<br>
[Issue #9067](https://github.com/pgadmin-org/pgadmin4/issues/9067) -  Ensure that disabling "Save Application State" in Preferences prevents tool data from being saved and stops it from being restored on application restart.<br>
