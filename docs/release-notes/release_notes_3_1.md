# Version 3.1

Release date: 2018-06-28

This release contains a number of features and fixes reported since the release of pgAdmin4 3.0

# Features

[Issue #1447](https://redmine.postgresql.org/issues/1447) - Add support for SSH tunneled connections<br>
[Issue #2686](https://redmine.postgresql.org/issues/2686) - Add an option to auto-complete keywords in upper case<br>
[Issue #3204](https://redmine.postgresql.org/issues/3204) - Add support for LISTEN/NOTIFY in the Query Tool<br>
[Issue #3273](https://redmine.postgresql.org/issues/3273) - Allow sorting in the file dialogue<br>
[Issue #3362](https://redmine.postgresql.org/issues/3362) - Function and procedure support for PG11<br>
[Issue #3388](https://redmine.postgresql.org/issues/3388) - Allow the connection timeout to be configured on a per-server basis<br>

# Bug fixes

[Issue #1220](https://redmine.postgresql.org/issues/1220) - Backup and Restore should not be started if database name contains "=" symbol<br>
[Issue #1221](https://redmine.postgresql.org/issues/1221) - Maintenance should not be started if database name contains "=" symbol<br>
[Issue #3179](https://redmine.postgresql.org/issues/3179) - Fix an error generating SQL for trigger functions<br>
[Issue #3238](https://redmine.postgresql.org/issues/3238) - Standardise the error handling for parsing of JSON response messages from the server<br>
[Issue #3250](https://redmine.postgresql.org/issues/3250) - Fix handling of SQL_ASCII data in the Query Tool<br>
[Issue #3257](https://redmine.postgresql.org/issues/3257) - Catch errors when trying to EXPLAIN an invalid query<br>
[Issue #3277](https://redmine.postgresql.org/issues/3277) - Ensure server cleanup on exit only happens if the server actually started up<br>
[Issue #3284](https://redmine.postgresql.org/issues/3284) - F5 key should work to refresh Browser tree<br>
[Issue #3289](https://redmine.postgresql.org/issues/3289) - Fix handling of SQL_ASCII data in the Query Tool<br>
[Issue #3290](https://redmine.postgresql.org/issues/3290) - Close button added to the alertify message box, which pops up in case of backend error<br>
[Issue #3295](https://redmine.postgresql.org/issues/3295) - Ensure the debugger gets focus when loaded so shortcut keys work as expected<br>
[Issue #3298](https://redmine.postgresql.org/issues/3298) - Fixed Query Tool keyboard issue where arrow keys were not behaving as expected for execute options dropdown<br>
[Issue #3303](https://redmine.postgresql.org/issues/3303) - Fix a Japanese translation error that could prevent the server starting up<br>
[Issue #3306](https://redmine.postgresql.org/issues/3306) - Fixed display SQL of table with index for Greenplum database<br>
[Issue #3307](https://redmine.postgresql.org/issues/3307) - Allow connections to servers with port numbers < 1024 which may be seen in container environments<br>
[Issue #3308](https://redmine.postgresql.org/issues/3308) - Fixed issue where icon for Partitioned tables was the same as Non Partitioned tables for Greenplum database<br>
[Issue #3310](https://redmine.postgresql.org/issues/3310) - Fixed layout of the alertify error message in the Query Tool<br>
[Issue #3324](https://redmine.postgresql.org/issues/3324) - Fix the template loader to work reliably under Windows (fixing external tables under Greenplum)<br>
[Issue #3333](https://redmine.postgresql.org/issues/3333) - Ensure the runtime core application is setup before trying to access any settings<br>
[Issue #3342](https://redmine.postgresql.org/issues/3342) - Set SESSION_COOKIE_SAMESITE='Lax' per Flask recommendation to prevents sending cookies with CSRF-prone requests from external sites, such as submitting a form<br>
[Issue #3353](https://redmine.postgresql.org/issues/3353) - Handle errors properly if they occur when renaming a database<br>
[Issue #3356](https://redmine.postgresql.org/issues/3356) - Include the schema name on RE-SQL for packages<br>
[Issue #3374](https://redmine.postgresql.org/issues/3374) - Fix autocomplete<br>
[Issue #3392](https://redmine.postgresql.org/issues/3392) - Fix IPv6 support in the container build<br>
[Issue #3409](https://redmine.postgresql.org/issues/3409) - Avoid an exception on GreenPlum when retrieving RE-SQL on a table<br>
[Issue #3411](https://redmine.postgresql.org/issues/3411) - Fix a French translation error that could prevent the server starting up<br>
[Issue #3431](https://redmine.postgresql.org/issues/3431) - Fix the RE-SQL generation for GreenPlum external tables<br>
