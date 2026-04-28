## PostGIS 3.5.1


2024/12/22


This version requires PostgreSQL 12-17, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.12+ is needed. To take advantage of all SFCGAL features, SFCGAL 1.5.0+ is needed.


## Breaking Changes


[#5677](https://trac.osgeo.org/postgis/ticket/5677), Retain SRID during unary union (Paul Ramsey)


[#5792](https://trac.osgeo.org/postgis/ticket/5792), [topology] Prevent topology corruption with TopoGeo_addPoint near almost collinear edges (Sandro Santilli)


[#5795](https://trac.osgeo.org/postgis/ticket/5795), [topology] Fix ST_NewEdgesSplit can cause invalid topology (Björn Harrtell)


[#5794](https://trac.osgeo.org/postgis/ticket/5794), [topology] Fix crash in TopoGeo_addPoint (Sandro Santilli)


[#5785](https://trac.osgeo.org/postgis/ticket/5785), [raster] ST_MapAlgebra segfaults when expression references a supernumerary rast argument (Dian M Fay)


[#5787](https://trac.osgeo.org/postgis/ticket/5787), Check that ST_ChangeEdgeGeom doesn't change winding of rings (Sandro Santilli)


[#5791](https://trac.osgeo.org/postgis/ticket/5791), Add legacy stubs for old transaction functions to allow pg_upgrade (Regina Obe)


[#5800](https://trac.osgeo.org/postgis/ticket/5800), PROJ compiled version reading the wrong minor and micro (Regina Obe)


[#5790](https://trac.osgeo.org/postgis/ticket/5790), Non-schema qualified calls causing issue with materialized views (Regina Obe)


[#5812](https://trac.osgeo.org/postgis/ticket/5812), Performance regression in ST_Within (Paul Ramsey)


[#5815](https://trac.osgeo.org/postgis/ticket/5815), Remove hash/merge promise from <> operator (Paul Ramsey)


[#5823](https://trac.osgeo.org/postgis/ticket/5823), Build support for Pg18 (Paul Ramsey)


## Enhancements


[#5782](https://trac.osgeo.org/postgis/ticket/5782), Improve robustness of min distance calculation (Sandro Santilli)


[topology] Speedup topology building when closing large rings with many holes (Björn Harrtell)


[#5810](https://trac.osgeo.org/postgis/ticket/5810), Update tiger geocoder to handle TIGER 2024 data (Regina Obe)


## Breaking Changes


[#5799](https://trac.osgeo.org/postgis/ticket/5799), make ST_TileEnvelope clip envelopes to tile plane extent (Paul Ramsey)
