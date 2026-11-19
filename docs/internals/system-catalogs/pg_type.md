<a id="catalog-pg-type"></a>

## `pg_type`


 The catalog `pg_type` stores information about data types. Base types and enum types (scalar types) are created with [`CREATE TYPE`](../../reference/sql-commands/create-type.md#sql-createtype), and domains with [`CREATE DOMAIN`](../../reference/sql-commands/create-domain.md#sql-createdomain). A composite type is automatically created for each table in the database, to represent the row structure of the table. It is also possible to create composite types with `CREATE TYPE AS`.


**Table: `pg_type` Columns**

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
<td><p><code>typname</code> <code>name</code></p>
<p>Data type name</p></td>
</tr>
<tr>
<td><p><code>typnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this type</p></td>
</tr>
<tr>
<td><p><code>typowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the type</p></td>
</tr>
<tr>
<td><p><code>typlen</code> <code>int2</code></p>
<p>For a fixed-size type, <code>typlen</code> is the number of bytes in the internal representation of the type. But for a variable-length type, <code>typlen</code> is negative. -1 indicates a “varlena” type (one that has a length word), -2 indicates a null-terminated C string.</p></td>
</tr>
<tr>
<td><p><code>typbyval</code> <code>bool</code></p>
<p><code>typbyval</code> determines whether internal routines pass a value of this type by value or by reference. <code>typbyval</code> had better be false if <code>typlen</code> is not 1, 2, or 4 (or 8 on machines where Datum is 8 bytes). Variable-length types are always passed by reference. Note that <code>typbyval</code> can be false even if the length would allow pass-by-value.</p></td>
</tr>
<tr>
<td><p><code>typtype</code> <code>char</code></p>
<p><code>typtype</code> is <code>b</code> for a base type, <code>c</code> for a composite type (e.g., a table's row type), <code>d</code> for a domain, <code>e</code> for an enum type, <code>p</code> for a pseudo-type, <code>r</code> for a range type, or <code>m</code> for a multirange type. See also <code>typrelid</code> and <code>typbasetype</code>.</p></td>
</tr>
<tr>
<td><p><code>typcategory</code> <code>char</code></p>
<p><code>typcategory</code> is an arbitrary classification of data types that is used by the parser to determine which implicit casts should be “preferred”. See <a href="#catalog-typcategory-table"><code>typcategory</code> Codes</a>.</p></td>
</tr>
<tr>
<td><p><code>typispreferred</code> <code>bool</code></p>
<p>True if the type is a preferred cast target within its <code>typcategory</code></p></td>
</tr>
<tr>
<td><p><code>typisdefined</code> <code>bool</code></p>
<p>True if the type is defined, false if this is a placeholder entry for a not-yet-defined type. When <code>typisdefined</code> is false, nothing except the type name, namespace, and OID can be relied on.</p></td>
</tr>
<tr>
<td><p><code>typdelim</code> <code>char</code></p>
<p>Character that separates two values of this type when parsing array input. Note that the delimiter is associated with the array element data type, not the array data type.</p></td>
</tr>
<tr>
<td><p><code>typrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>If this is a composite type (see <code>typtype</code>), then this column points to the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry that defines the corresponding table. (For a free-standing composite type, the <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a> entry doesn't really represent a table, but it is needed anyway for the type's <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a> entries to link to.) Zero for non-composite types.</p></td>
</tr>
<tr>
<td><p><code>typsubscript</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Subscripting handler function's OID, or zero if this type doesn't support subscripting. Types that are “true” array types have <code>typsubscript</code> = <code>array_subscript_handler</code>, but other types may have other handler functions to implement specialized subscripting behavior.</p></td>
</tr>
<tr>
<td><p><code>typelem</code> <code>oid</code> (references <a href="#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>If <code>typelem</code> is not zero then it identifies another row in <code>pg_type</code>, defining the type yielded by subscripting. This should be zero if <code>typsubscript</code> is zero. However, it can be zero when <code>typsubscript</code> isn't zero, if the handler doesn't need <code>typelem</code> to determine the subscripting result type. Note that a <code>typelem</code> dependency is considered to imply physical containment of the element type in this type; so DDL changes on the element type might be restricted by the presence of this type.</p></td>
</tr>
<tr>
<td><p><code>typarray</code> <code>oid</code> (references <a href="#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>If <code>typarray</code> is not zero then it identifies another row in <code>pg_type</code>, which is the “true” array type having this type as element</p></td>
</tr>
<tr>
<td><p><code>typinput</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Input conversion function (text format)</p></td>
</tr>
<tr>
<td><p><code>typoutput</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Output conversion function (text format)</p></td>
</tr>
<tr>
<td><p><code>typreceive</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Input conversion function (binary format), or zero if none</p></td>
</tr>
<tr>
<td><p><code>typsend</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Output conversion function (binary format), or zero if none</p></td>
</tr>
<tr>
<td><p><code>typmodin</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Type modifier input function, or zero if type does not support modifiers</p></td>
</tr>
<tr>
<td><p><code>typmodout</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Type modifier output function, or zero to use the standard format</p></td>
</tr>
<tr>
<td><p><code>typanalyze</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Custom <a href="../../reference/sql-commands/analyze.md#sql-analyze">sql-analyze</a> function, or zero to use the standard function</p></td>
</tr>
<tr>
<td><p><code>typalign</code> <code>char</code></p>
<p><code>typalign</code> is the alignment required when storing a value of this type. It applies to storage on disk as well as most representations of the value inside PostgreSQL. When multiple values are stored consecutively, such as in the representation of a complete row on disk, padding is inserted before a datum of this type so that it begins on the specified boundary. The alignment reference is the beginning of the first datum in the sequence. Possible values are:</p>
<p>- <code>c</code> = <code>char</code> alignment, i.e., no alignment needed.<br>
- <code>s</code> = <code>short</code> alignment (2 bytes on most machines).<br>
- <code>i</code> = <code>int</code> alignment (4 bytes on most machines).<br>
- <code>d</code> = <code>double</code> alignment (8 bytes on many machines, but by no means all).</p></td>
</tr>
<tr>
<td><p><code>typstorage</code> <code>char</code></p>
<p><code>typstorage</code> tells for varlena types (those with <code>typlen</code> = -1) if the type is prepared for toasting and what the default strategy for attributes of this type should be. Possible values are:</p>
<p>-  <code>p</code> (plain): Values must always be stored plain (non-varlena types always use this value). <br>
-  <code>e</code> (external): Values can be stored in a secondary “TOAST” relation (if relation has one, see <code>pg_class.reltoastrelid</code>). <br>
-  <code>m</code> (main): Values can be compressed and stored inline. <br>
-  <code>x</code> (extended): Values can be compressed and/or moved to a secondary relation. <br>
 <code>x</code> is the usual choice for toast-able types. Note that <code>m</code> values can also be moved out to secondary storage, but only as a last resort (<code>e</code> and <code>x</code> values are moved first).</p></td>
</tr>
<tr>
<td><p><code>typnotnull</code> <code>bool</code></p>
<p><code>typnotnull</code> represents a not-null constraint on a type. Used for domains only.</p></td>
</tr>
<tr>
<td><p><code>typbasetype</code> <code>oid</code> (references <a href="#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>If this is a domain (see <code>typtype</code>), then <code>typbasetype</code> identifies the type that this one is based on. Zero if this type is not a domain.</p></td>
</tr>
<tr>
<td><p><code>typtypmod</code> <code>int4</code></p>
<p>Domains use <code>typtypmod</code> to record the <code>typmod</code> to be applied to their base type (-1 if base type does not use a <code>typmod</code>). -1 if this type is not a domain.</p></td>
</tr>
<tr>
<td><p><code>typndims</code> <code>int4</code></p>
<p><code>typndims</code> is the number of array dimensions for a domain over an array (that is, <code>typbasetype</code> is an array type). Zero for types other than domains over array types.</p></td>
</tr>
<tr>
<td><p><code>typcollation</code> <code>oid</code> (references <a href="pg_collation.md#catalog-pg-collation"><code>pg_collation</code></a>.<code>oid</code>)</p>
<p><code>typcollation</code> specifies the collation of the type. If the type does not support collations, this will be zero. A base type that supports collations will have a nonzero value here, typically <code>DEFAULT_COLLATION_OID</code>. A domain over a collatable type can have a collation OID different from its base type's, if one was specified for the domain.</p></td>
</tr>
<tr>
<td><p><code>typdefaultbin</code> <code>pg_node_tree</code></p>
<p>If <code>typdefaultbin</code> is not null, it is the <code>nodeToString()</code> representation of a default expression for the type. This is only used for domains.</p></td>
</tr>
<tr>
<td><p><code>typdefault</code> <code>text</code></p>
<p><code>typdefault</code> is null if the type has no associated default value. If <code>typdefaultbin</code> is not null, <code>typdefault</code> must contain a human-readable version of the default expression represented by <code>typdefaultbin</code>. If <code>typdefaultbin</code> is null and <code>typdefault</code> is not, then <code>typdefault</code> is the external representation of the type's default value, which can be fed to the type's input converter to produce a constant.</p></td>
</tr>
<tr>
<td><p><code>typacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>


!!! note

    For fixed-width types used in system tables, it is critical that the size and alignment defined in `pg_type` agree with the way that the compiler will lay out the column in a structure representing a table row.


 [`typcategory` Codes](#catalog-typcategory-table) lists the system-defined values of `typcategory`. Any future additions to this list will also be upper-case ASCII letters. All other ASCII characters are reserved for user-defined categories.
 <a id="catalog-typcategory-table"></a>

**Table: `typcategory` Codes**

| Code | Category |
| --- | --- |
| `A` | Array types |
| `B` | Boolean types |
| `C` | Composite types |
| `D` | Date/time types |
| `E` | Enum types |
| `G` | Geometric types |
| `I` | Network address types |
| `N` | Numeric types |
| `P` | Pseudo-types |
| `R` | Range types |
| `S` | String types |
| `T` | Timespan types |
| `U` | User-defined types |
| `V` | Bit-string types |
| `X` | `unknown` type |
| `Z` | Internal-use types |
