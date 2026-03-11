<a id="infoschema-check-constraints"></a>

## `check_constraints`


 The view `check_constraints` contains all check constraints, either defined on a table or on a domain, that are owned by a currently enabled role. (The owner of the table or domain is the owner of the constraint.)


 The SQL standard considers not-null constraints to be check constraints with a <code>CHECK (</code><em>column_name</em><code> IS NOT
   NULL)</code> expression. So not-null constraints are also included here and don't have a separate view.


**Table: `check_constraints` Columns**

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
<td><p><code>check_clause</code> <code>character_data</code></p>
<p>The check expression of the check constraint</p></td>
</tr>
</tbody>
</table>
