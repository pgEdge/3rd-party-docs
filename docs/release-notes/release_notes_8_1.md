# Version 8.1

Release date: 2023-12-14

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v8.0.

# Supported Database Servers

**PostgreSQL**: 12, 13, 14, 15, and 16

**EDB Advanced Server**: 12, 13, 14, 15, and 16

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 16.0

# New features

[Issue #4580](https://github.com/pgadmin-org/pgadmin4/issues/4580) -  Add support for generating ERD for a schema.<br>
[Issue #6854](https://github.com/pgadmin-org/pgadmin4/issues/6854) -  Add support for creating a function with custom return type.<br>

# Housekeeping

[Issue #6991](https://github.com/pgadmin-org/pgadmin4/issues/6991) -  Fixed several accessibility-related issues for enhanced usability.<br>

# Bug fixes

[Issue #5471](https://github.com/pgadmin-org/pgadmin4/issues/5471) -  Ensure focus is not changed to ssh tunnel password input when user explicitly focus on server password input.<br>
[Issue #6095](https://github.com/pgadmin-org/pgadmin4/issues/6095) -  Provide a way to bypass the SSL cert verification for OAuth2 provider.<br>
[Issue #6488](https://github.com/pgadmin-org/pgadmin4/issues/6488) -  Fixed an issue where database name was missing in an error message if name contains any special characters.<br>
[Issue #6717](https://github.com/pgadmin-org/pgadmin4/issues/6717) -  Ensure that indexes created by constraints are visible in the object explorer when "Show system objects" is enabled.<br>
[Issue #6803](https://github.com/pgadmin-org/pgadmin4/issues/6803) -  Fixed an issue where reading process logs throws an error when DATA_DIR is moved to a networked drive.<br>
[Issue #6814](https://github.com/pgadmin-org/pgadmin4/issues/6814) -  Remove the 'Close Window' submenu specifically for OSX to prevent unintended closure of the entire application.<br>
[Issue #6842](https://github.com/pgadmin-org/pgadmin4/issues/6842) -  Rename all references of 'Execute query' to 'Execute script' to be more relevant.<br>
[Issue #6887](https://github.com/pgadmin-org/pgadmin4/issues/6887) -  Fixed an issue where syntax error was not highlighting in query tool.<br>
[Issue #6921](https://github.com/pgadmin-org/pgadmin4/issues/6921) -  Fixed an issue where on entering full screen, the option label is not changed to 'Exit Full Screen' in desktop mode.<br>
[Issue #6950](https://github.com/pgadmin-org/pgadmin4/issues/6950) -  Ensure that the Authentication Source in the drop-down of the UserManagement dialog aligns with the entries specified for AUTHENTICATION_SOURCES in the configuration file.<br>
[Issue #6958](https://github.com/pgadmin-org/pgadmin4/issues/6958) -  Reverse engineer serial columns when generating ERD for database/table.<br>
[Issue #6964](https://github.com/pgadmin-org/pgadmin4/issues/6964) -  Fixed an issue where the Schema was not visible in the dropdown for table properties or when creating a new table.<br>
[Issue #6968](https://github.com/pgadmin-org/pgadmin4/issues/6968) -  Fixed an issue where option key was not registering in PSQL tool.<br>
[Issue #6984](https://github.com/pgadmin-org/pgadmin4/issues/6984) -  Fixed an issue where the Vacuum option INDEX_CLEANUP have an incorrect value ('AUTO') for database versions < 14.<br>
[Issue #6989](https://github.com/pgadmin-org/pgadmin4/issues/6989) -  Fixed an issue where the pgAdmin page went blank when clicking the delete button in the User Management dialog.<br>
[Issue #7000](https://github.com/pgadmin-org/pgadmin4/issues/7000) -  Ensure that correct timezone is set for Docker deployments.<br>
[Issue #7011](https://github.com/pgadmin-org/pgadmin4/issues/7011) -  Fixed an issue where all rows and filter rows buttons of object explorer toolbar were disabled for views and other supported nodes.<br>
[Issue #7017](https://github.com/pgadmin-org/pgadmin4/issues/7017) -  Fixed an issue where schema diff tool is not loading preferences on start.<br>
