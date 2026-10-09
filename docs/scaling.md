# Scaling

For a smaller working set:

1. Use the `halfvec` type instead of `vector` for tables
2. Use [binary quantization](binary-quantization.md) for indexes (with re-ranking for search)

Scale vertically by increasing memory, CPU, and storage on a single instance. Use existing tools to [tune parameters](performance.md#tuning) and [monitor performance](monitoring.md).

Scale horizontally with [replicas](https://www.postgresql.org/docs/current/hot-standby.html), or use [Citus](https://github.com/citusdata/citus), [PgDog](https://github.com/pgdogdev/pgdog), or another approach for sharding ([example](https://github.com/pgvector/pgvector-python/blob/master/examples/citus/example.py)).

