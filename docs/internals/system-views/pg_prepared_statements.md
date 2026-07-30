<a id="view-pg-prepared-statements"></a>

## `pg_prepared_statements`


 The `pg_prepared_statements` view displays all the prepared statements that are available in the current session. See [sql-prepare](../../reference/sql-commands/prepare.md#sql-prepare) for more information about prepared statements.


 `pg_prepared_statements` contains one row for each prepared statement. Rows are added to the view when a new prepared statement is created and removed when a prepared statement is released (for example, via the [`DEALLOCATE`](../../reference/sql-commands/deallocate.md#sql-deallocate) command).


**Table: `pg_prepared_statements` Columns**

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
<p>The identifier of the prepared statement</p></td>
</tr>
<tr>
<td><p><code>statement</code> <code>text</code></p>
<p>The query string submitted by the client to create this prepared statement. For prepared statements created via SQL, this is the <code>PREPARE</code> statement submitted by the client. For prepared statements created via the frontend/backend protocol, this is the text of the prepared statement itself.</p></td>
</tr>
<tr>
<td><p><code>prepare_time</code> <code>timestamptz</code></p>
<p>The time at which the prepared statement was created</p></td>
</tr>
<tr>
<td><p><code>parameter_types</code> <code>regtype[]</code></p>
<p>The expected parameter types for the prepared statement in the form of an array of <code>regtype</code>. The OID corresponding to an element of this array can be obtained by casting the <code>regtype</code> value to <code>oid</code>.</p></td>
</tr>
<tr>
<td><p><code>result_types</code> <code>regtype[]</code></p>
<p>The types of the columns returned by the prepared statement in the form of an array of <code>regtype</code>. The OID corresponding to an element of this array can be obtained by casting the <code>regtype</code> value to <code>oid</code>. If the prepared statement does not provide a result (e.g., a DML statement), then this field will be null.</p></td>
</tr>
<tr>
<td><p><code>from_sql</code> <code>bool</code></p>
<p><code>true</code> if the prepared statement was created via the <code>PREPARE</code> SQL command; <code>false</code> if the statement was prepared via the frontend/backend protocol</p></td>
</tr>
<tr>
<td><p><code>generic_plans</code> <code>int8</code></p>
<p>Number of times generic plan was chosen</p></td>
</tr>
<tr>
<td><p><code>custom_plans</code> <code>int8</code></p>
<p>Number of times custom plan was chosen</p></td>
</tr>
</tbody>
</table>


 The `pg_prepared_statements` view is read-only.
