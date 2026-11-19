<a id="catalog-pg-range"></a>

## `pg_range`


 The catalog `pg_range` stores information about range types. This is in addition to the types' entries in [`pg_type`](pg_type.md#catalog-pg-type).


**Table: `pg_range` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>rngtypid</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>OID of the range type</p></td>
</tr>
<tr>
<td><p><code>rngsubtype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>OID of the element type (subtype) of this range type</p></td>
</tr>
<tr>
<td><p><code>rngmultitypid</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>OID of the multirange type for this range type</p></td>
</tr>
<tr>
<td><p><code>rngcollation</code> <code>oid</code> (references <a href="pg_collation.md#catalog-pg-collation"><code>pg_collation</code></a>.<code>oid</code>)</p>
<p>OID of the collation used for range comparisons, or zero if none</p></td>
</tr>
<tr>
<td><p><code>rngsubopc</code> <code>oid</code> (references <a href="pg_opclass.md#catalog-pg-opclass"><code>pg_opclass</code></a>.<code>oid</code>)</p>
<p>OID of the subtype's operator class used for range comparisons</p></td>
</tr>
<tr>
<td><p><code>rngconstruct2</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the 2-argument range constructor function (lower and upper)</p></td>
</tr>
<tr>
<td><p><code>rngconstruct3</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the 3-argument range constructor function (lower, upper, and flags)</p></td>
</tr>
<tr>
<td><p><code>rngmltconstruct0</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the 0-argument multirange constructor function (constructs empty range)</p></td>
</tr>
<tr>
<td><p><code>rngmltconstruct1</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the 1-argument multirange constructor function (constructs multirange from single range, also used as cast function)</p></td>
</tr>
<tr>
<td><p><code>rngmltconstruct2</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the 2-argument multirange constructor function (constructs multirange from array of ranges)</p></td>
</tr>
<tr>
<td><p><code>rngcanonical</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the function to convert a range value into canonical form, or zero if none</p></td>
</tr>
<tr>
<td><p><code>rngsubdiff</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of the function to return the difference between two element values as <code>double precision</code>, or zero if none</p></td>
</tr>
</tbody>
</table>


 `rngsubopc` (plus `rngcollation`, if the element type is collatable) determines the sort ordering used by the range type. `rngcanonical` is used when the element type is discrete. `rngsubdiff` is optional but should be supplied to improve performance of GiST indexes on the range type.
