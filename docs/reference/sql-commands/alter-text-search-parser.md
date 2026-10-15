<a id="sql-altertsparser"></a>

# ALTER TEXT SEARCH PARSER

change the definition of a text search parser

## Synopsis


```

ALTER TEXT SEARCH PARSER NAME RENAME TO NEW_NAME
ALTER TEXT SEARCH PARSER NAME SET SCHEMA NEW_SCHEMA
```


## Description


 `ALTER TEXT SEARCH PARSER` changes the definition of a text search parser. Currently, the only supported functionality is to change the parser's name.


 You must be a superuser to use `ALTER TEXT SEARCH PARSER`.


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing text search parser.

*new_name*
:   The new name of the text search parser.

*new_schema*
:   The new schema for the text search parser.


## Compatibility


 There is no `ALTER TEXT SEARCH PARSER` statement in the SQL standard.


## See Also
  [sql-createtsparser](create-text-search-parser.md#sql-createtsparser), [sql-droptsparser](drop-text-search-parser.md#sql-droptsparser)
