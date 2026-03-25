# Dedicated Repository Host
<a name="repo-host"></a>

The configuration described in [Quickstart](quickstart.md#quickstart) is suitable for simple installations but for enterprise configurations it is more typical to have a dedicated `repository` host where the backups and WAL archive files are stored. This separates the backups and WAL archive from the database server so `database` host failures have less impact. It is still a good idea to employ traditional backup software to backup the `repository` host.

On PostgreSQL hosts, `pg1-path` is required to be the path of the local PostgreSQL cluster and no `pg1-host` should be configured. When configuring a repository host, the pgbackrest configuration file must have the `pg-host` option configured to connect to the primary and standby (if any) hosts. The repository host has the only pgbackrest configuration that should be aware of more than one PostgreSQL host. Order does not matter, e.g. pg1-path/pg1-host, pg2-path/pg2-host can be primary or standby.

## Installation
<a name="install"></a>

A new host named `repository` is created to store the cluster backups.

!!! note

    The pgBackRest version installed on the `repository` host must exactly match the version installed on the PostgreSQL host.

The `pgbackrest` user is created to own the pgBackRest repository. Any user can own the repository but it is best not to use `postgres` (if it exists) to avoid confusion.

**Create pgbackrest user**

```bash
adduser --disabled-password --gecos "" pgbackrest
```

```bash
groupadd pgbackrest
```

```bash
adduser -gpgbackrest -n pgbackrest
```

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

**Create the  repository**

```bash
mkdir -p /var/lib/pgbackrest
```

```bash
chmod 750 /var/lib/pgbackrest
```

```bash
chown <br-install-user>:<br-install-group> /var/lib/pgbackrest
```

**Update permissions on the  repository**

```bash
chown <br-install-user>:<br-install-group> /var/lib/pgbackrest
```

## Setup Passwordless SSH
<a name="setup-ssh"></a>

pgBackRest can use passwordless SSH to enable communication between the hosts. It is also possible to use TLS, see [Setup TLS](https://pgbackrest.org/user-guide-rhel.html#repo-host/config).

**Create repository host key pair**

```bash
mkdir -m 750 /home/pgbackrest/.ssh
```

```bash
ssh-keygen -f /home/pgbackrest/.ssh/id_rsa
                        -t rsa -b 4096 -N ""
```

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

!!! note

    ssh has been configured to only allow pgBackRest to be run via passwordless ssh. This enhances security in the event that one of the service accounts is hijacked.

## Configuration
<a name="config"></a>

pgBackRest can use TLS with client certificates to enable communication between the hosts. It is also possible to use SSH, see [Setup SSH](https://pgbackrest.org/user-guide.html#repo-host/setup-ssh).

pgBackRest expects client/server certificates to be generated in the same way as PostgreSQL. See [Secure TCP/IP Connections with TLS](https://www.postgresql.org/docs/current/ssl-tcp.html) for detailed instructions on generating certificates.

The `repository` host must be configured with the `pg-primary` host/user and database path. The primary will be configured as `pg1` to allow a standby to be added later.

**Configure pg1-host/pg1-host-user and pg1-path**

```ini
[demo]
pg1-path=/var/lib/postgresql/17/demo
[demo]
pg1-host=pg-primary
[demo]
pg1-host-type=tls
[demo]
pg1-host-ca-file=/etc/pgbackrest/cert/ca.crt
[demo]
pg1-host-cert-file=/etc/pgbackrest/cert/client.crt
[demo]
pg1-host-key-file=/etc/pgbackrest/cert/client.key
[global]
tls-server-auth=pgbackrest-client=*
[global]
tls-server-address=*
[global]
tls-server-ca-file=/etc/pgbackrest/cert/ca.crt
[global]
tls-server-cert-file=/etc/pgbackrest/cert/server.crt
[global]
tls-server-key-file=/etc/pgbackrest/cert/server.key
[global]
start-fast=y
[global]
repo1-retention-full=2
[global]
log-timestamp=n
```

The database host must be configured with the repository host/user. The default for the `repo1-host-user` option is `pgbackrest`. If the `postgres` user does restores on the repository host it is best not to also allow the `postgres` user to perform backups. However, the `postgres` user can read the repository directly if it is in the same group as the `pgbackrest` user.

**Configure repo1-host/repo1-host-user**

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

PostgreSQL configuration may be found in the [Configure Archiving](quickstart.md#configure-archiving) section.

Commands are run the same as on a single host configuration except that some commands such as `backup` and `expire` are run from the `repository` host instead of the `database` host.

Configure Azure-compatible object store if required.

pgBackRest supports locating repositories in Azure-compatible object stores. The container used to store the repository must be created in advance &mdash; pgBackRest will not do it automatically. The repository can be located in the container root (`/`) but it's usually best to place it in a subpath so object store logs or other data can also be stored in the container without conflicts.

!!! warning

    Do not enable "hierarchical namespace" as this will cause errors during expire.

**Configure Azure**

```ini
[global]
repo{[azure-setup-repo-id]}-type=azure
[global]
repo{[azure-setup-repo-id]}-path=/demo-repo
[global]
repo{[azure-setup-repo-id]}-azure-account=pgbackrest
[global]
repo{[azure-setup-repo-id]}-azure-key-type=shared
[global]
repo{[azure-setup-repo-id]}-azure-key=YXpLZXk=
[global]
repo{[azure-setup-repo-id]}-azure-container=demo-container
[global]
repo{[azure-setup-repo-id]}-retention-full=4
```

Shared access signatures may be used by setting the `repo<azure-setup-repo-id>-azure-key-type` option to `sas` and the `repo<azure-setup-repo-id>-azure-key` option to the shared access signature token.

Configure GCS-compatible object store if required.

pgBackRest supports locating repositories in GCS-compatible object stores. The bucket used to store the repository must be created in advance &mdash; pgBackRest will not do it automatically. The repository can be located in the bucket root (`/`) but it's usually best to place it in a subpath so object store logs or other data can also be stored in the bucket without conflicts.

**Configure GCS**

```ini
[global]
repo{[gcs-setup-repo-id]}-type=gcs
[global]
repo{[gcs-setup-repo-id]}-path=/demo-repo
[global]
repo{[gcs-setup-repo-id]}-gcs-key-type=service
[global]
repo{[gcs-setup-repo-id]}-gcs-key=/etc/pgbackrest/gcs-key.json
[global]
repo{[gcs-setup-repo-id]}-gcs-bucket=demo-bucket
```

When running in GCE set `repo<gcs-setup-repo-id>-gcs-key-type=auto` to automatically authenticate using the instance service account.

Configure S3-compatible object store if required.

pgBackRest supports locating repositories in S3-compatible object stores. The bucket used to store the repository must be created in advance &mdash; pgBackRest will not do it automatically. The repository can be located in the bucket root (`/`) but it's usually best to place it in a subpath so object store logs or other data can also be stored in the bucket without conflicts.

**Configure S3**

```ini
[global]
repo{[s3-setup-repo-id]}-type=s3
[global]
repo{[s3-setup-repo-id]}-path=/demo-repo
[global]
repo{[s3-setup-repo-id]}-s3-key=accessKey1
[global]
repo{[s3-setup-repo-id]}-s3-key-secret=verySecretKey1
[global]
repo{[s3-setup-repo-id]}-s3-bucket=demo-bucket
[global]
repo{[s3-setup-repo-id]}-s3-endpoint=s3.us-east-1.amazonaws.com
[global]
repo{[s3-setup-repo-id]}-s3-region=us-east-1
[global]
repo{[s3-setup-repo-id]}-retention-full=4
```

!!! note

    The region and endpoint will need to be configured to where the bucket is located. The values given here are for the `us-east-1` region.

Configure SFTP storage if required.

pgBackRest supports locating repositories on SFTP hosts. SFTP file transfer is relatively slow so commands benefit by increasing `process-max` to parallelize file transfer.

**Configure SFTP**

```ini
[global]
repo{[sftp-setup-repo-id]}-type=sftp
[global]
repo{[sftp-setup-repo-id]}-path=/demo-repo
[global]
repo{[sftp-setup-repo-id]}-bundle=y
[global]
repo{[sftp-setup-repo-id]}-sftp-host=sftp-server
[global]
repo{[sftp-setup-repo-id]}-sftp-host-key-hash-type=sha1
[global]
repo{[sftp-setup-repo-id]}-sftp-host-user=pgbackrest
[global]
repo{[sftp-setup-repo-id]}-sftp-private-key-file=<sftp-setup-user-home-path>/.ssh/id_rsa_sftp
[global]
repo{[sftp-setup-repo-id]}-sftp-public-key-file=<sftp-setup-user-home-path>/.ssh/id_rsa_sftp.pub
[global]
process-max=4
```

When utilizing SFTP, if libssh2 is compiled against OpenSSH then `repo<sftp-setup-repo-id>-sftp-public-key-file` is optional.

**Generate SSH keypair for SFTP backup**

```bash
mkdir -m 750 -p <sftp-setup-user-home-path>/.ssh
```

```bash
ssh-keygen -f <sftp-setup-user-home-path>/.ssh/id_rsa_sftp
                    -t rsa -b 4096 -N "" -m PEM
```

**Copy pg-primary SFTP backup public key to sftp-server**

```bash
mkdir -m 750 -p /home/pgbackrest/.ssh
```

```bash
(sudo ssh root@<sftp-setup-host> cat <sftp-setup-user-home-path>/.ssh/id_rsa_sftp.pub) |
                    sudo -u pgbackrest tee -a /home/pgbackrest/.ssh/authorized_keys
```

## Setup TLS Server
<a name="setup-tls"></a>

The pgBackRest TLS server must be configured and started on each host.

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

## Create and Check Stanza
<a name="stanza-create"></a>

Create the stanza in the new repository.

**Create the stanza**

```bash
pgbackrest --stanza=demo stanza-create
```

Check that the configuration is correct on both the `database` and `repository` hosts. More information about the `check` command can be found in [Check the Configuration](quickstart.md#check-configuration).

**Check the configuration**

```bash
pgbackrest --stanza=demo check
```

**Check the configuration**

```bash
pgbackrest --stanza=demo check
```

## Perform a Backup
<a name="perform-backup"></a>

To perform a backup of the PostgreSQL cluster run pgBackRest with the `backup` command on the `repository` host.

**Backup the demo cluster**

```bash
pgbackrest --stanza=demo backup
```

Since a new repository was created on the `repository` host the warning about the incremental backup changing to a full backup was emitted.

## Restore a Backup
<a name="perform-restore"></a>

To perform a restore of the PostgreSQL cluster run pgBackRest with the `restore` command on the `database` host.

**Stop the demo cluster, restore, and restart**

```bash
pg_ctlcluster 17 demo stop
```

```bash
pgbackrest --stanza=demo --delta restore
```

```bash
pg_ctlcluster 17 demo start
```
