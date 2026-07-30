<a id="logicaldecoding-synchronous"></a>

## Synchronous Replication Support for Logical Decoding
  <a id="logicaldecoding-synchronous-overview"></a>

### Overview


 Logical decoding can be used to build [synchronous replication](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#synchronous-replication) solutions with the same user interface as synchronous replication for [streaming replication](../../server-administration/high-availability-load-balancing-and-replication/log-shipping-standby-servers.md#streaming-replication). To do this, the streaming replication interface (see [Streaming Replication Protocol Interface](streaming-replication-protocol-interface.md#logicaldecoding-walsender)) must be used to stream out data. Clients have to send `Standby status update (F)` (see [Streaming Replication Protocol](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication)) messages, just like streaming replication clients do.


!!! note

    A synchronous replica receiving changes via logical decoding will work in the scope of a single database. Since, in contrast to that, `synchronous_standby_names` currently is server wide, this means this technique will not work properly if more than one database is actively used.
  <a id="logicaldecoding-synchronous-caveats"></a>

### Caveats


 In synchronous replication setup, a deadlock can happen, if the transaction has locked [user] catalog tables exclusively. See [Capabilities](writing-logical-decoding-output-plugins.md#logicaldecoding-capabilities) for information on user catalog tables. This is because logical decoding of transactions can lock catalog tables to access them. To avoid this users must refrain from taking an exclusive lock on [user] catalog tables. This can happen in the following ways:

-  Issuing an explicit `LOCK` on `pg_class` in a transaction.
-  Perform `CLUSTER` on `pg_class` in a transaction.
-  `PREPARE TRANSACTION` after `LOCK` command on `pg_class` and allow logical decoding of two-phase transactions.
-  `PREPARE TRANSACTION` after `CLUSTER` command on `pg_trigger` and allow logical decoding of two-phase transactions. This will lead to deadlock only when published table have a trigger.
-  Executing `TRUNCATE` on [user] catalog table in a transaction.
 Note that these commands can cause deadlocks not only for the system catalog tables listed above but for other catalog tables.
