# Version 6.16

Release date: 2022-11-18

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.15.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13, 14 and 15

**EDB Advanced Server**: 10, 11, 12, 13, 14 and 15

# New features

[Issue #1832](https://github.com/pgadmin-org/pgadmin4/issues/1832) -  Added support for storing configurations of pgAdmin in an external database.<br>
[Issue #4756](https://github.com/pgadmin-org/pgadmin4/issues/4756) -  Added the ability to generate ERDs for tables.<br>
[Issue #5468](https://github.com/pgadmin-org/pgadmin4/issues/5468) -  Add the possibility to configure the Oauth2 claim which is used for the pgAdmin username.<br>

# Housekeeping

# Bug fixes

[Issue #2174](https://github.com/pgadmin-org/pgadmin4/issues/2174) -  Ensure that the browser tree should auto scroll to the selected node when expanding the server node.<br>
[Issue #4841](https://github.com/pgadmin-org/pgadmin4/issues/4841) -  Use SocketIO instead of REST for schema diff compare.<br>
[Issue #5066](https://github.com/pgadmin-org/pgadmin4/issues/5066) -  Ensure that users can use custom characters as CSV field separators/CSV quotes when downloading query results.<br>
[Issue #5058](https://github.com/pgadmin-org/pgadmin4/issues/5058) -  Ensure that the save button should be disabled by default on the Sort/Filter dialog in the query tool.<br>
[Issue #5098](https://github.com/pgadmin-org/pgadmin4/issues/5098) -  Fix an issue where the save button is enabled when the table properties dialog is opened.<br>
[Issue #5122](https://github.com/pgadmin-org/pgadmin4/issues/5122) -  Ensure that the spinner should be visible on the browser tree on node refresh.<br>
[Issue #5149](https://github.com/pgadmin-org/pgadmin4/issues/5149) -  Ensure the Generate ERD option is hidden if the connection to the database is not allowed.<br>
[Issue #5206](https://github.com/pgadmin-org/pgadmin4/issues/5206) -  Reposition the select dropdown when the browser is resized.<br>
[Issue #5281](https://github.com/pgadmin-org/pgadmin4/issues/5281) -  Ensure that autocomplete works properly with objects starting with double quotes.<br>
[Issue #5344](https://github.com/pgadmin-org/pgadmin4/issues/5344) -  Ensure that pgAdmin routes should have the SCRIPT_NAME prefix.<br>
[Issue #5424](https://github.com/pgadmin-org/pgadmin4/issues/5424) -  Ensure that the appropriate permissions are set on the key file before trying an SSL connection with the server in server mode.<br>
[Issue #5429](https://github.com/pgadmin-org/pgadmin4/issues/5429) -  Fixed an issue where parameters for roles were not visible.<br>
[Issue #5452](https://github.com/pgadmin-org/pgadmin4/issues/5452) -  The container deployment document should include the server.json file format.<br>
[Issue #5455](https://github.com/pgadmin-org/pgadmin4/issues/5455) -  Fixed an issue where the dependents tab wasn't working for PG 15.<br>
[Issue #5458](https://github.com/pgadmin-org/pgadmin4/issues/5458) -  Ensure that the browser path column in the search object shows the complete path.<br>
[Issue #5463](https://github.com/pgadmin-org/pgadmin4/issues/5463) -  Fixed an issue where the result grid was not working properly while trying to edit data by hitting Enter key.<br>
[Issue #5465](https://github.com/pgadmin-org/pgadmin4/issues/5465) -  Fixed an issue where the screen was freezing while closing the wcDocker panel.<br>
[Issue #5473](https://github.com/pgadmin-org/pgadmin4/issues/5473) -  Fixed an issue where AutoComplete was not working correctly due to incorrect regex.<br>
[Issue #5475](https://github.com/pgadmin-org/pgadmin4/issues/5475) -  Fixed an issue where the 'Confirm on close or refresh' setting was ignored when closing the query/ERD tool opened in the new tab.<br>
[Issue #5507](https://github.com/pgadmin-org/pgadmin4/issues/5507) -  Fixed an issue where pgadmin does not respect reverse proxy any more.<br>
[Issue #5521](https://github.com/pgadmin-org/pgadmin4/issues/5521) -  Fixed SocketIO calls when pgAdmin 4 server is running from a sub directory.<br>
[Issue #5522](https://github.com/pgadmin-org/pgadmin4/issues/5522) -  Ensure that the load file paths are children of the storage directory.<br>
[Issue #5533](https://github.com/pgadmin-org/pgadmin4/issues/5533) -  Use the shared server username when opening query tool.<br>
[Issue #5535](https://github.com/pgadmin-org/pgadmin4/issues/5535) -  Fixed an issue where the 'save_password' column threw an error for the shared server when using an external database.<br>
[Issue #5537](https://github.com/pgadmin-org/pgadmin4/issues/5537) -  Ensure that the correct error message in ERD for permission denied should be shown.<br>
