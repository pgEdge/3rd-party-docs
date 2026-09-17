# Version 6.11

Release date: 2022-06-30

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.10.

# New features

[Issue #2647](https://redmine.postgresql.org/issues/2647) -  Added mouse over indication for breakpoint area in the Debugger.<br>
[Issue #2648](https://redmine.postgresql.org/issues/2648) -  Added search text option to the Debugger panel.<br>
[Issue #7178](https://redmine.postgresql.org/issues/7178) -  Added capability to deploy PostgreSQL servers on Microsoft Azure.<br>
[Issue #7332](https://redmine.postgresql.org/issues/7332) -  Added support for passing password using Docker Secret to Docker images.<br>
[Issue #7351](https://redmine.postgresql.org/issues/7351) -  Added the option 'Show template databases?' to display template databases regardless of the setting of 'Show system objects?'.<br>
[Issue #7485](https://redmine.postgresql.org/issues/7485) -  Added support for visualise the graph using a Line chart in the query tool.<br>

# Housekeeping

[Issue #6132](https://redmine.postgresql.org/issues/6132) -  Port Debugger to React.<br>
[Issue #7315](https://redmine.postgresql.org/issues/7315) -  Updates documentation for the Traefik v2 container deployment.<br>
[Issue #7411](https://redmine.postgresql.org/issues/7411) -  Update pgcli to latest release 3.4.1.<br>
[Issue #7469](https://redmine.postgresql.org/issues/7469) -  Upgrade Chartjs to the latest 3.8.0.<br>

# Bug fixes

[Issue #7423](https://redmine.postgresql.org/issues/7423) -  Fixed an issue where there is no setting to turn off notifications in the Query Tool.<br>
[Issue #7440](https://redmine.postgresql.org/issues/7440) -  Fixed an issue where passwords entered in the 'Connect To Server' dialog were truncated.<br>
[Issue #7441](https://redmine.postgresql.org/issues/7441) -  Ensure that the Query Editor should be focused when switching between query tool tabs.<br>
[Issue #7443](https://redmine.postgresql.org/issues/7443) -  Fixed and issue where 'Use spaces' not working in the query tool.<br>
[Issue #7453](https://redmine.postgresql.org/issues/7453) -  Fixed an issue where the Database restriction is not working.<br>
[Issue #7460](https://redmine.postgresql.org/issues/7460) -  Fixed an issue where pgAdmin stuck while creating a new index.<br>
[Issue #7461](https://redmine.postgresql.org/issues/7461) -  Fixed an issue where the connection wasn't being closed when the user switched to a new connection and closed the query tool.<br>
[Issue #7468](https://redmine.postgresql.org/issues/7468) -  Skip the history records if the JSON info can't be parsed instead of showing 'No history'.<br>
[Issue #7502](https://redmine.postgresql.org/issues/7502) -  Fixed an issue where an error message is displayed when creating the new database.<br>
[Issue #7506](https://redmine.postgresql.org/issues/7506) -  Fixed permission denied error when deploying PostgreSQL in Azure using Docker.<br>
