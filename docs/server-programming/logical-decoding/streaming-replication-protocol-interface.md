<a id="logicaldecoding-walsender"></a>

## Streaming Replication Protocol Interface


 The commands

- <code>CREATE_REPLICATION_SLOT </code><em>slot_name</em><code> LOGICAL </code><em>output_plugin</em>
- <code>DROP_REPLICATION_SLOT </code><em>slot_name</em> [ `WAIT` ]
- <code>START_REPLICATION SLOT </code><em>slot_name</em><code> LOGICAL ...</code>
 are used to create, drop, and stream changes from a replication slot, respectively. These commands are only available over a replication connection; they cannot be used via SQL. See [Streaming Replication Protocol](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication) for details on these commands.


 The command [app-pgrecvlogical](../../reference/postgresql-client-applications/pg_recvlogical.md#app-pgrecvlogical) can be used to control logical decoding over a streaming replication connection. (It uses these commands internally.)
