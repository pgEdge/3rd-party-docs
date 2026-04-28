<a id="TopoGeometry_Outputs"></a>

## TopoGeometry Outputs
  <a id="AsGML"></a>

# AsGML

Returns the GML representation of a topogeometry.

## Synopsis


```sql
text AsGML(topogeometry  tg)
text AsGML(topogeometry  tg, text  nsprefix_in)
text AsGML(topogeometry  tg, regclass  visitedTable)
text AsGML(topogeometry  tg, regclass  visitedTable, text  nsprefix)
text AsGML(topogeometry  tg, text  nsprefix_in, integer  precision, integer  options)
text AsGML(topogeometry  tg, text  nsprefix_in, integer  precision, integer  options, regclass  visitedTable)
text AsGML(topogeometry  tg, text  nsprefix_in, integer  precision, integer  options, regclass  visitedTable, text  idprefix)
text AsGML(topogeometry  tg, text  nsprefix_in, integer  precision, integer  options, regclass  visitedTable, text  idprefix, int  gmlversion)
```


## Description


Returns the GML representation of a topogeometry in version GML3 format. If no `nsprefix_in` is specified then `gml` is used. Pass in an empty string for nsprefix to get a non-qualified name space. The precision (default: 15) and options (default 1) parameters, if given, are passed untouched to the underlying call to ST_AsGML.


 The `visitedTable` parameter, if given, is used for keeping track of the visited Node and Edge elements so to use cross-references (xlink:xref) rather than duplicating definitions. The table is expected to have (at least) two integer fields: 'element_type' and 'element_id'. The calling user must have both read and write privileges on the given table. For best performance, an index should be defined on `element_type` and `element_id`, in that order. Such index would be created automatically by adding a unique constraint to the fields. Example:

```sql

CREATE TABLE visited (
  element_type integer, element_id integer,
  unique(element_type, element_id)
);
```


The `idprefix` parameter, if given, will be prepended to Edge and Node tag identifiers.


The `gmlver` parameter, if given, will be passed to the underlying ST_AsGML. Defaults to 3.


Availability: 2.0.0


## Examples


This uses the topo geometry we created in [CreateTopoGeom](topogeometry-constructors.md#CreateTopoGeom)


```sql

SELECT topology.AsGML(topo) As rdgml
  FROM ri.roads
  WHERE road_name = 'Unknown';

-- rdgml--
<gml:TopoCurve>
    <gml:directedEdge>
        <gml:Edge gml:id="E1">
            <gml:directedNode orientation="-">
                <gml:Node gml:id="N1"/>
            </gml:directedNode>
            <gml:directedNode></gml:directedNode>
            <gml:curveProperty>
                <gml:Curve srsName="urn:ogc:def:crs:EPSG::3438">
                    <gml:segments>
                        <gml:LineStringSegment>
                            <gml:posList srsDimension="2">384744 236928 384750 236923 384769 236911 384799 236895 384811 236890
                            384833 236884 384844 236882 384866 236881 384879 236883 384954 236898 385087 236932 385117 236938
                            385167 236938 385203 236941 385224 236946 385233 236950 385241 236956 385254 236971
                            385260 236979 385268 236999 385273 237018 385273 237037 385271 237047 385267 237057 385225 237125
                            385210 237144 385192 237161 385167 237192 385162 237202 385159 237214 385159 237227 385162 237241
                            385166 237256 385196 237324 385209 237345 385234 237375 385237 237383 385238 237399 385236 237407
                            385227 237419 385213 237430 385193 237439 385174 237451 385170 237455 385169 237460 385171 237475
                            385181 237503 385190 237521 385200 237533 385206 237538 385213 237541 385221 237542 385235 237540 385242 237541
                            385249 237544 385260 237555 385270 237570 385289 237584 385292 237589 385291 237596 385284 237630</gml:posList>
                        </gml:LineStringSegment>
                    </gml:segments>
                </gml:Curve>
            </gml:curveProperty>
        </gml:Edge>
    </gml:directedEdge>
</gml:TopoCurve>
```


Same exercise as previous without namespace


```sql

SELECT topology.AsGML(topo,'') As rdgml
  FROM ri.roads
  WHERE road_name = 'Unknown';

-- rdgml--
<TopoCurve>
    <directedEdge>
        <Edge id="E1">
            <directedNode orientation="-">
                <Node id="N1"/>
            </directedNode>
            <directedNode></directedNode>
            <curveProperty>
                <Curve srsName="urn:ogc:def:crs:EPSG::3438">
                    <segments>
                        <LineStringSegment>
                            <posList srsDimension="2">384744 236928 384750 236923 384769 236911 384799 236895 384811 236890
                            384833 236884 384844 236882 384866 236881 384879 236883 384954 236898 385087 236932 385117 236938
                            385167 236938 385203 236941 385224 236946 385233 236950 385241 236956 385254 236971
                            385260 236979 385268 236999 385273 237018 385273 237037 385271 237047 385267 237057 385225 237125
                            385210 237144 385192 237161 385167 237192 385162 237202 385159 237214 385159 237227 385162 237241
                            385166 237256 385196 237324 385209 237345 385234 237375 385237 237383 385238 237399 385236 237407
                            385227 237419 385213 237430 385193 237439 385174 237451 385170 237455 385169 237460 385171 237475
                            385181 237503 385190 237521 385200 237533 385206 237538 385213 237541 385221 237542 385235 237540 385242 237541
                            385249 237544 385260 237555 385270 237570 385289 237584 385292 237589 385291 237596 385284 237630</posList>
                         </LineStringSegment>
                    </segments>
                </Curve>
            </curveProperty>
        </Edge>
    </directedEdge>
</TopoCurve>
```


## See Also


[CreateTopoGeom](topogeometry-constructors.md#CreateTopoGeom), [ST_CreateTopoGeo](topology-constructors.md#ST_CreateTopoGeo)
  <a id="AsTopoJSON"></a>

# AsTopoJSON

Returns the TopoJSON representation of a topogeometry.

## Synopsis


```sql
text AsTopoJSON(topogeometry  tg, regclass  edgeMapTable)
```


## Description


Returns the TopoJSON representation of a topogeometry. If `edgeMapTable` is not null, it will be used as a lookup/storage mapping of edge identifiers to arc indices. This is to be able to allow for a compact "arcs" array in the final document.


 The table, if given, is expected to have an "arc_id" field of type "serial" and an "edge_id" of type integer; the code will query the table for "edge_id" so it is recommended to add an index on that field.


!!! note

    Arc indices in the TopoJSON output are 0-based but they are 1-based in the "edgeMapTable" table.


 A full TopoJSON document will be need to contain, in addition to the snippets returned by this function, the actual arcs plus some headers. See the [TopoJSON specification](http://github.com/mbostock/topojson-specification/blob/master/README.md).


Availability: 2.1.0


Enhanced: 2.2.1 added support for puntal inputs


## See Also


[ST_AsGeoJSON](../postgis-reference/geometry-output.md#ST_AsGeoJSON)


## Examples


```sql

CREATE TEMP TABLE edgemap(arc_id serial, edge_id int unique);

-- header
SELECT '{ "type": "Topology", "transform": { "scale": [1,1], "translate": [0,0] }, "objects": {'

-- objects
UNION ALL SELECT '"' || feature_name || '": ' || AsTopoJSON(feature, 'edgemap')
FROM features.big_parcels WHERE feature_name = 'P3P4';

-- arcs
WITH edges AS (
  SELECT m.arc_id, e.geom FROM edgemap m, city_data.edge e
  WHERE e.edge_id = m.edge_id
), points AS (
  SELECT arc_id, (st_dumppoints(geom)).* FROM edges
), compare AS (
  SELECT p2.arc_id,
         CASE WHEN p1.path IS NULL THEN p2.geom
              ELSE ST_Translate(p2.geom, -ST_X(p1.geom), -ST_Y(p1.geom))
         END AS geom
  FROM points p2 LEFT OUTER JOIN points p1
  ON ( p1.arc_id = p2.arc_id AND p2.path[1] = p1.path[1]+1 )
  ORDER BY arc_id, p2.path
), arcsdump AS (
  SELECT arc_id, (regexp_matches( ST_AsGeoJSON(geom), '\[.*\]'))[1] as t
  FROM compare
), arcs AS (
  SELECT arc_id, '[' || array_to_string(array_agg(t), ',') || ']' as a FROM arcsdump
  GROUP BY arc_id
  ORDER BY arc_id
)
SELECT '}, "arcs": [' UNION ALL
SELECT array_to_string(array_agg(a), E',\n') from arcs

-- footer
UNION ALL SELECT ']}'::text as t;

-- Result:
{ "type": "Topology", "transform": { "scale": [1,1], "translate": [0,0] }, "objects": {
"P3P4": { "type": "MultiPolygon", "arcs": [[[-1]],[[6,5,-5,-4,-3,1]]]}
}, "arcs": [
 [[25,30],[6,0],[0,10],[-14,0],[0,-10],[8,0]],
 [[35,6],[0,8]],
 [[35,6],[12,0]],
 [[47,6],[0,8]],
 [[47,14],[0,8]],
 [[35,22],[12,0]],
 [[35,14],[0,8]]
 ]}
```
