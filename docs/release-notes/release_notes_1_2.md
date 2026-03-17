# Version 1.2

Release date: 2017-02-10

This release contains a number of features and fixes reported since the release of pgAdmin4 1.1.

# Features

[Issue #1375](https://redmine.postgresql.org/issues/1375) - Migrate the runtime to QtWebEngine from QtWebKit<br>
[Issue #1765](https://redmine.postgresql.org/issues/1765) - Find and replace functionality with regexp and group replacement<br>
[Issue #1789](https://redmine.postgresql.org/issues/1789) - Column width of data output panel should fit to data (as pgAdmin III)<br>
[Issue #1790](https://redmine.postgresql.org/issues/1790) - [Web] Support setting a field's value to "null"<br>
[Issue #1848](https://redmine.postgresql.org/issues/1848) - macOS appbundle is missing postgresql binaries for import etc.<br>
[Issue #1910](https://redmine.postgresql.org/issues/1910) - Remember last used directory in the file manager<br>
[Issue #1911](https://redmine.postgresql.org/issues/1911) - Direct path navigation in the file manager<br>
[Issue #1922](https://redmine.postgresql.org/issues/1922) - Improve handling of corrupt configuration databases<br>
[Issue #1963](https://redmine.postgresql.org/issues/1963) - Add a Chinese (Simplified) translation<br>
[Issue #1964](https://redmine.postgresql.org/issues/1964) - Create a docs tarball along with the source tarball<br>
[Issue #2025](https://redmine.postgresql.org/issues/2025) - Allow the SQL Editors to word-wrap<br>
[Issue #2124](https://redmine.postgresql.org/issues/2124) - Create a template loader to simplify SQL template location, and remove duplicate templates<br>

# Bug fixes

[Issue #1227](https://redmine.postgresql.org/issues/1227) - Display improved error message for Debugger listener starting error and reset between executions<br>
[Issue #1267](https://redmine.postgresql.org/issues/1267) - Fix issue where MINIFY_HTML doesn't work with the docs<br>
[Issue #1364](https://redmine.postgresql.org/issues/1364) - Ensure dialogue control buttons are consistent<br>
[Issue #1394](https://redmine.postgresql.org/issues/1394) - Fix Table dialogue column specification issues<br>
[Issue #1432](https://redmine.postgresql.org/issues/1432) - Enhanced OSX File Browser<br>
[Issue #1585](https://redmine.postgresql.org/issues/1585) - Cannot save scripts to the network<br>
[Issue #1599](https://redmine.postgresql.org/issues/1599) - Ensure the grant wizard works with objects with special characters in the name<br>
[Issue #1603](https://redmine.postgresql.org/issues/1603) - Fix quoting of objects names for external utilities.<br>
[Issue #1679](https://redmine.postgresql.org/issues/1679) - Re-engineer the background process executor to avoid using sqlite as some builds of components it relies on do not support working in forked children<br>
[Issue #1680](https://redmine.postgresql.org/issues/1680) - Render column headers at the correct width in the Query Tool under Firefox<br>
[Issue #1729](https://redmine.postgresql.org/issues/1729) - Improve display of role options<br>
[Issue #1730](https://redmine.postgresql.org/issues/1730) - Improve the display of role membership on both the properties panel and role dialogue<br>
[Issue #1745](https://redmine.postgresql.org/issues/1745) - Ensure breakpoints are cleared properly when working with Debugger<br>
[Issue #1747](https://redmine.postgresql.org/issues/1747) - Add newly created triggers to the treeview<br>
[Issue #1780](https://redmine.postgresql.org/issues/1780) - Properly size the SQL Editor gutter as the width of the line numbers increases<br>
[Issue #1792](https://redmine.postgresql.org/issues/1792) - List files and folders alphabetically<br>
[Issue #1800](https://redmine.postgresql.org/issues/1800) - Handle the template property on databases appropriately<br>
[Issue #1801](https://redmine.postgresql.org/issues/1801) - Handle databases with datallowconn == false<br>
[Issue #1807](https://redmine.postgresql.org/issues/1807) - Properly detect when files have changed in the Query Tool and set flag accordingly<br>
[Issue #1830](https://redmine.postgresql.org/issues/1830) - Fix a SQL error when reverse-engineering ROLE SQL on EPAS servers<br>
[Issue #1832](https://redmine.postgresql.org/issues/1832) - Prevent attempts to access what may be an empty list in Dependancies tab<br>
[Issue #1840](https://redmine.postgresql.org/issues/1840) - Enable/disable NULLs and ASC/DESC options for index columns and exclusion constraints appropriately<br>
[Issue #1842](https://redmine.postgresql.org/issues/1842) - Show index columns in the correct order in RE-SQL<br>
[Issue #1855](https://redmine.postgresql.org/issues/1855) - Ensure dialogue panels show their errors themselves, and not in the properties panel when creating Trigger Function<br>
[Issue #1865](https://redmine.postgresql.org/issues/1865) - Properly schema qualify domains when reverse engineering SQL<br>
[Issue #1874](https://redmine.postgresql.org/issues/1874) - Add file resources to the windows runtime<br>
[Issue #1893](https://redmine.postgresql.org/issues/1893) - Fix refreshing of Unique constraints<br>
[Issue #1896](https://redmine.postgresql.org/issues/1896) - Use the correct OID for retrieving properties of freshly created exclusion constraints<br>
[Issue #1899](https://redmine.postgresql.org/issues/1899) - Properly quote role names when specifying function ownership<br>
[Issue #1909](https://redmine.postgresql.org/issues/1909) - Handle startup errors more gracefully in the runtime<br>
[Issue #1912](https://redmine.postgresql.org/issues/1912) - Properly format arguments passed by triggers to functions<br>
[Issue #1919](https://redmine.postgresql.org/issues/1919) - Ensure all changes to rows are stored in the data editor<br>
[Issue #1924](https://redmine.postgresql.org/issues/1924) - Ensure the check_option is only set when editing views when appropriate<br>
[Issue #1936](https://redmine.postgresql.org/issues/1936) - Don't strip \r\n from "Download as CSV" batches of rows, as it leads to malformed data<br>
[Issue #1937](https://redmine.postgresql.org/issues/1937) - Generate mSQL for new schemas correctly<br>
[Issue #1938](https://redmine.postgresql.org/issues/1938) - Fix sorting of numerics in the statistics grids<br>
[Issue #1939](https://redmine.postgresql.org/issues/1939) - Updated dynamic default for the window size (90% x 90%)<br>
[Issue #1949](https://redmine.postgresql.org/issues/1949) - Ensure trigger function names are schema qualified in trigger RE-SQL<br>
[Issue #1951](https://redmine.postgresql.org/issues/1951) - Fix issue where nnable to browse table columns when oid values exceeed max int<br>
[Issue #1953](https://redmine.postgresql.org/issues/1953) - Add display messages and notices received in the Query Tool<br>
[Issue #1961](https://redmine.postgresql.org/issues/1961) - Fix upgrade check on Python 3<br>
[Issue #1962](https://redmine.postgresql.org/issues/1962) - Ensure treeview collection nodes are translated in the UI<br>
[Issue #1967](https://redmine.postgresql.org/issues/1967) - Store layout changes on each adjustment<br>
[Issue #1976](https://redmine.postgresql.org/issues/1976) - Prevent users selecting elements of the UI that shouldn't be selectable<br>
[Issue #1979](https://redmine.postgresql.org/issues/1979) - Deal with Function arguments correctly in the properties dialogue<br>
[Issue #1986](https://redmine.postgresql.org/issues/1986) - Fix various encoding issues with multibyte paths and filenames resulting in empty file save<br>
[Issue #1992](https://redmine.postgresql.org/issues/1992) - Quote identifiers correctly in auto-complete<br>
[Issue #1994](https://redmine.postgresql.org/issues/1994) - Update to show modifications in edit grid<br>
[Issue #2000](https://redmine.postgresql.org/issues/2000) - Allow setting of effective_io_concurrency on tablespaces in 9.6+<br>
[Issue #2005](https://redmine.postgresql.org/issues/2005) - Fix various mis-spellings of VACUUM<br>
[Issue #2006](https://redmine.postgresql.org/issues/2006) - Fix error when modifying table name or set schema on tables with postgis geometry column<br>
[Issue #2007](https://redmine.postgresql.org/issues/2007) - Correctly sort rows by the pkey when viewing first/last 100<br>
[Issue #2009](https://redmine.postgresql.org/issues/2009) - Reset the column list properly if the access method is changed on an index to ensure error handling works correctly<br>
[Issue #2012](https://redmine.postgresql.org/issues/2012) - Prevent attempts to create server groups with no name<br>
[Issue #2015](https://redmine.postgresql.org/issues/2015) - Enable trigger option when user tries to change Row trigger value through properties section<br>
[Issue #2024](https://redmine.postgresql.org/issues/2024) - Properly handle setting comments and other options on databases with allowconn = False<br>
[Issue #2026](https://redmine.postgresql.org/issues/2026) - Improve detection of the pldbgapi extension and functions before allowing debugging<br>
[Issue #2027](https://redmine.postgresql.org/issues/2027) - Fix inconsistent table styling<br>
[Issue #2028](https://redmine.postgresql.org/issues/2028) - Fix display of double scrollbars on the grant wizard<br>
[Issue #2032](https://redmine.postgresql.org/issues/2032) - Fix time formatting on dashboards<br>
[Issue #2033](https://redmine.postgresql.org/issues/2033) - Show icons for unique and exclusion constraints in the dependency/dependents panels<br>
[Issue #2045](https://redmine.postgresql.org/issues/2045) - Update copyright year on doc page<br>
[Issue #2046](https://redmine.postgresql.org/issues/2046) - Fix error when setting up regression on Windows for pgadmin4<br>
[Issue #2047](https://redmine.postgresql.org/issues/2047) - Ensure dialogues cannot be moved under the navbar<br>
[Issue #2061](https://redmine.postgresql.org/issues/2061) - Enable/disable NULLs and ASC/DESC options for index columns and exclusion constraints appropriately<br>
[Issue #2065](https://redmine.postgresql.org/issues/2065) - Improve display of columns of exclusion contraints and foreign keys in the properties lists<br>
[Issue #2069](https://redmine.postgresql.org/issues/2069) - Correct tablespace displayed in table properties<br>
[Issue #2076](https://redmine.postgresql.org/issues/2076) - Handle sized time/timestamp columns correctly<br>
[Issue #2109](https://redmine.postgresql.org/issues/2109) - Update copyright year<br>
[Issue #2110](https://redmine.postgresql.org/issues/2110) - Handle saved directories that no longer exist gracefully<br>
[Issue #2112](https://redmine.postgresql.org/issues/2026) - Enable comments on Initial database through right Click<br>
[Issue #2133](https://redmine.postgresql.org/issues/2133) - Fix display of graphical query plans for UPDATE/DELETE queries<br>
[Issue #2138](https://redmine.postgresql.org/issues/2138) - Fix display of zeros in read-only grid editors<br>
[Issue #2139](https://redmine.postgresql.org/issues/2139) - Fixed issue causing Message (Connection to the server has been lost.) displayed with Materialized view and view under sql tab<br>
[Issue #2152](https://redmine.postgresql.org/issues/2152) - Fix handling of "char" columns<br>
[Issue #2156](https://redmine.postgresql.org/issues/2156) - Added compatibility fixes for newer versions of Jinja2 (e.g. 2.9.5+)<br>
