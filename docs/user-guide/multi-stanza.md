# Multiple Stanzas
<a name="multi-stanza"></a>

pgBackRest supports multiple stanzas. The most common usage is sharing a `repository` host among multiple stanzas.

## Installation
<a name="installation"></a>

A new host named `pg-alt` is created to run the new primary.

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

## Configuration
<a name="configuration"></a>

pgBackRest configuration is nearly identical to `pg-primary` except that the `demo-alt` stanza will be used so backups and archive will be stored in a separate location.

**Configure  on the new primary**

```ini
[demo-alt]
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
tls-server-auth=pgbackrest-client=demo-alt
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

**Configure pg1-host/pg1-host-user and pg1-path**

```ini
[demo-alt]
pg1-path=/var/lib/postgresql/17/demo
[demo-alt]
pg1-host=pg-alt
[demo-alt]
pg1-host-type=tls
[demo-alt]
pg1-host-ca-file=/etc/pgbackrest/cert/ca.crt
[demo-alt]
pg1-host-cert-file=/etc/pgbackrest/cert/client.crt
[demo-alt]
pg1-host-key-file=/etc/pgbackrest/cert/client.key
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

## Setup Demo Cluster
<a name="setup-demo-cluster"></a>

**Create the demo cluster**

```bash
/usr/lib/postgresql/17/bin/initdb
                            -D /var/lib/postgresql/17/demo -k -A peer
```

```bash
pg_createcluster 17 demo
```

**Configure  settings**

```ini
archive_command = 'pgbackrest --stanza=demo-alt archive-push %p'
archive_mode = on
log_filename = 'postgresql.log'
```

**Start the demo cluster**

```bash
pg_ctlcluster 17 demo restart
```

## Create the Stanza and Check Configuration
<a name="create-stanza"></a>

The `stanza-create` command must be run to initialize the stanza. It is recommended that the `check` command be run after `stanza-create` to ensure archiving and backups are properly configured.

**Create the stanza and check the configuration**

```bash
pgbackrest --stanza=demo-alt --log-level-console=info stanza-create
```

```bash
pgbackrest --log-level-console=info check
```

If the `check` command is run from the `repository` host then all stanzas will be checked.

**Check the configuration for all stanzas**

```bash
pgbackrest --log-level-console=info check
```
