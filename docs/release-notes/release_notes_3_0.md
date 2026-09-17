# Version 3.0

Release date: 2018-03-22

This release contains a number of features and fixes reported since the release of pgAdmin4 2.1

# Features

[Issue #1894](https://redmine.postgresql.org/issues/1894) - Allow sorting when viewing/editing data<br>
[Issue #1978](https://redmine.postgresql.org/issues/1978) - Add the ability to enable/disable UI animations<br>
[Issue #2895](https://redmine.postgresql.org/issues/2895) - Add keyboard navigation options for the main browser windows<br>
[Issue #2896](https://redmine.postgresql.org/issues/2896) - Add keyboard navigation in Query tool module via Tab/Shift-Tab key<br>
[Issue #2897](https://redmine.postgresql.org/issues/2897) - Support keyboard navigation in the debugger<br>
[Issue #2898](https://redmine.postgresql.org/issues/2898) - Support tab navigation in dialogs<br>
[Issue #2899](https://redmine.postgresql.org/issues/2899) - Add configurable shortcut keys for various common options in the main window<br>
[Issue #2901](https://redmine.postgresql.org/issues/2901) - Configurable shortcuts in the Debugger<br>
[Issue #2904](https://redmine.postgresql.org/issues/2904) - Ensure clickable images/buttons have appropriate tooltips for screen readers<br>
[Issue #2950](https://redmine.postgresql.org/issues/2950) - Add a marker (/*pga4dash*/) to the dashboard queries to allow them to be more easily filtered from server logs<br>
[Issue #2951](https://redmine.postgresql.org/issues/2951) - Allow dashboard tables and charts to be enabled/disabled<br>
[Issue #3004](https://redmine.postgresql.org/issues/3004) - Support server and database statistics on Greenplum<br>
[Issue #3036](https://redmine.postgresql.org/issues/3036) - Display partitions in Greenplum<br>
[Issue #3044](https://redmine.postgresql.org/issues/3044) - Display functions in Greenplum<br>
[Issue #3086](https://redmine.postgresql.org/issues/3086) - Rewrite the runtime as a tray-based server which can launch a web browser<br>
[Issue #3097](https://redmine.postgresql.org/issues/3097) - Support EXPLAIN on Greenplum<br>
[Issue #3098](https://redmine.postgresql.org/issues/3098) - Unvendorize REACT so no longer required in our source tree<br>
[Issue #3107](https://redmine.postgresql.org/issues/3107) - Hide tablespace node on GPDB<br>
[Issue #3140](https://redmine.postgresql.org/issues/3140) - Add support for connecting using pg_service.conf files<br>
[Issue #3168](https://redmine.postgresql.org/issues/3168) - Support for external tables in GPDB<br>
[Issue #3182](https://redmine.postgresql.org/issues/3182) - Update Jasmine to v3<br>
[Issue #3184](https://redmine.postgresql.org/issues/3184) - Add a French translation<br>
[Issue #3195](https://redmine.postgresql.org/issues/3195) - Pass the service name to external processes<br>
[Issue #3246](https://redmine.postgresql.org/issues/3246) - Update container build to use Alpine Linux and Gunicorn instead of CentOS/Apache<br>

`In addition, various changes were made for PEP8 compliance`<br>

# Bug fixes

[Issue #1173](https://redmine.postgresql.org/issues/1173) - Add a comment to the existing node<br>
[Issue #1925](https://redmine.postgresql.org/issues/1925) - Fix issue resizing column widths not resizable in Query Tool after first query<br>
[Issue #2104](https://redmine.postgresql.org/issues/2104) - Runtime update display file version and copyright year under installers properties<br>
[Issue #2249](https://redmine.postgresql.org/issues/2249) - Application no longer hangs after reload in runtime<br>
[Issue #2251](https://redmine.postgresql.org/issues/2251) - Runtime fixed OSX html scroll direction ignored in MacOS setup<br>
[Issue #2309](https://redmine.postgresql.org/issues/2309) - Allow text selection/copying from disabled CodeMirror instances<br>
[Issue #2480](https://redmine.postgresql.org/issues/2480) - Runtime update fix to Context Menus on Mac that do not work<br>
[Issue #2578](https://redmine.postgresql.org/issues/2578) - Runtime update fix to HTML access keys that don't work<br>
[Issue #2581](https://redmine.postgresql.org/issues/2581) - Fix keyboard shortcut for text selection<br>
[Issue #2677](https://redmine.postgresql.org/issues/2677) - Update Elephant icon for pgAdmin4 on Windows<br>
[Issue #2776](https://redmine.postgresql.org/issues/2776) - Fix unreadable font via Remote Desktop<br>
[Issue #2777](https://redmine.postgresql.org/issues/2777) - Fix spacing issue on server tree<br>
[Issue #2783](https://redmine.postgresql.org/issues/2783) - Runtime update fixed blank screen on Windows Desktop<br>
[Issue #2906](https://redmine.postgresql.org/issues/2906) - Correct display issues on HiDPI screens<br>
[Issue #2961](https://redmine.postgresql.org/issues/2961) - Issues when creating a pgAgent Schedule<br>
[Issue #2963](https://redmine.postgresql.org/issues/2963) - Fix unicode handling in the external process tools and show the complete command in the process viewer<br>
[Issue #2980](https://redmine.postgresql.org/issues/2980) - Copy text from the Query tool into the clipboard adds invisible characters<br>
[Issue #2981](https://redmine.postgresql.org/issues/2981) - Support keyboard navigation in the debugger<br>
[Issue #2983](https://redmine.postgresql.org/issues/2983) - Fix intermittent specified_version_number ValueError issue on restart<br>
[Issue #2985](https://redmine.postgresql.org/issues/2985) - Fix drag and drop issues<br>
[Issue #2998](https://redmine.postgresql.org/issues/2998) - Don't listen on port 443 if TLS is not enabled when launching the container<br>
[Issue #3001](https://redmine.postgresql.org/issues/3001) - Runtime update fix scrolling with mouse wheel on mac pgAdmin 4.2.1<br>
[Issue #3002](https://redmine.postgresql.org/issues/3002) - Fix block indent/outdent with configurable width<br>
[Issue #3003](https://redmine.postgresql.org/issues/3003) - Runtime update fix copy to clipboard<br>
[Issue #3005](https://redmine.postgresql.org/issues/3005) - Runtime update fix unable to select tabs in pgAdmin 4.2.1<br>
[Issue #3013](https://redmine.postgresql.org/issues/3013) - Fix a minor UI issue on dashboard while displaying subnode control in Backgrid<br>
[Issue #3014](https://redmine.postgresql.org/issues/3014) - Fix validation of sequence parameters<br>
[Issue #3015](https://redmine.postgresql.org/issues/3015) - Support Properties on Greenplum databases<br>
[Issue #3016](https://redmine.postgresql.org/issues/3016) - Ensure debug messages are available in "messages" window when error occurs<br>
[Issue #3021](https://redmine.postgresql.org/issues/3021) - Update scan and index scan EXPLAIN icons for greater clarity<br>
[Issue #3027](https://redmine.postgresql.org/issues/3027) - Ensure we capture notices raised by queries<br>
[Issue #3031](https://redmine.postgresql.org/issues/3031) - Runtime issue causing double and single quotes not to work<br>
[Issue #3039](https://redmine.postgresql.org/issues/3039) - Runtime issue causing wrong row counts on count column<br>
[Issue #3042](https://redmine.postgresql.org/issues/3042) - Runtime issue causing empty dialog box when refreshing<br>
[Issue #3043](https://redmine.postgresql.org/issues/3043) - Runtime issue causing word sizing in macOS High Sierra<br>
[Issue #3045](https://redmine.postgresql.org/issues/3045) - Runtime issue causing copy cells issues copying cells for key binding<br>
[Issue #3046](https://redmine.postgresql.org/issues/3046) - Fix connection status indicator on IE/FF<br>
[Issue #3050](https://redmine.postgresql.org/issues/3050) - Correct display of RE-SQL for partitioned tables in Greenplum<br>
[Issue #3052](https://redmine.postgresql.org/issues/3052) - Don't include sizes on primitive data types that shouldn't have them when modifying columns<br>
[Issue #3054](https://redmine.postgresql.org/issues/3054) - Ensure the user can use keyboard shortcuts after using button controls such as Cancel, Open and Save<br>
[Issue #3057](https://redmine.postgresql.org/issues/3057) - Update the regression tests to fix issues with Python 3.5 and PG 9.2<br>
[Issue #3058](https://redmine.postgresql.org/issues/3058) - Fix on-click handling of treeview nodes that wasn't refreshing SQL/Dependencies/Dependents in some circumstances<br>
[Issue #3059](https://redmine.postgresql.org/issues/3059) - Fix table statistics for Greenplum<br>
[Issue #3060](https://redmine.postgresql.org/issues/3060) - Fix quoting of function names in RE-SQL<br>
[Issue #3066](https://redmine.postgresql.org/issues/3066) - Ensure column names on indexes on views are properly quoted in RE-SQL<br>
[Issue #3067](https://redmine.postgresql.org/issues/3067) - Prevent the filter dialog CodeMirror from overflowing onto the button bar of the dialog<br>
[Issue #3072](https://redmine.postgresql.org/issues/3072) - Add a (configurable) limit to the number of pgAgent job history rows displayed on the statistics tab<br>
[Issue #3073](https://redmine.postgresql.org/issues/3073) - Ensure the pgAgent job start/end time grid fields synchronise with the subnode control and validate correctly<br>
[Issue #3075](https://redmine.postgresql.org/issues/3075) - Runtime issue causing Select, Update, and Insert script generation for a table fails to load<br>
[Issue #3077](https://redmine.postgresql.org/issues/3077) - Remove dependency on standards_conforming_strings being enabled<br>
[Issue #3079](https://redmine.postgresql.org/issues/3079) - Fix handling of tie/datetime array types when adding columns to a table<br>
[Issue #3080](https://redmine.postgresql.org/issues/3080) - Fix alignment issues in keyboard shortcut options<br>
[Issue #3081](https://redmine.postgresql.org/issues/3081) - Add missing reverse-engineered SQL header and drop statement for sequences<br>
[Issue #3090](https://redmine.postgresql.org/issues/3090) - Ensure message severity is decoded when necessary by the driver<br>
[Issue #3094](https://redmine.postgresql.org/issues/3094) - Ensure all messages are retrieved from the server in the Query Tool<br>
[Issue #3099](https://redmine.postgresql.org/issues/3099) - Fix creation of tables and columns in GPDB<br>
[Issue #3105](https://redmine.postgresql.org/issues/3105) - Ensure we can properly update rows with upper-case primary key columns<br>
[Issue #3135](https://redmine.postgresql.org/issues/3135) - Insert rows correctly when a table has OIDs and a Primary Key in uppercase<br>
[Issue #3122](https://redmine.postgresql.org/issues/3122) - Ensure SSL options are pushed down to external tools like pg_dump<br>
[Issue #3129](https://redmine.postgresql.org/issues/3129) - Handle opening of non-UTF8 compatible files<br>
[Issue #3137](https://redmine.postgresql.org/issues/3137) - Allow copying of SQL from the dashboard tables<br>
[Issue #3138](https://redmine.postgresql.org/issues/3138) - Fix tablespace tests for Python 3.x<br>
[Issue #3150](https://redmine.postgresql.org/issues/3150) - Fix function reserve SQL for GPDB<br>
[Issue #3157](https://redmine.postgresql.org/issues/3157) - Fix unicode handling in the external process tools and show the complete command in the process viewer<br>
[Issue #3171](https://redmine.postgresql.org/issues/3171) - Runtime issue causing inability to scroll in File Selector with trackpad on OSX<br>
[Issue #3176](https://redmine.postgresql.org/issues/3176) - Disable function statistics on Greenplum<br>
[Issue #3180](https://redmine.postgresql.org/issues/3180) - Ensure Indexes are displayed on PG 10 tables<br>
[Issue #3190](https://redmine.postgresql.org/issues/3190) - Skip tests where appropriate on GPDB<br>
[Issue #3196](https://redmine.postgresql.org/issues/3196) - Ensure the file manager properly escapes file & directory names<br>
[Issue #3197](https://redmine.postgresql.org/issues/3197) - Appropriately set the cookie path<br>
[Issue #3200](https://redmine.postgresql.org/issues/3200) - Ensure the host parameter is correctly pickup up from the service file<br>
[Issue #3219](https://redmine.postgresql.org/issues/3219) - Update required ChromeDriver version for current versions of Chrome<br>
[Issue #3226](https://redmine.postgresql.org/issues/3226) - Move the field error indicators in front of the affected fields so they don't obscure spinners or drop downs etc.<br>
[Issue #3244](https://redmine.postgresql.org/issues/3244) - Show more granular timing info in the Query Tool history panel<br>
[Issue #3248](https://redmine.postgresql.org/issues/3248) - Ensure Alertify dialogues are modal to prevent them being closed by mis-click<br>
