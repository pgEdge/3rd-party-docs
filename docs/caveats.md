# Caveats

Object renames are logged under the name they were renamed to. For example, renaming a table will produce the following result:
```
ALTER TABLE test RENAME TO test2;

AUDIT: SESSION,36,1,DDL,ALTER TABLE,TABLE,public.test2,ALTER TABLE test RENAME TO test2,<not logged>
```
It is possible to have a command logged more than once. For example, when a table is created with a primary key specified at creation time the index for the primary key will be logged independently and another audit log will be made for the index under the create entry. The multiple entries will however be contained within one statement ID.

Autovacuum and Autoanalyze are not logged.

Statements that are executed after a transaction enters an aborted state will not be audit logged. However, the statement that caused the error and any subsequent statements executed in the aborted transaction will be logged as ERRORs by the standard logging facility.

It is not possible to reliably audit superusers with pgAudit. One solution is to restrict access to superuser accounts and use the [set_user](https://github.com/pgaudit/set_user) extension to escalate permissions when required.

