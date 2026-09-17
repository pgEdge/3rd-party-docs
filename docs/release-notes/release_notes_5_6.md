# Version 5.6

Release date: 2021-08-12

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.5.

# New features

[Issue #4904](https://redmine.postgresql.org/issues/4904) -  Added support to copy SQL from main window to query tool.<br>
[Issue #5198](https://redmine.postgresql.org/issues/5198) -  Added support for formatted JSON viewer/editor when interacting with data in a JSON column.<br>

# Housekeeping

[Issue #6622](https://redmine.postgresql.org/issues/6622) -  Rename the "Resize by data?" to "Columns sized by" and disabled the 'Maximum column width' button if 'Columns sized by' is set to 'Column data'.<br>

# Bug fixes

[Issue #6337](https://redmine.postgresql.org/issues/6337) -  Ensure that the login account should be locked after N number of attempts. N is configurable using the 'MAX_LOGIN_ATTEMPTS' parameter.<br>
[Issue #6369](https://redmine.postgresql.org/issues/6369) -  Fixed CSRF errors for stale sessions by increasing the session expiration time for desktop mode.<br>
[Issue #6448](https://redmine.postgresql.org/issues/6448) -  Fixed an issue in the search object when searching in 'all types' or 'subscription' if the user doesn't have access to the subscription.<br>
[Issue #6574](https://redmine.postgresql.org/issues/6574) -  Fixed an issue where paste is not working through Right-Click option on PSQL.<br>
[Issue #6580](https://redmine.postgresql.org/issues/6580) -  Fixed TypeError 'NoneType' object is not sub scriptable.<br>
[Issue #6586](https://redmine.postgresql.org/issues/6586) -  Fixed incorrect tablespace options in the drop-down for move objects dialog.<br>
[Issue #6618](https://redmine.postgresql.org/issues/6618) -  Fixed an issue where the titles in query tabs are different.<br>
[Issue #6619](https://redmine.postgresql.org/issues/6619) -  Fixed incorrect binary path issue when the user deletes the binary path from the preferences.<br>
[Issue #6643](https://redmine.postgresql.org/issues/6643) -  Ensure that all the required options should be loaded when the Range data type is selected while creating a custom data type.<br>
[Issue #6650](https://redmine.postgresql.org/issues/6650) -  Fixed dashboard server activity issue when active_since parameter is None.<br>
[Issue #6664](https://redmine.postgresql.org/issues/6664) -  Fixed an issue where even if the user is locked, he can reset the password and can login into pgAdmin.<br>
