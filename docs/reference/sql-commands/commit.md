<a id="sql-commit"></a>

# COMMIT

commit the current transaction

## Synopsis


```

COMMIT [ WORK | TRANSACTION ] [ AND [ NO ] CHAIN ]
```


## Description


 `COMMIT` commits the current transaction. All changes made by the transaction become visible to others and are guaranteed to be durable if a crash occurs.


## Parameters


<a id="sql-commit-transaction"></a>

`WORK`, `TRANSACTION`
:   Optional key words. They have no effect.
<a id="sql-commit-chain"></a>

`AND CHAIN`
:   If `AND CHAIN` is specified, a new transaction is immediately started with the same transaction characteristics (see [sql-set-transaction](set-transaction.md#sql-set-transaction)) as the just finished one. Otherwise, no new transaction is started.


## Notes


 Use [sql-rollback](rollback.md#sql-rollback) to abort a transaction.


 Issuing `COMMIT` when not inside a transaction does no harm, but it will provoke a warning message. `COMMIT AND CHAIN` when not inside a transaction is an error.


## Examples


 To commit the current transaction and make all changes permanent:

```sql

COMMIT;
```


## Compatibility


 The command `COMMIT` conforms to the SQL standard. The form `COMMIT TRANSACTION` is a PostgreSQL extension.


## See Also
  [sql-begin](begin.md#sql-begin), [sql-rollback](rollback.md#sql-rollback)
