<a id="event-log-registration"></a>

## Registering Event Log on `Windows`


 To register a `Windows` event log library with the operating system, issue this command:

```

regsvr32 PGSQL_LIBRARY_DIRECTORY/pgevent.dll
```
 This creates registry entries used by the event viewer, under the default event source named `PostgreSQL`.


 To specify a different event source name (see [event_source](../server-configuration/error-reporting-and-logging.md#guc-event-source)), use the `/n` and `/i` options:

```

regsvr32 /n /i:EVENT_SOURCE_NAME PGSQL_LIBRARY_DIRECTORY/pgevent.dll
```


 To unregister the event log library from the operating system, issue this command:

```

regsvr32 /u [/i:EVENT_SOURCE_NAME] PGSQL_LIBRARY_DIRECTORY/pgevent.dll
```


!!! note

    To enable event logging in the database server, modify [log_destination](../server-configuration/error-reporting-and-logging.md#guc-log-destination) to include `eventlog` in `postgresql.conf`.
