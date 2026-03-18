# Quick Start
<a name="quickstart"></a>

The Quick Start section will cover basic configuration of pgBackRest and PostgreSQL and introduce the `backup`, `restore`, and `info` commands.

## Setup Demo Cluster
<a name="setup-demo-cluster"></a>

Creating the demo cluster is optional but is strongly recommended, especially for new users, since the example commands in the user guide reference the demo cluster; the examples assume the demo cluster is running on the default port (i.e. 5432). The cluster will not be started until a later section because there is still some configuration to do.

**Create the demo cluster**

```bash
/usr/lib/postgresql/16/bin/initdb
                            -D /var/lib/postgresql/16/demo -k -A peer
```

```bash
pg_createcluster 16 demo
```

By default Debian/Ubuntu includes the day of the week in the log filename. This makes the user guide a bit more complicated so the `log_filename` is set to a constant.

**Set log_filename**

```ini
log_filename = 'postgresql.log'
```

## Configure Cluster Stanza
<a name="configure-stanza"></a>
`stanza`

The name 'demo' describes the purpose of this cluster accurately so that will also make a good stanza name.

pgBackRest needs to know where the base data directory for the PostgreSQL cluster is located. The path can be requested from PostgreSQL directly but in a recovery scenario the PostgreSQL process will not be available. During backups the value supplied to pgBackRest will be compared against the path that PostgreSQL is running on and they must be equal or the backup will return an error. Make sure that `pg-path` is exactly equal to `data_directory` as reported by PostgreSQL.

By default Debian/Ubuntu stores clusters in `/var/lib/postgresql/[version]/[cluster]` so it is easy to determine the correct path for the data directory.

When creating the `/etc/pgbackrest/pgbackrest.conf` file, the database owner (usually `postgres`) must be granted read privileges.

**Configure the  cluster data directory**

```ini
[demo]
pg1-path=/var/lib/postgresql/16/demo
[global]
log-timestamp=n
```

pgBackRest configuration files follow a Windows INI-like convention. Sections are denoted by text in brackets and key/value pairs are contained in each section. Lines beginning with `#` are ignored and can be used as comments. Quoting is not supported and whitespace is trimmed from keys and values. Sections will be merged if they appear more than once.

There are multiple ways the pgBackRest configuration files can be loaded:

- `config` and `config-include-path` are default: the default config file will be loaded, if it exists, and `*.conf` files in the default config include path will be appended, if they exist.
- `config` option is specified: only the specified config file will be loaded and is expected to exist.
- `config-include-path` is specified: `*.conf` files in the config include path will be loaded and the path is required to exist. The default config file will be be loaded if it exists. If it is desirable to load only the files in the specified config include path, then the `--no-config` option can also be passed.
- `config` and `config-include-path` are specified: using the user-specified values, the config file will be loaded and `*.conf` files in the config include path will be appended. The files are expected to exist.
- `config-path` is specified: this setting will override the base path for the default location of the config file and/or the base path of the default config-include-path setting unless the config and/or config-include-path option is explicitly set.

Files are concatenated as if they were one big file and each file must be valid individually. This means sections must be specified in each file where they are needed to store a key/value. Order doesn't matter but there is precedence based on sections. The precedence (highest to lowest) is:

- [*stanza*:*command*]
- [*stanza*]
- [global:*command*]
- [global]

!!! note

    `--config`, `--config-include-path` and `--config-path` are command-line only options.

pgBackRest can also be configured using environment variables as described in the [command reference](command.html).

**Configure log-path using the environment**

```bash
bash -c '
                        export PGBACKREST_LOG_PATH=/path/set/by/env &&
                        pgbackrest --log-level-console=error help backup log-path'
```

## Create the Repository
<a name="create-repository"></a>
`repo-path`

For this demonstration the repository will be stored on the same host as the PostgreSQL server. This is the simplest configuration and is useful in cases where traditional backup software is employed to backup the database host.

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

The repository path must be configured so pgBackRest knows where to find it.

**Configure the  repository path**

```ini
[global]
repo1-path=/var/lib/pgbackrest
```

Multiple repositories may also be configured. See [Multiple Repositories](#/multi-repo) for details.

## Azure-Compatible Object Store Support
<a name="azure-support"></a>

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

## GCS-Compatible Object Store Support
<a name="gcs-support"></a>

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

## S3-Compatible Object Store Support
<a name="s3-support"></a>

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

## SFTP Storage Support
<a name="sftp-support"></a>

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

## Configure Archiving
<a name="configure-archiving"></a>

Backing up a running PostgreSQL cluster requires WAL archiving to be enabled. Note that *at least* one WAL segment will be created during the backup process even if no explicit writes are made to the cluster.

**Configure archive settings**

```ini
archive_command = 'pgbackrest --stanza=demo archive-push %p'
archive_mode = on
wal_level = replica
max_wal_senders = 3
```

`%p` is how PostgreSQL specifies the location of the WAL segment to be archived. Setting `wal_level` to at least `replica` and increasing `max_wal_senders` is a good idea even if there are currently no replicas as this will allow them to be added later without restarting the primary cluster.

The PostgreSQL cluster must be restarted after making these changes and before performing a backup.

**Restart the demo cluster**

```bash
pg_ctlcluster 16 demo restart
```

When archiving a WAL segment is expected to take more than 60 seconds (the default) to reach the pgBackRest repository, then the pgBackRest `archive-timeout` option should be increased. Note that this option is not the same as the PostgreSQL `archive_timeout` option which is used to force a WAL segment switch; useful for databases where there are long periods of inactivity. For more information on the PostgreSQL `archive_timeout` option, see PostgreSQL [Write Ahead Log](https://www.postgresql.org/docs/current/static/runtime-config-wal.html).

The `archive-push` command can be configured with its own options. For example, a lower compression level may be set to speed archiving without affecting the compression used for backups.

**Config archive-push to use a lower compression level**

```ini
[global:archive-push]
compress-level=3
```

This configuration technique can be used for any command and can even target a specific stanza, e.g. `demo:archive-push`.

## Configure Retention
<a name="retention"></a>

pgBackRest expires backups based on retention options.

**Configure retention to 2 full backups**

```ini
[global]
repo1-retention-full=2
```

More information about retention can be found in the [Retention](#/retention) section.

## Configure Repository Encryption
<a name="configure-encryption"></a>

The repository will be configured with a cipher type and key to demonstrate encryption. Encryption is always performed client-side even if the repository type (e.g. S3 or other object store) supports encryption.

It is important to use a long, random passphrase for the cipher key. A good way to generate one is to run: `openssl rand -base64 48`.

**Configure  repository encryption**

```ini
[global]
repo1-cipher-type=aes-256-cbc
[global]
repo1-cipher-pass=zWaf6XtpjIVZC5444yXB+cgFDFl7MxGlgkZSaoPvTGirhPygu4jOKOXf9LO4vjfO
```

Once the repository has been configured and the stanza created and checked, the repository encryption settings cannot be changed.

## Create the Stanza
<a name="create-stanza"></a>

The `stanza-create` command must be run to initialize the stanza. It is recommended that the `check` command be run after `stanza-create` to ensure archiving and backups are properly configured.

**Create the stanza and check the configuration**

```bash
pgbackrest --stanza=demo --log-level-console=info stanza-create
```

## Check the Configuration
<a name="check-configuration"></a>
`check`

**Check the configuration**

```bash
pgbackrest --stanza=demo --log-level-console=info check
```

**Example of an invalid configuration**

```bash
pgbackrest --stanza=demo --archive-timeout=.1 check
```

## Performance Tuning
<a name="performance-tuning"></a>

pgBackRest has a number of performance options that are not enabled by default to maintain backward compatibility in the repository. However, when creating a new repository the following options are recommended. They can also be used on an existing repository with the caveat that older versions of pgBackRest will not be able to read the repository. This incompatibility depends on when the feature was introduced, which will be noted in the list below.

- `compress-type` - determines the compression algorithm used by the `backup` and `archive-push` commands. The default is `gz` (Gzip) but `zst` (Zstandard) is recommended because it is much faster and provides compression similar to `gz`. `zst` has been supported by the `compress-type` option since [v2.27](release.html#2.27). See [Compress Type](configuration.md#/section-general/option-compress-type) for more details.
- `repo-bundle` - combines small files during backup to save space and improve the speed of both the `backup` and `restore` commands, especially on object stores such as S3. The `repo-bundle` option was introduced in [v2.39](release.html#2.39). See [File Bundling](#/backup/bundle) for more details.
- `repo-block` - stores only the portions of files that have changed rather than the entire file during `diff`/`incr``backup`. This saves space and increases the speed of the `backup`. The `repo-block` option was introduced in [v2.46](release.html#2.46) but at least [v2.52.1](release.html#2.52.1) is recommended. See [Block Incremental](#/backup/block) for more details.

There are other performance options that are not enabled by default because they require additional configuration or because the default is safe (but not optimal). These options are available in all v2 versions of pgBackRest.

- `process-max` - determines how many processes will be used for commands. The default is 1, which is almost never the appropriate value. Each command uses `process-max` differently so refer to each command's documentation for details on usage.
- `archive-async` - archives WAL files to the repository in batch which greatly increases archiving speed. It is not enabled by default because it requires a spool path to be created. See [Asynchronous Archiving](#/async-archiving) for more details.
- `backup-standby` - performs the backup on a standby rather than the primary to reduce load on the primary. It is not enabled by default because it requires additional configuration and the presence of one or more standby hosts. See [Backup from a Standby](#/standby-backup) for more details.

## Perform a Backup
<a name="perform-backup"></a>

By default pgBackRest will wait for the next regularly scheduled checkpoint before starting a backup. Depending on the `checkpoint_timeout` and `checkpoint_segments` settings in PostgreSQL it may be quite some time before a checkpoint completes and the backup can begin. Generally, it is best to set `start-fast=y` so that the backup starts immediately. This forces a checkpoint, but since backups are usually run once a day an additional checkpoint should not have a noticeable impact on performance. However, on very busy clusters it may be best to pass `--start-fast` on the command-line as needed.

**Configure backup fast start**

```ini
[global]
start-fast=y
```

To perform a backup of the PostgreSQL cluster run pgBackRest with the `backup` command.

**Backup the demo cluster**

```bash
pgbackrest --stanza=demo
                        --log-level-console=info backup
```

By default pgBackRest will attempt to perform an incremental backup. However, an incremental backup must be based on a full backup and since no full backup existed pgBackRest ran a full backup instead.

The `type` option can be used to specify a full or differential backup.

**Differential backup of the demo cluster**

```bash
pgbackrest --stanza=demo --type=diff
                        --log-level-console=info backup
```

This time there was no warning because a full backup already existed. While incremental backups can be based on a full *or* differential backup, differential backups must be based on a full backup. A full backup can be performed by running the `backup` command with `--type=full`.

During an online backup pgBackRest waits for WAL segments that are required for backup consistency to be archived. This wait time is governed by the pgBackRest `archive-timeout` option which defaults to 60 seconds. If archiving an individual segment is known to take longer then this option should be increased.

## Schedule a Backup
<a name="schedule-backup"></a>

Backups can be scheduled with utilities such as cron.

In the following example, two cron jobs are configured to run; full backups are scheduled for 6:30 AM every Sunday with differential backups scheduled for 6:30 AM Monday through Saturday. If this crontab is installed for the first time mid-week, then pgBackRest will run a full backup the first time the differential job is executed, followed the next day by a differential backup.

**crontab**

```

                #m h   dom mon dow   command
                30 06  *   *   0     pgbackrest --type=full --stanza=demo backup
                30 06  *   *   1-6   pgbackrest --type=diff --stanza=demo backup
            
```

Once backups are scheduled it's important to configure retention so backups are expired on a regular schedule, see [Retention](#/retention).

## Backup Information
<a name="backup-info"></a>

Use the `info` command to get information about backups.

**Get info for the demo cluster**

```bash
pgbackrest info
```
`info`

## Restore a Backup
<a name="perform-restore"></a>

Backups can protect you from a number of disaster scenarios, the most common of which are hardware failure and data corruption. The easiest way to simulate data corruption is to remove an important PostgreSQL cluster file.

**Stop the demo cluster and delete the pg_control file**

```bash
pg_ctlcluster 16 demo stop
```

```bash
rm /var/lib/postgresql/16/demo/global/pg_control
```

Starting the cluster without this important file will result in an error.

**Attempt to start the corrupted demo cluster**

```bash
pg_ctlcluster 16 demo start
```

```bash
pg_ctlcluster 16 demo start
```

```bash
pg_lsclusters
```

To restore a backup of the PostgreSQL cluster run pgBackRest with the `restore` command. The cluster needs to be stopped (in this case it is already stopped) and all files must be removed from the PostgreSQL data directory.

**Remove old files from demo cluster**

```bash
find /var/lib/postgresql/16/demo -mindepth 1 -delete
```

**Restore the demo cluster and start**

```bash
pgbackrest --stanza=demo restore
```

```bash
pg_ctlcluster 16 demo start
```

This time the cluster started successfully since the restore replaced the missing `pg_control` file.

More information about the `restore` command can be found in the [Restore](#/restore) section.
