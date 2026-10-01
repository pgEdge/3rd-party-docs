# Version 4.22

Release date: 2020-05-28

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.21.

# New features

[Issue #5452](https://redmine.postgresql.org/issues/5452) -  Added connected pgAdmin user and connection name in the log file.<br>
[Issue #5489](https://redmine.postgresql.org/issues/5489) -  Show the startup log as well as the server log in the runtime's log viewer.<br>

# Housekeeping

[Issue #5255](https://redmine.postgresql.org/issues/5255) -  Implement Selenium Grid to run multiple tests across different browsers, operating systems, and machines in parallel.<br>
[Issue #5333](https://redmine.postgresql.org/issues/5333) -  Improve code coverage and API test cases for Indexes.<br>
[Issue #5334](https://redmine.postgresql.org/issues/5334) -  Improve code coverage and API test cases for the Rules module.<br>
[Issue #5335](https://redmine.postgresql.org/issues/5335) -  Improve code coverage and API test cases for Triggers and Compound Triggers.<br>
[Issue #5443](https://redmine.postgresql.org/issues/5443) -  Remove support for Python 2.<br>
[Issue #5444](https://redmine.postgresql.org/issues/5444) -  Cleanup Python detection in the runtime project file.<br>
[Issue #5455](https://redmine.postgresql.org/issues/5455) -  Refactor pgAdmin4.py so it can be imported and is a lot more readable.<br>
[Issue #5493](https://redmine.postgresql.org/issues/5493) -  Search object UI improvements.<br>
[Issue #5525](https://redmine.postgresql.org/issues/5525) -  Cleanup and refactor the macOS build scripts.<br>
[Issue #5552](https://redmine.postgresql.org/issues/5552) -  Update dependencies in the Docker container.<br>
[Issue #5553](https://redmine.postgresql.org/issues/5553) -  Remove PG 9.4 utilities from the Docker container as it's now out of support.<br>

# Bug fixes

[Issue #3694](https://redmine.postgresql.org/issues/3694) -  Gracefully informed the user that the database is already connected when they click on "Connect Database...".<br>
[Issue #4033](https://redmine.postgresql.org/issues/4033) -  Fixed an issue where clicking on the cross button of the alert box on the login page is not working.<br>
[Issue #4099](https://redmine.postgresql.org/issues/4099) -  Fixed the SQL help issue for EDB Postgres Advanced Server.<br>
[Issue #4223](https://redmine.postgresql.org/issues/4223) -  Ensure that maintenance job should be worked properly for indexes under a materialized view.<br>
[Issue #4279](https://redmine.postgresql.org/issues/4279) -  Ensure that file browse "home" button should point to $HOME rather than /.<br>
[Issue #4840](https://redmine.postgresql.org/issues/4840) -  Ensure that 'With OID' option should be disabled while taking backup of database server version 12 and above.<br>
[Issue #5001](https://redmine.postgresql.org/issues/5001) -  Fixed invalid literal issue when removing the connection limit for the existing role.<br>
[Issue #5052](https://redmine.postgresql.org/issues/5052) -  Fixed internal server error when clicking on Triggers -> 'Enable All' for partitions.<br>
[Issue #5398](https://redmine.postgresql.org/issues/5398) -  Fixed generated SQL issue for auto vacuum options.<br>
[Issue #5422](https://redmine.postgresql.org/issues/5422) -  Ensure that the dependencies tab shows correct information for Synonyms.<br>
[Issue #5434](https://redmine.postgresql.org/issues/5434) -  Fixed an issue where the newly added table is not alphabetically added to the tree.<br>
[Issue #5440](https://redmine.postgresql.org/issues/5440) -  Fixed list sorting issue in the schema diff tool.<br>
[Issue #5449](https://redmine.postgresql.org/issues/5449) -  Fixed an issue while comparing the two identical schemas using the schema diff tool.<br>
[Issue #5450](https://redmine.postgresql.org/issues/5450) -  Fixed an issue when renaming the column not added in the proper order.<br>
[Issue #5466](https://redmine.postgresql.org/issues/5466) -  Correct ipv4 "all interfaces" address in the container docs, per Frank Limpert.<br>
[Issue #5469](https://redmine.postgresql.org/issues/5469) -  Fixed an issue where select2 hover is inconsistent for the SSL field in create server dialog.<br>
[Issue #5473](https://redmine.postgresql.org/issues/5473) -  Fixed post-login redirect location when running in server mode under a non-default root.<br>
[Issue #5480](https://redmine.postgresql.org/issues/5480) -  Fixed an issue where the background job creation fails if there is only a version-specific python binary available in PATH.<br>
[Issue #5481](https://redmine.postgresql.org/issues/5481) -  Fixed data truncation issue when updating the data of type character with length.<br>
[Issue #5487](https://redmine.postgresql.org/issues/5487) -  Fixed an issue where if LDAP_SEARCH_BASE_DN is not set then, the value for LDAP_BASE_DN will be considered.<br>
[Issue #5496](https://redmine.postgresql.org/issues/5496) -  Fixed an issue where clicking on Select All button, not selecting all the options in pgAgent job scheduler.<br>
[Issue #5503](https://redmine.postgresql.org/issues/5503) -  Clarify and correct the docs on enabling the pl/debugger plugin on the server.<br>
[Issue #5510](https://redmine.postgresql.org/issues/5510) -  Fixed Unicode decode error 'utf-8' codec can't decode byte.<br>
