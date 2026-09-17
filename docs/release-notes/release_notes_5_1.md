# Version 5.1

Release date: 2021-03-25

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.0.

# New features

[Issue #5404](https://redmine.postgresql.org/issues/5404) -  Show the login roles that are members of a group role be shown when examining a group role.<br>
[Issue #6212](https://redmine.postgresql.org/issues/6212) -  Make the container distribution a multi-arch build with x86_64 and Arm64 support.<br>
[Issue #6268](https://redmine.postgresql.org/issues/6268) -  Make 'kerberos' an optional feature in the Python wheel, to avoid the need to install MIT Kerberos on the system by default.<br>
[Issue #6270](https://redmine.postgresql.org/issues/6270) -  Added '--replace' option in Import server to replace the list of servers with the newly imported one.<br>
[Issue #6271](https://redmine.postgresql.org/issues/6271) -  Added zoom scaling options with keyboard shortcuts in runtime.<br>

# Housekeeping

[Issue #3976](https://redmine.postgresql.org/issues/3976) -  Use schema qualification while accessing the catalog objects.<br>
[Issue #6176](https://redmine.postgresql.org/issues/6176) -  Make the 'Save Data Changes' icon to be more intuitive.<br>

# Bug fixes

[Issue #4014](https://redmine.postgresql.org/issues/4014) -  Fixed alignment issue under preferences for the German language.<br>
[Issue #4020](https://redmine.postgresql.org/issues/4020) -  Fixed color issue on the statistics tab for collection node in the safari browser.<br>
[Issue #4438](https://redmine.postgresql.org/issues/4438) -  Fixed an issue where adding/updating records fails if the table name contains percent sign.<br>
[Issue #4784](https://redmine.postgresql.org/issues/4784) -  Ensure that autovacuum and analyze scale factors should be editable with more than two decimals.<br>
[Issue #4847](https://redmine.postgresql.org/issues/4847) -  Fixed an issue where % displayed twice in explain analyze for query and table.<br>
[Issue #4849](https://redmine.postgresql.org/issues/4849) -  Rename text 'table' with 'relation' in the statistic tab for explain analyze.<br>
[Issue #4959](https://redmine.postgresql.org/issues/4959) -  Fixed an issue where the properties tab for collection nodes is unresponsive after switching the tabs.<br>
[Issue #5073](https://redmine.postgresql.org/issues/5073) -  Fixed an issue where the Save button is enabled for functions/procedures by default when open the properties dialog.<br>
[Issue #5119](https://redmine.postgresql.org/issues/5119) -  Fixed an issue where hanging symlinks in a directory cause select file dialog to break.<br>
[Issue #5467](https://redmine.postgresql.org/issues/5467) -  Allow underscores in the Windows installation path.<br>
[Issue #5628](https://redmine.postgresql.org/issues/5628) -  Remove the "launch now" option in the Windows installer, as UAC could cause it to run as an elevated user.<br>
[Issue #5810](https://redmine.postgresql.org/issues/5810) -  Ensure that cell content being auto selected when editing the cell data.<br>
[Issue #5869](https://redmine.postgresql.org/issues/5869) -  Ensure that SQL formatter should not add extra tabs and format the SQL correctly.<br>
[Issue #6018](https://redmine.postgresql.org/issues/6018) -  Fixed encoding issue when database encoding set to SQL_ASCII and name of the column is in ASCII character.<br>
[Issue #6159](https://redmine.postgresql.org/issues/6159) -  Ensure that the user should be able to kill the session from Dashboard if the user has a 'pg_signal_backend' role.<br>
[Issue #6206](https://redmine.postgresql.org/issues/6206) -  Ensure that the view/edit data panel should not be opened for unsupported nodes using the keyboard shortcut.<br>
[Issue #6227](https://redmine.postgresql.org/issues/6227) -  Ensure PGADMIN_DEFAULT_EMAIL looks sane when initialising a container deployment.<br>
[Issue #6228](https://redmine.postgresql.org/issues/6228) -  Improve the web setup script for Linux to make the platform detection more robust and overrideable.<br>
[Issue #6233](https://redmine.postgresql.org/issues/6233) -  Ensure that SQL formatter should not use tab size if 'Use spaces?' set to false.<br>
[Issue #6253](https://redmine.postgresql.org/issues/6253) -  Fixed an issue where the user is unable to create a subscription if the host/IP address for connection is 127.0.0.1.<br>
[Issue #6259](https://redmine.postgresql.org/issues/6259) -  Ensure that proper error message should be shown on the properties and statistics tab in case of insufficient privileges for a subscription.<br>
[Issue #6260](https://redmine.postgresql.org/issues/6260) -  Fixed an issue where the 'Create Slot' option is disabled in case of the same IP/host provided but the port is different.<br>
[Issue #6269](https://redmine.postgresql.org/issues/6269) -  Ensure the Python interpreter used by the runtime ignores user site-packages.<br>
[Issue #6272](https://redmine.postgresql.org/issues/6272) -  Fixed an issue where the user is not able to change the connection in Query Tool when any SQL file is opened.<br>
[Issue #6279](https://redmine.postgresql.org/issues/6279) -  Ensure that the venv activation scripts have the correct path in them on Linux.<br>
[Issue #6281](https://redmine.postgresql.org/issues/6281) -  Fixed an issue where schema diff showing wrong SQL when comparing triggers with different when clause.<br>
[Issue #6286](https://redmine.postgresql.org/issues/6286) -  Ensure that the template database should be visible while creating the database.<br>
[Issue #6292](https://redmine.postgresql.org/issues/6292) -  Fixed string index out of range error where the dependent tab is in focus and selecting any publication or table.<br>
[Issue #6294](https://redmine.postgresql.org/issues/6294) -  Fixed an issue where the dependent tab throwing an error when selecting any login/group role.<br>
[Issue #6307](https://redmine.postgresql.org/issues/6307) -  Fixed an issue where the incorrect values visible in the dependents tab for publication.<br>
[Issue #6312](https://redmine.postgresql.org/issues/6312) -  Fixed an issue where copy/paste rows in view data paste the wrong value for boolean type.<br>
[Issue #6316](https://redmine.postgresql.org/issues/6316) -  Ensure that the primary key should be visible properly in the table dialog.<br>
[Issue #6317](https://redmine.postgresql.org/issues/6317) -  Ensure that toggle buttons are accessible by most screen readers.<br>
[Issue #6322](https://redmine.postgresql.org/issues/6322) -  Fixed an issue where the top menu disappears when entering into the full screen for minimum screen resolution.<br>
[Issue #6323](https://redmine.postgresql.org/issues/6323) -  Ensure that the grantor name should be visible properly for the security tab in the table dialog.<br>
