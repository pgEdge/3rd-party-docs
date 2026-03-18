# Format

Audit entries are written to the standard logging facility and contain the following columns in comma-separated format. Output is compliant CSV format only if the log line prefix portion of each log entry is removed.

- **AUDIT_TYPE** - `SESSION` or `OBJECT`.

- **STATEMENT_ID** - Unique statement ID for this session. Each statement ID represents a backend call. Statement IDs are sequential even if some statements are not logged. There may be multiple entries for a statement ID when more than one relation is logged.

- **SUBSTATEMENT_ID** - Sequential ID for each sub-statement within the main statement. For example, calling a function from a query. Sub-statement IDs are continuous even if some sub-statements are not logged. There may be multiple entries for a sub-statement ID when more than one relation is logged.

- **CLASS** - e.g. `READ`, `ROLE` (see [pgaudit.log](settings.md#pgauditlog)).

- **COMMAND** - e.g. `ALTER TABLE`, `SELECT`.

- **OBJECT_TYPE** - `TABLE`, `INDEX`, `VIEW`, etc. Available for `SELECT`, `DML` and most `DDL` statements.

- **OBJECT_NAME** - The fully-qualified object name (e.g. public.account). Available for `SELECT`, `DML` and most `DDL` statements.

- **STATEMENT** - Statement executed on the backend.

- **PARAMETER** - If `pgaudit.log_parameter` is set then this field will contain the statement parameters as quoted CSV or `<none>` if there are no parameters. Otherwise, the field is `<not logged>`.

Use [log_line_prefix](http://www.postgresql.org/docs/17/runtime-config-logging.html#GUC-LOG-LINE-PREFIX) to add any other fields that are needed to satisfy your audit log requirements. A typical log line prefix might be `'%m %u %d [%p]: '` which would provide the date/time, user name, database name, and process id for each audit log.

