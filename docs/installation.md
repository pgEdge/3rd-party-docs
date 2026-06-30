# Installation

Clone this repo and run

```bash
make && make install
```

To make supautils available to the whole cluster, you can add the following to `postgresql.conf` (use `SHOW config_file` for finding the location).

```
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

Or to make it available only on some PostgreSQL roles use `session_preload_libraries`.

```
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

