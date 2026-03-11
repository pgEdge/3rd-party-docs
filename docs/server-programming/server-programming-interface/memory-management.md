<a id="spi-memory"></a>

## Memory Management


  PostgreSQL allocates memory within *memory contexts*, which provide a convenient method of managing allocations made in many different places that need to live for differing amounts of time. Destroying a context releases all the memory that was allocated in it. Thus, it is not necessary to keep track of individual objects to avoid memory leaks; instead only a relatively small number of contexts have to be managed. `palloc` and related functions allocate memory from the “current” context.


 `SPI_connect` creates a new memory context and makes it current. `SPI_finish` restores the previous current memory context and destroys the context created by `SPI_connect`. These actions ensure that transient memory allocations made inside your C function are reclaimed at C function exit, avoiding memory leakage.


 However, if your C function needs to return an object in allocated memory (such as a value of a pass-by-reference data type), you cannot allocate that memory using `palloc`, at least not while you are connected to SPI. If you try, the object will be deallocated by `SPI_finish`, and your C function will not work reliably. To solve this problem, use `SPI_palloc` to allocate memory for your return object. `SPI_palloc` allocates memory in the “upper executor context”, that is, the memory context that was current when `SPI_connect` was called, which is precisely the right context for a value returned from your C function. Several of the other utility functions described in this section also return objects created in the upper executor context.


 When `SPI_connect` is called, the private context of the C function, which is created by `SPI_connect`, is made the current context. All allocations made by `palloc`, `repalloc`, or SPI utility functions (except as described in this section) are made in this context. When a C function disconnects from the SPI manager (via `SPI_finish`) the current context is restored to the upper executor context, and all allocations made in the C function memory context are freed and cannot be used any more.
  <a id="spi-spi-palloc"></a>

# SPI_palloc

allocate memory in the upper executor context

## Synopsis


```

void * SPI_palloc(Size size)
```


## Description


 `SPI_palloc` allocates memory in the upper executor context.


 This function can only be used while connected to SPI. Otherwise, it throws an error.


## Arguments


`Size `size``
:   size in bytes of storage to allocate


## Return Value


 pointer to new storage space of the specified size
   <a id="spi-realloc"></a>

# SPI_repalloc

reallocate memory in the upper executor context

## Synopsis


```

void * SPI_repalloc(void * pointer, Size size)
```


## Description


 `SPI_repalloc` changes the size of a memory segment previously allocated using `SPI_palloc`.


 This function is no longer different from plain `repalloc`. It's kept just for backward compatibility of existing code.


## Arguments


`void * `pointer``
:   pointer to existing storage to change

`Size `size``
:   size in bytes of storage to allocate


## Return Value


 pointer to new storage space of specified size with the contents copied from the existing area
   <a id="spi-spi-pfree"></a>

# SPI_pfree

free memory in the upper executor context

## Synopsis


```

void SPI_pfree(void * pointer)
```


## Description


 `SPI_pfree` frees memory previously allocated using `SPI_palloc` or `SPI_repalloc`.


 This function is no longer different from plain `pfree`. It's kept just for backward compatibility of existing code.


## Arguments


`void * `pointer``
:   pointer to existing storage to free
   <a id="spi-spi-copytuple"></a>

# SPI_copytuple

make a copy of a row in the upper executor context

## Synopsis


```

HeapTuple SPI_copytuple(HeapTuple row)
```


## Description


 `SPI_copytuple` makes a copy of a row in the upper executor context. This is normally used to return a modified row from a trigger. In a function declared to return a composite type, use `SPI_returntuple` instead.


 This function can only be used while connected to SPI. Otherwise, it returns NULL and sets `SPI_result` to `SPI_ERROR_UNCONNECTED`.


## Arguments


`HeapTuple `row``
:   row to be copied


## Return Value


 the copied row, or `NULL` on error (see `SPI_result` for an error indication)
   <a id="spi-spi-returntuple"></a>

# SPI_returntuple

prepare to return a tuple as a Datum

## Synopsis


```

HeapTupleHeader SPI_returntuple(HeapTuple row, TupleDesc rowdesc)
```


## Description


 `SPI_returntuple` makes a copy of a row in the upper executor context, returning it in the form of a row type `Datum`. The returned pointer need only be converted to `Datum` via `PointerGetDatum` before returning.


 This function can only be used while connected to SPI. Otherwise, it returns NULL and sets `SPI_result` to `SPI_ERROR_UNCONNECTED`.


 Note that this should be used for functions that are declared to return composite types. It is not used for triggers; use `SPI_copytuple` for returning a modified row in a trigger.


## Arguments


`HeapTuple `row``
:   row to be copied

`TupleDesc `rowdesc``
:   descriptor for row (pass the same descriptor each time for most effective caching)


## Return Value


 `HeapTupleHeader` pointing to copied row, or `NULL` on error (see `SPI_result` for an error indication)
   <a id="spi-spi-modifytuple"></a>

# SPI_modifytuple

create a row by replacing selected fields of a given row

## Synopsis


```

HeapTuple SPI_modifytuple(Relation rel, HeapTuple row, int ncols,
                          int * colnum, Datum * values, const char * nulls)
```


## Description


 `SPI_modifytuple` creates a new row by substituting new values for selected columns, copying the original row's columns at other positions. The input row is not modified. The new row is returned in the upper executor context.


 This function can only be used while connected to SPI. Otherwise, it returns NULL and sets `SPI_result` to `SPI_ERROR_UNCONNECTED`.


## Arguments


`Relation `rel``
:   Used only as the source of the row descriptor for the row. (Passing a relation rather than a row descriptor is a misfeature.)

`HeapTuple `row``
:   row to be modified

`int `ncols``
:   number of columns to be changed

`int * `colnum``
:   an array of length `ncols`, containing the numbers of the columns that are to be changed (column numbers start at 1)

`Datum * `values``
:   an array of length `ncols`, containing the new values for the specified columns

`const char * `nulls``
:   an array of length `ncols`, describing which new values are null


     If `nulls` is `NULL` then `SPI_modifytuple` assumes that no new values are null. Otherwise, each entry of the `nulls` array should be `' '` if the corresponding new value is non-null, or `'n'` if the corresponding new value is null. (In the latter case, the actual value in the corresponding `values` entry doesn't matter.) Note that `nulls` is not a text string, just an array: it does not need a `'\0'` terminator.


## Return Value


 new row with modifications, allocated in the upper executor context, or `NULL` on error (see `SPI_result` for an error indication)


 On error, `SPI_result` is set as follows:

`SPI_ERROR_ARGUMENT`
:   if `rel` is `NULL`, or if `row` is `NULL`, or if `ncols` is less than or equal to 0, or if `colnum` is `NULL`, or if `values` is `NULL`.

`SPI_ERROR_NOATTRIBUTE`
:   if `colnum` contains an invalid column number (less than or equal to 0 or greater than the number of columns in `row`)

`SPI_ERROR_UNCONNECTED`
:   if SPI is not active

   <a id="spi-spi-freetuple"></a>

# SPI_freetuple

free a row allocated in the upper executor context

## Synopsis


```

void SPI_freetuple(HeapTuple row)
```


## Description


 `SPI_freetuple` frees a row previously allocated in the upper executor context.


 This function is no longer different from plain `heap_freetuple`. It's kept just for backward compatibility of existing code.


## Arguments


`HeapTuple `row``
:   row to free
   <a id="spi-spi-freetupletable"></a>

# SPI_freetuptable

free a row set created by `SPI_execute` or a similar function

## Synopsis


```

void SPI_freetuptable(SPITupleTable * tuptable)
```


## Description


 `SPI_freetuptable` frees a row set created by a prior SPI command execution function, such as `SPI_execute`. Therefore, this function is often called with the global variable `SPI_tuptable` as argument.


 This function is useful if an SPI-using C function needs to execute multiple commands and does not want to keep the results of earlier commands around until it ends. Note that any unfreed row sets will be freed anyway at `SPI_finish`. Also, if a subtransaction is started and then aborted within execution of an SPI-using C function, SPI automatically frees any row sets created while the subtransaction was running.


 Beginning in PostgreSQL 9.3, `SPI_freetuptable` contains guard logic to protect against duplicate deletion requests for the same row set. In previous releases, duplicate deletions would lead to crashes.


## Arguments


`SPITupleTable * `tuptable``
:   pointer to row set to free, or NULL to do nothing
   <a id="spi-spi-freeplan"></a>

# SPI_freeplan

free a previously saved prepared statement

## Synopsis


```

int SPI_freeplan(SPIPlanPtr plan)
```


## Description


 `SPI_freeplan` releases a prepared statement previously returned by `SPI_prepare` or saved by `SPI_keepplan` or `SPI_saveplan`.


## Arguments


`SPIPlanPtr `plan``
:   pointer to statement to free


## Return Value


 0 on success; `SPI_ERROR_ARGUMENT` if `plan` is `NULL` or invalid
