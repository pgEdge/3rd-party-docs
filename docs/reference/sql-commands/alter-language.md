<a id="sql-alterlanguage"></a>

# ALTER LANGUAGE

change the definition of a procedural language

## Synopsis


```

ALTER [ PROCEDURAL ] LANGUAGE NAME RENAME TO NEW_NAME
ALTER [ PROCEDURAL ] LANGUAGE NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
```


## Description


 `ALTER LANGUAGE` changes the definition of a procedural language. The only functionality is to rename the language or assign a new owner. You must be superuser or owner of the language to use `ALTER LANGUAGE`.


## Parameters


*name*
:   Name of a language

*new_name*
:   The new name of the language

*new_owner*
:   The new owner of the language


## Compatibility


 There is no `ALTER LANGUAGE` statement in the SQL standard.


## See Also
  [sql-createlanguage](create-language.md#sql-createlanguage), [sql-droplanguage](drop-language.md#sql-droplanguage)
