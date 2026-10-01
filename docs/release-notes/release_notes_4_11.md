# Version 4.11

Release date: 2019-07-25

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.10.

# New features

[Issue #1760](https://redmine.postgresql.org/issues/1760) -  Add support for editing of resultsets in the Query Tool, if the data can be identified as updatable.<br>
[Issue #4318](https://redmine.postgresql.org/issues/4318) -  Set the mouse cursor appropriately based on the layout lock state.<br>
[Issue #4335](https://redmine.postgresql.org/issues/4335) -  Add EXPLAIN options for SETTINGS and SUMMARY.<br>

# Housekeeping

[Issue #4415](https://redmine.postgresql.org/issues/4415) -  Add Reverse Engineered SQL tests for Roles and Resource Groups.<br>
[Issue #4441](https://redmine.postgresql.org/issues/4441) -  Add Reverse Engineered SQL tests for FDWs.<br>
[Issue #4452](https://redmine.postgresql.org/issues/4452) -  Add Reverse Engineered SQL tests for Languages.<br>
[Issue #4453](https://redmine.postgresql.org/issues/4453) -  Add Reverse Engineered SQL tests for Extensions.<br>
[Issue #4454](https://redmine.postgresql.org/issues/4454) -  Add Reverse Engineered SQL tests for FTS Configurations.<br>
[Issue #4456](https://redmine.postgresql.org/issues/4456) -  Add Reverse Engineered SQL tests for Packages.<br>
[Issue #4460](https://redmine.postgresql.org/issues/4460) -  Add Reverse Engineered SQL tests for FTS Dictionaries.<br>
[Issue #4463](https://redmine.postgresql.org/issues/4463) -  Add Reverse Engineered SQL tests for Domains.<br>
[Issue #4464](https://redmine.postgresql.org/issues/4464) -  Add Reverse Engineered SQL tests for Collations.<br>
[Issue #4468](https://redmine.postgresql.org/issues/4468) -  Add Reverse Engineered SQL tests for Types.<br>
[Issue #4469](https://redmine.postgresql.org/issues/4469) -  Add Reverse Engineered SQL tests for Sequences.<br>
[Issue #4471](https://redmine.postgresql.org/issues/4471) -  Add Reverse Engineered SQL tests for FTS Parsers.<br>
[Issue #4475](https://redmine.postgresql.org/issues/4475) -  Add Reverse Engineered SQL tests for Constraints.<br>

# Bug fixes

[Issue #3919](https://redmine.postgresql.org/issues/3919) -  Allow keyboard navigation of all controls on subnode grids.<br>
[Issue #3996](https://redmine.postgresql.org/issues/3996) -  Fix dropping of pgAgent schedules through the Job properties.<br>
[Issue #4224](https://redmine.postgresql.org/issues/4224) -  Prevent flickering of large tooltips on the Graphical EXPLAIN canvas.<br>
[Issue #4389](https://redmine.postgresql.org/issues/4389) -  Fix an error that could be seen when editing column privileges.<br>
[Issue #4393](https://redmine.postgresql.org/issues/4393) -  Ensure parameter values are quoted when needed when editing roles.<br>
[Issue #4395](https://redmine.postgresql.org/issues/4395) -  EXPLAIN options should be Query Tool instance-specific.<br>
[Issue #4427](https://redmine.postgresql.org/issues/4427) -  Fix an error while retrieving json data from the table.<br>
[Issue #4428](https://redmine.postgresql.org/issues/4428) -  Fix 'malformed array literal' error when updating a pgAgent job.<br>
[Issue #4429](https://redmine.postgresql.org/issues/4429) -  Ensure drag/drop from the treeview works as expected on Firefox.<br>
[Issue #4437](https://redmine.postgresql.org/issues/4437) -  Fix table icon issue when updating any existing field.<br>
[Issue #4442](https://redmine.postgresql.org/issues/4442) -  Ensure browser should not be started by Selenium when feature tests are excluded from a test run.<br>
[Issue #4446](https://redmine.postgresql.org/issues/4446) -  Use ROLE consistently when generating RE-SQL for roles, not USER.<br>
[Issue #4448](https://redmine.postgresql.org/issues/4448) -  Fix an error seen when updating a connection string in a pgAgent job step.<br>
[Issue #4450](https://redmine.postgresql.org/issues/4450) -  Fix reverse engineered sql for Foreign Data Wrapper created on EPAS server in redwood mode.<br>
[Issue #4462](https://redmine.postgresql.org/issues/4462) -  Fix some minor UI issues on IE11.<br>
[Issue #4470](https://redmine.postgresql.org/issues/4470) -  Fix sequence reverse engineered SQL generation with quoted names on PG/EPAS 10+.<br>
[Issue #4484](https://redmine.postgresql.org/issues/4484) -  Fix an issue where Explain and Explain Analyze are not working, it's regression of #1760.<br>
[Issue #4485](https://redmine.postgresql.org/issues/4485) -  Fix an issue where Filter toolbar button is not working in view/edit data, it's regression of keyboard navigation.<br>
