# Starting and Stopping
<a name="start-stop"></a>

If a standby is promoted for testing, or a test cluster is restored from a production backup, then it is a good idea to prevent those clusters from writing to pgBackRest repositories. This can be accomplished with the `stop` command.

The commands that write and are blocked by `stop` are: `archive-push`, `backup`, `expire`, `stanza-create`, and `stanza-upgrade`. Note that `stanza-delete` is an exception to this rule (see [Delete a Stanza](user-guide/delete-stanza.md#delete-stanza) for more details).

**Stop  write commands**

```bash
pgbackrest stop
```

New pgBackRest write commands will no longer run.

**Attempt a backup**

```bash
pgbackrest --stanza=demo backup
```

Specify the `--force` option to terminate any pgBackRest write commands that are currently running. This includes asynchronous archive-get (though it will run again if PostgreSQL requires it). If pgBackRest is already stopped then stopping again will generate a warning.

**Stop the  services again**

```bash
pgbackrest stop
```

Start pgBackRest write commands again with the `start` command. Write commands that were in progress before the stop will not automatically start again, but they are now allowed to start.

**Start  write commands**

```bash
pgbackrest start
```

It is also possible to stop pgBackRest for a single stanza.

**Stop  write commands for the demo stanza**

```bash
pgbackrest --stanza=demo stop
```

New pgBackRest write commands for the specified stanza will no longer run.

**Attempt a backup**

```bash
pgbackrest --stanza=demo backup
```

The stanza must also be specified when starting pgBackRest write commands for a single stanza.

**Start  write commands for the demo stanza**

```bash
pgbackrest --stanza=demo start
```
