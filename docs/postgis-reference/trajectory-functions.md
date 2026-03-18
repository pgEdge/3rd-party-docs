<a id="Temporal"></a>

## Trajectory Functions
  <a id="ST_IsValidTrajectory"></a>

# ST_IsValidTrajectory

Tests if the geometry is a valid trajectory.

## Synopsis


```sql
boolean ST_IsValidTrajectory(geometry  line)
```


## Description


 Tests if a geometry encodes a valid trajectory. A valid trajectory is represented as a `LINESTRING` with measures (M values). The measure values must increase from each vertex to the next.


 Valid trajectories are expected as input to spatio-temporal functions like [ST_ClosestPointOfApproach](#ST_ClosestPointOfApproach)


Availability: 2.2.0


## Examples


```

-- A valid trajectory
SELECT ST_IsValidTrajectory(ST_MakeLine(
  ST_MakePointM(0,0,1),
  ST_MakePointM(0,1,2))
);
 t

-- An invalid trajectory
SELECT ST_IsValidTrajectory(ST_MakeLine(ST_MakePointM(0,0,1), ST_MakePointM(0,1,0)));
NOTICE:  Measure of vertex 1 (0) not bigger than measure of vertex 0 (1)
 st_isvalidtrajectory
----------------------
 f
```


## See Also


 [ST_ClosestPointOfApproach](#ST_ClosestPointOfApproach)
  <a id="ST_ClosestPointOfApproach"></a>

# ST_ClosestPointOfApproach

Returns a measure at the closest point of approach of two trajectories.

## Synopsis


```sql
float8 ST_ClosestPointOfApproach(geometry  track1, geometry  track2)
```


## Description


 Returns the smallest measure at which points interpolated along the given trajectories are the least distance apart.


 Inputs must be valid trajectories as checked by [ST_IsValidTrajectory](#ST_IsValidTrajectory). Null is returned if the trajectories do not overlap in their M ranges.


 To obtain the actual points at the computed measure use [ST_LocateAlong](linear-referencing.md#ST_LocateAlong) .


Availability: 2.2.0


## Examples


```

-- Return the time in which two objects moving between 10:00 and 11:00
-- are closest to each other and their distance at that point
WITH inp AS ( SELECT
  ST_AddMeasure('LINESTRING Z (0 0 0, 10 0 5)'::geometry,
    extract(epoch from '2015-05-26 10:00'::timestamptz),
    extract(epoch from '2015-05-26 11:00'::timestamptz)
  ) a,
  ST_AddMeasure('LINESTRING Z (0 2 10, 12 1 2)'::geometry,
    extract(epoch from '2015-05-26 10:00'::timestamptz),
    extract(epoch from '2015-05-26 11:00'::timestamptz)
  ) b
), cpa AS (
  SELECT ST_ClosestPointOfApproach(a,b) m FROM inp
), points AS (
  SELECT ST_GeometryN(ST_LocateAlong(a,m),1) pa,
         ST_GeometryN(ST_LocateAlong(b,m),1) pb
  FROM inp, cpa
)
SELECT to_timestamp(m) t,
       ST_3DDistance(pa,pb) distance,
       ST_AsText(pa, 2) AS pa, ST_AsText(pb, 2) AS pb
FROM points, cpa;

               t               |      distance      |                  pa                  |                   pb
-------------------------------+--------------------+--------------------------------------+----------------------------------------
 2015-05-26 10:45:31.034483-07 | 1.9652147377620688 | POINT ZM (7.59 0 3.79 1432662331.03) | POINT ZM (9.1 1.24 3.93 1432662331.03)

```


## See Also


 [ST_IsValidTrajectory](#ST_IsValidTrajectory), [ST_DistanceCPA](#ST_DistanceCPA), [ST_LocateAlong](linear-referencing.md#ST_LocateAlong), [ST_AddMeasure](linear-referencing.md#ST_AddMeasure)
  <a id="ST_DistanceCPA"></a>

# ST_DistanceCPA

Returns the distance between the closest point of approach of two trajectories.

## Synopsis


```sql
float8 ST_DistanceCPA(geometry  track1, geometry  track2)
```


## Description


 Returns the distance (in 2D) between two trajectories at their closest point of approach.


 Inputs must be valid trajectories as checked by [ST_IsValidTrajectory](#ST_IsValidTrajectory). Null is returned if the trajectories do not overlap in their M ranges.


Availability: 2.2.0


## Examples


```

-- Return the minimum distance of two objects moving between 10:00 and 11:00
WITH inp AS ( SELECT
  ST_AddMeasure('LINESTRING Z (0 0 0, 10 0 5)'::geometry,
    extract(epoch from '2015-05-26 10:00'::timestamptz),
    extract(epoch from '2015-05-26 11:00'::timestamptz)
  ) a,
  ST_AddMeasure('LINESTRING Z (0 2 10, 12 1 2)'::geometry,
    extract(epoch from '2015-05-26 10:00'::timestamptz),
    extract(epoch from '2015-05-26 11:00'::timestamptz)
  ) b
)
SELECT ST_DistanceCPA(a,b) distance FROM inp;

     distance
-------------------
 1.965214737762069
```


## See Also


 [ST_IsValidTrajectory](#ST_IsValidTrajectory), [ST_ClosestPointOfApproach](#ST_ClosestPointOfApproach), [ST_AddMeasure](linear-referencing.md#ST_AddMeasure), [geometry_distance_cpa](operators.md#geometry_distance_cpa)
  <a id="ST_CPAWithin"></a>

# ST_CPAWithin

Tests if the closest point of approach of two trajectories is within the specified distance.

## Synopsis


```sql
boolean ST_CPAWithin(geometry  track1, geometry  track2, float8  dist)
```


## Description


 Tests whether two moving objects have ever been closer than the specified distance.


 Inputs must be valid trajectories as checked by [ST_IsValidTrajectory](#ST_IsValidTrajectory). False is returned if the trajectories do not overlap in their M ranges.


Availability: 2.2.0


## Examples


```sql

WITH inp AS ( SELECT
  ST_AddMeasure('LINESTRING Z (0 0 0, 10 0 5)'::geometry,
    extract(epoch from '2015-05-26 10:00'::timestamptz),
    extract(epoch from '2015-05-26 11:00'::timestamptz)
  ) a,
  ST_AddMeasure('LINESTRING Z (0 2 10, 12 1 2)'::geometry,
    extract(epoch from '2015-05-26 10:00'::timestamptz),
    extract(epoch from '2015-05-26 11:00'::timestamptz)
  ) b
)
SELECT ST_CPAWithin(a,b,2), ST_DistanceCPA(a,b) distance FROM inp;

 st_cpawithin |     distance
--------------+------------------
 t            | 1.96521473776207
```


## See Also


 [ST_IsValidTrajectory](#ST_IsValidTrajectory), [ST_ClosestPointOfApproach](#ST_ClosestPointOfApproach), [ST_DistanceCPA](#ST_DistanceCPA), [geometry_distance_cpa](operators.md#geometry_distance_cpa)
