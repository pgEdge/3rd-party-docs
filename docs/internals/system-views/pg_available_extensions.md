<a id="view-pg-available-extensions"></a>

## `pg_available_extensions`


 The `pg_available_extensions` view lists the extensions that are available for installation. See also the [`pg_extension`](../system-catalogs/pg_extension.md#catalog-pg-extension) catalog, which shows the extensions currently installed.


**Table: `pg_available_extensions` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>name</code> <code>name</code></p>
<p>Extension name</p></td>
</tr>
<tr>
<td><p><code>default_version</code> <code>text</code></p>
<p>Name of default version, or <code>NULL</code> if none is specified</p></td>
</tr>
<tr>
<td><p><code>installed_version</code> <code>text</code></p>
<p>Currently installed version of the extension, or <code>NULL</code> if not installed</p></td>
</tr>
<tr>
<td><p><code>comment</code> <code>text</code></p>
<p>Comment string from the extension's control file</p></td>
</tr>
</tbody>
</table>


 The `pg_available_extensions` view is read-only.
