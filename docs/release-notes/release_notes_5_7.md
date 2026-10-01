# Version 5.7

Release date: 2021-09-09

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.6.

# New features

[Issue #2538](https://redmine.postgresql.org/issues/2538) -  Added support for the truncate table with restart identity.<br>
[Issue #4264](https://redmine.postgresql.org/issues/4264) -  Make code folding case insensitive in the code mirror.<br>
[Issue #4629](https://redmine.postgresql.org/issues/4629) -  Added database and server information on the Maintenance process watcher dialog.<br>
[Issue #6495](https://redmine.postgresql.org/issues/6495) -  Allow the referenced table to be the same as the local table in one to many relationship for ERD Tool.<br>
[Issue #6625](https://redmine.postgresql.org/issues/6625) -  Make closing tabs to be smarter by focusing on the appropriate tab when the user closed a tab.<br>
[Issue #6691](https://redmine.postgresql.org/issues/6691) -  Set PSQLRC and PSQL_HISTORY env vars to apt. user storage path in the server mode.<br>

# Housekeeping

# Bug fixes

[Issue #4567](https://redmine.postgresql.org/issues/4567) -  Fixed an issue where privileges were revoked using SQL query on objects like tables that do not correctly show in SQL tab.<br>
[Issue #4815](https://redmine.postgresql.org/issues/4815) -  Fixed an issue where the user can not paste the updated table header in safari 12 and 13 browsers.<br>
[Issue #5849](https://redmine.postgresql.org/issues/5849) -  Ensure that trigger function SQL should have 'create or replace function' instead of 'create function' only.<br>
[Issue #6419](https://redmine.postgresql.org/issues/6419) -  Fixed blank screen issue on windows and also made changes to use NWjs manifest for remembering window size.<br>
[Issue #6531](https://redmine.postgresql.org/issues/6531) -  Fixed the export image issue where relation lines are over the nodes.<br>
[Issue #6544](https://redmine.postgresql.org/issues/6544) -  Fixed width limitation issue in PSQL tool window.<br>
[Issue #6564](https://redmine.postgresql.org/issues/6564) -  Fixed an issue where columns with sequences get altered unnecessarily with a schema diff tool.<br>
[Issue #6570](https://redmine.postgresql.org/issues/6570) -  Ensure that the lock panel should not be blocked for larger records.<br>
[Issue #6572](https://redmine.postgresql.org/issues/6572) -  Partially fixes the data output panel display issue.<br>
[Issue #6620](https://redmine.postgresql.org/issues/6620) -  Fixed an issue where whitespace in function bodies was not applied while generating the script using Schema Diff.<br>
[Issue #6627](https://redmine.postgresql.org/issues/6627) -  Introduced OAUTH2_SCOPE variable for the Oauth2 scope configuration.<br>
[Issue #6641](https://redmine.postgresql.org/issues/6641) -  Enables pgAdmin to retrieve user permissions in case of nested roles which helps to terminate the session for AWS RDS.<br>
[Issue #6663](https://redmine.postgresql.org/issues/6663) -  Fixed no attribute '_asdict' error when connecting the database server.<br>
[Issue #6668](https://redmine.postgresql.org/issues/6668) -  Fixed errors related to HTML tags shown in the error message for JSON editor.<br>
[Issue #6671](https://redmine.postgresql.org/issues/6671) -  Fixed UnboundLocalError where local variable 'user_id' referenced before assignment.<br>
[Issue #6682](https://redmine.postgresql.org/issues/6682) -  Renamed 'Auto rollback?' to 'Auto rollback on error?'.<br>
[Issue #6684](https://redmine.postgresql.org/issues/6684) -  Fixed the JSON editor issue of hiding the first record.<br>
[Issue #6685](https://redmine.postgresql.org/issues/6685) -  Ensure that deleting a database should not automatically connect to the next database.<br>
[Issue #6704](https://redmine.postgresql.org/issues/6704) -  Ensure that pgAdmin should not fail at login due to a special character in the hostname.<br>
[Issue #6710](https://redmine.postgresql.org/issues/6710) -  Fixed an issue where multiple query tool tabs getting closed for the single close event.<br>
