<a id="view-pg-backend-memory-contexts"></a>

## `pg_backend_memory_contexts`


 The view `pg_backend_memory_contexts` displays all the memory contexts of the server process attached to the current session.


 `pg_backend_memory_contexts` contains one row for each memory context.


**Table: `pg_backend_memory_contexts` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>name</code> <code>text</code></p>
<p>Name of the memory context</p></td>
</tr>
<tr>
<td><p><code>ident</code> <code>text</code></p>
<p>Identification information of the memory context. This field is truncated at 1024 bytes</p></td>
</tr>
<tr>
<td><p><code>type</code> <code>text</code></p>
<p>Type of the memory context</p></td>
</tr>
<tr>
<td><p><code>level</code> <code>int4</code></p>
<p>The 1-based level of the context in the memory context hierarchy. The level of a context also shows the position of that context in the <code>path</code> column.</p></td>
</tr>
<tr>
<td><p><code>path</code> <code>int4[]</code></p>
<p>Array of transient numerical identifiers to describe the memory context hierarchy. The first element is for <code>TopMemoryContext</code>, subsequent elements contain intermediate parents and the final element contains the identifier for the current context.</p></td>
</tr>
<tr>
<td><p><code>total_bytes</code> <code>int8</code></p>
<p>Total bytes allocated for this memory context</p></td>
</tr>
<tr>
<td><p><code>total_nblocks</code> <code>int8</code></p>
<p>Total number of blocks allocated for this memory context</p></td>
</tr>
<tr>
<td><p><code>free_bytes</code> <code>int8</code></p>
<p>Free space in bytes</p></td>
</tr>
<tr>
<td><p><code>free_chunks</code> <code>int8</code></p>
<p>Total number of free chunks</p></td>
</tr>
<tr>
<td><p><code>used_bytes</code> <code>int8</code></p>
<p>Used space in bytes</p></td>
</tr>
</tbody>
</table>


 By default, the `pg_backend_memory_contexts` view can be read only by superusers or roles with the privileges of the `pg_read_all_stats` role.


 Since memory contexts are created and destroyed during the running of a query, the identifiers stored in the `path` column can be unstable between multiple invocations of the view in the same query. The example below demonstrates an effective usage of this column and calculates the total number of bytes used by `CacheMemoryContext` and all of its children:

```sql

WITH memory_contexts AS (
    SELECT * FROM pg_backend_memory_contexts
)
SELECT sum(c1.total_bytes)
FROM memory_contexts c1, memory_contexts c2
WHERE c2.name = 'CacheMemoryContext'
AND c1.path[c2.level] = c2.path[c2.level];
```
 The [Common Table Expression](../../the-sql-language/queries/with-queries-common-table-expressions.md#queries-with) is used to ensure the context IDs in the `path` column match between both evaluations of the view.
