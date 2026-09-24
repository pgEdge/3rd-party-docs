<a id="view-pg-available-extension-versions"></a>

## `pg_available_extension_versions`


 The `pg_available_extension_versions` view lists the specific extension versions that are available for installation. See also the [`pg_extension`](../system-catalogs/pg_extension.md#catalog-pg-extension) catalog, which shows the extensions currently installed.


**Table: `pg_available_extension_versions` Columns**

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
<td><p><code>version</code> <code>text</code></p>
<p>Version name</p></td>
</tr>
<tr>
<td><p><code>installed</code> <code>bool</code></p>
<p>True if this version of this extension is currently installed</p></td>
</tr>
<tr>
<td><p><code>superuser</code> <code>bool</code></p>
<p>True if only superusers are allowed to install this extension (but see <code>trusted</code>)</p></td>
</tr>
<tr>
<td><p><code>trusted</code> <code>bool</code></p>
<p>True if the extension can be installed by non-superusers with appropriate privileges</p></td>
</tr>
<tr>
<td><p><code>relocatable</code> <code>bool</code></p>
<p>True if extension can be relocated to another schema</p></td>
</tr>
<tr>
<td><p><code>schema</code> <code>name</code></p>
<p>Name of the schema that the extension must be installed into, or <code>NULL</code> if partially or fully relocatable</p></td>
</tr>
<tr>
<td><p><code>requires</code> <code>name[]</code></p>
<p>Names of prerequisite extensions, or <code>NULL</code> if none</p></td>
</tr>
<tr>
<td><p><code>location</code> <code>text</code></p>
<p>The location where the extension is installed. If it is in the standard system location, then the value will be <code>$system</code>, while if it is found in the path specified by the <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-extension-control-path"><code>extension_control_path</code></a> GUC then the full path will be shown. Only superusers can see this information.</p></td>
</tr>
<tr>
<td><p><code>comment</code> <code>text</code></p>
<p>Comment string from the extension's control file</p></td>
</tr>
</tbody>
</table>


 The `pg_available_extension_versions` view is read-only.
