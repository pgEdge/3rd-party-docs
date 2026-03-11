<a id="view-pg-replication-origin-status"></a>

## `pg_replication_origin_status`


 The `pg_replication_origin_status` view contains information about how far replay for a certain origin has progressed. For more on replication origins see [Replication Progress Tracking](../../server-programming/replication-progress-tracking.md#replication-origins).


**Table: `pg_replication_origin_status` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>local_id</code> <code>oid</code> (references <a href="../system-catalogs/pg_replication_origin.md#catalog-pg-replication-origin"><code>pg_replication_origin</code></a>.<code>roident</code>)</p>
<p>internal node identifier</p></td>
</tr>
<tr>
<td><p><code>external_id</code> <code>text</code> (references <a href="../system-catalogs/pg_replication_origin.md#catalog-pg-replication-origin"><code>pg_replication_origin</code></a>.<code>roname</code>)</p>
<p>external node identifier</p></td>
</tr>
<tr>
<td><p><code>remote_lsn</code> <code>pg_lsn</code></p>
<p>The origin node's LSN up to which data has been replicated.</p></td>
</tr>
<tr>
<td><p><code>local_lsn</code> <code>pg_lsn</code></p>
<p>This node's LSN at which <code>remote_lsn</code> has been replicated. Used to flush commit records before persisting data to disk when using asynchronous commits.</p></td>
</tr>
</tbody>
</table>
