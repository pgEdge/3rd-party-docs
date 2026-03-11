<a id="indexam"></a>

# Index Access Method Interface Definition

 This chapter defines the interface between the core PostgreSQL system and *index access methods*, which manage individual index types. The core system knows nothing about indexes beyond what is specified here, so it is possible to develop entirely new index types by writing add-on code.

 All indexes in PostgreSQL are what are known technically as *secondary indexes*; that is, the index is physically separate from the table file that it describes. Each index is stored as its own physical *relation* and so is described by an entry in the `pg_class` catalog. The contents of an index are entirely under the control of its index access method. In practice, all index access methods divide indexes into standard-size pages so that they can use the regular storage manager and buffer manager to access the index contents. (All the existing index access methods furthermore use the standard page layout described in [Database Page Layout](../database-physical-storage/database-page-layout.md#storage-page-layout), and most use the same format for index tuple headers; but these decisions are not forced on an access method.)

 An index is effectively a mapping from some data key values to *tuple identifiers*, or TIDs, of row versions (tuples) in the index's parent table. A TID consists of a block number and an item number within that block (see [Database Page Layout](../database-physical-storage/database-page-layout.md#storage-page-layout)). This is sufficient information to fetch a particular row version from the table. Indexes are not directly aware that under MVCC, there might be multiple extant versions of the same logical row; to an index, each tuple is an independent object that needs its own index entry. Thus, an update of a row always creates all-new index entries for the row, even if the key values did not change. ([HOT tuples](../database-physical-storage/heap-only-tuples-hot.md#storage-hot) are an exception to this statement; but indexes do not deal with those, either.) Index entries for dead tuples are reclaimed (by vacuuming) when the dead tuples themselves are reclaimed.

- [Basic API Structure for Indexes](basic-api-structure-for-indexes.md#index-api)
- [Index Access Method Functions](index-access-method-functions.md#index-functions)
- [Index Scanning](index-scanning.md#index-scanning)
- [Index Locking Considerations](index-locking-considerations.md#index-locking)
- [Index Uniqueness Checks](index-uniqueness-checks.md#index-unique-checks)
- [Index Cost Estimation Functions](index-cost-estimation-functions.md#index-cost-estimation)
