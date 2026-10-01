# Version 7.7

Release date: 2023-09-21

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v7.6.

# Supported Database Servers

**PostgreSQL**: 11, 12, 13, 14 and 15

**EDB Advanced Server**: 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.4

# New features

[Issue #642](https://github.com/pgadmin-org/pgadmin4/issues/642)   -  Added support to select/deselect objects in the Backup dialog.<br>
[Issue #4805](https://github.com/pgadmin-org/pgadmin4/issues/4805) -  Added all the new options of the 'WITH' clause in the subscription dialog.<br>
[Issue #6378](https://github.com/pgadmin-org/pgadmin4/issues/6378) -  Added USING method while creating the table.<br>
[Issue #6379](https://github.com/pgadmin-org/pgadmin4/issues/6379) -  Added compression method option while creating a column.<br>
[Issue #6383](https://github.com/pgadmin-org/pgadmin4/issues/6383) -  Added Strategy, Locale Provider, ICU Locale, ICU Rules, and OID options while creating a database.<br>
[Issue #6400](https://github.com/pgadmin-org/pgadmin4/issues/6400) -  Added USING method while creating the materialized view.<br>
[Issue #6736](https://github.com/pgadmin-org/pgadmin4/issues/6736) -  Add support for additional ID token claim checks for OAuth 2 authentication.<br>

# Housekeeping

[Issue #2411](https://github.com/pgadmin-org/pgadmin4/issues/2411) -  Added the 'data type' column in the properties tab of the Columns collection node.<br>

# Bug fixes

[Issue #6274](https://github.com/pgadmin-org/pgadmin4/issues/6274) -  Fix an issue where user is not able to change the password when SMTP is not configured.<br>
[Issue #6704](https://github.com/pgadmin-org/pgadmin4/issues/6704) -  Ensure user is redirected to login page after failed login.<br>
[Issue #6712](https://github.com/pgadmin-org/pgadmin4/issues/6712) -  Ensure that Materialized view size fields in "Statistics" should be human-readable.<br>
[Issue #6730](https://github.com/pgadmin-org/pgadmin4/issues/6730) -  Fix an issue where changing the password shows success but the new password is not working.<br>
[Issue #6738](https://github.com/pgadmin-org/pgadmin4/issues/6738) -  Fix an issue where login form doesn't appear if internal auth source is removed.<br>
[Issue #6764](https://github.com/pgadmin-org/pgadmin4/issues/6764) -  Fix a security related issue where an authenticated user can run remote command using validate binary path API (CVE-2023-5002).<br>
