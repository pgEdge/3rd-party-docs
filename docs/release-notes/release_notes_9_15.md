# Version 9.15

Release date: 2026-05-11

This release contains a number of bug fixes and new features since the release of pgAdmin 4 v9.14.

# Supported Database Servers

**PostgreSQL**: 13, 14, 15, 16, 17 and 18

**EDB Advanced Server**: 13, 14, 15, 16, 17 and 18

# Bundled PostgreSQL Utilities

**psql**, **pg_dump**, **pg_dumpall**, **pg_restore**: 18.2

# New features

[Issue #9657](https://github.com/pgadmin-org/pgadmin4/issues/9657) -  Allow the container image to run as a non-default user via the PUID and PGID environment variables.<br>

# Housekeeping

[Issue #9764](https://github.com/pgadmin-org/pgadmin4/issues/9764) -  Update the Swedish translation.<br>
[Issue #9827](https://github.com/pgadmin-org/pgadmin4/issues/9827) -  Bump Python and JavaScript dependencies.<br>
[Issue #9832](https://github.com/pgadmin-org/pgadmin4/issues/9832) -  Fix the Czech translation for 'Refresh'.<br>
[Issue #9834](https://github.com/pgadmin-org/pgadmin4/issues/9834) -  Bump runtime dependencies and upgrade ESLint to v10.<br>
[Issue #9839](https://github.com/pgadmin-org/pgadmin4/issues/9839) -  Update the Russian translation.<br>
[Issue #9870](https://github.com/pgadmin-org/pgadmin4/issues/9870) -  Bump runtime and development dependencies.<br>
[Issue #9873](https://github.com/pgadmin-org/pgadmin4/issues/9873) -  Use an `<OWNER>` placeholder in resql tests instead of a hardcoded 'postgres' role to support non-default superuser names.<br>
[Issue #9893](https://github.com/pgadmin-org/pgadmin4/issues/9893) -  Update the Spanish translation.<br>
[Issue #9906](https://github.com/pgadmin-org/pgadmin4/issues/9906) -  Update the Italian translation.<br>

# Bug fixes

[Issue #9656](https://github.com/pgadmin-org/pgadmin4/issues/9656) -  Use absolute paths for `a2enmod` and `a2enconf` in the Debian setup script so it works when `/usr/sbin` is not on PATH.<br>
[Issue #9830](https://github.com/pgadmin-org/pgadmin4/issues/9830) -  Fix cross-user data access and shared-server privilege escalation in server mode (CVE-2026-7813). Also applies the `@with_object_filters` access-control decorator to `ServerNode.list`.<br>
[Issue #9835](https://github.com/pgadmin-org/pgadmin4/issues/9835) -  Tighten Shared Server feature parity, owner-only field handling, and write guards as a follow-up to the data-isolation hardening.<br>
[Issue #9865](https://github.com/pgadmin-org/pgadmin4/issues/9865) -  Fix stored cross-site scripting (XSS) via crafted PostgreSQL object names rendered in the Browser Tree and Explain Visualizer (CVE-2026-7814). Reported by Fahar Abbas.<br>
[Issue #9898](https://github.com/pgadmin-org/pgadmin4/issues/9898) -  Fix SQL injection in Maintenance tool option values (CVE-2026-7815). Reported by j3seer.<br>
[Issue #9899](https://github.com/pgadmin-org/pgadmin4/issues/9899) -  Fix OS command injection in Import/Export query export (CVE-2026-7816). Reported by Chung Kim (chungkn), OneMount Group.<br>
[Issue #9900](https://github.com/pgadmin-org/pgadmin4/issues/9900) -  Fix local-file inclusion and server-side request forgery in LLM API configuration endpoints (CVE-2026-7817). Reported by j3seer.<br>
[Issue #9901](https://github.com/pgadmin-org/pgadmin4/issues/9901) -  Fix unsafe deserialization in the session manager that could lead to remote code execution (CVE-2026-7818). Also encrypts session files at rest using Fernet, restricts session-file permissions to 0o600, switches the session-digest default from SHA-1 to SHA-256, drops several non-roundtrippable live objects from the session (`AuthSourceManager` and the Azure, RDS, Google Cloud, and BigAnimal cloud-provider instances), tightens DATA_DIR file and directory permissions at creation, creates `pgadmin4.log` with mode 0o600, hardens `EnhancedRotatingFileHandler._open` against rotation failures, and bounds the `user_info_server` prompt retry loop so a non-interactive caller cannot spin forever. Reported by Fernando Bortotti.<br>
[Issue #9902](https://github.com/pgadmin-org/pgadmin4/issues/9902) -  Fix symlink-based path traversal in the file manager (CVE-2026-7819). Reported by Fernando Bortotti.<br>
[Issue #9904](https://github.com/pgadmin-org/pgadmin4/issues/9904) -  Fix account-lockout bypass on Flask-Security's default `/login` view by overriding `User.is_active` and `User.is_locked()` so the `locked` field is honored on every authentication path (CVE-2026-7820). Reported by Fernando Bortotti.<br>

# Additional changes (no associated issue)

The commits below did not have a dedicated GitHub issue. They are listed here for transparency.

## Bug fixes

`1518b0828` - Restore the SERVER_MODE python-test path and fix two endpoint regressions surfaced by it.<br>
`d57acce35` - Harden validation, preference, and connection-params paths against pre-existing edge cases.<br>

## Test-suite stability

`a11d289bd` - Harden `click_modal` backdrop wait and `open_query_tool` stale-element retry in feature tests.<br>
`a50a553b0` - Feature tests use `sys.executable`; sync `yarn.lock` to `package.json`.<br>
`0fad04de8` - PSQL socket tests use the authenticated tester; the role-dependencies test skips cleanly on auth failure.<br>
`1f7194924` - Harden six regression tests against environmental drift.<br>
`dc61039e9` - Quote the username in the views/mview test helper for dotted local roles.<br>
`9b29bc203` - Quote the username in the types/compound-triggers test helpers for dotted local roles.<br>
`504775de8` - Quote the username in the user-mappings test helper for dotted local roles.<br>
`208541cc4` - `ImportExportServersTestCase` uses `sys.executable` instead of a bare `python` and surfaces subprocess errors instead of swallowing them as a misleading JSON-parse failure.<br>

## Refactoring

`6f4f28def` - Factor the WTForms-error-to-JSON conversion into a helper and drop a dead import.<br>

## Documentation

`9923eefca` - Clarify in `login.rst` and `ldap.rst` that `MAX_LOGIN_ATTEMPTS` applies only to the `INTERNAL` authentication source. Operators using LDAP, OAuth2, Kerberos, or Webserver auth should rely on the upstream identity provider's lockout policy and reverse-proxy request rate-limiting.<br>

## Dependencies

Non-breaking `dependabot` updates aggregated for v9.15.

Python:

`boto3` 1.42 -> 1.43 (`python_version > '3.9'`; Python 3.9 stays on 1.42)<br>
`cryptography` 46.0 -> 47.0<br>
`psycopg` 3.3.3 -> 3.3.4 (`python_version >= '3.10'`)<br>
`pycodestyle` >=2.5.0 -> >=2.14.0<br>
`requests` >=2.21.0 -> >=2.33.1<br>
`safety` >=1.9.0 -> >=3.7.0<br>
`testtools` 2.8.7 -> 2.9.1<br>
`typer` 0.24 -> 0.25 (`python > '3.9'`)<br>

JavaScript (`web/`):

`@tanstack/react-query` 5.90 -> 5.100.5<br>
`axios` 1.15.2 -> 1.16.0<br>
`moment-timezone` 0.6.0 -> 0.6.2<br>
`postcss` 8.5.6 -> 8.5.12<br>

JavaScript (`runtime/`):

`axios` 1.15.2 -> 1.16.0<br>
`electron` 41.3.0 -> 41.5.0<br>
`eslint` 10.2.1 -> 10.3.0<br>
`globals` 17.5.0 -> 17.6.0<br>
