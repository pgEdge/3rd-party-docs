# Backup from a Standby
<a name="standby-backup"></a>

pgBackRest can perform backups on a standby instead of the primary. Standby backups require the `pg-standby` host to be configured and the `backup-standby` option enabled. If more than one standby is configured then the first running standby found will be used for the backup.

**Configure pg2-host/pg2-host-user and pg2-path**

```ini
[demo]
pg2-path=/var/lib/postgresql/17/demo
[demo]
pg2-host-type=tls
[demo]
pg2-host=pg-standby
[demo]
pg2-host-ca-file=/etc/pgbackrest/cert/ca.crt
[demo]
pg2-host-cert-file=/etc/pgbackrest/cert/client.crt
[demo]
pg2-host-key-file=/etc/pgbackrest/cert/client.key
[global]
backup-standby=y
```

Both the primary and standby databases are required to perform the backup, though the vast majority of the files will be copied from the standby to reduce load on the primary. The database hosts can be configured in any order. pgBackRest will automatically determine which is the primary and which is the standby.

**Backup the demo cluster from pg2**

```bash
pgbackrest --stanza=demo --log-level-console=detail backup
```

This incremental backup shows that most of the files are copied from the `pg-standby` host and only a few are copied from the `pg-primary` host.

pgBackRest creates a standby backup that is identical to a backup performed on the primary. It does this by starting/stopping the backup on the `pg-primary` host, copying only files that are replicated from the `pg-standby` host, then copying the remaining few files from the `pg-primary` host. This means that logs and statistics from the primary database will be included in the backup.
