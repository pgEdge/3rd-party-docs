<a id="infoschema-column-domain-usage"></a>

## `column_domain_usage`


 The view `column_domain_usage` identifies all columns (of a table or a view) that make use of some domain defined in the current database and owned by a currently enabled role.


**Table: `column_domain_usage` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>domain_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the domain (always the current database)</p></td>
</tr>
<tr>
<td><p><code>domain_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the domain</p></td>
</tr>
<tr>
<td><p><code>domain_name</code> <code>sql_identifier</code></p>
<p>Name of the domain</p></td>
</tr>
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
<p>Name of the column</p></td>
</tr>
</tbody>
</table>
