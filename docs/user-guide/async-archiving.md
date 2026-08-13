# Asynchronous Archiving
<a name="async-archiving"></a>

Asynchronous archiving is enabled with the `archive-async` option. This option enables asynchronous operation for both the `archive-push` and `archive-get` commands.

A spool path is required. The commands will store transient data here but each command works quite a bit differently so spool path usage is described in detail in each section.

**Create the spool directory**

```bash
mkdir -p -m 750 /var/spool/pgbackrest
```

```bash
chown postgres:postgres /var/spool/pgbackrest
```

**Create the spool directory**

```bash
mkdir -p -m 750 /var/spool/pgbackrest
```

```bash
chown postgres:postgres /var/spool/pgbackrest
```

The spool path must be configured and asynchronous archiving enabled. Asynchronous archiving automatically confers some benefit by reducing the number of connections made to remote storage, but setting `process-max` can drastically improve performance by parallelizing operations. Be sure not to set `process-max` so high that it affects normal database operations.

**Configure the spool path and asynchronous archiving**

```ini
[global]
spool-path=/var/spool/pgbackrest
[global]
archive-async=y
[global:archive-push]
process-max=2
[global:archive-get]
process-max=2
```

**Configure the spool path and asynchronous archiving**

```ini
[global]
spool-path=/var/spool/pgbackrest
[global]
archive-async=y
[global:archive-push]
process-max=2
[global:archive-get]
process-max=2
```

!!! note

    `process-max` is configured using command sections so that the option is not used by backup and restore. This also allows different values for `archive-push` and `archive-get`.

For demonstration purposes streaming replication will be broken to force PostgreSQL to get WAL using the `restore_command`.

**Break streaming replication by changing the replication password**

```bash
psql -c "alter user replicator password 'bogus'"
```

**Restart standby to break connection**

```bash
pg_ctlcluster 17 demo restart
```

## Archive Push
<a name="async-archive-push"></a>

The asynchronous `archive-push` command offloads WAL archiving to a separate process (or processes) to improve throughput. It works by "looking ahead" to see which WAL segments are ready to be archived beyond the request that PostgreSQL is currently making via the `archive_command`. WAL segments are transferred to the archive directly from the `pg_xlog`/`pg_wal` directory and success is only returned by the `archive_command` when the WAL segment has been safely stored in the archive.

The spool path holds the current status of WAL archiving. Status files written into the spool directory are typically zero length and should consume a minimal amount of space (a few MB at most) and very little IO. All the information in this directory can be recreated so it is not necessary to preserve the spool directory if the cluster is moved to new hardware.

!!! important

    In the original implementation of asynchronous archiving, WAL segments were copied to the spool directory before compression and transfer. The new implementation copies WAL directly from the `pg_xlog` directory. If asynchronous archiving was utilized in v1.12 or prior, read the v1.13 release notes carefully before upgrading.

The `[stanza]-archive-push-async.log` file can be used to monitor the activity of the asynchronous process. A good way to test this is to quickly push a number of WAL segments.

**Test parallel asynchronous archiving**

```bash
psql -c "
                            select pg_create_restore_point('test async push'); select pg_switch_wal();
                            select pg_create_restore_point('test async push'); select pg_switch_wal();
                            select pg_create_restore_point('test async push'); select pg_switch_wal();
                            select pg_create_restore_point('test async push'); select pg_switch_wal();
                            select pg_create_restore_point('test async push'); select pg_switch_wal();"
```

```bash
pgbackrest --stanza=demo --log-level-console=info check
```

Now the log file will contain parallel, asynchronous activity.

**Check results in the log**

```bash
cat /var/log/pgbackrest/demo-archive-push-async.log
```

## Archive Get
<a name="async-archive-get"></a>

The asynchronous `archive-get` command maintains a local queue of WAL to improve throughput. If a WAL segment is not found in the queue it is fetched from the repository along with enough consecutive WAL to fill the queue. The maximum size of the queue is defined by `archive-get-queue-max`. Whenever the queue is less than half full more WAL will be fetched to fill it.

Asynchronous operation is most useful in environments that generate a lot of WAL or have a high latency connection to the repository storage (i.e., S3 or other object stores). In the case of a high latency connection it may be a good idea to increase `process-max`.

The `[stanza]-archive-get-async.log` file can be used to monitor the activity of the asynchronous process.

**Check results in the log**

```bash
cat /var/log/pgbackrest/demo-archive-get-async.log
```

**Fix streaming replication by changing the replication password**

```bash
psql -c "alter user replicator password 'jw8s0F4'"
```
