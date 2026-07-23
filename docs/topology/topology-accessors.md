<a id="Topology_Accessors"></a>

## Topology Accessors
  <a id="GetEdgeByPoint"></a>

# GetEdgeByPoint

Finds the edge-id of an edge that intersects a given point.

## Synopsis


```sql
integer GetEdgeByPoint(varchar  atopology, geometry  apoint, float8  tol1)
```


## Description


Retrieves the id of an edge that intersects a Point.


The function returns an integer (id-edge) given a topology, a POINT and a tolerance. If tolerance = 0 then the point has to intersect the edge.


If `apoint` doesn't intersect an edge, returns 0 (zero).


If use tolerance > 0 and there is more than one edge near the point then an exception is thrown.


!!! note

    If tolerance = 0, the function uses ST_Intersects otherwise uses ST_DWithin.


Performed by the GEOS module.


Availability: 2.0.0


## Examples


These examples use edges we created in [AddEdge](topology-processing.md#AddEdge)


```sql
SELECT topology.GetEdgeByPoint('ma_topo',geom, 1) As with1mtol, topology.GetEdgeByPoint('ma_topo',geom,0) As withnotol
FROM ST_GeomFromEWKT('SRID=26986;POINT(227622.6 893843)') As geom;
 with1mtol | withnotol
-----------+-----------
         2 |         0
```


```sql
SELECT topology.GetEdgeByPoint('ma_topo',geom, 1) As nearnode
FROM ST_GeomFromEWKT('SRID=26986;POINT(227591.9 893900.4)') As geom;

-- get error --
ERROR:  Two or more edges found
```


## See Also


 [AddEdge](topology-processing.md#AddEdge), [GetNodeByPoint](#GetNodeByPoint), [GetFaceByPoint](#GetFaceByPoint)
  <a id="GetFaceByPoint"></a>

# GetFaceByPoint

Finds face intersecting a given point.

## Synopsis


```sql
integer GetFaceByPoint(varchar  atopology, geometry  apoint, float8  tol1)
```


## Description


 Finds a face referenced by a Point, with given tolerance.


 The function will effectively look for a face intersecting a circle having the point as center and the tolerance as radius.


 If no face intersects the given query location, 0 is returned (universal face).


 If more than one face intersect the query location an exception is thrown.


Availability: 2.0.0


Enhanced: 3.2.0 more efficient implementation and clearer contract, stops working with invalid topologies.


## Examples


```sql
SELECT topology.GetFaceByPoint('ma_topo',geom, 10) As with1mtol, topology.GetFaceByPoint('ma_topo',geom,0) As withnotol
	FROM ST_GeomFromEWKT('POINT(234604.6 899382.0)') As geom;

	 with1mtol | withnotol
	-----------+-----------
			 1 |         0
```


```sql
SELECT topology.GetFaceByPoint('ma_topo',geom, 1) As nearnode
	FROM ST_GeomFromEWKT('POINT(227591.9 893900.4)') As geom;

-- get error --
ERROR:  Two or more faces found
```


## See Also


 [GetFaceContainingPoint](#GetFaceContainingPoint), [AddFace](topology-processing.md#AddFace), [GetNodeByPoint](#GetNodeByPoint), [GetEdgeByPoint](#GetEdgeByPoint)
  <a id="GetFaceContainingPoint"></a>

# GetFaceContainingPoint

Finds the face containing a point.

## Synopsis


```sql
integer GetFaceContainingPoint(text  atopology, geometry  apoint)
```


## Description


Returns the id of the face containing a point.


An exception is thrown if the point falls on a face boundary.


!!! note

    The function relies on a valid topology, using edge linking and face labeling.


Availability: 3.2.0


## See Also


 [ST_GetFaceGeometry](#ST_GetFaceGeometry)
  <a id="GetNodeByPoint"></a>

# GetNodeByPoint

Finds the node-id of a node at a point location.

## Synopsis


```sql
integer GetNodeByPoint(varchar  atopology, geometry  apoint, float8  tol1)
```


## Description


Retrieves the id of a node at a point location.


The function returns an integer (id-node) given a topology, a POINT and a tolerance. If tolerance = 0 means exact intersection, otherwise retrieves the node from an interval.


If `apoint` doesn't intersect a node, returns 0 (zero).


If use tolerance > 0 and there is more than one node near the point then an exception is thrown.


!!! note

    If tolerance = 0, the function uses ST_Intersects otherwise uses ST_DWithin.


Performed by the GEOS module.


Availability: 2.0.0


## Examples


These examples use edges we created in [AddEdge](topology-processing.md#AddEdge)


```sql
SELECT topology.GetNodeByPoint('ma_topo',geom, 1) As nearnode
 FROM ST_GeomFromEWKT('SRID=26986;POINT(227591.9 893900.4)') As geom;
  nearnode
----------
        2

```


```sql
SELECT topology.GetNodeByPoint('ma_topo',geom, 1000) As too_much_tolerance
 FROM ST_GeomFromEWKT('SRID=26986;POINT(227591.9 893900.4)') As geom;

 ----get error--
 ERROR:  Two or more nodes found

```


## See Also


 [AddEdge](topology-processing.md#AddEdge), [GetEdgeByPoint](#GetEdgeByPoint), [GetFaceByPoint](#GetFaceByPoint)
  <a id="GetTopologyID"></a>

# GetTopologyID

Returns the id of a topology in the topology.topology table given the name of the topology.

## Synopsis


```sql
integer GetTopologyID(varchar toponame)
```


## Description


Returns the id of a topology in the topology.topology table given the name of the topology.


Availability: 1.1


## Examples


```sql
SELECT topology.GetTopologyID('ma_topo') As topo_id;
 topo_id
---------
       1
```


## See Also


 [CreateTopology](topology-constructors.md#CreateTopology), [DropTopology](topology-and-topogeometry-management.md#DropTopology), [GetTopologyName](#GetTopologyName), [GetTopologySRID](#GetTopologySRID)
  <a id="GetTopologySRID"></a>

# GetTopologySRID

Returns the SRID of a topology in the topology.topology table given the name of the topology.

## Synopsis


```sql
integer GetTopologyID(varchar toponame)
```


## Description


Returns the spatial reference id of a topology in the topology.topology table given the name of the topology.


Availability: 2.0.0


## Examples


```sql
SELECT topology.GetTopologySRID('ma_topo') As SRID;
 SRID
-------
  4326
```


## See Also


 [CreateTopology](topology-constructors.md#CreateTopology), [DropTopology](topology-and-topogeometry-management.md#DropTopology), [GetTopologyName](#GetTopologyName), [GetTopologyID](#GetTopologyID)
  <a id="GetTopologyName"></a>

# GetTopologyName

Returns the name of a topology (schema) given the id of the topology.

## Synopsis


```sql
varchar GetTopologyName(integer topology_id)
```


## Description


Returns the topology name (schema) of a topology from the topology.topology table given the topology id of the topology.


Availability: 1.1


## Examples


```sql
SELECT topology.GetTopologyName(1) As topo_name;
 topo_name
-----------
 ma_topo
```


## See Also


 [CreateTopology](topology-constructors.md#CreateTopology), [DropTopology](topology-and-topogeometry-management.md#DropTopology), [GetTopologyID](#GetTopologyID), [GetTopologySRID](#GetTopologySRID)
  <a id="ST_GetFaceEdges"></a>

# ST_GetFaceEdges

Returns a set of ordered edges that bound `aface`.

## Synopsis


```sql
getfaceedges_returntype ST_GetFaceEdges(varchar  atopology, integer  aface)
```


## Description


Returns a set of ordered edges that bound `aface`. Each output consists of a sequence and edgeid. Sequence numbers start with value 1.


 Enumeration of each ring edges start from the edge with smallest identifier. Order of edges follows a left-hand-rule (bound face is on the left of each directed edge).


Availability: 2.0


 SQL-MM 3 Topo-Geo and Topo-Net 3: Routine Details: X.3.5


## Examples


```

-- Returns the edges bounding face 1
SELECT (topology.ST_GetFaceEdges('tt', 1)).*;
-- result --
 sequence | edge
----------+------
        1 |   -4
        2 |    5
        3 |    7
        4 |   -6
        5 |    1
        6 |    2
        7 |    3
(7 rows)
```


```

-- Returns the sequence, edge id
-- and geometry of the edges that bound face 1
-- If you just need geom and seq, can use ST_GetFaceGeometry
SELECT t.seq, t.edge, geom
FROM topology.ST_GetFaceEdges('tt',1) As t(seq,edge)
	INNER JOIN tt.edge AS e ON abs(t.edge) = e.edge_id;
```


## See Also


 [GetRingEdges](#GetRingEdges), [AddFace](topology-processing.md#AddFace), [ST_GetFaceGeometry](#ST_GetFaceGeometry)
  <a id="ST_GetFaceGeometry"></a>

# ST_GetFaceGeometry

Returns the polygon in the given topology with the specified face id.

## Synopsis


```sql
geometry ST_GetFaceGeometry(varchar  atopology, integer  aface)
```


## Description


Returns the polygon in the given topology with the specified face id. Builds the polygon from the edges making up the face.


Availability: 1.1


 SQL-MM 3 Topo-Geo and Topo-Net 3: Routine Details: X.3.16


## Examples


```

-- Returns the wkt of the polygon added with AddFace
SELECT ST_AsText(topology.ST_GetFaceGeometry('ma_topo', 1)) As facegeomwkt;
-- result --
               facegeomwkt

--------------------------------------------------------------------------------
 POLYGON((234776.9 899563.7,234896.5 899456.7,234914 899436.4,234946.6 899356.9,
234872.5 899328.7,234891 899285.4,234992.5 899145,234890.6 899069,
234755.2 899255.4,234612.7 899379.4,234776.9 899563.7))
```


## See Also


[AddFace](topology-processing.md#AddFace)
  <a id="GetRingEdges"></a>

# GetRingEdges

Returns the ordered set of signed edge identifiers met by walking on an a given edge side.

## Synopsis


```sql
getfaceedges_returntype GetRingEdges(varchar  atopology, integer  aring, integer  max_edges=null)
```


## Description


 Returns the ordered set of signed edge identifiers met by walking on an a given edge side. Each output consists of a sequence and a signed edge id. Sequence numbers start with value 1.


 If you pass a positive edge id, the walk starts on the left side of the corresponding edge and follows the edge direction. If you pass a negative edge id, the walk starts on the right side of it and goes backward.


 If `max_edges` is not null no more than those records are returned by that function. This is meant to be a safety parameter when dealing with possibly invalid topologies.


!!! note

    This function uses edge ring linking metadata.


Availability: 2.0.0


## See Also


 [ST_GetFaceEdges](#ST_GetFaceEdges), [GetNodeEdges](#GetNodeEdges)
  <a id="GetNodeEdges"></a>

# GetNodeEdges

Returns an ordered set of edges incident to the given node.

## Synopsis


```sql
getfaceedges_returntype GetNodeEdges(varchar  atopology, integer  anode)
```


## Description


 Returns an ordered set of edges incident to the given node. Each output consists of a sequence and a signed edge id. Sequence numbers start with value 1. A positive edge starts at the given node. A negative edge ends into the given node. Closed edges will appear twice (with both signs). Order is clockwise starting from northbound.


!!! note

    This function computes ordering rather than deriving from metadata and is thus usable to build edge ring linking.


Availability: 2.0


## See Also


 [getfaceedges_returntype](topology-types.md#getfaceedges_returntype), [GetRingEdges](#GetRingEdges), [ST_Azimuth](../postgis-reference/measurement-functions.md#ST_Azimuth)
