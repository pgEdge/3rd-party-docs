# Version 6.20

Release date: 2023-02-09

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.19.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13, 14 and 15

**EDB Advanced Server**: 10, 11, 12, 13, 14 and 15

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 15.1

# New features

[Issue #4728](https://github.com/pgadmin-org/pgadmin4/issues/4728) -  Added support for setting PostgreSQL connection parameters.<br>

# Housekeeping

[Issue #4320](https://github.com/pgadmin-org/pgadmin4/issues/4320) -  Added bundled PG utilities in the release note.<br>
[Issue #5525](https://github.com/pgadmin-org/pgadmin4/issues/5525) -  Upgrade Flask-Migrate to 4.x.<br>
[Issue #5723](https://github.com/pgadmin-org/pgadmin4/issues/5723) -  Improve performance by removing signal-based zoom-in, zoom-out, etc functionality from the runtime environment.<br>
[Issue #5794](https://github.com/pgadmin-org/pgadmin4/issues/5794) -  Use uplot for Dashboard graphs to reduce CPU usage.<br>

# Bug fixes

[Issue #5532](https://github.com/pgadmin-org/pgadmin4/issues/5532) -  Fixed an issue where the client cert location was not stored on the shared servers.<br>
[Issue #5567](https://github.com/pgadmin-org/pgadmin4/issues/5567) -  Fix orphan database connections resulting in an inability to connect to databases.<br>
[Issue #5702](https://github.com/pgadmin-org/pgadmin4/issues/5702) -  Fix an issue where role is used as username for newly added servers when opening query tool.<br>
[Issue #5705](https://github.com/pgadmin-org/pgadmin4/issues/5705) -  Ensure that all parts of the application recommend and enforce the same length of passwords.<br>
[Issue #5732](https://github.com/pgadmin-org/pgadmin4/issues/5732) -  Fixed an issue where Kerberos authentication to the server is not imported/exported.<br>
[Issue #5733](https://github.com/pgadmin-org/pgadmin4/issues/5733) -  Ensure that the system columns should not visible in the import/export data.<br>
[Issue #5746](https://github.com/pgadmin-org/pgadmin4/issues/5746) -  Increase the length of the value column of the setting table.<br>
[Issue #5748](https://github.com/pgadmin-org/pgadmin4/issues/5748) -  Fixed console error while attaching the partition.<br>
[Issue #5751](https://github.com/pgadmin-org/pgadmin4/issues/5751) -  Fix failing import servers CLI due to vulnerability fix.<br>
[Issue #5761](https://github.com/pgadmin-org/pgadmin4/issues/5761) -  Fix an issue where drag and drop object names is not working.<br>
[Issue #5763](https://github.com/pgadmin-org/pgadmin4/issues/5763) -  Ensure that keyboard hotkey to open query tool and search object should work properly.<br>
[Issue #5781](https://github.com/pgadmin-org/pgadmin4/issues/5781) -  Fixed an issue where Query history is not getting loaded with external database.<br>
[Issue #5796](https://github.com/pgadmin-org/pgadmin4/issues/5796) -  Ensure nested menu items are shown in quick search result.<br>
