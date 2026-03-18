# psycopg2.errorcodes -- Error codes defined by PostgreSQL

*module* `psycopg2.errorcodes`

*New in version 2.0.6.*

This module contains symbolic names for all PostgreSQL error codes and error classes codes.  Subclasses of `Error` make the PostgreSQL error code available in the `pgcode` attribute.

From PostgreSQL documentation:

All messages emitted by the PostgreSQL server are assigned five-character error codes that follow the SQL standard's conventions for `SQLSTATE` codes.  Applications that need to know which error condition has occurred should usually test the error code, rather than looking at the textual error message.  The error codes are less likely to change across PostgreSQL releases, and also are not subject to change due to localization of error messages. Note that some, but not all, of the error codes produced by PostgreSQL are defined by the SQL standard; some additional error codes for conditions not defined by the standard have been invented or borrowed from other databases.

According to the standard, the first two characters of an error code denote a class of errors, while the last three characters indicate a specific condition within that class. Thus, an application that does not recognize the specific error code can still be able to infer what to do from the error class.

!!! tip "See Also"

    [PostgreSQL Error Codes table](https://www.postgresql.org/docs/current/static/errcodes-appendix.html#ERRCODES-TABLE)

An example of the available constants defined in the module:

```python
>>> errorcodes.CLASS_SYNTAX_ERROR_OR_ACCESS_RULE_VIOLATION
'42'
>>> errorcodes.UNDEFINED_TABLE
'42P01'
```

Constants representing all the error values defined by PostgreSQL versions between 8.1 and 15 are included in the module.

*function* `lookup(code)`

```python
 >>> try:
 ...     cur.execute("SELECT ouch FROM aargh;")
 ... except Exception as e:
 ...     pass
 ...
 >>> errorcodes.lookup(e.pgcode[:2])
 'CLASS_SYNTAX_ERROR_OR_ACCESS_RULE_VIOLATION'
 >>> errorcodes.lookup(e.pgcode)
 'UNDEFINED_TABLE'
```

*New in version 2.0.14.*
