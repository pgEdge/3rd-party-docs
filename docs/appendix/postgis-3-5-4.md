## PostGIS 3.5.4


2025/10/16


This version requires PostgreSQL 12-18beta1, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.12+ is needed. To take advantage of all SFCGAL features, SFCGAL 1.5.0+ is needed.


## Bug Fixes


[5977](https://trac.osgeo.org/postgis/ticket/5977), Fix downgrade protection with standard conforming strings off (Sandro Santilli)


[5951](https://trac.osgeo.org/postgis/ticket/5951), Fix crash in ST_GetFaceEdges with corrupted topology (Sandro Santilli)


[5947](https://trac.osgeo.org/postgis/ticket/5947), [topology] Fix crash in ST_ModEdgeHeal (Sandro Santilli)


[5925](https://trac.osgeo.org/postgis/ticket/5925), [5946](https://trac.osgeo.org/postgis/ticket/5946), [topology] Have GetFaceContainingPoint survive EMPTY edges (Sandro Santilli)


[5936](https://trac.osgeo.org/postgis/ticket/5936), [topology] Do script-based upgrade in a single transaction (Sandro Santilli)


[5908](https://trac.osgeo.org/postgis/ticket/5908), [topology] Fix crash in GetFaceContainingPoint (Sandro Santilli)


[5907](https://trac.osgeo.org/postgis/ticket/5907), [topology] Fix crash in TopoGeo_AddPolygon with EMPTY input (Sandro Santilli)


[5922](https://trac.osgeo.org/postgis/ticket/5922), [topology] Fix crash in TopoGeo_AddLinestring with EMPTY input (Sandro Santilli)


[5921](https://trac.osgeo.org/postgis/ticket/5921), Crash freeing uninitialized pointer (Arsenii Mukhin)


[5912](https://trac.osgeo.org/postgis/ticket/5912), Crash on GML with xlink and no prefix (Paul Ramsey)


[5905](https://trac.osgeo.org/postgis/ticket/5905), Crash on deeply nested geometries (Paul Ramsey)


[5909](https://trac.osgeo.org/postgis/ticket/5909), ST_ValueCount crashes on empty table (Paul Ramsey)


[5917](https://trac.osgeo.org/postgis/ticket/5917), ST_Relate becomes unresponsive (Paul Ramsey)


[5923](https://trac.osgeo.org/postgis/ticket/5923), CG_ExtrudeStraightSkeleton crashes on empty polygon (Loïc Bartoletti)


[5935](https://trac.osgeo.org/postgis/ticket/5935), Require GDAL 2.4 for postgis_raster and switch to GDALGetDataTypeSizeBytes (Laurențiu Nicola)


GT-257, fix issue with xsltproc with path has spaces (Laurențiu Nicola)


[5938](https://trac.osgeo.org/postgis/ticket/5938), incorrect parameter order in ST_Relate caching (Paul Ramsey)


[5927](https://trac.osgeo.org/postgis/ticket/5927), ST_IsCollection throwing exception (Paul Ramsey)


[5902](https://trac.osgeo.org/postgis/ticket/5902), ST_PointFromText cannot create geometries with M (Paul Ramsey)


[5943](https://trac.osgeo.org/postgis/ticket/5943), Memory leak in handling GEOS GeometryFactory (Megan Ma)


[5407](https://trac.osgeo.org/postgis/ticket/5407), Use memset in place of bzero (Paul Ramsey)


[5082](https://trac.osgeo.org/postgis/ticket/5082), LRS proportions clamped to [0,1] (Pawel Ostrowski)


[5985](https://trac.osgeo.org/postgis/ticket/5985), Fix configure issue with Debian 12 and 13 (Regina Obe, Sandro Santilli)


[5991](https://trac.osgeo.org/postgis/ticket/5991), CircularString distance error (Paul Ramsey)


[5994](https://trac.osgeo.org/postgis/ticket/5994), Null pointer in ST_AsGeoJsonRow (Alexander Kukushkin)


[5989](https://trac.osgeo.org/postgis/ticket/5989), ST_Distance error on CurvePolygon (Paul Ramsey)


[5962](https://trac.osgeo.org/postgis/ticket/5962), Consistent clipping of MULTI/POINT (Paul Ramsey)


[5754](https://trac.osgeo.org/postgis/ticket/5754), ST_ForcePolygonCCW reverses lines (Paul Ramsey)
