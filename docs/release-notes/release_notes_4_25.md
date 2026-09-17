# Version 4.25

Release date: 2020-08-20

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.24.

# New features

[Issue #3904](https://redmine.postgresql.org/issues/3904) -  Replace charting library Flotr2 with ChartJS using React.<br>
[Issue #5126](https://redmine.postgresql.org/issues/5126) -  Modified schema diff tool to compare two databases instead of two schemas.<br>
[Issue #5610](https://redmine.postgresql.org/issues/5610) -  Add a --yes command line option to setup-web.sh to allow non-interactive use.<br>

# Housekeeping

[Issue #5324](https://redmine.postgresql.org/issues/5324) -  Improve code coverage and API test cases for Foreign Servers and User Mappings.<br>
[Issue #5327](https://redmine.postgresql.org/issues/5327) -  Improve code coverage and API test cases for Schemas.<br>
[Issue #5336](https://redmine.postgresql.org/issues/5336) -  Improve code coverage and API test cases for Types.<br>
[Issue #5700](https://redmine.postgresql.org/issues/5700) -  Remove old Python 2 compatibility code.<br>
[Issue #5731](https://redmine.postgresql.org/issues/5731) -  Upgrade font awesome from v4 to v5.<br>

# Bug fixes

[Issue #3767](https://redmine.postgresql.org/issues/3767) -  Ensure that the original file format should be retained when saving the same file in SQL editor.<br>
[Issue #3791](https://redmine.postgresql.org/issues/3791) -  Added missing comments in reverse engineering SQL for each column of a View.<br>
[Issue #4123](https://redmine.postgresql.org/issues/4123) -  Fixed an issue where debugger doesn't work if the search path is set other than 'public'.<br>
[Issue #4361](https://redmine.postgresql.org/issues/4361) -  Fixed ssh tunnel hang issue when the user tries to disconnect the server.<br>
[Issue #4387](https://redmine.postgresql.org/issues/4387) -  Fixed an issue where the user is not able to insert the data if the table and columns name contains special characters.<br>
[Issue #4810](https://redmine.postgresql.org/issues/4810) -  Fixed an issue where the user is not able to save the new row if the table is empty.<br>
[Issue #5429](https://redmine.postgresql.org/issues/5429) -  Ensure that the Dictionaries drop-down shows all the dictionaries in the FTS configuration dialog.<br>
[Issue #5490](https://redmine.postgresql.org/issues/5490) -  Make the runtime configuration dialog non-modal.<br>
[Issue #5526](https://redmine.postgresql.org/issues/5526) -  Fixed an issue where copying and pasting a cell with multiple line data will result in multiple rows.<br>
[Issue #5567](https://redmine.postgresql.org/issues/5567) -  Fixed an issue where conversion of bytea to the binary string results in an error.<br>
[Issue #5604](https://redmine.postgresql.org/issues/5604) -  Fixed an issue where the entire logs is in red text when the user runs backup and restore.<br>
[Issue #5632](https://redmine.postgresql.org/issues/5632) -  Ensure that the user will be able to modify the start value of the Identity column.<br>
[Issue #5646](https://redmine.postgresql.org/issues/5646) -  Ensure that RLS Policy node should be searchable using search object.<br>
[Issue #5664](https://redmine.postgresql.org/issues/5664) -  Fixed an issue where 'ALTER VIEW' statement is missing when the user sets the default value of a column for View.<br>
[Issue #5670](https://redmine.postgresql.org/issues/5670) -  Fixed an issue where the error message does not have a close button on utility dialogs.<br>
[Issue #5689](https://redmine.postgresql.org/issues/5689) -  Added the 'ORDER BY' clause for the privileges type to fix schema diff issue.<br>
[Issue #5708](https://redmine.postgresql.org/issues/5708) -  Correct TLS certificate filename in the container deployment docs.<br>
[Issue #5710](https://redmine.postgresql.org/issues/5710) -  Fixed an issue when comparing the table with a trigger throwing error in schema diff.<br>
[Issue #5713](https://redmine.postgresql.org/issues/5713) -  Corrected DROP SQL syntax for catalog.<br>
[Issue #5716](https://redmine.postgresql.org/issues/5716) -  Fixed an issue where ajax call continues to fire even after disconnecting the database server.<br>
[Issue #5724](https://redmine.postgresql.org/issues/5724) -  Clarify some of the differences when running in server mode in the docs.<br>
[Issue #5730](https://redmine.postgresql.org/issues/5730) -  Resolve schema diff dependencies by selecting the appropriate node automatically and maintain the order in the generated script.<br>
