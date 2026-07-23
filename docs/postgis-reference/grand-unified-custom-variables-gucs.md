<a id="PostGIS_GUC"></a>

## Grand Unified Custom Variables (GUCs)
  <a id="postgis_gdal_datapath"></a>

# postgis.gdal_datapath

A configuration option to assign the value of GDAL's GDAL_DATA option. If not set, the environmentally set GDAL_DATA variable is used.

## Description


 A PostgreSQL GUC variable for setting the value of GDAL's GDAL_DATA option. The `postgis.gdal_datapath` value should be the complete physical path to GDAL's data files.


 This configuration option is of most use for Windows platforms where GDAL's data files path is not hard-coded. This option should also be set when GDAL's data files are not located in GDAL's expected path.


!!! note

    This option can be set in PostgreSQL's configuration file postgresql.conf. It can also be set by connection or transaction.


Availability: 2.2.0


!!! note

    Additional information about GDAL_DATA is available at GDAL's [Configuration Options](http://trac.osgeo.org/gdal/wiki/ConfigOptions).


## Examples


Set and reset `postgis.gdal_datapath`


```sql

SET postgis.gdal_datapath TO '/usr/local/share/gdal.hidden';
SET postgis.gdal_datapath TO default;

```


Setting on windows for a particular database


```sql
ALTER DATABASE gisdb
SET postgis.gdal_datapath = 'C:/Program Files/PostgreSQL/9.3/gdal-data';
```


## See Also


 [RT_PostGIS_GDAL_Version](../raster-reference/raster-management.md#RT_PostGIS_GDAL_Version), [RT_ST_Transform](../raster-reference/raster-editors.md#RT_ST_Transform)
  <a id="postgis_gdal_enabled_drivers"></a>

# postgis.gdal_enabled_drivers

A configuration option to set the enabled GDAL drivers in the PostGIS environment. Affects the GDAL configuration variable GDAL_SKIP.

## Description


 A configuration option to set the enabled GDAL drivers in the PostGIS environment. Affects the GDAL configuration variable GDAL_SKIP. This option can be set in PostgreSQL's configuration file: postgresql.conf. It can also be set by connection or transaction.


 The initial value of `postgis.gdal_enabled_drivers` may also be set by passing the environment variable `POSTGIS_GDAL_ENABLED_DRIVERS` with the list of enabled drivers to the process starting PostgreSQL.


 Enabled GDAL specified drivers can be specified by the driver's short-name or code. Driver short-names or codes can be found at [GDAL Raster Formats](http://www.gdal.org/formats_list.html). Multiple drivers can be specified by putting a space between each driver.


!!! note

    There are three special codes available for `postgis.gdal_enabled_drivers`. The codes are case-sensitive.

    - `DISABLE_ALL` disables all GDAL drivers. If present, `DISABLE_ALL` overrides all other values in `postgis.gdal_enabled_drivers`.
    - `ENABLE_ALL` enables all GDAL drivers.
    - `VSICURL` enables GDAL's `/vsicurl/` virtual file system.


     When `postgis.gdal_enabled_drivers` is set to DISABLE_ALL, attempts to use out-db rasters, ST_FromGDALRaster(), ST_AsGDALRaster(), ST_AsTIFF(), ST_AsJPEG() and ST_AsPNG() will result in error messages.


!!! note

    In the standard PostGIS installation, `postgis.gdal_enabled_drivers` is set to DISABLE_ALL.


!!! note

    Additional information about GDAL_SKIP is available at GDAL's [Configuration Options](http://trac.osgeo.org/gdal/wiki/ConfigOptions).


Availability: 2.2.0


## Examples


To set and reset `postgis.gdal_enabled_drivers` for current session


```sql

SET postgis.gdal_enabled_drivers = 'ENABLE_ALL';
SET postgis.gdal_enabled_drivers = default;

```


Set for all new connections to a specific database to specific drivers


```sql
ALTER DATABASE mygisdb SET postgis.gdal_enabled_drivers TO 'GTiff PNG JPEG';
```


Setting for whole database cluster to enable all drivers. Requires super user access. Also note that database, session, and user settings override this.


```

 --writes to postgres.auto.conf
ALTER SYSTEM SET postgis.gdal_enabled_drivers TO 'ENABLE_ALL';
 --Reloads postgres conf
SELECT pg_reload_conf();

```


## See Also


 [RT_ST_FromGDALRaster](../raster-reference/raster-constructors.md#RT_ST_FromGDALRaster), [RT_ST_AsGDALRaster](../raster-reference/raster-outputs.md#RT_ST_AsGDALRaster), [RT_ST_AsTIFF](../raster-reference/raster-outputs.md#RT_ST_AsTIFF), [RT_ST_AsPNG](../raster-reference/raster-outputs.md#RT_ST_AsPNG), [RT_ST_AsJPEG](../raster-reference/raster-outputs.md#RT_ST_AsJPEG), [postgis_enable_outdb_rasters](#postgis_enable_outdb_rasters)
  <a id="postgis_enable_outdb_rasters"></a>

# postgis.enable_outdb_rasters

A boolean configuration option to enable access to out-db raster bands.

## Description


 A boolean configuration option to enable access to out-db raster bands. This option can be set in PostgreSQL's configuration file: postgresql.conf. It can also be set by connection or transaction.


 The initial value of `postgis.enable_outdb_rasters` may also be set by passing the environment variable `POSTGIS_ENABLE_OUTDB_RASTERS` with a non-zero value to the process starting PostgreSQL.


!!! note

    Even if `postgis.enable_outdb_rasters` is True, the GUC `postgis.gdal_enabled_drivers` determines the accessible raster formats.


!!! note

    In the standard PostGIS installation, `postgis.enable_outdb_rasters` is set to False.


Availability: 2.2.0


## Examples


Set and reset `postgis.enable_outdb_rasters` for current session


```sql

SET postgis.enable_outdb_rasters TO True;
SET postgis.enable_outdb_rasters = default;
SET postgis.enable_outdb_rasters = True;
SET postgis.enable_outdb_rasters = False;

```


Set for all new connections to a specific database


```sql

ALTER DATABASE gisdb SET postgis.enable_outdb_rasters = true;

```


Setting for whole database cluster. Requires super user access. Also note that database, session, and user settings override this.


```

 --writes to postgres.auto.conf
ALTER SYSTEM SET postgis.enable_outdb_rasters = true;
 --Reloads postgres conf
SELECT pg_reload_conf();

```


## See Also


 [postgis_gdal_enabled_drivers](#postgis_gdal_enabled_drivers) [postgis_gdal_vsi_options](#postgis_gdal_vsi_options)
  <a id="postgis_gdal_vsi_options"></a>

# postgis.gdal_vsi_options

A string configuration to set options used when working with an out-db raster.

## Description


 A string configuration to set options used when working with an out-db raster. [Configuration options](http://trac.osgeo.org/gdal/wiki/ConfigOptions) control things like how much space GDAL allocates to local data cache, whether to read overviews, and what access keys to use for remote out-db data sources.


Availability: 3.2.0


## Examples


Set `postgis.gdal_vsi_options` for current session:


```sql

SET postgis.gdal_vsi_options = 'AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxx AWS_SECRET_ACCESS_KEY=yyyyyyyyyyyyyyyyyyyyyyyyyy';

```


Set `postgis.gdal_vsi_options` just for the *current transaction* using the `LOCAL` keyword:


```sql

SET LOCAL postgis.gdal_vsi_options = 'AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxx AWS_SECRET_ACCESS_KEY=yyyyyyyyyyyyyyyyyyyyyyyyyy';

```


## See Also


 [postgis_enable_outdb_rasters](#postgis_enable_outdb_rasters) [postgis_gdal_enabled_drivers](#postgis_gdal_enabled_drivers)
