<a id="infoschema-tables"></a>

## `tables`


 The view `tables` contains all tables and views defined in the current database. Only those tables and views are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `tables` Columns**

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
<td><p><code>table_type</code> <code>character_data</code></p>
<p>Type of the table: <code>BASE TABLE</code> for a persistent base table (the normal table type), <code>VIEW</code> for a view, <code>FOREIGN</code> for a foreign table, or <code>LOCAL TEMPORARY</code> for a temporary table</p></td>
</tr>
<tr>
<td><p><code>self_referencing_column_name</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>reference_generation</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>user_defined_type_catalog</code> <code>sql_identifier</code></p>
<p>If the table is a typed table, the name of the database that contains the underlying data type (always the current database), else null.</p></td>
</tr>
<tr>
<td><p><code>user_defined_type_schema</code> <code>sql_identifier</code></p>
<p>If the table is a typed table, the name of the schema that contains the underlying data type, else null.</p></td>
</tr>
<tr>
<td><p><code>user_defined_type_name</code> <code>sql_identifier</code></p>
<p>If the table is a typed table, the name of the underlying data type, else null.</p></td>
</tr>
<tr>
<td><p><code>is_insertable_into</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the table is insertable into, <code>NO</code> if not (Base tables are always insertable into, views not necessarily.)</p></td>
</tr>
<tr>
<td><p><code>is_typed</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the table is a typed table, <code>NO</code> if not</p></td>
</tr>
<tr>
<td><p><code>commit_action</code> <code>character_data</code></p>
<p>Not yet implemented</p></td>
</tr>
</tbody>
</table>
