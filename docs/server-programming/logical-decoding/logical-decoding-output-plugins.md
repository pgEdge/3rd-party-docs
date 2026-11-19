<a id="logicaldecoding-output-plugin"></a>

## Logical Decoding Output Plugins


 PostgreSQL provides two logical decoding output plugins, [pgoutput](#logicaldecoding-pgoutput) and [test_decoding](../../appendixes/additional-supplied-modules-and-extensions/test_decoding-sql-based-test-example-module-for-wal-logical-decoding.md#test-decoding). You can also develop custom output plugins (see [Writing Logical Decoding Output Plugins](writing-logical-decoding-output-plugins.md#logicaldecoding-output-plugin-writing) for details).
 <a id="logicaldecoding-pgoutput"></a>

### pgoutput — Standard Logical Decoding Output Plugin


 `pgoutput` is the standard logical decoding output plugin provided by PostgreSQL. It's used for the built-in [logical replication](../../server-administration/logical-replication/index.md#logical-replication).
 <a id="logicaldecoding-pgoutput-options"></a>

#### Options


<a id="pgoutput-options-proto-version"></a>

`proto_version` (`integer`)
:   Specifies the protocol version. Currently versions `1`, `2`, `3`, and `4` are supported. A valid version is required.


     Version `2` is supported on server version 14 and above, and is required when `streaming` is set to `on` to stream large in-progress transactions.


     Version `3` is supported on server version 15 and above, and is required when `two_phase` is enabled to stream two-phase commits.


     Version `4` is supported on server version 16 and above, and is required when `streaming` is set to `parallel` to stream large in-progress transactions to be applied in parallel.
<a id="pgoutput-options-publication-names"></a>

`publication_names` (`string`)
:   A comma-separated list of publication names to subscribe to. The individual publication names are treated as standard objects names and can be quoted the same as needed. At least one publication name is required.
<a id="pgoutput-options-binary"></a>

`binary` (`boolean`)
:   Enables binary transfer mode. Binary mode is faster than the text mode but slightly less robust. The default is `off`.
<a id="pgoutput-options-messages"></a>

`messages` (`boolean`)
:   Enables sending the messages that are written by `[pg_logical_emit_message](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-emit-message)`. The default is `off`.
<a id="pgoutput-options-streaming"></a>

`streaming` (`enum`)
:   Enables streaming of in-progress transactions. Valid values are `off` (the default), `on` and `parallel`.


     When set to `off`, `pgoutput` fully decodes a transaction before sending it as a whole. This mode works with any protocol version.


     When set to `on`, `pgoutput` streams large in-progress transactions. This requires protocol version 2 or higher.


     When set to `parallel`, `pgoutput` streams large in-progress transactions and also sends extra information in some messages to support parallel processing. This requires protocol version 4 or higher.
<a id="pgoutput-options-two-phase"></a>

`two_phase` (`boolean`)
:   Enables sending two-phase transactions. Minimum protocol version 3 is required to turn it on. The default is `off`.
<a id="pgoutput-options-origin"></a>

`origin` (`enum`)
:   Specifies whether to send changes by their origin. Possible values are `none` to only send the changes that have no origin associated, or `any` to send the changes regardless of their origin. This can be used to avoid loops (infinite replication of the same data) among replication nodes. The default is `any`.
  <a id="logicaldecoding-pgoutput-notes"></a>

#### Notes


 `pgoutput` produces binary output, so functions expecting textual data ( `[pg_logical_slot_peek_changes](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-slot-peek-changes)` and `[pg_logical_slot_get_changes](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-slot-get-changes)`) cannot be used with it. Use `[pg_logical_slot_peek_binary_changes](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-slot-peek-binary-changes)` or `[pg_logical_slot_get_binary_changes](../../the-sql-language/functions-and-operators/system-administration-functions.md#pg-logical-slot-get-binary-changes)` instead.
