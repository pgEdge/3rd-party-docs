<a id="infoschema-column-column-usage"></a>

## `column_column_usage`


 The view `column_column_usage` identifies all generated columns that depend on another base column in the same table. Only tables owned by a currently enabled role are included.


**Table: `column_column_usage` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the table (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the table</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the table</p></td>
</tr>
<tr>
<td><p><code>column_name</code> <code>sql_identifier</code></p>
<p>Name of the base column that a generated column depends on</p></td>
</tr>
<tr>
<td><p><code>dependent_column</code> <code>sql_identifier</code></p>
<p>Name of the generated column</p></td>
</tr>
</tbody>
</table>
