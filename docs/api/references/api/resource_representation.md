# Resource Representation

PostgREST uses proper HTTP content negotiation ([RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-5.3)) to deliver a resource representation. That is to say the same API endpoint can respond in different formats like JSON or CSV depending on the request.
<a id="res_format"></a>

## Response Format

Use the Accept request header to specify the acceptable format (or formats) for the response:

```bash
curl "http://localhost:3000/people" \
  -H "Accept: application/json"
```

!!! note

    The ordering of columns in the response isn't guaranteed to align with the order specified in the `select` clause. For example, with resource embedding:

    ```bash
    http://localhost:3000/films?select=directors(last_name,id),title
    ```

    We may get:

    ```bash
    [
      {
        "title": "title",
        "directors": {
          "id": 5,
          "last_name": "name"
        }
      }
    ]
    ```

    This is in line with the [JSON schema spec](https://json-schema.org/draft/2020-12/json-schema-core#name-instance-data-model):

    *"object: An unordered set of properties mapping a string to an instance"*
<a id="builtin_media"></a>

## Builtin Media Type Handlers

Builtin handlers are offered for common standard media types.

- `text/csv` and `application/json`, for all API endpoints. See [Tables and Views](tables_views.md#tables_views) and [Functions as RPC](functions.md#functions).

- `application/openapi+json`, for the root endpoint. See [OpenAPI](openapi.md#open-api).

- `application/geo+json`, see [PostGIS](../../../how-tos/working-with-postgresql-data-types.md#ww_postgis).

- `*/*`, resolves to `application/json` for API endpoints and to `application/openapi+json` for the root endpoint.

The following vendor media types handlers are also supported.

- `application/vnd.pgrst.plan`, see [Execution plan](../../../references/observability.md#explain_plan).

- `application/vnd.pgrst.object` and `application/vnd.pgrst.array`, see [Singular or Plural](#singular_plural) and [Stripped Nulls](#stripped_nulls).

Any unrecognized media type will throw an error.

```bash
curl "http://localhost:3000/people" \
  -H "Accept: unknown/unknown"
```

```http
HTTP/1.1 415 Unsupported Media Type

{"code":"PGRST107","details":null,"hint":null,"message":"None of these media types are available: unknown/unknown"}
```

To extend the accepted media types, you can use [Media Type Handlers](media_type_handlers.md#custom_media).
<a id="singular_plural"></a>

### Singular or Plural

By default PostgREST returns all JSON results in an array, even when there is only one item. For example, requesting `/items?id=eq.1` returns

```json
[
  { "id": 1 }
]
```

This can be inconvenient for client code. To return the first result as an object unenclosed by an array, specify `vnd.pgrst.object` as part of the `Accept` header

```bash
curl "http://localhost:3000/items?id=eq.1" \
  -H "Accept: application/vnd.pgrst.object+json"
```

This returns

```json
{ "id": 1 }
```

When a singular response is requested but no entries are found, the server responds with an error message and 406 Not Acceptable status code rather than the usual empty array and 200 status:

```json
{
  "code": "PGRST116",
  "message": "Cannot coerce the result to a single JSON object",
  "details": "The result contains 0 rows",
  "hint": null
}
```

!!! note

    Many APIs distinguish plural and singular resources using a special nested URL convention e.g. `/stories` vs `/stories/1`. Why do we use `/stories?id=eq.1`? The answer is because a singular resource is (for us) a row determined by a primary key, and primary keys can be compound (meaning defined across more than one column). The more familiar nested urls consider only a degenerate case of simple and overwhelmingly numeric primary keys. These so-called artificial keys are often introduced automatically by Object Relational Mapping libraries.

    Admittedly PostgREST could detect when there is an equality condition holding on all columns constituting the primary key and automatically convert to singular. However this could lead to a surprising change of format that breaks unwary client code just by filtering on an extra column. Instead we allow manually specifying singular vs plural to decouple that choice from the URL format.
<a id="stripped_nulls"></a>

### Stripped Nulls

By default PostgREST returns all JSON null values. For example, requesting `/projects?id=gt.10` returns

```json
[
  { "id": 11, "name": "OSX",      "client_id": 1,    "another_col": "val" },
  { "id": 12, "name": "ProjectX", "client_id": null, "another_col": null },
  { "id": 13, "name": "Y",        "client_id": null, "another_col": null }
]
```

On large result sets, the unused keys with `null` values can waste bandwidth unnecessarily. To remove them, specify `nulls=stripped` as a parameter of `application/vnd.pgrst.array`:

```bash
curl "http://localhost:3000/projects?id=gt.10" \
  -H "Accept: application/vnd.pgrst.array+json;nulls=stripped"
```

This returns

```json
[
  { "id": 11, "name": "OSX", "client_id": 1, "another_col": "val" },
  { "id": 12, "name": "ProjectX" },
  { "id": 13, "name": "Y"}
]
```
<a id="req_body"></a>

## Request Body

The server handles the following request body media types:

- `application/json`

- `application/x-www-form-urlencoded`

- `text/csv`

For [Tables and Views](tables_views.md#tables_views) this works on `POST`, `PATCH` and `PUT` methods. For [Functions as RPC](functions.md#functions), it works on `POST` methods.

For functions there are three additional types:

- `application/octet-stream`

- `text/plain`

- `text/xml`

See [Functions with a single unnamed parameter](functions.md#function_single_unnamed).
