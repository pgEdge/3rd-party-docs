<a id="parallel-query"></a>

# Parallel Query

 PostgreSQL can devise query plans that can leverage multiple CPUs in order to answer queries faster. This feature is known as parallel query. Many queries cannot benefit from parallel query, either due to limitations of the current implementation or because there is no imaginable query plan that is any faster than the serial query plan. However, for queries that can benefit, the speedup from parallel query is often very significant. Many queries can run more than twice as fast when using parallel query, and some queries can run four times faster or even more. Queries that touch a large amount of data but return only a few rows to the user will typically benefit most. This chapter explains some details of how parallel query works and in which situations it can be used so that users who wish to make use of it can understand what to expect.

- [How Parallel Query Works](how-parallel-query-works.md#how-parallel-query-works)
- [When Can Parallel Query Be Used?](when-can-parallel-query-be-used.md#when-can-parallel-query-be-used)
- [Parallel Plans](parallel-plans.md#parallel-plans)
- [Parallel Safety](parallel-safety.md#parallel-safety)
