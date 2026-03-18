# Additional Installation Methods

## Docker

Get the [Docker image](https://hub.docker.com/r/pgvector/pgvector) with:

```sh
docker pull pgvector/pgvector:pg17-trixie
```

This adds pgvector to the [Postgres image](https://hub.docker.com/_/postgres) (replace `17` with your Postgres server version, and run it the same way).

Supported tags are:

- `pg17-trixie`, `0.8.1-pg17-trixie`
- `pg17-bookworm`, `0.8.1-pg17-bookworm`, `pg17`, `0.8.1-pg17`
- `pg16-trixie`, `0.8.1-pg16-trixie`
- `pg16-bookworm`, `0.8.1-pg16-bookworm`, `pg16`, `0.8.1-pg16`
- `pg15-trixie`, `0.8.1-pg15-trixie`
- `pg15-bookworm`, `0.8.1-pg15-bookworm`, `pg15`, `0.8.1-pg15`
- `pg14-trixie`, `0.8.1-pg14-trixie`
- `pg14-bookworm`, `0.8.1-pg14-bookworm`, `pg14`, `0.8.1-pg14`
- `pg13-trixie`, `0.8.1-pg13-trixie`
- `pg13-bookworm`, `0.8.1-pg13-bookworm`, `pg13`, `0.8.1-pg13`

You can also build the image manually:

```sh
git clone --branch v0.8.1 https://github.com/pgvector/pgvector.git
cd pgvector
docker build --pull --build-arg PG_MAJOR=17 -t myuser/pgvector .
```

If you increase `maintenance_work_mem`, make sure `--shm-size` is at least that size to avoid an error with parallel HNSW index builds.

```sh
docker run --shm-size=1g ...
```

## Homebrew

With Homebrew Postgres, you can use:

```sh
brew install pgvector
```

Note: This only adds it to the `postgresql@17` and `postgresql@14` formulas

## PGXN

Install from the [PostgreSQL Extension Network](https://pgxn.org/dist/vector) with:

```sh
pgxn install vector
```

## APT

Debian and Ubuntu packages are available from the [PostgreSQL APT Repository](https://wiki.postgresql.org/wiki/Apt). Follow the [setup instructions](https://wiki.postgresql.org/wiki/Apt#Quickstart) and run:

```sh
sudo apt install postgresql-17-pgvector
```

Note: Replace `17` with your Postgres server version

## Yum

RPM packages are available from the [PostgreSQL Yum Repository](https://yum.postgresql.org/). Follow the [setup instructions](https://www.postgresql.org/download/linux/redhat/) for your distribution and run:

```sh
sudo yum install pgvector_17
# or
sudo dnf install pgvector_17
```

Note: Replace `17` with your Postgres server version

## pkg

Install the FreeBSD package with:

```sh
pkg install postgresql17-pgvector
```

or the port with:

```sh
cd /usr/ports/databases/pgvector
make install
```

## conda-forge

With Conda Postgres, install from [conda-forge](https://anaconda.org/conda-forge/pgvector) with:

```sh
conda install -c conda-forge pgvector
```

This method is [community-maintained](https://github.com/conda-forge/pgvector-feedstock) by [@mmcauliffe](https://github.com/mmcauliffe)

## Postgres.app

Download the [latest release](https://postgresapp.com/downloads.html) with Postgres 15+.

