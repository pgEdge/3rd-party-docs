<a id="manage-ag-overview"></a>

## Overview


 A small number of objects, like role, database, and tablespace names, are defined at the cluster level and stored in the `pg_global` tablespace. Inside the cluster are multiple databases, which are isolated from each other but can access cluster-level objects. Inside each database are multiple schemas, which contain objects like tables and functions. So the full hierarchy is: cluster, database, schema, table (or some other kind of object, such as a function).


 When connecting to the database server, a client must specify the database name in its connection request. It is not possible to access more than one database per connection. However, clients can open multiple connections to the same database, or different databases. Database-level security has two components: access control (see [The `pg_hba.conf` File](../client-authentication/the-pg_hba-conf-file.md#auth-pg-hba-conf)), managed at the connection level, and authorization control (see [Privileges](../../the-sql-language/data-definition/privileges.md#ddl-priv)), managed via the grant system. Foreign data wrappers (see [postgres_fdw](../../appendixes/additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw)) allow for objects within one database to act as proxies for objects in other database or clusters. The older dblink module (see [dblink](../../appendixes/additional-supplied-modules-and-extensions/dblink-connect-to-other-postgresql-databases.md#dblink)) provides a similar capability. By default, all users can connect to all databases using all connection methods.


 If one PostgreSQL server cluster is planned to contain unrelated projects or users that should be, for the most part, unaware of each other, it is recommended to put them into separate databases and adjust authorizations and access controls accordingly. If the projects or users are interrelated, and thus should be able to use each other's resources, they should be put in the same database but probably into separate schemas; this provides a modular structure with namespace isolation and authorization control. More information about managing schemas is in [Schemas](../../the-sql-language/data-definition/schemas.md#ddl-schemas).


 While multiple databases can be created within a single cluster, it is advised to consider carefully whether the benefits outweigh the risks and limitations. In particular, the impact that having a shared WAL (see [Reliability and the Write-Ahead Log](../reliability-and-the-write-ahead-log/index.md#wal)) has on backup and recovery options. While individual databases in the cluster are isolated when considered from the user's perspective, they are closely bound from the database administrator's point-of-view.


 Databases are created with the `CREATE DATABASE` command (see [Creating a Database](creating-a-database.md#manage-ag-createdb)) and destroyed with the `DROP DATABASE` command (see [Destroying a Database](destroying-a-database.md#manage-ag-dropdb)). To determine the set of existing databases, examine the `pg_database` system catalog, for example

```

SELECT datname FROM pg_database;
```
 The [app-psql](../../reference/postgresql-client-applications/psql.md#app-psql) program's `\l` meta-command and `-l` command-line option are also useful for listing the existing databases.


!!! note

    The SQL standard calls databases “catalogs”, but there is no difference in practice.
