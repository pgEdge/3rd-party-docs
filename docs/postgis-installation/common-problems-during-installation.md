## Common Problems during installation


 There are several things to check when your installation or upgrade doesn't go as you expected.


1.  Check that you have installed PostgreSQL 12 or newer, and that you are compiling against the same version of the PostgreSQL source as the version of PostgreSQL that is running. Mix-ups can occur when your (Linux) distribution has already installed PostgreSQL, or you have otherwise installed PostgreSQL before and forgotten about it. PostGIS will only work with PostgreSQL 12 or newer, and strange, unexpected error messages will result if you use an older version. To check the version of PostgreSQL which is running, connect to the database using psql and run this query:

```sql
SELECT version();
```

    If you are running an RPM based distribution, you can check for the existence of pre-installed packages using the `rpm` command as follows: `rpm -qa | grep postgresql`
2. If your upgrade fails, make sure you are restoring into a database that already has PostGIS installed.

```sql
SELECT postgis_full_version();
```


 Also check that configure has correctly detected the location and version of PostgreSQL, the Proj library and the GEOS library.


1.  The output from configure is used to generate the `postgis_config.h` file. Check that the `POSTGIS_PGSQL_VERSION`, `POSTGIS_PROJ_VERSION` and `POSTGIS_GEOS_VERSION` variables have been set correctly.
