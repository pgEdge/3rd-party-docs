<a id="view-pg-seclabels"></a>

## `pg_seclabels`


 The view `pg_seclabels` provides information about security labels. It as an easier-to-query version of the [`pg_seclabel`](../system-catalogs/pg_seclabel.md#catalog-pg-seclabel) catalog.


**Table: `pg_seclabels` Columns**

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
<td><p><code>classoid</code> <code>oid</code> (references <a href="../system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the system catalog this object appears in</p></td>
</tr>
<tr>
<td><p><code>objsubid</code> <code>int4</code></p>
<p>For a security label on a table column, this is the column number (the <code>objoid</code> and <code>classoid</code> refer to the table itself). For all other object types, this column is zero.</p></td>
</tr>
<tr>
<td><p><code>objtype</code> <code>text</code></p>
<p>The type of object to which this label applies, as text.</p></td>
</tr>
<tr>
<td><p><code>objnamespace</code> <code>oid</code> (references <a href="../system-catalogs/pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace for this object, if applicable; otherwise NULL.</p></td>
</tr>
<tr>
<td><p><code>objname</code> <code>text</code></p>
<p>The name of the object to which this label applies, as text.</p></td>
</tr>
<tr>
<td><p><code>provider</code> <code>text</code> (references <a href="../system-catalogs/pg_seclabel.md#catalog-pg-seclabel"><code>pg_seclabel</code></a>.<code>provider</code>)</p>
<p>The label provider associated with this label.</p></td>
</tr>
<tr>
<td><p><code>label</code> <code>text</code> (references <a href="../system-catalogs/pg_seclabel.md#catalog-pg-seclabel"><code>pg_seclabel</code></a>.<code>label</code>)</p>
<p>The security label applied to this object.</p></td>
</tr>
</tbody>
</table>
