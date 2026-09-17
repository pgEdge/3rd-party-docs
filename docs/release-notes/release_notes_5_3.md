# Version 5.3

Release date: 2021-05-20

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.2.

# New features

[Issue #5954](https://redmine.postgresql.org/issues/5954) -  Added support to set auto width of columns by content size in the data output window.<br>
[Issue #6158](https://redmine.postgresql.org/issues/6158) -  Added support to connect PostgreSQL servers via Kerberos authentication.<br>
[Issue #6397](https://redmine.postgresql.org/issues/6397) -  Added "IF NOT EXISTS" clause while creating tables and partition tables which is convenient while using the ERD tool.<br>

# Housekeeping

# Bug fixes

[Issue #4436](https://redmine.postgresql.org/issues/4436) -  Fixed an issue where drag and drop object is not correct in codemirror for properties dialog.<br>
[Issue #5477](https://redmine.postgresql.org/issues/5477) -  Added support for cache bust webpack chunk files.<br>
[Issue #5555](https://redmine.postgresql.org/issues/5555) -  Fixed an issue where data is displayed in the wrong order when executing the query repeatedly.<br>
[Issue #5776](https://redmine.postgresql.org/issues/5776) -  Ensure that while connecting to the server using SSPI login, it should not prompt for the password.<br>
[Issue #6329](https://redmine.postgresql.org/issues/6329) -  Fixed an issue where the wrong SQL is showing for the child partition tables.<br>
[Issue #6341](https://redmine.postgresql.org/issues/6341) -  Fixed an issue where CSV download quotes the numeric columns.<br>
[Issue #6355](https://redmine.postgresql.org/issues/6355) -  Ensure that pgAdmin should not allow opening external files that are dragged into it.<br>
[Issue #6377](https://redmine.postgresql.org/issues/6377) -  Fixed an issue where schema diff does not create DROP DEFAULT statement for columns.<br>
[Issue #6385](https://redmine.postgresql.org/issues/6385) -  Ensure that Backup and Restore should work on shared servers.<br>
[Issue #6392](https://redmine.postgresql.org/issues/6392) -  Fixed an issue where the filter 'Include/Exclude By Selection' not working for null values.<br>
[Issue #6399](https://redmine.postgresql.org/issues/6399) -  Ensure that the user should not be able to add duplicate panels.<br>
[Issue #6407](https://redmine.postgresql.org/issues/6407) -  Added support for the creation of Nested Table and Varying Array Type for Advanced Server.<br>
[Issue #6408](https://redmine.postgresql.org/issues/6408) -  Fixed ModuleNotFoundError when running setup.py from outside of the root.<br>
[Issue #6409](https://redmine.postgresql.org/issues/6409) -  Fixed an issue where the current debug line is not visible in the 'Dark' theme.<br>
[Issue #6413](https://redmine.postgresql.org/issues/6413) -  Fixed an issue where duplicate columns are visible in the browser tree, which is owned by two sequences.<br>
[Issue #6414](https://redmine.postgresql.org/issues/6414) -  Fixed an issue where the Help message not displaying correctly on Login/Group role.<br>
[Issue #6416](https://redmine.postgresql.org/issues/6416) -  Added comment column in the properties panel for View and Materialized View collection node.<br>
[Issue #6417](https://redmine.postgresql.org/issues/6417) -  Fixed an issue where query editor is not being closed if the user clicks on the 'Don't Save' button.<br>
[Issue #6420](https://redmine.postgresql.org/issues/6420) -  Ensure that pgAdmin4 shut down completely on the Quit command.<br>
[Issue #6443](https://redmine.postgresql.org/issues/6443) -  Fixed an issue where file dialog showing incorrect files for the selected file types.<br>
[Issue #6444](https://redmine.postgresql.org/issues/6444) -  Fixed an issue where the user is not warned if Kerberos ticket expiration is less than 30 min while initiating a global backup.<br>
[Issue #6445](https://redmine.postgresql.org/issues/6445) -  Ensure that proper identification should be there when the server is connected using Kerberos or without Kerberos.<br>
