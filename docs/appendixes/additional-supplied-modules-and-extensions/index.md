<a id="contrib"></a>

# Additional Supplied Modules and Extensions

 This appendix and the next one contain information on the optional components found in the `contrib` directory of the PostgreSQL distribution. These include porting tools, analysis utilities, and plug-in features that are not part of the core PostgreSQL system. They are separate mainly because they address a limited audience or are too experimental to be part of the main source tree. This does not preclude their usefulness.

 This appendix covers extensions and other server plug-in module libraries found in `contrib`. [Additional Supplied Programs](../additional-supplied-programs/index.md#contrib-prog) covers utility programs.

 When building from the source distribution, these optional components are not built automatically, unless you build the "world" target (see [Build](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#build)). You can build and install all of them by running:

```

make
make install
```
 in the `contrib` directory of a configured source tree; or to build and install just one selected module, do the same in that module's subdirectory. Many of the modules have regression tests, which can be executed by running:

```

make check
```
 before installation or

```

make installcheck
```
 once you have a PostgreSQL server running.

 If you are using a pre-packaged version of PostgreSQL, these components are typically made available as a separate subpackage, such as `postgresql-contrib`.

 Many components supply new user-defined functions, operators, or types, packaged as *extensions*. To make use of one of these extensions, after you have installed the code you need to register the new SQL objects in the database system. This is done by executing a [sql-createextension](../../reference/sql-commands/create-extension.md#sql-createextension) command. In a fresh database, you can simply do

```sql

CREATE EXTENSION EXTENSION_NAME;
```
 This command registers the new SQL objects in the current database only, so you need to run it in every database in which you want the extension's facilities to be available. Alternatively, run it in database `template1` so that the extension will be copied into subsequently-created databases by default.

 For all extensions, the `CREATE EXTENSION` command must be run by a database superuser, unless the extension is considered “trusted”. Trusted extensions can be run by any user who has `CREATE` privilege on the current database. Extensions that are trusted are identified as such in the sections that follow. Generally, trusted extensions are ones that cannot provide access to outside-the-database functionality.
<a id="contrib-trusted-extensions"></a>

 The following extensions are trusted in a default installation:

- [btree_gin](btree_gin-gin-operator-classes-with-b-tree-behavior.md#btree-gin)
- [btree_gist](btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist)
- [citext](citext-a-case-insensitive-character-string-type.md#citext)
- [cube](cube-a-multi-dimensional-cube-data-type.md#cube)
- [dict_int](dict_int-example-full-text-search-dictionary-for-integers.md#dict-int)
- [fuzzystrmatch](fuzzystrmatch-determine-string-similarities-and-distance.md#fuzzystrmatch)
- [hstore](hstore-hstore-key-value-datatype.md#hstore)
- [intarray](intarray-manipulate-arrays-of-integers.md#intarray)
- [isn](isn-data-types-for-international-standard-numbers-isbn-ean-upc-etc.md#isn)
- [lo](lo-manage-large-objects.md#lo)
- [ltree](ltree-hierarchical-tree-like-data-type.md#ltree)
- [pgcrypto](pgcrypto-cryptographic-functions.md#pgcrypto)
- [pg_trgm](pg_trgm-support-for-similarity-of-text-using-trigram-matching.md#pgtrgm)
- [seg](seg-a-datatype-for-line-segments-or-floating-point-intervals.md#seg)
- [tablefunc](tablefunc-functions-that-return-tables-crosstab-and-others.md#tablefunc)
- [tcn](tcn-a-trigger-function-to-notify-listeners-of-changes-to-table-content.md#tcn)
- [tsm_system_rows](tsm_system_rows-the-system_rows-sampling-method-for-tablesample.md#tsm-system-rows)
- [tsm_system_time](tsm_system_time-the-system_time-sampling-method-for-tablesample.md#tsm-system-time)
- [unaccent](unaccent-a-text-search-dictionary-which-removes-diacritics.md#unaccent)
- [uuid-ossp](uuid-ossp-a-uuid-generator.md#uuid-ossp)


 Many extensions allow you to install their objects in a schema of your choice. To do that, add <code>SCHEMA
  </code><em>schema_name</em> to the `CREATE EXTENSION` command. By default, the objects will be placed in your current creation target schema, which in turn defaults to `public`.

 Note, however, that some of these components are not “extensions” in this sense, but are loaded into the server in some other way, for instance by way of [shared_preload_libraries](../../server-administration/server-configuration/client-connection-defaults.md#guc-shared-preload-libraries). See the documentation of each component for details.

- [amcheck — tools to verify table and index consistency](amcheck-tools-to-verify-table-and-index-consistency.md#amcheck)
- [auth_delay — pause on authentication failure](auth_delay-pause-on-authentication-failure.md#auth-delay)
- [auto_explain — log execution plans of slow queries](auto_explain-log-execution-plans-of-slow-queries.md#auto-explain)
- [basebackup_to_shell — example "shell" pg_basebackup module](basebackup_to_shell-example-shell-pg_basebackup-module.md#basebackup-to-shell)
- [basic_archive — an example WAL archive module](basic_archive-an-example-wal-archive-module.md#basic-archive)
- [bloom — bloom filter index access method](bloom-bloom-filter-index-access-method.md#bloom)
- [btree_gin — GIN operator classes with B-tree behavior](btree_gin-gin-operator-classes-with-b-tree-behavior.md#btree-gin)
- [btree_gist — GiST operator classes with B-tree behavior](btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist)
- [citext — a case-insensitive character string type](citext-a-case-insensitive-character-string-type.md#citext)
- [cube — a multi-dimensional cube data type](cube-a-multi-dimensional-cube-data-type.md#cube)
- [dblink — connect to other PostgreSQL databases](dblink-connect-to-other-postgresql-databases.md#dblink)
- [dict_int — example full-text search dictionary for integers](dict_int-example-full-text-search-dictionary-for-integers.md#dict-int)
- [dict_xsyn — example synonym full-text search dictionary](dict_xsyn-example-synonym-full-text-search-dictionary.md#dict-xsyn)
- [earthdistance — calculate great-circle distances](earthdistance-calculate-great-circle-distances.md#earthdistance)
- [file_fdw — access data files in the server's file system](file_fdw-access-data-files-in-the-servers-file-system.md#file-fdw)
- [fuzzystrmatch — determine string similarities and distance](fuzzystrmatch-determine-string-similarities-and-distance.md#fuzzystrmatch)
- [hstore — hstore key/value datatype](hstore-hstore-key-value-datatype.md#hstore)
- [intagg — integer aggregator and enumerator](intagg-integer-aggregator-and-enumerator.md#intagg)
- [intarray — manipulate arrays of integers](intarray-manipulate-arrays-of-integers.md#intarray)
- [isn — data types for international standard numbers (ISBN, EAN, UPC, etc.)](isn-data-types-for-international-standard-numbers-isbn-ean-upc-etc.md#isn)
- [lo — manage large objects](lo-manage-large-objects.md#lo)
- [ltree — hierarchical tree-like data type](ltree-hierarchical-tree-like-data-type.md#ltree)
- [pageinspect — low-level inspection of database pages](pageinspect-low-level-inspection-of-database-pages.md#pageinspect)
- [passwordcheck — verify password strength](passwordcheck-verify-password-strength.md#passwordcheck)
- [pg_buffercache — inspect PostgreSQL buffer cache state](pg_buffercache-inspect-postgresql-buffer-cache-state.md#pgbuffercache)
- [pgcrypto — cryptographic functions](pgcrypto-cryptographic-functions.md#pgcrypto)
- [pg_freespacemap — examine the free space map](pg_freespacemap-examine-the-free-space-map.md#pgfreespacemap)
- [pg_logicalinspect — logical decoding components inspection](pg_logicalinspect-logical-decoding-components-inspection.md#pglogicalinspect)
- [pg_overexplain — allow EXPLAIN to dump even more details](pg_overexplain-allow-explain-to-dump-even-more-details.md#pgoverexplain)
- [pg_prewarm — preload relation data into buffer caches](pg_prewarm-preload-relation-data-into-buffer-caches.md#pgprewarm)
- [pgrowlocks — show a table's row locking information](pgrowlocks-show-a-tables-row-locking-information.md#pgrowlocks)
- [pg_stat_statements — track statistics of SQL planning and execution](pg_stat_statements-track-statistics-of-sql-planning-and-execution.md#pgstatstatements)
- [pgstattuple — obtain tuple-level statistics](pgstattuple-obtain-tuple-level-statistics.md#pgstattuple)
- [pg_surgery — perform low-level surgery on relation data](pg_surgery-perform-low-level-surgery-on-relation-data.md#pgsurgery)
- [pg_trgm — support for similarity of text using trigram matching](pg_trgm-support-for-similarity-of-text-using-trigram-matching.md#pgtrgm)
- [pg_visibility — visibility map information and utilities](pg_visibility-visibility-map-information-and-utilities.md#pgvisibility)
- [pg_walinspect — low-level WAL inspection](pg_walinspect-low-level-wal-inspection.md#pgwalinspect)
- [postgres_fdw — access data stored in external PostgreSQL servers](postgres_fdw-access-data-stored-in-external-postgresql-servers.md#postgres-fdw)
- [seg — a datatype for line segments or floating point intervals](seg-a-datatype-for-line-segments-or-floating-point-intervals.md#seg)
- [sepgsql — SELinux-, label-based mandatory access control (MAC) security module](sepgsql-selinux-label-based-mandatory-access-control-mac-security-module.md#sepgsql)
- [spi — Server Programming Interface features/examples](spi-server-programming-interface-features-examples.md#contrib-spi)
- [sslinfo — obtain client SSL information](sslinfo-obtain-client-ssl-information.md#sslinfo)
- [tablefunc — functions that return tables (`crosstab` and others)](tablefunc-functions-that-return-tables-crosstab-and-others.md#tablefunc)
- [tcn — a trigger function to notify listeners of changes to table content](tcn-a-trigger-function-to-notify-listeners-of-changes-to-table-content.md#tcn)
- [test_decoding — SQL-based test/example module for WAL logical decoding](test_decoding-sql-based-test-example-module-for-wal-logical-decoding.md#test-decoding)
- [tsm_system_rows — the `SYSTEM_ROWS` sampling method for `TABLESAMPLE`](tsm_system_rows-the-system_rows-sampling-method-for-tablesample.md#tsm-system-rows)
- [tsm_system_time — the `SYSTEM_TIME` sampling method for `TABLESAMPLE`](tsm_system_time-the-system_time-sampling-method-for-tablesample.md#tsm-system-time)
- [unaccent — a text search dictionary which removes diacritics](unaccent-a-text-search-dictionary-which-removes-diacritics.md#unaccent)
- [uuid-ossp — a UUID generator](uuid-ossp-a-uuid-generator.md#uuid-ossp)
- [xml2 — XPath querying and XSLT functionality](xml2-xpath-querying-and-xslt-functionality.md#xml2)
