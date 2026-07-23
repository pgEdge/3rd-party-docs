<a id="TopoGeometry_Constructors"></a>

## TopoGeometry Constructors
  <a id="CreateTopoGeom"></a>

# CreateTopoGeom

Creates a new topo geometry object from topo element array - tg_type: 1:[multi]point, 2:[multi]line, 3:[multi]poly, 4:collection

## Synopsis


```sql
topogeometry CreateTopoGeom(varchar  toponame, integer  tg_type, integer layer_id, topoelementarray tg_objs)
topogeometry CreateTopoGeom(varchar  toponame, integer  tg_type, integer layer_id)
```


## Description


Creates a topogeometry object for layer denoted by `layer_id` and registers it in the relations table in the `toponame` schema.


`tg_type` is an integer: 1:[multi]point (punctal), 2:[multi]line (lineal), 3:[multi]poly (areal), 4:collection. `layer_id` is the layer id in the topology.layer table.


punctal layers are formed from set of nodes, lineal layers are formed from a set of edges, areal layers are formed from a set of faces, and collections can be formed from a mixture of nodes, edges, and faces.


Omitting the array of components generates an empty TopoGeometry object.


Availability: 1.1


## Examples: Form from existing edges


Create a topogeom in ri_topo schema for layer 2 (our ri_roads), of type (2) LINE, for the first edge (we loaded in `ST_CreateTopoGeo`).


```sql
INSERT INTO ri.ri_roads(road_name, topo) VALUES('Unknown', topology.CreateTopoGeom('ri_topo',2,2,'{{1,2}}'::topology.topoelementarray);
```


## Examples: Convert an areal geometry to best guess topogeometry


Lets say we have geometries that should be formed from a collection of faces. We have for example blockgroups table and want to know the topo geometry of each block group. If our data was perfectly aligned, we could do this:


```

-- create our topo geometry column --
SELECT topology.AddTopoGeometryColumn(
	'topo_boston',
	'boston', 'blockgroups', 'topo', 'POLYGON');

-- addtopgeometrycolumn --
1

-- update our column assuming
-- everything is perfectly aligned with our edges
UPDATE boston.blockgroups AS bg
	SET topo = topology.CreateTopoGeom('topo_boston'
        ,3,1
        , foo.bfaces)
FROM (SELECT b.gid,  topology.TopoElementArray_Agg(ARRAY[f.face_id,3]) As bfaces
	FROM boston.blockgroups As b
            INNER JOIN topo_boston.face As f ON b.geom && f.mbr
        WHERE ST_Covers(b.geom, topology.ST_GetFaceGeometry('topo_boston', f.face_id))
            GROUP BY b.gid) As foo
WHERE foo.gid = bg.gid;
```


```

--the world is rarely perfect allow for some error
--count the face if 50% of it falls
-- within what we think is our blockgroup boundary
UPDATE boston.blockgroups AS bg
	SET topo = topology.CreateTopoGeom('topo_boston'
        ,3,1
        , foo.bfaces)
FROM (SELECT b.gid,  topology.TopoElementArray_Agg(ARRAY[f.face_id,3]) As bfaces
	FROM boston.blockgroups As b
            INNER JOIN topo_boston.face As f ON b.geom && f.mbr
        WHERE ST_Covers(b.geom, topology.ST_GetFaceGeometry('topo_boston', f.face_id))
	OR
 (  ST_Intersects(b.geom, topology.ST_GetFaceGeometry('topo_boston', f.face_id))
            AND ST_Area(ST_Intersection(b.geom, topology.ST_GetFaceGeometry('topo_boston', f.face_id) ) ) >
                ST_Area(topology.ST_GetFaceGeometry('topo_boston', f.face_id))*0.5
                )
            GROUP BY b.gid) As foo
WHERE foo.gid = bg.gid;

-- and if we wanted to convert our topogeometry back
-- to a denormalized geometry aligned with our faces and edges
-- cast the topo to a geometry
-- The really cool thing is my new geometries
-- are now aligned with my tiger street centerlines
UPDATE boston.blockgroups SET new_geom = topo::geometry;
```


## See Also


 [AddTopoGeometryColumn](topology-and-topogeometry-management.md#AddTopoGeometryColumn), [toTopoGeom](#toTopoGeom) [ST_CreateTopoGeo](topology-constructors.md#ST_CreateTopoGeo), [ST_GetFaceGeometry](topology-accessors.md#ST_GetFaceGeometry), [topoelementarray](topology-domains.md#topoelementarray), [TopoElementArray_Agg](#TopoElementArray_Agg)
  <a id="toTopoGeom"></a>

# toTopoGeom

Converts a simple Geometry into a topo geometry.

## Synopsis


```sql
topogeometry toTopoGeom(geometry  geom, varchar  toponame, integer layer_id, float8 tolerance)
topogeometry toTopoGeom(geometry  geom, topogeometry  topogeom, float8 tolerance)
```


## Description


 Converts a simple Geometry into a [topogeometry](topology-types.md#topogeometry).


 Topological primitives required to represent the input geometry will be added to the underlying topology, possibly splitting existing ones, and they will be associated with the output TopoGeometry in the `relation` table.


 Existing TopoGeometry objects (with the possible exception of `topogeom`, if given) will retain their shapes.


 When `tolerance` is given it will be used to snap the input geometry to existing primitives.


 In the first form a new TopoGeometry will be created for the given layer (`layer_id`) of the given topology (`toponame`).


 In the second form the primitives resulting from the conversion will be added to the pre-existing TopoGeometry (`topogeom`), possibly adding space to its final shape. To have the new shape completely replace the old one see [clearTopoGeom](topogeometry-editors.md#clearTopoGeom).


Availability: 2.0


Enhanced: 2.1.0 adds the version taking an existing TopoGeometry.


## Examples


This is a full self-contained workflow


```
 -- do this if you don't have a topology setup already
-- creates topology not allowing any tolerance
SELECT topology.CreateTopology('topo_boston_test', 2249);
-- create a new table
CREATE TABLE nei_topo(gid serial primary key, nei varchar(30));
--add a topogeometry column to it
SELECT topology.AddTopoGeometryColumn('topo_boston_test', 'public', 'nei_topo', 'topo', 'MULTIPOLYGON') As new_layer_id;
new_layer_id
-----------
1

--use new layer id in populating the new topogeometry column
-- we add the topogeoms to the new layer with 0 tolerance
INSERT INTO nei_topo(nei, topo)
SELECT nei,  topology.toTopoGeom(geom, 'topo_boston_test', 1)
FROM neighborhoods
WHERE gid BETWEEN 1 and 15;

--use to verify what has happened --
SELECT * FROM
    topology.TopologySummary('topo_boston_test');

-- summary--
Topology topo_boston_test (5), SRID 2249, precision 0
61 nodes, 87 edges, 35 faces, 15 topogeoms in 1 layers
Layer 1, type Polygonal (3), 15 topogeoms
 Deploy: public.nei_topo.topo
```


```


-- Shrink all TopoGeometry polygons by 10 meters
UPDATE nei_topo SET topo = ST_Buffer(clearTopoGeom(topo), -10);

-- Get the no-one-lands left by the above operation
-- I think GRASS calls this "polygon0 layer"
SELECT ST_GetFaceGeometry('topo_boston_test', f.face_id)
  FROM topo_boston_test.face f
  WHERE f.face_id > 0 -- don't consider the universe face
  AND NOT EXISTS ( -- check that no TopoGeometry references the face
    SELECT * FROM topo_boston_test.relation
    WHERE layer_id = 1 AND element_id = f.face_id
  );

```


## See Also


 [CreateTopology](topology-constructors.md#CreateTopology), [AddTopoGeometryColumn](topology-and-topogeometry-management.md#AddTopoGeometryColumn), [CreateTopoGeom](#CreateTopoGeom), [TopologySummary](topology-and-topogeometry-management.md#TopologySummary), [clearTopoGeom](topogeometry-editors.md#clearTopoGeom)
  <a id="TopoElementArray_Agg"></a>

# TopoElementArray_Agg

Returns a `topoelementarray` for a set of element_id, type arrays (topoelements).

## Synopsis


```sql
topoelementarray TopoElementArray_Agg(topoelement set tefield)
```


## Description


Used to create a [topoelementarray](topology-domains.md#topoelementarray) from a set of [topoelement](topology-domains.md#topoelement).


Availability: 2.0.0


## Examples


```sql
SELECT topology.TopoElementArray_Agg(ARRAY[e,t]) As tea
  FROM generate_series(1,3) As e CROSS JOIN generate_series(1,4) As t;
  tea
--------------------------------------------------------------------------
{{1,1},{1,2},{1,3},{1,4},{2,1},{2,2},{2,3},{2,4},{3,1},{3,2},{3,3},{3,4}}
```


## See Also


[topoelement](topology-domains.md#topoelement), [topoelementarray](topology-domains.md#topoelementarray)
  <a id="TopoElement"></a>

# TopoElement

Converts a topogeometry to a topoelement.

## Synopsis


```sql
topoelement TopoElement(topogeometry  topo)
```


## Description


Converts a [topogeometry](topology-types.md#topogeometry) to a [topoelement](topology-domains.md#topoelement).


Availability: 3.4.0


## Examples


This is a full self-contained workflow


```
-- do this if you don't have a topology setup already
-- Creates topology not allowing any tolerance
SELECT TopoElement(topo)
FROM neighborhoods;
```


```
-- using as cast
SELECT topology.TopoElementArray_Agg(topo::topoelement)
FROM neighborhoods
GROUP BY city;
```


## See Also


[TopoElementArray_Agg](#TopoElementArray_Agg), [topogeometry](topology-types.md#topogeometry), [topoelement](topology-domains.md#topoelement)
