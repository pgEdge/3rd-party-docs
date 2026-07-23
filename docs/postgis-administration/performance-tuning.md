<a id="database_tuning_configuration"></a>

## Performance Tuning


Tuning for PostGIS performance is much like tuning for any PostgreSQL workload. The only additional consideration is that geometries and rasters are usually large, so memory-related optimizations generally have more of an impact on PostGIS than other types of PostgreSQL queries.


For general details about optimizing PostgreSQL, refer to [Tuning your PostgreSQL Server](https://wiki.postgresql.org/wiki/Tuning_Your_PostgreSQL_Server).


For PostgreSQL 9.4+ configuration can be set at the server level without touching <code>postgresql.conf</code> or <code>postgresql.auto.conf</code> by using the <code>ALTER SYSTEM</code> command.


```sql
ALTER SYSTEM SET work_mem = '256MB';
-- this forces non-startup configs to take effect for new connections
SELECT pg_reload_conf();
-- show current setting value
-- use SHOW ALL to see all settings
SHOW work_mem;
```


In addition to the Postgres settings, PostGIS has some custom settings which are listed in [Grand Unified Custom Variables (GUCs)](../postgis-reference/grand-unified-custom-variables-gucs.md#PostGIS_GUC).


## Startup


 These settings are configured in <code>postgresql.conf</code>:


 [constraint_exclusion](http://www.postgresql.org/docs/current/static/runtime-config-query.html#GUC-CONSTRAINT-EXCLUSION)


-  Default: partition
-  This is generally used for table partitioning. The default for this is set to "partition" which is ideal for PostgreSQL 8.4 and above since it will force the planner to only analyze tables for constraint consideration if they are in an inherited hierarchy and not pay the planner penalty otherwise.


 [shared_buffers](http://www.postgresql.org/docs/current/static/runtime-config-resource.html#GUC-SHARED-BUFFERS)


-  Default: ~128MB in PostgreSQL 9.6
-  Set to about 25% to 40% of available RAM. On windows you may not be able to set as high.


 [max_worker_processes](https://www.postgresql.org/docs/current/static/runtime-config-resource.html#GUC-MAX-WORKER-PROCESSES) This setting is only available for PostgreSQL 9.4+. For PostgreSQL 9.6+ this setting has additional importance in that it controls the max number of processes you can have for parallel queries.


-  Default: 8
-  Sets the maximum number of background processes that the system can support. This parameter can only be set at server start.


## Runtime


 [work_mem](http://www.postgresql.org/docs/current/static/runtime-config-resource.html#GUC-WORK-MEM) - sets the size of memory used for sort operations and complex queries


-  Default: 1-4MB
-  Adjust up for large dbs, complex queries, lots of RAM
-  Adjust down for many concurrent users or low RAM.
-  If you have lots of RAM and few developers:

```sql
SET work_mem TO '256MB';
```


 [maintenance_work_mem](http://www.postgresql.org/docs/current/static/runtime-config-resource.html#GUC-MAINTENANCE-WORK-MEM) - the memory size used for VACUUM, CREATE INDEX, etc.


-  Default: 16-64MB
-  Generally too low - ties up I/O, locks objects while swapping memory
-  Recommend 32MB to 1GB on production servers w/lots of RAM, but depends on the # of concurrent users. If you have lots of RAM and few developers:

```sql
SET maintenance_work_mem TO '1GB';
```


 [max_parallel_workers_per_gather](https://www.postgresql.org/docs/current/static/runtime-config-resource.html#GUC-MAX-PARALLEL-WORKERS-PER-GATHER)


 This setting is only available for PostgreSQL 9.6+ and will only affect PostGIS 2.3+, since only PostGIS 2.3+ supports parallel queries. If set to higher than 0, then some queries such as those involving relation functions like <code>ST_Intersects</code> can use multiple processes and can run more than twice as fast when doing so. If you have a lot of processors to spare, you should change the value of this to as many processors as you have. Also make sure to bump up <code>max_worker_processes</code> to at least as high as this number.


-  Default: 0
-  Sets the maximum number of workers that can be started by a single `Gather` node. Parallel workers are taken from the pool of processes established by `max_worker_processes`. Note that the requested number of workers may not actually be available at run time. If this occurs, the plan will run with fewer workers than expected, which may be inefficient. Setting this value to 0, which is the default, disables parallel query execution.
