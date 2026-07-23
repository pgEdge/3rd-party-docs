## PostGIS 3.5.5


2026/02/09


This version requires PostgreSQL 12-18, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.12+ is needed. To take advantage of all SFCGAL features, SFCGAL 1.5.0+ is needed.


## Bug Fixes


[5959](https://trac.osgeo.org/postgis/ticket/5959), #5984, Prevent histogram target overflow when analysing massive tables (Darafei Praliaskouski)


[6020](https://trac.osgeo.org/postgis/ticket/6020), schema qualify call in ST_MPointFromText (Paul Ramsey)


[6023](https://trac.osgeo.org/postgis/ticket/6023), Fix robustness issue in ptarray_contains_point (Sandro Santilli)


[6027](https://trac.osgeo.org/postgis/ticket/6027), Fix RemoveUnusedPrimitives without topology in search_path (Sandro Santilli)


[6028](https://trac.osgeo.org/postgis/ticket/6028), crash indexing malformed empty polygon (Paul Ramsey)


[GH-841](https://github.com/postgis/postgis/pull/841), small memory leak in address_standardizer (Maxim Korotkov)


[5998](https://trac.osgeo.org/postgis/ticket/5998), [tiger_geocoder] [security] CVE-2022-2625, make sure tables requires by extension are owned by extension


[5853](https://trac.osgeo.org/postgis/ticket/5853), Issue with topology and tiger geocoder upgrade scripts (Regina Obe, Spencer Bryson)
