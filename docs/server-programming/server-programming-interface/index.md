<a id="spi"></a>

# Server Programming Interface

 The *Server Programming Interface* (SPI) gives writers of user-defined C functions the ability to run SQL commands inside their functions or procedures. SPI is a set of interface functions to simplify access to the parser, planner, and executor. SPI also does some memory management.

!!! note

    The available procedural languages provide various means to execute SQL commands from functions. Most of these facilities are based on SPI, so this documentation might be of use for users of those languages as well.

 Note that if a command invoked via SPI fails, then control will not be returned to your C function. Rather, the transaction or subtransaction in which your C function executes will be rolled back. (This might seem surprising given that the SPI functions mostly have documented error-return conventions. Those conventions only apply for errors detected within the SPI functions themselves, however.) It is possible to recover control after an error by establishing your own subtransaction surrounding SPI calls that might fail.

 SPI functions return a nonnegative result on success (either via a returned integer value or in the global variable `SPI_result`, as described below). On error, a negative result or `NULL` will be returned.

 Source code files that use SPI must include the header file `executor/spi.h`.

- [Interface Functions](interface-functions.md#spi-interface)
- [Interface Support Functions](interface-support-functions.md#spi-interface-support)
- [Memory Management](memory-management.md#spi-memory)
- [Transaction Management](transaction-management.md#spi-transaction)
- [Visibility of Data Changes](visibility-of-data-changes.md#spi-visibility)
- [Examples](examples.md#spi-examples)
