## Canceling Queries in Progress { #libpq-cancel }


 A client application can request cancellation of a command that is still being processed by the server, using the functions described in this section.

<a id="libpq-PQgetCancel"></a>

`PQgetCancel`
:   Creates a data structure containing the information needed to cancel a command issued through a particular database connection.

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
:   Requests that the server abandon processing of the current command.

    ```

    int PQcancel(PGcancel *cancel, char *errbuf, int errbufsize);
    ```


     The return value is 1 if the cancel request was successfully dispatched and 0 if not. If not, `errbuf` is filled with an explanatory error message. `errbuf` must be a char array of size `errbufsize` (the recommended size is 256 bytes).


     Successful dispatch is no guarantee that the request will have any effect, however. If the cancellation is effective, the current command will terminate early and return an error result. If the cancellation fails (say, because the server was already done processing the command), then there will be no visible result at all.


     [PQcancel](#libpq-PQcancel) can safely be invoked from a signal handler, if the `errbuf` is a local variable in the signal handler. The `PGcancel` object is read-only as far as [PQcancel](#libpq-PQcancel) is concerned, so it can also be invoked from a thread that is separate from the one manipulating the `PGconn` object.


<a id="libpq-PQrequestCancel"></a>

`PQrequestCancel`
:   [PQrequestCancel](#libpq-PQrequestCancel) is a deprecated variant of [PQcancel](#libpq-PQcancel).

    ```

    int PQrequestCancel(PGconn *conn);
    ```


     Requests that the server abandon processing of the current command. It operates directly on the `PGconn` object, and in case of failure stores the error message in the `PGconn` object (whence it can be retrieved by [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage)). Although the functionality is the same, this approach is not safe within multiple-thread programs or signal handlers, since it is possible that overwriting the `PGconn`'s error message will mess up the operation currently in progress on the connection.
