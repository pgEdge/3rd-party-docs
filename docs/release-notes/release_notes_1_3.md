# Version 1.3

Release date: 2017-03-10

This release contains a number of features and fixes reported since the release of pgAdmin4 1.2.

# Features

[Issue #2036](https://redmine.postgresql.org/issues/2036) - Query tool efficiency - SlickGrid result set format efficiency<br>
[Issue #2038](https://redmine.postgresql.org/issues/2038) - Query tool efficiency - Incremental back off when polling<br>
[Issue #2163](https://redmine.postgresql.org/issues/2163) - Make syntax highlighting more visible<br>
[Issue #2210](https://redmine.postgresql.org/issues/2210) - Build a universal Python wheel instead of per-python-version ones<br>
[Issue #2215](https://redmine.postgresql.org/issues/2215) - Improve visibility of syntax highlighting colours<br>

# Bug fixes

[Issue #1796](https://redmine.postgresql.org/issues/1796) - Add missing "Run Now" option for pgAdmin jobs<br>
[Issue #1797](https://redmine.postgresql.org/issues/1797) - Resolve encoding issues with DATA_DIR<br>
[Issue #1914](https://redmine.postgresql.org/issues/1914) - Resolved error utf8' codec can't decode byte<br>
[Issue #1983](https://redmine.postgresql.org/issues/1983) - Fix bug in Sql query contains Arabic Charaters<br>
[Issue #2089](https://redmine.postgresql.org/issues/2089) - Add PARALLEL SAFE|UNSAFE|RESTRICTED support<br>
[Issue #2115](https://redmine.postgresql.org/issues/2115) - Fix exclusion constraint reverse engineered SQL<br>
[Issue #2119](https://redmine.postgresql.org/issues/2119) - Fix display of long integers and decimals<br>
[Issue #2126](https://redmine.postgresql.org/issues/2126) - Correct node labels in Preferences for EDB functions and procedures<br>
[Issue #2151](https://redmine.postgresql.org/issues/2151) - Display un-sized varlen column types correctly in the Query Tool<br>
[Issue #2154](https://redmine.postgresql.org/issues/2154) - Fix display of long integers and decimals<br>
[Issue #2159](https://redmine.postgresql.org/issues/2159) - Resolve issue where Query editor is not working with Python2.6<br>
[Issue #2160](https://redmine.postgresql.org/issues/2160) - Various encoding fixes to allow 'ascii' codec to decode byte 0xc3 in position 66: ordinal not in range(128)<br>
[Issue #2166](https://redmine.postgresql.org/issues/2166) - Resolved import/Export issue for a table<br>
[Issue #2173](https://redmine.postgresql.org/issues/2173) - Resolved issues where Sequences API test cases are not working in PG9.2 and PPAS9.2<br>
[Issue #2174](https://redmine.postgresql.org/issues/2174) - Resolved various file-system encoding/decoding related cases<br>
[Issue #2185](https://redmine.postgresql.org/issues/2185) - Removed sorting columns on the treeview<br>
[Issue #2192](https://redmine.postgresql.org/issues/2192) - Fix startup complete tests to ensure we properly poll the server for completed startup<br>
[Issue #2198](https://redmine.postgresql.org/issues/2198) - Fix function arguments when generating create SQL<br>
[Issue #2200](https://redmine.postgresql.org/issues/2200) - Properly handle event trigger functions in different schemas<br>
[Issue #2201](https://redmine.postgresql.org/issues/2201) - Fix renaming of check constraints when the table name is changed at the same time<br>
[Issue #2202](https://redmine.postgresql.org/issues/2202) - Fix issue where Dependents query fails due to non ascii characters<br>
[Issue #2204](https://redmine.postgresql.org/issues/2204) - Fixed issue where pgadmin 4 jobs not showing any activity<br>
[Issue #2205](https://redmine.postgresql.org/issues/2205) - Fix display of boolean nulls in the Query Tool<br>
[Issue #2208](https://redmine.postgresql.org/issues/2208) - Ensure primary key column names are quoted in View Data mode of the Query Tool<br>
[Issue #2212](https://redmine.postgresql.org/issues/2212) - Ensure servers are deleted when their parent group is deleted<br>
[Issue #2213](https://redmine.postgresql.org/issues/2213) - Enable right click on browser tree<br>
[Issue #2218](https://redmine.postgresql.org/issues/2218) - Show the correct indeterminate state when editing new boolean values<br>
[Issue #2228](https://redmine.postgresql.org/issues/2228) - Authenticate the runtime to the server<br>
[Issue #2230](https://redmine.postgresql.org/issues/2230) - Prevent the Slonik logo obscuring the login dialogue on small displays in server mode<br>
