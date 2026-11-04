# Monitoring
<a name="monitor"></a>

Monitoring is an important part of any production system. There are many tools available and pgBackRest can be monitored on any of them with a little work.

pgBackRest can output information about the repository in JSON format which includes a list of all backups for each stanza and WAL archive info.

## In
<a name="postgresql"></a>

The PostgreSQL `COPY` command allows pgBackRest info to be loaded into a table. The following example wraps that logic in a function that can be used to perform real-time queries.

**Load  info function for**

```bash
cat
                        /var/lib/postgresql/pgbackrest/doc/example/pgsql-pgbackrest-info.sql
```

```bash
psql -f
                        /var/lib/postgresql/pgbackrest/doc/example/pgsql-pgbackrest-info.sql
```

Now the `monitor.pgbackrest_info()` function can be used to determine the last successful backup time and archived WAL for a stanza.

**Query last successful backup time and archived WAL**

```bash
cat
                        /var/lib/postgresql/pgbackrest/doc/example/pgsql-pgbackrest-query.sql
```

```bash
psql -f
                        /var/lib/postgresql/pgbackrest/doc/example/pgsql-pgbackrest-query.sql
```

## Using jq
<a name="jq"></a>

jq is a command-line utility that can easily extract data from JSON.

**Install jq utility**

```bash
apt-get install jq
```

Now jq can be used to query the last successful backup time for a stanza.

**Query last successful backup time**

```bash
pgbackrest --output=json --stanza=demo info |
                          jq '.[0] | .backup[-1] | .timestamp.stop'
```

Or the last archived WAL.

**Query last archived WAL**

```bash
pgbackrest --output=json --stanza=demo info |
                          jq '.[0] | .archive[-1] | .max'
```

!!! note

    This syntax requires jq v1.5.

!!! note

    jq may round large numbers such as system identifiers. Test your queries carefully.
