<a id="BBox_Functions"></a>

## Bounding Box Functions
  <a id="Box2D"></a>

# Box2D

Returns a BOX2D representing the 2D extent of a geometry.

## Synopsis


```sql
box2d Box2D(geometry  geom)
```


## Description


Returns a [box2d_type](postgis-geometry-geography-box-data-types.md#box2d_type) representing the 2D extent of the geometry.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


## Examples


```sql
SELECT Box2D(ST_GeomFromText('LINESTRING(1 2, 3 4, 5 6)'));

box2d
---------
BOX(1 2,5 6)
```


```sql

SELECT Box2D(ST_GeomFromText('CIRCULARSTRING(220268 150415,220227 150505,220227 150406)'));

box2d
--------
BOX(220186.984375 150406,220288.25 150506.140625)
```


## See Also


[Box3D](#Box3D), [ST_GeomFromText](geometry-input.md#ST_GeomFromText)
  <a id="Box3D"></a>

# Box3D

Returns a BOX3D representing the 3D extent of a geometry.

## Synopsis


```sql
box3d Box3D(geometry  geom)
```


## Description


Returns a [box3d_type](postgis-geometry-geography-box-data-types.md#box3d_type) representing the 3D extent of the geometry.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


## Examples


```sql
SELECT Box3D(ST_GeomFromEWKT('LINESTRING(1 2 3, 3 4 5, 5 6 5)'));

Box3d
---------
BOX3D(1 2 3,5 6 5)
```


```sql

SELECT Box3D(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 1,220227 150406 1)'));

Box3d
--------
BOX3D(220227 150406 1,220268 150415 1)
```


## See Also


[Box2D](#Box2D), [ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT)
  <a id="ST_EstimatedExtent"></a>

# ST_EstimatedExtent

Returns the estimated extent of a spatial table.

## Synopsis


```sql
box2d ST_EstimatedExtent(text  schema_name, text  table_name, text  geocolumn_name, boolean  parent_only)
box2d ST_EstimatedExtent(text  schema_name, text  table_name, text  geocolumn_name)
box2d ST_EstimatedExtent(text  table_name, text  geocolumn_name)
```


## Description


Returns the estimated extent of a spatial table as a [box2d_type](postgis-geometry-geography-box-data-types.md#box2d_type). The current schema is used if not specified. The estimated extent is taken from the geometry column's statistics. This is usually much faster than computing the exact extent of the table using [ST_Extent](#ST_Extent) or [ST_3DExtent](#ST_3DExtent).


 The default behavior is to also use statistics collected from child tables (tables with INHERITS) if available. If `parent_only` is set to TRUE, only statistics for the given table are used and child tables are ignored.


For PostgreSQL >= 8.0.0 statistics are gathered by VACUUM ANALYZE and the result extent will be about 95% of the actual one. For PostgreSQL < 8.0.0 statistics are gathered by running <code>update_geometry_stats()</code> and the result extent is exact.


!!! note

    In the absence of statistics (empty table or no ANALYZE called) this function returns NULL. Prior to version 1.5.4 an exception was thrown instead.


!!! note

    Escaping names for tables and/or namespaces that include special characters and quotes may require special handling. A user notes: "For schemas and tables, use identifier escaping rules to produce a double-quoted string, and afterwards remove the first and last double-quote character. For geometry column pass as is."


Availability: 1.0.0


Changed: 2.1.0. Up to 2.0.x this was called ST_Estimated_Extent.


## Examples


```sql
SELECT ST_EstimatedExtent('ny', 'edges', 'geom');
--result--
BOX(-8877653 4912316,-8010225.5 5589284)

SELECT ST_EstimatedExtent('feature_poly', 'geom');
--result--
BOX(-124.659652709961 24.6830825805664,-67.7798080444336 49.0012092590332)

```


## See Also


[ST_Extent](#ST_Extent), [ST_3DExtent](#ST_3DExtent)
  <a id="ST_Expand"></a>

# ST_Expand

Returns a bounding box expanded from another bounding box or a geometry.

## Synopsis


```sql
geometry ST_Expand(geometry  geom, float units_to_expand)
geometry ST_Expand(geometry  geom, float dx, float dy, float dz=0, float dm=0)
box2d ST_Expand(box2d  box, float units_to_expand)
box2d ST_Expand(box2d  box, float dx, float dy)
box3d ST_Expand(box3d  box, float units_to_expand)
box3d ST_Expand(box3d  box, float  dx, float  dy, float  dz=0)
```


## Description


Returns a bounding box expanded from the bounding box of the input, either by specifying a single distance with which the box should be expanded on both axes, or by specifying an expansion distance for each axis. Uses double-precision. Can be used for distance queries, or to add a bounding box filter to a query to take advantage of a spatial index.


In addition to the version of ST_Expand accepting and returning a geometry, variants are provided that accept and return [box2d_type](postgis-geometry-geography-box-data-types.md#box2d_type) and [box3d_type](postgis-geometry-geography-box-data-types.md#box3d_type) data types.


Distances are in the units of the spatial reference system of the input.


ST_Expand is similar to [ST_Buffer](geometry-processing.md#ST_Buffer), except while buffering expands a geometry in all directions, ST_Expand expands the bounding box along each axis.


!!! note

    Pre version 1.3, ST_Expand was used in conjunction with [ST_Distance](measurement-functions.md#ST_Distance) to do indexable distance queries. For example, <code>geom && ST_Expand('POINT(10 20)', 10) AND ST_Distance(geom, 'POINT(10 20)') < 10</code>. This has been replaced by the simpler and more efficient [ST_DWithin](spatial-relationships.md#ST_DWithin) function.


Availability: 1.5.0 behavior changed to output double precision instead of float4 coordinates.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Enhanced: 2.3.0 support was added to expand a box by different amounts in different dimensions.


## Examples


!!! note

    Examples below use US National Atlas Equal Area (SRID=2163) which is a meter projection


```


--10 meter expanded box around bbox of a linestring
SELECT CAST(ST_Expand(ST_GeomFromText('LINESTRING(2312980 110676,2312923 110701,2312892 110714)', 2163),10) As box2d);
					 st_expand
------------------------------------
 BOX(2312882 110666,2312990 110724)

--10 meter expanded 3D box of a 3D box
SELECT ST_Expand(CAST('BOX3D(778783 2951741 1,794875 2970042.61545891 10)' As box3d),10)
							  st_expand
-----------------------------------------------------
 BOX3D(778773 2951731 -9,794885 2970052.61545891 20)

 --10 meter geometry astext rep of a expand box around a point geometry
 SELECT ST_AsEWKT(ST_Expand(ST_GeomFromEWKT('SRID=2163;POINT(2312980 110676)'),10));
											st_asewkt
-------------------------------------------------------------------------------------------------
 SRID=2163;POLYGON((2312970 110666,2312970 110686,2312990 110686,2312990 110666,2312970 110666))


```


## See Also


[ST_Buffer](geometry-processing.md#ST_Buffer), [ST_DWithin](spatial-relationships.md#ST_DWithin), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_Extent"></a>

# ST_Extent

Aggregate function that returns the bounding box of geometries.

## Synopsis


```sql
box2d ST_Extent(geometry set geomfield)
```


## Description


An aggregate function that returns a [box2d_type](postgis-geometry-geography-box-data-types.md#box2d_type) bounding box that bounds a set of geometries.


The bounding box coordinates are in the spatial reference system of the input geometries.


ST_Extent is similar in concept to Oracle Spatial/Locator's SDO_AGGR_MBR.


!!! note

    ST_Extent returns boxes with only X and Y ordinates even with 3D geometries. To return XYZ ordinates use [ST_3DExtent](#ST_3DExtent).


!!! note

    The returned <code>box3d</code> value does not include a SRID. Use [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID) to convert it into a geometry with SRID metadata. The SRID is the same as the input geometries.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


## Examples


!!! note

    Examples below use Massachusetts State Plane ft (SRID=2249)


```sql


SELECT ST_Extent(geom) as bextent FROM sometable;
					 st_bextent
------------------------------------
BOX(739651.875 2908247.25,794875.8125 2970042.75)


--Return extent of each category of geometries
SELECT ST_Extent(geom) as bextent
FROM sometable
GROUP BY category ORDER BY category;

					  bextent                       |         name
----------------------------------------------------+----------------
 BOX(778783.5625 2951741.25,794875.8125 2970042.75) | A
 BOX(751315.8125 2919164.75,765202.6875 2935417.25) | B
 BOX(739651.875 2917394.75,756688.375 2935866)      | C

 --Force back into a geometry
 -- and render the extended text representation of that geometry
SELECT ST_SetSRID(ST_Extent(geom),2249) as bextent FROM sometable;

				bextent
--------------------------------------------------------------------------------
 SRID=2249;POLYGON((739651.875 2908247.25,739651.875 2970042.75,794875.8125 2970042.75,
 794875.8125 2908247.25,739651.875 2908247.25))

```


## See Also


 [ST_EstimatedExtent](#ST_EstimatedExtent), [ST_3DExtent](#ST_3DExtent), [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID)
  <a id="ST_3DExtent"></a>

# ST_3DExtent

Aggregate function that returns the 3D bounding box of geometries.

## Synopsis


```sql
box3d ST_3DExtent(geometry set geomfield)
```


## Description


An aggregate function that returns a [box3d_type](postgis-geometry-geography-box-data-types.md#box3d_type) (includes Z ordinate) bounding box that bounds a set of geometries.


The bounding box coordinates are in the spatial reference system of the input geometries.


!!! note

    The returned <code>box3d</code> value does not include a SRID. Use [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID) to convert it into a geometry with SRID metadata. The SRID is the same as the input geometries.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Changed: 2.0.0 In prior versions this used to be called ST_Extent3D


## Examples


```sql

SELECT ST_3DExtent(foo.geom) As b3extent
FROM (SELECT ST_MakePoint(x,y,z) As geom
	FROM generate_series(1,3) As x
		CROSS JOIN generate_series(1,2) As y
		CROSS JOIN generate_series(0,2) As Z) As foo;
	  b3extent
--------------------
 BOX3D(1 1 0,3 2 2)

--Get the extent of various elevated circular strings
SELECT ST_3DExtent(foo.geom) As b3extent
FROM (SELECT ST_Translate(ST_Force_3DZ(ST_LineToCurve(ST_Buffer(ST_Point(x,y),1))),0,0,z) As geom
	FROM generate_series(1,3) As x
		CROSS JOIN generate_series(1,2) As y
		CROSS JOIN generate_series(0,2) As Z) As foo;

	b3extent
--------------------
 BOX3D(1 0 0,4 2 2)

```


## See Also


[ST_Extent](#ST_Extent), [ST_Force_3DZ](geometry-editors.md#ST_Force_3DZ), [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID)
  <a id="ST_MakeBox2D"></a>

# ST_MakeBox2D

Creates a BOX2D defined by two 2D point geometries.

## Synopsis


```sql
box2d ST_MakeBox2D(geometry  pointLowLeft, geometry  pointUpRight)
```


## Description


Creates a [box2d_type](postgis-geometry-geography-box-data-types.md#box2d_type) defined by two Point geometries. This is useful for doing range queries.


## Examples


```


--Return all features that fall reside or partly reside in a US national atlas coordinate bounding box
--It is assumed here that the geometries are stored with SRID = 2163 (US National atlas equal area)
SELECT feature_id, feature_name, geom
FROM features
WHERE geom && ST_SetSRID(ST_MakeBox2D(ST_Point(-989502.1875, 528439.5625),
	ST_Point(-987121.375 ,529933.1875)),2163)
```


## See Also


[ST_Point](geometry-constructors.md#ST_Point), [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_3DMakeBox"></a>

# ST_3DMakeBox

Creates a BOX3D defined by two 3D point geometries.

## Synopsis


```sql
box3d ST_3DMakeBox(geometry  point3DLowLeftBottom, geometry  point3DUpRightTop)
```


## Description


Creates a [box3d_type](postgis-geometry-geography-box-data-types.md#box3d_type) defined by two 3D Point geometries.


![image](images/check.png)
   This function supports 3D and will not drop the z-index.


Changed: 2.0.0 In prior versions this used to be called ST_MakeBox3D


## Examples


```sql

SELECT ST_3DMakeBox(ST_MakePoint(-989502.1875, 528439.5625, 10),
	ST_MakePoint(-987121.375 ,529933.1875, 10)) As abb3d

--bb3d--
--------
BOX3D(-989502.1875 528439.5625 10,-987121.375 529933.1875 10)

```


## See Also


[ST_MakePoint](geometry-constructors.md#ST_MakePoint), [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_XMax"></a>

# ST_XMax

Returns the X maxima of a 2D or 3D bounding box or a geometry.

## Synopsis


```sql
float ST_XMax(box3d  aGeomorBox2DorBox3D)
```


## Description


Returns the X maxima of a 2D or 3D bounding box or a geometry.


!!! note

    Although this function is only defined for box3d, it also works for box2d and geometry values due to automatic casting. However, it will not accept a geometry or box2d text representation, since those do not auto-cast.


## Examples


```sql
SELECT ST_XMax('BOX3D(1 2 3, 4 5 6)');
st_xmax
-------
4

SELECT ST_XMax(ST_GeomFromText('LINESTRING(1 3 4, 5 6 7)'));
st_xmax
-------
5

SELECT ST_XMax(CAST('BOX(-3 2, 3 4)' As box2d));
st_xmax
-------
3
--Observe THIS DOES NOT WORK because it will try to auto-cast the string representation to a BOX3D
SELECT ST_XMax('LINESTRING(1 3, 5 6)');

--ERROR:  BOX3D parser - doesn't start with BOX3D(

SELECT ST_XMax(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)'));
st_xmax
--------
220288.248780547

```


## See Also


[ST_XMin](#ST_XMin), [ST_YMax](#ST_YMax), [ST_YMin](#ST_YMin), [ST_ZMax](#ST_ZMax), [ST_ZMin](#ST_ZMin)
  <a id="ST_XMin"></a>

# ST_XMin

Returns the X minima of a 2D or 3D bounding box or a geometry.

## Synopsis


```sql
float ST_XMin(box3d  aGeomorBox2DorBox3D)
```


## Description


Returns the X minima of a 2D or 3D bounding box or a geometry.


!!! note

    Although this function is only defined for box3d, it also works for box2d and geometry values due to automatic casting. However it will not accept a geometry or box2d text representation, since those do not auto-cast.


## Examples


```sql
SELECT ST_XMin('BOX3D(1 2 3, 4 5 6)');
st_xmin
-------
1

SELECT ST_XMin(ST_GeomFromText('LINESTRING(1 3 4, 5 6 7)'));
st_xmin
-------
1

SELECT ST_XMin(CAST('BOX(-3 2, 3 4)' As box2d));
st_xmin
-------
-3
--Observe THIS DOES NOT WORK because it will try to auto-cast the string representation to a BOX3D
SELECT ST_XMin('LINESTRING(1 3, 5 6)');

--ERROR:  BOX3D parser - doesn't start with BOX3D(

SELECT ST_XMin(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)'));
st_xmin
--------
220186.995121892

```


## See Also


[ST_XMax](#ST_XMax), [ST_YMax](#ST_YMax), [ST_YMin](#ST_YMin), [ST_ZMax](#ST_ZMax), [ST_ZMin](#ST_ZMin)
  <a id="ST_YMax"></a>

# ST_YMax

Returns the Y maxima of a 2D or 3D bounding box or a geometry.

## Synopsis


```sql
float ST_YMax(box3d  aGeomorBox2DorBox3D)
```


## Description


Returns the Y maxima of a 2D or 3D bounding box or a geometry.


!!! note

    Although this function is only defined for box3d, it also works for box2d and geometry values due to automatic casting. However it will not accept a geometry or box2d text representation, since those do not auto-cast.


## Examples


```sql
SELECT ST_YMax('BOX3D(1 2 3, 4 5 6)');
st_ymax
-------
5

SELECT ST_YMax(ST_GeomFromText('LINESTRING(1 3 4, 5 6 7)'));
st_ymax
-------
6

SELECT ST_YMax(CAST('BOX(-3 2, 3 4)' As box2d));
st_ymax
-------
4
--Observe THIS DOES NOT WORK because it will try to auto-cast the string representation to a BOX3D
SELECT ST_YMax('LINESTRING(1 3, 5 6)');

--ERROR:  BOX3D parser - doesn't start with BOX3D(

SELECT ST_YMax(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)'));
st_ymax
--------
150506.126829327

```


## See Also


[ST_XMin](#ST_XMin), [ST_XMax](#ST_XMax), [ST_YMin](#ST_YMin), [ST_ZMax](#ST_ZMax), [ST_ZMin](#ST_ZMin)
  <a id="ST_YMin"></a>

# ST_YMin

Returns the Y minima of a 2D or 3D bounding box or a geometry.

## Synopsis


```sql
float ST_YMin(box3d  aGeomorBox2DorBox3D)
```


## Description


Returns the Y minima of a 2D or 3D bounding box or a geometry.


!!! note

    Although this function is only defined for box3d, it also works for box2d and geometry values due to automatic casting. However it will not accept a geometry or box2d text representation, since those do not auto-cast.


## Examples


```sql
SELECT ST_YMin('BOX3D(1 2 3, 4 5 6)');
st_ymin
-------
2

SELECT ST_YMin(ST_GeomFromText('LINESTRING(1 3 4, 5 6 7)'));
st_ymin
-------
3

SELECT ST_YMin(CAST('BOX(-3 2, 3 4)' As box2d));
st_ymin
-------
2
--Observe THIS DOES NOT WORK because it will try to auto-cast the string representation to a BOX3D
SELECT ST_YMin('LINESTRING(1 3, 5 6)');

--ERROR:  BOX3D parser - doesn't start with BOX3D(

SELECT ST_YMin(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)'));
st_ymin
--------
150406

```


## See Also


[ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT), [ST_XMin](#ST_XMin), [ST_XMax](#ST_XMax), [ST_YMax](#ST_YMax), [ST_ZMax](#ST_ZMax), [ST_ZMin](#ST_ZMin)
  <a id="ST_ZMax"></a>

# ST_ZMax

Returns the Z maxima of a 2D or 3D bounding box or a geometry.

## Synopsis


```sql
float ST_ZMax(box3d  aGeomorBox2DorBox3D)
```


## Description


Returns the Z maxima of a 2D or 3D bounding box or a geometry.


!!! note

    Although this function is only defined for box3d, it also works for box2d and geometry values due to automatic casting. However it will not accept a geometry or box2d text representation, since those do not auto-cast.


## Examples


```sql
SELECT ST_ZMax('BOX3D(1 2 3, 4 5 6)');
st_zmax
-------
6

SELECT ST_ZMax(ST_GeomFromEWKT('LINESTRING(1 3 4, 5 6 7)'));
st_zmax
-------
7

SELECT ST_ZMax('BOX3D(-3 2 1, 3 4 1)' );
st_zmax
-------
1
--Observe THIS DOES NOT WORK because it will try to auto-cast the string representation to a BOX3D
SELECT ST_ZMax('LINESTRING(1 3 4, 5 6 7)');

--ERROR:  BOX3D parser - doesn't start with BOX3D(

SELECT ST_ZMax(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)'));
st_zmax
--------
3

```


## See Also


[ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT), [ST_XMin](#ST_XMin), [ST_XMax](#ST_XMax), [ST_YMax](#ST_YMax), [ST_YMin](#ST_YMin), [ST_ZMax](#ST_ZMax)
  <a id="ST_ZMin"></a>

# ST_ZMin

Returns the Z minima of a 2D or 3D bounding box or a geometry.

## Synopsis


```sql
float ST_ZMin(box3d  aGeomorBox2DorBox3D)
```


## Description


Returns the Z minima of a 2D or 3D bounding box or a geometry.


!!! note

    Although this function is only defined for box3d, it also works for box2d and geometry values due to automatic casting. However it will not accept a geometry or box2d text representation, since those do not auto-cast.


## Examples


```sql
SELECT ST_ZMin('BOX3D(1 2 3, 4 5 6)');
st_zmin
-------
3

SELECT ST_ZMin(ST_GeomFromEWKT('LINESTRING(1 3 4, 5 6 7)'));
st_zmin
-------
4

SELECT ST_ZMin('BOX3D(-3 2 1, 3 4 1)' );
st_zmin
-------
1
--Observe THIS DOES NOT WORK because it will try to auto-cast the string representation to a BOX3D
SELECT ST_ZMin('LINESTRING(1 3 4, 5 6 7)');

--ERROR:  BOX3D parser - doesn't start with BOX3D(

SELECT ST_ZMin(ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)'));
st_zmin
--------
1

```


## See Also


[ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT), [ST_GeomFromText](geometry-input.md#ST_GeomFromText), [ST_XMin](#ST_XMin), [ST_XMax](#ST_XMax), [ST_YMax](#ST_YMax), [ST_YMin](#ST_YMin), [ST_ZMax](#ST_ZMax)
