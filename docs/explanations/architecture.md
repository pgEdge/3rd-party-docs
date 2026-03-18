# Architecture

This page describes the architecture of PostgREST.

## Bird's Eye View

You can click on the components to navigate to their respective documentation.

<object width="100%" data="../_static/arch.svg" type="image/svg+xml"></object>

## Code Map

This section talks briefly about various important modules.

### Main

The starting point of the program is [Main.hs](https://github.com/PostgREST/postgrest/blob/main/main/Main.hs).

### CLI

Main then calls [CLI.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/CLI.hs), which is in charge of [CLI](../references/cli.md#cli).

### App

[App.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/App.hs) is then in charge of composing the different modules.

### Auth

[Auth.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/Auth.hs) is in charge  of [Authentication](../references/auth.md#authn).

### Api Request

[ApiRequest.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/ApiRequest.hs) is in charge of parsing the URL query string (following PostgREST syntax), the request headers, and the request body.

A request might be rejected at this level if it's invalid. For example when providing an unknown media type to PostgREST or using an unknown HTTP method.

### Plan

Using the Schema Cache, [Plan.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/Plan.hs) generates an internal AST, filling out-of-band SQL details (like an `ON CONFLICT (pk)` clause) required to complete the user request.

A request might be rejected at this level if it's invalid. For example when doing resource embedding on a nonexistent resource.

### Query

[Query.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/Query.hs) generates the SQL queries (parametrized and prepared) required to satisfy the user request.

Only at this stage a connection from the pool might be used.

### Schema Cache

[SchemaCache.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/SchemaCache.hs) is in charge of [Schema Cache](../references/schema_cache.md#schema_cache).

### Config

[Config.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/Config.hs) is in charge of [Configuration](../references/configuration.md#configuration).

### Admin

[Admin.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/Admin.hs) is in charge of the [Admin Server](../references/admin_server.md#admin_server).

### HTTP

The HTTP server is provided by [Warp](https://aosabook.org/en/posa/warp.html).

### Listener

[Listener.hs](https://github.com/PostgREST/postgrest/blob/main/src/PostgREST/Listener.hs) is in charge of the [Listener](../references/listener.md#listener).
