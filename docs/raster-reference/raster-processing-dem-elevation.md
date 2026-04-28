<a id="Raster_Processing_DEM"></a>

## Raster Processing: DEM (Elevation)
  <a id="RT_ST_Aspect"></a>

# ST_Aspect

Returns the aspect (in degrees by default) of an elevation raster band. Useful for analyzing terrain.

## Synopsis


```sql
raster ST_Aspect(raster  rast, integer  band=1, text  pixeltype=32BF, text  units=DEGREES, boolean  interpolate_nodata=FALSE)
raster ST_Aspect(raster  rast, integer  band, raster  customextent, text  pixeltype=32BF, text  units=DEGREES, boolean  interpolate_nodata=FALSE)
```


## Description


Returns the aspect (in degrees by default) of an elevation raster band. Utilizes map algebra and applies the aspect equation to neighboring pixels.


 `units` indicates the units of the aspect. Possible values are: RADIANS, DEGREES (default).


 When `units` = RADIANS, values are between 0 and 2 * pi radians measured clockwise from North.


 When `units` = DEGREES, values are between 0 and 360 degrees measured clockwise from North.


 If slope of pixel is zero, aspect of pixel is -1.


!!! note

    For more information about Slope, Aspect and Hillshade, please refer to [ESRI - How hillshade works](http://webhelp.esri.com/arcgisdesktop/9.3/index.cfm?TopicName=How%20Hillshade%20works) and [ERDAS Field Guide - Aspect Images](http://e2b.erdas.com/fieldguide/wwhelp/wwhimpl/common/html/wwhelp.htm?context=FieldGuide&file=Aspect_Images.html).


Availability: 2.0.0


Enhanced: 2.1.0 Uses ST_MapAlgebra() and added optional `interpolate_nodata` function parameter


Changed: 2.1.0 In prior versions, return values were in radians. Now, return values default to degrees


## Examples: Variant 1


```sql

WITH foo AS (
    SELECT ST_SetValues(
        ST_AddBand(ST_MakeEmptyRaster(5, 5, 0, 0, 1, -1, 0, 0, 0), 1, '32BF', 0, -9999),
        1, 1, 1, ARRAY[
            [1, 1, 1, 1, 1],
            [1, 2, 2, 2, 1],
            [1, 2, 3, 2, 1],
            [1, 2, 2, 2, 1],
            [1, 1, 1, 1, 1]
        ]::double precision[][]
    ) AS rast
)
SELECT
    ST_DumpValues(ST_Aspect(rast, 1, '32BF'))
FROM foo

                                                                                                    st_dumpvalues

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------
 (1,"{{315,341.565063476562,0,18.4349479675293,45},{288.434936523438,315,0,45,71.5650482177734},{270,270,-1,90,90},{251.565048217773,225,180,135,108.434951782227},{225,198.43495178
2227,180,161.565048217773,135}}")
(1 row)

```


## Examples: Variant 2


Complete example of tiles of a coverage. This query only works with PostgreSQL 9.1 or higher.


```sql

WITH foo AS (
    SELECT ST_Tile(
        ST_SetValues(
            ST_AddBand(
                ST_MakeEmptyRaster(6, 6, 0, 0, 1, -1, 0, 0, 0),
                1, '32BF', 0, -9999
            ),
            1, 1, 1, ARRAY[
                [1, 1, 1, 1, 1, 1],
                [1, 1, 1, 1, 2, 1],
                [1, 2, 2, 3, 3, 1],
                [1, 1, 3, 2, 1, 1],
                [1, 2, 2, 1, 2, 1],
                [1, 1, 1, 1, 1, 1]
            ]::double precision[]
        ),
        2, 2
    ) AS rast
)
SELECT
    t1.rast,
    ST_Aspect(ST_Union(t2.rast), 1, t1.rast)
FROM foo t1
CROSS JOIN foo t2
WHERE ST_Intersects(t1.rast, t2.rast)
GROUP BY t1.rast;

```


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra), [RT_ST_TRI](#RT_ST_TRI), [RT_ST_TPI](#RT_ST_TPI), [RT_ST_Roughness](#RT_ST_Roughness), [RT_ST_HillShade](#RT_ST_HillShade), [RT_ST_Slope](#RT_ST_Slope)
  <a id="RT_ST_HillShade"></a>

# ST_HillShade

Returns the hypothetical illumination of an elevation raster band using provided azimuth, altitude, brightness and scale inputs.

## Synopsis


```sql
raster ST_HillShade(raster  rast, integer  band=1, text  pixeltype=32BF, double precision  azimuth=315, double precision  altitude=45, double precision  max_bright=255, double precision  scale=1.0, boolean  interpolate_nodata=FALSE)
raster ST_HillShade(raster  rast, integer  band, raster  customextent, text  pixeltype=32BF, double precision  azimuth=315, double precision  altitude=45, double precision  max_bright=255, double precision  scale=1.0, boolean  interpolate_nodata=FALSE)
```


## Description


Returns the hypothetical illumination of an elevation raster band using the azimuth, altitude, brightness, and scale inputs. Utilizes map algebra and applies the hill shade equation to neighboring pixels. Return pixel values are between 0 and 255.


 `azimuth` is a value between 0 and 360 degrees measured clockwise from North.


 `altitude` is a value between 0 and 90 degrees where 0 degrees is at the horizon and 90 degrees is directly overhead.


 `max_bright` is a value between 0 and 255 with 0 as no brightness and 255 as max brightness.


 `scale` is the ratio of vertical units to horizontal. For Feet:LatLon use scale=370400, for Meters:LatLon use scale=111120.


 If `interpolate_nodata` is TRUE, values for NODATA pixels from the input raster will be interpolated using [RT_ST_InvDistWeight4ma](built-in-map-algebra-callback-functions.md#RT_ST_InvDistWeight4ma) before computing the hillshade illumination.


!!! note

    For more information about Hillshade, please refer to [How hillshade works](http://webhelp.esri.com/arcgisdesktop/9.3/index.cfm?TopicName=How%20Hillshade%20works).


Availability: 2.0.0


Enhanced: 2.1.0 Uses ST_MapAlgebra() and added optional `interpolate_nodata` function parameter


Changed: 2.1.0 In prior versions, azimuth and altitude were expressed in radians. Now, azimuth and altitude are expressed in degrees


## Examples: Variant 1


```sql

WITH foo AS (
    SELECT ST_SetValues(
        ST_AddBand(ST_MakeEmptyRaster(5, 5, 0, 0, 1, -1, 0, 0, 0), 1, '32BF', 0, -9999),
        1, 1, 1, ARRAY[
            [1, 1, 1, 1, 1],
            [1, 2, 2, 2, 1],
            [1, 2, 3, 2, 1],
            [1, 2, 2, 2, 1],
            [1, 1, 1, 1, 1]
        ]::double precision[][]
    ) AS rast
)
SELECT
    ST_DumpValues(ST_Hillshade(rast, 1, '32BF'))
FROM foo

                                                                                                                       st_dumpvalues

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
-----------------------------------------------------------------------
 (1,"{{NULL,NULL,NULL,NULL,NULL},{NULL,251.32763671875,220.749786376953,147.224319458008,NULL},{NULL,220.749786376953,180.312225341797,67.7497863769531,NULL},{NULL,147.224319458008
,67.7497863769531,43.1210060119629,NULL},{NULL,NULL,NULL,NULL,NULL}}")
(1 row)

```


## Examples: Variant 2


Complete example of tiles of a coverage. This query only works with PostgreSQL 9.1 or higher.


```sql

WITH foo AS (
    SELECT ST_Tile(
        ST_SetValues(
            ST_AddBand(
                ST_MakeEmptyRaster(6, 6, 0, 0, 1, -1, 0, 0, 0),
                1, '32BF', 0, -9999
            ),
            1, 1, 1, ARRAY[
                [1, 1, 1, 1, 1, 1],
                [1, 1, 1, 1, 2, 1],
                [1, 2, 2, 3, 3, 1],
                [1, 1, 3, 2, 1, 1],
                [1, 2, 2, 1, 2, 1],
                [1, 1, 1, 1, 1, 1]
            ]::double precision[]
        ),
        2, 2
    ) AS rast
)
SELECT
    t1.rast,
    ST_Hillshade(ST_Union(t2.rast), 1, t1.rast)
FROM foo t1
CROSS JOIN foo t2
WHERE ST_Intersects(t1.rast, t2.rast)
GROUP BY t1.rast;

```


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra), [RT_ST_TRI](#RT_ST_TRI), [RT_ST_TPI](#RT_ST_TPI), [RT_ST_Roughness](#RT_ST_Roughness), [RT_ST_Aspect](#RT_ST_Aspect), [RT_ST_Slope](#RT_ST_Slope)
  <a id="RT_ST_Roughness"></a>

# ST_Roughness

Returns a raster with the calculated "roughness" of a DEM.

## Synopsis


```sql
raster ST_Roughness(raster  rast, integer  nband, raster  customextent, text  pixeltype="32BF", boolean   interpolate_nodata=FALSE)
```


## Description


Calculates the "roughness" of a DEM, by subtracting the maximum from the minimum for a given area.


Availability: 2.1.0


## Examples


```

-- needs examples

```


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra), [RT_ST_TRI](#RT_ST_TRI), [RT_ST_TPI](#RT_ST_TPI), [RT_ST_Slope](#RT_ST_Slope), [RT_ST_HillShade](#RT_ST_HillShade), [RT_ST_Aspect](#RT_ST_Aspect)
  <a id="RT_ST_Slope"></a>

# ST_Slope

Returns the slope (in degrees by default) of an elevation raster band. Useful for analyzing terrain.

## Synopsis


```sql
raster ST_Slope(raster  rast, integer  nband=1, text  pixeltype=32BF, text  units=DEGREES, double precision  scale=1.0, boolean  interpolate_nodata=FALSE)
raster ST_Slope(raster  rast, integer  nband, raster  customextent, text  pixeltype=32BF, text  units=DEGREES, double precision  scale=1.0, boolean  interpolate_nodata=FALSE)
```


## Description


Returns the slope (in degrees by default) of an elevation raster band. Utilizes map algebra and applies the slope equation to neighboring pixels.


 `units` indicates the units of the slope. Possible values are: RADIANS, DEGREES (default), PERCENT.


 `scale` is the ratio of vertical units to horizontal. For Feet:LatLon use scale=370400, for Meters:LatLon use scale=111120.


 If `interpolate_nodata` is TRUE, values for NODATA pixels from the input raster will be interpolated using [RT_ST_InvDistWeight4ma](built-in-map-algebra-callback-functions.md#RT_ST_InvDistWeight4ma) before computing the surface slope.


!!! note

    For more information about Slope, Aspect and Hillshade, please refer to [ESRI - How hillshade works](http://webhelp.esri.com/arcgisdesktop/9.3/index.cfm?TopicName=How%20Hillshade%20works) and [ERDAS Field Guide - Slope Images](http://e2b.erdas.com/fieldguide/wwhelp/wwhimpl/common/html/wwhelp.htm?context=FieldGuide&file=Slope_Images.html).


Availability: 2.0.0


Enhanced: 2.1.0 Uses ST_MapAlgebra() and added optional `units`, `scale`, `interpolate_nodata` function parameters


Changed: 2.1.0 In prior versions, return values were in radians. Now, return values default to degrees


## Examples: Variant 1


```sql

WITH foo AS (
    SELECT ST_SetValues(
        ST_AddBand(ST_MakeEmptyRaster(5, 5, 0, 0, 1, -1, 0, 0, 0), 1, '32BF', 0, -9999),
        1, 1, 1, ARRAY[
            [1, 1, 1, 1, 1],
            [1, 2, 2, 2, 1],
            [1, 2, 3, 2, 1],
            [1, 2, 2, 2, 1],
            [1, 1, 1, 1, 1]
        ]::double precision[][]
    ) AS rast
)
SELECT
    ST_DumpValues(ST_Slope(rast, 1, '32BF'))
FROM foo

                            st_dumpvalues

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------
 (1,"{{10.0249881744385,21.5681285858154,26.5650520324707,21.5681285858154,10.0249881744385},{21.5681285858154,35.2643890380859,36.8698959350586,35.2643890380859,21.5681285858154},
{26.5650520324707,36.8698959350586,0,36.8698959350586,26.5650520324707},{21.5681285858154,35.2643890380859,36.8698959350586,35.2643890380859,21.5681285858154},{10.0249881744385,21.
5681285858154,26.5650520324707,21.5681285858154,10.0249881744385}}")
(1 row)

```


## Examples: Variant 2


Complete example of tiles of a coverage. This query only works with PostgreSQL 9.1 or higher.


```sql

WITH foo AS (
    SELECT ST_Tile(
        ST_SetValues(
            ST_AddBand(
                ST_MakeEmptyRaster(6, 6, 0, 0, 1, -1, 0, 0, 0),
                1, '32BF', 0, -9999
            ),
            1, 1, 1, ARRAY[
                [1, 1, 1, 1, 1, 1],
                [1, 1, 1, 1, 2, 1],
                [1, 2, 2, 3, 3, 1],
                [1, 1, 3, 2, 1, 1],
                [1, 2, 2, 1, 2, 1],
                [1, 1, 1, 1, 1, 1]
            ]::double precision[]
        ),
        2, 2
    ) AS rast
)
SELECT
    t1.rast,
    ST_Slope(ST_Union(t2.rast), 1, t1.rast)
FROM foo t1
CROSS JOIN foo t2
WHERE ST_Intersects(t1.rast, t2.rast)
GROUP BY t1.rast;

```


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra), [RT_ST_TRI](#RT_ST_TRI), [RT_ST_TPI](#RT_ST_TPI), [RT_ST_Roughness](#RT_ST_Roughness), [RT_ST_HillShade](#RT_ST_HillShade), [RT_ST_Aspect](#RT_ST_Aspect)
  <a id="RT_ST_TPI"></a>

# ST_TPI

Returns a raster with the calculated Topographic Position Index.

## Synopsis


```sql
raster ST_TPI(raster  rast, integer  nband, raster  customextent, text  pixeltype="32BF", boolean   interpolate_nodata=FALSE)
```


## Description


Calculates the Topographic Position Index, which is defined as the focal mean with radius of one minus the center cell.


!!! note

    This function only supports a focalmean radius of one.


Availability: 2.1.0


## Examples


```

-- needs examples

```


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra), [RT_ST_TRI](#RT_ST_TRI), [RT_ST_Roughness](#RT_ST_Roughness), [RT_ST_Slope](#RT_ST_Slope), [RT_ST_HillShade](#RT_ST_HillShade), [RT_ST_Aspect](#RT_ST_Aspect)
  <a id="RT_ST_TRI"></a>

# ST_TRI

Returns a raster with the calculated Terrain Ruggedness Index.

## Synopsis


```sql
raster ST_TRI(raster  rast, integer  nband, raster  customextent, text  pixeltype="32BF", boolean   interpolate_nodata=FALSE)
```


## Description


 Terrain Ruggedness Index is calculated by comparing a central pixel with its neighbors, taking the absolute values of the differences, and averaging the result.


!!! note

    This function only supports a focalmean radius of one.


Availability: 2.1.0


## Examples


```

-- needs examples

```


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra), [RT_ST_Roughness](#RT_ST_Roughness), [RT_ST_TPI](#RT_ST_TPI), [RT_ST_Slope](#RT_ST_Slope), [RT_ST_HillShade](#RT_ST_HillShade), [RT_ST_Aspect](#RT_ST_Aspect)
