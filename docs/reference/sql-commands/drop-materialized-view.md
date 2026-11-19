<a id="sql-dropmaterializedview"></a>

# DROP MATERIALIZED VIEW

remove a materialized view

## Synopsis


```

DROP MATERIALIZED VIEW [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP MATERIALIZED VIEW` drops an existing materialized view. To execute this command you must be the owner of the materialized view.


## Parameters


`IF EXISTS`
:   Do not throw an error if the materialized view does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of the materialized view to remove.

`CASCADE`
:   Automatically drop objects that depend on the materialized view (such as other materialized views, or regular views), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the materialized view if any objects depend on it. This is the default.


## Examples


 This command will remove the materialized view called `order_summary`:

```sql

DROP MATERIALIZED VIEW order_summary;
```


## Compatibility


 `DROP MATERIALIZED VIEW` is a PostgreSQL extension.


## See Also
  [sql-creatematerializedview](create-materialized-view.md#sql-creatematerializedview), [sql-altermaterializedview](alter-materialized-view.md#sql-altermaterializedview), [sql-refreshmaterializedview](refresh-materialized-view.md#sql-refreshmaterializedview)
