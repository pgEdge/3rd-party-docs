<a id="Topology_ManagementFunctions"></a>

## Topology and TopoGeometry Management
  <a id="AddTopoGeometryColumn"></a>

# AddTopoGeometryColumn

Adds a topogeometry column to an existing table, registers this new column as a layer in topology.layer and returns the new layer_id.

## Synopsis


```sql
integer AddTopoGeometryColumn(varchar
                        topology_name, varchar
                        schema_name, varchar
                        table_name, varchar
                        column_name, varchar
                        feature_type)
integer AddTopoGeometryColumn(varchar
                        topology_name, varchar
                        schema_name, varchar
                        table_name, varchar
                        column_name, varchar
                        feature_type, integer
                        child_layer)
```


## Description


Each TopoGeometry object belongs to a specific Layer of a specific Topology. Before creating a TopoGeometry object you need to create its TopologyLayer. A Topology Layer is an association of a feature-table with the topology. It also contain type and hierarchy information. We create a layer using the AddTopoGeometryColumn() function:


This function will both add the requested column to the table and add a record to the topology.layer table with all the given info.


If you don't specify [child_layer] (or set it to NULL) this layer would contain Basic TopoGeometries (composed by primitive topology elements). Otherwise this layer will contain hierarchical TopoGeometries (composed by TopoGeometries from the child_layer).


Once the layer is created (its id is returned by the AddTopoGeometryColumn function) you're ready to construct TopoGeometry objects in it


Valid `feature_type`s are: POINT, MULTIPOINT, LINE, MULTILINE, POLYGON, MULTIPOLYGON, COLLECTION


Availability: 1.1


## Examples


```
-- Note for this example we created our new table in the ma_topo schema
-- though we could have created it in a different schema -- in which case topology_name and schema_name would be different
CREATE SCHEMA ma;
CREATE TABLE ma.parcels(gid serial, parcel_id varchar(20) PRIMARY KEY, address text);
SELECT topology.AddTopoGeometryColumn('ma_topo', 'ma', 'parcels', 'topo', 'POLYGON');
```


```sql

CREATE SCHEMA ri;
CREATE TABLE ri.roads(gid serial PRIMARY KEY, road_name text);
SELECT topology.AddTopoGeometryColumn('ri_topo', 'ri', 'roads', 'topo', 'LINE');
```


## See Also


 [DropTopoGeometryColumn](#DropTopoGeometryColumn), [toTopoGeom](topogeometry-constructors.md#toTopoGeom), [CreateTopology](topology-constructors.md#CreateTopology), [CreateTopoGeom](topogeometry-constructors.md#CreateTopoGeom)
  <a id="RenameTopoGeometryColumn"></a>

# RenameTopoGeometryColumn

Renames a topogeometry column

## Synopsis


```sql
topology.layer RenameTopoGeometryColumn(regclass
                        layer_table, name
                        feature_column, name
                        new_name)
```


## Description


 This function changes the name of an existing TopoGeometry column ensuring metadata information about it is updated accordingly.


Availability: 3.4.0


## Examples


```sql

SELECT topology.RenameTopoGeometryColumn('public.parcels', 'topogeom', 'tgeom');

```


## See Also


 [AddTopoGeometryColumn](#AddTopoGeometryColumn), [RenameTopology](#RenameTopology)
  <a id="DropTopology"></a>

# DropTopology

Use with caution: Drops a topology schema and deletes its reference from topology.topology table and references to tables in that schema from the geometry_columns table.

## Synopsis


```sql
integer DropTopology(varchar  topology_schema_name)
```


## Description


Drops a topology schema and deletes its reference from topology.topology table and references to tables in that schema from the geometry_columns table. This function should be USED WITH CAUTION, as it could destroy data you care about. If the schema does not exist, it just removes reference entries the named schema.


Availability: 1.1


## Examples


Cascade drops the ma_topo schema and removes all references to it in topology.topology and geometry_columns.


```sql
SELECT topology.DropTopology('ma_topo');
```


## See Also


[DropTopoGeometryColumn](#DropTopoGeometryColumn)
  <a id="RenameTopology"></a>

# RenameTopology

Renames a topology

## Synopsis


```sql
varchar RenameTopology(varchar  old_name, varchar  new_name)
```


## Description


 Renames a topology schema, updating its metadata record in the `topology.topology` table.


Availability: 3.4.0


## Examples


 Rename a topology from `topo_stage` to `topo_prod`.


```sql
SELECT topology.RenameTopology('topo_stage', 'topo_prod');
```


## See Also


 [CopyTopology](topology-constructors.md#CopyTopology), [RenameTopoGeometryColumn](#RenameTopoGeometryColumn)
  <a id="DropTopoGeometryColumn"></a>

# DropTopoGeometryColumn

Drops the topogeometry column from the table named `table_name` in schema `schema_name` and unregisters the columns from topology.layer table.

## Synopsis


```sql
text DropTopoGeometryColumn(varchar  schema_name, varchar  table_name, varchar  column_name)
```


## Description


Drops the topogeometry column from the table named `table_name` in schema `schema_name` and unregisters the columns from topology.layer table. Returns summary of drop status. NOTE: it first sets all values to NULL before dropping to bypass referential integrity checks.


Availability: 1.1


## Examples


```sql
SELECT topology.DropTopoGeometryColumn('ma_topo', 'parcel_topo', 'topo');
```


## See Also


[AddTopoGeometryColumn](#AddTopoGeometryColumn)
  <a id="Populate_Topology_Layer"></a>

# Populate_Topology_Layer

Adds missing entries to topology.layer table by reading metadata from topo tables.

## Synopsis


```sql
setof record Populate_Topology_Layer()
```


## Description


Adds missing entries to the `topology.layer` table by inspecting topology constraints on tables. This function is useful for fixing up entries in topology catalog after restores of schemas with topo data.


It returns the list of entries created. Returned columns are `schema_name`, `table_name`, `feature_column`.


Availability: 2.3.0


## Examples


```sql
SELECT CreateTopology('strk_topo');
CREATE SCHEMA strk;
CREATE TABLE strk.parcels(gid serial, parcel_id varchar(20) PRIMARY KEY, address text);
SELECT topology.AddTopoGeometryColumn('strk_topo', 'strk', 'parcels', 'topo', 'POLYGON');
-- this will return no records because this feature is already registered
SELECT *
  FROM topology.Populate_Topology_Layer();

-- let's rebuild
TRUNCATE TABLE topology.layer;

SELECT *
  FROM topology.Populate_Topology_Layer();

SELECT topology_id,layer_id, schema_name As sn, table_name As tn, feature_column As fc
FROM topology.layer;


```


```
 schema_name | table_name | feature_column
-------------+------------+----------------
 strk        | parcels    | topo
(1 row)

 topology_id | layer_id |  sn  |   tn    |  fc
-------------+----------+------+---------+------
           2 |        2 | strk | parcels | topo
(1 row)
```


## See Also


[AddTopoGeometryColumn](#AddTopoGeometryColumn)
  <a id="TopologySummary"></a>

# TopologySummary

Takes a topology name and provides summary totals of types of objects in topology.

## Synopsis


```sql
text TopologySummary(varchar  topology_schema_name)
```


## Description


Takes a topology name and provides summary totals of types of objects in topology.


Availability: 2.0.0


## Examples


```sql
SELECT topology.topologysummary('city_data');
                    topologysummary
--------------------------------------------------------
 Topology city_data (329), SRID 4326, precision: 0
 22 nodes, 24 edges, 10 faces, 29 topogeoms in 5 layers
 Layer 1, type Polygonal (3), 9 topogeoms
  Deploy: features.land_parcels.feature
 Layer 2, type Puntal (1), 8 topogeoms
  Deploy: features.traffic_signs.feature
 Layer 3, type Lineal (2), 8 topogeoms
  Deploy: features.city_streets.feature
 Layer 4, type Polygonal (3), 3 topogeoms
  Hierarchy level 1, child layer 1
  Deploy: features.big_parcels.feature
 Layer 5, type Puntal (1), 1 topogeoms
  Hierarchy level 1, child layer 2
  Deploy: features.big_signs.feature
```


## See Also


[Topology_Load_Tiger](../postgis-extras/tiger-geocoder.md#Topology_Load_Tiger)
  <a id="ValidateTopology"></a>

# ValidateTopology

Returns a set of validatetopology_returntype objects detailing issues with topology.

## Synopsis


```sql
setof validatetopology_returntype ValidateTopology(varchar  toponame, geometry bbox)
```


## Description


 Returns a set of [validatetopology_returntype](topology-types.md#validatetopology_returntype) objects detailing issues with topology, optionally limiting the check to the area specified by the `bbox` parameter.


List of possible errors, what they mean and what the returned ids represent are displayed below:


| Error | id1 | id2 | Meaning |
| --- | --- | --- | --- |
| coincident nodes | Identifier of first node. | Identifier of second node. | Two nodes have the same geometry. |
| edge crosses node | Identifier of the edge. | Identifier of the node. | An edge has a node in its interior. See [ST_Relate](../postgis-reference/spatial-relationships.md#ST_Relate). |
| invalid edge | Identifier of the edge. |  | An edge geometry is invalid. See [ST_IsValid](../postgis-reference/geometry-validation.md#ST_IsValid). |
| edge not simple | Identifier of the edge. |  | An edge geometry has self-intersections. See [ST_IsSimple](../postgis-reference/geometry-accessors.md#ST_IsSimple). |
| edge crosses edge | Identifier of first edge. | Identifier of second edge. | Two edges have an interior intersection. See [ST_Relate](../postgis-reference/spatial-relationships.md#ST_Relate). |
| edge start node geometry mismatch | Identifier of the edge. | Identifier of the indicated start node. | The geometry of the node indicated as the starting node for an edge does not match the first point of the edge geometry. See [ST_StartPoint](../postgis-reference/geometry-accessors.md#ST_StartPoint). |
| edge end node geometry mismatch | Identifier of the edge. | Identifier of the indicated end node. | The geometry of the node indicated as the ending node for an edge does not match the last point of the edge geometry. See [ST_EndPoint](../postgis-reference/geometry-accessors.md#ST_EndPoint). |
| face without edges | Identifier of the orphaned face. |  | No edge reports an existing face on either of its sides (left_face, right_face). |
| face has no rings | Identifier of the partially-defined face. |  | Edges reporting a face on their sides do not form a ring. |
| face has wrong mbr | Identifier of the face with wrong mbr cache. |  | Minimum bounding rectangle of a face does not match minimum bounding box of the collection of edges reporting the face on their sides. |
| hole not in advertised face | Signed identifier of an edge, identifying the ring. See [GetRingEdges](topology-accessors.md#GetRingEdges). |  | A ring of edges reporting a face on its exterior is contained in different face. |
| not-isolated node has not- containing_face | Identifier of the ill-defined node. |  | A node which is reported as being on the boundary of one or more edges is indicating a containing face. |
| isolated node has containing_face | Identifier of the ill-defined node. |  | A node which is not reported as being on the boundary of any edges is lacking the indication of a containing face. |
| isolated node has wrong containing_face | Identifier of the misrepresented node. |  | A node which is not reported as being on the boundary of any edges indicates a containing face which is not the actual face containing it. See [GetFaceContainingPoint](topology-accessors.md#GetFaceContainingPoint). |
| invalid next_right_edge | Identifier of the misrepresented edge. | Signed id of the edge which should be indicated as the next right edge. | The edge indicated as the next edge encountered walking on the right side of an edge is wrong. |
| invalid next_left_edge | Identifier of the misrepresented edge. | Signed id of the edge which should be indicated as the next left edge. | The edge indicated as the next edge encountered walking on the left side of an edge is wrong. |
| mixed face labeling in ring | Signed identifier of an edge, identifying the ring. See [GetRingEdges](topology-accessors.md#GetRingEdges). |  | Edges in a ring indicate conflicting faces on the walking side. This is also known as a "Side Location Conflict". |
| non-closed ring | Signed identifier of an edge, identifying the ring. See [GetRingEdges](topology-accessors.md#GetRingEdges). |  | A ring of edges formed by following next_left_edge/next_right_edge attributes starts and ends on different nodes. |
| face has multiple shells | Identifier of the contended face. | Signed identifier of an edge, identifying the ring. See [GetRingEdges](topology-accessors.md#GetRingEdges). | More than a one ring of edges indicate the same face on its interior. |


Availability: 1.0.0


Enhanced: 2.0.0 more efficient edge crossing detection and fixes for false positives that were existent in prior versions.


Changed: 2.2.0 values for id1 and id2 were swapped for 'edge crosses node' to be consistent with error description.


Changed: 3.2.0 added optional bbox parameter, perform face labeling and edge linking checks.


## Examples


```sql
SELECT * FROM  topology.ValidateTopology('ma_topo');
      error        | id1 | id2
-------------------+-----+-----
face without edges |   1 |

```


## See Also


[validatetopology_returntype](topology-types.md#validatetopology_returntype), [Topology_Load_Tiger](../postgis-extras/tiger-geocoder.md#Topology_Load_Tiger)
  <a id="ValidateTopologyRelation"></a>

# ValidateTopologyRelation

Returns info about invalid topology relation records

## Synopsis


```sql
setof record ValidateTopologyRelation(varchar  toponame)
```


## Description


 Returns a set records giving information about invalidities in the relation table of the topology.


Availability: 3.2.0


## See Also


[ValidateTopology](#ValidateTopology)
  <a id="FindTopology"></a>

# FindTopology

Returns a topology record by different means.

## Synopsis


```sql
topology FindTopology(TopoGeometry topogeom)
topology FindTopology(regclass layerTable, name layerColumn)
topology FindTopology(name layerSchema, name layerTable, name layerColumn)
topology FindTopology(text topoName)
topology FindTopology(int id)
```


## Description


Takes a topology identifier or the identifier of a topology-related object and returns a topology.topology record.


Availability: 3.2.0


## Examples


```sql

SELECT name(findTopology('features.land_parcels', 'feature'));
   name
-----------
 city_data
(1 row)
```


## See Also


[FindLayer](#FindLayer)
  <a id="FindLayer"></a>

# FindLayer

Returns a topology.layer record by different means.

## Synopsis


```sql
topology.layer FindLayer(TopoGeometry tg)
topology.layer FindLayer(regclass layer_table, name feature_column)
topology.layer FindLayer(name schema_name, name table_name, name feature_column)
topology.layer FindLayer(integer topology_id, integer layer_id)
```


## Description


Takes a layer identifier or the identifier of a topology-related object and returns a topology.layer record.


Availability: 3.2.0


## Examples


```sql

SELECT layer_id(findLayer('features.land_parcels', 'feature'));
 layer_id
----------
        1
(1 row)
```


## See Also


[FindTopology](#FindTopology)
