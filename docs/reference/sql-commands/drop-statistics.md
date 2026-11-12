<a id="sql-dropstatistics"></a>

# DROP STATISTICS

remove extended statistics

## Synopsis


```

DROP STATISTICS [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP STATISTICS` removes statistics object(s) from the database. Only the statistics object's owner, the schema owner, or a superuser can drop a statistics object.


## Parameters


`IF EXISTS`
:   Do not throw an error if the statistics object does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of the statistics object to drop.

`CASCADE`, `RESTRICT`
:   These key words do not have any effect, since there are no dependencies on statistics.


## Examples


 To destroy two statistics objects in different schemas, without failing if they don't exist:

```sql

DROP STATISTICS IF EXISTS
    accounting.users_uid_creation,
    public.grants_user_role;
```


## Compatibility


 There is no `DROP STATISTICS` command in the SQL standard.


## See Also
  [sql-alterstatistics](alter-statistics.md#sql-alterstatistics), [sql-createstatistics](create-statistics.md#sql-createstatistics)
