# Build
<a name="build"></a>

Installing pgBackRest from a package is preferable to building from source. See [Installation](user-guide/multi-stanza.md#installation) for more information about packages.

When building from source it is best to use a build host rather than building on production. Many of the tools required for the build should generally not be installed in production. pgBackRest consists of a single executable so it is easy to copy to a new host once it is built.

**Download version v2.58.0 of  to /build path**

```bash
mkdir -p /build
```

```bash
wget -q -O -
                    https://github.com/pgbackrest/pgbackrest/archive/release/v2.58.0.tar.gz |
                    tar zx -C /build
```

**Install build dependencies**

```bash
apt-get install python3-distutils meson gcc libpq-dev libssl-dev libxml2-dev
                    pkg-config liblz4-dev libzstd-dev libbz2-dev libz-dev libyaml-dev libssh2-1-dev
```

```bash
yum install meson gcc postgresql<pg-version-nodot>-devel openssl-devel
                    libxml2-devel lz4-devel libzstd-devel bzip2-devel libyaml-devel libssh2-devel
```

**Configure and compile**

```bash
meson setup /build/pgbackrest /build/pgbackrest-release-v2.58.0
```

```bash
ninja -C /build/pgbackrest
```
