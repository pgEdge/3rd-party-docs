<a id="sql-altertablespace"></a>

# ALTER TABLESPACE

change the definition of a tablespace

## Synopsis


```

ALTER TABLESPACE NAME RENAME TO NEW_NAME
ALTER TABLESPACE NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER TABLESPACE NAME SET ( TABLESPACE_OPTION = VALUE [, ... ] )
ALTER TABLESPACE NAME RESET ( TABLESPACE_OPTION [, ... ] )
```


## Description


 `ALTER TABLESPACE` can be used to change the definition of a tablespace.


 You must own the tablespace to change the definition of a tablespace. To alter the owner, you must also be able to `SET ROLE` to the new owning role. (Note that superusers have these privileges automatically.)


## Parameters


*name*
:   The name of an existing tablespace.

*new_name*
:   The new name of the tablespace. The new name cannot begin with `pg_`, as such names are reserved for system tablespaces.

*new_owner*
:   The new owner of the tablespace.

*tablespace_option*
:   A tablespace parameter to be set or reset. Currently, the only available parameters are `seq_page_cost`, `random_page_cost`, `effective_io_concurrency` and `maintenance_io_concurrency`. Setting these values for a particular tablespace will override the planner's usual estimate of the cost of reading pages from tables in that tablespace, and the executor's prefetching behavior, as established by the configuration parameters of the same name (see [seq_page_cost](../../server-administration/server-configuration/query-planning.md#guc-seq-page-cost), [random_page_cost](../../server-administration/server-configuration/query-planning.md#guc-random-page-cost), [effective_io_concurrency](../../server-administration/server-configuration/resource-consumption.md#guc-effective-io-concurrency), [maintenance_io_concurrency](../../server-administration/server-configuration/resource-consumption.md#guc-maintenance-io-concurrency)). This may be useful if one tablespace is located on a disk which is faster or slower than the remainder of the I/O subsystem.


## Examples


 Rename tablespace `index_space` to `fast_raid`:

```sql

ALTER TABLESPACE index_space RENAME TO fast_raid;
```


 Change the owner of tablespace `index_space`:

```sql

ALTER TABLESPACE index_space OWNER TO mary;
```


## Compatibility


 There is no `ALTER TABLESPACE` statement in the SQL standard.


## See Also
  [sql-createtablespace](create-tablespace.md#sql-createtablespace), [sql-droptablespace](drop-tablespace.md#sql-droptablespace)
