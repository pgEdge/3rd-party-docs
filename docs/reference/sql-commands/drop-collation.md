# DROP COLLATION { #sql-dropcollation }

remove a collation

## Synopsis


```

DROP COLLATION [ IF EXISTS ] NAME [ CASCADE | RESTRICT ]
```


## Description { #sql-dropcollation-description }


 `DROP COLLATION` removes a previously defined collation. To be able to drop a collation, you must own the collation.


## Parameters


`IF EXISTS`
:   Do not throw an error if the collation does not exist. A notice is issued in this case.

*name*
:   The name of the collation. The collation name can be schema-qualified.

`CASCADE`
:   Automatically drop objects that depend on the collation, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the collation if any objects depend on it. This is the default.


## Examples { #sql-dropcollation-examples }


 To drop the collation named `german`:

```sql

DROP COLLATION german;
```


## Compatibility { #sql-dropcollation-compat }


 The `DROP COLLATION` command conforms to the SQL standard, apart from the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also
  [sql-altercollation](alter-collation.md#sql-altercollation), [sql-createcollation](create-collation.md#sql-createcollation)
