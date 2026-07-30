<a id="logical-replication-quick-setup"></a>

## Quick Setup


 First set the configuration options in `postgresql.conf`:

```

wal_level = logical
```
 The other required settings have default values that are sufficient for a basic setup.


 `pg_hba.conf` needs to be adjusted to allow replication (the values here depend on your actual network configuration and user you want to use for connecting):

```

host     all     repuser     0.0.0.0/0     scram-sha-256
```


 Then on the publisher database:

```sql

CREATE PUBLICATION mypub FOR TABLE users, departments;
```


 And on the subscriber database:

```sql

CREATE SUBSCRIPTION mysub CONNECTION 'dbname=foo host=bar user=repuser' PUBLICATION mypub;
```


 The above will start the replication process, which synchronizes the initial table contents of the tables `users` and `departments` and then starts replicating incremental changes to those tables.
