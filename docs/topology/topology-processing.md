<a id="Topology_Processing"></a>

## Topology Processing
  <a id="TopologyPolygonize"></a>

# Polygonize

Finds and registers all faces defined by topology edges.

## Synopsis


```sql
text Polygonize(varchar  toponame)
```


## Description


Registers all faces that can be built out a topology edge primitives.


The target topology is assumed to contain no self-intersecting edges.


!!! note

    Already known faces are recognized, so it is safe to call Polygonize multiple times on the same topology.


!!! note

    This function does not use nor set the next_left_edge and next_right_edge fields of the edge table.


Availability: 2.0.0


## See Also


[AddFace](#AddFace), [ST_Polygonize](../postgis-reference/geometry-processing.md#ST_Polygonize)
  <a id="AddNode"></a>

# AddNode

Adds a point node to the node table in the specified topology schema and returns the nodeid of new node. If point already exists as node, the existing nodeid is returned.

## Synopsis


```sql
integer AddNode(varchar  toponame, geometry  apoint, boolean  allowEdgeSplitting=false, boolean  computeContainingFace=false)
```


## Description


 Adds a point node to the node table in the specified topology schema. The [AddEdge](#AddEdge) function automatically adds start and end points of an edge when called so not necessary to explicitly add nodes of an edge.


 If any edge crossing the node is found either an exception is raised or the edge is split, depending on the `allowEdgeSplitting` parameter value.


 If `computeContainingFace` is true a newly added node would get the correct containing face computed.


!!! note

    If the `apoint` geometry already exists as a node, the node is not added but the existing nodeid is returned.


Availability: 2.0.0


## Examples


```sql
SELECT topology.AddNode('ma_topo', ST_GeomFromText('POINT(227641.6 893816.5)', 26986) ) As nodeid;
-- result --
nodeid
--------
 4
```


## See Also


[AddEdge](#AddEdge), [CreateTopology](topology-constructors.md#CreateTopology)
  <a id="AddEdge"></a>

# AddEdge

Adds a linestring edge to the edge table and associated start and end points to the point nodes table of the specified topology schema using the specified linestring geometry and returns the edgeid of the new (or existing) edge.

## Synopsis


```sql
integer AddEdge(varchar  toponame, geometry  aline)
```


## Description


Adds an edge to the edge table and associated nodes to the nodes table of the specified `toponame` schema using the specified linestring geometry and returns the edgeid of the new or existing record. The newly added edge has "universe" face on both sides and links to itself.


!!! note

    If the `aline` geometry crosses, overlaps, contains or is contained by an existing linestring edge, then an error is thrown and the edge is not added.


!!! note

    The geometry of `aline` must have the same `srid` as defined for the topology otherwise an invalid spatial reference sys error will be thrown.


Performed by the GEOS module.


!!! warning

    [AddEdge](#AddEdge) is deprecated as of 3.5.0. Use [TopoGeo_AddLineString](topology-constructors.md#TopoGeo_AddLineString) instead.


Availability: 2.0.0


## Examples


```sql
SELECT topology.AddEdge('ma_topo', ST_GeomFromText('LINESTRING(227575.8 893917.2,227591.9 893900.4)', 26986) ) As edgeid;
-- result-
edgeid
--------
 1

SELECT topology.AddEdge('ma_topo', ST_GeomFromText('LINESTRING(227591.9 893900.4,227622.6 893844.2,227641.6 893816.5,
 227704.5 893778.5)', 26986) ) As edgeid;
-- result --
edgeid
--------
 2

 SELECT topology.AddEdge('ma_topo', ST_GeomFromText('LINESTRING(227591.2 893900, 227591.9 893900.4,
  227704.5 893778.5)', 26986) ) As edgeid;
 -- gives error --
 ERROR:  Edge intersects (not on endpoints) with existing edge 1
```


## See Also


 [TopoGeo_AddLineString](topology-constructors.md#TopoGeo_AddLineString), [CreateTopology](topology-constructors.md#CreateTopology), [Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys)
  <a id="AddFace"></a>

# AddFace

Registers a face primitive to a topology and gets its identifier.

## Synopsis


```sql
integer AddFace(varchar  toponame, geometry  apolygon, boolean  force_new=false)
```


## Description


 Registers a face primitive to a topology and gets its identifier.


 For a newly added face, the edges forming its boundaries and the ones contained in the face will be updated to have correct values in the left_face and right_face fields. Isolated nodes contained in the face will also be updated to have a correct containing_face field value.


!!! note

    This function does not use nor set the next_left_edge and next_right_edge fields of the edge table.


The target topology is assumed to be valid (containing no self-intersecting edges). An exception is raised if: The polygon boundary is not fully defined by existing edges or the polygon overlaps an existing face.


 If the `apolygon` geometry already exists as a face, then: if `force_new` is false (the default) the face id of the existing face is returned; if `force_new` is true a new id will be assigned to the newly registered face.


!!! note

    When a new registration of an existing face is performed (force_new=true), no action will be taken to resolve dangling references to the existing face in the edge, node an relation tables, nor will the MBR field of the existing face record be updated. It is up to the caller to deal with that.


!!! note

    The `apolygon` geometry must have the same `srid` as defined for the topology otherwise an invalid spatial reference sys error will be thrown.


Availability: 2.0.0


## Examples


```


-- first add the edges we use generate_series as an iterator (the below
-- will only work for polygons with < 10000 points because of our max in gs)
SELECT topology.AddEdge('ma_topo', ST_MakeLine(ST_PointN(geom,i), ST_PointN(geom, i + 1) )) As edgeid
    FROM (SELECT  ST_NPoints(geom) AS npt, geom
            FROM
                (SELECT ST_Boundary(ST_GeomFromText('POLYGON((234896.5 899456.7,234914 899436.4,234946.6 899356.9,234872.5 899328.7,
                234891 899285.4,234992.5 899145, 234890.6 899069,234755.2 899255.4,
                234612.7 899379.4,234776.9 899563.7,234896.5 899456.7))', 26986) )  As geom
            )  As geoms) As facen CROSS JOIN generate_series(1,10000) As i
         WHERE i < npt;
-- result --
 edgeid
--------
      3
      4
      5
      6
      7
      8
      9
     10
     11
     12
(10 rows)
-- then add the face -

SELECT topology.AddFace('ma_topo',
    ST_GeomFromText('POLYGON((234896.5 899456.7,234914 899436.4,234946.6 899356.9,234872.5 899328.7,
    234891 899285.4,234992.5 899145, 234890.6 899069,234755.2 899255.4,
    234612.7 899379.4,234776.9 899563.7,234896.5 899456.7))', 26986) ) As faceid;
-- result --
faceid
--------
 1
```


## See Also


[AddEdge](#AddEdge), [CreateTopology](topology-constructors.md#CreateTopology), [Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys)
  <a id="TP_ST_Simplify"></a>

# ST_Simplify

Returns a "simplified" geometry version of the given TopoGeometry using the Douglas-Peucker algorithm.

## Synopsis


```sql
geometry ST_Simplify(TopoGeometry tg, float8 tolerance)
```


## Description


Returns a "simplified" geometry version of the given TopoGeometry using the Douglas-Peucker algorithm on each component edge.


!!! note

    The returned geometry may be non-simple or non-valid.


    Splitting component edges may help retaining simplicity/validity.


Performed by the GEOS module.


Availability: 2.1.0


## See Also


Geometry [ST_Simplify](../postgis-reference/geometry-processing.md#ST_Simplify), [ST_IsSimple](../postgis-reference/geometry-accessors.md#ST_IsSimple), [ST_IsValid](../postgis-reference/geometry-validation.md#ST_IsValid), [ST_ModEdgeSplit](topology-editors.md#ST_ModEdgeSplit)
  <a id="TP_RemoveUnusedPrimitives"></a>

# RemoveUnusedPrimitives

Removes topology primitives which not needed to define existing TopoGeometry objects.

## Synopsis


```sql
int RemoveUnusedPrimitives(text topology_name, geometry bbox)
```


## Description


 Finds all primitives (nodes, edges, faces) that are not strictly needed to represent existing TopoGeometry objects and removes them, maintaining topology validity (edge linking, face labeling) and TopoGeometry space occupation.


 No new primitive identifiers are created, but rather existing primitives are expanded to include merged faces (upon removing edges) or healed edges (upon removing nodes).


Availability: 3.3.0


## See Also


 [ST_ModEdgeHeal](topology-editors.md#ST_ModEdgeHeal), [ST_RemEdgeModFace](topology-editors.md#ST_RemEdgeModFace)
