# Version 3.3

Release date: 2018-09-06

This release contains a number of features and fixes reported since the release of pgAdmin4 3.2

# Features

[Issue #1407](https://redmine.postgresql.org/issues/1407) - Add a geometry viewer that can render PostGIS data on a blank canvas or various map sources.<br>
[Issue #3503](https://redmine.postgresql.org/issues/3503) - Added new backup/restore options for PostgreSQL 11. Added dump options for 'pg_dumpall'.<br>
[Issue #3553](https://redmine.postgresql.org/issues/3553) - Add a Spanish translation.<br>

# Bug fixes

[Issue #3136](https://redmine.postgresql.org/issues/3136) - Stabilise feature tests for continuous running on CI systems.<br>
[Issue #3191](https://redmine.postgresql.org/issues/3191) - Fixed debugger execution issues.<br>
[Issue #3313](https://redmine.postgresql.org/issues/3313) - Ensure 'select all' and 'unselect all' working properly for pgAgent schedule.<br>
[Issue #3325](https://redmine.postgresql.org/issues/3325) - Fix sort/filter dialog issue where it incorrectly requires ASC/DESC.<br>
[Issue #3347](https://redmine.postgresql.org/issues/3347) - Ensure backup should work with '--data-only' and '--schema-only' for any format.<br>
[Issue #3407](https://redmine.postgresql.org/issues/3407) - Fix keyboard shortcuts layout in the preferences panel.<br>
[Issue #3420](https://redmine.postgresql.org/issues/3420) - Merge pgcli code with version 1.10.3, which is used for auto complete feature.<br>
[Issue #3461](https://redmine.postgresql.org/issues/3461) - Ensure that refreshing a node also updates the Property list.<br>
[Issue #3525](https://redmine.postgresql.org/issues/3525) - Ensure that refresh button on dashboard should refresh the table.<br>
[Issue #3528](https://redmine.postgresql.org/issues/3528) - Handle connection errors properly in the Query Tool.<br>
[Issue #3547](https://redmine.postgresql.org/issues/3547) - Make session implementation thread safe<br>
[Issue #3548](https://redmine.postgresql.org/issues/3548) - Ensure external table node should be visible only for GPDB.<br>
[Issue #3554](https://redmine.postgresql.org/issues/3554) - Fix auto scrolling issue in debugger on step in and step out.<br>
[Issue #3558](https://redmine.postgresql.org/issues/3558) - Fix sort/filter dialog editing issue.<br>
[Issue #3561](https://redmine.postgresql.org/issues/3561) - Ensure sort/filter dialog should display proper message after losing database connection.<br>
[Issue #3578](https://redmine.postgresql.org/issues/3578) - Ensure sql for Role should be visible in SQL panel for GPDB.<br>
[Issue #3579](https://redmine.postgresql.org/issues/3579) - When building the Windows installer, copy system Python packages before installing dependencies to ensure we don't end up with older versions than intended.<br>
[Issue #3604](https://redmine.postgresql.org/issues/3604) - Correct the documentation of View/Edit data.<br>
