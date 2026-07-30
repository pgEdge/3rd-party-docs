<a id="catalog-pg-opfamily"></a>

## `pg_opfamily`


 The catalog `pg_opfamily` defines operator families. Each operator family is a collection of operators and associated support routines that implement the semantics specified for a particular index access method. Furthermore, the operators in a family are all “compatible”, in a way that is specified by the access method. The operator family concept allows cross-data-type operators to be used with indexes and to be reasoned about using knowledge of access method semantics.


 Operator families are described at length in [Interfacing Extensions to Indexes](../../server-programming/extending-sql/interfacing-extensions-to-indexes.md#xindex).


**Table: `pg_opfamily` Columns**

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
<td><p><code>opfmethod</code> <code>oid</code> (references <a href="pg_am.md#catalog-pg-am"><code>pg_am</code></a>.<code>oid</code>)</p>
<p>Index access method operator family is for</p></td>
</tr>
<tr>
<td><p><code>opfname</code> <code>name</code></p>
<p>Name of this operator family</p></td>
</tr>
<tr>
<td><p><code>opfnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>Namespace of this operator family</p></td>
</tr>
<tr>
<td><p><code>opfowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the operator family</p></td>
</tr>
</tbody>
</table>


 The majority of the information defining an operator family is not in its `pg_opfamily` row, but in the associated rows in [`pg_amop`](pg_amop.md#catalog-pg-amop), [`pg_amproc`](pg_amproc.md#catalog-pg-amproc), and [`pg_opclass`](pg_opclass.md#catalog-pg-opclass).
