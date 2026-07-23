<a id="Version_Functions"></a>

## Version Functions
  <a id="PostGIS_Extensions_Upgrade"></a>

# PostGIS_Extensions_Upgrade

Packages and upgrades PostGIS extensions (e.g. postgis_raster, postgis_topology, postgis_sfcgal) to given or latest version.

## Synopsis


```sql
text PostGIS_Extensions_Upgrade(text target_version=null)
```


## Description


Packages and upgrades PostGIS extensions to given or latest version. Only extensions you have installed in the database will be packaged and upgraded if needed. Reports full PostGIS version and build configuration infos after. This is short-hand for doing multiple CREATE EXTENSION .. FROM unpackaged and ALTER EXTENSION .. UPDATE for each PostGIS extension. Currently only tries to upgrade extensions postgis, postgis_raster, postgis_sfcgal, postgis_topology, and postgis_tiger_geocoder.


Availability: 2.5.0


!!! note

    Changed: 3.4.0 to add target_version argument.


    Changed: 3.3.0 support for upgrades from any PostGIS version. Does not work on all systems.


    Changed: 3.0.0 to repackage loose extensions and support postgis_raster.


## Examples


```sql
SELECT PostGIS_Extensions_Upgrade();
```


```

NOTICE:  Packaging extension postgis
NOTICE:  Packaging extension postgis_raster
NOTICE:  Packaging extension postgis_sfcgal
NOTICE:  Extension postgis_topology is not available or not packagable for some reason
NOTICE:  Extension postgis_tiger_geocoder is not available or not packagable for some reason

                    postgis_extensions_upgrade
-------------------------------------------------------------------
 Upgrade completed, run SELECT postgis_full_version(); for details
(1 row)
```


## See Also


 [Upgrading spatial databases](../postgis-administration/upgrading-spatial-databases.md#upgrading), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_Full_Version"></a>

# PostGIS_Full_Version

Reports full PostGIS version and build configuration infos.

## Synopsis


```sql
text PostGIS_Full_Version()
```


## Description


Reports full PostGIS version and build configuration infos. Also informs about synchronization between libraries and scripts suggesting upgrades as needed.


Enhanced: 3.4.0 now includes extra PROJ configurations NETWORK_ENABLED, URL_ENDPOINT and DATABASE_PATH of proj.db location


## Examples


```sql
SELECT PostGIS_Full_Version();
							   postgis_full_version
----------------------------------------------------------------------------------
POSTGIS="3.4.0dev 3.3.0rc2-993-g61bdf43a7" [EXTENSION] PGSQL="160" GEOS="3.12.0dev-CAPI-1.18.0" SFCGAL="1.3.8" PROJ="7.2.1 NETWORK_ENABLED=OFF URL_ENDPOINT=https://cdn.proj.org USER_WRITABLE_DIRECTORY=/tmp/proj DATABASE_PATH=/usr/share/proj/proj.db" GDAL="GDAL 3.2.2, released 2021/03/05" LIBXML="2.9.10" LIBJSON="0.15" LIBPROTOBUF="1.3.3" WAGYU="0.5.0 (Internal)" TOPOLOGY RASTER
(1 row)
```


## See Also


 [Upgrading spatial databases](../postgis-administration/upgrading-spatial-databases.md#upgrading), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Wagyu_Version](#PostGIS_Wagyu_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_GEOS_Version"></a>

# PostGIS_GEOS_Version

Returns the version number of the GEOS library.

## Synopsis


```sql
text PostGIS_GEOS_Version()
```


## Description


Returns the version number of the GEOS library, or `NULL` if GEOS support is not enabled.


## Examples


```sql
SELECT PostGIS_GEOS_Version();
 postgis_geos_version
----------------------
3.12.0dev-CAPI-1.18.0
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_GEOS_Compiled_Version"></a>

# PostGIS_GEOS_Compiled_Version

Returns the version number of the GEOS library against which PostGIS was built.

## Synopsis


```sql
text PostGIS_GEOS_Compiled_Version()
```


## Description


Returns the version number of the GEOS library, or against which PostGIS was built.


Availability: 3.4.0


## Examples


```sql
SELECT PostGIS_GEOS_Compiled_Version();
 postgis_geos_compiled_version
-------------------------------
 3.12.0
(1 row)
```


## See Also


 [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Full_Version](#PostGIS_Full_Version)
  <a id="PostGIS_Liblwgeom_Version"></a>

# PostGIS_Liblwgeom_Version

Returns the version number of the liblwgeom library. This should match the version of PostGIS.

## Synopsis


```sql
text PostGIS_Liblwgeom_Version()
```


## Description


Returns the version number of the liblwgeom library/


## Examples


```sql
SELECT PostGIS_Liblwgeom_Version();
postgis_liblwgeom_version
--------------------------
3.4.0dev 3.3.0rc2-993-g61bdf43a7
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_LibXML_Version"></a>

# PostGIS_LibXML_Version

Returns the version number of the libxml2 library.

## Synopsis


```sql
text PostGIS_LibXML_Version()
```


## Description


Returns the version number of the LibXML2 library.


Availability: 1.5


## Examples


```sql
SELECT PostGIS_LibXML_Version();
 postgis_libxml_version
----------------------
 2.9.10
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_LibJSON_Version"></a>

# PostGIS_LibJSON_Version

Returns the version number of the libjson-c library.

## Synopsis


```sql
text PostGIS_LibJSON_Version()
```


## Description


Returns the version number of the LibJSON-C library.


## Examples


```sql
SELECT PostGIS_LibJSON_Version();
 postgis_libjson_version
-------------------------
 0.17
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_Lib_Build_Date"></a>

# PostGIS_Lib_Build_Date

Returns build date of the PostGIS library.

## Synopsis


```sql
text PostGIS_Lib_Build_Date()
```


## Description


Returns build date of the PostGIS library.


## Examples


```sql
SELECT PostGIS_Lib_Build_Date();
 postgis_lib_build_date
------------------------
 2023-06-22 03:56:11
(1 row)
```
  <a id="PostGIS_Lib_Version"></a>

# PostGIS_Lib_Version

Returns the version number of the PostGIS library.

## Synopsis


```sql
text PostGIS_Lib_Version()
```


## Description


Returns the version number of the PostGIS library.


## Examples


```sql
SELECT PostGIS_Lib_Version();
 postgis_lib_version
---------------------
 3.4.0dev
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_PROJ_Version"></a>

# PostGIS_PROJ_Version

Returns the version number of the PROJ4 library.

## Synopsis


```sql
text PostGIS_PROJ_Version()
```


## Description


Returns the version number of the PROJ library and some configuration options of proj.


Enhanced: 3.4.0 now includes NETWORK_ENABLED, URL_ENDPOINT and DATABASE_PATH of proj.db location


## Examples


```sql
SELECT PostGIS_PROJ_Version();
  postgis_proj_version
-------------------------
7.2.1 NETWORK_ENABLED=OFF URL_ENDPOINT=https://cdn.proj.org USER_WRITABLE_DIRECTORY=/tmp/proj DATABASE_PATH=/usr/share/proj/proj.db
(1 row)
```


## See Also


 [PostGIS_PROJ_Compiled_Version](#PostGIS_PROJ_Compiled_Version), [PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_PROJ_Compiled_Version"></a>

# PostGIS_PROJ_Compiled_Version

Returns the version number of the PROJ library against which PostGIS was built.

## Synopsis


```sql
text PostGIS_PROJ_Compiled_Version()
```


## Description


Returns the version number of the PROJ library, or against which PostGIS was built.


Availability: 3.5.0


## Examples


```sql
SELECT PostGIS_PROJ_Compiled_Version();
 postgis_proj_compiled_version
-------------------------------
 9.1.1
(1 row)
```


## See Also


 [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Full_Version](#PostGIS_Full_Version)
  <a id="PostGIS_Wagyu_Version"></a>

# PostGIS_Wagyu_Version

Returns the version number of the internal Wagyu library.

## Synopsis


```sql
text PostGIS_Wagyu_Version()
```


## Description


Returns the version number of the internal Wagyu library, or `NULL` if Wagyu support is not enabled.


## Examples


```sql
SELECT PostGIS_Wagyu_Version();
 postgis_wagyu_version
-----------------------
 0.5.0 (Internal)
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_Scripts_Build_Date"></a>

# PostGIS_Scripts_Build_Date

Returns build date of the PostGIS scripts.

## Synopsis


```sql
text PostGIS_Scripts_Build_Date()
```


## Description


Returns build date of the PostGIS scripts.


Availability: 1.0.0RC1


## Examples


```sql
SELECT PostGIS_Scripts_Build_Date();
  postgis_scripts_build_date
-------------------------
 2023-06-22 03:56:11
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_Scripts_Installed"></a>

# PostGIS_Scripts_Installed

Returns version of the PostGIS scripts installed in this database.

## Synopsis


```sql
text PostGIS_Scripts_Installed()
```


## Description


Returns version of the PostGIS scripts installed in this database.


!!! note

    If the output of this function doesn't match the output of [PostGIS_Scripts_Released](#PostGIS_Scripts_Released) you probably missed to properly upgrade an existing database. See the [Upgrading](../postgis-administration/upgrading-spatial-databases.md#upgrading) section for more info.


Availability: 0.9.0


## Examples


```sql
SELECT PostGIS_Scripts_Installed();
  postgis_scripts_installed
-------------------------
 3.4.0dev 3.3.0rc2-993-g61bdf43a7
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_Scripts_Released](#PostGIS_Scripts_Released), [PostGIS_Version](#PostGIS_Version)
  <a id="PostGIS_Scripts_Released"></a>

# PostGIS_Scripts_Released

Returns the version number of the postgis.sql script released with the installed PostGIS lib.

## Synopsis


```sql
text PostGIS_Scripts_Released()
```


## Description


Returns the version number of the postgis.sql script released with the installed PostGIS lib.


!!! note

    Starting with version 1.1.0 this function returns the same value of [PostGIS_Lib_Version](#PostGIS_Lib_Version). Kept for backward compatibility.


Availability: 0.9.0


## Examples


```sql
SELECT PostGIS_Scripts_Released();
  postgis_scripts_released
-------------------------
 3.4.0dev 3.3.0rc2-993-g61bdf43a7
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_Scripts_Installed](#PostGIS_Scripts_Installed), [PostGIS_Lib_Version](#PostGIS_Lib_Version)
  <a id="PostGIS_Version"></a>

# PostGIS_Version

Returns PostGIS version number and compile-time options.

## Synopsis


```sql
text PostGIS_Version()
```


## Description


Returns PostGIS version number and compile-time options.


## Examples


```sql
SELECT PostGIS_Version();
			postgis_version
---------------------------------------
 3.4 USE_GEOS=1 USE_PROJ=1 USE_STATS=1
(1 row)
```


## See Also


[PostGIS_Full_Version](#PostGIS_Full_Version), [PostGIS_GEOS_Version](#PostGIS_GEOS_Version), [PostGIS_Lib_Version](#PostGIS_Lib_Version), [PostGIS_LibXML_Version](#PostGIS_LibXML_Version), [PostGIS_PROJ_Version](#PostGIS_PROJ_Version)
