# Version 4.16

Release date: 2019-12-12

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.15.

!!! warning

    Warning: This release includes a change to the container distribution to run pgAdmin as a non-root user. Those users of the container who are running with mapped storage directories may need to change the ownership on the host machine, for example:

    ```bash
     sudo chown -R 5050:5050 <host_directory>
    ```

# New features

[Issue #4396](https://redmine.postgresql.org/issues/4396) -  Warn the user on changing the definition of Materialized View about the loss of data and its dependent objects.<br>
[Issue #4435](https://redmine.postgresql.org/issues/4435) -  Allow drag and drop functionality for all the nodes under the database node, excluding collection nodes.<br>
[Issue #4711](https://redmine.postgresql.org/issues/4711) -  Use a 'play' icon for the Execute Query button in the Query Tool for greater consistency with other applications.<br>
[Issue #4772](https://redmine.postgresql.org/issues/4772) -  Added aria-label to provide an invisible label where a visible label cannot be used.<br>
[Issue #4773](https://redmine.postgresql.org/issues/4773) -  Added role="status" attribute to all the status messages for accessibility.<br>
[Issue #4939](https://redmine.postgresql.org/issues/4939) -  Run pgAdmin in the container as a non-root user (pgadmin, UID: 5050)<br>
[Issue #4944](https://redmine.postgresql.org/issues/4944) -  Allow Gunicorn logs in the container to be directed to a file specified through GUNICORN_ACCESS_LOGFILE.<br>
[Issue #4990](https://redmine.postgresql.org/issues/4990) -  Changed the open query tool and data filter icons.<br>

# Housekeeping

[Issue #4696](https://redmine.postgresql.org/issues/4696) -  Add Reverse Engineered and Modified SQL tests for Materialized Views.<br>
[Issue #4807](https://redmine.postgresql.org/issues/4807) -  Refactored code of table and it's child nodes.<br>
[Issue #4938](https://redmine.postgresql.org/issues/4938) -  Refactored code of columns node.<br>

# Bug fixes

[Issue #3538](https://redmine.postgresql.org/issues/3538) -  Fix issue where the Reset button does not get enabled till all the mandatory fields are provided in the dialog.<br>
[Issue #4220](https://redmine.postgresql.org/issues/4220) -  Fix scrolling issue in 'Users' dialog.<br>
[Issue #4516](https://redmine.postgresql.org/issues/4516) -  Remove the sorting of table headers with no labels.<br>
[Issue #4659](https://redmine.postgresql.org/issues/4659) -  Updated documentation for default privileges to clarify more on the grantor.<br>
[Issue #4674](https://redmine.postgresql.org/issues/4674) -  Fix query tool launch error if user name contains HTML characters. It's a regression.<br>
[Issue #4724](https://redmine.postgresql.org/issues/4724) -  Fix network disconnect issue while establishing the connection via SSH Tunnel and it impossible to expand the Servers node.<br>
[Issue #4761](https://redmine.postgresql.org/issues/4761) -  Fix an issue where the wrong type is displayed when changing the datatype from timestamp with time zone to timestamp without time zone.<br>
[Issue #4792](https://redmine.postgresql.org/issues/4792) -  Ensure that the superuser should be able to create database, as the superuser overrides all the access restrictions.<br>
[Issue #4818](https://redmine.postgresql.org/issues/4818) -  Fix server connection drops out issue in query tool.<br>
[Issue #4836](https://redmine.postgresql.org/issues/4836) -  Updated the json file name from 'servers.json' to 'pgadmin4/servers.json' in the container deployment section of the documentation.<br>
[Issue #4878](https://redmine.postgresql.org/issues/4878) -  Ensure that the superuser should be able to create role, as the superuser overrides all the access restrictions.<br>
[Issue #4893](https://redmine.postgresql.org/issues/4893) -  Fix reverse engineering SQL issue for partitions when specifying digits as comments.<br>
[Issue #4923](https://redmine.postgresql.org/issues/4923) -  Enhance the logic to change the label from 'Delete/Drop' to 'Remove' for the server and server group node.<br>
[Issue #4925](https://redmine.postgresql.org/issues/4925) -  Shown some text on process watcher till the initial logs are loaded.<br>
[Issue #4926](https://redmine.postgresql.org/issues/4926) -  Fix VPN network disconnect issue where pgAdmin4 hangs on expanding the Servers node.<br>
[Issue #4930](https://redmine.postgresql.org/issues/4930) -  Fix main window tab navigation accessibility issue.<br>
[Issue #4933](https://redmine.postgresql.org/issues/4933) -  Ensure that the Servers collection node should expand independently of server connections.<br>
[Issue #4934](https://redmine.postgresql.org/issues/4934) -  Fix the help button link on the User Management dialog.<br>
[Issue #4935](https://redmine.postgresql.org/issues/4935) -  Fix accessibility issues.<br>
[Issue #4947](https://redmine.postgresql.org/issues/4947) -  Fix XSS issue in explain and explain analyze for table and type which contain HTML.<br>
[Issue #4952](https://redmine.postgresql.org/issues/4952) -  Fix an issue of retrieving properties for Compound Triggers. It's a regression of #4006.<br>
[Issue #4953](https://redmine.postgresql.org/issues/4953) -  Fix an issue where pgAdmin4 unable to retrieve table node if the trigger is already disabled and the user clicks on Enable All.<br>
[Issue #4958](https://redmine.postgresql.org/issues/4958) -  Fix reverse engineering SQL issue for triggers when passed a single argument to trigger function.<br>
[Issue #4964](https://redmine.postgresql.org/issues/4964) -  Fix an issue where length and precision are not removed from table/column dialog.<br>
[Issue #4965](https://redmine.postgresql.org/issues/4965) -  Fix an issue where the Interval data type is not displayed in the properties dialog of table/column.<br>
[Issue #4966](https://redmine.postgresql.org/issues/4966) -  Fix 'Could not find the object on the server.' error while refreshing the check constraint.<br>
[Issue #4975](https://redmine.postgresql.org/issues/4975) -  Fix issue where the user can not switch the UI language. It's a regression of #4348.<br>
[Issue #4976](https://redmine.postgresql.org/issues/4976) -  Fix reverse engineering SQL issue where when clause is not visible for PG/EPAS 12.<br>
[Issue #4978](https://redmine.postgresql.org/issues/4978) -  Fix pgAdmin4 failed to start issue after upgrading to version 4.15.<br>
[Issue #4982](https://redmine.postgresql.org/issues/4982) -  Added statistics and storage information in reverse engineering SQL of table/column.<br>
[Issue #4985](https://redmine.postgresql.org/issues/4985) -  Fix an issue where the inherited table name with quotes did not escape correctly.<br>
[Issue #4991](https://redmine.postgresql.org/issues/4991) -  Fix an issue where context menu is open along with submenu and the focus is not on context menu or submenu.<br>
