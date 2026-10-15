<a id="plpython-event-trigger"></a>

## Event Trigger Functions


 PL/Python can be used to define event triggers (see also [Event Triggers](../event-triggers/index.md#event-triggers)). PostgreSQL requires that a function that is to be called as an event trigger must be declared as a function with no arguments and a return type of `event_trigger`.


 When a function is used as an event trigger, the dictionary `TD` contains trigger-related values:

`TD["event"]`
:   The event the trigger was fired for, as a string, for example `ddl_command_start`.

`TD["tag"]`
:   The command tag for which the trigger was fired, as a string, for example `DROP TABLE`.


 [A PL/Python Event Trigger Function](#plpython-event-trigger-example) shows an example of an event trigger function in PL/Python.
 <a id="plpython-event-trigger-example"></a>

**Example: A PL/Python Event Trigger Function**


 This example trigger simply raises a `NOTICE` message each time a supported command is executed.


```sql

CREATE OR REPLACE FUNCTION pysnitch() RETURNS event_trigger
LANGUAGE plpython3u
AS $$
  plpy.notice("TD[event] => " + TD["event"] + " ; TD[tag] => " + TD["tag"]);
$$;

CREATE EVENT TRIGGER pysnitch ON ddl_command_start EXECUTE FUNCTION pysnitch();
```
