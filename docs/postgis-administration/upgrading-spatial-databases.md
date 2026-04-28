<a id="upgrading"></a>

## Upgrading spatial databases


 Upgrading existing spatial databases can be tricky as it requires replacement or introduction of new PostGIS object definitions.


 Unfortunately not all definitions can be easily replaced in a live database, so sometimes your best bet is a dump/reload process.


 PostGIS provides a SOFT UPGRADE procedure for minor or bugfix releases, and a HARD UPGRADE procedure for major releases.


 Before attempting to upgrade PostGIS, it is always worth to backup your data. If you use the -Fc flag to pg_dump you will always be able to restore the dump with a HARD UPGRADE.
 <a id="soft_upgrade"></a>

## Soft upgrade


If you installed your database using extensions, you'll need to upgrade using the extension model as well. If you installed using the old sql script way, you are advised to switch your install to extensions because the script way is no longer supported.
 <a id="soft_upgrade_extensions"></a>

## Soft Upgrade 9.1+ using extensions


If you originally installed PostGIS with extensions, then you need to upgrade using extensions as well. Doing a minor upgrade with extensions, is fairly painless.


If you are running PostGIS 3 or above, then you should use the [PostGIS_Extensions_Upgrade](../postgis-reference/version-functions.md#PostGIS_Extensions_Upgrade) function to upgrade to the latest version you have installed.


```sql
SELECT postgis_extensions_upgrade();
```


If you are running PostGIS 2.5 or lower, then do the following:


```sql
ALTER EXTENSION postgis UPDATE;
SELECT postgis_extensions_upgrade();
-- This second call is needed to rebundle postgis_raster extension
SELECT postgis_extensions_upgrade();
```


If you have multiple versions of PostGIS installed, and you don't want to upgrade to the latest, you can explicitly specify the version as follows:


```sql
ALTER EXTENSION postgis UPDATE TO "3.5.6";
ALTER EXTENSION postgis_topology UPDATE TO "3.5.6";
```


If you get an error notice something like:


```
No migration path defined for … to 3.5.6
```


Then you'll need to backup your database, create a fresh one as described in [Spatially enable database using EXTENSION](creating-spatial-databases.md#create_new_db_extensions) and then restore your backup on top of this new database.


If you get a notice message like:


```
Version "3.5.6" of extension "postgis" is already installed
```


 Then everything is already up to date and you can safely ignore it. **UNLESS** you're attempting to upgrade from an development version to the next (which doesn't get a new version number); in that case you can append "next" to the version string, and next time you'll need to drop the "next" suffix again:


```sql
ALTER EXTENSION postgis UPDATE TO "3.5.6next";
ALTER EXTENSION postgis_topology UPDATE TO "3.5.6next";
```


!!! note

    If you installed PostGIS originally without a version specified, you can often skip the reinstallation of postgis extension before restoring since the backup just has <code>CREATE EXTENSION postgis</code> and thus picks up the newest latest version during restore.


!!! note

    If you are upgrading PostGIS extension from a version prior to 3.0.0, you will have a new extension *postgis_raster* which you can safely drop, if you don't need raster support. You can drop as follows:


    ```sql
    DROP EXTENSION postgis_raster;
    ```
  <a id="soft_upgrade_sql_script"></a>

## Soft Upgrade Pre 9.1+ or without extensions


This section applies only to those who installed PostGIS not using extensions. If you have extensions and try to upgrade with this approach you'll get messages like:


```
can't drop … because postgis extension depends on it
```


NOTE: if you are moving from PostGIS 1.* to PostGIS 2.* or from PostGIS 2.* prior to r7409, you cannot use this procedure but would rather need to do a [HARD UPGRADE](#hard_upgrade).


 After compiling and installing (make install) you should find a set of `*_upgrade.sql` files in the installation folders. You can list them all with:


```
ls `pg_config --sharedir`/contrib/postgis-3.5.6/*_upgrade.sql
```


 Load them all in turn, starting from `postgis_upgrade.sql`.


```
psql -f postgis_upgrade.sql -d your_spatial_database
```


 The same procedure applies to raster, topology and sfcgal extensions, with upgrade files named `rtpostgis_upgrade.sql`, `topology_upgrade.sql` and `sfcgal_upgrade.sql` respectively. If you need them:


```
psql -f rtpostgis_upgrade.sql -d your_spatial_database
```


```
psql -f topology_upgrade.sql -d your_spatial_database
```


```
psql -f sfcgal_upgrade.sql -d your_spatial_database
```


You are advised to switch to an extension based install by running


```
psql -c "SELECT postgis_extensions_upgrade();"
```


!!! note

    If you can't find the `postgis_upgrade.sql` specific for upgrading your version you are using a version too early for a soft upgrade and need to do a [HARD UPGRADE](#hard_upgrade).


 The [PostGIS_Full_Version](../postgis-reference/version-functions.md#PostGIS_Full_Version) function should inform you about the need to run this kind of upgrade using a "procs need upgrade" message.
   <a id="hard_upgrade"></a>

## Hard upgrade


 By HARD UPGRADE we mean full dump/reload of postgis-enabled databases. You need a HARD UPGRADE when PostGIS objects' internal storage changes or when SOFT UPGRADE is not possible. The [Release Notes](../appendix/index.md#release_notes) appendix reports for each version whether you need a dump/reload (HARD UPGRADE) to upgrade.


 The dump/reload process is assisted by the postgis_restore script which takes care of skipping from the dump all definitions which belong to PostGIS (including old ones), allowing you to restore your schemas and data into a database with PostGIS installed without getting duplicate symbol errors or bringing forward deprecated objects.


Supplementary instructions for windows users are available at [Windows Hard upgrade](http://trac.osgeo.org/postgis/wiki/UsersWikiWinUpgrade).


 The Procedure is as follows:


1.  Create a "custom-format" dump of the database you want to upgrade (let's call it `olddb`) include binary blobs (-b) and verbose (-v) output. The user can be the owner of the db, need not be postgres super account.

```
pg_dump -h localhost -p 5432 -U postgres -Fc -b -v -f "/somepath/olddb.backup" olddb
```
2.  Do a fresh install of PostGIS in a new database -- we'll refer to this database as `newdb`. Please refer to [Spatially enable database without using EXTENSION (discouraged)](creating-spatial-databases.md#create_new_db) and [Spatially enable database using EXTENSION](creating-spatial-databases.md#create_new_db_extensions) for instructions on how to do this.

    The spatial_ref_sys entries found in your dump will be restored, but they will not override existing ones in spatial_ref_sys. This is to ensure that fixes in the official set will be properly propagated to restored databases. If for any reason you really want your own overrides of standard entries just don't load the spatial_ref_sys.sql file when creating the new db.

    If your database is really old or you know you've been using long deprecated functions in your views and functions, you might need to load `legacy.sql` for all your functions and views etc. to properly come back. Only do this if _really_ needed. Consider upgrading your views and functions before dumping instead, if possible. The deprecated functions can be later removed by loading `uninstall_legacy.sql`.
3.  Restore your backup into your fresh `newdb` database using postgis_restore. Unexpected errors, if any, will be printed to the standard error stream by psql. Keep a log of those.

```
postgis_restore "/somepath/olddb.backup" | psql -h localhost -p 5432 -U postgres newdb 2> errors.txt
```


 Errors may arise in the following cases:


1.  Some of your views or functions make use of deprecated PostGIS objects. In order to fix this you may try loading `legacy.sql` script prior to restore or you'll have to restore to a version of PostGIS which still contains those objects and try a migration again after porting your code. If the `legacy.sql` way works for you, don't forget to fix your code to stop using deprecated functions and drop them loading `uninstall_legacy.sql`.
2.  Some custom records of spatial_ref_sys in dump file have an invalid SRID value. Valid SRID values are bigger than 0 and smaller than 999000. Values in the 999000.999999 range are reserved for internal use while values > 999999 can't be used at all. All your custom records with invalid SRIDs will be retained, with those > 999999 moved into the reserved range, but the spatial_ref_sys table would lose a check constraint guarding for that invariant to hold and possibly also its primary key ( when multiple invalid SRIDS get converted to the same reserved SRID value ).

    In order to fix this you should copy your custom SRS to a SRID with a valid value (maybe in the 910000..910999 range), convert all your tables to the new srid (see [UpdateGeometrySRID](../postgis-reference/table-management-functions.md#UpdateGeometrySRID)), delete the invalid entry from spatial_ref_sys and re-construct the check(s) with:

```sql

ALTER TABLE spatial_ref_sys ADD CONSTRAINT spatial_ref_sys_srid_check check (srid > 0 AND srid < 999000 );
```

```sql

ALTER TABLE spatial_ref_sys ADD PRIMARY KEY(srid));

```

    If you are upgrading an old database containing french [IGN](https://en.wikipedia.org/wiki/Institut_g%C3%A9ographique_national) cartography, you will have probably SRIDs out of range and you will see, when importing your database, issues like this :

```
 WARNING: SRID 310642222 converted to 999175 (in reserved zone)
```

    In this case, you can try following steps : first throw out completely the IGN from the sql which is resulting from postgis_restore. So, after having run :

```
postgis_restore "/somepath/olddb.backup" > olddb.sql
```

    run this command :

```
grep -v IGNF olddb.sql > olddb-without-IGN.sql
```

    Create then your newdb, activate the required Postgis extensions, and insert properly the french system IGN with : [this script](https://raw.githubusercontent.com/Remi-C/IGN_spatial_ref_for_PostGIS/master/Put_IGN_SRS_into_Postgis.sql) After these operations, import your data :

```
psql -h localhost -p 5432 -U postgres -d newdb -f olddb-without-IGN.sql  2> errors.txt
```
