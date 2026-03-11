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
<td><p><code>parent</code> <code>text</code></p>
<p>Name of the parent of this memory context</p></td>
</tr>
<tr>
<td><p><code>level</code> <code>int4</code></p>
<p>Distance from TopMemoryContext in context tree</p></td>
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
