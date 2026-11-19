<a id="infoschema-key-column-usage"></a>

## `key_column_usage`


 The view `key_column_usage` identifies all columns in the current database that are restricted by some unique, primary key, or foreign key constraint. Check constraints are not included in this view. Only those columns are shown that the current user has access to, by way of being the owner or having some privilege.


**Table: `key_column_usage` Columns**

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
<p>Name of the database that contains the table that contains the column that is restricted by this constraint (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the table that contains the column that is restricted by this constraint</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the table that contains the column that is restricted by this constraint</p></td>
</tr>
<tr>
<td><p><code>column_name</code> <code>sql_identifier</code></p>
<p>Name of the column that is restricted by this constraint</p></td>
</tr>
<tr>
<td><p><code>ordinal_position</code> <code>cardinal_number</code></p>
<p>Ordinal position of the column within the constraint key (count starts at 1)</p></td>
</tr>
<tr>
<td><p><code>position_in_unique_constraint</code> <code>cardinal_number</code></p>
<p>For a foreign-key constraint, ordinal position of the referenced column within its unique constraint (count starts at 1); otherwise null</p></td>
</tr>
</tbody>
</table>
