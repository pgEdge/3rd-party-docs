# Version 9.10

Release date: 2025-11-13

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.9.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16, 17 and 18

**EDB Advanced Server**: 13, 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.0

# New features

[Issue #4306](https://github.com/pgadmin-org/pgadmin4/issues/4306) -  Added the ability to search for tables and automatically bring them into view in the ERD tool.<br>
[Issue #6391](https://github.com/pgadmin-org/pgadmin4/issues/6391) -  Add support of DEPENDS/NO DEPENDS ON EXTENSION for PROCEDURE.<br>
[Issue #6698](https://github.com/pgadmin-org/pgadmin4/issues/6698) -  Add support for setting image download resolution in the ERD tool.<br>
[Issue #7885](https://github.com/pgadmin-org/pgadmin4/issues/7885) -  Add support for displaying detailed Citus query plans instead of 'Custom Scan' placeholder.<br>
[Issue #8912](https://github.com/pgadmin-org/pgadmin4/issues/8912) -  Add support for formatting .pgerd ERD project file.<br>

# Housekeeping

[Issue #8676](https://github.com/pgadmin-org/pgadmin4/issues/8676) -  Migrate pgAdmin UI to use React 19.<br>

# Bug fixes

[Issue #8504](https://github.com/pgadmin-org/pgadmin4/issues/8504) -  Fixed an issue where data output column resize is not sticking in Safari.<br>
[Issue #9117](https://github.com/pgadmin-org/pgadmin4/issues/9117) -  Fixed an issue where Schema Diff does not ignore Tablespace for indexes.<br>
[Issue #9132](https://github.com/pgadmin-org/pgadmin4/issues/9132) -  Fixed an issue where the 2FA window redirected to the login page after session expiration.<br>
[Issue #9233](https://github.com/pgadmin-org/pgadmin4/issues/9233) -  Fixed an issue where the Select All option on the columns tab of import/export data was not working in languages other than English.<br>
[Issue #9240](https://github.com/pgadmin-org/pgadmin4/issues/9240) -  Fixed an issue where the Debian build process failed with a "Sphinx module not found" error when using a Python virtual environment.<br>
[Issue #9281](https://github.com/pgadmin-org/pgadmin4/issues/9281) -  Fixed an issue where the last used storage directory was reset to blank, leading to access denied errors during backup or restore operations.<br>
[Issue #9304](https://github.com/pgadmin-org/pgadmin4/issues/9304) -  Fixed an issue that prevented assigning multiple users to an RLS policy.<br>
[Issue #9320](https://github.com/pgadmin-org/pgadmin4/issues/9320) -  Fixed remote code execution vulnerability when restoring PLAIN-format SQL dumps in server mode (CVE-2025-12762).<br>
[Issue #9323](https://github.com/pgadmin-org/pgadmin4/issues/9323) -  Fixed Command injection vulnerability allowing arbitrary command execution on Windows (CVE-2025-12763).<br>
[Issue #9324](https://github.com/pgadmin-org/pgadmin4/issues/9324) -  Fixed LDAP authentication flow vulnerable to TLS certificate verification bypass (CVE-2025-12765).<br>
[Issue #9325](https://github.com/pgadmin-org/pgadmin4/issues/9325) -  Fixed LDAP injection vulnerability in LDAP authentication flow (CVE-2025-12764).<br>
