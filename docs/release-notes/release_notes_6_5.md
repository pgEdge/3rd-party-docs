# Version 6.5

Release date: 2022-02-11

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.4.

# New features

[Issue #7139](https://redmine.postgresql.org/issues/7139) -  Added support to open SQL help, Dialog help, and online help in an external web browser.<br>

# Housekeeping

[Issue #7016](https://redmine.postgresql.org/issues/7016) -  Port Dependent, dependencies, statistics panel to React.<br>
[Issue #7017](https://redmine.postgresql.org/issues/7017) -  Port Import/Export dialog to React.<br>
[Issue #7163](https://redmine.postgresql.org/issues/7163) -  Rename the menu 'Disconnect Server' to 'Disconnect from server'.<br>

# Bug fixes

[Issue #6916](https://redmine.postgresql.org/issues/6916) -  Added flag in runtime to disable GPU hardware acceleration.<br>
[Issue #7035](https://redmine.postgresql.org/issues/7035) -  Fixed an issue where connections keep open to (closed) connections on the initial connection to the database server.<br>
[Issue #7085](https://redmine.postgresql.org/issues/7085) -  Ensure that Partitioned tables should be visible correctly when creating multiple partition levels.<br>
[Issue #7086](https://redmine.postgresql.org/issues/7086) -  Correct documentation for 'Add named restore point'.<br>
[Issue #7100](https://redmine.postgresql.org/issues/7100) -  Fixed an issue where the Browser tree gets disappears when scrolling sequences.<br>
[Issue #7109](https://redmine.postgresql.org/issues/7109) -  Make the size blank for all the directories in the file select dialog.<br>
[Issue #7110](https://redmine.postgresql.org/issues/7110) -  Ensure that cursor should be focused on the first options of the Utility dialogs.<br>
[Issue #7118](https://redmine.postgresql.org/issues/7118) -  Ensure that JSON files should be downloaded properly from the storage manager.<br>
[Issue #7123](https://redmine.postgresql.org/issues/7123) -  Fixed an issue where restore generates incorrect options for the schema.<br>
[Issue #7126](https://redmine.postgresql.org/issues/7126) -  Fixed an issue where the F2 Function key removes browser panel contents.<br>
[Issue #7127](https://redmine.postgresql.org/issues/7127) -  Added validation for Hostname in the server dialog.<br>
[Issue #7135](https://redmine.postgresql.org/issues/7135) -  Enforce the minimum Windows version that the installer will run on.<br>
[Issue #7136](https://redmine.postgresql.org/issues/7136) -  Fixed an issue where the query tool is displaying an incorrect label.<br>
[Issue #7142](https://redmine.postgresql.org/issues/7142) -  Fixed an issue where a warning message was shown after database creation/modification.<br>
[Issue #7145](https://redmine.postgresql.org/issues/7145) -  Ensure that owner should be ignored while comparing extensions.<br>
[Issue #7146](https://redmine.postgresql.org/issues/7146) -  Fixed event trigger comparing issue in Schema Diff tool.<br>
[Issue #7150](https://redmine.postgresql.org/issues/7150) -  Fixed an issue when uploading a CSV throwing an error in the Desktop mode.<br>
[Issue #7151](https://redmine.postgresql.org/issues/7151) -  Fixed value error in the restore dialog.<br>
[Issue #7154](https://redmine.postgresql.org/issues/7154) -  Ensure that the layout should not be reset if a query tool is opened and pgAdmin is restarted.<br>
