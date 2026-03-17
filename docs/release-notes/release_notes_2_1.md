# Version 2.1

Release date: 2018-01-11

This release contains a number of features and fixes reported since the release of pgAdmin4 2.0

# Features

[Issue #1383](https://redmine.postgresql.org/issues/1383) - Allow connections to be coloured in the treeview and Query Tool<br>
[Issue #1489](https://redmine.postgresql.org/issues/1489) - Improve user interface for selection query in Data Filter window<br>
[Issue #2368](https://redmine.postgresql.org/issues/2368) - Improve data entry in Query Tool<br>
[Issue #2781](https://redmine.postgresql.org/issues/2781) - Allow configuration of CSV and clipboard formatting of query results<br>
[Issue #2802](https://redmine.postgresql.org/issues/2802) - Allow connections to be coloured in the treeview and Query Tool.<br>
[Issue #2810](https://redmine.postgresql.org/issues/2810) - Allow files to be opened by double clicking on them within Query Tool<br>
[Issue #2845](https://redmine.postgresql.org/issues/2845) - Make the "Save Changes" prompts in the Query Tool optional<br>
[Issue #2849](https://redmine.postgresql.org/issues/2849) - Add support for editing data in tables with OIDs but no primary keys and updates the editor to retrieve all row values on save, thus immediately showing default values and allowing subsequent editing without a refresh<br>

# Bug fixes

[Issue #1365](https://redmine.postgresql.org/issues/1365) - Prevent the Windows installer accepting paths containing invalid characters<br>
[Issue #1366](https://redmine.postgresql.org/issues/1366) - Fix /NOICONS switch in the windows installer<br>
[Issue #1436](https://redmine.postgresql.org/issues/1436) - Fix issue with debugger which is failing for sub - procedure on PPAS 9.6<br>
[Issue #1749](https://redmine.postgresql.org/issues/1749) - Fixes in pgAgent module including; 1) allowing start date earlier than end date when scheduling job, 2) Datetime picker not displaying in grid and 3) validation error not displaying propertly for Datetime control<br>
[Issue #2094](https://redmine.postgresql.org/issues/2094) - Display relevant error messages when access is denied creating a schema<br>
[Issue #2098](https://redmine.postgresql.org/issues/2098) - Cleanup some inconsistent error dialog titles<br>
[Issue #2258](https://redmine.postgresql.org/issues/2258) - Fix handling of DATERANGE[] type<br>
[Issue #2278](https://redmine.postgresql.org/issues/2278) - Display long names appropriately in dialogue headers<br>
[Issue #2443](https://redmine.postgresql.org/issues/2443) - Confirm with the user before exiting the runtime<br>
[Issue #2524](https://redmine.postgresql.org/issues/2524) - Fix debugging of self-referencing functions<br>
[Issue #2566](https://redmine.postgresql.org/issues/2566) - Fix the Pause/Resume Replay of WAL files for PostgreSQL 10<br>
[Issue #2624](https://redmine.postgresql.org/issues/2624) - Ensure the switch animation is consistent on the table dialogue and avoid displaying an error incorrectly<br>
[Issue #2651](https://redmine.postgresql.org/issues/2651) - Ensure estimated rows are included correctly in CREATE script for functions<br>
[Issue #2679](https://redmine.postgresql.org/issues/2679) - Getting started links does not open second time if User open any URL and Click on Close button with cross bar<br>
[Issue #2705](https://redmine.postgresql.org/issues/2705) - User can add expirty date on Windows<br>
[Issue #2715](https://redmine.postgresql.org/issues/2715) - Ensure we can download large files and keep the user informed about progress<br>
[Issue #2720](https://redmine.postgresql.org/issues/2720) - Ensure password changes are successful if authenticating using a pgpass file<br>
[Issue #2726](https://redmine.postgresql.org/issues/2726) - Ensure the auto-complete selection list can display longer names<br>
[Issue #2738](https://redmine.postgresql.org/issues/2738) - Ensure line numbers form CodeMirror don't appear on top of menus<br>
[Issue #2748](https://redmine.postgresql.org/issues/2748) - Format JSON/JSONB nicely when displaying it in the grid editor pop-up<br>
[Issue #2760](https://redmine.postgresql.org/issues/2760) - When selecting an SSL cert or key, update only the expected path in the UI, not all of them<br>
[Issue #2765](https://redmine.postgresql.org/issues/2765) - Do not decrypt the password when the password is 'None'.  This should avoid the common but harmless exception "ValueError: IV must be 16 bytes long while decrypting the password."<br>
[Issue #2768](https://redmine.postgresql.org/issues/2768) - Only allow specification of a pgpass file if libpq >= 10<br>
[Issue #2769](https://redmine.postgresql.org/issues/2769) - Correct keyboard shortcut. Don't un-comment code with alt+. in the Query Tool. It's only supposed to respond to ctrl/cmd+<br>
[Issue #2772](https://redmine.postgresql.org/issues/2772) - Remove external links from Panel's context menu<br>
[Issue #2778](https://redmine.postgresql.org/issues/2778) - Ensure the datatype cache is updated when a domain is added<br>
[Issue #2779](https://redmine.postgresql.org/issues/2779) - Ensure column collation isn't lost when changing field size<br>
[Issue #2780](https://redmine.postgresql.org/issues/2780) - Ensure auto-indent honours the spaces/tabs config setting<br>
[Issue #2782](https://redmine.postgresql.org/issues/2782) - Re-hash the way that we handle rendering of special types such as arrays<br>
[Issue #2787](https://redmine.postgresql.org/issues/2787) - Quote the owner name when creating types<br>
[Issue #2806](https://redmine.postgresql.org/issues/2806) - Attempt to decode database errors based on lc_messages<br>
[Issue #2811](https://redmine.postgresql.org/issues/2811) - Display process output as it happens<br>
[Issue #2820](https://redmine.postgresql.org/issues/2820) - Logs available when executing backup and restore<br>
[Issue #2821](https://redmine.postgresql.org/issues/2821) - Attempt to decode database errors based on lc_messages<br>
[Issue #2822](https://redmine.postgresql.org/issues/2822) - Re-hash the way that we handle rendering of special types such as arrays.<br>
[Issue #2824](https://redmine.postgresql.org/issues/2824) - Fix a number of graphical explain rendering issues<br>
[Issue #2836](https://redmine.postgresql.org/issues/2636) - Fix counted rows display in table properties<br>
[Issue #2842](https://redmine.postgresql.org/issues/2842) - Fix a number of graphical explain rendering issues<br>
[Issue #2846](https://redmine.postgresql.org/issues/2846) - Add an option to manually count rows in tables to render the properties<br>
[Issue #2854](https://redmine.postgresql.org/issues/2854) - Fix utility output capture encoding<br>
[Issue #2859](https://redmine.postgresql.org/issues/2859) - Allow form validation messages to be close in case the eclipse anything on the form<br>
[Issue #2866](https://redmine.postgresql.org/issues/2866) - Ensure we don't show the full path on the server when using virtual filesystem roots in server mode for SSL certs<br>
[Issue #2875](https://redmine.postgresql.org/issues/2875) - Ensure the scroll location is retains in the Query Tool data grid if the user changes tab and then returns<br>
[Issue #2877](https://redmine.postgresql.org/issues/2877) - Remove the artificial limit of 4000 characters from text areas<br>
[Issue #2880](https://redmine.postgresql.org/issues/2880) - Honour whitespace properly in the data grid<br>
[Issue #2881](https://redmine.postgresql.org/issues/2881) - Fix support for time without timezone<br>
[Issue #2886](https://redmine.postgresql.org/issues/2886) - Resolve issue where Insert failed when tried with default primary key value<br>
[Issue #2891](https://redmine.postgresql.org/issues/2891) - Allow changing of the users password without leaving the app<br>
[Issue #2892](https://redmine.postgresql.org/issues/2892) - Refuse password changes (and tell the user) if the notification email cannot be sent<br>
[Issue #2908](https://redmine.postgresql.org/issues/2908) - Fix bundle creation on Windows which was failing due to \r\n line endings in code mirror<br>
[Issue #2918](https://redmine.postgresql.org/issues/2918) - Add missing init.py to backports.csv when building the MSVC windows build<br>
[Issue #2920](https://redmine.postgresql.org/issues/2920) - Push HTTPD logs to container stdout/stderr as appropriate<br>
[Issue #2921](https://redmine.postgresql.org/issues/2921) - Fixes in pgAgent module including; 1) allowing start date earlier than end date when scheduling job, 2) Datetime picker not displaying in grid and 3) validation error not displaying propertly for Datetime control<br>
[Issue #2922](https://redmine.postgresql.org/issues/2922) - Don't login the user with every request in desktop mode. Just do it once<br>
[Issue #2923](https://redmine.postgresql.org/issues/2923) - Prevent the user pressing the select button in the file manager when it is supposed to be disabled<br>
[Issue #2924](https://redmine.postgresql.org/issues/2924) - Cleanup the layout of the filter data dialogue<br>
[Issue #2928](https://redmine.postgresql.org/issues/2928) - Prevent multiple connections to new slow-to-respond servers being initiated in error<br>
[Issue #2934](https://redmine.postgresql.org/issues/2934) - Fix a reference before assignment error in the file dialogue<br>
[Issue #2937](https://redmine.postgresql.org/issues/2937) - Prevent attempts to select directories as files in the file dialogue<br>
[Issue #2945](https://redmine.postgresql.org/issues/2945) - Ensure invalid options can't be selected on triggers on views<br>
[Issue #2949](https://redmine.postgresql.org/issues/2949) - Display complete SQL for FTS dictionaries<br>
[Issue #2952](https://redmine.postgresql.org/issues/2952) - Don't try to render security URLs in desktop mode<br>
[Issue #2954](https://redmine.postgresql.org/issues/2954) - Allow selection of validation error text<br>
[Issue #2974](https://redmine.postgresql.org/issues/2974) - Clear the messages tab when running EXPLAIN/EXPLAIN ANALYZE<br>
[Issue #2993](https://redmine.postgresql.org/issues/2993) - Fix view data for views/mat views<br>
