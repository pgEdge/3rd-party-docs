<a id="sql-commit-prepared"></a>

# COMMIT PREPARED

commit a transaction that was earlier prepared for two-phase commit

## Synopsis


```

COMMIT PREPARED TRANSACTION_ID
```


## Description


 `COMMIT PREPARED` commits a transaction that is in prepared state.


## Parameters


*transaction_id*
:   The transaction identifier of the transaction that is to be committed.


## Notes


 To commit a prepared transaction, you must be either the same user that executed the transaction originally, or a superuser. But you do not have to be in the same session that executed the transaction.


 This command cannot be executed inside a transaction block. The prepared transaction is committed immediately.


 All currently available prepared transactions are listed in the [`pg_prepared_xacts`](../../internals/system-views/pg_prepared_xacts.md#view-pg-prepared-xacts) system view.
 <a id="sql-commit-prepared-examples"></a>

## Examples


 Commit the transaction identified by the transaction identifier `foobar`:

```sql

COMMIT PREPARED 'foobar';
```


## Compatibility


 `COMMIT PREPARED` is a PostgreSQL extension. It is intended for use by external transaction management systems, some of which are covered by standards (such as X/Open XA), but the SQL side of those systems is not standardized.


## See Also
  [sql-prepare-transaction](prepare-transaction.md#sql-prepare-transaction), [sql-rollback-prepared](rollback-prepared.md#sql-rollback-prepared)
