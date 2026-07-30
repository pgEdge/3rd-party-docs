<a id="catalog-pg-database"></a>

## `pg_database`


 The catalog `pg_database` stores information about the available databases. Databases are created with the [`CREATE DATABASE`](../../reference/sql-commands/create-database.md#sql-createdatabase) command. Consult [Managing Databases](../../server-administration/managing-databases/index.md#managing-databases) for details about the meaning of some of the parameters.


 Unlike most system catalogs, `pg_database` is shared across all databases of a cluster: there is only one copy of `pg_database` per cluster, not one per database.


**Table: `pg_database` Columns**

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
<td><p><code>datname</code> <code>name</code></p>
<p>Database name</p></td>
</tr>
<tr>
<td><p><code>datdba</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the database, usually the user who created it</p></td>
</tr>
<tr>
<td><p><code>encoding</code> <code>int4</code></p>
<p>Character encoding for this database (<a href="../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#pg-encoding-to-char"><code>pg_encoding_to_char()</code></a> can translate this number to the encoding name)</p></td>
</tr>
<tr>
<td><p><code>datlocprovider</code> <code>char</code></p>
<p>Locale provider for this database: <code>b</code> = builtin, <code>c</code> = libc, <code>i</code> = icu</p></td>
</tr>
<tr>
<td><p><code>datistemplate</code> <code>bool</code></p>
<p>If true, then this database can be cloned by any user with <code>CREATEDB</code> privileges; if false, then only superusers or the owner of the database can clone it.</p></td>
</tr>
<tr>
<td><p><code>datallowconn</code> <code>bool</code></p>
<p>If false then no one can connect to this database. This is used to protect the <code>template0</code> database from being altered.</p></td>
</tr>
<tr>
<td><p><code>dathasloginevt</code> <code>bool</code></p>
<p>Indicates that there are login event triggers defined for this database. This flag is used to avoid extra lookups on the <code>pg_event_trigger</code> table during each backend startup. This flag is used internally by PostgreSQL and should not be manually altered or read for monitoring purposes.</p></td>
</tr>
<tr>
<td><p><code>datconnlimit</code> <code>int4</code></p>
<p>Sets maximum number of concurrent connections that can be made to this database. -1 means no limit, -2 indicates the database is invalid.</p></td>
</tr>
<tr>
<td><p><code>datfrozenxid</code> <code>xid</code></p>
<p>All transaction IDs before this one have been replaced with a permanent (“frozen”) transaction ID in this database. This is used to track whether the database needs to be vacuumed in order to prevent transaction ID wraparound or to allow <code>pg_xact</code> to be shrunk. It is the minimum of the per-table <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relfrozenxid</code> values.</p></td>
</tr>
<tr>
<td><p><code>datminmxid</code> <code>xid</code></p>
<p>All multixact IDs before this one have been replaced with a transaction ID in this database. This is used to track whether the database needs to be vacuumed in order to prevent multixact ID wraparound or to allow <code>pg_multixact</code> to be shrunk. It is the minimum of the per-table <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relminmxid</code> values.</p></td>
</tr>
<tr>
<td><p><code>dattablespace</code> <code>oid</code> (references <a href="pg_tablespace.md#catalog-pg-tablespace"><code>pg_tablespace</code></a>.<code>oid</code>)</p>
<p>The default tablespace for the database. Within this database, all tables for which <a href="pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>reltablespace</code> is zero will be stored in this tablespace; in particular, all the non-shared system catalogs will be there.</p></td>
</tr>
<tr>
<td><p><code>datcollate</code> <code>text</code></p>
<p>LC_COLLATE for this database (ignored unless <code>datlocprovider</code> is <code>c</code>)</p></td>
</tr>
<tr>
<td><p><code>datctype</code> <code>text</code></p>
<p>LC_CTYPE for this database</p></td>
</tr>
<tr>
<td><p><code>datlocale</code> <code>text</code></p>
<p>Collation provider locale name for this database. If the provider is <code>libc</code>, <code>datlocale</code> is <code>NULL</code>; <code>datcollate</code> and <code>datctype</code> are used instead.</p></td>
</tr>
<tr>
<td><p><code>daticurules</code> <code>text</code></p>
<p>ICU collation rules for this database</p></td>
</tr>
<tr>
<td><p><code>datcollversion</code> <code>text</code></p>
<p>Provider-specific version of the collation. This is recorded when the database is created and then checked when it is used, to detect changes in the collation definition that could lead to data corruption.</p></td>
</tr>
<tr>
<td><p><code>datacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>
