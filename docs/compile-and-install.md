# Compile and Install

pgAudit can be compiled against an installed copy of PostgreSQL with development packages using `PGXS`.

The following instructions are for RHEL 7.

Clone the pgAudit extension:
```
git clone https://github.com/pgaudit/pgaudit.git
```
Change to pgAudit directory:
```
cd pgaudit
```
Checkout `REL_17_STABLE` branch (note that the stable branch may not exist for unreleased versions of PostgreSQL):
```
git checkout REL_17_STABLE
```
Build and install pgAudit:
```
make install USE_PGXS=1 PG_CONFIG=/usr/pgsql-17/bin/pg_config
```
Instructions for testing and development may be found in `test`.

