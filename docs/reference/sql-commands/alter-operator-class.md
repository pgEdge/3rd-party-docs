<a id="sql-alteropclass"></a>

# ALTER OPERATOR CLASS

change the definition of an operator class

## Synopsis


```

ALTER OPERATOR CLASS NAME USING INDEX_METHOD
    RENAME TO NEW_NAME

ALTER OPERATOR CLASS NAME USING INDEX_METHOD
    OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }

ALTER OPERATOR CLASS NAME USING INDEX_METHOD
    SET SCHEMA NEW_SCHEMA
```


## Description


 `ALTER OPERATOR CLASS` changes the definition of an operator class.


 You must own the operator class to use `ALTER OPERATOR CLASS`. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the operator class's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the operator class. However, a superuser can alter ownership of any operator class anyway.)


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing operator class.

*index_method*
:   The name of the index method this operator class is for.

*new_name*
:   The new name of the operator class.

*new_owner*
:   The new owner of the operator class.

*new_schema*
:   The new schema for the operator class.


## Compatibility


 There is no `ALTER OPERATOR CLASS` statement in the SQL standard.


## See Also
  [sql-createopclass](create-operator-class.md#sql-createopclass), [sql-dropopclass](drop-operator-class.md#sql-dropopclass), [sql-alteropfamily](alter-operator-family.md#sql-alteropfamily)
