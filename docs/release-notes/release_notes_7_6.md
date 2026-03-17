# Version 7.6

Release date: 2023-08-24

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v7.5.

# Supported Database Servers

**PostgreSQL**: 11, 12, 13, 14 and 15

**EDB Advanced Server**: 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.3

# New features

[Issue #2595](https://github.com/pgadmin-org/pgadmin4/issues/2595) -  Added Expression to CREATE INDEX.<br>
[Issue #3942](https://github.com/pgadmin-org/pgadmin4/issues/3942) -  Added cascade option while creating an extension.<br>
[Issue #5759](https://github.com/pgadmin-org/pgadmin4/issues/5759) -  Added 'Ignore Grants' option in the schema diff tool.<br>
[Issue #6004](https://github.com/pgadmin-org/pgadmin4/issues/6004) -  Added 'Ignore Tablespace' option in the schema diff tool.<br>
[Issue #6375](https://github.com/pgadmin-org/pgadmin4/issues/6375) -  Added support for ALTER INDEX column statistics.<br>
[Issue #6376](https://github.com/pgadmin-org/pgadmin4/issues/6376) -  Added unlogged option while creating a sequence.<br>
[Issue #6377](https://github.com/pgadmin-org/pgadmin4/issues/6377) -  Added all like options while creating a table.<br>
[Issue #6381](https://github.com/pgadmin-org/pgadmin4/issues/6381) -  Added support for SYSTEM, CONCURRENTLY and TABLESPACE options in REINDEX.<br>
[Issue #6382](https://github.com/pgadmin-org/pgadmin4/issues/6382) -  Added WAL option to EXPLAIN ANALYZE command.<br>
[Issue #6397](https://github.com/pgadmin-org/pgadmin4/issues/6397) -  Added new/missing options to the VACUUM command.<br>
[Issue #6415](https://github.com/pgadmin-org/pgadmin4/issues/6415) -  Added SKIP_LOCKED and BUFFER_USAGE_LIMIT option to Analyze command.<br>
[Issue #6448](https://github.com/pgadmin-org/pgadmin4/issues/6448) -  Add support for TRUNCATE trigger in foreign table.<br>
[Issue #6595](https://github.com/pgadmin-org/pgadmin4/issues/6595) -  Ensure that Schema Diff comparison results should be displayed in the sorted order.<br>

# Housekeeping

[Issue #3702](https://github.com/pgadmin-org/pgadmin4/issues/3702) -  Generate software bill of materials as part of the package builds.<br>
[Issue #6588](https://github.com/pgadmin-org/pgadmin4/issues/6588) -  Added support for PostgreSQL and EPAS 16 to ensure it works without any errors.<br>

# Bug fixes

[Issue #5454](https://github.com/pgadmin-org/pgadmin4/issues/5454) -  Fix incorrect redirection URL after authentication by removing fixed value set to SCRIPT_NAME environment variable in pgAdmin4.wsgi file.<br>
[Issue #6208](https://github.com/pgadmin-org/pgadmin4/issues/6208) -  Allow changing the POOL_SIZE and MAX_OVERFLOW config values of the pgAdmin config DB connection pool.<br>
[Issue #6252](https://github.com/pgadmin-org/pgadmin4/issues/6252) -  Fix an issue where query tool on shared server is throwing error if the pgAdmin config DB is external.<br>
[Issue #6420](https://github.com/pgadmin-org/pgadmin4/issues/6420) -  Fix the query tool issue where raise Notice from func/proc or code blocks are no longer displayed live.<br>
[Issue #6500](https://github.com/pgadmin-org/pgadmin4/issues/6500) -  Fix the issue where query tool window turns blank if the user tries to generate a graph on the result.<br>
[Issue #6624](https://github.com/pgadmin-org/pgadmin4/issues/6624) -  Fix an issue where changing MFA_SUPPORTED_METHODS breaks the MFA validation.<br>
[Issue #6630](https://github.com/pgadmin-org/pgadmin4/issues/6630) -  Fix an issue where pgAdmin 7.5 fails to render table SQL with extension loaded index method.<br>
[Issue #6639](https://github.com/pgadmin-org/pgadmin4/issues/6639) -  Fix an issue where cycle syntax was not added in SQL when creating new sequence from UI.<br>
[Issue #6651](https://github.com/pgadmin-org/pgadmin4/issues/6651) -  Fix an issue where the SET directive is excluded from the function header in the schema diff tool.<br>
[Issue #6660](https://github.com/pgadmin-org/pgadmin4/issues/6660) -  Fix a query tool error 'pgAdminThread' object has no attribute 'native_id'.<br>
[Issue #6664](https://github.com/pgadmin-org/pgadmin4/issues/6664) -  Ensure keyboard shortcut for query execution is disabled when query execution is in progress.<br>
