# Version 8.7

Release date: 2024-05-30

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.6.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, and 16

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.3

# New features

[Issue #4735](https://github.com/pgadmin-org/pgadmin4/issues/4735) -  Added support for the 'Add to macros' feature and fixed various usability issues.<br>
[Issue #6841](https://github.com/pgadmin-org/pgadmin4/issues/6841) -  Added support for executing the query at the cursor position in the query tool.<br>
[Issue #7192](https://github.com/pgadmin-org/pgadmin4/issues/7192) -  Changes in Query Tool, Debugger, and ERD Tool shortcuts to remove the use of Accesskey which will allow them to be customized.<br>
[Issue #7215](https://github.com/pgadmin-org/pgadmin4/issues/7215) -  Added support for viewing PGD Clusters.<br>
[Issue #7411](https://github.com/pgadmin-org/pgadmin4/issues/7411) -  Enhance the Delete dialog by highlighting the names of the objects to be deleted in bold.<br>
[Issue #7437](https://github.com/pgadmin-org/pgadmin4/issues/7437) -  Added support for platform Ubuntu 24.04 and Fedora 40.<br>

# Housekeeping

[Issue #7419](https://github.com/pgadmin-org/pgadmin4/issues/7419) -  Upgrade react-table from v7 to v8.<br>
[Issue #7472](https://github.com/pgadmin-org/pgadmin4/issues/7472) -  Replace the current FontAwesome based PSQL tool icon with MUI Terminal icon.<br>

# Bug fixes

[Issue #5762](https://github.com/pgadmin-org/pgadmin4/issues/5762) -  Ensure that Schema Diff does not indicate a table as different when the trigger names are the same but the trigger function body is different.<br>
[Issue #5849](https://github.com/pgadmin-org/pgadmin4/issues/5849) -  Disable ERD for system catalogs.<br>
[Issue #6060](https://github.com/pgadmin-org/pgadmin4/issues/6060) -  Disable Debugger for system catalogs.<br>
[Issue #6086](https://github.com/pgadmin-org/pgadmin4/issues/6086) -  Fixed an issue where drag and drop publication and subscription name in SQL editors was not working.<br>
[Issue #6464](https://github.com/pgadmin-org/pgadmin4/issues/6464) -  Fixed an issue of the pgAdmin window size increasing each time it was reopened.<br>
[Issue #7349](https://github.com/pgadmin-org/pgadmin4/issues/7349) -  Update the documentation for preferences dialog and keyboard shortcuts.<br>
[Issue #7439](https://github.com/pgadmin-org/pgadmin4/issues/7439) -  Fixed an issue where pgAdmin fails to start when Ubuntu OS is upgraded to a major version.<br>
[Issue #7458](https://github.com/pgadmin-org/pgadmin4/issues/7458) -  Remove query info notifier timeout field from Query Tool Preferences Dialog.<br>
[Issue #7485](https://github.com/pgadmin-org/pgadmin4/issues/7485) -  Fixed incorrect highlighting for C-Style escape strings in SQL editor.<br>
[Issue #7487](https://github.com/pgadmin-org/pgadmin4/issues/7487) -  Fixed an issue where the recover password button was enabled even when no email was provided.<br>
[Issue #7500](https://github.com/pgadmin-org/pgadmin4/issues/7500) -  Fixed an issue where resetting the password from the password reset link was not working.<br>
