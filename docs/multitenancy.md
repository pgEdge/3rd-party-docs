# Multitenancy

For applications with multiple tenants, sharing an approximate index between tenants means vectors from one tenant can affect recall (and speed) for other tenants.

For tenant isolation, use [list partitioning](https://www.postgresql.org/docs/current/ddl-partitioning.html) or separate tables.

```sql
CREATE TABLE items (customer_id int, embedding vector(3)) PARTITION BY LIST(customer_id);
```

