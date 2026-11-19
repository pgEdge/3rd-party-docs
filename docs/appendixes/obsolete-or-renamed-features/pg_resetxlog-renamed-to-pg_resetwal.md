<a id="app-pgresetxlog"></a>

## `pg_resetxlog` renamed to `pg_resetwal`


 PostgreSQL 9.6 and below provided a command named `pg_resetxlog`  to reset the write-ahead-log (WAL) files. This command was renamed to `pg_resetwal`, see [app-pgresetwal](../../reference/postgresql-server-applications/pg_resetwal.md#app-pgresetwal) for documentation of `pg_resetwal` and see [the release notes for PostgreSQL 10](../release-notes/prior-releases.md#release-prior) for details on this change.
