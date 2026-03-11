<a id="infoschema-user-defined-types"></a>

## `user_defined_types`


 The view `user_defined_types` currently contains all composite types defined in the current database. Only those types are shown that the current user has access to (by way of being the owner or having some privilege).


 SQL knows about two kinds of user-defined types: structured types (also known as composite types in PostgreSQL) and distinct types (not implemented in PostgreSQL). To be future-proof, use the column `user_defined_type_category` to differentiate between these. Other user-defined types such as base types and enums, which are PostgreSQL extensions, are not shown here. For domains, see [`domains`](domains.md#infoschema-domains) instead.


**Table: `user_defined_types` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>user_defined_type_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the type (always the current database)</p></td>
</tr>
<tr>
<td><p><code>user_defined_type_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the type</p></td>
</tr>
<tr>
<td><p><code>user_defined_type_name</code> <code>sql_identifier</code></p>
<p>Name of the type</p></td>
</tr>
<tr>
<td><p><code>user_defined_type_category</code> <code>character_data</code></p>
<p>Currently always <code>STRUCTURED</code></p></td>
</tr>
<tr>
<td><p><code>is_instantiable</code> <code>yes_or_no</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>is_final</code> <code>yes_or_no</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>ordering_form</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>ordering_category</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>ordering_routine_catalog</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>ordering_routine_schema</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>ordering_routine_name</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>reference_type</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>data_type</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>character_maximum_length</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>character_octet_length</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
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
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>collation_schema</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>collation_name</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>numeric_precision</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>numeric_precision_radix</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>numeric_scale</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>datetime_precision</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>interval_type</code> <code>character_data</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>interval_precision</code> <code>cardinal_number</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>source_dtd_identifier</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>ref_dtd_identifier</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
</tbody>
</table>
