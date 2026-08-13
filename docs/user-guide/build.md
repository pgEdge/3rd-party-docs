# Build
<a name="build"></a>

Installing pgBackRest from a package is preferable to building from source. See [Installation](multi-stanza.md#installation) for more information about packages.

When building from source it is best to use a build host rather than building on production. Many of the tools required for the build should generally not be installed in production. pgBackRest consists of a single executable so it is easy to copy to a new host once it is built.

**Download version 2.59.0 of  to /build path**

```bash
mkdir -p /build
```

```bash
curl -fsSL
                    https://github.com/pgbackrest/pgbackrest/releases/download/release%2F2.59.0/pgbackrest-2.59.0.tar.gz |
                    tar zx -C /build
```

**Install build dependencies**

```bash
apt-get install python3-distutils meson gcc libpq-dev libssl-dev libxml2-dev
                    pkg-config liblz4-dev libzstd-dev libbz2-dev libz-dev libssh2-1-dev libsystemd-dev
```

```bash
yum install meson gcc postgresql17-devel openssl-devel libxml2-devel
                    lz4-devel libzstd-devel bzip2-devel libssh2-devel systemd-devel
```

**Configure and compile**

```bash
meson setup /build/pgbackrest /build/pgbackrest-2.59.0
```

```bash
ninja -C /build/pgbackrest
```

**Optionally run smoke tests to verify  was built correctly**

```bash
meson test -C /build/pgbackrest --suite smoke
```
