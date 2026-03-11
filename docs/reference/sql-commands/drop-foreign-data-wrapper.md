<a id="sql-dropforeigndatawrapper"></a>

# DROP FOREIGN DATA WRAPPER

remove a foreign-data wrapper

## Synopsis


```

DROP FOREIGN DATA WRAPPER [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP FOREIGN DATA WRAPPER` removes an existing foreign-data wrapper. To execute this command, the current user must be the owner of the foreign-data wrapper.


## Parameters


`IF EXISTS`
:   Do not throw an error if the foreign-data wrapper does not exist. A notice is issued in this case.

*name*
:   The name of an existing foreign-data wrapper.

`CASCADE`
:   Automatically drop objects that depend on the foreign-data wrapper (such as foreign tables and servers), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the foreign-data wrapper if any objects depend on it. This is the default.


## Examples


 Drop the foreign-data wrapper `dbi`:

```sql

DROP FOREIGN DATA WRAPPER dbi;
```


## Compatibility


 `DROP FOREIGN DATA WRAPPER` conforms to ISO/IEC 9075-9 (SQL/MED). The `IF EXISTS` clause is a PostgreSQL extension.


## See Also
  [sql-createforeigndatawrapper](create-foreign-data-wrapper.md#sql-createforeigndatawrapper), [sql-alterforeigndatawrapper](alter-foreign-data-wrapper.md#sql-alterforeigndatawrapper)
