<a id="listener"></a>

# Listener

PostgREST uses [LISTEN](https://www.postgresql.org/docs/current/sql-listen.html) to reload its [Schema Cache](schema_cache.md#schema_reloading_notify) and [Configuration](configuration.md#config_reloading_notify) via [NOTIFY](https://www.postgresql.org/docs/current/sql-notify.html). This is useful in environments where you can't send SIGUSR1 or SIGUSR2 Unix Signals. Like on cloud managed containers or on Windows systems.

```postgresql
NOTIFY pgrst, 'reload schema'; -- reload schema cache
NOTIFY pgrst, 'reload config'; -- reload config
NOTIFY pgrst;                  -- reload both
```

By default, the LISTEN channel is enabled ([db-channel-enabled](configuration.md#db-channel-enabled)) and named `pgrst` ([db-channel](configuration.md#db-channel)).

## Listener on Read Replicas

The `LISTEN` and `NOTIFY` commands do not work on PostgreSQL read replicas. Thus, if you connect PostgREST to a read replica the Listener will fail to start.

```psql
-- check if the instance is a replica
postgres=# select pg_is_in_recovery();
 pg_is_in_recovery
-------------------
 t
(1 row)

postgres=# LISTEN pgrst;
ERROR:  cannot execute LISTEN during recovery
```

To work around this, you can connect the Listener to the primary while still using the [Connection Pool](connection_pool.md#connection_pool) on the replica.

This can be done by using the standard [libpq multiple hosts](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-MULTIPLE-HOSTS) and [target_session_attrs](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNECT-TARGET-SESSION-ATTRS) in your [connection string](configuration.md#db-uri).

```bash
db-uri = "postgres://read_replica.host,primary.host/mydb?target_session_attrs=read-only"
```

This will cause the [Connection Pool](connection_pool.md#connection_pool) to connect to the read replica host and `LISTEN` on the fallback primary host.

!!! note

    Under the hood, PostgREST forces [target_session_attrs=read-write](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNECT-TARGET-SESSION-ATTRS) for the `LISTEN` session.
<a id="listener_automatic_recovery"></a>

## Automatic Recovery

The listener will retry reconnecting to the database if connection loss happens.

- It will retry forever with exponential backoff, with a maximum backoff time of 32 seconds between retries. Each of these attempts are [logged](observability.md#pgrst_logging).

- Automatic recovery can be disabled by setting [db-pool-automatic-recovery](configuration.md#db-pool-automatic-recovery) to `false`.

- To ensure a valid state, the listener reloads the [Schema Cache](schema_cache.md#schema_cache) and [Configuration](configuration.md#configuration) when recovering.
