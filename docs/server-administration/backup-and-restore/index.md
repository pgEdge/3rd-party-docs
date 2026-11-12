<a id="backup"></a>

# Backup and Restore

 As with everything that contains valuable data, PostgreSQL databases should be backed up regularly. While the procedure is essentially simple, it is important to have a clear understanding of the underlying techniques and assumptions.

 There are three fundamentally different approaches to backing up PostgreSQL data:

- SQL dump
- File system level backup
- Continuous archiving
 Each has its own strengths and weaknesses; each is discussed in turn in the following sections.

- [SQL Dump](sql-dump.md#backup-dump)
- [File System Level Backup](file-system-level-backup.md#backup-file)
- [Continuous Archiving and Point-in-Time Recovery (PITR)](continuous-archiving-and-point-in-time-recovery-pitr.md#continuous-archiving)
