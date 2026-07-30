<a id="sql-dropserver"></a>

# DROP SERVER

remove a foreign server descriptor

## Synopsis


```

DROP SERVER [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP SERVER` removes an existing foreign server descriptor. To execute this command, the current user must be the owner of the server.


## Parameters


`IF EXISTS`
:   Do not throw an error if the server does not exist. A notice is issued in this case.

*name*
:   The name of an existing server.

`CASCADE`
:   Automatically drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the server if any objects depend on it. This is the default.


## Examples


 Drop a server `foo` if it exists:

```sql

DROP SERVER IF EXISTS foo;
```


## Compatibility


 `DROP SERVER` conforms to ISO/IEC 9075-9 (SQL/MED). The `IF EXISTS` clause is a PostgreSQL extension.


## See Also
  [sql-createserver](create-server.md#sql-createserver), [sql-alterserver](alter-server.md#sql-alterserver)
