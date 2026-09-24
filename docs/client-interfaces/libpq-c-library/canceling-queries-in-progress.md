<a id="libpq-cancel"></a>

## Canceling Queries in Progress
    <a id="libpq-cancel-functions"></a>

### Functions for Sending Cancel Requests


<a id="libpq-PQcancelCreate"></a>

`PQcancelCreate`
:   Prepares a connection over which a cancel request can be sent.

    ```

    PGcancelConn *PQcancelCreate(PGconn *conn);
    ```


     [PQcancelCreate](#libpq-PQcancelCreate) creates a `PGcancelConn` object, but it won't instantly start sending a cancel request over this connection. A cancel request can be sent over this connection in a blocking manner using [PQcancelBlocking](#libpq-PQcancelBlocking) and in a non-blocking manner using [PQcancelStart](#libpq-PQcancelStart). The return value can be passed to [PQcancelStatus](#libpq-PQcancelStatus) to check if the `PGcancelConn` object was created successfully. The `PGcancelConn` object is an opaque structure that is not meant to be accessed directly by the application. This `PGcancelConn` object can be used to cancel the query that's running on the original connection in a thread-safe way.


     Many connection parameters of the original client will be reused when setting up the connection for the cancel request. Importantly, if the original connection requires encryption of the connection and/or verification of the target host (using `sslmode` or `gssencmode`), then the connection for the cancel request is made with these same requirements. Any connection options that are only used during authentication or after authentication of the client are ignored though, because cancellation requests do not require authentication and the connection is closed right after the cancellation request is submitted.


     Note that when `PQcancelCreate` returns a non-null pointer, you must call [PQcancelFinish](#libpq-PQcancelFinish) when you are finished with it, in order to dispose of the structure and any associated memory blocks. This must be done even if the cancel request failed or was abandoned.
<a id="libpq-PQcancelBlocking"></a>

`PQcancelBlocking`
:   Requests that the server abandons processing of the current command in a blocking manner.

    ```

    int PQcancelBlocking(PGcancelConn *cancelConn);
    ```


     The request is made over the given `PGcancelConn`, which needs to be created with [PQcancelCreate](#libpq-PQcancelCreate). The return value of [PQcancelBlocking](#libpq-PQcancelBlocking) is 1 if the cancel request was successfully dispatched and 0 if not. If it was unsuccessful, the error message can be retrieved using [PQcancelErrorMessage](#libpq-PQcancelErrorMessage).


     Successful dispatch of the cancellation is no guarantee that the request will have any effect, however. If the cancellation is effective, the command being canceled will terminate early and return an error result. If the cancellation fails (say, because the server was already done processing the command), then there will be no visible result at all.
<a id="libpq-PQcancelStart"></a>

`PQcancelStart`, `PQcancelPoll`
:   Requests that the server abandons processing of the current command in a non-blocking manner.

    ```

    int PQcancelStart(PGcancelConn *cancelConn);

    PostgresPollingStatusType PQcancelPoll(PGcancelConn *cancelConn);
    ```


     The request is made over the given `PGcancelConn`, which needs to be created with [PQcancelCreate](#libpq-PQcancelCreate). The return value of [PQcancelStart](#libpq-PQcancelStart) is 1 if the cancellation request could be started and 0 if not. If it was unsuccessful, the error message can be retrieved using [PQcancelErrorMessage](#libpq-PQcancelErrorMessage).


     If `PQcancelStart` succeeds, the next stage is to poll libpq so that it can proceed with the cancel connection sequence. Use [PQcancelSocket](#libpq-PQcancelSocket) to obtain the descriptor of the socket underlying the database connection. (Caution: do not assume that the socket remains the same across `PQcancelPoll` calls.) Loop thus: If `PQcancelPoll(cancelConn)` last returned `PGRES_POLLING_READING`, wait until the socket is ready to read (as indicated by `select()`, `poll()`, or similar system function). Then call `PQcancelPoll(cancelConn)` again. Conversely, if `PQcancelPoll(cancelConn)` last returned `PGRES_POLLING_WRITING`, wait until the socket is ready to write, then call `PQcancelPoll(cancelConn)` again. On the first iteration, i.e., if you have yet to call `PQcancelPoll(cancelConn)`, behave as if it last returned `PGRES_POLLING_WRITING`. Continue this loop until `PQcancelPoll(cancelConn)` returns `PGRES_POLLING_FAILED`, indicating the connection procedure has failed, or `PGRES_POLLING_OK`, indicating cancel request was successfully dispatched.


     Successful dispatch of the cancellation is no guarantee that the request will have any effect, however. If the cancellation is effective, the command being canceled will terminate early and return an error result. If the cancellation fails (say, because the server was already done processing the command), then there will be no visible result at all.


     At any time during connection, the status of the connection can be checked by calling [PQcancelStatus](#libpq-PQcancelStatus). If this call returns `CONNECTION_BAD`, then the cancel procedure has failed; if the call returns `CONNECTION_OK`, then cancel request was successfully dispatched. Both of these states are equally detectable from the return value of `PQcancelPoll`, described above. Other states might also occur during (and only during) an asynchronous connection procedure. These indicate the current stage of the connection procedure and might be useful to provide feedback to the user for example. These statuses are:

    <a id="libpq-cancel-connection-allocated"></a>

    `CONNECTION_ALLOCATED`
    :   Waiting for a call to [PQcancelStart](#libpq-PQcancelStart) or [PQcancelBlocking](#libpq-PQcancelBlocking), to actually open the socket. This is the connection state right after calling [PQcancelCreate](#libpq-PQcancelCreate) or [PQcancelReset](#libpq-PQcancelReset). No connection to the server has been initiated yet at this point. To actually start sending the cancel request use [PQcancelStart](#libpq-PQcancelStart) or [PQcancelBlocking](#libpq-PQcancelBlocking).
    <a id="libpq-cancel-connection-started"></a>

    `CONNECTION_STARTED`
    :   Waiting for connection to be made.
    <a id="libpq-cancel-connection-made"></a>

    `CONNECTION_MADE`
    :   Connection OK; waiting to send.
    <a id="libpq-cancel-connection-awaiting-response"></a>

    `CONNECTION_AWAITING_RESPONSE`
    :   Waiting for a response from the server.
    <a id="libpq-cancel-connection-ssl-startup"></a>

    `CONNECTION_SSL_STARTUP`
    :   Negotiating SSL encryption.
    <a id="libpq-cancel-connection-gss-startup"></a>

    `CONNECTION_GSS_STARTUP`
    :   Negotiating GSS encryption.
     Note that, although these constants will remain (in order to maintain compatibility), an application should never rely upon these occurring in a particular order, or at all, or on the status always being one of these documented values. An application might do something like this:

    ```

    switch(PQcancelStatus(conn))
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


     The `connect_timeout` connection parameter is ignored when using `PQcancelPoll`; it is the application's responsibility to decide whether an excessive amount of time has elapsed. Otherwise, `PQcancelStart` followed by a `PQcancelPoll` loop is equivalent to [PQcancelBlocking](#libpq-PQcancelBlocking).
<a id="libpq-PQcancelStatus"></a>

`PQcancelStatus`
:   Returns the status of the cancel connection.

    ```

    ConnStatusType PQcancelStatus(const PGcancelConn *cancelConn);
    ```


     The status can be one of a number of values. However, only three of these are seen outside of an asynchronous cancel procedure: `CONNECTION_ALLOCATED`, `CONNECTION_OK` and `CONNECTION_BAD`. The initial state of a `PGcancelConn` that's successfully created using [PQcancelCreate](#libpq-PQcancelCreate) is `CONNECTION_ALLOCATED`. A cancel request that was successfully dispatched has the status `CONNECTION_OK`. A failed cancel attempt is signaled by status `CONNECTION_BAD`. An OK status will remain so until [PQcancelFinish](#libpq-PQcancelFinish) or [PQcancelReset](#libpq-PQcancelReset) is called.


     See the entry for [PQcancelStart](#libpq-PQcancelStart) with regards to other status codes that might be returned.


     Successful dispatch of the cancellation is no guarantee that the request will have any effect, however. If the cancellation is effective, the command being canceled will terminate early and return an error result. If the cancellation fails (say, because the server was already done processing the command), then there will be no visible result at all.
<a id="libpq-PQcancelSocket"></a>

`PQcancelSocket`
:   Obtains the file descriptor number of the cancel connection socket to the server.

    ```

    int PQcancelSocket(const PGcancelConn *cancelConn);
    ```


     A valid descriptor will be greater than or equal to 0; a result of -1 indicates that no server connection is currently open. This might change as a result of calling any of the functions in this section on the `PGcancelConn` (except for [PQcancelErrorMessage](#libpq-PQcancelErrorMessage) and `PQcancelSocket` itself).
<a id="libpq-PQcancelErrorMessage"></a>

`PQcancelErrorMessage`
:   Returns the error message most recently generated by an operation on the cancel connection.

    ```

    char *PQcancelErrorMessage(const PGcancelConn *cancelconn);
    ```


     Nearly all libpq functions that take a `PGcancelConn` will set a message for [PQcancelErrorMessage](#libpq-PQcancelErrorMessage) if they fail. Note that by libpq convention, a nonempty [PQcancelErrorMessage](#libpq-PQcancelErrorMessage) result can consist of multiple lines, and will include a trailing newline. The caller should not free the result directly. It will be freed when the associated `PGcancelConn` handle is passed to [PQcancelFinish](#libpq-PQcancelFinish). The result string should not be expected to remain the same across operations on the `PGcancelConn` structure.
<a id="libpq-PQcancelFinish"></a>

`PQcancelFinish`
:   Closes the cancel connection (if it did not finish sending the cancel request yet). Also frees memory used by the `PGcancelConn` object.

    ```

    void PQcancelFinish(PGcancelConn *cancelConn);
    ```


     Note that even if the cancel attempt fails (as indicated by [PQcancelStatus](#libpq-PQcancelStatus)), the application should call [PQcancelFinish](#libpq-PQcancelFinish) to free the memory used by the `PGcancelConn` object. The `PGcancelConn` pointer must not be used again after [PQcancelFinish](#libpq-PQcancelFinish) has been called.
<a id="libpq-PQcancelReset"></a>

`PQcancelReset`
:   Resets the `PGcancelConn` so it can be reused for a new cancel connection.

    ```

    void PQcancelReset(PGcancelConn *cancelConn);
    ```


     If the `PGcancelConn` is currently used to send a cancel request, then this connection is closed. It will then prepare the `PGcancelConn` object such that it can be used to send a new cancel request.


     This can be used to create one `PGcancelConn` for a `PGconn` and reuse it multiple times throughout the lifetime of the original `PGconn`.
  <a id="libpq-cancel-deprecated"></a>

### Obsolete Functions for Sending Cancel Requests


 These functions represent older methods of sending cancel requests. Although they still work, they are deprecated due to not sending the cancel requests in an encrypted manner, even when the original connection specified `sslmode` or `gssencmode` to require encryption. Thus these older methods are heavily discouraged from being used in new code, and it is recommended to change existing code to use the new functions instead.


<a id="libpq-PQgetCancel"></a>

`PQgetCancel`
:   Creates a data structure containing the information needed to cancel a command using [PQcancel](#libpq-PQcancel).

    ```

    PGcancel *PQgetCancel(PGconn *conn);
    ```


     [PQgetCancel](#libpq-PQgetCancel) creates a `PGcancel` object given a `PGconn` connection object. It will return `NULL` if the given `conn` is `NULL` or an invalid connection. The `PGcancel` object is an opaque structure that is not meant to be accessed directly by the application; it can only be passed to [PQcancel](#libpq-PQcancel) or [PQfreeCancel](#libpq-PQfreeCancel).
<a id="libpq-PQfreeCancel"></a>

`PQfreeCancel`
:   Frees a data structure created by [PQgetCancel](#libpq-PQgetCancel).

    ```

    void PQfreeCancel(PGcancel *cancel);
    ```


     [PQfreeCancel](#libpq-PQfreeCancel) frees a data object previously created by [PQgetCancel](#libpq-PQgetCancel).
<a id="libpq-PQcancel"></a>

`PQcancel`
:   [PQcancel](#libpq-PQcancel) is a deprecated and insecure variant of [PQcancelBlocking](#libpq-PQcancelBlocking), but one that can be used safely from within a signal handler.

    ```

    int PQcancel(PGcancel *cancel, char *errbuf, int errbufsize);
    ```


     [PQcancel](#libpq-PQcancel) only exists because of backwards compatibility reasons. [PQcancelBlocking](#libpq-PQcancelBlocking) should be used instead. The only benefit that [PQcancel](#libpq-PQcancel) has is that it can be safely invoked from a signal handler, if the `errbuf` is a local variable in the signal handler. However, this is generally not considered a big enough benefit to be worth the security issues that this function has.


     The `PGcancel` object is read-only as far as [PQcancel](#libpq-PQcancel) is concerned, so it can also be invoked from a thread that is separate from the one manipulating the `PGconn` object.


     The return value of [PQcancel](#libpq-PQcancel) is 1 if the cancel request was successfully dispatched and 0 if not. If not, `errbuf` is filled with an explanatory error message. `errbuf` must be a char array of size `errbufsize` (the recommended size is 256 bytes).


<a id="libpq-PQrequestCancel"></a>

`PQrequestCancel`
:   [PQrequestCancel](#libpq-PQrequestCancel) is a deprecated and insecure variant of [PQcancelBlocking](#libpq-PQcancelBlocking).

    ```

    int PQrequestCancel(PGconn *conn);
    ```


     [PQrequestCancel](#libpq-PQrequestCancel) only exists because of backwards compatibility reasons. [PQcancelBlocking](#libpq-PQcancelBlocking) should be used instead. There is no benefit to using [PQrequestCancel](#libpq-PQrequestCancel) over [PQcancelBlocking](#libpq-PQcancelBlocking).


     Requests that the server abandon processing of the current command. It operates directly on the `PGconn` object, and in case of failure stores the error message in the `PGconn` object (whence it can be retrieved by [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage)). Although the functionality is the same, this approach is not safe within multiple-thread programs or signal handlers, since it is possible that overwriting the `PGconn`'s error message will mess up the operation currently in progress on the connection.
