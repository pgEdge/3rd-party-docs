<a id="Raster_Outputs"></a>

## Raster Outputs
  <a id="RT_ST_AsBinary"></a>

# ST_AsBinary/ST_AsWKB

Return the Well-Known Binary (WKB) representation of the raster.

## Synopsis


```sql
bytea ST_AsBinary(raster  rast, boolean  outasin=FALSE)
bytea ST_AsWKB(raster  rast, boolean  outasin=FALSE)
```


## Description


 Returns the Binary representation of the raster. If `outasin` is TRUE, out-db bands are treated as in-db. Refer to raster/doc/RFC2-WellKnownBinaryFormat located in the PostGIS source folder for details of the representation.


 This is useful in binary cursors to pull data out of the database without converting it to a string representation.


!!! note

    By default, WKB output contains the external file path for out-db bands. If the client does not have access to the raster file underlying an out-db band, set `outasin` to TRUE.


Enhanced: 2.1.0 Addition of `outasin`


Enhanced: 2.5.0 Addition of `ST_AsWKB`


## Examples


```sql

SELECT ST_AsBinary(rast) As rastbin FROM dummy_rast WHERE rid=1;

                     rastbin
---------------------------------------------------------------------------------
\001\000\000\000\000\000\000\000\000\000\000\000@\000\000\000\000\000\000\010@\000\000\000\000\000\000\340?\000\000\000\000\000\000\340?\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\012\000\000\000\012\000\024\000

```


## See Also


 [RT_ST_RastFromWKB](raster-inputs.md#RT_ST_RastFromWKB), [RT_ST_AsHexWKB](#RT_ST_AsHexWKB)
  <a id="RT_ST_AsHexWKB"></a>

# ST_AsHexWKB

Return the Well-Known Binary (WKB) in Hex representation of the raster.

## Synopsis


```sql
bytea ST_AsHexWKB(raster  rast, boolean  outasin=FALSE)
```


## Description


 Returns the Binary representation in Hex representation of the raster. If `outasin` is TRUE, out-db bands are treated as in-db. Refer to raster/doc/RFC2-WellKnownBinaryFormat located in the PostGIS source folder for details of the representation.


!!! note

    By default, Hex WKB output contains the external file path for out-db bands. If the client does not have access to the raster file underlying an out-db band, set `outasin` to TRUE.


Availability: 2.5.0


## Examples


```sql

SELECT ST_AsHexWKB(rast) As rastbin FROM dummy_rast WHERE rid=1;

                                                        st_ashexwkb
----------------------------------------------------------------------------------------------------------------------------
 010000000000000000000000400000000000000840000000000000E03F000000000000E03F000000000000000000000000000000000A0000000A001400

```


## See Also


 [RT_ST_RastFromHexWKB](raster-inputs.md#RT_ST_RastFromHexWKB), [RT_ST_AsBinary](#RT_ST_AsBinary)
  <a id="RT_ST_AsGDALRaster"></a>

# ST_AsGDALRaster

Return the raster tile in the designated GDAL Raster format. Raster formats are one of those supported by your compiled library. Use ST_GDALDrivers() to get a list of formats supported by your library.

## Synopsis


```sql
bytea ST_AsGDALRaster(raster  rast, text  format, text[]  options=NULL, integer  srid=sameassource)
```


## Description


Returns the raster tile in the designated format. Arguments are itemized below:


-  `format` format to output. This is dependent on the drivers compiled in your libgdal library. Generally available are 'JPEG', 'GTIff', 'PNG'. Use [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers) to get a list of formats supported by your library.
-  `options` text array of GDAL options. Valid options are dependent on the format. Refer to [GDAL Raster format options](http://www.gdal.org/frmt_various.html) for more details.
-  `srs` The proj4text or srtext (from spatial_ref_sys) to embed in the image


Availability: 2.0.0 - requires GDAL >= 1.6.0.


## JPEG Output Example, multiple tiles as single raster


```sql

SELECT ST_AsGDALRaster(ST_Union(rast), 'JPEG', ARRAY['QUALITY=50']) As rastjpg
FROM dummy_rast
WHERE rast && ST_MakeEnvelope(10, 10, 11, 11);
```


## Using PostgreSQL Large Object Support to export raster


One way to export raster into another format is using [PostgreSQL large object export functions](https://www.postgresql.org/docs/current/static/lo-funcs.html). We'lll repeat the prior example but also exporting. Note for this you'll need to have super user access to db since it uses server side lo functions. It will also export to path on server network. If you need export locally, use the psql equivalent lo_ functions which export to the local file system instead of the server file system.


```sql

DROP TABLE IF EXISTS tmp_out ;

CREATE TABLE tmp_out AS
SELECT lo_from_bytea(0,
       ST_AsGDALRaster(ST_Union(rast), 'JPEG', ARRAY['QUALITY=50'])
        ) AS loid
  FROM dummy_rast
WHERE rast && ST_MakeEnvelope(10, 10, 11, 11);

SELECT lo_export(loid, '/tmp/dummy.jpg')
   FROM tmp_out;

SELECT lo_unlink(loid)
  FROM tmp_out;
```


## GTIFF Output Examples


```sql
SELECT ST_AsGDALRaster(rast, 'GTiff') As rastjpg
FROM dummy_rast WHERE rid=2;

-- Out GeoTiff with jpeg compression, 90% quality
SELECT ST_AsGDALRaster(rast, 'GTiff',
  ARRAY['COMPRESS=JPEG', 'JPEG_QUALITY=90'],
  4269) As rasttiff
FROM dummy_rast WHERE rid=2;

```


## See Also


[Building Custom Applications with PostGIS Raster](../raster-data-management-queries-and-applications/building-custom-applications-with-postgis-raster.md#RT_Raster_Applications), [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers), [RT_ST_SRID](raster-accessors.md#RT_ST_SRID)
  <a id="RT_ST_AsJPEG"></a>

# ST_AsJPEG

Return the raster tile selected bands as a single Joint Photographic Exports Group (JPEG) image (byte array). If no band is specified and 1 or more than 3 bands, then only the first band is used. If only 3 bands then all 3 bands are used and mapped to RGB.

## Synopsis


```sql
bytea ST_AsJPEG(raster  rast, text[]  options=NULL)
bytea ST_AsJPEG(raster  rast, integer  nband, integer  quality)
bytea ST_AsJPEG(raster  rast, integer  nband, text[]  options=NULL)
bytea ST_AsJPEG(raster  rast, integer[]  nbands, text[]  options=NULL)
bytea ST_AsJPEG(raster  rast, integer[]  nbands, integer  quality)
```


## Description


Returns the selected bands of the raster as a single Joint Photographic Exports Group Image (JPEG). Use [RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster) if you need to export as less common raster types. If no band is specified and 1 or more than 3 bands, then only the first band is used. If 3 bands then all 3 bands are used. There are many variants of the function with many options. These are itemized below:


-  `nband` is for single band exports.
-  `nbands` is an array of bands to export (note that max is 3 for JPEG) and the order of the bands is RGB. e.g ARRAY[3,2,1] means map band 3 to Red, band 2 to green and band 1 to blue
-  `quality` number from 0 to 100. The higher the number the crisper the image.
-  `options` text Array of GDAL options as defined for JPEG (look at create_options for JPEG [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers)). For JPEG valid ones are `PROGRESSIVE` ON or OFF and `QUALITY` a range from 0 to 100 and default to 75. Refer to [GDAL Raster format options](http://www.gdal.org/frmt_various.html) for more details.


Availability: 2.0.0 - requires GDAL >= 1.6.0.


## Examples: Output


```
-- output first 3 bands 75% quality
SELECT ST_AsJPEG(rast) As rastjpg
    FROM dummy_rast WHERE rid=2;

-- output only first band as 90% quality
SELECT ST_AsJPEG(rast,1,90) As rastjpg
    FROM dummy_rast WHERE rid=2;

-- output first 3 bands (but make band 2 Red, band 1 green, and band 3 blue, progressive and 90% quality
SELECT ST_AsJPEG(rast,ARRAY[2,1,3],ARRAY['QUALITY=90','PROGRESSIVE=ON']) As rastjpg
    FROM dummy_rast WHERE rid=2;
```


## See Also


[Building Custom Applications with PostGIS Raster](../raster-data-management-queries-and-applications/building-custom-applications-with-postgis-raster.md#RT_Raster_Applications), [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers), [RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster), [RT_ST_AsPNG](#RT_ST_AsPNG), [RT_ST_AsTIFF](#RT_ST_AsTIFF)
  <a id="RT_ST_AsPNG"></a>

# ST_AsPNG

Return the raster tile selected bands as a single portable network graphics (PNG) image (byte array). If 1, 3, or 4 bands in raster and no bands are specified, then all bands are used. If more 2 or more than 4 bands and no bands specified, then only band 1 is used. Bands are mapped to RGB or RGBA space.

## Synopsis


```sql
bytea ST_AsPNG(raster  rast, text[]  options=NULL)
bytea ST_AsPNG(raster  rast, integer  nband, integer  compression)
bytea ST_AsPNG(raster  rast, integer  nband, text[]  options=NULL)
bytea ST_AsPNG(raster  rast, integer[]  nbands, integer  compression)
bytea ST_AsPNG(raster  rast, integer[]  nbands, text[]  options=NULL)
```


## Description


Returns the selected bands of the raster as a single Portable Network Graphics Image (PNG). Use [RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster) if you need to export as less common raster types. If no band is specified, then the first 3 bands are exported. There are many variants of the function with many options. If no `srid` is specified then then srid of the raster is used. These are itemized below:


-  `nband` is for single band exports.
-  `nbands` is an array of bands to export (note that max is 4 for PNG) and the order of the bands is RGBA. e.g ARRAY[3,2,1] means map band 3 to Red, band 2 to green and band 1 to blue
-  `compression` number from 1 to 9. The higher the number the greater the compression.
-  `options` text Array of GDAL options as defined for PNG (look at create_options for PNG of [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers)). For PNG valid one is only ZLEVEL (amount of time to spend on compression -- default 6) e.g. ARRAY['ZLEVEL=9']. WORLDFILE is not allowed since the function would have to output two outputs. Refer to [GDAL Raster format options](http://www.gdal.org/frmt_various.html) for more details.


Availability: 2.0.0 - requires GDAL >= 1.6.0.


## Examples


```sql
SELECT ST_AsPNG(rast) As rastpng
FROM dummy_rast WHERE rid=2;

-- export the first 3 bands and map band 3 to Red, band 1 to Green, band 2 to blue
SELECT ST_AsPNG(rast, ARRAY[3,1,2]) As rastpng
FROM dummy_rast WHERE rid=2;

```


## See Also


[RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster), [RT_ST_ColorMap](raster-processing-map-algebra.md#RT_ST_ColorMap), [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers), [Building Custom Applications with PostGIS Raster](../raster-data-management-queries-and-applications/building-custom-applications-with-postgis-raster.md#RT_Raster_Applications)
  <a id="RT_ST_AsTIFF"></a>

# ST_AsTIFF

Return the raster selected bands as a single TIFF image (byte array). If no band is specified or any of specified bands does not exist in the raster, then will try to use all bands.

## Synopsis


```sql
bytea ST_AsTIFF(raster  rast, text[]  options='', integer  srid=sameassource)
bytea ST_AsTIFF(raster  rast, text  compression='', integer  srid=sameassource)
bytea ST_AsTIFF(raster  rast, integer[]  nbands, text  compression='', integer  srid=sameassource)
bytea ST_AsTIFF(raster  rast, integer[]  nbands, text[]  options, integer  srid=sameassource)
```


## Description


Returns the selected bands of the raster as a single Tagged Image File Format (TIFF). If no band is specified, will try to use all bands. This is a wrapper around [RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster). Use [RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster) if you need to export as less common raster types. There are many variants of the function with many options. If no spatial reference SRS text is present, the spatial reference of the raster is used. These are itemized below:


-  `nbands` is an array of bands to export (note that max is 3 for PNG) and the order of the bands is RGB. e.g ARRAY[3,2,1] means map band 3 to Red, band 2 to green and band 1 to blue
-  `compression` Compression expression -- JPEG90 (or some other percent), LZW, JPEG, DEFLATE9.
-  `options` text Array of GDAL create options as defined for GTiff (look at create_options for GTiff of [RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers)). or refer to [GDAL Raster format options](http://www.gdal.org/frmt_various.html) for more details.
-  `srid` srid of spatial_ref_sys of the raster. This is used to populate the georeference information


Availability: 2.0.0 - requires GDAL >= 1.6.0.


## Examples: Use jpeg compression 90%


```sql
SELECT ST_AsTIFF(rast, 'JPEG90') As rasttiff
FROM dummy_rast WHERE rid=2;

```


## See Also


[RT_ST_GDALDrivers](raster-management.md#RT_ST_GDALDrivers), [RT_ST_AsGDALRaster](#RT_ST_AsGDALRaster), [RT_ST_SRID](raster-accessors.md#RT_ST_SRID)
