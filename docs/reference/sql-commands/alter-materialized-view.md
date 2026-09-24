<a id="sql-altermaterializedview"></a>

# ALTER MATERIALIZED VIEW

change the definition of a materialized view

## Synopsis


```

ALTER MATERIALIZED VIEW [ IF EXISTS ] NAME
    ACTION [, ... ]
ALTER MATERIALIZED VIEW NAME
    [ NO ] DEPENDS ON EXTENSION EXTENSION_NAME
ALTER MATERIALIZED VIEW [ IF EXISTS ] NAME
    RENAME [ COLUMN ] COLUMN_NAME TO NEW_COLUMN_NAME
ALTER MATERIALIZED VIEW [ IF EXISTS ] NAME
    RENAME TO NEW_NAME
ALTER MATERIALIZED VIEW [ IF EXISTS ] NAME
    SET SCHEMA NEW_SCHEMA
ALTER MATERIALIZED VIEW ALL IN TABLESPACE NAME [ OWNED BY ROLE_NAME [, ... ] ]
    SET TABLESPACE NEW_TABLESPACE [ NOWAIT ]

where ACTION is one of:

    ALTER [ COLUMN ] COLUMN_NAME SET STATISTICS INTEGER
    ALTER [ COLUMN ] COLUMN_NAME SET ( ATTRIBUTE_OPTION = VALUE [, ... ] )
    ALTER [ COLUMN ] COLUMN_NAME RESET ( ATTRIBUTE_OPTION [, ... ] )
    ALTER [ COLUMN ] COLUMN_NAME SET STORAGE { PLAIN | EXTERNAL | EXTENDED | MAIN | DEFAULT }
    ALTER [ COLUMN ] COLUMN_NAME SET COMPRESSION COMPRESSION_METHOD
    CLUSTER ON INDEX_NAME
    SET WITHOUT CLUSTER
    SET ACCESS METHOD NEW_ACCESS_METHOD
    SET TABLESPACE NEW_TABLESPACE
    SET ( STORAGE_PARAMETER [= VALUE] [, ... ] )
    RESET ( STORAGE_PARAMETER [, ... ] )
    OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
```


## Description


 `ALTER MATERIALIZED VIEW` changes various auxiliary properties of an existing materialized view.


 You must own the materialized view to use `ALTER MATERIALIZED VIEW`. To change a materialized view's schema, you must also have `CREATE` privilege on the new schema. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the materialized view's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the materialized view. However, a superuser can alter ownership of any view anyway.)


 The statement subforms and actions available for `ALTER MATERIALIZED VIEW` are a subset of those available for `ALTER TABLE`, and have the same meaning when used for materialized views. See the descriptions for [`ALTER TABLE`](alter-table.md#sql-altertable) for details.


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing materialized view.

*column_name*
:   Name of an existing column.

*extension_name*
:   The name of the extension that the materialized view is to depend on (or no longer dependent on, if `NO` is specified). A materialized view that's marked as dependent on an extension is automatically dropped when the extension is dropped.

*new_column_name*
:   New name for an existing column.

*new_owner*
:   The user name of the new owner of the materialized view.

*new_name*
:   The new name for the materialized view.

*new_schema*
:   The new schema for the materialized view.


## Examples


 To rename the materialized view `foo` to `bar`:

```sql

ALTER MATERIALIZED VIEW foo RENAME TO bar;
```


## Compatibility


 `ALTER MATERIALIZED VIEW` is a PostgreSQL extension.


## See Also
  [sql-creatematerializedview](create-materialized-view.md#sql-creatematerializedview), [sql-dropmaterializedview](drop-materialized-view.md#sql-dropmaterializedview), [sql-refreshmaterializedview](refresh-materialized-view.md#sql-refreshmaterializedview)
