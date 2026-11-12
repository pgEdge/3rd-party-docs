# Version 2.0

Release date: 2017-10-05

This release contains a number of features and fixes reported since the release of pgAdmin4 1.6

# Features

[Issue #1918](https://redmine.postgresql.org/issues/1918) - Add a field to the Server Dialogue allowing users to specify a subset of databases they'd like to see in the treeview<br>
[Issue #2135](https://redmine.postgresql.org/issues/2135) - Significantly speed up loading of the application<br>
[Issue #2556](https://redmine.postgresql.org/issues/2556) - Allow for slow vs. fast connection failures<br>
[Issue #2579](https://redmine.postgresql.org/issues/2579) - Default the file browser view to list, and make it configurable<br>
[Issue #2597](https://redmine.postgresql.org/issues/2597) - Allow queries to be cancelled from the dashboard and display additional info in the subnode control<br>
[Issue #2649](https://redmine.postgresql.org/issues/2649) - Support use of SSL certificates for authentication<br>
[Issue #2650](https://redmine.postgresql.org/issues/2650) - Support use of pgpass files<br>
[Issue #2662](https://redmine.postgresql.org/issues/2662) - Ship with pre-configured paths that can work in both Server and Desktop modes out of the box<br>
[Issue #2689](https://redmine.postgresql.org/issues/2689) - Update icons with new designs and remove from menus to de-clutter the UI<br>

# Bug fixes

[Issue #1165](https://redmine.postgresql.org/issues/1165) - Prevent continual polling for graph data on the dashboard if the server is disconnected<br>
[Issue #1697](https://redmine.postgresql.org/issues/1697) - Update CodeMirror version<br>
[Issue #2043](https://redmine.postgresql.org/issues/2043) - Properly handle trigger functions with parameters<br>
[Issue #2074](https://redmine.postgresql.org/issues/2074) - Make $ quoting consistent<br>
[Issue #2080](https://redmine.postgresql.org/issues/2080) - Fix issue where Browser hangs/crashes when loading data (using sql editor) from table which contains large blob data<br>
[Issue #2153](https://redmine.postgresql.org/issues/2153) - Fix handline of large file uploads and properly show any errors that may occur<br>
[Issue #2168](https://redmine.postgresql.org/issues/2168) - Update CodeMirror version<br>
[Issue #2170](https://redmine.postgresql.org/issues/2170) - Support SSL in the regression tests<br>
[Issue #2324](https://redmine.postgresql.org/issues/2324) - Fix PostGIS Datatypes in SQL tab, Create / Update dialogues for Table, Column, Foreign Table and Type node<br>
[Issue #2447](https://redmine.postgresql.org/issues/2447) - Update CodeMirror version<br>
[Issue #2452](https://redmine.postgresql.org/issues/2452) - Install pgadmin4-v1 1.5 on Centos7<br>
[Issue #2501](https://redmine.postgresql.org/issues/2501) - Fix collation tests on Windows, replace use of default 'POSIX' collation with 'C' collation for testing<br>
[Issue #2541](https://redmine.postgresql.org/issues/2541) - Fix issues using special keys on MacOS<br>
[Issue #2544](https://redmine.postgresql.org/issues/2544) - Correct malformed query generated when using custom type<br>
[Issue #2551](https://redmine.postgresql.org/issues/2551) - Show tablespace on partitions<br>
[Issue #2555](https://redmine.postgresql.org/issues/2555) - Fix issue in Query Tool where messages were not displaying from functions/procedures properly<br>
[Issue #2557](https://redmine.postgresql.org/issues/2557) - Tidy up tab styling<br>
[Issue #2558](https://redmine.postgresql.org/issues/2558) - Prevent the tab bar being hidden when detached tabs are being closed<br>
[Issue #2559](https://redmine.postgresql.org/issues/2559) - Stop tool buttons from changing their styling unexpectedly<br>
[Issue #2560](https://redmine.postgresql.org/issues/2560) - Fix View 'CREATE Script' Problem<br>
[Issue #2562](https://redmine.postgresql.org/issues/2562) - Update CodeMirror version<br>
[Issue #2563](https://redmine.postgresql.org/issues/2563) - Fix paths under non-standard virtual directories<br>
[Issue #2566](https://redmine.postgresql.org/issues/2566) - Fix Pause/Resume Replay of WAL files for PostgreSQL 10<br>
[Issue #2567](https://redmine.postgresql.org/issues/2567) - Use the proper database connection to fetch the default priviledges in the properties tab of the database<br>
[Issue #2582](https://redmine.postgresql.org/issues/2582) - Unset compression ratio if it is an empty string in Backup module<br>
[Issue #2586](https://redmine.postgresql.org/issues/2586) - Cleanup feature tests<br>
[Issue #2590](https://redmine.postgresql.org/issues/2590) - Allow navigation of query history using the arrow keys<br>
[Issue #2592](https://redmine.postgresql.org/issues/2592) - Stop Flask from initialising service twice in Debug mode<br>
[Issue #2593](https://redmine.postgresql.org/issues/2593) - Ensure babel-polyfill is loaded in older qWebKits<br>
[Issue #2594](https://redmine.postgresql.org/issues/2594) - Fix disconnection of new databases<br>
[Issue #2596](https://redmine.postgresql.org/issues/2596) - Define the proper NODE_ENV environment during running the webpack<br>
[Issue #2606](https://redmine.postgresql.org/issues/2606) - Ensure role names are escaped in the membership control<br>
[Issue #2616](https://redmine.postgresql.org/issues/2616) - Domain create dialog do not open and Font size issue in Security label control<br>
[Issue #2617](https://redmine.postgresql.org/issues/2617) - Add missing pgagent file in webpack.config.js<br>
[Issue #2619](https://redmine.postgresql.org/issues/2619) - Fix quoting of index column names on tables<br>
[Issue #2620](https://redmine.postgresql.org/issues/2620) - Set database name to blank('') when job type is set to batch, while creating pgAgent job<br>
[Issue #2631](https://redmine.postgresql.org/issues/2631) - Change mapping of cell from 'numeric' to 'integer' for integer control as numeric cell has been removed from the code<br>
[Issue #2633](https://redmine.postgresql.org/issues/2633) - Fix pgAgent job step issues<br>
[Issue #2634](https://redmine.postgresql.org/issues/2634) - Add New Server through Quick links<br>
[Issue #2637](https://redmine.postgresql.org/issues/2637) - Fix Copy so it still works after query results have been copied<br>
[Issue #2641](https://redmine.postgresql.org/issues/2641) - User management issues - styling and inability to edit users properly<br>
[Issue #2644](https://redmine.postgresql.org/issues/2644) - Fix alertify notification messages where checkmark box disconnected from frame<br>
[Issue #2646](https://redmine.postgresql.org/issues/2646) - Fix the path reference of load-node.gif which was referencing to vendor directory<br>
[Issue #2654](https://redmine.postgresql.org/issues/2654) - Update datetime picker<br>
[Issue #2655](https://redmine.postgresql.org/issues/2655) - Fix connection string validation for pgAgent jobs<br>
[Issue #2656](https://redmine.postgresql.org/issues/2656) - Change Datetimepicker to expand from bottom in pgAgent so calendar does not get hidden<br>
[Issue #2657](https://redmine.postgresql.org/issues/2657) - Fix syntax error while saving changes for start/end time, weekdays, monthdays, month, hours, minutes while updating the pgAgent Job<br>
[Issue #2659](https://redmine.postgresql.org/issues/2659) - Fix issue where unable to add/update variables for columns of a table<br>
[Issue #2660](https://redmine.postgresql.org/issues/2660) - Not able to select rows in History Tab<br>
[Issue #2668](https://redmine.postgresql.org/issues/2668) - Fix RE-SQL for triggers with a single arg<br>
[Issue #2670](https://redmine.postgresql.org/issues/2670) - Improve datamodel validations for default Validator if user (developer) does not implement validate function in datamodel<br>
[Issue #2671](https://redmine.postgresql.org/issues/2671) - Fix array data type formating for bigint, real, float, double precision<br>
[Issue #2681](https://redmine.postgresql.org/issues/2681) - Reset Query Tool options before running tests<br>
[Issue #2684](https://redmine.postgresql.org/issues/2684) - Fix layout of password prompt dialogue<br>
[Issue #2691](https://redmine.postgresql.org/issues/2691) - View data option is missing from pgAdmin4 2.0 version<br>
[Issue #2692](https://redmine.postgresql.org/issues/2692) - Base type is missing for Domain on pgAdmin4<br>
[Issue #2693](https://redmine.postgresql.org/issues/2693) - User list is not available on User mapping pgAdmin4<br>
[Issue #2698](https://redmine.postgresql.org/issues/2698) - User can not create function due to missing return type<br>
[Issue #2699](https://redmine.postgresql.org/issues/2699) - Filtered Rows issue on pgAdmin4<br>
[Issue #2700](https://redmine.postgresql.org/issues/2700) - Cancel button is visible after query executed succesfully<br>
[Issue #2707](https://redmine.postgresql.org/issues/2707) - Disable trigger button does not work on pgAdmin4<br>
[Issue #2708](https://redmine.postgresql.org/issues/2708) - Tablespace name should displayed instead of %s(new_tablespace)s with Move Objects to another tablespace<br>
[Issue #2709](https://redmine.postgresql.org/issues/2709) - Display user relations in schema prefixed by 'pg'<br>
[Issue #2713](https://redmine.postgresql.org/issues/2713) - Fix an exception seen sometimes when the server is restarted<br>
[Issue #2742](https://redmine.postgresql.org/issues/2742) - Ensure using an alternate role to connect to a database doesn't cause an error when checking recovery state.<br>
