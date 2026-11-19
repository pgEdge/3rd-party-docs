<a id="sql-dropschema"></a>

# DROP SCHEMA

remove a schema

## Synopsis


```

DROP SCHEMA [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP SCHEMA` removes schemas from the database.


 A schema can only be dropped by its owner or a superuser. Note that the owner can drop the schema (and thereby all contained objects) even if they do not own some of the objects within the schema.


## Parameters


`IF EXISTS`
:   Do not throw an error if the schema does not exist. A notice is issued in this case.

*name*
:   The name of a schema.

`CASCADE`
:   Automatically drop objects (tables, functions, etc.) that are contained in the schema, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the schema if it contains any objects. This is the default.


## Notes


 Using the `CASCADE` option might make the command remove objects in other schemas besides the one(s) named.


## Examples


 To remove schema `mystuff` from the database, along with everything it contains:

```sql

DROP SCHEMA mystuff CASCADE;
```


## Compatibility


 `DROP SCHEMA` is fully conforming with the SQL standard, except that the standard only allows one schema to be dropped per command, and apart from the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also
  [sql-alterschema](alter-schema.md#sql-alterschema), [sql-createschema](create-schema.md#sql-createschema)
