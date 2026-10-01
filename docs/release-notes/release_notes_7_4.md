# Version 7.4

Release date: 2023-06-29

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v7.3.

# Supported Database Servers

**PostgreSQL**: 11, 12, 13, 14 and 15

**EDB Advanced Server**: 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.3

# New features

[Issue #6367](https://github.com/pgadmin-org/pgadmin4/issues/6367) -  Added support to drop databases using the 'WITH (FORCE)' option.<br>
[Issue #6368](https://github.com/pgadmin-org/pgadmin4/issues/6368) -  Add "[ NULLS [ NOT ] DISTINCT ]" option while creating an Index.<br>
[Issue #6370](https://github.com/pgadmin-org/pgadmin4/issues/6370) -  Added 'OR REPLACE' clause while creating trigger.<br>
[Issue #6371](https://github.com/pgadmin-org/pgadmin4/issues/6371) -  Added security_invoker option while creating a view.<br>

# Housekeeping

# Bug fixes

[Issue #5306](https://github.com/pgadmin-org/pgadmin4/issues/5306) -  Fix an issue where object explorer tree crashes occasionally.<br>
[Issue #6065](https://github.com/pgadmin-org/pgadmin4/issues/6065) -  Ensure that query tool shortcuts are working properly.<br>
[Issue #6258](https://github.com/pgadmin-org/pgadmin4/issues/6258) -  Add Password exec command and Expiration time to server export JSON and also allow them to import.<br>
[Issue #6266](https://github.com/pgadmin-org/pgadmin4/issues/6266) -  When opening pgAdmin the layout should be auto reset if it is corrupted. Reset layout menu should work if layout is corrupted while using pgAdmin.<br>
[Issue #6291](https://github.com/pgadmin-org/pgadmin4/issues/6291) -  Fix an issue where loading more rows indicator will not disappear if connection is lost.<br>
[Issue #6340](https://github.com/pgadmin-org/pgadmin4/issues/6340) -  Fix an encoding error when connecting through Pgbouncer.<br>
[Issue #6352](https://github.com/pgadmin-org/pgadmin4/issues/6352) -  Fix an issue where explain plan details is showing HTML escaped characters.<br>
[Issue #6354](https://github.com/pgadmin-org/pgadmin4/issues/6354) -  Fixed an issue where queries with temporary tables in the same transaction is not working.<br>
[Issue #6363](https://github.com/pgadmin-org/pgadmin4/issues/6363) -  Fixed an issue where preview images for themes were not loading.<br>
[Issue #6420](https://github.com/pgadmin-org/pgadmin4/issues/6420) -  Fix raise notice from func/proc or code blocks are no longer displayed live.<br>
[Issue #6431](https://github.com/pgadmin-org/pgadmin4/issues/6431) -  Fix an issue where PSQL is not working if the database name have quotes or double quotes.<br>
[Issue #6435](https://github.com/pgadmin-org/pgadmin4/issues/6435) -  Fix an issue where all the menus are enabled when pgAdmin is opened and no object is selected in the object explorer.<br>
