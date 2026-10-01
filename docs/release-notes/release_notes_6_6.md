# Version 6.6

Release date: 2022-03-10

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.5.

# New features

[Issue #7177](https://redmine.postgresql.org/issues/7177) -  Added capability to deploy PostgreSQL servers on Amazon RDS.<br>

# Housekeeping

[Issue #7180](https://redmine.postgresql.org/issues/7180) -  Rename the menu 'Disconnect Database' to 'Disconnect from database'.<br>

# Bug fixes

[Issue #6956](https://redmine.postgresql.org/issues/6956) -  Fixed a schema diff issue in which user mappings were not compared correctly.<br>
[Issue #6991](https://redmine.postgresql.org/issues/6991) -  Fixed an issue where pgadmin cannot connect to LDAP when STARTTLS is required before bind.<br>
[Issue #6999](https://redmine.postgresql.org/issues/6999) -  Fixed an issue where a warning is flashed every time for an email address when authentication sources are internal and ldap.<br>
[Issue #7105](https://redmine.postgresql.org/issues/7105) -  Fixed an issue where the parent partition table was not displayed during autocomplete.<br>
[Issue #7124](https://redmine.postgresql.org/issues/7124) -  Fixed the schema diff issue where tables have different column positions and a column has a default value.<br>
[Issue #7152](https://redmine.postgresql.org/issues/7152) -  Added comments column for the functions collection node.<br>
[Issue #7172](https://redmine.postgresql.org/issues/7172) -  Allow users to scroll and enter input when there is a validation error.<br>
[Issue #7173](https://redmine.postgresql.org/issues/7173) -  Fixed an issue where the User Management dialog is not opening.<br>
[Issue #7181](https://redmine.postgresql.org/issues/7181) -  Ensure that the user should be able to add new server with unix socket connection.<br>
[Issue #7186](https://redmine.postgresql.org/issues/7186) -  Fixes an issue where the connect server/database menu was not updated correctly.<br>
[Issue #7202](https://redmine.postgresql.org/issues/7202) -  Ensure that Flask-Security-Too is using the latest version.<br>
