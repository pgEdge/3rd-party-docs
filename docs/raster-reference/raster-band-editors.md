<a id="RasterBand_Editors"></a>

## Raster Band Editors
  <a id="RT_ST_SetBandNoDataValue"></a>

# ST_SetBandNoDataValue

Sets the value for the given band that represents no data. Band 1 is assumed if no band is specified. To mark a band as having no nodata value, set the nodata value = NULL.

## Synopsis


```sql
raster ST_SetBandNoDataValue(raster  rast, double precision  nodatavalue)
raster ST_SetBandNoDataValue(raster  rast, integer  band, double precision  nodatavalue, boolean  forcechecking=false)
```


## Description


Sets the value that represents no data for the band. Band 1 is assumed if not specified. This will affect results from [RT_ST_Polygon](raster-processing-raster-to-geometry.md#RT_ST_Polygon), [RT_ST_DumpAsPolygons](raster-processing-raster-to-geometry.md#RT_ST_DumpAsPolygons), and the ST_PixelAs...() functions.


## Examples


```
-- change just first band no data value
UPDATE dummy_rast
    SET rast = ST_SetBandNoDataValue(rast,1, 254)
WHERE rid = 2;

-- change no data band value of bands 1,2,3
UPDATE dummy_rast
    SET rast =
        ST_SetBandNoDataValue(
            ST_SetBandNoDataValue(
                ST_SetBandNoDataValue(
                    rast,1, 254)
                ,2,99),
                3,108)
        WHERE rid = 2;

-- wipe out the nodata value this will ensure all pixels are considered for all processing functions
UPDATE dummy_rast
    SET rast = ST_SetBandNoDataValue(rast,1, NULL)
WHERE rid = 2;

```


## See Also


[RT_ST_BandNoDataValue](raster-band-accessors.md#RT_ST_BandNoDataValue), [RT_ST_NumBands](raster-accessors.md#RT_ST_NumBands)
  <a id="RT_ST_SetBandIsNoData"></a>

# ST_SetBandIsNoData

Sets the isnodata flag of the band to TRUE.

## Synopsis


```sql
raster ST_SetBandIsNoData(raster  rast, integer  band=1)
```


## Description


Sets the isnodata flag for the band to true. Band 1 is assumed if not specified. This function should be called only when the flag is considered dirty. That is, when the result calling [RT_ST_BandIsNoData](raster-band-accessors.md#RT_ST_BandIsNoData) is different using TRUE as last argument and without using it


Availability: 2.0.0


## Examples


```

-- Create dummy table with one raster column
create table dummy_rast (rid integer, rast raster);

-- Add raster with two bands, one pixel/band. In the first band, nodatavalue = pixel value = 3.
-- In the second band, nodatavalue = 13, pixel value = 4
insert into dummy_rast values(1,
(
'01' -- little endian (uint8 ndr)
||
'0000' -- version (uint16 0)
||
'0200' -- nBands (uint16 0)
||
'17263529ED684A3F' -- scaleX (float64 0.000805965234044584)
||
'F9253529ED684ABF' -- scaleY (float64 -0.00080596523404458)
||
'1C9F33CE69E352C0' -- ipX (float64 -75.5533328537098)
||
'718F0E9A27A44840' -- ipY (float64 49.2824585505576)
||
'ED50EB853EC32B3F' -- skewX (float64 0.000211812383858707)
||
'7550EB853EC32B3F' -- skewY (float64 0.000211812383858704)
||
'E6100000' -- SRID (int32 4326)
||
'0100' -- width (uint16 1)
||
'0100' -- height (uint16 1)
||
'4' -- hasnodatavalue set to true, isnodata value set to false (when it should be true)
||
'2' -- first band type (4BUI)
||
'03' -- novalue==3
||
'03' -- pixel(0,0)==3 (same that nodata)
||
'0' -- hasnodatavalue set to false
||
'5' -- second band type (16BSI)
||
'0D00' -- novalue==13
||
'0400' -- pixel(0,0)==4
)::raster
);

select st_bandisnodata(rast, 1) from dummy_rast where rid = 1; -- Expected false
select st_bandisnodata(rast, 1, TRUE) from dummy_rast where rid = 1; -- Expected true

-- The isnodata flag is dirty. We are going to set it to true
update dummy_rast set rast = st_setbandisnodata(rast, 1) where rid = 1;


select st_bandisnodata(rast, 1) from dummy_rast where rid = 1; -- Expected true


```


## See Also


[RT_ST_BandNoDataValue](raster-band-accessors.md#RT_ST_BandNoDataValue), [RT_ST_NumBands](raster-accessors.md#RT_ST_NumBands), [RT_ST_SetBandNoDataValue](#RT_ST_SetBandNoDataValue), [RT_ST_BandIsNoData](raster-band-accessors.md#RT_ST_BandIsNoData)
  <a id="RT_ST_SetBandPath"></a>

# ST_SetBandPath

Update the external path and band number of an out-db band

## Synopsis


```sql
raster ST_SetBandPath(raster  rast, integer  band, text  outdbpath, integer  outdbindex, boolean  force=false)
```


## Description


Updates an out-db band's external raster file path and external band number.


!!! note

    If `force` is set to true, no tests are done to ensure compatibility (e.g. alignment, pixel support) between the external raster file and the PostGIS raster. This mode is intended for file system changes where the external raster resides.


Availability: 2.5.0


## Examples


```sql

WITH foo AS (
    SELECT
        ST_AddBand(NULL::raster, '/home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif', NULL::int[]) AS rast
)
SELECT
    1 AS query,
    *
FROM ST_BandMetadata(
    (SELECT rast FROM foo),
    ARRAY[1,3,2]::int[]
)
UNION ALL
SELECT
    2,
    *
FROM ST_BandMetadata(
    (
        SELECT
            ST_SetBandPath(
                rast,
                2,
                '/home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected2.tif',
                1
            ) AS rast
        FROM foo
    ),
    ARRAY[1,3,2]::int[]
)
ORDER BY 1, 2;

 query | bandnum | pixeltype | nodatavalue | isoutdb |                                      path                                       | outdbbandnum
-------+---------+-----------+-------------+---------+---------------------------------------------------------------------------------+--------------
     1 |       1 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            1
     1 |       2 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            2
     1 |       3 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            3
     2 |       1 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            1
     2 |       2 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected2.tif      |            1
     2 |       3 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            3

```


## See Also


 [RT_ST_BandMetaData](raster-band-accessors.md#RT_ST_BandMetaData), [RT_ST_SetBandIndex](#RT_ST_SetBandIndex)
  <a id="RT_ST_SetBandIndex"></a>

# ST_SetBandIndex

Update the external band number of an out-db band

## Synopsis


```sql
raster ST_SetBandIndex(raster  rast, integer  band, integer  outdbindex, boolean  force=false)
```


## Description


Updates an out-db band's external band number. This does not touch the external raster file associated with the out-db band


!!! note

    If `force` is set to true, no tests are done to ensure compatibility (e.g. alignment, pixel support) between the external raster file and the PostGIS raster. This mode is intended for where bands are moved around in the external raster file.


!!! note

    Internally, this method replaces the PostGIS raster's band at index `band` with a new band instead of updating the existing path information.


Availability: 2.5.0


## Examples


```sql

WITH foo AS (
    SELECT
        ST_AddBand(NULL::raster, '/home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif', NULL::int[]) AS rast
)
SELECT
    1 AS query,
    *
FROM ST_BandMetadata(
    (SELECT rast FROM foo),
    ARRAY[1,3,2]::int[]
)
UNION ALL
SELECT
    2,
    *
FROM ST_BandMetadata(
    (
        SELECT
            ST_SetBandIndex(
                rast,
                2,
                1
            ) AS rast
        FROM foo
    ),
    ARRAY[1,3,2]::int[]
)
ORDER BY 1, 2;

 query | bandnum | pixeltype | nodatavalue | isoutdb |                                      path                                       | outdbbandnum
-------+---------+-----------+-------------+---------+---------------------------------------------------------------------------------+--------------
     1 |       1 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            1
     1 |       2 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            2
     1 |       3 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            3
     2 |       1 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            1
     2 |       2 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            1
     2 |       3 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif       |            3

```


## See Also


 [RT_ST_BandMetaData](raster-band-accessors.md#RT_ST_BandMetaData), [RT_ST_SetBandPath](#RT_ST_SetBandPath)
