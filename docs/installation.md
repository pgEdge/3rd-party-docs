# Installation

## Linux and Mac

Compile and install the extension (supports Postgres 13+)

```sh
cd /tmp
git clone --branch v0.8.1 https://github.com/pgvector/pgvector.git
cd pgvector
make
make install # may need sudo
```

See the [installation notes](installation-notes-linux-and-mac.md) if you run into issues

You can also install it with [Docker](additional-installation-methods.md#docker), [Homebrew](additional-installation-methods.md#homebrew), [PGXN](additional-installation-methods.md#pgxn), [APT](additional-installation-methods.md#apt), [Yum](additional-installation-methods.md#yum), [pkg](additional-installation-methods.md#pkg), or [conda-forge](additional-installation-methods.md#conda-forge), and it comes preinstalled with [Postgres.app](additional-installation-methods.md#postgresapp) and many [hosted providers](hosted-postgres.md). There are also instructions for [GitHub Actions](https://github.com/pgvector/setup-pgvector).

## Windows

Ensure [C++ support in Visual Studio](https://learn.microsoft.com/en-us/cpp/build/building-on-the-command-line?view=msvc-170#download-and-install-the-tools) is installed and run `x64 Native Tools Command Prompt for VS [version]` as administrator. Then use `nmake` to build:

```cmd
set "PGROOT=C:\Program Files\PostgreSQL\17"
cd %TEMP%
git clone --branch v0.8.1 https://github.com/pgvector/pgvector.git
cd pgvector
nmake /F Makefile.win
nmake /F Makefile.win install
```

See the [installation notes](installation-notes-windows.md) if you run into issues

You can also install it with [Docker](additional-installation-methods.md#docker) or [conda-forge](additional-installation-methods.md#conda-forge).

