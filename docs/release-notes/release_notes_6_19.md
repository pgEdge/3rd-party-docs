# Version 6.19

Release date: 2023-01-17

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v6.18.

# Supported Database Servers

**PostgreSQL**: 10, 11, 12, 13, 14 and 15

**EDB Advanced Server**: 10, 11, 12, 13, 14 and 15

# New features

[Issue #5569](https://github.com/pgadmin-org/pgadmin4/issues/5569) -  Added support of AWS provider for BigAnimal cloud deployment.<br>

# Housekeeping

[Issue #5563](https://github.com/pgadmin-org/pgadmin4/issues/5563) -  Allow YouTube video demo links to be added to appropriate pgAdmin documentation.<br>
[Issue #5615](https://github.com/pgadmin-org/pgadmin4/issues/5615) -  Rewrite pgAdmin main menu bar to use React.<br>

# Bug fixes

[Issue #5487](https://github.com/pgadmin-org/pgadmin4/issues/5487) -  Fixed an issue where incorrect password used with shared server.<br>
[Issue #5541](https://github.com/pgadmin-org/pgadmin4/issues/5541) -  Ensure the browser tree does not freeze while rendering 10k+ nodes/objects.<br>
[Issue #5542](https://github.com/pgadmin-org/pgadmin4/issues/5542) -  Fixed an issue updating the schema node de-select the node in the tree if only one schema is present in the collection node.<br>
[Issue #5559](https://github.com/pgadmin-org/pgadmin4/issues/5559) -  Fixed tree flickering issue on scroll.<br>
[Issue #5577](https://github.com/pgadmin-org/pgadmin4/issues/5577) -  Fixed an issue where the default value of string for columns should wrap in quotes in the create script.<br>
[Issue #5586](https://github.com/pgadmin-org/pgadmin4/issues/5586) -  Fix the webserver and internal authentication setup issue.<br>
[Issue #5613](https://github.com/pgadmin-org/pgadmin4/issues/5613) -  Ensure the appbundle has correct permissions so that pgAdmin can be accessed by users other than owner.<br>
[Issue #5622](https://github.com/pgadmin-org/pgadmin4/issues/5622) -  Fixed an issue where the ignore owner flag is not working for some cases in the Schema Diff.<br>
[Issue #5626](https://github.com/pgadmin-org/pgadmin4/issues/5626) -  Fixed an issue where actions performed on the tree node should update the context menu options.<br>
[Issue #5627](https://github.com/pgadmin-org/pgadmin4/issues/5627) -  Ensure that the submenus under the trigger's context menu are enabled/disabled correctly.<br>
[Issue #5640](https://github.com/pgadmin-org/pgadmin4/issues/5640) -  Update boto3 & botocore to the latest version.<br>
[Issue #5641](https://github.com/pgadmin-org/pgadmin4/issues/5641) -  Fixed an issue where Geometry viewer does not show popup when columns are less than 3.<br>
[Issue #5647](https://github.com/pgadmin-org/pgadmin4/issues/5647) -  Fixed an issue where row count notification was disappearing automatically.<br>
[Issue #5661](https://github.com/pgadmin-org/pgadmin4/issues/5661) -  Fix select dropdown border issue.<br>
[Issue #5666](https://github.com/pgadmin-org/pgadmin4/issues/5666) -  Fixed a missing "jwks_uri" in metadata error that occurred when logging in with an oAuth2 provider like Azure or Google.<br>
[Issue #5675](https://github.com/pgadmin-org/pgadmin4/issues/5675) -  Fixed an issue where rename panel was losing focus when trying to add name if input box is empty.<br>
[Issue #5734](https://github.com/pgadmin-org/pgadmin4/issues/5734) -  Ensure that the authenticated users can't access each other's directories and files by providing relative paths.(CVE-2023-0241)<br>
