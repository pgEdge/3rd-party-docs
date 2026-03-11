<a id="infoschema-sql-features"></a>

## `sql_features`


 The table `sql_features` contains information about which formal features defined in the SQL standard are supported by PostgreSQL. This is the same information that is presented in [SQL Conformance](../../appendixes/sql-conformance/index.md#features). There you can also find some additional background information.


**Table: `sql_features` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>feature_id</code> <code>character_data</code></p>
<p>Identifier string of the feature</p></td>
</tr>
<tr>
<td><p><code>feature_name</code> <code>character_data</code></p>
<p>Descriptive name of the feature</p></td>
</tr>
<tr>
<td><p><code>sub_feature_id</code> <code>character_data</code></p>
<p>Identifier string of the subfeature, or a zero-length string if not a subfeature</p></td>
</tr>
<tr>
<td><p><code>sub_feature_name</code> <code>character_data</code></p>
<p>Descriptive name of the subfeature, or a zero-length string if not a subfeature</p></td>
</tr>
<tr>
<td><p><code>is_supported</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the feature is fully supported by the current version of PostgreSQL, <code>NO</code> if not</p></td>
</tr>
<tr>
<td><p><code>is_verified_by</code> <code>character_data</code></p>
<p>Always null, since the PostgreSQL development group does not perform formal testing of feature conformance</p></td>
</tr>
<tr>
<td><p><code>comments</code> <code>character_data</code></p>
<p>Possibly a comment about the supported status of the feature</p></td>
</tr>
</tbody>
</table>
