# Version 5.2

Release date: 2021-04-22

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.1.

# New features

# Housekeeping

[Issue #5319](https://redmine.postgresql.org/issues/5319) -  Improve code coverage and API test cases for Server module.<br>

# Bug fixes

[Issue #4001](https://redmine.postgresql.org/issues/4001) -  Updated docs and screenshots to cover the Notifications tab on the Query Tool.<br>
[Issue #5519](https://redmine.postgresql.org/issues/5519) -  Ensure that the query tool tab should be closed after server disconnection when auto-commit/auto-rollback is set to false.<br>
[Issue #5908](https://redmine.postgresql.org/issues/5908) -  Fixed an issue where shortcut keys are not working with manage macro.<br>
[Issue #6076](https://redmine.postgresql.org/issues/6076) -  Fixed an issue where correct error not thrown while importing servers and JSON file has incorrect/insufficient keys.<br>
[Issue #6082](https://redmine.postgresql.org/issues/6082) -  Ensure that the user should not be to change the connection when a long query is running.<br>
[Issue #6107](https://redmine.postgresql.org/issues/6107) -  Fixed flickering issue of the input box on check constraints.<br>
[Issue #6161](https://redmine.postgresql.org/issues/6161) -  Fixed an issue where the cursor shifts its focus to the wrong window for all the query tool related model dialogs.<br>
[Issue #6220](https://redmine.postgresql.org/issues/6220) -  Corrected the syntax for 'CREATE TRIGGER', use 'EXECUTE FUNCTION' instead of 'EXECUTE PROCEDURE' from v11 onwards.<br>
[Issue #6274](https://redmine.postgresql.org/issues/6274) -  Ensure that the strings in the LDAP auth module are translatable.<br>
[Issue #6293](https://redmine.postgresql.org/issues/6293) -  Fixed an issue where the procedure creation is failed when providing the Volatility option.<br>
[Issue #6306](https://redmine.postgresql.org/issues/6306) -  Fixed an issue while selecting the row which was deleted just before the selection operation.<br>
[Issue #6325](https://redmine.postgresql.org/issues/6325) -  Ensure that the file format for the storage manager should be 'All files' and for other dialogs, it should remember the last selected format.<br>
[Issue #6327](https://redmine.postgresql.org/issues/6327) -  Ensure that while comparing domains check function dependencies should be considered in schema diff.<br>
[Issue #6333](https://redmine.postgresql.org/issues/6333) -  Fixed sizing issue of help dialog for Query Tool and ERD Tool when open in the new browser tab.<br>
[Issue #6334](https://redmine.postgresql.org/issues/6334) -  Fixed SQL panel black screen issue when detaching it in runtime.<br>
[Issue #6338](https://redmine.postgresql.org/issues/6338) -  Added missing dependency 'xdg-utils' for the desktop packages in RPM and Debian.<br>
[Issue #6344](https://redmine.postgresql.org/issues/6344) -  Fixed cannot unpack non-iterable response object error when selecting any partition.<br>
[Issue #6356](https://redmine.postgresql.org/issues/6356) -  Mark the Apache HTTPD config file as such in the web DEB and RPM packages.<br>
[Issue #6367](https://redmine.postgresql.org/issues/6367) -  Fixed an issue where the Save button is enabled by default when open the table's properties dialog on PG 9.5.<br>
[Issue #6375](https://redmine.postgresql.org/issues/6375) -  Fixed an issue where users are unable to see data of the partitions using the View/Edit data option.<br>
[Issue #6376](https://redmine.postgresql.org/issues/6376) -  Fixed an issue where a connection warning should be displayed on the user clicks on explain or explain analyze and the database server is disconnected from the browser tree.<br>
[Issue #6379](https://redmine.postgresql.org/issues/6379) -  Fixed an issue where foreign data wrapper properties are not visible if the host option contains two host addresses.<br>
