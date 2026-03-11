<a id="earthdistance"></a>

## earthdistance — calculate great-circle distances


 The `earthdistance` module provides two different approaches to calculating great circle distances on the surface of the Earth. The one described first depends on the `cube` module. The second one is based on the built-in `point` data type, using longitude and latitude for the coordinates.


 In this module, the Earth is assumed to be perfectly spherical. (If that's too inaccurate for you, you might want to look at the [PostGIS](https://postgis.net/) project.)


 The `cube` module must be installed before `earthdistance` can be installed (although you can use the `CASCADE` option of `CREATE EXTENSION` to install both in one command).


!!! caution

    It is strongly recommended that `earthdistance` and `cube` be installed in the same schema, and that that schema be one for which CREATE privilege has not been and will not be granted to any untrusted users. Otherwise there are installation-time security hazards if `earthdistance`'s schema contains objects defined by a hostile user. Furthermore, when using `earthdistance`'s functions after installation, the entire search path should contain only trusted schemas.
 <a id="earthdistance-cube-based"></a>

### Cube-Based Earth Distances


 Data is stored in cubes that are points (both corners are the same) using 3 coordinates representing the x, y, and z distance from the center of the Earth. A *domain* `earth` over type `cube` is provided, which includes constraint checks that the value meets these restrictions and is reasonably close to the actual surface of the Earth.


 The radius of the Earth is obtained from the `earth()` function. It is given in meters. But by changing this one function you can change the module to use some other units, or to use a different value of the radius that you feel is more appropriate.


 This package has applications to astronomical databases as well. Astronomers will probably want to change `earth()` to return a radius of `180/pi()` so that distances are in degrees.


 Functions are provided to support input in latitude and longitude (in degrees), to support output of latitude and longitude, to calculate the great circle distance between two points and to easily specify a bounding box usable for index searches.


 The provided functions are shown in [Cube-Based Earthdistance Functions](#earthdistance-cube-functions).
 <a id="earthdistance-cube-functions"></a>

**Table: Cube-Based Earthdistance Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>earth</code> () <code>float8</code></td>
<td>Returns the assumed radius of the Earth.</td>
<td></td>
</tr>
<tr>
<td><code>sec_to_gc</code> ( <code>float8</code> ) <code>float8</code></td>
<td>Converts the normal straight line (secant) distance between two points on the surface of the Earth to the great circle distance between them.</td>
<td></td>
</tr>
<tr>
<td><code>gc_to_sec</code> ( <code>float8</code> ) <code>float8</code></td>
<td>Converts the great circle distance between two points on the surface of the Earth to the normal straight line (secant) distance between them.</td>
<td></td>
</tr>
<tr>
<td><code>ll_to_earth</code> ( <code>float8</code>, <code>float8</code> ) <code>earth</code></td>
<td>Returns the location of a point on the surface of the Earth given its latitude (argument 1) and longitude (argument 2) in degrees.</td>
<td></td>
</tr>
<tr>
<td><code>latitude</code> ( <code>earth</code> ) <code>float8</code></td>
<td>Returns the latitude in degrees of a point on the surface of the Earth.</td>
<td></td>
</tr>
<tr>
<td><code>longitude</code> ( <code>earth</code> ) <code>float8</code></td>
<td>Returns the longitude in degrees of a point on the surface of the Earth.</td>
<td></td>
</tr>
<tr>
<td><code>earth_distance</code> ( <code>earth</code>, <code>earth</code> ) <code>float8</code></td>
<td>Returns the great circle distance between two points on the surface of the Earth.</td>
<td></td>
</tr>
<tr>
<td><code>earth_box</code> ( <code>earth</code>, <code>float8</code> ) <code>cube</code></td>
<td>Returns a box suitable for an indexed search using the <code>cube</code> <code>@&gt;</code> operator for points within a given great circle distance of a location. Some points in this box are further than the specified great circle distance from the location, so a second check using <code>earth_distance</code> should be included in the query.</td>
<td></td>
</tr>
</tbody>
</table>
  <a id="earthdistance-point-based"></a>

### Point-Based Earth Distances


 The second part of the module relies on representing Earth locations as values of type `point`, in which the first component is taken to represent longitude in degrees, and the second component is taken to represent latitude in degrees. Points are taken as (longitude, latitude) and not vice versa because longitude is closer to the intuitive idea of x-axis and latitude to y-axis.


 A single operator is provided, shown in [Point-Based Earthdistance Operators](#earthdistance-point-operators).
 <a id="earthdistance-point-operators"></a>

**Table: Point-Based Earthdistance Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>point</code> <code>&lt;@&gt;</code> <code>point</code> <code>float8</code></td>
<td>Computes the distance in statute miles between two points on the Earth's surface.</td>
<td></td>
</tr>
</tbody>
</table>


 Note that unlike the `cube`-based part of the module, units are hardwired here: changing the `earth()` function will not affect the results of this operator.


 One disadvantage of the longitude/latitude representation is that you need to be careful about the edge conditions near the poles and near +/- 180 degrees of longitude. The `cube`-based representation avoids these discontinuities.
