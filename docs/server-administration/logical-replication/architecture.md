<a id="logical-replication-architecture"></a>

## Architecture


 Logical replication is built with an architecture similar to physical streaming replication (see [Streaming Replication](../high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication)). It is implemented by `walsender` and `apply` processes. The walsender process starts logical decoding (described in [Logical Decoding](../../server-programming/logical-decoding/index.md#logicaldecoding)) of the WAL and loads the standard logical decoding output plugin ([pgoutput](../../server-programming/logical-decoding/logical-decoding-output-plugins.md#logicaldecoding-pgoutput)). The plugin transforms the changes read from WAL to the logical replication protocol (see [Logical Streaming Replication Protocol](../../internals/frontend-backend-protocol/logical-streaming-replication-protocol.md#protocol-logical-replication)) and filters the data according to the publication specification. The data is then continuously transferred using the streaming replication protocol to the apply worker, which maps the data to local tables and applies the individual changes as they are received, in correct transactional order.


 The apply process on the subscriber database always runs with [`session_replication_role`](../server-configuration/client-connection-defaults.md#guc-session-replication-role) set to `replica`. This means that, by default, triggers and rules will not fire on a subscriber. Users can optionally choose to enable triggers and rules on a table using the [`ALTER TABLE`](../../reference/sql-commands/alter-table.md#sql-altertable) command and the `ENABLE TRIGGER` and `ENABLE RULE` clauses.


 The logical replication apply process currently only fires row triggers, not statement triggers. The initial table synchronization, however, is implemented like a `COPY` command and thus fires both row and statement triggers for `INSERT`.
 <a id="logical-replication-snapshot"></a>

### Initial Snapshot


 The initial data in existing subscribed tables is snapshotted and copied in parallel instances of a special kind of apply process. These special apply processes are dedicated table synchronization workers, spawned for each table to be synchronized. Each table synchronization process will create its own replication slot and copy the existing data. As soon as the copy is finished the table contents will become visible to other backends. Once existing data is copied, the worker enters synchronization mode, which ensures that the table is brought up to a synchronized state with the main apply process by streaming any changes that happened during the initial data copy using standard logical replication. During this synchronization phase, the changes are applied and committed in the same order as they happened on the publisher. Once synchronization is done, control of the replication of the table is given back to the main apply process where replication continues as normal.


!!! note

    The publication [`publish`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-with-publish) parameter only affects what DML operations will be replicated. The initial data synchronization does not take this parameter into account when copying the existing table data.


!!! note

    If a table synchronization worker fails during copy, the apply worker detects the failure and respawns the table synchronization worker to continue the synchronization process. This behaviour ensures that transient errors do not permanently disrupt the replication setup. See also [`wal_retrieve_retry_interval`](../server-configuration/replication.md#guc-wal-retrieve-retry-interval).
