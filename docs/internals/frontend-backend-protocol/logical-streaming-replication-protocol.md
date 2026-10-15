<a id="protocol-logical-replication"></a>

## Logical Streaming Replication Protocol


 This section describes the logical replication protocol, which is the message flow started by the `START_REPLICATION` `SLOT` *slot_name* `LOGICAL` replication command.


 The logical streaming replication protocol builds on the primitives of the physical streaming replication protocol.
 <a id="protocol-logical-replication-params"></a>

### Logical Streaming Replication Parameters


 The `START_REPLICATION` command can pass options to the logical decoding output plugin associated with the specified replication slot. See [Options](../../server-programming/logical-decoding/logical-decoding-output-plugins.md#logicaldecoding-pgoutput-options) for options that are accepted by the standard (`pgoutput`) plugin.
  <a id="protocol-logical-messages"></a>

### Logical Replication Protocol Messages


 The individual protocol messages are discussed in the following subsections. Individual messages are described in [Logical Replication Message Formats](logical-replication-message-formats.md#protocol-logicalrep-message-formats).


 All top-level protocol messages begin with a message type byte. While represented in code as a character, this is a signed byte with no associated encoding.


 Since the streaming replication protocol supplies a message length there is no need for top-level protocol messages to embed a length in their header.
  <a id="protocol-logical-messages-flow"></a>

### Logical Replication Protocol Message Flow


 With the exception of the `START_REPLICATION` command and the replay progress messages, all information flows only from the backend to the frontend.


 The logical replication protocol sends individual transactions one by one. This means that all messages between a pair of Begin and Commit messages belong to the same transaction. Similarly, all messages between a pair of Begin Prepare and Prepare messages belong to the same transaction. It also sends changes of large in-progress transactions between a pair of Stream Start and Stream Stop messages. The last stream of such a transaction contains a Stream Commit or Stream Abort message.


 Every sent transaction contains zero or more DML messages (Insert, Update, Delete). In case of a cascaded setup it can also contain Origin messages. The origin message indicates that the transaction originated on different replication node. Since a replication node in the scope of logical replication protocol can be pretty much anything, the only identifier is the origin name. It's downstream's responsibility to handle this as needed (if needed). The Origin message is always sent before any DML messages in the transaction.


 Every DML message contains a relation OID, identifying the publisher's relation that was acted on. Before the first DML message for a given relation OID, a Relation message will be sent, describing the schema of that relation. Subsequently, a new Relation message will be sent if the relation's definition has changed since the last Relation message was sent for it. (The protocol assumes that the client is capable of remembering this metadata for as many relations as needed.)


 Relation messages identify column types by their OIDs. In the case of a built-in type, it is assumed that the client can look up that type OID locally, so no additional data is needed. For a non-built-in type OID, a Type message will be sent before the Relation message, to provide the type name associated with that OID. Thus, a client that needs to specifically identify the types of relation columns should cache the contents of Type messages, and first consult that cache to see if the type OID is defined there. If not, look up the type OID locally.
