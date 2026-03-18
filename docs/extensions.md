# psycopg2.extensions -- Extensions to the DB API

*module* `psycopg2.extensions`

The module contains a few objects and function extending the minimum set of functionalities defined by the DBAPI.

## Classes definitions

Instances of these classes are usually returned by factory functions or attributes. Their definitions are exposed here to allow subclassing, introspection etc.

*class* `connection(dsn, async=False)`

Is the class usually returned by the `connect()` function. It is exposed by the `extensions` module in order to allow subclassing to extend its behaviour: the subclass should be passed to the `connect()` function using the `connection_factory` parameter. See also [subclassing-connection](advanced.md#subclassing-connection).

For a complete description of the class, see `connection`.

*Changed in version 2.7.* *async_* can be used as alias for *async*.

*class* `cursor(conn, name=None)`

It is the class usually returned by the `connection.cursor()` method. It is exposed by the `extensions` module in order to allow subclassing to extend its behaviour: the subclass should be passed to the `cursor()` method using the `cursor_factory` parameter. See also [Connection and cursor factories](advanced.md#subclassing-cursor).

For a complete description of the class, see `cursor`.

*class* `lobject(conn [, oid [, mode [, new_oid [, new_file ]]]])`

Wrapper for a PostgreSQL large object. See [Access to PostgreSQL large objects](usage.md#large-objects) for an overview.

The class can be subclassed: see the `connection.lobject()` to know how to specify a `lobject` subclass.

*New in version 2.0.8.*

*attribute* `oid`

Database OID of the object.

*attribute* `mode`

The mode the database was open. See `connection.lobject()` for a description of the available modes.

*method* `read(bytes=-1)`

Read a chunk of data from the current file position. If -1 (default) read all the remaining data.

The result is an Unicode string (decoded according to `connection.encoding`) if the file was open in `t` mode, a bytes string for `b` mode.

*Changed in version 2.4.* added Unicode support.

*method* `write(str)`

Write a string to the large object. Return the number of bytes written. Unicode strings are encoded in the `connection.encoding` before writing.

*Changed in version 2.4.* added Unicode support.

*method* `export(file_name)`

Export the large object content to the file system.

The method uses the efficient [`lo_export()`](https://www.postgresql.org/docs/current/static/lo-interfaces.html#LO-EXPORT) libpq function.
<a id="lo_export"></a>

*method* `seek(offset, whence=0)`

Set the lobject current position.

*Changed in version 2.6.* added support for *offset* > 2GB.

*method* `tell()`

Return the lobject current position.

*New in version 2.2.*

*Changed in version 2.6.* added support for return value > 2GB.

*method* `truncate(len=0)`

Truncate the lobject to the given size.

The method will only be available if Psycopg has been built against libpq from PostgreSQL 8.3 or later and can only be used with PostgreSQL servers running these versions. It uses the [`lo_truncate()`](https://www.postgresql.org/docs/current/static/lo-interfaces.html#LO-TRUNCATE) libpq function.
<a id="lo_truncate"></a>

*New in version 2.2.*

*Changed in version 2.6.* added support for *len* > 2GB.

!!! warning

    If Psycopg is built with `lo_truncate()` support or with the 64 bits API support (resp. from PostgreSQL versions 8.3 and 9.3) but at runtime an older version of the dynamic library is found, the `psycopg2` module will fail to import.  See [the lo_truncate FAQ](faq.md#faq-lo_truncate) about the problem.

*method* `close()`

Close the object.

*attribute* `closed`

Boolean attribute specifying if the object is closed.

*method* `unlink()`

Close the object and remove it from the database.

*class* `ConnectionInfo(connection)`

*New in version 2.8.*

*attribute* `dbname`

*attribute* `user`

*attribute* `password`

*attribute* `host`

*attribute* `port`

*attribute* `options`

*attribute* `dsn_parameters`

Example:

```
>>> conn.info.dsn_parameters
{'dbname': 'test', 'user': 'postgres', 'port': '5432', 'sslmode': 'prefer'}
```

Requires libpq >= 9.3.

*attribute* `status`

*attribute* `transaction_status`

*method* `parameter_status(name)`

*attribute* `protocol_version`

Currently Psycopg supports only protocol 3, which allows connection to PostgreSQL server from version 7.4. Psycopg versions previous than 2.3 support both protocols 2 and 3.

*attribute* `server_version`

The number is formed by converting the major, minor, and revision numbers into two-decimal-digit numbers and appending them together. After PostgreSQL 10 the minor version was dropped, so the second group of digits is always `00`. For example, version 9.3.5 will be returned as `90305`, version 10.2 as `100002`.

*attribute* `error_message`

*attribute* `socket`

*attribute* `backend_pid`

*attribute* `needs_password`

*attribute* `used_password`

*attribute* `ssl_in_use`

*method* `ssl_attribute(name)`

*attribute* `ssl_attribute_names`

*class* `Column(*args, **kwargs)`

Description of one result column, exposed as items of the `cursor.description` sequence.

*New in version 2.8.*  in previous version the `description` attribute was a sequence of
 simple tuples or namedtuples.

*attribute* `name`

The name of the column returned.

*attribute* `type_code`

The PostgreSQL OID of the column. You can use the [`pg_type`](https://www.postgresql.org/docs/current/static/catalog-pg-type.html) system table to get more informations about the type.  This is the value used by Psycopg to decide what Python type use to represent the value.  See also [Type casting of SQL types into Python objects](advanced.md#type-casting-from-sql-to-python).

*attribute* `display_size`

Supposed to be the actual length of the column in bytes.  Obtaining this value is computationally intensive, so it is always `None`.

*Changed in version 2.8.* It was previously possible to obtain this value using a compiler
flag at builtin.

*attribute* `internal_size`

The size in bytes of the column associated to this column on the server. Set to a negative value for variable-size types See also [PQfsize](https://www.postgresql.org/docs/current/static/libpq-exec.html#LIBPQ-PQFSIZE).

*attribute* `precision`

Total number of significant digits in columns of type [`NUMERIC`](https://www.postgresql.org/docs/current/static/datatype-numeric.html#DATATYPE-NUMERIC-DECIMAL). `None` for other types.

*attribute* `scale`

Count of decimal digits in the fractional part in columns of type :sql:`NUMERIC`. `None` for other types.

*attribute* `null_ok`

Always `None` as not easy to retrieve from the libpq.

*attribute* `table_oid`

The oid of the table from which the column was fetched (matching `pg_class.oid`). `None` if the column is not a simple reference to a table column. See also [PQftable](https://www.postgresql.org/docs/current/static/libpq-exec.html#LIBPQ-PQFTABLE).

*New in version 2.8.*

*attribute* `table_column`

The number of the column (within its table) making up the result (matching `pg_attribute.attnum`, so it will start from 1). `None` if the column is not a simple reference to a table column. See also [PQftablecol](https://www.postgresql.org/docs/current/static/libpq-exec.html#LIBPQ-PQFTABLECOL).

*New in version 2.8.*
<a id="pg_type"></a>
<a id="PQgetlength"></a>
<a id="PQfsize"></a>
<a id="PQftable"></a>
<a id="PQftablecol"></a>
<a id="NUMERIC"></a>

*class* `Notify(pid, channel, payload='')`

*New in version 2.3.*

*class* `Xid(format_id, gtrid, bqual)`

*New in version 2.3.*

*method* `from_string(s)`

*class* `Diagnostics(exception)`

*New in version 2.5.*

The attributes currently available are:

column_name constraint_name context datatype_name internal_position internal_query message_detail message_hint message_primary schema_name severity severity_nonlocalized source_file source_function source_line sqlstate statement_position table_name

A string with the error field if available; `None` if not available. The attribute value is available only if the error sent by the server: not all the fields are available for all the errors and for all the server versions.

*New in version 2.8.* The `severity_nonlocalized` attribute.
<a id="sql-adaptation-objects"></a>

## SQL adaptation protocol objects

Psycopg provides a flexible system to adapt Python objects to the SQL syntax (inspired to the [PEP 246](https://peps.python.org/pep-246/)), allowing serialization in PostgreSQL. See [Adapting new Python types to SQL syntax](advanced.md#adapting-new-types) for a detailed description.  The following objects deal with Python objects adaptation:

*function* `adapt(obj)`

Return the SQL representation of *obj* as an `ISQLQuote`.  Raise a `ProgrammingError` if how to adapt the object is unknown. In order to allow new objects to be adapted, register a new adapter for it using the `register_adapter()` function.

The function is the entry point of the adaptation mechanism: it can be used to write adapters for complex objects by recursively calling `adapt()` on its components.

*function* `register_adapter(class, adapter)`

Register a new adapter for the objects of class *class*.

*adapter* should be a function taking a single argument (the object to adapt) and returning an object conforming to the `ISQLQuote` protocol (e.g. exposing a `getquoted()` method).  The `AsIs` is often useful for this task.

Once an object is registered, it can be safely used in SQL queries and by the `adapt()` function.

*class* `ISQLQuote(wrapped_object)`

Represents the SQL adaptation protocol.  Objects conforming this protocol should implement a `getquoted()` and optionally a `prepare()` method.

Adapters may subclass `ISQLQuote`, but is not necessary: it is enough to expose a `getquoted()` method to be conforming.

*attribute* `_wrapped`

The wrapped object passes to the constructor

*method* `getquoted()`

Subclasses or other conforming objects should return a valid SQL string representing the wrapped object. In Python 3 the SQL must be returned in a `bytes` object. The `ISQLQuote` implementation does nothing.

*method* `prepare(conn)`

Prepare the adapter for a connection.  The method is optional: if implemented, it will be invoked before `getquoted()` with the connection to adapt for as argument.

A conform object can implement this method if the SQL representation depends on any server parameter, such as the server version or the `standard_conforming_string` setting.  Container objects may store the connection and use it to recursively prepare contained objects: see the implementation for `psycopg2.extensions.SQL_IN` for a simple example.

*class* `AsIs(object)`

Adapter conform to the `ISQLQuote` protocol useful for objects whose string representation is already valid as SQL representation.

*method* `getquoted()`

Return the `str()` conversion of the wrapped object.

```python
>>> AsIs(42).getquoted()
'42'
```

*class* `QuotedString(str)`

Adapter conform to the `ISQLQuote` protocol for string-like objects.

*method* `getquoted()`

Return the string enclosed in single quotes. Any single quote appearing in the string is escaped by doubling it according to SQL string constants syntax. Backslashes are escaped too.

```python
>>> QuotedString(r"O'Reilly").getquoted()
"'O''Reilly'"
```

*class* `Binary(str)`

Adapter conform to the `ISQLQuote` protocol for binary objects.

*method* `getquoted()`

Return the string enclosed in single quotes.  It performs the same escaping of the `QuotedString` adapter, plus it knows how to escape non-printable chars.

```python
>>> Binary("\x00\x08\x0F").getquoted()
"'\\\\000\\\\010\\\\017'"
```

*Changed in version 2.0.14.* previously the adapter was not exposed by the `extensions`
module. In older versions it can be imported from the implementation
module `psycopg2._psycopg`.

*class* `Boolean`

Float SQL_IN

Specialized adapters for builtin objects.

*class* `DateFromPy`

TimeFromPy TimestampFromPy IntervalFromPy

Specialized adapters for Python datetime objects.

*data* `adapters`

Dictionary of the currently registered object adapters.  Use `register_adapter()` to add an adapter for a new type.

## Database types casting functions

These functions are used to manipulate type casters to convert from PostgreSQL types to Python objects.  See [Type casting of SQL types into Python objects](advanced.md#type-casting-from-sql-to-python) for details.

*function* `new_type(oids, name, adapter)`

Create a new type caster to convert from a PostgreSQL type to a Python object.  The object created must be registered using `register_type()` to be used.

**param oids**
:   tuple of OIDs of the PostgreSQL type to convert. :param name: the name of the new type adapter. :param adapter: the adaptation function.


The object OID can be read from the `cursor.description` attribute or by querying from the PostgreSQL catalog.

*adapter* should have signature `fun({value}, {cur})` where *value* is the string representation returned by PostgreSQL and *cur* is the cursor from which data are read. In case of `NULL`, *value* will be `None`. The adapter should return the converted object.

See [Type casting of SQL types into Python objects](advanced.md#type-casting-from-sql-to-python) for an usage example.

*function* `new_array_type(oids, name, base_caster)`

Create a new type caster to convert from a PostgreSQL array type to a list of Python object.  The object created must be registered using `register_type()` to be used.

**param oids**
:   tuple of OIDs of the PostgreSQL type to convert. It should probably contain the oid of the array type (e.g. the `typarray` field in the `pg_type` table). :param name: the name of the new type adapter. :param base_caster: a Psycopg typecaster, e.g. created using the `new_type()` function. The caster should be able to parse a single item of the desired type.


*New in version 2.4.3.*
<a id="cast-array-unknown"></a>

!!! note

    The function can be used to create a generic array typecaster, returning a list of strings: just use `psycopg2.STRING` as base typecaster. For instance, if you want to receive an array of `macaddr` from the database, each address represented by string, you can use:

    ```
    # select typarray from pg_type where typname = 'macaddr' -> 1040
    psycopg2.extensions.register_type(
        psycopg2.extensions.new_array_type(
            (1040,), 'MACADDR[]', psycopg2.STRING))
    ```

*function* `register_type(obj [, scope])`

Register a type caster created using `new_type()`.

If *scope* is specified, it should be a `connection` or a `cursor`: the type caster will be effective only limited to the specified object.  Otherwise it will be globally registered.

*data* `string_types`

The global register of type casters.

*data* `encodings`

Mapping from `PostgreSQL encoding`__ to `Python encoding`__ names. Used by Psycopg when adapting or casting unicode strings. See [unicode-handling](usage.md#unicode-handling).
<a id="_"></a>
<a id="_"></a>
<a id="extension-exceptions"></a>

## Additional exceptions

The module exports a few exceptions in addition to the [standard ones](module.md#dbapi-exceptions) defined by the DBAPI.

!!! note

    From psycopg 2.8 these error classes are also exposed by the `psycopg2.errors` module.

*exception* `QueryCanceledError`

(subclasses `OperationalError`)

Error related to SQL query cancellation.  It can be trapped specifically to detect a timeout.

*New in version 2.0.7.*

*exception* `TransactionRollbackError`

(subclasses `OperationalError`)

Error causing transaction rollback (deadlocks, serialization failures, etc).  It can be trapped specifically to detect a deadlock.

*New in version 2.0.7.*
<a id="coroutines-functions"></a>

## Coroutines support functions

These functions are used to set and retrieve the callback function for [cooperation with coroutine libraries](advanced.md#green-support).

*New in version 2.2.*

*function* `set_wait_callback(f)`

*function* `get_wait_callback()`

## Other functions

*function* `libpq_version()`

Return the version number of the `libpq` dynamic library loaded as an integer, in the same format of `server_version`.

Raise `NotSupportedError` if the `psycopg2` module was compiled with a `libpq` version lesser than 9.1 (which can be detected by the `__libpq_version__` constant).

*New in version 2.7.*

!!! tip "See Also"

    libpq docs for [PQlibVersion()](https://www.postgresql.org/docs/current/static/libpq-misc.html#LIBPQ-PQLIBVERSION).

*function* `make_dsn(dsn=None, **kwargs)`

Create a valid connection string from arguments.

Put together the arguments in *kwargs* into a connection string. If *dsn* is specified too, merge the arguments coming from both the sources. If the same argument name is specified in both the sources, the *kwargs* value overrides the *dsn* value.

The input arguments are validated: the output should always be a valid connection string (as far as `parse_dsn()` is concerned). If not raise `ProgrammingError`.

Example:

```
>>> from psycopg2.extensions import make_dsn
>>> make_dsn('dbname=foo host=example.com', password="s3cr3t")
'host=example.com password=s3cr3t dbname=foo'
```

*New in version 2.7.*

*function* `parse_dsn(dsn)`

Parse connection string into a dictionary of keywords and values.

Parsing is delegated to the libpq: different versions of the client library may support different formats or parameters (for example, `connection URIs`__ are only supported from libpq 9.2).  Raise `ProgrammingError` if the *dsn* is not valid.
<a id="_"></a>

Example:

```
>>> from psycopg2.extensions import parse_dsn
>>> parse_dsn('dbname=test user=postgres password=secret')
{'password': 'secret', 'user': 'postgres', 'dbname': 'test'}
>>> parse_dsn("postgresql://someone@example.com/somedb?connect_timeout=10")
{'host': 'example.com', 'user': 'someone', 'dbname': 'somedb', 'connect_timeout': '10'}
```

*New in version 2.7.*

!!! tip "See Also"

    libpq docs for [PQconninfoParse()](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-PQCONNINFOPARSE).

*function* `quote_ident(str, scope)`

Return quoted identifier according to PostgreSQL quoting rules.

The *scope* must be a `connection` or a `cursor`, the underlying connection encoding is used for any necessary character conversion.

*New in version 2.7.*

!!! tip "See Also"

    libpq docs for [PQescapeIdentifier()](https://www.postgresql.org/docs/current/static/libpq-exec.html#LIBPQ-PQESCAPEIDENTIFIER)

*method* `encrypt_password(password, user, scope=None, algorithm=None)`

Return the encrypted form of a PostgreSQL password.

**param password**
:   the cleartext password to encrypt :param user: the name of the user to use the password for :param scope: the scope to encrypt the password into; if *algorithm* is `md5` it can be `None` :type scope: `connection` or `cursor` :param algorithm: the password encryption algorithm to use


The *algorithm* `md5` is always supported. Other algorithms are only supported if the client libpq version is at least 10 and may require a compatible server version: check the `PostgreSQL encryption documentation`__ to know the algorithms supported by your server.
<a id="_"></a>

Using `None` as *algorithm* will result in querying the server to know the current server password encryption setting, which is a blocking operation: query the server separately and specify a value for *algorithm* if you want to maintain a non-blocking behaviour.

*New in version 2.8.*

!!! tip "See Also"

    PostgreSQL docs for the [password_encryption](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION) setting, libpq [PQencryptPasswordConn()](https://www.postgresql.org/docs/current/static/libpq-misc.html#LIBPQ-PQENCRYPTPASSWORDCONN), [PQencryptPassword()](https://www.postgresql.org/docs/current/static/libpq-misc.html#LIBPQ-PQENCRYPTPASSWORD) functions.
<a id="isolation-level-constants"></a>

## Isolation level constants

Psycopg2 `connection` objects hold informations about the PostgreSQL transaction isolation level  By default Psycopg doesn't change the default configuration of the server (`ISOLATION_LEVEL_DEFAULT`); the default for PostgreSQL servers is typically `READ COMMITTED`, but this may be changed in the server configuration files.  A different isolation level can be set through the `set_isolation_level()` or `set_session()` methods.  The level can be set to one of the following constants:

*data* `ISOLATION_LEVEL_AUTOCOMMIT`

No transaction is started when commands are executed and no `commit()` or `rollback()` is required. Some PostgreSQL command such as `CREATE DATABASE` or `VACUUM` can't run into a transaction: to run such command use:

```
>>> conn.set_isolation_level(ISOLATION_LEVEL_AUTOCOMMIT)
```

See also [Transactions control](usage.md#transactions-control).

*data* `ISOLATION_LEVEL_READ_UNCOMMITTED`

The `READ UNCOMMITTED` isolation level is defined in the SQL standard but not available in the |MVCC| model of PostgreSQL: it is replaced by the stricter `READ COMMITTED`.

*data* `ISOLATION_LEVEL_READ_COMMITTED`

This is usually the default PostgreSQL value, but a different default may be set in the database configuration.

A new transaction is started at the first `execute()` command on a cursor and at each new `execute()` after a `commit()` or a `rollback()`.  The transaction runs in the PostgreSQL `READ COMMITTED` isolation level: a `SELECT` query sees only data committed before the query began; it never sees either uncommitted data or changes committed during query execution by concurrent transactions.

!!! tip "See Also"

    [Read Committed Isolation Level](https://www.postgresql.org/docs/current/static/transaction-iso.html#XACT-READ-COMMITTED) in PostgreSQL documentation.

*data* `ISOLATION_LEVEL_REPEATABLE_READ`

As in `ISOLATION_LEVEL_READ_COMMITTED`, a new transaction is started at the first `execute()` command.  Transactions run at a `REPEATABLE READ` isolation level: all the queries in a transaction see a snapshot as of the start of the transaction, not as of the start of the current query within the transaction.  However applications using this level must be prepared to retry transactions due to serialization failures.

While this level provides a guarantee that each transaction sees a completely stable view of the database, this view will not necessarily always be consistent with some serial (one at a time) execution of concurrent transactions of the same level.

*Changed in version 2.4.2.* The value was an alias for `ISOLATION_LEVEL_SERIALIZABLE` before. The
two levels are distinct since PostgreSQL 9.1

!!! tip "See Also"

    [Repeatable Read Isolation Level](https://www.postgresql.org/docs/current/static/transaction-iso.html#XACT-REPEATABLE-READ) in PostgreSQL documentation.

*data* `ISOLATION_LEVEL_SERIALIZABLE`

As in `ISOLATION_LEVEL_READ_COMMITTED`, a new transaction is started at the first `execute()` command.  Transactions run at a `SERIALIZABLE` isolation level. This is the strictest transactions isolation level, equivalent to having the transactions executed serially rather than concurrently. However applications using this level must be prepared to retry transactions due to serialization failures.

Starting from PostgreSQL 9.1, this mode monitors for conditions which could make execution of a concurrent set of serializable transactions behave in a manner inconsistent with all possible serial (one at a time) executions of those transaction. In previous version the behaviour was the same of the `REPEATABLE READ` isolation level.

!!! tip "See Also"

    [Serializable Isolation Level](https://www.postgresql.org/docs/current/static/transaction-iso.html#XACT-SERIALIZABLE) in PostgreSQL documentation.

*data* `ISOLATION_LEVEL_DEFAULT`

A new transaction is started at the first `execute()` command, but the isolation level is not explicitly selected by Psycopg: the server will use whatever level is defined in its configuration or by statements executed within the session outside Pyscopg control.  If you want to know what the value is you can use a query such as `show transaction_isolation`.

*New in version 2.7.*
<a id="transaction-status-constants"></a>

## Transaction status constants

These values represent the possible status of a transaction: the current value can be read using the `connection.info.transaction_status` property.

*data* `TRANSACTION_STATUS_IDLE`

The session is idle and there is no current transaction.

*data* `TRANSACTION_STATUS_ACTIVE`

A command is currently in progress.

*data* `TRANSACTION_STATUS_INTRANS`

The session is idle in a valid transaction block.

*data* `TRANSACTION_STATUS_INERROR`

The session is idle in a failed transaction block.

*data* `TRANSACTION_STATUS_UNKNOWN`

Reported if the connection with the server is bad.
<a id="connection-status-constants"></a>

## Connection status constants

These values represent the possible status of a connection: the current value can be read from the `status` attribute.

It is possible to find the connection in other status than the one shown below. Those are the only states in which a working connection is expected to be found during the execution of regular Python client code: other states are for internal usage and Python code should not rely on them.

*data* `STATUS_READY`

Connection established. No transaction in progress.

*data* `STATUS_BEGIN`

Connection established. A transaction is currently in progress.

*data* `STATUS_IN_TRANSACTION`

An alias for `STATUS_BEGIN`

*data* `STATUS_PREPARED`

The connection has been prepared for the second phase in a [two-phase commit](usage.md#tpc) transaction. The connection can't be used to send commands to the database until the transaction is finished with `tpc_commit()` or `tpc_rollback()`.

*New in version 2.3.*
<a id="poll-constants"></a>

## Poll constants

*New in version 2.2.*

These values can be returned by `connection.poll()` during asynchronous connection and communication.  They match the values in the libpq enum `PostgresPollingStatusType`.  See [Asynchronous support](advanced.md#async-support) and [Support for coroutine libraries](advanced.md#green-support).

*data* `POLL_OK`

The data being read is available, or the file descriptor is ready for writing: reading or writing will not block.

*data* `POLL_READ`

Some data is being read from the backend, but it is not available yet on the client and reading would block.  Upon receiving this value, the client should wait for the connection file descriptor to be ready *for reading*. For example:

```
select.select([conn.fileno()], [], [])
```

*data* `POLL_WRITE`

Some data is being sent to the backend but the connection file descriptor can't currently accept new data.  Upon receiving this value, the client should wait for the connection file descriptor to be ready *for writing*. For example:

```
select.select([], [conn.fileno()], [])
```

*data* `POLL_ERROR`

There was a problem during connection polling. This value should actually never be returned: in case of poll error usually an exception containing the relevant details is raised.

## Additional database types

The `extensions` module includes typecasters for many standard PostgreSQL types.  These objects allow the conversion of returned data into Python objects.  All the typecasters are automatically registered, except `UNICODE` and `UNICODEARRAY`: you can register them using `register_type()` in order to receive Unicode objects instead of strings from the database.  See [unicode-handling](usage.md#unicode-handling) for details.

*data* `BOOLEAN`

BYTES DATE DECIMAL FLOAT INTEGER INTERVAL LONGINTEGER TIME UNICODE

Typecasters for basic types. Note that a few other ones (`BINARY`, `DATETIME`, `NUMBER`, `ROWID`, `STRING`) are exposed by the `psycopg2` module for DBAPI compliance.

*data* `BINARYARRAY`

BOOLEANARRAY BYTESARRAY DATEARRAY DATETIMEARRAY DECIMALARRAY FLOATARRAY INTEGERARRAY INTERVALARRAY LONGINTEGERARRAY ROWIDARRAY STRINGARRAY TIMEARRAY UNICODEARRAY

Typecasters to convert arrays of sql types into Python lists.

*data* `PYDATE`

PYDATETIME PYDATETIMETZ PYINTERVAL PYTIME PYDATEARRAY PYDATETIMEARRAY PYDATETIMETZARRAY PYINTERVALARRAY PYTIMEARRAY

Typecasters to convert time-related data types to Python `datetime` objects.

*Changed in version 2.2.* previously the `DECIMAL` typecaster and the specific time-related
typecasters (`PY*` and `MX*`) were not exposed by the `extensions`
module. In older versions they can be imported from the implementation
module `psycopg2._psycopg`.

*New in version 2.7.2.* the `*DATETIMETZ*` objects.

*New in version 2.8.* the `BYTES` and `BYTESARRAY` objects.
