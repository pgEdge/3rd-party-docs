<a id="sql-droptstemplate"></a>

# DROP TEXT SEARCH TEMPLATE

remove a text search template

## Synopsis


```

DROP TEXT SEARCH TEMPLATE [ IF EXISTS ] NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP TEXT SEARCH TEMPLATE` drops an existing text search template. You must be a superuser to use this command.


## Parameters


`IF EXISTS`
:   Do not throw an error if the text search template does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of an existing text search template.

`CASCADE`
:   Automatically drop objects that depend on the text search template, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the text search template if any objects depend on it. This is the default.


## Examples


 Remove the text search template `thesaurus`:

```sql

DROP TEXT SEARCH TEMPLATE thesaurus;
```
 This command will not succeed if there are any existing text search dictionaries that use the template. Add `CASCADE` to drop such dictionaries along with the template.


## Compatibility


 There is no `DROP TEXT SEARCH TEMPLATE` statement in the SQL standard.


## See Also
  [sql-altertstemplate](alter-text-search-template.md#sql-altertstemplate), [sql-createtstemplate](create-text-search-template.md#sql-createtstemplate)
