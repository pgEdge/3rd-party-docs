# Version 6.0

Release date: 2021-10-07

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.7.

# New features

[Issue #4211](https://redmine.postgresql.org/issues/4211) -  Added support for OWNED BY Clause for sequences.<br>

# Housekeeping

[Issue #5741](https://redmine.postgresql.org/issues/5741) -  Revisit all the CREATE and DROP DDL's to add appropriate 'IF EXISTS', 'CASCADE' and 'CREATE OR REPLACE'.<br>
[Issue #6129](https://redmine.postgresql.org/issues/6129) -  Port browser tree to React.<br>
[Issue #6588](https://redmine.postgresql.org/issues/6588) -  Port object nodes and properties dialogs to React.<br>
[Issue #6687](https://redmine.postgresql.org/issues/6687) -  Port Grant Wizard to react.<br>
[Issue #6692](https://redmine.postgresql.org/issues/6692) -  Remove GPDB support completely.<br>

# Bug fixes

[Issue #2097](https://redmine.postgresql.org/issues/2097) -  Fixed an issue where grant wizard is unresponsive if the database size is huge.<br>
[Issue #2546](https://redmine.postgresql.org/issues/2546) -  Added support to create the Partitioned table using COLLATE and opclass.<br>
[Issue #3827](https://redmine.postgresql.org/issues/3827) -  Ensure that in the Query History tab, query details should be scrollable.<br>
[Issue #6712](https://redmine.postgresql.org/issues/6712) -  Fixed an issue where collapse and expand arrows mismatch in case of nested IF.<br>
[Issue #6713](https://redmine.postgresql.org/issues/6713) -  Fixed an issue where the last message is not visible in the Debugger.<br>
[Issue #6723](https://redmine.postgresql.org/issues/6723) -  Updated query error row selection color as per dark theme style guide.<br>
[Issue #6724](https://redmine.postgresql.org/issues/6724) -  Fixed an issue where the drop cascade button enables for Databases.<br>
[Issue #6736](https://redmine.postgresql.org/issues/6736) -  Fixed an issue where Refresh view options are not working for materialized view.<br>
[Issue #6755](https://redmine.postgresql.org/issues/6755) -  Fixed keyerror issue in schema diff for 'attnum' and 'edit_types' parameter.<br>
[Issue #6759](https://redmine.postgresql.org/issues/6759) -  Ensure that RLS names should not be editable in the collection node of properties tab.<br>
[Issue #6798](https://redmine.postgresql.org/issues/6798) -  Fixed an issue where Execute button of the query tool gets disabled once we change anything in the data grid.<br>
[Issue #6834](https://redmine.postgresql.org/issues/6834) -  Ensure that SQL help should work for EPAS servers.<br>
