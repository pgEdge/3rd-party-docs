# psycopg2.sql -- SQL string composition

*module* `psycopg2.sql`

*New in version 2.7.*

The module contains objects and functions useful to generate SQL dynamically, in a convenient and safe way. SQL identifiers (e.g. names of tables and fields) cannot be passed to the `execute()` method like query arguments:

```
# This will not work
table_name = 'my_table'
cur.execute("insert into %s values (%s, %s)", [table_name, 10, 20])
```

The SQL query should be composed before the arguments are merged, for instance:

```
# This works, but it is not optimal
table_name = 'my_table'
cur.execute(
    "insert into %s values (%%s, %%s)" % table_name,
    [10, 20])
```

This sort of works, but it is an accident waiting to happen: the table name may be an invalid SQL literal and need quoting; even more serious is the security problem in case the table name comes from an untrusted source. The name should be escaped using `quote_ident()`:

```
# This works, but it is not optimal
table_name = 'my_table'
cur.execute(
    "insert into %s values (%%s, %%s)" % ext.quote_ident(table_name, cur),
    [10, 20])
```

This is now safe, but it somewhat ad-hoc. In case, for some reason, it is necessary to include a value in the query string (as opposite as in a value) the merging rule is still different (`adapt()` should be used...). It is also still relatively dangerous: if `quote_ident()` is forgotten somewhere, the program will usually work, but will eventually crash in the presence of a table or field name with containing characters to escape, or will present a potentially exploitable weakness.

The objects exposed by the `psycopg2.sql` module allow generating SQL statements on the fly, separating clearly the variable parts of the statement from the query parameters:

```
from psycopg2 import sql

cur.execute(
    sql.SQL("insert into {} values (%s, %s)")
        .format(sql.Identifier('my_table')),
    [10, 20])
```

## Module usage

Usually you should express the template of your query as an `SQL` instance with `{}`-style placeholders and use `format()` to merge the variable parts into them, all of which must be `Composable` subclasses. You can still have `%s` -style placeholders in your query and pass values to `execute()`: such value placeholders will be untouched by `format()`:

```
query = sql.SQL("select {field} from {table} where {pkey} = %s").format(
    field=sql.Identifier('my_name'),
    table=sql.Identifier('some_table'),
    pkey=sql.Identifier('id'))
```

The resulting object is meant to be passed directly to cursor methods such as `execute()`, `executemany()`, `copy_expert()`, but can also be used to compose a query as a Python string, using the `as_string()` method:

```
cur.execute(query, (42,))
```

If part of your query is a variable sequence of arguments, such as a comma-separated list of field names, you can use the `SQL.join()` method to pass them to the query:

```
query = sql.SQL("select {fields} from {table}").format(
    fields=sql.SQL(',').join([
        sql.Identifier('field1'),
        sql.Identifier('field2'),
        sql.Identifier('field3'),
    ]),
    table=sql.Identifier('some_table'))
```

## !sql objects

The `sql` objects are in the following inheritance hierarchy:

  `Composable`: the base class exposing the common interface<br>
  `|__` `SQL`: a literal snippet of an SQL query<br>
  `|__` `Identifier`: a PostgreSQL identifier or dot-separated sequence of identifiers<br>
  `|__` `Literal`: a value hardcoded into a query<br>
  `|__` `Placeholder`: a `%s` -style placeholder whose value will be added later e.g. by `execute()`<br>
  `|__` `Composed`: a sequence of `Composable` instances.<br>

*class* `Composable`

*method* `as_string`

*class* `SQL`

*attribute* `string`

*method* `format`

*method* `join`

*class* `Identifier`

*Changed in version 2.8.* added support for multiple strings.

*attribute* `strings`

*New in version 2.8.* previous verions only had a `string` attribute. The attribute
still exists but is deprecate and will only work if the
`Identifier` wraps a single string.

*class* `Literal`

*attribute* `wrapped`

*class* `Placeholder`

*attribute* `name`

*class* `Composed`

*attribute* `seq`

*method* `join`
