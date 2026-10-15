<a id="catalog-pg-seclabel"></a>

## `pg_seclabel`


 The catalog `pg_seclabel` stores security labels on database objects. Security labels can be manipulated with the [`SECURITY LABEL`](../../reference/sql-commands/security-label.md#sql-security-label) command. For an easier way to view security labels, see [`pg_seclabels`](../system-views/pg_seclabels.md#view-pg-seclabels).


 See also [`pg_shseclabel`](pg_shseclabel.md#catalog-pg-shseclabel), which performs a similar function for security labels of database objects that are shared across a database cluster.


**Table: `pg_seclabel` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>objoid</code> <code>oid</code> (references any OID column)</p>
<p>The OID of the object this security label pertains to</p></td>
</tr>
<tr>
<td><p><code>classoid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the system catalog this object appears in</p></td>
</tr>
<tr>
<td><p><code>objsubid</code> <code>int4</code></p>
<p>For a security label on a table column, this is the column number (the <code>objoid</code> and <code>classoid</code> refer to the table itself). For all other object types, this column is zero.</p></td>
</tr>
<tr>
<td><p><code>provider</code> <code>text</code></p>
<p>The label provider associated with this label.</p></td>
</tr>
<tr>
<td><p><code>label</code> <code>text</code></p>
<p>The security label applied to this object.</p></td>
</tr>
</tbody>
</table>
