<a id="ecpg-sql-commands"></a>

## Embedded SQL Commands


 This section describes all SQL commands that are specific to embedded SQL. Also refer to the SQL commands listed in [SQL Commands](../../reference/sql-commands/index.md#sql-commands), which can also be used in embedded SQL, unless stated otherwise.
 <a id="ecpg-sql-allocate-descriptor"></a>

# ALLOCATE DESCRIPTOR

allocate an SQL descriptor area

## Synopsis


```

ALLOCATE DESCRIPTOR NAME
```


## Description


 `ALLOCATE DESCRIPTOR` allocates a new named SQL descriptor area, which can be used to exchange data between the PostgreSQL server and the host program.


 Descriptor areas should be freed after use using the `DEALLOCATE DESCRIPTOR` command.


## Parameters


<a id="ecpg-sql-allocate-descriptor-name"></a>

*name*
:   A name of SQL descriptor, case sensitive. This can be an SQL identifier or a host variable.


## Examples


```

EXEC SQL ALLOCATE DESCRIPTOR mydesc;
```


## Compatibility


 `ALLOCATE DESCRIPTOR` is specified in the SQL standard.


## See Also
  [ecpg-sql-deallocate-descriptor](#ecpg-sql-deallocate-descriptor), [ecpg-sql-get-descriptor](#ecpg-sql-get-descriptor), [ecpg-sql-set-descriptor](#ecpg-sql-set-descriptor)  <a id="ecpg-sql-connect"></a>

# CONNECT

establish a database connection

## Synopsis


```

CONNECT TO CONNECTION_TARGET [ AS CONNECTION_NAME ] [ USER CONNECTION_USER ]
CONNECT TO DEFAULT
CONNECT CONNECTION_USER
DATABASE CONNECTION_TARGET
```


## Description


 The `CONNECT` command establishes a connection between the client and the PostgreSQL server.


## Parameters


<a id="ecpg-sql-connect-connection-target"></a>

*connection_target*
:   *connection_target* specifies the target server of the connection on one of several forms.

    <a id="ecpg-sql-connect-connection-target-database-name"></a>

    [ *database_name* ] [ `@`*host* ] [ `:`*port* ]
    :   Connect over TCP/IP
    <a id="ecpg-sql-connect-connection-target-unix-domain-sockets"></a>

    `unix:postgresql://`*host* [ `:`*port* ] `/` [ *database_name* ] [ `?`*connection_option* ]
    :   Connect over Unix-domain sockets
    <a id="ecpg-sql-connect-connection-target-tcp-ip"></a>

    `tcp:postgresql://`*host* [ `:`*port* ] `/` [ *database_name* ] [ `?`*connection_option* ]
    :   Connect over TCP/IP
    <a id="ecpg-sql-connect-connection-target-constant"></a>

    SQL string constant
    :   containing a value in one of the above forms
    <a id="ecpg-sql-connect-connection-target-host-variable"></a>

    host variable
    :   host variable of type `char[]` or `VARCHAR[]` containing a value in one of the above forms
<a id="ecpg-sql-connect-connection-name"></a>

*connection_name*
:   An optional identifier for the connection, so that it can be referred to in other commands. This can be an SQL identifier or a host variable.
<a id="ecpg-sql-connect-connection-user"></a>

*connection_user*
:   The user name for the database connection.


     This parameter can also specify user name and password, using one the forms <em>user_name</em><code>/</code><em>password</em>, <em>user_name</em><code> IDENTIFIED BY </code><em>password</em>, or <em>user_name</em><code> USING </code><em>password</em>.


     User name and password can be SQL identifiers, string constants, or host variables.
<a id="ecpg-sql-connect-default"></a>

`DEFAULT`
:   Use all default connection parameters, as defined by libpq.


## Examples


 Here a several variants for specifying connection parameters:

```

EXEC SQL CONNECT TO "connectdb" AS main;
EXEC SQL CONNECT TO "connectdb" AS second;
EXEC SQL CONNECT TO "unix:postgresql://200.46.204.71/connectdb" AS main USER connectuser;
EXEC SQL CONNECT TO "unix:postgresql://localhost/connectdb" AS main USER connectuser;
EXEC SQL CONNECT TO 'connectdb' AS main;
EXEC SQL CONNECT TO 'unix:postgresql://localhost/connectdb' AS main USER :user;
EXEC SQL CONNECT TO :db AS :id;
EXEC SQL CONNECT TO :db USER connectuser USING :pw;
EXEC SQL CONNECT TO @localhost AS main USER connectdb;
EXEC SQL CONNECT TO REGRESSDB1 as main;
EXEC SQL CONNECT TO AS main USER connectdb;
EXEC SQL CONNECT TO connectdb AS :id;
EXEC SQL CONNECT TO connectdb AS main USER connectuser/connectdb;
EXEC SQL CONNECT TO connectdb AS main;
EXEC SQL CONNECT TO connectdb@localhost AS main;
EXEC SQL CONNECT TO tcp:postgresql://localhost/ USER connectdb;
EXEC SQL CONNECT TO tcp:postgresql://localhost/connectdb USER connectuser IDENTIFIED BY connectpw;
EXEC SQL CONNECT TO tcp:postgresql://localhost:20/connectdb USER connectuser IDENTIFIED BY connectpw;
EXEC SQL CONNECT TO unix:postgresql://localhost/ AS main USER connectdb;
EXEC SQL CONNECT TO unix:postgresql://localhost/connectdb AS main USER connectuser;
EXEC SQL CONNECT TO unix:postgresql://localhost/connectdb USER connectuser IDENTIFIED BY "connectpw";
EXEC SQL CONNECT TO unix:postgresql://localhost/connectdb USER connectuser USING "connectpw";
EXEC SQL CONNECT TO unix:postgresql://localhost/connectdb?connect_timeout=14 USER connectuser;
```


 Here is an example program that illustrates the use of host variables to specify connection parameters:

```

int
main(void)
{
EXEC SQL BEGIN DECLARE SECTION;
    char *dbname     = "testdb";    /* database name */
    char *user       = "testuser";  /* connection user name */
    char *connection = "tcp:postgresql://localhost:5432/testdb";
                                    /* connection string */
    char ver[256];                  /* buffer to store the version string */
EXEC SQL END DECLARE SECTION;

    ECPGdebug(1, stderr);

    EXEC SQL CONNECT TO :dbname USER :user;
    EXEC SQL SELECT pg_catalog.set_config('search_path', '', false); EXEC SQL COMMIT;
    EXEC SQL SELECT version() INTO :ver;
    EXEC SQL DISCONNECT;

    printf("version: %s\n", ver);

    EXEC SQL CONNECT TO :connection USER :user;
    EXEC SQL SELECT pg_catalog.set_config('search_path', '', false); EXEC SQL COMMIT;
    EXEC SQL SELECT version() INTO :ver;
    EXEC SQL DISCONNECT;

    printf("version: %s\n", ver);

    return 0;
}
```


## Compatibility


 `CONNECT` is specified in the SQL standard, but the format of the connection parameters is implementation-specific.


## See Also
  [ecpg-sql-disconnect](#ecpg-sql-disconnect), [ecpg-sql-set-connection](#ecpg-sql-set-connection)  <a id="ecpg-sql-deallocate-descriptor"></a>

# DEALLOCATE DESCRIPTOR

deallocate an SQL descriptor area

## Synopsis


```

DEALLOCATE DESCRIPTOR NAME
```


## Description


 `DEALLOCATE DESCRIPTOR` deallocates a named SQL descriptor area.


## Parameters


<a id="ecpg-sql-deallocate-descriptor-name"></a>

*name*
:   The name of the descriptor which is going to be deallocated. It is case sensitive. This can be an SQL identifier or a host variable.


## Examples


```

EXEC SQL DEALLOCATE DESCRIPTOR mydesc;
```


## Compatibility


 `DEALLOCATE DESCRIPTOR` is specified in the SQL standard.


## See Also
  [ecpg-sql-allocate-descriptor](#ecpg-sql-allocate-descriptor), [ecpg-sql-get-descriptor](#ecpg-sql-get-descriptor), [ecpg-sql-set-descriptor](#ecpg-sql-set-descriptor)  <a id="ecpg-sql-declare"></a>

# DECLARE

define a cursor

## Synopsis


```

DECLARE CURSOR_NAME [ BINARY ] [ ASENSITIVE | INSENSITIVE ] [ [ NO ] SCROLL ] CURSOR [ { WITH | WITHOUT } HOLD ] FOR PREPARED_NAME
DECLARE CURSOR_NAME [ BINARY ] [ ASENSITIVE | INSENSITIVE ] [ [ NO ] SCROLL ] CURSOR [ { WITH | WITHOUT } HOLD ] FOR QUERY
```


## Description


 `DECLARE` declares a cursor for iterating over the result set of a prepared statement. This command has slightly different semantics from the direct SQL command `DECLARE`: Whereas the latter executes a query and prepares the result set for retrieval, this embedded SQL command merely declares a name as a “loop variable” for iterating over the result set of a query; the actual execution happens when the cursor is opened with the `OPEN` command.


## Parameters


<a id="ecpg-sql-declare-cursor-name"></a>

*cursor_name*
:   A cursor name, case sensitive. This can be an SQL identifier or a host variable.
<a id="ecpg-sql-declare-prepared-name"></a>

*prepared_name*
:   The name of a prepared query, either as an SQL identifier or a host variable.
<a id="ecpg-sql-declare-query"></a>

*query*
:   A [sql-select](../../reference/sql-commands/select.md#sql-select) or [sql-values](../../reference/sql-commands/values.md#sql-values) command which will provide the rows to be returned by the cursor.


 For the meaning of the cursor options, see [sql-declare](../../reference/sql-commands/declare.md#sql-declare).


## Examples


 Examples declaring a cursor for a query:

```

EXEC SQL DECLARE C CURSOR FOR SELECT * FROM My_Table;
EXEC SQL DECLARE C CURSOR FOR SELECT Item1 FROM T;
EXEC SQL DECLARE cur1 CURSOR FOR SELECT version();
```


 An example declaring a cursor for a prepared statement:

```

EXEC SQL PREPARE stmt1 AS SELECT version();
EXEC SQL DECLARE cur1 CURSOR FOR stmt1;
```


## Compatibility


 `DECLARE` is specified in the SQL standard.


## See Also
  [ecpg-sql-open](#ecpg-sql-open), [sql-close](../../reference/sql-commands/close.md#sql-close), [sql-declare](../../reference/sql-commands/declare.md#sql-declare)  <a id="ecpg-sql-declare-statement"></a>

# DECLARE STATEMENT

declare SQL statement identifier

## Synopsis


```

EXEC SQL [ AT CONNECTION_NAME ] DECLARE STATEMENT_NAME STATEMENT
```


## Description


 `DECLARE STATEMENT` declares an SQL statement identifier. SQL statement identifier can be associated with the connection. When the identifier is used by dynamic SQL statements, the statements are executed using the associated connection. The namespace of the declaration is the precompile unit, and multiple declarations to the same SQL statement identifier are not allowed. Note that if the precompiler runs in Informix compatibility mode and some SQL statement is declared, "database" can not be used as a cursor name.


## Parameters


<a id="ecpg-sql-declare-statement-connection-name"></a>

*connection_name*
:   A database connection name established by the `CONNECT` command.


     AT clause can be omitted, but such statement has no meaning.


<a id="ecpg-sql-declare-statement-statement-name"></a>

*statement_name*
:   The name of an SQL statement identifier, either as an SQL identifier or a host variable.


## Notes


 This association is valid only if the declaration is physically placed on top of a dynamic statement.


## Examples


```

EXEC SQL CONNECT TO postgres AS con1;
EXEC SQL AT con1 DECLARE sql_stmt STATEMENT;
EXEC SQL DECLARE cursor_name CURSOR FOR sql_stmt;
EXEC SQL PREPARE sql_stmt FROM :dyn_string;
EXEC SQL OPEN cursor_name;
EXEC SQL FETCH cursor_name INTO :column1;
EXEC SQL CLOSE cursor_name;
```


## Compatibility


 `DECLARE STATEMENT` is an extension of the SQL standard, but can be used in famous DBMSs.


## See Also
  [ecpg-sql-connect](#ecpg-sql-connect), [ecpg-sql-declare](#ecpg-sql-declare), [ecpg-sql-open](#ecpg-sql-open)  <a id="ecpg-sql-describe"></a>

# DESCRIBE

obtain information about a prepared statement or result set

## Synopsis


```

DESCRIBE [ OUTPUT ] PREPARED_NAME USING [ SQL ] DESCRIPTOR DESCRIPTOR_NAME
DESCRIBE [ OUTPUT ] PREPARED_NAME INTO [ SQL ] DESCRIPTOR DESCRIPTOR_NAME
DESCRIBE [ OUTPUT ] PREPARED_NAME INTO SQLDA_NAME
```


## Description


 `DESCRIBE` retrieves metadata information about the result columns contained in a prepared statement, without actually fetching a row.


## Parameters


<a id="ecpg-sql-describe-prepared-name"></a>

*prepared_name*
:   The name of a prepared statement. This can be an SQL identifier or a host variable.
<a id="ecpg-sql-describe-descriptor-name"></a>

*descriptor_name*
:   A descriptor name. It is case sensitive. It can be an SQL identifier or a host variable.
<a id="ecpg-sql-describe-sqlda-name"></a>

*sqlda_name*
:   The name of an SQLDA variable.


## Examples


```

EXEC SQL ALLOCATE DESCRIPTOR mydesc;
EXEC SQL PREPARE stmt1 FROM :sql_stmt;
EXEC SQL DESCRIBE stmt1 INTO SQL DESCRIPTOR mydesc;
EXEC SQL GET DESCRIPTOR mydesc VALUE 1 :charvar = NAME;
EXEC SQL DEALLOCATE DESCRIPTOR mydesc;
```


## Compatibility


 `DESCRIBE` is specified in the SQL standard.


## See Also
  [ecpg-sql-allocate-descriptor](#ecpg-sql-allocate-descriptor), [ecpg-sql-get-descriptor](#ecpg-sql-get-descriptor)  <a id="ecpg-sql-disconnect"></a>

# DISCONNECT

terminate a database connection

## Synopsis


```

DISCONNECT CONNECTION_NAME
DISCONNECT [ CURRENT ]
DISCONNECT ALL
```


## Description


 `DISCONNECT` closes a connection (or all connections) to the database.


## Parameters


<a id="ecpg-sql-disconnect-connection-name"></a>

*connection_name*
:   A database connection name established by the `CONNECT` command.
<a id="ecpg-sql-disconnect-current"></a>

`CURRENT`
:   Close the “current” connection, which is either the most recently opened connection, or the connection set by the `SET CONNECTION` command. This is also the default if no argument is given to the `DISCONNECT` command.
<a id="ecpg-sql-disconnect-all"></a>

`ALL`
:   Close all open connections.


## Examples


```

int
main(void)
{
    EXEC SQL CONNECT TO testdb AS con1 USER testuser;
    EXEC SQL CONNECT TO testdb AS con2 USER testuser;
    EXEC SQL CONNECT TO testdb AS con3 USER testuser;

    EXEC SQL DISCONNECT CURRENT;  /* close con3          */
    EXEC SQL DISCONNECT ALL;      /* close con2 and con1 */

    return 0;
}
```


## Compatibility


 `DISCONNECT` is specified in the SQL standard.


## See Also
  [ecpg-sql-connect](#ecpg-sql-connect), [ecpg-sql-set-connection](#ecpg-sql-set-connection)  <a id="ecpg-sql-execute-immediate"></a>

# EXECUTE IMMEDIATE

dynamically prepare and execute a statement

## Synopsis


```

EXECUTE IMMEDIATE STRING
```


## Description


 `EXECUTE IMMEDIATE` immediately prepares and executes a dynamically specified SQL statement, without retrieving result rows.


## Parameters


<a id="ecpg-sql-execute-immediate-string"></a>

*string*
:   A literal string or a host variable containing the SQL statement to be executed.


## Notes


 In typical usage, the *string* is a host variable reference to a string containing a dynamically-constructed SQL statement. The case of a literal string is not very useful; you might as well just write the SQL statement directly, without the extra typing of `EXECUTE IMMEDIATE`.


 If you do use a literal string, keep in mind that any double quotes you might wish to include in the SQL statement must be written as octal escapes (`\042`) not the usual C idiom `\"`. This is because the string is inside an `EXEC SQL` section, so the ECPG lexer parses it according to SQL rules not C rules. Any embedded backslashes will later be handled according to C rules; but `\"` causes an immediate syntax error because it is seen as ending the literal.


## Examples


 Here is an example that executes an `INSERT` statement using `EXECUTE IMMEDIATE` and a host variable named `command`:

```

sprintf(command, "INSERT INTO test (name, amount, letter) VALUES ('db: ''r1''', 1, 'f')");
EXEC SQL EXECUTE IMMEDIATE :command;
```


## Compatibility


 `EXECUTE IMMEDIATE` is specified in the SQL standard.
  <a id="ecpg-sql-get-descriptor"></a>

# GET DESCRIPTOR

get information from an SQL descriptor area

## Synopsis


```

GET DESCRIPTOR DESCRIPTOR_NAME :CVARIABLE = DESCRIPTOR_HEADER_ITEM [, ... ]
GET DESCRIPTOR DESCRIPTOR_NAME VALUE COLUMN_NUMBER :CVARIABLE = DESCRIPTOR_ITEM [, ... ]
```


## Description


 `GET DESCRIPTOR` retrieves information about a query result set from an SQL descriptor area and stores it into host variables. A descriptor area is typically populated using `FETCH` or `SELECT` before using this command to transfer the information into host language variables.


 This command has two forms: The first form retrieves descriptor “header” items, which apply to the result set in its entirety. One example is the row count. The second form, which requires the column number as additional parameter, retrieves information about a particular column. Examples are the column name and the actual column value.


## Parameters


<a id="ecpg-sql-get-descriptor-descriptor-name"></a>

*descriptor_name*
:   A descriptor name.
<a id="ecpg-sql-get-descriptor-descriptor-header-item"></a>

*descriptor_header_item*
:   A token identifying which header information item to retrieve. Only `COUNT`, to get the number of columns in the result set, is currently supported.
<a id="ecpg-sql-get-descriptor-column-number"></a>

*column_number*
:   The number of the column about which information is to be retrieved. The count starts at 1.
<a id="ecpg-sql-get-descriptor-descriptor-item"></a>

*descriptor_item*
:   A token identifying which item of information about a column to retrieve. See [Named SQL Descriptor Areas](using-descriptor-areas.md#ecpg-named-descriptors) for a list of supported items.
<a id="ecpg-sql-get-descriptor-cvariable"></a>

*cvariable*
:   A host variable that will receive the data retrieved from the descriptor area.


## Examples


 An example to retrieve the number of columns in a result set:

```

EXEC SQL GET DESCRIPTOR d :d_count = COUNT;
```


 An example to retrieve a data length in the first column:

```

EXEC SQL GET DESCRIPTOR d VALUE 1 :d_returned_octet_length = RETURNED_OCTET_LENGTH;
```


 An example to retrieve the data body of the second column as a string:

```

EXEC SQL GET DESCRIPTOR d VALUE 2 :d_data = DATA;
```


 Here is an example for a whole procedure of executing `SELECT current_database();` and showing the number of columns, the column data length, and the column data:

```

int
main(void)
{
EXEC SQL BEGIN DECLARE SECTION;
    int  d_count;
    char d_data[1024];
    int  d_returned_octet_length;
EXEC SQL END DECLARE SECTION;

    EXEC SQL CONNECT TO testdb AS con1 USER testuser;
    EXEC SQL SELECT pg_catalog.set_config('search_path', '', false); EXEC SQL COMMIT;
    EXEC SQL ALLOCATE DESCRIPTOR d;

    /* Declare, open a cursor, and assign a descriptor to the cursor  */
    EXEC SQL DECLARE cur CURSOR FOR SELECT current_database();
    EXEC SQL OPEN cur;
    EXEC SQL FETCH NEXT FROM cur INTO SQL DESCRIPTOR d;

    /* Get a number of total columns */
    EXEC SQL GET DESCRIPTOR d :d_count = COUNT;
    printf("d_count                 = %d\n", d_count);

    /* Get length of a returned column */
    EXEC SQL GET DESCRIPTOR d VALUE 1 :d_returned_octet_length = RETURNED_OCTET_LENGTH;
    printf("d_returned_octet_length = %d\n", d_returned_octet_length);

    /* Fetch the returned column as a string */
    EXEC SQL GET DESCRIPTOR d VALUE 1 :d_data = DATA;
    printf("d_data                  = %s\n", d_data);

    /* Closing */
    EXEC SQL CLOSE cur;
    EXEC SQL COMMIT;

    EXEC SQL DEALLOCATE DESCRIPTOR d;
    EXEC SQL DISCONNECT ALL;

    return 0;
}
```
 When the example is executed, the result will look like this:

```

d_count                 = 1
d_returned_octet_length = 6
d_data                  = testdb
```


## Compatibility


 `GET DESCRIPTOR` is specified in the SQL standard.


## See Also
  [ecpg-sql-allocate-descriptor](#ecpg-sql-allocate-descriptor), [ecpg-sql-set-descriptor](#ecpg-sql-set-descriptor)  <a id="ecpg-sql-open"></a>

# OPEN

open a dynamic cursor

## Synopsis


```

OPEN CURSOR_NAME
OPEN CURSOR_NAME USING VALUE [, ... ]
OPEN CURSOR_NAME USING SQL DESCRIPTOR DESCRIPTOR_NAME
```


## Description


 `OPEN` opens a cursor and optionally binds actual values to the placeholders in the cursor's declaration. The cursor must previously have been declared with the `DECLARE` command. The execution of `OPEN` causes the query to start executing on the server.


## Parameters


<a id="ecpg-sql-open-cursor-name"></a>

*cursor_name*
:   The name of the cursor to be opened. This can be an SQL identifier or a host variable.
<a id="ecpg-sql-open-value"></a>

*value*
:   A value to be bound to a placeholder in the cursor. This can be an SQL constant, a host variable, or a host variable with indicator.
<a id="ecpg-sql-open-descriptor-name"></a>

*descriptor_name*
:   The name of a descriptor containing values to be bound to the placeholders in the cursor. This can be an SQL identifier or a host variable.


## Examples


```

EXEC SQL OPEN a;
EXEC SQL OPEN d USING 1, 'test';
EXEC SQL OPEN c1 USING SQL DESCRIPTOR mydesc;
EXEC SQL OPEN :curname1;
```


## Compatibility


 `OPEN` is specified in the SQL standard.


## See Also
  [ecpg-sql-declare](#ecpg-sql-declare), [sql-close](../../reference/sql-commands/close.md#sql-close)  <a id="ecpg-sql-prepare"></a>

# PREPARE

prepare a statement for execution

## Synopsis


```

PREPARE PREPARED_NAME FROM STRING
```


## Description


 `PREPARE` prepares a statement dynamically specified as a string for execution. This is different from the direct SQL statement [sql-prepare](../../reference/sql-commands/prepare.md#sql-prepare), which can also be used in embedded programs. The [sql-execute](../../reference/sql-commands/execute.md#sql-execute) command is used to execute either kind of prepared statement.


## Parameters


<a id="ecpg-sql-prepare-prepared-name"></a>

*prepared_name*
:   An identifier for the prepared query.
<a id="ecpg-sql-prepare-string"></a>

*string*
:   A literal string or a host variable containing a preparable SQL statement, one of SELECT, INSERT, UPDATE, or DELETE. Use question marks (`?`) for parameter values to be supplied at execution.


## Notes


 In typical usage, the *string* is a host variable reference to a string containing a dynamically-constructed SQL statement. The case of a literal string is not very useful; you might as well just write a direct SQL `PREPARE` statement.


 If you do use a literal string, keep in mind that any double quotes you might wish to include in the SQL statement must be written as octal escapes (`\042`) not the usual C idiom `\"`. This is because the string is inside an `EXEC SQL` section, so the ECPG lexer parses it according to SQL rules not C rules. Any embedded backslashes will later be handled according to C rules; but `\"` causes an immediate syntax error because it is seen as ending the literal.


## Examples


```

char *stmt = "SELECT * FROM test1 WHERE a = ? AND b = ?";

EXEC SQL ALLOCATE DESCRIPTOR outdesc;
EXEC SQL PREPARE foo FROM :stmt;

EXEC SQL EXECUTE foo USING SQL DESCRIPTOR indesc INTO SQL DESCRIPTOR outdesc;
```


## Compatibility


 `PREPARE` is specified in the SQL standard.


## See Also
  [sql-execute](../../reference/sql-commands/execute.md#sql-execute)  <a id="ecpg-sql-set-autocommit"></a>

# SET AUTOCOMMIT

set the autocommit behavior of the current session

## Synopsis


```

SET AUTOCOMMIT { = | TO } { ON | OFF }
```


## Description


 `SET AUTOCOMMIT` sets the autocommit behavior of the current database session. By default, embedded SQL programs are *not* in autocommit mode, so `COMMIT` needs to be issued explicitly when desired. This command can change the session to autocommit mode, where each individual statement is committed implicitly.


## Compatibility


 `SET AUTOCOMMIT` is an extension of PostgreSQL ECPG.
  <a id="ecpg-sql-set-connection"></a>

# SET CONNECTION

select a database connection

## Synopsis


```

SET CONNECTION [ TO | = ] CONNECTION_NAME
```


## Description


 `SET CONNECTION` sets the “current” database connection, which is the one that all commands use unless overridden.


## Parameters


<a id="ecpg-sql-set-connection-connection-name"></a>

*connection_name*
:   A database connection name established by the `CONNECT` command.
<a id="ecpg-sql-set-connection-current"></a>

`CURRENT`
:   Set the connection to the current connection (thus, nothing happens).


## Examples


```

EXEC SQL SET CONNECTION TO con2;
EXEC SQL SET CONNECTION = con1;
```


## Compatibility


 `SET CONNECTION` is specified in the SQL standard.


## See Also
  [ecpg-sql-connect](#ecpg-sql-connect), [ecpg-sql-disconnect](#ecpg-sql-disconnect)  <a id="ecpg-sql-set-descriptor"></a>

# SET DESCRIPTOR

set information in an SQL descriptor area

## Synopsis


```

SET DESCRIPTOR DESCRIPTOR_NAME DESCRIPTOR_HEADER_ITEM = VALUE [, ... ]
SET DESCRIPTOR DESCRIPTOR_NAME VALUE NUMBER DESCRIPTOR_ITEM = VALUE [, ...]
```


## Description


 `SET DESCRIPTOR` populates an SQL descriptor area with values. The descriptor area is then typically used to bind parameters in a prepared query execution.


 This command has two forms: The first form applies to the descriptor “header”, which is independent of a particular datum. The second form assigns values to particular datums, identified by number.


## Parameters


<a id="ecpg-sql-set-descriptor-descriptor-name"></a>

*descriptor_name*
:   A descriptor name.
<a id="ecpg-sql-set-descriptor-descriptor-header-item"></a>

*descriptor_header_item*
:   A token identifying which header information item to set. Only `COUNT`, to set the number of descriptor items, is currently supported.
<a id="ecpg-sql-set-descriptor-number"></a>

*number*
:   The number of the descriptor item to set. The count starts at 1.
<a id="ecpg-sql-set-descriptor-descriptor-item"></a>

*descriptor_item*
:   A token identifying which item of information to set in the descriptor. See [Named SQL Descriptor Areas](using-descriptor-areas.md#ecpg-named-descriptors) for a list of supported items.
<a id="ecpg-sql-set-descriptor-value"></a>

*value*
:   A value to store into the descriptor item. This can be an SQL constant or a host variable.


## Examples


```

EXEC SQL SET DESCRIPTOR indesc COUNT = 1;
EXEC SQL SET DESCRIPTOR indesc VALUE 1 DATA = 2;
EXEC SQL SET DESCRIPTOR indesc VALUE 1 DATA = :val1;
EXEC SQL SET DESCRIPTOR indesc VALUE 2 INDICATOR = :val1, DATA = 'some string';
EXEC SQL SET DESCRIPTOR indesc VALUE 2 INDICATOR = :val2null, DATA = :val2;
```


## Compatibility


 `SET DESCRIPTOR` is specified in the SQL standard.


## See Also
  [ecpg-sql-allocate-descriptor](#ecpg-sql-allocate-descriptor), [ecpg-sql-get-descriptor](#ecpg-sql-get-descriptor)  <a id="ecpg-sql-type"></a>

# TYPE

define a new data type

## Synopsis


```

TYPE TYPE_NAME IS CTYPE
```


## Description


 The `TYPE` command defines a new C type. It is equivalent to putting a `typedef` into a declare section.


 This command is only recognized when `ecpg` is run with the `-c` option.


## Parameters


<a id="ecpg-sql-type-type-name"></a>

*type_name*
:   The name for the new type. It must be a valid C type name.
<a id="ecpg-sql-type-ctype"></a>

*ctype*
:   A C type specification.


## Examples


```

EXEC SQL TYPE customer IS
    struct
    {
        varchar name[50];
        int     phone;
    };

EXEC SQL TYPE cust_ind IS
    struct ind
    {
        short   name_ind;
        short   phone_ind;
    };

EXEC SQL TYPE c IS char reference;
EXEC SQL TYPE ind IS union { int integer; short smallint; };
EXEC SQL TYPE intarray IS int[AMOUNT];
EXEC SQL TYPE str IS varchar[BUFFERSIZ];
EXEC SQL TYPE string IS char[11];
```


 Here is an example program that uses `EXEC SQL TYPE`:

```

EXEC SQL WHENEVER SQLERROR SQLPRINT;

EXEC SQL TYPE tt IS
    struct
    {
        varchar v[256];
        int     i;
    };

EXEC SQL TYPE tt_ind IS
    struct ind {
        short   v_ind;
        short   i_ind;
    };

int
main(void)
{
EXEC SQL BEGIN DECLARE SECTION;
    tt t;
    tt_ind t_ind;
EXEC SQL END DECLARE SECTION;

    EXEC SQL CONNECT TO testdb AS con1;
    EXEC SQL SELECT pg_catalog.set_config('search_path', '', false); EXEC SQL COMMIT;

    EXEC SQL SELECT current_database(), 256 INTO :t:t_ind LIMIT 1;

    printf("t.v = %s\n", t.v.arr);
    printf("t.i = %d\n", t.i);

    printf("t_ind.v_ind = %d\n", t_ind.v_ind);
    printf("t_ind.i_ind = %d\n", t_ind.i_ind);

    EXEC SQL DISCONNECT con1;

    return 0;
}
```
 The output from this program looks like this:

```

t.v = testdb
t.i = 256
t_ind.v_ind = 0
t_ind.i_ind = 0
```


## Compatibility


 The `TYPE` command is a PostgreSQL extension.
  <a id="ecpg-sql-var"></a>

# VAR

define a variable

## Synopsis


```

VAR VARNAME IS CTYPE
```


## Description


 The `VAR` command assigns a new C data type to a host variable. The host variable must be previously declared in a declare section.


## Parameters


<a id="ecpg-sql-var-varname"></a>

*varname*
:   A C variable name.
<a id="ecpg-sql-var-ctype"></a>

*ctype*
:   A C type specification.


## Examples


```

Exec sql begin declare section;
short a;
exec sql end declare section;
EXEC SQL VAR a IS int;
```


## Compatibility


 The `VAR` command is a PostgreSQL extension.
  <a id="ecpg-sql-whenever"></a>

# WHENEVER

specify the action to be taken when an SQL statement causes a specific class condition to be raised

## Synopsis


```

WHENEVER { NOT FOUND | SQLERROR | SQLWARNING } ACTION
```


## Description


 Define a behavior which is called on the special cases (Rows not found, SQL warnings or errors) in the result of SQL execution.


## Parameters


 See [Setting Callbacks](error-handling.md#ecpg-whenever) for a description of the parameters.


## Examples


```

EXEC SQL WHENEVER NOT FOUND CONTINUE;
EXEC SQL WHENEVER NOT FOUND DO BREAK;
EXEC SQL WHENEVER NOT FOUND DO CONTINUE;
EXEC SQL WHENEVER SQLWARNING SQLPRINT;
EXEC SQL WHENEVER SQLWARNING DO warn();
EXEC SQL WHENEVER SQLERROR sqlprint;
EXEC SQL WHENEVER SQLERROR CALL print2();
EXEC SQL WHENEVER SQLERROR DO handle_error("select");
EXEC SQL WHENEVER SQLERROR DO sqlnotice(NULL, NONO);
EXEC SQL WHENEVER SQLERROR DO sqlprint();
EXEC SQL WHENEVER SQLERROR GOTO error_label;
EXEC SQL WHENEVER SQLERROR STOP;
```


 A typical application is the use of `WHENEVER NOT FOUND BREAK` to handle looping through result sets:

```

int
main(void)
{
    EXEC SQL CONNECT TO testdb AS con1;
    EXEC SQL SELECT pg_catalog.set_config('search_path', '', false); EXEC SQL COMMIT;
    EXEC SQL ALLOCATE DESCRIPTOR d;
    EXEC SQL DECLARE cur CURSOR FOR SELECT current_database(), 'hoge', 256;
    EXEC SQL OPEN cur;

    /* when end of result set reached, break out of while loop */
    EXEC SQL WHENEVER NOT FOUND DO BREAK;

    while (1)
    {
        EXEC SQL FETCH NEXT FROM cur INTO SQL DESCRIPTOR d;
        ...
    }

    EXEC SQL CLOSE cur;
    EXEC SQL COMMIT;

    EXEC SQL DEALLOCATE DESCRIPTOR d;
    EXEC SQL DISCONNECT ALL;

    return 0;
}
```


## Compatibility


 `WHENEVER` is specified in the SQL standard, but most of the actions are PostgreSQL extensions.
