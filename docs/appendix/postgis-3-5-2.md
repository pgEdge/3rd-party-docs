## PostGIS 3.5.2


2025/01/18


This version requires PostgreSQL 12-17, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.12+ is needed. To take advantage of all SFCGAL features, SFCGAL 1.5.0+ is needed.


## Bug Fixes


[#5677](https://trac.osgeo.org/postgis/ticket/5677), Retain SRID during unary union (Paul Ramsey)


[#5833](https://trac.osgeo.org/postgis/ticket/5833), pg_upgrade fix for postgis_sfcgal (Regina Obe)


[#5564](https://trac.osgeo.org/postgis/ticket/5564), BRIN crash fix and support for parallel in PG17+ (Paul Ramsey, Regina Obe)
