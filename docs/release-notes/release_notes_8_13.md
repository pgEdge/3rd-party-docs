# Version 8.13

Release date: 2024-11-14

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.12.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, 16 and 17

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.0

# New features

[Issue #1780](https://github.com/pgadmin-org/pgadmin4/issues/1780) -  Replace infinite scrolling with pagination in query tool data output for better UX and performance.<br>
[Issue #1984](https://github.com/pgadmin-org/pgadmin4/issues/1984) -  Add an object menu option to disconnect all server and database connections.<br>
[Issue #2874](https://github.com/pgadmin-org/pgadmin4/issues/2874) -  Alternate the background color of rows to enhance readability.<br>
[Issue #5869](https://github.com/pgadmin-org/pgadmin4/issues/5869) -  Allow to pass PGADMIN_CONFIG_CONFIG_DATABASE_URI from docker secrets.<br>
[Issue #6016](https://github.com/pgadmin-org/pgadmin4/issues/6016) -  Format dashboard graph metrics for better readability.<br>
[Issue #7393](https://github.com/pgadmin-org/pgadmin4/issues/7393) -  Added support for auto-detecting and setting the End-of-line character (LF/CRLF) in the query tool editor.<br>
[Issue #7863](https://github.com/pgadmin-org/pgadmin4/issues/7863) -  Make the background color for the serial number and header cells distinct.<br>

# Housekeeping

[Issue #8061](https://github.com/pgadmin-org/pgadmin4/issues/8061) -  Updated the react-menu library from v2 to v4.<br>
[Issue #8077](https://github.com/pgadmin-org/pgadmin4/issues/8077) -  Remove bootstrap from the runtime code.<br>
[Issue #8102](https://github.com/pgadmin-org/pgadmin4/issues/8077) -  Remove the dependency on python packages eventlet and greenlet.<br>

# Bug fixes

[Issue #5307](https://github.com/pgadmin-org/pgadmin4/issues/5307) -  Fixed an issue to allow the Enter key to save data in the result grid text editor.<br>
[Issue #7289](https://github.com/pgadmin-org/pgadmin4/issues/7289) -  Move 'About pgAdmin 4' to app menu on macOS.<br>
[Issue #7655](https://github.com/pgadmin-org/pgadmin4/issues/7655) -  Fixed an issue where the query tool was crashing when an empty geometry was being rendered.<br>
[Issue #7837](https://github.com/pgadmin-org/pgadmin4/issues/7837) -  Fixed an issue where role properties were not loading.<br>
[Issue #7883](https://github.com/pgadmin-org/pgadmin4/issues/7883) -  Fix multiple issues related to debugger params dialog input.<br>
[Issue #7919](https://github.com/pgadmin-org/pgadmin4/issues/7919) -  Fixed an issue where the dock layout was not saved upon closing a tab.<br>
[Issue #7920](https://github.com/pgadmin-org/pgadmin4/issues/7920) -  Fixed an issue where the copy shortcut CTRL +C was not working in the Query Tool data grid.<br>
[Issue #7955](https://github.com/pgadmin-org/pgadmin4/issues/7955) -  Fixed an issue where Dashboard tab showing 'Something went wrong'.<br>
[Issue #7957](https://github.com/pgadmin-org/pgadmin4/issues/7957) -  Fixed an issue where cursor selection is not visible in the PSQL tool.<br>
[Issue #7965](https://github.com/pgadmin-org/pgadmin4/issues/7965) -  Allow OAuth2 params OAUTH2_NAME, OAUTH2_DISPLAY_NAME, OAUTH2_ICON, OAUTH2_BUTTON_COLOR to be optional.<br>
[Issue #7988](https://github.com/pgadmin-org/pgadmin4/issues/7988) -  Add appropriate minimum width and height setting for desktop apps to avoid app accidentally going below visibility levels.<br>
[Issue #7993](https://github.com/pgadmin-org/pgadmin4/issues/7993) -  Fixed an issue where graphical explain keeps the node details open even after plan changed.<br>
[Issue #8007](https://github.com/pgadmin-org/pgadmin4/issues/8007) -  Fixed an issue in the debugger where function arguments of the character data type were being truncated.<br>
[Issue #8013](https://github.com/pgadmin-org/pgadmin4/issues/8013) -  Fix the issue where rows do not retain shading when scrolling.<br>
[Issue #8022](https://github.com/pgadmin-org/pgadmin4/issues/8022) -  Use PG17 as the default in the container.<br>
