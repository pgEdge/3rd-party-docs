<a id="sql-altereventtrigger"></a>

# ALTER EVENT TRIGGER

change the definition of an event trigger

## Synopsis


```

ALTER EVENT TRIGGER NAME DISABLE
ALTER EVENT TRIGGER NAME ENABLE [ REPLICA | ALWAYS ]
ALTER EVENT TRIGGER NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER EVENT TRIGGER NAME RENAME TO NEW_NAME
```


## Description


 `ALTER EVENT TRIGGER` changes properties of an existing event trigger.


 You must be superuser to alter an event trigger.


## Parameters


*name*
:   The name of an existing trigger to alter.

*new_owner*
:   The user name of the new owner of the event trigger.

*new_name*
:   The new name of the event trigger.

`DISABLE`/`ENABLE [ REPLICA | ALWAYS ]`
:   These forms configure the firing of event triggers. A disabled trigger is still known to the system, but is not executed when its triggering event occurs. See also [session_replication_role](../../server-administration/server-configuration/client-connection-defaults.md#guc-session-replication-role).
 <a id="sql-alterventtrigger-compatibility"></a>

## Compatibility


 There is no `ALTER EVENT TRIGGER` statement in the SQL standard.


## See Also
  [sql-createeventtrigger](create-event-trigger.md#sql-createeventtrigger), [sql-dropeventtrigger](drop-event-trigger.md#sql-dropeventtrigger)
