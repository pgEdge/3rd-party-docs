# Stress Testing
<a name="stress"></a>

## Configuration
<a name="configuration"></a>

**Configure repository for stress testing**

```ini
[global]
process-max=8
[global]
compress-type=lz4
[global]
compress-level=1
[global]
repo1-retention-full=1
[global]
log-timestamp=y
```

**Create the pg-primary spool directory**

```bash
mkdir -p -m 750 /var/spool/pgbackrest
```

```bash
chown postgres:postgres /var/spool/pgbackrest
```

**Configure pg-primary for stress testing**

```ini
[global]
process-max=8
[global]
log-timestamp=y
[global]
compress-type=lz4
[global]
compress-level=1
[global]
spool-path=/var/spool/pgbackrest
[global]
archive-async=y
[global:archive-push]
process-max=4
[global:archive-get]
process-max=4
```

**Create the pg-standby spool directory**

```bash
mkdir -p -m 750 /var/spool/pgbackrest
```

```bash
chown postgres:postgres /var/spool/pgbackrest
```

**Configure pg-standby for stress testing**

```ini
[global]
process-max=8
[global]
log-timestamp=y
[global]
compress-type=lz4
[global]
compress-level=1
[global]
spool-path=/var/spool/pgbackrest
[global]
archive-async=y
[global:archive-push]
process-max=4
[global:archive-get]
process-max=4
```

## Create Tables and Load Data
<a name="data-load"></a>

### Break Streaming Replication
<a name="streaming-break"></a>

Break streaming replication to force the standby to replicate from the archive during data load.

**Break streaming replication by changing the replication password**

```bash
psql -c "alter user replicator password 'bogus'"
```

**Restart standby to break connection**

```bash
pg_ctlcluster 17 demo restart
```

### Create Tables
<a name="table-create"></a>

**Create tables**

```bash
bash -c 'for i in {1..1};
                                do psql -c "select create_test_table(${i?}, 1000, true)";
                                done'
```

### Load Data
<a name="data-load"></a>

**Load data**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -i -s 1
```

### Fix Streaming Replication
<a name="streaming-fix"></a>

Fix streaming replication so backups will work. Note that streaming replication will not start again until all WAL in the archive has been exhausted.

**Fix streaming replication by changing the replication password**

```bash
psql -c "alter user replicator password 'jw8s0F4'"
```

## Testing
<a name="test"></a>

### Full Backup
<a name="backup-full"></a>

**Full backup**

```bash
pgbackrest --stanza=demo --type=full
                            --log-level-console=info --log-level-file=detail backup
```

### Diff Backup with Delta and Block Incremental
<a name="backup-diff-delta"></a>

**Database updates**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Diff backup**

```bash
pgbackrest --stanza=demo --type=diff --delta --repo1-bundle
                            --repo1-block --log-level-console=info --log-level-file=detail backup
```

### Incr Backup with Block Incremental
<a name="backup-incr"></a>

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

**Database updates <stress-update-incr-count>**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Incr backup <stress-update-incr-count>**

```bash
pgbackrest --stanza=demo --type=incr --repo1-bundle
                    --repo1-block --log-level-console=info backup
```

### Restore with Delta
<a name="restore-delta"></a>

**Database updates so delta has something to restore**

```bash
/usr/lib/postgresql/17/bin/pgbench -n -b simple-update -t 1
```

**Stop**

```bash
pg_ctlcluster 17 demo stop
```

**Restore**

```bash
pgbackrest --stanza=demo --type=standby --delta
                            --log-level-console=info --log-level-file=detail restore
```

### Restore
<a name="restore"></a>

**Remove data**

```bash
rm -rf /var/lib/postgresql/17/demo
```

**Restore**

```bash
pgbackrest --stanza=demo --type=standby
                            --log-level-console=info --log-level-file=detail restore
```

**Start**

```bash
pg_ctlcluster 17 demo start
```

**Check cluster**

```bash
psql -c "select count(*) from pg_class"
```
