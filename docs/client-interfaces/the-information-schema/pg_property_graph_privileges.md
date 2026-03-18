<a id="infoschema-pg-property-graph-privileges"></a>

## `pg_property_graph_privileges`


 The view `pg_property_graph_privileges` identifies all privileges granted on property graphs to a currently enabled role or by a currently enabled role. There is one row for each combination of property graph, grantor, and grantee.


**Table: `pg_property_graph_privileges` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>grantor</code> <code>sql_identifier</code></p>
<p>Name of the role that granted the privilege</p></td>
</tr>
<tr>
<td><p><code>grantee</code> <code>sql_identifier</code></p>
<p>Name of the role that the privilege was granted to</p></td>
</tr>
<tr>
<td><p><code>property_graph_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the property graph (always the current database)</p></td>
</tr>
<tr>
<td><p><code>property_graph_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the property graph</p></td>
</tr>
<tr>
<td><p><code>property_graph_name</code> <code>sql_identifier</code></p>
<p>Name of the property graph</p></td>
</tr>
<tr>
<td><p><code>privilege_type</code> <code>character_data</code></p>
<p>Type of the privilege: <code>SELECT</code> is the only privilege type applicable to property graphs.</p></td>
</tr>
<tr>
<td><p><code>is_grantable</code> <code>yes_or_no</code></p>
<p><code>YES</code> if the privilege is grantable, <code>NO</code> if not</p></td>
</tr>
</tbody>
</table>
