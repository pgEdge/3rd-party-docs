<a id="catalog-pg-opclass"></a>

## `pg_opclass`


 The catalog `pg_opclass` defines index access method operator classes. Each operator class defines semantics for index columns of a particular data type and a particular index access method. An operator class essentially specifies that a particular operator family is applicable to a particular indexable column data type. The set of operators from the family that are actually usable with the indexed column are whichever ones accept the column's data type as their left-hand input.


 Operator classes are described at length in [Interfacing Extensions to Indexes](../../server-programming/extending-sql/interfacing-extensions-to-indexes.md#xindex).


**Table: `pg_opclass` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>oid</code> <code>oid</code></p>
<p>Row identifier</p></td>
</tr>
<tr>
<td><p><code>opcmethod</code> <code>oid</code> (references <a href="pg_am.md#catalog-pg-am"><code>pg_am</code></a>.<code>oid</code>)</p>
<p>Index access method operator class is for</p></td>
</tr>
<tr>
<td><p><code>opcname</code> <code>name</code></p>
<p>Name of this operator class</p></td>
</tr>
<tr>
<td><p><code>opcnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>Namespace of this operator class</p></td>
</tr>
<tr>
<td><p><code>opcowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the operator class</p></td>
</tr>
<tr>
<td><p><code>opcfamily</code> <code>oid</code> (references <a href="pg_opfamily.md#catalog-pg-opfamily"><code>pg_opfamily</code></a>.<code>oid</code>)</p>
<p>Operator family containing the operator class</p></td>
</tr>
<tr>
<td><p><code>opcintype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type that the operator class indexes</p></td>
</tr>
<tr>
<td><p><code>opcdefault</code> <code>bool</code></p>
<p>True if this operator class is the default for <code>opcintype</code></p></td>
</tr>
<tr>
<td><p><code>opckeytype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Type of data stored in index, or zero if same as <code>opcintype</code></p></td>
</tr>
</tbody>
</table>


 An operator class's `opcmethod` must match the `opfmethod` of its containing operator family. Also, there must be no more than one `pg_opclass` row having `opcdefault` true for any given combination of `opcmethod` and `opcintype`.
