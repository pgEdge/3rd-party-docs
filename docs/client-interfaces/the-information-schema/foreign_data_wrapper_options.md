<a id="infoschema-foreign-data-wrapper-options"></a>

## `foreign_data_wrapper_options`


 The view `foreign_data_wrapper_options` contains all the options defined for foreign-data wrappers in the current database. Only those foreign-data wrappers are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `foreign_data_wrapper_options` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>foreign_data_wrapper_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that the foreign-data wrapper is defined in (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_data_wrapper_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign-data wrapper</p></td>
</tr>
<tr>
<td><p><code>option_name</code> <code>sql_identifier</code></p>
<p>Name of an option</p></td>
</tr>
<tr>
<td><p><code>option_value</code> <code>character_data</code></p>
<p>Value of the option</p></td>
</tr>
</tbody>
</table>
