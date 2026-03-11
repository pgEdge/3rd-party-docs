## Environment Variables { #libpq-envars }


 The following environment variables can be used to select default connection parameter values, which will be used by [PQconnectdb](database-connection-control-functions.md#libpq-PQconnectdb), [PQsetdbLogin](database-connection-control-functions.md#libpq-PQsetdbLogin) and [PQsetdb](database-connection-control-functions.md#libpq-PQsetdb) if no value is directly specified by the calling code. These are useful to avoid hard-coding database connection information into simple client applications, for example.

-   `PGHOST` behaves the same as the [host](database-connection-control-functions.md#libpq-connect-host) connection parameter.
-   `PGSSLNEGOTIATION` behaves the same as the [sslnegotiation](database-connection-control-functions.md#libpq-connect-sslnegotiation) connection parameter.
-   `PGHOSTADDR` behaves the same as the [hostaddr](database-connection-control-functions.md#libpq-connect-hostaddr) connection parameter. This can be set instead of or in addition to `PGHOST` to avoid DNS lookup overhead.
-   `PGPORT` behaves the same as the [port](database-connection-control-functions.md#libpq-connect-port) connection parameter.
-   `PGDATABASE` behaves the same as the [dbname](database-connection-control-functions.md#libpq-connect-dbname) connection parameter.
-   `PGUSER` behaves the same as the [user](database-connection-control-functions.md#libpq-connect-user) connection parameter.
-   `PGPASSWORD` behaves the same as the [password](database-connection-control-functions.md#libpq-connect-password) connection parameter. Use of this environment variable is not recommended for security reasons, as some operating systems allow non-root users to see process environment variables via ps; instead consider using a password file (see [The Password File](the-password-file.md#libpq-pgpass)).
-   `PGPASSFILE` behaves the same as the [passfile](database-connection-control-functions.md#libpq-connect-passfile) connection parameter.
-   `PGREQUIREAUTH` behaves the same as the [require_auth](database-connection-control-functions.md#libpq-connect-require-auth) connection parameter.
-   `PGCHANNELBINDING` behaves the same as the [channel_binding](database-connection-control-functions.md#libpq-connect-channel-binding) connection parameter.
-   `PGSERVICE` behaves the same as the [service](database-connection-control-functions.md#libpq-connect-service) connection parameter.
-   `PGSERVICEFILE` behaves the same as the [servicefile](database-connection-control-functions.md#libpq-connect-servicefile) connection parameter.
-   `PGOPTIONS` behaves the same as the [options](database-connection-control-functions.md#libpq-connect-options) connection parameter.
-   `PGAPPNAME` behaves the same as the [application_name](database-connection-control-functions.md#libpq-connect-application-name) connection parameter.
-   `PGSSLMODE` behaves the same as the [sslmode](database-connection-control-functions.md#libpq-connect-sslmode) connection parameter.
-   `PGREQUIRESSL` behaves the same as the [requiressl](database-connection-control-functions.md#libpq-connect-requiressl) connection parameter. This environment variable is deprecated in favor of the `PGSSLMODE` variable; setting both variables suppresses the effect of this one.
-   `PGSSLCOMPRESSION` behaves the same as the [sslcompression](database-connection-control-functions.md#libpq-connect-sslcompression) connection parameter.
-   `PGSSLCERT` behaves the same as the [sslcert](database-connection-control-functions.md#libpq-connect-sslcert) connection parameter.
-   `PGSSLKEY` behaves the same as the [sslkey](database-connection-control-functions.md#libpq-connect-sslkey) connection parameter.
-   `PGSSLCERTMODE` behaves the same as the [sslcertmode](database-connection-control-functions.md#libpq-connect-sslcertmode) connection parameter.
-   `PGSSLROOTCERT` behaves the same as the [sslrootcert](database-connection-control-functions.md#libpq-connect-sslrootcert) connection parameter.
-   `PGSSLCRL` behaves the same as the [sslcrl](database-connection-control-functions.md#libpq-connect-sslcrl) connection parameter.
-   `PGSSLCRLDIR` behaves the same as the [sslcrldir](database-connection-control-functions.md#libpq-connect-sslcrldir) connection parameter.
-   `PGSSLSNI` behaves the same as the [sslsni](database-connection-control-functions.md#libpq-connect-sslsni) connection parameter.
-   `PGREQUIREPEER` behaves the same as the [requirepeer](database-connection-control-functions.md#libpq-connect-requirepeer) connection parameter.
-   `PGSSLMINPROTOCOLVERSION` behaves the same as the [ssl_min_protocol_version](database-connection-control-functions.md#libpq-connect-ssl-min-protocol-version) connection parameter.
-   `PGSSLMAXPROTOCOLVERSION` behaves the same as the [ssl_max_protocol_version](database-connection-control-functions.md#libpq-connect-ssl-max-protocol-version) connection parameter.
-   `PGGSSENCMODE` behaves the same as the [gssencmode](database-connection-control-functions.md#libpq-connect-gssencmode) connection parameter.
-   `PGKRBSRVNAME` behaves the same as the [krbsrvname](database-connection-control-functions.md#libpq-connect-krbsrvname) connection parameter.
-   `PGGSSLIB` behaves the same as the [gsslib](database-connection-control-functions.md#libpq-connect-gsslib) connection parameter.
-   `PGGSSDELEGATION` behaves the same as the [gssdelegation](database-connection-control-functions.md#libpq-connect-gssdelegation) connection parameter.
-   `PGCONNECT_TIMEOUT` behaves the same as the [connect_timeout](database-connection-control-functions.md#libpq-connect-connect-timeout) connection parameter.
-   `PGCLIENTENCODING` behaves the same as the [client_encoding](database-connection-control-functions.md#libpq-connect-client-encoding) connection parameter.
-   `PGTARGETSESSIONATTRS` behaves the same as the [target_session_attrs](database-connection-control-functions.md#libpq-connect-target-session-attrs) connection parameter.
-   `PGLOADBALANCEHOSTS` behaves the same as the [load_balance_hosts](database-connection-control-functions.md#libpq-connect-load-balance-hosts) connection parameter.
-   `PGMINPROTOCOLVERSION` behaves the same as the [min_protocol_version](database-connection-control-functions.md#libpq-connect-min-protocol-version) connection parameter.
-   `PGMAXPROTOCOLVERSION` behaves the same as the [max_protocol_version](database-connection-control-functions.md#libpq-connect-max-protocol-version) connection parameter.


 The following environment variables can be used to specify default behavior for each PostgreSQL session. (See also the [sql-alterrole](../../reference/sql-commands/alter-role.md#sql-alterrole) and [sql-alterdatabase](../../reference/sql-commands/alter-database.md#sql-alterdatabase) commands for ways to set default behavior on a per-user or per-database basis.)

-   `PGDATESTYLE` sets the default style of date/time representation. (Equivalent to `SET datestyle TO ...`.)
-   `PGTZ` sets the default time zone. (Equivalent to `SET timezone TO ...`.)
-   `PGGEQO` sets the default mode for the genetic query optimizer. (Equivalent to `SET geqo TO ...`.)
 Refer to the SQL command [sql-set](../../reference/sql-commands/set.md#sql-set) for information on correct values for these environment variables.


 The following environment variables determine internal behavior of libpq; they override compiled-in defaults.

-   `PGSYSCONFDIR` sets the directory containing the `pg_service.conf` file and in a future version possibly other system-wide configuration files.
-   `PGLOCALEDIR` sets the directory containing the `locale` files for message localization.
