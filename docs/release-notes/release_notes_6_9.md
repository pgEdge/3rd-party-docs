# Version 6.9

Release date: 2022-05-12

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.8.

# New features

[Issue #3253](https://redmine.postgresql.org/issues/3253) -  Added status bar to the Query Tool.<br>
[Issue #3989](https://redmine.postgresql.org/issues/3989) -  Ensure that row numbers should be visible in view when scrolling horizontally.<br>
[Issue #6830](https://redmine.postgresql.org/issues/6830) -  Relocate GIS Viewer Button to the Left Side of the Results Table.<br>
[Issue #7179](https://redmine.postgresql.org/issues/7179) -  Added capability to deploy PostgreSQL servers on EDB BigAnimal.<br>
[Issue #7282](https://redmine.postgresql.org/issues/7282) -  Added options 'Ignore owner' and 'Ignore whitespace' to the schema diff panel.<br>
[Issue #7325](https://redmine.postgresql.org/issues/7325) -  Added support for Azure AD OAUTH2 authentication.<br>

# Housekeeping

[Issue #6131](https://redmine.postgresql.org/issues/6131) -  Port query tool to React.<br>
[Issue #6746](https://redmine.postgresql.org/issues/6746) -  Improve the Kerberos Documentation.<br>
[Issue #7255](https://redmine.postgresql.org/issues/7255) -  Ensure the database and schema restriction controls are not shown as a drop-down.<br>
[Issue #7340](https://redmine.postgresql.org/issues/7340) -  Port data filter dialog to React.<br>

# Bug fixes

[Issue #6725](https://redmine.postgresql.org/issues/6725) -  Fixed an issue where the Query tool opens on minimum size if the user opens multiple query tool Window quickly.<br>
[Issue #6958](https://redmine.postgresql.org/issues/6958) -  Only set permissions on the storage directory upon creation.<br>
[Issue #7026](https://redmine.postgresql.org/issues/7026) -  Fixed an issue where the Browser panel is not completely viewable.<br>
[Issue #7168](https://redmine.postgresql.org/issues/7168) -  Improvement to the Geometry Viewer popup to change the size of the result tables when column names are quite long.<br>
[Issue #7187](https://redmine.postgresql.org/issues/7187) -  Fixed an issue where the downloaded ERD diagram was 0 bytes.<br>
[Issue #7188](https://redmine.postgresql.org/issues/7188) -  Fixed an issue where the connection bar is not visible.<br>
[Issue #7231](https://redmine.postgresql.org/issues/7231) -  Don't strip binaries when packaging them in the server RPM as this might break cpython modules.<br>
[Issue #7252](https://redmine.postgresql.org/issues/7252) -  Ensure that Columns should always be visible in the import/export dialog.<br>
[Issue #7260](https://redmine.postgresql.org/issues/7260) -  Fixed an issue where an Empty message popup after running a query.<br>
[Issue #7262](https://redmine.postgresql.org/issues/7262) -  Ensure that Autocomplete should work after changing the connection.<br>
[Issue #7294](https://redmine.postgresql.org/issues/7294) -  Fixed an issue where the copy and paste row does not work if the first column contains no data.<br>
[Issue #7296](https://redmine.postgresql.org/issues/7296) -  Ensure that after deleting multiple objects from the properties panel, the browser tree should be refreshed.<br>
[Issue #7299](https://redmine.postgresql.org/issues/7299) -  Fixed sorting issue in the statistics panel.<br>
[Issue #7305](https://redmine.postgresql.org/issues/7305) -  Fixed an issue where the Dashboard Server Activity was showing old queries as active.<br>
[Issue #7307](https://redmine.postgresql.org/issues/7307) -  Fixed an issue where the table showed duplicate columns when creating multiple sequences on the same column.<br>
[Issue #7308](https://redmine.postgresql.org/issues/7308) -  Ensure that sorting should be preserved on refresh for Server Activity.<br>
[Issue #7322](https://redmine.postgresql.org/issues/7322) -  Fixed an issue while creating a new database throwing an error that failed to retrieve data.<br>
[Issue #7333](https://redmine.postgresql.org/issues/7333) -  Fixed an issue where the drag and drop table in ERD throws an error.<br>
[Issue #7339](https://redmine.postgresql.org/issues/7339) -  Ensure that the Dashboard column sort order should be remembered when the refresh button is clicked.<br>
