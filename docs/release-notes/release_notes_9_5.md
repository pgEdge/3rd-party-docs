# Version 9.5

Release date: 2025-06-30

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.4.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.2

# New features

[Issue #1926](https://github.com/pgadmin-org/pgadmin4/issues/1926) -  Add a new permission to allow disabling "Change Password" feature for a pgAdmin role.<br>
[Issue #1947](https://github.com/pgadmin-org/pgadmin4/issues/1947) -  Added role-based restrictions for editing server connections.<br>
[Issue #2659](https://github.com/pgadmin-org/pgadmin4/issues/2659) -  Added support for customizing keyboard shortcuts in the Query Tool's Edit menu.<br>
[Issue #2864](https://github.com/pgadmin-org/pgadmin4/issues/2864) -  Add a search box to enable searching within the preferences tab.<br>
[Issue #3319](https://github.com/pgadmin-org/pgadmin4/issues/3319) -  Added support to preserve the workspace, query windows, and pgAdmin state during an abrupt shutdown or restart.<br>
[Issue #6743](https://github.com/pgadmin-org/pgadmin4/issues/6743) -  Open preferences in a new tab instead of a dialog for better user experience.<br>
[Issue #8665](https://github.com/pgadmin-org/pgadmin4/issues/8665) -  Supports JSON logging for gunicorn process within Docker.<br>

# Housekeeping

# Bug fixes

[Issue #6118](https://github.com/pgadmin-org/pgadmin4/issues/6118) -  Improved PL/pgSQL code folding and support nested blocks.<br>
[Issue #7173](https://github.com/pgadmin-org/pgadmin4/issues/7173) -  Add a flag to allow access to system Python packages on recent Linux distributions.<br>
[Issue #7466](https://github.com/pgadmin-org/pgadmin4/issues/7466) -  Fixed an issue where utilities such as pg_dump and pg_restore failed to log error messages when required dependency files were missing.<br>
[Issue #8032](https://github.com/pgadmin-org/pgadmin4/issues/8032) -  Fixed an issue where the Schema Diff Tool incorrectly reported differences due to variations in the order of the privileges.<br>
[Issue #8235](https://github.com/pgadmin-org/pgadmin4/issues/8235) -  Fixed an issue in SQL syntax highlighting where the same color was used for both variable names and datatypes.<br>
[Issue #8691](https://github.com/pgadmin-org/pgadmin4/issues/8691) -  Fixed an issue in the query tool where using multiple cursors to copy text resulted in only the first line being copied.<br>
[Issue #8803](https://github.com/pgadmin-org/pgadmin4/issues/8803) -  Ensure that Keyboard shortcuts for save and download actions should not called when their respective UI buttons are disabled.<br>
[Issue #8808](https://github.com/pgadmin-org/pgadmin4/issues/8808) -  Fixed an issue where data export using a query opened the wrong dialog type.<br>
[Issue #8809](https://github.com/pgadmin-org/pgadmin4/issues/8809) -  Fixed an issue where data export using a query failed when the query contained a newline character.<br>
[Issue #8830](https://github.com/pgadmin-org/pgadmin4/issues/8830) -  Fixed a UI layout issue that occurred after upgrading from pgAdmin v9.2 when all tabs had been closed prior to the upgrade.<br>
[Issue #8834](https://github.com/pgadmin-org/pgadmin4/issues/8834) -  Fixed an issue where the Columns node was not visible under Catalog Objects.<br>
