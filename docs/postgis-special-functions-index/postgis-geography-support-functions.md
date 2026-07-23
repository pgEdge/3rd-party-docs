<a id="PostGIS_GeographyFunctions"></a>

## PostGIS Geography Support Functions


The functions and operators given below are PostGIS functions/operators that take as input or return as output a [geography](../data-management/geography-data-type.md#PostGIS_Geography) data type object.


!!! note

    Functions with a (T) are not native geodetic functions, and use a ST_Transform call to and from geometry to do the operation. As a result, they may not behave as expected when going over dateline, poles, and for large geometries or geometry pairs that cover more than one UTM zone. Basic transform - (favoring UTM, Lambert Azimuthal (North/South), and falling back on mercator in worst case scenario)
