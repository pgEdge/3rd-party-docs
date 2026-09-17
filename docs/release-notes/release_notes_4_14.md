# Version 4.14

Release date: 2019-10-17

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.13.

# New features

[Issue #3009](https://redmine.postgresql.org/issues/3009) -  Added Copy with headers functionality when copy data from Query Tool/View Data.<br>
[Issue #4778](https://redmine.postgresql.org/issues/4778) -  Implemented the Query Plan Analyser.<br>
[Issue #4823](https://redmine.postgresql.org/issues/4823) -  Include PostgreSQL 12 binaries in the container.<br>

# Housekeeping

[Issue #4472](https://redmine.postgresql.org/issues/4472) -  Add Reverse Engineered and Modified SQL tests for Synonyms.<br>
[Issue #4628](https://redmine.postgresql.org/issues/4628) -  Add Reverse Engineered and Modified SQL tests for Unique Constraints.<br>
[Issue #4701](https://redmine.postgresql.org/issues/4701) -  Optimize Webpack to improve overall performance.<br>

# Bug fixes

[Issue #3386](https://redmine.postgresql.org/issues/3386) -  Ensure backup a partition table should not backup the whole database.<br>
[Issue #4199](https://redmine.postgresql.org/issues/4199) -  Ensure that 'ENTER' key in the data filter should not run the query.<br>
[Issue #4590](https://redmine.postgresql.org/issues/4590) -  Fix issue where backup fails for schema name that needs quoting.<br>
[Issue #4728](https://redmine.postgresql.org/issues/4728) -  Highlighted the color of closing or opening parenthesis when user select them in CodeMirror.<br>
[Issue #4751](https://redmine.postgresql.org/issues/4751) -  Fix issue where export job fails when deselecting all the columns.<br>
[Issue #4753](https://redmine.postgresql.org/issues/4753) -  Fix an error where 'false' string is displayed when we add a new parameter in the Parameters tab, also clear the old value when the user changes the parameter name.<br>
[Issue #4755](https://redmine.postgresql.org/issues/4755) -  Ensure that pgAdmin should work behind reverse proxy if the inbuilt server is used as it is.<br>
[Issue #4756](https://redmine.postgresql.org/issues/4756) -  Fix issue where pgAdmin does not load completely if loaded in an iframe.<br>
[Issue #4760](https://redmine.postgresql.org/issues/4760) -  Ensure the search path should not be quoted for Database.<br>
[Issue #4768](https://redmine.postgresql.org/issues/4768) -  Ensure pgAdmin should work behind reverse proxy on a non standard port.<br>
[Issue #4769](https://redmine.postgresql.org/issues/4769) -  Fix query tool open issue on Internet Explorer.<br>
[Issue #4777](https://redmine.postgresql.org/issues/4777) -  Fix issue where query history is not visible in the query history tab.<br>
[Issue #4780](https://redmine.postgresql.org/issues/4780) -  Ensure the search path should not be quoted for Function, Procedure and Trigger Function.<br>
[Issue #4791](https://redmine.postgresql.org/issues/4791) -  Fix issue where VALID foreign keys show as NOT VALID in the SQL tab for tables.<br>
[Issue #4817](https://redmine.postgresql.org/issues/4817) -  Ensure the MAC OSX app should be notarized for Catalina.<br>
