<a id="Raster_Management_Functions"></a>

## Raster Management
  <a id="RT_AddRasterConstraints"></a>

# AddRasterConstraints

Adds raster constraints to a loaded raster table for a specific column that constrains spatial ref, scaling, blocksize, alignment, bands, band type and a flag to denote if raster column is regularly blocked. The table must be loaded with data for the constraints to be inferred. Returns true if the constraint setting was accomplished and issues a notice otherwise.

## Synopsis


```sql
boolean AddRasterConstraints(name
            rasttable, name
            rastcolumn, boolean
            srid=true, boolean
            scale_x=true, boolean
            scale_y=true, boolean
            blocksize_x=true, boolean
            blocksize_y=true, boolean
            same_alignment=true, boolean
            regular_blocking=false, boolean
                num_bands=true, boolean
                pixel_types=true, boolean
                nodata_values=true, boolean
                out_db=true, boolean
                extent=true)
boolean AddRasterConstraints(name
            rasttable, name
            rastcolumn, text[]
            VARIADIC constraints)
boolean AddRasterConstraints(name
            rastschema, name
            rasttable, name
            rastcolumn, text[]
            VARIADIC constraints)
boolean AddRasterConstraints(name
            rastschema, name
            rasttable, name
            rastcolumn, boolean
            srid=true, boolean
            scale_x=true, boolean
            scale_y=true, boolean
            blocksize_x=true, boolean
            blocksize_y=true, boolean
            same_alignment=true, boolean
            regular_blocking=false, boolean
            num_bands=true, boolean
            pixel_types=true, boolean
                nodata_values=true, boolean
                out_db=true, boolean
                extent=true)
```


## Description


Generates constraints on a raster column that are used to display information in the `raster_columns` raster catalog. The `rastschema` is the name of the table schema the table resides in. The `srid` must be an integer value reference to an entry in the SPATIAL_REF_SYS table.


`raster2pgsql` loader uses this function to register raster tables


Valid constraint names to pass in: refer to [Raster Columns Catalog](../raster-data-management-queries-and-applications/raster-catalogs.md#RT_Raster_Columns) for more details.


- `blocksize` sets both X and Y blocksize
- `blocksize_x` sets X tile (width in pixels of each tile)
- `blocksize_y` sets Y tile (height in pixels of each tile)
- `extent` computes extent of whole table and applies constraint all rasters must be within that extent
- `num_bands` number of bands
- `pixel_types` reads array of pixel types for each band ensure all band n have same pixel type
- `regular_blocking` sets spatially unique (no two rasters can be spatially the same) and coverage tile (raster is aligned to a coverage) constraints
- `same_alignment` ensures they all have same alignment meaning any two tiles you compare will return true for. Refer to [RT_ST_SameAlignment](raster-and-raster-band-spatial-relationships.md#RT_ST_SameAlignment).
- `srid` ensures all have same srid
- More -- any listed as inputs into the above functions


!!! note

    This function infers the constraints from the data already present in the table. As such for it to work, you must create the raster column first and then load it with data.


!!! note

    If you need to load more data in your tables after you have already applied constraints, you may want to run the DropRasterConstraints if the extent of your data has changed.


Availability: 2.0.0


## Examples: Apply all possible constraints on column based on data


```sql
CREATE TABLE myrasters(rid SERIAL primary key, rast raster);
INSERT INTO myrasters(rast)
SELECT ST_AddBand(ST_MakeEmptyRaster(1000, 1000, 0.3, -0.3, 2, 2, 0, 0,4326), 1, '8BSI'::text, -129, NULL);

SELECT AddRasterConstraints('myrasters'::name, 'rast'::name);


-- verify if registered correctly in the raster_columns view --
SELECT srid, scale_x, scale_y, blocksize_x, blocksize_y, num_bands, pixel_types, nodata_values
    FROM raster_columns
    WHERE r_table_name = 'myrasters';

 srid | scale_x | scale_y | blocksize_x | blocksize_y | num_bands | pixel_types| nodata_values
------+---------+---------+-------------+-------------+-----------+-------------+---------------
 4326 |       2 |       2 |        1000 |        1000 |         1 | {8BSI}      | {0}

```


## Examples: Apply single constraint


```sql
CREATE TABLE public.myrasters2(rid SERIAL primary key, rast raster);
INSERT INTO myrasters2(rast)
SELECT ST_AddBand(ST_MakeEmptyRaster(1000, 1000, 0.3, -0.3, 2, 2, 0, 0,4326), 1, '8BSI'::text, -129, NULL);

SELECT AddRasterConstraints('public'::name, 'myrasters2'::name, 'rast'::name,'regular_blocking', 'blocksize');
-- get notice--
NOTICE:  Adding regular blocking constraint
NOTICE:  Adding blocksize-X constraint
NOTICE:  Adding blocksize-Y constraint
```


## See Also


[Raster Columns Catalog](../raster-data-management-queries-and-applications/raster-catalogs.md#RT_Raster_Columns), [RT_ST_AddBand](raster-constructors.md#RT_ST_AddBand), [RT_ST_MakeEmptyRaster](raster-constructors.md#RT_ST_MakeEmptyRaster), [RT_DropRasterConstraints](#RT_DropRasterConstraints), [RT_ST_BandPixelType](raster-band-accessors.md#RT_ST_BandPixelType), [RT_ST_SRID](raster-accessors.md#RT_ST_SRID)
  <a id="RT_DropRasterConstraints"></a>

# DropRasterConstraints

Drops PostGIS raster constraints that refer to a raster table column. Useful if you need to reload data or update your raster column data.

## Synopsis


```sql
boolean DropRasterConstraints(name
            rasttable, name
            rastcolumn, boolean
            srid, boolean
            scale_x, boolean
            scale_y, boolean
            blocksize_x, boolean
            blocksize_y, boolean
            same_alignment, boolean
            regular_blocking, boolean
            num_bands=true, boolean
            pixel_types=true, boolean
            nodata_values=true, boolean
                out_db=true, boolean
            extent=true)
boolean DropRasterConstraints(name
            rastschema, name
            rasttable, name
            rastcolumn, boolean
            srid=true, boolean
            scale_x=true, boolean
            scale_y=true, boolean
            blocksize_x=true, boolean
            blocksize_y=true, boolean
            same_alignment=true, boolean
            regular_blocking=false, boolean
            num_bands=true, boolean
            pixel_types=true, boolean
            nodata_values=true, boolean
                out_db=true, boolean
            extent=true)
boolean DropRasterConstraints(name
            rastschema, name
            rasttable, name
            rastcolumn, text[]
            constraints)
```


## Description


Drops PostGIS raster constraints that refer to a raster table column that were added by [RT_AddRasterConstraints](#RT_AddRasterConstraints). Useful if you need to load more data or update your raster column data. You do not need to do this if you want to get rid of a raster table or a raster column.


To drop a raster table use the standard

```sql
DROP TABLE mytable
```


To drop just a raster column and leave the rest of the table, use standard SQL

```sql
ALTER TABLE mytable DROP COLUMN rast
```


the table will disappear from the `raster_columns` catalog if the column or table is dropped. However if only the constraints are dropped, the raster column will still be listed in the `raster_columns` catalog, but there will be no other information about it aside from the column name and table.


Availability: 2.0.0


## Examples


```sql

SELECT DropRasterConstraints ('myrasters','rast');
----RESULT output ---
t

-- verify change in raster_columns --
SELECT srid, scale_x, scale_y, blocksize_x, blocksize_y, num_bands, pixel_types, nodata_values
    FROM raster_columns
    WHERE r_table_name = 'myrasters';

 srid | scale_x | scale_y | blocksize_x | blocksize_y | num_bands | pixel_types| nodata_values
------+---------+---------+-------------+-------------+-----------+-------------+---------------
    0 |         |         |             |             |           |             |

```


## See Also


[RT_AddRasterConstraints](#RT_AddRasterConstraints)
  <a id="RT_AddOverviewConstraints"></a>

# AddOverviewConstraints

Tag a raster column as being an overview of another.

## Synopsis


```sql
boolean AddOverviewConstraints(name
        ovschema, name
        ovtable, name
        ovcolumn, name
        refschema, name
        reftable, name
        refcolumn, int
        ovfactor)
boolean AddOverviewConstraints(name
        ovtable, name
        ovcolumn, name
        reftable, name
        refcolumn, int
        ovfactor)
```


## Description


 Adds constraints on a raster column that are used to display information in the `raster_overviews` raster catalog.


 The `ovfactor` parameter represents the scale multiplier in the overview column: higher overview factors have lower resolution.


 When the `ovschema` and `refschema` parameters are omitted, the first table found scanning the `search_path` will be used.


Availability: 2.0.0


## Examples


```sql

CREATE TABLE res1 AS SELECT
ST_AddBand(
  ST_MakeEmptyRaster(1000, 1000, 0, 0, 2),
  1, '8BSI'::text, -129, NULL
) r1;

CREATE TABLE res2 AS SELECT
ST_AddBand(
  ST_MakeEmptyRaster(500, 500, 0, 0, 4),
  1, '8BSI'::text, -129, NULL
) r2;

SELECT AddOverviewConstraints('res2', 'r2', 'res1', 'r1', 2);

-- verify if registered correctly in the raster_overviews view --
SELECT o_table_name ot, o_raster_column oc,
       r_table_name rt, r_raster_column rc,
       overview_factor f
FROM raster_overviews WHERE o_table_name = 'res2';
  ot  | oc |  rt  | rc | f
------+----+------+----+---
 res2 | r2 | res1 | r1 | 2
(1 row)

```


## See Also


 [Raster Overviews](../raster-data-management-queries-and-applications/raster-catalogs.md#RT_Raster_Overviews), [RT_DropOverviewConstraints](#RT_DropOverviewConstraints), [RT_CreateOverview](#RT_CreateOverview), [RT_AddRasterConstraints](#RT_AddRasterConstraints)
  <a id="RT_DropOverviewConstraints"></a>

# DropOverviewConstraints

Untag a raster column from being an overview of another.

## Synopsis


```sql
boolean DropOverviewConstraints(name
        ovschema, name
        ovtable, name
        ovcolumn)
boolean DropOverviewConstraints(name
        ovtable, name
        ovcolumn)
```


## Description


 Remove from a raster column the constraints used to show it as being an overview of another in the `raster_overviews` raster catalog.


 When the `ovschema` parameter is omitted, the first table found scanning the `search_path` will be used.


Availability: 2.0.0


## See Also


 [Raster Overviews](../raster-data-management-queries-and-applications/raster-catalogs.md#RT_Raster_Overviews), [RT_AddOverviewConstraints](#RT_AddOverviewConstraints), [RT_DropRasterConstraints](#RT_DropRasterConstraints)
  <a id="RT_PostGIS_GDAL_Version"></a>

# PostGIS_GDAL_Version

Reports the version of the GDAL library in use by PostGIS.

## Synopsis


```sql
text PostGIS_GDAL_Version()
```


## Description


Reports the version of the GDAL library in use by PostGIS. Will also check and report if GDAL can find its data files.


## Examples


```sql

SELECT PostGIS_GDAL_Version();
       postgis_gdal_version
-----------------------------------
 GDAL 1.11dev, released 2013/04/13

```


## See Also


 [postgis_gdal_datapath](../postgis-reference/grand-unified-custom-variables-gucs.md#postgis_gdal_datapath)
  <a id="RT_PostGIS_Raster_Lib_Build_Date"></a>

# PostGIS_Raster_Lib_Build_Date

Reports full raster library build date.

## Synopsis


```sql
text PostGIS_Raster_Lib_Build_Date()
```


## Description


Reports raster build date


## Examples


```sql
SELECT PostGIS_Raster_Lib_Build_Date();
postgis_raster_lib_build_date
-----------------------------
2010-04-28 21:15:10
```


## See Also


 [RT_PostGIS_Raster_Lib_Version](#RT_PostGIS_Raster_Lib_Version)
  <a id="RT_PostGIS_Raster_Lib_Version"></a>

# PostGIS_Raster_Lib_Version

Reports full raster version and build configuration infos.

## Synopsis


```sql
text PostGIS_Raster_Lib_Version()
```


## Description


Reports full raster version and build configuration infos.


## Examples


```sql
SELECT PostGIS_Raster_Lib_Version();
postgis_raster_lib_version
-----------------------------
 2.0.0
```


## See Also


 [PostGIS_Lib_Version](../postgis-reference/version-functions.md#PostGIS_Lib_Version)
  <a id="RT_ST_GDALDrivers"></a>

# ST_GDALDrivers

Returns a list of raster formats supported by PostGIS through GDAL. Only those formats with can_write=True can be used by ST_AsGDALRaster

## Synopsis


```sql
setof record ST_GDALDrivers(integer  OUT idx, text  OUT short_name, text  OUT long_name, text  OUT can_read, text  OUT can_write, text  OUT create_options)
```


## Description


 Returns a list of raster formats short_name,long_name and creator options of each format supported by GDAL. Use the short_name as input in the `format` parameter of [RT_ST_AsGDALRaster](raster-outputs.md#RT_ST_AsGDALRaster). Options vary depending on what drivers your libgdal was compiled with. `create_options` returns an xml formatted set of CreationOptionList/Option consisting of name and optional `type`, `description` and set of `VALUE` for each creator option for the specific driver.


Changed: 2.5.0 - add can_read and can_write columns.


Changed: 2.0.6, 2.1.3 - by default no drivers are enabled, unless GUC or Environment variable gdal_enabled_drivers is set.


Availability: 2.0.0 - requires GDAL >= 1.6.0.


## Examples: List of Drivers


```sql
SET postgis.gdal_enabled_drivers = 'ENABLE_ALL';
SELECT short_name, long_name, can_write
FROM st_gdaldrivers()
ORDER BY short_name;

   short_name    |                          long_name                          | can_write
-----------------+-------------------------------------------------------------+-----------
 AAIGrid         | Arc/Info ASCII Grid                                         | t
 ACE2            | ACE2                                                        | f
 ADRG            | ARC Digitized Raster Graphics                               | f
 AIG             | Arc/Info Binary Grid                                        | f
 AirSAR          | AirSAR Polarimetric Image                                   | f
 ARG             | Azavea Raster Grid format                                   | t
 BAG             | Bathymetry Attributed Grid                                  | f
 BIGGIF          | Graphics Interchange Format (.gif)                          | f
 BLX             | Magellan topo (.blx)                                        | t
 BMP             | MS Windows Device Independent Bitmap                        | f
 BSB             | Maptech BSB Nautical Charts                                 | f
 PAux            | PCI .aux Labelled                                           | f
 PCIDSK          | PCIDSK Database File                                        | f
 PCRaster        | PCRaster Raster File                                        | f
 PDF             | Geospatial PDF                                              | f
 PDS             | NASA Planetary Data System                                  | f
 PDS4            | NASA Planetary Data System 4                                | t
 PLMOSAIC        | Planet Labs Mosaics API                                     | f
 PLSCENES        | Planet Labs Scenes API                                      | f
 PNG             | Portable Network Graphics                                   | t
 PNM             | Portable Pixmap Format (netpbm)                             | f
 PRF             | Racurs PHOTOMOD PRF                                         | f
 R               | R Object Data Store                                         | t
 Rasterlite      | Rasterlite                                                  | t
 RDA             | DigitalGlobe Raster Data Access driver                      | f
 RIK             | Swedish Grid RIK (.rik)                                     | f
 RMF             | Raster Matrix Format                                        | f
 ROI_PAC         | ROI_PAC raster                                              | f
 RPFTOC          | Raster Product Format TOC format                            | f
 RRASTER         | R Raster                                                    | f
 RS2             | RadarSat 2 XML Product                                      | f
 RST             | Idrisi Raster A.1                                           | t
 SAFE            | Sentinel-1 SAR SAFE Product                                 | f
 SAGA            | SAGA GIS Binary Grid (.sdat, .sg-grd-z)                     | t
 SAR_CEOS        | CEOS SAR Image                                              | f
 SDTS            | SDTS Raster                                                 | f
 SENTINEL2       | Sentinel 2                                                  | f
 SGI             | SGI Image File Format 1.0                                   | f
 SNODAS          | Snow Data Assimilation System                               | f
 SRP             | Standard Raster Product (ASRP/USRP)                         | f
 SRTMHGT         | SRTMHGT File Format                                         | t
 Terragen        | Terragen heightfield                                        | f
 TIL             | EarthWatch .TIL                                             | f
 TSX             | TerraSAR-X Product                                          | f
 USGSDEM         | USGS Optional ASCII DEM (and CDED)                          | t
 VICAR           | MIPL VICAR file                                             | f
 VRT             | Virtual Raster                                              | t
 WCS             | OGC Web Coverage Service                                    | f
 WMS             | OGC Web Map Service                                         | t
 WMTS            | OGC Web Map Tile Service                                    | t
 XPM             | X11 PixMap Format                                           | t
 XYZ             | ASCII Gridded XYZ                                           | t
 ZMap            | ZMap Plus Grid                                              | t
```


## Example: List of options for each driver


```
-- Output the create options XML column of JPEG as a table  --
-- Note you can use these creator options in ST_AsGDALRaster options argument
SELECT (xpath('@name', g.opt))[1]::text As oname,
       (xpath('@type', g.opt))[1]::text As otype,
       (xpath('@description', g.opt))[1]::text As descrip
FROM (SELECT unnest(xpath('/CreationOptionList/Option', create_options::xml)) As opt
FROM  st_gdaldrivers()
WHERE short_name = 'JPEG') As g;

       oname        |  otype  |      descrip
--------------------+---------+-------------------------------------------------
 PROGRESSIVE        | boolean | whether to generate a progressive JPEG
 QUALITY            | int     | good=100, bad=0, default=75
 WORLDFILE          | boolean | whether to generate a worldfile
 INTERNAL_MASK      | boolean | whether to generate a validity mask
 COMMENT            | string  | Comment
 SOURCE_ICC_PROFILE | string  | ICC profile encoded in Base64
 EXIF_THUMBNAIL     | boolean | whether to generate an EXIF thumbnail(overview).
                                By default its max dimension will be 128
 THUMBNAIL_WIDTH    | int     | Forced thumbnail width
 THUMBNAIL_HEIGHT   | int     | Forced thumbnail height
(9 rows)
```


```


-- raw xml output for creator options for GeoTiff --
SELECT create_options
FROM st_gdaldrivers()
WHERE short_name = 'GTiff';

<CreationOptionList>
    <Option name="COMPRESS" type="string-select">
        <Value>NONE</Value>
        <Value>LZW</Value>
        <Value>PACKBITS</Value>
        <Value>JPEG</Value>
        <Value>CCITTRLE</Value>
        <Value>CCITTFAX3</Value>
        <Value>CCITTFAX4</Value>
        <Value>DEFLATE</Value>
    </Option>
    <Option name="PREDICTOR" type="int" description="Predictor Type"/>
    <Option name="JPEG_QUALITY" type="int" description="JPEG quality 1-100" default="75"/>
    <Option name="ZLEVEL" type="int" description="DEFLATE compression level 1-9" default="6"/>
    <Option name="NBITS" type="int" description="BITS for sub-byte files (1-7), sub-uint16 (9-15), sub-uint32 (17-31)"/>
    <Option name="INTERLEAVE" type="string-select" default="PIXEL">
        <Value>BAND</Value>
        <Value>PIXEL</Value>
    </Option>
    <Option name="TILED" type="boolean" description="Switch to tiled format"/>
    <Option name="TFW" type="boolean" description="Write out world file"/>
    <Option name="RPB" type="boolean" description="Write out .RPB (RPC) file"/>
    <Option name="BLOCKXSIZE" type="int" description="Tile Width"/>
    <Option name="BLOCKYSIZE" type="int" description="Tile/Strip Height"/>
    <Option name="PHOTOMETRIC" type="string-select">
        <Value>MINISBLACK</Value>
        <Value>MINISWHITE</Value>
        <Value>PALETTE</Value>
        <Value>RGB</Value>
        <Value>CMYK</Value>
        <Value>YCBCR</Value>
        <Value>CIELAB</Value>
        <Value>ICCLAB</Value>
        <Value>ITULAB</Value>
    </Option>
    <Option name="SPARSE_OK" type="boolean" description="Can newly created files have missing blocks?" default="FALSE"/>
    <Option name="ALPHA" type="boolean" description="Mark first extrasample as being alpha"/>
    <Option name="PROFILE" type="string-select" default="GDALGeoTIFF">
        <Value>GDALGeoTIFF</Value>
        <Value>GeoTIFF</Value>
        <Value>BASELINE</Value>
    </Option>
    <Option name="PIXELTYPE" type="string-select">
        <Value>DEFAULT</Value>
        <Value>SIGNEDBYTE</Value>
    </Option>
    <Option name="BIGTIFF" type="string-select" description="Force creation of BigTIFF file">
        <Value>YES</Value>
        <Value>NO</Value>
        <Value>IF_NEEDED</Value>
        <Value>IF_SAFER</Value>
    </Option>
    <Option name="ENDIANNESS" type="string-select" default="NATIVE" description="Force endianness of created file. For DEBUG purpose mostly">
        <Value>NATIVE</Value>
        <Value>INVERTED</Value>
        <Value>LITTLE</Value>
        <Value>BIG</Value>
    </Option>
    <Option name="COPY_SRC_OVERVIEWS" type="boolean" default="NO" description="Force copy of overviews of source dataset (CreateCopy())"/>
</CreationOptionList>

-- Output the create options XML column for GTiff as a table  --
SELECT (xpath('@name', g.opt))[1]::text As oname,
       (xpath('@type', g.opt))[1]::text As otype,
       (xpath('@description', g.opt))[1]::text As descrip,
       array_to_string(xpath('Value/text()', g.opt),', ') As vals
FROM (SELECT unnest(xpath('/CreationOptionList/Option', create_options::xml)) As opt
FROM  st_gdaldrivers()
WHERE short_name = 'GTiff') As g;

       oname        |     otype     |                               descrip                                |                                   vals
--------------------+---------------+----------------------------------------------------------------------+---------------------------------------------------------------------------
 COMPRESS           | string-select |                                                                      | NONE, LZW, PACKBITS, JPEG, CCITTRLE, CCITTFAX3, CCITTFAX4, DEFLATE
 PREDICTOR          | int           | Predictor Type                                                       |
 JPEG_QUALITY       | int           | JPEG quality 1-100                                                   |
 ZLEVEL             | int           | DEFLATE compression level 1-9                                        |
 NBITS              | int           | BITS for sub-byte files (1-7), sub-uint16 (9-15), sub-uint32 (17-31) |
 INTERLEAVE         | string-select |                                                                      | BAND, PIXEL
 TILED              | boolean       | Switch to tiled format                                               |
 TFW                | boolean       | Write out world file                                                 |
 RPB                | boolean       | Write out .RPB (RPC) file                                            |
 BLOCKXSIZE         | int           | Tile Width                                                           |
 BLOCKYSIZE         | int           | Tile/Strip Height                                                    |
 PHOTOMETRIC        | string-select |                                                                      | MINISBLACK, MINISWHITE, PALETTE, RGB, CMYK, YCBCR, CIELAB, ICCLAB, ITULAB
 SPARSE_OK          | boolean       | Can newly created files have missing blocks?                         |
 ALPHA              | boolean       | Mark first extrasample as being alpha                                |
 PROFILE            | string-select |                                                                      | GDALGeoTIFF, GeoTIFF, BASELINE
 PIXELTYPE          | string-select |                                                                      | DEFAULT, SIGNEDBYTE
 BIGTIFF            | string-select | Force creation of BigTIFF file                                       | YES, NO, IF_NEEDED, IF_SAFER
 ENDIANNESS         | string-select | Force endianness of created file. For DEBUG purpose mostly           | NATIVE, INVERTED, LITTLE, BIG
 COPY_SRC_OVERVIEWS | boolean       | Force copy of overviews of source dataset (CreateCopy())             |
(19 rows)
```


## See Also


[RT_ST_AsGDALRaster](raster-outputs.md#RT_ST_AsGDALRaster), [ST_SRID](../postgis-reference/spatial-reference-system-functions.md#ST_SRID), [postgis_gdal_enabled_drivers](../postgis-reference/grand-unified-custom-variables-gucs.md#postgis_gdal_enabled_drivers)
  <a id="RT_ST_Contour"></a>

# ST_Contour

Generates a set of vector contours from the provided raster band, using the [GDAL contouring algorithm](https://gdal.org/api/gdal_alg.html?highlight=contour#_CPPv421GDALContourGenerateEx15GDALRasterBandHPv12CSLConstList16GDALProgressFuncPv).

## Synopsis


```sql
setof record ST_Contour(raster rast, integer bandnumber=1, double precision level_interval=100.0, double precision level_base=0.0, double precision[] fixed_levels=ARRAY[], boolean polygonize=false)
```


## Description


 Generates a set of vector contours from the provided raster band, using the [GDAL contouring algorithm](https://gdal.org/api/gdal_alg.html?highlight=contour#_CPPv421GDALContourGenerateEx15GDALRasterBandHPv12CSLConstList16GDALProgressFuncPv).


 When the `fixed_levels` parameter is a non-empty array, the `level_interval` and `level_base` parameters are ignored.


 Input parameters are:


`rast`
:   The raster to generate the contour of

`bandnumber`
:   The band to generate the contour of

`level_interval`
:   The elevation interval between contours generated

`level_base`
:   The "base" relative to which contour intervals are applied, this is normally zero, but could be different. To generate 10m contours at 5, 15, 25, ... the LEVEL_BASE would be 5.

`fixed_levels`
:   The elevation interval between contours generated

`polygonize`
:   If `true`, contour polygons will be created, rather than polygon lines.


 Return values are a set of records with the following attributes:


`geom`
:   The geometry of the contour line.

`id`
:   A unique identifier given to the contour line by GDAL.

`value`
:   The raster value the line represents. For an elevation DEM input, this would be the elevation of the output contour.


Availability: 3.2.0


## Example


```sql
WITH c AS (
SELECT (ST_Contour(rast, 1, fixed_levels => ARRAY[100.0, 200.0, 300.0])).*
FROM dem_grid WHERE rid = 1
)
SELECT st_astext(geom), id, value
FROM c;
```


## See Also


 [RT_ST_InterpolateRaster](#RT_ST_InterpolateRaster)
  <a id="RT_ST_InterpolateRaster"></a>

# ST_InterpolateRaster

Interpolates a gridded surface based on an input set of 3-d points, using the X- and Y-values to position the points on the grid and the Z-value of the points as the surface elevation.

## Synopsis


```sql
raster ST_InterpolateRaster(geometry input_points, text algorithm_options, raster template, integer template_band_num=1)
```


## Description


 Interpolates a gridded surface based on an input set of 3-d points, using the X- and Y-values to position the points on the grid and the Z-value of the points as the surface elevation. There are five interpolation algorithms available: inverse distance, inverse distance nearest-neighbor, moving average, nearest neighbor, and linear interpolation. See the [gdal_grid documentation](https://gdal.org/programs/gdal_grid.html#interpolation-algorithms) for more details on the algorithms and their parameters. For more information on how interpolations are calculated, see the [GDAL grid tutorial](https://gdal.org/tutorials/gdal_grid_tut.html).


 Input parameters are:


`input_points`
:   The points to drive the interpolation. Any geometry with Z-values is acceptable, all points in the input will be used.

`algorithm_options`
:   A string defining the algorithm and algorithm options, in the format used by [gdal_grid](https://gdal.org/programs/gdal_grid.html#interpolation-algorithms). For example, for an inverse-distance interpolation with a smoothing of 2, you would use "invdist:smoothing=2.0"

`template`
:   A raster template to drive the geometry of the output raster. The width, height, pixel size, spatial extent and pixel type will be read from this template.

`template_band_num`
:   By default the first band in the template raster is used to drive the output raster, but that can be adjusted with this parameter.


Availability: 3.2.0


## Example


```sql
SELECT ST_InterpolateRaster(
    'MULTIPOINT(10.5 9.5 1000, 11.5 8.5 1000, 10.5 8.5 500, 11.5 9.5 500)'::geometry,
    'invdist:smoothing:2.0',
    ST_AddBand(ST_MakeEmptyRaster(200, 400, 10, 10, 0.01, -0.005, 0, 0), '16BSI')
)
```


## See Also


 [RT_ST_Contour](#RT_ST_Contour)
  <a id="RT_UpdateRasterSRID"></a>

# UpdateRasterSRID

Change the SRID of all rasters in the user-specified column and table.

## Synopsis


```sql
raster UpdateRasterSRID(name  schema_name, name  table_name, name  column_name, integer  new_srid)
raster UpdateRasterSRID(name  table_name, name  column_name, integer  new_srid)
```


## Description


 Change the SRID of all rasters in the user-specified column and table. The function will drop all appropriate column constraints (extent, alignment and SRID) before changing the SRID of the specified column's rasters.


!!! note

    The data (band pixel values) of the rasters are not touched by this function. Only the raster's metadata is changed.


Availability: 2.1.0


## See Also


 [UpdateGeometrySRID](../postgis-reference/table-management-functions.md#UpdateGeometrySRID)
  <a id="RT_CreateOverview"></a>

# ST_CreateOverview

Create an reduced resolution version of a given raster coverage.

## Synopsis


```sql
regclass ST_CreateOverview(regclass  tab, name  col, int  factor, text  algo='NearestNeighbor')
```


## Description


 Create an overview table with resampled tiles from the source table. Output tiles will have the same size of input tiles and cover the same spatial extent with a lower resolution (pixel size will be 1/`factor` of the original in both directions).


 The overview table will be made available in the `raster_overviews` catalog and will have raster constraints enforced.


Algorithm options are: 'NearestNeighbor', 'Bilinear', 'Cubic', 'CubicSpline', and 'Lanczos'. Refer to: [GDAL Warp resampling methods](http://www.gdal.org/gdalwarp.html) for more details.


Availability: 2.2.0


## Example


Output to generally better quality but slower to product format


```sql
SELECT ST_CreateOverview('mydata.mytable'::regclass, 'rast', 2, 'Lanczos');
```


Output to faster to process default nearest neighbor


```sql
SELECT ST_CreateOverview('mydata.mytable'::regclass, 'rast', 2);
```


## See Also


 [RT_Retile](raster-constructors.md#RT_Retile), [RT_AddOverviewConstraints](#RT_AddOverviewConstraints), [RT_AddRasterConstraints](#RT_AddRasterConstraints), [Raster Overviews](../raster-data-management-queries-and-applications/raster-catalogs.md#RT_Raster_Overviews)
