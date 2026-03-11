## Building the Documentation with Meson { #docguide-build-meson }


 To build the documentation using Meson, change to the `build` directory before running one of these commands, or add `-C build` to the command.


 To build just the HTML version of the documentation:

```

build$ ninja html
```
 For a list of other documentation targets see [Documentation Targets](../../server-administration/installation-from-source-code/building-and-installation-with-meson.md#targets-meson-documentation). The output appears in the subdirectory `build/doc/src/sgml`.
