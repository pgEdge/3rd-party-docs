# Version 6.21

Release date: 2023-03-09

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.20.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13, 14 and 15

**EDB Advanced Server**: 10, 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.1

# New features

[Issue #5832](https://github.com/pgadmin-org/pgadmin4/issues/5832) -  Allow changing cardinality notation in ERD to use Chen notation.<br>
[Issue #5842](https://github.com/pgadmin-org/pgadmin4/issues/5842) -  Add additional logging for successful logins and user creation.<br>

# Housekeeping

# Bug fixes

[Issue #5269](https://github.com/pgadmin-org/pgadmin4/issues/5269) -  Ensure that the schema diff tool should pick up the change in the column grants.<br>
[Issue #5685](https://github.com/pgadmin-org/pgadmin4/issues/5685) -  Ensure that Grant column permission to a view is visible in the SQL tab.<br>
[Issue #5747](https://github.com/pgadmin-org/pgadmin4/issues/5747) -  Ensure that content on the DDL comparison panel should get refreshed on selecting the object using the up and down arrow keys.<br>
[Issue #5756](https://github.com/pgadmin-org/pgadmin4/issues/5756) -  Fix for query tool prompting for unsaved changes even if no changes have been made.<br>
[Issue #5758](https://github.com/pgadmin-org/pgadmin4/issues/5758) -  Fixed an issue where lock layout menu was not in sync with preferences.<br>
[Issue #5760](https://github.com/pgadmin-org/pgadmin4/issues/5760) -  Fixed an issue where query was not pasted to editor after trojan source warning.<br>
[Issue #5764](https://github.com/pgadmin-org/pgadmin4/issues/5764) -  Fix an issue where the maintenance dialog for Materialized View gives an error.<br>
[Issue #5773](https://github.com/pgadmin-org/pgadmin4/issues/5773) -  Fixed an issue where Clear Saved Password should be disabled if the password is already cleared.<br>
[Issue #5790](https://github.com/pgadmin-org/pgadmin4/issues/5790) -  Fixed an issue where the user can't create trigger AFTER UPDATE OF.<br>
[Issue #5803](https://github.com/pgadmin-org/pgadmin4/issues/5803) -  Fix an issue where query tool is stripping spaces in grid cell.<br>
[Issue #5810](https://github.com/pgadmin-org/pgadmin4/issues/5810) -  Fix an issue where sequence owner is remove on sequence edit.<br>
[Issue #5822](https://github.com/pgadmin-org/pgadmin4/issues/5822) -  Do not allow to save invalid JSON in query tool JSON editor.<br>
[Issue #5847](https://github.com/pgadmin-org/pgadmin4/issues/5847) -  Fixed an issue where pgAdmin failed to connect when the Postgres password included special characters.<br>
[Issue #5870](https://github.com/pgadmin-org/pgadmin4/issues/5870) -  Ensure that the database migration does not fail with a NoSuchTableError exception.<br>
[Issue #5872](https://github.com/pgadmin-org/pgadmin4/issues/5872) -  Handle MERGE operation in query tool explain introduced in PostgreSQL 15.<br>
[Issue #5889](https://github.com/pgadmin-org/pgadmin4/issues/5889) -  Fixed an issue where the database server is not connected using a service file.<br>
[Issue #5909](https://github.com/pgadmin-org/pgadmin4/issues/5909) -  Fixed an issue where the file name in the query tool tab was not changing if another file is opened.<br>
