# Version 9.17

Release date: 2026-07-30

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.16.

# Supported Database Servers

**PostgreSQL**: 14, 15, 16, 17 and 18

**EDB Advanced Server**: 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.4

# New features

[Issue #9990](https://github.com/pgadmin-org/pgadmin4/issues/9990) -  Include the authenticated user's identity in the HTTP access log.<br>
[Issue #9942](https://github.com/pgadmin-org/pgadmin4/issues/9942) -  Add an opt-in Gateway API HTTPRoute template to the Helm chart as an alternative to the existing Ingress.<br>
[Issue #10104](https://github.com/pgadmin-org/pgadmin4/issues/10104) -  Add a preference to cap the row count fetched by the plain "View Data" action, so it is usable on large tables without always doing a full `SELECT *`.<br>
[Issue #10142](https://github.com/pgadmin-org/pgadmin4/issues/10142) -  Add support for a custom XYZ tile provider (URL, name, CRS, attribution, max zoom) in the Geometry Viewer, alongside the existing built-in base layers.<br>

# Housekeeping

[Issue #10077](https://github.com/pgadmin-org/pgadmin4/issues/10077) -  Centralize shared-server-group visibility and access-control logic, and adjust the ServerGroup-to-Server/SharedServer model relationships.<br>
[Issue #10093](https://github.com/pgadmin-org/pgadmin4/issues/10093) -  Add a basic starter contributing guide.<br>
[Issue #10135](https://github.com/pgadmin-org/pgadmin4/issues/10135) -  Fail the macOS appbundle build if any bundled library links outside the bundle, and scan all Mach-O binaries for bundle linkage rather than only files named `.so`/`.dylib`.<br>
[Issue #10154](https://github.com/pgadmin-org/pgadmin4/issues/10154) -  Pin the sonarqube-scan-action GitHub workflow to a full commit SHA instead of a mutable branch ref, to harden against supply-chain tampering.<br>
[Issue #10138](https://github.com/pgadmin-org/pgadmin4/issues/10138) -  Update the Simplified Chinese (zh_Hans_CN) translation to match the wording used in the interface.<br>
[Issue #10156](https://github.com/pgadmin-org/pgadmin4/issues/10156) -  Pin the Yarn version used by the build scripts to the `packageManager` field, rather than whatever Yarn happens to be on the builder's PATH.<br>
[Issue #10176](https://github.com/pgadmin-org/pgadmin4/issues/10176) -  Bump JavaScript and Python third-party dependencies, including axios, webpack, react, electron, and certifi.<br>

# Bug fixes

[Issue #10053](https://github.com/pgadmin-org/pgadmin4/issues/10053) -  Warn when OAuth2 provider settings are misplaced at the top level of the config instead of under `OAUTH2_CONFIG`.<br>
[Issue #10100](https://github.com/pgadmin-org/pgadmin4/issues/10100) -  Fix Schema Diff's "Generate Script" and the browser tree's CREATE Script view emitting wrong SQL for SERIAL/identity columns, by detecting column-owned sequences via pg_depend instead of guessing the sequence name.<br>
[Issue #10102](https://github.com/pgadmin-org/pgadmin4/issues/10102) -  Fix a Schema Diff result-status filter chip showing "No difference found" after being toggled off and back on, even when real differences exist.<br>
[Issue #10106](https://github.com/pgadmin-org/pgadmin4/issues/10106) -  Fix the Object Explorer briefly showing the literal HTML markup instead of a greyed-out "[Disconnecting...]" label when disconnecting a server or database.<br>
[Issue #10107](https://github.com/pgadmin-org/pgadmin4/issues/10107) -  Detect a selected-but-unusable OS keyring and fall back gracefully instead of failing.<br>
[Issue #10109](https://github.com/pgadmin-org/pgadmin4/issues/10109) -  Fix ALT+F5 ("Execute query at cursor") doing nothing when the cursor is on or near a statement that is not highlighted, in a Query Tool tab with multiple statements separated by blank lines.<br>
[Issue #10110](https://github.com/pgadmin-org/pgadmin4/issues/10110) -  Stop the googleapiclient cloud-deployment dependency from pulling in the system oauth2client package.<br>
[Issue #10117](https://github.com/pgadmin-org/pgadmin4/issues/10117) -  Fix a "'Response' object is not iterable" crash when expanding a Trigger node under a Table to view its trigger function, and report the correct error message on a node.sql failure in that flow.<br>
[Issue #10158](https://github.com/pgadmin-org/pgadmin4/issues/10158) -  Honor the selected EOL sequence when copying query text to the clipboard.<br>
[Issue #10187](https://github.com/pgadmin-org/pgadmin4/issues/10187) -  Fix the object browser's extension UI breaking under PostgreSQL 19's extension catalog changes.<br>
[Issue #10190](https://github.com/pgadmin-org/pgadmin4/issues/10190) -  Fix a tool-permission bypass where a user denied the Query Tool, Grant Wizard, or Schema Diff permission could still drive that tool's backend routes and Socket.IO handlers directly, since the permission check was applied only to a single "front door" route per tool. Also fixes a non-owner triggering an adhoc connection against another user's shared server persisting a new server record still owned by that other user (CVE-2026-17350). Reported by LXY.<br>
[Issue #10191](https://github.com/pgadmin-org/pgadmin4/issues/10191) -  Fix OS command injection in the MASTER_PASSWORD_HOOK feature, where an externally-sourced username (e.g. via OAuth2/OIDC, Kerberos, or webserver authentication) containing shell metacharacters could execute arbitrary commands as the pgAdmin service account when the configured hook string uses `%u` (CVE-2026-17347). Reported by Thiago Pereira.<br>
[Issue #10192](https://github.com/pgadmin-org/pgadmin4/issues/10192) -  Fix a lexer-differential bypass of the AI Assistant's read-only transaction guard, where sqlparse's string-literal lexing disagreed with PostgreSQL's own parser under `standard_conforming_strings = on`, letting a crafted multi-statement payload smuggle a COMMIT past the intended read-only wrapper, an incomplete fix for CVE-2026-12045 (CVE-2026-17351). Reported by Kai Aizen.<br>
[Issue #10193](https://github.com/pgadmin-org/pgadmin4/issues/10193) -  Fix SQL injection in the Index Statistics all-indexes listing and the Publications/Subscriptions Dependencies views, where an apostrophe in a table, index, publication, or subscription name broke out of an unescaped template interpolation, an incomplete fix for CVE-2026-12044 (CVE-2026-17346). Reported by Hung Tran Quoc (@rampage0010).<br>
[Issue #10194](https://github.com/pgadmin-org/pgadmin4/issues/10194) -  Fix several Constraints, Preferences, Debugger, and Schema Diff routes missing the `@pga_login_required` decorator, making them reachable without authentication in server mode, an incomplete fix for CVE-2026-12046 (CVE-2026-17348). Reported by Hung Tran Quoc (@rampage0010).<br>
[Issue #10200](https://github.com/pgadmin-org/pgadmin4/issues/10200) -  Fix an adhoc server connection cloning another user's stored database credentials (password, save password flag, tunnel password) alongside ownership, letting a non-owner who cloned another user's shared server connect using that user's saved database password (CVE-2026-17349).<br>
[Issue #10213](https://github.com/pgadmin-org/pgadmin4/issues/10213) -  Fix OS command injection in the Import/Export Data tool, where a query-based export could pass a crafted query string past the `\copy (...)` parenthesis-balance guard by exploiting a backslash-escape mismatch with psql's default `standard_conforming_strings = on` behaviour, exposing a live `TO PROGRAM` clause for arbitrary command execution (CVE-2026-17566). Reported by Arpit Jain.<br>

# Additional changes (no associated issue)

## Dependencies

Non-breaking `dependabot` and audit-driven bumps aggregated for v9.17.

Python:

`certifi` 2026.5.20 -> 2026.6.17<br>

JavaScript (`web/`):

`@babel/core` 7.29.0 -> 7.29.6<br>
`@date-io/date-fns` 3.x -> 3.2.1<br>
`@szhsin/react-menu` 4.5.1 -> 4.5.2<br>
`@tanstack/react-query` 5.100.9 -> 5.101.4<br>
`@types/react` 19.2.14 -> 19.2.17<br>
`ajv` 8.18.0 -> 8.20.0<br>
`anti-trojan-source` 1.8.1 -> 1.12.0<br>
`autoprefixer` 10.5.0 -> 10.5.4<br>
`axios` 1.16.0 -> 1.18.1<br>
`core-js` added as an explicit direct dependency (3.49.0), previously relied on transitively via pickr<br>
`date-fns` 4.1.0 -> 4.4.0<br>
`dompurify` 3.4.1 -> 3.4.12<br>
`eslint-plugin-jest` 29.15.2 -> 29.15.5<br>
`globals` 17.5.0 -> 17.7.0<br>
`hotkeys-js` 4.0.3 -> 4.0.4<br>
`html-to-image` 1.11.11 -> 1.11.13<br>
`ip-address` 10.1.1 -> 10.2.0<br>
`jest` / `jest-environment-jsdom` 30.3.0 -> 30.4.2 / 30.4.1<br>
`marked` 18.0.2 -> 18.0.7<br>
`moment-timezone` 0.6.2 -> 0.6.3<br>
`papaparse` 5.5.3 -> 5.5.4<br>
`postcss` 8.5.10/8.5.14 -> 8.5.20<br>
`react` / `react-dom` 19.2.5 -> 19.2.7<br>
`react-checkbox-tree` 2.0.1 -> 2.0.2<br>
`react-draggable` 4.5.0 -> 4.7.0<br>
`react-timer-hook` 4.0.5 -> 4.0.6<br>
`sharp` 0.34.4 -> 0.35.3<br>
`sql-formatter` 15.7.3 -> 15.8.2<br>
`svgo` 4.0.1 -> 4.0.2<br>
`terser-webpack-plugin` 5.5.0 -> 5.6.1<br>
`typescript-eslint` 8.59.0 -> 8.65.0<br>
`webpack` 5.106.2 -> 5.108.4<br>
`webpack-bundle-analyzer` 5.3.0 -> 5.3.1<br>
`zustand` 5.0.12 -> 5.0.14<br>
Transitive advisory bumps: `brace-expansion`, `form-data`, `undici`, `js-yaml`<br>
Reverted `azure-mgmt-resource` to avoid a breaking import-path change in `pgadmin/misc/cloud/azure`<br>

JavaScript (`runtime/`):

`electron` 42.3.3 -> 43.1.1<br>
`eslint` 10.4.1 -> 10.7.0<br>
`globals` 17.6.0 -> 17.7.0<br>
`axios` 1.18.0 -> 1.18.1<br>
`ip-address` 10.1.1 -> 10.2.0<br>
`postcss` 8.5.10 -> 8.5.20<br>
Pinned `packageManager` to `yarn@4.15.0`<br>
