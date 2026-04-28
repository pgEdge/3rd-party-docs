# Frequently Asked Questions

Here are a few gotchas you may encounter using `psycopg2`.  Feel free to suggest new entries!

## Meta
<a id="faq-question"></a>

How do I ask a question?

- Have you first checked if your question is answered already in the
    documentation?

- If your question is about installing psycopg, have you checked the
    [install FAQ](#faq-compile) and the [install docs](install.md#installation)?

- Have you googled for your error message?

- If you haven't found an answer yet, please write to the Mailing List

- If you haven't found a bug, DO NOT write to the bug tracker to ask
    questions. You will only get piro grumpy.
<a id="mailing list"></a>
<a id="faq-transactions"></a>

## Problems with transactions handling
<a id="faq-idle-in-transaction"></a>

Why does `psycopg2` leave database sessions "idle in transaction"? Psycopg normally starts a new transaction the first time a query is executed, e.g. calling `cursor.execute()`, even if the command is a `SELECT`.  The transaction is not closed until an explicit `commit()` or `rollback()`.

If you are writing a long-living program, you should probably make sure to call one of the transaction closing methods before leaving the connection unused for a long time (which may also be a few seconds, depending on the concurrency level in your database).  Alternatively you can use a connection in `autocommit` mode to avoid a new transaction to be started at the first command.
<a id="faq-transaction-aborted"></a>

I receive the error *current transaction is aborted, commands ignored until end of transaction block* and can't do anything else! There was a problem *in the previous* command to the database, which resulted in an error.  The database will not recover automatically from this condition: you must run a `rollback()` before sending new commands to the session (if this seems too harsh, remember that PostgreSQL supports nested transactions using the [`SAVEPOINT`](https://www.postgresql.org/docs/current/static/sql-savepoint.html) command).
<a id="SAVEPOINT"></a>
<a id="faq-transaction-aborted-multiprocess"></a>

Why do I get the error *current transaction is aborted, commands ignored until end of transaction block* when I use `multiprocessing` (or any other forking system) and not when use `threading`? Psycopg's connections can't be shared across processes (but are thread safe).  If you are forking the Python process make sure to create a new connection in each forked child. See [Thread and process safety](usage.md#thread-safety) for further informations.
<a id="faq-types"></a>

## Problems with type conversions
<a id="faq-cant-adapt"></a>

Why does `cursor.execute()` raise the exception *can't adapt*? Psycopg converts Python objects in a SQL string representation by looking at the object class.  The exception is raised when you are trying to pass as query parameter an object for which there is no adapter registered for its class.  See [Adapting new Python types to SQL syntax](advanced.md#adapting-new-types) for informations.
<a id="faq-number-required"></a>

I can't pass an integer or a float parameter to my query: it says *a number is required*, but *it is* a number! In your query string, you always have to use `%s`  placeholders, even when passing a number.  All Python objects are converted by Psycopg in their SQL representation, so they get passed to the query as strings. See [Passing parameters to SQL queries](usage.md#query-parameters).:

```
>>> cur.execute("INSERT INTO numbers VALUES (%d)", (42,)) # WRONG
>>> cur.execute("INSERT INTO numbers VALUES (%s)", (42,)) # correct
```
<a id="faq-not-all-arguments-converted"></a>

I try to execute a query but it fails with the error *not all arguments converted during string formatting* (or *object does not support indexing*). Why? Psycopg always require positional arguments to be passed as a sequence, even when the query takes a single parameter.  And remember that to make a single item tuple in Python you need a comma!  See [Passing parameters to SQL queries](usage.md#query-parameters).:

```
>>> cur.execute("INSERT INTO foo VALUES (%s)", "bar")    # WRONG
>>> cur.execute("INSERT INTO foo VALUES (%s)", ("bar"))  # WRONG
>>> cur.execute("INSERT INTO foo VALUES (%s)", ("bar",)) # correct
>>> cur.execute("INSERT INTO foo VALUES (%s)", ["bar"])  # correct
```
<a id="faq-unicode"></a>

My database is Unicode, but I receive all the strings as UTF-8 `str`. Can I receive `unicode` objects instead? The following magic formula will do the trick:

```
psycopg2.extensions.register_type(psycopg2.extensions.UNICODE)
psycopg2.extensions.register_type(psycopg2.extensions.UNICODEARRAY)
```

See [unicode-handling](usage.md#unicode-handling) for the gory details.
<a id="faq-bytes"></a>

My database is in mixed encoding. My program was working on Python 2 but Python 3 fails decoding the strings. How do I avoid decoding? From psycopg 2.8 you can use the following adapters to always return bytes from strings:

```
psycopg2.extensions.register_type(psycopg2.extensions.BYTES)
psycopg2.extensions.register_type(psycopg2.extensions.BYTESARRAY)
```

See [unicode-handling](usage.md#unicode-handling) for an example.
<a id="faq-float"></a>

Psycopg converts `decimal`/ `numeric` database types into Python `Decimal` objects. Can I have `float` instead? You can register a customized adapter for PostgreSQL decimal type:

```
DEC2FLOAT = psycopg2.extensions.new_type(
    psycopg2.extensions.DECIMAL.values,
    'DEC2FLOAT',
    lambda value, curs: float(value) if value is not None else None)
psycopg2.extensions.register_type(DEC2FLOAT)
```

See [Type casting of SQL types into Python objects](advanced.md#type-casting-from-sql-to-python) to read the relevant documentation. If you find `psycopg2.extensions.DECIMAL` not available, use `psycopg2._psycopg.DECIMAL` instead.
<a id="faq-json-adapt"></a>

Psycopg automatically converts PostgreSQL `json` data into Python objects. How can I receive strings instead? The easiest way to avoid JSON parsing is to register a no-op function with `register_default_json()`:

```
psycopg2.extras.register_default_json(loads=lambda x: x)
```

See [JSON_ adaptation](extras.md#adapt-json) for further details.
<a id="faq-jsonb-adapt"></a>

Psycopg converts `json` values into Python objects but `jsonb` values are returned as strings. Can `jsonb` be converted automatically? Automatic conversion of `jsonb` values is supported from Psycopg release 2.5.4. For previous versions you can register the `json` typecaster on the `jsonb` oids (which are known and not supposed to change in future PostgreSQL versions):

```
psycopg2.extras.register_json(oid=3802, array_oid=3807, globally=True)
```

See [JSON_ adaptation](extras.md#adapt-json) for further details.
<a id="faq-identifier"></a>

How can I pass field/table names to a query? The arguments in the `execute()` methods can only represent data to pass to the query: they cannot represent a table or field name:

```
# This doesn't work
cur.execute("insert into %s values (%s)", ["my_table", 42])
```

If you want to build a query dynamically you can use the objects exposed by the `psycopg2.sql` module:

```
cur.execute(
    sql.SQL("insert into %s values (%%s)") % [sql.Identifier("my_table")],
    [42])
```
<a id="faq-bytea-9.0"></a>

Transferring binary data from PostgreSQL 9.0 doesn't work. PostgreSQL 9.0 uses by default `the "hex" format`__ to transfer `bytea` data: the format can't be parsed by the libpq 8.4 and earlier. The problem is solved in Psycopg 2.4.1, that uses its own parser for the `bytea` format. For previous Psycopg releases, three options to solve the problem are:

- set the bytea_output__ parameter to `escape` in the server;

- execute the database command `SET bytea_output TO escape;` in the
    session before reading binary data;

- upgrade the libpq library on the client to at least 9.0.
<a id="_"></a>
<a id="_"></a>
<a id="faq-array"></a>

Arrays of *TYPE* are not casted to list. Arrays are only casted to list when their oid is known, and an array typecaster is registered for them. If there is no typecaster, the array is returned unparsed from PostgreSQL (e.g. `{a,b,c}`). It is easy to create a generic arrays typecaster, returning a list of array: an example is provided in the `new_array_type()` documentation.
<a id="faq-best-practices"></a>

## Best practices
<a id="faq-reuse-cursors"></a>

When should I save and re-use a cursor as opposed to creating a new one as needed? Cursors are lightweight objects and creating lots of them should not pose any kind of problem. But note that cursors used to fetch result sets will cache the data and use memory in proportion to the result set size. Our suggestion is to almost always create a new cursor and dispose old ones as soon as the data is not required anymore (call `close()` on them.) The only exception are tight loops where one usually use the same cursor for a whole bunch of `INSERT`s or `UPDATE`s.
<a id="faq-reuse-connections"></a>

When should I save and re-use a connection as opposed to creating a new one as needed? Creating a connection can be slow (think of SSL over TCP) so the best practice is to create a single connection and keep it open as long as required. It is also good practice to rollback or commit frequently (even after a single `SELECT` statement) to make sure the backend is never left "idle in transaction".  See also `psycopg2.pool` for lightweight connection pooling.
<a id="faq-named-cursors"></a>

What are the advantages or disadvantages of using named cursors? The only disadvantages is that they use up resources on the server and that there is a little overhead because at least two queries (one to create the cursor and one to fetch the initial result set) are issued to the backend. The advantage is that data is fetched one chunk at a time: using small `fetchmany()` values it is possible to use very little memory on the client and to skip or discard parts of the result set.
<a id="faq-interrupt-query"></a>

How do I interrupt a long-running query in an interactive shell? Normally the interactive shell becomes unresponsive to `Ctrl-C` when running a query. Using a connection in green mode allows Python to receive and handle the interrupt, although it may leave the connection broken, if the async callback doesn't handle the `KeyboardInterrupt` correctly.

Starting from psycopg 2.6.2, the `wait_select` callback can handle a `Ctrl-C` correctly. For previous versions, you can use `this implementation`__.
<a id="_"></a>

```pycon
 >>> psycopg2.extensions.set_wait_callback(psycopg2.extras.wait_select)
 >>> cnn = psycopg2.connect('')
 >>> cur = cnn.cursor()
 >>> cur.execute("select pg_sleep(10)")
 ^C
 Traceback (most recent call last):
   File "<stdin>", line 1, in <module>
   QueryCanceledError: canceling statement due to user request

 >>> cnn.rollback()
 >>> # You can use the connection and cursor again from here
```
<a id="faq-compile"></a>

## Problems compiling and installing psycopg2
<a id="faq-wheels"></a>

Psycopg 2.8 fails to install, Psycopg 2.7 was working fine. With Psycopg 2.7 you were installing binary packages, but they have proven unreliable so now you have to install them explicitly using the `psycopg2-binary` package. See [Quick Install](install.md#binary-packages) for all the details.
<a id="faq-python-h"></a>

I can't compile `psycopg2`: the compiler says *error: Python.h: No such file or directory*. What am I missing? You need to install a Python development package: it is usually called `python-dev` or `python3-dev` according to your Python version.
<a id="faq-libpq-fe-h"></a>

I can't compile `psycopg2`: the compiler says *error: libpq-fe.h: No such file or directory*. What am I missing? You need to install the development version of the libpq: the package is usually called `libpq-dev`.
<a id="faq-lo_truncate"></a>

`psycopg2` raises `ImportError` with message *_psycopg.so: undefined symbol: lo_truncate* when imported. This means that Psycopg was compiled with [`lo_truncate()`](https://www.postgresql.org/docs/current/static/lo-interfaces.html#LO-TRUNCATE) support (*i.e.* the libpq used at compile time was version >= 8.3) but at runtime an older libpq dynamic library is found.

Fast-forward several years, if the message reports *undefined symbol: lo_truncate64* it means that Psycopg was built with large objects 64 bits API support (*i.e.* the libpq used at compile time was at least 9.3) but at runtime an older libpq dynamic library is found.

You can use:

```shell
 $ ldd /path/to/packages/psycopg2/_psycopg.so | grep libpq
```

to find what is the libpq dynamic library used at runtime.

You can avoid the problem by using the same version of the `pg_config` at install time and the libpq at runtime.
<a id="lo_truncate"></a>
<a id="faq-import-mod_wsgi"></a>

Psycopg raises *ImportError: cannot import name tz* on import in mod_wsgi / ASP, but it works fine otherwise. If `psycopg2` is installed in an [egg](http://peak.telecommunity.com/DevCenter/PythonEggs) (e.g. because installed by `easy_install`), the user running the program may be unable to write in the `eggs cache`__. Set the env variable `PYTHON_EGG_CACHE` to a writable directory. With modwsgi you can use the WSGIPythonEggs__ directive.
<a id="egg"></a>
<a id="_"></a>
<a id="_"></a>
