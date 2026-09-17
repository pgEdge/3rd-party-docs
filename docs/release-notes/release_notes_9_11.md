# Version 9.11

Release date: 2025-12-11

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.10.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16, 17 and 18

**EDB Advanced Server**: 13, 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.0

# New features

[Issue #6388](https://github.com/pgadmin-org/pgadmin4/issues/6388) -  Add support of DEPENDS/NO DEPENDS ON EXTENSION for INDEX.<br>
[Issue #6390](https://github.com/pgadmin-org/pgadmin4/issues/6390) -  Add support of DEPENDS/NO DEPENDS ON EXTENSION for MATERIALIZED VIEW.<br>
[Issue #8968](https://github.com/pgadmin-org/pgadmin4/issues/8968) -  Add support for showing the column data type beside column name in the object explorer.<br>
[Issue #9263](https://github.com/pgadmin-org/pgadmin4/issues/9263) -  Added Ubuntu 25.10 and Fedora 43<br>
[Issue #9345](https://github.com/pgadmin-org/pgadmin4/issues/9345) -  Add conditional TLS support for the Ingress in the Helm chart.<br>

# Housekeeping

# Bug fixes

[Issue #9036](https://github.com/pgadmin-org/pgadmin4/issues/9036) -  Fixed an issue on type casting of default string function/procedure arguments in debugger tool.<br>
[Issue #9155](https://github.com/pgadmin-org/pgadmin4/issues/9155) -  Fix pkg_resources deprecation warning by migrating Docker base image to python:3-alpine from alpine:latest.<br>
[Issue #9297](https://github.com/pgadmin-org/pgadmin4/issues/9297) -  Fixed an issue where EXPLAIN should run on query under cursor if no text is selected.<br>
[Issue #9351](https://github.com/pgadmin-org/pgadmin4/issues/9351) -  Fixed an issue where opening file in Query Tool does not retain file name in tab.<br>
[Issue #9354](https://github.com/pgadmin-org/pgadmin4/issues/9354) -  Fixed an issue where connection is failing via Query Tool/PSQL Tool workspaces.<br>
[Issue #9368](https://github.com/pgadmin-org/pgadmin4/issues/9368) -  Plain SQL restore runs with 'restrict' option to prevent harmful psql meta-commands (CVE-2025-13780).<br>
[Issue #9372](https://github.com/pgadmin-org/pgadmin4/issues/9372) -  Fixed an issue where copying highlighted text in the query tool data output cell editor would copy the complete string.<br>
[Issue #9373](https://github.com/pgadmin-org/pgadmin4/issues/9373) -  Fixed an issue where copying a single cell should not add quoting.<br>
[Issue #9393](https://github.com/pgadmin-org/pgadmin4/issues/9393) -  Fix the Helm chart server definition and change the app version.<br>
[Issue #9399](https://github.com/pgadmin-org/pgadmin4/issues/9399) -  Specify the correct hostname placeholder for the Password Exec command.<br>
[Issue #9408](https://github.com/pgadmin-org/pgadmin4/issues/9408) -  Ensure the proper handling of extra volume mount configurations in the Helm deployment template by correcting the configuration value references.<br>
