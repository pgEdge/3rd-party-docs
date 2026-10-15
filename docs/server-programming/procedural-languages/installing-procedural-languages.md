<a id="xplang-install"></a>

## Installing Procedural Languages


 A procedural language must be “installed” into each database where it is to be used. But procedural languages installed in the database `template1` are automatically available in all subsequently created databases, since their entries in `template1` will be copied by `CREATE DATABASE`. So the database administrator can decide which languages are available in which databases and can make some languages available by default if desired.


 For the languages supplied with the standard distribution, it is only necessary to execute `CREATE EXTENSION` *language_name* to install the language into the current database. The manual procedure described below is only recommended for installing languages that have not been packaged as extensions.


**Manual Procedural Language Installation**

<a id="xplang-install-cr1"></a>
1.  The shared object for the language handler must be compiled and installed into an appropriate library directory. This works in the same way as building and installing modules with regular user-defined C functions does; see [Compiling and Linking Dynamically-Loaded Functions](../extending-sql/c-language-functions.md#dfunc). Often, the language handler will depend on an external library that provides the actual programming language engine; if so, that must be installed as well.
<a id="xplang-install-cr2"></a>
2.  The handler must be declared with the command

```

CREATE FUNCTION HANDLER_FUNCTION_NAME()
    RETURNS language_handler
    AS 'PATH-TO-SHARED-OBJECT'
    LANGUAGE C;
```
    The special return type of `language_handler` tells the database system that this function does not return one of the defined SQL data types and is not directly usable in SQL statements.
<a id="xplang-install-cr3"></a>
3.  Optionally, the language handler can provide an “inline” handler function that executes anonymous code blocks ([`DO`](../../reference/sql-commands/do.md#sql-do) commands) written in this language. If an inline handler function is provided by the language, declare it with a command like

```

CREATE FUNCTION INLINE_FUNCTION_NAME(internal)
    RETURNS void
    AS 'PATH-TO-SHARED-OBJECT'
    LANGUAGE C;
```

<a id="xplang-install-cr4"></a>
4.  Optionally, the language handler can provide a “validator” function that checks a function definition for correctness without actually executing it. The validator function is called by `CREATE FUNCTION` if it exists. If a validator function is provided by the language, declare it with a command like

```

CREATE FUNCTION VALIDATOR_FUNCTION_NAME(oid)
    RETURNS void
    AS 'PATH-TO-SHARED-OBJECT'
    LANGUAGE C STRICT;
```

<a id="xplang-install-cr5"></a>
5.  Finally, the PL must be declared with the command

```

CREATE [TRUSTED] LANGUAGE LANGUAGE_NAME
    HANDLER HANDLER_FUNCTION_NAME
    [INLINE INLINE_FUNCTION_NAME]
    [VALIDATOR VALIDATOR_FUNCTION_NAME] ;
```
    The optional key word `TRUSTED` specifies that the language does not grant access to data that the user would not otherwise have. Trusted languages are designed for ordinary database users (those without superuser privilege) and allows them to safely create functions and procedures. Since PL functions are executed inside the database server, the `TRUSTED` flag should only be given for languages that do not allow access to database server internals or the file system. The languages PL/pgSQL, PL/Tcl, and PL/Perl are considered trusted; the languages PL/TclU, PL/PerlU, and PL/PythonU are designed to provide unlimited functionality and should *not* be marked trusted.


 [Manual Installation of PL/Perl](#xplang-install-example) shows how the manual installation procedure would work with the language PL/Perl.
 <a id="xplang-install-example"></a>

**Example: Manual Installation of PL/Perl**


 The following command tells the database server where to find the shared object for the PL/Perl language's call handler function:

```sql

CREATE FUNCTION plperl_call_handler() RETURNS language_handler AS
    '$libdir/plperl' LANGUAGE C;
```


 PL/Perl has an inline handler function and a validator function, so we declare those too:

```sql

CREATE FUNCTION plperl_inline_handler(internal) RETURNS void AS
    '$libdir/plperl' LANGUAGE C STRICT;

CREATE FUNCTION plperl_validator(oid) RETURNS void AS
    '$libdir/plperl' LANGUAGE C STRICT;
```


 The command:

```sql

CREATE TRUSTED LANGUAGE plperl
    HANDLER plperl_call_handler
    INLINE plperl_inline_handler
    VALIDATOR plperl_validator;
```
 then defines that the previously declared functions should be invoked for functions and procedures where the language attribute is `plperl`.


 In a default PostgreSQL installation, the handler for the PL/pgSQL language is built and installed into the “library” directory; furthermore, the PL/pgSQL language itself is installed in all databases. If Tcl support is configured in, the handlers for PL/Tcl and PL/TclU are built and installed in the library directory, but the language itself is not installed in any database by default. Likewise, the PL/Perl and PL/PerlU handlers are built and installed if Perl support is configured, and the PL/PythonU handler is installed if Python support is configured, but these languages are not installed by default.
