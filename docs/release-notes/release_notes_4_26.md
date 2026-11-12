# Version 4.26

Release date: 2020-09-17

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.25.

# New features

[Issue #2042](https://redmine.postgresql.org/issues/2042) -  Added SQL Formatter support in Query Tool.<br>
[Issue #4059](https://redmine.postgresql.org/issues/4059) -  Added a new button to the query tool toolbar to open a new query tool window.<br>
[Issue #4979](https://redmine.postgresql.org/issues/4979) -  Added shared server support for admin users.<br>
[Issue #5772](https://redmine.postgresql.org/issues/5772) -  Warn the user when connecting to a server that is older than pgAdmin supports.<br>

# Housekeeping

[Issue #5332](https://redmine.postgresql.org/issues/5332) -  Improve code coverage and API test cases for Columns and Constraints (Index, Foreign Key, Check, Exclusion).<br>
[Issue #5344](https://redmine.postgresql.org/issues/5344) -  Improve code coverage and API test cases for Grant Wizard.<br>
[Issue #5774](https://redmine.postgresql.org/issues/5774) -  Improve code coverage and API test cases for Tables.<br>
[Issue #5792](https://redmine.postgresql.org/issues/5792) -  Added documentation for shared server support.<br>

# Bug fixes

[Issue #4216](https://redmine.postgresql.org/issues/4216) -  Ensure that schema names starting with 'pg' should be visible in browser tree when standard_conforming_strings is set to off.<br>
[Issue #5426](https://redmine.postgresql.org/issues/5426) -  Adjusted the height of jobstep code block to use maximum space.<br>
[Issue #5652](https://redmine.postgresql.org/issues/5652) -  Modified the 'Commit' and 'Rollback' query tool button icons.<br>
[Issue #5722](https://redmine.postgresql.org/issues/5722) -  Ensure that the user should be able to drop the database even if it is connected.<br>
[Issue #5732](https://redmine.postgresql.org/issues/5732) -  Fixed some accessibility issues.<br>
[Issue #5734](https://redmine.postgresql.org/issues/5734) -  Update the description of GIN and GiST indexes in the documentation.<br>
[Issue #5746](https://redmine.postgresql.org/issues/5746) -  Fixed an issue where --load-server does not allow loading connections that use pg_services.<br>
[Issue #5748](https://redmine.postgresql.org/issues/5748) -  Fixed incorrect reverse engineering SQL for Foreign key when creating a table.<br>
[Issue #5751](https://redmine.postgresql.org/issues/5751) -  Enable the 'Configure' and 'View log' menu option when the server taking longer than usual time to start.<br>
[Issue #5754](https://redmine.postgresql.org/issues/5754) -  Fixed an issue where schema diff is not working when providing the options to Foreign Data Wrapper, Foreign Server, and User Mapping.<br>
[Issue #5764](https://redmine.postgresql.org/issues/5764) -  Fixed SQL for Row Level Security which is incorrectly generated.<br>
[Issue #5765](https://redmine.postgresql.org/issues/5765) -  Fixed an issue in the query tool when columns are having the same name as javascript object internal functions.<br>
[Issue #5766](https://redmine.postgresql.org/issues/5766) -  Fixed string indices must be integers issue for PostgreSQL < 9.3.<br>
[Issue #5773](https://redmine.postgresql.org/issues/5773) -  Fixed an issue where the application ignores the fixed port configuration value.<br>
[Issue #5775](https://redmine.postgresql.org/issues/5775) -  Ensure that 'setup-web.sh' should work in Debian 10.<br>
[Issue #5779](https://redmine.postgresql.org/issues/5779) -  Remove illegal argument from trigger function in trigger DDL statement.<br>
[Issue #5794](https://redmine.postgresql.org/issues/5794) -  Fixed excessive CPU usage by stopping the indefinite growth of the graph dataset.<br>
[Issue #5815](https://redmine.postgresql.org/issues/5815) -  Fixed an issue where clicking on the 'Generate script' button shows a forever spinner due to pop up blocker.<br>
[Issue #5816](https://redmine.postgresql.org/issues/5816) -  Ensure that the 'CREATE SCHEMA' statement should be present in the generated script if the schema is not present in the target database.<br>
[Issue #5820](https://redmine.postgresql.org/issues/5820) -  Fixed an issue while refreshing Resource Group.<br>
[Issue #5833](https://redmine.postgresql.org/issues/5833) -  Fixed an issue where custom sequences are not visible when show system objects are set to false.<br>
[Issue #5834](https://redmine.postgresql.org/issues/5834) -  Ensure that the 'Remove Server Group' option is available in the context menu.<br>
