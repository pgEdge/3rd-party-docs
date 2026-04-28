<a id="configuration"></a>

# Configuration

Configuration parameters can be provided via:

- [Config File](#file_config).

- [Environment Variables](#env_variables_config), overriding values from the config file.

- [In-Database Configuration](#in_db_config), overriding values from both the config file and environment variables.

Using [Configuration Reloading](#config_reloading) you can modify the parameters without restarting the server.

## Minimum parameters

The server is able to start without any config parameters, but it won't be able to serve requests unless it has [a role to serve anonymous requests with](#db-anon-role) - or [a secret to use for JWT authentication](#jwt-secret).
<a id="file_config"></a>

## Config File

There is no predefined location for the config file, you must specify the file path as the one and only argument to the server:

```bash
./postgrest /path/to/postgrest.conf
```

The configuration file must contain a set of key value pairs:

```
# postgrest.conf

# The standard connection URI format, documented at
# https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
db-uri       = "postgres://user:pass@host:5432/dbname"

# The database role to use when no client authentication is provided.
# Should differ from authenticator
db-anon-role = "anon"

# The secret to verify the JWT for authenticated requests with.
# Needs to be 32 characters minimum.
jwt-secret           = "reallyreallyreallyreallyverysafe"
jwt-secret-is-base64 = false

# Port the postgrest process is listening on for http requests
server-port = 3000
```

You can run `postgrest --example` to display all possible configuration parameters and how to use them in a configuration file.
<a id="env_variables_config"></a>

## Environment Variables

Environment variables are capitalized, have a `PGRST_` prefix, and use underscores. For example: `PGRST_DB_URI` corresponds to `db-uri` and `PGRST_APP_SETTINGS_*` to `app.settings.*`.

[libpq environment variables](https://www.postgresql.org/docs/current/libpq-envars.html) are also supported for constructing the connection string, see [db-uri](#db-uri).

See the full list of environment variable names on [List of parameters](#config_full_list).
<a id="in_db_config"></a>

## In-Database Configuration

You can also configure the server with database settings by using a [pre-config](#db-pre-config) function. For example, you can configure [db-schemas](#db-schemas) and [jwt-secret](#jwt-secret) like this:

```
# postgrest.conf

db-pre-config  = "postgrest.pre_config"

# or env vars

PGRST_DB_PRE_CONFIG = "postgrest.pre_config"
```

```postgres
-- create a dedicated schema, hidden from the API
create schema postgrest;
-- grant usage on this schema to the authenticator
grant usage on schema postgrest to authenticator;

-- the function can configure postgREST by using set_config
create or replace function postgrest.pre_config()
returns void as $$
  select
      set_config('pgrst.db_schemas', 'schema1, schema2', true)
    , set_config('pgrst.jwt_secret', 'REALLYREALLYREALLYREALLYVERYSAFE', true);
$$ language sql;
```

Note that underscores(`_`) need to be used instead of dashes(`-`) for the in-database config parameters. See the full list of in-database names on [List of parameters](#config_full_list).

You can disable the in-database configuration by setting [db-config](#db-config) to `false`.

!!! note

    For backwards compatibility, you can do in-db config by modifying the [authenticator role](auth.md#roles). This is no longer recommended as it requires SUPERUSER.

    ```postgresql
    ALTER ROLE authenticator SET pgrst.db_schemas = "tenant1, tenant2, tenant3"
    ALTER ROLE authenticator IN DATABASE <your_database_name> SET pgrst.db_schemas = "tenant4, tenant5" -- database-specific setting, overrides the previous setting
    ```
<a id="config_reloading"></a>

## Configuration Reloading

It's possible to reload PostgREST's configuration without restarting the server. You can do this [via signal](#config_reloading_signal) or [via notification](#config_reloading_notify).

- Any modification to the [Config File](#file_config) will be applied during reload.

- Any modification to the [In-Database Configuration](#in_db_config) will be applied during reload.

- Not all settings are reloadable, see the reloadable list on [List of parameters](#config_full_list).

- It's not possible to change [Environment Variables](#env_variables_config) for a running process, hence reloading a Docker container configuration will not work. In these cases, you can restart the process or use [In-Database Configuration](#in_db_config).
<a id="config_reloading_signal"></a>

### Configuration Reload with signal

To reload the configuration via signal, send a SIGUSR2 signal to the server process.

```bash
killall -SIGUSR2 postgrest
```
<a id="config_reloading_notify"></a>

### Configuration Reload with NOTIFY

To reload the configuration from within the database, you can use the `NOTIFY` command. See [Listener](listener.md#listener).

```postgresql
NOTIFY pgrst, 'reload config'
```
<a id="config_full_list"></a>

## List of parameters
<a id="admin-server-host"></a>

### admin-server-host

=============== ======================= **Type**        String **Default**     `server-host` value **Reloadable**  N **Environment** PGRST_ADMIN_SERVER_HOST **In-Database** `n/a` =============== =======================

Specifies the host for the [Admin Server](admin_server.md#admin_server). Defaults to [server-host](#server-host) value.
<a id="admin-server-port"></a>

### admin-server-port

=============== ======================= **Type**        Int **Default**     `n/a` **Reloadable**  N **Environment** PGRST_ADMIN_SERVER_PORT **In-Database** `n/a` =============== =======================

Specifies the port for the [Admin Server](admin_server.md#admin_server). Cannot be equal to [server-port](#server-port).
<a id="app.settings.*"></a>

### app.settings.*

=============== ======================= **Type**        String **Default**     `n/a` **Reloadable**  & **Environment** PGRST_APP_SETTINGS_* **In-Database** `n/a` =============== =======================

Arbitrary settings that can be used to pass in secret keys directly as strings, or via OS environment variables. For instance: `app.settings.jwt_secret = "$(MYAPP_JWT_SECRET)"` will take `MYAPP_JWT_SECRET` from the environment and make it available to PostgreSQL functions as `current_setting('app.settings.jwt_secret')`.

When using the environment variable `PGRST_APP_SETTINGS_*` form, the remainder of the variable is used as the new name. Case is not important : `PGRST_APP_SETTINGS_MY_ENV_VARIABLE=some_value` can be accessed in postgres as `current_setting('app.settings.my_env_variable')`.

The [current_setting` function has `an optional boolean second](https://www.postgresql.org/docs/current/functions-admin.html#FUNCTIONS-ADMIN-SET) argument to avoid it from raising an error if the value was not defined. Default values to `app.settings` can then be given by combining this argument with `coalesce` and `nullif` : `coalesce(nullif(current_setting('app.settings.my_custom_variable', true), ''), 'default value')`. The use of `nullif` is necessary because if set in a transaction, the setting is sometimes not "rolled back" to `null`. See also [this section](transactions.md#guc_req_headers_cookies_claims) for more information on this behaviour.
<a id="db-aggregates-enabled"></a>

### db-aggregates-enabled

=============== ======================= **Type**        Boolean **Default**     False **Reloadable**  Y **Environment** PGRST_DB_AGGREGATES_ENABLED **In-Database** pgrst.db_aggregates_enabled =============== =======================

When this is set to `true`, the use of [Aggregate Functions](../api/references/api/aggregate_functions.md#aggregate_functions) is allowed.

It is recommended that this be set to `false` unless proper safeguards are in place to prevent potential performance problems from arising. For example, it is possible that a user may request the `max()` of an unindexed column in a table with millions of rows. At best, this would result in a slow query, and at worst, it could be abused to prevent other users from accessing your API (i.e. a form of denial-of-service attack.)

Proper safeguards could include:

- Use of a statement timeout. See [Impersonated Role Settings](transactions.md#impersonated_settings).

- Use of the [pg_plan_filter extension](https://github.com/pgexperts/pg_plan_filter) to block excessively expensive queries.
<a id="db-anon-role"></a>

### db-anon-role

=============== ======================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_DB_ANON_ROLE **In-Database** pgrst.db_anon_role =============== =======================

The database role to use when executing commands on behalf of unauthenticated clients. For more information, see [Overview of role system](auth.md#roles).

When unset anonymous access will be blocked.
<a id="db-channel"></a>

### db-channel

=============== ======================= **Type**        String **Default**     pgrst **Reloadable**  Y **Environment** PGRST_DB_CHANNEL **In-Database** `n/a` =============== =======================

The name of the notification channel that PostgREST uses for [Schema Cache Reloading with NOTIFY](schema_cache.md#schema_reloading_notify) and [Configuration Reload with NOTIFY](#config_reloading_notify).
<a id="db-channel-enabled"></a>

### db-channel-enabled

=============== ======================= **Type**        Boolean **Default**     True **Reloadable**  Y **Environment** PGRST_DB_CHANNEL_ENABLED **In-Database** `n/a` =============== =======================

When this is set to `true`, the notification channel specified in [db-channel](#db-channel) is enabled.

You should set this to `false` when using PostgresSQL behind an external connection pooler such as PgBouncer working in transaction pooling mode. See [this section](connection_pool.md#external_connection_poolers) for more information.
<a id="db-config"></a>

### db-config

=============== ======================= **Type**        Boolean **Default**     True **Reloadable**  Y **Environment** PGRST_DB_CONFIG **In-Database** `n/a` =============== =======================

Enables the in-database configuration.
<a id="db-pre-config"></a>

### db-pre-config

=============== ======================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_DB_PRE_CONFIG **In-Database** pgrst.db_pre_config =============== =======================

Name of the function that does [In-Database Configuration](#in_db_config).
<a id="db-extra-search-path"></a>

### db-extra-search-path

=============== ========================== **Type**        String **Default**     public **Reloadable**  Y **Environment** PGRST_DB_EXTRA_SEARCH_PATH **In-Database** pgrst.db_extra_search_path =============== ==========================

Extra schemas to add to the [search_path](https://www.postgresql.org/docs/current/ddl-schemas.html#DDL-SCHEMAS-PATH) of every request. These schemas tables, views and functions **don't get API endpoints**, they can only be referred from the database objects inside your [db-schemas](#db-schemas).

This parameter was meant to make it easier to use **PostgreSQL extensions** (like PostGIS) that are outside of the [db-schemas](#db-schemas).

Multiple schemas can be added in a comma-separated string, e.g. `public, extensions`.

!!! important

    We default this config to `public` because it is the most common schema used to install PostgreSQL extensions such as [PostGIS](../how-tos/working-with-postgresql-data-types.md#ww_postgis). You can disable this by setting this config to `""`.
<a id="db-hoisted-tx-settings"></a>

### db-hoisted-tx-settings

=============== ================================================================================== **Type**        String **Default**     statement_timeout, plan_filter.statement_cost_limit, default_transaction_isolation **Reloadable**  Y **Environment** PGRST_DB_HOISTED_TX_SETTINGS **In-Database** pgrst.db_hoisted_tx_settings =============== ==================================================================================

Hoisted settings are allowed to be applied as transaction-scoped function settings. Multiple settings can be added in a comma-separated string, e.g. `work_mem, statement_timeout`.
<a id="db-max-rows"></a>

### db-max-rows

=============== ========================== **Type**        Int **Default**     ∞ **Reloadable**  Y **Environment** PGRST_DB_MAX_ROWS **In-Database** pgrst.db_max_rows =============== ==========================

*For backwards compatibility, this config parameter is also available without prefix as "max-rows".*

A hard limit to the number of rows PostgREST will fetch from a view, table, or function. Limits payload size for accidental or malicious requests.
<a id="db-plan-enabled"></a>

### db-plan-enabled

=============== ========================== **Type**        Boolean **Default**     False **Reloadable**  Y **Environment** PGRST_DB_PLAN_ENABLED **In-Database** pgrst.db_plan_enabled =============== ==========================

When this is set to `true`, the execution plan of a request can be retrieved by using the `Accept: application/vnd.pgrst.plan` header. See [Execution plan](observability.md#explain_plan).
<a id="db-pool"></a>

### db-pool

=============== ========================== **Type**        Int **Default**     10 **Reloadable**  N **Environment** PGRST_DB_POOL **In-Database** n/a =============== ==========================

Number of maximum connections to keep open in PostgREST's database pool.
<a id="db-pool-acquisition-timeout"></a>

### db-pool-acquisition-timeout

=============== ================================= **Type**        Int **Default**     10 **Reloadable**  N **Environment** PGRST_DB_POOL_ACQUISITION_TIMEOUT **In-Database** `n/a` =============== =================================

Specifies the maximum time in seconds that the request will wait for the pool to free up a connection slot to the database.
<a id="db-pool-max-idletime"></a>

### db-pool-max-idletime

=============== ================================= **Type**        Int **Default**     30 **Reloadable**  N **Environment** PGRST_DB_POOL_MAX_IDLETIME **In-Database** `n/a` =============== =================================

*For backwards compatibility, this config parameter is also available as "db-pool-timeout".*

Time in seconds to close idle pool connections.
<a id="db-pool-max-lifetime"></a>

### db-pool-max-lifetime

=============== ================================= **Type**        Int **Default**     1800 **Reloadable**  N **Environment** PGRST_DB_POOL_MAX_LIFETIME **In-Database** `n/a` =============== =================================

Specifies the maximum time in seconds of an existing connection in the pool.
<a id="db-pool-automatic-recovery"></a>

### db-pool-automatic-recovery

=============== ================================= **Type**        Boolean **Default**     True **Reloadable**  Y **Environment** PGRST_DB_POOL_AUTOMATIC_RECOVERY **In-Database** `n/a` =============== =================================

Enables or disables connection retrying.

When disabled, PostgREST would terminate immediately after connection loss instead of retrying indefinitely. See [this section](connection_pool.md#automatic_recovery) for more information.
<a id="db-pre-request"></a>

### db-pre-request

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_DB_PRE_REQUEST **In-Database** pgrst.db_pre_request =============== =================================

*For backwards compatibility, this config parameter is also available without prefix as "pre-request".*

A schema-qualified function name to call right after the [Transaction-Scoped Settings](transactions.md#tx_settings) are set. See [Pre-Request](transactions.md#pre-request).
<a id="db-prepared-statements"></a>

### db-prepared-statements

=============== ================================= **Type**        Boolean **Default**     True **Reloadable**  Y **Environment** PGRST_DB_PREPARED_STATEMENTS **In-Database** pgrst.db_prepared_statements =============== =================================

Enables or disables prepared statements.

When disabled, the generated queries will be parameterized (invulnerable to SQL injection) but they will not be prepared (cached in the database session). Not using prepared statements will noticeably decrease performance, so it's recommended to always have this setting enabled.

You should only set this to `false` when using PostgresSQL behind an external connection pooler such as PgBouncer working in transaction pooling mode. See [this section](connection_pool.md#external_connection_poolers) for more information.
<a id="db-root-spec"></a>

### db-root-spec

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_DB_ROOT_SPEC **In-Database** pgrst.db_root_spec =============== =================================

Function to override the OpenAPI response. See [Overriding Full OpenAPI Response](../api/references/api/openapi.md#override_openapi).
<a id="db-schemas"></a>

### db-schemas

=============== ================================= **Type**        String **Default**     public **Reloadable**  Y **Environment** PGRST_DB_SCHEMAS **In-Database** pgrst.db_schemas =============== =================================

*For backwards compatibility, this config parameter is also available in singular as "db-schema".*

The list of database schemas to expose to clients. See [Schemas](../api/references/api/schemas.md#schemas).
<a id="db-tx-end"></a>

### db-tx-end

=============== ================================= **Type**        String **Default**     commit **Reloadable**  N **Environment** PGRST_DB_TX_END **In-Database** pgrst.db_tx_end =============== =================================

Specifies how to terminate the database transactions.

```bash
# The transaction is always committed
db-tx-end = "commit"

# The transaction is committed unless a "Prefer: tx=rollback" header is sent
db-tx-end = "commit-allow-override"

# The transaction is always rolled back
db-tx-end = "rollback"

# The transaction is rolled back unless a "Prefer: tx=commit" header is sent
db-tx-end = "rollback-allow-override"
```
<a id="db-uri"></a>

### db-uri

=============== ================================= **Type**        String **Default**     postgresql:// **Reloadable**  N **Environment** PGRST_DB_URI **In-Database** `n/a` =============== =================================

The standard [PostgreSQL connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING), there are different ways to specify it:

#### URI Format

```
"postgres://authenticator:mysecretpassword@localhost:5433/postgres?parameters=val"
```

- Under this format symbols and unusual characters in the password or other fields should be percent encoded to avoid a parse error.

- If enforcing an SSL connection to the database is required you can use [sslmode](https://www.postgresql.org/docs/current/libpq-ssl.html#LIBPQ-SSL-SSLMODE-STATEMENTS) in the URI, for example `postgres://user:pass@host:5432/dbname?sslmode=require`.

- The user with whom PostgREST connects to the database is also known as the `authenticator` role. For more information see [Overview of role system](auth.md#roles).

- When running PostgREST on the same machine as PostgreSQL, it is also possible to connect to the database using a [Unix socket](https://en.wikipedia.org/wiki/Unix_domain_socket) and the [Peer Authentication method](https://www.postgresql.org/docs/current/auth-peer.html) as an alternative to TCP/IP communication and authentication with a password, this also grants higher performance.  To do this you can omit the host and the password, e.g. `postgres://user@/dbname`, see the [libpq connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) documentation for more details.

#### Keyword/Value Format

```
"host=localhost port=5433 user=authenticator password=mysecretpassword dbname=postgres"
```

#### LIBPQ Environment Variables

```
PGHOST=localhost PGPORT=5433 PGUSER=authenticator PGDATABASE=postgres
```

Any parameter that is not set in the above formats is read from [libpq environment variables](https://www.postgresql.org/docs/current/libpq-envars.html). The default connection string is `postgresql://`, which reads **all** parameters from the environment.

#### External config file

Choosing a value for this parameter beginning with the at sign such as `@filename` (e.g. `@./configs/my-config`) loads the connection string out of an external file.
<a id="jwt-aud"></a>

### jwt-aud

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_JWT_AUD **In-Database** pgrst.jwt_aud =============== =================================

Specifies an audience for the JWT `aud` claim. See [aud validation](auth.md#jwt_aud).
<a id="jwt-role-claim-key"></a>

### jwt-role-claim-key

=============== ================================= **Type**        String **Default**     .role **Reloadable**  Y **Environment** PGRST_JWT_ROLE_CLAIM_KEY **In-Database** pgrst.jwt_role_claim_key =============== =================================

*For backwards compatibility, this config parameter is also available without prefix as "role-claim-key".*

See [JWT Role Extraction](auth.md#jwt_role_extract) on how to specify key paths and usage examples.
<a id="jwt-secret"></a>

### jwt-secret

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_JWT_SECRET **In-Database** pgrst.jwt_secret =============== =================================

The secret or [JSON Web Key (JWK) (or set)](https://datatracker.ietf.org/doc/html/rfc7517) used to decode JWT tokens clients provide for authentication. For security the key must be **at least 32 characters long**. If this parameter is not specified then PostgREST refuses authentication requests. Choosing a value for this parameter beginning with the at sign such as `@filename` loads the secret out of an external file. This is useful for automating deployments. Note that any binary secrets must be base64 encoded. Both symmetric and asymmetric cryptography are supported. For more info see [Asymmetric Keys](auth.md#asym_keys).

Choosing a value for this parameter beginning with the at sign such as `@filename` (e.g. `@./configs/my-config`) loads the secret out of an external file.

!!! warning

    Only when using the [Config File](#file_config), if the `jwt-secret` contains a `$` character by itself it will give errors. In this case, use `$$` and PostgREST will interpret it as a single `$` character.
<a id="jwt-secret-is-base64"></a>

### jwt-secret-is-base64

=============== ================================= **Type**        Boolean **Default**     False **Reloadable**  Y **Environment** PGRST_JWT_SECRET_IS_BASE64 **In-Database** pgrst.jwt_secret_is_base64 =============== =================================

When this is set to `true`, the value derived from `jwt-secret` will be treated as a base64 encoded secret.
<a id="jwt-cache-max-entries"></a>

### jwt-cache-max-entries

=============== ================================= **Type**        Int **Default**     1000 **Reloadable**  Y **Environment** PGRST_JWT_CACHE_MAX_ENTRIES **In-Database** pgrst.jwt_cache_max_entries =============== =================================

Maximum number of entries in JWT cache. The value `0` disables JWT caching. See [JWT Cache](auth.md#jwt_caching).
<a id="log-level"></a>

### log-level

=============== ================================= **Type**        String **Default**     error **Reloadable**  N **Environment** PGRST_LOG_LEVEL **In-Database** `n/a` =============== =================================

Specifies the level of information to be logged while running PostgREST.

```bash
 # Only startup and db connection recovery messages are logged
 log-level = "crit"

 # All the "crit" level events plus server errors (status 5xx) are logged
 log-level = "error"

 # All the "error" level events plus request errors (status 4xx) are logged
 log-level = "warn"

 # All the "warn" level events plus all requests (every status code) are logged
 log-level = "info"

 # All the above plus events for development purposes are logged
 # Logs connection pool events and the schema cache parsing time
 log-level = "debug"
```

Because currently there's no buffering for logging, the levels with minimal logging(`crit/error`) will increase throughput.
<a id="log-query"></a>

### log-query

=============== ================================= **Type**        Boolean **Default**     False **Reloadable**  Y **Environment** PGRST_LOG_QUERY **In-Database** `n/a` =============== =================================

Logs the SQL query for the corresponding request at the current [log-level](#log-level). See [SQL Query Logs](observability.md#sql_query_logs).
<a id="openapi-mode"></a>

### openapi-mode

=============== ================================= **Type**        String **Default**     follow-privileges **Reloadable**  Y **Environment** PGRST_OPENAPI_MODE **In-Database** pgrst.openapi_mode =============== =================================

Specifies how the OpenAPI output should be displayed.

```bash
# Follows the privileges of the JWT role claim (or from db-anon-role if the JWT is not sent)
# Shows information depending on the permissions that the role making the request has
openapi-mode = "follow-privileges"

# Ignores the privileges of the JWT role claim (or from db-anon-role if the JWT is not sent)
# Shows all the exposed information, regardless of the permissions that the role making the request has
openapi-mode = "ignore-privileges"

# Disables the OpenApi output altogether.
# Throws a `404 Not Found` error when accessing the API root path
openapi-mode = "disabled"
```
<a id="openapi-security-active"></a>

### openapi-security-active

=============== ================================= **Type**        Boolean **Default**     False **Reloadable**  Y **Environment** PGRST_OPENAPI_SECURITY_ACTIVE **In-Database** pgrst.openapi_security_active =============== =================================

When this is set to `true`, security options are included in the [OpenAPI output](../api/references/api/openapi.md#open-api).
<a id="openapi-server-proxy-uri"></a>

### openapi-server-proxy-uri

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  N **Environment** PGRST_OPENAPI_SERVER_PROXY_URI **In-Database** pgrst.openapi_server_proxy_uri =============== =================================

Overrides the base URL used within the OpenAPI self-documentation hosted at the API root path. Use a complete URI syntax `scheme:[//[user:password@]host[:port]][/]path[?query][#fragment]`. Ex. `https://postgrest.com`

```json
{
  "swagger": "2.0",
  "info": {
    "version": "0.4.3.0",
    "title": "PostgREST API",
    "description": "This is a dynamic API generated by PostgREST"
  },
  "host": "postgrest.com:443",
  "basePath": "/",
  "schemes": [
    "https"
  ]
}
```
<a id="server_cors_allowed_origins"></a>

### server-cors-allowed-origins

=============== =================================== **Type**        String **Default**     `n/a` **Reloadable**  N **Environment** PGRST_SERVER_CORS_ALLOWED_ORIGINS **In-Database** `pgrst.server_cors_allowed_origins` =============== ===================================

Specifies allowed CORS origins in this config. See cors.

When this is not set or set to `""`, PostgREST **accepts** CORS requests from any domain.
<a id="server-host"></a>

### server-host

=============== ================================= **Type**        String **Default**     !4 **Reloadable**  N **Environment** PGRST_SERVER_HOST **In-Database** `n/a` =============== =================================

Where to bind the PostgREST web server. In addition to the usual address options, PostgREST interprets these reserved addresses with special meanings:

- `*` - any IPv4 or IPv6 hostname

- `*4` - any IPv4 or IPv6 hostname, IPv4 preferred

- `4` - any IPv4 hostname

- `*6` - any IPv4 or IPv6 hostname, IPv6 preferred

- `6` - any IPv6 hostname

Examples:

```bash
server-host = "127.0.0.1"
```
<a id="server-port"></a>

### server-port

=============== ================================= **Type**        Int **Default**     3000 **Reloadable**  N **Environment** PGRST_SERVER_PORT **In-Database** `n/a` =============== =================================

The TCP port to bind the web server. Use `0` to automatically assign a port.
<a id="server-trace-header"></a>

### server-trace-header

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  Y **Environment** PGRST_SERVER_TRACE_HEADER **In-Database** pgrst.server_trace_header =============== =================================

The header name used to trace HTTP requests. See [Trace Header](observability.md#trace_header).
<a id="server-timing-enabled"></a>

### server-timing-enabled

=============== ================================= **Type**        Boolean **Default**     False **Reloadable**  Y **Environment** PGRST_SERVER_TIMING_ENABLED **In-Database** pgrst.server_timing_enabled =============== =================================

Enables the [Server-Timing](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Server-Timing) header. See [Server-Timing Header](observability.md#server-timing_header).
<a id="server-unix-socket"></a>

### server-unix-socket

=============== ================================= **Type**        String **Default**     `n/a` **Reloadable**  N **Environment** PGRST_SERVER_UNIX_SOCKET **In-Database** `n/a` =============== =================================

[Unix domain socket](https://en.wikipedia.org/wiki/Unix_domain_socket) where to bind the PostgREST web server. If specified, this takes precedence over [server-port](#server-port). Example:

```bash
server-unix-socket = "/tmp/pgrst.sock"
```
<a id="server-unix-socket-mode"></a>

### server-unix-socket-mode

=============== ================================= **Type**        String **Default**     660 **Reloadable**  N **Environment** PGRST_SERVER_UNIX_SOCKET_MODE **In-Database** `n/a` =============== =================================

[Unix file mode](https://en.wikipedia.org/wiki/File_system_permissions) to be set for the socket specified in [server-unix-socket](#server-unix-socket) Needs to be a valid octal between 600 and 777.

```bash
server-unix-socket-mode = "660"
```
