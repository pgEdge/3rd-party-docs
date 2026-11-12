# Version 9.8

Release date: 2025-09-04

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.7.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.5

# New features

[Issue #6396](https://github.com/pgadmin-org/pgadmin4/issues/6396) -  Add menu items for truncating foreign tables.<br>
[Issue #7928](https://github.com/pgadmin-org/pgadmin4/issues/7928) -  Added Debian Trixie as a supported platform for pgAdmin.<br>
[Issue #8891](https://github.com/pgadmin-org/pgadmin4/issues/8891) -  Allow user to configure security related gunicorn parameters.<br>
[Issue #9093](https://github.com/pgadmin-org/pgadmin4/issues/9093) -  Change the default pgAdmin theme to System.<br>

# Housekeeping

[Issue #7448](https://github.com/pgadmin-org/pgadmin4/issues/7448) -  Remove usage of BrowserFS as it is deprecated.<br>

# Bug fixes

[Issue #9090](https://github.com/pgadmin-org/pgadmin4/issues/9090) -  Pin Paramiko to version 3.5.1 to fix the DSSKey error introduced in the latest release.<br>
[Issue #9095](https://github.com/pgadmin-org/pgadmin4/issues/9095) -  Fixed an issue where pgAdmin config migration was failing while upgrading to v9.7.<br>
[Issue #9114](https://github.com/pgadmin-org/pgadmin4/issues/9114) -  Fixed Cross-Origin Opener Policy (COOP) vulnerability in the OAuth 2.0 authentication flow (CVE-2025-9636).<br>
[Issue #9116](https://github.com/pgadmin-org/pgadmin4/issues/9116) -  Fixed an issue where editor shortcuts fail when using Option key combinations on macOS, due to macOS treating Option+Key as a different key input.<br>
