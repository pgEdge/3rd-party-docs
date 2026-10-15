<a id="sql-droptype"></a>

# DROP TYPE

remove a data type

## Synopsis


```

DROP TYPE [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP TYPE` removes a user-defined data type. Only the owner of a type can remove it.


## Parameters


`IF EXISTS`
:   Do not throw an error if the type does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of the data type to remove.

`CASCADE`
:   Automatically drop objects that depend on the type (such as table columns, functions, and operators), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the type if any objects depend on it. This is the default.
 <a id="sql-droptype-examples"></a>

## Examples


 To remove the data type `box`:

```sql

DROP TYPE box;
```
 <a id="sql-droptype-compatibility"></a>

## Compatibility


 This command is similar to the corresponding command in the SQL standard, apart from the `IF EXISTS` option, which is a PostgreSQL extension. But note that much of the `CREATE TYPE` command and the data type extension mechanisms in PostgreSQL differ from the SQL standard.
 <a id="sql-droptype-see-also"></a>

## See Also
  [sql-altertype](alter-type.md#sql-altertype), [sql-createtype](create-type.md#sql-createtype)
