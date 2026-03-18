<a id="PostGIS_Types"></a>

## PostGIS Geometry/Geography/Box Data Types
  <a id="box2d_type"></a>

# box2d

The type representing a 2-dimensional bounding box.

## Description


`box2d` is a spatial data type used to represent the two-dimensional bounding box enclosing a geometry or collection of geometries. For example, the [ST_Extent](bounding-box-functions.md#ST_Extent) aggregate function returns a `box2d` object.


The representation contains the values `xmin, ymin, xmax, ymax`. These are the minimum and maximum values of the X and Y extents.


`box2d` objects have a text representation which looks like <code>BOX(1 2,5 6)</code>.


## Casting Behavior


This table lists the automatic and explicit casts allowed for this data type:


| Cast To | Behavior |
| box3d | automatic |
| geometry | automatic |


## See Also


[PostGIS Box Functions](../postgis-special-functions-index/postgis-box-functions.md#PostGIS_BoxFunctions)
  <a id="box3d_type"></a>

# box3d

The type representing a 3-dimensional bounding box.

## Description


`box3d` is a PostGIS spatial data type used to represent the three-dimensional bounding box enclosing a geometry or collection of geometries. For example, the [ST_3DExtent](bounding-box-functions.md#ST_3DExtent) aggregate function returns a `box3d` object.


The representation contains the values `xmin, ymin, zmin, xmax, ymax, zmax`. These are the minimum and maximum values of the X, Y and Z extents.


`box3d` objects have a text representation which looks like <code>BOX3D(1 2 3,5 6 5)</code>.


## Casting Behavior


This table lists the automatic and explicit casts allowed for this data type:


| Cast To | Behavior |
| box | automatic |
| box2d | automatic |
| geometry | automatic |


## See Also


[PostGIS Box Functions](../postgis-special-functions-index/postgis-box-functions.md#PostGIS_BoxFunctions)
  <a id="geometry"></a>

# geometry

The type representing spatial features with planar coordinate systems.

## Description


`geometry` is a fundamental PostGIS spatial data type used to represent a feature in planar (Euclidean) coordinate systems.


All spatial operations on geometry use the units of the Spatial Reference System the geometry is in.


## Casting Behavior


This table lists the automatic and explicit casts allowed for this data type:


| Cast To | Behavior |
| box | automatic |
| box2d | automatic |
| box3d | automatic |
| bytea | automatic |
| geography | automatic |
| text | automatic |


## See Also


[Spatial Data Model](../data-management/spatial-data-model.md#RefObject), [PostGIS SQL-MM Compliant Functions](../postgis-special-functions-index/postgis-sql-mm-compliant-functions.md#PostGIS_SQLMM_Functions)
  <a id="geometry_dump"></a>

# geometry_dump

A composite type used to describe the parts of complex geometry.

## Description


`geometry_dump` is a [composite data type](https://www.postgresql.org/docs/current/rowtypes.html) containing the fields:


- `geom` - a geometry representing a component of the dumped geometry. The geometry type depends on the originating function.
- `path[]` - an integer array that defines the navigation path within the dumped geometry to the `geom` component. The path array is 1-based (i.e. `path[1]` is the first element.)


 It is used by the `ST_Dump*` family of functions as an output type to explode a complex geometry into its constituent parts.


## See Also


[PostGIS Geometry / Geography / Raster Dump Functions](../postgis-special-functions-index/postgis-geometry-geography-raster-dump-functions.md#PostGIS_Geometry_DumpFunctions)
  <a id="geography"></a>

# geography

The type representing spatial features with geodetic (ellipsoidal) coordinate systems.

## Description


`geography` is a spatial data type used to represent a feature in geodetic coordinate systems. Geodetic coordinate systems model the earth using an ellipsoid.


 Spatial operations on the geography type provide more accurate results by taking the ellipsoidal model into account.


## Casting Behavior


This table lists the automatic and explicit casts allowed for this data type:


| Cast To | Behavior |
| geometry | explicit |


## See Also


[Geography Data Type](../data-management/geography-data-type.md#PostGIS_Geography), [PostGIS Geography Support Functions](../postgis-special-functions-index/postgis-geography-support-functions.md#PostGIS_GeographyFunctions)
