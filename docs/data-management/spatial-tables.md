## Spatial Tables
  <a id="Create_Spatial_Table"></a>

## Creating a Spatial Table


You can create a table to store geometry data using the [CREATE TABLE](https://www.postgresql.org/docs/current/sql-createtable.html) SQL statement with a column of type `geometry`. The following example creates a table with a geometry column storing 2D (XY) LineStrings in the BC-Albers coordinate system (SRID 3005):


```sql
CREATE TABLE roads (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    geom geometry(LINESTRING,3005)
  );
```


The `geometry` type supports two optional **type modifiers**:


- the **spatial type modifier** restricts the kind of shapes and dimensions allowed in the column. The value can be any of the supported [geometry subtypes](spatial-data-model.md#RefObject) (e.g. POINT, LINESTRING, POLYGON, MULTIPOINT, MULTILINESTRING, MULTIPOLYGON, GEOMETRYCOLLECTION, etc). The modifier supports coordinate dimensionality restrictions by adding suffixes: Z, M and ZM. For example, a modifier of 'LINESTRINGM' allows only linestrings with three dimensions, and treats the third dimension as a measure. Similarly, 'POINTZM' requires four dimensional (XYZM) data.
- the **SRID modifier** restricts the [spatial reference system](spatial-reference-systems.md#spatial_ref_sys) SRID to a particular number. If omitted, the SRID defaults to 0.


Examples of creating tables with geometry columns:


- Create a table holding any kind of geometry with the default SRID:

```sql
CREATE TABLE geoms(gid serial PRIMARY KEY, geom geometry );
```
- Create a table with 2D POINT geometry with the default SRID:

```sql
CREATE TABLE pts(gid serial PRIMARY KEY, geom geometry(POINT) );
```
- Create a table with 3D (XYZ) POINTs and an explicit SRID of 3005:

```sql
CREATE TABLE pts(gid serial PRIMARY KEY, geom geometry(POINTZ,3005) );
```
- Create a table with 4D (XYZM) LINESTRING geometry with the default SRID:

```sql
CREATE TABLE lines(gid serial PRIMARY KEY, geom geometry(LINESTRINGZM) );
```
- Create a table with 2D POLYGON geometry with the SRID 4267 (NAD 1927 long lat):

```sql
CREATE TABLE polys(gid serial PRIMARY KEY, geom geometry(POLYGON,4267) );
```


It is possible to have more than one geometry column in a table. This can be specified when the table is created, or a column can be added using the [ALTER TABLE](https://www.postgresql.org/docs/current/sql-altertable.html) SQL statement. This example adds a column that can hold 3D LineStrings:


```sql
ALTER TABLE roads ADD COLUMN geom2 geometry(LINESTRINGZ,4326);
```
  <a id="geometry_columns"></a>

## GEOMETRY_COLUMNS View


The OGC *Simple Features Specification for SQL* defines the `GEOMETRY_COLUMNS` metadata table to describe geometry table structure. In PostGIS `geometry_columns` is a view reading from database system catalog tables. This ensures that the spatial metadata information is always consistent with the currently defined tables and views. The view structure is:


```
\d geometry_columns
```


```
             View "public.geometry_columns"
      Column       |          Type          | Modifiers
-------------------+------------------------+-----------
 f_table_catalog   | character varying(256) |
 f_table_schema    | character varying(256) |
 f_table_name      | character varying(256) |
 f_geometry_column | character varying(256) |
 coord_dimension   | integer                |
 srid              | integer                |
 type              | character varying(30)  |
```


The columns are:


`f_table_catalog, f_table_schema, f_table_name`
:   The fully qualified name of the feature table containing the geometry column. There is no PostgreSQL analogue of "catalog" so that column is left blank. For "schema" the PostgreSQL schema name is used (`public` is the default).

`f_geometry_column`
:   The name of the geometry column in the feature table.

`coord_dimension`
:   The coordinate dimension (2, 3 or 4) of the column.

`srid`
:   The ID of the spatial reference system used for the coordinate geometry in this table. It is a foreign key reference to the `spatial_ref_sys` table (see [SPATIAL_REF_SYS Table](spatial-reference-systems.md#spatial_ref_sys_table)).

`type`
:   The type of the spatial object. To restrict the spatial column to a single type, use one of: POINT, LINESTRING, POLYGON, MULTIPOINT, MULTILINESTRING, MULTIPOLYGON, GEOMETRYCOLLECTION or corresponding XYM versions POINTM, LINESTRINGM, POLYGONM, MULTIPOINTM, MULTILINESTRINGM, MULTIPOLYGONM, GEOMETRYCOLLECTIONM. For heterogeneous (mixed-type) collections, you can use "GEOMETRY" as the type.
  <a id="Manual_Register_Spatial_Column"></a>

## Manually Registering Geometry Columns


Two of the cases where you may need this are the case of SQL Views and bulk inserts. For bulk insert case, you can correct the registration in the geometry_columns table by constraining the column or doing an alter table. For views, you could expose using a CAST operation. Note, if your column is typmod based, the creation process would register it correctly, so no need to do anything. Also views that have no spatial function applied to the geometry will register the same as the underlying table geometry column.


```
-- Lets say you have a view created like this
CREATE VIEW public.vwmytablemercator AS
	SELECT gid, ST_Transform(geom, 3395) As geom, f_name
	FROM public.mytable;

-- For it to register correctly
-- You need to cast the geometry
--
DROP VIEW public.vwmytablemercator;
CREATE VIEW  public.vwmytablemercator AS
	SELECT gid, ST_Transform(geom, 3395)::geometry(Geometry, 3395) As geom, f_name
	FROM public.mytable;

-- If you know the geometry type for sure is a 2D POLYGON then you could do
DROP VIEW public.vwmytablemercator;
CREATE VIEW  public.vwmytablemercator AS
	SELECT gid, ST_Transform(geom,3395)::geometry(Polygon, 3395) As geom, f_name
	FROM public.mytable;
```


```
--Lets say you created a derivative table by doing a bulk insert
SELECT poi.gid, poi.geom, citybounds.city_name
INTO myschema.my_special_pois
FROM poi INNER JOIN citybounds ON ST_Intersects(citybounds.geom, poi.geom);

-- Create 2D index on new table
CREATE INDEX idx_myschema_myspecialpois_geom_gist
  ON myschema.my_special_pois USING gist(geom);

-- If your points are 3D points or 3M points,
-- then you might want to create an nd index instead of a 2D index
CREATE INDEX my_special_pois_geom_gist_nd
	ON my_special_pois USING gist(geom gist_geometry_ops_nd);

-- To manually register this new table's geometry column in geometry_columns.
-- Note it will also change the underlying structure of the table to
-- to make the column typmod based.
SELECT populate_geometry_columns('myschema.my_special_pois'::regclass);

-- If you are using PostGIS 2.0 and for whatever reason, you
-- you need the constraint based definition behavior
-- (such as case of inherited tables where all children do not have the same type and srid)
-- set optional use_typmod argument to false
SELECT populate_geometry_columns('myschema.my_special_pois'::regclass, false);
```


Although the old-constraint based method is still supported, a constraint-based geometry column used directly in a view, will not register correctly in geometry_columns, as will a typmod one. In this example we define a column using typmod and another using constraints.


```sql
CREATE TABLE pois_ny(gid SERIAL PRIMARY KEY, poi_name text, cat text, geom geometry(POINT,4326));
SELECT AddGeometryColumn('pois_ny', 'geom_2160', 2160, 'POINT', 2, false);
```


If we run in psql


```
\d pois_ny;
```


We observe they are defined differently -- one is typmod, one is constraint


```
                                  Table "public.pois_ny"
  Column   |         Type          |                       Modifiers

-----------+-----------------------+------------------------------------------------------
 gid       | integer               | not null default nextval('pois_ny_gid_seq'::regclass)
 poi_name  | text                  |
 cat       | character varying(20) |
 geom      | geometry(Point,4326)  |
 geom_2160 | geometry              |
Indexes:
    "pois_ny_pkey" PRIMARY KEY, btree (gid)
Check constraints:
    "enforce_dims_geom_2160" CHECK (st_ndims(geom_2160) = 2)
    "enforce_geotype_geom_2160" CHECK (geometrytype(geom_2160) = 'POINT'::text
        OR geom_2160 IS NULL)
    "enforce_srid_geom_2160" CHECK (st_srid(geom_2160) = 2160)
```


In geometry_columns, they both register correctly


```sql
SELECT f_table_name, f_geometry_column, srid, type
	FROM geometry_columns
	WHERE f_table_name = 'pois_ny';
```


```
f_table_name | f_geometry_column | srid | type
-------------+-------------------+------+-------
pois_ny      | geom              | 4326 | POINT
pois_ny      | geom_2160         | 2160 | POINT
```


However -- if we were to create a view like this


```sql
CREATE VIEW vw_pois_ny_parks AS
SELECT *
  FROM pois_ny
  WHERE cat='park';

SELECT f_table_name, f_geometry_column, srid, type
	FROM geometry_columns
	WHERE f_table_name = 'vw_pois_ny_parks';
```


The typmod based geom view column registers correctly, but the constraint based one does not.


```
   f_table_name   | f_geometry_column | srid |   type
------------------+-------------------+------+----------
 vw_pois_ny_parks | geom              | 4326 | POINT
 vw_pois_ny_parks | geom_2160         |    0 | GEOMETRY
```


This may change in future versions of PostGIS, but for now to force the constraint-based view column to register correctly, you need to do this:


```sql
DROP VIEW vw_pois_ny_parks;
CREATE VIEW vw_pois_ny_parks AS
SELECT gid, poi_name, cat,
  geom,
  geom_2160::geometry(POINT,2160) As geom_2160
  FROM pois_ny
  WHERE cat = 'park';
SELECT f_table_name, f_geometry_column, srid, type
	FROM geometry_columns
	WHERE f_table_name = 'vw_pois_ny_parks';
```


```
   f_table_name   | f_geometry_column | srid | type
------------------+-------------------+------+-------
 vw_pois_ny_parks | geom              | 4326 | POINT
 vw_pois_ny_parks | geom_2160         | 2160 | POINT
```
