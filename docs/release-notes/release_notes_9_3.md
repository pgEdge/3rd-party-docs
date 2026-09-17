# Version 9.3

Release date: 2025-04-30

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.2.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.2

# New features

[Issue #2256](https://github.com/pgadmin-org/pgadmin4/issues/2256) -  Add support for type constructors for PostGIS spatial types.<br>
[Issue #2767](https://github.com/pgadmin-org/pgadmin4/issues/2767) -  Added ability to use SQL in the "DB Restriction" field.<br>
[Issue #7310](https://github.com/pgadmin-org/pgadmin4/issues/7310) -  Add support for custom roles and role permissions management in pgAdmin.<br>
[Issue #8629](https://github.com/pgadmin-org/pgadmin4/issues/8629) -  Added support for font ligatures.<br>

# Housekeeping

# Bug fixes

[Issue #3688](https://github.com/pgadmin-org/pgadmin4/issues/3688) -  Add Cache-Control no-cache,no-store,must-revalidate header to dynamically generated utils.js file.<br>
[Issue #5266](https://github.com/pgadmin-org/pgadmin4/issues/5266) -  Fixed an issue where shift + click on rows/columns for range selection did not work in the query tool data output window.<br>
[Issue #8443](https://github.com/pgadmin-org/pgadmin4/issues/8443) -  Fixed an issue where the debugger hangs when stepping into nested function/procedure.<br>
[Issue #8497](https://github.com/pgadmin-org/pgadmin4/issues/8497) -  Fixed an issue where the scroll position in the Object Explorer was not retained when switching workspaces.<br>
[Issue #8556](https://github.com/pgadmin-org/pgadmin4/issues/8556) -  Ensure that graph data is updated even when the Dashboard tab is inactive.<br>
[Issue #8572](https://github.com/pgadmin-org/pgadmin4/issues/8572) -  Fixed an issue where Ctrl/Cmd+A in cell editor would select all rows.<br>
[Issue #8613](https://github.com/pgadmin-org/pgadmin4/issues/8613) -  Fixed an issue where drag and drop function with no parameters does not work.<br>
[Issue #8627](https://github.com/pgadmin-org/pgadmin4/issues/8627) -  Fixed an issue where changes to foreign key constraints were not detected in the schema diff.<br>
[Issue #8628](https://github.com/pgadmin-org/pgadmin4/issues/8628) -  Change the shortcut for canceling a running query as it conflicts with the shortcut to open a new query tool.<br>
[Issue #8630](https://github.com/pgadmin-org/pgadmin4/issues/8630) -  Fixed an issue where filtering on a view caused an error.<br>
[Issue #8632](https://github.com/pgadmin-org/pgadmin4/issues/8632) -  Fixed an issue where the query tool went blank when converting history dates to the appropriate locale format.<br>
[Issue #8636](https://github.com/pgadmin-org/pgadmin4/issues/8636) -  Ensure that the server list is alphabetically sorted in the Query Tool/PSQL workspace.<br>
[Issue #8651](https://github.com/pgadmin-org/pgadmin4/issues/8651) -  Fixed an issue where the user management tab is not opening in the classic layout.<br>
[Issue #8652](https://github.com/pgadmin-org/pgadmin4/issues/8652) -  Allow OAuth 2 login using the username claim if the email is absent from the user profile.<br>
