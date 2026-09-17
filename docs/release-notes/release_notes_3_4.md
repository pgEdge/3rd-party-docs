# Version 3.4

Release date: 2018-10-04

This release contains a number of features and fixes reported since the release of pgAdmin4 3.3

# Features

[Issue #2927](https://redmine.postgresql.org/issues/2927) - Move all CSS into SCSS files for consistency and ease of colour maintenance etc.<br>
[Issue #3514](https://redmine.postgresql.org/issues/3514) - Add optional data point markers and mouse-over tooltips to display values on graphs.<br>
[Issue #3564](https://redmine.postgresql.org/issues/3564) - Add shortcuts for View Data and the Query tool to the Browser header bar.<br>

# Bug fixes

[Issue #3464](https://redmine.postgresql.org/issues/3464) - Ensure the runtime can startup properly if there are wide characters in the logfile path on Windows.<br>
[Issue #3551](https://redmine.postgresql.org/issues/3551) - Fix handling of backslashes in the edit grid.<br>
[Issue #3576](https://redmine.postgresql.org/issues/3576) - Ensure queries are no longer executed when dashboards are closed.<br>
[Issue #3596](https://redmine.postgresql.org/issues/3596) - Fix support for the CLOB datatype in EPAS.<br>
[Issue #3607](https://redmine.postgresql.org/issues/3607) - Fix logic around validation and highlighting of Sort/Filter in the Query Tool.<br>
[Issue #3630](https://redmine.postgresql.org/issues/3630) - Ensure auto-complete works for objects in schemas other than public and pg_catalog.<br>
[Issue #3657](https://redmine.postgresql.org/issues/3657) - Ensure changes to Query Tool settings from the Preferences dialogue are applied before executing queries.<br>
[Issue #3658](https://redmine.postgresql.org/issues/3658) - Swap the Schema and Schemas icons and Catalog and Catalogs icons that had been used the wrong way around.<br>
