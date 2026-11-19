<a id="manage-ag-createdb"></a>

## Creating a Database


 In order to create a database, the PostgreSQL server must be up and running (see [Starting the Database Server](../server-setup-and-operation/starting-the-database-server.md#server-start)).


 Databases are created with the SQL command [sql-createdatabase](../../reference/sql-commands/create-database.md#sql-createdatabase):

```

CREATE DATABASE NAME;
```
 where *name* follows the usual rules for SQL identifiers. The current role automatically becomes the owner of the new database. It is the privilege of the owner of a database to remove it later (which also removes all the objects in it, even if they have a different owner).


 The creation of databases is a restricted operation. See [Role Attributes](../database-roles/role-attributes.md#role-attributes) for how to grant permission.


 Since you need to be connected to the database server in order to execute the `CREATE DATABASE` command, the question remains how the *first* database at any given site can be created. The first database is always created by the `initdb` command when the data storage area is initialized. (See [Creating a Database Cluster](../server-setup-and-operation/creating-a-database-cluster.md#creating-cluster).) This database is called `postgres`. So to create the first “ordinary” database you can connect to `postgres`.


 Two additional databases, `template1` and `template0`, are also created during database cluster initialization. Whenever a new database is created within the cluster, `template1` is essentially cloned. This means that any changes you make in `template1` are propagated to all subsequently created databases. Because of this, avoid creating objects in `template1` unless you want them propagated to every newly created database. `template0` is meant as a pristine copy of the original contents of `template1`. It can be cloned instead of `template1` when it is important to make a database without any such site-local additions. More details appear in [Template Databases](template-databases.md#manage-ag-templatedbs).


 As a convenience, there is a program you can execute from the shell to create new databases, `createdb`.

```

createdb DBNAME
```
 `createdb` does no magic. It connects to the `postgres` database and issues the `CREATE DATABASE` command, exactly as described above. The [app-createdb](../../reference/postgresql-client-applications/createdb.md#app-createdb) reference page contains the invocation details. Note that `createdb` without any arguments will create a database with the current user name.


!!! note

    [Client Authentication](../client-authentication/index.md#client-authentication) contains information about how to restrict who can connect to a given database.


 Sometimes you want to create a database for someone else, and have them become the owner of the new database, so they can configure and manage it themselves. To achieve that, use one of the following commands:

```sql

CREATE DATABASE DBNAME OWNER ROLENAME;
```
 from the SQL environment, or:

```

createdb -O ROLENAME DBNAME
```
 from the shell. Only the superuser is allowed to create a database for someone else (that is, for a role you are not a member of).
