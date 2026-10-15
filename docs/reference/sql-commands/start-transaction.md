<a id="sql-start-transaction"></a>

# START TRANSACTION

start a transaction block

## Synopsis


```

START TRANSACTION [ TRANSACTION_MODE [, ...] ]

where TRANSACTION_MODE is one of:

    ISOLATION LEVEL { SERIALIZABLE | REPEATABLE READ | READ COMMITTED | READ UNCOMMITTED }
    READ WRITE | READ ONLY
    [ NOT ] DEFERRABLE
```


## Description


 This command begins a new transaction block. If the isolation level, read/write mode, or deferrable mode is specified, the new transaction has those characteristics, as if [`SET TRANSACTION`](set-transaction.md#sql-set-transaction) was executed. This is the same as the [`BEGIN`](begin.md#sql-begin) command.


## Parameters


 Refer to [sql-set-transaction](set-transaction.md#sql-set-transaction) for information on the meaning of the parameters to this statement.


## Compatibility


 In the standard, it is not necessary to issue `START TRANSACTION` to start a transaction block: any SQL command implicitly begins a block. PostgreSQL's behavior can be seen as implicitly issuing a `COMMIT` after each command that does not follow `START TRANSACTION` (or `BEGIN`), and it is therefore often called “autocommit”. Other relational database systems might offer an autocommit feature as a convenience.


 The `DEFERRABLE` *transaction_mode* is a PostgreSQL language extension.


 The SQL standard requires commas between successive *transaction_modes*, but for historical reasons PostgreSQL allows the commas to be omitted.


 See also the compatibility section of [sql-set-transaction](set-transaction.md#sql-set-transaction).


## See Also
  [sql-begin](begin.md#sql-begin), [sql-commit](commit.md#sql-commit), [sql-rollback](rollback.md#sql-rollback), [sql-savepoint](savepoint.md#sql-savepoint), [sql-set-transaction](set-transaction.md#sql-set-transaction)
