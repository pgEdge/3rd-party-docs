# Version 6.13

Release date: 2022-08-25

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.12.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13 and 14

**EDB Advanced Server**: 10, 11, 12, 13 and 14

# New features

[Issue #3709](https://redmine.postgresql.org/issues/3709) -  Added support to show all background processes in separate panel.<br>
[Issue #7387](https://redmine.postgresql.org/issues/7387) -  Added support to create triggers from existing trigger functions in EPAS.<br>

# Housekeeping

[Issue #7344](https://redmine.postgresql.org/issues/7344) -  Port Role Reassign dialog to React.<br>
[Issue #7345](https://redmine.postgresql.org/issues/7345) -  Port User Management dialog to React.<br>
[Issue #7404](https://redmine.postgresql.org/issues/7404) -  Port process watcher to React.<br>
[Issue #7462](https://redmine.postgresql.org/issues/7462) -  Remove the SQL files for the unsupported versions of the database server.<br>
[Issue #7567](https://redmine.postgresql.org/issues/7567) -  Port About dialog to React.<br>
[Issue #7568](https://redmine.postgresql.org/issues/7568) -  Port change user password and 2FA dialog to React.<br>
[Issue #7590](https://redmine.postgresql.org/issues/7590) -  Port change ownership dialog to React.<br>
[Issue #7595](https://redmine.postgresql.org/issues/7595) -  Update the container base image to Alpine 3.16 (with Python 3.10.5).<br>
[Issue #7602](https://redmine.postgresql.org/issues/7602) -  Fixed improper parsing of HTTP requests in Pallets Werkzeug v2.1.0 and below (CVE-2022-29361).<br>

# Bug fixes

[Issue #7452](https://redmine.postgresql.org/issues/7452) -  Ensure that an error is thrown if clipboard access is not provided and change the copy rows shortcut.<br>
[Issue #7468](https://redmine.postgresql.org/issues/7468) -  Fixed an issue where the History tab is getting blank and showing an error after some queries are executed.<br>
[Issue #7481](https://redmine.postgresql.org/issues/7481) -  Fixed an issue where OWNED BY was incorrectly set to NONE when adding user privileges on the sequence.<br>
[Issue #7497](https://redmine.postgresql.org/issues/7497) -  Fixed an issue with the error message being displayed at the right place for Azure deployments.<br>
[Issue #7521](https://redmine.postgresql.org/issues/7521) -  Fixed an issue where the Query Editor loses focus when saving a query (Alt+s).<br>
[Issue #7527](https://redmine.postgresql.org/issues/7527) -  Fixed API test cases for Postgres 14.4.<br>
[Issue #7540](https://redmine.postgresql.org/issues/7540) -  Ensure that rename panel should work on view/edit panels.<br>
[Issue #7563](https://redmine.postgresql.org/issues/7563) -  Fixed an issue where autocomplete is not working after clearing the query editor.<br>
[Issue #7573](https://redmine.postgresql.org/issues/7573) -  Ensure that autocomplete does not appear when navigating code using arrow keys.<br>
[Issue #7575](https://redmine.postgresql.org/issues/7575) -  Fixed an issue where Alt-Shift-Q didn't work after creating a new query.<br>
[Issue #7579](https://redmine.postgresql.org/issues/7579) -  Fixed an issue where copy and pasting a row in the results grid doesn't set the default for boolean.<br>
[Issue #7586](https://redmine.postgresql.org/issues/7586) -  Fixed an issue with rendering geometry when selecting a complete column.<br>
[Issue #7587](https://redmine.postgresql.org/issues/7587) -  Ensure that the children of information_schema and pg_catalog node should be displayed.<br>
[Issue #7591](https://redmine.postgresql.org/issues/7591) -  Fixed column "none" does not exist issue, while comparing schema objects.<br>
[Issue #7596](https://redmine.postgresql.org/issues/7596) -  Fixed an issue where schema diff did not pick up the change in RLS policy.<br>
[Issue #7608](https://redmine.postgresql.org/issues/7608) -  Fixed an issue where the cloud deployment wizard creates the cluster with the High Availability even if that option is not selected.<br>
[Issue #7611](https://redmine.postgresql.org/issues/7611) -  Ensure that schema diff maintains view ownership when view definitions are modified.<br>
[Issue #7614](https://redmine.postgresql.org/issues/7614) -  Fixed crypt key is missing issue when logout from the pgAdmin.<br>
[Issue #7616](https://redmine.postgresql.org/issues/7616) -  Ensure that the next button should be disabled if the password did not match for Azure deployment.<br>
[Issue #7617](https://redmine.postgresql.org/issues/7617) -  Fixed an issue where Azure cloud deployment failed.<br>
[Issue #7625](https://redmine.postgresql.org/issues/7625) -  Fixed Spanish translations typo.<br>
[Issue #7630](https://redmine.postgresql.org/issues/7630) -  Ensure that If the trigger function definition is changed, drop and recreate the trigger in the schema diff.<br>
[Issue #7632](https://redmine.postgresql.org/issues/7632) -  Fixed an issue where a user could not authenticate using Azure CLI on OSX.<br>
[Issue #7633](https://redmine.postgresql.org/issues/7633) -  Ensure that the autofocus is on the input control for the master password and server password dialogs.<br>
[Issue #7641](https://redmine.postgresql.org/issues/7641) -  Pin Flask-SocketIO <= v5.2.0. The latest version does not support Werkzeug in production environments.<br>
