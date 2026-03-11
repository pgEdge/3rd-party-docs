<a id="sql-dropforeigntable"></a>

# DROP FOREIGN TABLE

remove a foreign table

## Synopsis


```

DROP FOREIGN TABLE [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP FOREIGN TABLE` removes a foreign table. Only the owner of a foreign table can remove it.


## Parameters


`IF EXISTS`
:   Do not throw an error if the foreign table does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of the foreign table to drop.

`CASCADE`
:   Automatically drop objects that depend on the foreign table (such as views), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the foreign table if any objects depend on it. This is the default.


## Examples


 To destroy two foreign tables, `films` and `distributors`:

```sql

DROP FOREIGN TABLE films, distributors;
```


## Compatibility


 This command conforms to ISO/IEC 9075-9 (SQL/MED), except that the standard only allows one foreign table to be dropped per command, and apart from the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also
  [sql-alterforeigntable](alter-foreign-table.md#sql-alterforeigntable), [sql-createforeigntable](create-foreign-table.md#sql-createforeigntable)
