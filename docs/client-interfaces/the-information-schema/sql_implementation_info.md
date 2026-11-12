<a id="infoschema-sql-implementation-info"></a>

## `sql_implementation_info`


 The table `sql_implementation_info` contains information about various aspects that are left implementation-defined by the SQL standard. This information is primarily intended for use in the context of the ODBC interface; users of other interfaces will probably find this information to be of little use. For this reason, the individual implementation information items are not described here; you will find them in the description of the ODBC interface.


**Table: `sql_implementation_info` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>implementation_info_id</code> <code>character_data</code></p>
<p>Identifier string of the implementation information item</p></td>
</tr>
<tr>
<td><p><code>implementation_info_name</code> <code>character_data</code></p>
<p>Descriptive name of the implementation information item</p></td>
</tr>
<tr>
<td><p><code>integer_value</code> <code>cardinal_number</code></p>
<p>Value of the implementation information item, or null if the value is contained in the column <code>character_value</code></p></td>
</tr>
<tr>
<td><p><code>character_value</code> <code>character_data</code></p>
<p>Value of the implementation information item, or null if the value is contained in the column <code>integer_value</code></p></td>
</tr>
<tr>
<td><p><code>comments</code> <code>character_data</code></p>
<p>Possibly a comment pertaining to the implementation information item</p></td>
</tr>
</tbody>
</table>
