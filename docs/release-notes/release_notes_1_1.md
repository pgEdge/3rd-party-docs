# Version 1.1

Release date: 2016-10-27

This release contains a number of features and fixes reported since the release of pgAdmin4 1.0;

# Features

[Issue #1328](https://redmine.postgresql.org/issues/1328) - Add Python 3.5 Support<br>
[Issue #1859](https://redmine.postgresql.org/issues/1859) - Include wait information on the activity tab of the dashboards<br>

# Bug fixes

[Issue #1155](https://redmine.postgresql.org/issues/1155) - Display the start value when the user creates sequence<br>
[Issue #1531](https://redmine.postgresql.org/issues/1531) - Fix to update privileges for Views and Materials Views where "string indices must be integers error" displayed<br>
[Issue #1574](https://redmine.postgresql.org/issues/1574) - Display SQL in SQL pane for security label in PG and EPAS server<br>
[Issue #1576](https://redmine.postgresql.org/issues/1576) - Make security label option available in procedure properties<br>
[Issue #1577](https://redmine.postgresql.org/issues/1577) - Make debug option available for package function and procedure<br>
[Issue #1596](https://redmine.postgresql.org/issues/1596) - Correct spelling error from evnt_turncate to evnt_truncate<br>
[Issue #1599](https://redmine.postgresql.org/issues/1599) - Ensure the grant wizard works with objects with special characters in the name<br>
[Issue #1622](https://redmine.postgresql.org/issues/1622) - Fix issue using special characters when creating synonym<br>
[Issue #1728](https://redmine.postgresql.org/issues/1728) - Properties refreshing after objects are edited<br>
[Issue #1739](https://redmine.postgresql.org/issues/1739) - Prevent the user from trying to.....<br>
[Issue #1785](https://redmine.postgresql.org/issues/1785) - Correctly identify server type upon first connection<br>
[Issue #1786](https://redmine.postgresql.org/issues/1786) - Ensure errorModel unset property is set correctly when adding a new server<br>
[Issue #1808](https://redmine.postgresql.org/issues/1808) - Set seconds to valid value in pgAgent job schedule<br>
[Issue #1817](https://redmine.postgresql.org/issues/1817) - Display message "server does not support ssl" if server with ca-cert or ca-full added<br>
[Issue #1821](https://redmine.postgresql.org/issues/1821) - Optionally sign both the Mac app bundle and the disk image<br>
[Issue #1822](https://redmine.postgresql.org/issues/1822) - Handle non-ascii responses from the server when connecting<br>
[Issue #1823](https://redmine.postgresql.org/issues/1823) - Attempt to sign the Windows installer, failing with a warning if there's no cert available<br>
[Issue #1824](https://redmine.postgresql.org/issues/1824) - Add documenation for pgAgent<br>
[Issue #1835](https://redmine.postgresql.org/issues/1835) - Allow users to choose SELECT permissions for tables and sequences in the grant wizard<br>
[Issue #1837](https://redmine.postgresql.org/issues/1837) - Fix refreshing of FTS dictionaries which was causing error "Connection to the server has been lost"<br>
[Issue #1838](https://redmine.postgresql.org/issues/1838) - Don't append new objects with the wrong parent in tree browser if the correct one isn't loaded<br>
[Issue #1843](https://redmine.postgresql.org/issues/1843) - Function definition matches value returned from pg_get_functiondef()<br>
[Issue #1845](https://redmine.postgresql.org/issues/1845) - Allow refreshing synonym node.  Does not display message "Unimplemented method (node) for this url (/browser/synonym/nodes/1/7/14301/2200/test)"<br>
[Issue #1847](https://redmine.postgresql.org/issues/1847) - Identify the collation correctly when reverse engineering table SQL.  ERROR:  schema "default" does not exist no longer displayed<br>
[Issue #1849](https://redmine.postgresql.org/issues/1849) - Remove security keys from config.py/config_local.py<br>
[Issue #1857](https://redmine.postgresql.org/issues/1857) - Fix error while renaming FTS dictionary and FTS template nodes<br>
[Issue #1858](https://redmine.postgresql.org/issues/1858) - Ensure the File Manager honours the file type while traversing the directories.<br>
[Issue #1861](https://redmine.postgresql.org/issues/1861) - Properly generate exclusion constraint SQL.<br>
[Issue #1863](https://redmine.postgresql.org/issues/1863) - Correctly quote type names in reverse engineered SQL for tables<br>
[Issue #1864](https://redmine.postgresql.org/issues/1864) - Fix layout of DateTimePicker control help message.<br>
[Issue #1867](https://redmine.postgresql.org/issues/1867) - Allow package bodies to be dropped.<br>
[Issue #1868](https://redmine.postgresql.org/issues/1868) - Resolved issue where Integer type of preferences are not updated<br>
[Issue #1872](https://redmine.postgresql.org/issues/1872) - Fix the file manager when used under Python 3.<br>
[Issue #1877](https://redmine.postgresql.org/issues/1877) - Ensure preferences values are stored properly.<br>
[Issue #1878](https://redmine.postgresql.org/issues/1878) - Ensure steps and schedules can be created in empty jobs.  ProgrammingError: can't adapt type 'Undefined' was displayed<br>
[Issue #1880](https://redmine.postgresql.org/issues/1880) - Add new indexes to the correct parent on the treeview.<br>
