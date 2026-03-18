# Settings

Settings may be modified only by a superuser. Allowing normal users to change their settings would defeat the point of an audit log.

Settings can be specified globally (in `postgresql.conf` or using `ALTER SYSTEM ... SET`), at the database level (using `ALTER DATABASE ... SET`), or at the role level (using `ALTER ROLE ... SET`). Note that settings are not inherited through normal role inheritance and `SET ROLE` will not alter a user's pgAudit settings. This is a limitation of the roles system and not inherent to pgAudit.

The pgAudit extension must be loaded in [shared_preload_libraries](http://www.postgresql.org/docs/17/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES). Otherwise, an error will be raised at load time and no audit logging will occur.

In addition, `CREATE EXTENSION pgaudit` must be called before `pgaudit.log` is set to ensure proper pgaudit functionality. The extension installs event triggers which add additional auditing for DDL. pgAudit will work without the extension installed but DDL statements will not have information about the object type and name.

If the `pgaudit` extension is dropped and needs to be recreated then `pgaudit.log` must be unset first otherwise an error will be raised.

## pgaudit.log

Specifies which classes of statements will be logged by session audit logging. Possible values are:

- **READ**: `SELECT` and `COPY` when the source is a relation or a query.

- **WRITE**: `INSERT`, `UPDATE`, `DELETE`, `TRUNCATE`, and `COPY` when the destination is a relation.

- **FUNCTION**: Function calls and `DO` blocks.

- **ROLE**: Statements related to roles and privileges: `GRANT`, `REVOKE`, `CREATE/ALTER/DROP ROLE`.

- **DDL**: All `DDL` that is not included in the `ROLE` class.

- **MISC**: Miscellaneous commands, e.g. `DISCARD`, `FETCH`, `CHECKPOINT`, `VACUUM`, `SET`.

- **MISC_SET**: Miscellaneous `SET` commands, e.g. `SET ROLE`.

- **ALL**: Include all of the above.

Multiple classes can be provided using a comma-separated list and classes can be subtracted by prefacing the class with a `-` sign (see [Session Audit Logging](session-audit-logging.md)).

The default is `none`.

## pgaudit.log_catalog

Specifies that session logging should be enabled in the case where all relations in a statement are in pg_catalog. Disabling this setting will reduce noise in the log from tools like psql and PgAdmin that query the catalog heavily.

The default is `on`.

## pgaudit.log_client

Specifies whether log messages will be visible to a client process such as psql. This setting should generally be left disabled but may be useful for debugging or other purposes.

Note that `pgaudit.log_level` is only enabled when `pgaudit.log_client` is `on`.

The default is `off`.

## pgaudit.log_level

Specifies the log level that will be used for log entries (see [Message Severity Levels](http://www.postgresql.org/docs/17/runtime-config-logging.html#RUNTIME-CONFIG-SEVERITY-LEVELS) for valid levels) but note that `ERROR`, `FATAL`, and `PANIC` are not allowed). This setting is used for regression testing and may also be useful to end users for testing or other purposes.

Note that `pgaudit.log_level` is only enabled when `pgaudit.log_client` is `on`; otherwise the default will be used.

The default is `log`.

## pgaudit.log_parameter

Specifies that audit logging should include the parameters that were passed with the statement. When parameters are present they will be included in `CSV` format after the statement text.

The default is `off`.

## pgaudit.log_parameter_max_size

Specifies that parameter values longer than this setting (in bytes) should not be logged, but replaced with `<long param suppressed>`. This is set in bytes, not characters, so does not account for multi-byte characters in a text parameters's encoding. This setting has no effect if `log_parameter` is `off`. If this setting is 0 (the default), all parameters are logged regardless of length

The default is `0`.

## pgaudit.log_relation

Specifies whether session audit logging should create a separate log entry for each relation (`TABLE`, `VIEW`, etc.) referenced in a `SELECT` or `DML` statement. This is a useful shortcut for exhaustive logging without using object audit logging.

The default is `off`.

## pgaudit.log_rows

Specifies that audit logging should include the number of rows retrieved or affected by a statement. When enabled the rows field will be included after the parameter field.

The default is `off`.

## pgaudit.log_statement

Specifies whether logging will include the statement text and parameters (if enabled). Depending on requirements, an audit log might not require this and it makes the logs less verbose.

The default is `on`.

## pgaudit.log_statement_once

Specifies whether logging will include the statement text and parameters with the first log entry for a statement/substatement combination or with every entry. Enabling this setting will result in less verbose logging but may make it more difficult to determine the statement that generated a log entry, though the statement/substatement pair along with the process id should suffice to identify the statement text logged with a previous entry.

The default is `off`.

## pgaudit.role

Specifies the master role to use for object audit logging. Multiple audit roles can be defined by granting them to the master role. This allows multiple groups to be in charge of different aspects of audit logging.

There is no default.

