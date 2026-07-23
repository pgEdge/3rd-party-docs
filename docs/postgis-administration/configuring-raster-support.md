<a id="raster_configuration"></a>

## Configuring raster support


 If you enabled raster support you may want to read below how to properly configure it.


As of PostGIS 2.1.3, out-of-db rasters and all raster drivers are disabled by default. In order to re-enable these, you need to set the following environment variables `POSTGIS_GDAL_ENABLED_DRIVERS` and `POSTGIS_ENABLE_OUTDB_RASTERS` in the server environment. For PostGIS 2.2, you can use the more cross-platform approach of setting the corresponding [Grand Unified Custom Variables (GUCs)](../postgis-reference/grand-unified-custom-variables-gucs.md#PostGIS_GUC).


If you want to enable offline raster:


```
POSTGIS_ENABLE_OUTDB_RASTERS=1
```


Any other setting or no setting at all will disable out of db rasters.


In order to enable all GDAL drivers available in your GDAL install, set this environment variable as follows


```
POSTGIS_GDAL_ENABLED_DRIVERS=ENABLE_ALL
```


If you want to only enable specific drivers, set your environment variable as follows:


```
POSTGIS_GDAL_ENABLED_DRIVERS="GTiff PNG JPEG GIF XYZ"
```


!!! note

    If you are on windows, do not quote the driver list


Setting environment variables varies depending on OS. For PostgreSQL installed on Ubuntu or Debian via apt-postgresql, the preferred way is to edit <code>/etc/postgresql/</code><em>10</em><code>/</code><em>main</em><code>/environment</code> where 10 refers to version of PostgreSQL and main refers to the cluster.


On windows, if you are running as a service, you can set via System variables which for Windows 7 you can get to by right-clicking on Computer->Properties Advanced System Settings or in explorer navigating to `Control Panel\All Control Panel Items\System`. Then clicking *Advanced System Settings ->Advanced->Environment Variables* and adding new system variables.


After you set the environment variables, you'll need to restart your PostgreSQL service for the changes to take effect.
