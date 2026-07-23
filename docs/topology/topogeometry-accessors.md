<a id="TopoGeom_Accessors"></a>

## TopoGeometry Accessors
  <a id="GetTopoGeomElementArray"></a>

# GetTopoGeomElementArray

Returns a `topoelementarray` (an array of topoelements) containing the topological elements and type of the given TopoGeometry (primitive elements).

## Synopsis


```sql
topoelementarray GetTopoGeomElementArray(varchar  toponame, integer  layer_id, integer tg_id)
```


```sql
topoelementarray GetTopoGeomElementArray(topogeometry  tg)
```


## Description


Returns a [topoelementarray](topology-domains.md#topoelementarray) containing the topological elements and type of the given TopoGeometry (primitive elements). This is similar to GetTopoGeomElements except it returns the elements as an array rather than as a dataset.


tg_id is the topogeometry id of the topogeometry object in the topology in the layer denoted by `layer_id` in the topology.layer table.


Availability: 1.1


## Examples


## See Also


[GetTopoGeomElements](#GetTopoGeomElements), [topoelementarray](topology-domains.md#topoelementarray)
  <a id="GetTopoGeomElements"></a>

# GetTopoGeomElements

Returns a set of `topoelement` objects containing the topological element_id,element_type of the given TopoGeometry (primitive elements).

## Synopsis


```sql
setof topoelement GetTopoGeomElements(varchar  toponame, integer  layer_id, integer tg_id)
```


```sql
setof topoelement GetTopoGeomElements(topogeometry  tg)
```


## Description


Returns a set of element_id,element_type (topoelements) corresponding to primitive topology elements [topoelement](topology-domains.md#topoelement) (1: nodes, 2: edges, 3: faces) that a given topogeometry object in `toponame` schema is composed of.


tg_id is the topogeometry id of the topogeometry object in the topology in the layer denoted by `layer_id` in the topology.layer table.


Availability: 2.0.0


## Examples


## See Also


 [GetTopoGeomElementArray](#GetTopoGeomElementArray), [topoelement](topology-domains.md#topoelement), [TopoGeom_addElement](topogeometry-editors.md#TopoGeom_addElement), [TopoGeom_remElement](topogeometry-editors.md#TopoGeom_remElement)
  <a id="TG_ST_SRID"></a>

# ST_SRID

Returns the spatial reference identifier for a topogeometry.

## Synopsis


```sql
integer ST_SRID(topogeometry  tg)
```


## Description


Returns the spatial reference identifier for the ST_Geometry as defined in spatial_ref_sys table. [Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys)


!!! note

    spatial_ref_sys table is a table that catalogs all spatial reference systems known to PostGIS and is used for transformations from one spatial reference system to another. So verifying you have the right spatial reference system identifier is important if you plan to ever transform your geometries.


Availability: 3.2.0


 SQL-MM 3: 14.1.5


## Examples


```sql
SELECT ST_SRID(ST_GeomFromText('POINT(-71.1043 42.315)',4326));
		--result
		4326

```


## See Also


[Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys), [ST_SetSRID](../postgis-reference/spatial-reference-system-functions.md#ST_SetSRID), [ST_Transform](../postgis-reference/spatial-reference-system-functions.md#ST_Transform), [ST_SRID](../postgis-reference/spatial-reference-system-functions.md#ST_SRID)
