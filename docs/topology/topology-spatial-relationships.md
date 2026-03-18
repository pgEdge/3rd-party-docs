<a id="Topology_Relationships"></a>

## Topology Spatial Relationships
  <a id="TG_Equals"></a>

# Equals

Returns true if two topogeometries are composed of the same topology primitives.

## Synopsis


```sql
boolean Equals(topogeometry  tg1, topogeometry  tg2)
```


## Description


Returns true if two topogeometries are composed of the same topology primitives: faces, edges, nodes.


!!! note

    This function not supported for topogeometries that are geometry collections. It also can not compare topogeometries from different topologies.


Availability: 1.1.0


## Examples


```
```


## See Also


[GetTopoGeomElements](topogeometry-accessors.md#GetTopoGeomElements), [ST_Equals](../postgis-reference/spatial-relationships.md#ST_Equals)
  <a id="TG_Intersects"></a>

# Intersects

Returns true if any pair of primitives from the two topogeometries intersect.

## Synopsis


```sql
boolean Intersects(topogeometry  tg1, topogeometry  tg2)
```


## Description


 Returns true if any pair of primitives from the two topogeometries intersect.


!!! note

    This function not supported for topogeometries that are geometry collections. It also can not compare topogeometries from different topologies. Also not currently supported for hierarchical topogeometries (topogeometries composed of other topogeometries).


Availability: 1.1.0


## Examples


```
```


## See Also


[ST_Intersects](../postgis-reference/spatial-relationships.md#ST_Intersects)
