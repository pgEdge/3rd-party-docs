<a id="runtime-config-short"></a>

## Short Options


 For convenience there are also single letter command-line option switches available for some parameters. They are described in [Short Option Key](#runtime-config-short-table). Some of these options exist for historical reasons, and their presence as a single-letter option does not necessarily indicate an endorsement to use the option heavily.
 <a id="runtime-config-short-table"></a>

**Table: Short Option Key**

| Short Option | Equivalent |
| --- | --- |
| <code>-B </code><em>x</em> | <code>shared_buffers = </code><em>x</em> |
| <code>-d </code><em>x</em> | <code>log_min_messages = DEBUG</code><em>x</em> |
| `-e` | `datestyle = euro` |
| `-fb`, `-fh`, `-fi`, `-fm`, `-fn`, `-fo`, `-fs`, `-ft` | `enable_bitmapscan = off`, `enable_hashjoin = off`, `enable_indexscan = off`, `enable_mergejoin = off`, `enable_nestloop = off`, `enable_indexonlyscan = off`, `enable_seqscan = off`, `enable_tidscan = off` |
| `-F` | `fsync = off` |
| <code>-h </code><em>x</em> | <code>listen_addresses = </code><em>x</em> |
| `-i` | `listen_addresses = '*'` |
| <code>-k </code><em>x</em> | <code>unix_socket_directories = </code><em>x</em> |
| `-l` | `ssl = on` |
| <code>-N </code><em>x</em> | <code>max_connections = </code><em>x</em> |
| `-O` | `allow_system_table_mods = on` |
| <code>-p </code><em>x</em> | <code>port = </code><em>x</em> |
| `-P` | `ignore_system_indexes = on` |
| `-s` | `log_statement_stats = on` |
| <code>-S </code><em>x</em> | <code>work_mem = </code><em>x</em> |
| `-tpa`, `-tpl`, `-te` | `log_parser_stats = on`, `log_planner_stats = on`, `log_executor_stats = on` |
| <code>-W </code><em>x</em> | <code>post_auth_delay = </code><em>x</em> |
