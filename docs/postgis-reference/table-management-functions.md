<a id="Management_Functions"></a>

## Table Management Functions
  <a id="AddGeometryColumn"></a>

# AddGeometryColumn

Adds a geometry column to an existing table.

## Synopsis


```sql
text AddGeometryColumn(varchar
			table_name, varchar
			column_name, integer
			srid, varchar
			type, integer
			dimension, boolean
			use_typmod=true)
text AddGeometryColumn(varchar
			schema_name, varchar
			table_name, varchar
			column_name, integer
			srid, varchar
			type, integer
			dimension, boolean
			use_typmod=true)
text AddGeometryColumn(varchar
			catalog_name, varchar
			schema_name, varchar
			table_name, varchar
			column_name, integer
			srid, varchar
			type, integer
			dimension, boolean
			use_typmod=true)
```


## Description


Adds a geometry column to an existing table of attributes. The `schema_name` is the name of the table schema. The `srid` must be an integer value reference to an entry in the SPATIAL_REF_SYS table. The `type` must be a string corresponding to the geometry type, eg, 'POLYGON' or 'MULTILINESTRING' . An error is thrown if the schemaname doesn't exist (or not visible in the current search_path) or the specified SRID, geometry type, or dimension is invalid.


!!! note

    Changed: 2.0.0 This function no longer updates geometry_columns since geometry_columns is a view that reads from system catalogs. It by default also does not create constraints, but instead uses the built in type modifier behavior of PostgreSQL. So for example building a wgs84 POINT column with this function is now equivalent to: <code>ALTER TABLE some_table ADD COLUMN geom geometry(Point,4326);</code>


    Changed: 2.0.0 If you require the old behavior of constraints use the default `use_typmod`, but set it to false.


!!! note

    Changed: 2.0.0 Views can no longer be manually registered in geometry_columns, however views built against geometry typmod tables geometries and used without wrapper functions will register themselves correctly because they inherit the typmod behavior of their parent table column. Views that use geometry functions that output other geometries will need to be cast to typmod geometries for these view geometry columns to be registered correctly in geometry_columns. Refer to [Manually Registering Geometry Columns](../data-management/spatial-tables.md#Manual_Register_Spatial_Column).


Enhanced: 2.0.0 use_typmod argument introduced. Defaults to creating typmod geometry column instead of constraint-based.


## Examples


```
-- Create schema to hold data
CREATE SCHEMA my_schema;
-- Create a new simple PostgreSQL table
CREATE TABLE my_schema.my_spatial_table (id serial);

-- Describing the table shows a simple table with a single "id" column.
postgis=# \d my_schema.my_spatial_table
							 Table "my_schema.my_spatial_table"
 Column |  Type   |                                Modifiers
--------+---------+-------------------------------------------------------------------------
 id     | integer | not null default nextval('my_schema.my_spatial_table_id_seq'::regclass)

-- Add a spatial column to the table
SELECT AddGeometryColumn ('my_schema','my_spatial_table','geom',4326,'POINT',2);

-- Add a point using the old constraint based behavior
SELECT AddGeometryColumn ('my_schema','my_spatial_table','geom_c',4326,'POINT',2, false);

--Add a curvepolygon using old constraint behavior
SELECT AddGeometryColumn ('my_schema','my_spatial_table','geomcp_c',4326,'CURVEPOLYGON',2, false);

-- Describe the table again reveals the addition of a new geometry columns.
\d my_schema.my_spatial_table
                            addgeometrycolumn
-------------------------------------------------------------------------
 my_schema.my_spatial_table.geomcp_c SRID:4326 TYPE:CURVEPOLYGON DIMS:2
(1 row)

                                    Table "my_schema.my_spatial_table"
  Column  |         Type         |                                Modifiers
----------+----------------------+-------------------------------------------------------------------------
 id       | integer              | not null default nextval('my_schema.my_spatial_table_id_seq'::regclass)
 geom     | geometry(Point,4326) |
 geom_c   | geometry             |
 geomcp_c | geometry             |
Check constraints:
    "enforce_dims_geom_c" CHECK (st_ndims(geom_c) = 2)
    "enforce_dims_geomcp_c" CHECK (st_ndims(geomcp_c) = 2)
    "enforce_geotype_geom_c" CHECK (geometrytype(geom_c) = 'POINT'::text OR geom_c IS NULL)
    "enforce_geotype_geomcp_c" CHECK (geometrytype(geomcp_c) = 'CURVEPOLYGON'::text OR geomcp_c IS NULL)
    "enforce_srid_geom_c" CHECK (st_srid(geom_c) = 4326)
    "enforce_srid_geomcp_c" CHECK (st_srid(geomcp_c) = 4326)

-- geometry_columns view also registers the new columns --
SELECT f_geometry_column As col_name, type, srid, coord_dimension As ndims
    FROM geometry_columns
    WHERE f_table_name = 'my_spatial_table' AND f_table_schema = 'my_schema';

 col_name |     type     | srid | ndims
----------+--------------+------+-------
 geom     | Point        | 4326 |     2
 geom_c   | Point        | 4326 |     2
 geomcp_c | CurvePolygon | 4326 |     2
```


## See Also


[DropGeometryColumn](#DropGeometryColumn), [DropGeometryTable](#DropGeometryTable), [GEOMETRY_COLUMNS View](../data-management/spatial-tables.md#geometry_columns), [Manually Registering Geometry Columns](../data-management/spatial-tables.md#Manual_Register_Spatial_Column)
  <a id="DropGeometryColumn"></a>

# DropGeometryColumn

Removes a geometry column from a spatial table.

## Synopsis


```sql
text DropGeometryColumn(varchar
			table_name, varchar
			column_name)
text DropGeometryColumn(varchar
			schema_name, varchar
			table_name, varchar
			column_name)
text DropGeometryColumn(varchar
			catalog_name, varchar
			schema_name, varchar
			table_name, varchar
			column_name)
```


## Description


Removes a geometry column from a spatial table. Note that schema_name will need to match the f_table_schema field of the table's row in the geometry_columns table.


!!! note

    Changed: 2.0.0 This function is provided for backward compatibility. Now that since geometry_columns is now a view against the system catalogs, you can drop a geometry column like any other table column using <code>ALTER TABLE</code>


## Examples


```sql

			SELECT DropGeometryColumn ('my_schema','my_spatial_table','geom');
			----RESULT output ---
			                  dropgeometrycolumn
------------------------------------------------------
 my_schema.my_spatial_table.geom effectively removed.

-- In PostGIS 2.0+ the above is also equivalent to the standard
-- the standard alter table.  Both will deregister from geometry_columns
ALTER TABLE my_schema.my_spatial_table DROP column geom;

```


## See Also


[AddGeometryColumn](#AddGeometryColumn), [DropGeometryTable](#DropGeometryTable), [GEOMETRY_COLUMNS View](../data-management/spatial-tables.md#geometry_columns)
  <a id="DropGeometryTable"></a>

# DropGeometryTable

Drops a table and all its references in geometry_columns.

## Synopsis


```sql
boolean DropGeometryTable(varchar
			table_name)
boolean DropGeometryTable(varchar
			schema_name, varchar
			table_name)
boolean DropGeometryTable(varchar
			catalog_name, varchar
			schema_name, varchar
			table_name)
```


## Description


Drops a table and all its references in geometry_columns. Note: uses current_schema() on schema-aware pgsql installations if schema is not provided.


!!! note

    Changed: 2.0.0 This function is provided for backward compatibility. Now that since geometry_columns is now a view against the system catalogs, you can drop a table with geometry columns like any other table using <code>DROP TABLE</code>


## Examples


```sql
SELECT DropGeometryTable ('my_schema','my_spatial_table');
----RESULT output ---
my_schema.my_spatial_table dropped.

-- The above is now equivalent to --
DROP TABLE my_schema.my_spatial_table;

```


## See Also


[AddGeometryColumn](#AddGeometryColumn), [DropGeometryColumn](#DropGeometryColumn), [GEOMETRY_COLUMNS View](../data-management/spatial-tables.md#geometry_columns)
  <a id="Find_SRID"></a>

# Find_SRID

Returns the SRID defined for a geometry column.

## Synopsis


```sql
integer Find_SRID(varchar  a_schema_name, varchar  a_table_name, varchar  a_geomfield_name)
```


## Description


Returns the integer SRID of the specified geometry column by searching through the GEOMETRY_COLUMNS table. If the geometry column has not been properly added (e.g. with the [AddGeometryColumn](#AddGeometryColumn) function), this function will not work.


## Examples


```sql
 SELECT Find_SRID('public', 'tiger_us_state_2007', 'geom_4269');
find_srid
----------
4269
```


## See Also


[ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="Populate_Geometry_Columns"></a>

# Populate_Geometry_Columns

Ensures geometry columns are defined with type modifiers or have appropriate spatial constraints.

## Synopsis


```sql
text Populate_Geometry_Columns(boolean  use_typmod=true)
int Populate_Geometry_Columns(oid relation_oid, boolean  use_typmod=true)
```


## Description


Ensures geometry columns have appropriate type modifiers or spatial constraints to ensure they are registered correctly in the `geometry_columns` view. By default will convert all geometry columns with no type modifier to ones with type modifiers.


For backwards compatibility and for spatial needs such as table inheritance where each child table may have different geometry type, the old check constraint behavior is still supported. If you need the old behavior, you need to pass in the new optional argument as false `use_typmod=false`. When this is done geometry columns will be created with no type modifiers but will have 3 constraints defined. In particular, this means that every geometry column belonging to a table has at least three constraints:


- `enforce_dims_geom` - ensures every geometry has the same dimension (see [ST_NDims](geometry-accessors.md#ST_NDims))
- `enforce_geotype_geom` - ensures every geometry is of the same type (see [GeometryType](geometry-accessors.md#GeometryType))
- `enforce_srid_geom` - ensures every geometry is in the same projection (see [ST_SRID](spatial-reference-system-functions.md#ST_SRID))


If a table `oid` is provided, this function tries to determine the srid, dimension, and geometry type of all geometry columns in the table, adding constraints as necessary. If successful, an appropriate row is inserted into the geometry_columns table, otherwise, the exception is caught and an error notice is raised describing the problem.


If the `oid` of a view is provided, as with a table oid, this function tries to determine the srid, dimension, and type of all the geometries in the view, inserting appropriate entries into the `geometry_columns` table, but nothing is done to enforce constraints.


The parameterless variant is a simple wrapper for the parameterized variant that first truncates and repopulates the geometry_columns table for every spatial table and view in the database, adding spatial constraints to tables where appropriate. It returns a summary of the number of geometry columns detected in the database and the number that were inserted into the `geometry_columns` table. The parameterized version simply returns the number of rows inserted into the `geometry_columns` table.


Availability: 1.4.0


Changed: 2.0.0 By default, now uses type modifiers instead of check constraints to constrain geometry types. You can still use check constraint behavior instead by using the new `use_typmod` and setting it to false.


Enhanced: 2.0.0 `use_typmod` optional argument was introduced that allows controlling if columns are created with typmodifiers or with check constraints.


## Examples


```sql

CREATE TABLE public.myspatial_table(gid serial, geom geometry);
INSERT INTO myspatial_table(geom) VALUES(ST_GeomFromText('LINESTRING(1 2, 3 4)',4326) );
-- This will now use typ modifiers.  For this to work, there must exist data
SELECT Populate_Geometry_Columns('public.myspatial_table'::regclass);

populate_geometry_columns
--------------------------
                        1


\d myspatial_table

                                   Table "public.myspatial_table"
 Column |           Type            |                           Modifiers
--------+---------------------------+---------------------------------------------------------------
 gid    | integer                   | not null default nextval('myspatial_table_gid_seq'::regclass)
 geom   | geometry(LineString,4326) |
```


```
-- This will change the geometry columns to use constraints if they are not typmod or have constraints already.
--For this to work, there must exist data
CREATE TABLE public.myspatial_table_cs(gid serial, geom geometry);
INSERT INTO myspatial_table_cs(geom) VALUES(ST_GeomFromText('LINESTRING(1 2, 3 4)',4326) );
SELECT Populate_Geometry_Columns('public.myspatial_table_cs'::regclass, false);
populate_geometry_columns
--------------------------
                        1
\d myspatial_table_cs

                          Table "public.myspatial_table_cs"
 Column |   Type   |                            Modifiers
--------+----------+------------------------------------------------------------------
 gid    | integer  | not null default nextval('myspatial_table_cs_gid_seq'::regclass)
 geom   | geometry |
Check constraints:
    "enforce_dims_geom" CHECK (st_ndims(geom) = 2)
    "enforce_geotype_geom" CHECK (geometrytype(geom) = 'LINESTRING'::text OR geom IS NULL)
    "enforce_srid_geom" CHECK (st_srid(geom) = 4326)
```
  <a id="UpdateGeometrySRID"></a>

# UpdateGeometrySRID

Updates the SRID of all features in a geometry column, and the table metadata.

## Synopsis


```sql
text UpdateGeometrySRID(varchar
			table_name, varchar
			column_name, integer
			srid)
text UpdateGeometrySRID(varchar
			schema_name, varchar
			table_name, varchar
			column_name, integer
			srid)
text UpdateGeometrySRID(varchar
			catalog_name, varchar
			schema_name, varchar
			table_name, varchar
			column_name, integer
			srid)
```


## Description


Updates the SRID of all features in a geometry column, updating constraints and reference in geometry_columns. If the column was enforced by a type definition, the type definition will be changed. Note: uses current_schema() on schema-aware pgsql installations if schema is not provided.


## Examples


Insert geometries into roads table with a SRID set already using [EWKT format](geometry-input.md#ST_GeomFromEWKT):


```
COPY roads (geom) FROM STDIN;
SRID=4326;LINESTRING(0 0, 10 10)
SRID=4326;LINESTRING(10 10, 15 0)
\.

```


This will change the srid of the roads table to 4326 from whatever it was before:


```sql
SELECT UpdateGeometrySRID('roads','geom',4326);
```


The prior example is equivalent to this DDL statement:


```sql
ALTER TABLE roads
  ALTER COLUMN geom TYPE geometry(MULTILINESTRING, 4326)
    USING ST_SetSRID(geom,4326);
```


If you got the projection wrong (or brought it in as unknown) in load and you wanted to transform to web mercator all in one shot you can do this with DDL but there is no equivalent PostGIS management function to do so in one go.


```sql
ALTER TABLE roads
 ALTER COLUMN geom TYPE geometry(MULTILINESTRING, 3857) USING ST_Transform(ST_SetSRID(geom,4326),3857) ;
```


## See Also


 [RT_UpdateRasterSRID](../raster-reference/raster-management.md#RT_UpdateRasterSRID), [ST_SetSRID](spatial-reference-system-functions.md#ST_SetSRID), [ST_Transform](spatial-reference-system-functions.md#ST_Transform), [ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT)
