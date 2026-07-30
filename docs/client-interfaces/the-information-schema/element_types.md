<a id="infoschema-element-types"></a>

## `element_types`


 The view `element_types` contains the data type descriptors of the elements of arrays. When a table column, composite-type attribute, domain, function parameter, or function return value is defined to be of an array type, the respective information schema view only contains `ARRAY` in the column `data_type`. To obtain information on the element type of the array, you can join the respective view with this view. For example, to show the columns of a table with data types and array element types, if applicable, you could do:

```sql

SELECT c.column_name, c.data_type, e.data_type AS element_type
FROM information_schema.columns c LEFT JOIN information_schema.element_types e
     ON ((c.table_catalog, c.table_schema, c.table_name, 'TABLE', c.dtd_identifier)
       = (e.object_catalog, e.object_schema, e.object_name, e.object_type, e.collection_type_identifier))
WHERE c.table_schema = '...' AND c.table_name = '...'
ORDER BY c.ordinal_position;
```
 This view only includes objects that the current user has access to, by way of being the owner or having some privilege.


**Table: `element_types` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>object_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the object that uses the array being described (always the current database)</p></td>
</tr>
<tr>
<td><p><code>object_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the object that uses the array being described</p></td>
</tr>
<tr>
<td><p><code>object_name</code> <code>sql_identifier</code></p>
<p>Name of the object that uses the array being described</p></td>
</tr>
<tr>
<td><p><code>object_type</code> <code>character_data</code></p>
<p>The type of the object that uses the array being described: one of <code>TABLE</code> (the array is used by a column of that table), <code>USER-DEFINED TYPE</code> (the array is used by an attribute of that composite type), <code>DOMAIN</code> (the array is used by that domain), <code>ROUTINE</code> (the array is used by a parameter or the return data type of that function).</p></td>
</tr>
<tr>
<td><p><code>collection_type_identifier</code> <code>sql_identifier</code></p>
<p>The identifier of the data type descriptor of the array being described. Use this to join with the <code>dtd_identifier</code> columns of other information schema views.</p></td>
</tr>
<tr>
<td><p><code>data_type</code> <code>character_data</code></p>
<p>Data type of the array elements, if it is a built-in type, else <code>USER-DEFINED</code> (in that case, the type is identified in <code>udt_name</code> and associated columns).</p></td>
</tr>
<tr>
<td><p><code>character_maximum_length</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>character_octet_length</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>character_set_catalog</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>character_set_schema</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>character_set_name</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>collation_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the collation of the element type (always the current database), null if default or the data type of the element is not collatable</p></td>
</tr>
<tr>
<td><p><code>collation_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the collation of the element type, null if default or the data type of the element is not collatable</p></td>
</tr>
<tr>
<td><p><code>collation_name</code> <code>sql_identifier</code></p>
<p>Name of the collation of the element type, null if default or the data type of the element is not collatable</p></td>
</tr>
<tr>
<td><p><code>numeric_precision</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>numeric_precision_radix</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>numeric_scale</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>datetime_precision</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>interval_type</code> <code>character_data</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>interval_precision</code> <code>cardinal_number</code></p>
<p>Always null, since this information is not applied to array element data types in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>udt_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the data type of the elements is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>udt_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that the data type of the elements is defined in</p></td>
</tr>
<tr>
<td><p><code>udt_name</code> <code>sql_identifier</code></p>
<p>Name of the data type of the elements</p></td>
</tr>
<tr>
<td><p><code>scope_catalog</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>scope_schema</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>scope_name</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>maximum_cardinality</code> <code>cardinal_number</code></p>
<p>Always null, because arrays always have unlimited maximum cardinality in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>dtd_identifier</code> <code>sql_identifier</code></p>
<p>An identifier of the data type descriptor of the element. This is currently not useful.</p></td>
</tr>
</tbody>
</table>
