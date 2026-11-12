# Version 8.4

Release date: 2024-03-07

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.3.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, and 16

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.1

# New features

[Issue #6058](https://github.com/pgadmin-org/pgadmin4/issues/6058) -  Allow preferences customization using a configuration file.<br>
[Issue #7138](https://github.com/pgadmin-org/pgadmin4/issues/7138) -  Add support for JSON log format.<br>
[Issue #7204](https://github.com/pgadmin-org/pgadmin4/issues/7204) -  Add --yes option for skipping the confirmation prompt while deleting the user via CLI for scripting purpose.<br>

# Housekeeping

[Issue #7097](https://github.com/pgadmin-org/pgadmin4/issues/7097) -  Upgrade CodeMirror from version 5 to 6.<br>
[Issue #7148](https://github.com/pgadmin-org/pgadmin4/issues/7148) -  Added documentation for Dashboard's System Statistics tab.<br>
[Issue #7187](https://github.com/pgadmin-org/pgadmin4/issues/7187) -  Separate the application name, branding & version information from the configuration file.<br>
[Issue #7234](https://github.com/pgadmin-org/pgadmin4/issues/7234) -  Upgrade python packages cryptography to 42.0.x and Flask-Security-Too to 5.3.x.<br>

# Bug fixes

[Issue #6792](https://github.com/pgadmin-org/pgadmin4/issues/6792) -  Fix multiple issues where PasswordExecCommand was not working in server mode and PasswordExecCommand was not loaded when importing servers.<br>
[Issue #6808](https://github.com/pgadmin-org/pgadmin4/issues/6808) -  Fix the tabbed panel backward/forward shortcut for tabs.<br>
[Issue #7027](https://github.com/pgadmin-org/pgadmin4/issues/7027) -  Fixed an issue where dependencies and dependents were not showing if a composite type is used as an attribute in another composite type.<br>
[Issue #7164](https://github.com/pgadmin-org/pgadmin4/issues/7164) -  Fix an issue where constraint check control is enabled in the edit table dialog.<br>
[Issue #7165](https://github.com/pgadmin-org/pgadmin4/issues/7165) -  Fix an issue where the scripts created by generate script of Schema diff for Table with sequence was not working earlier.<br>
[Issue #7193](https://github.com/pgadmin-org/pgadmin4/issues/7193) -  Ensure that the OAuth2 session is logged out when users log out from pgAdmin.<br>
[Issue #7217](https://github.com/pgadmin-org/pgadmin4/issues/7217) -  Remove role related checks on the UI dashboard when terminating session/query and let PostgreSQL take care of it.<br>
[Issue #7225](https://github.com/pgadmin-org/pgadmin4/issues/7225) -  Fix an issue where type column in dependents/dependencies tab is not showing correct label.<br>
[Issue #7258](https://github.com/pgadmin-org/pgadmin4/issues/7258) -  Unsafe Deserialization and Remote Code Execution by an Authenticated user in pgAdmin 4 (CVE-2024-2044).<br>
