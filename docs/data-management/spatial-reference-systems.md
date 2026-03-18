<a id="spatial_ref_sys"></a>

## Spatial Reference Systems


A [Spatial Reference System](https://en.wikipedia.org/wiki/Spatial_reference_system) (SRS) (also called a Coordinate Reference System (CRS)) defines how geometry is referenced to locations on the Earth's surface. There are three types of SRS:


- A **geodetic** SRS uses angular coordinates (longitude and latitude) which map directly to the surface of the earth.
- A **projected** SRS uses a mathematical projection transformation to "flatten" the surface of the spheroidal earth onto a plane. It assigns location coordinates in a way that allows direct measurement of quantities such as distance, area, and angle. The coordinate system is Cartesian, which means it has a defined origin point and two perpendicular axes (usually oriented North and East). Each projected SRS uses a stated length unit (usually metres or feet). A projected SRS may be limited in its area of applicability to avoid distortion and fit within the defined coordinate bounds.
- A **local** SRS is a Cartesian coordinate system which is not referenced to the earth's surface. In PostGIS this is specified by a SRID value of 0.


 There are many different spatial reference systems in use. Common SRSes are standardized in the European Petroleum Survey Group [EPSG database](http://www.epsg.org/). For convenience PostGIS (and many other spatial systems) refers to SRS definitions using an integer identifier called a SRID.


A geometry is associated with a Spatial Reference System by its SRID value, which is accessed by [ST_SRID](../postgis-reference/spatial-reference-system-functions.md#ST_SRID). The SRID for a geometry can be assigned using [ST_SetSRID](../postgis-reference/spatial-reference-system-functions.md#ST_SetSRID). Some geometry constructor functions allow supplying a SRID (such as [ST_Point](../postgis-reference/geometry-constructors.md#ST_Point) and [ST_MakeEnvelope](../postgis-reference/geometry-constructors.md#ST_MakeEnvelope)). The [EWKT](geometry-data-type.md#EWKB_EWKT) format supports SRIDs with the <code>SRID=n;</code> prefix.


 Spatial functions processing pairs of geometries (such as [overlay](../postgis-reference/overlay-functions.md#Overlay_Functions) and [relationship](../postgis-reference/spatial-relationships.md#Spatial_Relationships) functions) require that the input geometries are in the same spatial reference system (have the same SRID). Geometry data can be transformed into a different spatial reference system using [ST_Transform](../postgis-reference/spatial-reference-system-functions.md#ST_Transform) and [ST_TransformPipeline](../postgis-reference/spatial-reference-system-functions.md#ST_TransformPipeline). Geometry returned from functions has the same SRS as the input geometries.
 <a id="spatial_ref_sys_table"></a>

## SPATIAL_REF_SYS Table


The `SPATIAL_REF_SYS` table used by PostGIS is an OGC-compliant database table that defines the available spatial reference systems. It holds the numeric SRIDs and textual descriptions of the coordinate systems.


The `spatial_ref_sys` table definition is:


```sql
CREATE TABLE spatial_ref_sys (
  srid       INTEGER NOT NULL PRIMARY KEY,
  auth_name  VARCHAR(256),
  auth_srid  INTEGER,
  srtext     VARCHAR(2048),
  proj4text  VARCHAR(2048)
)
```


The columns are:


`srid`
:   An integer code that uniquely identifies the [Spatial Reference System](http://en.wikipedia.org/wiki/SRID) (SRS) within the database.

`auth_name`
:   The name of the standard or standards body that is being cited for this reference system. For example, "EPSG" is a valid `auth_name`.

`auth_srid`
:   The ID of the Spatial Reference System as defined by the Authority cited in the `auth_name`. In the case of EPSG, this is the EPSG code.

`srtext`
:   The Well-Known Text representation of the Spatial Reference System. An example of a WKT SRS representation is:


    ```
    PROJCS["NAD83 / UTM Zone 10N",
      GEOGCS["NAD83",
    	DATUM["North_American_Datum_1983",
    	  SPHEROID["GRS 1980",6378137,298.257222101]
    	],
    	PRIMEM["Greenwich",0],
    	UNIT["degree",0.0174532925199433]
      ],
      PROJECTION["Transverse_Mercator"],
      PARAMETER["latitude_of_origin",0],
      PARAMETER["central_meridian",-123],
      PARAMETER["scale_factor",0.9996],
      PARAMETER["false_easting",500000],
      PARAMETER["false_northing",0],
      UNIT["metre",1]
    ]
    ```


    For a discussion of SRS WKT, see the OGC standard [Well-known text representation of coordinate reference systems](http://docs.opengeospatial.org/is/12-063r5/12-063r5.html).

`proj4text`
:   PostGIS uses the PROJ library to provide coordinate transformation capabilities. The `proj4text` column contains the PROJ coordinate definition string for a particular SRID. For example:


    ```
    +proj=utm +zone=10 +ellps=clrk66 +datum=NAD27 +units=m
    ```


    For more information see the [PROJ web site](https://proj.org/). The `spatial_ref_sys.sql` file contains both `srtext` and `proj4text` definitions for all EPSG projections.


When retrieving spatial reference system definitions for use in transformations, PostGIS uses fhe following strategy:


- If `auth_name` and `auth_srid` are present (non-NULL) use the PROJ SRS based on those entries (if one exists).
- If `srtext` is present create a SRS using it, if possible.
- If `proj4text` is present create a SRS using it, if possible.
  <a id="user-spatial-ref-sys"></a>

## User-Defined Spatial Reference Systems


The PostGIS `spatial_ref_sys` table contains over 3000 of the most common spatial reference system definitions that are handled by the [PROJ](https://proj.org) projection library. But there are many coordinate systems that it does not contain. You can add SRS definitions to the table if you have the required information about the spatial reference system. Or, you can define your own custom spatial reference system if you are familiar with PROJ constructs. Keep in mind that most spatial reference systems are regional and have no meaning when used outside of the bounds they were intended for.


A resource for finding spatial reference systems not defined in the core set is [http://spatialreference.org/](http://spatialreference.org/)


Some commonly used spatial reference systems are: [4326 - WGS 84 Long Lat](http://spatialreference.org/ref/epsg/4326/), [4269 - NAD 83 Long Lat](http://spatialreference.org/ref/epsg/4269/), [3395 - WGS 84 World Mercator](http://spatialreference.org/ref/epsg/3395/), [2163 - US National Atlas Equal Area](http://spatialreference.org/ref/epsg/2163/), and the 60 WGS84 UTM zones. UTM zones are one of the most ideal for measurement, but only cover 6-degree regions. (To determine which UTM zone to use for your area of interest, see the [utmzone PostGIS plpgsql helper function](http://trac.osgeo.org/postgis/wiki/UsersWikiplpgsqlfunctionsDistance).)


 US states use State Plane spatial reference systems (meter or feet based) - usually one or 2 exists per state. Most of the meter-based ones are in the core set, but many of the feet-based ones or ESRI-created ones will need to be copied from [spatialreference.org](http://spatialreference.org).


You can even define non-Earth-based coordinate systems, such as [Mars 2000](http://spatialreference.org/ref/iau2000/mars-2000/) This Mars coordinate system is non-planar (it's in degrees spheroidal), but you can use it with the `geography` type to obtain length and proximity measurements in meters instead of degrees.


Here is an example of loading a custom coordinate system using an unassigned SRID and the PROJ definition for a US-centric Lambert Conformal projection:


```sql

INSERT INTO spatial_ref_sys (srid, proj4text)
VALUES ( 990000,
  '+proj=lcc  +lon_0=-95 +lat_0=25 +lat_1=25 +lat_2=25 +x_0=0 +y_0=0 +datum=WGS84 +units=m +no_defs'
);
```
