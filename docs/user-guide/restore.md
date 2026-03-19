# Restore
<a name="restore"></a>
`restore`

The following sections introduce additional `restore` command features.

## File Ownership
<a name="ownership"></a>

If a `restore` is run as a non-root user (the typical scenario) then all files restored will belong to the user/group executing pgBackRest. If existing files are not owned by the executing user/group then an error will result if the ownership cannot be updated to the executing user/group. In that case the file ownership will need to be updated by a privileged user before the restore can be retried.

If a `restore` is run as the `root` user then pgBackRest will attempt to recreate the ownership recorded in the manifest when the backup was made. Only user/group **names** are stored in the manifest so the same names must exist on the restore host for this to work. If the user/group name cannot be found locally then the user/group of the PostgreSQL data directory will be used and finally `root` if the data directory user/group cannot be mapped to a name.

## Delta Option
<a name="option-delta"></a>

[Restore a Backup](user-guide/repo-host.md#perform-restore) in [Quick Start](user-guide/quickstart.md#quickstart) required the database cluster directory to be cleaned before the `restore` could be performed. The `delta` option allows pgBackRest to automatically determine which files in the database cluster directory can be preserved and which ones need to be restored from the backup &mdash; it also *removes* files not present in the backup manifest so it will dispose of divergent changes. This is accomplished by calculating a [SHA-1](https://en.wikipedia.org/wiki/SHA-1) cryptographic hash for each file in the database cluster directory. If the `SHA-1` hash does not match the hash stored in the backup then that file will be restored. This operation is very efficient when combined with the `process-max` option. Since the PostgreSQL server is shut down during the restore, a larger number of processes can be used than might be desirable during a backup when the PostgreSQL server is running.

**Stop the demo cluster, perform delta restore**

```bash
pg_ctlcluster 16 demo stop
```

```bash
pgbackrest --stanza=demo --delta
                        --log-level-console=detail restore
```

**Restart**

```bash
pg_ctlcluster 16 demo start
```

## Restore Selected Databases
<a name="option-db-include"></a>

There may be cases where it is desirable to selectively restore specific databases from a cluster backup. This could be done for performance reasons or to move selected databases to a machine that does not have enough space to restore the entire cluster backup.

To demonstrate this feature two databases are created: test1 and test2.

**Create two test databases**

```bash
psql -c "create database test1;"
```

```bash
psql -c "create database test2;"
```

Each test database will be seeded with tables and data to demonstrate that recovery works with selective restore.

**Create a test table in each database**

```bash
psql -c "create table test1_table (id int);
                                 insert into test1_table (id) values (1);" test1
```

```bash
psql -c "create table test2_table (id int);
                                 insert into test2_table (id) values (2);" test2
```

A fresh backup is run so pgBackRest is aware of the new databases.

**Perform a backup**

```bash
pgbackrest --stanza=demo --type=incr backup
```

One of the main reasons to use selective restore is to save space. The size of the test1 database is shown here so it can be compared with the disk utilization after a selective restore.

**Show space used by test1 database**

```bash
du -sh /var/lib/postgresql/16/demo/base/<database-test1-oid>
```

If the database to restore is not known, use the `info` command `set` option to discover databases that are part of the backup set.

**Show database list for backup**

```bash
pgbackrest --stanza=demo
                        --set=<backup-last-incr> info
```

Stop the cluster and restore only the test2 database. Built-in databases (`template0`, `template1`, and `postgres`) are always restored.

!!! warning

    Recovery may error unless `--type=immediate` is specified. This is because after consistency is reached PostgreSQL will flag zeroed pages as errors even for a full-page write. For PostgreSQL &ge; 13 the `ignore_invalid_pages` setting may be used to ignore invalid pages. In this case it is important to check the logs after recovery to ensure that no invalid pages were reported in the selected databases.

**Restore from last backup including only the test2 database**

```bash
pg_ctlcluster 16 demo stop
```

```bash
pgbackrest --stanza=demo --delta
                        --db-include=test2 --type=immediate --target-action=promote restore
```

```bash
pg_ctlcluster 16 demo start
```

Once recovery is complete the test2 database will contain all previously created tables and data.

**Demonstrate that the test2 database was recovered**

```bash
psql -c "select * from test2_table;" test2
```

The test1 database, despite successful recovery, is not accessible. This is because the entire database was restored as sparse, zeroed files. PostgreSQL can successfully apply WAL on the zeroed files but the database as a whole will not be valid because key files contain no data. This is purposeful to prevent the database from being accidentally used when it might contain partial data that was applied during WAL replay.

**Attempting to connect to the test1 database will produce an error**

```bash
psql -c "select * from test1_table;" test1
```

Since the test1 database is restored with sparse, zeroed files it will only require as much space as the amount of WAL that is written during recovery. While the amount of WAL generated during a backup and applied during recovery can be significant it will generally be a small fraction of the total database size, especially for large databases where this feature is most likely to be useful.

It is clear that the test1 database uses far less disk space during the selective restore than it would have if the entire database had been restored.

**Show space used by test1 database after recovery**

```bash
du -sh /var/lib/postgresql/16/demo/base/<database-test1-oid>
```

At this point the only action that can be taken on the invalid test1 database is `drop database`. pgBackRest does not automatically drop the database since this cannot be done until recovery is complete and the cluster is accessible.

**Drop the test1 database**

```bash
psql -c "drop database test1;"
```

Now that the invalid test1 database has been dropped only the test2 and built-in databases remain.

**List remaining databases**

```bash
psql -c "select oid, datname from pg_database order by oid;"
```
