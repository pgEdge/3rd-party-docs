<a id="sql-altertrigger"></a>

# ALTER TRIGGER

change the definition of a trigger

## Synopsis


```

ALTER TRIGGER NAME ON TABLE_NAME RENAME TO NEW_NAME
ALTER TRIGGER NAME ON TABLE_NAME [ NO ] DEPENDS ON EXTENSION EXTENSION_NAME
```


## Description


 `ALTER TRIGGER` changes properties of an existing trigger.


 The `RENAME` clause changes the name of the given trigger without otherwise changing the trigger definition. If the table that the trigger is on is a partitioned table, then corresponding clone triggers in the partitions are renamed too.


 The `DEPENDS ON EXTENSION` clause marks the trigger as dependent on an extension, such that if the extension is dropped, the trigger will automatically be dropped as well.


 You must own the table on which the trigger acts to be allowed to change its properties.


## Parameters


*name*
:   The name of an existing trigger to alter.

*table_name*
:   The name of the table on which this trigger acts.

*new_name*
:   The new name for the trigger.

*extension_name*
:   The name of the extension that the trigger is to depend on (or no longer dependent on, if `NO` is specified). A trigger that's marked as dependent on an extension is automatically dropped when the extension is dropped.


## Notes


 The ability to temporarily enable or disable a trigger is provided by [`ALTER TABLE`](alter-table.md#sql-altertable), not by `ALTER TRIGGER`, because `ALTER TRIGGER` has no convenient way to express the option of enabling or disabling all of a table's triggers at once.


## Examples


 To rename an existing trigger:

```sql

ALTER TRIGGER emp_stamp ON emp RENAME TO emp_track_chgs;
```


 To mark a trigger as being dependent on an extension:

```sql

ALTER TRIGGER emp_stamp ON emp DEPENDS ON EXTENSION emplib;
```


## Compatibility


 `ALTER TRIGGER` is a PostgreSQL extension of the SQL standard.


## See Also
  [sql-altertable](alter-table.md#sql-altertable)
