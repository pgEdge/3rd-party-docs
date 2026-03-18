# psycopg2.errors -- Exception classes mapping PostgreSQL errors

*module* `psycopg2.errors`

*New in version 2.8.*

*Changed in version 2.8.4 added errors introduced in PostgreSQL 12.*

*Changed in version 2.8.6 added errors introduced in PostgreSQL 13.*

*Changed in version 2.9.2 added errors introduced in PostgreSQL 14.*

*Changed in version 2.9.4 added errors introduced in PostgreSQL 15.*

*Changed in version 2.9.10 added errors introduced in PostgreSQL 17.*

This module exposes the classes psycopg raises upon receiving an error from the database with a `SQLSTATE` value attached (available in the `pgcode` attribute). The content of the module is generated from the PostgreSQL source code and includes classes for every error defined by PostgreSQL in versions between 9.1 and 15.

Every class in the module is named after what referred as "condition name" `in the documentation`__, converted to CamelCase: e.g. the error 22012, `division_by_zero` is exposed by this module as the class `DivisionByZero`.
<a id="_"></a>

Every exception class is a subclass of one of the [standard DB-API exception](module.md#dbapi-exceptions) and expose the `Error` interface. Each class' superclass is what used to be raised by psycopg in versions before the introduction of this module, so everything should be compatible with previously written code catching one the DB-API class: if your code used to catch `IntegrityError` to detect a duplicate entry, it will keep on working even if a more specialised subclass such as `UniqueViolation` is raised.

The new classes allow a more idiomatic way to check and process a specific error among the many the database may return. For instance, in order to check that a table is locked, the following code could have been used previously:

```python
 try:
     cur.execute("LOCK TABLE mytable IN ACCESS EXCLUSIVE MODE NOWAIT")
 except psycopg2.OperationalError as e:
     if e.pgcode == psycopg2.errorcodes.LOCK_NOT_AVAILABLE:
         locked = True
     else:
         raise
```

While this method is still available, the specialised class allows for a more idiomatic error handler:

```python
 try:
     cur.execute("LOCK TABLE mytable IN ACCESS EXCLUSIVE MODE NOWAIT")
 except psycopg2.errors.LockNotAvailable:
     locked = True
```

*function* `lookup`

```python
 try:
     cur.execute("LOCK TABLE mytable IN ACCESS EXCLUSIVE MODE NOWAIT")
 except psycopg2.errors.lookup("55P03"):
     locked = True
```

## SQLSTATE exception classes

The following table contains the list of all the SQLSTATE classes exposed by the module.

Note that, for completeness, the module also exposes all the [DB-API-defined exceptions](module.md#dbapi-exceptions) and [a few psycopg-specific ones](extensions.md#extension-exceptions) exposed by the `extensions` module, which are not listed here.

*See: `sqlstate_errors.rst`*
