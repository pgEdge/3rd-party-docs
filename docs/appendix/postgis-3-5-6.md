## PostGIS 3.5.6


2026/04/14


## Bug Fixes


[6055](https://trac.osgeo.org/postgis/ticket/6055), Remove rare extension priv escalation case. Reported by Sven Klemm (Tiger Data), Allistair Ishmael Hakim (allistair.sh) and Daniel Bakker


[GH-850](https://github.com/postgis/postgis/pull/850), Use quote_identifier to build tables in pgis_tablefromflatgeobuf (Ariel Mashraki)


[6058](https://trac.osgeo.org/postgis/ticket/6058), Use Pg composite_to_json() function in 19+ (Paul Ramsey)


[6060](https://trac.osgeo.org/postgis/ticket/6060), fully quality calls to helper functions (Paul Ramsey)


[6026](https://trac.osgeo.org/postgis/ticket/6026), KNN failure in rare IEEE double rounding case (Paul Ramsey)


[6061](https://trac.osgeo.org/postgis/ticket/6061), WKT parser produces incorrect error locations (Paul Ramsey)
