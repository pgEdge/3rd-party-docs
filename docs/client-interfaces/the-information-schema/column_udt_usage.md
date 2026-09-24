<a id="infoschema-column-udt-usage"></a>

## `column_udt_usage`


 The view `column_udt_usage` identifies all columns that use data types owned by a currently enabled role. Note that in PostgreSQL, built-in data types behave like user-defined types, so they are included here as well. See also [`columns`](columns.md#infoschema-columns) for details.


**Table: `column_udt_usage` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>udt_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the column data type (the underlying type of the domain, if applicable) is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>udt_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that the column data type (the underlying type of the domain, if applicable) is defined in</p></td>
</tr>
<tr>
<td><p><code>udt_name</code> <code>sql_identifier</code></p>
<p>Name of the column data type (the underlying type of the domain, if applicable)</p></td>
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
