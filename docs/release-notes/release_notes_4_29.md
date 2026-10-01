# Version 4.29

Release date: 2020-12-10

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.28.

# New features

# Housekeeping

[Issue #5328](https://redmine.postgresql.org/issues/5328) -  Improve code coverage and API test cases for Foreign Tables.<br>
[Issue #5337](https://redmine.postgresql.org/issues/5337) -  Improve code coverage and API test cases for Views and Materialized Views.<br>
[Issue #5343](https://redmine.postgresql.org/issues/5343) -  Improve code coverage and API test cases for Debugger.<br>
[Issue #6062](https://redmine.postgresql.org/issues/6062) -  Ensure that code coverage should cover class and function declarations.<br>

# Bug fixes

[Issue #5886](https://redmine.postgresql.org/issues/5886) -  Fixed false error is shown while adding a new foreign key from the table dialog when a foreign key already exists with Auto FK Index set to true.<br>
[Issue #5943](https://redmine.postgresql.org/issues/5943) -  Ensure that folder rename should work properly in Storage Manager.<br>
[Issue #5974](https://redmine.postgresql.org/issues/5974) -  Fixed an issue where the debugger's custom tab title not applied when opened in the new browser tab.<br>
[Issue #5978](https://redmine.postgresql.org/issues/5978) -  Fixed an issue where dynamic tab title has not applied the first time for debugger panel.<br>
[Issue #5982](https://redmine.postgresql.org/issues/5982) -  Fixed documentation issue where JSON is not valid.<br>
[Issue #5983](https://redmine.postgresql.org/issues/5983) -  Added the appropriate server icon based on the server type in the new connection dialog.<br>
[Issue #5985](https://redmine.postgresql.org/issues/5985) -  Fixed an issue where the process watcher dialog throws an error for the database server which is already removed.<br>
[Issue #5991](https://redmine.postgresql.org/issues/5991) -  Ensure that dirty indicator (*) should not be visible when renaming the tabs.<br>
[Issue #5992](https://redmine.postgresql.org/issues/5992) -  Fixed an issue where escape character is shown when the server/database name has some special characters.<br>
[Issue #5998](https://redmine.postgresql.org/issues/5998) -  Fixed an issue where schema diff doesn't show the result of compare if source schema has tables with RLS.<br>
[Issue #6003](https://redmine.postgresql.org/issues/6003) -  Fixed an issue where an illegal argument is showing for trigger SQL when a trigger is created for View.<br>
[Issue #6022](https://redmine.postgresql.org/issues/6022) -  Fixed an issue where shared servers import is failing.<br>
[Issue #6072](https://redmine.postgresql.org/issues/6072) -  Fixed DLL load failed while importing bcrypt.<br>
