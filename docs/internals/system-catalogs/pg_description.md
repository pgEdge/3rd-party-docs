<a id="catalog-pg-description"></a>

## `pg_description`


 The catalog `pg_description` stores optional descriptions (comments) for each database object. Descriptions can be manipulated with the [`COMMENT`](../../reference/sql-commands/comment.md#sql-comment) command and viewed with psql's `\d` commands. Descriptions of many built-in system objects are provided in the initial contents of `pg_description`.


 See also [`pg_shdescription`](pg_shdescription.md#catalog-pg-shdescription), which performs a similar function for descriptions involving objects that are shared across a database cluster.


**Table: `pg_description` Columns**

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
<p>The OID of the object this description pertains to</p></td>
</tr>
<tr>
<td><p><code>classoid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The OID of the system catalog this object appears in</p></td>
</tr>
<tr>
<td><p><code>objsubid</code> <code>int4</code></p>
<p>For a comment on a table column, this is the column number (the <code>objoid</code> and <code>classoid</code> refer to the table itself). For all other object types, this column is zero.</p></td>
</tr>
<tr>
<td><p><code>description</code> <code>text</code></p>
<p>Arbitrary text that serves as the description of this object</p></td>
</tr>
</tbody>
</table>
