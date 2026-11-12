# Version 9.13

Release date: 2026-03-05

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.12.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16, 17 and 18

**EDB Advanced Server**: 13, 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.0

# New features

[Issue #229](https://github.com/pgadmin-org/pgadmin4/issues/229)   - Allow users to customize "OF TYPE" columns during table creation.<br>
[Issue #5578](https://github.com/pgadmin-org/pgadmin4/issues/5578) - No FK are shown in diagram created from existing tables in the ERD Tool.<br>
[Issue #6386](https://github.com/pgadmin-org/pgadmin4/issues/6386) - Add support for 'ONLY' in Index creation dialog.<br>
[Issue #8198](https://github.com/pgadmin-org/pgadmin4/issues/8198) - Allow "drag-n-drop" for only user chosen tables, and show relations between them.<br>
[Issue #9229](https://github.com/pgadmin-org/pgadmin4/issues/9229) - Load predefined users from a JSON file through command line.<br>
[Issue #9641](https://github.com/pgadmin-org/pgadmin4/issues/9641) - Core LLM integration infrastructure, AI reports for security, schema, and performance, AI chat for the Query Tool, and AI Insights for EXPLAIN.<br>

# Housekeeping

# Bug fixes

[Issue #7578](https://github.com/pgadmin-org/pgadmin4/issues/7578) - Fixed an issue where the 'Quote strings only' configuration was ignored when downloading the result set.<br>
[Issue #8988](https://github.com/pgadmin-org/pgadmin4/issues/8988) - Fixed an issue where tools settings changed by the users were not restored on application relaunch.<br>
[Issue #9258](https://github.com/pgadmin-org/pgadmin4/issues/9258) - Fixed an issue where modifying a shared server incorrectly updated the original server details.<br>
[Issue #9484](https://github.com/pgadmin-org/pgadmin4/issues/9484) - Fixed an issue where a long name in ERD table node was not breaking into multiple lines.<br>
[Issue #9486](https://github.com/pgadmin-org/pgadmin4/issues/9486) - Fixed an issue where column comments were not displayed in the SQL tab for materialised views.<br>
[Issue #9572](https://github.com/pgadmin-org/pgadmin4/issues/9572) - Fix an issue where deployment of helm chart crashing with operation not permitted.<br>
[Issue #9583](https://github.com/pgadmin-org/pgadmin4/issues/9583) - Fix translation compilation.<br>
[Issue #9649](https://github.com/pgadmin-org/pgadmin4/issues/9649) - Fix broken checkbox selection in backup dialog objects tree.<br>
[Issue #9651](https://github.com/pgadmin-org/pgadmin4/issues/9651) - Fixed an issue in file dialog where rename was not working.<br>
