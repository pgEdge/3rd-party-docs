# Version 4.24

Release date: 2020-07-23

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.23.

# New features

[Issue #5235](https://redmine.postgresql.org/issues/5235) -  Support configuration files that are external to the application installation.<br>
[Issue #5484](https://redmine.postgresql.org/issues/5484) -  Added support for LDAP authentication with different DN by setting the dedicated user for the LDAP connection.<br>
[Issue #5583](https://redmine.postgresql.org/issues/5583) -  Added support for schema level restriction.<br>
[Issue #5601](https://redmine.postgresql.org/issues/5601) -  Added RLS Policy support in Schema Diff.<br>
[Issue #5622](https://redmine.postgresql.org/issues/5622) -  Added support for permissive/restricted policy type while creating RLS Policy.<br>
[Issue #5650](https://redmine.postgresql.org/issues/5650) -  Added support for LDAP anonymous binding.<br>
[Issue #5653](https://redmine.postgresql.org/issues/5653) -  Added High Contrast theme support.<br>

# Housekeeping

[Issue #5323](https://redmine.postgresql.org/issues/5323) -  Improve code coverage and API test cases for Foreign Data Wrapper.<br>
[Issue #5326](https://redmine.postgresql.org/issues/5326) -  Improve code coverage and API test cases for Domain and Domain Constraints.<br>
[Issue #5329](https://redmine.postgresql.org/issues/5329) -  Improve code coverage and API test cases for FTS Configuration, FTS Parser, FTS Dictionaries, and FTS Template.<br>

# Bug fixes

[Issue #3814](https://redmine.postgresql.org/issues/3814) -  Fixed issue of error message not getting displayed when filename is empty for backup, restore, and import/export.<br>
[Issue #3851](https://redmine.postgresql.org/issues/3851) -  Add proper indentation to the code while generating functions, procedures, and trigger functions.<br>
[Issue #4235](https://redmine.postgresql.org/issues/4235) -  Fixed tab indent issue on a selection of lines is deleting the content when 'use spaces == true' in the preferences.<br>
[Issue #5137](https://redmine.postgresql.org/issues/5137) -  Fixed save button enable issue when focusing in and out of numeric input field.<br>
[Issue #5287](https://redmine.postgresql.org/issues/5287) -  Fixed dark theme-related CSS and modify the color codes.<br>
[Issue #5414](https://redmine.postgresql.org/issues/5414) -  Use QStandardPaths::AppLocalDataLocation in the runtime to determine where to store runtime logs.<br>
[Issue #5463](https://redmine.postgresql.org/issues/5463) -  Fixed an issue where CSV download quotes numeric columns.<br>
[Issue #5470](https://redmine.postgresql.org/issues/5470) -  Fixed backgrid row hover issue where on hover background color is set for edit and delete cell only.<br>
[Issue #5530](https://redmine.postgresql.org/issues/5530) -  Ensure that the referenced table should be displayed on foreign key constraints.<br>
[Issue #5554](https://redmine.postgresql.org/issues/5554) -  Replace the runtime themes with ones that don't have sizing issues.<br>
[Issue #5569](https://redmine.postgresql.org/issues/5569) -  Fixed reverse engineered SQL for partitions when storage parameters are specified.<br>
[Issue #5577](https://redmine.postgresql.org/issues/5577) -  Include LICENSE and DEPENDENCIES [inventory] files in official packages.<br>
[Issue #5621](https://redmine.postgresql.org/issues/5621) -  Remove extra brackets from reverse engineering SQL of RLS Policy.<br>
[Issue #5629](https://redmine.postgresql.org/issues/5629) -  Fixed an issue where the user is able to edit properties when some of the collection nodes are selected.<br>
[Issue #5630](https://redmine.postgresql.org/issues/5630) -  Fixed an issue where installation of pgadmin4 not working on 32-bit Windows.<br>
[Issue #5631](https://redmine.postgresql.org/issues/5631) -  Fixed 'cant execute empty query' issue when remove the value of 'USING' or 'WITH CHECK' option of RLS Policy.<br>
[Issue #5633](https://redmine.postgresql.org/issues/5633) -  Ensure that create RLS Policy menu should not be visible for catalog objects.<br>
[Issue #5647](https://redmine.postgresql.org/issues/5647) -  Fixed an issue where difference DDL is showing the wrong SQL when changing the policy owner.<br>
[Issue #5662](https://redmine.postgresql.org/issues/5662) -  Fixed accessibility issue where few dialogs are not rendering properly when we zoomed in browser window 200% and screen resolution is low.<br>
[Issue #5666](https://redmine.postgresql.org/issues/5666) -  Added missing dependencies/dependent and corrected some wrongly identified.<br>
[Issue #5673](https://redmine.postgresql.org/issues/5673) -  Fixed an issue where fetching the schema throws an error if the database is not connected in Schema Diff.<br>
[Issue #5675](https://redmine.postgresql.org/issues/5675) -  Fixed CSRF errors when pgAdmin opened in an iframe on safari browser.<br>
[Issue #5677](https://redmine.postgresql.org/issues/5677) -  Fixed text color issue in explain analyze for the Dark theme.<br>
[Issue #5686](https://redmine.postgresql.org/issues/5686) -  Fixed issue where the user was not able to update policy if the policy is created with space.<br>
