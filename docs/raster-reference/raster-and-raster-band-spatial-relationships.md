<a id="Raster_Relationships"></a>

## Raster and Raster Band Spatial Relationships
  <a id="RT_ST_Contains"></a>

# ST_Contains

Return true if no points of raster rastB lie in the exterior of raster rastA and at least one point of the interior of rastB lies in the interior of rastA.

## Synopsis


```sql
boolean ST_Contains(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Contains(raster
                        rastA, raster
                        rastB)
```


## Description


 Raster rastA contains rastB if and only if no points of rastB lie in the exterior of rastA and at least one point of the interior of rastB lies in the interior of rastA. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_Contains(ST_Polygon(raster), geometry) or ST_Contains(geometry, ST_Polygon(raster)).


!!! note

    ST_Contains() is the inverse of ST_Within(). So, ST_Contains(rastA, rastB) implies ST_Within(rastB, rastA).


Availability: 2.1.0


## Examples


```

-- specified band numbers
SELECT r1.rid, r2.rid, ST_Contains(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 1;

NOTICE:  The first raster provided has no bands
 rid | rid | st_contains
-----+-----+-------------
   1 |   1 |
   1 |   2 | f

```


```

-- no band numbers specified
SELECT r1.rid, r2.rid, ST_Contains(r1.rast, r2.rast) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 1;
 rid | rid | st_contains
-----+-----+-------------
   1 |   1 | t
   1 |   2 | f

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects), [RT_ST_Within](#RT_ST_Within)
  <a id="RT_ST_ContainsProperly"></a>

# ST_ContainsProperly

Return true if rastB intersects the interior of rastA but not the boundary or exterior of rastA.

## Synopsis


```sql
boolean ST_ContainsProperly(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_ContainsProperly(raster
                        rastA, raster
                        rastB)
```


## Description


 Raster rastA contains properly rastB if rastB intersects the interior of rastA but not the boundary or exterior of rastA. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


 Raster rastA does not contain properly itself but does contain itself.


!!! note

    This function will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_ContainsProperly(ST_Polygon(raster), geometry) or ST_ContainsProperly(geometry, ST_Polygon(raster)).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_ContainsProperly(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_containsproperly
-----+-----+---------------------
   2 |   1 | f
   2 |   2 | f

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects), [RT_ST_Contains](#RT_ST_Contains)
  <a id="RT_ST_Covers"></a>

# ST_Covers

Return true if no points of raster rastB lie outside raster rastA.

## Synopsis


```sql
boolean ST_Covers(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Covers(raster
                        rastA, raster
                        rastB)
```


## Description


 Raster rastA covers rastB if and only if no points of rastB lie in the exterior of rastA. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_Covers(ST_Polygon(raster), geometry) or ST_Covers(geometry, ST_Polygon(raster)).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_Covers(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_covers
-----+-----+-----------
   2 |   1 | f
   2 |   2 | t

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects), [RT_ST_CoveredBy](#RT_ST_CoveredBy)
  <a id="RT_ST_CoveredBy"></a>

# ST_CoveredBy

Return true if no points of raster rastA lie outside raster rastB.

## Synopsis


```sql
boolean ST_CoveredBy(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_CoveredBy(raster
                        rastA, raster
                        rastB)
```


## Description


 Raster rastA is covered by rastB if and only if no points of rastA lie in the exterior of rastB. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_CoveredBy(ST_Polygon(raster), geometry) or ST_CoveredBy(geometry, ST_Polygon(raster)).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_CoveredBy(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_coveredby
-----+-----+--------------
   2 |   1 | f
   2 |   2 | t

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects), [RT_ST_Covers](#RT_ST_Covers)
  <a id="RT_ST_Disjoint"></a>

# ST_Disjoint

Return true if raster rastA does not spatially intersect rastB.

## Synopsis


```sql
boolean ST_Disjoint(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Disjoint(raster
                        rastA, raster
                        rastB)
```


## Description


 Raster rastA and rastB are disjointed if they do not share any space together. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function does NOT use any indexes.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_Disjoint(ST_Polygon(raster), geometry).


Availability: 2.1.0


## Examples


```

-- rid = 1 has no bands, hence the NOTICE and the NULL value for st_disjoint
SELECT r1.rid, r2.rid, ST_Disjoint(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

NOTICE:  The second raster provided has no bands
 rid | rid | st_disjoint
-----+-----+-------------
   2 |   1 |
   2 |   2 | f

```


```

-- this time, without specifying band numbers
SELECT r1.rid, r2.rid, ST_Disjoint(r1.rast, r2.rast) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_disjoint
-----+-----+-------------
   2 |   1 | t
   2 |   2 | f

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects)
  <a id="RT_ST_Intersects"></a>

# ST_Intersects

Return true if raster rastA spatially intersects raster rastB.

## Synopsis


```sql
boolean ST_Intersects(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Intersects(raster
                        rastA, raster
                        rastB)
boolean ST_Intersects(raster
                        rast, integer
                        nband, geometry
                        geommin)
boolean ST_Intersects(raster
                        rast, geometry
                        geommin, integer
                        nband=NULL)
boolean ST_Intersects(geometry
                        geommin, raster
                        rast, integer
                        nband=NULL)
```


## Description


 Return true if raster rastA spatially intersects raster rastB. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function will make use of any indexes that may be available on the rasters.


 Enhanced: 2.0.0 support raster/raster intersects was introduced.


!!! warning

    Changed: 2.1.0 The behavior of the ST_Intersects(raster, geometry) variants changed to match that of ST_Intersects(geometry, raster).


## Examples


```

-- different bands of same raster
SELECT ST_Intersects(rast, 2, rast, 3) FROM dummy_rast WHERE rid = 2;

 st_intersects
---------------
 t

```


## See Also


 [RT_ST_Intersection](raster-processing-map-algebra.md#RT_ST_Intersection), [RT_ST_Disjoint](#RT_ST_Disjoint)
  <a id="RT_ST_Overlaps"></a>

# ST_Overlaps

Return true if raster rastA and rastB intersect but one does not completely contain the other.

## Synopsis


```sql
boolean ST_Overlaps(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Overlaps(raster
                        rastA, raster
                        rastB)
```


## Description


 Return true if raster rastA spatially overlaps raster rastB. This means that rastA and rastB intersect but one does not completely contain the other. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_Overlaps(ST_Polygon(raster), geometry).


Availability: 2.1.0


## Examples


```

-- comparing different bands of same raster
SELECT ST_Overlaps(rast, 1, rast, 2) FROM dummy_rast WHERE rid = 2;

 st_overlaps
-------------
 f

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects)
  <a id="RT_ST_Touches"></a>

# ST_Touches

Return true if raster rastA and rastB have at least one point in common but their interiors do not intersect.

## Synopsis


```sql
boolean ST_Touches(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Touches(raster
                        rastA, raster
                        rastB)
```


## Description


 Return true if raster rastA spatially touches raster rastB. This means that rastA and rastB have at least one point in common but their interiors do not intersect. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This function will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_Touches(ST_Polygon(raster), geometry).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_Touches(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_touches
-----+-----+------------
   2 |   1 | f
   2 |   2 | f

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects)
  <a id="RT_ST_SameAlignment"></a>

# ST_SameAlignment

Returns true if rasters have same skew, scale, spatial ref, and offset (pixels can be put on same grid without cutting into pixels) and false if they don't with notice detailing issue.

## Synopsis


```sql
boolean ST_SameAlignment(raster
                  rastA, raster
                  rastB)
boolean ST_SameAlignment(double precision
                  ulx1, double precision
                  uly1, double precision
                  scalex1, double precision
                  scaley1, double precision
                  skewx1, double precision
                  skewy1, double precision
                  ulx2, double precision
                  uly2, double precision
                  scalex2, double precision
                  scaley2, double precision
                  skewx2, double precision
                  skewy2)
boolean ST_SameAlignment(raster set
                        rastfield)
```


## Description


 Non-Aggregate version (Variants 1 and 2): Returns true if the two rasters (either provided directly or made using the values for upperleft, scale, skew and srid) have the same scale, skew, srid and at least one of any of the four corners of any pixel of one raster falls on any corner of the grid of the other raster. Returns false if they don't and a NOTICE detailing the alignment issue.


 Aggregate version (Variant 3): From a set of rasters, returns true if all rasters in the set are aligned. The ST_SameAlignment() function is an "aggregate" function in the terminology of PostgreSQL. That means that it operates on rows of data, in the same way the SUM() and AVG() functions do.


Availability: 2.0.0


Enhanced: 2.1.0 addition of Aggegrate variant


## Examples: Rasters


```sql
SELECT ST_SameAlignment(
    ST_MakeEmptyRaster(1, 1, 0, 0, 1, 1, 0, 0),
    ST_MakeEmptyRaster(1, 1, 0, 0, 1, 1, 0, 0)
) as sm;

sm
----
t
```


```sql

SELECT ST_SameAlignment(A.rast,b.rast)
 FROM dummy_rast AS A CROSS JOIN dummy_rast AS B;

 NOTICE:  The two rasters provided have different SRIDs
NOTICE:  The two rasters provided have different SRIDs
 st_samealignment
------------------
 t
 f
 f
 f
```


## See Also


 [Loading and Creating Rasters](../raster-data-management-queries-and-applications/loading-and-creating-rasters.md#RT_Loading_Rasters), [RT_ST_NotSameAlignmentReason](#RT_ST_NotSameAlignmentReason), [RT_ST_MakeEmptyRaster](raster-constructors.md#RT_ST_MakeEmptyRaster)
  <a id="RT_ST_NotSameAlignmentReason"></a>

# ST_NotSameAlignmentReason

Returns text stating if rasters are aligned and if not aligned, a reason why.

## Synopsis


```sql
text ST_NotSameAlignmentReason(raster rastA, raster rastB)
```


## Description


Returns text stating if rasters are aligned and if not aligned, a reason why.


!!! note

    If there are several reasons why the rasters are not aligned, only one reason (the first test to fail) will be returned.


Availability: 2.1.0


## Examples


```sql

SELECT
    ST_SameAlignment(
        ST_MakeEmptyRaster(1, 1, 0, 0, 1, 1, 0, 0),
        ST_MakeEmptyRaster(1, 1, 0, 0, 1.1, 1.1, 0, 0)
    ),
    ST_NotSameAlignmentReason(
        ST_MakeEmptyRaster(1, 1, 0, 0, 1, 1, 0, 0),
        ST_MakeEmptyRaster(1, 1, 0, 0, 1.1, 1.1, 0, 0)
    )
;

 st_samealignment |            st_notsamealignmentreason
------------------+-------------------------------------------------
 f                | The rasters have different scales on the X axis
(1 row)

```


## See Also


 [Loading and Creating Rasters](../raster-data-management-queries-and-applications/loading-and-creating-rasters.md#RT_Loading_Rasters), [RT_ST_SameAlignment](#RT_ST_SameAlignment)
  <a id="RT_ST_Within"></a>

# ST_Within

Return true if no points of raster rastA lie in the exterior of raster rastB and at least one point of the interior of rastA lies in the interior of rastB.

## Synopsis


```sql
boolean ST_Within(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB)
boolean ST_Within(raster
                        rastA, raster
                        rastB)
```


## Description


 Raster rastA is within rastB if and only if no points of rastA lie in the exterior of rastB and at least one point of the interior of rastA lies in the interior of rastB. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


!!! note

    This operand will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_Within(ST_Polygon(raster), geometry) or ST_Within(geometry, ST_Polygon(raster)).


!!! note

    ST_Within() is the inverse of ST_Contains(). So, ST_Within(rastA, rastB) implies ST_Contains(rastB, rastA).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_Within(r1.rast, 1, r2.rast, 1) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_within
-----+-----+-----------
   2 |   1 | f
   2 |   2 | t

```


## See Also


 [RT_ST_Intersects](#RT_ST_Intersects), [RT_ST_Contains](#RT_ST_Contains), [RT_ST_DWithin](#RT_ST_DWithin), [RT_ST_DFullyWithin](#RT_ST_DFullyWithin)
  <a id="RT_ST_DWithin"></a>

# ST_DWithin

Return true if rasters rastA and rastB are within the specified distance of each other.

## Synopsis


```sql
boolean ST_DWithin(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB, double precision
                        distance_of_srid)
boolean ST_DWithin(raster
                        rastA, raster
                        rastB, double precision
                        distance_of_srid)
```


## Description


 Return true if rasters rastA and rastB are within the specified distance of each other. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


 The distance is specified in units defined by the spatial reference system of the rasters. For this function to make sense, the source rasters must both be of the same coordinate projection, having the same SRID.


!!! note

    This operand will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_DWithin(ST_Polygon(raster), geometry).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_DWithin(r1.rast, 1, r2.rast, 1, 3.14) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_dwithin
-----+-----+------------
   2 |   1 | f
   2 |   2 | t

```


## See Also


 [RT_ST_Within](#RT_ST_Within), [RT_ST_DFullyWithin](#RT_ST_DFullyWithin)
  <a id="RT_ST_DFullyWithin"></a>

# ST_DFullyWithin

Return true if rasters rastA and rastB are fully within the specified distance of each other.

## Synopsis


```sql
boolean ST_DFullyWithin(raster
                        rastA, integer
                        nbandA, raster
                        rastB, integer
                        nbandB, double precision
                        distance_of_srid)
boolean ST_DFullyWithin(raster
                        rastA, raster
                        rastB, double precision
                        distance_of_srid)
```


## Description


 Return true if rasters rastA and rastB are fully within the specified distance of each other. If the band number is not provided (or set to NULL), only the convex hull of the raster is considered in the test. If the band number is provided, only those pixels with value (not NODATA) are considered in the test.


 The distance is specified in units defined by the spatial reference system of the rasters. For this function to make sense, the source rasters must both be of the same coordinate projection, having the same SRID.


!!! note

    This operand will make use of any indexes that may be available on the rasters.


!!! note

    To test the spatial relationship of a raster and a geometry, use ST_Polygon on the raster, e.g. ST_DFullyWithin(ST_Polygon(raster), geometry).


Availability: 2.1.0


## Examples


```sql

SELECT r1.rid, r2.rid, ST_DFullyWithin(r1.rast, 1, r2.rast, 1, 3.14) FROM dummy_rast r1 CROSS JOIN dummy_rast r2 WHERE r1.rid = 2;

 rid | rid | st_dfullywithin
-----+-----+-----------------
   2 |   1 | f
   2 |   2 | t

```


## See Also


 [RT_ST_Within](#RT_ST_Within), [RT_ST_DWithin](#RT_ST_DWithin)
