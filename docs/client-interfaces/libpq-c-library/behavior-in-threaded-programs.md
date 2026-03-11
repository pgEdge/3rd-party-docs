## Behavior in Threaded Programs { #libpq-threading }


 libpq is reentrant and thread-safe by default. You might need to use special compiler command-line options when you compile your application code. Refer to your system's documentation for information about how to build thread-enabled applications, or look in `src/Makefile.global` for `PTHREAD_CFLAGS` and `PTHREAD_LIBS`. This function allows the querying of libpq's thread-safe status:


<a id="libpq-PQisthreadsafe"></a>

`PQisthreadsafe`
:   Returns the thread safety status of the libpq library.

    ```

    int PQisthreadsafe();
    ```


     Returns 1 if the libpq is thread-safe and 0 if it is not.


 One thread restriction is that no two threads attempt to manipulate the same `PGconn` object at the same time. In particular, you cannot issue concurrent commands from different threads through the same connection object. (If you need to run concurrent commands, use multiple connections.)


 `PGresult` objects are normally read-only after creation, and so can be passed around freely between threads. However, if you use any of the `PGresult`-modifying functions described in [Miscellaneous Functions](miscellaneous-functions.md#libpq-misc) or [Event System](event-system.md#libpq-events), it's up to you to avoid concurrent operations on the same `PGresult`, too.


 The deprecated functions [PQrequestCancel](canceling-queries-in-progress.md#libpq-PQrequestCancel) and [PQoidStatus](command-execution-functions.md#libpq-PQoidStatus) are not thread-safe and should not be used in multithread programs. [PQrequestCancel](canceling-queries-in-progress.md#libpq-PQrequestCancel) can be replaced by [PQcancel](canceling-queries-in-progress.md#libpq-PQcancel). [PQoidStatus](command-execution-functions.md#libpq-PQoidStatus) can be replaced by [PQoidValue](command-execution-functions.md#libpq-PQoidValue).


 If you are using Kerberos inside your application (in addition to inside libpq), you will need to do locking around Kerberos calls because Kerberos functions are not thread-safe. See function `PQregisterThreadLock` in the libpq source code for a way to do cooperative locking between libpq and your application.
