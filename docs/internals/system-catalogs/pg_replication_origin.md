<a id="catalog-pg-replication-origin"></a>

## `pg_replication_origin`


 The `pg_replication_origin` catalog contains all replication origins created. For more on replication origins see [Replication Progress Tracking](../../server-programming/replication-progress-tracking.md#replication-origins).


 Unlike most system catalogs, `pg_replication_origin` is shared across all databases of a cluster: there is only one copy of `pg_replication_origin` per cluster, not one per database.


**Table: `pg_replication_origin` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>roident</code> <code>oid</code></p>
<p>A unique, cluster-wide identifier for the replication origin. Should never leave the system.</p></td>
</tr>
<tr>
<td><p><code>roname</code> <code>text</code></p>
<p>The external, user defined, name of a replication origin.</p></td>
</tr>
</tbody>
</table>
