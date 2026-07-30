<a id="sql-createtsparser"></a>

# CREATE TEXT SEARCH PARSER

define a new text search parser

## Synopsis


```

CREATE TEXT SEARCH PARSER NAME (
    START = START_FUNCTION ,
    GETTOKEN = GETTOKEN_FUNCTION ,
    END = END_FUNCTION ,
    LEXTYPES = LEXTYPES_FUNCTION
    [, HEADLINE = HEADLINE_FUNCTION ]
)
```


## Description


 `CREATE TEXT SEARCH PARSER` creates a new text search parser. A text search parser defines a method for splitting a text string into tokens and assigning types (categories) to the tokens. A parser is not particularly useful by itself, but must be bound into a text search configuration along with some text search dictionaries to be used for searching.


 If a schema name is given then the text search parser is created in the specified schema. Otherwise it is created in the current schema.


 You must be a superuser to use `CREATE TEXT SEARCH PARSER`. (This restriction is made because an erroneous text search parser definition could confuse or even crash the server.)


 Refer to [Full Text Search](../../the-sql-language/full-text-search/index.md#textsearch) for further information.


## Parameters


*name*
:   The name of the text search parser to be created. The name can be schema-qualified.

*start_function*
:   The name of the start function for the parser.

*gettoken_function*
:   The name of the get-next-token function for the parser.

*end_function*
:   The name of the end function for the parser.

*lextypes_function*
:   The name of the lextypes function for the parser (a function that returns information about the set of token types it produces).

*headline_function*
:   The name of the headline function for the parser (a function that summarizes a set of tokens).


 The function names can be schema-qualified if necessary. Argument types are not given, since the argument list for each type of function is predetermined. All except the headline function are required.


 The arguments can appear in any order, not only the one shown above.


## Compatibility


 There is no `CREATE TEXT SEARCH PARSER` statement in the SQL standard.


## See Also
  [sql-altertsparser](alter-text-search-parser.md#sql-altertsparser), [sql-droptsparser](drop-text-search-parser.md#sql-droptsparser)
