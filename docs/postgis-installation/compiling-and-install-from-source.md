<a id="PGInstall"></a>

## Compiling and Install from Source


!!! note

    Many OS systems now include pre-built packages for PostgreSQL/PostGIS. In many cases compilation is only necessary if you want the most bleeding edge versions or you are a package maintainer.


    This section includes general compilation instructions, if you are compiling for Windows etc or another OS, you may find additional more detailed help at [PostGIS User contributed compile guides](https://trac.osgeo.org/postgis/wiki/UsersWikiInstall) and [PostGIS Dev Wiki](http://trac.osgeo.org/postgis/wiki/DevWikiMain).


    Pre-Built Packages for various OS are listed in [PostGIS Pre-built Packages](https://trac.osgeo.org/postgis/wiki/UsersWikiPackages)


    If you are a windows user, you can get stable builds via Stackbuilder or [PostGIS Windows download site](https://postgis.net/windows_downloads) We also have [very bleeding-edge windows experimental builds](https://postgis.net/windows_downloads) that are built usually once or twice a week or whenever anything exciting happens. You can use these to experiment with the in progress releases of PostGIS


 The PostGIS module is an extension to the PostgreSQL backend server. As such, PostGIS 3.6.2 *requires* full PostgreSQL server headers access in order to compile. It can be built against PostgreSQL versions 12 - 18. Earlier versions of PostgreSQL are *not* supported.


 Refer to the PostgreSQL installation guides if you haven't already installed PostgreSQL. [https://www.postgresql.org](https://www.postgresql.org) .


!!! note

    For GEOS functionality, when you install PostgreSQL you may need to explicitly link PostgreSQL against the standard C++ library:


    ```
    LDFLAGS=-lstdc++ ./configure [YOUR OPTIONS HERE]
    ```


     This is a workaround for bogus C++ exceptions interaction with older development tools. If you experience weird problems (backend unexpectedly closed or similar things) try this trick. This will require recompiling your PostgreSQL from scratch, of course.


 The following steps outline the configuration and compilation of the PostGIS source. They are written for Linux users and will not work on Windows or Mac.
 <a id="install_getting_source"></a>

## Getting the Source


 Retrieve the PostGIS source archive from the downloads website [https://download.osgeo.org/postgis/source/postgis-3.6.2.tar.gz](https://download.osgeo.org/postgis/source/postgis-3.6.2.tar.gz)


```
wget https://download.osgeo.org/postgis/source/postgis-3.6.2.tar.gz
tar -xvzf postgis-3.6.2.tar.gz
cd postgis-3.6.2
```


 This will create a directory called `postgis-3.6.2` in the current working directory.


 Alternatively, checkout the source from the [git](https://git-scm.com/) repository [https://git.osgeo.org/gitea/postgis/postgis/](https://git.osgeo.org/gitea/postgis/postgis/) .


```
git clone https://git.osgeo.org/gitea/postgis/postgis.git postgis
cd postgis
sh autogen.sh

```


 Change into the newly created `postgis` directory to continue the installation.


```
./configure
```
  <a id="install_requirements"></a>

## Install Requirements


 PostGIS has the following requirements for building and usage:


 **Required**


-  PostgreSQL 12 - 18. A complete installation of PostgreSQL (including server headers) is required. PostgreSQL is available from [https://www.postgresql.org](https://www.postgresql.org)18 .

  For a full PostgreSQL / PostGIS support matrix and PostGIS/GEOS support matrix refer to [https://trac.osgeo.org/postgis/wiki/UsersWikiPostgreSQLPostGIS](https://trac.osgeo.org/postgis/wiki/UsersWikiPostgreSQLPostGIS)
-  GNU C compiler (`gcc`). Some other ANSI C compilers can be used to compile PostGIS, but we find far fewer problems when compiling with `gcc`.
-  GNU Make (`gmake` or `make`). For many systems, GNU `make` is the default version of make. Check the version by invoking `make -v`. Other versions of `make` may not process the PostGIS `Makefile` properly.
-  Proj reprojection library. Proj 6.1 or above is required. The Proj library is used to provide coordinate reprojection support within PostGIS. Proj is available for download from [https://proj.org/](https://proj.org/) .
-  GEOS geometry library, version 3.8.0 or greater, but GEOS 3.14+ is required to take full advantage of all the new functions and features. GEOS is available for download from [https://libgeos.org](https://libgeos.org/).
-  LibXML2, version 2.5.x or higher. LibXML2 is currently used in some imports functions (ST_GeomFromGML and ST_GeomFromKML). LibXML2 is available for download from [https://gitlab.gnome.org/GNOME/libxml2/-/releases](https://gitlab.gnome.org/GNOME/libxml2/-/releases).
-  JSON-C, version 0.9 or higher. JSON-C is currently used to import GeoJSON via the function ST_GeomFromGeoJson. JSON-C is available for download from [https://github.com/json-c/json-c/releases/](https://github.com/json-c/json-c/releases).
-  GDAL, version 3+ is preferred. This is required for raster support. [https://gdal.org/download.html](https://gdal.org/download.html).
-  If compiling with PostgreSQL+JIT, LLVM version >=6 is required [https://trac.osgeo.org/postgis/ticket/4125](https://trac.osgeo.org/postgis/ticket/4125).


 **Optional**


-  GDAL (pseudo optional) only if you don't want raster you can leave it out. Also make sure to enable the drivers you want to use as described in [Configuring raster support](../postgis-administration/configuring-raster-support.md#raster_configuration).
-  GTK (requires GTK+2.0, 2.8+) to compile the shp2pgsql-gui shape file loader. [http://www.gtk.org/](http://www.gtk.org/) .
-  SFCGAL, 1.4.1 or higher is required and 2.1+ is needed to be able to use all functionality. SFCGAL can be used to provide additional 2D and 3D advanced analysis functions to PostGIS cf [SFCGAL Functions Reference](../sfcgal-functions-reference/index.md#reference_sfcgal). And also allow to use SFCGAL rather than GEOS for some 2D functions provided by both backends (like ST_Intersection or ST_Area, for instance). A PostgreSQL configuration variable <code>postgis.backend</code> allow end user to control which backend he want to use if SFCGAL is installed (GEOS by default). Nota: SFCGAL 1.2 require at least CGAL 4.3 and Boost 1.54 (cf: [https://sfcgal.org](https://sfcgal.org)) [https://gitlab.com/sfcgal/SFCGAL/](https://gitlab.com/sfcgal/SFCGAL/).
-  In order to build the [Address Standardizer](../postgis-extras/address-standardizer.md#Address_Standardizer) you will also need PCRE 1 or 2 [http://www.pcre.org](http://www.pcre.org) (which generally is already installed on nix systems). [Address Standardizer](../postgis-extras/address-standardizer.md#Address_Standardizer) will automatically be built if it detects a PCRE library, or you pass in a valid `--with-pcre-dir=/path/to/pcre` during configure.
-  To enable ST_AsMVT protobuf-c library 1.1.0 or higher (for usage) and the protoc-c compiler (for building) are required. Also, pkg-config is required to verify the correct minimum version of protobuf-c. See [protobuf-c](https://github.com/protobuf-c/protobuf-c). By default, Postgis will use Wagyu to validate MVT polygons faster which requires a c++11 compiler. It will use CXXFLAGS and the same compiler as the PostgreSQL installation. To disable this and use GEOS instead use the `--without-wagyu` during the configure step.
-  CUnit (`CUnit`). This is needed for regression testing. [http://cunit.sourceforge.net/](http://cunit.sourceforge.net/)
-  DocBook (`xsltproc`) is required for building the documentation. Docbook is available from [http://www.docbook.org/](http://www.docbook.org/) .
-  DBLatex (`dblatex`) is required for building the documentation in PDF format. DBLatex is available from [http://dblatex.sourceforge.net/](http://dblatex.sourceforge.net/) .
-  ImageMagick (`convert`) is required to generate the images used in the documentation. ImageMagick is available from [http://www.imagemagick.org/](http://www.imagemagick.org/) .
  <a id="installation_configuration"></a>

## Build configuration


 As with most linux installations, the first step is to generate the Makefile that will be used to build the source code. This is done by running the shell script


 `./configure`


 With no additional parameters, this command will attempt to automatically locate the required components and libraries needed to build the PostGIS source code on your system. Although this is the most common usage of `./configure`, the script accepts several parameters for those who have the required libraries and programs in non-standard locations.


 The following list shows only the most commonly used parameters. For a complete list, use the `--help` or `--help=short` parameters.


`--with-library-minor-version`
:   Starting with PostGIS 3.0, the library files generated by default will no longer have the minor version as part of the file name. This means all PostGIS 3 libs will end in <code>postgis-3</code>. This was done to make pg_upgrade easier, with downside that you can only install one version PostGIS 3 series in your server. To get the old behavior of file including the minor version: e.g. <code>postgis-3.0</code> add this switch to your configure statement.

`--prefix=PREFIX`
:   This is the location the PostGIS loader executables and shared libs will be installed. By default, this location is the same as the detected PostgreSQL installation.


    !!! caution

        This parameter is currently broken, as the package will only install into the PostgreSQL installation directory. Visit [http://trac.osgeo.org/postgis/ticket/635](http://trac.osgeo.org/postgis/ticket/635) to track this bug.

`--with-pgconfig=FILE`
:   PostgreSQL provides a utility called `pg_config` to enable extensions like PostGIS to locate the PostgreSQL installation directory. Use this parameter (`--with-pgconfig=/path/to/pg_config`) to manually specify a particular PostgreSQL installation that PostGIS will build against.

`--with-gdalconfig=FILE`
:   GDAL, a required library, provides functionality needed for raster support `gdal-config` to enable software installations to locate the GDAL installation directory. Use this parameter (`--with-gdalconfig=/path/to/gdal-config`) to manually specify a particular GDAL installation that PostGIS will build against.

`--with-geosconfig=FILE`
:   GEOS, a required geometry library, provides a utility called `geos-config` to enable software installations to locate the GEOS installation directory. Use this parameter (`--with-geosconfig=/path/to/geos-config`) to manually specify a particular GEOS installation that PostGIS will build against.

`--with-xml2config=FILE`
:   LibXML is the library required for doing GeomFromKML/GML processes. It normally is found if you have libxml installed, but if not or you want a specific version used, you'll need to point PostGIS at a specific `xml2-config` confi file to enable software installations to locate the LibXML installation directory. Use this parameter (`>--with-xml2config=/path/to/xml2-config`) to manually specify a particular LibXML installation that PostGIS will build against.

`--with-projdir=DIR`
:   Proj is a reprojection library required by PostGIS. Use this parameter (`--with-projdir=/path/to/projdir`) to manually specify a particular Proj installation directory that PostGIS will build against.

`--with-libiconv=DIR`
:   Directory where iconv is installed.

`--with-jsondir=DIR`
:   [JSON-C](http://oss.metaparadigm.com/json-c/) is an MIT-licensed JSON library required by PostGIS ST_GeomFromJSON support. Use this parameter (`--with-jsondir=/path/to/jsondir`) to manually specify a particular JSON-C installation directory that PostGIS will build against.

`--with-pcredir=DIR`
:   [PCRE](http://www.pcre.org/) is an BSD-licensed Perl Compatible Regular Expression library required by address_standardizer extension. Use this parameter (`--with-pcredir=/path/to/pcredir`) to manually specify a particular PCRE installation directory that PostGIS will build against.

`--with-gui`
:   Compile the data import GUI (requires GTK+2.0). This will create shp2pgsql-gui graphical interface to shp2pgsql.

`--without-raster`
:   Compile without raster support.

`--without-tiger`
:   Disables tiger geocoder support.

`--without-topology`
:   Compile without topology support.

`--with-gettext=no`
:   By default PostGIS will try to detect gettext support and compile with it, however if you run into incompatibility issues that cause breakage of loader, you can disable it entirely with this command. Refer to ticket [http://trac.osgeo.org/postgis/ticket/748](http://trac.osgeo.org/postgis/ticket/748) for an example issue solved by configuring with this. NOTE: that you aren't missing much by turning this off. This is used for international help/label support for the GUI loader which is not yet documented and still experimental.

`--with-sfcgal=PATH`
:   By default PostGIS will not install with sfcgal support without this switch. `PATH` is an optional argument that allows to specify an alternate PATH to sfcgal-config.

`--without-phony-revision`
:   Disable updating postgis_revision.h to match current HEAD of the git repository.


!!! note

    If you obtained PostGIS from the [code repository](https://trac.osgeo.org/postgis/wiki/CodeRepository) , the first step is really to run the script


     `./autogen.sh`


     This script will generate the `configure` script that in turn is used to customize the installation of PostGIS.


     If you instead obtained PostGIS as a tarball, running `./autogen.sh` is not necessary as `configure` has already been generated.


## Building


 Once the Makefile has been generated, building PostGIS is as simple as running


 `make`


 The last line of the output should be "<code>PostGIS was built successfully. Ready to install.</code>"


 As of PostGIS v1.4.0, all the functions have comments generated from the documentation. If you wish to install these comments into your spatial databases later, run the command which requires docbook. The postgis_comments.sql and other package comments files raster_comments.sql, topology_comments.sql are also packaged in the tar.gz distribution in the doc folder so no need to make comments if installing from the tar ball. Comments are also included as part of the CREATE EXTENSION install.


 `make comments`


 Introduced in PostGIS 2.0. This generates html cheat sheets suitable for quick reference or for student handouts. This requires xsltproc to build and will generate 4 files in doc folder `topology_cheatsheet.html`, `tiger_geocoder_cheatsheet.html`, `raster_cheatsheet.html`, `postgis_cheatsheet.html`


You can download some pre-built ones available in html and pdf from [PostGIS / PostgreSQL Study Guides](http://www.postgis.us/study_guides)


 `make cheatsheets`
  <a id="make_install_postgis_extensions"></a>

## Building PostGIS Extensions and Deploying them


 The PostGIS extensions are built and installed automatically if you are using PostgreSQL 9.1+.


If you are building from source repository, you need to build the function descriptions first. These get built if you have docbook installed. You can also manually build with the statement:


 `make comments`


Building the comments is not necessary if you are building from a release tar ball since these are packaged pre-built with the tar ball already.


The extensions should automatically build as part of the make install process. You can if needed build from the extensions folders or copy files if you need them on a different server.


```
cd extensions
cd postgis
make clean
make
export PGUSER=postgres #overwrite psql variables
make check #to test before install
make install
# to test extensions
make check RUNTESTFLAGS=--extension
```


!!! note

    <code>make check</code> uses psql to run tests and as such can use psql environment variables. Common ones useful to override are `PGUSER`,`PGPORT`, and `PGHOST`. Refer to [psql environment variables](https://www.postgresql.org/docs/current/libpq-envars.html)


The extension files will always be the same for the same version of PostGIS and PostgreSQL regardless of OS, so it is fine to copy over the extension files from one OS to another as long as you have the PostGIS binaries already installed on your servers.


If you want to install the extensions manually on a separate server different from your development, You need to copy the following files from the extensions folder into the `PostgreSQL / share / extension` folder of your PostgreSQL install as well as the needed binaries for regular PostGIS if you don't have them already on the server.


-  These are the control files that denote information such as the version of the extension to install if not specified. `postgis.control, postgis_topology.control`.
-  All the files in the /sql folder of each extension. Note that these need to be copied to the root of the PostgreSQL share/extension folder `extensions/postgis/sql/*.sql`, `extensions/postgis_topology/sql/*.sql`


Once you do that, you should see `postgis`, `postgis_topology` as available extensions in PgAdmin -> extensions.


If you are using psql, you can verify that the extensions are installed by running this query:


```sql
SELECT name, default_version,installed_version
FROM pg_available_extensions WHERE name LIKE 'postgis%' or name LIKE 'address%';

             name             | default_version | installed_version
------------------------------+-----------------+-------------------
 address_standardizer         | 3.6.2         | 3.6.2
 address_standardizer_data_us | 3.6.2         | 3.6.2
 postgis                      | 3.6.2         | 3.6.2
 postgis_raster               | 3.6.2         | 3.6.2
 postgis_sfcgal               | 3.6.2         |
 postgis_tiger_geocoder       | 3.6.2         | 3.6.2
 postgis_topology             | 3.6.2         |
(6 rows)
```


If you have the extension installed in the database you are querying, you'll see mention in the `installed_version` column. If you get no records back, it means you don't have postgis extensions installed on the server at all. PgAdmin III 1.14+ will also provide this information in the `extensions` section of the database browser tree and will even allow upgrade or uninstall by right-clicking.


If you have the extensions available, you can install postgis extension in your database of choice by either using pgAdmin extension interface or running these sql commands:


```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION postgis_sfcgal;
CREATE EXTENSION fuzzystrmatch; --needed for postgis_tiger_geocoder
--optional used by postgis_tiger_geocoder, or can be used standalone
CREATE EXTENSION address_standardizer;
CREATE EXTENSION address_standardizer_data_us;
CREATE EXTENSION postgis_tiger_geocoder;
CREATE EXTENSION postgis_topology;
```


In psql you can use to see what versions you have installed and also what schema they are installed.


```
\connect mygisdb
\x
\dx postgis*
```


```
List of installed extensions
-[ RECORD 1 ]-------------------------------------------------
Name        | postgis
Version     | 3.6.2
Schema      | public
Description | PostGIS geometry, geography, and raster spat..
-[ RECORD 2 ]-------------------------------------------------
Name        | postgis_raster
Version     | 3.0.0dev
Schema      | public
Description | PostGIS raster types and functions
-[ RECORD 3 ]-------------------------------------------------
Name        | postgis_tiger_geocoder
Version     | 3.6.2
Schema      | tiger
Description | PostGIS tiger geocoder and reverse geocoder
-[ RECORD 4 ]-------------------------------------------------
Name        | postgis_topology
Version     | 3.6.2
Schema      | topology
Description | PostGIS topology spatial types and functions
```


!!! warning

    Extension tables `spatial_ref_sys`, `layer`, `topology` can not be explicitly backed up. They can only be backed up when the respective `postgis` or `postgis_topology` extension is backed up, which only seems to happen when you backup the whole database. As of PostGIS 2.0.1, only srid records not packaged with PostGIS are backed up when the database is backed up so don't go around changing srids we package and expect your changes to be there. Put in a ticket if you find an issue. The structures of extension tables are never backed up since they are created with <code>CREATE EXTENSION</code> and assumed to be the same for a given version of an extension. These behaviors are built into the current PostgreSQL extension model, so nothing we can do about it.


If you installed 3.6.2, without using our wonderful extension system, you can change it to be extension based by running the below commands to package the functions in their respective extension. Installing using `unpackaged` was removed in PostgreSQL 13, so you are advised to switch to an extension build before upgrading to PostgreSQL 13.


```sql

CREATE EXTENSION postgis FROM unpackaged;
CREATE EXTENSION postgis_raster FROM unpackaged;
CREATE EXTENSION postgis_topology FROM unpackaged;
CREATE EXTENSION postgis_tiger_geocoder FROM unpackaged;
```


## Testing


 If you wish to test the PostGIS build, run


 `make check`


 The above command will run through various checks and regression tests using the generated library against an actual PostgreSQL database.


!!! note

    If you configured PostGIS using non-standard PostgreSQL, GEOS, or Proj locations, you may need to add their library locations to the `LD_LIBRARY_PATH` environment variable.


!!! caution

    Currently, the `make check` relies on the <code>PATH</code> and <code>PGPORT</code> environment variables when performing the checks - it does *not* use the PostgreSQL version that may have been specified using the configuration parameter `--with-pgconfig`. So make sure to modify your PATH to match the detected PostgreSQL installation during configuration or be prepared to deal with the impending headaches.


 If successful, make check will produce the output of almost 500 tests. The results will look similar to the following (numerous lines omitted below):


```


     CUnit - A unit testing framework for C - Version 2.1-3
     http://cunit.sourceforge.net/

	.
	.
	.

Run Summary:    Type  Total    Ran Passed Failed Inactive
              suites     44     44    n/a      0        0
               tests    300    300    300      0        0
             asserts   4215   4215   4215      0      n/a
Elapsed time =    0.229 seconds

	.
	.
	.

Running tests

	.
	.
	.

Run tests: 134
Failed: 0


-- if you build with SFCGAL

	.
	.
	.

Running tests

	.
	.
	.

Run tests: 13
Failed: 0

-- if you built with raster support

	.
	.
	.

Run Summary:    Type  Total    Ran Passed Failed Inactive
              suites     12     12    n/a      0        0
               tests     65     65     65      0        0
             asserts  45896  45896  45896      0      n/a


	.
	.
	.

Running tests

	.
	.
	.

Run tests: 101
Failed: 0

-- topology regress

.
.
.

Running tests

	.
	.
	.

Run tests: 51
Failed: 0

-- if you built --with-gui, you should see this too

     CUnit - A unit testing framework for C - Version 2.1-2
     http://cunit.sourceforge.net/

	.
	.
	.

Run Summary:    Type  Total    Ran Passed Failed Inactive
              suites      2      2    n/a      0        0
               tests      4      4      4      0        0
             asserts      4      4      4      0      n/a
```


The `postgis_tiger_geocoder` and `address_standardizer` extensions, currently only support the standard PostgreSQL installcheck. To test these use the below. Note: the make install is not necessary if you already did make install at root of PostGIS code folder.


For address_standardizer:

```
cd extensions/address_standardizer
make install
make installcheck

```


Output should look like:

```
============== dropping database "contrib_regression" ==============
DROP DATABASE
============== creating database "contrib_regression" ==============
CREATE DATABASE
ALTER DATABASE
============== running regression test queries        ==============
test test-init-extensions     ... ok
test test-parseaddress        ... ok
test test-standardize_address_1 ... ok
test test-standardize_address_2 ... ok

=====================
 All 4 tests passed.
=====================
```


For tiger geocoder, make sure you have postgis and fuzzystrmatch extensions available in your PostgreSQL instance. The address_standardizer tests will also kick in if you built postgis with address_standardizer support:

```
cd extensions/postgis_tiger_geocoder
make install
make installcheck

```


output should look like:

```
============== dropping database "contrib_regression" ==============
DROP DATABASE
============== creating database "contrib_regression" ==============
CREATE DATABASE
ALTER DATABASE
============== installing fuzzystrmatch               ==============
CREATE EXTENSION
============== installing postgis                     ==============
CREATE EXTENSION
============== installing postgis_tiger_geocoder      ==============
CREATE EXTENSION
============== installing address_standardizer        ==============
CREATE EXTENSION
============== running regression test queries        ==============
test test-normalize_address   ... ok
test test-pagc_normalize_address ... ok

=====================
All 2 tests passed.
=====================
```


## Installation


 To install PostGIS, type


 `make install`


 This will copy the PostGIS installation files into their appropriate subdirectory specified by the `--prefix` configuration parameter. In particular:


-  The loader and dumper binaries are installed in `[prefix]/bin`.
-  The SQL files, such as `postgis.sql`, are installed in `[prefix]/share/contrib`.
-  The PostGIS libraries are installed in `[prefix]/lib`.


 If you previously ran the `make comments` command to generate the `postgis_comments.sql`, `raster_comments.sql` file, install the sql file by running


 `make comments-install`


!!! note

    `postgis_comments.sql`, `raster_comments.sql`, `topology_comments.sql` was separated from the typical build and installation targets since with it comes the extra dependency of `xsltproc`.
