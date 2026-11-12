# Version 6.8

Release date: 2022-04-07

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.7.

# New features

[Issue #7215](https://redmine.postgresql.org/issues/7215) -  Added transaction start time to Server activity sessions view.<br>
[Issue #7249](https://redmine.postgresql.org/issues/7249) -  Added support for unique keys in ERD.<br>
[Issue #7257](https://redmine.postgresql.org/issues/7257) -  Support running the container under OpenShift with alternate UIDs.<br>

# Housekeeping

[Issue #7132](https://redmine.postgresql.org/issues/7132) -  Port Properties panel for collection node, Dashboard, and SQL panel in React.<br>
[Issue #7149](https://redmine.postgresql.org/issues/7149) -  Port preferences dialog to React.<br>

# Bug fixes

[Issue #4256](https://redmine.postgresql.org/issues/4256) -  Fixed an issue where SQL for revoke statements are not shown for databases.<br>
[Issue #5836](https://redmine.postgresql.org/issues/5836) -  Adds a new LDAP authentication configuration parameter that indicates the case sensitivity of the LDAP schema/server.<br>
[Issue #6960](https://redmine.postgresql.org/issues/6960) -  Ensure that the master password dialog is popped up if the crypt key is missing.<br>
[Issue #7059](https://redmine.postgresql.org/issues/7059) -  Fixed an issue where the error is shown on logout when the authentication source is oauth2.<br>
[Issue #7176](https://redmine.postgresql.org/issues/7176) -  Fixed an issue where the browser tree state was not preserved correctly.<br>
[Issue #7197](https://redmine.postgresql.org/issues/7197) -  Fixed an issue where foreign key relationships do not update when the primary key is modified.<br>
[Issue #7216](https://redmine.postgresql.org/issues/7216) -  Ensure that the values of certain fields are prettified in the statistics tab for collection nodes.<br>
[Issue #7221](https://redmine.postgresql.org/issues/7221) -  Ensure objects depending on extensions are not displayed in Schema Diff.<br>
[Issue #7238](https://redmine.postgresql.org/issues/7238) -  Fixed an issue where foreign key is not removed even if the referred table is removed in ERD.<br>
[Issue #7239](https://redmine.postgresql.org/issues/7239) -  Fixed an issue where the newly added table is not visible under the Tables node on refresh.<br>
[Issue #7261](https://redmine.postgresql.org/issues/7261) -  Correct typo in the documentation.<br>
[Issue #7263](https://redmine.postgresql.org/issues/7263) -  Fixed schema diff issue where function's difference DDL was showing incorrectly when arguments had default values with commas.<br>
[Issue #7264](https://redmine.postgresql.org/issues/7264) -  Ensure that the correct user should be selected in the new connection dialog.<br>
[Issue #7265](https://redmine.postgresql.org/issues/7265) -  Fixed schema diff issue in which the option 'null' doesn't appear in the DDL statement for the foreign table.<br>
[Issue #7267](https://redmine.postgresql.org/issues/7267) -  Fixed an issue where unexpected error messages are displayed when users change the language via preferences.<br>
[Issue #7269](https://redmine.postgresql.org/issues/7269) -  Ensure that pgAdmin4 should work with latest jinja2 version.<br>
[Issue #7275](https://redmine.postgresql.org/issues/7275) -  Fixed 'Cannot read properties of undefined' error while creating the table via the ERD tool.<br>
