<a id="libpq"></a>

# libpq — C Library

 libpq is the C application programmer's interface to PostgreSQL. libpq is a set of library functions that allow client programs to pass queries to the PostgreSQL backend server and to receive the results of these queries.

 libpq is also the underlying engine for several other PostgreSQL application interfaces, including those written for C++, Perl, Python, Tcl and ECPG. So some aspects of libpq's behavior will be important to you if you use one of those packages. In particular, [Environment Variables](environment-variables.md#libpq-envars), [The Password File](the-password-file.md#libpq-pgpass) and [SSL Support](ssl-support.md#libpq-ssl) describe behavior that is visible to the user of any application that uses libpq.

 Some short programs are included at the end of this chapter ([Example Programs](example-programs.md#libpq-example)) to show how to write programs that use libpq. There are also several complete examples of libpq applications in the directory `src/test/examples` in the source code distribution.

 Client programs that use libpq must include the header file `libpq-fe.h` and must link with the libpq library.

- [Database Connection Control Functions](database-connection-control-functions.md#libpq-connect)
- [Connection Status Functions](connection-status-functions.md#libpq-status)
- [Command Execution Functions](command-execution-functions.md#libpq-exec)
- [Asynchronous Command Processing](asynchronous-command-processing.md#libpq-async)
- [Pipeline Mode](pipeline-mode.md#libpq-pipeline-mode)
- [Retrieving Query Results in Chunks](retrieving-query-results-in-chunks.md#libpq-single-row-mode)
- [Canceling Queries in Progress](canceling-queries-in-progress.md#libpq-cancel)
- [The Fast-Path Interface](the-fast-path-interface.md#libpq-fastpath)
- [Asynchronous Notification](asynchronous-notification.md#libpq-notify)
- [Functions Associated with the `COPY` Command](functions-associated-with-the-copy-command.md#libpq-copy)
- [Control Functions](control-functions.md#libpq-control)
- [Miscellaneous Functions](miscellaneous-functions.md#libpq-misc)
- [Notice Processing](notice-processing.md#libpq-notice-processing)
- [Event System](event-system.md#libpq-events)
- [Environment Variables](environment-variables.md#libpq-envars)
- [The Password File](the-password-file.md#libpq-pgpass)
- [The Connection Service File](the-connection-service-file.md#libpq-pgservice)
- [LDAP Lookup of Connection Parameters](ldap-lookup-of-connection-parameters.md#libpq-ldap)
- [SSL Support](ssl-support.md#libpq-ssl)
- [OAuth Support](oauth-support.md#libpq-oauth)
- [Behavior in Threaded Programs](behavior-in-threaded-programs.md#libpq-threading)
- [Building libpq Programs](building-libpq-programs.md#libpq-build)
- [Example Programs](example-programs.md#libpq-example)
