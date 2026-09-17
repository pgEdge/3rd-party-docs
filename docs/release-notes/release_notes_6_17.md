# Version 6.17

Release date: 2022-12-02

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.16.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13, 14 and 15

**EDB Advanced Server**: 10, 11, 12, 13, 14 and 15

# New features

# Housekeeping

[Issue #5147](https://github.com/pgadmin-org/pgadmin4/issues/5147) -  Update the BigAnimal API version to V2.<br>
[Issue #5493](https://github.com/pgadmin-org/pgadmin4/issues/5493) -  Remove all traces of Backbone and Underscore.<br>

# Bug fixes

[Issue #5488](https://github.com/pgadmin-org/pgadmin4/issues/5488) -  Fixed an issue where the wrong schema is displayed for a foreign key in the schema diff tool.<br>
[Issue #5495](https://github.com/pgadmin-org/pgadmin4/issues/5495) -  Ensure that the query history date format in Desktop mode matches the format of the locale of the pgadmin server.<br>
[Issue #5505](https://github.com/pgadmin-org/pgadmin4/issues/5505) -  Fixed an issue where the CSV file would not download if the CSV quote character length exceeded 1.<br>
[Issue #5513](https://github.com/pgadmin-org/pgadmin4/issues/5513) -  Ensure that DATA_DIR dependent folders/files are automatically created inside the specified DATA_DIR if they are not specified separately in the configuration file.<br>
[Issue #5539](https://github.com/pgadmin-org/pgadmin4/issues/5539) -  Improved error message to make it easier for users to understand.<br>
[Issue #5548](https://github.com/pgadmin-org/pgadmin4/issues/5548) -  Fixed an issue where editor position was wrong when editing data from result grid.<br>
[Issue #5575](https://github.com/pgadmin-org/pgadmin4/issues/5575) -  Ensure the query tool is launched successfully for the servers registered with the PostgreSQL service.<br>
[Issue #5593](https://github.com/pgadmin-org/pgadmin4/issues/5593) -  Ensure that only authorized and authenticated users can validate binary paths (CVE-2022-4223).<br>
