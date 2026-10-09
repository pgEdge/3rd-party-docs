# Hybrid Search

Use together with Postgres [full-text search](https://www.postgresql.org/docs/current/textsearch-intro.html) for hybrid search.

```sql
SELECT id, content FROM items, plainto_tsquery('hello search') query
    WHERE textsearch @@ query ORDER BY ts_rank_cd(textsearch, query) DESC LIMIT 5;
```

You can use [Reciprocal Rank Fusion](https://github.com/pgvector/pgvector-python/blob/master/examples/hybrid_search/rrf.py) or a [cross-encoder](https://github.com/pgvector/pgvector-python/blob/master/examples/hybrid_search/cross_encoder.py) to combine results.

