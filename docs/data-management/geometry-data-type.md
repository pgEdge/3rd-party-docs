<a id="PostGIS_Geometry"></a>

## Geometry Data Type


PostGIS implements the OGC Simple Features model by defining a PostgreSQL data type called `geometry`. It represents all of the geometry subtypes by using an internal type code (see [GeometryType](../postgis-reference/geometry-accessors.md#GeometryType) and [ST_GeometryType](../postgis-reference/geometry-accessors.md#ST_GeometryType)). This allows modelling spatial features as rows of tables defined with a column of type `geometry`.


The `geometry` data type is *opaque*, which means that all access is done via invoking functions on geometry values. Functions allow creating geometry objects, accessing or updating all internal fields, and compute new geometry values. PostGIS supports all the functions specified in the OGC [*Simple feature access - Part 2: SQL option*](https://portal.ogc.org/files/?artifact_id=25354) (SFS) specification, as well many others. See [PostGIS Reference](../postgis-reference/index.md#reference) for the full list of functions.


!!! note

    PostGIS follows the SFA standard by prefixing spatial functions with "ST_". This was intended to stand for "Spatial and Temporal", but the temporal part of the standard was never developed. Instead it can be interpreted as "Spatial Type".


The SFA standard specifies that spatial objects include a Spatial Reference System identifier (SRID). The SRID is required when creating spatial objects for insertion into the database (it may be defaulted to 0). See [ST_SRID](../postgis-reference/spatial-reference-system-functions.md#ST_SRID) and [Spatial Reference Systems](spatial-reference-systems.md#spatial_ref_sys)


To make querying geometry efficient PostGIS defines various kinds of spatial indexes, and spatial operators to use them. See [Spatial Indexes](spatial-indexes.md#build-indexes) and [Using Spatial Indexes](../spatial-queries/using-spatial-indexes.md#using-query-indexes) for details.
 <a id="EWKB_EWKT"></a>

## PostGIS EWKB and EWKT


OGC SFA specifications initially supported only 2D geometries, and the geometry SRID is not included in the input/output representations. The OGC SFA specification 1.2.1 (which aligns with the ISO 19125 standard) adds support for 3D (ZYZ) and measured (XYM and XYZM) coordinates, but still does not include the SRID value.


Because of these limitations PostGIS defined extended EWKB and EWKT formats. They provide 3D (XYZ and XYM) and 4D (XYZM) coordinate support and include SRID information. Including all geometry information allows PostGIS to use EWKB as the format of record (e.g. in DUMP files).


EWKB and EWKT are used for the "canonical forms" of PostGIS data objects. For input, the canonical form for binary data is EWKB, and for text data either EWKB or EWKT is accepted. This allows geometry values to be created by casting a text value in either HEXEWKB or EWKT to a geometry value using `::geometry`. For output, the canonical form for binary is EWKB, and for text it is HEXEWKB (hex-encoded EWKB).


For example this statement creates a geometry by casting from an EWKT text value, and outputs it using the canonical form of HEXEWKB:


```sql
SELECT 'SRID=4;POINT(0 0)'::geometry;
  geometry
  ----------------------------------------------------
  01010000200400000000000000000000000000000000000000
```


PostGIS EWKT output has a few differences to OGC WKT:


- For 3DZ geometries the Z qualifier is omitted:

  OGC: POINT Z (1 2 3)

  EWKT: POINT (1 2 3)
- For 3DM geometries the M qualifier is included:

  OGC: POINT M (1 2 3)

  EWKT: POINTM (1 2 3)
- For 4D geometries the ZM qualifier is omitted:

  OGC: POINT ZM (1 2 3 4)

  EWKT: POINT (1 2 3 4)


EWKT avoids over-specifying dimensionality and the inconsistencies that can occur with the OGC/ISO format, such as:


- POINT ZM (1 1)
- POINT ZM (1 1 1)
- POINT (1 1 1 1)


!!! caution

    PostGIS extended formats are currently a superset of the OGC ones, so that every valid OGC WKB/WKT is also valid EWKB/EWKT. However, this might vary in the future, if the OGC extends a format in a way that conflicts with the PosGIS definition. Thus you SHOULD NOT rely on this compatibility!


Examples of the EWKT text representation of spatial objects are:


- POINT(0 0 0) -- XYZ
- SRID=32632;POINT(0 0) -- XY with SRID
- POINTM(0 0 0) -- XYM
- POINT(0 0 0 0) -- XYZM
- SRID=4326;MULTIPOINTM(0 0 0,1 2 1) -- XYM with SRID
- MULTILINESTRING((0 0 0,1 1 0,1 2 1),(2 3 1,3 2 1,5 4 1))
- POLYGON((0 0 0,4 0 0,4 4 0,0 4 0,0 0 0),(1 1 0,2 1 0,2 2 0,1 2 0,1 1 0))
- MULTIPOLYGON(((0 0 0,4 0 0,4 4 0,0 4 0,0 0 0),(1 1 0,2 1 0,2 2 0,1 2 0,1 1 0)),((-1 -1 0,-1 -2 0,-2 -2 0,-2 -1 0,-1 -1 0)))
- GEOMETRYCOLLECTIONM( POINTM(2 3 9), LINESTRINGM(2 3 4, 3 4 5) )
- MULTICURVE( (0 0, 5 5), CIRCULARSTRING(4 0, 4 4, 8 4) )
- POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)), ((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)), ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)), ((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)), ((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)), ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )
- TRIANGLE ((0 0, 0 10, 10 0, 0 0))
- TIN( ((0 0 0, 0 0 1, 0 1 0, 0 0 0)), ((0 0 0, 0 1 0, 1 1 0, 0 0 0)) )


Input and output using these formats is available using the following functions:


```
bytea EWKB = ST_AsEWKB(geometry);
text EWKT = ST_AsEWKT(geometry);
geometry = ST_GeomFromEWKB(bytea EWKB);
geometry = ST_GeomFromEWKT(text EWKT);
```


For example, a statement to create and insert a PostGIS spatial object using EWKT is:


```sql
INSERT INTO geotable ( geom, name )
  VALUES ( ST_GeomFromEWKT('SRID=312;POINTM(-126.4 45.32 15)'), 'A Place' )
```
