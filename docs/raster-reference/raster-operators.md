<a id="RT_Operators"></a>

## Raster Operators
  <a id="RT_Raster_Intersect"></a>

# &&

Returns `TRUE` if A's bounding box intersects B's bounding box.

## Synopsis


```sql
boolean &&(raster
                        A, raster
                        B)
boolean &&(raster
                        A, geometry
                        B)
boolean &&(geometry
                        B, raster
                        A)
```


## Description


The `&&` operator returns `TRUE` if the bounding box of raster/geometr A intersects the bounding box of raster/geometr B.


!!! note

    This operand will make use of any indexes that may be available on the rasters.


Availability: 2.0.0


## Examples


```sql
SELECT A.rid As a_rid, B.rid As b_rid, A.rast && B.rast As intersect
 FROM dummy_rast AS A CROSS JOIN dummy_rast AS B LIMIT 3;

 a_rid | b_rid | intersect
-------+-------+---------
     2 |     2 | t
     2 |     3 | f
     2 |     1 | f
```
  <a id="RT_Raster_OverLeft"></a>

# &<

Returns `TRUE` if A's bounding box is to the left of B's.

## Synopsis


```sql
boolean &<(raster

                  A, raster

                  B)
```


## Description


The `&<` operator returns `TRUE` if the bounding box of raster A overlaps or is to the left of the bounding box of raster B, or more accurately, overlaps or is NOT to the right of the bounding box of raster B.


!!! note

    This operand will make use of any indexes that may be available on the rasters.


## Examples


```sql

SELECT A.rid As a_rid, B.rid As b_rid, A.rast &< B.rast As overleft
 FROM dummy_rast AS A CROSS JOIN dummy_rast AS B;

a_rid | b_rid | overleft
------+-------+----------
    2 |     2 | t
    2 |     3 | f
    2 |     1 | f
    3 |     2 | t
    3 |     3 | t
    3 |     1 | f
    1 |     2 | t
    1 |     3 | t
    1 |     1 | t
```
  <a id="RT_Raster_OverRight"></a>

# &>

Returns `TRUE` if A's bounding box is to the right of B's.

## Synopsis


```sql
boolean &>(raster

                  A, raster

                  B)
```


## Description


The `&>` operator returns `TRUE` if the bounding box of raster A overlaps or is to the right of the bounding box of raster B, or more accurately, overlaps or is NOT to the left of the bounding box of raster B.


!!! note

    This operand will make use of any indexes that may be available on the geometries.


## Examples


```sql

SELECT A.rid As a_rid, B.rid As b_rid, A.rast &> B.rast As overright
 FROM dummy_rast AS A CROSS JOIN dummy_rast AS B;

 a_rid | b_rid | overright
-------+-------+----------
     2 |     2 | t
     2 |     3 | t
     2 |     1 | t
     3 |     2 | f
     3 |     3 | t
     3 |     1 | f
     1 |     2 | f
     1 |     3 | t
     1 |     1 | t
```
  <a id="RT_Raster_EQ"></a>

# =

Returns `TRUE` if A's bounding box is the same as B's. Uses double precision bounding box.

## Synopsis


```sql
boolean =(raster

                  A, raster

                  B)
```


## Description


The `=` operator returns `TRUE` if the bounding box of raster A is the same as the bounding box of raster B. PostgreSQL uses the =, <, and > operators defined for rasters to perform internal orderings and comparison of rasters (ie. in a GROUP BY or ORDER BY clause).


!!! caution

    This operand will NOT make use of any indexes that may be available on the rasters. Use [RT_Raster_Same](#RT_Raster_Same) instead. This operator exists mostly so one can group by the raster column.


Availability: 2.1.0


## See Also


[RT_Raster_Same](#RT_Raster_Same)
  <a id="RT_Raster_Contained"></a>

# @

Returns `TRUE` if A's bounding box is contained by B's. Uses double precision bounding box.

## Synopsis


```sql
boolean @(raster
              A, raster
              B)
boolean @(geometry
              A, raster
              B)
boolean @(raster
              B, geometry
              A)
```


## Description


The `@` operator returns `TRUE` if the bounding box of raster/geometry A is contained by bounding box of raster/geometr B.


!!! note

    This operand will use spatial indexes on the rasters.


Availability: 2.0.0 raster @ raster, raster @ geometry introduced


Availability: 2.0.5 geometry @ raster introduced


## See Also


[RT_Raster_Contains](#RT_Raster_Contains)
  <a id="RT_Raster_Same"></a>

# ~=

Returns `TRUE` if A's bounding box is the same as B's.

## Synopsis


```sql
boolean ~=(raster

                  A, raster

                  B)
```


## Description


The `~=` operator returns `TRUE` if the bounding box of raster A is the same as the bounding box of raster B.


!!! note

    This operand will make use of any indexes that may be available on the rasters.


Availability: 2.0.0


## Examples


Very useful usecase is for taking two sets of single band rasters that are of the same chunk but represent different themes and creating a multi-band raster


```sql
SELECT ST_AddBand(prec.rast, alt.rast) As new_rast
    FROM prec INNER JOIN alt ON (prec.rast ~= alt.rast);

```


## See Also


[RT_ST_AddBand](raster-constructors.md#RT_ST_AddBand), [RT_Raster_EQ](#RT_Raster_EQ)
  <a id="RT_Raster_Contains"></a>

# ~

Returns `TRUE` if A's bounding box is contains B's. Uses double precision bounding box.

## Synopsis


```sql
boolean ~(raster
              A, raster
              B)
boolean ~(geometry
              A, raster
              B)
boolean ~(raster
              B, geometry
              A)
```


## Description


The `~` operator returns `TRUE` if the bounding box of raster/geometry A is contains bounding box of raster/geometr B.


!!! note

    This operand will use spatial indexes on the rasters.


Availability: 2.0.0


## See Also


[RT_Raster_Contained](#RT_Raster_Contained)
