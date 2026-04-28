## PostGIS 3.6.0


2025/09/01


This version requires PostgreSQL 12-18beta3, GEOS 3.8 or higher, and Proj 6.1+. To take advantage of all features, GEOS 3.14+ is needed. To take advantage of all SFCGAL features, SFCGAL 2.2+ is needed.


Many thanks to our translation teams, in particular:


Teramoto Ikuhiro (Japanese Team)


Daniel Nylander (Swedish Team)


Dapeng Wang, Zuo Chenwei from HighGo (Chinese Team)


Denys Kovshun (Ukrainian Team)


## Breaking Changes


[#5799](https://trac.osgeo.org/postgis/ticket/5799), make ST_TileEnvelope clips envelopes to tile plane extent (Paul Ramsey)


[#5829](https://trac.osgeo.org/postgis/ticket/5829), remove constraint checking from geometry_columns view (Paul Ramsey)


[#3373](https://trac.osgeo.org/postgis/ticket/3373), [GT-255](https://git.osgeo.org/gitea/postgis/postgis/pulls/255), [topology] Support for upgrading domains (Ayo Adesugba, U.S. Census Bureau)


[GT-252](https://git.osgeo.org/gitea/postgis/postgis/pulls/252), ST_NumGeometries/ST_GeometryN treat TIN and PolyhedralSurface as unitary geometries, use ST_NumPatches/ST_PatchN for patch access (Loïc Bartoletti)


[#3110](https://trac.osgeo.org/postgis/ticket/3110), [GT-242](https://git.osgeo.org/gitea/postgis/postgis/pulls/242), [topology] Support for bigint (Ayo Adesugba, U.S. Census Bureau)


[#5359](https://trac.osgeo.org/postgis/ticket/5359), [#5897](https://trac.osgeo.org/postgis/ticket/5897), [GT-260](https://git.osgeo.org/gitea/postgis/postgis/pulls/260) [tiger_geocoder] Use @extschema:extension@ for PG >= 16 to schema qualify dependent extensions, switch to use typmod for tiger tables (Regina Obe)


## Removed / Deprecate signatures


[#3110](https://trac.osgeo.org/postgis/ticket/3110), [GT-242](https://git.osgeo.org/gitea/postgis/postgis/pulls/242), [topology] Support for bigint (Ayo Adesugba, U.S. Census Bureau)


[#5498](https://trac.osgeo.org/postgis/ticket/5498) Drop st_approxquantile(raster, double precision), wasn't usable as it triggered is not unique error when used (Regina Obe)


## New Features


[GH-803](https://github.com/postgis/postgis/pull/803), [sfcgal] ADD CG_Simplify function (Loïc Bartoletti)


[GH-805](https://github.com/postgis/postgis/pull/805), [sfcgal] Add M support for SFCGAL >= 1.5.0 (Loïc Bartoletti)


[GH-801](https://github.com/postgis/postgis/pull/801), [sfcgal] ADD CG_3DAlphaWrapping function (Jean Felder)


[#5894](https://trac.osgeo.org/postgis/ticket/5894), [topology] TotalTopologySize (Sandro Santilli)


[#5890](https://trac.osgeo.org/postgis/ticket/5890), [topology] ValidateTopologyPrecision, MakeTopologyPrecise (Sandro Santilli)


[#5861](https://trac.osgeo.org/postgis/ticket/5861), [topology] Add --drop-topology switch to pgtopo_import (Sandro Santilli)


[#1247](https://trac.osgeo.org/postgis/ticket/1247), [raster] ST_AsRasterAgg (Sandro Santilli)


[#5784](https://trac.osgeo.org/postgis/ticket/5784), [GT-223](https://git.osgeo.org/gitea/postgis/postgis/pulls/223) Export circ_tree_distance_tree_internal for mobilitydb use (Maxime Schoemans)


[GT-228](https://git.osgeo.org/gitea/postgis/postgis/pulls/228) [sfcgal] Add new functions (Scale, Translate, Rotate, Buffer 3D and Straight Skeleton Partition) from SFCGAL 2 (Loïc Bartoletti)


[raster] New GUC postgis.gdal_cpl_debug, enables GDAL debugging messages and routes them into the PostgreSQL logging system. (Paul Ramsey)


[#5841](https://trac.osgeo.org/postgis/ticket/5841), Change interrupt handling to remove use of pqsignal to support PG 18 (Paul Ramsey)


Add ST_CoverageClean to edge match and gap remove polygonal coverages (Paul Ramsey) from GEOS 3.14 (Martin Davis)


[#3110](https://trac.osgeo.org/postgis/ticket/3110), [GT-242](https://git.osgeo.org/gitea/postgis/postgis/pulls/242) [topology] Support for bigint (Ayo Adesugba, U.S. Census Bureau)


[raster] Add ST_ReclassExact to quickly remap values in raster (Paul Ramsey)


[#5971](https://trac.osgeo.org/postgis/ticket/5971), [tiger] Option to build --without-tiger (Regina Obe)
