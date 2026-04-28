<a id="computed_cols"></a>

# Computed Fields

Computed fields are virtual columns that are not stored in a table. PostgreSQL makes it possible to implement them using functions on table types.

```postgres
CREATE TABLE people (
  first_name text
, last_name  text
, job        text
);

-- a computed field that combines data from two columns
CREATE FUNCTION full_name(people)
RETURNS text AS $$
  SELECT $1.first_name || ' ' || $1.last_name;
$$ LANGUAGE SQL;
```

## Horizontal Filtering on Computed Fields

[Horizontal Filtering](tables_views.md#h_filter) can be applied to computed fields. For example, we can do a [Full-Text Search](tables_views.md#fts) on `full_name`:

```postgres
-- (optional) you can add an index on the computed field to speed up the query
CREATE INDEX people_full_name_idx ON people
  USING GIN (to_tsvector('english', full_name(people)));
```

```bash
curl "http://localhost:3000/people?full_name=fts.Beckett"
```

```json
[
  {"first_name": "Samuel", "last_name": "Beckett", "job": "novelist"}
]
```

## Vertical Filtering on Computed Fields

Computed fields won't appear on the response by default but you can use [Vertical Filtering](tables_views.md#v_filter) to include them:

```bash
curl "http://localhost:3000/people?select=full_name,job"
```

```json
[
  {"full_name": "Samuel Beckett", "job": "novelist"}
]
```

## Ordering on Computed Fields

[Ordering](tables_views.md#ordering) on computed fields is also possible:

```bash
curl "http://localhost:3000/people?order=full_name.desc"
```

!!! important

    Computed fields must be created in the [exposed schema](../../../references/configuration.md#db-schemas) or in a schema in the [extra search path](../../../references/configuration.md#db-extra-search-path) to be used in this way. When placing the computed field in the [exposed schema](../../../references/configuration.md#db-schemas) you can use an **unnamed** parameter, as in the example above, to prevent it from being exposed as an [RPC](functions.md#functions) under `/rpc`.

!!! note

    - PostgreSQL 12 introduced [generated columns](https://www.postgresql.org/docs/12/ddl-generated-columns.html), which can also compute a value based on other columns. However they're stored, not virtual.

    - "computed fields" are documented on https://www.postgresql.org/docs/current/rowtypes.html#ROWTYPES-USAGE (search for "computed fields")

    - On previous PostgREST versions this feature was documented with the name of "computed columns".
