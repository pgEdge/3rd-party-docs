<a id="Geometry_Inputs"></a>

## Geometry Input


## Well-Known Text (WKT)
  <a id="ST_BdPolyFromText"></a>

# ST_BdPolyFromText

Construct a Polygon given an arbitrary collection of closed linestrings as a MultiLineString Well-Known text representation.

## Synopsis


```sql
geometry ST_BdPolyFromText(text  WKT, integer  srid)
```


## Description


Construct a Polygon given an arbitrary collection of closed linestrings as a MultiLineString Well-Known text representation.


!!! note

    Throws an error if WKT is not a MULTILINESTRING. Throws an error if output is a MULTIPOLYGON; use ST_BdMPolyFromText in that case, or see ST_BuildArea() for a postgis-specific approach.


 s3.2.6.2


Performed by the GEOS module.


Availability: 1.1.0


## See Also


[ST_BuildArea](geometry-processing.md#ST_BuildArea), [ST_BdMPolyFromText](#ST_BdMPolyFromText)
  <a id="ST_BdMPolyFromText"></a>

# ST_BdMPolyFromText

Construct a MultiPolygon given an arbitrary collection of closed linestrings as a MultiLineString text representation Well-Known text representation.

## Synopsis


```sql
geometry ST_BdMPolyFromText(text  WKT, integer  srid)
```


## Description


Construct a Polygon given an arbitrary collection of closed linestrings, polygons, MultiLineStrings as Well-Known text representation.


!!! note

    Throws an error if WKT is not a MULTILINESTRING. Forces MULTIPOLYGON output even when result is really only composed by a single POLYGON; use [ST_BdPolyFromText](#ST_BdPolyFromText) if you're sure a single POLYGON will result from operation, or see [ST_BuildArea()](geometry-processing.md#ST_BuildArea) for a postgis-specific approach.


 s3.2.6.2


Performed by the GEOS module.


Availability: 1.1.0


## See Also


[ST_BuildArea](geometry-processing.md#ST_BuildArea), [ST_BdPolyFromText](#ST_BdPolyFromText)
  <a id="ST_GeogFromText"></a>

# ST_GeogFromText

Return a specified geography value from Well-Known Text representation or extended (WKT).

## Synopsis


```sql
geography ST_GeogFromText(text  EWKT)
```


## Description


Returns a geography object from the well-known text or extended well-known representation. SRID 4326 is assumed if unspecified. This is an alias for ST_GeographyFromText. Points are always expressed in long lat form.


## Examples


```

--- converting lon lat coords to geography
ALTER TABLE sometable ADD COLUMN geog geography(POINT,4326);
UPDATE sometable SET geog = ST_GeogFromText('SRID=4326;POINT(' || lon || ' ' || lat || ')');

--- specify a geography point using EPSG:4267, NAD27
SELECT ST_AsEWKT(ST_GeogFromText('SRID=4267;POINT(-77.0092 38.889588)'));

```


## See Also


[ST_AsText](geometry-output.md#ST_AsText), [ST_GeographyFromText](#ST_GeographyFromText)
  <a id="ST_GeographyFromText"></a>

# ST_GeographyFromText

Return a specified geography value from Well-Known Text representation or extended (WKT).

## Synopsis


```sql
geography ST_GeographyFromText(text  EWKT)
```


## Description


Returns a geography object from the well-known text representation. SRID 4326 is assumed if unspecified.


## See Also


[ST_GeogFromText](#ST_GeogFromText), [ST_AsText](geometry-output.md#ST_AsText)
  <a id="ST_GeomCollFromText"></a>

# ST_GeomCollFromText

Makes a collection Geometry from collection WKT with the given SRID. If SRID is not given, it defaults to 0.

## Synopsis


```sql
geometry ST_GeomCollFromText(text  WKT, integer  srid)
geometry ST_GeomCollFromText(text  WKT)
```


## Description


Makes a collection Geometry from the Well-Known-Text (WKT) representation with the given SRID. If SRID is not given, it defaults to 0.


OGC SPEC 3.2.6.2 - option SRID is from the conformance suite


Returns null if the WKT is not a GEOMETRYCOLLECTION


!!! note

    If you are absolutely sure all your WKT geometries are collections, don't use this function. It is slower than ST_GeomFromText since it adds an additional validation step.


 s3.2.6.2


## Examples


```sql
SELECT ST_GeomCollFromText('GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(1 2, 3 4))');
```


## See Also


[ST_GeomFromText](#ST_GeomFromText), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_GeomFromEWKT"></a>

# ST_GeomFromEWKT

Return a specified ST_Geometry value from Extended Well-Known Text representation (EWKT).

## Synopsis


```sql
geometry ST_GeomFromEWKT(text  EWKT)
```


## Description


Constructs a PostGIS ST_Geometry object from the OGC Extended Well-Known text (EWKT) representation.


!!! note

    The EWKT format is not an OGC standard, but an PostGIS specific format that includes the spatial reference system (SRID) identifier


Enhanced: 2.0.0 support for Polyhedral surfaces and TIN was introduced.


## Examples


```sql
SELECT ST_GeomFromEWKT('SRID=4269;LINESTRING(-71.160281 42.258729,-71.160837 42.259113,-71.161144 42.25932)');
SELECT ST_GeomFromEWKT('SRID=4269;MULTILINESTRING((-71.160281 42.258729,-71.160837 42.259113,-71.161144 42.25932))');

SELECT ST_GeomFromEWKT('SRID=4269;POINT(-71.064544 42.28787)');

SELECT ST_GeomFromEWKT('SRID=4269;POLYGON((-71.1776585052917 42.3902909739571,-71.1776820268866 42.3903701743239,
-71.1776063012595 42.3903825660754,-71.1775826583081 42.3903033653531,-71.1776585052917 42.3902909739571))');

SELECT ST_GeomFromEWKT('SRID=4269;MULTIPOLYGON(((-71.1031880899493 42.3152774590236,
-71.1031627617667 42.3152960829043,-71.102923838298 42.3149156848307,
-71.1023097974109 42.3151969047397,-71.1019285062273 42.3147384934248,
-71.102505233663 42.3144722937587,-71.10277487471 42.3141658254797,
-71.103113945163 42.3142739188902,-71.10324876416 42.31402489987,
-71.1033002961013 42.3140393340215,-71.1033488797549 42.3139495090772,
-71.103396240451 42.3138632439557,-71.1041521907712 42.3141153348029,
-71.1041411411543 42.3141545014533,-71.1041287795912 42.3142114839058,
-71.1041188134329 42.3142693656241,-71.1041112482575 42.3143272556118,
-71.1041072845732 42.3143851580048,-71.1041057218871 42.3144430686681,
-71.1041065602059 42.3145009876017,-71.1041097995362 42.3145589148055,
-71.1041166403905 42.3146168544148,-71.1041258822717 42.3146748022936,
-71.1041375307579 42.3147318674446,-71.1041492906949 42.3147711126569,
-71.1041598612795 42.314808571739,-71.1042515013869 42.3151287620809,
-71.1041173835118 42.3150739481917,-71.1040809891419 42.3151344119048,
-71.1040438678912 42.3151191367447,-71.1040194562988 42.3151832057859,
-71.1038734225584 42.3151140942995,-71.1038446938243 42.3151006300338,
-71.1038315271889 42.315094347535,-71.1037393329282 42.315054824985,
-71.1035447555574 42.3152608696313,-71.1033436658644 42.3151648370544,
-71.1032580383161 42.3152269126061,-71.103223066939 42.3152517403219,
-71.1031880899493 42.3152774590236)),
((-71.1043632495873 42.315113108546,-71.1043583974082 42.3151211109857,
-71.1043443253471 42.3150676015829,-71.1043850704575 42.3150793250568,-71.1043632495873 42.315113108546)))');
```


```

--3d circular string
SELECT ST_GeomFromEWKT('CIRCULARSTRING(220268 150415 1,220227 150505 2,220227 150406 3)');
```


```

--Polyhedral Surface example
SELECT ST_GeomFromEWKT('POLYHEDRALSURFACE(
	((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
	((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)),
	((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
	((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
	((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)),
	((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1))
)');
```


## See Also


[ST_AsEWKT](geometry-output.md#ST_AsEWKT), [ST_GeomFromText](#ST_GeomFromText)
  <a id="ST_GeomFromMARC21"></a>

# ST_GeomFromMARC21

Takes MARC21/XML geographic data as input and returns a PostGIS geometry object.

## Synopsis


```sql
geometry
                        ST_GeomFromMARC21(text
                        marcxml)
```


## Description


 This function creates a PostGIS geometry from a MARC21/XML record, which can contain a `POINT` or a `POLYGON`. In case of multiple geographic data entries in the same MARC21/XML record, a `MULTIPOINT` or `MULTIPOLYGON` will be returned. If the record contains mixed geometry types, a `GEOMETRYCOLLECTION` will be returned. It returns NULL if the MARC21/XML record does not contain any geographic data (datafield:034).


 LOC MARC21/XML versions supported:

- [MARC21/XML 1.1](https://www.loc.gov/standards/marcxml/)


Availability: 3.3.0, requires libxml2 2.6+


!!! note

    The MARC21/XML Coded Cartographic Mathematical Data currently does not provide any means to describe the Spatial Reference System of the encoded coordinates, so this function will always return a geometry with `SRID 0`.


!!! note

    Returned `POLYGON` geometries will always be clockwise oriented.


## Examples


Converting MARC21/XML geographic data containing a single `POINT` encoded as `hddd.dddddd`


```sql


                SELECT
                ST_AsText(
                    ST_GeomFromMARC21('
                        <record xmlns="http://www.loc.gov/MARC21/slim">
                            <leader>00000nz a2200000nc 4500</leader>
                            <controlfield tag="001">040277569</controlfield>
                            <datafield tag="034" ind1=" " ind2=" ">
                                <subfield code="d">W004.500000</subfield>
                                <subfield code="e">W004.500000</subfield>
                                <subfield code="f">N054.250000</subfield>
                                <subfield code="g">N054.250000</subfield>
                            </datafield>
                        </record>'));

                st_astext
                -------------------
                POINT(-4.5 54.25)
                (1 row)


```


Converting MARC21/XML geographic data containing a single `POLYGON` encoded as `hdddmmss`


```sql


                SELECT
                ST_AsText(
                    ST_GeomFromMARC21('
                        <record xmlns="http://www.loc.gov/MARC21/slim">
                            <leader>01062cem a2200241 a 4500</leader>
                            <controlfield tag="001">   84696781 </controlfield>
                            <datafield tag="034" ind1="1" ind2=" ">
                                <subfield code="a">a</subfield>
                                <subfield code="b">50000</subfield>
                                <subfield code="d">E0130600</subfield>
                                <subfield code="e">E0133100</subfield>
                                <subfield code="f">N0523900</subfield>
                                <subfield code="g">N0522300</subfield>
                            </datafield>
                        </record>'));

                st_astext
                -----------------------------------------------------------------------------------------------------------------------
                POLYGON((13.1 52.65,13.516666666666667 52.65,13.516666666666667 52.38333333333333,13.1 52.38333333333333,13.1 52.65))
                (1 row)


```


Converting MARC21/XML geographic data containing a `POLYGON` and a `POINT`:


```sql


                SELECT
                ST_AsText(
                    ST_GeomFromMARC21('
                <record xmlns="http://www.loc.gov/MARC21/slim">
                    <datafield tag="034" ind1="1" ind2=" ">
                        <subfield code="a">a</subfield>
                        <subfield code="b">50000</subfield>
                        <subfield code="d">E0130600</subfield>
                        <subfield code="e">E0133100</subfield>
                        <subfield code="f">N0523900</subfield>
                        <subfield code="g">N0522300</subfield>
                    </datafield>
                    <datafield tag="034" ind1=" " ind2=" ">
                        <subfield code="d">W004.500000</subfield>
                        <subfield code="e">W004.500000</subfield>
                        <subfield code="f">N054.250000</subfield>
                        <subfield code="g">N054.250000</subfield>
                    </datafield>
                </record>'));
                                                                                        st_astext
                -------------------------------------------------------------------------------------------------------------------------------------------------------------
                GEOMETRYCOLLECTION(POLYGON((13.1 52.65,13.516666666666667 52.65,13.516666666666667 52.38333333333333,13.1 52.38333333333333,13.1 52.65)),POINT(-4.5 54.25))
                (1 row)

```


## See Also


 [ST_AsMARC21](geometry-output.md#ST_AsMARC21)
  <a id="ST_GeometryFromText"></a>

# ST_GeometryFromText

Return a specified ST_Geometry value from Well-Known Text representation (WKT). This is an alias name for ST_GeomFromText

## Synopsis


```sql
geometry ST_GeometryFromText(text  WKT)
geometry ST_GeometryFromText(text  WKT, integer  srid)
```


## Description


 SQL-MM 3: 5.1.40


## See Also


[ST_GeomFromText](#ST_GeomFromText)
  <a id="ST_GeomFromText"></a>

# ST_GeomFromText

Return a specified ST_Geometry value from Well-Known Text representation (WKT).

## Synopsis


```sql
geometry ST_GeomFromText(text  WKT)
geometry ST_GeomFromText(text  WKT, integer  srid)
```


## Description


Constructs a PostGIS ST_Geometry object from the OGC Well-Known text representation.


!!! note

    There are two variants of ST_GeomFromText function. The first takes no SRID and returns a geometry with no defined spatial reference system (SRID=0). The second takes a SRID as the second argument and returns a geometry that includes this SRID as part of its metadata.


 s3.2.6.2 - option SRID is from the conformance suite.


 SQL-MM 3: 5.1.40


!!! note

    While not OGC-compliant, [ST_MakePoint](geometry-constructors.md#ST_MakePoint) is faster than ST_GeomFromText and ST_PointFromText. It is also easier to use for numeric coordinate values. [ST_Point](geometry-constructors.md#ST_Point) is another option similar in speed to [ST_MakePoint](geometry-constructors.md#ST_MakePoint) and is OGC-compliant, but doesn't support anything but 2D points.


!!! warning

    Changed: 2.0.0 In prior versions of PostGIS ST_GeomFromText('GEOMETRYCOLLECTION(EMPTY)') was allowed. This is now illegal in PostGIS 2.0.0 to better conform with SQL/MM standards. This should now be written as ST_GeomFromText('GEOMETRYCOLLECTION EMPTY')


## Examples


```sql
SELECT ST_GeomFromText('LINESTRING(-71.160281 42.258729,-71.160837 42.259113,-71.161144 42.25932)');
SELECT ST_GeomFromText('LINESTRING(-71.160281 42.258729,-71.160837 42.259113,-71.161144 42.25932)',4269);

SELECT ST_GeomFromText('MULTILINESTRING((-71.160281 42.258729,-71.160837 42.259113,-71.161144 42.25932))');

SELECT ST_GeomFromText('POINT(-71.064544 42.28787)');

SELECT ST_GeomFromText('POLYGON((-71.1776585052917 42.3902909739571,-71.1776820268866 42.3903701743239,
-71.1776063012595 42.3903825660754,-71.1775826583081 42.3903033653531,-71.1776585052917 42.3902909739571))');

SELECT ST_GeomFromText('MULTIPOLYGON(((-71.1031880899493 42.3152774590236,
-71.1031627617667 42.3152960829043,-71.102923838298 42.3149156848307,
-71.1023097974109 42.3151969047397,-71.1019285062273 42.3147384934248,
-71.102505233663 42.3144722937587,-71.10277487471 42.3141658254797,
-71.103113945163 42.3142739188902,-71.10324876416 42.31402489987,
-71.1033002961013 42.3140393340215,-71.1033488797549 42.3139495090772,
-71.103396240451 42.3138632439557,-71.1041521907712 42.3141153348029,
-71.1041411411543 42.3141545014533,-71.1041287795912 42.3142114839058,
-71.1041188134329 42.3142693656241,-71.1041112482575 42.3143272556118,
-71.1041072845732 42.3143851580048,-71.1041057218871 42.3144430686681,
-71.1041065602059 42.3145009876017,-71.1041097995362 42.3145589148055,
-71.1041166403905 42.3146168544148,-71.1041258822717 42.3146748022936,
-71.1041375307579 42.3147318674446,-71.1041492906949 42.3147711126569,
-71.1041598612795 42.314808571739,-71.1042515013869 42.3151287620809,
-71.1041173835118 42.3150739481917,-71.1040809891419 42.3151344119048,
-71.1040438678912 42.3151191367447,-71.1040194562988 42.3151832057859,
-71.1038734225584 42.3151140942995,-71.1038446938243 42.3151006300338,
-71.1038315271889 42.315094347535,-71.1037393329282 42.315054824985,
-71.1035447555574 42.3152608696313,-71.1033436658644 42.3151648370544,
-71.1032580383161 42.3152269126061,-71.103223066939 42.3152517403219,
-71.1031880899493 42.3152774590236)),
((-71.1043632495873 42.315113108546,-71.1043583974082 42.3151211109857,
-71.1043443253471 42.3150676015829,-71.1043850704575 42.3150793250568,-71.1043632495873 42.315113108546)))',4326);

SELECT ST_GeomFromText('CIRCULARSTRING(220268 150415,220227 150505,220227 150406)');

```


## See Also


[ST_GeomFromEWKT](#ST_GeomFromEWKT), [ST_GeomFromWKB](#ST_GeomFromWKB), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_LineFromText"></a>

# ST_LineFromText

Makes a Geometry from WKT representation with the given SRID. If SRID is not given, it defaults to 0.

## Synopsis


```sql
geometry ST_LineFromText(text  WKT)
geometry ST_LineFromText(text  WKT, integer  srid)
```


## Description


Makes a Geometry from WKT with the given SRID. If SRID is not given, it defaults to 0. If WKT passed in is not a LINESTRING, then null is returned.


!!! note

    OGC SPEC 3.2.6.2 - option SRID is from the conformance suite.


!!! note

    If you know all your geometries are LINESTRINGS, its more efficient to just use ST_GeomFromText. This just calls ST_GeomFromText and adds additional validation that it returns a linestring.


 s3.2.6.2


 SQL-MM 3: 7.2.8


## Examples


```sql
SELECT ST_LineFromText('LINESTRING(1 2, 3 4)') AS aline, ST_LineFromText('POINT(1 2)') AS null_return;
aline                            | null_return
------------------------------------------------
010200000002000000000000000000F ... | t

```


## See Also


[ST_GeomFromText](#ST_GeomFromText)
  <a id="ST_MLineFromText"></a>

# ST_MLineFromText

Return a specified ST_MultiLineString value from WKT representation.

## Synopsis


```sql
geometry ST_MLineFromText(text  WKT, integer  srid)
geometry ST_MLineFromText(text  WKT)
```


## Description


Makes a Geometry from Well-Known-Text (WKT) with the given SRID. If SRID is not given, it defaults to 0.


OGC SPEC 3.2.6.2 - option SRID is from the conformance suite


Returns null if the WKT is not a MULTILINESTRING


!!! note

    If you are absolutely sure all your WKT geometries are points, don't use this function. It is slower than ST_GeomFromText since it adds an additional validation step.


 s3.2.6.2


SQL-MM 3: 9.4.4


## Examples


```sql
SELECT ST_MLineFromText('MULTILINESTRING((1 2, 3 4), (4 5, 6 7))');
```


## See Also


[ST_GeomFromText](#ST_GeomFromText)
  <a id="ST_MPointFromText"></a>

# ST_MPointFromText

Makes a Geometry from WKT with the given SRID. If SRID is not given, it defaults to 0.

## Synopsis


```sql
geometry ST_MPointFromText(text  WKT, integer  srid)
geometry ST_MPointFromText(text  WKT)
```


## Description


Makes a Geometry from WKT with the given SRID. If SRID is not given, it defaults to 0.


OGC SPEC 3.2.6.2 - option SRID is from the conformance suite


Returns null if the WKT is not a MULTIPOINT


!!! note

    If you are absolutely sure all your WKT geometries are points, don't use this function. It is slower than ST_GeomFromText since it adds an additional validation step.


 3.2.6.2


 SQL-MM 3: 9.2.4


## Examples


```sql
SELECT ST_MPointFromText('MULTIPOINT((1 2),(3 4))');
SELECT ST_MPointFromText('MULTIPOINT((-70.9590 42.1180),(-70.9611 42.1223))', 4326);
```


## See Also


[ST_GeomFromText](#ST_GeomFromText)
  <a id="ST_MPolyFromText"></a>

# ST_MPolyFromText

Makes a MultiPolygon Geometry from WKT with the given SRID. If SRID is not given, it defaults to 0.

## Synopsis


```sql
geometry ST_MPolyFromText(text  WKT, integer  srid)
geometry ST_MPolyFromText(text  WKT)
```


## Description


Makes a MultiPolygon from WKT with the given SRID. If SRID is not given, it defaults to 0.


OGC SPEC 3.2.6.2 - option SRID is from the conformance suite


Throws an error if the WKT is not a MULTIPOLYGON


!!! note

    If you are absolutely sure all your WKT geometries are multipolygons, don't use this function. It is slower than ST_GeomFromText since it adds an additional validation step.


 s3.2.6.2


 SQL-MM 3: 9.6.4


## Examples


```sql
SELECT ST_MPolyFromText('MULTIPOLYGON(((0 0 1,20 0 1,20 20 1,0 20 1,0 0 1),(5 5 3,5 7 3,7 7 3,7 5 3,5 5 3)))');
SELECt ST_MPolyFromText('MULTIPOLYGON(((-70.916 42.1002,-70.9468 42.0946,-70.9765 42.0872,-70.9754 42.0875,-70.9749 42.0879,-70.9752 42.0881,-70.9754 42.0891,-70.9758 42.0894,-70.9759 42.0897,-70.9759 42.0899,-70.9754 42.0902,-70.9756 42.0906,-70.9753 42.0907,-70.9753 42.0917,-70.9757 42.0924,-70.9755 42.0928,-70.9755 42.0942,-70.9751 42.0948,-70.9755 42.0953,-70.9751 42.0958,-70.9751 42.0962,-70.9759 42.0983,-70.9767 42.0987,-70.9768 42.0991,-70.9771 42.0997,-70.9771 42.1003,-70.9768 42.1005,-70.977 42.1011,-70.9766 42.1019,-70.9768 42.1026,-70.9769 42.1033,-70.9775 42.1042,-70.9773 42.1043,-70.9776 42.1043,-70.9778 42.1048,-70.9773 42.1058,-70.9774 42.1061,-70.9779 42.1065,-70.9782 42.1078,-70.9788 42.1085,-70.9798 42.1087,-70.9806 42.109,-70.9807 42.1093,-70.9806 42.1099,-70.9809 42.1109,-70.9808 42.1112,-70.9798 42.1116,-70.9792 42.1127,-70.979 42.1129,-70.9787 42.1134,-70.979 42.1139,-70.9791 42.1141,-70.9987 42.1116,-71.0022 42.1273,
	-70.9408 42.1513,-70.9315 42.1165,-70.916 42.1002)))',4326);
```


## See Also


[ST_GeomFromText](#ST_GeomFromText), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_PointFromText"></a>

# ST_PointFromText

Makes a point Geometry from WKT with the given SRID. If SRID is not given, it defaults to unknown.

## Synopsis


```sql
geometry ST_PointFromText(text  WKT)
geometry ST_PointFromText(text  WKT, integer  srid)
```


## Description


Constructs a PostGIS ST_Geometry point object from the OGC Well-Known text representation. If SRID is not given, it defaults to unknown (currently 0). If geometry is not a WKT point representation, returns null. If completely invalid WKT, then throws an error.


!!! note

    There are 2 variants of ST_PointFromText function, the first takes no SRID and returns a geometry with no defined spatial reference system. The second takes a spatial reference id as the second argument and returns an ST_Geometry that includes this srid as part of its meta-data. The srid must be defined in the spatial_ref_sys table.


!!! note

    If you are absolutely sure all your WKT geometries are points, don't use this function. It is slower than ST_GeomFromText since it adds an additional validation step. If you are building points from long lat coordinates and care more about performance and accuracy than OGC compliance, use [ST_MakePoint](geometry-constructors.md#ST_MakePoint) or OGC compliant alias [ST_Point](geometry-constructors.md#ST_Point).


 s3.2.6.2 - option SRID is from the conformance suite.


 SQL-MM 3: 6.1.8


## Examples


```sql

SELECT ST_PointFromText('POINT(-71.064544 42.28787)');
SELECT ST_PointFromText('POINT(-71.064544 42.28787)', 4326);

```


## See Also


[ST_GeomFromText](#ST_GeomFromText), [ST_MakePoint](geometry-constructors.md#ST_MakePoint), [ST_Point](geometry-constructors.md#ST_Point), [ST_SRID](spatial-reference-system-functions.md#ST_SRID)
  <a id="ST_PolygonFromText"></a>

# ST_PolygonFromText

Makes a Geometry from WKT with the given SRID. If SRID is not given, it defaults to 0.

## Synopsis


```sql
geometry ST_PolygonFromText(text  WKT)
geometry ST_PolygonFromText(text  WKT, integer  srid)
```


## Description


Makes a Geometry from WKT with the given SRID. If SRID is not given, it defaults to 0. Returns null if WKT is not a polygon.


OGC SPEC 3.2.6.2 - option SRID is from the conformance suite


!!! note

    If you are absolutely sure all your WKT geometries are polygons, don't use this function. It is slower than ST_GeomFromText since it adds an additional validation step.


 s3.2.6.2


 SQL-MM 3: 8.3.6


## Examples


```sql
SELECT ST_PolygonFromText('POLYGON((-71.1776585052917 42.3902909739571,-71.1776820268866 42.3903701743239,
-71.1776063012595 42.3903825660754,-71.1775826583081 42.3903033653531,-71.1776585052917 42.3902909739571))');
st_polygonfromtext
------------------
010300000001000000050000006...


SELECT ST_PolygonFromText('POINT(1 2)') IS NULL as point_is_notpoly;

point_is_not_poly
----------
t
```


## See Also


[ST_GeomFromText](#ST_GeomFromText)
  <a id="ST_WKTToSQL"></a>

# ST_WKTToSQL

Return a specified ST_Geometry value from Well-Known Text representation (WKT). This is an alias name for ST_GeomFromText

## Synopsis


```sql
geometry ST_WKTToSQL(text  WKT)
```


## Description


 SQL-MM 3: 5.1.34


## See Also


[ST_GeomFromText](#ST_GeomFromText)


## Well-Known Binary (WKB)
  <a id="ST_GeogFromWKB"></a>

# ST_GeogFromWKB

Creates a geography instance from a Well-Known Binary geometry representation (WKB) or extended Well Known Binary (EWKB).

## Synopsis


```sql
geography ST_GeogFromWKB(bytea  wkb)
```


## Description


The `ST_GeogFromWKB` function, takes a well-known binary representation (WKB) of a geometry or PostGIS Extended WKB and creates an instance of the appropriate geography type. This function plays the role of the Geometry Factory in SQL.


If SRID is not specified, it defaults to 4326 (WGS 84 long lat).


## Examples


```
--Although bytea rep contains single \, these need to be escaped when inserting into a table
SELECT ST_AsText(
ST_GeogFromWKB(E'\\001\\002\\000\\000\\000\\002\\000\\000\\000\\037\\205\\353Q\\270~\\\\\\300\\323Mb\\020X\\231C@\\020X9\\264\\310~\\\\\\300)\\\\\\217\\302\\365\\230C@')
);
					  st_astext
------------------------------------------------------
 LINESTRING(-113.98 39.198,-113.981 39.195)
(1 row)
```


## See Also


[ST_GeogFromText](#ST_GeogFromText), [ST_AsBinary](geometry-output.md#ST_AsBinary)
  <a id="ST_GeomFromEWKB"></a>

# ST_GeomFromEWKB

Return a specified ST_Geometry value from Extended Well-Known Binary representation (EWKB).

## Synopsis


```sql
geometry ST_GeomFromEWKB(bytea  EWKB)
```


## Description


Constructs a PostGIS ST_Geometry object from the OGC Extended Well-Known binary (EWKT) representation.


!!! note

    The EWKB format is not an OGC standard, but a PostGIS specific format that includes the spatial reference system (SRID) identifier


Enhanced: 2.0.0 support for Polyhedral surfaces and TIN was introduced.


## Examples


line string binary rep 0f LINESTRING(-71.160281 42.258729,-71.160837 42.259113,-71.161144 42.25932) in NAD 83 long lat (4269).


!!! note

    NOTE: Even though byte arrays are delimited with \ and may have ', we need to escape both out with \ and '' if standard_conforming_strings is off. So it does not look exactly like its AsEWKB representation.


```sql
SELECT ST_GeomFromEWKB(E'\\001\\002\\000\\000 \\255\\020\\000\\000\\003\\000\\000\\000\\344J=
\\013B\\312Q\\300n\\303(\\010\\036!E@''\\277E''K
\\312Q\\300\\366{b\\235*!E@\\225|\\354.P\\312Q
\\300p\\231\\323e1!E@');
```


!!! note

    In PostgreSQL 9.1+ - standard_conforming_strings is set to on by default, where as in past versions it was set to off. You can change defaults as needed for a single query or at the database or server level. Below is how you would do it with standard_conforming_strings = on. In this case we escape the ' with standard ansi ', but slashes are not escaped


```

	    set standard_conforming_strings = on;
SELECT ST_GeomFromEWKB('\001\002\000\000 \255\020\000\000\003\000\000\000\344J=\012\013B
    \312Q\300n\303(\010\036!E@''\277E''K\012\312Q\300\366{b\235*!E@\225|\354.P\312Q\012\300p\231\323e1')
```


## See Also


[ST_AsBinary](geometry-output.md#ST_AsBinary), [ST_AsEWKB](geometry-output.md#ST_AsEWKB), [ST_GeomFromWKB](#ST_GeomFromWKB)
  <a id="ST_GeomFromWKB"></a>

# ST_GeomFromWKB

Creates a geometry instance from a Well-Known Binary geometry representation (WKB) and optional SRID.

## Synopsis


```sql
geometry ST_GeomFromWKB(bytea  geom)
geometry ST_GeomFromWKB(bytea  geom, integer  srid)
```


## Description


The `ST_GeomFromWKB` function, takes a well-known binary representation of a geometry and a Spatial Reference System ID (`SRID`) and creates an instance of the appropriate geometry type. This function plays the role of the Geometry Factory in SQL. This is an alternate name for ST_WKBToSQL.


If SRID is not specified, it defaults to 0 (Unknown).


 s3.2.7.2 - the optional SRID is from the conformance suite


 SQL-MM 3: 5.1.41


## Examples


```
--Although bytea rep contains single \, these need to be escaped when inserting into a table
		-- unless standard_conforming_strings is set to on.
SELECT ST_AsEWKT(
ST_GeomFromWKB(E'\\001\\002\\000\\000\\000\\002\\000\\000\\000\\037\\205\\353Q\\270~\\\\\\300\\323Mb\\020X\\231C@\\020X9\\264\\310~\\\\\\300)\\\\\\217\\302\\365\\230C@',4326)
);
					  st_asewkt
------------------------------------------------------
 SRID=4326;LINESTRING(-113.98 39.198,-113.981 39.195)
(1 row)

SELECT
  ST_AsText(
	ST_GeomFromWKB(
	  ST_AsEWKB('POINT(2 5)'::geometry)
	)
  );
 st_astext
------------
 POINT(2 5)
(1 row)
```


## See Also


[ST_WKBToSQL](#ST_WKBToSQL), [ST_AsBinary](geometry-output.md#ST_AsBinary), [ST_GeomFromEWKB](#ST_GeomFromEWKB)
  <a id="ST_LineFromWKB"></a>

# ST_LineFromWKB

Makes a `LINESTRING` from WKB with the given SRID

## Synopsis


```sql
geometry ST_LineFromWKB(bytea  WKB)
geometry ST_LineFromWKB(bytea  WKB, integer  srid)
```


## Description


The `ST_LineFromWKB` function, takes a well-known binary representation of geometry and a Spatial Reference System ID (`SRID`) and creates an instance of the appropriate geometry type - in this case, a `LINESTRING` geometry. This function plays the role of the Geometry Factory in SQL.


If an SRID is not specified, it defaults to 0. `NULL` is returned if the input `bytea` does not represent a `LINESTRING`.


!!! note

    OGC SPEC 3.2.6.2 - option SRID is from the conformance suite.


!!! note

    If you know all your geometries are `LINESTRING`s, its more efficient to just use [ST_GeomFromWKB](#ST_GeomFromWKB). This function just calls [ST_GeomFromWKB](#ST_GeomFromWKB) and adds additional validation that it returns a linestring.


 s3.2.6.2


 SQL-MM 3: 7.2.9


## Examples


```sql
SELECT ST_LineFromWKB(ST_AsBinary(ST_GeomFromText('LINESTRING(1 2, 3 4)'))) AS aline,
		ST_LineFromWKB(ST_AsBinary(ST_GeomFromText('POINT(1 2)'))) IS NULL AS null_return;
aline                            | null_return
------------------------------------------------
010200000002000000000000000000F ... | t

```


## See Also


[ST_GeomFromWKB](#ST_GeomFromWKB), [ST_LinestringFromWKB](#ST_LinestringFromWKB)
  <a id="ST_LinestringFromWKB"></a>

# ST_LinestringFromWKB

Makes a geometry from WKB with the given SRID.

## Synopsis


```sql
geometry ST_LinestringFromWKB(bytea  WKB)
geometry ST_LinestringFromWKB(bytea  WKB, integer  srid)
```


## Description


The `ST_LinestringFromWKB` function, takes a well-known binary representation of geometry and a Spatial Reference System ID (`SRID`) and creates an instance of the appropriate geometry type - in this case, a `LINESTRING` geometry. This function plays the role of the Geometry Factory in SQL.


If an SRID is not specified, it defaults to 0. `NULL` is returned if the input `bytea` does not represent a `LINESTRING` geometry. This an alias for [ST_LineFromWKB](#ST_LineFromWKB).


!!! note

    OGC SPEC 3.2.6.2 - optional SRID is from the conformance suite.


!!! note

    If you know all your geometries are `LINESTRING`s, it's more efficient to just use [ST_GeomFromWKB](#ST_GeomFromWKB). This function just calls [ST_GeomFromWKB](#ST_GeomFromWKB) and adds additional validation that it returns a `LINESTRING`.


 s3.2.6.2


 SQL-MM 3: 7.2.9


## Examples


```sql
SELECT
  ST_LineStringFromWKB(
	ST_AsBinary(ST_GeomFromText('LINESTRING(1 2, 3 4)'))
  ) AS aline,
  ST_LinestringFromWKB(
	ST_AsBinary(ST_GeomFromText('POINT(1 2)'))
  ) IS NULL AS null_return;
   aline                            | null_return
------------------------------------------------
010200000002000000000000000000F ... | t
```


## See Also


[ST_GeomFromWKB](#ST_GeomFromWKB), [ST_LineFromWKB](#ST_LineFromWKB)
  <a id="ST_PointFromWKB"></a>

# ST_PointFromWKB

Makes a geometry from WKB with the given SRID

## Synopsis


```sql
geometry ST_GeomFromWKB(bytea  geom)
geometry ST_GeomFromWKB(bytea  geom, integer  srid)
```


## Description


The `ST_PointFromWKB` function, takes a well-known binary representation of geometry and a Spatial Reference System ID (`SRID`) and creates an instance of the appropriate geometry type - in this case, a `POINT` geometry. This function plays the role of the Geometry Factory in SQL.


If an SRID is not specified, it defaults to 0. `NULL` is returned if the input `bytea` does not represent a `POINT` geometry.


 s3.2.7.2


 SQL-MM 3: 6.1.9


## Examples


```sql
SELECT
  ST_AsText(
	ST_PointFromWKB(
	  ST_AsEWKB('POINT(2 5)'::geometry)
	)
  );
 st_astext
------------
 POINT(2 5)
(1 row)

SELECT
  ST_AsText(
	ST_PointFromWKB(
	  ST_AsEWKB('LINESTRING(2 5, 2 6)'::geometry)
	)
  );
 st_astext
-----------

(1 row)
```


## See Also


[ST_GeomFromWKB](#ST_GeomFromWKB), [ST_LineFromWKB](#ST_LineFromWKB)
  <a id="ST_WKBToSQL"></a>

# ST_WKBToSQL

Return a specified ST_Geometry value from Well-Known Binary representation (WKB). This is an alias name for ST_GeomFromWKB that takes no srid

## Synopsis


```sql
geometry ST_WKBToSQL(bytea  WKB)
```


## Description


 SQL-MM 3: 5.1.36


## See Also


[ST_GeomFromWKB](#ST_GeomFromWKB)


## Other Formats
  <a id="ST_Box2dFromGeoHash"></a>

# ST_Box2dFromGeoHash

Return a BOX2D from a GeoHash string.

## Synopsis


```sql
box2d ST_Box2dFromGeoHash(text  geohash, integer  precision=full_precision_of_geohash)
```


## Description


Return a BOX2D from a GeoHash string.


If no `precision` is specified ST_Box2dFromGeoHash returns a BOX2D based on full precision of the input GeoHash string.


If `precision` is specified ST_Box2dFromGeoHash will use that many characters from the GeoHash to create the BOX2D. Lower precision values results in larger BOX2Ds and larger values increase the precision.


Availability: 2.1.0


## Examples


```sql
SELECT ST_Box2dFromGeoHash('9qqj7nmxncgyy4d0dbxqz0');

                st_geomfromgeohash
--------------------------------------------------
 BOX(-115.172816 36.114646,-115.172816 36.114646)

SELECT ST_Box2dFromGeoHash('9qqj7nmxncgyy4d0dbxqz0', 0);

 st_box2dfromgeohash
----------------------
 BOX(-180 -90,180 90)

 SELECT ST_Box2dFromGeoHash('9qqj7nmxncgyy4d0dbxqz0', 10);
                            st_box2dfromgeohash
---------------------------------------------------------------------------
 BOX(-115.17282128334 36.1146408319473,-115.172810554504 36.1146461963654)


```


## See Also


[ST_GeoHash](geometry-output.md#ST_GeoHash), [ST_GeomFromGeoHash](#ST_GeomFromGeoHash), [ST_PointFromGeoHash](#ST_PointFromGeoHash)
  <a id="ST_GeomFromGeoHash"></a>

# ST_GeomFromGeoHash

Return a geometry from a GeoHash string.

## Synopsis


```sql
geometry ST_GeomFromGeoHash(text  geohash, integer  precision=full_precision_of_geohash)
```


## Description


Return a geometry from a GeoHash string. The geometry will be a polygon representing the GeoHash bounds.


If no `precision` is specified ST_GeomFromGeoHash returns a polygon based on full precision of the input GeoHash string.


If `precision` is specified ST_GeomFromGeoHash will use that many characters from the GeoHash to create the polygon.


Availability: 2.1.0


## Examples


```sql
SELECT ST_AsText(ST_GeomFromGeoHash('9qqj7nmxncgyy4d0dbxqz0'));
                                                        st_astext
--------------------------------------------------------------------------------------------------------------------------
 POLYGON((-115.172816 36.114646,-115.172816 36.114646,-115.172816 36.114646,-115.172816 36.114646,-115.172816 36.114646))

SELECT ST_AsText(ST_GeomFromGeoHash('9qqj7nmxncgyy4d0dbxqz0', 4));
                                                          st_astext
------------------------------------------------------------------------------------------------------------------------------
 POLYGON((-115.3125 36.03515625,-115.3125 36.2109375,-114.9609375 36.2109375,-114.9609375 36.03515625,-115.3125 36.03515625))

SELECT ST_AsText(ST_GeomFromGeoHash('9qqj7nmxncgyy4d0dbxqz0', 10));
                                                                                       st_astext
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
 POLYGON((-115.17282128334 36.1146408319473,-115.17282128334 36.1146461963654,-115.172810554504 36.1146461963654,-115.172810554504 36.1146408319473,-115.17282128334 36.1146408319473))


```


## See Also


[ST_GeoHash](geometry-output.md#ST_GeoHash),[ST_Box2dFromGeoHash](#ST_Box2dFromGeoHash), [ST_PointFromGeoHash](#ST_PointFromGeoHash)
  <a id="ST_GeomFromGML"></a>

# ST_GeomFromGML

Takes as input GML representation of geometry and outputs a PostGIS geometry object

## Synopsis


```sql
geometry ST_GeomFromGML(text  geomgml)
geometry ST_GeomFromGML(text  geomgml, integer  srid)
```


## Description


Constructs a PostGIS ST_Geometry object from the OGC GML representation.


ST_GeomFromGML works only for GML Geometry fragments. It throws an error if you try to use it on a whole GML document.


 OGC GML versions supported:

- GML 3.2.1 Namespace
- GML 3.1.1 Simple Features profile SF-2 (with GML 3.1.0 and 3.0.0 backward compatibility)
- GML 2.1.2
 OGC GML standards, cf: [http://www.opengeospatial.org/standards/gml](http://www.opengeospatial.org/standards/gml):


Availability: 1.5, requires libxml2 1.6+


Enhanced: 2.0.0 support for Polyhedral surfaces and TIN was introduced.


Enhanced: 2.0.0 default srid optional parameter added.


GML allow mixed dimensions (2D and 3D inside the same MultiGeometry for instance). As PostGIS geometries don't, ST_GeomFromGML convert the whole geometry to 2D if a missing Z dimension is found once.


GML support mixed SRS inside the same MultiGeometry. As PostGIS geometries don't, ST_GeomFromGML, in this case, reproject all subgeometries to the SRS root node. If no srsName attribute available for the GML root node, the function throw an error.


ST_GeomFromGML function is not pedantic about an explicit GML namespace. You could avoid to mention it explicitly for common usages. But you need it if you want to use XLink feature inside GML.


!!! note

    ST_GeomFromGML function not support SQL/MM curves geometries.


## Examples - A single geometry with srsName


```sql

SELECT ST_GeomFromGML($$
    <gml:LineString xmlns:gml="http://www.opengis.net/gml"
			srsName="EPSG:4269">
        <gml:coordinates>
            -71.16028,42.258729 -71.160837,42.259112 -71.161143,42.25932
        </gml:coordinates>
    </gml:LineString>
$$);


```


## Examples - XLink usage


```sql

SELECT ST_GeomFromGML($$
    <gml:LineString xmlns:gml="http://www.opengis.net/gml"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            srsName="urn:ogc:def:crs:EPSG::4269">
        <gml:pointProperty>
            <gml:Point gml:id="p1"><gml:pos>42.258729 -71.16028</gml:pos></gml:Point>
        </gml:pointProperty>
        <gml:pos>42.259112 -71.160837</gml:pos>
        <gml:pointProperty>
            <gml:Point xlink:type="simple" xlink:href="#p1"/>
        </gml:pointProperty>
    </gml:LineString>
$$);


```


## Examples - Polyhedral Surface


```sql

SELECT ST_AsEWKT(ST_GeomFromGML('
<gml:PolyhedralSurface xmlns:gml="http://www.opengis.net/gml">
<gml:polygonPatches>
  <gml:PolygonPatch>
    <gml:exterior>
      <gml:LinearRing><gml:posList srsDimension="3">0 0 0 0 0 1 0 1 1 0 1 0 0 0 0</gml:posList></gml:LinearRing>
    </gml:exterior>
  </gml:PolygonPatch>
  <gml:PolygonPatch>
    <gml:exterior>
	<gml:LinearRing><gml:posList srsDimension="3">0 0 0 0 1 0 1 1 0 1 0 0 0 0 0</gml:posList></gml:LinearRing>
    </gml:exterior>
  </gml:PolygonPatch>
  <gml:PolygonPatch>
    <gml:exterior>
	<gml:LinearRing><gml:posList srsDimension="3">0 0 0 1 0 0 1 0 1 0 0 1 0 0 0</gml:posList></gml:LinearRing>
    </gml:exterior>
  </gml:PolygonPatch>
  <gml:PolygonPatch>
    <gml:exterior>
	<gml:LinearRing><gml:posList srsDimension="3">1 1 0 1 1 1 1 0 1 1 0 0 1 1 0</gml:posList></gml:LinearRing>
    </gml:exterior>
  </gml:PolygonPatch>
  <gml:PolygonPatch>
    <gml:exterior>
	<gml:LinearRing><gml:posList srsDimension="3">0 1 0 0 1 1 1 1 1 1 1 0 0 1 0</gml:posList></gml:LinearRing>
    </gml:exterior>
  </gml:PolygonPatch>
  <gml:PolygonPatch>
    <gml:exterior>
	<gml:LinearRing><gml:posList srsDimension="3">0 0 1 1 0 1 1 1 1 0 1 1 0 0 1</gml:posList></gml:LinearRing>
    </gml:exterior>
  </gml:PolygonPatch>
</gml:polygonPatches>
</gml:PolyhedralSurface>'));

-- result --
 POLYHEDRALSURFACE(((0 0 0,0 0 1,0 1 1,0 1 0,0 0 0)),
 ((0 0 0,0 1 0,1 1 0,1 0 0,0 0 0)),
 ((0 0 0,1 0 0,1 0 1,0 0 1,0 0 0)),
 ((1 1 0,1 1 1,1 0 1,1 0 0,1 1 0)),
 ((0 1 0,0 1 1,1 1 1,1 1 0,0 1 0)),
 ((0 0 1,1 0 1,1 1 1,0 1 1,0 0 1)))

```


## See Also


[Build configuration](../postgis-installation/compiling-and-install-from-source.md#installation_configuration), [ST_AsGML](geometry-output.md#ST_AsGML), [ST_GMLToSQL](#ST_GMLToSQL)
  <a id="ST_GeomFromGeoJSON"></a>

# ST_GeomFromGeoJSON

Takes as input a geojson representation of a geometry and outputs a PostGIS geometry object

## Synopsis


```sql
geometry ST_GeomFromGeoJSON(text  geomjson)
geometry ST_GeomFromGeoJSON(json  geomjson)
geometry ST_GeomFromGeoJSON(jsonb  geomjson)
```


## Description


Constructs a PostGIS geometry object from the GeoJSON representation.


ST_GeomFromGeoJSON works only for JSON Geometry fragments. It throws an error if you try to use it on a whole JSON document.


Enhanced: 3.0.0 parsed geometry defaults to SRID=4326 if not specified otherwise.


Enhanced: 2.5.0 can now accept json and jsonb as inputs.


Availability: 2.0.0 requires - JSON-C >= 0.9


!!! note

    If you do not have JSON-C enabled, support you will get an error notice instead of seeing an output. To enable JSON-C, run configure --with-jsondir=/path/to/json-c. See [Build configuration](../postgis-installation/compiling-and-install-from-source.md#installation_configuration) for details.


## Examples


```sql
SELECT ST_AsText(ST_GeomFromGeoJSON('{"type":"Point","coordinates":[-48.23456,20.12345]}')) As wkt;
wkt
------
POINT(-48.23456 20.12345)
```


```
-- a 3D linestring
SELECT ST_AsText(ST_GeomFromGeoJSON('{"type":"LineString","coordinates":[[1,2,3],[4,5,6],[7,8,9]]}')) As wkt;

wkt
-------------------
LINESTRING(1 2,4 5,7 8)
```


## See Also


[ST_AsText](geometry-output.md#ST_AsText), [ST_AsGeoJSON](geometry-output.md#ST_AsGeoJSON), [Build configuration](../postgis-installation/compiling-and-install-from-source.md#installation_configuration)
  <a id="ST_GeomFromKML"></a>

# ST_GeomFromKML

Takes as input KML representation of geometry and outputs a PostGIS geometry object

## Synopsis


```sql
geometry ST_GeomFromKML(text  geomkml)
```


## Description


Constructs a PostGIS ST_Geometry object from the OGC KML representation.


ST_GeomFromKML works only for KML Geometry fragments. It throws an error if you try to use it on a whole KML document.


 OGC KML versions supported:

- KML 2.2.0 Namespace
 OGC KML standards, cf: [http://www.opengeospatial.org/standards/kml](http://www.opengeospatial.org/standards/kml):


Availability: 1.5, requires libxml2 2.6+


!!! note

    ST_GeomFromKML function not support SQL/MM curves geometries.


## Examples - A single geometry with srsName


```sql

SELECT ST_GeomFromKML($$
    <LineString>
        <coordinates>-71.1663,42.2614
            -71.1667,42.2616</coordinates>
    </LineString>
$$);


```


## See Also


[Build configuration](../postgis-installation/compiling-and-install-from-source.md#installation_configuration), [ST_AsKML](geometry-output.md#ST_AsKML)
  <a id="ST_GeomFromTWKB"></a>

# ST_GeomFromTWKB

Creates a geometry instance from a TWKB ("[Tiny Well-Known Binary](https://github.com/TWKB/Specification/blob/master/twkb.md)") geometry representation.

## Synopsis


```sql
geometry ST_GeomFromTWKB(bytea  twkb)
```


## Description


The `ST_GeomFromTWKB` function, takes a a TWKB ("[Tiny Well-Known Binary](https://github.com/TWKB/Specification/blob/master/twkb.md)") geometry representation (WKB) and creates an instance of the appropriate geometry type.


## Examples


```sql

SELECT ST_AsText(ST_GeomFromTWKB(ST_AsTWKB('LINESTRING(126 34, 127 35)'::geometry)));

         st_astext
-----------------------------
 LINESTRING(126 34, 127 35)
(1 row)


SELECT ST_AsEWKT(
  ST_GeomFromTWKB(E'\\x620002f7f40dbce4040105')
);
					  st_asewkt
------------------------------------------------------
LINESTRING(-113.98 39.198,-113.981 39.195)
(1 row)
```


## See Also


[ST_AsTWKB](geometry-output.md#ST_AsTWKB)
  <a id="ST_GMLToSQL"></a>

# ST_GMLToSQL

Return a specified ST_Geometry value from GML representation. This is an alias name for ST_GeomFromGML

## Synopsis


```sql
geometry ST_GMLToSQL(text  geomgml)
geometry ST_GMLToSQL(text  geomgml, integer  srid)
```


## Description


 SQL-MM 3: 5.1.50 (except for curves support).


Availability: 1.5, requires libxml2 1.6+


Enhanced: 2.0.0 support for Polyhedral surfaces and TIN was introduced.


Enhanced: 2.0.0 default srid optional parameter added.


## See Also


[Build configuration](../postgis-installation/compiling-and-install-from-source.md#installation_configuration), [ST_GeomFromGML](#ST_GeomFromGML), [ST_AsGML](geometry-output.md#ST_AsGML)
  <a id="ST_LineFromEncodedPolyline"></a>

# ST_LineFromEncodedPolyline

Creates a LineString from an Encoded Polyline.

## Synopsis


```sql
geometry ST_LineFromEncodedPolyline(text  polyline, integer  precision=5)
```


## Description


Creates a LineString from an Encoded Polyline string.


Optional `precision` specifies how many decimal places will be preserved in Encoded Polyline. Value should be the same on encoding and decoding, or coordinates will be incorrect.


See http://developers.google.com/maps/documentation/utilities/polylinealgorithm


Availability: 2.2.0


## Examples


```

-- Create a line string from a polyline
SELECT ST_AsEWKT(ST_LineFromEncodedPolyline('_p~iF~ps|U_ulLnnqC_mqNvxq`@'));
-- result --
SRID=4326;LINESTRING(-120.2 38.5,-120.95 40.7,-126.453 43.252)

-- Select different precision that was used for polyline encoding
SELECT ST_AsEWKT(ST_LineFromEncodedPolyline('_p~iF~ps|U_ulLnnqC_mqNvxq`@',6));
-- result --
SRID=4326;LINESTRING(-12.02 3.85,-12.095 4.07,-12.6453 4.3252)


```


## See Also


[ST_AsEncodedPolyline](geometry-output.md#ST_AsEncodedPolyline)
  <a id="ST_PointFromGeoHash"></a>

# ST_PointFromGeoHash

Return a point from a GeoHash string.

## Synopsis


```sql
point ST_PointFromGeoHash(text  geohash, integer  precision=full_precision_of_geohash)
```


## Description


Return a point from a GeoHash string. The point represents the center point of the GeoHash.


If no `precision` is specified ST_PointFromGeoHash returns a point based on full precision of the input GeoHash string.


If `precision` is specified ST_PointFromGeoHash will use that many characters from the GeoHash to create the point.


Availability: 2.1.0


## Examples


```sql
SELECT ST_AsText(ST_PointFromGeoHash('9qqj7nmxncgyy4d0dbxqz0'));
          st_astext
------------------------------
 POINT(-115.172816 36.114646)

SELECT ST_AsText(ST_PointFromGeoHash('9qqj7nmxncgyy4d0dbxqz0', 4));
             st_astext
-----------------------------------
 POINT(-115.13671875 36.123046875)

SELECT ST_AsText(ST_PointFromGeoHash('9qqj7nmxncgyy4d0dbxqz0', 10));
                 st_astext
-------------------------------------------
 POINT(-115.172815918922 36.1146435141563)


```


## See Also


 [ST_GeoHash](geometry-output.md#ST_GeoHash), [ST_Box2dFromGeoHash](#ST_Box2dFromGeoHash), [ST_GeomFromGeoHash](#ST_GeomFromGeoHash)
  <a id="ST_FromFlatGeobufToTable"></a>

# ST_FromFlatGeobufToTable

Creates a table based on the structure of FlatGeobuf data.

## Synopsis


```sql
void ST_FromFlatGeobufToTable(text  schemaname, text  tablename, bytea  FlatGeobuf input data)
```


## Description


 Creates a table based on the structure of FlatGeobuf data. ([http://flatgeobuf.org](http://flatgeobuf.org)).


`schema` Schema name.


`table` Table name.


`data` Input FlatGeobuf data.


Availability: 3.2.0
  <a id="ST_FromFlatGeobuf"></a>

# ST_FromFlatGeobuf

Reads FlatGeobuf data.

## Synopsis


```sql
setof anyelement ST_FromFlatGeobuf(anyelement  Table reference, bytea  FlatGeobuf input data)
```


## Description


 Reads FlatGeobuf data ([http://flatgeobuf.org](http://flatgeobuf.org)). NOTE: PostgreSQL bytea cannot exceed 1GB.


`tabletype` reference to a table type.


`data` input FlatGeobuf data.


Availability: 3.2.0
