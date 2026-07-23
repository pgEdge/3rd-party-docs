<a id="Raster_Types"></a>

## Raster Support Data types
  <a id="geomval"></a>

# geomval

A spatial datatype with two fields - geom (holding a geometry object) and val (holding a double precision pixel value from a raster band).

## Description


geomval is a compound data type consisting of a geometry object referenced by the .geom field and val, a double precision value that represents the pixel value at a particular geometric location in a raster band. It is used by the ST_DumpAsPolygon and Raster intersection family of functions as an output type to explode a raster band into geometry polygons.


## See Also


[PostGIS Geometry / Geography / Raster Dump Functions](../postgis-special-functions-index/postgis-geometry-geography-raster-dump-functions.md#PostGIS_Geometry_DumpFunctions)
  <a id="addbandarg"></a>

# addbandarg

A composite type used as input into the ST_AddBand function defining the attributes and initial value of the new band.

## Description


 A composite type used as input into the ST_AddBand function defining the attributes and initial value of the new band.

` index integer `
:   1-based value indicating the position where the new band will be added amongst the raster's bands. If NULL, the new band will be added at the end of the raster's bands.

` pixeltype text `
:   Pixel type of the new band. One of defined pixel types as described in [RT_ST_BandPixelType](raster-band-accessors.md#RT_ST_BandPixelType).

` initialvalue double precision `
:   Initial value that all pixels of new band will be set to.

` nodataval double precision `
:   NODATA value of the new band. If NULL, the new band will not have a NODATA value assigned.


## See Also


 [RT_ST_AddBand](raster-constructors.md#RT_ST_AddBand)
  <a id="rastbandarg"></a>

# rastbandarg

A composite type for use when needing to express a raster and a band index of that raster.

## Description


 A composite type for use when needing to express a raster and a band index of that raster.

` rast raster `
:   The raster in question/

` nband integer `
:   1-based value indicating the band of raster


## See Also


 [RT_ST_MapAlgebra](raster-processing-map-algebra.md#RT_ST_MapAlgebra)
  <a id="raster"></a>

# raster

raster spatial data type.

## Description


raster is a spatial data type used to represent raster data such as those imported from JPEGs, TIFFs, PNGs, digital elevation models. Each raster has 1 or more bands each having a set of pixel values. Rasters can be georeferenced.


!!! note

    Requires PostGIS be compiled with GDAL support. Currently rasters can be implicitly converted to geometry type, but the conversion returns the [RT_ST_ConvexHull](raster-processing-raster-to-geometry.md#RT_ST_ConvexHull) of the raster. This auto casting may be removed in the near future so don't rely on it.


## Casting Behavior


This section lists the automatic as well as explicit casts allowed for this data type


| Cast To | Behavior |
| geometry | automatic |


## See Also


[Raster Reference](index.md#RT_reference)
  <a id="reclassarg"></a>

# reclassarg

A composite type used as input into the ST_Reclass function defining the behavior of reclassification.

## Description


A composite type used as input into the ST_Reclass function defining the behavior of reclassification.


`nband integer`
:   The band number of band to reclassify.

`reclassexpr text`
:   range expression consisting of comma delimited range:map_range mappings. : to define mapping that defines how to map old band values to new band values. ( means >, ) means less than, ] < or equal, [ means > or equal


    ```

    1. [a-b] = a <= x <= b

    2. (a-b] = a < x <= b

    3. [a-b) = a <= x < b

    4. (a-b) = a < x < b
    ```


    ( notation is optional so a-b means the same as (a-b)

`pixeltype text`
:   One of defined pixel types as described in [RT_ST_BandPixelType](raster-band-accessors.md#RT_ST_BandPixelType)

`nodataval double precision`
:   Value to treat as no data. For image outputs that support transparency, these will be blank.


## Example: Reclassify band 2 as an 8BUI where 255 is nodata value


```sql
SELECT ROW(2, '0-100:1-10, 101-500:11-150,501 - 10000: 151-254', '8BUI', 255)::reclassarg;
```


## Example: Reclassify band 1 as an 1BB and no nodata value defined


```sql
SELECT ROW(1, '0-100]:0, (100-255:1', '1BB', NULL)::reclassarg;
```


## See Also


[RT_ST_Reclass](raster-processing-map-algebra.md#RT_ST_Reclass)
  <a id="summarystats"></a>

# summarystats

A composite type returned by the ST_SummaryStats and ST_SummaryStatsAgg functions.

## Description


 A composite type returned by the [RT_ST_SummaryStats](raster-band-statistics-and-analytics.md#RT_ST_SummaryStats) and [RT_ST_SummaryStatsAgg](raster-band-statistics-and-analytics.md#RT_ST_SummaryStatsAgg) functions.

count integer
:   Number of pixels counted for the summary statistics.

sum double precision
:   Sum of all counted pixel values.

mean double precision
:   Arithmetic mean of all counted pixel values.

stddev double precision
:   Standard deviation of all counted pixel values.

min double precision
:   Minimum value of counted pixel values.

max double precision
:   Maximum value of counted pixel values.


## See Also


 [RT_ST_SummaryStats](raster-band-statistics-and-analytics.md#RT_ST_SummaryStats), [RT_ST_SummaryStatsAgg](raster-band-statistics-and-analytics.md#RT_ST_SummaryStatsAgg)
  <a id="unionarg"></a>

# unionarg

A composite type used as input into the ST_Union function defining the bands to be processed and behavior of the UNION operation.

## Description


 A composite type used as input into the ST_Union function defining the bands to be processed and behavior of the UNION operation.

` nband integer `
:   1-based value indicating the band of each input raster to be processed.

` uniontype text `
:   Type of UNION operation. One of defined types as described in [RT_ST_Union](raster-processing-map-algebra.md#RT_ST_Union).


## See Also


 [RT_ST_Union](raster-processing-map-algebra.md#RT_ST_Union)
