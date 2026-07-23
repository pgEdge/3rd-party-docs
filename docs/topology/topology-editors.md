<a id="Topology_Editing"></a>

## Topology Editors
  <a id="ST_AddIsoNode"></a>

# ST_AddIsoNode

Adds an isolated node to a face in a topology and returns the nodeid of the new node. If face is null, the node is still created.

## Synopsis


```sql
integer ST_AddIsoNode(varchar  atopology, integer  aface, geometry  apoint)
```


## Description


Adds an isolated node with point location `apoint` to an existing face with faceid `aface` to a topology `atopology` and returns the nodeid of the new node.


If the spatial reference system (srid) of the point geometry is not the same as the topology, the `apoint` is not a point geometry, the point is null, or the point intersects an existing edge (even at the boundaries) then an exception is thrown. If the point already exists as a node, an exception is thrown.


If `aface` is not null and the `apoint` is not within the face, then an exception is thrown.


Availability: 1.1


 SQL-MM: Topo-Net Routines: X+1.3.1


## Examples


## See Also


[AddNode](topology-processing.md#AddNode), [CreateTopology](topology-constructors.md#CreateTopology), [DropTopology](topology-and-topogeometry-management.md#DropTopology), [ST_Intersects](../postgis-reference/spatial-relationships.md#ST_Intersects)
  <a id="ST_AddIsoEdge"></a>

# ST_AddIsoEdge

Adds an isolated edge defined by geometry `alinestring` to a topology connecting two existing isolated nodes `anode` and `anothernode` and returns the edge id of the new edge.

## Synopsis


```sql
integer ST_AddIsoEdge(varchar  atopology, integer  anode, integer  anothernode, geometry  alinestring)
```


## Description


Adds an isolated edge defined by geometry `alinestring` to a topology connecting two existing isolated nodes `anode` and `anothernode` and returns the edge id of the new edge.


If the spatial reference system (srid) of the `alinestring` geometry is not the same as the topology, any of the input arguments are null, or the nodes are contained in more than one face, or the nodes are start or end nodes of an existing edge, then an exception is thrown.


If the `alinestring` is not within the face of the face the `anode` and `anothernode` belong to, then an exception is thrown.


If the `anode` and `anothernode` are not the start and end points of the `alinestring` then an exception is thrown.


Availability: 1.1


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.4


## Examples


## See Also


[ST_AddIsoNode](#ST_AddIsoNode), [ST_IsSimple](../postgis-reference/geometry-accessors.md#ST_IsSimple), [ST_Within](../postgis-reference/spatial-relationships.md#ST_Within)
  <a id="ST_AddEdgeNewFaces"></a>

# ST_AddEdgeNewFaces

Add a new edge and, if in doing so it splits a face, delete the original face and replace it with two new faces.

## Synopsis


```sql
integer ST_AddEdgeNewFaces(varchar  atopology, integer  anode, integer  anothernode, geometry  acurve)
```


## Description


 Add a new edge and, if in doing so it splits a face, delete the original face and replace it with two new faces. Returns the id of the newly added edge.


 Updates all existing joined edges and relationships accordingly.


If any arguments are null, the given nodes are unknown (must already exist in the `node` table of the topology schema) , the `acurve` is not a `LINESTRING`, the `anode` and `anothernode` are not the start and endpoints of `acurve` then an error is thrown.


If the spatial reference system (srid) of the `acurve` geometry is not the same as the topology an exception is thrown.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.12


## Examples


## See Also


[ST_RemEdgeNewFace](#ST_RemEdgeNewFace)


[ST_AddEdgeModFace](#ST_AddEdgeModFace)
  <a id="ST_AddEdgeModFace"></a>

# ST_AddEdgeModFace

Add a new edge and, if in doing so it splits a face, modify the original face and add a new face.

## Synopsis


```sql
integer ST_AddEdgeModFace(varchar  atopology, integer  anode, integer  anothernode, geometry  acurve)
```


## Description


 Add a new edge and, if doing so splits a face, modify the original face and add a new one.


!!! note

    If possible, the new face will be created on left side of the new edge. This will not be possible if the face on the left side will need to be the Universe face (unbounded).


 Returns the id of the newly added edge.


 Updates all existing joined edges and relationships accordingly.


If any arguments are null, the given nodes are unknown (must already exist in the `node` table of the topology schema) , the `acurve` is not a `LINESTRING`, the `anode` and `anothernode` are not the start and endpoints of `acurve` then an error is thrown.


If the spatial reference system (srid) of the `acurve` geometry is not the same as the topology an exception is thrown.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.13


## Examples


## See Also


[ST_RemEdgeModFace](#ST_RemEdgeModFace)


[ST_AddEdgeNewFaces](#ST_AddEdgeNewFaces)
  <a id="ST_RemEdgeNewFace"></a>

# ST_RemEdgeNewFace

Removes an edge and, if the removed edge separated two faces, delete the original faces and replace them with a new face.

## Synopsis


```sql
integer ST_RemEdgeNewFace(varchar  atopology, integer  anedge)
```


## Description


 Removes an edge and, if the removed edge separated two faces, delete the original faces and replace them with a new face.


 Returns the id of a newly created face or NULL, if no new face is created. No new face is created when the removed edge is dangling or isolated or confined with the universe face (possibly making the universe flood into the face on the other side).


 Updates all existing joined edges and relationships accordingly.


 Refuses to remove an edge participating in the definition of an existing TopoGeometry. Refuses to heal two faces if any TopoGeometry is defined by only one of them (and not the other).


 If any arguments are null, the given edge is unknown (must already exist in the `edge` table of the topology schema), the topology name is invalid then an error is thrown.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.14


## Examples


## See Also


[ST_RemEdgeModFace](#ST_RemEdgeModFace)


[ST_AddEdgeNewFaces](#ST_AddEdgeNewFaces)
  <a id="ST_RemEdgeModFace"></a>

# ST_RemEdgeModFace

Removes an edge, and if the edge separates two faces deletes one face and modifies the other face to cover the space of both.

## Synopsis


```sql
integer ST_RemEdgeModFace(varchar  atopology, integer  anedge)
```


## Description


 Removes an edge, and if the removed edge separates two faces deletes one face and modifies the other face to cover the space of both. Preferentially keeps the face on the right, to be consistent with [ST_AddEdgeModFace](#ST_AddEdgeModFace). Returns the id of the face which is preserved.


 Updates all existing joined edges and relationships accordingly.


 Refuses to remove an edge participating in the definition of an existing TopoGeometry. Refuses to heal two faces if any TopoGeometry is defined by only one of them (and not the other).


 If any arguments are null, the given edge is unknown (must already exist in the `edge` table of the topology schema), the topology name is invalid then an error is thrown.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.15


## Examples


## See Also


[ST_AddEdgeModFace](#ST_AddEdgeModFace)


[ST_RemEdgeNewFace](#ST_RemEdgeNewFace)
  <a id="ST_ChangeEdgeGeom"></a>

# ST_ChangeEdgeGeom

Changes the shape of an edge without affecting the topology structure.

## Synopsis


```sql
text ST_ChangeEdgeGeom(varchar  atopology, integer  anedge, geometry  acurve)
```


## Description


 Changes the shape of an edge without affecting the topology structure.


 If any arguments are null, the given edge does not exist in the `edge` table of the topology schema, the `acurve` is not a `LINESTRING`, or the modification would change the underlying topology then an error is thrown.


If the spatial reference system (srid) of the `acurve` geometry is not the same as the topology an exception is thrown.


If the new `acurve` is not simple, then an error is thrown.


 If moving the edge from old to new position would hit an obstacle then an error is thrown.


Availability: 1.1.0


 Enhanced: 2.0.0 adds topological consistency enforcement


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details X.3.6


## Examples


```sql
SELECT topology.ST_ChangeEdgeGeom('ma_topo', 1,
		ST_GeomFromText('LINESTRING(227591.9 893900.4,227622.6 893844.3,227641.6 893816.6, 227704.5 893778.5)', 26986) );
 ----
 Edge 1 changed
```


## See Also


[ST_AddEdgeModFace](#ST_AddEdgeModFace)


[ST_RemEdgeModFace](#ST_RemEdgeModFace)


[ST_ModEdgeSplit](#ST_ModEdgeSplit)
  <a id="ST_ModEdgeSplit"></a>

# ST_ModEdgeSplit

Split an edge by creating a new node along an existing edge, modifying the original edge and adding a new edge.

## Synopsis


```sql
integer ST_ModEdgeSplit(varchar  atopology, integer  anedge, geometry  apoint)
```


## Description


 Split an edge by creating a new node along an existing edge, modifying the original edge and adding a new edge. Updates all existing joined edges and relationships accordingly. Returns the identifier of the newly added node.


Availability: 1.1


Changed: 2.0 - In prior versions, this was misnamed ST_ModEdgesSplit


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.9


## Examples


```

-- Add an edge --
 SELECT topology.AddEdge('ma_topo', ST_GeomFromText('LINESTRING(227592 893910, 227600 893910)', 26986) ) As edgeid;

-- edgeid-
3


-- Split the edge  --
SELECT topology.ST_ModEdgeSplit('ma_topo',  3, ST_SetSRID(ST_Point(227594,893910),26986)  ) As node_id;
        node_id
-------------------------
7
```


## See Also


 [ST_NewEdgesSplit](#ST_NewEdgesSplit), [ST_ModEdgeHeal](#ST_ModEdgeHeal), [ST_NewEdgeHeal](#ST_NewEdgeHeal), [AddEdge](topology-processing.md#AddEdge)
  <a id="ST_ModEdgeHeal"></a>

# ST_ModEdgeHeal

Heals two edges by deleting the node connecting them, modifying the first edge and deleting the second edge. Returns the id of the deleted node.

## Synopsis


```sql
int ST_ModEdgeHeal(varchar  atopology, integer  anedge, integer  anotheredge)
```


## Description


 Heals two edges by deleting the node connecting them, modifying the first edge and deleting the second edge. Returns the id of the deleted node. Updates all existing joined edges and relationships accordingly.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.9


## See Also


 [ST_ModEdgeSplit](#ST_ModEdgeSplit) [ST_NewEdgesSplit](#ST_NewEdgesSplit)
  <a id="ST_NewEdgeHeal"></a>

# ST_NewEdgeHeal

Heals two edges by deleting the node connecting them, deleting both edges, and replacing them with an edge whose direction is the same as the first edge provided.

## Synopsis


```sql
int ST_NewEdgeHeal(varchar  atopology, integer  anedge, integer  anotheredge)
```


## Description


 Heals two edges by deleting the node connecting them, deleting both edges, and replacing them with an edge whose direction is the same as the first edge provided. Returns the id of the new edge replacing the healed ones. Updates all existing joined edges and relationships accordingly.


Availability: 2.0


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X.3.9


## See Also


 [ST_ModEdgeHeal](#ST_ModEdgeHeal) [ST_ModEdgeSplit](#ST_ModEdgeSplit) [ST_NewEdgesSplit](#ST_NewEdgesSplit)
  <a id="ST_MoveIsoNode"></a>

# ST_MoveIsoNode

Moves an isolated node in a topology from one point to another. If new `apoint` geometry exists as a node an error is thrown. Returns description of move.

## Synopsis


```sql
text ST_MoveIsoNode(varchar  atopology, integer  anode, geometry  apoint)
```


## Description


Moves an isolated node in a topology from one point to another. If new `apoint` geometry exists as a node an error is thrown.


If any arguments are null, the `apoint` is not a point, the existing node is not isolated (is a start or end point of an existing edge), new node location intersects an existing edge (even at the end points) or the new location is in a different face (since 3.2.0) then an exception is thrown.


If the spatial reference system (srid) of the point geometry is not the same as the topology an exception is thrown.


Availability: 2.0.0


 Enhanced: 3.2.0 ensures the nod cannot be moved in a different face


 SQL-MM: Topo-Net Routines: X.3.2


## Examples


```

-- Add an isolated node with no face  --
SELECT topology.ST_AddIsoNode('ma_topo',  NULL, ST_GeomFromText('POINT(227579 893916)', 26986) ) As nodeid;
 nodeid
--------
      7
-- Move the new node --
SELECT topology.ST_MoveIsoNode('ma_topo', 7,  ST_GeomFromText('POINT(227579.5 893916.5)', 26986) ) As descrip;
                      descrip
----------------------------------------------------
Isolated Node 7 moved to location 227579.5,893916.5
```


## See Also


[ST_AddIsoNode](#ST_AddIsoNode)
  <a id="ST_NewEdgesSplit"></a>

# ST_NewEdgesSplit

Split an edge by creating a new node along an existing edge, deleting the original edge and replacing it with two new edges. Returns the id of the new node created that joins the new edges.

## Synopsis


```sql
integer ST_NewEdgesSplit(varchar  atopology, integer  anedge, geometry  apoint)
```


## Description


 Split an edge with edge id `anedge` by creating a new node with point location `apoint` along current edge, deleting the original edge and replacing it with two new edges. Returns the id of the new node created that joins the new edges. Updates all existing joined edges and relationships accordingly.


If the spatial reference system (srid) of the point geometry is not the same as the topology, the `apoint` is not a point geometry, the point is null, the point already exists as a node, the edge does not correspond to an existing edge or the point is not within the edge then an exception is thrown.


Availability: 1.1


 SQL-MM: Topo-Net Routines: X.3.8


## Examples


```

-- Add an edge  --
SELECT topology.AddEdge('ma_topo', ST_GeomFromText('LINESTRING(227575 893917,227592 893900)', 26986) ) As edgeid;
-- result-
edgeid
------
	2
-- Split the new edge --
SELECT topology.ST_NewEdgesSplit('ma_topo', 2,  ST_GeomFromText('POINT(227578.5 893913.5)', 26986) ) As newnodeid;
 newnodeid
---------
       6
```


## See Also


 [ST_ModEdgeSplit](#ST_ModEdgeSplit) [ST_ModEdgeHeal](#ST_ModEdgeHeal) [ST_NewEdgeHeal](#ST_NewEdgeHeal) [AddEdge](topology-processing.md#AddEdge)
  <a id="ST_RemoveIsoNode"></a>

# ST_RemoveIsoNode

Removes an isolated node and returns description of action. If the node is not isolated (is start or end of an edge), then an exception is thrown.

## Synopsis


```sql
text ST_RemoveIsoNode(varchar  atopology, integer  anode)
```


## Description


Removes an isolated node and returns description of action. If the node is not isolated (is start or end of an edge), then an exception is thrown.


Availability: 1.1


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X+1.3.3


## Examples


```

-- Remove an isolated node with no face  --
SELECT topology.ST_RemoveIsoNode('ma_topo',  7 ) As result;
         result
-------------------------
 Isolated node 7 removed
```


## See Also


[ST_AddIsoNode](#ST_AddIsoNode)
  <a id="ST_RemoveIsoEdge"></a>

# ST_RemoveIsoEdge

Removes an isolated edge and returns description of action. If the edge is not isolated, then an exception is thrown.

## Synopsis


```sql
text ST_RemoveIsoEdge(varchar  atopology, integer  anedge)
```


## Description


Removes an isolated edge and returns description of action. If the edge is not isolated, then an exception is thrown.


Availability: 1.1


 SQL-MM: Topo-Geo and Topo-Net 3: Routine Details: X+1.3.3


## Examples


```

-- Remove an isolated node with no face  --
SELECT topology.ST_RemoveIsoNode('ma_topo',  7 ) As result;
         result
-------------------------
 Isolated node 7 removed
```


## See Also


[ST_AddIsoNode](#ST_AddIsoNode)
