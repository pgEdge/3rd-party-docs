# Version 9.0

Release date: 2025-02-06

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.14.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.0

# New features

[Issue #6513](https://github.com/pgadmin-org/pgadmin4/issues/6513) -  Change button labels and color in delete confirmation dialog for all objects to improve UX.<br>
[Issue #7708](https://github.com/pgadmin-org/pgadmin4/issues/7708) -  Enhanced pgAdmin 4 with support for Workspace layouts.<br>
[Issue #8332](https://github.com/pgadmin-org/pgadmin4/issues/8332) -  Added the MAINTAIN privilege for PostgreSQL version 17 and above.<br>
[Issue #8391](https://github.com/pgadmin-org/pgadmin4/issues/8391) -  Add support for OAuth2 profile array response, which also takes care of the GitHub Private Email ID issue.<br>

# Housekeeping

[Issue #8249](https://github.com/pgadmin-org/pgadmin4/issues/8249) -  Show the python version used for the pgAdmin server in the about dialog.<br>

# Bug fixes

[Issue #3273](https://github.com/pgadmin-org/pgadmin4/issues/3273) -  Change the logic of setval function, so that the next nextval of sequence will return exactly the specified value.<br>
[Issue #5204](https://github.com/pgadmin-org/pgadmin4/issues/5204) -  Fixed an issue where pgadmin cannot install into path with non ASCII characters.<br>
[Issue #6044](https://github.com/pgadmin-org/pgadmin4/issues/6044) -  Fixed an issue where filter dialog save fails when the PostgreSQL server/database connection is lost.<br>
[Issue #6968](https://github.com/pgadmin-org/pgadmin4/issues/6968) -  Fixed an issue where option key was not registering in PSQL tool.<br>
[Issue #8072](https://github.com/pgadmin-org/pgadmin4/issues/8072) -  Fixed an issue where Schema Diff not produce difference script for Index definition with where condition.<br>
[Issue #8142](https://github.com/pgadmin-org/pgadmin4/issues/8142) -  Correct the documentation for the MFA configuration.<br>
[Issue #8165](https://github.com/pgadmin-org/pgadmin4/issues/8165) -  Fixed an issue where error message from the database server need space between two sentences.<br>
[Issue #8199](https://github.com/pgadmin-org/pgadmin4/issues/8199) -  Fixed an issue where query tool throws utf-8 decode error when using cursor with binary data.<br>
[Issue #8208](https://github.com/pgadmin-org/pgadmin4/issues/8208) -  Allow deleting the entry while creating/adding new label to enumeration type.<br>
[Issue #8209](https://github.com/pgadmin-org/pgadmin4/issues/8209) -  Fixed an issue where properties dialog throwing an error for Materialized View.<br>
[Issue #8254](https://github.com/pgadmin-org/pgadmin4/issues/8254) -  Fix a formatting issue in View/Edit tool generated SQL where some filters are applied.<br>
[Issue #8255](https://github.com/pgadmin-org/pgadmin4/issues/8255) -  Fixed an issue where tooltip on a dropdown button is blocking access to dropdown menu.<br>
[Issue #8256](https://github.com/pgadmin-org/pgadmin4/issues/8256) -  Fix the error occurring while loading preferences on startup.<br>
[Issue #8273](https://github.com/pgadmin-org/pgadmin4/issues/8273) -  Fixed an issue where copying query tool output cell is not working if any SQL text is selected.<br>
[Issue #8299](https://github.com/pgadmin-org/pgadmin4/issues/8299) -  Ensure master password pop up is not shown on setting MASTER_PASSWORD_REQUIRED to false.<br>
[Issue #8309](https://github.com/pgadmin-org/pgadmin4/issues/8309) -  Remove the option "With no data (concurrently)" from Refresh MATERIALIZED VIEW context menu.<br>
[Issue #8320](https://github.com/pgadmin-org/pgadmin4/issues/8320) -  Fix an issue where wrong information is shown after using the filter on the Dashboard> State tab.<br>
[Issue #8365](https://github.com/pgadmin-org/pgadmin4/issues/8365) -  Fixed an issue where PSQL tool is not opening if database name have HTML characters in the name.<br>
[Issue #8369](https://github.com/pgadmin-org/pgadmin4/issues/8369) -  Fixed an issue where Default Privileges and Privileges not working correctly.<br>
[Issue #8408](https://github.com/pgadmin-org/pgadmin4/issues/8408) -  Fixed an issue where quotes were missing in the CREATE script for the tablespace.<br>
