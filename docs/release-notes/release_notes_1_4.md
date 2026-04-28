# Version 1.4

Release date: 2017-04-13

This release contains a number of features and fixes reported since the release of pgAdmin4 1.3.

# Features

[Issue #2232](https://redmine.postgresql.org/issues/2232) - Add the ability to gray-out/disable the "Save Password" option when creating a connection to a server<br>
[Issue #2261](https://redmine.postgresql.org/issues/2261) - Display table DDL for Greenplum in SQL tab<br>
[Issue #2320](https://redmine.postgresql.org/issues/2163) - Added German translation<br>

# Bug fixes

[Issue #2077](https://redmine.postgresql.org/issues/2077) - Add missing "Run Now" option for pgAdmin jobs<br>
[Issue #2105](https://redmine.postgresql.org/issues/2105) - Fix validation on the table dialogue so the Save button isn't enabled if the name is removed and autovac custom settings are enabled<br>
[Issue #2145](https://redmine.postgresql.org/issues/2145) - Resolve the issue for restoring the table from the backup<br>
[Issue #2187](https://redmine.postgresql.org/issues/2187) - Ensure the web/ directory is cleared before upgrading Windows installations<br>
[Issue #2190](https://redmine.postgresql.org/issues/2190) - Allow users to select UI language at login or from Preferences rather than unpredictable behaviour from browsers<br>
[Issue #2226](https://redmine.postgresql.org/issues/2226) - Show tooltips for disabled buttons to help user learning<br>
[Issue #2241](https://redmine.postgresql.org/issues/2241) - Fix numeric control validation in nested schemas<br>
[Issue #2243](https://redmine.postgresql.org/issues/2243) - Fix dropping of databases with Unicode names<br>
[Issue #2244](https://redmine.postgresql.org/issues/2244) - Prevent an error being displayed if the user views data on a table with no columns<br>
[Issue #2246](https://redmine.postgresql.org/issues/2246) - Add missing braces to reverse engineered SQL header block for Functions<br>
[Issue #2258](https://redmine.postgresql.org/issues/2258) - Fix handling of DATERANGE[] type<br>
[Issue #2264](https://redmine.postgresql.org/issues/2264) - Resolve error message *ExtDeprecationWarning* displayed on new pgAdmin4 setup for Python 3.4 on ubuntu 14.04 Linux 64<br>
[Issue #2265](https://redmine.postgresql.org/issues/2265) - Resolved import/Export issue for a table<br>
[Issue #2274](https://redmine.postgresql.org/issues/2274) - Properly handle truncated table names<br>
[Issue #2277](https://redmine.postgresql.org/issues/2277) - Resolved various file-system encoding/decoding related cases<br>
[Issue #2281](https://redmine.postgresql.org/issues/2281) - Ensure menus are updated after disconnecting a server<br>
[Issue #2283](https://redmine.postgresql.org/issues/2283) - Check if cell is in multiselect mode before setting default selection of multiple values<br>
[Issue #2287](https://redmine.postgresql.org/issues/2287) - Properly handle EXPLAIN queries entered directly by the user in the Query Tool<br>
[Issue #2291](https://redmine.postgresql.org/issues/2291) - Fix error highlighting in the Query Tool<br>
[Issue #2299](https://redmine.postgresql.org/issues/2299) - Fix usage of QString<br>
[Issue #2303](https://redmine.postgresql.org/issues/2303) - Fix ascending/descending sort order in backgrid while clicking on the headers<br>
[Issue #2304](https://redmine.postgresql.org/issues/2304) - Resolve the issue for restoring the table from the backup<br>
[Issue #2305](https://redmine.postgresql.org/issues/2305) - Resolve the issue where Generic function qtLiteral was not adapting values properly when they contain non ascii characters<br>
[Issue #2310](https://redmine.postgresql.org/issues/2310) - Fix Dialog Help where Query Tool/Debugger opens in new browser tab<br>
[Issue #2319](https://redmine.postgresql.org/issues/2319) - Resolve issue where Click on pgAdmin4 logo leads to unauthorized error<br>
[Issue #2321](https://redmine.postgresql.org/issues/2321) - Improved functionality of browser tree when adding new nodes if parent collection node has not loaded<br>
[Issue #2330](https://redmine.postgresql.org/issues/2330) - Ensure the Query Tool displays but does not render HTML returned by the server in the results grid<br>
