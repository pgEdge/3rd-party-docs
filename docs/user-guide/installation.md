# Installation
<a name="installation"></a>

A new host named `pg-primary` is created to contain the demo cluster and run pgBackRest examples.

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

pgBackRest should now be properly installed but it is best to check. If any dependencies were missed then you will get an error when running pgBackRest from the command line.

**Make sure the installation worked**

```bash
pgbackrest
```
