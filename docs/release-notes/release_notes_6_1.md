# Version 6.1

Release date: 2021-10-21

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.0.

# New features

[Issue #4596](https://redmine.postgresql.org/issues/4596) -  Added support for indent guides in the browser tree.<br>
[Issue #6081](https://redmine.postgresql.org/issues/6081) -  Added support for advanced table fields like the foreign key, primary key in the ERD tool.<br>
[Issue #6241](https://redmine.postgresql.org/issues/6241) -  Added support to allow tables to be dragged to ERD Tool.<br>
[Issue #6529](https://redmine.postgresql.org/issues/6529) -  Added index creation when generating SQL in the ERD tool.<br>
[Issue #6657](https://redmine.postgresql.org/issues/6657) -  Added support for authentication via the webserver (REMOTE_USER).<br>
[Issue #6794](https://redmine.postgresql.org/issues/6794) -  Added support to enable/disable rules.<br>

# Housekeeping

# Bug fixes

[Issue #6719](https://redmine.postgresql.org/issues/6719) -  Fixed OAuth2 integration redirect issue.<br>
[Issue #6754](https://redmine.postgresql.org/issues/6754) -  Ensure that query highlighting color in the query tool should be less intensive.<br>
[Issue #6776](https://redmine.postgresql.org/issues/6776) -  Changed the label 'Inherits Tables?' to 'Is inherited?' as it misleading in the properties panel.<br>
[Issue #6790](https://redmine.postgresql.org/issues/6790) -  Fixed an issue where the user is unable to create an index with concurrently keyword.<br>
[Issue #6797](https://redmine.postgresql.org/issues/6797) -  Remove an extra blank line at the start of the SQL for function, procedure, and trigger function.<br>
[Issue #6802](https://redmine.postgresql.org/issues/6802) -  Fixed the issue of editing triggers for advanced servers.<br>
[Issue #6828](https://redmine.postgresql.org/issues/6828) -  Fixed an issue where the tree is not scrolling to the object selected from the search result.<br>
[Issue #6858](https://redmine.postgresql.org/issues/6858) -  Fixed object delete issue from the properties tab for the collection nodes.<br>
[Issue #6876](https://redmine.postgresql.org/issues/6876) -  Ensure that the Dashboard should get updated after connecting to the server.<br>
[Issue #6881](https://redmine.postgresql.org/issues/6881) -  Fixed an issue where the browser tree doesn't show all contents on changing resolution.<br>
[Issue #6882](https://redmine.postgresql.org/issues/6882) -  Ensure that columns should be displayed in the order of creation instead of alphabetical order in the browser tree.<br>
[Issue #6890](https://redmine.postgresql.org/issues/6890) -  Fixed background colour issue in the browser tree.<br>
[Issue #6891](https://redmine.postgresql.org/issues/6891) -  Added support for composite foreign keys in the ERD tool.<br>
[Issue #6900](https://redmine.postgresql.org/issues/6900) -  Fixed an issue where exclusion constraint cannot be created from table dialog if the access method name is changed once.<br>
[Issue #6905](https://redmine.postgresql.org/issues/6905) -  Fixed an issue where the users are unable to load the databases behind an HTTP reverse proxy.<br>
[Issue #6908](https://redmine.postgresql.org/issues/6908) -  Fixed an issue where each click to refresh the collection node, the number of objects decreasing by tens or more.<br>
[Issue #6912](https://redmine.postgresql.org/issues/6912) -  Fixed browser tree sort order regression issue.<br>
[Issue #6915](https://redmine.postgresql.org/issues/6915) -  Fixed an issue where the blank string is stored instead of NULL in the server table of SQLite database.<br>
[Issue #6928](https://redmine.postgresql.org/issues/6928) -  Ensure that the master password should be prompt when MASTER_PASSWORD_REQUIRED is set to True and AUTHENTICATION_SOURCES is webserver.<br>
[Issue #6929](https://redmine.postgresql.org/issues/6929) -  Ensure that only the table node should be allowed to drop on the ERD tool.<br>
[Issue #6930](https://redmine.postgresql.org/issues/6930) -  Fixed an issue where the existing server group is disappeared on rename it.<br>
[Issue #6935](https://redmine.postgresql.org/issues/6935) -  Fixed an issue where the wrong SQL is generated when deleting and renaming table columns together.<br>
