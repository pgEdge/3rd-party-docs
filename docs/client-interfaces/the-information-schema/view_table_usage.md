<a id="infoschema-view-table-usage"></a>

## `view_table_usage`


 The view `view_table_usage` identifies all tables that are used in the query expression of a view (the `SELECT` statement that defines the view). A table is only included if that table is owned by a currently enabled role.


!!! note

    System tables are not included. This should be fixed sometime.


**Table: `view_table_usage` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>view_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the view (always the current database)</p></td>
</tr>
<tr>
<td><p><code>view_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the view</p></td>
</tr>
<tr>
<td><p><code>view_name</code> <code>sql_identifier</code></p>
<p>Name of the view</p></td>
</tr>
<tr>
<td><p><code>table_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the table that is used by the view (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the table that is used by the view</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the table that is used by the view</p></td>
</tr>
</tbody>
</table>
