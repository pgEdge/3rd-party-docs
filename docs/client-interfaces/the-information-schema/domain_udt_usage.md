<a id="infoschema-domain-udt-usage"></a>

## `domain_udt_usage`


 The view `domain_udt_usage` identifies all domains that are based on data types owned by a currently enabled role. Note that in PostgreSQL, built-in data types behave like user-defined types, so they are included here as well.


**Table: `domain_udt_usage` Columns**

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
<p>Name of the database that the domain data type is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>udt_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that the domain data type is defined in</p></td>
</tr>
<tr>
<td><p><code>udt_name</code> <code>sql_identifier</code></p>
<p>Name of the domain data type</p></td>
</tr>
<tr>
<td><p><code>domain_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the domain (always the current database)</p></td>
</tr>
<tr>
<td><p><code>domain_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the domain</p></td>
</tr>
<tr>
<td><p><code>domain_name</code> <code>sql_identifier</code></p>
<p>Name of the domain</p></td>
</tr>
</tbody>
</table>
