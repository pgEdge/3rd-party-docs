# Point-in-Time Recovery
<a name="pitr"></a>

[Restore a Backup](user-guide/repo-host.md#perform-restore) in [Quick Start](user-guide/quickstart.md#quickstart) performed default recovery, which is to play all the way to the end of the WAL stream. In the case of a hardware failure this is usually the best choice but for data corruption scenarios (whether machine or human in origin) Point-in-Time Recovery (PITR) is often more appropriate.

Point-in-Time Recovery (PITR) allows the WAL to be played from a backup to a specified lsn, time, transaction id, or recovery point. For common recovery scenarios time-based recovery is arguably the most useful. A typical recovery scenario is to restore a table that was accidentally dropped or data that was accidentally deleted. Recovering a dropped table is more dramatic so that's the example given here but deleted data would be recovered in exactly the same way.

**Create a table with very important data**

```bash
psql -c "begin;
                             create table important_table (message text);
                             insert into important_table values ('Important Data');
                             commit;
                             select * from important_table;"
```

It is important to represent the time as reckoned by PostgreSQL and to include timezone offsets. This reduces the possibility of unintended timezone conversions and an unexpected recovery result.

**Get the time from**

```bash
psql -Atc "select current_timestamp"
```

Now that the time has been recorded the table is dropped. In practice finding the exact time that the table was dropped is a lot harder than in this example. It may not be possible to find the exact time, but some forensic work should be able to get you close.

**Drop the important table**

```bash
psql -c "begin;
                                  drop table important_table;
                                  commit;
                                  select * from important_table;"
```

If the wrong backup is selected for restore then recovery to the required time target will fail. To demonstrate this a new incremental backup is performed where `important_table` does not exist.

**Perform an incremental backup**

```bash
pgbackrest --stanza=demo --type=incr backup
```

```bash
pgbackrest info
```

It will not be possible to recover the lost table from this backup since PostgreSQL can only play forward, not backward.

**Attempt recovery from an incorrect backup**

```bash
pg_ctlcluster 16 demo stop
```

```bash
pgbackrest --stanza=demo --delta
                     --set=<backup-last> --target-timeline=current
                     --type=time "--target=<time-recovery-timestamp>" --target-action=promote restore
```

```bash
pg_ctlcluster 16 demo start
```

```bash
pg_ctlcluster 16 demo start
```

```bash
cat /var/log/postgresql/postgresql-16-demo.log
```

```bash
pg_ctlcluster 16 demo start
```

```bash
psql -c "select * from important_table"
```

Looking at the log output it's not obvious that recovery failed to restore the table. The key is to look for the presence of the "recovery stopping before..." and "last completed transaction..." log messages. If they are not present then the recovery to the specified point-in-time was not successful.

**Examine the  log output to discover the recovery was not successful**

```bash
cat /var/log/postgresql/postgresql-16-demo.log
```

A reliable method is to allow pgBackRest to automatically select a backup capable of recovery to the time target, i.e. a backup that ended before the specified time.

!!! note

    pgBackRest cannot automatically select a backup when the restore type is `xid` or `name`.

**Restore the demo cluster to <time-recovery-timestamp>**

```bash
pg_ctlcluster 16 demo stop
```

```bash
pgbackrest --stanza=demo --delta
                    --type=time "--target=<time-recovery-timestamp>"
                    --target-action=promote restore
```

```bash
cat /var/lib/postgresql/16/demo/postgresql.auto.conf
```

pgBackRest has generated the recovery settings in `postgresql.auto.conf` so PostgreSQL can be started immediately. `%f` is how PostgreSQL specifies the WAL segment it needs and `%p` is the location where it should be copied. Once PostgreSQL has finished recovery the table will exist again and can be queried.

**Start  and check that the important table exists**

```bash
pg_ctlcluster 16 demo start
```

```bash
psql -c "select * from important_table"
```

The PostgreSQL log also contains valuable information. It will indicate the time and transaction where the recovery stopped and also give the time of the last transaction to be applied.

**Examine the  log output**

```bash
cat /var/log/postgresql/postgresql-16-demo.log
```
