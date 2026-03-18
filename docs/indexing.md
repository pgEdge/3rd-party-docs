# Indexing

By default, pgvector performs exact nearest neighbor search, which provides perfect recall.

You can add an index to use approximate nearest neighbor search, which trades some recall for speed. Unlike typical indexes, you will see different results for queries after adding an approximate index.

Supported index types are:

- [HNSW](iterative-index-scans.md#hnsw)
- [IVFFlat](iterative-index-scans.md#ivfflat)

