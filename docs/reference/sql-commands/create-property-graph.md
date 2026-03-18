<a id="sql-create-property-graph"></a>

# CREATE PROPERTY GRAPH

define an SQL-property graph

## Synopsis


```

CREATE [ TEMP | TEMPORARY ] PROPERTY GRAPH NAME
    [ {VERTEX|NODE} TABLES ( VERTEX_TABLE_DEFINITION [, ...] ) ]
    [ {EDGE|RELATIONSHIP} TABLES ( EDGE_TABLE_DEFINITION [, ...] ) ]

where VERTEX_TABLE_DEFINITION is:

    VERTEX_TABLE_NAME [ AS ALIAS ] [ KEY ( COLUMN_NAME [, ...] ) ] [ ELEMENT_TABLE_LABEL_AND_PROPERTIES ]

and EDGE_TABLE_DEFINITION is:

    EDGE_TABLE_NAME [ AS ALIAS ] [ KEY ( COLUMN_NAME [, ...] ) ]
        SOURCE [ KEY ( COLUMN_NAME [, ...] ) REFERENCES ] SOURCE_TABLE [ ( COLUMN_NAME [, ...] ) ]
        DESTINATION [ KEY ( COLUMN_NAME [, ...] ) REFERENCES ] DEST_TABLE [ ( COLUMN_NAME [, ...] ) ]
        [ ELEMENT_TABLE_LABEL_AND_PROPERTIES ]

and ELEMENT_TABLE_LABEL_AND_PROPERTIES is either:

    NO PROPERTIES | PROPERTIES ALL COLUMNS | PROPERTIES ( { EXPRESSION [ AS PROPERTY_NAME ] } [, ...] )

or:

   { { LABEL LABEL_NAME | DEFAULT LABEL } [ NO PROPERTIES | PROPERTIES ALL COLUMNS | PROPERTIES ( { EXPRESSION [ AS PROPERTY_NAME ] } [, ...] ) ] } [...]
```


## Description


 `CREATE PROPERTY GRAPH` defines a property graph. A property graph consists of vertices and edges, together called elements, each with associated labels and properties, and can be queried using the `GRAPH_TABLE` clause of [sql-select](select.md#sql-select) with a special path matching syntax. The data in the graph is stored in regular tables (or views, foreign tables, etc.). Each vertex or edge corresponds to a table. The property graph definition links these tables together into a graph structure that can be queried using graph query techniques.


 `CREATE PROPERTY GRAPH` does not physically materialize a graph. It is thus similar to `CREATE VIEW` in that it records a structure that is used only when the defined object is queried.


 If a schema name is given (for example, `CREATE PROPERTY GRAPH myschema.mygraph ...`) then the property graph is created in the specified schema. Otherwise it is created in the current schema. Temporary property graphs exist in a special schema, so a schema name cannot be given when creating a temporary property graph. Property graphs share a namespace with tables and other relation types, so the name of the property graph must be distinct from the name of any other relation (table, sequence, index, view, materialized view, or foreign table) in the same schema.


## Parameters


*name*
:   The name (optionally schema-qualified) of the new property graph.

`VERTEX`/`NODE`, `EDGE`/`RELATIONSHIP`
:   These keywords are synonyms, respectively.

*vertex_table_name*
:   The name of a table that will contain vertices in the new property graph.

*edge_table_name*
:   The name of a table that will contain edges in the new property graph.

*alias*
:   A unique identifier for the vertex or edge table. This defaults to the name of the table. Aliases must be unique in a property graph definition (across all vertex table and edge table definitions). (Therefore, if a table is used more than once as a vertex or edge table, then an explicit alias must be specified for at least one of them to distinguish them.)

<code>KEY ( </code><em>column_name</em><code> [, ...] )</code>
:   A set of columns that uniquely identifies a row in the vertex or edge table. Defaults to the primary key.

*source_table*, *dest_table*
:   The vertex tables that the edge table is linked to. These refer to the aliases of the source and destination vertex tables respectively.

<code>KEY ( </code><em>column_name</em><code> [, ...] ) REFERENCES ... ( </code><em>column_name</em><code> [, ...] )</code>
:   Two sets of columns that connect the edge table and the source or destination vertex table, like in a foreign-key relationship. If a foreign-key constraint between the two tables exists, it is used by default.

*element_table_label_and_properties*
:   Defines the labels and properties for the element (vertex or edge) table. Each element has at least one label. By default, the label is the same as the element table alias. This can be specified explicitly as `DEFAULT LABEL`. Alternatively, one or more freely chosen label names can be specified. (Label names do not have to be unique across a property graph. It can be useful to assign the same label to different elements.) Each label has a list (possibly empty) of properties. By default, all columns of a table are automatically exposed as properties. This can be specified explicitly as `PROPERTIES ALL COLUMNS`. Alternatively, a list of expressions, which can refer to the columns of the underlying table, can be specified as properties. If the expressions are not a plain column reference, then an explicit property name must also be specified.
 <a id="sql-create-property-graph-notes"></a>

## Notes


 The following consistency checks must be satisfied by a property graph definition:

-  In a property graph, labels with the same name applied to different property graph elements must have the same number of properties and those properties must have the same names. For example, the following would be allowed:

```sql

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (
        v1 LABEL foo PROPERTIES (x, y),
        v2 LABEL foo PROPERTIES (x, y)
    ) ...
```
   but this would not:

```sql

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (
        v1 LABEL foo PROPERTIES (x, y),
        v2 LABEL foo PROPERTIES (z)
    ) ...
```

-  In a property graph, all properties with the same name must have the same data type, independent of which label they are on. For example, this would be allowed:

```sql

CREATE TABLE v1 (a int, b int);
CREATE TABLE v2 (a int, b int);

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (
        v1 LABEL foo PROPERTIES (a, b),
        v2 LABEL bar PROPERTIES (a, b)
    ) ...
```
   but this would not:

```sql

CREATE TABLE v1 (a int, b int);
CREATE TABLE v2 (a int, b varchar);

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (
        v1 LABEL foo PROPERTIES (a, b),
        v2 LABEL bar PROPERTIES (a, b)
    ) ...
```

-  For each property graph element, all properties with the same name must have the same expression for each label. For example, this would be allowed:

```sql

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (
        v1 LABEL foo PROPERTIES (a * 2 AS x) LABEL bar PROPERTIES (a * 2 AS x)
    ) ...
```
   but this would not:

```sql

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (
        v1 LABEL foo PROPERTIES (a * 2 AS x) LABEL bar PROPERTIES (a * 10 AS x)
    ) ...
```


 Property graphs are queried using the `GRAPH_TABLE` clause of [sql-select](select.md#sql-select).


 Access to the base relations underlying the `GRAPH_TABLE` clause is determined by the permissions of the user executing the query, rather than the property graph owner. Thus, the user of a property graph must have the relevant permissions on the property graph and base relations underlying the `GRAPH_TABLE` clause.


## Examples


```sql

CREATE PROPERTY GRAPH g1
    VERTEX TABLES (v1, v2, v3)
    EDGE TABLES (e1 SOURCE v1 DESTINATION v2,
                 e2 SOURCE v1 DESTINATION v3);
```


## Compatibility


 `CREATE PROPERTY GRAPH` conforms to ISO/IEC 9075-16 (SQL/PGQ).


## See Also
  [sql-alter-property-graph](alter-property-graph.md#sql-alter-property-graph), [sql-drop-property-graph](drop-property-graph.md#sql-drop-property-graph)
