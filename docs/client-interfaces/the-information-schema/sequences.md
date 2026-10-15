<a id="infoschema-sequences"></a>

## `sequences`


 The view `sequences` contains all sequences defined in the current database. Only those sequences are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `sequences` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>sequence_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the sequence (always the current database)</p></td>
</tr>
<tr>
<td><p><code>sequence_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the sequence</p></td>
</tr>
<tr>
<td><p><code>sequence_name</code> <code>sql_identifier</code></p>
<p>Name of the sequence</p></td>
</tr>
<tr>
<td><p><code>data_type</code> <code>character_data</code></p>
<p>The data type of the sequence.</p></td>
</tr>
<tr>
<td><p><code>numeric_precision</code> <code>cardinal_number</code></p>
<p>This column contains the (declared or implicit) precision of the sequence data type (see above). The precision indicates the number of significant digits. It can be expressed in decimal (base 10) or binary (base 2) terms, as specified in the column <code>numeric_precision_radix</code>.</p></td>
</tr>
<tr>
<td><p><code>numeric_precision_radix</code> <code>cardinal_number</code></p>
<p>This column indicates in which base the values in the columns <code>numeric_precision</code> and <code>numeric_scale</code> are expressed. The value is either 2 or 10.</p></td>
</tr>
<tr>
<td><p><code>numeric_scale</code> <code>cardinal_number</code></p>
<p>This column contains the (declared or implicit) scale of the sequence data type (see above). The scale indicates the number of significant digits to the right of the decimal point. It can be expressed in decimal (base 10) or binary (base 2) terms, as specified in the column <code>numeric_precision_radix</code>.</p></td>
</tr>
<tr>
<td><p><code>start_value</code> <code>character_data</code></p>
<p>The start value of the sequence</p></td>
</tr>
<tr>
<td><p><code>minimum_value</code> <code>character_data</code></p>
<p>The minimum value of the sequence</p></td>
</tr>
<tr>
<td><p><code>maximum_value</code> <code>character_data</code></p>
<p>The maximum value of the sequence</p></td>
</tr>
<tr>
<td><p><code>increment</code> <code>character_data</code></p>
<p>The increment of the sequence</p></td>
</tr>
<tr>
<td><p><code>cycle_option</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the sequence cycles, else <code>NO</code></p></td>
</tr>
</tbody>
</table>


 Note that in accordance with the SQL standard, the start, minimum, maximum, and increment values are returned as character strings.
