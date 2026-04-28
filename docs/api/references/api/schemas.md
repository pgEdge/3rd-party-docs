<a id="schemas"></a>

# Schemas

PostgREST can expose a single or multiple schema's tables, views and functions. The [active database role](../../../references/auth.md#roles) must have the usage privilege on the schemas to access them.

## Single schema

To expose a single schema, specify a single value in [db-schemas](../../../references/configuration.md#db-schemas).

```bash
db-schemas = "api"
```

This schema is added to the [search_path](https://www.postgresql.org/docs/current/ddl-schemas.html#DDL-SCHEMAS-PATH) of every request using [Transaction-Scoped Settings](../../../references/transactions.md#tx_settings).
<a id="multiple-schemas"></a>

## Multiple schemas

To expose multiple schemas, specify a comma-separated list on [db-schemas](../../../references/configuration.md#db-schemas):

```bash
db-schemas = "tenant1, tenant2"
```

To switch schemas, use the `Accept-Profile` and `Content-Profile` headers.

If you don't specify a Profile header, the first schema in the list(`tenant1` here) is selected as the default schema.

Only the selected schema gets added to the [search_path](https://www.postgresql.org/docs/current/ddl-schemas.html#DDL-SCHEMAS-PATH) of every request.

!!! note

    These headers are based on the "Content Negotiation by Profile" spec: https://www.w3.org/TR/dx-prof-conneg

### GET/HEAD

For GET or HEAD, select the schema with `Accept-Profile`.

```bash
curl "http://localhost:3000/items" \
  -H "Accept-Profile: tenant2"
```

### Other methods

For POST, PATCH, PUT and DELETE, select the schema with `Content-Profile`.

```bash
curl "http://localhost:3000/items" \
  -X POST -H "Content-Type: application/json" \
  -H "Content-Profile: tenant2" \
  -d '{...}'
```

You can also select the schema for [Functions as RPC](functions.md#functions) and [OpenAPI](openapi.md#open-api).

### Restricted schemas

You can only switch to a schema included in [db-schemas](../../../references/configuration.md#db-schemas). Using another schema will result in an error:

```bash
curl "http://localhost:3000/items" \
  -H "Accept-Profile: tenant3"
```

```
{
  "code":"PGRST106",
  "details":null,
  "hint":null,
  "message":"The schema must be one of the following: tenant1, tenant2"
}
```

### Dynamic schemas

To add schemas dynamically, you can use [In-Database Configuration](../../../references/configuration.md#in_db_config) plus [config reloading](../../../references/configuration.md#config_reloading_notify) and [schema cache reloading](../../../references/schema_cache.md#schema_reloading_notify). Here are some options for how to do this:

- If the schemas' names have a pattern, like a `tenant_` prefix, do:

```postgres
create or replace function postgrest.pre_config()
returns void as $$
  select
    set_config('pgrst.db_schemas', string_agg(nspname, ','), true)
  from pg_namespace
  where nspname like 'tenant_%';
$$ language sql;
```

- If there's no name pattern but they're created with a particular role (`CREATE SCHEMA mine AUTHORIZATION joe`), do:

```postgres
create or replace function postgrest.pre_config()
returns void as $$
  select
    set_config('pgrst.db_schemas', string_agg(nspname, ','), true)
  from pg_namespace
  where nspowner = 'joe'::regrole;
$$ language sql;
```

- Otherwise, you might need to create a table that stores the allowed schemas.

```postgres
create table postgrest.config (schemas text);

create or replace function postgrest.pre_config()
returns void as $$
  select
    set_config('pgrst.db_schemas', schemas, true)
  from postgrest.config;
$$ language sql;
```

Then each time you add an schema, do:

```postgres
NOTIFY pgrst, 'reload config';
NOTIFY pgrst, 'reload schema';
```
