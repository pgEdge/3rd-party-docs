<a id="Topology_Domains"></a>

## Topology Domains
  <a id="topoelement"></a>

# TopoElement

An array of 2 integers generally used to identify a TopoGeometry component.

## Description


 An array of 2 integers used to represent one component of a simple or hierarchical [topogeometry](topology-types.md#topogeometry).


 In the case of a simple TopoGeometry the first element of the array represents the identifier of a topological primitive and the second element represents its type (1:node, 2:edge, 3:face). In the case of a hierarchical TopoGeometry the first element of the array represents the identifier of a child TopoGeometry and the second element represents its layer identifier.


!!! note

    For any given hierarchical TopoGeometry all child TopoGeometry elements will come from the same child layer, as specified in the topology.layer record for the layer of the TopoGeometry being defined.


## Examples


```sql

SELECT te[1] AS id, te[2] AS type FROM
( SELECT ARRAY[1,2]::topology.topoelement AS te ) f;
 id | type
----+------
  1 |    2

```


```sql
SELECT ARRAY[1,2]::topology.topoelement;
  te
-------
 {1,2}

```


```

--Example of what happens when you try to case a 3 element array to topoelement
-- NOTE: topoement has to be a 2 element array so fails dimension check
SELECT ARRAY[1,2,3]::topology.topoelement;
ERROR:  value for domain topology.topoelement violates check constraint "dimensions"

```


## See Also


 [GetTopoGeomElements](topogeometry-accessors.md#GetTopoGeomElements), [topoelementarray](#topoelementarray), [topogeometry](topology-types.md#topogeometry), [TopoGeom_addElement](topogeometry-editors.md#TopoGeom_addElement), [TopoGeom_remElement](topogeometry-editors.md#TopoGeom_remElement)
  <a id="topoelementarray"></a>

# TopoElementArray

An array of TopoElement objects.

## Description


An array of 1 or more TopoElement objects, generally used to pass around components of TopoGeometry objects.


## Examples


```sql
SELECT '{{1,2},{4,3}}'::topology.topoelementarray As tea;
  tea
-------
{{1,2},{4,3}}

-- more verbose equivalent --
SELECT ARRAY[ARRAY[1,2], ARRAY[4,3]]::topology.topoelementarray As tea;

  tea
-------
{{1,2},{4,3}}

--using the array agg function packaged with topology --
SELECT topology.TopoElementArray_Agg(ARRAY[e,t]) As tea
  FROM generate_series(1,4) As e CROSS JOIN generate_series(1,3) As t;
  tea
--------------------------------------------------------------------------
{{1,1},{1,2},{1,3},{2,1},{2,2},{2,3},{3,1},{3,2},{3,3},{4,1},{4,2},{4,3}}

```


```sql
SELECT '{{1,2,4},{3,4,5}}'::topology.topoelementarray As tea;
ERROR:  value for domain topology.topoelementarray violates check constraint "dimensions"

```


## See Also


 [topoelement](#topoelement), [GetTopoGeomElementArray](topogeometry-accessors.md#GetTopoGeomElementArray), [TopoElementArray_Agg](topogeometry-constructors.md#TopoElementArray_Agg)
