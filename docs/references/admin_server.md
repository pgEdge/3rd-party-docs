<a id="admin_server"></a>

# Admin Server

PostgREST provides an admin server that can be enabled by setting [admin-server-port](configuration.md#admin-server-port).
<a id="health_check"></a>

## Health Check

You can enable a health check to verify if PostgREST is available for client requests. Also to check the status of its internal state.

Two endpoints `live` and `ready` will then be available. Both these endpoints reply with a status code and empty response body.

!!! important

    If you have a machine with multiple network interfaces and multiple PostgREST instances in the same port, you need to specify a unique [hostname](configuration.md#server-host) in the configuration of each PostgREST instance for the health check to work correctly. Don't use the special values(`!4`, `*`, etc) in this case because the health check could report a false positive.

### Live

The `live` endpoint verifies if PostgREST is running on its configured port. A request will return `200 OK` if PostgREST is alive or `500` otherwise.

For instance, to verify if PostgREST is running while the `admin-server-port` is set to `3001`:

```bash
curl -I "http://localhost:3001/live"
```

```http
HTTP/1.1 200 OK
```

### Ready

Additionally to the `live` check, the `ready` endpoint checks the state of the [Connection Pool](connection_pool.md#connection_pool) and the [Schema Cache](schema_cache.md#schema_cache). A request will return `200 OK` if both are good or `503` if not.

```bash
curl -I "http://localhost:3001/ready"
```

```http
HTTP/1.1 200 OK
```

PostgREST will try to recover from the `503` state with [Automatic Recovery](connection_pool.md#automatic_recovery).

## Metrics

Provides [Metrics](observability.md#metrics).

## Runtime Schema Cache

Provides the `schema_cache` endpoint that prints the runtime [Schema Cache](schema_cache.md#schema_cache).

```bash
curl "http://localhost:3001/schema_cache"
```

```json
{
  "dbMediaHandlers": ["..."],
  "dbRelationships": ["..."],
  "dbRepresentations": ["..."],
  "dbRoutines": ["..."],
  "dbTables": ["..."],
  "dbTimezones": ["..."]
}
```
