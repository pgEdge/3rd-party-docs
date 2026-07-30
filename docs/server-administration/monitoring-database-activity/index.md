<a id="monitoring"></a>

# Monitoring Database Activity

 A database administrator frequently wonders, “What is the system doing right now?” This chapter discusses how to find that out.

 Several tools are available for monitoring database activity and analyzing performance. Most of this chapter is devoted to describing PostgreSQL's cumulative statistics system, but one should not neglect regular Unix monitoring programs such as `ps`, `top`, `iostat`, and `vmstat`. Also, once one has identified a poorly-performing query, further investigation might be needed using PostgreSQL's [`EXPLAIN`](../../reference/sql-commands/explain.md#sql-explain) command. [Using `EXPLAIN`](../../the-sql-language/performance-tips/using-explain.md#using-explain) discusses `EXPLAIN` and other methods for understanding the behavior of an individual query.

- [Standard Unix Tools](standard-unix-tools.md#monitoring-ps)
- [The Cumulative Statistics System](the-cumulative-statistics-system.md#monitoring-stats)
- [Viewing Locks](viewing-locks.md#monitoring-locks)
- [Progress Reporting](progress-reporting.md#progress-reporting)
- [Dynamic Tracing](dynamic-tracing.md#dynamic-trace)
- [Monitoring Disk Usage](monitoring-disk-usage.md#diskusage)
