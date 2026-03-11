<a id="catalog-pg-amop"></a>

## `pg_amop`


 The catalog `pg_amop` stores information about operators associated with access method operator families. There is one row for each operator that is a member of an operator family. A family member can be either a *search* operator or an *ordering* operator. An operator can appear in more than one family, but cannot appear in more than one search position nor more than one ordering position within a family. (It is allowed, though unlikely, for an operator to be used for both search and ordering purposes.)


**Table: `pg_amop` Columns**

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
<td><p><code>amopfamily</code> <code>oid</code> (references <a href="pg_opfamily.md#catalog-pg-opfamily"><code>pg_opfamily</code></a>.<code>oid</code>)</p>
<p>The operator family this entry is for</p></td>
</tr>
<tr>
<td><p><code>amoplefttype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Left-hand input data type of operator</p></td>
</tr>
<tr>
<td><p><code>amoprighttype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Right-hand input data type of operator</p></td>
</tr>
<tr>
<td><p><code>amopstrategy</code> <code>int2</code></p>
<p>Operator strategy number</p></td>
</tr>
<tr>
<td><p><code>amoppurpose</code> <code>char</code></p>
<p>Operator purpose, either <code>s</code> for search or <code>o</code> for ordering</p></td>
</tr>
<tr>
<td><p><code>amopopr</code> <code>oid</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>OID of the operator</p></td>
</tr>
<tr>
<td><p><code>amopmethod</code> <code>oid</code> (references <a href="pg_am.md#catalog-pg-am"><code>pg_am</code></a>.<code>oid</code>)</p>
<p>Index access method operator family is for</p></td>
</tr>
<tr>
<td><p><code>amopsortfamily</code> <code>oid</code> (references <a href="pg_opfamily.md#catalog-pg-opfamily"><code>pg_opfamily</code></a>.<code>oid</code>)</p>
<p>The B-tree operator family this entry sorts according to, if an ordering operator; zero if a search operator</p></td>
</tr>
</tbody>
</table>


 A “search” operator entry indicates that an index of this operator family can be searched to find all rows satisfying `WHERE` *indexed_column* *operator* *constant*. Obviously, such an operator must return `boolean`, and its left-hand input type must match the index's column data type.


 An “ordering” operator entry indicates that an index of this operator family can be scanned to return rows in the order represented by `ORDER BY` *indexed_column* *operator* *constant*. Such an operator could return any sortable data type, though again its left-hand input type must match the index's column data type. The exact semantics of the `ORDER BY` are specified by the `amopsortfamily` column, which must reference a B-tree operator family for the operator's result type.


!!! note

    At present, it's assumed that the sort order for an ordering operator is the default for the referenced operator family, i.e., `ASC NULLS LAST`. This might someday be relaxed by adding additional columns to specify sort options explicitly.


 An entry's `amopmethod` must match the `opfmethod` of its containing operator family (including `amopmethod` here is an intentional denormalization of the catalog structure for performance reasons). Also, `amoplefttype` and `amoprighttype` must match the `oprleft` and `oprright` fields of the referenced [`pg_operator`](pg_operator.md#catalog-pg-operator) entry.
