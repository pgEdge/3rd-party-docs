<a id="SRS_Functions"></a>

## Spatial Reference System Functions
  <a id="ST_InverseTransformPipeline"></a>

# ST_InverseTransformPipeline

Return a new geometry with coordinates transformed to a different spatial reference system using the inverse of a defined coordinate transformation pipeline.

## Synopsis


```sql
geometry ST_InverseTransformPipeline(geometry  geom, text  pipeline, integer  to_srid)
```


## Description


 Return a new geometry with coordinates transformed to a different spatial reference system using a defined coordinate transformation pipeline to go in the inverse direction.


 Refer to [ST_TransformPipeline](#ST_TransformPipeline) for details on writing a transformation pipeline.


Availability: 3.4.0


 The SRID of the input geometry is ignored, and the SRID of the output geometry will be set to zero unless a value is provided via the optional `to_srid` parameter. When using [ST_TransformPipeline](#ST_TransformPipeline) the pipeline is executed in a forward direction. Using `ST_InverseTransformPipeline()` the pipeline is executed in the inverse direction.


Transforms using pipelines are a specialised version of [ST_Transform](#ST_Transform). In most cases `ST_Transform` will choose the correct operations to convert between coordinate systems, and should be preferred.


## Examples


Change WGS 84 long lat to UTM 31N using the EPSG:16031 conversion


```

-- Inverse direction
SELECT ST_AsText(ST_InverseTransformPipeline('POINT(426857.9877165967 5427937.523342293)'::geometry,
  'urn:ogc:def:coordinateOperation:EPSG::16031')) AS wgs_geom;

          wgs_geom
----------------------------
 POINT(2 48.99999999999999)
(1 row)

```


GDA2020 example.


```

-- using ST_Transform with automatic selection of a conversion pipeline.
SELECT ST_AsText(ST_Transform('SRID=4939;POINT(143.0 -37.0)'::geometry, 7844)) AS gda2020_auto;

                 gda2020_auto
-----------------------------------------------
 POINT(143.00000635638918 -36.999986706128176)
(1 row)

```


## See Also


[ST_Transform](#ST_Transform), [ST_TransformPipeline](#ST_TransformPipeline)
  <a id="ST_SetSRID"></a>

# ST_SetSRID

Set the SRID on a geometry.

## Synopsis


```sql
geometry ST_SetSRID(geometry
      geom, integer
      srid)
```


## Description


Sets the SRID on a geometry to a particular integer value. Useful in constructing bounding boxes for queries.


!!! note

    This function does not transform the geometry coordinates in any way - it simply sets the meta data defining the spatial reference system the geometry is assumed to be in. Use [ST_Transform](#ST_Transform) if you want to transform the geometry into a new projection.


## Examples


-- Mark a point as WGS 84 long lat --


```sql
SELECT ST_SetSRID(ST_Point(-123.365556, 48.428611),4326) As wgs84long_lat;
-- the ewkt representation (wrap with ST_AsEWKT) -
SRID=4326;POINT(-123.365556 48.428611)

```


-- Mark a point as WGS 84 long lat and then transform to web mercator (Spherical Mercator) --


```sql
SELECT ST_Transform(ST_SetSRID(ST_Point(-123.365556, 48.428611),4326),3785) As spere_merc;
-- the ewkt representation (wrap with ST_AsEWKT) -
SRID=3785;POINT(-13732990.8753491 6178458.96425423)

```


## See Also


[Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys), [ST_SRID](#ST_SRID), [ST_Transform](#ST_Transform), [UpdateGeometrySRID](table-management-functions.md#UpdateGeometrySRID)
  <a id="ST_SRID"></a>

# ST_SRID

Returns the spatial reference identifier for a geometry.

## Synopsis


```sql
integer ST_SRID(geometry  g1)
```


## Description


Returns the spatial reference identifier for the ST_Geometry as defined in spatial_ref_sys table. [Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys)


!!! note

    spatial_ref_sys table is a table that catalogs all spatial reference systems known to PostGIS and is used for transformations from one spatial reference system to another. So verifying you have the right spatial reference system identifier is important if you plan to ever transform your geometries.


 s2.1.1.1


 SQL-MM 3: 5.1.5


## Examples


```sql
SELECT ST_SRID(ST_GeomFromText('POINT(-71.1043 42.315)',4326));
    --result
    4326

```


## See Also


[Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys), [ST_SetSRID](#ST_SetSRID), [ST_Transform](#ST_Transform), [RT_ST_SRID](../raster-reference/raster-accessors.md#RT_ST_SRID), [TG_ST_SRID](../topology/topogeometry-accessors.md#TG_ST_SRID)
  <a id="ST_Transform"></a>

# ST_Transform

Return a new geometry with coordinates transformed to a different spatial reference system.

## Synopsis


```sql
geometry ST_Transform(geometry  g1, integer  srid)
geometry ST_Transform(geometry  geom, text  to_proj)
geometry ST_Transform(geometry  geom, text  from_proj, text  to_proj)
geometry ST_Transform(geometry  geom, text  from_proj, integer  to_srid)
```


## Description


Returns a new geometry with its coordinates transformed to a different spatial reference system. The destination spatial reference `to_srid` may be identified by a valid SRID integer parameter (i.e. it must exist in the `spatial_ref_sys` table). Alternatively, a spatial reference defined as a PROJ.4 string can be used for `to_proj` and/or `from_proj`, however these methods are not optimized. If the destination spatial reference system is expressed with a PROJ.4 string instead of an SRID, the SRID of the output geometry will be set to zero. With the exception of functions with `from_proj`, input geometries must have a defined SRID.


ST_Transform is often confused with [ST_SetSRID](#ST_SetSRID). ST_Transform actually changes the coordinates of a geometry from one spatial reference system to another, while ST_SetSRID() simply changes the SRID identifier of the geometry.


ST_Transform automatically selects a suitable conversion pipeline given the source and target spatial reference systems. To use a specific conversion method, use [ST_TransformPipeline](#ST_TransformPipeline).


!!! note

    Requires PostGIS be compiled with PROJ support. Use [PostGIS_Full_Version](version-functions.md#PostGIS_Full_Version) to confirm you have PROJ support compiled in.


!!! note

    If using more than one transformation, it is useful to have a functional index on the commonly used transformations to take advantage of index usage.


!!! note

    Prior to 1.3.4, this function crashes if used with geometries that contain CURVES. This is fixed in 1.3.4+


Enhanced: 2.0.0 support for Polyhedral surfaces was introduced.


Enhanced: 2.3.0 support for direct PROJ.4 text was introduced.


 SQL-MM 3: 5.1.6


## Examples


Change Massachusetts state plane US feet geometry to WGS 84 long lat


```sql

SELECT ST_AsText(ST_Transform(ST_GeomFromText('POLYGON((743238 2967416,743238 2967450,
  743265 2967450,743265.625 2967416,743238 2967416))',2249),4326)) As wgs_geom;

 wgs_geom
---------------------------
 POLYGON((-71.1776848522251 42.3902896512902,-71.1776843766326 42.3903829478009,
-71.1775844305465 42.3903826677917,-71.1775825927231 42.3902893647987,-71.177684
8522251 42.3902896512902));
(1 row)

--3D Circular String example
SELECT ST_AsEWKT(ST_Transform(ST_GeomFromEWKT('SRID=2249;CIRCULARSTRING(743238 2967416 1,743238 2967450 2,743265 2967450 3,743265.625 2967416 3,743238 2967416 4)'),4326));

         st_asewkt
--------------------------------------------------------------------------------------
 SRID=4326;CIRCULARSTRING(-71.1776848522251 42.3902896512902 1,-71.1776843766326 42.3903829478009 2,
 -71.1775844305465 42.3903826677917 3,
 -71.1775825927231 42.3902893647987 3,-71.1776848522251 42.3902896512902 4)


```


Example of creating a partial functional index. For tables where you are not sure all the geometries will be filled in, its best to use a partial index that leaves out null geometries which will both conserve space and make your index smaller and more efficient.


```sql

CREATE INDEX idx_geom_26986_parcels
  ON parcels
  USING gist
  (ST_Transform(geom, 26986))
  WHERE geom IS NOT NULL;

```


Examples of using PROJ.4 text to transform with custom spatial references.


```

-- Find intersection of two polygons near the North pole, using a custom Gnomic projection
-- See http://boundlessgeo.com/2012/02/flattening-the-peel/
 WITH data AS (
   SELECT
     ST_GeomFromText('POLYGON((170 50,170 72,-130 72,-130 50,170 50))', 4326) AS p1,
     ST_GeomFromText('POLYGON((-170 68,-170 90,-141 90,-141 68,-170 68))', 4326) AS p2,
     '+proj=gnom +ellps=WGS84 +lat_0=70 +lon_0=-160 +no_defs'::text AS gnom
 )
 SELECT ST_AsText(
   ST_Transform(
     ST_Intersection(ST_Transform(p1, gnom), ST_Transform(p2, gnom)),
   gnom, 4326))
 FROM data;
                                          st_astext
 --------------------------------------------------------------------------------
  POLYGON((-170 74.053793645338,-141 73.4268621378904,-141 68,-170 68,-170 74.053793645338))

```


## Configuring transformation behavior


Sometimes coordinate transformation involving a grid-shift can fail, for example if PROJ.4 has not been built with grid-shift files or the coordinate does not lie within the range for which the grid shift is defined. By default, PostGIS will throw an error if a grid shift file is not present, but this behavior can be configured on a per-SRID basis either by testing different `to_proj` values of PROJ.4 text, or altering the `proj4text` value within the `spatial_ref_sys` table.


For example, the proj4text parameter +datum=NAD87 is a shorthand form for the following +nadgrids parameter:


```
+nadgrids=@conus,@alaska,@ntv2_0.gsb,@ntv1_can.dat
```


The @ prefix means no error is reported if the files are not present, but if the end of the list is reached with no file having been appropriate (ie. found and overlapping) then an error is issued.


If, conversely, you wanted to ensure that at least the standard files were present, but that if all files were scanned without a hit a null transformation is applied you could use:


```
+nadgrids=@conus,@alaska,@ntv2_0.gsb,@ntv1_can.dat,null
```


The null grid shift file is a valid grid shift file covering the whole world and applying no shift. So for a complete example, if you wanted to alter PostGIS so that transformations to SRID 4267 that didn't lie within the correct range did not throw an ERROR, you would use the following:


```sql
UPDATE spatial_ref_sys SET proj4text = '+proj=longlat +ellps=clrk66 +nadgrids=@conus,@alaska,@ntv2_0.gsb,@ntv1_can.dat,null +no_defs' WHERE srid = 4267;
```


## See Also


[Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys), [ST_SetSRID](#ST_SetSRID), [ST_SRID](#ST_SRID), [UpdateGeometrySRID](table-management-functions.md#UpdateGeometrySRID), [ST_TransformPipeline](#ST_TransformPipeline)
  <a id="ST_TransformPipeline"></a>

# ST_TransformPipeline

Return a new geometry with coordinates transformed to a different spatial reference system using a defined coordinate transformation pipeline.

## Synopsis


```sql
geometry ST_TransformPipeline(geometry  g1, text  pipeline, integer  to_srid)
```


## Description


 Return a new geometry with coordinates transformed to a different spatial reference system using a defined coordinate transformation pipeline.


 Transformation pipelines are defined using any of the following string formats:

-  `urn:ogc:def:coordinateOperation:AUTHORITY::CODE`. Note that a simple `EPSG:CODE` string does not uniquely identify a coordinate operation: the same EPSG code can be used for a CRS definition.
-  A PROJ pipeline string of the form: `+proj=pipeline ...`. Automatic axis normalisation will not be applied, and if necessary the caller will need to add an additional pipeline step, or remove `axisswap` steps.
-  Concatenated operations of the form: `urn:ogc:def:coordinateOperation,coordinateOperation:EPSG::3895,coordinateOperation:EPSG::1618`.


Availability: 3.4.0


 The SRID of the input geometry is ignored, and the SRID of the output geometry will be set to zero unless a value is provided via the optional `to_srid` parameter. When using `ST_TransformPipeline()` the pipeline is executed in a forward direction. Using [ST_InverseTransformPipeline](#ST_InverseTransformPipeline) the pipeline is executed in the inverse direction.


Transforms using pipelines are a specialised version of [ST_Transform](#ST_Transform). In most cases `ST_Transform` will choose the correct operations to convert between coordinate systems, and should be preferred.


## Examples


Change WGS 84 long lat to UTM 31N using the EPSG:16031 conversion


```

-- Forward direction
SELECT ST_AsText(ST_TransformPipeline('SRID=4326;POINT(2 49)'::geometry,
  'urn:ogc:def:coordinateOperation:EPSG::16031')) AS utm_geom;

                  utm_geom
--------------------------------------------
 POINT(426857.9877165967 5427937.523342293)
(1 row)

-- Inverse direction
SELECT ST_AsText(ST_InverseTransformPipeline('POINT(426857.9877165967 5427937.523342293)'::geometry,
  'urn:ogc:def:coordinateOperation:EPSG::16031')) AS wgs_geom;

          wgs_geom
----------------------------
 POINT(2 48.99999999999999)
(1 row)

```


GDA2020 example.


```

-- using ST_Transform with automatic selection of a conversion pipeline.
SELECT ST_AsText(ST_Transform('SRID=4939;POINT(143.0 -37.0)'::geometry, 7844)) AS gda2020_auto;

                 gda2020_auto
-----------------------------------------------
 POINT(143.00000635638918 -36.999986706128176)
(1 row)

-- using a defined conversion (EPSG:8447)
SELECT ST_AsText(ST_TransformPipeline('SRID=4939;POINT(143.0 -37.0)'::geometry,
  'urn:ogc:def:coordinateOperation:EPSG::8447')) AS gda2020_code;

                   gda2020_code
----------------------------------------------
 POINT(143.0000063280214 -36.999986718287545)
(1 row)

-- using a PROJ pipeline definition matching EPSG:8447, as returned from
-- 'projinfo -s EPSG:4939 -t EPSG:7844'.
-- NOTE: any 'axisswap' steps must be removed.
SELECT ST_AsText(ST_TransformPipeline('SRID=4939;POINT(143.0 -37.0)'::geometry,
  '+proj=pipeline
   +step +proj=unitconvert +xy_in=deg +xy_out=rad
   +step +proj=hgridshift +grids=au_icsm_GDA94_GDA2020_conformal_and_distortion.tif
   +step +proj=unitconvert +xy_in=rad +xy_out=deg')) AS gda2020_pipeline;

                   gda2020_pipeline
----------------------------------------------
 POINT(143.0000063280214 -36.999986718287545)
(1 row)

```


## See Also


[ST_Transform](#ST_Transform), [ST_InverseTransformPipeline](#ST_InverseTransformPipeline)
  <a id="postgis_srs_codes"></a>

# postgis_srs_codes

Return the list of SRS codes associated with the given authority.

## Synopsis


```sql
setof text postgis_srs_codes(text  auth_name)
```


## Description


Returns a set of all `auth_srid` for the given `auth_name`.


Availability: 3.4.0


Proj version 6+


## Examples


List the first ten codes associated with the EPSG authority.


```sql

SELECT * FROM postgis_srs_codes('EPSG') LIMIT 10;

 postgis_srs_codes
-------------------
 2000
 20004
 20005
 20006
 20007
 20008
 20009
 2001
 20010
 20011

```


## See Also


[postgis_srs](#postgis_srs), [postgis_srs_all](#postgis_srs_all), [postgis_srs_search](#postgis_srs_search)
  <a id="postgis_srs"></a>

# postgis_srs

Return a metadata record for the requested authority and srid.

## Synopsis


```sql
setof record postgis_srs(text  auth_name, text  auth_srid)
```


## Description


Returns a metadata record for the requested `auth_srid` for the given `auth_name`. The record will have the `auth_name`, `auth_srid`, `srname`, `srtext`, `proj4text`, and the corners of the area of usage, `point_sw` and `point_ne`.


Availability: 3.4.0


Proj version 6+


## Examples


Get the metadata for EPSG:3005.


```sql

SELECT * FROM postgis_srs('EPSG', '3005');

auth_name | EPSG
auth_srid | 3005
srname    | NAD83 / BC Albers
srtext    | PROJCS["NAD83 / BC Albers", ... ]]
proj4text | +proj=aea +lat_0=45 +lon_0=-126 +lat_1=50 +lat_2=58.5 +x_0=1000000 +y_0=0 +datum=NAD83 +units=m +no_defs +type=crs
point_sw  | 0101000020E6100000E17A14AE476161C00000000000204840
point_ne  | 0101000020E610000085EB51B81E855CC0E17A14AE47014E40

```


## See Also


[postgis_srs_codes](#postgis_srs_codes), [postgis_srs_all](#postgis_srs_all), [postgis_srs_search](#postgis_srs_search)
  <a id="postgis_srs_all"></a>

# postgis_srs_all

Return metadata records for every spatial reference system in the underlying Proj database.

## Synopsis


```sql
setof record postgis_srs_all()
```


## Description


Returns a set of all metadata records in the underlying Proj database. The records will have the `auth_name`, `auth_srid`, `srname`, `srtext`, `proj4text`, and the corners of the area of usage, `point_sw` and `point_ne`.


Availability: 3.4.0


Proj version 6+


## Examples


Get the first 10 metadata records from the Proj database.


```sql

SELECT auth_name, auth_srid, srname FROM postgis_srs_all() LIMIT 10;

 auth_name | auth_srid |                  srname
-----------+-----------+------------------------------------------
 EPSG      | 2000      | Anguilla 1957 / British West Indies Grid
 EPSG      | 20004     | Pulkovo 1995 / Gauss-Kruger zone 4
 EPSG      | 20005     | Pulkovo 1995 / Gauss-Kruger zone 5
 EPSG      | 20006     | Pulkovo 1995 / Gauss-Kruger zone 6
 EPSG      | 20007     | Pulkovo 1995 / Gauss-Kruger zone 7
 EPSG      | 20008     | Pulkovo 1995 / Gauss-Kruger zone 8
 EPSG      | 20009     | Pulkovo 1995 / Gauss-Kruger zone 9
 EPSG      | 2001      | Antigua 1943 / British West Indies Grid
 EPSG      | 20010     | Pulkovo 1995 / Gauss-Kruger zone 10
 EPSG      | 20011     | Pulkovo 1995 / Gauss-Kruger zone 11
```


## See Also


[postgis_srs_codes](#postgis_srs_codes), [postgis_srs](#postgis_srs), [postgis_srs_search](#postgis_srs_search)
  <a id="postgis_srs_search"></a>

# postgis_srs_search

Return metadata records for projected coordinate systems that have areas of usage that fully contain the bounds parameter.

## Synopsis


```sql
setof record postgis_srs_search(geometry  bounds, text  auth_name=EPSG)
```


## Description


Return a set of metadata records for projected coordinate systems that have areas of usage that fully contain the bounds parameter. Each record will have the `auth_name`, `auth_srid`, `srname`, `srtext`, `proj4text`, and the corners of the area of usage, `point_sw` and `point_ne`.


The search only looks for projected coordinate systems, and is intended for users to explore the possible systems that work for the extent of their data.


Availability: 3.4.0


Proj version 6+


## Examples


Search for projected coordinate systems in Louisiana.


```sql

SELECT auth_name, auth_srid, srname,
  ST_AsText(point_sw) AS point_sw,
  ST_AsText(point_ne) AS point_ne
FROM postgis_srs_search('SRID=4326;LINESTRING(-90 30, -91 31)')
LIMIT 3;

 auth_name | auth_srid |                srname                |      point_sw       |      point_ne
-----------+-----------+--------------------------------------+---------------------+---------------------
 EPSG      | 2801      | NAD83(HARN) / Louisiana South        | POINT(-93.94 28.85) | POINT(-88.75 31.07)
 EPSG      | 3452      | NAD83 / Louisiana South (ftUS)       | POINT(-93.94 28.85) | POINT(-88.75 31.07)
 EPSG      | 3457      | NAD83(HARN) / Louisiana South (ftUS) | POINT(-93.94 28.85) | POINT(-88.75 31.07)
```


Scan a table for max extent and find projected coordinate systems that might suit.


```sql

WITH ext AS (
  SELECT ST_Extent(geom) AS geom, Max(ST_SRID(geom)) AS srid
  FROM foo
)
SELECT auth_name, auth_srid, srname,
  ST_AsText(point_sw) AS point_sw,
  ST_AsText(point_ne) AS point_ne
FROM ext
CROSS JOIN postgis_srs_search(ST_SetSRID(ext.geom, ext.srid))
LIMIT 3;
```


## See Also


[postgis_srs_codes](#postgis_srs_codes), [postgis_srs_all](#postgis_srs_all), [postgis_srs](#postgis_srs)
