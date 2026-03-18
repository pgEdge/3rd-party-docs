<a id="OGC_Validity"></a>

## Geometry Validation


PostGIS is compliant with the Open Geospatial Consortium’s (OGC) Simple Features specification. That standard defines the concepts of geometry being *simple* and *valid*. These definitions allow the Simple Features geometry model to represent spatial objects in a consistent and unambiguous way that supports efficient computation. (Note: the OGC SF and SQL/MM have the same definitions for simple and valid.)
 <a id="Simple_Geometry"></a>

## Simple Geometry


A *simple* geometry is one that has no anomalous geometric points, such as self intersection or self tangency.


A `POINT` is inherently *simple* as a 0-dimensional geometry object.


`MULTIPOINT`s are *simple* if no two coordinates (`POINT`s) are equal (have identical coordinate values).


A `LINESTRING` is *simple* if it does not pass through the same point twice, except for the endpoints. If the endpoints of a simple LineString are identical it is called *closed* and referred to as a Linear Ring.


| * **(a)** and **(c)** are simple `LINESTRING`s. **(b)** and **(d)** are not simple. **(c)** is a closed Linear Ring. * |


A `MULTILINESTRING` is *simple* only if all of its elements are simple and the only intersection between any two elements occurs at points that are on the boundaries of both elements.


| * **(e)** and **(f)** are simple `MULTILINESTRING`s. **(g)** is not simple. * |


`POLYGON`s are formed from linear rings, so valid polygonal geometry is always *simple*.


To test if a geometry is simple use the [ST_IsSimple](../postgis-reference/geometry-accessors.md#ST_IsSimple) function:


```sql

SELECT
   ST_IsSimple('LINESTRING(0 0, 100 100)') AS straight,
   ST_IsSimple('LINESTRING(0 0, 100 100, 100 0, 0 100)') AS crossing;

 straight | crossing
----------+----------
 t        | f
```


Generally, PostGIS functions do not require geometric arguments to be simple. Simplicity is primarily used as a basis for defining geometric validity. It is also a requirement for some kinds of spatial data models (for example, linear networks often disallow lines that cross). Multipoint and linear geometry can be made simple using [ST_UnaryUnion](../postgis-reference/overlay-functions.md#ST_UnaryUnion).
  <a id="Valid_Geometry"></a>

## Valid Geometry


Geometry validity primarily applies to 2-dimensional geometries (`POLYGON`s and `MULTIPOLYGON`s) . Validity is defined by rules that allow polygonal geometry to model planar areas unambiguously.


A `POLYGON` is *valid* if:


1.  the polygon boundary rings (the exterior shell ring and interior hole rings) are *simple* (do not cross or self-touch). Because of this a polygon cannot have cut lines, spikes or loops. This implies that polygon holes must be represented as interior rings, rather than by the exterior ring self-touching (a so-called "inverted hole").
2.  boundary rings do not cross
3.  boundary rings may touch at points but only as a tangent (i.e. not in a line)
4.  interior rings are contained in the exterior ring
5.  the polygon interior is simply connected (i.e. the rings must not touch in a way that splits the polygon into more than one part)


| * **(h)** and **(i)** are valid `POLYGON`s. **(j-m)** are invalid. **(j)** can be represented as a valid `MULTIPOLYGON`. * |


A `MULTIPOLYGON` is *valid* if:


1.  its element `POLYGON`s are valid
2.  elements do not overlap (i.e. their interiors must not intersect)
3.  elements touch only at points (i.e. not along a line)


| * **(n)** is a valid `MULTIPOLYGON`. **(o)** and **(p)** are invalid. * |


These rules mean that valid polygonal geometry is also *simple*.


For linear geometry the only validity rule is that `LINESTRING`s must have at least two points and have non-zero length (or equivalently, have at least two distinct points.) Note that non-simple (self-intersecting) lines are valid.


```sql

SELECT
   ST_IsValid('LINESTRING(0 0, 1 1)') AS len_nonzero,
   ST_IsValid('LINESTRING(0 0, 0 0, 0 0)') AS len_zero,
   ST_IsValid('LINESTRING(10 10, 150 150, 180 50, 20 130)') AS self_int;

 len_nonzero | len_zero | self_int
-------------+----------+----------
 t           | f        | t
```


`POINT` and `MULTIPOINT` geometries have no validity rules.
  <a id="Managing_Validity"></a>

## Managing Validity


PostGIS allows creating and storing both valid and invalid Geometry. This allows invalid geometry to be detected and flagged or fixed. There are also situations where the OGC validity rules are stricter than desired (examples of this are zero-length linestrings and polygons with inverted holes.)


Many of the functions provided by PostGIS rely on the assumption that geometry arguments are valid. For example, it does not make sense to calculate the area of a polygon that has a hole defined outside of the polygon, or to construct a polygon from a non-simple boundary line. Assuming valid geometric inputs allows functions to operate more efficiently, since they do not need to check for topological correctness. (Notable exceptions are that zero-length lines and polygons with inversions are generally handled correctly.) Also, most PostGIS functions produce valid geometry output if the inputs are valid. This allows PostGIS functions to be chained together safely.


If you encounter unexpected error messages when calling PostGIS functions (such as "GEOS Intersection() threw an error!"), you should first confirm that the function arguments are valid. If they are not, then consider using one of the techniques below to ensure the data you are processing is valid.


!!! note

    If a function reports an error with valid inputs, then you may have found an error in either PostGIS or one of the libraries it uses, and you should report this to the PostGIS project. The same is true if a PostGIS function returns an invalid geometry for valid input.


To test if a geometry is valid use the [ST_IsValid](../postgis-reference/geometry-validation.md#ST_IsValid) function:


```sql

SELECT ST_IsValid('POLYGON ((20 180, 180 180, 180 20, 20 20, 20 180))');
-----------------
 t
```


Information about the nature and location of an geometry invalidity are provided by the [ST_IsValidDetail](../postgis-reference/geometry-validation.md#ST_IsValidDetail) function:


```sql

SELECT valid, reason, ST_AsText(location) AS location
    FROM ST_IsValidDetail('POLYGON ((20 20, 120 190, 50 190, 170 50, 20 20))') AS t;

 valid |      reason       |                  location
-------+-------------------+---------------------------------------------
 f     | Self-intersection | POINT(91.51162790697674 141.56976744186045)
```


In some situations it is desirable to correct invalid geometry automatically. Use the [ST_MakeValid](../postgis-reference/geometry-validation.md#ST_MakeValid) function to do this. (<code>ST_MakeValid</code> is a case of a spatial function that *does* allow invalid input!)


By default, PostGIS does not check for validity when loading geometry, because validity testing can take a lot of CPU time for complex geometries. If you do not trust your data sources, you can enforce a validity check on your tables by adding a check constraint:


```sql
ALTER TABLE mytable
  ADD CONSTRAINT geometry_valid_check
	CHECK (ST_IsValid(geom));
```
