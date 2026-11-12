# Version 4.7

Release date: 2019-05-30

This release contains a number of bug fixes since the release of pgAdmin4 4.6.

# Bug fixes

[Issue #3377](https://redmine.postgresql.org/issues/3377) - In server mode, update all the saved server credentials when user password is changed.<br>
[Issue #3885](https://redmine.postgresql.org/issues/3885) - Fix the responsive layout of the main menu bar.<br>
[Issue #4162](https://redmine.postgresql.org/issues/4162) - Fix syntax error when adding more than one column to the existing table.<br>
[Issue #4164](https://redmine.postgresql.org/issues/4164) - Fix file browser path issue which occurs when client is on Windows and server is on Mac/Linux.<br>
[Issue #4184](https://redmine.postgresql.org/issues/4184) - Added Master Password to increase the security of saved passwords.<br>
[Issue #4194](https://redmine.postgresql.org/issues/4194) - Fix accessibility issue for menu navigation.<br>
[Issue #4208](https://redmine.postgresql.org/issues/4208) - Update the UI logo.<br>
[Issue #4217](https://redmine.postgresql.org/issues/4217) - Fixed CSRF security vulnerability issue, per Alvin Lindstam<br>
[Issue #4218](https://redmine.postgresql.org/issues/4218) - Properly assign dropdownParent in Select2 controls.<br>
[Issue #4219](https://redmine.postgresql.org/issues/4219) - Ensure popper.js is installed when needed.<br>
[Issue #4227](https://redmine.postgresql.org/issues/4227) - Fixed Tab key navigation for Maintenance dialog.<br>
[Issue #4244](https://redmine.postgresql.org/issues/4244) - Fix Tab key issue for Toggle switch controls and button on the dialog footer in Safari browser.<br>
[Issue #4245](https://redmine.postgresql.org/issues/4245) - Ensure that element should get highlighted when they get focus on using Tab key.<br>
[Issue #4246](https://redmine.postgresql.org/issues/4246) - Fixed console error when subnode control is used in panels.<br>
[Issue #4261](https://redmine.postgresql.org/issues/4261) - Stop using application/x-javascript as a mime type and use the RFC-compliant application/javascript instead.<br>
[Issue #4262](https://redmine.postgresql.org/issues/4262) - Fixed error on displaying table properties of a table partitioned by list having a default partition.<br>
[Issue #4263](https://redmine.postgresql.org/issues/4263) - Fix handling of JSON in the Query Tool with NULL elements.<br>
[Issue #4269](https://redmine.postgresql.org/issues/4269) - Fix navigation of switch cells in grids.<br>
[Issue #4275](https://redmine.postgresql.org/issues/4275) - Clarify wording for the NO INHERIT option on constraints, per Michel Feinstein.<br>
[Issue #4276](https://redmine.postgresql.org/issues/4276) - Relax the permission check on the directory containing the config database, as it may fail in some environments such as OpenShift.<br>
[Issue #4278](https://redmine.postgresql.org/issues/4278) - Prevent Backgrid Password cells from losing focus if the browser opens an autocomplete list.<br>
[Issue #4284](https://redmine.postgresql.org/issues/4284) - Fix syntax error when creating a table with a serial column.<br>
