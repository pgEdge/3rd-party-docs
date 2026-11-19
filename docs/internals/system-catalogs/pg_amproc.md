<a id="catalog-pg-amproc"></a>

## `pg_amproc`


 The catalog `pg_amproc` stores information about support functions associated with access method operator families. There is one row for each support function belonging to an operator family.


**Table: `pg_amproc` Columns**

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
<td><p><code>amprocfamily</code> <code>oid</code> (references <a href="pg_opfamily.md#catalog-pg-opfamily"><code>pg_opfamily</code></a>.<code>oid</code>)</p>
<p>The operator family this entry is for</p></td>
</tr>
<tr>
<td><p><code>amproclefttype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Left-hand input data type of associated operator</p></td>
</tr>
<tr>
<td><p><code>amprocrighttype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Right-hand input data type of associated operator</p></td>
</tr>
<tr>
<td><p><code>amprocnum</code> <code>int2</code></p>
<p>Support function number</p></td>
</tr>
<tr>
<td><p><code>amproc</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the function</p></td>
</tr>
</tbody>
</table>


 The usual interpretation of the `amproclefttype` and `amprocrighttype` fields is that they identify the left and right input types of the operator(s) that a particular support function supports. For some access methods these match the input data type(s) of the support function itself, for others not. There is a notion of “default” support functions for an index, which are those with `amproclefttype` and `amprocrighttype` both equal to the index operator class's `opcintype`.
