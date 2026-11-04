# Parallel Backup / Restore
<a name="parallel-backup-restore"></a>

pgBackRest offers parallel processing to improve performance of compression and transfer. The number of processes to be used for this feature is set using the `--process-max` option.

It is usually best not to use more than 25% of available CPUs for the `backup` command. Backups don't have to run that fast as long as they are performed regularly and the backup process should not impact database performance, if at all possible.

The restore command can and should use all available CPUs because during a restore the PostgreSQL cluster is shut down and there is generally no other important work being done on the host. If the host contains multiple clusters then that should be considered when setting restore parallelism.

**Perform a backup with single process**

```bash
pgbackrest --stanza=demo --type=full backup
```

**Configure  to use multiple backup processes**

```ini
[global]
process-max=3
```

**Perform a backup with multiple processes**

```bash
pgbackrest --stanza=demo --type=full backup
```

**Get backup info for the demo cluster**

```bash
pgbackrest info
```

The performance of the last backup should be improved by using multiple processes. For very small backups the difference may not be very apparent, but as the size of the database increases so will time savings.
