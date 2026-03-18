# Upgrading pgBackRest
<a name="upgrading"></a>

## Upgrading pgBackRest from v1 to v2
<a name="v1-v2"></a>

Upgrading from v1 to v2 is fairly straight-forward. The repository format has not changed and all non-deprecated options from v1 are accepted, so for most installations it is simply a matter of installing the new version.

However, there are a few caveats:

- The deprecated `thread-max` option is no longer valid. Use `process-max` instead.
- The deprecated `archive-max-mb` option is no longer valid. This has been replaced with the `archive-push-queue-max` option which has different semantics.
- The default for the `backup-user` option has changed from `backrest` to `pgbackrest`.
- In v2.02 the default location of the pgBackRest configuration file has changed from `/etc/pgbackrest.conf` to `/etc/pgbackrest/pgbackrest.conf`. If `/etc/pgbackrest/pgbackrest.conf` does not exist, the `/etc/pgbackrest.conf` file will be loaded instead, if it exists.

Many option names have changed to improve consistency although the old names from v1 are still accepted. In general, `db-*` options have been renamed to `pg-*` and `backup-*`/`retention-*` options have been renamed to `repo-*` when appropriate.

PostgreSQL and repository options must be indexed when using the new names introduced in v2, e.g. `pg1-host`, `pg1-path`, `repo1-path`, `repo1-type`, etc.

## Upgrading pgBackRest from v2.x to v2.y
<a name="v2.x"></a>

Upgrading from v2.x to v2.y is straight-forward. The repository format has not changed, so for most installations it is simply a matter of installing binaries for the new version. It is also possible to downgrade if you have not used new features that are unsupported by the older version.

!!! important

    The local and remote pgBackRest versions must match exactly so they should be upgraded together. If there is a mismatch, WAL archiving and backups will not function until the versions match. In such a case, the following error will be reported: `[ProtocolError] expected value '2.x' for greeting key 'version' but got '2.y'`.
