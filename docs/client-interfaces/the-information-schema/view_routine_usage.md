<a id="infoschema-view-routine-usage"></a>

## `view_routine_usage`


 The view `view_routine_usage` identifies all routines (functions and procedures) that are used in the query expression of a view (the `SELECT` statement that defines the view). A routine is only included if that routine is owned by a currently enabled role.


**Table: `view_routine_usage` Columns**

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
<p>Name of the database containing the view (always the current database)</p></td>
</tr>
<tr>
<td><p><code>table_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the view</p></td>
</tr>
<tr>
<td><p><code>table_name</code> <code>sql_identifier</code></p>
<p>Name of the view</p></td>
</tr>
<tr>
<td><p><code>specific_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the function (always the current database)</p></td>
</tr>
<tr>
<td><p><code>specific_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the function</p></td>
</tr>
<tr>
<td><p><code>specific_name</code> <code>sql_identifier</code></p>
<p>The “specific name” of the function. See <a href="routines.md#infoschema-routines"><code>routines</code></a> for more information.</p></td>
</tr>
</tbody>
</table>
