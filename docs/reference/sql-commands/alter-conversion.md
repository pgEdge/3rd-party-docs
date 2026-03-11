<a id="sql-alterconversion"></a>

# ALTER CONVERSION

change the definition of a conversion

## Synopsis


```

ALTER CONVERSION NAME RENAME TO NEW_NAME
ALTER CONVERSION NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER CONVERSION NAME SET SCHEMA NEW_SCHEMA
```


## Description


 `ALTER CONVERSION` changes the definition of a conversion.


 You must own the conversion to use `ALTER CONVERSION`. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the conversion's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the conversion. However, a superuser can alter ownership of any conversion anyway.)


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing conversion.

*new_name*
:   The new name of the conversion.

*new_owner*
:   The new owner of the conversion.

*new_schema*
:   The new schema for the conversion.


## Examples


 To rename the conversion `iso_8859_1_to_utf8` to `latin1_to_unicode`:

```sql

ALTER CONVERSION iso_8859_1_to_utf8 RENAME TO latin1_to_unicode;
```


 To change the owner of the conversion `iso_8859_1_to_utf8` to `joe`:

```sql

ALTER CONVERSION iso_8859_1_to_utf8 OWNER TO joe;
```


## Compatibility


 There is no `ALTER CONVERSION` statement in the SQL standard.


## See Also
  [sql-createconversion](create-conversion.md#sql-createconversion), [sql-dropconversion](drop-conversion.md#sql-dropconversion)
