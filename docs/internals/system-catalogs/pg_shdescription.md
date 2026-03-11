<a id="catalog-pg-shdescription"></a>

## `pg_shdescription`


 The catalog `pg_shdescription` stores optional descriptions (comments) for shared database objects. Descriptions can be manipulated with the [`COMMENT`](../../reference/sql-commands/comment.md#sql-comment) command and viewed with psql's `\d` commands.


 See also [`pg_description`](pg_description.md#catalog-pg-description), which performs a similar function for descriptions involving objects within a single database.


 Unlike most system catalogs, `pg_shdescription` is shared across all databases of a cluster: there is only one copy of `pg_shdescription` per cluster, not one per database.


**Table: `pg_shdescription` Columns**

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
<td><p><code>description</code> <code>text</code></p>
<p>Arbitrary text that serves as the description of this object</p></td>
</tr>
</tbody>
</table>
