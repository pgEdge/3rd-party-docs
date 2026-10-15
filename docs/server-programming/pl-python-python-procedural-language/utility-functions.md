<a id="plpython-util"></a>

## Utility Functions


 The `plpy` module also provides the functions

- <code>plpy.debug(</code><em>msg, **kwargs</em><code>)</code>
- <code>plpy.log(</code><em>msg, **kwargs</em><code>)</code>
- <code>plpy.info(</code><em>msg, **kwargs</em><code>)</code>
- <code>plpy.notice(</code><em>msg, **kwargs</em><code>)</code>
- <code>plpy.warning(</code><em>msg, **kwargs</em><code>)</code>
- <code>plpy.error(</code><em>msg, **kwargs</em><code>)</code>
- <code>plpy.fatal(</code><em>msg, **kwargs</em><code>)</code>
  `plpy.error` and `plpy.fatal` actually raise a Python exception which, if uncaught, propagates out to the calling query, causing the current transaction or subtransaction to be aborted. <code>raise plpy.Error(</code><em>msg</em><code>)</code> and <code>raise plpy.Fatal(</code><em>msg</em><code>)</code> are equivalent to calling <code>plpy.error(</code><em>msg</em><code>)</code> and <code>plpy.fatal(</code><em>msg</em><code>)</code>, respectively but the `raise` form does not allow passing keyword arguments. The other functions only generate messages of different priority levels. Whether messages of a particular priority are reported to the client, written to the server log, or both is controlled by the [log_min_messages](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-min-messages) and [client_min_messages](../../server-administration/server-configuration/client-connection-defaults.md#guc-client-min-messages) configuration variables. See [Server Configuration](../../server-administration/server-configuration/index.md#runtime-config) for more information.


 The *msg* argument is given as a positional argument. For backward compatibility, more than one positional argument can be given. In that case, the string representation of the tuple of positional arguments becomes the message reported to the client.


 The following keyword-only arguments are accepted:

- `detail`
- `hint`
- `sqlstate`
- `schema_name`
- `table_name`
- `column_name`
- `datatype_name`
- `constraint_name`
 The string representation of the objects passed as keyword-only arguments is used to enrich the messages reported to the client. For example:

```sql

CREATE FUNCTION raise_custom_exception() RETURNS void AS $$
plpy.error("custom exception message",
           detail="some info about exception",
           hint="hint for users")
$$ LANGUAGE plpython3u;

=# SELECT raise_custom_exception();
ERROR:  plpy.Error: custom exception message
DETAIL:  some info about exception
HINT:  hint for users
CONTEXT:  Traceback (most recent call last):
  PL/Python function "raise_custom_exception", line 4, in
    hint="hint for users")
PL/Python function "raise_custom_exception"
```


 Another set of utility functions are <code>plpy.quote_literal(</code><em>string</em><code>)</code>, <code>plpy.quote_nullable(</code><em>string</em><code>)</code>, and <code>plpy.quote_ident(</code><em>string</em><code>)</code>. They are equivalent to the built-in quoting functions described in [String Functions and Operators](../../the-sql-language/functions-and-operators/string-functions-and-operators.md#functions-string). They are useful when constructing ad-hoc queries. A PL/Python equivalent of dynamic SQL from [Quoting Values in Dynamic Queries](../pl-pgsql-sql-procedural-language/basic-statements.md#plpgsql-quote-literal-example) would be:

```

plpy.execute("UPDATE tbl SET %s = %s WHERE key = %s" % (
    plpy.quote_ident(colname),
    plpy.quote_nullable(newvalue),
    plpy.quote_literal(keyvalue)))
```
