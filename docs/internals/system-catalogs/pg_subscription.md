<a id="catalog-pg-subscription"></a>

## `pg_subscription`


 The catalog `pg_subscription` contains all existing logical replication subscriptions. For more information about logical replication see [Logical Replication](../../server-administration/logical-replication/index.md#logical-replication).


 Unlike most system catalogs, `pg_subscription` is shared across all databases of a cluster: there is only one copy of `pg_subscription` per cluster, not one per database.


 Access to the column `subconninfo` is revoked from normal users, because it could contain plain-text passwords.


**Table: `pg_subscription` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>oid</code> <code>oid</code></p>
<p>Row identifier</p></td>
</tr>
<tr>
<td><p><code>subdbid</code> <code>oid</code> (references <a href="pg_database.md#catalog-pg-database"><code>pg_database</code></a>.<code>oid</code>)</p>
<p>OID of the database that the subscription resides in</p></td>
</tr>
<tr>
<td><p><code>subskiplsn</code> <code>pg_lsn</code></p>
<p>Finish LSN of the transaction whose changes are to be skipped, if a valid LSN; otherwise <code>0/0000000</code>.</p></td>
</tr>
<tr>
<td><p><code>subname</code> <code>name</code></p>
<p>Name of the subscription</p></td>
</tr>
<tr>
<td><p><code>subowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the subscription</p></td>
</tr>
<tr>
<td><p><code>subenabled</code> <code>bool</code></p>
<p>If true, the subscription is enabled and should be replicating</p></td>
</tr>
<tr>
<td><p><code>subbinary</code> <code>bool</code></p>
<p>If true, the subscription will request that the publisher send data in binary format</p></td>
</tr>
<tr>
<td><p><code>substream</code> <code>char</code></p>
<p>Controls how to handle the streaming of in-progress transactions: <code>f</code> = disallow streaming of in-progress transactions, <code>t</code> = spill the changes of in-progress transactions to disk and apply at once after the transaction is committed on the publisher and received by the subscriber, <code>p</code> = apply changes directly using a parallel apply worker if available (same as <code>t</code> if no worker is available)</p></td>
</tr>
<tr>
<td><p><code>subtwophasestate</code> <code>char</code></p>
<p>State codes for two-phase mode: <code>d</code> = disabled, <code>p</code> = pending enablement, <code>e</code> = enabled</p></td>
</tr>
<tr>
<td><p><code>subdisableonerr</code> <code>bool</code></p>
<p>If true, the subscription will be disabled if one of its workers detects an error</p></td>
</tr>
<tr>
<td><p><code>subpasswordrequired</code> <code>bool</code></p>
<p>If true, the subscription will be required to specify a password for authentication</p></td>
</tr>
<tr>
<td><p><code>subrunasowner</code> <code>bool</code></p>
<p>If true, the subscription will be run with the permissions of the subscription owner</p></td>
</tr>
<tr>
<td><p><code>subfailover</code> <code>bool</code></p>
<p>If true, the associated replication slots (i.e. the main slot and the table synchronization slots) in the upstream database are enabled to be synchronized to the standbys</p></td>
</tr>
<tr>
<td><p><code>subretaindeadtuples</code> <code>bool</code></p>
<p>If true, the detection of <a href="../../server-administration/logical-replication/conflicts.md#conflict-update-deleted">update_deleted</a> is enabled and the information (e.g., dead tuples, commit timestamps, and origins) on the subscriber that is useful for conflict detection is retained.</p></td>
</tr>
<tr>
<td><p><code>submaxretention</code> <code>int4</code></p>
<p>The maximum duration (in milliseconds) for which information (e.g., dead tuples, commit timestamps, and origins) useful for conflict detection can be retained.</p></td>
</tr>
<tr>
<td><p><code>subretentionactive</code> <code>bool</code></p>
<p>The retention status of information (e.g., dead tuples, commit timestamps, and origins) useful for conflict detection. True if <a href="../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-retain-dead-tuples"><code>retain_dead_tuples</code></a> is enabled, and the retention duration has not exceeded <a href="../../reference/sql-commands/create-subscription.md#sql-createsubscription-params-with-max-retention-duration"><code>max_retention_duration</code></a>, when defined.</p></td>
</tr>
<tr>
<td><p><code>subserver</code> <code>oid</code> (references <a href="pg_foreign_server.md#catalog-pg-foreign-server"><code>pg_foreign_server</code></a>.<code>oid</code>)</p>
<p>Foreign server to use for the connection string. Zero if <code>subconninfo</code> is nonnull.</p></td>
</tr>
<tr>
<td><p><code>subconninfo</code> <code>text</code></p>
<p>Connection string to the upstream database. <code>NULL</code> if <code>subserver</code> is nonzero.</p></td>
</tr>
<tr>
<td><p><code>subslotname</code> <code>name</code></p>
<p>Name of the replication slot in the upstream database (also used for the local replication origin name); null represents <code>NONE</code></p></td>
</tr>
<tr>
<td><p><code>subsynccommit</code> <code>text</code></p>
<p>The <code>synchronous_commit</code> setting for the subscription's workers to use</p></td>
</tr>
<tr>
<td><p><code>subwalrcvtimeout</code> <code>text</code></p>
<p>The <code>wal_receiver_timeout</code> setting for the subscription's workers to use</p></td>
</tr>
<tr>
<td><p><code>subpublications</code> <code>text[]</code></p>
<p>Array of subscribed publication names. These reference publications defined in the upstream database. For more on publications see <a href="../../server-administration/logical-replication/publication.md#logical-replication-publication">Publication</a>.</p></td>
</tr>
<tr>
<td><p><code>suborigin</code> <code>text</code></p>
<p>The origin value must be either <code>none</code> or <code>any</code>. The default is <code>any</code>. If <code>none</code>, the subscription will request the publisher to only send changes that don't have an origin. If <code>any</code>, the publisher sends changes regardless of their origin.</p></td>
</tr>
</tbody>
</table>
