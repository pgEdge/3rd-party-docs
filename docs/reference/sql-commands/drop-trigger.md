<a id="sql-droptrigger"></a>

# DROP TRIGGER

remove a trigger

## Synopsis


```

DROP TRIGGER [ IF EXISTS ] NAME ON TABLE_NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP TRIGGER` removes an existing trigger definition. To execute this command, the current user must be the owner of the table for which the trigger is defined.


## Parameters


`IF EXISTS`
:   Do not throw an error if the trigger does not exist. A notice is issued in this case.

*name*
:   The name of the trigger to remove.

*table_name*
:   The name (optionally schema-qualified) of the table for which the trigger is defined.

`CASCADE`
:   Automatically drop objects that depend on the trigger, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the trigger if any objects depend on it. This is the default.
 <a id="sql-droptrigger-examples"></a>

## Examples


 Destroy the trigger `if_dist_exists` on the table `films`:

```sql

DROP TRIGGER if_dist_exists ON films;
```
 <a id="sql-droptrigger-compatibility"></a>

## Compatibility


 The `DROP TRIGGER` statement in PostgreSQL is incompatible with the SQL standard. In the SQL standard, trigger names are not local to tables, so the command is simply <code>DROP TRIGGER
   </code><em>name</em>.


## See Also
  [sql-createtrigger](create-trigger.md#sql-createtrigger)
