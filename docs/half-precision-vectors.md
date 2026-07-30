# Half-Precision Vectors

Use the `halfvec` type to store half-precision vectors

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding halfvec(3));
```

