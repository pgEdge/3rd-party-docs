# Version 6.14

Release date: 2022-09-22

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.13.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13 and 14

**EDB Advanced Server**: 10, 11, 12, 13 and 14

# New features

# Housekeeping

[Issue #4059](https://github.com/pgadmin-org/pgadmin4/issues/4059) -  Port schema diff to React. (RM #6133)<br>
[Issue #4060](https://github.com/pgadmin-org/pgadmin4/issues/4060) -  Remove Backgrid and Backform. (RM #6134)<br>
[Issue #5035](https://github.com/pgadmin-org/pgadmin4/issues/5035) -  Port the remaining components of the ERD Tool to React. (RM #7343)<br>
[Issue #5120](https://github.com/pgadmin-org/pgadmin4/issues/5120) -  Updated keyboard shortcut documentation. (RM #7446)<br>
[Issue #5260](https://github.com/pgadmin-org/pgadmin4/issues/5260) -  Remove Alertify from pgAdmin completely. (RM #7619)<br>
[Issue #5263](https://github.com/pgadmin-org/pgadmin4/issues/5263) -  Port search object dialog to React. (RM #7622)<br>

# Bug fixes

[Issue #5144](https://github.com/pgadmin-org/pgadmin4/issues/5144) -  Ensure that if BigAnimal authentication is aborted, API calls should be stopped. (RM #7472)<br>
[Issue #5209](https://github.com/pgadmin-org/pgadmin4/issues/5209) -  Fixed an issue where pgAdmin failed to start due to bin path migration. (RM #7557)<br>
[Issue #5230](https://github.com/pgadmin-org/pgadmin4/issues/5230) -  Fixed an issue where backup does not work due to parameter 'preexec_fn' no longer being supported. (RM #7580)<br>
[Issue #5250](https://github.com/pgadmin-org/pgadmin4/issues/5250) -  Ensure that the browser tree should be refreshed after changing the ownership. (RM #7607)<br>
[Issue #5274](https://github.com/pgadmin-org/pgadmin4/issues/5274) -  Fixed the error message displayed when clicking the cloud server for which deployment is in progress. (RM #7636)<br>
[Issue #5275](https://github.com/pgadmin-org/pgadmin4/issues/5275) -  Fixed an issue where the wrong SQL displayed in difference if the user create an RLS policy on the table without a column. (RM #7637)<br>
[Issue #5282](https://github.com/pgadmin-org/pgadmin4/issues/5282) -  Ensure that the dump servers functionality works from setup.py. (RM #7644)<br>
[Issue #5284](https://github.com/pgadmin-org/pgadmin4/issues/5284) -  Ensure that the Import/Export server menu option is visible. (RM #7646)<br>
[Issue #5286](https://github.com/pgadmin-org/pgadmin4/issues/5286) -  Fixed API test case for change password in the server mode. (RM #7648)<br>
[Issue #5287](https://github.com/pgadmin-org/pgadmin4/issues/5287) -  Fixed an issue with the non-visibility of columns added prior to import/export data. (RM #7649)<br>
[Issue #5292](https://github.com/pgadmin-org/pgadmin4/issues/5292) -  Fixed an issue where textarea of the JSON Editor does not resize with dialog. (RM #7656)<br>
[Issue #5299](https://github.com/pgadmin-org/pgadmin4/issues/5299) -  Fixed ModuleNotFoundError when running setup.py to load/dump servers. (RM #7663)<br>
[Issue #5323](https://github.com/pgadmin-org/pgadmin4/issues/5323) -  Replace the language selection 'Brazilian' with 'Portuguese (Brazilian). (RM #7693)<br>
[Issue #5325](https://github.com/pgadmin-org/pgadmin4/issues/5325) -  Fixed an issue where server names with special characters are not displayed correctly in the process tab. (RM #7695)<br>
[Issue #5333](https://github.com/pgadmin-org/pgadmin4/issues/5333) -  Fixed an issue where ERD throws an error if variable is added to the column. (RM #7709)<br>
[Issue #5342](https://github.com/pgadmin-org/pgadmin4/issues/5342) -  Fixed an error while saving changes to the ERD table.<br>
[Issue #5343](https://github.com/pgadmin-org/pgadmin4/issues/5343) -  Fixes a redirect vulnerability when the user opens the pgAdmin URL.<br>
