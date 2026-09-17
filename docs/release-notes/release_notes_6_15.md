# Version 6.15

Release date: 2022-10-20

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.14.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13, 14 and 15

**EDB Advanced Server**: 10, 11, 12, 13, 14 and 15

# New features

[Issue #3491](https://github.com/pgadmin-org/pgadmin4/issues/3491) -  Added support for IAM token based authentication for AWS RDS or Azure DB.<br>
[Issue #4392](https://github.com/pgadmin-org/pgadmin4/issues/4392) -  Added support to specify the background fill color to the table nodes in the ERD tool.<br>
[Issue #4994](https://github.com/pgadmin-org/pgadmin4/issues/4994) -  Allow reordering table columns using drag and drop in ERD Tool.<br>
[Issue #4997](https://github.com/pgadmin-org/pgadmin4/issues/4997) -  Add option to generate SQL with DROP table DDL in ERD Tool.<br>
[Issue #5304](https://github.com/pgadmin-org/pgadmin4/issues/5304) -  Added high availability options to AWS deployment.<br>
[Issue #5390](https://github.com/pgadmin-org/pgadmin4/issues/5390) -  Expose the Gunicorn limit_request_line parameter in the container, with the default set to the maximum 8190.<br>

# Housekeeping

[Issue #5065](https://github.com/pgadmin-org/pgadmin4/issues/5065) -  Use SocketIO instead of REST for fetching database tables data in ERD.<br>
[Issue #5293](https://github.com/pgadmin-org/pgadmin4/issues/5293) -  Ensure that the tooltips are consistent throughout the entire application.<br>
[Issue #5357](https://github.com/pgadmin-org/pgadmin4/issues/5357) -  Remove Python's 'Six' package completely.<br>

# Bug fixes

[Issue #5101](https://github.com/pgadmin-org/pgadmin4/issues/5101) -  Ensure consistent orderings for ACLS when comparing schemas in the schema diff.<br>
[Issue #5132](https://github.com/pgadmin-org/pgadmin4/issues/5132) -  Ensure that the result grid column should take width as pre preferences setting on first execution.<br>
[Issue #5133](https://github.com/pgadmin-org/pgadmin4/issues/5133) -  Fixed an exception occur while taking backup and SSL certificates/keys are not found in the specified path.<br>
[Issue #5145](https://github.com/pgadmin-org/pgadmin4/issues/5145) -  Fixed intermittent error shown while OAuth2 login.<br>
[Issue #5167](https://github.com/pgadmin-org/pgadmin4/issues/5167) -  Ensure that the path to the psqlrc file is correct when multiple users open the PSQL tool at the same time.<br>
[Issue #5188](https://github.com/pgadmin-org/pgadmin4/issues/5188) -  Ensure that the continue/start button should be disabled if the user stops the Debugger for the procedures.<br>
[Issue #5210](https://github.com/pgadmin-org/pgadmin4/issues/5210) -  Ensure that the query tool creates a new tab with the appropriate user when pressing Alt+Shift+Q.<br>
[Issue #5212](https://github.com/pgadmin-org/pgadmin4/issues/5212) -  Added the close button for all the notifications of the notistack.<br>
[Issue #5249](https://github.com/pgadmin-org/pgadmin4/issues/5249) -  Added the ability to display the selected text from the query tool in the find/replace box.<br>
[Issue #5261](https://github.com/pgadmin-org/pgadmin4/issues/5261) -  Ensure that the search filter should be cleared when a new row is added to the user management.<br>
[Issue #5262](https://github.com/pgadmin-org/pgadmin4/issues/5262) -  Ensure that the user management dialog should not allow the same email addresses with different letter casings when creating users.<br>
[Issue #5277](https://github.com/pgadmin-org/pgadmin4/issues/5277) -  Fixed XSS vulnerability issues.<br>
[Issue #5296](https://github.com/pgadmin-org/pgadmin4/issues/5296) -  Ensure that the scroll position should be preserved for the result set in the query tool on tab change.<br>
[Issue #5308](https://github.com/pgadmin-org/pgadmin4/issues/5308) -  Ensure that the default value for a column should be used if it is made empty.<br>
[Issue #5327](https://github.com/pgadmin-org/pgadmin4/issues/5327) -  Fixed an issue where user was unable to select privileges in Safari.<br>
[Issue #5332](https://github.com/pgadmin-org/pgadmin4/issues/5332) -  Fixed console warning shown while updating database node from browser tree.<br>
[Issue #5338](https://github.com/pgadmin-org/pgadmin4/issues/5338) -  Fixed an issue where the prompt is not visible when clicking on the 'save results to file' button on the large data.<br>
[Issue #5352](https://github.com/pgadmin-org/pgadmin4/issues/5352) -  Fixed error occurring while LDAP authentication for a user with multiple email attributes.<br>
[Issue #5364](https://github.com/pgadmin-org/pgadmin4/issues/5364) -  Fixed an issue where notifications disappeared quickly.<br>
[Issue #5367](https://github.com/pgadmin-org/pgadmin4/issues/5367) -  Ensure that the correct value should be returned if an exception occurs while decoding the password.<br>
[Issue #5368](https://github.com/pgadmin-org/pgadmin4/issues/5368) -  Fixed the issue while downloading the file from the file manager.<br>
[Issue #5386](https://github.com/pgadmin-org/pgadmin4/issues/5386) -  Ensure that the login form is hidden if the authentication source is OAuth2 or Kerberos.<br>
[Issue #5397](https://github.com/pgadmin-org/pgadmin4/issues/5397) -  Fixed an issue where the password recovery link was not working.<br>
[Issue #5402](https://github.com/pgadmin-org/pgadmin4/issues/5402) -  Ensure that scroll bar on browser tree should be visible on windows resize.<br>
[Issue #5405](https://github.com/pgadmin-org/pgadmin4/issues/5405) -  Fixed the cross-site scripting vulnerability.<br>
[Issue #5427](https://github.com/pgadmin-org/pgadmin4/issues/5427) -  Fixed an issue where filtered rows were not working.<br>
