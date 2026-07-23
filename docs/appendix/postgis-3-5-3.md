## PostGIS 3.5.3


2025/05/17


This version requires PostgreSQL 12-18beta1, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.12+ is needed. To take advantage of all SFCGAL features, SFCGAL 1.5.0+ is needed.


## Bug Fixes


Do not complain about illegal option when calling shp2pgsql -? (Sandro Santilli, Giovanni Zezza)


[5862](https://trac.osgeo.org/postgis/ticket/5862), [topology] Prevent another topology corruption with TopoGeo_addPoint near almost collinear edges (Sandro Santilli)


[5841](https://trac.osgeo.org/postgis/ticket/5841), Change approach to interrupt handling to conform to PgSQL recommended practice (Paul Ramsey)


[5855](https://trac.osgeo.org/postgis/ticket/5855), Fix index binding in ST_DFullyWithin (Paul Ramsey)


[5819](https://trac.osgeo.org/postgis/ticket/5819), Support longer names in estimated extent (Paul Ramsey)


Fix misassignment of result in _lwt_HealEdges (Maxim Korotkov)


[5876](https://trac.osgeo.org/postgis/ticket/5876), ST_AddPoint with empty argument adds garbage (Paul Ramsey)


[5874](https://trac.osgeo.org/postgis/ticket/5874), Line substring returns wrong answer (Paul Ramsey)


[5829](https://trac.osgeo.org/postgis/ticket/5829), geometry_columns with non-standard constraints (Paul Ramsey)


[5818](https://trac.osgeo.org/postgis/ticket/5818), [GT-244](https://git.osgeo.org/gitea/postgis/postgis/pulls/244) Fix CG_IsSolid function (Loïc Bartoletti)


[5885](https://trac.osgeo.org/postgis/ticket/5885), Fix documentation about grid-based overlay operations (Sandro Santilli)


For SFCGAL 2.1.0+ prevent using deprecated functions (Regina Obe)
