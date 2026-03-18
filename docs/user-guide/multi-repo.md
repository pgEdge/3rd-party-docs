# Multiple Repositories
<a name="multi-repo"></a>

Multiple repositories may be configured as demonstrated in [S3 Support](#/s3-support). A potential benefit is the ability to have a local repository for fast restores and a remote repository for redundancy.

Some commands, e.g. `stanza-create`/`stanza-upgrade`, will automatically work with all configured repositories while others, e.g. [stanza-delete](#/delete-stanza), will require a repository to be specified using the `repo` option. See the [command reference](command.html) for details on which commands require the repository to be specified.

Note that the `repo` option is not required when only `repo1` is configured in order to maintain backward compatibility. However, the `repo` option *is* required when a single repo is configured as, e.g. `repo2`. This is to prevent command breakage if a new repository is added later.

The `archive-push` command will always push WAL to the archive in all configured repositories. When a repository cannot be reached, WAL will still be pushed to other repositories. However, for this to work effectively, `archive-async=y` must be enabled; otherwise, the other repositories can only get one WAL segment ahead of the unreachable repository. Also, note that if WAL cannot be pushed to any repository, then PostgreSQL will not remove it from the `pg_wal` directory, which may cause the volume to run out of space.

Backups need to be scheduled individually for each repository. In many cases this is desirable since backup types and retention will vary by repository. Likewise, restores must specify a repository. It is generally better to specify a repository for restores that has low latency/cost even if that means more recovery time. Only restore testing can determine which repository will be most efficient.
