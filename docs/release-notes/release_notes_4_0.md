# Version 4.0

Release date: 2019-01-10

This release contains a number of features and fixes reported since the release of pgAdmin4 3.6

# Features

[Issue #3589](https://redmine.postgresql.org/issues/3589) - Allow query plans to be downloaded as an SVG file.<br>
[Issue #3692](https://redmine.postgresql.org/issues/3692) - New UI design.<br>
[Issue #3801](https://redmine.postgresql.org/issues/3801) - Allow servers to be pre-loaded into container deployments.<br>

# Bug fixes

[Issue #3083](https://redmine.postgresql.org/issues/3083) - Increase the size of the resize handle of the edit grid text pop-out.<br>
[Issue #3354](https://redmine.postgresql.org/issues/3354) - Fix handling of array types as inputs to the debugger.<br>
[Issue #3433](https://redmine.postgresql.org/issues/3433) - Fix an issue that could cause the Query Tool to fail to render.<br>
[Issue #3549](https://redmine.postgresql.org/issues/3549) - Display event trigger functions correctly on EPAS.<br>
[Issue #3559](https://redmine.postgresql.org/issues/3559) - Further improvements to treeview restoration.<br>
[Issue #3599](https://redmine.postgresql.org/issues/3599) - Run Postfix in the container build so passwords can be reset etc.<br>
[Issue #3619](https://redmine.postgresql.org/issues/3619) - Add titles to the code areas of the Query Tool and Debugger to ensure that panels can be re-docked within them.<br>
[Issue #3679](https://redmine.postgresql.org/issues/3679) - Fix a webpack issue that could cause the Query Tool to fail to render.<br>
[Issue #3702](https://redmine.postgresql.org/issues/3702) - Ensure we display the relation name (and not the OID) in the locks table wherever possible.<br>
[Issue #3711](https://redmine.postgresql.org/issues/3711) - Fix an encoding issue in the Query Tool.<br>
[Issue #3726](https://redmine.postgresql.org/issues/3726) - Include the WHERE clause on EXCLUDE constraints in RE-SQL.<br>
[Issue #3753](https://redmine.postgresql.org/issues/3753) - Fix an issue when user define Cast from smallint->text is created.<br>
[Issue #3757](https://redmine.postgresql.org/issues/3757) - Hide Radio buttons that should not be shown on the maintenance dialogue.<br>
[Issue #3780](https://redmine.postgresql.org/issues/3780) - Ensure that null values handled properly in CSV download.<br>
[Issue #3796](https://redmine.postgresql.org/issues/3796) - Tweak the wording on the Grant Wizard.<br>
[Issue #3797](https://redmine.postgresql.org/issues/3797) - Prevent attempts to bulk-drop schema objects.<br>
[Issue #3798](https://redmine.postgresql.org/issues/3798) - Ensure the browser toolbar buttons work in languages other than English.<br>
[Issue #3805](https://redmine.postgresql.org/issues/3805) - Allow horizontal sizing of the edit grid text pop-out.<br>
[Issue #3809](https://redmine.postgresql.org/issues/3809) - Ensure auto complete should works when first identifier in the FROM clause needs quoting.<br>
[Issue #3810](https://redmine.postgresql.org/issues/3810) - Ensure auto complete should works for columns from a schema-qualified table.<br>
[Issue #3821](https://redmine.postgresql.org/issues/3821) - Ensure identifiers are properly displayed in the plan viewer.<br>
[Issue #3830](https://redmine.postgresql.org/issues/3830) - Make the setup process more robust against aborted executions.<br>
[Issue #3856](https://redmine.postgresql.org/issues/3856) - Fixed an issue while creating export job.<br>
