<a id="catalog-pg-operator"></a>

## `pg_operator`


 The catalog `pg_operator` stores information about operators. See [sql-createoperator](../../reference/sql-commands/create-operator.md#sql-createoperator) and [User-Defined Operators](../../server-programming/extending-sql/user-defined-operators.md#xoper) for more information.


**Table: `pg_operator` Columns**

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
<td><p><code>oprname</code> <code>name</code></p>
<p>Name of the operator</p></td>
</tr>
<tr>
<td><p><code>oprnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this operator</p></td>
</tr>
<tr>
<td><p><code>oprowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the operator</p></td>
</tr>
<tr>
<td><p><code>oprkind</code> <code>char</code></p>
<p><code>b</code> = infix operator (“both”), or <code>l</code> = prefix operator (“left”)</p></td>
</tr>
<tr>
<td><p><code>oprcanmerge</code> <code>bool</code></p>
<p>This operator supports merge joins</p></td>
</tr>
<tr>
<td><p><code>oprcanhash</code> <code>bool</code></p>
<p>This operator supports hash joins</p></td>
</tr>
<tr>
<td><p><code>oprleft</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Type of the left operand (zero for a prefix operator)</p></td>
</tr>
<tr>
<td><p><code>oprright</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Type of the right operand</p></td>
</tr>
<tr>
<td><p><code>oprresult</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Type of the result (zero for a not-yet-defined “shell” operator)</p></td>
</tr>
<tr>
<td><p><code>oprcom</code> <code>oid</code> (references <a href="#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>Commutator of this operator (zero if none)</p></td>
</tr>
<tr>
<td><p><code>oprnegate</code> <code>oid</code> (references <a href="#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>Negator of this operator (zero if none)</p></td>
</tr>
<tr>
<td><p><code>oprcode</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Function that implements this operator (zero for a not-yet-defined “shell” operator)</p></td>
</tr>
<tr>
<td><p><code>oprrest</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Restriction selectivity estimation function for this operator (zero if none)</p></td>
</tr>
<tr>
<td><p><code>oprjoin</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Join selectivity estimation function for this operator (zero if none)</p></td>
</tr>
</tbody>
</table>
