<a id="Topology"></a>

# Topology

The PostGIS Topology types and functions are used to manage topological objects such as faces, edges and nodes.

Sandro Santilli's presentation at PostGIS Day Paris 2011 conference gives a good synopsis of PostGIS Topology and where it is headed [Topology with PostGIS 2.0 slide deck](http://strk.kbt.io/projects/postgis/Paris2011_TopologyWithPostGIS_2_0.pdf).

 Vincent Picavet provides a good synopsis and overview of what is Topology, how is it used, and various FOSS4G tools that support it in [PostGIS Topology PGConf EU 2012](https://gitlab.com/Oslandia/documentation/presentations/-/blob/master/2012/pgconf_eu_2012/pgconfeu2012_vincent_picavet_postgis_topology.pdf).

An example of a topologically based GIS database is the [US Census Topologically Integrated Geographic Encoding and Referencing System (TIGER)](https://www.census.gov/geo/maps-data/data/tiger.html) database. If you want to experiment with PostGIS topology and need some data, check out [Topology_Load_Tiger](../postgis-extras/tiger-geocoder.md#Topology_Load_Tiger).

The PostGIS topology module has existed in prior versions of PostGIS but was never part of the Official PostGIS documentation. In PostGIS 2.0.0 major cleanup is going on to remove use of all deprecated functions in it, fix known usability issues, better document the features and functions, add new functions, and enhance to closer conform to SQL-MM standards.

Details of this project can be found at [PostGIS Topology Wiki](http://trac.osgeo.org/postgis/wiki/UsersWikiPostgisTopology)

All functions and tables associated with this module are installed in a schema called `topology`.

Functions that are defined in SQL/MM standard are prefixed with ST_ and functions specific to PostGIS are not prefixed.

Topology support is build by default starting with PostGIS 2.0, and can be disabled specifying --without-topology configure option at build time as described in [PostGIS Installation](../postgis-installation/index.md#postgis_installation)

- [Topology Types](topology-types.md#Topology_Types)
- [Topology Domains](topology-domains.md#Topology_Domains)
- [Topology and TopoGeometry Management](topology-and-topogeometry-management.md#Topology_ManagementFunctions)
- [Topology Statistics Management](topology-statistics-management.md#Topology_StatsManagement)
- [Topology Constructors](topology-constructors.md#Topology_Constructors)
- [Topology Editors](topology-editors.md#Topology_Editing)
- [Topology Accessors](topology-accessors.md#Topology_Accessors)
- [Topology Processing](topology-processing.md#Topology_Processing)
- [TopoGeometry Constructors](topogeometry-constructors.md#TopoGeometry_Constructors)
- [TopoGeometry Editors](topogeometry-editors.md#TopoGeometry_Editors)
- [TopoGeometry Accessors](topogeometry-accessors.md#TopoGeom_Accessors)
- [TopoGeometry Outputs](topogeometry-outputs.md#TopoGeometry_Outputs)
- [Topology Spatial Relationships](topology-spatial-relationships.md#Topology_Relationships)
- [Importing and exporting Topologies](importing-and-exporting-topologies.md#Topology_Import_Export)
