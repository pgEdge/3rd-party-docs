<a id="sql-droptable"></a>

# DROP TABLE

remove a table

## Synopsis


```

DROP TABLE [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP TABLE` removes tables from the database. Only the table owner, the schema owner, and superuser can drop a table. To empty a table of rows without destroying the table, use [`DELETE`](delete.md#sql-delete) or [`TRUNCATE`](truncate.md#sql-truncate).


 `DROP TABLE` always removes any indexes, rules, triggers, and constraints that exist for the target table. However, to drop a table that is referenced by a view or a foreign-key constraint of another table, `CASCADE` must be specified. (`CASCADE` will remove a dependent view entirely, but in the foreign-key case it will only remove the foreign-key constraint, not the other table entirely.)


## Parameters


`IF EXISTS`
:   Do not throw an error if the table does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of the table to drop.

`CASCADE`
:   Automatically drop objects that depend on the table (such as views), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the table if any objects depend on it. This is the default.


## Examples


 To destroy two tables, `films` and `distributors`:

```sql

DROP TABLE films, distributors;
```


## Compatibility


 This command conforms to the SQL standard, except that the standard only allows one table to be dropped per command, and apart from the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also
  [sql-altertable](alter-table.md#sql-altertable), [sql-createtable](create-table.md#sql-createtable)
