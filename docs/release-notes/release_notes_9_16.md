# Version 9.16

Release date: 2026-06-18

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.15.

!!! warning

    Starting with this release, when a server connection is configured with both an external password-exec command and a passfile, pgAdmin now uses the passfile and ignores the password-exec command. Previously the password-exec command took precedence. A warning is written to the log when the password-exec command is ignored in favour of the passfile.

# Supported Database Servers

**PostgreSQL**: 14, 15, 16, 17 and 18

**EDB Advanced Server**: 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.4

# New features

[Issue #2431](https://github.com/pgadmin-org/pgadmin4/issues/2431) -  Added an option to colourize panel and tab headers based on the connected server's colour, making it easier to tell which server a tab is connected to at a glance.<br>
[Issue #7641](https://github.com/pgadmin-org/pgadmin4/issues/7641) -  Allow the OAuth2 login button icon to use any Font Awesome style (e.g. `fas fa-key`), not only brand icons.<br>
[Issue #9301](https://github.com/pgadmin-org/pgadmin4/issues/9301) -  Added a "Back to login" link to the Forgot Password and Reset Password pages.<br>
[Issue #9626](https://github.com/pgadmin-org/pgadmin4/issues/9626) -  Add support for the TOAST tuple target storage parameter in the Materialized View dialog.<br>
[Issue #9646](https://github.com/pgadmin-org/pgadmin4/issues/9646) -  Make the init container security context in the Helm chart configurable via containerSecurityContext, consistent with the main container.<br>
[Issue #9699](https://github.com/pgadmin-org/pgadmin4/issues/9699) -  Add support for closing a tab with a middle-click on its title.<br>

# Housekeeping

[Issue #9817](https://github.com/pgadmin-org/pgadmin4/issues/9817) -  Correct the macOS Replace keyboard shortcut documented in the Query Tool toolbar reference.<br>
[Issue #9866](https://github.com/pgadmin-org/pgadmin4/issues/9866) -  Fix a typo in the PGPASS_FILE environment-variable documentation and clarify its behaviour.<br>
[Issue #9917](https://github.com/pgadmin-org/pgadmin4/issues/9917) -  Update the Italian translation.<br>
[Issue #9959](https://github.com/pgadmin-org/pgadmin4/issues/9959) -  Bump Electron in the desktop runtime from 41.5.0 to 42.1.0 and pin the packaged Electron version.<br>
[Issue #9981](https://github.com/pgadmin-org/pgadmin4/issues/9981) -  Clarify the SSH tunnel "Prompt for identity file password?" switch label and help text to indicate it applies only to identity-file authentication.<br>
[Issue #10014](https://github.com/pgadmin-org/pgadmin4/issues/10014) -  Document the Kubernetes init-container tag-mismatch and RollingUpdate hazards on shared data volumes.<br>
[Issue #10018](https://github.com/pgadmin-org/pgadmin4/issues/10018) -  Remove the EDB BigAnimal cloud deployment support.<br>
[Issue #10023](https://github.com/pgadmin-org/pgadmin4/issues/10023) -  Bump JavaScript and Python third-party dependencies, including cryptography, axios, shell-quote, and assorted dependabot lockfile updates.<br>
[Issue #10049](https://github.com/pgadmin-org/pgadmin4/issues/10049) -  Bump vulnerable transitive dependencies (tar, flatted) to clear security advisories.<br>
[Issue #10050](https://github.com/pgadmin-org/pgadmin4/issues/10050) -  Rebase the version-specific SQL templates so the default targets PostgreSQL 14, the oldest supported server version, dropping the obsolete sub-14 template buckets.<br>
[Issue #10063](https://github.com/pgadmin-org/pgadmin4/issues/10063) -  Strip the foreign-architecture slice from the macOS bundle so single-arch builds no longer ship the universal2 Python framework's unused arm64/x86_64 code.<br>
[Issue #10094](https://github.com/pgadmin-org/pgadmin4/issues/10094) -  Update the Italian translation for v9.16.<br>

# Bug fixes

[Issue #6308](https://github.com/pgadmin-org/pgadmin4/issues/6308) -  Fix the infinite loading spinner after an idle database connection is silently dropped, by detecting stale connections and offering a reconnect dialog.<br>
[Issue #7346](https://github.com/pgadmin-org/pgadmin4/issues/7346) -  Fixed an issue where preferences set via the CLI (setup.py set-prefs) were not validated, so invalid values were stored silently; CLI preference values are now validated against the preference type and rejected (and reported) if invalid.<br>
[Issue #7596](https://github.com/pgadmin-org/pgadmin4/issues/7596) -  Fix the Query Tool turning into a blank white screen when the runtime has a malformed default locale, by guarding the Query History date/time formatting against the resulting RangeError.<br>
[Issue #8318](https://github.com/pgadmin-org/pgadmin4/issues/8318) -  Fixed an error ("i.default.find(...) is undefined") that prevented deleting a table or relationship link in the ERD tool when a foreign key referenced a column that had been renamed.<br>
[Issue #9060](https://github.com/pgadmin-org/pgadmin4/issues/9060) -  Fixed an issue in the Create Table dialog where renaming a column did not update the column references in foreign key and unique constraint definitions for the new table.<br>
[Issue #9091](https://github.com/pgadmin-org/pgadmin4/issues/9091) -  Fix the Query Tool re-prompting for an unsaved password in a loop and rejecting the re-entered password, by caching the entered password on the server manager when the primary connection is already established.<br>
[Issue #9128](https://github.com/pgadmin-org/pgadmin4/issues/9128) -  Fixed an issue where the object breadcrumbs popup blocked clicks on the object explorer items beneath it.<br>
[Issue #9595](https://github.com/pgadmin-org/pgadmin4/issues/9595) -  Fix missing ALTER ... SET DEFAULT statements for inherited columns in the generated table SQL/EDIT script.<br>
[Issue #9677](https://github.com/pgadmin-org/pgadmin4/issues/9677) -  Fix the Unlogged table toggle in table properties not generating any ALTER TABLE ... SET LOGGED/UNLOGGED statement.<br>
[Issue #9701](https://github.com/pgadmin-org/pgadmin4/issues/9701) -  Ensure pgAdmin uses psycopg3 rather than psycopg2 when connecting to the configuration database via PGADMIN_CONFIG_CONFIG_DATABASE_URI.<br>
[Issue #9744](https://github.com/pgadmin-org/pgadmin4/issues/9744) -  Fix a View/Edit Data crash when the session contains a transaction object that is not filter-capable (e.g. left by the Query Tool or persisted by an older version), which could prevent the desktop application from loading after an upgrade.<br>
[Issue #9762](https://github.com/pgadmin-org/pgadmin4/issues/9762) -  Fix the "Cannot read properties of undefined (reading 'map')" crash in the desktop runtime when a menu refresh is triggered before the application menus have been received.<br>
[Issue #9766](https://github.com/pgadmin-org/pgadmin4/issues/9766) -  Fixed an issue where a server's custom foreground colour was not applied to the object counts (children count) and column type labels shown in the object explorer.<br>
[Issue #9782](https://github.com/pgadmin-org/pgadmin4/issues/9782) -  Optimise pgAdmin startup time by replacing the fixed server-ping interval with adaptive pinging that backs off when no servers are configured.<br>
[Issue #9806](https://github.com/pgadmin-org/pgadmin4/issues/9806) -  Fixed an issue where generated font and image filenames contained a double dot (e.g. `Roboto-Bold..ttf`) due to the webpack asset filename template.<br>
[Issue #9828](https://github.com/pgadmin-org/pgadmin4/issues/9828) -  Fix tool calls failing against OpenAI-compatible providers that emit empty/null name, arguments, or id fields in streaming continuation deltas.<br>
[Issue #9829](https://github.com/pgadmin-org/pgadmin4/issues/9829) -  Fixed installation on Python 3.9 (e.g. RHEL/Rocky/AlmaLinux 8 and 9) failing with "No module named 'pkg_resources'" by pinning setuptools below the version that dropped pkg_resources for Python 3.9.<br>
[Issue #9854](https://github.com/pgadmin-org/pgadmin4/issues/9854) -  Fix the JSON editor stripping trailing fractional zeros (e.g. 10.00) and rewriting large integers in jsonb values, which corrupted unmodified numbers when saving.<br>
[Issue #9864](https://github.com/pgadmin-org/pgadmin4/issues/9864) -  Fixed a regression where the rectangular (block/column) text selection in the code editor stopped working; restored the default Alt+drag selection without re-introducing the crosshair cursor on Alt+F5 (#9570).<br>
[Issue #9868](https://github.com/pgadmin-org/pgadmin4/issues/9868) -  Warn before opening a very large JSON/JSONB value in the data grid cell editor, which could freeze pgAdmin, and let the user choose whether to proceed.<br>
[Issue #9875](https://github.com/pgadmin-org/pgadmin4/issues/9875) -  Fixed an issue where EXPLAIN and EXPLAIN ANALYZE failed to execute when blank lines separated clauses in the SQL query.<br>
[Issue #9810](https://github.com/pgadmin-org/pgadmin4/issues/9810) -  Use the ServerManager's passfile for the credential gate in connect() so the check matches the passfile actually used for the connection, and warn on conflicting passfile/passexec settings.<br>
[Issue #9892](https://github.com/pgadmin-org/pgadmin4/issues/9892) -  Fix blank difference counts on the top-level group rows in Schema Diff.<br>
[Issue #9896](https://github.com/pgadmin-org/pgadmin4/issues/9896) -  Fix invalid DDL reconstruction for SERIAL columns in Schema Diff and the generated SQL/CREATE Script so the output round-trips on a clean target.<br>
[Issue #9933](https://github.com/pgadmin-org/pgadmin4/issues/9933) -  Remove the administrator-role bypass from the server-access helpers so the access-control checks added in 9.15 (CVE-2026-7813) are enforced uniformly. The Administrator role manages pgAdmin itself, not other users' database connections.<br>
[Issue #9935](https://github.com/pgadmin-org/pgadmin4/issues/9935) -  Fix "Illegal instruction" crash on startup of the Linux DEB and RPM packages on older x86_64 CPUs by pinning the psycopg C extension build to the x86-64 baseline.<br>
[Issue #9936](https://github.com/pgadmin-org/pgadmin4/issues/9936) -  Fix the AI panel silently falling back to the default provider when a custom LLM API URL or key file was set, and allow self-hosted LLM endpoints on any loopback port.<br>
[Issue #9939](https://github.com/pgadmin-org/pgadmin4/issues/9939) -  Fix saving a newly-added row in the Query Tool failing when the result set includes expression or alias columns that are not real columns of the underlying table.<br>
[Issue #9952](https://github.com/pgadmin-org/pgadmin4/issues/9952) -  Ship libpq-oauth-18.so and libcurl in the Docker image so PostgreSQL 18 OAuth connections work out of the box.<br>
[Issue #9976](https://github.com/pgadmin-org/pgadmin4/issues/9976) -  Fix a startup migration crash (NoSuchTableError) when an old configuration database contains a stale foreign-key reference.<br>
[Issue #9984](https://github.com/pgadmin-org/pgadmin4/issues/9984) -  Fix the Docker entrypoint mishandling a quoted PGADMIN_CONFIG_CONFIG_DATABASE_URI, which caused a SQLAlchemy parse error and silently skipped PGADMIN_DEFAULT_EMAIL/PASSWORD setup.<br>
[Issue #9985](https://github.com/pgadmin-org/pgadmin4/issues/9985) -  Make CAP_NET_BIND_SERVICE optional in the Docker image so the container works on restricted runtimes that disable that capability.<br>
[Issue #9987](https://github.com/pgadmin-org/pgadmin4/issues/9987) -  Fix "AttributeError: 'PgAdmin' object has no attribute 'login_manager'" crash when running setup.py user-management commands (add-user, update-user) from the CLI.<br>
[Issue #9988](https://github.com/pgadmin-org/pgadmin4/issues/9988) -  Provide an actionable error when 'openid' is in OAUTH2_SCOPE but OAUTH2_SERVER_METADATA_URL is not set, instead of a cryptic Authlib failure.<br>
[Issue #10013](https://github.com/pgadmin-org/pgadmin4/issues/10013) -  Fix pg_attribute filtering by attname returning the wrong attnum when the same column name exists in multiple tables.<br>
[Issue #10022](https://github.com/pgadmin-org/pgadmin4/issues/10022) -  Fix AI Assistant read-only transaction bypass that allowed prompt-injected multi-statement payloads to commit out of the READ ONLY wrapper and execute arbitrary SQL, chaining to RCE via COPY ... TO PROGRAM on a superuser connection (CVE-2026-12045). Reported by Isaac Chen.<br>
[Issue #10026](https://github.com/pgadmin-org/pgadmin4/issues/10026) -  Fix SQL injection in the named restore point endpoint where the user-supplied restore point name was interpolated into SQL via str.format() instead of being passed as a bound parameter (CVE-2026-12050). Reported by Geo.<br>
[Issue #10027](https://github.com/pgadmin-org/pgadmin4/issues/10027) -  Fix the spurious "Crypt key is missing" error and logged traceback in the Query Tool new-connection endpoints after a backend restart, by surfacing it as the standard CRYPTKEY_MISSING response so the client recovers transparently.<br>
[Issue #10028](https://github.com/pgadmin-org/pgadmin4/issues/10028) -  Fix open redirect in the multi-factor authentication flow via an unvalidated 'next' parameter that allowed a crafted link to redirect an authenticated victim to an attacker-controlled host (CVE-2026-12049). Reported by Mai Phạm Hiền.<br>
[Issue #10029](https://github.com/pgadmin-org/pgadmin4/issues/10029) -  Fixed a regression where the rectangular (block/column) text selection in the code editor stopped working; restored the default Alt+drag selection without re-introducing the crosshair cursor on Alt+F5 (#9570).<br>
[Issue #10030](https://github.com/pgadmin-org/pgadmin4/issues/10030) -  Accept prepare and binary keyword arguments in DictCursor.execute() so callers using the modern psycopg API no longer raise TypeError.<br>
[Issue #10059](https://github.com/pgadmin-org/pgadmin4/issues/10059) -  Fix the generated SQL for editing a SQL-language function/procedure whose body contains the word "return" (e.g. a RETURNING clause), which was wrongly treated as a SQL-standard body and produced a statement without the AS $BODY$ wrapper.<br>
[Issue #10068](https://github.com/pgadmin-org/pgadmin4/issues/10068) -  Fix critical stored cross-site scripting where PostgreSQL server error text and Explain plan-node content passed through html-react-parser across notifier toasts, form errors, modal alerts, and the Explain visualiser. Because pgAdmin's default Content-Security-Policy allows inline script and an iframe `srcdoc` inherits the embedding origin, the injected JavaScript ran same-origin to the victim's authenticated pgAdmin session and could read every saved server connection credential and issue arbitrary SQL against every server the victim was connected to (CVE-2026-12048). Reported by Fernando Bortotti.<br>
[Issue #10069](https://github.com/pgadmin-org/pgadmin4/issues/10069) -  Fix HTML injection in the cloud deployment module (RDS, Azure, Google) where SDK exception text was forwarded to the browser unsanitised and rendered through html-react-parser in the Cloud Wizard (CVE-2026-12047). Reported by Fernando Bortotti.<br>
[Issue #10072](https://github.com/pgadmin-org/pgadmin4/issues/10072) -  Fix two SQL Editor endpoints (close and update_connection) missing the `@pga_login_required` decorator, making them reachable without authentication in server mode and exposing a pickle deserialization sink (CVE-2026-12046). Reported by Fernando Bortotti.<br>
[Issue #10078](https://github.com/pgadmin-org/pgadmin4/issues/10078) -  Fix SQL injection across sixteen dialog templates that rendered `COMMENT ON ... IS '<description>'` and the related `pgstattuple`/`pgstatindex` stats sinks, where a low-privilege user could plant a table or index name containing an apostrophe and a superuser viewing statistics on the object would trigger SQL execution under the superuser role. The fix switches the affected templates to the `qtLiteral` escape filter and rewrites the `pgstattuple`/`pgstatindex` calls to address the relation via OID with a `::oid::regclass` cast, eliminating the embedded string-literal call form entirely. Also hardens `qtLiteral` to raise instead of silently returning raw values when `conn` is falsy (CVE-2026-12044). Reported by Jasser Chebbi.<br>

# Additional changes (no associated issue)

## Dependencies

Non-breaking `dependabot` updates aggregated for v9.16.

Python:

`certifi` 2026.4.22 -> 2026.5.20<br>
`cryptography` 47.0 -> 49.0<br>
`Flask-Security-Too` 5.4 -> 5.6 (`python_version <= '3.9'`)<br>
`google-auth-oauthlib` 1.3.1 -> 1.4.0 (`python_version > '3.9'`; Python 3.9 stays on 1.3.1)<br>
`requests` >=2.33.1 -> >=2.34.2 (`python_version > '3.9'`)<br>
`safety` >=3.7.0 -> >=3.8.1<br>
`selenium` 4.43.0 -> 4.44.0<br>
`testscenarios` 0.6.1 -> 0.6.2 (`python_version > '3.9'`)<br>
`typer` 0.25 -> 0.26 (`python_version > '3.9'`)<br>
`urllib3` 2.6 -> 2.7 (`python_version > '3.9'`)<br>

JavaScript (`web/`):

`@tanstack/react-query` 5.100.5 -> 5.100.9<br>
`ip-address` 10.1.0 -> 10.1.1<br>
`postcss` 8.5.12 -> 8.5.14<br>

JavaScript (`runtime/`):

`axios` 1.16.0 -> 1.18.0<br>
`electron` 41.5.0 -> 42.3.3<br>
`eslint` 10.3.0 -> 10.4.1<br>
