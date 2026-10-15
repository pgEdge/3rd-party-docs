<a id="infoschema-referential-constraints"></a>

## `referential_constraints`


 The view `referential_constraints` contains all referential (foreign key) constraints in the current database. Only those constraints are shown for which the current user has write access to the referencing table (by way of being the owner or having some privilege other than `SELECT`).


**Table: `referential_constraints` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>constraint_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the constraint (always the current database)</p></td>
</tr>
<tr>
<td><p><code>constraint_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the constraint</p></td>
</tr>
<tr>
<td><p><code>constraint_name</code> <code>sql_identifier</code></p>
<p>Name of the constraint</p></td>
</tr>
<tr>
<td><p><code>unique_constraint_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the unique or primary key constraint that the foreign key constraint references (always the current database)</p></td>
</tr>
<tr>
<td><p><code>unique_constraint_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the unique or primary key constraint that the foreign key constraint references</p></td>
</tr>
<tr>
<td><p><code>unique_constraint_name</code> <code>sql_identifier</code></p>
<p>Name of the unique or primary key constraint that the foreign key constraint references</p></td>
</tr>
<tr>
<td><p><code>match_option</code> <code>character_data</code></p>
<p>Match option of the foreign key constraint: <code>FULL</code>, <code>PARTIAL</code>, or <code>NONE</code>.</p></td>
</tr>
<tr>
<td><p><code>update_rule</code> <code>character_data</code></p>
<p>Update rule of the foreign key constraint: <code>CASCADE</code>, <code>SET NULL</code>, <code>SET DEFAULT</code>, <code>RESTRICT</code>, or <code>NO ACTION</code>.</p></td>
</tr>
<tr>
<td><p><code>delete_rule</code> <code>character_data</code></p>
<p>Delete rule of the foreign key constraint: <code>CASCADE</code>, <code>SET NULL</code>, <code>SET DEFAULT</code>, <code>RESTRICT</code>, or <code>NO ACTION</code>.</p></td>
</tr>
</tbody>
</table>
