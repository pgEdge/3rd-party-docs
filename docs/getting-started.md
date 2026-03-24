# Getting Started
For new users, we recommend using `tensorchord/vchord-suite` image to get started quickly, you can find more details in the [VectorChord-images](https://github.com/tensorchord/VectorChord-images) repository.

```sh
docker run \
  --name vchord-suite \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  -d tensorchord/vchord-suite:pg18-latest
  # If you want to use ghcr image, you can change the image to `ghcr.io/tensorchord/vchord-suite:pg18-latest`.
  # if you want to use the specific version, you can use the tag `pg17-20250414`, supported version can be found in the support matrix.
```

Once everything’s set up, you can connect to the database using the `psql` command line tool. The default username is `postgres`, and the default password is `postgres`. Here’s how to connect:

```sh
psql -h localhost -p 5432 -U postgres
```

After connecting, run the following SQL to make sure the extension is enabled:

```sql
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;  -- for tokenizer
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;   -- for bm25 ranking
```

