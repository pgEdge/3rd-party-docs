# Version 6.4

Release date: 2022-01-13

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.3.

# New features

[Issue #4803](https://redmine.postgresql.org/issues/4803) -  Added support to import/export server groups and servers from GUI.<br>

# Housekeeping

[Issue #7018](https://redmine.postgresql.org/issues/7018) -  Port Restore dialog to React.<br>
[Issue #7019](https://redmine.postgresql.org/issues/7019) -  Port Maintenance dialog to React.<br>

# Bug fixes

[Issue #6745](https://redmine.postgresql.org/issues/6745) -  Fixed an issue where Tablespace is created though an error is shown on the dialog.<br>
[Issue #7003](https://redmine.postgresql.org/issues/7003) -  Fixed an issue where Explain Analyze shows negative exclusive time.<br>
[Issue #7034](https://redmine.postgresql.org/issues/7034) -  Fixed an issue where Columns with default value not showing when adding a new row.<br>
[Issue #7075](https://redmine.postgresql.org/issues/7075) -  Ensure that help should be visible properly for Procedures.<br>
[Issue #7077](https://redmine.postgresql.org/issues/7077) -  Fixed an issue where the Owner is not displayed in the reverse engineering SQL for Procedures.<br>
[Issue #7078](https://redmine.postgresql.org/issues/7078) -  Fixed an issue where an operation error message pop up showing the database object's name incorrectly.<br>
[Issue #7081](https://redmine.postgresql.org/issues/7081) -  Fixed an issue in SQL generation for PostgreSQL-14 functions.<br>
[Issue #7093](https://redmine.postgresql.org/issues/7093) -  Fixed an issue where SubPlans may overlap other nodes & make them inaccessible in Graphical EXPLAIN View.<br>
[Issue #7096](https://redmine.postgresql.org/issues/7096) -  Ensure that Truncate and Reset statistics should work.<br>
[Issue #7102](https://redmine.postgresql.org/issues/7102) -  Fixed a schema diff issue where generated script adds unwanted line endings for Functions/Procedures/Trigger Functions.<br>
