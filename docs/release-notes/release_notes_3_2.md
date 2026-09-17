# Version 3.2

Release date: 2018-08-09

This release contains a number of features and fixes reported since the release of pgAdmin4 3.1

# Features

[Issue #2136](https://redmine.postgresql.org/issues/2136) - Added version number for URL's to ensure that files are only cached on a per-version basis.<br>
[Issue #2214](https://redmine.postgresql.org/issues/2214) - Add support for SCRAM password changes (requires psycopg2 >= 2.8).<br>
[Issue #3074](https://redmine.postgresql.org/issues/3074) - Add support for reset saved password.<br>
[Issue #3397](https://redmine.postgresql.org/issues/3397) - Add support for Trigger and JIT stats in the graphical query plan viewer.<br>
[Issue #3412](https://redmine.postgresql.org/issues/3412) - Add support for primary key, foreign key, unique key, indexes and triggers on partitioned tables for PG/EPAS 11.<br>
[Issue #3506](https://redmine.postgresql.org/issues/3506) - Allow the user to specify a fixed port number in the runtime to aid cookie whitelisting etc.<br>
[Issue #3510](https://redmine.postgresql.org/issues/3510) - Add a menu option to the runtime to copy the appserver URL to the clipboard.<br>

# Bug fixes

[Issue #3185](https://redmine.postgresql.org/issues/3185) - Fix the upgrade check on macOS.<br>
[Issue #3191](https://redmine.postgresql.org/issues/3191) - Fix a number of debugger execution issues.<br>
[Issue #3294](https://redmine.postgresql.org/issues/3294) - Infrastructure (and changes to the Query Tool, Dashboards and Debugger) for realtime preference handling.<br>
[Issue #3309](https://redmine.postgresql.org/issues/3309) - Fix Directory format support for backups.<br>
[Issue #3316](https://redmine.postgresql.org/issues/3316) - Support running on systems without a system tray.<br>
[Issue #3319](https://redmine.postgresql.org/issues/3319) - Cleanup and fix handling of Query Tool Cancel button status.<br>
[Issue #3363](https://redmine.postgresql.org/issues/3363) - Fix restoring of restore options for sections.<br>
[Issue #3371](https://redmine.postgresql.org/issues/3371) - Don't create a session when the /misc/ping test endpoint is called.<br>
[Issue #3446](https://redmine.postgresql.org/issues/3446) - Various procedure/function related fixes for EPAS/PG 11.<br>
[Issue #3448](https://redmine.postgresql.org/issues/3448) - Exclude system columns in Import/Export.<br>
[Issue #3457](https://redmine.postgresql.org/issues/3457) - Fix debugging of procedures in EPAS packages.<br>
[Issue #3458](https://redmine.postgresql.org/issues/3458) - pgAdmin4 should work with python 3.7.<br>
[Issue #3468](https://redmine.postgresql.org/issues/3468) - Support SSH tunneling with keys that don't have a passphrase.<br>
[Issue #3471](https://redmine.postgresql.org/issues/3471) - Ensure the SSH tunnel port number is honoured.<br>
[Issue #3511](https://redmine.postgresql.org/issues/3511) - Add support to save and clear SSH Tunnel password.<br>
[Issue #3526](https://redmine.postgresql.org/issues/3526) - COST statement should not be automatically duplicated after creating trigger function.<br>
[Issue #3527](https://redmine.postgresql.org/issues/3527) - View Data->Filtered Rows dialog should be displayed.<br>
