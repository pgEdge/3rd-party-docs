
# What is pg_cron?

pg_cron is a simple cron-based job scheduler for PostgreSQL (10 or higher) that runs inside the database as an extension. 

---

# Contents
- [How pg_cron works](#how-pg_cron-works)
- [Cron syntax](#cron-syntax)
- [Managing and creating jobs](#managing-and-creating-jobs)
	- [Creating a cron job](#creating-a-cron-job)
	- [Creating a cron job in a different database](#creating-a-cron-job-in-a-different-database)
	- [Removing a cron job](#removing-a-cron-job)
	- [Altering a cron job](#altering-a-cron-job)
- [Installing pg_cron](#installing-pg_cron)
- [Setting up pg_cron](windows.md#setting-up-pg_cron)
- [Monitoring jobs](extension-settings.md#monitoring-jobs)
- [Example use cases](extension-settings.md#example-use-cases)
- [Managed services](extension-settings.md#managed-services)
- [Code of Conduct](extension-settings.md#code-of-conduct)

---

# How pg_cron works

The extension creates a background worker that tracks jobs in the `cron.job` table.

```sql
CREATE TABLE cron.job (
	jobid bigint primary key default pg_catalog.nextval('cron.jobid_seq'),
	schedule text not null,
	command text not null,
	nodename text not null default 'localhost',
	nodeport int not null default pg_catalog.inet_server_port(),
	database text not null default pg_catalog.current_database(),
	username text not null default current_user
);
```

Based on your configurations, to execute a job, the extension establishes a Postgres connection or spawns a database worker. 

pg_cron can run multiple jobs in parallel, but only one instance of each specific job at a time. If a second instance is triggered before the first finishes, it’s queued and starts as soon as the first one completes.

# Cron syntax

The code in pg_cron that handles parsing and scheduling comes directly from the [cron source code by Paul Vixie](https://github.com/vixie/cron), hence the same options are supported.
```
 ┌───────────── min (0 - 59)
 │ ┌────────────── hour (0 - 23)
 │ │ ┌─────────────── day of month (1 - 31) or last day of the month ($)
 │ │ │ ┌──────────────── month (1 - 12)
 │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
 │ │ │ │ │                  Saturday, or use names; 7 is also Sunday)
 │ │ │ │ │
 │ │ │ │ │
 * * * * *
```

An easy way to create a cron schedule is: [crontab.guru](http://crontab.guru/).

pg_cron also allows you:
- to use `$` to indicate last day of the month.
- to use `[1-59] seconds` to schedule a job based on an interval. Note, you cannot use seconds with the other time units.


Example cron schedules:

```
'10 seconds'  # every 10 seconds
* * * * *     # every minute
*/5 * * * *   # every 5 minutes
0 * * * *     # every hour
0 0 * * *     # daily at 12AM
0 0 * * 1-5   # 12AM every weekday
0 1 * * 0     # 1AM every Sunday
0 13 2 6 *    # 1PM on the 2nd of June
```

# Managing and creating jobs

Cron jobs can be managed by directly interacting with the `cron.job` table if you have the required permissions. However, it is recommended to use the cron functions:
- [`cron.schedule`](#creating-a-cron-job)
- [`cron.schedule_in_database`](#creating-a-cron-job-in-a-different-database)
- [`cron.unschedule`](#removing-a-cron-job)
- [`cron.alter_job`](#altering-a-cron-job)

> Note, an [RLS policy](https://www.postgresql.org/docs/current/ddl-rowsecurity.html) ensures that jobs can only be seen and modified by the user that created them, unless the user is a superuser or has the `bypassrls` attribute.

### Creating a cron job

#### `cron.schedule` signatures
```sql
-- create job, return jobid
CREATE OR REPLACE FUNCTION cron.schedule(schedule text, command text)
RETURNS bigint;

-- create named job, return jobid
CREATE OR REPLACE FUNCTION cron.schedule(job_name text, schedule text, command text)
RETURNS bigint
```

#### Examples

##### Create a cron job

```sql
-- Delete old data on Saturday at 3:30am (GMT)
SELECT cron.schedule(
       '30 3 * * 6', 
       $$DELETE FROM events WHERE event_time < now() - interval '1 week'$$
);
-- returns cron id
```

##### Create a named cron job

```sql
-- Vacuum every day at 10:00am (GMT)
SELECT cron.schedule(
       'nightly-vacuum', 
       '0 10 * * *', 
       'VACUUM'
);
-- returns cron id
```

##### Create a job that runs every 30 seconds

```sql
-- run SELECT 1 every 30 seconds
SELECT cron.schedule(
       'run_every_30_seconds', 
       '30 seconds', 
       'SELECT 1'
);
-- returns cron id
```

##### Create a job that calls a stored procedure every 5 seconds

```sql
-- Call a stored procedure every 5 seconds
SELECT cron.schedule(
       'process-updates',
       '5 seconds',
       'CALL process_updates()'
);
-- returns cron id
```

##### Create a job that processes payroll at 12:00 of the last day of each month

```sql
-- Process payroll at 12:00 of the last day of each month
SELECT cron.schedule(
       'process-payroll',
       '0 12 $ * *',
       'CALL process_payroll()'
);
-- returns cron id
```

### Creating a cron job in a different database

#### `cron.schedule_in_database` signature
```sql
-- create job, return jobid
CREATE OR REPLACE FUNCTION cron.schedule_in_database(
       job_name text, 
       schedule text, 
       command text, 
       database text, 
       username text DEFAULT NULL::text, 
       active boolean DEFAULT true
)
RETURNS bigint
```

#### Example

##### Create a cron job in a different database

```sql
-- Delete old data on Saturday at 3:30am (GMT)
SELECT cron.schedule_in_database(
       'delete_old_data', 
       '30 3 * * 6', 
       $$DELETE FROM events WHERE event_time < now() - interval '1 week'$$,
       'some_other_database'
);
-- returns cron id
```

### Removing a cron job

#### `cron.unschedule` signatures
```sql
-- remove job by name, return true if job was removed
CREATE OR REPLACE FUNCTION cron.unschedule(job_name text)
RETURNS boolean

-- remove job by id, return true if job was removed
CREATE OR REPLACE FUNCTION cron.unschedule(job_id bigint)
RETURNS boolean
```

#### Examples

##### Remove a named cron job

```sql
-- delete job by name
SELECT cron.unschedule('nightly-vacuum');
-- returns true if job was removed
```

##### Remove a cron job by id

```sql
-- delete job by id
SELECT cron.unschedule(42);
-- returns true if job was removed
```

### Altering a cron job

#### `cron.alter_job` signature
```sql
CREATE OR REPLACE FUNCTION cron.alter_job(
       job_id bigint, 
       schedule text DEFAULT NULL::text, 
       command text DEFAULT NULL::text, 
       database text DEFAULT NULL::text, 
       username text DEFAULT NULL::text, 
       active boolean DEFAULT NULL::boolean
)
RETURNS void
```

#### Examples

##### Change a job's schedule

```sql
-- change job's schedule
SELECT cron.alter_job(42, '0 10 * * *');
-- returns void
```

##### Change a job's, schedule, command, and username

```sql
-- change job's command
SELECT cron.alter_job(
       42,
       '0 10 * * *',
       'VACUUM',
       username := 'some_other_user'
);
-- returns void
```

##### Deactivate a job

```sql
-- deactivate job
SELECT cron.alter_job(42, active := false);
-- returns void
```

# Installing pg_cron

Install on Red Hat, CentOS, Fedora, Amazon Linux with PostgreSQL 18 using [PGDG](https://yum.postgresql.org/repopackages/):

```bash
# Install the pg_cron extension
sudo yum install -y pg_cron_18
```

Install on Debian, Ubuntu with PostgreSQL 18 using [apt.postgresql.org](https://wiki.postgresql.org/wiki/Apt):

```bash
# Install the pg_cron extension
sudo apt-get -y install postgresql-18-cron
```

You can also install pg_cron by building it from source:

```bash
git clone https://github.com/citusdata/pg_cron.git
cd pg_cron
# Ensure pg_config is in your path, e.g.
export PATH=/usr/pgsql-18/bin:$PATH
make && sudo PATH=$PATH make install
```
