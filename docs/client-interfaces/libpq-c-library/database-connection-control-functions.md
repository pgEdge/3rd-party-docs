<a id="libpq-connect"></a>

## Database Connection Control Functions


 The following functions deal with making a connection to a PostgreSQL backend server. An application program can have several backend connections open at one time. (One reason to do that is to access more than one database.) Each connection is represented by a `PGconn` object, which is obtained from the function [PQconnectdb](#libpq-PQconnectdb), [PQconnectdbParams](#libpq-PQconnectdbParams), or [PQsetdbLogin](#libpq-PQsetdbLogin). Note that these functions will always return a non-null object pointer, unless perhaps there is too little memory even to allocate the `PGconn` object. The [PQstatus](connection-status-functions.md#libpq-PQstatus) function should be called to check the return value for a successful connection before queries are sent via the connection object.

!!! warning

    If untrusted users have access to a database that has not adopted a [secure schema usage pattern](../../the-sql-language/data-definition/schemas.md#ddl-schemas-patterns), begin each session by removing publicly-writable schemas from `search_path`. One can set parameter key word `options` to value `-csearch_path=`. Alternately, one can issue <code>PQexec(</code><em>conn</em><code>, "SELECT
         pg_catalog.set_config('search_path', '', false)")</code> after connecting. This consideration is not specific to libpq; it applies to every interface for executing arbitrary SQL commands.


!!! warning

    On Unix, forking a process with open libpq connections can lead to unpredictable results because the parent and child processes share the same sockets and operating system resources. For this reason, such usage is not recommended, though doing an `exec` from the child process to load a new executable is safe.


<a id="libpq-PQconnectdbParams"></a>

`PQconnectdbParams`
:   Makes a new connection to the database server.

    ```

    PGconn *PQconnectdbParams(const char * const *keywords,
                              const char * const *values,
                              int expand_dbname);
    ```


     This function opens a new database connection using the parameters taken from two `NULL`-terminated arrays. The first, `keywords`, is defined as an array of strings, each one being a key word. The second, `values`, gives the value for each key word. Unlike [PQsetdbLogin](#libpq-PQsetdbLogin) below, the parameter set can be extended without changing the function signature, so use of this function (or its nonblocking analogs [PQconnectStartParams](#libpq-PQconnectStartParams) and `PQconnectPoll`) is preferred for new application programming.


     The currently recognized parameter key words are listed in [Parameter Key Words](#libpq-paramkeywords).


     The passed arrays can be empty to use all default parameters, or can contain one or more parameter settings. They must be matched in length. Processing will stop at the first `NULL` entry in the `keywords` array. Also, if the `values` entry associated with a non-`NULL` `keywords` entry is `NULL` or an empty string, that entry is ignored and processing continues with the next pair of array entries.


     When `expand_dbname` is non-zero, the value for the first `dbname` key word is checked to see if it is a *connection string*. If so, it is “expanded” into the individual connection parameters extracted from the string. The value is considered to be a connection string, rather than just a database name, if it contains an equal sign (`=`) or it begins with a URI scheme designator. (More details on connection string formats appear in [Connection Strings](#libpq-connstring).) Only the first occurrence of `dbname` is treated in this way; any subsequent `dbname` parameter is processed as a plain database name.


     In general the parameter arrays are processed from start to end. If any key word is repeated, the last value (that is not `NULL` or empty) is used. This rule applies in particular when a key word found in a connection string conflicts with one appearing in the `keywords` array. Thus, the programmer may determine whether array entries can override or be overridden by values taken from a connection string. Array entries appearing before an expanded `dbname` entry can be overridden by fields of the connection string, and in turn those fields are overridden by array entries appearing after `dbname` (but, again, only if those entries supply non-empty values).


     After processing all the array entries and any expanded connection string, any connection parameters that remain unset are filled with default values. If an unset parameter's corresponding environment variable (see [Environment Variables](environment-variables.md#libpq-envars)) is set, its value is used. If the environment variable is not set either, then the parameter's built-in default value is used.
<a id="libpq-PQconnectdb"></a>

`PQconnectdb`
:   Makes a new connection to the database server.

    ```

    PGconn *PQconnectdb(const char *conninfo);
    ```


     This function opens a new database connection using the parameters taken from the string `conninfo`.


     The passed string can be empty to use all default parameters, or it can contain one or more parameter settings separated by whitespace, or it can contain a URI. See [Connection Strings](#libpq-connstring) for details.
<a id="libpq-PQsetdbLogin"></a>

`PQsetdbLogin`
:   Makes a new connection to the database server.

    ```

    PGconn *PQsetdbLogin(const char *pghost,
                         const char *pgport,
                         const char *pgoptions,
                         const char *pgtty,
                         const char *dbName,
                         const char *login,
                         const char *pwd);
    ```


     This is the predecessor of [PQconnectdb](#libpq-PQconnectdb) with a fixed set of parameters. It has the same functionality except that the missing parameters will always take on default values. Write `NULL` or an empty string for any one of the fixed parameters that is to be defaulted.


     If the `dbName` contains an `=` sign or has a valid connection URI prefix, it is taken as a `conninfo` string in exactly the same way as if it had been passed to [PQconnectdb](#libpq-PQconnectdb), and the remaining parameters are then applied as specified for [PQconnectdbParams](#libpq-PQconnectdbParams).


     `pgtty` is no longer used and any value passed will be ignored.
<a id="libpq-PQsetdb"></a>

`PQsetdb`
:   Makes a new connection to the database server.

    ```

    PGconn *PQsetdb(char *pghost,
                    char *pgport,
                    char *pgoptions,
                    char *pgtty,
                    char *dbName);
    ```


     This is a macro that calls [PQsetdbLogin](#libpq-PQsetdbLogin) with null pointers for the `login` and `pwd` parameters. It is provided for backward compatibility with very old programs.
<a id="libpq-PQconnectStartParams"></a>

`PQconnectStartParams`, `PQconnectStart`<a id="libpq-PQconnectPoll"></a>
, `PQconnectPoll`
:   Make a connection to the database server in a nonblocking manner.

    ```

    PGconn *PQconnectStartParams(const char * const *keywords,
                                 const char * const *values,
                                 int expand_dbname);

    PGconn *PQconnectStart(const char *conninfo);

    PostgresPollingStatusType PQconnectPoll(PGconn *conn);
    ```


     These three functions are used to open a connection to a database server such that your application's thread of execution is not blocked on remote I/O whilst doing so. The point of this approach is that the waits for I/O to complete can occur in the application's main loop, rather than down inside [PQconnectdbParams](#libpq-PQconnectdbParams) or [PQconnectdb](#libpq-PQconnectdb), and so the application can manage this operation in parallel with other activities.


     With [PQconnectStartParams](#libpq-PQconnectStartParams), the database connection is made using the parameters taken from the `keywords` and `values` arrays, and controlled by `expand_dbname`, as described above for [PQconnectdbParams](#libpq-PQconnectdbParams).


     With `PQconnectStart`, the database connection is made using the parameters taken from the string `conninfo` as described above for [PQconnectdb](#libpq-PQconnectdb).


     Neither [PQconnectStartParams](#libpq-PQconnectStartParams) nor `PQconnectStart` nor `PQconnectPoll` will block, so long as a number of restrictions are met:

    -  The `hostaddr` parameter must be used appropriately to prevent DNS queries from being made. See the documentation of this parameter in [Parameter Key Words](#libpq-paramkeywords) for details.
    -  If you call [PQtrace](control-functions.md#libpq-PQtrace), ensure that the stream object into which you trace will not block.
    -  You must ensure that the socket is in the appropriate state before calling `PQconnectPoll`, as described below.


     To begin a nonblocking connection request, call `PQconnectStart` or [PQconnectStartParams](#libpq-PQconnectStartParams). If the result is null, then libpq has been unable to allocate a new `PGconn` structure. Otherwise, a valid `PGconn` pointer is returned (though not yet representing a valid connection to the database). Next call `PQstatus(conn)`. If the result is `CONNECTION_BAD`, the connection attempt has already failed, typically because of invalid connection parameters.


     If `PQconnectStart` or [PQconnectStartParams](#libpq-PQconnectStartParams) succeeds, the next stage is to poll libpq so that it can proceed with the connection sequence. Use `PQsocket(conn)` to obtain the descriptor of the socket underlying the database connection. (Caution: do not assume that the socket remains the same across `PQconnectPoll` calls.) Loop thus: If `PQconnectPoll(conn)` last returned `PGRES_POLLING_READING`, wait until the socket is ready to read (as indicated by `select()`, `poll()`, or similar system function). Note that `PQsocketPoll` can help reduce boilerplate by abstracting the setup of `select(2)` or `poll(2)` if it is available on your system. Then call `PQconnectPoll(conn)` again. Conversely, if `PQconnectPoll(conn)` last returned `PGRES_POLLING_WRITING`, wait until the socket is ready to write, then call `PQconnectPoll(conn)` again. On the first iteration, i.e., if you have yet to call `PQconnectPoll`, behave as if it last returned `PGRES_POLLING_WRITING`. Continue this loop until `PQconnectPoll(conn)` returns `PGRES_POLLING_FAILED`, indicating the connection procedure has failed, or `PGRES_POLLING_OK`, indicating the connection has been successfully made.


     At any time during connection, the status of the connection can be checked by calling [PQstatus](connection-status-functions.md#libpq-PQstatus). If this call returns `CONNECTION_BAD`, then the connection procedure has failed; if the call returns `CONNECTION_OK`, then the connection is ready. Both of these states are equally detectable from the return value of `PQconnectPoll`, described above. Other states might also occur during (and only during) an asynchronous connection procedure. These indicate the current stage of the connection procedure and might be useful to provide feedback to the user for example. These statuses are:

    <a id="libpq-connection-started"></a>

    `CONNECTION_STARTED`
    :   Waiting for connection to be made.
    <a id="libpq-connection-made"></a>

    `CONNECTION_MADE`
    :   Connection OK; waiting to send.
    <a id="libpq-connection-awaiting-response"></a>

    `CONNECTION_AWAITING_RESPONSE`
    :   Waiting for a response from the server.
    <a id="libpq-connection-auth-ok"></a>

    `CONNECTION_AUTH_OK`
    :   Received authentication; waiting for backend start-up to finish.
    <a id="libpq-connection-ssl-startup"></a>

    `CONNECTION_SSL_STARTUP`
    :   Negotiating SSL encryption.
    <a id="libpq-connection-gss-startup"></a>

    `CONNECTION_GSS_STARTUP`
    :   Negotiating GSS encryption.
    <a id="libpq-connection-check-writable"></a>

    `CONNECTION_CHECK_WRITABLE`
    :   Checking if connection is able to handle write transactions.
    <a id="libpq-connection-check-standby"></a>

    `CONNECTION_CHECK_STANDBY`
    :   Checking if connection is to a server in standby mode.
    <a id="libpq-connection-consume"></a>

    `CONNECTION_CONSUME`
    :   Consuming any remaining response messages on connection.
     Note that, although these constants will remain (in order to maintain compatibility), an application should never rely upon these occurring in a particular order, or at all, or on the status always being one of these documented values. An application might do something like this:

    ```

    switch(PQstatus(conn))
    {
            case CONNECTION_STARTED:
                feedback = "Connecting...";
                break;

            case CONNECTION_MADE:
                feedback = "Connected to server...";
                break;
    .
    .
    .
            default:
                feedback = "Connecting...";
    }
    ```


     The `connect_timeout` connection parameter is ignored when using `PQconnectPoll`; it is the application's responsibility to decide whether an excessive amount of time has elapsed. Otherwise, `PQconnectStart` followed by a `PQconnectPoll` loop is equivalent to [PQconnectdb](#libpq-PQconnectdb).


     Note that when `PQconnectStart` or [PQconnectStartParams](#libpq-PQconnectStartParams) returns a non-null pointer, you must call [PQfinish](#libpq-PQfinish) when you are finished with it, in order to dispose of the structure and any associated memory blocks. This must be done even if the connection attempt fails or is abandoned.
<a id="libpq-PQsocketPoll"></a>

`PQsocketPoll`
:   Poll a connection's underlying socket descriptor retrieved with [PQsocket](connection-status-functions.md#libpq-PQsocket). The primary use of this function is iterating through the connection sequence described in the documentation of [PQconnectStartParams](#libpq-PQconnectStartParams).

    ```

    typedef int64_t pg_usec_time_t;

    int PQsocketPoll(int sock, int forRead, int forWrite,
                     pg_usec_time_t end_time);
    ```


     This function performs polling of a file descriptor, optionally with a timeout. If `forRead` is nonzero, the function will terminate when the socket is ready for reading. If `forWrite` is nonzero, the function will terminate when the socket is ready for writing.


     The timeout is specified by `end_time`, which is the time to stop waiting expressed as a number of microseconds since the Unix epoch (that is, `time_t` times 1 million). Timeout is infinite if `end_time` is `-1`. Timeout is immediate (no blocking) if `end_time` is `0` (or indeed, any time before now). Timeout values can be calculated conveniently by adding the desired number of microseconds to the result of [PQgetCurrentTimeUSec](miscellaneous-functions.md#libpq-PQgetCurrentTimeUSec). Note that the underlying system calls may have less than microsecond precision, so that the actual delay may be imprecise.


     The function returns a value greater than `0` if the specified condition is met, `0` if a timeout occurred, or `-1` if an error occurred. The error can be retrieved by checking the `errno(3)` value. In the event both `forRead` and `forWrite` are zero, the function immediately returns a timeout indication.


     `PQsocketPoll` is implemented using either `poll(2)` or `select(2)`, depending on platform. See `POLLIN` and `POLLOUT` from `poll(2)`, or `readfds` and `writefds` from `select(2)`, for more information.
<a id="libpq-PQconndefaults"></a>

`PQconndefaults`
:   Returns the default connection options.

    ```

    PQconninfoOption *PQconndefaults(void);

    typedef struct
    {
        char   *keyword;   /* The keyword of the option */
        char   *envvar;    /* Fallback environment variable name */
        char   *compiled;  /* Fallback compiled in default value */
        char   *val;       /* Option's current value, or NULL */
        char   *label;     /* Label for field in connect dialog */
        char   *dispchar;  /* Indicates how to display this field
                              in a connect dialog. Values are:
                              ""        Display entered value as is
                              "*"       Password field - hide value
                              "D"       Debug option - don't show by default */
        int     dispsize;  /* Field size in characters for dialog */
    } PQconninfoOption;
    ```


     Returns a connection options array. This can be used to determine all possible [PQconnectdb](#libpq-PQconnectdb) options and their current default values. The return value points to an array of `PQconninfoOption` structures, which ends with an entry having a null `keyword` pointer. The null pointer is returned if memory could not be allocated. Note that the current default values (`val` fields) will depend on environment variables and other context. A missing or invalid service file will be silently ignored. Callers must treat the connection options data as read-only.


     After processing the options array, free it by passing it to [PQconninfoFree](miscellaneous-functions.md#libpq-PQconninfoFree). If this is not done, a small amount of memory is leaked for each call to [PQconndefaults](#libpq-PQconndefaults).
<a id="libpq-PQconninfo"></a>

`PQconninfo`
:   Returns the connection options used by a live connection.

    ```

    PQconninfoOption *PQconninfo(PGconn *conn);
    ```


     Returns a connection options array. This can be used to determine all possible [PQconnectdb](#libpq-PQconnectdb) options and the values that were used to connect to the server. The return value points to an array of `PQconninfoOption` structures, which ends with an entry having a null `keyword` pointer. All notes above for [PQconndefaults](#libpq-PQconndefaults) also apply to the result of [PQconninfo](#libpq-PQconninfo).
<a id="libpq-PQconninfoParse"></a>

`PQconninfoParse`
:   Returns parsed connection options from the provided connection string.

    ```

    PQconninfoOption *PQconninfoParse(const char *conninfo, char **errmsg);
    ```


     Parses a connection string and returns the resulting options as an array; or returns `NULL` if there is a problem with the connection string. This function can be used to extract the [PQconnectdb](#libpq-PQconnectdb) options in the provided connection string. The return value points to an array of `PQconninfoOption` structures, which ends with an entry having a null `keyword` pointer.


     All legal options will be present in the result array, but the `PQconninfoOption` for any option not present in the connection string will have `val` set to `NULL`; default values are not inserted.


     If `errmsg` is not `NULL`, then `*errmsg` is set to `NULL` on success, else to a `malloc`'d error string explaining the problem. (It is also possible for `*errmsg` to be set to `NULL` and the function to return `NULL`; this indicates an out-of-memory condition.)


     After processing the options array, free it by passing it to [PQconninfoFree](miscellaneous-functions.md#libpq-PQconninfoFree). If this is not done, some memory is leaked for each call to [PQconninfoParse](#libpq-PQconninfoParse). Conversely, if an error occurs and `errmsg` is not `NULL`, be sure to free the error string using [PQfreemem](miscellaneous-functions.md#libpq-PQfreemem).
<a id="libpq-PQfinish"></a>

`PQfinish`
:   Closes the connection to the server. Also frees memory used by the `PGconn` object.

    ```

    void PQfinish(PGconn *conn);
    ```


     Note that even if the server connection attempt fails (as indicated by [PQstatus](connection-status-functions.md#libpq-PQstatus)), the application should call [PQfinish](#libpq-PQfinish) to free the memory used by the `PGconn` object. The `PGconn` pointer must not be used again after [PQfinish](#libpq-PQfinish) has been called.
<a id="libpq-PQreset"></a>

`PQreset`
:   Resets the communication channel to the server.

    ```

    void PQreset(PGconn *conn);
    ```


     This function will close the connection to the server and attempt to establish a new connection, using all the same parameters previously used. This might be useful for error recovery if a working connection is lost.
<a id="libpq-PQresetStart"></a>

`PQresetStart`, `PQresetPoll`
:   Reset the communication channel to the server, in a nonblocking manner.

    ```

    int PQresetStart(PGconn *conn);

    PostgresPollingStatusType PQresetPoll(PGconn *conn);
    ```


     These functions will close the connection to the server and attempt to establish a new connection, using all the same parameters previously used. This can be useful for error recovery if a working connection is lost. They differ from [PQreset](#libpq-PQreset) (above) in that they act in a nonblocking manner. These functions suffer from the same restrictions as [PQconnectStartParams](#libpq-PQconnectStartParams), `PQconnectStart` and `PQconnectPoll`.


     To initiate a connection reset, call [PQresetStart](#libpq-PQresetStart). If it returns 0, the reset has failed. If it returns 1, poll the reset using `PQresetPoll` in exactly the same way as you would create the connection using `PQconnectPoll`.
<a id="libpq-PQpingParams"></a>

`PQpingParams`
:   [PQpingParams](#libpq-PQpingParams) reports the status of the server. It accepts connection parameters identical to those of [PQconnectdbParams](#libpq-PQconnectdbParams), described above. It is not necessary to supply correct user name, password, or database name values to obtain the server status; however, if incorrect values are provided, the server will log a failed connection attempt.

    ```

    PGPing PQpingParams(const char * const *keywords,
                        const char * const *values,
                        int expand_dbname);
    ```
     The function returns one of the following values:

    <a id="libpq-PQpingParams-PQPING_OK"></a>

    `PQPING_OK`
    :   The server is running and appears to be accepting connections.
    <a id="libpq-PQpingParams-PQPING_REJECT"></a>

    `PQPING_REJECT`
    :   The server is running but is in a state that disallows connections (startup, shutdown, or crash recovery).
    <a id="libpq-PQpingParams-PQPING_NO_RESPONSE"></a>

    `PQPING_NO_RESPONSE`
    :   The server could not be contacted. This might indicate that the server is not running, or that there is something wrong with the given connection parameters (for example, wrong port number), or that there is a network connectivity problem (for example, a firewall blocking the connection request).
    <a id="libpq-PQpingParams-PQPING_NO_ATTEMPT"></a>

    `PQPING_NO_ATTEMPT`
    :   No attempt was made to contact the server, because the supplied parameters were obviously incorrect or there was some client-side problem (for example, out of memory).
<a id="libpq-PQping"></a>

`PQping`
:   [PQping](#libpq-PQping) reports the status of the server. It accepts connection parameters identical to those of [PQconnectdb](#libpq-PQconnectdb), described above. It is not necessary to supply correct user name, password, or database name values to obtain the server status; however, if incorrect values are provided, the server will log a failed connection attempt.

    ```

    PGPing PQping(const char *conninfo);
    ```


     The return values are the same as for [PQpingParams](#libpq-PQpingParams).
<a id="libpq-pqsetsslkeypasshook-openssl"></a>

`PQsetSSLKeyPassHook_OpenSSL`
:   `PQsetSSLKeyPassHook_OpenSSL` lets an application override libpq's [default handling of encrypted client certificate key files](ssl-support.md#libpq-ssl-clientcert) using [sslpassword](#libpq-connect-sslpassword) or interactive prompting.

    ```

    void PQsetSSLKeyPassHook_OpenSSL(PQsslKeyPassHook_OpenSSL_type hook);
    ```
     The application passes a pointer to a callback function with signature:

    ```

    int callback_fn(char *buf, int size, PGconn *conn);
    ```
     which libpq will then call *instead of* its default `PQdefaultSSLKeyPassHook_OpenSSL` handler. The callback should determine the password for the key and copy it to result-buffer `buf` of size `size`. The string in `buf` must be null-terminated. The callback must return the length of the password stored in `buf` excluding the null terminator. On failure, the callback should set `buf[0] = '\0'` and return 0. See `PQdefaultSSLKeyPassHook_OpenSSL` in libpq's source code for an example.


     If the user specified an explicit key location, its path will be in `conn->sslkey` when the callback is invoked. This will be empty if the default key path is being used. For keys that are engine specifiers, it is up to engine implementations whether they use the OpenSSL password callback or define their own handling.


     The app callback may choose to delegate unhandled cases to `PQdefaultSSLKeyPassHook_OpenSSL`, or call it first and try something else if it returns 0, or completely override it.


     The callback *must not* escape normal flow control with exceptions, `longjmp(...)`, etc. It must return normally.
<a id="libpq-pqgetsslkeypasshook-openssl"></a>

`PQgetSSLKeyPassHook_OpenSSL`
:   `PQgetSSLKeyPassHook_OpenSSL` returns the current client certificate key password hook, or `NULL` if none has been set.

    ```

    PQsslKeyPassHook_OpenSSL_type PQgetSSLKeyPassHook_OpenSSL(void);
    ```

 <a id="libpq-connstring"></a>

### Connection Strings


 Several libpq functions parse a user-specified string to obtain connection parameters. There are two accepted formats for these strings: plain keyword/value strings and URIs. URIs generally follow [RFC 3986](https://datatracker.ietf.org/doc/html/rfc3986), except that multi-host connection strings are allowed as further described below.
 <a id="libpq-connstring-keyword-value"></a>

#### Keyword/Value Connection Strings


 In the keyword/value format, each parameter setting is in the form *keyword* `=` *value*, with space(s) between settings. Spaces around a setting's equal sign are optional. To write an empty value, or a value containing spaces, surround it with single quotes, for example `keyword = 'a value'`. Single quotes and backslashes within a value must be escaped with a backslash, i.e., `\'` and `\\`.


 Example:

```

host=localhost port=5432 dbname=mydb connect_timeout=10
```


 The recognized parameter key words are listed in [Parameter Key Words](#libpq-paramkeywords).
  <a id="libpq-connstring-uris"></a>

#### Connection URIs


 The general form for a connection URI is:

```

postgresql://[USERSPEC@][HOSTSPEC][/DBNAME][?PARAMSPEC]

where USERSPEC is:

USER[:PASSWORD]

and HOSTSPEC is:

[HOST][:PORT][,...]

and PARAMSPEC is:

NAME=VALUE[&...]
```


 The URI scheme designator can be either `postgresql://` or `postgres://`. Each of the remaining URI parts is optional. The following examples illustrate valid URI syntax:

```

postgresql://
postgresql://localhost
postgresql://localhost:5433
postgresql://localhost/mydb
postgresql://user@localhost
postgresql://user:secret@localhost
postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
postgresql://host1:123,host2:456/somedb?target_session_attrs=any&application_name=myapp
```
 Values that would normally appear in the hierarchical part of the URI can alternatively be given as named parameters. For example:

```

postgresql:///mydb?host=localhost&port=5433
```
 All named parameters must match key words listed in [Parameter Key Words](#libpq-paramkeywords), except that for compatibility with JDBC connection URIs, instances of `ssl=true` are translated into `sslmode=require`.


 The connection URI needs to be encoded with [percent-encoding](https://datatracker.ietf.org/doc/html/rfc3986#section-2.1) if it includes symbols with special meaning in any of its parts. Here is an example where the equal sign (`=`) is replaced with `%3D` and the space character with `%20`:

```

postgresql://user@localhost:5433/mydb?options=-c%20synchronous_commit%3Doff
```


 The host part may be either a host name or an IP address. To specify an IPv6 address, enclose it in square brackets:

```

postgresql://[2001:db8::1234]/database
```


 The host part is interpreted as described for the parameter [host](#libpq-connect-host). In particular, a Unix-domain socket connection is chosen if the host part is either empty or looks like an absolute path name, otherwise a TCP/IP connection is initiated. Note, however, that the slash is a reserved character in the hierarchical part of the URI. So, to specify a non-standard Unix-domain socket directory, either omit the host part of the URI and specify the host as a named parameter, or percent-encode the path in the host part of the URI:

```

postgresql:///dbname?host=/var/lib/postgresql
postgresql://%2Fvar%2Flib%2Fpostgresql/dbname
```


 It is possible to specify multiple host components, each with an optional port component, in a single URI. A URI of the form `postgresql://host1:port1,host2:port2,host3:port3/` is equivalent to a connection string of the form `host=host1,host2,host3 port=port1,port2,port3`. As further described below, each host will be tried in turn until a connection is successfully established.
  <a id="libpq-multiple-hosts"></a>

#### Specifying Multiple Hosts


 It is possible to specify multiple hosts to connect to, so that they are tried in the given order. In the Keyword/Value format, the `host`, `hostaddr`, and `port` options accept comma-separated lists of values. The same number of elements must be given in each option that is specified, such that e.g., the first `hostaddr` corresponds to the first host name, the second `hostaddr` corresponds to the second host name, and so forth. As an exception, if only one `port` is specified, it applies to all the hosts.


 In the connection URI format, you can list multiple `host:port` pairs separated by commas in the `host` component of the URI.


 In either format, a single host name can translate to multiple network addresses. A common example of this is a host that has both an IPv4 and an IPv6 address.


 When multiple hosts are specified, or when a single host name is translated to multiple addresses, all the hosts and addresses will be tried in order, until one succeeds. If none of the hosts can be reached, the connection fails. If a connection is established successfully, but authentication fails, the remaining hosts in the list are not tried.


 If a password file is used, you can have different passwords for different hosts. All the other connection options are the same for every host in the list; it is not possible to e.g., specify different usernames for different hosts.
   <a id="libpq-paramkeywords"></a>

### Parameter Key Words


 The currently recognized parameter key words are:

<a id="libpq-connect-host"></a>

`host`
:   Name of host to connect to. If a host name looks like an absolute path name, it specifies Unix-domain communication rather than TCP/IP communication; the value is the name of the directory in which the socket file is stored. (On Unix, an absolute path name begins with a slash. On Windows, paths starting with drive letters are also recognized.) If the host name starts with `@`, it is taken as a Unix-domain socket in the abstract namespace (currently supported on Linux and Windows). The default behavior when `host` is not specified, or is empty, is to connect to a Unix-domain socket in `/tmp` (or whatever socket directory was specified when PostgreSQL was built). On Windows, the default is to connect to `localhost`.


     A comma-separated list of host names is also accepted, in which case each host name in the list is tried in order; an empty item in the list selects the default behavior as explained above. See [Specifying Multiple Hosts](#libpq-multiple-hosts) for details.
<a id="libpq-connect-hostaddr"></a>

`hostaddr`
:   Numeric IP address of host to connect to. This should be in the standard IPv4 address format, e.g., `172.28.40.9`. If your machine supports IPv6, you can also use those addresses. TCP/IP communication is always used when a nonempty string is specified for this parameter. If this parameter is not specified, the value of `host` will be looked up to find the corresponding IP address — or, if `host` specifies an IP address, that value will be used directly.


     Using `hostaddr` allows the application to avoid a host name look-up, which might be important in applications with time constraints. However, a host name is required for GSSAPI or SSPI authentication methods, as well as for `verify-full` SSL certificate verification. The following rules are used:

    -  If `host` is specified without `hostaddr`, a host name lookup occurs. (When using `PQconnectPoll`, the lookup occurs when `PQconnectPoll` first considers this host name, and it may cause `PQconnectPoll` to block for a significant amount of time.)
    -  If `hostaddr` is specified without `host`, the value for `hostaddr` gives the server network address. The connection attempt will fail if the authentication method requires a host name.
    -  If both `host` and `hostaddr` are specified, the value for `hostaddr` gives the server network address. The value for `host` is ignored unless the authentication method requires it, in which case it will be used as the host name.
     Note that authentication is likely to fail if `host` is not the name of the server at network address `hostaddr`. Also, when both `host` and `hostaddr` are specified, `host` is used to identify the connection in a password file (see [The Password File](the-password-file.md#libpq-pgpass)).


     A comma-separated list of `hostaddr` values is also accepted, in which case each host in the list is tried in order. An empty item in the list causes the corresponding host name to be used, or the default host name if that is empty as well. See [Specifying Multiple Hosts](#libpq-multiple-hosts) for details.


     Without either a host name or host address, libpq will connect using a local Unix-domain socket; or on Windows, it will attempt to connect to `localhost`.
<a id="libpq-connect-port"></a>

`port`
:   Port number to connect to at the server host, or socket file name extension for Unix-domain connections. If multiple hosts were given in the `host` or `hostaddr` parameters, this parameter may specify a comma-separated list of ports of the same length as the host list, or it may specify a single port number to be used for all hosts. An empty string, or an empty item in a comma-separated list, specifies the default port number established when PostgreSQL was built.
<a id="libpq-connect-dbname"></a>

`dbname`
:   The database name. Defaults to be the same as the user name. In certain contexts, the value is checked for extended formats; see [Connection Strings](#libpq-connstring) for more details on those.
<a id="libpq-connect-user"></a>

`user`
:   PostgreSQL user name to connect as. Defaults to be the same as the operating system name of the user running the application.
<a id="libpq-connect-password"></a>

`password`
:   Password to be used if the server demands password authentication.
<a id="libpq-connect-passfile"></a>

`passfile`
:   Specifies the name of the file used to store passwords (see [The Password File](the-password-file.md#libpq-pgpass)). Defaults to `~/.pgpass`, or `%APPDATA%\postgresql\pgpass.conf` on Microsoft Windows. (No error is reported if this file does not exist.)
<a id="libpq-connect-require-auth"></a>

`require_auth`
:   Specifies the authentication method that the client requires from the server. If the server does not use the required method to authenticate the client, or if the authentication handshake is not fully completed by the server, the connection will fail. A comma-separated list of methods may also be provided, of which the server must use exactly one in order for the connection to succeed. By default, any authentication method is accepted, and the server is free to skip authentication altogether.


     Methods may be negated with the addition of a `!` prefix, in which case the server must *not* attempt the listed method; any other method is accepted, and the server is free not to authenticate the client at all. If a comma-separated list is provided, the server may not attempt *any* of the listed negated methods. Negated and non-negated forms may not be combined in the same setting.


     As a final special case, the `none` method requires the server not to use an authentication challenge. (It may also be negated, to require some form of authentication.)


     The following methods may be specified:

    `password`
    :   The server must request plaintext password authentication.

    `md5`
    :   The server must request MD5 hashed password authentication.


        !!! warning

            Support for MD5-encrypted passwords is deprecated and will be removed in a future release of PostgreSQL. Refer to [Password Authentication](../../server-administration/client-authentication/password-authentication.md#auth-password) for details about migrating to another password type.

    `gss`
    :   The server must either request a Kerberos handshake via GSSAPI or establish a GSS-encrypted channel (see also [gssencmode](#libpq-connect-gssencmode)).

    `sspi`
    :   The server must request Windows SSPI authentication.

    `scram-sha-256`
    :   The server must successfully complete a SCRAM-SHA-256 authentication exchange with the client.

    `oauth`
    :   The server must request an OAuth bearer token from the client.

    `none`
    :   The server must not prompt the client for an authentication exchange. (This does not prohibit client certificate authentication via TLS, nor GSS authentication via its encrypted transport.)
<a id="libpq-connect-channel-binding"></a>

`channel_binding`
:   This option controls the client's use of channel binding. A setting of `require` means that the connection must employ channel binding, `prefer` means that the client will choose channel binding if available, and `disable` prevents the use of channel binding. The default is `prefer` if PostgreSQL is compiled with SSL support; otherwise the default is `disable`.


     Channel binding is a method for the server to authenticate itself to the client. It is only supported over SSL connections with PostgreSQL 11 or later servers using the `SCRAM` authentication method.
<a id="libpq-connect-connect-timeout"></a>

`connect_timeout`
:   Maximum time to wait while connecting, in seconds (write as a decimal integer, e.g., `10`). Zero, negative, or not specified means wait indefinitely. This timeout applies separately to each host name or IP address. For example, if you specify two hosts and `connect_timeout` is 5, each host will time out if no connection is made within 5 seconds, so the total time spent waiting for a connection might be up to 10 seconds.
<a id="libpq-connect-client-encoding"></a>

`client_encoding`
:   This sets the `client_encoding` configuration parameter for this connection. In addition to the values accepted by the corresponding server option, you can use `auto` to determine the right encoding from the current locale in the client (`LC_CTYPE` environment variable on Unix systems).
<a id="libpq-connect-options"></a>

`options`
:   Specifies command-line options to send to the server at connection start. For example, setting this to `-c geqo=off` or `--geqo=off` sets the session's value of the `geqo` parameter to `off`. Spaces within this string are considered to separate command-line arguments, unless escaped with a backslash (`\`); write `\\` to represent a literal backslash. For a detailed discussion of the available options, consult [Server Configuration](../../server-administration/server-configuration/index.md#runtime-config).
<a id="libpq-connect-application-name"></a>

`application_name`
:   Specifies a value for the [application_name](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-application-name) configuration parameter.
<a id="libpq-connect-fallback-application-name"></a>

`fallback_application_name`
:   Specifies a fallback value for the [application_name](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-application-name) configuration parameter. This value will be used if no value has been given for `application_name` via a connection parameter or the `PGAPPNAME` environment variable. Specifying a fallback name is useful in generic utility programs that wish to set a default application name but allow it to be overridden by the user.
<a id="libpq-keepalives"></a>

`keepalives`
:   Controls whether client-side TCP keepalives are used. The default value is 1, meaning on, but you can change this to 0, meaning off, if keepalives are not wanted. This parameter is ignored for connections made via a Unix-domain socket.
<a id="libpq-keepalives-idle"></a>

`keepalives_idle`
:   Controls the number of seconds of inactivity after which TCP should send a keepalive message to the server. A value of zero uses the system default. This parameter is ignored for connections made via a Unix-domain socket, or if keepalives are disabled. It is only supported on systems where `TCP_KEEPIDLE` or an equivalent socket option is available, and on Windows; on other systems, it has no effect.
<a id="libpq-keepalives-interval"></a>

`keepalives_interval`
:   Controls the number of seconds after which a TCP keepalive message that is not acknowledged by the server should be retransmitted. A value of zero uses the system default. This parameter is ignored for connections made via a Unix-domain socket, or if keepalives are disabled. It is only supported on systems where `TCP_KEEPINTVL` or an equivalent socket option is available, and on Windows; on other systems, it has no effect.
<a id="libpq-keepalives-count"></a>

`keepalives_count`
:   Controls the number of TCP keepalives that can be lost before the client's connection to the server is considered dead. A value of zero uses the system default. This parameter is ignored for connections made via a Unix-domain socket, or if keepalives are disabled. It is only supported on systems where `TCP_KEEPCNT` or an equivalent socket option is available; on other systems, it has no effect.
<a id="libpq-tcp-user-timeout"></a>

`tcp_user_timeout`
:   Controls the number of milliseconds that transmitted data may remain unacknowledged before a connection is forcibly closed. A value of zero uses the system default. This parameter is ignored for connections made via a Unix-domain socket. It is only supported on systems where `TCP_USER_TIMEOUT` is available; on other systems, it has no effect.
<a id="libpq-connect-replication"></a>

`replication`
:   This option determines whether the connection should use the replication protocol instead of the normal protocol. This is what PostgreSQL replication connections as well as tools such as pg_basebackup use internally, but it can also be used by third-party applications. For a description of the replication protocol, consult [Streaming Replication Protocol](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication).


     The following values, which are case-insensitive, are supported:

    `true`, `on`, `yes`, `1`
    :   The connection goes into physical replication mode.

    `database`
    :   The connection goes into logical replication mode, connecting to the database specified in the `dbname` parameter.

    `false`, `off`, `no`, `0`
    :   The connection is a regular one, which is the default behavior.


     In physical or logical replication mode, only the simple query protocol can be used.
<a id="libpq-connect-gssencmode"></a>

`gssencmode`
:   This option determines whether or with what priority a secure GSS TCP/IP connection will be negotiated with the server. There are three modes:

    `disable`
    :   only try a non-GSSAPI-encrypted connection

    `prefer` (default)
    :   if there are GSSAPI credentials present (i.e., in a credentials cache), first try a GSSAPI-encrypted connection; if that fails or there are no credentials, try a non-GSSAPI-encrypted connection. This is the default when PostgreSQL has been compiled with GSSAPI support.

    `require`
    :   only try a GSSAPI-encrypted connection


     `gssencmode` is ignored for Unix domain socket communication. If PostgreSQL is compiled without GSSAPI support, using the `require` option will cause an error, while `prefer` will be accepted but libpq will not actually attempt a GSSAPI-encrypted connection.
<a id="libpq-connect-sslmode"></a>

`sslmode`
:   This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the server. There are six modes:

    `disable`
    :   only try a non-SSL connection

    `allow`
    :   first try a non-SSL connection; if that fails, try an SSL connection

    `prefer` (default)
    :   first try an SSL connection; if that fails, try a non-SSL connection

    `require`
    :   only try an SSL connection. If a root CA file is present, verify the certificate in the same way as if `verify-ca` was specified

    `verify-ca`
    :   only try an SSL connection, and verify that the server certificate is issued by a trusted certificate authority (CA)

    `verify-full`
    :   only try an SSL connection, verify that the server certificate is issued by a trusted CA and that the requested server host name matches that in the certificate
     See [SSL Support](ssl-support.md#libpq-ssl) for a detailed description of how these options work.


     `sslmode` is ignored for Unix domain socket communication. If PostgreSQL is compiled without SSL support, using options `require`, `verify-ca`, or `verify-full` will cause an error, while options `allow` and `prefer` will be accepted but libpq will not actually attempt an SSL connection.


     Note that if GSSAPI encryption is possible, that will be used in preference to SSL encryption, regardless of the value of `sslmode`. To force use of SSL encryption in an environment that has working GSSAPI infrastructure (such as a Kerberos server), also set `gssencmode` to `disable`.
<a id="libpq-connect-requiressl"></a>

`requiressl`
:   This option is deprecated in favor of the `sslmode` setting.


     If set to 1, an SSL connection to the server is required (this is equivalent to `sslmode` `require`). libpq will then refuse to connect if the server does not accept an SSL connection. If set to 0 (default), libpq will negotiate the connection type with the server (equivalent to `sslmode` `prefer`). This option is only available if PostgreSQL is compiled with SSL support.
<a id="libpq-connect-sslnegotiation"></a>

`sslnegotiation`
:   This option controls how SSL encryption is negotiated with the server, if SSL is used. In the default `postgres` mode, the client first asks the server if SSL is supported. In `direct` mode, the client starts the standard SSL handshake directly after establishing the TCP/IP connection. Traditional PostgreSQL protocol negotiation is the most flexible with different server configurations. If the server is known to support direct SSL connections then the latter requires one fewer round trip reducing connection latency and also allows the use of protocol agnostic SSL network tools. The direct SSL option was introduced in PostgreSQL version 17.


    `postgres`
    :   perform PostgreSQL protocol negotiation. This is the default if the option is not provided.

    `direct`
    :   start SSL handshake directly after establishing the TCP/IP connection. This is only allowed with `sslmode=require` or higher, because the weaker settings could lead to unintended fallback to plaintext authentication when the server does not support direct SSL handshake.
<a id="libpq-connect-sslcompression"></a>

`sslcompression`
:   If set to 1, data sent over SSL connections will be compressed. If set to 0, compression will be disabled. The default is 0. This parameter is ignored if a connection without SSL is made.


     SSL compression is nowadays considered insecure and its use is no longer recommended. OpenSSL 1.1.0 disabled compression by default, and many operating system distributions disabled it in prior versions as well, so setting this parameter to on will not have any effect if the server does not accept compression. PostgreSQL 14 disabled compression completely in the backend.


     If security is not a primary concern, compression can improve throughput if the network is the bottleneck. Disabling compression can improve response time and throughput if CPU performance is the limiting factor.
<a id="libpq-connect-sslcert"></a>

`sslcert`
:   This parameter specifies the file name of the client SSL certificate, replacing the default `~/.postgresql/postgresql.crt`. This parameter is ignored if an SSL connection is not made.
<a id="libpq-connect-sslkey"></a>

`sslkey`
:   This parameter specifies the location for the secret key used for the client certificate. It can either specify a file name that will be used instead of the default `~/.postgresql/postgresql.key`, or it can specify a key obtained from an external “engine” (engines are OpenSSL loadable modules). An external engine specification should consist of a colon-separated engine name and an engine-specific key identifier. This parameter is ignored if an SSL connection is not made.
<a id="libpq-connect-sslkeylogfile"></a>

`sslkeylogfile`
:   This parameter specifies the location where libpq will log keys used in this SSL context. This is useful for debugging PostgreSQL protocol interactions or client connections using network inspection tools like Wireshark. This parameter is ignored if an SSL connection is not made, or if LibreSSL is used (LibreSSL does not support key logging). Keys are logged using the NSS format.

    !!! warning

        Key logging will expose potentially sensitive information in the keylog file. Keylog files should be handled with the same care as [sslkey](#libpq-connect-sslkey) files.
<a id="libpq-connect-sslpassword"></a>

`sslpassword`
:   This parameter specifies the password for the secret key specified in `sslkey`, allowing client certificate private keys to be stored in encrypted form on disk even when interactive passphrase input is not practical.


     Specifying this parameter with any non-empty value suppresses the `Enter PEM pass phrase:` prompt that OpenSSL will emit by default when an encrypted client certificate key is provided to libpq.


     If the key is not encrypted this parameter is ignored. The parameter has no effect on keys specified by OpenSSL engines unless the engine uses the OpenSSL password callback mechanism for prompts.


     There is no environment variable equivalent to this option, and no facility for looking it up in `.pgpass`. It can be used in a service file connection definition. Users with more sophisticated uses should consider using OpenSSL engines and tools like PKCS#11 or USB crypto offload devices.
<a id="libpq-connect-sslcertmode"></a>

`sslcertmode`
:   This option determines whether a client certificate may be sent to the server, and whether the server is required to request one. There are three modes:

    `disable`
    :   A client certificate is never sent, even if one is available (default location or provided via [sslcert](#libpq-connect-sslcert)).

    `allow` (default)
    :   A certificate may be sent, if the server requests one and the client has one to send.

    `require`
    :   The server *must* request a certificate. The connection will fail if the client does not send a certificate and the server successfully authenticates the client anyway.


    !!! note

        `sslcertmode=require` doesn't add any additional security, since there is no guarantee that the server is validating the certificate correctly; PostgreSQL servers generally request TLS certificates from clients whether they validate them or not. The option may be useful when troubleshooting more complicated TLS setups.
<a id="libpq-connect-sslrootcert"></a>

`sslrootcert`
:   This parameter specifies the name of a file containing SSL certificate authority (CA) certificate(s). If the file exists, the server's certificate will be verified to be signed by one of these authorities. The default is `~/.postgresql/root.crt`.


     The special value `system` may be specified instead, in which case the trusted CA roots from the SSL implementation will be loaded. The exact locations of these root certificates differ by SSL implementation and platform. For OpenSSL in particular, the locations may be further modified by the `SSL_CERT_DIR` and `SSL_CERT_FILE` environment variables.


    !!! note

        When using `sslrootcert=system`, the default `sslmode` is changed to `verify-full`, and any weaker setting will result in an error. In most cases it is trivial for anyone to obtain a certificate trusted by the system for a hostname they control, rendering `verify-ca` and all weaker modes useless.


         The magic `system` value will take precedence over a local certificate file with the same name. If for some reason you find yourself in this situation, use an alternative path like `sslrootcert=./system` instead.
<a id="libpq-connect-sslcrl"></a>

`sslcrl`
:   This parameter specifies the file name of the SSL server certificate revocation list (CRL). Certificates listed in this file, if it exists, will be rejected while attempting to authenticate the server's certificate. If neither [sslcrl](#libpq-connect-sslcrl) nor [sslcrldir](#libpq-connect-sslcrldir) is set, this setting is taken as `~/.postgresql/root.crl`.
<a id="libpq-connect-sslcrldir"></a>

`sslcrldir`
:   This parameter specifies the directory name of the SSL server certificate revocation list (CRL). Certificates listed in the files in this directory, if it exists, will be rejected while attempting to authenticate the server's certificate.


     The directory needs to be prepared with the OpenSSL command `openssl rehash` or `c_rehash`. See its documentation for details.


     Both `sslcrl` and `sslcrldir` can be specified together.
<a id="libpq-connect-sslsni"></a>

`sslsni`
:   If set to 1 (default), libpq sets the TLS extension “Server Name Indication” (SNI) on SSL-enabled connections. By setting this parameter to 0, this is turned off.


     The Server Name Indication can be used by SSL-aware proxies to route connections without having to decrypt the SSL stream. (Note that unless the proxy is aware of the PostgreSQL protocol handshake this would require setting `sslnegotiation` to `direct`.) However, SNI makes the destination host name appear in cleartext in the network traffic, so it might be undesirable in some cases.
<a id="libpq-connect-requirepeer"></a>

`requirepeer`
:   This parameter specifies the operating-system user name of the server, for example `requirepeer=postgres`. When making a Unix-domain socket connection, if this parameter is set, the client checks at the beginning of the connection that the server process is running under the specified user name; if it is not, the connection is aborted with an error. This parameter can be used to provide server authentication similar to that available with SSL certificates on TCP/IP connections. (Note that if the Unix-domain socket is in `/tmp` or another publicly writable location, any user could start a server listening there. Use this parameter to ensure that you are connected to a server run by a trusted user.) This option is only supported on platforms for which the `peer` authentication method is implemented; see [Peer Authentication](../../server-administration/client-authentication/peer-authentication.md#auth-peer).
<a id="libpq-connect-ssl-min-protocol-version"></a>

`ssl_min_protocol_version`
:   This parameter specifies the minimum SSL/TLS protocol version to allow for the connection. Valid values are `TLSv1`, `TLSv1.1`, `TLSv1.2` and `TLSv1.3`. The supported protocols depend on the version of OpenSSL used, older versions not supporting the most modern protocol versions. If not specified, the default is `TLSv1.2`, which satisfies industry best practices as of this writing.
<a id="libpq-connect-ssl-max-protocol-version"></a>

`ssl_max_protocol_version`
:   This parameter specifies the maximum SSL/TLS protocol version to allow for the connection. Valid values are `TLSv1`, `TLSv1.1`, `TLSv1.2` and `TLSv1.3`. The supported protocols depend on the version of OpenSSL used, older versions not supporting the most modern protocol versions. If not set, this parameter is ignored and the connection will use the maximum bound defined by the backend, if set. Setting the maximum protocol version is mainly useful for testing or if some component has issues working with a newer protocol.
<a id="libpq-connect-min-protocol-version"></a>

`min_protocol_version`
:   Specifies the minimum protocol version to allow for the connection. The default is to allow any version of the PostgreSQL protocol supported by libpq, which currently means `3.0`. If the server does not support at least this protocol version the connection will be closed.


     The current supported values are `3.0`, `3.2`, and `latest`. The `latest` value is equivalent to the latest protocol version supported by the libpq version being used, which is currently `3.2`.
<a id="libpq-connect-max-protocol-version"></a>

`max_protocol_version`
:   Specifies the protocol version to request from the server. The default is to use version `3.0` of the PostgreSQL protocol, unless the connection string specifies a feature that relies on a higher protocol version, in which case the latest version supported by libpq is used. If the server does not support the protocol version requested by the client, the connection is automatically downgraded to a lower minor protocol version that the server supports. After the connection attempt has completed you can use [PQfullProtocolVersion](connection-status-functions.md#libpq-PQfullProtocolVersion) to find out which exact protocol version was negotiated.


     The current supported values are `3.0`, `3.2`, and `latest`. The `latest` value is equivalent to the latest protocol version supported by the libpq version being used, which is currently `3.2`.
<a id="libpq-connect-krbsrvname"></a>

`krbsrvname`
:   Kerberos service name to use when authenticating with GSSAPI. This must match the service name specified in the server configuration for Kerberos authentication to succeed. (See also [GSSAPI Authentication](../../server-administration/client-authentication/gssapi-authentication.md#gssapi-auth).) The default value is normally `postgres`, but that can be changed when building PostgreSQL via the `--with-krb-srvnam` option of configure. In most environments, this parameter never needs to be changed. Some Kerberos implementations might require a different service name, such as Microsoft Active Directory which requires the service name to be in upper case (`POSTGRES`).
<a id="libpq-connect-gsslib"></a>

`gsslib`
:   GSS library to use for GSSAPI authentication. Currently this is disregarded except on Windows builds that include both GSSAPI and SSPI support. In that case, set this to `gssapi` to cause libpq to use the GSSAPI library for authentication instead of the default SSPI.
<a id="libpq-connect-gssdelegation"></a>

`gssdelegation`
:   Forward (delegate) GSS credentials to the server. The default is `0` which means credentials will not be forwarded to the server. Set this to `1` to have credentials forwarded when possible.
<a id="libpq-connect-scram-client-key"></a>

`scram_client_key`
:   The base64-encoded SCRAM client key. This can be used by foreign-data wrappers or similar middleware to enable pass-through SCRAM authentication. See [Connection Management Options](../../appendixes/additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-options-connection-management) for one such implementation. It is not meant to be specified directly by users or client applications.
<a id="libpq-connect-scram-server-key"></a>

`scram_server_key`
:   The base64-encoded SCRAM server key. This can be used by foreign-data wrappers or similar middleware to enable pass-through SCRAM authentication. See [Connection Management Options](../../appendixes/additional-supplied-modules-and-extensions/postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw-options-connection-management) for one such implementation. It is not meant to be specified directly by users or client applications.
<a id="libpq-connect-service"></a>

`service`
:   Service name to use for additional parameters. It specifies a service name in `pg_service.conf` that holds additional connection parameters. This allows applications to specify only a service name so connection parameters can be centrally maintained. See [The Connection Service File](the-connection-service-file.md#libpq-pgservice).
<a id="libpq-connect-target-session-attrs"></a>

`target_session_attrs`
:   This option determines whether the session must have certain properties to be acceptable. It's typically used in combination with multiple host names to select the first acceptable alternative among several hosts. There are six modes:

    `any` (default)
    :   any successful connection is acceptable

    `read-write`
    :   session must accept read-write transactions by default (that is, the server must not be in hot standby mode and the `default_transaction_read_only` parameter must be `off`)

    `read-only`
    :   session must not accept read-write transactions by default (the converse)

    `primary`
    :   server must not be in hot standby mode

    `standby`
    :   server must be in hot standby mode

    `prefer-standby`
    :   first try to find a standby server, but if none of the listed hosts is a standby server, try again in `any` mode
<a id="libpq-connect-load-balance-hosts"></a>

`load_balance_hosts`
:   Controls the order in which the client tries to connect to the available hosts and addresses. Once a connection attempt is successful no other hosts and addresses will be tried. This parameter is typically used in combination with multiple host names or a DNS record that returns multiple IPs. This parameter can be used in combination with [target_session_attrs](#libpq-connect-target-session-attrs) to, for example, load balance over standby servers only. Once successfully connected, subsequent queries on the returned connection will all be sent to the same server. There are currently two modes:

    `disable` (default)
    :   No load balancing across hosts is performed. Hosts are tried in the order in which they are provided and addresses are tried in the order they are received from DNS or a hosts file.

    `random`
    :   Hosts and addresses are tried in random order. This value is mostly useful when opening multiple connections at the same time, possibly from different machines. This way connections can be load balanced across multiple PostgreSQL servers.


         While random load balancing, due to its random nature, will almost never result in a completely uniform distribution, it statistically gets quite close. One important aspect here is that this algorithm uses two levels of random choices: First the hosts will be resolved in random order. Then secondly, before resolving the next host, all resolved addresses for the current host will be tried in random order. This behaviour can skew the amount of connections each node gets greatly in certain cases, for instance when some hosts resolve to more addresses than others. But such a skew can also be used on purpose, e.g. to increase the number of connections a larger server gets by providing its hostname multiple times in the host string.


         When using this value it's recommended to also configure a reasonable value for [connect_timeout](#libpq-connect-connect-timeout). Because then, if one of the nodes that are used for load balancing is not responding, a new node will be tried.
<a id="libpq-connect-oauth-issuer"></a>

`oauth_issuer`
:   The HTTPS URL of a trusted issuer to contact if the server requests an OAuth token for the connection. This parameter is required for all OAuth connections; it should exactly match the `issuer` setting in [the server's HBA configuration](../../server-administration/client-authentication/oauth-authorization-authentication.md#auth-oauth).


     As part of the standard authentication handshake, libpq will ask the server for a *discovery document:* a URL providing a set of OAuth configuration parameters. The server must provide a URL that is directly constructed from the components of the `oauth_issuer`, and this value must exactly match the issuer identifier that is declared in the discovery document itself, or the connection will fail. This is required to prevent a class of ["mix-up attacks"](https://mailarchive.ietf.org/arch/msg/oauth/JIVxFBGsJBVtm7ljwJhPUm3Fr-w/) on OAuth clients.


     You may also explicitly set `oauth_issuer` to the `/.well-known/` URI used for OAuth discovery. In this case, if the server asks for a different URL, the connection will fail, but a [custom OAuth flow](oauth-support.md#libpq-oauth-authdata-hooks) may be able to speed up the standard handshake by using previously cached tokens. (In this case, it is recommended that [oauth_scope](#libpq-connect-oauth-scope) be set as well, since the client will not have a chance to ask the server for a correct scope setting, and the default scopes for a token may not be sufficient to connect.) libpq currently supports the following well-known endpoints:

    - `/.well-known/openid-configuration`
    - `/.well-known/oauth-authorization-server`


    !!! warning

        Issuers are highly privileged during the OAuth connection handshake. As a rule of thumb, if you would not trust the operator of a URL to handle access to your servers, or to impersonate you directly, that URL should not be trusted as an `oauth_issuer`.
<a id="libpq-connect-oauth-client-id"></a>

`oauth_client_id`
:   An OAuth 2.0 client identifier, as issued by the authorization server. If the PostgreSQL server [requests an OAuth token](../../server-administration/client-authentication/oauth-authorization-authentication.md#auth-oauth) for the connection (and if no [custom OAuth hook](oauth-support.md#libpq-oauth-authdata-hooks) is installed to provide one), then this parameter must be set; otherwise, the connection will fail.
<a id="libpq-connect-oauth-client-secret"></a>

`oauth_client_secret`
:   The client password, if any, to use when contacting the OAuth authorization server. Whether this parameter is required or not is determined by the OAuth provider; "public" clients generally do not use a secret, whereas "confidential" clients generally do.
<a id="libpq-connect-oauth-scope"></a>

`oauth_scope`
:   The scope of the access request sent to the authorization server, specified as a (possibly empty) space-separated list of OAuth scope identifiers. This parameter is optional and intended for advanced usage.


     Usually the client will obtain appropriate scope settings from the PostgreSQL server. If this parameter is used, the server's requested scope list will be ignored. This can prevent a less-trusted server from requesting inappropriate access scopes from the end user. However, if the client's scope setting does not contain the server's required scopes, the server is likely to reject the issued token, and the connection will fail.


     The meaning of an empty scope list is provider-dependent. An OAuth authorization server may choose to issue a token with "default scope", whatever that happens to be, or it may reject the token request entirely.
