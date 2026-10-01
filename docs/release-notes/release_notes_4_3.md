# Version 4.3

Release date: 2019-03-07

This release contains a number of new features and fixes reported since the release of pgAdmin4 4.2

# Features

[Issue #1825](https://redmine.postgresql.org/issues/1825) - Install a script to start pgAdmin (pgadmin4) from the command line when installed from the Python wheel.<br>
[Issue #2233](https://redmine.postgresql.org/issues/2233) - Add a "scratch pad" to the Query Tool to hold text snippets whilst editing.<br>
[Issue #2418](https://redmine.postgresql.org/issues/2418) - Add Commit and Rollback buttons to the Query Tool.<br>
[Issue #3439](https://redmine.postgresql.org/issues/3439) - Allow X-FRAME-OPTIONS to be set for security. Default to SAMEORIGIN.<br>
[Issue #3559](https://redmine.postgresql.org/issues/3559) - Automatically expand child nodes as well as the selected node on the treeview if there is only one.<br>
[Issue #3886](https://redmine.postgresql.org/issues/3886) - Include multiple versions of the PG utilties in containers.<br>
[Issue #3991](https://redmine.postgresql.org/issues/3991) - Update Alpine Linux version in the docker container.<br>
[Issue #4034](https://redmine.postgresql.org/issues/4034) - Support double-click on Query Tool result grid column resize handles to auto-size to the content.<br>

# Bug fixes

[Bug #3096](https://redmine.postgresql.org/issues/3096) - Ensure size stats are prettified on the statistics tab when the UI language is not English.<br>
[Bug #3352](https://redmine.postgresql.org/issues/3352) - Handle display of roles with expiration set to infinity correctly.<br>
[Bug #3418](https://redmine.postgresql.org/issues/3418) - Allow editing of values in columns with the oid datatype which are not an actual row OID.<br>
[Bug #3544](https://redmine.postgresql.org/issues/3544) - Make the Query Tool tab titles more concise and useful.<br>
[Bug #3587](https://redmine.postgresql.org/issues/3587) - Fix support for bigint's in JSONB data.<br>
[Bug #3583](https://redmine.postgresql.org/issues/3583) - Update CodeMirror to 5.43.0 to resolve issues with auto-indent.<br>
[Bug #3600](https://redmine.postgresql.org/issues/3600) - Ensure JSON data isn't modified in-flight by psycopg2 when using View/Edit data.<br>
[Bug #3673](https://redmine.postgresql.org/issues/3673) - Modify the Download as CSV option to use the same connection as the Query Tool its running in so temporary tables etc. can be used.<br>
[Bug #3873](https://redmine.postgresql.org/issues/3873) - Fix context sub-menu alignment on Safari.<br>
[Bug #3890](https://redmine.postgresql.org/issues/3890) - Update documentation screenshots as per new design.<br>
[Bug #3906](https://redmine.postgresql.org/issues/3906) - Fix alignment of Close and Maximize button of Grant Wizard.<br>
[Bug #3911](https://redmine.postgresql.org/issues/3911) - Add full support and testsfor all PG server side encodings.<br>
[Bug #3912](https://redmine.postgresql.org/issues/3912) - Fix editing of table data with a JSON primary key.<br>
[Bug #3933](https://redmine.postgresql.org/issues/3933) - Ignore exceptions in the logger.<br>
[Bug #3942](https://redmine.postgresql.org/issues/3942) - Close connections gracefully when the user logs out of pgAdmin.<br>
[Bug #3946](https://redmine.postgresql.org/issues/3946) - Fix alignment of checkbox to drop multiple schedules of pgAgent job.<br>
[Bug #3958](https://redmine.postgresql.org/issues/3958) - Don't exclude SELECT statements from transaction management in the Query Tool in case they call data-modifying functions.<br>
[Bug #3959](https://redmine.postgresql.org/issues/3959) - Optimise display of Dependencies and Dependents, and use on-demand loading of rows in batches of 100.<br>
[Bug #3963](https://redmine.postgresql.org/issues/3963) - Fix alignment of import/export toggle switch.<br>
[Bug #3970](https://redmine.postgresql.org/issues/3970) - Prevent an error when closing the Sort/Filter dialogue with an empty filter string.<br>
[Bug #3974](https://redmine.postgresql.org/issues/3974) - Fix alignment of Connection type toggle switch of pgagent.<br>
[Bug #3981](https://redmine.postgresql.org/issues/3981) - Fix the query to set bytea_output so that read-only standbys don't consider it a write query.<br>
[Bug #3982](https://redmine.postgresql.org/issues/3982) - Add full support and testsfor all PG server side encodings.<br>
[Bug #3985](https://redmine.postgresql.org/issues/3985) - Don't embed docs and external sites in iframes, to allow the external sites to set X-FRAME-OPTIONS = DENY for security.<br>
[Bug #3992](https://redmine.postgresql.org/issues/3992) - Add full support and testsfor all PG server side encodings.<br>
[Bug #3998](https://redmine.postgresql.org/issues/3998) - Custom-encode forward slashes in URL parameters as Apache HTTPD doesn't allow them in some cases.<br>
[Bug #4000](https://redmine.postgresql.org/issues/4000) - Update CodeMirror to 5.43.0 to resolve issues with tab indent with use spaces enabled.<br>
[Bug #4013](https://redmine.postgresql.org/issues/4013) - Ensure long queries don't cause errors when downloading CSV in the Query Tool.<br>
[Bug #4021](https://redmine.postgresql.org/issues/4021) - Disable the editor and execute functions whilst queries are executing.<br>
[Bug #4022](https://redmine.postgresql.org/issues/4022) - Fix an issue where importing servers fails if a group already exists for a different user.<br>
