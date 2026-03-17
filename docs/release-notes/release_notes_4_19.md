# Version 4.19

Release date: 2020-03-05

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.18.

# New features

[Issue #5154](https://redmine.postgresql.org/issues/5154) -  Added accessibility support in AlertifyJS.<br>
[Issue #5170](https://redmine.postgresql.org/issues/5170) -  Added Czech language support.<br>
[Issue #5179](https://redmine.postgresql.org/issues/5179) -  Added Python 3.8 support.<br>

# Housekeeping

[Issue #5088](https://redmine.postgresql.org/issues/5088) -  Improve code coverage and API test cases for the Event Trigger module.<br>
[Issue #5133](https://redmine.postgresql.org/issues/5133) -  Improvements in the UI for both default and dark themes.<br>
[Issue #5176](https://redmine.postgresql.org/issues/5176) -  Enhance logging by tracking stdout and stderr of subprocess when log level set to DEBUG.<br>
[Issue #5185](https://redmine.postgresql.org/issues/5185) -  Added option to override the class name of a label tag for select2 control.<br>

# Bug fixes

[Issue #4955](https://redmine.postgresql.org/issues/4955) -  Changed the color of selected and hovered item for Select2 dropdown.<br>
[Issue #4996](https://redmine.postgresql.org/issues/4996) -  Improve the style of the highlighted code after query execution for Dark mode.<br>
[Issue #5058](https://redmine.postgresql.org/issues/5058) -  Ensure that AlertifyJS should not be visible as a title for alert dialog.<br>
[Issue #5077](https://redmine.postgresql.org/issues/5077) -  Changed background pattern for geometry viewer to use #fff for all themes.<br>
[Issue #5101](https://redmine.postgresql.org/issues/5101) -  Fix an issue where debugger not showing all arguments anymore after hitting SQL error while debugging.<br>
[Issue #5107](https://redmine.postgresql.org/issues/5107) -  Set proper focus on tab navigation for file manager dialog.<br>
[Issue #5115](https://redmine.postgresql.org/issues/5115) -  Fix an issue where command and statements were parsed incorrectly for Rules.<br>
[Issue #5142](https://redmine.postgresql.org/issues/5142) -  Ensure that all the transactions should be canceled before closing the connections when a server is disconnected using pgAdmin.<br>
[Issue #5184](https://redmine.postgresql.org/issues/5184) -  Fixed Firefox monospaced issue by updating the font to the latest version.<br>
[Issue #5214](https://redmine.postgresql.org/issues/5214) -  Update Flask-SQLAlchemy and SQLAlchemy package which is not working on Windows with Python 3.8.<br>
[Issue #5215](https://redmine.postgresql.org/issues/5215) -  Fix syntax error when changing the event type for the existing rule.<br>
