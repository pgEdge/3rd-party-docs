# pgBackRest Releases

## v2.58.0 — Object Storage Improvements

*Released: 2026-01-19*

### Core Bug Fixes

- Fix deadlock due to logging in signal handler. 

### Core Features

- HTTP support for S3, GCS, and Azure. 

- Allow expiration of oldest full backup regardless of current retention. 

- Support for Azure managed identities. 

- ***Experimental*** support for S3 EKS pod identity. 

- Allow configuration of TLS cipher suites. 

- Allow process priority to be set. 

### Core Improvements

- Allow dots in S3 bucket names when using path-style URIs. 

- Require TLS >= 1.2 unless verification is disabled. 

- Dynamically size S3/GCS/Azure chunks for large uploads. 

- Optimize S3/GCS/Azure chunk size for small files. 

- Remove support for PostgreSQL `9.5`. 

- Improve logging of default for options with an unresolved dependency. 

### Documentation Improvements

- Remove explicit `max_wal_senders`/`wal_level` configuration from user guide. 

- Clarify that bundling is useful for filesystems with large block sizes. 

## v2.57.0 — Suppress Repository Symlinks

*Released: 2025-10-18*

### Core Bug Fixes

- Unnest HTTP/TLS/socket timeouts. 

- Fix possible segfault in page checksum error message. 

### Core Features

- Add `repo-symlink` option to suppress creation of repository symlinks. 

### Core Improvements

- Add HTTP retries for 408 and 429 errors. 

## v2.56.0 — Progress Info Improvements

*Released: 2025-07-21*

### Core Bug Fixes

- Fix issue with adhoc expiration when no backups in a repository. 

### Core Features

- Add restore progress to `info` command output. 

- Add progress-only detail level for `info` command output. 

### Core Improvements

- Retry failed reads on object stores. 

- Fix defaults in command-line help. 

### Core Development

- Improve the predictability of floating point numbers formatted as strings. 

### Documentation Improvements

- Describe discrete option values in a list where appropriate. 

- Fix "less than" in help output for `archive-mode` option. 

## v2.55.1 — Bug Fixes

*Released: 2025-05-05*

### Core Bug Fixes

- Revert "calculate content-md5 on S3 only when required". 

- Fix lower bounds checking for option keys. 

## v2.55.0 — Verification Improvements and PostgreSQL 18 Support

*Released: 2025-04-21*

### Core Bug Fixes

- Fix block incremental restore issue on non-default repository. 

- Do not set `recovery_target_timeline=current` for PostgreSQL &lt; 12. 

- Fix expire archive range logging. 

- Fix error reporting for queries with no results. 

### Core Features

- Verify recovery target timeline. 

- Allow verification of a specified backup. 

- Add support for S3/GCS requester pays. 

- PostgreSQL 18 support. 

- Allow connections to PostgreSQL on abstract domain sockets. 

- Add numeric output to `version` command. 

### Core Improvements

- Allow backup command to operate on remote repositories. 

- Use lz4 for protocol compression. 

- Calculate `content-md5` on S3 only when required. 

- Warn when a value for a multi-key option is overwritten. 

- Add detail logging for expired archive path. 

- Remove support for PostgreSQL `9.4`. 

- Remove autoconf/make build. 

### Core Development

- Revert "full/incremental backup method." 

- Improve hex encode performance with bytewise lookup. 

- Fix missing `return` in `FUNCTION_LOG_RETURN_VOID()`. 

- Remove extraneous const from `cvtIntToZ()` prototype. 

### Documentation Improvements

- Fix documentation for specifying multiple stanzas with `tls-server-auth`. 

- Clarify incremental backup expiration. 

- Clarify requirement for local/remote pgBackRest versions to match. 

- Add FAQ about exporting self-contained cluster. 

- Caveat `--tablespace-map-all` regarding tablespace creation. 

- Clarify behavior of `--repo-retention-full-type`. 

- Change `--process-max` recommendation for object stores to `--repo-bundle`. 

- Update `unix_socket_directory` to `unix_socket_directories`. 

- Recommend not placing `spool-path` within `pg_xlog`/`pg_wal`. 

### Documentation Development

- Fix instances of "of of". 

### Test Development

- Fix tests, logging, and comments in unit tests. 

## v2.54.2 — Bug Fix

*Released: 2025-01-20*

### Core Bug Fixes

- Fix issue after disabling bundling with block incremental enabled. 

### Documentation Improvements

- Clarify behavior of multiple configuration files. 

## v2.54.1 — Bug Fix

*Released: 2024-12-16*

### Core Bug Fixes

- Fix issue with `version`/`help` commands attempting to load `pgbackrest.conf`. 

### Test Improvements

- Stabilize async archiving in integration tests. 

## v2.54.0 — Target Time for Versioned Storage

*Released: 2024-10-21*

### Core Bug Fixes

- Fix PostgreSQL query performance for large datasets. 

### Core Features

- Allow repositories on versioned storage to be read at a target time. 

- Allow requested standby backup to proceed with no standby. 

### Core Improvements

- Summarize backup reference list for `info` command text output. 

- Refresh web-id token for each S3 authentication. 

- Correctly display current values for indexed options in help. 

- Save `backup.info` only when contents have changed. 

- Remove limitation on reading files in parallel during restore. 

- Improve SFTP error messages. 

### Core Development

- Protocol command multiplexing. 

### Documentation Features

- Add performance tuning section to user guide. 

### Documentation Improvements

- Clarify source for `data_directory`. 

- Better logic for deciding when a summary should be lower-cased. 

## v2.53.1 — PostgreSQL 17 Support

*Released: 2024-08-19*

### Core Bug Fixes

- Fix permissions when `restore` run as root user. 

- Fix segfault on delayed connection errors. 

- Skip local repository duplicate check for SFTP. 

### Core Improvements

- PostgreSQL 17 support. 

## v2.53 — Concurrent Backups

*Released: 2024-07-22*

### Core Bug Fixes

- Fix SFTP renaming failure when file already exists. 

### Core Features

- Allow backups to run concurrently on different repositories. 

- Support IP-based SANs for TLS certificate validation. 

### Core Improvements

- Default `log-level-stderr` option to `off`. 

- Allow alternative WAL segment sizes for PostgreSQL &le; 10. 

- Add hint to check SFTP authorization log. 

### Documentation Improvements

- Clarify `archive-push` multi-repo behavior. 

## v2.52.1 — Bug Fix

*Released: 2024-06-25*

### Core Bug Fixes

- Fix issue with files larger on the replica than on the primary. 

## v2.52 — PostgreSQL 17beta1 Support

*Released: 2024-05-27*

### Core Features

- Add GCS batch delete support. 

- S3 `SSE-C` encryption support. 

- PostgreSQL 17beta1 support. 

### Core Improvements

- Allow explicit disabling of optional dependencies in meson builds. 

- Dynamically find python in meson build. 

- Tag `pgbackrest` build target in meson as installable. 

### Core Development

- Update Fedora test image to Fedora 38. 

### Documentation Improvements

- Update `start`/`stop` documentation to reflect actual functionality. 

## v2.51 — Meson Build System

*Released: 2024-03-25*

### Core Bug Fixes

- Skip zero-length files for block incremental delta restore. 

- Fix performance regression in storage list. 

- Fix progress logging when file size changes during backup. 

### Core Improvements

- Improved support for dual stack connections. 

- Make meson the primary build system. 

- Detect files that have not changed during non-delta incremental backup. 

- Prevent invalid recovery when `backup_label` removed. 

- Improve `archive-push` WAL segment queue handling. 

- Limit resume functionality to full backups. 

- Update resume functionality for block incremental. 

- Allow `--version` and `--help` for version and help. 

- Add detailed backtrace to autoconf/make build. 

### Core Development

- Improve sort comparators. 

### Documentation Improvements

- Update references to `recovery.conf`. 

## v2.50 — Performance Improvements and Bug Fixes

*Released: 2024-01-22*

### Core Bug Fixes

- Fix short read in block incremental restore. 

- Fix overflow suppressing backup progress in `info` output. 

### Core Improvements

- Preserve partial files during block incremental delta restore. 

- Add support for alternate compile-time page sizes. 

- Skip files truncated during backup when bundling. 

- Improve SFTP storage error messages. 

### Core Development

- Refactor backup incremental manifest generation. 

## v2.49 — Remove PostgreSQL 9.3 Support

*Released: 2023-11-27*

### Core Bug Fixes

- Fix regression in retries. 

- Fix recursive path remove in SFTP storage driver. 

### Core Improvements

- Remove support for PostgreSQL `9.3`. 

### Core Development

- Fix `storageReadRemote()` to return actual read bytes. 

### Documentation Features

- Document maintainer options. 

- Update point-in-time recovery documentation for PostgreSQL >= 13. 

### Documentation Development

- Build command and configuration reference in C. 

### Test Improvements

- Allow `config/load` unit test to run without `libssh2` installed. 

### Test Development

- Remove unused fields from `backupJobResult()` test in command/backup unit. 

## v2.48 — Repository Storage Tags

*Released: 2023-09-25*

### Core Bug Fixes

- Fix issue restoring block incremental without a block list. 

### Core Features

- Add `--repo-storage-tag` option to create object tags. 

- Add known hosts checking for SFTP storage driver. 

- Support for dual stack connections. 

- Add backup size completed/total to `info` command JSON output. 

### Core Improvements

- Multi-stanza check command. 

- Retry reads of `pg_control` until checksum is valid. 

- Optimize WAL segment check after successful backup. 

- Improve GCS multi-part performance. 

- Allow `archive-get` command to run when stanza is stopped. 

- Accept leading tilde in paths for SFTP public/private keys. 

- Reload GCS credentials before renewing authentication token. 

### Core Development

- Add report option to `check` command. 

### Documentation Bug Fixes

- Fix configuration reference example for the `tls-server-address` option. 

- Fix command reference example for the `filter` option. 

### Test Improvements

- Allow `storage/sftp` unit test to run without `libssh2` installed. 

## v2.47 — Performance Improvements and Bug Fixes

*Released: 2023-07-24*

### Core Bug Fixes

- Preserve block incremental info in manifest during delta backup. 

- Fix block incremental file names in `verify` command. 

- Fix spurious automatic delta backup on backup from standby. 

- Skip `recovery.signal` for PostgreSQL >= 12 when recovery `type=none`. 

- Fix unique label generation for diff/incr backup. 

- Fix time-based archive expiration when no backups are expired. 

### Core Improvements

- Improve performance of SFTP storage driver. 

- Add timezone offset to `info` command date/time output. 

- Centralize error handling for unsupported features. 

### Documentation Improvements

- Clarify preference to install from packages in the user guide. 

## v2.46 — Block Incremental Backup and SFTP Storage

*Released: 2023-05-22*

### Core Features

- Block incremental backup. 

- SFTP support for repository storage. 

- PostgreSQL 16 support. 

### Core Improvements

- Allow page header checks to be skipped. 

- Avoid `chown()` on recovery files during restore. 

- Add error retry detail for HTTP retries. 

### Core Development

- Improve internal `manifest` command. 

### Documentation Improvements

- Add warning about using recovery `type=none`. 

- Add note about running `stanza-create` on already-created repositories. 

## v2.45 — Block Incremental Backup (BETA)

*Released: 2023-03-20*

### Core Bug Fixes

- Skip writing `recovery.signal` by default for restores of offline backups. 

### Core Features

- Block incremental backup (BETA). 

### Core Improvements

- Keep only one all-default group index. 

### Core Development

- Add `pg-version-force` option for fork integration. 

### Documentation Improvements

- Add explicit instructions for upgrading between `2.x` versions. 

- Remove references to SSH made obsolete when TLS was introduced. 

## v2.44 — Remove PostgreSQL 9.0/9.1/9.2 Support

*Released: 2023-01-30*

### Core Improvements

- Remove support for PostgreSQL `9.0`/`9.1`/`9.2`. 

- Restore errors when no backup matches the current version of PostgreSQL. 

- Add `compress-level` range checking for each `compress-type`. 

### Core Development

- Block-level incremental backup. 

- Add repository checksum to make verify and resume more efficient. 

- Add meson unity build and tests. 

- Refactor `common/ini` module to remove callbacks and duplicated code. 

### Documentation Improvements

- Add warning about enabling "hierarchical namespace" on Azure storage. 

- Add replacement for linefeeds in monitoring example. 

- Clarify `target-action` behavior on various PostgreSQL versions. 

- Updates and clarifications to index page. 

- Add dark mode to the website. 

## v2.43 — Bug Fix

*Released: 2022-11-28*

### Core Bug Fixes

- Fix missing reference in diff/incr backup. 

### Core Improvements

- Add hint when an option is specified without an index. 

## v2.42 — Bug Fixes

*Released: 2022-11-22*

### Core Bug Fixes

- Fix memory leak in file bundle `backup`/`restore`. 

- Fix protocol error on short read of remote file. 

### Core Improvements

- Do not store references for zero-length files when bundling. 

- Use more generic descriptions for `pg_start_backup()`/`pg_stop_backup()`. 

### Core Development

- Block incremental backup. 

- Move link creation to storage interface. 

### Test Improvements

- Update `test.pl` `--psql-bin` option to match command-line help. 

## v2.41 — Backup Annotations

*Released: 2022-09-19*

### Core Bug Fixes

- Fix incorrect time expiration being used for non-default repositories. 

- Fix issue when listing directories recursively with a filter. 

### Core Features

- Backup key/value annotations. 

### Core Improvements

- Support `--set` in JSON output for `info` command. 

- Allow upload chunk size to be configured for object stores. 

- Update archive.info timestamps after a successful backup. 

- Move standby timeline check after checkpoint. 

- Improve warning message on backup resume. 

### Core Development

- Skip mem context cleanup in `CATCH_FATAL()` block. 

- Do not allow `CATCH()` to catch a fatal error. 

- Add `FN_INLINE_ALWAYS` macro. 

### Documentation Improvements

- Add absolute path for `kill` in `pgbackrest.service`. 

## v2.40 — OpenSSL 3 Support

*Released: 2022-07-18*

### Core Improvements

- OpenSSL 3 support. 

- Create snapshot when listing contents of a path. 

- Force `target-timeline=current` when restore `type=immediate`. 

- Truncate files during delta `restore` when they are larger than expected. 

- Disable incremental manifest save when `resume=n`. 

- Set backup percent complete to zero before copy start. 

- Use S3 `IsTruncated` flag to determine list continuation. 

### Core Development

- Automatically create PostgreSQL version interfaces. 

- Improve memory usage of mem contexts. 

- Add `FN_NO_RETURN` macro. 

- Add cast to handle compilers (e.g. MSVC) that coerce to signed int. 

- Fix hard-coded WAL size assertion in `walSegmentNext()`. 

### Documentation Bug Fixes

- Skip internal options in the configuration reference. 

### Documentation Improvements

- Add link to PostgreSQL configuration in repository host section. 

### Test Improvements

- Add experimental Meson build. 

- Allow any path to be passed to the `--test-path` option. 

- Fix compile error when `DEBUG_EXEC_TIME` is defined without `DEBUG`. 

## v2.39 — Verify and File Bundling

*Released: 2022-05-16*

### Core Bug Fixes

- Fix error thrown from `FINALLY()` causing an infinite loop. 

- Error on all lock failures except another process holding the lock. 

### Core Features

- Backup file bundling for improved small file support. 

- Verify command to validate the contents of a repository. 

- PostgreSQL 15 support. 

- Show backup percent complete in `info` output. 

- Auto-select backup for `restore` command `--type=lsn`. 

- Suppress existing WAL warning when `archive-mode-check` is disabled. 

- Add AWS IMDSv2 support. 

### Core Improvements

- Allow `repo-hardlink` option to be changed after full backup. 

- Increase precision of percent complete logging for `backup` and `restore`. 

- Improve path validation for `repo-*` commands. 

- Improve `stop` command to honor `stanza` option. 

- Improve error message for invalid `repo-azure-key`. 

- Add hint to check the log on `archive-get`/`archive-push` async error. 

- Add `ClockError` for unexpected clock skew and timezone changes. 

- Strip extensions from history manifest before showing in error message. 

- Add user:group to lock permission error. 

### Core Development

- Improve JSON handling. 

- Add default for boolean options with unresolved dependencies. 

- Lock module refactoring. 

- Prevent `memContextFree()` from freeing memory needed by callbacks. 

- Refactor `PgClient` to return results in `Pack` format. 

### Documentation Bug Fixes

- Fix incorrect reference to `stanza-update` in the user guide. 

- Fix example for `repo-gcs-key-type` option in configuration reference. 

- Fix `tls-server-auth` example and add clarifications. 

### Documentation Improvements

- Simplify messaging around supported versions in the documentation. 

- Add option type descriptions. 

- Add FAQ about backup types and restore speed. 

- Document required base branch for pull requests. 

### Test Development

- Remove integration expect log testing. 

## v2.38 — Minor Bug Fixes and Improvements

*Released: 2022-03-06*

### Core Bug Fixes

- Retry errors in S3 batch file delete. 

- Allow case-insensitive matching of HTTP `connection` header values. 

### Core Features

- Add support for AWS S3 server-side encryption using KMS. 

- Add `archive-missing-retry` option. 

- Add backup type filter to `info` command. 

### Core Improvements

- Retry on page validation failure during `backup`. 

- Handle TLS servers that do not close connections gracefully. 

- Add backup LSNs to `info` command output. 

- Automatically strip trailing slashes for `repo-ls` paths. 

- Do not retry fatal errors. 

- Remove support for PostgreSQL `8.3`/`8.4`. 

- Remove logic that tried to determine additional file system compression. 

### Core Development

- Disconnect `help.auto.c` from `build-help` in `Makefile`. 

- Improve small file support. 

- Add percent complete to lock file. 

- Configuration refactoring. 

### Documentation Bug Fixes

- Move `repo` options in TLS documentation to the `global` section. 

- Remove unused `backup-standby` option from stanza commands. 

- Fix typos in help and release notes. 

### Documentation Improvements

- Add aliveness check to systemd service configuration. 

- Add FAQ explaining WAL archive suffix. 

- Note that replications slots are not restored. 

## v2.37 — TLS Server

*Released: 2022-01-03*

### Core Bug Fixes

- Fix `restore` delta link mapping when path/file already exists. 

- Fix socket leak on connection retries. 

### Core Features

- Add TLS server. 

- Add `--cmd` option. 

### Core Improvements

- Check archive immediately after backup start. 

- Add timeline and checkpoint checks to backup. 

- Check that clusters are alive and correctly configured during a backup. 

- Error when `restore` is unable to find a backup to match the time target. 

- Parse protocol/port in S3/Azure endpoints. 

- Add warning when `checkpoint_timeout` exceeds `db-timeout`. 

- Add verb to HTTP error output. 

- Allow y/n arguments for boolean command-line options. 

- Make backup size logging exactly match `info` command output. 

### Core Development

- Simplify error handler. 

- Add `StringId` as an option type. 

### Documentation Improvements

- Display size option default and allowed values with appropriate units. 

- Fix typos and improve documentation for the `tablespace-map-all` option. 

- Remove obsolete statement about future multi-repository support. 

## v2.36 — Minor Bug Fixes and Improvements

*Released: 2021-11-01*

### Core Bug Fixes

- Allow "global" as a stanza prefix. 

- Fix segfault on invalid GCS key file. 

### Core Improvements

- Allow `link-map` option to create new links. 

- Increase max index allowed for `pg`/`repo` options to 256. 

- Add `WebIdentity` authentication for AWS S3. 

- Report backup file validation errors in `backup.info`. 

- Add recovery start time to online backup restore log. 

- Report original error and retries on local job failure. 

- Rename page checksum error to error list in info text output. 

- Add hints to standby replay timeout message. 

### Core Development

- Add TLS server. 

- Store config values as a union instead of a variant. 

### Documentation Development

- Eliminate linefeed formatting from documentation. 

## v2.35 — Binary Protocol

*Released: 2021-08-23*

### Core Bug Fixes

- Detect errors in S3 multi-part upload finalize. 

- Fix detection of circular symlinks. 

- Only pass selected `repo` options to the remote. 

### Core Improvements

- Binary protocol. 

- Automatically create data directory on `restore`. 

- Allow `restore` `--type=lsn`. 

- Change level of `backup`/`restore` copied file logging to detail. 

- Loop while waiting for checkpoint LSN to reach replay LSN. 

- Log `backup` file total and `restore` size/file total. 

### Core Development

- Add support for more `Pack` types. 

- Clear error when a `CATCH()` block finishes. 

### Documentation Bug Fixes

- Fix incorrect host names in user guide. 

### Documentation Improvements

- Update contributing documentation and add pull request template. 

- Rearrange backup documentation in user guide. 

- Clarify `restore` `--type` behavior in command reference. 

- Fix documentation and comment typos. 

### Test Improvements

- Add check for test path inside repo path. 

- Add CodeQL static code analysis. 

- Update tests to use standard patterns. 

## v2.34 — PostgreSQL 14 Support

*Released: 2021-06-07*

### Core Bug Fixes

- Fix issues with leftover spool files from a prior `restore`. 

- Fix issue when checking links for large numbers of tablespaces. 

- Free no longer needed remotes so they do not timeout during `restore`. 

- Fix `help` when a valid option is invalid for the specified command. 

### Core Features

- Add PostgreSQL 14 support. 

- Add automatic GCS authentication for GCE instances. 

- Add `repo-retention-history` option to expire backup history. 

- Add `db-exclude` option. 

### Core Improvements

- Change archive expiration logging from detail to info level. 

- Remove stanza archive spool path on `restore`. 

- Do not write files atomically or sync paths during `backup` copy. 

### Core Development

- Replace `getopt_long()` with custom implementation. 

- Add help for all internal options valid for default roles. 

- Add `StringId` type. 

- Add local process shim. 

- Add `cfgOptionDisplay()`/`cfgOptionIdxDisplay()`. 

- Revert to checking catalog version for all PostgreSQL versions. 

- Rename default command role to main. 

- Simplify defaults for `--force` option. 

### Documentation Improvements

- Update contributing documentation. 

- Consolidate RHEL/CentOS user guide into a single document. 

- Clarify that `repo-s3-role` is not an `ARN`. 

## v2.33 — Multi-Repository and GCS Support

*Released: 2021-04-05*

### Core Bug Fixes

- Fix option warnings breaking async `archive-get`/`archive-push`. 

- Fix memory leak in backup during archive copy. 

- Fix stack overflow in cipher passphrase generation. 

- Fix `repo-ls` `/` on S3 repositories. 

### Core Features

- Multiple repository support. 

- GCS support for repository storage. 

- Add `archive-header-check` option. 

### Core Improvements

- Include recreated system databases during selective restore. 

- Exclude `content-length` from S3 signed headers. 

- Consolidate less commonly used repository storage options. 

- Allow custom `config-path` default with `./configure --with-configdir`. 

- Log archive copy during `backup`. 

### Core Development

- Improve protocol handlers. 

- Add `HttpUrl` object. 

### Documentation Improvements

- Update reference to include links to user guide examples. 

- Update selective restore documentation with caveats. 

- Add `compress-type` clarification to `archive-copy` documentation. 

- Add `compress-level` defaults per `compress-type` value. 

- Add note about required NFS settings being the same as PostgreSQL. 

## v2.32 — Repository Commands

*Released: 2021-02-08*

### Core Bug Fixes

- Fix resume after partial delete of backup by prior resume. 

### Core Features

- Add `repo-ls` command. 

- Add `repo-get` command. 

- Add `archive-mode-check` option. 

### Core Improvements

- Improve `archive-get` performance. 

### Core Development

- Partial multi-repository implementation. 

- Add backup verification to internal verify command. 

- Add `pack` type. 

- Allow option validity to be determined by command role. 

- Add `job-retry` and `job-retry-interval` options. 

- Replace `double` type with `time` in `config` module. 

### Documentation Improvements

- Improve `expire` command documentation. 

## v2.31 — Minor Bug Fixes and Improvements

*Released: 2020-12-07*

### Core Bug Fixes

- Allow `[`, `#`, and `space` as the first character in database names. 

- Create `standby.signal` only on PostgreSQL 12 when `restore` type is `standby`. 

### Core Features

- Expire history files. 

- Report page checksum errors in `info` command `text` output. 

- Add `repo-azure-endpoint` option. 

- Add `pg-database` option. 

### Core Improvements

- Improve `info` command output when a stanza is specified but missing. 

- Improve performance of large file lists in `backup`/`restore` commands. 

- Add retries to PostgreSQL sleep when starting a backup. 

### Core Development

- Prepare configuration module for multi-repository support. 

- Optimize small reads in `IoRead`. 

- Allow multiple remote locks from the same main process. 

- Conform retry in `lockAcquireFile()` to the common retry pattern. 

- Assert when buffer used is greater than size limit. 

### Documentation Improvements

- Replace RHEL/CentOS 6 documentation with RHEL/CentOS 8. 

## v2.30 — PostgreSQL 13 Support

*Released: 2020-10-05*

### Core Bug Fixes

- Error with hints when backup user cannot read `pg_settings`. 

### Core Features

- PostgreSQL 13 support. 

### Core Improvements

- Improve PostgreSQL version identification. 

- Improve working directory error message. 

- Add hint about starting the stanza when WAL segment not found. 

- Add hint for protocol version mismatch. 

### Core Development

- Add internal verify command. 

- Allow `ProtocolParallel` to complete with no jobs. 

### Documentation Improvements

- Add note that pgBackRest versions must match when running remotely. 

- Move info command text to the reference and link to user guide. 

- Update yum repository path for CentOS/RHEL user guide. 

## v2.29 — Auto S3 Credentials on AWS

*Released: 2020-08-31*

### Core Bug Fixes

- Suppress errors when closing `local`/`remote` processes. Since the command has completed it is counterproductive to throw an error but still **warn** to indicate that something unusual happened. 

- Fix issue with `=` character in file or database names. 

### Core Features

- Automatically retrieve temporary S3 credentials on AWS instances. 

- Add `archive-mode` option to disable archiving on restore. 

### Core Improvements

- PostgreSQL 13 beta3 support. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

- Asynchronous `list`/`remove` for S3/Azure storage. 

- Improve memory usage of unlogged relation detection in manifest build. 

- Proactively close file descriptors after forking async process. 

- Delay backup remote connection close until after archive check. 

- Improve detailed error output. 

- Improve TLS error reporting. 

### Core Development

- Allow `HttpClient`/`HttpSession` to work on plain sockets. 

- Add support for `HTTP/1.0`. 

- Add general-purpose statistics collector. 

- Add `user-agent` to HTTP requests. 

### Documentation Bug Fixes

- Add `none` to `compress-type` option reference and fix example. 

- Add missing `azure` type in `repo-type` option reference. 

- Fix typo in `repo-cipher-type` option reference. 

### Documentation Improvements

- Clarify that `expire` must be run regularly when `expire-auto` is disabled. 

## v2.28 — Azure Repository Storage

*Released: 2020-07-20*

### Core Bug Fixes

- Fix `restore` `--force` acting like `--force --delta`. This caused `restore` to replace files based on timestamp and size rather than overwriting, which meant some files that should have been updated were left unchanged. Normal `restore` and `restore` `--delta` were not affected by this issue. 

### Core Features

- Azure support for repository storage. 

- Add `expire-auto` option. This allows automatic expiration after a successful backup to be disabled. 

### Core Improvements

- Asynchronous S3 multipart upload. 

- Automatic retry for `backup`, `restore`, `archive-get`, and `archive-push`. 

- Disable query parallelism in PostgreSQL sessions used for backup control. 

- PostgreSQL 13 beta2 support. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

- Improve handling of invalid HTTP response status. 

- Improve error when `pg1-path` option missing for `archive-get` command. 

- Add hint when checksum delta is enabled after a timeline switch. 

- Use PostgreSQL instead of `postmaster` where appropriate. 

### Core Development

- Automatically determine cipher passphrase in `repo-get` command. 

- Fix expression when recursion enabled in `storageInfoListP()`. 

- Improve behavior of the `repo-ls` command. 

- Inline `strPtr()` to increase profiling accuracy. 

- Add `pgLsnFromWalSegment()`. 

### Documentation Bug Fixes

- Fix incorrect example for `repo-retention-full-type` option. 

- Remove internal commands from HTML and man command references. 

### Documentation Improvements

- Update PostgreSQL versions used to build user guides. Also add version ranges to indicate that a user guide is accurate for a range of PostgreSQL versions even if it was built for a specific version. 

- Update FAQ for expiring a specific backup set. 

- Update FAQ to clarify default PITR behavior. 

### Test Development

- Remove `real/all` integration tests now covered by unit tests. 

- Rename most instances of master to primary in tests. 

## v2.27 — Expiration Improvements and Compression Drivers

*Released: 2020-05-26*

### Core Bug Fixes

- Fix issue checking if file links are contained in path links. 

- Allow `pg-path1` to be optional for synchronous `archive-push`. 

- The `expire` command now checks if a stop file is present. 

- Handle missing reason phrase in HTTP response. 

- Increase buffer size for lz4 compression flush. 

- Ignore `pg-host*` and `repo-host*` options for the `remote` command. 

- Fix possibly missing `pg1-*` options for the `remote` command. 

### Core Features

- Time-based retention for full backups. The `--repo-retention-full-type` option allows retention of full backups based on a time period, specified in days. 

- Ad hoc backup expiration. Allow the user to remove a specified backup regardless of retention settings. 

- Zstandard compression support. Note that setting `compress-type=zst` will make new backups and archive incompatible (unrestorable) with prior versions of pgBackRest. 

- bzip2 compression support. Note that setting `compress-type=bz2` will make new backups and archive incompatible (unrestorable) with prior versions of pgBackRest. 

- Add `backup`/`expire` running status to the `info` command. 

### Core Improvements

- Expire WAL archive only when `repo-retention-archive` threshold is met. WAL prior to the first full backup was previously expired after the first full backup. Now it is preserved according to retention settings. 

- Add local MD5 implementation so S3 works when FIPS is enabled. 

- PostgreSQL 13 beta1 support. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

- Reduce `buffer-size` default to `1MiB`. 

- Throw user-friendly error if `expire` is not run on repository host. 

### Core Development

- Enforce non-null for most string options. 

## v2.26 — Non-blocking TLS

*Released: 2020-04-20*

### Core Bug Fixes

- Remove empty subexpression from manifest regular expression. MacOS was not happy about this though other platforms seemed to work fine. 

### Core Improvements

- Non-blocking TLS implementation. 

- Only limit backup copy size for WAL-logged files. The prior behavior could possibly lead to `postgresql.conf` or `postgresql.auto.conf` being truncated in the backup. 

- TCP keep-alive options are configurable. 

- Add `io-timeout` option. 

### Core Development

- Simplify storage driver info and list functions. 

- Split session functionality of `SocketClient` out into `SocketSession`. 

- Split session functionality of `TlsClient` out into `TlsSession`. 

- Use `poll()` instead of `select()` for monitoring socket read/write ready. 

- Use `SocketSession`/`TlsSession` for test servers. 

- Always throw error when OpenSSL returns `SSL_ERROR_SYSCALL`. 

- Use `__noreturn_` on error functions when coverage testing. 

## v2.25 — LZ4 Compression Support

*Released: 2020-03-26*

### Core Features

- Add lz4 compression support. Note that setting `compress-type=lz4` will make new backups and archive incompatible (unrestorable) with prior versions of pgBackRest. 

- Add `--dry-run` option to the `expire` command. Use dry-run to see which backups/archive would be removed by the `expire` command without actually removing anything. 

### Core Improvements

- Improve performance of remote manifest build. 

- Fix detection of keepalive options on Linux. 

- Add configure host detection to set standards flags correctly. 

- Remove `compress`/`compress-level` options from commands where unused. These commands (e.g. `restore`, `archive-get`) never used the compress options but allowed them to be passed on the command line. Now they will error when these options are passed on the command line. If these errors occur then remove the unused options. 

- Limit backup file copy size to size reported at backup start. If a file grows during the backup it will be reconstructed by WAL replay during recovery so there is no need to copy the additional data. 

### Core Development

- Add infrastructure for multiple compression type support. 

- Improve performance of `MEM_CONTEXT*()` macros. 

- Allow storage reads to be limited by bytes. 

## v2.24 — Auto-Select Backup Set for Time Target

*Released: 2020-02-25*

### Core Bug Fixes

- Prevent defunct processes in asynchronous archive commands. 

- Error when `archive-get`/`archive-push`/`restore` are not run on a PostgreSQL host. 

- Read HTTP content to eof when size/encoding not specified. 

- Fix resume when the resumable backup was created by Perl. In this case the resumable backup should be ignored, but the C code was not able to load the partial manifest written by Perl since the format differs slightly. Add validations to catch this case and continue gracefully. 

### Core Features

- Auto-select backup set on restore when time target is specified. Auto-selection is performed only when `--set` is not specified. If a backup set for the given target time cannot not be found, the latest (default) backup set will be used. 

### Core Improvements

- Skip `pg_internal.init` temp file during backup. 

- Add more validations to the manifest on `backup`. 

### Documentation Improvements

- Prevent lock-bot from adding comments to locked issues. 

## v2.23 — Bug Fix

*Released: 2020-01-27*

### Core Bug Fixes

- Fix missing files corrupting the manifest. If a file was removed by PostgreSQL during the backup (or was missing from the standby) then the next file might not be copied and updated in the manifest. If this happened then the backup would error when restored. 

### Core Improvements

- Use `pkg-config` instead of `xml2-config` for libxml2 build options. 

- Validate checksums are set in the manifest on `backup`/`restore`. 

## v2.22 — Bug Fix

*Released: 2020-01-21*

### Core Bug Fixes

- Fix error in timeline conversion. The timeline is required to verify WAL segments in the archive after a backup. The conversion was performed base `10` instead of `16`, which led to errors when the timeline was &ge; `0xA`. 

## v2.21 — C Migration Complete

*Released: 2020-01-15*

### Core Bug Fixes

- Fix options being ignored by asynchronous commands. The asynchronous `archive-get`/`archive-push` processes were not loading options configured in command configuration sections, e.g. `[global:archive-get]`. 

- Fix handling of `\` in filenames. `\` was not being properly escaped when calculating the manifest checksum which prevented the manifest from loading. Since instances of `\` in cluster filenames should be rare to nonexistent this does not seem likely to be a serious problem in the field. 

### Core Features

- pgBackRest is now pure C. 

- Add `pg-user` option. Specifies the database user name when connecting to PostgreSQL. If not specified pgBackRest will connect with the local OS user or `PGUSER`, which was the previous behavior. 

- Allow path-style URIs in S3 driver. 

### Core Improvements

- The `backup` command is implemented entirely in C. 

### Core Development

- Add basic time management functions. 

- Add `httpLastModifiedToTime()` to parse HTTP `last-modified` header. 

- Parse dates in `storageS3InfoList()` and `storageS3Info()`. 

## v2.20 — Bug Fixes

*Released: 2019-12-12*

### Core Bug Fixes

- Fix archive-push/archive-get when `PGDATA` is symlinked. These commands tried to use `cwd()` as `PGDATA` but this would disagree with the path configured in pgBackRest if `PGDATA` was symlinked. If `cwd()` does not match the pgBackRest path then `chdir()` to the path and make sure the next `cwd()` matches the result from the first call. 

- Fix reference list when `backup.info` is reconstructed in `expire` command. Since the `backup` command is still using the Perl version of reconstruct this issue will not express unless **1)** there is a backup missing from `backup.info` and **2)** the `expire` command is run directly instead of running after `backup` as usual. This unlikely combination of events means this is probably not a problem in the field. 

- Fix segfault on unexpected EOF in gzip decompression. 

### Core Development

- Add manifest build for new backups. 

## v2.19 — C Migrations and Bug Fixes

*Released: 2019-11-12*

### Core Bug Fixes

- Fix remote timeout in delta restore. When performing a delta restore on a largely unchanged cluster the remote could timeout if no files were fetched from the repository within `protocol-timeout`. Add keep-alives to prevent remote timeout. 

- Fix handling of repeated HTTP headers. When HTTP headers are repeated they should be considered equivalent to a single comma-separated header rather than generating an error, which was the prior behavior. 

### Core Improvements

- JSON output from the `info` command is no longer pretty-printed. Monitoring systems can more easily ingest the JSON without linefeeds. External tools such as `jq` can be used to pretty-print if desired. 

- The `check` command is implemented entirely in C. 

### Documentation Improvements

- Document how to contribute to pgBackRest. 

- Document maximum version for `auto-stop` option. 

### Test Improvements

- Fix container test path being used when `--vm=none`. 

- Fix mismatched timezone in expect test. 

- Don't autogenerate embedded libc code by default. 

## v2.18 — PostgreSQL 12 Support

*Released: 2019-10-01*

### Core Features

- PostgreSQL 12 support. 

- Add `info` command `set` option for detailed text output. The additional details include databases that can be used for selective restore and a list of tablespaces and symlinks with their default destinations. 

- Add `standby` restore type. This restore type automatically adds `standby_mode=on` to recovery.conf for PostgreSQL < 12 and creates `standby.signal` for PostgreSQL &ge; 12, creating a common interface between PostgreSQL versions. 

### Core Improvements

- The `restore` command is implemented entirely in C. 

### Core Development

- Migrate backup manifest load/save to C. 

- Improve performance of info file load/save. 

- Add helper function for adding `CipherBlock` filters to groups. 

### Documentation Improvements

- Document the relationship between `db-timeout` and `protocol-timeout`. 

- Add documentation clarifications regarding standby repositories. 

- Add FAQ for time-based Point-in-Time Recovery. 

## v2.17 — C Migrations and Bug Fixes

*Released: 2019-09-03*

### Core Bug Fixes

- Improve slow manifest build for very large quantities of tables/segments. 

- Fix exclusions for special files. 

### Core Improvements

- The `stanza-create/update/delete` commands are implemented entirely in C. 

- The `start`/`stop` commands are implemented entirely in C. 

- Create log directories/files with `0750`/`0640` mode. 

### Core Development

- Move info file checksum to the end of the file. 

### Documentation Bug Fixes

- Fix `yum.p.o` package being installed when custom package specified. 

### Documentation Improvements

- Build pgBackRest as an unprivileged user. 

## v2.16 — C Migrations and Bug Fixes

*Released: 2019-08-05*

### Core Bug Fixes

- Retry S3 `RequestTimeTooSkewed` errors instead of immediately terminating. 

- Fix incorrect handling of `transfer-encoding` response to `HEAD` request. 

- Fix scoping violations exposed by optimizations in gcc 9. 

### Core Features

- Add `repo-s3-port` option for setting a non-standard S3 service port. 

### Core Improvements

- The `local` command for `backup` is implemented entirely in C. 

- The `check` command is implemented partly in C. 

### Core Development

- Add Perl interface to C storage layer. 

- Add `Db` object to encapsulate PostgreSQL queries and commands. 

- Add PostgreSQL query client. 

## v2.15 — C Implementation of Expire

*Released: 2019-06-25*

### Core Bug Fixes

- Fix archive retention expiring too aggressively. 

### Core Improvements

- The `expire` command is implemented entirely in C. 

- The `local` command for restore is implemented entirely in C. 

- Remove hard-coded PostgreSQL user so `$PGUSER` works. 

- Honor configure `--prefix` option. 

- Rename `repo-s3-verify-ssl` option to `repo-s3-verify-tls`. The new name is preferred because pgBackRest does not support any SSL protocol versions (they are all considered to be insecure). The old name will continue to be accepted. 

### Core Development

- Add most unimplemented functions to the remote storage driver. 

- Rename `info*New()` functions to `info*NewLoad()`. 

- Add backup management functions to `InfoBackup`. 

### Documentation Improvements

- Add FAQ to the documentation. 

- Use `wal_level=replica` in the documentation for PostgreSQL &ge; 9.6. 

## v2.14 — Bug Fix and Improvements

*Released: 2019-05-20*

### Core Bug Fixes

- Fix segfault when `process-max` > 8 for `archive-push`/`archive-get`. 

### Core Improvements

- Bypass database checks when `stanza-delete` issued with `force`. 

- Add `configure` script for improved multi-platform support. 

### Core Development

- Filter improvements. Only process next filter in `IoFilterGroup` when input buffer is full or flushing. Improve filter's notion of "done" to optimize filter processing. 

- Improve performance of non-blocking reads by using maximum buffer size. 

- Add `storageInfoList()` to get detailed info about all entries in a path. 

- Allow `storageInfo()` to follow links. 

- Allow `StorageFileWrite` to set user, group, and modification time. 

- Add `pathExists()` to `Storage` object. 

- Improve zero-length content handling in `HttpClient` object. 

- Don't append `strerror()` to error message when `errno` is 0. 

- Improve type safety of interfaces and drivers. 

- Add `--c` option to request a C remote. 

- Add `common/macro.h` for general-purpose macros. 

- Add macros for object free functions. 

- Various `MemContext` callback improvements. 

- Various `Buffer` improvements. 

- Simplify storage object names. 

- Add `ioWriteStr()` and `ioWriteStrLine()`. 

- Add separate functions to encode/decode each JSON type. 

- Add macros to create constant `Buffer` objects. 

- Add missing `httpUriEncode()` in S3 request. 

- Add `unsigned int` `Variant` type and update code to use it. 

- Expose handle (file descriptor) from `IoWrite` when applicable. 

- Add `iniSave()` and `iniMove()` to `Ini` object. 

- Add `*Save()` functions to most `Info` objects. 

- Extern `infoHash()` so it can be used by other modules. 

- `varNewKv()` accepts a `KeyValue` object rather than creating one. 

- Add constant for maximum buffer sizes required by `cvt*()` functions. 

- Add `true` and `false` `String` constants. 

- Refactor `Ini` interface to expose `String` values instead of `Variant`. 

- Refactor `main()` as a `switch()` statement. 

- Add `cfgOptionUInt()` and `cfgOptionUInt64()` and update code to use them. 

- Improve log performance, simplify macros, rename `logWill()` to `logAny()`. 

- Improve coverage in `perl/exec`, `config/config`, and `config/parse` modules. 

- Remove `-Wswitch-enum` compiler option. 

- Error on multiple option alternate names and simplify help command. 

- Use `THROW_ON_SYS_ERROR*()` to improve code coverage. 

- Improve macros and coverage rules that were hiding missing coverage. 

- Improve efficiency of `FUNCTION_LOG*()` macros. 

### Documentation Features

- Add user guides for CentOS/RHEL 6/7. 

### Documentation Development

- Automate coverage summary report generation. 

- Add `--out-preserve` to preserve contents of output path. 

- Restore index menu url default lost in b85e51d6. 

### Test Development

- Add `harnessInfoChecksum/Z()` to ease creation of test info files. 

- Update containers with PostgreSQL minor releases and `liblz4`. 

- Add `testUser()` and `testGroup()`. 

- Add `build-max` option to set max build processes. 

- Reduce ScalityS3 processes since only two are needed. 

- Update `mock/expire` module test matrix so expect tests output. 

## v2.13 — Bug Fixes

*Released: 2019-04-18*

### Core Bug Fixes

- Fix zero-length reads causing problems for IO filters that did not expect them. 

- Fix reliability of error reporting from `local`/`remote` processes. 

- Fix Posix/CIFS error messages reporting the wrong filename on write/sync/close. 

### Core Development

- Harden IO filters against zero input and optimize zero output case. 

- Move `lockRelease()` to the end of `exitSafe()`. 

- Add `CHECK()` macro for production assertions. 

- Automatically generate constants for command and option names. 

- Use a macro instead of a nested struct to create common `String` variables. 

- Add `STR()` macro to create constant `String` objects from runtime strings. 

- Add macros to create constant `Variant` types. 

- Migrate `backupRegExp()` to C. 

### Documentation Development

- Option to build documentation from current apt.postgres.org packages. 

## v2.12 — C Implementation of Archive Push

*Released: 2019-04-11*

### Core Bug Fixes

- Fix issues when a path option is / terminated. 

- Fix issues when `log-level-file=off` is set for the `archive-get` command. 

- Fix C code to recognize `host:port` option format like Perl does. 

- Fix issues with `remote`/`local` command logging options. 

### Core Improvements

- The `archive-push` command is implemented entirely in C. 

- Increase `process-max` limit to `999`. 

- Improve error message when an S3 bucket name contains dots. 

### Core Development

- Add separate `archive-push-async` command. 

- `CryptoHash` improvements and fixes. Fix incorrect buffer size used in `cryptoHashOne()`. Add missing `const` to `cryptoHashOne()` and `cryptoHashOneStr()`. Add hash size constants. Extern hash type constant. 

- Add `CIFS` storage driver. 

- Add file write to the remote and S3 storage drivers. 

- Add `storageRepoWrite()` and `storagePg()`/`storagePgWrite()` to storage helper. 

- Use a single file to handle global errors in async archiving. 

- Add document creation to XML objects. 

- Remove redundant documentation from PostgreSQL interface files and clarify ambiguous function names. 

- Add WAL info to PostgreSQL interface. 

- Refactor PostgreSQL interface to remove most code duplication. 

- Logging improvements. Allow three-digit process IDs in logging. Allow process id in C logging. 

- Add process id to `ProtocolParallelJob`. 

- Add process id to C `archive-get` and `archive-push` logging. 

- Close log file before `exec()`. 

- Allow warnings to be written by `archiveAsyncStatusOkWrite()`. 

- Move WAL path prefix logic into `walPath()`. 

- Make notion of current PostgreSQL info ID in C align with Perl. 

- Add locking capability to the remote command. 

- Add `forkSafe()` to handle fork errors. 

- Add `httpHeaderDup()`. 

- `httpClientRequest()` accepts a body parameter. 

- Add `protocolKeepAlive()` to send `noops` to all remotes. 

- Make `strLstDup()` null-tolerant. 

- Add `strLstMergeAnti()` for merge anti-joins. 

- Add `cvtSSizeToZ()` and debug macros. 

- Remove unused `infoArchiveCheckPg()` function. 

- Add constants for `.ok`/`.error` status extensions. 

### Documentation Improvements

- Clarify that S3-compatible object stores are supported. 

### Test Development

- Build test harness with the same warnings as code being tested. 

- Add `TEST_64BIT()` macro to detect 64-bit platforms. 

- Skip coverage for macros with numbers in their name. 

- Use `restore` command for remote performances tests. 

## v2.11 — C Implementation of Archive Get

*Released: 2019-03-11*

### Core Bug Fixes

- Fix possible truncated WAL segments when an error occurs mid-write. 

- Fix info command missing WAL min/max when stanza specified. 

- Fix non-compliant JSON for options passed from C to Perl. 

### Core Improvements

- The `archive-get` command is implemented entirely in C. 

- Enable socket keep-alive on older Perl versions. 

- Error when parameters are passed to a command that does not accept parameters. 

- Add hints when unable to find a WAL segment in the archive. 

- Improve error when hostname cannot be found in a certificate. 

- Add additional options to `backup.manifest` for debugging purposes. 

### Core Development

- Migrate `local` and `remote` commands to C. 

- Add separate `archive-get-async` command. 

- Add `ProtocolParallel*` objects for parallelizing commands. 

- Add `ProtocolCommand` object. 

- Add `exists()` to remote storage. 

- Resolve storage path expressions before passing to remote. 

- Expose handle (file descriptor) from `IoRead` when applicable. 

- Add `storageHelperFree()` to storage helper. 

- Add `kvKeyExists()` to `KeyValue` object. 

- Add `lstRemove()` to `List` object. 

- Allow `cfgExecParam()` to exclude options. 

- `MemContext` improvements. Improve performance of context and memory allocations. Use `contextTop`/`contextCurrent` instead of `memContextTop()`/`memContextCurrent()`. Don't make a copy of the context name. 

- Make `DESTDIR` fully-configurable in the `Makefile`. 

- Add note for `CSTD` settings on BSD variants. 

- Add `clean` and `uninstall` targets to `Makefile`. 

- Change `execRead()` to return a `size_t`. 

- Prevent option warning from being output when running help command. 

- Improve null-handling of `strToLog()` and `varToLog()`. 

- Increase per-call stack trace size to `4096`. 

- Move `compress` module to `common/compress`. 

- Move `crypto` module to `common/crypto`. 

### Documentation Improvements

- Update default documentation version to PostgreSQL 10. 

### Documentation Development

- Add instructions for building the coverage report. 

- Documentation builds on PostgreSQL 9.4-10. 

### Test Development

- Create test matrix for `mock/archive`, `mock/archive-stop`, `mock/all`, `mock/expire`, and `mock/stanza` to increase coverage and reduce tests. 

- Improve fork harness to allow multiple children and setup pipes automatically. 

- Reduce expect log level in `mock/archive` and `mock/stanza` tests. 

- Rename test modules for consistency. 

- Only run test-level stack trace by default for unit-tested modules. 

- Add missing ToLog() coverage to `String`, `List`, and `PgControl`. 

- Create aliases for test VMs ordered by age. 

## v2.10 — Bug Fixes

*Released: 2019-02-09*

### Core Bug Fixes

- Add unimplemented S3 driver method required for `archive-get`. 

- Fix check for improperly configured `pg-path`. 

### Core Development

- JSON improvements. Optimize parser implementation. Make the renderer more null tolerant. 

- Automatically adjust `db-timeout` when `protocol-timeout` is smaller. 

## v2.09 — Minor Improvements and Bug Fixes

*Released: 2019-01-30*

### Core Bug Fixes

- Fix issue with multiple async status files causing a hard error. 

### Core Improvements

- The `info` command is implemented entirely in C. 

- Simplify `info` command text message when no stanzas are present. Replace the repository path with "the repository". 

- Add `_DARWIN_C_SOURCE` flag to Makefile for MacOS builds. 

- Update address lookup in C TLS client to use modern methods. 

- Include Posix-compliant header for `strcasecmp()` and `fd_set`. 

### Core Development

- Add remote storage objects. 

- Add `ProtocolClient` object and helper functions. 

- Add `Exec` object. 

- Add `IoHandleRead` and `IoHandleWrite` objects. 

- Add `cfgExecParam()` to generate parameters for executing commands. 

- Ignore `SIGPIPE` signals and check `EPIPE` result instead. 

- Function log macro improvements. Rename FUNCTION_DEBUG_* and consolidate ASSERT_* macros for consistency. Improve `CONST` and `P`/`PP` type macro handling. Move `MACRO_TO_STR()` to `common/debug.h`. Remove unused type parameter from `FUNCTION_TEST_RETURN()`. 

- Make the C version of the `info` command conform to the Perl version. 

- Improve accuracy of `strSizeFormat()`. 

- Add `ioReadBuf()` to easily read into a buffer. 

- JSON improvements. Allow empty arrays in JSON parser. Fix null output in JSON renderer. Fix escaping in JSON string parser/renderer. 

- Allocate extra space for concatenations in the `String` object. 

- Return `UnknownError` from `errorTypeFromCode()` for invalid error codes. 

- Update Perl repo rules to work when stanza is not specified. 

- Update `Storage::Local->list()` to accept an undefined path. 

- Null-terminate list returned by `strLstPtr()`. 

- Add `kvMove()` and `varLstMove()`. 

- Replace `FileOpenError` with `HostConnectError` in `TlsClient`. 

- Allow string `Variant` objects to contain `null`. 

- Rename `common/io/handle` module to `common/io/handleWrite`. 

- Add `const VariantList *` debug type. 

### Documentation Bug Fixes

- Fix hard-coded repository path. 

### Documentation Improvements

- Clarify that encryption is always performed client-side. 

- Add examples for building a documentation host. 

- Allow `if` in manifest variables, lists, and list items. 

### Test Development

- Move C module include in `test.c` above headers included for testing. 

- Allow primary `gid` for the test user to be different from `uid`. 

- Increase timeout in `storage/s3` module to improve reliability. 

## v2.08 — Minor Improvements and Bug Fixes

*Released: 2019-01-02*

### Core Bug Fixes

- Remove request for S3 object info directly after putting it. 

- Correct `archive-get-queue-max` to be `size` type. 

- Add error message when current user `uid`/`gid` does not map to a name. 

- Error when `--target-action=shutdown` specified for PostgreSQL < 9.5. 

### Core Improvements

- Set TCP keepalives on S3 connections. 

- Reorder `info` command text output so most recent backup is output last. 

- Change file ownership only when required. 

- Redact `authentication` header when throwing S3 errors. 

### Core Development

- Enable S3 storage and encryption for `archive-get` command in C. 

- Migrate local `info` command to C. 

- Add S3 storage driver. 

- Add `HttpClient` object. 

- Add `TlsClient` object. 

- Add interface objects for libxml2. 

- Add encryption capability to `Info*` objects. 

- Add `IoFilter` interface to `CipherBlock` object. 

- Allow arbitrary `InOut` filters to be chained in `IoFilterGroup`. 

- Add `infoBackup` object to encapsulate the `backup.info` file. 

- Improve JSON to `Variant` conversion and add `Variant` to JSON conversion. 

- Storage helper improvements. Allow `NULL` stanza in storage helper. Add path expression for repository backup. 

- Info module improvements. Rename constants in `Info` module for consistency. Remove `#define` statements in the `InfoPg` module to conform with newly-adopted coding standards. Use cast to make for loop more readable in `InfoPg` module. Add `infoArchiveIdHistoryMatch()` to the `InfoArchive` object. 

- Allow I/O read interface to explicitly request blocking reads. 

- Improve error messages when info files are missing/corrupt. 

- Add `strSizeFormat()` to `String` object. 

- Add `strLstInsert()` and `strLstInsertZ()` to `StringList` object. 

- Rename `PGBACKREST`/`BACKREST` constants to `PROJECT`. 

- Require S3 key options except for `local`/`remote` commands. 

- Explicitly compile with Posix 2001 standard. 

- Add `ServiceError` for errors from a service that can be retried. 

- Conditional compilation of Perl logic in `exit.c`. 

- Merge `cipher.h` into `crypto.h`. 

- Remove extraneous `use`/`include` statements. 

- Remove embedded semicolon from `String` constant macros. 

- Reduce debug level for `infoIni()` to test. 

- Return `IoFilterGroup *` from `ioFilterGroupAdd()`. 

- Add coding standards for `String` constants. 

- Add missing `LOG_DEBUG()` macro. 

### Documentation Improvements

- Clarify when `target-action` is effective and PostgreSQL version support. 

- Clarify that region/endpoint must be configured correctly for the bucket. 

- Add documentation for building the documentation. 

### Documentation Development

- Add `admonitions` to all documentation renderers (HTML, PDF, Markdown and help text) and update `xml` files accordingly. 

- Add HTML table rendering and update PDF/Markdown renderers to support header-less tables. Add optional table captions. 

- Escape special characters in latex when not in a code block. 

- Base menu ordering on natural ordering in the manifest. 

- Replace keywords with more flexible if statements. 

- Pre-build containers for any `execute` elements marked `pre`. 

- Documentation may be built with user-specified packages. 

- Add Centos/RHEL 7 option to documentation build. 

- Allow custom logo for PDF documentation. 

- Modify general document elements to allow any child element. 

- Use absolute paths so that `./doc.pl` runs. 

- Pick `pg_switch_wal()`/`pg_switch_xlog()` based on PostgreSQL version. 

- Add configuration to the standby so it works as a primary when promoted. 

- Create common `if` expressions for testing `os-type`. 

- Add `zlib1g-dev` to Debian builds. 

### Test Development

- New test containers with static test certificates. 

- Fix test binary name for `gprof`. 

- Allow arbitrary multiplier and flush character in `IoTestFilterMultiply`. 

- Update URL for Docker install. 

- Add `testRepoPath()` to let C unit tests know where the code repository is located. 

- Merge `common/typeStringListTest` module into `common/typeStringTest`. 

- Merge `common/typeVariantListTest` module into `common/typeVariantTest`. 

## v2.07 — Automatic Backup Checksum Delta

*Released: 2018-11-16*

### Core Bug Fixes

- Fix issue with `archive-push-queue-max` not being honored on connection error. 

- Fix static WAL segment size used to determine if `archive-push-queue-max` has been exceeded. 

- Fix error after log file open failure when processing should continue. 

### Core Features

- Automatically enable backup checksum delta when anomalies (e.g. timeline switch) are detected. 

### Core Improvements

- Retry all S3 `5xx` errors rather than just `500` internal errors. 

### Core Development

- Correct current history item in `InfoPg` to always be in position 0. 

- Make `ioReadLine()` read less aggressively. 

- Add `ioWriteFlush()` to flush pending output. 

- Add destructors to `IoRead` and `IoWrite` objects. 

- Add `base` variants to all integer to string conversion functions. 

- Add `lstInsert()` to `List` object. 

- Add `strCatChr()`, `strEmpty()`, and constant macros to `String` object. 

- Add `regExpPrefix()` to aid in static prefix searches. 

- Correct `cfgDefDataFind()` to use `UINTP` instead of `VOIDPP`. 

- Change `infoArchiveCheckPg()` to display the PostgreSQL version as a string (e.g. 9.4) instead of the integer representation (e.g. 90400) when throwing an error. 

- Allow storage path and file mode to be 0. 

- Limit usable `Buffer` size without changing allocated size. 

- Construct `Wait` object in milliseconds instead of fractional seconds. 

- Add `THROW*_ON_SYS_ERROR*` macros to test and throw system errors. 

- Add `KernelError` to report miscellaneous kernel errors. 

- Use `THROW_ON_SYS_ERROR` macro to improve `fork` code coverage. 

- `Storage` interface methods no longer declare the driver as const. 

- Add `memContextCallbackClear()` to prevent double `free()` calls. 

- Merge `crypto/random` module into `crypto/crypto`. 

- Add `cryptoError()` and update crypto code to use it. 

- Rename `CipherError` to `CryptoError`. 

- Reword misleading message in stack trace when parameter buffer is full. 

- Add logging macros for `TimeMSec` type. 

- Modify value of `PERL_EMBED_ERROR` macro. 

### Documentation Development

- Add new HTML tags and `strExtra` to `DocHtmlElement`. 

- Remove error suppression for pgBackRest `make`. 

### Test Development

- New test containers. Add libxml2 library needed for S3 development. Include new minor version upgrades for PostgreSQL. Remove PostgreSQL 11 beta/rc repository. 

- Test speed improvements. Mount `tmpfs` in `Vagrantfile` instead `test.pl`. Preserve contents of C unit test build directory between `test.pl` executions. Improve efficiency of code generation. 

- New, concise coverage report for C. 

- Add `TEST_LOG()` and `TEST_LOG_FMT()` macros. 

- Improve alignment of expected vs. actual error test results. 

- Add time since the beginning of the run to each test statement. 

## v2.06 — Checksum Delta Backup and PostgreSQL 11 Support

*Released: 2018-10-15*

### Core Bug Fixes

- Fix missing URI encoding in S3 driver. 

- Fix incorrect error message for duplicate options in configuration files. 

- Fix incorrectly reported error return in `info` logging. A return code of 1 from the `archive-get` was being logged as an error message at `info` level but otherwise worked correctly. 

### Core Features

- Add checksum delta for incremental backups. Checksum delta backups uses checksums rather than timestamps to determine if files have changed. 

- PostgreSQL 11 support, including configurable WAL segment size. 

### Core Improvements

- Ignore all files in a linked tablespace directory except the subdirectory for the current version of PostgreSQL. Previously an error would be generated if other files were present and not owned by the PostgreSQL user. 

- Improve `info` command to display the stanza cipher type. 

- Improve support for special characters in filenames. 

- Allow `delta` option to be specified in the pgBackRest configuration file. 

### Core Development

- Migrate local, unencrypted, non-S3 `archive-get` command to C. 

- Storage refactoring. Posix file functions now differentiate between open and missing errors. Don't use negations in objects below Storage. Rename posix driver files/functions for consistency. Full abstraction of storage driver interface. Merge protocol storage helper into storage helper. Add CIFS driver to storage helper for read-only repositories. 

- Update all interfaces to use variable parameter constructors. 

- Info objects now parse JSON and use specified storage. 

- Add `ioReadLine()`/`ioWriteLine()` to `IoRead`/`IoWrite` objects. 

- Add helper for repository storage. 

- Add `cryptoHmacOne()` for HMAC support. 

- Add `cfgDefOptionMulti()` to identify multi-value options. 

- Add `bufNewZ()` and `bufHex()` to `Buffer` object. 

- Allow `hashSize()` to run on remote storage. 

- Restore `bIgnoreMissing` flag in `backupFile()` lost in storage refactor. 

- Migrate `walIsPartial()`, `walIsSegment()`, and `walSegmentFind()` from Perl to C. 

- Migrate control functions to detect stop files to C. 

- Make archive-get info messages consistent between C and Perl implementations. 

- Change locking around async process forking to be more test friendly. 

- Simplify debug logging by allowing log functions to return `String` objects. 

- Improve documentation in `filter.h` and `filter.internal.h`. 

### Documentation Improvements

- Use `command` in `authorized_hosts` to improve SSH security. 

- List allowable values for the `buffer-size` option in the configuration reference. 

### Documentation Development

- Update introduction to be pithy. 

### Test Development

- Install nodejs from `deb.nodesource.com`. 

- Disable flapping `archive/get` unit on CentOS 6. 

- Move test expect log out of the regular test directory. 

- Fix buffer underrun in configuration test harness. 

- Fix missing test caused by a misplaced YAML tag. 

- Make Valgrind return an error even when a non-fatal issue is detected. Update some minor issues discovered in the tests as a result. 

- Add `-ftree-coalesce-vars` option to unit test compilation. 

- Clear test directory between test runs. 

- Allow C or Perl coverage to run on more than one VM. 

- Don't perform valgrind when requested. 

- Remove compiler warnings that are not valid for u16. 

- Merge all posix storage tests into a single unit. 

- Add `.gitignore` to C `src` directory. 

- Fix typo in unit test error messages, EXECTED => EXPECTED. 

- Make comment blocks consistent across all tests. 

## v2.05 — Environment Variable Options and Exclude Temporary/Unlogged Relations

*Released: 2018-08-31*

### Core Bug Fixes

- Fix issue where *relative* links in `$PGDATA` could be stored in the backup with the wrong path. This issue did not affect absolute links and relative tablespace links were caught by other checks. 

- Remove incompletely implemented `online` option from the `check` command. Offline operation runs counter to the purpose of this command, which is to check if archiving and backups are working correctly. 

- Fix issue where errors raised in C were not logged when called from Perl. pgBackRest properly terminated with the correct error code but lacked an error message to aid in debugging. 

- Fix issue when a boolean option (e.g. `delta`) was specified more than once. 

### Core Features

- Allow any option to be set in an environment variable. This includes options that previously could only be specified on the command line, e.g. `stanza`, and secret options that could not be specified on the command-line, e.g. `repo1-s3-key-secret`. 

- Exclude temporary and unlogged relation (table/index) files from backup. Implemented using the same logic as the patches adding this feature to PostgreSQL, [8694cc96](https://git.postgresql.org/pg/commitdiff/8694cc96b52a967a49725f32be7aa77fd3b6ac25) and [920a5e50](https://git.postgresql.org/pg/commitdiff/920a5e500a119b03356fb1fb64a677eb1aa5fc6f). Temporary relation exclusion is enabled in PostgreSQL &ge; `9.0`. Unlogged relation exclusion is enabled in PostgreSQL &ge; `9.1`, where the feature was introduced. 

- Allow arbitrary directories and/or files to be excluded from a backup. Misuse of this feature can lead to inconsistent backups so read the `--exclude` documentation carefully before using. 

- Add `log-subprocess` option to allow file logging for `local` and `remote` subprocesses. 

- PostgreSQL 11 Beta 3 support. 

### Core Improvements

- Allow zero-size files in backup manifest to reference a prior manifest regardless of timestamp delta. 

- Improve asynchronous `archive-get`/`archive-push` performance by directly checking status files. 

- Improve error message when a command is missing the `stanza` option. 

### Core Development

- Validate configuration options in a single pass. By pre-calculating and storing the option dependencies in `parse.auto.c` validation can be completed in a single pass, which is both simpler and faster. 

- Add gzip compression/decompression filters for C. 

- Improve performance of string to int conversion. Use `strtoll()` instead of `sprintf()` for conversion. Also use available integer min/max constants rather than hard-coded values. 

- Add `uint64` variant type and supporting conversion functions. 

- Add basic C JSON parser. 

- Migrate minimum set of code for reading `archive.info` files from Perl to C. 

- Allow `Buffer` object "used size" to be different than "allocated size". Add functions to manage used size and remaining size and update automatically when possible. 

- Abstract IO layer out of the storage layer. This allows the routines to be used for IO objects that do not have a storage representation. Implement buffer read and write IO objects. Implement filters and update `cryptoHash` to use the new interface. Implement size and buffer filters. 

- `storageFileRead()` accepts a buffer for output rather than creating one. This is more efficient overall and allows the caller to specify how many bytes will be read on each call. Reads are appended if the buffer already contains data but the buffer size will never increase. 

- Add `iniSectionList()` to `Ini` object and remove dead code. 

- Manifest improvements. Require PostgreSQL catalog version when instantiating a `Manifest` object (and not loading it from disk). Prevent manifest from being built more than once. Limit manifest build recursion (i.e. links followed) to sixteen levels to detect link loops. 

- Do nothing in `memContextMove()` when the context is already in the specified parent. 

- Allow command/option constants to autonumber in both C and Perl to reduce churn when a new command/option is added. 

- Show exact log level required for stack trace param output instead of just "debug". 

- Update `Archive::Info->archiveIdList()` to return a valid error code instead of unknown. 

- Add `cvtBoolToConstZ()` to simplify conversion of boolean to string. 

- Add `cvtZToUInt()` to convert string to unsigned int. 

- Enable `-Wstrict-prototypes`, `-Wpointer-arith`, `-Wduplicated-branches`, `-Wvla`, and `-Wduplicated-cond` and update code to conform where necessary. 

- Rename error-handling variables in `Main.pm` to conform to standard. 

- Remove redundant lines from embedded Perl by combining blank lines. 

- Define cipher magic size with `sizeof()` rather than using a constant. 

- Add `cvtCharToZ()` and macro for debugging `char` params. 

- Add `strReplaceChr()` to `String` object. 

- Correct `OptionInvalidError` to `OptionInvalidValueError` in boolean option parsing. 

### Documentation Bug Fixes

- Fix invalid log level in `log-path` option reference. 

### Documentation Improvements

- Stop trying to arrange contributors in `release.xml` by last/first name. Contributor names have always been presented in the release notes exactly as given, but we tried to assign internal IDs based on last/first name which can be hard to determine and ultimately doesn't make sense. Inspired by Christophe's PostgresOpen 2017 talk, "Human Beings Do Not Have a Primary Key". 

### Documentation Development

- Allow containers to be defined in a document. The defined containers are built before the document build begins which allows them to be reused. 

- Move most host setup to containers defined in the documentation. This includes PostgreSQL installation which had previously been included in the documentation. This way produces faster builds and there is no need for us to document PostgreSQL installation. 

### Test Improvements

- Error if LibC build is performed outside the test environment. LibC is no longer required for production builds. 

### Test Development

- Use pre-built images from Docker Hub when the container definition has not changed. Downloading an image is quite a bit faster than building a new image from scratch and saves minutes per test run in CI. 

- Refactor the `common/log` tests to not depend on `common/harnessLog`. `common/harnessLog` was not ideally suited for general testing and made all the tests quite awkward. Instead, move all code used to test the `common/log` module into the `logTest` module and repurpose `common/harnessLog` to do log expect testing for all other tests in a cleaner way. Add a few exceptions for config testing since the log levels are reset by default in `config/parse`. 

- Add `--log-level-test` option. This allows setting the test log level independently from the general test harness setting, but current only works for the C tests. It is useful for seeing log output from functions on the console while a test is running. 

- Improve error reporting for `TEST_ASSIGN()` and `TEST_RESULT_VOID()` macros. 

- Update code count for new file types and exclusions. 

## v2.04 — Critical Bug Fix for Backup Resume

*Released: 2018-07-05*

### Core Bug Fixes

- Fix critical bug in resume that resulted in inconsistent backups. A regression in `v0.82` removed the timestamp comparison when deciding which files from the aborted backup to keep on resume. See note above for more details. 

- Fix error in selective restore when only one user database exists in the cluster. 

- Fix non-compliant ISO-8601 timestamp format in S3 authorization headers. AWS and some gateways were tolerant of space rather than zero-padded hours while others were not. 

### Core Features

- PostgreSQL 11 Beta 2 support. 

### Core Improvements

- Improve the HTTP client to set `content-length` to 0 when not specified by the server. S3 (and gateways) always set `content-length` or `transfer-encoding` but `HTTP 1.1` does not require it and proxies (e.g. HAProxy) may not include either. 

- Set `search_path = 'pg_catalog'` on PostgreSQL connections. 

### Core Development

- Move cryptographic hash functions to C using OpenSSL. 

- Split log levels into separate header file. Many modules that use `debug.h` do not need to do logging so this reduces dependencies for those modules. 

- Auto-generate Makefile with dependencies. 

- Rename `cipher` module to the more general `crypto`. 

- Update Debian package to add debug symbols to pgBackRest executable. 

- Convert the not very portable `uint` type to `unsigned int`. 

### Documentation Improvements

- Create a new section to describe building pgBackRest and build on a separate host. 

- Add sample S3 policy to restrict bucket privileges. 

### Documentation Development

- Fix default location of `pgbackrest.conf` in option reference. 

- Preliminary documentation for PostgreSQL 11 unprivileged user backup. 

- Remove call to `lscpu` which can vary widely by build host. 

- Build containers from scratch for more accurate testing. Use a prebuilt s3 server container. 

- Document generator improvements. Allow parameters to be passed when a container is created. Allow `/etc/hosts` update to be skipped (for containers without bash). Allow environment load to be skipped. Allow bash wrapping to be skipped. Allow forcing a command to run as a user without sudo. Allow an entire execute list to be hidden. 

### Test Development

- Add zero-length file to `mock`/`all` test. 

- Update primary test environment (Vagrant and Docker) to Ubuntu 18.04. 

- Improve efficiency of C library builds now that they are used only for testing. 

- Remove RHEL and Debian package patches since they have been committed upstream. 

- Update parameters for `VBoxService` start. 

- Make `ls` ordering deterministic in `mock`/`all` test. 

## v2.03 — Single Executable to Deploy

*Released: 2018-05-22*

### Core Bug Fixes

- Fix potential buffer overrun in error message handling. 

- Fix archive write lock being taken for the synchronous `archive-get` command. 

### Core Improvements

- Embed exported C functions and Perl modules directly into the pgBackRest executable. 

- Use `time_t` instead of `__time_t` for better portability. 

- Print total runtime in milliseconds at command end. 

### Core Development

- Add stack trace macros to all functions. Low-level functions only include stack trace in test builds while higher-level functions ship with stack trace built-in. Stack traces include all parameters passed to the function but production builds only create the parameter list when the log level is set high enough, i.e. `debug` or `trace` depending on the function. 

- Build `libc` using links rather than referencing the C files in `src` directly. The C library builds with different options which should not be reused for the C binary or vice versa. 

### Test Development

- Test harness improvements. Allow more than one test to provide coverage for the same module. Add option to disable valgrind. Add option to disabled coverage. Add option to disable debug build. Add option to disable compiler optimization. Add `--dev-test` mode. 

- Update SSL error message test on CentOS 7. 

- Set `log-timestamp=n` for integration tests. This means less filtering of logs needs to be done and new timestamps can be added without adding new filters. 

## v2.02 — Parallel Asynchronous Archive Get and Configuration Includes

*Released: 2018-05-06*

### Core Bug Fixes

- Fix directory syncs running recursively when only the specified directory should be synced. 

- Fix `archive-copy` throwing "path not found" error for incr/diff backups. 

- Fix failure in manifest build when two or more files in `PGDATA` are linked to the same directory. 

- Fix delta restore failing when a linked file is missing. 

- Fix rendering of key/value and list options in help. 

### Core Features

- Add asynchronous, parallel `archive-get`. This feature maintains a queue of WAL segments to help reduce latency when PostgreSQL requests a WAL segment with `restore_command`. 

- Add support for additional pgBackRest configuration files. The directory is specified by the `--config-include-path` option. Add `--config-path` option for overriding the default base path of the `--config` and `--config-include-path` option. 

- Add `repo-s3-token` option to allow temporary credentials tokens to be configured. pgBackRest currently has no way to request new credentials so the entire command (e.g. `backup`, `restore`) must complete before the credentials expire. 

### Core Improvements

- Update the `archive-push-queue-max`, `manifest-save-threshold`, and `buffer-size` options to accept values in `KB`, `MB`, `GB`, `TB`, or `PB` where the multiplier is a power of `1024`. 

- Make backup/restore path sync more efficient. Scanning the entire directory can be very expensive if there are a lot of small tables. The backup manifest contains the path list so use it to perform syncs instead of scanning the backup/restore path. 

- Show command parameters as well as command options in initial info log message. 

- Rename archive-queue-max option to archive-push-queue-max. This is consistent with the new `archive-get-queue-max` option. The old option name will continue to be accepted. 

### Core Development

- Make `backup.history` sync more efficient. Only the `backup.history/[year]` directory was being synced, so check if the `backup.history` is newly created and sync it as well. 

- Move async forking and more error handling to C. The Perl process was exiting directly when called but that interfered with proper locking for the forked async process. Now Perl returns results to the C process which handles all errors, including signals. 

- Improved lock implementation written in C. Now only two types of locks can be taken: `archive` and `backup`. Most commands use one or the other but the `stanza-*` commands acquire both locks. This provides better protection than the old command-based locking scheme. 

- Storage object improvements. Convert all functions to variadic functions. Enforce read-only storage. Add `storageLocalWrite()` helper function. Add `storageCopy()`, `storageExists()`, `storageMove()`, `storageNewRead()`/`storageNewWrite()`, `storagePathCreate()`, `storagePathRemove()`, `storagePathSync()`, and `storageRemove()`. Add `StorageFileRead` and `StorageFileWrite` objects. Abstract Posix driver code into a separate module. Call `storagePathRemove()` from the Perl Posix driver. 

- Improve `String` and `StringList` objects. Add `strUpper()`, `strLower()`, `strLstExists()`, `strLstExistsZ()`, `strChr()`, `strSub()`, `strSubN()`, and `strTrunc()`. 

- Improve `Buffer` object. Add `bufNewC()`, `bufEq()` and `bufCat()`. Only reallocate buffer when the size has changed. 

- Add `pgControlInfo()` to read `pg_control` and determine the PostgreSQL version. 

- Add `walSegmentNext()` and `walSegmentRange()`. 

- Error handling improvements. Add `THROWP_`* macro variants for error handling. These macros allow an `ErrorType` pointer to be passed and are required for functions that may return different errors based on a parameter. Add `_FMT` variants for all `THROW` macros so format types are checked by the compiler. 

- Split `cfgLoad()` into multiple functions to make testing easier. Mainly this helps with unit tests that need to do log expect testing. 

- Allow `MemContext` objects to be copied to a new parent. This makes it easier to create objects and then copy them to another context when they are complete without having to worry about freeing them on error. Update `List`, `StringList`, and `Buffer` to allow moves. Update `Ini` and `Storage` to take advantage of moves. 

- Full branch coverage in C code. 

- Refactor `usec` to `msec` in `common/time.c`. The implementation provides `usec` resolution but this is not needed in practice and it makes the interface more complicated due to the extra zeros. 

- Replace `THROW_ON_SYS_ERROR()` with `THROW_SYS_ERROR()`. The former macro was hiding missing branch coverage for critical error handling. 

- Start work on C handle io object and use it to output help. 

- Don't copy `CFGDEF_NAME_ALT` or `CFGDEF_INHERIT` when processing config option inheritance. 

- Split debug and assert code into separate headers. Assert can be used earlier because it only depends on the error-handler and not logging. Add `ASSERT()` macro which is preserved in production builds. 

- Cleanup C types. Remove `typec.h`. Order all typdefs above local includes. 

- Fix header exclusion defines that do not match the general pattern. 

### Documentation Bug Fixes

- Update docs with 32-bit support and caveats. 32-bit support was added in v1.26. 

### Documentation Improvements

- Add monitoring examples using PostgreSQL and jq. 

- Add example of command section usage to archiving configuration. 

- Remove documentation describing `info --output=json` as experimental. 

- Update out-of-date description for the `spool-path` option. 

### Documentation Development

- Add logic to find the real oid of the `test1` database during restore testing. 

- Document build improvements. Perform `apt-get update` to ensure packages are up to date before installing. Add `-p` to the repository `mkdir` so it won't fail if the directory already exists, handy for testing packages. 

### Test Features

- Use lcov for C unit test coverage reporting. Switch from Devel::Cover because it would not report on branch coverage for reports converted from gcov. Incomplete branch coverage for a module now generates an error. Coverage of unit tests is not displayed in the report unless they are incomplete for either statement or branch coverage. 

### Test Development

- Move test definitions to `test/define.yaml`. The location is better because it is no longer buried in the Perl test libs. Also, the data can be easily accessed from C. 

- Move help/version integration tests to `mock/all`. Help and version are covered by unit tests, so we really just to need to make sure there is output when called from the command line. 

- Move `archive-stop` and `expire` tests to the `mock` module. These are mock integration tests so they should be grouped with the other mock integration tests. 

- Add `harnessCfgLoad()` test function, which allows a new config to be loaded for unit testing without resetting log functions, opening a log file, or taking locks. 

- Add `HARNESS_FORK` macros for tests that require fork(). A standard pattern for tests makes fork() easier to use and should help prevent some common mistakes. 

- Add `TEST_ERROR_FMT` macro to simplify testing of formatted error messages. 

- Generate code counts for all source files. The source files are also classified by type and purpose. 

- Include VM type in `gcov` path to avoid conflicts between VMs with different architectures. 

- Improve logic for smart builds to include version changes. Skip version checks when testing in `--dev` mode. 

- Use pip 9.03 in test VMs. pip 10 drops support for Python 2.6 which is still used by the older test VMs. 

- Allow `-DDEBUG_UNIT` to be suppressed to test how debug macros behave. 

- Rename Perl tests so they don't conflict with their C counterparts. 

- Divide tests into three types (`unit`, `integration`, `performance`). Many options that were set per test can instead be inferred from the types, i.e. `container`, `c`, `expect`, and `individual`. 

- Try tweaking time sync settings to prevent clock drift rather than restarting `VBoxService` on every test run. 

## v2.01 — Minor Bug Fixes and Improvements

*Released: 2018-03-19*

### Core Bug Fixes

- Fix `--target-action` and `--recovery-option` options being reported as invalid when restoring with `--type=immediate`. 

- Immediately error when a secure option (e.g. `repo1-s3-key`) is passed on the command line. Since pgBackRest would not pass secure options on to sub-processes an obscure error was thrown. The new error is much clearer and provides hints about how to fix the problem. Update command documentation to omit secure options that cannot be specified on the command-line. 

- Fix issue passing `--no-config` to embedded Perl. 

- Fix issue where specifying `log-level-stderr` > `warn` would cause a `local`/`remote` process to error on exit due to output found on stderr when none was expected. The max value for a `local`/`remote` process is now `error` since there is no reason for these processes to emit warnings. 

- Fix manifest test in the `check` command when tablespaces are present. 

### Core Improvements

- Error when multiple arguments are set in the config file for an option that does not accept multiple arguments. 

- Remove extraneous sudo commands from `src/Makefile`. 

### Core Development

- Improve Perl configuration. Set config before `Main::main()` call to avoid secrets being exposed in a stack trace. Move logic for setting defaults to C. 

- Improve logging. Move command begin to C except when it must be called after another command in Perl (e.g. `expire` after `backup`). Command begin logs correctly for complex data types like hash and list. Specify which commands will log to file immediately and set the default log level for log messages that are common to all commands. File logging is initiated from C. 

- Port most of `Config::Config::configLoad()` from Perl to C. 

- Fix incorrect enum types in `config.c` that throw warnings under clang. 

- Enable `-Wswitch-enum`, `-Wconversion`, `-Wformat=2`, `-Wformat-nonliteral`, and `-Wformat-signedness` and silence new warnings. 

- Improve code documentation in `config` module. 

- Improve debugging. Add `ASSERT_DEBUG()` macro for debugging and replace all current `assert()` calls except in tests that can't use the debug code. Replace remaining NDEBUG blocks with the more granular DEBUG_UNIT. Remove some debug `memset()` calls in `MemContext` since valgrind is more useful for these checks. 

- Add `cfgOptionTest()` and update `cfgOption()` calls that are better implemented as `cfgOptionTest()`. 

- Build with `-DNDEBUG` by default but disable for testing. 

- Check `int` size in `common/type.h`. This ensures that integers are at least 32-bits without having to run the test suite. 

- Improve conversion of C exceptions to `Exception` objects. Colons in the message would prevent all of the message from being loaded into the `Exception` object. 

### Documentation Improvements

- Show index in examples for indexed options, i.e. `repo-*`, `pg-*`. 

- Simplify table of contents on command page by only listing commands. 

- Remove references to the C library being optional. 

### Test Features

- Add CentOS/RHEL package builds. 

- Use clang for static code analysis. Nothing found initially except for some functions that should have been marked `__noreturn__`. 

### Test Development

- Build performance improvements. Improve bin and libc build performance. Improve code generation performance. 

- Config test code writes secure options to a file instead of passing on the command-line. 

- Disable console display of coverage for C files since `Devel::Cover` does not handle it well. 

- Add new test for `Common::Io::Process` to show that output on stderr will raise an exception on `close()` even if the exit code is 0. 

- Update `pip` before installing `awscli`. 

- Remove `--smart` from `--expect` tests. This ensures that new binaries are built before running the tests. 

- Remove Debian package patch now that it has been merged upstream. 

## v2.00 — Performance Improvements for Archive Push

*Released: 2018-02-23*

### Core Features

- The `archive-push` command is now partially coded in C which allows the PostgreSQL `archive_command` to run significantly faster when processing status messages from the asynchronous archive process. 

### Core Improvements

- Improve `check` command to verify that the backup manifest can be built. 

- Improve performance of HTTPS client. Buffering now takes the `pending` bytes on the socket into account (when present) rather than relying entirely on `select()`. In some instances the final bytes would not be flushed until the connection was closed. 

- Improve S3 delete performance. The constant `S3_BATCH_MAX` had been replaced with a hard-coded value of 2, probably during testing. 

- Allow any non-command-line option to be reset to default on the command-line. This allows options in `pgbackrest.conf` to be reset to default which reduces the need to write new configuration files for specific needs. 

- The C library is now required. This eliminates conditional loading and eases development of new library features. 

- The `pgbackrest` executable is now a C binary instead of Perl. This allows certain time-critical commands (like async `archive-push`) to run more quickly. 

- Rename `db-*` options to `pg-*` and `backup-*` options to `repo-*` to improve consistency. `repo-*` options are now indexed although currently only one is allowed. 

### Core Development

- Implement `help` command in C. 

- Implement `version` command in C. 

- Config parsing implemented in C and passed to Perl as JSON. 

- Add `Buffer`, `Ini`, `KeyValue`, `List`, `RegExp`, `Storage`, `String`, `StringList`, `Variant`, `VariantList`, and `Wait` objects. 

- Add `command`, `exit`, `log`, and `time` modules. 

- Remove deprecated `archive-max-mb` option. 

- Improve `MemContext` module. Add temporary context blocks and refactor allocation arrays to include allocation size. 

- Improve `error` module. Add functions to convert error codes to C errors and handle system errors. 

- Create a master list of errors in `build/error.yaml`. The C and Perl errors lists are created automatically by `Build.pm` so they stay up to date. 

- Move lock release later in exitSafe() to reduce the chance of a new process starting and acquiring a lock before the old process has exited. 

- Add 30 second wait loop to lockAcquire() when fail on no lock enabled. This should help prevent processes that are shutting down from interfering with processes that are starting up. 

- Replace `cfgCommandTotal()`/`cfgOptionTotal()` functions with constants. The constants are applicable in more cases and allow the compiler to optimize certain loops more efficiently. 

- Cleanup usage of internal options. Apply internal to options that need to be read to determine locality but should not appear in the help. 

- Refactor code to make valgrind happy. 

- Fix non-compliant formatting for function declarations. 

### Documentation Features

- All clusters in the documentation are initialized with checksums. 

### Documentation Improvements

- List deprecated option names in documentation and command-line help. 

- Clarify that S3 buckets must be created by the user. 

### Documentation Development

- Add coding standards document. 

- Improve section source feature to not require a title or content. The title will be pulled from the source document. 

- Allow code blocks to have a type. Currently this is only rendered in Markdown. 

- Add table render for Markdown format. 

- PDF rendering improvements. Check both `doc-path` and `bin-path` for logo. Allow PDF to be output to a location other than the `output` directory. Use PDF-specific version variable for more flexible formatting. Allow sections to be excluded from table of contents. More flexible replacements for titles and footers. Fill is now the default for table columns. Column width is specified as a percentage rather that using latex-specific notation. Fix missing variable replace for `code-block` title. 

- Add `id` param for hosts created with `host-add`. The `host-*-ip` variable is created from the `id` param so the `name` param can be changed without affecting the `host-*-ip` variable. If `id` is not specified then it is copied from `name`. 

- Deploy historical documentation to `prior` rather than the root directory. 

### Test Development

- Run valgrind on all C unit tests. 

- Only build C binary/library for Perl unit/integration tests or C unit tests that require Perl. 

- Improve speed of C unit tests. Preserve object files between tests and use a Makefile to avoid rebuilding object files. 

- Report coverage errors via the console. This helps with debugging coverage issues on remote services like Travis. 

- No longer run `master` branch through CI. The `integration` branch will be run through CI and then pushed to `master` with github status checks. 

- Rename Perl tests so they don't conflict with their C counterparts. 

- Update URL for Debian package repository. 

## v1.29 — Critical Bug Fix for Backup Resume

*Released: 2018-07-05*

### Core Bug Fixes

- Fix critical bug in resume that resulted in inconsistent backups. A regression in `v0.82` removed the timestamp comparison when deciding which files from the aborted backup to keep on resume. See note above for more details. 

- Fix non-compliant ISO-8601 timestamp format in S3 authorization headers. AWS and some gateways were tolerant of space rather than zero-padded hours while others were not. 

- Fix directory syncs running recursively when only the specified directory should be synced. 

- Fix `--target-action` and `--recovery-option` options being reported as invalid when restoring with `--type=immediate`. 

- Fix `archive-copy` throwing "path not found" error for incr/diff backups. 

- Fix failure in manifest build when two or more files in `PGDATA` are linked to the same directory. 

- Fix delta restore failing when a linked file was missing. 

- Fix error in selective restore when only one user database exists in the cluster. 

### Core Improvements

- Improve the HTTP client to set `content-length` to 0 when not specified by the server. S3 (and gateways) always set `content-length` or `transfer-encoding` but `HTTP 1.1` does not require it and proxies (e.g. HAProxy) may not include either. 

- Improve performance of HTTPS client. Buffering now takes the `pending` bytes on the socket into account (when present) rather than relying entirely on `select()`. In some instances the final bytes would not be flushed until the connection was closed. 

- Improve S3 delete performance. The constant `S3_BATCH_MAX` had been replaced with a hard-coded value of 2, probably during testing. 

- Make backup/restore path sync more efficient. Scanning the entire directory can be very expensive if there are a lot of small tables. The backup manifest contains the path list so use it to perform syncs instead of scanning the backup/restore path. Remove recursive path sync functionality since it is no longer used. 

### Core Development

- Make `backup.history` sync more efficient. Only the `backup.history/[year]` directory was being synced, so check if the `backup.history` is newly created and sync it as well. 

- Add log-level-stderr option for stanza-* commands. 

### Documentation Bug Fixes

- Update docs with 32-bit support and caveats. 32-bit support was added in v1.26. 

### Documentation Improvements

- Clarify that S3 buckets must be created by the user. 

- Update out-of-date description for the `spool-path` option. 

### Documentation Development

- Remove call to `lscpu` which can vary widely by build host. 

### Test Development

- Add new test for `Common::Io::Process` to show that output on stderr will raise an exception on `close()` even if the exit code is 0. 

- Add zero-length file to `mock`/`all` test. 

- Disable package build tests since `v1` will no longer be packaged. Users installing packages should update to `v2`. `v1` builds are intended for users installing from source. 

- Update SSL error message test on CentOS 7. 

- Update URL for Debian package repository. 

- Make `ls` ordering deterministic in `mock`/`all` test. 

- Change backup test user from `backrest` to `pgbackrest`. 

## v1.28 — Stanza Delete

*Released: 2018-02-01*

### Core Bug Fixes

- Fixed inability to restore a single database contained in a tablespace using --db-include. 

- Ensure latest `db-id` is selected on when matching `archive.info` to `backup.info`. This provides correct matching in the event there are `system-id` and `db-version` duplicates (e.g. after reverting a `pg_upgrade`). 

- Fixed overly chatty error message when reporting an invalid command. 

### Core Features

- Add `stanza-delete` command to cleanup unused stanzas. 

### Core Improvements

- Improve `stanza-create` command so that it does not error when the stanza already exists. 

### Core Development

- Minor changes to `Manifest` module, mostly for test reproducibility. 

- Fix non-compliant formatting for function declarations. 

### Documentation Improvements

- Update `stanza-create --force` documentation to urge caution when using. 

### Test Development

- Add unit tests for the `Manifest` module. 

## v1.27 — Bug Fixes and Documentation

*Released: 2017-12-19*

### Core Bug Fixes

- Fixed an issue that suppressed locality errors for `backup` and `restore`. When a backup host is present, backups should only be allowed on the backup host and restores should only be allowed on the database host unless an alternate configuration is created that ignores the remote host. 

- Fixed an issue where WAL was not expired on PostgreSQL 10. This was caused by a faulty regex that expected all PostgreSQL major versions to be X.X. 

- Fixed an issue where the `--no-config` option was not passed to child processes. This meant the child processes would still read the local config file and possibly cause unexpected behaviors. 

- Fixed `info` command to eliminate `"db (prior)"` output if no backups or archives exist for a prior version of the cluster. 

### Core Development

- Add `memGrowRaw()` to memory context module. 

### Documentation Features

- Document the relationship between the `archive-copy` and `archive-check` options. 

- Improve `archive-copy` reference documentation. 

### Documentation Development

- Relax permissions set by `release.pl`. 

- Split "refactor" sections into "improvements" and "development" in the release notes. Many development notes are not relevant to users and simply clutter the release notes, so they are no longer shown on the website. 

- Allow internal options that do not show up in the documentation. Used for test options initially but other use cases are on the horizon. 

### Test Development

- Update CI branches to `release/1` and `release/1-integration`. 

- No longer run `release/1` branch through CI. The `release/1-integration` branch will be run through CI and then pushed to `release/1` with github status checks. 

- Move restore test infrastructure to `HostBackup.pm`. Required to test restores on the backup server, a fairly common scenario. Improve the restore function to accept optional parameters rather than a long list of parameters. In passing, clean up extraneous use of `strType` and `strComment` variables. 

- Sync time to prevent build failures when running on VirtualBox. 

## v1.26 — Repository Encryption

*Released: 2017-11-21*

### Core Bug Fixes

- Fixed an issue that could cause copying large manifests to fail during restore. 

- Fixed incorrect WAL offset for 32-bit architectures. 

- Fixed an issue retrieving WAL for old database versions. After a `stanza-upgrade` it should still be possible to restore backups from the previous version and perform recovery with `archive-get`. However, archive-get only checked the most recent db version/id and failed. Also clean up some issues when the same db version/id appears multiple times in the history. 

- Fixed an issue with invalid backup groups being set correctly on restore. If the backup cannot map a group to a name it stores the group in the manifest as `false` then uses either the owner of $PGDATA to set the group during restore or failing that the group of the current user. This logic was not working correctly because the selected group was overwriting the user on restore leaving the group undefined and the user incorrectly set to the group. 

- Fixed an issue passing parameters to remotes. When more than one db was specified the path, port, and socket path would for db1 were passed no matter which db was actually being addressed. 

### Core Features

- Repository encryption support. 

### Core Improvements

- Disable gzip filter when `--compress-level-network=0`. The filter was used with compress level set to 0 which added overhead without any benefit. 

- Inflate performance improvement for gzip filter. 

### Core Development

- Refactor protocol param generation into a new function. This allows the code to be tested more precisely and doesn't require executing a remote process. 

- Add `list` type for options. The `hash` type was being used for lists with an additional flag (`value-hash`) to indicate that it was not really a hash. 

- Remove configurable option hints. `db-path` was the only option with a hint so the feature seemed wasteful. All missing stanza options now output the same hint without needing configuration. 

- Convert configuration definitions from auto-generated functions to auto-generated data structures. 

- Add `eof` to S3 file driver (required for encryption support). 

- Enable additional warnings for C builds. 

- Simplify try..catch..finally names. Also wrap in a do...while loop to make sure that no random else is attached to the main if block. 

- Improve base64 implementation. Different encoded strings could be generated based on compiler optimizations. Even though decoding was still successful the encoded strings did not match the standard. 

- Disable `-Wclobber` compiler warning because it is mostly useless but keep the rest of `-Wextra`. 

### Documentation Features

- Add template to improve initial information gathered for issue submissions. 

### Documentation Improvements

- Clarify usage of the `archive-timeout` option and describe how it is distinct from the PostgreSQL `archive_timeout` setting. 

### Documentation Development

- Update `release.pl` to push data to site repository. 

### Test Features

- Automated tests for 32-bit i386/i686 architecture. 

### Test Development

- Update Debian/Ubuntu containers to download latest version of `pip`. 

- Full unit test coverage for gzip filter. 

- Only check expect logs on CentOS 7. Variations in distros cause false negatives in tests but don't add much value. 

- Fix flapping protocol timeout test. It only matters that the correct error code is returned, so disable logging to prevent message ordering from failing the expect test. 

- Designate a single distro (Ubuntu 16.04) for coverage testing. Running coverage testing on multiple distros takes time but doesn't add significant value. Also ensure that the distro designated to run coverage tests is one of the default test distros. For C tests, enable optimizations on the distros that don't do coverage testing. 

- Automate generation of WAL and `pg_control` test files. The existing static files would not work with 32-bit or big-endian systems so create functions to generate these files dynamically rather than creating a bunch of new static files. 

- Refactor C unit test macros so they compile with `-Wstrict-aliasing`. 

- Refactor C page checksum unit test to compile with `-Wstrict-aliasing`. 

## v1.25 — S3 Performance Improvements

*Released: 2017-10-24*

### Core Bug Fixes

- Fix custom settings for `compress-level` option being ignored. 

- Remove error when overlapping timelines are detected. Overlapping timelines are valid in many Point-in-Time-Recovery (PITR) scenarios. 

- Fix instances where `database-id` was not rendered as an integer in JSON info output. 

### Core Features

- Improve performance of list requests on S3. Any beginning literal portion of a filter expression is used to generate a search prefix which often helps keep the request small enough to avoid rate limiting. 

### Core Development

- Improve protocol error handling. In particular, "stop" errors are no longer reported as "unexpected". 

- Allow functions with sensitive options to be logged at debug level with redactions. Previously, functions with sensitive options had to be logged at trace level to avoid exposing them. Trace level logging may still expose secrets so use with caution. 

- Replace dynamically built class hierarchies in I/O layer with fixed `parent()` calls. 

- Improve labeling for errors in helper processes. 

- Update C naming conventions. 

- Use `int` datatype wherever possible. 

- Better separation of C source from Perl interface. 

- Add `LibC.template.pm` to simplify LibC module generation. 

- Add C error handler. 

- Perl error handler recognizes errors thrown from the C library. 

- Page checksum module uses new C error handler. 

- Add C memory contexts. 

- Add base64 encode/decode. 

### Test Features

- Add I/O performance tests. 

### Test Development

- Add C unit test infrastructure. 

- Add test macros for C results and errors. 

- Warnings in C builds treated as errors. 

- Run all tests on tempfs rather than local disk. 

- Improve performance of test code. Wait when all tests have been assigned to reduce CPU load. 

- Remove Debian test repo after PostgreSQL 10 release. 

- Convert config and page checksum tests into C unit tests. 

- Add PostgreSQL versions to Debian VMs for testing. 

## v1.24 — New Backup Exclusions

*Released: 2017-09-28*

### Core Bug Fixes

- Fixed an issue where warnings were being emitted in place of lower priority log messages during backup from standby initialization. 

- Fixed an issue where some `db-*` options (e.g. `db-port`) were not being passed to remotes. 

### Core Features

- Exclude contents of `pg_snapshots`, `pg_serial`, `pg_notify`, and `pg_dynshmem` from backup since they are rebuilt on startup. 

- Exclude `pg_internal.init` files from backup since they are rebuilt on startup. 

### Core Improvements

- Open log file after async process is completely separated from the main process to prevent the main process from also logging to the file. 

### Core Development

- Dynamically generate list of files for C library build. 

- Break up `LibC.xs` into separate module files. 

### Documentation Features

- Add passwordless SSH configuration. 

### Documentation Improvements

- Rename master to primary in documentation to align with PostgreSQL convention. 

### Documentation Development

- Add full installation where required and remove doc containers that included parts of the installation. 

### Test Development

- Improve C library smart build by ignoring changes outside of `/lib/pgBackRest/Config`. 

## v1.23 — Multiple Standbys and PostgreSQL 10 Support

*Released: 2017-09-03*

### Core Bug Fixes

- Fixed an issue that could cause compression to abort on growing files. 

- Fixed an issue with keep-alives not being sent to the remote from the local process. 

### Core Features

- Up to seven standbys can be configured for backup from standby. 

- PostgreSQL 10 support. 

- Allow `content-length` (in addition to chunked encoding) when reading XML data to improve compatibility with third-party S3 gateways. 

### Core Improvements

- Increase HTTP timeout for S3. 

- Add HTTP retries to harden against transient S3 network errors. 

### Core Development

- Configuration definitions are now pulled from the C library when present. 

### Documentation Bug Fixes

- Fixed document generation to include section summaries on the Configuration page. 

### Documentation Development

- Move contributor list to the end of `release.xml` for convenience. 

### Test Development

- Change log test order to ignore unimportant log errors while shutting down PostgreSQL. 

- Drain `stderr` during test process execution as well as termination to prevent lockups if there is a lot of output. 

- Update Docker build in `Vagrantfile`. 

- Update containers to support C library builds in the documentation. 

- Simplify smart logic for C Library and package builds. 

## v1.22 — Fixed S3 Retry

*Released: 2017-08-09*

### Core Bug Fixes

- Fixed authentication issue in S3 retry. 

## v1.21 — Improved Info Output and SSH Port Option

*Released: 2017-08-08*

### Core Bug Fixes

- The `archive_status` directory is now recreated on restore to support PostgreSQL 8.3 which does not recreate it automatically like more recent versions do. 

- Fixed an issue that could cause the empty archive directory for an old PostgreSQL version to be left behind after a `stanza-upgrade`. 

### Core Features

- Modified the `info` command (both text and JSON output) to display the archive ID and minimum/maximum WAL currently present in the archive for the current and prior, if any, database cluster version. 

- Added `--backup-ssh-port` and `--db-ssh-port` options to support non-default SSH ports. 

### Core Improvements

- Retry when S3 returns an internal error (500). 

### Core Development

- Add `bIgnoreMissing` parameter to `Local->manifest()`. 

### Documentation Bug Fixes

- Fix description of `--online` based on the command context. 

### Documentation Features

- Add creation of `/etc/pgbackrest.conf` to manual installation instructions. 

### Documentation Improvements

- Move repository options into a separate section in command/command-line help. 

### Documentation Development

- Reduce log verbosity when building documentation by only logging sections that contain an execute list directly or in a child section. 

- Debian/Ubuntu documentation now builds on Ubuntu 16. 

- Remove vestigial repository options from `backup` command. 

### Test Development

- Fix log checking after PostgreSQL shuts down to include `FATAL` messages and disallow immediate shutdowns which can throw `FATAL` errors in the log. 

- Use Google DNS in test environment for consistency. 

- Use new Travis Trusty image. 

- Generate global fake cert in containers for testing. 

- Consolidate `stanza-create` and `stanza-upgrade` tests into new `stanza` test. 

## v1.20 — Critical 8.3/8.4 Bug Fix

*Released: 2017-06-27*

### Core Bug Fixes

- Fixed an issue that prevented tablespaces from being backed up on PostgreSQL &le; 8.4. 

- Fixed missing flag in C library build that resulted in a mismatched binary on 32-bit systems. 

### Core Features

- Add `s3-repo-ca-path` and `s3-repo-ca-file` options to accommodate systems where CAs are not automatically found by `IO::Socket::SSL`, i.e. RHEL7, or to load custom CAs. 

### Core Development

- Harden protocol handshake to handle race conditions. 

- Fixed misleading error message when a file was opened for write in a missing directory. 

- Change log level of hardlink logging to `detail`. 

- Cast size in S3 manifest to integer. 

- Rename `Archive` modules to remove redundancy. 

- Improve S3 error reporting. 

- Minor optimizations to package loads and ordering for `archive-get` and `archive-push` commands. 

### Documentation Development

- Remove exhaustive version list from Stable Releases TOC. 

- Improve S3 server implementation in documentation. 

- Update CentOS 6 documentation to build on PostgreSQL 9.5. 

- Remove `mount` from host `cache-key` because it can vary by system. 

### Test Features

- Add documentation builds to CI. 

### Test Development

- Fix timeouts in `ExecuteTest` to speed multi-process testing. 

- Remove patch directory before Debian package builds. 

- Combine hardlink and non/compressed in synthetic tests to reduce test time and improve coverage. 

- Split `full` module into `mock` and `real` to allow better test combinations and save time in CI. 

- Consolidate `archive-push` and `archive-get` tests into new `archive` test. 

- Eliminate redundancy in `real` tests. 

- Install `sudo` in base containers rather than on demand. 

- More optimized container suite that greatly improves build time. 

- Added static Debian packages for `Devel::Cover` to reduce build time. 

- Add `deprecated` state for containers. Deprecated containers may only be used to build packages. 

- Remove Debian 8 from CI because it does not provide additional coverage over Ubuntu 12.04, 14.04, 16.04. 

- Add Debian 9 to test suite. 

- Remove `process-max` option. Parallelism is now tested in a more targeted manner and the high level option is no longer needed. 

- Balance database versions between VMs to minimize test duration. 

- Automatically check that all supported PostgreSQL versions are being tested on a single default VM. 

- Add `performance` module and basic performance test for `archive-push`. 

## v1.19 — S3 Support

*Released: 2017-06-12*

### Core Bug Fixes

- Fixed the `info` command so the WAL archive min/max displayed is for the current database version. 

- Fixed the `backup` command so the `backup-standby` option is reset (and the backup proceeds on the primary) if the standby is not configured and/or reachable. 

- Fixed config warnings raised from a remote process causing errors in the master process. 

### Core Features

- Amazon S3 repository support. 

### Core Development

- Refactor storage layer to allow for new repository filesystems using drivers. 

- Refactor IO layer to allow for new compression formats, checksum types, and other capabilities using filters. 

- Move modules in `Protocol` directory in subdirectories. 

- Move backup modules into `Backup` directory. 

### Documentation Bug Fixes

- Changed invalid `max-archive-mb` option in configuration reference to `archive-queue-max`. 

- Fixed missing `sudo` in installation section. 

### Test Development

- Fixed an undefined variable when a module had no uncoverable code exceptions. 

- Fixed issue with `--dry-run` requiring `--vm-out` to work properly. 

- Moved test and env modules to new directories to avoid namespace conflicts with common tests. 

- Set `--vm-max=2` for CI. 

- Remove flapping protocol timeout test that will be replaced in the upcoming storage patch. 

## v1.18 — Stanza Upgrade, Refactoring, and Locking Improvements

*Released: 2017-04-12*

### Core Bug Fixes

- Fixed an issue where read-only operations that used local worker processes (i.e. `restore`) were creating write locks that could interfere with parallel `archive-push`. 

### Core Features

- Added the stanza-upgrade command to provide a mechanism for upgrading a stanza after upgrading to a new major version of PostgreSQL. 

- Added validation of `pgbackrest.conf` to display warnings if options are not valid or are not in the correct section. 

### Core Improvements

- Simplify locking scheme. Now, only the master process will hold write locks (for `archive-push` and `backup` commands) and not all local and remote worker processes as before. 

- Do not set timestamps of files in the backup directories to match timestamps in the cluster directory. This was originally done to enable backup resume, but that process is now implemented with checksums. 

- Improved error message when the `restore` command detects the presence of `postmaster.pid`. 

- Renumber return codes between 25 and 125 to avoid PostgreSQL interpreting some as fatal signal exceptions. 

### Core Development

- Refactor `Ini.pm` to facilitate testing. 

- The `backup` and `restore` commands no longer copy via temp files. In both cases the files are checksummed on resume so there's no danger of partial copies. 

- Allow functions to accept optional parameters as a hash. 

- Refactor `File->list()` and `fileList()` to accept optional parameters. 

- Refactor `backupLabel()` and add unit tests. 

- Silence some perl critic warnings. 

### Documentation Development

- Update wording for release note sections. 

- Ignore clock skew in container libc/package builds using make. It is common for containers to have clock skew so the build process takes care of this issue independently. 

### Test Development

- Complete statement/branch coverage for `Ini.pm`. 

- Improved functions used to test/munge manifest and info files. 

- Coverage testing always enabled on Debian-based containers. 

- Require description in every call to `testResult()`. 

- Make `iWaitSeconds` an optional parameter for `testResult()`. 

- Updated vagrant to new version and image. 

- Fixed flapping archive stop tests. 

- Added ability to test warning messages. 

## v1.17 — Page Checksum Bug Fix

*Released: 2017-03-13*

### Core Bug Fixes

- Fixed an issue where newly initialized (but unused) pages would cause page checksum warnings. 

## v1.16 — Page Checksum Improvements, CI, and Package Testing

*Released: 2017-03-02*

### Core Bug Fixes

- Fixed an issue where tables over 1GB would report page checksum warnings after the first segment. 

- Fixed an issue where databases created with a non-default tablespace would raise bogus warnings about `pg_filenode.map` and `pg_internal.init` not being page aligned. 

### Core Development

- Improved the code and tests for `fileManifest()` to prevent a possible race condition when files are removed by the database while the manifest is being built. 

### Documentation Development

- Container executions now load the user's environment. 

### Test Features

- Continuous integration using `travis-ci`. 

- Automated builds of Debian packages for all supported distributions. 

### Test Development

- Added `--dev` option to aggregate commonly used dev options. 

- Added `--retry` option. 

- Added `--no-package` option to skip package builds. 

- C library and packages are built by default, added `-smart` option to rebuild only when file changes are detected. 

- The `--libc-only` option has been changed to `--build-only` now that packages builds have been added. 

- Improved formatting of `testResult()` output. 

- Improved truncation when outputting errors logs in the `ExecuteTest` module. 

- Fixed flapping archive-stop test with `testResult()` retries. 

- Added final test of archive contents to archive-push test. 

- Temporarily disable flapping keep-alive test. 

## v1.15 — Refactoring and Bug Fixes

*Released: 2017-02-13*

### Core Bug Fixes

- Fixed a regression introduced in v1.13 that could cause backups to fail if files were removed (e.g. tables dropped) while the manifest was being built. 

### Core Development

- Refactor `FileCommon::fileManifest()` and `FileCommon::fileStat` to be more modular to allow complete branch/statement level coverage testing. 

### Test Development

- Complete branch/statement level coverage testing for `FileCommon::fileManifest()` and `FileCommon::fileStat` functions and helper functions. 

## v1.14 — Refactoring and Bug Fixes

*Released: 2017-02-13*

### Core Bug Fixes

- Fixed an issue where an archive-push error would not be retried and would instead return errors to PostgreSQL indefinitely (unless the `.error` file was manually deleted). 

- Fixed a race condition in parallel archiving where creation of new paths generated an error when multiple processes attempted to do so at the same time. 

### Core Improvements

- Improved performance of `wal archive min/max` provided by the `info` command. 

### Documentation Features

- Updated async archiving documentation to more accurately describe how the new method works and how it differs from the old method. 

### Documentation Development

- Documentation can now be built with reusable blocks to reduce duplication. 

- Improved support for `--require` option and section depends now default to the previous section. 

- Added ability to pass options to containers within the documentation. 

- Add `proper` tag to slightly emphasize proper nouns. 

## v1.13 — Parallel Archiving, Stanza Create, Improved Info and Check

*Released: 2017-02-05*

### Core Bug Fixes

- Fixed const assignment giving compiler warning in C library. 

- Fixed a few directory syncs that were missed for the `--repo-sync` option. 

- Fixed an issue where a missing user/group on restore could cause an "uninitialized value" error in `File->owner()`. 

- Fixed an issue where protocol mismatch errors did not output the expected value. 

- Fixed a spurious `archive-get` log message that indicated an exit code of 1 was an abnormal termination. 

### Core Features

- Improved, multi-process implementation of asynchronous archiving. 

- Improved `stanza-create` command so that it can repair broken repositories in most cases and is robust enough to be made mandatory. 

- Improved `check` command to run on a standby, though only basic checks are done because `pg_switch_xlog()` cannot be executed on a replica. 

- Added archive and backup WAL ranges to the `info` command. 

- Added warning to update `pg_tablespace.spclocation` when remapping tablespaces in PostgreSQL < 9.2. 

- Remove remote lock requirements for the `archive-get`, `restore`, `info`, and `check` commands since they are read-only operations. 

### Core Improvements

- Log file banner is not output until the first log entry is written. 

- Reduced the likelihood of torn pages causing a false positive in page checksums by filtering on start backup LSN. 

- Remove Intel-specific optimization from C library build flags. 

- Remove `--lock` option. This option was introduced before the lock directory could be located outside the repository and is now obsolete. 

- Added `--log-timestamp` option to allow timestamps to be suppressed in logging. This is primarily used to avoid filters in the automated documentation. 

- Return proper error code when unable to convert a relative path to an absolute path. 

### Core Development

- Refactor `File` and `BackupCommon` modules to improve test coverage. 

- Moved `File->manifest()` into the `FileCommon.pm` module. 

- Moved the `Archive` modules to the `Archive` directory and split the `archive-get` and `archive-push` commands into separate modules. 

- Split the `check` command out of the `Archive.pm` module. 

- Allow logging to be suppressed via `logDisable()` and `logEnable()`. 

- Allow for locks to be taken more than once in the same process without error. 

- Lock directories can be created when more than one directory level is required. 

- Clean up `optionValid()`/`optionTest()` logic in `Lock.pm`. 

- Added `Exception::exceptionCode()` and `Exception::exceptionMessage()` to simplify error handling logic. 

- Represent `.gz` extension with a constant. 

- Allow empty files to be created with `FileCommon::fileStringWrite()` and use temp files to avoid partial reads. 

- Refactor process IO and process master/minion code out from the common protocol code. 

- Fixed alignment issues with multiline logging. 

### Documentation Features

- Added documentation to the User Guide for the `process-max` option. 

### Documentation Development

- Update LICENSE.txt for 2017. 

### Test Development

- Fixed `--no-online` tests to suppress expected errors. 

- Added integration for testing coverage with `Devel::Cover`. 

- Added unit tests for low-level functions in the `File` and `BackupCommon` modules. 

- C Library builds only run when C library has actually changed. 

- Added more flexibility in initializing and cleaning up after modules and tests. 

- `testResult()` suppresses logging and reports exceptions. 

- `testException()` allows messages to be matched with regular expressions. 

- Split test modules into separate files to make the code more maintainable. Tests are dynamically loaded by name rather than requiring an if-else block. 

- Allow multiple `--module`, `--test`, and `--run` options to be used for `test.pl`. 

- Added expect log expression to replace year subdirectories in `backup.history`. 

- Refactor name/locations of common modules that setup test environments. 

## v1.12 — Page Checksums, Configuration, and Bug Fixes

*Released: 2016-12-12*

### Core Bug Fixes

- Fixed an issue where options that were invalid for the specified command could be provided on the command-line without generating an error. The options were ignored and did not cause any change in behavior, but it did lead to some confusion. Invalid options will now generate an error. 

- Fixed an issue where internal symlinks were not being created for tablespaces in the repository. This issue was only apparent when trying to bring up clusters in-place manually using filesystem snapshots and did not affect normal backup and restore. 

- Fixed an issue that prevented errors from being output to the console before the logging system was initialized, i.e. while parsing options. Error codes were still being returned accurately so this would not have made a process look like it succeeded when it did not. 

- Fixed an issue where the `db-port` option specified on the backup server would not be properly passed to the remote unless it was from the first configured database. 

### Core Features

- Added the `--checksum-page` option to allow pgBackRest to validate page checksums in data files when checksums are enabled on PostgreSQL >= 9.3. Note that this functionality requires a C library which may not initially be available in OS packages. The option will automatically be enabled when the library is present and checksums are enabled on the cluster. 

- Added the `--repo-link` option to allow internal symlinks to be suppressed when the repository is located on a filesystem that does not support symlinks. This does not affect any pgBackRest functionality, but the convenience link `latest` will not be created and neither will internal tablespace symlinks, which will affect the ability to bring up clusters in-place manually using filesystem snapshots. 

- Added the `--repo-sync` option to allow directory syncs in the repository to be disabled for file systems that do not support them, e.g. NTFS. 

- Added a predictable log entry to signal that a command has completed successfully. For example a backup ends successfully with: `INFO: backup command end: completed successfully`. 

### Core Improvements

- For simplicity, the `pg_control` file is now copied with the rest of the files instead of by itself of at the end of the process. The `backup` command does not require this behavior and the `restore` copies to a temporary file which is renamed at the end of the restore. 

### Core Development

- Abstracted code to determine which database cluster is the primary and which are standbys. 

- Improved consistency and flexibility of the protocol layer by using JSON for all messages. 

- File copy protocol now accepts a function that can do additional processing on the copy buffers and return a result to the calling process. 

- Improved `IO->bufferRead` to always return requested number of bytes until EOF. 

- Simplified the result hash of `File->manifest()`, `Db->tablespaceMapGet()`, and `Db->databaseMapGet()`. 

- Improved errors returned from child processes by removing redundant error level and code. 

- Code cleanup in preparation for improved `stanza-create` command. 

- Improved parameter/result logging in debug/trace functions. 

### Documentation Bug Fixes

- Fixed an issue that suppressed exceptions in PDF builds. 

- Fixed regression in section links introduced in v1.10. 

### Documentation Features

- Added Retention to QuickStart section. 

### Documentation Development

- Allow a source to be included as a section so large documents can be broken up. 

- Added section link support to Markdown output. 

- Added list support to PDF output. 

- Added `include` option to explicitly build sources (complements the `exclude` option though both cannot be used in the same invocation). 

- Added `keyword-add` option to add keywords without overriding the `default` keyword. 

- Added `debug` option to `doc.pl` to easily add the `debug` keyword to documentation builds. 

- Added `pre` option to `doc.pl` to easily add the `pre` keyword to documentation builds. 

- Builds in `release.pl` now remove all docker containers to get consistent IP address assignments. 

- Improvements to markdown rendering. 

- Remove code dependency on `project` variable, instead use `title` param. 

### Test Development

- Removed erroneous `--no-config` option in `help` test module. 

- Update control and WAL test files to `9.4` with matching system identifiers. 

- Improved exception handling in file unit tests. 

- Changed the `--no-fork` test option to `--fork` with negation to match all other boolean parameters. 

- Various improvements to validation of backup and restore. 

- Add more realistic data files to synthetic backup and restore tests. 

## v1.11 — Bug Fix for Asynchronous Archiving Efficiency

*Released: 2016-11-17*

### Core Bug Fixes

- Fixed an issue where asynchronous archiving was transferring one file per execution instead of transferring files in batches. This regression was introduced in v1.09 and affected efficiency only, all WAL segments were correctly archived in asynchronous mode. 

## v1.10 — Stanza Creation and Minor Bug Fixes

*Released: 2016-11-08*

### Core Bug Fixes

- Fixed an issue where a backup could error if no changes were made to a database between backups and only `pg_control` changed. 

- Fixed an issue where tablespace paths with the same prefix would cause an invalid link error. 

### Core Features

- Added the `stanza-create` command to formalize creation of stanzas in the repository. 

### Core Improvements

- Removed extraneous `use lib` directives from Perl modules. 

### Documentation Development

- Fixed missing variable replacements. 

- Removed hard-coded host names from configuration file paths. 

- Allow command-line length to be configured using `cmd-line-len` param. 

- Added `compact` param to allow CSS to be embedded in HTML file. 

- Added `pretty` param to produce HTML with proper indenting. 

- Only generate HTML menu when required and don't require index page. 

- Assign numbers to sections by default. 

- VM mount points are now optional. 

## v1.09 — 9.6 Support, Configurability, and Bug Fixes

*Released: 2016-10-10*

### Core Bug Fixes

- Fixed the `check` command to prevent an error message from being logged if the backup directory does not exist. 

- Fixed error message to properly display the archive command when an invalid archive command is detected. 

- Fixed an issue where the async archiver would not be started if `archive-push` did not have enough space to queue a new WAL segment. This meant that the queue would never be cleared without manual intervention (such as calling `archive-push` directly). PostgreSQL now receives errors when there is not enough space to store new WAL segments but the async process will still be started so that space is eventually freed. 

- Fixed a remote timeout that occurred when a local process generated checksums (during resume or restore) but did not copy files, allowing the remote to go idle. 

### Core Features

- Non-exclusive backups will automatically be used on PostgreSQL 9.6. 

- Added the `cmd-ssh` option to allow the ssh client to be specified. 

- Added the `log-level-stderr` option to control whether console log messages are sent to `stderr` or `stdout`. By default this is set to `warn` which represents a change in behavior from previous versions, even though it may be more intuitive. Setting `log-level-stderr=off` will preserve the old behavior. 

- Set `application_name` to `"pgBackRest [command]"` for database connections. 

- Check that archive_mode is enabled when `archive-check` option enabled. 

### Core Improvements

- Clarified error message when unable to acquire pgBackRest advisory lock to make it clear that it is not a PostgreSQL backup lock. 

- pgBackRest version number included in command start INFO log output. 

- Process ID logged for local process start/stop INFO log output. 

### Documentation Features

- Added `archive-timeout` option documentation to the user guide. 

### Documentation Development

- Added `dev` option to `doc.pl` to easily add the `dev` keyword to documentation builds. 

### Test Development

- Update CentOS/Debian package definitions. 

- Fixed missing expect output for help module. 

- Fixed broken `vm-max` option in `test.pl`. 

- Regression tests can now be run as any properly-configured user, not just vagrant. 

- Minimize TeXLive package list to save time during VM builds. 

## v1.08 — Bug Fixes and Log Improvements

*Released: 2016-09-14*

### Core Bug Fixes

- Fixed an issue where local processes were not disconnecting when complete and could later timeout. 

- Fixed an issue where the protocol layer could timeout while waiting for WAL segments to arrive in the archive. 

### Core Improvements

- Cache file log output until the file is created to create a more complete log. 

### Documentation Development

- Show Process ID in output instead of filtering it out with the timestamp. 

### Test Development

- Suppress "dpkg-reconfigure: unable to re-open stdin: No file or directory" warning in Vagrant VM build. 

- Show Process ID in expect logs instead of filtering it out with the timestamp. 

## v1.07 — Thread to Process Conversion and Bug Fixes

*Released: 2016-09-07*

### Core Bug Fixes

- Fixed an issue where tablespaces were copied from the primary during standby backup. 

- Fixed the `check` command so backup info is checked remotely and not just locally. 

- Fixed an issue where `retention-archive` was not automatically being set when `retention-archive-type=diff`, resulting in a less aggressive than intended expiration of archive. 

### Core Features

- Converted Perl threads to processes to improve compatibility and performance. 

- Exclude contents of `$PGDATA/pg_replslot` directory so that replication slots on the primary do not become part of the backup. 

- The `archive-start` and `archive-stop` settings are now filled in `backup.manifest` even when `archive-check=n`. 

- Additional warnings when archive retention settings may not have the intended effect or would allow indefinite retention. 

- Experimental support for non-exclusive backups in PostgreSQL 9.6 rc1. Changes to the control/catalog/WAL versions in subsequent release candidates may break compatibility but pgBackRest will be updated with each release to keep pace. 

### Core Development

- Refactor of protocol minions in preparation for the new local minion. 

- Remove obsolete thread index variable from `File()` module. 

- Changed temporary file names to consistently use the `.pgbackrest.tmp` extension even if the destination file is compressed or has an appended checksum. 

- Improve ASSERT error handling, safely check eval blocks, and convert `$@` to `$EVAL_ERROR`. 

### Documentation Bug Fixes

- Fixed minor documentation reproducibility issues related to binary paths. 

### Documentation Features

- Documentation for archive retention. 

### Documentation Development

- Suppress TOC for unsupported versions of pgBackRest. 

### Test Development

- New vagrant base box and make uid/gid selection for containers dynamic. 

## v1.06 — Backup from Standby and Bug Fixes

*Released: 2016-08-25*

### Core Bug Fixes

- Fixed an issue where a tablespace link that referenced another link would not produce an error, but instead skip the tablespace entirely. 

- Fixed an issue where options that should not allow multiple values could be specified multiple times in `pgbackrest.conf` without an error being raised. 

- Fixed an issue where the `protocol-timeout` option was not automatically increased when the `db-timeout` option was increased. 

### Core Features

- Backup from a standby cluster. A connection to the primary cluster is still required to start/stop the backup and copy files that are not replicated, but the vast majority of files are copied from the standby in order to reduce load on the primary. 

- More flexible configuration for databases. Master and standby can both be configured on the backup server and pgBackRest will automatically determine which is the primary. This means no configuration changes for backup are required after failing over from a primary to standby when a separate backup server is used. 

- Exclude directories during backup that are cleaned, recreated, or zeroed by PostgreSQL at startup. These include `pgsql_tmp` and `pg_stat_tmp`. The `postgresql.auto.conf.tmp` file is now excluded in addition to files that were already excluded: `backup_label.old`, `postmaster.opts`, `postmaster.pid`, `recovery.conf`, `recovery.done`. 

- Experimental support for non-exclusive backups in PostgreSQL 9.6 beta4. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

### Core Improvements

- Improve error message for links that reference links in manifest build. 

- Added hints to error message when relative paths are detected in `archive-push` or `archive-get`. 

- Improve backup log messages to indicate which host the files are being copied from. 

### Core Development

- Simplify protocol creation and identifying which host is local/remote. 

- Removed all `OP_*` function constants that were used only for debugging, not in the protocol, and replaced with `__PACKAGE__`. 

- Improvements in `Db` module: separated out `connect()` function, allow `executeSql()` calls that do not return data, and improve error handling. 

### Documentation Development

- Improve host tag rendering. 

### Test Development

- Refactor db version constants into a separate module. 

- Update synthetic backup tests to PostgreSQL 9.4. 

## v1.05 — Bug Fix for Tablespace Link Checking

*Released: 2016-08-09*

### Core Bug Fixes

- Fixed an issue where tablespace paths that had $PGDATA as a substring would be identified as a subdirectories of $PGDATA even when they were not. Also hardened relative path checking a bit. 

### Documentation Features

- Added documentation for scheduling backups with cron. 

### Documentation Improvements

- Moved the backlog from the pgBackRest website to the GitHub repository wiki. 

### Documentation Development

- Improved rendering of spaces in code blocks. 

## v1.04 — Various Bug Fixes

*Released: 2016-07-30*

### Core Bug Fixes

- Fixed an issue an where an extraneous remote was created causing threaded backup/restore to possibly timeout and/or throw a lock conflict. 

- Fixed an issue where db-path was not required for the `check` command so an assert was raised when it was missing rather than a polite error message. 

- Fixed `check` command to throw an error when database version/id does not match that of the archive. 

- Fixed an issue where a remote could try to start its own remote when the `backup-host` option was not present in `pgbackrest.conf` on the database server. 

- Fixed an issue where the contents of `pg_xlog` were being backed up if the directory was symlinked. This didn't cause any issues during restore but was a waste of space. 

- Fixed an invalid `log()` call in lock routines. 

### Core Features

- Experimental support for non-exclusive backups in PostgreSQL 9.6 beta3. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

### Core Improvements

- Suppress banners on SSH protocol connections. 

- Improved remote error messages to identify the host where the error was raised. 

- All remote types now take locks. The exceptions date to when the test harness and pgBackRest were running in the same VM and no longer apply. 

### Core Development

- Enhancements to the protocol layer for improved reliability and error handling. 

- Exceptions are now passed back from threads as messages when possible rather than raised directly. 

- Temp files created during backup are now placed in the same directory as the target file. 

- Output lock file name when a lock cannot be acquired to aid in debugging. 

- Reduce calls to `protocolGet()` in backup/restore. 

### Documentation Features

- Added clarification on why the default for the `backrest-user` option is `backrest`. 

- Updated information about package availability on supported platforms. 

### Documentation Development

- Added `release.pl` to make releases reproducible. For now this only includes building and deploying documentation. 

- HTML footer dates are statically created in English in order to be reproducible. 

### Test Development

- Fixed a version checking issue in `test.pl`. 

- Fixed an issue where multi-threaded tests were not being run when requested. 

- Reduce the frequency that certain tests are run to save time in regression. 

- Disable control master for older OS versions where it is less stable. 

## v1.03 — Check Command and Bug Fixes

*Released: 2016-07-02*

### Core Bug Fixes

- Fixed an issue where `keep-alives` could be starved out by lots of small files during multi-threaded `backup`. They were also completely absent from single/multi-threaded `backup` resume and `restore` checksumming. 

- Fixed an issue where the `expire` command would refuse to run when explicitly called from the command line if the `db-host` option was set. This was not an issue when `expire` was run automatically after a `backup` 

- Fixed an issue where validation was being running on `archive_command` even when the `archive-check` option was disabled. 

### Core Features

- Added `check` command to validate that pgBackRest is configured correctly for archiving and backups. 

- Added the `protocol-timeout` option. Previously `protocol-timeout` was set as `db-timeout` + 30 seconds. 

- Failure to shutdown remotes at the end of the backup no longer throws an exception. Instead a warning is generated that recommends a higher `protocol-timeout`. 

- Experimental support for non-exclusive backups in PostgreSQL 9.6 beta2. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

### Core Improvements

- Improved handling of users/groups captured during backup that do not exist on the restore host. Also explicitly handle the case where user/group is not mapped to a name. 

- Option handling is now far more strict. Previously it was possible for a command to use an option that was not explicitly assigned to it. This was especially true for the `backup-host` and `db-host` options which are used to determine locality. 

### Core Development

- The `pg_xlogfile_name()` function is no longer used to construct WAL filenames from LSNs. While this function is convenient it is not available on a standby. Instead, the archive is searched for the LSN in order to find the timeline. If due to some misadventure the LSN appears on multiple timelines then an error will be thrown, whereas before this condition would have passed unnoticed. 

- Changed version variable to a constant. It had originally been designed to play nice with a specific packaging tool but that tool was never used. 

### Documentation Improvements

- Allow a static date to be used for documentation to generate reproducible builds. 

- Added documentation for asynchronous archiving to the user guide. 

- Recommended install location for pgBackRest modules is now `/usr/share/perl5` since `/usr/lib/perl5` has been removed from the search path in newer versions of Perl. 

- Added instructions for removing prior versions of pgBackRest. 

### Documentation Development

- Fixed DTD search path that did not work properly when `--doc-path` was used. 

- Fixed pgBackRest-specific xml that was loaded for non-pgBackRest projects. 

- Fixed section names being repeated in the info output when multiple `--require` options depended on the same sections. 

- Fixed pgBackRest config sections being blank in the output when not loaded from cache. 

- Allow hidden options to be added to a command. This allows certain commands (like `apt-get`) to be forced during the build without making that a part of the documentation. 

- Allow command summaries to be inserted anywhere in the documentation to avoid duplication. 

- Update TeX Live to 2016 version. 

- New, consolidated implementation for link rendering. 

- PostgreSQL version is now a variable to allow multi-version documentation. 

### Test Development

- Obsolete containers are removed by the `--vm-force` option. 

- Major refactor of the test suite to make it more modular and object-oriented. Multiple Docker containers can now be created for a single test to simulate more realistic environments. Tests paths have been renamed for clarity. 

- Greatly reduced the quantity of Docker containers built by default. Containers are only built for PostgreSQL versions specified in `db-minimal` and those required to build documentation. Additional containers can be built with `--db-version=all` or by specifying a version, e.g. `--db-version=9.4`. 

## v1.02 — Bug Fix for Perl 5.22

*Released: 2016-06-02*

### Core Bug Fixes

- Fix usage of sprintf() due to new constraints in Perl 5.22. Parameters not referenced in the format string are no longer allowed. 

### Core Development

- Log directory create and file open now using FileCommon functions which produce more detailed error messages on failure. 

### Documentation Bug Fixes

- Fixed syntax that was not compatible with Perl 5.2X. 

- Fixed absolute paths that were used for the PDF logo. 

### Documentation Features

- Release notes are now broken into sections so that bugs, features, and refactors are clearly delineated. An "Additional Notes" section has been added for changes to documentation and the test suite that do not affect the core code. 

- Added man page generation. 

- The change log was the last piece of documentation to be rendered in Markdown only. Wrote a converter so the document can be output by the standard renderers. The change log will now be located on the website and has been renamed to "Releases". 

### Documentation Development

- Added an execution cache so that documentation can be generated without setting up the full container environment. This is useful for packaging, keeps the documentation consistent for a release, and speeds up generation when no changes are made in the execution list. 

- Remove function constants and pass strings directly to logDebugParam(). The function names were only used once so creating constants for them was wasteful. 

- Lists can now be used outside of `p` and `text` tags for more flexible document structuring. 

### Test Development

- Replaced overzealous `perl -cW` check which failed on Perl 5.22 with `perl -cw`. 

- Added Ubuntu 16.04 (Xenial) and Debian 8 (Jessie) to the regression suite. 

- Upgraded doc/test VM to Ubuntu 16.04. This will help catch Perl errors in the doc code since it is not run across multiple distributions like the core and test code. It is also to be hoped that a newer kernel will make Docker more stable. 

- Test release version against the executable using `change-log.xml` instead of `CHANGELOG.md`. 

## v1.01 — Enhanced Info, Selective Restore, and 9.6 Support

*Released: 2016-05-17*

### Core Features

- Enhanced text output of `info` command to include timestamps, sizes, and the reference list for all backups. 

- Allow selective restore of databases from a cluster backup. This feature can result in major space and time savings when only specific databases are restored. Unrestored databases will not be accessible but must be manually dropped before they will be removed from the shared catalogue. 

- Experimental support for non-exclusive backups in PostgreSQL 9.6 beta1. Changes to the control/catalog/WAL versions in subsequent betas may break compatibility but pgBackRest will be updated with each release to keep pace. 

## v1.00 — New Repository Format and Configuration Scheme, Link Support

*Released: 2016-04-14*

### Core Features

- Implemented a new configuration scheme which should be far simpler to use. See the User Guide and Configuration Reference for details but for a simple configuration all options can now be placed in the `stanza` section. Options that are shared between stanzas can be placed in the `[global]` section. More complex configurations can still make use of command sections though this should be a rare use case. 

- The `repo-path` option now always refers to the repository where backups and archive are stored, whether local or remote, so the `repo-remote-path` option has been removed. The new `spool-path` option can be used to define a location for queueing WAL segments when archiving asynchronously. A local repository is no longer required. 

- The default configuration filename is now `pgbackrest.conf` instead of `pg_backrest.conf`. This was done for consistency with other naming changes but also to prevent old config files from being loaded accidentally when migrating to `1.00`. 

- The default repository name was changed from `/var/lib/backup` to `/var/lib/pgbackrest`. 

- Lock files are now stored in `/tmp/pgbackrest` by default. These days `/run/pgbackrest` is the preferred location but that would require init scripts which are not part of this release. The `lock-path` option can be used to configure the lock directory. 

- Log files are now stored in `/var/log/pgbackrest` by default and no longer have the date appended so they can be managed with `logrotate`. The `log-path` option can be used to configure the log directory. 

- Executable filename changed from `pg_backrest` to `pgbackrest`. 

- All files and directories linked from PGDATA are now included in the backup. By default links will be restored directly into PGDATA as files or directories. The `--link-all` option can be used to restore all links to their original locations. The `--link-map` option can be used to remap a link to a new location. 

- Removed `--tablespace` option and replaced with `--tablespace-map-all` option which should more clearly indicate its function. 

- Added `detail` log level which will output more information than `info` without being as verbose as `debug`. 

## v0.92 — Command-line Repository Path Fix

*Released: 2016-04-06*

### Core Bug Fixes

- Fixed an issue where the master process was passing `--repo-remote-path` instead of `--repo-path` to the remote and causing the lock files to be created in the default repository directory (`/var/lib/backup`), generally ending in failure. This was only an issue when `--repo-remote-path` was defined on the command line rather than in `pg_backrest.conf`. 

## v0.91 — Tablespace Bug Fix and Minor Enhancements

*Released: 2016-03-22*

### Core Bug Fixes

- Fixed repository incompatibility introduced in pgBackRest 0.90. 

### Core Features

- Copy `global/pg_control` last during backups. 

- Write `.info` and `.manifest` files to temp before moving them to their final locations and fsync'ing. 

- Rename `--no-start-stop` option to `--no-online`. 

### Test Features

- Static source analysis using Perl-Critic, currently passes on gentle. 

## v0.90 — 9.5 Support, Various Enhancements, and Minor Bug Fixes

*Released: 2016-02-07*

### Core Bug Fixes

- Fixed an issue where specifying `--no-archive-check` would throw a configuration error. 

- Fixed an issue where a temp WAL file left over after a well-timed system crash could cause the next `archive-push` to fail. 

- The `retention-archive` option can now be be safely set to less than backup retention (`retention-full` or `retention-diff`) without also specifying `archive-copy=n`. The WAL required to make the backups that fall outside of archive retention consistent will be preserved in the archive. However, in this case PITR will not be possible for the backups that fall outside of archive retention. 

### Core Features

- When backing up and restoring tablespaces pgBackRest only operates on the subdirectory created for the version of PostgreSQL being run against. Since multiple versions can live in a tablespace (especially during a binary upgrade) this prevents too many files from being copied during a backup and other versions possibly being wiped out during a restore. This only applies to PostgreSQL >= 9.0 &mdash; prior versions of PostgreSQL could not share a tablespace directory. 

- Generate an error when `archive-check=y` but `archive_command` does not execute `pg_backrest`. 

- Improved error message when `repo-path` or `repo-remote-path` does not exist. 

- Added checks for `--delta` and `--force` restore options to ensure that the destination is a valid $PGDATA directory. pgBackRest will check for the presence of `PG_VERSION` or `backup.manifest` (left over from an aborted restore). If neither file is found then `--delta` and `--force` will be disabled but the restore will proceed unless there are files in the $PGDATA directory (or any tablespace directories) in which case the operation will be aborted. 

- When restore `--set=latest` (the default) the actual backup restored will be output to the log. 

- Support for PostgreSQL 9.5 partial WAL segments and `recovery_target_action` setting. The `archive_mode = 'always'` setting is not yet supported. 

- Support for `recovery_target = 'immediate'` recovery setting introduced in PostgreSQL 9.4. 

- The following tablespace checks have been added: paths or files in pg_tblspc, relative links in pg_tblspc, tablespaces in $PGDATA. All three will generate errors. 

### Documentation Development

- Fixed an issue where document generation failed because some OSs are not tolerant of having multiple installed versions of PostgreSQL. A separate VM is now created for each version. Also added a sleep after database starts during document generation to ensure the database is running before the next command runs. 

## v0.89 — Timeout Bug Fix and Restore Read-Only Repositories

*Released: 2015-12-24*

### Core Bug Fixes

- Fixed an issue where longer-running backups/restores would timeout when remote and threaded. Keepalives are now used to make sure the remote for the main process does not timeout while the thread remotes do all the work. The error message for timeouts was also improved to make debugging easier. 

### Core Features

- Allow restores to be performed on a read-only repository by using `--no-lock` and `--log-level-file=off`. The `--no-lock` option can only be used with restores. 

### Documentation Development

- Minor styling changes, clarifications and rewording in the user guide. 

### Test Development

- The dev branch has been renamed to master and for the time being the master branch has renamed to release, though it will probably be removed at some point -- thus ends the gitflow experiment for pgBackRest. It is recommended that any forks get re-forked and clones get re-cloned. 

## v0.88 — Documentation and Minor Bug Fixes

*Released: 2015-11-22*

### Core Bug Fixes

- Fixed an issue where the `start`/`stop` commands required the `--config` option. 

- Fixed an issue where log files were being overwritten instead of appended. 

- Fixed an issue where `backup-user` was not optional. 

### Core Features

- Symlinks are no longer created in backup directories in the repository. These symlinks could point virtually anywhere and potentially be dangerous. Symlinks are still recreated during a restore. 

- Added better messaging for backup expiration. Full and differential backup expirations are logged on a single line along with a list of all dependent backups expired. 

- Archive retention is automatically set to full backup retention if not explicitly configured. 

### Documentation Features

- Added documentation in the user guide for delta restores, expiration, dedicated backup hosts, starting and stopping pgBackRest, and replication. 

## v0.87 — Website and User Guide

*Released: 2015-10-28*

### Core Features

- The `backup_label.old` and `recovery.done` files are now excluded from backups. 

### Documentation Features

- Added a new user guide that covers pgBackRest basics and some advanced topics including PITR. Much more to come, but it's a start. 

### Documentation Development

- The website, markdown, and command-line help are now all generated from the same XML source. 

## v0.85 — Start/Stop Commands and Minor Bug Fixes

*Released: 2015-10-08*

### Core Bug Fixes

- Fixed an issue where an error could be returned after a backup or restore completely successfully. 

- Fixed an issue where a resume would fail if temp files were left in the root backup directory when the backup failed. This scenario was likely if the backup process got terminated during the copy phase. 

### Core Features

- Added `stop` and `start` commands to prevent pgBackRest processes from running on a system where PostgreSQL is shutdown or the system needs to be quiesced for some other reason. 

- Experimental support for PostgreSQL 9.5 beta1. This may break when the control version or WAL magic changes in future versions but will be updated in each pgBackRest release to keep pace. All regression tests pass except for `--target-resume` tests (this functionality has changed in 9.5) and there is no testing yet for `.partial` WAL segments. 

### Core Development

- Removed dependency on `IO::String` module. 

## v0.82 — Refactoring, Command-line Help, and Minor Bug Fixes

*Released: 2015-09-14*

### Core Bug Fixes

- Fixed an issue where resumed compressed backups were not preserving existing files. 

- Fixed an issue where resume and incr/diff would not ensure that the prior backup had the same compression and hardlink settings. 

- Fixed an issue where a cold backup using `--no-start-stop` could be started on a running PostgreSQL cluster without `--force` specified. 

- Fixed an issue where a thread could be started even when none were requested. 

- Fixed an issue where the pgBackRest version number was not being updated in `backup.info` and `archive.info` after an upgrade/downgrade. 

- Fixed an issue where the `info` command was throwing an exception when the repository contained no stanzas. 

- Fixed an issue where the PostgreSQL `pg_stop_backup()` NOTICEs were being output to `stderr`. 

### Core Features

- Experimental support for PostgreSQL 9.5 alpha2. This may break when the control version or WAL magic changes in future versions but will be updated in each pgBackRest release to keep pace. All regression tests pass except for `--target-resume` tests (this functionality has changed in 9.5) and there is no testing yet for `.partial` WAL segments. 

### Core Improvements

- Renamed `recovery-setting` option and section to `recovery-option` to be more consistent with pgBackRest naming conventions. 

- Added dynamic module loading to speed up commands, especially asynchronous archiving. 

### Core Development

- Code cleanup and refactoring to standardize on patterns that have evolved over time. 

### Documentation Features

- Command-line help is now extracted from the same XML source that is used for the other documentation and includes much more detail. 

### Test Development

- Expiration tests are now synthetic rather than based on actual backups. This will allow development of more advanced expiration features. 

## v0.80 — DBI Support, Stability, and Convenience Features

*Released: 2015-08-09*

### Core Bug Fixes

- Fixed an issue that caused the formatted timestamp for both the oldest and newest backups to be reported as the current time by the `info` command. Only `text` output was affected -- `json` output reported the correct epoch values. 

- Fixed protocol issue that was preventing ssh errors (especially on connection) from being logged. 

### Core Features

- The repository is now created and updated with consistent directory and file modes. By default `umask` is set to `0000` but this can be disabled with the `neutral-umask` setting. 

- Added the `stop-auto` option to allow failed backups to automatically be stopped when a new backup starts. 

- Added the `db-timeout` option to limit the amount of time pgBackRest will wait for `pg_start_backup()` and `pg_stop_backup()` to return. 

- Remove `pg_control` file at the beginning of the restore and copy it back at the very end. This prevents the possibility that a partial restore can be started by PostgreSQL. 

- Added checks to be sure the `db-path` setting is consistent with `db-port` by comparing the `data_directory` as reported by the cluster against the `db-path` setting and the version as reported by the cluster against the value read from `pg_control`. The `db-socket-path` setting is checked to be sure it is an absolute path. 

- Experimental support for PostgreSQL 9.5 alpha1. This may break when the control version or WAL magic changes in future versions but will be updated in each pgBackRest release to keep pace. All regression tests pass except for `--target-resume` tests (this functionality has changed in 9.5) and there is no testing yet for `.partial` WAL segments. 

### Core Improvements

- Now using Perl `DBI` and `DBD::Pg` for connections to PostgreSQL rather than `psql`. The `cmd-psql` and `cmd-psql-option` settings have been removed and replaced with `db-port` and `db-socket-path`. Follow the instructions in the Installation Guide to install `DBD::Pg` on your operating system. 

### Core Development

- Major refactoring of the protocol layer to support future development. 

### Documentation Development

- Split most of `README.md` out into `USERGUIDE.md` and `CHANGELOG.md` because it was becoming unwieldy. Changed most references to "database" in the user guide to "database cluster" for clarity. 

- Changed most references to "database" in the user guide to "database cluster" for clarity. 

### Test Features

- Added vagrant test configurations for Ubuntu 14.04 and CentOS 7. 

## v0.78 — Remove CPAN Dependencies, Stability Improvements

*Released: 2015-07-13*

### Core Improvements

- Removed dependency on CPAN packages for multi-threaded operation. While it might not be a bad idea to update the `threads` and `Thread::Queue` packages, it is no longer necessary. 

- Modified wait backoff to use a Fibonacci rather than geometric sequence. This will make wait time grow less aggressively while still giving reasonable values. 

### Test Features

- Added vagrant test configurations for Ubuntu 12.04 and CentOS 6. 

### Test Development

- More options for regression tests and improved code to run in a variety of environments. 

## v0.77 — CentOS/RHEL 6 Support and Protocol Improvements

*Released: 2015-06-30*

### Core Features

- Added file and directory syncs to the `File` object for additional safety during backup/restore and archiving. 

- Added support for Perl 5.10.1 and OpenSSH 5.3 which are default for CentOS/RHEL 6. 

- Improved error message when backup is run without `archive_command` set and without `--no-archive-check` specified. 

### Core Development

- Removed `pg_backrest_remote` and added the functionality to `pg_backrest` as the `remote` command. 

- Moved version number out of the `VERSION` file to `Version.pm` to better support packaging. 

- Replaced `IPC::System::Simple` and `Net::OpenSSH` with `IPC::Open3` to eliminate CPAN dependency for multiple operating systems. 

## v0.75 — New Repository Format, Info Command and Experimental 9.5 Support

*Released: 2015-06-14*

### Core Features

- Added the `info` command. 

- Logging now uses unbuffered output. This should make log files that are being written by multiple threads less chaotic. 

- Experimental support for PostgreSQL 9.5. This may break when the control version or WAL magic changes but will be updated in each release. 

### Core Improvements

- More efficient file ordering for `backup`. Files are copied in descending size order so a single thread does not end up copying a large file at the end. This had already been implemented for `restore`. 

## v0.70 — Stability Improvements for Archiving, Improved Logging and Help

*Released: 2015-06-01*

### Core Bug Fixes

- Fixed an issue where `archive-copy` would fail on an incr/diff backup when `hardlink=n`. In this case the `pg_xlog` path does not already exist and must be created. 

- Fixed an issue in async archiving where `archive-push` was not properly returning 0 when `archive-max-mb` was reached and moved the async check after transfer to avoid having to remove the stop file twice. Also added unit tests for this case and improved error messages to make it clearer to the user what went wrong. 

- Fixed a locking issue that could allow multiple operations of the same type against a single stanza. This appeared to be benign in terms of data integrity but caused spurious errors while archiving and could lead to errors in backup/restore. 

### Core Features

- Allow duplicate WAL segments to be archived when the checksum matches. This is necessary for some recovery scenarios. 

- Allow comments/disabling in `pg_backrest.conf` using the `#` character. Only `#` characters in the first character of the line are honored. 

- Better logging before `pg_start_backup()` to make it clear when the backup is waiting on a checkpoint. 

- Various command behavior and logging fixes. 

### Core Improvements

- Replaced `JSON` module with `JSON::PP` which ships with core Perl. 

### Documentation Bug Fixes

- Various help fixes. 

## v0.65 — Improved Resume and Restore Logging, Compact Restores

*Released: 2015-05-11*

### Core Bug Fixes

- Fixed an issue where an absolute path was not written into `recovery.conf` when the restore was run with a relative path. 

### Core Features

- Better resume support. Resumed files are checked to be sure they have not been modified and the manifest is saved more often to preserve checksums as the backup progresses. More unit tests to verify each resume case. 

- Resume is now optional. Use the `resume` setting or `--no-resume` from the command line to disable. 

- More info messages during restore. Previously, most of the restore messages were debug level so not a lot was output in the log. 

- Added `tablespace` setting to allow tablespaces to be restored into the `pg_tblspc` path. This produces compact restores that are convenient for development, staging, etc. Currently these restores cannot be backed up as pgBackRest expects only links in the `pg_tblspc` path. 

## v0.61 — Bug Fix for Uncompressed Remote Destination

*Released: 2015-04-21*

### Core Bug Fixes

- Fixed a buffering error that could occur on large, highly-compressible files when copying to an uncompressed remote destination. The error was detected in the decompression code and resulted in a failed backup rather than corruption so it should not affect successful backups made with previous versions. 

## v0.60 — Better Version Support and WAL Improvements

*Released: 2015-04-19*

### Core Bug Fixes

- Pushing duplicate WAL now generates an error. This worked before only if checksums were disabled. 

### Core Features

- Database System IDs are used to make sure that all WAL in an archive matches up. This should help prevent misconfigurations that send WAL from multiple clusters to the same archive. 

### Core Development

- Improved threading model by starting threads early and terminating them late. 

### Test Features

- Regression tests working back to PostgreSQL 8.3. 

## v0.50 — Restore and Much More

*Released: 2015-03-25*

### Core Bug Fixes

- Fixed broken checksums and now they work with normal and resumed backups. Finally realized that checksums and checksum deltas should be functionally separated and this simplified a number of things. Issue #28 has been created for checksum deltas. 

- Fixed an issue where a backup could be resumed from an aborted backup that didn't have the same type and prior backup. 

### Core Features

- Added restore functionality. 

- All options can now be set on the command-line making `pg_backrest.conf` optional. 

- De/compression is now performed without threads and checksum/size is calculated in stream. That means file checksums are no longer optional. 

- Added option `--no-start-stop` to allow backups when Postgres is shut down. If `postmaster.pid` is present then `--force` is required to make the backup run (though if Postgres is running an inconsistent backup will likely be created). This option was added primarily for the purpose of unit testing, but there may be applications in the real world as well. 

- Checksum for `backup.manifest` to detect a corrupted/modified manifest. 

- Link `latest` always points to the last backup. This has been added for convenience and to make restores simpler. 

### Core Development

- Removed dependency on `Moose`. It wasn't being used extensively and makes for longer startup times. 

### Test Features

- More comprehensive unit tests in all areas. 

## v0.30 — Core Restructuring and Unit Tests

*Released: 2014-10-05*

### Core Development

- Complete rewrite of `BackRest::File` module to use a custom protocol for remote operations and Perl native GZIP and SHA operations. Compression is performed in threads rather than forked processes. 

- Removed dependency on `Storable` and replaced with a custom ini file implementation. 

- Numerous other changes that can only be identified with a diff. 

### Documentation Features

- Added much needed documentation 

### Test Features

- Fairly comprehensive unit tests for all the basic operations. More work to be done here for sure, but then there is always more work to be done on unit tests. 

## v0.19 — Improved Error Reporting/Handling

*Released: 2014-05-13*

### Core Bug Fixes

- Found and squashed a nasty bug where `file_copy()` was defaulted to ignore errors. There was also an issue in `file_exists()` that was causing the test to fail when the file actually did exist. Together they could have resulted in a corrupt backup with no errors, though it is very unlikely. 

### Core Development

- Worked on improving error handling in the `File` object. This is not complete, but works well enough to find a few errors that have been causing us problems (notably, find is occasionally failing building the archive async manifest when system is under load). 

## v0.18 — Return Soft Error When Archive Missing

*Released: 2014-04-13*

### Core Bug Fixes

- The `archive-get` command now returns a 1 when the archive file is missing to differentiate from hard errors (ssh connection failure, file copy error, etc.) This lets PostgreSQL know that the archive stream has terminated normally. However, this does not take into account possible holes in the archive stream. 

## v0.17 — Warn When Archive Directories Cannot Be Deleted

*Released: 2014-04-03*

### Core Bug Fixes

- If an archive directory which should be empty could not be deleted backrest was throwing an error. There's a good fix for that coming, but for the time being it has been changed to a warning so processing can continue. This was impacting backups as sometimes the final archive file would not get pushed if the first archive file had been in a different directory (plus some bad luck). 

## v0.16 — RequestTTY=yes for SSH Sessions

*Released: 2014-04-01*

### Core Bug Fixes

- Added `RequestTTY=yes` to ssh sessions. Hoping this will prevent random lockups. 

## v0.15 — Added archive-get

*Released: 2014-03-29*

### Core Features

- Added `archive-get` functionality to aid in restores. 

- Added option to force a checkpoint when starting the backup, `start-fast=y`. 

## v0.11 — Minor Fixes

*Released: 2014-03-26*

### Core Bug Fixes

- Removed `master_stderr_discard` option on database SSH connections. There have been occasional lockups and they could be related to issues originally seen in the file code. 

- Changed lock file conflicts on `backup` and `expire` commands to `ERROR`. They were set to `DEBUG` due to a copy-and-paste from the archive locks. 

## v0.10 — Backup and Archiving are Functional

*Released: 2014-03-05*

### Core Features

- No restore functionality, but the backup directories are consistent PostgreSQL data directories. You'll need to either uncompress the files or turn off compression in the backup. Uncompressed backups on a ZFS (or similar) filesystem are a good option because backups can be restored locally via a snapshot to create logical backups or do spot data recovery. 

- Archiving is single-threaded. This has not posed an issue on our multi-terabyte databases with heavy write volume. Recommend a large WAL volume or to use the async option with a large volume nearby. 

- Backups are multi-threaded, but the `Net::OpenSSH` library does not appear to be 100% thread-safe so it will very occasionally lock up on a thread. There is an overall process timeout that resolves this issue by killing the process. Yes, very ugly. 

- Checksums are lost on any resumed backup. Only the final backup will record checksum on multiple resumes. Checksums from previous backups are correctly recorded and a full backup will reset everything. 

- The `backup.manifest` is being written as `Storable` because `Config::IniFile` does not seem to handle large files well. Would definitely like to save these as human-readable text. 

### Documentation Features

- Absolutely no documentation (outside the code). Well, excepting these release notes.
