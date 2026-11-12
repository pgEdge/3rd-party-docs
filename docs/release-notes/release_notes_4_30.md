# Version 4.30

Release date: 2021-01-28

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.29.

# New features

[Issue #1802](https://redmine.postgresql.org/issues/1802) -  Added ERD Diagram support with basic table fields, primary key, foreign key, and DDL SQL generation.<br>
[Issue #5457](https://redmine.postgresql.org/issues/5457) -  Added support for Kerberos authentication, using SPNEGO to forward the Kerberos tickets through a browser.<br>
[Issue #6147](https://redmine.postgresql.org/issues/6147) -  Documentation of Kerberos support.<br>
[Issue #6152](https://redmine.postgresql.org/issues/6152) -  Documentation of ERD Diagram support.<br>
[Issue #6160](https://redmine.postgresql.org/issues/6160) -  Add a container option (PGADMIN_DISABLE_POSTFIX) to disable the Postfix server.<br>

# Housekeeping

[Issue #5338](https://redmine.postgresql.org/issues/5338) -  Improve code coverage and API test cases for pgAgent.<br>
[Issue #6052](https://redmine.postgresql.org/issues/6052) -  Added connected pgAdmin user and connection name in the log file.<br>
[Issue #6079](https://redmine.postgresql.org/issues/6079) -  Updated mimetype from 'text/javascript' to 'application/javascript' as 'text/javascript' is obsolete.<br>
[Issue #6162](https://redmine.postgresql.org/issues/6162) -  Include PostgreSQL 13 utilities in the container.<br>

# Bug fixes

[Issue #5282](https://redmine.postgresql.org/issues/5282) -  Added 'Count Rows' option to the partition sub tables.<br>
[Issue #5488](https://redmine.postgresql.org/issues/5488) -  Improve the explain plan details by showing popup instead of tooltip on clicking of the specified node.<br>
[Issue #5571](https://redmine.postgresql.org/issues/5571) -  Added support for expression in exclusion constraints.<br>
[Issue #5829](https://redmine.postgresql.org/issues/5829) -  Fixed incorrect log information for AUTHENTICATION_SOURCES.<br>
[Issue #5875](https://redmine.postgresql.org/issues/5875) -  Ensure that the 'template1' database should not be visible after pg_upgrade.<br>
[Issue #5905](https://redmine.postgresql.org/issues/5905) -  Fixed an issue where the Save button is enabled by default in Macro.<br>
[Issue #5906](https://redmine.postgresql.org/issues/5906) -  Remove extra line after Manage Macros menu while clearing all macros.<br>
[Issue #5907](https://redmine.postgresql.org/issues/5907) -  Ensure that 'Clear All Rows' should not work if there is no existing macro available and the user does not specify any value.<br>
[Issue #5929](https://redmine.postgresql.org/issues/5929) -  Fixed an issue where the server is disconnected error message displayed if the user creates Macro with invalid SQL.<br>
[Issue #5965](https://redmine.postgresql.org/issues/5965) -  Ensure that the macro query result should be download properly.<br>
[Issue #5973](https://redmine.postgresql.org/issues/5973) -  Added appropriate help message and a placeholder for letting users know about the account password expiry for Login/Group Role.<br>
[Issue #5997](https://redmine.postgresql.org/issues/5997) -  Updated Flask-BabelEx to the latest.<br>
[Issue #6046](https://redmine.postgresql.org/issues/6046) -  Fixed an issue where the state of the Save File icon does not match the dirty editor indicator.<br>
[Issue #6047](https://redmine.postgresql.org/issues/6047) -  Fixed an issue where the dirty indicator stays active even if all changes were undone.<br>
[Issue #6058](https://redmine.postgresql.org/issues/6058) -  Ensure that the rename panel should be disabled when the SQL file opened in the query tool.<br>
[Issue #6061](https://redmine.postgresql.org/issues/6061) -  Fixed extra parentheses issue around joins for Views.<br>
[Issue #6065](https://redmine.postgresql.org/issues/6065) -  Fixed accessibility issues in schema diff module.<br>
[Issue #6069](https://redmine.postgresql.org/issues/6069) -  Fixed an issue on refreshing files in Query Tool.<br>
[Issue #6075](https://redmine.postgresql.org/issues/6075) -  Fixed an issue where Non-admin user is unable to view shared server created using service.<br>
[Issue #6077](https://redmine.postgresql.org/issues/6077) -  Fixed accessibility issues in various dialogs.<br>
[Issue #6084](https://redmine.postgresql.org/issues/6084) -  Fixed TypeError exception in schema diff when selected any identical object.<br>
[Issue #6096](https://redmine.postgresql.org/issues/6096) -  Updated deployment documentation, refer correctly to uWSGI where Gunicorn had been referenced.<br>
[Issue #6098](https://redmine.postgresql.org/issues/6098) -  Fixed an issue of deleting records when the user tries to delete multiple records.<br>
[Issue #6120](https://redmine.postgresql.org/issues/6120) -  Ensure that the user should be able to specify an older date for the account expiration of the role/user.<br>
[Issue #6121](https://redmine.postgresql.org/issues/6121) -  Fixed an issue where the database list in the new connection window is not visible.<br>
[Issue #6122](https://redmine.postgresql.org/issues/6122) -  Added informative message when there is no difference found for schema diff.<br>
[Issue #6128](https://redmine.postgresql.org/issues/6128) -  Fixed an issue where sequences are not created.<br>
[Issue #6140](https://redmine.postgresql.org/issues/6140) -  Ensure that verbose logs should be visible for Utility(Backup, Maintenance) jobs.<br>
[Issue #6144](https://redmine.postgresql.org/issues/6144) -  Ensure that the current value of the sequence should be ignored while comparing using schema diff.<br>
