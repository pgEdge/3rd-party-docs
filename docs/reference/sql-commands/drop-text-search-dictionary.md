<a id="sql-droptsdictionary"></a>

# DROP TEXT SEARCH DICTIONARY

remove a text search dictionary

## Synopsis


```

DROP TEXT SEARCH DICTIONARY [ IF EXISTS ] NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP TEXT SEARCH DICTIONARY` drops an existing text search dictionary. To execute this command you must be the owner of the dictionary.


## Parameters


`IF EXISTS`
:   Do not throw an error if the text search dictionary does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of an existing text search dictionary.

`CASCADE`
:   Automatically drop objects that depend on the text search dictionary, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the text search dictionary if any objects depend on it. This is the default.


## Examples


 Remove the text search dictionary `english`:

```sql

DROP TEXT SEARCH DICTIONARY english;
```
 This command will not succeed if there are any existing text search configurations that use the dictionary. Add `CASCADE` to drop such configurations along with the dictionary.


## Compatibility


 There is no `DROP TEXT SEARCH DICTIONARY` statement in the SQL standard.


## See Also
  [sql-altertsdictionary](alter-text-search-dictionary.md#sql-altertsdictionary), [sql-createtsdictionary](create-text-search-dictionary.md#sql-createtsdictionary)
