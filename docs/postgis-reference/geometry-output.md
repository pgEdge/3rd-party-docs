<a id="Geometry_Outputs"></a>

## Geometry Output


## Well-Known Text (WKT)
  <a id="ST_AsEWKT"></a>

# ST_AsEWKT

Return the Well-Known Text (WKT) representation of the geometry with SRID meta data.

## Synopsis


```sql
text ST_AsEWKT(geometry  g1)
text ST_AsEWKT(geometry  g1, integer  maxdecimaldigits=15)
text ST_AsEWKT(geography  g1)
text ST_AsEWKT(geography  g1, integer  maxdecimaldigits=15)
```


## Description


Returns the Well-Known Text representation of the geometry prefixed with the SRID. The optional `maxdecimaldigits` argument may be used to reduce the maximum number of decimal digits after floating point used in output (defaults to 15).


To perform the inverse conversion of EWKT representation to PostGIS geometry use [ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT).


!!! warning

    Using the `maxdecimaldigits` parameter can cause output geometry to become invalid. To avoid this use [ST_ReducePrecision](geometry-processing.md#ST_ReducePrecision) with a suitable gridsize first.


!!! note

    The WKT spec does not include the SRID. To get the OGC WKT format use [ST_AsText](#ST_AsText).


!!! warning

    WKT format does not maintain precision so to prevent floating truncation, use [ST_AsBinary](#ST_AsBinary) or [ST_AsEWKB](#ST_AsEWKB) format for transport.


Enhanced: 3.1.0 support for optional precision parameter.


Enhanced: 2.0.0 support for Geography, Polyhedral surfaces, Triangles and TIN was introduced.


## Examples


```sql
SELECT ST_AsEWKT('0103000020E61000000100000005000000000000
      000000000000000000000000000000000000000000000000000000
      F03F000000000000F03F000000000000F03F000000000000F03
      F000000000000000000000000000000000000000000000000'::geometry);

       st_asewkt
--------------------------------
SRID=4326;POLYGON((0 0,0 1,1 1,1 0,0 0))
(1 row)

SELECT ST_AsEWKT('0108000080030000000000000060E30A4100000000785C0241000000000000F03F0000000018
E20A4100000000485F024100000000000000400000000018
E20A4100000000305C02410000000000000840')

--st_asewkt---
CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)
```


## See Also


 [ST_AsBinary](#ST_AsBinary), [ST_AsEWKB](#ST_AsEWKB), [ST_AsText](#ST_AsText), [ST_GeomFromEWKT](geometry-input.md#ST_GeomFromEWKT)
  <a id="ST_AsText"></a>

# ST_AsText

Return the Well-Known Text (WKT) representation of the geometry/geography without SRID metadata.

## Synopsis


```sql
text ST_AsText(geometry  g1)
text ST_AsText(geometry  g1, integer  maxdecimaldigits = 15)
text ST_AsText(geography  g1)
text ST_AsText(geography  g1, integer  maxdecimaldigits = 15)
```


## Description


Returns the OGC [Well-Known Text](../data-management/spatial-data-model.md#OpenGISWKBWKT) (WKT) representation of the geometry/geography. The optional `maxdecimaldigits` argument may be used to limit the number of digits after the decimal point in output ordinates (defaults to 15).


To perform the inverse conversion of WKT representation to PostGIS geometry use [ST_GeomFromText](geometry-input.md#ST_GeomFromText).


!!! note

    The standard OGC WKT representation does not include the SRID. To include the SRID as part of the output representation, use the non-standard PostGIS function [ST_AsEWKT](#ST_AsEWKT)


!!! warning

    The textual representation of numbers in WKT may not maintain full floating-point precision. To ensure full accuracy for data storage or transport it is best to use [Well-Known Binary](../data-management/spatial-data-model.md#OpenGISWKBWKT) (WKB) format (see [ST_AsBinary](#ST_AsBinary) and `maxdecimaldigits`).


!!! warning

    Using the `maxdecimaldigits` parameter can cause output geometry to become invalid. To avoid this use [ST_ReducePrecision](geometry-processing.md#ST_ReducePrecision) with a suitable gridsize first.


Availability: 1.5 - support for geography was introduced.


Enhanced: 2.5 - optional parameter precision introduced.


 s2.1.1.1


 SQL-MM 3: 5.1.25


## Examples


```sql
SELECT ST_AsText('01030000000100000005000000000000000000
000000000000000000000000000000000000000000000000
F03F000000000000F03F000000000000F03F000000000000F03
F000000000000000000000000000000000000000000000000');

    st_astext
--------------------------------
 POLYGON((0 0,0 1,1 1,1 0,0 0))
```


Full precision output is the default.


```sql
SELECT ST_AsText('POINT(111.1111111 1.1111111)'));
    st_astext
------------------------------
 POINT(111.1111111 1.1111111)
```


The `maxdecimaldigits` argument can be used to limit output precision.


```sql
SELECT ST_AsText('POINT(111.1111111 1.1111111)'), 2);
    st_astext
--------------------
 POINT(111.11 1.11)
```


## See Also


[ST_AsBinary](#ST_AsBinary), [ST_AsEWKB](#ST_AsEWKB), [ST_AsEWKT](#ST_AsEWKT), [ST_GeomFromText](geometry-input.md#ST_GeomFromText)


## Well-Known Binary (WKB)
  <a id="ST_AsBinary"></a>

# ST_AsBinary

Return the OGC/ISO Well-Known Binary (WKB) representation of the geometry/geography without SRID meta data.

## Synopsis


```sql
bytea ST_AsBinary(geometry  g1)
bytea ST_AsBinary(geometry  g1, text NDR_or_XDR)
bytea ST_AsBinary(geography  g1)
bytea ST_AsBinary(geography  g1, text NDR_or_XDR)
```


## Description


Returns the OGC/ISO [Well-Known Binary](../data-management/spatial-data-model.md#OpenGISWKBWKT) (WKB) representation of the geometry. The first function variant defaults to encoding using server machine endian. The second function variant takes a text argument specifying the endian encoding: either 'NDR' for little-endian; or 'XDR' for big-endian. Supplying unknown arguments will result in little-endian output.


WKB format is useful to read geometry data from the database and maintaining full numeric precision. This avoids the precision rounding that can happen with text formats such as WKT.


To perform the inverse conversion of WKB to PostGIS geometry use [ST_GeomFromWKB](geometry-input.md#ST_GeomFromWKB).


!!! note

    The OGC/ISO WKB format does not include the SRID. To get the EWKB format which does include the SRID use [ST_AsEWKB](#ST_AsEWKB)


!!! note

    The default behavior in PostgreSQL 9.0 has been changed to output bytea in hex encoding. If your GUI tools require the old behavior, then SET bytea_output='escape' in your database.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Enhanced: 2.0.0 support for higher coordinate dimensions was introduced.


Enhanced: 2.0.0 support for specifying endian with geography was introduced.


Availability: 1.5.0 geography support was introduced.


Changed: 2.0.0 Inputs to this function can not be unknown -- must be geometry. Constructs such as <code>ST_AsBinary('POINT(1 2)')</code> are no longer valid and you will get an <code>n st_asbinary(unknown) is not unique error</code>. Code like that needs to be changed to <code>ST_AsBinary('POINT(1 2)'::geometry);</code>. If that is not possible, then install `legacy.sql`.


 s2.1.1.1


 SQL-MM 3: 5.1.37


## Examples


```sql
SELECT ST_AsBinary(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326));

       st_asbinary
--------------------------------
\x01030000000100000005000000000000000000000000000000000000000000000000000000000000
000000f03f000000000000f03f000000000000f03f000000000000f03f0000000000000000000000
00000000000000000000000000
```


```sql
SELECT ST_AsBinary(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326), 'XDR');
       st_asbinary
--------------------------------
\x000000000300000001000000050000000000000000000000000000000000000000000000003ff000
00000000003ff00000000000003ff00000000000003ff00000000000000000000000000000000000
00000000000000000000000000
```


## See Also


 [ST_GeomFromWKB](geometry-input.md#ST_GeomFromWKB), [ST_AsEWKB](#ST_AsEWKB), [ST_AsTWKB](#ST_AsTWKB), [ST_AsText](#ST_AsText),
  <a id="ST_AsEWKB"></a>

# ST_AsEWKB

Return the Extended Well-Known Binary (EWKB) representation of the geometry with SRID meta data.

## Synopsis


```sql
bytea ST_AsEWKB(geometry  g1)
bytea ST_AsEWKB(geometry  g1, text NDR_or_XDR)
```


## Description


Returns the [Extended Well-Known Binary](../data-management/geometry-data-type.md#EWKB_EWKT) (EWKB) representation of the geometry with SRID metadata. The first function variant defaults to encoding using server machine endian. The second function variant takes a text argument specifying the endian encoding: either 'NDR' for little-endian; or 'XDR' for big-endian. Supplying unknown arguments will result in little-endian output.


WKB format is useful to read geometry data from the database and maintaining full numeric precision. This avoids the precision rounding that can happen with text formats such as WKT.


To perform the inverse conversion of EWKB to PostGIS geometry use [ST_GeomFromEWKB](geometry-input.md#ST_GeomFromEWKB).


!!! note

    To get the OGC/ISO WKB format use [ST_AsBinary](#ST_AsBinary). Note that OGC/ISO WKB format does not include the SRID.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


## Examples


```sql
SELECT ST_AsEWKB(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326));

       st_asewkb
--------------------------------
\x0103000020e610000001000000050000000000000000000000000000000000000000000000000000
00000000000000f03f000000000000f03f000000000000f03f000000000000f03f00000000000000
0000000000000000000000000000000000
```


```sql

      SELECT ST_AsEWKB(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326), 'XDR');
       st_asewkb
--------------------------------
\x0020000003000010e600000001000000050000000000000000000000000000000000000000000000
003ff00000000000003ff00000000000003ff00000000000003ff000000000000000000000000000
0000000000000000000000000000000000

```


## See Also


[ST_AsBinary](#ST_AsBinary), [ST_GeomFromEWKB](geometry-input.md#ST_GeomFromEWKB), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_AsHEXEWKB"></a>

# ST_AsHEXEWKB

Returns a Geometry in HEXEWKB format (as text) using either little-endian (NDR) or big-endian (XDR) encoding.

## Synopsis


```sql
text ST_AsHEXEWKB(geometry  g1, text  NDRorXDR)
text ST_AsHEXEWKB(geometry  g1)
```


## Description


Returns a Geometry in HEXEWKB format (as text) using either little-endian (NDR) or big-endian (XDR) encoding. If no encoding is specified, then NDR is used.


!!! note

    Availability: 1.2.2


## Examples


```sql
SELECT ST_AsHEXEWKB(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326));
    which gives same answer as

    SELECT ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326)::text;

    st_ashexewkb
    --------
    0103000020E6100000010000000500
    00000000000000000000000000000000
    00000000000000000000000000000000F03F
    000000000000F03F000000000000F03F000000000000F03
    F000000000000000000000000000000000000000000000000
```


## Other Formats
  <a id="ST_AsEncodedPolyline"></a>

# ST_AsEncodedPolyline

Returns an Encoded Polyline from a LineString geometry.

## Synopsis


```sql
text ST_AsEncodedPolyline(geometry geom, integer  precision=5)
```


## Description


Returns the geometry as an Encoded Polyline. This format is used by Google Maps with precision=5 and by Open Source Routing Machine with precision=5 and 6.


Optional `precision` specifies how many decimal places will be preserved in Encoded Polyline. Value should be the same on encoding and decoding, or coordinates will be incorrect.


Availability: 2.2.0


## Examples


Basic


```sql

  SELECT ST_AsEncodedPolyline(GeomFromEWKT('SRID=4326;LINESTRING(-120.2 38.5,-120.95 40.7,-126.453 43.252)'));
  --result--
  |_p~iF~ps|U_ulLnnqC_mqNvxq`@

```


Use in conjunction with geography linestring and geography segmentize, and put on google maps


```
-- the SQL for Boston to San Francisco, segments every 100 KM
  SELECT ST_AsEncodedPolyline(
    ST_Segmentize(
      ST_GeogFromText('LINESTRING(-71.0519 42.4935,-122.4483 37.64)'),
        100000)::geometry) As encodedFlightPath;
```


javascript will look something like this where $ variable you replace with query result


```

<script type="text/javascript" src="http://maps.googleapis.com/maps/api/js?libraries=geometry"></script>
<script type="text/javascript">
   flightPath = new google.maps.Polyline({
      path:  google.maps.geometry.encoding.decodePath("$encodedFlightPath"),
      map: map,
      strokeColor: '#0000CC',
      strokeOpacity: 1.0,
      strokeWeight: 4
    });
</script>
```


## See Also


[ST_LineFromEncodedPolyline](geometry-input.md#ST_LineFromEncodedPolyline), [ST_Segmentize](geometry-editors.md#ST_Segmentize)
  <a id="ST_AsFlatGeobuf"></a>

# ST_AsFlatGeobuf

Return a FlatGeobuf representation of a set of rows.

## Synopsis


```sql
bytea ST_AsFlatGeobuf(anyelement set  row)
bytea ST_AsFlatGeobuf(anyelement  row, bool  index)
bytea ST_AsFlatGeobuf(anyelement  row, bool  index, text  geom_name)
```


## Description


 Return a FlatGeobuf representation ([http://flatgeobuf.org](http://flatgeobuf.org)) of a set of rows corresponding to a FeatureCollection. NOTE: PostgreSQL bytea cannot exceed 1GB.


`row` row data with at least a geometry column.


`index` toggle spatial index creation. Default is false.


`geom_name` is the name of the geometry column in the row data. If NULL it will default to the first found geometry column.


Availability: 3.2.0
  <a id="ST_AsGeobuf"></a>

# ST_AsGeobuf

Return a Geobuf representation of a set of rows.

## Synopsis


```sql
bytea ST_AsGeobuf(anyelement set  row)
bytea ST_AsGeobuf(anyelement  row, text  geom_name)
```


## Description


 Return a Geobuf representation ([https://github.com/mapbox/geobuf](https://github.com/mapbox/geobuf)) of a set of rows corresponding to a FeatureCollection. Every input geometry is analyzed to determine maximum precision for optimal storage. Note that Geobuf in its current form cannot be streamed so the full output will be assembled in memory.


`row` row data with at least a geometry column.


`geom_name` is the name of the geometry column in the row data. If NULL it will default to the first found geometry column.


Availability: 2.4.0


## Examples


```sql
SELECT encode(ST_AsGeobuf(q, 'geom'), 'base64')
    FROM (SELECT ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))') AS geom) AS q;
 st_asgeobuf
----------------------------------
 GAAiEAoOCgwIBBoIAAAAAgIAAAE=


```
  <a id="ST_AsGeoJSON"></a>

# ST_AsGeoJSON

Return a geometry or feature in GeoJSON format.

## Synopsis


```sql
text ST_AsGeoJSON(record  feature, text  geom_column="", integer  maxdecimaldigits=9, boolean  pretty_bool=false, text  id_column='')
text ST_AsGeoJSON(geometry  geom, integer  maxdecimaldigits=9, integer  options=8)
text ST_AsGeoJSON(geography  geog, integer  maxdecimaldigits=9, integer  options=0)
```


## Description


 Returns a geometry as a GeoJSON "geometry" object, or a row as a GeoJSON "feature" object.


 The resulting GeoJSON geometry and feature representations conform with the [GeoJSON specifications RFC 7946](https://tools.ietf.org/html/rfc7946), except when the parsed geometries are referenced with a CRS other than WGS84 longitude and latitude ([EPSG:4326](https://epsg.io/4326), [urn:ogc:def:crs:OGC::CRS84](http://epsg.io/4326.gml)); the GeoJSON geometry object will then have a short CRS SRID identifier attached by default. 2D and 3D Geometries are both supported. GeoJSON only supports SFS 1.1 geometry types (no curve support for example).


 The `geom_column` parameter is used to distinguish between multiple geometry columns. If omitted, the first geometry column in the record will be determined. Conversely, passing the parameter will save column type lookups.


The `maxdecimaldigits` argument may be used to reduce the maximum number of decimal places used in output (defaults to 9). If you are using EPSG:4326 and are outputting the geometry only for display, `maxdecimaldigits`=6 can be a good choice for many maps.


!!! warning

    Using the `maxdecimaldigits` parameter can cause output geometry to become invalid. To avoid this use [ST_ReducePrecision](geometry-processing.md#ST_ReducePrecision) with a suitable gridsize first.


The `options` argument can be used to add BBOX or CRS in GeoJSON output:

- 0: means no option
- 1: GeoJSON BBOX
- 2: GeoJSON Short CRS (e.g EPSG:4326)
- 4: GeoJSON Long CRS (e.g urn:ogc:def:crs:EPSG::4326)
- 8: GeoJSON Short CRS if not EPSG:4326 (default)


The `id_column` parameter is used to set the "id" member of the returned GeoJSON features. As per GeoJSON RFC, this SHOULD be used whenever a feature has a commonly used identifier, such as a primary key. When not specified, the produced features will not get an "id" member and any columns other than the geometry, including any potential keys, will just end up inside the feature’s "properties" member.


The GeoJSON specification states that polygons are oriented using the Right-Hand Rule, and some clients require this orientation. This can be ensured by using [ST_ForcePolygonCCW](geometry-editors.md#ST_ForcePolygonCCW). The specification also requires that geometry be in the WGS84 coordinate system (SRID = 4326). If necessary geometry can be projected into WGS84 using [ST_Transform](spatial-reference-system-functions.md#ST_Transform): <code>ST_Transform( geom, 4326 )</code>.


GeoJSON can be tested and viewed online at [geojson.io](http://geojson.io/) and [geojsonlint.com](http://geojson.io/). It is widely supported by web mapping frameworks:

- [OpenLayers GeoJSON Example](https://openlayers.org/en/latest/examples/geojson.html)
- [Leaflet GeoJSON Example](https://leafletjs.com/examples/geojson/)
- [Mapbox GL GeoJSON Example](https://www.mapbox.com/mapbox-gl-js/example/multiple-geometries/)


Availability: 1.3.4


Availability: 1.5.0 geography support was introduced.


Changed: 2.0.0 support default args and named args.


Changed: 3.0.0 support records as input


Changed: 3.0.0 output SRID if not EPSG:4326.


Changed: 3.5.0 allow specifying the column containing the feature id


## Examples


Generate a FeatureCollection:


```sql
SELECT json_build_object(
    'type', 'FeatureCollection',
    'features', json_agg(ST_AsGeoJSON(t.*, id_column => 'id')::json)
    )
FROM ( VALUES (1, 'one', 'POINT(1 1)'::geometry),
              (2, 'two', 'POINT(2 2)'),
              (3, 'three', 'POINT(3 3)')
     ) as t(id, name, geom);
```


```
{"type" : "FeatureCollection", "features" : [{"type": "Feature", "geometry": {"type":"Point","coordinates":[1,1]}, "id": 1, "properties": {"name": "one"}}, {"type": "Feature", "geometry": {"type":"Point","coordinates":[2,2]}, "id": 2, "properties": {"name": "two"}}, {"type": "Feature", "geometry": {"type":"Point","coordinates":[3,3]}, "id": 3, "properties": {"name": "three"}}]}
```


Generate a Feature:


```sql
SELECT ST_AsGeoJSON(t.*, id_column => 'id')
FROM (VALUES (1, 'one', 'POINT(1 1)'::geometry)) AS t(id, name, geom);
```


```
                                                  st_asgeojson
-----------------------------------------------------------------------------------------------------------------
 {"type": "Feature", "geometry": {"type":"Point","coordinates":[1,1]}, "id": 1, "properties": {"name": "one"}}
```


Don't forget to transform your data to WGS84 longitude, latitude to conform with the GeoJSON specification:


```sql
SELECT ST_AsGeoJSON(ST_Transform(geom,4326)) from fe_edges limit 1;
```


```
             st_asgeojson
-----------------------------------------------------------------------------------------------------------

{"type":"MultiLineString","coordinates":[[[-89.734634999999997,31.492072000000000],
[-89.734955999999997,31.492237999999997]]]}
```


3D geometries are supported:


```sql
SELECT ST_AsGeoJSON('LINESTRING(1 2 3, 4 5 6)');
```


```
{"type":"LineString","coordinates":[[1,2,3],[4,5,6]]}
```


Options argument can be used to add BBOX and CRS in GeoJSON output:


```sql
 SELECT ST_AsGeoJSON(ST_SetSRID('POINT(1 1)'::geometry, 4326), 9, 4|1);
```


```

  {"type":"Point","crs":{"type":"name","properties":{"name":"urn:ogc:def:crs:EPSG::4326"}},"bbox":[1.000000000,1.000000000,1.000000000,1.000000000],"coordinates":[1,1]}
```


## See Also


[ST_GeomFromGeoJSON](geometry-input.md#ST_GeomFromGeoJSON), [ST_ForcePolygonCCW](geometry-editors.md#ST_ForcePolygonCCW), [ST_Transform](spatial-reference-system-functions.md#ST_Transform)
  <a id="ST_AsGML"></a>

# ST_AsGML

Return the geometry as a GML version 2 or 3 element.

## Synopsis


```sql
text ST_AsGML(geometry  geom, integer  maxdecimaldigits=15, integer  options=0)
text ST_AsGML(geography  geog, integer  maxdecimaldigits=15, integer  options=0, text  nprefix=null, text  id=null)
text ST_AsGML(integer  version, geometry  geom, integer  maxdecimaldigits=15, integer  options=0, text  nprefix=null, text  id=null)
text ST_AsGML(integer  version, geography  geog, integer  maxdecimaldigits=15, integer  options=0, text  nprefix=null, text  id=null)
```


## Description


Return the geometry as a Geography Markup Language (GML) element. The version parameter, if specified, may be either 2 or 3. If no version parameter is specified then the default is assumed to be 2. The `maxdecimaldigits` argument may be used to reduce the maximum number of decimal places used in output (defaults to 15).


!!! warning

    Using the `maxdecimaldigits` parameter can cause output geometry to become invalid. To avoid this use [ST_ReducePrecision](geometry-processing.md#ST_ReducePrecision) with a suitable gridsize first.


GML 2 refer to 2.1.2 version, GML 3 to 3.1.1 version


The 'options' argument is a bitfield. It could be used to define CRS output type in GML output, and to declare data as lat/lon:

- 0: GML Short CRS (e.g EPSG:4326), default value
- 1: GML Long CRS (e.g urn:ogc:def:crs:EPSG::4326)
- 2: For GML 3 only, remove srsDimension attribute from output.
- 4: For GML 3 only, use <linestring> rather than <curve> tag for lines.
- 16: Declare that data are lat/lon (e.g srid=4326). Default is to assume that data are planars. This option is useful for GML 3.1.1 output only, related to axis order. So if you set it, it will swap the coordinates so order is lat lon instead of database lon lat.
- 32: Output the box of the geometry (envelope).


The 'namespace prefix' argument may be used to specify a custom namespace prefix or no prefix (if empty). If null or omitted 'gml' prefix is used


Availability: 1.3.2


Availability: 1.5.0 geography support was introduced.


Enhanced: 2.0.0 prefix support was introduced. Option 4 for GML3 was introduced to allow using LineString instead of Curve tag for lines. GML3 Support for Polyhedral surfaces and TINS was introduced. Option 32 was introduced to output the box.


Changed: 2.0.0 use default named args


Enhanced: 2.1.0 id support was introduced, for GML 3.


!!! note

    Only version 3+ of ST_AsGML supports Polyhedral Surfaces and TINS.


 SQL-MM IEC 13249-3: 17.2


## Examples: Version 2


```sql

SELECT ST_AsGML(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326));
    st_asgml
    --------
    <gml:Polygon srsName="EPSG:4326"><gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>0,0 0,1 1,1 1,0 0,0</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs></gml:Polygon>
```


## Examples: Version 3


```

-- Flip coordinates and output extended EPSG (16 | 1)--
SELECT ST_AsGML(3, ST_GeomFromText('POINT(5.234234233242 6.34534534534)',4326), 5, 17);
      st_asgml
      --------
    <gml:Point srsName="urn:ogc:def:crs:EPSG::4326"><gml:pos>6.34535 5.23423</gml:pos></gml:Point>
```


```

-- Output the envelope (32) --
SELECT ST_AsGML(3, ST_GeomFromText('LINESTRING(1 2, 3 4, 10 20)',4326), 5, 32);
    st_asgml
    --------
  <gml:Envelope srsName="EPSG:4326">
    <gml:lowerCorner>1 2</gml:lowerCorner>
    <gml:upperCorner>10 20</gml:upperCorner>
  </gml:Envelope>
```


```

-- Output the envelope (32) , reverse (lat lon instead of lon lat) (16), long srs (1)= 32 | 16 | 1 = 49 --
SELECT ST_AsGML(3, ST_GeomFromText('LINESTRING(1 2, 3 4, 10 20)',4326), 5, 49);
  st_asgml
  --------
<gml:Envelope srsName="urn:ogc:def:crs:EPSG::4326">
  <gml:lowerCorner>2 1</gml:lowerCorner>
  <gml:upperCorner>20 10</gml:upperCorner>
</gml:Envelope>
```


```

-- Polyhedral Example --
SELECT ST_AsGML(3, ST_GeomFromEWKT('POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)), ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)), ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )'));
  st_asgml
  --------
 <gml:PolyhedralSurface>
<gml:polygonPatches>
   <gml:PolygonPatch>
    <gml:exterior>
        <gml:LinearRing>
           <gml:posList srsDimension="3">0 0 0 0 0 1 0 1 1 0 1 0 0 0 0</gml:posList>
        </gml:LinearRing>
    </gml:exterior>
   </gml:PolygonPatch>
   <gml:PolygonPatch>
    <gml:exterior>
        <gml:LinearRing>
           <gml:posList srsDimension="3">0 0 0 0 1 0 1 1 0 1 0 0 0 0 0</gml:posList>
        </gml:LinearRing>
    </gml:exterior>
   </gml:PolygonPatch>
   <gml:PolygonPatch>
    <gml:exterior>
        <gml:LinearRing>
           <gml:posList srsDimension="3">0 0 0 1 0 0 1 0 1 0 0 1 0 0 0</gml:posList>
        </gml:LinearRing>
    </gml:exterior>
   </gml:PolygonPatch>
   <gml:PolygonPatch>
    <gml:exterior>
        <gml:LinearRing>
           <gml:posList srsDimension="3">1 1 0 1 1 1 1 0 1 1 0 0 1 1 0</gml:posList>
        </gml:LinearRing>
    </gml:exterior>
   </gml:PolygonPatch>
   <gml:PolygonPatch>
    <gml:exterior>
        <gml:LinearRing>
           <gml:posList srsDimension="3">0 1 0 0 1 1 1 1 1 1 1 0 0 1 0</gml:posList>
        </gml:LinearRing>
    </gml:exterior>
   </gml:PolygonPatch>
   <gml:PolygonPatch>
    <gml:exterior>
        <gml:LinearRing>
           <gml:posList srsDimension="3">0 0 1 1 0 1 1 1 1 0 1 1 0 0 1</gml:posList>
        </gml:LinearRing>
    </gml:exterior>
   </gml:PolygonPatch>
</gml:polygonPatches>
</gml:PolyhedralSurface>
```


## See Also


[ST_GeomFromGML](geometry-input.md#ST_GeomFromGML)
  <a id="ST_AsKML"></a>

# ST_AsKML

Return the geometry as a KML element.

## Synopsis


```sql
text ST_AsKML(geometry  geom, integer  maxdecimaldigits=15, text  nprefix=NULL)
text ST_AsKML(geography  geog, integer  maxdecimaldigits=15, text  nprefix=NULL)
```


## Description


Return the geometry as a Keyhole Markup Language (KML) element. default maximum number of decimal places is 15, default namespace is no prefix.


!!! warning

    Using the `maxdecimaldigits` parameter can cause output geometry to become invalid. To avoid this use [ST_ReducePrecision](geometry-processing.md#ST_ReducePrecision) with a suitable gridsize first.


!!! note

    Requires PostGIS be compiled with Proj support. Use [PostGIS_Full_Version](version-functions.md#PostGIS_Full_Version) to confirm you have proj support compiled in.


!!! note

    Availability: 1.2.2 - later variants that include version param came in 1.3.2


!!! note

    Enhanced: 2.0.0 - Add prefix namespace, use default and named args


!!! note

    Changed: 3.0.0 - Removed the "versioned" variant signature


!!! note

    AsKML output will not work with geometries that do not have an SRID


## Examples


```sql

SELECT ST_AsKML(ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))',4326));

    st_askml
    --------
    <Polygon><outerBoundaryIs><LinearRing><coordinates>0,0 0,1 1,1 1,0 0,0</coordinates></LinearRing></outerBoundaryIs></Polygon>

    --3d linestring
    SELECT ST_AsKML('SRID=4326;LINESTRING(1 2 3, 4 5 6)');
    <LineString><coordinates>1,2,3 4,5,6</coordinates></LineString>
```


## See Also


[ST_AsSVG](#ST_AsSVG), [ST_AsGML](#ST_AsGML)
  <a id="ST_AsLatLonText"></a>

# ST_AsLatLonText

Return the Degrees, Minutes, Seconds representation of the given point.

## Synopsis


```sql
text ST_AsLatLonText(geometry  pt, text  format='')
```


## Description


Returns the Degrees, Minutes, Seconds representation of the point.


!!! note

    It is assumed the point is in a lat/lon projection. The X (lon) and Y (lat) coordinates are normalized in the output to the "normal" range (-180 to +180 for lon, -90 to +90 for lat).


 The text parameter is a format string containing the format for the resulting text, similar to a date format string. Valid tokens are "D" for degrees, "M" for minutes, "S" for seconds, and "C" for cardinal direction (NSEW). DMS tokens may be repeated to indicate desired width and precision ("SSS.SSSS" means " 1.0023").


 "M", "S", and "C" are optional. If "C" is omitted, degrees are shown with a "-" sign if south or west. If "S" is omitted, minutes will be shown as decimal with as many digits of precision as you specify. If "M" is also omitted, degrees are shown as decimal with as many digits precision as you specify.


 If the format string is omitted (or zero-length) a default format will be used.


Availability: 2.0


## Examples


Default format.


```sql

SELECT (ST_AsLatLonText('POINT (-3.2342342 -2.32498)'));
      st_aslatlontext
----------------------------
 2°19'29.928"S 3°14'3.243"W
```


Providing a format (same as the default).


```sql

SELECT (ST_AsLatLonText('POINT (-3.2342342 -2.32498)', 'D°M''S.SSS"C'));
      st_aslatlontext
----------------------------
 2°19'29.928"S 3°14'3.243"W
```


Characters other than D, M, S, C and . are just passed through.


```sql

SELECT (ST_AsLatLonText('POINT (-3.2342342 -2.32498)', 'D degrees, M minutes, S seconds to the C'));
                                   st_aslatlontext
--------------------------------------------------------------------------------------
 2 degrees, 19 minutes, 30 seconds to the S 3 degrees, 14 minutes, 3 seconds to the W
```


Signed degrees instead of cardinal directions.


```sql

SELECT (ST_AsLatLonText('POINT (-3.2342342 -2.32498)', 'D°M''S.SSS"'));
      st_aslatlontext
----------------------------
 -2°19'29.928" -3°14'3.243"
```


Decimal degrees.


```sql

SELECT (ST_AsLatLonText('POINT (-3.2342342 -2.32498)', 'D.DDDD degrees C'));
          st_aslatlontext
-----------------------------------
 2.3250 degrees S 3.2342 degrees W
```


Excessively large values are normalized.


```sql

SELECT (ST_AsLatLonText('POINT (-302.2342342 -792.32498)'));
        st_aslatlontext
-------------------------------
 72°19'29.928"S 57°45'56.757"E
```
  <a id="ST_AsMARC21"></a>

# ST_AsMARC21

Returns geometry as a MARC21/XML record with a geographic datafield (034).

## Synopsis


```sql
text
                        ST_AsMARC21(geometry
                        geom, text
                        format='hdddmmss')
```


## Description


This function returns a MARC21/XML record with [Coded Cartographic Mathematical Data](https://www.loc.gov/marc/bibliographic/bd034.html) representing the bounding box of a given geometry. The `format` parameter allows to encode the coordinates in subfields `$d`,`$e`,`$f` and `$g` in all formats supported by the MARC21/XML standard. Valid formats are:


- cardinal direction, degrees, minutes and seconds (default): `hdddmmss`
- decimal degrees with cardinal direction: `hddd.dddddd`
- decimal degrees without cardinal direction: `ddd.dddddd`
- decimal minutes with cardinal direction: `hdddmm.mmmm`
- decimal minutes without cardinal direction: `dddmm.mmmm`
- decimal seconds with cardinal direction: `hdddmmss.sss`


The decimal sign may be also a comma, e.g. `hdddmm,mmmm`.


The precision of decimal formats can be limited by the number of characters after the decimal sign, e.g. `hdddmm.mm` for decimal minutes with a precision of two decimals.


This function ignores the Z and M dimensions.


 LOC MARC21/XML versions supported:

- [MARC21/XML 1.1](https://www.loc.gov/standards/marcxml/)


Availability: 3.3.0


!!! note

    This function does not support non lon/lat geometries, as they are not supported by the MARC21/XML standard (Coded Cartographic Mathematical Data).


!!! note

    The MARC21/XML Standard does not provide any means to annotate the spatial reference system for Coded Cartographic Mathematical Data, which means that this information will be lost after conversion to MARC21/XML.


## Examples


Converting a `POINT` to MARC21/XML formatted as hdddmmss (default)


```sql


                SELECT ST_AsMARC21('SRID=4326;POINT(-4.504289 54.253312)'::geometry);

                                st_asmarc21
                -------------------------------------------------
                <record xmlns="http://www.loc.gov/MARC21/slim">
                    <datafield tag="034" ind1="1" ind2=" ">
                        <subfield code="a">a</subfield>
                        <subfield code="d">W0043015</subfield>
                        <subfield code="e">W0043015</subfield>
                        <subfield code="f">N0541512</subfield>
                        <subfield code="g">N0541512</subfield>
                    </datafield>
                </record>
```


Converting a `POLYGON` to MARC21/XML formatted in decimal degrees


```sql


                SELECT ST_AsMARC21('SRID=4326;POLYGON((-4.5792388916015625 54.18172660239091,-4.56756591796875 54.196993557130355,-4.546623229980469 54.18313300502024,-4.5792388916015625 54.18172660239091))'::geometry,'hddd.dddd');

                <record xmlns="http://www.loc.gov/MARC21/slim">
                    <datafield tag="034" ind1="1" ind2=" ">
                        <subfield code="a">a</subfield>
                        <subfield code="d">W004.5792</subfield>
                        <subfield code="e">W004.5466</subfield>
                        <subfield code="f">N054.1970</subfield>
                        <subfield code="g">N054.1817</subfield>
                    </datafield>
                </record>
```


Converting a `GEOMETRYCOLLECTION` to MARC21/XML formatted in decimal minutes. The geometries order in the MARC21/XML output correspond to their order in the collection.


```sql


                SELECT ST_AsMARC21('SRID=4326;GEOMETRYCOLLECTION(POLYGON((13.1 52.65,13.516666666666667 52.65,13.516666666666667 52.38333333333333,13.1 52.38333333333333,13.1 52.65)),POINT(-4.5 54.25))'::geometry,'hdddmm.mmmm');

                                st_asmarc21
                -------------------------------------------------
                <record xmlns="http://www.loc.gov/MARC21/slim">
                    <datafield tag="034" ind1="1" ind2=" ">
                        <subfield code="a">a</subfield>
                        <subfield code="d">E01307.0000</subfield>
                        <subfield code="e">E01331.0000</subfield>
                        <subfield code="f">N05240.0000</subfield>
                        <subfield code="g">N05224.0000</subfield>
                    </datafield>
                    <datafield tag="034" ind1="1" ind2=" ">
                        <subfield code="a">a</subfield>
                        <subfield code="d">W00430.0000</subfield>
                        <subfield code="e">W00430.0000</subfield>
                        <subfield code="f">N05415.0000</subfield>
                        <subfield code="g">N05415.0000</subfield>
                    </datafield>
                </record>
```


## See Also


[ST_GeomFromMARC21](geometry-input.md#ST_GeomFromMARC21)
  <a id="ST_AsMVTGeom"></a>

# ST_AsMVTGeom

Transforms a geometry into the coordinate space of a MVT tile.

## Synopsis


```sql
geometry ST_AsMVTGeom(geometry  geom, box2d  bounds, integer  extent=4096, integer  buffer=256, boolean  clip_geom=true)
```


## Description


Transforms a geometry into the coordinate space of a MVT ([Mapbox Vector Tile](https://www.mapbox.com/vector-tiles/)) tile, clipping it to the tile bounds if required. The geometry must be in the coordinate system of the target map (using [ST_Transform](spatial-reference-system-functions.md#ST_Transform) if needed). Commonly this is [Web Mercator](https://en.wikipedia.org/wiki/Web_Mercator_projection) (SRID:3857).


The function attempts to preserve geometry validity, and corrects it if needed. This may cause the result geometry to collapse to a lower dimension.


The rectangular bounds of the tile in the target map coordinate space must be provided, so the geometry can be transformed, and clipped if required. The bounds can be generated using [ST_TileEnvelope](geometry-constructors.md#ST_TileEnvelope).


 This function is used to convert geometry into the tile coordinate space required by [ST_AsMVT](#ST_AsMVT).


`geom` is the geometry to transform, in the coordinate system of the target map.


`bounds` is the rectangular bounds of the tile in map coordinate space, with no buffer.


`extent` is the tile extent size in tile coordinate space as defined by the [MVT specification](https://www.mapbox.com/vector-tiles/specification/). Defaults to 4096.


`buffer` is the buffer size in tile coordinate space for geometry clippig. Defaults to 256.


`clip_geom` is a boolean to control if geometries are clipped or encoded as-is. Defaults to true.


Availability: 2.4.0


!!! note

    From 3.0, Wagyu can be chosen at configure time to clip and validate MVT polygons. This library is faster and produces more correct results than the GEOS default, but it might drop small polygons.


## Examples


```sql

SELECT ST_AsText(ST_AsMVTGeom(
  ST_GeomFromText('POLYGON ((0 0, 10 0, 10 5, 0 -5, 0 0))'),
  ST_MakeBox2D(ST_Point(0, 0), ST_Point(4096, 4096)),
  4096, 0, false));
                              st_astext
--------------------------------------------------------------------
 MULTIPOLYGON(((5 4096,10 4091,10 4096,5 4096)),((5 4096,0 4101,0 4096,5 4096)))

```


Canonical example for a Web Mercator tile using a computed tile bounds to query and clip geometry. This assumes the data.geom column has srid of 4326.


```sql


SELECT ST_AsMVTGeom(
            ST_Transform( geom, 3857 ),
            ST_TileEnvelope(12, 513, 412), extent => 4096, buffer => 64) AS geom
  FROM data
  WHERE geom && ST_Transform(ST_TileEnvelope(12, 513, 412, margin => (64.0 / 4096)),4326)
```


## See Also


 [ST_AsMVT](#ST_AsMVT), [ST_TileEnvelope](geometry-constructors.md#ST_TileEnvelope), [PostGIS_Wagyu_Version](version-functions.md#PostGIS_Wagyu_Version)
  <a id="ST_AsMVT"></a>

# ST_AsMVT

Aggregate function returning a MVT representation of a set of rows.

## Synopsis


```sql
bytea ST_AsMVT(anyelement set  row)
bytea ST_AsMVT(anyelement  row, text  name)
bytea ST_AsMVT(anyelement  row, text  name, integer  extent)
bytea ST_AsMVT(anyelement  row, text  name, integer  extent, text  geom_name)
bytea ST_AsMVT(anyelement  row, text  name, integer  extent, text  geom_name, text  feature_id_name)
```


## Description


An aggregate function which returns a binary [Mapbox Vector Tile](https://www.mapbox.com/vector-tiles/) representation of a set of rows corresponding to a tile layer. The rows must contain a geometry column which will be encoded as a feature geometry. The geometry must be in tile coordinate space and valid as per the [MVT specification](https://www.mapbox.com/vector-tiles/specification/). [ST_AsMVTGeom](#ST_AsMVTGeom) can be used to transform geometry into tile coordinate space. Other row columns are encoded as feature attributes.


The [Mapbox Vector Tile](https://www.mapbox.com/vector-tiles/) format can store features with varying sets of attributes. To use this capability supply a JSONB column in the row data containing Json objects one level deep. The keys and values in the JSONB values will be encoded as feature attributes.


 Tiles with multiple layers can be created by concatenating multiple calls to this function using `||` or `STRING_AGG`.


!!! important

    Do not call with a `GEOMETRYCOLLECTION` as an element in the row. However you can use [ST_AsMVTGeom](#ST_AsMVTGeom) to prepare a geometry collection for inclusion.


`row` row data with at least a geometry column.


`name` is the name of the layer. Default is the string "default".


`extent` is the tile extent in screen space as defined by the specification. Default is 4096.


`geom_name` is the name of the geometry column in the row data. Default is the first geometry column. Note that PostgreSQL by default automatically [folds unquoted identifiers to lower case](https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS), which means that unless the geometry column is quoted, e.g. `"MyMVTGeom"`, this parameter must be provided as lowercase.


`feature_id_name` is the name of the Feature ID column in the row data. If NULL or negative the Feature ID is not set. The first column matching name and valid type (smallint, integer, bigint) will be used as Feature ID, and any subsequent column will be added as a property. JSON properties are not supported.


Enhanced: 3.0 - added support for Feature ID.


Enhanced: 2.5.0 - added support parallel query.


Availability: 2.4.0


## Examples


```sql

WITH mvtgeom AS
(
  SELECT ST_AsMVTGeom(geom, ST_TileEnvelope(12, 513, 412), extent => 4096, buffer => 64) AS geom, name, description
  FROM points_of_interest
  WHERE geom && ST_TileEnvelope(12, 513, 412, margin => (64.0 / 4096))
)
SELECT ST_AsMVT(mvtgeom.*)
FROM mvtgeom;


```


## See Also


 [ST_AsMVTGeom](#ST_AsMVTGeom), [ST_TileEnvelope](geometry-constructors.md#ST_TileEnvelope)
  <a id="ST_AsSVG"></a>

# ST_AsSVG

Returns SVG path data for a geometry.

## Synopsis


```sql
text ST_AsSVG(geometry  geom, integer  rel=0, integer  maxdecimaldigits=15)
text ST_AsSVG(geography  geog, integer  rel=0, integer  maxdecimaldigits=15)
```


## Description


Return the geometry as Scalar Vector Graphics (SVG) path data. Use 1 as second argument to have the path data implemented in terms of relative moves, the default (or 0) uses absolute moves. Third argument may be used to reduce the maximum number of decimal digits used in output (defaults to 15). Point geometries will be rendered as cx/cy when 'rel' arg is 0, x/y when 'rel' is 1. Multipoint geometries are delimited by commas (","), GeometryCollection geometries are delimited by semicolons (";").


For working with PostGIS SVG graphics, checkout [pg_svg](https://github.com/dr-jts/pg_svg) library which provides plpgsql functions for working with outputs from ST_AsSVG.


Enhanced: 3.4.0 to support all curve types


Changed: 2.0.0 to use default args and support named args


!!! note

    Availability: 1.2.2. Availability: 1.4.0 Changed in PostGIS 1.4.0 to include L command in absolute path to conform to [http://www.w3.org/TR/SVG/paths.html#PathDataBNF](http://www.w3.org/TR/SVG/paths.html#PathDataBNF)


## Examples


```sql
SELECT ST_AsSVG('POLYGON((0 0,0 1,1 1,1 0,0 0))'::geometry);

st_assvg
--------
M 0 0 L 0 -1 1 -1 1 0 Z
```


Circular string


```sql
SELECT ST_AsSVG( ST_GeomFromText('CIRCULARSTRING(-2 0,0 2,2 0,0 2,2 4)') );

st_assvg
--------
M -2 0 A 2 2 0 0 1 2 0 A 2 2 0 0 1 2 -4
```


Multi-curve


```sql
SELECT ST_AsSVG('MULTICURVE((5 5,3 5,3 3,0 3),
 CIRCULARSTRING(0 0,2 1,2 2))'::geometry, 0, 0);
 st_assvg
------------------------------------------------
 M 5 -5 L 3 -5 3 -3 0 -3 M 0 0 A 2 2 0 0 0 2 -2

```


Multi-surface


```sql
SELECT ST_AsSVG('MULTISURFACE(
CURVEPOLYGON(CIRCULARSTRING(-2 0,-1 -1,0 0,1 -1,2 0,0 2,-2 0),
    (-1 0,0 0.5,1 0,0 1,-1 0)),
((7 8,10 10,6 14,4 11,7 8)))'::geometry, 0, 2);

st_assvg
---------------------------------------------------------
M -2 0 A 1 1 0 0 0 0 0 A 1 1 0 0 0 2 0 A 2 2 0 0 0 -2 0 Z
M -1 0 L 0 -0.5 1 0 0 -1 -1 0 Z
M 7 -8 L 10 -10 6 -14 4 -11 Z

```
  <a id="ST_AsTWKB"></a>

# ST_AsTWKB

Returns the geometry as TWKB, aka "Tiny Well-Known Binary"

## Synopsis


```sql
bytea ST_AsTWKB(geometry  geom, integer  prec=0, integer  prec_z=0, integer  prec_m=0, boolean  with_sizes=false, boolean  with_boxes=false)
bytea ST_AsTWKB(geometry[]  geom, bigint[]  ids, integer  prec=0, integer  prec_z=0, integer  prec_m=0, boolean  with_sizes=false, boolean  with_boxes=false)
```


## Description


Returns the geometry in TWKB (Tiny Well-Known Binary) format. TWKB is a [compressed binary format](https://github.com/TWKB/Specification/blob/master/twkb.md) with a focus on minimizing the size of the output.


The decimal digits parameters control how much precision is stored in the output. By default, values are rounded to the nearest unit before encoding. If you want to transfer more precision, increase the number. For example, a value of 1 implies that the first digit to the right of the decimal point will be preserved.


The sizes and bounding boxes parameters control whether optional information about the encoded length of the object and the bounds of the object are included in the output. By default they are not. Do not turn them on unless your client software has a use for them, as they just use up space (and saving space is the point of TWKB).


The array-input form of the function is used to convert a collection of geometries and unique identifiers into a TWKB collection that preserves the identifiers. This is useful for clients that expect to unpack a collection and then access further information about the objects inside. You can create the arrays using the [array_agg](https://www.postgresql.org/docs/current/functions-aggregate.html) function. The other parameters operate the same as for the simple form of the function.


!!! note

    The format specification is available online at [https://github.com/TWKB/Specification](https://github.com/TWKB/Specification), and code for building a JavaScript client can be found at [https://github.com/TWKB/twkb.js](https://github.com/TWKB/twkb.js).


Enhanced: 2.4.0 memory and speed improvements.


Availability: 2.2.0


## Examples


```sql

SELECT ST_AsTWKB('LINESTRING(1 1,5 5)'::geometry);
                 st_astwkb
--------------------------------------------
\x02000202020808
```


To create an aggregate TWKB object including identifiers aggregate the desired geometries and objects first, using "array_agg()", then call the appropriate TWKB function.


```sql

SELECT ST_AsTWKB(array_agg(geom), array_agg(gid)) FROM mytable;
                 st_astwkb
--------------------------------------------
\x040402020400000202
```


## See Also


[ST_GeomFromTWKB](geometry-input.md#ST_GeomFromTWKB), [ST_AsBinary](#ST_AsBinary), [ST_AsEWKB](#ST_AsEWKB), [ST_AsEWKT](#ST_AsEWKT), [ST_GeomFromText](geometry-input.md#ST_GeomFromText)
  <a id="ST_AsX3D"></a>

# ST_AsX3D

Returns a Geometry in X3D xml node element format: ISO-IEC-19776-1.2-X3DEncodings-XML

## Synopsis


```sql
text ST_AsX3D(geometry  g1, integer  maxdecimaldigits=15, integer  options=0)
```


## Description


Returns a geometry as an X3D xml formatted node element [http://www.web3d.org/standards/number/19776-1](http://www.web3d.org/standards/number/19776-1). If `maxdecimaldigits` (precision) is not specified then defaults to 15.


!!! note

    There are various options for translating PostGIS geometries to X3D since X3D geometry types don't map directly to PostGIS geometry types and some newer X3D types that might be better mappings we have avoided since most rendering tools don't currently support them. These are the mappings we have settled on. Feel free to post a bug ticket if you have thoughts on the idea or ways we can allow people to denote their preferred mappings.


    Below is how we currently map PostGIS 2D/3D types to X3D types


The 'options' argument is a bitfield. For PostGIS 2.2+, this is used to denote whether to represent coordinates with X3D GeoCoordinates Geospatial node and also whether to flip the x/y axis. By default, <code>ST_AsX3D</code> outputs in database form (long,lat or X,Y), but X3D default of lat/lon, y/x may be preferred.


- 0: X/Y in database order (e.g. long/lat = X,Y is standard database order), default value, and non-spatial coordinates (just regular old Coordinate tag).
- 1: Flip X and Y. If used in conjunction with the GeoCoordinate option switch, then output will be default "latitude_first" and coordinates will be flipped as well.
- 2: Output coordinates in GeoSpatial GeoCoordinates. This option will throw an error if geometries are not in WGS 84 long lat (srid: 4326). This is currently the only GeoCoordinate type supported. [Refer to X3D specs specifying a spatial reference system.](http://www.web3d.org/documents/specifications/19775-1/V3.2/Part01/components/geodata.html#Specifyingaspatialreference). Default output will be <code>GeoCoordinate geoSystem='"GD" "WE" "longitude_first"'</code>. If you prefer the X3D default of <code>GeoCoordinate geoSystem='"GD" "WE" "latitude_first"'</code> use <code>(2 + 1)</code> = <code>3</code>


| PostGIS Type | 2D X3D Type | 3D X3D Type |
| --- | --- | --- |
| LINESTRING | not yet implemented - will be PolyLine2D | LineSet |
| MULTILINESTRING | not yet implemented - will be PolyLine2D | IndexedLineSet |
| MULTIPOINT | Polypoint2D | PointSet |
| POINT | outputs the space delimited coordinates | outputs the space delimited coordinates |
| (MULTI) POLYGON, POLYHEDRALSURFACE | Invalid X3D markup | IndexedFaceSet (inner rings currently output as another faceset) |
| TIN | TriangleSet2D (Not Yet Implemented) | IndexedTriangleSet |


!!! note

    2D geometry support not yet complete. Inner rings currently just drawn as separate polygons. We are working on these.


Lots of advancements happening in 3D space particularly with [X3D Integration with HTML5](https://www.web3d.org/wiki/index.php/X3D_and_HTML5)


There is also a nice open source X3D viewer you can use to view rendered geometries. Free Wrl [http://freewrl.sourceforge.net/](http://freewrl.sourceforge.net/) binaries available for Mac, Linux, and Windows. Use the FreeWRL_Launcher packaged to view the geometries.


Also check out [PostGIS minimalist X3D viewer](https://git.osgeo.org/gitea/robe/postgis_x3d_viewer) that utilizes this function and [x3dDom html/js open source toolkit](http://www.x3dom.org/).


Availability: 2.0.0: ISO-IEC-19776-1.2-X3DEncodings-XML


Enhanced: 2.2.0: Support for GeoCoordinates and axis (x/y, long/lat) flipping. Look at options for details.


## Example: Create a fully functional X3D document - This will generate a cube that is viewable in FreeWrl and other X3D viewers.


```sql

SELECT '<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE X3D PUBLIC "ISO//Web3D//DTD X3D 3.0//EN" "http://www.web3d.org/specifications/x3d-3.0.dtd">
<X3D>
  <Scene>
    <Transform>
      <Shape>
       <Appearance>
            <Material emissiveColor=''0 0 1''/>
       </Appearance> ' ||
       ST_AsX3D( ST_GeomFromEWKT('POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)), ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)), ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )')) ||
      '</Shape>
    </Transform>
  </Scene>
</X3D>' As x3ddoc;

    x3ddoc
    --------
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE X3D PUBLIC "ISO//Web3D//DTD X3D 3.0//EN" "http://www.web3d.org/specifications/x3d-3.0.dtd">
<X3D>
  <Scene>
    <Transform>
      <Shape>
       <Appearance>
            <Material emissiveColor='0 0 1'/>
       </Appearance>
       <IndexedFaceSet  coordIndex='0 1 2 3 -1 4 5 6 7 -1 8 9 10 11 -1 12 13 14 15 -1 16 17 18 19 -1 20 21 22 23'>
            <Coordinate point='0 0 0 0 0 1 0 1 1 0 1 0 0 0 0 0 1 0 1 1 0 1 0 0 0 0 0 1 0 0 1 0 1 0 0 1 1 1 0 1 1 1 1 0 1 1 0 0 0 1 0 0 1 1 1 1 1 1 1 0 0 0 1 1 0 1 1 1 1 0 1 1' />
      </IndexedFaceSet>
      </Shape>
    </Transform>
  </Scene>
</X3D>
```


## PostGIS buildings


 Copy and paste the output of this query to [x3d scene viewer](http://postgis.net/docs/support/viewers/x3d_viewer.htm) and click Show


```sql

SELECT string_agg('<Shape>' || ST_AsX3D(ST_Extrude(geom, 0,0, i*0.5)) ||
    '<Appearance>
          <Material diffuseColor="' || (0.01*i)::text || ' 0.8 0.2" specularColor="' || (0.05*i)::text || ' 0 0.5"/>
        </Appearance>
    </Shape>', '')
FROM ST_Subdivide(ST_Letters('PostGIS'),20) WITH ORDINALITY AS f(geom,i);
```


![image](images/st_asx3d01.png)


Buildings formed by subdividing PostGIS and extrusion


## Example: An Octagon elevated 3 Units and decimal precision of 6


```sql

SELECT ST_AsX3D(
ST_Translate(
    ST_Force_3d(
        ST_Buffer(ST_Point(10,10),5, 'quad_segs=2')), 0,0,
    3)
  ,6) As x3dfrag;

x3dfrag
--------
<IndexedFaceSet coordIndex="0 1 2 3 4 5 6 7">
    <Coordinate point="15 10 3 13.535534 6.464466 3 10 5 3 6.464466 6.464466 3 5 10 3 6.464466 13.535534 3 10 15 3 13.535534 13.535534 3 " />
</IndexedFaceSet>
```


## Example: TIN


```sql

SELECT ST_AsX3D(ST_GeomFromEWKT('TIN (((
                0 0 0,
                0 0 1,
                0 1 0,
                0 0 0
            )), ((
                0 0 0,
                0 1 0,
                1 1 0,
                0 0 0
            ))
            )')) As x3dfrag;

    x3dfrag
    --------
<IndexedTriangleSet  index='0 1 2 3 4 5'><Coordinate point='0 0 0 0 0 1 0 1 0 0 0 0 0 1 0 1 1 0'/></IndexedTriangleSet>
```


## Example: Closed multilinestring (the boundary of a polygon with holes)


```sql

SELECT ST_AsX3D(
        ST_GeomFromEWKT('MULTILINESTRING((20 0 10,16 -12 10,0 -16 10,-12 -12 10,-20 0 10,-12 16 10,0 24 10,16 16 10,20 0 10),
  (12 0 10,8 8 10,0 12 10,-8 8 10,-8 0 10,-8 -4 10,0 -8 10,8 -4 10,12 0 10))')
) As x3dfrag;

    x3dfrag
    --------
<IndexedLineSet  coordIndex='0 1 2 3 4 5 6 7 0 -1 8 9 10 11 12 13 14 15 8'>
    <Coordinate point='20 0 10 16 -12 10 0 -16 10 -12 -12 10 -20 0 10 -12 16 10 0 24 10 16 16 10 12 0 10 8 8 10 0 12 10 -8 8 10 -8 0 10 -8 -4 10 0 -8 10 8 -4 10 ' />
 </IndexedLineSet>
```
  <a id="ST_GeoHash"></a>

# ST_GeoHash

Return a GeoHash representation of the geometry.

## Synopsis


```sql
text ST_GeoHash(geometry  geom, integer  maxchars=full_precision_of_point)
```


## Description


Computes a [GeoHash](http://en.wikipedia.org/wiki/Geohash) representation of a geometry. A GeoHash encodes a geographic Point into a text form that is sortable and searchable based on prefixing. A shorter GeoHash is a less precise representation of a point. It can be thought of as a box that contains the point.


Non-point geometry values with non-zero extent can also be mapped to GeoHash codes. The precision of the code depends on the geographic extent of the geometry.


If `maxchars` is not specified, the returned GeoHash code is for the smallest cell containing the input geometry. Points return a GeoHash with 20 characters of precision (about enough to hold the full double precision of the input). Other geometric types may return a GeoHash with less precision, depending on the extent of the geometry. Larger geometries are represented with less precision, smaller ones with more precision. The box determined by the GeoHash code always contains the input feature.


If `maxchars` is specified the returned GeoHash code has at most that many characters. It maps to a (possibly) lower precision representation of the input geometry. For non-points, the starting point of the calculation is the center of the bounding box of the geometry.


Availability: 1.4.0


!!! note

    ST_GeoHash requires input geometry to be in geographic (lon/lat) coordinates.


## Examples


```sql
SELECT ST_GeoHash( ST_Point(-126,48) );

   st_geohash
----------------------
 c0w3hf1s70w3hf1s70w3

SELECT ST_GeoHash( ST_Point(-126,48), 5);

 st_geohash
------------
 c0w3h

-- This line contains the point, so the GeoHash is a prefix of the point code
SELECT ST_GeoHash('LINESTRING(-126 48, -126.1 48.1)'::geometry);

 st_geohash
------------
 c0w3


```


## See Also


[ST_GeomFromGeoHash](geometry-input.md#ST_GeomFromGeoHash), [ST_PointFromGeoHash](geometry-input.md#ST_PointFromGeoHash), [ST_Box2dFromGeoHash](geometry-input.md#ST_Box2dFromGeoHash)
