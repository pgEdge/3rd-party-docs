# Windows

Ensure [C++ support in Visual Studio](https://learn.microsoft.com/en-us/cpp/build/building-on-the-command-line?view=msvc-170#download-and-install-the-tools) is installed and run `x64 Native Tools Command Prompt for VS [version]` as administrator. Then use `nmake` to build:

```cmd
set "PGROOT=C:\Program Files\PostgreSQL\18"
cd %TMP%
git clone https://github.com/citusdata/pg_cron.git
cd pg_cron
nmake /F Makefile.win
nmake /F Makefile.win install
```

## Installation Notes - Windows

### Missing Header

If compilation fails with `Cannot open include file: 'postgres.h': No such file or directory`, make sure `PGROOT` is correct.

### Permissions

If installation fails with `Access is denied`, re-run the installation instructions as an administrator.

# Setting up pg_cron

To start the pg_cron background worker, you need to add pg_cron to `shared_preload_libraries` in postgresql.conf. Note that pg_cron does not run any jobs as a long a server is in [hot standby](https://www.postgresql.org/docs/current/static/hot-standby.html) mode, but it automatically starts when the server is promoted.

```
# add to postgresql.conf

# required to load pg_cron background worker on start-up
shared_preload_libraries = 'pg_cron'
```

By default, the pg_cron background worker expects its metadata tables to be created in the "postgres" database. However, you can configure this by setting the `cron.database_name` configuration parameter in postgresql.conf.
```
# add to postgresql.conf

# optionally, specify the database in which the pg_cron background worker should run (defaults to postgres)
cron.database_name = 'postgres'
```
`pg_cron` may only be installed to one database in a cluster. If you need to run jobs in multiple databases, use `cron.schedule_in_database()`.

Previously pg_cron could only use GMT time, but now you can adapt your time by setting `cron.timezone` in postgresql.conf.
```
# add to postgresql.conf

# optionally, specify the timezone in which the pg_cron background worker should run (defaults to GMT). E.g:
cron.timezone = 'PRC'
```

After restarting PostgreSQL, you can create the pg_cron functions and metadata tables using `CREATE EXTENSION pg_cron`.

```sql
-- run as superuser:
CREATE EXTENSION pg_cron;

-- optionally, grant usage to regular users:
GRANT USAGE ON SCHEMA cron TO marco;
```

## Ensuring pg_cron can start jobs

**Important**: By default, pg_cron uses libpq to open a new connection to the local database, which needs to be allowed by [pg_hba.conf](https://www.postgresql.org/docs/current/static/auth-pg-hba-conf.html). 
It may be necessary to enable `trust` authentication for connections coming from localhost in  for the user running the cron job, or you can add the password to a [.pgpass file](https://www.postgresql.org/docs/current/static/libpq-pgpass.html), which libpq will use when opening a connection. 

You can also use a unix domain socket directory as the hostname and enable `trust` authentication for local connections in [pg_hba.conf](https://www.postgresql.org/docs/current/static/auth-pg-hba-conf.html), which is normally safe:
```
# Connect via a unix domain socket:
cron.host = '/tmp'

# Can also be an empty string to look for the default directory:
cron.host = ''
```

Alternatively, pg_cron can be configured to use background workers. In that case, the number of concurrent jobs is limited by the `max_worker_processes` setting, so you may need to raise that.

```
# Schedule jobs via background workers instead of localhost connections
cron.use_background_workers = on
# Increase the number of available background workers from the default of 8
max_worker_processes = 20
```

For security, jobs are executed in the database in which the `cron.schedule` function is called with the same permissions as the current user. In addition, users are only able to see their own jobs in the `cron.job` table.

```sql
-- View active jobs
select * from cron.job;
```

