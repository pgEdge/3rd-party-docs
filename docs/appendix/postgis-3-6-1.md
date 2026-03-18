## PostGIS 3.6.1


2025/11/13


If you are upgrading postgis_topology and have topogeometry columns, make sure to run after the upgrade to fix topogeometry corruption: <code>SELECT topology.FixCorruptTopoGeometryColumn(schema_name, table_name, feature_column) FROM topology.layer</code>


## Fixes


[#5978](https://trac.osgeo.org/postgis/ticket/5978), geometry_columns needs to still parse table constraints (Paul Ramsey)


[#5987](https://trac.osgeo.org/postgis/ticket/5987), ST_GeometryN fails for non-collections (Paul Ramsey)


[#5991](https://trac.osgeo.org/postgis/ticket/5991), CircularString distance error (Paul Ramsey)


[#5994](https://trac.osgeo.org/postgis/ticket/5994), Null pointer in ST_AsGeoJsonRow (Alexander Kukushkin)


[#5998](https://trac.osgeo.org/postgis/ticket/5998), ST_Distance error on CurvePolygon (Paul Ramsey)


[#5962](https://trac.osgeo.org/postgis/ticket/5962), Consistent clipping of MULTI/POINT (Paul Ramsey)


[#5998](https://trac.osgeo.org/postgis/ticket/5998), [tiger_geocoder] [security] CVE-2022-2625, make sure tables requires by extension are owned by extension authored: Andrey Borodin (Yandex), reported by Sergey Bobrov (Kaspersky)


[#5754](https://trac.osgeo.org/postgis/ticket/5754), ST_ForcePolygonCCW reverses lines (Paul Ramsey)


[#5959](https://trac.osgeo.org/postgis/ticket/5959), [#5984](https://trac.osgeo.org/postgis/ticket/5984), Prevent histogram target overflow when analysing massive tables (Darafei Praliaskouski)


[#6012](https://trac.osgeo.org/postgis/ticket/6012), Remove memory leak from lwcircstring_from_lwpointarray (Paul Ramsey)


[#6013](https://trac.osgeo.org/postgis/ticket/6013), [tiger_geocoder] Load Tiger 2025 data (Regina Obe)


[#5983](https://trac.osgeo.org/postgis/ticket/5983), [topology] topology.FixCorruptTopoGeometryColumn to fix corruption caused by 3.6.0 upgrade (Regina Obe, Francois Bonzon)
