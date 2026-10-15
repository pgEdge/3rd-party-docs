<a id="infoschema-column-privileges"></a>

## `column_privileges`


 The view `column_privileges` identifies all privileges granted on columns to a currently enabled role or by a currently enabled role. There is one row for each combination of column, grantor, and grantee.


 If a privilege has been granted on an entire table, it will show up in this view as a grant for each column, but only for the privilege types where column granularity is possible: `SELECT`, `INSERT`, `UPDATE`, `REFERENCES`.


**Table: `column_privileges` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>grantor</code> <code>sql_identifier</code></p>
<p>Name of the role that granted the privilege</p></td>
</tr>
<tr>
<td><p><code>grantee</code> <code>sql_identifier</code></p>
<p>Name of the role that the privilege was granted to</p></td>
</tr>
<tr>
<td><p><code>table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the table that contains the column (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the table that contains the column</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the table that contains the column</p></td>
</tr>
<tr>
<td><p><code>column_name</code> <code>sql_identifier</code></p>
<p>Name of the column</p></td>
</tr>
<tr>
<td><p><code>privilege_type</code> <code>character_data</code></p>
<p>Type of the privilege: <code>SELECT</code>, <code>INSERT</code>, <code>UPDATE</code>, or <code>REFERENCES</code></p></td>
</tr>
<tr>
<td><p><code>is_grantable</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the privilege is grantable, <code>NO</code> if not</p></td>
</tr>
</tbody>
</table>
