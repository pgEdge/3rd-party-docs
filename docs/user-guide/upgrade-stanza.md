# Upgrading
<a name="upgrade-stanza"></a>
`stanza-upgrade`

The following instructions are not meant to be a comprehensive guide for upgrading PostgreSQL, rather they outline the general process for upgrading a primary and standby with the intent of demonstrating the steps required to reconfigure pgBackRest. It is recommended that a backup be taken prior to upgrading.

**Stop old cluster**

```bash
pg_ctlcluster 17 demo stop
```

Stop the old cluster on the standby since it will be restored from the newly upgraded cluster.

**Stop old cluster**

```bash
pg_ctlcluster 17 demo stop
```

Create the new cluster and perform upgrade.

**Create new cluster and perform the upgrade**

```bash
/usr/lib/postgresql/18/bin/initdb
                        -D /var/lib/postgresql/18/demo -k -A peer
```

```bash
pg_createcluster 18 demo
```

```bash
sh -c 'cd /var/lib/postgresql &&
                        /usr/lib/postgresql/18/bin/pg_upgrade
                            --old-bindir=/usr/lib/postgresql/17/bin
                            --new-bindir=/usr/lib/postgresql/18/bin
                            --old-datadir=/var/lib/postgresql/17/demo
                            --new-datadir=/var/lib/postgresql/18/demo
                            --old-options=" -c config_file=/etc/postgresql/17/demo/postgresql.conf"
                            --new-options=" -c config_file=/etc/postgresql/18/demo/postgresql.conf"'
```

```bash
sh -c 'cd /var/lib/pgsql &&
                        /usr/pgsql-18/bin/pg_upgrade
                            --old-bindir=/usr/pgsql-17/bin
                            --new-bindir=/usr/pgsql-18/bin
                            --old-datadir=/var/lib/postgresql/17/demo
                            --new-datadir=/var/lib/postgresql/18/demo
                            --old-options=" -c config_file=/etc/postgresql/17/demo/postgresql.conf"
                            --new-options=" -c config_file=/etc/postgresql/18/demo/postgresql.conf"'
```

Configure the new cluster settings and port.

**Configure**

```ini
archive_command = 'pgbackrest --stanza=demo archive-push %p'
archive_mode = on
log_filename = 'postgresql.log'
```

Update the pgBackRest configuration on all systems to point to the new cluster.

**Upgrade the pg1-path**

```ini
[demo]
pg1-path=/var/lib/postgresql/18/demo
```

**Upgrade the pg-path**

```ini
[demo]
pg1-path=/var/lib/postgresql/18/demo
```

**Upgrade pg1-path and pg2-path, disable backup from standby**

```ini
[demo]
pg1-path=/var/lib/postgresql/18/demo
[demo]
pg2-path=/var/lib/postgresql/18/demo
[global]
backup-standby=n
```

**Copy hba configuration**

```bash
cp /etc/postgresql/17/demo/pg_hba.conf
                    /etc/postgresql/18/demo/pg_hba.conf
```

Before starting the new cluster, the `stanza-upgrade` command must be run.

**Upgrade the stanza**

```bash
pgbackrest --stanza=demo --no-online
                    --log-level-console=info stanza-upgrade
```

Start the new cluster and confirm it is successfully installed.

**Start new cluster**

```bash
pg_ctlcluster 18 demo start
```

Test configuration using the `check` command.

**Check configuration**

```bash
pg_lsclusters
```

```bash
pgbackrest --stanza=demo check
```

Remove the old cluster.

**Remove old cluster**

```bash
pg_dropcluster 17 demo
```

```bash
rm -rf /var/lib/postgresql/17/demo
```

Install the new PostgreSQL binaries on the standby and create the cluster.

**Remove old cluster and create the new cluster**

```bash
pg_dropcluster 17 demo
```

```bash
rm -rf /var/lib/postgresql/17/demo
```

```bash
mkdir -p -m 700 /usr/lib/postgresql/18/bin
```

```bash
pg_createcluster 18 demo
```

Run the `check` on the repository host. The warning regarding the standby being down is expected since the standby cluster is down. Running this command demonstrates that the repository server is aware of the standby and is configured properly for the primary server.

**Check configuration**

```bash
pgbackrest --stanza=demo check
```

Run a full backup on the new cluster and then restore the standby from the backup. The backup type will automatically be changed to `full` if `incr` or `diff` is requested.

**Run a full backup**

```bash
pgbackrest --stanza=demo --type=full backup
```

**Restore the demo standby cluster**

```bash
pgbackrest --stanza=demo --delta --type=standby restore
```

```bash
pgbackrest --stanza=demo --type=standby restore
```

**Configure**

```ini
hot_standby = on
```

**Start  and check the  configuration**

```bash
pg_ctlcluster 18 demo start
```

```bash
pgbackrest --stanza=demo check
```

Backup from standby can be enabled now that the standby is restored.

**Reenable backup from standby**

```ini
[global]
backup-standby=y
```
