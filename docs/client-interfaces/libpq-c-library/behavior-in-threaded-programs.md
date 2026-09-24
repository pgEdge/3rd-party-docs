<a id="libpq-threading"></a>

## Behavior in Threaded Programs


 As of version 17, libpq is always reentrant and thread-safe. However, one restriction is that no two threads attempt to manipulate the same `PGconn` object at the same time. In particular, you cannot issue concurrent commands from different threads through the same connection object. (If you need to run concurrent commands, use multiple connections.)


 `PGresult` objects are normally read-only after creation, and so can be passed around freely between threads. However, if you use any of the `PGresult`-modifying functions described in [Miscellaneous Functions](miscellaneous-functions.md#libpq-misc) or [Event System](event-system.md#libpq-events), it's up to you to avoid concurrent operations on the same `PGresult`, too.


 In earlier versions, libpq could be compiled with or without thread support, depending on compiler options. This function allows the querying of libpq's thread-safe status:


<a id="libpq-PQisthreadsafe"></a>

`PQisthreadsafe`
:   Returns the thread safety status of the libpq library.

    ```

    int PQisthreadsafe();
    ```


     Returns 1 if the libpq is thread-safe and 0 if it is not. Always returns 1 on version 17 and above.


 The deprecated functions [PQrequestCancel](canceling-queries-in-progress.md#libpq-PQrequestCancel) and [PQoidStatus](command-execution-functions.md#libpq-PQoidStatus) are not thread-safe and should not be used in multithread programs. [PQrequestCancel](canceling-queries-in-progress.md#libpq-PQrequestCancel) can be replaced by [PQcancelBlocking](canceling-queries-in-progress.md#libpq-PQcancelBlocking). [PQoidStatus](command-execution-functions.md#libpq-PQoidStatus) can be replaced by [PQoidValue](command-execution-functions.md#libpq-PQoidValue).


 If you are using Kerberos inside your application (in addition to inside libpq), you will need to do locking around Kerberos calls because Kerberos functions are not thread-safe. See function `PQregisterThreadLock` in the libpq source code for a way to do cooperative locking between libpq and your application. (Note that it is only safe to call `PQregisterThreadLock` when there are no open connections.) Clients may retrieve the current locking callback with `PQgetThreadLock`; if no custom callback has been registered, a default is used.


 Similarly, if you are using Curl inside your application, *and* you do not already [initialize libcurl globally](https://curl.se/libcurl/c/curl_global_init.html) before starting new threads, you will need to cooperatively lock (again via `PQregisterThreadLock`) around any code that may initialize libcurl. This restriction is lifted for more recent versions of Curl that are built to support thread-safe initialization; those builds can be identified by the advertisement of a `threadsafe` feature in their version metadata.
