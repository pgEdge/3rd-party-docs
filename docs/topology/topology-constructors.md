<a id="Topology_Constructors"></a>

## Topology Constructors
  <a id="CreateTopology"></a>

# CreateTopology

Creates a new topology schema and registers it in the topology.topology table.

## Synopsis


```sql
integer CreateTopology(varchar  topology_schema_name)
integer CreateTopology(varchar  topology_schema_name, integer  srid)
integer CreateTopology(varchar  topology_schema_name, integer  srid, double precision  prec)
integer CreateTopology(varchar  topology_schema_name, integer  srid, double precision  prec, boolean  hasz)
```


## Description


Creates a new topology schema with name `topology_name` and registers it in the `topology.topology` table. Topologies must be uniquely named. The topology tables (`edge_data`, `face`, `node`,and `relation` are created in the schema. It returns the id of the topology.


The `srid` is the [spatial reference system](../data-management/spatial-reference-systems.md#spatial_ref_sys) SRID for the topology.


The tolerance `prec` is measured in the units of the spatial reference system. The tolerance defaults to 0.


`hasz` defaults to false if not specified.


This is similar to the SQL/MM [ST_InitTopoGeo](#ST_InitTopoGeo) but has more functionality.


Availability: 1.1


Enhanced: 2.0 added the signature accepting hasZ


## Examples


Create a topology schema called `ma_topo` that stores edges and nodes in Massachusetts State Plane-meters (SRID = 26986). The tolerance represents 0.5 meters since the spatial reference system is meter-based.


```sql
SELECT topology.CreateTopology('ma_topo', 26986, 0.5);
```


Create a topology for Rhode Island called `ri_topo` in spatial reference system State Plane-feet (SRID = 3438)


```sql
SELECT topology.CreateTopology('ri_topo', 3438) AS topoid;
topoid
------
2
```


## See Also


[Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys), [ST_InitTopoGeo](#ST_InitTopoGeo), [Topology_Load_Tiger](../postgis-extras/tiger-geocoder.md#Topology_Load_Tiger)
  <a id="CopyTopology"></a>

# CopyTopology

Makes a copy of a topology (nodes, edges, faces, layers and TopoGeometries) into a new schema

## Synopsis


```sql
integer CopyTopology(varchar  existing_topology_name, varchar  new_name)
```


## Description


 Creates a new topology with name `new_name`, with SRID and precision copied from `existing_topology_name` The nodes, edges and faces in `existing_topology_name` are copied into the new topology, as well as Layers and their associated TopoGeometries.


!!! note

    The new rows in the `topology.layer` table contain synthetic values for schema_name, table_name and feature_column. This is because the TopoGeometry objects exist only as a definition and are not yet available in a user-defined table.


Availability: 2.0.0


## Examples


 Make a backup of a topology called `ma_topo`.


```sql
SELECT topology.CopyTopology('ma_topo', 'ma_topo_backup');
```


## See Also


 [Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys), [CreateTopology](#CreateTopology), [RenameTopology](topology-and-topogeometry-management.md#RenameTopology)
  <a id="ST_InitTopoGeo"></a>

# ST_InitTopoGeo

Creates a new topology schema and registers it in the topology.topology table.

## Synopsis


```sql
text ST_InitTopoGeo(varchar  topology_schema_name)
```


## Description


This is the SQL-MM equivalent of [CreateTopology](#CreateTopology). It lacks options for spatial reference system and tolerance. it returns a text description of the topology creation, instead of the topology id.


Availability: 1.1


 SQL-MM 3 Topo-Geo and Topo-Net 3: Routine Details: X.3.17


## Examples


```sql
SELECT topology.ST_InitTopoGeo('topo_schema_to_create') AS topocreation;
                      astopocreation
------------------------------------------------------------
 Topology-Geometry 'topo_schema_to_create' (id:7) created.

```


## See Also


[CreateTopology](#CreateTopology)
  <a id="ST_CreateTopoGeo"></a>

# ST_CreateTopoGeo

Adds a collection of geometries to a given empty topology and returns a message detailing success.

## Synopsis


```sql
text ST_CreateTopoGeo(varchar  atopology, geometry  acollection)
```


## Description


 Adds a collection of geometries to a given empty topology and returns a message detailing success.


Useful for populating an empty topology.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details -- X.3.18


## Examples


```

-- Populate topology --
SELECT topology.ST_CreateTopoGeo('ri_topo',
 ST_GeomFromText('MULTILINESTRING((384744 236928,384750 236923,384769 236911,384799 236895,384811 236890,384833 236884,
  384844 236882,384866 236881,384879 236883,384954 236898,385087 236932,385117 236938,
  385167 236938,385203 236941,385224 236946,385233 236950,385241 236956,385254 236971,
  385260 236979,385268 236999,385273 237018,385273 237037,385271 237047,385267 237057,
  385225 237125,385210 237144,385192 237161,385167 237192,385162 237202,385159 237214,
  385159 237227,385162 237241,385166 237256,385196 237324,385209 237345,385234 237375,
  385237 237383,385238 237399,385236 237407,385227 237419,385213 237430,385193 237439,
  385174 237451,385170 237455,385169 237460,385171 237475,385181 237503,385190 237521,
  385200 237533,385206 237538,385213 237541,385221 237542,385235 237540,385242 237541,
  385249 237544,385260 237555,385270 237570,385289 237584,385292 237589,385291 237596,385284 237630))',3438)
  );

      st_createtopogeo
----------------------------
 Topology ri_topo populated


-- create tables and topo geometries --
CREATE TABLE ri.roads(gid serial PRIMARY KEY, road_name text);

SELECT topology.AddTopoGeometryColumn('ri_topo', 'ri', 'roads', 'topo', 'LINE');

```


## See Also


 [TopoGeo_LoadGeometry](#TopoGeo_LoadGeometry), [AddTopoGeometryColumn](topology-and-topogeometry-management.md#AddTopoGeometryColumn), [CreateTopology](#CreateTopology), [DropTopology](topology-and-topogeometry-management.md#DropTopology)
  <a id="TopoGeo_AddPoint"></a>

# TopoGeo_AddPoint

Adds a point to an existing topology using a tolerance and possibly splitting an existing edge.

## Synopsis


```sql
integer TopoGeo_AddPoint(varchar  atopology, geometry  apoint, float8  tolerance)
```


## Description


 Adds a point to an existing topology and returns its identifier. The given point will snap to existing nodes or edges within given tolerance. An existing edge may be split by the snapped point.


Availability: 2.0.0


## See Also


 [TopoGeo_AddLineString](#TopoGeo_AddLineString), [TopoGeo_AddPolygon](#TopoGeo_AddPolygon), [TopoGeo_LoadGeometry](#TopoGeo_LoadGeometry), [AddNode](topology-processing.md#AddNode), [CreateTopology](#CreateTopology)
  <a id="TopoGeo_AddLineString"></a>

# TopoGeo_AddLineString

Adds a linestring to an existing topology using a tolerance and possibly splitting existing edges/faces.

## Synopsis


```sql
SETOF integer TopoGeo_AddLineString(varchar  atopology, geometry  aline, float8  tolerance)
```


## Description


 Adds a linestring to an existing topology and returns a set of signed edge identifiers forming it up (negative identifies mean the edge goes in the opposite direction of the input linestring). The given line will snap to existing nodes or edges within given tolerance. Existing edges and faces may be split by the line. New nodes and faces may be added.


!!! note

    Updating statistics about topologies being loaded via this function is up to caller, see [maintaining statistics during topology editing and population](topology-statistics-management.md#Topology_StatsManagement).


Availability: 2.0.0


Enhanced: 3.2.0 added support for returning signed identifier.


## See Also


 [TopoGeo_AddPoint](#TopoGeo_AddPoint), [TopoGeo_AddPolygon](#TopoGeo_AddPolygon), [TopoGeo_LoadGeometry](#TopoGeo_LoadGeometry), [AddEdge](topology-processing.md#AddEdge), [CreateTopology](#CreateTopology)
  <a id="TopoGeo_AddPolygon"></a>

# TopoGeo_AddPolygon

Adds a polygon to an existing topology using a tolerance and possibly splitting existing edges/faces. Returns face identifiers.

## Synopsis


```sql
SETOF integer TopoGeo_AddPolygon(varchar  atopology, geometry  apoly, float8  tolerance)
```


## Description


 Adds a polygon to an existing topology and returns a set of face identifiers forming it up. The boundary of the given polygon will snap to existing nodes or edges within given tolerance. Existing edges and faces may be split by the boundary of the new polygon.


!!! note

    Updating statistics about topologies being loaded via this function is up to caller, see [maintaining statistics during topology editing and population](topology-statistics-management.md#Topology_StatsManagement).


Availability: 2.0.0


## See Also


 [TopoGeo_AddPoint](#TopoGeo_AddPoint), [TopoGeo_AddLineString](#TopoGeo_AddLineString), [TopoGeo_LoadGeometry](#TopoGeo_LoadGeometry), [AddFace](topology-processing.md#AddFace), [CreateTopology](#CreateTopology)
  <a id="TopoGeo_LoadGeometry"></a>

# TopoGeo_LoadGeometry

Load a geometry into an existing topology, snapping and splitting as needed.

## Synopsis


```sql
void TopoGeo_LoadGeometry(varchar  atopology, geometry  ageom, float8  tolerance)
```


## Description


 Loads a geometry into an existing topology. The given geometry will snap to existing nodes or edges within given tolerance. Existing edges and faces may be split as a consequence of the load.


!!! note

    Updating statistics about topologies being loaded via this function is up to caller, see [maintaining statistics during topology editing and population](topology-statistics-management.md#Topology_StatsManagement).


Availability: 3.5.0


## See Also


 [TopoGeo_AddPoint](#TopoGeo_AddPoint), [TopoGeo_AddLineString](#TopoGeo_AddLineString), [TopoGeo_AddPolygon](#TopoGeo_AddPolygon), [CreateTopology](#CreateTopology)
