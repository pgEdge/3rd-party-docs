<a id="using-query-indexes"></a>

## Using Spatial Indexes


When constructing queries using spatial conditions, for best performance it is important to ensure that a spatial index is used, if one exists (see [Spatial Indexes](../data-management/spatial-indexes.md#build-indexes)). To do this, a spatial operator or index-aware function must be used in a <code>WHERE</code> or <code>ON</code> clause of the query.


Spatial operators include the bounding box operators (of which the most commonly used is [geometry_overlaps](../postgis-reference/operators.md#geometry_overlaps); see [Bounding Box Operators](../postgis-reference/operators.md#operators-bbox) for the full list) and the distance operators used in nearest-neighbor queries (the most common being [geometry_distance_knn](../postgis-reference/operators.md#geometry_distance_knn); see [Distance Operators](../postgis-reference/operators.md#operators-distance) for the full list.)


Index-aware functions automatically add a bounding box operator to the spatial condition. Index-aware functions include the named spatial relationship predicates [ST_Contains](../postgis-reference/spatial-relationships.md#ST_Contains), [ST_ContainsProperly](../postgis-reference/spatial-relationships.md#ST_ContainsProperly), [ST_CoveredBy](../postgis-reference/spatial-relationships.md#ST_CoveredBy), [ST_Covers](../postgis-reference/spatial-relationships.md#ST_Covers), [ST_Crosses](../postgis-reference/spatial-relationships.md#ST_Crosses), [ST_Intersects](../postgis-reference/spatial-relationships.md#ST_Intersects), [ST_Overlaps](../postgis-reference/spatial-relationships.md#ST_Overlaps), [ST_Touches](../postgis-reference/spatial-relationships.md#ST_Touches), [ST_Within](../postgis-reference/spatial-relationships.md#ST_Within), [ST_Within](../postgis-reference/spatial-relationships.md#ST_Within), and [ST_3DIntersects](../postgis-reference/spatial-relationships.md#ST_3DIntersects), and the distance predicates [ST_DWithin](../postgis-reference/spatial-relationships.md#ST_DWithin), [ST_DFullyWithin](../postgis-reference/spatial-relationships.md#ST_DFullyWithin), [ST_3DDFullyWithin](../postgis-reference/spatial-relationships.md#ST_3DDFullyWithin), and [ST_3DDWithin](../postgis-reference/spatial-relationships.md#ST_3DDWithin) .)


Functions such as [ST_Distance](../postgis-reference/measurement-functions.md#ST_Distance) do *not* use indexes to optimize their operation. For example, the following query would be quite slow on a large table:


```sql

SELECT geom
FROM geom_table
WHERE ST_Distance( geom, 'SRID=312;POINT(100000 200000)' ) < 100
```


This query selects all the geometries in <code>geom_table</code> which are within 100 units of the point (100000, 200000). It will be slow because it is calculating the distance between each point in the table and the specified point, ie. one `ST_Distance()` calculation is computed for **every** row in the table.


 The number of rows processed can be reduced substantially by using the index-aware function [ST_DWithin](../postgis-reference/spatial-relationships.md#ST_DWithin):


```sql
SELECT geom
FROM geom_table
WHERE ST_DWithin( geom, 'SRID=312;POINT(100000 200000)', 100 )
```


This query selects the same geometries, but it does it in a more efficient way. This is enabled by `ST_DWithin()` using the `&&` operator internally on an expanded bounding box of the query geometry. If there is a spatial index on <code>geom</code>, the query planner will recognize that it can use the index to reduce the number of rows scanned before calculating the distance. The spatial index allows retrieving only records with geometries whose bounding boxes overlap the expanded extent and hence which *might* be within the required distance. The actual distance is then computed to confirm whether to include the record in the result set.


For more information and examples see the [PostGIS Workshop](https://postgis.net/workshops/postgis-intro/indexing.html).
