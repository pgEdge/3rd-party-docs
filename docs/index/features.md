# Features
<a name="features"></a>

## Parallel Backup & Restore
<a name="parallel-backup-restore"></a>

Compression is usually the bottleneck during backup operations so pgBackRest solves this problem with parallel processing and more efficient compression algorithms such as lz4 and zstd.

## Local or Remote Operation
<a name="local-or-remote"></a>

A custom protocol allows pgBackRest to backup, restore, and archive locally or remotely via TLS/SSH with minimal configuration. An interface to query PostgreSQL is also provided via the protocol layer so that remote access to PostgreSQL is never required, which enhances security.

## Multiple Repositories
<a name="multi-repo"></a>

Multiple repositories allow, for example, a local repository with minimal retention for fast restores and a remote repository with a longer retention for redundancy and access across the enterprise.

## Full, Differential, & Incremental Backups (at File or Block Level)
<a name="backup-types"></a>

Full, differential, and incremental backups are supported. pgBackRest is not susceptible to the time resolution issues of rsync, making differential and incremental backups safe without the requirement to checksum each file. Block-level backups save space by only copying the parts of files that have changed.

## Backup Rotation & Archive Expiration
<a name="backup-rotation"></a>

Retention polices can be set for full and differential backups to create coverage for any time frame. The WAL archive can be maintained for all backups or strictly for the most recent backups. In the latter case WAL required to make older backups consistent will be maintained in the archive.

## Backup Integrity
<a name="backup-integrity"></a>

Checksums are calculated for every file in the backup and rechecked during a restore or verify. After a backup finishes copying files, it waits until every WAL segment required to make the backup consistent reaches the repository.

Backups in the repository may be stored in the same format as a standard PostgreSQL cluster (including tablespaces). If compression is disabled and hard links are enabled it is possible to snapshot a backup in the repository and bring up a PostgreSQL cluster directly on the snapshot. This is advantageous for terabyte-scale databases that are time consuming to restore in the traditional way.

All operations utilize file and directory level fsync to ensure durability.

## Page Checksums
<a name="page-checksum"></a>

If page checksums are enabled pgBackRest will validate the checksums for every file that is copied during a backup. All page checksums are validated during a full backup and checksums in files that have changed are validated during differential and incremental backups.

Validation failures do not stop the backup process, but warnings with details of exactly which pages have failed validation are output to the console and file log.

This feature allows page-level corruption to be detected early, before backups that contain valid copies of the data have expired.

## Backup Resume
<a name="backup-resume"></a>

An interrupted backup can be resumed from the point where it was stopped. Files that were already copied are compared with the checksums in the manifest to ensure integrity. Since this operation can take place entirely on the repository host, it reduces load on the PostgreSQL host and saves time since checksum calculation is faster than compressing and retransmitting data.

## Streaming Compression & Checksums
<a name="stream-compression-checksums"></a>

Compression and checksum calculations are performed in stream while files are being copied to the repository, whether the repository is located locally or remotely.

If the repository is on a repository host, compression is performed on the PostgreSQL host and files are transmitted in a compressed format and simply stored on the repository host. When compression is disabled a lower level of compression is utilized to make efficient use of available bandwidth while keeping CPU cost to a minimum.

## Delta Restore
<a name="delta-restore"></a>

The manifest contains checksums for every file in the backup so that during a restore it is possible to use these checksums to speed processing enormously. On a delta restore any files not present in the backup are first removed and then checksums are generated for the remaining files. Files that match the backup are left in place and the rest of the files are restored as usual. Parallel processing can lead to a dramatic reduction in restore times.

## Parallel, Asynchronous WAL Push & Get
<a name="parallel-archiving"></a>

Dedicated commands are included for pushing WAL to the archive and getting WAL from the archive. Both commands support parallelism to accelerate processing and run asynchronously to provide the fastest possible response time to PostgreSQL.

WAL push automatically detects WAL segments that are pushed multiple times and de-duplicates when the segment is identical, otherwise an error is raised. Asynchronous WAL push allows transfer to be offloaded to another process which compresses WAL segments in parallel for maximum throughput. This can be a critical feature for databases with extremely high write volume.

Asynchronous WAL get maintains a local queue of WAL segments that are decompressed and ready for replay. This reduces the time needed to provide WAL to PostgreSQL which maximizes replay speed. Higher-latency connections and storage (such as S3) benefit the most.

The push and get commands both ensure that the database and repository match by comparing PostgreSQL versions and system identifiers. This virtually eliminates the possibility of misconfiguring the WAL archive location.

## Tablespace & Link Support
<a name="tablespace-link-support"></a>

Tablespaces are fully supported and on restore tablespaces can be remapped to any location. It is also possible to remap all tablespaces to one location with a single command which is useful for development restores.

File and directory links are supported for any file or directory in the PostgreSQL cluster. When restoring it is possible to restore all links to their original locations, remap some or all links, or restore some or all links as normal files or directories within the cluster directory.

## S3, Azure, and GCS Compatible Object Store Support
<a name="object-store-support"></a>

pgBackRest repositories can be located in S3, Azure, and GCS compatible object stores to allow for virtually unlimited capacity and retention.

## Encryption
<a name="encryption"></a>

pgBackRest can encrypt the repository to secure backups wherever they are stored.

## Compatibility with ten versions of
<a name="postgres-compatibility"></a>

pgBackRest includes support for ten versions of PostgreSQL, the five supported versions and the last five EOL versions. This allows ample time to upgrade to a supported version.
