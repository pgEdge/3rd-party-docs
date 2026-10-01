# Version 4.15

Release date: 2019-11-14

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.14.

# New features

[Issue #1974](https://redmine.postgresql.org/issues/1974) -  Added encrypted password in reverse engineered SQL for roles.<br>
[Issue #3741](https://redmine.postgresql.org/issues/3741) -  Added Dark(Beta) UI Theme option.<br>
[Issue #4006](https://redmine.postgresql.org/issues/4006) -  Support Enable Always and Enable Replica on triggers.<br>
[Issue #4351](https://redmine.postgresql.org/issues/4351) -  Add an option to request confirmation before cancelling/resetting changes on a Properties dialog.<br>
[Issue #4348](https://redmine.postgresql.org/issues/4348) -  Added support for custom theme creation and selection.<br>

# Housekeeping

# Bug fixes

[Issue #3130](https://redmine.postgresql.org/issues/3130) -  Ensure create new object dialog should be opened when alt+shift+n key is pressed on the collection node.<br>
[Issue #3279](https://redmine.postgresql.org/issues/3279) -  Fixed issue where Drop and Disconnect connection menu points are too close to each other.<br>
[Issue #3789](https://redmine.postgresql.org/issues/3789) -  Ensure context menus never get hidden below the menu bar.<br>
[Issue #3859](https://redmine.postgresql.org/issues/3859) -  Rename the context menu from 'Drop Server' to 'Remove Server'.<br>
[Issue #3913](https://redmine.postgresql.org/issues/3913) -  Ensure the correct "running at" agent is shown when a pgAgent job is executing.<br>
[Issue #3915](https://redmine.postgresql.org/issues/3915) -  Fix an issue in the Query Tool where shortcut keys could be ignored following a query error.<br>
[Issue #3999](https://redmine.postgresql.org/issues/3999) -  Fix the toggle case shortcut key combination.<br>
[Issue #4173](https://redmine.postgresql.org/issues/4173) -  Fix an issue where a black arrow-kind image is displaying at the background of browser tree images.<br>
[Issue #4191](https://redmine.postgresql.org/issues/4191) -  Ensure comments are shown in reverse engineered SQL for table partitions.<br>
[Issue #4242](https://redmine.postgresql.org/issues/4242) -  Handle NULL values appropriately when sorting backgrid tables.<br>
[Issue #4341](https://redmine.postgresql.org/issues/4341) -  Give appropriate error messages when the user tries to use an blank master password.<br>
[Issue #4451](https://redmine.postgresql.org/issues/4451) -  Remove arbitrary (and incorrect) requirement that composite types must have at least two members.<br>
[Issue #4459](https://redmine.postgresql.org/issues/4459) -  Don't quote bigints when copying them from the Query Tool results grid.<br>
[Issue #4482](https://redmine.postgresql.org/issues/4482) -  Ensure compression level is passed to pg_dump when backing up in directory format.<br>
[Issue #4483](https://redmine.postgresql.org/issues/4483) -  Ensure the number of jobs can be specified when backing up in directory format.<br>
[Issue #4564](https://redmine.postgresql.org/issues/4564) -  Ensure Javascript errors during Query Tool execution are reported as such and not as Ajax errors.<br>
[Issue #4610](https://redmine.postgresql.org/issues/4610) -  Suppress Enter key presses in Alertify dialogues when the come from Select2 controls to allow item selection with Enter.<br>
[Issue #4647](https://redmine.postgresql.org/issues/4647) -  Ensure that units are respected when sorting by file size in the File dialog.<br>
[Issue #4730](https://redmine.postgresql.org/issues/4730) -  Ensure all messages are retained in the Query Tool from long running queries.<br>
[Issue #4734](https://redmine.postgresql.org/issues/4734) -  Updated documentation for the delete row button that only strikeout the row instead of deleting it.<br>
[Issue #4779](https://redmine.postgresql.org/issues/4779) -  Updated documentation for the query tool toolbar buttons.<br>
[Issue #4835](https://redmine.postgresql.org/issues/4835) -  Fixed an issue where psql of v12 throwing "symbol not found" error while running Maintenance and Import/Export.<br>
[Issue #4845](https://redmine.postgresql.org/issues/4845) -  Fixed potential error in the properties dialog for the Code tab.<br>
[Issue #4850](https://redmine.postgresql.org/issues/4850) -  Fixed an issue where Datetimepicker control opens when clicking on the label.<br>
[Issue #4895](https://redmine.postgresql.org/issues/4895) -  Fixed potential issue in reset function for nested objects.<br>
[Issue #4896](https://redmine.postgresql.org/issues/4896) -  Fixed an issue where escape key not working to close the open/save file dialog.<br>
[Issue #4906](https://redmine.postgresql.org/issues/4906) -  Fixed an issue where keyboard shortcut for context menu is not working when using Firefox on CentOS7.<br>
[Issue #4924](https://redmine.postgresql.org/issues/4924) -  Fixed docker container exit issue occurs due to change in Gunicorn's latest version.<br>
