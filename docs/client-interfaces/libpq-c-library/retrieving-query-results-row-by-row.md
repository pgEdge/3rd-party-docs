## Retrieving Query Results Row-by-Row { #libpq-single-row-mode }


 Ordinarily, libpq collects an SQL command's entire result and returns it to the application as a single `PGresult`. This can be unworkable for commands that return a large number of rows. For such cases, applications can use [PQsendQuery](asynchronous-command-processing.md#libpq-PQsendQuery) and [PQgetResult](asynchronous-command-processing.md#libpq-PQgetResult) in *single-row mode*. In this mode, the result row(s) are returned to the application one at a time, as they are received from the server.


 To enter single-row mode, call [PQsetSingleRowMode](#libpq-PQsetSingleRowMode) immediately after a successful call of [PQsendQuery](asynchronous-command-processing.md#libpq-PQsendQuery) (or a sibling function). This mode selection is effective only for the currently executing query. Then call [PQgetResult](asynchronous-command-processing.md#libpq-PQgetResult) repeatedly, until it returns null, as documented in [Asynchronous Command Processing](asynchronous-command-processing.md#libpq-async). If the query returns any rows, they are returned as individual `PGresult` objects, which look like normal query results except for having status code `PGRES_SINGLE_TUPLE` instead of `PGRES_TUPLES_OK`. After the last row, or immediately if the query returns zero rows, a zero-row object with status `PGRES_TUPLES_OK` is returned; this is the signal that no more rows will arrive. (But note that it is still necessary to continue calling [PQgetResult](asynchronous-command-processing.md#libpq-PQgetResult) until it returns null.) All of these `PGresult` objects will contain the same row description data (column names, types, etc.) that an ordinary `PGresult` object for the query would have. Each object should be freed with [PQclear](command-execution-functions.md#libpq-PQclear) as usual.


 When using pipeline mode, single-row mode needs to be activated for each query in the pipeline before retrieving results for that query with `PQgetResult`. See [Pipeline Mode](pipeline-mode.md#libpq-pipeline-mode) for more information.


<a id="libpq-PQsetSingleRowMode"></a>

`PQsetSingleRowMode`
:   Select single-row mode for the currently-executing query.

    ```

    int PQsetSingleRowMode(PGconn *conn);
    ```


     This function can only be called immediately after [PQsendQuery](asynchronous-command-processing.md#libpq-PQsendQuery) or one of its sibling functions, before any other operation on the connection such as [PQconsumeInput](asynchronous-command-processing.md#libpq-PQconsumeInput) or [PQgetResult](asynchronous-command-processing.md#libpq-PQgetResult). If called at the correct time, the function activates single-row mode for the current query and returns 1. Otherwise the mode stays unchanged and the function returns 0. In any case, the mode reverts to normal after completion of the current query.


!!! caution

    While processing a query, the server may return some rows and then encounter an error, causing the query to be aborted. Ordinarily, libpq discards any such rows and reports only the error. But in single-row mode, those rows will have already been returned to the application. Hence, the application will see some `PGRES_SINGLE_TUPLE` `PGresult` objects followed by a `PGRES_FATAL_ERROR` object. For proper transactional behavior, the application must be designed to discard or undo whatever has been done with the previously-processed rows, if the query ultimately fails.
