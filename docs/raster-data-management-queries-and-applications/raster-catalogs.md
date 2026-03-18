<a id="RT_Raster_Catalog"></a>

## Raster Catalogs


There are two raster catalog views that come packaged with PostGIS. Both views utilize information embedded in the constraints of the raster tables. As a result the catalog views are always consistent with the raster data in the tables since the constraints are enforced.


1. `raster_columns` this view catalogs all the raster table columns in your database.
2. `raster_overviews` this view catalogs all the raster table columns in your database that serve as overviews for a finer grained table. Tables of this type are generated when you use the `-l` switch during load.
 <a id="RT_Raster_Columns"></a>

## Raster Columns Catalog


The `raster_columns` is a catalog of all raster table columns in your database that are of type raster. It is a view utilizing the constraints on the tables so the information is always consistent even if you restore one raster table from a backup of another database. The following columns exist in the `raster_columns` catalog.


If you created your tables not with the loader or forgot to specify the `-C` flag during load, you can enforce the constraints after the fact using [RT_AddRasterConstraints](../raster-reference/raster-management.md#RT_AddRasterConstraints) so that the `raster_columns` catalog registers the common information about your raster tiles.


- `r_table_catalog` The database the table is in. This will always read the current database.
- `r_table_schema` The database schema the raster table belongs to.
- `r_table_name` raster table
- `r_raster_column` the column in the `r_table_name` table that is of type raster. There is nothing in PostGIS preventing you from having multiple raster columns per table so its possible to have a raster table listed multiple times with a different raster column for each.
- `srid` The spatial reference identifier of the raster. Should be an entry in the [Spatial Reference Systems](../data-management/spatial-reference-systems.md#spatial_ref_sys).
- `scale_x` The scaling between geometric spatial coordinates and pixel. This is only available if all tiles in the raster column have the same `scale_x` and this constraint is applied. Refer to [RT_ST_ScaleX](../raster-reference/raster-accessors.md#RT_ST_ScaleX) for more details.
- `scale_y` The scaling between geometric spatial coordinates and pixel. This is only available if all tiles in the raster column have the same `scale_y` and the `scale_y` constraint is applied. Refer to [RT_ST_ScaleY](../raster-reference/raster-accessors.md#RT_ST_ScaleY) for more details.
- `blocksize_x` The width (number of pixels across) of each raster tile . Refer to [RT_ST_Width](../raster-reference/raster-accessors.md#RT_ST_Width) for more details.
- `blocksize_y` The width (number of pixels down) of each raster tile . Refer to [RT_ST_Height](../raster-reference/raster-accessors.md#RT_ST_Height) for more details.
- `same_alignment` A boolean that is true if all the raster tiles have the same alignment . Refer to [RT_ST_SameAlignment](../raster-reference/raster-and-raster-band-spatial-relationships.md#RT_ST_SameAlignment) for more details.
- `regular_blocking` If the raster column has the spatially unique and coverage tile constraints, the value with be TRUE. Otherwise, it will be FALSE.
- `num_bands` The number of bands in each tile of your raster set. This is the same information as what is provided by [RT_ST_NumBands](../raster-reference/raster-accessors.md#RT_ST_NumBands)
- `pixel_types` An array defining the pixel type for each band. You will have the same number of elements in this array as you have number of bands. The pixel_types are one of the following defined in [RT_ST_BandPixelType](../raster-reference/raster-band-accessors.md#RT_ST_BandPixelType).
- `nodata_values` An array of double precision numbers denoting the `nodata_value` for each band. You will have the same number of elements in this array as you have number of bands. These numbers define the pixel value for each band that should be ignored for most operations. This is similar information provided by [RT_ST_BandNoDataValue](../raster-reference/raster-band-accessors.md#RT_ST_BandNoDataValue).
- `out_db` An array of boolean flags indicating if the raster bands data is maintained outside the database. You will have the same number of elements in this array as you have number of bands.
- `extent` This is the extent of all the raster rows in your raster set. If you plan to load more data that will change the extent of the set, you'll want to run the [RT_DropRasterConstraints](../raster-reference/raster-management.md#RT_DropRasterConstraints) function before load and then reapply constraints with [RT_AddRasterConstraints](../raster-reference/raster-management.md#RT_AddRasterConstraints) after load.
- `spatial_index` A boolean that is true if raster column has a spatial index.
  <a id="RT_Raster_Overviews"></a>

## Raster Overviews


`raster_overviews` catalogs information about raster table columns used for overviews and additional information about them that is useful to know when utilizing overviews. Overview tables are cataloged in both `raster_columns` and `raster_overviews` because they are rasters in their own right but also serve an additional special purpose of being a lower resolution caricature of a higher resolution table. These are generated along-side the main raster table when you use the `-l` switch in raster loading or can be generated manually using [RT_AddOverviewConstraints](../raster-reference/raster-management.md#RT_AddOverviewConstraints).


Overview tables contain the same constraints as other raster tables as well as additional informational only constraints specific to overviews.


!!! note

    The information in `raster_overviews` does not duplicate the information in `raster_columns`. If you need the information about an overview table present in `raster_columns` you can join the `raster_overviews` and `raster_columns` together to get the full set of information you need.


Two main reasons for overviews are:


1. Low resolution representation of the core tables commonly used for fast mapping zoom-out.
2. Computations are generally faster to do on them than their higher resolution parents because there are fewer records and each pixel covers more territory. Though the computations are not as accurate as the high-res tables they support, they can be sufficient in many rule-of-thumb computations.


The `raster_overviews` catalog contains the following columns of information.


- `o_table_catalog` The database the overview table is in. This will always read the current database.
- `o_table_schema` The database schema the overview raster table belongs to.
- `o_table_name` raster overview table name
- `o_raster_column` the raster column in the overview table.
- `r_table_catalog` The database the raster table that this overview services is in. This will always read the current database.
- `r_table_schema` The database schema the raster table that this overview services belongs to.
- `r_table_name` raster table that this overview services.
- `r_raster_column` the raster column that this overview column services.
- `overview_factor` - this is the pyramid level of the overview table. The higher the number the lower the resolution of the table. raster2pgsql if given a folder of images, will compute overview of each image file and load separately. Level 1 is assumed and always the original file. Level 2 is will have each tile represent 4 of the original. So for example if you have a folder of 5000x5000 pixel image files that you chose to chunk 125x125, for each image file your base table will have (5000*5000)/(125*125) records = 1600, your (l=2) `o_2` table will have ceiling(1600/Power(2,2)) = 400 rows, your (l=3) `o_3` will have ceiling(1600/Power(2,3) ) = 200 rows. If your pixels aren't divisible by the size of your tiles, you'll get some scrap tiles (tiles not completely filled). Note that each overview tile generated by raster2pgsql has the same number of pixels as its parent, but is of a lower resolution where each pixel of it represents (Power(2,overview_factor) pixels of the original).
