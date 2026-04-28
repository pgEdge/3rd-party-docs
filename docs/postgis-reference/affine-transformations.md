<a id="Affine_Transformation"></a>

## Affine Transformations
  <a id="ST_Affine"></a>

# ST_Affine

Apply a 3D affine transformation to a geometry.

## Synopsis


```sql
geometry ST_Affine(geometry  geomA, float  a, float  b, float  c, float  d, float  e, float  f, float  g, float  h, float  i, float  xoff, float  yoff, float  zoff)
geometry ST_Affine(geometry  geomA, float  a, float  b, float  d, float  e, float  xoff, float  yoff)
```


## Description


Applies a 3D affine transformation to the geometry to do things like translate, rotate, scale in one step.


 Version 1: The call

```
ST_Affine(geom, a, b, c, d, e, f, g, h, i, xoff, yoff, zoff)
```
 represents the transformation matrix

```
/ a  b  c  xoff \
| d  e  f  yoff |
| g  h  i  zoff |
\ 0  0  0     1 /
```
 and the vertices are transformed as follows:

```
x' = a*x + b*y + c*z + xoff
y' = d*x + e*y + f*z + yoff
z' = g*x + h*y + i*z + zoff
```
 All of the translate / scale functions below are expressed via such an affine transformation.


Version 2: Applies a 2d affine transformation to the geometry. The call

```
ST_Affine(geom, a, b, d, e, xoff, yoff)
```
 represents the transformation matrix

```
/  a  b  0  xoff  \       /  a  b  xoff  \
|  d  e  0  yoff  | rsp.  |  d  e  yoff  |
|  0  0  1     0  |       \  0  0     1  /
\  0  0  0     1  /
```
 and the vertices are transformed as follows:

```
x' = a*x + b*y + xoff
y' = d*x + e*y + yoff
z' = z
```
 This method is a subcase of the 3D method above.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Availability: 1.1.2. Name changed from Affine to ST_Affine in 1.2.2


!!! note

    Prior to 1.3.4, this function crashes if used with geometries that contain CURVES. This is fixed in 1.3.4+


## Examples


```

--Rotate a 3d line 180 degrees about the z axis.  Note this is long-hand for doing ST_Rotate();
 SELECT ST_AsEWKT(ST_Affine(geom,  cos(pi()), -sin(pi()), 0,  sin(pi()), cos(pi()), 0,  0, 0, 1,  0, 0, 0)) As using_affine,
	 ST_AsEWKT(ST_Rotate(geom, pi())) As using_rotate
	FROM (SELECT ST_GeomFromEWKT('LINESTRING(1 2 3, 1 4 3)') As geom) As foo;
        using_affine         |        using_rotate
-----------------------------+-----------------------------
 LINESTRING(-1 -2 3,-1 -4 3) | LINESTRING(-1 -2 3,-1 -4 3)
(1 row)

--Rotate a 3d line 180 degrees in both the x and z axis
SELECT ST_AsEWKT(ST_Affine(geom, cos(pi()), -sin(pi()), 0, sin(pi()), cos(pi()), -sin(pi()), 0, sin(pi()), cos(pi()), 0, 0, 0))
	FROM (SELECT ST_GeomFromEWKT('LINESTRING(1 2 3, 1 4 3)') As geom) As foo;
           st_asewkt
-------------------------------
 LINESTRING(-1 -2 -3,-1 -4 -3)
(1 row)

```


## See Also


[ST_Rotate](#ST_Rotate), [ST_Scale](#ST_Scale), [ST_Translate](#ST_Translate), [ST_TransScale](#ST_TransScale)
  <a id="ST_Rotate"></a>

# ST_Rotate

Rotates a geometry about an origin point.

## Synopsis


```sql
geometry ST_Rotate(geometry geomA, float rotRadians)
geometry ST_Rotate(geometry geomA, float rotRadians, float x0, float y0)
geometry ST_Rotate(geometry geomA, float rotRadians, geometry pointOrigin)
```


## Description


Rotates geometry rotRadians counter-clockwise about the origin point. The rotation origin can be specified either as a POINT geometry, or as x and y coordinates. If the origin is not specified, the geometry is rotated about POINT(0 0).


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Enhanced: 2.0.0 additional parameters for specifying the origin of rotation were added.


Availability: 1.1.2. Name changed from Rotate to ST_Rotate in 1.2.2


## Examples


```

--Rotate 180 degrees
SELECT ST_AsEWKT(ST_Rotate('LINESTRING (50 160, 50 50, 100 50)', pi()));
               st_asewkt
---------------------------------------
 LINESTRING(-50 -160,-50 -50,-100 -50)
(1 row)

--Rotate 30 degrees counter-clockwise at x=50, y=160
SELECT ST_AsEWKT(ST_Rotate('LINESTRING (50 160, 50 50, 100 50)', pi()/6, 50, 160));
                                 st_asewkt
---------------------------------------------------------------------------
 LINESTRING(50 160,105 64.7372055837117,148.301270189222 89.7372055837117)
(1 row)

--Rotate 60 degrees clockwise from centroid
SELECT ST_AsEWKT(ST_Rotate(geom, -pi()/3, ST_Centroid(geom)))
FROM (SELECT 'LINESTRING (50 160, 50 50, 100 50)'::geometry AS geom) AS foo;
                           st_asewkt
--------------------------------------------------------------
 LINESTRING(116.4225 130.6721,21.1597 75.6721,46.1597 32.3708)
(1 row)

```


## See Also


[ST_Affine](#ST_Affine), [ST_RotateX](#ST_RotateX), [ST_RotateY](#ST_RotateY), [ST_RotateZ](#ST_RotateZ)
  <a id="ST_RotateX"></a>

# ST_RotateX

Rotates a geometry about the X axis.

## Synopsis


```sql
geometry ST_RotateX(geometry geomA, float rotRadians)
```


## Description


Rotates a geometry geomA - rotRadians about the X axis.


!!! note

    <code>ST_RotateX(geomA, rotRadians)</code> is short-hand for <code>ST_Affine(geomA, 1, 0, 0, 0, cos(rotRadians), -sin(rotRadians), 0, sin(rotRadians), cos(rotRadians), 0, 0, 0)</code>.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Availability: 1.1.2. Name changed from RotateX to ST_RotateX in 1.2.2


## Examples


```

--Rotate a line 90 degrees along x-axis
SELECT ST_AsEWKT(ST_RotateX(ST_GeomFromEWKT('LINESTRING(1 2 3, 1 1 1)'), pi()/2));
		 st_asewkt
---------------------------
 LINESTRING(1 -3 2,1 -1 1)
```


## See Also


[ST_Affine](#ST_Affine), [ST_RotateY](#ST_RotateY), [ST_RotateZ](#ST_RotateZ)
  <a id="ST_RotateY"></a>

# ST_RotateY

Rotates a geometry about the Y axis.

## Synopsis


```sql
geometry ST_RotateY(geometry geomA, float rotRadians)
```


## Description


Rotates a geometry geomA - rotRadians about the y axis.


!!! note

    <code>ST_RotateY(geomA, rotRadians)</code> is short-hand for <code>ST_Affine(geomA, cos(rotRadians), 0, sin(rotRadians), 0, 1, 0, -sin(rotRadians), 0, cos(rotRadians), 0, 0, 0)</code>.


Availability: 1.1.2. Name changed from RotateY to ST_RotateY in 1.2.2


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


## Examples


```

--Rotate a line 90 degrees along y-axis
 SELECT ST_AsEWKT(ST_RotateY(ST_GeomFromEWKT('LINESTRING(1 2 3, 1 1 1)'), pi()/2));
		 st_asewkt
---------------------------
 LINESTRING(3 2 -1,1 1 -1)
```


## See Also


[ST_Affine](#ST_Affine), [ST_RotateX](#ST_RotateX), [ST_RotateZ](#ST_RotateZ)
  <a id="ST_RotateZ"></a>

# ST_RotateZ

Rotates a geometry about the Z axis.

## Synopsis


```sql
geometry ST_RotateZ(geometry geomA, float rotRadians)
```


## Description


Rotates a geometry geomA - rotRadians about the Z axis.


!!! note

    This is a synonym for ST_Rotate


!!! note

    <code>ST_RotateZ(geomA, rotRadians)</code> is short-hand for <code>SELECT ST_Affine(geomA, cos(rotRadians), -sin(rotRadians), 0, sin(rotRadians), cos(rotRadians), 0, 0, 0, 1, 0, 0, 0)</code>.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Availability: 1.1.2. Name changed from RotateZ to ST_RotateZ in 1.2.2


!!! note

    Prior to 1.3.4, this function crashes if used with geometries that contain CURVES. This is fixed in 1.3.4+


## Examples


```

--Rotate a line 90 degrees along z-axis
SELECT ST_AsEWKT(ST_RotateZ(ST_GeomFromEWKT('LINESTRING(1 2 3, 1 1 1)'), pi()/2));
		 st_asewkt
---------------------------
 LINESTRING(-2 1 3,-1 1 1)

 --Rotate a curved circle around z-axis
SELECT ST_AsEWKT(ST_RotateZ(geom, pi()/2))
FROM (SELECT ST_LineToCurve(ST_Buffer(ST_GeomFromText('POINT(234 567)'), 3)) As geom) As foo;

													   st_asewkt
----------------------------------------------------------------------------------------------------------------------------
 CURVEPOLYGON(CIRCULARSTRING(-567 237,-564.87867965644 236.12132034356,-564 234,-569.12132034356 231.87867965644,-567 237))
```


## See Also


[ST_Affine](#ST_Affine), [ST_RotateX](#ST_RotateX), [ST_RotateY](#ST_RotateY)
  <a id="ST_Scale"></a>

# ST_Scale

Scales a geometry by given factors.

## Synopsis


```sql
geometry ST_Scale(geometry  geomA, float XFactor, float YFactor, float ZFactor)
geometry ST_Scale(geometry  geomA, float XFactor, float YFactor)
geometry ST_Scale(geometry  geom, geometry factor)
geometry ST_Scale(geometry  geom, geometry factor, geometry origin)
```


## Description


Scales the geometry to a new size by multiplying the ordinates with the corresponding factor parameters.


 The version taking a geometry as the `factor` parameter allows passing a 2d, 3dm, 3dz or 4d point to set scaling factor for all supported dimensions. Missing dimensions in the `factor` point are equivalent to no scaling the corresponding dimension.


 The three-geometry variant allows a "false origin" for the scaling to be passed in. This allows "scaling in place", for example using the centroid of the geometry as the false origin. Without a false origin, scaling takes place relative to the actual origin, so all coordinates are just multiplied by the scale factor.


!!! note

    Prior to 1.3.4, this function crashes if used with geometries that contain CURVES. This is fixed in 1.3.4+


Availability: 1.1.0.


Enhanced: 2.0.0 support for Polyhedral surfaces, Triangles and TIN was introduced.


Enhanced: 2.2.0 support for scaling all dimension (`factor` parameter) was introduced.


Enhanced: 2.5.0 support for scaling relative to a local origin (`origin` parameter) was introduced.


## Examples


```
--Version 1: scale X, Y, Z
SELECT ST_AsEWKT(ST_Scale(ST_GeomFromEWKT('LINESTRING(1 2 3, 1 1 1)'), 0.5, 0.75, 0.8));
			  st_asewkt
--------------------------------------
 LINESTRING(0.5 1.5 2.4,0.5 0.75 0.8)

--Version 2: Scale X Y
 SELECT ST_AsEWKT(ST_Scale(ST_GeomFromEWKT('LINESTRING(1 2 3, 1 1 1)'), 0.5, 0.75));
			st_asewkt
----------------------------------
 LINESTRING(0.5 1.5 3,0.5 0.75 1)

--Version 3: Scale X Y Z M
 SELECT ST_AsEWKT(ST_Scale(ST_GeomFromEWKT('LINESTRING(1 2 3 4, 1 1 1 1)'),
   ST_MakePoint(0.5, 0.75, 2, -1)));
			       st_asewkt
----------------------------------------
 LINESTRING(0.5 1.5 6 -4,0.5 0.75 2 -1)

--Version 4: Scale X Y using false origin
SELECT ST_AsText(ST_Scale('LINESTRING(1 1, 2 2)', 'POINT(2 2)', 'POINT(1 1)'::geometry));
      st_astext
---------------------
 LINESTRING(1 1,3 3)
```


## See Also


[ST_Affine](#ST_Affine), [ST_TransScale](#ST_TransScale)
  <a id="ST_Translate"></a>

# ST_Translate

Translates a geometry by given offsets.

## Synopsis


```sql
geometry ST_Translate(geometry  g1, float  deltax, float  deltay)
geometry ST_Translate(geometry  g1, float  deltax, float  deltay, float  deltaz)
```


## Description


Returns a new geometry whose coordinates are translated delta x,delta y,delta z units. Units are based on the units defined in spatial reference (SRID) for this geometry.


!!! note

    Prior to 1.3.4, this function crashes if used with geometries that contain CURVES. This is fixed in 1.3.4+


Availability: 1.2.2


## Examples


Move a point 1 degree longitude


```sql

	SELECT ST_AsText(ST_Translate(ST_GeomFromText('POINT(-71.01 42.37)',4326),1,0)) As wgs_transgeomtxt;

	wgs_transgeomtxt
	---------------------
	POINT(-70.01 42.37)

```


Move a linestring 1 degree longitude and 1/2 degree latitude


```sql
SELECT ST_AsText(ST_Translate(ST_GeomFromText('LINESTRING(-71.01 42.37,-71.11 42.38)',4326),1,0.5)) As wgs_transgeomtxt;
		   wgs_transgeomtxt
	---------------------------------------
	LINESTRING(-70.01 42.87,-70.11 42.88)

```


Move a 3d point


```sql
SELECT ST_AsEWKT(ST_Translate(CAST('POINT(0 0 0)' As geometry), 5, 12,3));
	st_asewkt
	---------
	POINT(5 12 3)

```


Move a curve and a point


```sql
SELECT ST_AsText(ST_Translate(ST_Collect('CURVEPOLYGON(CIRCULARSTRING(4 3,3.12 0.878,1 0,-1.121 5.1213,6 7, 8 9,4 3))','POINT(1 3)'),1,2));
														 st_astext
------------------------------------------------------------------------------------------------------------
 GEOMETRYCOLLECTION(CURVEPOLYGON(CIRCULARSTRING(5 5,4.12 2.878,2 2,-0.121 7.1213,7 9,9 11,5 5)),POINT(2 5))
```


## See Also


[ST_Affine](#ST_Affine), [ST_AsText](geometry-output.md#ST_AsText), [ST_GeomFromText](geometry-input.md#ST_GeomFromText)
  <a id="ST_TransScale"></a>

# ST_TransScale

Translates and scales a geometry by given offsets and factors.

## Synopsis


```sql
geometry ST_TransScale(geometry  geomA, float deltaX, float deltaY, float XFactor, float YFactor)
```


## Description


Translates the geometry using the deltaX and deltaY args, then scales it using the XFactor, YFactor args, working in 2D only.


!!! note

    <code>ST_TransScale(geomA, deltaX, deltaY, XFactor, YFactor)</code> is short-hand for <code>ST_Affine(geomA, XFactor, 0, 0, 0, YFactor, 0, 0, 0, 1, deltaX*XFactor, deltaY*YFactor, 0)</code>.


!!! note

    Prior to 1.3.4, this function crashes if used with geometries that contain CURVES. This is fixed in 1.3.4+


Availability: 1.1.0.


## Examples


```sql
SELECT ST_AsEWKT(ST_TransScale(ST_GeomFromEWKT('LINESTRING(1 2 3, 1 1 1)'), 0.5, 1, 1, 2));
		  st_asewkt
-----------------------------
 LINESTRING(1.5 6 3,1.5 4 1)


--Buffer a point to get an approximation of a circle, convert to curve and then translate 1,2 and scale it 3,4
  SELECT ST_AsText(ST_Transscale(ST_LineToCurve(ST_Buffer('POINT(234 567)', 3)),1,2,3,4));
														  st_astext
------------------------------------------------------------------------------------------------------------------------------
 CURVEPOLYGON(CIRCULARSTRING(714 2276,711.363961030679 2267.51471862576,705 2264,698.636038969321 2284.48528137424,714 2276))
```


## See Also


[ST_Affine](#ST_Affine), [ST_Translate](#ST_Translate)
