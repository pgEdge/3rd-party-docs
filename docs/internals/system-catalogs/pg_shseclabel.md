<a id="catalog-pg-shseclabel"></a>

## `pg_shseclabel`


 The catalog `pg_shseclabel` stores security labels on shared database objects. Security labels can be manipulated with the [`SECURITY LABEL`](../../reference/sql-commands/security-label.md#sql-security-label) command. For an easier way to view security labels, see [`pg_seclabels`](../system-views/pg_seclabels.md#view-pg-seclabels).


 See also [`pg_seclabel`](pg_seclabel.md#catalog-pg-seclabel), which performs a similar function for security labels involving objects within a single database.


 Unlike most system catalogs, `pg_shseclabel` is shared across all databases of a cluster: there is only one copy of `pg_shseclabel` per cluster, not one per database.


**Table: `pg_shseclabel` Columns**

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
<td><p><code>provider</code> <code>text</code></p>
<p>The label provider associated with this label.</p></td>
</tr>
<tr>
<td><p><code>label</code> <code>text</code></p>
<p>The security label applied to this object.</p></td>
</tr>
</tbody>
</table>
