<a id="sql-rollback"></a>

# ROLLBACK

abort the current transaction

## Synopsis


```

ROLLBACK [ WORK | TRANSACTION ] [ AND [ NO ] CHAIN ]
```


## Description


 `ROLLBACK` rolls back the current transaction and causes all the updates made by the transaction to be discarded.


## Parameters


<a id="sql-rollback-transaction"></a>

`WORK`, `TRANSACTION`
:   Optional key words. They have no effect.
<a id="sql-rollback-chain"></a>

`AND CHAIN`
:   If `AND CHAIN` is specified, a new (not aborted) transaction is immediately started with the same transaction characteristics (see [sql-set-transaction](set-transaction.md#sql-set-transaction)) as the just finished one. Otherwise, no new transaction is started.


## Notes


 Use [`COMMIT`](commit.md#sql-commit) to successfully terminate a transaction.


 Issuing `ROLLBACK` outside of a transaction block emits a warning and otherwise has no effect. `ROLLBACK AND CHAIN` outside of a transaction block is an error.


## Examples


 To abort all changes:

```

ROLLBACK;
```


## Compatibility


 The command `ROLLBACK` conforms to the SQL standard. The form `ROLLBACK TRANSACTION` is a PostgreSQL extension.


## See Also
  [sql-begin](begin.md#sql-begin), [sql-commit](commit.md#sql-commit), [sql-rollback-to](rollback-to-savepoint.md#sql-rollback-to)
