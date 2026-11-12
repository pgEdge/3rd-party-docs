<a id="infoschema-table-constraints"></a>

## `table_constraints`


 The view `table_constraints` contains all constraints belonging to tables that the current user owns or has some privilege other than `SELECT` on.


**Table: `table_constraints` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>constraint_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the constraint (always the current database)</p></td>
</tr>
<tr>
<td><p><code>constraint_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the constraint</p></td>
</tr>
<tr>
<td><p><code>constraint_name</code> <code>sql_identifier</code></p>
<p>Name of the constraint</p></td>
</tr>
<tr>
<td><p><code>table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the table (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the table</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the table</p></td>
</tr>
<tr>
<td><p><code>constraint_type</code> <code>character_data</code></p>
<p>Type of the constraint: <code>CHECK</code> (includes not-null constraints), <code>FOREIGN KEY</code>, <code>PRIMARY KEY</code>, or <code>UNIQUE</code></p></td>
</tr>
<tr>
<td><p><code>is_deferrable</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the constraint is deferrable, <code>NO</code> if not</p></td>
</tr>
<tr>
<td><p><code>initially_deferred</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the constraint is deferrable and initially deferred, <code>NO</code> if not</p></td>
</tr>
<tr>
<td><p><code>enforced</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the constraint is enforced, <code>NO</code> if not</p></td>
</tr>
<tr>
<td><p><code>nulls_distinct</code> <code>yes_or_no</code></p>
<p>If the constraint is a unique constraint, then <code>YES</code> if the constraint treats nulls as distinct or <code>NO</code> if it treats nulls as not distinct, otherwise null for other types of constraints.</p></td>
</tr>
</tbody>
</table>
