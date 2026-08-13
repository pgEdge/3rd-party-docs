# Replication
<a name="replication"></a>

Replication allows multiple copies of a PostgreSQL cluster (called standbys) to be created from a single primary. The standbys are useful for balancing reads and to provide redundancy in case the primary host fails.

## Installation
<a name="installation"></a>

A new host named `pg-standby` is created to run the standby.

**Install  from package**

```bash
apt-get install pgbackrest
```

```bash
apt-get install pgbackrest
```

```bash
yum install pgbackrest
```

```bash
yum install pgbackrest
```

**Update permissions on configuration file and directories**

```bash
chown <br-install-user>:<br-install-group> /var/log/pgbackrest
```

```bash
chown <br-install-user>:<br-install-group> /etc/pgbackrest/pgbackrest.conf
```

## Setup Passwordless SSH
<a name="setup-ssh"></a>

pgBackRest can use passwordless SSH to enable communication between the hosts. It is also possible to use TLS, see [Setup TLS](https://pgbackrest.org/user-guide-rhel.html#repo-host/config).

**Create <setup-ssh-host> host key pair**

```bash
mkdir -m 750 -p <setup-ssh-user-home-path>/.ssh
```

```bash
ssh-keygen -f <setup-ssh-user-home-path>/.ssh/id_rsa
                    -t rsa -b 4096 -N ""
```

Exchange keys between `repository` and `<setup-ssh-host>`.

**Copy <setup-ssh-host> public key to repository**

```bash
(echo -n 'no-agent-forwarding,no-X11-forwarding,no-port-forwarding,' &&
                    echo -n 'command="/usr/bin/pgbackrest ${SSH_ORIGINAL_COMMAND#* }" ' &&
                    sudo ssh root@<setup-ssh-host> cat <setup-ssh-user-home-path>/.ssh/id_rsa.pub) |
                    sudo -u pgbackrest tee -a /home/pgbackrest/.ssh/authorized_keys
```

**Copy repository public key to <setup-ssh-host>**

```bash
(echo -n 'no-agent-forwarding,no-X11-forwarding,no-port-forwarding,' &&
                    echo -n 'command="/usr/bin/pgbackrest ${SSH_ORIGINAL_COMMAND#* }" ' &&
                    sudo ssh root@repository cat /home/pgbackrest/.ssh/id_rsa.pub) |
                    sudo -u <setup-ssh-user> tee -a <setup-ssh-user-home-path>/.ssh/authorized_keys
```

Test that connections can be made from `repository` to `<setup-ssh-host>` and vice versa.

**Test connection from repository to <setup-ssh-host>**

```bash
ssh <setup-ssh-user>@<setup-ssh-host>
```

**Test connection from <setup-ssh-host> to repository**

```bash
ssh pgbackrest@repository
```

## Hot Standby
<a name="hot-standby"></a>

A hot standby performs replication using the WAL archive and allows read-only queries.

pgBackRest configuration is very similar to `pg-primary` except that the `standby` recovery type will be used to keep the cluster in recovery mode when the end of the WAL stream has been reached.

**Configure  on the standby**

```ini
[demo]
pg1-path=/var/lib/postgresql/17/demo
[global]
repo1-host=repository
[global]
repo1-host-type=tls
[global]
repo1-host-ca-file=/etc/pgbackrest/cert/ca.crt
[global]
repo1-host-cert-file=/etc/pgbackrest/cert/client.crt
[global]
repo1-host-key-file=/etc/pgbackrest/cert/client.key
[global]
tls-server-auth=pgbackrest-client=demo
[global]
tls-server-address=*
[global]
tls-server-ca-file=/etc/pgbackrest/cert/ca.crt
[global]
tls-server-cert-file=/etc/pgbackrest/cert/server.crt
[global]
tls-server-key-file=/etc/pgbackrest/cert/server.key
[global]
log-level-file=detail
[global]
log-timestamp=n
```

**Setup pgBackRest Server**

```bash
cat /etc/systemd/system/pgbackrest.service
```

```bash
systemctl enable pgbackrest
```

```bash
systemctl start pgbackrest
```

The demo cluster must be created (even though it will be overwritten on restore) in order to create the PostgreSQL configuration files.

**Create demo cluster**

```bash
pg_createcluster 17 demo
```

Create the path where PostgreSQL will be restored.

**Create  path**

```bash
mkdir -p -m 700 /var/lib/postgresql/17/demo
```

Now the standby can be created with the `restore` command.

!!! important

    If the cluster is intended to be promoted without becoming the new primary (e.g. for reporting or testing), use `--archive-mode=off` or set `archive_mode=off` in `postgresql.conf` to disable archiving. If archiving is not disabled then the repository may be polluted with WAL that can make restores more difficult.

**Restore the demo standby cluster**

```bash
pgbackrest --stanza=demo --delta --type=standby restore
```

```bash
pgbackrest --stanza=demo --type=standby restore
```

```bash
cat /var/lib/postgresql/17/demo/postgresql.auto.conf
```

The `hot_standby` setting must be enabled before starting PostgreSQL to allow read-only connections on `pg-standby`. Otherwise, connection attempts will be refused. The rest of the configuration is in case the standby is promoted to a primary.

**Configure**

```ini
hot_standby = on
archive_command = 'pgbackrest --stanza=demo archive-push %p'
archive_mode = on
log_filename = 'postgresql.log'
```

**Start**

```bash
pg_ctlcluster 17 demo start
```

The PostgreSQL log gives valuable information about the recovery. Note especially that the cluster has entered standby mode and is ready to accept read-only connections.

**Examine the  log output for log messages indicating success**

```bash
cat /var/log/postgresql/postgresql-17-demo.log
```

An easy way to test that replication is properly configured is to create a table on `pg-primary`.

**Create a new table on the primary**

```bash
psql -c "
                                 begin;
                                 create table replicated_table (message text);
                                 insert into replicated_table values ('Important Data');
                                 commit;
                                 select * from replicated_table";
```

And then query the same table on `pg-standby`.

**Query new table on the standby**

```bash
psql -c "select * from replicated_table;"
```

So, what went wrong? Since PostgreSQL is pulling WAL segments from the archive to perform replication, changes won't be seen on the standby until the WAL segment that contains those changes is pushed from `pg-primary`.

This can be done manually by calling `pg_switch_wal()` which pushes the current WAL segment to the archive (a new WAL segment is created to contain further changes).

**Call pg_switch_wal()**

```bash
psql -c "select *, current_timestamp from pg_switch_wal()";
```

Now after a short delay the table will appear on `pg-standby`.

**Now the new table exists on the standby (may require a few retries)**

```bash
psql -c "
                        select *, current_timestamp from replicated_table"
```

Check the standby configuration for access to the repository.

**Check the configuration**

```bash
pgbackrest --stanza=demo --log-level-console=info check
```

## Streaming Replication
<a name="streaming"></a>

Instead of relying solely on the WAL archive, streaming replication makes a direct connection to the primary and applies changes as soon as they are made on the primary. This results in much less lag between the primary and standby.

Streaming replication requires a user with the replication privilege.

**Create replication user**

```bash
psql -c "
                            create user replicator password 'jw8s0F4' replication";
```

The `pg_hba.conf` file must be updated to allow the standby to connect as the replication user. Be sure to replace the IP address below with the actual IP address of your `pg-standby`. A reload will be required after modifying the `pg_hba.conf` file.

**Create pg_hba.conf entry for replication user**

```bash
sh -c 'echo
                        "host    replication     replicator      <host-pg2-ip>/32           md5"
                        >> /etc/postgresql/17/demo/pg_hba.conf'
```

```bash
pg_ctlcluster 17 demo reload
```

The standby needs to know how to contact the primary so the `primary_conninfo` setting will be configured in pgBackRest.

**Set primary_conninfo**

```ini
[demo]
recovery-option=primary_conninfo=host=<host-pg1-ip> port=5432 user=replicator
```

It is possible to configure a password in the `primary_conninfo` setting but using a `.pgpass` file is more flexible and secure.

**Configure the replication password in the .pgpass file.**

```bash
sh -c 'echo
                        "<host-pg1-ip>:*:replication:replicator:jw8s0F4"
                        >> /var/lib/postgresql/.pgpass'
```

```bash
chmod 600 /var/lib/postgresql/.pgpass
```

Now the standby can be created with the `restore` command.

**Stop  and restore the demo standby cluster**

```bash
pg_ctlcluster 17 demo stop
```

```bash
pgbackrest --stanza=demo --delta --type=standby restore
```

```bash
cat /var/lib/postgresql/17/demo/postgresql.auto.conf
```

!!! note

    The `primary_conninfo` setting has been written into the `postgresql.auto.conf` file because it was configured as a `recovery-option` in `pgbackrest.conf`. The `--type=preserve` option can be used with the `restore` to leave the existing `postgresql.auto.conf` file in place if that behavior is preferred.

By default Debian/Ubuntu stores the `postgresql.conf` file in the PostgreSQL data directory. That means the change made to `postgresql.conf` was overwritten by the last restore and the `hot_standby` setting must be enabled again. Other solutions to this problem are to store the `postgresql.conf` file elsewhere or to enable the `hot_standby` setting on the `pg-primary` host where it will be ignored.

**Enable hot_standby**

```ini
hot_standby = on
```

**Start**

```bash
pg_ctlcluster 17 demo start
```

The PostgreSQL log will confirm that streaming replication has started.

**Examine the  log output for log messages indicating success**

```bash
cat /var/log/postgresql/postgresql-17-demo.log
```

Now when a table is created on `pg-primary` it will appear on `pg-standby` quickly and without the need to call `pg_switch_wal()`.

**Create a new table on the primary**

```bash
psql -c "
                                 begin;
                                 create table stream_table (message text);
                                 insert into stream_table values ('Important Data');
                                 commit;
                                 select *, current_timestamp from stream_table";
```

**Query table on the standby**

```bash
psql -c "
                        select *, current_timestamp from stream_table"
```
