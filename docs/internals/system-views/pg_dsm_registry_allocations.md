<a id="view-pg-dsm-registry-allocations"></a>

## `pg_dsm_registry_allocations`


 The `pg_dsm_registry_allocations` view shows shared memory allocations tracked in the dynamic shared memory (DSM) registry. This includes memory allocated by extensions using the mechanisms detailed in [Requesting Shared Memory After Startup with `ShmemRequestStruct`](../../server-programming/extending-sql/c-language-functions.md#xfunc-shared-addin-after-startup).


**Table: `pg_dsm_registry_allocations` Columns**

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
<p>The name of the allocation in the DSM registry.</p></td>
</tr>
<tr>
<td><p><code>type</code> <code>text</code></p>
<p>The type of allocation. Possible values are <code>segment</code>, <code>area</code>, and <code>hash</code>, which correspond to dynamic shared memory segments, areas, and hash tables, respectively.</p></td>
</tr>
<tr>
<td><p><code>size</code> <code>int8</code></p>
<p>Size of the allocation in bytes. NULL for entries that failed initialization.</p></td>
</tr>
</tbody>
</table>


 By default, the `pg_dsm_registry_allocations` view can be read only by superusers or roles with privileges of the `pg_read_all_stats` role.
