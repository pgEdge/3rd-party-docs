# Installation from Source Code on Windows { #install-windows }

 It is recommended that most users download the binary distribution for Windows, available as a graphical installer package from the PostgreSQL website at [https://www.postgresql.org/download/](https://www.postgresql.org/download/). Building from source is only intended for people developing PostgreSQL or extensions.

 There are several different ways of building PostgreSQL on Windows. The simplest way to build with Microsoft tools is to install Visual Studio 2022 and use the included compiler. It is also possible to build with the full Microsoft Visual C++ 2015 to 2022. In some cases that requires the installation of the Windows SDK in addition to the compiler.

 It is also possible to build PostgreSQL using the GNU compiler tools provided by MinGW, or using Cygwin for older versions of Windows.

 Building using MinGW or Cygwin uses the normal build system, see [Installation from Source Code](../installation-from-source-code/index.md#installation) and the specific notes in [MinGW/Native Windows](../installation-from-source-code/platform-specific-notes.md#installation-notes-mingw) and [Cygwin](../installation-from-source-code/platform-specific-notes.md#installation-notes-cygwin). To produce native 64 bit binaries in these environments, use the tools from MinGW-w64. These tools can also be used to cross-compile for 32 bit and 64 bit Windows targets on other hosts, such as Linux and macOS. Cygwin is not recommended for running a production server, and it should only be used for running on older versions of Windows where the native build does not work. The official binaries are built using Visual Studio.

 Native builds of psql don't support command line editing. The Cygwin build does support command line editing, so it should be used where psql is needed for interactive use on Windows.

- [Building with Visual C++ or the Microsoft Windows SDK](building-with-visual-c-or-the-microsoft-windows-sdk.md#install-windows-full)
