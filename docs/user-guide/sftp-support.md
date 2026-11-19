# SFTP Support
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

Commands are run exactly as if the repository were stored on a local disk.

**Add sftp-server fingerprint to known_hosts file since repo4-sftp-host-key-check-type defaults to strict**

```bash
ssh-keyscan -H sftp-server >> /var/lib/postgresql/.ssh/known_hosts 2>/dev/null
```

**Create the stanza**

```bash
pgbackrest --stanza=demo --log-level-console=info stanza-create
```

**Backup the demo cluster**

```bash
pgbackrest --stanza=demo --repo=4
                    --log-level-console=info backup
```
