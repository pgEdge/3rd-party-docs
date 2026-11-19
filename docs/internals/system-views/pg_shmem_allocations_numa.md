<a id="view-pg-shmem-allocations-numa"></a>

## `pg_shmem_allocations_numa`


 The `pg_shmem_allocations_numa` shows how shared memory allocations in the server's main shared memory segment are distributed across NUMA nodes. This includes both memory allocated by PostgreSQL itself and memory allocated by extensions using the mechanisms detailed in [Shared Memory](../../server-programming/extending-sql/c-language-functions.md#xfunc-shared-addin). This view will output multiple rows for each of the shared memory segments provided that they are spread across multiple NUMA nodes. This view should not be queried by monitoring systems as it is very slow and may end up allocating shared memory in case it was not used earlier. Current limitation for this view is that won't show anonymous shared memory allocations.


 Note that this view does not include memory allocated using the dynamic shared memory infrastructure.


!!! warning

    When determining the NUMA node, the view touches all memory pages for the shared memory segment. This will force allocation of the shared memory, if it wasn't allocated already, and the memory may get allocated in a single NUMA node (depending on system configuration).


**Table: `pg_shmem_allocations_numa` Columns**

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
<p>The name of the shared memory allocation.</p></td>
</tr>
<tr>
<td><p><code>numa_node</code> <code>int4</code></p>
<p>ID of NUMA node</p></td>
</tr>
<tr>
<td><p><code>size</code> <code>int8</code></p>
<p>Size of the allocation on this particular NUMA memory node in bytes</p></td>
</tr>
</tbody>
</table>


 By default, the `pg_shmem_allocations_numa` view can be read only by superusers or roles with privileges of the `pg_read_all_stats` role.
