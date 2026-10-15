<a id="catalog-pg-class"></a>

## `pg_class`


 The catalog `pg_class` describes tables and other objects that have columns or are otherwise similar to a table. This includes indexes (but see also [`pg_index`](pg_index.md#catalog-pg-index)), sequences (but see also [`pg_sequence`](pg_sequence.md#catalog-pg-sequence)), views, materialized views, composite types, and TOAST tables; see `relkind`. Below, when we mean all of these kinds of objects we speak of “relations”. Not all of `pg_class`'s columns are meaningful for all relation kinds.


**Table: `pg_class` Columns**

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
<td><p><code>relname</code> <code>name</code></p>
<p>Name of the table, index, view, etc.</p></td>
</tr>
<tr>
<td><p><code>relnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this relation</p></td>
</tr>
<tr>
<td><p><code>reltype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>The OID of the data type that corresponds to this table's row type, if any; zero for indexes, sequences, and TOAST tables, which have no <code>pg_type</code> entry</p></td>
</tr>
<tr>
<td><p><code>reloftype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>For typed tables, the OID of the underlying composite type; zero for all other relations</p></td>
</tr>
<tr>
<td><p><code>relowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the relation</p></td>
</tr>
<tr>
<td><p><code>relam</code> <code>oid</code> (references <a href="pg_am.md#catalog-pg-am"><code>pg_am</code></a>.<code>oid</code>)</p>
<p>The access method used to access this table or index. Not meaningful if the relation is a sequence or has no on-disk file, except for partitioned tables, where, if set, it takes precedence over <code>default_table_access_method</code> when determining the access method to use for partitions created when one is not specified in the creation command.</p></td>
</tr>
<tr>
<td><p><code>relfilenode</code> <code>oid</code></p>
<p>Name of the on-disk file of this relation; zero means this is a “mapped” relation whose disk file name is determined by low-level state</p></td>
</tr>
<tr>
<td><p><code>reltablespace</code> <code>oid</code> (references <a href="pg_tablespace.md#catalog-pg-tablespace"><code>pg_tablespace</code></a>.<code>oid</code>)</p>
<p>The tablespace in which this relation is stored. If zero, the database's default tablespace is implied. Not meaningful if the relation has no on-disk file, except for partitioned tables, where this is the tablespace in which partitions will be created when one is not specified in the creation command.</p></td>
</tr>
<tr>
<td><p><code>relpages</code> <code>int4</code></p>
<p>Size of the on-disk representation of this table in pages (of size <code>BLCKSZ</code>). This is only an estimate used by the planner. It is updated by <a href="../../reference/sql-commands/vacuum.md#sql-vacuum"><code>VACUUM</code></a>, <a href="../../reference/sql-commands/analyze.md#sql-analyze"><code>ANALYZE</code></a>, and a few DDL commands such as <a href="../../reference/sql-commands/create-index.md#sql-createindex"><code>CREATE INDEX</code></a>.</p></td>
</tr>
<tr>
<td><p><code>reltuples</code> <code>float4</code></p>
<p>Number of live rows in the table. This is only an estimate used by the planner. It is updated by <a href="../../reference/sql-commands/vacuum.md#sql-vacuum"><code>VACUUM</code></a>, <a href="../../reference/sql-commands/analyze.md#sql-analyze"><code>ANALYZE</code></a>, and a few DDL commands such as <a href="../../reference/sql-commands/create-index.md#sql-createindex"><code>CREATE INDEX</code></a>. If the table has never yet been vacuumed or analyzed, <code>reltuples</code> contains <code>-1</code> indicating that the row count is unknown.</p></td>
</tr>
<tr>
<td><p><code>relallvisible</code> <code>int4</code></p>
<p>Number of pages that are marked all-visible in the table's visibility map. This is only an estimate used by the planner. It is updated by <a href="../../reference/sql-commands/vacuum.md#sql-vacuum"><code>VACUUM</code></a>, <a href="../../reference/sql-commands/analyze.md#sql-analyze"><code>ANALYZE</code></a>, and a few DDL commands such as <a href="../../reference/sql-commands/create-index.md#sql-createindex"><code>CREATE INDEX</code></a>.</p></td>
</tr>
<tr>
<td><p><code>relallfrozen</code> <code>int4</code></p>
<p>Number of pages that are marked all-frozen in the table's visibility map. This is only an estimate used for triggering autovacuums. It can also be used along with <code>relallvisible</code> for scheduling manual vacuums and tuning <a href="../../server-administration/server-configuration/vacuuming.md#runtime-config-vacuum-freezing">vacuum's freezing behavior</a>. It is updated by <a href="../../reference/sql-commands/vacuum.md#sql-vacuum"><code>VACUUM</code></a>, <a href="../../reference/sql-commands/analyze.md#sql-analyze"><code>ANALYZE</code></a>, and a few DDL commands such as <a href="../../reference/sql-commands/create-index.md#sql-createindex"><code>CREATE INDEX</code></a>.</p></td>
</tr>
<tr>
<td><p><code>reltoastrelid</code> <code>oid</code> (references <a href="#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>OID of the TOAST table associated with this table, zero if none. The TOAST table stores large attributes “out of line” in a secondary table.</p></td>
</tr>
<tr>
<td><p><code>relhasindex</code> <code>bool</code></p>
<p>True if this is a table and it has (or recently had) any indexes</p></td>
</tr>
<tr>
<td><p><code>relisshared</code> <code>bool</code></p>
<p>True if this table is shared across all databases in the cluster. Only certain system catalogs (such as <a href="pg_database.md#catalog-pg-database"><code>pg_database</code></a>) are shared.</p></td>
</tr>
<tr>
<td><p><code>relpersistence</code> <code>char</code></p>
<p><code>p</code> = permanent table/sequence, <code>u</code> = unlogged table/sequence, <code>t</code> = temporary table/sequence</p></td>
</tr>
<tr>
<td><p><code>relkind</code> <code>char</code></p>
<p><code>r</code> = ordinary table, <code>i</code> = index, <code>S</code> = sequence, <code>t</code> = TOAST table, <code>v</code> = view, <code>m</code> = materialized view, <code>c</code> = composite type, <code>f</code> = foreign table, <code>p</code> = partitioned table, <code>I</code> = partitioned index</p></td>
</tr>
<tr>
<td><p><code>relnatts</code> <code>int2</code></p>
<p>Number of user columns in the relation (system columns not counted). There must be this many corresponding entries in <a href="pg_attribute.md#catalog-pg-attribute"><code>pg_attribute</code></a>. See also <code>pg_attribute</code>.<code>attnum</code>.</p></td>
</tr>
<tr>
<td><p><code>relchecks</code> <code>int2</code></p>
<p>Number of <code>CHECK</code> constraints on the table; see <a href="pg_constraint.md#catalog-pg-constraint"><code>pg_constraint</code></a> catalog</p></td>
</tr>
<tr>
<td><p><code>relhasrules</code> <code>bool</code></p>
<p>True if table has (or once had) rules; see <a href="pg_rewrite.md#catalog-pg-rewrite"><code>pg_rewrite</code></a> catalog</p></td>
</tr>
<tr>
<td><p><code>relhastriggers</code> <code>bool</code></p>
<p>True if table has (or once had) triggers; see <a href="pg_trigger.md#catalog-pg-trigger"><code>pg_trigger</code></a> catalog</p></td>
</tr>
<tr>
<td><p><code>relhassubclass</code> <code>bool</code></p>
<p>True if table or index has (or once had) any inheritance children or partitions</p></td>
</tr>
<tr>
<td><p><code>relrowsecurity</code> <code>bool</code></p>
<p>True if table has row-level security enabled; see <a href="pg_policy.md#catalog-pg-policy"><code>pg_policy</code></a> catalog</p></td>
</tr>
<tr>
<td><p><code>relforcerowsecurity</code> <code>bool</code></p>
<p>True if row-level security (when enabled) will also apply to table owner; see <a href="pg_policy.md#catalog-pg-policy"><code>pg_policy</code></a> catalog</p></td>
</tr>
<tr>
<td><p><code>relispopulated</code> <code>bool</code></p>
<p>True if relation is populated (this is true for all relations other than some materialized views)</p></td>
</tr>
<tr>
<td><p><code>relreplident</code> <code>char</code></p>
<p>Columns used to form “replica identity” for rows: <code>d</code> = default (primary key, if any), <code>n</code> = nothing, <code>f</code> = all columns, <code>i</code> = index with <code>indisreplident</code> set (same as nothing if the index used has been dropped)</p></td>
</tr>
<tr>
<td><p><code>relispartition</code> <code>bool</code></p>
<p>True if table or index is a partition</p></td>
</tr>
<tr>
<td><p><code>relrewrite</code> <code>oid</code> (references <a href="#catalog-pg-class"><code>pg_class</code></a>.<code>oid</code>)</p>
<p>For new relations being written during a DDL operation that requires a table rewrite, this contains the OID of the original relation; otherwise zero. That state is only visible internally; this field should never contain anything other than zero for a user-visible relation.</p></td>
</tr>
<tr>
<td><p><code>relfrozenxid</code> <code>xid</code></p>
<p>All transaction IDs before this one have been replaced with a permanent (“frozen”) transaction ID in this table. This is used to track whether the table needs to be vacuumed in order to prevent transaction ID wraparound or to allow <code>pg_xact</code> to be shrunk. Zero (<code>InvalidTransactionId</code>) if the relation is not a table.</p></td>
</tr>
<tr>
<td><p><code>relminmxid</code> <code>xid</code></p>
<p>All multixact IDs before this one have been replaced by a transaction ID in this table. This is used to track whether the table needs to be vacuumed in order to prevent multixact ID wraparound or to allow <code>pg_multixact</code> to be shrunk. Zero (<code>InvalidMultiXactId</code>) if the relation is not a table.</p></td>
</tr>
<tr>
<td><p><code>relacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
<tr>
<td><p><code>reloptions</code> <code>text[]</code></p>
<p>Access-method-specific options, as “keyword=value” strings</p></td>
</tr>
<tr>
<td><p><code>relpartbound</code> <code>pg_node_tree</code></p>
<p>If table is a partition (see <code>relispartition</code>), internal representation of the partition bound</p></td>
</tr>
</tbody>
</table>


 Several of the Boolean flags in `pg_class` are maintained lazily: they are guaranteed to be true if that's the correct state, but may not be reset to false immediately when the condition is no longer true. For example, `relhasindex` is set by [`CREATE INDEX`](../../reference/sql-commands/create-index.md#sql-createindex), but it is never cleared by [`DROP INDEX`](../../reference/sql-commands/drop-index.md#sql-dropindex). Instead, [`VACUUM`](../../reference/sql-commands/vacuum.md#sql-vacuum) clears `relhasindex` if it finds the table has no indexes. This arrangement avoids race conditions and improves concurrency.
