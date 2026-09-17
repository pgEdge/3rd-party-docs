# Version 8.3

Release date: 2024-02-09

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.2.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, and 16

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.0

# New features

[Issue #4419](https://github.com/pgadmin-org/pgadmin4/issues/4419) -  Allow drag-n-drop columns collection tree node as comma separated columns.<br>
[Issue #6380](https://github.com/pgadmin-org/pgadmin4/issues/6380) -  Added support to rename columns in Views.<br>
[Issue #6392](https://github.com/pgadmin-org/pgadmin4/issues/6392) -  Added BYPASSRLS|NOBYPASSRLS option while creating a Role.<br>
[Issue #6450](https://github.com/pgadmin-org/pgadmin4/issues/6450) -  Added support for column storage syntax while creating table.<br>
[Issue #6557](https://github.com/pgadmin-org/pgadmin4/issues/6557) -  Use COOKIE_DEFAULT_PATH or SCRIPT_NAME in session cookie path.<br>
[Issue #6792](https://github.com/pgadmin-org/pgadmin4/issues/6792) -  Added configurable parameter to enable support for PasswordExecCommand in server mode.<br>

# Housekeeping

# Bug fixes

[Issue #5083](https://github.com/pgadmin-org/pgadmin4/issues/5083) -  Fixed an issue where format sql was messing up operator. Included many other feature changes, more details [here](https://github.com/pgadmin-org/pgadmin4/commit/f7045b58d4d1b98b6a2f035267d2dd01c7235aa6)<br>
[Issue #6785](https://github.com/pgadmin-org/pgadmin4/issues/6785) -  Fixed an issue where formatting inserts empty lines in specific case.<br>
[Issue #7053](https://github.com/pgadmin-org/pgadmin4/issues/7053) -  Add support for selecting a schema in the backup database dialog with no tables, mviews, views or foreign tables.<br>
[Issue #7055](https://github.com/pgadmin-org/pgadmin4/issues/7055) -  Fixed a UI border issue on the dependencies tab for columns with icon.<br>
[Issue #7073](https://github.com/pgadmin-org/pgadmin4/issues/7073) -  Fixed an issue where multiple errors were showing if user does not have connect privileges.<br>
[Issue #7085](https://github.com/pgadmin-org/pgadmin4/issues/7085) -  Fixed an issue where group membership information was displayed incorrectly.<br>
[Issue #7113](https://github.com/pgadmin-org/pgadmin4/issues/7113) -  Ensure that the correct SQL is generated when changing the column data type to "char".<br>
[Issue #7153](https://github.com/pgadmin-org/pgadmin4/issues/7153) -  Fixed an issue pgAdmin fails to launch due to inconsistent PATH variable name on windows.<br>
[Issue #7166](https://github.com/pgadmin-org/pgadmin4/issues/7166) -  Fix a backend server error when creating a named restore point.<br>
[Issue #7175](https://github.com/pgadmin-org/pgadmin4/issues/7175) -  Fix wrong default for FORCE_LOWER in the FLASK_GRAVATAR module.<br>
