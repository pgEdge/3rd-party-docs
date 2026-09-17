# Version 1.6

Release date: 2017-07-13

This release contains a number of features and fixes reported since the release of pgAdmin4 1.5

# Features

[Issue #1344](https://redmine.postgresql.org/issues/1344) - Allow the Query Tool, Debugger and web browser tabs to be moved to different monitors as desired<br>
[Issue #1533](https://redmine.postgresql.org/issues/1533) - Set focus on the first enabled field when a dialogue is opened<br>
[Issue #1535](https://redmine.postgresql.org/issues/1535) - Teach dialogues about Escape to cancel, Enter to Save/OK, and F1 for help<br>
[Issue #1971](https://redmine.postgresql.org/issues/1971) - Retain column sizing in the Query Tool results grid when the same query is re-run multiple times in a row<br>
[Issue #1972](https://redmine.postgresql.org/issues/1972) - Prompt the user to save dirty queries rather than discard them for a more natural workflow<br>
[Issue #2137](https://redmine.postgresql.org/issues/2137) - On-demand loading for the Query Tool results<br>
[Issue #2191](https://redmine.postgresql.org/issues/2191) - Add support for the hostaddr connection parameter. This helps us play nicely with Kerberos/SSPI and friends<br>
[Issue #2282](https://redmine.postgresql.org/issues/2282) - Overhaul the query history tab to allow browsing of the history and full query text<br>
[Issue #2379](https://redmine.postgresql.org/issues/2379) - Support inserting multiple new rows into a table without clicking Save for each row<br>
[Issue #2485](https://redmine.postgresql.org/issues/2485) - Add a shortcut to reset the zoom level in the runtime<br>
[Issue #2506](https://redmine.postgresql.org/issues/2506) - Allow the user to close the dashboard panel<br>
[Issue #2513](https://redmine.postgresql.org/issues/2513) - Add preferences to enable brace matching and brace closing in the SQL editors<br>

# Bug fixes

[Issue #1126](https://redmine.postgresql.org/issues/1126) - Various FTS dictionary cleanups<br>
[Issue #1229](https://redmine.postgresql.org/issues/1229) - Fix default values and SQL formatting for event triggers<br>
[Issue #1466](https://redmine.postgresql.org/issues/1466) - Prevent attempts to debug procedures with variadic arguments<br>
[Issue #1525](https://redmine.postgresql.org/issues/1525) - Make $ quoting consistent<br>
[Issue #1575](https://redmine.postgresql.org/issues/1575) - Properly display security labels on EPAS 9.2+<br>
[Issue #1795](https://redmine.postgresql.org/issues/1795) - Fix validation for external and range types<br>
[Issue #1813](https://redmine.postgresql.org/issues/1813) - List packages in PPAS 9.2-9.4 when creating synonyms<br>
[Issue #1831](https://redmine.postgresql.org/issues/1831) - Fix server stats display for EPAS 9.2, where inet needs casting to text for concatenation<br>
[Issue #1851](https://redmine.postgresql.org/issues/1851) - Reverse engineer SQL for table-returning functions correctly<br>
[Issue #1860](https://redmine.postgresql.org/issues/1860) - Ensure default values are honoured when adding/editing columns<br>
[Issue #1888](https://redmine.postgresql.org/issues/1888) - Fix various issues with pgAgent job steps and schedules<br>
[Issue #1889](https://redmine.postgresql.org/issues/1889) - Fix various issues with pgAgent job steps and schedules<br>
[Issue #1890](https://redmine.postgresql.org/issues/1890) - Fix various issues with pgAgent job steps and schedules<br>
[Issue #1920](https://redmine.postgresql.org/issues/1920) - Ensure saved passwords are effective immediately, not just following a restart when first saved<br>
[Issue #1928](https://redmine.postgresql.org/issues/1928) - Fix the handling of double precision[] type<br>
[Issue #1934](https://redmine.postgresql.org/issues/1934) - Fix import/export to work as expected with TSV data<br>
[Issue #1999](https://redmine.postgresql.org/issues/1999) - Handle warning correctly when saving query results to an unmounted USB drive<br>
[Issue #2013](https://redmine.postgresql.org/issues/2013) - Increase the default size of the Grant Wizard to enable it to properly display privileges at the default size on smaller displays<br>
[Issue #2014](https://redmine.postgresql.org/issues/2014) - To fix unexpected behaviour displayed if user stops debugging on package/procedure fire_emp<br>
[Issue #2043](https://redmine.postgresql.org/issues/2043) - Properly handle trigger functions with parameters<br>
[Issue #2078](https://redmine.postgresql.org/issues/2078) - Refresh the SQL editor view on resize to ensure the contents are re-rendered for the new viewport<br>
[Issue #2086](https://redmine.postgresql.org/issues/2086) - Allow editing of the WITH ADMIN option of role membership<br>
[Issue #2113](https://redmine.postgresql.org/issues/2113) - Correct the validation logic when modifying indexes/exclusion constraints<br>
[Issue #2116](https://redmine.postgresql.org/issues/2116) - Enable dialogue help buttons on Language and Foreign Table dialogues<br>
[Issue #2142](https://redmine.postgresql.org/issues/2142) - Fix canceling of Grant Wizard on Windows<br>
[Issue #2155](https://redmine.postgresql.org/issues/2155) - Fix removal of sizes from column definitions<br>
[Issue #2162](https://redmine.postgresql.org/issues/2162) - Allow non-superusers to debug their own functions and prevent them from setting global breakpoints<br>
[Issue #2242](https://redmine.postgresql.org/issues/2242) - Fix an issue in NodeAjaxControl caching with cache-node field and add cache-node field in Trigger & Event trigger node so that whenever the user creates new Trigger Function we get new data from server in NodeAjaxControl<br>
[Issue #2280](https://redmine.postgresql.org/issues/2280) - Handle procedure flags (IMMUTABLE STRICT SECURITY DEFINER PARALLEL RESTRICTED) properly in RE-SQL on EPAS<br>
[Issue #2324](https://redmine.postgresql.org/issues/2324) - Fix the PostGIS Datatypes in SQL tab, Create / Update dialogues for Table, Column, Foreign Table and Type node<br>
[Issue #2344](https://redmine.postgresql.org/issues/2344) - Fix issue with ctrl-c / ctrl-v not working in Query Tool<br>
[Issue #2348](https://redmine.postgresql.org/issues/2348) - Fix issue when resizing columns in Query Too/View Data where all row/colums will select/deselect<br>
[Issue #2355](https://redmine.postgresql.org/issues/2355) - Properly refresh the parent node when renaming children<br>
[Issue #2357](https://redmine.postgresql.org/issues/2355) - Cache statistics more reliably<br>
[Issue #2381](https://redmine.postgresql.org/issues/2381) - Fix the RE-SQL for for views to properly qualify trigger function names<br>
[Issue #2386](https://redmine.postgresql.org/issues/2386) - Display and allow toggling of trigger enable/disable status from the trigger dialogue<br>
[Issue #2398](https://redmine.postgresql.org/issues/2398) - Bypass the proxy server for local addresses on Windows<br>
[Issue #2400](https://redmine.postgresql.org/issues/2400) - Cleanup handling of default/null values when data editing<br>
[Issue #2414](https://redmine.postgresql.org/issues/2414) - Improve error handling in cases where the user tries to rename or create a server group that would duplicate an existing group<br>
[Issue #2417](https://redmine.postgresql.org/issues/2417) - Order columns in multi-column pkeys correctly<br>
[Issue #2422](https://redmine.postgresql.org/issues/2422) - Fix RE-SQL for rules which got the table name wrong in the header and DROP statement<br>
[Issue #2425](https://redmine.postgresql.org/issues/2425) - Handle composite primary keys correctly when deleting rows in the Edit Grid<br>
[Issue #2426](https://redmine.postgresql.org/issues/2426) - Allow creation of ENUM types with no members<br>
[Issue #2427](https://redmine.postgresql.org/issues/2427) - Add numerous missing checks to ensure objects really exist when we think they do<br>
[Issue #2435](https://redmine.postgresql.org/issues/2435) - Pass the database ID to the Query Tool when using the Script options<br>
[Issue #2436](https://redmine.postgresql.org/issues/2436) - Ensure the last placeholder is included when generating UPDATE scripts for tables<br>
[Issue #2448](https://redmine.postgresql.org/issues/2448) - Ensure that boolean checkboxes cycle values in the correct order<br>
[Issue #2450](https://redmine.postgresql.org/issues/2450) - Fix error on the stats tab with PG10. Also, rename the 10.0_plus template directory to 10_plus to match the new versioning<br>
[Issue #2461](https://redmine.postgresql.org/issues/2461) - Allow users to remove default values from columns properly<br>
[Issue #2468](https://redmine.postgresql.org/issues/2468) - Fix issue where function create script won't compile<br>
[Issue #2470](https://redmine.postgresql.org/issues/2470) - Fix an intermittent error seen during result polling<br>
[Issue #2476](https://redmine.postgresql.org/issues/2476) - Improvements to the Query Results grid including improvements to the UI and allow copy/paste from sets of rows, columns or arbitrary blocks of cells<br>
[Issue #2477](https://redmine.postgresql.org/issues/2477) - Ensure text editors render in an appropriate place on the results grid<br>
[Issue #2479](https://redmine.postgresql.org/issues/2479) - No need for the menu icon to link to the homepage, as pgAdmin is a SPA<br>
[Issue #2482](https://redmine.postgresql.org/issues/2482) - Use a more sensible name for Query Tool tabs<br>
[Issue #2486](https://redmine.postgresql.org/issues/2486) - Ensure the feature tests use the correct test settings database<br>
[Issue #2487](https://redmine.postgresql.org/issues/2487) - Maintain a client-side cache of preference values, populated using an async call<br>
[Issue #2489](https://redmine.postgresql.org/issues/2489) - Fix clipboard handling with large datasets<br>
[Issue #2492](https://redmine.postgresql.org/issues/2492) - Ensure the initial password is properly hashed during setup in web mode<br>
[Issue #2498](https://redmine.postgresql.org/issues/2498) - Properly handle bytea[], and 'infinity'::real/real[]<br>
[Issue #2502](https://redmine.postgresql.org/issues/2502) - Properly handle bytea[], and 'infinity'::real/real[]<br>
[Issue #2503](https://redmine.postgresql.org/issues/2503) - Handle missing/dropped synonyms gracefully<br>
[Issue #2504](https://redmine.postgresql.org/issues/2504) - Update MatView and pgAgent modules to work with recent integer/numeric changes<br>
[Issue #2507](https://redmine.postgresql.org/issues/2507) - Ensure revoked public privileges are displayed in the RE-SQL for functions<br>
[Issue #2518](https://redmine.postgresql.org/issues/2518) - Fix encoding issue when saving servers<br>
[Issue #2522](https://redmine.postgresql.org/issues/2522) - Improve speed of Select All in the results grid<br>
[Issue #2527](https://redmine.postgresql.org/issues/2527) - Fix deletion of table rows with the column definition having NOT NULL TRUE and HAS NO DEFAULT VALUE<br>
[Issue #2528](https://redmine.postgresql.org/issues/2528) - Allow breakpoints to be set on triggers on views<br>
[Issue #2529](https://redmine.postgresql.org/issues/2529) - Resolve a number of issues with domains and domain constraints<br>
[Issue #2532](https://redmine.postgresql.org/issues/2532) - Refresh nodes correctly when there is a single child that is updated<br>
[Issue #2534](https://redmine.postgresql.org/issues/2534) - Fix handling of CREATE TABLE OF <type><br>
[Issue #2535](https://redmine.postgresql.org/issues/2535) - Fix clear history functionality<br>
[Issue #2540](https://redmine.postgresql.org/issues/2540) - Ensure the save password option is enabled when creating a server<br>
