<a id="Raster_Processing_Geometry"></a>

## Raster Processing: Raster to Geometry
  <a id="RT_Box3D"></a>

# Box3D

Returns the box 3d representation of the enclosing box of the raster.

## Synopsis


```sql
box3d Box3D(raster  rast)
```


## Description


Returns the box representing the extent of the raster.


 The polygon is defined by the corner points of the bounding box ((`MINX`, `MINY`), (`MAXX`, `MAXY`))


Changed: 2.0.0 In pre-2.0 versions, there used to be a box2d instead of box3d. Since box2d is a deprecated type, this was changed to box3d.


## Examples


```sql

SELECT
    rid,
    Box3D(rast) AS rastbox
FROM dummy_rast;

rid |        rastbox
----+-------------------------------------------------
1   | BOX3D(0.5 0.5 0,20.5 60.5 0)
2   | BOX3D(3427927.75 5793243.5 0,3427928 5793244 0)

```


## See Also


 [RT_ST_Envelope](#RT_ST_Envelope)
  <a id="RT_ST_ConvexHull"></a>

# ST_ConvexHull

Return the convex hull geometry of the raster including pixel values equal to BandNoDataValue. For regular shaped and non-skewed rasters, this gives the same result as ST_Envelope so only useful for irregularly shaped or skewed rasters.

## Synopsis


```sql
geometry ST_ConvexHull(raster  rast)
```


## Description


Return the convex hull geometry of the raster including the NoDataBandValue band pixels. For regular shaped and non-skewed rasters, this gives more or less the same result as ST_Envelope so only useful for irregularly shaped or skewed rasters.


!!! note

    ST_Envelope floors the coordinates and hence add a little buffer around the raster so the answer is subtly different from ST_ConvexHull which does not floor.


## Examples


Refer to [PostGIS Raster Specification](http://trac.osgeo.org/postgis/wiki/WKTRaster/SpecificationWorking01) for a diagram of this.


```

-- Note envelope and convexhull are more or less the same
SELECT ST_AsText(ST_ConvexHull(rast)) As convhull,
    ST_AsText(ST_Envelope(rast)) As env
FROM dummy_rast WHERE rid=1;

                        convhull                        |                env
--------------------------------------------------------+------------------------------------
 POLYGON((0.5 0.5,20.5 0.5,20.5 60.5,0.5 60.5,0.5 0.5)) | POLYGON((0 0,20 0,20 60,0 60,0 0))

```


```

-- now we skew the raster
-- note how the convex hull and envelope are now different
SELECT ST_AsText(ST_ConvexHull(rast)) As convhull,
    ST_AsText(ST_Envelope(rast)) As env
FROM (SELECT ST_SetRotation(rast, 0.1, 0.1) As rast
    FROM dummy_rast WHERE rid=1) As foo;

                        convhull                        |                env
--------------------------------------------------------+------------------------------------
 POLYGON((0.5 0.5,20.5 1.5,22.5 61.5,2.5 60.5,0.5 0.5)) | POLYGON((0 0,22 0,22 61,0 61,0 0))

```


## See Also


 [RT_ST_Envelope](#RT_ST_Envelope), [RT_ST_MinConvexHull](#RT_ST_MinConvexHull), [ST_ConvexHull](../postgis-reference/geometry-processing.md#ST_ConvexHull), [ST_AsText](../postgis-reference/geometry-output.md#ST_AsText)
  <a id="RT_ST_DumpAsPolygons"></a>

# ST_DumpAsPolygons

Returns a set of geomval (geom,val) rows, from a given raster band. If no band number is specified, band num defaults to 1.

## Synopsis


```sql
setof geomval ST_DumpAsPolygons(raster  rast, integer  band_num=1, boolean  exclude_nodata_value=TRUE)
```


## Description


This is a set-returning function (SRF). It returns a set of geomval rows, formed by a geometry (geom) and a pixel band value (val). Each polygon is the union of all pixels for that band that have the same pixel value denoted by val.


ST_DumpAsPolygon is useful for polygonizing rasters. It is the reverse of a GROUP BY in that it creates new rows. For example it can be used to expand a single raster into multiple POLYGONS/MULTIPOLYGONS.


Changed 3.3.0, validation and fixing is disabled to improve performance. May result invalid geometries.


Availability: Requires GDAL 1.7 or higher.


!!! note

    If there is a no data value set for a band, pixels with that value will not be returned except in the case of exclude_nodata_value=false.


!!! note

    If you only care about count of pixels with a given value in a raster, it is faster to use [RT_ST_ValueCount](raster-band-statistics-and-analytics.md#RT_ST_ValueCount).


!!! note

    This is different than ST_PixelAsPolygons where one geometry is returned for each pixel regardless of pixel value.


## Examples


```
 -- this syntax requires PostgreSQL 9.3+
SELECT val, ST_AsText(geom) As geomwkt
FROM (
SELECT dp.*
FROM dummy_rast, LATERAL ST_DumpAsPolygons(rast) AS dp
WHERE rid = 2
) As foo
WHERE val BETWEEN 249 and 251
ORDER BY val;

 val |                                                       geomwkt
-----+--------------------------------------------------------------------------
 249 | POLYGON((3427927.95 5793243.95,3427927.95 5793243.85,3427928 5793243.85,
        3427928 5793243.95,3427927.95 5793243.95))
 250 | POLYGON((3427927.75 5793243.9,3427927.75 5793243.85,3427927.8 5793243.85,
        3427927.8 5793243.9,3427927.75 5793243.9))
 250 | POLYGON((3427927.8 5793243.8,3427927.8 5793243.75,3427927.85 5793243.75,
        3427927.85 5793243.8, 3427927.8 5793243.8))
 251 | POLYGON((3427927.75 5793243.85,3427927.75 5793243.8,3427927.8 5793243.8,
        3427927.8 5793243.85,3427927.75 5793243.85))

```


## See Also


 [geomval](raster-support-data-types.md#geomval), [RT_ST_Value](raster-pixel-accessors-and-setters.md#RT_ST_Value), [RT_ST_Polygon](#RT_ST_Polygon), [RT_ST_ValueCount](raster-band-statistics-and-analytics.md#RT_ST_ValueCount)
  <a id="RT_ST_Envelope"></a>

# ST_Envelope

Returns the polygon representation of the extent of the raster.

## Synopsis


```sql
geometry ST_Envelope(raster  rast)
```


## Description


Returns the polygon representation of the extent of the raster in spatial coordinate units defined by srid. It is a float8 minimum bounding box represented as a polygon.


The polygon is defined by the corner points of the bounding box ((`MINX`, `MINY`), (`MINX`, `MAXY`), (`MAXX`, `MAXY`), (`MAXX`, `MINY`), (`MINX`, `MINY`))


## Examples


```sql

SELECT rid, ST_AsText(ST_Envelope(rast)) As envgeomwkt
FROM dummy_rast;

 rid |                                         envgeomwkt
-----+--------------------------------------------------------------------
   1 | POLYGON((0 0,20 0,20 60,0 60,0 0))
   2 | POLYGON((3427927 5793243,3427928 5793243,
        3427928 5793244,3427927 5793244, 3427927 5793243))

```


## See Also


 [ST_Envelope](../postgis-reference/geometry-accessors.md#ST_Envelope), [ST_AsText](../postgis-reference/geometry-output.md#ST_AsText), [RT_ST_SRID](raster-accessors.md#RT_ST_SRID)
  <a id="RT_ST_MinConvexHull"></a>

# ST_MinConvexHull

Return the convex hull geometry of the raster excluding NODATA pixels.

## Synopsis


```sql
geometry ST_MinConvexHull(raster  rast, integer  nband=NULL)
```


## Description


 Return the convex hull geometry of the raster excluding NODATA pixels. If `nband` is NULL, all bands of the raster are considered.


Availability: 2.1.0


## Examples


```sql

WITH foo AS (
    SELECT
        ST_SetValues(
            ST_SetValues(
                ST_AddBand(ST_AddBand(ST_MakeEmptyRaster(9, 9, 0, 0, 1, -1, 0, 0, 0), 1, '8BUI', 0, 0), 2, '8BUI', 1, 0),
                1, 1, 1,
                ARRAY[
                    [0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 1, 0, 0, 0, 0, 1],
                    [0, 0, 0, 1, 1, 0, 0, 0, 0],
                    [0, 0, 0, 1, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0]
                ]::double precision[][]
            ),
            2, 1, 1,
            ARRAY[
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [1, 0, 0, 0, 0, 1, 0, 0, 0],
                [0, 0, 0, 0, 1, 1, 0, 0, 0],
                [0, 0, 0, 0, 0, 1, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 1, 0, 0, 0, 0, 0, 0]
            ]::double precision[][]
        ) AS rast
)
SELECT
    ST_AsText(ST_ConvexHull(rast)) AS hull,
    ST_AsText(ST_MinConvexHull(rast)) AS mhull,
    ST_AsText(ST_MinConvexHull(rast, 1)) AS mhull_1,
    ST_AsText(ST_MinConvexHull(rast, 2)) AS mhull_2
FROM foo

               hull               |                mhull                |               mhull_1               |               mhull_2
----------------------------------+-------------------------------------+-------------------------------------+-------------------------------------
 POLYGON((0 0,9 0,9 -9,0 -9,0 0)) | POLYGON((0 -3,9 -3,9 -9,0 -9,0 -3)) | POLYGON((3 -3,9 -3,9 -6,3 -6,3 -3)) | POLYGON((0 -3,6 -3,6 -9,0 -9,0 -3))

```


## See Also


 [RT_ST_Envelope](#RT_ST_Envelope), [RT_ST_ConvexHull](#RT_ST_ConvexHull), [ST_ConvexHull](../postgis-reference/geometry-processing.md#ST_ConvexHull), [ST_AsText](../postgis-reference/geometry-output.md#ST_AsText)
  <a id="RT_ST_Polygon"></a>

# ST_Polygon

Returns a multipolygon geometry formed by the union of pixels that have a pixel value that is not no data value. If no band number is specified, band num defaults to 1.

## Synopsis


```sql
geometry ST_Polygon(raster  rast, integer  band_num=1)
```


## Description


Changed 3.3.0, validation and fixing is disabled to improve performance. May result invalid geometries.


Availability: 0.1.6 Requires GDAL 1.7 or higher.


Enhanced: 2.1.0 Improved Speed (fully C-Based) and the returning multipolygon is ensured to be valid.


Changed: 2.1.0 In prior versions would sometimes return a polygon, changed to always return multipolygon.


## Examples


```

-- by default no data band value is 0 or not set, so polygon will return a square polygon
SELECT ST_AsText(ST_Polygon(rast)) As geomwkt
FROM dummy_rast
WHERE rid = 2;

geomwkt
--------------------------------------------
MULTIPOLYGON(((3427927.75 5793244,3427928 5793244,3427928 5793243.75,3427927.75 5793243.75,3427927.75 5793244)))


-- now we change the no data value of first band
UPDATE dummy_rast SET rast = ST_SetBandNoDataValue(rast,1,254)
WHERE rid = 2;
SELECt rid, ST_BandNoDataValue(rast)
from dummy_rast where rid = 2;

-- ST_Polygon excludes the pixel value 254 and returns a multipolygon
SELECT ST_AsText(ST_Polygon(rast)) As geomwkt
FROM dummy_rast
WHERE rid = 2;

geomwkt
---------------------------------------------------------
MULTIPOLYGON(((3427927.9 5793243.95,3427927.85 5793243.95,3427927.85 5793244,3427927.9 5793244,3427927.9 5793243.95)),((3427928 5793243.85,3427928 5793243.8,3427927.95 5793243.8,3427927.95 5793243.85,3427927.9 5793243.85,3427927.9 5793243.9,3427927.9 5793243.95,3427927.95 5793243.95,3427928 5793243.95,3427928 5793243.85)),((3427927.8 5793243.75,3427927.75 5793243.75,3427927.75 5793243.8,3427927.75 5793243.85,3427927.75 5793243.9,3427927.75 5793244,3427927.8 5793244,3427927.8 5793243.9,3427927.8 5793243.85,3427927.85 5793243.85,3427927.85 5793243.8,3427927.85 5793243.75,3427927.8 5793243.75)))

-- Or if you want the no data value different for just one time

SELECT ST_AsText(
    ST_Polygon(
        ST_SetBandNoDataValue(rast,1,252)
        )
    ) As geomwkt
FROM dummy_rast
WHERE rid =2;

geomwkt
---------------------------------
MULTIPOLYGON(((3427928 5793243.85,3427928 5793243.8,3427928 5793243.75,3427927.85 5793243.75,3427927.8 5793243.75,3427927.8 5793243.8,3427927.75 5793243.8,3427927.75 5793243.85,3427927.75 5793243.9,3427927.75 5793244,3427927.8 5793244,3427927.85 5793244,3427927.9 5793244,3427928 5793244,3427928 5793243.95,3427928 5793243.85),(3427927.9 5793243.9,3427927.9 5793243.85,3427927.95 5793243.85,3427927.95 5793243.9,3427927.9 5793243.9)))

```


## See Also


 [RT_ST_Value](raster-pixel-accessors-and-setters.md#RT_ST_Value), [RT_ST_DumpAsPolygons](#RT_ST_DumpAsPolygons)
