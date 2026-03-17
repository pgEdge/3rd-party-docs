# Version 6.3

Release date: 2021-12-16

This release contains a number of bug fixes and new features since the release of pgAdmin4 6.2.

# New features

[Issue #6543](https://redmine.postgresql.org/issues/6543) -  Added support for Two-factor authentication for improving security.<br>
[Issue #6872](https://redmine.postgresql.org/issues/6872) -  Include GSSAPI support in the PostgreSQL libraries and utilities on macOS.<br>
[Issue #7039](https://redmine.postgresql.org/issues/7039) -  Added support to disable the auto-discovery of the database servers.<br>

# Housekeeping

[Issue #6088](https://redmine.postgresql.org/issues/6088) -  Replace Flask-BabelEx with Flask-Babel.<br>
[Issue #6984](https://redmine.postgresql.org/issues/6984) -  Port Backup Global, Backup Server, and Backup object dialog in React.<br>
[Issue #7004](https://redmine.postgresql.org/issues/7004) -  Replaced alertifyjs notifiers with React-based notistack.<br>
[Issue #7010](https://redmine.postgresql.org/issues/7010) -  Upgrade Flask to version 2.<br>
[Issue #7053](https://redmine.postgresql.org/issues/7053) -  Replace Alertify alert and confirm with React-based model dialog.<br>

# Bug fixes

[Issue #6840](https://redmine.postgresql.org/issues/6840) -  Fixed an issue where tooltip data are not displaying on downloaded graphical explain plan.<br>
[Issue #6877](https://redmine.postgresql.org/issues/6877) -  Fixed schema diff owner related issue.<br>
[Issue #6906](https://redmine.postgresql.org/issues/6906) -  Fixed an issue where referenced table drop-down should be disabled in foreign key -> columns after one row is added.<br>
[Issue #6955](https://redmine.postgresql.org/issues/6955) -  Ensure that sort order should be maintained when renaming a server group.<br>
[Issue #6957](https://redmine.postgresql.org/issues/6957) -  Fixed schema diff related some issues.<br>
[Issue #6963](https://redmine.postgresql.org/issues/6963) -  Ensure that the user should be allowed to set the schema of an extension while creating it.<br>
[Issue #6978](https://redmine.postgresql.org/issues/6978) -  Increase the width of the scrollbars.<br>
[Issue #6986](https://redmine.postgresql.org/issues/6986) -  Fixed an issue where the user can't debug function with timestamp parameter.<br>
[Issue #6989](https://redmine.postgresql.org/issues/6989) -  Fixed an issue where the Change Password menu option is missing for internal authentication source when more than one authentication source is defined.<br>
[Issue #7005](https://redmine.postgresql.org/issues/7005) -  Fixed an issue where On-demand rows throw an error when any row cell is edited and saved it then scroll to get more rows.<br>
[Issue #7006](https://redmine.postgresql.org/issues/7006) -  Ensure that Python 3.10 and the latest eventlet dependency should not break the application.<br>
[Issue #7013](https://redmine.postgresql.org/issues/7013) -  Fix an RPM build issue that could lead to a conflict with python3 at installation.<br>
[Issue #7015](https://redmine.postgresql.org/issues/7015) -  Fixed an issue where the error is thrown while creating a new server using Add New Server from the dashboard while tree item is not selected.<br>
[Issue #7020](https://redmine.postgresql.org/issues/7020) -  Ensure that statue message should not hide the last line of messages when running a long query.<br>
[Issue #7024](https://redmine.postgresql.org/issues/7024) -  Fixed an issue where reverse engineering SQL is wrong for Aggregate.<br>
[Issue #7029](https://redmine.postgresql.org/issues/7029) -  Correct the SQL definition for function/procedure with the Atomic keyword in PG14.<br>
[Issue #7031](https://redmine.postgresql.org/issues/7031) -  Fixed an issue where SQLite database definition is wrong because the USER_ID FK references the table user_old which is not available.<br>
[Issue #7040](https://redmine.postgresql.org/issues/7040) -  Add "section" to the Debian package control files.<br>
[Issue #7044](https://redmine.postgresql.org/issues/7044) -  Update the dropzone version to 5.9.3 and Flask-SQLAlchemy to 2.5.*.<br>
[Issue #7046](https://redmine.postgresql.org/issues/7046) -  Fixed some accessibility issues.<br>
[Issue #7048](https://redmine.postgresql.org/issues/7048) -  Fixed unhashable type issue while opening the about dialog.<br>
[Issue #7064](https://redmine.postgresql.org/issues/7064) -  Ensure that the Owner should not be disabled while creating the procedure.<br>
[Issue #7071](https://redmine.postgresql.org/issues/7071) -  Fixed an issue where confirmation pop-up is hidden behind Reassign/Drop Owned Dialog.<br>
