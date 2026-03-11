<a id="infoschema-routine-routine-usage"></a>

## `routine_routine_usage`


 The view `routine_routine_usage` identifies all functions or procedures that are used by another (or the same) function or procedure, either in the SQL body or in parameter default expressions. (This only works for unquoted SQL bodies, not quoted bodies or functions in other languages.) An entry is included here only if the used function is owned by a currently enabled role. (There is no such restriction on the using function.)


 Note that the entries for both functions in the view refer to the “specific” name of the routine, even though the column names are used in a way that is inconsistent with other information schema views about routines. This is per SQL standard, although it is arguably a misdesign. See [`routines`](routines.md#infoschema-routines) for more information about specific names.


**Table: `routine_routine_usage` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>specific_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database containing the using function (always the current database)</p></td>
</tr>
<tr>
<td><p><code>specific_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema containing the using function</p></td>
</tr>
<tr>
<td><p><code>specific_name</code> <code>sql_identifier</code></p>
<p>The “specific name” of the using function.</p></td>
</tr>
<tr>
<td><p><code>routine_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the function that is used by the first function (always the current database)</p></td>
</tr>
<tr>
<td><p><code>routine_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the function that is used by the first function</p></td>
</tr>
<tr>
<td><p><code>routine_name</code> <code>sql_identifier</code></p>
<p>The “specific name” of the function that is used by the first function.</p></td>
</tr>
</tbody>
</table>
