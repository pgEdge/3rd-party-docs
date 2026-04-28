<a id="error_source"></a>

# Errors

PostgREST error messages follow the PostgreSQL error structure. It includes `MESSAGE`, `DETAIL`, `HINT`, `ERRCODE` and will add an HTTP status code to the response.
<a id="postgresql_errors"></a>

## Errors from PostgreSQL

PostgREST will forward errors coming from PostgreSQL. For instance, on a failed constraint:

```http
POST /projects HTTP/1.1
```

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
```

```json

{
    "code": "23502",
    "details": "Failing row contains (null, foo, null).",
    "hint": null,
    "message": "null value in column \"id\" of relation \"projects\" violates not-null constraint"
}
```
<a id="status_codes"></a>

### HTTP Status Codes

PostgREST translates [PostgreSQL error codes](https://www.postgresql.org/docs/current/errcodes-appendix.html) into HTTP status as follows:

| PostgreSQL error code(s) | HTTP status | Error description |
|---|---|---|
| 08* | 503 | pg connection err |
| 09* | 500 | triggered action exception |
| 0L* | 403 | invalid grantor |
| 0P* | 403 | invalid role specification |
| 23503 | 409 | foreign key violation |
| 23505 | 409 | uniqueness violation |
| 25006 | 405 | read only sql transaction |
| 25* | 500 | invalid transaction state |
| 28* | 403 | invalid auth specification |
| 2D* | 500 | invalid transaction termination |
| 38* | 500 | external routine exception |
| 39* | 500 | external routine invocation |
| 3B* | 500 | savepoint exception |
| 40* | 500 | transaction rollback |
| 53400 | 500 | config limit exceeded |
| 53* | 503 | insufficient resources |
| 54* | 500 | too complex |
| 55* | 500 | obj not in prerequisite state |
| 57* | 500 | operator intervention |
| 58* | 500 | system error |
| F0* | 500 | config file error |
| HV* | 500 | foreign data wrapper error |
| P0001 | 400 | default code for "raise" |
| P0* | 500 | PL/pgSQL error |
| XX* | 500 | internal error |
| 42883 | 404 | undefined function |
| 42P01 | 404 | undefined table |
| 42P17 | 500 | infinite recursion |
| 42501 | | if authenticated 403, | else 401 | insufficient privileges |
| other | 400 |  |

## Errors from PostgREST

Errors that come from PostgREST itself maintain the same structure but differ in the `PGRST` prefix in the `code` field. For instance, when querying a function that does not exist in the schema cache:

```http
POST /rpc/nonexistent_function HTTP/1.1
```

```http
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
```

```json
{
  "hint": "...",
  "details": null
  "code": "PGRST202",
  "message": "Could not find the api.nonexistent_function() function in the schema cache"
}
```
<a id="pgrst_errors"></a>

### PostgREST Error Codes

PostgREST error codes have the form `PGRSTgxx`.

- `PGRST` is the prefix that differentiates the error from a PostgreSQL error.

- `g` is the error group

- `xx` is the error identifier in the group.
<a id="pgrst0**"></a>

#### Group 0 - Connection

Related to the connection with the database.

<table>
<thead>
<tr>
  <th>Code</th>
  <th>HTTP status</th>
  <th>Description</th>
</tr>
</thead>
<tbody>
<tr>
  <td>.. _pgrst000:PGRST000</td>
  <td>503</td>
  <td>Could not connect with the database due to an incorrect <a href="configuration.md#db-uri">db-uri</a> or due to the PostgreSQL service not running.</td>
</tr>
<tr>
  <td>.. _pgrst001:PGRST001</td>
  <td>503</td>
  <td>Could not connect with the database due to an internal error.</td>
</tr>
<tr>
  <td>.. _pgrst002:PGRST002</td>
  <td>503</td>
  <td>Could not connect with the database when building the Schema Cache due to the PostgreSQL service not running.</td>
</tr>
<tr>
  <td>.. _pgrst003:PGRST003</td>
  <td>504</td>
  <td>The request timed out waiting for a pool connection to be available. See <a href="configuration.md#db-pool-acquisition-timeout">db-pool-acquisition-timeout</a>.</td>
</tr>
</tbody>
</table>
<a id="pgrst1**"></a>

#### Group 1 - Api Request

Related to the HTTP request elements.

<table>
<thead>
<tr>
  <th>Code</th>
  <th>HTTP status</th>
  <th>Description</th>
</tr>
</thead>
<tbody>
<tr>
  <td>.. _pgrst100:PGRST100</td>
  <td>400</td>
  <td>Parsing error in the query string parameter. See <a href="../api/references/api/tables_views.md#h_filter">Horizontal Filtering</a>, <a href="../api/references/api/tables_views.md#operators">Operators</a> and <a href="../api/references/api/tables_views.md#ordering">Ordering</a>.</td>
</tr>
<tr>
  <td>.. _pgrst101:PGRST101</td>
  <td>405</td>
  <td>For <a href="../api/references/api/functions.md#functions">functions</a>, only <code>GET</code> and <code>POST</code> verbs are allowed. Any other verb will throw this error.</td>
</tr>
<tr>
  <td>.. _pgrst102:PGRST102</td>
  <td>400</td>
  <td>An invalid request body was sent(e.g. an empty body or malformed JSON).</td>
</tr>
<tr>
  <td>.. _pgrst103:PGRST103</td>
  <td>416</td>
  <td>An invalid range was specified for <a href="../api/references/api/pagination_count.md#limits">Limits and Pagination</a>.</td>
</tr>
<tr>
  <td>.. _pgrst105:PGRST105</td>
  <td>405</td>
  <td>An invalid <a href="../api/references/api/tables_views.md#upsert_put">PUT</a> request was done</td>
</tr>
<tr>
  <td>.. _pgrst106:PGRST106</td>
  <td>406</td>
  <td>The schema specified when <a href="../api/references/api/schemas.md#multiple-schemas">switching schemas</a> is not present in the <a href="configuration.md#db-schemas">db-schemas</a> configuration variable.</td>
</tr>
<tr>
  <td>.. _pgrst107:PGRST107</td>
  <td>415</td>
  <td>The <code>Content-Type</code> sent in the request is invalid.</td>
</tr>
<tr>
  <td>.. _pgrst108:PGRST108</td>
  <td>400</td>
  <td>The filter is applied to a embedded resource that is not specified in the <code>select</code> part of the query string. See <a href="../api/references/api/resource_embedding.md#embed_filters">Embedded Filters</a>.</td>
</tr>
<tr>
  <td>.. _pgrst111:PGRST111</td>
  <td>500</td>
  <td>An invalid <code>response.headers</code> was set. See <a href="transactions.md#guc_resp_hdrs">Response Headers</a>.</td>
</tr>
<tr>
  <td>.. _pgrst112:PGRST112</td>
  <td>500</td>
  <td>The status code must be a positive integer. See <a href="transactions.md#guc_resp_status">Response Status Code</a>.</td>
</tr>
<tr>
  <td>.. _pgrst114:PGRST114</td>
  <td>400</td>
  <td>For an <a href="../api/references/api/tables_views.md#upsert_put">UPSERT using PUT</a>, when <a href="../api/references/api/pagination_count.md#limits">limits and offsets</a> are used.</td>
</tr>
<tr>
  <td>.. _pgrst115:PGRST115</td>
  <td>400</td>
  <td>For an <a href="../api/references/api/tables_views.md#upsert_put">UPSERT using PUT</a>, when the primary key in the query string and the body are different.</td>
</tr>
<tr>
  <td>.. _pgrst116:PGRST116</td>
  <td>406</td>
  <td>More than 1 or no items where returned when requesting a singular response. See <a href="../api/references/api/resource_representation.md#singular_plural">Singular or Plural</a>.</td>
</tr>
<tr>
  <td>.. _pgrst117:PGRST117</td>
  <td>405</td>
  <td>The HTTP verb used in the request in not supported.</td>
</tr>
<tr>
  <td>.. _pgrst118:PGRST118</td>
  <td>400</td>
  <td>Could not order the result using the related table because there is no many-to-one or one-to-one relationship between them.</td>
</tr>
<tr>
  <td>.. _pgrst120:PGRST120</td>
  <td>400</td>
  <td>An embedded resource can only be filtered using the <code>is.null</code> or <code>not.is.null</code> <a href="../api/references/api/tables_views.md#operators">operators</a>.</td>
</tr>
<tr>
  <td>.. _pgrst121:PGRST121</td>
  <td>500</td>
  <td>PostgREST can't parse the JSON objects in RAISE <code>PGRST</code> error. See <a href="#raise_headers">raise headers</a>.</td>
</tr>
<tr>
  <td>.. _pgrst122:PGRST122</td>
  <td>400</td>
  <td>Invalid preferences found in <code>Prefer</code> header with <code>Prefer: handling=strict</code>. See <a href="../api/references/api/preferences.md#prefer_handling">Strict or Lenient Handling</a>.</td>
</tr>
<tr>
  <td>.. _pgrst123:PGRST123</td>
  <td>400</td>
  <td>Aggregate functions are disabled. See <a href="configuration.md#db-aggregates-enabled">db-aggregates-enabled</a>.</td>
</tr>
<tr>
  <td>.. _pgrst124:PGRST124</td>
  <td>400</td>
  <td><code>max-affected</code> preference is violated. See <a href="../api/references/api/preferences.md#prefer_max_affected">Max Affected</a>.</td>
</tr>
<tr>
  <td>.. _pgrst125:PGRST125</td>
  <td>404</td>
  <td>Invalid path is specified in request URL.</td>
</tr>
<tr>
  <td>.. _pgrst126:PGRST126</td>
  <td>404</td>
  <td>Open API config is disabled but API root path is accessed. See <a href="configuration.md#openapi-mode">openapi-mode</a>.</td>
</tr>
<tr>
  <td>.. _pgrst127:PGRST127</td>
  <td>400</td>
  <td>The feature specified in the <code>details</code> field is not implemented.</td>
</tr>
<tr>
  <td>.. _pgrst128:PGRST128</td>
  <td>400</td>
  <td><code>max-affected</code> preference is violated with <code>RPC</code> call. See <a href="../api/references/api/preferences.md#prefer_max_affected">Max Affected</a>.</td>
</tr>
</tbody>
</table>
<a id="pgrst2**"></a>

#### Group 2 - Schema Cache

Related to a [Schema Cache](schema_cache.md#schema_cache). Most of the time, these errors are solved by [Schema Cache Reloading](schema_cache.md#schema_reloading).

<table>
<thead>
<tr>
  <th>Code</th>
  <th>HTTP status</th>
  <th>Description</th>
</tr>
</thead>
<tbody>
<tr>
  <td>.. _pgrst200:PGRST200</td>
  <td>400</td>
  <td>Caused by stale foreign key relationships, otherwise any of the embedding resources or the relationship itself may not exist in the database.</td>
</tr>
<tr>
  <td>.. _pgrst201:PGRST201</td>
  <td>300</td>
  <td>An ambiguous embedding request was made. See <a href="../api/references/api/resource_embedding.md#complex_rels">Foreign Key Joins on Multiple Foreign Key Relationships</a>.</td>
</tr>
<tr>
  <td>.. _pgrst202:PGRST202</td>
  <td>404</td>
  <td>Caused by a stale function signature, otherwise the function may not exist in the database.</td>
</tr>
<tr>
  <td>.. _pgrst203:PGRST203</td>
  <td>300</td>
  <td>Caused by requesting overloaded functions with the same argument names but different types, or by using a <code>POST</code> verb to request overloaded functions with a <code>JSON</code> or <code>JSONB</code> type unnamed parameter. The solution is to rename the function or add/modify the names of the arguments.</td>
</tr>
<tr>
  <td>.. _pgrst204:PGRST204</td>
  <td>400</td>
  <td>Caused when the <a href="../api/references/api/tables_views.md#specify_columns">column specified</a> in the <code>columns</code> query parameter is not found.</td>
</tr>
<tr>
  <td>.. _pgrst205:PGRST205</td>
  <td>404</td>
  <td>Caused when the <a href="../api/references/api/tables_views.md#tables_views">table specified</a> in the URI is not found.</td>
</tr>
</tbody>
</table>
<a id="pgrst3**"></a>

#### Group 3 - JWT

Related to the authentication process using JWT. You can follow the [Tutorial 1 - The Golden Key](../tutorials/tut1.md#tut1) for an example on how to implement authentication and the Authentication page for more information on this process.

<table>
<thead>
<tr>
  <th>Code</th>
  <th>HTTP status</th>
  <th>Description</th>
</tr>
</thead>
<tbody>
<tr>
  <td>.. _pgrst300:PGRST300</td>
  <td>500</td>
  <td>A <a href="configuration.md#jwt-secret">JWT secret</a> is missing from the configuration.</td>
</tr>
<tr>
  <td>.. _pgrst301:PGRST301</td>
  <td>401</td>
  <td>Provided JWT couldn't be decoded or it is invalid.</td>
</tr>
<tr>
  <td>.. _pgrst302:PGRST302</td>
  <td>401</td>
  <td>Attempted to do a request without <a href="auth.md#bearer_auth">Bearer Authentication</a> when the anonymous role is disabled by not setting it in <a href="configuration.md#db-anon-role">db-anon-role</a>.</td>
</tr>
<tr>
  <td>.. _pgrst303:PGRST303</td>
  <td>401</td>
  <td><a href="auth.md#jwt_claims_validation">JWT claims validation</a> or parsing failed.</td>
</tr>
</tbody>
</table>
<a id="pgrst_X**"></a>

#### Group X - Internal

Internal errors. If you encounter any of these, you may have stumbled on a PostgREST bug, please [open an issue](https://github.com/PostgREST/postgrest/issues) and we'll be glad to fix it.

<table>
<thead>
<tr>
  <th>Code</th>
  <th>HTTP status</th>
  <th>Description</th>
</tr>
</thead>
<tbody>
<tr>
  <td>.. _pgrstX00:PGRSTX00</td>
  <td>500</td>
  <td>Internal errors related to the library used for connecting to the database.</td>
</tr>
</tbody>
</table>
<a id="custom_errors"></a>

## Custom Errors

You can customize the errors by using the [RAISE statement](https://www.postgresql.org/docs/current/plpgsql-errors-and-messages.html#PLPGSQL-STATEMENTS-RAISE)  on functions.
<a id="raise_error"></a>

### RAISE errors with HTTP Status Codes

Custom status codes can be done by raising SQL exceptions inside [functions](../api/references/api/functions.md#functions). For instance, here's a saucy function that always responds with an error:

```postgres
CREATE OR REPLACE FUNCTION just_fail() RETURNS void
  LANGUAGE plpgsql
  AS $$
BEGIN
  RAISE EXCEPTION 'I refuse!'
    USING DETAIL = 'Pretty simple',
          HINT = 'There is nothing you can do.';
END
$$;
```

Calling the function returns HTTP 400 with the body

```json
{
  "message":"I refuse!",
  "details":"Pretty simple",
  "hint":"There is nothing you can do.",
  "code":"P0001"
}
```

One way to customize the HTTP status code is by raising particular exceptions according to the PostgREST [error to status code mapping](#status_codes). For example, `RAISE insufficient_privilege` will respond with HTTP 401/403 as appropriate.

For even greater control of the HTTP status code, raise an exception of the `PTxyz` type. For instance to respond with HTTP 402, raise `PT402`:

```postgres
RAISE sqlstate 'PT402' using
  message = 'Payment Required',
  detail = 'Quota exceeded',
  hint = 'Upgrade your plan';
```

Returns:

```http
HTTP/1.1 402 Payment Required
Content-Type: application/json; charset=utf-8

{
  "message": "Payment Required",
  "details": "Quota exceeded",
  "hint": "Upgrade your plan",
  "code": "PT402"
}
```
<a id="raise_headers"></a>

### Add HTTP Headers with RAISE

For full control over headers and status you can raise a `PGRST` SQLSTATE error. You can achieve this by adding the `code`, `message`, `detail` and `hint` in the PostgreSQL error message field as a JSON object. Here, the `details` and `hint` are optional. Similarly, the `status` and `headers` must be added to the SQL error detail field as a JSON object. For instance:

```postgres
RAISE sqlstate 'PGRST' USING
    message = '{"code":"123","message":"Payment Required","details":"Quota exceeded","hint":"Upgrade your plan"}',
    detail = '{"status":402,"headers":{"X-Powered-By":"Nerd Rage"}}';
```

Returns:

```http
HTTP/1.1 402 Payment Required
Content-Type: application/json; charset=utf-8
X-Powered-By: Nerd Rage

{
  "message": "Payment Required",
  "details": "Quota exceeded",
  "hint": "Upgrade your plan",
  "code": "123"
}
```

For non standard HTTP status, you can optionally add `status_text` to describe the status code. For status code `419` the detail field may look like this:

```postgres
detail = '{"status":419,"status_text":"Page Expired","headers":{"X-Powered-By":"Nerd Rage"}}';
```

If PostgREST can't parse the JSON objects `message` and `detail`, it will throw a `PGRST121` error. See [Errors from PostgREST](#pgrst1**).
<a id="proxy-status_header"></a>

## Proxy-Status Header

For error cases, the standard [Proxy-Status](https://www.rfc-editor.org/rfc/rfc9209.html#name-the-proxy-status-http-field) header is returned with the error code. The error code comes from either [PostgREST](#pgrst_errors), [PostgreSQL](#postgresql_errors) or [Custom](#custom_errors) errors. This is useful when doing `HEAD` requests where the HTTP status is not descriptive enough.

For example, doing a request on a table with high count (say 30_000_000), we get:

```http
HEAD /table HTTP/1.1
Prefer: count=exact
```

```http
HTTP/1.1 500 Internal Server Error
Proxy-Status: PostgREST; error=57014
```

The PostgreSQL error code `57014` ([ref](https://www.postgresql.org/docs/current/errcodes-appendix.html)) reveals that the error is due to a short `statement_timeout` value.
