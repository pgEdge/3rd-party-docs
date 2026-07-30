<a id="archive-modules"></a>

# Archive Modules

 PostgreSQL provides infrastructure to create custom modules for continuous archiving (see [Continuous Archiving and Point-in-Time Recovery (PITR)](../../server-administration/backup-and-restore/continuous-archiving-and-point-in-time-recovery-pitr.md#continuous-archiving)). While archiving via a shell command (i.e., [archive_command](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-command)) is much simpler, a custom archive module will often be considerably more robust and performant.

 When a custom [archive_library](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-library) is configured, PostgreSQL will submit completed WAL files to the module, and the server will avoid recycling or removing these WAL files until the module indicates that the files were successfully archived. It is ultimately up to the module to decide what to do with each WAL file, but many recommendations are listed at [Setting Up WAL Archiving](../../server-administration/backup-and-restore/continuous-archiving-and-point-in-time-recovery-pitr.md#backup-archiving-wal).

 Archiving modules must at least consist of an initialization function (see [Initialization Functions](initialization-functions.md#archive-module-init)) and the required callbacks (see [Archive Module Callbacks](archive-module-callbacks.md#archive-module-callbacks)). However, archive modules are also permitted to do much more (e.g., declare GUCs and register background workers).

 The `contrib/basic_archive` module contains a working example, which demonstrates some useful techniques.

- [Initialization Functions](initialization-functions.md#archive-module-init)
- [Archive Module Callbacks](archive-module-callbacks.md#archive-module-callbacks)
