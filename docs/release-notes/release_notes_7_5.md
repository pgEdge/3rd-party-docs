# Version 7.5

Release date: 2023-07-27

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v7.4.

# Supported Database Servers

**PostgreSQL**: 11, 12, 13, 14 and 15

**EDB Advanced Server**: 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.3

# New features

[Issue #6369](https://github.com/pgadmin-org/pgadmin4/issues/6369) -  Added support to detach partitions using concurrently and finalize.<br>
[Issue #6374](https://github.com/pgadmin-org/pgadmin4/issues/6374) -  Added all supported index storage parameters while creating an index.<br>
[Issue #6416](https://github.com/pgadmin-org/pgadmin4/issues/6416) -  Added new/missing parameters to pg_dumpall (Backup Server).<br>
[Issue #6417](https://github.com/pgadmin-org/pgadmin4/issues/6417) -  Added new/missing parameters to pg_dump (Backup Objects).<br>
[Issue #6562](https://github.com/pgadmin-org/pgadmin4/issues/6562) -  Added new/missing parameters to pg_restore.<br>

# Housekeeping

[Issue #6295](https://github.com/pgadmin-org/pgadmin4/issues/6295) -  Remove Bootstrap and jQuery from authentication pages and rewrite them in ReactJS.<br>
[Issue #6323](https://github.com/pgadmin-org/pgadmin4/issues/6323) -  Enable cluster deployment with gp3 volume for AWS & BigAnimal cloud providers.<br>
[Issue #6423](https://github.com/pgadmin-org/pgadmin4/issues/6423) -  Clarify the LICENSE file to indicate that it is the PostgreSQL Licence.<br>
[Issue #6532](https://github.com/pgadmin-org/pgadmin4/issues/6532) -  Remove unsupported PostgreSQL versions from the container.<br>

# Bug fixes

[Issue #6163](https://github.com/pgadmin-org/pgadmin4/issues/6163) -  Fix an issue where queries can't complete execution.<br>
[Issue #6165](https://github.com/pgadmin-org/pgadmin4/issues/6165) -  Fixed an issue where Import Export not working when using pgpassfile.<br>
[Issue #6317](https://github.com/pgadmin-org/pgadmin4/issues/6317) -  Fix an issue where queries longer than 1 minute get stuck - Container 7.1<br>
[Issue #6356](https://github.com/pgadmin-org/pgadmin4/issues/6356) -  Fix an issue where queries get stuck with auto-completion enabled.<br>
[Issue #6364](https://github.com/pgadmin-org/pgadmin4/issues/6364) -  Fixed Query Tool/ PSQL tool tab title not getting updated on database rename.<br>
[Issue #6406](https://github.com/pgadmin-org/pgadmin4/issues/6406) -  Ensure user gets proper error if incorrect credentials are entered while authenticating AWS.<br>
[Issue #6489](https://github.com/pgadmin-org/pgadmin4/issues/6489) -  Fix an issue where the edit server fails in desktop mode if the server password is not stored.<br>
[Issue #6499](https://github.com/pgadmin-org/pgadmin4/issues/6499) -  Ensure that Backup, Restore, and Maintenance should work properly when pgpass file is used.<br>
[Issue #6501](https://github.com/pgadmin-org/pgadmin4/issues/6501) -  Fix the query tool auto-complete issue on the server reconnection.<br>
[Issue #6502](https://github.com/pgadmin-org/pgadmin4/issues/6502) -  Fix the query tool restore connection issue.<br>
[Issue #6509](https://github.com/pgadmin-org/pgadmin4/issues/6509) -  Fix the reconnecton issue if the PostgreSQL server is restarted from the backend.<br>
[Issue #6514](https://github.com/pgadmin-org/pgadmin4/issues/6514) -  Fix the connection and stability issues since v7, possibly related to background schema changes.<br>
[Issue #6515](https://github.com/pgadmin-org/pgadmin4/issues/6515) -  Fixed an issue where the query tool is unable to execute a query on Postgres 10 and below versions.<br>
[Issue #6524](https://github.com/pgadmin-org/pgadmin4/issues/6524) -  Fix the lost connection error in v7.4.<br>
[Issue #6531](https://github.com/pgadmin-org/pgadmin4/issues/6531) -  Fixed an issue where pgAdmin failed to setup role with hyphens in name.<br>
[Issue #6537](https://github.com/pgadmin-org/pgadmin4/issues/6537) -  Fixed an issue where filters are not working and query history shows empty queries.<br>
[Issue #6544](https://github.com/pgadmin-org/pgadmin4/issues/6544) -  Fix an issue where adding a sub-folder inside a folder is not working as expected in File Manager.<br>
[Issue #6556](https://github.com/pgadmin-org/pgadmin4/issues/6556) -  Fix an error 'list' object has no attribute 'strip' while attempting to populate auto-complete manually the first time.<br>
[Issue #6558](https://github.com/pgadmin-org/pgadmin4/issues/6558) -  Fixed an issue where ERD Tool can't load the saved pgerd file from Shared Storage.<br>
[Issue #6582](https://github.com/pgadmin-org/pgadmin4/issues/6582) -  Fix an issue where inserting more than 10 rows does not work correctly in View Data; only parts end up in the table.<br>
