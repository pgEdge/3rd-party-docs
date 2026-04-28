<a id="TopoGeometry_Editors"></a>

## TopoGeometry Editors
  <a id="clearTopoGeom"></a>

# clearTopoGeom

Clears the content of a topo geometry.

## Synopsis


```sql
topogeometry clearTopoGeom(topogeometry  topogeom)
```


## Description


 Clears the content a [topogeometry](topology-types.md#topogeometry) turning it into an empty one. Mostly useful in conjunction with [toTopoGeom](topogeometry-constructors.md#toTopoGeom) to replace the shape of existing objects and any dependent object in higher hierarchical levels.


Availability: 2.1


## Examples


```

-- Shrink all TopoGeometry polygons by 10 meters
UPDATE nei_topo SET topo = ST_Buffer(clearTopoGeom(topo), -10);

```


## See Also


 [toTopoGeom](topogeometry-constructors.md#toTopoGeom)
  <a id="TopoGeom_addElement"></a>

# TopoGeom_addElement

Adds an element to the definition of a TopoGeometry.

## Synopsis


```sql
topogeometry TopoGeom_addElement(topogeometry  tg, topoelement  el)
```


## Description


 Adds a [topoelement](topology-domains.md#topoelement) to the definition of a TopoGeometry object. Does not error out if the element is already part of the definition.


Availability: 2.3


## Examples


```

-- Add edge 5 to TopoGeometry tg
UPDATE mylayer SET tg = TopoGeom_addElement(tg, '{5,2}');

```


## See Also


 [TopoGeom_remElement](#TopoGeom_remElement), [CreateTopoGeom](topogeometry-constructors.md#CreateTopoGeom)
  <a id="TopoGeom_remElement"></a>

# TopoGeom_remElement

Removes an element from the definition of a TopoGeometry.

## Synopsis


```sql
topogeometry TopoGeom_remElement(topogeometry  tg, topoelement  el)
```


## Description


 Removes a [topoelement](topology-domains.md#topoelement) from the definition of a TopoGeometry object.


Availability: 2.3


## Examples


```

-- Remove face 43 from TopoGeometry tg
UPDATE mylayer SET tg = TopoGeom_remElement(tg, '{43,3}');

```


## See Also


 [TopoGeom_addElement](#TopoGeom_addElement), [CreateTopoGeom](topogeometry-constructors.md#CreateTopoGeom)
  <a id="TopoGeom_addTopoGeom"></a>

# TopoGeom_addTopoGeom

Adds element of a TopoGeometry to the definition of another TopoGeometry.

## Synopsis


```sql
topogeometry TopoGeom_addTopoGeom(topogeometry  tgt, topogeometry  src)
```


## Description


 Adds the elements of a [topogeometry](topology-types.md#topogeometry) to the definition of another TopoGeometry, possibly changing its cached type (type attribute) to a collection, if needed to hold all elements in the source object.


 The two TopoGeometry objects need be defined against the *same* topology and, if hierarchically defined, need be composed by elements of the same child layer.


Availability: 3.2


## Examples


```

-- Set an "overall" TopoGeometry value to be composed by all
-- elements of specific TopoGeometry values
UPDATE mylayer SET tg_overall = TopoGeom_addTopogeom(
    TopoGeom_addTopoGeom(
        clearTopoGeom(tg_overall),
        tg_specific1
    ),
    tg_specific2
);

```


## See Also


 [TopoGeom_addElement](#TopoGeom_addElement), [clearTopoGeom](#clearTopoGeom), [CreateTopoGeom](topogeometry-constructors.md#CreateTopoGeom)
  <a id="toTopoGeom_editor_proxy"></a>

# toTopoGeom

Adds a geometry shape to an existing topo geometry.

## Description


 Refer to [toTopoGeom](topogeometry-constructors.md#toTopoGeom).
