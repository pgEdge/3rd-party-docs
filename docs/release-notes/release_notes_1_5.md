# Version 1.5

Release date: 2017-05-19

This release contains a number of features and fixes reported since the release of pgAdmin4 1.4.

# Features

[Issue #2216](https://redmine.postgresql.org/issues/2216) - Allow column or row selection in the Query Tool<br>

# Bug fixes

[Issue #2225](https://redmine.postgresql.org/issues/2225) - Hide menu options for creating objects, if the object type is set to hidden. Includes Jasmine tests<br>
[Issue #2253](https://redmine.postgresql.org/issues/2253) - Fix various issues in CSV file download feature<br>
[Issue #2257](https://redmine.postgresql.org/issues/2257) - Improve handling of nulls and default values in the data editor<br>
[Issue #2271](https://redmine.postgresql.org/issues/2271) - Don't change the trigger icon back to "enabled" when the trigger is updated when it's disabled<br>
[Issue #2284](https://redmine.postgresql.org/issues/2284) - Allow creation of tables with pure numeric names<br>
[Issue #2292](https://redmine.postgresql.org/issues/2292) - Only reconnect to databases that were previously connected<br>
[Issue #2314](https://redmine.postgresql.org/issues/2314) - Fix various issues in CSV file download feature<br>
[Issue #2315](https://redmine.postgresql.org/issues/2315) - Fix sorting of sizes on the statistics views by sorting raw values and prettifying on the client side. Includes Jasmine tests for the prettyfying function<br>
[Issue #2318](https://redmine.postgresql.org/issues/2318) - Order foreign table columns correctly<br>
[Issue #2331](https://redmine.postgresql.org/issues/2331) - Fix binary search algorithm so new treeview nodes are added in the correct position<br>
[Issue #2336](https://redmine.postgresql.org/issues/2336) - Update inode info when refreshing treeview nodes.<br>
[Issue #2339](https://redmine.postgresql.org/issues/2339) - Ensure the treeview can be scrolled horizontally<br>
[Issue #2350](https://redmine.postgresql.org/issues/2350) - Fix handling of default parameters ordering in functions<br>
[Issue #2354](https://redmine.postgresql.org/issues/2354) - Fix the Backup module where it was not working if user changes its preference language other than English<br>
[Issue #2356](https://redmine.postgresql.org/issues/2356) - Ensure errors thrown when deleting rows in the Query Tool in edit mode are shown properly<br>
[Issue #2360](https://redmine.postgresql.org/issues/2360) - Fix various issues in CSV file download feature<br>
[Issue #2369](https://redmine.postgresql.org/issues/2369) - Support loading files with Unicode BOMs<br>
[Issue #2377](https://redmine.postgresql.org/issues/2377) - Update psycopg2 version for PostgreSQL 10 compatibility<br>
[Issue #2379](https://redmine.postgresql.org/issues/2379) - Make various improvements to the NULL/DEFAULT handling in the data editor<br>
[Issue #2405](https://redmine.postgresql.org/issues/2405) - Ensure object names are properly escaped for external process management<br>
[Issue #2410](https://redmine.postgresql.org/issues/2410) - Fix PostgreSQL 10.0 compatibility issues<br>
