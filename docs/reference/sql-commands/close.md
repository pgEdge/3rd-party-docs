<a id="sql-close"></a>

# CLOSE

close a cursor

## Synopsis


```

CLOSE { NAME | ALL }
```


## Description


 `CLOSE` frees the resources associated with an open cursor. After the cursor is closed, no subsequent operations are allowed on it. A cursor should be closed when it is no longer needed.


 Every non-holdable open cursor is implicitly closed when a transaction is terminated by `COMMIT` or `ROLLBACK`. A holdable cursor is implicitly closed if the transaction that created it aborts via `ROLLBACK`. If the creating transaction successfully commits, the holdable cursor remains open until an explicit `CLOSE` is executed, or the client disconnects.


## Parameters


*name*
:   The name of an open cursor to close.

`ALL`
:   Close all open cursors.


## Notes


 PostgreSQL does not have an explicit `OPEN` cursor statement; a cursor is considered open when it is declared. Use the [`DECLARE`](declare.md#sql-declare) statement to declare a cursor.


 You can see all available cursors by querying the [`pg_cursors`](../../internals/system-views/pg_cursors.md#view-pg-cursors) system view.


 If a cursor is closed after a savepoint which is later rolled back, the `CLOSE` is not rolled back; that is, the cursor remains closed.


## Examples


 Close the cursor `liahona`:

```

CLOSE liahona;
```


## Compatibility


 `CLOSE` is fully conforming with the SQL standard. `CLOSE ALL` is a PostgreSQL extension.


## See Also
  [sql-declare](declare.md#sql-declare), [sql-fetch](fetch.md#sql-fetch), [sql-move](move.md#sql-move)
