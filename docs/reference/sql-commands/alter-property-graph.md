<a id="sql-alter-property-graph"></a>

# ALTER PROPERTY GRAPH

change the definition of an SQL-property graph

## Synopsis


```

ALTER PROPERTY GRAPH NAME ADD
    [ {VERTEX|NODE} TABLES ( VERTEX_TABLE_DEFINITION [, ...] ) ]
    [ {EDGE|RELATIONSHIP} TABLES ( EDGE_TABLE_DEFINITION [, ...] ) ]

ALTER PROPERTY GRAPH NAME DROP
    {VERTEX|NODE} TABLES ( VERTEX_TABLE_ALIAS [, ...] ) [ CASCADE | RESTRICT ]

ALTER PROPERTY GRAPH NAME DROP
    {EDGE|RELATIONSHIP} TABLES ( EDGE_TABLE_ALIAS [, ...] ) [ CASCADE | RESTRICT ]

ALTER PROPERTY GRAPH NAME ALTER
    {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ELEMENT_TABLE_ALIAS
    { ADD LABEL LABEL_NAME [ NO PROPERTIES | PROPERTIES ALL COLUMNS | PROPERTIES ( { EXPRESSION [ AS PROPERTY_NAME ] } [, ...] ) ] } [ ... ]

ALTER PROPERTY GRAPH NAME ALTER
    {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ELEMENT_TABLE_ALIAS
    DROP LABEL LABEL_NAME [ CASCADE | RESTRICT ]

ALTER PROPERTY GRAPH NAME ALTER
    {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ELEMENT_TABLE_ALIAS
    ALTER LABEL LABEL_NAME ADD PROPERTIES ( { EXPRESSION [ AS PROPERTY_NAME ] } [, ...] )

ALTER PROPERTY GRAPH NAME ALTER
    {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ELEMENT_TABLE_ALIAS
    ALTER LABEL LABEL_NAME DROP PROPERTIES ( PROPERTY_NAME [, ...] ) [ CASCADE | RESTRICT ]

ALTER PROPERTY GRAPH NAME OWNER TO { NEW_OWNER | CURRENT_USER | SESSION_USER }
ALTER PROPERTY GRAPH NAME RENAME TO NEW_NAME
ALTER PROPERTY GRAPH [ IF EXISTS ] NAME SET SCHEMA NEW_SCHEMA
```


## Description


 `ALTER PROPERTY GRAPH` changes the definition of an existing property graph. There are several subforms:

`ADD {VERTEX|NODE|EDGE|RELATIONSHIP} TABLES`
:   This form adds new vertex or edge tables to the property graph, using the same syntax as [`CREATE PROPERTY GRAPH`](create-property-graph.md#sql-create-property-graph).

`DROP {VERTEX|NODE|EDGE|RELATIONSHIP} TABLES`
:   This form removes vertex or edge tables from the property graph. (Only the association of the tables with the graph is removed. The tables themselves are not dropped.)

`ALTER {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ... ADD LABEL`
:   This form adds a new label to an existing vertex or edge table, using the same syntax as [`CREATE PROPERTY GRAPH`](create-property-graph.md#sql-create-property-graph).

`ALTER {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ... DROP LABEL`
:   This form removes a label from an existing vertex or edge table.

`ALTER {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ... ALTER LABEL ... ADD PROPERTIES`
:   This form adds new properties to an existing label on an existing vertex or edge table.

`ALTER {VERTEX|NODE|EDGE|RELATIONSHIP} TABLE ... ALTER LABEL ... DROP PROPERTIES`
:   This form removes properties from an existing label on an existing vertex or edge table.

`OWNER`
:   This form changes the owner of the property graph to the specified user.

`RENAME`
:   This form changes the name of a property graph.

`SET SCHEMA`
:   This form moves the property graph into another schema.


 You must own the property graph to use `ALTER PROPERTY GRAPH`. To change a property graph's schema, you must also have `CREATE` privilege on the new schema. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the property graph's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the property graph. However, a superuser can alter ownership of any property graph anyway.)


## Parameters


*name*
:   The name (optionally schema-qualified) of a property graph to be altered.

`IF EXISTS`
:   Do not throw an error if the property graph does not exist. A notice is issued in this case.

*vertex_table_definition*, *edge_table_definition*
:   See [`CREATE PROPERTY GRAPH`](create-property-graph.md#sql-create-property-graph).

*vertex_table_alias*, *edge_table_alias*
:   The alias of an existing vertex or edge table to operate on. (Note that the alias is potentially different from the name of the underlying table, if the vertex or edge table was created with <code>AS
          </code><em>alias</em>.)

*label_name*, *property_name*, *expression*
:   See [`CREATE PROPERTY GRAPH`](create-property-graph.md#sql-create-property-graph).

*new_owner*
:   The user name of the new owner of the property graph.

*new_name*
:   The new name for the property graph.

*new_schema*
:   The new schema for the property graph.


## Notes


 The consistency checks on a property graph described at [Notes](create-property-graph.md#sql-create-property-graph-notes) must be maintained by `ALTER PROPERTY GRAPH` operations. In some cases, it might be necessary to make multiple alterations in a single command to satisfy the checks.


## Examples


```sql

ALTER PROPERTY GRAPH g1 ADD VERTEX TABLES (v2);

ALTER PROPERTY GRAPH g1 ALTER VERTEX TABLE v1 DROP LABEL foo;

ALTER PROPERTY GRAPH g1 RENAME TO g2;
```


## Compatibility


 `ALTER PROPERTY GRAPH` conforms to ISO/IEC 9075-16 (SQL/PGQ).


## See Also
  [sql-create-property-graph](create-property-graph.md#sql-create-property-graph), [sql-drop-property-graph](drop-property-graph.md#sql-drop-property-graph)
