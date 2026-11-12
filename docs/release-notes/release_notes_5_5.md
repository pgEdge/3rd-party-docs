# Version 5.5

Release date: 2021-07-15

This release contains a number of bug fixes and new features since the release of pgAdmin4 5.4.

# New features

[Issue #1975](https://redmine.postgresql.org/issues/1975) -  Highlighted long running queries on the dashboards.<br>
[Issue #3893](https://redmine.postgresql.org/issues/3893) -  Added support for Reassign/Drop Owned for login roles.<br>
[Issue #3920](https://redmine.postgresql.org/issues/3920) -  Do not block the query editor window when running a query.<br>
[Issue #5940](https://redmine.postgresql.org/issues/5940) -  Added support for OAuth 2 authentication.<br>
[Issue #6559](https://redmine.postgresql.org/issues/6559) -  Added option to provide maximum width of the column when 'Resize by data?’ option in the preferences is set to True.<br>

# Housekeeping

# Bug fixes

[Issue #4189](https://redmine.postgresql.org/issues/4189) -  Ensure that the Data Output panel can be snapped back after it is detached.<br>
[Issue #6388](https://redmine.postgresql.org/issues/6388) -  Fixed replace keyboard shortcut issue in the query tool on the normal keyboard layout.<br>
[Issue #6398](https://redmine.postgresql.org/issues/6398) -  Fixed an issue where detaching the query editor panel gives a blank white panel.<br>
[Issue #6427](https://redmine.postgresql.org/issues/6427) -  Remove leading whitespace and replace it with '[...] ' in the Query Tool data grid so cells don't look empty.<br>
[Issue #6448](https://redmine.postgresql.org/issues/6448) -  Fixed an issue in the search object when searching in 'all types' or 'subscription' if the user doesn't have access to the subscription.<br>
[Issue #6489](https://redmine.postgresql.org/issues/6489) -  Fixed an issue where Execute/Refresh button should not be disabled when we run the empty query.<br>
[Issue #6505](https://redmine.postgresql.org/issues/6505) -  Fixed an issue where the New Connection Drop Down has lost default maintenance database, auto-select, and tab-through functionality.<br>
[Issue #6536](https://redmine.postgresql.org/issues/6536) -  Fixed directory selection issue with the folder dialog.<br>
[Issue #6541](https://redmine.postgresql.org/issues/6541) -  Ensure that setting 'Open in new browser tab' should be visible, it should not be based on the value of 'ENABLE_PSQL'.<br>
[Issue #6547](https://redmine.postgresql.org/issues/6547) -  Fixed copy/paste issues for PSQL tool terminal.<br>
[Issue #6550](https://redmine.postgresql.org/issues/6550) -  Disable email deliverability check that was introduced in flask-security-too by default to maintain backwards compatibility.<br>
[Issue #6555](https://redmine.postgresql.org/issues/6555) -  Fixed Czech translation string for 'Login' keyword.<br>
[Issue #6557](https://redmine.postgresql.org/issues/6557) -  Fixed an issue where incorrect column name listed in the properties of Index.<br>
