# Retention
<a name="retention"></a>

Generally it is best to retain as many backups as possible to provide a greater window for [Point-in-Time Recovery](pitr.md#pitr), but practical concerns such as disk space must also be considered. Retention options remove older backups once they are no longer needed.

`expire`

## Full Backup Retention
<a name="full"></a>

The `repo1-retention-full-type` determines how the option `repo1-retention-full` is interpreted; either as the count of full backups to be retained or how many days to retain full backups. New backups must be completed before expiration will occur &mdash; that means if `repo1-retention-full-type=count` and `repo1-retention-full=2` then there will be three full backups stored before the oldest one is expired, or if `repo1-retention-full-type=time` and `repo1-retention-full=20` then there must be one full backup that is at least 20 days old before expiration can occur.

**Configure repo1-retention-full**

```ini
[global]
repo1-retention-full=2
```

Backup `repo1-retention-full=2` but currently there is only one full backup so the next full backup to run will not expire any full backups.

**Perform a full backup**

```bash
pgbackrest --stanza=demo --type=full
                        --log-level-console=detail backup
```

Archive *is* expired because WAL segments were generated before the oldest backup. These are not useful for recovery &mdash; only WAL segments generated after a backup can be used to recover that backup.

**Perform a full backup**

```bash
pgbackrest --stanza=demo --type=full
                        --log-level-console=info backup
```

The `<backup-full-first>` full backup is expired and archive retention is based on the `<backup-full-second>` which is now the oldest full backup.

## Differential Backup Retention
<a name="diff"></a>

Set `repo1-retention-diff` to the number of differential backups required. Differentials only rely on the prior full backup so it is possible to create a "rolling" set of differentials for the last day or more. This allows quick restores to recent points-in-time but reduces overall space consumption.

**Configure repo1-retention-diff**

```ini
[global]
repo1-retention-diff=1
```

Backup `repo1-retention-diff=1` so two differentials will need to be performed before one is expired. An incremental backup is added to demonstrate incremental expiration, which in this case depends on the differential expiration.

**Perform differential and incremental backups**

```bash
pgbackrest --stanza=demo --type=diff backup
```

```bash
pgbackrest --stanza=demo --type=incr backup
```

Now performing a differential backup will expire the previous differential and incremental backups leaving only one differential backup.

**Perform a differential backup**

```bash
pgbackrest --stanza=demo --type=diff
                        --log-level-console=info backup
```

## Archive Retention
<a name="archive"></a>

Although pgBackRest automatically removes archived WAL segments when expiring backups (the default expires WAL for full backups based on the `repo1-retention-full` option), it may be useful to expire archive more aggressively to save disk space. Note that full backups are treated as differential backups for the purpose of differential archive retention.

Expiring archive will never remove WAL segments that are required to make a backup consistent. However, since Point-in-Time-Recovery (PITR) only works on a continuous WAL stream, care should be taken when aggressively expiring archive outside of the normal backup expiration process. To determine what will be expired without actually expiring anything, the `dry-run` option can be provided on the command line with the `expire` command.

**Configure repo1-retention-diff**

```ini
[global]
repo1-retention-diff=2
```

**Perform differential backup**

```bash
pgbackrest --stanza=demo --type=diff
                        --log-level-console=info backup
```

**Expire archive**

```bash
pgbackrest --stanza=demo --log-level-console=detail
                        --repo1-retention-archive-type=diff --repo1-retention-archive=1 expire
```

The `<backup-diff-first>` differential backup has archived WAL segments that must be retained to make the older backups consistent even though they cannot be played any further forward with PITR. WAL segments generated after `<backup-diff-first>` but before `<backup-diff-second>` are removed. WAL segments generated after the new backup `<backup-diff-second>` remain and can be used for PITR.

Since full backups are considered differential backups for the purpose of differential archive retention, if a full backup is now performed with the same settings, only the archive for that full backup is retained for PITR.
