<a id="glossary"></a>

# Glossary


 This is a list of terms and their meaning in the context of PostgreSQL and relational database systems in general.
  <a id="glossary-acid"></a>

**ACID**
:   *Atomicity*, *Consistency*, *Isolation*, and *Durability*. This set of properties of database transactions is intended to guarantee validity in concurrent operation and even in event of errors, power failures, etc.
 <a id="glossary-aggregate"></a>

**Aggregate function (routine)**
:   A *function* that combines (*aggregates*) multiple input values, for example by counting, averaging or adding, yielding a single output value.


     For more information, see [Aggregate Functions](../the-sql-language/functions-and-operators/aggregate-functions.md#functions-aggregate).
         *See also: *
 <a id="glossary-am"></a>

**Access Method**
:   Interfaces which PostgreSQL use in order to access data in tables and indexes. This abstraction allows for adding support for new types of data storage.


     For more information, see [Table Access Method Interface Definition](../internals/table-access-method-interface-definition.md#tableam) and [Index Access Method Interface Definition](../internals/index-access-method-interface-definition/index.md#indexam).


**Analytic function**
    *See: *
 <a id="glossary-analyze"></a>

**Analyze (operation)**
:   The act of collecting statistics from data in *tables* and other *relations* to help the *query planner* to make decisions about how to execute *queries*.


     (Don't confuse this term with the `ANALYZE` option to the [sql-explain](../reference/sql-commands/explain.md#sql-explain) command.)


     For more information, see [sql-analyze](../reference/sql-commands/analyze.md#sql-analyze).
 <a id="glossary-application-time"></a>

**Application time**
:   In a *temporal table*, the dimension of time that represents when the entity described by the table changed (as opposed to the table itself).


     For more information, see [Temporal Tables](../the-sql-language/data-definition/temporal-tables.md#ddl-temporal-tables).
 <a id="glossary-aio"></a>

**Asynchronous I/O**
:   Asynchronous I/O (AIO) describes performing I/O in a non-blocking way (asynchronously), in contrast to synchronous I/O, which blocks for the entire duration of the I/O.


     With AIO, starting an I/O operation is separated from waiting for the result of the operation, allowing multiple I/O operations to be initiated concurrently, as well as performing CPU heavy operations concurrently with I/O. The price for that increased concurrency is increased complexity.
         *See also: *
 <a id="glossary-atomic"></a>

**Atomic**
:   In reference to a *datum*: the fact that its value cannot be broken down into smaller components.
 <a id="glossary-atomicity"></a>

**Atomicity**
:   The property of a *transaction* that either all its operations complete as a single unit or none do. In addition, if a system failure occurs during the execution of a transaction, no partial results are visible after recovery. This is one of the ACID properties.
 <a id="glossary-attribute"></a>

**Attribute**
:   An element with a certain name and data type found within a *tuple*.
 <a id="glossary-autovacuum"></a>

**Autovacuum (process)**
:   A set of background processes that routinely perform *vacuum* and *analyze* operations. The *auxiliary process* that coordinates the work and is always present (unless autovacuum is disabled) is known as the *autovacuum launcher*, and the processes that carry out the tasks are known as the *autovacuum workers*.


     For more information, see [The Autovacuum Daemon](../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#autovacuum).
 <a id="glossary-auxiliary-proc"></a>

**Auxiliary process**
:   A process within an *instance* that is in charge of some specific background task for the instance. The auxiliary processes consist of   the *autovacuum launcher* (but not the autovacuum workers), the *background writer*, the *checkpointer*, the *logger*, the *startup process*, the *WAL archiver*, the *WAL receiver* (but not the *WAL senders*), the *WAL summarizer*, and the *WAL writer*.
 <a id="glossary-backend"></a>

**Backend (process)**
:   Process of an *instance* which acts on behalf of a *client session* and handles its requests.


     (Don't confuse this term with the similar terms *Background Worker* or *Background Writer*).
 <a id="glossary-background-worker"></a>

**Background worker (process)**
:   Process within an *instance*, which runs system- or user-supplied code. Serves as infrastructure for several features in PostgreSQL, such as *logical replication* and *parallel queries*. In addition, *Extensions* can add custom background worker processes.


     For more information, see [Background Worker Processes](../server-programming/background-worker-processes.md#bgworker).
 <a id="glossary-background-writer"></a>

**Background writer (process)**
:   An *auxiliary process* that writes dirty *data pages* from *shared memory* to the file system. It wakes up periodically, but works only for a short period in order to distribute its expensive I/O activity over time to avoid generating larger I/O peaks which could block other processes.


     For more information, see [Background Writer](../server-administration/server-configuration/resource-consumption.md#runtime-config-resource-background-writer).
 <a id="glossary-basebackup"></a>

**Base Backup**
:   A binary copy of all *database cluster* files. It is generated by the tool [app-pgbasebackup](../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup). In combination with WAL files it can be used as the starting point for recovery, log shipping, or streaming replication.
 <a id="glossary-bloat"></a>

**Bloat**
:   Space in data pages which does not contain current row versions, such as unused (free) space or outdated row versions.
 <a id="glossary-bootstrap-superuser"></a>

**Bootstrap superuser**
:   The first *user* initialized in a *database cluster*.


     This user owns all system catalog tables in each database. It is also the role from which all granted permissions originate. Because of these things, this role may not be dropped.


     This role also behaves as a normal *database superuser*, and its superuser status cannot be removed.
 <a id="glossary-buffer-access-strategy"></a>

**Buffer Access Strategy**
:   Some operations will access a large number of *pages*. A *Buffer Access Strategy* helps to prevent these operations from evicting too many pages from *shared buffers*.


     A Buffer Access Strategy sets up references to a limited number of *shared buffers* and reuses them circularly. When the operation requires a new page, a victim buffer is chosen from the buffers in the strategy ring, which may require flushing the page's dirty data and possibly also unflushed *WAL* to permanent storage.


     Buffer Access Strategies are used for various operations such as sequential scans of large tables, `VACUUM`, `COPY`, `CREATE TABLE AS SELECT`, `ALTER TABLE`, `CREATE DATABASE`, `CREATE INDEX`, and `CLUSTER`.
 <a id="glossary-cast"></a>

**Cast**
:   A conversion of a *datum* from its current data type to another data type.


     For more information, see [sql-createcast](../reference/sql-commands/create-cast.md#sql-createcast).
 <a id="glossary-catalog"></a>

**Catalog**
:   The SQL standard uses this term to indicate what is called a *database* in PostgreSQL's terminology.


     (Don't confuse this term with *system catalog*).


     For more information, see [Overview](../server-administration/managing-databases/overview.md#manage-ag-overview).
 <a id="glossary-check-constraint"></a>

**Check constraint**
:   A type of *constraint* defined on a *relation* which restricts the values allowed in one or more *attributes*. The check constraint can make reference to any attribute of the same row in the relation, but cannot reference other rows of the same relation or other relations.


     For more information, see [Constraints](../the-sql-language/data-definition/constraints.md#ddl-constraints).
 <a id="glossary-checkpoint"></a>

**Checkpoint**
:   A point in the *WAL* sequence at which it is guaranteed that the heap and index data files have been updated with all information from *shared memory* modified before that checkpoint; a *checkpoint record* is written and flushed to WAL to mark that point.


     A checkpoint is also the act of carrying out all the actions that are necessary to reach a checkpoint as defined above. This process is initiated when predefined conditions are met, such as a specified amount of time has passed, or a certain volume of records has been written; or it can be invoked by the user with the command `CHECKPOINT`.


     For more information, see [WAL Configuration](../server-administration/reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration).
 <a id="glossary-checkpointer"></a>

**Checkpointer (process)**
:   An *auxiliary process* that is responsible for executing *checkpoints*.


**Class (archaic)**
    *See: *
 <a id="glossary-client"></a>

**Client (process)**
:   Any process, possibly remote, that establishes a *session* by *connecting* to an *instance* to interact with a *database*.
 <a id="glossary-cluster-owner"></a>

**Cluster owner**
:   The operating system user that owns the *data directory* and under which the `postgres` process is run. It is required that this user exist prior to creating a new *database cluster*.


     On operating systems with a `root` user, said user is not allowed to be the cluster owner.
 <a id="glossary-column"></a>

**Column**
:   An *attribute* found in a *table* or *view*.
 <a id="glossary-commit"></a>

**Commit**
:   The act of finalizing a *transaction* within the *database*, which makes it visible to other transactions and assures its *durability*.


     For more information, see [sql-commit](../reference/sql-commands/commit.md#sql-commit).
 <a id="glossary-concurrency"></a>

**Concurrency**
:   The concept that multiple independent operations happen within the *database* at the same time. In PostgreSQL, concurrency is controlled by the *multiversion concurrency control* mechanism.
 <a id="glossary-connection"></a>

**Connection**
:   An established line of communication between a client process and a *backend* process, usually over a network, supporting a *session*. This term is sometimes used as a synonym for session.


     For more information, see [Connections and Authentication](../server-administration/server-configuration/connections-and-authentication.md#runtime-config-connection).
 <a id="glossary-consistency"></a>

**Consistency**
:   The property that the data in the *database* is always in compliance with *integrity constraints*. Transactions may be allowed to violate some of the constraints transiently before it commits, but if such violations are not resolved by the time it commits, such a transaction is automatically *rolled back*. This is one of the ACID properties.
 <a id="glossary-constraint"></a>

**Constraint**
:   A restriction on the values of data allowed within a *table*, or in attributes of a *domain*.


     For more information, see [Constraints](../the-sql-language/data-definition/constraints.md#ddl-constraints).
 <a id="glossary-cumulative-statistics"></a>

**Cumulative Statistics System**
:   A system which, if enabled, accumulates statistical information about the *instance*'s activities.


     For more information, see [The Cumulative Statistics System](../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats).


**Data area**
    *See: *
 <a id="glossary-database"></a>

**Database**
:   A named collection of *local SQL objects*.


     For more information, see [Overview](../server-administration/managing-databases/overview.md#manage-ag-overview).
 <a id="glossary-db-cluster"></a>

**Database cluster**
:   A collection of databases and global SQL objects, and their common static and dynamic metadata. Sometimes referred to as a *cluster*. A database cluster is created using the [app-initdb](../reference/postgresql-server-applications/initdb.md#app-initdb) program.


     In PostgreSQL, the term *cluster* is also sometimes used to refer to an instance. (Don't confuse this term with the SQL command `CLUSTER`.)


     See also *cluster owner*, the operating-system owner of a cluster, and *bootstrap superuser*, the PostgreSQL owner of a cluster.


**Database server**
    *See: *
 <a id="glossary-database-superuser"></a>

**Database superuser**
:   A role having *superuser status* (see [Role Attributes](../server-administration/database-roles/role-attributes.md#role-attributes)).


     Frequently referred to as *superuser*.
 <a id="glossary-data-directory"></a>

**Data directory**
:   The base directory on the file system of a *server* that contains all data files and subdirectories associated with a *database cluster* (with the exception of *tablespaces*, and optionally *WAL*). The environment variable `PGDATA` is commonly used to refer to the data directory.


     A *cluster*'s storage space comprises the data directory plus any additional tablespaces.


     For more information, see [Database File Layout](../internals/database-physical-storage/database-file-layout.md#storage-file-layout).
 <a id="glossary-data-page"></a>

**Data page**
:   The basic structure used to store relation data. All pages are of the same size. Data pages are typically stored on disk, each in a specific file, and can be read to *shared buffers* where they can be modified, becoming *dirty*. They become clean when written to disk. New pages, which initially exist in memory only, are also dirty until written.
 <a id="glossary-datum"></a>

**Datum**
:   The internal representation of one value of an SQL data type.
 <a id="glossary-delete"></a>

**Delete**
:   An SQL command which removes *rows* from a given *table* or *relation*.


     For more information, see [sql-delete](../reference/sql-commands/delete.md#sql-delete).
 <a id="glossary-domain"></a>

**Domain**
:   A user-defined data type that is based on another underlying data type. It acts the same as the underlying type except for possibly restricting the set of allowed values.


     For more information, see [Domain Types](../the-sql-language/data-types/domain-types.md#domains).
 <a id="glossary-durability"></a>

**Durability**
:   The assurance that once a *transaction* has been *committed*, the changes remain even after a system failure or crash. This is one of the ACID properties.


**Epoch**
    *See: *
 <a id="glossary-extension"></a>

**Extension**
:   A software add-on package that can be installed on an *instance* to get extra features.


     For more information, see [Packaging Related Objects into an Extension](../server-programming/extending-sql/packaging-related-objects-into-an-extension.md#extend-extensions).
 <a id="glossary-file-segment"></a>

**File segment**
:   A physical file which stores data for a given *relation*. File segments are limited in size by a configuration value (typically 1 gigabyte), so if a relation exceeds that size, it is split into multiple segments.


     For more information, see [Database File Layout](../internals/database-physical-storage/database-file-layout.md#storage-file-layout).


     (Don't confuse this term with the similar term *WAL segment*).
 <a id="glossary-foreign-data-wrapper"></a>

**Foreign data wrapper**
:   A means of representing data that is not contained in the local *database* so that it appears as if were in local *table(s)*. With a foreign data wrapper it is possible to define a *foreign server* and *foreign tables*.


     For more information, see [sql-createforeigndatawrapper](../reference/sql-commands/create-foreign-data-wrapper.md#sql-createforeigndatawrapper).
 <a id="glossary-foreign-key"></a>

**Foreign key**
:   A type of *constraint* defined on one or more *columns* in a *table* which requires the value(s) in those *columns* to identify zero or one *row* in another (or, infrequently, the same) *table*.
 <a id="glossary-foreign-server"></a>

**Foreign server**
:   A named collection of *foreign tables* which all use the same *foreign data wrapper* and have other configuration values in common.


     For more information, see [sql-createserver](../reference/sql-commands/create-server.md#sql-createserver).
 <a id="glossary-foreign-table"></a>

**Foreign table (relation)**
:   A *relation* which appears to have *rows* and *columns* similar to a regular *table*, but will forward requests for data through its *foreign data wrapper*, which will return *result sets* structured according to the definition of the *foreign table*.


     For more information, see [sql-createforeigntable](../reference/sql-commands/create-foreign-table.md#sql-createforeigntable).
 <a id="glossary-fork"></a>

**Fork**
:   Each of the separate segmented file sets in which a relation is stored. The *main fork* is where the actual data resides. There also exist two secondary forks for metadata: the *free space map* and the *visibility map*. *Unlogged relations* also have an *init fork*.
 <a id="glossary-fsm"></a>

**Free space map (fork)**
:   A storage structure that keeps metadata about each data page of a table's main fork. The free space map entry for each page stores the amount of free space that's available for future tuples, and is structured to be efficiently searched for available space for a new tuple of a given size.


     For more information, see [Free Space Map](../internals/database-physical-storage/free-space-map.md#storage-fsm).
 <a id="glossary-function"></a>

**Function (routine)**
:   A type of routine that receives zero or more arguments, returns zero or more output values, and is constrained to run within one transaction. Functions are invoked as part of a query, for example via `SELECT`. Certain functions can return *sets*; those are called *set-returning functions*.


     Functions can also be used for *triggers* to invoke.


     For more information, see [sql-createfunction](../reference/sql-commands/create-function.md#sql-createfunction).


**GMT**
    *See: *
 <a id="glossary-grant"></a>

**Grant**
:   An SQL command that is used to allow a *user* or *role* to access specific objects within the *database*.


     For more information, see [sql-grant](../reference/sql-commands/grant.md#sql-grant).
 <a id="glossary-guc"></a>

**GUC**
:   Based on a short-hand for Grand Unified Configuration, the subsystem which controls server configuration, it is now a commonly used term for the individual configuration parameters themselves. Typically meant to refer to user settable parameters, it can also refer to internal or build-time parameters.


     For information on working with configuration parameters, see [Setting Parameters](../server-administration/server-configuration/setting-parameters.md#config-setting).
 <a id="glossary-heap"></a>

**Heap**
:   Contains the values of *row* attributes (i.e., the data) for a *relation*. The heap is realized within one or more *file segments* in the relation's *main fork*.
 <a id="glossary-host"></a>

**Host**
:   A computer that communicates with other computers over a network. This is sometimes used as a synonym for *server*. It is also used to refer to a computer where *client processes* run.
 <a id="glossary-index"></a>

**Index (relation)**
:   A *relation* that contains data derived from a *table* or *materialized view*. Its internal structure supports fast retrieval of and access to the original data.


     For more information, see [sql-createindex](../reference/sql-commands/create-index.md#sql-createindex).
 <a id="glossary-incremental-backup"></a>

**Incremental backup**
:   A special *base backup* that for some files may contain only those pages that were modified since a previous backup, as opposed to the full contents of every file. Like base backups, it is generated by the tool [app-pgbasebackup](../reference/postgresql-client-applications/pg_basebackup.md#app-pgbasebackup).


     To restore incremental backups the tool [app-pgcombinebackup](../reference/postgresql-client-applications/pg_combinebackup.md#app-pgcombinebackup) is used, which combines incremental backups with a base backup. Afterwards, recovery can use *WAL* to bring the *database cluster* to a consistent state.


     For more information, see [Making an Incremental Backup](../server-administration/backup-and-restore/continuous-archiving-and-point-in-time-recovery-pitr.md#backup-incremental-backup).
 <a id="glossary-io"></a>

**Input/Output**
:   Input/Output (I/O) describes the communication between a program and peripheral devices. In the context of database systems, I/O commonly, but not exclusively, refers to interaction with storage devices or the network.
         *See also: *
 <a id="glossary-insert"></a>

**Insert**
:   An SQL command used to add new data into a *table*.


     For more information, see [sql-insert](../reference/sql-commands/insert.md#sql-insert).
 <a id="glossary-instance"></a>

**Instance**
:   A group of *backend* and *auxiliary processes* that communicate using a common shared memory area. One *postmaster process* manages the instance; one instance manages exactly one *database cluster* with all its databases. Many instances can run on the same *server* as long as their TCP ports do not conflict.


     The instance handles all key features of a DBMS: read and write access to files and shared memory, assurance of the ACID properties, *connections* to *client processes*, privilege verification, crash recovery, replication, etc.
 <a id="glossary-isolation"></a>

**Isolation**
:   The property that the effects of a transaction are not visible to *concurrent transactions* before it commits. This is one of the ACID properties.


     For more information, see [Transaction Isolation](../the-sql-language/concurrency-control/transaction-isolation.md#transaction-iso).
 <a id="glossary-join"></a>

**Join**
:   An operation and SQL keyword used in *queries* for combining data from multiple *relations*.
 <a id="glossary-key"></a>

**Key**
:   A means of identifying a *row* within a *table* or other *relation* by values contained within one or more *attributes* in that relation.
 <a id="glossary-lock"></a>

**Lock**
:   A mechanism that allows a process to limit or prevent simultaneous access to a resource.
 <a id="glossary-log-file"></a>

**Log file**
:   Log files contain human-readable text lines about events. Examples include login failures, long-running queries, etc.


     For more information, see [Log File Maintenance](../server-administration/routine-database-maintenance-tasks/log-file-maintenance.md#logfile-maintenance).
 <a id="glossary-logged"></a>

**Logged**
:   A *table* is considered *logged* if changes to it are sent to the *WAL*. By default, all regular tables are logged. A table can be specified as *unlogged* either at creation time or via the `ALTER TABLE` command.
 <a id="glossary-logger"></a>

**Logger (process)**
:   An *auxiliary process* which, if enabled, writes information about database events into the current *log file*. When reaching certain time- or volume-dependent criteria, a new log file is created. Also called *syslogger*.


     For more information, see [Error Reporting and Logging](../server-administration/server-configuration/error-reporting-and-logging.md#runtime-config-logging).
 <a id="glossary-logical-replication-cluster"></a>

**Logical replication cluster**
:   A set of publisher and subscriber instances with the publisher instance replicating changes to the subscriber instance.
 <a id="glossary-log-record"></a>

**Log record**
:   Archaic term for a *WAL record*.
 <a id="glossary-log-sequence-number"></a>

**Log sequence number**
:   Byte offset into the *WAL*, increasing monotonically with each new *WAL record*.


     For more information, see [`pg_lsn`](../the-sql-language/data-types/pg_lsn-type.md#datatype-pg-lsn) and [WAL Internals](../server-administration/reliability-and-the-write-ahead-log/wal-internals.md#wal-internals).


**LSN**
    *See: *


**Master (server)**
    *See: *
 <a id="glossary-materialized"></a>

**Materialized**
:   The property that some information has been pre-computed and stored for later use, rather than computing it on-the-fly.


     This term is used in *materialized view*, to mean that the data derived from the view's query is stored on disk separately from the sources of that data.


     This term is also used to refer to some multi-step queries to mean that the data resulting from executing a given step is stored in memory (with the possibility of spilling to disk), so that it can be read multiple times by another step.
 <a id="glossary-materialized-view"></a>

**Materialized view (relation)**
:   A *relation* that is defined by a `SELECT` statement (just like a *view*), but stores data in the same way that a *table* does. It cannot be modified via `INSERT`, `UPDATE`, `DELETE`, or `MERGE` operations.


     For more information, see [sql-creatematerializedview](../reference/sql-commands/create-materialized-view.md#sql-creatematerializedview).
 <a id="glossary-merge"></a>

**Merge**
:   An SQL command used to conditionally add, modify, or remove *rows* in a given *table*, using data from a source *relation*.


     For more information, see [sql-merge](../reference/sql-commands/merge.md#sql-merge).
 <a id="glossary-mvcc"></a>

**Multi-version concurrency control (MVCC)**
:   A mechanism designed to allow several *transactions* to be reading and writing the same rows without one process causing other processes to stall. In PostgreSQL, MVCC is implemented by creating copies (*versions*) of *tuples* as they are modified; after transactions that can see the old versions terminate, those old versions need to be removed.
 <a id="glossary-null"></a>

**Null**
:   A concept of non-existence that is a central tenet of relational database theory. It represents the absence of a definite value.


**Optimizer**
    *See: *
 <a id="glossary-parallel-query"></a>

**Parallel query**
:   The ability to handle parts of executing a *query* to take advantage of parallel processes on servers with multiple CPUs.
 <a id="glossary-partition"></a>

**Partition**
:   One of several disjoint (not overlapping) subsets of a larger set.
 <a id="glossary-partitioned-table"></a>

**Partitioned table (relation)**
:   A *relation* that is in semantic terms the same as a *table*, but whose storage is distributed across several *partitions*.
 <a id="glossary-postmaster"></a>

**Postmaster (process)**
:   The very first process of an *instance*. It starts and manages the *auxiliary processes* and creates *backend processes* on demand.


     For more information, see [Starting the Database Server](../server-administration/server-setup-and-operation/starting-the-database-server.md#server-start).
 <a id="glossary-primary-key"></a>

**Primary key**
:   A special case of a *unique constraint* defined on a *table* or other *relation* that also guarantees that all of the *attributes* within the *primary key* do not have *null* values. As the name implies, there can be only one primary key per table, though it is possible to have multiple unique constraints that also have no null-capable attributes.
 <a id="glossary-primary-server"></a>

**Primary (server)**
:   When two or more *databases* are linked via *replication*, the *server* that is considered the authoritative source of information is called the *primary*, also known as a *master*.
 <a id="glossary-procedure"></a>

**Procedure (routine)**
:   A type of routine. Their distinctive qualities are that they do not return values, and that they are allowed to make transactional statements such as `COMMIT` and `ROLLBACK`. They are invoked via the `CALL` command.


     For more information, see [sql-createprocedure](../reference/sql-commands/create-procedure.md#sql-createprocedure).
 <a id="glossary-query"></a>

**Query**
:   A request sent by a client to a *backend*, usually to return results or to modify data on the database.
 <a id="glossary-planner"></a>

**Query planner**
:   The part of PostgreSQL that is devoted to determining (*planning*) the most efficient way to execute *queries*. Also known as *query optimizer*, *optimizer*, or simply *planner*.


**Record**
    *See: *


**Recycling**
    *See: *
 <a id="glossary-referential-integrity"></a>

**Referential integrity**
:   A means of restricting data in one *relation* by a *foreign key* so that it must have matching data in another *relation*.
 <a id="glossary-relation"></a>

**Relation**
:   Mathematically, a *relation* is a set of *tuples*; this is the sense meant in the term "relational database".


     In PostgreSQL, "relation" is commonly used to mean an *SQL object* that has a name and a list of attributes defined in a specific order. *Tables*, *sequences*, *views*, *foreign tables*, *materialized views*, composite types, and *indexes* are all relations. A relation in this sense is a container or a descriptor for a set of tuples.


     *Class* is an alternative but archaic term. The system catalog [`pg_class`](../internals/system-catalogs/pg_class.md#catalog-pg-class) holds an entry for each PostgreSQL relation.
 <a id="glossary-replica"></a>

**Replica (server)**
:   A *database* that is paired with a *primary* database and is maintaining a copy of some or all of the primary database's data. The foremost reasons for doing this are to allow for greater access to that data, and to maintain availability of the data in the event that the *primary* becomes unavailable.
 <a id="glossary-replication"></a>

**Replication**
:   The act of reproducing data on one *server* onto another server called a *replica*. This can take the form of *physical replication*, where all file changes from one server are copied verbatim, or *logical replication* where a defined subset of data changes are conveyed using a higher-level representation.
 <a id="glossary-restartpoint"></a>

**Restartpoint**
:   A variant of a *checkpoint* performed on a *replica*.


     For more information, see [WAL Configuration](../server-administration/reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration).
 <a id="glossary-result-set"></a>

**Result set**
:   A *relation* transmitted from a *backend process* to a *client* upon the completion of an SQL command, usually a `SELECT` but it can be an `INSERT`, `UPDATE`, `DELETE`, or `MERGE` command if the `RETURNING` clause is specified.


     The fact that a result set is a relation means that a query can be used in the definition of another query, becoming a *subquery*.
 <a id="glossary-revoke"></a>

**Revoke**
:   A command to prevent access to a named set of *database* objects for a named list of *roles*.


     For more information, see [sql-revoke](../reference/sql-commands/revoke.md#sql-revoke).
 <a id="glossary-role"></a>

**Role**
:   A collection of access privileges to the *instance*. Roles are themselves a privilege that can be granted to other roles. This is often done for convenience or to ensure completeness when multiple *users* need the same privileges.


     For more information, see [sql-createrole](../reference/sql-commands/create-role.md#sql-createrole).
 <a id="glossary-rollback"></a>

**Rollback**
:   A command to undo all of the operations performed since the beginning of a *transaction*.


     For more information, see [sql-rollback](../reference/sql-commands/rollback.md#sql-rollback).
 <a id="glossary-routine"></a>

**Routine**
:   A defined set of instructions stored in the database system that can be invoked for execution. A routine can be written in a variety of programming languages. Routines can be *functions* (including set-returning functions and *trigger functions*), *aggregate functions*, and *procedures*.


     Many routines are already defined within PostgreSQL itself, but user-defined ones can also be added.


**Row**
    *See: *
 <a id="glossary-savepoint"></a>

**Savepoint**
:   A special mark in the sequence of steps in a *transaction*. Data modifications after this point in time may be reverted to the time of the savepoint.


     For more information, see [sql-savepoint](../reference/sql-commands/savepoint.md#sql-savepoint).
 <a id="glossary-schema"></a>

**Schema**
:   A schema is a namespace for *SQL objects*, which all reside in the same *database*. Each SQL object must reside in exactly one schema.


     All system-defined SQL objects reside in schema `pg_catalog`.


**Segment**
    *See: *
 <a id="glossary-select"></a>

**Select**
:   The SQL command used to request data from a *database*. Normally, `SELECT` commands are not expected to modify the *database* in any way, but it is possible that *functions* invoked within the query could have side effects that do modify data.


     For more information, see [sql-select](../reference/sql-commands/select.md#sql-select).
 <a id="glossary-sequence"></a>

**Sequence (relation)**
:   A type of relation that is used to generate values. Typically the generated values are sequential non-repeating numbers. They are commonly used to generate surrogate *primary key* values.
  <a id="glossary-server"></a>

**Server**
:   A computer on which PostgreSQL *instances* run. The term *server* denotes real hardware, a container, or a *virtual machine*.


     This term is sometimes used to refer to an instance or to a host.
 <a id="glossary-session"></a>

**Session**
:   A state that allows a client and a backend to interact, communicating over a *connection*.
 <a id="glossary-shared-memory"></a>

**Shared memory**
:   RAM which is used by the processes common to an *instance*. It mirrors parts of *database* files, provides a transient area for *WAL records*, and stores additional common information. Note that shared memory belongs to the complete instance, not to a single database.


     The largest part of shared memory is known as *shared buffers* and is used to mirror part of data files, organized into pages. When a page is modified, it is called a dirty page until it is written back to the file system.


     For more information, see [Memory](../server-administration/server-configuration/resource-consumption.md#runtime-config-resource-memory).
 <a id="glossary-sql-object"></a>

**SQL object**
:   Any object that can be created with a `CREATE` command. Most objects are specific to one database, and are commonly known as *local objects*.


     Most local objects reside in a specific *schema* in their containing database, such as *relations* (all types), *routines* (all types), data types, etc. The names of such objects of the same type in the same schema are enforced to be unique.


     There also exist local objects that do not reside in schemas; some examples are *extensions*, *data type casts*, and *foreign data wrappers*. The names of such objects of the same type are enforced to be unique within the database.


     Other object types, such as *roles*, *tablespaces*, replication origins, subscriptions for logical replication, and databases themselves are not local SQL objects since they exist entirely outside of any specific database; they are called *global objects*. The names of such objects are enforced to be unique within the whole database cluster.


     For more information, see [Overview](../server-administration/managing-databases/overview.md#manage-ag-overview).
 <a id="glossary-sql-standard"></a>

**SQL standard**
:   A series of documents that define the SQL language.


**Standby (server)**
    *See: *
 <a id="glossary-startup-process"></a>

**Startup process**
:   An *auxiliary process* that replays WAL during crash recovery and in a *physical replica*.


     (The name is historical: the startup process was named before replication was implemented; the name refers to its task as it relates to the server startup following a crash.)
 <a id="glossary-superuser"></a>

**Superuser**
:   As used in this documentation, it is a synonym for *database superuser*.
 <a id="glossary-system-catalog"></a>

**System catalog**
:   A collection of *tables* which describe the structure of all *SQL objects* of the instance. The system catalog resides in the schema `pg_catalog`. These tables contain data in internal representation and are not typically considered useful for user examination; a number of user-friendlier *views*, also in schema `pg_catalog`, offer more convenient access to some of that information, while additional tables and views exist in schema `information_schema` (see [The Information Schema](../client-interfaces/the-information-schema/index.md#information-schema)) that expose some of the same and additional information as mandated by the *SQL standard*.


     For more information, see [Schemas](../the-sql-language/data-definition/schemas.md#ddl-schemas).
 <a id="glossary-system-time"></a>

**System time**
:   In a *temporal table*, the dimension of time that represents when the table itself was changed (as opposed to the entity the table describes). Often used for auditing, compliance, and debugging.


     For more information, see [Temporal Tables](../the-sql-language/data-definition/temporal-tables.md#ddl-temporal-tables).
 <a id="glossary-table"></a>

**Table**
:   A collection of *tuples* having a common data structure (the same number of *attributes*, in the same order, having the same name and type per position). A table is the most common form of *relation* in PostgreSQL.


     For more information, see [sql-createtable](../reference/sql-commands/create-table.md#sql-createtable).
 <a id="glossary-tablespace"></a>

**Tablespace**
:   A named location on the server file system. All *SQL objects* which require storage beyond their definition in the *system catalog* must belong to a single tablespace. Initially, a database cluster contains a single usable tablespace which is used as the default for all SQL objects, called `pg_default`.


     For more information, see [Tablespaces](../server-administration/managing-databases/tablespaces.md#manage-ag-tablespaces).
 <a id="glossary-temporal-table"></a>

**Temporal table**
:   *Tables* that track *application time* or *system time* (or both). Not to be confused with *temporary tables*.


     For more information, see [Temporal Tables](../the-sql-language/data-definition/temporal-tables.md#ddl-temporal-tables).
 <a id="glossary-temporary-table"></a>

**Temporary table**
:   *Tables* that exist either for the lifetime of a *session* or a *transaction*, as specified at the time of creation. The data in them is not visible to other sessions, and is not *logged*. Temporary tables are often used to store intermediate data for a multi-step operation.


     For more information, see [sql-createtable](../reference/sql-commands/create-table.md#sql-createtable).
 <a id="glossary-toast"></a>

**TOAST**
:   A mechanism by which large attributes of table rows are split and stored in a secondary table, called the *TOAST table*. Each relation with large attributes has its own TOAST table.


     For more information, see [TOAST](../internals/database-physical-storage/toast.md#storage-toast).
 <a id="glossary-transaction"></a>

**Transaction**
:   A combination of commands that must act as a single *atomic* command: they all succeed or all fail as a single unit, and their effects are not visible to other *sessions* until the transaction is complete, and possibly even later, depending on the isolation level.


     For more information, see [Transaction Isolation](../the-sql-language/concurrency-control/transaction-isolation.md#transaction-iso).
 <a id="glossary-xid"></a>

**Transaction ID**
:   The numerical, unique, sequentially-assigned identifier that each transaction receives when it first causes a database modification. Frequently abbreviated as *xid*. When stored on disk, xids are only 32-bits wide, so only approximately four billion write transaction IDs can be generated; to permit the system to run for longer than that, *epochs* are used, also 32 bits wide. When the counter reaches the maximum xid value, it starts over at `3` (values under that are reserved) and the epoch value is incremented by one. In some contexts, the epoch and xid values are considered together as a single 64-bit value; see [Transactions and Identifiers](../internals/transaction-processing/transactions-and-identifiers.md#transaction-id) for more details.


     For more information, see [Object Identifier Types](../the-sql-language/data-types/object-identifier-types.md#datatype-oid).
 <a id="glossary-tps"></a>

**Transactions per second (TPS)**
:   Average number of transactions that are executed per second, totaled across all sessions active for a measured run. This is used as a measure of the performance characteristics of an instance.
 <a id="glossary-trigger"></a>

**Trigger**
:   A *function* which can be defined to execute whenever a certain operation (`INSERT`, `UPDATE`, `DELETE`, `TRUNCATE`) is applied to a *relation*. A trigger executes within the same *transaction* as the statement which invoked it, and if the function fails, then the invoking statement also fails.


     For more information, see [sql-createtrigger](../reference/sql-commands/create-trigger.md#sql-createtrigger).
 <a id="glossary-tuple"></a>

**Tuple**
:   A collection of *attributes* in a fixed order. That order may be defined by the *table* (or other *relation*) where the tuple is contained, in which case the tuple is often called a *row*. It may also be defined by the structure of a result set, in which case it is sometimes called a *record*.
 <a id="glossary-unique-constraint"></a>

**Unique constraint**
:   A type of *constraint* defined on a *relation* which restricts the values allowed in one or a combination of columns so that each value or combination of values can only appear once in the relation — that is, no other row in the relation contains values that are equal to those.


     Because *null values* are not considered equal to each other, multiple rows with null values are allowed to exist without violating the unique constraint.
 <a id="glossary-unlogged"></a>

**Unlogged**
:   The property of certain *relations* that the changes to them are not reflected in the *WAL*. This disables replication and crash recovery for these relations.


     The primary use of unlogged tables is for storing transient work data that must be shared across processes.


     *Temporary tables* are always unlogged.
 <a id="glossary-update"></a>

**Update**
:   An SQL command used to modify *rows* that may already exist in a specified *table*. It cannot create or remove rows.


     For more information, see [sql-update](../reference/sql-commands/update.md#sql-update).
 <a id="glossary-user"></a>

**User**
:   A *role* that has the *login privilege* (see [Role Attributes](../server-administration/database-roles/role-attributes.md#role-attributes)).
 <a id="glossary-user-mapping"></a>

**User mapping**
:   The translation of login credentials in the local *database* to credentials in a remote data system defined by a *foreign data wrapper*.


     For more information, see [sql-createusermapping](../reference/sql-commands/create-user-mapping.md#sql-createusermapping).
 <a id="glossary-utc"></a>

**UTC**
:   Universal Coordinated Time, the primary global time reference, approximately the time prevailing at the zero meridian of longitude. Often but inaccurately referred to as GMT (Greenwich Mean Time).
 <a id="glossary-vacuum"></a>

**Vacuum**
:   The process of removing outdated *tuple versions* from tables or materialized views, and other closely related processing required by PostgreSQL's implementation of *MVCC*. This can be initiated through the use of the `VACUUM` command, but can also be handled automatically via *autovacuum* processes.


     For more information, see [Routine Vacuuming](../server-administration/routine-database-maintenance-tasks/routine-vacuuming.md#routine-vacuuming) .
 <a id="glossary-view"></a>

**View**
:   A *relation* that is defined by a `SELECT` statement, but has no storage of its own. Any time a query references a view, the definition of the view is substituted into the query as if the user had typed it as a subquery instead of the name of the view.


     For more information, see [sql-createview](../reference/sql-commands/create-view.md#sql-createview).
 <a id="glossary-vm"></a>

**Visibility map (fork)**
:   A storage structure that keeps metadata about each data page of a table's main fork. The visibility map entry for each page stores two bits: the first one (`all-visible`) indicates that all tuples in the page are visible to all transactions. The second one (`all-frozen`) indicates that all tuples in the page are marked frozen.


**WAL**
    *See: *
 <a id="glossary-wal-archiver"></a>

**WAL archiver (process)**
:   An *auxiliary process* which, if enabled, saves copies of *WAL files* for the purpose of creating backups or keeping *replicas* current.


     For more information, see [Continuous Archiving and Point-in-Time Recovery (PITR)](../server-administration/backup-and-restore/continuous-archiving-and-point-in-time-recovery-pitr.md#continuous-archiving).
 <a id="glossary-wal-file"></a>

**WAL file**
:   Also known as *WAL segment* or *WAL segment file*. Each of the sequentially-numbered files that provide storage space for *WAL*. The files are all of the same predefined size and are written in sequential order, interspersing changes as they occur in multiple simultaneous sessions. If the system crashes, the files are read in order, and each of the changes is replayed to restore the system to the state it was in before the crash.


     Each WAL file can be released after a *checkpoint* writes all the changes in it to the corresponding data files. Releasing the file can be done either by deleting it, or by changing its name so that it will be used in the future, which is called *recycling*.


     For more information, see [WAL Internals](../server-administration/reliability-and-the-write-ahead-log/wal-internals.md#wal-internals).
 <a id="glossary-wal-record"></a>

**WAL record**
:   A low-level description of an individual data change. It contains sufficient information for the data change to be re-executed (*replayed*) in case a system failure causes the change to be lost. WAL records use a non-printable binary format.


     For more information, see [WAL Internals](../server-administration/reliability-and-the-write-ahead-log/wal-internals.md#wal-internals).
 <a id="glossary-wal-receiver"></a>

**WAL receiver (process)**
:   An *auxiliary process* that runs on a *replica* to receive WAL from the *primary server* for replay by the *startup process*.


     For more information, see [Log-Shipping Standby Servers](../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#warm-standby).


**WAL segment**
    *See: *
 <a id="glossary-wal-sender"></a>

**WAL sender (process)**
:   A special *backend process* that streams WAL over a network. The receiving end can be a *WAL receiver* in a *replica*, [app-pgreceivewal](../reference/postgresql-client-applications/pg_receivewal.md#app-pgreceivewal), or any other client program that speaks the replication protocol.
 <a id="glossary-wal-summarizer"></a>

**WAL summarizer (process)**
:   An *auxiliary process* that summarizes WAL data for *incremental backups*.


     For more information, see [WAL Summarization](../server-administration/server-configuration/write-ahead-log.md#runtime-config-wal-summarization).
 <a id="glossary-wal-writer"></a>

**WAL writer (process)**
:   An *auxiliary process* that writes *WAL records* from *shared memory* to *WAL files*.


     For more information, see [Write Ahead Log](../server-administration/server-configuration/write-ahead-log.md#runtime-config-wal).
 <a id="glossary-window-function"></a>

**Window function (routine)**
:   A type of *function* used in a *query* that applies to a *partition* of the query's *result set*; the function's result is based on values found in *rows* of the same partition or frame.


     All *aggregate functions* can be used as window functions, but window functions can also be used to, for example, give ranks to each of the rows in the partition. Also known as *analytic functions*.


     For more information, see [Window Functions](../tutorial/advanced-features/window-functions.md#tutorial-window).
 <a id="glossary-wal"></a>

**Write-ahead log**
:   The journal that keeps track of the changes in the *database cluster* as user- and system-invoked operations take place. It comprises many individual *WAL records* written sequentially to *WAL files*.
