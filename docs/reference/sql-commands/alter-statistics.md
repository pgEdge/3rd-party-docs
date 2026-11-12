<a id="sql-alterstatistics"></a>

# ALTER STATISTICS

change the definition of an extended statistics object

## Synopsis


```

ALTER STATISTICS NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER STATISTICS NAME RENAME TO NEW_NAME
ALTER STATISTICS NAME SET SCHEMA NEW_SCHEMA
ALTER STATISTICS NAME SET STATISTICS { NEW_TARGET | DEFAULT }
```


## Description


 `ALTER STATISTICS` changes the parameters of an existing extended statistics object. Any parameters not specifically set in the `ALTER STATISTICS` command retain their prior settings.


 You must own the statistics object to use `ALTER STATISTICS`. To change a statistics object's schema, you must also have `CREATE` privilege on the new schema. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the statistics object's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the statistics object. However, a superuser can alter ownership of any statistics object anyway.)


## Parameters


*name*
:   The name (optionally schema-qualified) of the statistics object to be altered.

*new_owner*
:   The user name of the new owner of the statistics object.

*new_name*
:   The new name for the statistics object.

*new_schema*
:   The new schema for the statistics object.

*new_target*
:   The statistic-gathering target for this statistics object for subsequent [`ANALYZE`](analyze.md#sql-analyze) operations. The target can be set in the range 0 to 10000. Set it to `DEFAULT` to revert to using the system default statistics target ([default_statistics_target](../../server-administration/server-configuration/query-planning.md#guc-default-statistics-target)). (Setting to a value of -1 is an obsolete way spelling to get the same outcome.) For more information on the use of statistics by the PostgreSQL query planner, refer to [Statistics Used by the Planner](../../the-sql-language/performance-tips/statistics-used-by-the-planner.md#planner-stats).


## Compatibility


 There is no `ALTER STATISTICS` command in the SQL standard.


## See Also
  [sql-createstatistics](create-statistics.md#sql-createstatistics), [sql-dropstatistics](drop-statistics.md#sql-dropstatistics)
