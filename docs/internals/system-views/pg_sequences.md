<a id="view-pg-sequences"></a>

## `pg_sequences`


 The view `pg_sequences` provides access to useful information about each sequence in the database.


**Table: `pg_sequences` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>schemaname</code> <code>name</code> (references <a href="../system-catalogs/pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>nspname</code>)</p>
<p>Name of schema containing sequence</p></td>
</tr>
<tr>
<td><p><code>sequencename</code> <code>name</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relname</code>)</p>
<p>Name of sequence</p></td>
</tr>
<tr>
<td><p><code>sequenceowner</code> <code>name</code> (references <a href="../system-catalogs/pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>rolname</code>)</p>
<p>Name of sequence's owner</p></td>
</tr>
<tr>
<td><p><code>data_type</code> <code>regtype</code> (references <a href="../system-catalogs/pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type of the sequence</p></td>
</tr>
<tr>
<td><p><code>start_value</code> <code>int8</code></p>
<p>Start value of the sequence</p></td>
</tr>
<tr>
<td><p><code>min_value</code> <code>int8</code></p>
<p>Minimum value of the sequence</p></td>
</tr>
<tr>
<td><p><code>max_value</code> <code>int8</code></p>
<p>Maximum value of the sequence</p></td>
</tr>
<tr>
<td><p><code>increment_by</code> <code>int8</code></p>
<p>Increment value of the sequence</p></td>
</tr>
<tr>
<td><p><code>cycle</code> <code>bool</code></p>
<p>Whether the sequence cycles</p></td>
</tr>
<tr>
<td><p><code>cache_size</code> <code>int8</code></p>
<p>Cache size of the sequence</p></td>
</tr>
<tr>
<td><p><code>last_value</code> <code>int8</code></p>
<p>The last sequence value written to disk. If caching is used, this value can be greater than the last value handed out from the sequence.</p></td>
</tr>
</tbody>
</table>


 The `last_value` column will read as null if any of the following are true:

-  The sequence has not been read from yet.
-  The current user does not have `USAGE` or `SELECT` privilege on the sequence.
-  The sequence is unlogged and the server is a standby.
