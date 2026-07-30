<a id="sql-drop-property-graph"></a>

# DROP PROPERTY GRAPH

remove an SQL-property graph

## Synopsis


```

DROP PROPERTY GRAPH [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP PROPERTY GRAPH` drops an existing property graph. To execute this command you must be the owner of the property graph.


## Parameters


`IF EXISTS`
:   Do not throw an error if the property graph does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of the property graph to remove.

`CASCADE`
:   Automatically drop objects that depend on the property graph, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the property graph if any objects depend on it. This is the default.


## Examples


```sql

DROP PROPERTY GRAPH g1;
```


## Compatibility


 `DROP PROPERTY GRAPH` conforms to ISO/IEC 9075-16 (SQL/PGQ), except that the standard only allows one property graph to be dropped per command, and apart from the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also
  [sql-create-property-graph](create-property-graph.md#sql-create-property-graph), [sql-alter-property-graph](alter-property-graph.md#sql-alter-property-graph)
