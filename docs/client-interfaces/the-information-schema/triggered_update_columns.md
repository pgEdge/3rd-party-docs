<a id="infoschema-triggered-update-columns"></a>

## `triggered_update_columns`


 For triggers in the current database that specify a column list (like `UPDATE OF column1, column2`), the view `triggered_update_columns` identifies these columns. Triggers that do not specify a column list are not included in this view. Only those columns are shown that the current user owns or has some privilege other than `SELECT` on.


**Table: `triggered_update_columns` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>trigger_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the trigger (always the current database)</p></td>
</tr>
<tr>
<td><p><code>trigger_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the trigger</p></td>
</tr>
<tr>
<td><p><code>trigger_name</code> <code>sql_identifier</code></p>
<p>Name of the trigger</p></td>
</tr>
<tr>
<td><p><code>event_object_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the table that the trigger is defined on (always the current database)</p></td>
</tr>
<tr>
<td><p><code>event_object_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the table that the trigger is defined on</p></td>
</tr>
<tr>
<td><p><code>event_object_table</code> <code>sql_identifier</code></p>
<p>Name of the table that the trigger is defined on</p></td>
</tr>
<tr>
<td><p><code>event_object_column</code> <code>sql_identifier</code></p>
<p>Name of the column that the trigger is defined on</p></td>
</tr>
</tbody>
</table>
