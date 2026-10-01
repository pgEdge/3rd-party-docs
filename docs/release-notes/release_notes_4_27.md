# Version 4.27

Release date: 2020-10-15

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.26.

# New features

[Issue #1402](https://redmine.postgresql.org/issues/1402) -  Added Macro support.<br>
[Issue #2519](https://redmine.postgresql.org/issues/2519) -  Added support to view trigger function under the respective trigger node.<br>
[Issue #3794](https://redmine.postgresql.org/issues/3794) -  Allow user to change the database connection from an open query tool tab.<br>
[Issue #5200](https://redmine.postgresql.org/issues/5200) -  Added support to ignore the owner while comparing objects in the Schema Diff tool.<br>
[Issue #5857](https://redmine.postgresql.org/issues/5857) -  Added documentation for Macro support.<br>

# Housekeeping

[Issue #5330](https://redmine.postgresql.org/issues/5330) -  Improve code coverage and API test cases for Functions.<br>
[Issue #5395](https://redmine.postgresql.org/issues/5395) -  Added RESQL/MSQL test cases for Functions.<br>
[Issue #5497](https://redmine.postgresql.org/issues/5497) -  Merged the latest code of 'pgcli' used for the autocomplete feature.<br>

# Bug fixes

[Issue #4806](https://redmine.postgresql.org/issues/4806) -  Added useful message when the explain plan is not used and empty.<br>
[Issue #4855](https://redmine.postgresql.org/issues/4855) -  Fixed an issue where file extension is stripped on renaming a file.<br>
[Issue #5131](https://redmine.postgresql.org/issues/5131) -  Ensure that 'ctrl + a' shortcut does not move the cursor in SQL editor.<br>
[Issue #5417](https://redmine.postgresql.org/issues/5417) -  Fixed and improve API test cases for the schema diff tool.<br>
[Issue #5739](https://redmine.postgresql.org/issues/5739) -  Ensure that the import/export feature should work with SSH Tunnel.<br>
[Issue #5802](https://redmine.postgresql.org/issues/5802) -  Remove maximum length on the password field in the server dialog.<br>
[Issue #5807](https://redmine.postgresql.org/issues/5807) -  Fixed an issue where a column is renamed and then removed, then the drop SQL query takes the wrong column name.<br>
[Issue #5826](https://redmine.postgresql.org/issues/5826) -  Fixed an issue where schema diff is showing identical table as different due to default vacuum settings.<br>
[Issue #5830](https://redmine.postgresql.org/issues/5830) -  Fixed reverse engineering SQL where parenthesis is not properly arranged for View/MView definition.<br>
[Issue #5835](https://redmine.postgresql.org/issues/5835) -  Fixed 'can't execute an empty query' message if the user change the option of Auto FK Index.<br>
[Issue #5839](https://redmine.postgresql.org/issues/5839) -  Ensure that multiple extensions can be dropped from the properties tab.<br>
[Issue #5841](https://redmine.postgresql.org/issues/5841) -  Fixed an issue where the server is not able to connect using the service.<br>
[Issue #5843](https://redmine.postgresql.org/issues/5843) -  Fixed an issue where the 'PARALLEL UNSAFE' option is missing from reverse engineering SQL of function/procedure.<br>
[Issue #5845](https://redmine.postgresql.org/issues/5845) -  Fixed an issue where the query tool is not fetching more than 1000 rows for the table does not have any primary key.<br>
[Issue #5853](https://redmine.postgresql.org/issues/5853) -  Fixed an issue where 'Rows X' column values were not visible properly for Explain Analyze in Dark theme.<br>
[Issue #5855](https://redmine.postgresql.org/issues/5855) -  Ensure that the user should be able to change the start value of the existing sequence.<br>
[Issue #5861](https://redmine.postgresql.org/issues/5861) -  Ensure that the 'Remove Server' option should be visible in the context menu.<br>
[Issue #5867](https://redmine.postgresql.org/issues/5867) -  Fixed an issue where some properties are not being updated correctly for the shared server.<br>
[Issue #5882](https://redmine.postgresql.org/issues/5882) -  Fixed invalid literal issue when fetching dependencies for Materialized View.<br>
[Issue #5885](https://redmine.postgresql.org/issues/5885) -  Fixed an issue where the user is unable to change the macro name.<br>
