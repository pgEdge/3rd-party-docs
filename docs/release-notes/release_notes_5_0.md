# Version 5.0

Release date: 2021-02-25

This release contains a number of bug fixes and new features since the release of pgAdmin4 4.30.

# New features

[Issue #5091](https://redmine.postgresql.org/issues/5091) -  Make Statistics, Dependencies, Dependants tabs closable and the user can add them back using the 'Add panel' option.<br>
[Issue #5912](https://redmine.postgresql.org/issues/5912) -  Added support for Logical Replication.<br>
[Issue #5967](https://redmine.postgresql.org/issues/5967) -  Implemented runtime using NWjs to open pgAdmin4 in a standalone window instead of the system tray and web browser.<br>
[Issue #6148](https://redmine.postgresql.org/issues/6148) -  Added Quick Search functionality for menu items and help articles.<br>
[Issue #6153](https://redmine.postgresql.org/issues/6153) -  Added publication and subscription support in Schema Diff.<br>
[Issue #6170](https://redmine.postgresql.org/issues/6170) -  Ensure logs are not stored in the container, and only sent to the console.<br>

# Housekeeping

[Issue #5017](https://redmine.postgresql.org/issues/5017) -  Use cheroot as the default production server for pgAdmin4.<br>
[Issue #6145](https://redmine.postgresql.org/issues/6145) -  Documentation of Logical Replication.<br>
[Issue #6195](https://redmine.postgresql.org/issues/6195) -  Documentation of runtime with NWjs.<br>
[Issue #6196](https://redmine.postgresql.org/issues/6196) -  Documentation of Quick Search support.<br>
[Issue #6207](https://redmine.postgresql.org/issues/6207) -  Updated the JS dependencies to the latest.<br>

# Bug fixes

[Issue #5809](https://redmine.postgresql.org/issues/5809) -  Fixed an issue where the focus is not properly set on the filter text editor after closing the error dialog.<br>
[Issue #5871](https://redmine.postgresql.org/issues/5871) -  Ensure that username should be visible in the 'Connect to Server' popup when service and user name both specified.<br>
[Issue #6045](https://redmine.postgresql.org/issues/6045) -  Fixed autocomplete issue where it is not showing any suggestions if the schema name contains escape characters.<br>
[Issue #6087](https://redmine.postgresql.org/issues/6087) -  Fixed an issue where the dependencies tab showing multiple owners for the objects having shared dependencies.<br>
[Issue #6117](https://redmine.postgresql.org/issues/6117) -  Fixed an issue where the user is unable to update column-level privileges from the security tab.<br>
[Issue #6143](https://redmine.postgresql.org/issues/6143) -  Fixed an issue where shared server entries not getting deleted from SQLite database if the user gets deleted.<br>
[Issue #6157](https://redmine.postgresql.org/issues/6157) -  Fixed an issue where strike-through is not visible for rows selected for deletion after scrolling.<br>
[Issue #6163](https://redmine.postgresql.org/issues/6163) -  Fixed an issue where Zoom to fit button only works if the diagram is larger than the canvas.<br>
[Issue #6164](https://redmine.postgresql.org/issues/6164) -  Ensure that the diagram should not vanish entirely if zooming out too far in ERD.<br>
[Issue #6177](https://redmine.postgresql.org/issues/6177) -  Fixed an issue while downloading ERD images in Safari and Firefox.<br>
[Issue #6178](https://redmine.postgresql.org/issues/6178) -  Fixed an issue where the user unable to change the background color for a server.<br>
[Issue #6179](https://redmine.postgresql.org/issues/6179) -  Fixed an issue where Generate SQL displayed twice in the ERD tool.<br>
[Issue #6180](https://redmine.postgresql.org/issues/6180) -  Updated missing documentation for the 'Download Image' option in ERD.<br>
[Issue #6187](https://redmine.postgresql.org/issues/6187) -  Limit the upgrade check to run once per day.<br>
[Issue #6193](https://redmine.postgresql.org/issues/6193) -  Ensure that ERD throws a warning before closing unsaved changes if open in a new tab.<br>
[Issue #6197](https://redmine.postgresql.org/issues/6197) -  Fixed an issue where the ERD image is not properly downloaded.<br>
[Issue #6201](https://redmine.postgresql.org/issues/6201) -  Added SSL support for creating a subscription.<br>
[Issue #6208](https://redmine.postgresql.org/issues/6208) -  Fixed an issue where utility(Backup, Maintenance, ...) jobs are failing when the log level is set to DEBUG.<br>
[Issue #6230](https://redmine.postgresql.org/issues/6230) -  Fixed an issue where the user is not able to create the subscription.<br>
[Issue #6250](https://redmine.postgresql.org/issues/6250) -  Ensure DEB/RPM packages depend on the same version of each other.<br>
