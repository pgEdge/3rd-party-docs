# Version 4.4

Release date: 2019-04-04

This release contains a number of new features and fixes reported since the release of pgAdmin4 4.3.

!!! warning

    This release includes a bug fix ([Issue #3887](https://redmine.postgresql.org/issues/3887)) which will rename the per-user storage directories for existing users when running in server mode. Previously, saved SQL queries were stored under the *STORAGE_DIR* in a sub-directory named after the username part of the user's email address. From this version onwards, the full email address is used, with the @ replaced with an underscore. For example, in v.4.3 with *STORAGE_DIR* set to */var/lib/pgadmin4* user files may be stored in:

    ```bash
     /var/lib/pgadmin4/storage/username/
    ```

    With the fix, that directory will be renamed (or created for new users) as:

    ```bash
     /var/lib/pgadmin4/storage/username_example.com/
    ```

# Features

[Issue #2001](https://redmine.postgresql.org/issues/2001) - Add support for reverse proxied setups with Gunicorn, and document Gunicorn, uWSGI & NGINX configurations.<br>
[Issue #4017](https://redmine.postgresql.org/issues/4017) - Make the Query Tool history persistent across sessions.<br>
[Issue #4018](https://redmine.postgresql.org/issues/4018) - Remove the large and unnecessary dependency on React and 87 other related libraries.<br>
[Issue #4030](https://redmine.postgresql.org/issues/4030) - Add support for IDENTITY columns.<br>
[Issue #4075](https://redmine.postgresql.org/issues/4075) - Add an ePub doc build target.<br>

# Bug fixes

[Issue #1269](https://redmine.postgresql.org/issues/1269) - Fix naming inconsistency for the column and FTS parser modules.<br>
[Issue #2627](https://redmine.postgresql.org/issues/2627) - Include inherited column comments and defaults in reverse engineered table SQL.<br>
[Issue #3104](https://redmine.postgresql.org/issues/3104) - Improve a couple of German translations.<br>
[Issue #3887](https://redmine.postgresql.org/issues/3887) - Use the user's full email address (not just the username part) as the basis for the storage directory name.<br>
[Issue #3968](https://redmine.postgresql.org/issues/3968) - Update wcDocker to fix the issue where the Scratch Pad grows in size if the results panel is resized.<br>
[Issue #3995](https://redmine.postgresql.org/issues/3995) - Avoid 'bogus varno' message from Postgres when viewing the SQL for a table with triggers.<br>
[Issue #4019](https://redmine.postgresql.org/issues/4019) - Update all Python and JavaScript dependencies.<br>
[Issue #4037](https://redmine.postgresql.org/issues/4037) - Include comment SQL for inherited columns in reverse engineered table SQL.<br>
[Issue #4050](https://redmine.postgresql.org/issues/4050) - Make the WHEN field a CodeMirror control on the Event Trigger dialogue.<br>
[Issue #4052](https://redmine.postgresql.org/issues/4052) - Fix the online help button on the resource group dialogue.<br>
[Issue #4053](https://redmine.postgresql.org/issues/4053) - Enable the online help button on the index dialogue.<br>
[Issue #4054](https://redmine.postgresql.org/issues/4054) - Handle resultsets with zero columns correctly in the Query Tool.<br>
[Issue #4058](https://redmine.postgresql.org/issues/4058) - Include inherited columns in SELECT scripts.<br>
[Issue #4060](https://redmine.postgresql.org/issues/4060) - Fix the latexpdf doc build.<br>
[Issue #4062](https://redmine.postgresql.org/issues/4062) - Fix handling of numeric arrays in View/Edit Data.<br>
[Issue #4063](https://redmine.postgresql.org/issues/4063) - Enlarge the grab handles for resizing dialogs etc.<br>
[Issue #4069](https://redmine.postgresql.org/issues/4069) - Append the file suffix to filenames when needed in the File Create dialogue.<br>
[Issue #4071](https://redmine.postgresql.org/issues/4071) - Ensure that Firefox prompts for a filename/location when downloading query results as a CSV file.<br>
[Issue #4073](https://redmine.postgresql.org/issues/4073) - Change the CodeMirror active line background colour to $color-danger-lighter so it doesn't conflict with the selection colour.<br>
[Issue #4081](https://redmine.postgresql.org/issues/4081) - Fix the RE-SQL syntax for roles with a VALID UNTIL clause.<br>
[Issue #4082](https://redmine.postgresql.org/issues/4082) - Prevent an empty error message being shown when "downloading" a CREATE script using the CSV download.<br>
[Issue #4084](https://redmine.postgresql.org/issues/4084) - Overhaul the layout saving code so it includes the Query Tool and Debugger, and stores the layout when change events are detected rather than (unreliably) on exit.<br>
[Issue #4085](https://redmine.postgresql.org/issues/4085) - Display errors during CSV download from the Query Tool in the UI rather than putting them in the CSV file.<br>
[Issue #4090](https://redmine.postgresql.org/issues/4090) - Improve the German translation for Backup Server.<br>
[Issue #4096](https://redmine.postgresql.org/issues/4096) - Ensure the toolbar buttons are properly reset following a CSV download in the Query Tool.<br>
[Issue #4099](https://redmine.postgresql.org/issues/4099) - Fix SQL help for EPAS 10+, and refactor the URL generation code into a testable function.<br>
[Issue #4100](https://redmine.postgresql.org/issues/4100) - Ensure sequences can be created with increment, start, minimum and maximum options set.<br>
[Issue #4105](https://redmine.postgresql.org/issues/4105) - Fix an issue where JSON data would not be rendered in the Query Tool.<br>
[Issue #4109](https://redmine.postgresql.org/issues/4109) - Ensure View/Materialized View node should be visible after updating any property.<br>
[Issue #4110](https://redmine.postgresql.org/issues/4110) - Fix custom autovacuum configuration for Materialized Views.<br>
