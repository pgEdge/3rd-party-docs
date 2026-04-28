<a id="extracting-data"></a>

## Extracting Spatial Data


Spatial data can be extracted from the database using either SQL or the Shapefile dumper. The section on SQL presents some of the functions available to do comparisons and queries on spatial tables.
 <a id="extract-data-sql"></a>

## Using SQL to Extract Data


The most straightforward way of extracting spatial data out of the database is to use a SQL <code>SELECT</code> query to define the data set to be extracted and dump the resulting columns into a parsable text file:


```
db=# SELECT road_id, ST_AsText(road_geom) AS geom, road_name FROM roads;

road_id | geom                                    | road_name
--------+-----------------------------------------+-----------
	  1 | LINESTRING(191232 243118,191108 243242) | Jeff Rd
	  2 | LINESTRING(189141 244158,189265 244817) | Geordie Rd
	  3 | LINESTRING(192783 228138,192612 229814) | Paul St
	  4 | LINESTRING(189412 252431,189631 259122) | Graeme Ave
	  5 | LINESTRING(190131 224148,190871 228134) | Phil Tce
	  6 | LINESTRING(198231 263418,198213 268322) | Dave Cres
	  7 | LINESTRING(218421 284121,224123 241231) | Chris Way
(6 rows)
```


There will be times when some kind of restriction is necessary to cut down the number of records returned. In the case of attribute-based restrictions, use the same SQL syntax as used with a non-spatial table. In the case of spatial restrictions, the following functions are useful:


`ST_Intersects`
:   This function tells whether two geometries share any space.

<code>=</code>
:   This tests whether two geometries are geometrically identical. For example, if 'POLYGON((0 0,1 1,1 0,0 0))' is the same as 'POLYGON((0 0,1 1,1 0,0 0))' (it is).


Next, you can use these operators in queries. Note that when specifying geometries and boxes on the SQL command line, you must explicitly turn the string representations into geometries function. The 312 is a fictitious spatial reference system that matches our data. So, for example:


```sql
SELECT road_id, road_name
  FROM roads
  WHERE roads_geom='SRID=312;LINESTRING(191232 243118,191108 243242)'::geometry;
```


The above query would return the single record from the "ROADS_GEOM" table in which the geometry was equal to that value.


To check whether some of the roads passes in the area defined by a polygon:


```sql
SELECT road_id, road_name
FROM roads
WHERE ST_Intersects(roads_geom, 'SRID=312;POLYGON((...))');
```


The most common spatial query will probably be a "frame-based" query, used by client software, like data browsers and web mappers, to grab a "map frame" worth of data for display.


When using the "&&" operator, you can specify either a BOX3D as the comparison feature or a GEOMETRY. When you specify a GEOMETRY, however, its bounding box will be used for the comparison.


Using a "BOX3D" object for the frame, such a query looks like this:


```sql

SELECT ST_AsText(roads_geom) AS geom
FROM roads
WHERE
  roads_geom && ST_MakeEnvelope(191232, 243117,191232, 243119,312);
```


Note the use of the SRID 312, to specify the projection of the envelope.
  <a id="pgsql2shp-usage"></a>

## Using the Shapefile Dumper


The `pgsql2shp` table dumper connects to the database and converts a table (possibly defined by a query) into a shape file. The basic syntax is:


```

pgsql2shp [<options>] <database> [<schema>.]<table>
```


```

pgsql2shp [<options>] <database> <query>
```


The commandline options are:


`-f `
:   Write the output to a particular filename.

`-h <host>`
:   The database host to connect to.

`-p <port>`
:   The port to connect to on the database host.

`-P <password>`
:   The password to use when connecting to the database.

`-u <user>`
:   The username to use when connecting to the database.

`-g <geometry column="column">`
:   In the case of tables with multiple geometry columns, the geometry column to use when writing the shape file.

`-b`
:   Use a binary cursor. This will make the operation faster, but will not work if any NON-geometry attribute in the table lacks a cast to text.

`-r`
:   Raw mode. Do not drop the `gid` field, or escape column names.

`-m filename`
:   Remap identifiers to ten character names. The content of the file is lines of two symbols separated by a single white space and no trailing or leading space: VERYLONGSYMBOL SHORTONE ANOTHERVERYLONGSYMBOL SHORTER etc.
