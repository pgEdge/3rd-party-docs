<a id="Topology_Import_Export"></a>

## Importing and exporting Topologies


 Once you have created topologies, and maybe associated topological layers, you might want to export them into a file-based format for backup or transfer into another database.


 Using the standard dump/restore tools of PostgreSQL is problematic because topologies are composed by a set of tables (4 for primitives, an arbitrary number for layers) and records in metadata tables (topology.topology and topology.layer). Additionally, topology identifiers are not univoque across databases so that parameter of your topology will need to be changes upon restoring it.


 In order to simplify export/restore of topologies a pair of executables are provided: `pgtopo_export` and `pgtopo_import`. Example usage:


```

pgtopo_export dev_db topo1 | pgtopo_import topo1 | psql staging_db
```
 <a id="pgtopo_export"></a>

## Using the Topology exporter


 The `pgtopo_export` script takes the name of a database and a topology and outputs a dump file which can be used to import the topology (and associated layers) into a new database.


 By default `pgtopo_export` writes the dump file to the standard output so that it can be piped to `pgtopo_import` or redirected to a file (refusing to write to terminal). You can optionally specify an output filename with the <code>-f</code> commandline switch.


 By default `pgtopo_export` includes a dump of all layers defined against the given topology. This may be more data than you need, or may be non-working (in case your layer tables have complex dependencies) in which case you can request skipping the layers with the <code>--skip-layers</code> switch and deal with those separately.


 Invoking `pgtopo_export` with the <code>--help</code> (or <code>-h</code> for short) switch will always print short usage string.


 The dump file format is a compressed tar archive of a `pgtopo_export` directory containing at least a `pgtopo_dump_version` file with format version info. As of version <code>1</code> the directory contains tab-delimited CSV files with data of the topology primitive tables (node, edge_data, face, relation), the topology and layer records associated with it and (unless <code>--skip-layers</code> is given) a custom-format PostgreSQL dump of tables reported as being layers of the given topology.
   <a id="pgtopo_import"></a>

## Using the Topology importer


 The `pgtopo_import` script takes a <code>pgtopo_export</code> format topology dump and a name to give to the topology to be created and outputs an SQL script reconstructing the topology and associated layers.


 The generated SQL file will contain statements that create a topology with the given name, load primitive data in it, restores and registers all topology layers by properly linking all TopoGeometry values to their correct topology.


 By default `pgtopo_import` reads the dump from the standard input so that it can be used in conjunction with `pgtopo_export` in a pipeline. You can optionally specify an input filename with the <code>-f</code> commandline switch.


 By default `pgtopo_import` includes in the output SQL file the code to restore all layers found in the dump.


 This may be unwanted or non-working in case your target database already have tables with the same name as the ones in the dump. In that case you can request skipping the layers with the <code>--skip-layers</code> switch and deal with those separately (or later).


 SQL to only load and link layers to a named topology can be generated using the <code>--only-layers</code> switch. This can be useful to load layers AFTER resolving the naming conflicts or to link layers to a different topology (say a spatially-simplified version of the starting topology).
