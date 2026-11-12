<a id="libpq-exec"></a>

## Command Execution Functions


 Once a connection to a database server has been successfully established, the functions described here are used to perform SQL queries and commands.
 <a id="libpq-exec-main"></a>

### Main Functions


<a id="libpq-PQexec"></a>

`PQexec`
:   Submits a command to the server and waits for the result.

    ```

    PGresult *PQexec(PGconn *conn, const char *command);
    ```


     Returns a `PGresult` pointer or possibly a null pointer. A non-null pointer will generally be returned except in out-of-memory conditions or serious errors such as inability to send the command to the server. The [PQresultStatus](#libpq-PQresultStatus) function should be called to check the return value for any errors (including the value of a null pointer, in which case it will return `PGRES_FATAL_ERROR`). Use [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage) to get more information about such errors.
 The command string can include multiple SQL commands (separated by semicolons). Multiple queries sent in a single [PQexec](#libpq-PQexec) call are processed in a single transaction, unless there are explicit `BEGIN`/`COMMIT` commands included in the query string to divide it into multiple transactions. (See [Multiple Statements in a Simple Query](../../internals/frontend-backend-protocol/message-flow.md#protocol-flow-multi-statement) for more details about how the server handles multi-query strings.) Note however that the returned `PGresult` structure describes only the result of the last command executed from the string. Should one of the commands fail, processing of the string stops with it and the returned `PGresult` describes the error condition.


<a id="libpq-PQexecParams"></a>

`PQexecParams`
:   Submits a command to the server and waits for the result, with the ability to pass parameters separately from the SQL command text.

    ```

    PGresult *PQexecParams(PGconn *conn,
                           const char *command,
                           int nParams,
                           const Oid *paramTypes,
                           const char * const *paramValues,
                           const int *paramLengths,
                           const int *paramFormats,
                           int resultFormat);
    ```


     [PQexecParams](#libpq-PQexecParams) is like [PQexec](#libpq-PQexec), but offers additional functionality: parameter values can be specified separately from the command string proper, and query results can be requested in either text or binary format.


     The function arguments are:

    `conn`
    :   The connection object to send the command through.

    `command`
    :   The SQL command string to be executed. If parameters are used, they are referred to in the command string as `$1`, `$2`, etc.

    `nParams`
    :   The number of parameters supplied; it is the length of the arrays `paramTypes[]`, `paramValues[]`, `paramLengths[]`, and `paramFormats[]`. (The array pointers can be `NULL` when `nParams` is zero.)

    `paramTypes[]`
    :   Specifies, by OID, the data types to be assigned to the parameter symbols. If `paramTypes` is `NULL`, or any particular element in the array is zero, the server infers a data type for the parameter symbol in the same way it would do for an untyped literal string.

    `paramValues[]`
    :   Specifies the actual values of the parameters. A null pointer in this array means the corresponding parameter is null; otherwise the pointer points to a zero-terminated text string (for text format) or binary data in the format expected by the server (for binary format).

    `paramLengths[]`
    :   Specifies the actual data lengths of binary-format parameters. It is ignored for null parameters and text-format parameters. The array pointer can be null when there are no binary parameters.

    `paramFormats[]`
    :   Specifies whether parameters are text (put a zero in the array entry for the corresponding parameter) or binary (put a one in the array entry for the corresponding parameter). If the array pointer is null then all parameters are presumed to be text strings.


         Values passed in binary format require knowledge of the internal representation expected by the backend. For example, integers must be passed in network byte order. Passing `numeric` values requires knowledge of the server storage format, as implemented in `src/backend/utils/adt/numeric.c::numeric_send()` and `src/backend/utils/adt/numeric.c::numeric_recv()`.

    `resultFormat`
    :   Specify zero to obtain results in text format, or one to obtain results in binary format. (There is not currently a provision to obtain different result columns in different formats, although that is possible in the underlying protocol.)


 The primary advantage of [PQexecParams](#libpq-PQexecParams) over [PQexec](#libpq-PQexec) is that parameter values can be separated from the command string, thus avoiding the need for tedious and error-prone quoting and escaping.


 Unlike [PQexec](#libpq-PQexec), [PQexecParams](#libpq-PQexecParams) allows at most one SQL command in the given string. (There can be semicolons in it, but not more than one nonempty command.) This is a limitation of the underlying protocol, but has some usefulness as an extra defense against SQL-injection attacks.


!!! tip

    Specifying parameter types via OIDs is tedious, particularly if you prefer not to hard-wire particular OID values into your program. However, you can avoid doing so even in cases where the server by itself cannot determine the type of the parameter, or chooses a different type than you want. In the SQL command text, attach an explicit cast to the parameter symbol to show what data type you will send. For example:

    ```sql

    SELECT * FROM mytable WHERE x = $1::bigint;
    ```
     This forces parameter `$1` to be treated as `bigint`, whereas by default it would be assigned the same type as `x`. Forcing the parameter type decision, either this way or by specifying a numeric type OID, is strongly recommended when sending parameter values in binary format, because binary format has less redundancy than text format and so there is less chance that the server will detect a type mismatch mistake for you.


<a id="libpq-PQprepare"></a>

`PQprepare`
:   Submits a request to create a prepared statement with the given parameters, and waits for completion.

    ```

    PGresult *PQprepare(PGconn *conn,
                        const char *stmtName,
                        const char *query,
                        int nParams,
                        const Oid *paramTypes);
    ```


     [PQprepare](#libpq-PQprepare) creates a prepared statement for later execution with [PQexecPrepared](#libpq-PQexecPrepared). This feature allows commands to be executed repeatedly without being parsed and planned each time; see [sql-prepare](../../reference/sql-commands/prepare.md#sql-prepare) for details.


     The function creates a prepared statement named `stmtName` from the `query` string, which must contain a single SQL command. `stmtName` can be `""` to create an unnamed statement, in which case any pre-existing unnamed statement is automatically replaced; otherwise it is an error if the statement name is already defined in the current session. If any parameters are used, they are referred to in the query as `$1`, `$2`, etc. `nParams` is the number of parameters for which types are pre-specified in the array `paramTypes[]`. (The array pointer can be `NULL` when `nParams` is zero.) `paramTypes[]` specifies, by OID, the data types to be assigned to the parameter symbols. If `paramTypes` is `NULL`, or any particular element in the array is zero, the server assigns a data type to the parameter symbol in the same way it would do for an untyped literal string. Also, the query can use parameter symbols with numbers higher than `nParams`; data types will be inferred for these symbols as well. (See [PQdescribePrepared](#libpq-PQdescribePrepared) for a means to find out what data types were inferred.)


     As with [PQexec](#libpq-PQexec), the result is normally a `PGresult` object whose contents indicate server-side success or failure. A null result indicates out-of-memory or inability to send the command at all. Use [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage) to get more information about such errors.
 Prepared statements for use with [PQexecPrepared](#libpq-PQexecPrepared) can also be created by executing SQL [sql-prepare](../../reference/sql-commands/prepare.md#sql-prepare) statements.


<a id="libpq-PQexecPrepared"></a>

`PQexecPrepared`
:   Sends a request to execute a prepared statement with given parameters, and waits for the result.

    ```

    PGresult *PQexecPrepared(PGconn *conn,
                             const char *stmtName,
                             int nParams,
                             const char * const *paramValues,
                             const int *paramLengths,
                             const int *paramFormats,
                             int resultFormat);
    ```


     [PQexecPrepared](#libpq-PQexecPrepared) is like [PQexecParams](#libpq-PQexecParams), but the command to be executed is specified by naming a previously-prepared statement, instead of giving a query string. This feature allows commands that will be used repeatedly to be parsed and planned just once, rather than each time they are executed. The statement must have been prepared previously in the current session.


     The parameters are identical to [PQexecParams](#libpq-PQexecParams), except that the name of a prepared statement is given instead of a query string, and the `paramTypes[]` parameter is not present (it is not needed since the prepared statement's parameter types were determined when it was created).
<a id="libpq-PQdescribePrepared"></a>

`PQdescribePrepared`
:   Submits a request to obtain information about the specified prepared statement, and waits for completion.

    ```

    PGresult *PQdescribePrepared(PGconn *conn, const char *stmtName);
    ```


     [PQdescribePrepared](#libpq-PQdescribePrepared) allows an application to obtain information about a previously prepared statement.


     `stmtName` can be `""` or `NULL` to reference the unnamed statement, otherwise it must be the name of an existing prepared statement. On success, a `PGresult` with status `PGRES_COMMAND_OK` is returned. The functions [PQnparams](#libpq-PQnparams) and [PQparamtype](#libpq-PQparamtype) can be applied to this `PGresult` to obtain information about the parameters of the prepared statement, and the functions [PQnfields](#libpq-PQnfields), [PQfname](#libpq-PQfname), [PQftype](#libpq-PQftype), etc. provide information about the result columns (if any) of the statement.
<a id="libpq-PQdescribePortal"></a>

`PQdescribePortal`
:   Submits a request to obtain information about the specified portal, and waits for completion.

    ```

    PGresult *PQdescribePortal(PGconn *conn, const char *portalName);
    ```


     [PQdescribePortal](#libpq-PQdescribePortal) allows an application to obtain information about a previously created portal. (libpq does not provide any direct access to portals, but you can use this function to inspect the properties of a cursor created with a `DECLARE CURSOR` SQL command.)


     `portalName` can be `""` or `NULL` to reference the unnamed portal, otherwise it must be the name of an existing portal. On success, a `PGresult` with status `PGRES_COMMAND_OK` is returned. The functions [PQnfields](#libpq-PQnfields), [PQfname](#libpq-PQfname), [PQftype](#libpq-PQftype), etc. can be applied to the `PGresult` to obtain information about the result columns (if any) of the portal.
<a id="libpq-PQclosePrepared"></a>

`PQclosePrepared`
:   Submits a request to close the specified prepared statement, and waits for completion.

    ```

    PGresult *PQclosePrepared(PGconn *conn, const char *stmtName);
    ```


     [PQclosePrepared](#libpq-PQclosePrepared) allows an application to close a previously prepared statement. Closing a statement releases all of its associated resources on the server and allows its name to be reused.


     `stmtName` can be `""` or `NULL` to reference the unnamed statement. It is fine if no statement exists with this name, in that case the operation is a no-op. On success, a `PGresult` with status `PGRES_COMMAND_OK` is returned.
<a id="libpq-PQclosePortal"></a>

`PQclosePortal`
:   Submits a request to close the specified portal, and waits for completion.

    ```

    PGresult *PQclosePortal(PGconn *conn, const char *portalName);
    ```


     [PQclosePortal](#libpq-PQclosePortal) allows an application to trigger a close of a previously created portal. Closing a portal releases all of its associated resources on the server and allows its name to be reused. (libpq does not provide any direct access to portals, but you can use this function to close a cursor created with a `DECLARE CURSOR` SQL command.)


     `portalName` can be `""` or `NULL` to reference the unnamed portal. It is fine if no portal exists with this name, in that case the operation is a no-op. On success, a `PGresult` with status `PGRES_COMMAND_OK` is returned.


 The `PGresult` structure encapsulates the result returned by the server. libpq application programmers should be careful to maintain the `PGresult` abstraction. Use the accessor functions below to get at the contents of `PGresult`. Avoid directly referencing the fields of the `PGresult` structure because they are subject to change in the future.

<a id="libpq-PQresultStatus"></a>

`PQresultStatus`
:   Returns the result status of the command.

    ```

    ExecStatusType PQresultStatus(const PGresult *res);
    ```


     [PQresultStatus](#libpq-PQresultStatus) can return one of the following values:

    <a id="libpq-pgres-empty-query"></a>

    `PGRES_EMPTY_QUERY`
    :   The string sent to the server was empty.
    <a id="libpq-pgres-command-ok"></a>

    `PGRES_COMMAND_OK`
    :   Successful completion of a command returning no data.
    <a id="libpq-pgres-tuples-ok"></a>

    `PGRES_TUPLES_OK`
    :   Successful completion of a command returning data (such as a `SELECT` or `SHOW`).
    <a id="libpq-pgres-copy-out"></a>

    `PGRES_COPY_OUT`
    :   Copy Out (from server) data transfer started.
    <a id="libpq-pgres-copy-in"></a>

    `PGRES_COPY_IN`
    :   Copy In (to server) data transfer started.
    <a id="libpq-pgres-bad-response"></a>

    `PGRES_BAD_RESPONSE`
    :   The server's response was not understood.
    <a id="libpq-pgres-nonfatal-error"></a>

    `PGRES_NONFATAL_ERROR`
    :   A nonfatal error (a notice or warning) occurred.
    <a id="libpq-pgres-fatal-error"></a>

    `PGRES_FATAL_ERROR`
    :   A fatal error occurred.
    <a id="libpq-pgres-copy-both"></a>

    `PGRES_COPY_BOTH`
    :   Copy In/Out (to and from server) data transfer started. This feature is currently used only for streaming replication, so this status should not occur in ordinary applications.
    <a id="libpq-pgres-single-tuple"></a>

    `PGRES_SINGLE_TUPLE`
    :   The `PGresult` contains a single result tuple from the current command. This status occurs only when single-row mode has been selected for the query (see [Retrieving Query Results in Chunks](retrieving-query-results-in-chunks.md#libpq-single-row-mode)).
    <a id="libpq-pgres-tuples-chunk"></a>

    `PGRES_TUPLES_CHUNK`
    :   The `PGresult` contains several result tuples from the current command. This status occurs only when chunked mode has been selected for the query (see [Retrieving Query Results in Chunks](retrieving-query-results-in-chunks.md#libpq-single-row-mode)). The number of tuples will not exceed the limit passed to [PQsetChunkedRowsMode](retrieving-query-results-in-chunks.md#libpq-PQsetChunkedRowsMode).
    <a id="libpq-pgres-pipeline-sync"></a>

    `PGRES_PIPELINE_SYNC`
    :   The `PGresult` represents a synchronization point in pipeline mode, requested by either [PQpipelineSync](pipeline-mode.md#libpq-PQpipelineSync) or [PQsendPipelineSync](pipeline-mode.md#libpq-PQsendPipelineSync). This status occurs only when pipeline mode has been selected.
    <a id="libpq-pgres-pipeline-aborted"></a>

    `PGRES_PIPELINE_ABORTED`
    :   The `PGresult` represents a pipeline that has received an error from the server. `PQgetResult` must be called repeatedly, and each time it will return this status code until the end of the current pipeline, at which point it will return `PGRES_PIPELINE_SYNC` and normal processing can resume.
     If the result status is `PGRES_TUPLES_OK`, `PGRES_SINGLE_TUPLE`, or `PGRES_TUPLES_CHUNK`, then the functions described below can be used to retrieve the rows returned by the query. Note that a `SELECT` command that happens to retrieve zero rows still shows `PGRES_TUPLES_OK`. `PGRES_COMMAND_OK` is for commands that can never return rows (`INSERT` or `UPDATE` without a `RETURNING` clause, etc.). A response of `PGRES_EMPTY_QUERY` might indicate a bug in the client software.


     A result of status `PGRES_NONFATAL_ERROR` will never be returned directly by [PQexec](#libpq-PQexec) or other query execution functions; results of this kind are instead passed to the notice processor (see [Notice Processing](notice-processing.md#libpq-notice-processing)).
<a id="libpq-PQresStatus"></a>

`PQresStatus`
:   Converts the enumerated type returned by [PQresultStatus](#libpq-PQresultStatus) into a string constant describing the status code. The caller should not free the result.

    ```

    char *PQresStatus(ExecStatusType status);
    ```
<a id="libpq-PQresultErrorMessage"></a>

`PQresultErrorMessage`
:   Returns the error message associated with the command, or an empty string if there was no error.

    ```

    char *PQresultErrorMessage(const PGresult *res);
    ```
     If there was an error, the returned string will include a trailing newline. The caller should not free the result directly. It will be freed when the associated `PGresult` handle is passed to [PQclear](#libpq-PQclear).


     Immediately following a [PQexec](#libpq-PQexec) or [PQgetResult](asynchronous-command-processing.md#libpq-PQgetResult) call, [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage) (on the connection) will return the same string as [PQresultErrorMessage](#libpq-PQresultErrorMessage) (on the result). However, a `PGresult` will retain its error message until destroyed, whereas the connection's error message will change when subsequent operations are done. Use [PQresultErrorMessage](#libpq-PQresultErrorMessage) when you want to know the status associated with a particular `PGresult`; use [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage) when you want to know the status from the latest operation on the connection.
<a id="libpq-PQresultVerboseErrorMessage"></a>

`PQresultVerboseErrorMessage`
:   Returns a reformatted version of the error message associated with a `PGresult` object.

    ```

    char *PQresultVerboseErrorMessage(const PGresult *res,
                                      PGVerbosity verbosity,
                                      PGContextVisibility show_context);
    ```
     In some situations a client might wish to obtain a more detailed version of a previously-reported error. [PQresultVerboseErrorMessage](#libpq-PQresultVerboseErrorMessage) addresses this need by computing the message that would have been produced by [PQresultErrorMessage](#libpq-PQresultErrorMessage) if the specified verbosity settings had been in effect for the connection when the given `PGresult` was generated. If the `PGresult` is not an error result, “PGresult is not an error result” is reported instead. The returned string includes a trailing newline.


     Unlike most other functions for extracting data from a `PGresult`, the result of this function is a freshly allocated string. The caller must free it using `PQfreemem()` when the string is no longer needed.


     A NULL return is possible if there is insufficient memory.
<a id="libpq-PQresultErrorField"></a>

`PQresultErrorField`
:   Returns an individual field of an error report.

    ```

    char *PQresultErrorField(const PGresult *res, int fieldcode);
    ```
     `fieldcode` is an error field identifier; see the symbols listed below. `NULL` is returned if the `PGresult` is not an error or warning result, or does not include the specified field. Field values will normally not include a trailing newline. The caller should not free the result directly. It will be freed when the associated `PGresult` handle is passed to [PQclear](#libpq-PQclear).


     The following field codes are available:

    <a id="libpq-pg-diag-severity"></a>

    `PG_DIAG_SEVERITY`
    :   The severity; the field contents are `ERROR`, `FATAL`, or `PANIC` (in an error message), or `WARNING`, `NOTICE`, `DEBUG`, `INFO`, or `LOG` (in a notice message), or a localized translation of one of these. Always present.
    <a id="libpq-PG-diag-severity-nonlocalized"></a>

    `PG_DIAG_SEVERITY_NONLOCALIZED`
    :   The severity; the field contents are `ERROR`, `FATAL`, or `PANIC` (in an error message), or `WARNING`, `NOTICE`, `DEBUG`, `INFO`, or `LOG` (in a notice message). This is identical to the `PG_DIAG_SEVERITY` field except that the contents are never localized. This is present only in reports generated by PostgreSQL versions 9.6 and later.
    <a id="libpq-pg-diag-sqlstate"></a>

    `PG_DIAG_SQLSTATE`
    :   The SQLSTATE code for the error. The SQLSTATE code identifies the type of error that has occurred; it can be used by front-end applications to perform specific operations (such as error handling) in response to a particular database error. For a list of the possible SQLSTATE codes, see [PostgreSQL Error Codes](../../appendixes/postgresql-error-codes.md#errcodes-appendix). This field is not localizable, and is always present.
    <a id="libpq-pg-diag-message-primary"></a>

    `PG_DIAG_MESSAGE_PRIMARY`
    :   The primary human-readable error message (typically one line). Always present.
    <a id="libpq-pg-diag-message-detail"></a>

    `PG_DIAG_MESSAGE_DETAIL`
    :   Detail: an optional secondary error message carrying more detail about the problem. Might run to multiple lines.
    <a id="libpq-pg-diag-message-hint"></a>

    `PG_DIAG_MESSAGE_HINT`
    :   Hint: an optional suggestion what to do about the problem. This is intended to differ from detail in that it offers advice (potentially inappropriate) rather than hard facts. Might run to multiple lines.
    <a id="libpq-pg-diag-statement-position"></a>

    `PG_DIAG_STATEMENT_POSITION`
    :   A string containing a decimal integer indicating an error cursor position as an index into the original statement string. The first character has index 1, and positions are measured in characters not bytes.
    <a id="libpq-pg-diag-internal-position"></a>

    `PG_DIAG_INTERNAL_POSITION`
    :   This is defined the same as the `PG_DIAG_STATEMENT_POSITION` field, but it is used when the cursor position refers to an internally generated command rather than the one submitted by the client. The `PG_DIAG_INTERNAL_QUERY` field will always appear when this field appears.
    <a id="libpq-pg-diag-internal-query"></a>

    `PG_DIAG_INTERNAL_QUERY`
    :   The text of a failed internally-generated command. This could be, for example, an SQL query issued by a PL/pgSQL function.
    <a id="libpq-pg-diag-context"></a>

    `PG_DIAG_CONTEXT`
    :   An indication of the context in which the error occurred. Presently this includes a call stack traceback of active procedural language functions and internally-generated queries. The trace is one entry per line, most recent first.
    <a id="libpq-pg-diag-schema-name"></a>

    `PG_DIAG_SCHEMA_NAME`
    :   If the error was associated with a specific database object, the name of the schema containing that object, if any.
    <a id="libpq-pg-diag-table-name"></a>

    `PG_DIAG_TABLE_NAME`
    :   If the error was associated with a specific table, the name of the table. (Refer to the schema name field for the name of the table's schema.)
    <a id="libpq-pg-diag-column-name"></a>

    `PG_DIAG_COLUMN_NAME`
    :   If the error was associated with a specific table column, the name of the column. (Refer to the schema and table name fields to identify the table.)
    <a id="libpq-pg-diag-datatype-name"></a>

    `PG_DIAG_DATATYPE_NAME`
    :   If the error was associated with a specific data type, the name of the data type. (Refer to the schema name field for the name of the data type's schema.)
    <a id="libpq-pg-diag-constraint-name"></a>

    `PG_DIAG_CONSTRAINT_NAME`
    :   If the error was associated with a specific constraint, the name of the constraint. Refer to fields listed above for the associated table or domain. (For this purpose, indexes are treated as constraints, even if they weren't created with constraint syntax.)
    <a id="libpq-pg-diag-source-file"></a>

    `PG_DIAG_SOURCE_FILE`
    :   The file name of the source-code location where the error was reported.
    <a id="libpq-pg-diag-source-line"></a>

    `PG_DIAG_SOURCE_LINE`
    :   The line number of the source-code location where the error was reported.
    <a id="libpq-pg-diag-source-function"></a>

    `PG_DIAG_SOURCE_FUNCTION`
    :   The name of the source-code function reporting the error.


    !!! note

        The fields for schema name, table name, column name, data type name, and constraint name are supplied only for a limited number of error types; see [PostgreSQL Error Codes](../../appendixes/postgresql-error-codes.md#errcodes-appendix). Do not assume that the presence of any of these fields guarantees the presence of another field. Core error sources observe the interrelationships noted above, but user-defined functions may use these fields in other ways. In the same vein, do not assume that these fields denote contemporary objects in the current database.


     The client is responsible for formatting displayed information to meet its needs; in particular it should break long lines as needed. Newline characters appearing in the error message fields should be treated as paragraph breaks, not line breaks.


     Errors generated internally by libpq will have severity and primary message, but typically no other fields.


     Note that error fields are only available from `PGresult` objects, not `PGconn` objects; there is no `PQerrorField` function.
<a id="libpq-PQclear"></a>

`PQclear`
:   Frees the storage associated with a `PGresult`. Every command result should be freed via [PQclear](#libpq-PQclear) when it is no longer needed.

    ```

    void PQclear(PGresult *res);
    ```
     If the argument is a `NULL` pointer, no operation is performed.


     You can keep a `PGresult` object around for as long as you need it; it does not go away when you issue a new command, nor even if you close the connection. To get rid of it, you must call [PQclear](#libpq-PQclear). Failure to do this will result in memory leaks in your application.

  <a id="libpq-exec-select-info"></a>

### Retrieving Query Result Information


 These functions are used to extract information from a `PGresult` object that represents a successful query result (that is, one that has status `PGRES_TUPLES_OK`, `PGRES_SINGLE_TUPLE`, or `PGRES_TUPLES_CHUNK`). They can also be used to extract information from a successful Describe operation: a Describe's result has all the same column information that actual execution of the query would provide, but it has zero rows. For objects with other status values, these functions will act as though the result has zero rows and zero columns.


<a id="libpq-PQntuples"></a>

`PQntuples`
:   Returns the number of rows (tuples) in the query result. (Note that `PGresult` objects are limited to no more than `INT_MAX` rows, so an `int` result is sufficient.)

    ```

    int PQntuples(const PGresult *res);
    ```
<a id="libpq-PQnfields"></a>

`PQnfields`
:   Returns the number of columns (fields) in each row of the query result.

    ```

    int PQnfields(const PGresult *res);
    ```
<a id="libpq-PQfname"></a>

`PQfname`
:   Returns the column name associated with the given column number. Column numbers start at 0. The caller should not free the result directly. It will be freed when the associated `PGresult` handle is passed to [PQclear](#libpq-PQclear).

    ```

    char *PQfname(const PGresult *res,
                  int column_number);
    ```


     `NULL` is returned if the column number is out of range.
<a id="libpq-PQfnumber"></a>

`PQfnumber`
:   Returns the column number associated with the given column name.

    ```

    int PQfnumber(const PGresult *res,
                  const char *column_name);
    ```


     -1 is returned if the given name does not match any column.


     The given name is treated like an identifier in an SQL command, that is, it is downcased unless double-quoted. For example, given a query result generated from the SQL command:

    ```sql

    SELECT 1 AS FOO, 2 AS "BAR";
    ```
     we would have the results:

    ```

    PQfname(res, 0)              foo
    PQfname(res, 1)              BAR
    PQfnumber(res, "FOO")        0
    PQfnumber(res, "foo")        0
    PQfnumber(res, "BAR")        -1
    PQfnumber(res, "\"BAR\"")    1
    ```
<a id="libpq-PQftable"></a>

`PQftable`
:   Returns the OID of the table from which the given column was fetched. Column numbers start at 0.

    ```

    Oid PQftable(const PGresult *res,
                 int column_number);
    ```


     `InvalidOid` is returned if the column number is out of range, or if the specified column is not a simple reference to a table column. You can query the system table `pg_class` to determine exactly which table is referenced.


     The type `Oid` and the constant `InvalidOid` will be defined when you include the libpq header file. They will both be some integer type.
<a id="libpq-PQftablecol"></a>

`PQftablecol`
:   Returns the column number (within its table) of the column making up the specified query result column. Query-result column numbers start at 0, but table columns have nonzero numbers.

    ```

    int PQftablecol(const PGresult *res,
                    int column_number);
    ```


     Zero is returned if the column number is out of range, or if the specified column is not a simple reference to a table column.
<a id="libpq-PQfformat"></a>

`PQfformat`
:   Returns the format code indicating the format of the given column. Column numbers start at 0.

    ```

    int PQfformat(const PGresult *res,
                  int column_number);
    ```


     Format code zero indicates textual data representation, while format code one indicates binary representation. (Other codes are reserved for future definition.)
<a id="libpq-PQftype"></a>

`PQftype`
:   Returns the data type associated with the given column number. The integer returned is the internal OID number of the type. Column numbers start at 0.

    ```

    Oid PQftype(const PGresult *res,
                int column_number);
    ```


     You can query the system table `pg_type` to obtain the names and properties of the various data types. The OIDs of the built-in data types are defined in the file `catalog/pg_type_d.h` in the PostgreSQL installation's `include` directory.
<a id="libpq-PQfmod"></a>

`PQfmod`
:   Returns the type modifier of the column associated with the given column number. Column numbers start at 0.

    ```

    int PQfmod(const PGresult *res,
               int column_number);
    ```


     The interpretation of modifier values is type-specific; they typically indicate precision or size limits. The value -1 is used to indicate “no information available”. Most data types do not use modifiers, in which case the value is always -1.
<a id="libpq-PQfsize"></a>

`PQfsize`
:   Returns the size in bytes of the column associated with the given column number. Column numbers start at 0.

    ```

    int PQfsize(const PGresult *res,
                int column_number);
    ```


     [PQfsize](#libpq-PQfsize) returns the space allocated for this column in a database row, in other words the size of the server's internal representation of the data type. (Accordingly, it is not really very useful to clients.) A negative value indicates the data type is variable-length.
<a id="libpq-PQbinaryTuples"></a>

`PQbinaryTuples`
:   Returns 1 if the `PGresult` contains binary data and 0 if it contains text data.

    ```

    int PQbinaryTuples(const PGresult *res);
    ```


     This function is deprecated (except for its use in connection with `COPY`), because it is possible for a single `PGresult` to contain text data in some columns and binary data in others. [PQfformat](#libpq-PQfformat) is preferred. [PQbinaryTuples](#libpq-PQbinaryTuples) returns 1 only if all columns of the result are binary (format 1).
<a id="libpq-PQgetvalue"></a>

`PQgetvalue`
:   Returns a single field value of one row of a `PGresult`. Row and column numbers start at 0. The caller should not free the result directly. It will be freed when the associated `PGresult` handle is passed to [PQclear](#libpq-PQclear).

    ```

    char *PQgetvalue(const PGresult *res,
                     int row_number,
                     int column_number);
    ```


     For data in text format, the value returned by [PQgetvalue](#libpq-PQgetvalue) is a null-terminated character string representation of the field value. For data in binary format, the value is in the binary representation determined by the data type's `typsend` and `typreceive` functions. (The value is actually followed by a zero byte in this case too, but that is not ordinarily useful, since the value is likely to contain embedded nulls.)


     An empty string is returned if the field value is null. See [PQgetisnull](#libpq-PQgetisnull) to distinguish null values from empty-string values.


     The pointer returned by [PQgetvalue](#libpq-PQgetvalue) points to storage that is part of the `PGresult` structure. One should not modify the data it points to, and one must explicitly copy the data into other storage if it is to be used past the lifetime of the `PGresult` structure itself.
<a id="libpq-PQgetisnull"></a>

`PQgetisnull`
:   Tests a field for a null value. Row and column numbers start at 0.

    ```

    int PQgetisnull(const PGresult *res,
                    int row_number,
                    int column_number);
    ```


     This function returns 1 if the field is null and 0 if it contains a non-null value. (Note that [PQgetvalue](#libpq-PQgetvalue) will return an empty string, not a null pointer, for a null field.)
<a id="libpq-PQgetlength"></a>

`PQgetlength`
:   Returns the actual length of a field value in bytes. Row and column numbers start at 0.

    ```

    int PQgetlength(const PGresult *res,
                    int row_number,
                    int column_number);
    ```


     This is the actual data length for the particular data value, that is, the size of the object pointed to by [PQgetvalue](#libpq-PQgetvalue). For text data format this is the same as `strlen()`. For binary format this is essential information. Note that one should *not* rely on [PQfsize](#libpq-PQfsize) to obtain the actual data length.
<a id="libpq-PQnparams"></a>

`PQnparams`
:   Returns the number of parameters of a prepared statement.

    ```

    int PQnparams(const PGresult *res);
    ```


     This function is only useful when inspecting the result of [PQdescribePrepared](#libpq-PQdescribePrepared). For other types of results it will return zero.
<a id="libpq-PQparamtype"></a>

`PQparamtype`
:   Returns the data type of the indicated statement parameter. Parameter numbers start at 0.

    ```

    Oid PQparamtype(const PGresult *res, int param_number);
    ```


     This function is only useful when inspecting the result of [PQdescribePrepared](#libpq-PQdescribePrepared). For other types of results it will return zero.
<a id="libpq-PQprint"></a>

`PQprint`
:   Prints out all the rows and, optionally, the column names to the specified output stream.

    ```

    void PQprint(FILE *fout,      /* output stream */
                 const PGresult *res,
                 const PQprintOpt *po);
    typedef struct
    {
        pqbool  header;      /* print output field headings and row count */
        pqbool  align;       /* fill align the fields */
        pqbool  standard;    /* old brain dead format */
        pqbool  html3;       /* output HTML tables */
        pqbool  expanded;    /* expand tables */
        pqbool  pager;       /* use pager for output if needed */
        char    *fieldSep;   /* field separator */
        char    *tableOpt;   /* attributes for HTML table element */
        char    *caption;    /* HTML table caption */
        char    **fieldName; /* null-terminated array of replacement field names */
    } PQprintOpt;
    ```


     This function was formerly used by psql to print query results, but this is no longer the case. Note that it assumes all the data is in text format.
  <a id="libpq-exec-nonselect"></a>

### Retrieving Other Result Information


 These functions are used to extract other information from `PGresult` objects.


<a id="libpq-PQcmdStatus"></a>

`PQcmdStatus`
:   Returns the command status tag from the SQL command that generated the `PGresult`.

    ```

    char *PQcmdStatus(PGresult *res);
    ```


     Commonly this is just the name of the command, but it might include additional data such as the number of rows processed. The caller should not free the result directly. It will be freed when the associated `PGresult` handle is passed to [PQclear](#libpq-PQclear).
<a id="libpq-PQcmdTuples"></a>

`PQcmdTuples`
:   Returns the number of rows affected by the SQL command.

    ```

    char *PQcmdTuples(PGresult *res);
    ```


     This function returns a string containing the number of rows affected by the SQL statement that generated the `PGresult`. This function can only be used following the execution of a `SELECT`, `CREATE TABLE AS`, `INSERT`, `UPDATE`, `DELETE`, `MERGE`, `MOVE`, `FETCH`, or `COPY` statement, or an `EXECUTE` of a prepared query that contains an `INSERT`, `UPDATE`, `DELETE`, or `MERGE` statement. If the command that generated the `PGresult` was anything else, [PQcmdTuples](#libpq-PQcmdTuples) returns an empty string. The caller should not free the return value directly. It will be freed when the associated `PGresult` handle is passed to [PQclear](#libpq-PQclear).
<a id="libpq-PQoidValue"></a>

`PQoidValue`
:   Returns the OID of the inserted row, if the SQL command was an `INSERT` that inserted exactly one row into a table that has OIDs, or a `EXECUTE` of a prepared query containing a suitable `INSERT` statement. Otherwise, this function returns `InvalidOid`. This function will also return `InvalidOid` if the table affected by the `INSERT` statement does not contain OIDs.

    ```

    Oid PQoidValue(const PGresult *res);
    ```
<a id="libpq-PQoidStatus"></a>

`PQoidStatus`
:   This function is deprecated in favor of [PQoidValue](#libpq-PQoidValue) and is not thread-safe. It returns a string with the OID of the inserted row, while [PQoidValue](#libpq-PQoidValue) returns the OID value.

    ```

    char *PQoidStatus(const PGresult *res);
    ```
  <a id="libpq-exec-escape-string"></a>

### Escaping Strings for Inclusion in SQL Commands


<a id="libpq-PQescapeLiteral"></a>

`PQescapeLiteral`
:   ```

    char *PQescapeLiteral(PGconn *conn, const char *str, size_t length);
    ```


     [PQescapeLiteral](#libpq-PQescapeLiteral) escapes a string for use within an SQL command. This is useful when inserting data values as literal constants in SQL commands. Certain characters (such as quotes and backslashes) must be escaped to prevent them from being interpreted specially by the SQL parser. [PQescapeLiteral](#libpq-PQescapeLiteral) performs this operation.


     [PQescapeLiteral](#libpq-PQescapeLiteral) returns an escaped version of the `str` parameter in memory allocated with `malloc()`. This memory should be freed using `PQfreemem()` when the result is no longer needed. A terminating zero byte is not required, and should not be counted in `length`. (If a terminating zero byte is found before `length` bytes are processed, [PQescapeLiteral](#libpq-PQescapeLiteral) stops at the zero; the behavior is thus rather like `strncpy`.) The return string has all special characters replaced so that they can be properly processed by the PostgreSQL string literal parser. A terminating zero byte is also added. The single quotes that must surround PostgreSQL string literals are included in the result string.


     On error, [PQescapeLiteral](#libpq-PQescapeLiteral) returns `NULL` and a suitable message is stored in the `conn` object.


    !!! tip

        It is especially important to do proper escaping when handling strings that were received from an untrustworthy source. Otherwise there is a security risk: you are vulnerable to “SQL injection” attacks wherein unwanted SQL commands are fed to your database.


     Note that it is neither necessary nor correct to do escaping when a data value is passed as a separate parameter in [PQexecParams](#libpq-PQexecParams) or its sibling routines.
<a id="libpq-PQescapeIdentifier"></a>

`PQescapeIdentifier`
:   ```

    char *PQescapeIdentifier(PGconn *conn, const char *str, size_t length);
    ```


     [PQescapeIdentifier](#libpq-PQescapeIdentifier) escapes a string for use as an SQL identifier, such as a table, column, or function name. This is useful when a user-supplied identifier might contain special characters that would otherwise not be interpreted as part of the identifier by the SQL parser, or when the identifier might contain upper case characters whose case should be preserved.


     [PQescapeIdentifier](#libpq-PQescapeIdentifier) returns a version of the `str` parameter escaped as an SQL identifier in memory allocated with `malloc()`. This memory must be freed using `PQfreemem()` when the result is no longer needed. A terminating zero byte is not required, and should not be counted in `length`. (If a terminating zero byte is found before `length` bytes are processed, [PQescapeIdentifier](#libpq-PQescapeIdentifier) stops at the zero; the behavior is thus rather like `strncpy`.) The return string has all special characters replaced so that it will be properly processed as an SQL identifier. A terminating zero byte is also added. The return string will also be surrounded by double quotes.


     On error, [PQescapeIdentifier](#libpq-PQescapeIdentifier) returns `NULL` and a suitable message is stored in the `conn` object.


    !!! tip

        As with string literals, to prevent SQL injection attacks, SQL identifiers must be escaped when they are received from an untrustworthy source.
<a id="libpq-PQescapeStringConn"></a>

`PQescapeStringConn`
:   ```

    size_t PQescapeStringConn(PGconn *conn,
                              char *to, const char *from, size_t length,
                              int *error);
    ```


     [PQescapeStringConn](#libpq-PQescapeStringConn) escapes string literals, much like [PQescapeLiteral](#libpq-PQescapeLiteral). Unlike [PQescapeLiteral](#libpq-PQescapeLiteral), the caller is responsible for providing an appropriately sized buffer. Furthermore, [PQescapeStringConn](#libpq-PQescapeStringConn) does not generate the single quotes that must surround PostgreSQL string literals; they should be provided in the SQL command that the result is inserted into. The parameter `from` points to the first character of the string that is to be escaped, and the `length` parameter gives the number of bytes in this string. A terminating zero byte is not required, and should not be counted in `length`. (If a terminating zero byte is found before `length` bytes are processed, [PQescapeStringConn](#libpq-PQescapeStringConn) stops at the zero; the behavior is thus rather like `strncpy`.) `to` shall point to a buffer that is able to hold at least one more byte than twice the value of `length`, otherwise the behavior is undefined. Behavior is likewise undefined if the `to` and `from` strings overlap.


     If the `error` parameter is not `NULL`, then `*error` is set to zero on success, nonzero on error. Presently the only possible error conditions involve invalid multibyte encoding in the source string. The output string is still generated on error, but it can be expected that the server will reject it as malformed. On error, a suitable message is stored in the `conn` object, whether or not `error` is `NULL`.


     [PQescapeStringConn](#libpq-PQescapeStringConn) returns the number of bytes written to `to`, not including the terminating zero byte.
<a id="libpq-PQescapeString"></a>

`PQescapeString`
:   [PQescapeString](#libpq-PQescapeString) is an older, deprecated version of [PQescapeStringConn](#libpq-PQescapeStringConn).

    ```

    size_t PQescapeString (char *to, const char *from, size_t length);
    ```


     The only difference from [PQescapeStringConn](#libpq-PQescapeStringConn) is that [PQescapeString](#libpq-PQescapeString) does not take `PGconn` or `error` parameters. Because of this, it cannot adjust its behavior depending on the connection properties (such as character encoding) and therefore *it might give the wrong results*. Also, it has no way to report error conditions.


     [PQescapeString](#libpq-PQescapeString) can be used safely in client programs that work with only one PostgreSQL connection at a time (in this case it can find out what it needs to know “behind the scenes”). In other contexts it is a security hazard and should be avoided in favor of [PQescapeStringConn](#libpq-PQescapeStringConn).
<a id="libpq-PQescapeByteaConn"></a>

`PQescapeByteaConn`
:   Escapes binary data for use within an SQL command with the type `bytea`. As with [PQescapeStringConn](#libpq-PQescapeStringConn), this is only used when inserting data directly into an SQL command string.

    ```

    unsigned char *PQescapeByteaConn(PGconn *conn,
                                     const unsigned char *from,
                                     size_t from_length,
                                     size_t *to_length);
    ```


     Certain byte values must be escaped when used as part of a `bytea` literal in an SQL statement. [PQescapeByteaConn](#libpq-PQescapeByteaConn) escapes bytes using either hex encoding or backslash escaping. See [Binary Data Types](../../the-sql-language/data-types/binary-data-types.md#datatype-binary) for more information.


     The `from` parameter points to the first byte of the string that is to be escaped, and the `from_length` parameter gives the number of bytes in this binary string. (A terminating zero byte is neither necessary nor counted.) The `to_length` parameter points to a variable that will hold the resultant escaped string length. This result string length includes the terminating zero byte of the result.


     [PQescapeByteaConn](#libpq-PQescapeByteaConn) returns an escaped version of the `from` parameter binary string in memory allocated with `malloc()`. This memory should be freed using `PQfreemem()` when the result is no longer needed. The return string has all special characters replaced so that they can be properly processed by the PostgreSQL string literal parser, and the `bytea` input function. A terminating zero byte is also added. The single quotes that must surround PostgreSQL string literals are not part of the result string.


     On error, a null pointer is returned, and a suitable error message is stored in the `conn` object. Currently, the only possible error is insufficient memory for the result string.
<a id="libpq-PQescapeBytea"></a>

`PQescapeBytea`
:   [PQescapeBytea](#libpq-PQescapeBytea) is an older, deprecated version of [PQescapeByteaConn](#libpq-PQescapeByteaConn).

    ```

    unsigned char *PQescapeBytea(const unsigned char *from,
                                 size_t from_length,
                                 size_t *to_length);
    ```


     The only difference from [PQescapeByteaConn](#libpq-PQescapeByteaConn) is that [PQescapeBytea](#libpq-PQescapeBytea) does not take a `PGconn` parameter. Because of this, [PQescapeBytea](#libpq-PQescapeBytea) can only be used safely in client programs that use a single PostgreSQL connection at a time (in this case it can find out what it needs to know “behind the scenes”). It *might give the wrong results* if used in programs that use multiple database connections (use [PQescapeByteaConn](#libpq-PQescapeByteaConn) in such cases).
<a id="libpq-PQunescapeBytea"></a>

`PQunescapeBytea`
:   Converts a string representation of binary data into binary data — the reverse of [PQescapeBytea](#libpq-PQescapeBytea). This is needed when retrieving `bytea` data in text format, but not when retrieving it in binary format.

    ```

    unsigned char *PQunescapeBytea(const unsigned char *from, size_t *to_length);
    ```


     The `from` parameter points to a string such as might be returned by [PQgetvalue](#libpq-PQgetvalue) when applied to a `bytea` column. [PQunescapeBytea](#libpq-PQunescapeBytea) converts this string representation into its binary representation. It returns a pointer to a buffer allocated with `malloc()`, or `NULL` on error, and puts the size of the buffer in `to_length`. The result must be freed using [PQfreemem](miscellaneous-functions.md#libpq-PQfreemem) when it is no longer needed.


     This conversion is not exactly the inverse of [PQescapeBytea](#libpq-PQescapeBytea), because the string is not expected to be “escaped” when received from [PQgetvalue](#libpq-PQgetvalue). In particular this means there is no need for string quoting considerations, and so no need for a `PGconn` parameter.
