# Version 4.9

Release date: 2019-06-27

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.8.

# New features

[Issue #3174](https://redmine.postgresql.org/issues/3174) - Visually distinguish simple tables from tables that are inherited and from which other tables are inherited.<br>

# Housekeeping

[Issue #4202](https://redmine.postgresql.org/issues/4202) - Add a framework for testing reversed engineered SQL and CRUD API endpoints.<br>

# Bug fixes

[Issue #3994](https://redmine.postgresql.org/issues/3994) - Fix issue where the dependencies tab for inherited tables/foreign keys shows partial text.<br>
[Issue #4036](https://redmine.postgresql.org/issues/4036) - Allow editing of data where a primary key column includes a % sign in the value.<br>
[Issue #4171](https://redmine.postgresql.org/issues/4171) - Fix issue where reverse engineered SQL was failing for foreign tables, if it had "=" in the options.<br>
[Issue #4195](https://redmine.postgresql.org/issues/4195) - Fix keyboard navigation in "inner" tabsets such as the Query Tool and Debugger.<br>
[Issue #4228](https://redmine.postgresql.org/issues/4228) - Ensure the correct label is used in panel headers when viewing filtered rows.<br>
[Issue #4253](https://redmine.postgresql.org/issues/4253) - Fix issue where new column should be created with Default value.<br>
[Issue #4283](https://redmine.postgresql.org/issues/4283) - Initial support for PostgreSQL 12.<br>
[Issue #4288](https://redmine.postgresql.org/issues/4288) - Initial support for PostgreSQL 12.<br>
[Issue #4290](https://redmine.postgresql.org/issues/4290) - Initial support for PostgreSQL 12.<br>
[Issue #4255](https://redmine.postgresql.org/issues/4255) - Prevent the geometry viewer grabbing key presses when not in focus under Firefox, IE and Edge.<br>
[Issue #4306](https://redmine.postgresql.org/issues/4306) - Prevent the "Please login to access this page" message displaying multiple times.<br>
[Issue #4310](https://redmine.postgresql.org/issues/4310) - Ensure that the Return key can be used to submit the Master Password dialogue.<br>
[Issue #4317](https://redmine.postgresql.org/issues/4317) - Ensure that browser auto-fill doesn't cause Help pages to be opened unexpectedly.<br>
[Issue #4320](https://redmine.postgresql.org/issues/4320) - Fix issue where SSH tunnel connection using password is failing, it's regression of Master Password.<br>
[Issue #4329](https://redmine.postgresql.org/issues/4329) - Fix an initialisation error when two functions with parameters are debugged in parallel.<br>
[Issue #4343](https://redmine.postgresql.org/issues/4343) - Fix issue where property dialog of column should open properly for EPAS v12.<br>
[Issue #4345](https://redmine.postgresql.org/issues/4345) - Capitalize the word 'export' used in Import/Export module.<br>
[Issue #4349](https://redmine.postgresql.org/issues/4349) - Ensure strings are properly encoded in the Query History.<br>
[Issue #4350](https://redmine.postgresql.org/issues/4350) - Ensure we include the CSRF token when uploading files.<br>
[Issue #4357](https://redmine.postgresql.org/issues/4357) - Fix connection restoration issue when pgAdmin server is restarted and the page is refreshed.<br>
[Issue #4360](https://redmine.postgresql.org/issues/4360) - Ensure the debugger control buttons are only enabled once initialisation is complete.<br>
[Issue #4362](https://redmine.postgresql.org/issues/4362) - Remove additional "SETOF" included when generating CREATE scripts for trigger functions.<br>
[Issue #4365](https://redmine.postgresql.org/issues/4365) - Fix help links for backup globals and backup server.<br>
[Issue #4367](https://redmine.postgresql.org/issues/4367) - Fix an XSS issue seen in View/Edit data mode if a column name includes HTML.<br>
[Issue #4378](https://redmine.postgresql.org/issues/4378) - Ensure Python escaping matched JS escaping and fix a minor XSS issue in the Query Tool that required superuser access to trigger.<br>
[Issue #4380](https://redmine.postgresql.org/issues/4380) - Ensure that both columns and partitions can be edited at the same time in the table dialog.<br>
[Issue #4386](https://redmine.postgresql.org/issues/4386) - Fix an XSS issue when username contains XSS vulnerable text.<br>
