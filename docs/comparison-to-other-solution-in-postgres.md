# Comparison to other solution in Postgres
PostgreSQL supports full-text search using the tsvector data type and GIN indexes. Text is transformed into a tsvector, which tokenizes content into standardized lexemes, and a GIN index accelerates searches—even on large text fields. However, PostgreSQL lacks modern relevance scoring methods like BM25; it retrieves all matching documents and re-ranks them using ts_rank, which is inefficient and can obscure the most relevant results.

ParadeDB is an alternative that functions as a full-featured PostgreSQL replacement for ElasticSearch. It offloads full-text search and filtering operations to Tantivy and includes BM25 among its features, though it uses a different query and filter syntax than PostgreSQL's native indexes.

In contrast, Vectorchord-bm25 focuses exclusively on BM25 ranking within PostgreSQL. We implemented the BM25 ranking algorithm Block WeakAnd from scratch and built it as a custom operator and index (similar to pgvector) to accelerate queries. It is designed to be lightweight and a more native and intuitive API for better full-text search and ranking in PostgreSQL.

