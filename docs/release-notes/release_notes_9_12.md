# Version 9.12

Release date: 2026-02-05

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.11.

!!! warning

    Starting with this release, pgAdmin Windows installers are signed with a new code signing certificate. When installing or running pgAdmin on Windows, you should verify that the digital signature shows the certificate name as **"Open Source Developer, David John Page"**. This certificate will be used for this and future releases.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16, 17 and 18

**EDB Advanced Server**: 13, 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.0

# New features

[Issue #6451](https://github.com/pgadmin-org/pgadmin4/issues/6451) - Add new options like INHERIT and SET to the Role's membership tab.<br>
[Issue #8890](https://github.com/pgadmin-org/pgadmin4/issues/8890) - Add a new button in the query tool data output toolbar to get entire range of data.<br>
[Issue #9292](https://github.com/pgadmin-org/pgadmin4/issues/9292) - Enhance OAUTH2 and OIDC authentication support with improved claims handling and configuration options.<br>

# Housekeeping

# Bug fixes

[Issue #8916](https://github.com/pgadmin-org/pgadmin4/issues/8916) - Warn user of unsaved data output edits before page navigation.<br>
[Issue #8987](https://github.com/pgadmin-org/pgadmin4/issues/8987) - Fix Query Tool state restoration for new connections and queries.<br>
[Issue #9074](https://github.com/pgadmin-org/pgadmin4/issues/9074) - Fix pg_restore logs to distinguish UI sync issues from actual failures.<br>
[Issue #9110](https://github.com/pgadmin-org/pgadmin4/issues/9110) - Optimize checkbox selection logic in backup dialog objects tree.<br>
[Issue #9196](https://github.com/pgadmin-org/pgadmin4/issues/9196) - Fixed an issue where double click to open a file in the file manager is not working.<br>
[Issue #9223](https://github.com/pgadmin-org/pgadmin4/issues/9223) - Upgrade ID column in the database table to BigInteger to support large OID values.<br>
[Issue #9235](https://github.com/pgadmin-org/pgadmin4/issues/9235) - Fixed an issue where "View/Edit Data" shortcut opened "First 100 rows" instead of "All Rows".<br>
[Issue #9258](https://github.com/pgadmin-org/pgadmin4/issues/9258) - Ensure saved shared server passwords are re-encrypted on password change.<br>
[Issue #9260](https://github.com/pgadmin-org/pgadmin4/issues/9260) - Fixed an issue where data filter dialog removes newline character when sending SQL to the query tool.<br>
[Issue #9285](https://github.com/pgadmin-org/pgadmin4/issues/9285) - Fixed an issue where the dashboard freezes on initial render when there is a high number of locks.<br>
[Issue #9293](https://github.com/pgadmin-org/pgadmin4/issues/9293) - Fixed the SSL certificate issue while checking for the upgrade.<br>
[Issue #9332](https://github.com/pgadmin-org/pgadmin4/issues/9332) - Fixed a sorting issue in the system stats memory usage table.<br>
[Issue #9350](https://github.com/pgadmin-org/pgadmin4/issues/9350) - Disable Parameters and Membership fields when object is not new for Login and group roles.<br>
[Issue #9380](https://github.com/pgadmin-org/pgadmin4/issues/9380) - Fixed an issue where the Query History panel would auto-scroll to the top and did not preserve the scroll bar position for the selected entry.<br>
[Issue #9402](https://github.com/pgadmin-org/pgadmin4/issues/9402) - Fixed an issue where pgAdmin4 app on macOS cannot auto-update while running on a read-only volume even if present in the Applications folder.<br>
[Issue #9500](https://github.com/pgadmin-org/pgadmin4/issues/9500) - Fixed an issue where connection parameters were using localized values instead of literal values, causing connection failures.<br>
[Issue #9518](https://github.com/pgadmin-org/pgadmin4/issues/9518) - Mask the secret key for restrict option in the process watcher when restoring plain SQL file (CVE-2026-1707).<br>
[Issue #9522](https://github.com/pgadmin-org/pgadmin4/issues/9522) - Ensure the container deployment supports boolean values in yaml format.<br>
[Issue #9552](https://github.com/pgadmin-org/pgadmin4/issues/9552) - Ensure that the tooltip for the password cell is not visible.<br>
[Issue #9553](https://github.com/pgadmin-org/pgadmin4/issues/9553) - Fix pgAdmin fails when performing Backup/Restore on a PostgreSQL connection defined exclusively via pg_service.conf.<br>
[Issue #9567](https://github.com/pgadmin-org/pgadmin4/issues/9567) - Update menu bar documentation.<br>
