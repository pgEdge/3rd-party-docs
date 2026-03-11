<a id="functions"></a>

# Functions and Operators

 PostgreSQL provides a large number of functions and operators for the built-in data types. This chapter describes most of them, although additional special-purpose functions appear in relevant sections of the manual. Users can also define their own functions and operators, as described in [Server Programming](../../server-programming/index.md#server-programming). The psql commands `\df` and `\do` can be used to list all available functions and operators, respectively.

 The notation used throughout this chapter to describe the argument and result data types of a function or operator is like this:

```

repeat ( text, integer ) text
```
 which says that the function `repeat` takes one text and one integer argument and returns a result of type text. The right arrow is also used to indicate the result of an example, thus:

```

repeat('Pg', 4) PgPgPgPg
```


 If you are concerned about portability then note that most of the functions and operators described in this chapter, with the exception of the most trivial arithmetic and comparison operators and some explicitly marked functions, are not specified by the SQL standard. Some of this extended functionality is present in other SQL database management systems, and in many cases this functionality is compatible and consistent between the various implementations.

- [Logical Operators](logical-operators.md#functions-logical)
- [Comparison Functions and Operators](comparison-functions-and-operators.md#functions-comparison)
- [Mathematical Functions and Operators](mathematical-functions-and-operators.md#functions-math)
- [String Functions and Operators](string-functions-and-operators.md#functions-string)
- [Binary String Functions and Operators](binary-string-functions-and-operators.md#functions-binarystring)
- [Bit String Functions and Operators](bit-string-functions-and-operators.md#functions-bitstring)
- [Pattern Matching](pattern-matching.md#functions-matching)
- [Data Type Formatting Functions](data-type-formatting-functions.md#functions-formatting)
- [Date/Time Functions and Operators](date-time-functions-and-operators.md#functions-datetime)
- [Enum Support Functions](enum-support-functions.md#functions-enum)
- [Geometric Functions and Operators](geometric-functions-and-operators.md#functions-geometry)
- [Network Address Functions and Operators](network-address-functions-and-operators.md#functions-net)
- [Text Search Functions and Operators](text-search-functions-and-operators.md#functions-textsearch)
- [UUID Functions](uuid-functions.md#functions-uuid)
- [XML Functions](xml-functions.md#functions-xml)
- [JSON Functions and Operators](json-functions-and-operators.md#functions-json)
- [Sequence Manipulation Functions](sequence-manipulation-functions.md#functions-sequence)
- [Conditional Expressions](conditional-expressions.md#functions-conditional)
- [Array Functions and Operators](array-functions-and-operators.md#functions-array)
- [Range/Multirange Functions and Operators](range-multirange-functions-and-operators.md#functions-range)
- [Aggregate Functions](aggregate-functions.md#functions-aggregate)
- [Window Functions](window-functions.md#functions-window)
- [Merge Support Functions](merge-support-functions.md#functions-merge-support)
- [Subquery Expressions](subquery-expressions.md#functions-subquery)
- [Row and Array Comparisons](row-and-array-comparisons.md#functions-comparisons)
- [Set Returning Functions](set-returning-functions.md#functions-srf)
- [System Information Functions and Operators](system-information-functions-and-operators.md#functions-info)
- [System Administration Functions](system-administration-functions.md#functions-admin)
- [Trigger Functions](trigger-functions.md#functions-trigger)
- [Event Trigger Functions](event-trigger-functions.md#functions-event-triggers)
- [Statistics Information Functions](statistics-information-functions.md#functions-statistics)
