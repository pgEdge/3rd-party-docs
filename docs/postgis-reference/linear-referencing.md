<a id="Linear_Referencing"></a>

## Linear Referencing
  <a id="ST_LineInterpolatePoint"></a>

# ST_LineInterpolatePoint

Returns a point interpolated along a line at a fractional location.

## Synopsis


```sql
geometry ST_LineInterpolatePoint(geometry  a_linestring, float8  a_fraction)
geography ST_LineInterpolatePoint(geography  a_linestring, float8  a_fraction, boolean use_spheroid = true)
```


## Description


Returns a point interpolated along a line at a fractional location. First argument must be a LINESTRING. Second argument is a float between 0 and 1 representing the fraction of line length where the point is to be located. The Z and M values are interpolated if present.


See [ST_LineLocatePoint](#ST_LineLocatePoint) for computing the line location nearest to a Point.


!!! note

    This function computes points in 2D and then interpolates values for Z and M, while [ST_3DLineInterpolatePoint](#ST_3DLineInterpolatePoint) computes points in 3D and only interpolates the M value.


!!! note

    Since release 1.1.1 this function also interpolates M and Z values (when present), while prior releases set them to 0.0.


Availability: 0.8.2, Z and M supported added in 1.1.1


Changed: 2.1.0. Up to 2.0.x this was called ST_Line_Interpolate_Point.


## Examples


![image](images/st_line_interpolate_point01.png)


A LineString with the interpolated point at 20% position (0.20)


```
-- The point 20% along a line

SELECT ST_AsEWKT(  ST_LineInterpolatePoint(
        'LINESTRING(25 50, 100 125, 150 190)',
        0.2 ));
----------------
 POINT(51.5974135047432 76.5974135047432)
```


The mid-point of a 3D line:


```sql

SELECT ST_AsEWKT(  ST_LineInterpolatePoint('
        LINESTRING(1 2 3, 4 5 6, 6 7 8)',
        0.5 ));
--------------------
 POINT(3.5 4.5 5.5)
```


The closest point on a line to a point:


```sql

SELECT ST_AsText( ST_LineInterpolatePoint( line.geom,
                      ST_LineLocatePoint( line.geom, 'POINT(4 3)')))
FROM (SELECT ST_GeomFromText('LINESTRING(1 2, 4 5, 6 7)') As geom) AS line;
------------
 POINT(3 4)
```


## See Also


 [ST_LineInterpolatePoints](#ST_LineInterpolatePoints), [ST_3DLineInterpolatePoint](#ST_3DLineInterpolatePoint), [ST_LineLocatePoint](#ST_LineLocatePoint)
  <a id="ST_3DLineInterpolatePoint"></a>

# ST_3DLineInterpolatePoint

Returns a point interpolated along a 3D line at a fractional location.

## Synopsis


```sql
geometry ST_3DLineInterpolatePoint(geometry  a_linestring, float8  a_fraction)
```


## Description


Returns a point interpolated along a 3D line at a fractional location. First argument must be a LINESTRING. Second argument is a float between 0 and 1 representing the point location as a fraction of line length. The M value is interpolated if present.


!!! note

    [ST_LineInterpolatePoint](#ST_LineInterpolatePoint) computes points in 2D and then interpolates the values for Z and M, while this function computes points in 3D and only interpolates the M value.


Availability: 3.0.0


## Examples


Return point 20% along 3D line


```sql

SELECT ST_AsText(
    ST_3DLineInterpolatePoint('LINESTRING(25 50 70, 100 125 90, 150 190 200)',
        0.20));

   st_asetext
----------------
 POINT Z (59.0675892910822 84.0675892910822 79.0846904776219)
```


## See Also


 [ST_LineInterpolatePoint](#ST_LineInterpolatePoint), [ST_LineInterpolatePoints](#ST_LineInterpolatePoints), [ST_LineLocatePoint](#ST_LineLocatePoint)
  <a id="ST_LineInterpolatePoints"></a>

# ST_LineInterpolatePoints

Returns points interpolated along a line at a fractional interval.

## Synopsis


```sql
geometry ST_LineInterpolatePoints(geometry  a_linestring, float8  a_fraction, boolean  repeat)
geography ST_LineInterpolatePoints(geography  a_linestring, float8  a_fraction, boolean use_spheroid = true, boolean  repeat = true)
```


## Description


Returns one or more points interpolated along a line at a fractional interval. The first argument must be a LINESTRING. The second argument is a float8 between 0 and 1 representing the spacing between the points as a fraction of line length. If the third argument is false, at most one point will be constructed (which is equivalent to [ST_LineInterpolatePoint](#ST_LineInterpolatePoint).)


 If the result has zero or one points, it is returned as a POINT. If it has two or more points, it is returned as a MULTIPOINT.


Availability: 2.5.0


## Examples


![image](images/st_line_interpolate_points01.png)


A LineString with points interpolated every 20%


```
--Return points each 20% along a 2D line
SELECT ST_AsText(ST_LineInterpolatePoints('LINESTRING(25 50, 100 125, 150 190)', 0.20))
----------------
 MULTIPOINT((51.5974135047432 76.5974135047432),(78.1948270094864 103.194827009486),(104.132163186446 130.37181214238),(127.066081593223 160.18590607119),(150 190))
```


## See Also


 [ST_LineInterpolatePoint](#ST_LineInterpolatePoint), [ST_LineLocatePoint](#ST_LineLocatePoint)
  <a id="ST_LineLocatePoint"></a>

# ST_LineLocatePoint

Returns the fractional location of the closest point on a line to a point.

## Synopsis


```sql
float8 ST_LineLocatePoint(geometry  a_linestring, geometry  a_point)
float8 ST_LineLocatePoint(geography  a_linestring, geography  a_point, boolean use_spheroid = true)
```


## Description


Returns a float between 0 and 1 representing the location of the closest point on a LineString to the given Point, as a fraction of [2d line](measurement-functions.md#ST_Length2D) length.


You can use the returned location to extract a Point ([ST_LineInterpolatePoint](#ST_LineInterpolatePoint)) or a substring ([ST_LineSubstring](#ST_LineSubstring)).


This is useful for approximating numbers of addresses


Availability: 1.1.0


Changed: 2.1.0. Up to 2.0.x this was called ST_Line_Locate_Point.


## Examples


```

--Rough approximation of finding the street number of a point along the street
--Note the whole foo thing is just to generate dummy data that looks
--like house centroids and street
--We use ST_DWithin to exclude
--houses too far away from the street to be considered on the street
SELECT ST_AsText(house_loc) As as_text_house_loc,
	startstreet_num +
		CAST( (endstreet_num - startstreet_num)
			* ST_LineLocatePoint(street_line, house_loc) As integer) As street_num
FROM
(SELECT ST_GeomFromText('LINESTRING(1 2, 3 4)') As street_line,
	ST_Point(x*1.01,y*1.03) As house_loc, 10 As startstreet_num,
		20 As endstreet_num
FROM generate_series(1,3) x CROSS JOIN generate_series(2,4) As y)
As foo
WHERE ST_DWithin(street_line, house_loc, 0.2);

 as_text_house_loc | street_num
-------------------+------------
 POINT(1.01 2.06)  |         10
 POINT(2.02 3.09)  |         15
 POINT(3.03 4.12)  |         20

 --find closest point on a line to a point or other geometry
 SELECT ST_AsText(ST_LineInterpolatePoint(foo.the_line, ST_LineLocatePoint(foo.the_line, ST_GeomFromText('POINT(4 3)'))))
FROM (SELECT ST_GeomFromText('LINESTRING(1 2, 4 5, 6 7)') As the_line) As foo;
   st_astext
----------------
 POINT(3 4)
```


## See Also


[ST_DWithin](spatial-relationships.md#ST_DWithin), [ST_Length2D](measurement-functions.md#ST_Length2D), [ST_LineInterpolatePoint](#ST_LineInterpolatePoint), [ST_LineSubstring](#ST_LineSubstring)
  <a id="ST_LineSubstring"></a>

# ST_LineSubstring

Returns the part of a line between two fractional locations.

## Synopsis


```sql
geometry ST_LineSubstring(geometry  a_linestring, float8  startfraction, float8  endfraction)
geography ST_LineSubstring(geography  a_linestring, float8  startfraction, float8  endfraction)
```


## Description


Computes the line which is the section of the input line starting and ending at the given fractional locations. The first argument must be a LINESTRING. The second and third arguments are values in the range [0, 1] representing the start and end locations as fractions of line length. The Z and M values are interpolated for added endpoints if present.


If `startfraction` and `endfraction` have the same value this is equivalent to [ST_LineInterpolatePoint](#ST_LineInterpolatePoint).


!!! note

    This only works with LINESTRINGs. To use on contiguous MULTILINESTRINGs first join them with [ST_LineMerge](geometry-processing.md#ST_LineMerge).


!!! note

    Since release 1.1.1 this function interpolates M and Z values. Prior releases set Z and M to unspecified values.


Enhanced: 3.4.0 - Support for geography was introduced.


Changed: 2.1.0. Up to 2.0.x this was called ST_Line_Substring.


Availability: 1.1.0, Z and M supported added in 1.1.1


## Examples


![image](images/st_line_substring01.png)


A LineString seen with 1/3 midrange overlaid (0.333, 0.666)


```sql

SELECT ST_AsText(ST_LineSubstring( 'LINESTRING (20 180, 50 20, 90 80, 120 40, 180 150)', 0.333, 0.666));
------------------------------------------------------------------------------------------------
LINESTRING (45.17311810399485 45.74337011202746, 50 20, 90 80, 112.97593050157862 49.36542599789519)
```


 If start and end locations are the same, the result is a POINT.


```sql

SELECT ST_AsText(ST_LineSubstring( 'LINESTRING(25 50, 100 125, 150 190)', 0.333, 0.333));
------------------------------------------
 POINT(69.2846934853974 94.2846934853974)
```


 A query to cut a LineString into sections of length 100 or shorter. It uses `generate_series()` with a CROSS JOIN LATERAL to produce the equivalent of a FOR loop.


```sql


WITH data(id, geom) AS (VALUES
        ( 'A', 'LINESTRING( 0 0, 200 0)'::geometry ),
        ( 'B', 'LINESTRING( 0 100, 350 100)'::geometry ),
        ( 'C', 'LINESTRING( 0 200, 50 200)'::geometry )
    )
SELECT id, i,
       ST_AsText( ST_LineSubstring( geom, startfrac, LEAST( endfrac, 1 )) ) AS geom
FROM (
    SELECT id, geom, ST_Length(geom) len, 100 sublen FROM data
    ) AS d
CROSS JOIN LATERAL (
    SELECT i, (sublen * i) / len AS startfrac,
              (sublen * (i+1)) / len AS endfrac
    FROM generate_series(0, floor( len / sublen )::integer ) AS t(i)
    -- skip last i if line length is exact multiple of sublen
    WHERE (sublen * i) / len <> 1.0
    ) AS d2;

 id | i |            geom
----+---+-----------------------------
 A  | 0 | LINESTRING(0 0,100 0)
 A  | 1 | LINESTRING(100 0,200 0)
 B  | 0 | LINESTRING(0 100,100 100)
 B  | 1 | LINESTRING(100 100,200 100)
 B  | 2 | LINESTRING(200 100,300 100)
 B  | 3 | LINESTRING(300 100,350 100)
 C  | 0 | LINESTRING(0 200,50 200)
```


Geography implementation measures along a spheroid, geometry along a line


```sql

SELECT ST_AsText(ST_LineSubstring( 'LINESTRING(-118.2436 34.0522, -71.0570 42.3611)'::geography, 0.333, 0.666),6) AS geog_sub
 , ST_AsText(ST_LineSubstring('LINESTRING(-118.2436 34.0522, -71.0570 42.3611)'::geometry, 0.333, 0.666),6) AS geom_sub;
---------------------------------------------------------------
geog_sub | LINESTRING(-104.167064 38.854691,-87.674646 41.849854)
geom_sub | LINESTRING(-102.530462 36.819064,-86.817324 39.585927)
```


## See Also


[ST_Length](measurement-functions.md#ST_Length), [ST_LineExtend](geometry-editors.md#ST_LineExtend), [ST_LineInterpolatePoint](#ST_LineInterpolatePoint), [ST_LineMerge](geometry-processing.md#ST_LineMerge)
  <a id="ST_LocateAlong"></a>

# ST_LocateAlong

Returns the point(s) on a geometry that match a measure value.

## Synopsis


```sql
geometry ST_LocateAlong(geometry  geom_with_measure, float8  measure, float8  offset = 0)
```


## Description


Returns the location(s) along a measured geometry that have the given measure values. The result is a Point or MultiPoint. Polygonal inputs are not supported.


If `offset` is provided, the result is offset to the left or right of the input line by the specified distance. A positive offset will be to the left, and a negative one to the right.


!!! note

    Use this function only for linear geometries with an M component


The semantic is specified by the *ISO/IEC 13249-3 SQL/MM Spatial* standard.


Availability: 1.1.0 by old name ST_Locate_Along_Measure.


Changed: 2.0.0 in prior versions this used to be called ST_Locate_Along_Measure.


 SQL-MM IEC 13249-3: 5.1.13


## Examples


```sql

SELECT ST_AsText(
  ST_LocateAlong(
    'MULTILINESTRINGM((1 2 3, 3 4 2, 9 4 3),(1 2 3, 5 4 5))'::geometry,
    3 ));

----------------------------------
 MULTIPOINT M ((1 2 3),(9 4 3),(1 2 3))
```


## See Also


[ST_LocateBetween](#ST_LocateBetween), [ST_LocateBetweenElevations](#ST_LocateBetweenElevations), [ST_InterpolatePoint](#ST_InterpolatePoint)
  <a id="ST_LocateBetween"></a>

# ST_LocateBetween

Returns the portions of a geometry that match a measure range.

## Synopsis


```sql
geometry ST_LocateBetween(geometry  geom, float8  measure_start, float8  measure_end, float8  offset = 0)
```


## Description


Return a geometry (collection) with the portions of the input measured geometry that match the specified measure range (inclusively).


If the `offset` is provided, the result is offset to the left or right of the input line by the specified distance. A positive offset will be to the left, and a negative one to the right.


Clipping a non-convex POLYGON may produce invalid geometry.


The semantic is specified by the *ISO/IEC 13249-3 SQL/MM Spatial* standard.


Availability: 1.1.0 by old name ST_Locate_Between_Measures.


Changed: 2.0.0 - in prior versions this used to be called ST_Locate_Between_Measures.


Enhanced: 3.0.0 - added support for POLYGON, TIN, TRIANGLE.


 SQL-MM IEC 13249-3: 5.1


## Examples


```sql

SELECT ST_AsText(
  ST_LocateBetween(
       'MULTILINESTRING M ((1 2 3, 3 4 2, 9 4 3),(1 2 3, 5 4 5))':: geometry,
       1.5, 3 ));
------------------------------------------------------------------------
 GEOMETRYCOLLECTION M (LINESTRING M (1 2 3,3 4 2,9 4 3),POINT M (1 2 3))
```


![image](images/st_locatebetween01.png)


A LineString with the section between measures 2 and 8, offset to the left


```sql

SELECT ST_AsText( ST_LocateBetween(
  ST_AddMeasure('LINESTRING (20 180, 50 20, 100 120, 180 20)', 0, 10),
  2, 8,
  20
));
------------------------------------------------------------------------
MULTILINESTRING((54.49835019899045 104.53426957938231,58.70056060327303 82.12248075654186,69.16695286779743 103.05526528559065,82.11145618000168 128.94427190999915,84.24893681714357 132.32493442618113,87.01636951231555 135.21267035596549,90.30307285299679 137.49198684843182,93.97759758337769 139.07172433557758,97.89298381958797 139.8887023914453,101.89263860095893 139.9102465862721,105.81659870902816 139.13549527600819,109.50792827749828 137.5954340631298,112.81899532549731 135.351656550512,115.6173761888606 132.49390095108848,145.31017306064817 95.37790486135405))
```


## See Also


[ST_LocateAlong](#ST_LocateAlong), [ST_LocateBetweenElevations](#ST_LocateBetweenElevations)
  <a id="ST_LocateBetweenElevations"></a>

# ST_LocateBetweenElevations

Returns the portions of a geometry that lie in an elevation (Z) range.

## Synopsis


```sql
geometry ST_LocateBetweenElevations(geometry  geom, float8  elevation_start, float8  elevation_end)
```


## Description


Returns a geometry (collection) with the portions of a geometry that lie in an elevation (Z) range.


Clipping a non-convex POLYGON may produce invalid geometry.


Availability: 1.4.0


Enhanced: 3.0.0 - added support for POLYGON, TIN, TRIANGLE.


## Examples


```sql
SELECT ST_AsText(
  ST_LocateBetweenElevations(
    'LINESTRING(1 2 3, 4 5 6)'::geometry,
    2, 4 ));

             st_astext
-----------------------------------
 MULTILINESTRING Z ((1 2 3,2 3 4))

SELECT ST_AsText(
    ST_LocateBetweenElevations(
      'LINESTRING(1 2 6, 4 5 -1, 7 8 9)',
      6, 9)) As ewelev;

                                ewelev
-----------------------------------------------------------------------
 GEOMETRYCOLLECTION Z (POINT Z (1 2 6),LINESTRING Z (6.1 7.1 6,7 8 9))
```


## See Also


[ST_Dump](geometry-accessors.md#ST_Dump), [ST_LocateBetween](#ST_LocateBetween)
  <a id="ST_InterpolatePoint"></a>

# ST_InterpolatePoint

Returns the interpolated measure of a geometry closest to a point.

## Synopsis


```sql
float8 ST_InterpolatePoint(geometry  linear_geom_with_measure, geometry  point)
```


## Description


Returns an interpolated measure value of a linear measured geometry at the location closest to the given point.


!!! note

    Use this function only for linear geometries with an M component


Availability: 2.0.0


## Examples


```sql
SELECT ST_InterpolatePoint('LINESTRING M (0 0 0, 10 0 20)', 'POINT(5 5)');
 ---------------------
         10

```


## See Also


[ST_AddMeasure](#ST_AddMeasure), [ST_LocateAlong](#ST_LocateAlong), [ST_LocateBetween](#ST_LocateBetween)
  <a id="ST_AddMeasure"></a>

# ST_AddMeasure

Interpolates measures along a linear geometry.

## Synopsis


```sql
geometry ST_AddMeasure(geometry  geom_mline, float8  measure_start, float8  measure_end)
```


## Description


Return a derived geometry with measure values linearly interpolated between the start and end points. If the geometry has no measure dimension, one is added. If the geometry has a measure dimension, it is over-written with new values. Only LINESTRINGS and MULTILINESTRINGS are supported.


Availability: 1.5.0


## Examples


```sql
SELECT ST_AsText(ST_AddMeasure(
ST_GeomFromEWKT('LINESTRING(1 0, 2 0, 4 0)'),1,4)) As ewelev;
           ewelev
--------------------------------
 LINESTRINGM(1 0 1,2 0 2,4 0 4)

SELECT ST_AsText(ST_AddMeasure(
ST_GeomFromEWKT('LINESTRING(1 0 4, 2 0 4, 4 0 4)'),10,40)) As ewelev;
                 ewelev
----------------------------------------
 LINESTRING(1 0 4 10,2 0 4 20,4 0 4 40)

SELECT ST_AsText(ST_AddMeasure(
ST_GeomFromEWKT('LINESTRINGM(1 0 4, 2 0 4, 4 0 4)'),10,40)) As ewelev;
                 ewelev
----------------------------------------
 LINESTRINGM(1 0 10,2 0 20,4 0 40)

SELECT ST_AsText(ST_AddMeasure(
ST_GeomFromEWKT('MULTILINESTRINGM((1 0 4, 2 0 4, 4 0 4),(1 0 4, 2 0 4, 4 0 4))'),10,70)) As ewelev;
                             ewelev
-----------------------------------------------------------------
 MULTILINESTRINGM((1 0 10,2 0 20,4 0 40),(1 0 40,2 0 50,4 0 70))
```
