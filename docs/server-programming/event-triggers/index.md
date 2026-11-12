<a id="event-triggers"></a>

# Event Triggers

 To supplement the trigger mechanism discussed in [Triggers](../triggers/index.md#triggers), PostgreSQL also provides event triggers. Unlike regular triggers, which are attached to a single table and capture only DML events, event triggers are global to a particular database and are capable of capturing DDL events.

 Like regular triggers, event triggers can be written in any procedural language that includes event trigger support, or in C, but not in plain SQL.

- [Overview of Event Trigger Behavior](overview-of-event-trigger-behavior.md#event-trigger-definition)
- [Writing Event Trigger Functions in C](writing-event-trigger-functions-in-c.md#event-trigger-interface)
- [A Complete Event Trigger Example](a-complete-event-trigger-example.md#event-trigger-example)
- [A Table Rewrite Event Trigger Example](a-table-rewrite-event-trigger-example.md#event-trigger-table-rewrite-example)
- [A Database Login Event Trigger Example](a-database-login-event-trigger-example.md#event-trigger-database-login-example)
