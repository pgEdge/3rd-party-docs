<a id="infoschema-foreign-data-wrappers"></a>

## `foreign_data_wrappers`


 The view `foreign_data_wrappers` contains all foreign-data wrappers defined in the current database. Only those foreign-data wrappers are shown that the current user has access to (by way of being the owner or having some privilege).


**Table: `foreign_data_wrappers` Columns**

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
<p>Name of the database that contains the foreign-data wrapper (always the current database)</p></td>
</tr>
<tr>
<td><p><code>foreign_data_wrapper_name</code> <code>sql_identifier</code></p>
<p>Name of the foreign-data wrapper</p></td>
</tr>
<tr>
<td><p><code>authorization_identifier</code> <code>sql_identifier</code></p>
<p>Name of the owner of the foreign server</p></td>
</tr>
<tr>
<td><p><code>library_name</code> <code>character_data</code></p>
<p>File name of the library that implementing this foreign-data wrapper</p></td>
</tr>
<tr>
<td><p><code>foreign_data_wrapper_language</code> <code>character_data</code></p>
<p>Language used to implement this foreign-data wrapper</p></td>
</tr>
</tbody>
</table>
