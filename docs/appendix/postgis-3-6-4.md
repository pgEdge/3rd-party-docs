## PostGIS 3.6.4


2026/06/08


## Fixes


[6076](https://trac.osgeo.org/postgis/ticket/6076), ST_Length(geography) for GeometryCollection not consistent with geometry implementation


[GH-862](https://github.com/postgis/postgis/pull/862), Use SFCGAL undeprecated function sfcgal_geometry_tessellate() (Jean Felder)


Flatgeobuf schema mismatch vulnerability (NeuroWinter)


[5899](https://trac.osgeo.org/postgis/ticket/5899), pg_upgrade issue for non-standard geography SRID (Paul Ramsey)


[5916](https://trac.osgeo.org/postgis/ticket/5916), ST_Split hang with very large ordinates (Paul Ramsey)
