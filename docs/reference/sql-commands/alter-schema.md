<a id="sql-alterschema"></a>

# ALTER SCHEMA

change the definition of a schema

## Synopsis


```

ALTER SCHEMA NAME RENAME TO NEW_NAME
ALTER SCHEMA NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
```


## Description


 `ALTER SCHEMA` changes the definition of a schema.


 You must own the schema to use `ALTER SCHEMA`. To rename a schema you must also have the `CREATE` privilege for the database. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have the `CREATE` privilege for the database. (Note that superusers have all these privileges automatically.)


## Parameters


*name*
:   The name of an existing schema.

*new_name*
:   The new name of the schema. The new name cannot begin with `pg_`, as such names are reserved for system schemas.

*new_owner*
:   The new owner of the schema.


## Compatibility


 There is no `ALTER SCHEMA` statement in the SQL standard.


## See Also
  [sql-createschema](create-schema.md#sql-createschema), [sql-dropschema](drop-schema.md#sql-dropschema)
