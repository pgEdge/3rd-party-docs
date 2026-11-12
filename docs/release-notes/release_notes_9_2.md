# Version 9.2

Release date: 2025-04-03

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.1.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.2

# New features

[Issue #0392](https://github.com/pgadmin-org/pgadmin4/issues/0392) -  Added an ability to configure the font family for SQL editors.<br>
[Issue #4194](https://github.com/pgadmin-org/pgadmin4/issues/4194) -  Added support to automatically open a file after it is downloaded in the desktop mode.<br>
[Issue #4503](https://github.com/pgadmin-org/pgadmin4/issues/4503) -  Added support for post-connection SQL execution, which will be run automatically on each connection made to any database of the server.<br>
[Issue #5871](https://github.com/pgadmin-org/pgadmin4/issues/5871) -  Add support for restoring plain SQL database dumps.<br>
[Issue #8034](https://github.com/pgadmin-org/pgadmin4/issues/8034) -  Added support for creating Directory nodes in EPAS.<br>
[Issue #8449](https://github.com/pgadmin-org/pgadmin4/issues/8449) -  Change icon buttons to show tooltip even when disabled.<br>
[Issue #8540](https://github.com/pgadmin-org/pgadmin4/issues/8540) -  Add an option to load/replace the servers.json file on each container startup.<br>
[Issue #8574](https://github.com/pgadmin-org/pgadmin4/issues/8574) -  Open user management in a separate tab instead of a dialog to enhance UI/UX.<br>

# Housekeeping

[Issue #8545](https://github.com/pgadmin-org/pgadmin4/issues/8545) -  Added Security.md file.<br>

# Bug fixes

[Issue #8316](https://github.com/pgadmin-org/pgadmin4/issues/8316) -  Ensure that modal dialogs are not triggered more than once to avoid duplicates.<br>
[Issue #8355](https://github.com/pgadmin-org/pgadmin4/issues/8355) -  Change session files garbage collection strategy.<br>
[Issue #8362](https://github.com/pgadmin-org/pgadmin4/issues/8362) -  Fixed an issue where pgAdmin should fallback to main screen if the last opened screen is disconnected.<br>
[Issue #8437](https://github.com/pgadmin-org/pgadmin4/issues/8437) -  Fixed an issue where the PSQL terminal displays keyname for non alphanumeric keys.<br>
[Issue #8462](https://github.com/pgadmin-org/pgadmin4/issues/8462) -  Fixed an issue where geometries in the geometry viewer will render partially when the container was resized.<br>
[Issue #8473](https://github.com/pgadmin-org/pgadmin4/issues/8473) -  Change the stop/terminate icon at all the places for better UX.<br>
[Issue #8479](https://github.com/pgadmin-org/pgadmin4/issues/8479) -  Fixed an issue where the Schema Diff was not displaying the difference query when a table had a UNIQUE NULLS NOT DISTINCT constraint.<br>
[Issue #8483](https://github.com/pgadmin-org/pgadmin4/issues/8483) -  Fixed an issue where the query tool data grid did not respect the default value for columns of domain type when the domain had a default value.<br>
[Issue #8514](https://github.com/pgadmin-org/pgadmin4/issues/8514) -  Ensure the newly added parameters in the server dialog are incorporated into the Import/Export Servers functionality.<br>
[Issue #8546](https://github.com/pgadmin-org/pgadmin4/issues/8546) -  Fixed an issue where updating the grantee was not correctly applying the privileges.<br>
[Issue #8577](https://github.com/pgadmin-org/pgadmin4/issues/8577) -  Fixed an issue where the upgrade_check API returned an unexpected keyword argument 'cafile' due to changes in the urllib package supporting Python v3.13.<br>
[Issue #8597](https://github.com/pgadmin-org/pgadmin4/issues/8597) -  Fixed an issue where delete/rename was done on wrong file after sorting in Storage Manager.<br>
[Issue #8602](https://github.com/pgadmin-org/pgadmin4/issues/8602) -  Fixed an XSS vulnerability issue in the Query Tool and View/Edit Data (CVE-2025-2946).<br>
[Issue #8603](https://github.com/pgadmin-org/pgadmin4/issues/8603) -  Fixed a remote code execution issue in the Query Tool and Cloud Deployment (CVE-2025-2945).<br>
[Issue #8623](https://github.com/pgadmin-org/pgadmin4/issues/8623) -  Fixed an issue where query tool is crashing on macOS 15.4 due to a locale issue.<br>
