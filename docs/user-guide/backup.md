# Backup
<a name="backup"></a>
`backup`

## File Bundling
<a name="bundle"></a>

Bundling files together in the repository saves time during the backup and some space in the repository. This is especially pronounced when the repository is stored on an object store such as S3 or file systems with large block sizes. Per-file creation time on object stores is higher and very small files might cost as much to store as larger files.

The file bundling feature is enabled with the `repo-bundle` option.

**Configure repo1-bundle**

```ini
[global]
repo1-bundle=y
```

A full backup without file bundling will have 1000+ files in the backup path, but with bundling the total number of files is greatly reduced. An additional benefit is that zero-length files are not stored (except in the manifest), whereas in a normal backup each zero-length file is stored individually.

**Perform a full backup**

```bash
pgbackrest --stanza=demo --type=full backup
```

**Check file total**

```bash
find /var/lib/pgbackrest/backup/demo/latest/ -type f | wc -l
```

The `repo-bundle-size` and `repo-bundle-limit` options can be used for tuning, though the defaults should be optimal in most cases.

While file bundling is generally more efficient, the downside is that it is more difficult to manually retrieve files from the repository. It may not be ideal for deduplicated storage since each full backup will arrange files in the bundles differently. Lastly, file bundles cannot be resumed, so be careful not to set `repo-bundle-limit` too high.

## Block Incremental
<a name="block"></a>

Block incremental backups save space by only storing the parts of a file that have changed since the prior backup rather than storing the entire file.

The block incremental feature is enabled with the `repo-block` option and it works best when enabled for all backup types. File bundling must also be enabled.

**Configure repo1-block**

```ini
[global]
repo1-block=y
```

## Backup Annotations
<a name="annotate"></a>

Users can attach informative key/value pairs to the backup. This option may be used multiple times to attach multiple annotations.

**Perform a full backup with annotations**

```bash
pgbackrest --stanza=demo --annotation=source="demo backup"
                        --annotation=key=value --type=full backup
```

Annotations are output by the `info` command text output when a backup is specified with `--set` and always appear in the JSON output.

**Get info for the demo cluster**

```bash
pgbackrest --stanza=demo --set=<backup-annotate-last> info
```

Annotations included with the `backup` command can be added, modified, or removed afterwards using the `annotate` command.

**Change backup annotations**

```bash
pgbackrest --stanza=demo --set=<backup-annotate-last>
                        --annotation=key= --annotation=new_key=new_value annotate
```

```bash
pgbackrest --stanza=demo --set=<backup-annotate-last> info
```
