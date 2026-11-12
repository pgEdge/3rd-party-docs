# Version 4.21

Release date: 2020-04-30

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.20.

# New features

[Issue #2172](https://redmine.postgresql.org/issues/2172) -  Added search object functionality.<br>
[Issue #2186](https://redmine.postgresql.org/issues/2186) -  Added LDAP authentication support.<br>
[Issue #4636](https://redmine.postgresql.org/issues/4636) -  Added job step and job schedule disable icons to identify it quickly within the browser tree.<br>
[Issue #5181](https://redmine.postgresql.org/issues/5181) -  Added support for parameter toast_tuple_target and parallel_workers of the table.<br>
[Issue #5263](https://redmine.postgresql.org/issues/5263) -  Added support of Foreign Tables to the Schema Diff.<br>
[Issue #5264](https://redmine.postgresql.org/issues/5264) -  Added support of Packages, Sequences and Synonyms to the Schema Diff.<br>
[Issue #5348](https://redmine.postgresql.org/issues/5348) -  Documentation of LDAP authentication support.<br>
[Issue #5353](https://redmine.postgresql.org/issues/5353) -  Added an option to prevent a browser tab being opened at startup.<br>
[Issue #5399](https://redmine.postgresql.org/issues/5399) -  Warn the user if an unsupported, deprecated or unknown browser is detected.<br>

# Housekeeping

[Issue #4620](https://redmine.postgresql.org/issues/4620) -  Add Reverse Engineered and Modified SQL tests for procedures.<br>
[Issue #4623](https://redmine.postgresql.org/issues/4623) -  Add Reverse Engineered and Modified SQL tests for pgAgent jobs.<br>

# Bug fixes

[Issue #1257](https://redmine.postgresql.org/issues/1257) -  Ensure all object types have a "System XXX?" property.<br>
[Issue #2813](https://redmine.postgresql.org/issues/2813) -  Ensure that the password prompt should not be visible if the database server is in trust authentication mode.<br>
[Issue #3495](https://redmine.postgresql.org/issues/3495) -  Fixed an issue where the query tool unable to load the file which contains the BOM marker.<br>
[Issue #3523](https://redmine.postgresql.org/issues/3523) -  Fixed an issue where right-clicking a browser object does not apply to the object on which right-click was fired.<br>
[Issue #3645](https://redmine.postgresql.org/issues/3645) -  Ensure that the start and end date should be deleted when clear the selection for pgAgent Job.<br>
[Issue #3900](https://redmine.postgresql.org/issues/3900) -  Added multiple drop/delete functionality for the table constraints.<br>
[Issue #3947](https://redmine.postgresql.org/issues/3947) -  Fixed copy-paste row issues in View/Edit Data.<br>
[Issue #3972](https://redmine.postgresql.org/issues/3972) -  Modified keyboard shortcuts in Query Tool for OSX native support.<br>
[Issue #3988](https://redmine.postgresql.org/issues/3988) -  Fixed cursor disappeared issue in the query editor for some of the characters when zoomed out.<br>
[Issue #4180](https://redmine.postgresql.org/issues/4180) -  Fixed mouse click issue where it does not select an object in Browser unless the pointer is over the object.<br>
[Issue #4206](https://redmine.postgresql.org/issues/4206) -  Ensure that the grant wizard should be closed on pressing the ESC key.<br>
[Issue #4292](https://redmine.postgresql.org/issues/4292) -  Added dark mode support for the configuration dialog on Windows/macOS runtime.<br>
[Issue #4440](https://redmine.postgresql.org/issues/4440) -  Ensure the DROP statements in reverse engineered SQL are properly quoted for all objects.<br>
[Issue #4445](https://redmine.postgresql.org/issues/4445) -  Ensure all object names in the title line of the reverse-engineered SQL are not quoted.<br>
[Issue #4504](https://redmine.postgresql.org/issues/4504) -  Fixed an issue where like options should be disabled if the relation is not selected while creating a table.<br>
[Issue #4512](https://redmine.postgresql.org/issues/4512) -  Fixed calendar opening issue on the exception tab inside the schedules tab of pgAgent.<br>
[Issue #4545](https://redmine.postgresql.org/issues/4545) -  Fixed an issue wherein grant wizard the last object is not selectable.<br>
[Issue #4573](https://redmine.postgresql.org/issues/4573) -  Ensure that if the delimiter is set other than comma then download the file as '.txt' file.<br>
[Issue #4684](https://redmine.postgresql.org/issues/4684) -  Fixed encoding issue while saving data in encoded charset other than 'utf-8'.<br>
[Issue #4709](https://redmine.postgresql.org/issues/4709) -  Added schema-qualified dictionary names in FTS configuration to avoid confusion of duplicate names.<br>
[Issue #4856](https://redmine.postgresql.org/issues/4856) -  Enable the save button by default when a query tool is opened with CREATE or other scripts.<br>
[Issue #4858](https://redmine.postgresql.org/issues/4858) -  Fixed python exception error when user tries to download the CSV and there is a connection issue.<br>
[Issue #4864](https://redmine.postgresql.org/issues/4864) -  Make the configuration window in runtime to auto-resize.<br>
[Issue #4873](https://redmine.postgresql.org/issues/4873) -  Fixed an issue when changing the comments of the procedure with arguments gives error in case of overloading.<br>
[Issue #4946](https://redmine.postgresql.org/issues/4946) -  Fixed an issue when the user creates a temporary table with 'on commit drop as' clause.<br>
[Issue #4957](https://redmine.postgresql.org/issues/4957) -  Ensure that Constraint Trigger, Deferrable, Deferred option should be disabled when the user selects EDB-SPL function for the trigger.<br>
[Issue #4969](https://redmine.postgresql.org/issues/4969) -  Fixed an issue where changing the values of columns with JSONB or JSON types to NULL.<br>
[Issue #5007](https://redmine.postgresql.org/issues/5007) -  Ensure index dropdown should have existing indexes while creating unique constraints.<br>
[Issue #5043](https://redmine.postgresql.org/issues/5043) -  Fixed an issue where columns names should be visible in the order of their creation in the browser tree.<br>
[Issue #5053](https://redmine.postgresql.org/issues/5053) -  Fixed an issue where changing the columns in the existing view throws an error.<br>
[Issue #5157](https://redmine.postgresql.org/issues/5157) -  Ensure that default sort order should be using the primary key in View/Edit data.<br>
[Issue #5180](https://redmine.postgresql.org/issues/5180) -  Fixed an issue where the autovacuum_enabled parameter is added automatically in the RE-SQL when the table has been created using the WITH clause.<br>
[Issue #5210](https://redmine.postgresql.org/issues/5210) -  Ensure that text larger than underlying field size should not be truncated automatically.<br>
[Issue #5213](https://redmine.postgresql.org/issues/5213) -  Fixed an issue when the user performs refresh on a large size materialized view.<br>
[Issue #5227](https://redmine.postgresql.org/issues/5227) -  Fixed an issue where user cannot be added if many users are already exists.<br>
[Issue #5268](https://redmine.postgresql.org/issues/5268) -  Fixed generated SQL when any token in FTS Configuration or any option in FTS Dictionary is changed.<br>
[Issue #5270](https://redmine.postgresql.org/issues/5270) -  Ensure that OID should be shown in properties for Synonyms.<br>
[Issue #5275](https://redmine.postgresql.org/issues/5275) -  Fixed tab key navigation issue for parameters in table dialog.<br>
[Issue #5302](https://redmine.postgresql.org/issues/5302) -  Fixed an issue where difference SQL is not seen in the schema diff tool for Types.<br>
[Issue #5314](https://redmine.postgresql.org/issues/5314) -  Ensure that switch cell is in sync with switch control for accessibility.<br>
[Issue #5315](https://redmine.postgresql.org/issues/5315) -  Fixed an issue where schema diff showing changes in the identical domain constraints.<br>
[Issue #5350](https://redmine.postgresql.org/issues/5350) -  Fixed an issue where schema diff marks an identical table as different.<br>
[Issue #5351](https://redmine.postgresql.org/issues/5351) -  Fixed compilation warnings while building pgAdmin.<br>
[Issue #5352](https://redmine.postgresql.org/issues/5352) -  Fixed the rightmost and bottom tooltip crop issues in the explain query plan.<br>
[Issue #5356](https://redmine.postgresql.org/issues/5356) -  Fixed modified SQL issue while adding an exception in pgAgent job schedule.<br>
[Issue #5361](https://redmine.postgresql.org/issues/5361) -  Fixes an issue where pgAdmin4 GUI does not display properly in IE 11.<br>
[Issue #5362](https://redmine.postgresql.org/issues/5362) -  Fixed an issue where the identical packages and sequences visible as different in the schema diff tool.<br>
[Issue #5366](https://redmine.postgresql.org/issues/5366) -  Added alert message to Reset Layout if any of the panels from Query Tool failed to load.<br>
[Issue #5371](https://redmine.postgresql.org/issues/5371) -  Fixed tab key navigation for some dialogs.<br>
[Issue #5375](https://redmine.postgresql.org/issues/5375) -  Fixed an issue where the Mode cell of argument grid does not appear completely in the Functions dialog.<br>
[Issue #5383](https://redmine.postgresql.org/issues/5383) -  Fixed syntax error while refreshing the existing synonyms.<br>
[Issue #5387](https://redmine.postgresql.org/issues/5387) -  Fixed an issue where the mode is not shown in the properties dialog of functions/procedures if all the arguments are "IN" arguments.<br>
[Issue #5396](https://redmine.postgresql.org/issues/5396) -  Fixed an issue where the search object module unable to locate the object in the browser tree.<br>
[Issue #5400](https://redmine.postgresql.org/issues/5400) -  Fixed internal server error when the database server is logged in with non-super user.<br>
[Issue #5401](https://redmine.postgresql.org/issues/5401) -  Fixed search object issue when the object name contains special characters.<br>
[Issue #5402](https://redmine.postgresql.org/issues/5402) -  Fixed an issue where the checkbox is not visible on Configuration dialog in runtime for the dark theme.<br>
[Issue #5409](https://redmine.postgresql.org/issues/5409) -  Fixed validation issue in Synonyms node.<br>
[Issue #5410](https://redmine.postgresql.org/issues/5410) -  Fixed an issue while removing the package body showing wrong modified SQL.<br>
[Issue #5415](https://redmine.postgresql.org/issues/5415) -  Ensure that the query tool context menu should work on the collection nodes.<br>
[Issue #5419](https://redmine.postgresql.org/issues/5419) -  Ensure that the user should not be able to change the authentication source.<br>
[Issue #5420](https://redmine.postgresql.org/issues/5420) -  Ensure error should be handled properly when LDAP user is created with the same name.<br>
[Issue #5430](https://redmine.postgresql.org/issues/5430) -  Added title to the login page.<br>
[Issue #5432](https://redmine.postgresql.org/issues/5432) -  Fixed an issue where an internal user is not created if the authentication source is set to internal and ldap.<br>
[Issue #5439](https://redmine.postgresql.org/issues/5439) -  Fixed an issue where the user is not able to create a server if login with an LDAP account.<br>
[Issue #5441](https://redmine.postgresql.org/issues/5441) -  Fixed an issue where the search object not able to locate pg_toast_* tables in the pg_toast schema.<br>
[Issue #5447](https://redmine.postgresql.org/issues/5447) -  Fixed failed to fetch utility error when click on refresh(any option) materialized view.<br>
