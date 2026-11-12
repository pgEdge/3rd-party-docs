# Version 5.4

Release date: 2021-06-17

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.3.

# New features

[Issue #1561](https://redmine.postgresql.org/issues/1561) -  Added browse button to select the binary path in the Preferences.<br>
[Issue #1591](https://redmine.postgresql.org/issues/1591) -  Added Grant Wizard option under Package node.<br>
[Issue #2341](https://redmine.postgresql.org/issues/2341) -  Added support to launch PSQL for the connected database server.<br>
[Issue #4064](https://redmine.postgresql.org/issues/4064) -  Added window maximize/restore functionality for properties dialog.<br>
[Issue #5370](https://redmine.postgresql.org/issues/5370) -  Added support to set the binary path for the different database server versions.<br>
[Issue #6231](https://redmine.postgresql.org/issues/6231) -  Added OS, Browser, Configuration details in the About dialog.<br>
[Issue #6395](https://redmine.postgresql.org/issues/6395) -  Added support for rotating the pgAdmin log file on the basis of size and age.<br>
[Issue #6524](https://redmine.postgresql.org/issues/6524) -  Support non-admin installation on Windows.<br>

# Housekeeping

[Issue #4622](https://redmine.postgresql.org/issues/4622) -  Added RESQL/MSQL test cases for Table and its child nodes.<br>
[Issue #6225](https://redmine.postgresql.org/issues/6225) -  Updated Flask-Security-Too to the latest v4.<br>
[Issue #6460](https://redmine.postgresql.org/issues/6460) -  Added a mechanism to detect a corrupt/broken config database file.<br>

# Bug fixes

[Issue #4203](https://redmine.postgresql.org/issues/4203) -  Fixed the issue of renaming the database by another user.<br>
[Issue #6404](https://redmine.postgresql.org/issues/6404) -  Ensure that the Query Tool connection string should not be changed as per the 'Query Tool tab title'.<br>
[Issue #6466](https://redmine.postgresql.org/issues/6466) -  Ensure that the user should be able to add members in Login/Role group while creating it.<br>
[Issue #6469](https://redmine.postgresql.org/issues/6469) -  Ensure that the calendar control should be disabled in the properties panel for Role.<br>
[Issue #6473](https://redmine.postgresql.org/issues/6473) -  Disable browser password saving in the runtime.<br>
[Issue #6478](https://redmine.postgresql.org/issues/6478) -  Fixed duplicate SQL issue for tables with more than one partition.<br>
[Issue #6482](https://redmine.postgresql.org/issues/6482) -  Fixed an issue where the Foreground Color property of server dialog does not work.<br>
[Issue #6513](https://redmine.postgresql.org/issues/6513) -  Fixed an issue where pgAdmin does not open after password reset in server mode.<br>
[Issue #6520](https://redmine.postgresql.org/issues/6520) -  Fixed an issue where a decimal number is appended for character varying fields while downloading the data in CSV format.<br>
