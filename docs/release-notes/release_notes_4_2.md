# Version 4.2

Release date: 2019-02-07

This release contains a number of fixes reported since the release of pgAdmin4 4.1

# Bug fixes

[Issue #3051](https://redmine.postgresql.org/issues/3051) - Replace Bootstrap switch with Bootstrap4 toggle to improve the performance.<br>
[Issue #3272](https://redmine.postgresql.org/issues/3272) - Replace the PyCrypto module with the cryptography module.<br>
[Issue #3453](https://redmine.postgresql.org/issues/3453) - Fixed SQL for foreign table options.<br>
[Issue #3475](https://redmine.postgresql.org/issues/3475) - Fixed execution time to show Hours part for long running queries in Query Tool.<br>
[Issue #3608](https://redmine.postgresql.org/issues/3608) - Messages tab of Query Tool should be clear on subsequent execution of table/view using View/Edit Data.<br>
[Issue #3609](https://redmine.postgresql.org/issues/3609) - Clear drop-down menu should be disabled for View/Edit Data.<br>
[Issue #3664](https://redmine.postgresql.org/issues/3664) - Fixed Statistics panel hang issue for 1000+ tables.<br>
[Issue #3693](https://redmine.postgresql.org/issues/3693) - Proper error should be thrown when server group is created with existing name.<br>
[Issue #3695](https://redmine.postgresql.org/issues/3695) - Ensure long string should be wrap in alertify dialogs.<br>
[Issue #3697](https://redmine.postgresql.org/issues/3697) - Ensure that output of the query should be displayed even if Data Output window is detached from the Query Tool.<br>
[Issue #3740](https://redmine.postgresql.org/issues/3740) - Inline edbspl trigger functions should not be visible in Grant Wizard.<br>
[Issue #3774](https://redmine.postgresql.org/issues/3774) - Proper SQL should be generated when create function with return type as custom type argument.<br>
[Issue #3800](https://redmine.postgresql.org/issues/3800) - Ensure that database restriction of server dialog should work with special characters.<br>
[Issue #3811](https://redmine.postgresql.org/issues/3811) - Ensure that Backup/Restore button should work on single click.<br>
[Issue #3837](https://redmine.postgresql.org/issues/3837) - Fixed SQL for when clause while creating Trigger.<br>
[Issue #3838](https://redmine.postgresql.org/issues/3838) - Proper SQL should be generated when creating/changing column with custom type argument.<br>
[Issue #3840](https://redmine.postgresql.org/issues/3840) - Ensure that file format combo box value should be retained when hidden files checkbox is toggled.<br>
[Issue #3846](https://redmine.postgresql.org/issues/3846) - Proper SQL should be generated when create procedure with custom type arguments.<br>
[Issue #3849](https://redmine.postgresql.org/issues/3849) - Ensure that browser should warn before close or refresh.<br>
[Issue #3850](https://redmine.postgresql.org/issues/3850) - Fixed EXEC script for procedures.<br>
[Issue #3853](https://redmine.postgresql.org/issues/3853) - Proper SQL should be generated when create domain of type interval with precision.<br>
[Issue #3858](https://redmine.postgresql.org/issues/3858) - Drop-down should be closed when click on any other toolbar button.<br>
[Issue #3862](https://redmine.postgresql.org/issues/3862) - Fixed keyboard navigation for dialog tabs.<br>
[Issue #3865](https://redmine.postgresql.org/issues/3865) - Increase frames splitter mouse hover area to make it easier to resize.<br>
[Issue #3871](https://redmine.postgresql.org/issues/3871) - Fixed alignment of tree arrow icons for Internet Explorer.<br>
[Issue #3872](https://redmine.postgresql.org/issues/3872) - Ensure object names in external process dialogues are properly escaped.<br>
[Issue #3891](https://redmine.postgresql.org/issues/3891) - Correct order of Save and Cancel button for json/jsonb editing.<br>
[Issue #3897](https://redmine.postgresql.org/issues/3897) - Data should be updated properly for FTS Configurations, FTS Dictionaries, FTS Parsers and FTS Templates.<br>
[Issue #3899](https://redmine.postgresql.org/issues/3899) - Fixed unable to drop multiple Rules and Foreign Tables from properties tab.<br>
[Issue #3903](https://redmine.postgresql.org/issues/3903) - Fixed Query Tool Initialization Error.<br>
[Issue #3908](https://redmine.postgresql.org/issues/3908) - Fixed keyboard navigation for Select2 and Privilege cell in Backgrid.<br>
[Issue #3916](https://redmine.postgresql.org/issues/3916) - Correct schema should be displayed in Materialized View dialog.<br>
[Issue #3927](https://redmine.postgresql.org/issues/3927) - Fixed debugger issue for procedure inside package for EPAS servers.<br>
[Issue #3929](https://redmine.postgresql.org/issues/3929) - Fix alignment of help messages in properties panels.<br>
[Issue #3932](https://redmine.postgresql.org/issues/3932) - Fix alignment of submenu for Internet Explorer.<br>
[Issue #3935](https://redmine.postgresql.org/issues/3935) - Ensure that grant wizard should list down functions for EPAS server running with no-redwood-compat mode.<br>
[Issue #3941](https://redmine.postgresql.org/issues/3941) - Dashboard graph optimization.<br>
[Issue #3954](https://redmine.postgresql.org/issues/3954) - Remove Python 2.6 code that's now obsolete.<br>
[Issue #3955](https://redmine.postgresql.org/issues/3955) - Expose the bind address in the Docker container via PGADMIN_BIND_ADDRESS.<br>
[Issue #3961](https://redmine.postgresql.org/issues/3961) - Exclude HTTPExceptions from the all_exception_handler as they should be returned as-is.<br>
