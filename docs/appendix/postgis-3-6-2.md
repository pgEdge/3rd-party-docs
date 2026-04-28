## PostGIS 3.6.2


2026/02/09


If you are upgrading postgis_topology and have topogeometry columns, make sure to run after the upgrade to fix topogeometry corruption: <code>SELECT topology.FixCorruptTopoGeometryColumn(schema_name, table_name, feature_column) FROM topology.layer</code>


## Fixes


[6023](https://trac.osgeo.org/postgis/ticket/6023), Fix robustness issue in ptarray_contains_point (Sandro Santilli)


[6027](https://trac.osgeo.org/postgis/ticket/6027), Fix RemoveUnusedPrimitives without topology in search_path (Sandro Santilli)


[6019](https://trac.osgeo.org/postgis/ticket/6019), make clean does not remove cunit generated files (Bas Couwenberg)


[6020](https://trac.osgeo.org/postgis/ticket/6020), schema qualify call in ST_MPointFromText (Paul Ramsey)


[6028](https://trac.osgeo.org/postgis/ticket/6028), crash indexing malformed empty polygon (Paul Ramsey)


[GH-841](https://github.com/postgis/postgis/pull/841), small memory leak in address_standardizer (Maxim Korotkov)


[5853](https://trac.osgeo.org/postgis/ticket/5853), Issue with topology and tiger geocoder upgrade scripts (Regina Obe, Spencer Bryson)


[6032](https://trac.osgeo.org/postgis/ticket/6032), Fix postgis_tiger_geocoder upgrade for PostgreSQL < 16 (Regina Obe)
