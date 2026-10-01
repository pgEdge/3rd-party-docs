# Version 4.13

Release date: 2019-09-19

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.12.

# New features

[Issue #2828](https://redmine.postgresql.org/issues/2828) -  Added Gather Merge, Named Tuple Store Scan and Table Function Scan icon for explain module.<br>
[Issue #4553](https://redmine.postgresql.org/issues/4553) -  Don't wait for the database connection before rendering the Query Tool UI, for improved UX.<br>
[Issue #4651](https://redmine.postgresql.org/issues/4651) -  Allow configuration options to be set from the environment in the container distribution.<br>
[Issue #4667](https://redmine.postgresql.org/issues/4667) -  Ensure editable and read-only columns in Query Tool should be identified by icons and tooltips in the column header.<br>
[Issue #4691](https://redmine.postgresql.org/issues/4691) -  Add an Italian translation.<br>
[Issue #4752](https://redmine.postgresql.org/issues/4752) -  Refactor Dockerfile to avoid needing to run supporting scripts (i.e. 'docker build .' will work) and minimise layers.<br>

# Housekeeping

[Issue #4575](https://redmine.postgresql.org/issues/4575) -  Add Reverse Engineered SQL tests for Schemas.<br>
[Issue #4576](https://redmine.postgresql.org/issues/4576) -  Add Reverse Engineered SQL tests for Views.<br>
[Issue #4600](https://redmine.postgresql.org/issues/4600) -  Add Reverse Engineered SQL tests for Rules.<br>
[Issue #4616](https://redmine.postgresql.org/issues/4616) -  Add Reverse Engineered and Modified SQL tests for Foreign Keys.<br>
[Issue #4617](https://redmine.postgresql.org/issues/4617) -  Add Reverse Engineered and Modified SQL tests for Foreign Servers.<br>
[Issue #4618](https://redmine.postgresql.org/issues/4618) -  Add Reverse Engineered and Modified SQL tests for Foreign Tables.<br>
[Issue #4619](https://redmine.postgresql.org/issues/4619) -  Add Reverse Engineered and Modified SQL tests for FTS Templates.<br>
[Issue #4621](https://redmine.postgresql.org/issues/4621) -  Add Reverse Engineered and Modified SQL tests for Indexes.<br>
[Issue #4624](https://redmine.postgresql.org/issues/4624) -  Add Reverse Engineered and Modified SQL tests for Primary Keys.<br>
[Issue #4627](https://redmine.postgresql.org/issues/4627) -  Add Reverse Engineered and Modified SQL tests for User Mappings.<br>
[Issue #4690](https://redmine.postgresql.org/issues/4690) -  Add Modified SQL tests for Resource Group.<br>

# Bug fixes

[Issue #2706](https://redmine.postgresql.org/issues/2706) -  Added ProjectSet icon for explain module.<br>
[Issue #3778](https://redmine.postgresql.org/issues/3778) -  Ensure Boolean columns should be editable using keyboard keys.<br>
[Issue #3936](https://redmine.postgresql.org/issues/3936) -  Further code refactoring to stabilise the Feature Tests.<br>
[Issue #4381](https://redmine.postgresql.org/issues/4381) -  Fix an issue where oid column should not be pasted when copy/paste row is used on query output containing the oid column.<br>
[Issue #4408](https://redmine.postgresql.org/issues/4408) -  Fix display of validation error message in SlickGrid cells.<br>
[Issue #4412](https://redmine.postgresql.org/issues/4412) -  Fix issue where Validated switch option is inverted for the Foreign Key.<br>
[Issue #4419](https://redmine.postgresql.org/issues/4419) -  Fix a debugger error when using Python 2.7.<br>
[Issue #4461](https://redmine.postgresql.org/issues/4461) -  Fix error while importing data to a table using Import/Export dialog and providing "Not null columns" option.<br>
[Issue #4486](https://redmine.postgresql.org/issues/4486) -  Ensure View should be created with special characters.<br>
[Issue #4487](https://redmine.postgresql.org/issues/4487) -  Ensure Boolean columns should be editable in View/Edit data and Query Tool.<br>
[Issue #4577](https://redmine.postgresql.org/issues/4577) -  Fix an error that could be seen when click on any system column of a table.<br>
[Issue #4584](https://redmine.postgresql.org/issues/4584) -  Unescape HTML entities in database names in the Query Tool title bar.<br>
[Issue #4631](https://redmine.postgresql.org/issues/4631) -  Add editor options for plain text mode and to disable block folding to workaround rendering speed issues in CodeMirror with very large scripts.<br>
[Issue #4642](https://redmine.postgresql.org/issues/4642) -  Ensure port and username should not be mandatory when a service is provided.<br>
[Issue #4643](https://redmine.postgresql.org/issues/4643) -  Fix Truncate option deselect issue for compound triggers.<br>
[Issue #4644](https://redmine.postgresql.org/issues/4644) -  Fix length and precision enable/disable issue when changing the data type for Domain node.<br>
[Issue #4650](https://redmine.postgresql.org/issues/4650) -  Fix SQL tab issue for Views. It's a regression of compound triggers.<br>
[Issue #4657](https://redmine.postgresql.org/issues/4657) -  Fix PGADMIN_SERVER_JSON_FILE environment variable support in the container.<br>
[Issue #4663](https://redmine.postgresql.org/issues/4663) -  Fix exception in query history for python 2.7.<br>
[Issue #4674](https://redmine.postgresql.org/issues/4674) -  Fix query tool launch error if user name contain html characters.<br>
[Issue #4681](https://redmine.postgresql.org/issues/4681) -  Increase cache control max age for static files to improve performance over longer run.<br>
[Issue #4698](https://redmine.postgresql.org/issues/4698) -  Fix SQL issue of length and precision when changing the data type of Column.<br>
[Issue #4702](https://redmine.postgresql.org/issues/4702) -  Fix modified SQL for Index when reset the value of Fill factor and Clustered?.<br>
[Issue #4703](https://redmine.postgresql.org/issues/4703) -  Fix reversed engineered SQL for btree Index when provided sort order and NULLs.<br>
[Issue #4726](https://redmine.postgresql.org/issues/4726) -  Ensure sequence with negative value should be created.<br>
[Issue #4727](https://redmine.postgresql.org/issues/4727) -  Fix issue where EXEC script doesn't write the complete script for Procedures.<br>
[Issue #4736](https://redmine.postgresql.org/issues/4736) -  Fix query tool and view data issue with the Italian language.<br>
[Issue #4742](https://redmine.postgresql.org/issues/4742) -  Ensure Primary Key should be created with Index.<br>
[Issue #4750](https://redmine.postgresql.org/issues/4750) -  Fix query history exception for Python 3.6.<br>
