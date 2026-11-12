<a id="datatype-geometric"></a>

## Geometric Types


 Geometric data types represent two-dimensional spatial objects. [Geometric Types](#datatype-geo-table) shows the geometric types available in PostgreSQL.
 <a id="datatype-geo-table"></a>

**Table: Geometric Types**

| Name | Storage Size | Description | Representation |
| --- | --- | --- | --- |
| `point` | 16 bytes | Point on a plane | (x,y) |
| `line` | 24 bytes | Infinite line | {A,B,C} |
| `lseg` | 32 bytes | Finite line segment | [(x1,y1),(x2,y2)] |
| `box` | 32 bytes | Rectangular box | (x1,y1),(x2,y2) |
| `path` | 16+16n bytes | Closed path (similar to polygon) | ((x1,y1),...) |
| `path` | 16+16n bytes | Open path | [(x1,y1),...] |
| `polygon` | 40+16n bytes | Polygon (similar to closed path) | ((x1,y1),...) |
| `circle` | 24 bytes | Circle | <(x,y),r> (center point and radius) |


 In all these types, the individual coordinates are stored as `double precision` (`float8`) numbers.


 A rich set of functions and operators is available to perform various geometric operations such as scaling, translation, rotation, and determining intersections. They are explained in [Geometric Functions and Operators](../functions-and-operators/geometric-functions-and-operators.md#functions-geometry).
 <a id="datatype-geometric-points"></a>

### Points


 Points are the fundamental two-dimensional building block for geometric types. Values of type `point` are specified using either of the following syntaxes:

```

( X , Y )
  X , Y
```
 where *x* and *y* are the respective coordinates, as floating-point numbers.


 Points are output using the first syntax.
  <a id="datatype-line"></a>

### Lines


 Lines are represented by the linear equation *A*x + *B*y + *C* = 0, where *A* and *B* are not both zero. Values of type `line` are input and output in the following form:

```

{ A, B, C }
```
 Alternatively, any of the following forms can be used for input:

```

[ ( X1 , Y1 ) , ( X2 , Y2 ) ]
( ( X1 , Y1 ) , ( X2 , Y2 ) )
  ( X1 , Y1 ) , ( X2 , Y2 )
    X1 , Y1   ,   X2 , Y2
```
 where <code>(</code><em>x1</em><code>,</code><em>y1</em><code>)</code> and <code>(</code><em>x2</em><code>,</code><em>y2</em><code>)</code> are two different points on the line.
  <a id="datatype-lseg"></a>

### Line Segments


 Line segments are represented by pairs of points that are the endpoints of the segment. Values of type `lseg` are specified using any of the following syntaxes:

```

[ ( X1 , Y1 ) , ( X2 , Y2 ) ]
( ( X1 , Y1 ) , ( X2 , Y2 ) )
  ( X1 , Y1 ) , ( X2 , Y2 )
    X1 , Y1   ,   X2 , Y2
```
 where <code>(</code><em>x1</em><code>,</code><em>y1</em><code>)</code> and <code>(</code><em>x2</em><code>,</code><em>y2</em><code>)</code> are the end points of the line segment.


 Line segments are output using the first syntax.
  <a id="datatype-geometric-boxes"></a>

### Boxes


 Boxes are represented by pairs of points that are opposite corners of the box. Values of type `box` are specified using any of the following syntaxes:

```

( ( X1 , Y1 ) , ( X2 , Y2 ) )
  ( X1 , Y1 ) , ( X2 , Y2 )
    X1 , Y1   ,   X2 , Y2
```
 where <code>(</code><em>x1</em><code>,</code><em>y1</em><code>)</code> and <code>(</code><em>x2</em><code>,</code><em>y2</em><code>)</code> are any two opposite corners of the box.


 Boxes are output using the second syntax.


 Any two opposite corners can be supplied on input, but the values will be reordered as needed to store the upper right and lower left corners, in that order.
  <a id="datatype-geometric-paths"></a>

### Paths


 Paths are represented by lists of connected points. Paths can be *open*, where the first and last points in the list are considered not connected, or *closed*, where the first and last points are considered connected.


 Values of type `path` are specified using any of the following syntaxes:

```

[ ( X1 , Y1 ) , ... , ( XN , YN ) ]
( ( X1 , Y1 ) , ... , ( XN , YN ) )
  ( X1 , Y1 ) , ... , ( XN , YN )
  ( X1 , Y1   , ... ,   XN , YN )
    X1 , Y1   , ... ,   XN , YN
```
 where the points are the end points of the line segments comprising the path. Square brackets (`[]`) indicate an open path, while parentheses (`()`) indicate a closed path. When the outermost parentheses are omitted, as in the third through fifth syntaxes, a closed path is assumed.


 Paths are output using the first or second syntax, as appropriate.
  <a id="datatype-polygon"></a>

### Polygons


 Polygons are represented by lists of points (the vertices of the polygon). Polygons are very similar to closed paths; the essential semantic difference is that a polygon is considered to include the area within it, while a path is not.


 An important implementation difference between polygons and paths is that the stored representation of a polygon includes its smallest bounding box. This speeds up certain search operations, although computing the bounding box adds overhead while constructing new polygons.


 Values of type `polygon` are specified using any of the following syntaxes:

```

( ( X1 , Y1 ) , ... , ( XN , YN ) )
  ( X1 , Y1 ) , ... , ( XN , YN )
  ( X1 , Y1   , ... ,   XN , YN )
    X1 , Y1   , ... ,   XN , YN
```
 where the points are the end points of the line segments comprising the boundary of the polygon.


 Polygons are output using the first syntax.
  <a id="datatype-circle"></a>

### Circles


 Circles are represented by a center point and radius. Values of type `circle` are specified using any of the following syntaxes:

```

< ( X , Y ) , R >
( ( X , Y ) , R )
  ( X , Y ) , R
    X , Y   , R
```
 where <code>(</code><em>x</em><code>,</code><em>y</em><code>)</code> is the center point and *r* is the radius of the circle.


 Circles are output using the first syntax.
