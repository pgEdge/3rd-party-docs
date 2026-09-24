<a id="catalog-pg-enum"></a>

## `pg_enum`


 The `pg_enum` catalog contains entries showing the values and labels for each enum type. The internal representation of a given enum value is actually the OID of its associated row in `pg_enum`.


**Table: `pg_enum` Columns**

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
<td><p><code>enumtypid</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>The OID of the <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a> entry owning this enum value</p></td>
</tr>
<tr>
<td><p><code>enumsortorder</code> <code>float4</code></p>
<p>The sort position of this enum value within its enum type</p></td>
</tr>
<tr>
<td><p><code>enumlabel</code> <code>name</code></p>
<p>The textual label for this enum value</p></td>
</tr>
</tbody>
</table>


 The OIDs for `pg_enum` rows follow a special rule: even-numbered OIDs are guaranteed to be ordered in the same way as the sort ordering of their enum type. That is, if two even OIDs belong to the same enum type, the smaller OID must have the smaller `enumsortorder` value. Odd-numbered OID values need bear no relationship to the sort order. This rule allows the enum comparison routines to avoid catalog lookups in many common cases. The routines that create and alter enum types attempt to assign even OIDs to enum values whenever possible.


 When an enum type is created, its members are assigned sort-order positions 1..*n*. But members added later might be given negative or fractional values of `enumsortorder`. The only requirement on these values is that they be correctly ordered and unique within each enum type.
