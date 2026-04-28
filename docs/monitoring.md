# Monitoring

Use existing tools like [pg_stat_statements](https://www.postgresql.org/docs/current/pgstatstatements.html) or [PgHero](https://github.com/ankane/pghero) to monitor performance.

Monitor recall by comparing results from approximate search with exact search.

```sql
BEGIN;
SET LOCAL enable_indexscan = off; -- use exact search
SELECT ...
COMMIT;
```

