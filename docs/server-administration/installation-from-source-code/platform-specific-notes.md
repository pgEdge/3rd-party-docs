<a id="installation-platform-notes"></a>

## Platform-Specific Notes


 This section documents additional platform-specific issues regarding the installation and setup of PostgreSQL. Be sure to read the installation instructions, and in particular [Requirements](requirements.md#install-requirements) as well. Also, check [Regression Tests](../regression-tests/index.md#regress) regarding the interpretation of regression test results.


 Platforms that are not covered here have no known platform-specific installation issues.
 <a id="installation-notes-aix"></a>

### AIX


 You must use GCC to build PostgreSQL on AIX. The native IBM compiler `xlc` is not supported.


 Also, only 64-bit builds are supported. While the make-based build system will automatically create 64-bit executables and libraries, the meson build system requires you to set `OBJECT_MODE` before building:

```

export OBJECT_MODE=64
meson setup ...
```
 Failure to do that will usually manifest as complaints from ar about files having the wrong object file mode.


 AIX versions before 7.2 are no longer tested nor supported by the PostgreSQL community.
  <a id="installation-notes-cygwin"></a>

### Cygwin


 PostgreSQL can be built using Cygwin, a Linux-like environment for Windows, but that method is inferior to the native Windows build and running a server under Cygwin is no longer recommended.


 When building from source, proceed according to the Unix-style installation procedure (i.e., `./configure; make`; etc.), noting the following Cygwin-specific differences:

-  Set your path to use the Cygwin bin directory before the Windows utilities. This will help prevent problems with compilation.
-  The `adduser` command is not supported; use the appropriate user management application on Windows. Otherwise, skip this step.
-  The `su` command is not supported; use ssh to simulate su on Windows. Otherwise, skip this step.
-  OpenSSL is not supported.
-  Start `cygserver` for shared memory support. To do this, enter the command `/usr/sbin/cygserver &`. This program needs to be running anytime you start the PostgreSQL server or initialize a database cluster (`initdb`). The default `cygserver` configuration may need to be changed (e.g., increase `SEMMNS`) to prevent PostgreSQL from failing due to a lack of system resources.
-  Building might fail on some systems where a locale other than C is in use. To fix this, set the locale to C by doing `export LANG=C.utf8` before building, and then setting it back to the previous setting after you have installed PostgreSQL.
-  The parallel regression tests (`make check`) can generate spurious regression test failures due to overflowing the `listen()` backlog queue which causes connection refused errors or hangs. You can limit the number of connections using the make variable `MAX_CONNECTIONS` thus:

```

make MAX_CONNECTIONS=5 check
```
   (On some systems you can have up to about 10 simultaneous connections.)


 It is possible to install `cygserver` and the PostgreSQL server as Windows NT services. For information on how to do this, please refer to the `README` document included with the PostgreSQL binary package on Cygwin. It is installed in the directory `/usr/share/doc/Cygwin`.
  <a id="installation-notes-macos"></a>

### macOS


 To build PostgreSQL from source on macOS, you will need to install Apple's command line developer tools, which can be done by issuing

```

xcode-select --install
```
 (note that this will pop up a GUI dialog window for confirmation). You may or may not wish to also install Xcode.


 On recent macOS releases, it's necessary to embed the “sysroot” path in the include switches used to find some system header files. This results in the outputs of the configure script varying depending on which SDK version was used during configure. That shouldn't pose any problem in simple scenarios, but if you are trying to do something like building an extension on a different machine than the server code was built on, you may need to force use of a different sysroot path. To do that, set `PG_SYSROOT`, for example

```

make PG_SYSROOT=/DESIRED/PATH all
```
 To find out the appropriate path on your machine, run

```

xcrun --show-sdk-path
```
 Note that building an extension using a different sysroot version than was used to build the core server is not really recommended; in the worst case it could result in hard-to-debug ABI inconsistencies.


 You can also select a non-default sysroot path when configuring, by specifying `PG_SYSROOT` to configure:

```

./configure ... PG_SYSROOT=/DESIRED/PATH
```
 This would primarily be useful to cross-compile for some other macOS version. There is no guarantee that the resulting executables will run on the current host.


 To suppress the `-isysroot` options altogether, use

```

./configure ... PG_SYSROOT=none
```
 (any nonexistent pathname will work). This might be useful if you wish to build with a non-Apple compiler, but beware that that case is not tested or supported by the PostgreSQL developers.


 macOS's “System Integrity Protection” (SIP) feature breaks `make check`, because it prevents passing the needed setting of `DYLD_LIBRARY_PATH` down to the executables being tested. You can work around that by doing `make install` before `make check`. Most PostgreSQL developers just turn off SIP, though.
  <a id="installation-notes-mingw"></a>

### MinGW


 PostgreSQL for Windows can be built using MinGW, a Unix-like build environment for Windows. It is recommended to use the [MSYS2](https://www.msys2.org/) environment for this and also to install any prerequisite packages.
 <a id="mingw-crash-dumps"></a>

#### Collecting Crash Dumps


 If PostgreSQL on Windows crashes, it has the ability to generate minidumps that can be used to track down the cause for the crash, similar to core dumps on Unix. These dumps can be read using the Windows Debugger Tools or using Visual Studio. To enable the generation of dumps on Windows, create a subdirectory named `crashdumps` inside the cluster data directory. The dumps will then be written into this directory with a unique name based on the identifier of the crashing process and the current time of the crash.
   <a id="installation-notes-solaris"></a>

### Solaris


 PostgreSQL is well-supported on Solaris. The more up to date your operating system, the fewer issues you will experience.
 <a id="installation-notes-solaris-req-tools"></a>

#### Required Tools


 Only GCC is supported as the compiler. Sun's compiler suite is no longer supported.


 Many additional dependencies can be installed via the package management system.
  <a id="installation-notes-solaris-configure-complains"></a>

#### configure Complains About a Failed Test Program


 If `configure` complains about a failed test program, this is probably a case of the run-time linker being unable to find some library, probably libz, libreadline or some other non-standard library such as libssl. To point it to the right location, set the `LDFLAGS` environment variable on the `configure` command line, e.g.,

```

configure ... LDFLAGS="-R /usr/sfw/lib:/opt/sfw/lib:/usr/local/lib"
```
 See the `ld`(1) man page for more information.
  <a id="installation-notes-solaris-using-dtrace"></a>

#### Using DTrace for Tracing PostgreSQL


 Yes, using DTrace is possible. See [Dynamic Tracing](../monitoring-database-activity/dynamic-tracing.md#dynamic-trace) for further information.
   <a id="installation-notes-visual-studio"></a>

### Visual Studio


 It is recommended that most users download the binary distribution for Windows, available as a graphical installer package from the PostgreSQL website at [https://www.postgresql.org/download/](https://www.postgresql.org/download/). Building from source is only intended for people developing PostgreSQL or extensions.


 PostgreSQL for Windows with Visual Studio can be built using Meson, as described in [Building and Installation with Meson](building-and-installation-with-meson.md#install-meson). The native Windows port requires a 32 or 64-bit version of Windows 10 or later.


 Native builds of psql don't support command line editing. The Cygwin build does support command line editing, so it should be used where psql is needed for interactive use on Windows.


 PostgreSQL can be built using the Visual C++ compiler suite from Microsoft. These compilers can be either from Visual Studio, Visual Studio Express or some versions of the Microsoft Windows SDK. If you do not already have a Visual Studio environment set up, the easiest ways are to use the compilers from Visual Studio 2022 or those in the Windows SDK 10, which are both free downloads from Microsoft.


 Both 32-bit and 64-bit builds are possible with the Microsoft Compiler suite. 32-bit PostgreSQL builds are possible with Visual Studio 2019 to Visual Studio 2022, as well as standalone Windows SDK releases 10 and above. 64-bit PostgreSQL builds are supported with Microsoft Windows SDK version 10 and above or Visual Studio 2019 and above.


 If your build environment doesn't ship with a supported version of the Microsoft Windows SDK it is recommended that you upgrade to the latest version (currently version 10), available for download from [https://www.microsoft.com/download](https://www.microsoft.com/download).


 You must always include the Windows Headers and Libraries part of the SDK. If you install a Windows SDK including the Visual C++ Compilers, you don't need Visual Studio to build. Note that as of Version 8.0a the Windows SDK no longer ships with a complete command-line build environment.
 <a id="windows-requirements"></a>

#### Requirements


 The following additional products are required to build PostgreSQL on Windows.

Strawberry Perl
:   Strawberry Perl is required to run the build generation scripts. MinGW or Cygwin Perl will not work. It must also be present in the PATH. Binaries can be downloaded from [https://strawberryperl.com](https://strawberryperl.com).

Bison and Flex
:   Binaries for Bison and Flex can be downloaded from [https://github.com/lexxmark/winflexbison](https://github.com/lexxmark/winflexbison).


 The following additional products are not required to get started, but are required to build the complete package.

Magicsplat Tcl
:   Required for building PL/Tcl. Binaries can be downloaded from [https://www.magicsplat.com/tcl-installer/index.html](https://www.magicsplat.com/tcl-installer/index.html).

Diff
:   Diff is required to run the regression tests, and can be downloaded from [http://gnuwin32.sourceforge.net](http://gnuwin32.sourceforge.net).

Gettext
:   Gettext is required to build with NLS support, and can be downloaded from [http://gnuwin32.sourceforge.net](http://gnuwin32.sourceforge.net). Note that binaries, dependencies and developer files are all needed.

MIT Kerberos
:   Required for GSSAPI authentication support. MIT Kerberos can be downloaded from [https://web.mit.edu/Kerberos/dist/index.html](https://web.mit.edu/Kerberos/dist/index.html).

libxml2 and libxslt
:   Required for XML support. Binaries can be downloaded from [https://zlatkovic.com/pub/libxml](https://zlatkovic.com/pub/libxml) or source from [http://xmlsoft.org](http://xmlsoft.org). Note that libxml2 requires iconv, which is available from the same download location.

LZ4
:   Required for supporting LZ4 compression. Binaries and source can be downloaded from [https://github.com/lz4/lz4/releases](https://github.com/lz4/lz4/releases).

Zstandard
:   Required for supporting Zstandard compression. Binaries and source can be downloaded from [https://github.com/facebook/zstd/releases](https://github.com/facebook/zstd/releases).

OpenSSL
:   Required for SSL support. Binaries can be downloaded from [https://slproweb.com/products/Win32OpenSSL.html](https://slproweb.com/products/Win32OpenSSL.html) or source from [https://www.openssl.org](https://www.openssl.org).

ossp-uuid
:   Required for UUID-OSSP support (contrib only). Source can be downloaded from [http://www.ossp.org/pkg/lib/uuid/](http://www.ossp.org/pkg/lib/uuid/).

Python
:   Required for building PL/Python. Binaries can be downloaded from [https://www.python.org](https://www.python.org).

zlib
:   Required for compression support in pg_dump and pg_restore. Binaries can be downloaded from [https://www.zlib.net](https://www.zlib.net).

  <a id="install-windows-full-64-bit"></a>

#### Special Considerations for 64-Bit Windows


 PostgreSQL will only build for the x64 and AArch64 architectures on 64-bit Windows.


 Mixing 32- and 64-bit versions in the same build tree is not supported. The build system will automatically detect if it's running in a 32- or 64-bit environment, and build PostgreSQL accordingly. For this reason, it is important to start the correct command prompt before building.


 To use a server-side third party library such as Python or OpenSSL, this library *must* also be 64-bit. There is no support for loading a 32-bit library in a 64-bit server. Several of the third party libraries that PostgreSQL supports may only be available in 32-bit versions, in which case they cannot be used with 64-bit PostgreSQL.
  <a id="windows-crash-dumps"></a>

#### Collecting Crash Dumps


 If PostgreSQL on Windows crashes, it has the ability to generate minidumps that can be used to track down the cause for the crash, similar to core dumps on Unix. These dumps can be read using the Windows Debugger Tools or using Visual Studio. To enable the generation of dumps on Windows, create a subdirectory named `crashdumps` inside the cluster data directory. The dumps will then be written into this directory with a unique name based on the identifier of the crashing process and the current time of the crash.
