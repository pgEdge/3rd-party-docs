# Extension settings

The pg_cron extension supports the following configuration parameters:

| Setting                          | Default     | Description                                                                              |
| ---------------------------------| ----------- | ---------------------------------------------------------------------------------------- |
| `cron.database_name`             | `postgres`  | Database in which the pg_cron background worker should run.                              |
| `cron.enable_superuser_jobs`     | `on`        | Allow jobs to be scheduled as superusers.                                                |
| `cron.host`                      | `localhost` | Hostname to connect to postgres.                                                         |
| `cron.launch_active_jobs`        | `on`        | When off, disables all active jobs without requiring a server restart                    |
| `cron.log_min_messages`          | `WARNING`   | log_min_messages for the launcher bgworker.                                              |
| `cron.log_run`                   | `on`        | Log all run details in the`cron.job_run_details` table.                                  |
| `cron.log_statement`             | `on`        | Log all cron statements prior to execution.                                              |
| `cron.max_running_jobs`          | `32`        | Maximum number of jobs that can be running at the same time.                             |
| `cron.timezone`                  | `GMT`       | Timezone in which the pg_cron background worker should run.                              |
| `cron.use_background_workers`    | `off`       | Use background workers instead of client connections.                                    |

## Changing settings

To view setting configurations, run:

```sql
SELECT * FROM pg_settings WHERE name LIKE 'cron.%';
```

Setting can be changed in the postgresql.conf file or with the below command:

```sql
ALTER SYSTEM SET cron.<parameter> TO '<value>';
```

`cron.log_min_messages` and `cron.launch_active_jobs` have a [setting context](https://www.postgresql.org/docs/current/view-pg-settings.html#VIEW-PG-SETTINGS) of `sighup`. They can be finalized by executing `SELECT pg_reload_conf();`.

All the other settings have a postmaster context and only take effect after a server restart.

# Monitoring jobs

## Reviewing the `cron.job_run_details` table

You can view job activity in the `cron.job_run_details` table:

```sql
select * from cron.job_run_details order by start_time desc limit 5;
┌───────┬───────┬─────────┬──────────┬──────────┬───────────────────┬───────────┬──────────────────┬───────────────────────────────┬───────────────────────────────┐
│ jobid │ runid │ job_pid │ database │ username │      command      │  status   │  return_message  │          start_time           │           end_time            │
├───────┼───────┼─────────┼──────────┼──────────┼───────────────────┼───────────┼──────────────────┼───────────────────────────────┼───────────────────────────────┤
│    11 │  4328 │    2610 │ postgres │ marco    │ select pg_sleep(3)│ running   │ NULL             │ 2023-02-07 09:30:00.098164+01 │ NULL                          │
│    10 │  4327 │    2609 │ postgres │ marco    │ select process()  │ succeeded │ SELECT 1         │ 2023-02-07 09:29:00.015168+01 │ 2023-02-07 09:29:00.832308+01 │
│    10 │  4321 │    2603 │ postgres │ marco    │ select process()  │ succeeded │ SELECT 1         │ 2023-02-07 09:28:00.011965+01 │ 2023-02-07 09:28:01.420901+01 │
│    10 │  4320 │    2602 │ postgres │ marco    │ select process()  │ failed    │ server restarted │ 2023-02-07 09:27:00.011833+01 │ 2023-02-07 09:27:00.72121+01  │
│     9 │  4320 │    2602 │ postgres │ marco    │ select do_stuff() │ failed    │ job canceled     │ 2023-02-07 09:26:00.011833+01 │ 2023-02-07 09:26:00.22121+01  │
└───────┴───────┴─────────┴──────────┴──────────┴───────────────────┴───────────┴──────────────────┴───────────────────────────────┴───────────────────────────────┘
(10 rows)
```

The records in the table are not cleaned automatically, but every user that can schedule cron jobs also has permission to delete their own `cron.job_run_details` records. 

Especially when you have jobs that run every few seconds, it can be a good idea to clean up regularly, which can easily be done using pg_cron itself:

```sql
-- Delete old cron.job_run_details records of the current user every day at noon
SELECT  cron.schedule('delete-job-run-details', '0 12 * * *', $$DELETE FROM cron.job_run_details WHERE end_time < now() - interval '7 days'$$);
```

If you do not want to use `cron.job_run_details` at all, then you can add `cron.log_run = off` to `postgresql.conf`.

## Other cron logging settings

If the `cron.log_statement` setting is configured, jobs will be logged before execution. The `cron.log_min_messages` setting controls the [minimum level of messages](https://www.postgresql.org/docs/current/runtime-config-logging.html#RUNTIME-CONFIG-SEVERITY-LEVELS) that will be recorded.

# Example use cases

Articles showing possible ways of using pg_cron:

* [Auto-partitioning using pg_partman](https://www.citusdata.com/blog/2018/01/24/citus-and-pg-partman-creating-a-scalable-time-series-database-on-postgresql/)
* [Computing rollups in an analytical dashboard](https://www.citusdata.com/blog/2017/12/27/real-time-analytics-dashboards-with-citus/)
* [Deleting old data, vacuum](https://www.citusdata.com/blog/2016/09/09/pgcron-run-periodic-jobs-in-postgres/)
* [Feeding cats](http://bonesmoses.org/2016/09/09/pg-phriday-irrelevant-inclinations/)
* [Routinely invoking a function](https://fluca1978.github.io/2019/05/21/pgcron.html)
* [Postgres as a cron server](https://supabase.io/blog/2021/03/05/postgres-as-a-cron-server)

# Managed services

The following table keeps track of which of the major managed Postgres services support pg_cron.

| Service       | Supported     |
| ------------- |:-------------:|
| [Aiven](https://aiven.io/postgresql) | ✔️ |
| [Alibaba Cloud](https://www.alibabacloud.com/help/doc-detail/150355.htm) | ✔️ |
| [Amazon RDS](https://aws.amazon.com/rds/postgresql/)     | ✔️      |          |
| [Azure](https://azure.microsoft.com/en-us/services/postgresql/) | ✔️  |
| [Crunchy Bridge](https://www.crunchydata.com/products/crunchy-bridge/?ref=producthunt) | ✔️ |
| [DigitalOcean](https://www.digitalocean.com/products/managed-databases/) | ✔️ |
| [Google Cloud](https://cloud.google.com/sql/postgresql/) | ✔️ |
| [Heroku](https://elements.heroku.com/addons/heroku-postgresql) | ❌ |
| [Instaclustr](https://instaclustr.com) | ✔️  |
| [Neon](https://neon.tech/docs/extensions/extensions-intro#tooling-admin) | ✔️ |
| [PlanetScale](https://planetscale.com/docs/postgres/extensions) | ✔️ |
| [ScaleGrid](https://scalegrid.io/postgresql.html) | ✔️  |
| [Scaleway](https://www.scaleway.com/en/database/) | ✔️  |
| [Supabase](https://supabase.io/docs/guides/database) | ✔️  |
| [YugabyteDB](https://www.yugabyte.com/) | ✔️  |

# Code of Conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

