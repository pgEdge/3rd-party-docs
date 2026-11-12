<a id="ecpg-library"></a>

## Library Functions


 The `libecpg` library primarily contains “hidden” functions that are used to implement the functionality expressed by the embedded SQL commands. But there are some functions that can usefully be called directly. Note that this makes your code unportable.


-  <code>ECPGdebug(int </code><em>on</em><code>, FILE
     *</code><em>stream</em><code>)</code> turns on debug logging if called with the first argument non-zero. Debug logging is done on *stream*. The log contains all SQL statements with all the input variables inserted, and the results from the PostgreSQL server. This can be very useful when searching for errors in your SQL statements.

!!! note

    On Windows, if the ecpg libraries and an application are compiled with different flags, this function call will crash the application because the internal representation of the `FILE` pointers differ. Specifically, multithreaded/single-threaded, release/debug, and static/dynamic flags should be the same for the library and all applications using that library.
-  <code>ECPGget_PGconn(const char *</code><em>connection_name</em><code>)
       </code> returns the library database connection handle identified by the given name. If *connection_name* is set to `NULL`, the current connection handle is returned. If no connection handle can be identified, the function returns `NULL`. The returned connection handle can be used to call any other functions from libpq, if necessary.

!!! note

    It is a bad idea to manipulate database connection handles made from ecpg directly with libpq routines.
-  <code>ECPGtransactionStatus(const char *</code><em>connection_name</em><code>)</code> returns the current transaction status of the given connection identified by *connection_name*. See [Connection Status Functions](../libpq-c-library/connection-status-functions.md#libpq-status) and libpq's [PQtransactionStatus](../libpq-c-library/connection-status-functions.md#libpq-PQtransactionStatus) for details about the returned status codes.
-  <code>ECPGstatus(int </code><em>lineno</em><code>,
     const char* </code><em>connection_name</em><code>)</code> returns true if you are connected to a database and false if not. *connection_name* can be `NULL` if a single connection is being used.
