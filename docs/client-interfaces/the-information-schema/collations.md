<a id="infoschema-collations"></a>

## `collations`


 The view `collations` contains the collations available in the current database.


**Table: `collations` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>collation_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the collation (always the current database)</p></td>
</tr>
<tr>
<td><p><code>collation_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the collation</p></td>
</tr>
<tr>
<td><p><code>collation_name</code> <code>sql_identifier</code></p>
<p>Name of the default collation</p></td>
</tr>
<tr>
<td><p><code>pad_attribute</code> <code>character_data</code></p>
<p>Always <code>NO PAD</code> (The alternative <code>PAD SPACE</code> is not supported by PostgreSQL.)</p></td>
</tr>
</tbody>
</table>
