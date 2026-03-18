# Target Time for Repository
<a name="repo-target-time"></a>
`repo-target-time`

To demonstrate this feature the `demo` stanza in the S3 repo is deleted.

**Delete stanza in S3 repository**

```bash
pg_ctlcluster 16 demo stop
```

```bash
pgbackrest --stanza=demo stop
```

```bash
pgbackrest --stanza=demo --repo=3 stanza-delete
```

Once the stanza is deleted the `info` command will show the repository in an error state.

**Error on info**

```bash
pgbackrest --stanza=demo --repo=3 info
```

However, since the storage is versioned, it is possible to look at the repository at a time before the stanza was deleted. Finding the target time can be tricky depending on the situation, but in this case the time when the stanza was deleted can be determined by checking when `backup.info` was deleted.

**Use mc to list versions of backup.info in the bucket**

```bash
mc ls --versions s3/demo-bucket/demo-repo/backup/demo/backup.info
```

Now the `info` command can be run with a target time that will show the repository before it was deleted.

**Info with target time**

```bash
pgbackrest --stanza=demo --repo=3
                --repo-target-time="<limit-recovery-timestamp>" info
```

If the required backup is shown by the `info` command then it can be restored using the same target time.

**Restore with target time**

```bash
pgbackrest --stanza=demo --repo=3 --delta
                --repo-target-time="<limit-recovery-timestamp>" --log-level-console=info restore
```

```bash
pg_ctlcluster 16 demo start
```
