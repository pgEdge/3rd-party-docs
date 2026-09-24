<a id="infoschema-collation-character-set-applicab"></a>

## `collation_character_set_​applicability`


 The view `collation_character_set_applicability` identifies which character set the available collations are applicable to. In PostgreSQL, there is only one character set per database (see explanation in [`character_sets`](character_sets.md#infoschema-character-sets)), so this view does not provide much useful information.


**Table: `collation_character_set_applicability` Columns**

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
<td><p><code>character_set_catalog</code> <code>sql_identifier</code></p>
<p>Character sets are currently not implemented as schema objects, so this column is null</p></td>
</tr>
<tr>
<td><p><code>character_set_schema</code> <code>sql_identifier</code></p>
<p>Character sets are currently not implemented as schema objects, so this column is null</p></td>
</tr>
<tr>
<td><p><code>character_set_name</code> <code>sql_identifier</code></p>
<p>Name of the character set</p></td>
</tr>
</tbody>
</table>
