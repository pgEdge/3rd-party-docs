<a id="infoschema-transforms"></a>

## `transforms`


 The view `transforms` contains information about the transforms defined in the current database. More precisely, it contains a row for each function contained in a transform (the “from SQL” or “to SQL” function).


**Table: `transforms` Columns**

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
<p>Name of the database that contains the type the transform is for (always the current database)</p></td>
</tr>
<tr>
<td><p><code>udt_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the type the transform is for</p></td>
</tr>
<tr>
<td><p><code>udt_name</code> <code>sql_identifier</code></p>
<p>Name of the type the transform is for</p></td>
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
<tr>
<td><p><code>group_name</code> <code>sql_identifier</code></p>
<p>The SQL standard allows defining transforms in “groups”, and selecting a group at run time. PostgreSQL does not support this. Instead, transforms are specific to a language. As a compromise, this field contains the language the transform is for.</p></td>
</tr>
<tr>
<td><p><code>transform_type</code> <code>character_data</code></p>
<p><code>FROM SQL</code> or <code>TO SQL</code></p></td>
</tr>
</tbody>
</table>
