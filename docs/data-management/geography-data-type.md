<a id="PostGIS_Geography"></a>

## Geography Data Type


The PostGIS `geography` data type provides native support for spatial features represented on "geographic" coordinates (sometimes called "geodetic" coordinates, or "lat/lon", or "lon/lat"). Geographic coordinates are spherical coordinates expressed in angular units (degrees).


The basis for the PostGIS geometry data type is a plane. The shortest path between two points on the plane is a straight line. That means functions on geometries (areas, distances, lengths, intersections, etc) are calculated using straight line vectors and cartesian mathematics. This makes them simpler to implement and faster to execute, but also makes them inaccurate for data on the spheroidal surface of the earth.


The PostGIS geography data type is based on a spherical model. The shortest path between two points on the sphere is a great circle arc. Functions on geographies (areas, distances, lengths, intersections, etc) are calculated using arcs on the sphere. By taking the spheroidal shape of the world into account, the functions provide more accurate results.


Because the underlying mathematics is more complicated, there are fewer functions defined for the geography type than for the geometry type. Over time, as new algorithms are added the capabilities of the geography type will expand. As a workaround one can convert back and forth between geometry and geography types.


Like the geometry data type, geography data is associated with a spatial reference system via a spatial reference system identifier (SRID). Any geodetic (long/lat based) spatial reference system defined in the `spatial_ref_sys` table can be used. (Prior to PostGIS 2.2, the geography type supported only WGS 84 geodetic (SRID:4326)). You can add your own custom geodetic spatial reference system as described in [User-Defined Spatial Reference Systems](spatial-reference-systems.md#user-spatial-ref-sys).


For all spatial reference systems the units returned by measurement functions (e.g. [ST_Distance](../postgis-reference/measurement-functions.md#ST_Distance), [ST_Length](../postgis-reference/measurement-functions.md#ST_Length), [ST_Perimeter](../postgis-reference/measurement-functions.md#ST_Perimeter), [ST_Area](../postgis-reference/measurement-functions.md#ST_Area)) and for the distance argument of [ST_DWithin](../postgis-reference/spatial-relationships.md#ST_DWithin) are in meters.
 <a id="Create_Geography_Tables"></a>

## Creating Geography Tables


You can create a table to store geography data using the [CREATE TABLE](https://www.postgresql.org/docs/current/sql-createtable.html) SQL statement with a column of type `geography`. The following example creates a table with a geography column storing 2D LineStrings in the WGS84 geodetic coordinate system (SRID 4326):


```sql
CREATE TABLE global_points (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    location geography(POINT,4326)
  );
```


The geography type supports two optional type modifiers:


- the spatial type modifier restricts the kind of shapes and dimensions allowed in the column. Values allowed for the spatial type are: POINT, LINESTRING, POLYGON, MULTIPOINT, MULTILINESTRING, MULTIPOLYGON, GEOMETRYCOLLECTION. The geography type does not support curves, TINS, or POLYHEDRALSURFACEs. The modifier supports coordinate dimensionality restrictions by adding suffixes: Z, M and ZM. For example, a modifier of 'LINESTRINGM' only allows linestrings with three dimensions, and treats the third dimension as a measure. Similarly, 'POINTZM' requires four dimensional (XYZM) data.
- the SRID modifier restricts the spatial reference system SRID to a particular number. If omitted, the SRID defaults to 4326 (WGS84 geodetic), and all calculations are performed using WGS84.


Examples of creating tables with geography columns:


- Create a table with 2D POINT geography with the default SRID 4326 (WGS84 long/lat):

```sql
CREATE TABLE ptgeogwgs(gid serial PRIMARY KEY, geog geography(POINT) );
```
- Create a table with 2D POINT geography in NAD83 longlat:

```sql
CREATE TABLE ptgeognad83(gid serial PRIMARY KEY, geog geography(POINT,4269) );
```
- Create a table with 3D (XYZ) POINTs and an explicit SRID of 4326:

```sql
CREATE TABLE ptzgeogwgs84(gid serial PRIMARY KEY, geog geography(POINTZ,4326) );
```
- Create a table with 2D LINESTRING geography with the default SRID 4326:

```sql
CREATE TABLE lgeog(gid serial PRIMARY KEY, geog geography(LINESTRING) );
```
- Create a table with 2D POLYGON geography with the SRID 4267 (NAD 1927 long lat):

```sql
CREATE TABLE lgeognad27(gid serial PRIMARY KEY, geog geography(POLYGON,4267) );
```


Geography fields are registered in the `geography_columns` system view. You can query the `geography_columns` view and see that the table is listed:


```sql

SELECT * FROM geography_columns;
```


Creating a spatial index works the same as for geometry columns. PostGIS will note that the column type is GEOGRAPHY and create an appropriate sphere-based index instead of the usual planar index used for GEOMETRY.


```
-- Index the test table with a spherical index
CREATE INDEX global_points_gix ON global_points USING GIST ( location );
```
  <a id="Use_Geography_Tables"></a>

## Using Geography Tables


You can insert data into geography tables in the same way as geometry. Geometry data will autocast to the geography type if it has SRID 4326. The [EWKT and EWKB](geometry-data-type.md#EWKB_EWKT) formats can also be used to specify geography values.


```
-- Add some data into the test table
INSERT INTO global_points (name, location) VALUES ('Town', 'SRID=4326;POINT(-110 30)');
INSERT INTO global_points (name, location) VALUES ('Forest', 'SRID=4326;POINT(-109 29)');
INSERT INTO global_points (name, location) VALUES ('London', 'SRID=4326;POINT(0 49)');
```


Any geodetic (long/lat) spatial reference system listed in `spatial_ref_sys` table may be specified as a geography SRID. Non-geodetic coordinate systems raise an error if used.


```
-- NAD 83 lon/lat
SELECT 'SRID=4269;POINT(-123 34)'::geography;
                    geography
----------------------------------------------------
 0101000020AD1000000000000000C05EC00000000000004140
```


```
-- NAD27 lon/lat
SELECT 'SRID=4267;POINT(-123 34)'::geography;
                    geography
----------------------------------------------------
 0101000020AB1000000000000000C05EC00000000000004140
```


```
-- NAD83 UTM zone meters - gives an error since it is a meter-based planar projection
SELECT 'SRID=26910;POINT(-123 34)'::geography;

ERROR:  Only lon/lat coordinate systems are supported in geography.
```


Query and measurement functions use units of meters. So distance parameters should be expressed in meters, and return values should be expected in meters (or square meters for areas).


```
-- A distance query using a 1000km tolerance
SELECT name FROM global_points WHERE ST_DWithin(location, 'SRID=4326;POINT(-110 29)'::geography, 1000000);
```


You can see the power of geography in action by calculating how close a plane flying a great circle route from Seattle to London (LINESTRING(-122.33 47.606, 0.0 51.5)) comes to Reykjavik (POINT(-21.96 64.15)) ([map the route](http://gc.kls2.com/cgi-bin/gc?PATH=SEA-LHR)).


The geography type calculates the true shortest distance of 122.235 km over the sphere between Reykjavik and the great circle flight path between Seattle and London.


```
-- Distance calculation using GEOGRAPHY
SELECT ST_Distance('LINESTRING(-122.33 47.606, 0.0 51.5)'::geography, 'POINT(-21.96 64.15)'::geography);
   st_distance
-----------------
 122235.23815667
```


 The geometry type calculates a meaningless cartesian distance between Reykjavik and the straight line path from Seattle to London plotted on a flat map of the world. The nominal units of the result is "degrees", but the result doesn't correspond to any true angular difference between the points, so even calling them "degrees" is inaccurate.


```
-- Distance calculation using GEOMETRY
SELECT ST_Distance('LINESTRING(-122.33 47.606, 0.0 51.5)'::geometry, 'POINT(-21.96 64.15)'::geometry);
      st_distance
--------------------
 13.342271221453624
```
  <a id="PostGIS_GeographyVSGeometry"></a>

## When to use the Geography data type


The geography data type allows you to store data in longitude/latitude coordinates, but at a cost: there are fewer functions defined on GEOGRAPHY than there are on GEOMETRY; those functions that are defined take more CPU time to execute.


The data type you choose should be determined by the expected working area of the application you are building. Will your data span the globe or a large continental area, or is it local to a state, county or municipality?


- If your data is contained in a small area, you might find that choosing an appropriate projection and using GEOMETRY is the best solution, in terms of performance and functionality available.
- If your data is global or covers a continental region, you may find that GEOGRAPHY allows you to build a system without having to worry about projection details. You store your data in longitude/latitude, and use the functions that have been defined on GEOGRAPHY.
- If you don't understand projections, and you don't want to learn about them, and you're prepared to accept the limitations in functionality available in GEOGRAPHY, then it might be easier for you to use GEOGRAPHY than GEOMETRY. Simply load your data up as longitude/latitude and go from there.


Refer to [PostGIS Function Support Matrix](../postgis-special-functions-index/postgis-function-support-matrix.md#PostGIS_TypeFunctionMatrix) for compare between what is supported for Geography vs. Geometry. For a brief listing and description of Geography functions, refer to [PostGIS Geography Support Functions](../postgis-special-functions-index/postgis-geography-support-functions.md#PostGIS_GeographyFunctions)
  <a id="PostGIS_Geography_AdvancedFAQ"></a>

## Geography Advanced FAQ


**Q: Do you calculate on the sphere or the spheroid?**


 By default, all distance and area calculations are done on the spheroid. You should find that the results of calculations in local areas match up will with local planar results in good local projections. Over larger areas, the spheroidal calculations will be more accurate than any calculation done on a projected plane.


All the geography functions have the option of using a sphere calculation, by setting a final boolean parameter to 'FALSE'. This will somewhat speed up calculations, particularly for cases where the geometries are very simple.


**Q: What about the date-line and the poles?**


 All the calculations have no conception of date-line or poles, the coordinates are spherical (longitude/latitude) so a shape that crosses the dateline is, from a calculation point of view, no different from any other shape.


**Q: What is the longest arc you can process?**


We use great circle arcs as the "interpolation line" between two points. That means any two points are actually joined up two ways, depending on which direction you travel along the great circle. All our code assumes that the points are joined by the *shorter* of the two paths along the great circle. As a consequence, shapes that have arcs of more than 180 degrees will not be correctly modelled.


**Q: Why is it so slow to calculate the area of Europe / Russia / insert big geographic region here ?**


Because the polygon is so darned huge! Big areas are bad for two reasons: their bounds are huge, so the index tends to pull the feature no matter what query you run; the number of vertices is huge, and tests (distance, containment) have to traverse the vertex list at least once and sometimes N times (with N being the number of vertices in the other candidate feature).


As with GEOMETRY, we recommend that when you have very large polygons, but are doing queries in small areas, you "denormalize" your geometric data into smaller chunks so that the index can effectively subquery parts of the object and so queries don't have to pull out the whole object every time. Please consult [ST_Subdivide](../postgis-reference/overlay-functions.md#ST_Subdivide) function documentation. Just because you *can* store all of Europe in one polygon doesn't mean you *should*.
