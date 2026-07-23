<a id="RT_reference"></a>

# Raster Reference

The functions given below are the ones which a user of PostGIS Raster is likely to need and which are currently available in PostGIS Raster. There are other functions which are required support functions to the raster objects which are not of use to a general user.

`raster` is a new PostGIS type for storing and analyzing raster data.

For loading rasters from raster files please refer to [Loading and Creating Rasters](../raster-data-management-queries-and-applications/loading-and-creating-rasters.md#RT_Loading_Rasters)

For the examples in this reference we will be using a raster table of dummy rasters - Formed with the following code

```sql
CREATE TABLE dummy_rast(rid integer, rast raster);
INSERT INTO dummy_rast(rid, rast)
VALUES (1,
('01' -- little endian (uint8 ndr)
||
'0000' -- version (uint16 0)
||
'0000' -- nBands (uint16 0)
||
'0000000000000040' -- scaleX (float64 2)
||
'0000000000000840' -- scaleY (float64 3)
||
'000000000000E03F' -- ipX (float64 0.5)
||
'000000000000E03F' -- ipY (float64 0.5)
||
'0000000000000000' -- skewX (float64 0)
||
'0000000000000000' -- skewY (float64 0)
||
'00000000' -- SRID (int32 0)
||
'0A00' -- width (uint16 10)
||
'1400' -- height (uint16 20)
)::raster
),
-- Raster: 5 x 5 pixels, 3 bands, PT_8BUI pixel type, NODATA = 0
(2,  ('01000003009A9999999999A93F9A9999999999A9BF000000E02B274A' ||
'41000000007719564100000000000000000000000000000000FFFFFFFF050005000400FDFEFDFEFEFDFEFEFDF9FAFEF' ||
'EFCF9FBFDFEFEFDFCFAFEFEFE04004E627AADD16076B4F9FE6370A9F5FE59637AB0E54F58617087040046566487A1506CA2E3FA5A6CAFFBFE4D566DA4CB3E454C5665')::raster);
```

- [Raster Support Data types](raster-support-data-types.md#Raster_Types)
- [Raster Management](raster-management.md#Raster_Management_Functions)
- [Raster Constructors](raster-constructors.md#Raster_Constructors)
- [Raster Accessors](raster-accessors.md#Raster_Accessors)
- [Raster Band Accessors](raster-band-accessors.md#RasterBand_Accessors)
- [Raster Pixel Accessors and Setters](raster-pixel-accessors-and-setters.md#Raster_Pixel_Accessors)
- [Raster Editors](raster-editors.md#Raster_Editors)
- [Raster Band Editors](raster-band-editors.md#RasterBand_Editors)
- [Raster Band Statistics and Analytics](raster-band-statistics-and-analytics.md#RasterBand_Stats)
- [Raster Inputs](raster-inputs.md#Raster_Inputs)
- [Raster Outputs](raster-outputs.md#Raster_Outputs)
- [Raster Processing: Map Algebra](raster-processing-map-algebra.md#Raster_Processing_MapAlgebra)
- [Built-in Map Algebra Callback Functions](built-in-map-algebra-callback-functions.md#Raster_Processing_MapAlgebra_Callbacks)
- [Raster Processing: DEM (Elevation)](raster-processing-dem-elevation.md#Raster_Processing_DEM)
- [Raster Processing: Raster to Geometry](raster-processing-raster-to-geometry.md#Raster_Processing_Geometry)
- [Raster Operators](raster-operators.md#RT_Operators)
- [Raster and Raster Band Spatial Relationships](raster-and-raster-band-spatial-relationships.md#Raster_Relationships)
- [Raster Tips](raster-tips.md#Raster_Tips)
