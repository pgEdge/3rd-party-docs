# Storing

Create a new table with a vector column

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
```

Or add a vector column to an existing table

```sql
ALTER TABLE items ADD COLUMN embedding vector(3);
```

Also supports [half-precision](half-precision-vectors.md), [binary](binary-vectors.md), and [sparse](sparse-vectors.md) vectors

Insert vectors

```sql
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
```

Or load vectors in bulk using `COPY` ([example](https://github.com/pgvector/pgvector-python/blob/master/examples/loading/example.py))

```sql
COPY items (embedding) FROM STDIN WITH (FORMAT BINARY);
```

Upsert vectors

```sql
INSERT INTO items (id, embedding) VALUES (1, '[1,2,3]'), (2, '[4,5,6]')
    ON CONFLICT (id) DO UPDATE SET embedding = EXCLUDED.embedding;
```

Update vectors

```sql
UPDATE items SET embedding = '[1,2,3]' WHERE id = 1;
```

Delete vectors

```sql
DELETE FROM items WHERE id = 1;
```

