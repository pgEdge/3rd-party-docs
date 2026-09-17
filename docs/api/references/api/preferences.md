<a id="preferences"></a>

# Prefer Header

PostgREST honors the Prefer HTTP header specified on [RFC 7240](https://www.rfc-editor.org/rfc/rfc7240.html). It allows clients to specify required and optional behaviors for their requests.

The following preferences are supported.

- `Prefer: handling`. See [Strict or Lenient Handling](#prefer_handling).

- `Prefer: timezone`. See [Timezone](#prefer_timezone).

- `Prefer: return`. See [Return Representation](#prefer_return).

- `Prefer: count`. See [Counting](pagination_count.md#prefer_count).

- `Prefer: resolution`. See [prefer_resolution](tables_views.md#prefer_resolution).

- `Prefer: missing`. See [Missing](#prefer_missing).

- `Prefer: max-affected`, See [Max Affected](#prefer_max_affected).

- `Prefer: tx`. See [Transaction End Preference](#prefer_tx).
<a id="prefer_handling"></a>

## Strict or Lenient Handling

The server ignores unrecognized or unfulfillable preferences by default. You can control this behavior with the `handling` preference. It can take two values: `lenient` (the default) or `strict`.

`handling=strict` will throw an error if you specify invalid preferences. For instance:

```bash
curl -i "http://localhost:3000/projects" \
  -H "Prefer: handling=strict, foo, bar"
```

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
```

```json
{
    "code": "PGRST122",
    "message": "Invalid preferences given with handling=strict",
    "details": "Invalid preferences: foo, bar",
    "hint": null
}
```

`handling=lenient` ignores invalid preferences.

```bash
curl -i "http://localhost:3000/projects" \
  -H "Prefer: handling=lenient, foo, bar"
```

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
```
<a id="prefer_timezone"></a>

## Timezone

The `timezone` preference allows you to change the [PostgreSQL timezone](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-TIMEZONE). It accepts all time zones in [pg_timezone_names](https://www.postgresql.org/docs/current/view-pg-timezone-names.html).

```bash
curl -i "http://localhost:3000/timestamps" \
  -H "Prefer: timezone=America/Los_Angeles"
```

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Preference-Applied: timezone=America/Los_Angeles
```

```json
[
  {"t":"2023-10-18T05:37:59.611-07:00"},
  {"t":"2023-10-18T07:37:59.611-07:00"},
  {"t":"2023-10-18T09:37:59.611-07:00"}
]
```

For an invalid time zone, PostgREST returns values with the default time zone (configured on `postgresql.conf` or as a setting on the [authenticator](../../../references/auth.md#roles)).

```bash
curl -i "http://localhost:3000/timestamps" \
  -H "Prefer: timezone=Jupiter/Red_Spot"
```

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
```

```json
[
  {"t":"2023-10-18T12:37:59.611+00:00"},
  {"t":"2023-10-18T14:37:59.611+00:00"},
  {"t":"2023-10-18T16:37:59.611+00:00"}
]
```

Note that there's no `Preference-Applied` in the response.

However, with `handling=strict`, an invalid time zone preference will throw an error.

```bash
curl -i "http://localhost:3000/timestamps" \
  -H "Prefer: handling=strict, timezone=Jupiter/Red_Spot"
```

```http
HTTP/1.1 400 Bad Request
```
<a id="prefer_return"></a>

## Return Representation

The `return` preference can be used to obtain information about affected resource when it's [inserted](tables_views.md#insert), [updated](tables_views.md#update) or [deleted](tables_views.md#delete). This helps avoid a subsequent GET request.

### Minimal

With `Prefer: return=minimal`, no response body will be returned. This is the default mode for all write requests.

### Headers Only

If the table has a primary key, the response can contain a `Location` header describing where to find the new object by including the header `Prefer: return=headers-only` in the request. Make sure that the table is not write-only, otherwise constructing the `Location` header will cause a permissions error.

```bash
curl -i "http://localhost:3000/projects" -X POST \
  -H "Content-Type: application/json" \
  -H "Prefer: return=headers-only" \
  -d '{"id":33, "name": "x"}'
```

```http
HTTP/1.1 201 Created
Location: /projects?id=eq.34
Preference-Applied: return=headers-only
```

### Full

On the other end of the spectrum you can get the full created object back in the response to your request by including the header `Prefer: return=representation`. That way you won't have to make another HTTP call to discover properties that may have been filled in on the server side. You can also apply the standard [Vertical Filtering](tables_views.md#v_filter) to these results.

```bash
curl -i "http://localhost:3000/projects" -X POST \
  -H "Content-Type: application/json" \
  -H "Prefer: return=representation" \
  -d '{"id":33, "name": "x"}'
```

```http
HTTP/1.1 201 Created
Preference-Applied: return=representation
```

```json
[
    {
        "id": 33,
        "name": "x"
    }
]
```
<a id="prefer_tx"></a>

## Transaction End Preference

The `tx` preference can be set to specify if the [transaction](../../../references/transactions.md#transactions) will end in a COMMIT or ROLLBACK. This preference is not enabled by default but can be activated with [db-tx-end](../../../references/configuration.md#db-tx-end).

```bash
curl -i "http://localhost:3000/projects" -X POST \
  -H "Content-Type: application/json" \
  -H "Prefer: tx=rollback, return=representation" \
  -d '{"name": "Project X"}'
```

```http
HTTP/1.1 200 OK
Preference-Applied: tx=rollback, return=representation

{"id": 35, "name": "Project X"}
```
<a id="prefer_missing"></a>

## Missing

When doing `POST` and `PATCH` requests, any missing columns in the payload will be inserted as `null` value by default. To use the `DEFAULT` column value instead, use the `Prefer: missing=default` header.

Having:

```postgres
create table foo (
  id bigint generated by default as identity primary key
, bar text
, baz int default 100
);
```

A request:

```bash
curl "http://localhost:3000/foo?columns=id,bar,baz" \
  -H "Content-Type: application/json" \
  -H "Prefer: missing=default, return=representation" \
  -d @- << EOF
    [
      { "bar": "val1" },
      { "bar": "val2", "baz": 15 }
    ]
EOF
```

Will result in:

```json
[
  { "id":  1, "bar": "val1", "baz": 100 },
  { "id":  2, "bar": "val2", "baz": 15 }
]
```
<a id="prefer_max_affected"></a>

## Max Affected

You can set a limit to the amount of resources affected in a request by sending `max-affected` preference. This feature works in combination with `handling=strict` preference. `max-affected` would be ignored with lenient handling. The "affected resources" are the number of rows returned by `DELETE` and `PATCH` requests.

To illustrate the use of this preference, consider the following scenario where the `items` table contains 14 rows.

```bash
curl -i "http://localhost:3000/items?id=lt.15 -X DELETE \
  -H "Content-Type: application/json" \
  -H "Prefer: handling=strict, max-affected=10"
```

```http
HTTP/1.1 400 Bad Request
```

```json
{
    "code": "PGRST124",
    "message": "Query result exceeds max-affected preference constraint",
    "details": "The query affects 14 rows",
    "hint": null
}
```

With [RPC](functions.md#functions), the preference is honored completely on the basis of the number of rows returned in the result set of the function. This can be useful for complex mutation queries using [data-modifying statements](https://www.postgresql.org/docs/current/queries-with.html#QUERIES-WITH-MODIFYING). A simple example:

```postgres
CREATE FUNCTION test.delete_items()
RETURNS SETOF items AS $$
  DELETE FROM items WHERE id < 15 RETURNING *;
$$ LANGUAGE SQL;
```

```bash
curl -i "http://localhost:3000/rpc/delete_items" \
  -H "Content-Type: application/json" \
  -H "Prefer: handling=strict, max-affected=10"
```

```http
HTTP/1.1 400 Bad Request
```

```json
{
    "code": "PGRST124",
    "message": "Query result exceeds max-affected preference constraint",
    "details": "The query affects 14 rows",
    "hint": null
}
```

!!! note

    It is important for functions to return `SETOF` or `TABLE` when called with `max-affected` preference. A violation of this would cause a PGRST128 error.
