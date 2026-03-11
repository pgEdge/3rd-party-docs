<a id="install-meson"></a>

## Building and Installation with Meson
  <a id="install-short-meson"></a>

### Short Version


```

meson setup build --prefix=/usr/local/pgsql
cd build
ninja
su
ninja install
adduser postgres
mkdir -p /usr/local/pgsql/data
chown postgres /usr/local/pgsql/data
su - postgres
/usr/local/pgsql/bin/initdb -D /usr/local/pgsql/data
/usr/local/pgsql/bin/pg_ctl -D /usr/local/pgsql/data -l logfile start
/usr/local/pgsql/bin/createdb test
/usr/local/pgsql/bin/psql test
```
 The long version is the rest of this section.
  <a id="install-procedure-meson"></a>

### Installation Procedure


<a id="meson-configure"></a>
1.

    The first step of the installation procedure is to configure the build tree for your system and choose the options you would like. To create and configure the build directory, you can start with the `meson setup` command.

```

meson setup build
```
    The setup command takes a `builddir` and a `srcdir` argument. If no `srcdir` is given, Meson will deduce the `srcdir` based on the current directory and the location of `meson.build`. The `builddir` is mandatory.

    Running `meson setup` loads the build configuration file and sets up the build directory. Additionally, you can also pass several build options to Meson. Some commonly used options are mentioned in the subsequent sections. For example:

```

# configure with a different installation prefix
meson setup build --prefix=/home/user/pg-install

# configure to generate a debug build
meson setup build --buildtype=debug

# configure to build with OpenSSL support
meson setup build -Dssl=openssl
```


    Setting up the build directory is a one-time step. To reconfigure before a new build, you can simply use the `meson configure` command

```

meson configure -Dcassert=true
```
    `meson configure`'s commonly used command-line options are explained in [`meson setup` Options](#meson-options).
<a id="meson-build"></a>
2.

    By default, Meson uses the [Ninja](https://ninja-build.org/) build tool. To build PostgreSQL from source using Meson, you can simply use the `ninja` command in the build directory.

```

ninja
```
    Ninja will automatically detect the number of CPUs in your computer and parallelize itself accordingly. You can override the number of parallel processes used with the command line argument `-j`.

    It should be noted that after the initial configure step, `ninja` is the only command you ever need to type to compile. No matter how you alter your source tree (short of moving it to a completely new location), Meson will detect the changes and regenerate itself accordingly. This is especially handy if you have multiple build directories. Often one of them is used for development (the "debug" build) and others only every now and then (such as a "static analysis" build). Any configuration can be built just by cd'ing to the corresponding directory and running Ninja.

    If you'd like to build with a backend other than ninja, you can use configure with the `--backend` option to select the one you want to use and then build using `meson compile`. To learn more about these backends and other arguments you can provide to ninja, you can refer to the [Meson documentation](https://mesonbuild.com/Running-Meson.html#building-from-the-source).
3.

    If you want to test the newly built server before you install it, you can run the regression tests at this point. The regression tests are a test suite to verify that PostgreSQL runs on your machine in the way the developers expected it to. Type:

```

meson test
```
    (This won't work as root; do it as an unprivileged user.) See [Regression Tests](../regression-tests/index.md#regress) for detailed information about interpreting the test results. You can repeat this test at any later time by issuing the same command.

    To run pg_regress and pg_isolation_regress tests against a running postgres instance, specify `--setup running` as an argument to `meson test`.
<a id="meson-install"></a>
4.

!!! note

    If you are upgrading an existing system be sure to read [Upgrading a PostgreSQL Cluster](../server-setup-and-operation/upgrading-a-postgresql-cluster.md#upgrading), which has instructions about upgrading a cluster.

    Once PostgreSQL is built, you can install it by simply running the `ninja install` command.

```

ninja install
```


    This will install files into the directories that were specified in [Configuration](#meson-configure). Make sure that you have appropriate permissions to write into that area. You might need to do this step as root. Alternatively, you can create the target directories in advance and arrange for appropriate permissions to be granted. The standard installation provides all the header files needed for client application development as well as for server-side program development, such as custom functions or data types written in C.

    `ninja install` should work for most cases, but if you'd like to use more options (such as `--quiet` to suppress extra output), you could also use `meson install` instead. You can learn more about [meson install](https://mesonbuild.com/Commands.html#install) and its options in the Meson documentation.


**Uninstallation:.**


 To undo the installation, you can use the `ninja uninstall` command.


**Cleaning:.**


 After the installation, you can free disk space by removing the built files from the source tree with the `ninja clean` command.
   <a id="meson-options"></a>

### `meson setup` Options


 `meson setup`'s command-line options are explained below. This list is not exhaustive (use `meson configure --help` to get one that is). The options not covered here are meant for advanced use-cases, and are documented in the standard [Meson documentation](https://mesonbuild.com/Commands.html#configure). These arguments can be used with `meson setup` as well.
 <a id="meson-options-locations"></a>

#### Installation Locations


 These options control where `ninja install` (or `meson install`) will put the files. The `--prefix` option (example [Short Version](#install-short-meson)) is sufficient for most cases. If you have special needs, you can customize the installation subdirectories with the other options described in this section. Beware however that changing the relative locations of the different subdirectories may render the installation non-relocatable, meaning you won't be able to move it after installation. (The `man` and `doc` locations are not affected by this restriction.) For relocatable installs, you might want to use the `-Drpath=false` option described later.


<a id="configure-prefix-meson"></a>

<code>--prefix=</code><em>PREFIX</em>
:   Install all files under the directory *PREFIX* instead of `/usr/local/pgsql` (on Unix based systems) or <em>current drive letter</em><code>:/usr/local/pgsql</code> (on Windows). The actual files will be installed into various subdirectories; no files will ever be installed directly into the *PREFIX* directory.
<a id="configure-bindir-meson"></a>

<code>--bindir=</code><em>DIRECTORY</em>
:   Specifies the directory for executable programs. The default is <em>PREFIX</em><code>/bin</code>.
<a id="configure-sysconfdir-meson"></a>

<code>--sysconfdir=</code><em>DIRECTORY</em>
:   Sets the directory for various configuration files, <em>PREFIX</em><code>/etc</code> by default.
<a id="configure-libdir-meson"></a>

<code>--libdir=</code><em>DIRECTORY</em>
:   Sets the location to install libraries and dynamically loadable modules. The default is <em>PREFIX</em><code>/lib</code>.
<a id="configure-includedir-meson"></a>

<code>--includedir=</code><em>DIRECTORY</em>
:   Sets the directory for installing C and C++ header files. The default is <em>PREFIX</em><code>/include</code>.
<a id="configure-datadir-meson"></a>

<code>--datadir=</code><em>DIRECTORY</em>
:   Sets the directory for read-only data files used by the installed programs. The default is <em>PREFIX</em><code>/share</code>. Note that this has nothing to do with where your database files will be placed.
<a id="configure-localedir-meson"></a>

<code>--localedir=</code><em>DIRECTORY</em>
:   Sets the directory for installing locale data, in particular message translation catalog files. The default is <em>DATADIR</em><code>/locale</code>.
<a id="configure-mandir-meson"></a>

<code>--mandir=</code><em>DIRECTORY</em>
:   The man pages that come with PostgreSQL will be installed under this directory, in their respective <code>man</code><em>x</em> subdirectories. The default is <em>DATADIR</em><code>/man</code>.


!!! note

    Care has been taken to make it possible to install PostgreSQL into shared installation locations (such as `/usr/local/include`) without interfering with the namespace of the rest of the system. First, the string “`/postgresql`” is automatically appended to `datadir`, `sysconfdir`, and `docdir`, unless the fully expanded directory name already contains the string “`postgres`” or “`pgsql`”. For example, if you choose `/usr/local` as prefix, the documentation will be installed in `/usr/local/doc/postgresql`, but if the prefix is `/opt/postgres`, then it will be in `/opt/postgres/doc`. The public C header files of the client interfaces are installed into `includedir` and are namespace-clean. The internal header files and the server header files are installed into private directories under `includedir`. See the documentation of each interface for information about how to access its header files. Finally, a private subdirectory will also be created, if appropriate, under `libdir` for dynamically loadable modules.
  <a id="meson-options-features"></a>

#### PostgreSQL Features


 The options described in this section enable building of various optional PostgreSQL features. Most of these require additional software, as described in [Requirements](requirements.md#install-requirements), and will be automatically enabled if the required software is found. You can change this behavior by manually setting these features to `enabled` to require them or `disabled` to not build with them.


 To specify PostgreSQL-specific options, the name of the option must be prefixed by `-D`.


<a id="configure-with-nls-meson"></a>

`-Dnls={ auto | enabled | disabled }`
:   Enables or disables Native Language Support (NLS), that is, the ability to display a program's messages in a language other than English. Defaults to auto and will be enabled automatically if an implementation of the Gettext API is found.
<a id="configure-with-plperl-meson"></a>

`-Dplperl={ auto | enabled | disabled }`
:   Build the PL/Perl server-side language. Defaults to auto.
<a id="configure-with-plpython-meson"></a>

`-Dplpython={ auto | enabled | disabled }`
:   Build the PL/Python server-side language. Defaults to auto.
<a id="configure-with-pltcl-meson"></a>

`-Dpltcl={ auto | enabled | disabled }`
:   Build the PL/Tcl server-side language. Defaults to auto.
<a id="configure-with-tcl-version-meson"></a>

<code>-Dtcl_version=</code><em>TCL_VERSION</em>
:   Specifies the Tcl version to use when building PL/Tcl.
<a id="configure-with-icu-meson"></a>

`-Dicu={ auto | enabled | disabled }`
:   Build with support for the ICU library, enabling use of ICU collation features (see [Collation Support](../localization/collation-support.md#collation)). Defaults to auto and requires the ICU4C package to be installed. The minimum required version of ICU4C is currently 4.2.
<a id="configure-with-llvm-meson"></a>

`-Dllvm={ auto | enabled | disabled }`
:   Build with support for LLVM based JIT compilation (see [Just-in-Time Compilation (JIT)](../just-in-time-compilation-jit/index.md#jit)). This requires the LLVM library to be installed. The minimum required version of LLVM is currently 14. Disabled by default.


     `llvm-config` will be used to find the required compilation options. `llvm-config`, and then `llvm-config-$version` for all supported versions, will be searched for in your `PATH`. If that would not yield the desired program, use `LLVM_CONFIG` to specify a path to the correct `llvm-config`.
<a id="configure-with-lz4-meson"></a>

`-Dlz4={ auto | enabled | disabled }`
:   Build with LZ4 compression support. Defaults to auto.
<a id="configure-with-zstd-meson"></a>

`-Dzstd={ auto | enabled | disabled }`
:   Build with Zstandard compression support. Defaults to auto.
<a id="configure-with-ssl-meson"></a>

<code>-Dssl={ auto | </code><em>LIBRARY</em><code> }</code>
:   Build with support for SSL (encrypted) connections. The only *LIBRARY* supported is `openssl`. This requires the OpenSSL package to be installed. Building with this will check for the required header files and libraries to make sure that your OpenSSL installation is sufficient before proceeding. The default for this option is auto.
<a id="configure-with-gssapi-meson"></a>

`-Dgssapi={ auto | enabled | disabled }`
:   Build with support for GSSAPI authentication. MIT Kerberos is required to be installed for GSSAPI. On many systems, the GSSAPI system (a part of the MIT Kerberos installation) is not installed in a location that is searched by default (e.g., `/usr/include`, `/usr/lib`). In those cases, PostgreSQL will query `pkg-config` to detect the required compiler and linker options. Defaults to auto. `meson configure` will check for the required header files and libraries to make sure that your GSSAPI installation is sufficient before proceeding.
<a id="configure-with-ldap-meson"></a>

`-Dldap={ auto | enabled | disabled }`
:   Build with LDAP support for authentication and connection parameter lookup (see <a id="install-ldap-links-meson"></a>
    [LDAP Lookup of Connection Parameters](../../client-interfaces/libpq-c-library/ldap-lookup-of-connection-parameters.md#libpq-ldap) and [LDAP Authentication](../client-authentication/ldap-authentication.md#auth-ldap) for more information). On Unix, this requires the OpenLDAP package to be installed. On Windows, the default WinLDAP library is used. Defaults to auto. `meson configure` will check for the required header files and libraries to make sure that your OpenLDAP installation is sufficient before proceeding.
<a id="configure-with-pam-meson"></a>

`-Dpam={ auto | enabled | disabled }`
:   Build with PAM (Pluggable Authentication Modules) support. Defaults to auto.
<a id="configure-with-bsd-auth-meson"></a>

`-Dbsd_auth={ auto | enabled | disabled }`
:   Build with BSD Authentication support. (The BSD Authentication framework is currently only available on OpenBSD.) Defaults to auto.
<a id="configure-with-systemd-meson"></a>

`-Dsystemd={ auto | enabled | disabled }`
:   Build with support for systemd service notifications. This improves integration if the server is started under systemd but has no impact otherwise; see [Starting the Database Server](../server-setup-and-operation/starting-the-database-server.md#server-start) for more information. Defaults to auto. libsystemd and the associated header files need to be installed to use this option.
<a id="configure-with-bonjour-meson"></a>

`-Dbonjour={ auto | enabled | disabled }`
:   Build with support for Bonjour automatic service discovery. Defaults to auto and requires Bonjour support in your operating system. Recommended on macOS.
<a id="configure-with-uuid-meson"></a>

<code>-Duuid=</code><em>LIBRARY</em>
:   Build the [uuid-ossp](../../appendixes/additional-supplied-modules-and-extensions/uuid-ossp-a-uuid-generator.md#uuid-ossp) module (which provides functions to generate UUIDs), using the specified UUID library. *LIBRARY* must be one of:


    -  `none` to not build the uuid module. This is the default.
    -  `bsd` to use the UUID functions found in FreeBSD, and some other BSD-derived systems
    -  `e2fs` to use the UUID library created by the `e2fsprogs` project; this library is present in most Linux systems and in macOS, and can be obtained for other platforms as well
    -  `ossp` to use the [OSSP UUID library](http://www.ossp.org/pkg/lib/uuid/)
<a id="configure-with-libcurl-meson"></a>

`-Dlibcurl={ auto | enabled | disabled }`
:   Build with libcurl support for OAuth 2.0 client flows. Libcurl version 7.61.0 or later is required for this feature. Building with this will check for the required header files and libraries to make sure that your Curl installation is sufficient before proceeding. The default for this option is auto.
<a id="configure-with-liburing-meson"></a>

`-Dliburing={ auto | enabled | disabled }`
:   Build with liburing, enabling io_uring support for asynchronous I/O. Defaults to auto.


     To use a liburing installation that is in an unusual location, you can set `pkg-config`-related environment variables (see its documentation).
<a id="configure-with-libnuma-meson"></a>

`-Dlibnuma={ auto | enabled | disabled }`
:   Build with libnuma support for basic NUMA support. Only supported on platforms for which the libnuma library is implemented. The default for this option is auto.
<a id="configure-with-libxml-meson"></a>

`-Dlibxml={ auto | enabled | disabled }`
:   Build with libxml2, enabling SQL/XML support. Defaults to auto. Libxml2 version 2.6.23 or later is required for this feature.


     To use a libxml2 installation that is in an unusual location, you can set `pkg-config`-related environment variables (see its documentation).
<a id="configure-with-libxslt-meson"></a>

`-Dlibxslt={ auto | enabled | disabled }`
:   Build with libxslt, enabling the [xml2](../../appendixes/additional-supplied-modules-and-extensions/xml2-xpath-querying-and-xslt-functionality.md#xml2) module to perform XSL transformations of XML. `-Dlibxml` must be specified as well. Defaults to auto.
<a id="configure-with-sepgsql-meson"></a>

`-Dselinux={ auto | enabled | disabled }`
:   Build with SElinux support, enabling the [sepgsql](../../appendixes/additional-supplied-modules-and-extensions/sepgsql-selinux-label-based-mandatory-access-control-mac-security-module.md#sepgsql) extension. Defaults to auto.
  <a id="meson-options-anti-features"></a>

#### Anti-Features


<a id="configure-readline-meson"></a>

`-Dreadline={ auto | enabled | disabled }`
:   Allows use of the Readline library (and libedit as well). This option defaults to auto and enables command-line editing and history in psql and is strongly recommended.
<a id="configure-libedit-preferred-meson"></a>

`-Dlibedit_preferred={ true | false }`
:   Setting this to true favors the use of the BSD-licensed libedit library rather than GPL-licensed Readline. This option is significant only if you have both libraries installed; the default is false, that is to use Readline.
<a id="configure-zlib-meson"></a>

`-Dzlib={ auto | enabled | disabled }`
:   Enables use of the Zlib library. It defaults to auto and enables support for compressed archives in pg_dump, pg_restore and pg_basebackup and is recommended.
  <a id="meson-options-build-process"></a>

#### Build Process Details


<a id="configure-auto-features-meson"></a>

`--auto-features={ auto | enabled | disabled }`
:   Setting this option allows you to override the value of all “auto” features (features that are enabled automatically if the required software is found). This can be useful when you want to disable or enable all the “optional” features at once without having to set each of them manually. The default value for this parameter is auto.
<a id="configure-backend-meson"></a>

<code>--backend=</code><em>BACKEND</em>
:   The default backend Meson uses is ninja and that should suffice for most use cases. However, if you'd like to fully integrate with Visual Studio, you can set the *BACKEND* to `vs`.
<a id="configure-c-args-meson"></a>

<code>-Dc_args=</code><em>OPTIONS</em>
:   This option can be used to pass extra options to the C compiler.
<a id="configure-c-link-args-meson"></a>

<code>-Dc_link_args=</code><em>OPTIONS</em>
:   This option can be used to pass extra options to the C linker.
<a id="configure-extra-include-dirs-meson"></a>

<code>-Dextra_include_dirs=</code><em>DIRECTORIES</em>
:   *DIRECTORIES* is a comma-separated list of directories that will be added to the list the compiler searches for header files. If you have optional packages (such as GNU Readline) installed in a non-standard location, you have to use this option and probably also the corresponding `-Dextra_lib_dirs` option.


     Example: `-Dextra_include_dirs=/opt/gnu/include,/usr/sup/include`.
<a id="configure-extra-lib-dirs-meson"></a>

<code>-Dextra_lib_dirs=</code><em>DIRECTORIES</em>
:   *DIRECTORIES* is a comma-separated list of directories to search for libraries. You will probably have to use this option (and the corresponding `-Dextra_include_dirs` option) if you have packages installed in non-standard locations.


     Example: `-Dextra_lib_dirs=/opt/gnu/lib,/usr/sup/lib`.
<a id="configure-system-tzdata-meson"></a>

<code>-Dsystem_tzdata=</code><em>DIRECTORY</em>
:   PostgreSQL includes its own time zone database, which it requires for date and time operations. This time zone database is in fact compatible with the IANA time zone database provided by many operating systems such as FreeBSD, Linux, and Solaris, so it would be redundant to install it again. When this option is used, the system-supplied time zone database in *DIRECTORY* is used instead of the one included in the PostgreSQL source distribution. *DIRECTORY* must be specified as an absolute path. `/usr/share/zoneinfo` is a likely directory on some operating systems. Note that the installation routine will not detect mismatching or erroneous time zone data. If you use this option, you are advised to run the regression tests to verify that the time zone data you have pointed to works correctly with PostgreSQL.


     This option is mainly aimed at binary package distributors who know their target operating system well. The main advantage of using this option is that the PostgreSQL package won't need to be upgraded whenever any of the many local daylight-saving time rules change. Another advantage is that PostgreSQL can be cross-compiled more straightforwardly if the time zone database files do not need to be built during the installation.
<a id="configure-extra-version-meson"></a>

<code>-Dextra_version=</code><em>STRING</em>
:   Append *STRING* to the PostgreSQL version number. You can use this, for example, to mark binaries built from unreleased Git snapshots or containing custom patches with an extra version string, such as a `git describe` identifier or a distribution package release number.
<a id="configure-rpath-meson"></a>

`-Drpath={ true | false }`
:   This option is set to true by default. If set to false, do not mark PostgreSQL's executables to indicate that they should search for shared libraries in the installation's library directory (see `--libdir`). On most platforms, this marking uses an absolute path to the library directory, so that it will be unhelpful if you relocate the installation later. However, you will then need to provide some other way for the executables to find the shared libraries. Typically this requires configuring the operating system's dynamic linker to search the library directory; see [Shared Libraries](post-installation-setup.md#install-post-shlibs) for more detail.
<a id="configure-binary-name-meson"></a>

<code>-D</code><em>BINARY_NAME</em><code>=</code><em>PATH</em>
:   If a program required to build PostgreSQL (with or without optional flags) is stored at a non-standard path, you can specify it manually to `meson configure`. The complete list of programs for which this is supported can be found by running `meson configure`. Example:

    ```
    meson configure -DBISON=PATH_TO_BISON
    ```
  <a id="meson-options-docs"></a>

#### Documentation


 See [Tool Sets](../../appendixes/documentation/tool-sets.md#docguide-toolsets) for the tools needed for building the documentation.


<a id="configure-docs-meson"></a>

`-Ddocs={ auto | enabled | disabled }`
:   Enables building the documentation in HTML and man format. It defaults to auto.
<a id="configure-docs-pdf-meson"></a>

`-Ddocs_pdf={ auto | enabled | disabled }`
:   Enables building the documentation in PDF format. It defaults to auto.
<a id="configure-docs-html-style"></a>

`-Ddocs_html_style={ simple | website }`
:   Controls which CSS stylesheet is used. The default is `simple`. If set to `website`, the HTML documentation will reference the stylesheet for [postgresql.org](https://www.postgresql.org/docs/current/).
  <a id="meson-options-misc"></a>

#### Miscellaneous


<a id="configure-pgport-meson"></a>

<code>-Dpgport=</code><em>NUMBER</em>
:   Set *NUMBER* as the default port number for server and clients. The default is 5432. The port can always be changed later on, but if you specify it here then both server and clients will have the same default compiled in, which can be very convenient. Usually the only good reason to select a non-default value is if you intend to run multiple PostgreSQL servers on the same machine.
<a id="configure-krb-srvnam-meson"></a>

<code>-Dkrb_srvnam=</code><em>NAME</em>
:   The default name of the Kerberos service principal used by GSSAPI. `postgres` is the default. There's usually no reason to change this unless you are building for a Windows environment, in which case it must be set to upper case `POSTGRES`.
<a id="configure-segsize-meson"></a>

<code>-Dsegsize=</code><em>SEGSIZE</em>
:   Set the *segment size*, in gigabytes. Large tables are divided into multiple operating-system files, each of size equal to the segment size. This avoids problems with file size limits that exist on many platforms. The default segment size, 1 gigabyte, is safe on all supported platforms. If your operating system has “largefile” support (which most do, nowadays), you can use a larger segment size. This can be helpful to reduce the number of file descriptors consumed when working with very large tables. But be careful not to select a value larger than is supported by your platform and the file systems you intend to use. Other tools you might wish to use, such as tar, could also set limits on the usable file size. It is recommended, though not absolutely required, that this value be a power of 2.
<a id="configure-blocksize-meson"></a>

<code>-Dblocksize=</code><em>BLOCKSIZE</em>
:   Set the *block size*, in kilobytes. This is the unit of storage and I/O within tables. The default, 8 kilobytes, is suitable for most situations; but other values may be useful in special cases. The value must be a power of 2 between 1 and 32 (kilobytes).
<a id="configure-wal-blocksize-meson"></a>

<code>-Dwal_blocksize=</code><em>BLOCKSIZE</em>
:   Set the *WAL block size*, in kilobytes. This is the unit of storage and I/O within the WAL log. The default, 8 kilobytes, is suitable for most situations; but other values may be useful in special cases. The value must be a power of 2 between 1 and 64 (kilobytes).
  <a id="meson-options-devel"></a>

#### Developer Options


 Most of the options in this section are only of interest for developing or debugging PostgreSQL. They are not recommended for production builds, except for `--debug`, which can be useful to enable detailed bug reports in the unlucky event that you encounter a bug. On platforms supporting DTrace, `-Ddtrace` may also be reasonable to use in production.


 When building an installation that will be used to develop code inside the server, it is recommended to use at least the `--buildtype=debug` and `-Dcassert` options.


<a id="configure-buildtype-meson"></a>

<code>--buildtype=</code><em>BUILDTYPE</em>
:   This option can be used to specify the buildtype to use; defaults to `debugoptimized`. If you'd like finer control on the debug symbols and optimization levels than what this option provides, you can refer to the `--debug` and `--optimization` flags.


     The following build types are generally used: `plain`, `debug`, `debugoptimized` and `release`. More information about them can be found in the [Meson documentation](https://mesonbuild.com/Running-Meson.html#configuring-the-build-directory).
<a id="configure-debug-meson"></a>

`--debug`
:   Compiles all programs and libraries with debugging symbols. This means that you can run the programs in a debugger to analyze problems. This enlarges the size of the installed executables considerably, and on non-GCC compilers it usually also disables compiler optimization, causing slowdowns. However, having the symbols available is extremely helpful for dealing with any problems that might arise. Currently, this option is recommended for production installations only if you use GCC. But you should always have it on if you are doing development work or running a beta version.
<a id="configure-optimization-meson"></a>

`--optimization`=*LEVEL*
:   Specify the optimization level. `LEVEL` can be set to any of {0,g,1,2,3,s}.
<a id="configure-werror-meson"></a>

`--werror`
:   Setting this option asks the compiler to treat warnings as errors. This can be useful for code development.
<a id="configure-cassert-meson"></a>

`-Dcassert={ true | false }`
:   Enables *assertion* checks in the server, which test for many “cannot happen” conditions. This is invaluable for code development purposes, but the tests slow down the server significantly. Also, having the tests turned on won't necessarily enhance the stability of your server! The assertion checks are not categorized for severity, and so what might be a relatively harmless bug will still lead to server restarts if it triggers an assertion failure. This option is not recommended for production use, but you should have it on for development work or when running a beta version.
<a id="configure-tap-tests-meson"></a>

`-Dtap_tests={ auto | enabled | disabled }`
:   Enable tests using the Perl TAP tools. Defaults to auto and requires a Perl installation and the Perl module `IPC::Run`. See [TAP Tests](../regression-tests/tap-tests.md#regress-tap) for more information.
<a id="configure-pg-test-extra-meson"></a>

<code>-DPG_TEST_EXTRA=</code><em>TEST_SUITES</em>
:   Enable additional test suites, which are not run by default because they are not secure to run on a multiuser system, require special software to run, or are resource intensive. The argument is a whitespace-separated list of tests to enable. See [Additional Test Suites](../regression-tests/running-the-tests.md#regress-additional) for details. If the `PG_TEST_EXTRA` environment variable is set when the tests are run, it overrides this setup-time option.
<a id="configure-b-coverage-meson"></a>

`-Db_coverage={ true | false }`
:   If using GCC, all programs and libraries are compiled with code coverage testing instrumentation. When run, they generate files in the build directory with code coverage metrics. See [Test Coverage Examination](../regression-tests/test-coverage-examination.md#regress-coverage) for more information. This option is for use only with GCC and when doing development work.
<a id="configure-dtrace-meson"></a>

`-Ddtrace={ auto | enabled | disabled }`
:   Enabling this compiles PostgreSQL with support for the dynamic tracing tool DTrace. See [Dynamic Tracing](../monitoring-database-activity/dynamic-tracing.md#dynamic-trace) for more information.


     To point to the `dtrace` program, the `DTRACE` option can be set. This will often be necessary because `dtrace` is typically installed under `/usr/sbin`, which might not be in your `PATH`.
<a id="configure-injection-points-meson"></a>

`-Dinjection_points={ true | false }`
:   Compiles PostgreSQL with support for injection points in the server. Injection points allow to run user-defined code from within the server in pre-defined code paths. This helps in testing and in the investigation of concurrency scenarios in a controlled fashion. This option is disabled by default. See [Injection Points](../../server-programming/extending-sql/c-language-functions.md#xfunc-addin-injection-points) for more details. This option is intended to be used only by developers for testing.
<a id="configure-segsize-blocks-meson"></a>

`-Dsegsize_blocks=SEGSIZE_BLOCKS`
:   Specify the relation segment size in blocks. If both `-Dsegsize` and this option are specified, this option wins. This option is only for developers, to test segment related code.
   <a id="targets-meson"></a>

### `meson` Build Targets


 Individual build targets can be built using `ninja` *target*. When no target is specified, everything except documentation is built. Individual build products can be built using the path/filename as *target*.
  <a id="targets-meson-code"></a>

#### Code Targets


<a id="meson-target-all"></a>

`all`
:   Build everything other than documentation
<a id="meson-target-backend"></a>

`backend`
:   Build backend and related modules
<a id="meson-target-bin"></a>

`bin`
:   Build frontend binaries
<a id="meson-target-contrib"></a>

`contrib`
:   Build contrib modules
<a id="meson-target-pl"></a>

`pl`
:   Build procedural languages
  <a id="targets-meson-developer"></a>

#### Developer Targets


<a id="meson-target-reformat-dat-files"></a>

`reformat-dat-files`
:   Rewrite catalog data files into standard format
<a id="meson-target-expand-dat-files"></a>

`expand-dat-files`
:   Expand all data files to include defaults
<a id="meson-target-update-unicode"></a>

`update-unicode`
:   Update unicode data to new version
  <a id="targets-meson-documentation"></a>

#### Documentation Targets


<a id="meson-target-html"></a>

`html`
:   Build documentation in multi-page HTML format
<a id="meson-target-man"></a>

`man`
:   Build documentation in man page format
<a id="meson-target-docs"></a>

`docs`
:   Build documentation in multi-page HTML and man page format
<a id="meson-target-doc-src-sgml-postgres-A4.pdf"></a>

`doc/src/sgml/postgres-A4.pdf`
:   Build documentation in PDF format, with A4 pages
<a id="meson-target-doc-src-sgml-postgres-US.pdf"></a>

`doc/src/sgml/postgres-US.pdf`
:   Build documentation in PDF format, with US letter pages
<a id="meson-target-doc-src-sgml-postgres.html"></a>

`doc/src/sgml/postgres.html`
:   Build documentation in single-page HTML format
<a id="meson-target-alldocs"></a>

`alldocs`
:   Build documentation in all supported formats
  <a id="targets-meson-installation"></a>

#### Installation Targets


<a id="meson-target-install"></a>

`install`
:   Install postgres, excluding documentation
<a id="meson-target-install-docs"></a>

`install-docs`
:   Install documentation in multi-page HTML and man page formats
<a id="meson-target-install-html"></a>

`install-html`
:   Install documentation in multi-page HTML format
<a id="meson-target-install-man"></a>

`install-man`
:   Install documentation in man page format
<a id="meson-target-install-quiet"></a>

`install-quiet`
:   Like "install", but installed files are not displayed
<a id="meson-target-install-world"></a>

`install-world`
:   Install postgres, including multi-page HTML and man page documentation
<a id="meson-target-uninstall"></a>

`uninstall`
:   Remove installed files
  <a id="targets-meson-other"></a>

#### Other Targets


<a id="meson-target-clean"></a>

`clean`
:   Remove all build products
<a id="meson-target-test"></a>

`test`
:   Run all enabled tests (including contrib)
<a id="meson-target-world"></a>

`world`
:   Build everything, including documentation
<a id="meson-target-help"></a>

`help`
:   List important targets
