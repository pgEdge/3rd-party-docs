<a id="ddl-foreign-data"></a>

## Foreign Data


 PostgreSQL implements portions of the SQL/MED specification, allowing you to access data that resides outside PostgreSQL using regular SQL queries. Such data is referred to as *foreign data*. (Note that this usage is not to be confused with foreign keys, which are a type of constraint within the database.)


 Foreign data is accessed with help from a *foreign data wrapper*. A foreign data wrapper is a library that can communicate with an external data source, hiding the details of connecting to the data source and obtaining data from it. There are some foreign data wrappers available as `contrib` modules; see [Additional Supplied Modules and Extensions](../../appendixes/additional-supplied-modules-and-extensions/index.md#contrib). Other kinds of foreign data wrappers might be found as third party products. If none of the existing foreign data wrappers suit your needs, you can write your own; see [Writing a Foreign Data Wrapper](../../internals/writing-a-foreign-data-wrapper/index.md#fdwhandler).


 To access foreign data, you need to create a *foreign server* object, which defines how to connect to a particular external data source according to the set of options used by its supporting foreign data wrapper. Then you need to create one or more *foreign tables*, which define the structure of the remote data. A foreign table can be used in queries just like a normal table, but a foreign table has no storage in the PostgreSQL server. Whenever it is used, PostgreSQL asks the foreign data wrapper to fetch data from the external source, or transmit data to the external source in the case of update commands.


 Accessing remote data may require authenticating to the external data source. This information can be provided by a *user mapping*, which can provide additional data such as user names and passwords based on the current PostgreSQL role.


 For additional information, see [sql-createforeigndatawrapper](../../reference/sql-commands/create-foreign-data-wrapper.md#sql-createforeigndatawrapper), [sql-createserver](../../reference/sql-commands/create-server.md#sql-createserver), [sql-createusermapping](../../reference/sql-commands/create-user-mapping.md#sql-createusermapping), [sql-createforeigntable](../../reference/sql-commands/create-foreign-table.md#sql-createforeigntable), and [sql-importforeignschema](../../reference/sql-commands/import-foreign-schema.md#sql-importforeignschema).
