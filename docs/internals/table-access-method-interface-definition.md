<a id="tableam"></a>

# Table Access Method Interface Definition


 This chapter explains the interface between the core PostgreSQL system and *table access methods*, which manage the storage for tables. The core system knows little about these access methods beyond what is specified here, so it is possible to develop entirely new access method types by writing add-on code.


 Each table access method is described by a row in the [`pg_am`](system-catalogs/pg_am.md#catalog-pg-am) system catalog. The `pg_am` entry specifies a name and a *handler function* for the table access method. These entries can be created and deleted using the [sql-create-access-method](../reference/sql-commands/create-access-method.md#sql-create-access-method) and [sql-drop-access-method](../reference/sql-commands/drop-access-method.md#sql-drop-access-method) SQL commands.


 A table access method handler function must be declared to accept a single argument of type `internal` and to return the pseudo-type `table_am_handler`. The argument is a dummy value that simply serves to prevent handler functions from being called directly from SQL commands.


 Here is how an extension SQL script file might create a table access method handler:


```sql

CREATE OR REPLACE FUNCTION my_tableam_handler(internal)
  RETURNS table_am_handler AS 'my_extension', 'my_tableam_handler'
  LANGUAGE C STRICT;

CREATE ACCESS METHOD myam TYPE TABLE HANDLER my_tableam_handler;
```


 The result of the function must be a pointer to a struct of type `TableAmRoutine`, which contains everything that the core code needs to know to make use of the table access method. The return value needs to be of server lifetime, which is typically achieved by defining it as a `static const` variable in global scope.


 Here is how a source file with the table access method handler might look like:


```

#include "postgres.h"

#include "access/tableam.h"
#include "fmgr.h"

PG_MODULE_MAGIC;

static const TableAmRoutine my_tableam_methods = {
    .type = T_TableAmRoutine,

    /* Methods of TableAmRoutine omitted from example, add them here. */
};

PG_FUNCTION_INFO_V1(my_tableam_handler);

Datum
my_tableam_handler(PG_FUNCTION_ARGS)
{
    PG_RETURN_POINTER(&my_tableam_methods);
}
```


 The `TableAmRoutine` struct, also called the access method's *API struct*, defines the behavior of the access method using callbacks. These callbacks are pointers to plain C functions and are not visible or callable at the SQL level. All the callbacks and their behavior is defined in the `TableAmRoutine` structure (with comments inside the struct defining the requirements for callbacks). Most callbacks have wrapper functions, which are documented from the point of view of a user (rather than an implementor) of the table access method. For details, please refer to the [`src/include/access/tableam.h`](https://git.postgresql.org/gitweb/?p=postgresql.git;a=blob;f=src/include/access/tableam.h;hb=HEAD) file.


 To implement an access method, an implementor will typically need to implement an AM-specific type of tuple table slot (see [`src/include/executor/tuptable.h`](https://git.postgresql.org/gitweb/?p=postgresql.git;a=blob;f=src/include/executor/tuptable.h;hb=HEAD)), which allows code outside the access method to hold references to tuples of the AM, and to access the columns of the tuple.


 Currently, the way an AM actually stores data is fairly unconstrained. For example, it's possible, but not required, to use postgres' shared buffer cache. In case it is used, it likely makes sense to use PostgreSQL's standard page layout as described in [Database Page Layout](database-physical-storage/database-page-layout.md#storage-page-layout).


 One fairly large constraint of the table access method API is that, currently, if the AM wants to support modifications and/or indexes, it is necessary for each tuple to have a tuple identifier (TID) consisting of a block number and an item number (see also [Database Page Layout](database-physical-storage/database-page-layout.md#storage-page-layout)). It is not strictly necessary that the sub-parts of TIDs have the same meaning they e.g., have for `heap`, but if bitmap scan support is desired (it is optional), the block number needs to provide locality.


 For crash safety, an AM can use postgres' [WAL](../server-administration/reliability-and-the-write-ahead-log/index.md#wal), or a custom implementation. If WAL is chosen, either [Generic WAL Records](write-ahead-logging-for-extensions/generic-wal-records.md#generic-wal) can be used, or a [Custom WAL Resource Manager](write-ahead-logging-for-extensions/custom-wal-resource-managers.md#custom-rmgr) can be implemented.


 To implement transactional support in a manner that allows different table access methods be accessed within a single transaction, it likely is necessary to closely integrate with the machinery in `src/backend/access/transam/xlog.c`.


 Any developer of a new `table access method` can refer to the existing `heap` implementation present in `src/backend/access/heap/heapam_handler.c` for details of its implementation.
