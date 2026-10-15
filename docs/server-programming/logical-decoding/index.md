<a id="logicaldecoding"></a>

# Logical Decoding

 PostgreSQL provides infrastructure to stream the modifications performed via SQL to external consumers. This functionality can be used for a variety of purposes, including replication solutions and auditing.

 Changes are sent out in streams identified by logical replication slots.

 The format in which those changes are streamed is determined by the output plugin used. An example plugin is provided in the PostgreSQL distribution. Additional plugins can be written to extend the choice of available formats without modifying any core code. Every output plugin has access to each individual new row produced by `INSERT` and the new row version created by `UPDATE`. Availability of old row versions for `UPDATE` and `DELETE` depends on the configured replica identity (see [REPLICA IDENTITY](../../reference/sql-commands/alter-table.md#sql-altertable-replica-identity)).

 Changes can be consumed either using the streaming replication protocol (see [Streaming Replication Protocol](../../internals/frontend-backend-protocol/streaming-replication-protocol.md#protocol-replication) and [Streaming Replication Protocol Interface](streaming-replication-protocol-interface.md#logicaldecoding-walsender)), or by calling functions via SQL (see [Logical Decoding SQL Interface](logical-decoding-sql-interface.md#logicaldecoding-sql)). It is also possible to write additional methods of consuming the output of a replication slot without modifying core code (see [Logical Decoding Output Writers](logical-decoding-output-writers.md#logicaldecoding-writer)).

- [Logical Decoding Examples](logical-decoding-examples.md#logicaldecoding-example)
- [Logical Decoding Concepts](logical-decoding-concepts.md#logicaldecoding-explanation)
- [Streaming Replication Protocol Interface](streaming-replication-protocol-interface.md#logicaldecoding-walsender)
- [Logical Decoding SQL Interface](logical-decoding-sql-interface.md#logicaldecoding-sql)
- [System Catalogs Related to Logical Decoding](system-catalogs-related-to-logical-decoding.md#logicaldecoding-catalogs)
- [Logical Decoding Output Plugins](logical-decoding-output-plugins.md#logicaldecoding-output-plugin)
- [Writing Logical Decoding Output Plugins](writing-logical-decoding-output-plugins.md#logicaldecoding-output-plugin-writing)
- [Logical Decoding Output Writers](logical-decoding-output-writers.md#logicaldecoding-writer)
- [Synchronous Replication Support for Logical Decoding](synchronous-replication-support-for-logical-decoding.md#logicaldecoding-synchronous)
- [Streaming of Large Transactions for Logical Decoding](streaming-of-large-transactions-for-logical-decoding.md#logicaldecoding-streaming)
- [Two-phase Commit Support for Logical Decoding](two-phase-commit-support-for-logical-decoding.md#logicaldecoding-two-phase-commits)
