# Installation Notes - Linux and Mac

## Postgres Location

If your machine has multiple Postgres installations, specify the path to [pg_config](https://www.postgresql.org/docs/current/app-pgconfig.html) with:

```sh
export PG_CONFIG=/Library/PostgreSQL/18/bin/pg_config
```

Then re-run the installation instructions (run `make clean` before `make` if needed). If `sudo` is needed for `make install`, use:

```sh
sudo --preserve-env=PG_CONFIG make install
```

A few common paths on Mac are:

- EDB installer - `/Library/PostgreSQL/18/bin/pg_config`
- Homebrew (arm64) - `/opt/homebrew/opt/postgresql@18/bin/pg_config`
- Homebrew (x86-64) - `/usr/local/opt/postgresql@18/bin/pg_config`

Note: Replace `18` with your Postgres server version

## Missing Header

If compilation fails with `fatal error: postgres.h: No such file or directory`, make sure Postgres development files are installed on the server.

For Ubuntu and Debian, use:

```sh
sudo apt install postgresql-server-dev-18
```

Note: Replace `18` with your Postgres server version

## Missing SDK

If compilation fails and the output includes `warning: no such sysroot directory` on Mac, your Postgres installation points to a path that no longer exists.

```sh
pg_config --cppflags
```

Reinstall Postgres to fix this.

## Portability

By default, pgvector compiles with `-march=native` on some platforms for best performance. However, this can lead to `Illegal instruction` errors if trying to run the compiled extension on a different machine.

To compile for portability, use:

```sh
make OPTFLAGS=""
```

