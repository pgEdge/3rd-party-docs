<a id="sql-dropeventtrigger"></a>

# DROP EVENT TRIGGER

remove an event trigger

## Synopsis


```

DROP EVENT TRIGGER [ IF EXISTS ] NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP EVENT TRIGGER` removes an existing event trigger. To execute this command, the current user must be the owner of the event trigger.


## Parameters


`IF EXISTS`
:   Do not throw an error if the event trigger does not exist. A notice is issued in this case.

*name*
:   The name of the event trigger to remove.

`CASCADE`
:   Automatically drop objects that depend on the trigger, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the trigger if any objects depend on it. This is the default.
 <a id="sql-dropeventtrigger-examples"></a>

## Examples


 Destroy the trigger `snitch`:

```sql

DROP EVENT TRIGGER snitch;
```
 <a id="sql-dropeventtrigger-compatibility"></a>

## Compatibility


 There is no `DROP EVENT TRIGGER` statement in the SQL standard.


## See Also
  [sql-createeventtrigger](create-event-trigger.md#sql-createeventtrigger), [sql-altereventtrigger](alter-event-trigger.md#sql-altereventtrigger)
