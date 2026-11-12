<a id="libpq-async"></a>

## Asynchronous Command Processing


 The [PQexec](command-execution-functions.md#libpq-PQexec) function is adequate for submitting commands in normal, synchronous applications. It has a few deficiencies, however, that can be of importance to some users:

-  [PQexec](command-execution-functions.md#libpq-PQexec) waits for the command to be completed. The application might have other work to do (such as maintaining a user interface), in which case it won't want to block waiting for the response.
-  Since the execution of the client application is suspended while it waits for the result, it is hard for the application to decide that it would like to try to cancel the ongoing command. (It can be done from a signal handler, but not otherwise.)
-  [PQexec](command-execution-functions.md#libpq-PQexec) can return only one `PGresult` structure. If the submitted command string contains multiple SQL commands, all but the last `PGresult` are discarded by [PQexec](command-execution-functions.md#libpq-PQexec).
-  [PQexec](command-execution-functions.md#libpq-PQexec) always collects the command's entire result, buffering it in a single `PGresult`. While this simplifies error-handling logic for the application, it can be impractical for results containing many rows.


 Applications that do not like these limitations can instead use the underlying functions that [PQexec](command-execution-functions.md#libpq-PQexec) is built from: [PQsendQuery](#libpq-PQsendQuery) and [PQgetResult](#libpq-PQgetResult). There are also [PQsendQueryParams](#libpq-PQsendQueryParams), [PQsendPrepare](#libpq-PQsendPrepare), [PQsendQueryPrepared](#libpq-PQsendQueryPrepared), [PQsendDescribePrepared](#libpq-PQsendDescribePrepared), [PQsendDescribePortal](#libpq-PQsendDescribePortal), [PQsendClosePrepared](#libpq-PQsendClosePrepared), and [PQsendClosePortal](#libpq-PQsendClosePortal), which can be used with [PQgetResult](#libpq-PQgetResult) to duplicate the functionality of [PQexecParams](command-execution-functions.md#libpq-PQexecParams), [PQprepare](command-execution-functions.md#libpq-PQprepare), [PQexecPrepared](command-execution-functions.md#libpq-PQexecPrepared), [PQdescribePrepared](command-execution-functions.md#libpq-PQdescribePrepared), [PQdescribePortal](command-execution-functions.md#libpq-PQdescribePortal), [PQclosePrepared](command-execution-functions.md#libpq-PQclosePrepared), and [PQclosePortal](command-execution-functions.md#libpq-PQclosePortal) respectively.

<a id="libpq-PQsendQuery"></a>

`PQsendQuery`
:   Submits a command to the server without waiting for the result(s). 1 is returned if the command was successfully dispatched and 0 if not (in which case, use [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage) to get more information about the failure).

    ```

    int PQsendQuery(PGconn *conn, const char *command);
    ```
     After successfully calling [PQsendQuery](#libpq-PQsendQuery), call [PQgetResult](#libpq-PQgetResult) one or more times to obtain the results. [PQsendQuery](#libpq-PQsendQuery) cannot be called again (on the same connection) until [PQgetResult](#libpq-PQgetResult) has returned a null pointer, indicating that the command is done.


     In pipeline mode, this function is disallowed.
<a id="libpq-PQsendQueryParams"></a>

`PQsendQueryParams`
:   Submits a command and separate parameters to the server without waiting for the result(s).

    ```

    int PQsendQueryParams(PGconn *conn,
                          const char *command,
                          int nParams,
                          const Oid *paramTypes,
                          const char * const *paramValues,
                          const int *paramLengths,
                          const int *paramFormats,
                          int resultFormat);
    ```
     This is equivalent to [PQsendQuery](#libpq-PQsendQuery) except that query parameters can be specified separately from the query string. The function's parameters are handled identically to [PQexecParams](command-execution-functions.md#libpq-PQexecParams). Like [PQexecParams](command-execution-functions.md#libpq-PQexecParams), it allows only one command in the query string.
<a id="libpq-PQsendPrepare"></a>

`PQsendPrepare`
:   Sends a request to create a prepared statement with the given parameters, without waiting for completion.

    ```

    int PQsendPrepare(PGconn *conn,
                      const char *stmtName,
                      const char *query,
                      int nParams,
                      const Oid *paramTypes);
    ```
     This is an asynchronous version of [PQprepare](command-execution-functions.md#libpq-PQprepare): it returns 1 if it was able to dispatch the request, and 0 if not. After a successful call, call [PQgetResult](#libpq-PQgetResult) to determine whether the server successfully created the prepared statement. The function's parameters are handled identically to [PQprepare](command-execution-functions.md#libpq-PQprepare).
<a id="libpq-PQsendQueryPrepared"></a>

`PQsendQueryPrepared`
:   Sends a request to execute a prepared statement with given parameters, without waiting for the result(s).

    ```

    int PQsendQueryPrepared(PGconn *conn,
                            const char *stmtName,
                            int nParams,
                            const char * const *paramValues,
                            const int *paramLengths,
                            const int *paramFormats,
                            int resultFormat);
    ```
     This is similar to [PQsendQueryParams](#libpq-PQsendQueryParams), but the command to be executed is specified by naming a previously-prepared statement, instead of giving a query string. The function's parameters are handled identically to [PQexecPrepared](command-execution-functions.md#libpq-PQexecPrepared).
<a id="libpq-PQsendDescribePrepared"></a>

`PQsendDescribePrepared`
:   Submits a request to obtain information about the specified prepared statement, without waiting for completion.

    ```

    int PQsendDescribePrepared(PGconn *conn, const char *stmtName);
    ```
     This is an asynchronous version of [PQdescribePrepared](command-execution-functions.md#libpq-PQdescribePrepared): it returns 1 if it was able to dispatch the request, and 0 if not. After a successful call, call [PQgetResult](#libpq-PQgetResult) to obtain the results. The function's parameters are handled identically to [PQdescribePrepared](command-execution-functions.md#libpq-PQdescribePrepared).
<a id="libpq-PQsendDescribePortal"></a>

`PQsendDescribePortal`
:   Submits a request to obtain information about the specified portal, without waiting for completion.

    ```

    int PQsendDescribePortal(PGconn *conn, const char *portalName);
    ```
     This is an asynchronous version of [PQdescribePortal](command-execution-functions.md#libpq-PQdescribePortal): it returns 1 if it was able to dispatch the request, and 0 if not. After a successful call, call [PQgetResult](#libpq-PQgetResult) to obtain the results. The function's parameters are handled identically to [PQdescribePortal](command-execution-functions.md#libpq-PQdescribePortal).
<a id="libpq-PQsendClosePrepared"></a>

`PQsendClosePrepared`
:   Submits a request to close the specified prepared statement, without waiting for completion.

    ```

    int PQsendClosePrepared(PGconn *conn, const char *stmtName);
    ```
     This is an asynchronous version of [PQclosePrepared](command-execution-functions.md#libpq-PQclosePrepared): it returns 1 if it was able to dispatch the request, and 0 if not. After a successful call, call [PQgetResult](#libpq-PQgetResult) to obtain the results. The function's parameters are handled identically to [PQclosePrepared](command-execution-functions.md#libpq-PQclosePrepared).
<a id="libpq-PQsendClosePortal"></a>

`PQsendClosePortal`
:   Submits a request to close specified portal, without waiting for completion.

    ```

    int PQsendClosePortal(PGconn *conn, const char *portalName);
    ```
     This is an asynchronous version of [PQclosePortal](command-execution-functions.md#libpq-PQclosePortal): it returns 1 if it was able to dispatch the request, and 0 if not. After a successful call, call [PQgetResult](#libpq-PQgetResult) to obtain the results. The function's parameters are handled identically to [PQclosePortal](command-execution-functions.md#libpq-PQclosePortal).
<a id="libpq-PQgetResult"></a>

`PQgetResult`
:   Waits for the next result from a prior [PQsendQuery](#libpq-PQsendQuery), [PQsendQueryParams](#libpq-PQsendQueryParams), [PQsendPrepare](#libpq-PQsendPrepare), [PQsendQueryPrepared](#libpq-PQsendQueryPrepared), [PQsendDescribePrepared](#libpq-PQsendDescribePrepared), [PQsendDescribePortal](#libpq-PQsendDescribePortal), [PQsendClosePrepared](#libpq-PQsendClosePrepared), [PQsendClosePortal](#libpq-PQsendClosePortal), [PQsendPipelineSync](pipeline-mode.md#libpq-PQsendPipelineSync), or [PQpipelineSync](pipeline-mode.md#libpq-PQpipelineSync) call, and returns it. A null pointer is returned when the command is complete and there will be no more results.

    ```

    PGresult *PQgetResult(PGconn *conn);
    ```


     [PQgetResult](#libpq-PQgetResult) must be called repeatedly until it returns a null pointer, indicating that the command is done. (If called when no command is active, [PQgetResult](#libpq-PQgetResult) will just return a null pointer at once.) Each non-null result from [PQgetResult](#libpq-PQgetResult) should be processed using the same `PGresult` accessor functions previously described. Don't forget to free each result object with [PQclear](command-execution-functions.md#libpq-PQclear) when done with it. Note that [PQgetResult](#libpq-PQgetResult) will block only if a command is active and the necessary response data has not yet been read by [PQconsumeInput](#libpq-PQconsumeInput).


     In pipeline mode, `PQgetResult` will return normally unless an error occurs; for any subsequent query sent after the one that caused the error until (and excluding) the next synchronization point, a special result of type `PGRES_PIPELINE_ABORTED` will be returned, and a null pointer will be returned after it. When the pipeline synchronization point is reached, a result of type `PGRES_PIPELINE_SYNC` will be returned. The result of the next query after the synchronization point follows immediately (that is, no null pointer is returned after the synchronization point).


    !!! note

        Even when [PQresultStatus](command-execution-functions.md#libpq-PQresultStatus) indicates a fatal error, [PQgetResult](#libpq-PQgetResult) should be called until it returns a null pointer, to allow libpq to process the error information completely.


 Using [PQsendQuery](#libpq-PQsendQuery) and [PQgetResult](#libpq-PQgetResult) solves one of [PQexec](command-execution-functions.md#libpq-PQexec)'s problems: If a command string contains multiple SQL commands, the results of those commands can be obtained individually. (This allows a simple form of overlapped processing, by the way: the client can be handling the results of one command while the server is still working on later queries in the same command string.)


 Another frequently-desired feature that can be obtained with [PQsendQuery](#libpq-PQsendQuery) and [PQgetResult](#libpq-PQgetResult) is retrieving large query results a limited number of rows at a time. This is discussed in [Retrieving Query Results in Chunks](retrieving-query-results-in-chunks.md#libpq-single-row-mode).


 By itself, calling [PQgetResult](#libpq-PQgetResult) will still cause the client to block until the server completes the next SQL command. This can be avoided by proper use of two more functions:

<a id="libpq-PQconsumeInput"></a>

`PQconsumeInput`
:   If input is available from the server, consume it.

    ```

    int PQconsumeInput(PGconn *conn);
    ```


     [PQconsumeInput](#libpq-PQconsumeInput) normally returns 1 indicating “no error”, but returns 0 if there was some kind of trouble (in which case [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage) can be consulted). Note that the result does not say whether any input data was actually collected. After calling [PQconsumeInput](#libpq-PQconsumeInput), the application can check [PQisBusy](#libpq-PQisBusy) and/or `PQnotifies` to see if their state has changed.


     [PQconsumeInput](#libpq-PQconsumeInput) can be called even if the application is not prepared to deal with a result or notification just yet. The function will read available data and save it in a buffer, thereby causing a `select()` read-ready indication to go away. The application can thus use [PQconsumeInput](#libpq-PQconsumeInput) to clear the `select()` condition immediately, and then examine the results at leisure.
<a id="libpq-PQisBusy"></a>

`PQisBusy`
:   Returns 1 if a command is busy, that is, [PQgetResult](#libpq-PQgetResult) would block waiting for input. A 0 return indicates that [PQgetResult](#libpq-PQgetResult) can be called with assurance of not blocking.

    ```

    int PQisBusy(PGconn *conn);
    ```


     [PQisBusy](#libpq-PQisBusy) will not itself attempt to read data from the server; therefore [PQconsumeInput](#libpq-PQconsumeInput) must be invoked first, or the busy state will never end.


 A typical application using these functions will have a main loop that uses `select()` or `poll()` to wait for all the conditions that it must respond to. One of the conditions will be input available from the server, which in terms of `select()` means readable data on the file descriptor identified by [PQsocket](connection-status-functions.md#libpq-PQsocket). When the main loop detects input ready, it should call [PQconsumeInput](#libpq-PQconsumeInput) to read the input. It can then call [PQisBusy](#libpq-PQisBusy), followed by [PQgetResult](#libpq-PQgetResult) if [PQisBusy](#libpq-PQisBusy) returns false (0). It can also call `PQnotifies` to detect `NOTIFY` messages (see [Asynchronous Notification](asynchronous-notification.md#libpq-notify)).


 A client that uses [PQsendQuery](#libpq-PQsendQuery)/[PQgetResult](#libpq-PQgetResult) can also attempt to cancel a command that is still being processed by the server; see [Canceling Queries in Progress](canceling-queries-in-progress.md#libpq-cancel). But regardless of the return value of [PQcancelBlocking](canceling-queries-in-progress.md#libpq-PQcancelBlocking), the application must continue with the normal result-reading sequence using [PQgetResult](#libpq-PQgetResult). A successful cancellation will simply cause the command to terminate sooner than it would have otherwise.


 By using the functions described above, it is possible to avoid blocking while waiting for input from the database server. However, it is still possible that the application will block waiting to send output to the server. This is relatively uncommon but can happen if very long SQL commands or data values are sent. (It is much more probable if the application sends data via `COPY IN`, however.) To prevent this possibility and achieve completely nonblocking database operation, the following additional functions can be used.

<a id="libpq-PQsetnonblocking"></a>

`PQsetnonblocking`
:   Sets the nonblocking status of the connection.

    ```

    int PQsetnonblocking(PGconn *conn, int arg);
    ```


     Sets the state of the connection to nonblocking if `arg` is 1, or blocking if `arg` is 0. Returns 0 if OK, -1 if error.


     In the nonblocking state, successful calls to [PQsendQuery](#libpq-PQsendQuery), [PQputline](functions-associated-with-the-copy-command.md#libpq-PQputline), [PQputnbytes](functions-associated-with-the-copy-command.md#libpq-PQputnbytes), [PQputCopyData](functions-associated-with-the-copy-command.md#libpq-PQputCopyData), and [PQendcopy](functions-associated-with-the-copy-command.md#libpq-PQendcopy) will not block; their changes are stored in the local output buffer until they are flushed. Unsuccessful calls will return an error and must be retried.


     Note that [PQexec](command-execution-functions.md#libpq-PQexec) does not honor nonblocking mode; if it is called, it will act in blocking fashion anyway.
<a id="libpq-PQisnonblocking"></a>

`PQisnonblocking`
:   Returns the blocking status of the database connection.

    ```

    int PQisnonblocking(const PGconn *conn);
    ```


     Returns 1 if the connection is set to nonblocking mode and 0 if blocking.
<a id="libpq-PQflush"></a>

`PQflush`
:   Attempts to flush any queued output data to the server. Returns 0 if successful (or if the send queue is empty), -1 if it failed for some reason, or 1 if it was unable to send all the data in the send queue yet (this case can only occur if the connection is nonblocking).

    ```

    int PQflush(PGconn *conn);
    ```


 After sending any command or data on a nonblocking connection, call [PQflush](#libpq-PQflush). If it returns 1, wait for the socket to become read- or write-ready. If it becomes write-ready, call [PQflush](#libpq-PQflush) again. If it becomes read-ready, call [PQconsumeInput](#libpq-PQconsumeInput), then call [PQflush](#libpq-PQflush) again. Repeat until [PQflush](#libpq-PQflush) returns 0. (It is necessary to check for read-ready and drain the input with [PQconsumeInput](#libpq-PQconsumeInput), because the server can block trying to send us data, e.g., NOTICE messages, and won't read our data until we read its.) Once [PQflush](#libpq-PQflush) returns 0, wait for the socket to be read-ready and then read the response as described above.
