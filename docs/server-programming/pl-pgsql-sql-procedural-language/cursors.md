<a id="plpgsql-cursors"></a>

## Cursors


 Rather than executing a whole query at once, it is possible to set up a *cursor* that encapsulates the query, and then read the query result a few rows at a time. One reason for doing this is to avoid memory overrun when the result contains a large number of rows. (However, PL/pgSQL users do not normally need to worry about that, since `FOR` loops automatically use a cursor internally to avoid memory problems.) A more interesting usage is to return a reference to a cursor that a function has created, allowing the caller to read the rows. This provides an efficient way to return large row sets from functions.
 <a id="plpgsql-cursor-declarations"></a>

### Declaring Cursor Variables


 All access to cursors in PL/pgSQL goes through cursor variables, which are always of the special data type `refcursor`. One way to create a cursor variable is just to declare it as a variable of type `refcursor`. Another way is to use the cursor declaration syntax, which in general is:

```

NAME [ [ NO ] SCROLL ] CURSOR [ ( ARGUMENTS ) ] FOR QUERY;
```
 (`FOR` can be replaced by `IS` for Oracle compatibility.) If `SCROLL` is specified, the cursor will be capable of scrolling backward; if `NO SCROLL` is specified, backward fetches will be rejected; if neither specification appears, it is query-dependent whether backward fetches will be allowed. *arguments*, if specified, is a comma-separated list of pairs <em>name</em><code>
     </code><em>datatype</em> that define names to be replaced by parameter values in the given query. The actual values to substitute for these names will be specified later, when the cursor is opened.


 Some examples:

```

DECLARE
    curs1 refcursor;
    curs2 CURSOR FOR SELECT * FROM tenk1;
    curs3 CURSOR (key integer) FOR SELECT * FROM tenk1 WHERE unique1 = key;
```
 All three of these variables have the data type `refcursor`, but the first can be used with any query, while the second has a fully specified query already *bound* to it, and the last has a parameterized query bound to it. (`key` will be replaced by an integer parameter value when the cursor is opened.) The variable `curs1` is said to be *unbound* since it is not bound to any particular query.


 The `SCROLL` option cannot be used when the cursor's query uses `FOR UPDATE/SHARE`. Also, it is best to use `NO SCROLL` with a query that involves volatile functions. The implementation of `SCROLL` assumes that re-reading the query's output will give consistent results, which a volatile function might not do.
  <a id="plpgsql-cursor-opening"></a>

### Opening Cursors


 Before a cursor can be used to retrieve rows, it must be *opened*. (This is the equivalent action to the SQL command [`DECLARE CURSOR`](../../reference/sql-commands/declare.md#sql-declare).) PL/pgSQL has three forms of the `OPEN` statement, two of which use unbound cursor variables while the third uses a bound cursor variable.


!!! note

    Bound cursor variables can also be used without explicitly opening the cursor, via the `FOR` statement described in [Looping through a Cursor's Result](#plpgsql-cursor-for-loop). A `FOR` loop will open the cursor and then close it again when the loop completes.


 Opening a cursor involves creating a server-internal data structure called a *portal*, which holds the execution state for the cursor's query. A portal has a name, which must be unique within the session for the duration of the portal's existence. By default, PL/pgSQL will assign a unique name to each portal it creates. However, if you assign a non-null string value to a cursor variable, that string will be used as its portal name. This feature can be used as described in [Returning Cursors](#plpgsql-cursor-returning).
 <a id="plpgsql-cursor-opening-open-for-query"></a>

#### `OPEN FOR` *query*


```

OPEN UNBOUND_CURSORVAR [ [ NO ] SCROLL ] FOR QUERY;
```


 The cursor variable is opened and given the specified query to execute. The cursor cannot be open already, and it must have been declared as an unbound cursor variable (that is, as a simple `refcursor` variable). The query must be a `SELECT`, or something else that returns rows (such as `EXPLAIN`). The query is treated in the same way as other SQL commands in PL/pgSQL: PL/pgSQL variable names are substituted, and the query plan is cached for possible reuse. When a PL/pgSQL variable is substituted into the cursor query, the value that is substituted is the one it has at the time of the `OPEN`; subsequent changes to the variable will not affect the cursor's behavior. The `SCROLL` and `NO SCROLL` options have the same meanings as for a bound cursor.


 An example:

```

OPEN curs1 FOR SELECT * FROM foo WHERE key = mykey;
```

  <a id="plpgsql-cursor-opening-open-for-execute"></a>

#### `OPEN FOR EXECUTE`


```

OPEN UNBOUND_CURSORVAR [ [ NO ] SCROLL ] FOR EXECUTE QUERY_STRING
                                     [ USING EXPRESSION [, ... ] ];
```


 The cursor variable is opened and given the specified query to execute. The cursor cannot be open already, and it must have been declared as an unbound cursor variable (that is, as a simple `refcursor` variable). The query is specified as a string expression, in the same way as in the `EXECUTE` command. As usual, this gives flexibility so the query plan can vary from one run to the next (see [Plan Caching](pl-pgsql-under-the-hood.md#plpgsql-plan-caching)), and it also means that variable substitution is not done on the command string. As with `EXECUTE`, parameter values can be inserted into the dynamic command via `format()` and `USING`. The `SCROLL` and `NO SCROLL` options have the same meanings as for a bound cursor.


 An example:

```

OPEN curs1 FOR EXECUTE format('SELECT * FROM %I WHERE col1 = $1',tabname) USING keyvalue;
```
 In this example, the table name is inserted into the query via `format()`. The comparison value for `col1` is inserted via a `USING` parameter, so it needs no quoting.
  <a id="plpgsql-open-bound-cursor"></a>

#### Opening a Bound Cursor


```

OPEN BOUND_CURSORVAR [ ( [ ARGUMENT_NAME := ] ARGUMENT_VALUE [, ...] ) ];
```


 This form of `OPEN` is used to open a cursor variable whose query was bound to it when it was declared. The cursor cannot be open already. A list of actual argument value expressions must appear if and only if the cursor was declared to take arguments. These values will be substituted in the query.


 The query plan for a bound cursor is always considered cacheable; there is no equivalent of `EXECUTE` in this case. Notice that `SCROLL` and `NO SCROLL` cannot be specified in `OPEN`, as the cursor's scrolling behavior was already determined.


 Argument values can be passed using either *positional* or *named* notation. In positional notation, all arguments are specified in order. In named notation, each argument's name is specified using `:=` to separate it from the argument expression. Similar to calling functions, described in [Calling Functions](../../the-sql-language/sql-syntax/calling-functions.md#sql-syntax-calling-funcs), it is also allowed to mix positional and named notation.


 Examples (these use the cursor declaration examples above):

```

OPEN curs2;
OPEN curs3(42);
OPEN curs3(key := 42);
```


 Because variable substitution is done on a bound cursor's query, there are really two ways to pass values into the cursor: either with an explicit argument to `OPEN`, or implicitly by referencing a PL/pgSQL variable in the query. However, only variables declared before the bound cursor was declared will be substituted into it. In either case the value to be passed is determined at the time of the `OPEN`. For example, another way to get the same effect as the `curs3` example above is

```

DECLARE
    key integer;
    curs4 CURSOR FOR SELECT * FROM tenk1 WHERE unique1 = key;
BEGIN
    key := 42;
    OPEN curs4;
```

   <a id="plpgsql-cursor-using"></a>

### Using Cursors


 Once a cursor has been opened, it can be manipulated with the statements described here.


 These manipulations need not occur in the same function that opened the cursor to begin with. You can return a `refcursor` value out of a function and let the caller operate on the cursor. (Internally, a `refcursor` value is simply the string name of the portal containing the active query for the cursor. This name can be passed around, assigned to other `refcursor` variables, and so on, without disturbing the portal.)


 All portals are implicitly closed at transaction end. Therefore a `refcursor` value is usable to reference an open cursor only until the end of the transaction.
 <a id="plpgsql-cursor-using-fetch"></a>

#### `FETCH`


```

FETCH [ DIRECTION { FROM | IN } ] CURSOR INTO TARGET;
```


 `FETCH` retrieves the next row (in the indicated direction) from the cursor into a target, which might be a row variable, a record variable, or a comma-separated list of simple variables, just like `SELECT INTO`. If there is no suitable row, the target is set to NULL(s). As with `SELECT INTO`, the special variable `FOUND` can be checked to see whether a row was obtained or not. If no row is obtained, the cursor is positioned after the last row or before the first row, depending on the movement direction.


 The *direction* clause can be any of the variants allowed in the SQL [sql-fetch](../../reference/sql-commands/fetch.md#sql-fetch) command except the ones that can fetch more than one row; namely, it can be `NEXT`, `PRIOR`, `FIRST`, `LAST`, `ABSOLUTE` *count*, `RELATIVE` *count*, `FORWARD`, or `BACKWARD`. Omitting *direction* is the same as specifying `NEXT`. In the forms using a *count*, the *count* can be any integer-valued expression (unlike the SQL `FETCH` command, which only allows an integer constant). *direction* values that require moving backward are likely to fail unless the cursor was declared or opened with the `SCROLL` option.


 *cursor* must be the name of a `refcursor` variable that references an open cursor portal.


 Examples:

```

FETCH curs1 INTO rowvar;
FETCH curs2 INTO foo, bar, baz;
FETCH LAST FROM curs3 INTO x, y;
FETCH RELATIVE -2 FROM curs4 INTO x;
```

  <a id="plpgsql-cursor-using-move"></a>

#### `MOVE`


```

MOVE [ DIRECTION { FROM | IN } ] CURSOR;
```


 `MOVE` repositions a cursor without retrieving any data. `MOVE` works like the `FETCH` command, except it only repositions the cursor and does not return the row moved to. The *direction* clause can be any of the variants allowed in the SQL [sql-fetch](../../reference/sql-commands/fetch.md#sql-fetch) command, including those that can fetch more than one row; the cursor is positioned to the last such row. (However, the case in which the *direction* clause is simply a *count* expression with no key word is deprecated in PL/pgSQL. That syntax is ambiguous with the case where the *direction* clause is omitted altogether, and hence it may fail if the *count* is not a constant.) As with `SELECT INTO`, the special variable `FOUND` can be checked to see whether there was a row to move to. If there is no such row, the cursor is positioned after the last row or before the first row, depending on the movement direction.


 Examples:

```

MOVE curs1;
MOVE LAST FROM curs3;
MOVE RELATIVE -2 FROM curs4;
MOVE FORWARD 2 FROM curs4;
```

  <a id="plpgsql-cursor-using-update-delete"></a>

#### `UPDATE/DELETE WHERE CURRENT OF`


```

UPDATE TABLE SET ... WHERE CURRENT OF CURSOR;
DELETE FROM TABLE WHERE CURRENT OF CURSOR;
```


 When a cursor is positioned on a table row, that row can be updated or deleted using the cursor to identify the row. There are restrictions on what the cursor's query can be (in particular, no grouping) and it's best to use `FOR UPDATE` in the cursor. For more information see the [sql-declare](../../reference/sql-commands/declare.md#sql-declare) reference page.


 An example:

```sql

UPDATE foo SET dataval = myval WHERE CURRENT OF curs1;
```

  <a id="plpgsql-cursor-using-close"></a>

#### `CLOSE`


```

CLOSE CURSOR;
```


 `CLOSE` closes the portal underlying an open cursor. This can be used to release resources earlier than end of transaction, or to free up the cursor variable to be opened again.


 An example:

```

CLOSE curs1;
```

  <a id="plpgsql-cursor-returning"></a>

#### Returning Cursors


 PL/pgSQL functions can return cursors to the caller. This is useful to return multiple rows or columns, especially with very large result sets. To do this, the function opens the cursor and returns the cursor name to the caller (or simply opens the cursor using a portal name specified by or otherwise known to the caller). The caller can then fetch rows from the cursor. The cursor can be closed by the caller, or it will be closed automatically when the transaction closes.


 The portal name used for a cursor can be specified by the programmer or automatically generated. To specify a portal name, simply assign a string to the `refcursor` variable before opening it. The string value of the `refcursor` variable will be used by `OPEN` as the name of the underlying portal. However, if the `refcursor` variable's value is null (as it will be by default), then `OPEN` automatically generates a name that does not conflict with any existing portal, and assigns it to the `refcursor` variable.


!!! note

    Prior to PostgreSQL 16, bound cursor variables were initialized to contain their own names, rather than being left as null, so that the underlying portal name would be the same as the cursor variable's name by default. This was changed because it created too much risk of conflicts between similarly-named cursors in different functions.


 The following example shows one way a cursor name can be supplied by the caller:

```sql

CREATE TABLE test (col text);
INSERT INTO test VALUES ('123');

CREATE FUNCTION reffunc(refcursor) RETURNS refcursor AS '
BEGIN
    OPEN $1 FOR SELECT col FROM test;
    RETURN $1;
END;
' LANGUAGE plpgsql;

BEGIN;
SELECT reffunc('funccursor');
FETCH ALL IN funccursor;
COMMIT;
```


 The following example uses automatic cursor name generation:

```sql

CREATE FUNCTION reffunc2() RETURNS refcursor AS '
DECLARE
    ref refcursor;
BEGIN
    OPEN ref FOR SELECT col FROM test;
    RETURN ref;
END;
' LANGUAGE plpgsql;

-- need to be in a transaction to use cursors.
BEGIN;
SELECT reffunc2();

      reffunc2
--------------------

(1 row)

FETCH ALL IN "";
COMMIT;
```


 The following example shows one way to return multiple cursors from a single function:

```sql

CREATE FUNCTION myfunc(refcursor, refcursor) RETURNS SETOF refcursor AS $$
BEGIN
    OPEN $1 FOR SELECT * FROM table_1;
    RETURN NEXT $1;
    OPEN $2 FOR SELECT * FROM table_2;
    RETURN NEXT $2;
END;
$$ LANGUAGE plpgsql;

-- need to be in a transaction to use cursors.
BEGIN;

SELECT * FROM myfunc('a', 'b');

FETCH ALL FROM a;
FETCH ALL FROM b;
COMMIT;
```

   <a id="plpgsql-cursor-for-loop"></a>

### Looping through a Cursor's Result


 There is a variant of the `FOR` statement that allows iterating through the rows returned by a cursor. The syntax is:

```

[ <<LABEL>> ]
FOR RECORDVAR IN BOUND_CURSORVAR [ ( [ ARGUMENT_NAME := ] ARGUMENT_VALUE [, ...] ) ] LOOP
    STATEMENTS
END LOOP [ LABEL ];
```
 The cursor variable must have been bound to some query when it was declared, and it *cannot* be open already. The `FOR` statement automatically opens the cursor, and it closes the cursor again when the loop exits. A list of actual argument value expressions must appear if and only if the cursor was declared to take arguments. These values will be substituted in the query, in just the same way as during an `OPEN` (see [Opening a Bound Cursor](#plpgsql-open-bound-cursor)).


 The variable *recordvar* is automatically defined as type `record` and exists only inside the loop (any existing definition of the variable name is ignored within the loop). Each row returned by the cursor is successively assigned to this record variable and the loop body is executed.
