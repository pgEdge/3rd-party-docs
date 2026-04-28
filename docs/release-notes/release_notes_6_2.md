# Version 6.2

Release date: 2021-11-18

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.1.

# New features

[Issue #3834](https://redmine.postgresql.org/issues/3834) -  Added support of Aggregate and Operator node in view-only mode.<br>
[Issue #6953](https://redmine.postgresql.org/issues/6953) -  Ensure that users should be able to modify the REMOTE_USER environment variable as per their environment by introducing the new config parameter WEBSERVER_REMOTE_USER.<br>

# Housekeeping

# Bug fixes

[Issue #5427](https://redmine.postgresql.org/issues/5427) -  Fixed pgAdmin freezing issue by providing the error message for the operation that can't perform due to lock on the particular table.<br>
[Issue #6780](https://redmine.postgresql.org/issues/6780) -  Ensure that columns should be merged if the newly added column is present in the parent table.<br>
[Issue #6809](https://redmine.postgresql.org/issues/6809) -  Fixed an issue where pgAdmin is not opening properly.<br>
[Issue #6832](https://redmine.postgresql.org/issues/6832) -  Ensure that internal authentication when combined with other authentication providers, the order of internal source should not matter while picking up the provider.<br>
[Issue #6845](https://redmine.postgresql.org/issues/6845) -  Ensure that inherit table icon should be visible properly in the tree view.<br>
[Issue #6859](https://redmine.postgresql.org/issues/6859) -  Fixed an issue where properties panel is not updated when any object is added from the browser tree.<br>
[Issue #6896](https://redmine.postgresql.org/issues/6896) -  Ensure that the user should be able to navigate browser tree objects using arrow keys from keyboard.<br>
[Issue #6905](https://redmine.postgresql.org/issues/6905) -  Fixed an issue where database nodes are not getting loaded behind a reverse proxy with SSL.<br>
[Issue #6925](https://redmine.postgresql.org/issues/6925) -  Fixed SQL syntax error if select "Custom auto-vacuum" option and not set Autovacuum option to Yes or No.<br>
[Issue #6939](https://redmine.postgresql.org/issues/6939) -  Fixed an issue where older server group name displayed in the confirmation pop-up when the user removes server group.<br>
[Issue #6944](https://redmine.postgresql.org/issues/6944) -  Fixed an issue where JSON editor preview colours have inappropriate contrast in dark mode.<br>
[Issue #6945](https://redmine.postgresql.org/issues/6945) -  Fixed JSON Editor scrolling issue in code mode.<br>
[Issue #6940](https://redmine.postgresql.org/issues/6940) -  Fixed an issue where user details are not shown when the non-admin user tries to connect to the shared server.<br>
[Issue #6949](https://redmine.postgresql.org/issues/6949) -  Ensure that dialog should be opened when clicking on Reassign/Drop owned menu.<br>
[Issue #6954](https://redmine.postgresql.org/issues/6954) -  Ensure that changing themes should work on Windows when system high contrast mode is enabled.<br>
[Issue #6972](https://redmine.postgresql.org/issues/6972) -  Ensure that the Binary path for PG14 should be visible in the preferences.<br>
[Issue #6974](https://redmine.postgresql.org/issues/6974) -  Added operators and aggregates in search objects.<br>
[Issue #6976](https://redmine.postgresql.org/issues/6976) -  Fixed an issue where textarea should be allowed to resize and have more than 255 chars.<br>
[Issue #6981](https://redmine.postgresql.org/issues/6981) -  Fixed an issue where SQL for index shows the same column multiple times.<br>
[Issue #6988](https://redmine.postgresql.org/issues/6988) -  Reset the layout if pgAdmin4 detects the layout is in an inconsistent state.<br>
