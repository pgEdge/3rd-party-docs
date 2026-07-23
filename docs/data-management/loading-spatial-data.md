<a id="loading-data"></a>

## Loading Spatial Data


Once you have created a spatial table, you are ready to upload spatial data to the database. There are two built-in ways to get spatial data into a PostGIS/PostgreSQL database: using formatted SQL statements or using the Shapefile loader.
 <a id="load-data-sql"></a>

## Using SQL to Load Data


If spatial data can be converted to a text representation (as either WKT or WKB), then using SQL might be the easiest way to get data into PostGIS. Data can be bulk-loaded into PostGIS/PostgreSQL by loading a text file of SQL <code>INSERT</code> statements using the <code>psql</code> SQL utility.


A SQL load file (`roads.sql` for example) might look like this:


```sql
BEGIN;
INSERT INTO roads (road_id, roads_geom, road_name)
  VALUES (1,'LINESTRING(191232 243118,191108 243242)','Jeff Rd');
INSERT INTO roads (road_id, roads_geom, road_name)
  VALUES (2,'LINESTRING(189141 244158,189265 244817)','Geordie Rd');
INSERT INTO roads (road_id, roads_geom, road_name)
  VALUES (3,'LINESTRING(192783 228138,192612 229814)','Paul St');
INSERT INTO roads (road_id, roads_geom, road_name)
  VALUES (4,'LINESTRING(189412 252431,189631 259122)','Graeme Ave');
INSERT INTO roads (road_id, roads_geom, road_name)
  VALUES (5,'LINESTRING(190131 224148,190871 228134)','Phil Tce');
INSERT INTO roads (road_id, roads_geom, road_name)
  VALUES (6,'LINESTRING(198231 263418,198213 268322)','Dave Cres');
COMMIT;
```


The SQL file can be loaded into PostgreSQL using <code>psql</code>:


```
psql -d [database] -f roads.sql
```
  <a id="shp2pgsql_usage"></a>

## Using the Shapefile Loader


 The `shp2pgsql` data loader converts Shapefiles into SQL suitable for insertion into a PostGIS/PostgreSQL database either in geometry or geography format. The loader has several operating modes selected by command line flags.


There is also a `shp2pgsql-gui` graphical interface with most of the options as the command-line loader. This may be easier to use for one-off non-scripted loading or if you are new to PostGIS. It can also be configured as a plugin to PgAdminIII.


(c|a|d|p) These are mutually exclusive options:
:   `-c`
    :   Creates a new table and populates it from the Shapefile. *This is the default mode.*

    `-a`
    :   Appends data from the Shapefile into the database table. Note that to use this option to load multiple files, the files must have the same attributes and same data types.

    `-d`
    :   Drops the database table before creating a new table with the data in the Shapefile.

    `-p`
    :   Only produces the table creation SQL code, without adding any actual data. This can be used if you need to completely separate the table creation and data loading steps.

`-?`
:   Display help screen.

`-D`
:   Use the PostgreSQL "dump" format for the output data. This can be combined with -a, -c and -d. It is much faster to load than the default "insert" SQL format. Use this for very large data sets.

`-s [<from_srid>:]<srid>`
:   Creates and populates the geometry tables with the specified SRID. Optionally specifies that the input shapefile uses the given FROM_SRID, in which case the geometries will be reprojected to the target SRID.

`-k`
:   Keep identifiers' case (column, schema and attributes). Note that attributes in Shapefile are all UPPERCASE.

`-i`
:   Coerce all integers to standard 32-bit integers, do not create 64-bit bigints, even if the DBF header signature appears to warrant it.

`-I`
:   Create a GiST index on the geometry column.

`-m`
:   -m `a_file_name` Specify a file containing a set of mappings of (long) column names to 10 character DBF column names. The content of the file is one or more lines of two names separated by white space and no trailing or leading space. For example:

    ```
    COLUMNNAME DBFFIELD1
    AVERYLONGCOLUMNNAME DBFFIELD2
    ```

`-S`
:   Generate simple geometries instead of MULTI geometries. Will only succeed if all the geometries are actually single (I.E. a MULTIPOLYGON with a single shell, or or a MULTIPOINT with a single vertex).

`-t <dimensionality>`
:   Force the output geometry to have the specified dimensionality. Use the following strings to indicate the dimensionality: 2D, 3DZ, 3DM, 4D.


     If the input has fewer dimensions that specified, the output will have those dimensions filled in with zeroes. If the input has more dimensions that specified, the unwanted dimensions will be stripped.

`-w`
:   Output WKT format, instead of WKB. Note that this can introduce coordinate drifts due to loss of precision.

`-e`
:   Execute each statement on its own, without using a transaction. This allows loading of the majority of good data when there are some bad geometries that generate errors. Note that this cannot be used with the -D flag as the "dump" format always uses a transaction.

`-W <encoding>`
:   Specify encoding of the input data (dbf file). When used, all attributes of the dbf are converted from the specified encoding to UTF8. The resulting SQL output will contain a <code>SET CLIENT_ENCODING to UTF8</code> command, so that the backend will be able to reconvert from UTF8 to whatever encoding the database is configured to use internally.

`-N <policy>`
:   NULL geometries handling policy (insert*,skip,abort)

`-n`
:   -n Only import DBF file. If your data has no corresponding shapefile, it will automatically switch to this mode and load just the dbf. So setting this flag is only needed if you have a full shapefile set, and you only want the attribute data and no geometry.

`-G`
:   Use geography type instead of geometry (requires lon/lat data) in WGS84 long lat (SRID=4326)

`-T <tablespace>`
:   Specify the tablespace for the new table. Indexes will still use the default tablespace unless the -X parameter is also used. The PostgreSQL documentation has a good description on when to use custom tablespaces.

`-X <tablespace>`
:   Specify the tablespace for the new table's indexes. This applies to the primary key index, and the GIST spatial index if -I is also used.

`-Z`
:   When used, this flag will prevent the generation of <code>ANALYZE</code> statements. Without the -Z flag (default behavior), the <code>ANALYZE</code> statements will be generated.


 An example session using the loader to create an input file and loading it might look like this:


```

# shp2pgsql -c -D -s 4269 -i -I shaperoads.shp myschema.roadstable > roads.sql
# psql -d roadsdb -f roads.sql
```


 A conversion and load can be done in one step using UNIX pipes:


```
# shp2pgsql shaperoads.shp myschema.roadstable | psql -d roadsdb
```
