<a id="sfcgal_accessors"></a>

## SFCGAL Accessors and Setters
   <a id="CG_ForceLHR"></a>

# CG_ForceLHR

Force LHR orientation

## Synopsis


```sql
geometry CG_ForceLHR(geometry geom)
```


## Description


Availability: 3.5.0


  <a id="CG_IsPlanar"></a>

# CG_IsPlanar

Check if a surface is or not planar

## Synopsis


```sql
boolean CG_IsPlanar(geometry geom)
```


## Description


Availability: 3.5.0


  <a id="CG_IsSolid"></a>

# CG_IsSolid

Test if the geometry is a solid. No validity check is performed.

## Synopsis


```sql
boolean CG_IsSolid(geometry geom1)
```


## Description


Availability: 3.5.0


  <a id="CG_MakeSolid"></a>

# CG_MakeSolid

Cast the geometry into a solid. No check is performed. To obtain a valid solid, the input geometry must be a closed Polyhedral Surface or a closed TIN.

## Synopsis


```sql
geometry CG_MakeSolid(geometry geom1)
```


## Description


Availability: 3.5.0


  <a id="CG_Orientation"></a>

# CG_Orientation

Determine surface orientation

## Synopsis


```sql
integer CG_Orientation(geometry geom)
```


## Description


The function only applies to polygons. It returns -1 if the polygon is counterclockwise oriented and 1 if the polygon is clockwise oriented.


Availability: 3.5.0


  <a id="CG_Area"></a>

# CG_Area

Calculates the area of a geometry

## Synopsis


```sql
double precision CG_Area(geometry
                            geom)
```


## Description


Calculates the area of a geometry.


Performed by the SFCGAL module


!!! note

    NOTE: this function returns a double precision value representing the area.


Availability: 3.5.0


## Geometry Examples


```sql
SELECT CG_Area('Polygon ((0 0, 0 5, 5 5, 5 0, 0 0), (1 1, 2 1, 2 2, 1 2, 1 1), (3 3, 4 3, 4 4, 3 4, 3 3))');
                cg_area
                --------
                25
                (1 row)
```


## See Also


[ST_3DArea](#ST_3DArea), [ST_Area](../postgis-reference/measurement-functions.md#ST_Area)
  <a id="CG_3DArea"></a>

# CG_3DArea

Computes area of 3D surface geometries. Will return 0 for solids.

## Synopsis


```sql
floatCG_3DArea(geometry geom1)
```


## Description


Availability: 3.5.0


 SQL-MM IEC 13249-3: 8.1, 10.5


## Examples


Note: By default a PolyhedralSurface built from WKT is a surface geometry, not solid. It therefore has surface area. Once converted to a solid, no area.


```sql
SELECT CG_3DArea(geom) As cube_surface_area,
            CG_3DArea(CG_MakeSolid(geom)) As solid_surface_area
            FROM (SELECT 'POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
            ((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)),
            ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
            ((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
            ((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)),
            ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )'::geometry) As f(geom);

            cube_surface_area | solid_surface_area
            -------------------+--------------------
            6 |                  0
```


## See Also


[CG_Area](#CG_Area), [CG_MakeSolid](#CG_MakeSolid), [CG_IsSolid](#CG_IsSolid), [CG_Area](#CG_Area)
  <a id="CG_Volume"></a>

# CG_Volume

Computes the volume of a 3D solid. If applied to surface (even closed) geometries will return 0.

## Synopsis


```sql
float CG_Volume(geometry geom1)
```


## Description


Availability: 3.5.0


 SQL-MM IEC 13249-3: 9.1 (same as CG_3DVolume)


## Example


When closed surfaces are created with WKT, they are treated as areal rather than solid. To make them solid, you need to use [CG_MakeSolid](#CG_MakeSolid). Areal geometries have no volume. Here is an example to demonstrate.


```sql
SELECT CG_Volume(geom) As cube_surface_vol,
    CG_Volume(CG_MakeSolid(geom)) As solid_surface_vol
    FROM (SELECT 'POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
    ((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)),
    ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
    ((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
    ((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)),
    ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )'::geometry) As f(geom);

    cube_surface_vol | solid_surface_vol
    ------------------+-------------------
    0 |                 1
```


## See Also


[CG_3DArea](#CG_3DArea), [CG_MakeSolid](#CG_MakeSolid), [CG_IsSolid](#CG_IsSolid)
  <a id="ST_ForceLHR"></a>

# ST_ForceLHR

Force LHR orientation

## Synopsis


```sql
geometry ST_ForceLHR(geometry geom)
```


## Description


!!! warning

    [ST_ForceLHR](#ST_ForceLHR) is deprecated as of 3.5.0. Use [CG_ForceLHR](#CG_ForceLHR) instead.


Availability: 2.1.0


  <a id="ST_IsPlanar"></a>

# ST_IsPlanar

Check if a surface is or not planar

## Synopsis


```sql
boolean ST_IsPlanar(geometry geom)
```


## Description


!!! warning

    [ST_IsPlanar](#ST_IsPlanar) is deprecated as of 3.5.0. Use [CG_IsPlanar](#CG_IsPlanar) instead.


Availability: 2.2.0: This was documented in 2.1.0 but got accidentally left out in 2.1 release.


  <a id="ST_IsSolid"></a>

# ST_IsSolid

Test if the geometry is a solid. No validity check is performed.

## Synopsis


```sql
boolean ST_IsSolid(geometry geom1)
```


## Description


!!! warning

    [ST_IsSolid](#ST_IsSolid) is deprecated as of 3.5.0. Use [CG_IsSolid](#CG_IsSolid) instead.


Availability: 2.2.0


  <a id="ST_MakeSolid"></a>

# ST_MakeSolid

Cast the geometry into a solid. No check is performed. To obtain a valid solid, the input geometry must be a closed Polyhedral Surface or a closed TIN.

## Synopsis


```sql
geometry ST_MakeSolid(geometry geom1)
```


## Description


!!! warning

    [ST_MakeSolid](#ST_MakeSolid) is deprecated as of 3.5.0. Use [CG_MakeSolid](#CG_MakeSolid) instead.


Availability: 2.2.0


  <a id="ST_Orientation"></a>

# ST_Orientation

Determine surface orientation

## Synopsis


```sql
integer ST_Orientation(geometry geom)
```


## Description


!!! warning

    [ST_Orientation](#ST_Orientation) is deprecated as of 3.5.0. Use [CG_Orientation](#CG_Orientation) instead.


The function only applies to polygons. It returns -1 if the polygon is counterclockwise oriented and 1 if the polygon is clockwise oriented.


Availability: 2.1.0


  <a id="ST_3DArea"></a>

# ST_3DArea

Computes area of 3D surface geometries. Will return 0 for solids.

## Synopsis


```sql
floatST_3DArea(geometry geom1)
```


## Description


!!! warning

    [ST_3DArea](#ST_3DArea) is deprecated as of 3.5.0. Use [CG_3DArea](#CG_3DArea) instead.


Availability: 2.1.0


 SQL-MM IEC 13249-3: 8.1, 10.5


## Examples


Note: By default a PolyhedralSurface built from WKT is a surface geometry, not solid. It therefore has surface area. Once converted to a solid, no area.


```sql
SELECT ST_3DArea(geom) As cube_surface_area,
            ST_3DArea(ST_MakeSolid(geom)) As solid_surface_area
            FROM (SELECT 'POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
            ((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)),
            ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
            ((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
            ((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)),
            ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )'::geometry) As f(geom);

            cube_surface_area | solid_surface_area
            -------------------+--------------------
            6 |                  0
```


## See Also


[ST_Area](../postgis-reference/measurement-functions.md#ST_Area), [ST_MakeSolid](#ST_MakeSolid), [ST_IsSolid](#ST_IsSolid), [ST_Area](../postgis-reference/measurement-functions.md#ST_Area)
  <a id="ST_Volume"></a>

# ST_Volume

Computes the volume of a 3D solid. If applied to surface (even closed) geometries will return 0.

## Synopsis


```sql
float ST_Volume(geometry geom1)
```


## Description


!!! warning

    [ST_Volume](#ST_Volume) is deprecated as of 3.5.0. Use [CG_Volume](#CG_Volume) instead.


Availability: 2.2.0


 SQL-MM IEC 13249-3: 9.1 (same as ST_3DVolume)


## Example


When closed surfaces are created with WKT, they are treated as areal rather than solid. To make them solid, you need to use [ST_MakeSolid](#ST_MakeSolid). Areal geometries have no volume. Here is an example to demonstrate.


```sql
SELECT ST_Volume(geom) As cube_surface_vol,
    ST_Volume(ST_MakeSolid(geom)) As solid_surface_vol
    FROM (SELECT 'POLYHEDRALSURFACE( ((0 0 0, 0 0 1, 0 1 1, 0 1 0, 0 0 0)),
    ((0 0 0, 0 1 0, 1 1 0, 1 0 0, 0 0 0)),
    ((0 0 0, 1 0 0, 1 0 1, 0 0 1, 0 0 0)),
    ((1 1 0, 1 1 1, 1 0 1, 1 0 0, 1 1 0)),
    ((0 1 0, 0 1 1, 1 1 1, 1 1 0, 0 1 0)),
    ((0 0 1, 1 0 1, 1 1 1, 0 1 1, 0 0 1)) )'::geometry) As f(geom);

    cube_surface_vol | solid_surface_vol
    ------------------+-------------------
    0 |                 1
```


## See Also


[ST_3DArea](#ST_3DArea), [ST_MakeSolid](#ST_MakeSolid), [ST_IsSolid](#ST_IsSolid)
