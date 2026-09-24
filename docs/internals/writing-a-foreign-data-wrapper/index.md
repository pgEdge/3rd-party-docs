<a id="fdwhandler"></a>

# Writing a Foreign Data Wrapper

 All operations on a foreign table are handled through its foreign data wrapper, which consists of a set of functions that the core server calls. The foreign data wrapper is responsible for fetching data from the remote data source and returning it to the PostgreSQL executor. If updating foreign tables is to be supported, the wrapper must handle that, too. This chapter outlines how to write a new foreign data wrapper.

 The foreign data wrappers included in the standard distribution are good references when trying to write your own. Look into the `contrib` subdirectory of the source tree. The [sql-createforeigndatawrapper](../../reference/sql-commands/create-foreign-data-wrapper.md#sql-createforeigndatawrapper) reference page also has some useful details.

!!! note

    The SQL standard specifies an interface for writing foreign data wrappers. However, PostgreSQL does not implement that API, because the effort to accommodate it into PostgreSQL would be large, and the standard API hasn't gained wide adoption anyway.

- [Foreign Data Wrapper Functions](foreign-data-wrapper-functions.md#fdw-functions)
- [Foreign Data Wrapper Callback Routines](foreign-data-wrapper-callback-routines.md#fdw-callbacks)
- [Foreign Data Wrapper Helper Functions](foreign-data-wrapper-helper-functions.md#fdw-helpers)
- [Foreign Data Wrapper Query Planning](foreign-data-wrapper-query-planning.md#fdw-planning)
- [Row Locking in Foreign Data Wrappers](row-locking-in-foreign-data-wrappers.md#fdw-row-locking)
