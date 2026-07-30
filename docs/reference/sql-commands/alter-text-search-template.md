<a id="sql-altertstemplate"></a>

# ALTER TEXT SEARCH TEMPLATE

change the definition of a text search template

## Synopsis


```

ALTER TEXT SEARCH TEMPLATE NAME RENAME TO NEW_NAME
ALTER TEXT SEARCH TEMPLATE NAME SET SCHEMA NEW_SCHEMA
```


## Description


 `ALTER TEXT SEARCH TEMPLATE` changes the definition of a text search template. Currently, the only supported functionality is to change the template's name.


 You must be a superuser to use `ALTER TEXT SEARCH TEMPLATE`.


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing text search template.

*new_name*
:   The new name of the text search template.

*new_schema*
:   The new schema for the text search template.


## Compatibility


 There is no `ALTER TEXT SEARCH TEMPLATE` statement in the SQL standard.


## See Also
  [sql-createtstemplate](create-text-search-template.md#sql-createtstemplate), [sql-droptstemplate](drop-text-search-template.md#sql-droptstemplate)
