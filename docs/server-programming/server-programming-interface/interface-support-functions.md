<a id="spi-interface-support"></a>

## Interface Support Functions


 The functions described here provide an interface for extracting information from result sets returned by `SPI_execute` and other SPI functions.


 All functions described in this section can be used by both connected and unconnected C functions.
  <a id="spi-spi-fname"></a>

# SPI_fname

determine the column name for the specified column number

## Synopsis


```

char * SPI_fname(TupleDesc rowdesc, int colnumber)
```


## Description


 `SPI_fname` returns a copy of the column name of the specified column. (You can use `pfree` to release the copy of the name when you don't need it anymore.)


## Arguments


`TupleDesc `rowdesc``
:   input row description

`int `colnumber``
:   column number (count starts at 1)


## Return Value


 The column name; `NULL` if `colnumber` is out of range. `SPI_result` set to `SPI_ERROR_NOATTRIBUTE` on error.
   <a id="spi-spi-fnumber"></a>

# SPI_fnumber

determine the column number for the specified column name

## Synopsis


```

int SPI_fnumber(TupleDesc rowdesc, const char * colname)
```


## Description


 `SPI_fnumber` returns the column number for the column with the specified name.


 If `colname` refers to a system column (e.g., `ctid`) then the appropriate negative column number will be returned. The caller should be careful to test the return value for exact equality to `SPI_ERROR_NOATTRIBUTE` to detect an error; testing the result for less than or equal to 0 is not correct unless system columns should be rejected.


## Arguments


`TupleDesc `rowdesc``
:   input row description

`const char * `colname``
:   column name


## Return Value


 Column number (count starts at 1 for user-defined columns), or `SPI_ERROR_NOATTRIBUTE` if the named column was not found.
   <a id="spi-spi-getvalue"></a>

# SPI_getvalue

return the string value of the specified column

## Synopsis


```

char * SPI_getvalue(HeapTuple row, TupleDesc rowdesc, int colnumber)
```


## Description


 `SPI_getvalue` returns the string representation of the value of the specified column.


 The result is returned in memory allocated using `palloc`. (You can use `pfree` to release the memory when you don't need it anymore.)


## Arguments


`HeapTuple `row``
:   input row to be examined

`TupleDesc `rowdesc``
:   input row description

`int `colnumber``
:   column number (count starts at 1)


## Return Value


 Column value, or `NULL` if the column is null, `colnumber` is out of range (`SPI_result` is set to `SPI_ERROR_NOATTRIBUTE`), or no output function is available (`SPI_result` is set to `SPI_ERROR_NOOUTFUNC`).
   <a id="spi-spi-getbinval"></a>

# SPI_getbinval

return the binary value of the specified column

## Synopsis


```

Datum SPI_getbinval(HeapTuple row, TupleDesc rowdesc, int colnumber,
                    bool * isnull)
```


## Description


 `SPI_getbinval` returns the value of the specified column in the internal form (as type `Datum`).


 This function does not allocate new space for the datum. In the case of a pass-by-reference data type, the return value will be a pointer into the passed row.


## Arguments


`HeapTuple `row``
:   input row to be examined

`TupleDesc `rowdesc``
:   input row description

`int `colnumber``
:   column number (count starts at 1)

`bool * `isnull``
:   flag for a null value in the column


## Return Value


 The binary value of the column is returned. The variable pointed to by `isnull` is set to true if the column is null, else to false.


 `SPI_result` is set to `SPI_ERROR_NOATTRIBUTE` on error.
   <a id="spi-spi-gettype"></a>

# SPI_gettype

return the data type name of the specified column

## Synopsis


```

char * SPI_gettype(TupleDesc rowdesc, int colnumber)
```


## Description


 `SPI_gettype` returns a copy of the data type name of the specified column. (You can use `pfree` to release the copy of the name when you don't need it anymore.)


## Arguments


`TupleDesc `rowdesc``
:   input row description

`int `colnumber``
:   column number (count starts at 1)


## Return Value


 The data type name of the specified column, or `NULL` on error. `SPI_result` is set to `SPI_ERROR_NOATTRIBUTE` on error.
   <a id="spi-spi-gettypeid"></a>

# SPI_gettypeid

return the data type OID of the specified column

## Synopsis


```

Oid SPI_gettypeid(TupleDesc rowdesc, int colnumber)
```


## Description


 `SPI_gettypeid` returns the OID of the data type of the specified column.


## Arguments


`TupleDesc `rowdesc``
:   input row description

`int `colnumber``
:   column number (count starts at 1)


## Return Value


 The OID of the data type of the specified column or `InvalidOid` on error. On error, `SPI_result` is set to `SPI_ERROR_NOATTRIBUTE`.
   <a id="spi-spi-getrelname"></a>

# SPI_getrelname

return the name of the specified relation

## Synopsis


```

char * SPI_getrelname(Relation rel)
```


## Description


 `SPI_getrelname` returns a copy of the name of the specified relation. (You can use `pfree` to release the copy of the name when you don't need it anymore.)


## Arguments


`Relation `rel``
:   input relation


## Return Value


 The name of the specified relation.
  <a id="spi-spi-getnspname"></a>

# SPI_getnspname

return the namespace of the specified relation

## Synopsis


```

char * SPI_getnspname(Relation rel)
```


## Description


 `SPI_getnspname` returns a copy of the name of the namespace that the specified `Relation` belongs to. This is equivalent to the relation's schema. You should `pfree` the return value of this function when you are finished with it.


## Arguments


`Relation `rel``
:   input relation


## Return Value


 The name of the specified relation's namespace.
  <a id="spi-spi-result-code-string"></a>

# SPI_result_code_string

return error code as string

## Synopsis


```

const char * SPI_result_code_string(int code);
```


## Description


 `SPI_result_code_string` returns a string representation of the result code returned by various SPI functions or stored in `SPI_result`.


## Arguments


`int `code``
:   result code


## Return Value


 A string representation of the result code.
