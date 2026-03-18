<a id="RT_Loading_Rasters"></a>

## Loading and Creating Rasters


For most use cases, you will create PostGIS rasters by loading existing raster files using the packaged `raster2pgsql` raster loader.
 <a id="RT_Raster_Loader"></a>

## Using raster2pgsql to load rasters


The `raster2pgsql` is a raster loader executable that loads GDAL supported raster formats into SQL suitable for loading into a PostGIS raster table. It is capable of loading folders of raster files as well as creating overviews of rasters.


Since the raster2pgsql is compiled as part of PostGIS most often (unless you compile your own GDAL library), the raster types supported by the executable will be the same as those compiled in the GDAL dependency library. To get a list of raster types your particular `raster2pgsql` supports use the `-G` switch.


!!! note

    When creating overviews of a specific factor from a set of rasters that are aligned, it is possible for the overviews to not align. Visit [http://trac.osgeo.org/postgis/ticket/1764](http://trac.osgeo.org/postgis/ticket/1764) for an example where the overviews do not align.


## Example Usage


An example session using the loader to create an input file and uploading it chunked in 100x100 tiles might look like this:


```

# -s use srid 4326
# -I create spatial index
# -C use standard raster constraints
# -M vacuum analyze after load
# *.tif load all these files
# -F include a filename column in the raster table
# -t tile the output 100x100
# public.demelevation load into this table
raster2pgsql -s 4326 -I -C -M -F -t 100x100 *.tif public.demelevation > elev.sql

# -d connect to this database
# -f read this file after connecting
psql -d gisdb -f elev.sql
```


!!! note

    If you do not specify the schema as part of the target table name, the table will be created in the default schema of the database or user you are connecting with.


A conversion and upload can be done all in one step using UNIX pipes:


```
raster2pgsql -s 4326 -I -C -M *.tif -F -t 100x100 public.demelevation | psql -d gisdb
```


Load rasters Massachusetts state plane meters aerial tiles into a schema called `aerial` and create a full view, 2 and 4 level overview tables, use copy mode for inserting (no intermediary file just straight to db), and -e don't force everything in a transaction (good if you want to see data in tables right away without waiting). Break up the rasters into 128x128 pixel tiles and apply raster constraints. Use copy mode instead of table insert. (-F) Include a field called filename to hold the name of the file the tiles were cut from.


```
raster2pgsql -I -C -e -Y -F -s 26986 -t 128x128  -l 2,4 bostonaerials2008/*.jpg aerials.boston | psql -U postgres -d gisdb -h localhost -p 5432
```


```
--get a list of raster types supported:
raster2pgsql -G
```


The -G commands outputs a list something like


```

Available GDAL raster formats:
  Virtual Raster
  GeoTIFF
  National Imagery Transmission Format
  Raster Product Format TOC format
  ECRG TOC format
  Erdas Imagine Images (.img)
  CEOS SAR Image
  CEOS Image
  ...
  Arc/Info Export E00 GRID
  ZMap Plus Grid
  NOAA NGS Geoid Height Grids
```


## raster2pgsql options


`-?`
:   Display help screen. Help will also display if you don't pass in any arguments.

`-G`
:   Print the supported raster formats.

(c|a|d|p) These are mutually exclusive options:
:   `-c`
    :   Create new table and populate it with raster(s), *this is the default mode*

    `-a`
    :   Append raster(s) to an existing table.

    `-d`
    :   Drop table, create new one and populate it with raster(s)

    `-p`
    :   Prepare mode, only create the table.

Raster processing: Applying constraints for proper registering in raster catalogs
:   `-C `
    :   Apply raster constraints -- srid, pixelsize etc. to ensure raster is properly registered in `raster_columns` view.

    `-x `
    :   Disable setting the max extent constraint. Only applied if -C flag is also used.

    `-r`
    :   Set the constraints (spatially unique and coverage tile) for regular blocking. Only applied if -C flag is also used.

Raster processing: Optional parameters used to manipulate input raster dataset
:   `-s <srid>`
    :   Assign output raster with specified SRID. If not provided or is zero, raster's metadata will be checked to determine an appropriate SRID.

    `-b BAND`
    :   Index (1-based) of band to extract from raster. For more than one band index, separate with comma (,). If unspecified, all bands of raster will be extracted.

    `-t TILE_SIZE`
    :   Cut raster into tiles to be inserted one per table row. `TILE_SIZE` is expressed as WIDTHxHEIGHT or set to the value "auto" to allow the loader to compute an appropriate tile size using the first raster and applied to all rasters.

    `-P`
    :   Pad right-most and bottom-most tiles to guarantee that all tiles have the same width and height.

    `-R, --register`
    :   Register the raster as a filesystem (out-db) raster.


        Only the metadata of the raster and path location to the raster is stored in the database (not the pixels).

    `-l OVERVIEW_FACTOR`
    :   Create overview of the raster. For more than one factor, separate with comma(,). Overview table name follows the pattern o_`overview factor`_`table`, where `overview factor` is a placeholder for numerical overview factor and `table` is replaced with the base table name. Created overview is stored in the database and is not affected by -R. Note that your generated sql file will contain both the main table and overview tables.

    `-N NODATA`
    :   NODATA value to use on bands without a NODATA value.

Optional parameters used to manipulate database objects
:   `-f COLUMN`
    :   Specify name of destination raster column, default is 'rast'

    `-F`
    :   Add a column with the name of the file

    `-n COLUMN`
    :   Specify the name of the filename column. Implies -F.

    `-q`
    :   Wrap PostgreSQL identifiers in quotes.

    `-I`
    :   Create a GiST index on the raster column.

    `-M`
    :   Vacuum analyze the raster table.

    `-k`
    :   Keeps empty tiles and skips NODATA value checks for each raster band. Note you save time in checking, but could end up with far more junk rows in your database and those junk rows are not marked as empty tiles.

    `-T tablespace`
    :   Specify the tablespace for the new table. Note that indices (including the primary key) will still use the default tablespace unless the -X flag is also used.

    `-X tablespace`
    :   Specify the tablespace for the table's new index. This applies to the primary key and the spatial index if the -I flag is used.

    `-Y max_rows_per_copy=50`
    :   Use copy statements instead of insert statements. Optionally specify `max_rows_per_copy`; default 50 when not specified.

`-e`
:   Execute each statement individually, do not use a transaction.

`-E ENDIAN`
:   Control endianness of generated binary output of raster; specify 0 for XDR and 1 for NDR (default); only NDR output is supported now

`-V version`
:   Specify version of output format. Default is 0. Only 0 is supported at this time.
   <a id="RT_Creating_Rasters"></a>

## Creating rasters using PostGIS raster functions


On many occasions, you'll want to create rasters and raster tables right in the database. There are a plethora of functions to do that. The general steps to follow.


1. Create a table with a raster column to hold the new raster records which can be accomplished with:

```sql
CREATE TABLE myrasters(rid serial primary key, rast raster);
```
2. There are many functions to help with that goal. If you are creating rasters not as a derivative of other rasters, you will want to start with: [RT_ST_MakeEmptyRaster](../raster-reference/raster-constructors.md#RT_ST_MakeEmptyRaster), followed by [RT_ST_AddBand](../raster-reference/raster-constructors.md#RT_ST_AddBand)

   You can also create rasters from geometries. To achieve that you'll want to use [RT_ST_AsRaster](../raster-reference/raster-constructors.md#RT_ST_AsRaster) perhaps accompanied with other functions such as [RT_ST_Union](../raster-reference/raster-processing-map-algebra.md#RT_ST_Union) or [RT_ST_MapAlgebraFct2](../raster-reference/raster-processing-map-algebra.md#RT_ST_MapAlgebraFct2) or any of the family of other map algebra functions.

   There are even many more options for creating new raster tables from existing tables. For example you can create a raster table in a different projection from an existing one using [RT_ST_Transform](../raster-reference/raster-editors.md#RT_ST_Transform)
3. Once you are done populating your table initially, you'll want to create a spatial index on the raster column with something like:

```sql
CREATE INDEX myrasters_rast_st_convexhull_idx ON myrasters USING gist( ST_ConvexHull(rast) );
```

   Note the use of [RT_ST_ConvexHull](../raster-reference/raster-processing-raster-to-geometry.md#RT_ST_ConvexHull) since most raster operators are based on the convex hull of the rasters.

!!! note

    Pre-2.0 versions of PostGIS raster were based on the envelop rather than the convex hull. For the spatial indexes to work properly you'll need to drop those and replace with convex hull based index.
4. Apply raster constraints using [RT_AddRasterConstraints](../raster-reference/raster-management.md#RT_AddRasterConstraints)
  <a id="RT_Cloud_Rasters"></a>

## Using "out db" cloud rasters


 The `raster2pgsql` tool uses GDAL to access raster data, and can take advantage of a key GDAL feature: the ability to read from rasters that are [stored remotely](https://gdal.org/user/virtual_file_systems.html#network-based-file-systems) in cloud "object stores" (e.g. AWS S3, Google Cloud Storage).


 Efficient use of cloud stored rasters requires the use of a "cloud optimized" format. The most well-known and widely used is the "[cloud optimized GeoTIFF](https://gdal.org/drivers/raster/cog.html)" format. Using a non-cloud format, like a JPEG, or an un-tiled TIFF will result in very poor performance, as the system will have to download the entire raster each time it needs to access a subset.


 First, load your raster into the cloud storage of your choice. Once it is loaded, you will have a URI to access it with, either an "http" URI, or sometimes a URI specific to the service. (e.g., "s3://bucket/object"). To access non-public buckets, you will need to supply GDAL config options to authenticate your connection. Note that this command is *reading* from the cloud raster and *writing* to the database.


```
AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx \
AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx \
raster2pgsql \
  -s 990000 \
  -t 256x256 \
  -I \
  -R \
  /vsis3/your.bucket.com/your_file.tif \
  your_table \
  | psql your_db
```


 Once the table is loaded, you need to give the database permission to read from remote rasters, by setting two permissions, [postgis_enable_outdb_rasters](../postgis-reference/grand-unified-custom-variables-gucs.md#postgis_enable_outdb_rasters) and [postgis_gdal_enabled_drivers](../postgis-reference/grand-unified-custom-variables-gucs.md#postgis_gdal_enabled_drivers).


```sql
SET postgis.enable_outdb_rasters = true;
SET postgis.gdal_enabled_drivers TO 'ENABLE_ALL';

```


 To make the changes sticky, set them directly on your database. You will need to re-connect to experience the new settings.


```sql
ALTER DATABASE your_db SET postgis.enable_outdb_rasters = true;
ALTER DATABASE your_db SET postgis.gdal_enabled_drivers TO 'ENABLE_ALL';

```


 For non-public rasters, you may have to provide access keys to read from the cloud rasters. The same keys you used to write the `raster2pgsql` call can be set for use inside the database, with the [postgis_gdal_vsi_options](../postgis-reference/grand-unified-custom-variables-gucs.md#postgis_gdal_vsi_options) configuration. Note that multiple options can be set by space-separating the `key=value` pairs.


```sql
SET postgis.gdal_vsi_options = 'AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx';
```


 Once you have the data loaded and permissions set you can interact with the raster table like any other raster table, using the same functions. The database will handle all the mechanics of connecting to the cloud data when it needs to read pixel data.
