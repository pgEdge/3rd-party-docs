<a id="catalog-pg-publication"></a>

## `pg_publication`


 The catalog `pg_publication` contains all publications created in the database. For more on publications see [Publication](../../server-administration/logical-replication/publication.md#logical-replication-publication).


**Table: `pg_publication` Columns**

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
<td><p><code>pubname</code> <code>name</code></p>
<p>Name of the publication</p></td>
</tr>
<tr>
<td><p><code>pubowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the publication</p></td>
</tr>
<tr>
<td><p><code>puballtables</code> <code>bool</code></p>
<p>If true, this publication automatically includes all tables in the database, including any that will be created in the future.</p></td>
</tr>
<tr>
<td><p><code>puballsequences</code> <code>bool</code></p>
<p>If true, this publication automatically includes all sequences in the database, including any that will be created in the future.</p></td>
</tr>
<tr>
<td><p><code>pubinsert</code> <code>bool</code></p>
<p>If true, <a href="../../reference/sql-commands/insert.md#sql-insert">sql-insert</a> operations are replicated for tables in the publication.</p></td>
</tr>
<tr>
<td><p><code>pubupdate</code> <code>bool</code></p>
<p>If true, <a href="../../reference/sql-commands/update.md#sql-update">sql-update</a> operations are replicated for tables in the publication.</p></td>
</tr>
<tr>
<td><p><code>pubdelete</code> <code>bool</code></p>
<p>If true, <a href="../../reference/sql-commands/delete.md#sql-delete">sql-delete</a> operations are replicated for tables in the publication.</p></td>
</tr>
<tr>
<td><p><code>pubtruncate</code> <code>bool</code></p>
<p>If true, <a href="../../reference/sql-commands/truncate.md#sql-truncate">sql-truncate</a> operations are replicated for tables in the publication.</p></td>
</tr>
<tr>
<td><p><code>pubviaroot</code> <code>bool</code></p>
<p>If true, operations on a leaf partition are replicated using the identity and schema of its topmost partitioned ancestor mentioned in the publication instead of its own.</p></td>
</tr>
<tr>
<td><p><code>pubgencols</code> <code>char</code></p>
<p>Controls how to handle generated column replication when there is no publication column list: <code>n</code> = generated columns in the tables associated with the publication should not be replicated, <code>s</code> = stored generated columns in the tables associated with the publication should be replicated.</p></td>
</tr>
</tbody>
</table>
