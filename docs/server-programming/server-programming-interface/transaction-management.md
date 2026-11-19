<a id="spi-transaction"></a>

## Transaction Management


 It is not possible to run transaction control commands such as `COMMIT` and `ROLLBACK` through SPI functions such as `SPI_execute`. There are, however, separate interface functions that allow transaction control through SPI.


 It is not generally safe and sensible to start and end transactions in arbitrary user-defined SQL-callable functions without taking into account the context in which they are called. For example, a transaction boundary in the middle of a function that is part of a complex SQL expression that is part of some SQL command will probably result in obscure internal errors or crashes. The interface functions presented here are primarily intended to be used by procedural language implementations to support transaction management in SQL-level procedures that are invoked by the `CALL` command, taking the context of the `CALL` invocation into account. SPI-using procedures implemented in C can implement the same logic, but the details of that are beyond the scope of this documentation.
  <a id="spi-spi-commit"></a>

# SPI_commit

commit the current transaction

## Synopsis


```

void SPI_commit(void)
```


```

void SPI_commit_and_chain(void)
```


## Description


 `SPI_commit` commits the current transaction. It is approximately equivalent to running the SQL command `COMMIT`. After the transaction is committed, a new transaction is automatically started using default transaction characteristics, so that the caller can continue using SPI facilities. If there is a failure during commit, the current transaction is instead rolled back and a new transaction is started, after which the error is thrown in the usual way.


 `SPI_commit_and_chain` is the same, but the new transaction is started with the same transaction characteristics as the just finished one, like with the SQL command `COMMIT AND CHAIN`.


 These functions can only be executed if the SPI connection has been set as nonatomic in the call to `SPI_connect_ext`.
   <a id="spi-spi-rollback"></a>

# SPI_rollback

abort the current transaction

## Synopsis


```

void SPI_rollback(void)
```


```

void SPI_rollback_and_chain(void)
```


## Description


 `SPI_rollback` rolls back the current transaction. It is approximately equivalent to running the SQL command `ROLLBACK`. After the transaction is rolled back, a new transaction is automatically started using default transaction characteristics, so that the caller can continue using SPI facilities.


 `SPI_rollback_and_chain` is the same, but the new transaction is started with the same transaction characteristics as the just finished one, like with the SQL command `ROLLBACK AND CHAIN`.


 These functions can only be executed if the SPI connection has been set as nonatomic in the call to `SPI_connect_ext`.
   <a id="spi-spi-start-transaction"></a>

# SPI_start_transaction

obsolete function

## Synopsis


```

void SPI_start_transaction(void)
```


## Description


 `SPI_start_transaction` does nothing, and exists only for code compatibility with earlier PostgreSQL releases. It used to be required after calling `SPI_commit` or `SPI_rollback`, but now those functions start a new transaction automatically.
