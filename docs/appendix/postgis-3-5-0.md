## PostGIS 3.5.0


2024/09/25


This version requires PostgreSQL 12-17, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.12+ is needed. To take advantage of all SFCGAL features, SFCGAL 1.5.0+ is needed.


Many thanks to our translation teams, in particular:


Dapeng Wang, Zuo Chenwei from HighGo (Chinese Team)


Teramoto Ikuhiro (Japanese Team)


Vincent Bre (French Team)


## Breaking Changes


[#5546](https://trac.osgeo.org/postgis/ticket/5546), TopoGeometry <> TopoGeometry is now ambiguous, to get the old behaviour, assuming your TopoGeometry objects are named tg1 and tg2, use: ( id(tg1) <> id(tg2) OR topology_id(tg1) <> topology_id(tg2) OR layer_id(tg1) <> layer_id(tg2) OR type(tg1) <> type(tg2) ) (Sandro Santilli)


[#5536](https://trac.osgeo.org/postgis/ticket/5536), comments are not anymore included in PostGIS extensions (Sandro Santilli)


xmllint is now required to build comments (Sandro Santilli)


DocBook5 XSL is now required to build html (Sandro Santilli)


[#5602](https://trac.osgeo.org/postgis/ticket/5602), Drop support for GEOS 3.6 and 3.7 (Regina Obe)


[#5571](https://trac.osgeo.org/postgis/ticket/5571), Improve ST_GeneratePoints performance, but old seeded pseudo random points will need to be regenerated.


[#5596](https://trac.osgeo.org/postgis/ticket/5596), GH-749, Allow promoting column as an id in ST_AsGeoJson(record,..). Views and materialized views that use the ST_AsGeoJSON(record ..) will need rebuilding to upgrade to new signature (Jan Tojnar)


[#5496](https://trac.osgeo.org/postgis/ticket/5496), ST_Clip all variants replaced, will require rebuilding of materialized views that use them (funding from The National Institute for Agricultural and Food Research and Technology (INIA-CSIC)), Regina Obe


[#5659](https://trac.osgeo.org/postgis/ticket/5659), ST_DFullyWithin behaviour has changed to be ST_Contains(ST_Buffer(A, R), B) (Paul Ramsey)


Remove the WFS_locks extra package. (Paul Ramsey)


[5747](https://trac.osgeo.org/postgis/ticket/5747), [GH-776](https://github.com/postgis/postgis/pull/776), ST_Length: Return 0 for CurvePolygon (Dan Baston)


[5770](https://trac.osgeo.org/postgis/ticket/5770), support for GEOS 3.13 and RelateNG. Most functionality remains the same, but new GEOS predicate implementation has a few small changes.


Boundary Node Rule relate matrices might be different when using the "multi-valent end point" rule.


Relate matrices for situations with invalid MultiPolygons with shared boundaries might be different. Run ST_MakeValid to get valid inputs to feed to the calculation.


Zero length LineStrings are treated as if they are the equivalent Point object.


## Deprecated signatures


[GH-761](https://github.com/postgis/postgis/pull/761), ST_StraightSkeleton => CG_StraightSkeleton (Loïc Bartoletti)


[GH-189](https://git.osgeo.org/gitea/postgis/postgis/pulls/189), All SFCGAL functions now use the prefix CG_, with the old ones using ST_ being deprecated. (Loïc Bartoletti)


## New features


Improvements in the 'postgis' script:

- new command list-enabled
- new command list-all
- command upgrade upgrades all databases that need to be
- command status reports status of all databases
 (Sandro Santilli)


[#5742](https://trac.osgeo.org/postgis/ticket/5742), expose version of PROJ at compile time (Sandro Santilli)


[#5721](https://trac.osgeo.org/postgis/ticket/5721), postgis_topology: Allow sharing sequences between different topologies (Lars Opsahl)


[#5667](https://trac.osgeo.org/postgis/ticket/5667), postgis_topology: TopoGeo_LoadGeometry (Sandro Santilli)


[#5055](https://trac.osgeo.org/postgis/ticket/5055), add explicit <> geometry operator to prevent non-unique error with <> and != (Paul Ramsey)


Add ST_HasZ/ST_HasM (Loïc Bartoletti)


[GT-123](https://git.osgeo.org/gitea/postgis/postgis/pulls/123), postgis_sfcgal: CG_YMonotonePartition, CG_ApproxConvexPartition, CG_GreeneApproxConvexPartition and CG_OptimalConvexPartition (Loïc Bartoletti)


[GT-156](https://git.osgeo.org/gitea/postgis/postgis/pulls/156), postgis_sfcgal: CG_Visibility (Loïc Bartoletti)


[GT-157](https://git.osgeo.org/gitea/postgis/postgis/pulls/157), postgis_sfcgal: Add ST_ExtrudeStraightSkeleton (Loïc Bartoletti)


[#5496](https://trac.osgeo.org/postgis/ticket/5496), postgis_raster: ST_Clip support for touched (Regina Obe)


[GH-760](https://github.com/postgis/postgis/pull/760), postgis_sfcgal: CG_Intersection, CG_3DIntersects, CG_Intersects, CG_Difference, CG_Union (and aggregate), CG_Triangulate, CG_Area, CG_3DDistance, CG_Distance (Loïc Bartoletti)


[#5687](https://trac.osgeo.org/postgis/ticket/5687), Don't rely on search_path to determine postgis schema Fix for PG17 security change (Regina Obe)


[#5705](https://trac.osgeo.org/postgis/ticket/5705), [GH-767](https://github.com/postgis/postgis/pull/767), ST_RemoveIrrelevantPointsForView (Sam Peters)


[#5706](https://trac.osgeo.org/postgis/ticket/5706), [GH-768](https://github.com/postgis/postgis/pull/768), ST_RemoveSmallParts (Sam Peters)


## Enhancements


[5550](https://trac.osgeo.org/postgis/ticket/5550), Fix upgrades from 2.x in sandboxed systems (Sandro Santilli)


[#3587](https://trac.osgeo.org/postgis/ticket/3587), postgis_topology: faster load of big lines in topologies (Sandro Santilli)


[#5670](https://trac.osgeo.org/postgis/ticket/5670), postgis_topology: faster ST_CreateTopoGeo (Sandro Santilli)


[#5531](https://trac.osgeo.org/postgis/ticket/5531), documentation format upgraded to DocBook 5 (Sandro Santilli)


[#5543](https://trac.osgeo.org/postgis/ticket/5543), allow building without documentation (Sandro Santilli)


[#5596](https://trac.osgeo.org/postgis/ticket/5596), [GH-749](https://github.com/postgis/postgis/pull/749), Allow promoting column as an id in ST_AsGeoJson(record,..). (Jan Tojnar)


[GH-744](https://github.com/postgis/postgis/pull/744), Don't create docbook.css for the HTML manual, use style.css instead (Chris Mayo)


Faster implementation of point-in-poly cached index (Paul Ramsey)


Improve performance of ST_GeneratePoints (Paul Ramsey)


[#5361](https://trac.osgeo.org/postgis/ticket/5361), ST_CurveN, ST_NumCurves and consistency in accessors on curved geometry (Paul Ramsey)


[GH-761](https://github.com/postgis/postgis/pull/761), postgis_sfcgal: Add an optional parameter to CG_StraightSkeleton (was ST_StraightSkeleton) to use m as a distance in result (Hannes Janetzek, Loïc Bartoletti)
