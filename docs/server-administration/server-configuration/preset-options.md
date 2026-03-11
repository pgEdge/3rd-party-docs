## Preset Options { #runtime-config-preset }


 The following “parameters” are read-only. As such, they have been excluded from the sample `postgresql.conf` file. These options report various aspects of PostgreSQL behavior that might be of interest to certain applications, particularly administrative front-ends. Most of them are determined when PostgreSQL is compiled or when it is installed.


<a id="guc-block-size"></a>

`block_size` (`integer`)
:   Reports the size of a disk block. It is determined by the value of `BLCKSZ` when building the server. The default value is 8192 bytes. The meaning of some configuration variables (such as [shared_buffers](resource-consumption.md#guc-shared-buffers)) is influenced by `block_size`. See [Resource Consumption](resource-consumption.md#runtime-config-resource) for information.
<a id="guc-data-checksums"></a>

`data_checksums` (`boolean`)
:   Reports whether data checksums are enabled for this cluster. See [-k](../../reference/postgresql-server-applications/initdb.md#app-initdb-data-checksums) for more information.
<a id="guc-data-directory-mode"></a>

`data_directory_mode` (`integer`)
:   On Unix systems this parameter reports the permissions the data directory (defined by [data_directory](file-locations.md#guc-data-directory)) had at server startup. (On Microsoft Windows this parameter will always display `0700`.) See [the initdb `-g` option](../../reference/postgresql-server-applications/initdb.md#app-initdb-allow-group-access) for more information.
<a id="guc-debug-assertions"></a>

`debug_assertions` (`boolean`)
:   Reports whether PostgreSQL has been built with assertions enabled. That is the case if the macro `USE_ASSERT_CHECKING` is defined when PostgreSQL is built (accomplished e.g., by the `configure` option `--enable-cassert`). By default PostgreSQL is built without assertions.
<a id="guc-debug-exec-backend"></a>

`debug_exec_backend` (`boolean`)
:   Reports whether PostgreSQL has been built with `EXEC_BACKEND` enabled. That is the case on `Windows` or if the macro `EXEC_BACKEND` is defined when PostgreSQL is built.
<a id="guc-effective-wal-level"></a>

`effective_wal_level` (`enum`)
:   Reports the actual WAL logging level currently in effect in the system. This parameter shares the same set of values as [wal_level](write-ahead-log.md#guc-wal-level), but reflects the operational WAL level rather than the configured setting. For descriptions of possible values, refer to the `wal_level` parameter documentation.


     The effective WAL level can differ from the configured `wal_level` in certain situations. For example, when `wal_level` is set to `replica` and the system has one or more logical replication slots, `effective_wal_level` will show `logical` to indicate that the system is maintaining WAL records at `logical` level equivalent.


     On standby servers, `effective_wal_level` matches the value of `effective_wal_level` from the most upstream server in the replication chain.
<a id="guc-huge-pages-status"></a>

`huge_pages_status` (`enum`)
:   Reports the state of huge pages in the current instance: `on`, `off`, or `unknown` (if displayed with `postgres -C`). This parameter is useful to determine whether allocation of huge pages was successful under `huge_pages=try`. See [huge_pages](resource-consumption.md#guc-huge-pages) for more information.
<a id="guc-integer-datetimes"></a>

`integer_datetimes` (`boolean`)
:   Reports whether PostgreSQL was built with support for 64-bit-integer dates and times. As of PostgreSQL 10, this is always `on`.
<a id="guc-in-hot-standby"></a>

`in_hot_standby` (`boolean`)
:   Reports whether the server is currently in hot standby mode. When this is `on`, all transactions are forced to be read-only. Within a session, this can change only if the server is promoted to be primary. See [Hot Standby](../high-availability-load-balancing-and-replication/hot-standby.md#hot-standby) for more information.
<a id="guc-max-function-args"></a>

`max_function_args` (`integer`)
:   Reports the maximum number of function arguments. It is determined by the value of `FUNC_MAX_ARGS` when building the server. The default value is 100 arguments.
<a id="guc-max-identifier-length"></a>

`max_identifier_length` (`integer`)
:   Reports the maximum identifier length. It is determined as one less than the value of `NAMEDATALEN` when building the server. The default value of `NAMEDATALEN` is 64; therefore the default `max_identifier_length` is 63 bytes, which can be less than 63 characters when using multibyte encodings.
<a id="guc-max-index-keys"></a>

`max_index_keys` (`integer`)
:   Reports the maximum number of index keys. It is determined by the value of `INDEX_MAX_KEYS` when building the server. The default value is 32 keys.
<a id="guc-num-os-semaphores"></a>

`num_os_semaphores` (`integer`)
:   Reports the number of semaphores that are needed for the server based on the configured number of allowed connections ([max_connections](connections-and-authentication.md#guc-max-connections)), allowed autovacuum worker processes ([autovacuum_max_workers](vacuuming.md#guc-autovacuum-max-workers)), allowed WAL sender processes ([max_wal_senders](replication.md#guc-max-wal-senders)), allowed background processes ([max_worker_processes](resource-consumption.md#guc-max-worker-processes)), etc.
<a id="guc-segment-size"></a>

`segment_size` (`integer`)
:   Reports the number of blocks (pages) that can be stored within a file segment. It is determined by the value of `RELSEG_SIZE` when building the server. The maximum size of a segment file in bytes is equal to `segment_size` multiplied by `block_size`; by default this is 1GB.
<a id="guc-server-encoding"></a>

`server_encoding` (`string`)
:   Reports the database encoding (character set). It is determined when the database is created. Ordinarily, clients need only be concerned with the value of [client_encoding](client-connection-defaults.md#guc-client-encoding).
<a id="guc-server-version"></a>

`server_version` (`string`)
:   Reports the version number of the server. It is determined by the value of `PG_VERSION` when building the server.
<a id="guc-server-version-num"></a>

`server_version_num` (`integer`)
:   Reports the version number of the server as an integer. It is determined by the value of `PG_VERSION_NUM` when building the server.
<a id="guc-shared-memory-size"></a>

`shared_memory_size` (`integer`)
:   Reports the size of the main shared memory area, rounded up to the nearest megabyte.
<a id="guc-shared-memory-size-in-huge-pages"></a>

`shared_memory_size_in_huge_pages` (`integer`)
:   Reports the number of huge pages that are needed for the main shared memory area based on the specified [huge_page_size](resource-consumption.md#guc-huge-page-size). If huge pages are not supported, this will be `-1`.


     This setting is supported only on Linux. It is always set to `-1` on other platforms. For more details about using huge pages on Linux, see [Linux Huge Pages](../server-setup-and-operation/managing-kernel-resources.md#linux-huge-pages).
<a id="guc-ssl-library"></a>

`ssl_library` (`string`)
:   Reports the name of the SSL library that this PostgreSQL server was built with (even if SSL is not currently configured or in use on this instance), for example `OpenSSL`, or an empty string if none.
<a id="guc-wal-block-size"></a>

`wal_block_size` (`integer`)
:   Reports the size of a WAL disk block. It is determined by the value of `XLOG_BLCKSZ` when building the server. The default value is 8192 bytes.
<a id="guc-wal-segment-size"></a>

`wal_segment_size` (`integer`)
:   Reports the size of write ahead log segments. The default value is 16MB. See [WAL Configuration](../reliability-and-the-write-ahead-log/wal-configuration.md#wal-configuration) for more information.
