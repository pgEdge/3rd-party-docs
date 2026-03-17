# Version 9.4

Release date: 2025-05-29

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.3.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16 and 17

**EDB Advanced Server**: 13, 14, 15, 16 and 17

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 17.2

# New features

[Issue #3369](https://github.com/pgadmin-org/pgadmin4/issues/3369) -  Enabled large file downloads for desktop users within the query tool.<br>
[Issue #8583](https://github.com/pgadmin-org/pgadmin4/issues/8583) -  Add all missing options to the Import/Export Data functionality, and update the syntax of the COPY command to align with the latest standards.<br>
[Issue #8681](https://github.com/pgadmin-org/pgadmin4/issues/8681) -  Add support for exporting table data based on a custom query.<br>

# Housekeeping

# Bug fixes

[Issue #6510](https://github.com/pgadmin-org/pgadmin4/issues/6510) -  Fixed an issue where the result grid slowed down when any column contained a large amount of data.<br>
[Issue #6564](https://github.com/pgadmin-org/pgadmin4/issues/6564) -  Fix the issue where an error is displayed when a table is dropped while a query is running.<br>
[Issue #6968](https://github.com/pgadmin-org/pgadmin4/issues/6968) -  Fixed an issue where the options key was not working as expected in the PSQL tool.<br>
[Issue #7926](https://github.com/pgadmin-org/pgadmin4/issues/7926) -  Fixed an issue where correct error message not displayed when sql statement contains Arabic letters.<br>
[Issue #8595](https://github.com/pgadmin-org/pgadmin4/issues/8595) -  Enhance contrast for selected and hovered items in the Object Explorer to improve visibility and accessibility.<br>
[Issue #8607](https://github.com/pgadmin-org/pgadmin4/issues/8607) -  Fixed an issue where the query tool returns "cannot unpack non-iterable Response object" when running any query with a database name change.<br>
[Issue #8608](https://github.com/pgadmin-org/pgadmin4/issues/8608) -  Handle result grid data changes in View/Edit Data mode by automatically reconnecting to the server if a disconnection occurs.<br>
[Issue #8668](https://github.com/pgadmin-org/pgadmin4/issues/8668) -  Implement API fetch error display for select dropdown.<br>
[Issue #8711](https://github.com/pgadmin-org/pgadmin4/issues/8711) -  Fixed an issue where light theme briefly appears when pgAdmin loads or tools open, even when a dark or system UI theme is preferred.<br>
[Issue #8713](https://github.com/pgadmin-org/pgadmin4/issues/8713) -  Fixed issues related to column range selection using shift + click.<br>
[Issue #8760](https://github.com/pgadmin-org/pgadmin4/issues/8760) -  Fixed an issue where pgAdmin failed to focus when previously unfocused and then quit.<br>
