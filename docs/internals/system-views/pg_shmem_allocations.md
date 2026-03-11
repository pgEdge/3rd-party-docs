<a id="view-pg-shmem-allocations"></a>

## `pg_shmem_allocations`


 The `pg_shmem_allocations` view shows allocations made from the server's main shared memory segment. This includes both memory allocated by PostgreSQL itself and memory allocated by extensions using the mechanisms detailed in [Shared Memory](../../server-programming/extending-sql/c-language-functions.md#xfunc-shared-addin).


 Note that this view does not include memory allocated using the dynamic shared memory infrastructure.


**Table: `pg_shmem_allocations` Columns**

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
<p>The name of the shared memory allocation. NULL for unused memory and <code>&lt;anonymous&gt;</code> for anonymous allocations.</p></td>
</tr>
<tr>
<td><p><code>off</code> <code>int8</code></p>
<p>The offset at which the allocation starts. NULL for anonymous allocations, since details related to them are not known.</p></td>
</tr>
<tr>
<td><p><code>size</code> <code>int8</code></p>
<p>Size of the allocation in bytes</p></td>
</tr>
<tr>
<td><p><code>allocated_size</code> <code>int8</code></p>
<p>Size of the allocation in bytes including padding. For anonymous allocations, no information about padding is available, so the <code>size</code> and <code>allocated_size</code> columns will always be equal. Padding is not meaningful for free memory, so the columns will be equal in that case also.</p></td>
</tr>
</tbody>
</table>


 Anonymous allocations are allocations that have been made with `ShmemAlloc()` directly, rather than via `ShmemInitStruct()` or `ShmemInitHash()`.


 By default, the `pg_shmem_allocations` view can be read only by superusers or roles with privileges of the `pg_read_all_stats` role.
