<a id="catalog-pg-constraint"></a>

## `pg_constraint`


 The catalog `pg_constraint` stores check, not-null, primary key, unique, foreign key, and exclusion constraints on tables. (Column constraints are not treated specially. Every column constraint is equivalent to some table constraint.)


 User-defined constraint triggers (created with [`CREATE CONSTRAINT TRIGGER`](../../reference/sql-commands/create-trigger.md#sql-createtrigger)) also give rise to an entry in this table.


 Check constraints on domains are stored here, too.


**Table: `pg_constraint` Columns**

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
<td><p><code>conname</code> <code>name</code></p>
<p>Constraint name (not necessarily unique!)</p></td>
</tr>
<tr>
<td><p><code>connamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this constraint</p></td>
</tr>
<tr>
<td><p><code>contype</code> <code>char</code></p>
<p><code>c</code> = check constraint, <code>f</code> = foreign key constraint, <code>n</code> = not-null constraint, <code>p</code> = primary key constraint, <code>u</code> = unique constraint, <code>t</code> = constraint trigger, <code>x</code> = exclusion constraint</p></td>
</tr>
<tr>
<td><p><code>condeferrable</code> <code>bool</code></p>
<p>Is the constraint deferrable?</p></td>
</tr>
<tr>
<td><p><code>condeferred</code> <code>bool</code></p>
<p>Is the constraint deferred by default?</p></td>
</tr>
<tr>
<td><p><code>conenforced</code> <code>bool</code></p>
<p>Is the constraint enforced?</p></td>
</tr>
<tr>
<td><p><code>convalidated</code> <code>bool</code></p>
<p>Has the constraint been validated?</p></td>
</tr>
<tr>
<td><p><code>conrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The table this constraint is on; zero if not a table constraint</p></td>
</tr>
<tr>
<td><p><code>contypid</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>The domain this constraint is on; zero if not a domain constraint</p></td>
</tr>
<tr>
<td><p><code>conindid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>The index supporting this constraint, if it's a unique, primary key, foreign key, or exclusion constraint; else zero</p></td>
</tr>
<tr>
<td><p><code>conparentid</code> <code>oid</code> (references <a href="#catalog-pg-constraint"><code>pg_constraint</code></a>.<code>oid</code>)</p>
<p>The corresponding constraint of the parent partitioned table, if this is a constraint on a partition; else zero</p></td>
</tr>
<tr>
<td><p><code>confrelid</code> <code>oid</code> (references <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>If a foreign key, the referenced table; else zero</p></td>
</tr>
<tr>
<td><p><code>confupdtype</code> <code>char</code></p>
<p>Foreign key update action code: <code>a</code> = no action, <code>r</code> = restrict, <code>c</code> = cascade, <code>n</code> = set null, <code>d</code> = set default</p></td>
</tr>
<tr>
<td><p><code>confdeltype</code> <code>char</code></p>
<p>Foreign key deletion action code: <code>a</code> = no action, <code>r</code> = restrict, <code>c</code> = cascade, <code>n</code> = set null, <code>d</code> = set default</p></td>
</tr>
<tr>
<td><p><code>confmatchtype</code> <code>char</code></p>
<p>Foreign key match type: <code>f</code> = full, <code>p</code> = partial, <code>s</code> = simple</p></td>
</tr>
<tr>
<td><p><code>conislocal</code> <code>bool</code></p>
<p>This constraint is defined locally for the relation. Note that a constraint can be locally defined and inherited simultaneously.</p></td>
</tr>
<tr>
<td><p><code>coninhcount</code> <code>int2</code></p>
<p>The number of direct inheritance ancestors this constraint has. A constraint with a nonzero number of ancestors cannot be dropped nor renamed.</p></td>
</tr>
<tr>
<td><p><code>connoinherit</code> <code>bool</code></p>
<p>This constraint is defined locally for the relation. It is a non-inheritable constraint.</p></td>
</tr>
<tr>
<td><p><code>conperiod</code> <code>bool</code></p>
<p>This constraint is defined with <code>WITHOUT OVERLAPS</code> (for primary keys and unique constraints) or <code>PERIOD</code> (for foreign keys).</p></td>
</tr>
<tr>
<td><p><code>conkey</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>If a table constraint (including foreign keys, but not constraint triggers), list of the constrained columns</p></td>
</tr>
<tr>
<td><p><code>confkey</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>If a foreign key, list of the referenced columns</p></td>
</tr>
<tr>
<td><p><code>conpfeqop</code> <code>oid[]</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>If a foreign key, list of the equality operators for PK = FK comparisons</p></td>
</tr>
<tr>
<td><p><code>conppeqop</code> <code>oid[]</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>If a foreign key, list of the equality operators for PK = PK comparisons</p></td>
</tr>
<tr>
<td><p><code>conffeqop</code> <code>oid[]</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>If a foreign key, list of the equality operators for FK = FK comparisons</p></td>
</tr>
<tr>
<td><p><code>confdelsetcols</code> <code>int2[]</code> (references <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>.<code>attnum</code>)</p>
<p>If a foreign key with a <code>SET NULL</code> or <code>SET DEFAULT</code> delete action, the columns that will be updated. If null, all of the referencing columns will be updated.</p></td>
</tr>
<tr>
<td><p><code>conexclop</code> <code>oid[]</code> (references <a href="pg_operator.md#catalog-pg-operator"><code>pg_operator</code></a>.<code>oid</code>)</p>
<p>If an exclusion constraint or <code>WITHOUT OVERLAPS</code> primary key/unique constraint, list of the per-column exclusion operators.</p></td>
</tr>
<tr>
<td><p><code>conbin</code> <code>pg_node_tree</code></p>
<p>If a check constraint, an internal representation of the expression. (It's recommended to use <code>pg_get_constraintdef()</code> to extract the definition of a check constraint.)</p></td>
</tr>
</tbody>
</table>


 In the case of an exclusion constraint, `conkey` is only useful for constraint elements that are simple column references. For other cases, a zero appears in `conkey` and the associated index must be consulted to discover the expression that is constrained. (`conkey` thus has the same contents as [`pg_index`](pg_index.md#catalog-pg-index).`indkey` for the index.)


!!! note

    `pg_class.relchecks` needs to agree with the number of check-constraint entries found in this table for each relation.
